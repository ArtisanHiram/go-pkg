package __obf_030d39f7a8de96c6

import (
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_4ae27c40a2fad41c = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var __obf_243dc45a11f19155 = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func __obf_8a2ee1f3ca9358c1(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	if __obf_8744d0b8c80ccdc1 == __obf_4ae27c40a2fad41c {
		return &__obf_d9b7e027154b2a2d{}
	}
	if __obf_8744d0b8c80ccdc1 == __obf_243dc45a11f19155 {
		return &__obf_a4917da526b99f67{}
	}
	return nil
}

func __obf_3d1edc0e2c1c1e07(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	if __obf_8744d0b8c80ccdc1 == __obf_4ae27c40a2fad41c {
		return &__obf_d9b7e027154b2a2d{}
	}
	if __obf_8744d0b8c80ccdc1 == __obf_243dc45a11f19155 {
		return &__obf_a4917da526b99f67{}
	}
	return nil
}

type __obf_d9b7e027154b2a2d struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_d9b7e027154b2a2d) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.ReadNil() {
		*((*json.RawMessage)(__obf_dbbf371b91136494)) = nil
	} else {
		*((*json.RawMessage)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.SkipAndReturnBytes()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_d9b7e027154b2a2d) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if *((*json.RawMessage)(__obf_dbbf371b91136494)) == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
	} else {
		__obf_8a2c51fe22775655.
			WriteRaw(string(*((*json.RawMessage)(__obf_dbbf371b91136494))))
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_d9b7e027154b2a2d) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(__obf_dbbf371b91136494))) == 0
}

type __obf_a4917da526b99f67 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_a4917da526b99f67) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.ReadNil() {
		*((*RawMessage)(__obf_dbbf371b91136494)) = nil
	} else {
		*((*RawMessage)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.SkipAndReturnBytes()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_a4917da526b99f67) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if *((*RawMessage)(__obf_dbbf371b91136494)) == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
	} else {
		__obf_8a2c51fe22775655.
			WriteRaw(string(*((*RawMessage)(__obf_dbbf371b91136494))))
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_a4917da526b99f67) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return len(*((*RawMessage)(__obf_dbbf371b91136494))) == 0
}
