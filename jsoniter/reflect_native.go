package __obf_5b802ce8d9ba56d6

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

const __obf_1b422ac0cf965f7a = 32 << uintptr(^uintptr(0)>>63)

func __obf_062c5f9f03d60bc9(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	if __obf_5efc66d8c338c133.Kind() == reflect.Slice && __obf_5efc66d8c338c133.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_fae02aecb00c25f9 := __obf_868f472bae9503f0(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
		return &__obf_a18e9d92ad8c34a9{__obf_fae02aecb00c25f9: __obf_fae02aecb00c25f9}
	}
	__obf_0d9188c359cc43b1 := __obf_5efc66d8c338c133.String()
	__obf_f1764122678433dc := __obf_5efc66d8c338c133.Kind()
	switch __obf_f1764122678433dc {
	case reflect.String:
		if __obf_0d9188c359cc43b1 != "string" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_6e64bbe5dcd439c3{}
	case reflect.Int:
		if __obf_0d9188c359cc43b1 != "int" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_aeb5e7cae1d3d27e{}
		}
		return &__obf_92c464f4f16e071c{}
	case reflect.Int8:
		if __obf_0d9188c359cc43b1 != "int8" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_31bb730b2e917d9d{}
	case reflect.Int16:
		if __obf_0d9188c359cc43b1 != "int16" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_5c68bf6a9fce0257{}
	case reflect.Int32:
		if __obf_0d9188c359cc43b1 != "int32" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_aeb5e7cae1d3d27e{}
	case reflect.Int64:
		if __obf_0d9188c359cc43b1 != "int64" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_92c464f4f16e071c{}
	case reflect.Uint:
		if __obf_0d9188c359cc43b1 != "uint" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_466737d69e243d07{}
		}
		return &__obf_cc9e5c5a9ce7e2d1{}
	case reflect.Uint8:
		if __obf_0d9188c359cc43b1 != "uint8" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_72de03bc5292cee4{}
	case reflect.Uint16:
		if __obf_0d9188c359cc43b1 != "uint16" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_6d8a7372e7b2174f{}
	case reflect.Uint32:
		if __obf_0d9188c359cc43b1 != "uint32" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_466737d69e243d07{}
	case reflect.Uintptr:
		if __obf_0d9188c359cc43b1 != "uintptr" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_1b422ac0cf965f7a == 32 {
			return &__obf_466737d69e243d07{}
		}
		return &__obf_cc9e5c5a9ce7e2d1{}
	case reflect.Uint64:
		if __obf_0d9188c359cc43b1 != "uint64" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_cc9e5c5a9ce7e2d1{}
	case reflect.Float32:
		if __obf_0d9188c359cc43b1 != "float32" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_267f314242d45e83{}
	case reflect.Float64:
		if __obf_0d9188c359cc43b1 != "float64" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_9faca152bf30e4b0{}
	case reflect.Bool:
		if __obf_0d9188c359cc43b1 != "bool" {
			return __obf_3bb380f7abc03208(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_d0a75414375dd73f{}
	}
	return nil
}

func __obf_145e07780e2db77f(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	if __obf_5efc66d8c338c133.Kind() == reflect.Slice && __obf_5efc66d8c338c133.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_fae02aecb00c25f9 := __obf_868f472bae9503f0(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
		return &__obf_a18e9d92ad8c34a9{__obf_fae02aecb00c25f9: __obf_fae02aecb00c25f9}
	}
	__obf_0d9188c359cc43b1 := __obf_5efc66d8c338c133.String()
	switch __obf_5efc66d8c338c133.Kind() {
	case reflect.String:
		if __obf_0d9188c359cc43b1 != "string" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_6e64bbe5dcd439c3{}
	case reflect.Int:
		if __obf_0d9188c359cc43b1 != "int" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_aeb5e7cae1d3d27e{}
		}
		return &__obf_92c464f4f16e071c{}
	case reflect.Int8:
		if __obf_0d9188c359cc43b1 != "int8" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_31bb730b2e917d9d{}
	case reflect.Int16:
		if __obf_0d9188c359cc43b1 != "int16" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_5c68bf6a9fce0257{}
	case reflect.Int32:
		if __obf_0d9188c359cc43b1 != "int32" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_aeb5e7cae1d3d27e{}
	case reflect.Int64:
		if __obf_0d9188c359cc43b1 != "int64" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_92c464f4f16e071c{}
	case reflect.Uint:
		if __obf_0d9188c359cc43b1 != "uint" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_466737d69e243d07{}
		}
		return &__obf_cc9e5c5a9ce7e2d1{}
	case reflect.Uint8:
		if __obf_0d9188c359cc43b1 != "uint8" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_72de03bc5292cee4{}
	case reflect.Uint16:
		if __obf_0d9188c359cc43b1 != "uint16" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_6d8a7372e7b2174f{}
	case reflect.Uint32:
		if __obf_0d9188c359cc43b1 != "uint32" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_466737d69e243d07{}
	case reflect.Uintptr:
		if __obf_0d9188c359cc43b1 != "uintptr" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_1b422ac0cf965f7a == 32 {
			return &__obf_466737d69e243d07{}
		}
		return &__obf_cc9e5c5a9ce7e2d1{}
	case reflect.Uint64:
		if __obf_0d9188c359cc43b1 != "uint64" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_cc9e5c5a9ce7e2d1{}
	case reflect.Float32:
		if __obf_0d9188c359cc43b1 != "float32" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_267f314242d45e83{}
	case reflect.Float64:
		if __obf_0d9188c359cc43b1 != "float64" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_9faca152bf30e4b0{}
	case reflect.Bool:
		if __obf_0d9188c359cc43b1 != "bool" {
			return __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_d0a75414375dd73f{}
	}
	return nil
}

type __obf_6e64bbe5dcd439c3 struct {
}

func (__obf_0099bc81efe356aa *__obf_6e64bbe5dcd439c3) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	*((*string)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadString()
}

func (__obf_0099bc81efe356aa *__obf_6e64bbe5dcd439c3) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_12c21b79fa86dcba := *((*string)(__obf_d3c919547bf7e47a))
	__obf_00fc491caa967a74.
		WriteString(__obf_12c21b79fa86dcba)
}

func (__obf_0099bc81efe356aa *__obf_6e64bbe5dcd439c3) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*string)(__obf_d3c919547bf7e47a)) == ""
}

type __obf_31bb730b2e917d9d struct {
}

func (__obf_0099bc81efe356aa *__obf_31bb730b2e917d9d) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*int8)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadInt8()
	}
}

func (__obf_0099bc81efe356aa *__obf_31bb730b2e917d9d) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteInt8(*((*int8)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_31bb730b2e917d9d) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*int8)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_5c68bf6a9fce0257 struct {
}

func (__obf_0099bc81efe356aa *__obf_5c68bf6a9fce0257) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*int16)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadInt16()
	}
}

func (__obf_0099bc81efe356aa *__obf_5c68bf6a9fce0257) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteInt16(*((*int16)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_5c68bf6a9fce0257) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*int16)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_aeb5e7cae1d3d27e struct {
}

func (__obf_0099bc81efe356aa *__obf_aeb5e7cae1d3d27e) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*int32)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadInt32()
	}
}

func (__obf_0099bc81efe356aa *__obf_aeb5e7cae1d3d27e) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteInt32(*((*int32)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_aeb5e7cae1d3d27e) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*int32)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_92c464f4f16e071c struct {
}

func (__obf_0099bc81efe356aa *__obf_92c464f4f16e071c) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*int64)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadInt64()
	}
}

func (__obf_0099bc81efe356aa *__obf_92c464f4f16e071c) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteInt64(*((*int64)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_92c464f4f16e071c) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*int64)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_72de03bc5292cee4 struct {
}

func (__obf_0099bc81efe356aa *__obf_72de03bc5292cee4) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*uint8)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadUint8()
	}
}

func (__obf_0099bc81efe356aa *__obf_72de03bc5292cee4) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteUint8(*((*uint8)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_72de03bc5292cee4) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*uint8)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_6d8a7372e7b2174f struct {
}

func (__obf_0099bc81efe356aa *__obf_6d8a7372e7b2174f) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*uint16)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadUint16()
	}
}

func (__obf_0099bc81efe356aa *__obf_6d8a7372e7b2174f) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteUint16(*((*uint16)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_6d8a7372e7b2174f) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*uint16)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_466737d69e243d07 struct {
}

func (__obf_0099bc81efe356aa *__obf_466737d69e243d07) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*uint32)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadUint32()
	}
}

func (__obf_0099bc81efe356aa *__obf_466737d69e243d07) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteUint32(*((*uint32)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_466737d69e243d07) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*uint32)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_cc9e5c5a9ce7e2d1 struct {
}

func (__obf_0099bc81efe356aa *__obf_cc9e5c5a9ce7e2d1) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*uint64)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadUint64()
	}
}

func (__obf_0099bc81efe356aa *__obf_cc9e5c5a9ce7e2d1) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteUint64(*((*uint64)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_cc9e5c5a9ce7e2d1) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*uint64)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_267f314242d45e83 struct {
}

func (__obf_0099bc81efe356aa *__obf_267f314242d45e83) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*float32)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadFloat32()
	}
}

func (__obf_0099bc81efe356aa *__obf_267f314242d45e83) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteFloat32(*((*float32)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_267f314242d45e83) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*float32)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_9faca152bf30e4b0 struct {
}

func (__obf_0099bc81efe356aa *__obf_9faca152bf30e4b0) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*float64)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadFloat64()
	}
}

func (__obf_0099bc81efe356aa *__obf_9faca152bf30e4b0) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteFloat64(*((*float64)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_9faca152bf30e4b0) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*float64)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_d0a75414375dd73f struct {
}

func (__obf_0099bc81efe356aa *__obf_d0a75414375dd73f) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if !__obf_67008a6a9e5ba828.ReadNil() {
		*((*bool)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.ReadBool()
	}
}

func (__obf_0099bc81efe356aa *__obf_d0a75414375dd73f) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteBool(*((*bool)(__obf_d3c919547bf7e47a)))
}

func (__obf_0099bc81efe356aa *__obf_d0a75414375dd73f) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return !(*((*bool)(__obf_d3c919547bf7e47a)))
}

type __obf_a18e9d92ad8c34a9 struct {
	__obf_2aefb0cc09305308 *reflect2.UnsafeSliceType
	__obf_fae02aecb00c25f9 ValDecoder
}

func (__obf_0099bc81efe356aa *__obf_a18e9d92ad8c34a9) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.ReadNil() {
		__obf_0099bc81efe356aa.__obf_2aefb0cc09305308.
			UnsafeSetNil(__obf_d3c919547bf7e47a)
		return
	}
	switch __obf_67008a6a9e5ba828.WhatIsNext() {
	case StringValue:
		__obf_9248e56f8d6040e1 := __obf_67008a6a9e5ba828.ReadString()
		__obf_d271a3a761283691, __obf_6d071d48f3b321ad := base64.StdEncoding.DecodeString(__obf_9248e56f8d6040e1)
		if __obf_6d071d48f3b321ad != nil {
			__obf_67008a6a9e5ba828.
				ReportError("decode base64", __obf_6d071d48f3b321ad.Error())
		} else {
			__obf_0099bc81efe356aa.__obf_2aefb0cc09305308.
				UnsafeSet(__obf_d3c919547bf7e47a, unsafe.Pointer(&__obf_d271a3a761283691))
		}
	case ArrayValue:
		__obf_0099bc81efe356aa.__obf_fae02aecb00c25f9.
			Decode(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
	default:
		__obf_67008a6a9e5ba828.
			ReportError("base64Codec", "invalid input")
	}
}

func (__obf_0099bc81efe356aa *__obf_a18e9d92ad8c34a9) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if __obf_0099bc81efe356aa.__obf_2aefb0cc09305308.UnsafeIsNil(__obf_d3c919547bf7e47a) {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	__obf_9248e56f8d6040e1 := *((*[]byte)(__obf_d3c919547bf7e47a))
	__obf_24f4a0ee87394bbd := base64.StdEncoding
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
	if len(__obf_9248e56f8d6040e1) != 0 {
		__obf_aedccbde4481459c := __obf_24f4a0ee87394bbd.EncodedLen(len(__obf_9248e56f8d6040e1))
		__obf_9fc06d9180f0daca := make([]byte, __obf_aedccbde4481459c)
		__obf_24f4a0ee87394bbd.
			Encode(__obf_9fc06d9180f0daca, __obf_9248e56f8d6040e1)
		__obf_00fc491caa967a74.__obf_9fc06d9180f0daca = append(__obf_00fc491caa967a74.__obf_9fc06d9180f0daca, __obf_9fc06d9180f0daca...)
	}
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
}

func (__obf_0099bc81efe356aa *__obf_a18e9d92ad8c34a9) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_d3c919547bf7e47a))) == 0
}
