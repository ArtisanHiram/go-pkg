// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_d8b53a99900c2ed7

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
type __obf_dbf6f541061eeeb5 struct {
	// Value of module at [y][x]. True is set.
	__obf_96e599f6f35e9507 [][]bool

	// True if the module at [y][x] is used (to either true or false).
	// Used to identify unused modules.
	__obf_5287a97af7a4fcb2 [][]bool

	// Combined width/height of the symbol and quiet zones.
	//
	// size = symbolSize + 2*quietZoneSize.
	__obf_658024c8e4ec9df7 int

	// Width/height of the symbol only.
	__obf_dc6945971f3283f0 int

	// Width/height of a single quiet zone.
	__obf_3252eb05fbf6197a int
}

// newSymbol constructs a symbol of size size*size, with a border of
// quietZoneSize.
func __obf_a45f602b0d096c7f(__obf_658024c8e4ec9df7 int, __obf_3252eb05fbf6197a int) *__obf_dbf6f541061eeeb5 {
	var __obf_77b51aa28d7a9d4e __obf_dbf6f541061eeeb5

	__obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507 = make([][]bool, __obf_658024c8e4ec9df7+2*__obf_3252eb05fbf6197a)
	__obf_77b51aa28d7a9d4e.__obf_5287a97af7a4fcb2 = make([][]bool, __obf_658024c8e4ec9df7+2*__obf_3252eb05fbf6197a)

	for __obf_6ffda43d9ecd66ed := range __obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507 {
		__obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507[__obf_6ffda43d9ecd66ed] = make([]bool, __obf_658024c8e4ec9df7+2*__obf_3252eb05fbf6197a)
		__obf_77b51aa28d7a9d4e.__obf_5287a97af7a4fcb2[__obf_6ffda43d9ecd66ed] = make([]bool, __obf_658024c8e4ec9df7+2*__obf_3252eb05fbf6197a)
	}

	__obf_77b51aa28d7a9d4e.__obf_658024c8e4ec9df7 = __obf_658024c8e4ec9df7 + 2*__obf_3252eb05fbf6197a
	__obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0 = __obf_658024c8e4ec9df7
	__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a = __obf_3252eb05fbf6197a

	return &__obf_77b51aa28d7a9d4e
}

// get returns the module value at (x, y).
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_34a639aea69ea768(__obf_2f6be407475f551f int, __obf_349392c72c5c7113 int) (__obf_4650eaa332d33944 bool) {
	__obf_4650eaa332d33944 = __obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507[__obf_349392c72c5c7113+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a][__obf_2f6be407475f551f+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a]
	return
}

// empty returns true if the module at (x, y) has not been set (to either true
// or false).
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_0c4c4c831e6aadec(__obf_2f6be407475f551f int, __obf_349392c72c5c7113 int) bool {
	return !__obf_77b51aa28d7a9d4e.__obf_5287a97af7a4fcb2[__obf_349392c72c5c7113+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a][__obf_2f6be407475f551f+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a]
}

// numEmptyModules returns the number of empty modules.
//
// Initially numEmptyModules is symbolSize * symbolSize. After every module has
// been set (to either true or false), the number of empty modules is zero.
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_1362019ac7bdc111() int {
	var __obf_979ce6ee8f29807e int
	for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
		for __obf_2f6be407475f551f := 0; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
			if !__obf_77b51aa28d7a9d4e.__obf_5287a97af7a4fcb2[__obf_349392c72c5c7113+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a][__obf_2f6be407475f551f+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a] {
				__obf_979ce6ee8f29807e++
			}
		}
	}

	return __obf_979ce6ee8f29807e
}

// set sets the module at (x, y) to v.
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_fadaba2037d229f1(__obf_2f6be407475f551f int, __obf_349392c72c5c7113 int, __obf_4650eaa332d33944 bool) {
	__obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507[__obf_349392c72c5c7113+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a][__obf_2f6be407475f551f+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a] = __obf_4650eaa332d33944
	__obf_77b51aa28d7a9d4e.__obf_5287a97af7a4fcb2[__obf_349392c72c5c7113+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a][__obf_2f6be407475f551f+__obf_77b51aa28d7a9d4e.__obf_3252eb05fbf6197a] = true
}

// set2dPattern sets a 2D array of modules, starting at (x, y).
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_5614fe3ef4f76bac(__obf_2f6be407475f551f int, __obf_349392c72c5c7113 int, __obf_4650eaa332d33944 [][]bool) {
	for __obf_b09a06cc5f53e0ed, __obf_1505718a25845425 := range __obf_4650eaa332d33944 {
		for __obf_6ffda43d9ecd66ed, __obf_895b3fc3ae24b62b := range __obf_1505718a25845425 {
			__obf_77b51aa28d7a9d4e.__obf_fadaba2037d229f1(__obf_2f6be407475f551f+__obf_6ffda43d9ecd66ed, __obf_349392c72c5c7113+__obf_b09a06cc5f53e0ed, __obf_895b3fc3ae24b62b)
		}
	}
}

// bitmap returns the entire symbol, including the quiet zone.
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_86e6f11fca3fe468() [][]bool {
	__obf_96e599f6f35e9507 := make([][]bool, len(__obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507))

	for __obf_6ffda43d9ecd66ed := range __obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507 {
		__obf_96e599f6f35e9507[__obf_6ffda43d9ecd66ed] = __obf_77b51aa28d7a9d4e.__obf_96e599f6f35e9507[__obf_6ffda43d9ecd66ed][:]
	}

	return __obf_96e599f6f35e9507
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
	__obf_e03a673b0bf260f8 = 3
	__obf_6bd28ac267c9bc6d = 3
	__obf_8b89de88158f720f = 40
	__obf_15ecf9c7708b64d7 = 10
)

// penaltyScore returns the penalty score of the symbol. The penalty score
// consists of the sum of the four individual penalty types.
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_2ce595e50aac84df() int {
	return __obf_77b51aa28d7a9d4e.__obf_f2d540fb27dfff9b() + __obf_77b51aa28d7a9d4e.__obf_08de380577e20124() + __obf_77b51aa28d7a9d4e.__obf_24d834c5361509ee() + __obf_77b51aa28d7a9d4e.__obf_2e21e5f612ff1009()
}

// penalty1 returns the penalty score for "adjacent modules in row/column with
// same colour".
//
// The numbers of adjacent matching modules and scores are:
// 0-5: score = 0
// 6+ : score = penaltyWeight1 + (numAdjacentModules - 5)
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_f2d540fb27dfff9b() int {
	__obf_e751bf876b38498b := 0

	for __obf_2f6be407475f551f := 0; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
		__obf_b774025f0a899f27 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, 0)
		__obf_979ce6ee8f29807e := 1

		for __obf_349392c72c5c7113 := 1; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
			__obf_4650eaa332d33944 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113)

			if __obf_4650eaa332d33944 != __obf_b774025f0a899f27 {
				__obf_979ce6ee8f29807e = 1
				__obf_b774025f0a899f27 = __obf_4650eaa332d33944
			} else {
				__obf_979ce6ee8f29807e++
				if __obf_979ce6ee8f29807e == 6 {
					__obf_e751bf876b38498b += __obf_e03a673b0bf260f8 + 1
				} else if __obf_979ce6ee8f29807e > 6 {
					__obf_e751bf876b38498b++
				}
			}
		}
	}

	for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
		__obf_b774025f0a899f27 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(0, __obf_349392c72c5c7113)
		__obf_979ce6ee8f29807e := 1

		for __obf_2f6be407475f551f := 1; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
			__obf_4650eaa332d33944 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113)

			if __obf_4650eaa332d33944 != __obf_b774025f0a899f27 {
				__obf_979ce6ee8f29807e = 1
				__obf_b774025f0a899f27 = __obf_4650eaa332d33944
			} else {
				__obf_979ce6ee8f29807e++
				if __obf_979ce6ee8f29807e == 6 {
					__obf_e751bf876b38498b += __obf_e03a673b0bf260f8 + 1
				} else if __obf_979ce6ee8f29807e > 6 {
					__obf_e751bf876b38498b++
				}
			}
		}
	}

	return __obf_e751bf876b38498b
}

// penalty2 returns the penalty score for "block of modules in the same colour".
//
// m*n: score = penaltyWeight2 * (m-1) * (n-1).
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_08de380577e20124() int {
	__obf_e751bf876b38498b := 0

	for __obf_349392c72c5c7113 := 1; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
		for __obf_2f6be407475f551f := 1; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
			__obf_27bc64a4619b64cf := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f-1, __obf_349392c72c5c7113-1)
			__obf_84ec1034564d6786 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113-1)
			__obf_fd0b11fb10448dc2 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f-1, __obf_349392c72c5c7113)
			__obf_d2c2796b22ac552c := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113)

			if __obf_d2c2796b22ac552c == __obf_fd0b11fb10448dc2 && __obf_d2c2796b22ac552c == __obf_84ec1034564d6786 && __obf_d2c2796b22ac552c == __obf_27bc64a4619b64cf {
				__obf_e751bf876b38498b++
			}
		}
	}

	return __obf_e751bf876b38498b * __obf_6bd28ac267c9bc6d
}

// penalty3 returns the penalty score for "1:1:3:1:1 ratio
// (dark:light:dark:light:dark) pattern in row/column, preceded or followed by
// light area 4 modules wide".
//
// Existence of the pattern scores penaltyWeight3.
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_24d834c5361509ee() int {
	__obf_e751bf876b38498b := 0

	for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
		var __obf_20ac96e15cea2c2b int16 = 0x00

		for __obf_2f6be407475f551f := 0; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
			__obf_20ac96e15cea2c2b <<= 1
			if __obf_4650eaa332d33944 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113); __obf_4650eaa332d33944 {
				__obf_20ac96e15cea2c2b |= 1
			}

			switch __obf_20ac96e15cea2c2b & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_e751bf876b38498b += __obf_8b89de88158f720f
				__obf_20ac96e15cea2c2b = 0xFF
			default:
				if __obf_2f6be407475f551f == __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0-1 && (__obf_20ac96e15cea2c2b&0x7f) == 0x5d {
					__obf_e751bf876b38498b += __obf_8b89de88158f720f
					__obf_20ac96e15cea2c2b = 0xFF
				}
			}
		}
	}

	for __obf_2f6be407475f551f := 0; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
		var __obf_20ac96e15cea2c2b int16 = 0x00

		for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
			__obf_20ac96e15cea2c2b <<= 1
			if __obf_4650eaa332d33944 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113); __obf_4650eaa332d33944 {
				__obf_20ac96e15cea2c2b |= 1
			}

			switch __obf_20ac96e15cea2c2b & 0x7ff {
			// 0b000 0101 1101 or 0b10111010000
			// 0x05d           or 0x5d0
			case 0x05d, 0x5d0:
				__obf_e751bf876b38498b += __obf_8b89de88158f720f
				__obf_20ac96e15cea2c2b = 0xFF
			default:
				if __obf_349392c72c5c7113 == __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0-1 && (__obf_20ac96e15cea2c2b&0x7f) == 0x5d {
					__obf_e751bf876b38498b += __obf_8b89de88158f720f
					__obf_20ac96e15cea2c2b = 0xFF
				}
			}
		}
	}

	return __obf_e751bf876b38498b
}

// penalty4 returns the penalty score...
func (__obf_77b51aa28d7a9d4e *__obf_dbf6f541061eeeb5) __obf_2e21e5f612ff1009() int {
	__obf_b02bc3693a8362e9 := __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0 * __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0
	__obf_c55121d0a6c5c64f := 0

	for __obf_2f6be407475f551f := 0; __obf_2f6be407475f551f < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_2f6be407475f551f++ {
		for __obf_349392c72c5c7113 := 0; __obf_349392c72c5c7113 < __obf_77b51aa28d7a9d4e.__obf_dc6945971f3283f0; __obf_349392c72c5c7113++ {
			if __obf_4650eaa332d33944 := __obf_77b51aa28d7a9d4e.__obf_34a639aea69ea768(__obf_2f6be407475f551f, __obf_349392c72c5c7113); __obf_4650eaa332d33944 {
				__obf_c55121d0a6c5c64f++
			}
		}
	}

	__obf_9a0980561be2d01a := __obf_b02bc3693a8362e9/2 - __obf_c55121d0a6c5c64f
	if __obf_9a0980561be2d01a < 0 {
		__obf_9a0980561be2d01a *= -1
	}

	return __obf_15ecf9c7708b64d7 * (__obf_9a0980561be2d01a / (__obf_b02bc3693a8362e9 / 20))
}
