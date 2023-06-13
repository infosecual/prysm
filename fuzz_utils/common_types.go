package fuzz_utils

import (
	"math/rand"
)

// https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/p2p-interface.md#responding-side
var responseCodeSuccess = byte(0x00)
var responseCodeInvalidRequest = byte(0x01)
var responseCodeServerError = byte(0x02)
var responseCodeResourceUnavailable = byte(0x03)

func FuzzRespCode(code *byte) {
	log.Info("Fuzzing response code")
	number := rand.Intn(4)
	switch number {
	case 0:
		*code = responseCodeSuccess
	case 1:
		*code = responseCodeInvalidRequest
	case 2:
		*code = responseCodeServerError
	case 3:
		*code = responseCodeResourceUnavailable
	}
	return
}
