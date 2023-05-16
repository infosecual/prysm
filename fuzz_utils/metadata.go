package fuzz_utils

import (
	"github.com/influxdata/influxdb-client-go/v2/log"
	"github.com/prysmaticlabs/go-bitfield"
)

// /////////////////////////////////////////////////////////////////////////////
// metatdata fuzzers
// /////////////////////////////////////////////////////////////////////////////
func FuzzMetadataV1(Attnets bitfield.Bitvector64, SeqNumber uint64, Syncnets bitfield.Bitvector4) (bitfield.Bitvector64, uint64, bitfield.Bitvector4) {
	// Attnets:   currMd.AttnetsBitfield(), // bitfield.Bitvector64
	// SeqNumber: currMd.SequenceNumber(), // uint64
	// Syncnets:  bitfield.Bitvector4{byte(0x00)} / bitfield.Bitvector4{byte[]}
	log.Log.Info("FUZZ - FuzzMetadataV1")
	return Attnets, SeqNumber, Syncnets
}

func FuzzMetadataV0(Attnets bitfield.Bitvector64, SeqNumber uint64) (bitfield.Bitvector64, uint64) {
	// Attnets:   currMd.AttnetsBitfield(), // bitfield.Bitvector64
	// SeqNumber: currMd.SequenceNumber() // uint64
	log.Log.Info("FUZZ - FuzzMetadataV0")
	return Attnets, SeqNumber
}
