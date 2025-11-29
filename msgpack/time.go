package __obf_3edaa49e853afa16

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

var __obf_76bb45f824925051 int8 = -1

func init() {
	RegisterExtEncoder(__obf_76bb45f824925051, time.Time{}, __obf_fbaef019f3f79e57)
	RegisterExtDecoder(__obf_76bb45f824925051, time.Time{}, __obf_b9f64c10a6f12261)
}

func __obf_fbaef019f3f79e57(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) ([]byte, error) {
	return __obf_84d0d31a8288f191.__obf_920bb677e77cf318(__obf_85d270343a0dfe14.Interface().(time.Time)), nil
}

func __obf_b9f64c10a6f12261(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value, __obf_69375c248ca4aa99 int) error {
	__obf_f4a1d268d7ae3788, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_e1fd4c2c8b85e4d1(__obf_69375c248ca4aa99)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_f4a1d268d7ae3788.IsZero() {
		__obf_f4a1d268d7ae3788 = // Zero time does not have timezone information.
			__obf_f4a1d268d7ae3788.UTC()
	}
	__obf_8f0c71619c0382f1 := __obf_85d270343a0dfe14.Addr().Interface().(*time.Time)
	*__obf_8f0c71619c0382f1 = __obf_f4a1d268d7ae3788

	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeTime(__obf_f4a1d268d7ae3788 time.Time) error {
	__obf_9b4a5a04bdad2347 := __obf_84d0d31a8288f191.__obf_920bb677e77cf318(__obf_f4a1d268d7ae3788)
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_257d74d0a2026131(len(__obf_9b4a5a04bdad2347)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.WriteByte(byte(__obf_76bb45f824925051)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_9b4a5a04bdad2347)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_920bb677e77cf318(__obf_f4a1d268d7ae3788 time.Time) []byte {
	if __obf_84d0d31a8288f191.__obf_c0eccc33acc16ac1 == nil {
		__obf_84d0d31a8288f191.__obf_c0eccc33acc16ac1 = make([]byte, 12)
	}
	__obf_ae5698575d282e5a := uint64(__obf_f4a1d268d7ae3788.Unix())
	if __obf_ae5698575d282e5a>>34 == 0 {
		__obf_2171bf5dcbdd2ac8 := uint64(__obf_f4a1d268d7ae3788.Nanosecond())<<34 | __obf_ae5698575d282e5a

		if __obf_2171bf5dcbdd2ac8&0xffffffff00000000 == 0 {
			__obf_9b4a5a04bdad2347 := __obf_84d0d31a8288f191.__obf_c0eccc33acc16ac1[:4]
			binary.BigEndian.PutUint32(__obf_9b4a5a04bdad2347, uint32(__obf_2171bf5dcbdd2ac8))
			return __obf_9b4a5a04bdad2347
		}
		__obf_9b4a5a04bdad2347 := __obf_84d0d31a8288f191.__obf_c0eccc33acc16ac1[:8]
		binary.BigEndian.PutUint64(__obf_9b4a5a04bdad2347, __obf_2171bf5dcbdd2ac8)
		return __obf_9b4a5a04bdad2347
	}
	__obf_9b4a5a04bdad2347 := __obf_84d0d31a8288f191.__obf_c0eccc33acc16ac1[:12]
	binary.BigEndian.PutUint32(__obf_9b4a5a04bdad2347, uint32(__obf_f4a1d268d7ae3788.Nanosecond()))
	binary.BigEndian.PutUint64(__obf_9b4a5a04bdad2347[4:], __obf_ae5698575d282e5a)
	return __obf_9b4a5a04bdad2347
}

func (__obf_015afbee33862a01 *Decoder) DecodeTime() (time.Time, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return time.Time{}, __obf_3aa78ad28f50ed46
	}

	// Legacy format.
	if __obf_145c7a7d25eea2bd == FixedArrayLow|2 {
		__obf_c398567bd3b96122, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
		if __obf_3aa78ad28f50ed46 != nil {
			return time.Time{}, __obf_3aa78ad28f50ed46
		}
		__obf_64ead7597ef9c86c, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
		if __obf_3aa78ad28f50ed46 != nil {
			return time.Time{}, __obf_3aa78ad28f50ed46
		}

		return time.Unix(__obf_c398567bd3b96122, __obf_64ead7597ef9c86c), nil
	}

	if IsString(__obf_145c7a7d25eea2bd) {
		__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.string(__obf_145c7a7d25eea2bd)
		if __obf_3aa78ad28f50ed46 != nil {
			return time.Time{}, __obf_3aa78ad28f50ed46
		}
		return time.Parse(time.RFC3339Nano, __obf_6827ff1b59d7ccec)
	}
	__obf_82dbf2151b1070df, __obf_69375c248ca4aa99, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_9d268a1ca1cebd45(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return time.Time{}, __obf_3aa78ad28f50ed46
	}

	// NodeJS seems to use extID 13.
	if __obf_82dbf2151b1070df != __obf_76bb45f824925051 && __obf_82dbf2151b1070df != 13 {
		return time.Time{}, fmt.Errorf("msgpack: invalid time ext id=%d", __obf_82dbf2151b1070df)
	}
	__obf_f4a1d268d7ae3788, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_e1fd4c2c8b85e4d1(__obf_69375c248ca4aa99)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_f4a1d268d7ae3788, __obf_3aa78ad28f50ed46
	}

	if __obf_f4a1d268d7ae3788.IsZero() {
		// Zero time does not have timezone information.
		return __obf_f4a1d268d7ae3788.UTC(), nil
	}
	return __obf_f4a1d268d7ae3788, nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_e1fd4c2c8b85e4d1(__obf_69375c248ca4aa99 int) (time.Time, error) {
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(__obf_69375c248ca4aa99)
	if __obf_3aa78ad28f50ed46 != nil {
		return time.Time{}, __obf_3aa78ad28f50ed46
	}

	switch len(__obf_9b4a5a04bdad2347) {
	case 4:
		__obf_c398567bd3b96122 := binary.BigEndian.Uint32(__obf_9b4a5a04bdad2347)
		return time.Unix(int64(__obf_c398567bd3b96122), 0), nil
	case 8:
		__obf_c398567bd3b96122 := binary.BigEndian.Uint64(__obf_9b4a5a04bdad2347)
		__obf_64ead7597ef9c86c := int64(__obf_c398567bd3b96122 >> 34)
		__obf_c398567bd3b96122 &= 0x00000003ffffffff
		return time.Unix(int64(__obf_c398567bd3b96122), __obf_64ead7597ef9c86c), nil
	case 12:
		__obf_64ead7597ef9c86c := binary.BigEndian.Uint32(__obf_9b4a5a04bdad2347)
		__obf_c398567bd3b96122 := binary.BigEndian.Uint64(__obf_9b4a5a04bdad2347[4:])
		return time.Unix(int64(__obf_c398567bd3b96122), int64(__obf_64ead7597ef9c86c)), nil
	default:
		__obf_3aa78ad28f50ed46 = fmt.Errorf("msgpack: invalid ext len=%d decoding time", __obf_69375c248ca4aa99)
		return time.Time{}, __obf_3aa78ad28f50ed46
	}
}
