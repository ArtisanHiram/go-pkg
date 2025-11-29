// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_540b74a9d13704fa

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
type __obf_ebb7a995f536d449 struct {
	__obf_00dc7454ffb4f350 [// Value of module at [y][x]. True is set.
	][]bool
	__obf_0182d987177ff9f7 [// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	][]bool
	__obf_5315f29e47197a10 int// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_7bb0ccb8474ed4db int// Width/height of the symbol only.
	__obf_914e29178d61e4a9 int// Width/height of a single quiet zone.

}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_60e59e63cc171237(__obf_5315f29e47197a10 int, __obf_914e29178d61e4a9 int) *__obf_ebb7a995f536d449 {
	var __obf_628996e2b70b8b90 __obf_ebb7a995f536d449
	__obf_628996e2b70b8b90.__obf_00dc7454ffb4f350 = make([][]bool, __obf_5315f29e47197a10+2*__obf_914e29178d61e4a9)
	__obf_628996e2b70b8b90.__obf_0182d987177ff9f7 = make([][]bool, __obf_5315f29e47197a10+2*__obf_914e29178d61e4a9)

	for __obf_2e2482997826bcc5 := range __obf_628996e2b70b8b90.__obf_00dc7454ffb4f350 {
		__obf_628996e2b70b8b90.__obf_00dc7454ffb4f350[__obf_2e2482997826bcc5] = make([]bool, __obf_5315f29e47197a10+2*__obf_914e29178d61e4a9)
		__obf_628996e2b70b8b90.__obf_0182d987177ff9f7[__obf_2e2482997826bcc5] = make([]bool, __obf_5315f29e47197a10+2*__obf_914e29178d61e4a9)
	}
	__obf_628996e2b70b8b90.__obf_5315f29e47197a10 = __obf_5315f29e47197a10 + 2*__obf_914e29178d61e4a9
	__obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db = __obf_5315f29e47197a10
	__obf_628996e2b70b8b90.__obf_914e29178d61e4a9 = __obf_914e29178d61e4a9

	return &__obf_628996e2b70b8b90
}

// get returns the module value at (x, y).
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea int, __obf_9637ef17b76683ca int) (__obf_75eba50799fccf42 bool) {
	__obf_75eba50799fccf42 = __obf_628996e2b70b8b90.__obf_00dc7454ffb4f350[__obf_9637ef17b76683ca+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9][__obf_a7fd5f72cb6fabea+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_55d4fd802a4df174(__obf_a7fd5f72cb6fabea int, __obf_9637ef17b76683ca int) bool {
	return !__obf_628996e2b70b8b90.__obf_0182d987177ff9f7[__obf_9637ef17b76683ca+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9][__obf_a7fd5f72cb6fabea+__obf_628996e2b70b8b90.

	// numEmptyModules returns the number of empty modules.
	//
	// Initially numEmptyModules is symbolSize * symbolSize. After every module has
	// been set (to either true or false), the number of empty modules is zero.
	__obf_914e29178d61e4a9]
}

func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_730a1c84bbe28f0b() int {
	var __obf_f630a33de5b03001 int
	for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
		for __obf_a7fd5f72cb6fabea := 0; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
			if !__obf_628996e2b70b8b90.__obf_0182d987177ff9f7[__obf_9637ef17b76683ca+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9][__obf_a7fd5f72cb6fabea+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9] {
				__obf_f630a33de5b03001++
			}
		}
	}

	return __obf_f630a33de5b03001
}

// set sets the module at (x, y) to v.
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_72f8084f2a80600a(__obf_a7fd5f72cb6fabea int, __obf_9637ef17b76683ca int, __obf_75eba50799fccf42 bool) {
	__obf_628996e2b70b8b90.__obf_00dc7454ffb4f350[__obf_9637ef17b76683ca+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9][__obf_a7fd5f72cb6fabea+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9] = __obf_75eba50799fccf42
	__obf_628996e2b70b8b90.__obf_0182d987177ff9f7[__obf_9637ef17b76683ca+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9][__obf_a7fd5f72cb6fabea+__obf_628996e2b70b8b90.__obf_914e29178d61e4a9] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_fc8a722f849192fd(__obf_a7fd5f72cb6fabea int, __obf_9637ef17b76683ca int, __obf_75eba50799fccf42 [][]bool) {
	for __obf_cf60660918b25aa8, __obf_1bb374b1838198fd := range __obf_75eba50799fccf42 {
		for __obf_2e2482997826bcc5, __obf_efff9ae9a1dfb06e := range __obf_1bb374b1838198fd {
			__obf_628996e2b70b8b90.__obf_72f8084f2a80600a(__obf_a7fd5f72cb6fabea+__obf_2e2482997826bcc5, __obf_9637ef17b76683ca+__obf_cf60660918b25aa8,

			// bitmap returns the entire symbol, including the quiet zone.
			__obf_efff9ae9a1dfb06e)
		}
	}
}

func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_0f6cf8adbc21238e() [][]bool {
	__obf_00dc7454ffb4f350 := make([][]bool, len(__obf_628996e2b70b8b90.__obf_00dc7454ffb4f350))

	for __obf_2e2482997826bcc5 := range __obf_628996e2b70b8b90.__obf_00dc7454ffb4f350 {
		__obf_00dc7454ffb4f350[__obf_2e2482997826bcc5] = __obf_628996e2b70b8b90.__obf_00dc7454ffb4f350[__obf_2e2482997826bcc5][:]
	}

	return __obf_00dc7454ffb4f350
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
	__obf_1d97bb4181a8cacd = 3
	__obf_80977e15f082fa2b = 3
	__obf_b0053fc605b69672 = 40
	__obf_b8e64c207fe0d21c = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_60fe70efba5f9777() int {
	return __obf_628996e2b70b8b90.__obf_f0ddb8dfc2eea2b2() + __obf_628996e2b70b8b90.__obf_4dea7ece8b599abb() + __obf_628996e2b70b8b90.__obf_b4d51e08db453e25() + __obf_628996e2b70b8b90.

	// penalty1 returns the penalty score for "adjacent modules in row/column with
	// same colour".
	//
	// The numbers of adjacent matching modules and scores are:
	// 0-5: score = 0
	// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
	__obf_eaaacd3069599c62()
}

func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_f0ddb8dfc2eea2b2() int {
	__obf_17e9e44595f51024 := 0

	for __obf_a7fd5f72cb6fabea := 0; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
		__obf_8e3708a70ffeeda1 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, 0)
		__obf_f630a33de5b03001 := 1

		for __obf_9637ef17b76683ca := 1; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
			__obf_75eba50799fccf42 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca)

			if __obf_75eba50799fccf42 != __obf_8e3708a70ffeeda1 {
				__obf_f630a33de5b03001 = 1
				__obf_8e3708a70ffeeda1 = __obf_75eba50799fccf42
			} else {
				__obf_f630a33de5b03001++
				if __obf_f630a33de5b03001 == 6 {
					__obf_17e9e44595f51024 += __obf_1d97bb4181a8cacd + 1
				} else if __obf_f630a33de5b03001 > 6 {
					__obf_17e9e44595f51024++
				}
			}
		}
	}

	for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
		__obf_8e3708a70ffeeda1 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(0, __obf_9637ef17b76683ca)
		__obf_f630a33de5b03001 := 1

		for __obf_a7fd5f72cb6fabea := 1; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
			__obf_75eba50799fccf42 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca)

			if __obf_75eba50799fccf42 != __obf_8e3708a70ffeeda1 {
				__obf_f630a33de5b03001 = 1
				__obf_8e3708a70ffeeda1 = __obf_75eba50799fccf42
			} else {
				__obf_f630a33de5b03001++
				if __obf_f630a33de5b03001 == 6 {
					__obf_17e9e44595f51024 += __obf_1d97bb4181a8cacd + 1
				} else if __obf_f630a33de5b03001 > 6 {
					__obf_17e9e44595f51024++
				}
			}
		}
	}

	return __obf_17e9e44595f51024
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_4dea7ece8b599abb() int {
	__obf_17e9e44595f51024 := 0

	for __obf_9637ef17b76683ca := 1; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
		for __obf_a7fd5f72cb6fabea := 1; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
			__obf_38cc0027cf0bd0bb := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea-1, __obf_9637ef17b76683ca-1)
			__obf_d98ad32574cc7a08 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca-1)
			__obf_b8f8f8c92aa7278a := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea-1, __obf_9637ef17b76683ca)
			__obf_c2ab6ab9c78cef29 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca)

			if __obf_c2ab6ab9c78cef29 == __obf_b8f8f8c92aa7278a && __obf_c2ab6ab9c78cef29 == __obf_d98ad32574cc7a08 && __obf_c2ab6ab9c78cef29 == __obf_38cc0027cf0bd0bb {
				__obf_17e9e44595f51024++
			}
		}
	}

	return __obf_17e9e44595f51024 * __obf_80977e15f082fa2b
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_b4d51e08db453e25() int {
	__obf_17e9e44595f51024 := 0

	for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
		var __obf_0aaa832b0a849f8e int16 = 0x00

		for __obf_a7fd5f72cb6fabea := 0; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
			__obf_0aaa832b0a849f8e <<= 1
			if __obf_75eba50799fccf42 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca); __obf_75eba50799fccf42 {
				__obf_0aaa832b0a849f8e |= 1
			}

			switch __obf_0aaa832b0a849f8e & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_17e9e44595f51024 += __obf_b0053fc605b69672
				__obf_0aaa832b0a849f8e = 0xFF
			default:
				if __obf_a7fd5f72cb6fabea == __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db-1 && (__obf_0aaa832b0a849f8e&0x7f) == 0x5d {
					__obf_17e9e44595f51024 += __obf_b0053fc605b69672
					__obf_0aaa832b0a849f8e = 0xFF
				}
			}
		}
	}

	for __obf_a7fd5f72cb6fabea := 0; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
		var __obf_0aaa832b0a849f8e int16 = 0x00

		for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
			__obf_0aaa832b0a849f8e <<= 1
			if __obf_75eba50799fccf42 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca); __obf_75eba50799fccf42 {
				__obf_0aaa832b0a849f8e |= 1
			}

			switch __obf_0aaa832b0a849f8e & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_17e9e44595f51024 += __obf_b0053fc605b69672
				__obf_0aaa832b0a849f8e = 0xFF
			default:
				if __obf_9637ef17b76683ca == __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db-1 && (__obf_0aaa832b0a849f8e&0x7f) == 0x5d {
					__obf_17e9e44595f51024 += __obf_b0053fc605b69672
					__obf_0aaa832b0a849f8e = 0xFF
				}
			}
		}
	}

	return __obf_17e9e44595f51024
}

// penalty4 returns the penalty score...
func (__obf_628996e2b70b8b90 *__obf_ebb7a995f536d449,) __obf_eaaacd3069599c62() int {
	__obf_4de44b07fb8380b0 := __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db * __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db
	__obf_b2964993ff9f9617 := 0

	for __obf_a7fd5f72cb6fabea := 0; __obf_a7fd5f72cb6fabea < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_a7fd5f72cb6fabea++ {
		for __obf_9637ef17b76683ca := 0; __obf_9637ef17b76683ca < __obf_628996e2b70b8b90.__obf_7bb0ccb8474ed4db; __obf_9637ef17b76683ca++ {
			if __obf_75eba50799fccf42 := __obf_628996e2b70b8b90.__obf_aa19d2af5aae478d(__obf_a7fd5f72cb6fabea, __obf_9637ef17b76683ca); __obf_75eba50799fccf42 {
				__obf_b2964993ff9f9617++
			}
		}
	}
	__obf_038e7829ede19311 := __obf_4de44b07fb8380b0/2 - __obf_b2964993ff9f9617
	if __obf_038e7829ede19311 < 0 {
		__obf_038e7829ede19311 *= -1
	}

	return __obf_b8e64c207fe0d21c * (__obf_038e7829ede19311 / (__obf_4de44b07fb8380b0 / 20))
}
