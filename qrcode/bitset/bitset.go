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
package __obf_440c491c36e8b418

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_b1f97aa1d99d020a = false
	__obf_2b26d1de1560cf2c = true
)

// Bitset stores an array of bits.
type Bitset struct {
	__obf_3dc0516df5d61e4a int // The number of bits stored.
	__obf_ce7e41377213bc67 []byte// Storage for individual bits.

}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_f567a88747ec4c65 ...bool) *Bitset {
	__obf_a82d6fa4013a0627 := &Bitset{__obf_3dc0516df5d61e4a: 0, __obf_ce7e41377213bc67: make([]byte, 0)}
	__obf_a82d6fa4013a0627.
		AppendBools(__obf_f567a88747ec4c65...)

	return __obf_a82d6fa4013a0627
}

// Clone returns a copy.
func Clone(__obf_2d60fd86a0b5c909 *Bitset) *Bitset {
	return &Bitset{__obf_3dc0516df5d61e4a: __obf_2d60fd86a0b5c909.__obf_3dc0516df5d61e4a,

		// Substr returns a substring, consisting of the bits from indexes start to end.
		__obf_ce7e41377213bc67: __obf_2d60fd86a0b5c909.__obf_ce7e41377213bc67[:]}
}

func (__obf_a82d6fa4013a0627 *Bitset) Substr(__obf_161d3d4b22654865 int, __obf_bef7186a28759839 int) *Bitset {
	if __obf_161d3d4b22654865 > __obf_bef7186a28759839 || __obf_bef7186a28759839 > __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_161d3d4b22654865, __obf_bef7186a28759839, __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a)
	}
	__obf_589be14fc1d66e6c := New()
	__obf_589be14fc1d66e6c.__obf_496a11ce74c6a31f(__obf_bef7186a28759839 - __obf_161d3d4b22654865)

	for __obf_a45ee9c5c39c4e4b := __obf_161d3d4b22654865; __obf_a45ee9c5c39c4e4b < __obf_bef7186a28759839; __obf_a45ee9c5c39c4e4b++ {
		if __obf_a82d6fa4013a0627.At(__obf_a45ee9c5c39c4e4b) {
			__obf_589be14fc1d66e6c.__obf_ce7e41377213bc67[__obf_589be14fc1d66e6c.__obf_3dc0516df5d61e4a/8] |= 0x80 >> uint(__obf_589be14fc1d66e6c.__obf_3dc0516df5d61e4a%8)
		}
		__obf_589be14fc1d66e6c.__obf_3dc0516df5d61e4a++
	}

	return __obf_589be14fc1d66e6c
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_c0192433ee3c6298 string) *Bitset {
	__obf_a82d6fa4013a0627 := &Bitset{__obf_3dc0516df5d61e4a: 0, __obf_ce7e41377213bc67: make([]byte, 0)}

	for _, __obf_919e1443e2c22cbf := range __obf_c0192433ee3c6298 {
		switch __obf_919e1443e2c22cbf {
		case '1':
			__obf_a82d6fa4013a0627.
				AppendBools(true)
		case '0':
			__obf_a82d6fa4013a0627.
				AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_919e1443e2c22cbf)
		}
	}

	return __obf_a82d6fa4013a0627
}

// AppendBytes appends a list of whole bytes.
func (__obf_a82d6fa4013a0627 *Bitset) AppendBytes(__obf_84996dc1e29283e5 []byte) {
	for _, __obf_2f13770ad0a46b04 := range __obf_84996dc1e29283e5 {
		__obf_a82d6fa4013a0627.
			AppendByte(__obf_2f13770ad0a46b04, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_a82d6fa4013a0627 *Bitset) AppendByte(__obf_55de05699bb11c0e byte, __obf_3dc0516df5d61e4a int) {
	__obf_a82d6fa4013a0627.__obf_496a11ce74c6a31f(__obf_3dc0516df5d61e4a)

	if __obf_3dc0516df5d61e4a > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_3dc0516df5d61e4a)
	}

	for __obf_a45ee9c5c39c4e4b := __obf_3dc0516df5d61e4a - 1; __obf_a45ee9c5c39c4e4b >= 0; __obf_a45ee9c5c39c4e4b-- {
		if __obf_55de05699bb11c0e&(1<<uint(__obf_a45ee9c5c39c4e4b)) != 0 {
			__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a/8] |= 0x80 >> uint(__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a%8)
		}
		__obf_a82d6fa4013a0627.

			// AppendUint32 appends the numBits least significant bits from value.
			__obf_3dc0516df5d61e4a++
	}
}

func (__obf_a82d6fa4013a0627 *Bitset) AppendUint32(__obf_55de05699bb11c0e uint32, __obf_3dc0516df5d61e4a int) {
	__obf_a82d6fa4013a0627.__obf_496a11ce74c6a31f(__obf_3dc0516df5d61e4a)

	if __obf_3dc0516df5d61e4a > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_3dc0516df5d61e4a)
	}

	for __obf_a45ee9c5c39c4e4b := __obf_3dc0516df5d61e4a - 1; __obf_a45ee9c5c39c4e4b >= 0; __obf_a45ee9c5c39c4e4b-- {
		if __obf_55de05699bb11c0e&(1<<uint(__obf_a45ee9c5c39c4e4b)) != 0 {
			__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a/8] |= 0x80 >> uint(__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a%8)
		}
		__obf_a82d6fa4013a0627.

			// ensureCapacity ensures the Bitset can store an additional |numBits|.
			//
			// The underlying array is expanded if necessary. To prevent frequent
			// reallocation, expanding the underlying array at least doubles its capacity.
			__obf_3dc0516df5d61e4a++
	}
}

func (__obf_a82d6fa4013a0627 *Bitset) __obf_496a11ce74c6a31f(__obf_3dc0516df5d61e4a int) {
	__obf_3dc0516df5d61e4a += __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a
	__obf_000d45cf3dfb5bc8 := __obf_3dc0516df5d61e4a / 8
	if __obf_3dc0516df5d61e4a%8 != 0 {
		__obf_000d45cf3dfb5bc8++
	}

	if len(__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67) >= __obf_000d45cf3dfb5bc8 {
		return
	}
	__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67 = append(__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67, make([]byte, __obf_000d45cf3dfb5bc8+2*len(__obf_a82d6fa4013a0627.

		// Append bits copied from |other|.
		//
		// The new length is b.Len() + other.Len().
		__obf_ce7e41377213bc67))...)
}

func (__obf_a82d6fa4013a0627 *Bitset) Append(__obf_ba72ef35a3a6e847 *Bitset) {
	__obf_a82d6fa4013a0627.__obf_496a11ce74c6a31f(__obf_ba72ef35a3a6e847.__obf_3dc0516df5d61e4a)

	for __obf_a45ee9c5c39c4e4b := 0; __obf_a45ee9c5c39c4e4b < __obf_ba72ef35a3a6e847.__obf_3dc0516df5d61e4a; __obf_a45ee9c5c39c4e4b++ {
		if __obf_ba72ef35a3a6e847.At(__obf_a45ee9c5c39c4e4b) {
			__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a/8] |= 0x80 >> uint(__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a%8)
		}
		__obf_a82d6fa4013a0627.

			// AppendBools appends bits to the Bitset.
			__obf_3dc0516df5d61e4a++
	}
}

func (__obf_a82d6fa4013a0627 *Bitset) AppendBools(__obf_ce7e41377213bc67 ...bool) {
	__obf_a82d6fa4013a0627.__obf_496a11ce74c6a31f(len(__obf_ce7e41377213bc67))

	for _, __obf_f567a88747ec4c65 := range __obf_ce7e41377213bc67 {
		if __obf_f567a88747ec4c65 {
			__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a/8] |= 0x80 >> uint(__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a%8)
		}
		__obf_a82d6fa4013a0627.

			// AppendNumBools appends num bits of value value.
			__obf_3dc0516df5d61e4a++
	}
}

func (__obf_a82d6fa4013a0627 *Bitset) AppendNumBools(__obf_718ff775fcdac031 int, __obf_55de05699bb11c0e bool) {
	for __obf_a45ee9c5c39c4e4b := 0; __obf_a45ee9c5c39c4e4b < __obf_718ff775fcdac031; __obf_a45ee9c5c39c4e4b++ {
		__obf_a82d6fa4013a0627.
			AppendBools(__obf_55de05699bb11c0e)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_a82d6fa4013a0627 *Bitset) String() string {
	var __obf_92cadb7eaf3c7009 string
	for __obf_a45ee9c5c39c4e4b := 0; __obf_a45ee9c5c39c4e4b < __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a; __obf_a45ee9c5c39c4e4b++ {
		if (__obf_a45ee9c5c39c4e4b % 8) == 0 {
			__obf_92cadb7eaf3c7009 += " "
		}

		if (__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a45ee9c5c39c4e4b/8] & (0x80 >> byte(__obf_a45ee9c5c39c4e4b%8))) != 0 {
			__obf_92cadb7eaf3c7009 += "1"
		} else {
			__obf_92cadb7eaf3c7009 += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a,

		// Len returns the length of the Bitset in bits.
		__obf_92cadb7eaf3c7009)
}

func (__obf_a82d6fa4013a0627 *Bitset) Len() int {
	return __obf_a82d6fa4013a0627.

		// Bits returns the contents of the Bitset.
		__obf_3dc0516df5d61e4a
}

func (__obf_a82d6fa4013a0627 *Bitset) Bits() []bool {
	__obf_589be14fc1d66e6c := make([]bool, __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a)

	var __obf_a45ee9c5c39c4e4b int
	for __obf_a45ee9c5c39c4e4b = 0; __obf_a45ee9c5c39c4e4b < __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a; __obf_a45ee9c5c39c4e4b++ {
		__obf_589be14fc1d66e6c[__obf_a45ee9c5c39c4e4b] = (__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a45ee9c5c39c4e4b/8] & (0x80 >> byte(__obf_a45ee9c5c39c4e4b%8))) != 0
	}

	return __obf_589be14fc1d66e6c
}

// At returns the value of the bit at |index|.
func (__obf_a82d6fa4013a0627 *Bitset) At(__obf_f931ba5bad6cd431 int) bool {
	if __obf_f931ba5bad6cd431 >= __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a {
		log.Panicf("Index %d out of range", __obf_f931ba5bad6cd431)
	}

	return (__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_f931ba5bad6cd431/8] & (0x80 >> byte(__obf_f931ba5bad6cd431%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_a82d6fa4013a0627 *Bitset) Equals(__obf_ba72ef35a3a6e847 *Bitset) bool {
	if __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a != __obf_ba72ef35a3a6e847.__obf_3dc0516df5d61e4a {
		return false
	}

	if !bytes.Equal(__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[0:__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a/8], __obf_ba72ef35a3a6e847.__obf_ce7e41377213bc67[0:__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a/8]) {
		return false
	}

	for __obf_a45ee9c5c39c4e4b := 8 * (__obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a / 8); __obf_a45ee9c5c39c4e4b < __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a; __obf_a45ee9c5c39c4e4b++ {
		__obf_6c46a4b568f24c7f := (__obf_a82d6fa4013a0627.__obf_ce7e41377213bc67[__obf_a45ee9c5c39c4e4b/8] & (0x80 >> byte(__obf_a45ee9c5c39c4e4b%8)))
		__obf_a82d6fa4013a0627 := (__obf_ba72ef35a3a6e847.__obf_ce7e41377213bc67[__obf_a45ee9c5c39c4e4b/8] & (0x80 >> byte(__obf_a45ee9c5c39c4e4b%8)))

		if __obf_6c46a4b568f24c7f != __obf_a82d6fa4013a0627 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_a82d6fa4013a0627 *Bitset) ByteAt(__obf_f931ba5bad6cd431 int) byte {
	if __obf_f931ba5bad6cd431 < 0 || __obf_f931ba5bad6cd431 >= __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a {
		log.Panicf("Index %d out of range", __obf_f931ba5bad6cd431)
	}

	var __obf_589be14fc1d66e6c byte

	for __obf_a45ee9c5c39c4e4b := __obf_f931ba5bad6cd431; __obf_a45ee9c5c39c4e4b < __obf_f931ba5bad6cd431+8 && __obf_a45ee9c5c39c4e4b < __obf_a82d6fa4013a0627.__obf_3dc0516df5d61e4a; __obf_a45ee9c5c39c4e4b++ {
		__obf_589be14fc1d66e6c <<= 1
		if __obf_a82d6fa4013a0627.At(__obf_a45ee9c5c39c4e4b) {
			__obf_589be14fc1d66e6c |= 1
		}
	}

	return __obf_589be14fc1d66e6c
}
