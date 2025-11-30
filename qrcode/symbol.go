// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_8d73621f5856b288

// symbol is a 2D array of bits representing a QR Code symbol.
//
// A symbol consists of size*size modules, with each module normally drawn as a
// black or white square. The symbol also has a border of quietZoneSize modules.
//
// A (fictional) size=2, quietZoneSize=1 QR Code looks like:
//
// +----+
// |    |
// | ab |
// | cd |
// |    |
// +----+
//
// For ease of implementation, the functions to set/get bits ignore the border,
// so (0,0)=a, (0,1)=b, (1,0)=c, and (1,1)=d. The entire symbol (including the
// border) is returned by bitmap().
type __obf_df7d9041d687f9b6 struct {
	__obf_c5516de648063f60 [// Value of module at [y][x]. True is set.
	][]bool
	__obf_d3cd552e7bb6a59e [// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	][]bool
	__obf_8ab5ee435c7948e3 int// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_3065d5392ebc294f int// Width/height of the symbol only.
	__obf_b4e60978957de1cf int// Width/height of a single quiet zone.

}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_cb5d75f8cf3fdde1(__obf_8ab5ee435c7948e3 int, __obf_b4e60978957de1cf int) *__obf_df7d9041d687f9b6 {
	var __obf_5dd86f102187ad6a __obf_df7d9041d687f9b6
	__obf_5dd86f102187ad6a.__obf_c5516de648063f60 = make([][]bool, __obf_8ab5ee435c7948e3+2*__obf_b4e60978957de1cf)
	__obf_5dd86f102187ad6a.__obf_d3cd552e7bb6a59e = make([][]bool, __obf_8ab5ee435c7948e3+2*__obf_b4e60978957de1cf)

	for __obf_015b1aa3060900e0 := range __obf_5dd86f102187ad6a.__obf_c5516de648063f60 {
		__obf_5dd86f102187ad6a.__obf_c5516de648063f60[__obf_015b1aa3060900e0] = make([]bool, __obf_8ab5ee435c7948e3+2*__obf_b4e60978957de1cf)
		__obf_5dd86f102187ad6a.__obf_d3cd552e7bb6a59e[__obf_015b1aa3060900e0] = make([]bool, __obf_8ab5ee435c7948e3+2*__obf_b4e60978957de1cf)
	}
	__obf_5dd86f102187ad6a.__obf_8ab5ee435c7948e3 = __obf_8ab5ee435c7948e3 + 2*__obf_b4e60978957de1cf
	__obf_5dd86f102187ad6a.__obf_3065d5392ebc294f = __obf_8ab5ee435c7948e3
	__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf = __obf_b4e60978957de1cf

	return &__obf_5dd86f102187ad6a
}

// get returns the module value at (x, y).
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f int, __obf_d70c25f3394ee124 int) (__obf_1b6e726e11cd9194 bool) {
	__obf_1b6e726e11cd9194 = __obf_5dd86f102187ad6a.__obf_c5516de648063f60[__obf_d70c25f3394ee124+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf][__obf_2db4671bf7cd078f+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_1b89945e00dc7e9a(__obf_2db4671bf7cd078f int, __obf_d70c25f3394ee124 int) bool {
	return !__obf_5dd86f102187ad6a.__obf_d3cd552e7bb6a59e[__obf_d70c25f3394ee124+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf][__obf_2db4671bf7cd078f+__obf_5dd86f102187ad6a.

	// numEmptyModules returns the number of empty modules.
	//
	// Initially numEmptyModules is symbolSize * symbolSize. After every module has
	// been set (to either true or false), the number of empty modules is zero.
	__obf_b4e60978957de1cf]
}

func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_a2c9171a634d8aa4() int {
	var __obf_3a7fbdb823fc7e90 int
	for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
		for __obf_2db4671bf7cd078f := 0; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
			if !__obf_5dd86f102187ad6a.__obf_d3cd552e7bb6a59e[__obf_d70c25f3394ee124+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf][__obf_2db4671bf7cd078f+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf] {
				__obf_3a7fbdb823fc7e90++
			}
		}
	}

	return __obf_3a7fbdb823fc7e90
}

// set sets the module at (x, y) to v.
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_d48f0bdef49b337a(__obf_2db4671bf7cd078f int, __obf_d70c25f3394ee124 int, __obf_1b6e726e11cd9194 bool) {
	__obf_5dd86f102187ad6a.__obf_c5516de648063f60[__obf_d70c25f3394ee124+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf][__obf_2db4671bf7cd078f+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf] = __obf_1b6e726e11cd9194
	__obf_5dd86f102187ad6a.__obf_d3cd552e7bb6a59e[__obf_d70c25f3394ee124+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf][__obf_2db4671bf7cd078f+__obf_5dd86f102187ad6a.__obf_b4e60978957de1cf] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_1f6dfffce05c0867(__obf_2db4671bf7cd078f int, __obf_d70c25f3394ee124 int, __obf_1b6e726e11cd9194 [][]bool) {
	for __obf_7c0eb8656083f52b, __obf_7c1397a937eac4ff := range __obf_1b6e726e11cd9194 {
		for __obf_015b1aa3060900e0, __obf_a5e034be5f59dea2 := range __obf_7c1397a937eac4ff {
			__obf_5dd86f102187ad6a.__obf_d48f0bdef49b337a(__obf_2db4671bf7cd078f+__obf_015b1aa3060900e0, __obf_d70c25f3394ee124+__obf_7c0eb8656083f52b,

			// bitmap returns the entire symbol, including the quiet zone.
			__obf_a5e034be5f59dea2)
		}
	}
}

func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_9c007cb7461851ed() [][]bool {
	__obf_c5516de648063f60 := make([][]bool, len(__obf_5dd86f102187ad6a.__obf_c5516de648063f60))

	for __obf_015b1aa3060900e0 := range __obf_5dd86f102187ad6a.__obf_c5516de648063f60 {
		__obf_c5516de648063f60[__obf_015b1aa3060900e0] = __obf_5dd86f102187ad6a.__obf_c5516de648063f60[__obf_015b1aa3060900e0][:]
	}

	return __obf_c5516de648063f60
}

// string returns a pictorial representation of the symbol, suitable for
// printing in a TTY.
// func (m *symbol) string() string {
// 	var result string

// 	for _, row := range m.module {
// 		for _, value := range row {
// 			switch value {
// 			case true:
// 				result += "  "
// 			case false:
// 				// Unicode 'FULL BLOCK' (U+2588).
// 				result += "██"
// 			}
// 		}
// 		result += "\n"
// 	}

// 	return result
// }

// Constants used to weight penalty calculations. Specified by ISO/IEC
// 18004:2006.
const (
	__obf_e2efb816b79db25e = 3
	__obf_d56b4d1ead8df742 = 3
	__obf_35ecd0f65f3656f8 = 40
	__obf_b91af3feea0ec8ee = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_5ff464158ac6e38e() int {
	return __obf_5dd86f102187ad6a.__obf_6998b552a3412183() + __obf_5dd86f102187ad6a.__obf_59c4c064efabf3ec() + __obf_5dd86f102187ad6a.__obf_646b48056c8eb2a9() + __obf_5dd86f102187ad6a.

	// penalty1 returns the penalty score for "adjacent modules in row/column with
	// same colour".
	//
	// The numbers of adjacent matching modules and scores are:
	// 0-5: score = 0
	// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
	__obf_4e05f9e705cba759()
}

func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_6998b552a3412183() int {
	__obf_cca1d7447027f3ba := 0

	for __obf_2db4671bf7cd078f := 0; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
		__obf_a3a80e652efddd8d := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, 0)
		__obf_3a7fbdb823fc7e90 := 1

		for __obf_d70c25f3394ee124 := 1; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
			__obf_1b6e726e11cd9194 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124)

			if __obf_1b6e726e11cd9194 != __obf_a3a80e652efddd8d {
				__obf_3a7fbdb823fc7e90 = 1
				__obf_a3a80e652efddd8d = __obf_1b6e726e11cd9194
			} else {
				__obf_3a7fbdb823fc7e90++
				if __obf_3a7fbdb823fc7e90 == 6 {
					__obf_cca1d7447027f3ba += __obf_e2efb816b79db25e + 1
				} else if __obf_3a7fbdb823fc7e90 > 6 {
					__obf_cca1d7447027f3ba++
				}
			}
		}
	}

	for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
		__obf_a3a80e652efddd8d := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(0, __obf_d70c25f3394ee124)
		__obf_3a7fbdb823fc7e90 := 1

		for __obf_2db4671bf7cd078f := 1; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
			__obf_1b6e726e11cd9194 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124)

			if __obf_1b6e726e11cd9194 != __obf_a3a80e652efddd8d {
				__obf_3a7fbdb823fc7e90 = 1
				__obf_a3a80e652efddd8d = __obf_1b6e726e11cd9194
			} else {
				__obf_3a7fbdb823fc7e90++
				if __obf_3a7fbdb823fc7e90 == 6 {
					__obf_cca1d7447027f3ba += __obf_e2efb816b79db25e + 1
				} else if __obf_3a7fbdb823fc7e90 > 6 {
					__obf_cca1d7447027f3ba++
				}
			}
		}
	}

	return __obf_cca1d7447027f3ba
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_59c4c064efabf3ec() int {
	__obf_cca1d7447027f3ba := 0

	for __obf_d70c25f3394ee124 := 1; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
		for __obf_2db4671bf7cd078f := 1; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
			__obf_05a66d7695ae5b84 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f-1, __obf_d70c25f3394ee124-1)
			__obf_ec1fa2260713c3d0 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124-1)
			__obf_5f06ffc44f58f97f := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f-1, __obf_d70c25f3394ee124)
			__obf_ac0beacb8695e028 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124)

			if __obf_ac0beacb8695e028 == __obf_5f06ffc44f58f97f && __obf_ac0beacb8695e028 == __obf_ec1fa2260713c3d0 && __obf_ac0beacb8695e028 == __obf_05a66d7695ae5b84 {
				__obf_cca1d7447027f3ba++
			}
		}
	}

	return __obf_cca1d7447027f3ba * __obf_d56b4d1ead8df742
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_646b48056c8eb2a9() int {
	__obf_cca1d7447027f3ba := 0

	for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
		var __obf_6bab3798868e9560 int16 = 0x00

		for __obf_2db4671bf7cd078f := 0; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
			__obf_6bab3798868e9560 <<= 1
			if __obf_1b6e726e11cd9194 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124); __obf_1b6e726e11cd9194 {
				__obf_6bab3798868e9560 |= 1
			}

			switch __obf_6bab3798868e9560 & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_cca1d7447027f3ba += __obf_35ecd0f65f3656f8
				__obf_6bab3798868e9560 = 0xFF
			default:
				if __obf_2db4671bf7cd078f == __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f-1 && (__obf_6bab3798868e9560&0x7f) == 0x5d {
					__obf_cca1d7447027f3ba += __obf_35ecd0f65f3656f8
					__obf_6bab3798868e9560 = 0xFF
				}
			}
		}
	}

	for __obf_2db4671bf7cd078f := 0; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
		var __obf_6bab3798868e9560 int16 = 0x00

		for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
			__obf_6bab3798868e9560 <<= 1
			if __obf_1b6e726e11cd9194 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124); __obf_1b6e726e11cd9194 {
				__obf_6bab3798868e9560 |= 1
			}

			switch __obf_6bab3798868e9560 & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_cca1d7447027f3ba += __obf_35ecd0f65f3656f8
				__obf_6bab3798868e9560 = 0xFF
			default:
				if __obf_d70c25f3394ee124 == __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f-1 && (__obf_6bab3798868e9560&0x7f) == 0x5d {
					__obf_cca1d7447027f3ba += __obf_35ecd0f65f3656f8
					__obf_6bab3798868e9560 = 0xFF
				}
			}
		}
	}

	return __obf_cca1d7447027f3ba
}

// penalty4 returns the penalty score...
func (__obf_5dd86f102187ad6a *__obf_df7d9041d687f9b6,) __obf_4e05f9e705cba759() int {
	__obf_82d3c0a9f810ea7f := __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f * __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f
	__obf_acef75697626332d := 0

	for __obf_2db4671bf7cd078f := 0; __obf_2db4671bf7cd078f < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_2db4671bf7cd078f++ {
		for __obf_d70c25f3394ee124 := 0; __obf_d70c25f3394ee124 < __obf_5dd86f102187ad6a.__obf_3065d5392ebc294f; __obf_d70c25f3394ee124++ {
			if __obf_1b6e726e11cd9194 := __obf_5dd86f102187ad6a.__obf_9e72f66032a8bc0e(__obf_2db4671bf7cd078f, __obf_d70c25f3394ee124); __obf_1b6e726e11cd9194 {
				__obf_acef75697626332d++
			}
		}
	}
	__obf_bc04745c8456858c := __obf_82d3c0a9f810ea7f/2 - __obf_acef75697626332d
	if __obf_bc04745c8456858c < 0 {
		__obf_bc04745c8456858c *= -1
	}

	return __obf_b91af3feea0ec8ee * (__obf_bc04745c8456858c / (__obf_82d3c0a9f810ea7f / 20))
}
