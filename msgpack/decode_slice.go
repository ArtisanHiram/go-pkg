package __obf_a4aad98aaf3178f4

import (
	"fmt"
	"reflect"
)

var __obf_d2ba50ac68c0af02 = reflect.TypeOf((*[]string)(nil))

// DecodeArrayLen decodes array length. Length is -1 when array is nil.
func (__obf_613397fefdec0ed0 *Decoder) DecodeArrayLen() (int, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_0b139571f0c923ec(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_0b139571f0c923ec(__obf_6dbe86b3d9d9b037 byte) (int, error) {
	if __obf_6dbe86b3d9d9b037 == Nil {
		return -1, nil
	} else if __obf_6dbe86b3d9d9b037 >= FixedArrayLow && __obf_6dbe86b3d9d9b037 <= FixedArrayHigh {
		return int(__obf_6dbe86b3d9d9b037 & FixedArrayMask), nil
	}
	switch __obf_6dbe86b3d9d9b037 {
	case Array16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Array32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding array length", __obf_6dbe86b3d9d9b037)
}

func __obf_89bc4f0b771b8379(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_cf3077802722ecc2 := __obf_a1f43267eeb48a1a.Addr().Convert(__obf_d2ba50ac68c0af02).Interface().(*[]string)
	return __obf_613397fefdec0ed0.__obf_3f7e7f51ea420948(__obf_cf3077802722ecc2)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_3f7e7f51ea420948(__obf_cf3077802722ecc2 *[]string) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeArrayLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return nil
	}
	__obf_a7ace2ee7c6566ec := __obf_8e1a71d747f1999e(*__obf_cf3077802722ecc2, __obf_99a74e41c9c640c0, __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_beecb71ff037aa1b != 0)
	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_99a74e41c9c640c0; __obf_097d8b6061c9592b++ {
		__obf_759f458bfa19abba, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeString()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		__obf_a7ace2ee7c6566ec = append(__obf_a7ace2ee7c6566ec, __obf_759f458bfa19abba)
	}
	*__obf_cf3077802722ecc2 = __obf_a7ace2ee7c6566ec

	return nil
}

func __obf_8e1a71d747f1999e(__obf_759f458bfa19abba []string, __obf_99a74e41c9c640c0 int, __obf_187c82cd1c42659f bool) []string {
	if !__obf_187c82cd1c42659f && __obf_99a74e41c9c640c0 > __obf_953f2fb1425dda7a {
		__obf_99a74e41c9c640c0 = __obf_953f2fb1425dda7a
	}

	if __obf_759f458bfa19abba == nil {
		return make([]string, 0, __obf_99a74e41c9c640c0)
	}

	if cap(__obf_759f458bfa19abba) >= __obf_99a74e41c9c640c0 {
		return __obf_759f458bfa19abba[:0]
	}
	__obf_759f458bfa19abba = __obf_759f458bfa19abba[:cap(__obf_759f458bfa19abba)]
	__obf_759f458bfa19abba = append(__obf_759f458bfa19abba, make([]string, __obf_99a74e41c9c640c0-len(__obf_759f458bfa19abba))...)
	return __obf_759f458bfa19abba[:0]
}

func __obf_6bcf41fd13df3eb0(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeArrayLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	if __obf_99a74e41c9c640c0 == -1 {
		__obf_a1f43267eeb48a1a.
			Set(reflect.Zero(__obf_a1f43267eeb48a1a.Type()))
		return nil
	}
	if __obf_99a74e41c9c640c0 == 0 && __obf_a1f43267eeb48a1a.IsNil() {
		__obf_a1f43267eeb48a1a.
			Set(reflect.MakeSlice(__obf_a1f43267eeb48a1a.Type(), 0, 0))
		return nil
	}

	if __obf_a1f43267eeb48a1a.Cap() >= __obf_99a74e41c9c640c0 {
		__obf_a1f43267eeb48a1a.
			Set(__obf_a1f43267eeb48a1a.Slice(0, __obf_99a74e41c9c640c0))
	} else if __obf_a1f43267eeb48a1a.Len() < __obf_a1f43267eeb48a1a.Cap() {
		__obf_a1f43267eeb48a1a.
			Set(__obf_a1f43267eeb48a1a.Slice(0, __obf_a1f43267eeb48a1a.Cap()))
	}
	__obf_187c82cd1c42659f := __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_beecb71ff037aa1b != 1

	if __obf_187c82cd1c42659f && __obf_99a74e41c9c640c0 > __obf_a1f43267eeb48a1a.Len() {
		__obf_a1f43267eeb48a1a.
			Set(__obf_72710338ef238337(__obf_a1f43267eeb48a1a, __obf_99a74e41c9c640c0, __obf_187c82cd1c42659f))
	}

	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_99a74e41c9c640c0; __obf_097d8b6061c9592b++ {
		if !__obf_187c82cd1c42659f && __obf_097d8b6061c9592b >= __obf_a1f43267eeb48a1a.Len() {
			__obf_a1f43267eeb48a1a.
				Set(__obf_72710338ef238337(__obf_a1f43267eeb48a1a, __obf_99a74e41c9c640c0, __obf_187c82cd1c42659f))
		}
		__obf_ed47b1177873a7da := __obf_a1f43267eeb48a1a.Index(__obf_097d8b6061c9592b)
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeValue(__obf_ed47b1177873a7da); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func __obf_72710338ef238337(__obf_a1f43267eeb48a1a reflect.Value, __obf_99a74e41c9c640c0 int, __obf_187c82cd1c42659f bool) reflect.Value {
	__obf_ae3db9ba55b627b3 := __obf_99a74e41c9c640c0 - __obf_a1f43267eeb48a1a.Len()
	if !__obf_187c82cd1c42659f && __obf_ae3db9ba55b627b3 > __obf_953f2fb1425dda7a {
		__obf_ae3db9ba55b627b3 = __obf_953f2fb1425dda7a
	}
	__obf_a1f43267eeb48a1a = reflect.AppendSlice(__obf_a1f43267eeb48a1a, reflect.MakeSlice(__obf_a1f43267eeb48a1a.Type(), __obf_ae3db9ba55b627b3, __obf_ae3db9ba55b627b3))
	return __obf_a1f43267eeb48a1a
}

func __obf_7e0a77c3faa7dfde(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeArrayLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	if __obf_99a74e41c9c640c0 == -1 {
		return nil
	}
	if __obf_99a74e41c9c640c0 > __obf_a1f43267eeb48a1a.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_a1f43267eeb48a1a.Type(), __obf_a1f43267eeb48a1a.Len(), __obf_99a74e41c9c640c0)
	}

	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_99a74e41c9c640c0; __obf_097d8b6061c9592b++ {
		__obf_2689da659d099aa3 := __obf_a1f43267eeb48a1a.Index(__obf_097d8b6061c9592b)
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeValue(__obf_2689da659d099aa3); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeSlice() ([]any, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_1020cdfcb81f8fa6(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_1020cdfcb81f8fa6(__obf_6dbe86b3d9d9b037 byte) ([]any, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0b139571f0c923ec(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return nil, nil
	}
	__obf_759f458bfa19abba := make([]any, 0, __obf_99a74e41c9c640c0)
	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_99a74e41c9c640c0; __obf_097d8b6061c9592b++ {
		__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		__obf_759f458bfa19abba = append(__obf_759f458bfa19abba, __obf_a1f43267eeb48a1a)
	}

	return __obf_759f458bfa19abba, nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_b037cd157d532089(__obf_6dbe86b3d9d9b037 byte) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0b139571f0c923ec(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_99a74e41c9c640c0; __obf_097d8b6061c9592b++ {
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}
