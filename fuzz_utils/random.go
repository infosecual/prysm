package fuzz_utils

import "math/rand"

// /////////////////////////////////////////////////////////////////////////////
// Random Type Generators
// /////////////////////////////////////////////////////////////////////////////

var (
	interesting8  = []uint8{0, 1, 16, 32, 64, 100, 127, 128, 255}
	interesting16 = []uint16{256, 512, 1000, 1024, 4096, 32767, 32768, 65535}
	interesting32 = []uint32{65536, 100663045, 2147483647, 2147483648, 4294967295}
	interesting64 = []uint64{4294967296, 9223372036854775807, 9223372036854775808, 18446744073709551615}
)

func init() {
	for _, b := range interesting8 {
		interesting16 = append(interesting16, uint16(b))
	}
	for _, b := range interesting16 {
		interesting32 = append(interesting32, uint32(b))
	}
	for _, b := range interesting32 {
		interesting64 = append(interesting64, uint64(b))
	}
}

func Prob(probability int) bool {
	return rand.Intn(100) < probability
}

func RandBool() bool {
	return rand.Intn(2) == 0
}

func RandUint8() uint8 {
	// 25% of the time pick one of the interesting values
	if Prob(25) {
		return interesting8[rand.Intn(len(interesting8))]
	}
	// 75% of the time pick a random value
	return uint8(rand.Intn(1 << 8))
}

func RandUint16() uint16 {
	// 25% of the time pick one of the interesting values
	if Prob(25) {
		return interesting16[rand.Intn(len(interesting16))]
	}
	// 75% of the time pick a random value
	return uint16(rand.Intn(1 << 16))
}

func RandUint32() uint32 {
	// 25% of the time pick one of the interesting values
	if Prob(25) {
		return interesting32[rand.Intn(len(interesting32))]
	}
	// 75% of the time pick a random value
	return uint32(rand.Intn(1 << 32))
}

func RandUint64() uint64 {
	// 25% of the time pick one of the interesting values
	if Prob(25) {
		return interesting64[rand.Intn(len(interesting64))]
	}
	// 75% of the time pick a random value
	return rand.Uint64()
}

func RandInt8() int8 {
	return int8(RandUint8())
}

func RandInt16() int16 {
	return int16(RandUint16())
}

func RandInt32() int32 {
	return int32(RandUint32())
}

func RandInt64() int64 {
	return int64(RandUint64())
}
