package __obf_3edaa49e853afa16

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
)

var (
	__obf_967569602a7564fc = reflect.TypeOf((*any)(nil)).Elem()
	__obf_2566dc21412a8f7b = reflect.TypeOf((*string)(nil)).Elem()
	__obf_e717f08f63c1455d = reflect.TypeOf((*bool)(nil)).Elem()
)

var __obf_baf87a19628dae6e []__obf_47bea90bc2ca8162

//nolint:gochecknoinits
func init() {
	__obf_baf87a19628dae6e = []__obf_47bea90bc2ca8162{
		reflect.Bool:          __obf_9fc5883588fe8960,
		reflect.Int:           __obf_69ca6ad8831c417d,
		reflect.Int8:          __obf_69ca6ad8831c417d,
		reflect.Int16:         __obf_69ca6ad8831c417d,
		reflect.Int32:         __obf_69ca6ad8831c417d,
		reflect.Int64:         __obf_69ca6ad8831c417d,
		reflect.Uint:          __obf_e4e267de4597ac4a,
		reflect.Uint8:         __obf_e4e267de4597ac4a,
		reflect.Uint16:        __obf_e4e267de4597ac4a,
		reflect.Uint32:        __obf_e4e267de4597ac4a,
		reflect.Uint64:        __obf_e4e267de4597ac4a,
		reflect.Float32:       __obf_fe8f515e5461c435,
		reflect.Float64:       __obf_0ccce5c30ace0e82,
		reflect.Complex64:     __obf_30ba0253fb330c0e,
		reflect.Complex128:    __obf_30ba0253fb330c0e,
		reflect.Array:         __obf_2858ebecafe0e9d9,
		reflect.Chan:          __obf_30ba0253fb330c0e,
		reflect.Func:          __obf_30ba0253fb330c0e,
		reflect.Interface:     __obf_6b9d6a0bab505ca5,
		reflect.Map:           __obf_943cb3215d99d848,
		reflect.Pointer:       __obf_30ba0253fb330c0e,
		reflect.Slice:         __obf_84291db4b7f5231d,
		reflect.String:        __obf_e018477bba540f48,
		reflect.Struct:        __obf_4f268d405fba2c41,
		reflect.UnsafePointer: __obf_30ba0253fb330c0e,
	}
}

func __obf_6253a63b14aba59e(__obf_af250e9784b6a92c reflect.Type) __obf_47bea90bc2ca8162 {
	if __obf_85d270343a0dfe14, __obf_ccfb92cc26e4569f := __obf_8d7a14362d6f6fc4.Load(__obf_af250e9784b6a92c); __obf_ccfb92cc26e4569f {
		return __obf_85d270343a0dfe14.(__obf_47bea90bc2ca8162)
	}
	__obf_c8137d8e88d5b1dd := _getDecoder(__obf_af250e9784b6a92c)
	__obf_8d7a14362d6f6fc4.
		Store(__obf_af250e9784b6a92c, __obf_c8137d8e88d5b1dd)
	return __obf_c8137d8e88d5b1dd
}

func _getDecoder(__obf_af250e9784b6a92c reflect.Type) __obf_47bea90bc2ca8162 {
	__obf_5819e74754f40cb2 := __obf_af250e9784b6a92c.Kind()

	if __obf_5819e74754f40cb2 == reflect.Ptr {
		if _, __obf_ccfb92cc26e4569f := __obf_8d7a14362d6f6fc4.Load(__obf_af250e9784b6a92c.Elem()); __obf_ccfb92cc26e4569f {
			return __obf_c20829a2061f17c4(__obf_af250e9784b6a92c)
		}
	}

	if __obf_af250e9784b6a92c.Implements(__obf_3fcf1c5b6db37260) {
		return __obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_28f6a0edb9a78b3b)
	}
	if __obf_af250e9784b6a92c.Implements(__obf_b2238d76404cb5bd) {
		return __obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_d42b547c18fe3ce4)
	}
	if __obf_af250e9784b6a92c.Implements(__obf_217405f1e132acef) {
		return __obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_7c8c572e8c0b9d8d)
	}
	if __obf_af250e9784b6a92c.Implements(__obf_9d24d27869bcc179) {
		return __obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_dcc4141ef0fecd3c)
	}

	// Addressable struct field value.
	if __obf_5819e74754f40cb2 != reflect.Pointer {
		__obf_8f0c71619c0382f1 := reflect.PointerTo(__obf_af250e9784b6a92c)
		if __obf_8f0c71619c0382f1.Implements(__obf_3fcf1c5b6db37260) {
			return __obf_1e5d4f1ad9c79696(__obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_28f6a0edb9a78b3b))
		}
		if __obf_8f0c71619c0382f1.Implements(__obf_b2238d76404cb5bd) {
			return __obf_1e5d4f1ad9c79696(__obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_d42b547c18fe3ce4))
		}
		if __obf_8f0c71619c0382f1.Implements(__obf_217405f1e132acef) {
			return __obf_1e5d4f1ad9c79696(__obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_7c8c572e8c0b9d8d))
		}
		if __obf_8f0c71619c0382f1.Implements(__obf_9d24d27869bcc179) {
			return __obf_1e5d4f1ad9c79696(__obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c, __obf_dcc4141ef0fecd3c))
		}
	}

	switch __obf_5819e74754f40cb2 {
	case reflect.Pointer:
		return __obf_c20829a2061f17c4(__obf_af250e9784b6a92c)
	case reflect.Slice:
		__obf_a3aef353a1d07512 := __obf_af250e9784b6a92c.Elem()
		if __obf_a3aef353a1d07512.Kind() == reflect.Uint8 {
			return __obf_9d816678b5cb9529
		}
		if __obf_a3aef353a1d07512 == __obf_2566dc21412a8f7b {
			return __obf_22486c0586fcfbc8
		}
	case reflect.Array:
		if __obf_af250e9784b6a92c.Elem().Kind() == reflect.Uint8 {
			return __obf_dcb87a3c2d394964
		}
	case reflect.Map:
		if __obf_af250e9784b6a92c.Key() == __obf_2566dc21412a8f7b {
			switch __obf_af250e9784b6a92c.Elem() {
			case __obf_2566dc21412a8f7b:
				return __obf_107f1a504832f4ec
			case __obf_967569602a7564fc:
				return __obf_aa82e0459275f272
			}
		}
	}

	return __obf_baf87a19628dae6e[__obf_5819e74754f40cb2]
}

func __obf_c20829a2061f17c4(__obf_af250e9784b6a92c reflect.Type) __obf_47bea90bc2ca8162 {
	__obf_2cb4408eb815d713 := __obf_6253a63b14aba59e(__obf_af250e9784b6a92c.Elem())
	return func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if __obf_015afbee33862a01.__obf_7773819fc7c589cb() {
			if !__obf_85d270343a0dfe14.IsNil() {
				__obf_85d270343a0dfe14.
					Set(__obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_af250e9784b6a92c).Elem())
			}
			return __obf_015afbee33862a01.DecodeNil()
		}
		if __obf_85d270343a0dfe14.IsNil() {
			__obf_85d270343a0dfe14.
				Set(__obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_af250e9784b6a92c.Elem()))
		}
		return __obf_2cb4408eb815d713(__obf_015afbee33862a01, __obf_85d270343a0dfe14.Elem())
	}
}

func __obf_1e5d4f1ad9c79696(__obf_c8137d8e88d5b1dd __obf_47bea90bc2ca8162) __obf_47bea90bc2ca8162 {
	return func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if !__obf_85d270343a0dfe14.CanAddr() {
			return fmt.Errorf("msgpack: Decode(nonaddressable %T)", __obf_85d270343a0dfe14.Interface())
		}
		return __obf_c8137d8e88d5b1dd(__obf_015afbee33862a01, __obf_85d270343a0dfe14.Addr())
	}
}

func __obf_fdf8f644be3d1bdc(__obf_af250e9784b6a92c reflect.Type, __obf_c8137d8e88d5b1dd __obf_47bea90bc2ca8162) __obf_47bea90bc2ca8162 {
	if __obf_6238d8f3c0669e3b(__obf_af250e9784b6a92c.Kind()) {
		return func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
			if __obf_015afbee33862a01.__obf_7773819fc7c589cb() {
				return __obf_015afbee33862a01.__obf_38b8c150d8f5b61b(__obf_85d270343a0dfe14)
			}
			if __obf_85d270343a0dfe14.IsNil() {
				__obf_85d270343a0dfe14.
					Set(__obf_015afbee33862a01.__obf_4c87e5dcce642cc7(__obf_af250e9784b6a92c.Elem()))
			}
			return __obf_c8137d8e88d5b1dd(__obf_015afbee33862a01, __obf_85d270343a0dfe14)
		}
	}

	return func(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
		if __obf_015afbee33862a01.__obf_7773819fc7c589cb() {
			return __obf_015afbee33862a01.__obf_38b8c150d8f5b61b(__obf_85d270343a0dfe14)
		}
		return __obf_c8137d8e88d5b1dd(__obf_015afbee33862a01, __obf_85d270343a0dfe14)
	}
}

func __obf_9fc5883588fe8960(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_e0c6c53aed9fbfd4, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeBool()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetBool(__obf_e0c6c53aed9fbfd4)
	return nil
}

func __obf_6b9d6a0bab505ca5(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_015afbee33862a01.__obf_04d610da04f1230b(__obf_85d270343a0dfe14)
	}
	return __obf_015afbee33862a01.DecodeValue(__obf_85d270343a0dfe14.Elem())
}

func (__obf_015afbee33862a01 *Decoder) __obf_04d610da04f1230b(__obf_85d270343a0dfe14 reflect.Value) error {
	__obf_ab0d31723a444d49, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if __obf_ab0d31723a444d49 != nil {
		if __obf_85d270343a0dfe14.Type() == __obf_6d051195c59d18d2 {
			if __obf_ab0d31723a444d49, __obf_ccfb92cc26e4569f := __obf_ab0d31723a444d49.(string); __obf_ccfb92cc26e4569f {
				__obf_85d270343a0dfe14.
					Set(reflect.ValueOf(errors.New(__obf_ab0d31723a444d49)))
				return nil
			}
		}
		__obf_85d270343a0dfe14.
			Set(reflect.ValueOf(__obf_ab0d31723a444d49))
	}

	return nil
}

func __obf_30ba0253fb330c0e(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	return fmt.Errorf("msgpack: Decode(unsupported %s)", __obf_85d270343a0dfe14.Type())
}

//------------------------------------------------------------------------------

func __obf_28f6a0edb9a78b3b(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_2cb4408eb815d713 := __obf_85d270343a0dfe14.Interface().(CustomDecoder)
	return __obf_2cb4408eb815d713.DecodeMsgpack(__obf_015afbee33862a01)
}

func __obf_d42b547c18fe3ce4(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	var __obf_9b4a5a04bdad2347 []byte
	__obf_015afbee33862a01.__obf_e11087a315c4c191 = make([]byte, 0, 64)
	if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_9b4a5a04bdad2347 = __obf_015afbee33862a01.__obf_e11087a315c4c191
	__obf_015afbee33862a01.__obf_e11087a315c4c191 = nil
	__obf_fff7241b89c79538 := __obf_85d270343a0dfe14.Interface().(Unmarshaler)
	return __obf_fff7241b89c79538.UnmarshalMsgpack(__obf_9b4a5a04bdad2347)
}

func __obf_7c8c572e8c0b9d8d(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_2171bf5dcbdd2ac8, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeBytes()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_fff7241b89c79538 := __obf_85d270343a0dfe14.Interface().(encoding.BinaryUnmarshaler)
	return __obf_fff7241b89c79538.UnmarshalBinary(__obf_2171bf5dcbdd2ac8)
}

func __obf_dcc4141ef0fecd3c(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_2171bf5dcbdd2ac8, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeBytes()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_fff7241b89c79538 := __obf_85d270343a0dfe14.Interface().(encoding.TextUnmarshaler)
	return __obf_fff7241b89c79538.UnmarshalText(__obf_2171bf5dcbdd2ac8)
}
