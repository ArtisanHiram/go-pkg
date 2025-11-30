package __obf_fd770cb9675903b0

import (
	"encoding"
	"fmt"
	"reflect"
)

var __obf_dfbdb30a1aae87b2 []__obf_dc120d370e164c91

//nolint:gochecknoinits
func init() {
	__obf_dfbdb30a1aae87b2 = []__obf_dc120d370e164c91{
		reflect.Bool:          __obf_8af2108e5f5a9336,
		reflect.Int:           __obf_3944b7d3af00ae61,
		reflect.Int8:          __obf_dad2c9af88c11156,
		reflect.Int16:         __obf_9728683ea899b38d,
		reflect.Int32:         __obf_d85919497ce72a5b,
		reflect.Int64:         __obf_546c6c5440e640a7,
		reflect.Uint:          __obf_d7a9f99d1b2d37cc,
		reflect.Uint8:         __obf_4a3a44a2fc6f392a,
		reflect.Uint16:        __obf_2647406311f2e1f7,
		reflect.Uint32:        __obf_d68e2f5e7bc5739e,
		reflect.Uint64:        __obf_51c4091845697c82,
		reflect.Float32:       __obf_709c5fa255b1f690,
		reflect.Float64:       __obf_f288d375e57ad4c1,
		reflect.Complex64:     __obf_d1d21523321fd9d0,
		reflect.Complex128:    __obf_d1d21523321fd9d0,
		reflect.Array:         __obf_2e9719daceaf70e4,
		reflect.Chan:          __obf_d1d21523321fd9d0,
		reflect.Func:          __obf_d1d21523321fd9d0,
		reflect.Interface:     __obf_3c71664b718cf2c4,
		reflect.Map:           __obf_50d653704a4e67b8,
		reflect.Ptr:           __obf_d1d21523321fd9d0,
		reflect.Slice:         __obf_352d72be21a6a0d7,
		reflect.String:        __obf_f3b7f586caae0bd6,
		reflect.Struct:        __obf_c52bd37d7024764c,
		reflect.UnsafePointer: __obf_d1d21523321fd9d0,
	}
}

func __obf_031e3a13e30d95dc(__obf_8733059f76088ffc reflect.Type) __obf_dc120d370e164c91 {
	if __obf_f328a048f76b7256, __obf_b00b3c6a10f90467 := __obf_56aff9039d970a50.Load(__obf_8733059f76088ffc); __obf_b00b3c6a10f90467 {
		return __obf_f328a048f76b7256.(__obf_dc120d370e164c91)
	}
	__obf_1e4576e8508b04bc := _getEncoder(__obf_8733059f76088ffc)
	__obf_56aff9039d970a50.
		Store(__obf_8733059f76088ffc, __obf_1e4576e8508b04bc)
	return __obf_1e4576e8508b04bc
}

func _getEncoder(__obf_8733059f76088ffc reflect.Type) __obf_dc120d370e164c91 {
	__obf_5cd51d276078fe9d := __obf_8733059f76088ffc.Kind()

	if __obf_5cd51d276078fe9d == reflect.Pointer {
		if _, __obf_b00b3c6a10f90467 := __obf_56aff9039d970a50.Load(__obf_8733059f76088ffc.Elem()); __obf_b00b3c6a10f90467 {
			return __obf_7a752c9836a9ee1e(__obf_8733059f76088ffc)
		}
	}

	if __obf_8733059f76088ffc.Implements(__obf_3dbe24e34d0a7693) {
		return __obf_f8f9ae3cc36e5f9d
	}
	if __obf_8733059f76088ffc.Implements(__obf_c3360c234e8a51cb) {
		return __obf_83060f2b2cde521c
	}
	if __obf_8733059f76088ffc.Implements(__obf_6e4a841f6465176f) {
		return __obf_03b1c3316fe255c4
	}
	if __obf_8733059f76088ffc.Implements(__obf_2fa03b50688c3358) {
		return __obf_afe2be14b2072c5c
	}

	// Addressable struct field value.
	if __obf_5cd51d276078fe9d != reflect.Pointer {
		__obf_2c49f9d2007cfaf6 := reflect.PointerTo(__obf_8733059f76088ffc)
		if __obf_2c49f9d2007cfaf6.Implements(__obf_3dbe24e34d0a7693) {
			return __obf_30978ab67c376dcd
		}
		if __obf_2c49f9d2007cfaf6.Implements(__obf_c3360c234e8a51cb) {
			return __obf_9e056da4627a0ac3
		}
		if __obf_2c49f9d2007cfaf6.Implements(__obf_6e4a841f6465176f) {
			return __obf_08cf0d4863d94784
		}
		if __obf_2c49f9d2007cfaf6.Implements(__obf_2fa03b50688c3358) {
			return __obf_69d23e854bc7930e
		}
	}

	if __obf_8733059f76088ffc == __obf_f1f3ff99eed138e6 {
		return __obf_56d6bcce6d349920
	}

	switch __obf_5cd51d276078fe9d {
	case reflect.Ptr:
		return __obf_7a752c9836a9ee1e(__obf_8733059f76088ffc)
	case reflect.Slice:
		__obf_d6cb414a13d9d0f7 := __obf_8733059f76088ffc.Elem()
		if __obf_d6cb414a13d9d0f7.Kind() == reflect.Uint8 {
			return __obf_898cd0032343e4d6
		}
		if __obf_d6cb414a13d9d0f7 == __obf_e7498c925721ef3a {
			return __obf_1d2042271514bd6b
		}
	case reflect.Array:
		if __obf_8733059f76088ffc.Elem().Kind() == reflect.Uint8 {
			return __obf_66038bf3c325450f
		}
	case reflect.Map:
		if __obf_8733059f76088ffc.Key() == __obf_e7498c925721ef3a {
			switch __obf_8733059f76088ffc.Elem() {
			case __obf_e7498c925721ef3a:
				return __obf_617b6f6590f5807c
			case __obf_a9624631014b694f:
				return __obf_e92628670d4494a8
			case __obf_533b257584aace5f:
				return __obf_d3a88247514432a8
			}
		}
	}

	return __obf_dfbdb30a1aae87b2[__obf_5cd51d276078fe9d]
}

func __obf_7a752c9836a9ee1e(__obf_8733059f76088ffc reflect.Type) __obf_dc120d370e164c91 {
	__obf_ff15dfff9c06a391 := __obf_031e3a13e30d95dc(__obf_8733059f76088ffc.Elem())
	return func(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
		if __obf_f328a048f76b7256.IsNil() {
			return __obf_e9038cda3b5cf711.EncodeNil()
		}
		return __obf_ff15dfff9c06a391(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256.Elem())
	}
}

func __obf_30978ab67c376dcd(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if !__obf_f328a048f76b7256.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_f328a048f76b7256.Interface())
	}
	__obf_ff15dfff9c06a391 := __obf_f328a048f76b7256.Addr().Interface().(CustomEncoder)
	return __obf_ff15dfff9c06a391.EncodeMsgpack(__obf_e9038cda3b5cf711)
}

func __obf_f8f9ae3cc36e5f9d(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_066825d3e1771287(__obf_f328a048f76b7256.Kind()) && __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	__obf_ff15dfff9c06a391 := __obf_f328a048f76b7256.Interface().(CustomEncoder)
	return __obf_ff15dfff9c06a391.EncodeMsgpack(__obf_e9038cda3b5cf711)
}

func __obf_9e056da4627a0ac3(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if !__obf_f328a048f76b7256.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_f328a048f76b7256.Interface())
	}
	return __obf_83060f2b2cde521c(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256.Addr())
}

func __obf_83060f2b2cde521c(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_066825d3e1771287(__obf_f328a048f76b7256.Kind()) && __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	__obf_10d6f05961239208 := __obf_f328a048f76b7256.Interface().(Marshaler)
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_10d6f05961239208.MarshalMsgpack()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	_, __obf_45342a3a754d12f5 = __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.Write(__obf_94333af0f0facd65)
	return __obf_45342a3a754d12f5
}

func __obf_8af2108e5f5a9336(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeBool(__obf_f328a048f76b7256.Bool())
}

func __obf_3c71664b718cf2c4(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	return __obf_e9038cda3b5cf711.EncodeValue(__obf_f328a048f76b7256.Elem())
}

func __obf_56d6bcce6d349920(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	return __obf_e9038cda3b5cf711.EncodeString(__obf_f328a048f76b7256.Interface().(error).Error())
}

func __obf_d1d21523321fd9d0(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return fmt.Errorf("msgpack: Encode(unsupported %s)", __obf_f328a048f76b7256.Type())
}

func __obf_066825d3e1771287(__obf_5cd51d276078fe9d reflect.Kind) bool {
	switch __obf_5cd51d276078fe9d {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func __obf_cdb09347bdcac603(__obf_0eb9b9619deca594 reflect.Type) bool {
	if __obf_0eb9b9619deca594.Kind() == reflect.Ptr {
		__obf_0eb9b9619deca594 = __obf_0eb9b9619deca594.Elem()
	}
	return __obf_066825d3e1771287(__obf_0eb9b9619deca594.Kind())
}

//------------------------------------------------------------------------------

func __obf_08cf0d4863d94784(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if !__obf_f328a048f76b7256.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_f328a048f76b7256.Interface())
	}
	return __obf_03b1c3316fe255c4(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256.Addr())
}

func __obf_03b1c3316fe255c4(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_066825d3e1771287(__obf_f328a048f76b7256.Kind()) && __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	__obf_10d6f05961239208 := __obf_f328a048f76b7256.Interface().(encoding.BinaryMarshaler)
	__obf_cc76e8ed73142f28, __obf_45342a3a754d12f5 := __obf_10d6f05961239208.MarshalBinary()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	return __obf_e9038cda3b5cf711.EncodeBytes(__obf_cc76e8ed73142f28)
}

//------------------------------------------------------------------------------

func __obf_69d23e854bc7930e(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if !__obf_f328a048f76b7256.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_f328a048f76b7256.Interface())
	}
	return __obf_afe2be14b2072c5c(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256.Addr())
}

func __obf_afe2be14b2072c5c(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_066825d3e1771287(__obf_f328a048f76b7256.Kind()) && __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	__obf_10d6f05961239208 := __obf_f328a048f76b7256.Interface().(encoding.TextMarshaler)
	__obf_cc76e8ed73142f28, __obf_45342a3a754d12f5 := __obf_10d6f05961239208.MarshalText()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	return __obf_e9038cda3b5cf711.EncodeBytes(__obf_cc76e8ed73142f28)
}
