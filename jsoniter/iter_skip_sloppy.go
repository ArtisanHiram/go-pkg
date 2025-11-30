//go:build jsoniter_sloppy
// +build jsoniter_sloppy

package __obf_030d39f7a8de96c6

// sloppy but faster implementation, do not validate the input json

func (__obf_4ab56a99965952e7 *Iterator) __obf_38c33066e8f2b292() {
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
			switch __obf_24309b3b0ff9ba22 {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58
				return
			}
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			return
		}
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_e8e6ee2ef6c27f13() {
	__obf_dcb1e230de7ebc6e := 1
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}
	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			switch __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58] {
			case '"':
				__obf_4ab56a99965952e7. // If inside string, skip it
				__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
				__obf_4ab56a99965952e7.__obf_ab668369e47b63de()
				__obf_82c6e05b9d226c58 = __obf_4ab56a99965952e7. // it will be i++ soon
				__obf_5e22d6270d27491f - 1
			case '[':
				__obf_dcb1e230de7ebc6e++// If open symbol, increase level

				if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
					return
				}
			case ']':
				__obf_dcb1e230de7ebc6e--// If close symbol, increase level

				if !__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_dcb1e230de7ebc6e == 0 {
					__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
					return
				}
			}
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			__obf_4ab56a99965952e7.
				ReportError("skipObject", "incomplete array")
			return
		}
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_56cafb6b2d847244() {
	__obf_dcb1e230de7ebc6e := 1
	if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
		return
	}

	for {
		for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
			switch __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58] {
			case '"':
				__obf_4ab56a99965952e7. // If inside string, skip it
				__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
				__obf_4ab56a99965952e7.__obf_ab668369e47b63de()
				__obf_82c6e05b9d226c58 = __obf_4ab56a99965952e7. // it will be i++ soon
				__obf_5e22d6270d27491f - 1
			case '{':
				__obf_dcb1e230de7ebc6e++// If open symbol, increase level

				if !__obf_4ab56a99965952e7.__obf_9cff3330bd1a56ea() {
					return
				}
			case '}':
				__obf_dcb1e230de7ebc6e--// If close symbol, increase level

				if !__obf_4ab56a99965952e7.__obf_6502e8959d5f85cf() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_dcb1e230de7ebc6e == 0 {
					__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_82c6e05b9d226c58 + 1
					return
				}
			}
		}
		if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			__obf_4ab56a99965952e7.
				ReportError("skipObject", "incomplete object")
			return
		}
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_ab668369e47b63de() {
	for {
		__obf_27a52b85a4f4f7ac, __obf_72f4ef889956217d := __obf_4ab56a99965952e7.__obf_19c716ca0c71a9ed()
		if __obf_27a52b85a4f4f7ac == -1 {
			if !__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
				__obf_4ab56a99965952e7.
					ReportError("skipString", "incomplete string")
				return
			}
			if __obf_72f4ef889956217d {
				__obf_4ab56a99965952e7.
				// skip the first char as last char read is \
				__obf_5e22d6270d27491f = 1
			}
		} else {
			__obf_4ab56a99965952e7.__obf_5e22d6270d27491f = __obf_27a52b85a4f4f7ac
			return
		}
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (__obf_4ab56a99965952e7 *Iterator) __obf_19c716ca0c71a9ed() (int, bool) {
	__obf_72f4ef889956217d := false
	for __obf_82c6e05b9d226c58 := __obf_4ab56a99965952e7.__obf_5e22d6270d27491f; __obf_82c6e05b9d226c58 < __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc; __obf_82c6e05b9d226c58++ {
		__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_82c6e05b9d226c58]
		if __obf_24309b3b0ff9ba22 == '"' {
			if !__obf_72f4ef889956217d {
				return __obf_82c6e05b9d226c58 + 1, false
			}
			__obf_354f1fc1cab1d938 := __obf_82c6e05b9d226c58 - 1
			for {
				if __obf_354f1fc1cab1d938 < __obf_4ab56a99965952e7.__obf_5e22d6270d27491f || __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_354f1fc1cab1d938] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return __obf_82c6e05b9d226c58 + 1, true
				}
				__obf_354f1fc1cab1d938--
				if __obf_354f1fc1cab1d938 < __obf_4ab56a99965952e7.__obf_5e22d6270d27491f || __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_354f1fc1cab1d938] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				__obf_354f1fc1cab1d938--
			}
		} else if __obf_24309b3b0ff9ba22 == '\\' {
			__obf_72f4ef889956217d = true
		}
	}
	__obf_354f1fc1cab1d938 := __obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc - 1
	for {
		if __obf_354f1fc1cab1d938 < __obf_4ab56a99965952e7.__obf_5e22d6270d27491f || __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_354f1fc1cab1d938] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		__obf_354f1fc1cab1d938--
		if __obf_354f1fc1cab1d938 < __obf_4ab56a99965952e7.__obf_5e22d6270d27491f || __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_354f1fc1cab1d938] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		__obf_354f1fc1cab1d938--

	}
	return -1, true // end with \
}
