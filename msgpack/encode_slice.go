package __obf_3edaa49e853afa16

import (
	"math"
	"reflect"
)

var __obf_9c24738fb7143a73 = reflect.TypeOf(([]string)(nil))

func __obf_a46e21fb53854fca(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeString(__obf_85d270343a0dfe14.String())
}

func __obf_636a2f8dbc830828(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeBytes(__obf_85d270343a0dfe14.Bytes())
}

func __obf_ddd3546e24350c75(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeBytesLen(__obf_85d270343a0dfe14.Len()); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_85d270343a0dfe14.CanAddr() {
		__obf_9b4a5a04bdad2347 := __obf_85d270343a0dfe14.Slice(0, __obf_85d270343a0dfe14.Len()).Bytes()
		return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_9b4a5a04bdad2347)
	}
	__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75 = __obf_b1e973ae2d19737b(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75, __obf_85d270343a0dfe14.Len())
	reflect.Copy(reflect.ValueOf(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75), __obf_85d270343a0dfe14)
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_84d0d31a8288f191.__obf_8f3bed0c23610f75)
}

func __obf_b1e973ae2d19737b(__obf_9b4a5a04bdad2347 []byte, __obf_56127cd370854f0b int) []byte {
	if cap(__obf_9b4a5a04bdad2347) >= __obf_56127cd370854f0b {
		return __obf_9b4a5a04bdad2347[:__obf_56127cd370854f0b]
	}
	__obf_9b4a5a04bdad2347 = __obf_9b4a5a04bdad2347[:cap(__obf_9b4a5a04bdad2347)]
	__obf_9b4a5a04bdad2347 = append(__obf_9b4a5a04bdad2347, make([]byte, __obf_56127cd370854f0b-len(__obf_9b4a5a04bdad2347))...)
	return __obf_9b4a5a04bdad2347
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeBytesLen(__obf_48b3b71f73d35432 int) error {
	if __obf_48b3b71f73d35432 < 256 {
		return __obf_84d0d31a8288f191.__obf_59b14e1d8bb8dea4(Bin8, uint8(__obf_48b3b71f73d35432))
	}
	if __obf_48b3b71f73d35432 <= math.MaxUint16 {
		return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Bin16, uint16(__obf_48b3b71f73d35432))
	}
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Bin32, uint32(__obf_48b3b71f73d35432))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_b27752720f0fd1f4(__obf_48b3b71f73d35432 int) error {
	if __obf_48b3b71f73d35432 < 32 {
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixedStrLow | byte(__obf_48b3b71f73d35432))
	}
	if __obf_48b3b71f73d35432 < 256 {
		return __obf_84d0d31a8288f191.__obf_59b14e1d8bb8dea4(Str8, uint8(__obf_48b3b71f73d35432))
	}
	if __obf_48b3b71f73d35432 <= math.MaxUint16 {
		return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Str16, uint16(__obf_48b3b71f73d35432))
	}
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Str32, uint32(__obf_48b3b71f73d35432))
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeString(__obf_85d270343a0dfe14 string) error {
	if __obf_60ddbad9c6f41c91 := __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_48e006ba02b70b78 != 0; __obf_60ddbad9c6f41c91 || len(__obf_84d0d31a8288f191.__obf_b014355f64826d7b) > 0 {
		return __obf_84d0d31a8288f191.__obf_fd9d10636b1b890e(__obf_85d270343a0dfe14, __obf_60ddbad9c6f41c91)
	}
	return __obf_84d0d31a8288f191.__obf_7b5be37728f8302b(__obf_85d270343a0dfe14)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_7b5be37728f8302b(__obf_85d270343a0dfe14 string) error {
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_b27752720f0fd1f4(len(__obf_85d270343a0dfe14)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return __obf_84d0d31a8288f191.__obf_938c3e39051aebef(__obf_85d270343a0dfe14)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeBytes(__obf_85d270343a0dfe14 []byte) error {
	if __obf_85d270343a0dfe14 == nil {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeBytesLen(len(__obf_85d270343a0dfe14)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_85d270343a0dfe14)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeArrayLen(__obf_48b3b71f73d35432 int) error {
	if __obf_48b3b71f73d35432 < 16 {
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixedArrayLow | byte(__obf_48b3b71f73d35432))
	}
	if __obf_48b3b71f73d35432 <= math.MaxUint16 {
		return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Array16, uint16(__obf_48b3b71f73d35432))
	}
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Array32, uint32(__obf_48b3b71f73d35432))
}

func __obf_1f2df5fa72bbce12(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_b3201977b79aabff := __obf_85d270343a0dfe14.Convert(__obf_9c24738fb7143a73).Interface().([]string)
	return __obf_84d0d31a8288f191.__obf_9d2d0560199e1b94(__obf_b3201977b79aabff)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_9d2d0560199e1b94(__obf_6827ff1b59d7ccec []string) error {
	if __obf_6827ff1b59d7ccec == nil {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeArrayLen(len(__obf_6827ff1b59d7ccec)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	for _, __obf_85d270343a0dfe14 := range __obf_6827ff1b59d7ccec {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeString(__obf_85d270343a0dfe14); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}

func __obf_35a62158d6000e10(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	return __obf_e7b538c88a16cf6c(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14)
}

func __obf_e7b538c88a16cf6c(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_48b3b71f73d35432 := __obf_85d270343a0dfe14.Len()
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeArrayLen(__obf_48b3b71f73d35432); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_48b3b71f73d35432; __obf_bd2da3a1d4616916++ {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeValue(__obf_85d270343a0dfe14.Index(__obf_bd2da3a1d4616916)); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}
