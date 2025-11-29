package __obf_a4aad98aaf3178f4

import (
	"fmt"
	"reflect"
)

func (__obf_613397fefdec0ed0 *Decoder) __obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037 byte) (int, error) {
	if __obf_6dbe86b3d9d9b037 == Nil {
		return -1, nil
	}

	if IsFixedString(__obf_6dbe86b3d9d9b037) {
		return int(__obf_6dbe86b3d9d9b037 & FixedStrMask), nil
	}

	switch __obf_6dbe86b3d9d9b037 {
	case Str8, Bin8:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Str16, Bin16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Str32, Bin32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	}

	return 0, fmt.Errorf("msgpack: invalid code=%x decoding string/bytes length", __obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeString() (string, error) {
	if __obf_0a0fa8c9dd348fed := __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_ca379a8882c3af10 != 0; __obf_0a0fa8c9dd348fed || len(__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c) > 0 {
		return __obf_613397fefdec0ed0.__obf_2fcfd6a9223bea0f(__obf_0a0fa8c9dd348fed)
	}
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.string(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) string(__obf_6dbe86b3d9d9b037 byte) (string, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_d6534b35621d5ef7(__obf_99a74e41c9c640c0)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_d6534b35621d5ef7(__obf_99a74e41c9c640c0 int) (string, error) {
	if __obf_99a74e41c9c640c0 <= 0 {
		return "", nil
	}
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(__obf_99a74e41c9c640c0)
	return string(__obf_f57209cfda8e17cf), __obf_4061ca0039150c39
}

func __obf_8b084829819fd8e7(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeString()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetString(__obf_759f458bfa19abba)
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_dc853c0361533caa() (string, error) {
	if __obf_0a0fa8c9dd348fed := __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_ca379a8882c3af10 != 0; __obf_0a0fa8c9dd348fed || len(__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c) > 0 {
		return __obf_613397fefdec0ed0.__obf_2fcfd6a9223bea0f(__obf_0a0fa8c9dd348fed)
	}
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return "", nil
	}
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(__obf_99a74e41c9c640c0)
	if __obf_4061ca0039150c39 != nil {
		return "", __obf_4061ca0039150c39
	}

	return __obf_d49c57ba49c6cfdc(__obf_f57209cfda8e17cf), nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_16359546865c48fc(__obf_cf3077802722ecc2 *[]byte) error {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_df2e52dc2ade8bd8(__obf_6dbe86b3d9d9b037, __obf_cf3077802722ecc2)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_df2e52dc2ade8bd8(__obf_6dbe86b3d9d9b037 byte, __obf_cf3077802722ecc2 *[]byte) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		*__obf_cf3077802722ecc2 = nil
		return nil
	}

	*__obf_cf3077802722ecc2, __obf_4061ca0039150c39 = __obf_4429146a43365de4(__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e, *__obf_cf3077802722ecc2, __obf_99a74e41c9c640c0)
	return __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_07d3af657b68439d(__obf_6dbe86b3d9d9b037 byte) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 <= 0 {
		return nil
	}
	return __obf_613397fefdec0ed0.__obf_2bfec4fb91ed09f6(__obf_99a74e41c9c640c0)
}

func __obf_5091976ad4bd5009(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return nil
	}
	if __obf_99a74e41c9c640c0 > __obf_a1f43267eeb48a1a.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_a1f43267eeb48a1a.Type(), __obf_a1f43267eeb48a1a.Len(), __obf_99a74e41c9c640c0)
	}
	__obf_f57209cfda8e17cf := __obf_a1f43267eeb48a1a.Slice(0, __obf_99a74e41c9c640c0).Bytes()
	return __obf_613397fefdec0ed0.__obf_6814a41007039154(__obf_f57209cfda8e17cf)
}
