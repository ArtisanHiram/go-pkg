package __obf_703d23ba520c3295

import (
	"fmt"
	"strconv"
)

type __obf_223737ac6afcb5d9 struct {
	__obf_b2a2e31336581ab8
	__obf_b924538fffe5fd64 string
}

func (any *__obf_223737ac6afcb5d9) Get(__obf_267bdf615cb1b310 ...any) Any {
	if len(__obf_267bdf615cb1b310) == 0 {
		return any
	}
	return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, fmt.Errorf("GetIndex %v from simple value", __obf_267bdf615cb1b310)}
}

func (any *__obf_223737ac6afcb5d9) Parse() *Iterator {
	return nil
}

func (any *__obf_223737ac6afcb5d9) ValueType() ValueType {
	return StringValue
}

func (any *__obf_223737ac6afcb5d9) MustBeValid() Any {
	return any
}

func (any *__obf_223737ac6afcb5d9) LastError() error {
	return nil
}

func (any *__obf_223737ac6afcb5d9) ToBool() bool {
	__obf_44b48c26051f8033 := any.ToString()
	if __obf_44b48c26051f8033 == "0" {
		return false
	}
	for _, __obf_bd08f5161fff294a := range __obf_44b48c26051f8033 {
		switch __obf_bd08f5161fff294a {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (any *__obf_223737ac6afcb5d9) ToInt() int {
	return int(any.ToInt64())

}

func (any *__obf_223737ac6afcb5d9) ToInt32() int32 {
	return int32(any.ToInt64())
}

func (any *__obf_223737ac6afcb5d9) ToInt64() int64 {
	if any.__obf_b924538fffe5fd64 == "" {
		return 0
	}
	__obf_ccfdd34190f74108 := 1
	__obf_582069b57dfac5bb := 0
	if any.__obf_b924538fffe5fd64[0] == '+' || any.__obf_b924538fffe5fd64[0] == '-' {
		__obf_582069b57dfac5bb = 1
	}

	if any.__obf_b924538fffe5fd64[0] == '-' {
		__obf_ccfdd34190f74108 = -1
	}
	__obf_a7774b733c556f32 := __obf_582069b57dfac5bb
	for __obf_b0a5d2bd48690f1d := __obf_582069b57dfac5bb; __obf_b0a5d2bd48690f1d < len(any.__obf_b924538fffe5fd64); __obf_b0a5d2bd48690f1d++ {
		if any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] >= '0' && any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] <= '9' {
			__obf_a7774b733c556f32 = __obf_b0a5d2bd48690f1d + 1
		} else {
			break
		}
	}
	__obf_18a085c1869fd128, _ := strconv.ParseInt(any.__obf_b924538fffe5fd64[__obf_582069b57dfac5bb:__obf_a7774b733c556f32], 10, 64)
	return int64(__obf_ccfdd34190f74108) * __obf_18a085c1869fd128
}

func (any *__obf_223737ac6afcb5d9) ToUint() uint {
	return uint(any.ToUint64())
}

func (any *__obf_223737ac6afcb5d9) ToUint32() uint32 {
	return uint32(any.ToUint64())
}

func (any *__obf_223737ac6afcb5d9) ToUint64() uint64 {
	if any.__obf_b924538fffe5fd64 == "" {
		return 0
	}
	__obf_582069b57dfac5bb := 0

	if any.__obf_b924538fffe5fd64[0] == '-' {
		return 0
	}
	if any.__obf_b924538fffe5fd64[0] == '+' {
		__obf_582069b57dfac5bb = 1
	}
	__obf_a7774b733c556f32 := __obf_582069b57dfac5bb
	for __obf_b0a5d2bd48690f1d := __obf_582069b57dfac5bb; __obf_b0a5d2bd48690f1d < len(any.__obf_b924538fffe5fd64); __obf_b0a5d2bd48690f1d++ {
		if any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] >= '0' && any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] <= '9' {
			__obf_a7774b733c556f32 = __obf_b0a5d2bd48690f1d + 1
		} else {
			break
		}
	}
	__obf_18a085c1869fd128, _ := strconv.ParseUint(any.__obf_b924538fffe5fd64[__obf_582069b57dfac5bb:__obf_a7774b733c556f32], 10, 64)
	return __obf_18a085c1869fd128
}

func (any *__obf_223737ac6afcb5d9) ToFloat32() float32 {
	return float32(any.ToFloat64())
}

func (any *__obf_223737ac6afcb5d9) ToFloat64() float64 {
	if len(any.__obf_b924538fffe5fd64) == 0 {
		return 0
	}

	// first char invalid
	if any.__obf_b924538fffe5fd64[0] != '+' && any.__obf_b924538fffe5fd64[0] != '-' && (any.__obf_b924538fffe5fd64[0] > '9' || any.__obf_b924538fffe5fd64[0] < '0') {
		return 0
	}
	__obf_a7774b733c556f32 := // extract valid num expression from string
		// eg 123true => 123, -12.12xxa => -12.12
		1
	for __obf_b0a5d2bd48690f1d := 1; __obf_b0a5d2bd48690f1d < len(any.__obf_b924538fffe5fd64); __obf_b0a5d2bd48690f1d++ {
		if any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] == '.' || any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] == 'e' || any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] == 'E' || any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] == '+' || any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] == '-' {
			__obf_a7774b733c556f32 = __obf_b0a5d2bd48690f1d + 1
			continue
		}

		// end position is the first char which is not digit
		if any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] >= '0' && any.__obf_b924538fffe5fd64[__obf_b0a5d2bd48690f1d] <= '9' {
			__obf_a7774b733c556f32 = __obf_b0a5d2bd48690f1d + 1
		} else {
			__obf_a7774b733c556f32 = __obf_b0a5d2bd48690f1d
			break
		}
	}
	__obf_18a085c1869fd128, _ := strconv.ParseFloat(any.__obf_b924538fffe5fd64[:__obf_a7774b733c556f32], 64)
	return __obf_18a085c1869fd128
}

func (any *__obf_223737ac6afcb5d9) ToString() string {
	return any.__obf_b924538fffe5fd64
}

func (any *__obf_223737ac6afcb5d9) WriteTo(__obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteString(any.__obf_b924538fffe5fd64)
}

func (any *__obf_223737ac6afcb5d9) GetInterface() any {
	return any.__obf_b924538fffe5fd64
}
