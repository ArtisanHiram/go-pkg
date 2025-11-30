package __obf_fd770cb9675903b0

import (
	"fmt"
	"math"
	"reflect"
)

type __obf_ae91b3f1af57dd05 struct {
	Type    reflect.Type
	Decoder func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value, __obf_803bccbf05dad29c int) error
}

var __obf_8adb5c6b5b689f26 = make(map[int8]*__obf_ae91b3f1af57dd05)

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

func RegisterExt(__obf_e728c4abb2dcc444 int8, __obf_28cbfc96ff0a56b6 MarshalerUnmarshaler) {
	RegisterExtEncoder(__obf_e728c4abb2dcc444, __obf_28cbfc96ff0a56b6, func(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) ([]byte, error) {
		__obf_10d6f05961239208 := __obf_f328a048f76b7256.Interface().(Marshaler)
		return __obf_10d6f05961239208.MarshalMsgpack()
	})
	RegisterExtDecoder(__obf_e728c4abb2dcc444, __obf_28cbfc96ff0a56b6, func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value, __obf_803bccbf05dad29c int) error {
		__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(__obf_803bccbf05dad29c)
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		return __obf_f328a048f76b7256.Interface().(Unmarshaler).UnmarshalMsgpack(__obf_94333af0f0facd65)
	})
}

func UnregisterExt(__obf_e728c4abb2dcc444 int8) {
	__obf_8d9dfbb85b2526a9(__obf_e728c4abb2dcc444)
	__obf_32c355ab911d27fc(__obf_e728c4abb2dcc444)
}

func RegisterExtEncoder(__obf_e728c4abb2dcc444 int8, __obf_28cbfc96ff0a56b6 any, __obf_ff15dfff9c06a391 func(__obf_a7b5a0ca650f48ee *Encoder, __obf_f328a048f76b7256 reflect.Value) ([]byte, error),
) {
	__obf_8d9dfbb85b2526a9(__obf_e728c4abb2dcc444)
	__obf_8733059f76088ffc := reflect.TypeOf(__obf_28cbfc96ff0a56b6)
	__obf_07575e084439607f := __obf_a94ad1ddc9ac81d9(__obf_e728c4abb2dcc444, __obf_8733059f76088ffc, __obf_ff15dfff9c06a391)
	__obf_56aff9039d970a50.
		Store(__obf_e728c4abb2dcc444, __obf_8733059f76088ffc)
	__obf_56aff9039d970a50.
		Store(__obf_8733059f76088ffc, __obf_07575e084439607f)
	if __obf_8733059f76088ffc.Kind() == reflect.Ptr {
		__obf_56aff9039d970a50.
			Store(__obf_8733059f76088ffc.Elem(), __obf_4bdf30f0240dfd45(__obf_07575e084439607f))
	}
}

func __obf_8d9dfbb85b2526a9(__obf_e728c4abb2dcc444 int8) {
	__obf_0eb9b9619deca594, __obf_b00b3c6a10f90467 := __obf_56aff9039d970a50.Load(__obf_e728c4abb2dcc444)
	if !__obf_b00b3c6a10f90467 {
		return
	}
	__obf_56aff9039d970a50.
		Delete(__obf_e728c4abb2dcc444)
	__obf_8733059f76088ffc := __obf_0eb9b9619deca594.(reflect.Type)
	__obf_56aff9039d970a50.
		Delete(__obf_8733059f76088ffc)
	if __obf_8733059f76088ffc.Kind() == reflect.Ptr {
		__obf_56aff9039d970a50.
			Delete(__obf_8733059f76088ffc.Elem())
	}
}

func __obf_a94ad1ddc9ac81d9(__obf_e728c4abb2dcc444 int8, __obf_8733059f76088ffc reflect.Type, __obf_ff15dfff9c06a391 func(__obf_a7b5a0ca650f48ee *Encoder, __obf_f328a048f76b7256 reflect.Value) ([]byte, error),
) __obf_dc120d370e164c91 {
	__obf_066825d3e1771287 := __obf_8733059f76088ffc.Kind() == reflect.Ptr

	return func(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
		if __obf_066825d3e1771287 && __obf_f328a048f76b7256.IsNil() {
			return __obf_e9038cda3b5cf711.EncodeNil()
		}
		__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_ff15dfff9c06a391(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256)
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}

		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeExtHeader(__obf_e728c4abb2dcc444, len(__obf_94333af0f0facd65)); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}

		return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_94333af0f0facd65)
	}
}

func __obf_4bdf30f0240dfd45(__obf_07575e084439607f __obf_dc120d370e164c91) __obf_dc120d370e164c91 {
	return func(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
		if !__obf_f328a048f76b7256.CanAddr() {
			return fmt.Errorf("msgpack: EncodeExt(nonaddressable %T)", __obf_f328a048f76b7256.Interface())
		}
		return __obf_07575e084439607f(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256.Addr())
	}
}

func RegisterExtDecoder(__obf_e728c4abb2dcc444 int8, __obf_28cbfc96ff0a56b6 any, __obf_25e78266d5ba1944 func(__obf_2f8f01810a02d049 *Decoder, __obf_f328a048f76b7256 reflect.Value, __obf_803bccbf05dad29c int) error,
) {
	__obf_32c355ab911d27fc(__obf_e728c4abb2dcc444)
	__obf_8733059f76088ffc := reflect.TypeOf(__obf_28cbfc96ff0a56b6)
	__obf_159a6ada5e3fef68 := __obf_6172e8d93973bb68(__obf_e728c4abb2dcc444, __obf_8733059f76088ffc, __obf_25e78266d5ba1944)
	__obf_8adb5c6b5b689f26[__obf_e728c4abb2dcc444] = &__obf_ae91b3f1af57dd05{
		Type:    __obf_8733059f76088ffc,
		Decoder: __obf_25e78266d5ba1944,
	}
	__obf_9ea14163dfc9e317.
		Store(__obf_e728c4abb2dcc444, __obf_8733059f76088ffc)
	__obf_9ea14163dfc9e317.
		Store(__obf_8733059f76088ffc, __obf_159a6ada5e3fef68)
	if __obf_8733059f76088ffc.Kind() == reflect.Ptr {
		__obf_9ea14163dfc9e317.
			Store(__obf_8733059f76088ffc.Elem(), __obf_6e7838ab34e309c9(__obf_159a6ada5e3fef68))
	}
}

func __obf_32c355ab911d27fc(__obf_e728c4abb2dcc444 int8) {
	__obf_0eb9b9619deca594, __obf_b00b3c6a10f90467 := __obf_9ea14163dfc9e317.Load(__obf_e728c4abb2dcc444)
	if !__obf_b00b3c6a10f90467 {
		return
	}
	__obf_9ea14163dfc9e317.
		Delete(__obf_e728c4abb2dcc444)
	delete(__obf_8adb5c6b5b689f26, __obf_e728c4abb2dcc444)
	__obf_8733059f76088ffc := __obf_0eb9b9619deca594.(reflect.Type)
	__obf_9ea14163dfc9e317.
		Delete(__obf_8733059f76088ffc)
	if __obf_8733059f76088ffc.Kind() == reflect.Ptr {
		__obf_9ea14163dfc9e317.
			Delete(__obf_8733059f76088ffc.Elem())
	}
}

func __obf_6172e8d93973bb68(__obf_a763e05470377242 int8, __obf_8733059f76088ffc reflect.Type, __obf_25e78266d5ba1944 func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value, __obf_803bccbf05dad29c int) error,
) __obf_fb10c6a4667e72a5 {
	return __obf_ca3ac1e2410413b5(__obf_8733059f76088ffc, func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
		__obf_e728c4abb2dcc444, __obf_803bccbf05dad29c, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeExtHeader()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		if __obf_e728c4abb2dcc444 != __obf_a763e05470377242 {
			return fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_e728c4abb2dcc444, __obf_a763e05470377242)
		}
		return __obf_25e78266d5ba1944(__obf_5d389b660eefb08c, __obf_f328a048f76b7256, __obf_803bccbf05dad29c)
	})
}

func __obf_6e7838ab34e309c9(__obf_159a6ada5e3fef68 __obf_fb10c6a4667e72a5) __obf_fb10c6a4667e72a5 {
	return func(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
		if !__obf_f328a048f76b7256.CanAddr() {
			return fmt.Errorf("msgpack: DecodeExt(nonaddressable %T)", __obf_f328a048f76b7256.Interface())
		}
		return __obf_159a6ada5e3fef68(__obf_5d389b660eefb08c, __obf_f328a048f76b7256.Addr())
	}
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeExtHeader(__obf_e728c4abb2dcc444 int8, __obf_803bccbf05dad29c int) error {
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_5926ca009cd80dd3(__obf_803bccbf05dad29c); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.WriteByte(byte(__obf_e728c4abb2dcc444)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_5926ca009cd80dd3(__obf_bddea2836c583aa2 int) error {
	switch __obf_bddea2836c583aa2 {
	case 1:
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt1)
	case 2:
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt2)
	case 4:
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt4)
	case 8:
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt8)
	case 16:
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixExt16)
	}
	if __obf_bddea2836c583aa2 <= math.MaxUint8 {
		return __obf_e9038cda3b5cf711.__obf_c58302e6ea7043c4(Ext8, uint8(__obf_bddea2836c583aa2))
	}
	if __obf_bddea2836c583aa2 <= math.MaxUint16 {
		return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Ext16, uint16(__obf_bddea2836c583aa2))
	}
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Ext32, uint32(__obf_bddea2836c583aa2))
}

func (__obf_5d389b660eefb08c *Decoder) DecodeExtHeader() (__obf_e728c4abb2dcc444 int8, __obf_803bccbf05dad29c int, __obf_45342a3a754d12f5 error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return
	}
	return __obf_5d389b660eefb08c.__obf_051bbe645d4e1aa5(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_051bbe645d4e1aa5(__obf_4148125b350dfea2 byte) (int8, int, error) {
	__obf_803bccbf05dad29c, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4b7d05f03050f9ec(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return 0, 0, __obf_45342a3a754d12f5
	}
	__obf_e728c4abb2dcc444, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, 0, __obf_45342a3a754d12f5
	}

	return int8(__obf_e728c4abb2dcc444), __obf_803bccbf05dad29c, nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_4b7d05f03050f9ec(__obf_4148125b350dfea2 byte) (int, error) {
	switch __obf_4148125b350dfea2 {
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
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Ext16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Ext32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	default:
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", __obf_4148125b350dfea2)
	}
}

func (__obf_5d389b660eefb08c *Decoder) __obf_fae8700b6c50321a(__obf_4148125b350dfea2 byte) (any, error) {
	__obf_e728c4abb2dcc444, __obf_803bccbf05dad29c, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_051bbe645d4e1aa5(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	__obf_4847e81f3ea28355, __obf_b00b3c6a10f90467 := __obf_8adb5c6b5b689f26[__obf_e728c4abb2dcc444]
	if !__obf_b00b3c6a10f90467 {
		return nil, fmt.Errorf("msgpack: unknown ext id=%d", __obf_e728c4abb2dcc444)
	}
	__obf_f328a048f76b7256 := __obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_4847e81f3ea28355.Type).Elem()
	if __obf_066825d3e1771287(__obf_f328a048f76b7256.Kind()) && __obf_f328a048f76b7256.IsNil() {
		__obf_f328a048f76b7256.
			Set(__obf_5d389b660eefb08c.__obf_9eb3ea28cca5995b(__obf_4847e81f3ea28355.Type.Elem()))
	}

	if __obf_45342a3a754d12f5 := __obf_4847e81f3ea28355.Decoder(__obf_5d389b660eefb08c, __obf_f328a048f76b7256, __obf_803bccbf05dad29c); __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}

	return __obf_f328a048f76b7256.Interface(), nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_d955aa85bcfecb15(__obf_4148125b350dfea2 byte) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4b7d05f03050f9ec(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_d93e03df64d057a0(__obf_d774e4844c74bfe9 + 1)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_57d324efa1b98131(__obf_4148125b350dfea2 byte) error {
	// Read ext type.
	_, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	// Read ext body len.
	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_ff144082e76a9d3b(__obf_4148125b350dfea2); __obf_4140b3ff04f75a36++ {
		_, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}

func __obf_ff144082e76a9d3b(__obf_4148125b350dfea2 byte) int {
	switch __obf_4148125b350dfea2 {
	case Ext8:
		return 1
	case Ext16:
		return 2
	case Ext32:
		return 4
	}
	return 0
}
