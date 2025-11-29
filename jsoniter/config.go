package __obf_91620b895eeff9ed

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
	MarshalToString(__obf_97f02ddd0770463f any) (string, error)
	Marshal(__obf_97f02ddd0770463f any) ([]byte, error)
	MarshalIndent(__obf_97f02ddd0770463f any, __obf_6e20c91e70903582, __obf_6ab410c037fcd950 string) ([]byte, error)
	UnmarshalFromString(__obf_e91bd2feb751e4f1 string, __obf_97f02ddd0770463f any) error
	Unmarshal(__obf_2894589095c323b4 []byte, __obf_97f02ddd0770463f any) error
	Get(__obf_2894589095c323b4 []byte, __obf_82c50fcbfc9ab432 ...any) Any
	NewEncoder(__obf_e4132b2c3ef9e14e io.Writer) *Encoder
	NewDecoder(__obf_fe90e88a3ba69cf3 io.Reader) *Decoder
	Valid(__obf_2894589095c323b4 []byte) bool
	RegisterExtension(__obf_8c9d73a8f8319687 Extension)
	DecoderOf(__obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder
	EncoderOf(__obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder
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

type __obf_972c2293804d141c struct {
	__obf_e05f98c9e7cc9089 Config
	__obf_e63708933643279b bool
	__obf_f186c052066ae54a int
	__obf_065d67ebe135fdba bool
	__obf_dc0073de34b8604d bool
	__obf_8d8c301f6a725e61 bool
	__obf_596a9743225af4ef *concurrent.Map
	__obf_8f33877e4e309b45 *concurrent.Map
	__obf_47929f7efe51b371 Extension
	__obf_920e2f1ddf47b5e1 Extension
	__obf_b4dfae156c7ac26c []Extension
	__obf_9e7b77251a4cf7c4 *sync.Pool
	__obf_145970d48796be8b *sync.Pool
	__obf_f087a7a47617f72a bool
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_55f171a032c0b522() {
	__obf_892632c77e859037.__obf_596a9743225af4ef = concurrent.NewMap()
	__obf_892632c77e859037.__obf_8f33877e4e309b45 = concurrent.NewMap()
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_70fb224d773d428b(__obf_f5a66ce7c84d060d uintptr, __obf_6fd3f72eb9b5574c ValDecoder) {
	__obf_892632c77e859037.__obf_596a9743225af4ef.
		Store(__obf_f5a66ce7c84d060d, __obf_6fd3f72eb9b5574c)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_56d16fae33267fc7(__obf_f5a66ce7c84d060d uintptr, __obf_96e65a4c8c4f2ce5 ValEncoder) {
	__obf_892632c77e859037.__obf_8f33877e4e309b45.
		Store(__obf_f5a66ce7c84d060d, __obf_96e65a4c8c4f2ce5)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_49679f79c9853606(__obf_f5a66ce7c84d060d uintptr) ValDecoder {
	__obf_6fd3f72eb9b5574c, __obf_93f189bffabaf5a4 := __obf_892632c77e859037.__obf_596a9743225af4ef.Load(__obf_f5a66ce7c84d060d)
	if __obf_93f189bffabaf5a4 {
		return __obf_6fd3f72eb9b5574c.(ValDecoder)
	}
	return nil
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_7892b430524f2484(__obf_f5a66ce7c84d060d uintptr) ValEncoder {
	__obf_96e65a4c8c4f2ce5, __obf_93f189bffabaf5a4 := __obf_892632c77e859037.__obf_8f33877e4e309b45.Load(__obf_f5a66ce7c84d060d)
	if __obf_93f189bffabaf5a4 {
		return __obf_96e65a4c8c4f2ce5.(ValEncoder)
	}
	return nil
}

var __obf_4694fb1f5a5f6a79 = concurrent.NewMap()

func __obf_56ab43fe764ccbfe(__obf_892632c77e859037 Config) *__obf_972c2293804d141c {
	__obf_35136ef2ac0f94e4, __obf_93f189bffabaf5a4 := __obf_4694fb1f5a5f6a79.Load(__obf_892632c77e859037)
	if __obf_93f189bffabaf5a4 {
		return __obf_35136ef2ac0f94e4.(*__obf_972c2293804d141c)
	}
	return nil
}

func __obf_fc7b524923d8fb17(__obf_892632c77e859037 Config, __obf_972c2293804d141c *__obf_972c2293804d141c) {
	__obf_4694fb1f5a5f6a79.
		Store(__obf_892632c77e859037,

			// Froze forge API from config
			__obf_972c2293804d141c)
}

func (__obf_892632c77e859037 Config) Froze() API {
	__obf_f23c57ee5021b4b7 := &__obf_972c2293804d141c{__obf_e63708933643279b: __obf_892632c77e859037.SortMapKeys, __obf_f186c052066ae54a: __obf_892632c77e859037.IndentionStep, __obf_065d67ebe135fdba: __obf_892632c77e859037.ObjectFieldMustBeSimpleString, __obf_dc0073de34b8604d: __obf_892632c77e859037.OnlyTaggedField, __obf_8d8c301f6a725e61: __obf_892632c77e859037.DisallowUnknownFields, __obf_f087a7a47617f72a: __obf_892632c77e859037.CaseSensitive}
	__obf_f23c57ee5021b4b7.__obf_9e7b77251a4cf7c4 = &sync.Pool{
		New: func() any {
			return NewStream(__obf_f23c57ee5021b4b7, nil, 512)
		},
	}
	__obf_f23c57ee5021b4b7.__obf_145970d48796be8b = &sync.Pool{
		New: func() any {
			return NewIterator(__obf_f23c57ee5021b4b7)
		},
	}
	__obf_f23c57ee5021b4b7.__obf_55f171a032c0b522()
	__obf_47929f7efe51b371 := EncoderExtension{}
	__obf_920e2f1ddf47b5e1 := DecoderExtension{}
	if __obf_892632c77e859037.MarshalFloatWith6Digits {
		__obf_f23c57ee5021b4b7.__obf_79f1cd8e38e8dcd8(__obf_47929f7efe51b371)
	}
	if __obf_892632c77e859037.EscapeHTML {
		__obf_f23c57ee5021b4b7.__obf_6ee6f0f6c0425676(__obf_47929f7efe51b371)
	}
	if __obf_892632c77e859037.UseNumber {
		__obf_f23c57ee5021b4b7.__obf_2b34ee0adffb133a(__obf_920e2f1ddf47b5e1)
	}
	if __obf_892632c77e859037.ValidateJsonRawMessage {
		__obf_f23c57ee5021b4b7.__obf_4f552bbce3a0e4b0(__obf_47929f7efe51b371)
	}
	__obf_f23c57ee5021b4b7.__obf_47929f7efe51b371 = __obf_47929f7efe51b371
	__obf_f23c57ee5021b4b7.__obf_920e2f1ddf47b5e1 = __obf_920e2f1ddf47b5e1
	__obf_f23c57ee5021b4b7.__obf_e05f98c9e7cc9089 = __obf_892632c77e859037
	return __obf_f23c57ee5021b4b7
}

func (__obf_892632c77e859037 Config) __obf_6f0de5d9837befa9(__obf_b4dfae156c7ac26c []Extension) *__obf_972c2293804d141c {
	__obf_f23c57ee5021b4b7 := __obf_56ab43fe764ccbfe(__obf_892632c77e859037)
	if __obf_f23c57ee5021b4b7 != nil {
		return __obf_f23c57ee5021b4b7
	}
	__obf_f23c57ee5021b4b7 = __obf_892632c77e859037.Froze().(*__obf_972c2293804d141c)
	for _, __obf_8c9d73a8f8319687 := range __obf_b4dfae156c7ac26c {
		__obf_f23c57ee5021b4b7.
			RegisterExtension(__obf_8c9d73a8f8319687)
	}
	__obf_fc7b524923d8fb17(__obf_892632c77e859037, __obf_f23c57ee5021b4b7)
	return __obf_f23c57ee5021b4b7
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_4f552bbce3a0e4b0(__obf_8c9d73a8f8319687 EncoderExtension) {
	__obf_96e65a4c8c4f2ce5 := &__obf_009fa1223393ec9a{func(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
		__obf_c319554f95d3a66c := *(*json.RawMessage)(__obf_2a1474febb02279b)
		__obf_1bb30e8a74ed8233 := __obf_892632c77e859037.BorrowIterator([]byte(__obf_c319554f95d3a66c))
		defer __obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
		__obf_1bb30e8a74ed8233.
			Read()
		if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
			__obf_850a7457bb739a32.
				WriteRaw("null")
		} else {
			__obf_850a7457bb739a32.
				WriteRaw(string(__obf_c319554f95d3a66c))
		}
	}, func(__obf_2a1474febb02279b unsafe.Pointer) bool {
		return len(*((*json.RawMessage)(__obf_2a1474febb02279b))) == 0
	}}
	__obf_8c9d73a8f8319687[reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()] = __obf_96e65a4c8c4f2ce5
	__obf_8c9d73a8f8319687[reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()] = __obf_96e65a4c8c4f2ce5
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_2b34ee0adffb133a(__obf_8c9d73a8f8319687 DecoderExtension) {
	__obf_8c9d73a8f8319687[reflect2.TypeOfPtr((*any)(nil)).Elem()] = &__obf_b33e9e2851d3e569{func(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
		__obf_538b34b0607baeff := *((*any)(__obf_2a1474febb02279b))
		if __obf_538b34b0607baeff != nil && reflect.TypeOf(__obf_538b34b0607baeff).Kind() == reflect.Ptr {
			__obf_1bb30e8a74ed8233.
				ReadVal(__obf_538b34b0607baeff)
			return
		}
		if __obf_1bb30e8a74ed8233.WhatIsNext() == NumberValue {
			*((*any)(__obf_2a1474febb02279b)) = json.Number(__obf_1bb30e8a74ed8233.__obf_9df99ebed87959cd())
		} else {
			*((*any)(__obf_2a1474febb02279b)) = __obf_1bb30e8a74ed8233.Read()
		}
	}}
}
func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_f6ba056498acdb88() string {
	__obf_cd3966e81a1b34e3 := __obf_892632c77e859037.__obf_e05f98c9e7cc9089.TagKey
	if __obf_cd3966e81a1b34e3 == "" {
		return "json"
	}
	return __obf_cd3966e81a1b34e3
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) RegisterExtension(__obf_8c9d73a8f8319687 Extension) {
	__obf_892632c77e859037.__obf_b4dfae156c7ac26c = append(__obf_892632c77e859037.__obf_b4dfae156c7ac26c, __obf_8c9d73a8f8319687)
	__obf_fce0adcefc66d2ad := __obf_892632c77e859037.__obf_e05f98c9e7cc9089
	__obf_892632c77e859037.__obf_e05f98c9e7cc9089 = __obf_fce0adcefc66d2ad
}

type __obf_001c5343c2e5281f struct {
}

func (__obf_96e65a4c8c4f2ce5 *__obf_001c5343c2e5281f) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteFloat32Lossy(*((*float32)(__obf_2a1474febb02279b)))
}

func (__obf_96e65a4c8c4f2ce5 *__obf_001c5343c2e5281f) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*float32)(__obf_2a1474febb02279b)) == 0
}

type __obf_fa97f7ef5a06b5ec struct {
}

func (__obf_96e65a4c8c4f2ce5 *__obf_fa97f7ef5a06b5ec) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteFloat64Lossy(*((*float64)(__obf_2a1474febb02279b)))
}

func (__obf_96e65a4c8c4f2ce5 *__obf_fa97f7ef5a06b5ec) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*float64)(__obf_2a1474febb02279b)) == 0
}

// EnableLossyFloatMarshalling keeps 10**(-6) precision
// for float variables for better performance.
func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_79f1cd8e38e8dcd8(__obf_8c9d73a8f8319687 EncoderExtension) {
	__obf_8c9d73a8f8319687[ // for better performance
	reflect2.TypeOfPtr((*float32)(nil)).Elem()] = &__obf_001c5343c2e5281f{}
	__obf_8c9d73a8f8319687[reflect2.TypeOfPtr((*float64)(nil)).Elem()] = &__obf_fa97f7ef5a06b5ec{}
}

type __obf_aea10c981842c332 struct {
}

func (__obf_96e65a4c8c4f2ce5 *__obf_aea10c981842c332) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_e91bd2feb751e4f1 := *((*string)(__obf_2a1474febb02279b))
	__obf_850a7457bb739a32.
		WriteStringWithHTMLEscaped(__obf_e91bd2feb751e4f1)
}

func (__obf_96e65a4c8c4f2ce5 *__obf_aea10c981842c332) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	return *((*string)(__obf_2a1474febb02279b)) == ""
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_6ee6f0f6c0425676(__obf_47929f7efe51b371 EncoderExtension) {
	__obf_47929f7efe51b371[reflect2.TypeOfPtr((*string)(nil)).Elem()] = &__obf_aea10c981842c332{}
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_0e9c8b94e9472b05() {
	__obf_c0b846a36a3844aa = map[string]ValDecoder{}
	__obf_c7a2a8e5abb2ff8a = map[string]ValDecoder{}
	*__obf_892632c77e859037 = *(__obf_892632c77e859037.__obf_e05f98c9e7cc9089.Froze().(*__obf_972c2293804d141c))
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) __obf_c202ab740a5110fc() {
	__obf_ad65186839258b68 = map[string]ValEncoder{}
	__obf_afeecc600805f008 = map[string]ValEncoder{}
	*__obf_892632c77e859037 = *(__obf_892632c77e859037.__obf_e05f98c9e7cc9089.Froze().(*__obf_972c2293804d141c))
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) MarshalToString(__obf_97f02ddd0770463f any) (string, error) {
	__obf_850a7457bb739a32 := __obf_892632c77e859037.BorrowStream(nil)
	defer __obf_892632c77e859037.ReturnStream(__obf_850a7457bb739a32)
	__obf_850a7457bb739a32.
		WriteVal(__obf_97f02ddd0770463f)
	if __obf_850a7457bb739a32.Error != nil {
		return "", __obf_850a7457bb739a32.Error
	}
	return string(__obf_850a7457bb739a32.Buffer()), nil
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) Marshal(__obf_97f02ddd0770463f any) ([]byte, error) {
	__obf_850a7457bb739a32 := __obf_892632c77e859037.BorrowStream(nil)
	defer __obf_892632c77e859037.ReturnStream(__obf_850a7457bb739a32)
	__obf_850a7457bb739a32.
		WriteVal(__obf_97f02ddd0770463f)
	if __obf_850a7457bb739a32.Error != nil {
		return nil, __obf_850a7457bb739a32.Error
	}
	__obf_350e93f99f49e64c := __obf_850a7457bb739a32.Buffer()
	__obf_fce0adcefc66d2ad := make([]byte, len(__obf_350e93f99f49e64c))
	copy(__obf_fce0adcefc66d2ad, __obf_350e93f99f49e64c)
	return __obf_fce0adcefc66d2ad, nil
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) MarshalIndent(__obf_97f02ddd0770463f any, __obf_6e20c91e70903582, __obf_6ab410c037fcd950 string) ([]byte, error) {
	if __obf_6e20c91e70903582 != "" {
		panic("prefix is not supported")
	}
	for _, __obf_44946063bfe54da7 := range __obf_6ab410c037fcd950 {
		if __obf_44946063bfe54da7 != ' ' {
			panic("indent can only be space")
		}
	}
	__obf_579b02cd2a2330d1 := __obf_892632c77e859037.__obf_e05f98c9e7cc9089
	__obf_579b02cd2a2330d1.
		IndentionStep = len(__obf_6ab410c037fcd950)
	return __obf_579b02cd2a2330d1.__obf_6f0de5d9837befa9(__obf_892632c77e859037.__obf_b4dfae156c7ac26c).Marshal(__obf_97f02ddd0770463f)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) UnmarshalFromString(__obf_e91bd2feb751e4f1 string, __obf_97f02ddd0770463f any) error {
	__obf_2894589095c323b4 := []byte(__obf_e91bd2feb751e4f1)
	__obf_1bb30e8a74ed8233 := __obf_892632c77e859037.BorrowIterator(__obf_2894589095c323b4)
	defer __obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadVal(__obf_97f02ddd0770463f)
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == 0 {
		if __obf_1bb30e8a74ed8233.Error == io.EOF {
			return nil
		}
		return __obf_1bb30e8a74ed8233.Error
	}
	__obf_1bb30e8a74ed8233.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_1bb30e8a74ed8233.Error
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) Get(__obf_2894589095c323b4 []byte, __obf_82c50fcbfc9ab432 ...any) Any {
	__obf_1bb30e8a74ed8233 := __obf_892632c77e859037.BorrowIterator(__obf_2894589095c323b4)
	defer __obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	return __obf_67ce74ddb71f8159(__obf_1bb30e8a74ed8233, __obf_82c50fcbfc9ab432)
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) Unmarshal(__obf_2894589095c323b4 []byte, __obf_97f02ddd0770463f any) error {
	__obf_1bb30e8a74ed8233 := __obf_892632c77e859037.BorrowIterator(__obf_2894589095c323b4)
	defer __obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		ReadVal(__obf_97f02ddd0770463f)
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == 0 {
		if __obf_1bb30e8a74ed8233.Error == io.EOF {
			return nil
		}
		return __obf_1bb30e8a74ed8233.Error
	}
	__obf_1bb30e8a74ed8233.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_1bb30e8a74ed8233.Error
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) NewEncoder(__obf_e4132b2c3ef9e14e io.Writer) *Encoder {
	__obf_850a7457bb739a32 := NewStream(__obf_892632c77e859037, __obf_e4132b2c3ef9e14e, 512)
	return &Encoder{__obf_850a7457bb739a32}
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) NewDecoder(__obf_fe90e88a3ba69cf3 io.Reader) *Decoder {
	__obf_1bb30e8a74ed8233 := Parse(__obf_892632c77e859037, __obf_fe90e88a3ba69cf3, 512)
	return &Decoder{__obf_1bb30e8a74ed8233}
}

func (__obf_892632c77e859037 *__obf_972c2293804d141c) Valid(__obf_2894589095c323b4 []byte) bool {
	__obf_1bb30e8a74ed8233 := __obf_892632c77e859037.BorrowIterator(__obf_2894589095c323b4)
	defer __obf_892632c77e859037.ReturnIterator(__obf_1bb30e8a74ed8233)
	__obf_1bb30e8a74ed8233.
		Skip()
	return __obf_1bb30e8a74ed8233.Error == nil
}
