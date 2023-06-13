package fuzz_utils

// /////////////////////////////////////////////////////////////////////////////
// error fuzzers
// /////////////////////////////////////////////////////////////////////////////

/// Spec citation: https://github.com/ethereum/consensus-specs/blob/dev/specs/phase0/p2p-interface.md?plain=1#L585-L601

func FuzzWriteErrorResponseToStream(originalResponseCode byte, originalReason string) (responseCode byte, reason string) {
	// var responseCodeSuccess = byte(0x00)
	// var responseCodeInvalidRequest = byte(0x01)
	// var responseCodeServerError = byte(0x02)
	//errCode, err = fuzz_utils.FuzzWriteErrorResponseToStream(errCode, types.ErrGeneric.Error())
	log.Info("FUZZ - FuzzError")
	return originalResponseCode, originalReason
}

func FuzzErrorResponseBytes(OriginalError []byte) []byte {
	// resp, err := createErrorResponse(responseCode, reason, s.cfg.p2p)
	log.Info("FUZZ - FuzzErrorResponseBytes before mutation: ", OriginalError[:1], OriginalError[1:])
	/// Spec
	/// The response code can have one of the following values, encoded as a single unsigned byte:
	//	-  0: **Success** -- a normal response follows, with contents matching the expected message schema and encoding specified in the request.
	//	-  1: **InvalidRequest** -- the contents of the request are semantically invalid, or the payload is malformed, or could not be understood.
	//	The response payload adheres to the `ErrorMessage` schema (described below).
	//	-  2: **ServerError** -- the responder encountered an error while processing the request.
	//	The response payload adheres to the `ErrorMessage` schema (described below).
	//	-  3: **ResourceUnavailable** -- the responder does not have requested resource.
	//  The response payload adheres to the `ErrorMessage` schema (described below).
	//
	// Clients MAY use response codes above `128` to indicate alternative, erroneous request-specific responses.
	// The range `[4, 127]` is RESERVED for future usages, and should be treated as error if not recognized expressly.
	var responseCode byte
	FuzzRespCode(&responseCode)
	payload := OriginalError[1:]
	MutateNBytes(&payload, 255)
	log.Info("FUZZ - FuzzErrorResponseBytes after mutation: ", responseCode, payload)
	return append([]byte{responseCode}, payload...)
}
