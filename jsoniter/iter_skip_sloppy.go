//go:build jsoniter_sloppy
// +build jsoniter_sloppy

package __obf_91620b895eeff9ed

// sloppy but faster implementation, do not validate the input json

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_d6a474bd319b770c() {
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
			switch __obf_f16b4157911bc9af {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182
				return
			}
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			return
		}
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_531ac82b5de39aec() {
	__obf_0daa2d992c59d1ec := 1
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			switch __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182] {
			case '"':
				__obf_1bb30e8a74ed8233. // If inside string, skip it
				__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
				__obf_1bb30e8a74ed8233.__obf_c5493c5b3916fc3f()
				__obf_5aa5c8829b97f182 = __obf_1bb30e8a74ed8233. // it will be i++ soon
				__obf_a657fb48fcb34e21 - 1
			case '[':
				__obf_0daa2d992c59d1ec++// If open symbol, increase level

				if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
					return
				}
			case ']':
				__obf_0daa2d992c59d1ec--// If close symbol, increase level

				if !__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_0daa2d992c59d1ec == 0 {
					__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
					return
				}
			}
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			__obf_1bb30e8a74ed8233.
				ReportError("skipObject", "incomplete array")
			return
		}
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_8d12ebf64f0517d2() {
	__obf_0daa2d992c59d1ec := 1
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}

	for {
		for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
			switch __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182] {
			case '"':
				__obf_1bb30e8a74ed8233. // If inside string, skip it
				__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
				__obf_1bb30e8a74ed8233.__obf_c5493c5b3916fc3f()
				__obf_5aa5c8829b97f182 = __obf_1bb30e8a74ed8233. // it will be i++ soon
				__obf_a657fb48fcb34e21 - 1
			case '{':
				__obf_0daa2d992c59d1ec++// If open symbol, increase level

				if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
					return
				}
			case '}':
				__obf_0daa2d992c59d1ec--// If close symbol, increase level

				if !__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_0daa2d992c59d1ec == 0 {
					__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_5aa5c8829b97f182 + 1
					return
				}
			}
		}
		if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			__obf_1bb30e8a74ed8233.
				ReportError("skipObject", "incomplete object")
			return
		}
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_c5493c5b3916fc3f() {
	for {
		__obf_81a7bfce22c2a9fa, __obf_23aafe4d90b76633 := __obf_1bb30e8a74ed8233.__obf_5898ed66ccba8f31()
		if __obf_81a7bfce22c2a9fa == -1 {
			if !__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
				__obf_1bb30e8a74ed8233.
					ReportError("skipString", "incomplete string")
				return
			}
			if __obf_23aafe4d90b76633 {
				__obf_1bb30e8a74ed8233.
				// skip the first char as last char read is \
				__obf_a657fb48fcb34e21 = 1
			}
		} else {
			__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 = __obf_81a7bfce22c2a9fa
			return
		}
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (__obf_1bb30e8a74ed8233 *Iterator) __obf_5898ed66ccba8f31() (int, bool) {
	__obf_23aafe4d90b76633 := false
	for __obf_5aa5c8829b97f182 := __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21; __obf_5aa5c8829b97f182 < __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae; __obf_5aa5c8829b97f182++ {
		__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_5aa5c8829b97f182]
		if __obf_f16b4157911bc9af == '"' {
			if !__obf_23aafe4d90b76633 {
				return __obf_5aa5c8829b97f182 + 1, false
			}
			__obf_7b29092764c6c9cb := __obf_5aa5c8829b97f182 - 1
			for {
				if __obf_7b29092764c6c9cb < __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 || __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_7b29092764c6c9cb] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return __obf_5aa5c8829b97f182 + 1, true
				}
				__obf_7b29092764c6c9cb--
				if __obf_7b29092764c6c9cb < __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 || __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_7b29092764c6c9cb] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				__obf_7b29092764c6c9cb--
			}
		} else if __obf_f16b4157911bc9af == '\\' {
			__obf_23aafe4d90b76633 = true
		}
	}
	__obf_7b29092764c6c9cb := __obf_1bb30e8a74ed8233.__obf_15d837671d2809ae - 1
	for {
		if __obf_7b29092764c6c9cb < __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 || __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_7b29092764c6c9cb] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		__obf_7b29092764c6c9cb--
		if __obf_7b29092764c6c9cb < __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 || __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_7b29092764c6c9cb] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		__obf_7b29092764c6c9cb--

	}
	return -1, true // end with \
}
