// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_63dc7f87a7750084

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
type __obf_a9c08ca051c7e6c3 struct {
	__obf_6992edcaa16b67fa [// Value of module at [y][x]. True is set.
	][]bool
	__obf_c79239acbf74f51f [// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	][]bool
	__obf_80f4ab6d4332abb9 int// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_f2408f5081085d49 int// Width/height of the symbol only.
	__obf_33b2b94fa59e1169 int// Width/height of a single quiet zone.

}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_1236b406b8a6e2f9(__obf_80f4ab6d4332abb9 int, __obf_33b2b94fa59e1169 int) *__obf_a9c08ca051c7e6c3 {
	var __obf_6c3057b1bb15c1d0 __obf_a9c08ca051c7e6c3
	__obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa = make([][]bool, __obf_80f4ab6d4332abb9+2*__obf_33b2b94fa59e1169)
	__obf_6c3057b1bb15c1d0.__obf_c79239acbf74f51f = make([][]bool, __obf_80f4ab6d4332abb9+2*__obf_33b2b94fa59e1169)

	for __obf_b39cd0fd88eaf53c := range __obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa {
		__obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa[__obf_b39cd0fd88eaf53c] = make([]bool, __obf_80f4ab6d4332abb9+2*__obf_33b2b94fa59e1169)
		__obf_6c3057b1bb15c1d0.__obf_c79239acbf74f51f[__obf_b39cd0fd88eaf53c] = make([]bool, __obf_80f4ab6d4332abb9+2*__obf_33b2b94fa59e1169)
	}
	__obf_6c3057b1bb15c1d0.__obf_80f4ab6d4332abb9 = __obf_80f4ab6d4332abb9 + 2*__obf_33b2b94fa59e1169
	__obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49 = __obf_80f4ab6d4332abb9
	__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169 = __obf_33b2b94fa59e1169

	return &__obf_6c3057b1bb15c1d0
}

// get returns the module value at (x, y).
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_472ab334bb09a22a(__obf_7529943f67ee44cb int, __obf_16bfdfd36c335382 int) (__obf_d3308b8b9ec8593a bool) {
	__obf_d3308b8b9ec8593a = __obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa[__obf_16bfdfd36c335382+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169][__obf_7529943f67ee44cb+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_497c9b705268f269(__obf_7529943f67ee44cb int, __obf_16bfdfd36c335382 int) bool {
	return !__obf_6c3057b1bb15c1d0.__obf_c79239acbf74f51f[__obf_16bfdfd36c335382+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169][__obf_7529943f67ee44cb+__obf_6c3057b1bb15c1d0.

	// numEmptyModules returns the number of empty modules.
	//
	// Initially numEmptyModules is symbolSize * symbolSize. After every module has
	// been set (to either true or false), the number of empty modules is zero.
	__obf_33b2b94fa59e1169]
}

func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_955c9aaf43109184() int {
	var __obf_475706a6f788d1e6 int
	for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
		for __obf_7529943f67ee44cb := 0; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
			if !__obf_6c3057b1bb15c1d0.__obf_c79239acbf74f51f[__obf_16bfdfd36c335382+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169][__obf_7529943f67ee44cb+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169] {
				__obf_475706a6f788d1e6++
			}
		}
	}

	return __obf_475706a6f788d1e6
}

// set sets the module at (x, y) to v.
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_f621c41973e6ce2f(__obf_7529943f67ee44cb int, __obf_16bfdfd36c335382 int, __obf_d3308b8b9ec8593a bool) {
	__obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa[__obf_16bfdfd36c335382+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169][__obf_7529943f67ee44cb+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169] = __obf_d3308b8b9ec8593a
	__obf_6c3057b1bb15c1d0.__obf_c79239acbf74f51f[__obf_16bfdfd36c335382+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169][__obf_7529943f67ee44cb+__obf_6c3057b1bb15c1d0.__obf_33b2b94fa59e1169] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_5a8662f11dc82931(__obf_7529943f67ee44cb int, __obf_16bfdfd36c335382 int, __obf_d3308b8b9ec8593a [][]bool) {
	for __obf_0d2b5442049f8295, __obf_79f6304f07a3b57b := range __obf_d3308b8b9ec8593a {
		for __obf_b39cd0fd88eaf53c, __obf_85088f87763437a5 := range __obf_79f6304f07a3b57b {
			__obf_6c3057b1bb15c1d0.__obf_f621c41973e6ce2f(__obf_7529943f67ee44cb+__obf_b39cd0fd88eaf53c, __obf_16bfdfd36c335382+__obf_0d2b5442049f8295,

			// bitmap returns the entire symbol, including the quiet zone.
			__obf_85088f87763437a5)
		}
	}
}

func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_94902da22ea0120e() [][]bool {
	__obf_6992edcaa16b67fa := make([][]bool, len(__obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa))

	for __obf_b39cd0fd88eaf53c := range __obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa {
		__obf_6992edcaa16b67fa[__obf_b39cd0fd88eaf53c] = __obf_6c3057b1bb15c1d0.__obf_6992edcaa16b67fa[__obf_b39cd0fd88eaf53c][:]
	}

	return __obf_6992edcaa16b67fa
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
	__obf_f180bfad19d9bd74 = 3
	__obf_402994f787875ea1 = 3
	__obf_cb529b4a7c83e52b = 40
	__obf_46d45806ce719fd3 = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_9dab437c9fcdbe57() int {
	return __obf_6c3057b1bb15c1d0.__obf_43f20d23e3f212f5() + __obf_6c3057b1bb15c1d0.__obf_b14db37818c24e1f() + __obf_6c3057b1bb15c1d0.__obf_e7c21d47ac986035() + __obf_6c3057b1bb15c1d0.

	// penalty1 returns the penalty score for "adjacent modules in row/column with
	// same colour".
	//
	// The numbers of adjacent matching modules and scores are:
	// 0-5: score = 0
	// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
	__obf_32c0dd100702ac4b()
}

func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_43f20d23e3f212f5() int {
	__obf_a1a2e98604fc16f7 := 0

	for __obf_7529943f67ee44cb := 0; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
		__obf_d69f92fc79c57b10 := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, 0)
		__obf_475706a6f788d1e6 := 1

		for __obf_16bfdfd36c335382 := 1; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
			__obf_d3308b8b9ec8593a := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382)

			if __obf_d3308b8b9ec8593a != __obf_d69f92fc79c57b10 {
				__obf_475706a6f788d1e6 = 1
				__obf_d69f92fc79c57b10 = __obf_d3308b8b9ec8593a
			} else {
				__obf_475706a6f788d1e6++
				if __obf_475706a6f788d1e6 == 6 {
					__obf_a1a2e98604fc16f7 += __obf_f180bfad19d9bd74 + 1
				} else if __obf_475706a6f788d1e6 > 6 {
					__obf_a1a2e98604fc16f7++
				}
			}
		}
	}

	for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
		__obf_d69f92fc79c57b10 := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(0, __obf_16bfdfd36c335382)
		__obf_475706a6f788d1e6 := 1

		for __obf_7529943f67ee44cb := 1; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
			__obf_d3308b8b9ec8593a := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382)

			if __obf_d3308b8b9ec8593a != __obf_d69f92fc79c57b10 {
				__obf_475706a6f788d1e6 = 1
				__obf_d69f92fc79c57b10 = __obf_d3308b8b9ec8593a
			} else {
				__obf_475706a6f788d1e6++
				if __obf_475706a6f788d1e6 == 6 {
					__obf_a1a2e98604fc16f7 += __obf_f180bfad19d9bd74 + 1
				} else if __obf_475706a6f788d1e6 > 6 {
					__obf_a1a2e98604fc16f7++
				}
			}
		}
	}

	return __obf_a1a2e98604fc16f7
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_b14db37818c24e1f() int {
	__obf_a1a2e98604fc16f7 := 0

	for __obf_16bfdfd36c335382 := 1; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
		for __obf_7529943f67ee44cb := 1; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
			__obf_e04a45ac50c97f36 := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb-1, __obf_16bfdfd36c335382-1)
			__obf_68dad60da097db07 := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382-1)
			__obf_86db66dd953a19c4 := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb-1, __obf_16bfdfd36c335382)
			__obf_4fad8740b4d5bfaf := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382)

			if __obf_4fad8740b4d5bfaf == __obf_86db66dd953a19c4 && __obf_4fad8740b4d5bfaf == __obf_68dad60da097db07 && __obf_4fad8740b4d5bfaf == __obf_e04a45ac50c97f36 {
				__obf_a1a2e98604fc16f7++
			}
		}
	}

	return __obf_a1a2e98604fc16f7 * __obf_402994f787875ea1
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_e7c21d47ac986035() int {
	__obf_a1a2e98604fc16f7 := 0

	for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
		var __obf_4c5d7fd40a283ec0 int16 = 0x00

		for __obf_7529943f67ee44cb := 0; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
			__obf_4c5d7fd40a283ec0 <<= 1
			if __obf_d3308b8b9ec8593a := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382); __obf_d3308b8b9ec8593a {
				__obf_4c5d7fd40a283ec0 |= 1
			}

			switch __obf_4c5d7fd40a283ec0 & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_a1a2e98604fc16f7 += __obf_cb529b4a7c83e52b
				__obf_4c5d7fd40a283ec0 = 0xFF
			default:
				if __obf_7529943f67ee44cb == __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49-1 && (__obf_4c5d7fd40a283ec0&0x7f) == 0x5d {
					__obf_a1a2e98604fc16f7 += __obf_cb529b4a7c83e52b
					__obf_4c5d7fd40a283ec0 = 0xFF
				}
			}
		}
	}

	for __obf_7529943f67ee44cb := 0; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
		var __obf_4c5d7fd40a283ec0 int16 = 0x00

		for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
			__obf_4c5d7fd40a283ec0 <<= 1
			if __obf_d3308b8b9ec8593a := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382); __obf_d3308b8b9ec8593a {
				__obf_4c5d7fd40a283ec0 |= 1
			}

			switch __obf_4c5d7fd40a283ec0 & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_a1a2e98604fc16f7 += __obf_cb529b4a7c83e52b
				__obf_4c5d7fd40a283ec0 = 0xFF
			default:
				if __obf_16bfdfd36c335382 == __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49-1 && (__obf_4c5d7fd40a283ec0&0x7f) == 0x5d {
					__obf_a1a2e98604fc16f7 += __obf_cb529b4a7c83e52b
					__obf_4c5d7fd40a283ec0 = 0xFF
				}
			}
		}
	}

	return __obf_a1a2e98604fc16f7
}

// penalty4 returns the penalty score...
func (__obf_6c3057b1bb15c1d0 *__obf_a9c08ca051c7e6c3,) __obf_32c0dd100702ac4b() int {
	__obf_6cb102c55684aa44 := __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49 * __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49
	__obf_767baccaacb54189 := 0

	for __obf_7529943f67ee44cb := 0; __obf_7529943f67ee44cb < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_7529943f67ee44cb++ {
		for __obf_16bfdfd36c335382 := 0; __obf_16bfdfd36c335382 < __obf_6c3057b1bb15c1d0.__obf_f2408f5081085d49; __obf_16bfdfd36c335382++ {
			if __obf_d3308b8b9ec8593a := __obf_6c3057b1bb15c1d0.__obf_472ab334bb09a22a(__obf_7529943f67ee44cb, __obf_16bfdfd36c335382); __obf_d3308b8b9ec8593a {
				__obf_767baccaacb54189++
			}
		}
	}
	__obf_65fcaa3ea11d5a03 := __obf_6cb102c55684aa44/2 - __obf_767baccaacb54189
	if __obf_65fcaa3ea11d5a03 < 0 {
		__obf_65fcaa3ea11d5a03 *= -1
	}

	return __obf_46d45806ce719fd3 * (__obf_65fcaa3ea11d5a03 / (__obf_6cb102c55684aa44 / 20))
}
