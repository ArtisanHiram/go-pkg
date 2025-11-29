package __obf_3edaa49e853afa16

import (
	"fmt"
	"math"
	"reflect"
)

type __obf_4e4d1cbf2e9fdd58 struct {
	Type    reflect.Type
	Decoder func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value, __obf_69375c248ca4aa99 int) error
}

var __obf_733cc19b5e047a7a = make(map[int8]*__obf_4e4d1cbf2e9fdd58)

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

func RegisterExt(__obf_82dbf2151b1070df int8, __obf_e6992bb5202a647c MarshalerUnmarshaler) {
	RegisterExtEncoder(__obf_82dbf2151b1070df, __obf_e6992bb5202a647c, func(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) ([]byte, error) {
		__obf_827aeedcead80d9b := __obf_85d270343a0dfe14.Interface().(Marshaler)
		return __obf_827aeedcead80d9b.MarshalMsgpack()
	})
	RegisterExtDecoder(__obf_82dbf2151b1070df, __obf_e6992bb5202a647c, func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value, __obf_69375c248ca4aa99 int) error {
		__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_2c3d4aa0a72cea84(__obf_69375c248ca4aa99)
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		return __obf_85d270343a0dfe14.Interface().(Unmarshaler).UnmarshalMsgpack(__obf_9b4a5a04bdad2347)
	})
}

func UnregisterExt(__obf_82dbf2151b1070df int8) {
	__obf_3a8d0a5e27d4d71c(__obf_82dbf2151b1070df)
	__obf_1ece0de2090495f0(__obf_82dbf2151b1070df)
}

func RegisterExtEncoder(__obf_82dbf2151b1070df int8, __obf_e6992bb5202a647c any, __obf_88d8471374f73dd1 func(__obf_2fa4ef37f70ce1de *Encoder, __obf_85d270343a0dfe14 reflect.Value) ([]byte, error),
) {
	__obf_3a8d0a5e27d4d71c(__obf_82dbf2151b1070df)
	__obf_af250e9784b6a92c := reflect.TypeOf(__obf_e6992bb5202a647c)
	__obf_3d777c53ca4f67ef := __obf_c949c1b381cc29b3(__obf_82dbf2151b1070df, __obf_af250e9784b6a92c, __obf_88d8471374f73dd1)
	__obf_126a303617e63b80.
		Store(__obf_82dbf2151b1070df, __obf_af250e9784b6a92c)
	__obf_126a303617e63b80.
		Store(__obf_af250e9784b6a92c, __obf_3d777c53ca4f67ef)
	if __obf_af250e9784b6a92c.Kind() == reflect.Ptr {
		__obf_126a303617e63b80.
			Store(__obf_af250e9784b6a92c.Elem(), __obf_58a6c22f2b301fc4(__obf_3d777c53ca4f67ef))
	}
}

func __obf_3a8d0a5e27d4d71c(__obf_82dbf2151b1070df int8) {
	__obf_862c8bdd067ceabe, __obf_ccfb92cc26e4569f := __obf_126a303617e63b80.Load(__obf_82dbf2151b1070df)
	if !__obf_ccfb92cc26e4569f {
		return
	}
	__obf_126a303617e63b80.
		Delete(__obf_82dbf2151b1070df)
	__obf_af250e9784b6a92c := __obf_862c8bdd067ceabe.(reflect.Type)
	__obf_126a303617e63b80.
		Delete(__obf_af250e9784b6a92c)
	if __obf_af250e9784b6a92c.Kind() == reflect.Ptr {
		__obf_126a303617e63b80.
			Delete(__obf_af250e9784b6a92c.Elem())
	}
}

func __obf_c949c1b381cc29b3(__obf_82dbf2151b1070df int8, __obf_af250e9784b6a92c reflect.Type, __obf_88d8471374f73dd1 func(__obf_2fa4ef37f70ce1de *Encoder, __obf_85d270343a0dfe14 reflect.Value) ([]byte, error),
) __obf_b2b0b98cf9f599e0 {
	__obf_6238d8f3c0669e3b := __obf_af250e9784b6a92c.Kind() == reflect.Ptr

	return func(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if __obf_6238d8f3c0669e3b && __obf_85d270343a0dfe14.IsNil() {
			return __obf_84d0d31a8288f191.EncodeNil()
		}
		__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_88d8471374f73dd1(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14)
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}

		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.EncodeExtHeader(__obf_82dbf2151b1070df, len(__obf_9b4a5a04bdad2347)); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}

		return __obf_84d0d31a8288f191.__obf_fe140b946d6a369e(__obf_9b4a5a04bdad2347)
	}
}

func __obf_58a6c22f2b301fc4(__obf_3d777c53ca4f67ef __obf_b2b0b98cf9f599e0) __obf_b2b0b98cf9f599e0 {
	return func(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if !__obf_85d270343a0dfe14.CanAddr() {
			return fmt.Errorf("msgpack: EncodeExt(nonaddressable %T)", __obf_85d270343a0dfe14.Interface())
		}
		return __obf_3d777c53ca4f67ef(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14.Addr())
	}
}

func RegisterExtDecoder(__obf_82dbf2151b1070df int8, __obf_e6992bb5202a647c any, __obf_2cb4408eb815d713 func(__obf_20b87cf22b03338d *Decoder, __obf_85d270343a0dfe14 reflect.Value, __obf_69375c248ca4aa99 int) error,
) {
	__obf_1ece0de2090495f0(__obf_82dbf2151b1070df)
	__obf_af250e9784b6a92c := reflect.TypeOf(__obf_e6992bb5202a647c)
	__obf_32412c616faf44f9 := __obf_9eccfb09103f978b(__obf_82dbf2151b1070df, __obf_af250e9784b6a92c, __obf_2cb4408eb815d713)
	__obf_733cc19b5e047a7a[__obf_82dbf2151b1070df] = &__obf_4e4d1cbf2e9fdd58{
		Type:    __obf_af250e9784b6a92c,
		Decoder: __obf_2cb4408eb815d713,
	}
	__obf_8d7a14362d6f6fc4.
		Store(__obf_82dbf2151b1070df, __obf_af250e9784b6a92c)
	__obf_8d7a14362d6f6fc4.
		Store(__obf_af250e9784b6a92c, __obf_32412c616faf44f9)
	if __obf_af250e9784b6a92c.Kind() == reflect.Ptr {
		__obf_8d7a14362d6f6fc4.
			Store(__obf_af250e9784b6a92c.Elem(), __obf_0cc24b595ff75be1(__obf_32412c616faf44f9))
	}
}

func __obf_1ece0de2090495f0(__obf_82dbf2151b1070df int8) {
	__obf_862c8bdd067ceabe, __obf_ccfb92cc26e4569f := __obf_8d7a14362d6f6fc4.Load(__obf_82dbf2151b1070df)
	if !__obf_ccfb92cc26e4569f {
		return
	}
	__obf_8d7a14362d6f6fc4.
		Delete(__obf_82dbf2151b1070df)
	delete(__obf_733cc19b5e047a7a, __obf_82dbf2151b1070df)
	__obf_af250e9784b6a92c := __obf_862c8bdd067ceabe.(reflect.Type)
	__obf_8d7a14362d6f6fc4.
		Delete(__obf_af250e9784b6a92c)
	if __obf_af250e9784b6a92c.Kind() == reflect.Ptr {
		__obf_8d7a14362d6f6fc4.
			Delete(__obf_af250e9784b6a92c.Elem())
	}
}

func __obf_9eccfb09103f978b(__obf_56eb4f4104917b2c int8, __obf_af250e9784b6a92c reflect.Type, __obf_2cb4408eb815d713 func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value, __obf_69375c248ca4aa99 int) error,
) __obf_47bea90bc2ca8162 {
	return __obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
		__obf_82dbf2151b1070df, __obf_69375c248ca4aa99, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeExtHeader()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		if __obf_82dbf2151b1070df != __obf_56eb4f4104917b2c {
			return fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_82dbf2151b1070df, __obf_56eb4f4104917b2c)
		}
		return __obf_2cb4408eb815d713(__obf_015afbee33862a01, __obf_85d270343a0dfe14, __obf_69375c248ca4aa99)
	})
}

func __obf_0cc24b595ff75be1(__obf_32412c616faf44f9 __obf_47bea90bc2ca8162) __obf_47bea90bc2ca8162 {
	return func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if !__obf_85d270343a0dfe14.CanAddr() {
			return fmt.Errorf("msgpack: DecodeExt(nonaddressable %T)", __obf_85d270343a0dfe14.Interface())
		}
		return __obf_32412c616faf44f9(__obf_015afbee33862a01, __obf_85d270343a0dfe14.Addr())
	}
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeExtHeader(__obf_82dbf2151b1070df int8, __obf_69375c248ca4aa99 int) error {
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_257d74d0a2026131(__obf_69375c248ca4aa99); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.WriteByte(byte(__obf_82dbf2151b1070df)); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_257d74d0a2026131(__obf_48b3b71f73d35432 int) error {
	switch __obf_48b3b71f73d35432 {
	case 1:
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt1)
	case 2:
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt2)
	case 4:
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt4)
	case 8:
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt8)
	case 16:
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(FixExt16)
	}
	if __obf_48b3b71f73d35432 <= math.MaxUint8 {
		return __obf_84d0d31a8288f191.__obf_59b14e1d8bb8dea4(Ext8, uint8(__obf_48b3b71f73d35432))
	}
	if __obf_48b3b71f73d35432 <= math.MaxUint16 {
		return __obf_84d0d31a8288f191.__obf_941ea548b75e478e(Ext16, uint16(__obf_48b3b71f73d35432))
	}
	return __obf_84d0d31a8288f191.__obf_7636556325987942(Ext32, uint32(__obf_48b3b71f73d35432))
}

func (__obf_015afbee33862a01 *Decoder) DecodeExtHeader() (__obf_82dbf2151b1070df int8, __obf_69375c248ca4aa99 int, __obf_3aa78ad28f50ed46 error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return
	}
	return __obf_015afbee33862a01.__obf_9d268a1ca1cebd45(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) __obf_9d268a1ca1cebd45(__obf_145c7a7d25eea2bd byte) (int8, int, error) {
	__obf_69375c248ca4aa99, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_acc2f24f73a5a428(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, 0, __obf_3aa78ad28f50ed46
	}
	__obf_82dbf2151b1070df, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, 0, __obf_3aa78ad28f50ed46
	}

	return int8(__obf_82dbf2151b1070df), __obf_69375c248ca4aa99, nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_acc2f24f73a5a428(__obf_145c7a7d25eea2bd byte) (int, error) {
	switch __obf_145c7a7d25eea2bd {
	case FixExt1:
		return 1, nil
	case FixExt2:
		return 2, nil
	case FixExt4:
		return 4, nil
	case FixExt8:
		return 8, nil
	case FixExt16:
		return 16, nil
	case Ext8:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint8()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Ext16:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint16()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	case Ext32:
		__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.uint32()
		return int(__obf_56127cd370854f0b), __obf_3aa78ad28f50ed46
	default:
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", __obf_145c7a7d25eea2bd)
	}
}

func (__obf_015afbee33862a01 *Decoder) __obf_8795dc4d75fc75c5(__obf_145c7a7d25eea2bd byte) (any, error) {
	__obf_82dbf2151b1070df, __obf_69375c248ca4aa99, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_9d268a1ca1cebd45(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	__obf_1302a1b4627a1eb2, __obf_ccfb92cc26e4569f := __obf_733cc19b5e047a7a[__obf_82dbf2151b1070df]
	if !__obf_ccfb92cc26e4569f {
		return nil, fmt.Errorf("msgpack: unknown ext id=%d", __obf_82dbf2151b1070df)
	}
	__obf_85d270343a0dfe14 := __obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_1302a1b4627a1eb2.Type).Elem()
	if __obf_6238d8f3c0669e3b(__obf_85d270343a0dfe14.Kind()) && __obf_85d270343a0dfe14.IsNil() {
		__obf_85d270343a0dfe14.
			Set(__obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_1302a1b4627a1eb2.Type.Elem()))
	}

	if __obf_3aa78ad28f50ed46 := __obf_1302a1b4627a1eb2.Decoder(__obf_015afbee33862a01, __obf_85d270343a0dfe14, __obf_69375c248ca4aa99); __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}

	return __obf_85d270343a0dfe14.Interface(), nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_31f60c6085a2cba9(__obf_145c7a7d25eea2bd byte) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_acc2f24f73a5a428(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_1a431295b897a1b7(__obf_56127cd370854f0b + 1)
}

func (__obf_015afbee33862a01 *Decoder) __obf_c7f8d446fe3a2767(__obf_145c7a7d25eea2bd byte) error {
	// Read ext type.
	_, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	// Read ext body len.
	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_4cfdbcbdc23dc75a(__obf_145c7a7d25eea2bd); __obf_bd2da3a1d4616916++ {
		_, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}

func __obf_4cfdbcbdc23dc75a(__obf_145c7a7d25eea2bd byte) int {
	switch __obf_145c7a7d25eea2bd {
	case Ext8:
		return 1
	case Ext16:
		return 2
	case Ext32:
		return 4
	}
	return 0
}
