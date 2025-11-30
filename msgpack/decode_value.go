package __obf_fd770cb9675903b0

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
)

var (
	__obf_533b257584aace5f = reflect.TypeOf((*any)(nil)).Elem()
	__obf_e7498c925721ef3a = reflect.TypeOf((*string)(nil)).Elem()
	__obf_a9624631014b694f = reflect.TypeOf((*bool)(nil)).Elem()
)

var __obf_def4798fa57068b9 []__obf_fb10c6a4667e72a5

//nolint:gochecknoinits
func init() {
	__obf_def4798fa57068b9 = []__obf_fb10c6a4667e72a5{
		reflect.Bool:          __obf_8eb4f218669b6f42,
		reflect.Int:           __obf_7801128403b4eb1b,
		reflect.Int8:          __obf_7801128403b4eb1b,
		reflect.Int16:         __obf_7801128403b4eb1b,
		reflect.Int32:         __obf_7801128403b4eb1b,
		reflect.Int64:         __obf_7801128403b4eb1b,
		reflect.Uint:          __obf_f5ab33312c20fcd4,
		reflect.Uint8:         __obf_f5ab33312c20fcd4,
		reflect.Uint16:        __obf_f5ab33312c20fcd4,
		reflect.Uint32:        __obf_f5ab33312c20fcd4,
		reflect.Uint64:        __obf_f5ab33312c20fcd4,
		reflect.Float32:       __obf_028e0e88467a2c89,
		reflect.Float64:       __obf_fdb9efbf3bc19d4f,
		reflect.Complex64:     __obf_1d64332cad2806a4,
		reflect.Complex128:    __obf_1d64332cad2806a4,
		reflect.Array:         __obf_b694339ec1e870f3,
		reflect.Chan:          __obf_1d64332cad2806a4,
		reflect.Func:          __obf_1d64332cad2806a4,
		reflect.Interface:     __obf_604515803d5e5b1f,
		reflect.Map:           __obf_274fc8cba465d312,
		reflect.Pointer:       __obf_1d64332cad2806a4,
		reflect.Slice:         __obf_bee2295884206fce,
		reflect.String:        __obf_5ac5b748094f0aee,
		reflect.Struct:        __obf_9fde65ded02d475f,
		reflect.UnsafePointer: __obf_1d64332cad2806a4,
	}
}

func __obf_538e4e278f06c9d6(__obf_8733059f76088ffc reflect.Type) __obf_fb10c6a4667e72a5 {
	if __obf_f328a048f76b7256, __obf_b00b3c6a10f90467 := __obf_9ea14163dfc9e317.Load(__obf_8733059f76088ffc); __obf_b00b3c6a10f90467 {
		return __obf_f328a048f76b7256.(__obf_fb10c6a4667e72a5)
	}
	__obf_1e4576e8508b04bc := _getDecoder(__obf_8733059f76088ffc)
	__obf_9ea14163dfc9e317.
		Store(__obf_8733059f76088ffc, __obf_1e4576e8508b04bc)
	return __obf_1e4576e8508b04bc
}

func _getDecoder(__obf_8733059f76088ffc reflect.Type) __obf_fb10c6a4667e72a5 {
	__obf_5cd51d276078fe9d := __obf_8733059f76088ffc.Kind()

	if __obf_5cd51d276078fe9d == reflect.Ptr {
		if _, __obf_b00b3c6a10f90467 := __obf_9ea14163dfc9e317.Load(__obf_8733059f76088ffc.Elem()); __obf_b00b3c6a10f90467 {
			return __obf_39068db44507eae5(__obf_8733059f76088ffc)
		}
	}

	if __obf_8733059f76088ffc.Implements(__obf_3370d035693fe504) {
		return __obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_a11fa4cb3df57f60)
	}
	if __obf_8733059f76088ffc.Implements(__obf_d8f141f48ff21664) {
		return __obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_1f872f397d848cdb)
	}
	if __obf_8733059f76088ffc.Implements(__obf_4f3a83c066a65856) {
		return __obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_08d3282a99017ddb)
	}
	if __obf_8733059f76088ffc.Implements(__obf_5d3e3730f188b5b3) {
		return __obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_4cde2ffec0b92fef)
	}

	// Addressable struct field value.
	if __obf_5cd51d276078fe9d != reflect.Pointer {
		__obf_2c49f9d2007cfaf6 := reflect.PointerTo(__obf_8733059f76088ffc)
		if __obf_2c49f9d2007cfaf6.Implements(__obf_3370d035693fe504) {
			return __obf_b20f2a65414ff8c5(__obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_a11fa4cb3df57f60))
		}
		if __obf_2c49f9d2007cfaf6.Implements(__obf_d8f141f48ff21664) {
			return __obf_b20f2a65414ff8c5(__obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_1f872f397d848cdb))
		}
		if __obf_2c49f9d2007cfaf6.Implements(__obf_4f3a83c066a65856) {
			return __obf_b20f2a65414ff8c5(__obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_08d3282a99017ddb))
		}
		if __obf_2c49f9d2007cfaf6.Implements(__obf_5d3e3730f188b5b3) {
			return __obf_b20f2a65414ff8c5(__obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, __obf_4cde2ffec0b92fef))
		}
	}

	switch __obf_5cd51d276078fe9d {
	case reflect.Pointer:
		return __obf_39068db44507eae5(__obf_8733059f76088ffc)
	case reflect.Slice:
		__obf_d6cb414a13d9d0f7 := __obf_8733059f76088ffc.Elem()
		if __obf_d6cb414a13d9d0f7.Kind() == reflect.Uint8 {
			return __obf_322f79ef2f26d64d
		}
		if __obf_d6cb414a13d9d0f7 == __obf_e7498c925721ef3a {
			return __obf_d77a8d5f087ab98c
		}
	case reflect.Array:
		if __obf_8733059f76088ffc.Elem().Kind() == reflect.Uint8 {
			return __obf_bcddc3fae95e8e05
		}
	case reflect.Map:
		if __obf_8733059f76088ffc.Key() == __obf_e7498c925721ef3a {
			switch __obf_8733059f76088ffc.Elem() {
			case __obf_e7498c925721ef3a:
				return __obf_9a941cbc62083f7f
			case __obf_533b257584aace5f:
				return __obf_80486e14ddfacbaf
			}
		}
	}

	return __obf_def4798fa57068b9[__obf_5cd51d276078fe9d]
}

func __obf_39068db44507eae5(__obf_8733059f76088ffc reflect.Type) __obf_fb10c6a4667e72a5 {
	__obf_25e78266d5ba1944 := __obf_538e4e278f06c9d6(__obf_8733059f76088ffc.Elem())
	return func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
		if __obf_5d389b660eefb08c.__obf_982d42aefd776ffb() {
			if !__obf_f328a048f76b7256.IsNil() {
				__obf_f328a048f76b7256.
					Set(__obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_8733059f76088ffc).Elem())
			}
			return __obf_5d389b660eefb08c.DecodeNil()
		}
		if __obf_f328a048f76b7256.IsNil() {
			__obf_f328a048f76b7256.
				Set(__obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_8733059f76088ffc.Elem()))
		}
		return __obf_25e78266d5ba1944(__obf_5d389b660eefb08c, __obf_f328a048f76b7256.Elem())
	}
}

func __obf_b20f2a65414ff8c5(__obf_1e4576e8508b04bc __obf_fb10c6a4667e72a5) __obf_fb10c6a4667e72a5 {
	return func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
		if !__obf_f328a048f76b7256.CanAddr() {
			return fmt.Errorf("msgpack: Decode(nonaddressable %T)", __obf_f328a048f76b7256.Interface())
		}
		return __obf_1e4576e8508b04bc(__obf_5d389b660eefb08c, __obf_f328a048f76b7256.Addr())
	}
}

func __obf_ca3ac1e2410413b5(__obf_8733059f76088ffc reflect.Type, __obf_1e4576e8508b04bc __obf_fb10c6a4667e72a5) __obf_fb10c6a4667e72a5 {
	if __obf_066825d3e1771287(__obf_8733059f76088ffc.Kind()) {
		return func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
			if __obf_5d389b660eefb08c.__obf_982d42aefd776ffb() {
				return __obf_5d389b660eefb08c.__obf_4e408f1c6737e5d3(__obf_f328a048f76b7256)
			}
			if __obf_f328a048f76b7256.IsNil() {
				__obf_f328a048f76b7256.
					Set(__obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_8733059f76088ffc.Elem()))
			}
			return __obf_1e4576e8508b04bc(__obf_5d389b660eefb08c, __obf_f328a048f76b7256)
		}
	}

	return func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
		if __obf_5d389b660eefb08c.__obf_982d42aefd776ffb() {
			return __obf_5d389b660eefb08c.__obf_4e408f1c6737e5d3(__obf_f328a048f76b7256)
		}
		return __obf_1e4576e8508b04bc(__obf_5d389b660eefb08c, __obf_f328a048f76b7256)
	}
}

func __obf_8eb4f218669b6f42(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_97f4b149581b6703, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeBool()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetBool(__obf_97f4b149581b6703)
	return nil
}

func __obf_604515803d5e5b1f(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_5d389b660eefb08c.__obf_82e46e35a168cbd0(__obf_f328a048f76b7256)
	}
	return __obf_5d389b660eefb08c.DecodeValue(__obf_f328a048f76b7256.Elem())
}

func (__obf_5d389b660eefb08c *Decoder) __obf_82e46e35a168cbd0(__obf_f328a048f76b7256 reflect.Value) error {
	__obf_e2c885d3da27fd46, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	if __obf_e2c885d3da27fd46 != nil {
		if __obf_f328a048f76b7256.Type() == __obf_f1f3ff99eed138e6 {
			if __obf_e2c885d3da27fd46, __obf_b00b3c6a10f90467 := __obf_e2c885d3da27fd46.(string); __obf_b00b3c6a10f90467 {
				__obf_f328a048f76b7256.
					Set(reflect.ValueOf(errors.New(__obf_e2c885d3da27fd46)))
				return nil
			}
		}
		__obf_f328a048f76b7256.
			Set(reflect.ValueOf(__obf_e2c885d3da27fd46))
	}

	return nil
}

func __obf_1d64332cad2806a4(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	return fmt.Errorf("msgpack: Decode(unsupported %s)", __obf_f328a048f76b7256.Type())
}

//------------------------------------------------------------------------------

func __obf_a11fa4cb3df57f60(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_25e78266d5ba1944 := __obf_f328a048f76b7256.Interface().(CustomDecoder)
	return __obf_25e78266d5ba1944.DecodeMsgpack(__obf_5d389b660eefb08c)
}

func __obf_1f872f397d848cdb(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	var __obf_94333af0f0facd65 []byte
	__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = make([]byte, 0, 64)
	if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_94333af0f0facd65 = __obf_5d389b660eefb08c.__obf_98388ae8c389d758
	__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = nil
	__obf_804b18a8a092f040 := __obf_f328a048f76b7256.Interface().(Unmarshaler)
	return __obf_804b18a8a092f040.UnmarshalMsgpack(__obf_94333af0f0facd65)
}

func __obf_08d3282a99017ddb(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_cc76e8ed73142f28, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeBytes()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_804b18a8a092f040 := __obf_f328a048f76b7256.Interface().(encoding.BinaryUnmarshaler)
	return __obf_804b18a8a092f040.UnmarshalBinary(__obf_cc76e8ed73142f28)
}

func __obf_4cde2ffec0b92fef(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_cc76e8ed73142f28, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeBytes()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_804b18a8a092f040 := __obf_f328a048f76b7256.Interface().(encoding.TextUnmarshaler)
	return __obf_804b18a8a092f040.UnmarshalText(__obf_cc76e8ed73142f28)
}
