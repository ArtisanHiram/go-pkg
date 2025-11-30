// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_63dc7f87a7750084

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
type __obf_f113bf2789577d1b uint8

const (
	__obf_09d69fba407499ff __obf_f113bf2789577d1b =
	// Each dataMode is a subset of the subsequent dataMode:
	// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
	//
	// This ordering is important for determining which data modes a character can
	// be encoded with. E.g. 'E' can be encoded in both dataModeAlphanumeric and
	// dataModeByte.
	1 << iota
	__obf_bd761c2ad5f9c89e
	__obf_2625eb1c563eabe3
	__obf_a22a481b3c30a3f5
)

// dataModeString returns d as a short printable string.
func __obf_37281d064a2c235d(__obf_4a17dbc21446e4bc __obf_f113bf2789577d1b) string {
	switch __obf_4a17dbc21446e4bc {
	case __obf_09d69fba407499ff:
		return "none"
	case __obf_bd761c2ad5f9c89e:
		return "numeric"
	case __obf_2625eb1c563eabe3:
		return "alphanumeric"
	case __obf_a22a481b3c30a3f5:
		return "byte"
	}

	return "unknown"
}

type __obf_40700b751bbe7fa9 uint8

const (
	__obf_4ef346b16a795892 __obf_40700b751bbe7fa9 = iota
	__obf_e891267bd1bcc38a
	__obf_054184439c33471a
)

// segment is a single segment of data.
type __obf_0ec4e265b7c002f6 struct {
	__obf_f113bf2789577d1b __obf_f113bf2789577d1b
	// Data Mode (e.g. numeric).
	__obf_eb8b3f015415859a []byte// segment data (e.g. "abc").

}

// A dataEncoder encodes data for a particular QR Code version.
type __obf_3ee8d058dd75e4b0 struct {
	__obf_2cbe93d8b68961b6 int // Minimum & maximum versions supported.
	__obf_88dda467e700acbf int
	__obf_b671eb4a7bf71957 * // Mode indicator bit sequences.
	bitset.Bitset
	__obf_4d2937cc207f7910 *bitset.Bitset
	__obf_43451274f40671d8 *bitset.Bitset
	__obf_5ab0099443a7662a int // Character count lengths.
	__obf_b03cc86bc036ee75 int
	__obf_8d232ed2d06f96c3 int
	__obf_eb8b3f015415859a []byte// The raw input data.

	__obf_f2013991a748f1fa []__obf_0ec4e265b7c002f6// The data classified into unoptimised segments.

	// The data classified into optimised segments.
	__obf_1f7e3df0f5ac38b1 []__obf_0ec4e265b7c002f6
}

// newDataEncoder constructs a dataEncoder.
func __obf_d8bb8ebef294bcff(__obf_494cef4b0d092224 __obf_40700b751bbe7fa9) *__obf_3ee8d058dd75e4b0 {
	__obf_4a17dbc21446e4bc := &__obf_3ee8d058dd75e4b0{}

	switch __obf_494cef4b0d092224 {
	case __obf_4ef346b16a795892:
		__obf_4a17dbc21446e4bc = &__obf_3ee8d058dd75e4b0{__obf_2cbe93d8b68961b6: 1, __obf_88dda467e700acbf: 9, __obf_b671eb4a7bf71957: bitset.New(__obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313), __obf_4d2937cc207f7910: bitset.New(__obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313, __obf_7d446e3e8b73acba), __obf_43451274f40671d8: bitset.New(__obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313, __obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba), __obf_5ab0099443a7662a: 10, __obf_b03cc86bc036ee75: 9, __obf_8d232ed2d06f96c3: 8}
	case __obf_e891267bd1bcc38a:
		__obf_4a17dbc21446e4bc = &__obf_3ee8d058dd75e4b0{__obf_2cbe93d8b68961b6: 10, __obf_88dda467e700acbf: 26, __obf_b671eb4a7bf71957: bitset.New(__obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313), __obf_4d2937cc207f7910: bitset.New(__obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313, __obf_7d446e3e8b73acba), __obf_43451274f40671d8: bitset.New(__obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313, __obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba), __obf_5ab0099443a7662a: 12, __obf_b03cc86bc036ee75: 11, __obf_8d232ed2d06f96c3: 16}
	case __obf_054184439c33471a:
		__obf_4a17dbc21446e4bc = &__obf_3ee8d058dd75e4b0{__obf_2cbe93d8b68961b6: 27, __obf_88dda467e700acbf: 40, __obf_b671eb4a7bf71957: bitset.New(__obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313), __obf_4d2937cc207f7910: bitset.New(__obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313, __obf_7d446e3e8b73acba), __obf_43451274f40671d8: bitset.New(__obf_7d446e3e8b73acba, __obf_ebcd5b842c56f313, __obf_7d446e3e8b73acba, __obf_7d446e3e8b73acba), __obf_5ab0099443a7662a: 14, __obf_b03cc86bc036ee75: 13, __obf_8d232ed2d06f96c3: 16}
	default:
		log.Panic("Unknown dataEncoderType")
	}

	return __obf_4a17dbc21446e4bc
}

// encode data as one or more segments and return the encoded data.
//
// The returned data does not include the terminator bit sequence.
func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_395c063b7326115c(__obf_eb8b3f015415859a []byte) (*bitset.Bitset, error) {
	__obf_4a17dbc21446e4bc.__obf_eb8b3f015415859a = __obf_eb8b3f015415859a
	__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa = nil
	__obf_4a17dbc21446e4bc.__obf_1f7e3df0f5ac38b1 = nil

	if len(__obf_eb8b3f015415859a) == 0 {
		return nil, errors.New("no data to encode")
	}
	__obf_d8de624a0636ff09 := // Classify data into unoptimised segments.
		__obf_4a17dbc21446e4bc.__obf_1348ab4acb01b761()
	__obf_b59cbf05240604a0 := // Optimise segments.
		__obf_4a17dbc21446e4bc.__obf_f03ceed1496f30aa()
	if __obf_b59cbf05240604a0 != nil {
		return nil, __obf_b59cbf05240604a0
	}
	__obf_9c76b3afaf06308b := // Check if a single byte encoded segment would be more efficient.
		0
	for _, __obf_ab8d068af1cb47d6 := range __obf_4a17dbc21446e4bc.__obf_1f7e3df0f5ac38b1 {
		__obf_a7c399b5ad460699, __obf_b59cbf05240604a0 := __obf_4a17dbc21446e4bc.__obf_fdd4b0cec1dcb8de(__obf_ab8d068af1cb47d6.__obf_f113bf2789577d1b, len(__obf_ab8d068af1cb47d6.__obf_eb8b3f015415859a))
		if __obf_b59cbf05240604a0 != nil {
			return nil, __obf_b59cbf05240604a0
		}
		__obf_9c76b3afaf06308b += __obf_a7c399b5ad460699
	}
	__obf_0f3bd979d1d8e811, __obf_b59cbf05240604a0 := __obf_4a17dbc21446e4bc.__obf_fdd4b0cec1dcb8de(__obf_d8de624a0636ff09, len(__obf_4a17dbc21446e4bc.__obf_eb8b3f015415859a))
	if __obf_b59cbf05240604a0 != nil {
		return nil, __obf_b59cbf05240604a0
	}

	if __obf_0f3bd979d1d8e811 <= __obf_9c76b3afaf06308b {
		__obf_4a17dbc21446e4bc.__obf_1f7e3df0f5ac38b1 = []__obf_0ec4e265b7c002f6{{__obf_f113bf2789577d1b: __obf_d8de624a0636ff09, __obf_eb8b3f015415859a: __obf_4a17dbc21446e4bc.

			// Encode data.
			__obf_eb8b3f015415859a}}
	}
	__obf_175ed6e946c14a9b := bitset.New()
	for _, __obf_ab8d068af1cb47d6 := range __obf_4a17dbc21446e4bc.__obf_1f7e3df0f5ac38b1 {
		__obf_4a17dbc21446e4bc.__obf_679ca941efa46cf5(__obf_ab8d068af1cb47d6.__obf_eb8b3f015415859a, __obf_ab8d068af1cb47d6.

			// classifyDataModes classifies the raw data into unoptimised segments.
			// e.g. "123ZZ#!#!" =>
			// [numeric, 3, "123"] [alphanumeric, 2, "ZZ"] [byte, 4, "#!#!"].
			//
			// Returns the highest data mode needed to encode the data. e.g. for a mixed
			// numeric/alphanumeric input, the highest is alphanumeric.
			//
			// dataModeNone < dataModeNumeric < dataModeAlphanumeric < dataModeByte
			__obf_f113bf2789577d1b, __obf_175ed6e946c14a9b)
	}

	return __obf_175ed6e946c14a9b, nil
}

func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_1348ab4acb01b761() __obf_f113bf2789577d1b {
	var __obf_3a15471c26131738 int
	__obf_d00404d041f6debf := __obf_09d69fba407499ff
	__obf_d8de624a0636ff09 := __obf_d00404d041f6debf

	for __obf_b39cd0fd88eaf53c, __obf_d3308b8b9ec8593a := range __obf_4a17dbc21446e4bc.__obf_eb8b3f015415859a {
		__obf_13f3ff8bee53850d := __obf_09d69fba407499ff
		switch {
		case __obf_d3308b8b9ec8593a >= 0x30 && __obf_d3308b8b9ec8593a <= 0x39:
			__obf_13f3ff8bee53850d = __obf_bd761c2ad5f9c89e
		case __obf_d3308b8b9ec8593a == 0x20 || __obf_d3308b8b9ec8593a == 0x24 || __obf_d3308b8b9ec8593a == 0x25 || __obf_d3308b8b9ec8593a == 0x2a || __obf_d3308b8b9ec8593a == 0x2b || __obf_d3308b8b9ec8593a ==
			0x2d || __obf_d3308b8b9ec8593a == 0x2e || __obf_d3308b8b9ec8593a == 0x2f || __obf_d3308b8b9ec8593a == 0x3a || (__obf_d3308b8b9ec8593a >= 0x41 && __obf_d3308b8b9ec8593a <= 0x5a):
			__obf_13f3ff8bee53850d = __obf_2625eb1c563eabe3
		default:
			__obf_13f3ff8bee53850d = __obf_a22a481b3c30a3f5
		}

		if __obf_13f3ff8bee53850d != __obf_d00404d041f6debf {
			if __obf_b39cd0fd88eaf53c > 0 {
				__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa = append(__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa, __obf_0ec4e265b7c002f6{__obf_f113bf2789577d1b: __obf_d00404d041f6debf, __obf_eb8b3f015415859a: __obf_4a17dbc21446e4bc.__obf_eb8b3f015415859a[__obf_3a15471c26131738:__obf_b39cd0fd88eaf53c]})
				__obf_3a15471c26131738 = __obf_b39cd0fd88eaf53c
			}
			__obf_d00404d041f6debf = __obf_13f3ff8bee53850d
		}

		if __obf_13f3ff8bee53850d > __obf_d8de624a0636ff09 {
			__obf_d8de624a0636ff09 = __obf_13f3ff8bee53850d
		}
	}
	__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa = append(__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa, __obf_0ec4e265b7c002f6{__obf_f113bf2789577d1b: __obf_d00404d041f6debf, __obf_eb8b3f015415859a: __obf_4a17dbc21446e4bc.__obf_eb8b3f015415859a[__obf_3a15471c26131738:len(__obf_4a17dbc21446e4bc.__obf_eb8b3f015415859a)]})

	return __obf_d8de624a0636ff09
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
func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_f03ceed1496f30aa() error {
	for __obf_b39cd0fd88eaf53c := 0; __obf_b39cd0fd88eaf53c < len(__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa); {
		__obf_d00404d041f6debf := __obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa[__obf_b39cd0fd88eaf53c].__obf_f113bf2789577d1b
		__obf_d519da9fbd506b4a := len(__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa[__obf_b39cd0fd88eaf53c].__obf_eb8b3f015415859a)
		__obf_0d2b5442049f8295 := __obf_b39cd0fd88eaf53c + 1
		for __obf_0d2b5442049f8295 < len(__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa) {
			__obf_d536a047f435c509 := len(__obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa[__obf_0d2b5442049f8295].__obf_eb8b3f015415859a)
			__obf_1666333903b34abb := __obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa[__obf_0d2b5442049f8295].__obf_f113bf2789577d1b

			if __obf_1666333903b34abb > __obf_d00404d041f6debf {
				break
			}
			__obf_48d41c226f6ede65, __obf_b59cbf05240604a0 := __obf_4a17dbc21446e4bc.__obf_fdd4b0cec1dcb8de(__obf_d00404d041f6debf, __obf_d519da9fbd506b4a+__obf_d536a047f435c509)

			if __obf_b59cbf05240604a0 != nil {
				return __obf_b59cbf05240604a0
			}
			__obf_3a47f6bd1e26dcc1, __obf_b59cbf05240604a0 := __obf_4a17dbc21446e4bc.__obf_fdd4b0cec1dcb8de(__obf_d00404d041f6debf, __obf_d519da9fbd506b4a)

			if __obf_b59cbf05240604a0 != nil {
				return __obf_b59cbf05240604a0
			}
			__obf_da2a9c230b0721b6, __obf_b59cbf05240604a0 := __obf_4a17dbc21446e4bc.__obf_fdd4b0cec1dcb8de(__obf_1666333903b34abb, __obf_d536a047f435c509)

			if __obf_b59cbf05240604a0 != nil {
				return __obf_b59cbf05240604a0
			}

			if __obf_48d41c226f6ede65 < __obf_3a47f6bd1e26dcc1+__obf_da2a9c230b0721b6 {
				__obf_0d2b5442049f8295++
				__obf_d519da9fbd506b4a += __obf_d536a047f435c509
			} else {
				break
			}
		}
		__obf_1f7e3df0f5ac38b1 := __obf_0ec4e265b7c002f6{__obf_f113bf2789577d1b: __obf_d00404d041f6debf, __obf_eb8b3f015415859a: make([]byte, 0, __obf_d519da9fbd506b4a)}

		for __obf_87439be06a2b9f41 := __obf_b39cd0fd88eaf53c; __obf_87439be06a2b9f41 < __obf_0d2b5442049f8295; __obf_87439be06a2b9f41++ {
			__obf_1f7e3df0f5ac38b1.__obf_eb8b3f015415859a = append(__obf_1f7e3df0f5ac38b1.__obf_eb8b3f015415859a, __obf_4a17dbc21446e4bc.__obf_f2013991a748f1fa[__obf_87439be06a2b9f41].__obf_eb8b3f015415859a...)
		}
		__obf_4a17dbc21446e4bc.__obf_1f7e3df0f5ac38b1 = append(__obf_4a17dbc21446e4bc.__obf_1f7e3df0f5ac38b1, __obf_1f7e3df0f5ac38b1)
		__obf_b39cd0fd88eaf53c = __obf_0d2b5442049f8295
	}

	return nil
}

// encodeDataRaw encodes data in dataMode. The encoded data is appended to
// encoded.
func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_679ca941efa46cf5(__obf_eb8b3f015415859a []byte, __obf_f113bf2789577d1b __obf_f113bf2789577d1b, __obf_175ed6e946c14a9b *bitset.Bitset) {
	__obf_f05a72d84cdea69c := __obf_4a17dbc21446e4bc.__obf_f05a72d84cdea69c(__obf_f113bf2789577d1b)
	__obf_feba9402f06a9f6e := __obf_4a17dbc21446e4bc.__obf_feba9402f06a9f6e(__obf_f113bf2789577d1b)
	__obf_175ed6e946c14a9b.

		// Append mode indicator.
		Append(__obf_f05a72d84cdea69c)
	__obf_175ed6e946c14a9b.

		// Append character count.
		AppendUint32(uint32(len(__obf_eb8b3f015415859a)), __obf_feba9402f06a9f6e)

	// Append data.
	switch __obf_f113bf2789577d1b {
	case __obf_bd761c2ad5f9c89e:
		for __obf_b39cd0fd88eaf53c := 0; __obf_b39cd0fd88eaf53c < len(__obf_eb8b3f015415859a); __obf_b39cd0fd88eaf53c += 3 {
			__obf_cc421a5ea11aba08 := len(__obf_eb8b3f015415859a) - __obf_b39cd0fd88eaf53c

			var __obf_85088f87763437a5 uint32
			__obf_19c527cf1334800f := 1

			for __obf_0d2b5442049f8295 := 0; __obf_0d2b5442049f8295 < __obf_cc421a5ea11aba08 && __obf_0d2b5442049f8295 < 3; __obf_0d2b5442049f8295++ {
				__obf_85088f87763437a5 *= 10
				__obf_85088f87763437a5 += uint32(__obf_eb8b3f015415859a[__obf_b39cd0fd88eaf53c+__obf_0d2b5442049f8295] - 0x30)
				__obf_19c527cf1334800f += 3
			}
			__obf_175ed6e946c14a9b.
				AppendUint32(__obf_85088f87763437a5, __obf_19c527cf1334800f)
		}
	case __obf_2625eb1c563eabe3:
		for __obf_b39cd0fd88eaf53c := 0; __obf_b39cd0fd88eaf53c < len(__obf_eb8b3f015415859a); __obf_b39cd0fd88eaf53c += 2 {
			__obf_cc421a5ea11aba08 := len(__obf_eb8b3f015415859a) - __obf_b39cd0fd88eaf53c

			var __obf_85088f87763437a5 uint32
			for __obf_0d2b5442049f8295 := 0; __obf_0d2b5442049f8295 < __obf_cc421a5ea11aba08 && __obf_0d2b5442049f8295 < 2; __obf_0d2b5442049f8295++ {
				__obf_85088f87763437a5 *= 45
				__obf_85088f87763437a5 += __obf_0f3a5da3dd2f260a(__obf_eb8b3f015415859a[__obf_b39cd0fd88eaf53c+__obf_0d2b5442049f8295])
			}
			__obf_19c527cf1334800f := 6
			if __obf_cc421a5ea11aba08 > 1 {
				__obf_19c527cf1334800f = 11
			}
			__obf_175ed6e946c14a9b.
				AppendUint32(__obf_85088f87763437a5, __obf_19c527cf1334800f)
		}
	case __obf_a22a481b3c30a3f5:
		for _, __obf_0de44d307506fdeb := range __obf_eb8b3f015415859a {
			__obf_175ed6e946c14a9b.
				AppendByte(__obf_0de44d307506fdeb, 8)
		}
	}
}

// modeIndicator returns the segment header bits for a segment of type dataMode.
func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_f05a72d84cdea69c(__obf_f113bf2789577d1b __obf_f113bf2789577d1b) *bitset.Bitset {
	switch __obf_f113bf2789577d1b {
	case __obf_bd761c2ad5f9c89e:
		return __obf_4a17dbc21446e4bc.__obf_b671eb4a7bf71957
	case __obf_2625eb1c563eabe3:
		return __obf_4a17dbc21446e4bc.__obf_4d2937cc207f7910
	case __obf_a22a481b3c30a3f5:
		return __obf_4a17dbc21446e4bc.__obf_43451274f40671d8
	default:
		log.Panic("Unknown data mode")
	}

	return nil
}

// charCountBits returns the number of bits used to encode the length of a data
// segment of type dataMode.
func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_feba9402f06a9f6e(__obf_f113bf2789577d1b __obf_f113bf2789577d1b) int {
	switch __obf_f113bf2789577d1b {
	case __obf_bd761c2ad5f9c89e:
		return __obf_4a17dbc21446e4bc.__obf_5ab0099443a7662a
	case __obf_2625eb1c563eabe3:
		return __obf_4a17dbc21446e4bc.__obf_b03cc86bc036ee75
	case __obf_a22a481b3c30a3f5:
		return __obf_4a17dbc21446e4bc.__obf_8d232ed2d06f96c3
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
func (__obf_4a17dbc21446e4bc *__obf_3ee8d058dd75e4b0) __obf_fdd4b0cec1dcb8de(__obf_f113bf2789577d1b __obf_f113bf2789577d1b, __obf_14c860a7e9df876a int) (int, error) {
	__obf_f05a72d84cdea69c := __obf_4a17dbc21446e4bc.__obf_f05a72d84cdea69c(__obf_f113bf2789577d1b)
	__obf_feba9402f06a9f6e := __obf_4a17dbc21446e4bc.__obf_feba9402f06a9f6e(__obf_f113bf2789577d1b)

	if __obf_f05a72d84cdea69c == nil {
		return 0, errors.New("mode not supported")
	}
	__obf_22296e9bf66c783a := (1 << uint8(__obf_feba9402f06a9f6e)) - 1

	if __obf_14c860a7e9df876a > __obf_22296e9bf66c783a {
		return 0, errors.New("length too long to be represented")
	}
	__obf_a7c399b5ad460699 := __obf_f05a72d84cdea69c.Len() + __obf_feba9402f06a9f6e

	switch __obf_f113bf2789577d1b {
	case __obf_bd761c2ad5f9c89e:
		__obf_a7c399b5ad460699 += 10 * (__obf_14c860a7e9df876a / 3)

		if __obf_14c860a7e9df876a%3 != 0 {
			__obf_a7c399b5ad460699 += 1 + 3*(__obf_14c860a7e9df876a%3)
		}
	case __obf_2625eb1c563eabe3:
		__obf_a7c399b5ad460699 += 11 * (__obf_14c860a7e9df876a / 2)
		__obf_a7c399b5ad460699 += 6 * (__obf_14c860a7e9df876a % 2)
	case __obf_a22a481b3c30a3f5:
		__obf_a7c399b5ad460699 += 8 * __obf_14c860a7e9df876a
	}

	return __obf_a7c399b5ad460699, nil
}

// encodeAlphanumericChar returns the QR Code encoded value of v.
//
// v must be a QR Code defined alphanumeric character: 0-9, A-Z, SP, $%*+-./ or
// :. The characters are mapped to values in the range 0-44 respectively.
func __obf_0f3a5da3dd2f260a(__obf_d3308b8b9ec8593a byte) uint32 {
	__obf_cbe674ab72969f9d := uint32(__obf_d3308b8b9ec8593a)

	switch {
	case __obf_cbe674ab72969f9d >= '0' && __obf_cbe674ab72969f9d <= '9':
		// 0-9 encoded as 0-9.
		return __obf_cbe674ab72969f9d - '0'
	case __obf_cbe674ab72969f9d >= 'A' && __obf_cbe674ab72969f9d <= 'Z':
		// A-Z encoded as 10-35.
		return __obf_cbe674ab72969f9d - 'A' + 10
	case __obf_cbe674ab72969f9d == ' ':
		return 36
	case __obf_cbe674ab72969f9d == '$':
		return 37
	case __obf_cbe674ab72969f9d == '%':
		return 38
	case __obf_cbe674ab72969f9d == '*':
		return 39
	case __obf_cbe674ab72969f9d == '+':
		return 40
	case __obf_cbe674ab72969f9d == '-':
		return 41
	case __obf_cbe674ab72969f9d == '.':
		return 42
	case __obf_cbe674ab72969f9d == '/':
		return 43
	case __obf_cbe674ab72969f9d == ':':
		return 44
	default:
		log.Panicf("encodeAlphanumericCharacter() with non alphanumeric char %v.", __obf_d3308b8b9ec8593a)
	}

	return 0
}
