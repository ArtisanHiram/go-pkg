//go:build jsoniter_sloppy
// +build jsoniter_sloppy

package __obf_703d23ba520c3295

// sloppy but faster implementation, do not validate the input json

func (__obf_47edab4c16a2d88d *Iterator) __obf_91b8504ad35d5b75() {
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
			switch __obf_bd08f5161fff294a {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d
				return
			}
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			return
		}
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_93a7c6915e88f56c() {
	__obf_c82dbf13dc44e3f2 := 1
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}
	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			switch __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d] {
			case '"':
				__obf_47edab4c16a2d88d. // If inside string, skip it
				__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
				__obf_47edab4c16a2d88d.__obf_3c5a32ecdaa844fb()
				__obf_b0a5d2bd48690f1d = __obf_47edab4c16a2d88d. // it will be i++ soon
				__obf_da6aa1cfbd772c11 - 1
			case '[':
				__obf_c82dbf13dc44e3f2++// If open symbol, increase level

				if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
					return
				}
			case ']':
				__obf_c82dbf13dc44e3f2--// If close symbol, increase level

				if !__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_c82dbf13dc44e3f2 == 0 {
					__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
					return
				}
			}
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			__obf_47edab4c16a2d88d.
				ReportError("skipObject", "incomplete array")
			return
		}
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_e53a692fb1024439() {
	__obf_c82dbf13dc44e3f2 := 1
	if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
		return
	}

	for {
		for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
			switch __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d] {
			case '"':
				__obf_47edab4c16a2d88d. // If inside string, skip it
				__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
				__obf_47edab4c16a2d88d.__obf_3c5a32ecdaa844fb()
				__obf_b0a5d2bd48690f1d = __obf_47edab4c16a2d88d. // it will be i++ soon
				__obf_da6aa1cfbd772c11 - 1
			case '{':
				__obf_c82dbf13dc44e3f2++// If open symbol, increase level

				if !__obf_47edab4c16a2d88d.__obf_0f044cac6de70712() {
					return
				}
			case '}':
				__obf_c82dbf13dc44e3f2--// If close symbol, increase level

				if !__obf_47edab4c16a2d88d.__obf_bf8dbc480744e1b2() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_c82dbf13dc44e3f2 == 0 {
					__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_b0a5d2bd48690f1d + 1
					return
				}
			}
		}
		if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			__obf_47edab4c16a2d88d.
				ReportError("skipObject", "incomplete object")
			return
		}
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_3c5a32ecdaa844fb() {
	for {
		__obf_605e27382c9ac6fb, __obf_56956d61483d3b5d := __obf_47edab4c16a2d88d.__obf_4fc29cabcd0e7a3f()
		if __obf_605e27382c9ac6fb == -1 {
			if !__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
				__obf_47edab4c16a2d88d.
					ReportError("skipString", "incomplete string")
				return
			}
			if __obf_56956d61483d3b5d {
				__obf_47edab4c16a2d88d.
				// skip the first char as last char read is \
				__obf_da6aa1cfbd772c11 = 1
			}
		} else {
			__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 = __obf_605e27382c9ac6fb
			return
		}
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (__obf_47edab4c16a2d88d *Iterator) __obf_4fc29cabcd0e7a3f() (int, bool) {
	__obf_56956d61483d3b5d := false
	for __obf_b0a5d2bd48690f1d := __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11; __obf_b0a5d2bd48690f1d < __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5; __obf_b0a5d2bd48690f1d++ {
		__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_b0a5d2bd48690f1d]
		if __obf_bd08f5161fff294a == '"' {
			if !__obf_56956d61483d3b5d {
				return __obf_b0a5d2bd48690f1d + 1, false
			}
			__obf_0e04c22ffdf339de := __obf_b0a5d2bd48690f1d - 1
			for {
				if __obf_0e04c22ffdf339de < __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 || __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_0e04c22ffdf339de] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return __obf_b0a5d2bd48690f1d + 1, true
				}
				__obf_0e04c22ffdf339de--
				if __obf_0e04c22ffdf339de < __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 || __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_0e04c22ffdf339de] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				__obf_0e04c22ffdf339de--
			}
		} else if __obf_bd08f5161fff294a == '\\' {
			__obf_56956d61483d3b5d = true
		}
	}
	__obf_0e04c22ffdf339de := __obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 - 1
	for {
		if __obf_0e04c22ffdf339de < __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 || __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_0e04c22ffdf339de] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		__obf_0e04c22ffdf339de--
		if __obf_0e04c22ffdf339de < __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 || __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_0e04c22ffdf339de] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		__obf_0e04c22ffdf339de--

	}
	return -1, true // end with \
}
