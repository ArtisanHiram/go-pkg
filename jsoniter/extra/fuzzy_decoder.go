package __obf_9a397ef96534ad45

import (
	"encoding/json"
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"github.com/modern-go/reflect2"
	"io"
	"math"
	"reflect"
	"strings"
	"unsafe"
)

const __obf_ce4be2113f8316d7 = ^uint(0)
const __obf_37464f9148a140d5 = int(__obf_ce4be2113f8316d7 >> 1)
const __obf_6644e31e952f150d = -__obf_37464f9148a140d5 - 1

// RegisterFuzzyDecoders decode input from PHP with tolerance.
// It will handle string/number auto conversation, and treat empty [] as empty struct.
func RegisterFuzzyDecoders() {
	jsoniter.RegisterExtension(&__obf_92049a5f2e29e12b{})
	jsoniter.RegisterTypeDecoder("string", &__obf_e39c14125eb180a8{})
	jsoniter.RegisterTypeDecoder("float32", &__obf_baf19b5333f84370{})
	jsoniter.RegisterTypeDecoder("float64", &__obf_77bcb70d03f33429{})
	jsoniter.RegisterTypeDecoder("int", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(__obf_37464f9148a140d5) || __obf_05b16ae687fabbb8 < float64(__obf_6644e31e952f150d) {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode int", "exceed range")
				return
			}
			*((*int)(__obf_77e2593d34cc3a6a)) = int(__obf_05b16ae687fabbb8)
		} else {
			*((*int)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadInt()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(__obf_ce4be2113f8316d7) || __obf_05b16ae687fabbb8 < 0 {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode uint", "exceed range")
				return
			}
			*((*uint)(__obf_77e2593d34cc3a6a)) = uint(__obf_05b16ae687fabbb8)
		} else {
			*((*uint)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadUint()
		}
	}})
	jsoniter.RegisterTypeDecoder("int8", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxInt8) || __obf_05b16ae687fabbb8 < float64(math.MinInt8) {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode int8", "exceed range")
				return
			}
			*((*int8)(__obf_77e2593d34cc3a6a)) = int8(__obf_05b16ae687fabbb8)
		} else {
			*((*int8)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadInt8()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint8", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxUint8) || __obf_05b16ae687fabbb8 < 0 {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode uint8", "exceed range")
				return
			}
			*((*uint8)(__obf_77e2593d34cc3a6a)) = uint8(__obf_05b16ae687fabbb8)
		} else {
			*((*uint8)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadUint8()
		}
	}})
	jsoniter.RegisterTypeDecoder("int16", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxInt16) || __obf_05b16ae687fabbb8 < float64(math.MinInt16) {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode int16", "exceed range")
				return
			}
			*((*int16)(__obf_77e2593d34cc3a6a)) = int16(__obf_05b16ae687fabbb8)
		} else {
			*((*int16)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadInt16()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint16", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxUint16) || __obf_05b16ae687fabbb8 < 0 {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode uint16", "exceed range")
				return
			}
			*((*uint16)(__obf_77e2593d34cc3a6a)) = uint16(__obf_05b16ae687fabbb8)
		} else {
			*((*uint16)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadUint16()
		}
	}})
	jsoniter.RegisterTypeDecoder("int32", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxInt32) || __obf_05b16ae687fabbb8 < float64(math.MinInt32) {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode int32", "exceed range")
				return
			}
			*((*int32)(__obf_77e2593d34cc3a6a)) = int32(__obf_05b16ae687fabbb8)
		} else {
			*((*int32)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadInt32()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint32", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxUint32) || __obf_05b16ae687fabbb8 < 0 {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode uint32", "exceed range")
				return
			}
			*((*uint32)(__obf_77e2593d34cc3a6a)) = uint32(__obf_05b16ae687fabbb8)
		} else {
			*((*uint32)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadUint32()
		}
	}})
	jsoniter.RegisterTypeDecoder("int64", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxInt64) || __obf_05b16ae687fabbb8 < float64(math.MinInt64) {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode int64", "exceed range")
				return
			}
			*((*int64)(__obf_77e2593d34cc3a6a)) = int64(__obf_05b16ae687fabbb8)
		} else {
			*((*int64)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadInt64()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint64", &__obf_65058ba6db155ee5{func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
		if __obf_9d597ff475786bf4 {
			__obf_05b16ae687fabbb8 := __obf_f2099f19d22a8175.ReadFloat64()
			if __obf_05b16ae687fabbb8 > float64(math.MaxUint64) || __obf_05b16ae687fabbb8 < 0 {
				__obf_f2099f19d22a8175.
					ReportError("fuzzy decode uint64", "exceed range")
				return
			}
			*((*uint64)(__obf_77e2593d34cc3a6a)) = uint64(__obf_05b16ae687fabbb8)
		} else {
			*((*uint64)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadUint64()
		}
	}})
}

type __obf_92049a5f2e29e12b struct {
	jsoniter.DummyExtension
}

func (__obf_454cba947156c7ed *__obf_92049a5f2e29e12b) DecorateDecoder(__obf_36cc3da433275e85 reflect2.Type, __obf_9cf25f620ff1cfc6 jsoniter.ValDecoder) jsoniter.ValDecoder {
	if __obf_36cc3da433275e85.Kind() == reflect.Struct || __obf_36cc3da433275e85.Kind() == reflect.Map {
		return &__obf_dac96f41f4f0abf3{__obf_9cf25f620ff1cfc6}
	}
	return __obf_9cf25f620ff1cfc6
}

type __obf_dac96f41f4f0abf3 struct {
	__obf_b25ae9db73fcde81 jsoniter.ValDecoder
}

func (__obf_9cf25f620ff1cfc6 *__obf_dac96f41f4f0abf3) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	if __obf_f2099f19d22a8175.WhatIsNext() == jsoniter.ArrayValue {
		__obf_f2099f19d22a8175.
			Skip()
		__obf_845aeb519f1c9c91 := __obf_f2099f19d22a8175.Pool().BorrowIterator([]byte("{}"))
		defer __obf_f2099f19d22a8175.Pool().ReturnIterator(__obf_845aeb519f1c9c91)
		__obf_9cf25f620ff1cfc6.__obf_b25ae9db73fcde81.
			Decode(__obf_77e2593d34cc3a6a, __obf_845aeb519f1c9c91)
	} else {
		__obf_9cf25f620ff1cfc6.__obf_b25ae9db73fcde81.
			Decode(__obf_77e2593d34cc3a6a, __obf_f2099f19d22a8175)
	}
}

type __obf_e39c14125eb180a8 struct {
}

func (__obf_9cf25f620ff1cfc6 *__obf_e39c14125eb180a8) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	__obf_e44d5b7ab327f6a2 := __obf_f2099f19d22a8175.WhatIsNext()
	switch __obf_e44d5b7ab327f6a2 {
	case jsoniter.NumberValue:
		var __obf_2875372c2e4650c0 json.Number
		__obf_f2099f19d22a8175.
			ReadVal(&__obf_2875372c2e4650c0)
		*((*string)(__obf_77e2593d34cc3a6a)) = string(__obf_2875372c2e4650c0)
	case jsoniter.StringValue:
		*((*string)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadString()
	case jsoniter.NilValue:
		__obf_f2099f19d22a8175.
			Skip()
		*((*string)(__obf_77e2593d34cc3a6a)) = ""
	default:
		__obf_f2099f19d22a8175.
			ReportError("fuzzyStringDecoder", "not number or string")
	}
}

type __obf_65058ba6db155ee5 struct {
	__obf_8ed3acd7ec510198 func(__obf_9d597ff475786bf4 bool, __obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator)
}

func (__obf_9cf25f620ff1cfc6 *__obf_65058ba6db155ee5) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	__obf_e44d5b7ab327f6a2 := __obf_f2099f19d22a8175.WhatIsNext()
	var __obf_b1c6f61eac17763a string
	switch __obf_e44d5b7ab327f6a2 {
	case jsoniter.NumberValue:
		var __obf_2875372c2e4650c0 json.Number
		__obf_f2099f19d22a8175.
			ReadVal(&__obf_2875372c2e4650c0)
		__obf_b1c6f61eac17763a = string(__obf_2875372c2e4650c0)
	case jsoniter.StringValue:
		__obf_b1c6f61eac17763a = __obf_f2099f19d22a8175.ReadString()
	case jsoniter.BoolValue:
		if __obf_f2099f19d22a8175.ReadBool() {
			__obf_b1c6f61eac17763a = "1"
		} else {
			__obf_b1c6f61eac17763a = "0"
		}
	case jsoniter.NilValue:
		__obf_f2099f19d22a8175.
			Skip()
		__obf_b1c6f61eac17763a = "0"
	default:
		__obf_f2099f19d22a8175.
			ReportError("fuzzyIntegerDecoder", "not number or string")
	}
	if len(__obf_b1c6f61eac17763a) == 0 {
		__obf_b1c6f61eac17763a = "0"
	}
	__obf_845aeb519f1c9c91 := __obf_f2099f19d22a8175.Pool().BorrowIterator([]byte(__obf_b1c6f61eac17763a))
	defer __obf_f2099f19d22a8175.Pool().ReturnIterator(__obf_845aeb519f1c9c91)
	__obf_9d597ff475786bf4 := strings.IndexByte(__obf_b1c6f61eac17763a, '.') != -1
	__obf_9cf25f620ff1cfc6.__obf_8ed3acd7ec510198(__obf_9d597ff475786bf4, __obf_77e2593d34cc3a6a, __obf_845aeb519f1c9c91)
	if __obf_845aeb519f1c9c91.Error != nil && __obf_845aeb519f1c9c91.Error != io.EOF {
		__obf_f2099f19d22a8175.
			Error = __obf_845aeb519f1c9c91.Error
	}
}

type __obf_baf19b5333f84370 struct {
}

func (__obf_9cf25f620ff1cfc6 *__obf_baf19b5333f84370) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	__obf_e44d5b7ab327f6a2 := __obf_f2099f19d22a8175.WhatIsNext()
	var __obf_b1c6f61eac17763a string
	switch __obf_e44d5b7ab327f6a2 {
	case jsoniter.NumberValue:
		*((*float32)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadFloat32()
	case jsoniter.StringValue:
		__obf_b1c6f61eac17763a = __obf_f2099f19d22a8175.ReadString()
		__obf_845aeb519f1c9c91 := __obf_f2099f19d22a8175.Pool().BorrowIterator([]byte(__obf_b1c6f61eac17763a))
		defer __obf_f2099f19d22a8175.Pool().ReturnIterator(__obf_845aeb519f1c9c91)
		*((*float32)(__obf_77e2593d34cc3a6a)) = __obf_845aeb519f1c9c91.ReadFloat32()
		if __obf_845aeb519f1c9c91.Error != nil && __obf_845aeb519f1c9c91.Error != io.EOF {
			__obf_f2099f19d22a8175.
				Error = __obf_845aeb519f1c9c91.Error
		}
	case jsoniter.BoolValue:
		// support bool to float32
		if __obf_f2099f19d22a8175.ReadBool() {
			*((*float32)(__obf_77e2593d34cc3a6a)) = 1
		} else {
			*((*float32)(__obf_77e2593d34cc3a6a)) = 0
		}
	case jsoniter.NilValue:
		__obf_f2099f19d22a8175.
			Skip()
		*((*float32)(__obf_77e2593d34cc3a6a)) = 0
	default:
		__obf_f2099f19d22a8175.
			ReportError("fuzzyFloat32Decoder", "not number or string")
	}
}

type __obf_77bcb70d03f33429 struct {
}

func (__obf_9cf25f620ff1cfc6 *__obf_77bcb70d03f33429) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	__obf_e44d5b7ab327f6a2 := __obf_f2099f19d22a8175.WhatIsNext()
	var __obf_b1c6f61eac17763a string
	switch __obf_e44d5b7ab327f6a2 {
	case jsoniter.NumberValue:
		*((*float64)(__obf_77e2593d34cc3a6a)) = __obf_f2099f19d22a8175.ReadFloat64()
	case jsoniter.StringValue:
		__obf_b1c6f61eac17763a = __obf_f2099f19d22a8175.ReadString()
		__obf_845aeb519f1c9c91 := __obf_f2099f19d22a8175.Pool().BorrowIterator([]byte(__obf_b1c6f61eac17763a))
		defer __obf_f2099f19d22a8175.Pool().ReturnIterator(__obf_845aeb519f1c9c91)
		*((*float64)(__obf_77e2593d34cc3a6a)) = __obf_845aeb519f1c9c91.ReadFloat64()
		if __obf_845aeb519f1c9c91.Error != nil && __obf_845aeb519f1c9c91.Error != io.EOF {
			__obf_f2099f19d22a8175.
				Error = __obf_845aeb519f1c9c91.Error
		}
	case jsoniter.BoolValue:
		// support bool to float64
		if __obf_f2099f19d22a8175.ReadBool() {
			*((*float64)(__obf_77e2593d34cc3a6a)) = 1
		} else {
			*((*float64)(__obf_77e2593d34cc3a6a)) = 0
		}
	case jsoniter.NilValue:
		__obf_f2099f19d22a8175.
			Skip()
		*((*float64)(__obf_77e2593d34cc3a6a)) = 0
	default:
		__obf_f2099f19d22a8175.
			ReportError("fuzzyFloat64Decoder", "not number or string")
	}
}
