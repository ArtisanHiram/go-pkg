// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_07675b7eb1c7284a

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
type __obf_e4275ce80c964a2d struct {
	__obf_56453d487990a23b [// Value of module at [y][x]. True is set.
	][]bool
	__obf_10494ad08bd1f3f9 [// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	][]bool
	__obf_0abbe8d67dfa1a14 int// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_567ab8938edd620e int// Width/height of the symbol only.
	__obf_a03b13c27e59ae4a int// Width/height of a single quiet zone.

}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_7ae41c427b32aafc(__obf_0abbe8d67dfa1a14 int, __obf_a03b13c27e59ae4a int) *__obf_e4275ce80c964a2d {
	var __obf_d1bf6b08b3ab56ed __obf_e4275ce80c964a2d
	__obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b = make([][]bool, __obf_0abbe8d67dfa1a14+2*__obf_a03b13c27e59ae4a)
	__obf_d1bf6b08b3ab56ed.__obf_10494ad08bd1f3f9 = make([][]bool, __obf_0abbe8d67dfa1a14+2*__obf_a03b13c27e59ae4a)

	for __obf_2e4170383a5ddec7 := range __obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b {
		__obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b[__obf_2e4170383a5ddec7] = make([]bool, __obf_0abbe8d67dfa1a14+2*__obf_a03b13c27e59ae4a)
		__obf_d1bf6b08b3ab56ed.__obf_10494ad08bd1f3f9[__obf_2e4170383a5ddec7] = make([]bool, __obf_0abbe8d67dfa1a14+2*__obf_a03b13c27e59ae4a)
	}
	__obf_d1bf6b08b3ab56ed.__obf_0abbe8d67dfa1a14 = __obf_0abbe8d67dfa1a14 + 2*__obf_a03b13c27e59ae4a
	__obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e = __obf_0abbe8d67dfa1a14
	__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a = __obf_a03b13c27e59ae4a

	return &__obf_d1bf6b08b3ab56ed
}

// get returns the module value at (x, y).
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_33a7bb8c092673d2(__obf_7107b684b2ee3954 int, __obf_21e5ea661cba6209 int) (__obf_283aa709fd2621d8 bool) {
	__obf_283aa709fd2621d8 = __obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b[__obf_21e5ea661cba6209+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a][__obf_7107b684b2ee3954+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_6dc1e71542360e8c(__obf_7107b684b2ee3954 int, __obf_21e5ea661cba6209 int) bool {
	return !__obf_d1bf6b08b3ab56ed.__obf_10494ad08bd1f3f9[__obf_21e5ea661cba6209+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a][__obf_7107b684b2ee3954+__obf_d1bf6b08b3ab56ed.

	// numEmptyModules returns the number of empty modules.
	//
	// Initially numEmptyModules is symbolSize * symbolSize. After every module has
	// been set (to either true or false), the number of empty modules is zero.
	__obf_a03b13c27e59ae4a]
}

func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_8933e3589fae1c15() int {
	var __obf_24286ef66b6ad0a9 int
	for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
		for __obf_7107b684b2ee3954 := 0; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
			if !__obf_d1bf6b08b3ab56ed.__obf_10494ad08bd1f3f9[__obf_21e5ea661cba6209+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a][__obf_7107b684b2ee3954+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a] {
				__obf_24286ef66b6ad0a9++
			}
		}
	}

	return __obf_24286ef66b6ad0a9
}

// set sets the module at (x, y) to v.
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_1b96d0c4df8d7fac(__obf_7107b684b2ee3954 int, __obf_21e5ea661cba6209 int, __obf_283aa709fd2621d8 bool) {
	__obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b[__obf_21e5ea661cba6209+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a][__obf_7107b684b2ee3954+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a] = __obf_283aa709fd2621d8
	__obf_d1bf6b08b3ab56ed.__obf_10494ad08bd1f3f9[__obf_21e5ea661cba6209+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a][__obf_7107b684b2ee3954+__obf_d1bf6b08b3ab56ed.__obf_a03b13c27e59ae4a] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_3f3b9c9a0d5abbc1(__obf_7107b684b2ee3954 int, __obf_21e5ea661cba6209 int, __obf_283aa709fd2621d8 [][]bool) {
	for __obf_95b5b5fef8a19647, __obf_158fbe0b9c041995 := range __obf_283aa709fd2621d8 {
		for __obf_2e4170383a5ddec7, __obf_6c9b4b130cbe971d := range __obf_158fbe0b9c041995 {
			__obf_d1bf6b08b3ab56ed.__obf_1b96d0c4df8d7fac(__obf_7107b684b2ee3954+__obf_2e4170383a5ddec7, __obf_21e5ea661cba6209+__obf_95b5b5fef8a19647,

			// bitmap returns the entire symbol, including the quiet zone.
			__obf_6c9b4b130cbe971d)
		}
	}
}

func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_e6043f786f7c3cb5() [][]bool {
	__obf_56453d487990a23b := make([][]bool, len(__obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b))

	for __obf_2e4170383a5ddec7 := range __obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b {
		__obf_56453d487990a23b[__obf_2e4170383a5ddec7] = __obf_d1bf6b08b3ab56ed.__obf_56453d487990a23b[__obf_2e4170383a5ddec7][:]
	}

	return __obf_56453d487990a23b
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
	__obf_f5c57a0d6771871c = 3
	__obf_d1a6567c448bcfec = 3
	__obf_fffc244ce1617099 = 40
	__obf_7a8dd816f3676702 = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_fcbea8a1aa2a09d3() int {
	return __obf_d1bf6b08b3ab56ed.__obf_34dd952ea7e245ca() + __obf_d1bf6b08b3ab56ed.__obf_60d0ee71de1b4a8e() + __obf_d1bf6b08b3ab56ed.__obf_ece7d9006e5372ee() + __obf_d1bf6b08b3ab56ed.

	// penalty1 returns the penalty score for "adjacent modules in row/column with
	// same colour".
	//
	// The numbers of adjacent matching modules and scores are:
	// 0-5: score = 0
	// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
	__obf_3e7d79ae46124487()
}

func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_34dd952ea7e245ca() int {
	__obf_a633b45f4e29f865 := 0

	for __obf_7107b684b2ee3954 := 0; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
		__obf_92b235440ed25a65 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, 0)
		__obf_24286ef66b6ad0a9 := 1

		for __obf_21e5ea661cba6209 := 1; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
			__obf_283aa709fd2621d8 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209)

			if __obf_283aa709fd2621d8 != __obf_92b235440ed25a65 {
				__obf_24286ef66b6ad0a9 = 1
				__obf_92b235440ed25a65 = __obf_283aa709fd2621d8
			} else {
				__obf_24286ef66b6ad0a9++
				if __obf_24286ef66b6ad0a9 == 6 {
					__obf_a633b45f4e29f865 += __obf_f5c57a0d6771871c + 1
				} else if __obf_24286ef66b6ad0a9 > 6 {
					__obf_a633b45f4e29f865++
				}
			}
		}
	}

	for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
		__obf_92b235440ed25a65 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(0, __obf_21e5ea661cba6209)
		__obf_24286ef66b6ad0a9 := 1

		for __obf_7107b684b2ee3954 := 1; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
			__obf_283aa709fd2621d8 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209)

			if __obf_283aa709fd2621d8 != __obf_92b235440ed25a65 {
				__obf_24286ef66b6ad0a9 = 1
				__obf_92b235440ed25a65 = __obf_283aa709fd2621d8
			} else {
				__obf_24286ef66b6ad0a9++
				if __obf_24286ef66b6ad0a9 == 6 {
					__obf_a633b45f4e29f865 += __obf_f5c57a0d6771871c + 1
				} else if __obf_24286ef66b6ad0a9 > 6 {
					__obf_a633b45f4e29f865++
				}
			}
		}
	}

	return __obf_a633b45f4e29f865
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_60d0ee71de1b4a8e() int {
	__obf_a633b45f4e29f865 := 0

	for __obf_21e5ea661cba6209 := 1; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
		for __obf_7107b684b2ee3954 := 1; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
			__obf_ca922d6ad9f3a693 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954-1, __obf_21e5ea661cba6209-1)
			__obf_0ef8ae9f6ac75029 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209-1)
			__obf_9728026a19b337ce := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954-1, __obf_21e5ea661cba6209)
			__obf_366fff63f95e2d31 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209)

			if __obf_366fff63f95e2d31 == __obf_9728026a19b337ce && __obf_366fff63f95e2d31 == __obf_0ef8ae9f6ac75029 && __obf_366fff63f95e2d31 == __obf_ca922d6ad9f3a693 {
				__obf_a633b45f4e29f865++
			}
		}
	}

	return __obf_a633b45f4e29f865 * __obf_d1a6567c448bcfec
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_ece7d9006e5372ee() int {
	__obf_a633b45f4e29f865 := 0

	for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
		var __obf_be02a994b3a662d5 int16 = 0x00

		for __obf_7107b684b2ee3954 := 0; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
			__obf_be02a994b3a662d5 <<= 1
			if __obf_283aa709fd2621d8 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209); __obf_283aa709fd2621d8 {
				__obf_be02a994b3a662d5 |= 1
			}

			switch __obf_be02a994b3a662d5 & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_a633b45f4e29f865 += __obf_fffc244ce1617099
				__obf_be02a994b3a662d5 = 0xFF
			default:
				if __obf_7107b684b2ee3954 == __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e-1 && (__obf_be02a994b3a662d5&0x7f) == 0x5d {
					__obf_a633b45f4e29f865 += __obf_fffc244ce1617099
					__obf_be02a994b3a662d5 = 0xFF
				}
			}
		}
	}

	for __obf_7107b684b2ee3954 := 0; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
		var __obf_be02a994b3a662d5 int16 = 0x00

		for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
			__obf_be02a994b3a662d5 <<= 1
			if __obf_283aa709fd2621d8 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209); __obf_283aa709fd2621d8 {
				__obf_be02a994b3a662d5 |= 1
			}

			switch __obf_be02a994b3a662d5 & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_a633b45f4e29f865 += __obf_fffc244ce1617099
				__obf_be02a994b3a662d5 = 0xFF
			default:
				if __obf_21e5ea661cba6209 == __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e-1 && (__obf_be02a994b3a662d5&0x7f) == 0x5d {
					__obf_a633b45f4e29f865 += __obf_fffc244ce1617099
					__obf_be02a994b3a662d5 = 0xFF
				}
			}
		}
	}

	return __obf_a633b45f4e29f865
}

// penalty4 returns the penalty score...
func (__obf_d1bf6b08b3ab56ed *__obf_e4275ce80c964a2d,) __obf_3e7d79ae46124487() int {
	__obf_5938e792887bb25d := __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e * __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e
	__obf_b4a826699b1e1e94 := 0

	for __obf_7107b684b2ee3954 := 0; __obf_7107b684b2ee3954 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_7107b684b2ee3954++ {
		for __obf_21e5ea661cba6209 := 0; __obf_21e5ea661cba6209 < __obf_d1bf6b08b3ab56ed.__obf_567ab8938edd620e; __obf_21e5ea661cba6209++ {
			if __obf_283aa709fd2621d8 := __obf_d1bf6b08b3ab56ed.__obf_33a7bb8c092673d2(__obf_7107b684b2ee3954, __obf_21e5ea661cba6209); __obf_283aa709fd2621d8 {
				__obf_b4a826699b1e1e94++
			}
		}
	}
	__obf_a8439080806800df := __obf_5938e792887bb25d/2 - __obf_b4a826699b1e1e94
	if __obf_a8439080806800df < 0 {
		__obf_a8439080806800df *= -1
	}

	return __obf_7a8dd816f3676702 * (__obf_a8439080806800df / (__obf_5938e792887bb25d / 20))
}
