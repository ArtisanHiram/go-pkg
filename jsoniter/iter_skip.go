package __obf_703d23ba520c3295

import "fmt"

// ReadNil reads a json object as nil and
// returns whether it's a nil or not
func (__obf_47edab4c16a2d88d *Iterator) ReadNil() (__obf_551cbec38242ce0c bool) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l') // null
		return true
	}
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	return false
}

// ReadBool reads a json object as BoolValue
func (__obf_47edab4c16a2d88d *Iterator) ReadBool() (__obf_551cbec38242ce0c bool) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == 't' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('r', 'u', 'e')
		return true
	}
	if __obf_bd08f5161fff294a == 'f' {
		__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('a', 'l', 's', 'e')
		return false
	}
	__obf_47edab4c16a2d88d.
		ReportError("ReadBool", "expect t or f, but found "+string([]byte{__obf_bd08f5161fff294a}))
	return
}

// SkipAndReturnBytes skip next JSON element, and return its content as []byte.
// The []byte can be kept, it is a copy of data.
func (__obf_47edab4c16a2d88d *Iterator) SkipAndReturnBytes() []byte {
	__obf_47edab4c16a2d88d.__obf_57875b3e9496553e(__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11)
	__obf_47edab4c16a2d88d.
		Skip()
	return __obf_47edab4c16a2d88d.__obf_842667b1c2db6627()
}

// SkipAndAppendBytes skips next JSON element and appends its content to
// buffer, returning the result.
func (__obf_47edab4c16a2d88d *Iterator) SkipAndAppendBytes(__obf_a065f8e0da5f5952 []byte) []byte {
	__obf_47edab4c16a2d88d.__obf_4faf0450c6bfd59f(__obf_a065f8e0da5f5952, __obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11)
	__obf_47edab4c16a2d88d.
		Skip()
	return __obf_47edab4c16a2d88d.__obf_842667b1c2db6627()
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_4faf0450c6bfd59f(__obf_a065f8e0da5f5952 []byte, __obf_2e0a81bee9ef3e76 int) {
	if __obf_47edab4c16a2d88d.__obf_4b160b169035b16e != nil {
		panic("already in capture mode")
	}
	__obf_47edab4c16a2d88d.__obf_2e0a81bee9ef3e76 = __obf_2e0a81bee9ef3e76
	__obf_47edab4c16a2d88d.__obf_4b160b169035b16e = __obf_a065f8e0da5f5952
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_57875b3e9496553e(__obf_2e0a81bee9ef3e76 int) {
	__obf_47edab4c16a2d88d.__obf_4faf0450c6bfd59f(make([]byte, 0, 32), __obf_2e0a81bee9ef3e76)
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_842667b1c2db6627() []byte {
	if __obf_47edab4c16a2d88d.__obf_4b160b169035b16e == nil {
		panic("not in capture mode")
	}
	__obf_4b160b169035b16e := __obf_47edab4c16a2d88d.__obf_4b160b169035b16e
	__obf_87dfc2abf099a36b := __obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_47edab4c16a2d88d.__obf_2e0a81bee9ef3e76:__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11]
	__obf_47edab4c16a2d88d.__obf_2e0a81bee9ef3e76 = -1
	__obf_47edab4c16a2d88d.__obf_4b160b169035b16e = nil
	return append(__obf_4b160b169035b16e, __obf_87dfc2abf099a36b...)
}

// Skip skips a json object and positions to relatively the next json object
func (__obf_47edab4c16a2d88d *Iterator) Skip() {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	switch __obf_bd08f5161fff294a {
	case '"':
		__obf_47edab4c16a2d88d.__obf_3c5a32ecdaa844fb()
	case 'n':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l') // null
	case 't':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('r', 'u', 'e') // true
	case 'f':
		__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('a', 'l', 's', 'e') // false
	case '0':
		__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
		__obf_47edab4c16a2d88d.
			ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		__obf_47edab4c16a2d88d.__obf_91b8504ad35d5b75()
	case '[':
		__obf_47edab4c16a2d88d.__obf_93a7c6915e88f56c()
	case '{':
		__obf_47edab4c16a2d88d.__obf_e53a692fb1024439()
	default:
		__obf_47edab4c16a2d88d.
			ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", __obf_bd08f5161fff294a))
		return
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_6f584222681dcca0(__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25, __obf_a8c4d645c69a3e1f byte) {
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_a31799cbcd265b2a {
		__obf_47edab4c16a2d88d.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25, __obf_a8c4d645c69a3e1f})))
		return
	}
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_fad0123fcf505d59 {
		__obf_47edab4c16a2d88d.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25, __obf_a8c4d645c69a3e1f})))
		return
	}
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_d78380e3a9ddce25 {
		__obf_47edab4c16a2d88d.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25, __obf_a8c4d645c69a3e1f})))
		return
	}
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_a8c4d645c69a3e1f {
		__obf_47edab4c16a2d88d.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25, __obf_a8c4d645c69a3e1f})))
		return
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_0704d4be12be5e96(__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25 byte) {
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_a31799cbcd265b2a {
		__obf_47edab4c16a2d88d.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25})))
		return
	}
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_fad0123fcf505d59 {
		__obf_47edab4c16a2d88d.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25})))
		return
	}
	if __obf_47edab4c16a2d88d.__obf_2fa919905fa99cc3() != __obf_d78380e3a9ddce25 {
		__obf_47edab4c16a2d88d.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_a31799cbcd265b2a, __obf_fad0123fcf505d59, __obf_d78380e3a9ddce25})))
		return
	}
}
