// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_07675b7eb1c7284a

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
type __obf_319b72319cecddec uint8

const (
	__obf_c24c42c85933c4c7 __obf_319b72319cecddec =
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	1 << iota
	__obf_917a5ab78c19aa97
	__obf_23eff39094784930
	__obf_105296cdda013cec
)

// dataModeString returns d as a short printable string.
func __obf_605f3c21dcf93517(__obf_f4afc55b009a3568 __obf_319b72319cecddec) string {
	switch __obf_f4afc55b009a3568 {
	case __obf_c24c42c85933c4c7:
		return "none"
	case __obf_917a5ab78c19aa97:
		return "numeric"
	case __obf_23eff39094784930:
		return "alphanumeric"
	case __obf_105296cdda013cec:
		return "byte"
	}

	return "unknown"
}

type __obf_d9ac8a0ee0d0f27f uint8

const (
	__obf_96e9dfef6bbb06ba __obf_d9ac8a0ee0d0f27f = iota
	__obf_bd631aba8d4005cf
	__obf_5da2f2a34edb06bc
)

// segment is a single segment of data.
type __obf_61a731c6456893e0 struct {
	__obf_319b72319cecddec __obf_319b72319cecddec
	// Data Mode (e.g. numeric).
	__obf_b08d86cf0c048279 []byte// segment data (e.g. "abc").

}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_a1f99035c479a58e struct {
	__obf_96356ace474cb30b int // Minimum & maximum versions supported.
	__obf_cefffff249e842b3 int
	__obf_b79a3e03cf940166 * // Mode indicator bit sequences.
	bitset.Bitset
	__obf_c261a78043989f54 *bitset.Bitset
	__obf_3f299aaed347bcdd *bitset.Bitset
	__obf_658acb240021e40d int // Character count lengths.
	__obf_785a2936c735d1ad int
	__obf_11f79fbb70f059e6 int
	__obf_b08d86cf0c048279 []byte// The raw input data.

	__obf_f3d1dd1a36be2be4 []__obf_61a731c6456893e0// The data classified into unoptimised segments.

	// The data classified into optimised segments.
	__obf_51ecedba44aa8ec2 []__obf_61a731c6456893e0
}

// newDataEncoder constructs a dataEncoder.
func __obf_11835c8cedb230b3(__obf_066728c6553a619e __obf_d9ac8a0ee0d0f27f) *__obf_a1f99035c479a58e {
	__obf_f4afc55b009a3568 := &__obf_a1f99035c479a58e{}

	switch __obf_066728c6553a619e {
	case __obf_96e9dfef6bbb06ba:
		__obf_f4afc55b009a3568 = &__obf_a1f99035c479a58e{__obf_96356ace474cb30b: 1, __obf_cefffff249e842b3: 9, __obf_b79a3e03cf940166: bitset.New(__obf_584bba5d11113643, __obf_584bba5d11113643, __obf_584bba5d11113643, __obf_0249afb330f83509), __obf_c261a78043989f54: bitset.New(__obf_584bba5d11113643, __obf_584bba5d11113643, __obf_0249afb330f83509, __obf_584bba5d11113643), __obf_3f299aaed347bcdd: bitset.New(__obf_584bba5d11113643, __obf_0249afb330f83509, __obf_584bba5d11113643, __obf_584bba5d11113643), __obf_658acb240021e40d: 10, __obf_785a2936c735d1ad: 9, __obf_11f79fbb70f059e6: 8}
	case __obf_bd631aba8d4005cf:
		__obf_f4afc55b009a3568 = &__obf_a1f99035c479a58e{__obf_96356ace474cb30b: 10, __obf_cefffff249e842b3: 26, __obf_b79a3e03cf940166: bitset.New(__obf_584bba5d11113643, __obf_584bba5d11113643, __obf_584bba5d11113643, __obf_0249afb330f83509), __obf_c261a78043989f54: bitset.New(__obf_584bba5d11113643, __obf_584bba5d11113643, __obf_0249afb330f83509, __obf_584bba5d11113643), __obf_3f299aaed347bcdd: bitset.New(__obf_584bba5d11113643, __obf_0249afb330f83509, __obf_584bba5d11113643, __obf_584bba5d11113643), __obf_658acb240021e40d: 12, __obf_785a2936c735d1ad: 11, __obf_11f79fbb70f059e6: 16}
	case __obf_5da2f2a34edb06bc:
		__obf_f4afc55b009a3568 = &__obf_a1f99035c479a58e{__obf_96356ace474cb30b: 27, __obf_cefffff249e842b3: 40, __obf_b79a3e03cf940166: bitset.New(__obf_584bba5d11113643, __obf_584bba5d11113643, __obf_584bba5d11113643, __obf_0249afb330f83509), __obf_c261a78043989f54: bitset.New(__obf_584bba5d11113643, __obf_584bba5d11113643, __obf_0249afb330f83509, __obf_584bba5d11113643), __obf_3f299aaed347bcdd: bitset.New(__obf_584bba5d11113643, __obf_0249afb330f83509, __obf_584bba5d11113643, __obf_584bba5d11113643), __obf_658acb240021e40d: 14, __obf_785a2936c735d1ad: 13, __obf_11f79fbb70f059e6: 16}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_f4afc55b009a3568
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_7626f31ec1adfcf5(__obf_b08d86cf0c048279 []byte) (*bitset.Bitset, error) {
	__obf_f4afc55b009a3568.__obf_b08d86cf0c048279 = __obf_b08d86cf0c048279
	__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4 = nil
	__obf_f4afc55b009a3568.__obf_51ecedba44aa8ec2 = nil

	if len(__obf_b08d86cf0c048279) == 0 {
		return nil, errors.New("no data to encode")
	}
	__obf_93d69090667e2328 := // Classify data into unoptimised segments.
		__obf_f4afc55b009a3568.__obf_1273ebd275828441()
	__obf_c88308f5c3b1b66a := // Optimise segments.
		__obf_f4afc55b009a3568.__obf_d011521c6fa18cab()
	if __obf_c88308f5c3b1b66a != nil {
		return nil, __obf_c88308f5c3b1b66a
	}
	__obf_3ab763a07cbdd9b4 := // Check if a single byte encoded segment would be more efficient.
		0
	for _, __obf_e988dab6de39cff8 := range __obf_f4afc55b009a3568.__obf_51ecedba44aa8ec2 {
		__obf_ebfaa7ff95572879, __obf_c88308f5c3b1b66a := __obf_f4afc55b009a3568.__obf_450a56f409c91c82(__obf_e988dab6de39cff8.__obf_319b72319cecddec, len(__obf_e988dab6de39cff8.__obf_b08d86cf0c048279))
		if __obf_c88308f5c3b1b66a != nil {
			return nil, __obf_c88308f5c3b1b66a
		}
		__obf_3ab763a07cbdd9b4 += __obf_ebfaa7ff95572879
	}
	__obf_4eb306bd5d7cb879, __obf_c88308f5c3b1b66a := __obf_f4afc55b009a3568.__obf_450a56f409c91c82(__obf_93d69090667e2328, len(__obf_f4afc55b009a3568.__obf_b08d86cf0c048279))
	if __obf_c88308f5c3b1b66a != nil {
		return nil, __obf_c88308f5c3b1b66a
	}

	if __obf_4eb306bd5d7cb879 <= __obf_3ab763a07cbdd9b4 {
		__obf_f4afc55b009a3568.__obf_51ecedba44aa8ec2 = []__obf_61a731c6456893e0{{__obf_319b72319cecddec: __obf_93d69090667e2328, __obf_b08d86cf0c048279: __obf_f4afc55b009a3568.

			// Encode data.
			__obf_b08d86cf0c048279}}
	}
	__obf_9b1dcfbd17fe9618 := bitset.New()
	for _, __obf_e988dab6de39cff8 := range __obf_f4afc55b009a3568.__obf_51ecedba44aa8ec2 {
		__obf_f4afc55b009a3568.__obf_995fa9722e1e3e94(__obf_e988dab6de39cff8.__obf_b08d86cf0c048279, __obf_e988dab6de39cff8.

			// classifyDataModes classifies the raw data into unoptimised segments.
			// e.g. "123ZZ#!#!" =>
			// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
			//
			// Returns the highest data mode needed to encode the data. e.g. for a mixed
			// numeric/alphanumeric input, the highest is alphanumeric.
			//
			// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
			__obf_319b72319cecddec, __obf_9b1dcfbd17fe9618)
	}

	return __obf_9b1dcfbd17fe9618, nil
}

func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_1273ebd275828441() __obf_319b72319cecddec {
	var __obf_6a6a77e65664e51f int
	__obf_80e7217120235b62 := __obf_c24c42c85933c4c7
	__obf_93d69090667e2328 := __obf_80e7217120235b62

	for __obf_2e4170383a5ddec7, __obf_283aa709fd2621d8 := range __obf_f4afc55b009a3568.__obf_b08d86cf0c048279 {
		__obf_0586bff182bd0146 := __obf_c24c42c85933c4c7
		switch {
		case __obf_283aa709fd2621d8 >= 0x30 && __obf_283aa709fd2621d8 <= 0x39:
			__obf_0586bff182bd0146 = __obf_917a5ab78c19aa97
		case __obf_283aa709fd2621d8 == 0x20 || __obf_283aa709fd2621d8 == 0x24 || __obf_283aa709fd2621d8 == 0x25 || __obf_283aa709fd2621d8 == 0x2a || __obf_283aa709fd2621d8 == 0x2b || __obf_283aa709fd2621d8 ==
			0x2d || __obf_283aa709fd2621d8 == 0x2e || __obf_283aa709fd2621d8 == 0x2f || __obf_283aa709fd2621d8 == 0x3a || (__obf_283aa709fd2621d8 >= 0x41 && __obf_283aa709fd2621d8 <= 0x5a):
			__obf_0586bff182bd0146 = __obf_23eff39094784930
		default:
			__obf_0586bff182bd0146 = __obf_105296cdda013cec
		}

		if __obf_0586bff182bd0146 != __obf_80e7217120235b62 {
			if __obf_2e4170383a5ddec7 > 0 {
				__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4 = append(__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4, __obf_61a731c6456893e0{__obf_319b72319cecddec: __obf_80e7217120235b62, __obf_b08d86cf0c048279: __obf_f4afc55b009a3568.__obf_b08d86cf0c048279[__obf_6a6a77e65664e51f:__obf_2e4170383a5ddec7]})
				__obf_6a6a77e65664e51f = __obf_2e4170383a5ddec7
			}
			__obf_80e7217120235b62 = __obf_0586bff182bd0146
		}

		if __obf_0586bff182bd0146 > __obf_93d69090667e2328 {
			__obf_93d69090667e2328 = __obf_0586bff182bd0146
		}
	}
	__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4 = append(__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4, __obf_61a731c6456893e0{__obf_319b72319cecddec: __obf_80e7217120235b62, __obf_b08d86cf0c048279: __obf_f4afc55b009a3568.__obf_b08d86cf0c048279[__obf_6a6a77e65664e51f:len(__obf_f4afc55b009a3568.__obf_b08d86cf0c048279)]})

	return __obf_93d69090667e2328
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
func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_d011521c6fa18cab() error {
	for __obf_2e4170383a5ddec7 := 0; __obf_2e4170383a5ddec7 < len(__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4); {
		__obf_80e7217120235b62 := __obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4[__obf_2e4170383a5ddec7].__obf_319b72319cecddec
		__obf_60d3991259b492b0 := len(__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4[__obf_2e4170383a5ddec7].__obf_b08d86cf0c048279)
		__obf_95b5b5fef8a19647 := __obf_2e4170383a5ddec7 + 1
		for __obf_95b5b5fef8a19647 < len(__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4) {
			__obf_0dfab613b3881e50 := len(__obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4[__obf_95b5b5fef8a19647].__obf_b08d86cf0c048279)
			__obf_81c02393a42f7561 := __obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4[__obf_95b5b5fef8a19647].__obf_319b72319cecddec

			if __obf_81c02393a42f7561 > __obf_80e7217120235b62 {
				break
			}
			__obf_7afff957137cb0ad, __obf_c88308f5c3b1b66a := __obf_f4afc55b009a3568.__obf_450a56f409c91c82(__obf_80e7217120235b62, __obf_60d3991259b492b0+__obf_0dfab613b3881e50)

			if __obf_c88308f5c3b1b66a != nil {
				return __obf_c88308f5c3b1b66a
			}
			__obf_e692ad6a050d4bc6, __obf_c88308f5c3b1b66a := __obf_f4afc55b009a3568.__obf_450a56f409c91c82(__obf_80e7217120235b62, __obf_60d3991259b492b0)

			if __obf_c88308f5c3b1b66a != nil {
				return __obf_c88308f5c3b1b66a
			}
			__obf_ac9b9d902c8cb11a, __obf_c88308f5c3b1b66a := __obf_f4afc55b009a3568.__obf_450a56f409c91c82(__obf_81c02393a42f7561, __obf_0dfab613b3881e50)

			if __obf_c88308f5c3b1b66a != nil {
				return __obf_c88308f5c3b1b66a
			}

			if __obf_7afff957137cb0ad < __obf_e692ad6a050d4bc6+__obf_ac9b9d902c8cb11a {
				__obf_95b5b5fef8a19647++
				__obf_60d3991259b492b0 += __obf_0dfab613b3881e50
			} else {
				break
			}
		}
		__obf_51ecedba44aa8ec2 := __obf_61a731c6456893e0{__obf_319b72319cecddec: __obf_80e7217120235b62, __obf_b08d86cf0c048279: make([]byte, 0, __obf_60d3991259b492b0)}

		for __obf_c4c99244a3ebad64 := __obf_2e4170383a5ddec7; __obf_c4c99244a3ebad64 < __obf_95b5b5fef8a19647; __obf_c4c99244a3ebad64++ {
			__obf_51ecedba44aa8ec2.__obf_b08d86cf0c048279 = append(__obf_51ecedba44aa8ec2.__obf_b08d86cf0c048279, __obf_f4afc55b009a3568.__obf_f3d1dd1a36be2be4[__obf_c4c99244a3ebad64].__obf_b08d86cf0c048279...)
		}
		__obf_f4afc55b009a3568.__obf_51ecedba44aa8ec2 = append(__obf_f4afc55b009a3568.__obf_51ecedba44aa8ec2, __obf_51ecedba44aa8ec2)
		__obf_2e4170383a5ddec7 = __obf_95b5b5fef8a19647
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_995fa9722e1e3e94(__obf_b08d86cf0c048279 []byte, __obf_319b72319cecddec __obf_319b72319cecddec, __obf_9b1dcfbd17fe9618 *bitset.Bitset) {
	__obf_a8afedf204c1c547 := __obf_f4afc55b009a3568.__obf_a8afedf204c1c547(__obf_319b72319cecddec)
	__obf_3458fa791dc0b9a0 := __obf_f4afc55b009a3568.__obf_3458fa791dc0b9a0(__obf_319b72319cecddec)
	__obf_9b1dcfbd17fe9618.

		// Append mode indicator.
		Append(__obf_a8afedf204c1c547)
	__obf_9b1dcfbd17fe9618.

		// Append character count.
		AppendUint32(uint32(len(__obf_b08d86cf0c048279)), __obf_3458fa791dc0b9a0)

	// Append data.
	switch __obf_319b72319cecddec {
	case __obf_917a5ab78c19aa97:
		for __obf_2e4170383a5ddec7 := 0; __obf_2e4170383a5ddec7 < len(__obf_b08d86cf0c048279); __obf_2e4170383a5ddec7 += 3 {
			__obf_a54e890fd8abec6a := len(__obf_b08d86cf0c048279) - __obf_2e4170383a5ddec7

			var __obf_6c9b4b130cbe971d uint32
			__obf_cce0522023e75233 := 1

			for __obf_95b5b5fef8a19647 := 0; __obf_95b5b5fef8a19647 < __obf_a54e890fd8abec6a && __obf_95b5b5fef8a19647 < 3; __obf_95b5b5fef8a19647++ {
				__obf_6c9b4b130cbe971d *= 10
				__obf_6c9b4b130cbe971d += uint32(__obf_b08d86cf0c048279[__obf_2e4170383a5ddec7+__obf_95b5b5fef8a19647] - 0x30)
				__obf_cce0522023e75233 += 3
			}
			__obf_9b1dcfbd17fe9618.
				AppendUint32(__obf_6c9b4b130cbe971d, __obf_cce0522023e75233)
		}
	case __obf_23eff39094784930:
		for __obf_2e4170383a5ddec7 := 0; __obf_2e4170383a5ddec7 < len(__obf_b08d86cf0c048279); __obf_2e4170383a5ddec7 += 2 {
			__obf_a54e890fd8abec6a := len(__obf_b08d86cf0c048279) - __obf_2e4170383a5ddec7

			var __obf_6c9b4b130cbe971d uint32
			for __obf_95b5b5fef8a19647 := 0; __obf_95b5b5fef8a19647 < __obf_a54e890fd8abec6a && __obf_95b5b5fef8a19647 < 2; __obf_95b5b5fef8a19647++ {
				__obf_6c9b4b130cbe971d *= 45
				__obf_6c9b4b130cbe971d += __obf_a95688ceda5c4487(__obf_b08d86cf0c048279[__obf_2e4170383a5ddec7+__obf_95b5b5fef8a19647])
			}
			__obf_cce0522023e75233 := 6
			if __obf_a54e890fd8abec6a > 1 {
				__obf_cce0522023e75233 = 11
			}
			__obf_9b1dcfbd17fe9618.
				AppendUint32(__obf_6c9b4b130cbe971d, __obf_cce0522023e75233)
		}
	case __obf_105296cdda013cec:
		for _, __obf_32f895d189b42b64 := range __obf_b08d86cf0c048279 {
			__obf_9b1dcfbd17fe9618.
				AppendByte(__obf_32f895d189b42b64, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_a8afedf204c1c547(__obf_319b72319cecddec __obf_319b72319cecddec) *bitset.Bitset {
	switch __obf_319b72319cecddec {
	case __obf_917a5ab78c19aa97:
		return __obf_f4afc55b009a3568.__obf_b79a3e03cf940166
	case __obf_23eff39094784930:
		return __obf_f4afc55b009a3568.__obf_c261a78043989f54
	case __obf_105296cdda013cec:
		return __obf_f4afc55b009a3568.__obf_3f299aaed347bcdd
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_3458fa791dc0b9a0(__obf_319b72319cecddec __obf_319b72319cecddec) int {
	switch __obf_319b72319cecddec {
	case __obf_917a5ab78c19aa97:
		return __obf_f4afc55b009a3568.__obf_658acb240021e40d
	case __obf_23eff39094784930:
		return __obf_f4afc55b009a3568.__obf_785a2936c735d1ad
	case __obf_105296cdda013cec:
		return __obf_f4afc55b009a3568.__obf_11f79fbb70f059e6
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
func (__obf_f4afc55b009a3568 *__obf_a1f99035c479a58e) __obf_450a56f409c91c82(__obf_319b72319cecddec __obf_319b72319cecddec, __obf_15c5ace1af53f46b int) (int, error) {
	__obf_a8afedf204c1c547 := __obf_f4afc55b009a3568.__obf_a8afedf204c1c547(__obf_319b72319cecddec)
	__obf_3458fa791dc0b9a0 := __obf_f4afc55b009a3568.__obf_3458fa791dc0b9a0(__obf_319b72319cecddec)

	if __obf_a8afedf204c1c547 == nil {
		return 0, errors.New("mode not supported")
	}
	__obf_e7360c82c37f2fb7 := (1 << uint8(__obf_3458fa791dc0b9a0)) - 1

	if __obf_15c5ace1af53f46b > __obf_e7360c82c37f2fb7 {
		return 0, errors.New("length too long to be represented")
	}
	__obf_ebfaa7ff95572879 := __obf_a8afedf204c1c547.Len() + __obf_3458fa791dc0b9a0

	switch __obf_319b72319cecddec {
	case __obf_917a5ab78c19aa97:
		__obf_ebfaa7ff95572879 += 10 * (__obf_15c5ace1af53f46b / 3)

		if __obf_15c5ace1af53f46b%3 != 0 {
			__obf_ebfaa7ff95572879 += 1 + 3*(__obf_15c5ace1af53f46b%3)
		}
	case __obf_23eff39094784930:
		__obf_ebfaa7ff95572879 += 11 * (__obf_15c5ace1af53f46b / 2)
		__obf_ebfaa7ff95572879 += 6 * (__obf_15c5ace1af53f46b % 2)
	case __obf_105296cdda013cec:
		__obf_ebfaa7ff95572879 += 8 * __obf_15c5ace1af53f46b
	}

	return __obf_ebfaa7ff95572879, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_a95688ceda5c4487(__obf_283aa709fd2621d8 byte) uint32 {
	__obf_2cb1880a1bbf7299 := uint32(__obf_283aa709fd2621d8)

	switch {
	case __obf_2cb1880a1bbf7299 >= '0' && __obf_2cb1880a1bbf7299 <= '9':
		// 0-9 encoded as 0-9.
		return __obf_2cb1880a1bbf7299 - '0'
	case __obf_2cb1880a1bbf7299 >= 'A' && __obf_2cb1880a1bbf7299 <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_2cb1880a1bbf7299 - 'A' + 10
	case __obf_2cb1880a1bbf7299 == ' ':
		return 36
	case __obf_2cb1880a1bbf7299 == '$':
		return 37
	case __obf_2cb1880a1bbf7299 == '%':
		return 38
	case __obf_2cb1880a1bbf7299 == '*':
		return 39
	case __obf_2cb1880a1bbf7299 == '+':
		return 40
	case __obf_2cb1880a1bbf7299 == '-':
		return 41
	case __obf_2cb1880a1bbf7299 == '.':
		return 42
	case __obf_2cb1880a1bbf7299 == '/':
		return 43
	case __obf_2cb1880a1bbf7299 == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_283aa709fd2621d8)
	}

	return 0
}
