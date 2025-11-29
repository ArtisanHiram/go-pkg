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
package __obf_61d9a54ff1b000b6

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_a08bc8c0e2297bfd = false
	__obf_4929e11ed627cba7 = true
)

// Bitset stores an array of bits.
type Bitset struct {
	__obf_3791bdde2bdd5e22 int // The number of bits stored.
	__obf_ed6888cbe081d7e8 []byte// Storage for individual bits.

}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_05ee4f102a378a25 ...bool) *Bitset {
	__obf_2f1ea0ddff159105 := &Bitset{__obf_3791bdde2bdd5e22: 0, __obf_ed6888cbe081d7e8: make([]byte, 0)}
	__obf_2f1ea0ddff159105.
		AppendBools(__obf_05ee4f102a378a25...)

	return __obf_2f1ea0ddff159105
}

// Clone returns a copy.
func Clone(__obf_d400259272a5fe28 *Bitset) *Bitset {
	return &Bitset{__obf_3791bdde2bdd5e22: __obf_d400259272a5fe28.__obf_3791bdde2bdd5e22,

		// Substr returns a substring, consisting of the bits from indexes start to end.
		__obf_ed6888cbe081d7e8: __obf_d400259272a5fe28.__obf_ed6888cbe081d7e8[:]}
}

func (__obf_2f1ea0ddff159105 *Bitset) Substr(__obf_3cccf3a781737e33 int, __obf_cf1789608b35fbaa int) *Bitset {
	if __obf_3cccf3a781737e33 > __obf_cf1789608b35fbaa || __obf_cf1789608b35fbaa > __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22 {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_3cccf3a781737e33, __obf_cf1789608b35fbaa, __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22)
	}
	__obf_da769d1f0fe6cafd := New()
	__obf_da769d1f0fe6cafd.__obf_bb404f1fa6fa8912(__obf_cf1789608b35fbaa - __obf_3cccf3a781737e33)

	for __obf_252ba549ea50313b := __obf_3cccf3a781737e33; __obf_252ba549ea50313b < __obf_cf1789608b35fbaa; __obf_252ba549ea50313b++ {
		if __obf_2f1ea0ddff159105.At(__obf_252ba549ea50313b) {
			__obf_da769d1f0fe6cafd.__obf_ed6888cbe081d7e8[__obf_da769d1f0fe6cafd.__obf_3791bdde2bdd5e22/8] |= 0x80 >> uint(__obf_da769d1f0fe6cafd.__obf_3791bdde2bdd5e22%8)
		}
		__obf_da769d1f0fe6cafd.__obf_3791bdde2bdd5e22++
	}

	return __obf_da769d1f0fe6cafd
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_1f7cc5fe5b141109 string) *Bitset {
	__obf_2f1ea0ddff159105 := &Bitset{__obf_3791bdde2bdd5e22: 0, __obf_ed6888cbe081d7e8: make([]byte, 0)}

	for _, __obf_6ab66469a0f8b0d8 := range __obf_1f7cc5fe5b141109 {
		switch __obf_6ab66469a0f8b0d8 {
		case '1':
			__obf_2f1ea0ddff159105.
				AppendBools(true)
		case '0':
			__obf_2f1ea0ddff159105.
				AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_6ab66469a0f8b0d8)
		}
	}

	return __obf_2f1ea0ddff159105
}

// AppendBytes appends a list of whole bytes.
func (__obf_2f1ea0ddff159105 *Bitset) AppendBytes(__obf_e48c18108a5a4e01 []byte) {
	for _, __obf_318e56ae2c9c42f0 := range __obf_e48c18108a5a4e01 {
		__obf_2f1ea0ddff159105.
			AppendByte(__obf_318e56ae2c9c42f0, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_2f1ea0ddff159105 *Bitset) AppendByte(__obf_b03ba9fd7dceb500 byte, __obf_3791bdde2bdd5e22 int) {
	__obf_2f1ea0ddff159105.__obf_bb404f1fa6fa8912(__obf_3791bdde2bdd5e22)

	if __obf_3791bdde2bdd5e22 > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_3791bdde2bdd5e22)
	}

	for __obf_252ba549ea50313b := __obf_3791bdde2bdd5e22 - 1; __obf_252ba549ea50313b >= 0; __obf_252ba549ea50313b-- {
		if __obf_b03ba9fd7dceb500&(1<<uint(__obf_252ba549ea50313b)) != 0 {
			__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22/8] |= 0x80 >> uint(__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22%8)
		}
		__obf_2f1ea0ddff159105.

			// AppendUint32 appends the numBits least significant bits from value.
			__obf_3791bdde2bdd5e22++
	}
}

func (__obf_2f1ea0ddff159105 *Bitset) AppendUint32(__obf_b03ba9fd7dceb500 uint32, __obf_3791bdde2bdd5e22 int) {
	__obf_2f1ea0ddff159105.__obf_bb404f1fa6fa8912(__obf_3791bdde2bdd5e22)

	if __obf_3791bdde2bdd5e22 > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_3791bdde2bdd5e22)
	}

	for __obf_252ba549ea50313b := __obf_3791bdde2bdd5e22 - 1; __obf_252ba549ea50313b >= 0; __obf_252ba549ea50313b-- {
		if __obf_b03ba9fd7dceb500&(1<<uint(__obf_252ba549ea50313b)) != 0 {
			__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22/8] |= 0x80 >> uint(__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22%8)
		}
		__obf_2f1ea0ddff159105.

			// ensureCapacity ensures the Bitset can store an additional |numBits|.
			//
			// The underlying array is expanded if necessary. To prevent frequent
			// reallocation, expanding the underlying array at least doubles its capacity.
			__obf_3791bdde2bdd5e22++
	}
}

func (__obf_2f1ea0ddff159105 *Bitset) __obf_bb404f1fa6fa8912(__obf_3791bdde2bdd5e22 int) {
	__obf_3791bdde2bdd5e22 += __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22
	__obf_0ad45f572f3a6af6 := __obf_3791bdde2bdd5e22 / 8
	if __obf_3791bdde2bdd5e22%8 != 0 {
		__obf_0ad45f572f3a6af6++
	}

	if len(__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8) >= __obf_0ad45f572f3a6af6 {
		return
	}
	__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8 = append(__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8, make([]byte, __obf_0ad45f572f3a6af6+2*len(__obf_2f1ea0ddff159105.

		// Append bits copied from |other|.
		//
		// The new length is b.Len() + other.Len().
		__obf_ed6888cbe081d7e8))...)
}

func (__obf_2f1ea0ddff159105 *Bitset) Append(__obf_649f5c1c658a2687 *Bitset) {
	__obf_2f1ea0ddff159105.__obf_bb404f1fa6fa8912(__obf_649f5c1c658a2687.__obf_3791bdde2bdd5e22)

	for __obf_252ba549ea50313b := 0; __obf_252ba549ea50313b < __obf_649f5c1c658a2687.__obf_3791bdde2bdd5e22; __obf_252ba549ea50313b++ {
		if __obf_649f5c1c658a2687.At(__obf_252ba549ea50313b) {
			__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22/8] |= 0x80 >> uint(__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22%8)
		}
		__obf_2f1ea0ddff159105.

			// AppendBools appends bits to the Bitset.
			__obf_3791bdde2bdd5e22++
	}
}

func (__obf_2f1ea0ddff159105 *Bitset) AppendBools(__obf_ed6888cbe081d7e8 ...bool) {
	__obf_2f1ea0ddff159105.__obf_bb404f1fa6fa8912(len(__obf_ed6888cbe081d7e8))

	for _, __obf_05ee4f102a378a25 := range __obf_ed6888cbe081d7e8 {
		if __obf_05ee4f102a378a25 {
			__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22/8] |= 0x80 >> uint(__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22%8)
		}
		__obf_2f1ea0ddff159105.

			// AppendNumBools appends num bits of value value.
			__obf_3791bdde2bdd5e22++
	}
}

func (__obf_2f1ea0ddff159105 *Bitset) AppendNumBools(__obf_1c388782f6c3c5f8 int, __obf_b03ba9fd7dceb500 bool) {
	for __obf_252ba549ea50313b := 0; __obf_252ba549ea50313b < __obf_1c388782f6c3c5f8; __obf_252ba549ea50313b++ {
		__obf_2f1ea0ddff159105.
			AppendBools(__obf_b03ba9fd7dceb500)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_2f1ea0ddff159105 *Bitset) String() string {
	var __obf_6b75a80baaa801a7 string
	for __obf_252ba549ea50313b := 0; __obf_252ba549ea50313b < __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22; __obf_252ba549ea50313b++ {
		if (__obf_252ba549ea50313b % 8) == 0 {
			__obf_6b75a80baaa801a7 += " "
		}

		if (__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_252ba549ea50313b/8] & (0x80 >> byte(__obf_252ba549ea50313b%8))) != 0 {
			__obf_6b75a80baaa801a7 += "1"
		} else {
			__obf_6b75a80baaa801a7 += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22,

		// Len returns the length of the Bitset in bits.
		__obf_6b75a80baaa801a7)
}

func (__obf_2f1ea0ddff159105 *Bitset) Len() int {
	return __obf_2f1ea0ddff159105.

		// Bits returns the contents of the Bitset.
		__obf_3791bdde2bdd5e22
}

func (__obf_2f1ea0ddff159105 *Bitset) Bits() []bool {
	__obf_da769d1f0fe6cafd := make([]bool, __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22)

	var __obf_252ba549ea50313b int
	for __obf_252ba549ea50313b = 0; __obf_252ba549ea50313b < __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22; __obf_252ba549ea50313b++ {
		__obf_da769d1f0fe6cafd[__obf_252ba549ea50313b] = (__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_252ba549ea50313b/8] & (0x80 >> byte(__obf_252ba549ea50313b%8))) != 0
	}

	return __obf_da769d1f0fe6cafd
}

// At returns the value of the bit at |index|.
func (__obf_2f1ea0ddff159105 *Bitset) At(__obf_d1e12186c8016861 int) bool {
	if __obf_d1e12186c8016861 >= __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22 {
		log.Panicf("Index %d out of range", __obf_d1e12186c8016861)
	}

	return (__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_d1e12186c8016861/8] & (0x80 >> byte(__obf_d1e12186c8016861%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_2f1ea0ddff159105 *Bitset) Equals(__obf_649f5c1c658a2687 *Bitset) bool {
	if __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22 != __obf_649f5c1c658a2687.__obf_3791bdde2bdd5e22 {
		return false
	}

	if !bytes.Equal(__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[0:__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22/8], __obf_649f5c1c658a2687.__obf_ed6888cbe081d7e8[0:__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22/8]) {
		return false
	}

	for __obf_252ba549ea50313b := 8 * (__obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22 / 8); __obf_252ba549ea50313b < __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22; __obf_252ba549ea50313b++ {
		__obf_43107b12625f1e41 := (__obf_2f1ea0ddff159105.__obf_ed6888cbe081d7e8[__obf_252ba549ea50313b/8] & (0x80 >> byte(__obf_252ba549ea50313b%8)))
		__obf_2f1ea0ddff159105 := (__obf_649f5c1c658a2687.__obf_ed6888cbe081d7e8[__obf_252ba549ea50313b/8] & (0x80 >> byte(__obf_252ba549ea50313b%8)))

		if __obf_43107b12625f1e41 != __obf_2f1ea0ddff159105 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_2f1ea0ddff159105 *Bitset) ByteAt(__obf_d1e12186c8016861 int) byte {
	if __obf_d1e12186c8016861 < 0 || __obf_d1e12186c8016861 >= __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22 {
		log.Panicf("Index %d out of range", __obf_d1e12186c8016861)
	}

	var __obf_da769d1f0fe6cafd byte

	for __obf_252ba549ea50313b := __obf_d1e12186c8016861; __obf_252ba549ea50313b < __obf_d1e12186c8016861+8 && __obf_252ba549ea50313b < __obf_2f1ea0ddff159105.__obf_3791bdde2bdd5e22; __obf_252ba549ea50313b++ {
		__obf_da769d1f0fe6cafd <<= 1
		if __obf_2f1ea0ddff159105.At(__obf_252ba549ea50313b) {
			__obf_da769d1f0fe6cafd |= 1
		}
	}

	return __obf_da769d1f0fe6cafd
}
