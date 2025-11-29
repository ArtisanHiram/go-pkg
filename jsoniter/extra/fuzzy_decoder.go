package __obf_060749efdc6ad522

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

const __obf_1fd731fa91da407c = ^uint(0)
const __obf_44ed70768bd63ac5 = int(__obf_1fd731fa91da407c >> 1)
const __obf_4831ef4fd7564572 = -__obf_44ed70768bd63ac5 - 1

// RegisterFuzzyDecoders decode input from PHP with tolerance.
// It will handle string/number auto conversation, and treat empty [] as empty struct.
func RegisterFuzzyDecoders() {
	jsoniter.RegisterExtension(&__obf_8ab576d45823230f{})
	jsoniter.RegisterTypeDecoder("string", &__obf_8cc0c69a5829ac7a{})
	jsoniter.RegisterTypeDecoder("float32", &__obf_cc332a89c3b9b51c{})
	jsoniter.RegisterTypeDecoder("float64", &__obf_ed6cd047d4dee3a9{})
	jsoniter.RegisterTypeDecoder("int", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(__obf_44ed70768bd63ac5) || __obf_d1f9a46fb5537823 < float64(__obf_4831ef4fd7564572) {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode int", "exceed range")
				return
			}
			*((*int)(__obf_deff95ef54922509)) = int(__obf_d1f9a46fb5537823)
		} else {
			*((*int)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadInt()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(__obf_1fd731fa91da407c) || __obf_d1f9a46fb5537823 < 0 {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode uint", "exceed range")
				return
			}
			*((*uint)(__obf_deff95ef54922509)) = uint(__obf_d1f9a46fb5537823)
		} else {
			*((*uint)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadUint()
		}
	}})
	jsoniter.RegisterTypeDecoder("int8", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxInt8) || __obf_d1f9a46fb5537823 < float64(math.MinInt8) {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode int8", "exceed range")
				return
			}
			*((*int8)(__obf_deff95ef54922509)) = int8(__obf_d1f9a46fb5537823)
		} else {
			*((*int8)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadInt8()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint8", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxUint8) || __obf_d1f9a46fb5537823 < 0 {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode uint8", "exceed range")
				return
			}
			*((*uint8)(__obf_deff95ef54922509)) = uint8(__obf_d1f9a46fb5537823)
		} else {
			*((*uint8)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadUint8()
		}
	}})
	jsoniter.RegisterTypeDecoder("int16", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxInt16) || __obf_d1f9a46fb5537823 < float64(math.MinInt16) {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode int16", "exceed range")
				return
			}
			*((*int16)(__obf_deff95ef54922509)) = int16(__obf_d1f9a46fb5537823)
		} else {
			*((*int16)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadInt16()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint16", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxUint16) || __obf_d1f9a46fb5537823 < 0 {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode uint16", "exceed range")
				return
			}
			*((*uint16)(__obf_deff95ef54922509)) = uint16(__obf_d1f9a46fb5537823)
		} else {
			*((*uint16)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadUint16()
		}
	}})
	jsoniter.RegisterTypeDecoder("int32", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxInt32) || __obf_d1f9a46fb5537823 < float64(math.MinInt32) {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode int32", "exceed range")
				return
			}
			*((*int32)(__obf_deff95ef54922509)) = int32(__obf_d1f9a46fb5537823)
		} else {
			*((*int32)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadInt32()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint32", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxUint32) || __obf_d1f9a46fb5537823 < 0 {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode uint32", "exceed range")
				return
			}
			*((*uint32)(__obf_deff95ef54922509)) = uint32(__obf_d1f9a46fb5537823)
		} else {
			*((*uint32)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadUint32()
		}
	}})
	jsoniter.RegisterTypeDecoder("int64", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxInt64) || __obf_d1f9a46fb5537823 < float64(math.MinInt64) {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode int64", "exceed range")
				return
			}
			*((*int64)(__obf_deff95ef54922509)) = int64(__obf_d1f9a46fb5537823)
		} else {
			*((*int64)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadInt64()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint64", &__obf_fb552808d4aa9a09{func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
		if __obf_07b90faf10e7dc2e {
			__obf_d1f9a46fb5537823 := __obf_7eafecc31dc47100.ReadFloat64()
			if __obf_d1f9a46fb5537823 > float64(math.MaxUint64) || __obf_d1f9a46fb5537823 < 0 {
				__obf_7eafecc31dc47100.
					ReportError("fuzzy decode uint64", "exceed range")
				return
			}
			*((*uint64)(__obf_deff95ef54922509)) = uint64(__obf_d1f9a46fb5537823)
		} else {
			*((*uint64)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadUint64()
		}
	}})
}

type __obf_8ab576d45823230f struct {
	jsoniter.DummyExtension
}

func (__obf_1ffe89e9c54d7879 *__obf_8ab576d45823230f) DecorateDecoder(__obf_036de3429d9233a8 reflect2.Type, __obf_991658a42ec507d8 jsoniter.ValDecoder) jsoniter.ValDecoder {
	if __obf_036de3429d9233a8.Kind() == reflect.Struct || __obf_036de3429d9233a8.Kind() == reflect.Map {
		return &__obf_ab5e098e47ad0cf6{__obf_991658a42ec507d8}
	}
	return __obf_991658a42ec507d8
}

type __obf_ab5e098e47ad0cf6 struct {
	__obf_76c025ad7c5266a8 jsoniter.ValDecoder
}

func (__obf_991658a42ec507d8 *__obf_ab5e098e47ad0cf6) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	if __obf_7eafecc31dc47100.WhatIsNext() == jsoniter.ArrayValue {
		__obf_7eafecc31dc47100.
			Skip()
		__obf_09f108a2da4976ae := __obf_7eafecc31dc47100.Pool().BorrowIterator([]byte("{}"))
		defer __obf_7eafecc31dc47100.Pool().ReturnIterator(__obf_09f108a2da4976ae)
		__obf_991658a42ec507d8.__obf_76c025ad7c5266a8.
			Decode(__obf_deff95ef54922509, __obf_09f108a2da4976ae)
	} else {
		__obf_991658a42ec507d8.__obf_76c025ad7c5266a8.
			Decode(__obf_deff95ef54922509, __obf_7eafecc31dc47100)
	}
}

type __obf_8cc0c69a5829ac7a struct {
}

func (__obf_991658a42ec507d8 *__obf_8cc0c69a5829ac7a) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	__obf_18008d63e6242204 := __obf_7eafecc31dc47100.WhatIsNext()
	switch __obf_18008d63e6242204 {
	case jsoniter.NumberValue:
		var __obf_134961554b8eb5d8 json.Number
		__obf_7eafecc31dc47100.
			ReadVal(&__obf_134961554b8eb5d8)
		*((*string)(__obf_deff95ef54922509)) = string(__obf_134961554b8eb5d8)
	case jsoniter.StringValue:
		*((*string)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadString()
	case jsoniter.NilValue:
		__obf_7eafecc31dc47100.
			Skip()
		*((*string)(__obf_deff95ef54922509)) = ""
	default:
		__obf_7eafecc31dc47100.
			ReportError("fuzzyStringDecoder", "not number or string")
	}
}

type __obf_fb552808d4aa9a09 struct {
	__obf_e0191c851af1c1d7 func(__obf_07b90faf10e7dc2e bool, __obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator)
}

func (__obf_991658a42ec507d8 *__obf_fb552808d4aa9a09) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	__obf_18008d63e6242204 := __obf_7eafecc31dc47100.WhatIsNext()
	var __obf_4358ba91b616fb23 string
	switch __obf_18008d63e6242204 {
	case jsoniter.NumberValue:
		var __obf_134961554b8eb5d8 json.Number
		__obf_7eafecc31dc47100.
			ReadVal(&__obf_134961554b8eb5d8)
		__obf_4358ba91b616fb23 = string(__obf_134961554b8eb5d8)
	case jsoniter.StringValue:
		__obf_4358ba91b616fb23 = __obf_7eafecc31dc47100.ReadString()
	case jsoniter.BoolValue:
		if __obf_7eafecc31dc47100.ReadBool() {
			__obf_4358ba91b616fb23 = "1"
		} else {
			__obf_4358ba91b616fb23 = "0"
		}
	case jsoniter.NilValue:
		__obf_7eafecc31dc47100.
			Skip()
		__obf_4358ba91b616fb23 = "0"
	default:
		__obf_7eafecc31dc47100.
			ReportError("fuzzyIntegerDecoder", "not number or string")
	}
	if len(__obf_4358ba91b616fb23) == 0 {
		__obf_4358ba91b616fb23 = "0"
	}
	__obf_09f108a2da4976ae := __obf_7eafecc31dc47100.Pool().BorrowIterator([]byte(__obf_4358ba91b616fb23))
	defer __obf_7eafecc31dc47100.Pool().ReturnIterator(__obf_09f108a2da4976ae)
	__obf_07b90faf10e7dc2e := strings.IndexByte(__obf_4358ba91b616fb23, '.') != -1
	__obf_991658a42ec507d8.__obf_e0191c851af1c1d7(__obf_07b90faf10e7dc2e, __obf_deff95ef54922509, __obf_09f108a2da4976ae)
	if __obf_09f108a2da4976ae.Error != nil && __obf_09f108a2da4976ae.Error != io.EOF {
		__obf_7eafecc31dc47100.
			Error = __obf_09f108a2da4976ae.Error
	}
}

type __obf_cc332a89c3b9b51c struct {
}

func (__obf_991658a42ec507d8 *__obf_cc332a89c3b9b51c) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	__obf_18008d63e6242204 := __obf_7eafecc31dc47100.WhatIsNext()
	var __obf_4358ba91b616fb23 string
	switch __obf_18008d63e6242204 {
	case jsoniter.NumberValue:
		*((*float32)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadFloat32()
	case jsoniter.StringValue:
		__obf_4358ba91b616fb23 = __obf_7eafecc31dc47100.ReadString()
		__obf_09f108a2da4976ae := __obf_7eafecc31dc47100.Pool().BorrowIterator([]byte(__obf_4358ba91b616fb23))
		defer __obf_7eafecc31dc47100.Pool().ReturnIterator(__obf_09f108a2da4976ae)
		*((*float32)(__obf_deff95ef54922509)) = __obf_09f108a2da4976ae.ReadFloat32()
		if __obf_09f108a2da4976ae.Error != nil && __obf_09f108a2da4976ae.Error != io.EOF {
			__obf_7eafecc31dc47100.
				Error = __obf_09f108a2da4976ae.Error
		}
	case jsoniter.BoolValue:
		// support bool to float32
		if __obf_7eafecc31dc47100.ReadBool() {
			*((*float32)(__obf_deff95ef54922509)) = 1
		} else {
			*((*float32)(__obf_deff95ef54922509)) = 0
		}
	case jsoniter.NilValue:
		__obf_7eafecc31dc47100.
			Skip()
		*((*float32)(__obf_deff95ef54922509)) = 0
	default:
		__obf_7eafecc31dc47100.
			ReportError("fuzzyFloat32Decoder", "not number or string")
	}
}

type __obf_ed6cd047d4dee3a9 struct {
}

func (__obf_991658a42ec507d8 *__obf_ed6cd047d4dee3a9) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	__obf_18008d63e6242204 := __obf_7eafecc31dc47100.WhatIsNext()
	var __obf_4358ba91b616fb23 string
	switch __obf_18008d63e6242204 {
	case jsoniter.NumberValue:
		*((*float64)(__obf_deff95ef54922509)) = __obf_7eafecc31dc47100.ReadFloat64()
	case jsoniter.StringValue:
		__obf_4358ba91b616fb23 = __obf_7eafecc31dc47100.ReadString()
		__obf_09f108a2da4976ae := __obf_7eafecc31dc47100.Pool().BorrowIterator([]byte(__obf_4358ba91b616fb23))
		defer __obf_7eafecc31dc47100.Pool().ReturnIterator(__obf_09f108a2da4976ae)
		*((*float64)(__obf_deff95ef54922509)) = __obf_09f108a2da4976ae.ReadFloat64()
		if __obf_09f108a2da4976ae.Error != nil && __obf_09f108a2da4976ae.Error != io.EOF {
			__obf_7eafecc31dc47100.
				Error = __obf_09f108a2da4976ae.Error
		}
	case jsoniter.BoolValue:
		// support bool to float64
		if __obf_7eafecc31dc47100.ReadBool() {
			*((*float64)(__obf_deff95ef54922509)) = 1
		} else {
			*((*float64)(__obf_deff95ef54922509)) = 0
		}
	case jsoniter.NilValue:
		__obf_7eafecc31dc47100.
			Skip()
		*((*float64)(__obf_deff95ef54922509)) = 0
	default:
		__obf_7eafecc31dc47100.
			ReportError("fuzzyFloat64Decoder", "not number or string")
	}
}
