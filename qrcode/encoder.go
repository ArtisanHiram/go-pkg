// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_540b74a9d13704fa

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
type __obf_e4998b76d9af03cb uint8

const (
	__obf_c6c4fd29e9fdeabf __obf_e4998b76d9af03cb =
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	1 << iota
	__obf_cc5dc580e228ac73
	__obf_740f785829efce56
	__obf_5233e2bd38dee79e
)

// dataModeString returns d as a short printable string.
func __obf_e4d7bb8e7a4fda6d(__obf_cf6bdd4e63393758 __obf_e4998b76d9af03cb) string {
	switch __obf_cf6bdd4e63393758 {
	case __obf_c6c4fd29e9fdeabf:
		return "none"
	case __obf_cc5dc580e228ac73:
		return "numeric"
	case __obf_740f785829efce56:
		return "alphanumeric"
	case __obf_5233e2bd38dee79e:
		return "byte"
	}

	return "unknown"
}

type __obf_ede9586c20e29cd6 uint8

const (
	__obf_36e336d7eab6b5c8 __obf_ede9586c20e29cd6 = iota
	__obf_ce4133b003340f88
	__obf_d9027dc4f9acd3b0
)

// segment is a single segment of data.
type __obf_844346755473b968 struct {
	__obf_e4998b76d9af03cb __obf_e4998b76d9af03cb
	// Data Mode (e.g. numeric).
	__obf_3bcdb4f6c983dc7c []byte// segment data (e.g. "abc").

}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_5aac2ee6cac292a2 struct {
	__obf_84c65bc29348e4af int // Minimum & maximum versions supported.
	__obf_69cd1e2223127914 int
	__obf_6d929cb133549473 * // Mode indicator bit sequences.
	bitset.Bitset
	__obf_a888d4e41fcdd672 *bitset.Bitset
	__obf_d0aedc11216c6c97 *bitset.Bitset
	__obf_0c701a3535d5efa0 int // Character count lengths.
	__obf_25086482e80aca1d int
	__obf_d90f2f786faa0425 int
	__obf_3bcdb4f6c983dc7c []byte// The raw input data.

	__obf_6c49d9351542545c []__obf_844346755473b968// The data classified into unoptimised segments.

	// The data classified into optimised segments.
	__obf_20e21a51ad6add71 []__obf_844346755473b968
}

// newDataEncoder constructs a dataEncoder.
func __obf_46cf60d6f5b35f0a(__obf_56c65516aa2ae2b8 __obf_ede9586c20e29cd6) *__obf_5aac2ee6cac292a2 {
	__obf_cf6bdd4e63393758 := &__obf_5aac2ee6cac292a2{}

	switch __obf_56c65516aa2ae2b8 {
	case __obf_36e336d7eab6b5c8:
		__obf_cf6bdd4e63393758 = &__obf_5aac2ee6cac292a2{__obf_84c65bc29348e4af: 1, __obf_69cd1e2223127914: 9, __obf_6d929cb133549473: bitset.New(__obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_55459a2c0723f89c), __obf_a888d4e41fcdd672: bitset.New(__obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_55459a2c0723f89c, __obf_f646277736ff27c0), __obf_d0aedc11216c6c97: bitset.New(__obf_f646277736ff27c0, __obf_55459a2c0723f89c, __obf_f646277736ff27c0, __obf_f646277736ff27c0), __obf_0c701a3535d5efa0: 10, __obf_25086482e80aca1d: 9, __obf_d90f2f786faa0425: 8}
	case __obf_ce4133b003340f88:
		__obf_cf6bdd4e63393758 = &__obf_5aac2ee6cac292a2{__obf_84c65bc29348e4af: 10, __obf_69cd1e2223127914: 26, __obf_6d929cb133549473: bitset.New(__obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_55459a2c0723f89c), __obf_a888d4e41fcdd672: bitset.New(__obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_55459a2c0723f89c, __obf_f646277736ff27c0), __obf_d0aedc11216c6c97: bitset.New(__obf_f646277736ff27c0, __obf_55459a2c0723f89c, __obf_f646277736ff27c0, __obf_f646277736ff27c0), __obf_0c701a3535d5efa0: 12, __obf_25086482e80aca1d: 11, __obf_d90f2f786faa0425: 16}
	case __obf_d9027dc4f9acd3b0:
		__obf_cf6bdd4e63393758 = &__obf_5aac2ee6cac292a2{__obf_84c65bc29348e4af: 27, __obf_69cd1e2223127914: 40, __obf_6d929cb133549473: bitset.New(__obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_55459a2c0723f89c), __obf_a888d4e41fcdd672: bitset.New(__obf_f646277736ff27c0, __obf_f646277736ff27c0, __obf_55459a2c0723f89c, __obf_f646277736ff27c0), __obf_d0aedc11216c6c97: bitset.New(__obf_f646277736ff27c0, __obf_55459a2c0723f89c, __obf_f646277736ff27c0, __obf_f646277736ff27c0), __obf_0c701a3535d5efa0: 14, __obf_25086482e80aca1d: 13, __obf_d90f2f786faa0425: 16}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_cf6bdd4e63393758
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_c7821a6383b7b718(__obf_3bcdb4f6c983dc7c []byte) (*bitset.Bitset, error) {
	__obf_cf6bdd4e63393758.__obf_3bcdb4f6c983dc7c = __obf_3bcdb4f6c983dc7c
	__obf_cf6bdd4e63393758.__obf_6c49d9351542545c = nil
	__obf_cf6bdd4e63393758.__obf_20e21a51ad6add71 = nil

	if len(__obf_3bcdb4f6c983dc7c) == 0 {
		return nil, errors.New("no data to encode")
	}
	__obf_2f6b294b0c3ca81c := // Classify data into unoptimised segments.
		__obf_cf6bdd4e63393758.__obf_d39fadaedea83dc0()
	__obf_800077dfc326e0b7 := // Optimise segments.
		__obf_cf6bdd4e63393758.__obf_4e39bb9528e6a96b()
	if __obf_800077dfc326e0b7 != nil {
		return nil, __obf_800077dfc326e0b7
	}
	__obf_c8400a4881da1b06 := // Check if a single byte encoded segment would be more efficient.
		0
	for _, __obf_6900aee466231dfd := range __obf_cf6bdd4e63393758.__obf_20e21a51ad6add71 {
		__obf_9b6d33e9ef886064, __obf_800077dfc326e0b7 := __obf_cf6bdd4e63393758.__obf_7f4688f820f83b55(__obf_6900aee466231dfd.__obf_e4998b76d9af03cb, len(__obf_6900aee466231dfd.__obf_3bcdb4f6c983dc7c))
		if __obf_800077dfc326e0b7 != nil {
			return nil, __obf_800077dfc326e0b7
		}
		__obf_c8400a4881da1b06 += __obf_9b6d33e9ef886064
	}
	__obf_52b6663ade315ed9, __obf_800077dfc326e0b7 := __obf_cf6bdd4e63393758.__obf_7f4688f820f83b55(__obf_2f6b294b0c3ca81c, len(__obf_cf6bdd4e63393758.__obf_3bcdb4f6c983dc7c))
	if __obf_800077dfc326e0b7 != nil {
		return nil, __obf_800077dfc326e0b7
	}

	if __obf_52b6663ade315ed9 <= __obf_c8400a4881da1b06 {
		__obf_cf6bdd4e63393758.__obf_20e21a51ad6add71 = []__obf_844346755473b968{{__obf_e4998b76d9af03cb: __obf_2f6b294b0c3ca81c, __obf_3bcdb4f6c983dc7c: __obf_cf6bdd4e63393758.

			// Encode data.
			__obf_3bcdb4f6c983dc7c}}
	}
	__obf_cb404255cbbc8689 := bitset.New()
	for _, __obf_6900aee466231dfd := range __obf_cf6bdd4e63393758.__obf_20e21a51ad6add71 {
		__obf_cf6bdd4e63393758.__obf_35a67d4d0550c9d2(__obf_6900aee466231dfd.__obf_3bcdb4f6c983dc7c, __obf_6900aee466231dfd.

			// classifyDataModes classifies the raw data into unoptimised segments.
			// e.g. "123ZZ#!#!" =>
			// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
			//
			// Returns the highest data mode needed to encode the data. e.g. for a mixed
			// numeric/alphanumeric input, the highest is alphanumeric.
			//
			// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
			__obf_e4998b76d9af03cb, __obf_cb404255cbbc8689)
	}

	return __obf_cb404255cbbc8689, nil
}

func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_d39fadaedea83dc0() __obf_e4998b76d9af03cb {
	var __obf_711a399c8c9ecede int
	__obf_0db6273905c0920a := __obf_c6c4fd29e9fdeabf
	__obf_2f6b294b0c3ca81c := __obf_0db6273905c0920a

	for __obf_2e2482997826bcc5, __obf_75eba50799fccf42 := range __obf_cf6bdd4e63393758.__obf_3bcdb4f6c983dc7c {
		__obf_803d0dd04a33b90c := __obf_c6c4fd29e9fdeabf
		switch {
		case __obf_75eba50799fccf42 >= 0x30 && __obf_75eba50799fccf42 <= 0x39:
			__obf_803d0dd04a33b90c = __obf_cc5dc580e228ac73
		case __obf_75eba50799fccf42 == 0x20 || __obf_75eba50799fccf42 == 0x24 || __obf_75eba50799fccf42 == 0x25 || __obf_75eba50799fccf42 == 0x2a || __obf_75eba50799fccf42 == 0x2b || __obf_75eba50799fccf42 ==
			0x2d || __obf_75eba50799fccf42 == 0x2e || __obf_75eba50799fccf42 == 0x2f || __obf_75eba50799fccf42 == 0x3a || (__obf_75eba50799fccf42 >= 0x41 && __obf_75eba50799fccf42 <= 0x5a):
			__obf_803d0dd04a33b90c = __obf_740f785829efce56
		default:
			__obf_803d0dd04a33b90c = __obf_5233e2bd38dee79e
		}

		if __obf_803d0dd04a33b90c != __obf_0db6273905c0920a {
			if __obf_2e2482997826bcc5 > 0 {
				__obf_cf6bdd4e63393758.__obf_6c49d9351542545c = append(__obf_cf6bdd4e63393758.__obf_6c49d9351542545c, __obf_844346755473b968{__obf_e4998b76d9af03cb: __obf_0db6273905c0920a, __obf_3bcdb4f6c983dc7c: __obf_cf6bdd4e63393758.__obf_3bcdb4f6c983dc7c[__obf_711a399c8c9ecede:__obf_2e2482997826bcc5]})
				__obf_711a399c8c9ecede = __obf_2e2482997826bcc5
			}
			__obf_0db6273905c0920a = __obf_803d0dd04a33b90c
		}

		if __obf_803d0dd04a33b90c > __obf_2f6b294b0c3ca81c {
			__obf_2f6b294b0c3ca81c = __obf_803d0dd04a33b90c
		}
	}
	__obf_cf6bdd4e63393758.__obf_6c49d9351542545c = append(__obf_cf6bdd4e63393758.__obf_6c49d9351542545c, __obf_844346755473b968{__obf_e4998b76d9af03cb: __obf_0db6273905c0920a, __obf_3bcdb4f6c983dc7c: __obf_cf6bdd4e63393758.__obf_3bcdb4f6c983dc7c[__obf_711a399c8c9ecede:len(__obf_cf6bdd4e63393758.__obf_3bcdb4f6c983dc7c)]})

	return __obf_2f6b294b0c3ca81c
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
func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_4e39bb9528e6a96b() error {
	for __obf_2e2482997826bcc5 := 0; __obf_2e2482997826bcc5 < len(__obf_cf6bdd4e63393758.__obf_6c49d9351542545c); {
		__obf_0db6273905c0920a := __obf_cf6bdd4e63393758.__obf_6c49d9351542545c[__obf_2e2482997826bcc5].__obf_e4998b76d9af03cb
		__obf_40f308943292c43b := len(__obf_cf6bdd4e63393758.__obf_6c49d9351542545c[__obf_2e2482997826bcc5].__obf_3bcdb4f6c983dc7c)
		__obf_cf60660918b25aa8 := __obf_2e2482997826bcc5 + 1
		for __obf_cf60660918b25aa8 < len(__obf_cf6bdd4e63393758.__obf_6c49d9351542545c) {
			__obf_4574956bc623148f := len(__obf_cf6bdd4e63393758.__obf_6c49d9351542545c[__obf_cf60660918b25aa8].__obf_3bcdb4f6c983dc7c)
			__obf_2c09b3b58355438d := __obf_cf6bdd4e63393758.__obf_6c49d9351542545c[__obf_cf60660918b25aa8].__obf_e4998b76d9af03cb

			if __obf_2c09b3b58355438d > __obf_0db6273905c0920a {
				break
			}
			__obf_90528cb98cb09e20, __obf_800077dfc326e0b7 := __obf_cf6bdd4e63393758.__obf_7f4688f820f83b55(__obf_0db6273905c0920a, __obf_40f308943292c43b+__obf_4574956bc623148f)

			if __obf_800077dfc326e0b7 != nil {
				return __obf_800077dfc326e0b7
			}
			__obf_02838ea9bc883c69, __obf_800077dfc326e0b7 := __obf_cf6bdd4e63393758.__obf_7f4688f820f83b55(__obf_0db6273905c0920a, __obf_40f308943292c43b)

			if __obf_800077dfc326e0b7 != nil {
				return __obf_800077dfc326e0b7
			}
			__obf_b08c2fa107b9a0ea, __obf_800077dfc326e0b7 := __obf_cf6bdd4e63393758.__obf_7f4688f820f83b55(__obf_2c09b3b58355438d, __obf_4574956bc623148f)

			if __obf_800077dfc326e0b7 != nil {
				return __obf_800077dfc326e0b7
			}

			if __obf_90528cb98cb09e20 < __obf_02838ea9bc883c69+__obf_b08c2fa107b9a0ea {
				__obf_cf60660918b25aa8++
				__obf_40f308943292c43b += __obf_4574956bc623148f
			} else {
				break
			}
		}
		__obf_20e21a51ad6add71 := __obf_844346755473b968{__obf_e4998b76d9af03cb: __obf_0db6273905c0920a, __obf_3bcdb4f6c983dc7c: make([]byte, 0, __obf_40f308943292c43b)}

		for __obf_d6474094611100b5 := __obf_2e2482997826bcc5; __obf_d6474094611100b5 < __obf_cf60660918b25aa8; __obf_d6474094611100b5++ {
			__obf_20e21a51ad6add71.__obf_3bcdb4f6c983dc7c = append(__obf_20e21a51ad6add71.__obf_3bcdb4f6c983dc7c, __obf_cf6bdd4e63393758.__obf_6c49d9351542545c[__obf_d6474094611100b5].__obf_3bcdb4f6c983dc7c...)
		}
		__obf_cf6bdd4e63393758.__obf_20e21a51ad6add71 = append(__obf_cf6bdd4e63393758.__obf_20e21a51ad6add71, __obf_20e21a51ad6add71)
		__obf_2e2482997826bcc5 = __obf_cf60660918b25aa8
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_35a67d4d0550c9d2(__obf_3bcdb4f6c983dc7c []byte, __obf_e4998b76d9af03cb __obf_e4998b76d9af03cb, __obf_cb404255cbbc8689 *bitset.Bitset) {
	__obf_75e79779e9efe9f3 := __obf_cf6bdd4e63393758.__obf_75e79779e9efe9f3(__obf_e4998b76d9af03cb)
	__obf_0dd945db5e2c46e9 := __obf_cf6bdd4e63393758.__obf_0dd945db5e2c46e9(__obf_e4998b76d9af03cb)
	__obf_cb404255cbbc8689.

		// Append mode indicator.
		Append(__obf_75e79779e9efe9f3)
	__obf_cb404255cbbc8689.

		// Append character count.
		AppendUint32(uint32(len(__obf_3bcdb4f6c983dc7c)), __obf_0dd945db5e2c46e9)

	// Append data.
	switch __obf_e4998b76d9af03cb {
	case __obf_cc5dc580e228ac73:
		for __obf_2e2482997826bcc5 := 0; __obf_2e2482997826bcc5 < len(__obf_3bcdb4f6c983dc7c); __obf_2e2482997826bcc5 += 3 {
			__obf_6c65ab0013a6f4e2 := len(__obf_3bcdb4f6c983dc7c) - __obf_2e2482997826bcc5

			var __obf_efff9ae9a1dfb06e uint32
			__obf_6d33ebfaab5ccae6 := 1

			for __obf_cf60660918b25aa8 := 0; __obf_cf60660918b25aa8 < __obf_6c65ab0013a6f4e2 && __obf_cf60660918b25aa8 < 3; __obf_cf60660918b25aa8++ {
				__obf_efff9ae9a1dfb06e *= 10
				__obf_efff9ae9a1dfb06e += uint32(__obf_3bcdb4f6c983dc7c[__obf_2e2482997826bcc5+__obf_cf60660918b25aa8] - 0x30)
				__obf_6d33ebfaab5ccae6 += 3
			}
			__obf_cb404255cbbc8689.
				AppendUint32(__obf_efff9ae9a1dfb06e, __obf_6d33ebfaab5ccae6)
		}
	case __obf_740f785829efce56:
		for __obf_2e2482997826bcc5 := 0; __obf_2e2482997826bcc5 < len(__obf_3bcdb4f6c983dc7c); __obf_2e2482997826bcc5 += 2 {
			__obf_6c65ab0013a6f4e2 := len(__obf_3bcdb4f6c983dc7c) - __obf_2e2482997826bcc5

			var __obf_efff9ae9a1dfb06e uint32
			for __obf_cf60660918b25aa8 := 0; __obf_cf60660918b25aa8 < __obf_6c65ab0013a6f4e2 && __obf_cf60660918b25aa8 < 2; __obf_cf60660918b25aa8++ {
				__obf_efff9ae9a1dfb06e *= 45
				__obf_efff9ae9a1dfb06e += __obf_bde054ee3212e0a3(__obf_3bcdb4f6c983dc7c[__obf_2e2482997826bcc5+__obf_cf60660918b25aa8])
			}
			__obf_6d33ebfaab5ccae6 := 6
			if __obf_6c65ab0013a6f4e2 > 1 {
				__obf_6d33ebfaab5ccae6 = 11
			}
			__obf_cb404255cbbc8689.
				AppendUint32(__obf_efff9ae9a1dfb06e, __obf_6d33ebfaab5ccae6)
		}
	case __obf_5233e2bd38dee79e:
		for _, __obf_aa760948f0d0122d := range __obf_3bcdb4f6c983dc7c {
			__obf_cb404255cbbc8689.
				AppendByte(__obf_aa760948f0d0122d, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_75e79779e9efe9f3(__obf_e4998b76d9af03cb __obf_e4998b76d9af03cb) *bitset.Bitset {
	switch __obf_e4998b76d9af03cb {
	case __obf_cc5dc580e228ac73:
		return __obf_cf6bdd4e63393758.__obf_6d929cb133549473
	case __obf_740f785829efce56:
		return __obf_cf6bdd4e63393758.__obf_a888d4e41fcdd672
	case __obf_5233e2bd38dee79e:
		return __obf_cf6bdd4e63393758.__obf_d0aedc11216c6c97
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_0dd945db5e2c46e9(__obf_e4998b76d9af03cb __obf_e4998b76d9af03cb) int {
	switch __obf_e4998b76d9af03cb {
	case __obf_cc5dc580e228ac73:
		return __obf_cf6bdd4e63393758.__obf_0c701a3535d5efa0
	case __obf_740f785829efce56:
		return __obf_cf6bdd4e63393758.__obf_25086482e80aca1d
	case __obf_5233e2bd38dee79e:
		return __obf_cf6bdd4e63393758.__obf_d90f2f786faa0425
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
func (__obf_cf6bdd4e63393758 *__obf_5aac2ee6cac292a2) __obf_7f4688f820f83b55(__obf_e4998b76d9af03cb __obf_e4998b76d9af03cb, __obf_5996f361acddec74 int) (int, error) {
	__obf_75e79779e9efe9f3 := __obf_cf6bdd4e63393758.__obf_75e79779e9efe9f3(__obf_e4998b76d9af03cb)
	__obf_0dd945db5e2c46e9 := __obf_cf6bdd4e63393758.__obf_0dd945db5e2c46e9(__obf_e4998b76d9af03cb)

	if __obf_75e79779e9efe9f3 == nil {
		return 0, errors.New("mode not supported")
	}
	__obf_96964d69d119e097 := (1 << uint8(__obf_0dd945db5e2c46e9)) - 1

	if __obf_5996f361acddec74 > __obf_96964d69d119e097 {
		return 0, errors.New("length too long to be represented")
	}
	__obf_9b6d33e9ef886064 := __obf_75e79779e9efe9f3.Len() + __obf_0dd945db5e2c46e9

	switch __obf_e4998b76d9af03cb {
	case __obf_cc5dc580e228ac73:
		__obf_9b6d33e9ef886064 += 10 * (__obf_5996f361acddec74 / 3)

		if __obf_5996f361acddec74%3 != 0 {
			__obf_9b6d33e9ef886064 += 1 + 3*(__obf_5996f361acddec74%3)
		}
	case __obf_740f785829efce56:
		__obf_9b6d33e9ef886064 += 11 * (__obf_5996f361acddec74 / 2)
		__obf_9b6d33e9ef886064 += 6 * (__obf_5996f361acddec74 % 2)
	case __obf_5233e2bd38dee79e:
		__obf_9b6d33e9ef886064 += 8 * __obf_5996f361acddec74
	}

	return __obf_9b6d33e9ef886064, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_bde054ee3212e0a3(__obf_75eba50799fccf42 byte) uint32 {
	__obf_484efef73eb4b41d := uint32(__obf_75eba50799fccf42)

	switch {
	case __obf_484efef73eb4b41d >= '0' && __obf_484efef73eb4b41d <= '9':
		// 0-9 encoded as 0-9.
		return __obf_484efef73eb4b41d - '0'
	case __obf_484efef73eb4b41d >= 'A' && __obf_484efef73eb4b41d <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_484efef73eb4b41d - 'A' + 10
	case __obf_484efef73eb4b41d == ' ':
		return 36
	case __obf_484efef73eb4b41d == '$':
		return 37
	case __obf_484efef73eb4b41d == '%':
		return 38
	case __obf_484efef73eb4b41d == '*':
		return 39
	case __obf_484efef73eb4b41d == '+':
		return 40
	case __obf_484efef73eb4b41d == '-':
		return 41
	case __obf_484efef73eb4b41d == '.':
		return 42
	case __obf_484efef73eb4b41d == '/':
		return 43
	case __obf_484efef73eb4b41d == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_75eba50799fccf42)
	}

	return 0
}
