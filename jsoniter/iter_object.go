package __obf_030d39f7a8de96c6

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (__obf_4ab56a99965952e7 *Iterator) ReadObject() (__obf_59c72400ec1a6110 string) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	switch __obf_24309b3b0ff9ba22 {
	case 'n':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return "" // null
	case '{':
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 == '"' {
			__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
			__obf_cd4d02f264b18d55 := __obf_4ab56a99965952e7.ReadString()
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			if __obf_24309b3b0ff9ba22 != ':' {
				__obf_4ab56a99965952e7.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
			}
			return __obf_cd4d02f264b18d55
		}
		if __obf_24309b3b0ff9ba22 == '}' {
			return "" // end of object
		}
		__obf_4ab56a99965952e7.
			ReportError("ReadObject", `expect " after {, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	case ',':
		__obf_cd4d02f264b18d55 := __obf_4ab56a99965952e7.ReadString()
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 != ':' {
			__obf_4ab56a99965952e7.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		}
		return __obf_cd4d02f264b18d55
	case '}':
		return "" // end of object
	default:
		__obf_4ab56a99965952e7.
			ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{__obf_24309b3b0ff9ba22})))
		return
	}
}

// CaseInsensitive
func (__obf_4ab56a99965952e7 *Iterator) __obf_98be9914bcefeaa7() int64 {
	__obf_025480a3c400c918 := int64(0x811c9dc5)
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 != '"' {
		__obf_4ab56a99965952e7.
			ReportError("readFieldHash", `expect ", but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		return 0
	}
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f;
		// require ascii string and no escape
		__obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_ea107e11b66371dd := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			if __obf_ea107e11b66371dd == '\\' {
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
				for _, __obf_ea107e11b66371dd := range __obf_4ab56a99965952e7.__obf_506d8d9b7d95ca67() {
					if 'A' <= __obf_ea107e11b66371dd && __obf_ea107e11b66371dd <= 'Z' && !__obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_af13e3babb780a4e {
						__obf_ea107e11b66371dd += 'a' - 'A'
					}
					__obf_025480a3c400c918 ^= int64(__obf_ea107e11b66371dd)
					__obf_025480a3c400c918 *= 0x1000193
				}
				__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
				if __obf_24309b3b0ff9ba22 != ':' {
					__obf_4ab56a99965952e7.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
					return 0
				}
				return __obf_025480a3c400c918
			}
			if __obf_ea107e11b66371dd == '"' {
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
				__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
				if __obf_24309b3b0ff9ba22 != ':' {
					__obf_4ab56a99965952e7.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
					return 0
				}
				return __obf_025480a3c400c918
			}
			if 'A' <= __obf_ea107e11b66371dd && __obf_ea107e11b66371dd <= 'Z' && !__obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_af13e3babb780a4e {
				__obf_ea107e11b66371dd += 'a' - 'A'
			}
			__obf_025480a3c400c918 ^= int64(__obf_ea107e11b66371dd)
			__obf_025480a3c400c918 *= 0x1000193
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			__obf_4ab56a99965952e7.
				ReportError("readFieldHash", `incomplete field name`)
			return 0
		}
	}
}

func __obf_47e34696e8c7ae78(__obf_428c3b4a95725c4a string, __obf_af13e3babb780a4e bool) int64 {
	if !__obf_af13e3babb780a4e {
		__obf_428c3b4a95725c4a = strings.ToLower(__obf_428c3b4a95725c4a)
	}
	__obf_025480a3c400c918 := int64(0x811c9dc5)
	for _, __obf_ea107e11b66371dd := range []byte(__obf_428c3b4a95725c4a) {
		__obf_025480a3c400c918 ^= int64(__obf_ea107e11b66371dd)
		__obf_025480a3c400c918 *= 0x1000193
	}
	return int64(__obf_025480a3c400c918)
}

// ReadObjectCB read object with callback, the key is ascii only and field name not copied
func (__obf_4ab56a99965952e7 *Iterator) ReadObjectCB(__obf_a2c838c9bdc87499 func(*Iterator, string) bool) bool {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	var __obf_cd4d02f264b18d55 string
	if __obf_24309b3b0ff9ba22 == '{' {
		if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
			return false
		}
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 == '"' {
			__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
			__obf_cd4d02f264b18d55 = __obf_4ab56a99965952e7.ReadString()
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			if __obf_24309b3b0ff9ba22 != ':' {
				__obf_4ab56a99965952e7.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
			}
			if !__obf_a2c838c9bdc87499(__obf_4ab56a99965952e7, __obf_cd4d02f264b18d55) {
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			for __obf_24309b3b0ff9ba22 == ',' {
				__obf_cd4d02f264b18d55 = __obf_4ab56a99965952e7.ReadString()
				__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
				if __obf_24309b3b0ff9ba22 != ':' {
					__obf_4ab56a99965952e7.
						ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
				}
				if !__obf_a2c838c9bdc87499(__obf_4ab56a99965952e7, __obf_cd4d02f264b18d55) {
					__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
					return false
				}
				__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			}
			if __obf_24309b3b0ff9ba22 != '}' {
				__obf_4ab56a99965952e7.
					ReportError("ReadObjectCB", `object not ended with }`)
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			return __obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		}
		if __obf_24309b3b0ff9ba22 == '}' {
			return __obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		}
		__obf_4ab56a99965952e7.
			ReportError("ReadObjectCB", `expect " after {, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		return false
	}
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return true // null
	}
	__obf_4ab56a99965952e7.
		ReportError("ReadObjectCB", `expect { or n, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
	return false
}

// ReadMapCB read map with callback, the key can be any string
func (__obf_4ab56a99965952e7 *Iterator) ReadMapCB(__obf_a2c838c9bdc87499 func(*Iterator, string) bool) bool {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '{' {
		if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
			return false
		}
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 == '"' {
			__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
			__obf_cd4d02f264b18d55 := __obf_4ab56a99965952e7.ReadString()
			if __obf_4ab56a99965952e7.__obf_61df301d1f67ad73() != ':' {
				__obf_4ab56a99965952e7.
					ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			if !__obf_a2c838c9bdc87499(__obf_4ab56a99965952e7, __obf_cd4d02f264b18d55) {
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			for __obf_24309b3b0ff9ba22 == ',' {
				__obf_cd4d02f264b18d55 = __obf_4ab56a99965952e7.ReadString()
				if __obf_4ab56a99965952e7.__obf_61df301d1f67ad73() != ':' {
					__obf_4ab56a99965952e7.
						ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
					__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
					return false
				}
				if !__obf_a2c838c9bdc87499(__obf_4ab56a99965952e7, __obf_cd4d02f264b18d55) {
					__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
					return false
				}
				__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			}
			if __obf_24309b3b0ff9ba22 != '}' {
				__obf_4ab56a99965952e7.
					ReportError("ReadMapCB", `object not ended with }`)
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			return __obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		}
		if __obf_24309b3b0ff9ba22 == '}' {
			return __obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		}
		__obf_4ab56a99965952e7.
			ReportError("ReadMapCB", `expect " after {, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
		__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		return false
	}
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return true // null
	}
	__obf_4ab56a99965952e7.
		ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_24309b3b0ff9ba22}))
	return false
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_2c0c8655dcf5f087() bool {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	switch __obf_24309b3b0ff9ba22 {
	case '{':
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 == '}' {
			return false
		}
		__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
		return true
	case 'n':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return false
	}
	__obf_4ab56a99965952e7.
		ReportError("readObjectStart", "expect { or n, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
	return false
}
