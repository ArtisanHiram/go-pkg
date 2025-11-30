package __obf_703d23ba520c3295

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type __obf_7f31cdbc4d682d50 struct {
	__obf_15da62f311934a45 reflect2.Type
}

func (__obf_cf840243a12ee308 *__obf_7f31cdbc4d682d50) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_02ba706f4f104d71 := __obf_cf840243a12ee308.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	__obf_9a34f62857fb1d1d.
		WriteVal(__obf_02ba706f4f104d71)
}

func (__obf_cf840243a12ee308 *__obf_7f31cdbc4d682d50) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c) == nil
}

type __obf_a2ef1129a194d05a struct {
}

func (__obf_c9719e68325f7d44 *__obf_a2ef1129a194d05a) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_d8e151a9d1e22c71 := (*any)(__obf_47fa53f3d161f45c)
	__obf_02ba706f4f104d71 := *__obf_d8e151a9d1e22c71
	if __obf_02ba706f4f104d71 == nil {
		*__obf_d8e151a9d1e22c71 = __obf_47edab4c16a2d88d.Read()
		return
	}
	__obf_3b7f6abbae19451e := reflect2.TypeOf(__obf_02ba706f4f104d71)
	if __obf_3b7f6abbae19451e.Kind() != reflect.Ptr {
		*__obf_d8e151a9d1e22c71 = __obf_47edab4c16a2d88d.Read()
		return
	}
	__obf_95093efb9321684e := __obf_3b7f6abbae19451e.(*reflect2.UnsafePtrType)
	__obf_1a739b127dbf2386 := __obf_95093efb9321684e.Elem()
	if __obf_47edab4c16a2d88d.WhatIsNext() == NilValue {
		if __obf_1a739b127dbf2386.Kind() != reflect.Ptr {
			__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('n', 'u', 'l', 'l')
			*__obf_d8e151a9d1e22c71 = nil
			return
		}
	}
	if reflect2.IsNil(__obf_02ba706f4f104d71) {
		__obf_02ba706f4f104d71 := __obf_1a739b127dbf2386.New()
		__obf_47edab4c16a2d88d.
			ReadVal(__obf_02ba706f4f104d71)
		*__obf_d8e151a9d1e22c71 = __obf_02ba706f4f104d71
		return
	}
	__obf_47edab4c16a2d88d.
		ReadVal(__obf_02ba706f4f104d71)
}

type __obf_a2ad861d69c8a6ea struct {
	__obf_15da62f311934a45 *reflect2.UnsafeIFaceType
}

func (__obf_c9719e68325f7d44 *__obf_a2ad861d69c8a6ea) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.ReadNil() {
		__obf_c9719e68325f7d44.__obf_15da62f311934a45.
			UnsafeSet(__obf_47fa53f3d161f45c, __obf_c9719e68325f7d44.__obf_15da62f311934a45.UnsafeNew())
		return
	}
	__obf_02ba706f4f104d71 := __obf_c9719e68325f7d44.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	if reflect2.IsNil(__obf_02ba706f4f104d71) {
		__obf_47edab4c16a2d88d.
			ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	__obf_47edab4c16a2d88d.
		ReadVal(__obf_02ba706f4f104d71)
}
