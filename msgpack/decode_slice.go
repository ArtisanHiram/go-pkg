package __obf_3edaa49e853afa16

import (
	"fmt"
	"reflect"
)

var __obf_9c401b40307a916c = reflect.TypeOf((*[]string)(nil))

// DecodeArrayLen decodes array length. Length is -1 when array is nil.
func (__obf_015afbee33862a01 *Decoder) DecodeArrayLen() (int, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_5102e56f561052c2(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) __obf_5102e56f561052c2(__obf_145c7a7d25eea2bd byte) (int, error) {
	if __obf_145c7a7d25eea2bd == Nil {
		return -1, nil
	} else if __obf_145c7a7d25eea2bd >= FixedArrayLow && __obf_145c7a7d25eea2bd <= FixedArrayHigh {
		return int(__obf_145c7a7d25eea2bd & FixedArrayMask), nil
	}
	switch __obf_145c7a7d25eea2bd {
	case Array16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Array32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding array length", __obf_145c7a7d25eea2bd)
}

func __obf_22486c0586fcfbc8(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_8f0c71619c0382f1 := __obf_85d270343a0dfe14.Addr().Convert(__obf_9c401b40307a916c).Interface().(*[]string)
	return __obf_015afbee33862a01.__obf_a75b379e5bfbda38(__obf_8f0c71619c0382f1)
}

func (__obf_015afbee33862a01 *Decoder) __obf_a75b379e5bfbda38(__obf_8f0c71619c0382f1 *[]string) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeArrayLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return nil
	}
	__obf_b3201977b79aabff := __obf_d8fe9309904e23bc(*__obf_8f0c71619c0382f1, __obf_56127cd370854f0b, __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_9d8a15eb3a073456 != 0)
	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_56127cd370854f0b; __obf_bd2da3a1d4616916++ {
		__obf_6827ff1b59d7ccec, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeString()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		__obf_b3201977b79aabff = append(__obf_b3201977b79aabff, __obf_6827ff1b59d7ccec)
	}
	*__obf_8f0c71619c0382f1 = __obf_b3201977b79aabff

	return nil
}

func __obf_d8fe9309904e23bc(__obf_6827ff1b59d7ccec []string, __obf_56127cd370854f0b int, __obf_f2f3e379039efbd6 bool) []string {
	if !__obf_f2f3e379039efbd6 && __obf_56127cd370854f0b > __obf_dc3dfe27c3ffba99 {
		__obf_56127cd370854f0b = __obf_dc3dfe27c3ffba99
	}

	if __obf_6827ff1b59d7ccec == nil {
		return make([]string, 0, __obf_56127cd370854f0b)
	}

	if cap(__obf_6827ff1b59d7ccec) >= __obf_56127cd370854f0b {
		return __obf_6827ff1b59d7ccec[:0]
	}
	__obf_6827ff1b59d7ccec = __obf_6827ff1b59d7ccec[:cap(__obf_6827ff1b59d7ccec)]
	__obf_6827ff1b59d7ccec = append(__obf_6827ff1b59d7ccec, make([]string, __obf_56127cd370854f0b-len(__obf_6827ff1b59d7ccec))...)
	return __obf_6827ff1b59d7ccec[:0]
}

func __obf_84291db4b7f5231d(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeArrayLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_56127cd370854f0b == -1 {
		__obf_85d270343a0dfe14.
			Set(reflect.Zero(__obf_85d270343a0dfe14.Type()))
		return nil
	}
	if __obf_56127cd370854f0b == 0 && __obf_85d270343a0dfe14.IsNil() {
		__obf_85d270343a0dfe14.
			Set(reflect.MakeSlice(__obf_85d270343a0dfe14.Type(), 0, 0))
		return nil
	}

	if __obf_85d270343a0dfe14.Cap() >= __obf_56127cd370854f0b {
		__obf_85d270343a0dfe14.
			Set(__obf_85d270343a0dfe14.Slice(0, __obf_56127cd370854f0b))
	} else if __obf_85d270343a0dfe14.Len() < __obf_85d270343a0dfe14.Cap() {
		__obf_85d270343a0dfe14.
			Set(__obf_85d270343a0dfe14.Slice(0, __obf_85d270343a0dfe14.Cap()))
	}
	__obf_f2f3e379039efbd6 := __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_9d8a15eb3a073456 != 1

	if __obf_f2f3e379039efbd6 && __obf_56127cd370854f0b > __obf_85d270343a0dfe14.Len() {
		__obf_85d270343a0dfe14.
			Set(__obf_6cff13575fbe2fde(__obf_85d270343a0dfe14, __obf_56127cd370854f0b, __obf_f2f3e379039efbd6))
	}

	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_56127cd370854f0b; __obf_bd2da3a1d4616916++ {
		if !__obf_f2f3e379039efbd6 && __obf_bd2da3a1d4616916 >= __obf_85d270343a0dfe14.Len() {
			__obf_85d270343a0dfe14.
				Set(__obf_6cff13575fbe2fde(__obf_85d270343a0dfe14, __obf_56127cd370854f0b, __obf_f2f3e379039efbd6))
		}
		__obf_a3aef353a1d07512 := __obf_85d270343a0dfe14.Index(__obf_bd2da3a1d4616916)
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeValue(__obf_a3aef353a1d07512); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func __obf_6cff13575fbe2fde(__obf_85d270343a0dfe14 reflect.Value, __obf_56127cd370854f0b int, __obf_f2f3e379039efbd6 bool) reflect.Value {
	__obf_e6ce44d146fbd3ed := __obf_56127cd370854f0b - __obf_85d270343a0dfe14.Len()
	if !__obf_f2f3e379039efbd6 && __obf_e6ce44d146fbd3ed > __obf_dc3dfe27c3ffba99 {
		__obf_e6ce44d146fbd3ed = __obf_dc3dfe27c3ffba99
	}
	__obf_85d270343a0dfe14 = reflect.AppendSlice(__obf_85d270343a0dfe14, reflect.MakeSlice(__obf_85d270343a0dfe14.Type(), __obf_e6ce44d146fbd3ed, __obf_e6ce44d146fbd3ed))
	return __obf_85d270343a0dfe14
}

func __obf_2858ebecafe0e9d9(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeArrayLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_56127cd370854f0b == -1 {
		return nil
	}
	if __obf_56127cd370854f0b > __obf_85d270343a0dfe14.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_85d270343a0dfe14.Type(), __obf_85d270343a0dfe14.Len(), __obf_56127cd370854f0b)
	}

	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_56127cd370854f0b; __obf_bd2da3a1d4616916++ {
		__obf_ca31aa4b94e97fb4 := __obf_85d270343a0dfe14.Index(__obf_bd2da3a1d4616916)
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeValue(__obf_ca31aa4b94e97fb4); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_015afbee33862a01 *Decoder) DecodeSlice() ([]any, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_977775eeaf115ac9(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) __obf_977775eeaf115ac9(__obf_145c7a7d25eea2bd byte) ([]any, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_5102e56f561052c2(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return nil, nil
	}
	__obf_6827ff1b59d7ccec := make([]any, 0, __obf_56127cd370854f0b)
	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_56127cd370854f0b; __obf_bd2da3a1d4616916++ {
		__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		__obf_6827ff1b59d7ccec = append(__obf_6827ff1b59d7ccec, __obf_85d270343a0dfe14)
	}

	return __obf_6827ff1b59d7ccec, nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_39e3b2fed086f4e4(__obf_145c7a7d25eea2bd byte) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_5102e56f561052c2(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_56127cd370854f0b; __obf_bd2da3a1d4616916++ {
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}
