package __obf_030d39f7a8de96c6

// ReadArray read array element, tells if the array has more element to read.
func (__obf_4ab56a99965952e7 *Iterator) ReadArray() (__obf_59c72400ec1a6110 bool) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	switch __obf_24309b3b0ff9ba22 {
	case 'n':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return false // null
	case '[':
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 != ']' {
			__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
			return true
		}
		return false
	case ']':
		return false
	case ',':
		return true
	default:
		__obf_4ab56a99965952e7.
			ReportError("ReadArray", "expect [ or , or ] or n, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
		return
	}
}

// ReadArrayCB read array with callback
func (__obf_4ab56a99965952e7 *Iterator) ReadArrayCB(__obf_a2c838c9bdc87499 func(*Iterator) bool) (__obf_59c72400ec1a6110 bool) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == '[' {
		if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
			return false
		}
		__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
		if __obf_24309b3b0ff9ba22 != ']' {
			__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
			if !__obf_a2c838c9bdc87499(__obf_4ab56a99965952e7) {
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			for __obf_24309b3b0ff9ba22 == ',' {
				if !__obf_a2c838c9bdc87499(__obf_4ab56a99965952e7) {
					__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
					return false
				}
				__obf_24309b3b0ff9ba22 = __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
			}
			if __obf_24309b3b0ff9ba22 != ']' {
				__obf_4ab56a99965952e7.
					ReportError("ReadArrayCB", "expect ] in the end, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
				__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
				return false
			}
			return __obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
		}
		return __obf_4ab56a99965952e7.__obf_6502e8959d5f85cf()
	}
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l')
		return true // null
	}
	__obf_4ab56a99965952e7.
		ReportError("ReadArrayCB", "expect [ or n, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
	return false
}
