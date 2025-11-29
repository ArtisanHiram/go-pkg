package __obf_789dc33d47f4ab2c

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

const __obf_ad58a90ebc3ce5a8 = ^uint(0)
const __obf_4361d2b9e7968597 = int(__obf_ad58a90ebc3ce5a8 >> 1)
const __obf_5471c30b1fab6f86 = -__obf_4361d2b9e7968597 - 1

// RegisterFuzzyDecoders decode input from PHP with tolerance.
// It will handle string/number auto conversation, and treat empty [] as empty struct.
func RegisterFuzzyDecoders() {
	jsoniter.RegisterExtension(&__obf_261c18649b164d66{})
	jsoniter.RegisterTypeDecoder("string", &__obf_a63cadfbec9f1eec{})
	jsoniter.RegisterTypeDecoder("float32", &__obf_0237154de208a535{})
	jsoniter.RegisterTypeDecoder("float64", &__obf_a9d446cd0615ce6a{})
	jsoniter.RegisterTypeDecoder("int", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(__obf_4361d2b9e7968597) || __obf_0fb751fedbd6c970 < float64(__obf_5471c30b1fab6f86) {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode int", "exceed range")
				return
			}
			*((*int)(__obf_4280d1a4ac42e464)) = int(__obf_0fb751fedbd6c970)
		} else {
			*((*int)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadInt()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(__obf_ad58a90ebc3ce5a8) || __obf_0fb751fedbd6c970 < 0 {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode uint", "exceed range")
				return
			}
			*((*uint)(__obf_4280d1a4ac42e464)) = uint(__obf_0fb751fedbd6c970)
		} else {
			*((*uint)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadUint()
		}
	}})
	jsoniter.RegisterTypeDecoder("int8", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxInt8) || __obf_0fb751fedbd6c970 < float64(math.MinInt8) {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode int8", "exceed range")
				return
			}
			*((*int8)(__obf_4280d1a4ac42e464)) = int8(__obf_0fb751fedbd6c970)
		} else {
			*((*int8)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadInt8()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint8", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxUint8) || __obf_0fb751fedbd6c970 < 0 {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode uint8", "exceed range")
				return
			}
			*((*uint8)(__obf_4280d1a4ac42e464)) = uint8(__obf_0fb751fedbd6c970)
		} else {
			*((*uint8)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadUint8()
		}
	}})
	jsoniter.RegisterTypeDecoder("int16", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxInt16) || __obf_0fb751fedbd6c970 < float64(math.MinInt16) {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode int16", "exceed range")
				return
			}
			*((*int16)(__obf_4280d1a4ac42e464)) = int16(__obf_0fb751fedbd6c970)
		} else {
			*((*int16)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadInt16()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint16", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxUint16) || __obf_0fb751fedbd6c970 < 0 {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode uint16", "exceed range")
				return
			}
			*((*uint16)(__obf_4280d1a4ac42e464)) = uint16(__obf_0fb751fedbd6c970)
		} else {
			*((*uint16)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadUint16()
		}
	}})
	jsoniter.RegisterTypeDecoder("int32", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxInt32) || __obf_0fb751fedbd6c970 < float64(math.MinInt32) {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode int32", "exceed range")
				return
			}
			*((*int32)(__obf_4280d1a4ac42e464)) = int32(__obf_0fb751fedbd6c970)
		} else {
			*((*int32)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadInt32()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint32", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxUint32) || __obf_0fb751fedbd6c970 < 0 {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode uint32", "exceed range")
				return
			}
			*((*uint32)(__obf_4280d1a4ac42e464)) = uint32(__obf_0fb751fedbd6c970)
		} else {
			*((*uint32)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadUint32()
		}
	}})
	jsoniter.RegisterTypeDecoder("int64", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxInt64) || __obf_0fb751fedbd6c970 < float64(math.MinInt64) {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode int64", "exceed range")
				return
			}
			*((*int64)(__obf_4280d1a4ac42e464)) = int64(__obf_0fb751fedbd6c970)
		} else {
			*((*int64)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadInt64()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint64", &__obf_10ac9a5db659756e{func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
		if __obf_d017eb5c434c50b7 {
			__obf_0fb751fedbd6c970 := __obf_e97e2ea2e3d1b0a2.ReadFloat64()
			if __obf_0fb751fedbd6c970 > float64(math.MaxUint64) || __obf_0fb751fedbd6c970 < 0 {
				__obf_e97e2ea2e3d1b0a2.
					ReportError("fuzzy decode uint64", "exceed range")
				return
			}
			*((*uint64)(__obf_4280d1a4ac42e464)) = uint64(__obf_0fb751fedbd6c970)
		} else {
			*((*uint64)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadUint64()
		}
	}})
}

type __obf_261c18649b164d66 struct {
	jsoniter.DummyExtension
}

func (__obf_369258fb13d82ac5 *__obf_261c18649b164d66) DecorateDecoder(__obf_487781debab9b36c reflect2.Type, __obf_50ab65d221b126b0 jsoniter.ValDecoder) jsoniter.ValDecoder {
	if __obf_487781debab9b36c.Kind() == reflect.Struct || __obf_487781debab9b36c.Kind() == reflect.Map {
		return &__obf_011a882534a4cd1a{__obf_50ab65d221b126b0}
	}
	return __obf_50ab65d221b126b0
}

type __obf_011a882534a4cd1a struct {
	__obf_2ae89546343db054 jsoniter.ValDecoder
}

func (__obf_50ab65d221b126b0 *__obf_011a882534a4cd1a) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	if __obf_e97e2ea2e3d1b0a2.WhatIsNext() == jsoniter.ArrayValue {
		__obf_e97e2ea2e3d1b0a2.
			Skip()
		__obf_a1b1f9b17228269a := __obf_e97e2ea2e3d1b0a2.Pool().BorrowIterator([]byte("{}"))
		defer __obf_e97e2ea2e3d1b0a2.Pool().ReturnIterator(__obf_a1b1f9b17228269a)
		__obf_50ab65d221b126b0.__obf_2ae89546343db054.
			Decode(__obf_4280d1a4ac42e464, __obf_a1b1f9b17228269a)
	} else {
		__obf_50ab65d221b126b0.__obf_2ae89546343db054.
			Decode(__obf_4280d1a4ac42e464, __obf_e97e2ea2e3d1b0a2)
	}
}

type __obf_a63cadfbec9f1eec struct {
}

func (__obf_50ab65d221b126b0 *__obf_a63cadfbec9f1eec) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	__obf_b22b8bfb4cccfd07 := __obf_e97e2ea2e3d1b0a2.WhatIsNext()
	switch __obf_b22b8bfb4cccfd07 {
	case jsoniter.NumberValue:
		var __obf_442e3f098a589686 json.Number
		__obf_e97e2ea2e3d1b0a2.
			ReadVal(&__obf_442e3f098a589686)
		*((*string)(__obf_4280d1a4ac42e464)) = string(__obf_442e3f098a589686)
	case jsoniter.StringValue:
		*((*string)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadString()
	case jsoniter.NilValue:
		__obf_e97e2ea2e3d1b0a2.
			Skip()
		*((*string)(__obf_4280d1a4ac42e464)) = ""
	default:
		__obf_e97e2ea2e3d1b0a2.
			ReportError("fuzzyStringDecoder", "not number or string")
	}
}

type __obf_10ac9a5db659756e struct {
	__obf_283631aa90ec16b1 func(__obf_d017eb5c434c50b7 bool, __obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator)
}

func (__obf_50ab65d221b126b0 *__obf_10ac9a5db659756e) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	__obf_b22b8bfb4cccfd07 := __obf_e97e2ea2e3d1b0a2.WhatIsNext()
	var __obf_f0bd3a090b87546b string
	switch __obf_b22b8bfb4cccfd07 {
	case jsoniter.NumberValue:
		var __obf_442e3f098a589686 json.Number
		__obf_e97e2ea2e3d1b0a2.
			ReadVal(&__obf_442e3f098a589686)
		__obf_f0bd3a090b87546b = string(__obf_442e3f098a589686)
	case jsoniter.StringValue:
		__obf_f0bd3a090b87546b = __obf_e97e2ea2e3d1b0a2.ReadString()
	case jsoniter.BoolValue:
		if __obf_e97e2ea2e3d1b0a2.ReadBool() {
			__obf_f0bd3a090b87546b = "1"
		} else {
			__obf_f0bd3a090b87546b = "0"
		}
	case jsoniter.NilValue:
		__obf_e97e2ea2e3d1b0a2.
			Skip()
		__obf_f0bd3a090b87546b = "0"
	default:
		__obf_e97e2ea2e3d1b0a2.
			ReportError("fuzzyIntegerDecoder", "not number or string")
	}
	if len(__obf_f0bd3a090b87546b) == 0 {
		__obf_f0bd3a090b87546b = "0"
	}
	__obf_a1b1f9b17228269a := __obf_e97e2ea2e3d1b0a2.Pool().BorrowIterator([]byte(__obf_f0bd3a090b87546b))
	defer __obf_e97e2ea2e3d1b0a2.Pool().ReturnIterator(__obf_a1b1f9b17228269a)
	__obf_d017eb5c434c50b7 := strings.IndexByte(__obf_f0bd3a090b87546b, '.') != -1
	__obf_50ab65d221b126b0.__obf_283631aa90ec16b1(__obf_d017eb5c434c50b7, __obf_4280d1a4ac42e464, __obf_a1b1f9b17228269a)
	if __obf_a1b1f9b17228269a.Error != nil && __obf_a1b1f9b17228269a.Error != io.EOF {
		__obf_e97e2ea2e3d1b0a2.
			Error = __obf_a1b1f9b17228269a.Error
	}
}

type __obf_0237154de208a535 struct {
}

func (__obf_50ab65d221b126b0 *__obf_0237154de208a535) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	__obf_b22b8bfb4cccfd07 := __obf_e97e2ea2e3d1b0a2.WhatIsNext()
	var __obf_f0bd3a090b87546b string
	switch __obf_b22b8bfb4cccfd07 {
	case jsoniter.NumberValue:
		*((*float32)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadFloat32()
	case jsoniter.StringValue:
		__obf_f0bd3a090b87546b = __obf_e97e2ea2e3d1b0a2.ReadString()
		__obf_a1b1f9b17228269a := __obf_e97e2ea2e3d1b0a2.Pool().BorrowIterator([]byte(__obf_f0bd3a090b87546b))
		defer __obf_e97e2ea2e3d1b0a2.Pool().ReturnIterator(__obf_a1b1f9b17228269a)
		*((*float32)(__obf_4280d1a4ac42e464)) = __obf_a1b1f9b17228269a.ReadFloat32()
		if __obf_a1b1f9b17228269a.Error != nil && __obf_a1b1f9b17228269a.Error != io.EOF {
			__obf_e97e2ea2e3d1b0a2.
				Error = __obf_a1b1f9b17228269a.Error
		}
	case jsoniter.BoolValue:
		// support bool to float32
		if __obf_e97e2ea2e3d1b0a2.ReadBool() {
			*((*float32)(__obf_4280d1a4ac42e464)) = 1
		} else {
			*((*float32)(__obf_4280d1a4ac42e464)) = 0
		}
	case jsoniter.NilValue:
		__obf_e97e2ea2e3d1b0a2.
			Skip()
		*((*float32)(__obf_4280d1a4ac42e464)) = 0
	default:
		__obf_e97e2ea2e3d1b0a2.
			ReportError("fuzzyFloat32Decoder", "not number or string")
	}
}

type __obf_a9d446cd0615ce6a struct {
}

func (__obf_50ab65d221b126b0 *__obf_a9d446cd0615ce6a) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	__obf_b22b8bfb4cccfd07 := __obf_e97e2ea2e3d1b0a2.WhatIsNext()
	var __obf_f0bd3a090b87546b string
	switch __obf_b22b8bfb4cccfd07 {
	case jsoniter.NumberValue:
		*((*float64)(__obf_4280d1a4ac42e464)) = __obf_e97e2ea2e3d1b0a2.ReadFloat64()
	case jsoniter.StringValue:
		__obf_f0bd3a090b87546b = __obf_e97e2ea2e3d1b0a2.ReadString()
		__obf_a1b1f9b17228269a := __obf_e97e2ea2e3d1b0a2.Pool().BorrowIterator([]byte(__obf_f0bd3a090b87546b))
		defer __obf_e97e2ea2e3d1b0a2.Pool().ReturnIterator(__obf_a1b1f9b17228269a)
		*((*float64)(__obf_4280d1a4ac42e464)) = __obf_a1b1f9b17228269a.ReadFloat64()
		if __obf_a1b1f9b17228269a.Error != nil && __obf_a1b1f9b17228269a.Error != io.EOF {
			__obf_e97e2ea2e3d1b0a2.
				Error = __obf_a1b1f9b17228269a.Error
		}
	case jsoniter.BoolValue:
		// support bool to float64
		if __obf_e97e2ea2e3d1b0a2.ReadBool() {
			*((*float64)(__obf_4280d1a4ac42e464)) = 1
		} else {
			*((*float64)(__obf_4280d1a4ac42e464)) = 0
		}
	case jsoniter.NilValue:
		__obf_e97e2ea2e3d1b0a2.
			Skip()
		*((*float64)(__obf_4280d1a4ac42e464)) = 0
	default:
		__obf_e97e2ea2e3d1b0a2.
			ReportError("fuzzyFloat64Decoder", "not number or string")
	}
}
