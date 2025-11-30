package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

var __obf_f31c7837deea7cf1 = map[string]ValDecoder{}
var __obf_6b3006d32bbe7ad3 = map[string]ValDecoder{}
var __obf_8061900e068c691e = map[string]ValEncoder{}
var __obf_319f7fa6b1b719bd = map[string]ValEncoder{}
var __obf_572808a33f9754ff = []Extension{}

// StructDescriptor describe how should we encode/decode the struct
type StructDescriptor struct {
	Type   reflect2.Type
	Fields []*Binding
}

// GetField get one field from the descriptor by its name.
// Can not use map here to keep field orders.
func (__obf_832a553f1a21f913 *StructDescriptor) GetField(__obf_35302245dbc18a32 string) *Binding {
	for _, __obf_e94d87391e53ee63 := range __obf_832a553f1a21f913.Fields {
		if __obf_e94d87391e53ee63.Field.Name() == __obf_35302245dbc18a32 {
			return __obf_e94d87391e53ee63
		}
	}
	return nil
}

// Binding describe how should we encode/decode the struct field
type Binding struct {
	__obf_74792d5daf7f8edc []int
	Field                  reflect2.StructField
	FromNames              []string
	ToNames                []string
	Encoder                ValEncoder
	Decoder                ValDecoder
}

// Extension the one for all SPI. Customize encoding/decoding by specifying alternate encoder/decoder.
// Can also rename fields by UpdateStructDescriptor.
type Extension interface {
	UpdateStructDescriptor(__obf_832a553f1a21f913 *StructDescriptor)
	CreateMapKeyDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder
	CreateMapKeyEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder
	CreateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder
	CreateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder
	DecorateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_924cc7ef585cdfb0 ValDecoder) ValDecoder
	DecorateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_c85b895560db528f ValEncoder) ValEncoder
}

// DummyExtension embed this type get dummy implementation for all methods of Extension
type DummyExtension struct {
}

// UpdateStructDescriptor No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) UpdateStructDescriptor(__obf_832a553f1a21f913 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) CreateMapKeyDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) CreateMapKeyEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) CreateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) CreateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) DecorateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_924cc7ef585cdfb0 ValDecoder) ValDecoder {
	return __obf_924cc7ef585cdfb0
}

// DecorateEncoder No-op
func (__obf_e4a614f491c1bb0c *DummyExtension) DecorateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_c85b895560db528f ValEncoder) ValEncoder {
	return __obf_c85b895560db528f
}

type EncoderExtension map[reflect2.Type]ValEncoder

// UpdateStructDescriptor No-op
func (__obf_e4a614f491c1bb0c EncoderExtension) UpdateStructDescriptor(__obf_832a553f1a21f913 *StructDescriptor) {
}

// CreateDecoder No-op
func (__obf_e4a614f491c1bb0c EncoderExtension) CreateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder get encoder from map
func (__obf_e4a614f491c1bb0c EncoderExtension) CreateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	return __obf_e4a614f491c1bb0c[__obf_8667d4fc2a95b0d7]
}

// CreateMapKeyDecoder No-op
func (__obf_e4a614f491c1bb0c EncoderExtension) CreateMapKeyDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_e4a614f491c1bb0c EncoderExtension) CreateMapKeyEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_e4a614f491c1bb0c EncoderExtension) DecorateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_924cc7ef585cdfb0 ValDecoder) ValDecoder {
	return __obf_924cc7ef585cdfb0
}

// DecorateEncoder No-op
func (__obf_e4a614f491c1bb0c EncoderExtension) DecorateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_c85b895560db528f ValEncoder) ValEncoder {
	return __obf_c85b895560db528f
}

type DecoderExtension map[reflect2.Type]ValDecoder

// UpdateStructDescriptor No-op
func (__obf_e4a614f491c1bb0c DecoderExtension) UpdateStructDescriptor(__obf_832a553f1a21f913 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_e4a614f491c1bb0c DecoderExtension) CreateMapKeyDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_e4a614f491c1bb0c DecoderExtension) CreateMapKeyEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder get decoder from map
func (__obf_e4a614f491c1bb0c DecoderExtension) CreateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	return __obf_e4a614f491c1bb0c[__obf_8667d4fc2a95b0d7]
}

// CreateEncoder No-op
func (__obf_e4a614f491c1bb0c DecoderExtension) CreateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_e4a614f491c1bb0c DecoderExtension) DecorateDecoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_924cc7ef585cdfb0 ValDecoder) ValDecoder {
	return __obf_924cc7ef585cdfb0
}

// DecorateEncoder No-op
func (__obf_e4a614f491c1bb0c DecoderExtension) DecorateEncoder(__obf_8667d4fc2a95b0d7 reflect2.Type, __obf_c85b895560db528f ValEncoder) ValEncoder {
	return __obf_c85b895560db528f
}

type __obf_c4dd4b1daf1fb426 struct {
	__obf_aca452c64ee5d281 DecoderFunc
}

func (__obf_924cc7ef585cdfb0 *__obf_c4dd4b1daf1fb426) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_924cc7ef585cdfb0.__obf_aca452c64ee5d281(__obf_753ab3fb72cdd659, __obf_edd9fe48d39445e4)
}

type __obf_f57166fb6630c677 struct {
	__obf_aca452c64ee5d281 EncoderFunc
	__obf_b7251e5615e794f5 func(__obf_753ab3fb72cdd659 unsafe.Pointer) bool
}

func (__obf_c85b895560db528f *__obf_f57166fb6630c677) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_c85b895560db528f.__obf_aca452c64ee5d281(__obf_753ab3fb72cdd659, __obf_2361f5214e84c60f)
}

func (__obf_c85b895560db528f *__obf_f57166fb6630c677) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	if __obf_c85b895560db528f.__obf_b7251e5615e794f5 == nil {
		return false
	}
	return __obf_c85b895560db528f.__obf_b7251e5615e794f5(__obf_753ab3fb72cdd659)
}

// DecoderFunc the function form of TypeDecoder
type DecoderFunc func(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator)

// EncoderFunc the function form of TypeEncoder
type EncoderFunc func(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream)

// RegisterTypeDecoderFunc register TypeDecoder for a type with function
func RegisterTypeDecoderFunc(__obf_8667d4fc2a95b0d7 string, __obf_aca452c64ee5d281 DecoderFunc) {
	__obf_f31c7837deea7cf1[__obf_8667d4fc2a95b0d7] = &__obf_c4dd4b1daf1fb426{__obf_aca452c64ee5d281}
}

// RegisterTypeDecoder register TypeDecoder for a typ
func RegisterTypeDecoder(__obf_8667d4fc2a95b0d7 string, __obf_924cc7ef585cdfb0 ValDecoder) {
	__obf_f31c7837deea7cf1[__obf_8667d4fc2a95b0d7] = __obf_924cc7ef585cdfb0
}

// RegisterFieldDecoderFunc register TypeDecoder for a struct field with function
func RegisterFieldDecoderFunc(__obf_8667d4fc2a95b0d7 string, __obf_f992c91036745ca0 string, __obf_aca452c64ee5d281 DecoderFunc) {
	RegisterFieldDecoder(__obf_8667d4fc2a95b0d7, __obf_f992c91036745ca0, &__obf_c4dd4b1daf1fb426{__obf_aca452c64ee5d281})
}

// RegisterFieldDecoder register TypeDecoder for a struct field
func RegisterFieldDecoder(__obf_8667d4fc2a95b0d7 string, __obf_f992c91036745ca0 string, __obf_924cc7ef585cdfb0 ValDecoder) {
	__obf_6b3006d32bbe7ad3[fmt.Sprintf("%s/%s", __obf_8667d4fc2a95b0d7, __obf_f992c91036745ca0)] = __obf_924cc7ef585cdfb0
}

// RegisterTypeEncoderFunc register TypeEncoder for a type with encode/isEmpty function
func RegisterTypeEncoderFunc(__obf_8667d4fc2a95b0d7 string, __obf_aca452c64ee5d281 EncoderFunc, __obf_b7251e5615e794f5 func(unsafe.Pointer) bool) {
	__obf_8061900e068c691e[__obf_8667d4fc2a95b0d7] = &__obf_f57166fb6630c677{__obf_aca452c64ee5d281,

		// RegisterTypeEncoder register TypeEncoder for a type
		__obf_b7251e5615e794f5}
}

func RegisterTypeEncoder(__obf_8667d4fc2a95b0d7 string, __obf_c85b895560db528f ValEncoder) {
	__obf_8061900e068c691e[__obf_8667d4fc2a95b0d7] = __obf_c85b895560db528f
}

// RegisterFieldEncoderFunc register TypeEncoder for a struct field with encode/isEmpty function
func RegisterFieldEncoderFunc(__obf_8667d4fc2a95b0d7 string, __obf_f992c91036745ca0 string, __obf_aca452c64ee5d281 EncoderFunc, __obf_b7251e5615e794f5 func(unsafe.Pointer) bool) {
	RegisterFieldEncoder(__obf_8667d4fc2a95b0d7, __obf_f992c91036745ca0, &__obf_f57166fb6630c677{__obf_aca452c64ee5d281,

		// RegisterFieldEncoder register TypeEncoder for a struct field
		__obf_b7251e5615e794f5})
}

func RegisterFieldEncoder(__obf_8667d4fc2a95b0d7 string, __obf_f992c91036745ca0 string, __obf_c85b895560db528f ValEncoder) {
	__obf_319f7fa6b1b719bd[fmt.Sprintf("%s/%s", __obf_8667d4fc2a95b0d7, __obf_f992c91036745ca0)] = __obf_c85b895560db528f
}

// RegisterExtension register extension
func RegisterExtension(__obf_e4a614f491c1bb0c Extension) {
	__obf_572808a33f9754ff = append(__obf_572808a33f9754ff, __obf_e4a614f491c1bb0c)
}

func __obf_e17667896a54b4c5(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	__obf_924cc7ef585cdfb0 := _getTypeDecoderFromExtension(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
			__obf_924cc7ef585cdfb0 = __obf_e4a614f491c1bb0c.DecorateDecoder(__obf_8667d4fc2a95b0d7, __obf_924cc7ef585cdfb0)
		}
		__obf_924cc7ef585cdfb0 = __obf_62435d89ebefd5aa.__obf_d86968bf26bed402.DecorateDecoder(__obf_8667d4fc2a95b0d7, __obf_924cc7ef585cdfb0)
		for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
			__obf_924cc7ef585cdfb0 = __obf_e4a614f491c1bb0c.DecorateDecoder(__obf_8667d4fc2a95b0d7, __obf_924cc7ef585cdfb0)
		}
	}
	return __obf_924cc7ef585cdfb0
}
func _getTypeDecoderFromExtension(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
		__obf_924cc7ef585cdfb0 := __obf_e4a614f491c1bb0c.CreateDecoder(__obf_8667d4fc2a95b0d7)
		if __obf_924cc7ef585cdfb0 != nil {
			return __obf_924cc7ef585cdfb0
		}
	}
	__obf_924cc7ef585cdfb0 := __obf_62435d89ebefd5aa.__obf_d86968bf26bed402.CreateDecoder(__obf_8667d4fc2a95b0d7)
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_924cc7ef585cdfb0 := __obf_e4a614f491c1bb0c.CreateDecoder(__obf_8667d4fc2a95b0d7)
		if __obf_924cc7ef585cdfb0 != nil {
			return __obf_924cc7ef585cdfb0
		}
	}
	__obf_9c9c915bb768f418 := __obf_8667d4fc2a95b0d7.String()
	__obf_924cc7ef585cdfb0 = __obf_f31c7837deea7cf1[__obf_9c9c915bb768f418]
	if __obf_924cc7ef585cdfb0 != nil {
		return __obf_924cc7ef585cdfb0
	}
	if __obf_8667d4fc2a95b0d7.Kind() == reflect.Ptr {
		__obf_9fd3068ebc25d7f9 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafePtrType)
		__obf_924cc7ef585cdfb0 := __obf_f31c7837deea7cf1[__obf_9fd3068ebc25d7f9.Elem().String()]
		if __obf_924cc7ef585cdfb0 != nil {
			return &OptionalDecoder{__obf_9fd3068ebc25d7f9.Elem(), __obf_924cc7ef585cdfb0}
		}
	}
	return nil
}

func __obf_31b5d333a775e58a(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	__obf_c85b895560db528f := _getTypeEncoderFromExtension(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
			__obf_c85b895560db528f = __obf_e4a614f491c1bb0c.DecorateEncoder(__obf_8667d4fc2a95b0d7, __obf_c85b895560db528f)
		}
		__obf_c85b895560db528f = __obf_62435d89ebefd5aa.__obf_912608e689e59f9b.DecorateEncoder(__obf_8667d4fc2a95b0d7, __obf_c85b895560db528f)
		for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
			__obf_c85b895560db528f = __obf_e4a614f491c1bb0c.DecorateEncoder(__obf_8667d4fc2a95b0d7, __obf_c85b895560db528f)
		}
	}
	return __obf_c85b895560db528f
}

func _getTypeEncoderFromExtension(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
		__obf_c85b895560db528f := __obf_e4a614f491c1bb0c.CreateEncoder(__obf_8667d4fc2a95b0d7)
		if __obf_c85b895560db528f != nil {
			return __obf_c85b895560db528f
		}
	}
	__obf_c85b895560db528f := __obf_62435d89ebefd5aa.__obf_912608e689e59f9b.CreateEncoder(__obf_8667d4fc2a95b0d7)
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_c85b895560db528f := __obf_e4a614f491c1bb0c.CreateEncoder(__obf_8667d4fc2a95b0d7)
		if __obf_c85b895560db528f != nil {
			return __obf_c85b895560db528f
		}
	}
	__obf_9c9c915bb768f418 := __obf_8667d4fc2a95b0d7.String()
	__obf_c85b895560db528f = __obf_8061900e068c691e[__obf_9c9c915bb768f418]
	if __obf_c85b895560db528f != nil {
		return __obf_c85b895560db528f
	}
	if __obf_8667d4fc2a95b0d7.Kind() == reflect.Ptr {
		__obf_5316766ba02bce0e := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafePtrType)
		__obf_c85b895560db528f := __obf_8061900e068c691e[__obf_5316766ba02bce0e.Elem().String()]
		if __obf_c85b895560db528f != nil {
			return &OptionalEncoder{__obf_c85b895560db528f}
		}
	}
	return nil
}

func __obf_0446b148b9ab4187(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) *StructDescriptor {
	__obf_2805906b6b1f0342 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafeStructType)
	__obf_bbf0f5581ea2751c := []*Binding{}
	__obf_e52c8293b0cb7ead := []*Binding{}
	for __obf_28d099df85f083a8 := 0; __obf_28d099df85f083a8 < __obf_2805906b6b1f0342.NumField(); __obf_28d099df85f083a8++ {
		__obf_f992c91036745ca0 := __obf_2805906b6b1f0342.Field(__obf_28d099df85f083a8)
		__obf_fe55139bedbe5312, __obf_fa8584453469a9b9 := __obf_f992c91036745ca0.Tag().Lookup(__obf_62435d89ebefd5aa.__obf_6cf02e65d7ac8e99())
		if __obf_62435d89ebefd5aa.__obf_600fbe2319e94519 && !__obf_fa8584453469a9b9 && !__obf_f992c91036745ca0.Anonymous() {
			continue
		}
		if __obf_fe55139bedbe5312 == "-" || __obf_f992c91036745ca0.Name() == "_" {
			continue
		}
		__obf_642511981143db58 := strings.Split(__obf_fe55139bedbe5312, ",")
		if __obf_f992c91036745ca0.Anonymous() && (__obf_fe55139bedbe5312 == "" || __obf_642511981143db58[0] == "") {
			if __obf_f992c91036745ca0.Type().Kind() == reflect.Struct {
				__obf_832a553f1a21f913 := __obf_0446b148b9ab4187(__obf_62435d89ebefd5aa, __obf_f992c91036745ca0.Type())
				for _, __obf_e94d87391e53ee63 := range __obf_832a553f1a21f913.Fields {
					__obf_e94d87391e53ee63.__obf_74792d5daf7f8edc = append([]int{__obf_28d099df85f083a8}, __obf_e94d87391e53ee63.__obf_74792d5daf7f8edc...)
					__obf_87e72ba5f6e231d4 := __obf_e94d87391e53ee63.Encoder.(*__obf_8541dcc7318e179a).__obf_87e72ba5f6e231d4
					__obf_e94d87391e53ee63.
						Encoder = &__obf_8541dcc7318e179a{__obf_f992c91036745ca0, __obf_e94d87391e53ee63.Encoder, __obf_87e72ba5f6e231d4}
					__obf_e94d87391e53ee63.
						Decoder = &__obf_ca6a50aada57ad9b{__obf_f992c91036745ca0, __obf_e94d87391e53ee63.Decoder}
					__obf_bbf0f5581ea2751c = append(__obf_bbf0f5581ea2751c, __obf_e94d87391e53ee63)
				}
				continue
			} else if __obf_f992c91036745ca0.Type().Kind() == reflect.Ptr {
				__obf_9fd3068ebc25d7f9 := __obf_f992c91036745ca0.Type().(*reflect2.UnsafePtrType)
				if __obf_9fd3068ebc25d7f9.Elem().Kind() == reflect.Struct {
					__obf_832a553f1a21f913 := __obf_0446b148b9ab4187(__obf_62435d89ebefd5aa, __obf_9fd3068ebc25d7f9.Elem())
					for _, __obf_e94d87391e53ee63 := range __obf_832a553f1a21f913.Fields {
						__obf_e94d87391e53ee63.__obf_74792d5daf7f8edc = append([]int{__obf_28d099df85f083a8}, __obf_e94d87391e53ee63.__obf_74792d5daf7f8edc...)
						__obf_87e72ba5f6e231d4 := __obf_e94d87391e53ee63.Encoder.(*__obf_8541dcc7318e179a).__obf_87e72ba5f6e231d4
						__obf_e94d87391e53ee63.
							Encoder = &__obf_09858d643dd5d97b{__obf_e94d87391e53ee63.Encoder}
						__obf_e94d87391e53ee63.
							Encoder = &__obf_8541dcc7318e179a{__obf_f992c91036745ca0, __obf_e94d87391e53ee63.Encoder, __obf_87e72ba5f6e231d4}
						__obf_e94d87391e53ee63.
							Decoder = &__obf_ff8360cd52b76a62{__obf_9fd3068ebc25d7f9.Elem(), __obf_e94d87391e53ee63.Decoder}
						__obf_e94d87391e53ee63.
							Decoder = &__obf_ca6a50aada57ad9b{__obf_f992c91036745ca0, __obf_e94d87391e53ee63.Decoder}
						__obf_bbf0f5581ea2751c = append(__obf_bbf0f5581ea2751c, __obf_e94d87391e53ee63)
					}
					continue
				}
			}
		}
		__obf_e3fb24faa04348dd := __obf_da44044919672f3f(__obf_f992c91036745ca0.Name(), __obf_642511981143db58[0], __obf_fe55139bedbe5312)
		__obf_6e06b61299f4f468 := fmt.Sprintf("%s/%s", __obf_8667d4fc2a95b0d7.String(), __obf_f992c91036745ca0.Name())
		__obf_924cc7ef585cdfb0 := __obf_6b3006d32bbe7ad3[__obf_6e06b61299f4f468]
		if __obf_924cc7ef585cdfb0 == nil {
			__obf_924cc7ef585cdfb0 = __obf_eddb00a5736b0fe7(__obf_62435d89ebefd5aa.append(__obf_f992c91036745ca0.Name()), __obf_f992c91036745ca0.Type())
		}
		__obf_c85b895560db528f := __obf_319f7fa6b1b719bd[__obf_6e06b61299f4f468]
		if __obf_c85b895560db528f == nil {
			__obf_c85b895560db528f = __obf_dd4ab965a9fde81c(__obf_62435d89ebefd5aa.append(__obf_f992c91036745ca0.Name()), __obf_f992c91036745ca0.Type())
		}
		__obf_e94d87391e53ee63 := &Binding{
			Field:     __obf_f992c91036745ca0,
			FromNames: __obf_e3fb24faa04348dd,
			ToNames:   __obf_e3fb24faa04348dd,
			Decoder:   __obf_924cc7ef585cdfb0,
			Encoder:   __obf_c85b895560db528f,
		}
		__obf_e94d87391e53ee63.__obf_74792d5daf7f8edc = []int{__obf_28d099df85f083a8}
		__obf_e52c8293b0cb7ead = append(__obf_e52c8293b0cb7ead, __obf_e94d87391e53ee63)
	}
	return __obf_fed7af3a3e026f6d(__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7, __obf_e52c8293b0cb7ead, __obf_bbf0f5581ea2751c)
}
func __obf_fed7af3a3e026f6d(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type, __obf_e52c8293b0cb7ead []*Binding, __obf_bbf0f5581ea2751c []*Binding) *StructDescriptor {
	__obf_832a553f1a21f913 := &StructDescriptor{
		Type:   __obf_8667d4fc2a95b0d7,
		Fields: __obf_e52c8293b0cb7ead,
	}
	for _, __obf_e4a614f491c1bb0c := range __obf_572808a33f9754ff {
		__obf_e4a614f491c1bb0c.
			UpdateStructDescriptor(__obf_832a553f1a21f913)
	}
	__obf_62435d89ebefd5aa.__obf_912608e689e59f9b.
		UpdateStructDescriptor(__obf_832a553f1a21f913)
	__obf_62435d89ebefd5aa.__obf_d86968bf26bed402.
		UpdateStructDescriptor(__obf_832a553f1a21f913)
	for _, __obf_e4a614f491c1bb0c := range __obf_62435d89ebefd5aa.__obf_a1eac24bcc56f0a6 {
		__obf_e4a614f491c1bb0c.
			UpdateStructDescriptor(__obf_832a553f1a21f913)
	}
	__obf_80ee851ba26b9111(__obf_832a553f1a21f913, __obf_62435d89ebefd5aa.
		// merge normal & embedded bindings & sort with original order
		__obf_3a25fb4c9a8342bb)
	__obf_93adf040e1a73826 := __obf_d87686cd025b56c4(append(__obf_bbf0f5581ea2751c, __obf_832a553f1a21f913.Fields...))
	sort.Sort(__obf_93adf040e1a73826)
	__obf_832a553f1a21f913.
		Fields = __obf_93adf040e1a73826
	return __obf_832a553f1a21f913
}

type __obf_d87686cd025b56c4 []*Binding

func (__obf_e52c8293b0cb7ead __obf_d87686cd025b56c4) Len() int {
	return len(__obf_e52c8293b0cb7ead)
}

func (__obf_e52c8293b0cb7ead __obf_d87686cd025b56c4) Less(__obf_28d099df85f083a8, __obf_57cd9c621e7db075 int) bool {
	__obf_61754b25a6c67984 := __obf_e52c8293b0cb7ead[__obf_28d099df85f083a8].__obf_74792d5daf7f8edc
	__obf_4cd660dbcd4937ac := __obf_e52c8293b0cb7ead[__obf_57cd9c621e7db075].__obf_74792d5daf7f8edc
	__obf_7814f735b261cef8 := 0
	for {
		if __obf_61754b25a6c67984[__obf_7814f735b261cef8] < __obf_4cd660dbcd4937ac[__obf_7814f735b261cef8] {
			return true
		} else if __obf_61754b25a6c67984[__obf_7814f735b261cef8] > __obf_4cd660dbcd4937ac[__obf_7814f735b261cef8] {
			return false
		}
		__obf_7814f735b261cef8++
	}
}

func (__obf_e52c8293b0cb7ead __obf_d87686cd025b56c4) Swap(__obf_28d099df85f083a8, __obf_57cd9c621e7db075 int) {
	__obf_e52c8293b0cb7ead[__obf_28d099df85f083a8], __obf_e52c8293b0cb7ead[__obf_57cd9c621e7db075] = __obf_e52c8293b0cb7ead[__obf_57cd9c621e7db075], __obf_e52c8293b0cb7ead[__obf_28d099df85f083a8]
}

func __obf_80ee851ba26b9111(__obf_832a553f1a21f913 *StructDescriptor, __obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) {
	for _, __obf_e94d87391e53ee63 := range __obf_832a553f1a21f913.Fields {
		__obf_f282c1bb474af133 := false
		__obf_642511981143db58 := strings.Split(__obf_e94d87391e53ee63.Field.Tag().Get(__obf_0c8065c219ccf0ab.__obf_6cf02e65d7ac8e99()), ",")
		for _, __obf_150afed91625cd27 := range __obf_642511981143db58[1:] {
			if __obf_150afed91625cd27 == "omitempty" {
				__obf_f282c1bb474af133 = true
			} else if __obf_150afed91625cd27 == "string" {
				if __obf_e94d87391e53ee63.Field.Type().Kind() == reflect.String {
					__obf_e94d87391e53ee63.
						Decoder = &__obf_ed19b45f4bdc0c0f{__obf_e94d87391e53ee63.Decoder, __obf_0c8065c219ccf0ab}
					__obf_e94d87391e53ee63.
						Encoder = &__obf_45b40254c768957c{__obf_e94d87391e53ee63.Encoder, __obf_0c8065c219ccf0ab}
				} else {
					__obf_e94d87391e53ee63.
						Decoder = &__obf_cabfece1b2ec841f{__obf_e94d87391e53ee63.Decoder}
					__obf_e94d87391e53ee63.
						Encoder = &__obf_2e0d36011b4e847e{__obf_e94d87391e53ee63.Encoder}
				}
			}
		}
		__obf_e94d87391e53ee63.
			Decoder = &__obf_ca6a50aada57ad9b{__obf_e94d87391e53ee63.Field, __obf_e94d87391e53ee63.Decoder}
		__obf_e94d87391e53ee63.
			Encoder = &__obf_8541dcc7318e179a{__obf_e94d87391e53ee63.Field, __obf_e94d87391e53ee63.Encoder, __obf_f282c1bb474af133}
	}
}

func __obf_da44044919672f3f(__obf_ef5d88fd29433948 string, __obf_86a1b449bdf71d90 string, __obf_59e0a1533aee6e1e string) []string {
	// ignore?
	if __obf_59e0a1533aee6e1e == "-" {
		return []string{}
	}
	// rename?
	var __obf_e3fb24faa04348dd []string
	if __obf_86a1b449bdf71d90 == "" {
		__obf_e3fb24faa04348dd = []string{__obf_ef5d88fd29433948}
	} else {
		__obf_e3fb24faa04348dd = []string{__obf_86a1b449bdf71d90}
	}
	__obf_c866ef28fe7e7ef8 := // private?
		unicode.IsLower(rune(__obf_ef5d88fd29433948[0])) || __obf_ef5d88fd29433948[0] == '_'
	if __obf_c866ef28fe7e7ef8 {
		__obf_e3fb24faa04348dd = []string{}
	}
	return __obf_e3fb24faa04348dd
}
