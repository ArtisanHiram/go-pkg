package __obf_c7b79b12b33d8238

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (__obf_0da8c843dad13139 *Iterator) ReadObject() (__obf_9a8638dac9e1c083 string) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	switch __obf_12577c03a12f0d2c {
	case 'n':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return "" // null
	case '{':
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c == '"' {
			__obf_0da8c843dad13139.__obf_a675366c80290b83()
			__obf_ad34f8a6de2b01cb := __obf_0da8c843dad13139.ReadString()
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			if __obf_12577c03a12f0d2c != ':' {
				__obf_0da8c843dad13139.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
			}
			return __obf_ad34f8a6de2b01cb
		}
		if __obf_12577c03a12f0d2c == '}' {
			return "" // end of object
		}
		__obf_0da8c843dad13139.
			ReportError("ReadObject", `expect " after {, but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return
	case ',':
		__obf_ad34f8a6de2b01cb := __obf_0da8c843dad13139.ReadString()
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c != ':' {
			__obf_0da8c843dad13139.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
		}
		return __obf_ad34f8a6de2b01cb
	case '}':
		return "" // end of object
	default:
		__obf_0da8c843dad13139.
			ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{__obf_12577c03a12f0d2c})))
		return
	}
}

// CaseInsensitive
func (__obf_0da8c843dad13139 *Iterator) __obf_a4ae174d25635aaf() int64 {
	__obf_f068265e9fce9202 := int64(0x811c9dc5)
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c != '"' {
		__obf_0da8c843dad13139.
			ReportError("readFieldHash", `expect ", but found `+string([]byte{__obf_12577c03a12f0d2c}))
		return 0
	}
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2;
		// require ascii string and no escape
		__obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_fa8a3db302183a92 := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			if __obf_fa8a3db302183a92 == '\\' {
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
				for _, __obf_fa8a3db302183a92 := range __obf_0da8c843dad13139.__obf_22b06ef3645a8e41() {
					if 'A' <= __obf_fa8a3db302183a92 && __obf_fa8a3db302183a92 <= 'Z' && !__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_3704f04b7ac67609 {
						__obf_fa8a3db302183a92 += 'a' - 'A'
					}
					__obf_f068265e9fce9202 ^= int64(__obf_fa8a3db302183a92)
					__obf_f068265e9fce9202 *= 0x1000193
				}
				__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
				if __obf_12577c03a12f0d2c != ':' {
					__obf_0da8c843dad13139.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_12577c03a12f0d2c}))
					return 0
				}
				return __obf_f068265e9fce9202
			}
			if __obf_fa8a3db302183a92 == '"' {
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
				__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
				if __obf_12577c03a12f0d2c != ':' {
					__obf_0da8c843dad13139.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_12577c03a12f0d2c}))
					return 0
				}
				return __obf_f068265e9fce9202
			}
			if 'A' <= __obf_fa8a3db302183a92 && __obf_fa8a3db302183a92 <= 'Z' && !__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_3704f04b7ac67609 {
				__obf_fa8a3db302183a92 += 'a' - 'A'
			}
			__obf_f068265e9fce9202 ^= int64(__obf_fa8a3db302183a92)
			__obf_f068265e9fce9202 *= 0x1000193
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			__obf_0da8c843dad13139.
				ReportError("readFieldHash", `incomplete field name`)
			return 0
		}
	}
}

func __obf_cedc170066626295(__obf_a3eaafc22faf1644 string, __obf_3704f04b7ac67609 bool) int64 {
	if !__obf_3704f04b7ac67609 {
		__obf_a3eaafc22faf1644 = strings.ToLower(__obf_a3eaafc22faf1644)
	}
	__obf_f068265e9fce9202 := int64(0x811c9dc5)
	for _, __obf_fa8a3db302183a92 := range []byte(__obf_a3eaafc22faf1644) {
		__obf_f068265e9fce9202 ^= int64(__obf_fa8a3db302183a92)
		__obf_f068265e9fce9202 *= 0x1000193
	}
	return int64(__obf_f068265e9fce9202)
}

// ReadObjectCB read object with callback, the key is ascii only and field name not copied
func (__obf_0da8c843dad13139 *Iterator) ReadObjectCB(__obf_5aafb4f722a58e01 func(*Iterator, string) bool) bool {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	var __obf_ad34f8a6de2b01cb string
	if __obf_12577c03a12f0d2c == '{' {
		if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
			return false
		}
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c == '"' {
			__obf_0da8c843dad13139.__obf_a675366c80290b83()
			__obf_ad34f8a6de2b01cb = __obf_0da8c843dad13139.ReadString()
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			if __obf_12577c03a12f0d2c != ':' {
				__obf_0da8c843dad13139.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
			}
			if !__obf_5aafb4f722a58e01(__obf_0da8c843dad13139, __obf_ad34f8a6de2b01cb) {
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			for __obf_12577c03a12f0d2c == ',' {
				__obf_ad34f8a6de2b01cb = __obf_0da8c843dad13139.ReadString()
				__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
				if __obf_12577c03a12f0d2c != ':' {
					__obf_0da8c843dad13139.
						ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
				}
				if !__obf_5aafb4f722a58e01(__obf_0da8c843dad13139, __obf_ad34f8a6de2b01cb) {
					__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
					return false
				}
				__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			}
			if __obf_12577c03a12f0d2c != '}' {
				__obf_0da8c843dad13139.
					ReportError("ReadObjectCB", `object not ended with }`)
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			return __obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		}
		if __obf_12577c03a12f0d2c == '}' {
			return __obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		}
		__obf_0da8c843dad13139.
			ReportError("ReadObjectCB", `expect " after {, but found `+string([]byte{__obf_12577c03a12f0d2c}))
		__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		return false
	}
	if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return true // null
	}
	__obf_0da8c843dad13139.
		ReportError("ReadObjectCB", `expect { or n, but found `+string([]byte{__obf_12577c03a12f0d2c}))
	return false
}

// ReadMapCB read map with callback, the key can be any string
func (__obf_0da8c843dad13139 *Iterator) ReadMapCB(__obf_5aafb4f722a58e01 func(*Iterator, string) bool) bool {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '{' {
		if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
			return false
		}
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c == '"' {
			__obf_0da8c843dad13139.__obf_a675366c80290b83()
			__obf_ad34f8a6de2b01cb := __obf_0da8c843dad13139.ReadString()
			if __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d() != ':' {
				__obf_0da8c843dad13139.
					ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			if !__obf_5aafb4f722a58e01(__obf_0da8c843dad13139, __obf_ad34f8a6de2b01cb) {
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			for __obf_12577c03a12f0d2c == ',' {
				__obf_ad34f8a6de2b01cb = __obf_0da8c843dad13139.ReadString()
				if __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d() != ':' {
					__obf_0da8c843dad13139.
						ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_12577c03a12f0d2c}))
					__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
					return false
				}
				if !__obf_5aafb4f722a58e01(__obf_0da8c843dad13139, __obf_ad34f8a6de2b01cb) {
					__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
					return false
				}
				__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			}
			if __obf_12577c03a12f0d2c != '}' {
				__obf_0da8c843dad13139.
					ReportError("ReadMapCB", `object not ended with }`)
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			return __obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		}
		if __obf_12577c03a12f0d2c == '}' {
			return __obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		}
		__obf_0da8c843dad13139.
			ReportError("ReadMapCB", `expect " after {, but found `+string([]byte{__obf_12577c03a12f0d2c}))
		__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		return false
	}
	if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return true // null
	}
	__obf_0da8c843dad13139.
		ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_12577c03a12f0d2c}))
	return false
}

func (__obf_0da8c843dad13139 *Iterator) __obf_0454a225ebb329d1() bool {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	switch __obf_12577c03a12f0d2c {
	case '{':
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c == '}' {
			return false
		}
		__obf_0da8c843dad13139.__obf_a675366c80290b83()
		return true
	case 'n':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return false
	}
	__obf_0da8c843dad13139.
		ReportError("readObjectStart", "expect { or n, but found "+string([]byte{__obf_12577c03a12f0d2c}))
	return false
}
