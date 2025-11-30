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
package __obf_398820583db81f0b

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_c9c4efc0550c45f4 = false
	__obf_cb72ac9cdf0d4086 = true
)

// Bitset stores an array of bits.
type Bitset struct {
	__obf_65933c337060b280 int // The number of bits stored.
	__obf_7cced0d732b68faa []byte// Storage for individual bits.

}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_5a0db15d4d044eaa ...bool) *Bitset {
	__obf_603da838f6d2be62 := &Bitset{__obf_65933c337060b280: 0, __obf_7cced0d732b68faa: make([]byte, 0)}
	__obf_603da838f6d2be62.
		AppendBools(__obf_5a0db15d4d044eaa...)

	return __obf_603da838f6d2be62
}

// Clone returns a copy.
func Clone(__obf_acae70369510e9ca *Bitset) *Bitset {
	return &Bitset{__obf_65933c337060b280: __obf_acae70369510e9ca.__obf_65933c337060b280,

		// Substr returns a substring, consisting of the bits from indexes start to end.
		__obf_7cced0d732b68faa: __obf_acae70369510e9ca.__obf_7cced0d732b68faa[:]}
}

func (__obf_603da838f6d2be62 *Bitset) Substr(__obf_6e2017ca8140cbf5 int, __obf_09ab9b819b7184da int) *Bitset {
	if __obf_6e2017ca8140cbf5 > __obf_09ab9b819b7184da || __obf_09ab9b819b7184da > __obf_603da838f6d2be62.__obf_65933c337060b280 {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_6e2017ca8140cbf5, __obf_09ab9b819b7184da, __obf_603da838f6d2be62.__obf_65933c337060b280)
	}
	__obf_ac2c247125beb5e3 := New()
	__obf_ac2c247125beb5e3.__obf_360842d0953f7673(__obf_09ab9b819b7184da - __obf_6e2017ca8140cbf5)

	for __obf_06aa0aefebeb4d0e := __obf_6e2017ca8140cbf5; __obf_06aa0aefebeb4d0e < __obf_09ab9b819b7184da; __obf_06aa0aefebeb4d0e++ {
		if __obf_603da838f6d2be62.At(__obf_06aa0aefebeb4d0e) {
			__obf_ac2c247125beb5e3.__obf_7cced0d732b68faa[__obf_ac2c247125beb5e3.__obf_65933c337060b280/8] |= 0x80 >> uint(__obf_ac2c247125beb5e3.__obf_65933c337060b280%8)
		}
		__obf_ac2c247125beb5e3.__obf_65933c337060b280++
	}

	return __obf_ac2c247125beb5e3
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_58feb0540864c644 string) *Bitset {
	__obf_603da838f6d2be62 := &Bitset{__obf_65933c337060b280: 0, __obf_7cced0d732b68faa: make([]byte, 0)}

	for _, __obf_fc1e37010fdeac08 := range __obf_58feb0540864c644 {
		switch __obf_fc1e37010fdeac08 {
		case '1':
			__obf_603da838f6d2be62.
				AppendBools(true)
		case '0':
			__obf_603da838f6d2be62.
				AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_fc1e37010fdeac08)
		}
	}

	return __obf_603da838f6d2be62
}

// AppendBytes appends a list of whole bytes.
func (__obf_603da838f6d2be62 *Bitset) AppendBytes(__obf_f562011de6b4ca24 []byte) {
	for _, __obf_258b0a3956abee9e := range __obf_f562011de6b4ca24 {
		__obf_603da838f6d2be62.
			AppendByte(__obf_258b0a3956abee9e, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_603da838f6d2be62 *Bitset) AppendByte(__obf_3e4678302f69bc01 byte, __obf_65933c337060b280 int) {
	__obf_603da838f6d2be62.__obf_360842d0953f7673(__obf_65933c337060b280)

	if __obf_65933c337060b280 > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_65933c337060b280)
	}

	for __obf_06aa0aefebeb4d0e := __obf_65933c337060b280 - 1; __obf_06aa0aefebeb4d0e >= 0; __obf_06aa0aefebeb4d0e-- {
		if __obf_3e4678302f69bc01&(1<<uint(__obf_06aa0aefebeb4d0e)) != 0 {
			__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_603da838f6d2be62.__obf_65933c337060b280/8] |= 0x80 >> uint(__obf_603da838f6d2be62.__obf_65933c337060b280%8)
		}
		__obf_603da838f6d2be62.

			// AppendUint32 appends the numBits least significant bits from value.
			__obf_65933c337060b280++
	}
}

func (__obf_603da838f6d2be62 *Bitset) AppendUint32(__obf_3e4678302f69bc01 uint32, __obf_65933c337060b280 int) {
	__obf_603da838f6d2be62.__obf_360842d0953f7673(__obf_65933c337060b280)

	if __obf_65933c337060b280 > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_65933c337060b280)
	}

	for __obf_06aa0aefebeb4d0e := __obf_65933c337060b280 - 1; __obf_06aa0aefebeb4d0e >= 0; __obf_06aa0aefebeb4d0e-- {
		if __obf_3e4678302f69bc01&(1<<uint(__obf_06aa0aefebeb4d0e)) != 0 {
			__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_603da838f6d2be62.__obf_65933c337060b280/8] |= 0x80 >> uint(__obf_603da838f6d2be62.__obf_65933c337060b280%8)
		}
		__obf_603da838f6d2be62.

			// ensureCapacity ensures the Bitset can store an additional |numBits|.
			//
			// The underlying array is expanded if necessary. To prevent frequent
			// reallocation, expanding the underlying array at least doubles its capacity.
			__obf_65933c337060b280++
	}
}

func (__obf_603da838f6d2be62 *Bitset) __obf_360842d0953f7673(__obf_65933c337060b280 int) {
	__obf_65933c337060b280 += __obf_603da838f6d2be62.__obf_65933c337060b280
	__obf_8a72a823317d8860 := __obf_65933c337060b280 / 8
	if __obf_65933c337060b280%8 != 0 {
		__obf_8a72a823317d8860++
	}

	if len(__obf_603da838f6d2be62.__obf_7cced0d732b68faa) >= __obf_8a72a823317d8860 {
		return
	}
	__obf_603da838f6d2be62.__obf_7cced0d732b68faa = append(__obf_603da838f6d2be62.__obf_7cced0d732b68faa, make([]byte, __obf_8a72a823317d8860+2*len(__obf_603da838f6d2be62.

		// Append bits copied from |other|.
		//
		// The new length is b.Len() + other.Len().
		__obf_7cced0d732b68faa))...)
}

func (__obf_603da838f6d2be62 *Bitset) Append(__obf_73b72717567272f8 *Bitset) {
	__obf_603da838f6d2be62.__obf_360842d0953f7673(__obf_73b72717567272f8.__obf_65933c337060b280)

	for __obf_06aa0aefebeb4d0e := 0; __obf_06aa0aefebeb4d0e < __obf_73b72717567272f8.__obf_65933c337060b280; __obf_06aa0aefebeb4d0e++ {
		if __obf_73b72717567272f8.At(__obf_06aa0aefebeb4d0e) {
			__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_603da838f6d2be62.__obf_65933c337060b280/8] |= 0x80 >> uint(__obf_603da838f6d2be62.__obf_65933c337060b280%8)
		}
		__obf_603da838f6d2be62.

			// AppendBools appends bits to the Bitset.
			__obf_65933c337060b280++
	}
}

func (__obf_603da838f6d2be62 *Bitset) AppendBools(__obf_7cced0d732b68faa ...bool) {
	__obf_603da838f6d2be62.__obf_360842d0953f7673(len(__obf_7cced0d732b68faa))

	for _, __obf_5a0db15d4d044eaa := range __obf_7cced0d732b68faa {
		if __obf_5a0db15d4d044eaa {
			__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_603da838f6d2be62.__obf_65933c337060b280/8] |= 0x80 >> uint(__obf_603da838f6d2be62.__obf_65933c337060b280%8)
		}
		__obf_603da838f6d2be62.

			// AppendNumBools appends num bits of value value.
			__obf_65933c337060b280++
	}
}

func (__obf_603da838f6d2be62 *Bitset) AppendNumBools(__obf_36d4e6f313c88552 int, __obf_3e4678302f69bc01 bool) {
	for __obf_06aa0aefebeb4d0e := 0; __obf_06aa0aefebeb4d0e < __obf_36d4e6f313c88552; __obf_06aa0aefebeb4d0e++ {
		__obf_603da838f6d2be62.
			AppendBools(__obf_3e4678302f69bc01)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_603da838f6d2be62 *Bitset) String() string {
	var __obf_7a3abef36de513ee string
	for __obf_06aa0aefebeb4d0e := 0; __obf_06aa0aefebeb4d0e < __obf_603da838f6d2be62.__obf_65933c337060b280; __obf_06aa0aefebeb4d0e++ {
		if (__obf_06aa0aefebeb4d0e % 8) == 0 {
			__obf_7a3abef36de513ee += " "
		}

		if (__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_06aa0aefebeb4d0e/8] & (0x80 >> byte(__obf_06aa0aefebeb4d0e%8))) != 0 {
			__obf_7a3abef36de513ee += "1"
		} else {
			__obf_7a3abef36de513ee += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_603da838f6d2be62.__obf_65933c337060b280,

		// Len returns the length of the Bitset in bits.
		__obf_7a3abef36de513ee)
}

func (__obf_603da838f6d2be62 *Bitset) Len() int {
	return __obf_603da838f6d2be62.

		// Bits returns the contents of the Bitset.
		__obf_65933c337060b280
}

func (__obf_603da838f6d2be62 *Bitset) Bits() []bool {
	__obf_ac2c247125beb5e3 := make([]bool, __obf_603da838f6d2be62.__obf_65933c337060b280)

	var __obf_06aa0aefebeb4d0e int
	for __obf_06aa0aefebeb4d0e = 0; __obf_06aa0aefebeb4d0e < __obf_603da838f6d2be62.__obf_65933c337060b280; __obf_06aa0aefebeb4d0e++ {
		__obf_ac2c247125beb5e3[__obf_06aa0aefebeb4d0e] = (__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_06aa0aefebeb4d0e/8] & (0x80 >> byte(__obf_06aa0aefebeb4d0e%8))) != 0
	}

	return __obf_ac2c247125beb5e3
}

// At returns the value of the bit at |index|.
func (__obf_603da838f6d2be62 *Bitset) At(__obf_d1d1e246f38f48a2 int) bool {
	if __obf_d1d1e246f38f48a2 >= __obf_603da838f6d2be62.__obf_65933c337060b280 {
		log.Panicf("Index %d out of range", __obf_d1d1e246f38f48a2)
	}

	return (__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_d1d1e246f38f48a2/8] & (0x80 >> byte(__obf_d1d1e246f38f48a2%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_603da838f6d2be62 *Bitset) Equals(__obf_73b72717567272f8 *Bitset) bool {
	if __obf_603da838f6d2be62.__obf_65933c337060b280 != __obf_73b72717567272f8.__obf_65933c337060b280 {
		return false
	}

	if !bytes.Equal(__obf_603da838f6d2be62.__obf_7cced0d732b68faa[0:__obf_603da838f6d2be62.__obf_65933c337060b280/8], __obf_73b72717567272f8.__obf_7cced0d732b68faa[0:__obf_603da838f6d2be62.__obf_65933c337060b280/8]) {
		return false
	}

	for __obf_06aa0aefebeb4d0e := 8 * (__obf_603da838f6d2be62.__obf_65933c337060b280 / 8); __obf_06aa0aefebeb4d0e < __obf_603da838f6d2be62.__obf_65933c337060b280; __obf_06aa0aefebeb4d0e++ {
		__obf_2c2ec262f0459cc4 := (__obf_603da838f6d2be62.__obf_7cced0d732b68faa[__obf_06aa0aefebeb4d0e/8] & (0x80 >> byte(__obf_06aa0aefebeb4d0e%8)))
		__obf_603da838f6d2be62 := (__obf_73b72717567272f8.__obf_7cced0d732b68faa[__obf_06aa0aefebeb4d0e/8] & (0x80 >> byte(__obf_06aa0aefebeb4d0e%8)))

		if __obf_2c2ec262f0459cc4 != __obf_603da838f6d2be62 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_603da838f6d2be62 *Bitset) ByteAt(__obf_d1d1e246f38f48a2 int) byte {
	if __obf_d1d1e246f38f48a2 < 0 || __obf_d1d1e246f38f48a2 >= __obf_603da838f6d2be62.__obf_65933c337060b280 {
		log.Panicf("Index %d out of range", __obf_d1d1e246f38f48a2)
	}

	var __obf_ac2c247125beb5e3 byte

	for __obf_06aa0aefebeb4d0e := __obf_d1d1e246f38f48a2; __obf_06aa0aefebeb4d0e < __obf_d1d1e246f38f48a2+8 && __obf_06aa0aefebeb4d0e < __obf_603da838f6d2be62.__obf_65933c337060b280; __obf_06aa0aefebeb4d0e++ {
		__obf_ac2c247125beb5e3 <<= 1
		if __obf_603da838f6d2be62.At(__obf_06aa0aefebeb4d0e) {
			__obf_ac2c247125beb5e3 |= 1
		}
	}

	return __obf_ac2c247125beb5e3
}
