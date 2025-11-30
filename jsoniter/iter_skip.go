package __obf_030d39f7a8de96c6

import "fmt"

// ReadNil reads a json object as nil and
// returns whether it's a nil or not
func (__obf_4ab56a99965952e7 *Iterator) ReadNil() (__obf_59c72400ec1a6110 bool) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == 'n' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l') // null
		return true
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	return false
}

// ReadBool reads a json object as BoolValue
func (__obf_4ab56a99965952e7 *Iterator) ReadBool() (__obf_59c72400ec1a6110 bool) {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == 't' {
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('r', 'u', 'e')
		return true
	}
	if __obf_24309b3b0ff9ba22 == 'f' {
		__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('a', 'l', 's', 'e')
		return false
	}
	__obf_4ab56a99965952e7.
		ReportError("ReadBool", "expect t or f, but found "+string([]byte{__obf_24309b3b0ff9ba22}))
	return
}

// SkipAndReturnBytes skip next JSON element, and return its content as []byte.
// The []byte can be kept, it is a copy of data.
func (__obf_4ab56a99965952e7 *Iterator) SkipAndReturnBytes() []byte {
	__obf_4ab56a99965952e7.__obf_f1e0f594888555ba(__obf_4ab56a99965952e7.__obf_5e22d6270d27491f)
	__obf_4ab56a99965952e7.
		Skip()
	return __obf_4ab56a99965952e7.__obf_4403a37119edc6b3()
}

// SkipAndAppendBytes skips next JSON element and appends its content to
// buffer, returning the result.
func (__obf_4ab56a99965952e7 *Iterator) SkipAndAppendBytes(__obf_0b1656d7ef4bd9df []byte) []byte {
	__obf_4ab56a99965952e7.__obf_8429981e182bd1e8(__obf_0b1656d7ef4bd9df, __obf_4ab56a99965952e7.__obf_5e22d6270d27491f)
	__obf_4ab56a99965952e7.
		Skip()
	return __obf_4ab56a99965952e7.__obf_4403a37119edc6b3()
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_8429981e182bd1e8(__obf_0b1656d7ef4bd9df []byte, __obf_85ab55dc89939a18 int) {
	if __obf_4ab56a99965952e7.__obf_4541c922f6acdb7d != nil {
		panic("already in capture mode")
	}
	__obf_4ab56a99965952e7.__obf_85ab55dc89939a18 = __obf_85ab55dc89939a18
	__obf_4ab56a99965952e7.__obf_4541c922f6acdb7d = __obf_0b1656d7ef4bd9df
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_f1e0f594888555ba(__obf_85ab55dc89939a18 int) {
	__obf_4ab56a99965952e7.__obf_8429981e182bd1e8(make([]byte, 0, 32), __obf_85ab55dc89939a18)
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_4403a37119edc6b3() []byte {
	if __obf_4ab56a99965952e7.__obf_4541c922f6acdb7d == nil {
		panic("not in capture mode")
	}
	__obf_4541c922f6acdb7d := __obf_4ab56a99965952e7.__obf_4541c922f6acdb7d
	__obf_2257b1725dfd65c5 := __obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_4ab56a99965952e7.__obf_85ab55dc89939a18:__obf_4ab56a99965952e7.__obf_5e22d6270d27491f]
	__obf_4ab56a99965952e7.__obf_85ab55dc89939a18 = -1
	__obf_4ab56a99965952e7.__obf_4541c922f6acdb7d = nil
	return append(__obf_4541c922f6acdb7d, __obf_2257b1725dfd65c5...)
}

// Skip skips a json object and positions to relatively the next json object
func (__obf_4ab56a99965952e7 *Iterator) Skip() {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	switch __obf_24309b3b0ff9ba22 {
	case '"':
		__obf_4ab56a99965952e7.__obf_ab668369e47b63de()
	case 'n':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l') // null
	case 't':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('r', 'u', 'e') // true
	case 'f':
		__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('a', 'l', 's', 'e') // false
	case '0':
		__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
		__obf_4ab56a99965952e7.
			ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		__obf_4ab56a99965952e7.__obf_38c33066e8f2b292()
	case '[':
		__obf_4ab56a99965952e7.__obf_e8e6ee2ef6c27f13()
	case '{':
		__obf_4ab56a99965952e7.__obf_56cafb6b2d847244()
	default:
		__obf_4ab56a99965952e7.
			ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", __obf_24309b3b0ff9ba22))
		return
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_aaf95589108acb4b(__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5, __obf_12099b3177c524c2 byte) {
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_9a19699afdc51406 {
		__obf_4ab56a99965952e7.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5, __obf_12099b3177c524c2})))
		return
	}
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_121f32b9e51592b3 {
		__obf_4ab56a99965952e7.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5, __obf_12099b3177c524c2})))
		return
	}
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_17bdad7c3468c7b5 {
		__obf_4ab56a99965952e7.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5, __obf_12099b3177c524c2})))
		return
	}
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_12099b3177c524c2 {
		__obf_4ab56a99965952e7.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5, __obf_12099b3177c524c2})))
		return
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_01038f7a5ba9ca56(__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5 byte) {
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_9a19699afdc51406 {
		__obf_4ab56a99965952e7.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5})))
		return
	}
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_121f32b9e51592b3 {
		__obf_4ab56a99965952e7.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5})))
		return
	}
	if __obf_4ab56a99965952e7.__obf_2b6cf70e8dd32b68() != __obf_17bdad7c3468c7b5 {
		__obf_4ab56a99965952e7.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_9a19699afdc51406, __obf_121f32b9e51592b3, __obf_17bdad7c3468c7b5})))
		return
	}
}
