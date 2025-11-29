package __obf_5b802ce8d9ba56d6

import (
	"encoding/json"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

var __obf_7779fb98616bb2f3 []int8

const __obf_22cd644f84c3de2e = int8(-1)
const __obf_9408e377f73c3b34 = int8(-2)
const __obf_1b23a538da85bcd3 = int8(-3)

func init() {
	__obf_7779fb98616bb2f3 = make([]int8, 256)
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < len(__obf_7779fb98616bb2f3); __obf_2deec7c38ffb6ae3++ {
		__obf_7779fb98616bb2f3[__obf_2deec7c38ffb6ae3] = __obf_22cd644f84c3de2e
	}
	for __obf_2deec7c38ffb6ae3 := int8('0'); __obf_2deec7c38ffb6ae3 <= int8('9'); __obf_2deec7c38ffb6ae3++ {
		__obf_7779fb98616bb2f3[__obf_2deec7c38ffb6ae3] = __obf_2deec7c38ffb6ae3 - int8('0')
	}
	__obf_7779fb98616bb2f3[','] = __obf_9408e377f73c3b34
	__obf_7779fb98616bb2f3[']'] = __obf_9408e377f73c3b34
	__obf_7779fb98616bb2f3['}'] = __obf_9408e377f73c3b34
	__obf_7779fb98616bb2f3[' '] = __obf_9408e377f73c3b34
	__obf_7779fb98616bb2f3['\t'] = __obf_9408e377f73c3b34
	__obf_7779fb98616bb2f3['\n'] = __obf_9408e377f73c3b34
	__obf_7779fb98616bb2f3['.'] = __obf_1b23a538da85bcd3
}

// ReadBigFloat read big.Float
func (__obf_67008a6a9e5ba828 *Iterator) ReadBigFloat() (__obf_5dabcdfee5097ed6 *big.Float) {
	__obf_12c21b79fa86dcba := __obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		return nil
	}
	__obf_fb7dfd12d2d8d625 := 64
	if len(__obf_12c21b79fa86dcba) > __obf_fb7dfd12d2d8d625 {
		__obf_fb7dfd12d2d8d625 = len(__obf_12c21b79fa86dcba)
	}
	__obf_5406b11e33b9d1d3, _, __obf_6d071d48f3b321ad := big.ParseFloat(__obf_12c21b79fa86dcba, 10, uint(__obf_fb7dfd12d2d8d625), big.ToZero)
	if __obf_6d071d48f3b321ad != nil {
		__obf_67008a6a9e5ba828.
			Error = __obf_6d071d48f3b321ad
		return nil
	}
	return __obf_5406b11e33b9d1d3
}

// ReadBigInt read big.Int
func (__obf_67008a6a9e5ba828 *Iterator) ReadBigInt() (__obf_5dabcdfee5097ed6 *big.Int) {
	__obf_12c21b79fa86dcba := __obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		return nil
	}
	__obf_5dabcdfee5097ed6 = big.NewInt(0)
	var __obf_96ae77751ddfcfbd bool
	__obf_5dabcdfee5097ed6, __obf_96ae77751ddfcfbd = __obf_5dabcdfee5097ed6.SetString(__obf_12c21b79fa86dcba, 10)
	if !__obf_96ae77751ddfcfbd {
		__obf_67008a6a9e5ba828.
			ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return __obf_5dabcdfee5097ed6
}

// ReadFloat32 read float32
func (__obf_67008a6a9e5ba828 *Iterator) ReadFloat32() (__obf_5dabcdfee5097ed6 float32) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '-' {
		return -__obf_67008a6a9e5ba828.__obf_98b730eea56da144()
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	return __obf_67008a6a9e5ba828.__obf_98b730eea56da144()
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_98b730eea56da144() (__obf_5dabcdfee5097ed6 float32) {
	__obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.
		// first char
		__obf_14babd6f9a55bd36

	if __obf_2deec7c38ffb6ae3 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
		return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
	}
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
	__obf_2deec7c38ffb6ae3++
	__obf_0664edf071899d0b := __obf_7779fb98616bb2f3[__obf_dab9baaadfa7c8c2]
	switch __obf_0664edf071899d0b {
	case __obf_22cd644f84c3de2e:
		return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
	case __obf_9408e377f73c3b34:
		__obf_67008a6a9e5ba828.
			ReportError("readFloat32", "empty number")
		return
	case __obf_1b23a538da85bcd3:
		__obf_67008a6a9e5ba828.
			ReportError("readFloat32", "leading dot is invalid")
		return
	case 0:
		if __obf_2deec7c38ffb6ae3 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
			return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
		}
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		switch __obf_dab9baaadfa7c8c2 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_67008a6a9e5ba828.
				ReportError("readFloat32", "leading zero is invalid")
			return
		}
	}
	__obf_c949477a4af2082d := uint64(__obf_0664edf071899d0b)
__obf_baae93481c6c2a4e: // chars before dot

	for ; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		__obf_0664edf071899d0b := __obf_7779fb98616bb2f3[__obf_dab9baaadfa7c8c2]
		switch __obf_0664edf071899d0b {
		case __obf_22cd644f84c3de2e:
			return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
		case __obf_9408e377f73c3b34:
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			return float32(__obf_c949477a4af2082d)
		case __obf_1b23a538da85bcd3:
			break __obf_baae93481c6c2a4e
		}
		if __obf_c949477a4af2082d > __obf_c6c8cab5bfa05789 {
			return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
		}
		__obf_c949477a4af2082d = (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint64(__obf_0664edf071899d0b) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_dab9baaadfa7c8c2 == '.' {
		__obf_2deec7c38ffb6ae3++
		__obf_457ce922f4d9189d := 0
		if __obf_2deec7c38ffb6ae3 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
			return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
		}
		for ; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			__obf_0664edf071899d0b := __obf_7779fb98616bb2f3[__obf_dab9baaadfa7c8c2]
			switch __obf_0664edf071899d0b {
			case __obf_9408e377f73c3b34:
				if __obf_457ce922f4d9189d > 0 && __obf_457ce922f4d9189d < len(__obf_532c8e692981b5b4) {
					__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
					return float32(float64(__obf_c949477a4af2082d) / float64(__obf_532c8e692981b5b4[__obf_457ce922f4d9189d]))
				}
				// too many decimal places
				return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
			case __obf_22cd644f84c3de2e, __obf_1b23a538da85bcd3:
				return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
			}
			__obf_457ce922f4d9189d++
			if __obf_c949477a4af2082d > __obf_c6c8cab5bfa05789 {
				return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
			}
			__obf_c949477a4af2082d = (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint64(__obf_0664edf071899d0b)
		}
	}
	return __obf_67008a6a9e5ba828.__obf_b5a584e054708c57()
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_bad2e5b8bc5a6a05() (__obf_5dabcdfee5097ed6 string) {
	__obf_f81afbe4e7a5b665 := [16]byte{}
	__obf_12c21b79fa86dcba := __obf_f81afbe4e7a5b665[0:0]
__obf_9b3ca22f2b2e6bf2:
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			switch __obf_dab9baaadfa7c8c2 {
			case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				__obf_12c21b79fa86dcba = append(__obf_12c21b79fa86dcba, __obf_dab9baaadfa7c8c2)
				continue
			default:
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
				break __obf_9b3ca22f2b2e6bf2
			}
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			break
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		return
	}
	if len(__obf_12c21b79fa86dcba) == 0 {
		__obf_67008a6a9e5ba828.
			ReportError("readNumberAsString", "invalid number")
	}
	return *(*string)(unsafe.Pointer(&__obf_12c21b79fa86dcba))
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_b5a584e054708c57() (__obf_5dabcdfee5097ed6 float32) {
	__obf_12c21b79fa86dcba := __obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		return
	}
	__obf_ed32c75cb220585d := __obf_ae28afcbb4a23370(__obf_12c21b79fa86dcba)
	if __obf_ed32c75cb220585d != "" {
		__obf_67008a6a9e5ba828.
			ReportError("readFloat32SlowPath", __obf_ed32c75cb220585d)
		return
	}
	__obf_5406b11e33b9d1d3, __obf_6d071d48f3b321ad := strconv.ParseFloat(__obf_12c21b79fa86dcba, 32)
	if __obf_6d071d48f3b321ad != nil {
		__obf_67008a6a9e5ba828.
			Error = __obf_6d071d48f3b321ad
		return
	}
	return float32(__obf_5406b11e33b9d1d3)
}

// ReadFloat64 read float64
func (__obf_67008a6a9e5ba828 *Iterator) ReadFloat64() (__obf_5dabcdfee5097ed6 float64) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '-' {
		return -__obf_67008a6a9e5ba828.__obf_5e10252297812838()
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	return __obf_67008a6a9e5ba828.__obf_5e10252297812838()
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_5e10252297812838() (__obf_5dabcdfee5097ed6 float64) {
	__obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.
		// first char
		__obf_14babd6f9a55bd36

	if __obf_2deec7c38ffb6ae3 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
		return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
	}
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
	__obf_2deec7c38ffb6ae3++
	__obf_0664edf071899d0b := __obf_7779fb98616bb2f3[__obf_dab9baaadfa7c8c2]
	switch __obf_0664edf071899d0b {
	case __obf_22cd644f84c3de2e:
		return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
	case __obf_9408e377f73c3b34:
		__obf_67008a6a9e5ba828.
			ReportError("readFloat64", "empty number")
		return
	case __obf_1b23a538da85bcd3:
		__obf_67008a6a9e5ba828.
			ReportError("readFloat64", "leading dot is invalid")
		return
	case 0:
		if __obf_2deec7c38ffb6ae3 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
			return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
		}
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		switch __obf_dab9baaadfa7c8c2 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_67008a6a9e5ba828.
				ReportError("readFloat64", "leading zero is invalid")
			return
		}
	}
	__obf_c949477a4af2082d := uint64(__obf_0664edf071899d0b)
__obf_baae93481c6c2a4e: // chars before dot

	for ; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		__obf_0664edf071899d0b := __obf_7779fb98616bb2f3[__obf_dab9baaadfa7c8c2]
		switch __obf_0664edf071899d0b {
		case __obf_22cd644f84c3de2e:
			return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
		case __obf_9408e377f73c3b34:
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
			return float64(__obf_c949477a4af2082d)
		case __obf_1b23a538da85bcd3:
			break __obf_baae93481c6c2a4e
		}
		if __obf_c949477a4af2082d > __obf_c6c8cab5bfa05789 {
			return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
		}
		__obf_c949477a4af2082d = (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint64(__obf_0664edf071899d0b) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_dab9baaadfa7c8c2 == '.' {
		__obf_2deec7c38ffb6ae3++
		__obf_457ce922f4d9189d := 0
		if __obf_2deec7c38ffb6ae3 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
			return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
		}
		for ; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			__obf_0664edf071899d0b := __obf_7779fb98616bb2f3[__obf_dab9baaadfa7c8c2]
			switch __obf_0664edf071899d0b {
			case __obf_9408e377f73c3b34:
				if __obf_457ce922f4d9189d > 0 && __obf_457ce922f4d9189d < len(__obf_532c8e692981b5b4) {
					__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
					return float64(__obf_c949477a4af2082d) / float64(__obf_532c8e692981b5b4[__obf_457ce922f4d9189d])
				}
				// too many decimal places
				return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
			case __obf_22cd644f84c3de2e, __obf_1b23a538da85bcd3:
				return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
			}
			__obf_457ce922f4d9189d++
			if __obf_c949477a4af2082d > __obf_c6c8cab5bfa05789 {
				return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
			}
			__obf_c949477a4af2082d = (__obf_c949477a4af2082d << 3) + (__obf_c949477a4af2082d << 1) + uint64(__obf_0664edf071899d0b)
			if __obf_c949477a4af2082d > __obf_d004373e847514ce {
				return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
			}
		}
	}
	return __obf_67008a6a9e5ba828.__obf_02b92fd2de11a663()
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_02b92fd2de11a663() (__obf_5dabcdfee5097ed6 float64) {
	__obf_12c21b79fa86dcba := __obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05()
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		return
	}
	__obf_ed32c75cb220585d := __obf_ae28afcbb4a23370(__obf_12c21b79fa86dcba)
	if __obf_ed32c75cb220585d != "" {
		__obf_67008a6a9e5ba828.
			ReportError("readFloat64SlowPath", __obf_ed32c75cb220585d)
		return
	}
	__obf_5406b11e33b9d1d3, __obf_6d071d48f3b321ad := strconv.ParseFloat(__obf_12c21b79fa86dcba, 64)
	if __obf_6d071d48f3b321ad != nil {
		__obf_67008a6a9e5ba828.
			Error = __obf_6d071d48f3b321ad
		return
	}
	return __obf_5406b11e33b9d1d3
}

func __obf_ae28afcbb4a23370(__obf_12c21b79fa86dcba string) string {
	// strconv.ParseFloat is not validating `1.` or `1.e1`
	if len(__obf_12c21b79fa86dcba) == 0 {
		return "empty number"
	}
	if __obf_12c21b79fa86dcba[0] == '-' {
		return "-- is not valid"
	}
	__obf_c304d59aeb22a4f2 := strings.IndexByte(__obf_12c21b79fa86dcba, '.')
	if __obf_c304d59aeb22a4f2 != -1 {
		if __obf_c304d59aeb22a4f2 == len(__obf_12c21b79fa86dcba)-1 {
			return "dot can not be last character"
		}
		switch __obf_12c21b79fa86dcba[__obf_c304d59aeb22a4f2+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return "missing digit after dot"
		}
	}
	return ""
}

// ReadNumber read json.Number
func (__obf_67008a6a9e5ba828 *Iterator) ReadNumber() (__obf_5dabcdfee5097ed6 json.Number) {
	return json.Number(__obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05())
}
