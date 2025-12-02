package __obf_a935eb7f548271a4

import (
	"fmt"
	"math"
	"reflect"
)

type __obf_9b64c962e55dc648 struct {
	Type    reflect.Type
	Decoder func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value, __obf_f9e21d535a3562ea int) error
}

var __obf_335898695997d965 = make(map[int8]*__obf_9b64c962e55dc648)

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

func RegisterExt(__obf_12dd79d80139d0ca int8, __obf_205ee5d820f118f1 MarshalerUnmarshaler) {
	RegisterExtEncoder(__obf_12dd79d80139d0ca, __obf_205ee5d820f118f1, func(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) ([]byte, error) {
		__obf_ccececcb28726e85 := __obf_6d570581f4b60dbc.Interface().(Marshaler)
		return __obf_ccececcb28726e85.MarshalMsgpack()
	})
	RegisterExtDecoder(__obf_12dd79d80139d0ca, __obf_205ee5d820f118f1, func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value, __obf_f9e21d535a3562ea int) error {
		__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_595cbf51b6653ebf(__obf_f9e21d535a3562ea)
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		return __obf_6d570581f4b60dbc.Interface().(Unmarshaler).UnmarshalMsgpack(__obf_f2ca794293605b73)
	})
}

func UnregisterExt(__obf_12dd79d80139d0ca int8) {
	__obf_5bc13467310d85c9(__obf_12dd79d80139d0ca)
	__obf_d1740bde7a6fca19(__obf_12dd79d80139d0ca)
}

func RegisterExtEncoder(__obf_12dd79d80139d0ca int8, __obf_205ee5d820f118f1 any, __obf_6a114719334e34d2 func(__obf_6ca0309463018ed3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) ([]byte, error),
) {
	__obf_5bc13467310d85c9(__obf_12dd79d80139d0ca)
	__obf_7bd64e8df8357cab := reflect.TypeOf(__obf_205ee5d820f118f1)
	__obf_a334b5e1ac01ab78 := __obf_835ab14689c00f59(__obf_12dd79d80139d0ca, __obf_7bd64e8df8357cab, __obf_6a114719334e34d2)
	__obf_a16fc4f2eeccec6d.
		Store(__obf_12dd79d80139d0ca, __obf_7bd64e8df8357cab)
	__obf_a16fc4f2eeccec6d.
		Store(__obf_7bd64e8df8357cab, __obf_a334b5e1ac01ab78)
	if __obf_7bd64e8df8357cab.Kind() == reflect.Ptr {
		__obf_a16fc4f2eeccec6d.
			Store(__obf_7bd64e8df8357cab.Elem(), __obf_8b0d16b3c322b730(__obf_a334b5e1ac01ab78))
	}
}

func __obf_5bc13467310d85c9(__obf_12dd79d80139d0ca int8) {
	__obf_14c99752d39cbe18, __obf_826ac2dbb957d7df := __obf_a16fc4f2eeccec6d.Load(__obf_12dd79d80139d0ca)
	if !__obf_826ac2dbb957d7df {
		return
	}
	__obf_a16fc4f2eeccec6d.
		Delete(__obf_12dd79d80139d0ca)
	__obf_7bd64e8df8357cab := __obf_14c99752d39cbe18.(reflect.Type)
	__obf_a16fc4f2eeccec6d.
		Delete(__obf_7bd64e8df8357cab)
	if __obf_7bd64e8df8357cab.Kind() == reflect.Ptr {
		__obf_a16fc4f2eeccec6d.
			Delete(__obf_7bd64e8df8357cab.Elem())
	}
}

func __obf_835ab14689c00f59(__obf_12dd79d80139d0ca int8, __obf_7bd64e8df8357cab reflect.Type, __obf_6a114719334e34d2 func(__obf_6ca0309463018ed3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) ([]byte, error),
) __obf_4982458b85add976 {
	__obf_8b51c57d6f65dcf5 := __obf_7bd64e8df8357cab.Kind() == reflect.Ptr

	return func(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if __obf_8b51c57d6f65dcf5 && __obf_6d570581f4b60dbc.IsNil() {
			return __obf_ed37a34c347049f3.EncodeNil()
		}
		__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_6a114719334e34d2(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc)
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}

		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeExtHeader(__obf_12dd79d80139d0ca, len(__obf_f2ca794293605b73)); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}

		return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_f2ca794293605b73)
	}
}

func __obf_8b0d16b3c322b730(__obf_a334b5e1ac01ab78 __obf_4982458b85add976) __obf_4982458b85add976 {
	return func(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if !__obf_6d570581f4b60dbc.CanAddr() {
			return fmt.Errorf("msgpack: EncodeExt(nonaddressable %T)", __obf_6d570581f4b60dbc.Interface())
		}
		return __obf_a334b5e1ac01ab78(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc.Addr())
	}
}

func RegisterExtDecoder(__obf_12dd79d80139d0ca int8, __obf_205ee5d820f118f1 any, __obf_a17f840f59d84bbc func(__obf_fbe570faa7707c02 *Decoder, __obf_6d570581f4b60dbc reflect.Value, __obf_f9e21d535a3562ea int) error,
) {
	__obf_d1740bde7a6fca19(__obf_12dd79d80139d0ca)
	__obf_7bd64e8df8357cab := reflect.TypeOf(__obf_205ee5d820f118f1)
	__obf_ffa55e9478d7248a := __obf_4aa115dfb75bc0a9(__obf_12dd79d80139d0ca, __obf_7bd64e8df8357cab, __obf_a17f840f59d84bbc)
	__obf_335898695997d965[__obf_12dd79d80139d0ca] = &__obf_9b64c962e55dc648{
		Type:    __obf_7bd64e8df8357cab,
		Decoder: __obf_a17f840f59d84bbc,
	}
	__obf_48ad70830e899de5.
		Store(__obf_12dd79d80139d0ca, __obf_7bd64e8df8357cab)
	__obf_48ad70830e899de5.
		Store(__obf_7bd64e8df8357cab, __obf_ffa55e9478d7248a)
	if __obf_7bd64e8df8357cab.Kind() == reflect.Ptr {
		__obf_48ad70830e899de5.
			Store(__obf_7bd64e8df8357cab.Elem(), __obf_12a5c0262dd0ccfd(__obf_ffa55e9478d7248a))
	}
}

func __obf_d1740bde7a6fca19(__obf_12dd79d80139d0ca int8) {
	__obf_14c99752d39cbe18, __obf_826ac2dbb957d7df := __obf_48ad70830e899de5.Load(__obf_12dd79d80139d0ca)
	if !__obf_826ac2dbb957d7df {
		return
	}
	__obf_48ad70830e899de5.
		Delete(__obf_12dd79d80139d0ca)
	delete(__obf_335898695997d965, __obf_12dd79d80139d0ca)
	__obf_7bd64e8df8357cab := __obf_14c99752d39cbe18.(reflect.Type)
	__obf_48ad70830e899de5.
		Delete(__obf_7bd64e8df8357cab)
	if __obf_7bd64e8df8357cab.Kind() == reflect.Ptr {
		__obf_48ad70830e899de5.
			Delete(__obf_7bd64e8df8357cab.Elem())
	}
}

func __obf_4aa115dfb75bc0a9(__obf_e198bc71948320d1 int8, __obf_7bd64e8df8357cab reflect.Type, __obf_a17f840f59d84bbc func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value, __obf_f9e21d535a3562ea int) error,
) __obf_8e00c4b112264dba {
	return __obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
		__obf_12dd79d80139d0ca, __obf_f9e21d535a3562ea, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeExtHeader()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		if __obf_12dd79d80139d0ca != __obf_e198bc71948320d1 {
			return fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_12dd79d80139d0ca, __obf_e198bc71948320d1)
		}
		return __obf_a17f840f59d84bbc(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc, __obf_f9e21d535a3562ea)
	})
}

func __obf_12a5c0262dd0ccfd(__obf_ffa55e9478d7248a __obf_8e00c4b112264dba) __obf_8e00c4b112264dba {
	return func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if !__obf_6d570581f4b60dbc.CanAddr() {
			return fmt.Errorf("msgpack: DecodeExt(nonaddressable %T)", __obf_6d570581f4b60dbc.Interface())
		}
		return __obf_ffa55e9478d7248a(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc.Addr())
	}
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeExtHeader(__obf_12dd79d80139d0ca int8, __obf_f9e21d535a3562ea int) error {
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_e3e9e2d425e2405b(__obf_f9e21d535a3562ea); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.WriteByte(byte(__obf_12dd79d80139d0ca)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_e3e9e2d425e2405b(__obf_3c213c92f0597e4d int) error {
	switch __obf_3c213c92f0597e4d {
	case 1:
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt1)
	case 2:
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt2)
	case 4:
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt4)
	case 8:
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt8)
	case 16:
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixExt16)
	}
	if __obf_3c213c92f0597e4d <= math.MaxUint8 {
		return __obf_ed37a34c347049f3.__obf_327ee6f3f671bfd3(Ext8, uint8(__obf_3c213c92f0597e4d))
	}
	if __obf_3c213c92f0597e4d <= math.MaxUint16 {
		return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Ext16, uint16(__obf_3c213c92f0597e4d))
	}
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Ext32, uint32(__obf_3c213c92f0597e4d))
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeExtHeader() (__obf_12dd79d80139d0ca int8, __obf_f9e21d535a3562ea int, __obf_4d327e1cd40c2e21 error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return
	}
	return __obf_a21885da2425f2b2.__obf_4dc04de1e73ebdac(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_4dc04de1e73ebdac(__obf_f5df560f4d67421b byte) (int8, int, error) {
	__obf_f9e21d535a3562ea, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_046c8402b7a8e825(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, 0, __obf_4d327e1cd40c2e21
	}
	__obf_12dd79d80139d0ca, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, 0, __obf_4d327e1cd40c2e21
	}

	return int8(__obf_12dd79d80139d0ca), __obf_f9e21d535a3562ea, nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_046c8402b7a8e825(__obf_f5df560f4d67421b byte) (int, error) {
	switch __obf_f5df560f4d67421b {
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
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint8()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Ext16:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint16()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	case Ext32:
		__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.uint32()
		return int(__obf_326af9bd942662ac), __obf_4d327e1cd40c2e21
	default:
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", __obf_f5df560f4d67421b)
	}
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_f93129cdacf36826(__obf_f5df560f4d67421b byte) (any, error) {
	__obf_12dd79d80139d0ca, __obf_f9e21d535a3562ea, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_4dc04de1e73ebdac(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	__obf_f7e3b1463f090cdc, __obf_826ac2dbb957d7df := __obf_335898695997d965[__obf_12dd79d80139d0ca]
	if !__obf_826ac2dbb957d7df {
		return nil, fmt.Errorf("msgpack: unknown ext id=%d", __obf_12dd79d80139d0ca)
	}
	__obf_6d570581f4b60dbc := __obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_f7e3b1463f090cdc.Type).Elem()
	if __obf_8b51c57d6f65dcf5(__obf_6d570581f4b60dbc.Kind()) && __obf_6d570581f4b60dbc.IsNil() {
		__obf_6d570581f4b60dbc.
			Set(__obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_f7e3b1463f090cdc.Type.Elem()))
	}

	if __obf_4d327e1cd40c2e21 := __obf_f7e3b1463f090cdc.Decoder(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc, __obf_f9e21d535a3562ea); __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}

	return __obf_6d570581f4b60dbc.Interface(), nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_fa04fccd68154ccb(__obf_f5df560f4d67421b byte) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_046c8402b7a8e825(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_b2541a0cb78c8e1f(__obf_326af9bd942662ac + 1)
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_b5db8a2c243ae15f(__obf_f5df560f4d67421b byte) error {
	// Read ext type.
	_, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	// Read ext body len.
	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_10aecf4d518a742e(__obf_f5df560f4d67421b); __obf_4a8e280ffaa52cf4++ {
		_, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}

func __obf_10aecf4d518a742e(__obf_f5df560f4d67421b byte) int {
	switch __obf_f5df560f4d67421b {
	case Ext8:
		return 1
	case Ext16:
		return 2
	case Ext32:
		return 4
	}
	return 0
}
