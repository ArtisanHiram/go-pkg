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
package __obf_b22b687246040e5f

import (
	"bytes"
	"fmt"
	"log"
)

const (
	__obf_d357870681608879 = false
	__obf_52de12b09a90ee06 = true
)

// Bitset stores an array of bits.
type Bitset struct {
	// The number of bits stored.
	__obf_fcf15eabd50d1941 int

	// Storage for individual bits.
	__obf_0ba96c121f3aad06 []byte
}

// New returns an initialised Bitset with optional initial bits v.
func New(__obf_52ca141448ff3e09 ...bool) *Bitset {
	__obf_9606d417ef6362f1 := &Bitset{__obf_fcf15eabd50d1941: 0, __obf_0ba96c121f3aad06: make([]byte, 0)}
	__obf_9606d417ef6362f1.AppendBools(__obf_52ca141448ff3e09...)

	return __obf_9606d417ef6362f1
}

// Clone returns a copy.
func Clone(__obf_534c4474aa6e4559 *Bitset) *Bitset {
	return &Bitset{__obf_fcf15eabd50d1941: __obf_534c4474aa6e4559.__obf_fcf15eabd50d1941, __obf_0ba96c121f3aad06: __obf_534c4474aa6e4559.__obf_0ba96c121f3aad06[:]}
}

// Substr returns a substring, consisting of the bits from indexes start to end.
func (__obf_9606d417ef6362f1 *Bitset) Substr(__obf_e1b3fb603ae18deb int, __obf_1733816dabe35812 int) *Bitset {
	if __obf_e1b3fb603ae18deb > __obf_1733816dabe35812 || __obf_1733816dabe35812 > __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941 {
		log.Panicf("Out of range start=%d end=%d numBits=%d", __obf_e1b3fb603ae18deb, __obf_1733816dabe35812, __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941)
	}

	__obf_7b25f00fae751e0c := New()
	__obf_7b25f00fae751e0c.__obf_34e9d404905ff33d(__obf_1733816dabe35812 - __obf_e1b3fb603ae18deb)

	for __obf_f8676eb2b8769cf2 := __obf_e1b3fb603ae18deb; __obf_f8676eb2b8769cf2 < __obf_1733816dabe35812; __obf_f8676eb2b8769cf2++ {
		if __obf_9606d417ef6362f1.At(__obf_f8676eb2b8769cf2) {
			__obf_7b25f00fae751e0c.__obf_0ba96c121f3aad06[__obf_7b25f00fae751e0c.__obf_fcf15eabd50d1941/8] |= 0x80 >> uint(__obf_7b25f00fae751e0c.__obf_fcf15eabd50d1941%8)
		}
		__obf_7b25f00fae751e0c.__obf_fcf15eabd50d1941++
	}

	return __obf_7b25f00fae751e0c
}

// NewFromBase2String constructs and returns a Bitset from a string. The string
// consists of '1', '0' or ' ' characters, e.g. "1010 0101". The '1' and '0'
// characters represent true/false bits respectively, and ' ' characters are
// ignored.
//
// The function panics if the input string contains other characters.
func NewFromBase2String(__obf_6ea6c95d31e21807 string) *Bitset {
	__obf_9606d417ef6362f1 := &Bitset{__obf_fcf15eabd50d1941: 0, __obf_0ba96c121f3aad06: make([]byte, 0)}

	for _, __obf_d8d831d4796e742e := range __obf_6ea6c95d31e21807 {
		switch __obf_d8d831d4796e742e {
		case '1':
			__obf_9606d417ef6362f1.AppendBools(true)
		case '0':
			__obf_9606d417ef6362f1.AppendBools(false)
		case ' ':
		default:
			log.Panicf("Invalid char %c in NewFromBase2String", __obf_d8d831d4796e742e)
		}
	}

	return __obf_9606d417ef6362f1
}

// AppendBytes appends a list of whole bytes.
func (__obf_9606d417ef6362f1 *Bitset) AppendBytes(__obf_2ed433626e133800 []byte) {
	for _, __obf_58fcd001f43772f3 := range __obf_2ed433626e133800 {
		__obf_9606d417ef6362f1.AppendByte(__obf_58fcd001f43772f3, 8)
	}
}

// AppendByte appends the numBits least significant bits from value.
func (__obf_9606d417ef6362f1 *Bitset) AppendByte(__obf_7ce5e14223dea0e5 byte, __obf_fcf15eabd50d1941 int) {
	__obf_9606d417ef6362f1.__obf_34e9d404905ff33d(__obf_fcf15eabd50d1941)

	if __obf_fcf15eabd50d1941 > 8 {
		log.Panicf("numBits %d out of range 0-8", __obf_fcf15eabd50d1941)
	}

	for __obf_f8676eb2b8769cf2 := __obf_fcf15eabd50d1941 - 1; __obf_f8676eb2b8769cf2 >= 0; __obf_f8676eb2b8769cf2-- {
		if __obf_7ce5e14223dea0e5&(1<<uint(__obf_f8676eb2b8769cf2)) != 0 {
			__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941/8] |= 0x80 >> uint(__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941%8)
		}

		__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941++
	}
}

// AppendUint32 appends the numBits least significant bits from value.
func (__obf_9606d417ef6362f1 *Bitset) AppendUint32(__obf_7ce5e14223dea0e5 uint32, __obf_fcf15eabd50d1941 int) {
	__obf_9606d417ef6362f1.__obf_34e9d404905ff33d(__obf_fcf15eabd50d1941)

	if __obf_fcf15eabd50d1941 > 32 {
		log.Panicf("numBits %d out of range 0-32", __obf_fcf15eabd50d1941)
	}

	for __obf_f8676eb2b8769cf2 := __obf_fcf15eabd50d1941 - 1; __obf_f8676eb2b8769cf2 >= 0; __obf_f8676eb2b8769cf2-- {
		if __obf_7ce5e14223dea0e5&(1<<uint(__obf_f8676eb2b8769cf2)) != 0 {
			__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941/8] |= 0x80 >> uint(__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941%8)
		}

		__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941++
	}
}

// ensureCapacity ensures the Bitset can store an additional |numBits|.
//
// The underlying array is expanded if necessary. To prevent frequent
// reallocation, expanding the underlying array at least doubles its capacity.
func (__obf_9606d417ef6362f1 *Bitset) __obf_34e9d404905ff33d(__obf_fcf15eabd50d1941 int) {
	__obf_fcf15eabd50d1941 += __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941

	__obf_5e7ae74ad6f52419 := __obf_fcf15eabd50d1941 / 8
	if __obf_fcf15eabd50d1941%8 != 0 {
		__obf_5e7ae74ad6f52419++
	}

	if len(__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06) >= __obf_5e7ae74ad6f52419 {
		return
	}

	__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06 = append(__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06, make([]byte, __obf_5e7ae74ad6f52419+2*len(__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06))...)
}

// Append bits copied from |other|.
//
// The new length is b.Len() + other.Len().
func (__obf_9606d417ef6362f1 *Bitset) Append(__obf_a2086c4431bdc29e *Bitset) {
	__obf_9606d417ef6362f1.__obf_34e9d404905ff33d(__obf_a2086c4431bdc29e.__obf_fcf15eabd50d1941)

	for __obf_f8676eb2b8769cf2 := 0; __obf_f8676eb2b8769cf2 < __obf_a2086c4431bdc29e.__obf_fcf15eabd50d1941; __obf_f8676eb2b8769cf2++ {
		if __obf_a2086c4431bdc29e.At(__obf_f8676eb2b8769cf2) {
			__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941/8] |= 0x80 >> uint(__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941%8)
		}
		__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941++
	}
}

// AppendBools appends bits to the Bitset.
func (__obf_9606d417ef6362f1 *Bitset) AppendBools(__obf_0ba96c121f3aad06 ...bool) {
	__obf_9606d417ef6362f1.__obf_34e9d404905ff33d(len(__obf_0ba96c121f3aad06))

	for _, __obf_52ca141448ff3e09 := range __obf_0ba96c121f3aad06 {
		if __obf_52ca141448ff3e09 {
			__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941/8] |= 0x80 >> uint(__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941%8)
		}
		__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941++
	}
}

// AppendNumBools appends num bits of value value.
func (__obf_9606d417ef6362f1 *Bitset) AppendNumBools(__obf_e1c63d5e783cf2b5 int, __obf_7ce5e14223dea0e5 bool) {
	for __obf_f8676eb2b8769cf2 := 0; __obf_f8676eb2b8769cf2 < __obf_e1c63d5e783cf2b5; __obf_f8676eb2b8769cf2++ {
		__obf_9606d417ef6362f1.AppendBools(__obf_7ce5e14223dea0e5)
	}
}

// String returns a human readable representation of the Bitset's contents.
func (__obf_9606d417ef6362f1 *Bitset) String() string {
	var __obf_325727a8bffab490 string
	for __obf_f8676eb2b8769cf2 := 0; __obf_f8676eb2b8769cf2 < __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941; __obf_f8676eb2b8769cf2++ {
		if (__obf_f8676eb2b8769cf2 % 8) == 0 {
			__obf_325727a8bffab490 += " "
		}

		if (__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_f8676eb2b8769cf2/8] & (0x80 >> byte(__obf_f8676eb2b8769cf2%8))) != 0 {
			__obf_325727a8bffab490 += "1"
		} else {
			__obf_325727a8bffab490 += "0"
		}
	}

	return fmt.Sprintf("numBits=%d, bits=%s", __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941, __obf_325727a8bffab490)
}

// Len returns the length of the Bitset in bits.
func (__obf_9606d417ef6362f1 *Bitset) Len() int {
	return __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941
}

// Bits returns the contents of the Bitset.
func (__obf_9606d417ef6362f1 *Bitset) Bits() []bool {
	__obf_7b25f00fae751e0c := make([]bool, __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941)

	var __obf_f8676eb2b8769cf2 int
	for __obf_f8676eb2b8769cf2 = 0; __obf_f8676eb2b8769cf2 < __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941; __obf_f8676eb2b8769cf2++ {
		__obf_7b25f00fae751e0c[__obf_f8676eb2b8769cf2] = (__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_f8676eb2b8769cf2/8] & (0x80 >> byte(__obf_f8676eb2b8769cf2%8))) != 0
	}

	return __obf_7b25f00fae751e0c
}

// At returns the value of the bit at |index|.
func (__obf_9606d417ef6362f1 *Bitset) At(__obf_55f97f5d28882ea4 int) bool {
	if __obf_55f97f5d28882ea4 >= __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941 {
		log.Panicf("Index %d out of range", __obf_55f97f5d28882ea4)
	}

	return (__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_55f97f5d28882ea4/8] & (0x80 >> byte(__obf_55f97f5d28882ea4%8))) != 0
}

// Equals returns true if the Bitset equals other.
func (__obf_9606d417ef6362f1 *Bitset) Equals(__obf_a2086c4431bdc29e *Bitset) bool {
	if __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941 != __obf_a2086c4431bdc29e.__obf_fcf15eabd50d1941 {
		return false
	}

	if !bytes.Equal(__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[0:__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941/8], __obf_a2086c4431bdc29e.__obf_0ba96c121f3aad06[0:__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941/8]) {
		return false
	}

	for __obf_f8676eb2b8769cf2 := 8 * (__obf_9606d417ef6362f1.__obf_fcf15eabd50d1941 / 8); __obf_f8676eb2b8769cf2 < __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941; __obf_f8676eb2b8769cf2++ {
		__obf_166e23067ab7ab7f := (__obf_9606d417ef6362f1.__obf_0ba96c121f3aad06[__obf_f8676eb2b8769cf2/8] & (0x80 >> byte(__obf_f8676eb2b8769cf2%8)))
		__obf_9606d417ef6362f1 := (__obf_a2086c4431bdc29e.__obf_0ba96c121f3aad06[__obf_f8676eb2b8769cf2/8] & (0x80 >> byte(__obf_f8676eb2b8769cf2%8)))

		if __obf_166e23067ab7ab7f != __obf_9606d417ef6362f1 {
			return false
		}
	}

	return true
}

// ByteAt returns a byte consisting of upto 8 bits starting at index.
func (__obf_9606d417ef6362f1 *Bitset) ByteAt(__obf_55f97f5d28882ea4 int) byte {
	if __obf_55f97f5d28882ea4 < 0 || __obf_55f97f5d28882ea4 >= __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941 {
		log.Panicf("Index %d out of range", __obf_55f97f5d28882ea4)
	}

	var __obf_7b25f00fae751e0c byte

	for __obf_f8676eb2b8769cf2 := __obf_55f97f5d28882ea4; __obf_f8676eb2b8769cf2 < __obf_55f97f5d28882ea4+8 && __obf_f8676eb2b8769cf2 < __obf_9606d417ef6362f1.__obf_fcf15eabd50d1941; __obf_f8676eb2b8769cf2++ {
		__obf_7b25f00fae751e0c <<= 1
		if __obf_9606d417ef6362f1.At(__obf_f8676eb2b8769cf2) {
			__obf_7b25f00fae751e0c |= 1
		}
	}

	return __obf_7b25f00fae751e0c
}
