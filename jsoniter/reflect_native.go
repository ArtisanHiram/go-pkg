package __obf_91620b895eeff9ed

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

const __obf_848460bf3389f562 = 32 << uintptr(^uintptr(0)>>63)

func __obf_6e137b9e634e5072(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	if __obf_29ebd0f2c324f5ea.Kind() == reflect.Slice && __obf_29ebd0f2c324f5ea.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_786f7abbe25edc1b := __obf_5f6f4a91bd6a3b9b(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
		return &__obf_6bedef1ccbf5157c{__obf_786f7abbe25edc1b: __obf_786f7abbe25edc1b}
	}
	__obf_d1d6ae4b00f1c128 := __obf_29ebd0f2c324f5ea.String()
	__obf_c2f7526bd7f1cf55 := __obf_29ebd0f2c324f5ea.Kind()
	switch __obf_c2f7526bd7f1cf55 {
	case reflect.String:
		if __obf_d1d6ae4b00f1c128 != "string" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_373056bb02a9d8c5{}
	case reflect.Int:
		if __obf_d1d6ae4b00f1c128 != "int" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_09218e3aefd11230{}
		}
		return &__obf_2bc397abeaf66110{}
	case reflect.Int8:
		if __obf_d1d6ae4b00f1c128 != "int8" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_afefd8a9579d4d18{}
	case reflect.Int16:
		if __obf_d1d6ae4b00f1c128 != "int16" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_06eb7ffd13f0c70d{}
	case reflect.Int32:
		if __obf_d1d6ae4b00f1c128 != "int32" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_09218e3aefd11230{}
	case reflect.Int64:
		if __obf_d1d6ae4b00f1c128 != "int64" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_2bc397abeaf66110{}
	case reflect.Uint:
		if __obf_d1d6ae4b00f1c128 != "uint" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_1dcc374b8b09d1e8{}
		}
		return &__obf_f2a80dc61008bc46{}
	case reflect.Uint8:
		if __obf_d1d6ae4b00f1c128 != "uint8" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_a6269d3758ec2960{}
	case reflect.Uint16:
		if __obf_d1d6ae4b00f1c128 != "uint16" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_a1a8233daacf7654{}
	case reflect.Uint32:
		if __obf_d1d6ae4b00f1c128 != "uint32" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_1dcc374b8b09d1e8{}
	case reflect.Uintptr:
		if __obf_d1d6ae4b00f1c128 != "uintptr" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_848460bf3389f562 == 32 {
			return &__obf_1dcc374b8b09d1e8{}
		}
		return &__obf_f2a80dc61008bc46{}
	case reflect.Uint64:
		if __obf_d1d6ae4b00f1c128 != "uint64" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_f2a80dc61008bc46{}
	case reflect.Float32:
		if __obf_d1d6ae4b00f1c128 != "float32" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_b91a249e1b5d5db1{}
	case reflect.Float64:
		if __obf_d1d6ae4b00f1c128 != "float64" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_0070bd3970835dfa{}
	case reflect.Bool:
		if __obf_d1d6ae4b00f1c128 != "bool" {
			return __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_7acff277aed4f6ea{}
	}
	return nil
}

func __obf_f1b7f12340262e83(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	if __obf_29ebd0f2c324f5ea.Kind() == reflect.Slice && __obf_29ebd0f2c324f5ea.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_786f7abbe25edc1b := __obf_5f6f4a91bd6a3b9b(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
		return &__obf_6bedef1ccbf5157c{__obf_786f7abbe25edc1b: __obf_786f7abbe25edc1b}
	}
	__obf_d1d6ae4b00f1c128 := __obf_29ebd0f2c324f5ea.String()
	switch __obf_29ebd0f2c324f5ea.Kind() {
	case reflect.String:
		if __obf_d1d6ae4b00f1c128 != "string" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_373056bb02a9d8c5{}
	case reflect.Int:
		if __obf_d1d6ae4b00f1c128 != "int" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_09218e3aefd11230{}
		}
		return &__obf_2bc397abeaf66110{}
	case reflect.Int8:
		if __obf_d1d6ae4b00f1c128 != "int8" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_afefd8a9579d4d18{}
	case reflect.Int16:
		if __obf_d1d6ae4b00f1c128 != "int16" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_06eb7ffd13f0c70d{}
	case reflect.Int32:
		if __obf_d1d6ae4b00f1c128 != "int32" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_09218e3aefd11230{}
	case reflect.Int64:
		if __obf_d1d6ae4b00f1c128 != "int64" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_2bc397abeaf66110{}
	case reflect.Uint:
		if __obf_d1d6ae4b00f1c128 != "uint" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_1dcc374b8b09d1e8{}
		}
		return &__obf_f2a80dc61008bc46{}
	case reflect.Uint8:
		if __obf_d1d6ae4b00f1c128 != "uint8" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_a6269d3758ec2960{}
	case reflect.Uint16:
		if __obf_d1d6ae4b00f1c128 != "uint16" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_a1a8233daacf7654{}
	case reflect.Uint32:
		if __obf_d1d6ae4b00f1c128 != "uint32" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_1dcc374b8b09d1e8{}
	case reflect.Uintptr:
		if __obf_d1d6ae4b00f1c128 != "uintptr" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_848460bf3389f562 == 32 {
			return &__obf_1dcc374b8b09d1e8{}
		}
		return &__obf_f2a80dc61008bc46{}
	case reflect.Uint64:
		if __obf_d1d6ae4b00f1c128 != "uint64" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_f2a80dc61008bc46{}
	case reflect.Float32:
		if __obf_d1d6ae4b00f1c128 != "float32" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_b91a249e1b5d5db1{}
	case reflect.Float64:
		if __obf_d1d6ae4b00f1c128 != "float64" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_0070bd3970835dfa{}
	case reflect.Bool:
		if __obf_d1d6ae4b00f1c128 != "bool" {
			return __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_7acff277aed4f6ea{}
	}
	return nil
}

type __obf_373056bb02a9d8c5 struct {
}

func (__obf_be93ede593e1d304 *__obf_373056bb02a9d8c5) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	*((*string)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadString()
}

func (__obf_be93ede593e1d304 *__obf_373056bb02a9d8c5) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_e91bd2feb751e4f1 := *((*string)(__obf_2a1474febb02279b))
	__obf_850a7457bb739a32.
		WriteString(__obf_e91bd2feb751e4f1)
}

func (__obf_be93ede593e1d304 *__obf_373056bb02a9d8c5) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*string)(__obf_2a1474febb02279b)) == ""
}

type __obf_afefd8a9579d4d18 struct {
}

func (__obf_be93ede593e1d304 *__obf_afefd8a9579d4d18) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*int8)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadInt8()
	}
}

func (__obf_be93ede593e1d304 *__obf_afefd8a9579d4d18) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteInt8(*((*int8)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_afefd8a9579d4d18) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*int8)(__obf_2a1474febb02279b)) == 0
}

type __obf_06eb7ffd13f0c70d struct {
}

func (__obf_be93ede593e1d304 *__obf_06eb7ffd13f0c70d) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*int16)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadInt16()
	}
}

func (__obf_be93ede593e1d304 *__obf_06eb7ffd13f0c70d) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteInt16(*((*int16)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_06eb7ffd13f0c70d) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*int16)(__obf_2a1474febb02279b)) == 0
}

type __obf_09218e3aefd11230 struct {
}

func (__obf_be93ede593e1d304 *__obf_09218e3aefd11230) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*int32)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadInt32()
	}
}

func (__obf_be93ede593e1d304 *__obf_09218e3aefd11230) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteInt32(*((*int32)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_09218e3aefd11230) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*int32)(__obf_2a1474febb02279b)) == 0
}

type __obf_2bc397abeaf66110 struct {
}

func (__obf_be93ede593e1d304 *__obf_2bc397abeaf66110) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*int64)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadInt64()
	}
}

func (__obf_be93ede593e1d304 *__obf_2bc397abeaf66110) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteInt64(*((*int64)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_2bc397abeaf66110) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*int64)(__obf_2a1474febb02279b)) == 0
}

type __obf_a6269d3758ec2960 struct {
}

func (__obf_be93ede593e1d304 *__obf_a6269d3758ec2960) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*uint8)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadUint8()
	}
}

func (__obf_be93ede593e1d304 *__obf_a6269d3758ec2960) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteUint8(*((*uint8)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_a6269d3758ec2960) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*uint8)(__obf_2a1474febb02279b)) == 0
}

type __obf_a1a8233daacf7654 struct {
}

func (__obf_be93ede593e1d304 *__obf_a1a8233daacf7654) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*uint16)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadUint16()
	}
}

func (__obf_be93ede593e1d304 *__obf_a1a8233daacf7654) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteUint16(*((*uint16)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_a1a8233daacf7654) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*uint16)(__obf_2a1474febb02279b)) == 0
}

type __obf_1dcc374b8b09d1e8 struct {
}

func (__obf_be93ede593e1d304 *__obf_1dcc374b8b09d1e8) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*uint32)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadUint32()
	}
}

func (__obf_be93ede593e1d304 *__obf_1dcc374b8b09d1e8) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteUint32(*((*uint32)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_1dcc374b8b09d1e8) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*uint32)(__obf_2a1474febb02279b)) == 0
}

type __obf_f2a80dc61008bc46 struct {
}

func (__obf_be93ede593e1d304 *__obf_f2a80dc61008bc46) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*uint64)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadUint64()
	}
}

func (__obf_be93ede593e1d304 *__obf_f2a80dc61008bc46) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteUint64(*((*uint64)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_f2a80dc61008bc46) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*uint64)(__obf_2a1474febb02279b)) == 0
}

type __obf_b91a249e1b5d5db1 struct {
}

func (__obf_be93ede593e1d304 *__obf_b91a249e1b5d5db1) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*float32)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadFloat32()
	}
}

func (__obf_be93ede593e1d304 *__obf_b91a249e1b5d5db1) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteFloat32(*((*float32)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_b91a249e1b5d5db1) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*float32)(__obf_2a1474febb02279b)) == 0
}

type __obf_0070bd3970835dfa struct {
}

func (__obf_be93ede593e1d304 *__obf_0070bd3970835dfa) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*float64)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadFloat64()
	}
}

func (__obf_be93ede593e1d304 *__obf_0070bd3970835dfa) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteFloat64(*((*float64)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_0070bd3970835dfa) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*float64)(__obf_2a1474febb02279b)) == 0
}

type __obf_7acff277aed4f6ea struct {
}

func (__obf_be93ede593e1d304 *__obf_7acff277aed4f6ea) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.ReadNil() {
		*((*bool)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.ReadBool()
	}
}

func (__obf_be93ede593e1d304 *__obf_7acff277aed4f6ea) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteBool(*((*bool)(__obf_2a1474febb02279b)))
}

func (__obf_be93ede593e1d304 *__obf_7acff277aed4f6ea) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return !(*((*bool)(__obf_2a1474febb02279b)))
}

type __obf_6bedef1ccbf5157c struct {
	__obf_56bad1827144ee34 *reflect2.UnsafeSliceType
	__obf_786f7abbe25edc1b ValDecoder
}

func (__obf_be93ede593e1d304 *__obf_6bedef1ccbf5157c) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.ReadNil() {
		__obf_be93ede593e1d304.__obf_56bad1827144ee34.
			UnsafeSetNil(__obf_2a1474febb02279b)
		return
	}
	switch __obf_1bb30e8a74ed8233.WhatIsNext() {
	case StringValue:
		__obf_0a3212be68b8bb4a := __obf_1bb30e8a74ed8233.ReadString()
		__obf_84e389f346f83d04, __obf_62967739bca1d11d := base64.StdEncoding.DecodeString(__obf_0a3212be68b8bb4a)
		if __obf_62967739bca1d11d != nil {
			__obf_1bb30e8a74ed8233.
				ReportError("decode base64", __obf_62967739bca1d11d.Error())
		} else {
			__obf_be93ede593e1d304.__obf_56bad1827144ee34.
				UnsafeSet(__obf_2a1474febb02279b, unsafe.Pointer(&__obf_84e389f346f83d04))
		}
	case ArrayValue:
		__obf_be93ede593e1d304.__obf_786f7abbe25edc1b.
			Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	default:
		__obf_1bb30e8a74ed8233.
			ReportError("base64Codec", "invalid input")
	}
}

func (__obf_be93ede593e1d304 *__obf_6bedef1ccbf5157c) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	if __obf_be93ede593e1d304.__obf_56bad1827144ee34.UnsafeIsNil(__obf_2a1474febb02279b) {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_0a3212be68b8bb4a := *((*[]byte)(__obf_2a1474febb02279b))
	__obf_db8658353e4b2677 := base64.StdEncoding
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
	if len(__obf_0a3212be68b8bb4a) != 0 {
		__obf_dc51447b63477324 := __obf_db8658353e4b2677.EncodedLen(len(__obf_0a3212be68b8bb4a))
		__obf_184433571fa55237 := make([]byte, __obf_dc51447b63477324)
		__obf_db8658353e4b2677.
			Encode(__obf_184433571fa55237, __obf_0a3212be68b8bb4a)
		__obf_850a7457bb739a32.__obf_184433571fa55237 = append(__obf_850a7457bb739a32.__obf_184433571fa55237, __obf_184433571fa55237...)
	}
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
}

func (__obf_be93ede593e1d304 *__obf_6bedef1ccbf5157c) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_2a1474febb02279b))) == 0
}
