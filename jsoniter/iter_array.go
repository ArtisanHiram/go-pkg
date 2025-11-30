package __obf_703d23ba520c3295

// ReadArray read array element, tells if the array has more element to read.
func (__obf_47edab4c16a2d88d *Iterator) ReadArray() (__obf_551cbec38242ce0c bool) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	switch __obf_bd08f5161fff294a {
	case 'n':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return false // null
	case '[':
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a != ']' {
			__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		__obf_47edab4c16a2d88d.
			ReportError("ReadArray", "expect [ or , or ] or n, but found "+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
}

// ReadArrayCB read array with callback
func (__obf_47edab4c16a2d88d *Iterator) ReadArrayCB(__obf_22c8e049442f16aa func(*Iterator) bool) (__obf_551cbec38242ce0c bool) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '[' {
		if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
			return false
		}
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a != ']' {
			__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
			if !__obf_22c8e049442f16aa(__obf_47edab4c16a2d88d) {
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			for __obf_bd08f5161fff294a == ',' {
				if !__obf_22c8e049442f16aa(__obf_47edab4c16a2d88d) {
					__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
					return false
				}
				__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
			}
			if __obf_bd08f5161fff294a != ']' {
				__obf_47edab4c16a2d88d.
					ReportError("ReadArrayCB", "expect ] in the end, but found "+string([]byte{__obf_bd08f5161fff294a}))
				__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
				return false
			}
			return __obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
		}
		return __obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2()
	}
	if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		return true // null
	}
	__obf_47edab4c16a2d88d.
		ReportError("ReadArrayCB", "expect [ or n, but found "+string([]byte{__obf_bd08f5161fff294a}))
	return false
}
