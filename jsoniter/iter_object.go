package __obf_91620b895eeff9ed

import (
	"fmt"
	"strings"
)

// ReadObject read one field from object.
// If object ended, returns empty string.
// Otherwise, returns the field name.
func (__obf_1bb30e8a74ed8233 *Iterator) ReadObject() (__obf_e46f5fe3db5036fe string) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	switch __obf_f16b4157911bc9af {
	case 'n':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return "" // null
	case '{':
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af == '"' {
			__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
			__obf_7e01b5b4651074d4 := __obf_1bb30e8a74ed8233.ReadString()
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			if __obf_f16b4157911bc9af != ':' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
			}
			return __obf_7e01b5b4651074d4
		}
		if __obf_f16b4157911bc9af == '}' {
			return "" // end of object
		}
		__obf_1bb30e8a74ed8233.
			ReportError("ReadObject", `expect " after {, but found `+string([]byte{__obf_f16b4157911bc9af}))
		return
	case ',':
		__obf_7e01b5b4651074d4 := __obf_1bb30e8a74ed8233.ReadString()
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af != ':' {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
		}
		return __obf_7e01b5b4651074d4
	case '}':
		return "" // end of object
	default:
		__obf_1bb30e8a74ed8233.
			ReportError("ReadObject", fmt.Sprintf(`expect { or , or } or n, but found %s`, string([]byte{__obf_f16b4157911bc9af})))
		return
	}
}

// CaseInsensitive
func (__obf_1bb30e8a74ed8233 *Iterator) __obf_d848eeb634118020() int64 {
	__obf_ca8afe3fa4d96d8f := int64(0x811c9dc5)
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af != '"' {
		__obf_1bb30e8a74ed8233.
			ReportError("readFieldHash", `expect ", but found `+string([]byte{__obf_f16b4157911bc9af}))
		return 0
	}
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21;
		// require ascii string and no escape
		__obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_1b9dc2dc0f9af749 := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			if __obf_1b9dc2dc0f9af749 == '\\' {
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
				for _, __obf_1b9dc2dc0f9af749 := range __obf_1bb30e8a74ed8233.__obf_e58c02da1e6cd884() {
					if 'A' <= __obf_1b9dc2dc0f9af749 && __obf_1b9dc2dc0f9af749 <= 'Z' && !__obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_f087a7a47617f72a {
						__obf_1b9dc2dc0f9af749 += 'a' - 'A'
					}
					__obf_ca8afe3fa4d96d8f ^= int64(__obf_1b9dc2dc0f9af749)
					__obf_ca8afe3fa4d96d8f *= 0x1000193
				}
				__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
				if __obf_f16b4157911bc9af != ':' {
					__obf_1bb30e8a74ed8233.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_f16b4157911bc9af}))
					return 0
				}
				return __obf_ca8afe3fa4d96d8f
			}
			if __obf_1b9dc2dc0f9af749 == '"' {
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
				__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
				if __obf_f16b4157911bc9af != ':' {
					__obf_1bb30e8a74ed8233.
						ReportError("readFieldHash", `expect :, but found `+string([]byte{__obf_f16b4157911bc9af}))
					return 0
				}
				return __obf_ca8afe3fa4d96d8f
			}
			if 'A' <= __obf_1b9dc2dc0f9af749 && __obf_1b9dc2dc0f9af749 <= 'Z' && !__obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_f087a7a47617f72a {
				__obf_1b9dc2dc0f9af749 += 'a' - 'A'
			}
			__obf_ca8afe3fa4d96d8f ^= int64(__obf_1b9dc2dc0f9af749)
			__obf_ca8afe3fa4d96d8f *= 0x1000193
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			__obf_1bb30e8a74ed8233.
				ReportError("readFieldHash", `incomplete field name`)
			return 0
		}
	}
}

func __obf_a34bad55ba1d026b(__obf_e91bd2feb751e4f1 string, __obf_f087a7a47617f72a bool) int64 {
	if !__obf_f087a7a47617f72a {
		__obf_e91bd2feb751e4f1 = strings.ToLower(__obf_e91bd2feb751e4f1)
	}
	__obf_ca8afe3fa4d96d8f := int64(0x811c9dc5)
	for _, __obf_1b9dc2dc0f9af749 := range []byte(__obf_e91bd2feb751e4f1) {
		__obf_ca8afe3fa4d96d8f ^= int64(__obf_1b9dc2dc0f9af749)
		__obf_ca8afe3fa4d96d8f *= 0x1000193
	}
	return int64(__obf_ca8afe3fa4d96d8f)
}

// ReadObjectCB read object with callback, the key is ascii only and field name not copied
func (__obf_1bb30e8a74ed8233 *Iterator) ReadObjectCB(__obf_acfb5a69efe27baa func(*Iterator, string) bool) bool {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	var __obf_7e01b5b4651074d4 string
	if __obf_f16b4157911bc9af == '{' {
		if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
			return false
		}
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af == '"' {
			__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
			__obf_7e01b5b4651074d4 = __obf_1bb30e8a74ed8233.ReadString()
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			if __obf_f16b4157911bc9af != ':' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
			}
			if !__obf_acfb5a69efe27baa(__obf_1bb30e8a74ed8233, __obf_7e01b5b4651074d4) {
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			for __obf_f16b4157911bc9af == ',' {
				__obf_7e01b5b4651074d4 = __obf_1bb30e8a74ed8233.ReadString()
				__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
				if __obf_f16b4157911bc9af != ':' {
					__obf_1bb30e8a74ed8233.
						ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
				}
				if !__obf_acfb5a69efe27baa(__obf_1bb30e8a74ed8233, __obf_7e01b5b4651074d4) {
					__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
					return false
				}
				__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			}
			if __obf_f16b4157911bc9af != '}' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadObjectCB", `object not ended with }`)
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			return __obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		}
		if __obf_f16b4157911bc9af == '}' {
			return __obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		}
		__obf_1bb30e8a74ed8233.
			ReportError("ReadObjectCB", `expect " after {, but found `+string([]byte{__obf_f16b4157911bc9af}))
		__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		return false
	}
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return true // null
	}
	__obf_1bb30e8a74ed8233.
		ReportError("ReadObjectCB", `expect { or n, but found `+string([]byte{__obf_f16b4157911bc9af}))
	return false
}

// ReadMapCB read map with callback, the key can be any string
func (__obf_1bb30e8a74ed8233 *Iterator) ReadMapCB(__obf_acfb5a69efe27baa func(*Iterator, string) bool) bool {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '{' {
		if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
			return false
		}
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af == '"' {
			__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
			__obf_7e01b5b4651074d4 := __obf_1bb30e8a74ed8233.ReadString()
			if __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049() != ':' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			if !__obf_acfb5a69efe27baa(__obf_1bb30e8a74ed8233, __obf_7e01b5b4651074d4) {
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			for __obf_f16b4157911bc9af == ',' {
				__obf_7e01b5b4651074d4 = __obf_1bb30e8a74ed8233.ReadString()
				if __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049() != ':' {
					__obf_1bb30e8a74ed8233.
						ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
					__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
					return false
				}
				if !__obf_acfb5a69efe27baa(__obf_1bb30e8a74ed8233, __obf_7e01b5b4651074d4) {
					__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
					return false
				}
				__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			}
			if __obf_f16b4157911bc9af != '}' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadMapCB", `object not ended with }`)
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			return __obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		}
		if __obf_f16b4157911bc9af == '}' {
			return __obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		}
		__obf_1bb30e8a74ed8233.
			ReportError("ReadMapCB", `expect " after {, but found `+string([]byte{__obf_f16b4157911bc9af}))
		__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		return false
	}
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return true // null
	}
	__obf_1bb30e8a74ed8233.
		ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_f16b4157911bc9af}))
	return false
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_613c71685c28c50b() bool {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	switch __obf_f16b4157911bc9af {
	case '{':
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af == '}' {
			return false
		}
		__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
		return true
	case 'n':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return false
	}
	__obf_1bb30e8a74ed8233.
		ReportError("readObjectStart", "expect { or n, but found "+string([]byte{__obf_f16b4157911bc9af}))
	return false
}
