package __obf_030d39f7a8de96c6

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
	MarshalToString(__obf_41b184145cfea25b any) (string, error)
	Marshal(__obf_41b184145cfea25b any) ([]byte, error)
	MarshalIndent(__obf_41b184145cfea25b any, __obf_834e1679b8081f46, __obf_2ee2bfac47d1636e string) ([]byte, error)
	UnmarshalFromString(__obf_428c3b4a95725c4a string, __obf_41b184145cfea25b any) error
	Unmarshal(__obf_c28f0e30eceb6d4a []byte, __obf_41b184145cfea25b any) error
	Get(__obf_c28f0e30eceb6d4a []byte, __obf_f77a9507782b919d ...any) Any
	NewEncoder(__obf_ae48508e054620bd io.Writer) *Encoder
	NewDecoder(__obf_7582c70e81d83895 io.Reader) *Decoder
	Valid(__obf_c28f0e30eceb6d4a []byte) bool
	RegisterExtension(__obf_621544a57e52000f Extension)
	DecoderOf(__obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder
	EncoderOf(__obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder
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

type __obf_81639538813814ff struct {
	__obf_fb97917857c6c8da Config
	__obf_e01b84a9a1b0f266 bool
	__obf_063460864bd2e9ef int
	__obf_8c3b82d5a5e45b5c bool
	__obf_67b8a9dee101ecbf bool
	__obf_5a45cc14ef9b494e bool
	__obf_e4055d7a0b1f1441 *concurrent.Map
	__obf_0623c185205156d7 *concurrent.Map
	__obf_7db1f0a55b319e45 Extension
	__obf_6313bfb9c913bd50 Extension
	__obf_6621255bc1f80c8e []Extension
	__obf_e195df3f11517e00 *sync.Pool
	__obf_4522de1f25b34c0b *sync.Pool
	__obf_af13e3babb780a4e bool
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_d961e64682eee099() {
	__obf_679611dc56ff465b.__obf_e4055d7a0b1f1441 = concurrent.NewMap()
	__obf_679611dc56ff465b.__obf_0623c185205156d7 = concurrent.NewMap()
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_0599525836bd5d66(__obf_5ddcbb0301e0cf46 uintptr, __obf_11a3f28116164d09 ValDecoder) {
	__obf_679611dc56ff465b.__obf_e4055d7a0b1f1441.
		Store(__obf_5ddcbb0301e0cf46, __obf_11a3f28116164d09)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_ec0f7f63f3dc6639(__obf_5ddcbb0301e0cf46 uintptr, __obf_008f61596d7e9523 ValEncoder) {
	__obf_679611dc56ff465b.__obf_0623c185205156d7.
		Store(__obf_5ddcbb0301e0cf46, __obf_008f61596d7e9523)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_dff673ee34435a0e(__obf_5ddcbb0301e0cf46 uintptr) ValDecoder {
	__obf_11a3f28116164d09, __obf_ccdc4ee8b9d2f52f := __obf_679611dc56ff465b.__obf_e4055d7a0b1f1441.Load(__obf_5ddcbb0301e0cf46)
	if __obf_ccdc4ee8b9d2f52f {
		return __obf_11a3f28116164d09.(ValDecoder)
	}
	return nil
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_f3694d09db16d770(__obf_5ddcbb0301e0cf46 uintptr) ValEncoder {
	__obf_008f61596d7e9523, __obf_ccdc4ee8b9d2f52f := __obf_679611dc56ff465b.__obf_0623c185205156d7.Load(__obf_5ddcbb0301e0cf46)
	if __obf_ccdc4ee8b9d2f52f {
		return __obf_008f61596d7e9523.(ValEncoder)
	}
	return nil
}

var __obf_6997ae10d198b8bb = concurrent.NewMap()

func __obf_d645fa76596a28e7(__obf_679611dc56ff465b Config) *__obf_81639538813814ff {
	__obf_eefa92392cc2442c, __obf_ccdc4ee8b9d2f52f := __obf_6997ae10d198b8bb.Load(__obf_679611dc56ff465b)
	if __obf_ccdc4ee8b9d2f52f {
		return __obf_eefa92392cc2442c.(*__obf_81639538813814ff)
	}
	return nil
}

func __obf_6e7c631b32cf29bb(__obf_679611dc56ff465b Config, __obf_81639538813814ff *__obf_81639538813814ff) {
	__obf_6997ae10d198b8bb.
		Store(__obf_679611dc56ff465b,

			// Froze forge API from config
			__obf_81639538813814ff)
}

func (__obf_679611dc56ff465b Config) Froze() API {
	__obf_f0a9df63f1b75800 := &__obf_81639538813814ff{__obf_e01b84a9a1b0f266: __obf_679611dc56ff465b.SortMapKeys, __obf_063460864bd2e9ef: __obf_679611dc56ff465b.IndentionStep, __obf_8c3b82d5a5e45b5c: __obf_679611dc56ff465b.ObjectFieldMustBeSimpleString, __obf_67b8a9dee101ecbf: __obf_679611dc56ff465b.OnlyTaggedField, __obf_5a45cc14ef9b494e: __obf_679611dc56ff465b.DisallowUnknownFields, __obf_af13e3babb780a4e: __obf_679611dc56ff465b.CaseSensitive}
	__obf_f0a9df63f1b75800.__obf_e195df3f11517e00 = &sync.Pool{
		New: func() any {
			return NewStream(__obf_f0a9df63f1b75800, nil, 512)
		},
	}
	__obf_f0a9df63f1b75800.__obf_4522de1f25b34c0b = &sync.Pool{
		New: func() any {
			return NewIterator(__obf_f0a9df63f1b75800)
		},
	}
	__obf_f0a9df63f1b75800.__obf_d961e64682eee099()
	__obf_7db1f0a55b319e45 := EncoderExtension{}
	__obf_6313bfb9c913bd50 := DecoderExtension{}
	if __obf_679611dc56ff465b.MarshalFloatWith6Digits {
		__obf_f0a9df63f1b75800.__obf_64fd14f9ac25d2d8(__obf_7db1f0a55b319e45)
	}
	if __obf_679611dc56ff465b.EscapeHTML {
		__obf_f0a9df63f1b75800.__obf_ea893fc01ce71521(__obf_7db1f0a55b319e45)
	}
	if __obf_679611dc56ff465b.UseNumber {
		__obf_f0a9df63f1b75800.__obf_683a502e6458a095(__obf_6313bfb9c913bd50)
	}
	if __obf_679611dc56ff465b.ValidateJsonRawMessage {
		__obf_f0a9df63f1b75800.__obf_968025c6c46d9c19(__obf_7db1f0a55b319e45)
	}
	__obf_f0a9df63f1b75800.__obf_7db1f0a55b319e45 = __obf_7db1f0a55b319e45
	__obf_f0a9df63f1b75800.__obf_6313bfb9c913bd50 = __obf_6313bfb9c913bd50
	__obf_f0a9df63f1b75800.__obf_fb97917857c6c8da = __obf_679611dc56ff465b
	return __obf_f0a9df63f1b75800
}

func (__obf_679611dc56ff465b Config) __obf_38802709ddf21ec6(__obf_6621255bc1f80c8e []Extension) *__obf_81639538813814ff {
	__obf_f0a9df63f1b75800 := __obf_d645fa76596a28e7(__obf_679611dc56ff465b)
	if __obf_f0a9df63f1b75800 != nil {
		return __obf_f0a9df63f1b75800
	}
	__obf_f0a9df63f1b75800 = __obf_679611dc56ff465b.Froze().(*__obf_81639538813814ff)
	for _, __obf_621544a57e52000f := range __obf_6621255bc1f80c8e {
		__obf_f0a9df63f1b75800.
			RegisterExtension(__obf_621544a57e52000f)
	}
	__obf_6e7c631b32cf29bb(__obf_679611dc56ff465b, __obf_f0a9df63f1b75800)
	return __obf_f0a9df63f1b75800
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_968025c6c46d9c19(__obf_621544a57e52000f EncoderExtension) {
	__obf_008f61596d7e9523 := &__obf_f97fbd765d3a66b4{func(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
		__obf_71f88de23c04aecb := *(*json.RawMessage)(__obf_dbbf371b91136494)
		__obf_4ab56a99965952e7 := __obf_679611dc56ff465b.BorrowIterator([]byte(__obf_71f88de23c04aecb))
		defer __obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
		__obf_4ab56a99965952e7.
			Read()
		if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
			__obf_8a2c51fe22775655.
				WriteRaw("null")
		} else {
			__obf_8a2c51fe22775655.
				WriteRaw(string(__obf_71f88de23c04aecb))
		}
	}, func(__obf_dbbf371b91136494 unsafe.Pointer) bool {
		return len(*((*json.RawMessage)(__obf_dbbf371b91136494))) == 0
	}}
	__obf_621544a57e52000f[reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()] = __obf_008f61596d7e9523
	__obf_621544a57e52000f[reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()] = __obf_008f61596d7e9523
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_683a502e6458a095(__obf_621544a57e52000f DecoderExtension) {
	__obf_621544a57e52000f[reflect2.TypeOfPtr((*any)(nil)).Elem()] = &__obf_3f78bd1ffd8b466d{func(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
		__obf_fdc21cb863bf3c1e := *((*any)(__obf_dbbf371b91136494))
		if __obf_fdc21cb863bf3c1e != nil && reflect.TypeOf(__obf_fdc21cb863bf3c1e).Kind() == reflect.Ptr {
			__obf_4ab56a99965952e7.
				ReadVal(__obf_fdc21cb863bf3c1e)
			return
		}
		if __obf_4ab56a99965952e7.WhatIsNext() == NumberValue {
			*((*any)(__obf_dbbf371b91136494)) = json.Number(__obf_4ab56a99965952e7.__obf_0ba08d27498a0d84())
		} else {
			*((*any)(__obf_dbbf371b91136494)) = __obf_4ab56a99965952e7.Read()
		}
	}}
}
func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_79fd0a7f8863c590() string {
	__obf_da44d0f7410dd554 := __obf_679611dc56ff465b.__obf_fb97917857c6c8da.TagKey
	if __obf_da44d0f7410dd554 == "" {
		return "json"
	}
	return __obf_da44d0f7410dd554
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) RegisterExtension(__obf_621544a57e52000f Extension) {
	__obf_679611dc56ff465b.__obf_6621255bc1f80c8e = append(__obf_679611dc56ff465b.__obf_6621255bc1f80c8e, __obf_621544a57e52000f)
	__obf_390a561ab5257f8f := __obf_679611dc56ff465b.__obf_fb97917857c6c8da
	__obf_679611dc56ff465b.__obf_fb97917857c6c8da = __obf_390a561ab5257f8f
}

type __obf_df39aa767def6a9d struct {
}

func (__obf_008f61596d7e9523 *__obf_df39aa767def6a9d) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteFloat32Lossy(*((*float32)(__obf_dbbf371b91136494)))
}

func (__obf_008f61596d7e9523 *__obf_df39aa767def6a9d) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*float32)(__obf_dbbf371b91136494)) == 0
}

type __obf_38179ae2d6730aa4 struct {
}

func (__obf_008f61596d7e9523 *__obf_38179ae2d6730aa4) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteFloat64Lossy(*((*float64)(__obf_dbbf371b91136494)))
}

func (__obf_008f61596d7e9523 *__obf_38179ae2d6730aa4) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*float64)(__obf_dbbf371b91136494)) == 0
}

// EnableLossyFloatMarshalling keeps 10**(-6) precision
// for float variables for better performance.
func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_64fd14f9ac25d2d8(__obf_621544a57e52000f EncoderExtension) {
	__obf_621544a57e52000f[ // for better performance
	reflect2.TypeOfPtr((*float32)(nil)).Elem()] = &__obf_df39aa767def6a9d{}
	__obf_621544a57e52000f[reflect2.TypeOfPtr((*float64)(nil)).Elem()] = &__obf_38179ae2d6730aa4{}
}

type __obf_9ec8c5eeea47f3b5 struct {
}

func (__obf_008f61596d7e9523 *__obf_9ec8c5eeea47f3b5) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_428c3b4a95725c4a := *((*string)(__obf_dbbf371b91136494))
	__obf_8a2c51fe22775655.
		WriteStringWithHTMLEscaped(__obf_428c3b4a95725c4a)
}

func (__obf_008f61596d7e9523 *__obf_9ec8c5eeea47f3b5) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return *((*string)(__obf_dbbf371b91136494)) == ""
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_ea893fc01ce71521(__obf_7db1f0a55b319e45 EncoderExtension) {
	__obf_7db1f0a55b319e45[reflect2.TypeOfPtr((*string)(nil)).Elem()] = &__obf_9ec8c5eeea47f3b5{}
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_252003baf66ab4d8() {
	__obf_af43d5efcf6ae11d = map[string]ValDecoder{}
	__obf_3b686c29105543bc = map[string]ValDecoder{}
	*__obf_679611dc56ff465b = *(__obf_679611dc56ff465b.__obf_fb97917857c6c8da.Froze().(*__obf_81639538813814ff))
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) __obf_9c92e1a3e8f44950() {
	__obf_8f43affe58f43115 = map[string]ValEncoder{}
	__obf_133872148688e4c1 = map[string]ValEncoder{}
	*__obf_679611dc56ff465b = *(__obf_679611dc56ff465b.__obf_fb97917857c6c8da.Froze().(*__obf_81639538813814ff))
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) MarshalToString(__obf_41b184145cfea25b any) (string, error) {
	__obf_8a2c51fe22775655 := __obf_679611dc56ff465b.BorrowStream(nil)
	defer __obf_679611dc56ff465b.ReturnStream(__obf_8a2c51fe22775655)
	__obf_8a2c51fe22775655.
		WriteVal(__obf_41b184145cfea25b)
	if __obf_8a2c51fe22775655.Error != nil {
		return "", __obf_8a2c51fe22775655.Error
	}
	return string(__obf_8a2c51fe22775655.Buffer()), nil
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) Marshal(__obf_41b184145cfea25b any) ([]byte, error) {
	__obf_8a2c51fe22775655 := __obf_679611dc56ff465b.BorrowStream(nil)
	defer __obf_679611dc56ff465b.ReturnStream(__obf_8a2c51fe22775655)
	__obf_8a2c51fe22775655.
		WriteVal(__obf_41b184145cfea25b)
	if __obf_8a2c51fe22775655.Error != nil {
		return nil, __obf_8a2c51fe22775655.Error
	}
	__obf_5ae3a53d1322d642 := __obf_8a2c51fe22775655.Buffer()
	__obf_390a561ab5257f8f := make([]byte, len(__obf_5ae3a53d1322d642))
	copy(__obf_390a561ab5257f8f, __obf_5ae3a53d1322d642)
	return __obf_390a561ab5257f8f, nil
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) MarshalIndent(__obf_41b184145cfea25b any, __obf_834e1679b8081f46, __obf_2ee2bfac47d1636e string) ([]byte, error) {
	if __obf_834e1679b8081f46 != "" {
		panic("prefix is not supported")
	}
	for _, __obf_9aca26302d2c5887 := range __obf_2ee2bfac47d1636e {
		if __obf_9aca26302d2c5887 != ' ' {
			panic("indent can only be space")
		}
	}
	__obf_4966bed951f8c4a3 := __obf_679611dc56ff465b.__obf_fb97917857c6c8da
	__obf_4966bed951f8c4a3.
		IndentionStep = len(__obf_2ee2bfac47d1636e)
	return __obf_4966bed951f8c4a3.__obf_38802709ddf21ec6(__obf_679611dc56ff465b.__obf_6621255bc1f80c8e).Marshal(__obf_41b184145cfea25b)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) UnmarshalFromString(__obf_428c3b4a95725c4a string, __obf_41b184145cfea25b any) error {
	__obf_c28f0e30eceb6d4a := []byte(__obf_428c3b4a95725c4a)
	__obf_4ab56a99965952e7 := __obf_679611dc56ff465b.BorrowIterator(__obf_c28f0e30eceb6d4a)
	defer __obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadVal(__obf_41b184145cfea25b)
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == 0 {
		if __obf_4ab56a99965952e7.Error == io.EOF {
			return nil
		}
		return __obf_4ab56a99965952e7.Error
	}
	__obf_4ab56a99965952e7.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_4ab56a99965952e7.Error
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) Get(__obf_c28f0e30eceb6d4a []byte, __obf_f77a9507782b919d ...any) Any {
	__obf_4ab56a99965952e7 := __obf_679611dc56ff465b.BorrowIterator(__obf_c28f0e30eceb6d4a)
	defer __obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	return __obf_4cd15630eeb40546(__obf_4ab56a99965952e7, __obf_f77a9507782b919d)
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) Unmarshal(__obf_c28f0e30eceb6d4a []byte, __obf_41b184145cfea25b any) error {
	__obf_4ab56a99965952e7 := __obf_679611dc56ff465b.BorrowIterator(__obf_c28f0e30eceb6d4a)
	defer __obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		ReadVal(__obf_41b184145cfea25b)
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == 0 {
		if __obf_4ab56a99965952e7.Error == io.EOF {
			return nil
		}
		return __obf_4ab56a99965952e7.Error
	}
	__obf_4ab56a99965952e7.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_4ab56a99965952e7.Error
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) NewEncoder(__obf_ae48508e054620bd io.Writer) *Encoder {
	__obf_8a2c51fe22775655 := NewStream(__obf_679611dc56ff465b, __obf_ae48508e054620bd, 512)
	return &Encoder{__obf_8a2c51fe22775655}
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) NewDecoder(__obf_7582c70e81d83895 io.Reader) *Decoder {
	__obf_4ab56a99965952e7 := Parse(__obf_679611dc56ff465b, __obf_7582c70e81d83895, 512)
	return &Decoder{__obf_4ab56a99965952e7}
}

func (__obf_679611dc56ff465b *__obf_81639538813814ff) Valid(__obf_c28f0e30eceb6d4a []byte) bool {
	__obf_4ab56a99965952e7 := __obf_679611dc56ff465b.BorrowIterator(__obf_c28f0e30eceb6d4a)
	defer __obf_679611dc56ff465b.ReturnIterator(__obf_4ab56a99965952e7)
	__obf_4ab56a99965952e7.
		Skip()
	return __obf_4ab56a99965952e7.Error == nil
}
