package __obf_a4aad98aaf3178f4

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"time"
)

var __obf_dc001522b79280df int8 = -1

func init() {
	RegisterExtEncoder(__obf_dc001522b79280df, time.Time{}, __obf_7b2ab8f2da3514e9)
	RegisterExtDecoder(__obf_dc001522b79280df, time.Time{}, __obf_d5bc030564efbf76)
}

func __obf_7b2ab8f2da3514e9(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) ([]byte, error) {
	return __obf_2c8e97779108ab17.__obf_f9a0e9cad49dbfad(__obf_a1f43267eeb48a1a.Interface().(time.Time)), nil
}

func __obf_d5bc030564efbf76(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value, __obf_f975a2b18cb9c21c int) error {
	__obf_a83e1c062c542c07, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ce8d8fe7f8f58988(__obf_f975a2b18cb9c21c)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	if __obf_a83e1c062c542c07.IsZero() {
		__obf_a83e1c062c542c07 = // Zero time does not have timezone information.
			__obf_a83e1c062c542c07.UTC()
	}
	__obf_cf3077802722ecc2 := __obf_a1f43267eeb48a1a.Addr().Interface().(*time.Time)
	*__obf_cf3077802722ecc2 = __obf_a83e1c062c542c07

	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeTime(__obf_a83e1c062c542c07 time.Time) error {
	__obf_f57209cfda8e17cf := __obf_2c8e97779108ab17.__obf_f9a0e9cad49dbfad(__obf_a83e1c062c542c07)
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_d1f2c919354a65de(len(__obf_f57209cfda8e17cf)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.WriteByte(byte(__obf_dc001522b79280df)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_f57209cfda8e17cf)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_f9a0e9cad49dbfad(__obf_a83e1c062c542c07 time.Time) []byte {
	if __obf_2c8e97779108ab17.__obf_ccee18a737f071fd == nil {
		__obf_2c8e97779108ab17.__obf_ccee18a737f071fd = make([]byte, 12)
	}
	__obf_58e031e2e2ceaf68 := uint64(__obf_a83e1c062c542c07.Unix())
	if __obf_58e031e2e2ceaf68>>34 == 0 {
		__obf_69e4c283179273ce := uint64(__obf_a83e1c062c542c07.Nanosecond())<<34 | __obf_58e031e2e2ceaf68

		if __obf_69e4c283179273ce&0xffffffff00000000 == 0 {
			__obf_f57209cfda8e17cf := __obf_2c8e97779108ab17.__obf_ccee18a737f071fd[:4]
			binary.BigEndian.PutUint32(__obf_f57209cfda8e17cf, uint32(__obf_69e4c283179273ce))
			return __obf_f57209cfda8e17cf
		}
		__obf_f57209cfda8e17cf := __obf_2c8e97779108ab17.__obf_ccee18a737f071fd[:8]
		binary.BigEndian.PutUint64(__obf_f57209cfda8e17cf, __obf_69e4c283179273ce)
		return __obf_f57209cfda8e17cf
	}
	__obf_f57209cfda8e17cf := __obf_2c8e97779108ab17.__obf_ccee18a737f071fd[:12]
	binary.BigEndian.PutUint32(__obf_f57209cfda8e17cf, uint32(__obf_a83e1c062c542c07.Nanosecond()))
	binary.BigEndian.PutUint64(__obf_f57209cfda8e17cf[4:], __obf_58e031e2e2ceaf68)
	return __obf_f57209cfda8e17cf
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeTime() (time.Time, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return time.Time{}, __obf_4061ca0039150c39
	}

	// Legacy format.
	if __obf_6dbe86b3d9d9b037 == FixedArrayLow|2 {
		__obf_dc08a27274a2538d, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
		if __obf_4061ca0039150c39 != nil {
			return time.Time{}, __obf_4061ca0039150c39
		}
		__obf_7a70aec96d8819a7, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
		if __obf_4061ca0039150c39 != nil {
			return time.Time{}, __obf_4061ca0039150c39
		}

		return time.Unix(__obf_dc08a27274a2538d, __obf_7a70aec96d8819a7), nil
	}

	if IsString(__obf_6dbe86b3d9d9b037) {
		__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.string(__obf_6dbe86b3d9d9b037)
		if __obf_4061ca0039150c39 != nil {
			return time.Time{}, __obf_4061ca0039150c39
		}
		return time.Parse(time.RFC3339Nano, __obf_759f458bfa19abba)
	}
	__obf_2f505346eb973ca1, __obf_f975a2b18cb9c21c, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_2227ad80d0435996(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return time.Time{}, __obf_4061ca0039150c39
	}

	// NodeJS seems to use extID 13.
	if __obf_2f505346eb973ca1 != __obf_dc001522b79280df && __obf_2f505346eb973ca1 != 13 {
		return time.Time{}, fmt.Errorf("msgpack: invalid time ext id=%d", __obf_2f505346eb973ca1)
	}
	__obf_a83e1c062c542c07, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ce8d8fe7f8f58988(__obf_f975a2b18cb9c21c)
	if __obf_4061ca0039150c39 != nil {
		return __obf_a83e1c062c542c07, __obf_4061ca0039150c39
	}

	if __obf_a83e1c062c542c07.IsZero() {
		// Zero time does not have timezone information.
		return __obf_a83e1c062c542c07.UTC(), nil
	}
	return __obf_a83e1c062c542c07, nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_ce8d8fe7f8f58988(__obf_f975a2b18cb9c21c int) (time.Time, error) {
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(__obf_f975a2b18cb9c21c)
	if __obf_4061ca0039150c39 != nil {
		return time.Time{}, __obf_4061ca0039150c39
	}

	switch len(__obf_f57209cfda8e17cf) {
	case 4:
		__obf_dc08a27274a2538d := binary.BigEndian.Uint32(__obf_f57209cfda8e17cf)
		return time.Unix(int64(__obf_dc08a27274a2538d), 0), nil
	case 8:
		__obf_dc08a27274a2538d := binary.BigEndian.Uint64(__obf_f57209cfda8e17cf)
		__obf_7a70aec96d8819a7 := int64(__obf_dc08a27274a2538d >> 34)
		__obf_dc08a27274a2538d &= 0x00000003ffffffff
		return time.Unix(int64(__obf_dc08a27274a2538d), __obf_7a70aec96d8819a7), nil
	case 12:
		__obf_7a70aec96d8819a7 := binary.BigEndian.Uint32(__obf_f57209cfda8e17cf)
		__obf_dc08a27274a2538d := binary.BigEndian.Uint64(__obf_f57209cfda8e17cf[4:])
		return time.Unix(int64(__obf_dc08a27274a2538d), int64(__obf_7a70aec96d8819a7)), nil
	default:
		__obf_4061ca0039150c39 = fmt.Errorf("msgpack: invalid ext len=%d decoding time", __obf_f975a2b18cb9c21c)
		return time.Time{}, __obf_4061ca0039150c39
	}
}
