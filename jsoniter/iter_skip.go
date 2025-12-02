package __obf_c7b79b12b33d8238

import "fmt"

// ReadNil reads a json object as nil and
// returns whether it's a nil or not
func (__obf_0da8c843dad13139 *Iterator) ReadNil() (__obf_9a8638dac9e1c083 bool) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == 'n' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l') // null
		return true
	}
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	return false
}

// ReadBool reads a json object as BoolValue
func (__obf_0da8c843dad13139 *Iterator) ReadBool() (__obf_9a8638dac9e1c083 bool) {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == 't' {
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('r', 'u', 'e')
		return true
	}
	if __obf_12577c03a12f0d2c == 'f' {
		__obf_0da8c843dad13139.__obf_62908d9424a8486b('a', 'l', 's', 'e')
		return false
	}
	__obf_0da8c843dad13139.
		ReportError("ReadBool", "expect t or f, but found "+string([]byte{__obf_12577c03a12f0d2c}))
	return
}

// SkipAndReturnBytes skip next JSON element, and return its content as []byte.
// The []byte can be kept, it is a copy of data.
func (__obf_0da8c843dad13139 *Iterator) SkipAndReturnBytes() []byte {
	__obf_0da8c843dad13139.__obf_611b08915abc95f3(__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2)
	__obf_0da8c843dad13139.
		Skip()
	return __obf_0da8c843dad13139.__obf_d96514238c6d6bd1()
}

// SkipAndAppendBytes skips next JSON element and appends its content to
// buffer, returning the result.
func (__obf_0da8c843dad13139 *Iterator) SkipAndAppendBytes(__obf_684d738bc3ac851a []byte) []byte {
	__obf_0da8c843dad13139.__obf_aef8389bd08e834c(__obf_684d738bc3ac851a, __obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2)
	__obf_0da8c843dad13139.
		Skip()
	return __obf_0da8c843dad13139.__obf_d96514238c6d6bd1()
}

func (__obf_0da8c843dad13139 *Iterator) __obf_aef8389bd08e834c(__obf_684d738bc3ac851a []byte, __obf_df8f35c95ef68415 int) {
	if __obf_0da8c843dad13139.__obf_228f49de295a387e != nil {
		panic("already in capture mode")
	}
	__obf_0da8c843dad13139.__obf_df8f35c95ef68415 = __obf_df8f35c95ef68415
	__obf_0da8c843dad13139.__obf_228f49de295a387e = __obf_684d738bc3ac851a
}

func (__obf_0da8c843dad13139 *Iterator) __obf_611b08915abc95f3(__obf_df8f35c95ef68415 int) {
	__obf_0da8c843dad13139.__obf_aef8389bd08e834c(make([]byte, 0, 32), __obf_df8f35c95ef68415)
}

func (__obf_0da8c843dad13139 *Iterator) __obf_d96514238c6d6bd1() []byte {
	if __obf_0da8c843dad13139.__obf_228f49de295a387e == nil {
		panic("not in capture mode")
	}
	__obf_228f49de295a387e := __obf_0da8c843dad13139.__obf_228f49de295a387e
	__obf_2565c9afb6012ebe := __obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_0da8c843dad13139.__obf_df8f35c95ef68415:__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2]
	__obf_0da8c843dad13139.__obf_df8f35c95ef68415 = -1
	__obf_0da8c843dad13139.__obf_228f49de295a387e = nil
	return append(__obf_228f49de295a387e, __obf_2565c9afb6012ebe...)
}

// Skip skips a json object and positions to relatively the next json object
func (__obf_0da8c843dad13139 *Iterator) Skip() {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	switch __obf_12577c03a12f0d2c {
	case '"':
		__obf_0da8c843dad13139.__obf_7dda9c5406ae47f0()
	case 'n':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l') // null
	case 't':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('r', 'u', 'e') // true
	case 'f':
		__obf_0da8c843dad13139.__obf_62908d9424a8486b('a', 'l', 's', 'e') // false
	case '0':
		__obf_0da8c843dad13139.__obf_a675366c80290b83()
		__obf_0da8c843dad13139.
			ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		__obf_0da8c843dad13139.__obf_9dd6f9cece1dabe0()
	case '[':
		__obf_0da8c843dad13139.__obf_a12f980c2c849354()
	case '{':
		__obf_0da8c843dad13139.__obf_f42b686c47b3acb5()
	default:
		__obf_0da8c843dad13139.
			ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", __obf_12577c03a12f0d2c))
		return
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_62908d9424a8486b(__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf, __obf_74da2ff524f360c1 byte) {
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_ec0e29010fd9c6a7 {
		__obf_0da8c843dad13139.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf, __obf_74da2ff524f360c1})))
		return
	}
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_dda50892fe7798e2 {
		__obf_0da8c843dad13139.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf, __obf_74da2ff524f360c1})))
		return
	}
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_d8fc4d94c4f666bf {
		__obf_0da8c843dad13139.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf, __obf_74da2ff524f360c1})))
		return
	}
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_74da2ff524f360c1 {
		__obf_0da8c843dad13139.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf, __obf_74da2ff524f360c1})))
		return
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_5da8d54e7a1e782c(__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf byte) {
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_ec0e29010fd9c6a7 {
		__obf_0da8c843dad13139.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf})))
		return
	}
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_dda50892fe7798e2 {
		__obf_0da8c843dad13139.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf})))
		return
	}
	if __obf_0da8c843dad13139.__obf_fe749dd9dadf32f8() != __obf_d8fc4d94c4f666bf {
		__obf_0da8c843dad13139.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_ec0e29010fd9c6a7, __obf_dda50892fe7798e2, __obf_d8fc4d94c4f666bf})))
		return
	}
}
