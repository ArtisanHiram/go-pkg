package __obf_c7b79b12b33d8238

import (
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_1c2de5de21e87b65 = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var __obf_5c3046a12fdb76f5 = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func __obf_f49908b292adac72(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	if __obf_edcded11af6ebc4c == __obf_1c2de5de21e87b65 {
		return &__obf_0c7b7acd5d922c87{}
	}
	if __obf_edcded11af6ebc4c == __obf_5c3046a12fdb76f5 {
		return &__obf_a9584e35b79c3dfe{}
	}
	return nil
}

func __obf_4360c3949f7a6e97(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	if __obf_edcded11af6ebc4c == __obf_1c2de5de21e87b65 {
		return &__obf_0c7b7acd5d922c87{}
	}
	if __obf_edcded11af6ebc4c == __obf_5c3046a12fdb76f5 {
		return &__obf_a9584e35b79c3dfe{}
	}
	return nil
}

type __obf_0c7b7acd5d922c87 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_0c7b7acd5d922c87) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.ReadNil() {
		*((*json.RawMessage)(__obf_575c04f2b9d91315)) = nil
	} else {
		*((*json.RawMessage)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.SkipAndReturnBytes()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_0c7b7acd5d922c87) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if *((*json.RawMessage)(__obf_575c04f2b9d91315)) == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
	} else {
		__obf_d8f50bcfe87d8b47.
			WriteRaw(string(*((*json.RawMessage)(__obf_575c04f2b9d91315))))
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_0c7b7acd5d922c87) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(__obf_575c04f2b9d91315))) == 0
}

type __obf_a9584e35b79c3dfe struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_a9584e35b79c3dfe) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.ReadNil() {
		*((*RawMessage)(__obf_575c04f2b9d91315)) = nil
	} else {
		*((*RawMessage)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.SkipAndReturnBytes()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_a9584e35b79c3dfe) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if *((*RawMessage)(__obf_575c04f2b9d91315)) == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
	} else {
		__obf_d8f50bcfe87d8b47.
			WriteRaw(string(*((*RawMessage)(__obf_575c04f2b9d91315))))
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_a9584e35b79c3dfe) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return len(*((*RawMessage)(__obf_575c04f2b9d91315))) == 0
}
