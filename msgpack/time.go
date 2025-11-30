package __obf_de86cdc8ae98b45b

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

var __obf_f3f11eddd7fd9e2b int8 = -1

func init() {
	RegisterExtEncoder(__obf_f3f11eddd7fd9e2b, time.Time{}, __obf_37ccfd32f90c0238)
	RegisterExtDecoder(__obf_f3f11eddd7fd9e2b, time.Time{}, __obf_db7958d2e1d4b9ba)
}

func __obf_37ccfd32f90c0238(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) ([]byte, error) {
	return __obf_7bae0b613da60dd3.__obf_4dc8385558b1778d(__obf_17732590722140e7.Interface().(time.Time)), nil
}

func __obf_db7958d2e1d4b9ba(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value, __obf_eef44dfac5af7de1 int) error {
	__obf_f1595279ca946a25, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_25e6dd2823658f7b(__obf_eef44dfac5af7de1)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	if __obf_f1595279ca946a25.IsZero() {
		__obf_f1595279ca946a25 = // Zero time does not have timezone information.
			__obf_f1595279ca946a25.UTC()
	}
	__obf_01f0154f1fce8e5a := __obf_17732590722140e7.Addr().Interface().(*time.Time)
	*__obf_01f0154f1fce8e5a = __obf_f1595279ca946a25

	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeTime(__obf_f1595279ca946a25 time.Time) error {
	__obf_a7fd7acf2bd4435f := __obf_7bae0b613da60dd3.__obf_4dc8385558b1778d(__obf_f1595279ca946a25)
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_0fe969e97a7f7c13(len(__obf_a7fd7acf2bd4435f)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.WriteByte(byte(__obf_f3f11eddd7fd9e2b)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_a7fd7acf2bd4435f)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_4dc8385558b1778d(__obf_f1595279ca946a25 time.Time) []byte {
	if __obf_7bae0b613da60dd3.__obf_985719a68b5833e6 == nil {
		__obf_7bae0b613da60dd3.__obf_985719a68b5833e6 = make([]byte, 12)
	}
	__obf_1ee8ccf6247d73aa := uint64(__obf_f1595279ca946a25.Unix())
	if __obf_1ee8ccf6247d73aa>>34 == 0 {
		__obf_50d002bc8b306017 := uint64(__obf_f1595279ca946a25.Nanosecond())<<34 | __obf_1ee8ccf6247d73aa

		if __obf_50d002bc8b306017&0xffffffff00000000 == 0 {
			__obf_a7fd7acf2bd4435f := __obf_7bae0b613da60dd3.__obf_985719a68b5833e6[:4]
			binary.BigEndian.PutUint32(__obf_a7fd7acf2bd4435f, uint32(__obf_50d002bc8b306017))
			return __obf_a7fd7acf2bd4435f
		}
		__obf_a7fd7acf2bd4435f := __obf_7bae0b613da60dd3.__obf_985719a68b5833e6[:8]
		binary.BigEndian.PutUint64(__obf_a7fd7acf2bd4435f, __obf_50d002bc8b306017)
		return __obf_a7fd7acf2bd4435f
	}
	__obf_a7fd7acf2bd4435f := __obf_7bae0b613da60dd3.__obf_985719a68b5833e6[:12]
	binary.BigEndian.PutUint32(__obf_a7fd7acf2bd4435f, uint32(__obf_f1595279ca946a25.Nanosecond()))
	binary.BigEndian.PutUint64(__obf_a7fd7acf2bd4435f[4:], __obf_1ee8ccf6247d73aa)
	return __obf_a7fd7acf2bd4435f
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeTime() (time.Time, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return time.Time{}, __obf_0feb3528b7b709ec
	}

	// Legacy format.
	if __obf_ecf6d2d3253a058d == FixedArrayLow|2 {
		__obf_67779f59b9b56bc2, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
		if __obf_0feb3528b7b709ec != nil {
			return time.Time{}, __obf_0feb3528b7b709ec
		}
		__obf_a114af7ffec83c3e, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
		if __obf_0feb3528b7b709ec != nil {
			return time.Time{}, __obf_0feb3528b7b709ec
		}

		return time.Unix(__obf_67779f59b9b56bc2, __obf_a114af7ffec83c3e), nil
	}

	if IsString(__obf_ecf6d2d3253a058d) {
		__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.string(__obf_ecf6d2d3253a058d)
		if __obf_0feb3528b7b709ec != nil {
			return time.Time{}, __obf_0feb3528b7b709ec
		}
		return time.Parse(time.RFC3339Nano, __obf_a93d004caac96500)
	}
	__obf_da58875072df271b, __obf_eef44dfac5af7de1, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_73f4503729af0402(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return time.Time{}, __obf_0feb3528b7b709ec
	}

	// NodeJS seems to use extID 13.
	if __obf_da58875072df271b != __obf_f3f11eddd7fd9e2b && __obf_da58875072df271b != 13 {
		return time.Time{}, fmt.Errorf("msgpack: invalid time ext id=%d", __obf_da58875072df271b)
	}
	__obf_f1595279ca946a25, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_25e6dd2823658f7b(__obf_eef44dfac5af7de1)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_f1595279ca946a25, __obf_0feb3528b7b709ec
	}

	if __obf_f1595279ca946a25.IsZero() {
		// Zero time does not have timezone information.
		return __obf_f1595279ca946a25.UTC(), nil
	}
	return __obf_f1595279ca946a25, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_25e6dd2823658f7b(__obf_eef44dfac5af7de1 int) (time.Time, error) {
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(__obf_eef44dfac5af7de1)
	if __obf_0feb3528b7b709ec != nil {
		return time.Time{}, __obf_0feb3528b7b709ec
	}

	switch len(__obf_a7fd7acf2bd4435f) {
	case 4:
		__obf_67779f59b9b56bc2 := binary.BigEndian.Uint32(__obf_a7fd7acf2bd4435f)
		return time.Unix(int64(__obf_67779f59b9b56bc2), 0), nil
	case 8:
		__obf_67779f59b9b56bc2 := binary.BigEndian.Uint64(__obf_a7fd7acf2bd4435f)
		__obf_a114af7ffec83c3e := int64(__obf_67779f59b9b56bc2 >> 34)
		__obf_67779f59b9b56bc2 &= 0x00000003ffffffff
		return time.Unix(int64(__obf_67779f59b9b56bc2), __obf_a114af7ffec83c3e), nil
	case 12:
		__obf_a114af7ffec83c3e := binary.BigEndian.Uint32(__obf_a7fd7acf2bd4435f)
		__obf_67779f59b9b56bc2 := binary.BigEndian.Uint64(__obf_a7fd7acf2bd4435f[4:])
		return time.Unix(int64(__obf_67779f59b9b56bc2), int64(__obf_a114af7ffec83c3e)), nil
	default:
		__obf_0feb3528b7b709ec = fmt.Errorf("msgpack: invalid ext len=%d decoding time", __obf_eef44dfac5af7de1)
		return time.Time{}, __obf_0feb3528b7b709ec
	}
}
