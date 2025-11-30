package __obf_703d23ba520c3295

import (
	"encoding/json"
	"io"
	"reflect"
	"sync"
	"unsafe"

	"github.com/modern-go/concurrent"
	"github.com/modern-go/reflect2"
)

// Config customize how the API should behave.
// The API is created from Config by Froze.
type Config struct {
	IndentionStep                 int
	MarshalFloatWith6Digits       bool
	EscapeHTML                    bool
	SortMapKeys                   bool
	UseNumber                     bool
	DisallowUnknownFields         bool
	TagKey                        string
	OnlyTaggedField               bool
	ValidateJsonRawMessage        bool
	ObjectFieldMustBeSimpleString bool
	CaseSensitive                 bool
}

// API the public interface of this package.
// Primary Marshal and Unmarshal.
type API interface {
	IteratorPool
	StreamPool
	MarshalToString(__obf_9cfb140618a077d4 any) (string, error)
	Marshal(__obf_9cfb140618a077d4 any) ([]byte, error)
	MarshalIndent(__obf_9cfb140618a077d4 any, __obf_c5fd316f3c4457de, __obf_0d119fc055fc9f67 string) ([]byte, error)
	UnmarshalFromString(__obf_44b48c26051f8033 string, __obf_9cfb140618a077d4 any) error
	Unmarshal(__obf_c180417ee204f8c5 []byte, __obf_9cfb140618a077d4 any) error
	Get(__obf_c180417ee204f8c5 []byte, __obf_267bdf615cb1b310 ...any) Any
	NewEncoder(__obf_8d0bd0f6d68b0c59 io.Writer) *Encoder
	NewDecoder(__obf_3943589abf9d9fb6 io.Reader) *Decoder
	Valid(__obf_c180417ee204f8c5 []byte) bool
	RegisterExtension(__obf_dfc6ab1ee0bfd79e Extension)
	DecoderOf(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder
	EncoderOf(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder
}

// ConfigDefault the default API
var ConfigDefault = Config{
	EscapeHTML: true,
}.Froze()

// ConfigCompatibleWithStandardLibrary tries to be 100% compatible with standard library behavior
var ConfigCompatibleWithStandardLibrary = Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

// ConfigFastest marshals float with only 6 digits precision
var ConfigFastest = Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

type __obf_ef74248d8ae9ba78 struct {
	__obf_4d2d9c0287da7aea Config
	__obf_e5ee8066c2c3cabf bool
	__obf_f5558fa1bf957433 int
	__obf_c8260e5b2e39ae59 bool
	__obf_eb6867c2411c1b8b bool
	__obf_f3aed2fd747dcb8e bool
	__obf_8fbc2f7cca658b1d *concurrent.Map
	__obf_d57e3bec6b3e58ba *concurrent.Map
	__obf_9289f2028bb085f2 Extension
	__obf_7c9972bb545a908d Extension
	__obf_324ee50b8c0b2f3b []Extension
	__obf_27ad6a094b7bcdbf *sync.Pool
	__obf_582fcdbe5914a904 *sync.Pool
	__obf_9332366e75c2e023 bool
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_8f98f7eb87f6bfcf() {
	__obf_25bd4f754a37b862.__obf_8fbc2f7cca658b1d = concurrent.NewMap()
	__obf_25bd4f754a37b862.__obf_d57e3bec6b3e58ba = concurrent.NewMap()
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_4bf6eb1a561eb9a1(__obf_5106f8b01c71afb8 uintptr, __obf_c9719e68325f7d44 ValDecoder) {
	__obf_25bd4f754a37b862.__obf_8fbc2f7cca658b1d.
		Store(__obf_5106f8b01c71afb8, __obf_c9719e68325f7d44)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_722868f5d55df1b8(__obf_5106f8b01c71afb8 uintptr, __obf_cf840243a12ee308 ValEncoder) {
	__obf_25bd4f754a37b862.__obf_d57e3bec6b3e58ba.
		Store(__obf_5106f8b01c71afb8, __obf_cf840243a12ee308)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_9d061f9912e56c90(__obf_5106f8b01c71afb8 uintptr) ValDecoder {
	__obf_c9719e68325f7d44, __obf_87a836f73d31233b := __obf_25bd4f754a37b862.__obf_8fbc2f7cca658b1d.Load(__obf_5106f8b01c71afb8)
	if __obf_87a836f73d31233b {
		return __obf_c9719e68325f7d44.(ValDecoder)
	}
	return nil
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_4295547e93f1b46e(__obf_5106f8b01c71afb8 uintptr) ValEncoder {
	__obf_cf840243a12ee308, __obf_87a836f73d31233b := __obf_25bd4f754a37b862.__obf_d57e3bec6b3e58ba.Load(__obf_5106f8b01c71afb8)
	if __obf_87a836f73d31233b {
		return __obf_cf840243a12ee308.(ValEncoder)
	}
	return nil
}

var __obf_25db02a14117411e = concurrent.NewMap()

func __obf_cbbd424bc3cc37d7(__obf_25bd4f754a37b862 Config) *__obf_ef74248d8ae9ba78 {
	__obf_02ba706f4f104d71, __obf_87a836f73d31233b := __obf_25db02a14117411e.Load(__obf_25bd4f754a37b862)
	if __obf_87a836f73d31233b {
		return __obf_02ba706f4f104d71.(*__obf_ef74248d8ae9ba78)
	}
	return nil
}

func __obf_b88cae1d27d0c0a1(__obf_25bd4f754a37b862 Config, __obf_ef74248d8ae9ba78 *__obf_ef74248d8ae9ba78) {
	__obf_25db02a14117411e.
		Store(__obf_25bd4f754a37b862,

			// Froze forge API from config
			__obf_ef74248d8ae9ba78)
}

func (__obf_25bd4f754a37b862 Config) Froze() API {
	__obf_d4e161deb5b48ae2 := &__obf_ef74248d8ae9ba78{__obf_e5ee8066c2c3cabf: __obf_25bd4f754a37b862.SortMapKeys, __obf_f5558fa1bf957433: __obf_25bd4f754a37b862.IndentionStep, __obf_c8260e5b2e39ae59: __obf_25bd4f754a37b862.ObjectFieldMustBeSimpleString, __obf_eb6867c2411c1b8b: __obf_25bd4f754a37b862.OnlyTaggedField, __obf_f3aed2fd747dcb8e: __obf_25bd4f754a37b862.DisallowUnknownFields, __obf_9332366e75c2e023: __obf_25bd4f754a37b862.CaseSensitive}
	__obf_d4e161deb5b48ae2.__obf_27ad6a094b7bcdbf = &sync.Pool{
		New: func() any {
			return NewStream(__obf_d4e161deb5b48ae2, nil, 512)
		},
	}
	__obf_d4e161deb5b48ae2.__obf_582fcdbe5914a904 = &sync.Pool{
		New: func() any {
			return NewIterator(__obf_d4e161deb5b48ae2)
		},
	}
	__obf_d4e161deb5b48ae2.__obf_8f98f7eb87f6bfcf()
	__obf_9289f2028bb085f2 := EncoderExtension{}
	__obf_7c9972bb545a908d := DecoderExtension{}
	if __obf_25bd4f754a37b862.MarshalFloatWith6Digits {
		__obf_d4e161deb5b48ae2.__obf_8470c94b9e001994(__obf_9289f2028bb085f2)
	}
	if __obf_25bd4f754a37b862.EscapeHTML {
		__obf_d4e161deb5b48ae2.__obf_3ea234a937346d00(__obf_9289f2028bb085f2)
	}
	if __obf_25bd4f754a37b862.UseNumber {
		__obf_d4e161deb5b48ae2.__obf_376a95a68fb52ebb(__obf_7c9972bb545a908d)
	}
	if __obf_25bd4f754a37b862.ValidateJsonRawMessage {
		__obf_d4e161deb5b48ae2.__obf_61ebf38cf7df6e3e(__obf_9289f2028bb085f2)
	}
	__obf_d4e161deb5b48ae2.__obf_9289f2028bb085f2 = __obf_9289f2028bb085f2
	__obf_d4e161deb5b48ae2.__obf_7c9972bb545a908d = __obf_7c9972bb545a908d
	__obf_d4e161deb5b48ae2.__obf_4d2d9c0287da7aea = __obf_25bd4f754a37b862
	return __obf_d4e161deb5b48ae2
}

func (__obf_25bd4f754a37b862 Config) __obf_06ac4a8116bedc14(__obf_324ee50b8c0b2f3b []Extension) *__obf_ef74248d8ae9ba78 {
	__obf_d4e161deb5b48ae2 := __obf_cbbd424bc3cc37d7(__obf_25bd4f754a37b862)
	if __obf_d4e161deb5b48ae2 != nil {
		return __obf_d4e161deb5b48ae2
	}
	__obf_d4e161deb5b48ae2 = __obf_25bd4f754a37b862.Froze().(*__obf_ef74248d8ae9ba78)
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_324ee50b8c0b2f3b {
		__obf_d4e161deb5b48ae2.
			RegisterExtension(__obf_dfc6ab1ee0bfd79e)
	}
	__obf_b88cae1d27d0c0a1(__obf_25bd4f754a37b862, __obf_d4e161deb5b48ae2)
	return __obf_d4e161deb5b48ae2
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_61ebf38cf7df6e3e(__obf_dfc6ab1ee0bfd79e EncoderExtension) {
	__obf_cf840243a12ee308 := &__obf_12cc6cbc7e204422{func(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
		__obf_c0771f7df74d1a0c := *(*json.RawMessage)(__obf_47fa53f3d161f45c)
		__obf_47edab4c16a2d88d := __obf_25bd4f754a37b862.BorrowIterator([]byte(__obf_c0771f7df74d1a0c))
		defer __obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
		__obf_47edab4c16a2d88d.
			Read()
		if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
			__obf_9a34f62857fb1d1d.
				WriteRaw("null")
		} else {
			__obf_9a34f62857fb1d1d.
				WriteRaw(string(__obf_c0771f7df74d1a0c))
		}
	}, func(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
		return len(*((*json.RawMessage)(__obf_47fa53f3d161f45c))) == 0
	}}
	__obf_dfc6ab1ee0bfd79e[reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()] = __obf_cf840243a12ee308
	__obf_dfc6ab1ee0bfd79e[reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()] = __obf_cf840243a12ee308
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_376a95a68fb52ebb(__obf_dfc6ab1ee0bfd79e DecoderExtension) {
	__obf_dfc6ab1ee0bfd79e[reflect2.TypeOfPtr((*any)(nil)).Elem()] = &__obf_b283015dd3beecea{func(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
		__obf_8a28d88040812a23 := *((*any)(__obf_47fa53f3d161f45c))
		if __obf_8a28d88040812a23 != nil && reflect.TypeOf(__obf_8a28d88040812a23).Kind() == reflect.Ptr {
			__obf_47edab4c16a2d88d.
				ReadVal(__obf_8a28d88040812a23)
			return
		}
		if __obf_47edab4c16a2d88d.WhatIsNext() == NumberValue {
			*((*any)(__obf_47fa53f3d161f45c)) = json.Number(__obf_47edab4c16a2d88d.__obf_4c009361bc8be406())
		} else {
			*((*any)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.Read()
		}
	}}
}
func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_75bbb1cf9d1d76ab() string {
	__obf_4c3fe1329e8275b1 := __obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea.TagKey
	if __obf_4c3fe1329e8275b1 == "" {
		return "json"
	}
	return __obf_4c3fe1329e8275b1
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) RegisterExtension(__obf_dfc6ab1ee0bfd79e Extension) {
	__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b = append(__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b, __obf_dfc6ab1ee0bfd79e)
	__obf_6fd9ea911f2d54de := __obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea
	__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea = __obf_6fd9ea911f2d54de
}

type __obf_7d3b57f5983fb4ff struct {
}

func (__obf_cf840243a12ee308 *__obf_7d3b57f5983fb4ff) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteFloat32Lossy(*((*float32)(__obf_47fa53f3d161f45c)))
}

func (__obf_cf840243a12ee308 *__obf_7d3b57f5983fb4ff) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*float32)(__obf_47fa53f3d161f45c)) == 0
}

type __obf_14e7ec92d20e3cf9 struct {
}

func (__obf_cf840243a12ee308 *__obf_14e7ec92d20e3cf9) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteFloat64Lossy(*((*float64)(__obf_47fa53f3d161f45c)))
}

func (__obf_cf840243a12ee308 *__obf_14e7ec92d20e3cf9) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*float64)(__obf_47fa53f3d161f45c)) == 0
}

// EnableLossyFloatMarshalling keeps 10**(-6) precision
// for float variables for better performance.
func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_8470c94b9e001994(__obf_dfc6ab1ee0bfd79e EncoderExtension) {
	__obf_dfc6ab1ee0bfd79e[ // for better performance
	reflect2.TypeOfPtr((*float32)(nil)).Elem()] = &__obf_7d3b57f5983fb4ff{}
	__obf_dfc6ab1ee0bfd79e[reflect2.TypeOfPtr((*float64)(nil)).Elem()] = &__obf_14e7ec92d20e3cf9{}
}

type __obf_e04b00a4d9052f04 struct {
}

func (__obf_cf840243a12ee308 *__obf_e04b00a4d9052f04) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_44b48c26051f8033 := *((*string)(__obf_47fa53f3d161f45c))
	__obf_9a34f62857fb1d1d.
		WriteStringWithHTMLEscaped(__obf_44b48c26051f8033)
}

func (__obf_cf840243a12ee308 *__obf_e04b00a4d9052f04) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return *((*string)(__obf_47fa53f3d161f45c)) == ""
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_3ea234a937346d00(__obf_9289f2028bb085f2 EncoderExtension) {
	__obf_9289f2028bb085f2[reflect2.TypeOfPtr((*string)(nil)).Elem()] = &__obf_e04b00a4d9052f04{}
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_8b93374d940855e2() {
	__obf_a04d76a786d6c58c = map[string]ValDecoder{}
	__obf_abe36267b721fea3 = map[string]ValDecoder{}
	*__obf_25bd4f754a37b862 = *(__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea.Froze().(*__obf_ef74248d8ae9ba78))
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) __obf_972c7c5e3d0b693e() {
	__obf_9a66f8c0ef9aea38 = map[string]ValEncoder{}
	__obf_9cd47905513ef558 = map[string]ValEncoder{}
	*__obf_25bd4f754a37b862 = *(__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea.Froze().(*__obf_ef74248d8ae9ba78))
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) MarshalToString(__obf_9cfb140618a077d4 any) (string, error) {
	__obf_9a34f62857fb1d1d := __obf_25bd4f754a37b862.BorrowStream(nil)
	defer __obf_25bd4f754a37b862.ReturnStream(__obf_9a34f62857fb1d1d)
	__obf_9a34f62857fb1d1d.
		WriteVal(__obf_9cfb140618a077d4)
	if __obf_9a34f62857fb1d1d.Error != nil {
		return "", __obf_9a34f62857fb1d1d.Error
	}
	return string(__obf_9a34f62857fb1d1d.Buffer()), nil
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) Marshal(__obf_9cfb140618a077d4 any) ([]byte, error) {
	__obf_9a34f62857fb1d1d := __obf_25bd4f754a37b862.BorrowStream(nil)
	defer __obf_25bd4f754a37b862.ReturnStream(__obf_9a34f62857fb1d1d)
	__obf_9a34f62857fb1d1d.
		WriteVal(__obf_9cfb140618a077d4)
	if __obf_9a34f62857fb1d1d.Error != nil {
		return nil, __obf_9a34f62857fb1d1d.Error
	}
	__obf_6a8f120b2918ae8a := __obf_9a34f62857fb1d1d.Buffer()
	__obf_6fd9ea911f2d54de := make([]byte, len(__obf_6a8f120b2918ae8a))
	copy(__obf_6fd9ea911f2d54de, __obf_6a8f120b2918ae8a)
	return __obf_6fd9ea911f2d54de, nil
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) MarshalIndent(__obf_9cfb140618a077d4 any, __obf_c5fd316f3c4457de, __obf_0d119fc055fc9f67 string) ([]byte, error) {
	if __obf_c5fd316f3c4457de != "" {
		panic("prefix is not supported")
	}
	for _, __obf_a44636b770caa848 := range __obf_0d119fc055fc9f67 {
		if __obf_a44636b770caa848 != ' ' {
			panic("indent can only be space")
		}
	}
	__obf_7f2283488a4bd284 := __obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea
	__obf_7f2283488a4bd284.
		IndentionStep = len(__obf_0d119fc055fc9f67)
	return __obf_7f2283488a4bd284.__obf_06ac4a8116bedc14(__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b).Marshal(__obf_9cfb140618a077d4)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) UnmarshalFromString(__obf_44b48c26051f8033 string, __obf_9cfb140618a077d4 any) error {
	__obf_c180417ee204f8c5 := []byte(__obf_44b48c26051f8033)
	__obf_47edab4c16a2d88d := __obf_25bd4f754a37b862.BorrowIterator(__obf_c180417ee204f8c5)
	defer __obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadVal(__obf_9cfb140618a077d4)
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == 0 {
		if __obf_47edab4c16a2d88d.Error == io.EOF {
			return nil
		}
		return __obf_47edab4c16a2d88d.Error
	}
	__obf_47edab4c16a2d88d.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_47edab4c16a2d88d.Error
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) Get(__obf_c180417ee204f8c5 []byte, __obf_267bdf615cb1b310 ...any) Any {
	__obf_47edab4c16a2d88d := __obf_25bd4f754a37b862.BorrowIterator(__obf_c180417ee204f8c5)
	defer __obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	return __obf_da70fd4b038d543a(__obf_47edab4c16a2d88d, __obf_267bdf615cb1b310)
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) Unmarshal(__obf_c180417ee204f8c5 []byte, __obf_9cfb140618a077d4 any) error {
	__obf_47edab4c16a2d88d := __obf_25bd4f754a37b862.BorrowIterator(__obf_c180417ee204f8c5)
	defer __obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		ReadVal(__obf_9cfb140618a077d4)
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == 0 {
		if __obf_47edab4c16a2d88d.Error == io.EOF {
			return nil
		}
		return __obf_47edab4c16a2d88d.Error
	}
	__obf_47edab4c16a2d88d.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_47edab4c16a2d88d.Error
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) NewEncoder(__obf_8d0bd0f6d68b0c59 io.Writer) *Encoder {
	__obf_9a34f62857fb1d1d := NewStream(__obf_25bd4f754a37b862, __obf_8d0bd0f6d68b0c59, 512)
	return &Encoder{__obf_9a34f62857fb1d1d}
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) NewDecoder(__obf_3943589abf9d9fb6 io.Reader) *Decoder {
	__obf_47edab4c16a2d88d := Parse(__obf_25bd4f754a37b862, __obf_3943589abf9d9fb6, 512)
	return &Decoder{__obf_47edab4c16a2d88d}
}

func (__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78) Valid(__obf_c180417ee204f8c5 []byte) bool {
	__obf_47edab4c16a2d88d := __obf_25bd4f754a37b862.BorrowIterator(__obf_c180417ee204f8c5)
	defer __obf_25bd4f754a37b862.ReturnIterator(__obf_47edab4c16a2d88d)
	__obf_47edab4c16a2d88d.
		Skip()
	return __obf_47edab4c16a2d88d.Error == nil
}
