package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"reflect"
	"unsafe"
)

func __obf_04c6c702681f689e(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	type __obf_929e209ad681acdc struct {
		__obf_55bb392806565443 *Binding
		__obf_15e31015b126fe45 string
		__obf_10c5525c7ccb34ea bool
	}
	__obf_397ff9ea32d284f1 := []*__obf_929e209ad681acdc{}
	__obf_d31a49f9cbbb8716 := __obf_3b77508c34fb1648(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	for _, __obf_55bb392806565443 := range __obf_d31a49f9cbbb8716.Fields {
		for _, __obf_15e31015b126fe45 := range __obf_55bb392806565443.ToNames {
			new := &__obf_929e209ad681acdc{__obf_55bb392806565443: __obf_55bb392806565443, __obf_15e31015b126fe45: __obf_15e31015b126fe45}
			for _, __obf_3668b60d03008ac5 := range __obf_397ff9ea32d284f1 {
				if __obf_3668b60d03008ac5.__obf_15e31015b126fe45 != __obf_15e31015b126fe45 {
					continue
				}
				__obf_3668b60d03008ac5.__obf_10c5525c7ccb34ea, new.__obf_10c5525c7ccb34ea = __obf_5902434c0363a1d9(__obf_08da24be66d0d13c.__obf_5d13d7f3bd06c6cf, __obf_3668b60d03008ac5.__obf_55bb392806565443, new.__obf_55bb392806565443)
			}
			__obf_397ff9ea32d284f1 = append(__obf_397ff9ea32d284f1, new)
		}
	}
	if len(__obf_397ff9ea32d284f1) == 0 {
		return &__obf_9c5603bc7a544f2e{}
	}
	__obf_86b58674c46829fd := []__obf_01618dc0a558c688{}
	for _, __obf_929e209ad681acdc := range __obf_397ff9ea32d284f1 {
		if !__obf_929e209ad681acdc.__obf_10c5525c7ccb34ea {
			__obf_86b58674c46829fd = append(__obf_86b58674c46829fd, __obf_01618dc0a558c688{__obf_29366c3ad76a8c51: __obf_929e209ad681acdc.__obf_55bb392806565443.Encoder.(*__obf_c3b424feff43237e), __obf_15e31015b126fe45: __obf_929e209ad681acdc.__obf_15e31015b126fe45})
		}
	}
	return &__obf_dc4ba147cc0cfc4f{__obf_5efc66d8c338c133, __obf_86b58674c46829fd}
}

func __obf_35a60ba99630ce31(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) __obf_9cd9bb2e92b7d8c2 {
	__obf_29366c3ad76a8c51 := __obf_062c5f9f03d60bc9(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	__obf_f1764122678433dc := __obf_5efc66d8c338c133.Kind()
	switch __obf_f1764122678433dc {
	case reflect.Interface:
		return &__obf_cb5eb67b8b3d51d0{__obf_5efc66d8c338c133}
	case reflect.Struct:
		return &__obf_dc4ba147cc0cfc4f{__obf_5efc66d8c338c133: __obf_5efc66d8c338c133}
	case reflect.Array:
		return &__obf_55abf1f4a08610a6{}
	case reflect.Slice:
		return &__obf_47f54f698749f41b{}
	case reflect.Map:
		return __obf_8fcbff4be3a3cf2e(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	case reflect.Ptr:
		return &OptionalEncoder{}
	default:
		return &__obf_5d9ef7d0fc378907{__obf_6d071d48f3b321ad: fmt.Errorf("unsupported type: %v", __obf_5efc66d8c338c133)}
	}
}

func __obf_5902434c0363a1d9(__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf, __obf_3668b60d03008ac5, new *Binding) (__obf_e29aa548946ea7a6, __obf_95fe2be3f6af92bc bool) {
	__obf_3eee7db1d4c56bb2 := new.Field.Tag().Get(__obf_dca106e1cdf85392.__obf_8deb8923819cc0a0()) != ""
	__obf_15d4022ec471d2c2 := __obf_3668b60d03008ac5.Field.Tag().Get(__obf_dca106e1cdf85392.__obf_8deb8923819cc0a0()) != ""
	if __obf_3eee7db1d4c56bb2 {
		if __obf_15d4022ec471d2c2 {
			if len(__obf_3668b60d03008ac5.__obf_1c953e9c0d983f1c) > len(new.__obf_1c953e9c0d983f1c) {
				return true, false
			} else if len(new.__obf_1c953e9c0d983f1c) > len(__obf_3668b60d03008ac5.__obf_1c953e9c0d983f1c) {
				return false, true
			} else {
				return true, true
			}
		} else {
			return true, false
		}
	} else {
		if __obf_15d4022ec471d2c2 {
			return true, false
		}
		if len(__obf_3668b60d03008ac5.__obf_1c953e9c0d983f1c) > len(new.__obf_1c953e9c0d983f1c) {
			return true, false
		} else if len(new.__obf_1c953e9c0d983f1c) > len(__obf_3668b60d03008ac5.__obf_1c953e9c0d983f1c) {
			return false, true
		} else {
			return true, true
		}
	}
}

type __obf_c3b424feff43237e struct {
	__obf_61998fb083387c3c reflect2.StructField
	__obf_5919e20e7ba09338 ValEncoder
	__obf_be63db246ff944d0 bool
}

func (__obf_29366c3ad76a8c51 *__obf_c3b424feff43237e) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_a6502f11c1c12aab := __obf_29366c3ad76a8c51.__obf_61998fb083387c3c.UnsafeGet(__obf_d3c919547bf7e47a)
	__obf_29366c3ad76a8c51.__obf_5919e20e7ba09338.
		Encode(__obf_a6502f11c1c12aab, __obf_00fc491caa967a74)
	if __obf_00fc491caa967a74.Error != nil && __obf_00fc491caa967a74.Error != io.EOF {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("%s: %s", __obf_29366c3ad76a8c51.__obf_61998fb083387c3c.Name(), __obf_00fc491caa967a74.Error.Error())
	}
}

func (__obf_29366c3ad76a8c51 *__obf_c3b424feff43237e) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_a6502f11c1c12aab := __obf_29366c3ad76a8c51.__obf_61998fb083387c3c.UnsafeGet(__obf_d3c919547bf7e47a)
	return __obf_29366c3ad76a8c51.__obf_5919e20e7ba09338.IsEmpty(__obf_a6502f11c1c12aab)
}

func (__obf_29366c3ad76a8c51 *__obf_c3b424feff43237e) IsEmbeddedPtrNil(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_8fcef6e2878720af, __obf_ecd3f2241015639e := __obf_29366c3ad76a8c51.__obf_5919e20e7ba09338.(IsEmbeddedPtrNil)
	if !__obf_ecd3f2241015639e {
		return false
	}
	__obf_a6502f11c1c12aab := __obf_29366c3ad76a8c51.__obf_61998fb083387c3c.UnsafeGet(__obf_d3c919547bf7e47a)
	return __obf_8fcef6e2878720af.IsEmbeddedPtrNil(__obf_a6502f11c1c12aab)
}

type IsEmbeddedPtrNil interface {
	IsEmbeddedPtrNil(__obf_d3c919547bf7e47a unsafe.Pointer) bool
}

type __obf_dc4ba147cc0cfc4f struct {
	__obf_5efc66d8c338c133 reflect2.Type
	__obf_db858406797b4036 []__obf_01618dc0a558c688
}

type __obf_01618dc0a558c688 struct {
	__obf_29366c3ad76a8c51 *__obf_c3b424feff43237e
	__obf_15e31015b126fe45 string
}

func (__obf_29366c3ad76a8c51 *__obf_dc4ba147cc0cfc4f) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteObjectStart()
	__obf_70dad9974e23dabd := false
	for _, __obf_61998fb083387c3c := range __obf_29366c3ad76a8c51.__obf_db858406797b4036 {
		if __obf_61998fb083387c3c.__obf_29366c3ad76a8c51.__obf_be63db246ff944d0 && __obf_61998fb083387c3c.__obf_29366c3ad76a8c51.IsEmpty(__obf_d3c919547bf7e47a) {
			continue
		}
		if __obf_61998fb083387c3c.__obf_29366c3ad76a8c51.IsEmbeddedPtrNil(__obf_d3c919547bf7e47a) {
			continue
		}
		if __obf_70dad9974e23dabd {
			__obf_00fc491caa967a74.
				WriteMore()
		}
		__obf_00fc491caa967a74.
			WriteObjectField(__obf_61998fb083387c3c.__obf_15e31015b126fe45)
		__obf_61998fb083387c3c.__obf_29366c3ad76a8c51.
			Encode(__obf_d3c919547bf7e47a, __obf_00fc491caa967a74)
		__obf_70dad9974e23dabd = true
	}
	__obf_00fc491caa967a74.
		WriteObjectEnd()
	if __obf_00fc491caa967a74.Error != nil && __obf_00fc491caa967a74.Error != io.EOF {
		__obf_00fc491caa967a74.
			Error = fmt.Errorf("%v.%s", __obf_29366c3ad76a8c51.__obf_5efc66d8c338c133, __obf_00fc491caa967a74.Error.Error())
	}
}

func (__obf_29366c3ad76a8c51 *__obf_dc4ba147cc0cfc4f) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return false
}

type __obf_9c5603bc7a544f2e struct {
}

func (__obf_29366c3ad76a8c51 *__obf_9c5603bc7a544f2e) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteEmptyObject()
}

func (__obf_29366c3ad76a8c51 *__obf_9c5603bc7a544f2e) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return false
}

type __obf_8ee4ab2968aab988 struct {
	__obf_e57e229f9a2faf9a ValEncoder
}

func (__obf_29366c3ad76a8c51 *__obf_8ee4ab2968aab988) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
	__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
		Encode(__obf_d3c919547bf7e47a, __obf_00fc491caa967a74)
	__obf_00fc491caa967a74.__obf_487ee8035c7dd8f6('"')
}

func (__obf_29366c3ad76a8c51 *__obf_8ee4ab2968aab988) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.IsEmpty(__obf_d3c919547bf7e47a)
}

type __obf_0c596498b5fb2fef struct {
	__obf_e57e229f9a2faf9a ValEncoder
	__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf
}

func (__obf_29366c3ad76a8c51 *__obf_0c596498b5fb2fef) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_bd6159819589a79a := __obf_29366c3ad76a8c51.__obf_dca106e1cdf85392.BorrowStream(nil)
	__obf_bd6159819589a79a.
		Attachment = __obf_00fc491caa967a74.Attachment
	defer __obf_29366c3ad76a8c51.__obf_dca106e1cdf85392.ReturnStream(__obf_bd6159819589a79a)
	__obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.
		Encode(__obf_d3c919547bf7e47a, __obf_bd6159819589a79a)
	__obf_00fc491caa967a74.
		WriteString(string(__obf_bd6159819589a79a.Buffer()))
}

func (__obf_29366c3ad76a8c51 *__obf_0c596498b5fb2fef) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return __obf_29366c3ad76a8c51.__obf_e57e229f9a2faf9a.IsEmpty(__obf_d3c919547bf7e47a)
}
