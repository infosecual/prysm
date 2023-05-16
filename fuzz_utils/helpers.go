package fuzz_utils

import (
	"math/rand"

	"github.com/influxdata/influxdb-client-go/v2/log"
	"github.com/prysmaticlabs/go-bitfield"
	"github.com/prysmaticlabs/prysm/v4/cmd/beacon-chain/flags"
	"github.com/urfave/cli/v2"
)

var Fuzziness int = 0

func SetFuzziness(ctx *cli.Context) {
	Fuzziness = ctx.Int(flags.FuzzinessFlag.Name)
}

func ShouldFuzz() bool {
	return Fuzziness > rand.Intn(100)
}

// /////////////////////////////////////////////////////////////////////////////
// metatdata objects
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
