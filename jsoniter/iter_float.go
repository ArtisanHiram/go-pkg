package __obf_703d23ba520c3295

import (
	"encoding/json"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

var __obf_30d25ab114970748 []int8

const __obf_0712b1d6d67b066b = int8(-1)
const __obf_a17e534b31a849ba = int8(-2)
const __obf_7dcae25e468b67af = int8(-3)

func init() {
	__obf_30d25ab114970748 = make([]int8, 256)
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < len(__obf_30d25ab114970748); __obf_b0a5d2bd48690f1d++ {
		__obf_30d25ab114970748[__obf_b0a5d2bd48690f1d] = __obf_0712b1d6d67b066b
	}
	for __obf_b0a5d2bd48690f1d := int8('0'); __obf_b0a5d2bd48690f1d <= int8('9'); __obf_b0a5d2bd48690f1d++ {
		__obf_30d25ab114970748[__obf_b0a5d2bd48690f1d] = __obf_b0a5d2bd48690f1d - int8('0')
	}
	__obf_30d25ab114970748[','] = __obf_a17e534b31a849ba
	__obf_30d25ab114970748[']'] = __obf_a17e534b31a849ba
	__obf_30d25ab114970748['}'] = __obf_a17e534b31a849ba
	__obf_30d25ab114970748[' '] = __obf_a17e534b31a849ba
	__obf_30d25ab114970748['\t'] = __obf_a17e534b31a849ba
	__obf_30d25ab114970748['\n'] = __obf_a17e534b31a849ba
	__obf_30d25ab114970748['.'] = __obf_7dcae25e468b67af
}

// ReadBigFloat read big.Float
func (__obf_47edab4c16a2d88d *Iterator) ReadBigFloat() (__obf_551cbec38242ce0c *big.Float) {
	__obf_44b48c26051f8033 := __obf_47edab4c16a2d88d.__obf_4c009361bc8be406()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		return nil
	}
	__obf_b8a26eb8f0c215e8 := 64
	if len(__obf_44b48c26051f8033) > __obf_b8a26eb8f0c215e8 {
		__obf_b8a26eb8f0c215e8 = len(__obf_44b48c26051f8033)
	}
	__obf_b924538fffe5fd64, _, __obf_e56742d6e52f642e := big.ParseFloat(__obf_44b48c26051f8033, 10, uint(__obf_b8a26eb8f0c215e8), big.ToZero)
	if __obf_e56742d6e52f642e != nil {
		__obf_47edab4c16a2d88d.
			Error = __obf_e56742d6e52f642e
		return nil
	}
	return __obf_b924538fffe5fd64
}

// ReadBigInt read big.Int
func (__obf_47edab4c16a2d88d *Iterator) ReadBigInt() (__obf_551cbec38242ce0c *big.Int) {
	__obf_44b48c26051f8033 := __obf_47edab4c16a2d88d.__obf_4c009361bc8be406()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		return nil
	}
	__obf_551cbec38242ce0c = big.NewInt(0)
	var __obf_a92d797d86808a2e bool
	__obf_551cbec38242ce0c, __obf_a92d797d86808a2e = __obf_551cbec38242ce0c.SetString(__obf_44b48c26051f8033, 10)
	if !__obf_a92d797d86808a2e {
		__obf_47edab4c16a2d88d.
			ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return __obf_551cbec38242ce0c
}

// ReadFloat32 read float32
func (__obf_47edab4c16a2d88d *Iterator) ReadFloat32() (__obf_551cbec38242ce0c float32) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '-' {
		return -__obf_47edab4c16a2d88d.__obf_6ad7452a5b0c6b03()
	}
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	return __obf_47edab4c16a2d88d.__obf_6ad7452a5b0c6b03()
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_6ad7452a5b0c6b03() (__obf_551cbec38242ce0c float32) {
	__obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.
		// first char
		__obf_da6aa1cfbd772c11

	if __obf_b0a5d2bd48690f1d == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
		return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
	}
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
	__obf_b0a5d2bd48690f1d++
	__obf_562d42737cc3cc4f := __obf_30d25ab114970748[__obf_bd08f5161fff294a]
	switch __obf_562d42737cc3cc4f {
	case __obf_0712b1d6d67b066b:
		return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
	case __obf_a17e534b31a849ba:
		__obf_47edab4c16a2d88d.
			ReportError("readFloat32", "empty number")
		return
	case __obf_7dcae25e468b67af:
		__obf_47edab4c16a2d88d.
			ReportError("readFloat32", "leading dot is invalid")
		return
	case 0:
		if __obf_b0a5d2bd48690f1d == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
			return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
		}
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		switch __obf_bd08f5161fff294a {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_47edab4c16a2d88d.
				ReportError("readFloat32", "leading zero is invalid")
			return
		}
	}
	__obf_ccb7d8cb07a5350f := uint64(__obf_562d42737cc3cc4f)
__obf_6ec2874ac5cf5e9d: // chars before dot

	for ; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		__obf_562d42737cc3cc4f := __obf_30d25ab114970748[__obf_bd08f5161fff294a]
		switch __obf_562d42737cc3cc4f {
		case __obf_0712b1d6d67b066b:
			return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
		case __obf_a17e534b31a849ba:
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			return float32(__obf_ccb7d8cb07a5350f)
		case __obf_7dcae25e468b67af:
			break __obf_6ec2874ac5cf5e9d
		}
		if __obf_ccb7d8cb07a5350f > __obf_470c0669890e2df4 {
			return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
		}
		__obf_ccb7d8cb07a5350f = (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint64(__obf_562d42737cc3cc4f) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_bd08f5161fff294a == '.' {
		__obf_b0a5d2bd48690f1d++
		__obf_4fc10775493a992f := 0
		if __obf_b0a5d2bd48690f1d == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
			return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
		}
		for ; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			__obf_562d42737cc3cc4f := __obf_30d25ab114970748[__obf_bd08f5161fff294a]
			switch __obf_562d42737cc3cc4f {
			case __obf_a17e534b31a849ba:
				if __obf_4fc10775493a992f > 0 && __obf_4fc10775493a992f < len(__obf_ae9fbf6f0c88dcb9) {
					__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
					return float32(float64(__obf_ccb7d8cb07a5350f) / float64(__obf_ae9fbf6f0c88dcb9[__obf_4fc10775493a992f]))
				}
				// too many decimal places
				return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
			case __obf_0712b1d6d67b066b, __obf_7dcae25e468b67af:
				return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
			}
			__obf_4fc10775493a992f++
			if __obf_ccb7d8cb07a5350f > __obf_470c0669890e2df4 {
				return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
			}
			__obf_ccb7d8cb07a5350f = (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint64(__obf_562d42737cc3cc4f)
		}
	}
	return __obf_47edab4c16a2d88d.__obf_31a4a6ca0a38a835()
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_4c009361bc8be406() (__obf_551cbec38242ce0c string) {
	__obf_d98d4c6039b4a3c2 := [16]byte{}
	__obf_44b48c26051f8033 := __obf_d98d4c6039b4a3c2[0:0]
__obf_99421cae95cacb4d:
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			switch __obf_bd08f5161fff294a {
			case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				__obf_44b48c26051f8033 = append(__obf_44b48c26051f8033, __obf_bd08f5161fff294a)
				continue
			default:
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
				break __obf_99421cae95cacb4d
			}
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			break
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		return
	}
	if len(__obf_44b48c26051f8033) == 0 {
		__obf_47edab4c16a2d88d.
			ReportError("readNumberAsString", "invalid number")
	}
	return *(*string)(unsafe.Pointer(&__obf_44b48c26051f8033))
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_31a4a6ca0a38a835() (__obf_551cbec38242ce0c float32) {
	__obf_44b48c26051f8033 := __obf_47edab4c16a2d88d.__obf_4c009361bc8be406()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		return
	}
	__obf_2125cfb29515da7d := __obf_e48bcf616b723ef2(__obf_44b48c26051f8033)
	if __obf_2125cfb29515da7d != "" {
		__obf_47edab4c16a2d88d.
			ReportError("readFloat32SlowPath", __obf_2125cfb29515da7d)
		return
	}
	__obf_b924538fffe5fd64, __obf_e56742d6e52f642e := strconv.ParseFloat(__obf_44b48c26051f8033, 32)
	if __obf_e56742d6e52f642e != nil {
		__obf_47edab4c16a2d88d.
			Error = __obf_e56742d6e52f642e
		return
	}
	return float32(__obf_b924538fffe5fd64)
}

// ReadFloat64 read float64
func (__obf_47edab4c16a2d88d *Iterator) ReadFloat64() (__obf_551cbec38242ce0c float64) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '-' {
		return -__obf_47edab4c16a2d88d.__obf_14dd23fe04c7dbb9()
	}
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	return __obf_47edab4c16a2d88d.__obf_14dd23fe04c7dbb9()
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_14dd23fe04c7dbb9() (__obf_551cbec38242ce0c float64) {
	__obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.
		// first char
		__obf_da6aa1cfbd772c11

	if __obf_b0a5d2bd48690f1d == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
		return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
	}
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
	__obf_b0a5d2bd48690f1d++
	__obf_562d42737cc3cc4f := __obf_30d25ab114970748[__obf_bd08f5161fff294a]
	switch __obf_562d42737cc3cc4f {
	case __obf_0712b1d6d67b066b:
		return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
	case __obf_a17e534b31a849ba:
		__obf_47edab4c16a2d88d.
			ReportError("readFloat64", "empty number")
		return
	case __obf_7dcae25e468b67af:
		__obf_47edab4c16a2d88d.
			ReportError("readFloat64", "leading dot is invalid")
		return
	case 0:
		if __obf_b0a5d2bd48690f1d == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
			return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
		}
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		switch __obf_bd08f5161fff294a {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_47edab4c16a2d88d.
				ReportError("readFloat64", "leading zero is invalid")
			return
		}
	}
	__obf_ccb7d8cb07a5350f := uint64(__obf_562d42737cc3cc4f)
__obf_6ec2874ac5cf5e9d: // chars before dot

	for ; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		__obf_562d42737cc3cc4f := __obf_30d25ab114970748[__obf_bd08f5161fff294a]
		switch __obf_562d42737cc3cc4f {
		case __obf_0712b1d6d67b066b:
			return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
		case __obf_a17e534b31a849ba:
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
			return float64(__obf_ccb7d8cb07a5350f)
		case __obf_7dcae25e468b67af:
			break __obf_6ec2874ac5cf5e9d
		}
		if __obf_ccb7d8cb07a5350f > __obf_470c0669890e2df4 {
			return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
		}
		__obf_ccb7d8cb07a5350f = (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint64(__obf_562d42737cc3cc4f) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_bd08f5161fff294a == '.' {
		__obf_b0a5d2bd48690f1d++
		__obf_4fc10775493a992f := 0
		if __obf_b0a5d2bd48690f1d == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
			return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
		}
		for ; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			__obf_562d42737cc3cc4f := __obf_30d25ab114970748[__obf_bd08f5161fff294a]
			switch __obf_562d42737cc3cc4f {
			case __obf_a17e534b31a849ba:
				if __obf_4fc10775493a992f > 0 && __obf_4fc10775493a992f < len(__obf_ae9fbf6f0c88dcb9) {
					__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
					return float64(__obf_ccb7d8cb07a5350f) / float64(__obf_ae9fbf6f0c88dcb9[__obf_4fc10775493a992f])
				}
				// too many decimal places
				return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
			case __obf_0712b1d6d67b066b, __obf_7dcae25e468b67af:
				return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
			}
			__obf_4fc10775493a992f++
			if __obf_ccb7d8cb07a5350f > __obf_470c0669890e2df4 {
				return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
			}
			__obf_ccb7d8cb07a5350f = (__obf_ccb7d8cb07a5350f << 3) + (__obf_ccb7d8cb07a5350f << 1) + uint64(__obf_562d42737cc3cc4f)
			if __obf_ccb7d8cb07a5350f > __obf_df0c999320bb22b2 {
				return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
			}
		}
	}
	return __obf_47edab4c16a2d88d.__obf_8d7a1ceb64398ebb()
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_8d7a1ceb64398ebb() (__obf_551cbec38242ce0c float64) {
	__obf_44b48c26051f8033 := __obf_47edab4c16a2d88d.__obf_4c009361bc8be406()
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		return
	}
	__obf_2125cfb29515da7d := __obf_e48bcf616b723ef2(__obf_44b48c26051f8033)
	if __obf_2125cfb29515da7d != "" {
		__obf_47edab4c16a2d88d.
			ReportError("readFloat64SlowPath", __obf_2125cfb29515da7d)
		return
	}
	__obf_b924538fffe5fd64, __obf_e56742d6e52f642e := strconv.ParseFloat(__obf_44b48c26051f8033, 64)
	if __obf_e56742d6e52f642e != nil {
		__obf_47edab4c16a2d88d.
			Error = __obf_e56742d6e52f642e
		return
	}
	return __obf_b924538fffe5fd64
}

func __obf_e48bcf616b723ef2(__obf_44b48c26051f8033 string) string {
	// strconv.ParseFloat is not validating `1.` or `1.e1`
	if len(__obf_44b48c26051f8033) == 0 {
		return "empty number"
	}
	if __obf_44b48c26051f8033[0] == '-' {
		return "-- is not valid"
	}
	__obf_ab1b5560fc0d0160 := strings.IndexByte(__obf_44b48c26051f8033, '.')
	if __obf_ab1b5560fc0d0160 != -1 {
		if __obf_ab1b5560fc0d0160 == len(__obf_44b48c26051f8033)-1 {
			return "dot can not be last character"
		}
		switch __obf_44b48c26051f8033[__obf_ab1b5560fc0d0160+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return "missing digit after dot"
		}
	}
	return ""
}

// ReadNumber read json.Number
func (__obf_47edab4c16a2d88d *Iterator) ReadNumber() (__obf_551cbec38242ce0c json.Number) {
	return json.Number(__obf_47edab4c16a2d88d.__obf_4c009361bc8be406())
}
