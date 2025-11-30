package __obf_030d39f7a8de96c6

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type __obf_3691828d7304f473 struct {
	__obf_a7610e23a63622fd reflect2.Type
}

func (__obf_008f61596d7e9523 *__obf_3691828d7304f473) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_eefa92392cc2442c := __obf_008f61596d7e9523.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	__obf_8a2c51fe22775655.
		WriteVal(__obf_eefa92392cc2442c)
}

func (__obf_008f61596d7e9523 *__obf_3691828d7304f473) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494) == nil
}

type __obf_b4fd1ed1693bf959 struct {
}

func (__obf_11a3f28116164d09 *__obf_b4fd1ed1693bf959) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_1d11bee5a14363a1 := (*any)(__obf_dbbf371b91136494)
	__obf_eefa92392cc2442c := *__obf_1d11bee5a14363a1
	if __obf_eefa92392cc2442c == nil {
		*__obf_1d11bee5a14363a1 = __obf_4ab56a99965952e7.Read()
		return
	}
	__obf_8744d0b8c80ccdc1 := reflect2.TypeOf(__obf_eefa92392cc2442c)
	if __obf_8744d0b8c80ccdc1.Kind() != reflect.Ptr {
		*__obf_1d11bee5a14363a1 = __obf_4ab56a99965952e7.Read()
		return
	}
	__obf_3a51157620f9910b := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafePtrType)
	__obf_b1437a4f3535fc68 := __obf_3a51157620f9910b.Elem()
	if __obf_4ab56a99965952e7.WhatIsNext() == NilValue {
		if __obf_b1437a4f3535fc68.Kind() != reflect.Ptr {
			__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('n', 'u', 'l', 'l')
			*__obf_1d11bee5a14363a1 = nil
			return
		}
	}
	if reflect2.IsNil(__obf_eefa92392cc2442c) {
		__obf_eefa92392cc2442c := __obf_b1437a4f3535fc68.New()
		__obf_4ab56a99965952e7.
			ReadVal(__obf_eefa92392cc2442c)
		*__obf_1d11bee5a14363a1 = __obf_eefa92392cc2442c
		return
	}
	__obf_4ab56a99965952e7.
		ReadVal(__obf_eefa92392cc2442c)
}

type __obf_ee8a1f4d7afb7af9 struct {
	__obf_a7610e23a63622fd *reflect2.UnsafeIFaceType
}

func (__obf_11a3f28116164d09 *__obf_ee8a1f4d7afb7af9) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	if __obf_4ab56a99965952e7.ReadNil() {
		__obf_11a3f28116164d09.__obf_a7610e23a63622fd.
			UnsafeSet(__obf_dbbf371b91136494, __obf_11a3f28116164d09.__obf_a7610e23a63622fd.UnsafeNew())
		return
	}
	__obf_eefa92392cc2442c := __obf_11a3f28116164d09.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	if reflect2.IsNil(__obf_eefa92392cc2442c) {
		__obf_4ab56a99965952e7.
			ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	__obf_4ab56a99965952e7.
		ReadVal(__obf_eefa92392cc2442c)
}
