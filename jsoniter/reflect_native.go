package __obf_030d39f7a8de96c6

import (
	"encoding/base64"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

const __obf_eb8e424502869741 = 32 << uintptr(^uintptr(0)>>63)

func __obf_438eebf6331ef15b(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	if __obf_8744d0b8c80ccdc1.Kind() == reflect.Slice && __obf_8744d0b8c80ccdc1.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_b122f1d695f8466c := __obf_6f35214e10cf09dc(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
		return &__obf_cb3ab562e0bc750f{__obf_b122f1d695f8466c: __obf_b122f1d695f8466c}
	}
	__obf_aa66e764b9a50b6b := __obf_8744d0b8c80ccdc1.String()
	__obf_30e494a4f2832c2f := __obf_8744d0b8c80ccdc1.Kind()
	switch __obf_30e494a4f2832c2f {
	case reflect.String:
		if __obf_aa66e764b9a50b6b != "string" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_77fd4e98864958b8{}
	case reflect.Int:
		if __obf_aa66e764b9a50b6b != "int" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_e22f6605dc28fd15{}
		}
		return &__obf_f9521cbe2b8f862b{}
	case reflect.Int8:
		if __obf_aa66e764b9a50b6b != "int8" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_e4640ed90fc2b426{}
	case reflect.Int16:
		if __obf_aa66e764b9a50b6b != "int16" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_10bddab489f72ae1{}
	case reflect.Int32:
		if __obf_aa66e764b9a50b6b != "int32" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_e22f6605dc28fd15{}
	case reflect.Int64:
		if __obf_aa66e764b9a50b6b != "int64" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_f9521cbe2b8f862b{}
	case reflect.Uint:
		if __obf_aa66e764b9a50b6b != "uint" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_1dd89735fd333751{}
		}
		return &__obf_0b943b84d45dcd50{}
	case reflect.Uint8:
		if __obf_aa66e764b9a50b6b != "uint8" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_7bba42e561382f51{}
	case reflect.Uint16:
		if __obf_aa66e764b9a50b6b != "uint16" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_118a03214ee56049{}
	case reflect.Uint32:
		if __obf_aa66e764b9a50b6b != "uint32" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_1dd89735fd333751{}
	case reflect.Uintptr:
		if __obf_aa66e764b9a50b6b != "uintptr" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_eb8e424502869741 == 32 {
			return &__obf_1dd89735fd333751{}
		}
		return &__obf_0b943b84d45dcd50{}
	case reflect.Uint64:
		if __obf_aa66e764b9a50b6b != "uint64" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_0b943b84d45dcd50{}
	case reflect.Float32:
		if __obf_aa66e764b9a50b6b != "float32" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_dada070e88eaa43a{}
	case reflect.Float64:
		if __obf_aa66e764b9a50b6b != "float64" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_df18f1273f2fb63e{}
	case reflect.Bool:
		if __obf_aa66e764b9a50b6b != "bool" {
			return __obf_37ef471fa5e40404(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_8bf4655000719b09{}
	}
	return nil
}

func __obf_506acfeb1c549d6a(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	if __obf_8744d0b8c80ccdc1.Kind() == reflect.Slice && __obf_8744d0b8c80ccdc1.(reflect2.SliceType).Elem().Kind() == reflect.Uint8 {
		__obf_b122f1d695f8466c := __obf_6f35214e10cf09dc(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
		return &__obf_cb3ab562e0bc750f{__obf_b122f1d695f8466c: __obf_b122f1d695f8466c}
	}
	__obf_aa66e764b9a50b6b := __obf_8744d0b8c80ccdc1.String()
	switch __obf_8744d0b8c80ccdc1.Kind() {
	case reflect.String:
		if __obf_aa66e764b9a50b6b != "string" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*string)(nil)).Elem())
		}
		return &__obf_77fd4e98864958b8{}
	case reflect.Int:
		if __obf_aa66e764b9a50b6b != "int" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_e22f6605dc28fd15{}
		}
		return &__obf_f9521cbe2b8f862b{}
	case reflect.Int8:
		if __obf_aa66e764b9a50b6b != "int8" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int8)(nil)).Elem())
		}
		return &__obf_e4640ed90fc2b426{}
	case reflect.Int16:
		if __obf_aa66e764b9a50b6b != "int16" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int16)(nil)).Elem())
		}
		return &__obf_10bddab489f72ae1{}
	case reflect.Int32:
		if __obf_aa66e764b9a50b6b != "int32" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int32)(nil)).Elem())
		}
		return &__obf_e22f6605dc28fd15{}
	case reflect.Int64:
		if __obf_aa66e764b9a50b6b != "int64" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*int64)(nil)).Elem())
		}
		return &__obf_f9521cbe2b8f862b{}
	case reflect.Uint:
		if __obf_aa66e764b9a50b6b != "uint" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint)(nil)).Elem())
		}
		if strconv.IntSize == 32 {
			return &__obf_1dd89735fd333751{}
		}
		return &__obf_0b943b84d45dcd50{}
	case reflect.Uint8:
		if __obf_aa66e764b9a50b6b != "uint8" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint8)(nil)).Elem())
		}
		return &__obf_7bba42e561382f51{}
	case reflect.Uint16:
		if __obf_aa66e764b9a50b6b != "uint16" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint16)(nil)).Elem())
		}
		return &__obf_118a03214ee56049{}
	case reflect.Uint32:
		if __obf_aa66e764b9a50b6b != "uint32" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint32)(nil)).Elem())
		}
		return &__obf_1dd89735fd333751{}
	case reflect.Uintptr:
		if __obf_aa66e764b9a50b6b != "uintptr" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uintptr)(nil)).Elem())
		}
		if __obf_eb8e424502869741 == 32 {
			return &__obf_1dd89735fd333751{}
		}
		return &__obf_0b943b84d45dcd50{}
	case reflect.Uint64:
		if __obf_aa66e764b9a50b6b != "uint64" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*uint64)(nil)).Elem())
		}
		return &__obf_0b943b84d45dcd50{}
	case reflect.Float32:
		if __obf_aa66e764b9a50b6b != "float32" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*float32)(nil)).Elem())
		}
		return &__obf_dada070e88eaa43a{}
	case reflect.Float64:
		if __obf_aa66e764b9a50b6b != "float64" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*float64)(nil)).Elem())
		}
		return &__obf_df18f1273f2fb63e{}
	case reflect.Bool:
		if __obf_aa66e764b9a50b6b != "bool" {
			return __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0, reflect2.TypeOfPtr((*bool)(nil)).Elem())
		}
		return &__obf_8bf4655000719b09{}
	}
	return nil
}

type __obf_77fd4e98864958b8 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_77fd4e98864958b8) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	*((*string)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadString()
}

func (__obf_7bfaa2f467491ab9 *__obf_77fd4e98864958b8) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_428c3b4a95725c4a := *((*string)(__obf_dbbf371b91136494))
	__obf_8a2c51fe22775655.
		WriteString(__obf_428c3b4a95725c4a)
}

func (__obf_7bfaa2f467491ab9 *__obf_77fd4e98864958b8) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*string)(__obf_dbbf371b91136494)) == ""
}

type __obf_e4640ed90fc2b426 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_e4640ed90fc2b426) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*int8)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadInt8()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_e4640ed90fc2b426) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteInt8(*((*int8)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_e4640ed90fc2b426) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*int8)(__obf_dbbf371b91136494)) == 0
}

type __obf_10bddab489f72ae1 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_10bddab489f72ae1) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*int16)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadInt16()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_10bddab489f72ae1) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteInt16(*((*int16)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_10bddab489f72ae1) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*int16)(__obf_dbbf371b91136494)) == 0
}

type __obf_e22f6605dc28fd15 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_e22f6605dc28fd15) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*int32)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadInt32()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_e22f6605dc28fd15) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteInt32(*((*int32)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_e22f6605dc28fd15) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*int32)(__obf_dbbf371b91136494)) == 0
}

type __obf_f9521cbe2b8f862b struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_f9521cbe2b8f862b) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*int64)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadInt64()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_f9521cbe2b8f862b) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteInt64(*((*int64)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_f9521cbe2b8f862b) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*int64)(__obf_dbbf371b91136494)) == 0
}

type __obf_7bba42e561382f51 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_7bba42e561382f51) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*uint8)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadUint8()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_7bba42e561382f51) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteUint8(*((*uint8)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_7bba42e561382f51) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*uint8)(__obf_dbbf371b91136494)) == 0
}

type __obf_118a03214ee56049 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_118a03214ee56049) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*uint16)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadUint16()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_118a03214ee56049) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteUint16(*((*uint16)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_118a03214ee56049) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*uint16)(__obf_dbbf371b91136494)) == 0
}

type __obf_1dd89735fd333751 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_1dd89735fd333751) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*uint32)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadUint32()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_1dd89735fd333751) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteUint32(*((*uint32)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_1dd89735fd333751) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*uint32)(__obf_dbbf371b91136494)) == 0
}

type __obf_0b943b84d45dcd50 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_0b943b84d45dcd50) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*uint64)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadUint64()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_0b943b84d45dcd50) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteUint64(*((*uint64)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_0b943b84d45dcd50) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*uint64)(__obf_dbbf371b91136494)) == 0
}

type __obf_dada070e88eaa43a struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_dada070e88eaa43a) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*float32)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadFloat32()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_dada070e88eaa43a) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteFloat32(*((*float32)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_dada070e88eaa43a) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*float32)(__obf_dbbf371b91136494)) == 0
}

type __obf_df18f1273f2fb63e struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_df18f1273f2fb63e) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*float64)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadFloat64()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_df18f1273f2fb63e) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteFloat64(*((*float64)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_df18f1273f2fb63e) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*float64)(__obf_dbbf371b91136494)) == 0
}

type __obf_8bf4655000719b09 struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_8bf4655000719b09) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if !__obf_4ab56a99965952e7.ReadNil() {
		*((*bool)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.ReadBool()
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_8bf4655000719b09) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteBool(*((*bool)(__obf_dbbf371b91136494)))
}

func (__obf_7bfaa2f467491ab9 *__obf_8bf4655000719b09) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return !(*((*bool)(__obf_dbbf371b91136494)))
}

type __obf_cb3ab562e0bc750f struct {
	__obf_c44e8ae6d57faefd *reflect2.UnsafeSliceType
	__obf_b122f1d695f8466c ValDecoder
}

func (__obf_7bfaa2f467491ab9 *__obf_cb3ab562e0bc750f) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.ReadNil() {
		__obf_7bfaa2f467491ab9.__obf_c44e8ae6d57faefd.
			UnsafeSetNil(__obf_dbbf371b91136494)
		return
	}
	switch __obf_4ab56a99965952e7.WhatIsNext() {
	case StringValue:
		__obf_6e58d24ed6dcb371 := __obf_4ab56a99965952e7.ReadString()
		__obf_ce9859b9c202328e, __obf_fcc907dd69879566 := base64.StdEncoding.DecodeString(__obf_6e58d24ed6dcb371)
		if __obf_fcc907dd69879566 != nil {
			__obf_4ab56a99965952e7.
				ReportError("decode base64", __obf_fcc907dd69879566.Error())
		} else {
			__obf_7bfaa2f467491ab9.__obf_c44e8ae6d57faefd.
				UnsafeSet(__obf_dbbf371b91136494, unsafe.Pointer(&__obf_ce9859b9c202328e))
		}
	case ArrayValue:
		__obf_7bfaa2f467491ab9.__obf_b122f1d695f8466c.
			Decode(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
	default:
		__obf_4ab56a99965952e7.
			ReportError("base64Codec", "invalid input")
	}
}

func (__obf_7bfaa2f467491ab9 *__obf_cb3ab562e0bc750f) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	if __obf_7bfaa2f467491ab9.__obf_c44e8ae6d57faefd.UnsafeIsNil(__obf_dbbf371b91136494) {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	__obf_6e58d24ed6dcb371 := *((*[]byte)(__obf_dbbf371b91136494))
	__obf_34390500a43343b1 := base64.StdEncoding
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
	if len(__obf_6e58d24ed6dcb371) != 0 {
		__obf_452ca3698eb735a5 := __obf_34390500a43343b1.EncodedLen(len(__obf_6e58d24ed6dcb371))
		__obf_0b1656d7ef4bd9df := make([]byte, __obf_452ca3698eb735a5)
		__obf_34390500a43343b1.
			Encode(__obf_0b1656d7ef4bd9df, __obf_6e58d24ed6dcb371)
		__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df = append(__obf_8a2c51fe22775655.__obf_0b1656d7ef4bd9df, __obf_0b1656d7ef4bd9df...)
	}
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
}

func (__obf_7bfaa2f467491ab9 *__obf_cb3ab562e0bc750f) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return len(*((*[]byte)(__obf_dbbf371b91136494))) == 0
}
