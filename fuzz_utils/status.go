package fuzz_utils

import (
	"github.com/ethereum/go-ethereum/log"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
)

// /////////////////////////////////////////////////////////////////////////////
// status fuzzers
// /////////////////////////////////////////////////////////////////////////////
func FuzzStatusResponse(forkDigest *[]byte, finalizedRoot *[]byte, finalizedEpoch *primitives.Epoch, headRoot *[]byte, HeadSlot *primitives.Slot) {
	// status: *StatusResponse
	log.Info("FUZZ - FuzzStatusResponse")
	return
}
