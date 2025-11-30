//go:build !jsoniter_sloppy
// +build !jsoniter_sloppy

package __obf_703d23ba520c3295

import (
	"fmt"
	"io"
)

func (__obf_47edab4c16a2d88d *Iterator) __obf_91b8504ad35d5b75() {
	if !__obf_47edab4c16a2d88d.__obf_d0aba758bfabdf1e() {
		__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
		if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
			return
		}
		__obf_47edab4c16a2d88d.
			ReadFloat64()
		if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
			__obf_47edab4c16a2d88d.
				Error = nil
			__obf_47edab4c16a2d88d.
				ReadBigFloat()
		}
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_d0aba758bfabdf1e() bool {
	__obf_d49bcac29d63f01b := false
	for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		switch __obf_bd08f5161fff294a {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if __obf_d49bcac29d63f01b {
				__obf_47edab4c16a2d88d.
					ReportError("validateNumber", `more than one dot found in number`)
				return true // already failed
			}
			if __obf_b0a5d2bd48690f1d+1 == __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 {
				return false
			}
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d+1]
			switch __obf_bd08f5161fff294a {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				__obf_47edab4c16a2d88d.
					ReportError("validateNumber", `missing digit after dot`)
				return true // already failed
			}
			__obf_d49bcac29d63f01b = true
		default:
			switch __obf_bd08f5161fff294a {
			case ',', ']', '}', ' ', '\t', '\n', '\r':
				if __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 == __obf_b0a5d2bd48690f1d {
					return false // if - without following digits
				}
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
				return true // must be valid
			}
			return false // may be invalid
		}
	}
	return false
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_3c5a32ecdaa844fb() {
	if !__obf_47edab4c16a2d88d.__obf_15d6a08bac884f13() {
		__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
		__obf_47edab4c16a2d88d.
			ReadString()
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_15d6a08bac884f13() bool {
	for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		if __obf_bd08f5161fff294a == '"' {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
			return true // valid
		} else if __obf_bd08f5161fff294a == '\\' {
			return false
		} else if __obf_bd08f5161fff294a < ' ' {
			__obf_47edab4c16a2d88d.
				ReportError("trySkipString",
					fmt.Sprintf(`invalid control character found: %d`, __obf_bd08f5161fff294a))
			return true // already failed
		}
	}
	return false
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_e53a692fb1024439() {
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	__obf_47edab4c16a2d88d.
		ReadObjectCB(func(__obf_47edab4c16a2d88d *Iterator, __obf_13f6440419533990 string) bool {
			__obf_47edab4c16a2d88d.
				Skip()
			return true
		})
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_93a7c6915e88f56c() {
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	__obf_47edab4c16a2d88d.
		ReadArrayCB(func(__obf_47edab4c16a2d88d *Iterator) bool {
			__obf_47edab4c16a2d88d.
				Skip()
			return true
		})
}
