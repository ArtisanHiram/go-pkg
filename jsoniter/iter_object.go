package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (__obf_edd9fe48d39445e4 *Iterator) ReadObject() (__obf_316af79472661247 string) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	switch __obf_0c1bc1e511a43120 {
	case 'n':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return "" // null
	case '{':
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 == '"' {
			__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
			__obf_f992c91036745ca0 := __obf_edd9fe48d39445e4.ReadString()
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			if __obf_0c1bc1e511a43120 != ':' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
			}
			return __obf_f992c91036745ca0
		}
		if __obf_0c1bc1e511a43120 == '}' {
			return "" // end of object
		}
		__obf_edd9fe48d39445e4.
			ReportError("ReadObject", `expect " after {, but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return
	case ',':
		__obf_f992c91036745ca0 := __obf_edd9fe48d39445e4.ReadString()
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 != ':' {
			__obf_edd9fe48d39445e4.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
		}
		return __obf_f992c91036745ca0
	case '}':
		return "" // end of object
	default:
		__obf_edd9fe48d39445e4.
			ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{__obf_0c1bc1e511a43120})))
		return
	}
}

// CaseInsensitive
func (__obf_edd9fe48d39445e4 *Iterator) __obf_18f2978e30c0e899() int64 {
	__obf_63f7a48112c37b0e := int64(0x811c9dc5)
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 != '"' {
		__obf_edd9fe48d39445e4.
			ReportError("readFieldHash", `expect ", but found `+string([]byte{__obf_0c1bc1e511a43120}))
		return 0
	}
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b;
		// require ascii string and no escape
		__obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_902d7026e8a09dd2 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			if __obf_902d7026e8a09dd2 == '\\' {
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
				for _, __obf_902d7026e8a09dd2 := range __obf_edd9fe48d39445e4.__obf_561396e75273c9da() {
					if 'A' <= __obf_902d7026e8a09dd2 && __obf_902d7026e8a09dd2 <= 'Z' && !__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_b600271a247f99b8 {
						__obf_902d7026e8a09dd2 += 'a' - 'A'
					}
					__obf_63f7a48112c37b0e ^= int64(__obf_902d7026e8a09dd2)
					__obf_63f7a48112c37b0e *= 0x1000193
				}
				__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
				if __obf_0c1bc1e511a43120 != ':' {
					__obf_edd9fe48d39445e4.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_0c1bc1e511a43120}))
					return 0
				}
				return __obf_63f7a48112c37b0e
			}
			if __obf_902d7026e8a09dd2 == '"' {
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
				__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
				if __obf_0c1bc1e511a43120 != ':' {
					__obf_edd9fe48d39445e4.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_0c1bc1e511a43120}))
					return 0
				}
				return __obf_63f7a48112c37b0e
			}
			if 'A' <= __obf_902d7026e8a09dd2 && __obf_902d7026e8a09dd2 <= 'Z' && !__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_b600271a247f99b8 {
				__obf_902d7026e8a09dd2 += 'a' - 'A'
			}
			__obf_63f7a48112c37b0e ^= int64(__obf_902d7026e8a09dd2)
			__obf_63f7a48112c37b0e *= 0x1000193
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			__obf_edd9fe48d39445e4.
				ReportError("readFieldHash", `incomplete field name`)
			return 0
		}
	}
}

func __obf_5651d62f410f485f(__obf_2d944450d48e5810 string, __obf_b600271a247f99b8 bool) int64 {
	if !__obf_b600271a247f99b8 {
		__obf_2d944450d48e5810 = strings.ToLower(__obf_2d944450d48e5810)
	}
	__obf_63f7a48112c37b0e := int64(0x811c9dc5)
	for _, __obf_902d7026e8a09dd2 := range []byte(__obf_2d944450d48e5810) {
		__obf_63f7a48112c37b0e ^= int64(__obf_902d7026e8a09dd2)
		__obf_63f7a48112c37b0e *= 0x1000193
	}
	return int64(__obf_63f7a48112c37b0e)
}

// ReadObjectCB read object with callback, the key is ascii only and field name not copied
func (__obf_edd9fe48d39445e4 *Iterator) ReadObjectCB(__obf_d1ec68aebedafce9 func(*Iterator, string) bool) bool {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	var __obf_f992c91036745ca0 string
	if __obf_0c1bc1e511a43120 == '{' {
		if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
			return false
		}
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 == '"' {
			__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
			__obf_f992c91036745ca0 = __obf_edd9fe48d39445e4.ReadString()
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			if __obf_0c1bc1e511a43120 != ':' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
			}
			if !__obf_d1ec68aebedafce9(__obf_edd9fe48d39445e4, __obf_f992c91036745ca0) {
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			for __obf_0c1bc1e511a43120 == ',' {
				__obf_f992c91036745ca0 = __obf_edd9fe48d39445e4.ReadString()
				__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
				if __obf_0c1bc1e511a43120 != ':' {
					__obf_edd9fe48d39445e4.
						ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
				}
				if !__obf_d1ec68aebedafce9(__obf_edd9fe48d39445e4, __obf_f992c91036745ca0) {
					__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
					return false
				}
				__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			}
			if __obf_0c1bc1e511a43120 != '}' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadObjectCB", `object not ended with }`)
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			return __obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		}
		if __obf_0c1bc1e511a43120 == '}' {
			return __obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		}
		__obf_edd9fe48d39445e4.
			ReportError("ReadObjectCB", `expect " after {, but found `+string([]byte{__obf_0c1bc1e511a43120}))
		__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		return false
	}
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return true // null
	}
	__obf_edd9fe48d39445e4.
		ReportError("ReadObjectCB", `expect { or n, but found `+string([]byte{__obf_0c1bc1e511a43120}))
	return false
}

// ReadMapCB read map with callback, the key can be any string
func (__obf_edd9fe48d39445e4 *Iterator) ReadMapCB(__obf_d1ec68aebedafce9 func(*Iterator, string) bool) bool {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '{' {
		if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
			return false
		}
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 == '"' {
			__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
			__obf_f992c91036745ca0 := __obf_edd9fe48d39445e4.ReadString()
			if __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11() != ':' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			if !__obf_d1ec68aebedafce9(__obf_edd9fe48d39445e4, __obf_f992c91036745ca0) {
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			for __obf_0c1bc1e511a43120 == ',' {
				__obf_f992c91036745ca0 = __obf_edd9fe48d39445e4.ReadString()
				if __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11() != ':' {
					__obf_edd9fe48d39445e4.
						ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_0c1bc1e511a43120}))
					__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
					return false
				}
				if !__obf_d1ec68aebedafce9(__obf_edd9fe48d39445e4, __obf_f992c91036745ca0) {
					__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
					return false
				}
				__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			}
			if __obf_0c1bc1e511a43120 != '}' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadMapCB", `object not ended with }`)
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			return __obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		}
		if __obf_0c1bc1e511a43120 == '}' {
			return __obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		}
		__obf_edd9fe48d39445e4.
			ReportError("ReadMapCB", `expect " after {, but found `+string([]byte{__obf_0c1bc1e511a43120}))
		__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		return false
	}
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return true // null
	}
	__obf_edd9fe48d39445e4.
		ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_0c1bc1e511a43120}))
	return false
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_870a32c6708136fd() bool {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	switch __obf_0c1bc1e511a43120 {
	case '{':
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 == '}' {
			return false
		}
		__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
		return true
	case 'n':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return false
	}
	__obf_edd9fe48d39445e4.
		ReportError("readObjectStart", "expect { or n, but found "+string([]byte{__obf_0c1bc1e511a43120}))
	return false
}
