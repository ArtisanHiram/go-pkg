package __obf_fd770cb9675903b0

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

var __obf_643e9bce4c8d9cbe int8 = -1

func init() {
	RegisterExtEncoder(__obf_643e9bce4c8d9cbe, time.Time{}, __obf_66d94ab08696d001)
	RegisterExtDecoder(__obf_643e9bce4c8d9cbe, time.Time{}, __obf_07d50de8ea00f946)
}

func __obf_66d94ab08696d001(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) ([]byte, error) {
	return __obf_e9038cda3b5cf711.__obf_5afccbe157f69d06(__obf_f328a048f76b7256.Interface().(time.Time)), nil
}

func __obf_07d50de8ea00f946(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value, __obf_803bccbf05dad29c int) error {
	__obf_9f3a870c077f71cc, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_b29b5dde3a7b779b(__obf_803bccbf05dad29c)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	if __obf_9f3a870c077f71cc.IsZero() {
		__obf_9f3a870c077f71cc = // Zero time does not have timezone information.
			__obf_9f3a870c077f71cc.UTC()
	}
	__obf_2c49f9d2007cfaf6 := __obf_f328a048f76b7256.Addr().Interface().(*time.Time)
	*__obf_2c49f9d2007cfaf6 = __obf_9f3a870c077f71cc

	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeTime(__obf_9f3a870c077f71cc time.Time) error {
	__obf_94333af0f0facd65 := __obf_e9038cda3b5cf711.__obf_5afccbe157f69d06(__obf_9f3a870c077f71cc)
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_5926ca009cd80dd3(len(__obf_94333af0f0facd65)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.WriteByte(byte(__obf_643e9bce4c8d9cbe)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_94333af0f0facd65)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_5afccbe157f69d06(__obf_9f3a870c077f71cc time.Time) []byte {
	if __obf_e9038cda3b5cf711.__obf_07a810e976869435 == nil {
		__obf_e9038cda3b5cf711.__obf_07a810e976869435 = make([]byte, 12)
	}
	__obf_7ec9f28e9a1e0977 := uint64(__obf_9f3a870c077f71cc.Unix())
	if __obf_7ec9f28e9a1e0977>>34 == 0 {
		__obf_cc76e8ed73142f28 := uint64(__obf_9f3a870c077f71cc.Nanosecond())<<34 | __obf_7ec9f28e9a1e0977

		if __obf_cc76e8ed73142f28&0xffffffff00000000 == 0 {
			__obf_94333af0f0facd65 := __obf_e9038cda3b5cf711.__obf_07a810e976869435[:4]
			binary.BigEndian.PutUint32(__obf_94333af0f0facd65, uint32(__obf_cc76e8ed73142f28))
			return __obf_94333af0f0facd65
		}
		__obf_94333af0f0facd65 := __obf_e9038cda3b5cf711.__obf_07a810e976869435[:8]
		binary.BigEndian.PutUint64(__obf_94333af0f0facd65, __obf_cc76e8ed73142f28)
		return __obf_94333af0f0facd65
	}
	__obf_94333af0f0facd65 := __obf_e9038cda3b5cf711.__obf_07a810e976869435[:12]
	binary.BigEndian.PutUint32(__obf_94333af0f0facd65, uint32(__obf_9f3a870c077f71cc.Nanosecond()))
	binary.BigEndian.PutUint64(__obf_94333af0f0facd65[4:], __obf_7ec9f28e9a1e0977)
	return __obf_94333af0f0facd65
}

func (__obf_5d389b660eefb08c *Decoder) DecodeTime() (time.Time, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return time.Time{}, __obf_45342a3a754d12f5
	}

	// Legacy format.
	if __obf_4148125b350dfea2 == FixedArrayLow|2 {
		__obf_0fdab1b51d94bd76, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
		if __obf_45342a3a754d12f5 != nil {
			return time.Time{}, __obf_45342a3a754d12f5
		}
		__obf_0561208db65409f6, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
		if __obf_45342a3a754d12f5 != nil {
			return time.Time{}, __obf_45342a3a754d12f5
		}

		return time.Unix(__obf_0fdab1b51d94bd76, __obf_0561208db65409f6), nil
	}

	if IsString(__obf_4148125b350dfea2) {
		__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.string(__obf_4148125b350dfea2)
		if __obf_45342a3a754d12f5 != nil {
			return time.Time{}, __obf_45342a3a754d12f5
		}
		return time.Parse(time.RFC3339Nano, __obf_fe1ace7a2eb8bf9f)
	}
	__obf_e728c4abb2dcc444, __obf_803bccbf05dad29c, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_051bbe645d4e1aa5(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return time.Time{}, __obf_45342a3a754d12f5
	}

	// NodeJS seems to use extID 13.
	if __obf_e728c4abb2dcc444 != __obf_643e9bce4c8d9cbe && __obf_e728c4abb2dcc444 != 13 {
		return time.Time{}, fmt.Errorf("msgpack: invalid time ext id=%d", __obf_e728c4abb2dcc444)
	}
	__obf_9f3a870c077f71cc, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_b29b5dde3a7b779b(__obf_803bccbf05dad29c)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_9f3a870c077f71cc, __obf_45342a3a754d12f5
	}

	if __obf_9f3a870c077f71cc.IsZero() {
		// Zero time does not have timezone information.
		return __obf_9f3a870c077f71cc.UTC(), nil
	}
	return __obf_9f3a870c077f71cc, nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_b29b5dde3a7b779b(__obf_803bccbf05dad29c int) (time.Time, error) {
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(__obf_803bccbf05dad29c)
	if __obf_45342a3a754d12f5 != nil {
		return time.Time{}, __obf_45342a3a754d12f5
	}

	switch len(__obf_94333af0f0facd65) {
	case 4:
		__obf_0fdab1b51d94bd76 := binary.BigEndian.Uint32(__obf_94333af0f0facd65)
		return time.Unix(int64(__obf_0fdab1b51d94bd76), 0), nil
	case 8:
		__obf_0fdab1b51d94bd76 := binary.BigEndian.Uint64(__obf_94333af0f0facd65)
		__obf_0561208db65409f6 := int64(__obf_0fdab1b51d94bd76 >> 34)
		__obf_0fdab1b51d94bd76 &= 0x00000003ffffffff
		return time.Unix(int64(__obf_0fdab1b51d94bd76), __obf_0561208db65409f6), nil
	case 12:
		__obf_0561208db65409f6 := binary.BigEndian.Uint32(__obf_94333af0f0facd65)
		__obf_0fdab1b51d94bd76 := binary.BigEndian.Uint64(__obf_94333af0f0facd65[4:])
		return time.Unix(int64(__obf_0fdab1b51d94bd76), int64(__obf_0561208db65409f6)), nil
	default:
		__obf_45342a3a754d12f5 = fmt.Errorf("msgpack: invalid ext len=%d decoding time", __obf_803bccbf05dad29c)
		return time.Time{}, __obf_45342a3a754d12f5
	}
}
