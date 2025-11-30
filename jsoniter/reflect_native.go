package __obf_c3cd04a15c56f32f

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

const __obf_26c0da34d2f86b82 = 32 << uintptr(^uintptr(0)>>63)

func __obf_3828217e60017848(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	if __obf_8667d4fc2a95b0d7.Kind() == reflect.Slice && __obf_8667d4fc2a95b0d7.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_c58a03309b11d437 := __obf_f7de3e7720df8df1(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
		return &__obf_c2e8b967bdcef95c{__obf_c58a03309b11d437: __obf_c58a03309b11d437}
	}
	__obf_9c9c915bb768f418 := __obf_8667d4fc2a95b0d7.String()
	__obf_0e6540508f0e8fa6 := __obf_8667d4fc2a95b0d7.Kind()
	switch __obf_0e6540508f0e8fa6 {
	case reflect.String:
		if __obf_9c9c915bb768f418 != "string" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_dc7b14d60550decb{}
	case reflect.Int:
		if __obf_9c9c915bb768f418 != "int" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_98e7a4395b43ed9c{}
		}
		return &__obf_55c91da78a533634{}
	case reflect.Int8:
		if __obf_9c9c915bb768f418 != "int8" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_f5b2ca5d7a8d04e5{}
	case reflect.Int16:
		if __obf_9c9c915bb768f418 != "int16" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_3e034395bc5a3361{}
	case reflect.Int32:
		if __obf_9c9c915bb768f418 != "int32" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_98e7a4395b43ed9c{}
	case reflect.Int64:
		if __obf_9c9c915bb768f418 != "int64" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_55c91da78a533634{}
	case reflect.Uint:
		if __obf_9c9c915bb768f418 != "uint" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_9c7fdde19acfed0b{}
		}
		return &__obf_91e1d00fca72da55{}
	case reflect.Uint8:
		if __obf_9c9c915bb768f418 != "uint8" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_7573621aea0825be{}
	case reflect.Uint16:
		if __obf_9c9c915bb768f418 != "uint16" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_5aedf2d833b33bc7{}
	case reflect.Uint32:
		if __obf_9c9c915bb768f418 != "uint32" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_9c7fdde19acfed0b{}
	case reflect.Uintptr:
		if __obf_9c9c915bb768f418 != "uintptr" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_26c0da34d2f86b82 == 32 {
			return &__obf_9c7fdde19acfed0b{}
		}
		return &__obf_91e1d00fca72da55{}
	case reflect.Uint64:
		if __obf_9c9c915bb768f418 != "uint64" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_91e1d00fca72da55{}
	case reflect.Float32:
		if __obf_9c9c915bb768f418 != "float32" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_895882821abf4d80{}
	case reflect.Float64:
		if __obf_9c9c915bb768f418 != "float64" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_8ab18ccbae1c975f{}
	case reflect.Bool:
		if __obf_9c9c915bb768f418 != "bool" {
			return __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_bd766e3e1b2d90c4{}
	}
	return nil
}

func __obf_985b053eae9c9342(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	if __obf_8667d4fc2a95b0d7.Kind() == reflect.Slice && __obf_8667d4fc2a95b0d7.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_c58a03309b11d437 := __obf_f7de3e7720df8df1(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
		return &__obf_c2e8b967bdcef95c{__obf_c58a03309b11d437: __obf_c58a03309b11d437}
	}
	__obf_9c9c915bb768f418 := __obf_8667d4fc2a95b0d7.String()
	switch __obf_8667d4fc2a95b0d7.Kind() {
	case reflect.String:
		if __obf_9c9c915bb768f418 != "string" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_dc7b14d60550decb{}
	case reflect.Int:
		if __obf_9c9c915bb768f418 != "int" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_98e7a4395b43ed9c{}
		}
		return &__obf_55c91da78a533634{}
	case reflect.Int8:
		if __obf_9c9c915bb768f418 != "int8" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_f5b2ca5d7a8d04e5{}
	case reflect.Int16:
		if __obf_9c9c915bb768f418 != "int16" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_3e034395bc5a3361{}
	case reflect.Int32:
		if __obf_9c9c915bb768f418 != "int32" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_98e7a4395b43ed9c{}
	case reflect.Int64:
		if __obf_9c9c915bb768f418 != "int64" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_55c91da78a533634{}
	case reflect.Uint:
		if __obf_9c9c915bb768f418 != "uint" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_9c7fdde19acfed0b{}
		}
		return &__obf_91e1d00fca72da55{}
	case reflect.Uint8:
		if __obf_9c9c915bb768f418 != "uint8" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_7573621aea0825be{}
	case reflect.Uint16:
		if __obf_9c9c915bb768f418 != "uint16" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_5aedf2d833b33bc7{}
	case reflect.Uint32:
		if __obf_9c9c915bb768f418 != "uint32" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_9c7fdde19acfed0b{}
	case reflect.Uintptr:
		if __obf_9c9c915bb768f418 != "uintptr" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_26c0da34d2f86b82 == 32 {
			return &__obf_9c7fdde19acfed0b{}
		}
		return &__obf_91e1d00fca72da55{}
	case reflect.Uint64:
		if __obf_9c9c915bb768f418 != "uint64" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_91e1d00fca72da55{}
	case reflect.Float32:
		if __obf_9c9c915bb768f418 != "float32" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_895882821abf4d80{}
	case reflect.Float64:
		if __obf_9c9c915bb768f418 != "float64" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_8ab18ccbae1c975f{}
	case reflect.Bool:
		if __obf_9c9c915bb768f418 != "bool" {
			return __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_bd766e3e1b2d90c4{}
	}
	return nil
}

type __obf_dc7b14d60550decb struct {
}

func (__obf_1569d098873d2228 *__obf_dc7b14d60550decb) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	*((*string)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadString()
}

func (__obf_1569d098873d2228 *__obf_dc7b14d60550decb) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2d944450d48e5810 := *((*string)(__obf_753ab3fb72cdd659))
	__obf_2361f5214e84c60f.
		WriteString(__obf_2d944450d48e5810)
}

func (__obf_1569d098873d2228 *__obf_dc7b14d60550decb) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*string)(__obf_753ab3fb72cdd659)) == ""
}

type __obf_f5b2ca5d7a8d04e5 struct {
}

func (__obf_1569d098873d2228 *__obf_f5b2ca5d7a8d04e5) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*int8)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadInt8()
	}
}

func (__obf_1569d098873d2228 *__obf_f5b2ca5d7a8d04e5) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteInt8(*((*int8)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_f5b2ca5d7a8d04e5) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*int8)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_3e034395bc5a3361 struct {
}

func (__obf_1569d098873d2228 *__obf_3e034395bc5a3361) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*int16)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadInt16()
	}
}

func (__obf_1569d098873d2228 *__obf_3e034395bc5a3361) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteInt16(*((*int16)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_3e034395bc5a3361) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*int16)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_98e7a4395b43ed9c struct {
}

func (__obf_1569d098873d2228 *__obf_98e7a4395b43ed9c) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*int32)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadInt32()
	}
}

func (__obf_1569d098873d2228 *__obf_98e7a4395b43ed9c) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteInt32(*((*int32)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_98e7a4395b43ed9c) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*int32)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_55c91da78a533634 struct {
}

func (__obf_1569d098873d2228 *__obf_55c91da78a533634) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*int64)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadInt64()
	}
}

func (__obf_1569d098873d2228 *__obf_55c91da78a533634) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteInt64(*((*int64)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_55c91da78a533634) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*int64)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_7573621aea0825be struct {
}

func (__obf_1569d098873d2228 *__obf_7573621aea0825be) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*uint8)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadUint8()
	}
}

func (__obf_1569d098873d2228 *__obf_7573621aea0825be) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteUint8(*((*uint8)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_7573621aea0825be) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*uint8)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_5aedf2d833b33bc7 struct {
}

func (__obf_1569d098873d2228 *__obf_5aedf2d833b33bc7) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*uint16)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadUint16()
	}
}

func (__obf_1569d098873d2228 *__obf_5aedf2d833b33bc7) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteUint16(*((*uint16)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_5aedf2d833b33bc7) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*uint16)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_9c7fdde19acfed0b struct {
}

func (__obf_1569d098873d2228 *__obf_9c7fdde19acfed0b) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*uint32)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadUint32()
	}
}

func (__obf_1569d098873d2228 *__obf_9c7fdde19acfed0b) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteUint32(*((*uint32)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_9c7fdde19acfed0b) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*uint32)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_91e1d00fca72da55 struct {
}

func (__obf_1569d098873d2228 *__obf_91e1d00fca72da55) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*uint64)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadUint64()
	}
}

func (__obf_1569d098873d2228 *__obf_91e1d00fca72da55) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteUint64(*((*uint64)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_91e1d00fca72da55) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*uint64)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_895882821abf4d80 struct {
}

func (__obf_1569d098873d2228 *__obf_895882821abf4d80) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*float32)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadFloat32()
	}
}

func (__obf_1569d098873d2228 *__obf_895882821abf4d80) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteFloat32(*((*float32)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_895882821abf4d80) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*float32)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_8ab18ccbae1c975f struct {
}

func (__obf_1569d098873d2228 *__obf_8ab18ccbae1c975f) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*float64)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadFloat64()
	}
}

func (__obf_1569d098873d2228 *__obf_8ab18ccbae1c975f) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteFloat64(*((*float64)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_8ab18ccbae1c975f) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*float64)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_bd766e3e1b2d90c4 struct {
}

func (__obf_1569d098873d2228 *__obf_bd766e3e1b2d90c4) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if !__obf_edd9fe48d39445e4.ReadNil() {
		*((*bool)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.ReadBool()
	}
}

func (__obf_1569d098873d2228 *__obf_bd766e3e1b2d90c4) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteBool(*((*bool)(__obf_753ab3fb72cdd659)))
}

func (__obf_1569d098873d2228 *__obf_bd766e3e1b2d90c4) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return !(*((*bool)(__obf_753ab3fb72cdd659)))
}

type __obf_c2e8b967bdcef95c struct {
	__obf_ada810583e2cffde *reflect2.UnsafeSliceType
	__obf_c58a03309b11d437 ValDecoder
}

func (__obf_1569d098873d2228 *__obf_c2e8b967bdcef95c) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.ReadNil() {
		__obf_1569d098873d2228.__obf_ada810583e2cffde.
			UnsafeSetNil(__obf_753ab3fb72cdd659)
		return
	}
	switch __obf_edd9fe48d39445e4.WhatIsNext() {
	case StringValue:
		__obf_40d9a6116b8d8e66 := __obf_edd9fe48d39445e4.ReadString()
		__obf_19c217aabaacbf31, __obf_5966ecc5edb9b17e := base64.StdEncoding.DecodeString(__obf_40d9a6116b8d8e66)
		if __obf_5966ecc5edb9b17e != nil {
			__obf_edd9fe48d39445e4.
				ReportError("decode base64", __obf_5966ecc5edb9b17e.Error())
		} else {
			__obf_1569d098873d2228.__obf_ada810583e2cffde.
				UnsafeSet(__obf_753ab3fb72cdd659, unsafe.Pointer(&__obf_19c217aabaacbf31))
		}
	case ArrayValue:
		__obf_1569d098873d2228.__obf_c58a03309b11d437.
			Decode(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
	default:
		__obf_edd9fe48d39445e4.
			ReportError("base64Codec", "invalid input")
	}
}

func (__obf_1569d098873d2228 *__obf_c2e8b967bdcef95c) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	if __obf_1569d098873d2228.__obf_ada810583e2cffde.UnsafeIsNil(__obf_753ab3fb72cdd659) {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	__obf_40d9a6116b8d8e66 := *((*[]byte)(__obf_753ab3fb72cdd659))
	__obf_a65e1eaf4e40ede4 := base64.StdEncoding
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
	if len(__obf_40d9a6116b8d8e66) != 0 {
		__obf_ecf95be2d6e27166 := __obf_a65e1eaf4e40ede4.EncodedLen(len(__obf_40d9a6116b8d8e66))
		__obf_ace979f6622823f3 := make([]byte, __obf_ecf95be2d6e27166)
		__obf_a65e1eaf4e40ede4.
			Encode(__obf_ace979f6622823f3, __obf_40d9a6116b8d8e66)
		__obf_2361f5214e84c60f.__obf_ace979f6622823f3 = append(__obf_2361f5214e84c60f.__obf_ace979f6622823f3, __obf_ace979f6622823f3...)
	}
	__obf_2361f5214e84c60f.__obf_c4fec0edfb3875ad('"')
}

func (__obf_1569d098873d2228 *__obf_c2e8b967bdcef95c) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_753ab3fb72cdd659))) == 0
}
