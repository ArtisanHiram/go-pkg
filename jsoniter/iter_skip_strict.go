//go:build !jsoniter_sloppy
// +build !jsoniter_sloppy

package __obf_91620b895eeff9ed

import (
	"fmt"
	"io"
)

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_d6a474bd319b770c() {
	if !__obf_1bb30e8a74ed8233.__obf_76a4e2d10cb4f2ba() {
		__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
		if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
			return
		}
		__obf_1bb30e8a74ed8233.
			ReadFloat64()
		if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
			__obf_1bb30e8a74ed8233.
				Error = nil
			__obf_1bb30e8a74ed8233.
				ReadBigFloat()
		}
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_76a4e2d10cb4f2ba() bool {
	__obf_4641fb766806f529 := false
	for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		switch __obf_f16b4157911bc9af {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if __obf_4641fb766806f529 {
				__obf_1bb30e8a74ed8233.
					ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			}
			if __obf_5aa5c8829b97f182+1 == __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae {
				return false
			}
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182+1]
			switch __obf_f16b4157911bc9af {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				__obf_1bb30e8a74ed8233.
					ReportError("validateNumber", `missing digit after dot`)
				return true // already failed
			}
			__obf_4641fb766806f529 = true
		default:
			switch __obf_f16b4157911bc9af {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				if __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 == __obf_5aa5c8829b97f182 {
					return false // if - without following digits
				}
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_c5493c5b3916fc3f() {
	if !__obf_1bb30e8a74ed8233.__obf_b4f60d9d56f1f38d() {
		__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
		__obf_1bb30e8a74ed8233.
			ReadString()
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_b4f60d9d56f1f38d() bool {
	for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		if __obf_f16b4157911bc9af == '"' {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
			return true // valid
		} else if __obf_f16b4157911bc9af == '\\' {
			return false
		} else if __obf_f16b4157911bc9af < ' ' {
			__obf_1bb30e8a74ed8233.
				ReportError("trySkipString",
					fmt.Sprintf(`invalid control character found: %d`, __obf_f16b4157911bc9af))
			return true // already failed
		}
	}
	return false
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_8d12ebf64f0517d2() {
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	__obf_1bb30e8a74ed8233.
		ReadObjectCB(func(__obf_1bb30e8a74ed8233 *Iterator, __obf_7e01b5b4651074d4 string) bool {
			__obf_1bb30e8a74ed8233.
				Skip()
			return true
		})
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_531ac82b5de39aec() {
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	__obf_1bb30e8a74ed8233.
		ReadArrayCB(func(__obf_1bb30e8a74ed8233 *Iterator) bool {
			__obf_1bb30e8a74ed8233.
				Skip()
			return true
		})
}
