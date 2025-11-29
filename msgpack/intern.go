package __obf_3edaa49e853afa16

import (
	"fmt"
	"math"
	"reflect"
)

const (
	__obf_fbea0ec178356358 = 3
	__obf_751a7901cea85832 = math.MaxUint16
)

var __obf_6481344dc83dee01 = int8(math.MinInt8)

func init() {
	__obf_733cc19b5e047a7a[__obf_6481344dc83dee01] = &__obf_4e4d1cbf2e9fdd58{
		Type:    __obf_2566dc21412a8f7b,
		Decoder: __obf_02668242b31df084,
	}
}

func __obf_02668242b31df084(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value, __obf_69375c248ca4aa99 int) error {
	__obf_eff265c728cf9640, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_1383c4deca08b1f8(__obf_69375c248ca4aa99)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_a05969ab9ce027e7(__obf_eff265c728cf9640)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetString(__obf_6827ff1b59d7ccec)
	return nil
}

//------------------------------------------------------------------------------

func __obf_7ad73e7d72971622(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Elem()
	if __obf_85d270343a0dfe14.Kind() == reflect.String {
		return __obf_84d0d31a8288f191.__obf_fd9d10636b1b890e(__obf_85d270343a0dfe14.String(), true)
	}
	return __obf_84d0d31a8288f191.EncodeValue(__obf_85d270343a0dfe14)
}

func __obf_25ae4057eed7711b(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.__obf_fd9d10636b1b890e(__obf_85d270343a0dfe14.String(), true)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_fd9d10636b1b890e(__obf_6827ff1b59d7ccec string, __obf_60ddbad9c6f41c91 bool) error {
	// Interned string takes at least 3 bytes. Plain string 1 byte + string len.
	if __obf_eff265c728cf9640, __obf_ccfb92cc26e4569f := __obf_84d0d31a8288f191.__obf_b014355f64826d7b[__obf_6827ff1b59d7ccec]; __obf_ccfb92cc26e4569f {
		return __obf_84d0d31a8288f191.__obf_2813f84dc17a3406(__obf_eff265c728cf9640)
	}

	if __obf_60ddbad9c6f41c91 && len(__obf_6827ff1b59d7ccec) >= __obf_fbea0ec178356358 && len(__obf_84d0d31a8288f191.__obf_b014355f64826d7b) < __obf_751a7901cea85832 {
		if __obf_84d0d31a8288f191.__obf_b014355f64826d7b == nil {
			__obf_84d0d31a8288f191.__obf_b014355f64826d7b = make(map[string]int)
		}
		__obf_eff265c728cf9640 := len(__obf_84d0d31a8288f191.__obf_b014355f64826d7b)
		__obf_84d0d31a8288f191.__obf_b014355f64826d7b[__obf_6827ff1b59d7ccec] = __obf_eff265c728cf9640
	}

	return __obf_84d0d31a8288f191.__obf_7b5be37728f8302b(__obf_6827ff1b59d7ccec)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_2813f84dc17a3406(__obf_eff265c728cf9640 int) error {
	if __obf_eff265c728cf9640 <= math.MaxUint8 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt1); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		return __obf_84d0d31a8288f191.__obf_59b14e1d8bb8dea4(byte(__obf_6481344dc83dee01), uint8(__obf_eff265c728cf9640))
	}

	if __obf_eff265c728cf9640 <= math.MaxUint16 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt2); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(byte(__obf_6481344dc83dee01), uint16(__obf_eff265c728cf9640))
	}

	if uint64(__obf_eff265c728cf9640) <= math.MaxUint32 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt4); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		return __obf_84d0d31a8288f191.__obf_7636556325987942(byte(__obf_6481344dc83dee01), uint32(__obf_eff265c728cf9640))
	}

	return fmt.Errorf("msgpack: interned string index=%d is too large", __obf_eff265c728cf9640)
}

//------------------------------------------------------------------------------

func __obf_867c9a7f5c7de360(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_bc64b9136455b38b(true)
	if __obf_3aa78ad28f50ed46 == nil {
		__obf_85d270343a0dfe14.
			Set(reflect.ValueOf(__obf_6827ff1b59d7ccec))
		return nil
	}
	if _, __obf_ccfb92cc26e4569f := __obf_3aa78ad28f50ed46.(__obf_6b688bac335393d7); !__obf_ccfb92cc26e4569f {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.UnreadByte(); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return __obf_6b9d6a0bab505ca5(__obf_015afbee33862a01, __obf_85d270343a0dfe14)
}

func __obf_0156ce491a786988(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_bc64b9136455b38b(true)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetString(__obf_6827ff1b59d7ccec)
	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_bc64b9136455b38b(__obf_60ddbad9c6f41c91 bool) (string, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}

	if IsFixedString(__obf_145c7a7d25eea2bd) {
		__obf_56127cd370854f0b := int(__obf_145c7a7d25eea2bd & FixedStrMask)
		return __obf_015afbee33862a01.__obf_666abe5e67801583(__obf_56127cd370854f0b, __obf_60ddbad9c6f41c91)
	}

	switch __obf_145c7a7d25eea2bd {
	case Nil:
		return "", nil
	case FixExt1, FixExt2, FixExt4:
		__obf_8f6a3347b425706d, __obf_69375c248ca4aa99, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_9d268a1ca1cebd45(__obf_145c7a7d25eea2bd)
		if __obf_3aa78ad28f50ed46 != nil {
			return "", __obf_3aa78ad28f50ed46
		}
		if __obf_8f6a3347b425706d != __obf_6481344dc83dee01 {
			__obf_3aa78ad28f50ed46 := fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_8f6a3347b425706d, __obf_6481344dc83dee01)
			return "", __obf_3aa78ad28f50ed46
		}
		__obf_eff265c728cf9640, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_1383c4deca08b1f8(__obf_69375c248ca4aa99)
		if __obf_3aa78ad28f50ed46 != nil {
			return "", __obf_3aa78ad28f50ed46
		}

		return __obf_015afbee33862a01.__obf_a05969ab9ce027e7(__obf_eff265c728cf9640)
	case Str8, Bin8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		if __obf_3aa78ad28f50ed46 != nil {
			return "", __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_666abe5e67801583(int(__obf_56127cd370854f0b), __obf_60ddbad9c6f41c91)
	case Str16, Bin16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		if __obf_3aa78ad28f50ed46 != nil {
			return "", __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_666abe5e67801583(int(__obf_56127cd370854f0b), __obf_60ddbad9c6f41c91)
	case Str32, Bin32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		if __obf_3aa78ad28f50ed46 != nil {
			return "", __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_666abe5e67801583(int(__obf_56127cd370854f0b), __obf_60ddbad9c6f41c91)
	}

	return "", __obf_6b688bac335393d7{__obf_508e2d6ff53655c0: __obf_145c7a7d25eea2bd, __obf_6e0dac6e0c9d5f66: "interned string"}
}

func (__obf_015afbee33862a01 *Decoder) __obf_1383c4deca08b1f8(__obf_69375c248ca4aa99 int) (int, error) {
	switch __obf_69375c248ca4aa99 {
	case 1:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		return int(__obf_56127cd370854f0b), nil
	case 2:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		return int(__obf_56127cd370854f0b), nil
	case 4:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		if __obf_3aa78ad28f50ed46 != nil {
			return 0, __obf_3aa78ad28f50ed46
		}
		return int(__obf_56127cd370854f0b), nil
	}
	__obf_3aa78ad28f50ed46 := fmt.Errorf("msgpack: unsupported ext len=%d decoding interned string", __obf_69375c248ca4aa99)
	return 0, __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) __obf_a05969ab9ce027e7(__obf_eff265c728cf9640 int) (string, error) {
	if __obf_eff265c728cf9640 >= len(__obf_015afbee33862a01.__obf_b014355f64826d7b) {
		__obf_3aa78ad28f50ed46 := fmt.Errorf("msgpack: interned string at index=%d does not exist", __obf_eff265c728cf9640)
		return "", __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_b014355f64826d7b[__obf_eff265c728cf9640], nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_666abe5e67801583(__obf_56127cd370854f0b int, __obf_60ddbad9c6f41c91 bool) (string, error) {
	if __obf_56127cd370854f0b <= 0 {
		return "", nil
	}
	__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_b85f6957ac631eeb(__obf_56127cd370854f0b)
	if __obf_3aa78ad28f50ed46 != nil {
		return "", __obf_3aa78ad28f50ed46
	}

	if __obf_60ddbad9c6f41c91 && len(__obf_6827ff1b59d7ccec) >= __obf_fbea0ec178356358 && len(__obf_015afbee33862a01.__obf_b014355f64826d7b) < __obf_751a7901cea85832 {
		__obf_015afbee33862a01.__obf_b014355f64826d7b = append(__obf_015afbee33862a01.__obf_b014355f64826d7b, __obf_6827ff1b59d7ccec)
	}

	return __obf_6827ff1b59d7ccec, nil
}
