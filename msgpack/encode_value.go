package __obf_a4aad98aaf3178f4

import (
	"encoding"
	"fmt"
	"reflect"
)

var __obf_033549c43dbd64fa []__obf_9d8023233a199a4f

//nolint:gochecknoinits
func init() {
	__obf_033549c43dbd64fa = []__obf_9d8023233a199a4f{
		reflect.Bool:          __obf_a56f1b1deeece1b6,
		reflect.Int:           __obf_82e4dec38ff6b0f0,
		reflect.Int8:          __obf_1f3e7325737988bb,
		reflect.Int16:         __obf_8438172b55bae72b,
		reflect.Int32:         __obf_3f87414aacb11a03,
		reflect.Int64:         __obf_b43612dd4961a390,
		reflect.Uint:          __obf_7356d4170a0d75c7,
		reflect.Uint8:         __obf_5234f218217cc405,
		reflect.Uint16:        __obf_707d2811334e2066,
		reflect.Uint32:        __obf_6a77622f3635270d,
		reflect.Uint64:        __obf_2e07f9acaad90e64,
		reflect.Float32:       __obf_947000f09dd30923,
		reflect.Float64:       __obf_2053638319e34fb7,
		reflect.Complex64:     __obf_610ab0bb055f21c6,
		reflect.Complex128:    __obf_610ab0bb055f21c6,
		reflect.Array:         __obf_5711cd74f1e4cfdc,
		reflect.Chan:          __obf_610ab0bb055f21c6,
		reflect.Func:          __obf_610ab0bb055f21c6,
		reflect.Interface:     __obf_284855b7eb206cc5,
		reflect.Map:           __obf_c6d71f18960117fe,
		reflect.Ptr:           __obf_610ab0bb055f21c6,
		reflect.Slice:         __obf_6abbdd3481da14fa,
		reflect.String:        __obf_52eb49ecc239146d,
		reflect.Struct:        __obf_8fcfe77f763e1de9,
		reflect.UnsafePointer: __obf_610ab0bb055f21c6,
	}
}

func __obf_f1f0aa6d8078637c(__obf_bbfd30fcc08dc1bc reflect.Type) __obf_9d8023233a199a4f {
	if __obf_a1f43267eeb48a1a, __obf_81b19f2a6159ab89 := __obf_a25a843d85821e5c.Load(__obf_bbfd30fcc08dc1bc); __obf_81b19f2a6159ab89 {
		return __obf_a1f43267eeb48a1a.(__obf_9d8023233a199a4f)
	}
	__obf_64fe495dbc0c543b := _getEncoder(__obf_bbfd30fcc08dc1bc)
	__obf_a25a843d85821e5c.
		Store(__obf_bbfd30fcc08dc1bc, __obf_64fe495dbc0c543b)
	return __obf_64fe495dbc0c543b
}

func _getEncoder(__obf_bbfd30fcc08dc1bc reflect.Type) __obf_9d8023233a199a4f {
	__obf_91ae1c24bc584c3a := __obf_bbfd30fcc08dc1bc.Kind()

	if __obf_91ae1c24bc584c3a == reflect.Pointer {
		if _, __obf_81b19f2a6159ab89 := __obf_a25a843d85821e5c.Load(__obf_bbfd30fcc08dc1bc.Elem()); __obf_81b19f2a6159ab89 {
			return __obf_67fd3414b2972b01(__obf_bbfd30fcc08dc1bc)
		}
	}

	if __obf_bbfd30fcc08dc1bc.Implements(__obf_a644aad6e1b3ff55) {
		return __obf_cd0fd08ee03fa431
	}
	if __obf_bbfd30fcc08dc1bc.Implements(__obf_1b905c6b55da13d6) {
		return __obf_b53f7a3574904a29
	}
	if __obf_bbfd30fcc08dc1bc.Implements(__obf_8b41e5fad6d95fea) {
		return __obf_3e5a30549997b534
	}
	if __obf_bbfd30fcc08dc1bc.Implements(__obf_efdcdb0d1ee35009) {
		return __obf_f2bfb354310e5a4e
	}

	// Addressable struct field value.
	if __obf_91ae1c24bc584c3a != reflect.Pointer {
		__obf_cf3077802722ecc2 := reflect.PointerTo(__obf_bbfd30fcc08dc1bc)
		if __obf_cf3077802722ecc2.Implements(__obf_a644aad6e1b3ff55) {
			return __obf_e538a3d3fcfe6fc1
		}
		if __obf_cf3077802722ecc2.Implements(__obf_1b905c6b55da13d6) {
			return __obf_0866d6dbeaee01bf
		}
		if __obf_cf3077802722ecc2.Implements(__obf_8b41e5fad6d95fea) {
			return __obf_97c3522eb32aa7a9
		}
		if __obf_cf3077802722ecc2.Implements(__obf_efdcdb0d1ee35009) {
			return __obf_7c02d200ea5c2a0d
		}
	}

	if __obf_bbfd30fcc08dc1bc == __obf_d66851dd29a0f414 {
		return __obf_e4fb09d82d9dcfe3
	}

	switch __obf_91ae1c24bc584c3a {
	case reflect.Ptr:
		return __obf_67fd3414b2972b01(__obf_bbfd30fcc08dc1bc)
	case reflect.Slice:
		__obf_ed47b1177873a7da := __obf_bbfd30fcc08dc1bc.Elem()
		if __obf_ed47b1177873a7da.Kind() == reflect.Uint8 {
			return __obf_7b11a1a1aaa56165
		}
		if __obf_ed47b1177873a7da == __obf_60ac326f01b0b63a {
			return __obf_7c02e0223b928d18
		}
	case reflect.Array:
		if __obf_bbfd30fcc08dc1bc.Elem().Kind() == reflect.Uint8 {
			return __obf_7f7b053e6838c23d
		}
	case reflect.Map:
		if __obf_bbfd30fcc08dc1bc.Key() == __obf_60ac326f01b0b63a {
			switch __obf_bbfd30fcc08dc1bc.Elem() {
			case __obf_60ac326f01b0b63a:
				return __obf_af0bc339c396d17f
			case __obf_5b19c38666a256ef:
				return __obf_a8b5c3d4b26ca8c6
			case __obf_3eea2b4982d29baa:
				return __obf_76c5385b22000a50
			}
		}
	}

	return __obf_033549c43dbd64fa[__obf_91ae1c24bc584c3a]
}

func __obf_67fd3414b2972b01(__obf_bbfd30fcc08dc1bc reflect.Type) __obf_9d8023233a199a4f {
	__obf_c4f748de4fd652e4 := __obf_f1f0aa6d8078637c(__obf_bbfd30fcc08dc1bc.Elem())
	return func(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if __obf_a1f43267eeb48a1a.IsNil() {
			return __obf_2c8e97779108ab17.EncodeNil()
		}
		return __obf_c4f748de4fd652e4(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a.Elem())
	}
}

func __obf_e538a3d3fcfe6fc1(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if !__obf_a1f43267eeb48a1a.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_a1f43267eeb48a1a.Interface())
	}
	__obf_c4f748de4fd652e4 := __obf_a1f43267eeb48a1a.Addr().Interface().(CustomEncoder)
	return __obf_c4f748de4fd652e4.EncodeMsgpack(__obf_2c8e97779108ab17)
}

func __obf_cd0fd08ee03fa431(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_178efdbad2797673(__obf_a1f43267eeb48a1a.Kind()) && __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	__obf_c4f748de4fd652e4 := __obf_a1f43267eeb48a1a.Interface().(CustomEncoder)
	return __obf_c4f748de4fd652e4.EncodeMsgpack(__obf_2c8e97779108ab17)
}

func __obf_0866d6dbeaee01bf(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if !__obf_a1f43267eeb48a1a.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_a1f43267eeb48a1a.Interface())
	}
	return __obf_b53f7a3574904a29(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a.Addr())
}

func __obf_b53f7a3574904a29(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_178efdbad2797673(__obf_a1f43267eeb48a1a.Kind()) && __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	__obf_f68fedb43a0f697e := __obf_a1f43267eeb48a1a.Interface().(Marshaler)
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_f68fedb43a0f697e.MarshalMsgpack()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	_, __obf_4061ca0039150c39 = __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.Write(__obf_f57209cfda8e17cf)
	return __obf_4061ca0039150c39
}

func __obf_a56f1b1deeece1b6(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return __obf_2c8e97779108ab17.EncodeBool(__obf_a1f43267eeb48a1a.Bool())
}

func __obf_284855b7eb206cc5(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	return __obf_2c8e97779108ab17.EncodeValue(__obf_a1f43267eeb48a1a.Elem())
}

func __obf_e4fb09d82d9dcfe3(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	return __obf_2c8e97779108ab17.EncodeString(__obf_a1f43267eeb48a1a.Interface().(error).Error())
}

func __obf_610ab0bb055f21c6(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return fmt.Errorf("msgpack: Encode(unsupported %s)", __obf_a1f43267eeb48a1a.Type())
}

func __obf_178efdbad2797673(__obf_91ae1c24bc584c3a reflect.Kind) bool {
	switch __obf_91ae1c24bc584c3a {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func __obf_98cd53591a84c7fa(__obf_8698f5d73996ef26 reflect.Type) bool {
	if __obf_8698f5d73996ef26.Kind() == reflect.Ptr {
		__obf_8698f5d73996ef26 = __obf_8698f5d73996ef26.Elem()
	}
	return __obf_178efdbad2797673(__obf_8698f5d73996ef26.Kind())
}

//------------------------------------------------------------------------------

func __obf_97c3522eb32aa7a9(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if !__obf_a1f43267eeb48a1a.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_a1f43267eeb48a1a.Interface())
	}
	return __obf_3e5a30549997b534(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a.Addr())
}

func __obf_3e5a30549997b534(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_178efdbad2797673(__obf_a1f43267eeb48a1a.Kind()) && __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	__obf_f68fedb43a0f697e := __obf_a1f43267eeb48a1a.Interface().(encoding.BinaryMarshaler)
	__obf_69e4c283179273ce, __obf_4061ca0039150c39 := __obf_f68fedb43a0f697e.MarshalBinary()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	return __obf_2c8e97779108ab17.EncodeBytes(__obf_69e4c283179273ce)
}

//------------------------------------------------------------------------------

func __obf_7c02d200ea5c2a0d(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if !__obf_a1f43267eeb48a1a.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_a1f43267eeb48a1a.Interface())
	}
	return __obf_f2bfb354310e5a4e(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a.Addr())
}

func __obf_f2bfb354310e5a4e(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_178efdbad2797673(__obf_a1f43267eeb48a1a.Kind()) && __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	__obf_f68fedb43a0f697e := __obf_a1f43267eeb48a1a.Interface().(encoding.TextMarshaler)
	__obf_69e4c283179273ce, __obf_4061ca0039150c39 := __obf_f68fedb43a0f697e.MarshalText()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	return __obf_2c8e97779108ab17.EncodeBytes(__obf_69e4c283179273ce)
}
