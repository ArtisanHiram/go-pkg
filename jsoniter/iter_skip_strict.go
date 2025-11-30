//go:build !jsoniter_sloppy
// +build !jsoniter_sloppy

package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"io"
)

func (__obf_edd9fe48d39445e4 *Iterator) __obf_fdeaf1c2cbe3aa15() {
	if !__obf_edd9fe48d39445e4.__obf_a949c49878771368() {
		__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
		if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
			return
		}
		__obf_edd9fe48d39445e4.
			ReadFloat64()
		if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
			__obf_edd9fe48d39445e4.
				Error = nil
			__obf_edd9fe48d39445e4.
				ReadBigFloat()
		}
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_a949c49878771368() bool {
	__obf_1d30c55eef676224 := false
	for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		switch __obf_0c1bc1e511a43120 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if __obf_1d30c55eef676224 {
				__obf_edd9fe48d39445e4.
					ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			}
			if __obf_28d099df85f083a8+1 == __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 {
				return false
			}
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8+1]
			switch __obf_0c1bc1e511a43120 {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				__obf_edd9fe48d39445e4.
					ReportError("validateNumber", `missing digit after dot`)
				return true // already failed
			}
			__obf_1d30c55eef676224 = true
		default:
			switch __obf_0c1bc1e511a43120 {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				if __obf_edd9fe48d39445e4.__obf_edd3c3885447229b == __obf_28d099df85f083a8 {
					return false // if - without following digits
				}
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_d743935ecad873a7() {
	if !__obf_edd9fe48d39445e4.__obf_774a800868ae1136() {
		__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
		__obf_edd9fe48d39445e4.
			ReadString()
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_774a800868ae1136() bool {
	for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		if __obf_0c1bc1e511a43120 == '"' {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
			return true // valid
		} else if __obf_0c1bc1e511a43120 == '\\' {
			return false
		} else if __obf_0c1bc1e511a43120 < ' ' {
			__obf_edd9fe48d39445e4.
				ReportError("trySkipString",
					fmt.Sprintf(`invalid control character found: %d`, __obf_0c1bc1e511a43120))
			return true // already failed
		}
	}
	return false
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_0704e073eacf05e6() {
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	__obf_edd9fe48d39445e4.
		ReadObjectCB(func(__obf_edd9fe48d39445e4 *Iterator, __obf_f992c91036745ca0 string) bool {
			__obf_edd9fe48d39445e4.
				Skip()
			return true
		})
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_5566b22788a0a705() {
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	__obf_edd9fe48d39445e4.
		ReadArrayCB(func(__obf_edd9fe48d39445e4 *Iterator) bool {
			__obf_edd9fe48d39445e4.
				Skip()
			return true
		})
}
