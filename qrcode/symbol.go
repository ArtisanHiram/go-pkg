// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_380fc15a74e6e942

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
type __obf_87d3969cdacfe5ff struct {
	__obf_dd9e344f41d7f12d [// Value of module at [y][x]. True is set.
	][]bool
	__obf_587efdba899e2aa8 [// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	][]bool
	__obf_5077244fac9670a3 int// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_108473f4f8f75beb int// Width/height of the symbol only.
	__obf_7674347e39061c69 int// Width/height of a single quiet zone.

}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_f4d308c42ad8a72a(__obf_5077244fac9670a3 int, __obf_7674347e39061c69 int) *__obf_87d3969cdacfe5ff {
	var __obf_e5f8610bf4c5dd2c __obf_87d3969cdacfe5ff
	__obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d = make([][]bool, __obf_5077244fac9670a3+2*__obf_7674347e39061c69)
	__obf_e5f8610bf4c5dd2c.__obf_587efdba899e2aa8 = make([][]bool, __obf_5077244fac9670a3+2*__obf_7674347e39061c69)

	for __obf_f1022e2071655fce := range __obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d {
		__obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d[__obf_f1022e2071655fce] = make([]bool, __obf_5077244fac9670a3+2*__obf_7674347e39061c69)
		__obf_e5f8610bf4c5dd2c.__obf_587efdba899e2aa8[__obf_f1022e2071655fce] = make([]bool, __obf_5077244fac9670a3+2*__obf_7674347e39061c69)
	}
	__obf_e5f8610bf4c5dd2c.__obf_5077244fac9670a3 = __obf_5077244fac9670a3 + 2*__obf_7674347e39061c69
	__obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb = __obf_5077244fac9670a3
	__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69 = __obf_7674347e39061c69

	return &__obf_e5f8610bf4c5dd2c
}

// get returns the module value at (x, y).
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_734fe92724eefb2b(__obf_e3f2eace849c7041 int, __obf_8eaa1145d0e6ec9a int) (__obf_3591b555f7f18238 bool) {
	__obf_3591b555f7f18238 = __obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d[__obf_8eaa1145d0e6ec9a+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69][__obf_e3f2eace849c7041+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_26ef80c094cd0692(__obf_e3f2eace849c7041 int, __obf_8eaa1145d0e6ec9a int) bool {
	return !__obf_e5f8610bf4c5dd2c.__obf_587efdba899e2aa8[__obf_8eaa1145d0e6ec9a+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69][__obf_e3f2eace849c7041+__obf_e5f8610bf4c5dd2c.

	// numEmptyModules returns the number of empty modules.
	//
	// Initially numEmptyModules is symbolSize * symbolSize. After every module has
	// been set (to either true or false), the number of empty modules is zero.
	__obf_7674347e39061c69]
}

func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_20c8c0869c8b5b37() int {
	var __obf_c039b15fe0c8cc4f int
	for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
		for __obf_e3f2eace849c7041 := 0; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
			if !__obf_e5f8610bf4c5dd2c.__obf_587efdba899e2aa8[__obf_8eaa1145d0e6ec9a+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69][__obf_e3f2eace849c7041+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69] {
				__obf_c039b15fe0c8cc4f++
			}
		}
	}

	return __obf_c039b15fe0c8cc4f
}

// set sets the module at (x, y) to v.
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_e551bdefdf23bb09(__obf_e3f2eace849c7041 int, __obf_8eaa1145d0e6ec9a int, __obf_3591b555f7f18238 bool) {
	__obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d[__obf_8eaa1145d0e6ec9a+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69][__obf_e3f2eace849c7041+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69] = __obf_3591b555f7f18238
	__obf_e5f8610bf4c5dd2c.__obf_587efdba899e2aa8[__obf_8eaa1145d0e6ec9a+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69][__obf_e3f2eace849c7041+__obf_e5f8610bf4c5dd2c.__obf_7674347e39061c69] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_f8c785d15edcc7d8(__obf_e3f2eace849c7041 int, __obf_8eaa1145d0e6ec9a int, __obf_3591b555f7f18238 [][]bool) {
	for __obf_e8fc5f84cbc468cf, __obf_0d0b344a1e699432 := range __obf_3591b555f7f18238 {
		for __obf_f1022e2071655fce, __obf_7c3cf7d3ea763ee4 := range __obf_0d0b344a1e699432 {
			__obf_e5f8610bf4c5dd2c.__obf_e551bdefdf23bb09(__obf_e3f2eace849c7041+__obf_f1022e2071655fce, __obf_8eaa1145d0e6ec9a+__obf_e8fc5f84cbc468cf,

			// bitmap returns the entire symbol, including the quiet zone.
			__obf_7c3cf7d3ea763ee4)
		}
	}
}

func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_0d0a40eb3c5b3d00() [][]bool {
	__obf_dd9e344f41d7f12d := make([][]bool, len(__obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d))

	for __obf_f1022e2071655fce := range __obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d {
		__obf_dd9e344f41d7f12d[__obf_f1022e2071655fce] = __obf_e5f8610bf4c5dd2c.__obf_dd9e344f41d7f12d[__obf_f1022e2071655fce][:]
	}

	return __obf_dd9e344f41d7f12d
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
	__obf_43fddee8270d06b8 = 3
	__obf_d1ae55b691c470e9 = 3
	__obf_a8900f69961c0c53 = 40
	__obf_780e085890761587 = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_a62e72882ae00260() int {
	return __obf_e5f8610bf4c5dd2c.__obf_3c0824592e23050c() + __obf_e5f8610bf4c5dd2c.__obf_e0d286fb5c78b072() + __obf_e5f8610bf4c5dd2c.__obf_d93febdc18acb275() + __obf_e5f8610bf4c5dd2c.

	// penalty1 returns the penalty score for "adjacent modules in row/column with
	// same colour".
	//
	// The numbers of adjacent matching modules and scores are:
	// 0-5: score = 0
	// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
	__obf_46d6f6adcf4a0c1f()
}

func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_3c0824592e23050c() int {
	__obf_d40e98e6d42d4c72 := 0

	for __obf_e3f2eace849c7041 := 0; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
		__obf_ef0e9247a4024f44 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, 0)
		__obf_c039b15fe0c8cc4f := 1

		for __obf_8eaa1145d0e6ec9a := 1; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
			__obf_3591b555f7f18238 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a)

			if __obf_3591b555f7f18238 != __obf_ef0e9247a4024f44 {
				__obf_c039b15fe0c8cc4f = 1
				__obf_ef0e9247a4024f44 = __obf_3591b555f7f18238
			} else {
				__obf_c039b15fe0c8cc4f++
				if __obf_c039b15fe0c8cc4f == 6 {
					__obf_d40e98e6d42d4c72 += __obf_43fddee8270d06b8 + 1
				} else if __obf_c039b15fe0c8cc4f > 6 {
					__obf_d40e98e6d42d4c72++
				}
			}
		}
	}

	for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
		__obf_ef0e9247a4024f44 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(0, __obf_8eaa1145d0e6ec9a)
		__obf_c039b15fe0c8cc4f := 1

		for __obf_e3f2eace849c7041 := 1; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
			__obf_3591b555f7f18238 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a)

			if __obf_3591b555f7f18238 != __obf_ef0e9247a4024f44 {
				__obf_c039b15fe0c8cc4f = 1
				__obf_ef0e9247a4024f44 = __obf_3591b555f7f18238
			} else {
				__obf_c039b15fe0c8cc4f++
				if __obf_c039b15fe0c8cc4f == 6 {
					__obf_d40e98e6d42d4c72 += __obf_43fddee8270d06b8 + 1
				} else if __obf_c039b15fe0c8cc4f > 6 {
					__obf_d40e98e6d42d4c72++
				}
			}
		}
	}

	return __obf_d40e98e6d42d4c72
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_e0d286fb5c78b072() int {
	__obf_d40e98e6d42d4c72 := 0

	for __obf_8eaa1145d0e6ec9a := 1; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
		for __obf_e3f2eace849c7041 := 1; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
			__obf_25d25bf1d34c7944 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041-1, __obf_8eaa1145d0e6ec9a-1)
			__obf_90baf826c417991d := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a-1)
			__obf_7c023da221985594 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041-1, __obf_8eaa1145d0e6ec9a)
			__obf_9775b1d953b0b135 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a)

			if __obf_9775b1d953b0b135 == __obf_7c023da221985594 && __obf_9775b1d953b0b135 == __obf_90baf826c417991d && __obf_9775b1d953b0b135 == __obf_25d25bf1d34c7944 {
				__obf_d40e98e6d42d4c72++
			}
		}
	}

	return __obf_d40e98e6d42d4c72 * __obf_d1ae55b691c470e9
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_d93febdc18acb275() int {
	__obf_d40e98e6d42d4c72 := 0

	for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
		var __obf_76cecf223cd47cff int16 = 0x00

		for __obf_e3f2eace849c7041 := 0; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
			__obf_76cecf223cd47cff <<= 1
			if __obf_3591b555f7f18238 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a); __obf_3591b555f7f18238 {
				__obf_76cecf223cd47cff |= 1
			}

			switch __obf_76cecf223cd47cff & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_d40e98e6d42d4c72 += __obf_a8900f69961c0c53
				__obf_76cecf223cd47cff = 0xFF
			default:
				if __obf_e3f2eace849c7041 == __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb-1 && (__obf_76cecf223cd47cff&0x7f) == 0x5d {
					__obf_d40e98e6d42d4c72 += __obf_a8900f69961c0c53
					__obf_76cecf223cd47cff = 0xFF
				}
			}
		}
	}

	for __obf_e3f2eace849c7041 := 0; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
		var __obf_76cecf223cd47cff int16 = 0x00

		for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
			__obf_76cecf223cd47cff <<= 1
			if __obf_3591b555f7f18238 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a); __obf_3591b555f7f18238 {
				__obf_76cecf223cd47cff |= 1
			}

			switch __obf_76cecf223cd47cff & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_d40e98e6d42d4c72 += __obf_a8900f69961c0c53
				__obf_76cecf223cd47cff = 0xFF
			default:
				if __obf_8eaa1145d0e6ec9a == __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb-1 && (__obf_76cecf223cd47cff&0x7f) == 0x5d {
					__obf_d40e98e6d42d4c72 += __obf_a8900f69961c0c53
					__obf_76cecf223cd47cff = 0xFF
				}
			}
		}
	}

	return __obf_d40e98e6d42d4c72
}

// penalty4 returns the penalty score...
func (__obf_e5f8610bf4c5dd2c *__obf_87d3969cdacfe5ff,) __obf_46d6f6adcf4a0c1f() int {
	__obf_d94a8fd987a5e1b6 := __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb * __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb
	__obf_b09ab1904b91e511 := 0

	for __obf_e3f2eace849c7041 := 0; __obf_e3f2eace849c7041 < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_e3f2eace849c7041++ {
		for __obf_8eaa1145d0e6ec9a := 0; __obf_8eaa1145d0e6ec9a < __obf_e5f8610bf4c5dd2c.__obf_108473f4f8f75beb; __obf_8eaa1145d0e6ec9a++ {
			if __obf_3591b555f7f18238 := __obf_e5f8610bf4c5dd2c.__obf_734fe92724eefb2b(__obf_e3f2eace849c7041, __obf_8eaa1145d0e6ec9a); __obf_3591b555f7f18238 {
				__obf_b09ab1904b91e511++
			}
		}
	}
	__obf_7bb5b1b93bcc030b := __obf_d94a8fd987a5e1b6/2 - __obf_b09ab1904b91e511
	if __obf_7bb5b1b93bcc030b < 0 {
		__obf_7bb5b1b93bcc030b *= -1
	}

	return __obf_780e085890761587 * (__obf_7bb5b1b93bcc030b / (__obf_d94a8fd987a5e1b6 / 20))
}
