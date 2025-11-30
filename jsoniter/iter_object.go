package __obf_703d23ba520c3295

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (__obf_47edab4c16a2d88d *Iterator) ReadObject() (__obf_551cbec38242ce0c string) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	switch __obf_bd08f5161fff294a {
	case 'n':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return "" // null
	case '{':
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a == '"' {
			__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
			__obf_13f6440419533990 := __obf_47edab4c16a2d88d.ReadString()
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			if __obf_bd08f5161fff294a != ':' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
			}
			return __obf_13f6440419533990
		}
		if __obf_bd08f5161fff294a == '}' {
			return "" // end of object
		}
		__obf_47edab4c16a2d88d.
			ReportError("ReadObject", `expect " after {, but found `+string([]byte{__obf_bd08f5161fff294a}))
		return
	case ',':
		__obf_13f6440419533990 := __obf_47edab4c16a2d88d.ReadString()
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a != ':' {
			__obf_47edab4c16a2d88d.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
		}
		return __obf_13f6440419533990
	case '}':
		return "" // end of object
	default:
		__obf_47edab4c16a2d88d.
			ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{__obf_bd08f5161fff294a})))
		return
	}
}

// CaseInsensitive
func (__obf_47edab4c16a2d88d *Iterator) __obf_0ede172a7a3a6a62() int64 {
	__obf_7b13e047c9275f66 := int64(0x811c9dc5)
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a != '"' {
		__obf_47edab4c16a2d88d.
			ReportError("readFieldHash", `expect ", but found `+string([]byte{__obf_bd08f5161fff294a}))
		return 0
	}
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11;
		// require ascii string and no escape
		__obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_85a417ca3a5d43c2 := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			if __obf_85a417ca3a5d43c2 == '\\' {
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
				for _, __obf_85a417ca3a5d43c2 := range __obf_47edab4c16a2d88d.__obf_61de6bc49a932c21() {
					if 'A' <= __obf_85a417ca3a5d43c2 && __obf_85a417ca3a5d43c2 <= 'Z' && !__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_9332366e75c2e023 {
						__obf_85a417ca3a5d43c2 += 'a' - 'A'
					}
					__obf_7b13e047c9275f66 ^= int64(__obf_85a417ca3a5d43c2)
					__obf_7b13e047c9275f66 *= 0x1000193
				}
				__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
				if __obf_bd08f5161fff294a != ':' {
					__obf_47edab4c16a2d88d.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_bd08f5161fff294a}))
					return 0
				}
				return __obf_7b13e047c9275f66
			}
			if __obf_85a417ca3a5d43c2 == '"' {
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
				__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
				if __obf_bd08f5161fff294a != ':' {
					__obf_47edab4c16a2d88d.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_bd08f5161fff294a}))
					return 0
				}
				return __obf_7b13e047c9275f66
			}
			if 'A' <= __obf_85a417ca3a5d43c2 && __obf_85a417ca3a5d43c2 <= 'Z' && !__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_9332366e75c2e023 {
				__obf_85a417ca3a5d43c2 += 'a' - 'A'
			}
			__obf_7b13e047c9275f66 ^= int64(__obf_85a417ca3a5d43c2)
			__obf_7b13e047c9275f66 *= 0x1000193
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			__obf_47edab4c16a2d88d.
				ReportError("readFieldHash", `incomplete field name`)
			return 0
		}
	}
}

func __obf_818f762e1b421403(__obf_44b48c26051f8033 string, __obf_9332366e75c2e023 bool) int64 {
	if !__obf_9332366e75c2e023 {
		__obf_44b48c26051f8033 = strings.ToLower(__obf_44b48c26051f8033)
	}
	__obf_7b13e047c9275f66 := int64(0x811c9dc5)
	for _, __obf_85a417ca3a5d43c2 := range []byte(__obf_44b48c26051f8033) {
		__obf_7b13e047c9275f66 ^= int64(__obf_85a417ca3a5d43c2)
		__obf_7b13e047c9275f66 *= 0x1000193
	}
	return int64(__obf_7b13e047c9275f66)
}

// ReadObjectCB read object with callback, the key is ascii only and field name not copied
func (__obf_47edab4c16a2d88d *Iterator) ReadObjectCB(__obf_22c8e049442f16aa func(*Iterator, string) bool) bool {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	var __obf_13f6440419533990 string
	if __obf_bd08f5161fff294a == '{' {
		if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
			return false
		}
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a == '"' {
			__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
			__obf_13f6440419533990 = __obf_47edab4c16a2d88d.ReadString()
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			if __obf_bd08f5161fff294a != ':' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
			}
			if !__obf_22c8e049442f16aa(__obf_47edab4c16a2d88d, __obf_13f6440419533990) {
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			for __obf_bd08f5161fff294a == ',' {
				__obf_13f6440419533990 = __obf_47edab4c16a2d88d.ReadString()
				__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
				if __obf_bd08f5161fff294a != ':' {
					__obf_47edab4c16a2d88d.
						ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
				}
				if !__obf_22c8e049442f16aa(__obf_47edab4c16a2d88d, __obf_13f6440419533990) {
					__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
					return false
				}
				__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			}
			if __obf_bd08f5161fff294a != '}' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadObjectCB", `object not ended with }`)
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			return __obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		}
		if __obf_bd08f5161fff294a == '}' {
			return __obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		}
		__obf_47edab4c16a2d88d.
			ReportError("ReadObjectCB", `expect " after {, but found `+string([]byte{__obf_bd08f5161fff294a}))
		__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		return false
	}
	if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return true // null
	}
	__obf_47edab4c16a2d88d.
		ReportError("ReadObjectCB", `expect { or n, but found `+string([]byte{__obf_bd08f5161fff294a}))
	return false
}

// ReadMapCB read map with callback, the key can be any string
func (__obf_47edab4c16a2d88d *Iterator) ReadMapCB(__obf_22c8e049442f16aa func(*Iterator, string) bool) bool {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '{' {
		if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
			return false
		}
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a == '"' {
			__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
			__obf_13f6440419533990 := __obf_47edab4c16a2d88d.ReadString()
			if __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec() != ':' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			if !__obf_22c8e049442f16aa(__obf_47edab4c16a2d88d, __obf_13f6440419533990) {
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			for __obf_bd08f5161fff294a == ',' {
				__obf_13f6440419533990 = __obf_47edab4c16a2d88d.ReadString()
				if __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec() != ':' {
					__obf_47edab4c16a2d88d.
						ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
					__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
					return false
				}
				if !__obf_22c8e049442f16aa(__obf_47edab4c16a2d88d, __obf_13f6440419533990) {
					__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
					return false
				}
				__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			}
			if __obf_bd08f5161fff294a != '}' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadMapCB", `object not ended with }`)
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			return __obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		}
		if __obf_bd08f5161fff294a == '}' {
			return __obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		}
		__obf_47edab4c16a2d88d.
			ReportError("ReadMapCB", `expect " after {, but found `+string([]byte{__obf_bd08f5161fff294a}))
		__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		return false
	}
	if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return true // null
	}
	__obf_47edab4c16a2d88d.
		ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_bd08f5161fff294a}))
	return false
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_000080914c527b80() bool {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	switch __obf_bd08f5161fff294a {
	case '{':
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a == '}' {
			return false
		}
		__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
		return true
	case 'n':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return false
	}
	__obf_47edab4c16a2d88d.
		ReportError("readObjectStart", "expect { or n, but found "+string([]byte{__obf_bd08f5161fff294a}))
	return false
}
