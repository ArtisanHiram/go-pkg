package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

var __obf_52b2f46f997a7928 = map[string]ValDecoder{}
var __obf_6f60d6128c9e434a = map[string]ValDecoder{}
var __obf_808b92c3db71ce90 = map[string]ValEncoder{}
var __obf_88154662404be67e = map[string]ValEncoder{}
var __obf_d12055ebaa226ca1 = []Extension{}

// StructDescriptor describe how should we encode/decode the struct
type StructDescriptor struct {
	Type   reflect2.Type
	Fields []*Binding
}

// GetField get one field from the descriptor by its name.
// Can not use map here to keep field orders.
func (__obf_d31a49f9cbbb8716 *StructDescriptor) GetField(__obf_d776b95c6ab08d8c string) *Binding {
	for _, __obf_55bb392806565443 := range __obf_d31a49f9cbbb8716.Fields {
		if __obf_55bb392806565443.Field.Name() == __obf_d776b95c6ab08d8c {
			return __obf_55bb392806565443
		}
	}
	return nil
}

// Binding describe how should we encode/decode the struct field
type Binding struct {
	__obf_1c953e9c0d983f1c []int
	Field                  reflect2.StructField
	FromNames              []string
	ToNames                []string
	Encoder                ValEncoder
	Decoder                ValDecoder
}

// Extension the one for all SPI. Customize encoding/decoding by specifying alternate encoder/decoder.
// Can also rename fields by UpdateStructDescriptor.
type Extension interface {
	UpdateStructDescriptor(__obf_d31a49f9cbbb8716 *StructDescriptor)
	CreateMapKeyDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder
	CreateMapKeyEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder
	CreateDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder
	CreateEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder
	DecorateDecoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_18f746d7b5b440ee ValDecoder) ValDecoder
	DecorateEncoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_29366c3ad76a8c51 ValEncoder) ValEncoder
}

// DummyExtension embed this type get dummy implementation for all methods of Extension
type DummyExtension struct {
}

// UpdateStructDescriptor No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) UpdateStructDescriptor(__obf_d31a49f9cbbb8716 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) CreateMapKeyDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) CreateMapKeyEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) CreateDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) CreateEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) DecorateDecoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_18f746d7b5b440ee ValDecoder) ValDecoder {
	return __obf_18f746d7b5b440ee
}

// DecorateEncoder No-op
func (__obf_6b17b29e9b6f5243 *DummyExtension) DecorateEncoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_29366c3ad76a8c51 ValEncoder) ValEncoder {
	return __obf_29366c3ad76a8c51
}

type EncoderExtension map[reflect2.Type]ValEncoder

// UpdateStructDescriptor No-op
func (__obf_6b17b29e9b6f5243 EncoderExtension) UpdateStructDescriptor(__obf_d31a49f9cbbb8716 *StructDescriptor) {
}

// CreateDecoder No-op
func (__obf_6b17b29e9b6f5243 EncoderExtension) CreateDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder get encoder from map
func (__obf_6b17b29e9b6f5243 EncoderExtension) CreateEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	return __obf_6b17b29e9b6f5243[__obf_5efc66d8c338c133]
}

// CreateMapKeyDecoder No-op
func (__obf_6b17b29e9b6f5243 EncoderExtension) CreateMapKeyDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_6b17b29e9b6f5243 EncoderExtension) CreateMapKeyEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_6b17b29e9b6f5243 EncoderExtension) DecorateDecoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_18f746d7b5b440ee ValDecoder) ValDecoder {
	return __obf_18f746d7b5b440ee
}

// DecorateEncoder No-op
func (__obf_6b17b29e9b6f5243 EncoderExtension) DecorateEncoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_29366c3ad76a8c51 ValEncoder) ValEncoder {
	return __obf_29366c3ad76a8c51
}

type DecoderExtension map[reflect2.Type]ValDecoder

// UpdateStructDescriptor No-op
func (__obf_6b17b29e9b6f5243 DecoderExtension) UpdateStructDescriptor(__obf_d31a49f9cbbb8716 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_6b17b29e9b6f5243 DecoderExtension) CreateMapKeyDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_6b17b29e9b6f5243 DecoderExtension) CreateMapKeyEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder get decoder from map
func (__obf_6b17b29e9b6f5243 DecoderExtension) CreateDecoder(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	return __obf_6b17b29e9b6f5243[__obf_5efc66d8c338c133]
}

// CreateEncoder No-op
func (__obf_6b17b29e9b6f5243 DecoderExtension) CreateEncoder(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_6b17b29e9b6f5243 DecoderExtension) DecorateDecoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_18f746d7b5b440ee ValDecoder) ValDecoder {
	return __obf_18f746d7b5b440ee
}

// DecorateEncoder No-op
func (__obf_6b17b29e9b6f5243 DecoderExtension) DecorateEncoder(__obf_5efc66d8c338c133 reflect2.Type, __obf_29366c3ad76a8c51 ValEncoder) ValEncoder {
	return __obf_29366c3ad76a8c51
}

type __obf_362de96b93b8b0c3 struct {
	__obf_c0e0f4b7ab1ee699 DecoderFunc
}

func (__obf_18f746d7b5b440ee *__obf_362de96b93b8b0c3) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	__obf_18f746d7b5b440ee.__obf_c0e0f4b7ab1ee699(__obf_d3c919547bf7e47a, __obf_67008a6a9e5ba828)
}

type __obf_f00ee3350abbd8b3 struct {
	__obf_c0e0f4b7ab1ee699 EncoderFunc
	__obf_e2b0d9acf85f52c6 func(__obf_d3c919547bf7e47a unsafe.Pointer) bool
}

func (__obf_29366c3ad76a8c51 *__obf_f00ee3350abbd8b3) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_29366c3ad76a8c51.__obf_c0e0f4b7ab1ee699(__obf_d3c919547bf7e47a, __obf_00fc491caa967a74)
}

func (__obf_29366c3ad76a8c51 *__obf_f00ee3350abbd8b3) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	if __obf_29366c3ad76a8c51.__obf_e2b0d9acf85f52c6 == nil {
		return false
	}
	return __obf_29366c3ad76a8c51.__obf_e2b0d9acf85f52c6(__obf_d3c919547bf7e47a)
}

// DecoderFunc the function form of TypeDecoder
type DecoderFunc func(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator)

// EncoderFunc the function form of TypeEncoder
type EncoderFunc func(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream)

// RegisterTypeDecoderFunc register TypeDecoder for a type with function
func RegisterTypeDecoderFunc(__obf_5efc66d8c338c133 string, __obf_c0e0f4b7ab1ee699 DecoderFunc) {
	__obf_52b2f46f997a7928[__obf_5efc66d8c338c133] = &__obf_362de96b93b8b0c3{__obf_c0e0f4b7ab1ee699}
}

// RegisterTypeDecoder register TypeDecoder for a typ
func RegisterTypeDecoder(__obf_5efc66d8c338c133 string, __obf_18f746d7b5b440ee ValDecoder) {
	__obf_52b2f46f997a7928[__obf_5efc66d8c338c133] = __obf_18f746d7b5b440ee
}

// RegisterFieldDecoderFunc register TypeDecoder for a struct field with function
func RegisterFieldDecoderFunc(__obf_5efc66d8c338c133 string, __obf_61998fb083387c3c string, __obf_c0e0f4b7ab1ee699 DecoderFunc) {
	RegisterFieldDecoder(__obf_5efc66d8c338c133, __obf_61998fb083387c3c, &__obf_362de96b93b8b0c3{__obf_c0e0f4b7ab1ee699})
}

// RegisterFieldDecoder register TypeDecoder for a struct field
func RegisterFieldDecoder(__obf_5efc66d8c338c133 string, __obf_61998fb083387c3c string, __obf_18f746d7b5b440ee ValDecoder) {
	__obf_6f60d6128c9e434a[fmt.Sprintf("%s/%s", __obf_5efc66d8c338c133, __obf_61998fb083387c3c)] = __obf_18f746d7b5b440ee
}

// RegisterTypeEncoderFunc register TypeEncoder for a type with encode/isEmpty function
func RegisterTypeEncoderFunc(__obf_5efc66d8c338c133 string, __obf_c0e0f4b7ab1ee699 EncoderFunc, __obf_e2b0d9acf85f52c6 func(unsafe.Pointer) bool) {
	__obf_808b92c3db71ce90[__obf_5efc66d8c338c133] = &__obf_f00ee3350abbd8b3{__obf_c0e0f4b7ab1ee699,

		// RegisterTypeEncoder register TypeEncoder for a type
		__obf_e2b0d9acf85f52c6}
}

func RegisterTypeEncoder(__obf_5efc66d8c338c133 string, __obf_29366c3ad76a8c51 ValEncoder) {
	__obf_808b92c3db71ce90[__obf_5efc66d8c338c133] = __obf_29366c3ad76a8c51
}

// RegisterFieldEncoderFunc register TypeEncoder for a struct field with encode/isEmpty function
func RegisterFieldEncoderFunc(__obf_5efc66d8c338c133 string, __obf_61998fb083387c3c string, __obf_c0e0f4b7ab1ee699 EncoderFunc, __obf_e2b0d9acf85f52c6 func(unsafe.Pointer) bool) {
	RegisterFieldEncoder(__obf_5efc66d8c338c133, __obf_61998fb083387c3c, &__obf_f00ee3350abbd8b3{__obf_c0e0f4b7ab1ee699,

		// RegisterFieldEncoder register TypeEncoder for a struct field
		__obf_e2b0d9acf85f52c6})
}

func RegisterFieldEncoder(__obf_5efc66d8c338c133 string, __obf_61998fb083387c3c string, __obf_29366c3ad76a8c51 ValEncoder) {
	__obf_88154662404be67e[fmt.Sprintf("%s/%s", __obf_5efc66d8c338c133, __obf_61998fb083387c3c)] = __obf_29366c3ad76a8c51
}

// RegisterExtension register extension
func RegisterExtension(__obf_6b17b29e9b6f5243 Extension) {
	__obf_d12055ebaa226ca1 = append(__obf_d12055ebaa226ca1, __obf_6b17b29e9b6f5243)
}

func __obf_8b38c87279d9c5d7(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	__obf_18f746d7b5b440ee := _getTypeDecoderFromExtension(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
			__obf_18f746d7b5b440ee = __obf_6b17b29e9b6f5243.DecorateDecoder(__obf_5efc66d8c338c133, __obf_18f746d7b5b440ee)
		}
		__obf_18f746d7b5b440ee = __obf_08da24be66d0d13c.__obf_de1968c22ba7e047.DecorateDecoder(__obf_5efc66d8c338c133, __obf_18f746d7b5b440ee)
		for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
			__obf_18f746d7b5b440ee = __obf_6b17b29e9b6f5243.DecorateDecoder(__obf_5efc66d8c338c133, __obf_18f746d7b5b440ee)
		}
	}
	return __obf_18f746d7b5b440ee
}
func _getTypeDecoderFromExtension(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
		__obf_18f746d7b5b440ee := __obf_6b17b29e9b6f5243.CreateDecoder(__obf_5efc66d8c338c133)
		if __obf_18f746d7b5b440ee != nil {
			return __obf_18f746d7b5b440ee
		}
	}
	__obf_18f746d7b5b440ee := __obf_08da24be66d0d13c.__obf_de1968c22ba7e047.CreateDecoder(__obf_5efc66d8c338c133)
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_18f746d7b5b440ee := __obf_6b17b29e9b6f5243.CreateDecoder(__obf_5efc66d8c338c133)
		if __obf_18f746d7b5b440ee != nil {
			return __obf_18f746d7b5b440ee
		}
	}
	__obf_0d9188c359cc43b1 := __obf_5efc66d8c338c133.String()
	__obf_18f746d7b5b440ee = __obf_52b2f46f997a7928[__obf_0d9188c359cc43b1]
	if __obf_18f746d7b5b440ee != nil {
		return __obf_18f746d7b5b440ee
	}
	if __obf_5efc66d8c338c133.Kind() == reflect.Ptr {
		__obf_d0cac7bfcf0092ea := __obf_5efc66d8c338c133.(*reflect2.UnsafePtrType)
		__obf_18f746d7b5b440ee := __obf_52b2f46f997a7928[__obf_d0cac7bfcf0092ea.Elem().String()]
		if __obf_18f746d7b5b440ee != nil {
			return &OptionalDecoder{__obf_d0cac7bfcf0092ea.Elem(), __obf_18f746d7b5b440ee}
		}
	}
	return nil
}

func __obf_4563eaa1c4c5e82b(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	__obf_29366c3ad76a8c51 := _getTypeEncoderFromExtension(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
			__obf_29366c3ad76a8c51 = __obf_6b17b29e9b6f5243.DecorateEncoder(__obf_5efc66d8c338c133, __obf_29366c3ad76a8c51)
		}
		__obf_29366c3ad76a8c51 = __obf_08da24be66d0d13c.__obf_f754da59a9c09bdc.DecorateEncoder(__obf_5efc66d8c338c133, __obf_29366c3ad76a8c51)
		for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
			__obf_29366c3ad76a8c51 = __obf_6b17b29e9b6f5243.DecorateEncoder(__obf_5efc66d8c338c133, __obf_29366c3ad76a8c51)
		}
	}
	return __obf_29366c3ad76a8c51
}

func _getTypeEncoderFromExtension(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
		__obf_29366c3ad76a8c51 := __obf_6b17b29e9b6f5243.CreateEncoder(__obf_5efc66d8c338c133)
		if __obf_29366c3ad76a8c51 != nil {
			return __obf_29366c3ad76a8c51
		}
	}
	__obf_29366c3ad76a8c51 := __obf_08da24be66d0d13c.__obf_f754da59a9c09bdc.CreateEncoder(__obf_5efc66d8c338c133)
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_29366c3ad76a8c51 := __obf_6b17b29e9b6f5243.CreateEncoder(__obf_5efc66d8c338c133)
		if __obf_29366c3ad76a8c51 != nil {
			return __obf_29366c3ad76a8c51
		}
	}
	__obf_0d9188c359cc43b1 := __obf_5efc66d8c338c133.String()
	__obf_29366c3ad76a8c51 = __obf_808b92c3db71ce90[__obf_0d9188c359cc43b1]
	if __obf_29366c3ad76a8c51 != nil {
		return __obf_29366c3ad76a8c51
	}
	if __obf_5efc66d8c338c133.Kind() == reflect.Ptr {
		__obf_3016aca6467014d9 := __obf_5efc66d8c338c133.(*reflect2.UnsafePtrType)
		__obf_29366c3ad76a8c51 := __obf_808b92c3db71ce90[__obf_3016aca6467014d9.Elem().String()]
		if __obf_29366c3ad76a8c51 != nil {
			return &OptionalEncoder{__obf_29366c3ad76a8c51}
		}
	}
	return nil
}

func __obf_3b77508c34fb1648(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) *StructDescriptor {
	__obf_aad1f6a06bd81b25 := __obf_5efc66d8c338c133.(*reflect2.UnsafeStructType)
	__obf_a1cef2f4afac2ca6 := []*Binding{}
	__obf_bde6dee98bfb656b := []*Binding{}
	for __obf_2deec7c38ffb6ae3 := 0; __obf_2deec7c38ffb6ae3 < __obf_aad1f6a06bd81b25.NumField(); __obf_2deec7c38ffb6ae3++ {
		__obf_61998fb083387c3c := __obf_aad1f6a06bd81b25.Field(__obf_2deec7c38ffb6ae3)
		__obf_3f2421d10ca63eab, __obf_322a19b0daefbbf8 := __obf_61998fb083387c3c.Tag().Lookup(__obf_08da24be66d0d13c.__obf_8deb8923819cc0a0())
		if __obf_08da24be66d0d13c.__obf_4f5d4eba90b029f2 && !__obf_322a19b0daefbbf8 && !__obf_61998fb083387c3c.Anonymous() {
			continue
		}
		if __obf_3f2421d10ca63eab == "-" || __obf_61998fb083387c3c.Name() == "_" {
			continue
		}
		__obf_8cacc8e7ae684fb3 := strings.Split(__obf_3f2421d10ca63eab, ",")
		if __obf_61998fb083387c3c.Anonymous() && (__obf_3f2421d10ca63eab == "" || __obf_8cacc8e7ae684fb3[0] == "") {
			if __obf_61998fb083387c3c.Type().Kind() == reflect.Struct {
				__obf_d31a49f9cbbb8716 := __obf_3b77508c34fb1648(__obf_08da24be66d0d13c, __obf_61998fb083387c3c.Type())
				for _, __obf_55bb392806565443 := range __obf_d31a49f9cbbb8716.Fields {
					__obf_55bb392806565443.__obf_1c953e9c0d983f1c = append([]int{__obf_2deec7c38ffb6ae3}, __obf_55bb392806565443.__obf_1c953e9c0d983f1c...)
					__obf_be63db246ff944d0 := __obf_55bb392806565443.Encoder.(*__obf_c3b424feff43237e).__obf_be63db246ff944d0
					__obf_55bb392806565443.
						Encoder = &__obf_c3b424feff43237e{__obf_61998fb083387c3c, __obf_55bb392806565443.Encoder, __obf_be63db246ff944d0}
					__obf_55bb392806565443.
						Decoder = &__obf_2afabbd5d9366177{__obf_61998fb083387c3c, __obf_55bb392806565443.Decoder}
					__obf_a1cef2f4afac2ca6 = append(__obf_a1cef2f4afac2ca6, __obf_55bb392806565443)
				}
				continue
			} else if __obf_61998fb083387c3c.Type().Kind() == reflect.Ptr {
				__obf_d0cac7bfcf0092ea := __obf_61998fb083387c3c.Type().(*reflect2.UnsafePtrType)
				if __obf_d0cac7bfcf0092ea.Elem().Kind() == reflect.Struct {
					__obf_d31a49f9cbbb8716 := __obf_3b77508c34fb1648(__obf_08da24be66d0d13c, __obf_d0cac7bfcf0092ea.Elem())
					for _, __obf_55bb392806565443 := range __obf_d31a49f9cbbb8716.Fields {
						__obf_55bb392806565443.__obf_1c953e9c0d983f1c = append([]int{__obf_2deec7c38ffb6ae3}, __obf_55bb392806565443.__obf_1c953e9c0d983f1c...)
						__obf_be63db246ff944d0 := __obf_55bb392806565443.Encoder.(*__obf_c3b424feff43237e).__obf_be63db246ff944d0
						__obf_55bb392806565443.
							Encoder = &__obf_b7df4ea38882225e{__obf_55bb392806565443.Encoder}
						__obf_55bb392806565443.
							Encoder = &__obf_c3b424feff43237e{__obf_61998fb083387c3c, __obf_55bb392806565443.Encoder, __obf_be63db246ff944d0}
						__obf_55bb392806565443.
							Decoder = &__obf_7ad15af757dc337b{__obf_d0cac7bfcf0092ea.Elem(), __obf_55bb392806565443.Decoder}
						__obf_55bb392806565443.
							Decoder = &__obf_2afabbd5d9366177{__obf_61998fb083387c3c, __obf_55bb392806565443.Decoder}
						__obf_a1cef2f4afac2ca6 = append(__obf_a1cef2f4afac2ca6, __obf_55bb392806565443)
					}
					continue
				}
			}
		}
		__obf_09fd10be268c5b03 := __obf_bbe1ccafb6b66f13(__obf_61998fb083387c3c.Name(), __obf_8cacc8e7ae684fb3[0], __obf_3f2421d10ca63eab)
		__obf_c54a098056ab8627 := fmt.Sprintf("%s/%s", __obf_5efc66d8c338c133.String(), __obf_61998fb083387c3c.Name())
		__obf_18f746d7b5b440ee := __obf_6f60d6128c9e434a[__obf_c54a098056ab8627]
		if __obf_18f746d7b5b440ee == nil {
			__obf_18f746d7b5b440ee = __obf_c3a46fc9dd10c84e(__obf_08da24be66d0d13c.append(__obf_61998fb083387c3c.Name()), __obf_61998fb083387c3c.Type())
		}
		__obf_29366c3ad76a8c51 := __obf_88154662404be67e[__obf_c54a098056ab8627]
		if __obf_29366c3ad76a8c51 == nil {
			__obf_29366c3ad76a8c51 = __obf_3bb380f7abc03208(__obf_08da24be66d0d13c.append(__obf_61998fb083387c3c.Name()), __obf_61998fb083387c3c.Type())
		}
		__obf_55bb392806565443 := &Binding{
			Field:     __obf_61998fb083387c3c,
			FromNames: __obf_09fd10be268c5b03,
			ToNames:   __obf_09fd10be268c5b03,
			Decoder:   __obf_18f746d7b5b440ee,
			Encoder:   __obf_29366c3ad76a8c51,
		}
		__obf_55bb392806565443.__obf_1c953e9c0d983f1c = []int{__obf_2deec7c38ffb6ae3}
		__obf_bde6dee98bfb656b = append(__obf_bde6dee98bfb656b, __obf_55bb392806565443)
	}
	return __obf_c9c37271195a746d(__obf_08da24be66d0d13c, __obf_5efc66d8c338c133, __obf_bde6dee98bfb656b, __obf_a1cef2f4afac2ca6)
}
func __obf_c9c37271195a746d(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type, __obf_bde6dee98bfb656b []*Binding, __obf_a1cef2f4afac2ca6 []*Binding) *StructDescriptor {
	__obf_d31a49f9cbbb8716 := &StructDescriptor{
		Type:   __obf_5efc66d8c338c133,
		Fields: __obf_bde6dee98bfb656b,
	}
	for _, __obf_6b17b29e9b6f5243 := range __obf_d12055ebaa226ca1 {
		__obf_6b17b29e9b6f5243.
			UpdateStructDescriptor(__obf_d31a49f9cbbb8716)
	}
	__obf_08da24be66d0d13c.__obf_f754da59a9c09bdc.
		UpdateStructDescriptor(__obf_d31a49f9cbbb8716)
	__obf_08da24be66d0d13c.__obf_de1968c22ba7e047.
		UpdateStructDescriptor(__obf_d31a49f9cbbb8716)
	for _, __obf_6b17b29e9b6f5243 := range __obf_08da24be66d0d13c.__obf_251d7593467966e4 {
		__obf_6b17b29e9b6f5243.
			UpdateStructDescriptor(__obf_d31a49f9cbbb8716)
	}
	__obf_6d28554199467f58(__obf_d31a49f9cbbb8716, __obf_08da24be66d0d13c.
		// merge normal & embedded bindings & sort with original order
		__obf_5d13d7f3bd06c6cf)
	__obf_3f30ad74edb73e7a := __obf_f79346984ea959fe(append(__obf_a1cef2f4afac2ca6, __obf_d31a49f9cbbb8716.Fields...))
	sort.Sort(__obf_3f30ad74edb73e7a)
	__obf_d31a49f9cbbb8716.
		Fields = __obf_3f30ad74edb73e7a
	return __obf_d31a49f9cbbb8716
}

type __obf_f79346984ea959fe []*Binding

func (__obf_bde6dee98bfb656b __obf_f79346984ea959fe) Len() int {
	return len(__obf_bde6dee98bfb656b)
}

func (__obf_bde6dee98bfb656b __obf_f79346984ea959fe) Less(__obf_2deec7c38ffb6ae3, __obf_973404809dee3093 int) bool {
	__obf_2d558c476a3b86b1 := __obf_bde6dee98bfb656b[__obf_2deec7c38ffb6ae3].__obf_1c953e9c0d983f1c
	__obf_bff22bbd393ddcf4 := __obf_bde6dee98bfb656b[__obf_973404809dee3093].__obf_1c953e9c0d983f1c
	__obf_3a6754e3ef0a86a0 := 0
	for {
		if __obf_2d558c476a3b86b1[__obf_3a6754e3ef0a86a0] < __obf_bff22bbd393ddcf4[__obf_3a6754e3ef0a86a0] {
			return true
		} else if __obf_2d558c476a3b86b1[__obf_3a6754e3ef0a86a0] > __obf_bff22bbd393ddcf4[__obf_3a6754e3ef0a86a0] {
			return false
		}
		__obf_3a6754e3ef0a86a0++
	}
}

func (__obf_bde6dee98bfb656b __obf_f79346984ea959fe) Swap(__obf_2deec7c38ffb6ae3, __obf_973404809dee3093 int) {
	__obf_bde6dee98bfb656b[__obf_2deec7c38ffb6ae3], __obf_bde6dee98bfb656b[__obf_973404809dee3093] = __obf_bde6dee98bfb656b[__obf_973404809dee3093], __obf_bde6dee98bfb656b[__obf_2deec7c38ffb6ae3]
}

func __obf_6d28554199467f58(__obf_d31a49f9cbbb8716 *StructDescriptor, __obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) {
	for _, __obf_55bb392806565443 := range __obf_d31a49f9cbbb8716.Fields {
		__obf_dde2e0c3fd0b999c := false
		__obf_8cacc8e7ae684fb3 := strings.Split(__obf_55bb392806565443.Field.Tag().Get(__obf_dca106e1cdf85392.__obf_8deb8923819cc0a0()), ",")
		for _, __obf_f593d723fbaff897 := range __obf_8cacc8e7ae684fb3[1:] {
			if __obf_f593d723fbaff897 == "omitempty" {
				__obf_dde2e0c3fd0b999c = true
			} else if __obf_f593d723fbaff897 == "string" {
				if __obf_55bb392806565443.Field.Type().Kind() == reflect.String {
					__obf_55bb392806565443.
						Decoder = &__obf_76d674defb583cf0{__obf_55bb392806565443.Decoder, __obf_dca106e1cdf85392}
					__obf_55bb392806565443.
						Encoder = &__obf_0c596498b5fb2fef{__obf_55bb392806565443.Encoder, __obf_dca106e1cdf85392}
				} else {
					__obf_55bb392806565443.
						Decoder = &__obf_dcbce58ec66b632d{__obf_55bb392806565443.Decoder}
					__obf_55bb392806565443.
						Encoder = &__obf_8ee4ab2968aab988{__obf_55bb392806565443.Encoder}
				}
			}
		}
		__obf_55bb392806565443.
			Decoder = &__obf_2afabbd5d9366177{__obf_55bb392806565443.Field, __obf_55bb392806565443.Decoder}
		__obf_55bb392806565443.
			Encoder = &__obf_c3b424feff43237e{__obf_55bb392806565443.Field, __obf_55bb392806565443.Encoder, __obf_dde2e0c3fd0b999c}
	}
}

func __obf_bbe1ccafb6b66f13(__obf_a222db94ad8fab13 string, __obf_645185cc0f904e2b string, __obf_5cff03ea05412923 string) []string {
	// ignore?
	if __obf_5cff03ea05412923 == "-" {
		return []string{}
	}
	// rename?
	var __obf_09fd10be268c5b03 []string
	if __obf_645185cc0f904e2b == "" {
		__obf_09fd10be268c5b03 = []string{__obf_a222db94ad8fab13}
	} else {
		__obf_09fd10be268c5b03 = []string{__obf_645185cc0f904e2b}
	}
	__obf_68032455dc0b7cf0 := // private?
		unicode.IsLower(rune(__obf_a222db94ad8fab13[0])) || __obf_a222db94ad8fab13[0] == '_'
	if __obf_68032455dc0b7cf0 {
		__obf_09fd10be268c5b03 = []string{}
	}
	return __obf_09fd10be268c5b03
}
