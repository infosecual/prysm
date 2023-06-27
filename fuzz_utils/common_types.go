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
	log.Info("FUZZ - FuzzRespCode before mutation: ", *code)
	number := rand.Intn(5)
	switch number {
	case 0:
		*code = responseCodeSuccess
	case 1:
		*code = responseCodeInvalidRequest
	case 2:
		*code = responseCodeServerError
	case 3:
		*code = responseCodeResourceUnavailable
	case 4:
		// The range [4, 127] is RESERVED for future usages, and should be treated as error if not recognized expressly.
		futureRespCodes := rand.Intn(127-4+1) + 4
		// Clients MAY use response codes above 128 to indicate alternative, erroneous request-specific responses.
		clientRespCodes := rand.Intn(255-128+1) + 128
		if RandBool() {
			*code = byte(futureRespCodes)
		} else {
			*code = byte(clientRespCodes)
		}
	}
	log.Info("FUZZ - FuzzRespCode after mutation: ", *code)
}
