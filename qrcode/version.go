// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_540b74a9d13704fa

import (
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// Error detection/recovery capacity.
//
// There are several levels of error detection/recovery capacity. Higher levels
// of error recovery are able to correct more errors, with the trade-off of
// increased symbol size.
type RecoveryLevel int

const (
	// Level L: 7% error recovery.
	Low RecoveryLevel = iota

	// Level M: 15% error recovery. Good default choice.
	Medium

	// Level Q: 25% error recovery.
	High

	// Level H: 30% error recovery.
	Highest
)

// qrCodeVersion describes the data length and encoding order of a single QR
// Code version. There are 40 versions numbers x 4 recovery levels == 160
// possible qrCodeVersion structures.
type __obf_52ef8fe7f69f3df7 struct {
	__obf_7bf1f2404ecb0c0a int           // Version number (1-40 inclusive).
	__obf_a221f051d10eb503 RecoveryLevel // Error recovery level.
	__obf_ede9586c20e29cd6 __obf_ede9586c20e29cd6

	// Encoded data can be split into multiple blocks. Each block contains data
	// and error recovery bytes.
	//
	// Larger QR Codes contain more blocks.
	__obf_ef7d737cbd163114 []__obf_ef7d737cbd163114

	// Number of bits required to pad the combined data & error correction bit
	// stream up to the symbol's full capacity.
	__obf_58489f07bc720472 int
}

type __obf_ef7d737cbd163114 struct {
	__obf_bbfd4dbabec7ff0b int
	__obf_b9de796f50df339e int // Total codewords (numCodewords == numErrorCodewords+numDataCodewords).
	__obf_ebdbba63d7655685 int // Number of data codewords.

}

var (
	__obf_60777f0e04dfdeb9 = []__obf_52ef8fe7f69f3df7{
		{
			1,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					26,
					19,
				},
			},
			0,
		},
		{
			1,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					26,
					16,
				},
			},
			0,
		},
		{
			1,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					26,
					13,
				},
			},
			0,
		},
		{
			1,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					26,
					9,
				},
			},
			0,
		},
		{
			2,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					44,
					34,
				},
			},
			7,
		},
		{
			2,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					44,
					28,
				},
			},
			7,
		},
		{
			2,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					44,
					22,
				},
			},
			7,
		},
		{
			2,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					44,
					16,
				},
			},
			7,
		},
		{
			3,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					70,
					55,
				},
			},
			7,
		},
		{
			3,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					70,
					44,
				},
			},
			7,
		},
		{
			3,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					35,
					17,
				},
			},
			7,
		},
		{
			3,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					35,
					13,
				},
			},
			7,
		},
		{
			4,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					100,
					80,
				},
			},
			7,
		},
		{
			4,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					50,
					32,
				},
			},
			7,
		},
		{
			4,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					50,
					24,
				},
			},
			7,
		},
		{
			4,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					25,
					9,
				},
			},
			7,
		},
		{
			5,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					1,
					134,
					108,
				},
			},
			7,
		},
		{
			5,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					67,
					43,
				},
			},
			7,
		},
		{
			5,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					33,
					15,
				},
				{
					2,
					34,
					16,
				},
			},
			7,
		},
		{
			5,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					33,
					11,
				},
				{
					2,
					34,
					12,
				},
			},
			7,
		},
		{
			6,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					86,
					68,
				},
			},
			7,
		},
		{
			6,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					43,
					27,
				},
			},
			7,
		},
		{
			6,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					43,
					19,
				},
			},
			7,
		},
		{
			6,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					43,
					15,
				},
			},
			7,
		},
		{
			7,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					98,
					78,
				},
			},
			0,
		},
		{
			7,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					49,
					31,
				},
			},
			0,
		},
		{
			7,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					32,
					14,
				},
				{
					4,
					33,
					15,
				},
			},
			0,
		},
		{
			7,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					39,
					13,
				},
				{
					1,
					40,
					14,
				},
			},
			0,
		},
		{
			8,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					121,
					97,
				},
			},
			0,
		},
		{
			8,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					60,
					38,
				},
				{
					2,
					61,
					39,
				},
			},
			0,
		},
		{
			8,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					40,
					18,
				},
				{
					2,
					41,
					19,
				},
			},
			0,
		},
		{
			8,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					40,
					14,
				},
				{
					2,
					41,
					15,
				},
			},
			0,
		},
		{
			9,
			Low, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					2,
					146,
					116,
				},
			},
			0,
		},
		{
			9,
			Medium, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					3,
					58,
					36,
				},
				{
					2,
					59,
					37,
				},
			},
			0,
		},
		{
			9,
			High, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					36,
					16,
				},
				{
					4,
					37,
					17,
				},
			},
			0,
		},
		{
			9,
			Highest, __obf_36e336d7eab6b5c8, []__obf_ef7d737cbd163114{
				{
					4,
					36,
					12,
				},
				{
					4,
					37,
					13,
				},
			},
			0,
		},
		{
			10,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					2,
					86,
					68,
				},
				{
					2,
					87,
					69,
				},
			},
			0,
		},
		{
			10,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					69,
					43,
				},
				{
					1,
					70,
					44,
				},
			},
			0,
		},
		{
			10,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					6,
					43,
					19,
				},
				{
					2,
					44,
					20,
				},
			},
			0,
		},
		{
			10,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					6,
					43,
					15,
				},
				{
					2,
					44,
					16,
				},
			},
			0,
		},
		{
			11,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					101,
					81,
				},
			},
			0,
		},
		{
			11,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					1,
					80,
					50,
				},
				{
					4,
					81,
					51,
				},
			},
			0,
		},
		{
			11,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					50,
					22,
				},
				{
					4,
					51,
					23,
				},
			},
			0,
		},
		{
			11,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					36,
					12,
				},
				{
					8,
					37,
					13,
				},
			},
			0,
		},
		{
			12,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					2,
					116,
					92,
				},
				{
					2,
					117,
					93,
				},
			},
			0,
		},
		{
			12,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					6,
					58,
					36,
				},
				{
					2,
					59,
					37,
				},
			},
			0,
		},
		{
			12,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					46,
					20,
				},
				{
					6,
					47,
					21,
				},
			},
			0,
		},
		{
			12,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					7,
					42,
					14,
				},
				{
					4,
					43,
					15,
				},
			},
			0,
		},
		{
			13,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					133,
					107,
				},
			},
			0,
		},
		{
			13,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					8,
					59,
					37,
				},
				{
					1,
					60,
					38,
				},
			},
			0,
		},
		{
			13,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					8,
					44,
					20,
				},
				{
					4,
					45,
					21,
				},
			},
			0,
		},
		{
			13,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					12,
					33,
					11,
				},
				{
					4,
					34,
					12,
				},
			},
			0,
		},
		{
			14,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					145,
					115,
				},
				{
					1,
					146,
					116,
				},
			},
			3,
		},
		{
			14,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					64,
					40,
				},
				{
					5,
					65,
					41,
				},
			},
			3,
		},
		{
			14,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					11,
					36,
					16,
				},
				{
					5,
					37,
					17,
				},
			},
			3,
		},
		{
			14,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					11,
					36,
					12,
				},
				{
					5,
					37,
					13,
				},
			},
			3,
		},
		{
			15,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					5,
					109,
					87,
				},
				{
					1,
					110,
					88,
				},
			},
			3,
		},
		{
			15,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					5,
					65,
					41,
				},
				{
					5,
					66,
					42,
				},
			},
			3,
		},
		{
			15,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					5,
					54,
					24,
				},
				{
					7,
					55,
					25,
				},
			},
			3,
		},
		{
			15,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					11,
					36,
					12,
				},
				{
					7,
					37,
					13,
				},
			},
			3,
		},
		{
			16,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					5,
					122,
					98,
				},
				{
					1,
					123,
					99,
				},
			},
			3,
		},
		{
			16,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					7,
					73,
					45,
				},
				{
					3,
					74,
					46,
				},
			},
			3,
		},
		{
			16,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					15,
					43,
					19,
				},
				{
					2,
					44,
					20,
				},
			},
			3,
		},
		{
			16,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					45,
					15,
				},
				{
					13,
					46,
					16,
				},
			},
			3,
		},
		{
			17,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					1,
					135,
					107,
				},
				{
					5,
					136,
					108,
				},
			},
			3,
		},
		{
			17,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					10,
					74,
					46,
				},
				{
					1,
					75,
					47,
				},
			},
			3,
		},
		{
			17,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					1,
					50,
					22,
				},
				{
					15,
					51,
					23,
				},
			},
			3,
		},
		{
			17,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					2,
					42,
					14,
				},
				{
					17,
					43,
					15,
				},
			},
			3,
		},
		{
			18,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					5,
					150,
					120,
				},
				{
					1,
					151,
					121,
				},
			},
			3,
		},
		{
			18,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					9,
					69,
					43,
				},
				{
					4,
					70,
					44,
				},
			},
			3,
		},
		{
			18,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					17,
					50,
					22,
				},
				{
					1,
					51,
					23,
				},
			},
			3,
		},
		{
			18,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					2,
					42,
					14,
				},
				{
					19,
					43,
					15,
				},
			},
			3,
		},
		{
			19,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					141,
					113,
				},
				{
					4,
					142,
					114,
				},
			},
			3,
		},
		{
			19,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					70,
					44,
				},
				{
					11,
					71,
					45,
				},
			},
			3,
		},
		{
			19,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					17,
					47,
					21,
				},
				{
					4,
					48,
					22,
				},
			},
			3,
		},
		{
			19,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					9,
					39,
					13,
				},
				{
					16,
					40,
					14,
				},
			},
			3,
		},
		{
			20,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					135,
					107,
				},
				{
					5,
					136,
					108,
				},
			},
			3,
		},
		{
			20,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					3,
					67,
					41,
				},
				{
					13,
					68,
					42,
				},
			},
			3,
		},
		{
			20,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					15,
					54,
					24,
				},
				{
					5,
					55,
					25,
				},
			},
			3,
		},
		{
			20,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					15,
					43,
					15,
				},
				{
					10,
					44,
					16,
				},
			},
			3,
		},
		{
			21,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					144,
					116,
				},
				{
					4,
					145,
					117,
				},
			},
			4,
		},
		{
			21,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					17,
					68,
					42,
				},
			},
			4,
		},
		{
			21,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					17,
					50,
					22,
				},
				{
					6,
					51,
					23,
				},
			},
			4,
		},
		{
			21,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					19,
					46,
					16,
				},
				{
					6,
					47,
					17,
				},
			},
			4,
		},
		{
			22,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					2,
					139,
					111,
				},
				{
					7,
					140,
					112,
				},
			},
			4,
		},
		{
			22,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					17,
					74,
					46,
				},
			},
			4,
		},
		{
			22,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					7,
					54,
					24,
				},
				{
					16,
					55,
					25,
				},
			},
			4,
		},
		{
			22,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					34,
					37,
					13,
				},
			},
			4,
		},
		{
			23,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					151,
					121,
				},
				{
					5,
					152,
					122,
				},
			},
			4,
		},
		{
			23,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					4,
					75,
					47,
				},
				{
					14,
					76,
					48,
				},
			},
			4,
		},
		{
			23,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					11,
					54,
					24,
				},
				{
					14,
					55,
					25,
				},
			},
			4,
		},
		{
			23,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					16,
					45,
					15,
				},
				{
					14,
					46,
					16,
				},
			},
			4,
		},
		{
			24,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					6,
					147,
					117,
				},
				{
					4,
					148,
					118,
				},
			},
			4,
		},
		{
			24,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					6,
					73,
					45,
				},
				{
					14,
					74,
					46,
				},
			},
			4,
		},
		{
			24,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					11,
					54,
					24,
				},
				{
					16,
					55,
					25,
				},
			},
			4,
		},
		{
			24,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					30,
					46,
					16,
				},
				{
					2,
					47,
					17,
				},
			},
			4,
		},
		{
			25,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					8,
					132,
					106,
				},
				{
					4,
					133,
					107,
				},
			},
			4,
		},
		{
			25,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					8,
					75,
					47,
				},
				{
					13,
					76,
					48,
				},
			},
			4,
		},
		{
			25,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					7,
					54,
					24,
				},
				{
					22,
					55,
					25,
				},
			},
			4,
		},
		{
			25,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					22,
					45,
					15,
				},
				{
					13,
					46,
					16,
				},
			},
			4,
		},
		{
			26,
			Low, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					10,
					142,
					114,
				},
				{
					2,
					143,
					115,
				},
			},
			4,
		},
		{
			26,
			Medium, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					19,
					74,
					46,
				},
				{
					4,
					75,
					47,
				},
			},
			4,
		},
		{
			26,
			High, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					28,
					50,
					22,
				},
				{
					6,
					51,
					23,
				},
			},
			4,
		},
		{
			26,
			Highest, __obf_ce4133b003340f88, []__obf_ef7d737cbd163114{
				{
					33,
					46,
					16,
				},
				{
					4,
					47,
					17,
				},
			},
			4,
		},
		{
			27,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					8,
					152,
					122,
				},
				{
					4,
					153,
					123,
				},
			},
			4,
		},
		{
			27,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					22,
					73,
					45,
				},
				{
					3,
					74,
					46,
				},
			},
			4,
		},
		{
			27,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					8,
					53,
					23,
				},
				{
					26,
					54,
					24,
				},
			},
			4,
		},
		{
			27,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					12,
					45,
					15,
				},
				{
					28,
					46,
					16,
				},
			},
			4,
		},
		{
			28,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					3,
					147,
					117,
				},
				{
					10,
					148,
					118,
				},
			},
			3,
		},
		{
			28,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					3,
					73,
					45,
				},
				{
					23,
					74,
					46,
				},
			},
			3,
		},
		{
			28,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					4,
					54,
					24,
				},
				{
					31,
					55,
					25,
				},
			},
			3,
		},
		{
			28,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					11,
					45,
					15,
				},
				{
					31,
					46,
					16,
				},
			},
			3,
		},
		{
			29,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					7,
					146,
					116,
				},
				{
					7,
					147,
					117,
				},
			},
			3,
		},
		{
			29,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					21,
					73,
					45,
				},
				{
					7,
					74,
					46,
				},
			},
			3,
		},
		{
			29,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					1,
					53,
					23,
				},
				{
					37,
					54,
					24,
				},
			},
			3,
		},
		{
			29,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					19,
					45,
					15,
				},
				{
					26,
					46,
					16,
				},
			},
			3,
		},
		{
			30,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					5,
					145,
					115,
				},
				{
					10,
					146,
					116,
				},
			},
			3,
		},
		{
			30,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					19,
					75,
					47,
				},
				{
					10,
					76,
					48,
				},
			},
			3,
		},
		{
			30,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					15,
					54,
					24,
				},
				{
					25,
					55,
					25,
				},
			},
			3,
		},
		{
			30,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					23,
					45,
					15,
				},
				{
					25,
					46,
					16,
				},
			},
			3,
		},
		{
			31,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					13,
					145,
					115,
				},
				{
					3,
					146,
					116,
				},
			},
			3,
		},
		{
			31,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					2,
					74,
					46,
				},
				{
					29,
					75,
					47,
				},
			},
			3,
		},
		{
			31,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					42,
					54,
					24,
				},
				{
					1,
					55,
					25,
				},
			},
			3,
		},
		{
			31,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					23,
					45,
					15,
				},
				{
					28,
					46,
					16,
				},
			},
			3,
		},
		{
			32,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					17,
					145,
					115,
				},
			},
			3,
		},
		{
			32,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					10,
					74,
					46,
				},
				{
					23,
					75,
					47,
				},
			},
			3,
		},
		{
			32,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					10,
					54,
					24,
				},
				{
					35,
					55,
					25,
				},
			},
			3,
		},
		{
			32,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					19,
					45,
					15,
				},
				{
					35,
					46,
					16,
				},
			},
			3,
		},
		{
			33,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					17,
					145,
					115,
				},
				{
					1,
					146,
					116,
				},
			},
			3,
		},
		{
			33,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					14,
					74,
					46,
				},
				{
					21,
					75,
					47,
				},
			},
			3,
		},
		{
			33,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					29,
					54,
					24,
				},
				{
					19,
					55,
					25,
				},
			},
			3,
		},
		{
			33,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					11,
					45,
					15,
				},
				{
					46,
					46,
					16,
				},
			},
			3,
		},
		{
			34,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					13,
					145,
					115,
				},
				{
					6,
					146,
					116,
				},
			},
			3,
		},
		{
			34,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					14,
					74,
					46,
				},
				{
					23,
					75,
					47,
				},
			},
			3,
		},
		{
			34,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					44,
					54,
					24,
				},
				{
					7,
					55,
					25,
				},
			},
			3,
		},
		{
			34,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					59,
					46,
					16,
				},
				{
					1,
					47,
					17,
				},
			},
			3,
		},
		{
			35,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					12,
					151,
					121,
				},
				{
					7,
					152,
					122,
				},
			},
			0,
		},
		{
			35,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					12,
					75,
					47,
				},
				{
					26,
					76,
					48,
				},
			},
			0,
		},
		{
			35,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					39,
					54,
					24,
				},
				{
					14,
					55,
					25,
				},
			},
			0,
		},
		{
			35,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					22,
					45,
					15,
				},
				{
					41,
					46,
					16,
				},
			},
			0,
		},
		{
			36,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					6,
					151,
					121,
				},
				{
					14,
					152,
					122,
				},
			},
			0,
		},
		{
			36,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					6,
					75,
					47,
				},
				{
					34,
					76,
					48,
				},
			},
			0,
		},
		{
			36,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					46,
					54,
					24,
				},
				{
					10,
					55,
					25,
				},
			},
			0,
		},
		{
			36,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					2,
					45,
					15,
				},
				{
					64,
					46,
					16,
				},
			},
			0,
		},
		{
			37,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					17,
					152,
					122,
				},
				{
					4,
					153,
					123,
				},
			},
			0,
		},
		{
			37,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					29,
					74,
					46,
				},
				{
					14,
					75,
					47,
				},
			},
			0,
		},
		{
			37,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					49,
					54,
					24,
				},
				{
					10,
					55,
					25,
				},
			},
			0,
		},
		{
			37,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					24,
					45,
					15,
				},
				{
					46,
					46,
					16,
				},
			},
			0,
		},
		{
			38,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					4,
					152,
					122,
				},
				{
					18,
					153,
					123,
				},
			},
			0,
		},
		{
			38,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					13,
					74,
					46,
				},
				{
					32,
					75,
					47,
				},
			},
			0,
		},
		{
			38,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					48,
					54,
					24,
				},
				{
					14,
					55,
					25,
				},
			},
			0,
		},
		{
			38,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					42,
					45,
					15,
				},
				{
					32,
					46,
					16,
				},
			},
			0,
		},
		{
			39,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					20,
					147,
					117,
				},
				{
					4,
					148,
					118,
				},
			},
			0,
		},
		{
			39,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					40,
					75,
					47,
				},
				{
					7,
					76,
					48,
				},
			},
			0,
		},
		{
			39,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					43,
					54,
					24,
				},
				{
					22,
					55,
					25,
				},
			},
			0,
		},
		{
			39,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					10,
					45,
					15,
				},
				{
					67,
					46,
					16,
				},
			},
			0,
		},
		{
			40,
			Low, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					19,
					148,
					118,
				},
				{
					6,
					149,
					119,
				},
			},
			0,
		},
		{
			40,
			Medium, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					18,
					75,
					47,
				},
				{
					31,
					76,
					48,
				},
			},
			0,
		},
		{
			40,
			High, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					34,
					54,
					24,
				},
				{
					34,
					55,
					25,
				},
			},
			0,
		},
		{
			40,
			Highest, __obf_d9027dc4f9acd3b0, []__obf_ef7d737cbd163114{
				{
					20,
					45,
					15,
				},
				{
					61,
					46,
					16,
				},
			},
			0,
		},
	}
)

var (
	__obf_8757cd18a70e2cec =
	// Each QR Code contains a 15-bit Format Information value.  The 15 bits
	// consist of 5 data bits concatenated with 10 error correction bits.
	//
	// The 5 data bits consist of:
	// - 2 bits for the error correction level (L=01, M=00, G=11, H=10).
	// - 3 bits for the data mask pattern identifier.
	//
	// formatBitSequence is a mapping from the 5 data bits to the completed 15-bit
	// Format Information value.
	//
	// For example, a QR Code using error correction level L, and data mask
	// pattern identifier 001:
	//
	// 01 | 001 = 01001 = 0x9
	// formatBitSequence[0x9].qrCode = 0x72f3 = 111001011110011
	[]struct {
		__obf_58fd024ca3fb4778 uint32
		__obf_5c09276c9ed39348 uint32
	}{
		{0x5412, 0x4445},
		{0x5125, 0x4172},
		{0x5e7c, 0x4e2b},
		{0x5b4b, 0x4b1c},
		{0x45f9, 0x55ae},
		{0x40ce, 0x5099},
		{0x4f97, 0x5fc0},
		{0x4aa0, 0x5af7},
		{0x77c4, 0x6793},
		{0x72f3, 0x62a4},
		{0x7daa, 0x6dfd},
		{0x789d, 0x68ca},
		{0x662f, 0x7678},
		{0x6318, 0x734f},
		{0x6c41, 0x7c16},
		{0x6976, 0x7921},
		{0x1689, 0x06de},
		{0x13be, 0x03e9},
		{0x1ce7, 0x0cb0},
		{0x19d0, 0x0987},
		{0x0762, 0x1735},
		{0x0255, 0x1202},
		{0x0d0c, 0x1d5b},
		{0x083b, 0x186c},
		{0x355f, 0x2508},
		{0x3068, 0x203f},
		{0x3f31, 0x2f66},
		{0x3a06, 0x2a51},
		{0x24b4, 0x34e3},
		{0x2183, 0x31d4},
		{0x2eda, 0x3e8d},
		{0x2bed, 0x3bba},
	}
	__obf_d1c726c6984fa48d =

	// QR Codes version 7 and higher contain an 18-bit Version Information value,
	// consisting of a 6 data bits and 12 error correction bits.
	//
	// versionBitSequence is a mapping from QR Code version to the completed
	// 18-bit Version Information value.
	//
	// For example, a QR code of version 7:
	// versionBitSequence[0x7] = 0x07c94 = 000111110010010100
	[]uint32{
		0x00000,
		0x00000,
		0x00000,
		0x00000,
		0x00000,
		0x00000,
		0x00000,
		0x07c94,
		0x085bc,
		0x09a99,
		0x0a4d3,
		0x0bbf6,
		0x0c762,
		0x0d847,
		0x0e60d,
		0x0f928,
		0x10b78,
		0x1145d,
		0x12a17,
		0x13532,
		0x149a6,
		0x15683,
		0x168c9,
		0x177ec,
		0x18ec4,
		0x191e1,
		0x1afab,
		0x1b08e,
		0x1cc1a,
		0x1d33f,
		0x1ed75,
		0x1f250,
		0x209d5,
		0x216f0,
		0x228ba,
		0x2379f,
		0x24b0b,
		0x2542e,
		0x26a64,
		0x27541,
		0x28c69,
	}
)

const (
	__obf_b1e8e072b4d898e1 = 15
	__obf_1b80056f638b8cfd = 18
)

// formatInfo returns the 15-bit Format Information value for a QR
// code.
func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_fd6a56f00c2dba37(__obf_25e0d8307970288e int) *bitset.Bitset {
	__obf_fa551e66a1ea9884 := 0

	switch __obf_75eba50799fccf42.__obf_a221f051d10eb503 {
	case Low:
		__obf_fa551e66a1ea9884 = 0x08 // 0b01000
	case Medium:
		__obf_fa551e66a1ea9884 = 0x00 // 0b00000
	case High:
		__obf_fa551e66a1ea9884 = 0x18 // 0b11000
	case Highest:
		__obf_fa551e66a1ea9884 = 0x10 // 0b10000
	default:
		log.Panicf("Invalid level %d", __obf_75eba50799fccf42.__obf_a221f051d10eb503)
	}

	if __obf_25e0d8307970288e < 0 || __obf_25e0d8307970288e > 7 {
		log.Panicf("Invalid maskPattern %d", __obf_25e0d8307970288e)
	}
	__obf_fa551e66a1ea9884 |= __obf_25e0d8307970288e & 0x7
	__obf_3ceb8341caf311f5 := bitset.New()
	__obf_3ceb8341caf311f5.
		AppendUint32(__obf_8757cd18a70e2cec[__obf_fa551e66a1ea9884].__obf_58fd024ca3fb4778, __obf_b1e8e072b4d898e1)

	return __obf_3ceb8341caf311f5
}

// versionInfo returns the 18-bit Version Information value for a QR Code.
//
// Version Information is applicable only to QR Codes versions 7-40 inclusive.
// nil is returned if Version Information is not required.
func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_10c4cf895310ca83() *bitset.Bitset {
	if __obf_75eba50799fccf42.__obf_7bf1f2404ecb0c0a < 7 {
		return nil
	}
	__obf_3ceb8341caf311f5 := bitset.New()
	__obf_3ceb8341caf311f5.
		AppendUint32(__obf_d1c726c6984fa48d[__obf_75eba50799fccf42.__obf_7bf1f2404ecb0c0a], 18)

	return __obf_3ceb8341caf311f5
}

// numDataBits returns the data capacity in bits.
func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_6ca7a7ab38ace36f() int {
	__obf_6ca7a7ab38ace36f := 0
	for _, __obf_aa760948f0d0122d := range __obf_75eba50799fccf42.__obf_ef7d737cbd163114 {
		__obf_6ca7a7ab38ace36f += 8 * __obf_aa760948f0d0122d.__obf_bbfd4dbabec7ff0b * __obf_aa760948f0d0122d. // 8 bits in a byte
															__obf_ebdbba63d7655685
	}

	return __obf_6ca7a7ab38ace36f
}

// chooseQRCodeVersion chooses the most suitable QR Code version for a stated
// data length in bits, the error recovery level required, and the data encoder
// used.
//
// The chosen QR Code version is the smallest version able to fit numDataBits
// and the optional terminator bits required by the specified encoder.
//
// On success the chosen QR Code version is returned.
func __obf_846df71d822a28c1(__obf_a221f051d10eb503 RecoveryLevel, __obf_2646c2f27fac28d6 *__obf_5aac2ee6cac292a2, __obf_6ca7a7ab38ace36f int) *__obf_52ef8fe7f69f3df7 {
	var __obf_2d8435b35e31964c *__obf_52ef8fe7f69f3df7

	for _, __obf_75eba50799fccf42 := range __obf_60777f0e04dfdeb9 {
		if __obf_75eba50799fccf42.__obf_a221f051d10eb503 != __obf_a221f051d10eb503 {
			continue
		} else if __obf_75eba50799fccf42.__obf_7bf1f2404ecb0c0a < __obf_2646c2f27fac28d6.__obf_84c65bc29348e4af {
			continue
		} else if __obf_75eba50799fccf42.__obf_7bf1f2404ecb0c0a > __obf_2646c2f27fac28d6.__obf_69cd1e2223127914 {
			break
		}
		__obf_88648f9cdefc270d := __obf_75eba50799fccf42.__obf_6ca7a7ab38ace36f() - __obf_6ca7a7ab38ace36f

		if __obf_88648f9cdefc270d >= 0 {
			__obf_2d8435b35e31964c = &__obf_75eba50799fccf42
			break
		}
	}

	return __obf_2d8435b35e31964c
}

func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_bdf7c8fa55e0922e(__obf_6ca7a7ab38ace36f int) int {
	__obf_88648f9cdefc270d := __obf_75eba50799fccf42.__obf_6ca7a7ab38ace36f() - __obf_6ca7a7ab38ace36f

	var __obf_14757c29e2bdea6e int

	switch {
	case __obf_88648f9cdefc270d >= 4:
		__obf_14757c29e2bdea6e = 4
	default:
		__obf_14757c29e2bdea6e = __obf_88648f9cdefc270d
	}

	return __obf_14757c29e2bdea6e
}

// numBlocks returns the number of blocks.
func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_bbfd4dbabec7ff0b() int {
	__obf_bbfd4dbabec7ff0b := 0

	for _, __obf_aa760948f0d0122d := range __obf_75eba50799fccf42.__obf_ef7d737cbd163114 {
		__obf_bbfd4dbabec7ff0b += __obf_aa760948f0d0122d.__obf_bbfd4dbabec7ff0b
	}

	return __obf_bbfd4dbabec7ff0b
}

// numBitsToPadToCodeword returns the number of bits required to pad data of
// length numDataBits upto the nearest codeword size.
func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_187f84545b0e27f4(__obf_6ca7a7ab38ace36f int) int {
	if __obf_6ca7a7ab38ace36f == __obf_75eba50799fccf42.__obf_6ca7a7ab38ace36f() {
		return 0
	}

	return (8 - __obf_6ca7a7ab38ace36f%8) % 8
}

// symbolSize returns the size of the QR Code symbol in number of modules (which
// is both the width and height, since QR codes are square). The QR Code has
// size symbolSize() x symbolSize() pixels. This does not include the quiet
// zone.
func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_7bb0ccb8474ed4db() int {
	return 21 + (__obf_75eba50799fccf42.

		// quietZoneSize returns the number of pixels of border space on each side of
		// the QR Code. The quiet space assists with decoding.
		__obf_7bf1f2404ecb0c0a-1)*4
}

func (__obf_75eba50799fccf42 __obf_52ef8fe7f69f3df7) __obf_914e29178d61e4a9() int {
	return 4
}

// getQRCodeVersion returns the QR Code version by version number and recovery
// level. Returns nil if the requested combination is not defined.
func __obf_23004ab726110f63(__obf_a221f051d10eb503 RecoveryLevel, __obf_7bf1f2404ecb0c0a int) *__obf_52ef8fe7f69f3df7 {
	for _, __obf_75eba50799fccf42 := range __obf_60777f0e04dfdeb9 {
		if __obf_75eba50799fccf42.__obf_a221f051d10eb503 == __obf_a221f051d10eb503 && __obf_75eba50799fccf42.__obf_7bf1f2404ecb0c0a == __obf_7bf1f2404ecb0c0a {
			return &__obf_75eba50799fccf42
		}
	}

	return nil
}
