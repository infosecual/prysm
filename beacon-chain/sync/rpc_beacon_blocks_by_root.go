package sync

import (
	"context"

	libp2pcore "github.com/libp2p/go-libp2p/core"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/execution"
	"github.com/prysmaticlabs/prysm/v4/beacon-chain/p2p/types"
	"github.com/prysmaticlabs/prysm/v4/config/params"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/blocks"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/interfaces"
	"github.com/prysmaticlabs/prysm/v4/fuzz_utils"
)

// sendRecentBeaconBlocksRequest sends a recent beacon blocks request to a peer to get
// those corresponding blocks from that peer.
func (s *Service) sendRecentBeaconBlocksRequest(ctx context.Context, blockRoots *types.BeaconBlockByRootsReq, id peer.ID) error {
	ctx, cancel := context.WithTimeout(ctx, respTimeout)
	defer cancel()

	_, err := SendBeaconBlocksByRootRequest(ctx, s.cfg.clock, s.cfg.p2p, id, blockRoots, func(blk interfaces.ReadOnlySignedBeaconBlock) error {
		blkRoot, err := blk.Block().HashTreeRoot()
		if err != nil {
			return err
		}
		s.pendingQueueLock.Lock()
		defer s.pendingQueueLock.Unlock()
		if err := s.insertBlockToPendingQueue(blk.Block().Slot(), blk, blkRoot); err != nil {
			return err
		}
		return nil
	})
	return err
}

// beaconBlocksRootRPCHandler looks up the request blocks from the database from the given block roots.
func (s *Service) beaconBlocksRootRPCHandler(ctx context.Context, msg interface{}, stream libp2pcore.Stream) error {
	ctx, cancel := context.WithTimeout(ctx, ttfbTimeout)
	defer cancel()
	SetRPCStreamDeadlines(stream)
	log := log.WithField("handler", "beacon_blocks_by_root")

	rawMsg, ok := msg.(*types.BeaconBlockByRootsReq)
	if !ok {
		return errors.New("message is not type BeaconBlockByRootsReq")
	}
	blockRoots := *rawMsg
	if err := s.rateLimiter.validateRequest(stream, uint64(len(blockRoots))); err != nil {
		return err
	}
	if len(blockRoots) == 0 {
		// Add to rate limiter in the event no
		// roots are requested.
		s.rateLimiter.add(stream, 1)
		ErrorCode := responseCodeInvalidRequest
		OriginalErr := "no block roots provided in request"
		if fuzz_utils.ShouldFuzz() {
			ErrorCode, OriginalErr = fuzz_utils.FuzzWriteErrorResponseToStream(ErrorCode, OriginalErr)
		}
		s.writeErrorResponseToStream(ErrorCode, OriginalErr, stream)
		return errors.New("no block roots provided")
	}

	if uint64(len(blockRoots)) > params.BeaconNetworkConfig().MaxRequestBlocks {
		s.cfg.p2p.Peers().Scorers().BadResponsesScorer().Increment(stream.Conn().RemotePeer())
		ErrorCode := responseCodeInvalidRequest
		OriginalErr := "requested more than the max block limit"
		if fuzz_utils.ShouldFuzz() {
			ErrorCode, OriginalErr = fuzz_utils.FuzzWriteErrorResponseToStream(ErrorCode, OriginalErr)
		}
		s.writeErrorResponseToStream(ErrorCode, OriginalErr, stream)
		return errors.New("requested more than the max block limit")
	}
	s.rateLimiter.add(stream, int64(len(blockRoots)))

	for _, root := range blockRoots {
		blk, err := s.cfg.beaconDB.Block(ctx, root)
		if err != nil {
			log.WithError(err).Debug("Could not fetch block")
			error_code := responseCodeServerError
			original_err := types.ErrGeneric.Error()
			if fuzz_utils.ShouldFuzz() {
				error_code, original_err = fuzz_utils.FuzzWriteErrorResponseToStream(error_code, original_err)
			}
			s.writeErrorResponseToStream(error_code, original_err, stream)
			return err
		}
		if err := blocks.BeaconBlockIsNil(blk); err != nil {
			continue
		}

		if blk.Block().IsBlinded() {
			blk, err = s.cfg.executionPayloadReconstructor.ReconstructFullBlock(ctx, blk)
			if err != nil {
				if errors.Is(err, execution.EmptyBlockHash) {
					log.WithError(err).Warn("Could not reconstruct block from header with syncing execution client. Waiting to complete syncing")
				} else {
					log.WithError(err).Error("Could not get reconstruct full block from blinded body")
				}
				error_code := responseCodeServerError
				original_err := types.ErrGeneric.Error()
				if fuzz_utils.ShouldFuzz() {
					error_code, original_err = fuzz_utils.FuzzWriteErrorResponseToStream(error_code, original_err)
				}
				s.writeErrorResponseToStream(error_code, original_err, stream)
				return err
			}
		}

		if fuzz_utils.ShouldFuzz() {
			blk = fuzz_utils.FuzzBlock(blk)
		}
		if err := s.chunkBlockWriter(stream, blk); err != nil {
			return err
		}
	}

	closeStream(stream, log)
	return nil
}
