package __obf_703d23ba520c3295

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

var __obf_a04d76a786d6c58c = map[string]ValDecoder{}
var __obf_abe36267b721fea3 = map[string]ValDecoder{}
var __obf_9a66f8c0ef9aea38 = map[string]ValEncoder{}
var __obf_9cd47905513ef558 = map[string]ValEncoder{}
var __obf_3e7b8f77e79afd5c = []Extension{}

// StructDescriptor describe how should we encode/decode the struct
type StructDescriptor struct {
	Type   reflect2.Type
	Fields []*Binding
}

// GetField get one field from the descriptor by its name.
// Can not use map here to keep field orders.
func (__obf_ab1ac896d4aa86ee *StructDescriptor) GetField(__obf_f8b76c1746a1e630 string) *Binding {
	for _, __obf_af9543e035a21def := range __obf_ab1ac896d4aa86ee.Fields {
		if __obf_af9543e035a21def.Field.Name() == __obf_f8b76c1746a1e630 {
			return __obf_af9543e035a21def
		}
	}
	return nil
}

// Binding describe how should we encode/decode the struct field
type Binding struct {
	__obf_6c2a2b1e7bd4d99b []int
	Field                  reflect2.StructField
	FromNames              []string
	ToNames                []string
	Encoder                ValEncoder
	Decoder                ValDecoder
}

// Extension the one for all SPI. Customize encoding/decoding by specifying alternate encoder/decoder.
// Can also rename fields by UpdateStructDescriptor.
type Extension interface {
	UpdateStructDescriptor(__obf_ab1ac896d4aa86ee *StructDescriptor)
	CreateMapKeyDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder
	CreateMapKeyEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder
	CreateDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder
	CreateEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder
	DecorateDecoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_c9719e68325f7d44 ValDecoder) ValDecoder
	DecorateEncoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_cf840243a12ee308 ValEncoder) ValEncoder
}

// DummyExtension embed this type get dummy implementation for all methods of Extension
type DummyExtension struct {
}

// UpdateStructDescriptor No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) UpdateStructDescriptor(__obf_ab1ac896d4aa86ee *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) CreateMapKeyDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) CreateMapKeyEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) CreateDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) CreateEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) DecorateDecoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_c9719e68325f7d44 ValDecoder) ValDecoder {
	return __obf_c9719e68325f7d44
}

// DecorateEncoder No-op
func (__obf_dfc6ab1ee0bfd79e *DummyExtension) DecorateEncoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_cf840243a12ee308 ValEncoder) ValEncoder {
	return __obf_cf840243a12ee308
}

type EncoderExtension map[reflect2.Type]ValEncoder

// UpdateStructDescriptor No-op
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) UpdateStructDescriptor(__obf_ab1ac896d4aa86ee *StructDescriptor) {
}

// CreateDecoder No-op
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) CreateDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder get encoder from map
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) CreateEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	return __obf_dfc6ab1ee0bfd79e[__obf_3b7f6abbae19451e]
}

// CreateMapKeyDecoder No-op
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) CreateMapKeyDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) CreateMapKeyEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) DecorateDecoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_c9719e68325f7d44 ValDecoder) ValDecoder {
	return __obf_c9719e68325f7d44
}

// DecorateEncoder No-op
func (__obf_dfc6ab1ee0bfd79e EncoderExtension) DecorateEncoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_cf840243a12ee308 ValEncoder) ValEncoder {
	return __obf_cf840243a12ee308
}

type DecoderExtension map[reflect2.Type]ValDecoder

// UpdateStructDescriptor No-op
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) UpdateStructDescriptor(__obf_ab1ac896d4aa86ee *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) CreateMapKeyDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) CreateMapKeyEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder get decoder from map
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) CreateDecoder(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	return __obf_dfc6ab1ee0bfd79e[__obf_3b7f6abbae19451e]
}

// CreateEncoder No-op
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) CreateEncoder(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) DecorateDecoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_c9719e68325f7d44 ValDecoder) ValDecoder {
	return __obf_c9719e68325f7d44
}

// DecorateEncoder No-op
func (__obf_dfc6ab1ee0bfd79e DecoderExtension) DecorateEncoder(__obf_3b7f6abbae19451e reflect2.Type, __obf_cf840243a12ee308 ValEncoder) ValEncoder {
	return __obf_cf840243a12ee308
}

type __obf_b283015dd3beecea struct {
	__obf_ca0a2cef2d1adae8 DecoderFunc
}

func (__obf_c9719e68325f7d44 *__obf_b283015dd3beecea) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_c9719e68325f7d44.__obf_ca0a2cef2d1adae8(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
}

type __obf_12cc6cbc7e204422 struct {
	__obf_ca0a2cef2d1adae8 EncoderFunc
	__obf_775debf3167de731 func(__obf_47fa53f3d161f45c unsafe.Pointer) bool
}

func (__obf_cf840243a12ee308 *__obf_12cc6cbc7e204422) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_cf840243a12ee308.__obf_ca0a2cef2d1adae8(__obf_47fa53f3d161f45c, __obf_9a34f62857fb1d1d)
}

func (__obf_cf840243a12ee308 *__obf_12cc6cbc7e204422) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	if __obf_cf840243a12ee308.__obf_775debf3167de731 == nil {
		return false
	}
	return __obf_cf840243a12ee308.__obf_775debf3167de731(__obf_47fa53f3d161f45c)
}

// DecoderFunc the function form of TypeDecoder
type DecoderFunc func(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator)

// EncoderFunc the function form of TypeEncoder
type EncoderFunc func(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream)

// RegisterTypeDecoderFunc register TypeDecoder for a type with function
func RegisterTypeDecoderFunc(__obf_3b7f6abbae19451e string, __obf_ca0a2cef2d1adae8 DecoderFunc) {
	__obf_a04d76a786d6c58c[__obf_3b7f6abbae19451e] = &__obf_b283015dd3beecea{__obf_ca0a2cef2d1adae8}
}

// RegisterTypeDecoder register TypeDecoder for a typ
func RegisterTypeDecoder(__obf_3b7f6abbae19451e string, __obf_c9719e68325f7d44 ValDecoder) {
	__obf_a04d76a786d6c58c[__obf_3b7f6abbae19451e] = __obf_c9719e68325f7d44
}

// RegisterFieldDecoderFunc register TypeDecoder for a struct field with function
func RegisterFieldDecoderFunc(__obf_3b7f6abbae19451e string, __obf_13f6440419533990 string, __obf_ca0a2cef2d1adae8 DecoderFunc) {
	RegisterFieldDecoder(__obf_3b7f6abbae19451e, __obf_13f6440419533990, &__obf_b283015dd3beecea{__obf_ca0a2cef2d1adae8})
}

// RegisterFieldDecoder register TypeDecoder for a struct field
func RegisterFieldDecoder(__obf_3b7f6abbae19451e string, __obf_13f6440419533990 string, __obf_c9719e68325f7d44 ValDecoder) {
	__obf_abe36267b721fea3[fmt.Sprintf("%s/%s", __obf_3b7f6abbae19451e, __obf_13f6440419533990)] = __obf_c9719e68325f7d44
}

// RegisterTypeEncoderFunc register TypeEncoder for a type with encode/isEmpty function
func RegisterTypeEncoderFunc(__obf_3b7f6abbae19451e string, __obf_ca0a2cef2d1adae8 EncoderFunc, __obf_775debf3167de731 func(unsafe.Pointer) bool) {
	__obf_9a66f8c0ef9aea38[__obf_3b7f6abbae19451e] = &__obf_12cc6cbc7e204422{__obf_ca0a2cef2d1adae8,

		// RegisterTypeEncoder register TypeEncoder for a type
		__obf_775debf3167de731}
}

func RegisterTypeEncoder(__obf_3b7f6abbae19451e string, __obf_cf840243a12ee308 ValEncoder) {
	__obf_9a66f8c0ef9aea38[__obf_3b7f6abbae19451e] = __obf_cf840243a12ee308
}

// RegisterFieldEncoderFunc register TypeEncoder for a struct field with encode/isEmpty function
func RegisterFieldEncoderFunc(__obf_3b7f6abbae19451e string, __obf_13f6440419533990 string, __obf_ca0a2cef2d1adae8 EncoderFunc, __obf_775debf3167de731 func(unsafe.Pointer) bool) {
	RegisterFieldEncoder(__obf_3b7f6abbae19451e, __obf_13f6440419533990, &__obf_12cc6cbc7e204422{__obf_ca0a2cef2d1adae8,

		// RegisterFieldEncoder register TypeEncoder for a struct field
		__obf_775debf3167de731})
}

func RegisterFieldEncoder(__obf_3b7f6abbae19451e string, __obf_13f6440419533990 string, __obf_cf840243a12ee308 ValEncoder) {
	__obf_9cd47905513ef558[fmt.Sprintf("%s/%s", __obf_3b7f6abbae19451e, __obf_13f6440419533990)] = __obf_cf840243a12ee308
}

// RegisterExtension register extension
func RegisterExtension(__obf_dfc6ab1ee0bfd79e Extension) {
	__obf_3e7b8f77e79afd5c = append(__obf_3e7b8f77e79afd5c, __obf_dfc6ab1ee0bfd79e)
}

func __obf_3c75139ba7ab27e4(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_c9719e68325f7d44 := _getTypeDecoderFromExtension(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
			__obf_c9719e68325f7d44 = __obf_dfc6ab1ee0bfd79e.DecorateDecoder(__obf_3b7f6abbae19451e, __obf_c9719e68325f7d44)
		}
		__obf_c9719e68325f7d44 = __obf_2aaf7367de3ff86d.__obf_7c9972bb545a908d.DecorateDecoder(__obf_3b7f6abbae19451e, __obf_c9719e68325f7d44)
		for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
			__obf_c9719e68325f7d44 = __obf_dfc6ab1ee0bfd79e.DecorateDecoder(__obf_3b7f6abbae19451e, __obf_c9719e68325f7d44)
		}
	}
	return __obf_c9719e68325f7d44
}
func _getTypeDecoderFromExtension(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
		__obf_c9719e68325f7d44 := __obf_dfc6ab1ee0bfd79e.CreateDecoder(__obf_3b7f6abbae19451e)
		if __obf_c9719e68325f7d44 != nil {
			return __obf_c9719e68325f7d44
		}
	}
	__obf_c9719e68325f7d44 := __obf_2aaf7367de3ff86d.__obf_7c9972bb545a908d.CreateDecoder(__obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_c9719e68325f7d44 := __obf_dfc6ab1ee0bfd79e.CreateDecoder(__obf_3b7f6abbae19451e)
		if __obf_c9719e68325f7d44 != nil {
			return __obf_c9719e68325f7d44
		}
	}
	__obf_48b633681f0b5b99 := __obf_3b7f6abbae19451e.String()
	__obf_c9719e68325f7d44 = __obf_a04d76a786d6c58c[__obf_48b633681f0b5b99]
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	if __obf_3b7f6abbae19451e.Kind() == reflect.Ptr {
		__obf_95093efb9321684e := __obf_3b7f6abbae19451e.(*reflect2.UnsafePtrType)
		__obf_c9719e68325f7d44 := __obf_a04d76a786d6c58c[__obf_95093efb9321684e.Elem().String()]
		if __obf_c9719e68325f7d44 != nil {
			return &OptionalDecoder{__obf_95093efb9321684e.Elem(), __obf_c9719e68325f7d44}
		}
	}
	return nil
}

func __obf_6e923984d5a5d1d7(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_cf840243a12ee308 := _getTypeEncoderFromExtension(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
			__obf_cf840243a12ee308 = __obf_dfc6ab1ee0bfd79e.DecorateEncoder(__obf_3b7f6abbae19451e, __obf_cf840243a12ee308)
		}
		__obf_cf840243a12ee308 = __obf_2aaf7367de3ff86d.__obf_9289f2028bb085f2.DecorateEncoder(__obf_3b7f6abbae19451e, __obf_cf840243a12ee308)
		for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
			__obf_cf840243a12ee308 = __obf_dfc6ab1ee0bfd79e.DecorateEncoder(__obf_3b7f6abbae19451e, __obf_cf840243a12ee308)
		}
	}
	return __obf_cf840243a12ee308
}

func _getTypeEncoderFromExtension(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
		__obf_cf840243a12ee308 := __obf_dfc6ab1ee0bfd79e.CreateEncoder(__obf_3b7f6abbae19451e)
		if __obf_cf840243a12ee308 != nil {
			return __obf_cf840243a12ee308
		}
	}
	__obf_cf840243a12ee308 := __obf_2aaf7367de3ff86d.__obf_9289f2028bb085f2.CreateEncoder(__obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_cf840243a12ee308 := __obf_dfc6ab1ee0bfd79e.CreateEncoder(__obf_3b7f6abbae19451e)
		if __obf_cf840243a12ee308 != nil {
			return __obf_cf840243a12ee308
		}
	}
	__obf_48b633681f0b5b99 := __obf_3b7f6abbae19451e.String()
	__obf_cf840243a12ee308 = __obf_9a66f8c0ef9aea38[__obf_48b633681f0b5b99]
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	if __obf_3b7f6abbae19451e.Kind() == reflect.Ptr {
		__obf_a469b4f7b1b50b53 := __obf_3b7f6abbae19451e.(*reflect2.UnsafePtrType)
		__obf_cf840243a12ee308 := __obf_9a66f8c0ef9aea38[__obf_a469b4f7b1b50b53.Elem().String()]
		if __obf_cf840243a12ee308 != nil {
			return &OptionalEncoder{__obf_cf840243a12ee308}
		}
	}
	return nil
}

func __obf_969b50c34ac12a33(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) *StructDescriptor {
	__obf_586e179934393b8f := __obf_3b7f6abbae19451e.(*reflect2.UnsafeStructType)
	__obf_b561645a2d19df28 := []*Binding{}
	__obf_655058327e5c073a := []*Binding{}
	for __obf_b0a5d2bd48690f1d := 0; __obf_b0a5d2bd48690f1d < __obf_586e179934393b8f.NumField(); __obf_b0a5d2bd48690f1d++ {
		__obf_13f6440419533990 := __obf_586e179934393b8f.Field(__obf_b0a5d2bd48690f1d)
		__obf_e83958b8d81aede0, __obf_2a5727c3734ca0e0 := __obf_13f6440419533990.Tag().Lookup(__obf_2aaf7367de3ff86d.__obf_75bbb1cf9d1d76ab())
		if __obf_2aaf7367de3ff86d.__obf_eb6867c2411c1b8b && !__obf_2a5727c3734ca0e0 && !__obf_13f6440419533990.Anonymous() {
			continue
		}
		if __obf_e83958b8d81aede0 == "-" || __obf_13f6440419533990.Name() == "_" {
			continue
		}
		__obf_d5f3f8deb9c6a9a8 := strings.Split(__obf_e83958b8d81aede0, ",")
		if __obf_13f6440419533990.Anonymous() && (__obf_e83958b8d81aede0 == "" || __obf_d5f3f8deb9c6a9a8[0] == "") {
			if __obf_13f6440419533990.Type().Kind() == reflect.Struct {
				__obf_ab1ac896d4aa86ee := __obf_969b50c34ac12a33(__obf_2aaf7367de3ff86d, __obf_13f6440419533990.Type())
				for _, __obf_af9543e035a21def := range __obf_ab1ac896d4aa86ee.Fields {
					__obf_af9543e035a21def.__obf_6c2a2b1e7bd4d99b = append([]int{__obf_b0a5d2bd48690f1d}, __obf_af9543e035a21def.__obf_6c2a2b1e7bd4d99b...)
					__obf_6744c81d52458bde := __obf_af9543e035a21def.Encoder.(*__obf_2dd807eec6f01cb2).__obf_6744c81d52458bde
					__obf_af9543e035a21def.
						Encoder = &__obf_2dd807eec6f01cb2{__obf_13f6440419533990, __obf_af9543e035a21def.Encoder, __obf_6744c81d52458bde}
					__obf_af9543e035a21def.
						Decoder = &__obf_5b7a5815a547f104{__obf_13f6440419533990, __obf_af9543e035a21def.Decoder}
					__obf_b561645a2d19df28 = append(__obf_b561645a2d19df28, __obf_af9543e035a21def)
				}
				continue
			} else if __obf_13f6440419533990.Type().Kind() == reflect.Ptr {
				__obf_95093efb9321684e := __obf_13f6440419533990.Type().(*reflect2.UnsafePtrType)
				if __obf_95093efb9321684e.Elem().Kind() == reflect.Struct {
					__obf_ab1ac896d4aa86ee := __obf_969b50c34ac12a33(__obf_2aaf7367de3ff86d, __obf_95093efb9321684e.Elem())
					for _, __obf_af9543e035a21def := range __obf_ab1ac896d4aa86ee.Fields {
						__obf_af9543e035a21def.__obf_6c2a2b1e7bd4d99b = append([]int{__obf_b0a5d2bd48690f1d}, __obf_af9543e035a21def.__obf_6c2a2b1e7bd4d99b...)
						__obf_6744c81d52458bde := __obf_af9543e035a21def.Encoder.(*__obf_2dd807eec6f01cb2).__obf_6744c81d52458bde
						__obf_af9543e035a21def.
							Encoder = &__obf_d6bbd0845c4e0abb{__obf_af9543e035a21def.Encoder}
						__obf_af9543e035a21def.
							Encoder = &__obf_2dd807eec6f01cb2{__obf_13f6440419533990, __obf_af9543e035a21def.Encoder, __obf_6744c81d52458bde}
						__obf_af9543e035a21def.
							Decoder = &__obf_af91af019702e067{__obf_95093efb9321684e.Elem(), __obf_af9543e035a21def.Decoder}
						__obf_af9543e035a21def.
							Decoder = &__obf_5b7a5815a547f104{__obf_13f6440419533990, __obf_af9543e035a21def.Decoder}
						__obf_b561645a2d19df28 = append(__obf_b561645a2d19df28, __obf_af9543e035a21def)
					}
					continue
				}
			}
		}
		__obf_12001d14d7add0ef := __obf_a10e8298e006fa5c(__obf_13f6440419533990.Name(), __obf_d5f3f8deb9c6a9a8[0], __obf_e83958b8d81aede0)
		__obf_f8d7e7a99b94d7cf := fmt.Sprintf("%s/%s", __obf_3b7f6abbae19451e.String(), __obf_13f6440419533990.Name())
		__obf_c9719e68325f7d44 := __obf_abe36267b721fea3[__obf_f8d7e7a99b94d7cf]
		if __obf_c9719e68325f7d44 == nil {
			__obf_c9719e68325f7d44 = __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d.append(__obf_13f6440419533990.Name()), __obf_13f6440419533990.Type())
		}
		__obf_cf840243a12ee308 := __obf_9cd47905513ef558[__obf_f8d7e7a99b94d7cf]
		if __obf_cf840243a12ee308 == nil {
			__obf_cf840243a12ee308 = __obf_906496c658b349cf(__obf_2aaf7367de3ff86d.append(__obf_13f6440419533990.Name()), __obf_13f6440419533990.Type())
		}
		__obf_af9543e035a21def := &Binding{
			Field:     __obf_13f6440419533990,
			FromNames: __obf_12001d14d7add0ef,
			ToNames:   __obf_12001d14d7add0ef,
			Decoder:   __obf_c9719e68325f7d44,
			Encoder:   __obf_cf840243a12ee308,
		}
		__obf_af9543e035a21def.__obf_6c2a2b1e7bd4d99b = []int{__obf_b0a5d2bd48690f1d}
		__obf_655058327e5c073a = append(__obf_655058327e5c073a, __obf_af9543e035a21def)
	}
	return __obf_fd06b54381b8f1da(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e, __obf_655058327e5c073a, __obf_b561645a2d19df28)
}
func __obf_fd06b54381b8f1da(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type, __obf_655058327e5c073a []*Binding, __obf_b561645a2d19df28 []*Binding) *StructDescriptor {
	__obf_ab1ac896d4aa86ee := &StructDescriptor{
		Type:   __obf_3b7f6abbae19451e,
		Fields: __obf_655058327e5c073a,
	}
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_3e7b8f77e79afd5c {
		__obf_dfc6ab1ee0bfd79e.
			UpdateStructDescriptor(__obf_ab1ac896d4aa86ee)
	}
	__obf_2aaf7367de3ff86d.__obf_9289f2028bb085f2.
		UpdateStructDescriptor(__obf_ab1ac896d4aa86ee)
	__obf_2aaf7367de3ff86d.__obf_7c9972bb545a908d.
		UpdateStructDescriptor(__obf_ab1ac896d4aa86ee)
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_dfc6ab1ee0bfd79e.
			UpdateStructDescriptor(__obf_ab1ac896d4aa86ee)
	}
	__obf_6d130218ee92e841(__obf_ab1ac896d4aa86ee, __obf_2aaf7367de3ff86d.
		// merge normal & embedded bindings & sort with original order
		__obf_ef74248d8ae9ba78)
	__obf_43dffb81f9069eb0 := __obf_5a99c79c931af5f9(append(__obf_b561645a2d19df28, __obf_ab1ac896d4aa86ee.Fields...))
	sort.Sort(__obf_43dffb81f9069eb0)
	__obf_ab1ac896d4aa86ee.
		Fields = __obf_43dffb81f9069eb0
	return __obf_ab1ac896d4aa86ee
}

type __obf_5a99c79c931af5f9 []*Binding

func (__obf_655058327e5c073a __obf_5a99c79c931af5f9) Len() int {
	return len(__obf_655058327e5c073a)
}

func (__obf_655058327e5c073a __obf_5a99c79c931af5f9) Less(__obf_b0a5d2bd48690f1d, __obf_0e04c22ffdf339de int) bool {
	__obf_c9c790a409e0ace2 := __obf_655058327e5c073a[__obf_b0a5d2bd48690f1d].__obf_6c2a2b1e7bd4d99b
	__obf_4d4a4291859f11c5 := __obf_655058327e5c073a[__obf_0e04c22ffdf339de].__obf_6c2a2b1e7bd4d99b
	__obf_a95ada6cc83385bf := 0
	for {
		if __obf_c9c790a409e0ace2[__obf_a95ada6cc83385bf] < __obf_4d4a4291859f11c5[__obf_a95ada6cc83385bf] {
			return true
		} else if __obf_c9c790a409e0ace2[__obf_a95ada6cc83385bf] > __obf_4d4a4291859f11c5[__obf_a95ada6cc83385bf] {
			return false
		}
		__obf_a95ada6cc83385bf++
	}
}

func (__obf_655058327e5c073a __obf_5a99c79c931af5f9) Swap(__obf_b0a5d2bd48690f1d, __obf_0e04c22ffdf339de int) {
	__obf_655058327e5c073a[__obf_b0a5d2bd48690f1d], __obf_655058327e5c073a[__obf_0e04c22ffdf339de] = __obf_655058327e5c073a[__obf_0e04c22ffdf339de], __obf_655058327e5c073a[__obf_b0a5d2bd48690f1d]
}

func __obf_6d130218ee92e841(__obf_ab1ac896d4aa86ee *StructDescriptor, __obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) {
	for _, __obf_af9543e035a21def := range __obf_ab1ac896d4aa86ee.Fields {
		__obf_d3842adbd9543233 := false
		__obf_d5f3f8deb9c6a9a8 := strings.Split(__obf_af9543e035a21def.Field.Tag().Get(__obf_25bd4f754a37b862.__obf_75bbb1cf9d1d76ab()), ",")
		for _, __obf_519b4a8fe56d9ba1 := range __obf_d5f3f8deb9c6a9a8[1:] {
			if __obf_519b4a8fe56d9ba1 == "omitempty" {
				__obf_d3842adbd9543233 = true
			} else if __obf_519b4a8fe56d9ba1 == "string" {
				if __obf_af9543e035a21def.Field.Type().Kind() == reflect.String {
					__obf_af9543e035a21def.
						Decoder = &__obf_bfcd3685f73f8364{__obf_af9543e035a21def.Decoder, __obf_25bd4f754a37b862}
					__obf_af9543e035a21def.
						Encoder = &__obf_9f3b818ecbe63ff5{__obf_af9543e035a21def.Encoder, __obf_25bd4f754a37b862}
				} else {
					__obf_af9543e035a21def.
						Decoder = &__obf_e2eefa9baff51ea8{__obf_af9543e035a21def.Decoder}
					__obf_af9543e035a21def.
						Encoder = &__obf_d9cbb1c64198bf84{__obf_af9543e035a21def.Encoder}
				}
			}
		}
		__obf_af9543e035a21def.
			Decoder = &__obf_5b7a5815a547f104{__obf_af9543e035a21def.Field, __obf_af9543e035a21def.Decoder}
		__obf_af9543e035a21def.
			Encoder = &__obf_2dd807eec6f01cb2{__obf_af9543e035a21def.Field, __obf_af9543e035a21def.Encoder, __obf_d3842adbd9543233}
	}
}

func __obf_a10e8298e006fa5c(__obf_db01ea7643ae2801 string, __obf_072e9e10c26516a6 string, __obf_502dbdfc54e3e05c string) []string {
	// ignore?
	if __obf_502dbdfc54e3e05c == "-" {
		return []string{}
	}
	// rename?
	var __obf_12001d14d7add0ef []string
	if __obf_072e9e10c26516a6 == "" {
		__obf_12001d14d7add0ef = []string{__obf_db01ea7643ae2801}
	} else {
		__obf_12001d14d7add0ef = []string{__obf_072e9e10c26516a6}
	}
	__obf_15eab5f5821c64c5 := // private?
		unicode.IsLower(rune(__obf_db01ea7643ae2801[0])) || __obf_db01ea7643ae2801[0] == '_'
	if __obf_15eab5f5821c64c5 {
		__obf_12001d14d7add0ef = []string{}
	}
	return __obf_12001d14d7add0ef
}
