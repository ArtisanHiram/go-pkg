package __obf_030d39f7a8de96c6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
	"sort"
	"strings"
	"unicode"
	"unsafe"
)

var __obf_af43d5efcf6ae11d = map[string]ValDecoder{}
var __obf_3b686c29105543bc = map[string]ValDecoder{}
var __obf_8f43affe58f43115 = map[string]ValEncoder{}
var __obf_133872148688e4c1 = map[string]ValEncoder{}
var __obf_04498d7bbd7ab2ba = []Extension{}

// StructDescriptor describe how should we encode/decode the struct
type StructDescriptor struct {
	Type   reflect2.Type
	Fields []*Binding
}

// GetField get one field from the descriptor by its name.
// Can not use map here to keep field orders.
func (__obf_ef1b847f404bee13 *StructDescriptor) GetField(__obf_07e0e32c6be0db36 string) *Binding {
	for _, __obf_04ea35bce3978e8f := range __obf_ef1b847f404bee13.Fields {
		if __obf_04ea35bce3978e8f.Field.Name() == __obf_07e0e32c6be0db36 {
			return __obf_04ea35bce3978e8f
		}
	}
	return nil
}

// Binding describe how should we encode/decode the struct field
type Binding struct {
	__obf_9cdf9b6aa25d4f68 []int
	Field                  reflect2.StructField
	FromNames              []string
	ToNames                []string
	Encoder                ValEncoder
	Decoder                ValDecoder
}

// Extension the one for all SPI. Customize encoding/decoding by specifying alternate encoder/decoder.
// Can also rename fields by UpdateStructDescriptor.
type Extension interface {
	UpdateStructDescriptor(__obf_ef1b847f404bee13 *StructDescriptor)
	CreateMapKeyDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder
	CreateMapKeyEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder
	CreateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder
	CreateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder
	DecorateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_11a3f28116164d09 ValDecoder) ValDecoder
	DecorateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_008f61596d7e9523 ValEncoder) ValEncoder
}

// DummyExtension embed this type get dummy implementation for all methods of Extension
type DummyExtension struct {
}

// UpdateStructDescriptor No-op
func (__obf_621544a57e52000f *DummyExtension) UpdateStructDescriptor(__obf_ef1b847f404bee13 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_621544a57e52000f *DummyExtension) CreateMapKeyDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_621544a57e52000f *DummyExtension) CreateMapKeyEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder No-op
func (__obf_621544a57e52000f *DummyExtension) CreateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder No-op
func (__obf_621544a57e52000f *DummyExtension) CreateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_621544a57e52000f *DummyExtension) DecorateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_11a3f28116164d09 ValDecoder) ValDecoder {
	return __obf_11a3f28116164d09
}

// DecorateEncoder No-op
func (__obf_621544a57e52000f *DummyExtension) DecorateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_008f61596d7e9523 ValEncoder) ValEncoder {
	return __obf_008f61596d7e9523
}

type EncoderExtension map[reflect2.Type]ValEncoder

// UpdateStructDescriptor No-op
func (__obf_621544a57e52000f EncoderExtension) UpdateStructDescriptor(__obf_ef1b847f404bee13 *StructDescriptor) {
}

// CreateDecoder No-op
func (__obf_621544a57e52000f EncoderExtension) CreateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	return nil
}

// CreateEncoder get encoder from map
func (__obf_621544a57e52000f EncoderExtension) CreateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	return __obf_621544a57e52000f[__obf_8744d0b8c80ccdc1]
}

// CreateMapKeyDecoder No-op
func (__obf_621544a57e52000f EncoderExtension) CreateMapKeyDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_621544a57e52000f EncoderExtension) CreateMapKeyEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_621544a57e52000f EncoderExtension) DecorateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_11a3f28116164d09 ValDecoder) ValDecoder {
	return __obf_11a3f28116164d09
}

// DecorateEncoder No-op
func (__obf_621544a57e52000f EncoderExtension) DecorateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_008f61596d7e9523 ValEncoder) ValEncoder {
	return __obf_008f61596d7e9523
}

type DecoderExtension map[reflect2.Type]ValDecoder

// UpdateStructDescriptor No-op
func (__obf_621544a57e52000f DecoderExtension) UpdateStructDescriptor(__obf_ef1b847f404bee13 *StructDescriptor) {
}

// CreateMapKeyDecoder No-op
func (__obf_621544a57e52000f DecoderExtension) CreateMapKeyDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	return nil
}

// CreateMapKeyEncoder No-op
func (__obf_621544a57e52000f DecoderExtension) CreateMapKeyEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	return nil
}

// CreateDecoder get decoder from map
func (__obf_621544a57e52000f DecoderExtension) CreateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	return __obf_621544a57e52000f[__obf_8744d0b8c80ccdc1]
}

// CreateEncoder No-op
func (__obf_621544a57e52000f DecoderExtension) CreateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	return nil
}

// DecorateDecoder No-op
func (__obf_621544a57e52000f DecoderExtension) DecorateDecoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_11a3f28116164d09 ValDecoder) ValDecoder {
	return __obf_11a3f28116164d09
}

// DecorateEncoder No-op
func (__obf_621544a57e52000f DecoderExtension) DecorateEncoder(__obf_8744d0b8c80ccdc1 reflect2.Type, __obf_008f61596d7e9523 ValEncoder) ValEncoder {
	return __obf_008f61596d7e9523
}

type __obf_3f78bd1ffd8b466d struct {
	__obf_20e59be43c1dcfec DecoderFunc
}

func (__obf_11a3f28116164d09 *__obf_3f78bd1ffd8b466d) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	__obf_11a3f28116164d09.__obf_20e59be43c1dcfec(__obf_dbbf371b91136494, __obf_4ab56a99965952e7)
}

type __obf_f97fbd765d3a66b4 struct {
	__obf_20e59be43c1dcfec EncoderFunc
	__obf_2f4b87f1ba276038 func(__obf_dbbf371b91136494 unsafe.Pointer) bool
}

func (__obf_008f61596d7e9523 *__obf_f97fbd765d3a66b4) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_008f61596d7e9523.__obf_20e59be43c1dcfec(__obf_dbbf371b91136494, __obf_8a2c51fe22775655)
}

func (__obf_008f61596d7e9523 *__obf_f97fbd765d3a66b4) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	if __obf_008f61596d7e9523.__obf_2f4b87f1ba276038 == nil {
		return false
	}
	return __obf_008f61596d7e9523.__obf_2f4b87f1ba276038(__obf_dbbf371b91136494)
}

// DecoderFunc the function form of TypeDecoder
type DecoderFunc func(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator)

// EncoderFunc the function form of TypeEncoder
type EncoderFunc func(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream)

// RegisterTypeDecoderFunc register TypeDecoder for a type with function
func RegisterTypeDecoderFunc(__obf_8744d0b8c80ccdc1 string, __obf_20e59be43c1dcfec DecoderFunc) {
	__obf_af43d5efcf6ae11d[__obf_8744d0b8c80ccdc1] = &__obf_3f78bd1ffd8b466d{__obf_20e59be43c1dcfec}
}

// RegisterTypeDecoder register TypeDecoder for a typ
func RegisterTypeDecoder(__obf_8744d0b8c80ccdc1 string, __obf_11a3f28116164d09 ValDecoder) {
	__obf_af43d5efcf6ae11d[__obf_8744d0b8c80ccdc1] = __obf_11a3f28116164d09
}

// RegisterFieldDecoderFunc register TypeDecoder for a struct field with function
func RegisterFieldDecoderFunc(__obf_8744d0b8c80ccdc1 string, __obf_cd4d02f264b18d55 string, __obf_20e59be43c1dcfec DecoderFunc) {
	RegisterFieldDecoder(__obf_8744d0b8c80ccdc1, __obf_cd4d02f264b18d55, &__obf_3f78bd1ffd8b466d{__obf_20e59be43c1dcfec})
}

// RegisterFieldDecoder register TypeDecoder for a struct field
func RegisterFieldDecoder(__obf_8744d0b8c80ccdc1 string, __obf_cd4d02f264b18d55 string, __obf_11a3f28116164d09 ValDecoder) {
	__obf_3b686c29105543bc[fmt.Sprintf("%s/%s", __obf_8744d0b8c80ccdc1, __obf_cd4d02f264b18d55)] = __obf_11a3f28116164d09
}

// RegisterTypeEncoderFunc register TypeEncoder for a type with encode/isEmpty function
func RegisterTypeEncoderFunc(__obf_8744d0b8c80ccdc1 string, __obf_20e59be43c1dcfec EncoderFunc, __obf_2f4b87f1ba276038 func(unsafe.Pointer) bool) {
	__obf_8f43affe58f43115[__obf_8744d0b8c80ccdc1] = &__obf_f97fbd765d3a66b4{__obf_20e59be43c1dcfec,

		// RegisterTypeEncoder register TypeEncoder for a type
		__obf_2f4b87f1ba276038}
}

func RegisterTypeEncoder(__obf_8744d0b8c80ccdc1 string, __obf_008f61596d7e9523 ValEncoder) {
	__obf_8f43affe58f43115[__obf_8744d0b8c80ccdc1] = __obf_008f61596d7e9523
}

// RegisterFieldEncoderFunc register TypeEncoder for a struct field with encode/isEmpty function
func RegisterFieldEncoderFunc(__obf_8744d0b8c80ccdc1 string, __obf_cd4d02f264b18d55 string, __obf_20e59be43c1dcfec EncoderFunc, __obf_2f4b87f1ba276038 func(unsafe.Pointer) bool) {
	RegisterFieldEncoder(__obf_8744d0b8c80ccdc1, __obf_cd4d02f264b18d55, &__obf_f97fbd765d3a66b4{__obf_20e59be43c1dcfec,

		// RegisterFieldEncoder register TypeEncoder for a struct field
		__obf_2f4b87f1ba276038})
}

func RegisterFieldEncoder(__obf_8744d0b8c80ccdc1 string, __obf_cd4d02f264b18d55 string, __obf_008f61596d7e9523 ValEncoder) {
	__obf_133872148688e4c1[fmt.Sprintf("%s/%s", __obf_8744d0b8c80ccdc1, __obf_cd4d02f264b18d55)] = __obf_008f61596d7e9523
}

// RegisterExtension register extension
func RegisterExtension(__obf_621544a57e52000f Extension) {
	__obf_04498d7bbd7ab2ba = append(__obf_04498d7bbd7ab2ba, __obf_621544a57e52000f)
}

func __obf_8a762691a84adf6f(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	__obf_11a3f28116164d09 := _getTypeDecoderFromExtension(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
			__obf_11a3f28116164d09 = __obf_621544a57e52000f.DecorateDecoder(__obf_8744d0b8c80ccdc1, __obf_11a3f28116164d09)
		}
		__obf_11a3f28116164d09 = __obf_a565fbaffca944f0.__obf_6313bfb9c913bd50.DecorateDecoder(__obf_8744d0b8c80ccdc1, __obf_11a3f28116164d09)
		for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
			__obf_11a3f28116164d09 = __obf_621544a57e52000f.DecorateDecoder(__obf_8744d0b8c80ccdc1, __obf_11a3f28116164d09)
		}
	}
	return __obf_11a3f28116164d09
}
func _getTypeDecoderFromExtension(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
		__obf_11a3f28116164d09 := __obf_621544a57e52000f.CreateDecoder(__obf_8744d0b8c80ccdc1)
		if __obf_11a3f28116164d09 != nil {
			return __obf_11a3f28116164d09
		}
	}
	__obf_11a3f28116164d09 := __obf_a565fbaffca944f0.__obf_6313bfb9c913bd50.CreateDecoder(__obf_8744d0b8c80ccdc1)
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_11a3f28116164d09 := __obf_621544a57e52000f.CreateDecoder(__obf_8744d0b8c80ccdc1)
		if __obf_11a3f28116164d09 != nil {
			return __obf_11a3f28116164d09
		}
	}
	__obf_aa66e764b9a50b6b := __obf_8744d0b8c80ccdc1.String()
	__obf_11a3f28116164d09 = __obf_af43d5efcf6ae11d[__obf_aa66e764b9a50b6b]
	if __obf_11a3f28116164d09 != nil {
		return __obf_11a3f28116164d09
	}
	if __obf_8744d0b8c80ccdc1.Kind() == reflect.Ptr {
		__obf_3a51157620f9910b := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafePtrType)
		__obf_11a3f28116164d09 := __obf_af43d5efcf6ae11d[__obf_3a51157620f9910b.Elem().String()]
		if __obf_11a3f28116164d09 != nil {
			return &OptionalDecoder{__obf_3a51157620f9910b.Elem(), __obf_11a3f28116164d09}
		}
	}
	return nil
}

func __obf_4f7cd8fb242103cc(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	__obf_008f61596d7e9523 := _getTypeEncoderFromExtension(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
			__obf_008f61596d7e9523 = __obf_621544a57e52000f.DecorateEncoder(__obf_8744d0b8c80ccdc1, __obf_008f61596d7e9523)
		}
		__obf_008f61596d7e9523 = __obf_a565fbaffca944f0.__obf_7db1f0a55b319e45.DecorateEncoder(__obf_8744d0b8c80ccdc1, __obf_008f61596d7e9523)
		for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
			__obf_008f61596d7e9523 = __obf_621544a57e52000f.DecorateEncoder(__obf_8744d0b8c80ccdc1, __obf_008f61596d7e9523)
		}
	}
	return __obf_008f61596d7e9523
}

func _getTypeEncoderFromExtension(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
		__obf_008f61596d7e9523 := __obf_621544a57e52000f.CreateEncoder(__obf_8744d0b8c80ccdc1)
		if __obf_008f61596d7e9523 != nil {
			return __obf_008f61596d7e9523
		}
	}
	__obf_008f61596d7e9523 := __obf_a565fbaffca944f0.__obf_7db1f0a55b319e45.CreateEncoder(__obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_008f61596d7e9523 := __obf_621544a57e52000f.CreateEncoder(__obf_8744d0b8c80ccdc1)
		if __obf_008f61596d7e9523 != nil {
			return __obf_008f61596d7e9523
		}
	}
	__obf_aa66e764b9a50b6b := __obf_8744d0b8c80ccdc1.String()
	__obf_008f61596d7e9523 = __obf_8f43affe58f43115[__obf_aa66e764b9a50b6b]
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	if __obf_8744d0b8c80ccdc1.Kind() == reflect.Ptr {
		__obf_c62cc9ccdf2bbe92 := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafePtrType)
		__obf_008f61596d7e9523 := __obf_8f43affe58f43115[__obf_c62cc9ccdf2bbe92.Elem().String()]
		if __obf_008f61596d7e9523 != nil {
			return &OptionalEncoder{__obf_008f61596d7e9523}
		}
	}
	return nil
}

func __obf_1219e3aba7e47ed0(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) *StructDescriptor {
	__obf_1baf30032abd054b := __obf_8744d0b8c80ccdc1.(*reflect2.UnsafeStructType)
	__obf_468dea9924558326 := []*Binding{}
	__obf_2a0baf04ebd1bcd4 := []*Binding{}
	for __obf_82c6e05b9d226c58 := 0; __obf_82c6e05b9d226c58 < __obf_1baf30032abd054b.NumField(); __obf_82c6e05b9d226c58++ {
		__obf_cd4d02f264b18d55 := __obf_1baf30032abd054b.Field(__obf_82c6e05b9d226c58)
		__obf_9ba49f7bc2bc210e, __obf_83459ec5ee6a818d := __obf_cd4d02f264b18d55.Tag().Lookup(__obf_a565fbaffca944f0.__obf_79fd0a7f8863c590())
		if __obf_a565fbaffca944f0.__obf_67b8a9dee101ecbf && !__obf_83459ec5ee6a818d && !__obf_cd4d02f264b18d55.Anonymous() {
			continue
		}
		if __obf_9ba49f7bc2bc210e == "-" || __obf_cd4d02f264b18d55.Name() == "_" {
			continue
		}
		__obf_5d3a213f4419d26a := strings.Split(__obf_9ba49f7bc2bc210e, ",")
		if __obf_cd4d02f264b18d55.Anonymous() && (__obf_9ba49f7bc2bc210e == "" || __obf_5d3a213f4419d26a[0] == "") {
			if __obf_cd4d02f264b18d55.Type().Kind() == reflect.Struct {
				__obf_ef1b847f404bee13 := __obf_1219e3aba7e47ed0(__obf_a565fbaffca944f0, __obf_cd4d02f264b18d55.Type())
				for _, __obf_04ea35bce3978e8f := range __obf_ef1b847f404bee13.Fields {
					__obf_04ea35bce3978e8f.__obf_9cdf9b6aa25d4f68 = append([]int{__obf_82c6e05b9d226c58}, __obf_04ea35bce3978e8f.__obf_9cdf9b6aa25d4f68...)
					__obf_ef6aa90a5fa198f2 := __obf_04ea35bce3978e8f.Encoder.(*__obf_06e12c234657ff27).__obf_ef6aa90a5fa198f2
					__obf_04ea35bce3978e8f.
						Encoder = &__obf_06e12c234657ff27{__obf_cd4d02f264b18d55, __obf_04ea35bce3978e8f.Encoder, __obf_ef6aa90a5fa198f2}
					__obf_04ea35bce3978e8f.
						Decoder = &__obf_7198db8bbb81224a{__obf_cd4d02f264b18d55, __obf_04ea35bce3978e8f.Decoder}
					__obf_468dea9924558326 = append(__obf_468dea9924558326, __obf_04ea35bce3978e8f)
				}
				continue
			} else if __obf_cd4d02f264b18d55.Type().Kind() == reflect.Ptr {
				__obf_3a51157620f9910b := __obf_cd4d02f264b18d55.Type().(*reflect2.UnsafePtrType)
				if __obf_3a51157620f9910b.Elem().Kind() == reflect.Struct {
					__obf_ef1b847f404bee13 := __obf_1219e3aba7e47ed0(__obf_a565fbaffca944f0, __obf_3a51157620f9910b.Elem())
					for _, __obf_04ea35bce3978e8f := range __obf_ef1b847f404bee13.Fields {
						__obf_04ea35bce3978e8f.__obf_9cdf9b6aa25d4f68 = append([]int{__obf_82c6e05b9d226c58}, __obf_04ea35bce3978e8f.__obf_9cdf9b6aa25d4f68...)
						__obf_ef6aa90a5fa198f2 := __obf_04ea35bce3978e8f.Encoder.(*__obf_06e12c234657ff27).__obf_ef6aa90a5fa198f2
						__obf_04ea35bce3978e8f.
							Encoder = &__obf_e3ea39e92054d646{__obf_04ea35bce3978e8f.Encoder}
						__obf_04ea35bce3978e8f.
							Encoder = &__obf_06e12c234657ff27{__obf_cd4d02f264b18d55, __obf_04ea35bce3978e8f.Encoder, __obf_ef6aa90a5fa198f2}
						__obf_04ea35bce3978e8f.
							Decoder = &__obf_174b577e33a12aa4{__obf_3a51157620f9910b.Elem(), __obf_04ea35bce3978e8f.Decoder}
						__obf_04ea35bce3978e8f.
							Decoder = &__obf_7198db8bbb81224a{__obf_cd4d02f264b18d55, __obf_04ea35bce3978e8f.Decoder}
						__obf_468dea9924558326 = append(__obf_468dea9924558326, __obf_04ea35bce3978e8f)
					}
					continue
				}
			}
		}
		__obf_7c06f2fa07a5e0ff := __obf_a69f7be6850b9319(__obf_cd4d02f264b18d55.Name(), __obf_5d3a213f4419d26a[0], __obf_9ba49f7bc2bc210e)
		__obf_89a3784f2fe05889 := fmt.Sprintf("%s/%s", __obf_8744d0b8c80ccdc1.String(), __obf_cd4d02f264b18d55.Name())
		__obf_11a3f28116164d09 := __obf_3b686c29105543bc[__obf_89a3784f2fe05889]
		if __obf_11a3f28116164d09 == nil {
			__obf_11a3f28116164d09 = __obf_729ef08c8b8e060e(__obf_a565fbaffca944f0.append(__obf_cd4d02f264b18d55.Name()), __obf_cd4d02f264b18d55.Type())
		}
		__obf_008f61596d7e9523 := __obf_133872148688e4c1[__obf_89a3784f2fe05889]
		if __obf_008f61596d7e9523 == nil {
			__obf_008f61596d7e9523 = __obf_37ef471fa5e40404(__obf_a565fbaffca944f0.append(__obf_cd4d02f264b18d55.Name()), __obf_cd4d02f264b18d55.Type())
		}
		__obf_04ea35bce3978e8f := &Binding{
			Field:     __obf_cd4d02f264b18d55,
			FromNames: __obf_7c06f2fa07a5e0ff,
			ToNames:   __obf_7c06f2fa07a5e0ff,
			Decoder:   __obf_11a3f28116164d09,
			Encoder:   __obf_008f61596d7e9523,
		}
		__obf_04ea35bce3978e8f.__obf_9cdf9b6aa25d4f68 = []int{__obf_82c6e05b9d226c58}
		__obf_2a0baf04ebd1bcd4 = append(__obf_2a0baf04ebd1bcd4, __obf_04ea35bce3978e8f)
	}
	return __obf_72ca34d64a6d272a(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1, __obf_2a0baf04ebd1bcd4, __obf_468dea9924558326)
}
func __obf_72ca34d64a6d272a(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type, __obf_2a0baf04ebd1bcd4 []*Binding, __obf_468dea9924558326 []*Binding) *StructDescriptor {
	__obf_ef1b847f404bee13 := &StructDescriptor{
		Type:   __obf_8744d0b8c80ccdc1,
		Fields: __obf_2a0baf04ebd1bcd4,
	}
	for _, __obf_621544a57e52000f := range __obf_04498d7bbd7ab2ba {
		__obf_621544a57e52000f.
			UpdateStructDescriptor(__obf_ef1b847f404bee13)
	}
	__obf_a565fbaffca944f0.__obf_7db1f0a55b319e45.
		UpdateStructDescriptor(__obf_ef1b847f404bee13)
	__obf_a565fbaffca944f0.__obf_6313bfb9c913bd50.
		UpdateStructDescriptor(__obf_ef1b847f404bee13)
	for _, __obf_621544a57e52000f := range __obf_a565fbaffca944f0.__obf_6621255bc1f80c8e {
		__obf_621544a57e52000f.
			UpdateStructDescriptor(__obf_ef1b847f404bee13)
	}
	__obf_5a01580f041685b9(__obf_ef1b847f404bee13, __obf_a565fbaffca944f0.
		// merge normal & embedded bindings & sort with original order
		__obf_81639538813814ff)
	__obf_3e35b11d9a21733a := __obf_500590c90545b4a1(append(__obf_468dea9924558326, __obf_ef1b847f404bee13.Fields...))
	sort.Sort(__obf_3e35b11d9a21733a)
	__obf_ef1b847f404bee13.
		Fields = __obf_3e35b11d9a21733a
	return __obf_ef1b847f404bee13
}

type __obf_500590c90545b4a1 []*Binding

func (__obf_2a0baf04ebd1bcd4 __obf_500590c90545b4a1) Len() int {
	return len(__obf_2a0baf04ebd1bcd4)
}

func (__obf_2a0baf04ebd1bcd4 __obf_500590c90545b4a1) Less(__obf_82c6e05b9d226c58, __obf_354f1fc1cab1d938 int) bool {
	__obf_cc01539eb61db3e0 := __obf_2a0baf04ebd1bcd4[__obf_82c6e05b9d226c58].__obf_9cdf9b6aa25d4f68
	__obf_5fd3f6c47b4d3fae := __obf_2a0baf04ebd1bcd4[__obf_354f1fc1cab1d938].__obf_9cdf9b6aa25d4f68
	__obf_4c6b7384a71b9c24 := 0
	for {
		if __obf_cc01539eb61db3e0[__obf_4c6b7384a71b9c24] < __obf_5fd3f6c47b4d3fae[__obf_4c6b7384a71b9c24] {
			return true
		} else if __obf_cc01539eb61db3e0[__obf_4c6b7384a71b9c24] > __obf_5fd3f6c47b4d3fae[__obf_4c6b7384a71b9c24] {
			return false
		}
		__obf_4c6b7384a71b9c24++
	}
}

func (__obf_2a0baf04ebd1bcd4 __obf_500590c90545b4a1) Swap(__obf_82c6e05b9d226c58, __obf_354f1fc1cab1d938 int) {
	__obf_2a0baf04ebd1bcd4[__obf_82c6e05b9d226c58], __obf_2a0baf04ebd1bcd4[__obf_354f1fc1cab1d938] = __obf_2a0baf04ebd1bcd4[__obf_354f1fc1cab1d938], __obf_2a0baf04ebd1bcd4[__obf_82c6e05b9d226c58]
}

func __obf_5a01580f041685b9(__obf_ef1b847f404bee13 *StructDescriptor, __obf_679611dc56ff465b *__obf_81639538813814ff) {
	for _, __obf_04ea35bce3978e8f := range __obf_ef1b847f404bee13.Fields {
		__obf_142dfe92147b1381 := false
		__obf_5d3a213f4419d26a := strings.Split(__obf_04ea35bce3978e8f.Field.Tag().Get(__obf_679611dc56ff465b.__obf_79fd0a7f8863c590()), ",")
		for _, __obf_9d6f668cc1cbc206 := range __obf_5d3a213f4419d26a[1:] {
			if __obf_9d6f668cc1cbc206 == "omitempty" {
				__obf_142dfe92147b1381 = true
			} else if __obf_9d6f668cc1cbc206 == "string" {
				if __obf_04ea35bce3978e8f.Field.Type().Kind() == reflect.String {
					__obf_04ea35bce3978e8f.
						Decoder = &__obf_73fb4b35802e3bb0{__obf_04ea35bce3978e8f.Decoder, __obf_679611dc56ff465b}
					__obf_04ea35bce3978e8f.
						Encoder = &__obf_d418a6a275dc78e1{__obf_04ea35bce3978e8f.Encoder, __obf_679611dc56ff465b}
				} else {
					__obf_04ea35bce3978e8f.
						Decoder = &__obf_6fee448b50200908{__obf_04ea35bce3978e8f.Decoder}
					__obf_04ea35bce3978e8f.
						Encoder = &__obf_e8dab09dfce32843{__obf_04ea35bce3978e8f.Encoder}
				}
			}
		}
		__obf_04ea35bce3978e8f.
			Decoder = &__obf_7198db8bbb81224a{__obf_04ea35bce3978e8f.Field, __obf_04ea35bce3978e8f.Decoder}
		__obf_04ea35bce3978e8f.
			Encoder = &__obf_06e12c234657ff27{__obf_04ea35bce3978e8f.Field, __obf_04ea35bce3978e8f.Encoder, __obf_142dfe92147b1381}
	}
}

func __obf_a69f7be6850b9319(__obf_dd1ac6f56ebcf84e string, __obf_114d8492e2366f56 string, __obf_80b22fad8f7d9785 string) []string {
	// ignore?
	if __obf_80b22fad8f7d9785 == "-" {
		return []string{}
	}
	// rename?
	var __obf_7c06f2fa07a5e0ff []string
	if __obf_114d8492e2366f56 == "" {
		__obf_7c06f2fa07a5e0ff = []string{__obf_dd1ac6f56ebcf84e}
	} else {
		__obf_7c06f2fa07a5e0ff = []string{__obf_114d8492e2366f56}
	}
	__obf_f7937d99026394ca := // private?
		unicode.IsLower(rune(__obf_dd1ac6f56ebcf84e[0])) || __obf_dd1ac6f56ebcf84e[0] == '_'
	if __obf_f7937d99026394ca {
		__obf_7c06f2fa07a5e0ff = []string{}
	}
	return __obf_7c06f2fa07a5e0ff
}
