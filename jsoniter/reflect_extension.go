package __obf_91620b895eeff9ed

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

var __obf_c0b846a36a3844aa = map[string]ValDecoder{}
var __obf_c7a2a8e5abb2ff8a = map[string]ValDecoder{}
var __obf_ad65186839258b68 = map[string]ValEncoder{}
var __obf_afeecc600805f008 = map[string]ValEncoder{}
var __obf_bb7ace2aaa9423b9 = []Extension{}

// StructDescriptor describe how should we encode/decode the struct
type StructDescriptor struct {
	Type   reflect2.Type
	Fields []*Binding
}

// GetField get one field from the descriptor by its name.
// Can not use map here to keep field orders.
func (__obf_e9a2f5e72a0f0289 *StructDescriptor) GetField(__obf_04fb01d7218c6db3 string) *Binding {
	for _, __obf_9a21bdc8399ba183 := range __obf_e9a2f5e72a0f0289.Fields {
		if __obf_9a21bdc8399ba183.Field.Name() == __obf_04fb01d7218c6db3 {
			return __obf_9a21bdc8399ba183
		}
	}
	return nil
}

// Binding describe how should we encode/decode the struct field
type Binding struct {
	__obf_82efe19241f5de3d []int
	Field                  reflect2.StructField
	FromNames              []string
	ToNames                []string
	Encoder                ValEncoder
	Decoder                ValDecoder
}

// Extension the one for all SPI. Customize encoding/decoding by specifying alternate encoder/decoder.
// Can also rename fields by UpdateStructDescriptor.
type Extension interface {
	UpdateStructDescriptor(__obf_e9a2f5e72a0f0289 *StructDescriptor)
	CreateMapKeyDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder
	CreateMapKeyEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder
	CreateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder
	CreateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder
	DecorateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_6fd3f72eb9b5574c ValDecoder) ValDecoder
	DecorateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_96e65a4c8c4f2ce5 ValEncoder) ValEncoder
}

// DummyExtension embed this type get dummy implementation for all methods of Extension
type DummyExtension struct {
}

// UpdateStructDescriptor No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) UpdateStructDescriptor(__obf_e9a2f5e72a0f0289 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) CreateMapKeyDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) CreateMapKeyEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) CreateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) CreateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) DecorateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_6fd3f72eb9b5574c ValDecoder) ValDecoder {
	return __obf_6fd3f72eb9b5574c
}

// DecorateEncoder No-op
func (__obf_8c9d73a8f8319687 *DummyExtension) DecorateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_96e65a4c8c4f2ce5 ValEncoder) ValEncoder {
	return __obf_96e65a4c8c4f2ce5
}

type EncoderExtension map[reflect2.Type]ValEncoder

// UpdateStructDescriptor No-op
func (__obf_8c9d73a8f8319687 EncoderExtension) UpdateStructDescriptor(__obf_e9a2f5e72a0f0289 *StructDescriptor) {
}

// CreateDecoder No-op
func (__obf_8c9d73a8f8319687 EncoderExtension) CreateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder get encoder from map
func (__obf_8c9d73a8f8319687 EncoderExtension) CreateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	return __obf_8c9d73a8f8319687[__obf_29ebd0f2c324f5ea]
}

// CreateMapKeyDecoder No-op
func (__obf_8c9d73a8f8319687 EncoderExtension) CreateMapKeyDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_8c9d73a8f8319687 EncoderExtension) CreateMapKeyEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_8c9d73a8f8319687 EncoderExtension) DecorateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_6fd3f72eb9b5574c ValDecoder) ValDecoder {
	return __obf_6fd3f72eb9b5574c
}

// DecorateEncoder No-op
func (__obf_8c9d73a8f8319687 EncoderExtension) DecorateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_96e65a4c8c4f2ce5 ValEncoder) ValEncoder {
	return __obf_96e65a4c8c4f2ce5
}

type DecoderExtension map[reflect2.Type]ValDecoder

// UpdateStructDescriptor No-op
func (__obf_8c9d73a8f8319687 DecoderExtension) UpdateStructDescriptor(__obf_e9a2f5e72a0f0289 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_8c9d73a8f8319687 DecoderExtension) CreateMapKeyDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_8c9d73a8f8319687 DecoderExtension) CreateMapKeyEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder get decoder from map
func (__obf_8c9d73a8f8319687 DecoderExtension) CreateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	return __obf_8c9d73a8f8319687[__obf_29ebd0f2c324f5ea]
}

// CreateEncoder No-op
func (__obf_8c9d73a8f8319687 DecoderExtension) CreateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_8c9d73a8f8319687 DecoderExtension) DecorateDecoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_6fd3f72eb9b5574c ValDecoder) ValDecoder {
	return __obf_6fd3f72eb9b5574c
}

// DecorateEncoder No-op
func (__obf_8c9d73a8f8319687 DecoderExtension) DecorateEncoder(__obf_29ebd0f2c324f5ea reflect2.Type, __obf_96e65a4c8c4f2ce5 ValEncoder) ValEncoder {
	return __obf_96e65a4c8c4f2ce5
}

type __obf_b33e9e2851d3e569 struct {
	__obf_32c2c59b26b6e25b DecoderFunc
}

func (__obf_6fd3f72eb9b5574c *__obf_b33e9e2851d3e569) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_6fd3f72eb9b5574c.__obf_32c2c59b26b6e25b(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
}

type __obf_009fa1223393ec9a struct {
	__obf_32c2c59b26b6e25b EncoderFunc
	__obf_0e7cdc7b54a85fff func(__obf_2a1474febb02279b unsafe.Pointer) bool
}

func (__obf_96e65a4c8c4f2ce5 *__obf_009fa1223393ec9a) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_96e65a4c8c4f2ce5.__obf_32c2c59b26b6e25b(__obf_2a1474febb02279b, __obf_850a7457bb739a32)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_009fa1223393ec9a) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	if __obf_96e65a4c8c4f2ce5.__obf_0e7cdc7b54a85fff == nil {
		return false
	}
	return __obf_96e65a4c8c4f2ce5.__obf_0e7cdc7b54a85fff(__obf_2a1474febb02279b)
}

// DecoderFunc the function form of TypeDecoder
type DecoderFunc func(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator)

// EncoderFunc the function form of TypeEncoder
type EncoderFunc func(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream)

// RegisterTypeDecoderFunc register TypeDecoder for a type with function
func RegisterTypeDecoderFunc(__obf_29ebd0f2c324f5ea string, __obf_32c2c59b26b6e25b DecoderFunc) {
	__obf_c0b846a36a3844aa[__obf_29ebd0f2c324f5ea] = &__obf_b33e9e2851d3e569{__obf_32c2c59b26b6e25b}
}

// RegisterTypeDecoder register TypeDecoder for a typ
func RegisterTypeDecoder(__obf_29ebd0f2c324f5ea string, __obf_6fd3f72eb9b5574c ValDecoder) {
	__obf_c0b846a36a3844aa[__obf_29ebd0f2c324f5ea] = __obf_6fd3f72eb9b5574c
}

// RegisterFieldDecoderFunc register TypeDecoder for a struct field with function
func RegisterFieldDecoderFunc(__obf_29ebd0f2c324f5ea string, __obf_7e01b5b4651074d4 string, __obf_32c2c59b26b6e25b DecoderFunc) {
	RegisterFieldDecoder(__obf_29ebd0f2c324f5ea, __obf_7e01b5b4651074d4, &__obf_b33e9e2851d3e569{__obf_32c2c59b26b6e25b})
}

// RegisterFieldDecoder register TypeDecoder for a struct field
func RegisterFieldDecoder(__obf_29ebd0f2c324f5ea string, __obf_7e01b5b4651074d4 string, __obf_6fd3f72eb9b5574c ValDecoder) {
	__obf_c7a2a8e5abb2ff8a[fmt.Sprintf("%s/%s", __obf_29ebd0f2c324f5ea, __obf_7e01b5b4651074d4)] = __obf_6fd3f72eb9b5574c
}

// RegisterTypeEncoderFunc register TypeEncoder for a type with encode/isEmpty function
func RegisterTypeEncoderFunc(__obf_29ebd0f2c324f5ea string, __obf_32c2c59b26b6e25b EncoderFunc, __obf_0e7cdc7b54a85fff func(unsafe.Pointer) bool) {
	__obf_ad65186839258b68[__obf_29ebd0f2c324f5ea] = &__obf_009fa1223393ec9a{__obf_32c2c59b26b6e25b,

		// RegisterTypeEncoder register TypeEncoder for a type
		__obf_0e7cdc7b54a85fff}
}

func RegisterTypeEncoder(__obf_29ebd0f2c324f5ea string, __obf_96e65a4c8c4f2ce5 ValEncoder) {
	__obf_ad65186839258b68[__obf_29ebd0f2c324f5ea] = __obf_96e65a4c8c4f2ce5
}

// RegisterFieldEncoderFunc register TypeEncoder for a struct field with encode/isEmpty function
func RegisterFieldEncoderFunc(__obf_29ebd0f2c324f5ea string, __obf_7e01b5b4651074d4 string, __obf_32c2c59b26b6e25b EncoderFunc, __obf_0e7cdc7b54a85fff func(unsafe.Pointer) bool) {
	RegisterFieldEncoder(__obf_29ebd0f2c324f5ea, __obf_7e01b5b4651074d4, &__obf_009fa1223393ec9a{__obf_32c2c59b26b6e25b,

		// RegisterFieldEncoder register TypeEncoder for a struct field
		__obf_0e7cdc7b54a85fff})
}

func RegisterFieldEncoder(__obf_29ebd0f2c324f5ea string, __obf_7e01b5b4651074d4 string, __obf_96e65a4c8c4f2ce5 ValEncoder) {
	__obf_afeecc600805f008[fmt.Sprintf("%s/%s", __obf_29ebd0f2c324f5ea, __obf_7e01b5b4651074d4)] = __obf_96e65a4c8c4f2ce5
}

// RegisterExtension register extension
func RegisterExtension(__obf_8c9d73a8f8319687 Extension) {
	__obf_bb7ace2aaa9423b9 = append(__obf_bb7ace2aaa9423b9, __obf_8c9d73a8f8319687)
}

func __obf_deb887de9e48895a(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_6fd3f72eb9b5574c := _getTypeDecoderFromExtension(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
			__obf_6fd3f72eb9b5574c = __obf_8c9d73a8f8319687.DecorateDecoder(__obf_29ebd0f2c324f5ea, __obf_6fd3f72eb9b5574c)
		}
		__obf_6fd3f72eb9b5574c = __obf_2f9c5aed866cce75.__obf_920e2f1ddf47b5e1.DecorateDecoder(__obf_29ebd0f2c324f5ea, __obf_6fd3f72eb9b5574c)
		for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
			__obf_6fd3f72eb9b5574c = __obf_8c9d73a8f8319687.DecorateDecoder(__obf_29ebd0f2c324f5ea, __obf_6fd3f72eb9b5574c)
		}
	}
	return __obf_6fd3f72eb9b5574c
}
func _getTypeDecoderFromExtension(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
		__obf_6fd3f72eb9b5574c := __obf_8c9d73a8f8319687.CreateDecoder(__obf_29ebd0f2c324f5ea)
		if __obf_6fd3f72eb9b5574c != nil {
			return __obf_6fd3f72eb9b5574c
		}
	}
	__obf_6fd3f72eb9b5574c := __obf_2f9c5aed866cce75.__obf_920e2f1ddf47b5e1.CreateDecoder(__obf_29ebd0f2c324f5ea)
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_6fd3f72eb9b5574c := __obf_8c9d73a8f8319687.CreateDecoder(__obf_29ebd0f2c324f5ea)
		if __obf_6fd3f72eb9b5574c != nil {
			return __obf_6fd3f72eb9b5574c
		}
	}
	__obf_d1d6ae4b00f1c128 := __obf_29ebd0f2c324f5ea.String()
	__obf_6fd3f72eb9b5574c = __obf_c0b846a36a3844aa[__obf_d1d6ae4b00f1c128]
	if __obf_6fd3f72eb9b5574c != nil {
		return __obf_6fd3f72eb9b5574c
	}
	if __obf_29ebd0f2c324f5ea.Kind() == reflect.Ptr {
		__obf_f2fdafeb141957bd := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafePtrType)
		__obf_6fd3f72eb9b5574c := __obf_c0b846a36a3844aa[__obf_f2fdafeb141957bd.Elem().String()]
		if __obf_6fd3f72eb9b5574c != nil {
			return &OptionalDecoder{__obf_f2fdafeb141957bd.Elem(), __obf_6fd3f72eb9b5574c}
		}
	}
	return nil
}

func __obf_ee643627fa1cfd77(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	__obf_96e65a4c8c4f2ce5 := _getTypeEncoderFromExtension(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
			__obf_96e65a4c8c4f2ce5 = __obf_8c9d73a8f8319687.DecorateEncoder(__obf_29ebd0f2c324f5ea, __obf_96e65a4c8c4f2ce5)
		}
		__obf_96e65a4c8c4f2ce5 = __obf_2f9c5aed866cce75.__obf_47929f7efe51b371.DecorateEncoder(__obf_29ebd0f2c324f5ea, __obf_96e65a4c8c4f2ce5)
		for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
			__obf_96e65a4c8c4f2ce5 = __obf_8c9d73a8f8319687.DecorateEncoder(__obf_29ebd0f2c324f5ea, __obf_96e65a4c8c4f2ce5)
		}
	}
	return __obf_96e65a4c8c4f2ce5
}

func _getTypeEncoderFromExtension(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
		__obf_96e65a4c8c4f2ce5 := __obf_8c9d73a8f8319687.CreateEncoder(__obf_29ebd0f2c324f5ea)
		if __obf_96e65a4c8c4f2ce5 != nil {
			return __obf_96e65a4c8c4f2ce5
		}
	}
	__obf_96e65a4c8c4f2ce5 := __obf_2f9c5aed866cce75.__obf_47929f7efe51b371.CreateEncoder(__obf_29ebd0f2c324f5ea)
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_96e65a4c8c4f2ce5 := __obf_8c9d73a8f8319687.CreateEncoder(__obf_29ebd0f2c324f5ea)
		if __obf_96e65a4c8c4f2ce5 != nil {
			return __obf_96e65a4c8c4f2ce5
		}
	}
	__obf_d1d6ae4b00f1c128 := __obf_29ebd0f2c324f5ea.String()
	__obf_96e65a4c8c4f2ce5 = __obf_ad65186839258b68[__obf_d1d6ae4b00f1c128]
	if __obf_96e65a4c8c4f2ce5 != nil {
		return __obf_96e65a4c8c4f2ce5
	}
	if __obf_29ebd0f2c324f5ea.Kind() == reflect.Ptr {
		__obf_46f821b04a80a108 := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafePtrType)
		__obf_96e65a4c8c4f2ce5 := __obf_ad65186839258b68[__obf_46f821b04a80a108.Elem().String()]
		if __obf_96e65a4c8c4f2ce5 != nil {
			return &OptionalEncoder{__obf_96e65a4c8c4f2ce5}
		}
	}
	return nil
}

func __obf_d2a025550d51a745(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) *StructDescriptor {
	__obf_77c4e788739dc65f := __obf_29ebd0f2c324f5ea.(*reflect2.UnsafeStructType)
	__obf_a306972ee06178a1 := []*Binding{}
	__obf_884ece42b376f6b2 := []*Binding{}
	for __obf_5aa5c8829b97f182 := 0; __obf_5aa5c8829b97f182 < __obf_77c4e788739dc65f.NumField(); __obf_5aa5c8829b97f182++ {
		__obf_7e01b5b4651074d4 := __obf_77c4e788739dc65f.Field(__obf_5aa5c8829b97f182)
		__obf_d177960a22cc0ad2, __obf_431572323a9b58d7 := __obf_7e01b5b4651074d4.Tag().Lookup(__obf_2f9c5aed866cce75.__obf_f6ba056498acdb88())
		if __obf_2f9c5aed866cce75.__obf_dc0073de34b8604d && !__obf_431572323a9b58d7 && !__obf_7e01b5b4651074d4.Anonymous() {
			continue
		}
		if __obf_d177960a22cc0ad2 == "-" || __obf_7e01b5b4651074d4.Name() == "_" {
			continue
		}
		__obf_bac5d25a755820f9 := strings.Split(__obf_d177960a22cc0ad2, ",")
		if __obf_7e01b5b4651074d4.Anonymous() && (__obf_d177960a22cc0ad2 == "" || __obf_bac5d25a755820f9[0] == "") {
			if __obf_7e01b5b4651074d4.Type().Kind() == reflect.Struct {
				__obf_e9a2f5e72a0f0289 := __obf_d2a025550d51a745(__obf_2f9c5aed866cce75, __obf_7e01b5b4651074d4.Type())
				for _, __obf_9a21bdc8399ba183 := range __obf_e9a2f5e72a0f0289.Fields {
					__obf_9a21bdc8399ba183.__obf_82efe19241f5de3d = append([]int{__obf_5aa5c8829b97f182}, __obf_9a21bdc8399ba183.__obf_82efe19241f5de3d...)
					__obf_75f6fd68d932045a := __obf_9a21bdc8399ba183.Encoder.(*__obf_65fe7282c99999d2).__obf_75f6fd68d932045a
					__obf_9a21bdc8399ba183.
						Encoder = &__obf_65fe7282c99999d2{__obf_7e01b5b4651074d4, __obf_9a21bdc8399ba183.Encoder, __obf_75f6fd68d932045a}
					__obf_9a21bdc8399ba183.
						Decoder = &__obf_e06930e4e8a0db9e{__obf_7e01b5b4651074d4, __obf_9a21bdc8399ba183.Decoder}
					__obf_a306972ee06178a1 = append(__obf_a306972ee06178a1, __obf_9a21bdc8399ba183)
				}
				continue
			} else if __obf_7e01b5b4651074d4.Type().Kind() == reflect.Ptr {
				__obf_f2fdafeb141957bd := __obf_7e01b5b4651074d4.Type().(*reflect2.UnsafePtrType)
				if __obf_f2fdafeb141957bd.Elem().Kind() == reflect.Struct {
					__obf_e9a2f5e72a0f0289 := __obf_d2a025550d51a745(__obf_2f9c5aed866cce75, __obf_f2fdafeb141957bd.Elem())
					for _, __obf_9a21bdc8399ba183 := range __obf_e9a2f5e72a0f0289.Fields {
						__obf_9a21bdc8399ba183.__obf_82efe19241f5de3d = append([]int{__obf_5aa5c8829b97f182}, __obf_9a21bdc8399ba183.__obf_82efe19241f5de3d...)
						__obf_75f6fd68d932045a := __obf_9a21bdc8399ba183.Encoder.(*__obf_65fe7282c99999d2).__obf_75f6fd68d932045a
						__obf_9a21bdc8399ba183.
							Encoder = &__obf_9146060c74844107{__obf_9a21bdc8399ba183.Encoder}
						__obf_9a21bdc8399ba183.
							Encoder = &__obf_65fe7282c99999d2{__obf_7e01b5b4651074d4, __obf_9a21bdc8399ba183.Encoder, __obf_75f6fd68d932045a}
						__obf_9a21bdc8399ba183.
							Decoder = &__obf_7a31d659954d3208{__obf_f2fdafeb141957bd.Elem(), __obf_9a21bdc8399ba183.Decoder}
						__obf_9a21bdc8399ba183.
							Decoder = &__obf_e06930e4e8a0db9e{__obf_7e01b5b4651074d4, __obf_9a21bdc8399ba183.Decoder}
						__obf_a306972ee06178a1 = append(__obf_a306972ee06178a1, __obf_9a21bdc8399ba183)
					}
					continue
				}
			}
		}
		__obf_313816ce86ba265c := __obf_3aa2914a500bdce4(__obf_7e01b5b4651074d4.Name(), __obf_bac5d25a755820f9[0], __obf_d177960a22cc0ad2)
		__obf_1de880dda3a5e321 := fmt.Sprintf("%s/%s", __obf_29ebd0f2c324f5ea.String(), __obf_7e01b5b4651074d4.Name())
		__obf_6fd3f72eb9b5574c := __obf_c7a2a8e5abb2ff8a[__obf_1de880dda3a5e321]
		if __obf_6fd3f72eb9b5574c == nil {
			__obf_6fd3f72eb9b5574c = __obf_0b44a7afc1523314(__obf_2f9c5aed866cce75.append(__obf_7e01b5b4651074d4.Name()), __obf_7e01b5b4651074d4.Type())
		}
		__obf_96e65a4c8c4f2ce5 := __obf_afeecc600805f008[__obf_1de880dda3a5e321]
		if __obf_96e65a4c8c4f2ce5 == nil {
			__obf_96e65a4c8c4f2ce5 = __obf_d1233db7a73cc96c(__obf_2f9c5aed866cce75.append(__obf_7e01b5b4651074d4.Name()), __obf_7e01b5b4651074d4.Type())
		}
		__obf_9a21bdc8399ba183 := &Binding{
			Field:     __obf_7e01b5b4651074d4,
			FromNames: __obf_313816ce86ba265c,
			ToNames:   __obf_313816ce86ba265c,
			Decoder:   __obf_6fd3f72eb9b5574c,
			Encoder:   __obf_96e65a4c8c4f2ce5,
		}
		__obf_9a21bdc8399ba183.__obf_82efe19241f5de3d = []int{__obf_5aa5c8829b97f182}
		__obf_884ece42b376f6b2 = append(__obf_884ece42b376f6b2, __obf_9a21bdc8399ba183)
	}
	return __obf_2bff893873cf632f(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea, __obf_884ece42b376f6b2, __obf_a306972ee06178a1)
}
func __obf_2bff893873cf632f(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type, __obf_884ece42b376f6b2 []*Binding, __obf_a306972ee06178a1 []*Binding) *StructDescriptor {
	__obf_e9a2f5e72a0f0289 := &StructDescriptor{
		Type:   __obf_29ebd0f2c324f5ea,
		Fields: __obf_884ece42b376f6b2,
	}
	for _, __obf_8c9d73a8f8319687 := range __obf_bb7ace2aaa9423b9 {
		__obf_8c9d73a8f8319687.
			UpdateStructDescriptor(__obf_e9a2f5e72a0f0289)
	}
	__obf_2f9c5aed866cce75.__obf_47929f7efe51b371.
		UpdateStructDescriptor(__obf_e9a2f5e72a0f0289)
	__obf_2f9c5aed866cce75.__obf_920e2f1ddf47b5e1.
		UpdateStructDescriptor(__obf_e9a2f5e72a0f0289)
	for _, __obf_8c9d73a8f8319687 := range __obf_2f9c5aed866cce75.__obf_b4dfae156c7ac26c {
		__obf_8c9d73a8f8319687.
			UpdateStructDescriptor(__obf_e9a2f5e72a0f0289)
	}
	__obf_683148925bdd955f(__obf_e9a2f5e72a0f0289, __obf_2f9c5aed866cce75.
		// merge normal & embedded bindings & sort with original order
		__obf_972c2293804d141c)
	__obf_0de78973ec45494d := __obf_5aaf07df93fd5365(append(__obf_a306972ee06178a1, __obf_e9a2f5e72a0f0289.Fields...))
	sort.Sort(__obf_0de78973ec45494d)
	__obf_e9a2f5e72a0f0289.
		Fields = __obf_0de78973ec45494d
	return __obf_e9a2f5e72a0f0289
}

type __obf_5aaf07df93fd5365 []*Binding

func (__obf_884ece42b376f6b2 __obf_5aaf07df93fd5365) Len() int {
	return len(__obf_884ece42b376f6b2)
}

func (__obf_884ece42b376f6b2 __obf_5aaf07df93fd5365) Less(__obf_5aa5c8829b97f182, __obf_7b29092764c6c9cb int) bool {
	__obf_43fd3c3e07e3bf9f := __obf_884ece42b376f6b2[__obf_5aa5c8829b97f182].__obf_82efe19241f5de3d
	__obf_db0c70142847728b := __obf_884ece42b376f6b2[__obf_7b29092764c6c9cb].__obf_82efe19241f5de3d
	__obf_4b23cd971318cca4 := 0
	for {
		if __obf_43fd3c3e07e3bf9f[__obf_4b23cd971318cca4] < __obf_db0c70142847728b[__obf_4b23cd971318cca4] {
			return true
		} else if __obf_43fd3c3e07e3bf9f[__obf_4b23cd971318cca4] > __obf_db0c70142847728b[__obf_4b23cd971318cca4] {
			return false
		}
		__obf_4b23cd971318cca4++
	}
}

func (__obf_884ece42b376f6b2 __obf_5aaf07df93fd5365) Swap(__obf_5aa5c8829b97f182, __obf_7b29092764c6c9cb int) {
	__obf_884ece42b376f6b2[__obf_5aa5c8829b97f182], __obf_884ece42b376f6b2[__obf_7b29092764c6c9cb] = __obf_884ece42b376f6b2[__obf_7b29092764c6c9cb], __obf_884ece42b376f6b2[__obf_5aa5c8829b97f182]
}

func __obf_683148925bdd955f(__obf_e9a2f5e72a0f0289 *StructDescriptor, __obf_892632c77e859037 *__obf_972c2293804d141c) {
	for _, __obf_9a21bdc8399ba183 := range __obf_e9a2f5e72a0f0289.Fields {
		__obf_e38106386a5f4d63 := false
		__obf_bac5d25a755820f9 := strings.Split(__obf_9a21bdc8399ba183.Field.Tag().Get(__obf_892632c77e859037.__obf_f6ba056498acdb88()), ",")
		for _, __obf_01ec78d146758040 := range __obf_bac5d25a755820f9[1:] {
			if __obf_01ec78d146758040 == "omitempty" {
				__obf_e38106386a5f4d63 = true
			} else if __obf_01ec78d146758040 == "string" {
				if __obf_9a21bdc8399ba183.Field.Type().Kind() == reflect.String {
					__obf_9a21bdc8399ba183.
						Decoder = &__obf_cd171cb12faf3492{__obf_9a21bdc8399ba183.Decoder, __obf_892632c77e859037}
					__obf_9a21bdc8399ba183.
						Encoder = &__obf_5f33847749b86fb1{__obf_9a21bdc8399ba183.Encoder, __obf_892632c77e859037}
				} else {
					__obf_9a21bdc8399ba183.
						Decoder = &__obf_246bb5830b114872{__obf_9a21bdc8399ba183.Decoder}
					__obf_9a21bdc8399ba183.
						Encoder = &__obf_1b71667627f7b8fe{__obf_9a21bdc8399ba183.Encoder}
				}
			}
		}
		__obf_9a21bdc8399ba183.
			Decoder = &__obf_e06930e4e8a0db9e{__obf_9a21bdc8399ba183.Field, __obf_9a21bdc8399ba183.Decoder}
		__obf_9a21bdc8399ba183.
			Encoder = &__obf_65fe7282c99999d2{__obf_9a21bdc8399ba183.Field, __obf_9a21bdc8399ba183.Encoder, __obf_e38106386a5f4d63}
	}
}

func __obf_3aa2914a500bdce4(__obf_ca4d4f4e20107e26 string, __obf_4d805ab05a124335 string, __obf_d1dd7f8cbf0e9cc3 string) []string {
	// ignore?
	if __obf_d1dd7f8cbf0e9cc3 == "-" {
		return []string{}
	}
	// rename?
	var __obf_313816ce86ba265c []string
	if __obf_4d805ab05a124335 == "" {
		__obf_313816ce86ba265c = []string{__obf_ca4d4f4e20107e26}
	} else {
		__obf_313816ce86ba265c = []string{__obf_4d805ab05a124335}
	}
	__obf_872e36c321b29161 := // private?
		unicode.IsLower(rune(__obf_ca4d4f4e20107e26[0])) || __obf_ca4d4f4e20107e26[0] == '_'
	if __obf_872e36c321b29161 {
		__obf_313816ce86ba265c = []string{}
	}
	return __obf_313816ce86ba265c
}
