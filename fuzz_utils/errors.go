package fuzz_utils

import "github.com/ethereum/go-ethereum/log"

// /////////////////////////////////////////////////////////////////////////////
// error fuzzers
// /////////////////////////////////////////////////////////////////////////////

func FuzzErrorResponse(originalResponseCode byte, originalReason string) (responseCode byte, reason string) {
	// var responseCodeSuccess = byte(0x00)
	// var responseCodeInvalidRequest = byte(0x01)
	// var responseCodeServerError = byte(0x02)
	//errCode, err = fuzz_utils.FuzzErrorResponse(errCode, types.ErrGeneric.Error())
	log.Info("FUZZ - FuzzError")
	return originalResponseCode, originalReason
}
