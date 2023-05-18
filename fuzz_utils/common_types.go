package fuzz_utils

import (
	"github.com/ethereum/go-ethereum/log"
)

var responseCodeSuccess = byte(0x00)
var responseCodeInvalidRequest = byte(0x01)
var responseCodeServerError = byte(0x02)

func FuzzRespCode(code *byte) {
	log.Info("Fuzzing response code")
	return
}
