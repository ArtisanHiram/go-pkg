package __obf_5b802ce8d9ba56d6

import "fmt"

// ReadNil reads a json object as nil and
// returns whether it's a nil or not
func (__obf_67008a6a9e5ba828 *Iterator) ReadNil() (__obf_5dabcdfee5097ed6 bool) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == 'n' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l') // null
		return true
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	return false
}

// ReadBool reads a json object as BoolValue
func (__obf_67008a6a9e5ba828 *Iterator) ReadBool() (__obf_5dabcdfee5097ed6 bool) {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == 't' {
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('r', 'u', 'e')
		return true
	}
	if __obf_dab9baaadfa7c8c2 == 'f' {
		__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('a', 'l', 's', 'e')
		return false
	}
	__obf_67008a6a9e5ba828.
		ReportError("ReadBool", "expect t or f, but found "+string([]byte{__obf_dab9baaadfa7c8c2}))
	return
}

// SkipAndReturnBytes skip next JSON element, and return its content as []byte.
// The []byte can be kept, it is a copy of data.
func (__obf_67008a6a9e5ba828 *Iterator) SkipAndReturnBytes() []byte {
	__obf_67008a6a9e5ba828.__obf_1b03992be8504c10(__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36)
	__obf_67008a6a9e5ba828.
		Skip()
	return __obf_67008a6a9e5ba828.__obf_3998584530a209bb()
}

// SkipAndAppendBytes skips next JSON element and appends its content to
// buffer, returning the result.
func (__obf_67008a6a9e5ba828 *Iterator) SkipAndAppendBytes(__obf_9fc06d9180f0daca []byte) []byte {
	__obf_67008a6a9e5ba828.__obf_599177a934256761(__obf_9fc06d9180f0daca, __obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36)
	__obf_67008a6a9e5ba828.
		Skip()
	return __obf_67008a6a9e5ba828.__obf_3998584530a209bb()
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_599177a934256761(__obf_9fc06d9180f0daca []byte, __obf_787a5df6d8c7d251 int) {
	if __obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b != nil {
		panic("already in capture mode")
	}
	__obf_67008a6a9e5ba828.__obf_787a5df6d8c7d251 = __obf_787a5df6d8c7d251
	__obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b = __obf_9fc06d9180f0daca
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_1b03992be8504c10(__obf_787a5df6d8c7d251 int) {
	__obf_67008a6a9e5ba828.__obf_599177a934256761(make([]byte, 0, 32), __obf_787a5df6d8c7d251)
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_3998584530a209bb() []byte {
	if __obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b == nil {
		panic("not in capture mode")
	}
	__obf_645fd9e4de8f601b := __obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b
	__obf_1a7ec2cfcbfcf807 := __obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_67008a6a9e5ba828.__obf_787a5df6d8c7d251:__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36]
	__obf_67008a6a9e5ba828.__obf_787a5df6d8c7d251 = -1
	__obf_67008a6a9e5ba828.__obf_645fd9e4de8f601b = nil
	return append(__obf_645fd9e4de8f601b, __obf_1a7ec2cfcbfcf807...)
}

// Skip skips a json object and positions to relatively the next json object
func (__obf_67008a6a9e5ba828 *Iterator) Skip() {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	switch __obf_dab9baaadfa7c8c2 {
	case '"':
		__obf_67008a6a9e5ba828.__obf_acb079d30d4527e3()
	case 'n':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l') // null
	case 't':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('r', 'u', 'e') // true
	case 'f':
		__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('a', 'l', 's', 'e') // false
	case '0':
		__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
		__obf_67008a6a9e5ba828.
			ReadFloat32()
	case '-', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		__obf_67008a6a9e5ba828.__obf_a64603a02c650185()
	case '[':
		__obf_67008a6a9e5ba828.__obf_1d506622058b77ef()
	case '{':
		__obf_67008a6a9e5ba828.__obf_7e3d8ea4ab2a8258()
	default:
		__obf_67008a6a9e5ba828.
			ReportError("Skip", fmt.Sprintf("do not know how to skip: %v", __obf_dab9baaadfa7c8c2))
		return
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_acc74c95f4492ff8(__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89, __obf_04e7e16162feaf64 byte) {
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_1ab8ed8983ef4e90 {
		__obf_67008a6a9e5ba828.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89, __obf_04e7e16162feaf64})))
		return
	}
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_b02ffc6f1fe55bba {
		__obf_67008a6a9e5ba828.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89, __obf_04e7e16162feaf64})))
		return
	}
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_4f4b35ec0ba0ae89 {
		__obf_67008a6a9e5ba828.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89, __obf_04e7e16162feaf64})))
		return
	}
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_04e7e16162feaf64 {
		__obf_67008a6a9e5ba828.
			ReportError("skipFourBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89, __obf_04e7e16162feaf64})))
		return
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_4aeb767e0be7277a(__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89 byte) {
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_1ab8ed8983ef4e90 {
		__obf_67008a6a9e5ba828.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89})))
		return
	}
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_b02ffc6f1fe55bba {
		__obf_67008a6a9e5ba828.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89})))
		return
	}
	if __obf_67008a6a9e5ba828.__obf_ea3ebd5c6789bccb() != __obf_4f4b35ec0ba0ae89 {
		__obf_67008a6a9e5ba828.
			ReportError("skipThreeBytes", fmt.Sprintf("expect %s", string([]byte{__obf_1ab8ed8983ef4e90, __obf_b02ffc6f1fe55bba, __obf_4f4b35ec0ba0ae89})))
		return
	}
}
