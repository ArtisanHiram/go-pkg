// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_8d73621f5856b288

import (
	"errors"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// Data encoding.
//
// The main data portion of a QR Code consists of one or more segments of data.
// A segment consists of:
//
// - The segment Data Mode: numeric, alphanumeric, or byte.
// - The length of segment in bits.
// - Encoded data.
//
// For example, the string "123ZZ#!#!" may be represented as:
//
// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"]
//
// Multiple data modes exist to minimise the size of encoded data. For example,
// 8-bit bytes require 8 bits to encode each, but base 10 numeric data can be
// encoded at a higher density of 3 numbers (e.g. 123) per 10 bits.
//
// Some data can be represented in multiple modes. Numeric data can be
// represented in all three modes, whereas alphanumeric data (e.g. 'A') can be
// represented in alphanumeric and byte mode.
//
// Starting a new segment (to use a different Data Mode) has a cost, the bits to
// state the new segment Data Mode and length. To minimise each QR Code's symbol
// size, an optimisation routine coalesces segment types where possible, to
// reduce the encoded data length.
//
// There are several other data modes available (e.g. Kanji mode) which are not
// implemented here.

// A segment encoding mode.
type __obf_51964cb459dfa8c2 uint8

const (
	__obf_7e29f0a1d309f15b __obf_51964cb459dfa8c2 =
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	1 << iota
	__obf_2d9adca192f6544c
	__obf_06a20756bb412f63
	__obf_98f9ed48ce6f730b
)

// dataModeString returns d as a short printable string.
func __obf_ea3b8b9622cbc6ee(__obf_bbb97c28ddd95a60 __obf_51964cb459dfa8c2) string {
	switch __obf_bbb97c28ddd95a60 {
	case __obf_7e29f0a1d309f15b:
		return "none"
	case __obf_2d9adca192f6544c:
		return "numeric"
	case __obf_06a20756bb412f63:
		return "alphanumeric"
	case __obf_98f9ed48ce6f730b:
		return "byte"
	}

	return "unknown"
}

type __obf_a9d4704514b6286c uint8

const (
	__obf_80807e9626ee48fa __obf_a9d4704514b6286c = iota
	__obf_9bb2ab43c09df8ad
	__obf_7f04905a311519fa
)

// segment is a single segment of data.
type __obf_e076203f11833315 struct {
	__obf_51964cb459dfa8c2 __obf_51964cb459dfa8c2
	// Data Mode (e.g. numeric).
	__obf_7a0637e203e4ea7a []byte// segment data (e.g. "abc").

}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_db79b5bc73d5a614 struct {
	__obf_1423e711ab2d7104 int // Minimum & maximum versions supported.
	__obf_a9a6d1fd09e8597b int
	__obf_8213f88df91d8826 * // Mode indicator bit sequences.
	bitset.Bitset
	__obf_94407096894974c2 *bitset.Bitset
	__obf_32106c8c49ebde5d *bitset.Bitset
	__obf_9a03c960dad69649 int // Character count lengths.
	__obf_d05ff65d77bfc31e int
	__obf_2a2cbaecb5ec2f61 int
	__obf_7a0637e203e4ea7a []byte// The raw input data.

	__obf_bd9515c26418e69f []__obf_e076203f11833315// The data classified into unoptimised segments.

	// The data classified into optimised segments.
	__obf_4fa8b1538f6edfcb []__obf_e076203f11833315
}

// newDataEncoder constructs a dataEncoder.
func __obf_2a2ab161febefc8e(__obf_33479cec7dd55141 __obf_a9d4704514b6286c) *__obf_db79b5bc73d5a614 {
	__obf_bbb97c28ddd95a60 := &__obf_db79b5bc73d5a614{}

	switch __obf_33479cec7dd55141 {
	case __obf_80807e9626ee48fa:
		__obf_bbb97c28ddd95a60 = &__obf_db79b5bc73d5a614{__obf_1423e711ab2d7104: 1, __obf_a9a6d1fd09e8597b: 9, __obf_8213f88df91d8826: bitset.New(__obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_81568e00075d439f), __obf_94407096894974c2: bitset.New(__obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_81568e00075d439f, __obf_968c0831b0eb9e22), __obf_32106c8c49ebde5d: bitset.New(__obf_968c0831b0eb9e22, __obf_81568e00075d439f, __obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22), __obf_9a03c960dad69649: 10, __obf_d05ff65d77bfc31e: 9, __obf_2a2cbaecb5ec2f61: 8}
	case __obf_9bb2ab43c09df8ad:
		__obf_bbb97c28ddd95a60 = &__obf_db79b5bc73d5a614{__obf_1423e711ab2d7104: 10, __obf_a9a6d1fd09e8597b: 26, __obf_8213f88df91d8826: bitset.New(__obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_81568e00075d439f), __obf_94407096894974c2: bitset.New(__obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_81568e00075d439f, __obf_968c0831b0eb9e22), __obf_32106c8c49ebde5d: bitset.New(__obf_968c0831b0eb9e22, __obf_81568e00075d439f, __obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22), __obf_9a03c960dad69649: 12, __obf_d05ff65d77bfc31e: 11, __obf_2a2cbaecb5ec2f61: 16}
	case __obf_7f04905a311519fa:
		__obf_bbb97c28ddd95a60 = &__obf_db79b5bc73d5a614{__obf_1423e711ab2d7104: 27, __obf_a9a6d1fd09e8597b: 40, __obf_8213f88df91d8826: bitset.New(__obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_81568e00075d439f), __obf_94407096894974c2: bitset.New(__obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22, __obf_81568e00075d439f, __obf_968c0831b0eb9e22), __obf_32106c8c49ebde5d: bitset.New(__obf_968c0831b0eb9e22, __obf_81568e00075d439f, __obf_968c0831b0eb9e22, __obf_968c0831b0eb9e22), __obf_9a03c960dad69649: 14, __obf_d05ff65d77bfc31e: 13, __obf_2a2cbaecb5ec2f61: 16}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_bbb97c28ddd95a60
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_35d7b1c8ddee5d52(__obf_7a0637e203e4ea7a []byte) (*bitset.Bitset, error) {
	__obf_bbb97c28ddd95a60.__obf_7a0637e203e4ea7a = __obf_7a0637e203e4ea7a
	__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f = nil
	__obf_bbb97c28ddd95a60.__obf_4fa8b1538f6edfcb = nil

	if len(__obf_7a0637e203e4ea7a) == 0 {
		return nil, errors.New("no data to encode")
	}
	__obf_5010ef42afe4c769 := // Classify data into unoptimised segments.
		__obf_bbb97c28ddd95a60.__obf_e8522880c489e1a9()
	__obf_a4be63d440a7a9e0 := // Optimise segments.
		__obf_bbb97c28ddd95a60.__obf_bbed381572c45556()
	if __obf_a4be63d440a7a9e0 != nil {
		return nil, __obf_a4be63d440a7a9e0
	}
	__obf_abd89548eb7561a1 := // Check if a single byte encoded segment would be more efficient.
		0
	for _, __obf_d4fa85a87405a625 := range __obf_bbb97c28ddd95a60.__obf_4fa8b1538f6edfcb {
		__obf_a10120810bad38e5, __obf_a4be63d440a7a9e0 := __obf_bbb97c28ddd95a60.__obf_3ce87097aab2f57d(__obf_d4fa85a87405a625.__obf_51964cb459dfa8c2, len(__obf_d4fa85a87405a625.__obf_7a0637e203e4ea7a))
		if __obf_a4be63d440a7a9e0 != nil {
			return nil, __obf_a4be63d440a7a9e0
		}
		__obf_abd89548eb7561a1 += __obf_a10120810bad38e5
	}
	__obf_8dec13f5cbdad298, __obf_a4be63d440a7a9e0 := __obf_bbb97c28ddd95a60.__obf_3ce87097aab2f57d(__obf_5010ef42afe4c769, len(__obf_bbb97c28ddd95a60.__obf_7a0637e203e4ea7a))
	if __obf_a4be63d440a7a9e0 != nil {
		return nil, __obf_a4be63d440a7a9e0
	}

	if __obf_8dec13f5cbdad298 <= __obf_abd89548eb7561a1 {
		__obf_bbb97c28ddd95a60.__obf_4fa8b1538f6edfcb = []__obf_e076203f11833315{{__obf_51964cb459dfa8c2: __obf_5010ef42afe4c769, __obf_7a0637e203e4ea7a: __obf_bbb97c28ddd95a60.

			// Encode data.
			__obf_7a0637e203e4ea7a}}
	}
	__obf_3efd75da4d4d8559 := bitset.New()
	for _, __obf_d4fa85a87405a625 := range __obf_bbb97c28ddd95a60.__obf_4fa8b1538f6edfcb {
		__obf_bbb97c28ddd95a60.__obf_ffa55e32a47ec895(__obf_d4fa85a87405a625.__obf_7a0637e203e4ea7a, __obf_d4fa85a87405a625.

			// classifyDataModes classifies the raw data into unoptimised segments.
			// e.g. "123ZZ#!#!" =>
			// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
			//
			// Returns the highest data mode needed to encode the data. e.g. for a mixed
			// numeric/alphanumeric input, the highest is alphanumeric.
			//
			// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
			__obf_51964cb459dfa8c2, __obf_3efd75da4d4d8559)
	}

	return __obf_3efd75da4d4d8559, nil
}

func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_e8522880c489e1a9() __obf_51964cb459dfa8c2 {
	var __obf_bb49f3bbe0778fb0 int
	__obf_3f619dc8695c0767 := __obf_7e29f0a1d309f15b
	__obf_5010ef42afe4c769 := __obf_3f619dc8695c0767

	for __obf_015b1aa3060900e0, __obf_1b6e726e11cd9194 := range __obf_bbb97c28ddd95a60.__obf_7a0637e203e4ea7a {
		__obf_e4198e3dd9ee417c := __obf_7e29f0a1d309f15b
		switch {
		case __obf_1b6e726e11cd9194 >= 0x30 && __obf_1b6e726e11cd9194 <= 0x39:
			__obf_e4198e3dd9ee417c = __obf_2d9adca192f6544c
		case __obf_1b6e726e11cd9194 == 0x20 || __obf_1b6e726e11cd9194 == 0x24 || __obf_1b6e726e11cd9194 == 0x25 || __obf_1b6e726e11cd9194 == 0x2a || __obf_1b6e726e11cd9194 == 0x2b || __obf_1b6e726e11cd9194 ==
			0x2d || __obf_1b6e726e11cd9194 == 0x2e || __obf_1b6e726e11cd9194 == 0x2f || __obf_1b6e726e11cd9194 == 0x3a || (__obf_1b6e726e11cd9194 >= 0x41 && __obf_1b6e726e11cd9194 <= 0x5a):
			__obf_e4198e3dd9ee417c = __obf_06a20756bb412f63
		default:
			__obf_e4198e3dd9ee417c = __obf_98f9ed48ce6f730b
		}

		if __obf_e4198e3dd9ee417c != __obf_3f619dc8695c0767 {
			if __obf_015b1aa3060900e0 > 0 {
				__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f = append(__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f, __obf_e076203f11833315{__obf_51964cb459dfa8c2: __obf_3f619dc8695c0767, __obf_7a0637e203e4ea7a: __obf_bbb97c28ddd95a60.__obf_7a0637e203e4ea7a[__obf_bb49f3bbe0778fb0:__obf_015b1aa3060900e0]})
				__obf_bb49f3bbe0778fb0 = __obf_015b1aa3060900e0
			}
			__obf_3f619dc8695c0767 = __obf_e4198e3dd9ee417c
		}

		if __obf_e4198e3dd9ee417c > __obf_5010ef42afe4c769 {
			__obf_5010ef42afe4c769 = __obf_e4198e3dd9ee417c
		}
	}
	__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f = append(__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f, __obf_e076203f11833315{__obf_51964cb459dfa8c2: __obf_3f619dc8695c0767, __obf_7a0637e203e4ea7a: __obf_bbb97c28ddd95a60.__obf_7a0637e203e4ea7a[__obf_bb49f3bbe0778fb0:len(__obf_bbb97c28ddd95a60.__obf_7a0637e203e4ea7a)]})

	return __obf_5010ef42afe4c769
}

// optimiseDataModes optimises the list of segments to reduce the overall output
// encoded data length.
//
// The algorithm coalesces adjacent segments. segments are only coalesced when
// the Data Modes are compatible, and when the coalesced segment has a shorter
// encoded length than separate segments.
//
// Multiple segments may be coalesced. For example a string of alternating
// alphanumeric/numeric segments ANANANANA can be optimised to just A.
func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_bbed381572c45556() error {
	for __obf_015b1aa3060900e0 := 0; __obf_015b1aa3060900e0 < len(__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f); {
		__obf_3f619dc8695c0767 := __obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f[__obf_015b1aa3060900e0].__obf_51964cb459dfa8c2
		__obf_53298a87b69261ec := len(__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f[__obf_015b1aa3060900e0].__obf_7a0637e203e4ea7a)
		__obf_7c0eb8656083f52b := __obf_015b1aa3060900e0 + 1
		for __obf_7c0eb8656083f52b < len(__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f) {
			__obf_4fc4c7588e2baca4 := len(__obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f[__obf_7c0eb8656083f52b].__obf_7a0637e203e4ea7a)
			__obf_29ce119b5376a27c := __obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f[__obf_7c0eb8656083f52b].__obf_51964cb459dfa8c2

			if __obf_29ce119b5376a27c > __obf_3f619dc8695c0767 {
				break
			}
			__obf_abf3c0f4414901f9, __obf_a4be63d440a7a9e0 := __obf_bbb97c28ddd95a60.__obf_3ce87097aab2f57d(__obf_3f619dc8695c0767, __obf_53298a87b69261ec+__obf_4fc4c7588e2baca4)

			if __obf_a4be63d440a7a9e0 != nil {
				return __obf_a4be63d440a7a9e0
			}
			__obf_63b0fbe68d4ffe5d, __obf_a4be63d440a7a9e0 := __obf_bbb97c28ddd95a60.__obf_3ce87097aab2f57d(__obf_3f619dc8695c0767, __obf_53298a87b69261ec)

			if __obf_a4be63d440a7a9e0 != nil {
				return __obf_a4be63d440a7a9e0
			}
			__obf_2cb8ff7fb69fa9b7, __obf_a4be63d440a7a9e0 := __obf_bbb97c28ddd95a60.__obf_3ce87097aab2f57d(__obf_29ce119b5376a27c, __obf_4fc4c7588e2baca4)

			if __obf_a4be63d440a7a9e0 != nil {
				return __obf_a4be63d440a7a9e0
			}

			if __obf_abf3c0f4414901f9 < __obf_63b0fbe68d4ffe5d+__obf_2cb8ff7fb69fa9b7 {
				__obf_7c0eb8656083f52b++
				__obf_53298a87b69261ec += __obf_4fc4c7588e2baca4
			} else {
				break
			}
		}
		__obf_4fa8b1538f6edfcb := __obf_e076203f11833315{__obf_51964cb459dfa8c2: __obf_3f619dc8695c0767, __obf_7a0637e203e4ea7a: make([]byte, 0, __obf_53298a87b69261ec)}

		for __obf_8f752cf8fb327ab3 := __obf_015b1aa3060900e0; __obf_8f752cf8fb327ab3 < __obf_7c0eb8656083f52b; __obf_8f752cf8fb327ab3++ {
			__obf_4fa8b1538f6edfcb.__obf_7a0637e203e4ea7a = append(__obf_4fa8b1538f6edfcb.__obf_7a0637e203e4ea7a, __obf_bbb97c28ddd95a60.__obf_bd9515c26418e69f[__obf_8f752cf8fb327ab3].__obf_7a0637e203e4ea7a...)
		}
		__obf_bbb97c28ddd95a60.__obf_4fa8b1538f6edfcb = append(__obf_bbb97c28ddd95a60.__obf_4fa8b1538f6edfcb, __obf_4fa8b1538f6edfcb)
		__obf_015b1aa3060900e0 = __obf_7c0eb8656083f52b
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_ffa55e32a47ec895(__obf_7a0637e203e4ea7a []byte, __obf_51964cb459dfa8c2 __obf_51964cb459dfa8c2, __obf_3efd75da4d4d8559 *bitset.Bitset) {
	__obf_934de79d2fc3b1d4 := __obf_bbb97c28ddd95a60.__obf_934de79d2fc3b1d4(__obf_51964cb459dfa8c2)
	__obf_622de2a9b942161c := __obf_bbb97c28ddd95a60.__obf_622de2a9b942161c(__obf_51964cb459dfa8c2)
	__obf_3efd75da4d4d8559.

		// Append mode indicator.
		Append(__obf_934de79d2fc3b1d4)
	__obf_3efd75da4d4d8559.

		// Append character count.
		AppendUint32(uint32(len(__obf_7a0637e203e4ea7a)), __obf_622de2a9b942161c)

	// Append data.
	switch __obf_51964cb459dfa8c2 {
	case __obf_2d9adca192f6544c:
		for __obf_015b1aa3060900e0 := 0; __obf_015b1aa3060900e0 < len(__obf_7a0637e203e4ea7a); __obf_015b1aa3060900e0 += 3 {
			__obf_3ed67346b65bc666 := len(__obf_7a0637e203e4ea7a) - __obf_015b1aa3060900e0

			var __obf_a5e034be5f59dea2 uint32
			__obf_21c23b12845c92dc := 1

			for __obf_7c0eb8656083f52b := 0; __obf_7c0eb8656083f52b < __obf_3ed67346b65bc666 && __obf_7c0eb8656083f52b < 3; __obf_7c0eb8656083f52b++ {
				__obf_a5e034be5f59dea2 *= 10
				__obf_a5e034be5f59dea2 += uint32(__obf_7a0637e203e4ea7a[__obf_015b1aa3060900e0+__obf_7c0eb8656083f52b] - 0x30)
				__obf_21c23b12845c92dc += 3
			}
			__obf_3efd75da4d4d8559.
				AppendUint32(__obf_a5e034be5f59dea2, __obf_21c23b12845c92dc)
		}
	case __obf_06a20756bb412f63:
		for __obf_015b1aa3060900e0 := 0; __obf_015b1aa3060900e0 < len(__obf_7a0637e203e4ea7a); __obf_015b1aa3060900e0 += 2 {
			__obf_3ed67346b65bc666 := len(__obf_7a0637e203e4ea7a) - __obf_015b1aa3060900e0

			var __obf_a5e034be5f59dea2 uint32
			for __obf_7c0eb8656083f52b := 0; __obf_7c0eb8656083f52b < __obf_3ed67346b65bc666 && __obf_7c0eb8656083f52b < 2; __obf_7c0eb8656083f52b++ {
				__obf_a5e034be5f59dea2 *= 45
				__obf_a5e034be5f59dea2 += __obf_b613e6a9388f1d17(__obf_7a0637e203e4ea7a[__obf_015b1aa3060900e0+__obf_7c0eb8656083f52b])
			}
			__obf_21c23b12845c92dc := 6
			if __obf_3ed67346b65bc666 > 1 {
				__obf_21c23b12845c92dc = 11
			}
			__obf_3efd75da4d4d8559.
				AppendUint32(__obf_a5e034be5f59dea2, __obf_21c23b12845c92dc)
		}
	case __obf_98f9ed48ce6f730b:
		for _, __obf_22b7af99cfb59bff := range __obf_7a0637e203e4ea7a {
			__obf_3efd75da4d4d8559.
				AppendByte(__obf_22b7af99cfb59bff, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_934de79d2fc3b1d4(__obf_51964cb459dfa8c2 __obf_51964cb459dfa8c2) *bitset.Bitset {
	switch __obf_51964cb459dfa8c2 {
	case __obf_2d9adca192f6544c:
		return __obf_bbb97c28ddd95a60.__obf_8213f88df91d8826
	case __obf_06a20756bb412f63:
		return __obf_bbb97c28ddd95a60.__obf_94407096894974c2
	case __obf_98f9ed48ce6f730b:
		return __obf_bbb97c28ddd95a60.__obf_32106c8c49ebde5d
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_622de2a9b942161c(__obf_51964cb459dfa8c2 __obf_51964cb459dfa8c2) int {
	switch __obf_51964cb459dfa8c2 {
	case __obf_2d9adca192f6544c:
		return __obf_bbb97c28ddd95a60.__obf_9a03c960dad69649
	case __obf_06a20756bb412f63:
		return __obf_bbb97c28ddd95a60.__obf_d05ff65d77bfc31e
	case __obf_98f9ed48ce6f730b:
		return __obf_bbb97c28ddd95a60.__obf_2a2cbaecb5ec2f61
	default:
		log.Panic("Unknown data mode")
	}

	return 0
}

// encodedLength returns the number of bits required to encode n symbols in
// dataMode.
//
// The number of bits required is affected by:
//   - QR code type - Mode Indicator length.
//   - Data mode - number of bits used to represent data length.
//   - Data mode - how the data is encoded.
//   - Number of symbols encoded.
//
// An error is returned if the mode is not supported, or the length requested is
// too long to be represented.
func (__obf_bbb97c28ddd95a60 *__obf_db79b5bc73d5a614) __obf_3ce87097aab2f57d(__obf_51964cb459dfa8c2 __obf_51964cb459dfa8c2, __obf_03aa75734a20845a int) (int, error) {
	__obf_934de79d2fc3b1d4 := __obf_bbb97c28ddd95a60.__obf_934de79d2fc3b1d4(__obf_51964cb459dfa8c2)
	__obf_622de2a9b942161c := __obf_bbb97c28ddd95a60.__obf_622de2a9b942161c(__obf_51964cb459dfa8c2)

	if __obf_934de79d2fc3b1d4 == nil {
		return 0, errors.New("mode not supported")
	}
	__obf_5d039bd3a406ae55 := (1 << uint8(__obf_622de2a9b942161c)) - 1

	if __obf_03aa75734a20845a > __obf_5d039bd3a406ae55 {
		return 0, errors.New("length too long to be represented")
	}
	__obf_a10120810bad38e5 := __obf_934de79d2fc3b1d4.Len() + __obf_622de2a9b942161c

	switch __obf_51964cb459dfa8c2 {
	case __obf_2d9adca192f6544c:
		__obf_a10120810bad38e5 += 10 * (__obf_03aa75734a20845a / 3)

		if __obf_03aa75734a20845a%3 != 0 {
			__obf_a10120810bad38e5 += 1 + 3*(__obf_03aa75734a20845a%3)
		}
	case __obf_06a20756bb412f63:
		__obf_a10120810bad38e5 += 11 * (__obf_03aa75734a20845a / 2)
		__obf_a10120810bad38e5 += 6 * (__obf_03aa75734a20845a % 2)
	case __obf_98f9ed48ce6f730b:
		__obf_a10120810bad38e5 += 8 * __obf_03aa75734a20845a
	}

	return __obf_a10120810bad38e5, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_b613e6a9388f1d17(__obf_1b6e726e11cd9194 byte) uint32 {
	__obf_6f119af4b0794054 := uint32(__obf_1b6e726e11cd9194)

	switch {
	case __obf_6f119af4b0794054 >= '0' && __obf_6f119af4b0794054 <= '9':
		// 0-9 encoded as 0-9.
		return __obf_6f119af4b0794054 - '0'
	case __obf_6f119af4b0794054 >= 'A' && __obf_6f119af4b0794054 <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_6f119af4b0794054 - 'A' + 10
	case __obf_6f119af4b0794054 == ' ':
		return 36
	case __obf_6f119af4b0794054 == '$':
		return 37
	case __obf_6f119af4b0794054 == '%':
		return 38
	case __obf_6f119af4b0794054 == '*':
		return 39
	case __obf_6f119af4b0794054 == '+':
		return 40
	case __obf_6f119af4b0794054 == '-':
		return 41
	case __obf_6f119af4b0794054 == '.':
		return 42
	case __obf_6f119af4b0794054 == '/':
		return 43
	case __obf_6f119af4b0794054 == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_1b6e726e11cd9194)
	}

	return 0
}
