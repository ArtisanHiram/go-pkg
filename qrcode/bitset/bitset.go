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
package __obf_09471402236bac99

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_40b6879463fc98c0 = false
	__obf_3d368a5192cfddb0 = true
)

// Bitset stores an array of bits.
type Bitset struct {
	__obf_b40d2843659b7788 int // The number of bits stored.
	__obf_a98e0f33b9379401 []byte// Storage for individual bits.

}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_88fd8473d9750335 ...bool) *Bitset {
	__obf_ce8d978dad4a08f9 := &Bitset{__obf_b40d2843659b7788: 0, __obf_a98e0f33b9379401: make([]byte, 0)}
	__obf_ce8d978dad4a08f9.
		AppendBools(__obf_88fd8473d9750335...)

	return __obf_ce8d978dad4a08f9
}

// Clone returns a copy.
func Clone(__obf_5a485949ab60b3af *Bitset) *Bitset {
	return &Bitset{__obf_b40d2843659b7788: __obf_5a485949ab60b3af.__obf_b40d2843659b7788,

		// Substr returns a substring, consisting of the bits from indexes start to end.
		__obf_a98e0f33b9379401: __obf_5a485949ab60b3af.__obf_a98e0f33b9379401[:]}
}

func (__obf_ce8d978dad4a08f9 *Bitset) Substr(__obf_ac2631eda8fb742e int, __obf_f2c58c08c4e51323 int) *Bitset {
	if __obf_ac2631eda8fb742e > __obf_f2c58c08c4e51323 || __obf_f2c58c08c4e51323 > __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788 {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_ac2631eda8fb742e, __obf_f2c58c08c4e51323, __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788)
	}
	__obf_7588920eb3600640 := New()
	__obf_7588920eb3600640.__obf_3891c5b40504d5b2(__obf_f2c58c08c4e51323 - __obf_ac2631eda8fb742e)

	for __obf_190a54ebba5c6dbd := __obf_ac2631eda8fb742e; __obf_190a54ebba5c6dbd < __obf_f2c58c08c4e51323; __obf_190a54ebba5c6dbd++ {
		if __obf_ce8d978dad4a08f9.At(__obf_190a54ebba5c6dbd) {
			__obf_7588920eb3600640.__obf_a98e0f33b9379401[__obf_7588920eb3600640.__obf_b40d2843659b7788/8] |= 0x80 >> uint(__obf_7588920eb3600640.__obf_b40d2843659b7788%8)
		}
		__obf_7588920eb3600640.__obf_b40d2843659b7788++
	}

	return __obf_7588920eb3600640
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_dd11079532ba21a3 string) *Bitset {
	__obf_ce8d978dad4a08f9 := &Bitset{__obf_b40d2843659b7788: 0, __obf_a98e0f33b9379401: make([]byte, 0)}

	for _, __obf_27f077878e760398 := range __obf_dd11079532ba21a3 {
		switch __obf_27f077878e760398 {
		case '1':
			__obf_ce8d978dad4a08f9.
				AppendBools(true)
		case '0':
			__obf_ce8d978dad4a08f9.
				AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_27f077878e760398)
		}
	}

	return __obf_ce8d978dad4a08f9
}

// AppendBytes appends a list of whole bytes.
func (__obf_ce8d978dad4a08f9 *Bitset) AppendBytes(__obf_b17981c2861ff761 []byte) {
	for _, __obf_c0cf69cb87dcb29a := range __obf_b17981c2861ff761 {
		__obf_ce8d978dad4a08f9.
			AppendByte(__obf_c0cf69cb87dcb29a, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_ce8d978dad4a08f9 *Bitset) AppendByte(__obf_302acd64ae07b1d1 byte, __obf_b40d2843659b7788 int) {
	__obf_ce8d978dad4a08f9.__obf_3891c5b40504d5b2(__obf_b40d2843659b7788)

	if __obf_b40d2843659b7788 > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_b40d2843659b7788)
	}

	for __obf_190a54ebba5c6dbd := __obf_b40d2843659b7788 - 1; __obf_190a54ebba5c6dbd >= 0; __obf_190a54ebba5c6dbd-- {
		if __obf_302acd64ae07b1d1&(1<<uint(__obf_190a54ebba5c6dbd)) != 0 {
			__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788/8] |= 0x80 >> uint(__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788%8)
		}
		__obf_ce8d978dad4a08f9.

			// AppendUint32 appends the numBits least significant bits from value.
			__obf_b40d2843659b7788++
	}
}

func (__obf_ce8d978dad4a08f9 *Bitset) AppendUint32(__obf_302acd64ae07b1d1 uint32, __obf_b40d2843659b7788 int) {
	__obf_ce8d978dad4a08f9.__obf_3891c5b40504d5b2(__obf_b40d2843659b7788)

	if __obf_b40d2843659b7788 > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_b40d2843659b7788)
	}

	for __obf_190a54ebba5c6dbd := __obf_b40d2843659b7788 - 1; __obf_190a54ebba5c6dbd >= 0; __obf_190a54ebba5c6dbd-- {
		if __obf_302acd64ae07b1d1&(1<<uint(__obf_190a54ebba5c6dbd)) != 0 {
			__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788/8] |= 0x80 >> uint(__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788%8)
		}
		__obf_ce8d978dad4a08f9.

			// ensureCapacity ensures the Bitset can store an additional |numBits|.
			//
			// The underlying array is expanded if necessary. To prevent frequent
			// reallocation, expanding the underlying array at least doubles its capacity.
			__obf_b40d2843659b7788++
	}
}

func (__obf_ce8d978dad4a08f9 *Bitset) __obf_3891c5b40504d5b2(__obf_b40d2843659b7788 int) {
	__obf_b40d2843659b7788 += __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788
	__obf_fcd23d37d0512d32 := __obf_b40d2843659b7788 / 8
	if __obf_b40d2843659b7788%8 != 0 {
		__obf_fcd23d37d0512d32++
	}

	if len(__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401) >= __obf_fcd23d37d0512d32 {
		return
	}
	__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401 = append(__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401, make([]byte, __obf_fcd23d37d0512d32+2*len(__obf_ce8d978dad4a08f9.

		// Append bits copied from |other|.
		//
		// The new length is b.Len() + other.Len().
		__obf_a98e0f33b9379401))...)
}

func (__obf_ce8d978dad4a08f9 *Bitset) Append(__obf_e8824b83eeeef871 *Bitset) {
	__obf_ce8d978dad4a08f9.__obf_3891c5b40504d5b2(__obf_e8824b83eeeef871.__obf_b40d2843659b7788)

	for __obf_190a54ebba5c6dbd := 0; __obf_190a54ebba5c6dbd < __obf_e8824b83eeeef871.__obf_b40d2843659b7788; __obf_190a54ebba5c6dbd++ {
		if __obf_e8824b83eeeef871.At(__obf_190a54ebba5c6dbd) {
			__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788/8] |= 0x80 >> uint(__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788%8)
		}
		__obf_ce8d978dad4a08f9.

			// AppendBools appends bits to the Bitset.
			__obf_b40d2843659b7788++
	}
}

func (__obf_ce8d978dad4a08f9 *Bitset) AppendBools(__obf_a98e0f33b9379401 ...bool) {
	__obf_ce8d978dad4a08f9.__obf_3891c5b40504d5b2(len(__obf_a98e0f33b9379401))

	for _, __obf_88fd8473d9750335 := range __obf_a98e0f33b9379401 {
		if __obf_88fd8473d9750335 {
			__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788/8] |= 0x80 >> uint(__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788%8)
		}
		__obf_ce8d978dad4a08f9.

			// AppendNumBools appends num bits of value value.
			__obf_b40d2843659b7788++
	}
}

func (__obf_ce8d978dad4a08f9 *Bitset) AppendNumBools(__obf_07450b319560c242 int, __obf_302acd64ae07b1d1 bool) {
	for __obf_190a54ebba5c6dbd := 0; __obf_190a54ebba5c6dbd < __obf_07450b319560c242; __obf_190a54ebba5c6dbd++ {
		__obf_ce8d978dad4a08f9.
			AppendBools(__obf_302acd64ae07b1d1)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_ce8d978dad4a08f9 *Bitset) String() string {
	var __obf_c1410dd79428c12b string
	for __obf_190a54ebba5c6dbd := 0; __obf_190a54ebba5c6dbd < __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788; __obf_190a54ebba5c6dbd++ {
		if (__obf_190a54ebba5c6dbd % 8) == 0 {
			__obf_c1410dd79428c12b += " "
		}

		if (__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_190a54ebba5c6dbd/8] & (0x80 >> byte(__obf_190a54ebba5c6dbd%8))) != 0 {
			__obf_c1410dd79428c12b += "1"
		} else {
			__obf_c1410dd79428c12b += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788,

		// Len returns the length of the Bitset in bits.
		__obf_c1410dd79428c12b)
}

func (__obf_ce8d978dad4a08f9 *Bitset) Len() int {
	return __obf_ce8d978dad4a08f9.

		// Bits returns the contents of the Bitset.
		__obf_b40d2843659b7788
}

func (__obf_ce8d978dad4a08f9 *Bitset) Bits() []bool {
	__obf_7588920eb3600640 := make([]bool, __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788)

	var __obf_190a54ebba5c6dbd int
	for __obf_190a54ebba5c6dbd = 0; __obf_190a54ebba5c6dbd < __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788; __obf_190a54ebba5c6dbd++ {
		__obf_7588920eb3600640[__obf_190a54ebba5c6dbd] = (__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_190a54ebba5c6dbd/8] & (0x80 >> byte(__obf_190a54ebba5c6dbd%8))) != 0
	}

	return __obf_7588920eb3600640
}

// At returns the value of the bit at |index|.
func (__obf_ce8d978dad4a08f9 *Bitset) At(__obf_d3883037cd77e231 int) bool {
	if __obf_d3883037cd77e231 >= __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788 {
		log.Panicf("Index %d out of range", __obf_d3883037cd77e231)
	}

	return (__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_d3883037cd77e231/8] & (0x80 >> byte(__obf_d3883037cd77e231%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_ce8d978dad4a08f9 *Bitset) Equals(__obf_e8824b83eeeef871 *Bitset) bool {
	if __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788 != __obf_e8824b83eeeef871.__obf_b40d2843659b7788 {
		return false
	}

	if !bytes.Equal(__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[0:__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788/8], __obf_e8824b83eeeef871.__obf_a98e0f33b9379401[0:__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788/8]) {
		return false
	}

	for __obf_190a54ebba5c6dbd := 8 * (__obf_ce8d978dad4a08f9.__obf_b40d2843659b7788 / 8); __obf_190a54ebba5c6dbd < __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788; __obf_190a54ebba5c6dbd++ {
		__obf_a4612b72fb7e5b62 := (__obf_ce8d978dad4a08f9.__obf_a98e0f33b9379401[__obf_190a54ebba5c6dbd/8] & (0x80 >> byte(__obf_190a54ebba5c6dbd%8)))
		__obf_ce8d978dad4a08f9 := (__obf_e8824b83eeeef871.__obf_a98e0f33b9379401[__obf_190a54ebba5c6dbd/8] & (0x80 >> byte(__obf_190a54ebba5c6dbd%8)))

		if __obf_a4612b72fb7e5b62 != __obf_ce8d978dad4a08f9 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_ce8d978dad4a08f9 *Bitset) ByteAt(__obf_d3883037cd77e231 int) byte {
	if __obf_d3883037cd77e231 < 0 || __obf_d3883037cd77e231 >= __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788 {
		log.Panicf("Index %d out of range", __obf_d3883037cd77e231)
	}

	var __obf_7588920eb3600640 byte

	for __obf_190a54ebba5c6dbd := __obf_d3883037cd77e231; __obf_190a54ebba5c6dbd < __obf_d3883037cd77e231+8 && __obf_190a54ebba5c6dbd < __obf_ce8d978dad4a08f9.__obf_b40d2843659b7788; __obf_190a54ebba5c6dbd++ {
		__obf_7588920eb3600640 <<= 1
		if __obf_ce8d978dad4a08f9.At(__obf_190a54ebba5c6dbd) {
			__obf_7588920eb3600640 |= 1
		}
	}

	return __obf_7588920eb3600640
}
