package __obf_5b802ce8d9ba56d6

import (
	"encoding/json"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Number string

// String returns the literal text of the number.
func (__obf_2c0ce853cb81f014 Number) String() string { return string(__obf_2c0ce853cb81f014) }

// Float64 returns the number as a float64.
func (__obf_2c0ce853cb81f014 Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(__obf_2c0ce853cb81f014), 64)
}

// Int64 returns the number as an int64.
func (__obf_2c0ce853cb81f014 Number) Int64() (int64, error) {
	return strconv.ParseInt(string(__obf_2c0ce853cb81f014), 10, 64)
}

func CastJsonNumber(__obf_5406b11e33b9d1d3 any) (string, bool) {
	switch __obf_6cf8e62213a6c8a1 := __obf_5406b11e33b9d1d3.(type) {
	case json.Number:
		return string(__obf_6cf8e62213a6c8a1), true
	case Number:
		return string(__obf_6cf8e62213a6c8a1), true
	}
	return "", false
}

var __obf_813655adfe611f16 = reflect2.TypeOfPtr((*json.Number)(nil)).Elem()
var __obf_8bdf27926c956099 = reflect2.TypeOfPtr((*Number)(nil)).Elem()

func __obf_3fe7eb60741bd404(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	if __obf_5efc66d8c338c133.AssignableTo(__obf_813655adfe611f16) {
		return &__obf_8f17744532c03fb1{}
	}
	if __obf_5efc66d8c338c133.AssignableTo(__obf_8bdf27926c956099) {
		return &__obf_fa525ebe50d04cd4{}
	}
	return nil
}

func __obf_4e38aa1ee94cd7cf(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	if __obf_5efc66d8c338c133.AssignableTo(__obf_813655adfe611f16) {
		return &__obf_8f17744532c03fb1{}
	}
	if __obf_5efc66d8c338c133.AssignableTo(__obf_8bdf27926c956099) {
		return &__obf_fa525ebe50d04cd4{}
	}
	return nil
}

type __obf_8f17744532c03fb1 struct {
}

func (__obf_0099bc81efe356aa *__obf_8f17744532c03fb1) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	switch __obf_67008a6a9e5ba828.WhatIsNext() {
	case StringValue:
		*((*json.Number)(__obf_d3c919547bf7e47a)) = json.Number(__obf_67008a6a9e5ba828.ReadString())
	case NilValue:
		__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('n', 'u', 'l', 'l')
		*((*json.Number)(__obf_d3c919547bf7e47a)) = ""
	default:
		*((*json.Number)(__obf_d3c919547bf7e47a)) = json.Number([]byte(__obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05()))
	}
}

func (__obf_0099bc81efe356aa *__obf_8f17744532c03fb1) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_553f22e0c3850071 := *((*json.Number)(__obf_d3c919547bf7e47a))
	if len(__obf_553f22e0c3850071) == 0 {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('0')
	} else {
		__obf_00fc491caa967a74.
			WriteRaw(string(__obf_553f22e0c3850071))
	}
}

func (__obf_0099bc81efe356aa *__obf_8f17744532c03fb1) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return len(*((*json.Number)(__obf_d3c919547bf7e47a))) == 0
}

type __obf_fa525ebe50d04cd4 struct {
}

func (__obf_0099bc81efe356aa *__obf_fa525ebe50d04cd4) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	switch __obf_67008a6a9e5ba828.WhatIsNext() {
	case StringValue:
		*((*Number)(__obf_d3c919547bf7e47a)) = Number(__obf_67008a6a9e5ba828.ReadString())
	case NilValue:
		__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('n', 'u', 'l', 'l')
		*((*Number)(__obf_d3c919547bf7e47a)) = ""
	default:
		*((*Number)(__obf_d3c919547bf7e47a)) = Number([]byte(__obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05()))
	}
}

func (__obf_0099bc81efe356aa *__obf_fa525ebe50d04cd4) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_553f22e0c3850071 := *((*Number)(__obf_d3c919547bf7e47a))
	if len(__obf_553f22e0c3850071) == 0 {
		__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('0')
	} else {
		__obf_00fc491caa967a74.
			WriteRaw(string(__obf_553f22e0c3850071))
	}
}

func (__obf_0099bc81efe356aa *__obf_fa525ebe50d04cd4) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return len(*((*Number)(__obf_d3c919547bf7e47a))) == 0
}
