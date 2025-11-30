package __obf_c3cd04a15c56f32f

import (
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_3839c59f7b1fc415 = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var __obf_a99d3c45c00f40a9 = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func __obf_5939b4177daf663d(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	if __obf_8667d4fc2a95b0d7 == __obf_3839c59f7b1fc415 {
		return &__obf_97c7eba96d9d09af{}
	}
	if __obf_8667d4fc2a95b0d7 == __obf_a99d3c45c00f40a9 {
		return &__obf_af6168e45506fa5d{}
	}
	return nil
}

func __obf_93af885996aef0fb(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	if __obf_8667d4fc2a95b0d7 == __obf_3839c59f7b1fc415 {
		return &__obf_97c7eba96d9d09af{}
	}
	if __obf_8667d4fc2a95b0d7 == __obf_a99d3c45c00f40a9 {
		return &__obf_af6168e45506fa5d{}
	}
	return nil
}

type __obf_97c7eba96d9d09af struct {
}

func (__obf_1569d098873d2228 *__obf_97c7eba96d9d09af) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.ReadNil() {
		*((*json.RawMessage)(__obf_753ab3fb72cdd659)) = nil
	} else {
		*((*json.RawMessage)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.SkipAndReturnBytes()
	}
}

func (__obf_1569d098873d2228 *__obf_97c7eba96d9d09af) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if *((*json.RawMessage)(__obf_753ab3fb72cdd659)) == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
	} else {
		__obf_2361f5214e84c60f.
			WriteRaw(string(*((*json.RawMessage)(__obf_753ab3fb72cdd659))))
	}
}

func (__obf_1569d098873d2228 *__obf_97c7eba96d9d09af) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(__obf_753ab3fb72cdd659))) == 0
}

type __obf_af6168e45506fa5d struct {
}

func (__obf_1569d098873d2228 *__obf_af6168e45506fa5d) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.ReadNil() {
		*((*RawMessage)(__obf_753ab3fb72cdd659)) = nil
	} else {
		*((*RawMessage)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.SkipAndReturnBytes()
	}
}

func (__obf_1569d098873d2228 *__obf_af6168e45506fa5d) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if *((*RawMessage)(__obf_753ab3fb72cdd659)) == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
	} else {
		__obf_2361f5214e84c60f.
			WriteRaw(string(*((*RawMessage)(__obf_753ab3fb72cdd659))))
	}
}

func (__obf_1569d098873d2228 *__obf_af6168e45506fa5d) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return len(*((*RawMessage)(__obf_753ab3fb72cdd659))) == 0
}
