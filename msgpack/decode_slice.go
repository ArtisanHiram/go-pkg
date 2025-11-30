package __obf_3e0c303610a19bd4

import (
	"fmt"
	"reflect"
)

var __obf_ec9fe546e17363b1 = reflect.TypeOf((*[]string)(nil))

// DecodeArrayLen decodes array length. Length is -1 when array is nil.
func (__obf_dc35117108ba8439 *Decoder) DecodeArrayLen() (int, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_2a74eb3d00d3189b(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_2a74eb3d00d3189b(__obf_e46289218af709bf byte) (int, error) {
	if __obf_e46289218af709bf == Nil {
		return -1, nil
	} else if __obf_e46289218af709bf >= FixedArrayLow && __obf_e46289218af709bf <= FixedArrayHigh {
		return int(__obf_e46289218af709bf & FixedArrayMask), nil
	}
	switch __obf_e46289218af709bf {
	case Array16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Array32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding array length", __obf_e46289218af709bf)
}

func __obf_f8b216cbef75a778(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_b5a4664807537c0d := __obf_63bbcee86d44fdde.Addr().Convert(__obf_ec9fe546e17363b1).Interface().(*[]string)
	return __obf_dc35117108ba8439.__obf_bdc4d45b04bc0799(__obf_b5a4664807537c0d)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_bdc4d45b04bc0799(__obf_b5a4664807537c0d *[]string) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeArrayLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return nil
	}
	__obf_66e8d01dfb8a3535 := __obf_3deb9518d63d9be0(*__obf_b5a4664807537c0d, __obf_4909ae60ffbb8e53, __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_4b9eaa01d630fce4 != 0)
	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_4909ae60ffbb8e53; __obf_39714879601f9b69++ {
		__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeString()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		__obf_66e8d01dfb8a3535 = append(__obf_66e8d01dfb8a3535, __obf_61027e0491b6dd3d)
	}
	*__obf_b5a4664807537c0d = __obf_66e8d01dfb8a3535

	return nil
}

func __obf_3deb9518d63d9be0(__obf_61027e0491b6dd3d []string, __obf_4909ae60ffbb8e53 int, __obf_2f0ca841c331a185 bool) []string {
	if !__obf_2f0ca841c331a185 && __obf_4909ae60ffbb8e53 > __obf_180374296ff44439 {
		__obf_4909ae60ffbb8e53 = __obf_180374296ff44439
	}

	if __obf_61027e0491b6dd3d == nil {
		return make([]string, 0, __obf_4909ae60ffbb8e53)
	}

	if cap(__obf_61027e0491b6dd3d) >= __obf_4909ae60ffbb8e53 {
		return __obf_61027e0491b6dd3d[:0]
	}
	__obf_61027e0491b6dd3d = __obf_61027e0491b6dd3d[:cap(__obf_61027e0491b6dd3d)]
	__obf_61027e0491b6dd3d = append(__obf_61027e0491b6dd3d, make([]string, __obf_4909ae60ffbb8e53-len(__obf_61027e0491b6dd3d))...)
	return __obf_61027e0491b6dd3d[:0]
}

func __obf_8047bcf5dd14321e(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeArrayLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	if __obf_4909ae60ffbb8e53 == -1 {
		__obf_63bbcee86d44fdde.
			Set(reflect.Zero(__obf_63bbcee86d44fdde.Type()))
		return nil
	}
	if __obf_4909ae60ffbb8e53 == 0 && __obf_63bbcee86d44fdde.IsNil() {
		__obf_63bbcee86d44fdde.
			Set(reflect.MakeSlice(__obf_63bbcee86d44fdde.Type(), 0, 0))
		return nil
	}

	if __obf_63bbcee86d44fdde.Cap() >= __obf_4909ae60ffbb8e53 {
		__obf_63bbcee86d44fdde.
			Set(__obf_63bbcee86d44fdde.Slice(0, __obf_4909ae60ffbb8e53))
	} else if __obf_63bbcee86d44fdde.Len() < __obf_63bbcee86d44fdde.Cap() {
		__obf_63bbcee86d44fdde.
			Set(__obf_63bbcee86d44fdde.Slice(0, __obf_63bbcee86d44fdde.Cap()))
	}
	__obf_2f0ca841c331a185 := __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_4b9eaa01d630fce4 != 1

	if __obf_2f0ca841c331a185 && __obf_4909ae60ffbb8e53 > __obf_63bbcee86d44fdde.Len() {
		__obf_63bbcee86d44fdde.
			Set(__obf_14bb6fedb0944576(__obf_63bbcee86d44fdde, __obf_4909ae60ffbb8e53, __obf_2f0ca841c331a185))
	}

	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_4909ae60ffbb8e53; __obf_39714879601f9b69++ {
		if !__obf_2f0ca841c331a185 && __obf_39714879601f9b69 >= __obf_63bbcee86d44fdde.Len() {
			__obf_63bbcee86d44fdde.
				Set(__obf_14bb6fedb0944576(__obf_63bbcee86d44fdde, __obf_4909ae60ffbb8e53, __obf_2f0ca841c331a185))
		}
		__obf_265b8ddf380ff77a := __obf_63bbcee86d44fdde.Index(__obf_39714879601f9b69)
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeValue(__obf_265b8ddf380ff77a); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func __obf_14bb6fedb0944576(__obf_63bbcee86d44fdde reflect.Value, __obf_4909ae60ffbb8e53 int, __obf_2f0ca841c331a185 bool) reflect.Value {
	__obf_ad8e81dbc1f2bf8c := __obf_4909ae60ffbb8e53 - __obf_63bbcee86d44fdde.Len()
	if !__obf_2f0ca841c331a185 && __obf_ad8e81dbc1f2bf8c > __obf_180374296ff44439 {
		__obf_ad8e81dbc1f2bf8c = __obf_180374296ff44439
	}
	__obf_63bbcee86d44fdde = reflect.AppendSlice(__obf_63bbcee86d44fdde, reflect.MakeSlice(__obf_63bbcee86d44fdde.Type(), __obf_ad8e81dbc1f2bf8c, __obf_ad8e81dbc1f2bf8c))
	return __obf_63bbcee86d44fdde
}

func __obf_99c138aaf5d3c77b(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeArrayLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	if __obf_4909ae60ffbb8e53 == -1 {
		return nil
	}
	if __obf_4909ae60ffbb8e53 > __obf_63bbcee86d44fdde.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_63bbcee86d44fdde.Type(), __obf_63bbcee86d44fdde.Len(), __obf_4909ae60ffbb8e53)
	}

	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_4909ae60ffbb8e53; __obf_39714879601f9b69++ {
		__obf_056c80694e0598bf := __obf_63bbcee86d44fdde.Index(__obf_39714879601f9b69)
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeValue(__obf_056c80694e0598bf); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_dc35117108ba8439 *Decoder) DecodeSlice() ([]any, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_562eb479f437208e(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_562eb479f437208e(__obf_e46289218af709bf byte) ([]any, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_2a74eb3d00d3189b(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return nil, nil
	}
	__obf_61027e0491b6dd3d := make([]any, 0, __obf_4909ae60ffbb8e53)
	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_4909ae60ffbb8e53; __obf_39714879601f9b69++ {
		__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		__obf_61027e0491b6dd3d = append(__obf_61027e0491b6dd3d, __obf_63bbcee86d44fdde)
	}

	return __obf_61027e0491b6dd3d, nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_3d61f79527fa3520(__obf_e46289218af709bf byte) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_2a74eb3d00d3189b(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_4909ae60ffbb8e53; __obf_39714879601f9b69++ {
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}
