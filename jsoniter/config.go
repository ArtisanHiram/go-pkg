package __obf_c7b79b12b33d8238

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
	MarshalToString(__obf_bc7196b82ffbf1d3 any) (string, error)
	Marshal(__obf_bc7196b82ffbf1d3 any) ([]byte, error)
	MarshalIndent(__obf_bc7196b82ffbf1d3 any, __obf_5de9ece8fa3a16e3, __obf_be0b45c15e61f86b string) ([]byte, error)
	UnmarshalFromString(__obf_a3eaafc22faf1644 string, __obf_bc7196b82ffbf1d3 any) error
	Unmarshal(__obf_1d34d01a9b83218a []byte, __obf_bc7196b82ffbf1d3 any) error
	Get(__obf_1d34d01a9b83218a []byte, __obf_5dde9fb4007294d4 ...any) Any
	NewEncoder(__obf_29b48e4cc46aebd6 io.Writer) *Encoder
	NewDecoder(__obf_801a7eebc4303617 io.Reader) *Decoder
	Valid(__obf_1d34d01a9b83218a []byte) bool
	RegisterExtension(__obf_b8f2f726e65c4d89 Extension)
	DecoderOf(__obf_edcded11af6ebc4c reflect2.Type) ValDecoder
	EncoderOf(__obf_edcded11af6ebc4c reflect2.Type) ValEncoder
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

type __obf_30fe5c95cabd69c0 struct {
	__obf_790a57f66978b3bc Config
	__obf_8571acf590409abc bool
	__obf_9f90ab8046926590 int
	__obf_82ca3cd044b095ae bool
	__obf_6e9ee582f04bfc43 bool
	__obf_b1dc099c042d4d26 bool
	__obf_a8556033d59ae33a *concurrent.Map
	__obf_ad1d36b2b67e17f0 *concurrent.Map
	__obf_3594368cedb76ac8 Extension
	__obf_1b8392ccffa3bed5 Extension
	__obf_3c82f4beb30882eb []Extension
	__obf_4409df5a50bd7d1b *sync.Pool
	__obf_4d02f79ac1d4d36b *sync.Pool
	__obf_3704f04b7ac67609 bool
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_ece0cf6130fbf668() {
	__obf_c52e0bcfad4b8b71.__obf_a8556033d59ae33a = concurrent.NewMap()
	__obf_c52e0bcfad4b8b71.__obf_ad1d36b2b67e17f0 = concurrent.NewMap()
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_84b44689116bd691(__obf_cb0de621f162c1ef uintptr, __obf_801f19702638809c ValDecoder) {
	__obf_c52e0bcfad4b8b71.__obf_a8556033d59ae33a.
		Store(__obf_cb0de621f162c1ef, __obf_801f19702638809c)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_fca11eaf15932cf2(__obf_cb0de621f162c1ef uintptr, __obf_c091c27480fae550 ValEncoder) {
	__obf_c52e0bcfad4b8b71.__obf_ad1d36b2b67e17f0.
		Store(__obf_cb0de621f162c1ef, __obf_c091c27480fae550)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_73853dd89dff1885(__obf_cb0de621f162c1ef uintptr) ValDecoder {
	__obf_801f19702638809c, __obf_22405d874667f998 := __obf_c52e0bcfad4b8b71.__obf_a8556033d59ae33a.Load(__obf_cb0de621f162c1ef)
	if __obf_22405d874667f998 {
		return __obf_801f19702638809c.(ValDecoder)
	}
	return nil
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_627bd91264fcb8f2(__obf_cb0de621f162c1ef uintptr) ValEncoder {
	__obf_c091c27480fae550, __obf_22405d874667f998 := __obf_c52e0bcfad4b8b71.__obf_ad1d36b2b67e17f0.Load(__obf_cb0de621f162c1ef)
	if __obf_22405d874667f998 {
		return __obf_c091c27480fae550.(ValEncoder)
	}
	return nil
}

var __obf_5a82191b59712916 = concurrent.NewMap()

func __obf_5c341b3862bbb97c(__obf_c52e0bcfad4b8b71 Config) *__obf_30fe5c95cabd69c0 {
	__obf_d6e2df4782353c64, __obf_22405d874667f998 := __obf_5a82191b59712916.Load(__obf_c52e0bcfad4b8b71)
	if __obf_22405d874667f998 {
		return __obf_d6e2df4782353c64.(*__obf_30fe5c95cabd69c0)
	}
	return nil
}

func __obf_6f8dea7bf018a80c(__obf_c52e0bcfad4b8b71 Config, __obf_30fe5c95cabd69c0 *__obf_30fe5c95cabd69c0) {
	__obf_5a82191b59712916.
		Store(__obf_c52e0bcfad4b8b71,

			// Froze forge API from config
			__obf_30fe5c95cabd69c0)
}

func (__obf_c52e0bcfad4b8b71 Config) Froze() API {
	__obf_c1b51f9c8b8a9878 := &__obf_30fe5c95cabd69c0{__obf_8571acf590409abc: __obf_c52e0bcfad4b8b71.SortMapKeys, __obf_9f90ab8046926590: __obf_c52e0bcfad4b8b71.IndentionStep, __obf_82ca3cd044b095ae: __obf_c52e0bcfad4b8b71.ObjectFieldMustBeSimpleString, __obf_6e9ee582f04bfc43: __obf_c52e0bcfad4b8b71.OnlyTaggedField, __obf_b1dc099c042d4d26: __obf_c52e0bcfad4b8b71.DisallowUnknownFields, __obf_3704f04b7ac67609: __obf_c52e0bcfad4b8b71.CaseSensitive}
	__obf_c1b51f9c8b8a9878.__obf_4409df5a50bd7d1b = &sync.Pool{
		New: func() any {
			return NewStream(__obf_c1b51f9c8b8a9878, nil, 512)
		},
	}
	__obf_c1b51f9c8b8a9878.__obf_4d02f79ac1d4d36b = &sync.Pool{
		New: func() any {
			return NewIterator(__obf_c1b51f9c8b8a9878)
		},
	}
	__obf_c1b51f9c8b8a9878.__obf_ece0cf6130fbf668()
	__obf_3594368cedb76ac8 := EncoderExtension{}
	__obf_1b8392ccffa3bed5 := DecoderExtension{}
	if __obf_c52e0bcfad4b8b71.MarshalFloatWith6Digits {
		__obf_c1b51f9c8b8a9878.__obf_bdf34cb6aae561ba(__obf_3594368cedb76ac8)
	}
	if __obf_c52e0bcfad4b8b71.EscapeHTML {
		__obf_c1b51f9c8b8a9878.__obf_3f302bac9ebc5e76(__obf_3594368cedb76ac8)
	}
	if __obf_c52e0bcfad4b8b71.UseNumber {
		__obf_c1b51f9c8b8a9878.__obf_3cb2d5265a5e80ea(__obf_1b8392ccffa3bed5)
	}
	if __obf_c52e0bcfad4b8b71.ValidateJsonRawMessage {
		__obf_c1b51f9c8b8a9878.__obf_bcee6676884660fe(__obf_3594368cedb76ac8)
	}
	__obf_c1b51f9c8b8a9878.__obf_3594368cedb76ac8 = __obf_3594368cedb76ac8
	__obf_c1b51f9c8b8a9878.__obf_1b8392ccffa3bed5 = __obf_1b8392ccffa3bed5
	__obf_c1b51f9c8b8a9878.__obf_790a57f66978b3bc = __obf_c52e0bcfad4b8b71
	return __obf_c1b51f9c8b8a9878
}

func (__obf_c52e0bcfad4b8b71 Config) __obf_ec5a9b5887aa3ec6(__obf_3c82f4beb30882eb []Extension) *__obf_30fe5c95cabd69c0 {
	__obf_c1b51f9c8b8a9878 := __obf_5c341b3862bbb97c(__obf_c52e0bcfad4b8b71)
	if __obf_c1b51f9c8b8a9878 != nil {
		return __obf_c1b51f9c8b8a9878
	}
	__obf_c1b51f9c8b8a9878 = __obf_c52e0bcfad4b8b71.Froze().(*__obf_30fe5c95cabd69c0)
	for _, __obf_b8f2f726e65c4d89 := range __obf_3c82f4beb30882eb {
		__obf_c1b51f9c8b8a9878.
			RegisterExtension(__obf_b8f2f726e65c4d89)
	}
	__obf_6f8dea7bf018a80c(__obf_c52e0bcfad4b8b71, __obf_c1b51f9c8b8a9878)
	return __obf_c1b51f9c8b8a9878
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_bcee6676884660fe(__obf_b8f2f726e65c4d89 EncoderExtension) {
	__obf_c091c27480fae550 := &__obf_fd208f64584e4e8f{func(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
		__obf_e8f5f5d4fc93347b := *(*json.RawMessage)(__obf_575c04f2b9d91315)
		__obf_0da8c843dad13139 := __obf_c52e0bcfad4b8b71.BorrowIterator([]byte(__obf_e8f5f5d4fc93347b))
		defer __obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
		__obf_0da8c843dad13139.
			Read()
		if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
			__obf_d8f50bcfe87d8b47.
				WriteRaw("null")
		} else {
			__obf_d8f50bcfe87d8b47.
				WriteRaw(string(__obf_e8f5f5d4fc93347b))
		}
	}, func(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
		return len(*((*json.RawMessage)(__obf_575c04f2b9d91315))) == 0
	}}
	__obf_b8f2f726e65c4d89[reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()] = __obf_c091c27480fae550
	__obf_b8f2f726e65c4d89[reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()] = __obf_c091c27480fae550
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_3cb2d5265a5e80ea(__obf_b8f2f726e65c4d89 DecoderExtension) {
	__obf_b8f2f726e65c4d89[reflect2.TypeOfPtr((*any)(nil)).Elem()] = &__obf_45b8a60966351e16{func(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
		__obf_c9edcb7f3bae4c1b := *((*any)(__obf_575c04f2b9d91315))
		if __obf_c9edcb7f3bae4c1b != nil && reflect.TypeOf(__obf_c9edcb7f3bae4c1b).Kind() == reflect.Ptr {
			__obf_0da8c843dad13139.
				ReadVal(__obf_c9edcb7f3bae4c1b)
			return
		}
		if __obf_0da8c843dad13139.WhatIsNext() == NumberValue {
			*((*any)(__obf_575c04f2b9d91315)) = json.Number(__obf_0da8c843dad13139.__obf_a678d14dd84cbb6b())
		} else {
			*((*any)(__obf_575c04f2b9d91315)) = __obf_0da8c843dad13139.Read()
		}
	}}
}
func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_2b09ab97bcd3be4f() string {
	__obf_da542d6e2b77dc17 := __obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc.TagKey
	if __obf_da542d6e2b77dc17 == "" {
		return "json"
	}
	return __obf_da542d6e2b77dc17
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) RegisterExtension(__obf_b8f2f726e65c4d89 Extension) {
	__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb = append(__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb, __obf_b8f2f726e65c4d89)
	__obf_a3c6c71d32db2eb1 := __obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc
	__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc = __obf_a3c6c71d32db2eb1
}

type __obf_dc203a0d80f885b0 struct {
}

func (__obf_c091c27480fae550 *__obf_dc203a0d80f885b0) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteFloat32Lossy(*((*float32)(__obf_575c04f2b9d91315)))
}

func (__obf_c091c27480fae550 *__obf_dc203a0d80f885b0) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*float32)(__obf_575c04f2b9d91315)) == 0
}

type __obf_07ef2459712d83ee struct {
}

func (__obf_c091c27480fae550 *__obf_07ef2459712d83ee) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteFloat64Lossy(*((*float64)(__obf_575c04f2b9d91315)))
}

func (__obf_c091c27480fae550 *__obf_07ef2459712d83ee) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*float64)(__obf_575c04f2b9d91315)) == 0
}

// EnableLossyFloatMarshalling keeps 10**(-6) precision
// for float variables for better performance.
func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_bdf34cb6aae561ba(__obf_b8f2f726e65c4d89 EncoderExtension) {
	__obf_b8f2f726e65c4d89[ // for better performance
	reflect2.TypeOfPtr((*float32)(nil)).Elem()] = &__obf_dc203a0d80f885b0{}
	__obf_b8f2f726e65c4d89[reflect2.TypeOfPtr((*float64)(nil)).Elem()] = &__obf_07ef2459712d83ee{}
}

type __obf_c07efd1007e74b09 struct {
}

func (__obf_c091c27480fae550 *__obf_c07efd1007e74b09) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_a3eaafc22faf1644 := *((*string)(__obf_575c04f2b9d91315))
	__obf_d8f50bcfe87d8b47.
		WriteStringWithHTMLEscaped(__obf_a3eaafc22faf1644)
}

func (__obf_c091c27480fae550 *__obf_c07efd1007e74b09) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return *((*string)(__obf_575c04f2b9d91315)) == ""
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_3f302bac9ebc5e76(__obf_3594368cedb76ac8 EncoderExtension) {
	__obf_3594368cedb76ac8[reflect2.TypeOfPtr((*string)(nil)).Elem()] = &__obf_c07efd1007e74b09{}
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_a04c2b03a6e6ae9b() {
	__obf_6117f44f95358fef = map[string]ValDecoder{}
	__obf_545223d75e12c929 = map[string]ValDecoder{}
	*__obf_c52e0bcfad4b8b71 = *(__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc.Froze().(*__obf_30fe5c95cabd69c0))
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) __obf_7f607e898cf7f663() {
	__obf_1cd2b048b2975dac = map[string]ValEncoder{}
	__obf_0d0da0d7a04287ff = map[string]ValEncoder{}
	*__obf_c52e0bcfad4b8b71 = *(__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc.Froze().(*__obf_30fe5c95cabd69c0))
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) MarshalToString(__obf_bc7196b82ffbf1d3 any) (string, error) {
	__obf_d8f50bcfe87d8b47 := __obf_c52e0bcfad4b8b71.BorrowStream(nil)
	defer __obf_c52e0bcfad4b8b71.ReturnStream(__obf_d8f50bcfe87d8b47)
	__obf_d8f50bcfe87d8b47.
		WriteVal(__obf_bc7196b82ffbf1d3)
	if __obf_d8f50bcfe87d8b47.Error != nil {
		return "", __obf_d8f50bcfe87d8b47.Error
	}
	return string(__obf_d8f50bcfe87d8b47.Buffer()), nil
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) Marshal(__obf_bc7196b82ffbf1d3 any) ([]byte, error) {
	__obf_d8f50bcfe87d8b47 := __obf_c52e0bcfad4b8b71.BorrowStream(nil)
	defer __obf_c52e0bcfad4b8b71.ReturnStream(__obf_d8f50bcfe87d8b47)
	__obf_d8f50bcfe87d8b47.
		WriteVal(__obf_bc7196b82ffbf1d3)
	if __obf_d8f50bcfe87d8b47.Error != nil {
		return nil, __obf_d8f50bcfe87d8b47.Error
	}
	__obf_b8a287f4c212fc4f := __obf_d8f50bcfe87d8b47.Buffer()
	__obf_a3c6c71d32db2eb1 := make([]byte, len(__obf_b8a287f4c212fc4f))
	copy(__obf_a3c6c71d32db2eb1, __obf_b8a287f4c212fc4f)
	return __obf_a3c6c71d32db2eb1, nil
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) MarshalIndent(__obf_bc7196b82ffbf1d3 any, __obf_5de9ece8fa3a16e3, __obf_be0b45c15e61f86b string) ([]byte, error) {
	if __obf_5de9ece8fa3a16e3 != "" {
		panic("prefix is not supported")
	}
	for _, __obf_464eaebd868bf0fe := range __obf_be0b45c15e61f86b {
		if __obf_464eaebd868bf0fe != ' ' {
			panic("indent can only be space")
		}
	}
	__obf_61796c5f334fc27d := __obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc
	__obf_61796c5f334fc27d.
		IndentionStep = len(__obf_be0b45c15e61f86b)
	return __obf_61796c5f334fc27d.__obf_ec5a9b5887aa3ec6(__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb).Marshal(__obf_bc7196b82ffbf1d3)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) UnmarshalFromString(__obf_a3eaafc22faf1644 string, __obf_bc7196b82ffbf1d3 any) error {
	__obf_1d34d01a9b83218a := []byte(__obf_a3eaafc22faf1644)
	__obf_0da8c843dad13139 := __obf_c52e0bcfad4b8b71.BorrowIterator(__obf_1d34d01a9b83218a)
	defer __obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadVal(__obf_bc7196b82ffbf1d3)
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == 0 {
		if __obf_0da8c843dad13139.Error == io.EOF {
			return nil
		}
		return __obf_0da8c843dad13139.Error
	}
	__obf_0da8c843dad13139.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_0da8c843dad13139.Error
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) Get(__obf_1d34d01a9b83218a []byte, __obf_5dde9fb4007294d4 ...any) Any {
	__obf_0da8c843dad13139 := __obf_c52e0bcfad4b8b71.BorrowIterator(__obf_1d34d01a9b83218a)
	defer __obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	return __obf_fedb6216061761ae(__obf_0da8c843dad13139, __obf_5dde9fb4007294d4)
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) Unmarshal(__obf_1d34d01a9b83218a []byte, __obf_bc7196b82ffbf1d3 any) error {
	__obf_0da8c843dad13139 := __obf_c52e0bcfad4b8b71.BorrowIterator(__obf_1d34d01a9b83218a)
	defer __obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		ReadVal(__obf_bc7196b82ffbf1d3)
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == 0 {
		if __obf_0da8c843dad13139.Error == io.EOF {
			return nil
		}
		return __obf_0da8c843dad13139.Error
	}
	__obf_0da8c843dad13139.
		ReportError("Unmarshal", "there are bytes left after unmarshal")
	return __obf_0da8c843dad13139.Error
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) NewEncoder(__obf_29b48e4cc46aebd6 io.Writer) *Encoder {
	__obf_d8f50bcfe87d8b47 := NewStream(__obf_c52e0bcfad4b8b71, __obf_29b48e4cc46aebd6, 512)
	return &Encoder{__obf_d8f50bcfe87d8b47}
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) NewDecoder(__obf_801a7eebc4303617 io.Reader) *Decoder {
	__obf_0da8c843dad13139 := Parse(__obf_c52e0bcfad4b8b71, __obf_801a7eebc4303617, 512)
	return &Decoder{__obf_0da8c843dad13139}
}

func (__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0) Valid(__obf_1d34d01a9b83218a []byte) bool {
	__obf_0da8c843dad13139 := __obf_c52e0bcfad4b8b71.BorrowIterator(__obf_1d34d01a9b83218a)
	defer __obf_c52e0bcfad4b8b71.ReturnIterator(__obf_0da8c843dad13139)
	__obf_0da8c843dad13139.
		Skip()
	return __obf_0da8c843dad13139.Error == nil
}
