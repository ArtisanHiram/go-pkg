package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (__obf_67008a6a9e5ba828 *Iterator) ReadObject() (__obf_5dabcdfee5097ed6 string) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	switch __obf_dab9baaadfa7c8c2 {
	case 'n':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return "" // null
	case '{':
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 == '"' {
			__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
			__obf_61998fb083387c3c := __obf_67008a6a9e5ba828.ReadString()
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			if __obf_dab9baaadfa7c8c2 != ':' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
			}
			return __obf_61998fb083387c3c
		}
		if __obf_dab9baaadfa7c8c2 == '}' {
			return "" // end of object
		}
		__obf_67008a6a9e5ba828.
			ReportError("ReadObject", `expect " after {, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	case ',':
		__obf_61998fb083387c3c := __obf_67008a6a9e5ba828.ReadString()
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 != ':' {
			__obf_67008a6a9e5ba828.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		}
		return __obf_61998fb083387c3c
	case '}':
		return "" // end of object
	default:
		__obf_67008a6a9e5ba828.
			ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{__obf_dab9baaadfa7c8c2})))
		return
	}
}

// CaseInsensitive
func (__obf_67008a6a9e5ba828 *Iterator) __obf_64030aba1e95f95b() int64 {
	__obf_73f9e506630e35b0 := int64(0x811c9dc5)
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 != '"' {
		__obf_67008a6a9e5ba828.
			ReportError("readFieldHash", `expect ", but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		return 0
	}
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36;
		// require ascii string and no escape
		__obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_1c7157183bc604f5 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			if __obf_1c7157183bc604f5 == '\\' {
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
				for _, __obf_1c7157183bc604f5 := range __obf_67008a6a9e5ba828.__obf_ddc6a694dea48ef4() {
					if 'A' <= __obf_1c7157183bc604f5 && __obf_1c7157183bc604f5 <= 'Z' && !__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_f48e3198f571baa9 {
						__obf_1c7157183bc604f5 += 'a' - 'A'
					}
					__obf_73f9e506630e35b0 ^= int64(__obf_1c7157183bc604f5)
					__obf_73f9e506630e35b0 *= 0x1000193
				}
				__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
				if __obf_dab9baaadfa7c8c2 != ':' {
					__obf_67008a6a9e5ba828.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
					return 0
				}
				return __obf_73f9e506630e35b0
			}
			if __obf_1c7157183bc604f5 == '"' {
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
				__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
				if __obf_dab9baaadfa7c8c2 != ':' {
					__obf_67008a6a9e5ba828.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
					return 0
				}
				return __obf_73f9e506630e35b0
			}
			if 'A' <= __obf_1c7157183bc604f5 && __obf_1c7157183bc604f5 <= 'Z' && !__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_f48e3198f571baa9 {
				__obf_1c7157183bc604f5 += 'a' - 'A'
			}
			__obf_73f9e506630e35b0 ^= int64(__obf_1c7157183bc604f5)
			__obf_73f9e506630e35b0 *= 0x1000193
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			__obf_67008a6a9e5ba828.
				ReportError("readFieldHash", `incomplete field name`)
			return 0
		}
	}
}

func __obf_7dc2f08aee807385(__obf_12c21b79fa86dcba string, __obf_f48e3198f571baa9 bool) int64 {
	if !__obf_f48e3198f571baa9 {
		__obf_12c21b79fa86dcba = strings.ToLower(__obf_12c21b79fa86dcba)
	}
	__obf_73f9e506630e35b0 := int64(0x811c9dc5)
	for _, __obf_1c7157183bc604f5 := range []byte(__obf_12c21b79fa86dcba) {
		__obf_73f9e506630e35b0 ^= int64(__obf_1c7157183bc604f5)
		__obf_73f9e506630e35b0 *= 0x1000193
	}
	return int64(__obf_73f9e506630e35b0)
}

// ReadObjectCB read object with callback, the key is ascii only and field name not copied
func (__obf_67008a6a9e5ba828 *Iterator) ReadObjectCB(__obf_5fde3c82e220905a func(*Iterator, string) bool) bool {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	var __obf_61998fb083387c3c string
	if __obf_dab9baaadfa7c8c2 == '{' {
		if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
			return false
		}
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 == '"' {
			__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
			__obf_61998fb083387c3c = __obf_67008a6a9e5ba828.ReadString()
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			if __obf_dab9baaadfa7c8c2 != ':' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
			}
			if !__obf_5fde3c82e220905a(__obf_67008a6a9e5ba828, __obf_61998fb083387c3c) {
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			for __obf_dab9baaadfa7c8c2 == ',' {
				__obf_61998fb083387c3c = __obf_67008a6a9e5ba828.ReadString()
				__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
				if __obf_dab9baaadfa7c8c2 != ':' {
					__obf_67008a6a9e5ba828.
						ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
				}
				if !__obf_5fde3c82e220905a(__obf_67008a6a9e5ba828, __obf_61998fb083387c3c) {
					__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
					return false
				}
				__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			}
			if __obf_dab9baaadfa7c8c2 != '}' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadObjectCB", `object not ended with }`)
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			return __obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		}
		if __obf_dab9baaadfa7c8c2 == '}' {
			return __obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		}
		__obf_67008a6a9e5ba828.
			ReportError("ReadObjectCB", `expect " after {, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		return false
	}
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return true // null
	}
	__obf_67008a6a9e5ba828.
		ReportError("ReadObjectCB", `expect { or n, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
	return false
}

// ReadMapCB read map with callback, the key can be any string
func (__obf_67008a6a9e5ba828 *Iterator) ReadMapCB(__obf_5fde3c82e220905a func(*Iterator, string) bool) bool {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '{' {
		if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
			return false
		}
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 == '"' {
			__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
			__obf_61998fb083387c3c := __obf_67008a6a9e5ba828.ReadString()
			if __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490() != ':' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			if !__obf_5fde3c82e220905a(__obf_67008a6a9e5ba828, __obf_61998fb083387c3c) {
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			for __obf_dab9baaadfa7c8c2 == ',' {
				__obf_61998fb083387c3c = __obf_67008a6a9e5ba828.ReadString()
				if __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490() != ':' {
					__obf_67008a6a9e5ba828.
						ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
					__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
					return false
				}
				if !__obf_5fde3c82e220905a(__obf_67008a6a9e5ba828, __obf_61998fb083387c3c) {
					__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
					return false
				}
				__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			}
			if __obf_dab9baaadfa7c8c2 != '}' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadMapCB", `object not ended with }`)
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			return __obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		}
		if __obf_dab9baaadfa7c8c2 == '}' {
			return __obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		}
		__obf_67008a6a9e5ba828.
			ReportError("ReadMapCB", `expect " after {, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
		__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		return false
	}
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return true // null
	}
	__obf_67008a6a9e5ba828.
		ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_dab9baaadfa7c8c2}))
	return false
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_b8169430213031d6() bool {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	switch __obf_dab9baaadfa7c8c2 {
	case '{':
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 == '}' {
			return false
		}
		__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
		return true
	case 'n':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return false
	}
	__obf_67008a6a9e5ba828.
		ReportError("readObjectStart", "expect { or n, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
	return false
}
