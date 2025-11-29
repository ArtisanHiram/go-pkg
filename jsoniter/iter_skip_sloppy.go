//go:build jsoniter_sloppy
// +build jsoniter_sloppy

package __obf_5b802ce8d9ba56d6

// sloppy but faster implementation, do not validate the input json

func (__obf_67008a6a9e5ba828 *Iterator) __obf_a64603a02c650185() {
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
			switch __obf_dab9baaadfa7c8c2 {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3
				return
			}
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			return
		}
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_1d506622058b77ef() {
	__obf_e85cc6c8350cdebb := 1
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}
	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			switch __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3] {
			case '"':
				__obf_67008a6a9e5ba828. // If inside string, skip it
				__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
				__obf_67008a6a9e5ba828.__obf_acb079d30d4527e3()
				__obf_2deec7c38ffb6ae3 = __obf_67008a6a9e5ba828. // it will be i++ soon
				__obf_14babd6f9a55bd36 - 1
			case '[':
				__obf_e85cc6c8350cdebb++// If open symbol, increase level

				if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
					return
				}
			case ']':
				__obf_e85cc6c8350cdebb--// If close symbol, increase level

				if !__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_e85cc6c8350cdebb == 0 {
					__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
					return
				}
			}
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			__obf_67008a6a9e5ba828.
				ReportError("skipObject", "incomplete array")
			return
		}
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_7e3d8ea4ab2a8258() {
	__obf_e85cc6c8350cdebb := 1
	if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
		return
	}

	for {
		for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
			switch __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3] {
			case '"':
				__obf_67008a6a9e5ba828. // If inside string, skip it
				__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
				__obf_67008a6a9e5ba828.__obf_acb079d30d4527e3()
				__obf_2deec7c38ffb6ae3 = __obf_67008a6a9e5ba828. // it will be i++ soon
				__obf_14babd6f9a55bd36 - 1
			case '{':
				__obf_e85cc6c8350cdebb++// If open symbol, increase level

				if !__obf_67008a6a9e5ba828.__obf_093e3593687574f7() {
					return
				}
			case '}':
				__obf_e85cc6c8350cdebb--// If close symbol, increase level

				if !__obf_67008a6a9e5ba828.__obf_c37c04dba052f09d() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_e85cc6c8350cdebb == 0 {
					__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_2deec7c38ffb6ae3 + 1
					return
				}
			}
		}
		if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			__obf_67008a6a9e5ba828.
				ReportError("skipObject", "incomplete object")
			return
		}
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_acb079d30d4527e3() {
	for {
		__obf_ffcf0bbba1f2673c, __obf_fa25ec78f14d411b := __obf_67008a6a9e5ba828.__obf_a461e328d5833193()
		if __obf_ffcf0bbba1f2673c == -1 {
			if !__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
				__obf_67008a6a9e5ba828.
					ReportError("skipString", "incomplete string")
				return
			}
			if __obf_fa25ec78f14d411b {
				__obf_67008a6a9e5ba828.
				// skip the first char as last char read is \
				__obf_14babd6f9a55bd36 = 1
			}
		} else {
			__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 = __obf_ffcf0bbba1f2673c
			return
		}
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (__obf_67008a6a9e5ba828 *Iterator) __obf_a461e328d5833193() (int, bool) {
	__obf_fa25ec78f14d411b := false
	for __obf_2deec7c38ffb6ae3 := __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36; __obf_2deec7c38ffb6ae3 < __obf_67008a6a9e5ba828.__obf_3a36550914545c79; __obf_2deec7c38ffb6ae3++ {
		__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_2deec7c38ffb6ae3]
		if __obf_dab9baaadfa7c8c2 == '"' {
			if !__obf_fa25ec78f14d411b {
				return __obf_2deec7c38ffb6ae3 + 1, false
			}
			__obf_973404809dee3093 := __obf_2deec7c38ffb6ae3 - 1
			for {
				if __obf_973404809dee3093 < __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 || __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_973404809dee3093] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return __obf_2deec7c38ffb6ae3 + 1, true
				}
				__obf_973404809dee3093--
				if __obf_973404809dee3093 < __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 || __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_973404809dee3093] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				__obf_973404809dee3093--
			}
		} else if __obf_dab9baaadfa7c8c2 == '\\' {
			__obf_fa25ec78f14d411b = true
		}
	}
	__obf_973404809dee3093 := __obf_67008a6a9e5ba828.__obf_3a36550914545c79 - 1
	for {
		if __obf_973404809dee3093 < __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 || __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_973404809dee3093] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		__obf_973404809dee3093--
		if __obf_973404809dee3093 < __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 || __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_973404809dee3093] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		__obf_973404809dee3093--

	}
	return -1, true // end with \
}
