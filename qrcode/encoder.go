// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_d8b53a99900c2ed7

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
type __obf_5c4639a52fd845a9 uint8

const (
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	__obf_8904eeb0e6070283 __obf_5c4639a52fd845a9 = 1 << iota
	__obf_4e400452bceac92f
	__obf_2c1e89b2a5d75af1
	__obf_dc644ee144ffe79d
)

// dataModeString returns d as a short printable string.
func __obf_0337d1690cdd0f4b(__obf_afa8ed5b526f190d __obf_5c4639a52fd845a9) string {
	switch __obf_afa8ed5b526f190d {
	case __obf_8904eeb0e6070283:
		return "none"
	case __obf_4e400452bceac92f:
		return "numeric"
	case __obf_2c1e89b2a5d75af1:
		return "alphanumeric"
	case __obf_dc644ee144ffe79d:
		return "byte"
	}

	return "unknown"
}

type __obf_80ac4aaf2f637970 uint8

const (
	__obf_c531a6f523406e39 __obf_80ac4aaf2f637970 = iota
	__obf_c0bbd372ece2268e
	__obf_14b09e7535ac520c
)

// segment is a single segment of data.
type __obf_625a739863f7df0b struct {
	// Data Mode (e.g. numeric).
	__obf_5c4639a52fd845a9 __obf_5c4639a52fd845a9

	// segment data (e.g. "abc").
	__obf_597c6046cd04cc8f []byte
}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_c5fc77e1fc362b78 struct {
	// Minimum & maximum versions supported.
	__obf_37d39dd366b6222a int
	__obf_f36d7698dc68d8b5 int

	// Mode indicator bit sequences.
	__obf_616a4bf40e113ebb *bitset.Bitset
	__obf_8bed5a222d237455 *bitset.Bitset
	__obf_de42629460bbfb17 *bitset.Bitset

	// Character count lengths.
	__obf_390d5f6e6bc86644 int
	__obf_8677a2566a502741 int
	__obf_daef7b2eb908826a int

	// The raw input data.
	__obf_597c6046cd04cc8f []byte

	// The data classified into unoptimised segments.
	__obf_103a42653732bfb0 []__obf_625a739863f7df0b

	// The data classified into optimised segments.
	__obf_5540ca6ce621a891 []__obf_625a739863f7df0b
}

// newDataEncoder constructs a dataEncoder.
func __obf_0a11f28667b2f0e3(__obf_c40c4d1c8e73185e __obf_80ac4aaf2f637970) *__obf_c5fc77e1fc362b78 {
	__obf_afa8ed5b526f190d := &__obf_c5fc77e1fc362b78{}

	switch __obf_c40c4d1c8e73185e {
	case __obf_c531a6f523406e39:
		__obf_afa8ed5b526f190d = &__obf_c5fc77e1fc362b78{
			__obf_37d39dd366b6222a: 1,
			__obf_f36d7698dc68d8b5: 9,
			__obf_616a4bf40e113ebb: bitset.New(__obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22),
			__obf_8bed5a222d237455: bitset.New(__obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22, __obf_8bb2569d08f1d0d7),
			__obf_de42629460bbfb17: bitset.New(__obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22, __obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7),
			__obf_390d5f6e6bc86644: 10,
			__obf_8677a2566a502741: 9,
			__obf_daef7b2eb908826a: 8,
		}
	case __obf_c0bbd372ece2268e:
		__obf_afa8ed5b526f190d = &__obf_c5fc77e1fc362b78{
			__obf_37d39dd366b6222a: 10,
			__obf_f36d7698dc68d8b5: 26,
			__obf_616a4bf40e113ebb: bitset.New(__obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22),
			__obf_8bed5a222d237455: bitset.New(__obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22, __obf_8bb2569d08f1d0d7),
			__obf_de42629460bbfb17: bitset.New(__obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22, __obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7),
			__obf_390d5f6e6bc86644: 12,
			__obf_8677a2566a502741: 11,
			__obf_daef7b2eb908826a: 16,
		}
	case __obf_14b09e7535ac520c:
		__obf_afa8ed5b526f190d = &__obf_c5fc77e1fc362b78{
			__obf_37d39dd366b6222a: 27,
			__obf_f36d7698dc68d8b5: 40,
			__obf_616a4bf40e113ebb: bitset.New(__obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22),
			__obf_8bed5a222d237455: bitset.New(__obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22, __obf_8bb2569d08f1d0d7),
			__obf_de42629460bbfb17: bitset.New(__obf_8bb2569d08f1d0d7, __obf_dd8a19138b3d4a22, __obf_8bb2569d08f1d0d7, __obf_8bb2569d08f1d0d7),
			__obf_390d5f6e6bc86644: 14,
			__obf_8677a2566a502741: 13,
			__obf_daef7b2eb908826a: 16,
		}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_afa8ed5b526f190d
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_d69ee0563ed868f4(__obf_597c6046cd04cc8f []byte) (*bitset.Bitset, error) {
	__obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f = __obf_597c6046cd04cc8f
	__obf_afa8ed5b526f190d.__obf_103a42653732bfb0 = nil
	__obf_afa8ed5b526f190d.__obf_5540ca6ce621a891 = nil

	if len(__obf_597c6046cd04cc8f) == 0 {
		return nil, errors.New("no data to encode")
	}

	// Classify data into unoptimised segments.
	__obf_1b2182c8ce3b3e61 := __obf_afa8ed5b526f190d.__obf_c88e0f992b68c6a7()

	// Optimise segments.
	__obf_f3cf37a038d12d80 := __obf_afa8ed5b526f190d.__obf_b0d25d23359160c2()
	if __obf_f3cf37a038d12d80 != nil {
		return nil, __obf_f3cf37a038d12d80
	}

	// Check if a single byte encoded segment would be more efficient.
	__obf_22a214620c73a3d1 := 0
	for _, __obf_da19f8f4aed522cb := range __obf_afa8ed5b526f190d.__obf_5540ca6ce621a891 {
		__obf_2b2e0ead9ee6917b, __obf_f3cf37a038d12d80 := __obf_afa8ed5b526f190d.__obf_babc4e2d32fa0aee(__obf_da19f8f4aed522cb.__obf_5c4639a52fd845a9, len(__obf_da19f8f4aed522cb.__obf_597c6046cd04cc8f))
		if __obf_f3cf37a038d12d80 != nil {
			return nil, __obf_f3cf37a038d12d80
		}
		__obf_22a214620c73a3d1 += __obf_2b2e0ead9ee6917b
	}

	__obf_7c89bd2ed253a1bd, __obf_f3cf37a038d12d80 := __obf_afa8ed5b526f190d.__obf_babc4e2d32fa0aee(__obf_1b2182c8ce3b3e61, len(__obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f))
	if __obf_f3cf37a038d12d80 != nil {
		return nil, __obf_f3cf37a038d12d80
	}

	if __obf_7c89bd2ed253a1bd <= __obf_22a214620c73a3d1 {
		__obf_afa8ed5b526f190d.__obf_5540ca6ce621a891 = []__obf_625a739863f7df0b{{__obf_5c4639a52fd845a9: __obf_1b2182c8ce3b3e61, __obf_597c6046cd04cc8f: __obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f}}
	}

	// Encode data.
	__obf_db73de0d3577b11c := bitset.New()
	for _, __obf_da19f8f4aed522cb := range __obf_afa8ed5b526f190d.__obf_5540ca6ce621a891 {
		__obf_afa8ed5b526f190d.__obf_ea8717f83acf7afb(__obf_da19f8f4aed522cb.__obf_597c6046cd04cc8f, __obf_da19f8f4aed522cb.__obf_5c4639a52fd845a9, __obf_db73de0d3577b11c)
	}

	return __obf_db73de0d3577b11c, nil
}

// classifyDataModes classifies the raw data into unoptimised segments.
// e.g. "123ZZ#!#!" =>
// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
//
// Returns the highest data mode needed to encode the data. e.g. for a mixed
// numeric/alphanumeric input, the highest is alphanumeric.
//
// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_c88e0f992b68c6a7() __obf_5c4639a52fd845a9 {
	var __obf_c6edaa3d3a18688d int
	__obf_d8f32e09042a3b32 := __obf_8904eeb0e6070283
	__obf_1b2182c8ce3b3e61 := __obf_d8f32e09042a3b32

	for __obf_6ffda43d9ecd66ed, __obf_4650eaa332d33944 := range __obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f {
		__obf_35a9b6929e13af76 := __obf_8904eeb0e6070283
		switch {
		case __obf_4650eaa332d33944 >= 0x30 && __obf_4650eaa332d33944 <= 0x39:
			__obf_35a9b6929e13af76 = __obf_4e400452bceac92f
		case __obf_4650eaa332d33944 == 0x20 || __obf_4650eaa332d33944 == 0x24 || __obf_4650eaa332d33944 == 0x25 || __obf_4650eaa332d33944 == 0x2a || __obf_4650eaa332d33944 == 0x2b || __obf_4650eaa332d33944 ==
			0x2d || __obf_4650eaa332d33944 == 0x2e || __obf_4650eaa332d33944 == 0x2f || __obf_4650eaa332d33944 == 0x3a || (__obf_4650eaa332d33944 >= 0x41 && __obf_4650eaa332d33944 <= 0x5a):
			__obf_35a9b6929e13af76 = __obf_2c1e89b2a5d75af1
		default:
			__obf_35a9b6929e13af76 = __obf_dc644ee144ffe79d
		}

		if __obf_35a9b6929e13af76 != __obf_d8f32e09042a3b32 {
			if __obf_6ffda43d9ecd66ed > 0 {
				__obf_afa8ed5b526f190d.__obf_103a42653732bfb0 = append(__obf_afa8ed5b526f190d.__obf_103a42653732bfb0, __obf_625a739863f7df0b{__obf_5c4639a52fd845a9: __obf_d8f32e09042a3b32, __obf_597c6046cd04cc8f: __obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f[__obf_c6edaa3d3a18688d:__obf_6ffda43d9ecd66ed]})

				__obf_c6edaa3d3a18688d = __obf_6ffda43d9ecd66ed
			}

			__obf_d8f32e09042a3b32 = __obf_35a9b6929e13af76
		}

		if __obf_35a9b6929e13af76 > __obf_1b2182c8ce3b3e61 {
			__obf_1b2182c8ce3b3e61 = __obf_35a9b6929e13af76
		}
	}

	__obf_afa8ed5b526f190d.__obf_103a42653732bfb0 = append(__obf_afa8ed5b526f190d.__obf_103a42653732bfb0, __obf_625a739863f7df0b{__obf_5c4639a52fd845a9: __obf_d8f32e09042a3b32, __obf_597c6046cd04cc8f: __obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f[__obf_c6edaa3d3a18688d:len(__obf_afa8ed5b526f190d.__obf_597c6046cd04cc8f)]})

	return __obf_1b2182c8ce3b3e61
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
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_b0d25d23359160c2() error {
	for __obf_6ffda43d9ecd66ed := 0; __obf_6ffda43d9ecd66ed < len(__obf_afa8ed5b526f190d.__obf_103a42653732bfb0); {
		__obf_d8f32e09042a3b32 := __obf_afa8ed5b526f190d.__obf_103a42653732bfb0[__obf_6ffda43d9ecd66ed].__obf_5c4639a52fd845a9
		__obf_26f8637738bf742d := len(__obf_afa8ed5b526f190d.__obf_103a42653732bfb0[__obf_6ffda43d9ecd66ed].__obf_597c6046cd04cc8f)

		__obf_b09a06cc5f53e0ed := __obf_6ffda43d9ecd66ed + 1
		for __obf_b09a06cc5f53e0ed < len(__obf_afa8ed5b526f190d.__obf_103a42653732bfb0) {
			__obf_f26beb3ccfdbc9d8 := len(__obf_afa8ed5b526f190d.__obf_103a42653732bfb0[__obf_b09a06cc5f53e0ed].__obf_597c6046cd04cc8f)
			__obf_37e2afe8f7d6ea97 := __obf_afa8ed5b526f190d.__obf_103a42653732bfb0[__obf_b09a06cc5f53e0ed].__obf_5c4639a52fd845a9

			if __obf_37e2afe8f7d6ea97 > __obf_d8f32e09042a3b32 {
				break
			}

			__obf_7619804ba9ea12f1, __obf_f3cf37a038d12d80 := __obf_afa8ed5b526f190d.__obf_babc4e2d32fa0aee(__obf_d8f32e09042a3b32, __obf_26f8637738bf742d+__obf_f26beb3ccfdbc9d8)

			if __obf_f3cf37a038d12d80 != nil {
				return __obf_f3cf37a038d12d80
			}

			__obf_fa4ca55cc5cde682, __obf_f3cf37a038d12d80 := __obf_afa8ed5b526f190d.__obf_babc4e2d32fa0aee(__obf_d8f32e09042a3b32, __obf_26f8637738bf742d)

			if __obf_f3cf37a038d12d80 != nil {
				return __obf_f3cf37a038d12d80
			}

			__obf_f89d40f635072397, __obf_f3cf37a038d12d80 := __obf_afa8ed5b526f190d.__obf_babc4e2d32fa0aee(__obf_37e2afe8f7d6ea97, __obf_f26beb3ccfdbc9d8)

			if __obf_f3cf37a038d12d80 != nil {
				return __obf_f3cf37a038d12d80
			}

			if __obf_7619804ba9ea12f1 < __obf_fa4ca55cc5cde682+__obf_f89d40f635072397 {
				__obf_b09a06cc5f53e0ed++
				__obf_26f8637738bf742d += __obf_f26beb3ccfdbc9d8
			} else {
				break
			}
		}

		__obf_5540ca6ce621a891 := __obf_625a739863f7df0b{__obf_5c4639a52fd845a9: __obf_d8f32e09042a3b32,
			__obf_597c6046cd04cc8f: make([]byte, 0, __obf_26f8637738bf742d)}

		for __obf_aec22229b859db6e := __obf_6ffda43d9ecd66ed; __obf_aec22229b859db6e < __obf_b09a06cc5f53e0ed; __obf_aec22229b859db6e++ {
			__obf_5540ca6ce621a891.__obf_597c6046cd04cc8f = append(__obf_5540ca6ce621a891.__obf_597c6046cd04cc8f, __obf_afa8ed5b526f190d.__obf_103a42653732bfb0[__obf_aec22229b859db6e].__obf_597c6046cd04cc8f...)
		}

		__obf_afa8ed5b526f190d.__obf_5540ca6ce621a891 = append(__obf_afa8ed5b526f190d.__obf_5540ca6ce621a891, __obf_5540ca6ce621a891)

		__obf_6ffda43d9ecd66ed = __obf_b09a06cc5f53e0ed
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_ea8717f83acf7afb(__obf_597c6046cd04cc8f []byte, __obf_5c4639a52fd845a9 __obf_5c4639a52fd845a9, __obf_db73de0d3577b11c *bitset.Bitset) {
	__obf_089136fa4773c997 := __obf_afa8ed5b526f190d.__obf_089136fa4773c997(__obf_5c4639a52fd845a9)
	__obf_06b32109ee6b3d5f := __obf_afa8ed5b526f190d.__obf_06b32109ee6b3d5f(__obf_5c4639a52fd845a9)

	// Append mode indicator.
	__obf_db73de0d3577b11c.Append(__obf_089136fa4773c997)

	// Append character count.
	__obf_db73de0d3577b11c.AppendUint32(uint32(len(__obf_597c6046cd04cc8f)), __obf_06b32109ee6b3d5f)

	// Append data.
	switch __obf_5c4639a52fd845a9 {
	case __obf_4e400452bceac92f:
		for __obf_6ffda43d9ecd66ed := 0; __obf_6ffda43d9ecd66ed < len(__obf_597c6046cd04cc8f); __obf_6ffda43d9ecd66ed += 3 {
			__obf_99809ffb5b750c9b := len(__obf_597c6046cd04cc8f) - __obf_6ffda43d9ecd66ed

			var __obf_895b3fc3ae24b62b uint32
			__obf_234eec744708ef21 := 1

			for __obf_b09a06cc5f53e0ed := 0; __obf_b09a06cc5f53e0ed < __obf_99809ffb5b750c9b && __obf_b09a06cc5f53e0ed < 3; __obf_b09a06cc5f53e0ed++ {
				__obf_895b3fc3ae24b62b *= 10
				__obf_895b3fc3ae24b62b += uint32(__obf_597c6046cd04cc8f[__obf_6ffda43d9ecd66ed+__obf_b09a06cc5f53e0ed] - 0x30)
				__obf_234eec744708ef21 += 3
			}
			__obf_db73de0d3577b11c.AppendUint32(__obf_895b3fc3ae24b62b, __obf_234eec744708ef21)
		}
	case __obf_2c1e89b2a5d75af1:
		for __obf_6ffda43d9ecd66ed := 0; __obf_6ffda43d9ecd66ed < len(__obf_597c6046cd04cc8f); __obf_6ffda43d9ecd66ed += 2 {
			__obf_99809ffb5b750c9b := len(__obf_597c6046cd04cc8f) - __obf_6ffda43d9ecd66ed

			var __obf_895b3fc3ae24b62b uint32
			for __obf_b09a06cc5f53e0ed := 0; __obf_b09a06cc5f53e0ed < __obf_99809ffb5b750c9b && __obf_b09a06cc5f53e0ed < 2; __obf_b09a06cc5f53e0ed++ {
				__obf_895b3fc3ae24b62b *= 45
				__obf_895b3fc3ae24b62b += __obf_09999fc6072e066a(__obf_597c6046cd04cc8f[__obf_6ffda43d9ecd66ed+__obf_b09a06cc5f53e0ed])
			}

			__obf_234eec744708ef21 := 6
			if __obf_99809ffb5b750c9b > 1 {
				__obf_234eec744708ef21 = 11
			}

			__obf_db73de0d3577b11c.AppendUint32(__obf_895b3fc3ae24b62b, __obf_234eec744708ef21)
		}
	case __obf_dc644ee144ffe79d:
		for _, __obf_86534ca54b6c1d7e := range __obf_597c6046cd04cc8f {
			__obf_db73de0d3577b11c.AppendByte(__obf_86534ca54b6c1d7e, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_089136fa4773c997(__obf_5c4639a52fd845a9 __obf_5c4639a52fd845a9) *bitset.Bitset {
	switch __obf_5c4639a52fd845a9 {
	case __obf_4e400452bceac92f:
		return __obf_afa8ed5b526f190d.__obf_616a4bf40e113ebb
	case __obf_2c1e89b2a5d75af1:
		return __obf_afa8ed5b526f190d.__obf_8bed5a222d237455
	case __obf_dc644ee144ffe79d:
		return __obf_afa8ed5b526f190d.__obf_de42629460bbfb17
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_06b32109ee6b3d5f(__obf_5c4639a52fd845a9 __obf_5c4639a52fd845a9) int {
	switch __obf_5c4639a52fd845a9 {
	case __obf_4e400452bceac92f:
		return __obf_afa8ed5b526f190d.__obf_390d5f6e6bc86644
	case __obf_2c1e89b2a5d75af1:
		return __obf_afa8ed5b526f190d.__obf_8677a2566a502741
	case __obf_dc644ee144ffe79d:
		return __obf_afa8ed5b526f190d.__obf_daef7b2eb908826a
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
func (__obf_afa8ed5b526f190d *__obf_c5fc77e1fc362b78) __obf_babc4e2d32fa0aee(__obf_5c4639a52fd845a9 __obf_5c4639a52fd845a9, __obf_76e5f47d6201804c int) (int, error) {
	__obf_089136fa4773c997 := __obf_afa8ed5b526f190d.__obf_089136fa4773c997(__obf_5c4639a52fd845a9)
	__obf_06b32109ee6b3d5f := __obf_afa8ed5b526f190d.__obf_06b32109ee6b3d5f(__obf_5c4639a52fd845a9)

	if __obf_089136fa4773c997 == nil {
		return 0, errors.New("mode not supported")
	}

	__obf_399a13dfdb49f744 := (1 << uint8(__obf_06b32109ee6b3d5f)) - 1

	if __obf_76e5f47d6201804c > __obf_399a13dfdb49f744 {
		return 0, errors.New("length too long to be represented")
	}

	__obf_2b2e0ead9ee6917b := __obf_089136fa4773c997.Len() + __obf_06b32109ee6b3d5f

	switch __obf_5c4639a52fd845a9 {
	case __obf_4e400452bceac92f:
		__obf_2b2e0ead9ee6917b += 10 * (__obf_76e5f47d6201804c / 3)

		if __obf_76e5f47d6201804c%3 != 0 {
			__obf_2b2e0ead9ee6917b += 1 + 3*(__obf_76e5f47d6201804c%3)
		}
	case __obf_2c1e89b2a5d75af1:
		__obf_2b2e0ead9ee6917b += 11 * (__obf_76e5f47d6201804c / 2)
		__obf_2b2e0ead9ee6917b += 6 * (__obf_76e5f47d6201804c % 2)
	case __obf_dc644ee144ffe79d:
		__obf_2b2e0ead9ee6917b += 8 * __obf_76e5f47d6201804c
	}

	return __obf_2b2e0ead9ee6917b, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_09999fc6072e066a(__obf_4650eaa332d33944 byte) uint32 {
	__obf_4b78f541c0594a32 := uint32(__obf_4650eaa332d33944)

	switch {
	case __obf_4b78f541c0594a32 >= '0' && __obf_4b78f541c0594a32 <= '9':
		// 0-9 encoded as 0-9.
		return __obf_4b78f541c0594a32 - '0'
	case __obf_4b78f541c0594a32 >= 'A' && __obf_4b78f541c0594a32 <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_4b78f541c0594a32 - 'A' + 10
	case __obf_4b78f541c0594a32 == ' ':
		return 36
	case __obf_4b78f541c0594a32 == '$':
		return 37
	case __obf_4b78f541c0594a32 == '%':
		return 38
	case __obf_4b78f541c0594a32 == '*':
		return 39
	case __obf_4b78f541c0594a32 == '+':
		return 40
	case __obf_4b78f541c0594a32 == '-':
		return 41
	case __obf_4b78f541c0594a32 == '.':
		return 42
	case __obf_4b78f541c0594a32 == '/':
		return 43
	case __obf_4b78f541c0594a32 == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_4650eaa332d33944)
	}

	return 0
}
