package __obf_3edaa49e853afa16

import (
	"encoding"
	"fmt"
	"reflect"
)

var __obf_4846e675fa6fb7a6 []__obf_b2b0b98cf9f599e0

//nolint:gochecknoinits
func init() {
	__obf_4846e675fa6fb7a6 = []__obf_b2b0b98cf9f599e0{
		reflect.Bool:          __obf_19ddf8a0020ded94,
		reflect.Int:           __obf_2143a6537eb57ebf,
		reflect.Int8:          __obf_408f73c1836414a3,
		reflect.Int16:         __obf_ad4975ca3a4d2b79,
		reflect.Int32:         __obf_7ee966f13d02d236,
		reflect.Int64:         __obf_33e67bf72b55646f,
		reflect.Uint:          __obf_ea3307b0a2ac01ef,
		reflect.Uint8:         __obf_55469737ffbcc25a,
		reflect.Uint16:        __obf_a78815c57f0236f5,
		reflect.Uint32:        __obf_287a9d470d9b65bf,
		reflect.Uint64:        __obf_b38acace80334530,
		reflect.Float32:       __obf_c3e585f35dc28005,
		reflect.Float64:       __obf_9c643c7031b55a89,
		reflect.Complex64:     __obf_f2e27ee6992d8676,
		reflect.Complex128:    __obf_f2e27ee6992d8676,
		reflect.Array:         __obf_e7b538c88a16cf6c,
		reflect.Chan:          __obf_f2e27ee6992d8676,
		reflect.Func:          __obf_f2e27ee6992d8676,
		reflect.Interface:     __obf_85f45bdd273b38e6,
		reflect.Map:           __obf_c2aec17963f8c911,
		reflect.Ptr:           __obf_f2e27ee6992d8676,
		reflect.Slice:         __obf_35a62158d6000e10,
		reflect.String:        __obf_a46e21fb53854fca,
		reflect.Struct:        __obf_bb1447402cee3139,
		reflect.UnsafePointer: __obf_f2e27ee6992d8676,
	}
}

func __obf_1e49c68fa4194051(__obf_af250e9784b6a92c reflect.Type) __obf_b2b0b98cf9f599e0 {
	if __obf_85d270343a0dfe14, __obf_ccfb92cc26e4569f := __obf_126a303617e63b80.Load(__obf_af250e9784b6a92c); __obf_ccfb92cc26e4569f {
		return __obf_85d270343a0dfe14.(__obf_b2b0b98cf9f599e0)
	}
	__obf_c8137d8e88d5b1dd := _getEncoder(__obf_af250e9784b6a92c)
	__obf_126a303617e63b80.
		Store(__obf_af250e9784b6a92c, __obf_c8137d8e88d5b1dd)
	return __obf_c8137d8e88d5b1dd
}

func _getEncoder(__obf_af250e9784b6a92c reflect.Type) __obf_b2b0b98cf9f599e0 {
	__obf_5819e74754f40cb2 := __obf_af250e9784b6a92c.Kind()

	if __obf_5819e74754f40cb2 == reflect.Pointer {
		if _, __obf_ccfb92cc26e4569f := __obf_126a303617e63b80.Load(__obf_af250e9784b6a92c.Elem()); __obf_ccfb92cc26e4569f {
			return __obf_df4182cb1e43434e(__obf_af250e9784b6a92c)
		}
	}

	if __obf_af250e9784b6a92c.Implements(__obf_3dd0ef54e702a457) {
		return __obf_9a565c6ab17a891d
	}
	if __obf_af250e9784b6a92c.Implements(__obf_6ff28f42337db34f) {
		return __obf_14e168d17d6102ca
	}
	if __obf_af250e9784b6a92c.Implements(__obf_a091a0cc441d3614) {
		return __obf_ba4cbbe1051a90bf
	}
	if __obf_af250e9784b6a92c.Implements(__obf_4acc62fe1992007c) {
		return __obf_4270b572933fcfeb
	}

	// Addressable struct field value.
	if __obf_5819e74754f40cb2 != reflect.Pointer {
		__obf_8f0c71619c0382f1 := reflect.PointerTo(__obf_af250e9784b6a92c)
		if __obf_8f0c71619c0382f1.Implements(__obf_3dd0ef54e702a457) {
			return __obf_05608f29442fda3a
		}
		if __obf_8f0c71619c0382f1.Implements(__obf_6ff28f42337db34f) {
			return __obf_74eda24509529e7e
		}
		if __obf_8f0c71619c0382f1.Implements(__obf_a091a0cc441d3614) {
			return __obf_c2017404f305d278
		}
		if __obf_8f0c71619c0382f1.Implements(__obf_4acc62fe1992007c) {
			return __obf_6e106acfdd7bcffd
		}
	}

	if __obf_af250e9784b6a92c == __obf_6d051195c59d18d2 {
		return __obf_a74fd11809985098
	}

	switch __obf_5819e74754f40cb2 {
	case reflect.Ptr:
		return __obf_df4182cb1e43434e(__obf_af250e9784b6a92c)
	case reflect.Slice:
		__obf_a3aef353a1d07512 := __obf_af250e9784b6a92c.Elem()
		if __obf_a3aef353a1d07512.Kind() == reflect.Uint8 {
			return __obf_636a2f8dbc830828
		}
		if __obf_a3aef353a1d07512 == __obf_2566dc21412a8f7b {
			return __obf_1f2df5fa72bbce12
		}
	case reflect.Array:
		if __obf_af250e9784b6a92c.Elem().Kind() == reflect.Uint8 {
			return __obf_ddd3546e24350c75
		}
	case reflect.Map:
		if __obf_af250e9784b6a92c.Key() == __obf_2566dc21412a8f7b {
			switch __obf_af250e9784b6a92c.Elem() {
			case __obf_2566dc21412a8f7b:
				return __obf_be42c64a29c7a474
			case __obf_e717f08f63c1455d:
				return __obf_a421402c0bc3fb7f
			case __obf_967569602a7564fc:
				return __obf_f18ed5a38c2a6b1a
			}
		}
	}

	return __obf_4846e675fa6fb7a6[__obf_5819e74754f40cb2]
}

func __obf_df4182cb1e43434e(__obf_af250e9784b6a92c reflect.Type) __obf_b2b0b98cf9f599e0 {
	__obf_88d8471374f73dd1 := __obf_1e49c68fa4194051(__obf_af250e9784b6a92c.Elem())
	return func(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if __obf_85d270343a0dfe14.IsNil() {
			return __obf_84d0d31a8288f191.EncodeNil()
		}
		return __obf_88d8471374f73dd1(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14.Elem())
	}
}

func __obf_05608f29442fda3a(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if !__obf_85d270343a0dfe14.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_85d270343a0dfe14.Interface())
	}
	__obf_88d8471374f73dd1 := __obf_85d270343a0dfe14.Addr().Interface().(CustomEncoder)
	return __obf_88d8471374f73dd1.EncodeMsgpack(__obf_84d0d31a8288f191)
}

func __obf_9a565c6ab17a891d(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_6238d8f3c0669e3b(__obf_85d270343a0dfe14.Kind()) && __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	__obf_88d8471374f73dd1 := __obf_85d270343a0dfe14.Interface().(CustomEncoder)
	return __obf_88d8471374f73dd1.EncodeMsgpack(__obf_84d0d31a8288f191)
}

func __obf_74eda24509529e7e(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if !__obf_85d270343a0dfe14.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_85d270343a0dfe14.Interface())
	}
	return __obf_14e168d17d6102ca(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14.Addr())
}

func __obf_14e168d17d6102ca(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_6238d8f3c0669e3b(__obf_85d270343a0dfe14.Kind()) && __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	__obf_827aeedcead80d9b := __obf_85d270343a0dfe14.Interface().(Marshaler)
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_827aeedcead80d9b.MarshalMsgpack()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	_, __obf_3aa78ad28f50ed46 = __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.Write(__obf_9b4a5a04bdad2347)
	return __obf_3aa78ad28f50ed46
}

func __obf_19ddf8a0020ded94(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return __obf_84d0d31a8288f191.EncodeBool(__obf_85d270343a0dfe14.Bool())
}

func __obf_85f45bdd273b38e6(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	return __obf_84d0d31a8288f191.EncodeValue(__obf_85d270343a0dfe14.Elem())
}

func __obf_a74fd11809985098(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	return __obf_84d0d31a8288f191.EncodeString(__obf_85d270343a0dfe14.Interface().(error).Error())
}

func __obf_f2e27ee6992d8676(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return fmt.Errorf("msgpack: Encode(unsupported %s)", __obf_85d270343a0dfe14.Type())
}

func __obf_6238d8f3c0669e3b(__obf_5819e74754f40cb2 reflect.Kind) bool {
	switch __obf_5819e74754f40cb2 {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func __obf_c2e4305473db18c5(__obf_862c8bdd067ceabe reflect.Type) bool {
	if __obf_862c8bdd067ceabe.Kind() == reflect.Ptr {
		__obf_862c8bdd067ceabe = __obf_862c8bdd067ceabe.Elem()
	}
	return __obf_6238d8f3c0669e3b(__obf_862c8bdd067ceabe.Kind())
}

//------------------------------------------------------------------------------

func __obf_c2017404f305d278(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if !__obf_85d270343a0dfe14.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_85d270343a0dfe14.Interface())
	}
	return __obf_ba4cbbe1051a90bf(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14.Addr())
}

func __obf_ba4cbbe1051a90bf(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_6238d8f3c0669e3b(__obf_85d270343a0dfe14.Kind()) && __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	__obf_827aeedcead80d9b := __obf_85d270343a0dfe14.Interface().(encoding.BinaryMarshaler)
	__obf_2171bf5dcbdd2ac8, __obf_3aa78ad28f50ed46 := __obf_827aeedcead80d9b.MarshalBinary()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	return __obf_84d0d31a8288f191.EncodeBytes(__obf_2171bf5dcbdd2ac8)
}

//------------------------------------------------------------------------------

func __obf_6e106acfdd7bcffd(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if !__obf_85d270343a0dfe14.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_85d270343a0dfe14.Interface())
	}
	return __obf_4270b572933fcfeb(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14.Addr())
}

func __obf_4270b572933fcfeb(__obf_84d0d31a8288f191 *Encoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_6238d8f3c0669e3b(__obf_85d270343a0dfe14.Kind()) && __obf_85d270343a0dfe14.IsNil() {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	__obf_827aeedcead80d9b := __obf_85d270343a0dfe14.Interface().(encoding.TextMarshaler)
	__obf_2171bf5dcbdd2ac8, __obf_3aa78ad28f50ed46 := __obf_827aeedcead80d9b.MarshalText()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	return __obf_84d0d31a8288f191.EncodeBytes(__obf_2171bf5dcbdd2ac8)
}
