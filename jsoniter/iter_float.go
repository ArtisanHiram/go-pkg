package __obf_c7b79b12b33d8238

import (
	"encoding/json"
	"io"
	"math/big"
	"strconv"
	"strings"
	"unsafe"
)

var __obf_50da0ade3c8b6b71 []int8

const __obf_905dd362f7428491 = int8(-1)
const __obf_950190363e7680e9 = int8(-2)
const __obf_ef4636e4ba3b2808 = int8(-3)

func init() {
	__obf_50da0ade3c8b6b71 = make([]int8, 256)
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < len(__obf_50da0ade3c8b6b71); __obf_a051269af3a541bb++ {
		__obf_50da0ade3c8b6b71[__obf_a051269af3a541bb] = __obf_905dd362f7428491
	}
	for __obf_a051269af3a541bb := int8('0'); __obf_a051269af3a541bb <= int8('9'); __obf_a051269af3a541bb++ {
		__obf_50da0ade3c8b6b71[__obf_a051269af3a541bb] = __obf_a051269af3a541bb - int8('0')
	}
	__obf_50da0ade3c8b6b71[','] = __obf_950190363e7680e9
	__obf_50da0ade3c8b6b71[']'] = __obf_950190363e7680e9
	__obf_50da0ade3c8b6b71['}'] = __obf_950190363e7680e9
	__obf_50da0ade3c8b6b71[' '] = __obf_950190363e7680e9
	__obf_50da0ade3c8b6b71['\t'] = __obf_950190363e7680e9
	__obf_50da0ade3c8b6b71['\n'] = __obf_950190363e7680e9
	__obf_50da0ade3c8b6b71['.'] = __obf_ef4636e4ba3b2808
}

// ReadBigFloat read big.Float
func (__obf_0da8c843dad13139 *Iterator) ReadBigFloat() (__obf_9a8638dac9e1c083 *big.Float) {
	__obf_a3eaafc22faf1644 := __obf_0da8c843dad13139.__obf_a678d14dd84cbb6b()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		return nil
	}
	__obf_276516e3fd0b1cbb := 64
	if len(__obf_a3eaafc22faf1644) > __obf_276516e3fd0b1cbb {
		__obf_276516e3fd0b1cbb = len(__obf_a3eaafc22faf1644)
	}
	__obf_35accd096612ebe4, _, __obf_ea0680f8b461a85b := big.ParseFloat(__obf_a3eaafc22faf1644, 10, uint(__obf_276516e3fd0b1cbb), big.ToZero)
	if __obf_ea0680f8b461a85b != nil {
		__obf_0da8c843dad13139.
			Error = __obf_ea0680f8b461a85b
		return nil
	}
	return __obf_35accd096612ebe4
}

// ReadBigInt read big.Int
func (__obf_0da8c843dad13139 *Iterator) ReadBigInt() (__obf_9a8638dac9e1c083 *big.Int) {
	__obf_a3eaafc22faf1644 := __obf_0da8c843dad13139.__obf_a678d14dd84cbb6b()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		return nil
	}
	__obf_9a8638dac9e1c083 = big.NewInt(0)
	var __obf_b145180ce2c394bf bool
	__obf_9a8638dac9e1c083, __obf_b145180ce2c394bf = __obf_9a8638dac9e1c083.SetString(__obf_a3eaafc22faf1644, 10)
	if !__obf_b145180ce2c394bf {
		__obf_0da8c843dad13139.
			ReportError("ReadBigInt", "invalid big int")
		return nil
	}
	return __obf_9a8638dac9e1c083
}

// ReadFloat32 read float32
func (__obf_0da8c843dad13139 *Iterator) ReadFloat32() (__obf_9a8638dac9e1c083 float32) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '-' {
		return -__obf_0da8c843dad13139.__obf_b55848e4e7e1bb21()
	}
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	return __obf_0da8c843dad13139.__obf_b55848e4e7e1bb21()
}

func (__obf_0da8c843dad13139 *Iterator) __obf_b55848e4e7e1bb21() (__obf_9a8638dac9e1c083 float32) {
	__obf_a051269af3a541bb := __obf_0da8c843dad13139.
		// first char
		__obf_6a63c9ac34fe84e2

	if __obf_a051269af3a541bb == __obf_0da8c843dad13139.__obf_840246080559848c {
		return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
	}
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
	__obf_a051269af3a541bb++
	__obf_83600c60799be31c := __obf_50da0ade3c8b6b71[__obf_12577c03a12f0d2c]
	switch __obf_83600c60799be31c {
	case __obf_905dd362f7428491:
		return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
	case __obf_950190363e7680e9:
		__obf_0da8c843dad13139.
			ReportError("readFloat32", "empty number")
		return
	case __obf_ef4636e4ba3b2808:
		__obf_0da8c843dad13139.
			ReportError("readFloat32", "leading dot is invalid")
		return
	case 0:
		if __obf_a051269af3a541bb == __obf_0da8c843dad13139.__obf_840246080559848c {
			return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
		}
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		switch __obf_12577c03a12f0d2c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_0da8c843dad13139.
				ReportError("readFloat32", "leading zero is invalid")
			return
		}
	}
	__obf_52321dce0d8db989 := uint64(__obf_83600c60799be31c)
__obf_f3eb43a3d7ddf79e: // chars before dot

	for ; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		__obf_83600c60799be31c := __obf_50da0ade3c8b6b71[__obf_12577c03a12f0d2c]
		switch __obf_83600c60799be31c {
		case __obf_905dd362f7428491:
			return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
		case __obf_950190363e7680e9:
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			return float32(__obf_52321dce0d8db989)
		case __obf_ef4636e4ba3b2808:
			break __obf_f3eb43a3d7ddf79e
		}
		if __obf_52321dce0d8db989 > __obf_888f28ae8b5c39dc {
			return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
		}
		__obf_52321dce0d8db989 = (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint64(__obf_83600c60799be31c) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_12577c03a12f0d2c == '.' {
		__obf_a051269af3a541bb++
		__obf_fff48a6ac660ba0e := 0
		if __obf_a051269af3a541bb == __obf_0da8c843dad13139.__obf_840246080559848c {
			return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
		}
		for ; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			__obf_83600c60799be31c := __obf_50da0ade3c8b6b71[__obf_12577c03a12f0d2c]
			switch __obf_83600c60799be31c {
			case __obf_950190363e7680e9:
				if __obf_fff48a6ac660ba0e > 0 && __obf_fff48a6ac660ba0e < len(__obf_e075611d0d831082) {
					__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
					return float32(float64(__obf_52321dce0d8db989) / float64(__obf_e075611d0d831082[__obf_fff48a6ac660ba0e]))
				}
				// too many decimal places
				return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
			case __obf_905dd362f7428491, __obf_ef4636e4ba3b2808:
				return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
			}
			__obf_fff48a6ac660ba0e++
			if __obf_52321dce0d8db989 > __obf_888f28ae8b5c39dc {
				return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
			}
			__obf_52321dce0d8db989 = (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint64(__obf_83600c60799be31c)
		}
	}
	return __obf_0da8c843dad13139.__obf_bf89968e1a143c58()
}

func (__obf_0da8c843dad13139 *Iterator) __obf_a678d14dd84cbb6b() (__obf_9a8638dac9e1c083 string) {
	__obf_a348ad18fd1fbc21 := [16]byte{}
	__obf_a3eaafc22faf1644 := __obf_a348ad18fd1fbc21[0:0]
__obf_4ac90cc32cc49538:
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			switch __obf_12577c03a12f0d2c {
			case '+', '-', '.', 'e', 'E', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				__obf_a3eaafc22faf1644 = append(__obf_a3eaafc22faf1644, __obf_12577c03a12f0d2c)
				continue
			default:
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
				break __obf_4ac90cc32cc49538
			}
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			break
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		return
	}
	if len(__obf_a3eaafc22faf1644) == 0 {
		__obf_0da8c843dad13139.
			ReportError("readNumberAsString", "invalid number")
	}
	return *(*string)(unsafe.Pointer(&__obf_a3eaafc22faf1644))
}

func (__obf_0da8c843dad13139 *Iterator) __obf_bf89968e1a143c58() (__obf_9a8638dac9e1c083 float32) {
	__obf_a3eaafc22faf1644 := __obf_0da8c843dad13139.__obf_a678d14dd84cbb6b()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		return
	}
	__obf_78ae98784d055850 := __obf_883acdb7883a5163(__obf_a3eaafc22faf1644)
	if __obf_78ae98784d055850 != "" {
		__obf_0da8c843dad13139.
			ReportError("readFloat32SlowPath", __obf_78ae98784d055850)
		return
	}
	__obf_35accd096612ebe4, __obf_ea0680f8b461a85b := strconv.ParseFloat(__obf_a3eaafc22faf1644, 32)
	if __obf_ea0680f8b461a85b != nil {
		__obf_0da8c843dad13139.
			Error = __obf_ea0680f8b461a85b
		return
	}
	return float32(__obf_35accd096612ebe4)
}

// ReadFloat64 read float64
func (__obf_0da8c843dad13139 *Iterator) ReadFloat64() (__obf_9a8638dac9e1c083 float64) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '-' {
		return -__obf_0da8c843dad13139.__obf_267353792e8f7446()
	}
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	return __obf_0da8c843dad13139.__obf_267353792e8f7446()
}

func (__obf_0da8c843dad13139 *Iterator) __obf_267353792e8f7446() (__obf_9a8638dac9e1c083 float64) {
	__obf_a051269af3a541bb := __obf_0da8c843dad13139.
		// first char
		__obf_6a63c9ac34fe84e2

	if __obf_a051269af3a541bb == __obf_0da8c843dad13139.__obf_840246080559848c {
		return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
	}
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
	__obf_a051269af3a541bb++
	__obf_83600c60799be31c := __obf_50da0ade3c8b6b71[__obf_12577c03a12f0d2c]
	switch __obf_83600c60799be31c {
	case __obf_905dd362f7428491:
		return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
	case __obf_950190363e7680e9:
		__obf_0da8c843dad13139.
			ReportError("readFloat64", "empty number")
		return
	case __obf_ef4636e4ba3b2808:
		__obf_0da8c843dad13139.
			ReportError("readFloat64", "leading dot is invalid")
		return
	case 0:
		if __obf_a051269af3a541bb == __obf_0da8c843dad13139.__obf_840246080559848c {
			return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
		}
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		switch __obf_12577c03a12f0d2c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			__obf_0da8c843dad13139.
				ReportError("readFloat64", "leading zero is invalid")
			return
		}
	}
	__obf_52321dce0d8db989 := uint64(__obf_83600c60799be31c)
__obf_f3eb43a3d7ddf79e: // chars before dot

	for ; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		__obf_83600c60799be31c := __obf_50da0ade3c8b6b71[__obf_12577c03a12f0d2c]
		switch __obf_83600c60799be31c {
		case __obf_905dd362f7428491:
			return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
		case __obf_950190363e7680e9:
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
			return float64(__obf_52321dce0d8db989)
		case __obf_ef4636e4ba3b2808:
			break __obf_f3eb43a3d7ddf79e
		}
		if __obf_52321dce0d8db989 > __obf_888f28ae8b5c39dc {
			return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
		}
		__obf_52321dce0d8db989 = (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint64(__obf_83600c60799be31c) // value = value * 10 + ind;
	}
	// chars after dot
	if __obf_12577c03a12f0d2c == '.' {
		__obf_a051269af3a541bb++
		__obf_fff48a6ac660ba0e := 0
		if __obf_a051269af3a541bb == __obf_0da8c843dad13139.__obf_840246080559848c {
			return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
		}
		for ; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			__obf_83600c60799be31c := __obf_50da0ade3c8b6b71[__obf_12577c03a12f0d2c]
			switch __obf_83600c60799be31c {
			case __obf_950190363e7680e9:
				if __obf_fff48a6ac660ba0e > 0 && __obf_fff48a6ac660ba0e < len(__obf_e075611d0d831082) {
					__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
					return float64(__obf_52321dce0d8db989) / float64(__obf_e075611d0d831082[__obf_fff48a6ac660ba0e])
				}
				// too many decimal places
				return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
			case __obf_905dd362f7428491, __obf_ef4636e4ba3b2808:
				return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
			}
			__obf_fff48a6ac660ba0e++
			if __obf_52321dce0d8db989 > __obf_888f28ae8b5c39dc {
				return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
			}
			__obf_52321dce0d8db989 = (__obf_52321dce0d8db989 << 3) + (__obf_52321dce0d8db989 << 1) + uint64(__obf_83600c60799be31c)
			if __obf_52321dce0d8db989 > __obf_a86df43efd02f0dc {
				return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
			}
		}
	}
	return __obf_0da8c843dad13139.__obf_57bbb083c2ea3b24()
}

func (__obf_0da8c843dad13139 *Iterator) __obf_57bbb083c2ea3b24() (__obf_9a8638dac9e1c083 float64) {
	__obf_a3eaafc22faf1644 := __obf_0da8c843dad13139.__obf_a678d14dd84cbb6b()
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		return
	}
	__obf_78ae98784d055850 := __obf_883acdb7883a5163(__obf_a3eaafc22faf1644)
	if __obf_78ae98784d055850 != "" {
		__obf_0da8c843dad13139.
			ReportError("readFloat64SlowPath", __obf_78ae98784d055850)
		return
	}
	__obf_35accd096612ebe4, __obf_ea0680f8b461a85b := strconv.ParseFloat(__obf_a3eaafc22faf1644, 64)
	if __obf_ea0680f8b461a85b != nil {
		__obf_0da8c843dad13139.
			Error = __obf_ea0680f8b461a85b
		return
	}
	return __obf_35accd096612ebe4
}

func __obf_883acdb7883a5163(__obf_a3eaafc22faf1644 string) string {
	// strconv.ParseFloat is not validating `1.` or `1.e1`
	if len(__obf_a3eaafc22faf1644) == 0 {
		return "empty number"
	}
	if __obf_a3eaafc22faf1644[0] == '-' {
		return "-- is not valid"
	}
	__obf_aac738e34aeab64b := strings.IndexByte(__obf_a3eaafc22faf1644, '.')
	if __obf_aac738e34aeab64b != -1 {
		if __obf_aac738e34aeab64b == len(__obf_a3eaafc22faf1644)-1 {
			return "dot can not be last character"
		}
		switch __obf_a3eaafc22faf1644[__obf_aac738e34aeab64b+1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		default:
			return "missing digit after dot"
		}
	}
	return ""
}

// ReadNumber read json.Number
func (__obf_0da8c843dad13139 *Iterator) ReadNumber() (__obf_9a8638dac9e1c083 json.Number) {
	return json.Number(__obf_0da8c843dad13139.__obf_a678d14dd84cbb6b())
}
