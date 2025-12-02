package __obf_c7b79b12b33d8238

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

const __obf_95a95e14d0f9c485 = 32 << uintptr(^uintptr(0)>>63)

func __obf_1f22c50e61ed6284(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	if __obf_edcded11af6ebc4c.Kind() == reflect.Slice && __obf_edcded11af6ebc4c.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_06c2733d7b043ec8 := __obf_978533a1047109a9(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
		return &__obf_31a7155b9ae7865d{__obf_06c2733d7b043ec8: __obf_06c2733d7b043ec8}
	}
	__obf_391099e75d7a5ccb := __obf_edcded11af6ebc4c.String()
	__obf_105d1cd1a59219cd := __obf_edcded11af6ebc4c.Kind()
	switch __obf_105d1cd1a59219cd {
	case reflect.String:
		if __obf_391099e75d7a5ccb != "string" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_589f8c68e65ce95b{}
	case reflect.Int:
		if __obf_391099e75d7a5ccb != "int" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_1ec130220d85a877{}
		}
		return &__obf_a8184120ec91c9ac{}
	case reflect.Int8:
		if __obf_391099e75d7a5ccb != "int8" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_7cba7f0525204214{}
	case reflect.Int16:
		if __obf_391099e75d7a5ccb != "int16" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_18df0482f03469ee{}
	case reflect.Int32:
		if __obf_391099e75d7a5ccb != "int32" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_1ec130220d85a877{}
	case reflect.Int64:
		if __obf_391099e75d7a5ccb != "int64" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_a8184120ec91c9ac{}
	case reflect.Uint:
		if __obf_391099e75d7a5ccb != "uint" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_c1780610c9f021bc{}
		}
		return &__obf_92526ab5b54d1c2b{}
	case reflect.Uint8:
		if __obf_391099e75d7a5ccb != "uint8" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_2b83f334903ad8fc{}
	case reflect.Uint16:
		if __obf_391099e75d7a5ccb != "uint16" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_f46b5561244cda91{}
	case reflect.Uint32:
		if __obf_391099e75d7a5ccb != "uint32" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_c1780610c9f021bc{}
	case reflect.Uintptr:
		if __obf_391099e75d7a5ccb != "uintptr" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_95a95e14d0f9c485 == 32 {
			return &__obf_c1780610c9f021bc{}
		}
		return &__obf_92526ab5b54d1c2b{}
	case reflect.Uint64:
		if __obf_391099e75d7a5ccb != "uint64" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_92526ab5b54d1c2b{}
	case reflect.Float32:
		if __obf_391099e75d7a5ccb != "float32" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_7b48d0b55a17470d{}
	case reflect.Float64:
		if __obf_391099e75d7a5ccb != "float64" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_e2388a7e9751f5ac{}
	case reflect.Bool:
		if __obf_391099e75d7a5ccb != "bool" {
			return __obf_85f0e4c352da4678(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_919cca597f9e30df{}
	}
	return nil
}

func __obf_8cd5deb18d9f5f28(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	if __obf_edcded11af6ebc4c.Kind() == reflect.Slice && __obf_edcded11af6ebc4c.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_06c2733d7b043ec8 := __obf_978533a1047109a9(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
		return &__obf_31a7155b9ae7865d{__obf_06c2733d7b043ec8: __obf_06c2733d7b043ec8}
	}
	__obf_391099e75d7a5ccb := __obf_edcded11af6ebc4c.String()
	switch __obf_edcded11af6ebc4c.Kind() {
	case reflect.String:
		if __obf_391099e75d7a5ccb != "string" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_589f8c68e65ce95b{}
	case reflect.Int:
		if __obf_391099e75d7a5ccb != "int" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_1ec130220d85a877{}
		}
		return &__obf_a8184120ec91c9ac{}
	case reflect.Int8:
		if __obf_391099e75d7a5ccb != "int8" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_7cba7f0525204214{}
	case reflect.Int16:
		if __obf_391099e75d7a5ccb != "int16" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_18df0482f03469ee{}
	case reflect.Int32:
		if __obf_391099e75d7a5ccb != "int32" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_1ec130220d85a877{}
	case reflect.Int64:
		if __obf_391099e75d7a5ccb != "int64" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_a8184120ec91c9ac{}
	case reflect.Uint:
		if __obf_391099e75d7a5ccb != "uint" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_c1780610c9f021bc{}
		}
		return &__obf_92526ab5b54d1c2b{}
	case reflect.Uint8:
		if __obf_391099e75d7a5ccb != "uint8" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_2b83f334903ad8fc{}
	case reflect.Uint16:
		if __obf_391099e75d7a5ccb != "uint16" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_f46b5561244cda91{}
	case reflect.Uint32:
		if __obf_391099e75d7a5ccb != "uint32" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_c1780610c9f021bc{}
	case reflect.Uintptr:
		if __obf_391099e75d7a5ccb != "uintptr" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_95a95e14d0f9c485 == 32 {
			return &__obf_c1780610c9f021bc{}
		}
		return &__obf_92526ab5b54d1c2b{}
	case reflect.Uint64:
		if __obf_391099e75d7a5ccb != "uint64" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_92526ab5b54d1c2b{}
	case reflect.Float32:
		if __obf_391099e75d7a5ccb != "float32" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_7b48d0b55a17470d{}
	case reflect.Float64:
		if __obf_391099e75d7a5ccb != "float64" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_e2388a7e9751f5ac{}
	case reflect.Bool:
		if __obf_391099e75d7a5ccb != "bool" {
			return __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_919cca597f9e30df{}
	}
	return nil
}

type __obf_589f8c68e65ce95b struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_589f8c68e65ce95b) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	*((*string)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadString()
}

func (__obf_26e7ab1d7cef55d9 *__obf_589f8c68e65ce95b) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_a3eaafc22faf1644 := *((*string)(__obf_575c04f2b9d91315))
	__obf_d8f50bcfe87d8b47.
		WriteString(__obf_a3eaafc22faf1644)
}

func (__obf_26e7ab1d7cef55d9 *__obf_589f8c68e65ce95b) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*string)(__obf_575c04f2b9d91315)) == ""
}

type __obf_7cba7f0525204214 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_7cba7f0525204214) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*int8)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadInt8()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_7cba7f0525204214) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteInt8(*((*int8)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_7cba7f0525204214) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*int8)(__obf_575c04f2b9d91315)) == 0
}

type __obf_18df0482f03469ee struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_18df0482f03469ee) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*int16)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadInt16()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_18df0482f03469ee) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteInt16(*((*int16)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_18df0482f03469ee) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*int16)(__obf_575c04f2b9d91315)) == 0
}

type __obf_1ec130220d85a877 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_1ec130220d85a877) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*int32)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadInt32()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_1ec130220d85a877) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteInt32(*((*int32)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_1ec130220d85a877) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*int32)(__obf_575c04f2b9d91315)) == 0
}

type __obf_a8184120ec91c9ac struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_a8184120ec91c9ac) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*int64)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadInt64()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_a8184120ec91c9ac) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteInt64(*((*int64)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_a8184120ec91c9ac) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*int64)(__obf_575c04f2b9d91315)) == 0
}

type __obf_2b83f334903ad8fc struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_2b83f334903ad8fc) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*uint8)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadUint8()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_2b83f334903ad8fc) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteUint8(*((*uint8)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_2b83f334903ad8fc) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*uint8)(__obf_575c04f2b9d91315)) == 0
}

type __obf_f46b5561244cda91 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_f46b5561244cda91) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*uint16)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadUint16()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_f46b5561244cda91) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteUint16(*((*uint16)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_f46b5561244cda91) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*uint16)(__obf_575c04f2b9d91315)) == 0
}

type __obf_c1780610c9f021bc struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_c1780610c9f021bc) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*uint32)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadUint32()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_c1780610c9f021bc) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteUint32(*((*uint32)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_c1780610c9f021bc) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*uint32)(__obf_575c04f2b9d91315)) == 0
}

type __obf_92526ab5b54d1c2b struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_92526ab5b54d1c2b) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*uint64)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadUint64()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_92526ab5b54d1c2b) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteUint64(*((*uint64)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_92526ab5b54d1c2b) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*uint64)(__obf_575c04f2b9d91315)) == 0
}

type __obf_7b48d0b55a17470d struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_7b48d0b55a17470d) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*float32)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadFloat32()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_7b48d0b55a17470d) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteFloat32(*((*float32)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_7b48d0b55a17470d) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*float32)(__obf_575c04f2b9d91315)) == 0
}

type __obf_e2388a7e9751f5ac struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_e2388a7e9751f5ac) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*float64)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadFloat64()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_e2388a7e9751f5ac) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteFloat64(*((*float64)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_e2388a7e9751f5ac) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*float64)(__obf_575c04f2b9d91315)) == 0
}

type __obf_919cca597f9e30df struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_919cca597f9e30df) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if !__obf_0da8c843dad13139.ReadNil() {
		*((*bool)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.ReadBool()
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_919cca597f9e30df) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteBool(*((*bool)(__obf_575c04f2b9d91315)))
}

func (__obf_26e7ab1d7cef55d9 *__obf_919cca597f9e30df) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return !(*((*bool)(__obf_575c04f2b9d91315)))
}

type __obf_31a7155b9ae7865d struct {
	__obf_782666043ff6c7ac *reflect2.UnsafeSliceType
	__obf_06c2733d7b043ec8 ValDecoder
}

func (__obf_26e7ab1d7cef55d9 *__obf_31a7155b9ae7865d) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	if __obf_0da8c843dad13139.ReadNil() {
		__obf_26e7ab1d7cef55d9.__obf_782666043ff6c7ac.
			UnsafeSetNil(__obf_575c04f2b9d91315)
		return
	}
	switch __obf_0da8c843dad13139.WhatIsNext() {
	case StringValue:
		__obf_f23c37b8d0da0240 := __obf_0da8c843dad13139.ReadString()
		__obf_c9d27a0593bc18e5, __obf_ea0680f8b461a85b := base64.StdEncoding.DecodeString(__obf_f23c37b8d0da0240)
		if __obf_ea0680f8b461a85b != nil {
			__obf_0da8c843dad13139.
				ReportError("decode base64", __obf_ea0680f8b461a85b.Error())
		} else {
			__obf_26e7ab1d7cef55d9.__obf_782666043ff6c7ac.
				UnsafeSet(__obf_575c04f2b9d91315, unsafe.Pointer(&__obf_c9d27a0593bc18e5))
		}
	case ArrayValue:
		__obf_26e7ab1d7cef55d9.__obf_06c2733d7b043ec8.
			Decode(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
	default:
		__obf_0da8c843dad13139.
			ReportError("base64Codec", "invalid input")
	}
}

func (__obf_26e7ab1d7cef55d9 *__obf_31a7155b9ae7865d) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	if __obf_26e7ab1d7cef55d9.__obf_782666043ff6c7ac.UnsafeIsNil(__obf_575c04f2b9d91315) {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_f23c37b8d0da0240 := *((*[]byte)(__obf_575c04f2b9d91315))
	__obf_4f2ad245ff71fb8f := base64.StdEncoding
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
	if len(__obf_f23c37b8d0da0240) != 0 {
		__obf_86901472dc2ab8e1 := __obf_4f2ad245ff71fb8f.EncodedLen(len(__obf_f23c37b8d0da0240))
		__obf_684d738bc3ac851a := make([]byte, __obf_86901472dc2ab8e1)
		__obf_4f2ad245ff71fb8f.
			Encode(__obf_684d738bc3ac851a, __obf_f23c37b8d0da0240)
		__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a = append(__obf_d8f50bcfe87d8b47.__obf_684d738bc3ac851a, __obf_684d738bc3ac851a...)
	}
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
}

func (__obf_26e7ab1d7cef55d9 *__obf_31a7155b9ae7865d) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_575c04f2b9d91315))) == 0
}
