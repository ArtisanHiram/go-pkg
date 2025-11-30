package __obf_3e0c303610a19bd4

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

var __obf_2b109b947f8a4e95 int8 = -1

func init() {
	RegisterExtEncoder(__obf_2b109b947f8a4e95, time.Time{}, __obf_11b95691ed44c1a0)
	RegisterExtDecoder(__obf_2b109b947f8a4e95, time.Time{}, __obf_695b65a0331f81bd)
}

func __obf_11b95691ed44c1a0(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) ([]byte, error) {
	return __obf_77240dc7776b60b4.__obf_e2fa58dcc22cf3a2(__obf_63bbcee86d44fdde.Interface().(time.Time)), nil
}

func __obf_695b65a0331f81bd(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value, __obf_7e99dd6aa706c57b int) error {
	__obf_75953c9a45f5db9e, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_7fa5d5dcdfa5e1ed(__obf_7e99dd6aa706c57b)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	if __obf_75953c9a45f5db9e.IsZero() {
		__obf_75953c9a45f5db9e = // Zero time does not have timezone information.
			__obf_75953c9a45f5db9e.UTC()
	}
	__obf_b5a4664807537c0d := __obf_63bbcee86d44fdde.Addr().Interface().(*time.Time)
	*__obf_b5a4664807537c0d = __obf_75953c9a45f5db9e

	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeTime(__obf_75953c9a45f5db9e time.Time) error {
	__obf_11bcc66cde095c11 := __obf_77240dc7776b60b4.__obf_e2fa58dcc22cf3a2(__obf_75953c9a45f5db9e)
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_7a9a2753ca4f49af(len(__obf_11bcc66cde095c11)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.WriteByte(byte(__obf_2b109b947f8a4e95)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_11bcc66cde095c11)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_e2fa58dcc22cf3a2(__obf_75953c9a45f5db9e time.Time) []byte {
	if __obf_77240dc7776b60b4.__obf_42dab97af35e666e == nil {
		__obf_77240dc7776b60b4.__obf_42dab97af35e666e = make([]byte, 12)
	}
	__obf_8630d8679f495d6c := uint64(__obf_75953c9a45f5db9e.Unix())
	if __obf_8630d8679f495d6c>>34 == 0 {
		__obf_4d827abff69b9b40 := uint64(__obf_75953c9a45f5db9e.Nanosecond())<<34 | __obf_8630d8679f495d6c

		if __obf_4d827abff69b9b40&0xffffffff00000000 == 0 {
			__obf_11bcc66cde095c11 := __obf_77240dc7776b60b4.__obf_42dab97af35e666e[:4]
			binary.BigEndian.PutUint32(__obf_11bcc66cde095c11, uint32(__obf_4d827abff69b9b40))
			return __obf_11bcc66cde095c11
		}
		__obf_11bcc66cde095c11 := __obf_77240dc7776b60b4.__obf_42dab97af35e666e[:8]
		binary.BigEndian.PutUint64(__obf_11bcc66cde095c11, __obf_4d827abff69b9b40)
		return __obf_11bcc66cde095c11
	}
	__obf_11bcc66cde095c11 := __obf_77240dc7776b60b4.__obf_42dab97af35e666e[:12]
	binary.BigEndian.PutUint32(__obf_11bcc66cde095c11, uint32(__obf_75953c9a45f5db9e.Nanosecond()))
	binary.BigEndian.PutUint64(__obf_11bcc66cde095c11[4:], __obf_8630d8679f495d6c)
	return __obf_11bcc66cde095c11
}

func (__obf_dc35117108ba8439 *Decoder) DecodeTime() (time.Time, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return time.Time{}, __obf_8882f6cf6e378ded
	}

	// Legacy format.
	if __obf_e46289218af709bf == FixedArrayLow|2 {
		__obf_ad3d72f686eb35a4, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
		if __obf_8882f6cf6e378ded != nil {
			return time.Time{}, __obf_8882f6cf6e378ded
		}
		__obf_6617171abe4523cd, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
		if __obf_8882f6cf6e378ded != nil {
			return time.Time{}, __obf_8882f6cf6e378ded
		}

		return time.Unix(__obf_ad3d72f686eb35a4, __obf_6617171abe4523cd), nil
	}

	if IsString(__obf_e46289218af709bf) {
		__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.string(__obf_e46289218af709bf)
		if __obf_8882f6cf6e378ded != nil {
			return time.Time{}, __obf_8882f6cf6e378ded
		}
		return time.Parse(time.RFC3339Nano, __obf_61027e0491b6dd3d)
	}
	__obf_927f1edb79ba3be1, __obf_7e99dd6aa706c57b, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_3ac0957b06c080b5(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return time.Time{}, __obf_8882f6cf6e378ded
	}

	// NodeJS seems to use extID 13.
	if __obf_927f1edb79ba3be1 != __obf_2b109b947f8a4e95 && __obf_927f1edb79ba3be1 != 13 {
		return time.Time{}, fmt.Errorf("msgpack: invalid time ext id=%d", __obf_927f1edb79ba3be1)
	}
	__obf_75953c9a45f5db9e, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_7fa5d5dcdfa5e1ed(__obf_7e99dd6aa706c57b)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_75953c9a45f5db9e, __obf_8882f6cf6e378ded
	}

	if __obf_75953c9a45f5db9e.IsZero() {
		// Zero time does not have timezone information.
		return __obf_75953c9a45f5db9e.UTC(), nil
	}
	return __obf_75953c9a45f5db9e, nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_7fa5d5dcdfa5e1ed(__obf_7e99dd6aa706c57b int) (time.Time, error) {
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(__obf_7e99dd6aa706c57b)
	if __obf_8882f6cf6e378ded != nil {
		return time.Time{}, __obf_8882f6cf6e378ded
	}

	switch len(__obf_11bcc66cde095c11) {
	case 4:
		__obf_ad3d72f686eb35a4 := binary.BigEndian.Uint32(__obf_11bcc66cde095c11)
		return time.Unix(int64(__obf_ad3d72f686eb35a4), 0), nil
	case 8:
		__obf_ad3d72f686eb35a4 := binary.BigEndian.Uint64(__obf_11bcc66cde095c11)
		__obf_6617171abe4523cd := int64(__obf_ad3d72f686eb35a4 >> 34)
		__obf_ad3d72f686eb35a4 &= 0x00000003ffffffff
		return time.Unix(int64(__obf_ad3d72f686eb35a4), __obf_6617171abe4523cd), nil
	case 12:
		__obf_6617171abe4523cd := binary.BigEndian.Uint32(__obf_11bcc66cde095c11)
		__obf_ad3d72f686eb35a4 := binary.BigEndian.Uint64(__obf_11bcc66cde095c11[4:])
		return time.Unix(int64(__obf_ad3d72f686eb35a4), int64(__obf_6617171abe4523cd)), nil
	default:
		__obf_8882f6cf6e378ded = fmt.Errorf("msgpack: invalid ext len=%d decoding time", __obf_7e99dd6aa706c57b)
		return time.Time{}, __obf_8882f6cf6e378ded
	}
}
