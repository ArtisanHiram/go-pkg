// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_48b3d686b985a20b

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
type __obf_14306cbbf0004950 struct {
	__obf_8ade07d880d3e6f3 [// Value of module at [y][x]. True is set.
	][]bool
	__obf_b0362872b0bcdee4 [// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	][]bool
	__obf_311d29b1744e83e0 int// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_934aef36207d3467 int// Width/height of the symbol only.
	__obf_dcb53264a169b9b1 int// Width/height of a single quiet zone.

}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_c86c0a5bb103d5b0(__obf_311d29b1744e83e0 int, __obf_dcb53264a169b9b1 int) *__obf_14306cbbf0004950 {
	var __obf_7b9309ca6f3e8da2 __obf_14306cbbf0004950
	__obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3 = make([][]bool, __obf_311d29b1744e83e0+2*__obf_dcb53264a169b9b1)
	__obf_7b9309ca6f3e8da2.__obf_b0362872b0bcdee4 = make([][]bool, __obf_311d29b1744e83e0+2*__obf_dcb53264a169b9b1)

	for __obf_8715c27aa2bf2586 := range __obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3 {
		__obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3[__obf_8715c27aa2bf2586] = make([]bool, __obf_311d29b1744e83e0+2*__obf_dcb53264a169b9b1)
		__obf_7b9309ca6f3e8da2.__obf_b0362872b0bcdee4[__obf_8715c27aa2bf2586] = make([]bool, __obf_311d29b1744e83e0+2*__obf_dcb53264a169b9b1)
	}
	__obf_7b9309ca6f3e8da2.__obf_311d29b1744e83e0 = __obf_311d29b1744e83e0 + 2*__obf_dcb53264a169b9b1
	__obf_7b9309ca6f3e8da2.__obf_934aef36207d3467 = __obf_311d29b1744e83e0
	__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1 = __obf_dcb53264a169b9b1

	return &__obf_7b9309ca6f3e8da2
}

// get returns the module value at (x, y).
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_55397a5408a37425(__obf_cf967a75716d8446 int, __obf_8650f6fe673c98a1 int) (__obf_f85dcc6c0aea860b bool) {
	__obf_f85dcc6c0aea860b = __obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3[__obf_8650f6fe673c98a1+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1][__obf_cf967a75716d8446+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_64c163370054c2c2(__obf_cf967a75716d8446 int, __obf_8650f6fe673c98a1 int) bool {
	return !__obf_7b9309ca6f3e8da2.__obf_b0362872b0bcdee4[__obf_8650f6fe673c98a1+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1][__obf_cf967a75716d8446+__obf_7b9309ca6f3e8da2.

	// numEmptyModules returns the number of empty modules.
	//
	// Initially numEmptyModules is symbolSize * symbolSize. After every module has
	// been set (to either true or false), the number of empty modules is zero.
	__obf_dcb53264a169b9b1]
}

func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_2a27c02f01243dcc() int {
	var __obf_65b7827e8fd2993e int
	for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
		for __obf_cf967a75716d8446 := 0; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
			if !__obf_7b9309ca6f3e8da2.__obf_b0362872b0bcdee4[__obf_8650f6fe673c98a1+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1][__obf_cf967a75716d8446+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1] {
				__obf_65b7827e8fd2993e++
			}
		}
	}

	return __obf_65b7827e8fd2993e
}

// set sets the module at (x, y) to v.
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_5f29f833d6c91939(__obf_cf967a75716d8446 int, __obf_8650f6fe673c98a1 int, __obf_f85dcc6c0aea860b bool) {
	__obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3[__obf_8650f6fe673c98a1+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1][__obf_cf967a75716d8446+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1] = __obf_f85dcc6c0aea860b
	__obf_7b9309ca6f3e8da2.__obf_b0362872b0bcdee4[__obf_8650f6fe673c98a1+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1][__obf_cf967a75716d8446+__obf_7b9309ca6f3e8da2.__obf_dcb53264a169b9b1] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_b795da39576b2a49(__obf_cf967a75716d8446 int, __obf_8650f6fe673c98a1 int, __obf_f85dcc6c0aea860b [][]bool) {
	for __obf_02acc1a50ace36fe, __obf_52dfea4eee5f731a := range __obf_f85dcc6c0aea860b {
		for __obf_8715c27aa2bf2586, __obf_04a7a2ca242933f3 := range __obf_52dfea4eee5f731a {
			__obf_7b9309ca6f3e8da2.__obf_5f29f833d6c91939(__obf_cf967a75716d8446+__obf_8715c27aa2bf2586, __obf_8650f6fe673c98a1+__obf_02acc1a50ace36fe,

			// bitmap returns the entire symbol, including the quiet zone.
			__obf_04a7a2ca242933f3)
		}
	}
}

func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_d7ceef4be5069a61() [][]bool {
	__obf_8ade07d880d3e6f3 := make([][]bool, len(__obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3))

	for __obf_8715c27aa2bf2586 := range __obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3 {
		__obf_8ade07d880d3e6f3[__obf_8715c27aa2bf2586] = __obf_7b9309ca6f3e8da2.__obf_8ade07d880d3e6f3[__obf_8715c27aa2bf2586][:]
	}

	return __obf_8ade07d880d3e6f3
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
	__obf_4ac9a4de04c75e28 = 3
	__obf_316225081b7ee3c9 = 3
	__obf_10754466dc8aa8d5 = 40
	__obf_65cb610837fbfcfc = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_0a829429859f1a12() int {
	return __obf_7b9309ca6f3e8da2.__obf_60ace556ccc30342() + __obf_7b9309ca6f3e8da2.__obf_4b928cf0b4f52d58() + __obf_7b9309ca6f3e8da2.__obf_209814b7d34c47b3() + __obf_7b9309ca6f3e8da2.

	// penalty1 returns the penalty score for "adjacent modules in row/column with
	// same colour".
	//
	// The numbers of adjacent matching modules and scores are:
	// 0-5: score = 0
	// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
	__obf_2afa6e59c7c10319()
}

func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_60ace556ccc30342() int {
	__obf_d5cbfb4fd6fc4841 := 0

	for __obf_cf967a75716d8446 := 0; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
		__obf_f9c860f3bae8f954 := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, 0)
		__obf_65b7827e8fd2993e := 1

		for __obf_8650f6fe673c98a1 := 1; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
			__obf_f85dcc6c0aea860b := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1)

			if __obf_f85dcc6c0aea860b != __obf_f9c860f3bae8f954 {
				__obf_65b7827e8fd2993e = 1
				__obf_f9c860f3bae8f954 = __obf_f85dcc6c0aea860b
			} else {
				__obf_65b7827e8fd2993e++
				if __obf_65b7827e8fd2993e == 6 {
					__obf_d5cbfb4fd6fc4841 += __obf_4ac9a4de04c75e28 + 1
				} else if __obf_65b7827e8fd2993e > 6 {
					__obf_d5cbfb4fd6fc4841++
				}
			}
		}
	}

	for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
		__obf_f9c860f3bae8f954 := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(0, __obf_8650f6fe673c98a1)
		__obf_65b7827e8fd2993e := 1

		for __obf_cf967a75716d8446 := 1; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
			__obf_f85dcc6c0aea860b := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1)

			if __obf_f85dcc6c0aea860b != __obf_f9c860f3bae8f954 {
				__obf_65b7827e8fd2993e = 1
				__obf_f9c860f3bae8f954 = __obf_f85dcc6c0aea860b
			} else {
				__obf_65b7827e8fd2993e++
				if __obf_65b7827e8fd2993e == 6 {
					__obf_d5cbfb4fd6fc4841 += __obf_4ac9a4de04c75e28 + 1
				} else if __obf_65b7827e8fd2993e > 6 {
					__obf_d5cbfb4fd6fc4841++
				}
			}
		}
	}

	return __obf_d5cbfb4fd6fc4841
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_4b928cf0b4f52d58() int {
	__obf_d5cbfb4fd6fc4841 := 0

	for __obf_8650f6fe673c98a1 := 1; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
		for __obf_cf967a75716d8446 := 1; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
			__obf_c2f50cd59688b5fe := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446-1, __obf_8650f6fe673c98a1-1)
			__obf_6aa5daf7f92e0fc7 := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1-1)
			__obf_06a4c2e162b50098 := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446-1, __obf_8650f6fe673c98a1)
			__obf_c539d4dea3888ab5 := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1)

			if __obf_c539d4dea3888ab5 == __obf_06a4c2e162b50098 && __obf_c539d4dea3888ab5 == __obf_6aa5daf7f92e0fc7 && __obf_c539d4dea3888ab5 == __obf_c2f50cd59688b5fe {
				__obf_d5cbfb4fd6fc4841++
			}
		}
	}

	return __obf_d5cbfb4fd6fc4841 * __obf_316225081b7ee3c9
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_209814b7d34c47b3() int {
	__obf_d5cbfb4fd6fc4841 := 0

	for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
		var __obf_3fb560de43842fbf int16 = 0x00

		for __obf_cf967a75716d8446 := 0; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
			__obf_3fb560de43842fbf <<= 1
			if __obf_f85dcc6c0aea860b := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1); __obf_f85dcc6c0aea860b {
				__obf_3fb560de43842fbf |= 1
			}

			switch __obf_3fb560de43842fbf & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_d5cbfb4fd6fc4841 += __obf_10754466dc8aa8d5
				__obf_3fb560de43842fbf = 0xFF
			default:
				if __obf_cf967a75716d8446 == __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467-1 && (__obf_3fb560de43842fbf&0x7f) == 0x5d {
					__obf_d5cbfb4fd6fc4841 += __obf_10754466dc8aa8d5
					__obf_3fb560de43842fbf = 0xFF
				}
			}
		}
	}

	for __obf_cf967a75716d8446 := 0; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
		var __obf_3fb560de43842fbf int16 = 0x00

		for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
			__obf_3fb560de43842fbf <<= 1
			if __obf_f85dcc6c0aea860b := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1); __obf_f85dcc6c0aea860b {
				__obf_3fb560de43842fbf |= 1
			}

			switch __obf_3fb560de43842fbf & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_d5cbfb4fd6fc4841 += __obf_10754466dc8aa8d5
				__obf_3fb560de43842fbf = 0xFF
			default:
				if __obf_8650f6fe673c98a1 == __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467-1 && (__obf_3fb560de43842fbf&0x7f) == 0x5d {
					__obf_d5cbfb4fd6fc4841 += __obf_10754466dc8aa8d5
					__obf_3fb560de43842fbf = 0xFF
				}
			}
		}
	}

	return __obf_d5cbfb4fd6fc4841
}

// penalty4 returns the penalty score...
func (__obf_7b9309ca6f3e8da2 *__obf_14306cbbf0004950,) __obf_2afa6e59c7c10319() int {
	__obf_412c3ba2c9c51436 := __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467 * __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467
	__obf_0ac09e0b14d833b8 := 0

	for __obf_cf967a75716d8446 := 0; __obf_cf967a75716d8446 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_cf967a75716d8446++ {
		for __obf_8650f6fe673c98a1 := 0; __obf_8650f6fe673c98a1 < __obf_7b9309ca6f3e8da2.__obf_934aef36207d3467; __obf_8650f6fe673c98a1++ {
			if __obf_f85dcc6c0aea860b := __obf_7b9309ca6f3e8da2.__obf_55397a5408a37425(__obf_cf967a75716d8446, __obf_8650f6fe673c98a1); __obf_f85dcc6c0aea860b {
				__obf_0ac09e0b14d833b8++
			}
		}
	}
	__obf_b37ad646960fda55 := __obf_412c3ba2c9c51436/2 - __obf_0ac09e0b14d833b8
	if __obf_b37ad646960fda55 < 0 {
		__obf_b37ad646960fda55 *= -1
	}

	return __obf_65cb610837fbfcfc * (__obf_b37ad646960fda55 / (__obf_412c3ba2c9c51436 / 20))
}
