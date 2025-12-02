package __obf_c7b79b12b33d8238

// ReadArray read array element, tells if the array has more element to read.
func (__obf_0da8c843dad13139 *Iterator) ReadArray() (__obf_9a8638dac9e1c083 bool) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	switch __obf_12577c03a12f0d2c {
	case 'n':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return false // null
	case '[':
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c != ']' {
			__obf_0da8c843dad13139.__obf_a675366c80290b83()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		__obf_0da8c843dad13139.
			ReportError("ReadArray", "expect [ or , or ] or n, but found "+string([]byte{__obf_12577c03a12f0d2c}))
		return
	}
}

// ReadArrayCB read array with callback
func (__obf_0da8c843dad13139 *Iterator) ReadArrayCB(__obf_5aafb4f722a58e01 func(*Iterator) bool) (__obf_9a8638dac9e1c083 bool) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == '[' {
		if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
			return false
		}
		__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
		if __obf_12577c03a12f0d2c != ']' {
			__obf_0da8c843dad13139.__obf_a675366c80290b83()
			if !__obf_5aafb4f722a58e01(__obf_0da8c843dad13139) {
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			for __obf_12577c03a12f0d2c == ',' {
				if !__obf_5aafb4f722a58e01(__obf_0da8c843dad13139) {
					__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
					return false
				}
				__obf_12577c03a12f0d2c = __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
			}
			if __obf_12577c03a12f0d2c != ']' {
				__obf_0da8c843dad13139.
					ReportError("ReadArrayCB", "expect ] in the end, but found "+string([]byte{__obf_12577c03a12f0d2c}))
				__obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
				return false
			}
			return __obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
		}
		return __obf_0da8c843dad13139.__obf_aed83e308a23ff9c()
	}
	if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l')
		return true // null
	}
	__obf_0da8c843dad13139.
		ReportError("ReadArrayCB", "expect [ or n, but found "+string([]byte{__obf_12577c03a12f0d2c}))
	return false
}
