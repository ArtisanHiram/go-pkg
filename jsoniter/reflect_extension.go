package __obf_c7b79b12b33d8238

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

var __obf_6117f44f95358fef = map[string]ValDecoder{}
var __obf_545223d75e12c929 = map[string]ValDecoder{}
var __obf_1cd2b048b2975dac = map[string]ValEncoder{}
var __obf_0d0da0d7a04287ff = map[string]ValEncoder{}
var __obf_0303af627a3db690 = []Extension{}

// StructDescriptor describe how should we encode/decode the struct
type StructDescriptor struct {
	Type   reflect2.Type
	Fields []*Binding
}

// GetField get one field from the descriptor by its name.
// Can not use map here to keep field orders.
func (__obf_31531899b0d36123 *StructDescriptor) GetField(__obf_dd29aaaf63f43a4a string) *Binding {
	for _, __obf_56b28257c2a1cc8d := range __obf_31531899b0d36123.Fields {
		if __obf_56b28257c2a1cc8d.Field.Name() == __obf_dd29aaaf63f43a4a {
			return __obf_56b28257c2a1cc8d
		}
	}
	return nil
}

// Binding describe how should we encode/decode the struct field
type Binding struct {
	__obf_82fe242d978d1bf2 []int
	Field                  reflect2.StructField
	FromNames              []string
	ToNames                []string
	Encoder                ValEncoder
	Decoder                ValDecoder
}

// Extension the one for all SPI. Customize encoding/decoding by specifying alternate encoder/decoder.
// Can also rename fields by UpdateStructDescriptor.
type Extension interface {
	UpdateStructDescriptor(__obf_31531899b0d36123 *StructDescriptor)
	CreateMapKeyDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder
	CreateMapKeyEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder
	CreateDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder
	CreateEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder
	DecorateDecoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_801f19702638809c ValDecoder) ValDecoder
	DecorateEncoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_c091c27480fae550 ValEncoder) ValEncoder
}

// DummyExtension embed this type get dummy implementation for all methods of Extension
type DummyExtension struct {
}

// UpdateStructDescriptor No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) UpdateStructDescriptor(__obf_31531899b0d36123 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) CreateMapKeyDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) CreateMapKeyEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) CreateDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) CreateEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) DecorateDecoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_801f19702638809c ValDecoder) ValDecoder {
	return __obf_801f19702638809c
}

// DecorateEncoder No-op
func (__obf_b8f2f726e65c4d89 *DummyExtension) DecorateEncoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_c091c27480fae550 ValEncoder) ValEncoder {
	return __obf_c091c27480fae550
}

type EncoderExtension map[reflect2.Type]ValEncoder

// UpdateStructDescriptor No-op
func (__obf_b8f2f726e65c4d89 EncoderExtension) UpdateStructDescriptor(__obf_31531899b0d36123 *StructDescriptor) {
}

// CreateDecoder No-op
func (__obf_b8f2f726e65c4d89 EncoderExtension) CreateDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder get encoder from map
func (__obf_b8f2f726e65c4d89 EncoderExtension) CreateEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	return __obf_b8f2f726e65c4d89[__obf_edcded11af6ebc4c]
}

// CreateMapKeyDecoder No-op
func (__obf_b8f2f726e65c4d89 EncoderExtension) CreateMapKeyDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_b8f2f726e65c4d89 EncoderExtension) CreateMapKeyEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_b8f2f726e65c4d89 EncoderExtension) DecorateDecoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_801f19702638809c ValDecoder) ValDecoder {
	return __obf_801f19702638809c
}

// DecorateEncoder No-op
func (__obf_b8f2f726e65c4d89 EncoderExtension) DecorateEncoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_c091c27480fae550 ValEncoder) ValEncoder {
	return __obf_c091c27480fae550
}

type DecoderExtension map[reflect2.Type]ValDecoder

// UpdateStructDescriptor No-op
func (__obf_b8f2f726e65c4d89 DecoderExtension) UpdateStructDescriptor(__obf_31531899b0d36123 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_b8f2f726e65c4d89 DecoderExtension) CreateMapKeyDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_b8f2f726e65c4d89 DecoderExtension) CreateMapKeyEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder get decoder from map
func (__obf_b8f2f726e65c4d89 DecoderExtension) CreateDecoder(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	return __obf_b8f2f726e65c4d89[__obf_edcded11af6ebc4c]
}

// CreateEncoder No-op
func (__obf_b8f2f726e65c4d89 DecoderExtension) CreateEncoder(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_b8f2f726e65c4d89 DecoderExtension) DecorateDecoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_801f19702638809c ValDecoder) ValDecoder {
	return __obf_801f19702638809c
}

// DecorateEncoder No-op
func (__obf_b8f2f726e65c4d89 DecoderExtension) DecorateEncoder(__obf_edcded11af6ebc4c reflect2.Type, __obf_c091c27480fae550 ValEncoder) ValEncoder {
	return __obf_c091c27480fae550
}

type __obf_45b8a60966351e16 struct {
	__obf_8a72d73bc63d2cad DecoderFunc
}

func (__obf_801f19702638809c *__obf_45b8a60966351e16) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_801f19702638809c.__obf_8a72d73bc63d2cad(__obf_575c04f2b9d91315, __obf_0da8c843dad13139)
}

type __obf_fd208f64584e4e8f struct {
	__obf_8a72d73bc63d2cad EncoderFunc
	__obf_564795f9f35b204e func(__obf_575c04f2b9d91315 unsafe.Pointer) bool
}

func (__obf_c091c27480fae550 *__obf_fd208f64584e4e8f) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c091c27480fae550.__obf_8a72d73bc63d2cad(__obf_575c04f2b9d91315, __obf_d8f50bcfe87d8b47)
}

func (__obf_c091c27480fae550 *__obf_fd208f64584e4e8f) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	if __obf_c091c27480fae550.__obf_564795f9f35b204e == nil {
		return false
	}
	return __obf_c091c27480fae550.__obf_564795f9f35b204e(__obf_575c04f2b9d91315)
}

// DecoderFunc the function form of TypeDecoder
type DecoderFunc func(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator)

// EncoderFunc the function form of TypeEncoder
type EncoderFunc func(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream)

// RegisterTypeDecoderFunc register TypeDecoder for a type with function
func RegisterTypeDecoderFunc(__obf_edcded11af6ebc4c string, __obf_8a72d73bc63d2cad DecoderFunc) {
	__obf_6117f44f95358fef[__obf_edcded11af6ebc4c] = &__obf_45b8a60966351e16{__obf_8a72d73bc63d2cad}
}

// RegisterTypeDecoder register TypeDecoder for a typ
func RegisterTypeDecoder(__obf_edcded11af6ebc4c string, __obf_801f19702638809c ValDecoder) {
	__obf_6117f44f95358fef[__obf_edcded11af6ebc4c] = __obf_801f19702638809c
}

// RegisterFieldDecoderFunc register TypeDecoder for a struct field with function
func RegisterFieldDecoderFunc(__obf_edcded11af6ebc4c string, __obf_ad34f8a6de2b01cb string, __obf_8a72d73bc63d2cad DecoderFunc) {
	RegisterFieldDecoder(__obf_edcded11af6ebc4c, __obf_ad34f8a6de2b01cb, &__obf_45b8a60966351e16{__obf_8a72d73bc63d2cad})
}

// RegisterFieldDecoder register TypeDecoder for a struct field
func RegisterFieldDecoder(__obf_edcded11af6ebc4c string, __obf_ad34f8a6de2b01cb string, __obf_801f19702638809c ValDecoder) {
	__obf_545223d75e12c929[fmt.Sprintf("%s/%s", __obf_edcded11af6ebc4c, __obf_ad34f8a6de2b01cb)] = __obf_801f19702638809c
}

// RegisterTypeEncoderFunc register TypeEncoder for a type with encode/isEmpty function
func RegisterTypeEncoderFunc(__obf_edcded11af6ebc4c string, __obf_8a72d73bc63d2cad EncoderFunc, __obf_564795f9f35b204e func(unsafe.Pointer) bool) {
	__obf_1cd2b048b2975dac[__obf_edcded11af6ebc4c] = &__obf_fd208f64584e4e8f{__obf_8a72d73bc63d2cad,

		// RegisterTypeEncoder register TypeEncoder for a type
		__obf_564795f9f35b204e}
}

func RegisterTypeEncoder(__obf_edcded11af6ebc4c string, __obf_c091c27480fae550 ValEncoder) {
	__obf_1cd2b048b2975dac[__obf_edcded11af6ebc4c] = __obf_c091c27480fae550
}

// RegisterFieldEncoderFunc register TypeEncoder for a struct field with encode/isEmpty function
func RegisterFieldEncoderFunc(__obf_edcded11af6ebc4c string, __obf_ad34f8a6de2b01cb string, __obf_8a72d73bc63d2cad EncoderFunc, __obf_564795f9f35b204e func(unsafe.Pointer) bool) {
	RegisterFieldEncoder(__obf_edcded11af6ebc4c, __obf_ad34f8a6de2b01cb, &__obf_fd208f64584e4e8f{__obf_8a72d73bc63d2cad,

		// RegisterFieldEncoder register TypeEncoder for a struct field
		__obf_564795f9f35b204e})
}

func RegisterFieldEncoder(__obf_edcded11af6ebc4c string, __obf_ad34f8a6de2b01cb string, __obf_c091c27480fae550 ValEncoder) {
	__obf_0d0da0d7a04287ff[fmt.Sprintf("%s/%s", __obf_edcded11af6ebc4c, __obf_ad34f8a6de2b01cb)] = __obf_c091c27480fae550
}

// RegisterExtension register extension
func RegisterExtension(__obf_b8f2f726e65c4d89 Extension) {
	__obf_0303af627a3db690 = append(__obf_0303af627a3db690, __obf_b8f2f726e65c4d89)
}

func __obf_62b401f4386d713a(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_801f19702638809c := _getTypeDecoderFromExtension(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
			__obf_801f19702638809c = __obf_b8f2f726e65c4d89.DecorateDecoder(__obf_edcded11af6ebc4c, __obf_801f19702638809c)
		}
		__obf_801f19702638809c = __obf_99ec45908bceabdb.__obf_1b8392ccffa3bed5.DecorateDecoder(__obf_edcded11af6ebc4c, __obf_801f19702638809c)
		for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
			__obf_801f19702638809c = __obf_b8f2f726e65c4d89.DecorateDecoder(__obf_edcded11af6ebc4c, __obf_801f19702638809c)
		}
	}
	return __obf_801f19702638809c
}
func _getTypeDecoderFromExtension(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
		__obf_801f19702638809c := __obf_b8f2f726e65c4d89.CreateDecoder(__obf_edcded11af6ebc4c)
		if __obf_801f19702638809c != nil {
			return __obf_801f19702638809c
		}
	}
	__obf_801f19702638809c := __obf_99ec45908bceabdb.__obf_1b8392ccffa3bed5.CreateDecoder(__obf_edcded11af6ebc4c)
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_801f19702638809c := __obf_b8f2f726e65c4d89.CreateDecoder(__obf_edcded11af6ebc4c)
		if __obf_801f19702638809c != nil {
			return __obf_801f19702638809c
		}
	}
	__obf_391099e75d7a5ccb := __obf_edcded11af6ebc4c.String()
	__obf_801f19702638809c = __obf_6117f44f95358fef[__obf_391099e75d7a5ccb]
	if __obf_801f19702638809c != nil {
		return __obf_801f19702638809c
	}
	if __obf_edcded11af6ebc4c.Kind() == reflect.Ptr {
		__obf_e2840a6d1d1a664b := __obf_edcded11af6ebc4c.(*reflect2.UnsafePtrType)
		__obf_801f19702638809c := __obf_6117f44f95358fef[__obf_e2840a6d1d1a664b.Elem().String()]
		if __obf_801f19702638809c != nil {
			return &OptionalDecoder{__obf_e2840a6d1d1a664b.Elem(), __obf_801f19702638809c}
		}
	}
	return nil
}

func __obf_558310ac59c2b1f5(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	__obf_c091c27480fae550 := _getTypeEncoderFromExtension(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
			__obf_c091c27480fae550 = __obf_b8f2f726e65c4d89.DecorateEncoder(__obf_edcded11af6ebc4c, __obf_c091c27480fae550)
		}
		__obf_c091c27480fae550 = __obf_99ec45908bceabdb.__obf_3594368cedb76ac8.DecorateEncoder(__obf_edcded11af6ebc4c, __obf_c091c27480fae550)
		for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
			__obf_c091c27480fae550 = __obf_b8f2f726e65c4d89.DecorateEncoder(__obf_edcded11af6ebc4c, __obf_c091c27480fae550)
		}
	}
	return __obf_c091c27480fae550
}

func _getTypeEncoderFromExtension(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
		__obf_c091c27480fae550 := __obf_b8f2f726e65c4d89.CreateEncoder(__obf_edcded11af6ebc4c)
		if __obf_c091c27480fae550 != nil {
			return __obf_c091c27480fae550
		}
	}
	__obf_c091c27480fae550 := __obf_99ec45908bceabdb.__obf_3594368cedb76ac8.CreateEncoder(__obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_c091c27480fae550 := __obf_b8f2f726e65c4d89.CreateEncoder(__obf_edcded11af6ebc4c)
		if __obf_c091c27480fae550 != nil {
			return __obf_c091c27480fae550
		}
	}
	__obf_391099e75d7a5ccb := __obf_edcded11af6ebc4c.String()
	__obf_c091c27480fae550 = __obf_1cd2b048b2975dac[__obf_391099e75d7a5ccb]
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	if __obf_edcded11af6ebc4c.Kind() == reflect.Ptr {
		__obf_bb77a39fa1f5f573 := __obf_edcded11af6ebc4c.(*reflect2.UnsafePtrType)
		__obf_c091c27480fae550 := __obf_1cd2b048b2975dac[__obf_bb77a39fa1f5f573.Elem().String()]
		if __obf_c091c27480fae550 != nil {
			return &OptionalEncoder{__obf_c091c27480fae550}
		}
	}
	return nil
}

func __obf_0572835446c88a68(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) *StructDescriptor {
	__obf_4f8debe2f352f35e := __obf_edcded11af6ebc4c.(*reflect2.UnsafeStructType)
	__obf_9d063d3afa521015 := []*Binding{}
	__obf_a61cdc4b92dc1d63 := []*Binding{}
	for __obf_a051269af3a541bb := 0; __obf_a051269af3a541bb < __obf_4f8debe2f352f35e.NumField(); __obf_a051269af3a541bb++ {
		__obf_ad34f8a6de2b01cb := __obf_4f8debe2f352f35e.Field(__obf_a051269af3a541bb)
		__obf_c60abe37f2d9246c, __obf_bd5750427688171c := __obf_ad34f8a6de2b01cb.Tag().Lookup(__obf_99ec45908bceabdb.__obf_2b09ab97bcd3be4f())
		if __obf_99ec45908bceabdb.__obf_6e9ee582f04bfc43 && !__obf_bd5750427688171c && !__obf_ad34f8a6de2b01cb.Anonymous() {
			continue
		}
		if __obf_c60abe37f2d9246c == "-" || __obf_ad34f8a6de2b01cb.Name() == "_" {
			continue
		}
		__obf_21f9b5c3b7dbc197 := strings.Split(__obf_c60abe37f2d9246c, ",")
		if __obf_ad34f8a6de2b01cb.Anonymous() && (__obf_c60abe37f2d9246c == "" || __obf_21f9b5c3b7dbc197[0] == "") {
			if __obf_ad34f8a6de2b01cb.Type().Kind() == reflect.Struct {
				__obf_31531899b0d36123 := __obf_0572835446c88a68(__obf_99ec45908bceabdb, __obf_ad34f8a6de2b01cb.Type())
				for _, __obf_56b28257c2a1cc8d := range __obf_31531899b0d36123.Fields {
					__obf_56b28257c2a1cc8d.__obf_82fe242d978d1bf2 = append([]int{__obf_a051269af3a541bb}, __obf_56b28257c2a1cc8d.__obf_82fe242d978d1bf2...)
					__obf_1f6282d06d1e12e1 := __obf_56b28257c2a1cc8d.Encoder.(*__obf_f750ce163e10b96f).__obf_1f6282d06d1e12e1
					__obf_56b28257c2a1cc8d.
						Encoder = &__obf_f750ce163e10b96f{__obf_ad34f8a6de2b01cb, __obf_56b28257c2a1cc8d.Encoder, __obf_1f6282d06d1e12e1}
					__obf_56b28257c2a1cc8d.
						Decoder = &__obf_7ca3f82b6bdb44f7{__obf_ad34f8a6de2b01cb, __obf_56b28257c2a1cc8d.Decoder}
					__obf_9d063d3afa521015 = append(__obf_9d063d3afa521015, __obf_56b28257c2a1cc8d)
				}
				continue
			} else if __obf_ad34f8a6de2b01cb.Type().Kind() == reflect.Ptr {
				__obf_e2840a6d1d1a664b := __obf_ad34f8a6de2b01cb.Type().(*reflect2.UnsafePtrType)
				if __obf_e2840a6d1d1a664b.Elem().Kind() == reflect.Struct {
					__obf_31531899b0d36123 := __obf_0572835446c88a68(__obf_99ec45908bceabdb, __obf_e2840a6d1d1a664b.Elem())
					for _, __obf_56b28257c2a1cc8d := range __obf_31531899b0d36123.Fields {
						__obf_56b28257c2a1cc8d.__obf_82fe242d978d1bf2 = append([]int{__obf_a051269af3a541bb}, __obf_56b28257c2a1cc8d.__obf_82fe242d978d1bf2...)
						__obf_1f6282d06d1e12e1 := __obf_56b28257c2a1cc8d.Encoder.(*__obf_f750ce163e10b96f).__obf_1f6282d06d1e12e1
						__obf_56b28257c2a1cc8d.
							Encoder = &__obf_a76744e81695632d{__obf_56b28257c2a1cc8d.Encoder}
						__obf_56b28257c2a1cc8d.
							Encoder = &__obf_f750ce163e10b96f{__obf_ad34f8a6de2b01cb, __obf_56b28257c2a1cc8d.Encoder, __obf_1f6282d06d1e12e1}
						__obf_56b28257c2a1cc8d.
							Decoder = &__obf_af9e91113c760212{__obf_e2840a6d1d1a664b.Elem(), __obf_56b28257c2a1cc8d.Decoder}
						__obf_56b28257c2a1cc8d.
							Decoder = &__obf_7ca3f82b6bdb44f7{__obf_ad34f8a6de2b01cb, __obf_56b28257c2a1cc8d.Decoder}
						__obf_9d063d3afa521015 = append(__obf_9d063d3afa521015, __obf_56b28257c2a1cc8d)
					}
					continue
				}
			}
		}
		__obf_59df7cf7cce855af := __obf_f88c0a71e016772b(__obf_ad34f8a6de2b01cb.Name(), __obf_21f9b5c3b7dbc197[0], __obf_c60abe37f2d9246c)
		__obf_6859fba6d58129c2 := fmt.Sprintf("%s/%s", __obf_edcded11af6ebc4c.String(), __obf_ad34f8a6de2b01cb.Name())
		__obf_801f19702638809c := __obf_545223d75e12c929[__obf_6859fba6d58129c2]
		if __obf_801f19702638809c == nil {
			__obf_801f19702638809c = __obf_bb14644cc3f030b3(__obf_99ec45908bceabdb.append(__obf_ad34f8a6de2b01cb.Name()), __obf_ad34f8a6de2b01cb.Type())
		}
		__obf_c091c27480fae550 := __obf_0d0da0d7a04287ff[__obf_6859fba6d58129c2]
		if __obf_c091c27480fae550 == nil {
			__obf_c091c27480fae550 = __obf_85f0e4c352da4678(__obf_99ec45908bceabdb.append(__obf_ad34f8a6de2b01cb.Name()), __obf_ad34f8a6de2b01cb.Type())
		}
		__obf_56b28257c2a1cc8d := &Binding{
			Field:     __obf_ad34f8a6de2b01cb,
			FromNames: __obf_59df7cf7cce855af,
			ToNames:   __obf_59df7cf7cce855af,
			Decoder:   __obf_801f19702638809c,
			Encoder:   __obf_c091c27480fae550,
		}
		__obf_56b28257c2a1cc8d.__obf_82fe242d978d1bf2 = []int{__obf_a051269af3a541bb}
		__obf_a61cdc4b92dc1d63 = append(__obf_a61cdc4b92dc1d63, __obf_56b28257c2a1cc8d)
	}
	return __obf_fea65b2a4142c5f2(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c, __obf_a61cdc4b92dc1d63, __obf_9d063d3afa521015)
}
func __obf_fea65b2a4142c5f2(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type, __obf_a61cdc4b92dc1d63 []*Binding, __obf_9d063d3afa521015 []*Binding) *StructDescriptor {
	__obf_31531899b0d36123 := &StructDescriptor{
		Type:   __obf_edcded11af6ebc4c,
		Fields: __obf_a61cdc4b92dc1d63,
	}
	for _, __obf_b8f2f726e65c4d89 := range __obf_0303af627a3db690 {
		__obf_b8f2f726e65c4d89.
			UpdateStructDescriptor(__obf_31531899b0d36123)
	}
	__obf_99ec45908bceabdb.__obf_3594368cedb76ac8.
		UpdateStructDescriptor(__obf_31531899b0d36123)
	__obf_99ec45908bceabdb.__obf_1b8392ccffa3bed5.
		UpdateStructDescriptor(__obf_31531899b0d36123)
	for _, __obf_b8f2f726e65c4d89 := range __obf_99ec45908bceabdb.__obf_3c82f4beb30882eb {
		__obf_b8f2f726e65c4d89.
			UpdateStructDescriptor(__obf_31531899b0d36123)
	}
	__obf_d4ff901f377329f0(__obf_31531899b0d36123, __obf_99ec45908bceabdb.
		// merge normal & embedded bindings & sort with original order
		__obf_30fe5c95cabd69c0)
	__obf_fe9daaa6db58bf73 := __obf_fc3b9603909bcd70(append(__obf_9d063d3afa521015, __obf_31531899b0d36123.Fields...))
	sort.Sort(__obf_fe9daaa6db58bf73)
	__obf_31531899b0d36123.
		Fields = __obf_fe9daaa6db58bf73
	return __obf_31531899b0d36123
}

type __obf_fc3b9603909bcd70 []*Binding

func (__obf_a61cdc4b92dc1d63 __obf_fc3b9603909bcd70) Len() int {
	return len(__obf_a61cdc4b92dc1d63)
}

func (__obf_a61cdc4b92dc1d63 __obf_fc3b9603909bcd70) Less(__obf_a051269af3a541bb, __obf_aac66f5ab0eab626 int) bool {
	__obf_f36fa24bd0ddf46a := __obf_a61cdc4b92dc1d63[__obf_a051269af3a541bb].__obf_82fe242d978d1bf2
	__obf_1bd5880624767bf8 := __obf_a61cdc4b92dc1d63[__obf_aac66f5ab0eab626].__obf_82fe242d978d1bf2
	__obf_60f36f72f8a4c230 := 0
	for {
		if __obf_f36fa24bd0ddf46a[__obf_60f36f72f8a4c230] < __obf_1bd5880624767bf8[__obf_60f36f72f8a4c230] {
			return true
		} else if __obf_f36fa24bd0ddf46a[__obf_60f36f72f8a4c230] > __obf_1bd5880624767bf8[__obf_60f36f72f8a4c230] {
			return false
		}
		__obf_60f36f72f8a4c230++
	}
}

func (__obf_a61cdc4b92dc1d63 __obf_fc3b9603909bcd70) Swap(__obf_a051269af3a541bb, __obf_aac66f5ab0eab626 int) {
	__obf_a61cdc4b92dc1d63[__obf_a051269af3a541bb], __obf_a61cdc4b92dc1d63[__obf_aac66f5ab0eab626] = __obf_a61cdc4b92dc1d63[__obf_aac66f5ab0eab626], __obf_a61cdc4b92dc1d63[__obf_a051269af3a541bb]
}

func __obf_d4ff901f377329f0(__obf_31531899b0d36123 *StructDescriptor, __obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) {
	for _, __obf_56b28257c2a1cc8d := range __obf_31531899b0d36123.Fields {
		__obf_32cd38b005394072 := false
		__obf_21f9b5c3b7dbc197 := strings.Split(__obf_56b28257c2a1cc8d.Field.Tag().Get(__obf_c52e0bcfad4b8b71.__obf_2b09ab97bcd3be4f()), ",")
		for _, __obf_f01e8c1000fb3b33 := range __obf_21f9b5c3b7dbc197[1:] {
			if __obf_f01e8c1000fb3b33 == "omitempty" {
				__obf_32cd38b005394072 = true
			} else if __obf_f01e8c1000fb3b33 == "string" {
				if __obf_56b28257c2a1cc8d.Field.Type().Kind() == reflect.String {
					__obf_56b28257c2a1cc8d.
						Decoder = &__obf_df799b1b6fef897a{__obf_56b28257c2a1cc8d.Decoder, __obf_c52e0bcfad4b8b71}
					__obf_56b28257c2a1cc8d.
						Encoder = &__obf_7de5e4e3e891d771{__obf_56b28257c2a1cc8d.Encoder, __obf_c52e0bcfad4b8b71}
				} else {
					__obf_56b28257c2a1cc8d.
						Decoder = &__obf_cdd0c9c1c11eaa23{__obf_56b28257c2a1cc8d.Decoder}
					__obf_56b28257c2a1cc8d.
						Encoder = &__obf_5e5315ac220d9e9b{__obf_56b28257c2a1cc8d.Encoder}
				}
			}
		}
		__obf_56b28257c2a1cc8d.
			Decoder = &__obf_7ca3f82b6bdb44f7{__obf_56b28257c2a1cc8d.Field, __obf_56b28257c2a1cc8d.Decoder}
		__obf_56b28257c2a1cc8d.
			Encoder = &__obf_f750ce163e10b96f{__obf_56b28257c2a1cc8d.Field, __obf_56b28257c2a1cc8d.Encoder, __obf_32cd38b005394072}
	}
}

func __obf_f88c0a71e016772b(__obf_274ed4a9a85908bc string, __obf_f9c8a71b51b323aa string, __obf_335e184da56cb76b string) []string {
	// ignore?
	if __obf_335e184da56cb76b == "-" {
		return []string{}
	}
	// rename?
	var __obf_59df7cf7cce855af []string
	if __obf_f9c8a71b51b323aa == "" {
		__obf_59df7cf7cce855af = []string{__obf_274ed4a9a85908bc}
	} else {
		__obf_59df7cf7cce855af = []string{__obf_f9c8a71b51b323aa}
	}
	__obf_4a95501227f3bc5c := // private?
		unicode.IsLower(rune(__obf_274ed4a9a85908bc[0])) || __obf_274ed4a9a85908bc[0] == '_'
	if __obf_4a95501227f3bc5c {
		__obf_59df7cf7cce855af = []string{}
	}
	return __obf_59df7cf7cce855af
}
