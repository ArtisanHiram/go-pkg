package __obf_91620b895eeff9ed

import (
	"encoding"
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_8907862fdaebc8e9 = reflect2.TypeOfPtr((*json.Marshaler)(nil)).Elem()
var __obf_731bba38e4f2c30a = reflect2.TypeOfPtr((*json.Unmarshaler)(nil)).Elem()
var __obf_c32cd3ec60132ef8 = reflect2.TypeOfPtr((*encoding.TextMarshaler)(nil)).Elem()
var __obf_39d8082c5a1ebfa8 = reflect2.TypeOfPtr((*encoding.TextUnmarshaler)(nil)).Elem()

func __obf_a0300a1b15f3fca6(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_f2fdafeb141957bd := reflect2.PtrTo(__obf_29ebd0f2c324f5ea)
	if __obf_f2fdafeb141957bd.Implements(__obf_731bba38e4f2c30a) {
		return &__obf_a82dfb83d40f1b70{
			&__obf_c90773957546cc29{__obf_f2fdafeb141957bd},
		}
	}
	if __obf_f2fdafeb141957bd.Implements(__obf_39d8082c5a1ebfa8) {
		return &__obf_a82dfb83d40f1b70{
			&__obf_ebde06d9f0cb70ce{__obf_f2fdafeb141957bd},
		}
	}
	return nil
}

func __obf_656f25d36a6603f4(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	if __obf_29ebd0f2c324f5ea == __obf_8907862fdaebc8e9 {
		__obf_7a8d7d05140d0aaa := __obf_25b8173ae763425a(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
		var __obf_96e65a4c8c4f2ce5 ValEncoder = &__obf_1d8fc0268e16ab96{__obf_7a8d7d05140d0aaa: __obf_7a8d7d05140d0aaa}
		return __obf_96e65a4c8c4f2ce5
	}
	if __obf_29ebd0f2c324f5ea.Implements(__obf_8907862fdaebc8e9) {
		__obf_7a8d7d05140d0aaa := __obf_25b8173ae763425a(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
		var __obf_96e65a4c8c4f2ce5 ValEncoder = &__obf_2b5dfd7c750e7ebc{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea, __obf_7a8d7d05140d0aaa: __obf_7a8d7d05140d0aaa}
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_f2fdafeb141957bd := reflect2.PtrTo(__obf_29ebd0f2c324f5ea)
	if __obf_2f9c5aed866cce75.__obf_6e20c91e70903582 != "" && __obf_f2fdafeb141957bd.Implements(__obf_8907862fdaebc8e9) {
		__obf_7a8d7d05140d0aaa := __obf_25b8173ae763425a(__obf_2f9c5aed866cce75, __obf_f2fdafeb141957bd)
		var __obf_96e65a4c8c4f2ce5 ValEncoder = &__obf_2b5dfd7c750e7ebc{__obf_5ea8ee327a6f7e0d: __obf_f2fdafeb141957bd, __obf_7a8d7d05140d0aaa: __obf_7a8d7d05140d0aaa}
		return &__obf_9d9e26186edb63a5{__obf_96e65a4c8c4f2ce5}
	}
	if __obf_29ebd0f2c324f5ea == __obf_c32cd3ec60132ef8 {
		__obf_7a8d7d05140d0aaa := __obf_25b8173ae763425a(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
		var __obf_96e65a4c8c4f2ce5 ValEncoder = &__obf_df238ef33130aad9{__obf_7a8d7d05140d0aaa: __obf_7a8d7d05140d0aaa, __obf_7ebb8230b2719231: __obf_2f9c5aed866cce75.EncoderOf(reflect2.TypeOf(""))}
		return __obf_96e65a4c8c4f2ce5
	}
	if __obf_29ebd0f2c324f5ea.Implements(__obf_c32cd3ec60132ef8) {
		__obf_7a8d7d05140d0aaa := __obf_25b8173ae763425a(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
		var __obf_96e65a4c8c4f2ce5 ValEncoder = &__obf_9f57b5893251d4b1{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea, __obf_7ebb8230b2719231: __obf_2f9c5aed866cce75.EncoderOf(reflect2.TypeOf("")), __obf_7a8d7d05140d0aaa: __obf_7a8d7d05140d0aaa}
		return __obf_96e65a4c8c4f2ce5
	}
	// if prefix is empty, the type is the root type
	if __obf_2f9c5aed866cce75.__obf_6e20c91e70903582 != "" && __obf_f2fdafeb141957bd.Implements(__obf_c32cd3ec60132ef8) {
		__obf_7a8d7d05140d0aaa := __obf_25b8173ae763425a(__obf_2f9c5aed866cce75, __obf_f2fdafeb141957bd)
		var __obf_96e65a4c8c4f2ce5 ValEncoder = &__obf_9f57b5893251d4b1{__obf_5ea8ee327a6f7e0d: __obf_f2fdafeb141957bd, __obf_7ebb8230b2719231: __obf_2f9c5aed866cce75.EncoderOf(reflect2.TypeOf("")), __obf_7a8d7d05140d0aaa: __obf_7a8d7d05140d0aaa}
		return &__obf_9d9e26186edb63a5{__obf_96e65a4c8c4f2ce5}
	}
	return nil
}

type __obf_2b5dfd7c750e7ebc struct {
	__obf_7a8d7d05140d0aaa __obf_7a8d7d05140d0aaa
	__obf_5ea8ee327a6f7e0d reflect2.Type
}

func (__obf_96e65a4c8c4f2ce5 *__obf_2b5dfd7c750e7ebc) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_35136ef2ac0f94e4 := __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	if __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.IsNullable() && reflect2.IsNil(__obf_35136ef2ac0f94e4) {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_b4cc496dc34e127e := __obf_35136ef2ac0f94e4.(json.Marshaler)
	__obf_cdcb9f992d7bbeb0, __obf_62967739bca1d11d := __obf_b4cc496dc34e127e.MarshalJSON()
	if __obf_62967739bca1d11d != nil {
		__obf_850a7457bb739a32.
			Error = __obf_62967739bca1d11d
	} else {
		__obf_c43a08f035e1ec01 := // html escape was already done by jsoniter
			// but the extra '\n' should be trimed
			len(__obf_cdcb9f992d7bbeb0)
		if __obf_c43a08f035e1ec01 > 0 && __obf_cdcb9f992d7bbeb0[__obf_c43a08f035e1ec01-1] == '\n' {
			__obf_cdcb9f992d7bbeb0 = __obf_cdcb9f992d7bbeb0[:__obf_c43a08f035e1ec01-1]
		}
		__obf_850a7457bb739a32.
			Write(__obf_cdcb9f992d7bbeb0)
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_2b5dfd7c750e7ebc) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_7a8d7d05140d0aaa.IsEmpty(__obf_2a1474febb02279b)
}

type __obf_1d8fc0268e16ab96 struct {
	__obf_7a8d7d05140d0aaa __obf_7a8d7d05140d0aaa
}

func (__obf_96e65a4c8c4f2ce5 *__obf_1d8fc0268e16ab96) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_b4cc496dc34e127e := *(*json.Marshaler)(__obf_2a1474febb02279b)
	if __obf_b4cc496dc34e127e == nil {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_cdcb9f992d7bbeb0, __obf_62967739bca1d11d := __obf_b4cc496dc34e127e.MarshalJSON()
	if __obf_62967739bca1d11d != nil {
		__obf_850a7457bb739a32.
			Error = __obf_62967739bca1d11d
	} else {
		__obf_850a7457bb739a32.
			Write(__obf_cdcb9f992d7bbeb0)
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_1d8fc0268e16ab96) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_7a8d7d05140d0aaa.IsEmpty(__obf_2a1474febb02279b)
}

type __obf_9f57b5893251d4b1 struct {
	__obf_5ea8ee327a6f7e0d reflect2.Type
	__obf_7ebb8230b2719231 ValEncoder
	__obf_7a8d7d05140d0aaa __obf_7a8d7d05140d0aaa
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9f57b5893251d4b1) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_35136ef2ac0f94e4 := __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	if __obf_96e65a4c8c4f2ce5.__obf_5ea8ee327a6f7e0d.IsNullable() && reflect2.IsNil(__obf_35136ef2ac0f94e4) {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_b4cc496dc34e127e := (__obf_35136ef2ac0f94e4).(encoding.TextMarshaler)
	__obf_cdcb9f992d7bbeb0, __obf_62967739bca1d11d := __obf_b4cc496dc34e127e.MarshalText()
	if __obf_62967739bca1d11d != nil {
		__obf_850a7457bb739a32.
			Error = __obf_62967739bca1d11d
	} else {
		__obf_e91bd2feb751e4f1 := string(__obf_cdcb9f992d7bbeb0)
		__obf_96e65a4c8c4f2ce5.__obf_7ebb8230b2719231.
			Encode(unsafe.Pointer(&__obf_e91bd2feb751e4f1), __obf_850a7457bb739a32)
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_9f57b5893251d4b1) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_7a8d7d05140d0aaa.IsEmpty(__obf_2a1474febb02279b)
}

type __obf_df238ef33130aad9 struct {
	__obf_7ebb8230b2719231 ValEncoder
	__obf_7a8d7d05140d0aaa __obf_7a8d7d05140d0aaa
}

func (__obf_96e65a4c8c4f2ce5 *__obf_df238ef33130aad9) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_b4cc496dc34e127e := *(*encoding.TextMarshaler)(__obf_2a1474febb02279b)
	if __obf_b4cc496dc34e127e == nil {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	__obf_cdcb9f992d7bbeb0, __obf_62967739bca1d11d := __obf_b4cc496dc34e127e.MarshalText()
	if __obf_62967739bca1d11d != nil {
		__obf_850a7457bb739a32.
			Error = __obf_62967739bca1d11d
	} else {
		__obf_e91bd2feb751e4f1 := string(__obf_cdcb9f992d7bbeb0)
		__obf_96e65a4c8c4f2ce5.__obf_7ebb8230b2719231.
			Encode(unsafe.Pointer(&__obf_e91bd2feb751e4f1), __obf_850a7457bb739a32)
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_df238ef33130aad9) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_7a8d7d05140d0aaa.IsEmpty(__obf_2a1474febb02279b)
}

type __obf_c90773957546cc29 struct {
	__obf_5ea8ee327a6f7e0d reflect2.Type
}

func (__obf_6fd3f72eb9b5574c *__obf_c90773957546cc29) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_5ea8ee327a6f7e0d := __obf_6fd3f72eb9b5574c.__obf_5ea8ee327a6f7e0d
	__obf_35136ef2ac0f94e4 := __obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	__obf_a0474f4b2ac5862e := __obf_35136ef2ac0f94e4.(json.Unmarshaler)
	__obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	__obf_1bb30e8a74ed8233.
		// skip spaces
		__obf_a163df67f9bb1c4b()
	__obf_cdcb9f992d7bbeb0 := __obf_1bb30e8a74ed8233.SkipAndReturnBytes()
	__obf_62967739bca1d11d := __obf_a0474f4b2ac5862e.UnmarshalJSON(__obf_cdcb9f992d7bbeb0)
	if __obf_62967739bca1d11d != nil {
		__obf_1bb30e8a74ed8233.
			ReportError("unmarshalerDecoder", __obf_62967739bca1d11d.Error())
	}
}

type __obf_ebde06d9f0cb70ce struct {
	__obf_5ea8ee327a6f7e0d reflect2.Type
}

func (__obf_6fd3f72eb9b5574c *__obf_ebde06d9f0cb70ce) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_5ea8ee327a6f7e0d := __obf_6fd3f72eb9b5574c.__obf_5ea8ee327a6f7e0d
	__obf_35136ef2ac0f94e4 := __obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	if reflect2.IsNil(__obf_35136ef2ac0f94e4) {
		__obf_f2fdafeb141957bd := __obf_5ea8ee327a6f7e0d.(*reflect2.UnsafePtrType)
		__obf_eef27b4522cbbab1 := __obf_f2fdafeb141957bd.Elem()
		__obf_97d7079fd160f22c := __obf_eef27b4522cbbab1.UnsafeNew()
		__obf_f2fdafeb141957bd.
			UnsafeSet(__obf_2a1474febb02279b, unsafe.Pointer(&__obf_97d7079fd160f22c))
		__obf_35136ef2ac0f94e4 = __obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	}
	__obf_a0474f4b2ac5862e := (__obf_35136ef2ac0f94e4).(encoding.TextUnmarshaler)
	__obf_e91bd2feb751e4f1 := __obf_1bb30e8a74ed8233.ReadString()
	__obf_62967739bca1d11d := __obf_a0474f4b2ac5862e.UnmarshalText([]byte(__obf_e91bd2feb751e4f1))
	if __obf_62967739bca1d11d != nil {
		__obf_1bb30e8a74ed8233.
			ReportError("textUnmarshalerDecoder", __obf_62967739bca1d11d.Error())
	}
}
