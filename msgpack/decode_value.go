package __obf_a4aad98aaf3178f4

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
)

var (
	__obf_3eea2b4982d29baa = reflect.TypeOf((*any)(nil)).Elem()
	__obf_60ac326f01b0b63a = reflect.TypeOf((*string)(nil)).Elem()
	__obf_5b19c38666a256ef = reflect.TypeOf((*bool)(nil)).Elem()
)

var __obf_89ac0afae69f93fe []__obf_c1c8351491cde4cc

//nolint:gochecknoinits
func init() {
	__obf_89ac0afae69f93fe = []__obf_c1c8351491cde4cc{
		reflect.Bool:          __obf_a50e59d6cfa3ae6c,
		reflect.Int:           __obf_ee0f5bc124a2f6b4,
		reflect.Int8:          __obf_ee0f5bc124a2f6b4,
		reflect.Int16:         __obf_ee0f5bc124a2f6b4,
		reflect.Int32:         __obf_ee0f5bc124a2f6b4,
		reflect.Int64:         __obf_ee0f5bc124a2f6b4,
		reflect.Uint:          __obf_4a99d79a1164a943,
		reflect.Uint8:         __obf_4a99d79a1164a943,
		reflect.Uint16:        __obf_4a99d79a1164a943,
		reflect.Uint32:        __obf_4a99d79a1164a943,
		reflect.Uint64:        __obf_4a99d79a1164a943,
		reflect.Float32:       __obf_60142f5a2be96546,
		reflect.Float64:       __obf_a14f56eca4c741e3,
		reflect.Complex64:     __obf_0979e87d6e5a231f,
		reflect.Complex128:    __obf_0979e87d6e5a231f,
		reflect.Array:         __obf_7e0a77c3faa7dfde,
		reflect.Chan:          __obf_0979e87d6e5a231f,
		reflect.Func:          __obf_0979e87d6e5a231f,
		reflect.Interface:     __obf_3b5214be8509686e,
		reflect.Map:           __obf_c97ee5b3b53fb2c5,
		reflect.Pointer:       __obf_0979e87d6e5a231f,
		reflect.Slice:         __obf_6bcf41fd13df3eb0,
		reflect.String:        __obf_8b084829819fd8e7,
		reflect.Struct:        __obf_cf4307adf926d203,
		reflect.UnsafePointer: __obf_0979e87d6e5a231f,
	}
}

func __obf_66268eb3a3deccf5(__obf_bbfd30fcc08dc1bc reflect.Type) __obf_c1c8351491cde4cc {
	if __obf_a1f43267eeb48a1a, __obf_81b19f2a6159ab89 := __obf_6a2362b607537b3d.Load(__obf_bbfd30fcc08dc1bc); __obf_81b19f2a6159ab89 {
		return __obf_a1f43267eeb48a1a.(__obf_c1c8351491cde4cc)
	}
	__obf_64fe495dbc0c543b := _getDecoder(__obf_bbfd30fcc08dc1bc)
	__obf_6a2362b607537b3d.
		Store(__obf_bbfd30fcc08dc1bc, __obf_64fe495dbc0c543b)
	return __obf_64fe495dbc0c543b
}

func _getDecoder(__obf_bbfd30fcc08dc1bc reflect.Type) __obf_c1c8351491cde4cc {
	__obf_91ae1c24bc584c3a := __obf_bbfd30fcc08dc1bc.Kind()

	if __obf_91ae1c24bc584c3a == reflect.Ptr {
		if _, __obf_81b19f2a6159ab89 := __obf_6a2362b607537b3d.Load(__obf_bbfd30fcc08dc1bc.Elem()); __obf_81b19f2a6159ab89 {
			return __obf_84d2b53583a6c8d1(__obf_bbfd30fcc08dc1bc)
		}
	}

	if __obf_bbfd30fcc08dc1bc.Implements(__obf_283fd634becb1665) {
		return __obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_ac9d19a5f42e3609)
	}
	if __obf_bbfd30fcc08dc1bc.Implements(__obf_26665c3f2693777c) {
		return __obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_1f81a6417e7f3e72)
	}
	if __obf_bbfd30fcc08dc1bc.Implements(__obf_4da005c7e58bcab6) {
		return __obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_b15255074c7c3f0d)
	}
	if __obf_bbfd30fcc08dc1bc.Implements(__obf_1e6abd2dc26ca3b3) {
		return __obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_83bfe173315833be)
	}

	// Addressable struct field value.
	if __obf_91ae1c24bc584c3a != reflect.Pointer {
		__obf_cf3077802722ecc2 := reflect.PointerTo(__obf_bbfd30fcc08dc1bc)
		if __obf_cf3077802722ecc2.Implements(__obf_283fd634becb1665) {
			return __obf_3f6e2cd139237db6(__obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_ac9d19a5f42e3609))
		}
		if __obf_cf3077802722ecc2.Implements(__obf_26665c3f2693777c) {
			return __obf_3f6e2cd139237db6(__obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_1f81a6417e7f3e72))
		}
		if __obf_cf3077802722ecc2.Implements(__obf_4da005c7e58bcab6) {
			return __obf_3f6e2cd139237db6(__obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_b15255074c7c3f0d))
		}
		if __obf_cf3077802722ecc2.Implements(__obf_1e6abd2dc26ca3b3) {
			return __obf_3f6e2cd139237db6(__obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, __obf_83bfe173315833be))
		}
	}

	switch __obf_91ae1c24bc584c3a {
	case reflect.Pointer:
		return __obf_84d2b53583a6c8d1(__obf_bbfd30fcc08dc1bc)
	case reflect.Slice:
		__obf_ed47b1177873a7da := __obf_bbfd30fcc08dc1bc.Elem()
		if __obf_ed47b1177873a7da.Kind() == reflect.Uint8 {
			return __obf_9c86656b7d9d5cc1
		}
		if __obf_ed47b1177873a7da == __obf_60ac326f01b0b63a {
			return __obf_89bc4f0b771b8379
		}
	case reflect.Array:
		if __obf_bbfd30fcc08dc1bc.Elem().Kind() == reflect.Uint8 {
			return __obf_5091976ad4bd5009
		}
	case reflect.Map:
		if __obf_bbfd30fcc08dc1bc.Key() == __obf_60ac326f01b0b63a {
			switch __obf_bbfd30fcc08dc1bc.Elem() {
			case __obf_60ac326f01b0b63a:
				return __obf_cc88c63d6bb28c3c
			case __obf_3eea2b4982d29baa:
				return __obf_a5e827593d54b583
			}
		}
	}

	return __obf_89ac0afae69f93fe[__obf_91ae1c24bc584c3a]
}

func __obf_84d2b53583a6c8d1(__obf_bbfd30fcc08dc1bc reflect.Type) __obf_c1c8351491cde4cc {
	__obf_815edc9179203b9a := __obf_66268eb3a3deccf5(__obf_bbfd30fcc08dc1bc.Elem())
	return func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if __obf_613397fefdec0ed0.__obf_d6d4a347783e80b3() {
			if !__obf_a1f43267eeb48a1a.IsNil() {
				__obf_a1f43267eeb48a1a.
					Set(__obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_bbfd30fcc08dc1bc).Elem())
			}
			return __obf_613397fefdec0ed0.DecodeNil()
		}
		if __obf_a1f43267eeb48a1a.IsNil() {
			__obf_a1f43267eeb48a1a.
				Set(__obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_bbfd30fcc08dc1bc.Elem()))
		}
		return __obf_815edc9179203b9a(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a.Elem())
	}
}

func __obf_3f6e2cd139237db6(__obf_64fe495dbc0c543b __obf_c1c8351491cde4cc) __obf_c1c8351491cde4cc {
	return func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if !__obf_a1f43267eeb48a1a.CanAddr() {
			return fmt.Errorf("msgpack: Decode(nonaddressable %T)", __obf_a1f43267eeb48a1a.Interface())
		}
		return __obf_64fe495dbc0c543b(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a.Addr())
	}
}

func __obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc reflect.Type, __obf_64fe495dbc0c543b __obf_c1c8351491cde4cc) __obf_c1c8351491cde4cc {
	if __obf_178efdbad2797673(__obf_bbfd30fcc08dc1bc.Kind()) {
		return func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
			if __obf_613397fefdec0ed0.__obf_d6d4a347783e80b3() {
				return __obf_613397fefdec0ed0.__obf_82c6f7a4267374a0(__obf_a1f43267eeb48a1a)
			}
			if __obf_a1f43267eeb48a1a.IsNil() {
				__obf_a1f43267eeb48a1a.
					Set(__obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_bbfd30fcc08dc1bc.Elem()))
			}
			return __obf_64fe495dbc0c543b(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a)
		}
	}

	return func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if __obf_613397fefdec0ed0.__obf_d6d4a347783e80b3() {
			return __obf_613397fefdec0ed0.__obf_82c6f7a4267374a0(__obf_a1f43267eeb48a1a)
		}
		return __obf_64fe495dbc0c543b(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a)
	}
}

func __obf_a50e59d6cfa3ae6c(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_0104318e07e25055, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeBool()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetBool(__obf_0104318e07e25055)
	return nil
}

func __obf_3b5214be8509686e(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_613397fefdec0ed0.__obf_d947db23dba7519b(__obf_a1f43267eeb48a1a)
	}
	return __obf_613397fefdec0ed0.DecodeValue(__obf_a1f43267eeb48a1a.Elem())
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_d947db23dba7519b(__obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_9f1e21befe556e87, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	if __obf_9f1e21befe556e87 != nil {
		if __obf_a1f43267eeb48a1a.Type() == __obf_d66851dd29a0f414 {
			if __obf_9f1e21befe556e87, __obf_81b19f2a6159ab89 := __obf_9f1e21befe556e87.(string); __obf_81b19f2a6159ab89 {
				__obf_a1f43267eeb48a1a.
					Set(reflect.ValueOf(errors.New(__obf_9f1e21befe556e87)))
				return nil
			}
		}
		__obf_a1f43267eeb48a1a.
			Set(reflect.ValueOf(__obf_9f1e21befe556e87))
	}

	return nil
}

func __obf_0979e87d6e5a231f(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	return fmt.Errorf("msgpack: Decode(unsupported %s)", __obf_a1f43267eeb48a1a.Type())
}

//------------------------------------------------------------------------------

func __obf_ac9d19a5f42e3609(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_815edc9179203b9a := __obf_a1f43267eeb48a1a.Interface().(CustomDecoder)
	return __obf_815edc9179203b9a.DecodeMsgpack(__obf_613397fefdec0ed0)
}

func __obf_1f81a6417e7f3e72(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	var __obf_f57209cfda8e17cf []byte
	__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = make([]byte, 0, 64)
	if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_f57209cfda8e17cf = __obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1
	__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = nil
	__obf_1fb38a3d24203f78 := __obf_a1f43267eeb48a1a.Interface().(Unmarshaler)
	return __obf_1fb38a3d24203f78.UnmarshalMsgpack(__obf_f57209cfda8e17cf)
}

func __obf_b15255074c7c3f0d(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_69e4c283179273ce, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeBytes()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_1fb38a3d24203f78 := __obf_a1f43267eeb48a1a.Interface().(encoding.BinaryUnmarshaler)
	return __obf_1fb38a3d24203f78.UnmarshalBinary(__obf_69e4c283179273ce)
}

func __obf_83bfe173315833be(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_69e4c283179273ce, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeBytes()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_1fb38a3d24203f78 := __obf_a1f43267eeb48a1a.Interface().(encoding.TextUnmarshaler)
	return __obf_1fb38a3d24203f78.UnmarshalText(__obf_69e4c283179273ce)
}
