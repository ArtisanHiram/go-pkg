package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"reflect"
	"unsafe"
)

func __obf_fca27e6ceeb592ed(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	type __obf_bd8a787985964aa7 struct {
		__obf_e94d87391e53ee63 *Binding
		__obf_5205590656760d45 string
		__obf_51580f7149e3d568 bool
	}
	__obf_8859db6368c47750 := []*__obf_bd8a787985964aa7{}
	__obf_832a553f1a21f913 := __obf_0446b148b9ab4187(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	for _, __obf_e94d87391e53ee63 := range __obf_832a553f1a21f913.Fields {
		for _, __obf_5205590656760d45 := range __obf_e94d87391e53ee63.ToNames {
			new := &__obf_bd8a787985964aa7{__obf_e94d87391e53ee63: __obf_e94d87391e53ee63, __obf_5205590656760d45: __obf_5205590656760d45}
			for _, __obf_b2f3b4aadfe3e860 := range __obf_8859db6368c47750 {
				if __obf_b2f3b4aadfe3e860.__obf_5205590656760d45 != __obf_5205590656760d45 {
					continue
				}
				__obf_b2f3b4aadfe3e860.__obf_51580f7149e3d568, new.__obf_51580f7149e3d568 = __obf_f52233c898eec9eb(__obf_62435d89ebefd5aa.__obf_3a25fb4c9a8342bb, __obf_b2f3b4aadfe3e860.__obf_e94d87391e53ee63, new.__obf_e94d87391e53ee63)
			}
			__obf_8859db6368c47750 = append(__obf_8859db6368c47750, new)
		}
	}
	if len(__obf_8859db6368c47750) == 0 {
		return &__obf_366a5a40d617c7e2{}
	}
	__obf_53b6376efd5ed746 := []__obf_a4fd68afc7e81fb7{}
	for _, __obf_bd8a787985964aa7 := range __obf_8859db6368c47750 {
		if !__obf_bd8a787985964aa7.__obf_51580f7149e3d568 {
			__obf_53b6376efd5ed746 = append(__obf_53b6376efd5ed746, __obf_a4fd68afc7e81fb7{__obf_c85b895560db528f: __obf_bd8a787985964aa7.__obf_e94d87391e53ee63.Encoder.(*__obf_8541dcc7318e179a), __obf_5205590656760d45: __obf_bd8a787985964aa7.__obf_5205590656760d45})
		}
	}
	return &__obf_44072a19950e97b8{__obf_8667d4fc2a95b0d7, __obf_53b6376efd5ed746}
}

func __obf_6ef591ca3f1a443c(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) __obf_7cb1bad9ea4a8234 {
	__obf_c85b895560db528f := __obf_3828217e60017848(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	__obf_0e6540508f0e8fa6 := __obf_8667d4fc2a95b0d7.Kind()
	switch __obf_0e6540508f0e8fa6 {
	case reflect.Interface:
		return &__obf_4fb4b91e56916c0b{__obf_8667d4fc2a95b0d7}
	case reflect.Struct:
		return &__obf_44072a19950e97b8{__obf_8667d4fc2a95b0d7: __obf_8667d4fc2a95b0d7}
	case reflect.Array:
		return &__obf_06afadf28a5822f4{}
	case reflect.Slice:
		return &__obf_9a75119645c258dc{}
	case reflect.Map:
		return __obf_f36df00362e2ff3a(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	case reflect.Ptr:
		return &OptionalEncoder{}
	default:
		return &__obf_a580dbbcc400f543{__obf_5966ecc5edb9b17e: fmt.Errorf("unsupported type: %v", __obf_8667d4fc2a95b0d7)}
	}
}

func __obf_f52233c898eec9eb(__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb, __obf_b2f3b4aadfe3e860, new *Binding) (__obf_265bf31f944f5bd2, __obf_60cf2537b91343db bool) {
	__obf_72a04513bf516f59 := new.Field.Tag().Get(__obf_0c8065c219ccf0ab.__obf_6cf02e65d7ac8e99()) != ""
	__obf_96d40fb180e7bf33 := __obf_b2f3b4aadfe3e860.Field.Tag().Get(__obf_0c8065c219ccf0ab.__obf_6cf02e65d7ac8e99()) != ""
	if __obf_72a04513bf516f59 {
		if __obf_96d40fb180e7bf33 {
			if len(__obf_b2f3b4aadfe3e860.__obf_74792d5daf7f8edc) > len(new.__obf_74792d5daf7f8edc) {
				return true, false
			} else if len(new.__obf_74792d5daf7f8edc) > len(__obf_b2f3b4aadfe3e860.__obf_74792d5daf7f8edc) {
				return false, true
			} else {
				return true, true
			}
		} else {
			return true, false
		}
	} else {
		if __obf_96d40fb180e7bf33 {
			return true, false
		}
		if len(__obf_b2f3b4aadfe3e860.__obf_74792d5daf7f8edc) > len(new.__obf_74792d5daf7f8edc) {
			return true, false
		} else if len(new.__obf_74792d5daf7f8edc) > len(__obf_b2f3b4aadfe3e860.__obf_74792d5daf7f8edc) {
			return false, true
		} else {
			return true, true
		}
	}
}

type __obf_8541dcc7318e179a struct {
	__obf_f992c91036745ca0 reflect2.StructField
	__obf_c7300fe99d67d416 ValEncoder
	__obf_87e72ba5f6e231d4 bool
}

func (__obf_c85b895560db528f *__obf_8541dcc7318e179a) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_f4916a65ae1a5a5a := __obf_c85b895560db528f.__obf_f992c91036745ca0.UnsafeGet(__obf_753ab3fb72cdd659)
	__obf_c85b895560db528f.__obf_c7300fe99d67d416.
		Encode(__obf_f4916a65ae1a5a5a, __obf_2361f5214e84c60f)
	if __obf_2361f5214e84c60f.Error != nil && __obf_2361f5214e84c60f.Error != io.EOF {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("%s: %s", __obf_c85b895560db528f.__obf_f992c91036745ca0.Name(), __obf_2361f5214e84c60f.Error.Error())
	}
}

func (__obf_c85b895560db528f *__obf_8541dcc7318e179a) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_f4916a65ae1a5a5a := __obf_c85b895560db528f.__obf_f992c91036745ca0.UnsafeGet(__obf_753ab3fb72cdd659)
	return __obf_c85b895560db528f.__obf_c7300fe99d67d416.IsEmpty(__obf_f4916a65ae1a5a5a)
}

func (__obf_c85b895560db528f *__obf_8541dcc7318e179a) IsEmbeddedPtrNil(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_defaafd546a4fefb, __obf_315a503b16d9c8a0 := __obf_c85b895560db528f.__obf_c7300fe99d67d416.(IsEmbeddedPtrNil)
	if !__obf_315a503b16d9c8a0 {
		return false
	}
	__obf_f4916a65ae1a5a5a := __obf_c85b895560db528f.__obf_f992c91036745ca0.UnsafeGet(__obf_753ab3fb72cdd659)
	return __obf_defaafd546a4fefb.IsEmbeddedPtrNil(__obf_f4916a65ae1a5a5a)
}

type IsEmbeddedPtrNil interface {
	IsEmbeddedPtrNil(__obf_753ab3fb72cdd659 unsafe.Pointer) bool
}

type __obf_44072a19950e97b8 struct {
	__obf_8667d4fc2a95b0d7 reflect2.Type
	__obf_ec908eb5c5599686 []__obf_a4fd68afc7e81fb7
}

type __obf_a4fd68afc7e81fb7 struct {
	__obf_c85b895560db528f *__obf_8541dcc7318e179a
	__obf_5205590656760d45 string
}

func (__obf_c85b895560db528f *__obf_44072a19950e97b8) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteObjectStart()
	__obf_4d0ffd1f01a4fd01 := false
	for _, __obf_f992c91036745ca0 := range __obf_c85b895560db528f.__obf_ec908eb5c5599686 {
		if __obf_f992c91036745ca0.__obf_c85b895560db528f.__obf_87e72ba5f6e231d4 && __obf_f992c91036745ca0.__obf_c85b895560db528f.IsEmpty(__obf_753ab3fb72cdd659) {
			continue
		}
		if __obf_f992c91036745ca0.__obf_c85b895560db528f.IsEmbeddedPtrNil(__obf_753ab3fb72cdd659) {
			continue
		}
		if __obf_4d0ffd1f01a4fd01 {
			__obf_2361f5214e84c60f.
				WriteMore()
		}
		__obf_2361f5214e84c60f.
			WriteObjectField(__obf_f992c91036745ca0.__obf_5205590656760d45)
		__obf_f992c91036745ca0.__obf_c85b895560db528f.
			Encode(__obf_753ab3fb72cdd659, __obf_2361f5214e84c60f)
		__obf_4d0ffd1f01a4fd01 = true
	}
	__obf_2361f5214e84c60f.
		WriteObjectEnd()
	if __obf_2361f5214e84c60f.Error != nil && __obf_2361f5214e84c60f.Error != io.EOF {
		__obf_2361f5214e84c60f.
			Error = fmt.Errorf("%v.%s", __obf_c85b895560db528f.__obf_8667d4fc2a95b0d7, __obf_2361f5214e84c60f.Error.Error())
	}
}

func (__obf_c85b895560db528f *__obf_44072a19950e97b8) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return false
}

type __obf_366a5a40d617c7e2 struct {
}

func (__obf_c85b895560db528f *__obf_366a5a40d617c7e2) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteEmptyObject()
}

func (__obf_c85b895560db528f *__obf_366a5a40d617c7e2) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return false
}

type __obf_2e0d36011b4e847e struct {
	__obf_2b4ba80da80f8306 ValEncoder
}

func (__obf_c85b895560db528f *__obf_2e0d36011b4e847e) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
	__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
		Encode(__obf_753ab3fb72cdd659, __obf_2361f5214e84c60f)
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
}

func (__obf_c85b895560db528f *__obf_2e0d36011b4e847e) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_2b4ba80da80f8306.IsEmpty(__obf_753ab3fb72cdd659)
}

type __obf_45b40254c768957c struct {
	__obf_2b4ba80da80f8306 ValEncoder
	__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb
}

func (__obf_c85b895560db528f *__obf_45b40254c768957c) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_8703f1e8bcaff394 := __obf_c85b895560db528f.__obf_0c8065c219ccf0ab.BorrowStream(nil)
	__obf_8703f1e8bcaff394.
		Attachment = __obf_2361f5214e84c60f.Attachment
	defer __obf_c85b895560db528f.__obf_0c8065c219ccf0ab.ReturnStream(__obf_8703f1e8bcaff394)
	__obf_c85b895560db528f.__obf_2b4ba80da80f8306.
		Encode(__obf_753ab3fb72cdd659, __obf_8703f1e8bcaff394)
	__obf_2361f5214e84c60f.
		WriteString(string(__obf_8703f1e8bcaff394.Buffer()))
}

func (__obf_c85b895560db528f *__obf_45b40254c768957c) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_2b4ba80da80f8306.IsEmpty(__obf_753ab3fb72cdd659)
}
