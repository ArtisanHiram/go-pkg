// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_48b3d686b985a20b

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
type __obf_1241f5f5324d4773 uint8

const (
	__obf_c26ae8e2022a969a __obf_1241f5f5324d4773 =
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	1 << iota
	__obf_e0eedf57edbfa53a
	__obf_45466721b316305e
	__obf_66671368f679acff
)

// dataModeString returns d as a short printable string.
func __obf_6f71cd1d7cbc8220(__obf_1b7831cf10b4f76e __obf_1241f5f5324d4773) string {
	switch __obf_1b7831cf10b4f76e {
	case __obf_c26ae8e2022a969a:
		return "none"
	case __obf_e0eedf57edbfa53a:
		return "numeric"
	case __obf_45466721b316305e:
		return "alphanumeric"
	case __obf_66671368f679acff:
		return "byte"
	}

	return "unknown"
}

type __obf_2818b04b6d379c6a uint8

const (
	__obf_2ddbe5e0d50f3fdb __obf_2818b04b6d379c6a = iota
	__obf_d4c11dce5ccd0394
	__obf_e21e9a063b84f8a4
)

// segment is a single segment of data.
type __obf_fff5688077fd97ac struct {
	__obf_1241f5f5324d4773 __obf_1241f5f5324d4773
	// Data Mode (e.g. numeric).
	__obf_e4df198c6616985f []byte// segment data (e.g. "abc").

}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_5a535fd3d1104862 struct {
	__obf_d1067ff1d326e575 int // Minimum & maximum versions supported.
	__obf_5ff771e03c3e7b0e int
	__obf_386495099f57611a * // Mode indicator bit sequences.
	bitset.Bitset
	__obf_cc5d7e2e10037557 *bitset.Bitset
	__obf_143be839a85c11ab *bitset.Bitset
	__obf_820dd657c45d6563 int // Character count lengths.
	__obf_3fe1cfd3ed20cee3 int
	__obf_8917c71548f7da32 int
	__obf_e4df198c6616985f []byte// The raw input data.

	__obf_deaa58bc0ec97958 []__obf_fff5688077fd97ac// The data classified into unoptimised segments.

	// The data classified into optimised segments.
	__obf_03f9ed3f4a58dcfc []__obf_fff5688077fd97ac
}

// newDataEncoder constructs a dataEncoder.
func __obf_fccacc671db5d5f3(__obf_05af1e63275cf6cc __obf_2818b04b6d379c6a) *__obf_5a535fd3d1104862 {
	__obf_1b7831cf10b4f76e := &__obf_5a535fd3d1104862{}

	switch __obf_05af1e63275cf6cc {
	case __obf_2ddbe5e0d50f3fdb:
		__obf_1b7831cf10b4f76e = &__obf_5a535fd3d1104862{__obf_d1067ff1d326e575: 1, __obf_5ff771e03c3e7b0e: 9, __obf_386495099f57611a: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614), __obf_cc5d7e2e10037557: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614, __obf_4c3efb5a9adb8cf4), __obf_143be839a85c11ab: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614, __obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4), __obf_820dd657c45d6563: 10, __obf_3fe1cfd3ed20cee3: 9, __obf_8917c71548f7da32: 8}
	case __obf_d4c11dce5ccd0394:
		__obf_1b7831cf10b4f76e = &__obf_5a535fd3d1104862{__obf_d1067ff1d326e575: 10, __obf_5ff771e03c3e7b0e: 26, __obf_386495099f57611a: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614), __obf_cc5d7e2e10037557: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614, __obf_4c3efb5a9adb8cf4), __obf_143be839a85c11ab: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614, __obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4), __obf_820dd657c45d6563: 12, __obf_3fe1cfd3ed20cee3: 11, __obf_8917c71548f7da32: 16}
	case __obf_e21e9a063b84f8a4:
		__obf_1b7831cf10b4f76e = &__obf_5a535fd3d1104862{__obf_d1067ff1d326e575: 27, __obf_5ff771e03c3e7b0e: 40, __obf_386495099f57611a: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614), __obf_cc5d7e2e10037557: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614, __obf_4c3efb5a9adb8cf4), __obf_143be839a85c11ab: bitset.New(__obf_4c3efb5a9adb8cf4, __obf_42aad36989ea5614, __obf_4c3efb5a9adb8cf4, __obf_4c3efb5a9adb8cf4), __obf_820dd657c45d6563: 14, __obf_3fe1cfd3ed20cee3: 13, __obf_8917c71548f7da32: 16}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_1b7831cf10b4f76e
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_5fa94477cdc64b8f(__obf_e4df198c6616985f []byte) (*bitset.Bitset, error) {
	__obf_1b7831cf10b4f76e.__obf_e4df198c6616985f = __obf_e4df198c6616985f
	__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958 = nil
	__obf_1b7831cf10b4f76e.__obf_03f9ed3f4a58dcfc = nil

	if len(__obf_e4df198c6616985f) == 0 {
		return nil, errors.New("no data to encode")
	}
	__obf_ee4c7f92d1cfd415 := // Classify data into unoptimised segments.
		__obf_1b7831cf10b4f76e.__obf_768537528ffaa10d()
	__obf_3f84ccce09b63a68 := // Optimise segments.
		__obf_1b7831cf10b4f76e.__obf_1465094d6e87eaa1()
	if __obf_3f84ccce09b63a68 != nil {
		return nil, __obf_3f84ccce09b63a68
	}
	__obf_a28dad8a88ebbb13 := // Check if a single byte encoded segment would be more efficient.
		0
	for _, __obf_d73c540e0bdff21b := range __obf_1b7831cf10b4f76e.__obf_03f9ed3f4a58dcfc {
		__obf_2a2edcd8f8edb79b, __obf_3f84ccce09b63a68 := __obf_1b7831cf10b4f76e.__obf_894c98e72d7dd31f(__obf_d73c540e0bdff21b.__obf_1241f5f5324d4773, len(__obf_d73c540e0bdff21b.__obf_e4df198c6616985f))
		if __obf_3f84ccce09b63a68 != nil {
			return nil, __obf_3f84ccce09b63a68
		}
		__obf_a28dad8a88ebbb13 += __obf_2a2edcd8f8edb79b
	}
	__obf_010d2a63044bd00d, __obf_3f84ccce09b63a68 := __obf_1b7831cf10b4f76e.__obf_894c98e72d7dd31f(__obf_ee4c7f92d1cfd415, len(__obf_1b7831cf10b4f76e.__obf_e4df198c6616985f))
	if __obf_3f84ccce09b63a68 != nil {
		return nil, __obf_3f84ccce09b63a68
	}

	if __obf_010d2a63044bd00d <= __obf_a28dad8a88ebbb13 {
		__obf_1b7831cf10b4f76e.__obf_03f9ed3f4a58dcfc = []__obf_fff5688077fd97ac{{__obf_1241f5f5324d4773: __obf_ee4c7f92d1cfd415, __obf_e4df198c6616985f: __obf_1b7831cf10b4f76e.

			// Encode data.
			__obf_e4df198c6616985f}}
	}
	__obf_ba2373eb7d967534 := bitset.New()
	for _, __obf_d73c540e0bdff21b := range __obf_1b7831cf10b4f76e.__obf_03f9ed3f4a58dcfc {
		__obf_1b7831cf10b4f76e.__obf_726a9adbacdf5468(__obf_d73c540e0bdff21b.__obf_e4df198c6616985f, __obf_d73c540e0bdff21b.

			// classifyDataModes classifies the raw data into unoptimised segments.
			// e.g. "123ZZ#!#!" =>
			// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
			//
			// Returns the highest data mode needed to encode the data. e.g. for a mixed
			// numeric/alphanumeric input, the highest is alphanumeric.
			//
			// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
			__obf_1241f5f5324d4773, __obf_ba2373eb7d967534)
	}

	return __obf_ba2373eb7d967534, nil
}

func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_768537528ffaa10d() __obf_1241f5f5324d4773 {
	var __obf_b1d9b56dd18d11cd int
	__obf_a2f2a0229f69760c := __obf_c26ae8e2022a969a
	__obf_ee4c7f92d1cfd415 := __obf_a2f2a0229f69760c

	for __obf_8715c27aa2bf2586, __obf_f85dcc6c0aea860b := range __obf_1b7831cf10b4f76e.__obf_e4df198c6616985f {
		__obf_736f4fa4daaf4bf6 := __obf_c26ae8e2022a969a
		switch {
		case __obf_f85dcc6c0aea860b >= 0x30 && __obf_f85dcc6c0aea860b <= 0x39:
			__obf_736f4fa4daaf4bf6 = __obf_e0eedf57edbfa53a
		case __obf_f85dcc6c0aea860b == 0x20 || __obf_f85dcc6c0aea860b == 0x24 || __obf_f85dcc6c0aea860b == 0x25 || __obf_f85dcc6c0aea860b == 0x2a || __obf_f85dcc6c0aea860b == 0x2b || __obf_f85dcc6c0aea860b ==
			0x2d || __obf_f85dcc6c0aea860b == 0x2e || __obf_f85dcc6c0aea860b == 0x2f || __obf_f85dcc6c0aea860b == 0x3a || (__obf_f85dcc6c0aea860b >= 0x41 && __obf_f85dcc6c0aea860b <= 0x5a):
			__obf_736f4fa4daaf4bf6 = __obf_45466721b316305e
		default:
			__obf_736f4fa4daaf4bf6 = __obf_66671368f679acff
		}

		if __obf_736f4fa4daaf4bf6 != __obf_a2f2a0229f69760c {
			if __obf_8715c27aa2bf2586 > 0 {
				__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958 = append(__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958, __obf_fff5688077fd97ac{__obf_1241f5f5324d4773: __obf_a2f2a0229f69760c, __obf_e4df198c6616985f: __obf_1b7831cf10b4f76e.__obf_e4df198c6616985f[__obf_b1d9b56dd18d11cd:__obf_8715c27aa2bf2586]})
				__obf_b1d9b56dd18d11cd = __obf_8715c27aa2bf2586
			}
			__obf_a2f2a0229f69760c = __obf_736f4fa4daaf4bf6
		}

		if __obf_736f4fa4daaf4bf6 > __obf_ee4c7f92d1cfd415 {
			__obf_ee4c7f92d1cfd415 = __obf_736f4fa4daaf4bf6
		}
	}
	__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958 = append(__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958, __obf_fff5688077fd97ac{__obf_1241f5f5324d4773: __obf_a2f2a0229f69760c, __obf_e4df198c6616985f: __obf_1b7831cf10b4f76e.__obf_e4df198c6616985f[__obf_b1d9b56dd18d11cd:len(__obf_1b7831cf10b4f76e.__obf_e4df198c6616985f)]})

	return __obf_ee4c7f92d1cfd415
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
func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_1465094d6e87eaa1() error {
	for __obf_8715c27aa2bf2586 := 0; __obf_8715c27aa2bf2586 < len(__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958); {
		__obf_a2f2a0229f69760c := __obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958[__obf_8715c27aa2bf2586].__obf_1241f5f5324d4773
		__obf_e6cf716fe53ced3f := len(__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958[__obf_8715c27aa2bf2586].__obf_e4df198c6616985f)
		__obf_02acc1a50ace36fe := __obf_8715c27aa2bf2586 + 1
		for __obf_02acc1a50ace36fe < len(__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958) {
			__obf_d107203f2f4ed37c := len(__obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958[__obf_02acc1a50ace36fe].__obf_e4df198c6616985f)
			__obf_149425be1d0004d9 := __obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958[__obf_02acc1a50ace36fe].__obf_1241f5f5324d4773

			if __obf_149425be1d0004d9 > __obf_a2f2a0229f69760c {
				break
			}
			__obf_62dba9acbe3da0e9, __obf_3f84ccce09b63a68 := __obf_1b7831cf10b4f76e.__obf_894c98e72d7dd31f(__obf_a2f2a0229f69760c, __obf_e6cf716fe53ced3f+__obf_d107203f2f4ed37c)

			if __obf_3f84ccce09b63a68 != nil {
				return __obf_3f84ccce09b63a68
			}
			__obf_f10d926f0d7181fa, __obf_3f84ccce09b63a68 := __obf_1b7831cf10b4f76e.__obf_894c98e72d7dd31f(__obf_a2f2a0229f69760c, __obf_e6cf716fe53ced3f)

			if __obf_3f84ccce09b63a68 != nil {
				return __obf_3f84ccce09b63a68
			}
			__obf_467168e9b1a8d1ef, __obf_3f84ccce09b63a68 := __obf_1b7831cf10b4f76e.__obf_894c98e72d7dd31f(__obf_149425be1d0004d9, __obf_d107203f2f4ed37c)

			if __obf_3f84ccce09b63a68 != nil {
				return __obf_3f84ccce09b63a68
			}

			if __obf_62dba9acbe3da0e9 < __obf_f10d926f0d7181fa+__obf_467168e9b1a8d1ef {
				__obf_02acc1a50ace36fe++
				__obf_e6cf716fe53ced3f += __obf_d107203f2f4ed37c
			} else {
				break
			}
		}
		__obf_03f9ed3f4a58dcfc := __obf_fff5688077fd97ac{__obf_1241f5f5324d4773: __obf_a2f2a0229f69760c, __obf_e4df198c6616985f: make([]byte, 0, __obf_e6cf716fe53ced3f)}

		for __obf_cc7a8cff799972da := __obf_8715c27aa2bf2586; __obf_cc7a8cff799972da < __obf_02acc1a50ace36fe; __obf_cc7a8cff799972da++ {
			__obf_03f9ed3f4a58dcfc.__obf_e4df198c6616985f = append(__obf_03f9ed3f4a58dcfc.__obf_e4df198c6616985f, __obf_1b7831cf10b4f76e.__obf_deaa58bc0ec97958[__obf_cc7a8cff799972da].__obf_e4df198c6616985f...)
		}
		__obf_1b7831cf10b4f76e.__obf_03f9ed3f4a58dcfc = append(__obf_1b7831cf10b4f76e.__obf_03f9ed3f4a58dcfc, __obf_03f9ed3f4a58dcfc)
		__obf_8715c27aa2bf2586 = __obf_02acc1a50ace36fe
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_726a9adbacdf5468(__obf_e4df198c6616985f []byte, __obf_1241f5f5324d4773 __obf_1241f5f5324d4773, __obf_ba2373eb7d967534 *bitset.Bitset) {
	__obf_d43a5152ff67bc44 := __obf_1b7831cf10b4f76e.__obf_d43a5152ff67bc44(__obf_1241f5f5324d4773)
	__obf_41e0bb7ff38f3067 := __obf_1b7831cf10b4f76e.__obf_41e0bb7ff38f3067(__obf_1241f5f5324d4773)
	__obf_ba2373eb7d967534.

		// Append mode indicator.
		Append(__obf_d43a5152ff67bc44)
	__obf_ba2373eb7d967534.

		// Append character count.
		AppendUint32(uint32(len(__obf_e4df198c6616985f)), __obf_41e0bb7ff38f3067)

	// Append data.
	switch __obf_1241f5f5324d4773 {
	case __obf_e0eedf57edbfa53a:
		for __obf_8715c27aa2bf2586 := 0; __obf_8715c27aa2bf2586 < len(__obf_e4df198c6616985f); __obf_8715c27aa2bf2586 += 3 {
			__obf_a095280398dea06f := len(__obf_e4df198c6616985f) - __obf_8715c27aa2bf2586

			var __obf_04a7a2ca242933f3 uint32
			__obf_66c1cf6ec50c5766 := 1

			for __obf_02acc1a50ace36fe := 0; __obf_02acc1a50ace36fe < __obf_a095280398dea06f && __obf_02acc1a50ace36fe < 3; __obf_02acc1a50ace36fe++ {
				__obf_04a7a2ca242933f3 *= 10
				__obf_04a7a2ca242933f3 += uint32(__obf_e4df198c6616985f[__obf_8715c27aa2bf2586+__obf_02acc1a50ace36fe] - 0x30)
				__obf_66c1cf6ec50c5766 += 3
			}
			__obf_ba2373eb7d967534.
				AppendUint32(__obf_04a7a2ca242933f3, __obf_66c1cf6ec50c5766)
		}
	case __obf_45466721b316305e:
		for __obf_8715c27aa2bf2586 := 0; __obf_8715c27aa2bf2586 < len(__obf_e4df198c6616985f); __obf_8715c27aa2bf2586 += 2 {
			__obf_a095280398dea06f := len(__obf_e4df198c6616985f) - __obf_8715c27aa2bf2586

			var __obf_04a7a2ca242933f3 uint32
			for __obf_02acc1a50ace36fe := 0; __obf_02acc1a50ace36fe < __obf_a095280398dea06f && __obf_02acc1a50ace36fe < 2; __obf_02acc1a50ace36fe++ {
				__obf_04a7a2ca242933f3 *= 45
				__obf_04a7a2ca242933f3 += __obf_db3095340d293096(__obf_e4df198c6616985f[__obf_8715c27aa2bf2586+__obf_02acc1a50ace36fe])
			}
			__obf_66c1cf6ec50c5766 := 6
			if __obf_a095280398dea06f > 1 {
				__obf_66c1cf6ec50c5766 = 11
			}
			__obf_ba2373eb7d967534.
				AppendUint32(__obf_04a7a2ca242933f3, __obf_66c1cf6ec50c5766)
		}
	case __obf_66671368f679acff:
		for _, __obf_21b947e29586a23f := range __obf_e4df198c6616985f {
			__obf_ba2373eb7d967534.
				AppendByte(__obf_21b947e29586a23f, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_d43a5152ff67bc44(__obf_1241f5f5324d4773 __obf_1241f5f5324d4773) *bitset.Bitset {
	switch __obf_1241f5f5324d4773 {
	case __obf_e0eedf57edbfa53a:
		return __obf_1b7831cf10b4f76e.__obf_386495099f57611a
	case __obf_45466721b316305e:
		return __obf_1b7831cf10b4f76e.__obf_cc5d7e2e10037557
	case __obf_66671368f679acff:
		return __obf_1b7831cf10b4f76e.__obf_143be839a85c11ab
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_41e0bb7ff38f3067(__obf_1241f5f5324d4773 __obf_1241f5f5324d4773) int {
	switch __obf_1241f5f5324d4773 {
	case __obf_e0eedf57edbfa53a:
		return __obf_1b7831cf10b4f76e.__obf_820dd657c45d6563
	case __obf_45466721b316305e:
		return __obf_1b7831cf10b4f76e.__obf_3fe1cfd3ed20cee3
	case __obf_66671368f679acff:
		return __obf_1b7831cf10b4f76e.__obf_8917c71548f7da32
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
func (__obf_1b7831cf10b4f76e *__obf_5a535fd3d1104862) __obf_894c98e72d7dd31f(__obf_1241f5f5324d4773 __obf_1241f5f5324d4773, __obf_64a624c625796c8e int) (int, error) {
	__obf_d43a5152ff67bc44 := __obf_1b7831cf10b4f76e.__obf_d43a5152ff67bc44(__obf_1241f5f5324d4773)
	__obf_41e0bb7ff38f3067 := __obf_1b7831cf10b4f76e.__obf_41e0bb7ff38f3067(__obf_1241f5f5324d4773)

	if __obf_d43a5152ff67bc44 == nil {
		return 0, errors.New("mode not supported")
	}
	__obf_4f47575cda2dd9f0 := (1 << uint8(__obf_41e0bb7ff38f3067)) - 1

	if __obf_64a624c625796c8e > __obf_4f47575cda2dd9f0 {
		return 0, errors.New("length too long to be represented")
	}
	__obf_2a2edcd8f8edb79b := __obf_d43a5152ff67bc44.Len() + __obf_41e0bb7ff38f3067

	switch __obf_1241f5f5324d4773 {
	case __obf_e0eedf57edbfa53a:
		__obf_2a2edcd8f8edb79b += 10 * (__obf_64a624c625796c8e / 3)

		if __obf_64a624c625796c8e%3 != 0 {
			__obf_2a2edcd8f8edb79b += 1 + 3*(__obf_64a624c625796c8e%3)
		}
	case __obf_45466721b316305e:
		__obf_2a2edcd8f8edb79b += 11 * (__obf_64a624c625796c8e / 2)
		__obf_2a2edcd8f8edb79b += 6 * (__obf_64a624c625796c8e % 2)
	case __obf_66671368f679acff:
		__obf_2a2edcd8f8edb79b += 8 * __obf_64a624c625796c8e
	}

	return __obf_2a2edcd8f8edb79b, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_db3095340d293096(__obf_f85dcc6c0aea860b byte) uint32 {
	__obf_2674a63b142b0a83 := uint32(__obf_f85dcc6c0aea860b)

	switch {
	case __obf_2674a63b142b0a83 >= '0' && __obf_2674a63b142b0a83 <= '9':
		// 0-9 encoded as 0-9.
		return __obf_2674a63b142b0a83 - '0'
	case __obf_2674a63b142b0a83 >= 'A' && __obf_2674a63b142b0a83 <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_2674a63b142b0a83 - 'A' + 10
	case __obf_2674a63b142b0a83 == ' ':
		return 36
	case __obf_2674a63b142b0a83 == '$':
		return 37
	case __obf_2674a63b142b0a83 == '%':
		return 38
	case __obf_2674a63b142b0a83 == '*':
		return 39
	case __obf_2674a63b142b0a83 == '+':
		return 40
	case __obf_2674a63b142b0a83 == '-':
		return 41
	case __obf_2674a63b142b0a83 == '.':
		return 42
	case __obf_2674a63b142b0a83 == '/':
		return 43
	case __obf_2674a63b142b0a83 == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_f85dcc6c0aea860b)
	}

	return 0
}
