package __obf_c3cd04a15c56f32f

import "fmt"

// ReadNil reads a json object as nil and
// returns whether it's a nil or not
func (__obf_edd9fe48d39445e4 *Iterator) ReadNil() (__obf_316af79472661247 bool) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == 'n' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l') // null
		return true
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	return false
}

// ReadBool reads a json object as BoolValue
func (__obf_edd9fe48d39445e4 *Iterator) ReadBool() (__obf_316af79472661247 bool) {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == 't' {
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('r', 'u', 'e')
		return true
	}
	if __obf_0c1bc1e511a43120 == 'f' {
		__obf_edd9fe48d39445e4.__obf_f22f308da0537336('a', 'l', 's', 'e')
		return false
	}
	__obf_edd9fe48d39445e4.
		ReportError("ReadBool", "expect t or f, but found "+string([]byte{__obf_0c1bc1e511a43120}))
	return
}

// SkipAndReturnBytes skip next JSON element, and return its content as []byte.
// The []byte can be kept, it is a copy of data.
func (__obf_edd9fe48d39445e4 *Iterator) SkipAndReturnBytes() []byte {
	__obf_edd9fe48d39445e4.__obf_e786284811b1ebdd(__obf_edd9fe48d39445e4.__obf_edd3c3885447229b)
	__obf_edd9fe48d39445e4.
		Skip()
	return __obf_edd9fe48d39445e4.__obf_e27ad1d64f9218e7()
}

// SkipAndAppendBytes skips next JSON element and appends its content to
// buffer, returning the result.
func (__obf_edd9fe48d39445e4 *Iterator) SkipAndAppendBytes(__obf_ace979f6622823f3 []byte) []byte {
	__obf_edd9fe48d39445e4.__obf_ac1ffebf099a785a(__obf_ace979f6622823f3, __obf_edd9fe48d39445e4.__obf_edd3c3885447229b)
	__obf_edd9fe48d39445e4.
		Skip()
	return __obf_edd9fe48d39445e4.__obf_e27ad1d64f9218e7()
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_ac1ffebf099a785a(__obf_ace979f6622823f3 []byte, __obf_df16664751f477bb int) {
	if __obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3 != nil {
		panic("already in capture mode")
	}
	__obf_edd9fe48d39445e4.__obf_df16664751f477bb = __obf_df16664751f477bb
	__obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3 = __obf_ace979f6622823f3
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_e786284811b1ebdd(__obf_df16664751f477bb int) {
	__obf_edd9fe48d39445e4.__obf_ac1ffebf099a785a(make([]byte, 0, 32), __obf_df16664751f477bb)
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_e27ad1d64f9218e7() []byte {
	if __obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3 == nil {
		panic("not in capture mode")
	}
	__obf_4db99b2aa14dd4a3 := __obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3
	__obf_17fc21b94a0188cf := __obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_edd9fe48d39445e4.__obf_df16664751f477bb:__obf_edd9fe48d39445e4.__obf_edd3c3885447229b]
	__obf_edd9fe48d39445e4.__obf_df16664751f477bb = -1
	__obf_edd9fe48d39445e4.__obf_4db99b2aa14dd4a3 = nil
	return append(__obf_4db99b2aa14dd4a3, __obf_17fc21b94a0188cf...)
}

// Skip skips a json object and positions to relatively the next json object
func (__obf_edd9fe48d39445e4 *Iterator) Skip() {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	switch __obf_0c1bc1e511a43120 {
	case '"':
		__obf_edd9fe48d39445e4.__obf_d743935ecad873a7()
	case 'n':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l') // null
	case 't':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('r', 'u', 'e') // true
	case 'f':
		__obf_edd9fe48d39445e4.__obf_f22f308da0537336('a', 'l', 's', 'e') // false
	case '0':
		__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
		__obf_edd9fe48d39445e4.
			ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		__obf_edd9fe48d39445e4.__obf_fdeaf1c2cbe3aa15()
	case '[':
		__obf_edd9fe48d39445e4.__obf_5566b22788a0a705()
	case '{':
		__obf_edd9fe48d39445e4.__obf_0704e073eacf05e6()
	default:
		__obf_edd9fe48d39445e4.
			ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", __obf_0c1bc1e511a43120))
		return
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_f22f308da0537336(__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe, __obf_6876881542173ffb byte) {
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_e7059202e5ecff53 {
		__obf_edd9fe48d39445e4.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe, __obf_6876881542173ffb})))
		return
	}
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_c6eebdcc01b11420 {
		__obf_edd9fe48d39445e4.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe, __obf_6876881542173ffb})))
		return
	}
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_8b95242c5191c1fe {
		__obf_edd9fe48d39445e4.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe, __obf_6876881542173ffb})))
		return
	}
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_6876881542173ffb {
		__obf_edd9fe48d39445e4.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe, __obf_6876881542173ffb})))
		return
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_8bc1f4b8d62f5247(__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe byte) {
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_e7059202e5ecff53 {
		__obf_edd9fe48d39445e4.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe})))
		return
	}
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_c6eebdcc01b11420 {
		__obf_edd9fe48d39445e4.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe})))
		return
	}
	if __obf_edd9fe48d39445e4.__obf_70c673c91de4f883() != __obf_8b95242c5191c1fe {
		__obf_edd9fe48d39445e4.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_e7059202e5ecff53, __obf_c6eebdcc01b11420, __obf_8b95242c5191c1fe})))
		return
	}
}
