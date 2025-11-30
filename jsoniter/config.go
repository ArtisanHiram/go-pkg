package __obf_c3cd04a15c56f32f

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
	MarshalToString(__obf_f967fc9e700d5e33 any) (string, error)
	Marshal(__obf_f967fc9e700d5e33 any) ([]byte, error)
	MarshalIndent(__obf_f967fc9e700d5e33 any, __obf_f517983aa5897f07, __obf_30653456d76688b8 string) ([]byte, error)
	UnmarshalFromString(__obf_2d944450d48e5810 string, __obf_f967fc9e700d5e33 any) error
	Unmarshal(__obf_905295f6bd29aafe []byte, __obf_f967fc9e700d5e33 any) error
	Get(__obf_905295f6bd29aafe []byte, __obf_b90e4ca332607787 ...any) Any
	NewEncoder(__obf_b1699a146b20b28e io.Writer) *Encoder
	NewDecoder(__obf_a539369e10727e42 io.Reader) *Decoder
	Valid(__obf_905295f6bd29aafe []byte) bool
	RegisterExtension(__obf_e4a614f491c1bb0c Extension)
	DecoderOf(__obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder
	EncoderOf(__obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder
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

type __obf_3a25fb4c9a8342bb struct {
	__obf_6f7054bdfbca17ab Config
	__obf_70c2e3a03299f745 bool
	__obf_1ee92c16ad6ec329 int
	__obf_ccb7c8499611626f bool
	__obf_600fbe2319e94519 bool
	__obf_97cd803c79c71ab5 bool
	__obf_2b44eede29f9c5c4 *concurrent.Map
	__obf_3f8a87aaa1438c70 *concurrent.Map
	__obf_912608e689e59f9b Extension
	__obf_d86968bf26bed402 Extension
	__obf_a1eac24bcc56f0a6 []Extension
	__obf_456a3a3023fd9656 *sync.Pool
	__obf_7e0b4a46cc5f35d8 *sync.Pool
	__obf_b600271a247f99b8 bool
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_b07c4f7fec3f409a() {
	__obf_0c8065c219ccf0ab.__obf_2b44eede29f9c5c4 = concurrent.NewMap()
	__obf_0c8065c219ccf0ab.__obf_3f8a87aaa1438c70 = concurrent.NewMap()
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_eb4fd71a722727bf(__obf_cb552b6dac1d7edc uintptr, __obf_924cc7ef585cdfb0 ValDecoder) {
	__obf_0c8065c219ccf0ab.__obf_2b44eede29f9c5c4.
		Store(__obf_cb552b6dac1d7edc, __obf_924cc7ef585cdfb0)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_7ac7db4cbd198195(__obf_cb552b6dac1d7edc uintptr, __obf_c85b895560db528f ValEncoder) {
	__obf_0c8065c219ccf0ab.__obf_3f8a87aaa1438c70.
		Store(__obf_cb552b6dac1d7edc, __obf_c85b895560db528f)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_70f630ceb9086844(__obf_cb552b6dac1d7edc uintptr) ValDecoder {
	__obf_924cc7ef585cdfb0, __obf_da9e40840d5dbfdf := __obf_0c8065c219ccf0ab.__obf_2b44eede29f9c5c4.Load(__obf_cb552b6dac1d7edc)
	if __obf_da9e40840d5dbfdf {
		return __obf_924cc7ef585cdfb0.(ValDecoder)
	}
	return nil
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_ca7fd3d0f377cbfd(__obf_cb552b6dac1d7edc uintptr) ValEncoder {
	__obf_c85b895560db528f, __obf_da9e40840d5dbfdf := __obf_0c8065c219ccf0ab.__obf_3f8a87aaa1438c70.Load(__obf_cb552b6dac1d7edc)
	if __obf_da9e40840d5dbfdf {
		return __obf_c85b895560db528f.(ValEncoder)
	}
	return nil
}

var __obf_ed61aea428af7698 = concurrent.NewMap()

func __obf_a92e2cd9ef0a239d(__obf_0c8065c219ccf0ab Config) *__obf_3a25fb4c9a8342bb {
	__obf_50acae8c0a871ba1, __obf_da9e40840d5dbfdf := __obf_ed61aea428af7698.Load(__obf_0c8065c219ccf0ab)
	if __obf_da9e40840d5dbfdf {
		return __obf_50acae8c0a871ba1.(*__obf_3a25fb4c9a8342bb)
	}
	return nil
}

func __obf_d97b125cf13e890c(__obf_0c8065c219ccf0ab Config, __obf_3a25fb4c9a8342bb *__obf_3a25fb4c9a8342bb) {
	__obf_ed61aea428af7698.
		Store(__obf_0c8065c219ccf0ab,

			// Froze forge API from config
			__obf_3a25fb4c9a8342bb)
}

func (__obf_0c8065c219ccf0ab Config) Froze() API {
	__obf_e7dfc0b74a6746ef := &__obf_3a25fb4c9a8342bb{__obf_70c2e3a03299f745: __obf_0c8065c219ccf0ab.SortMapKeys, __obf_1ee92c16ad6ec329: __obf_0c8065c219ccf0ab.IndentionStep, __obf_ccb7c8499611626f: __obf_0c8065c219ccf0ab.ObjectFieldMustBeSimpleString, __obf_600fbe2319e94519: __obf_0c8065c219ccf0ab.OnlyTaggedField, __obf_97cd803c79c71ab5: __obf_0c8065c219ccf0ab.DisallowUnknownFields, __obf_b600271a247f99b8: __obf_0c8065c219ccf0ab.CaseSensitive}
	__obf_e7dfc0b74a6746ef.__obf_456a3a3023fd9656 = &sync.Pool{
		New: func() any {
			return NewStream(__obf_e7dfc0b74a6746ef, nil, 512)
		},
	}
	__obf_e7dfc0b74a6746ef.__obf_7e0b4a46cc5f35d8 = &sync.Pool{
		New: func() any {
			return NewIterator(__obf_e7dfc0b74a6746ef)
		},
	}
	__obf_e7dfc0b74a6746ef.__obf_b07c4f7fec3f409a()
	__obf_912608e689e59f9b := EncoderExtension{}
	__obf_d86968bf26bed402 := DecoderExtension{}
	if __obf_0c8065c219ccf0ab.MarshalFloatWith6Digits {
		__obf_e7dfc0b74a6746ef.__obf_d34b631463eb4202(__obf_912608e689e59f9b)
	}
	if __obf_0c8065c219ccf0ab.EscapeHTML {
		__obf_e7dfc0b74a6746ef.__obf_98a068a278faf6e8(__obf_912608e689e59f9b)
	}
	if __obf_0c8065c219ccf0ab.UseNumber {
		__obf_e7dfc0b74a6746ef.__obf_ef36b5509db9a3be(__obf_d86968bf26bed402)
	}
	if __obf_0c8065c219ccf0ab.ValidateJsonRawMessage {
		__obf_e7dfc0b74a6746ef.__obf_d5964e0396fd3d5b(__obf_912608e689e59f9b)
	}
	__obf_e7dfc0b74a6746ef.__obf_912608e689e59f9b = __obf_912608e689e59f9b
	__obf_e7dfc0b74a6746ef.__obf_d86968bf26bed402 = __obf_d86968bf26bed402
	__obf_e7dfc0b74a6746ef.__obf_6f7054bdfbca17ab = __obf_0c8065c219ccf0ab
	return __obf_e7dfc0b74a6746ef
}

func (__obf_0c8065c219ccf0ab Config) __obf_c3344b4ada41a413(__obf_a1eac24bcc56f0a6 []Extension) *__obf_3a25fb4c9a8342bb {
	__obf_e7dfc0b74a6746ef := __obf_a92e2cd9ef0a239d(__obf_0c8065c219ccf0ab)
	if __obf_e7dfc0b74a6746ef != nil {
		return __obf_e7dfc0b74a6746ef
	}
	__obf_e7dfc0b74a6746ef = __obf_0c8065c219ccf0ab.Froze().(*__obf_3a25fb4c9a8342bb)
	for _, __obf_e4a614f491c1bb0c := range __obf_a1eac24bcc56f0a6 {
		__obf_e7dfc0b74a6746ef.
			RegisterExtension(__obf_e4a614f491c1bb0c)
	}
	__obf_d97b125cf13e890c(__obf_0c8065c219ccf0ab, __obf_e7dfc0b74a6746ef)
	return __obf_e7dfc0b74a6746ef
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_d5964e0396fd3d5b(__obf_e4a614f491c1bb0c EncoderExtension) {
	__obf_c85b895560db528f := &__obf_f57166fb6630c677{func(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
		__obf_f6f05d0629ad7fc3 := *(*json.RawMessage)(__obf_753ab3fb72cdd659)
		__obf_edd9fe48d39445e4 := __obf_0c8065c219ccf0ab.BorrowIterator([]byte(__obf_f6f05d0629ad7fc3))
		defer __obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
		__obf_edd9fe48d39445e4.
			Read()
		if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
			__obf_2361f5214e84c60f.
				WriteRaw("null")
		} else {
			__obf_2361f5214e84c60f.
				WriteRaw(string(__obf_f6f05d0629ad7fc3))
		}
	}, func(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
		return len(*((*json.RawMessage)(__obf_753ab3fb72cdd659))) == 0
	}}
	__obf_e4a614f491c1bb0c[reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()] = __obf_c85b895560db528f
	__obf_e4a614f491c1bb0c[reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()] = __obf_c85b895560db528f
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_ef36b5509db9a3be(__obf_e4a614f491c1bb0c DecoderExtension) {
	__obf_e4a614f491c1bb0c[reflect2.TypeOfPtr((*any)(nil)).Elem()] = &__obf_c4dd4b1daf1fb426{func(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
		__obf_fb4aa039caaad786 := *((*any)(__obf_753ab3fb72cdd659))
		if __obf_fb4aa039caaad786 != nil && reflect.TypeOf(__obf_fb4aa039caaad786).Kind() == reflect.Ptr {
			__obf_edd9fe48d39445e4.
				ReadVal(__obf_fb4aa039caaad786)
			return
		}
		if __obf_edd9fe48d39445e4.WhatIsNext() == NumberValue {
			*((*any)(__obf_753ab3fb72cdd659)) = json.Number(__obf_edd9fe48d39445e4.__obf_8ccccffe17f4c818())
		} else {
			*((*any)(__obf_753ab3fb72cdd659)) = __obf_edd9fe48d39445e4.Read()
		}
	}}
}
func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_6cf02e65d7ac8e99() string {
	__obf_719bbdb676261f81 := __obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab.TagKey
	if __obf_719bbdb676261f81 == "" {
		return "json"
	}
	return __obf_719bbdb676261f81
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) RegisterExtension(__obf_e4a614f491c1bb0c Extension) {
	__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6 = append(__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6, __obf_e4a614f491c1bb0c)
	__obf_4cb2b46ddea01d73 := __obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab
	__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab = __obf_4cb2b46ddea01d73
}

type __obf_27fb8f7841a79d4b struct {
}

func (__obf_c85b895560db528f *__obf_27fb8f7841a79d4b) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteFloat32Lossy(*((*float32)(__obf_753ab3fb72cdd659)))
}

func (__obf_c85b895560db528f *__obf_27fb8f7841a79d4b) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*float32)(__obf_753ab3fb72cdd659)) == 0
}

type __obf_186f07839b62fe8f struct {
}

func (__obf_c85b895560db528f *__obf_186f07839b62fe8f) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteFloat64Lossy(*((*float64)(__obf_753ab3fb72cdd659)))
}

func (__obf_c85b895560db528f *__obf_186f07839b62fe8f) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*float64)(__obf_753ab3fb72cdd659)) == 0
}

// EnableLossyFloatMarshalling keeps 10**(-6) precision
// for float variables for better performance.
func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_d34b631463eb4202(__obf_e4a614f491c1bb0c EncoderExtension) {
	__obf_e4a614f491c1bb0c[ // for better performance
	reflect2.TypeOfPtr((*float32)(nil)).Elem()] = &__obf_27fb8f7841a79d4b{}
	__obf_e4a614f491c1bb0c[reflect2.TypeOfPtr((*float64)(nil)).Elem()] = &__obf_186f07839b62fe8f{}
}

type __obf_2398fec1b5ce21c8 struct {
}

func (__obf_c85b895560db528f *__obf_2398fec1b5ce21c8) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_2d944450d48e5810 := *((*string)(__obf_753ab3fb72cdd659))
	__obf_2361f5214e84c60f.
		WriteStringWithHTMLEscaped(__obf_2d944450d48e5810)
}

func (__obf_c85b895560db528f *__obf_2398fec1b5ce21c8) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return *((*string)(__obf_753ab3fb72cdd659)) == ""
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_98a068a278faf6e8(__obf_912608e689e59f9b EncoderExtension) {
	__obf_912608e689e59f9b[reflect2.TypeOfPtr((*string)(nil)).Elem()] = &__obf_2398fec1b5ce21c8{}
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_304467a78133c439() {
	__obf_f31c7837deea7cf1 = map[string]ValDecoder{}
	__obf_6b3006d32bbe7ad3 = map[string]ValDecoder{}
	*__obf_0c8065c219ccf0ab = *(__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab.Froze().(*__obf_3a25fb4c9a8342bb))
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) __obf_5a7a774979415647() {
	__obf_8061900e068c691e = map[string]ValEncoder{}
	__obf_319f7fa6b1b719bd = map[string]ValEncoder{}
	*__obf_0c8065c219ccf0ab = *(__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab.Froze().(*__obf_3a25fb4c9a8342bb))
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) MarshalToString(__obf_f967fc9e700d5e33 any) (string, error) {
	__obf_2361f5214e84c60f := __obf_0c8065c219ccf0ab.BorrowStream(nil)
	defer __obf_0c8065c219ccf0ab.ReturnStream(__obf_2361f5214e84c60f)
	__obf_2361f5214e84c60f.
		WriteVal(__obf_f967fc9e700d5e33)
	if __obf_2361f5214e84c60f.Error != nil {
		return "", __obf_2361f5214e84c60f.Error
	}
	return string(__obf_2361f5214e84c60f.Buffer()), nil
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) Marshal(__obf_f967fc9e700d5e33 any) ([]byte, error) {
	__obf_2361f5214e84c60f := __obf_0c8065c219ccf0ab.BorrowStream(nil)
	defer __obf_0c8065c219ccf0ab.ReturnStream(__obf_2361f5214e84c60f)
	__obf_2361f5214e84c60f.
		WriteVal(__obf_f967fc9e700d5e33)
	if __obf_2361f5214e84c60f.Error != nil {
		return nil, __obf_2361f5214e84c60f.Error
	}
	__obf_c5a13e5fd49848a8 := __obf_2361f5214e84c60f.Buffer()
	__obf_4cb2b46ddea01d73 := make([]byte, len(__obf_c5a13e5fd49848a8))
	copy(__obf_4cb2b46ddea01d73, __obf_c5a13e5fd49848a8)
	return __obf_4cb2b46ddea01d73, nil
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) MarshalIndent(__obf_f967fc9e700d5e33 any, __obf_f517983aa5897f07, __obf_30653456d76688b8 string) ([]byte, error) {
	if __obf_f517983aa5897f07 != "" {
		panic("prefix is not supported")
	}
	for _, __obf_c9c2210cc516a267 := range __obf_30653456d76688b8 {
		if __obf_c9c2210cc516a267 != ' ' {
			panic("indent can only be space")
		}
	}
	__obf_65b77c7bd66c6c45 := __obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab
	__obf_65b77c7bd66c6c45.
		IndentionStep = len(__obf_30653456d76688b8)
	return __obf_65b77c7bd66c6c45.__obf_c3344b4ada41a413(__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6).Marshal(__obf_f967fc9e700d5e33)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) UnmarshalFromString(__obf_2d944450d48e5810 string, __obf_f967fc9e700d5e33 any) error {
	__obf_905295f6bd29aafe := []byte(__obf_2d944450d48e5810)
	__obf_edd9fe48d39445e4 := __obf_0c8065c219ccf0ab.BorrowIterator(__obf_905295f6bd29aafe)
	defer __obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadVal(__obf_f967fc9e700d5e33)
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == 0 {
		if __obf_edd9fe48d39445e4.Error == io.EOF {
			return nil
		}
		return __obf_edd9fe48d39445e4.Error
	}
	__obf_edd9fe48d39445e4.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_edd9fe48d39445e4.Error
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) Get(__obf_905295f6bd29aafe []byte, __obf_b90e4ca332607787 ...any) Any {
	__obf_edd9fe48d39445e4 := __obf_0c8065c219ccf0ab.BorrowIterator(__obf_905295f6bd29aafe)
	defer __obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	return __obf_9979e7313b51aeca(__obf_edd9fe48d39445e4, __obf_b90e4ca332607787)
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) Unmarshal(__obf_905295f6bd29aafe []byte, __obf_f967fc9e700d5e33 any) error {
	__obf_edd9fe48d39445e4 := __obf_0c8065c219ccf0ab.BorrowIterator(__obf_905295f6bd29aafe)
	defer __obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		ReadVal(__obf_f967fc9e700d5e33)
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == 0 {
		if __obf_edd9fe48d39445e4.Error == io.EOF {
			return nil
		}
		return __obf_edd9fe48d39445e4.Error
	}
	__obf_edd9fe48d39445e4.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_edd9fe48d39445e4.Error
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) NewEncoder(__obf_b1699a146b20b28e io.Writer) *Encoder {
	__obf_2361f5214e84c60f := NewStream(__obf_0c8065c219ccf0ab, __obf_b1699a146b20b28e, 512)
	return &Encoder{__obf_2361f5214e84c60f}
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) NewDecoder(__obf_a539369e10727e42 io.Reader) *Decoder {
	__obf_edd9fe48d39445e4 := Parse(__obf_0c8065c219ccf0ab, __obf_a539369e10727e42, 512)
	return &Decoder{__obf_edd9fe48d39445e4}
}

func (__obf_0c8065c219ccf0ab *__obf_3a25fb4c9a8342bb) Valid(__obf_905295f6bd29aafe []byte) bool {
	__obf_edd9fe48d39445e4 := __obf_0c8065c219ccf0ab.BorrowIterator(__obf_905295f6bd29aafe)
	defer __obf_0c8065c219ccf0ab.ReturnIterator(__obf_edd9fe48d39445e4)
	__obf_edd9fe48d39445e4.
		Skip()
	return __obf_edd9fe48d39445e4.Error == nil
}
