package fuzz_utils

import (
	"github.com/prysmaticlabs/go-bitfield"
)

// /////////////////////////////////////////////////////////////////////////////
// metatdata fuzzers
// /////////////////////////////////////////////////////////////////////////////
func FuzzMetadataV1(Attnets bitfield.Bitvector64, SeqNumber uint64, Syncnets bitfield.Bitvector4) (bitfield.Bitvector64, uint64, bitfield.Bitvector4) {
	// Attnets:   currMd.AttnetsBitfield(), // bitfield.Bitvector64
	// SeqNumber: currMd.SequenceNumber(), // uint64
	// Syncnets:  bitfield.Bitvector4{byte(0x00)} / bitfield.Bitvector4{byte[]}
	log.Info("FUZZ - FuzzMetadataV1 BEFORE:")
	log.Info("\tAttnets", Attnets)
	log.Info("\tSeqNumber", SeqNumber)
	log.Info("\tSyncnets", Syncnets)

	MutateNBytes((*[]byte)(&Attnets), 8)
	SeqNumber = RandUint64()
	MutateNBytes((*[]byte)(&Syncnets), 1)

	log.Info("FUZZ - FuzzMetadataV1 AFTER:")
	log.Info("\tAttnets", Attnets)
	log.Info("\tSeqNumber", SeqNumber)
	log.Info("\tSyncnets", Syncnets)

	return Attnets, SeqNumber, Syncnets
}

func FuzzMetadataV0(Attnets bitfield.Bitvector64, SeqNumber uint64) (bitfield.Bitvector64, uint64) {
	// Attnets:   currMd.AttnetsBitfield(), // bitfield.Bitvector64
	// SeqNumber: currMd.SequenceNumber() // uint64
	log.Info("FUZZ - FuzzMetadataV0 BEFORE:")
	log.Info("\tAttnets", Attnets)
	log.Info("\tSeqNumber", SeqNumber)

	MutateNBytes((*[]byte)(&Attnets), 8)
	SeqNumber = RandUint64()

	log.Info("FUZZ - FuzzMetadataV0 AFTER:")
	log.Info("\tAttnets", Attnets)
	log.Info("\tSeqNumber", SeqNumber)

	return Attnets, SeqNumber
}
