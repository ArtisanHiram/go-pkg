package __obf_a935eb7f548271a4

import (
	"fmt"
	"reflect"
)

var __obf_6be1e41337b56eb4 = reflect.TypeOf((*[]string)(nil))

// DecodeArrayLen decodes array length. Length is -1 when array is nil.
func (__obf_a21885da2425f2b2 *Decoder) DecodeArrayLen() (int, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_2aa510435f9160c6(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_2aa510435f9160c6(__obf_f5df560f4d67421b byte) (int, error) {
	if __obf_f5df560f4d67421b == Nil {
		return -1, nil
	} else if __obf_f5df560f4d67421b >= FixedArrayLow && __obf_f5df560f4d67421b <= FixedArrayHigh {
		return int(__obf_f5df560f4d67421b & FixedArrayMask), nil
	}
	switch __obf_f5df560f4d67421b {
	case Array16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Array32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding array length", __obf_f5df560f4d67421b)
}

func __obf_9b16848de3b7b97b(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_0d8a994785cda6df := __obf_6d570581f4b60dbc.Addr().Convert(__obf_6be1e41337b56eb4).Interface().(*[]string)
	return __obf_a21885da2425f2b2.__obf_95cef988e387ad44(__obf_0d8a994785cda6df)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_95cef988e387ad44(__obf_0d8a994785cda6df *[]string) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeArrayLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return nil
	}
	__obf_cc04a11b8213c893 := __obf_aa26c0ea0bddc075(*__obf_0d8a994785cda6df, __obf_326af9bd942662ac, __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_20f89e0d9adcdf29 != 0)
	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_326af9bd942662ac; __obf_4a8e280ffaa52cf4++ {
		__obf_b62c60fba2fd9788, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeString()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		__obf_cc04a11b8213c893 = append(__obf_cc04a11b8213c893, __obf_b62c60fba2fd9788)
	}
	*__obf_0d8a994785cda6df = __obf_cc04a11b8213c893

	return nil
}

func __obf_aa26c0ea0bddc075(__obf_b62c60fba2fd9788 []string, __obf_326af9bd942662ac int, __obf_dca3768572a4a8f5 bool) []string {
	if !__obf_dca3768572a4a8f5 && __obf_326af9bd942662ac > __obf_69adbcf4aa0f33cb {
		__obf_326af9bd942662ac = __obf_69adbcf4aa0f33cb
	}

	if __obf_b62c60fba2fd9788 == nil {
		return make([]string, 0, __obf_326af9bd942662ac)
	}

	if cap(__obf_b62c60fba2fd9788) >= __obf_326af9bd942662ac {
		return __obf_b62c60fba2fd9788[:0]
	}
	__obf_b62c60fba2fd9788 = __obf_b62c60fba2fd9788[:cap(__obf_b62c60fba2fd9788)]
	__obf_b62c60fba2fd9788 = append(__obf_b62c60fba2fd9788, make([]string, __obf_326af9bd942662ac-len(__obf_b62c60fba2fd9788))...)
	return __obf_b62c60fba2fd9788[:0]
}

func __obf_8ba462222fa73e4f(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeArrayLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_326af9bd942662ac == -1 {
		__obf_6d570581f4b60dbc.
			Set(reflect.Zero(__obf_6d570581f4b60dbc.Type()))
		return nil
	}
	if __obf_326af9bd942662ac == 0 && __obf_6d570581f4b60dbc.IsNil() {
		__obf_6d570581f4b60dbc.
			Set(reflect.MakeSlice(__obf_6d570581f4b60dbc.Type(), 0, 0))
		return nil
	}

	if __obf_6d570581f4b60dbc.Cap() >= __obf_326af9bd942662ac {
		__obf_6d570581f4b60dbc.
			Set(__obf_6d570581f4b60dbc.Slice(0, __obf_326af9bd942662ac))
	} else if __obf_6d570581f4b60dbc.Len() < __obf_6d570581f4b60dbc.Cap() {
		__obf_6d570581f4b60dbc.
			Set(__obf_6d570581f4b60dbc.Slice(0, __obf_6d570581f4b60dbc.Cap()))
	}
	__obf_dca3768572a4a8f5 := __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_20f89e0d9adcdf29 != 1

	if __obf_dca3768572a4a8f5 && __obf_326af9bd942662ac > __obf_6d570581f4b60dbc.Len() {
		__obf_6d570581f4b60dbc.
			Set(__obf_b6000d69e81f4df2(__obf_6d570581f4b60dbc, __obf_326af9bd942662ac, __obf_dca3768572a4a8f5))
	}

	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_326af9bd942662ac; __obf_4a8e280ffaa52cf4++ {
		if !__obf_dca3768572a4a8f5 && __obf_4a8e280ffaa52cf4 >= __obf_6d570581f4b60dbc.Len() {
			__obf_6d570581f4b60dbc.
				Set(__obf_b6000d69e81f4df2(__obf_6d570581f4b60dbc, __obf_326af9bd942662ac, __obf_dca3768572a4a8f5))
		}
		__obf_f50dd3cfb2ad5435 := __obf_6d570581f4b60dbc.Index(__obf_4a8e280ffaa52cf4)
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeValue(__obf_f50dd3cfb2ad5435); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func __obf_b6000d69e81f4df2(__obf_6d570581f4b60dbc reflect.Value, __obf_326af9bd942662ac int, __obf_dca3768572a4a8f5 bool) reflect.Value {
	__obf_48ef3de3dccd447e := __obf_326af9bd942662ac - __obf_6d570581f4b60dbc.Len()
	if !__obf_dca3768572a4a8f5 && __obf_48ef3de3dccd447e > __obf_69adbcf4aa0f33cb {
		__obf_48ef3de3dccd447e = __obf_69adbcf4aa0f33cb
	}
	__obf_6d570581f4b60dbc = reflect.AppendSlice(__obf_6d570581f4b60dbc, reflect.MakeSlice(__obf_6d570581f4b60dbc.Type(), __obf_48ef3de3dccd447e, __obf_48ef3de3dccd447e))
	return __obf_6d570581f4b60dbc
}

func __obf_a51ce3b2fcc836ae(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeArrayLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_326af9bd942662ac == -1 {
		return nil
	}
	if __obf_326af9bd942662ac > __obf_6d570581f4b60dbc.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_6d570581f4b60dbc.Type(), __obf_6d570581f4b60dbc.Len(), __obf_326af9bd942662ac)
	}

	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_326af9bd942662ac; __obf_4a8e280ffaa52cf4++ {
		__obf_2e221a78b342d397 := __obf_6d570581f4b60dbc.Index(__obf_4a8e280ffaa52cf4)
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeValue(__obf_2e221a78b342d397); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeSlice() ([]any, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_9d4786f56b1c142d(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_9d4786f56b1c142d(__obf_f5df560f4d67421b byte) ([]any, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_2aa510435f9160c6(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return nil, nil
	}
	__obf_b62c60fba2fd9788 := make([]any, 0, __obf_326af9bd942662ac)
	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_326af9bd942662ac; __obf_4a8e280ffaa52cf4++ {
		__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		__obf_b62c60fba2fd9788 = append(__obf_b62c60fba2fd9788, __obf_6d570581f4b60dbc)
	}

	return __obf_b62c60fba2fd9788, nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_533ee38d6a6706df(__obf_f5df560f4d67421b byte) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_2aa510435f9160c6(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_326af9bd942662ac; __obf_4a8e280ffaa52cf4++ {
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}
