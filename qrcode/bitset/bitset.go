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
package __obf_dfae9ab71459be69

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_695bdaade6ebefac = false
	__obf_d2ba4fef21b4a5c2 = true
)

// Bitset stores an array of bits.
type Bitset struct {
	__obf_39ec6fec1493133f int // The number of bits stored.
	__obf_098303cec33c6c28 []byte// Storage for individual bits.

}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_090a3107cf018c88 ...bool) *Bitset {
	__obf_28037982d5226392 := &Bitset{__obf_39ec6fec1493133f: 0, __obf_098303cec33c6c28: make([]byte, 0)}
	__obf_28037982d5226392.
		AppendBools(__obf_090a3107cf018c88...)

	return __obf_28037982d5226392
}

// Clone returns a copy.
func Clone(__obf_e799d92e8523d5b1 *Bitset) *Bitset {
	return &Bitset{__obf_39ec6fec1493133f: __obf_e799d92e8523d5b1.__obf_39ec6fec1493133f,

		// Substr returns a substring, consisting of the bits from indexes start to end.
		__obf_098303cec33c6c28: __obf_e799d92e8523d5b1.__obf_098303cec33c6c28[:]}
}

func (__obf_28037982d5226392 *Bitset) Substr(__obf_c00d3e424287b3b8 int, __obf_e9aece29df5a968f int) *Bitset {
	if __obf_c00d3e424287b3b8 > __obf_e9aece29df5a968f || __obf_e9aece29df5a968f > __obf_28037982d5226392.__obf_39ec6fec1493133f {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_c00d3e424287b3b8, __obf_e9aece29df5a968f, __obf_28037982d5226392.__obf_39ec6fec1493133f)
	}
	__obf_3c0068a9f5d72474 := New()
	__obf_3c0068a9f5d72474.__obf_732255969f9970ba(__obf_e9aece29df5a968f - __obf_c00d3e424287b3b8)

	for __obf_7342c91c4b631f5e := __obf_c00d3e424287b3b8; __obf_7342c91c4b631f5e < __obf_e9aece29df5a968f; __obf_7342c91c4b631f5e++ {
		if __obf_28037982d5226392.At(__obf_7342c91c4b631f5e) {
			__obf_3c0068a9f5d72474.__obf_098303cec33c6c28[__obf_3c0068a9f5d72474.__obf_39ec6fec1493133f/8] |= 0x80 >> uint(__obf_3c0068a9f5d72474.__obf_39ec6fec1493133f%8)
		}
		__obf_3c0068a9f5d72474.__obf_39ec6fec1493133f++
	}

	return __obf_3c0068a9f5d72474
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_bd270c6dbf7f9b99 string) *Bitset {
	__obf_28037982d5226392 := &Bitset{__obf_39ec6fec1493133f: 0, __obf_098303cec33c6c28: make([]byte, 0)}

	for _, __obf_457a6690ad2dd2ec := range __obf_bd270c6dbf7f9b99 {
		switch __obf_457a6690ad2dd2ec {
		case '1':
			__obf_28037982d5226392.
				AppendBools(true)
		case '0':
			__obf_28037982d5226392.
				AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_457a6690ad2dd2ec)
		}
	}

	return __obf_28037982d5226392
}

// AppendBytes appends a list of whole bytes.
func (__obf_28037982d5226392 *Bitset) AppendBytes(__obf_295a9d3ab715b537 []byte) {
	for _, __obf_d081c3f49fd8110f := range __obf_295a9d3ab715b537 {
		__obf_28037982d5226392.
			AppendByte(__obf_d081c3f49fd8110f, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_28037982d5226392 *Bitset) AppendByte(__obf_71e5b97a66b551af byte, __obf_39ec6fec1493133f int) {
	__obf_28037982d5226392.__obf_732255969f9970ba(__obf_39ec6fec1493133f)

	if __obf_39ec6fec1493133f > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_39ec6fec1493133f)
	}

	for __obf_7342c91c4b631f5e := __obf_39ec6fec1493133f - 1; __obf_7342c91c4b631f5e >= 0; __obf_7342c91c4b631f5e-- {
		if __obf_71e5b97a66b551af&(1<<uint(__obf_7342c91c4b631f5e)) != 0 {
			__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_28037982d5226392.__obf_39ec6fec1493133f/8] |= 0x80 >> uint(__obf_28037982d5226392.__obf_39ec6fec1493133f%8)
		}
		__obf_28037982d5226392.

			// AppendUint32 appends the numBits least significant bits from value.
			__obf_39ec6fec1493133f++
	}
}

func (__obf_28037982d5226392 *Bitset) AppendUint32(__obf_71e5b97a66b551af uint32, __obf_39ec6fec1493133f int) {
	__obf_28037982d5226392.__obf_732255969f9970ba(__obf_39ec6fec1493133f)

	if __obf_39ec6fec1493133f > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_39ec6fec1493133f)
	}

	for __obf_7342c91c4b631f5e := __obf_39ec6fec1493133f - 1; __obf_7342c91c4b631f5e >= 0; __obf_7342c91c4b631f5e-- {
		if __obf_71e5b97a66b551af&(1<<uint(__obf_7342c91c4b631f5e)) != 0 {
			__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_28037982d5226392.__obf_39ec6fec1493133f/8] |= 0x80 >> uint(__obf_28037982d5226392.__obf_39ec6fec1493133f%8)
		}
		__obf_28037982d5226392.

			// ensureCapacity ensures the Bitset can store an additional |numBits|.
			//
			// The underlying array is expanded if necessary. To prevent frequent
			// reallocation, expanding the underlying array at least doubles its capacity.
			__obf_39ec6fec1493133f++
	}
}

func (__obf_28037982d5226392 *Bitset) __obf_732255969f9970ba(__obf_39ec6fec1493133f int) {
	__obf_39ec6fec1493133f += __obf_28037982d5226392.__obf_39ec6fec1493133f
	__obf_f86d90d2f881f1f3 := __obf_39ec6fec1493133f / 8
	if __obf_39ec6fec1493133f%8 != 0 {
		__obf_f86d90d2f881f1f3++
	}

	if len(__obf_28037982d5226392.__obf_098303cec33c6c28) >= __obf_f86d90d2f881f1f3 {
		return
	}
	__obf_28037982d5226392.__obf_098303cec33c6c28 = append(__obf_28037982d5226392.__obf_098303cec33c6c28, make([]byte, __obf_f86d90d2f881f1f3+2*len(__obf_28037982d5226392.

		// Append bits copied from |other|.
		//
		// The new length is b.Len() + other.Len().
		__obf_098303cec33c6c28))...)
}

func (__obf_28037982d5226392 *Bitset) Append(__obf_e4de3d9f9684ad71 *Bitset) {
	__obf_28037982d5226392.__obf_732255969f9970ba(__obf_e4de3d9f9684ad71.__obf_39ec6fec1493133f)

	for __obf_7342c91c4b631f5e := 0; __obf_7342c91c4b631f5e < __obf_e4de3d9f9684ad71.__obf_39ec6fec1493133f; __obf_7342c91c4b631f5e++ {
		if __obf_e4de3d9f9684ad71.At(__obf_7342c91c4b631f5e) {
			__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_28037982d5226392.__obf_39ec6fec1493133f/8] |= 0x80 >> uint(__obf_28037982d5226392.__obf_39ec6fec1493133f%8)
		}
		__obf_28037982d5226392.

			// AppendBools appends bits to the Bitset.
			__obf_39ec6fec1493133f++
	}
}

func (__obf_28037982d5226392 *Bitset) AppendBools(__obf_098303cec33c6c28 ...bool) {
	__obf_28037982d5226392.__obf_732255969f9970ba(len(__obf_098303cec33c6c28))

	for _, __obf_090a3107cf018c88 := range __obf_098303cec33c6c28 {
		if __obf_090a3107cf018c88 {
			__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_28037982d5226392.__obf_39ec6fec1493133f/8] |= 0x80 >> uint(__obf_28037982d5226392.__obf_39ec6fec1493133f%8)
		}
		__obf_28037982d5226392.

			// AppendNumBools appends num bits of value value.
			__obf_39ec6fec1493133f++
	}
}

func (__obf_28037982d5226392 *Bitset) AppendNumBools(__obf_38f208ddaf1c7de4 int, __obf_71e5b97a66b551af bool) {
	for __obf_7342c91c4b631f5e := 0; __obf_7342c91c4b631f5e < __obf_38f208ddaf1c7de4; __obf_7342c91c4b631f5e++ {
		__obf_28037982d5226392.
			AppendBools(__obf_71e5b97a66b551af)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_28037982d5226392 *Bitset) String() string {
	var __obf_ead1bf0878c38198 string
	for __obf_7342c91c4b631f5e := 0; __obf_7342c91c4b631f5e < __obf_28037982d5226392.__obf_39ec6fec1493133f; __obf_7342c91c4b631f5e++ {
		if (__obf_7342c91c4b631f5e % 8) == 0 {
			__obf_ead1bf0878c38198 += " "
		}

		if (__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_7342c91c4b631f5e/8] & (0x80 >> byte(__obf_7342c91c4b631f5e%8))) != 0 {
			__obf_ead1bf0878c38198 += "1"
		} else {
			__obf_ead1bf0878c38198 += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_28037982d5226392.__obf_39ec6fec1493133f,

		// Len returns the length of the Bitset in bits.
		__obf_ead1bf0878c38198)
}

func (__obf_28037982d5226392 *Bitset) Len() int {
	return __obf_28037982d5226392.

		// Bits returns the contents of the Bitset.
		__obf_39ec6fec1493133f
}

func (__obf_28037982d5226392 *Bitset) Bits() []bool {
	__obf_3c0068a9f5d72474 := make([]bool, __obf_28037982d5226392.__obf_39ec6fec1493133f)

	var __obf_7342c91c4b631f5e int
	for __obf_7342c91c4b631f5e = 0; __obf_7342c91c4b631f5e < __obf_28037982d5226392.__obf_39ec6fec1493133f; __obf_7342c91c4b631f5e++ {
		__obf_3c0068a9f5d72474[__obf_7342c91c4b631f5e] = (__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_7342c91c4b631f5e/8] & (0x80 >> byte(__obf_7342c91c4b631f5e%8))) != 0
	}

	return __obf_3c0068a9f5d72474
}

// At returns the value of the bit at |index|.
func (__obf_28037982d5226392 *Bitset) At(__obf_727b9d53647adbea int) bool {
	if __obf_727b9d53647adbea >= __obf_28037982d5226392.__obf_39ec6fec1493133f {
		log.Panicf("Index %d out of range", __obf_727b9d53647adbea)
	}

	return (__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_727b9d53647adbea/8] & (0x80 >> byte(__obf_727b9d53647adbea%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_28037982d5226392 *Bitset) Equals(__obf_e4de3d9f9684ad71 *Bitset) bool {
	if __obf_28037982d5226392.__obf_39ec6fec1493133f != __obf_e4de3d9f9684ad71.__obf_39ec6fec1493133f {
		return false
	}

	if !bytes.Equal(__obf_28037982d5226392.__obf_098303cec33c6c28[0:__obf_28037982d5226392.__obf_39ec6fec1493133f/8], __obf_e4de3d9f9684ad71.__obf_098303cec33c6c28[0:__obf_28037982d5226392.__obf_39ec6fec1493133f/8]) {
		return false
	}

	for __obf_7342c91c4b631f5e := 8 * (__obf_28037982d5226392.__obf_39ec6fec1493133f / 8); __obf_7342c91c4b631f5e < __obf_28037982d5226392.__obf_39ec6fec1493133f; __obf_7342c91c4b631f5e++ {
		__obf_e525015d605cc293 := (__obf_28037982d5226392.__obf_098303cec33c6c28[__obf_7342c91c4b631f5e/8] & (0x80 >> byte(__obf_7342c91c4b631f5e%8)))
		__obf_28037982d5226392 := (__obf_e4de3d9f9684ad71.__obf_098303cec33c6c28[__obf_7342c91c4b631f5e/8] & (0x80 >> byte(__obf_7342c91c4b631f5e%8)))

		if __obf_e525015d605cc293 != __obf_28037982d5226392 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_28037982d5226392 *Bitset) ByteAt(__obf_727b9d53647adbea int) byte {
	if __obf_727b9d53647adbea < 0 || __obf_727b9d53647adbea >= __obf_28037982d5226392.__obf_39ec6fec1493133f {
		log.Panicf("Index %d out of range", __obf_727b9d53647adbea)
	}

	var __obf_3c0068a9f5d72474 byte

	for __obf_7342c91c4b631f5e := __obf_727b9d53647adbea; __obf_7342c91c4b631f5e < __obf_727b9d53647adbea+8 && __obf_7342c91c4b631f5e < __obf_28037982d5226392.__obf_39ec6fec1493133f; __obf_7342c91c4b631f5e++ {
		__obf_3c0068a9f5d72474 <<= 1
		if __obf_28037982d5226392.At(__obf_7342c91c4b631f5e) {
			__obf_3c0068a9f5d72474 |= 1
		}
	}

	return __obf_3c0068a9f5d72474
}
