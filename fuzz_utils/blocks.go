package fuzz_utils

import (
	"github.com/prysmaticlabs/prysm/v4/consensus-types/interfaces"
)

// /////////////////////////////////////////////////////////////////////////////
// block fuzzers
// /////////////////////////////////////////////////////////////////////////////

func FuzzBlock(blk interfaces.ReadOnlySignedBeaconBlock) interfaces.ReadOnlySignedBeaconBlock {
	// blk: interfaces.ReadOnlySignedBeaconBlock
	// from spec
	//class BeaconBlock(Container):
	//	slot: Slot
	//	proposer_index: ValidatorIndex
	//	parent_root: Root
	//	state_root: Root
	//	body: BeaconBlockBody
	//class BeaconBlockBody(Container):
	//	randao_reveal: BLSSignature
	//	eth1_data: Eth1Data  # Eth1 data vote
	//	graffiti: Bytes32  # Arbitrary data
	//	# Operations
	//	proposer_slashings: List[ProposerSlashing, MAX_PROPOSER_SLASHINGS]
	//	attester_slashings: List[AttesterSlashing, MAX_ATTESTER_SLASHINGS]
	//	attestations: List[Attestation, MAX_ATTESTATIONS]
	//	deposits: List[Deposit, MAX_DEPOSITS]
	//	voluntary_exits: List[SignedVoluntaryExit, MAX_VOLUNTARY_EXITS]

	log.Info("FUZZ - FuzzBlock")
	// print all fields in blk
	unsigned_blk, err := blk.Block().Copy()
	if err != nil {
		log.Info("FUZZ - FuzzBlock ReadOnlySignedBeaconBlock.Copy() returned error: ", err)
	}
	log.Info("FUZZ - FuzzBlock blk: ", unsigned_blk)
	// print all fields of blk
	log.Info("FUZZ - FuzzBlock blk.Slot(): ", unsigned_blk.Slot())
	log.Info("FUZZ - FuzzBlock blk.ProposerIndex(): ", unsigned_blk.ProposerIndex())
	log.Info("FUZZ - FuzzBlock blk.ParentRoot(): ", unsigned_blk.ParentRoot())
	log.Info("FUZZ - FuzzBlock blk.StateRoot(): ", unsigned_blk.StateRoot())
	log.Info("FUZZ - FuzzBlock blk.Body(): ", unsigned_blk.Body())
	// print all fields of blk.Body()
	log.Info("FUZZ - FuzzBlock blk.Body().RandaoReveal(): ", unsigned_blk.Body().RandaoReveal())
	log.Info("FUZZ - FuzzBlock blk.Body().Eth1Data(): ", unsigned_blk.Body().Eth1Data())
	log.Info("FUZZ - FuzzBlock blk.Body().Graffiti(): ", unsigned_blk.Body().Graffiti())
	log.Info("FUZZ - FuzzBlock blk.Body().ProposerSlashings(): ", unsigned_blk.Body().ProposerSlashings())
	log.Info("FUZZ - FuzzBlock blk.Body().AttesterSlashings(): ", unsigned_blk.Body().AttesterSlashings())
	log.Info("FUZZ - FuzzBlock blk.Body().Attestations(): ", unsigned_blk.Body().Attestations())
	log.Info("FUZZ - FuzzBlock blk.Body().Deposits(): ", unsigned_blk.Body().Deposits())
	log.Info("FUZZ - FuzzBlock blk.Body().VoluntaryExits(): ", unsigned_blk.Body().VoluntaryExits())

	return blk
}
