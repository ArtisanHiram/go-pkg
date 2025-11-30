package __obf_c3cd04a15c56f32f

import (
	"fmt"
	"strconv"
)

type __obf_7402622ee5d0a489 struct {
	__obf_86279f81acb381f3
	__obf_d59813f23ad740a8 string
}

func (any *__obf_7402622ee5d0a489) Get(__obf_b90e4ca332607787 ...any) Any {
	if len(__obf_b90e4ca332607787) == 0 {
		return any
	}
	return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, fmt.Errorf("GetIndex %v from simple value", __obf_b90e4ca332607787)}
}

func (any *__obf_7402622ee5d0a489) Parse() *Iterator {
	return nil
}

func (any *__obf_7402622ee5d0a489) ValueType() ValueType {
	return StringValue
}

func (any *__obf_7402622ee5d0a489) MustBeValid() Any {
	return any
}

func (any *__obf_7402622ee5d0a489) LastError() error {
	return nil
}

func (any *__obf_7402622ee5d0a489) ToBool() bool {
	__obf_2d944450d48e5810 := any.ToString()
	if __obf_2d944450d48e5810 == "0" {
		return false
	}
	for _, __obf_0c1bc1e511a43120 := range __obf_2d944450d48e5810 {
		switch __obf_0c1bc1e511a43120 {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (any *__obf_7402622ee5d0a489) ToInt() int {
	return int(any.ToInt64())

}

func (any *__obf_7402622ee5d0a489) ToInt32() int32 {
	return int32(any.ToInt64())
}

func (any *__obf_7402622ee5d0a489) ToInt64() int64 {
	if any.__obf_d59813f23ad740a8 == "" {
		return 0
	}
	__obf_8de0383e0b7d0b66 := 1
	__obf_c05c199f4a6e8f13 := 0
	if any.__obf_d59813f23ad740a8[0] == '+' || any.__obf_d59813f23ad740a8[0] == '-' {
		__obf_c05c199f4a6e8f13 = 1
	}

	if any.__obf_d59813f23ad740a8[0] == '-' {
		__obf_8de0383e0b7d0b66 = -1
	}
	__obf_e879c9e02a73ed1c := __obf_c05c199f4a6e8f13
	for __obf_28d099df85f083a8 := __obf_c05c199f4a6e8f13; __obf_28d099df85f083a8 < len(any.__obf_d59813f23ad740a8); __obf_28d099df85f083a8++ {
		if any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] >= '0' && any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] <= '9' {
			__obf_e879c9e02a73ed1c = __obf_28d099df85f083a8 + 1
		} else {
			break
		}
	}
	__obf_c6664354ca9dfd29, _ := strconv.ParseInt(any.__obf_d59813f23ad740a8[__obf_c05c199f4a6e8f13:__obf_e879c9e02a73ed1c], 10, 64)
	return int64(__obf_8de0383e0b7d0b66) * __obf_c6664354ca9dfd29
}

func (any *__obf_7402622ee5d0a489) ToUint() uint {
	return uint(any.ToUint64())
}

func (any *__obf_7402622ee5d0a489) ToUint32() uint32 {
	return uint32(any.ToUint64())
}

func (any *__obf_7402622ee5d0a489) ToUint64() uint64 {
	if any.__obf_d59813f23ad740a8 == "" {
		return 0
	}
	__obf_c05c199f4a6e8f13 := 0

	if any.__obf_d59813f23ad740a8[0] == '-' {
		return 0
	}
	if any.__obf_d59813f23ad740a8[0] == '+' {
		__obf_c05c199f4a6e8f13 = 1
	}
	__obf_e879c9e02a73ed1c := __obf_c05c199f4a6e8f13
	for __obf_28d099df85f083a8 := __obf_c05c199f4a6e8f13; __obf_28d099df85f083a8 < len(any.__obf_d59813f23ad740a8); __obf_28d099df85f083a8++ {
		if any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] >= '0' && any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] <= '9' {
			__obf_e879c9e02a73ed1c = __obf_28d099df85f083a8 + 1
		} else {
			break
		}
	}
	__obf_c6664354ca9dfd29, _ := strconv.ParseUint(any.__obf_d59813f23ad740a8[__obf_c05c199f4a6e8f13:__obf_e879c9e02a73ed1c], 10, 64)
	return __obf_c6664354ca9dfd29
}

func (any *__obf_7402622ee5d0a489) ToFloat32() float32 {
	return float32(any.ToFloat64())
}

func (any *__obf_7402622ee5d0a489) ToFloat64() float64 {
	if len(any.__obf_d59813f23ad740a8) == 0 {
		return 0
	}

	// first char invalid
	if any.__obf_d59813f23ad740a8[0] != '+' && any.__obf_d59813f23ad740a8[0] != '-' && (any.__obf_d59813f23ad740a8[0] > '9' || any.__obf_d59813f23ad740a8[0] < '0') {
		return 0
	}
	__obf_e879c9e02a73ed1c := // extract valid num expression from string
		// eg 123true => 123, -12.12xxa => -12.12
		1
	for __obf_28d099df85f083a8 := 1; __obf_28d099df85f083a8 < len(any.__obf_d59813f23ad740a8); __obf_28d099df85f083a8++ {
		if any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] == '.' || any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] == 'e' || any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] == 'E' || any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] == '+' || any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] == '-' {
			__obf_e879c9e02a73ed1c = __obf_28d099df85f083a8 + 1
			continue
		}

		// end position is the first char which is not digit
		if any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] >= '0' && any.__obf_d59813f23ad740a8[__obf_28d099df85f083a8] <= '9' {
			__obf_e879c9e02a73ed1c = __obf_28d099df85f083a8 + 1
		} else {
			__obf_e879c9e02a73ed1c = __obf_28d099df85f083a8
			break
		}
	}
	__obf_c6664354ca9dfd29, _ := strconv.ParseFloat(any.__obf_d59813f23ad740a8[:__obf_e879c9e02a73ed1c], 64)
	return __obf_c6664354ca9dfd29
}

func (any *__obf_7402622ee5d0a489) ToString() string {
	return any.__obf_d59813f23ad740a8
}

func (any *__obf_7402622ee5d0a489) WriteTo(__obf_2361f5214e84c60f *Stream) {
	__obf_2361f5214e84c60f.
		WriteString(any.__obf_d59813f23ad740a8)
}

func (any *__obf_7402622ee5d0a489) GetInterface() any {
	return any.__obf_d59813f23ad740a8
}
