package __obf_c3cd04a15c56f32f

import (
	"encoding/json"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

var __obf_317d0145726c1931 []int8

const __obf_234bd47da6a58289 = int8(-1)
const __obf_06719ac845597e38 = int8(-2)
const __obf_0e8edd179006cfa4 = int8(-3)

func init() {
	__obf_317d0145726c1931 = make([]int8, 256)
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < len(__obf_317d0145726c1931); __obf_28d099df85f083a8++ {
		__obf_317d0145726c1931[__obf_28d099df85f083a8] = __obf_234bd47da6a58289
	}
	for __obf_28d099df85f083a8 := int8('0'); __obf_28d099df85f083a8 <= int8('9'); __obf_28d099df85f083a8++ {
		__obf_317d0145726c1931[__obf_28d099df85f083a8] = __obf_28d099df85f083a8 - int8('0')
	}
	__obf_317d0145726c1931[','] = __obf_06719ac845597e38
	__obf_317d0145726c1931[']'] = __obf_06719ac845597e38
	__obf_317d0145726c1931['}'] = __obf_06719ac845597e38
	__obf_317d0145726c1931[' '] = __obf_06719ac845597e38
	__obf_317d0145726c1931['\t'] = __obf_06719ac845597e38
	__obf_317d0145726c1931['\n'] = __obf_06719ac845597e38
	__obf_317d0145726c1931['.'] = __obf_0e8edd179006cfa4
}

// ReadBigFloat read big.Float
func (__obf_edd9fe48d39445e4 *Iterator) ReadBigFloat() (__obf_316af79472661247 *big.Float) {
	__obf_2d944450d48e5810 := __obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		return nil
	}
	__obf_96e2888e59ab7d1d := 64
	if len(__obf_2d944450d48e5810) > __obf_96e2888e59ab7d1d {
		__obf_96e2888e59ab7d1d = len(__obf_2d944450d48e5810)
	}
	__obf_d59813f23ad740a8, _, __obf_5966ecc5edb9b17e := big.ParseFloat(__obf_2d944450d48e5810, 10, uint(__obf_96e2888e59ab7d1d), big.ToZero)
	if __obf_5966ecc5edb9b17e != nil {
		__obf_edd9fe48d39445e4.
			Error = __obf_5966ecc5edb9b17e
		return nil
	}
	return __obf_d59813f23ad740a8
}

// ReadBigInt read big.Int
func (__obf_edd9fe48d39445e4 *Iterator) ReadBigInt() (__obf_316af79472661247 *big.Int) {
	__obf_2d944450d48e5810 := __obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		return nil
	}
	__obf_316af79472661247 = big.NewInt(0)
	var __obf_e0847f682ce51fc4 bool
	__obf_316af79472661247, __obf_e0847f682ce51fc4 = __obf_316af79472661247.SetString(__obf_2d944450d48e5810, 10)
	if !__obf_e0847f682ce51fc4 {
		__obf_edd9fe48d39445e4.
			ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return __obf_316af79472661247
}

// ReadFloat32 read float32
func (__obf_edd9fe48d39445e4 *Iterator) ReadFloat32() (__obf_316af79472661247 float32) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '-' {
		return -__obf_edd9fe48d39445e4.__obf_832656a8a832920d()
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	return __obf_edd9fe48d39445e4.__obf_832656a8a832920d()
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_832656a8a832920d() (__obf_316af79472661247 float32) {
	__obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.
		// first char
		__obf_edd3c3885447229b

	if __obf_28d099df85f083a8 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
		return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
	}
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
	__obf_28d099df85f083a8++
	__obf_5e075812c416735d := __obf_317d0145726c1931[__obf_0c1bc1e511a43120]
	switch __obf_5e075812c416735d {
	case __obf_234bd47da6a58289:
		return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
	case __obf_06719ac845597e38:
		__obf_edd9fe48d39445e4.
			ReportError("readFloat32", "empty number")
		return
	case __obf_0e8edd179006cfa4:
		__obf_edd9fe48d39445e4.
			ReportError("readFloat32", "leading dot is invalid")
		return
	case 0:
		if __obf_28d099df85f083a8 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
			return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
		}
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		switch __obf_0c1bc1e511a43120 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_edd9fe48d39445e4.
				ReportError("readFloat32", "leading zero is invalid")
			return
		}
	}
	__obf_1a534fe904baaf4a := uint64(__obf_5e075812c416735d)
__obf_43eb439dd24b9446: // chars before dot

	for ; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		__obf_5e075812c416735d := __obf_317d0145726c1931[__obf_0c1bc1e511a43120]
		switch __obf_5e075812c416735d {
		case __obf_234bd47da6a58289:
			return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
		case __obf_06719ac845597e38:
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			return float32(__obf_1a534fe904baaf4a)
		case __obf_0e8edd179006cfa4:
			break __obf_43eb439dd24b9446
		}
		if __obf_1a534fe904baaf4a > __obf_61c0bbfc8fa559d6 {
			return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
		}
		__obf_1a534fe904baaf4a = (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint64(__obf_5e075812c416735d) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_0c1bc1e511a43120 == '.' {
		__obf_28d099df85f083a8++
		__obf_1d493182e4208194 := 0
		if __obf_28d099df85f083a8 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
			return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
		}
		for ; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			__obf_5e075812c416735d := __obf_317d0145726c1931[__obf_0c1bc1e511a43120]
			switch __obf_5e075812c416735d {
			case __obf_06719ac845597e38:
				if __obf_1d493182e4208194 > 0 && __obf_1d493182e4208194 < len(__obf_cef54dda9534fb7d) {
					__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
					return float32(float64(__obf_1a534fe904baaf4a) / float64(__obf_cef54dda9534fb7d[__obf_1d493182e4208194]))
				}
				// too many decimal places
				return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
			case __obf_234bd47da6a58289, __obf_0e8edd179006cfa4:
				return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
			}
			__obf_1d493182e4208194++
			if __obf_1a534fe904baaf4a > __obf_61c0bbfc8fa559d6 {
				return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
			}
			__obf_1a534fe904baaf4a = (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint64(__obf_5e075812c416735d)
		}
	}
	return __obf_edd9fe48d39445e4.__obf_6ec8653ebbd6c983()
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_8ccccffe17f4c818() (__obf_316af79472661247 string) {
	__obf_01a3dc10125ff22a := [16]byte{}
	__obf_2d944450d48e5810 := __obf_01a3dc10125ff22a[0:0]
__obf_694aa9a8197b089c:
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			switch __obf_0c1bc1e511a43120 {
			case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				__obf_2d944450d48e5810 = append(__obf_2d944450d48e5810, __obf_0c1bc1e511a43120)
				continue
			default:
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
				break __obf_694aa9a8197b089c
			}
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			break
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		return
	}
	if len(__obf_2d944450d48e5810) == 0 {
		__obf_edd9fe48d39445e4.
			ReportError("readNumberAsString", "invalid number")
	}
	return *(*string)(unsafe.Pointer(&__obf_2d944450d48e5810))
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_6ec8653ebbd6c983() (__obf_316af79472661247 float32) {
	__obf_2d944450d48e5810 := __obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		return
	}
	__obf_759e6867bb34418a := __obf_a9febc73c46e733d(__obf_2d944450d48e5810)
	if __obf_759e6867bb34418a != "" {
		__obf_edd9fe48d39445e4.
			ReportError("readFloat32SlowPath", __obf_759e6867bb34418a)
		return
	}
	__obf_d59813f23ad740a8, __obf_5966ecc5edb9b17e := strconv.ParseFloat(__obf_2d944450d48e5810, 32)
	if __obf_5966ecc5edb9b17e != nil {
		__obf_edd9fe48d39445e4.
			Error = __obf_5966ecc5edb9b17e
		return
	}
	return float32(__obf_d59813f23ad740a8)
}

// ReadFloat64 read float64
func (__obf_edd9fe48d39445e4 *Iterator) ReadFloat64() (__obf_316af79472661247 float64) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '-' {
		return -__obf_edd9fe48d39445e4.__obf_0b6fa5afe92faa37()
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	return __obf_edd9fe48d39445e4.__obf_0b6fa5afe92faa37()
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_0b6fa5afe92faa37() (__obf_316af79472661247 float64) {
	__obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.
		// first char
		__obf_edd3c3885447229b

	if __obf_28d099df85f083a8 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
		return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
	}
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
	__obf_28d099df85f083a8++
	__obf_5e075812c416735d := __obf_317d0145726c1931[__obf_0c1bc1e511a43120]
	switch __obf_5e075812c416735d {
	case __obf_234bd47da6a58289:
		return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
	case __obf_06719ac845597e38:
		__obf_edd9fe48d39445e4.
			ReportError("readFloat64", "empty number")
		return
	case __obf_0e8edd179006cfa4:
		__obf_edd9fe48d39445e4.
			ReportError("readFloat64", "leading dot is invalid")
		return
	case 0:
		if __obf_28d099df85f083a8 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
			return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
		}
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		switch __obf_0c1bc1e511a43120 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_edd9fe48d39445e4.
				ReportError("readFloat64", "leading zero is invalid")
			return
		}
	}
	__obf_1a534fe904baaf4a := uint64(__obf_5e075812c416735d)
__obf_43eb439dd24b9446: // chars before dot

	for ; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		__obf_5e075812c416735d := __obf_317d0145726c1931[__obf_0c1bc1e511a43120]
		switch __obf_5e075812c416735d {
		case __obf_234bd47da6a58289:
			return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
		case __obf_06719ac845597e38:
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
			return float64(__obf_1a534fe904baaf4a)
		case __obf_0e8edd179006cfa4:
			break __obf_43eb439dd24b9446
		}
		if __obf_1a534fe904baaf4a > __obf_61c0bbfc8fa559d6 {
			return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
		}
		__obf_1a534fe904baaf4a = (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint64(__obf_5e075812c416735d) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_0c1bc1e511a43120 == '.' {
		__obf_28d099df85f083a8++
		__obf_1d493182e4208194 := 0
		if __obf_28d099df85f083a8 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
			return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
		}
		for ; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			__obf_5e075812c416735d := __obf_317d0145726c1931[__obf_0c1bc1e511a43120]
			switch __obf_5e075812c416735d {
			case __obf_06719ac845597e38:
				if __obf_1d493182e4208194 > 0 && __obf_1d493182e4208194 < len(__obf_cef54dda9534fb7d) {
					__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
					return float64(__obf_1a534fe904baaf4a) / float64(__obf_cef54dda9534fb7d[__obf_1d493182e4208194])
				}
				// too many decimal places
				return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
			case __obf_234bd47da6a58289, __obf_0e8edd179006cfa4:
				return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
			}
			__obf_1d493182e4208194++
			if __obf_1a534fe904baaf4a > __obf_61c0bbfc8fa559d6 {
				return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
			}
			__obf_1a534fe904baaf4a = (__obf_1a534fe904baaf4a << 3) + (__obf_1a534fe904baaf4a << 1) + uint64(__obf_5e075812c416735d)
			if __obf_1a534fe904baaf4a > __obf_8d3d71fc19d51467 {
				return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
			}
		}
	}
	return __obf_edd9fe48d39445e4.__obf_31d2be1747d975a0()
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_31d2be1747d975a0() (__obf_316af79472661247 float64) {
	__obf_2d944450d48e5810 := __obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818()
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		return
	}
	__obf_759e6867bb34418a := __obf_a9febc73c46e733d(__obf_2d944450d48e5810)
	if __obf_759e6867bb34418a != "" {
		__obf_edd9fe48d39445e4.
			ReportError("readFloat64SlowPath", __obf_759e6867bb34418a)
		return
	}
	__obf_d59813f23ad740a8, __obf_5966ecc5edb9b17e := strconv.ParseFloat(__obf_2d944450d48e5810, 64)
	if __obf_5966ecc5edb9b17e != nil {
		__obf_edd9fe48d39445e4.
			Error = __obf_5966ecc5edb9b17e
		return
	}
	return __obf_d59813f23ad740a8
}

func __obf_a9febc73c46e733d(__obf_2d944450d48e5810 string) string {
	// strconv.ParseFloat is not validating `1.` or `1.e1`
	if len(__obf_2d944450d48e5810) == 0 {
		return "empty number"
	}
	if __obf_2d944450d48e5810[0] == '-' {
		return "-- is not valid"
	}
	__obf_487c21b17aed4185 := strings.IndexByte(__obf_2d944450d48e5810, '.')
	if __obf_487c21b17aed4185 != -1 {
		if __obf_487c21b17aed4185 == len(__obf_2d944450d48e5810)-1 {
			return "dot can not be last character"
		}
		switch __obf_2d944450d48e5810[__obf_487c21b17aed4185+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return "missing digit after dot"
		}
	}
	return ""
}

// ReadNumber read json.Number
func (__obf_edd9fe48d39445e4 *Iterator) ReadNumber() (__obf_316af79472661247 json.Number) {
	return json.Number(__obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818())
}
