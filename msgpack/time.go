package __obf_a935eb7f548271a4

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

var __obf_c7c703b8d4ff871b int8 = -1

func init() {
	RegisterExtEncoder(__obf_c7c703b8d4ff871b, time.Time{}, __obf_7ffdadd4ac85161b)
	RegisterExtDecoder(__obf_c7c703b8d4ff871b, time.Time{}, __obf_8d47a456193d6c4b)
}

func __obf_7ffdadd4ac85161b(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) ([]byte, error) {
	return __obf_ed37a34c347049f3.__obf_910a67dae1baa476(__obf_6d570581f4b60dbc.Interface().(time.Time)), nil
}

func __obf_8d47a456193d6c4b(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value, __obf_f9e21d535a3562ea int) error {
	__obf_7ddb8a02bb5e5afb, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_5369cac9df16de45(__obf_f9e21d535a3562ea)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_7ddb8a02bb5e5afb.IsZero() {
		__obf_7ddb8a02bb5e5afb = // Zero time does not have timezone information.
			__obf_7ddb8a02bb5e5afb.UTC()
	}
	__obf_0d8a994785cda6df := __obf_6d570581f4b60dbc.Addr().Interface().(*time.Time)
	*__obf_0d8a994785cda6df = __obf_7ddb8a02bb5e5afb

	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeTime(__obf_7ddb8a02bb5e5afb time.Time) error {
	__obf_f2ca794293605b73 := __obf_ed37a34c347049f3.__obf_910a67dae1baa476(__obf_7ddb8a02bb5e5afb)
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_e3e9e2d425e2405b(len(__obf_f2ca794293605b73)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.WriteByte(byte(__obf_c7c703b8d4ff871b)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_f2ca794293605b73)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_910a67dae1baa476(__obf_7ddb8a02bb5e5afb time.Time) []byte {
	if __obf_ed37a34c347049f3.__obf_f1da64bc48a2ea7f == nil {
		__obf_ed37a34c347049f3.__obf_f1da64bc48a2ea7f = make([]byte, 12)
	}
	__obf_b742f40a328cc53b := uint64(__obf_7ddb8a02bb5e5afb.Unix())
	if __obf_b742f40a328cc53b>>34 == 0 {
		__obf_652c67787b9c24c9 := uint64(__obf_7ddb8a02bb5e5afb.Nanosecond())<<34 | __obf_b742f40a328cc53b

		if __obf_652c67787b9c24c9&0xffffffff00000000 == 0 {
			__obf_f2ca794293605b73 := __obf_ed37a34c347049f3.__obf_f1da64bc48a2ea7f[:4]
			binary.BigEndian.PutUint32(__obf_f2ca794293605b73, uint32(__obf_652c67787b9c24c9))
			return __obf_f2ca794293605b73
		}
		__obf_f2ca794293605b73 := __obf_ed37a34c347049f3.__obf_f1da64bc48a2ea7f[:8]
		binary.BigEndian.PutUint64(__obf_f2ca794293605b73, __obf_652c67787b9c24c9)
		return __obf_f2ca794293605b73
	}
	__obf_f2ca794293605b73 := __obf_ed37a34c347049f3.__obf_f1da64bc48a2ea7f[:12]
	binary.BigEndian.PutUint32(__obf_f2ca794293605b73, uint32(__obf_7ddb8a02bb5e5afb.Nanosecond()))
	binary.BigEndian.PutUint64(__obf_f2ca794293605b73[4:], __obf_b742f40a328cc53b)
	return __obf_f2ca794293605b73
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeTime() (time.Time, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return time.Time{}, __obf_4d327e1cd40c2e21
	}

	// Legacy format.
	if __obf_f5df560f4d67421b == FixedArrayLow|2 {
		__obf_a855f991afef8313, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
		if __obf_4d327e1cd40c2e21 != nil {
			return time.Time{}, __obf_4d327e1cd40c2e21
		}
		__obf_81e61e7331187657, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
		if __obf_4d327e1cd40c2e21 != nil {
			return time.Time{}, __obf_4d327e1cd40c2e21
		}

		return time.Unix(__obf_a855f991afef8313, __obf_81e61e7331187657), nil
	}

	if IsString(__obf_f5df560f4d67421b) {
		__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.string(__obf_f5df560f4d67421b)
		if __obf_4d327e1cd40c2e21 != nil {
			return time.Time{}, __obf_4d327e1cd40c2e21
		}
		return time.Parse(time.RFC3339Nano, __obf_b62c60fba2fd9788)
	}
	__obf_12dd79d80139d0ca, __obf_f9e21d535a3562ea, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_4dc04de1e73ebdac(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return time.Time{}, __obf_4d327e1cd40c2e21
	}

	// NodeJS seems to use extID 13.
	if __obf_12dd79d80139d0ca != __obf_c7c703b8d4ff871b && __obf_12dd79d80139d0ca != 13 {
		return time.Time{}, fmt.Errorf("msgpack: invalid time ext id=%d", __obf_12dd79d80139d0ca)
	}
	__obf_7ddb8a02bb5e5afb, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_5369cac9df16de45(__obf_f9e21d535a3562ea)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_7ddb8a02bb5e5afb, __obf_4d327e1cd40c2e21
	}

	if __obf_7ddb8a02bb5e5afb.IsZero() {
		// Zero time does not have timezone information.
		return __obf_7ddb8a02bb5e5afb.UTC(), nil
	}
	return __obf_7ddb8a02bb5e5afb, nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_5369cac9df16de45(__obf_f9e21d535a3562ea int) (time.Time, error) {
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(__obf_f9e21d535a3562ea)
	if __obf_4d327e1cd40c2e21 != nil {
		return time.Time{}, __obf_4d327e1cd40c2e21
	}

	switch len(__obf_f2ca794293605b73) {
	case 4:
		__obf_a855f991afef8313 := binary.BigEndian.Uint32(__obf_f2ca794293605b73)
		return time.Unix(int64(__obf_a855f991afef8313), 0), nil
	case 8:
		__obf_a855f991afef8313 := binary.BigEndian.Uint64(__obf_f2ca794293605b73)
		__obf_81e61e7331187657 := int64(__obf_a855f991afef8313 >> 34)
		__obf_a855f991afef8313 &= 0x00000003ffffffff
		return time.Unix(int64(__obf_a855f991afef8313), __obf_81e61e7331187657), nil
	case 12:
		__obf_81e61e7331187657 := binary.BigEndian.Uint32(__obf_f2ca794293605b73)
		__obf_a855f991afef8313 := binary.BigEndian.Uint64(__obf_f2ca794293605b73[4:])
		return time.Unix(int64(__obf_a855f991afef8313), int64(__obf_81e61e7331187657)), nil
	default:
		__obf_4d327e1cd40c2e21 = fmt.Errorf("msgpack: invalid ext len=%d decoding time", __obf_f9e21d535a3562ea)
		return time.Time{}, __obf_4d327e1cd40c2e21
	}
}
