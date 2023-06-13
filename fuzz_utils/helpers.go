package fuzz_utils

import (
	"math/rand"

	"github.com/prysmaticlabs/prysm/v4/cmd/beacon-chain/flags"
	"github.com/urfave/cli/v2"
)

// default is to not fuzz
var Fuzziness int = 0

// this is called at boot time and sets the global fuzziness level
func SetFuzziness(ctx *cli.Context) {
	Fuzziness = ctx.Int(flags.FuzzinessFlag.Name)
}

// this is called in-line at runtime in various req/response message processing
// routines to determine if we should fuzz
func ShouldFuzz() bool {
	log.Info("Fuzziness", Fuzziness)
	if Fuzziness > rand.Intn(100) {
		log.Info("FUZZING")
		return true
	}
	log.Info("NOT FUZZING")
	return false
	//return Fuzziness > rand.Intn(100)
}
