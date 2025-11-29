package __obf_91620b895eeff9ed

// ReadArray read array element, tells if the array has more element to read.
func (__obf_1bb30e8a74ed8233 *Iterator) ReadArray() (__obf_e46f5fe3db5036fe bool) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	switch __obf_f16b4157911bc9af {
	case 'n':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return false // null
	case '[':
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af != ']' {
			__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		__obf_1bb30e8a74ed8233.
			ReportError("ReadArray", "expect [ or , or ] or n, but found "+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
}

// ReadArrayCB read array with callback
func (__obf_1bb30e8a74ed8233 *Iterator) ReadArrayCB(__obf_acfb5a69efe27baa func(*Iterator) bool) (__obf_e46f5fe3db5036fe bool) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == '[' {
		if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
			return false
		}
		__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af != ']' {
			__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
			if !__obf_acfb5a69efe27baa(__obf_1bb30e8a74ed8233) {
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			for __obf_f16b4157911bc9af == ',' {
				if !__obf_acfb5a69efe27baa(__obf_1bb30e8a74ed8233) {
					__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
					return false
				}
				__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
			}
			if __obf_f16b4157911bc9af != ']' {
				__obf_1bb30e8a74ed8233.
					ReportError("ReadArrayCB", "expect ] in the end, but found "+string([]byte{__obf_f16b4157911bc9af}))
				__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
				return false
			}
			return __obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
		}
		return __obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
	}
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l')
		return true // null
	}
	__obf_1bb30e8a74ed8233.
		ReportError("ReadArrayCB", "expect [ or n, but found "+string([]byte{__obf_f16b4157911bc9af}))
	return false
}
