//go:build jsoniter_sloppy
// +build jsoniter_sloppy

package __obf_c7b79b12b33d8238

// sloppy but faster implementation, do not validate the input json

func (__obf_0da8c843dad13139 *Iterator) __obf_9dd6f9cece1dabe0() {
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
			switch __obf_12577c03a12f0d2c {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb
				return
			}
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			return
		}
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_a12f980c2c849354() {
	__obf_9b059c758f2afc6b := 1
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}
	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			switch __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb] {
			case '"':
				__obf_0da8c843dad13139. // If inside string, skip it
				__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
				__obf_0da8c843dad13139.__obf_7dda9c5406ae47f0()
				__obf_a051269af3a541bb = __obf_0da8c843dad13139. // it will be i++ soon
				__obf_6a63c9ac34fe84e2 - 1
			case '[':
				__obf_9b059c758f2afc6b++// If open symbol, increase level

				if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
					return
				}
			case ']':
				__obf_9b059c758f2afc6b--// If close symbol, increase level

				if !__obf_0da8c843dad13139.__obf_aed83e308a23ff9c() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_9b059c758f2afc6b == 0 {
					__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
					return
				}
			}
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			__obf_0da8c843dad13139.
				ReportError("skipObject", "incomplete array")
			return
		}
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_f42b686c47b3acb5() {
	__obf_9b059c758f2afc6b := 1
	if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
		return
	}

	for {
		for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
			switch __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb] {
			case '"':
				__obf_0da8c843dad13139. // If inside string, skip it
				__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
				__obf_0da8c843dad13139.__obf_7dda9c5406ae47f0()
				__obf_a051269af3a541bb = __obf_0da8c843dad13139. // it will be i++ soon
				__obf_6a63c9ac34fe84e2 - 1
			case '{':
				__obf_9b059c758f2afc6b++// If open symbol, increase level

				if !__obf_0da8c843dad13139.__obf_b638f6394ae6a28c() {
					return
				}
			case '}':
				__obf_9b059c758f2afc6b--// If close symbol, increase level

				if !__obf_0da8c843dad13139.__obf_aed83e308a23ff9c() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_9b059c758f2afc6b == 0 {
					__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_a051269af3a541bb + 1
					return
				}
			}
		}
		if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			__obf_0da8c843dad13139.
				ReportError("skipObject", "incomplete object")
			return
		}
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_7dda9c5406ae47f0() {
	for {
		__obf_e5b0220fa04a0f83, __obf_4cccd34469a00d56 := __obf_0da8c843dad13139.__obf_856a142cec2e887c()
		if __obf_e5b0220fa04a0f83 == -1 {
			if !__obf_0da8c843dad13139.__obf_baaf768e13842431() {
				__obf_0da8c843dad13139.
					ReportError("skipString", "incomplete string")
				return
			}
			if __obf_4cccd34469a00d56 {
				__obf_0da8c843dad13139.
				// skip the first char as last char read is \
				__obf_6a63c9ac34fe84e2 = 1
			}
		} else {
			__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 = __obf_e5b0220fa04a0f83
			return
		}
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (__obf_0da8c843dad13139 *Iterator) __obf_856a142cec2e887c() (int, bool) {
	__obf_4cccd34469a00d56 := false
	for __obf_a051269af3a541bb := __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2; __obf_a051269af3a541bb < __obf_0da8c843dad13139.__obf_840246080559848c; __obf_a051269af3a541bb++ {
		__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_a051269af3a541bb]
		if __obf_12577c03a12f0d2c == '"' {
			if !__obf_4cccd34469a00d56 {
				return __obf_a051269af3a541bb + 1, false
			}
			__obf_aac66f5ab0eab626 := __obf_a051269af3a541bb - 1
			for {
				if __obf_aac66f5ab0eab626 < __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 || __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_aac66f5ab0eab626] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return __obf_a051269af3a541bb + 1, true
				}
				__obf_aac66f5ab0eab626--
				if __obf_aac66f5ab0eab626 < __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 || __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_aac66f5ab0eab626] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				__obf_aac66f5ab0eab626--
			}
		} else if __obf_12577c03a12f0d2c == '\\' {
			__obf_4cccd34469a00d56 = true
		}
	}
	__obf_aac66f5ab0eab626 := __obf_0da8c843dad13139.__obf_840246080559848c - 1
	for {
		if __obf_aac66f5ab0eab626 < __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 || __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_aac66f5ab0eab626] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		__obf_aac66f5ab0eab626--
		if __obf_aac66f5ab0eab626 < __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 || __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_aac66f5ab0eab626] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		__obf_aac66f5ab0eab626--

	}
	return -1, true // end with \
}
