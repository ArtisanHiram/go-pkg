package __obf_703d23ba520c3295

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

const __obf_6927eb49141afe0a = 32 << uintptr(^uintptr(0)>>63)

func __obf_3774fbf920f93937(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	if __obf_3b7f6abbae19451e.Kind() == reflect.Slice && __obf_3b7f6abbae19451e.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_b8463ec8d4aaf2be := __obf_aeff82b222c8ff5b(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
		return &__obf_d5e42f76266a8c66{__obf_b8463ec8d4aaf2be: __obf_b8463ec8d4aaf2be}
	}
	__obf_48b633681f0b5b99 := __obf_3b7f6abbae19451e.String()
	__obf_2bb4ea546eb4165c := __obf_3b7f6abbae19451e.Kind()
	switch __obf_2bb4ea546eb4165c {
	case reflect.String:
		if __obf_48b633681f0b5b99 != "string" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_b4b1ce99589eef7a{}
	case reflect.Int:
		if __obf_48b633681f0b5b99 != "int" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_35257b5680d92c49{}
		}
		return &__obf_fd2003dbeca94a10{}
	case reflect.Int8:
		if __obf_48b633681f0b5b99 != "int8" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_c699d8dd0f200d7b{}
	case reflect.Int16:
		if __obf_48b633681f0b5b99 != "int16" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_be52e794d1070cb5{}
	case reflect.Int32:
		if __obf_48b633681f0b5b99 != "int32" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_35257b5680d92c49{}
	case reflect.Int64:
		if __obf_48b633681f0b5b99 != "int64" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_fd2003dbeca94a10{}
	case reflect.Uint:
		if __obf_48b633681f0b5b99 != "uint" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_8ffbfcff07b962c4{}
		}
		return &__obf_8373e5cb39284d85{}
	case reflect.Uint8:
		if __obf_48b633681f0b5b99 != "uint8" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_72e9a7eb9be3840a{}
	case reflect.Uint16:
		if __obf_48b633681f0b5b99 != "uint16" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_f6bd81c248d21217{}
	case reflect.Uint32:
		if __obf_48b633681f0b5b99 != "uint32" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_8ffbfcff07b962c4{}
	case reflect.Uintptr:
		if __obf_48b633681f0b5b99 != "uintptr" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_6927eb49141afe0a == 32 {
			return &__obf_8ffbfcff07b962c4{}
		}
		return &__obf_8373e5cb39284d85{}
	case reflect.Uint64:
		if __obf_48b633681f0b5b99 != "uint64" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_8373e5cb39284d85{}
	case reflect.Float32:
		if __obf_48b633681f0b5b99 != "float32" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_7f1fdc3a9a639106{}
	case reflect.Float64:
		if __obf_48b633681f0b5b99 != "float64" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_40c5d6ad61432000{}
	case reflect.Bool:
		if __obf_48b633681f0b5b99 != "bool" {
			return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_b2791a8f9a46f9b8{}
	}
	return nil
}

func __obf_d310ff63d30070e1(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	if __obf_3b7f6abbae19451e.Kind() == reflect.Slice && __obf_3b7f6abbae19451e.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_b8463ec8d4aaf2be := __obf_aeff82b222c8ff5b(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
		return &__obf_d5e42f76266a8c66{__obf_b8463ec8d4aaf2be: __obf_b8463ec8d4aaf2be}
	}
	__obf_48b633681f0b5b99 := __obf_3b7f6abbae19451e.String()
	switch __obf_3b7f6abbae19451e.Kind() {
	case reflect.String:
		if __obf_48b633681f0b5b99 != "string" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_b4b1ce99589eef7a{}
	case reflect.Int:
		if __obf_48b633681f0b5b99 != "int" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_35257b5680d92c49{}
		}
		return &__obf_fd2003dbeca94a10{}
	case reflect.Int8:
		if __obf_48b633681f0b5b99 != "int8" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_c699d8dd0f200d7b{}
	case reflect.Int16:
		if __obf_48b633681f0b5b99 != "int16" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_be52e794d1070cb5{}
	case reflect.Int32:
		if __obf_48b633681f0b5b99 != "int32" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_35257b5680d92c49{}
	case reflect.Int64:
		if __obf_48b633681f0b5b99 != "int64" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_fd2003dbeca94a10{}
	case reflect.Uint:
		if __obf_48b633681f0b5b99 != "uint" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_8ffbfcff07b962c4{}
		}
		return &__obf_8373e5cb39284d85{}
	case reflect.Uint8:
		if __obf_48b633681f0b5b99 != "uint8" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_72e9a7eb9be3840a{}
	case reflect.Uint16:
		if __obf_48b633681f0b5b99 != "uint16" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_f6bd81c248d21217{}
	case reflect.Uint32:
		if __obf_48b633681f0b5b99 != "uint32" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_8ffbfcff07b962c4{}
	case reflect.Uintptr:
		if __obf_48b633681f0b5b99 != "uintptr" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_6927eb49141afe0a == 32 {
			return &__obf_8ffbfcff07b962c4{}
		}
		return &__obf_8373e5cb39284d85{}
	case reflect.Uint64:
		if __obf_48b633681f0b5b99 != "uint64" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_8373e5cb39284d85{}
	case reflect.Float32:
		if __obf_48b633681f0b5b99 != "float32" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_7f1fdc3a9a639106{}
	case reflect.Float64:
		if __obf_48b633681f0b5b99 != "float64" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_40c5d6ad61432000{}
	case reflect.Bool:
		if __obf_48b633681f0b5b99 != "bool" {
			return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_b2791a8f9a46f9b8{}
	}
	return nil
}

type __obf_b4b1ce99589eef7a struct {
}

func (__obf_4acad06eb5535907 *__obf_b4b1ce99589eef7a) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	*((*string)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadString()
}

func (__obf_4acad06eb5535907 *__obf_b4b1ce99589eef7a) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_44b48c26051f8033 := *((*string)(__obf_47fa53f3d161f45c))
	__obf_9a34f62857fb1d1d.
		WriteString(__obf_44b48c26051f8033)
}

func (__obf_4acad06eb5535907 *__obf_b4b1ce99589eef7a) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*string)(__obf_47fa53f3d161f45c)) == ""
}

type __obf_c699d8dd0f200d7b struct {
}

func (__obf_4acad06eb5535907 *__obf_c699d8dd0f200d7b) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*int8)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadInt8()
	}
}

func (__obf_4acad06eb5535907 *__obf_c699d8dd0f200d7b) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteInt8(*((*int8)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_c699d8dd0f200d7b) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*int8)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_be52e794d1070cb5 struct {
}

func (__obf_4acad06eb5535907 *__obf_be52e794d1070cb5) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*int16)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadInt16()
	}
}

func (__obf_4acad06eb5535907 *__obf_be52e794d1070cb5) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteInt16(*((*int16)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_be52e794d1070cb5) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*int16)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_35257b5680d92c49 struct {
}

func (__obf_4acad06eb5535907 *__obf_35257b5680d92c49) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*int32)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadInt32()
	}
}

func (__obf_4acad06eb5535907 *__obf_35257b5680d92c49) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteInt32(*((*int32)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_35257b5680d92c49) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*int32)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_fd2003dbeca94a10 struct {
}

func (__obf_4acad06eb5535907 *__obf_fd2003dbeca94a10) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*int64)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadInt64()
	}
}

func (__obf_4acad06eb5535907 *__obf_fd2003dbeca94a10) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteInt64(*((*int64)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_fd2003dbeca94a10) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*int64)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_72e9a7eb9be3840a struct {
}

func (__obf_4acad06eb5535907 *__obf_72e9a7eb9be3840a) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*uint8)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadUint8()
	}
}

func (__obf_4acad06eb5535907 *__obf_72e9a7eb9be3840a) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteUint8(*((*uint8)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_72e9a7eb9be3840a) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*uint8)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_f6bd81c248d21217 struct {
}

func (__obf_4acad06eb5535907 *__obf_f6bd81c248d21217) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*uint16)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadUint16()
	}
}

func (__obf_4acad06eb5535907 *__obf_f6bd81c248d21217) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteUint16(*((*uint16)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_f6bd81c248d21217) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*uint16)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_8ffbfcff07b962c4 struct {
}

func (__obf_4acad06eb5535907 *__obf_8ffbfcff07b962c4) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*uint32)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadUint32()
	}
}

func (__obf_4acad06eb5535907 *__obf_8ffbfcff07b962c4) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteUint32(*((*uint32)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_8ffbfcff07b962c4) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*uint32)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_8373e5cb39284d85 struct {
}

func (__obf_4acad06eb5535907 *__obf_8373e5cb39284d85) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*uint64)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadUint64()
	}
}

func (__obf_4acad06eb5535907 *__obf_8373e5cb39284d85) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteUint64(*((*uint64)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_8373e5cb39284d85) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*uint64)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_7f1fdc3a9a639106 struct {
}

func (__obf_4acad06eb5535907 *__obf_7f1fdc3a9a639106) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*float32)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadFloat32()
	}
}

func (__obf_4acad06eb5535907 *__obf_7f1fdc3a9a639106) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteFloat32(*((*float32)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_7f1fdc3a9a639106) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*float32)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_40c5d6ad61432000 struct {
}

func (__obf_4acad06eb5535907 *__obf_40c5d6ad61432000) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*float64)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadFloat64()
	}
}

func (__obf_4acad06eb5535907 *__obf_40c5d6ad61432000) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteFloat64(*((*float64)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_40c5d6ad61432000) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*float64)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_b2791a8f9a46f9b8 struct {
}

func (__obf_4acad06eb5535907 *__obf_b2791a8f9a46f9b8) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if !__obf_47edab4c16a2d88d.ReadNil() {
		*((*bool)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.ReadBool()
	}
}

func (__obf_4acad06eb5535907 *__obf_b2791a8f9a46f9b8) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteBool(*((*bool)(__obf_47fa53f3d161f45c)))
}

func (__obf_4acad06eb5535907 *__obf_b2791a8f9a46f9b8) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return !(*((*bool)(__obf_47fa53f3d161f45c)))
}

type __obf_d5e42f76266a8c66 struct {
	__obf_fdb50398a2b1e6b1 *reflect2.UnsafeSliceType
	__obf_b8463ec8d4aaf2be ValDecoder
}

func (__obf_4acad06eb5535907 *__obf_d5e42f76266a8c66) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.ReadNil() {
		__obf_4acad06eb5535907.__obf_fdb50398a2b1e6b1.
			UnsafeSetNil(__obf_47fa53f3d161f45c)
		return
	}
	switch __obf_47edab4c16a2d88d.WhatIsNext() {
	case StringValue:
		__obf_ed4ff8130f1b3f3b := __obf_47edab4c16a2d88d.ReadString()
		__obf_aef3bfdcada2e44a, __obf_e56742d6e52f642e := base64.StdEncoding.DecodeString(__obf_ed4ff8130f1b3f3b)
		if __obf_e56742d6e52f642e != nil {
			__obf_47edab4c16a2d88d.
				ReportError("decode base64", __obf_e56742d6e52f642e.Error())
		} else {
			__obf_4acad06eb5535907.__obf_fdb50398a2b1e6b1.
				UnsafeSet(__obf_47fa53f3d161f45c, unsafe.Pointer(&__obf_aef3bfdcada2e44a))
		}
	case ArrayValue:
		__obf_4acad06eb5535907.__obf_b8463ec8d4aaf2be.
			Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	default:
		__obf_47edab4c16a2d88d.
			ReportError("base64Codec", "invalid input")
	}
}

func (__obf_4acad06eb5535907 *__obf_d5e42f76266a8c66) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if __obf_4acad06eb5535907.__obf_fdb50398a2b1e6b1.UnsafeIsNil(__obf_47fa53f3d161f45c) {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_ed4ff8130f1b3f3b := *((*[]byte)(__obf_47fa53f3d161f45c))
	__obf_fe1f073939e1f7cc := base64.StdEncoding
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
	if len(__obf_ed4ff8130f1b3f3b) != 0 {
		__obf_0126ec6b3c37befb := __obf_fe1f073939e1f7cc.EncodedLen(len(__obf_ed4ff8130f1b3f3b))
		__obf_a065f8e0da5f5952 := make([]byte, __obf_0126ec6b3c37befb)
		__obf_fe1f073939e1f7cc.
			Encode(__obf_a065f8e0da5f5952, __obf_ed4ff8130f1b3f3b)
		__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952 = append(__obf_9a34f62857fb1d1d.__obf_a065f8e0da5f5952, __obf_a065f8e0da5f5952...)
	}
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
}

func (__obf_4acad06eb5535907 *__obf_d5e42f76266a8c66) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_47fa53f3d161f45c))) == 0
}
