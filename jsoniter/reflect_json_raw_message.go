package __obf_91620b895eeff9ed

import (
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_2fe32ebd686258b1 = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var __obf_5c597776869c8984 = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func __obf_70759df1070bb0f4(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	if __obf_29ebd0f2c324f5ea == __obf_2fe32ebd686258b1 {
		return &__obf_f73bdd1883284ee0{}
	}
	if __obf_29ebd0f2c324f5ea == __obf_5c597776869c8984 {
		return &__obf_bd17dda1a88cca5a{}
	}
	return nil
}

func __obf_f5b325d98e33bfe7(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	if __obf_29ebd0f2c324f5ea == __obf_2fe32ebd686258b1 {
		return &__obf_f73bdd1883284ee0{}
	}
	if __obf_29ebd0f2c324f5ea == __obf_5c597776869c8984 {
		return &__obf_bd17dda1a88cca5a{}
	}
	return nil
}

type __obf_f73bdd1883284ee0 struct {
}

func (__obf_be93ede593e1d304 *__obf_f73bdd1883284ee0) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.ReadNil() {
		*((*json.RawMessage)(__obf_2a1474febb02279b)) = nil
	} else {
		*((*json.RawMessage)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.SkipAndReturnBytes()
	}
}

func (__obf_be93ede593e1d304 *__obf_f73bdd1883284ee0) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if *((*json.RawMessage)(__obf_2a1474febb02279b)) == nil {
		__obf_850a7457bb739a32.
			WriteNil()
	} else {
		__obf_850a7457bb739a32.
			WriteRaw(string(*((*json.RawMessage)(__obf_2a1474febb02279b))))
	}
}

func (__obf_be93ede593e1d304 *__obf_f73bdd1883284ee0) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(__obf_2a1474febb02279b))) == 0
}

type __obf_bd17dda1a88cca5a struct {
}

func (__obf_be93ede593e1d304 *__obf_bd17dda1a88cca5a) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.ReadNil() {
		*((*RawMessage)(__obf_2a1474febb02279b)) = nil
	} else {
		*((*RawMessage)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.SkipAndReturnBytes()
	}
}

func (__obf_be93ede593e1d304 *__obf_bd17dda1a88cca5a) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if *((*RawMessage)(__obf_2a1474febb02279b)) == nil {
		__obf_850a7457bb739a32.
			WriteNil()
	} else {
		__obf_850a7457bb739a32.
			WriteRaw(string(*((*RawMessage)(__obf_2a1474febb02279b))))
	}
}

func (__obf_be93ede593e1d304 *__obf_bd17dda1a88cca5a) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return len(*((*RawMessage)(__obf_2a1474febb02279b))) == 0
}
