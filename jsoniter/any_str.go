package __obf_030d39f7a8de96c6

import (
	"fmt"
	"strconv"
)

type __obf_b2998b74708f8d59 struct {
	__obf_7834a13a259af712
	__obf_efaf2719b44ad8ac string
}

func (any *__obf_b2998b74708f8d59) Get(__obf_f77a9507782b919d ...any) Any {
	if len(__obf_f77a9507782b919d) == 0 {
		return any
	}
	return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, fmt.Errorf("GetIndex %v from simple value", __obf_f77a9507782b919d)}
}

func (any *__obf_b2998b74708f8d59) Parse() *Iterator {
	return nil
}

func (any *__obf_b2998b74708f8d59) ValueType() ValueType {
	return StringValue
}

func (any *__obf_b2998b74708f8d59) MustBeValid() Any {
	return any
}

func (any *__obf_b2998b74708f8d59) LastError() error {
	return nil
}

func (any *__obf_b2998b74708f8d59) ToBool() bool {
	__obf_428c3b4a95725c4a := any.ToString()
	if __obf_428c3b4a95725c4a == "0" {
		return false
	}
	for _, __obf_24309b3b0ff9ba22 := range __obf_428c3b4a95725c4a {
		switch __obf_24309b3b0ff9ba22 {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (any *__obf_b2998b74708f8d59) ToInt() int {
	return int(any.ToInt64())

}

func (any *__obf_b2998b74708f8d59) ToInt32() int32 {
	return int32(any.ToInt64())
}

func (any *__obf_b2998b74708f8d59) ToInt64() int64 {
	if any.__obf_efaf2719b44ad8ac == "" {
		return 0
	}
	__obf_a60939f13c5282c7 := 1
	__obf_adacdc4cf247c29a := 0
	if any.__obf_efaf2719b44ad8ac[0] == '+' || any.__obf_efaf2719b44ad8ac[0] == '-' {
		__obf_adacdc4cf247c29a = 1
	}

	if any.__obf_efaf2719b44ad8ac[0] == '-' {
		__obf_a60939f13c5282c7 = -1
	}
	__obf_192357dc4e126889 := __obf_adacdc4cf247c29a
	for __obf_82c6e05b9d226c58 := __obf_adacdc4cf247c29a; __obf_82c6e05b9d226c58 < len(any.__obf_efaf2719b44ad8ac); __obf_82c6e05b9d226c58++ {
		if any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] >= '0' && any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] <= '9' {
			__obf_192357dc4e126889 = __obf_82c6e05b9d226c58 + 1
		} else {
			break
		}
	}
	__obf_b2c4c7b1b90b2110, _ := strconv.ParseInt(any.__obf_efaf2719b44ad8ac[__obf_adacdc4cf247c29a:__obf_192357dc4e126889], 10, 64)
	return int64(__obf_a60939f13c5282c7) * __obf_b2c4c7b1b90b2110
}

func (any *__obf_b2998b74708f8d59) ToUint() uint {
	return uint(any.ToUint64())
}

func (any *__obf_b2998b74708f8d59) ToUint32() uint32 {
	return uint32(any.ToUint64())
}

func (any *__obf_b2998b74708f8d59) ToUint64() uint64 {
	if any.__obf_efaf2719b44ad8ac == "" {
		return 0
	}
	__obf_adacdc4cf247c29a := 0

	if any.__obf_efaf2719b44ad8ac[0] == '-' {
		return 0
	}
	if any.__obf_efaf2719b44ad8ac[0] == '+' {
		__obf_adacdc4cf247c29a = 1
	}
	__obf_192357dc4e126889 := __obf_adacdc4cf247c29a
	for __obf_82c6e05b9d226c58 := __obf_adacdc4cf247c29a; __obf_82c6e05b9d226c58 < len(any.__obf_efaf2719b44ad8ac); __obf_82c6e05b9d226c58++ {
		if any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] >= '0' && any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] <= '9' {
			__obf_192357dc4e126889 = __obf_82c6e05b9d226c58 + 1
		} else {
			break
		}
	}
	__obf_b2c4c7b1b90b2110, _ := strconv.ParseUint(any.__obf_efaf2719b44ad8ac[__obf_adacdc4cf247c29a:__obf_192357dc4e126889], 10, 64)
	return __obf_b2c4c7b1b90b2110
}

func (any *__obf_b2998b74708f8d59) ToFloat32() float32 {
	return float32(any.ToFloat64())
}

func (any *__obf_b2998b74708f8d59) ToFloat64() float64 {
	if len(any.__obf_efaf2719b44ad8ac) == 0 {
		return 0
	}

	// first char invalid
	if any.__obf_efaf2719b44ad8ac[0] != '+' && any.__obf_efaf2719b44ad8ac[0] != '-' && (any.__obf_efaf2719b44ad8ac[0] > '9' || any.__obf_efaf2719b44ad8ac[0] < '0') {
		return 0
	}
	__obf_192357dc4e126889 := // extract valid num expression from string
		// eg 123true => 123, -12.12xxa => -12.12
		1
	for __obf_82c6e05b9d226c58 := 1; __obf_82c6e05b9d226c58 < len(any.__obf_efaf2719b44ad8ac); __obf_82c6e05b9d226c58++ {
		if any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] == '.' || any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] == 'e' || any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] == 'E' || any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] == '+' || any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] == '-' {
			__obf_192357dc4e126889 = __obf_82c6e05b9d226c58 + 1
			continue
		}

		// end position is the first char which is not digit
		if any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] >= '0' && any.__obf_efaf2719b44ad8ac[__obf_82c6e05b9d226c58] <= '9' {
			__obf_192357dc4e126889 = __obf_82c6e05b9d226c58 + 1
		} else {
			__obf_192357dc4e126889 = __obf_82c6e05b9d226c58
			break
		}
	}
	__obf_b2c4c7b1b90b2110, _ := strconv.ParseFloat(any.__obf_efaf2719b44ad8ac[:__obf_192357dc4e126889], 64)
	return __obf_b2c4c7b1b90b2110
}

func (any *__obf_b2998b74708f8d59) ToString() string {
	return any.__obf_efaf2719b44ad8ac
}

func (any *__obf_b2998b74708f8d59) WriteTo(__obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteString(any.__obf_efaf2719b44ad8ac)
}

func (any *__obf_b2998b74708f8d59) GetInterface() any {
	return any.__obf_efaf2719b44ad8ac
}
