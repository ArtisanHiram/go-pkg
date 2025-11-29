package __obf_91620b895eeff9ed

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"reflect"
	"unsafe"
)

func __obf_64b71f72ded515d8(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	type __obf_028b8870771eb670 struct {
		__obf_9a21bdc8399ba183 *Binding
		__obf_d734c14548ca2d7e string
		__obf_4bd4bfa272f776f7 bool
	}
	__obf_f85b672ba0272563 := []*__obf_028b8870771eb670{}
	__obf_e9a2f5e72a0f0289 := __obf_d2a025550d51a745(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	for _, __obf_9a21bdc8399ba183 := range __obf_e9a2f5e72a0f0289.Fields {
		for _, __obf_d734c14548ca2d7e := range __obf_9a21bdc8399ba183.ToNames {
			new := &__obf_028b8870771eb670{__obf_9a21bdc8399ba183: __obf_9a21bdc8399ba183, __obf_d734c14548ca2d7e: __obf_d734c14548ca2d7e}
			for _, __obf_6d85186cba056552 := range __obf_f85b672ba0272563 {
				if __obf_6d85186cba056552.__obf_d734c14548ca2d7e != __obf_d734c14548ca2d7e {
					continue
				}
				__obf_6d85186cba056552.__obf_4bd4bfa272f776f7, new.__obf_4bd4bfa272f776f7 = __obf_7edb6a45370750ef(__obf_2f9c5aed866cce75.__obf_972c2293804d141c, __obf_6d85186cba056552.__obf_9a21bdc8399ba183, new.__obf_9a21bdc8399ba183)
			}
			__obf_f85b672ba0272563 = append(__obf_f85b672ba0272563, new)
		}
	}
	if len(__obf_f85b672ba0272563) == 0 {
		return &__obf_38a23476befe63b0{}
	}
	__obf_08c9a2401941aa4b := []__obf_e35381d857924e75{}
	for _, __obf_028b8870771eb670 := range __obf_f85b672ba0272563 {
		if !__obf_028b8870771eb670.__obf_4bd4bfa272f776f7 {
			__obf_08c9a2401941aa4b = append(__obf_08c9a2401941aa4b, __obf_e35381d857924e75{__obf_96e65a4c8c4f2ce5: __obf_028b8870771eb670.__obf_9a21bdc8399ba183.Encoder.(*__obf_65fe7282c99999d2), __obf_d734c14548ca2d7e: __obf_028b8870771eb670.__obf_d734c14548ca2d7e})
		}
	}
	return &__obf_0ce7fba9d110b0e8{__obf_29ebd0f2c324f5ea, __obf_08c9a2401941aa4b}
}

func __obf_25b8173ae763425a(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) __obf_7a8d7d05140d0aaa {
	__obf_96e65a4c8c4f2ce5 := __obf_6e137b9e634e5072(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	__obf_c2f7526bd7f1cf55 := __obf_29ebd0f2c324f5ea.Kind()
	switch __obf_c2f7526bd7f1cf55 {
	case reflect.Interface:
		return &__obf_bcab05fd1729d9fa{__obf_29ebd0f2c324f5ea}
	case reflect.Struct:
		return &__obf_0ce7fba9d110b0e8{__obf_29ebd0f2c324f5ea: __obf_29ebd0f2c324f5ea}
	case reflect.Array:
		return &__obf_8dfe2d1964f42840{}
	case reflect.Slice:
		return &__obf_a70647718c3a51db{}
	case reflect.Map:
		return __obf_00ff8b3760381bbe(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	case reflect.Ptr:
		return &OptionalEncoder{}
	default:
		return &__obf_01cb1af95bf874f9{__obf_62967739bca1d11d: fmt.Errorf("unsupported type: %v", __obf_29ebd0f2c324f5ea)}
	}
}

func __obf_7edb6a45370750ef(__obf_892632c77e859037 *__obf_972c2293804d141c, __obf_6d85186cba056552, new *Binding) (__obf_f771f31cc06d5af5, __obf_685e2d4ce26180b4 bool) {
	__obf_6f8e77b256893bfd := new.Field.Tag().Get(__obf_892632c77e859037.__obf_f6ba056498acdb88()) != ""
	__obf_358b55a05cf33ec0 := __obf_6d85186cba056552.Field.Tag().Get(__obf_892632c77e859037.__obf_f6ba056498acdb88()) != ""
	if __obf_6f8e77b256893bfd {
		if __obf_358b55a05cf33ec0 {
			if len(__obf_6d85186cba056552.__obf_82efe19241f5de3d) > len(new.__obf_82efe19241f5de3d) {
				return true, false
			} else if len(new.__obf_82efe19241f5de3d) > len(__obf_6d85186cba056552.__obf_82efe19241f5de3d) {
				return false, true
			} else {
				return true, true
			}
		} else {
			return true, false
		}
	} else {
		if __obf_358b55a05cf33ec0 {
			return true, false
		}
		if len(__obf_6d85186cba056552.__obf_82efe19241f5de3d) > len(new.__obf_82efe19241f5de3d) {
			return true, false
		} else if len(new.__obf_82efe19241f5de3d) > len(__obf_6d85186cba056552.__obf_82efe19241f5de3d) {
			return false, true
		} else {
			return true, true
		}
	}
}

type __obf_65fe7282c99999d2 struct {
	__obf_7e01b5b4651074d4 reflect2.StructField
	__obf_76cbbf10485d765b ValEncoder
	__obf_75f6fd68d932045a bool
}

func (__obf_96e65a4c8c4f2ce5 *__obf_65fe7282c99999d2) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_43dac9f05df8fcc6 := __obf_96e65a4c8c4f2ce5.__obf_7e01b5b4651074d4.UnsafeGet(__obf_2a1474febb02279b)
	__obf_96e65a4c8c4f2ce5.__obf_76cbbf10485d765b.
		Encode(__obf_43dac9f05df8fcc6, __obf_850a7457bb739a32)
	if __obf_850a7457bb739a32.Error != nil && __obf_850a7457bb739a32.Error != io.EOF {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("%s: %s", __obf_96e65a4c8c4f2ce5.__obf_7e01b5b4651074d4.Name(), __obf_850a7457bb739a32.Error.Error())
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_65fe7282c99999d2) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_43dac9f05df8fcc6 := __obf_96e65a4c8c4f2ce5.__obf_7e01b5b4651074d4.UnsafeGet(__obf_2a1474febb02279b)
	return __obf_96e65a4c8c4f2ce5.__obf_76cbbf10485d765b.IsEmpty(__obf_43dac9f05df8fcc6)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_65fe7282c99999d2) IsEmbeddedPtrNil(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_c96270edda65ae80, __obf_9fb2af2f6f03e773 := __obf_96e65a4c8c4f2ce5.__obf_76cbbf10485d765b.(IsEmbeddedPtrNil)
	if !__obf_9fb2af2f6f03e773 {
		return false
	}
	__obf_43dac9f05df8fcc6 := __obf_96e65a4c8c4f2ce5.__obf_7e01b5b4651074d4.UnsafeGet(__obf_2a1474febb02279b)
	return __obf_c96270edda65ae80.IsEmbeddedPtrNil(__obf_43dac9f05df8fcc6)
}

type IsEmbeddedPtrNil interface {
	IsEmbeddedPtrNil(__obf_2a1474febb02279b unsafe.Pointer) bool
}

type __obf_0ce7fba9d110b0e8 struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_26c5a8c732658174 []__obf_e35381d857924e75
}

type __obf_e35381d857924e75 struct {
	__obf_96e65a4c8c4f2ce5 *__obf_65fe7282c99999d2
	__obf_d734c14548ca2d7e string
}

func (__obf_96e65a4c8c4f2ce5 *__obf_0ce7fba9d110b0e8) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteObjectStart()
	__obf_6007993f763f210a := false
	for _, __obf_7e01b5b4651074d4 := range __obf_96e65a4c8c4f2ce5.__obf_26c5a8c732658174 {
		if __obf_7e01b5b4651074d4.__obf_96e65a4c8c4f2ce5.__obf_75f6fd68d932045a && __obf_7e01b5b4651074d4.__obf_96e65a4c8c4f2ce5.IsEmpty(__obf_2a1474febb02279b) {
			continue
		}
		if __obf_7e01b5b4651074d4.__obf_96e65a4c8c4f2ce5.IsEmbeddedPtrNil(__obf_2a1474febb02279b) {
			continue
		}
		if __obf_6007993f763f210a {
			__obf_850a7457bb739a32.
				WriteMore()
		}
		__obf_850a7457bb739a32.
			WriteObjectField(__obf_7e01b5b4651074d4.__obf_d734c14548ca2d7e)
		__obf_7e01b5b4651074d4.__obf_96e65a4c8c4f2ce5.
			Encode(__obf_2a1474febb02279b, __obf_850a7457bb739a32)
		__obf_6007993f763f210a = true
	}
	__obf_850a7457bb739a32.
		WriteObjectEnd()
	if __obf_850a7457bb739a32.Error != nil && __obf_850a7457bb739a32.Error != io.EOF {
		__obf_850a7457bb739a32.
			Error = fmt.Errorf("%v.%s", __obf_96e65a4c8c4f2ce5.__obf_29ebd0f2c324f5ea, __obf_850a7457bb739a32.Error.Error())
	}
}

func (__obf_96e65a4c8c4f2ce5 *__obf_0ce7fba9d110b0e8) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return false
}

type __obf_38a23476befe63b0 struct {
}

func (__obf_96e65a4c8c4f2ce5 *__obf_38a23476befe63b0) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteEmptyObject()
}

func (__obf_96e65a4c8c4f2ce5 *__obf_38a23476befe63b0) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return false
}

type __obf_1b71667627f7b8fe struct {
	__obf_11328395cf86586a ValEncoder
}

func (__obf_96e65a4c8c4f2ce5 *__obf_1b71667627f7b8fe) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
	__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
		Encode(__obf_2a1474febb02279b, __obf_850a7457bb739a32)
	__obf_850a7457bb739a32.__obf_72837f6fe737f078('"')
}

func (__obf_96e65a4c8c4f2ce5 *__obf_1b71667627f7b8fe) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.IsEmpty(__obf_2a1474febb02279b)
}

type __obf_5f33847749b86fb1 struct {
	__obf_11328395cf86586a ValEncoder
	__obf_892632c77e859037 *__obf_972c2293804d141c
}

func (__obf_96e65a4c8c4f2ce5 *__obf_5f33847749b86fb1) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_0fa19adbef75b5f9 := __obf_96e65a4c8c4f2ce5.__obf_892632c77e859037.BorrowStream(nil)
	__obf_0fa19adbef75b5f9.
		Attachment = __obf_850a7457bb739a32.Attachment
	defer __obf_96e65a4c8c4f2ce5.__obf_892632c77e859037.ReturnStream(__obf_0fa19adbef75b5f9)
	__obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.
		Encode(__obf_2a1474febb02279b, __obf_0fa19adbef75b5f9)
	__obf_850a7457bb739a32.
		WriteString(string(__obf_0fa19adbef75b5f9.Buffer()))
}

func (__obf_96e65a4c8c4f2ce5 *__obf_5f33847749b86fb1) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return __obf_96e65a4c8c4f2ce5.__obf_11328395cf86586a.IsEmpty(__obf_2a1474febb02279b)
}
