package __obf_030d39f7a8de96c6

import (
	"encoding/json"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

var __obf_6aaf9584a5258a3a []int8

const __obf_78b317d8e70f5d5a = int8(-1)
const __obf_d02f76ea6caf74af = int8(-2)
const __obf_426c4818e2baf97f = int8(-3)

func init() {
	__obf_6aaf9584a5258a3a = make([]int8, 256)
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < len(__obf_6aaf9584a5258a3a); __obf_82c6e05b9d226c58++ {
		__obf_6aaf9584a5258a3a[__obf_82c6e05b9d226c58] = __obf_78b317d8e70f5d5a
	}
	for __obf_82c6e05b9d226c58 := int8('0'); __obf_82c6e05b9d226c58 <= int8('9'); __obf_82c6e05b9d226c58++ {
		__obf_6aaf9584a5258a3a[__obf_82c6e05b9d226c58] = __obf_82c6e05b9d226c58 - int8('0')
	}
	__obf_6aaf9584a5258a3a[','] = __obf_d02f76ea6caf74af
	__obf_6aaf9584a5258a3a[']'] = __obf_d02f76ea6caf74af
	__obf_6aaf9584a5258a3a['}'] = __obf_d02f76ea6caf74af
	__obf_6aaf9584a5258a3a[' '] = __obf_d02f76ea6caf74af
	__obf_6aaf9584a5258a3a['\t'] = __obf_d02f76ea6caf74af
	__obf_6aaf9584a5258a3a['\n'] = __obf_d02f76ea6caf74af
	__obf_6aaf9584a5258a3a['.'] = __obf_426c4818e2baf97f
}

// ReadBigFloat read big.Float
func (__obf_4ab56a99965952e7 *Iterator) ReadBigFloat() (__obf_59c72400ec1a6110 *big.Float) {
	__obf_428c3b4a95725c4a := __obf_4ab56a99965952e7.__obf_0ba08d27498a0d84()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		return nil
	}
	__obf_213839e2d936ae64 := 64
	if len(__obf_428c3b4a95725c4a) > __obf_213839e2d936ae64 {
		__obf_213839e2d936ae64 = len(__obf_428c3b4a95725c4a)
	}
	__obf_efaf2719b44ad8ac, _, __obf_fcc907dd69879566 := big.ParseFloat(__obf_428c3b4a95725c4a, 10, uint(__obf_213839e2d936ae64), big.ToZero)
	if __obf_fcc907dd69879566 != nil {
		__obf_4ab56a99965952e7.
			Error = __obf_fcc907dd69879566
		return nil
	}
	return __obf_efaf2719b44ad8ac
}

// ReadBigInt read big.Int
func (__obf_4ab56a99965952e7 *Iterator) ReadBigInt() (__obf_59c72400ec1a6110 *big.Int) {
	__obf_428c3b4a95725c4a := __obf_4ab56a99965952e7.__obf_0ba08d27498a0d84()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		return nil
	}
	__obf_59c72400ec1a6110 = big.NewInt(0)
	var __obf_8ea174385fbfb807 bool
	__obf_59c72400ec1a6110, __obf_8ea174385fbfb807 = __obf_59c72400ec1a6110.SetString(__obf_428c3b4a95725c4a, 10)
	if !__obf_8ea174385fbfb807 {
		__obf_4ab56a99965952e7.
			ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return __obf_59c72400ec1a6110
}

// ReadFloat32 read float32
func (__obf_4ab56a99965952e7 *Iterator) ReadFloat32() (__obf_59c72400ec1a6110 float32) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '-' {
		return -__obf_4ab56a99965952e7.__obf_5416756ef982b9b6()
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	return __obf_4ab56a99965952e7.__obf_5416756ef982b9b6()
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_5416756ef982b9b6() (__obf_59c72400ec1a6110 float32) {
	__obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.
		// first char
		__obf_5e22d6270d27491f

	if __obf_82c6e05b9d226c58 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
		return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
	}
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
	__obf_82c6e05b9d226c58++
	__obf_0d7ccf72a2e94911 := __obf_6aaf9584a5258a3a[__obf_24309b3b0ff9ba22]
	switch __obf_0d7ccf72a2e94911 {
	case __obf_78b317d8e70f5d5a:
		return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
	case __obf_d02f76ea6caf74af:
		__obf_4ab56a99965952e7.
			ReportError("readFloat32", "empty number")
		return
	case __obf_426c4818e2baf97f:
		__obf_4ab56a99965952e7.
			ReportError("readFloat32", "leading dot is invalid")
		return
	case 0:
		if __obf_82c6e05b9d226c58 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
			return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
		}
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		switch __obf_24309b3b0ff9ba22 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_4ab56a99965952e7.
				ReportError("readFloat32", "leading zero is invalid")
			return
		}
	}
	__obf_b80d5217e1232943 := uint64(__obf_0d7ccf72a2e94911)
__obf_ea2f3f8a33ada2ab: // chars before dot

	for ; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		__obf_0d7ccf72a2e94911 := __obf_6aaf9584a5258a3a[__obf_24309b3b0ff9ba22]
		switch __obf_0d7ccf72a2e94911 {
		case __obf_78b317d8e70f5d5a:
			return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
		case __obf_d02f76ea6caf74af:
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			return float32(__obf_b80d5217e1232943)
		case __obf_426c4818e2baf97f:
			break __obf_ea2f3f8a33ada2ab
		}
		if __obf_b80d5217e1232943 > __obf_92013c25906febd6 {
			return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
		}
		__obf_b80d5217e1232943 = (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint64(__obf_0d7ccf72a2e94911) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_24309b3b0ff9ba22 == '.' {
		__obf_82c6e05b9d226c58++
		__obf_926fcf88b57e320e := 0
		if __obf_82c6e05b9d226c58 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
			return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
		}
		for ; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			__obf_0d7ccf72a2e94911 := __obf_6aaf9584a5258a3a[__obf_24309b3b0ff9ba22]
			switch __obf_0d7ccf72a2e94911 {
			case __obf_d02f76ea6caf74af:
				if __obf_926fcf88b57e320e > 0 && __obf_926fcf88b57e320e < len(__obf_faa2ce64bf143410) {
					__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
					return float32(float64(__obf_b80d5217e1232943) / float64(__obf_faa2ce64bf143410[__obf_926fcf88b57e320e]))
				}
				// too many decimal places
				return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
			case __obf_78b317d8e70f5d5a, __obf_426c4818e2baf97f:
				return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
			}
			__obf_926fcf88b57e320e++
			if __obf_b80d5217e1232943 > __obf_92013c25906febd6 {
				return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
			}
			__obf_b80d5217e1232943 = (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint64(__obf_0d7ccf72a2e94911)
		}
	}
	return __obf_4ab56a99965952e7.__obf_dbfa82d71ff369c6()
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_0ba08d27498a0d84() (__obf_59c72400ec1a6110 string) {
	__obf_1a067b8c29f662c5 := [16]byte{}
	__obf_428c3b4a95725c4a := __obf_1a067b8c29f662c5[0:0]
__obf_c7e3e88740a1507b:
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			switch __obf_24309b3b0ff9ba22 {
			case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				__obf_428c3b4a95725c4a = append(__obf_428c3b4a95725c4a, __obf_24309b3b0ff9ba22)
				continue
			default:
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
				break __obf_c7e3e88740a1507b
			}
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			break
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		return
	}
	if len(__obf_428c3b4a95725c4a) == 0 {
		__obf_4ab56a99965952e7.
			ReportError("readNumberAsString", "invalid number")
	}
	return *(*string)(unsafe.Pointer(&__obf_428c3b4a95725c4a))
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_dbfa82d71ff369c6() (__obf_59c72400ec1a6110 float32) {
	__obf_428c3b4a95725c4a := __obf_4ab56a99965952e7.__obf_0ba08d27498a0d84()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		return
	}
	__obf_a56e644550dbb87f := __obf_0dda531e93ce6cdb(__obf_428c3b4a95725c4a)
	if __obf_a56e644550dbb87f != "" {
		__obf_4ab56a99965952e7.
			ReportError("readFloat32SlowPath", __obf_a56e644550dbb87f)
		return
	}
	__obf_efaf2719b44ad8ac, __obf_fcc907dd69879566 := strconv.ParseFloat(__obf_428c3b4a95725c4a, 32)
	if __obf_fcc907dd69879566 != nil {
		__obf_4ab56a99965952e7.
			Error = __obf_fcc907dd69879566
		return
	}
	return float32(__obf_efaf2719b44ad8ac)
}

// ReadFloat64 read float64
func (__obf_4ab56a99965952e7 *Iterator) ReadFloat64() (__obf_59c72400ec1a6110 float64) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '-' {
		return -__obf_4ab56a99965952e7.__obf_f27d86ec145a76ff()
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	return __obf_4ab56a99965952e7.__obf_f27d86ec145a76ff()
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_f27d86ec145a76ff() (__obf_59c72400ec1a6110 float64) {
	__obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.
		// first char
		__obf_5e22d6270d27491f

	if __obf_82c6e05b9d226c58 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
		return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
	}
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
	__obf_82c6e05b9d226c58++
	__obf_0d7ccf72a2e94911 := __obf_6aaf9584a5258a3a[__obf_24309b3b0ff9ba22]
	switch __obf_0d7ccf72a2e94911 {
	case __obf_78b317d8e70f5d5a:
		return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
	case __obf_d02f76ea6caf74af:
		__obf_4ab56a99965952e7.
			ReportError("readFloat64", "empty number")
		return
	case __obf_426c4818e2baf97f:
		__obf_4ab56a99965952e7.
			ReportError("readFloat64", "leading dot is invalid")
		return
	case 0:
		if __obf_82c6e05b9d226c58 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
			return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
		}
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		switch __obf_24309b3b0ff9ba22 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_4ab56a99965952e7.
				ReportError("readFloat64", "leading zero is invalid")
			return
		}
	}
	__obf_b80d5217e1232943 := uint64(__obf_0d7ccf72a2e94911)
__obf_ea2f3f8a33ada2ab: // chars before dot

	for ; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		__obf_0d7ccf72a2e94911 := __obf_6aaf9584a5258a3a[__obf_24309b3b0ff9ba22]
		switch __obf_0d7ccf72a2e94911 {
		case __obf_78b317d8e70f5d5a:
			return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
		case __obf_d02f76ea6caf74af:
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
			return float64(__obf_b80d5217e1232943)
		case __obf_426c4818e2baf97f:
			break __obf_ea2f3f8a33ada2ab
		}
		if __obf_b80d5217e1232943 > __obf_92013c25906febd6 {
			return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
		}
		__obf_b80d5217e1232943 = (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint64(__obf_0d7ccf72a2e94911) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_24309b3b0ff9ba22 == '.' {
		__obf_82c6e05b9d226c58++
		__obf_926fcf88b57e320e := 0
		if __obf_82c6e05b9d226c58 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
			return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
		}
		for ; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			__obf_0d7ccf72a2e94911 := __obf_6aaf9584a5258a3a[__obf_24309b3b0ff9ba22]
			switch __obf_0d7ccf72a2e94911 {
			case __obf_d02f76ea6caf74af:
				if __obf_926fcf88b57e320e > 0 && __obf_926fcf88b57e320e < len(__obf_faa2ce64bf143410) {
					__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
					return float64(__obf_b80d5217e1232943) / float64(__obf_faa2ce64bf143410[__obf_926fcf88b57e320e])
				}
				// too many decimal places
				return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
			case __obf_78b317d8e70f5d5a, __obf_426c4818e2baf97f:
				return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
			}
			__obf_926fcf88b57e320e++
			if __obf_b80d5217e1232943 > __obf_92013c25906febd6 {
				return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
			}
			__obf_b80d5217e1232943 = (__obf_b80d5217e1232943 << 3) + (__obf_b80d5217e1232943 << 1) + uint64(__obf_0d7ccf72a2e94911)
			if __obf_b80d5217e1232943 > __obf_636d63e04d6021f5 {
				return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
			}
		}
	}
	return __obf_4ab56a99965952e7.__obf_1c4c785ea4a51fd2()
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_1c4c785ea4a51fd2() (__obf_59c72400ec1a6110 float64) {
	__obf_428c3b4a95725c4a := __obf_4ab56a99965952e7.__obf_0ba08d27498a0d84()
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		return
	}
	__obf_a56e644550dbb87f := __obf_0dda531e93ce6cdb(__obf_428c3b4a95725c4a)
	if __obf_a56e644550dbb87f != "" {
		__obf_4ab56a99965952e7.
			ReportError("readFloat64SlowPath", __obf_a56e644550dbb87f)
		return
	}
	__obf_efaf2719b44ad8ac, __obf_fcc907dd69879566 := strconv.ParseFloat(__obf_428c3b4a95725c4a, 64)
	if __obf_fcc907dd69879566 != nil {
		__obf_4ab56a99965952e7.
			Error = __obf_fcc907dd69879566
		return
	}
	return __obf_efaf2719b44ad8ac
}

func __obf_0dda531e93ce6cdb(__obf_428c3b4a95725c4a string) string {
	// strconv.ParseFloat is not validating `1.` or `1.e1`
	if len(__obf_428c3b4a95725c4a) == 0 {
		return "empty number"
	}
	if __obf_428c3b4a95725c4a[0] == '-' {
		return "-- is not valid"
	}
	__obf_a877872d1a11ebdc := strings.IndexByte(__obf_428c3b4a95725c4a, '.')
	if __obf_a877872d1a11ebdc != -1 {
		if __obf_a877872d1a11ebdc == len(__obf_428c3b4a95725c4a)-1 {
			return "dot can not be last character"
		}
		switch __obf_428c3b4a95725c4a[__obf_a877872d1a11ebdc+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return "missing digit after dot"
		}
	}
	return ""
}

// ReadNumber read json.Number
func (__obf_4ab56a99965952e7 *Iterator) ReadNumber() (__obf_59c72400ec1a6110 json.Number) {
	return json.Number(__obf_4ab56a99965952e7.__obf_0ba08d27498a0d84())
}
