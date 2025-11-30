package __obf_1f22c6b8dfc77bff

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

const __obf_8ffbd7224d91e8d3 = ^uint(0)
const __obf_1efee6edc2d0d39c = int(__obf_8ffbd7224d91e8d3 >> 1)
const __obf_953e10b1af5b8343 = -__obf_1efee6edc2d0d39c - 1

// RegisterFuzzyDecoders decode input from PHP with tolerance.
// It will handle string/number auto conversation, and treat empty [] as empty struct.
func RegisterFuzzyDecoders() {
	jsoniter.RegisterExtension(&__obf_417e5add1dcd79e7{})
	jsoniter.RegisterTypeDecoder("string", &__obf_65929a4dea6c84f3{})
	jsoniter.RegisterTypeDecoder("float32", &__obf_83c8a3913b2451e1{})
	jsoniter.RegisterTypeDecoder("float64", &__obf_4c2ac12073937402{})
	jsoniter.RegisterTypeDecoder("int", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(__obf_1efee6edc2d0d39c) || __obf_c26a6b64dc1ad5d5 < float64(__obf_953e10b1af5b8343) {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode int", "exceed range")
				return
			}
			*((*int)(__obf_a5271c65f4ae17af)) = int(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*int)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadInt()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(__obf_8ffbd7224d91e8d3) || __obf_c26a6b64dc1ad5d5 < 0 {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode uint", "exceed range")
				return
			}
			*((*uint)(__obf_a5271c65f4ae17af)) = uint(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*uint)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadUint()
		}
	}})
	jsoniter.RegisterTypeDecoder("int8", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxInt8) || __obf_c26a6b64dc1ad5d5 < float64(math.MinInt8) {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode int8", "exceed range")
				return
			}
			*((*int8)(__obf_a5271c65f4ae17af)) = int8(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*int8)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadInt8()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint8", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxUint8) || __obf_c26a6b64dc1ad5d5 < 0 {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode uint8", "exceed range")
				return
			}
			*((*uint8)(__obf_a5271c65f4ae17af)) = uint8(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*uint8)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadUint8()
		}
	}})
	jsoniter.RegisterTypeDecoder("int16", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxInt16) || __obf_c26a6b64dc1ad5d5 < float64(math.MinInt16) {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode int16", "exceed range")
				return
			}
			*((*int16)(__obf_a5271c65f4ae17af)) = int16(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*int16)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadInt16()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint16", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxUint16) || __obf_c26a6b64dc1ad5d5 < 0 {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode uint16", "exceed range")
				return
			}
			*((*uint16)(__obf_a5271c65f4ae17af)) = uint16(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*uint16)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadUint16()
		}
	}})
	jsoniter.RegisterTypeDecoder("int32", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxInt32) || __obf_c26a6b64dc1ad5d5 < float64(math.MinInt32) {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode int32", "exceed range")
				return
			}
			*((*int32)(__obf_a5271c65f4ae17af)) = int32(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*int32)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadInt32()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint32", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxUint32) || __obf_c26a6b64dc1ad5d5 < 0 {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode uint32", "exceed range")
				return
			}
			*((*uint32)(__obf_a5271c65f4ae17af)) = uint32(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*uint32)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadUint32()
		}
	}})
	jsoniter.RegisterTypeDecoder("int64", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxInt64) || __obf_c26a6b64dc1ad5d5 < float64(math.MinInt64) {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode int64", "exceed range")
				return
			}
			*((*int64)(__obf_a5271c65f4ae17af)) = int64(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*int64)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadInt64()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint64", &__obf_5f205394b3e6d8ef{func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
		if __obf_e364c91b5209017f {
			__obf_c26a6b64dc1ad5d5 := __obf_d021dab62946a708.ReadFloat64()
			if __obf_c26a6b64dc1ad5d5 > float64(math.MaxUint64) || __obf_c26a6b64dc1ad5d5 < 0 {
				__obf_d021dab62946a708.
					ReportError("fuzzy decode uint64", "exceed range")
				return
			}
			*((*uint64)(__obf_a5271c65f4ae17af)) = uint64(__obf_c26a6b64dc1ad5d5)
		} else {
			*((*uint64)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadUint64()
		}
	}})
}

type __obf_417e5add1dcd79e7 struct {
	jsoniter.DummyExtension
}

func (__obf_9382a03d04135598 *__obf_417e5add1dcd79e7) DecorateDecoder(__obf_66f97a1172af63eb reflect2.Type, __obf_571d1fed69e547a5 jsoniter.ValDecoder) jsoniter.ValDecoder {
	if __obf_66f97a1172af63eb.Kind() == reflect.Struct || __obf_66f97a1172af63eb.Kind() == reflect.Map {
		return &__obf_3b0883e44c3b1eed{__obf_571d1fed69e547a5}
	}
	return __obf_571d1fed69e547a5
}

type __obf_3b0883e44c3b1eed struct {
	__obf_58c1837ebdb9c576 jsoniter.ValDecoder
}

func (__obf_571d1fed69e547a5 *__obf_3b0883e44c3b1eed) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	if __obf_d021dab62946a708.WhatIsNext() == jsoniter.ArrayValue {
		__obf_d021dab62946a708.
			Skip()
		__obf_12c35f823a513ecb := __obf_d021dab62946a708.Pool().BorrowIterator([]byte("{}"))
		defer __obf_d021dab62946a708.Pool().ReturnIterator(__obf_12c35f823a513ecb)
		__obf_571d1fed69e547a5.__obf_58c1837ebdb9c576.
			Decode(__obf_a5271c65f4ae17af, __obf_12c35f823a513ecb)
	} else {
		__obf_571d1fed69e547a5.__obf_58c1837ebdb9c576.
			Decode(__obf_a5271c65f4ae17af, __obf_d021dab62946a708)
	}
}

type __obf_65929a4dea6c84f3 struct {
}

func (__obf_571d1fed69e547a5 *__obf_65929a4dea6c84f3) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	__obf_a38053be8cb310f3 := __obf_d021dab62946a708.WhatIsNext()
	switch __obf_a38053be8cb310f3 {
	case jsoniter.NumberValue:
		var __obf_c820f0e92ec2b899 json.Number
		__obf_d021dab62946a708.
			ReadVal(&__obf_c820f0e92ec2b899)
		*((*string)(__obf_a5271c65f4ae17af)) = string(__obf_c820f0e92ec2b899)
	case jsoniter.StringValue:
		*((*string)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadString()
	case jsoniter.NilValue:
		__obf_d021dab62946a708.
			Skip()
		*((*string)(__obf_a5271c65f4ae17af)) = ""
	default:
		__obf_d021dab62946a708.
			ReportError("fuzzyStringDecoder", "not number or string")
	}
}

type __obf_5f205394b3e6d8ef struct {
	__obf_3a64b80aee78e7df func(__obf_e364c91b5209017f bool, __obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator)
}

func (__obf_571d1fed69e547a5 *__obf_5f205394b3e6d8ef) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	__obf_a38053be8cb310f3 := __obf_d021dab62946a708.WhatIsNext()
	var __obf_5a1b5343e640b4d0 string
	switch __obf_a38053be8cb310f3 {
	case jsoniter.NumberValue:
		var __obf_c820f0e92ec2b899 json.Number
		__obf_d021dab62946a708.
			ReadVal(&__obf_c820f0e92ec2b899)
		__obf_5a1b5343e640b4d0 = string(__obf_c820f0e92ec2b899)
	case jsoniter.StringValue:
		__obf_5a1b5343e640b4d0 = __obf_d021dab62946a708.ReadString()
	case jsoniter.BoolValue:
		if __obf_d021dab62946a708.ReadBool() {
			__obf_5a1b5343e640b4d0 = "1"
		} else {
			__obf_5a1b5343e640b4d0 = "0"
		}
	case jsoniter.NilValue:
		__obf_d021dab62946a708.
			Skip()
		__obf_5a1b5343e640b4d0 = "0"
	default:
		__obf_d021dab62946a708.
			ReportError("fuzzyIntegerDecoder", "not number or string")
	}
	if len(__obf_5a1b5343e640b4d0) == 0 {
		__obf_5a1b5343e640b4d0 = "0"
	}
	__obf_12c35f823a513ecb := __obf_d021dab62946a708.Pool().BorrowIterator([]byte(__obf_5a1b5343e640b4d0))
	defer __obf_d021dab62946a708.Pool().ReturnIterator(__obf_12c35f823a513ecb)
	__obf_e364c91b5209017f := strings.IndexByte(__obf_5a1b5343e640b4d0, '.') != -1
	__obf_571d1fed69e547a5.__obf_3a64b80aee78e7df(__obf_e364c91b5209017f, __obf_a5271c65f4ae17af, __obf_12c35f823a513ecb)
	if __obf_12c35f823a513ecb.Error != nil && __obf_12c35f823a513ecb.Error != io.EOF {
		__obf_d021dab62946a708.
			Error = __obf_12c35f823a513ecb.Error
	}
}

type __obf_83c8a3913b2451e1 struct {
}

func (__obf_571d1fed69e547a5 *__obf_83c8a3913b2451e1) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	__obf_a38053be8cb310f3 := __obf_d021dab62946a708.WhatIsNext()
	var __obf_5a1b5343e640b4d0 string
	switch __obf_a38053be8cb310f3 {
	case jsoniter.NumberValue:
		*((*float32)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadFloat32()
	case jsoniter.StringValue:
		__obf_5a1b5343e640b4d0 = __obf_d021dab62946a708.ReadString()
		__obf_12c35f823a513ecb := __obf_d021dab62946a708.Pool().BorrowIterator([]byte(__obf_5a1b5343e640b4d0))
		defer __obf_d021dab62946a708.Pool().ReturnIterator(__obf_12c35f823a513ecb)
		*((*float32)(__obf_a5271c65f4ae17af)) = __obf_12c35f823a513ecb.ReadFloat32()
		if __obf_12c35f823a513ecb.Error != nil && __obf_12c35f823a513ecb.Error != io.EOF {
			__obf_d021dab62946a708.
				Error = __obf_12c35f823a513ecb.Error
		}
	case jsoniter.BoolValue:
		// support bool to float32
		if __obf_d021dab62946a708.ReadBool() {
			*((*float32)(__obf_a5271c65f4ae17af)) = 1
		} else {
			*((*float32)(__obf_a5271c65f4ae17af)) = 0
		}
	case jsoniter.NilValue:
		__obf_d021dab62946a708.
			Skip()
		*((*float32)(__obf_a5271c65f4ae17af)) = 0
	default:
		__obf_d021dab62946a708.
			ReportError("fuzzyFloat32Decoder", "not number or string")
	}
}

type __obf_4c2ac12073937402 struct {
}

func (__obf_571d1fed69e547a5 *__obf_4c2ac12073937402) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	__obf_a38053be8cb310f3 := __obf_d021dab62946a708.WhatIsNext()
	var __obf_5a1b5343e640b4d0 string
	switch __obf_a38053be8cb310f3 {
	case jsoniter.NumberValue:
		*((*float64)(__obf_a5271c65f4ae17af)) = __obf_d021dab62946a708.ReadFloat64()
	case jsoniter.StringValue:
		__obf_5a1b5343e640b4d0 = __obf_d021dab62946a708.ReadString()
		__obf_12c35f823a513ecb := __obf_d021dab62946a708.Pool().BorrowIterator([]byte(__obf_5a1b5343e640b4d0))
		defer __obf_d021dab62946a708.Pool().ReturnIterator(__obf_12c35f823a513ecb)
		*((*float64)(__obf_a5271c65f4ae17af)) = __obf_12c35f823a513ecb.ReadFloat64()
		if __obf_12c35f823a513ecb.Error != nil && __obf_12c35f823a513ecb.Error != io.EOF {
			__obf_d021dab62946a708.
				Error = __obf_12c35f823a513ecb.Error
		}
	case jsoniter.BoolValue:
		// support bool to float64
		if __obf_d021dab62946a708.ReadBool() {
			*((*float64)(__obf_a5271c65f4ae17af)) = 1
		} else {
			*((*float64)(__obf_a5271c65f4ae17af)) = 0
		}
	case jsoniter.NilValue:
		__obf_d021dab62946a708.
			Skip()
		*((*float64)(__obf_a5271c65f4ae17af)) = 0
	default:
		__obf_d021dab62946a708.
			ReportError("fuzzyFloat64Decoder", "not number or string")
	}
}
