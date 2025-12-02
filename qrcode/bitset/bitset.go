// go-qrcode
// Copyright 2014 Tom Harwood

// Package bitset implements an append only bit array.
//
// To create a Bitset and append some bits:
//
//	                                  // Bitset Contents
//	b := bitset.New()                 // {}
//	b.AppendBools(true, true, false)  // {1, 1, 0}
//	b.AppendBools(true)               // {1, 1, 0, 1}
//	b.AppendValue(0x02, 4)            // {1, 1, 0, 1, 0, 0, 1, 0}
//
// To read values:
//
//	len := b.Len()                    // 8
//	v := b.At(0)                      // 1
//	v = b.At(1)                       // 1
//	v = b.At(2)                       // 0
//	v = b.At(8)                       // 0
package __obf_50ca330d3b265607

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_b377d83febcb24eb = false
	__obf_503127ef0c093c78 = true
)

// Bitset stores an array of bits.
type Bitset struct {
	__obf_a3af651309e3498e int // The number of bits stored.
	__obf_6b2166ce5f3b90d5 []byte// Storage for individual bits.

}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_7ea27a41aef29670 ...bool) *Bitset {
	__obf_3d3ca6a14713b521 := &Bitset{__obf_a3af651309e3498e: 0, __obf_6b2166ce5f3b90d5: make([]byte, 0)}
	__obf_3d3ca6a14713b521.
		AppendBools(__obf_7ea27a41aef29670...)

	return __obf_3d3ca6a14713b521
}

// Clone returns a copy.
func Clone(__obf_16d25022355980e4 *Bitset) *Bitset {
	return &Bitset{__obf_a3af651309e3498e: __obf_16d25022355980e4.__obf_a3af651309e3498e,

		// Substr returns a substring, consisting of the bits from indexes start to end.
		__obf_6b2166ce5f3b90d5: __obf_16d25022355980e4.__obf_6b2166ce5f3b90d5[:]}
}

func (__obf_3d3ca6a14713b521 *Bitset) Substr(__obf_5500e31ff6756c59 int, __obf_1e858ef2436be756 int) *Bitset {
	if __obf_5500e31ff6756c59 > __obf_1e858ef2436be756 || __obf_1e858ef2436be756 > __obf_3d3ca6a14713b521.__obf_a3af651309e3498e {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_5500e31ff6756c59, __obf_1e858ef2436be756, __obf_3d3ca6a14713b521.__obf_a3af651309e3498e)
	}
	__obf_e075a5af4911717c := New()
	__obf_e075a5af4911717c.__obf_f5d932a04e5b7b94(__obf_1e858ef2436be756 - __obf_5500e31ff6756c59)

	for __obf_0a20d34e8e4c171b := __obf_5500e31ff6756c59; __obf_0a20d34e8e4c171b < __obf_1e858ef2436be756; __obf_0a20d34e8e4c171b++ {
		if __obf_3d3ca6a14713b521.At(__obf_0a20d34e8e4c171b) {
			__obf_e075a5af4911717c.__obf_6b2166ce5f3b90d5[__obf_e075a5af4911717c.__obf_a3af651309e3498e/8] |= 0x80 >> uint(__obf_e075a5af4911717c.__obf_a3af651309e3498e%8)
		}
		__obf_e075a5af4911717c.__obf_a3af651309e3498e++
	}

	return __obf_e075a5af4911717c
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_ac05f1d9bc3b226e string) *Bitset {
	__obf_3d3ca6a14713b521 := &Bitset{__obf_a3af651309e3498e: 0, __obf_6b2166ce5f3b90d5: make([]byte, 0)}

	for _, __obf_69ad54dc0899be24 := range __obf_ac05f1d9bc3b226e {
		switch __obf_69ad54dc0899be24 {
		case '1':
			__obf_3d3ca6a14713b521.
				AppendBools(true)
		case '0':
			__obf_3d3ca6a14713b521.
				AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_69ad54dc0899be24)
		}
	}

	return __obf_3d3ca6a14713b521
}

// AppendBytes appends a list of whole bytes.
func (__obf_3d3ca6a14713b521 *Bitset) AppendBytes(__obf_0b3087b9eb154e0e []byte) {
	for _, __obf_87c5182250643f1b := range __obf_0b3087b9eb154e0e {
		__obf_3d3ca6a14713b521.
			AppendByte(__obf_87c5182250643f1b, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_3d3ca6a14713b521 *Bitset) AppendByte(__obf_ec678b00a7acdae1 byte, __obf_a3af651309e3498e int) {
	__obf_3d3ca6a14713b521.__obf_f5d932a04e5b7b94(__obf_a3af651309e3498e)

	if __obf_a3af651309e3498e > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_a3af651309e3498e)
	}

	for __obf_0a20d34e8e4c171b := __obf_a3af651309e3498e - 1; __obf_0a20d34e8e4c171b >= 0; __obf_0a20d34e8e4c171b-- {
		if __obf_ec678b00a7acdae1&(1<<uint(__obf_0a20d34e8e4c171b)) != 0 {
			__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_3d3ca6a14713b521.__obf_a3af651309e3498e/8] |= 0x80 >> uint(__obf_3d3ca6a14713b521.__obf_a3af651309e3498e%8)
		}
		__obf_3d3ca6a14713b521.

			// AppendUint32 appends the numBits least significant bits from value.
			__obf_a3af651309e3498e++
	}
}

func (__obf_3d3ca6a14713b521 *Bitset) AppendUint32(__obf_ec678b00a7acdae1 uint32, __obf_a3af651309e3498e int) {
	__obf_3d3ca6a14713b521.__obf_f5d932a04e5b7b94(__obf_a3af651309e3498e)

	if __obf_a3af651309e3498e > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_a3af651309e3498e)
	}

	for __obf_0a20d34e8e4c171b := __obf_a3af651309e3498e - 1; __obf_0a20d34e8e4c171b >= 0; __obf_0a20d34e8e4c171b-- {
		if __obf_ec678b00a7acdae1&(1<<uint(__obf_0a20d34e8e4c171b)) != 0 {
			__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_3d3ca6a14713b521.__obf_a3af651309e3498e/8] |= 0x80 >> uint(__obf_3d3ca6a14713b521.__obf_a3af651309e3498e%8)
		}
		__obf_3d3ca6a14713b521.

			// ensureCapacity ensures the Bitset can store an additional |numBits|.
			//
			// The underlying array is expanded if necessary. To prevent frequent
			// reallocation, expanding the underlying array at least doubles its capacity.
			__obf_a3af651309e3498e++
	}
}

func (__obf_3d3ca6a14713b521 *Bitset) __obf_f5d932a04e5b7b94(__obf_a3af651309e3498e int) {
	__obf_a3af651309e3498e += __obf_3d3ca6a14713b521.__obf_a3af651309e3498e
	__obf_cdc010bfd47702cb := __obf_a3af651309e3498e / 8
	if __obf_a3af651309e3498e%8 != 0 {
		__obf_cdc010bfd47702cb++
	}

	if len(__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5) >= __obf_cdc010bfd47702cb {
		return
	}
	__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5 = append(__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5, make([]byte, __obf_cdc010bfd47702cb+2*len(__obf_3d3ca6a14713b521.

		// Append bits copied from |other|.
		//
		// The new length is b.Len() + other.Len().
		__obf_6b2166ce5f3b90d5))...)
}

func (__obf_3d3ca6a14713b521 *Bitset) Append(__obf_d592f73ea8084efb *Bitset) {
	__obf_3d3ca6a14713b521.__obf_f5d932a04e5b7b94(__obf_d592f73ea8084efb.__obf_a3af651309e3498e)

	for __obf_0a20d34e8e4c171b := 0; __obf_0a20d34e8e4c171b < __obf_d592f73ea8084efb.__obf_a3af651309e3498e; __obf_0a20d34e8e4c171b++ {
		if __obf_d592f73ea8084efb.At(__obf_0a20d34e8e4c171b) {
			__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_3d3ca6a14713b521.__obf_a3af651309e3498e/8] |= 0x80 >> uint(__obf_3d3ca6a14713b521.__obf_a3af651309e3498e%8)
		}
		__obf_3d3ca6a14713b521.

			// AppendBools appends bits to the Bitset.
			__obf_a3af651309e3498e++
	}
}

func (__obf_3d3ca6a14713b521 *Bitset) AppendBools(__obf_6b2166ce5f3b90d5 ...bool) {
	__obf_3d3ca6a14713b521.__obf_f5d932a04e5b7b94(len(__obf_6b2166ce5f3b90d5))

	for _, __obf_7ea27a41aef29670 := range __obf_6b2166ce5f3b90d5 {
		if __obf_7ea27a41aef29670 {
			__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_3d3ca6a14713b521.__obf_a3af651309e3498e/8] |= 0x80 >> uint(__obf_3d3ca6a14713b521.__obf_a3af651309e3498e%8)
		}
		__obf_3d3ca6a14713b521.

			// AppendNumBools appends num bits of value value.
			__obf_a3af651309e3498e++
	}
}

func (__obf_3d3ca6a14713b521 *Bitset) AppendNumBools(__obf_3e2843622fd2db52 int, __obf_ec678b00a7acdae1 bool) {
	for __obf_0a20d34e8e4c171b := 0; __obf_0a20d34e8e4c171b < __obf_3e2843622fd2db52; __obf_0a20d34e8e4c171b++ {
		__obf_3d3ca6a14713b521.
			AppendBools(__obf_ec678b00a7acdae1)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_3d3ca6a14713b521 *Bitset) String() string {
	var __obf_ccf4fdfc6ecefc81 string
	for __obf_0a20d34e8e4c171b := 0; __obf_0a20d34e8e4c171b < __obf_3d3ca6a14713b521.__obf_a3af651309e3498e; __obf_0a20d34e8e4c171b++ {
		if (__obf_0a20d34e8e4c171b % 8) == 0 {
			__obf_ccf4fdfc6ecefc81 += " "
		}

		if (__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_0a20d34e8e4c171b/8] & (0x80 >> byte(__obf_0a20d34e8e4c171b%8))) != 0 {
			__obf_ccf4fdfc6ecefc81 += "1"
		} else {
			__obf_ccf4fdfc6ecefc81 += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_3d3ca6a14713b521.__obf_a3af651309e3498e,

		// Len returns the length of the Bitset in bits.
		__obf_ccf4fdfc6ecefc81)
}

func (__obf_3d3ca6a14713b521 *Bitset) Len() int {
	return __obf_3d3ca6a14713b521.

		// Bits returns the contents of the Bitset.
		__obf_a3af651309e3498e
}

func (__obf_3d3ca6a14713b521 *Bitset) Bits() []bool {
	__obf_e075a5af4911717c := make([]bool, __obf_3d3ca6a14713b521.__obf_a3af651309e3498e)

	var __obf_0a20d34e8e4c171b int
	for __obf_0a20d34e8e4c171b = 0; __obf_0a20d34e8e4c171b < __obf_3d3ca6a14713b521.__obf_a3af651309e3498e; __obf_0a20d34e8e4c171b++ {
		__obf_e075a5af4911717c[__obf_0a20d34e8e4c171b] = (__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_0a20d34e8e4c171b/8] & (0x80 >> byte(__obf_0a20d34e8e4c171b%8))) != 0
	}

	return __obf_e075a5af4911717c
}

// At returns the value of the bit at |index|.
func (__obf_3d3ca6a14713b521 *Bitset) At(__obf_24294cac9cc73c0c int) bool {
	if __obf_24294cac9cc73c0c >= __obf_3d3ca6a14713b521.__obf_a3af651309e3498e {
		log.Panicf("Index %d out of range", __obf_24294cac9cc73c0c)
	}

	return (__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_24294cac9cc73c0c/8] & (0x80 >> byte(__obf_24294cac9cc73c0c%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_3d3ca6a14713b521 *Bitset) Equals(__obf_d592f73ea8084efb *Bitset) bool {
	if __obf_3d3ca6a14713b521.__obf_a3af651309e3498e != __obf_d592f73ea8084efb.__obf_a3af651309e3498e {
		return false
	}

	if !bytes.Equal(__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[0:__obf_3d3ca6a14713b521.__obf_a3af651309e3498e/8], __obf_d592f73ea8084efb.__obf_6b2166ce5f3b90d5[0:__obf_3d3ca6a14713b521.__obf_a3af651309e3498e/8]) {
		return false
	}

	for __obf_0a20d34e8e4c171b := 8 * (__obf_3d3ca6a14713b521.__obf_a3af651309e3498e / 8); __obf_0a20d34e8e4c171b < __obf_3d3ca6a14713b521.__obf_a3af651309e3498e; __obf_0a20d34e8e4c171b++ {
		__obf_c0253ef4775aee9f := (__obf_3d3ca6a14713b521.__obf_6b2166ce5f3b90d5[__obf_0a20d34e8e4c171b/8] & (0x80 >> byte(__obf_0a20d34e8e4c171b%8)))
		__obf_3d3ca6a14713b521 := (__obf_d592f73ea8084efb.__obf_6b2166ce5f3b90d5[__obf_0a20d34e8e4c171b/8] & (0x80 >> byte(__obf_0a20d34e8e4c171b%8)))

		if __obf_c0253ef4775aee9f != __obf_3d3ca6a14713b521 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_3d3ca6a14713b521 *Bitset) ByteAt(__obf_24294cac9cc73c0c int) byte {
	if __obf_24294cac9cc73c0c < 0 || __obf_24294cac9cc73c0c >= __obf_3d3ca6a14713b521.__obf_a3af651309e3498e {
		log.Panicf("Index %d out of range", __obf_24294cac9cc73c0c)
	}

	var __obf_e075a5af4911717c byte

	for __obf_0a20d34e8e4c171b := __obf_24294cac9cc73c0c; __obf_0a20d34e8e4c171b < __obf_24294cac9cc73c0c+8 && __obf_0a20d34e8e4c171b < __obf_3d3ca6a14713b521.__obf_a3af651309e3498e; __obf_0a20d34e8e4c171b++ {
		__obf_e075a5af4911717c <<= 1
		if __obf_3d3ca6a14713b521.At(__obf_0a20d34e8e4c171b) {
			__obf_e075a5af4911717c |= 1
		}
	}

	return __obf_e075a5af4911717c
}
