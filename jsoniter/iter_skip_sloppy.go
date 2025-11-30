//go:build jsoniter_sloppy
// +build jsoniter_sloppy

package __obf_c3cd04a15c56f32f

// sloppy but faster implementation, do not validate the input json

func (__obf_edd9fe48d39445e4 *Iterator) __obf_fdeaf1c2cbe3aa15() {
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
			switch __obf_0c1bc1e511a43120 {
			case ' ', '\n', '\r', '\t', ',', '}', ']':
				__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8
				return
			}
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			return
		}
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_5566b22788a0a705() {
	__obf_b9680d33d102638b := 1
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}
	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			switch __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8] {
			case '"':
				__obf_edd9fe48d39445e4. // If inside string, skip it
				__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
				__obf_edd9fe48d39445e4.__obf_d743935ecad873a7()
				__obf_28d099df85f083a8 = __obf_edd9fe48d39445e4. // it will be i++ soon
				__obf_edd3c3885447229b - 1
			case '[':
				__obf_b9680d33d102638b++// If open symbol, increase level

				if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
					return
				}
			case ']':
				__obf_b9680d33d102638b--// If close symbol, increase level

				if !__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_b9680d33d102638b == 0 {
					__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
					return
				}
			}
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			__obf_edd9fe48d39445e4.
				ReportError("skipObject", "incomplete array")
			return
		}
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_0704e073eacf05e6() {
	__obf_b9680d33d102638b := 1
	if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
		return
	}

	for {
		for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
			switch __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8] {
			case '"':
				__obf_edd9fe48d39445e4. // If inside string, skip it
				__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
				__obf_edd9fe48d39445e4.__obf_d743935ecad873a7()
				__obf_28d099df85f083a8 = __obf_edd9fe48d39445e4. // it will be i++ soon
				__obf_edd3c3885447229b - 1
			case '{':
				__obf_b9680d33d102638b++// If open symbol, increase level

				if !__obf_edd9fe48d39445e4.__obf_68afb6508328afc9() {
					return
				}
			case '}':
				__obf_b9680d33d102638b--// If close symbol, increase level

				if !__obf_edd9fe48d39445e4.__obf_0727968e0120c8a6() {
					return
				}

				// If we have returned to the original level, we're done
				if __obf_b9680d33d102638b == 0 {
					__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_28d099df85f083a8 + 1
					return
				}
			}
		}
		if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			__obf_edd9fe48d39445e4.
				ReportError("skipObject", "incomplete object")
			return
		}
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_d743935ecad873a7() {
	for {
		__obf_0b4b03a63da33b39, __obf_f83cc431d285a7be := __obf_edd9fe48d39445e4.__obf_c5e7eac0613854e1()
		if __obf_0b4b03a63da33b39 == -1 {
			if !__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
				__obf_edd9fe48d39445e4.
					ReportError("skipString", "incomplete string")
				return
			}
			if __obf_f83cc431d285a7be {
				__obf_edd9fe48d39445e4.
				// skip the first char as last char read is \
				__obf_edd3c3885447229b = 1
			}
		} else {
			__obf_edd9fe48d39445e4.__obf_edd3c3885447229b = __obf_0b4b03a63da33b39
			return
		}
	}
}

// adapted from: https://github.com/buger/jsonparser/blob/master/parser.go
// Tries to find the end of string
// Support if string contains escaped quote symbols.
func (__obf_edd9fe48d39445e4 *Iterator) __obf_c5e7eac0613854e1() (int, bool) {
	__obf_f83cc431d285a7be := false
	for __obf_28d099df85f083a8 := __obf_edd9fe48d39445e4.__obf_edd3c3885447229b; __obf_28d099df85f083a8 < __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57; __obf_28d099df85f083a8++ {
		__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_28d099df85f083a8]
		if __obf_0c1bc1e511a43120 == '"' {
			if !__obf_f83cc431d285a7be {
				return __obf_28d099df85f083a8 + 1, false
			}
			__obf_57cd9c621e7db075 := __obf_28d099df85f083a8 - 1
			for {
				if __obf_57cd9c621e7db075 < __obf_edd9fe48d39445e4.__obf_edd3c3885447229b || __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_57cd9c621e7db075] != '\\' {
					// even number of backslashes
					// either end of buffer, or " found
					return __obf_28d099df85f083a8 + 1, true
				}
				__obf_57cd9c621e7db075--
				if __obf_57cd9c621e7db075 < __obf_edd9fe48d39445e4.__obf_edd3c3885447229b || __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_57cd9c621e7db075] != '\\' {
					// odd number of backslashes
					// it is \" or \\\"
					break
				}
				__obf_57cd9c621e7db075--
			}
		} else if __obf_0c1bc1e511a43120 == '\\' {
			__obf_f83cc431d285a7be = true
		}
	}
	__obf_57cd9c621e7db075 := __obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 - 1
	for {
		if __obf_57cd9c621e7db075 < __obf_edd9fe48d39445e4.__obf_edd3c3885447229b || __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_57cd9c621e7db075] != '\\' {
			// even number of backslashes
			// either end of buffer, or " found
			return -1, false // do not end with \
		}
		__obf_57cd9c621e7db075--
		if __obf_57cd9c621e7db075 < __obf_edd9fe48d39445e4.__obf_edd3c3885447229b || __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_57cd9c621e7db075] != '\\' {
			// odd number of backslashes
			// it is \" or \\\"
			break
		}
		__obf_57cd9c621e7db075--

	}
	return -1, true // end with \
}
