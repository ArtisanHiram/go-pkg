package __obf_91620b895eeff9ed

import (
	"encoding/json"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Number string

// String returns the literal text of the number.
func (__obf_c80a670e818798d8 Number) String() string { return string(__obf_c80a670e818798d8) }

// Float64 returns the number as a float64.
func (__obf_c80a670e818798d8 Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(__obf_c80a670e818798d8), 64)
}

// Int64 returns the number as an int64.
func (__obf_c80a670e818798d8 Number) Int64() (int64, error) {
	return strconv.ParseInt(string(__obf_c80a670e818798d8), 10, 64)
}

func CastJsonNumber(__obf_bbfd2ac8618a6f0c any) (string, bool) {
	switch __obf_c216af0b988110a5 := __obf_bbfd2ac8618a6f0c.(type) {
	case json.Number:
		return string(__obf_c216af0b988110a5), true
	case Number:
		return string(__obf_c216af0b988110a5), true
	}
	return "", false
}

var __obf_f4b618955fc897b9 = reflect2.TypeOfPtr((*json.Number)(nil)).Elem()
var __obf_7ccf1dff904af924 = reflect2.TypeOfPtr((*Number)(nil)).Elem()

func __obf_5a852a950afc440b(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	if __obf_29ebd0f2c324f5ea.AssignableTo(__obf_f4b618955fc897b9) {
		return &__obf_e606fdbff86c4c4a{}
	}
	if __obf_29ebd0f2c324f5ea.AssignableTo(__obf_7ccf1dff904af924) {
		return &__obf_3c1e78a4ba0e29a6{}
	}
	return nil
}

func __obf_2d74fdc730e873e9(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	if __obf_29ebd0f2c324f5ea.AssignableTo(__obf_f4b618955fc897b9) {
		return &__obf_e606fdbff86c4c4a{}
	}
	if __obf_29ebd0f2c324f5ea.AssignableTo(__obf_7ccf1dff904af924) {
		return &__obf_3c1e78a4ba0e29a6{}
	}
	return nil
}

type __obf_e606fdbff86c4c4a struct {
}

func (__obf_be93ede593e1d304 *__obf_e606fdbff86c4c4a) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	switch __obf_1bb30e8a74ed8233.WhatIsNext() {
	case StringValue:
		*((*json.Number)(__obf_2a1474febb02279b)) = json.Number(__obf_1bb30e8a74ed8233.ReadString())
	case NilValue:
		__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('n', 'u', 'l', 'l')
		*((*json.Number)(__obf_2a1474febb02279b)) = ""
	default:
		*((*json.Number)(__obf_2a1474febb02279b)) = json.Number([]byte(__obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd()))
	}
}

func (__obf_be93ede593e1d304 *__obf_e606fdbff86c4c4a) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_7649f8370f217c73 := *((*json.Number)(__obf_2a1474febb02279b))
	if len(__obf_7649f8370f217c73) == 0 {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078('0')
	} else {
		__obf_850a7457bb739a32.
			WriteRaw(string(__obf_7649f8370f217c73))
	}
}

func (__obf_be93ede593e1d304 *__obf_e606fdbff86c4c4a) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return len(*((*json.Number)(__obf_2a1474febb02279b))) == 0
}

type __obf_3c1e78a4ba0e29a6 struct {
}

func (__obf_be93ede593e1d304 *__obf_3c1e78a4ba0e29a6) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	switch __obf_1bb30e8a74ed8233.WhatIsNext() {
	case StringValue:
		*((*Number)(__obf_2a1474febb02279b)) = Number(__obf_1bb30e8a74ed8233.ReadString())
	case NilValue:
		__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('n', 'u', 'l', 'l')
		*((*Number)(__obf_2a1474febb02279b)) = ""
	default:
		*((*Number)(__obf_2a1474febb02279b)) = Number([]byte(__obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd()))
	}
}

func (__obf_be93ede593e1d304 *__obf_3c1e78a4ba0e29a6) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_7649f8370f217c73 := *((*Number)(__obf_2a1474febb02279b))
	if len(__obf_7649f8370f217c73) == 0 {
		__obf_850a7457bb739a32.__obf_72837f6fe737f078('0')
	} else {
		__obf_850a7457bb739a32.
			WriteRaw(string(__obf_7649f8370f217c73))
	}
}

func (__obf_be93ede593e1d304 *__obf_3c1e78a4ba0e29a6) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return len(*((*Number)(__obf_2a1474febb02279b))) == 0
}
