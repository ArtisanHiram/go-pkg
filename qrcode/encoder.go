// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_380fc15a74e6e942

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
type __obf_5567dee69fba9724 uint8

const (
	__obf_a160bc77783ba3bc __obf_5567dee69fba9724 =
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	1 << iota
	__obf_ddbca48ee272f278
	__obf_52ec52d5dc95fc7e
	__obf_cbcd39bb47188d34
)

// dataModeString returns d as a short printable string.
func __obf_a29591f7fbfbf25b(__obf_f3d954fc8d7b3447 __obf_5567dee69fba9724) string {
	switch __obf_f3d954fc8d7b3447 {
	case __obf_a160bc77783ba3bc:
		return "none"
	case __obf_ddbca48ee272f278:
		return "numeric"
	case __obf_52ec52d5dc95fc7e:
		return "alphanumeric"
	case __obf_cbcd39bb47188d34:
		return "byte"
	}

	return "unknown"
}

type __obf_d05a9d69f4a51132 uint8

const (
	__obf_f0e68bac73c2a21c __obf_d05a9d69f4a51132 = iota
	__obf_b270254bf247f115
	__obf_602ce8ec28abacf4
)

// segment is a single segment of data.
type __obf_aec21873922e652f struct {
	__obf_5567dee69fba9724 __obf_5567dee69fba9724
	// Data Mode (e.g. numeric).
	__obf_75065e802bd2136f []byte// segment data (e.g. "abc").

}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_8a2e265fe487f1a2 struct {
	__obf_6ef648b259fb8ea6 int // Minimum & maximum versions supported.
	__obf_ff1062ba2e7ad054 int
	__obf_9bc0a4086d0acadd * // Mode indicator bit sequences.
	bitset.Bitset
	__obf_54260837ef0a396c *bitset.Bitset
	__obf_09570bce043b5e24 *bitset.Bitset
	__obf_3e47d47ae7f25fb3 int // Character count lengths.
	__obf_6315db64a8d841ac int
	__obf_4e68de5bab85a12c int
	__obf_75065e802bd2136f []byte// The raw input data.

	__obf_d0237257a0034029 []__obf_aec21873922e652f// The data classified into unoptimised segments.

	// The data classified into optimised segments.
	__obf_45b59fdd01de4170 []__obf_aec21873922e652f
}

// newDataEncoder constructs a dataEncoder.
func __obf_6bdf8930191977ce(__obf_bfc4cdf552630747 __obf_d05a9d69f4a51132) *__obf_8a2e265fe487f1a2 {
	__obf_f3d954fc8d7b3447 := &__obf_8a2e265fe487f1a2{}

	switch __obf_bfc4cdf552630747 {
	case __obf_f0e68bac73c2a21c:
		__obf_f3d954fc8d7b3447 = &__obf_8a2e265fe487f1a2{__obf_6ef648b259fb8ea6: 1, __obf_ff1062ba2e7ad054: 9, __obf_9bc0a4086d0acadd: bitset.New(__obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_78f6621b614504aa), __obf_54260837ef0a396c: bitset.New(__obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_78f6621b614504aa, __obf_08b3663024fb7581), __obf_09570bce043b5e24: bitset.New(__obf_08b3663024fb7581, __obf_78f6621b614504aa, __obf_08b3663024fb7581, __obf_08b3663024fb7581), __obf_3e47d47ae7f25fb3: 10, __obf_6315db64a8d841ac: 9, __obf_4e68de5bab85a12c: 8}
	case __obf_b270254bf247f115:
		__obf_f3d954fc8d7b3447 = &__obf_8a2e265fe487f1a2{__obf_6ef648b259fb8ea6: 10, __obf_ff1062ba2e7ad054: 26, __obf_9bc0a4086d0acadd: bitset.New(__obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_78f6621b614504aa), __obf_54260837ef0a396c: bitset.New(__obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_78f6621b614504aa, __obf_08b3663024fb7581), __obf_09570bce043b5e24: bitset.New(__obf_08b3663024fb7581, __obf_78f6621b614504aa, __obf_08b3663024fb7581, __obf_08b3663024fb7581), __obf_3e47d47ae7f25fb3: 12, __obf_6315db64a8d841ac: 11, __obf_4e68de5bab85a12c: 16}
	case __obf_602ce8ec28abacf4:
		__obf_f3d954fc8d7b3447 = &__obf_8a2e265fe487f1a2{__obf_6ef648b259fb8ea6: 27, __obf_ff1062ba2e7ad054: 40, __obf_9bc0a4086d0acadd: bitset.New(__obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_78f6621b614504aa), __obf_54260837ef0a396c: bitset.New(__obf_08b3663024fb7581, __obf_08b3663024fb7581, __obf_78f6621b614504aa, __obf_08b3663024fb7581), __obf_09570bce043b5e24: bitset.New(__obf_08b3663024fb7581, __obf_78f6621b614504aa, __obf_08b3663024fb7581, __obf_08b3663024fb7581), __obf_3e47d47ae7f25fb3: 14, __obf_6315db64a8d841ac: 13, __obf_4e68de5bab85a12c: 16}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_f3d954fc8d7b3447
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_d827976a41e20008(__obf_75065e802bd2136f []byte) (*bitset.Bitset, error) {
	__obf_f3d954fc8d7b3447.__obf_75065e802bd2136f = __obf_75065e802bd2136f
	__obf_f3d954fc8d7b3447.__obf_d0237257a0034029 = nil
	__obf_f3d954fc8d7b3447.__obf_45b59fdd01de4170 = nil

	if len(__obf_75065e802bd2136f) == 0 {
		return nil, errors.New("no data to encode")
	}
	__obf_2c1246d5f3a0b3c9 := // Classify data into unoptimised segments.
		__obf_f3d954fc8d7b3447.__obf_905d9afdf3aa7573()
	__obf_643ee744984d6f9b := // Optimise segments.
		__obf_f3d954fc8d7b3447.__obf_1b2f155b412bb830()
	if __obf_643ee744984d6f9b != nil {
		return nil, __obf_643ee744984d6f9b
	}
	__obf_c5d6301b8fdb221a := // Check if a single byte encoded segment would be more efficient.
		0
	for _, __obf_0f0d15721c6081af := range __obf_f3d954fc8d7b3447.__obf_45b59fdd01de4170 {
		__obf_bbeaced94428e75b, __obf_643ee744984d6f9b := __obf_f3d954fc8d7b3447.__obf_1b44ed8776456eba(__obf_0f0d15721c6081af.__obf_5567dee69fba9724, len(__obf_0f0d15721c6081af.__obf_75065e802bd2136f))
		if __obf_643ee744984d6f9b != nil {
			return nil, __obf_643ee744984d6f9b
		}
		__obf_c5d6301b8fdb221a += __obf_bbeaced94428e75b
	}
	__obf_f85c742b2b3ac98e, __obf_643ee744984d6f9b := __obf_f3d954fc8d7b3447.__obf_1b44ed8776456eba(__obf_2c1246d5f3a0b3c9, len(__obf_f3d954fc8d7b3447.__obf_75065e802bd2136f))
	if __obf_643ee744984d6f9b != nil {
		return nil, __obf_643ee744984d6f9b
	}

	if __obf_f85c742b2b3ac98e <= __obf_c5d6301b8fdb221a {
		__obf_f3d954fc8d7b3447.__obf_45b59fdd01de4170 = []__obf_aec21873922e652f{{__obf_5567dee69fba9724: __obf_2c1246d5f3a0b3c9, __obf_75065e802bd2136f: __obf_f3d954fc8d7b3447.

			// Encode data.
			__obf_75065e802bd2136f}}
	}
	__obf_fe987f8387318f16 := bitset.New()
	for _, __obf_0f0d15721c6081af := range __obf_f3d954fc8d7b3447.__obf_45b59fdd01de4170 {
		__obf_f3d954fc8d7b3447.__obf_a094def1f70532bd(__obf_0f0d15721c6081af.__obf_75065e802bd2136f, __obf_0f0d15721c6081af.

			// classifyDataModes classifies the raw data into unoptimised segments.
			// e.g. "123ZZ#!#!" =>
			// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
			//
			// Returns the highest data mode needed to encode the data. e.g. for a mixed
			// numeric/alphanumeric input, the highest is alphanumeric.
			//
			// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
			__obf_5567dee69fba9724, __obf_fe987f8387318f16)
	}

	return __obf_fe987f8387318f16, nil
}

func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_905d9afdf3aa7573() __obf_5567dee69fba9724 {
	var __obf_0fb8459f0ac45df2 int
	__obf_262c205cd6144d5e := __obf_a160bc77783ba3bc
	__obf_2c1246d5f3a0b3c9 := __obf_262c205cd6144d5e

	for __obf_f1022e2071655fce, __obf_3591b555f7f18238 := range __obf_f3d954fc8d7b3447.__obf_75065e802bd2136f {
		__obf_accb2ef0021152c3 := __obf_a160bc77783ba3bc
		switch {
		case __obf_3591b555f7f18238 >= 0x30 && __obf_3591b555f7f18238 <= 0x39:
			__obf_accb2ef0021152c3 = __obf_ddbca48ee272f278
		case __obf_3591b555f7f18238 == 0x20 || __obf_3591b555f7f18238 == 0x24 || __obf_3591b555f7f18238 == 0x25 || __obf_3591b555f7f18238 == 0x2a || __obf_3591b555f7f18238 == 0x2b || __obf_3591b555f7f18238 ==
			0x2d || __obf_3591b555f7f18238 == 0x2e || __obf_3591b555f7f18238 == 0x2f || __obf_3591b555f7f18238 == 0x3a || (__obf_3591b555f7f18238 >= 0x41 && __obf_3591b555f7f18238 <= 0x5a):
			__obf_accb2ef0021152c3 = __obf_52ec52d5dc95fc7e
		default:
			__obf_accb2ef0021152c3 = __obf_cbcd39bb47188d34
		}

		if __obf_accb2ef0021152c3 != __obf_262c205cd6144d5e {
			if __obf_f1022e2071655fce > 0 {
				__obf_f3d954fc8d7b3447.__obf_d0237257a0034029 = append(__obf_f3d954fc8d7b3447.__obf_d0237257a0034029, __obf_aec21873922e652f{__obf_5567dee69fba9724: __obf_262c205cd6144d5e, __obf_75065e802bd2136f: __obf_f3d954fc8d7b3447.__obf_75065e802bd2136f[__obf_0fb8459f0ac45df2:__obf_f1022e2071655fce]})
				__obf_0fb8459f0ac45df2 = __obf_f1022e2071655fce
			}
			__obf_262c205cd6144d5e = __obf_accb2ef0021152c3
		}

		if __obf_accb2ef0021152c3 > __obf_2c1246d5f3a0b3c9 {
			__obf_2c1246d5f3a0b3c9 = __obf_accb2ef0021152c3
		}
	}
	__obf_f3d954fc8d7b3447.__obf_d0237257a0034029 = append(__obf_f3d954fc8d7b3447.__obf_d0237257a0034029, __obf_aec21873922e652f{__obf_5567dee69fba9724: __obf_262c205cd6144d5e, __obf_75065e802bd2136f: __obf_f3d954fc8d7b3447.__obf_75065e802bd2136f[__obf_0fb8459f0ac45df2:len(__obf_f3d954fc8d7b3447.__obf_75065e802bd2136f)]})

	return __obf_2c1246d5f3a0b3c9
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
func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_1b2f155b412bb830() error {
	for __obf_f1022e2071655fce := 0; __obf_f1022e2071655fce < len(__obf_f3d954fc8d7b3447.__obf_d0237257a0034029); {
		__obf_262c205cd6144d5e := __obf_f3d954fc8d7b3447.__obf_d0237257a0034029[__obf_f1022e2071655fce].__obf_5567dee69fba9724
		__obf_44c72372b4ff4f7a := len(__obf_f3d954fc8d7b3447.__obf_d0237257a0034029[__obf_f1022e2071655fce].__obf_75065e802bd2136f)
		__obf_e8fc5f84cbc468cf := __obf_f1022e2071655fce + 1
		for __obf_e8fc5f84cbc468cf < len(__obf_f3d954fc8d7b3447.__obf_d0237257a0034029) {
			__obf_2fa6e31ca4f07a8d := len(__obf_f3d954fc8d7b3447.__obf_d0237257a0034029[__obf_e8fc5f84cbc468cf].__obf_75065e802bd2136f)
			__obf_6b91ccfb3ba62ea3 := __obf_f3d954fc8d7b3447.__obf_d0237257a0034029[__obf_e8fc5f84cbc468cf].__obf_5567dee69fba9724

			if __obf_6b91ccfb3ba62ea3 > __obf_262c205cd6144d5e {
				break
			}
			__obf_b1102d5cfc737da2, __obf_643ee744984d6f9b := __obf_f3d954fc8d7b3447.__obf_1b44ed8776456eba(__obf_262c205cd6144d5e, __obf_44c72372b4ff4f7a+__obf_2fa6e31ca4f07a8d)

			if __obf_643ee744984d6f9b != nil {
				return __obf_643ee744984d6f9b
			}
			__obf_aeb383ff50609c13, __obf_643ee744984d6f9b := __obf_f3d954fc8d7b3447.__obf_1b44ed8776456eba(__obf_262c205cd6144d5e, __obf_44c72372b4ff4f7a)

			if __obf_643ee744984d6f9b != nil {
				return __obf_643ee744984d6f9b
			}
			__obf_c17b90f4bb6fb7dd, __obf_643ee744984d6f9b := __obf_f3d954fc8d7b3447.__obf_1b44ed8776456eba(__obf_6b91ccfb3ba62ea3, __obf_2fa6e31ca4f07a8d)

			if __obf_643ee744984d6f9b != nil {
				return __obf_643ee744984d6f9b
			}

			if __obf_b1102d5cfc737da2 < __obf_aeb383ff50609c13+__obf_c17b90f4bb6fb7dd {
				__obf_e8fc5f84cbc468cf++
				__obf_44c72372b4ff4f7a += __obf_2fa6e31ca4f07a8d
			} else {
				break
			}
		}
		__obf_45b59fdd01de4170 := __obf_aec21873922e652f{__obf_5567dee69fba9724: __obf_262c205cd6144d5e, __obf_75065e802bd2136f: make([]byte, 0, __obf_44c72372b4ff4f7a)}

		for __obf_79e80b6ba607da66 := __obf_f1022e2071655fce; __obf_79e80b6ba607da66 < __obf_e8fc5f84cbc468cf; __obf_79e80b6ba607da66++ {
			__obf_45b59fdd01de4170.__obf_75065e802bd2136f = append(__obf_45b59fdd01de4170.__obf_75065e802bd2136f, __obf_f3d954fc8d7b3447.__obf_d0237257a0034029[__obf_79e80b6ba607da66].__obf_75065e802bd2136f...)
		}
		__obf_f3d954fc8d7b3447.__obf_45b59fdd01de4170 = append(__obf_f3d954fc8d7b3447.__obf_45b59fdd01de4170, __obf_45b59fdd01de4170)
		__obf_f1022e2071655fce = __obf_e8fc5f84cbc468cf
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_a094def1f70532bd(__obf_75065e802bd2136f []byte, __obf_5567dee69fba9724 __obf_5567dee69fba9724, __obf_fe987f8387318f16 *bitset.Bitset) {
	__obf_69a4a38df6721154 := __obf_f3d954fc8d7b3447.__obf_69a4a38df6721154(__obf_5567dee69fba9724)
	__obf_4d05995fcf7d6401 := __obf_f3d954fc8d7b3447.__obf_4d05995fcf7d6401(__obf_5567dee69fba9724)
	__obf_fe987f8387318f16.

		// Append mode indicator.
		Append(__obf_69a4a38df6721154)
	__obf_fe987f8387318f16.

		// Append character count.
		AppendUint32(uint32(len(__obf_75065e802bd2136f)), __obf_4d05995fcf7d6401)

	// Append data.
	switch __obf_5567dee69fba9724 {
	case __obf_ddbca48ee272f278:
		for __obf_f1022e2071655fce := 0; __obf_f1022e2071655fce < len(__obf_75065e802bd2136f); __obf_f1022e2071655fce += 3 {
			__obf_614eca1605792089 := len(__obf_75065e802bd2136f) - __obf_f1022e2071655fce

			var __obf_7c3cf7d3ea763ee4 uint32
			__obf_6960ee2b967ddae1 := 1

			for __obf_e8fc5f84cbc468cf := 0; __obf_e8fc5f84cbc468cf < __obf_614eca1605792089 && __obf_e8fc5f84cbc468cf < 3; __obf_e8fc5f84cbc468cf++ {
				__obf_7c3cf7d3ea763ee4 *= 10
				__obf_7c3cf7d3ea763ee4 += uint32(__obf_75065e802bd2136f[__obf_f1022e2071655fce+__obf_e8fc5f84cbc468cf] - 0x30)
				__obf_6960ee2b967ddae1 += 3
			}
			__obf_fe987f8387318f16.
				AppendUint32(__obf_7c3cf7d3ea763ee4, __obf_6960ee2b967ddae1)
		}
	case __obf_52ec52d5dc95fc7e:
		for __obf_f1022e2071655fce := 0; __obf_f1022e2071655fce < len(__obf_75065e802bd2136f); __obf_f1022e2071655fce += 2 {
			__obf_614eca1605792089 := len(__obf_75065e802bd2136f) - __obf_f1022e2071655fce

			var __obf_7c3cf7d3ea763ee4 uint32
			for __obf_e8fc5f84cbc468cf := 0; __obf_e8fc5f84cbc468cf < __obf_614eca1605792089 && __obf_e8fc5f84cbc468cf < 2; __obf_e8fc5f84cbc468cf++ {
				__obf_7c3cf7d3ea763ee4 *= 45
				__obf_7c3cf7d3ea763ee4 += __obf_14b6e3ce23b6b4c7(__obf_75065e802bd2136f[__obf_f1022e2071655fce+__obf_e8fc5f84cbc468cf])
			}
			__obf_6960ee2b967ddae1 := 6
			if __obf_614eca1605792089 > 1 {
				__obf_6960ee2b967ddae1 = 11
			}
			__obf_fe987f8387318f16.
				AppendUint32(__obf_7c3cf7d3ea763ee4, __obf_6960ee2b967ddae1)
		}
	case __obf_cbcd39bb47188d34:
		for _, __obf_b58cd35067d92ce7 := range __obf_75065e802bd2136f {
			__obf_fe987f8387318f16.
				AppendByte(__obf_b58cd35067d92ce7, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_69a4a38df6721154(__obf_5567dee69fba9724 __obf_5567dee69fba9724) *bitset.Bitset {
	switch __obf_5567dee69fba9724 {
	case __obf_ddbca48ee272f278:
		return __obf_f3d954fc8d7b3447.__obf_9bc0a4086d0acadd
	case __obf_52ec52d5dc95fc7e:
		return __obf_f3d954fc8d7b3447.__obf_54260837ef0a396c
	case __obf_cbcd39bb47188d34:
		return __obf_f3d954fc8d7b3447.__obf_09570bce043b5e24
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_4d05995fcf7d6401(__obf_5567dee69fba9724 __obf_5567dee69fba9724) int {
	switch __obf_5567dee69fba9724 {
	case __obf_ddbca48ee272f278:
		return __obf_f3d954fc8d7b3447.__obf_3e47d47ae7f25fb3
	case __obf_52ec52d5dc95fc7e:
		return __obf_f3d954fc8d7b3447.__obf_6315db64a8d841ac
	case __obf_cbcd39bb47188d34:
		return __obf_f3d954fc8d7b3447.__obf_4e68de5bab85a12c
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
func (__obf_f3d954fc8d7b3447 *__obf_8a2e265fe487f1a2) __obf_1b44ed8776456eba(__obf_5567dee69fba9724 __obf_5567dee69fba9724, __obf_0b01635812afe0bc int) (int, error) {
	__obf_69a4a38df6721154 := __obf_f3d954fc8d7b3447.__obf_69a4a38df6721154(__obf_5567dee69fba9724)
	__obf_4d05995fcf7d6401 := __obf_f3d954fc8d7b3447.__obf_4d05995fcf7d6401(__obf_5567dee69fba9724)

	if __obf_69a4a38df6721154 == nil {
		return 0, errors.New("mode not supported")
	}
	__obf_e3d57159429e9789 := (1 << uint8(__obf_4d05995fcf7d6401)) - 1

	if __obf_0b01635812afe0bc > __obf_e3d57159429e9789 {
		return 0, errors.New("length too long to be represented")
	}
	__obf_bbeaced94428e75b := __obf_69a4a38df6721154.Len() + __obf_4d05995fcf7d6401

	switch __obf_5567dee69fba9724 {
	case __obf_ddbca48ee272f278:
		__obf_bbeaced94428e75b += 10 * (__obf_0b01635812afe0bc / 3)

		if __obf_0b01635812afe0bc%3 != 0 {
			__obf_bbeaced94428e75b += 1 + 3*(__obf_0b01635812afe0bc%3)
		}
	case __obf_52ec52d5dc95fc7e:
		__obf_bbeaced94428e75b += 11 * (__obf_0b01635812afe0bc / 2)
		__obf_bbeaced94428e75b += 6 * (__obf_0b01635812afe0bc % 2)
	case __obf_cbcd39bb47188d34:
		__obf_bbeaced94428e75b += 8 * __obf_0b01635812afe0bc
	}

	return __obf_bbeaced94428e75b, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_14b6e3ce23b6b4c7(__obf_3591b555f7f18238 byte) uint32 {
	__obf_fcb7b90cb2d2fec1 := uint32(__obf_3591b555f7f18238)

	switch {
	case __obf_fcb7b90cb2d2fec1 >= '0' && __obf_fcb7b90cb2d2fec1 <= '9':
		// 0-9 encoded as 0-9.
		return __obf_fcb7b90cb2d2fec1 - '0'
	case __obf_fcb7b90cb2d2fec1 >= 'A' && __obf_fcb7b90cb2d2fec1 <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_fcb7b90cb2d2fec1 - 'A' + 10
	case __obf_fcb7b90cb2d2fec1 == ' ':
		return 36
	case __obf_fcb7b90cb2d2fec1 == '$':
		return 37
	case __obf_fcb7b90cb2d2fec1 == '%':
		return 38
	case __obf_fcb7b90cb2d2fec1 == '*':
		return 39
	case __obf_fcb7b90cb2d2fec1 == '+':
		return 40
	case __obf_fcb7b90cb2d2fec1 == '-':
		return 41
	case __obf_fcb7b90cb2d2fec1 == '.':
		return 42
	case __obf_fcb7b90cb2d2fec1 == '/':
		return 43
	case __obf_fcb7b90cb2d2fec1 == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_3591b555f7f18238)
	}

	return 0
}
