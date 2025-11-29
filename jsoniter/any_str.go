package __obf_5b802ce8d9ba56d6

import (
	"fmt"
	"strconv"
)

type __obf_6ff1a29b3162ef37 struct {
	__obf_fb75d4e4562132ae
	__obf_5406b11e33b9d1d3 string
}

func (any *__obf_6ff1a29b3162ef37) Get(__obf_c5d71353f4393b3c ...any) Any {
	if len(__obf_c5d71353f4393b3c) == 0 {
		return any
	}
	return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, fmt.Errorf("GetIndex %v from simple value", __obf_c5d71353f4393b3c)}
}

func (any *__obf_6ff1a29b3162ef37) Parse() *Iterator {
	return nil
}

func (any *__obf_6ff1a29b3162ef37) ValueType() ValueType {
	return StringValue
}

func (any *__obf_6ff1a29b3162ef37) MustBeValid() Any {
	return any
}

func (any *__obf_6ff1a29b3162ef37) LastError() error {
	return nil
}

func (any *__obf_6ff1a29b3162ef37) ToBool() bool {
	__obf_12c21b79fa86dcba := any.ToString()
	if __obf_12c21b79fa86dcba == "0" {
		return false
	}
	for _, __obf_dab9baaadfa7c8c2 := range __obf_12c21b79fa86dcba {
		switch __obf_dab9baaadfa7c8c2 {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (any *__obf_6ff1a29b3162ef37) ToInt() int {
	return int(any.ToInt64())

}

func (any *__obf_6ff1a29b3162ef37) ToInt32() int32 {
	return int32(any.ToInt64())
}

func (any *__obf_6ff1a29b3162ef37) ToInt64() int64 {
	if any.__obf_5406b11e33b9d1d3 == "" {
		return 0
	}
	__obf_f80c9cd13d76f191 := 1
	__obf_579936615d21cf2c := 0
	if any.__obf_5406b11e33b9d1d3[0] == '+' || any.__obf_5406b11e33b9d1d3[0] == '-' {
		__obf_579936615d21cf2c = 1
	}

	if any.__obf_5406b11e33b9d1d3[0] == '-' {
		__obf_f80c9cd13d76f191 = -1
	}
	__obf_b8fd060b0379b48e := __obf_579936615d21cf2c
	for __obf_2deec7c38ffb6ae3 := __obf_579936615d21cf2c; __obf_2deec7c38ffb6ae3 < len(any.__obf_5406b11e33b9d1d3); __obf_2deec7c38ffb6ae3++ {
		if any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] >= '0' && any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] <= '9' {
			__obf_b8fd060b0379b48e = __obf_2deec7c38ffb6ae3 + 1
		} else {
			break
		}
	}
	__obf_a7a6424513706965, _ := strconv.ParseInt(any.__obf_5406b11e33b9d1d3[__obf_579936615d21cf2c:__obf_b8fd060b0379b48e], 10, 64)
	return int64(__obf_f80c9cd13d76f191) * __obf_a7a6424513706965
}

func (any *__obf_6ff1a29b3162ef37) ToUint() uint {
	return uint(any.ToUint64())
}

func (any *__obf_6ff1a29b3162ef37) ToUint32() uint32 {
	return uint32(any.ToUint64())
}

func (any *__obf_6ff1a29b3162ef37) ToUint64() uint64 {
	if any.__obf_5406b11e33b9d1d3 == "" {
		return 0
	}
	__obf_579936615d21cf2c := 0

	if any.__obf_5406b11e33b9d1d3[0] == '-' {
		return 0
	}
	if any.__obf_5406b11e33b9d1d3[0] == '+' {
		__obf_579936615d21cf2c = 1
	}
	__obf_b8fd060b0379b48e := __obf_579936615d21cf2c
	for __obf_2deec7c38ffb6ae3 := __obf_579936615d21cf2c; __obf_2deec7c38ffb6ae3 < len(any.__obf_5406b11e33b9d1d3); __obf_2deec7c38ffb6ae3++ {
		if any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] >= '0' && any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] <= '9' {
			__obf_b8fd060b0379b48e = __obf_2deec7c38ffb6ae3 + 1
		} else {
			break
		}
	}
	__obf_a7a6424513706965, _ := strconv.ParseUint(any.__obf_5406b11e33b9d1d3[__obf_579936615d21cf2c:__obf_b8fd060b0379b48e], 10, 64)
	return __obf_a7a6424513706965
}

func (any *__obf_6ff1a29b3162ef37) ToFloat32() float32 {
	return float32(any.ToFloat64())
}

func (any *__obf_6ff1a29b3162ef37) ToFloat64() float64 {
	if len(any.__obf_5406b11e33b9d1d3) == 0 {
		return 0
	}

	// first char invalid
	if any.__obf_5406b11e33b9d1d3[0] != '+' && any.__obf_5406b11e33b9d1d3[0] != '-' && (any.__obf_5406b11e33b9d1d3[0] > '9' || any.__obf_5406b11e33b9d1d3[0] < '0') {
		return 0
	}
	__obf_b8fd060b0379b48e := // extract valid num expression from string
		// eg 123true => 123, -12.12xxa => -12.12
		1
	for __obf_2deec7c38ffb6ae3 := 1; __obf_2deec7c38ffb6ae3 < len(any.__obf_5406b11e33b9d1d3); __obf_2deec7c38ffb6ae3++ {
		if any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] == '.' || any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] == 'e' || any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] == 'E' || any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] == '+' || any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] == '-' {
			__obf_b8fd060b0379b48e = __obf_2deec7c38ffb6ae3 + 1
			continue
		}

		// end position is the first char which is not digit
		if any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] >= '0' && any.__obf_5406b11e33b9d1d3[__obf_2deec7c38ffb6ae3] <= '9' {
			__obf_b8fd060b0379b48e = __obf_2deec7c38ffb6ae3 + 1
		} else {
			__obf_b8fd060b0379b48e = __obf_2deec7c38ffb6ae3
			break
		}
	}
	__obf_a7a6424513706965, _ := strconv.ParseFloat(any.__obf_5406b11e33b9d1d3[:__obf_b8fd060b0379b48e], 64)
	return __obf_a7a6424513706965
}

func (any *__obf_6ff1a29b3162ef37) ToString() string {
	return any.__obf_5406b11e33b9d1d3
}

func (any *__obf_6ff1a29b3162ef37) WriteTo(__obf_00fc491caa967a74 *Stream) {
	__obf_00fc491caa967a74.
		WriteString(any.__obf_5406b11e33b9d1d3)
}

func (any *__obf_6ff1a29b3162ef37) GetInterface() any {
	return any.__obf_5406b11e33b9d1d3
}
