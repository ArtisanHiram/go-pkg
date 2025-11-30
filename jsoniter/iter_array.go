package __obf_c3cd04a15c56f32f

// ReadArray read array element, tells if the array has more element to read.
func (__obf_edd9fe48d39445e4 *Iterator) ReadArray() (__obf_316af79472661247 bool) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	switch __obf_0c1bc1e511a43120 {
	case 'n':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return false // null
	case '[':
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 != ']' {
			__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		__obf_edd9fe48d39445e4.
			ReportError("ReadArray", "expect [ or , or ] or n, but found "+string([]byte{__obf_0c1bc1e511a43120}))
		return
	}
}

// ReadArrayCB read array with callback
func (__obf_edd9fe48d39445e4 *Iterator) ReadArrayCB(__obf_d1ec68aebedafce9 func(*Iterator) bool) (__obf_316af79472661247 bool) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == '[' {
		if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
			return false
		}
		__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
		if __obf_0c1bc1e511a43120 != ']' {
			__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
			if !__obf_d1ec68aebedafce9(__obf_edd9fe48d39445e4) {
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			for __obf_0c1bc1e511a43120 == ',' {
				if !__obf_d1ec68aebedafce9(__obf_edd9fe48d39445e4) {
					__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
					return false
				}
				__obf_0c1bc1e511a43120 = __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
			}
			if __obf_0c1bc1e511a43120 != ']' {
				__obf_edd9fe48d39445e4.
					ReportError("ReadArrayCB", "expect ] in the end, but found "+string([]byte{__obf_0c1bc1e511a43120}))
				__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
				return false
			}
			return __obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
		}
		return __obf_edd9fe48d39445e4.__obf_0727968e0120c8a6()
	}
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l')
		return true // null
	}
	__obf_edd9fe48d39445e4.
		ReportError("ReadArrayCB", "expect [ or n, but found "+string([]byte{__obf_0c1bc1e511a43120}))
	return false
}
