//go:build !jsoniter_sloppy
// +build !jsoniter_sloppy

package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"io"
)

func (__obf_67008a6a9e5ba828 *Iterator) __obf_a64603a02c650185() {
	if !__obf_67008a6a9e5ba828.__obf_7719b9789b951612() {
		__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
		if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
			return
		}
		__obf_67008a6a9e5ba828.
			ReadFloat64()
		if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
			__obf_67008a6a9e5ba828.
				Error = nil
			__obf_67008a6a9e5ba828.
				ReadBigFloat()
		}
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_7719b9789b951612() bool {
	__obf_dc39df2b1294d628 := false
	for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		switch __obf_dab9baaadfa7c8c2 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if __obf_dc39df2b1294d628 {
				__obf_67008a6a9e5ba828.
					ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			}
			if __obf_2deec7c38ffb6ae3+1 == __obf_67008a6a9e5ba828.__obf_3a36550914545c79 {
				return false
			}
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3+1]
			switch __obf_dab9baaadfa7c8c2 {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				__obf_67008a6a9e5ba828.
					ReportError("validateNumber", `missing digit after dot`)
				return true // already failed
			}
			__obf_dc39df2b1294d628 = true
		default:
			switch __obf_dab9baaadfa7c8c2 {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				if __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 == __obf_2deec7c38ffb6ae3 {
					return false // if - without following digits
				}
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_acb079d30d4527e3() {
	if !__obf_67008a6a9e5ba828.__obf_58e557e4c92ca345() {
		__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
		__obf_67008a6a9e5ba828.
			ReadString()
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_58e557e4c92ca345() bool {
	for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		if __obf_dab9baaadfa7c8c2 == '"' {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
			return true // valid
		} else if __obf_dab9baaadfa7c8c2 == '\\' {
			return false
		} else if __obf_dab9baaadfa7c8c2 < ' ' {
			__obf_67008a6a9e5ba828.
				ReportError("trySkipString",
					fmt.Sprintf(`invalid control character found: %d`, __obf_dab9baaadfa7c8c2))
			return true // already failed
		}
	}
	return false
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_7e3d8ea4ab2a8258() {
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	__obf_67008a6a9e5ba828.
		ReadObjectCB(func(__obf_67008a6a9e5ba828 *Iterator, __obf_61998fb083387c3c string) bool {
			__obf_67008a6a9e5ba828.
				Skip()
			return true
		})
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_1d506622058b77ef() {
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	__obf_67008a6a9e5ba828.
		ReadArrayCB(func(__obf_67008a6a9e5ba828 *Iterator) bool {
			__obf_67008a6a9e5ba828.
				Skip()
			return true
		})
}
