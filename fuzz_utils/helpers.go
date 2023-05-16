package fuzz_utils

import (
	"math/rand"

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
