package __obf_91620b895eeff9ed

import "fmt"

// ReadNil reads a json object as nil and
// returns whether it's a nil or not
func (__obf_1bb30e8a74ed8233 *Iterator) ReadNil() (__obf_e46f5fe3db5036fe bool) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == 'n' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l') // null
		return true
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	return false
}

// ReadBool reads a json object as BoolValue
func (__obf_1bb30e8a74ed8233 *Iterator) ReadBool() (__obf_e46f5fe3db5036fe bool) {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == 't' {
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('r', 'u', 'e')
		return true
	}
	if __obf_f16b4157911bc9af == 'f' {
		__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('a', 'l', 's', 'e')
		return false
	}
	__obf_1bb30e8a74ed8233.
		ReportError("ReadBool", "expect t or f, but found "+string([]byte{__obf_f16b4157911bc9af}))
	return
}

// SkipAndReturnBytes skip next JSON element, and return its content as []byte.
// The []byte can be kept, it is a copy of data.
func (__obf_1bb30e8a74ed8233 *Iterator) SkipAndReturnBytes() []byte {
	__obf_1bb30e8a74ed8233.__obf_c704d25028ca6789(__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21)
	__obf_1bb30e8a74ed8233.
		Skip()
	return __obf_1bb30e8a74ed8233.__obf_2141069796e7e215()
}

// SkipAndAppendBytes skips next JSON element and appends its content to
// buffer, returning the result.
func (__obf_1bb30e8a74ed8233 *Iterator) SkipAndAppendBytes(__obf_184433571fa55237 []byte) []byte {
	__obf_1bb30e8a74ed8233.__obf_7bc9f2cbb69fe974(__obf_184433571fa55237, __obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21)
	__obf_1bb30e8a74ed8233.
		Skip()
	return __obf_1bb30e8a74ed8233.__obf_2141069796e7e215()
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_7bc9f2cbb69fe974(__obf_184433571fa55237 []byte, __obf_f69c5dadb6d744d6 int) {
	if __obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e != nil {
		panic("already in capture mode")
	}
	__obf_1bb30e8a74ed8233.__obf_f69c5dadb6d744d6 = __obf_f69c5dadb6d744d6
	__obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e = __obf_184433571fa55237
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_c704d25028ca6789(__obf_f69c5dadb6d744d6 int) {
	__obf_1bb30e8a74ed8233.__obf_7bc9f2cbb69fe974(make([]byte, 0, 32), __obf_f69c5dadb6d744d6)
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_2141069796e7e215() []byte {
	if __obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e == nil {
		panic("not in capture mode")
	}
	__obf_7a95a5ed5bcae59e := __obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e
	__obf_661d5d31d1b381f5 := __obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_1bb30e8a74ed8233.__obf_f69c5dadb6d744d6:__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21]
	__obf_1bb30e8a74ed8233.__obf_f69c5dadb6d744d6 = -1
	__obf_1bb30e8a74ed8233.__obf_7a95a5ed5bcae59e = nil
	return append(__obf_7a95a5ed5bcae59e, __obf_661d5d31d1b381f5...)
}

// Skip skips a json object and positions to relatively the next json object
func (__obf_1bb30e8a74ed8233 *Iterator) Skip() {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	switch __obf_f16b4157911bc9af {
	case '"':
		__obf_1bb30e8a74ed8233.__obf_c5493c5b3916fc3f()
	case 'n':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l') // null
	case 't':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('r', 'u', 'e') // true
	case 'f':
		__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('a', 'l', 's', 'e') // false
	case '0':
		__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
		__obf_1bb30e8a74ed8233.
			ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		__obf_1bb30e8a74ed8233.__obf_d6a474bd319b770c()
	case '[':
		__obf_1bb30e8a74ed8233.__obf_531ac82b5de39aec()
	case '{':
		__obf_1bb30e8a74ed8233.__obf_8d12ebf64f0517d2()
	default:
		__obf_1bb30e8a74ed8233.
			ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", __obf_f16b4157911bc9af))
		return
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_94a8e39928c8440c(__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d, __obf_2d0ea2daef3f1239 byte) {
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_3476619ed7f7bbc5 {
		__obf_1bb30e8a74ed8233.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d, __obf_2d0ea2daef3f1239})))
		return
	}
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_afc38b1fa6a4101a {
		__obf_1bb30e8a74ed8233.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d, __obf_2d0ea2daef3f1239})))
		return
	}
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_c9b933a92c45005d {
		__obf_1bb30e8a74ed8233.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d, __obf_2d0ea2daef3f1239})))
		return
	}
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_2d0ea2daef3f1239 {
		__obf_1bb30e8a74ed8233.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d, __obf_2d0ea2daef3f1239})))
		return
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_3e1d2ad9a54f0d22(__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d byte) {
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_3476619ed7f7bbc5 {
		__obf_1bb30e8a74ed8233.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d})))
		return
	}
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_afc38b1fa6a4101a {
		__obf_1bb30e8a74ed8233.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d})))
		return
	}
	if __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc() != __obf_c9b933a92c45005d {
		__obf_1bb30e8a74ed8233.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_3476619ed7f7bbc5, __obf_afc38b1fa6a4101a, __obf_c9b933a92c45005d})))
		return
	}
}
