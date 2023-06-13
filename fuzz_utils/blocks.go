package fuzz_utils

import (
	"github.com/prysmaticlabs/prysm/v4/consensus-types/interfaces"
)

// /////////////////////////////////////////////////////////////////////////////
// block fuzzers
// /////////////////////////////////////////////////////////////////////////////

func FuzzBlock(blk interfaces.ReadOnlySignedBeaconBlock) interfaces.ReadOnlySignedBeaconBlock {
	// blk: interfaces.ReadOnlySignedBeaconBlock
	log.Info("FUZZ - FuzzBlock")
	return blk
}
