//go:build !jsoniter_sloppy
// +build !jsoniter_sloppy

package __obf_030d39f7a8de96c6

import (
	"fmt"
	"io"
)

func (__obf_4ab56a99965952e7 *Iterator) __obf_38c33066e8f2b292() {
	if !__obf_4ab56a99965952e7.__obf_4deca1324d600f9e() {
		__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
		if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
			return
		}
		__obf_4ab56a99965952e7.
			ReadFloat64()
		if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
			__obf_4ab56a99965952e7.
				Error = nil
			__obf_4ab56a99965952e7.
				ReadBigFloat()
		}
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_4deca1324d600f9e() bool {
	__obf_eb25c0abd054cef0 := false
	for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		switch __obf_24309b3b0ff9ba22 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if __obf_eb25c0abd054cef0 {
				__obf_4ab56a99965952e7.
					ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			}
			if __obf_82c6e05b9d226c58+1 == __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc {
				return false
			}
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58+1]
			switch __obf_24309b3b0ff9ba22 {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				__obf_4ab56a99965952e7.
					ReportError("validateNumber", `missing digit after dot`)
				return true // already failed
			}
			__obf_eb25c0abd054cef0 = true
		default:
			switch __obf_24309b3b0ff9ba22 {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				if __obf_4ab56a99965952e7.__obf_5e22d6270d27491f == __obf_82c6e05b9d226c58 {
					return false // if - without following digits
				}
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_ab668369e47b63de() {
	if !__obf_4ab56a99965952e7.__obf_4246b78598a3743b() {
		__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
		__obf_4ab56a99965952e7.
			ReadString()
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_4246b78598a3743b() bool {
	for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		if __obf_24309b3b0ff9ba22 == '"' {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
			return true // valid
		} else if __obf_24309b3b0ff9ba22 == '\\' {
			return false
		} else if __obf_24309b3b0ff9ba22 < ' ' {
			__obf_4ab56a99965952e7.
				ReportError("trySkipString",
					fmt.Sprintf(`invalid control character found: %d`, __obf_24309b3b0ff9ba22))
			return true // already failed
		}
	}
	return false
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_56cafb6b2d847244() {
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	__obf_4ab56a99965952e7.
		ReadObjectCB(func(__obf_4ab56a99965952e7 *Iterator, __obf_cd4d02f264b18d55 string) bool {
			__obf_4ab56a99965952e7.
				Skip()
			return true
		})
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_e8e6ee2ef6c27f13() {
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	__obf_4ab56a99965952e7.
		ReadArrayCB(func(__obf_4ab56a99965952e7 *Iterator) bool {
			__obf_4ab56a99965952e7.
				Skip()
			return true
		})
}
