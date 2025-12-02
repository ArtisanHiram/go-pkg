//go:build !jsoniter_sloppy
// +build !jsoniter_sloppy

package __obf_c7b79b12b33d8238

import (
	"fmt"
	"io"
)

func (__obf_0da8c843dad13139 *Iterator) __obf_9dd6f9cece1dabe0() {
	if !__obf_0da8c843dad13139.__obf_6ecb89a198a8ba80() {
		__obf_0da8c843dad13139.__obf_a675366c80290b83()
		if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
			return
		}
		__obf_0da8c843dad13139.
			ReadFloat64()
		if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
			__obf_0da8c843dad13139.
				Error = nil
			__obf_0da8c843dad13139.
				ReadBigFloat()
		}
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_6ecb89a198a8ba80() bool {
	__obf_1d6a637641774016 := false
	for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		switch __obf_12577c03a12f0d2c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if __obf_1d6a637641774016 {
				__obf_0da8c843dad13139.
					ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			}
			if __obf_a051269af3a541bb+1 == __obf_0da8c843dad13139.__obf_840246080559848c {
				return false
			}
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb+1]
			switch __obf_12577c03a12f0d2c {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				__obf_0da8c843dad13139.
					ReportError("validateNumber", `missing digit after dot`)
				return true // already failed
			}
			__obf_1d6a637641774016 = true
		default:
			switch __obf_12577c03a12f0d2c {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				if __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 == __obf_a051269af3a541bb {
					return false // if - without following digits
				}
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (__obf_0da8c843dad13139 *Iterator) __obf_7dda9c5406ae47f0() {
	if !__obf_0da8c843dad13139.__obf_940d0be4cc6a4b6c() {
		__obf_0da8c843dad13139.__obf_a675366c80290b83()
		__obf_0da8c843dad13139.
			ReadString()
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_940d0be4cc6a4b6c() bool {
	for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		if __obf_12577c03a12f0d2c == '"' {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
			return true // valid
		} else if __obf_12577c03a12f0d2c == '\\' {
			return false
		} else if __obf_12577c03a12f0d2c < ' ' {
			__obf_0da8c843dad13139.
				ReportError("trySkipString",
					fmt.Sprintf(`invalid control character found: %d`, __obf_12577c03a12f0d2c))
			return true // already failed
		}
	}
	return false
}

func (__obf_0da8c843dad13139 *Iterator) __obf_f42b686c47b3acb5() {
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	__obf_0da8c843dad13139.
		ReadObjectCB(func(__obf_0da8c843dad13139 *Iterator, __obf_ad34f8a6de2b01cb string) bool {
			__obf_0da8c843dad13139.
				Skip()
			return true
		})
}

func (__obf_0da8c843dad13139 *Iterator) __obf_a12f980c2c849354() {
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	__obf_0da8c843dad13139.
		ReadArrayCB(func(__obf_0da8c843dad13139 *Iterator) bool {
			__obf_0da8c843dad13139.
				Skip()
			return true
		})
}
