// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_380fc15a74e6e942

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
type __obf_870ce853fd3b5622 struct {
	__obf_264d1de845406472 int           // Version number (1-40 inclusive).
	__obf_fc3c629b710f5957 RecoveryLevel // Error recovery level.
	__obf_d05a9d69f4a51132 __obf_d05a9d69f4a51132

	// Encoded data can be split into multiple blocks. Each block contains data
	// and error recovery bytes.
	//
	// Larger QR Codes contain more blocks.
	__obf_d754684ce53f7200 []__obf_d754684ce53f7200

	// Number of bits required to pad the combined data & error correction bit
	// stream up to the symbol's full capacity.
	__obf_059895faaf46eda6 int
}

type __obf_d754684ce53f7200 struct {
	__obf_3d59d9d22bc9936c int
	__obf_9be363f84174ff9f int // Total codewords (numCodewords == numErrorCodewords+numDataCodewords).
	__obf_896247f00fa0ee79 int // Number of data codewords.

}

var (
	__obf_c82737fb11ae5a50 = []__obf_870ce853fd3b5622{
		{
			1,
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Medium, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			High, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Highest, __obf_f0e68bac73c2a21c, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Medium, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			High, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Highest, __obf_b270254bf247f115, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Low, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Medium, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			High, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
			Highest, __obf_602ce8ec28abacf4, []__obf_d754684ce53f7200{
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
	__obf_9bbaa121fac53b17 =
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
		__obf_723b66302079c9e8 uint32
		__obf_db677989584f90b7 uint32
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
	__obf_591717f2fbd24754 =

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
	__obf_4193d55bd666deb1 = 15
	__obf_19ec830da4cebeb5 = 18
)

// formatInfo returns the 15-bit Format Information value for a QR
// code.
func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_cff0cd5bff199c02(__obf_9f9fce9ca5630ce5 int) *bitset.Bitset {
	__obf_0270d89b5e991bc2 := 0

	switch __obf_3591b555f7f18238.__obf_fc3c629b710f5957 {
	case Low:
		__obf_0270d89b5e991bc2 = 0x08 // 0b01000
	case Medium:
		__obf_0270d89b5e991bc2 = 0x00 // 0b00000
	case High:
		__obf_0270d89b5e991bc2 = 0x18 // 0b11000
	case Highest:
		__obf_0270d89b5e991bc2 = 0x10 // 0b10000
	default:
		log.Panicf("Invalid level %d", __obf_3591b555f7f18238.__obf_fc3c629b710f5957)
	}

	if __obf_9f9fce9ca5630ce5 < 0 || __obf_9f9fce9ca5630ce5 > 7 {
		log.Panicf("Invalid maskPattern %d", __obf_9f9fce9ca5630ce5)
	}
	__obf_0270d89b5e991bc2 |= __obf_9f9fce9ca5630ce5 & 0x7
	__obf_0f51954847d61555 := bitset.New()
	__obf_0f51954847d61555.
		AppendUint32(__obf_9bbaa121fac53b17[__obf_0270d89b5e991bc2].__obf_723b66302079c9e8, __obf_4193d55bd666deb1)

	return __obf_0f51954847d61555
}

// versionInfo returns the 18-bit Version Information value for a QR Code.
//
// Version Information is applicable only to QR Codes versions 7-40 inclusive.
// nil is returned if Version Information is not required.
func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_016ebbf815cdbbe6() *bitset.Bitset {
	if __obf_3591b555f7f18238.__obf_264d1de845406472 < 7 {
		return nil
	}
	__obf_0f51954847d61555 := bitset.New()
	__obf_0f51954847d61555.
		AppendUint32(__obf_591717f2fbd24754[__obf_3591b555f7f18238.__obf_264d1de845406472], 18)

	return __obf_0f51954847d61555
}

// numDataBits returns the data capacity in bits.
func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_35f302a87ce4a7f1() int {
	__obf_35f302a87ce4a7f1 := 0
	for _, __obf_b58cd35067d92ce7 := range __obf_3591b555f7f18238.__obf_d754684ce53f7200 {
		__obf_35f302a87ce4a7f1 += 8 * __obf_b58cd35067d92ce7.__obf_3d59d9d22bc9936c * __obf_b58cd35067d92ce7. // 8 bits in a byte
															__obf_896247f00fa0ee79
	}

	return __obf_35f302a87ce4a7f1
}

// chooseQRCodeVersion chooses the most suitable QR Code version for a stated
// data length in bits, the error recovery level required, and the data encoder
// used.
//
// The chosen QR Code version is the smallest version able to fit numDataBits
// and the optional terminator bits required by the specified encoder.
//
// On success the chosen QR Code version is returned.
func __obf_945c524e956c9051(__obf_fc3c629b710f5957 RecoveryLevel, __obf_c275e7e500011398 *__obf_8a2e265fe487f1a2, __obf_35f302a87ce4a7f1 int) *__obf_870ce853fd3b5622 {
	var __obf_67cb565eb7f8949d *__obf_870ce853fd3b5622

	for _, __obf_3591b555f7f18238 := range __obf_c82737fb11ae5a50 {
		if __obf_3591b555f7f18238.__obf_fc3c629b710f5957 != __obf_fc3c629b710f5957 {
			continue
		} else if __obf_3591b555f7f18238.__obf_264d1de845406472 < __obf_c275e7e500011398.__obf_6ef648b259fb8ea6 {
			continue
		} else if __obf_3591b555f7f18238.__obf_264d1de845406472 > __obf_c275e7e500011398.__obf_ff1062ba2e7ad054 {
			break
		}
		__obf_3236e98e05b89b04 := __obf_3591b555f7f18238.__obf_35f302a87ce4a7f1() - __obf_35f302a87ce4a7f1

		if __obf_3236e98e05b89b04 >= 0 {
			__obf_67cb565eb7f8949d = &__obf_3591b555f7f18238
			break
		}
	}

	return __obf_67cb565eb7f8949d
}

func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_7f895ab7638986af(__obf_35f302a87ce4a7f1 int) int {
	__obf_3236e98e05b89b04 := __obf_3591b555f7f18238.__obf_35f302a87ce4a7f1() - __obf_35f302a87ce4a7f1

	var __obf_52c948bfc84a0bb4 int

	switch {
	case __obf_3236e98e05b89b04 >= 4:
		__obf_52c948bfc84a0bb4 = 4
	default:
		__obf_52c948bfc84a0bb4 = __obf_3236e98e05b89b04
	}

	return __obf_52c948bfc84a0bb4
}

// numBlocks returns the number of blocks.
func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_3d59d9d22bc9936c() int {
	__obf_3d59d9d22bc9936c := 0

	for _, __obf_b58cd35067d92ce7 := range __obf_3591b555f7f18238.__obf_d754684ce53f7200 {
		__obf_3d59d9d22bc9936c += __obf_b58cd35067d92ce7.__obf_3d59d9d22bc9936c
	}

	return __obf_3d59d9d22bc9936c
}

// numBitsToPadToCodeword returns the number of bits required to pad data of
// length numDataBits upto the nearest codeword size.
func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_d5e0eb0d36587ec2(__obf_35f302a87ce4a7f1 int) int {
	if __obf_35f302a87ce4a7f1 == __obf_3591b555f7f18238.__obf_35f302a87ce4a7f1() {
		return 0
	}

	return (8 - __obf_35f302a87ce4a7f1%8) % 8
}

// symbolSize returns the size of the QR Code symbol in number of modules (which
// is both the width and height, since QR codes are square). The QR Code has
// size symbolSize() x symbolSize() pixels. This does not include the quiet
// zone.
func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_108473f4f8f75beb() int {
	return 21 + (__obf_3591b555f7f18238.

		// quietZoneSize returns the number of pixels of border space on each side of
		// the QR Code. The quiet space assists with decoding.
		__obf_264d1de845406472-1)*4
}

func (__obf_3591b555f7f18238 __obf_870ce853fd3b5622) __obf_7674347e39061c69() int {
	return 4
}

// getQRCodeVersion returns the QR Code version by version number and recovery
// level. Returns nil if the requested combination is not defined.
func __obf_c5839fc8a96e682a(__obf_fc3c629b710f5957 RecoveryLevel, __obf_264d1de845406472 int) *__obf_870ce853fd3b5622 {
	for _, __obf_3591b555f7f18238 := range __obf_c82737fb11ae5a50 {
		if __obf_3591b555f7f18238.__obf_fc3c629b710f5957 == __obf_fc3c629b710f5957 && __obf_3591b555f7f18238.__obf_264d1de845406472 == __obf_264d1de845406472 {
			return &__obf_3591b555f7f18238
		}
	}

	return nil
}
