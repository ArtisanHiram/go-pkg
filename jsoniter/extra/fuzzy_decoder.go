package __obf_eed9c5a643743c33

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

const __obf_a6cee2815965b472 = ^uint(0)
const __obf_19ee0cb418ef09ad = int(__obf_a6cee2815965b472 >> 1)
const __obf_8ca1157b15c33f3f = -__obf_19ee0cb418ef09ad - 1

// RegisterFuzzyDecoders decode input from PHP with tolerance.
// It will handle string/number auto conversation, and treat empty [] as empty struct.
func RegisterFuzzyDecoders() {
	jsoniter.RegisterExtension(&__obf_a10c8c219584974e{})
	jsoniter.RegisterTypeDecoder("string", &__obf_cab3ffafee538e6b{})
	jsoniter.RegisterTypeDecoder("float32", &__obf_51d82e4459cdbe7e{})
	jsoniter.RegisterTypeDecoder("float64", &__obf_e243f1ba7a4e8782{})
	jsoniter.RegisterTypeDecoder("int", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(__obf_19ee0cb418ef09ad) || __obf_34fe357f6baf2b37 < float64(__obf_8ca1157b15c33f3f) {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode int", "exceed range")
				return
			}
			*((*int)(__obf_e88eec03e62b77ec)) = int(__obf_34fe357f6baf2b37)
		} else {
			*((*int)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadInt()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(__obf_a6cee2815965b472) || __obf_34fe357f6baf2b37 < 0 {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode uint", "exceed range")
				return
			}
			*((*uint)(__obf_e88eec03e62b77ec)) = uint(__obf_34fe357f6baf2b37)
		} else {
			*((*uint)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadUint()
		}
	}})
	jsoniter.RegisterTypeDecoder("int8", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxInt8) || __obf_34fe357f6baf2b37 < float64(math.MinInt8) {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode int8", "exceed range")
				return
			}
			*((*int8)(__obf_e88eec03e62b77ec)) = int8(__obf_34fe357f6baf2b37)
		} else {
			*((*int8)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadInt8()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint8", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxUint8) || __obf_34fe357f6baf2b37 < 0 {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode uint8", "exceed range")
				return
			}
			*((*uint8)(__obf_e88eec03e62b77ec)) = uint8(__obf_34fe357f6baf2b37)
		} else {
			*((*uint8)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadUint8()
		}
	}})
	jsoniter.RegisterTypeDecoder("int16", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxInt16) || __obf_34fe357f6baf2b37 < float64(math.MinInt16) {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode int16", "exceed range")
				return
			}
			*((*int16)(__obf_e88eec03e62b77ec)) = int16(__obf_34fe357f6baf2b37)
		} else {
			*((*int16)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadInt16()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint16", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxUint16) || __obf_34fe357f6baf2b37 < 0 {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode uint16", "exceed range")
				return
			}
			*((*uint16)(__obf_e88eec03e62b77ec)) = uint16(__obf_34fe357f6baf2b37)
		} else {
			*((*uint16)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadUint16()
		}
	}})
	jsoniter.RegisterTypeDecoder("int32", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxInt32) || __obf_34fe357f6baf2b37 < float64(math.MinInt32) {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode int32", "exceed range")
				return
			}
			*((*int32)(__obf_e88eec03e62b77ec)) = int32(__obf_34fe357f6baf2b37)
		} else {
			*((*int32)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadInt32()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint32", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxUint32) || __obf_34fe357f6baf2b37 < 0 {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode uint32", "exceed range")
				return
			}
			*((*uint32)(__obf_e88eec03e62b77ec)) = uint32(__obf_34fe357f6baf2b37)
		} else {
			*((*uint32)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadUint32()
		}
	}})
	jsoniter.RegisterTypeDecoder("int64", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxInt64) || __obf_34fe357f6baf2b37 < float64(math.MinInt64) {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode int64", "exceed range")
				return
			}
			*((*int64)(__obf_e88eec03e62b77ec)) = int64(__obf_34fe357f6baf2b37)
		} else {
			*((*int64)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadInt64()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint64", &__obf_e0ca1fe42424825b{func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
		if __obf_f424b9c08d3cfad6 {
			__obf_34fe357f6baf2b37 := __obf_338eca60ccc15d82.ReadFloat64()
			if __obf_34fe357f6baf2b37 > float64(math.MaxUint64) || __obf_34fe357f6baf2b37 < 0 {
				__obf_338eca60ccc15d82.
					ReportError("fuzzy decode uint64", "exceed range")
				return
			}
			*((*uint64)(__obf_e88eec03e62b77ec)) = uint64(__obf_34fe357f6baf2b37)
		} else {
			*((*uint64)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadUint64()
		}
	}})
}

type __obf_a10c8c219584974e struct {
	jsoniter.DummyExtension
}

func (__obf_7ab69ccfb0d084de *__obf_a10c8c219584974e) DecorateDecoder(__obf_4db0904e845b7200 reflect2.Type, __obf_bdb89ac21da4a38c jsoniter.ValDecoder) jsoniter.ValDecoder {
	if __obf_4db0904e845b7200.Kind() == reflect.Struct || __obf_4db0904e845b7200.Kind() == reflect.Map {
		return &__obf_7e88cb05ebe96daf{__obf_bdb89ac21da4a38c}
	}
	return __obf_bdb89ac21da4a38c
}

type __obf_7e88cb05ebe96daf struct {
	__obf_6349c847813ef1e4 jsoniter.ValDecoder
}

func (__obf_bdb89ac21da4a38c *__obf_7e88cb05ebe96daf) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	if __obf_338eca60ccc15d82.WhatIsNext() == jsoniter.ArrayValue {
		__obf_338eca60ccc15d82.
			Skip()
		__obf_d91719946777be7e := __obf_338eca60ccc15d82.Pool().BorrowIterator([]byte("{}"))
		defer __obf_338eca60ccc15d82.Pool().ReturnIterator(__obf_d91719946777be7e)
		__obf_bdb89ac21da4a38c.__obf_6349c847813ef1e4.
			Decode(__obf_e88eec03e62b77ec, __obf_d91719946777be7e)
	} else {
		__obf_bdb89ac21da4a38c.__obf_6349c847813ef1e4.
			Decode(__obf_e88eec03e62b77ec, __obf_338eca60ccc15d82)
	}
}

type __obf_cab3ffafee538e6b struct {
}

func (__obf_bdb89ac21da4a38c *__obf_cab3ffafee538e6b) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	__obf_b2c92a9f9ab967df := __obf_338eca60ccc15d82.WhatIsNext()
	switch __obf_b2c92a9f9ab967df {
	case jsoniter.NumberValue:
		var __obf_2ee66b45a383db65 json.Number
		__obf_338eca60ccc15d82.
			ReadVal(&__obf_2ee66b45a383db65)
		*((*string)(__obf_e88eec03e62b77ec)) = string(__obf_2ee66b45a383db65)
	case jsoniter.StringValue:
		*((*string)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadString()
	case jsoniter.NilValue:
		__obf_338eca60ccc15d82.
			Skip()
		*((*string)(__obf_e88eec03e62b77ec)) = ""
	default:
		__obf_338eca60ccc15d82.
			ReportError("fuzzyStringDecoder", "not number or string")
	}
}

type __obf_e0ca1fe42424825b struct {
	__obf_290daadbaa52b3f1 func(__obf_f424b9c08d3cfad6 bool, __obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator)
}

func (__obf_bdb89ac21da4a38c *__obf_e0ca1fe42424825b) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	__obf_b2c92a9f9ab967df := __obf_338eca60ccc15d82.WhatIsNext()
	var __obf_c09d7d538e283c2e string
	switch __obf_b2c92a9f9ab967df {
	case jsoniter.NumberValue:
		var __obf_2ee66b45a383db65 json.Number
		__obf_338eca60ccc15d82.
			ReadVal(&__obf_2ee66b45a383db65)
		__obf_c09d7d538e283c2e = string(__obf_2ee66b45a383db65)
	case jsoniter.StringValue:
		__obf_c09d7d538e283c2e = __obf_338eca60ccc15d82.ReadString()
	case jsoniter.BoolValue:
		if __obf_338eca60ccc15d82.ReadBool() {
			__obf_c09d7d538e283c2e = "1"
		} else {
			__obf_c09d7d538e283c2e = "0"
		}
	case jsoniter.NilValue:
		__obf_338eca60ccc15d82.
			Skip()
		__obf_c09d7d538e283c2e = "0"
	default:
		__obf_338eca60ccc15d82.
			ReportError("fuzzyIntegerDecoder", "not number or string")
	}
	if len(__obf_c09d7d538e283c2e) == 0 {
		__obf_c09d7d538e283c2e = "0"
	}
	__obf_d91719946777be7e := __obf_338eca60ccc15d82.Pool().BorrowIterator([]byte(__obf_c09d7d538e283c2e))
	defer __obf_338eca60ccc15d82.Pool().ReturnIterator(__obf_d91719946777be7e)
	__obf_f424b9c08d3cfad6 := strings.IndexByte(__obf_c09d7d538e283c2e, '.') != -1
	__obf_bdb89ac21da4a38c.__obf_290daadbaa52b3f1(__obf_f424b9c08d3cfad6, __obf_e88eec03e62b77ec, __obf_d91719946777be7e)
	if __obf_d91719946777be7e.Error != nil && __obf_d91719946777be7e.Error != io.EOF {
		__obf_338eca60ccc15d82.
			Error = __obf_d91719946777be7e.Error
	}
}

type __obf_51d82e4459cdbe7e struct {
}

func (__obf_bdb89ac21da4a38c *__obf_51d82e4459cdbe7e) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	__obf_b2c92a9f9ab967df := __obf_338eca60ccc15d82.WhatIsNext()
	var __obf_c09d7d538e283c2e string
	switch __obf_b2c92a9f9ab967df {
	case jsoniter.NumberValue:
		*((*float32)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadFloat32()
	case jsoniter.StringValue:
		__obf_c09d7d538e283c2e = __obf_338eca60ccc15d82.ReadString()
		__obf_d91719946777be7e := __obf_338eca60ccc15d82.Pool().BorrowIterator([]byte(__obf_c09d7d538e283c2e))
		defer __obf_338eca60ccc15d82.Pool().ReturnIterator(__obf_d91719946777be7e)
		*((*float32)(__obf_e88eec03e62b77ec)) = __obf_d91719946777be7e.ReadFloat32()
		if __obf_d91719946777be7e.Error != nil && __obf_d91719946777be7e.Error != io.EOF {
			__obf_338eca60ccc15d82.
				Error = __obf_d91719946777be7e.Error
		}
	case jsoniter.BoolValue:
		// support bool to float32
		if __obf_338eca60ccc15d82.ReadBool() {
			*((*float32)(__obf_e88eec03e62b77ec)) = 1
		} else {
			*((*float32)(__obf_e88eec03e62b77ec)) = 0
		}
	case jsoniter.NilValue:
		__obf_338eca60ccc15d82.
			Skip()
		*((*float32)(__obf_e88eec03e62b77ec)) = 0
	default:
		__obf_338eca60ccc15d82.
			ReportError("fuzzyFloat32Decoder", "not number or string")
	}
}

type __obf_e243f1ba7a4e8782 struct {
}

func (__obf_bdb89ac21da4a38c *__obf_e243f1ba7a4e8782) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	__obf_b2c92a9f9ab967df := __obf_338eca60ccc15d82.WhatIsNext()
	var __obf_c09d7d538e283c2e string
	switch __obf_b2c92a9f9ab967df {
	case jsoniter.NumberValue:
		*((*float64)(__obf_e88eec03e62b77ec)) = __obf_338eca60ccc15d82.ReadFloat64()
	case jsoniter.StringValue:
		__obf_c09d7d538e283c2e = __obf_338eca60ccc15d82.ReadString()
		__obf_d91719946777be7e := __obf_338eca60ccc15d82.Pool().BorrowIterator([]byte(__obf_c09d7d538e283c2e))
		defer __obf_338eca60ccc15d82.Pool().ReturnIterator(__obf_d91719946777be7e)
		*((*float64)(__obf_e88eec03e62b77ec)) = __obf_d91719946777be7e.ReadFloat64()
		if __obf_d91719946777be7e.Error != nil && __obf_d91719946777be7e.Error != io.EOF {
			__obf_338eca60ccc15d82.
				Error = __obf_d91719946777be7e.Error
		}
	case jsoniter.BoolValue:
		// support bool to float64
		if __obf_338eca60ccc15d82.ReadBool() {
			*((*float64)(__obf_e88eec03e62b77ec)) = 1
		} else {
			*((*float64)(__obf_e88eec03e62b77ec)) = 0
		}
	case jsoniter.NilValue:
		__obf_338eca60ccc15d82.
			Skip()
		*((*float64)(__obf_e88eec03e62b77ec)) = 0
	default:
		__obf_338eca60ccc15d82.
			ReportError("fuzzyFloat64Decoder", "not number or string")
	}
}
