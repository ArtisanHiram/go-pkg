package __obf_c7b79b12b33d8238

import (
	"fmt"
	"strconv"
)

type __obf_755c14a10251d761 struct {
	__obf_3a44d9c5204c2e12
	__obf_35accd096612ebe4 string
}

func (any *__obf_755c14a10251d761) Get(__obf_5dde9fb4007294d4 ...any) Any {
	if len(__obf_5dde9fb4007294d4) == 0 {
		return any
	}
	return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, fmt.Errorf("GetIndex %v from simple value", __obf_5dde9fb4007294d4)}
}

func (any *__obf_755c14a10251d761) Parse() *Iterator {
	return nil
}

func (any *__obf_755c14a10251d761) ValueType() ValueType {
	return StringValue
}

func (any *__obf_755c14a10251d761) MustBeValid() Any {
	return any
}

func (any *__obf_755c14a10251d761) LastError() error {
	return nil
}

func (any *__obf_755c14a10251d761) ToBool() bool {
	__obf_a3eaafc22faf1644 := any.ToString()
	if __obf_a3eaafc22faf1644 == "0" {
		return false
	}
	for _, __obf_12577c03a12f0d2c := range __obf_a3eaafc22faf1644 {
		switch __obf_12577c03a12f0d2c {
		case ' ', '\n', '\r', '\t':
		default:
			return true
		}
	}
	return false
}

func (any *__obf_755c14a10251d761) ToInt() int {
	return int(any.ToInt64())

}

func (any *__obf_755c14a10251d761) ToInt32() int32 {
	return int32(any.ToInt64())
}

func (any *__obf_755c14a10251d761) ToInt64() int64 {
	if any.__obf_35accd096612ebe4 == "" {
		return 0
	}
	__obf_4e2a3de3b6caf261 := 1
	__obf_998122590957c742 := 0
	if any.__obf_35accd096612ebe4[0] == '+' || any.__obf_35accd096612ebe4[0] == '-' {
		__obf_998122590957c742 = 1
	}

	if any.__obf_35accd096612ebe4[0] == '-' {
		__obf_4e2a3de3b6caf261 = -1
	}
	__obf_bcce94b13441b16c := __obf_998122590957c742
	for __obf_a051269af3a541bb := __obf_998122590957c742; __obf_a051269af3a541bb < len(any.__obf_35accd096612ebe4); __obf_a051269af3a541bb++ {
		if any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] >= '0' && any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] <= '9' {
			__obf_bcce94b13441b16c = __obf_a051269af3a541bb + 1
		} else {
			break
		}
	}
	__obf_95728f859dd49e84, _ := strconv.ParseInt(any.__obf_35accd096612ebe4[__obf_998122590957c742:__obf_bcce94b13441b16c], 10, 64)
	return int64(__obf_4e2a3de3b6caf261) * __obf_95728f859dd49e84
}

func (any *__obf_755c14a10251d761) ToUint() uint {
	return uint(any.ToUint64())
}

func (any *__obf_755c14a10251d761) ToUint32() uint32 {
	return uint32(any.ToUint64())
}

func (any *__obf_755c14a10251d761) ToUint64() uint64 {
	if any.__obf_35accd096612ebe4 == "" {
		return 0
	}
	__obf_998122590957c742 := 0

	if any.__obf_35accd096612ebe4[0] == '-' {
		return 0
	}
	if any.__obf_35accd096612ebe4[0] == '+' {
		__obf_998122590957c742 = 1
	}
	__obf_bcce94b13441b16c := __obf_998122590957c742
	for __obf_a051269af3a541bb := __obf_998122590957c742; __obf_a051269af3a541bb < len(any.__obf_35accd096612ebe4); __obf_a051269af3a541bb++ {
		if any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] >= '0' && any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] <= '9' {
			__obf_bcce94b13441b16c = __obf_a051269af3a541bb + 1
		} else {
			break
		}
	}
	__obf_95728f859dd49e84, _ := strconv.ParseUint(any.__obf_35accd096612ebe4[__obf_998122590957c742:__obf_bcce94b13441b16c], 10, 64)
	return __obf_95728f859dd49e84
}

func (any *__obf_755c14a10251d761) ToFloat32() float32 {
	return float32(any.ToFloat64())
}

func (any *__obf_755c14a10251d761) ToFloat64() float64 {
	if len(any.__obf_35accd096612ebe4) == 0 {
		return 0
	}

	// first char invalid
	if any.__obf_35accd096612ebe4[0] != '+' && any.__obf_35accd096612ebe4[0] != '-' && (any.__obf_35accd096612ebe4[0] > '9' || any.__obf_35accd096612ebe4[0] < '0') {
		return 0
	}
	__obf_bcce94b13441b16c := // extract valid num expression from string
		// eg 123true => 123, -12.12xxa => -12.12
		1
	for __obf_a051269af3a541bb := 1; __obf_a051269af3a541bb < len(any.__obf_35accd096612ebe4); __obf_a051269af3a541bb++ {
		if any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] == '.' || any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] == 'e' || any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] == 'E' || any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] == '+' || any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] == '-' {
			__obf_bcce94b13441b16c = __obf_a051269af3a541bb + 1
			continue
		}

		// end position is the first char which is not digit
		if any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] >= '0' && any.__obf_35accd096612ebe4[__obf_a051269af3a541bb] <= '9' {
			__obf_bcce94b13441b16c = __obf_a051269af3a541bb + 1
		} else {
			__obf_bcce94b13441b16c = __obf_a051269af3a541bb
			break
		}
	}
	__obf_95728f859dd49e84, _ := strconv.ParseFloat(any.__obf_35accd096612ebe4[:__obf_bcce94b13441b16c], 64)
	return __obf_95728f859dd49e84
}

func (any *__obf_755c14a10251d761) ToString() string {
	return any.__obf_35accd096612ebe4
}

func (any *__obf_755c14a10251d761) WriteTo(__obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteString(any.__obf_35accd096612ebe4)
}

func (any *__obf_755c14a10251d761) GetInterface() any {
	return any.__obf_35accd096612ebe4
}
