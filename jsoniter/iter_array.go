package __obf_5b802ce8d9ba56d6

// ReadArray read array element, tells if the array has more element to read.
func (__obf_67008a6a9e5ba828 *Iterator) ReadArray() (__obf_5dabcdfee5097ed6 bool) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	switch __obf_dab9baaadfa7c8c2 {
	case 'n':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return false // null
	case '[':
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 != ']' {
			__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		__obf_67008a6a9e5ba828.
			ReportError("ReadArray", "expect [ or , or ] or n, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
		return
	}
}

// ReadArrayCB read array with callback
func (__obf_67008a6a9e5ba828 *Iterator) ReadArrayCB(__obf_5fde3c82e220905a func(*Iterator) bool) (__obf_5dabcdfee5097ed6 bool) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == '[' {
		if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
			return false
		}
		__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
		if __obf_dab9baaadfa7c8c2 != ']' {
			__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
			if !__obf_5fde3c82e220905a(__obf_67008a6a9e5ba828) {
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			for __obf_dab9baaadfa7c8c2 == ',' {
				if !__obf_5fde3c82e220905a(__obf_67008a6a9e5ba828) {
					__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
					return false
				}
				__obf_dab9baaadfa7c8c2 = __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
			}
			if __obf_dab9baaadfa7c8c2 != ']' {
				__obf_67008a6a9e5ba828.
					ReportError("ReadArrayCB", "expect ] in the end, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
				__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
				return false
			}
			return __obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
		}
		return __obf_67008a6a9e5ba828.__obf_c37c04dba052f09d()
	}
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l')
		return true // null
	}
	__obf_67008a6a9e5ba828.
		ReportError("ReadArrayCB", "expect [ or n, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
	return false
}
