package fuzz_utils

import "math/bits"

// /////////////////////////////////////////////////////////////////////////////
// Mutators
// /////////////////////////////////////////////////////////////////////////////

func MutateBytes(b *[]byte, n int) *[]byte {
	// make a copy of the slice
	//new_slice := make([]byte, len(*b))
	//new_ptr := &new_slice
	//copy(*new_ptr, *b)

	switch x := RandUint8(); {
	case x < 10:
		// mutate the copy
		for i := 0; i < n; i++ {
			// iterate through bytes, 50% of the time mutate the byte
			if RandBool() {
				(*b)[i] = MutateByte((*b)[i])
			}
		}
	case x < 100:
		// pick a random value from cache
		bytes := cache.GetBytes(n, int(RandInt64()))
		*b = bytes
	default:
		// do nothing
	}

	return b
}

func MutateByte(b byte) byte {
	switch x := RandUint8(); {
	case x < 10:
		// Bitflip
		pos := RandUint8() % 8
		return b ^ 1<<pos
	case x < 20:
		// XOR random
		return b ^ RandUint8()
	case x < 30:
		// Inverse
		return 0xff ^ b
	case x < 40:
		// Reverse ByteOrder
		return bits.Reverse8(b)
	case x < 50:
		return bits.RotateLeft8(b, int(RandInt8()))
	default:
		// Random
		return RandUint8()

	}
}
