package __obf_91620b895eeff9ed

import (
	"fmt"
	"strconv"
)

type __obf_e76c1c0259373083 struct {
	__obf_58563642f42f4a04
	__obf_bbfd2ac8618a6f0c string
}

func (any *__obf_e76c1c0259373083) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	if len(__obf_82c50fcbfc9ab432) == 0 {
		return any
	}
	return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, fmt.Errorf("GetIndex %v from simple value", __obf_82c50fcbfc9ab432)}
}

func (any *__obf_e76c1c0259373083) Parse() *Iterator {
	return nil
}

func (any *__obf_e76c1c0259373083) ValueType() ValueType {
	return StringValue
}

func (any *__obf_e76c1c0259373083) MustBeValid() Any {
	return any
}

func (any *__obf_e76c1c0259373083) LastError() error {
	return nil
}

func (any *__obf_e76c1c0259373083) ToBool() bool {
	__obf_e91bd2feb751e4f1 := any.ToString()
	if __obf_e91bd2feb751e4f1 == "0" {
		return false
	}
	for _, __obf_f16b4157911bc9af := range __obf_e91bd2feb751e4f1 {
		switch __obf_f16b4157911bc9af {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (any *__obf_e76c1c0259373083) ToInt() int {
	return int(any.ToInt64())

}

func (any *__obf_e76c1c0259373083) ToInt32() int32 {
	return int32(any.ToInt64())
}

func (any *__obf_e76c1c0259373083) ToInt64() int64 {
	if any.__obf_bbfd2ac8618a6f0c == "" {
		return 0
	}
	__obf_e21907a9c2b199a2 := 1
	__obf_dc00dc2962f7736e := 0
	if any.__obf_bbfd2ac8618a6f0c[0] == '+' || any.__obf_bbfd2ac8618a6f0c[0] == '-' {
		__obf_dc00dc2962f7736e = 1
	}

	if any.__obf_bbfd2ac8618a6f0c[0] == '-' {
		__obf_e21907a9c2b199a2 = -1
	}
	__obf_7171ef87a0db195c := __obf_dc00dc2962f7736e
	for __obf_5aa5c8829b97f182 := __obf_dc00dc2962f7736e; __obf_5aa5c8829b97f182 < len(any.__obf_bbfd2ac8618a6f0c); __obf_5aa5c8829b97f182++ {
		if any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] >= '0' && any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] <= '9' {
			__obf_7171ef87a0db195c = __obf_5aa5c8829b97f182 + 1
		} else {
			break
		}
	}
	__obf_3eed30596febc10c, _ := strconv.ParseInt(any.__obf_bbfd2ac8618a6f0c[__obf_dc00dc2962f7736e:__obf_7171ef87a0db195c], 10, 64)
	return int64(__obf_e21907a9c2b199a2) * __obf_3eed30596febc10c
}

func (any *__obf_e76c1c0259373083) ToUint() uint {
	return uint(any.ToUint64())
}

func (any *__obf_e76c1c0259373083) ToUint32() uint32 {
	return uint32(any.ToUint64())
}

func (any *__obf_e76c1c0259373083) ToUint64() uint64 {
	if any.__obf_bbfd2ac8618a6f0c == "" {
		return 0
	}
	__obf_dc00dc2962f7736e := 0

	if any.__obf_bbfd2ac8618a6f0c[0] == '-' {
		return 0
	}
	if any.__obf_bbfd2ac8618a6f0c[0] == '+' {
		__obf_dc00dc2962f7736e = 1
	}
	__obf_7171ef87a0db195c := __obf_dc00dc2962f7736e
	for __obf_5aa5c8829b97f182 := __obf_dc00dc2962f7736e; __obf_5aa5c8829b97f182 < len(any.__obf_bbfd2ac8618a6f0c); __obf_5aa5c8829b97f182++ {
		if any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] >= '0' && any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] <= '9' {
			__obf_7171ef87a0db195c = __obf_5aa5c8829b97f182 + 1
		} else {
			break
		}
	}
	__obf_3eed30596febc10c, _ := strconv.ParseUint(any.__obf_bbfd2ac8618a6f0c[__obf_dc00dc2962f7736e:__obf_7171ef87a0db195c], 10, 64)
	return __obf_3eed30596febc10c
}

func (any *__obf_e76c1c0259373083) ToFloat32() float32 {
	return float32(any.ToFloat64())
}

func (any *__obf_e76c1c0259373083) ToFloat64() float64 {
	if len(any.__obf_bbfd2ac8618a6f0c) == 0 {
		return 0
	}

	// first char invalid
	if any.__obf_bbfd2ac8618a6f0c[0] != '+' && any.__obf_bbfd2ac8618a6f0c[0] != '-' && (any.__obf_bbfd2ac8618a6f0c[0] > '9' || any.__obf_bbfd2ac8618a6f0c[0] < '0') {
		return 0
	}
	__obf_7171ef87a0db195c := // extract valid num expression from string
		// eg 123true => 123, -12.12xxa => -12.12
		1
	for __obf_5aa5c8829b97f182 := 1; __obf_5aa5c8829b97f182 < len(any.__obf_bbfd2ac8618a6f0c); __obf_5aa5c8829b97f182++ {
		if any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] == '.' || any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] == 'e' || any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] == 'E' || any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] == '+' || any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] == '-' {
			__obf_7171ef87a0db195c = __obf_5aa5c8829b97f182 + 1
			continue
		}

		// end position is the first char which is not digit
		if any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] >= '0' && any.__obf_bbfd2ac8618a6f0c[__obf_5aa5c8829b97f182] <= '9' {
			__obf_7171ef87a0db195c = __obf_5aa5c8829b97f182 + 1
		} else {
			__obf_7171ef87a0db195c = __obf_5aa5c8829b97f182
			break
		}
	}
	__obf_3eed30596febc10c, _ := strconv.ParseFloat(any.__obf_bbfd2ac8618a6f0c[:__obf_7171ef87a0db195c], 64)
	return __obf_3eed30596febc10c
}

func (any *__obf_e76c1c0259373083) ToString() string {
	return any.__obf_bbfd2ac8618a6f0c
}

func (any *__obf_e76c1c0259373083) WriteTo(__obf_850a7457bb739a32 *Stream) {
	__obf_850a7457bb739a32.
		WriteString(any.__obf_bbfd2ac8618a6f0c)
}

func (any *__obf_e76c1c0259373083) GetInterface() any {
	return any.__obf_bbfd2ac8618a6f0c
}
