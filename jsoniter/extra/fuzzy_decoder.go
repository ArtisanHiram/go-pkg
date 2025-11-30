package __obf_38f1d2f091ad74e0

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

const __obf_b43de09e9f6c9205 = ^uint(0)
const __obf_bae59f7754308a62 = int(__obf_b43de09e9f6c9205 >> 1)
const __obf_085cd86a52e802bc = -__obf_bae59f7754308a62 - 1

// RegisterFuzzyDecoders decode input from PHP with tolerance.
// It will handle string/number auto conversation, and treat empty [] as empty struct.
func RegisterFuzzyDecoders() {
	jsoniter.RegisterExtension(&__obf_794fec31d6971cbb{})
	jsoniter.RegisterTypeDecoder("string", &__obf_f1ee8d425ca14d84{})
	jsoniter.RegisterTypeDecoder("float32", &__obf_5879b4018281eb7e{})
	jsoniter.RegisterTypeDecoder("float64", &__obf_e59e28e083d85e5d{})
	jsoniter.RegisterTypeDecoder("int", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(__obf_bae59f7754308a62) || __obf_7d46d0382746f77e < float64(__obf_085cd86a52e802bc) {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode int", "exceed range")
				return
			}
			*((*int)(__obf_35567cf7daf6e12d)) = int(__obf_7d46d0382746f77e)
		} else {
			*((*int)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadInt()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(__obf_b43de09e9f6c9205) || __obf_7d46d0382746f77e < 0 {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode uint", "exceed range")
				return
			}
			*((*uint)(__obf_35567cf7daf6e12d)) = uint(__obf_7d46d0382746f77e)
		} else {
			*((*uint)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadUint()
		}
	}})
	jsoniter.RegisterTypeDecoder("int8", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxInt8) || __obf_7d46d0382746f77e < float64(math.MinInt8) {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode int8", "exceed range")
				return
			}
			*((*int8)(__obf_35567cf7daf6e12d)) = int8(__obf_7d46d0382746f77e)
		} else {
			*((*int8)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadInt8()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint8", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxUint8) || __obf_7d46d0382746f77e < 0 {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode uint8", "exceed range")
				return
			}
			*((*uint8)(__obf_35567cf7daf6e12d)) = uint8(__obf_7d46d0382746f77e)
		} else {
			*((*uint8)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadUint8()
		}
	}})
	jsoniter.RegisterTypeDecoder("int16", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxInt16) || __obf_7d46d0382746f77e < float64(math.MinInt16) {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode int16", "exceed range")
				return
			}
			*((*int16)(__obf_35567cf7daf6e12d)) = int16(__obf_7d46d0382746f77e)
		} else {
			*((*int16)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadInt16()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint16", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxUint16) || __obf_7d46d0382746f77e < 0 {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode uint16", "exceed range")
				return
			}
			*((*uint16)(__obf_35567cf7daf6e12d)) = uint16(__obf_7d46d0382746f77e)
		} else {
			*((*uint16)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadUint16()
		}
	}})
	jsoniter.RegisterTypeDecoder("int32", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxInt32) || __obf_7d46d0382746f77e < float64(math.MinInt32) {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode int32", "exceed range")
				return
			}
			*((*int32)(__obf_35567cf7daf6e12d)) = int32(__obf_7d46d0382746f77e)
		} else {
			*((*int32)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadInt32()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint32", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxUint32) || __obf_7d46d0382746f77e < 0 {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode uint32", "exceed range")
				return
			}
			*((*uint32)(__obf_35567cf7daf6e12d)) = uint32(__obf_7d46d0382746f77e)
		} else {
			*((*uint32)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadUint32()
		}
	}})
	jsoniter.RegisterTypeDecoder("int64", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxInt64) || __obf_7d46d0382746f77e < float64(math.MinInt64) {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode int64", "exceed range")
				return
			}
			*((*int64)(__obf_35567cf7daf6e12d)) = int64(__obf_7d46d0382746f77e)
		} else {
			*((*int64)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadInt64()
		}
	}})
	jsoniter.RegisterTypeDecoder("uint64", &__obf_5a51e45bca4e6e91{func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
		if __obf_db694893efb14880 {
			__obf_7d46d0382746f77e := __obf_113f80f39cc94185.ReadFloat64()
			if __obf_7d46d0382746f77e > float64(math.MaxUint64) || __obf_7d46d0382746f77e < 0 {
				__obf_113f80f39cc94185.
					ReportError("fuzzy decode uint64", "exceed range")
				return
			}
			*((*uint64)(__obf_35567cf7daf6e12d)) = uint64(__obf_7d46d0382746f77e)
		} else {
			*((*uint64)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadUint64()
		}
	}})
}

type __obf_794fec31d6971cbb struct {
	jsoniter.DummyExtension
}

func (__obf_9c280f4c530aab7f *__obf_794fec31d6971cbb) DecorateDecoder(__obf_3c1b235923760ab5 reflect2.Type, __obf_98f9bd3c7c5539a3 jsoniter.ValDecoder) jsoniter.ValDecoder {
	if __obf_3c1b235923760ab5.Kind() == reflect.Struct || __obf_3c1b235923760ab5.Kind() == reflect.Map {
		return &__obf_57505ed933cefa6a{__obf_98f9bd3c7c5539a3}
	}
	return __obf_98f9bd3c7c5539a3
}

type __obf_57505ed933cefa6a struct {
	__obf_5690015e7878691a jsoniter.ValDecoder
}

func (__obf_98f9bd3c7c5539a3 *__obf_57505ed933cefa6a) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	if __obf_113f80f39cc94185.WhatIsNext() == jsoniter.ArrayValue {
		__obf_113f80f39cc94185.
			Skip()
		__obf_33a5712a056829be := __obf_113f80f39cc94185.Pool().BorrowIterator([]byte("{}"))
		defer __obf_113f80f39cc94185.Pool().ReturnIterator(__obf_33a5712a056829be)
		__obf_98f9bd3c7c5539a3.__obf_5690015e7878691a.
			Decode(__obf_35567cf7daf6e12d, __obf_33a5712a056829be)
	} else {
		__obf_98f9bd3c7c5539a3.__obf_5690015e7878691a.
			Decode(__obf_35567cf7daf6e12d, __obf_113f80f39cc94185)
	}
}

type __obf_f1ee8d425ca14d84 struct {
}

func (__obf_98f9bd3c7c5539a3 *__obf_f1ee8d425ca14d84) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	__obf_09245e59dcfc3fa7 := __obf_113f80f39cc94185.WhatIsNext()
	switch __obf_09245e59dcfc3fa7 {
	case jsoniter.NumberValue:
		var __obf_c175143b5d0beab8 json.Number
		__obf_113f80f39cc94185.
			ReadVal(&__obf_c175143b5d0beab8)
		*((*string)(__obf_35567cf7daf6e12d)) = string(__obf_c175143b5d0beab8)
	case jsoniter.StringValue:
		*((*string)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadString()
	case jsoniter.NilValue:
		__obf_113f80f39cc94185.
			Skip()
		*((*string)(__obf_35567cf7daf6e12d)) = ""
	default:
		__obf_113f80f39cc94185.
			ReportError("fuzzyStringDecoder", "not number or string")
	}
}

type __obf_5a51e45bca4e6e91 struct {
	__obf_e67be8c7e81c0c6f func(__obf_db694893efb14880 bool, __obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator)
}

func (__obf_98f9bd3c7c5539a3 *__obf_5a51e45bca4e6e91) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	__obf_09245e59dcfc3fa7 := __obf_113f80f39cc94185.WhatIsNext()
	var __obf_69bdd887beaf0c09 string
	switch __obf_09245e59dcfc3fa7 {
	case jsoniter.NumberValue:
		var __obf_c175143b5d0beab8 json.Number
		__obf_113f80f39cc94185.
			ReadVal(&__obf_c175143b5d0beab8)
		__obf_69bdd887beaf0c09 = string(__obf_c175143b5d0beab8)
	case jsoniter.StringValue:
		__obf_69bdd887beaf0c09 = __obf_113f80f39cc94185.ReadString()
	case jsoniter.BoolValue:
		if __obf_113f80f39cc94185.ReadBool() {
			__obf_69bdd887beaf0c09 = "1"
		} else {
			__obf_69bdd887beaf0c09 = "0"
		}
	case jsoniter.NilValue:
		__obf_113f80f39cc94185.
			Skip()
		__obf_69bdd887beaf0c09 = "0"
	default:
		__obf_113f80f39cc94185.
			ReportError("fuzzyIntegerDecoder", "not number or string")
	}
	if len(__obf_69bdd887beaf0c09) == 0 {
		__obf_69bdd887beaf0c09 = "0"
	}
	__obf_33a5712a056829be := __obf_113f80f39cc94185.Pool().BorrowIterator([]byte(__obf_69bdd887beaf0c09))
	defer __obf_113f80f39cc94185.Pool().ReturnIterator(__obf_33a5712a056829be)
	__obf_db694893efb14880 := strings.IndexByte(__obf_69bdd887beaf0c09, '.') != -1
	__obf_98f9bd3c7c5539a3.__obf_e67be8c7e81c0c6f(__obf_db694893efb14880, __obf_35567cf7daf6e12d, __obf_33a5712a056829be)
	if __obf_33a5712a056829be.Error != nil && __obf_33a5712a056829be.Error != io.EOF {
		__obf_113f80f39cc94185.
			Error = __obf_33a5712a056829be.Error
	}
}

type __obf_5879b4018281eb7e struct {
}

func (__obf_98f9bd3c7c5539a3 *__obf_5879b4018281eb7e) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	__obf_09245e59dcfc3fa7 := __obf_113f80f39cc94185.WhatIsNext()
	var __obf_69bdd887beaf0c09 string
	switch __obf_09245e59dcfc3fa7 {
	case jsoniter.NumberValue:
		*((*float32)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadFloat32()
	case jsoniter.StringValue:
		__obf_69bdd887beaf0c09 = __obf_113f80f39cc94185.ReadString()
		__obf_33a5712a056829be := __obf_113f80f39cc94185.Pool().BorrowIterator([]byte(__obf_69bdd887beaf0c09))
		defer __obf_113f80f39cc94185.Pool().ReturnIterator(__obf_33a5712a056829be)
		*((*float32)(__obf_35567cf7daf6e12d)) = __obf_33a5712a056829be.ReadFloat32()
		if __obf_33a5712a056829be.Error != nil && __obf_33a5712a056829be.Error != io.EOF {
			__obf_113f80f39cc94185.
				Error = __obf_33a5712a056829be.Error
		}
	case jsoniter.BoolValue:
		// support bool to float32
		if __obf_113f80f39cc94185.ReadBool() {
			*((*float32)(__obf_35567cf7daf6e12d)) = 1
		} else {
			*((*float32)(__obf_35567cf7daf6e12d)) = 0
		}
	case jsoniter.NilValue:
		__obf_113f80f39cc94185.
			Skip()
		*((*float32)(__obf_35567cf7daf6e12d)) = 0
	default:
		__obf_113f80f39cc94185.
			ReportError("fuzzyFloat32Decoder", "not number or string")
	}
}

type __obf_e59e28e083d85e5d struct {
}

func (__obf_98f9bd3c7c5539a3 *__obf_e59e28e083d85e5d) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	__obf_09245e59dcfc3fa7 := __obf_113f80f39cc94185.WhatIsNext()
	var __obf_69bdd887beaf0c09 string
	switch __obf_09245e59dcfc3fa7 {
	case jsoniter.NumberValue:
		*((*float64)(__obf_35567cf7daf6e12d)) = __obf_113f80f39cc94185.ReadFloat64()
	case jsoniter.StringValue:
		__obf_69bdd887beaf0c09 = __obf_113f80f39cc94185.ReadString()
		__obf_33a5712a056829be := __obf_113f80f39cc94185.Pool().BorrowIterator([]byte(__obf_69bdd887beaf0c09))
		defer __obf_113f80f39cc94185.Pool().ReturnIterator(__obf_33a5712a056829be)
		*((*float64)(__obf_35567cf7daf6e12d)) = __obf_33a5712a056829be.ReadFloat64()
		if __obf_33a5712a056829be.Error != nil && __obf_33a5712a056829be.Error != io.EOF {
			__obf_113f80f39cc94185.
				Error = __obf_33a5712a056829be.Error
		}
	case jsoniter.BoolValue:
		// support bool to float64
		if __obf_113f80f39cc94185.ReadBool() {
			*((*float64)(__obf_35567cf7daf6e12d)) = 1
		} else {
			*((*float64)(__obf_35567cf7daf6e12d)) = 0
		}
	case jsoniter.NilValue:
		__obf_113f80f39cc94185.
			Skip()
		*((*float64)(__obf_35567cf7daf6e12d)) = 0
	default:
		__obf_113f80f39cc94185.
			ReportError("fuzzyFloat64Decoder", "not number or string")
	}
}
