package __obf_5b802ce8d9ba56d6

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type __obf_cb5eb67b8b3d51d0 struct {
	__obf_5e777ac7034ab220 reflect2.Type
}

func (__obf_29366c3ad76a8c51 *__obf_cb5eb67b8b3d51d0) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_f9b1526531f3c024 := __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	__obf_00fc491caa967a74.
		WriteVal(__obf_f9b1526531f3c024)
}

func (__obf_29366c3ad76a8c51 *__obf_cb5eb67b8b3d51d0) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a) == nil
}

type __obf_7560826a5f4553cb struct {
}

func (__obf_18f746d7b5b440ee *__obf_7560826a5f4553cb) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_d7641db63caa0309 := (*any)(__obf_d3c919547bf7e47a)
	__obf_f9b1526531f3c024 := *__obf_d7641db63caa0309
	if __obf_f9b1526531f3c024 == nil {
		*__obf_d7641db63caa0309 = __obf_67008a6a9e5ba828.Read()
		return
	}
	__obf_5efc66d8c338c133 := reflect2.TypeOf(__obf_f9b1526531f3c024)
	if __obf_5efc66d8c338c133.Kind() != reflect.Ptr {
		*__obf_d7641db63caa0309 = __obf_67008a6a9e5ba828.Read()
		return
	}
	__obf_d0cac7bfcf0092ea := __obf_5efc66d8c338c133.(*reflect2.UnsafePtrType)
	__obf_4f35ec98763f793f := __obf_d0cac7bfcf0092ea.Elem()
	if __obf_67008a6a9e5ba828.WhatIsNext() == NilValue {
		if __obf_4f35ec98763f793f.Kind() != reflect.Ptr {
			__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('n', 'u', 'l', 'l')
			*__obf_d7641db63caa0309 = nil
			return
		}
	}
	if reflect2.IsNil(__obf_f9b1526531f3c024) {
		__obf_f9b1526531f3c024 := __obf_4f35ec98763f793f.New()
		__obf_67008a6a9e5ba828.
			ReadVal(__obf_f9b1526531f3c024)
		*__obf_d7641db63caa0309 = __obf_f9b1526531f3c024
		return
	}
	__obf_67008a6a9e5ba828.
		ReadVal(__obf_f9b1526531f3c024)
}

type __obf_acb33c4a1597a4a2 struct {
	__obf_5e777ac7034ab220 *reflect2.UnsafeIFaceType
}

func (__obf_18f746d7b5b440ee *__obf_acb33c4a1597a4a2) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.ReadNil() {
		__obf_18f746d7b5b440ee.__obf_5e777ac7034ab220.
			UnsafeSet(__obf_d3c919547bf7e47a, __obf_18f746d7b5b440ee.__obf_5e777ac7034ab220.UnsafeNew())
		return
	}
	__obf_f9b1526531f3c024 := __obf_18f746d7b5b440ee.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	if reflect2.IsNil(__obf_f9b1526531f3c024) {
		__obf_67008a6a9e5ba828.
			ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	__obf_67008a6a9e5ba828.
		ReadVal(__obf_f9b1526531f3c024)
}
