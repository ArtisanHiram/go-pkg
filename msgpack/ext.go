package __obf_a4aad98aaf3178f4

import (
	"fmt"
	"math"
	"reflect"
)

type __obf_3346ed3406077a0b struct {
	Type    reflect.Type
	Decoder func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value, __obf_f975a2b18cb9c21c int) error
}

var __obf_b394b4262313f390 = make(map[int8]*__obf_3346ed3406077a0b)

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

func RegisterExt(__obf_2f505346eb973ca1 int8, __obf_9055e1481edcd436 MarshalerUnmarshaler) {
	RegisterExtEncoder(__obf_2f505346eb973ca1, __obf_9055e1481edcd436, func(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) ([]byte, error) {
		__obf_f68fedb43a0f697e := __obf_a1f43267eeb48a1a.Interface().(Marshaler)
		return __obf_f68fedb43a0f697e.MarshalMsgpack()
	})
	RegisterExtDecoder(__obf_2f505346eb973ca1, __obf_9055e1481edcd436, func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value, __obf_f975a2b18cb9c21c int) error {
		__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_4429146a43365de4(__obf_f975a2b18cb9c21c)
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		return __obf_a1f43267eeb48a1a.Interface().(Unmarshaler).UnmarshalMsgpack(__obf_f57209cfda8e17cf)
	})
}

func UnregisterExt(__obf_2f505346eb973ca1 int8) {
	__obf_f494d5150c84253a(__obf_2f505346eb973ca1)
	__obf_834d936b57ff70ba(__obf_2f505346eb973ca1)
}

func RegisterExtEncoder(__obf_2f505346eb973ca1 int8, __obf_9055e1481edcd436 any, __obf_c4f748de4fd652e4 func(__obf_dfba1ffd920163bc *Encoder, __obf_a1f43267eeb48a1a reflect.Value) ([]byte, error),
) {
	__obf_f494d5150c84253a(__obf_2f505346eb973ca1)
	__obf_bbfd30fcc08dc1bc := reflect.TypeOf(__obf_9055e1481edcd436)
	__obf_a1720eb9fcc587cb := __obf_cd46eccbefa4d337(__obf_2f505346eb973ca1, __obf_bbfd30fcc08dc1bc, __obf_c4f748de4fd652e4)
	__obf_a25a843d85821e5c.
		Store(__obf_2f505346eb973ca1, __obf_bbfd30fcc08dc1bc)
	__obf_a25a843d85821e5c.
		Store(__obf_bbfd30fcc08dc1bc, __obf_a1720eb9fcc587cb)
	if __obf_bbfd30fcc08dc1bc.Kind() == reflect.Ptr {
		__obf_a25a843d85821e5c.
			Store(__obf_bbfd30fcc08dc1bc.Elem(), __obf_6fc16d6c245a502a(__obf_a1720eb9fcc587cb))
	}
}

func __obf_f494d5150c84253a(__obf_2f505346eb973ca1 int8) {
	__obf_8698f5d73996ef26, __obf_81b19f2a6159ab89 := __obf_a25a843d85821e5c.Load(__obf_2f505346eb973ca1)
	if !__obf_81b19f2a6159ab89 {
		return
	}
	__obf_a25a843d85821e5c.
		Delete(__obf_2f505346eb973ca1)
	__obf_bbfd30fcc08dc1bc := __obf_8698f5d73996ef26.(reflect.Type)
	__obf_a25a843d85821e5c.
		Delete(__obf_bbfd30fcc08dc1bc)
	if __obf_bbfd30fcc08dc1bc.Kind() == reflect.Ptr {
		__obf_a25a843d85821e5c.
			Delete(__obf_bbfd30fcc08dc1bc.Elem())
	}
}

func __obf_cd46eccbefa4d337(__obf_2f505346eb973ca1 int8, __obf_bbfd30fcc08dc1bc reflect.Type, __obf_c4f748de4fd652e4 func(__obf_dfba1ffd920163bc *Encoder, __obf_a1f43267eeb48a1a reflect.Value) ([]byte, error),
) __obf_9d8023233a199a4f {
	__obf_178efdbad2797673 := __obf_bbfd30fcc08dc1bc.Kind() == reflect.Ptr

	return func(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if __obf_178efdbad2797673 && __obf_a1f43267eeb48a1a.IsNil() {
			return __obf_2c8e97779108ab17.EncodeNil()
		}
		__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_c4f748de4fd652e4(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a)
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}

		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.EncodeExtHeader(__obf_2f505346eb973ca1, len(__obf_f57209cfda8e17cf)); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}

		return __obf_2c8e97779108ab17.__obf_ab9f7a04111e0451(__obf_f57209cfda8e17cf)
	}
}

func __obf_6fc16d6c245a502a(__obf_a1720eb9fcc587cb __obf_9d8023233a199a4f) __obf_9d8023233a199a4f {
	return func(__obf_2c8e97779108ab17 *Encoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if !__obf_a1f43267eeb48a1a.CanAddr() {
			return fmt.Errorf("msgpack: EncodeExt(nonaddressable %T)", __obf_a1f43267eeb48a1a.Interface())
		}
		return __obf_a1720eb9fcc587cb(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a.Addr())
	}
}

func RegisterExtDecoder(__obf_2f505346eb973ca1 int8, __obf_9055e1481edcd436 any, __obf_815edc9179203b9a func(__obf_cc68ee77d25ca8f2 *Decoder, __obf_a1f43267eeb48a1a reflect.Value, __obf_f975a2b18cb9c21c int) error,
) {
	__obf_834d936b57ff70ba(__obf_2f505346eb973ca1)
	__obf_bbfd30fcc08dc1bc := reflect.TypeOf(__obf_9055e1481edcd436)
	__obf_d2b70e1bdd2533d1 := __obf_e09823b1718cc089(__obf_2f505346eb973ca1, __obf_bbfd30fcc08dc1bc, __obf_815edc9179203b9a)
	__obf_b394b4262313f390[__obf_2f505346eb973ca1] = &__obf_3346ed3406077a0b{
		Type:    __obf_bbfd30fcc08dc1bc,
		Decoder: __obf_815edc9179203b9a,
	}
	__obf_6a2362b607537b3d.
		Store(__obf_2f505346eb973ca1, __obf_bbfd30fcc08dc1bc)
	__obf_6a2362b607537b3d.
		Store(__obf_bbfd30fcc08dc1bc, __obf_d2b70e1bdd2533d1)
	if __obf_bbfd30fcc08dc1bc.Kind() == reflect.Ptr {
		__obf_6a2362b607537b3d.
			Store(__obf_bbfd30fcc08dc1bc.Elem(), __obf_3cd07572ecf064e3(__obf_d2b70e1bdd2533d1))
	}
}

func __obf_834d936b57ff70ba(__obf_2f505346eb973ca1 int8) {
	__obf_8698f5d73996ef26, __obf_81b19f2a6159ab89 := __obf_6a2362b607537b3d.Load(__obf_2f505346eb973ca1)
	if !__obf_81b19f2a6159ab89 {
		return
	}
	__obf_6a2362b607537b3d.
		Delete(__obf_2f505346eb973ca1)
	delete(__obf_b394b4262313f390, __obf_2f505346eb973ca1)
	__obf_bbfd30fcc08dc1bc := __obf_8698f5d73996ef26.(reflect.Type)
	__obf_6a2362b607537b3d.
		Delete(__obf_bbfd30fcc08dc1bc)
	if __obf_bbfd30fcc08dc1bc.Kind() == reflect.Ptr {
		__obf_6a2362b607537b3d.
			Delete(__obf_bbfd30fcc08dc1bc.Elem())
	}
}

func __obf_e09823b1718cc089(__obf_d4ca28cbbd5239fb int8, __obf_bbfd30fcc08dc1bc reflect.Type, __obf_815edc9179203b9a func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value, __obf_f975a2b18cb9c21c int) error,
) __obf_c1c8351491cde4cc {
	return __obf_e5dd8d0f0e7e2933(__obf_bbfd30fcc08dc1bc, func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		__obf_2f505346eb973ca1, __obf_f975a2b18cb9c21c, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeExtHeader()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		if __obf_2f505346eb973ca1 != __obf_d4ca28cbbd5239fb {
			return fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_2f505346eb973ca1, __obf_d4ca28cbbd5239fb)
		}
		return __obf_815edc9179203b9a(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a, __obf_f975a2b18cb9c21c)
	})
}

func __obf_3cd07572ecf064e3(__obf_d2b70e1bdd2533d1 __obf_c1c8351491cde4cc) __obf_c1c8351491cde4cc {
	return func(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
		if !__obf_a1f43267eeb48a1a.CanAddr() {
			return fmt.Errorf("msgpack: DecodeExt(nonaddressable %T)", __obf_a1f43267eeb48a1a.Interface())
		}
		return __obf_d2b70e1bdd2533d1(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a.Addr())
	}
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeExtHeader(__obf_2f505346eb973ca1 int8, __obf_f975a2b18cb9c21c int) error {
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_d1f2c919354a65de(__obf_f975a2b18cb9c21c); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.WriteByte(byte(__obf_2f505346eb973ca1)); __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_d1f2c919354a65de(__obf_4431cbe3c2f63394 int) error {
	switch __obf_4431cbe3c2f63394 {
	case 1:
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt1)
	case 2:
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt2)
	case 4:
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt4)
	case 8:
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt8)
	case 16:
		return __obf_2c8e97779108ab17.__obf_c087409347698442(FixExt16)
	}
	if __obf_4431cbe3c2f63394 <= math.MaxUint8 {
		return __obf_2c8e97779108ab17.__obf_2ebb4be6da23dcc7(Ext8, uint8(__obf_4431cbe3c2f63394))
	}
	if __obf_4431cbe3c2f63394 <= math.MaxUint16 {
		return __obf_2c8e97779108ab17.__obf_71321d3206264f3b(Ext16, uint16(__obf_4431cbe3c2f63394))
	}
	return __obf_2c8e97779108ab17.__obf_a40df49dbc3bfc37(Ext32, uint32(__obf_4431cbe3c2f63394))
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeExtHeader() (__obf_2f505346eb973ca1 int8, __obf_f975a2b18cb9c21c int, __obf_4061ca0039150c39 error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return
	}
	return __obf_613397fefdec0ed0.__obf_2227ad80d0435996(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_2227ad80d0435996(__obf_6dbe86b3d9d9b037 byte) (int8, int, error) {
	__obf_f975a2b18cb9c21c, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0cc0871e880ee750(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return 0, 0, __obf_4061ca0039150c39
	}
	__obf_2f505346eb973ca1, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, 0, __obf_4061ca0039150c39
	}

	return int8(__obf_2f505346eb973ca1), __obf_f975a2b18cb9c21c, nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_0cc0871e880ee750(__obf_6dbe86b3d9d9b037 byte) (int, error) {
	switch __obf_6dbe86b3d9d9b037 {
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
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint8()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Ext16:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint16()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	case Ext32:
		__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.uint32()
		return int(__obf_99a74e41c9c640c0), __obf_4061ca0039150c39
	default:
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", __obf_6dbe86b3d9d9b037)
	}
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_478bdbc02e8c61ec(__obf_6dbe86b3d9d9b037 byte) (any, error) {
	__obf_2f505346eb973ca1, __obf_f975a2b18cb9c21c, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_2227ad80d0435996(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	__obf_76a17a04fd62dc3b, __obf_81b19f2a6159ab89 := __obf_b394b4262313f390[__obf_2f505346eb973ca1]
	if !__obf_81b19f2a6159ab89 {
		return nil, fmt.Errorf("msgpack: unknown ext id=%d", __obf_2f505346eb973ca1)
	}
	__obf_a1f43267eeb48a1a := __obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_76a17a04fd62dc3b.Type).Elem()
	if __obf_178efdbad2797673(__obf_a1f43267eeb48a1a.Kind()) && __obf_a1f43267eeb48a1a.IsNil() {
		__obf_a1f43267eeb48a1a.
			Set(__obf_613397fefdec0ed0.__obf_7a63f73d6b3d6972(__obf_76a17a04fd62dc3b.Type.Elem()))
	}

	if __obf_4061ca0039150c39 := __obf_76a17a04fd62dc3b.Decoder(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a, __obf_f975a2b18cb9c21c); __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}

	return __obf_a1f43267eeb48a1a.Interface(), nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_a07e0b14e7edfb58(__obf_6dbe86b3d9d9b037 byte) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0cc0871e880ee750(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_2bfec4fb91ed09f6(__obf_99a74e41c9c640c0 + 1)
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_0428b2f11094e5f5(__obf_6dbe86b3d9d9b037 byte) error {
	// Read ext type.
	_, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	// Read ext body len.
	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_08018dfe02cd99a6(__obf_6dbe86b3d9d9b037); __obf_097d8b6061c9592b++ {
		_, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}

func __obf_08018dfe02cd99a6(__obf_6dbe86b3d9d9b037 byte) int {
	switch __obf_6dbe86b3d9d9b037 {
	case Ext8:
		return 1
	case Ext16:
		return 2
	case Ext32:
		return 4
	}
	return 0
}
