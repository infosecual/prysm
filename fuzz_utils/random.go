package fuzz_utils

import "math/rand"

// /////////////////////////////////////////////////////////////////////////////
// Random Type Generators
// /////////////////////////////////////////////////////////////////////////////

var (
	interesting8   = []int8{-128, -1, 0, 1, 16, 32, 64, 100, 127}
	interestingU8  = []uint8{0, 1, 16, 32, 64, 100, 127, 128, 255}
	interesting16  = []int16{-32768, -129, 128, 255, 256, 512, 1000, 1024, 4096, 32767}
	interestingU16 = []uint16{0, 1, 256, 512, 1000, 1024, 4096, 32767, 32768, 65535}
	interesting32  = []int32{-2147483648, -100663046, -32769, 32768, 65535, 65536, 100663045, 2147483647}
	interestingU32 = []uint32{0, 1, 65535, 65536, 100663045, 2147483647, 2147483648, 4294967295}
	interesting64  = []int64{-9223372036854775808, -4294967297, -4294967296, -2147483649, -2147483648, -100663047, 100663046, 2147483647, 2147483648, 4294967295, 4294967296, 9223372036854775807}
	interestingU64 = []uint64{0, 1, 2147483647, 2147483648, 4294967295, 4294967296, 9223372036854775807, 9223372036854775808, 18446744073709551615}
)

func TrueWithProb(propability int) bool {
	return rand.Intn(100) < propability
}

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandInt8() int8 {
	if rand.Intn(4) == 0 {
		return interesting8[rand.Intn(len(interesting8))]
	}
	// 75% of the time pick a random value
	return int8(rand.Intn(1 << 8))
}

func RandInt16() int16 {
	if rand.Intn(4) == 0 {
		return interesting16[rand.Intn(len(interesting16))]
	}
	// 75% of the time pick a random value
	return int16(rand.Intn(1 << 16))
}

func RandInt32() int32 {
	if rand.Intn(4) == 0 {
		return interesting32[rand.Intn(len(interesting32))]
	}
	// 75% of the time pick a random value
	return int32(rand.Intn(1 << 32))
}

func RandInt64() int64 {
	if rand.Intn(4) == 0 {
		return interesting64[rand.Intn(len(interesting64))]
	}
	// 75% of the time pick a random value
	return int64(rand.Uint64())
}

func RandUint8() uint8 {
	// 25% of the time pick one of the interesting values
	if rand.Intn(4) == 0 {
		return interestingU8[rand.Intn(len(interestingU8))]
	}
	// 75% of the time pick a random value
	return uint8(rand.Intn(1 << 8))
}

func RandUint16() uint16 {
	// 25% of the time pick one of the interesting values
	if rand.Intn(4) == 0 {
		return interestingU16[rand.Intn(len(interestingU16))]
	}
	// 75% of the time pick a random value
	return uint16(rand.Intn(1 << 16))
}

func RandUint32() uint32 {
	// 25% of the time pick one of the interesting values
	if rand.Intn(4) == 0 {
		return interestingU32[rand.Intn(len(interestingU32))]
	}
	// 75% of the time pick a random value
	return uint32(rand.Intn(1 << 32))
}

func RandUint64() uint64 {
	// 25% of the time pick one of the interesting values
	if rand.Intn(4) == 0 {
		return interestingU64[rand.Intn(len(interestingU64))]
	}
	// 75% of the time pick a random value
	return rand.Uint64()
}
