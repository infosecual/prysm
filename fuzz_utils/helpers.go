package fuzz_utils

import (
	"math/rand"

	"github.com/prysmaticlabs/prysm/v4/cmd/beacon-chain/flags"
	"github.com/urfave/cli/v2"
)

// default is to not fuzz
var Fuzziness int = 0
var FuzzForkDigest bool = false

// this is called at boot time and sets the global fuzziness level
func SetFuzzSettings(ctx *cli.Context) {
	Fuzziness = ctx.Int(flags.FuzzinessFlag.Name)
	FuzzForkDigest = ctx.Bool(flags.FuzzForkDigestFlag.Name)
}

// this is called in-line at runtime in various req/response message processing
// routines to determine if we should fuzz
func ShouldFuzz() bool {
	return Fuzziness > rand.Intn(100)
}
