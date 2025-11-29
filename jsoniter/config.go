package __obf_5b802ce8d9ba56d6

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
	MarshalToString(__obf_91fc460a138dcae7 any) (string, error)
	Marshal(__obf_91fc460a138dcae7 any) ([]byte, error)
	MarshalIndent(__obf_91fc460a138dcae7 any, __obf_f3dd783c02964acf, __obf_8e554c17d6bbd080 string) ([]byte, error)
	UnmarshalFromString(__obf_12c21b79fa86dcba string, __obf_91fc460a138dcae7 any) error
	Unmarshal(__obf_5d3dff745e2b805b []byte, __obf_91fc460a138dcae7 any) error
	Get(__obf_5d3dff745e2b805b []byte, __obf_c5d71353f4393b3c ...any) Any
	NewEncoder(__obf_3dbedd483fe5a38a io.Writer) *Encoder
	NewDecoder(__obf_1fcab394164d9b41 io.Reader) *Decoder
	Valid(__obf_5d3dff745e2b805b []byte) bool
	RegisterExtension(__obf_6b17b29e9b6f5243 Extension)
	DecoderOf(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder
	EncoderOf(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder
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

type __obf_5d13d7f3bd06c6cf struct {
	__obf_97c6754e53034071 Config
	__obf_805c0ed9960be9bd bool
	__obf_ee9b4b613c13fea6 int
	__obf_35f74bd8c14f5694 bool
	__obf_4f5d4eba90b029f2 bool
	__obf_c395639ca05d31b5 bool
	__obf_20a9142f6944f823 *concurrent.Map
	__obf_e5a089c2b6cc7909 *concurrent.Map
	__obf_f754da59a9c09bdc Extension
	__obf_de1968c22ba7e047 Extension
	__obf_251d7593467966e4 []Extension
	__obf_a774d126f4306dfb *sync.Pool
	__obf_f74e8f2afdb9df78 *sync.Pool
	__obf_f48e3198f571baa9 bool
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_7539bc7dc0c17336() {
	__obf_dca106e1cdf85392.__obf_20a9142f6944f823 = concurrent.NewMap()
	__obf_dca106e1cdf85392.__obf_e5a089c2b6cc7909 = concurrent.NewMap()
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_88c5e2d37a5eaf26(__obf_760ddf980dcc31a5 uintptr, __obf_18f746d7b5b440ee ValDecoder) {
	__obf_dca106e1cdf85392.__obf_20a9142f6944f823.
		Store(__obf_760ddf980dcc31a5, __obf_18f746d7b5b440ee)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_7662f98a8ed819d0(__obf_760ddf980dcc31a5 uintptr, __obf_29366c3ad76a8c51 ValEncoder) {
	__obf_dca106e1cdf85392.__obf_e5a089c2b6cc7909.
		Store(__obf_760ddf980dcc31a5, __obf_29366c3ad76a8c51)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_0b82b8c45af29ba7(__obf_760ddf980dcc31a5 uintptr) ValDecoder {
	__obf_18f746d7b5b440ee, __obf_98a3eefc4c187186 := __obf_dca106e1cdf85392.__obf_20a9142f6944f823.Load(__obf_760ddf980dcc31a5)
	if __obf_98a3eefc4c187186 {
		return __obf_18f746d7b5b440ee.(ValDecoder)
	}
	return nil
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_2ca4a0a018fe3cfc(__obf_760ddf980dcc31a5 uintptr) ValEncoder {
	__obf_29366c3ad76a8c51, __obf_98a3eefc4c187186 := __obf_dca106e1cdf85392.__obf_e5a089c2b6cc7909.Load(__obf_760ddf980dcc31a5)
	if __obf_98a3eefc4c187186 {
		return __obf_29366c3ad76a8c51.(ValEncoder)
	}
	return nil
}

var __obf_1cc0c4c545e98064 = concurrent.NewMap()

func __obf_a753887d78871667(__obf_dca106e1cdf85392 Config) *__obf_5d13d7f3bd06c6cf {
	__obf_f9b1526531f3c024, __obf_98a3eefc4c187186 := __obf_1cc0c4c545e98064.Load(__obf_dca106e1cdf85392)
	if __obf_98a3eefc4c187186 {
		return __obf_f9b1526531f3c024.(*__obf_5d13d7f3bd06c6cf)
	}
	return nil
}

func __obf_29429a184833ea8a(__obf_dca106e1cdf85392 Config, __obf_5d13d7f3bd06c6cf *__obf_5d13d7f3bd06c6cf) {
	__obf_1cc0c4c545e98064.
		Store(__obf_dca106e1cdf85392,

			// Froze forge API from config
			__obf_5d13d7f3bd06c6cf)
}

func (__obf_dca106e1cdf85392 Config) Froze() API {
	__obf_549824fe479f4337 := &__obf_5d13d7f3bd06c6cf{__obf_805c0ed9960be9bd: __obf_dca106e1cdf85392.SortMapKeys, __obf_ee9b4b613c13fea6: __obf_dca106e1cdf85392.IndentionStep, __obf_35f74bd8c14f5694: __obf_dca106e1cdf85392.ObjectFieldMustBeSimpleString, __obf_4f5d4eba90b029f2: __obf_dca106e1cdf85392.OnlyTaggedField, __obf_c395639ca05d31b5: __obf_dca106e1cdf85392.DisallowUnknownFields, __obf_f48e3198f571baa9: __obf_dca106e1cdf85392.CaseSensitive}
	__obf_549824fe479f4337.__obf_a774d126f4306dfb = &sync.Pool{
		New: func() any {
			return NewStream(__obf_549824fe479f4337, nil, 512)
		},
	}
	__obf_549824fe479f4337.__obf_f74e8f2afdb9df78 = &sync.Pool{
		New: func() any {
			return NewIterator(__obf_549824fe479f4337)
		},
	}
	__obf_549824fe479f4337.__obf_7539bc7dc0c17336()
	__obf_f754da59a9c09bdc := EncoderExtension{}
	__obf_de1968c22ba7e047 := DecoderExtension{}
	if __obf_dca106e1cdf85392.MarshalFloatWith6Digits {
		__obf_549824fe479f4337.__obf_99892ef5ca170a26(__obf_f754da59a9c09bdc)
	}
	if __obf_dca106e1cdf85392.EscapeHTML {
		__obf_549824fe479f4337.__obf_0e412ce07ebed81c(__obf_f754da59a9c09bdc)
	}
	if __obf_dca106e1cdf85392.UseNumber {
		__obf_549824fe479f4337.__obf_af434d6332137770(__obf_de1968c22ba7e047)
	}
	if __obf_dca106e1cdf85392.ValidateJsonRawMessage {
		__obf_549824fe479f4337.__obf_10699f1bd571c1d7(__obf_f754da59a9c09bdc)
	}
	__obf_549824fe479f4337.__obf_f754da59a9c09bdc = __obf_f754da59a9c09bdc
	__obf_549824fe479f4337.__obf_de1968c22ba7e047 = __obf_de1968c22ba7e047
	__obf_549824fe479f4337.__obf_97c6754e53034071 = __obf_dca106e1cdf85392
	return __obf_549824fe479f4337
}

func (__obf_dca106e1cdf85392 Config) __obf_25a90ed48fcf6327(__obf_251d7593467966e4 []Extension) *__obf_5d13d7f3bd06c6cf {
	__obf_549824fe479f4337 := __obf_a753887d78871667(__obf_dca106e1cdf85392)
	if __obf_549824fe479f4337 != nil {
		return __obf_549824fe479f4337
	}
	__obf_549824fe479f4337 = __obf_dca106e1cdf85392.Froze().(*__obf_5d13d7f3bd06c6cf)
	for _, __obf_6b17b29e9b6f5243 := range __obf_251d7593467966e4 {
		__obf_549824fe479f4337.
			RegisterExtension(__obf_6b17b29e9b6f5243)
	}
	__obf_29429a184833ea8a(__obf_dca106e1cdf85392, __obf_549824fe479f4337)
	return __obf_549824fe479f4337
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_10699f1bd571c1d7(__obf_6b17b29e9b6f5243 EncoderExtension) {
	__obf_29366c3ad76a8c51 := &__obf_f00ee3350abbd8b3{func(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
		__obf_fc4dce51e684d724 := *(*json.RawMessage)(__obf_d3c919547bf7e47a)
		__obf_67008a6a9e5ba828 := __obf_dca106e1cdf85392.BorrowIterator([]byte(__obf_fc4dce51e684d724))
		defer __obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
		__obf_67008a6a9e5ba828.
			Read()
		if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
			__obf_00fc491caa967a74.
				WriteRaw("null")
		} else {
			__obf_00fc491caa967a74.
				WriteRaw(string(__obf_fc4dce51e684d724))
		}
	}, func(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
		return len(*((*json.RawMessage)(__obf_d3c919547bf7e47a))) == 0
	}}
	__obf_6b17b29e9b6f5243[reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()] = __obf_29366c3ad76a8c51
	__obf_6b17b29e9b6f5243[reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()] = __obf_29366c3ad76a8c51
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_af434d6332137770(__obf_6b17b29e9b6f5243 DecoderExtension) {
	__obf_6b17b29e9b6f5243[reflect2.TypeOfPtr((*any)(nil)).Elem()] = &__obf_362de96b93b8b0c3{func(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
		__obf_61109545050551b0 := *((*any)(__obf_d3c919547bf7e47a))
		if __obf_61109545050551b0 != nil && reflect.TypeOf(__obf_61109545050551b0).Kind() == reflect.Ptr {
			__obf_67008a6a9e5ba828.
				ReadVal(__obf_61109545050551b0)
			return
		}
		if __obf_67008a6a9e5ba828.WhatIsNext() == NumberValue {
			*((*any)(__obf_d3c919547bf7e47a)) = json.Number(__obf_67008a6a9e5ba828.__obf_bad2e5b8bc5a6a05())
		} else {
			*((*any)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.Read()
		}
	}}
}
func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_8deb8923819cc0a0() string {
	__obf_9ad4d6c08ca4178c := __obf_dca106e1cdf85392.__obf_97c6754e53034071.TagKey
	if __obf_9ad4d6c08ca4178c == "" {
		return "json"
	}
	return __obf_9ad4d6c08ca4178c
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) RegisterExtension(__obf_6b17b29e9b6f5243 Extension) {
	__obf_dca106e1cdf85392.__obf_251d7593467966e4 = append(__obf_dca106e1cdf85392.__obf_251d7593467966e4, __obf_6b17b29e9b6f5243)
	__obf_cd8f569527424a26 := __obf_dca106e1cdf85392.__obf_97c6754e53034071
	__obf_dca106e1cdf85392.__obf_97c6754e53034071 = __obf_cd8f569527424a26
}

type __obf_e6bf3da99734d2da struct {
}

func (__obf_29366c3ad76a8c51 *__obf_e6bf3da99734d2da) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteFloat32Lossy(*((*float32)(__obf_d3c919547bf7e47a)))
}

func (__obf_29366c3ad76a8c51 *__obf_e6bf3da99734d2da) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*float32)(__obf_d3c919547bf7e47a)) == 0
}

type __obf_b97b066e41cff5df struct {
}

func (__obf_29366c3ad76a8c51 *__obf_b97b066e41cff5df) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteFloat64Lossy(*((*float64)(__obf_d3c919547bf7e47a)))
}

func (__obf_29366c3ad76a8c51 *__obf_b97b066e41cff5df) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*float64)(__obf_d3c919547bf7e47a)) == 0
}

// EnableLossyFloatMarshalling keeps 10**(-6) precision
// for float variables for better performance.
func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_99892ef5ca170a26(__obf_6b17b29e9b6f5243 EncoderExtension) {
	__obf_6b17b29e9b6f5243[ // for better performance
	reflect2.TypeOfPtr((*float32)(nil)).Elem()] = &__obf_e6bf3da99734d2da{}
	__obf_6b17b29e9b6f5243[reflect2.TypeOfPtr((*float64)(nil)).Elem()] = &__obf_b97b066e41cff5df{}
}

type __obf_5ac877a585fa7108 struct {
}

func (__obf_29366c3ad76a8c51 *__obf_5ac877a585fa7108) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_12c21b79fa86dcba := *((*string)(__obf_d3c919547bf7e47a))
	__obf_00fc491caa967a74.
		WriteStringWithHTMLEscaped(__obf_12c21b79fa86dcba)
}

func (__obf_29366c3ad76a8c51 *__obf_5ac877a585fa7108) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return *((*string)(__obf_d3c919547bf7e47a)) == ""
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_0e412ce07ebed81c(__obf_f754da59a9c09bdc EncoderExtension) {
	__obf_f754da59a9c09bdc[reflect2.TypeOfPtr((*string)(nil)).Elem()] = &__obf_5ac877a585fa7108{}
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_e24222fe2189fef7() {
	__obf_52b2f46f997a7928 = map[string]ValDecoder{}
	__obf_6f60d6128c9e434a = map[string]ValDecoder{}
	*__obf_dca106e1cdf85392 = *(__obf_dca106e1cdf85392.__obf_97c6754e53034071.Froze().(*__obf_5d13d7f3bd06c6cf))
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) __obf_5651829942018369() {
	__obf_808b92c3db71ce90 = map[string]ValEncoder{}
	__obf_88154662404be67e = map[string]ValEncoder{}
	*__obf_dca106e1cdf85392 = *(__obf_dca106e1cdf85392.__obf_97c6754e53034071.Froze().(*__obf_5d13d7f3bd06c6cf))
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) MarshalToString(__obf_91fc460a138dcae7 any) (string, error) {
	__obf_00fc491caa967a74 := __obf_dca106e1cdf85392.BorrowStream(nil)
	defer __obf_dca106e1cdf85392.ReturnStream(__obf_00fc491caa967a74)
	__obf_00fc491caa967a74.
		WriteVal(__obf_91fc460a138dcae7)
	if __obf_00fc491caa967a74.Error != nil {
		return "", __obf_00fc491caa967a74.Error
	}
	return string(__obf_00fc491caa967a74.Buffer()), nil
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) Marshal(__obf_91fc460a138dcae7 any) ([]byte, error) {
	__obf_00fc491caa967a74 := __obf_dca106e1cdf85392.BorrowStream(nil)
	defer __obf_dca106e1cdf85392.ReturnStream(__obf_00fc491caa967a74)
	__obf_00fc491caa967a74.
		WriteVal(__obf_91fc460a138dcae7)
	if __obf_00fc491caa967a74.Error != nil {
		return nil, __obf_00fc491caa967a74.Error
	}
	__obf_b2354bdea3b1a12a := __obf_00fc491caa967a74.Buffer()
	__obf_cd8f569527424a26 := make([]byte, len(__obf_b2354bdea3b1a12a))
	copy(__obf_cd8f569527424a26, __obf_b2354bdea3b1a12a)
	return __obf_cd8f569527424a26, nil
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) MarshalIndent(__obf_91fc460a138dcae7 any, __obf_f3dd783c02964acf, __obf_8e554c17d6bbd080 string) ([]byte, error) {
	if __obf_f3dd783c02964acf != "" {
		panic("prefix is not supported")
	}
	for _, __obf_14c0fc639db9b91c := range __obf_8e554c17d6bbd080 {
		if __obf_14c0fc639db9b91c != ' ' {
			panic("indent can only be space")
		}
	}
	__obf_519ba0f5118db09d := __obf_dca106e1cdf85392.__obf_97c6754e53034071
	__obf_519ba0f5118db09d.
		IndentionStep = len(__obf_8e554c17d6bbd080)
	return __obf_519ba0f5118db09d.__obf_25a90ed48fcf6327(__obf_dca106e1cdf85392.__obf_251d7593467966e4).Marshal(__obf_91fc460a138dcae7)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) UnmarshalFromString(__obf_12c21b79fa86dcba string, __obf_91fc460a138dcae7 any) error {
	__obf_5d3dff745e2b805b := []byte(__obf_12c21b79fa86dcba)
	__obf_67008a6a9e5ba828 := __obf_dca106e1cdf85392.BorrowIterator(__obf_5d3dff745e2b805b)
	defer __obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadVal(__obf_91fc460a138dcae7)
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == 0 {
		if __obf_67008a6a9e5ba828.Error == io.EOF {
			return nil
		}
		return __obf_67008a6a9e5ba828.Error
	}
	__obf_67008a6a9e5ba828.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_67008a6a9e5ba828.Error
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) Get(__obf_5d3dff745e2b805b []byte, __obf_c5d71353f4393b3c ...any) Any {
	__obf_67008a6a9e5ba828 := __obf_dca106e1cdf85392.BorrowIterator(__obf_5d3dff745e2b805b)
	defer __obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	return __obf_9c645421e8d71b4a(__obf_67008a6a9e5ba828, __obf_c5d71353f4393b3c)
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) Unmarshal(__obf_5d3dff745e2b805b []byte, __obf_91fc460a138dcae7 any) error {
	__obf_67008a6a9e5ba828 := __obf_dca106e1cdf85392.BorrowIterator(__obf_5d3dff745e2b805b)
	defer __obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		ReadVal(__obf_91fc460a138dcae7)
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == 0 {
		if __obf_67008a6a9e5ba828.Error == io.EOF {
			return nil
		}
		return __obf_67008a6a9e5ba828.Error
	}
	__obf_67008a6a9e5ba828.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_67008a6a9e5ba828.Error
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) NewEncoder(__obf_3dbedd483fe5a38a io.Writer) *Encoder {
	__obf_00fc491caa967a74 := NewStream(__obf_dca106e1cdf85392, __obf_3dbedd483fe5a38a, 512)
	return &Encoder{__obf_00fc491caa967a74}
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) NewDecoder(__obf_1fcab394164d9b41 io.Reader) *Decoder {
	__obf_67008a6a9e5ba828 := Parse(__obf_dca106e1cdf85392, __obf_1fcab394164d9b41, 512)
	return &Decoder{__obf_67008a6a9e5ba828}
}

func (__obf_dca106e1cdf85392 *__obf_5d13d7f3bd06c6cf) Valid(__obf_5d3dff745e2b805b []byte) bool {
	__obf_67008a6a9e5ba828 := __obf_dca106e1cdf85392.BorrowIterator(__obf_5d3dff745e2b805b)
	defer __obf_dca106e1cdf85392.ReturnIterator(__obf_67008a6a9e5ba828)
	__obf_67008a6a9e5ba828.
		Skip()
	return __obf_67008a6a9e5ba828.Error == nil
}
