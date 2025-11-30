package __obf_c3cd04a15c56f32f

import (
	"encoding/json"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Number string

// String returns the literal text of the number.
func (__obf_fd4464bb6b13cabd Number) String() string { return string(__obf_fd4464bb6b13cabd) }

// Float64 returns the number as a float64.
func (__obf_fd4464bb6b13cabd Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(__obf_fd4464bb6b13cabd), 64)
}

// Int64 returns the number as an int64.
func (__obf_fd4464bb6b13cabd Number) Int64() (int64, error) {
	return strconv.ParseInt(string(__obf_fd4464bb6b13cabd), 10, 64)
}

func CastJsonNumber(__obf_d59813f23ad740a8 any) (string, bool) {
	switch __obf_d9a4b362e245b912 := __obf_d59813f23ad740a8.(type) {
	case json.Number:
		return string(__obf_d9a4b362e245b912), true
	case Number:
		return string(__obf_d9a4b362e245b912), true
	}
	return "", false
}

var __obf_f1b6c81c265849e1 = reflect2.TypeOfPtr((*json.Number)(nil)).Elem()
var __obf_6897610178c6ad64 = reflect2.TypeOfPtr((*Number)(nil)).Elem()

func __obf_2ffedf0111a1f93e(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	if __obf_8667d4fc2a95b0d7.AssignableTo(__obf_f1b6c81c265849e1) {
		return &__obf_15af08b766d93982{}
	}
	if __obf_8667d4fc2a95b0d7.AssignableTo(__obf_6897610178c6ad64) {
		return &__obf_ab91a4d7d689c468{}
	}
	return nil
}

func __obf_28c9f319ac0c0137(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	if __obf_8667d4fc2a95b0d7.AssignableTo(__obf_f1b6c81c265849e1) {
		return &__obf_15af08b766d93982{}
	}
	if __obf_8667d4fc2a95b0d7.AssignableTo(__obf_6897610178c6ad64) {
		return &__obf_ab91a4d7d689c468{}
	}
	return nil
}

type __obf_15af08b766d93982 struct {
}

func (__obf_1569d098873d2228 *__obf_15af08b766d93982) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	switch __obf_edd9fe48d39445e4.WhatIsNext() {
	case StringValue:
		*((*json.Number)(__obf_753ab3fb72cdd659)) = json.Number(__obf_edd9fe48d39445e4.ReadString())
	case NilValue:
		__obf_edd9fe48d39445e4.__obf_f22f308da0537336('n', 'u', 'l', 'l')
		*((*json.Number)(__obf_753ab3fb72cdd659)) = ""
	default:
		*((*json.Number)(__obf_753ab3fb72cdd659)) = json.Number([]byte(__obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818()))
	}
}

func (__obf_1569d098873d2228 *__obf_15af08b766d93982) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_45a19a0c38cb4f7c := *((*json.Number)(__obf_753ab3fb72cdd659))
	if len(__obf_45a19a0c38cb4f7c) == 0 {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('0')
	} else {
		__obf_2361f5214e84c60f.
			WriteRaw(string(__obf_45a19a0c38cb4f7c))
	}
}

func (__obf_1569d098873d2228 *__obf_15af08b766d93982) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return len(*((*json.Number)(__obf_753ab3fb72cdd659))) == 0
}

type __obf_ab91a4d7d689c468 struct {
}

func (__obf_1569d098873d2228 *__obf_ab91a4d7d689c468) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	switch __obf_edd9fe48d39445e4.WhatIsNext() {
	case StringValue:
		*((*Number)(__obf_753ab3fb72cdd659)) = Number(__obf_edd9fe48d39445e4.ReadString())
	case NilValue:
		__obf_edd9fe48d39445e4.__obf_f22f308da0537336('n', 'u', 'l', 'l')
		*((*Number)(__obf_753ab3fb72cdd659)) = ""
	default:
		*((*Number)(__obf_753ab3fb72cdd659)) = Number([]byte(__obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818()))
	}
}

func (__obf_1569d098873d2228 *__obf_ab91a4d7d689c468) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_45a19a0c38cb4f7c := *((*Number)(__obf_753ab3fb72cdd659))
	if len(__obf_45a19a0c38cb4f7c) == 0 {
		__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('0')
	} else {
		__obf_2361f5214e84c60f.
			WriteRaw(string(__obf_45a19a0c38cb4f7c))
	}
}

func (__obf_1569d098873d2228 *__obf_ab91a4d7d689c468) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return len(*((*Number)(__obf_753ab3fb72cdd659))) == 0
}
