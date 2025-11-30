package __obf_030d39f7a8de96c6

import (
	"encoding/json"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Number string

// String returns the literal text of the number.
func (__obf_e40b3963a92ca4a0 Number) String() string { return string(__obf_e40b3963a92ca4a0) }

// Float64 returns the number as a float64.
func (__obf_e40b3963a92ca4a0 Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(__obf_e40b3963a92ca4a0), 64)
}

// Int64 returns the number as an int64.
func (__obf_e40b3963a92ca4a0 Number) Int64() (int64, error) {
	return strconv.ParseInt(string(__obf_e40b3963a92ca4a0), 10, 64)
}

func CastJsonNumber(__obf_efaf2719b44ad8ac any) (string, bool) {
	switch __obf_c37412231cc0308f := __obf_efaf2719b44ad8ac.(type) {
	case json.Number:
		return string(__obf_c37412231cc0308f), true
	case Number:
		return string(__obf_c37412231cc0308f), true
	}
	return "", false
}

var __obf_a07628042142fb03 = reflect2.TypeOfPtr((*json.Number)(nil)).Elem()
var __obf_a1a3c58280781de3 = reflect2.TypeOfPtr((*Number)(nil)).Elem()

func __obf_18e6804367792db0(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	if __obf_8744d0b8c80ccdc1.AssignableTo(__obf_a07628042142fb03) {
		return &__obf_7baf60b8c4c1075e{}
	}
	if __obf_8744d0b8c80ccdc1.AssignableTo(__obf_a1a3c58280781de3) {
		return &__obf_e51bfb67e188ae32{}
	}
	return nil
}

func __obf_21ace3604d05bc8e(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	if __obf_8744d0b8c80ccdc1.AssignableTo(__obf_a07628042142fb03) {
		return &__obf_7baf60b8c4c1075e{}
	}
	if __obf_8744d0b8c80ccdc1.AssignableTo(__obf_a1a3c58280781de3) {
		return &__obf_e51bfb67e188ae32{}
	}
	return nil
}

type __obf_7baf60b8c4c1075e struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_7baf60b8c4c1075e) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	switch __obf_4ab56a99965952e7.WhatIsNext() {
	case StringValue:
		*((*json.Number)(__obf_dbbf371b91136494)) = json.Number(__obf_4ab56a99965952e7.ReadString())
	case NilValue:
		__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('n', 'u', 'l', 'l')
		*((*json.Number)(__obf_dbbf371b91136494)) = ""
	default:
		*((*json.Number)(__obf_dbbf371b91136494)) = json.Number([]byte(__obf_4ab56a99965952e7.__obf_0ba08d27498a0d84()))
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_7baf60b8c4c1075e) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_83b1e63b792d512f := *((*json.Number)(__obf_dbbf371b91136494))
	if len(__obf_83b1e63b792d512f) == 0 {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53('0')
	} else {
		__obf_8a2c51fe22775655.
			WriteRaw(string(__obf_83b1e63b792d512f))
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_7baf60b8c4c1075e) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return len(*((*json.Number)(__obf_dbbf371b91136494))) == 0
}

type __obf_e51bfb67e188ae32 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_e51bfb67e188ae32) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	switch __obf_4ab56a99965952e7.WhatIsNext() {
	case StringValue:
		*((*Number)(__obf_dbbf371b91136494)) = Number(__obf_4ab56a99965952e7.ReadString())
	case NilValue:
		__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('n', 'u', 'l', 'l')
		*((*Number)(__obf_dbbf371b91136494)) = ""
	default:
		*((*Number)(__obf_dbbf371b91136494)) = Number([]byte(__obf_4ab56a99965952e7.__obf_0ba08d27498a0d84()))
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_e51bfb67e188ae32) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_83b1e63b792d512f := *((*Number)(__obf_dbbf371b91136494))
	if len(__obf_83b1e63b792d512f) == 0 {
		__obf_8a2c51fe22775655.__obf_41130daa346c0e53('0')
	} else {
		__obf_8a2c51fe22775655.
			WriteRaw(string(__obf_83b1e63b792d512f))
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_e51bfb67e188ae32) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return len(*((*Number)(__obf_dbbf371b91136494))) == 0
}
