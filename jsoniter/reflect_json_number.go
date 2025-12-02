package __obf_c7b79b12b33d8238

import (
	"encoding/json"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Number string

// String returns the literal text of the number.
func (__obf_ff4330e73e137d5c Number) String() string { return string(__obf_ff4330e73e137d5c) }

// Float64 returns the number as a float64.
func (__obf_ff4330e73e137d5c Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(__obf_ff4330e73e137d5c), 64)
}

// Int64 returns the number as an int64.
func (__obf_ff4330e73e137d5c Number) Int64() (int64, error) {
	return strconv.ParseInt(string(__obf_ff4330e73e137d5c), 10, 64)
}

func CastJsonNumber(__obf_35accd096612ebe4 any) (string, bool) {
	switch __obf_692bf907784e1cdc := __obf_35accd096612ebe4.(type) {
	case json.Number:
		return string(__obf_692bf907784e1cdc), true
	case Number:
		return string(__obf_692bf907784e1cdc), true
	}
	return "", false
}

var __obf_2ffc8416cba6db6b = reflect2.TypeOfPtr((*json.Number)(nil)).Elem()
var __obf_7ce26d2c13fb639a = reflect2.TypeOfPtr((*Number)(nil)).Elem()

func __obf_6bf3f557fcfe51c2(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	if __obf_edcded11af6ebc4c.AssignableTo(__obf_2ffc8416cba6db6b) {
		return &__obf_cf9ebb7f6cbcb886{}
	}
	if __obf_edcded11af6ebc4c.AssignableTo(__obf_7ce26d2c13fb639a) {
		return &__obf_076df0511d52b273{}
	}
	return nil
}

func __obf_84a8ea4f190dbe8e(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	if __obf_edcded11af6ebc4c.AssignableTo(__obf_2ffc8416cba6db6b) {
		return &__obf_cf9ebb7f6cbcb886{}
	}
	if __obf_edcded11af6ebc4c.AssignableTo(__obf_7ce26d2c13fb639a) {
		return &__obf_076df0511d52b273{}
	}
	return nil
}

type __obf_cf9ebb7f6cbcb886 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_cf9ebb7f6cbcb886) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	switch __obf_0da8c843dad13139.WhatIsNext() {
	case StringValue:
		*((*json.Number)(__obf_575c04f2b9d91315)) = json.Number(__obf_0da8c843dad13139.ReadString())
	case NilValue:
		__obf_0da8c843dad13139.__obf_62908d9424a8486b('n', 'u', 'l', 'l')
		*((*json.Number)(__obf_575c04f2b9d91315)) = ""
	default:
		*((*json.Number)(__obf_575c04f2b9d91315)) = json.Number([]byte(__obf_0da8c843dad13139.__obf_a678d14dd84cbb6b()))
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_cf9ebb7f6cbcb886) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c71547d6cfcb55da := *((*json.Number)(__obf_575c04f2b9d91315))
	if len(__obf_c71547d6cfcb55da) == 0 {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('0')
	} else {
		__obf_d8f50bcfe87d8b47.
			WriteRaw(string(__obf_c71547d6cfcb55da))
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_cf9ebb7f6cbcb886) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return len(*((*json.Number)(__obf_575c04f2b9d91315))) == 0
}

type __obf_076df0511d52b273 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_076df0511d52b273) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	switch __obf_0da8c843dad13139.WhatIsNext() {
	case StringValue:
		*((*Number)(__obf_575c04f2b9d91315)) = Number(__obf_0da8c843dad13139.ReadString())
	case NilValue:
		__obf_0da8c843dad13139.__obf_62908d9424a8486b('n', 'u', 'l', 'l')
		*((*Number)(__obf_575c04f2b9d91315)) = ""
	default:
		*((*Number)(__obf_575c04f2b9d91315)) = Number([]byte(__obf_0da8c843dad13139.__obf_a678d14dd84cbb6b()))
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_076df0511d52b273) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c71547d6cfcb55da := *((*Number)(__obf_575c04f2b9d91315))
	if len(__obf_c71547d6cfcb55da) == 0 {
		__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('0')
	} else {
		__obf_d8f50bcfe87d8b47.
			WriteRaw(string(__obf_c71547d6cfcb55da))
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_076df0511d52b273) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return len(*((*Number)(__obf_575c04f2b9d91315))) == 0
}
