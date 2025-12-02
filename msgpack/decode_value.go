package __obf_a935eb7f548271a4

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
)

var (
	__obf_24d1b614db17eb2b = reflect.TypeOf((*any)(nil)).Elem()
	__obf_b128fe7c7d10a668 = reflect.TypeOf((*string)(nil)).Elem()
	__obf_8887b1926a35f0b7 = reflect.TypeOf((*bool)(nil)).Elem()
)

var __obf_bc021b46631f5774 []__obf_8e00c4b112264dba

//nolint:gochecknoinits
func init() {
	__obf_bc021b46631f5774 = []__obf_8e00c4b112264dba{
		reflect.Bool:          __obf_f6915865a4c0b583,
		reflect.Int:           __obf_93b2295ef65970e5,
		reflect.Int8:          __obf_93b2295ef65970e5,
		reflect.Int16:         __obf_93b2295ef65970e5,
		reflect.Int32:         __obf_93b2295ef65970e5,
		reflect.Int64:         __obf_93b2295ef65970e5,
		reflect.Uint:          __obf_d5ee3308a7edaa80,
		reflect.Uint8:         __obf_d5ee3308a7edaa80,
		reflect.Uint16:        __obf_d5ee3308a7edaa80,
		reflect.Uint32:        __obf_d5ee3308a7edaa80,
		reflect.Uint64:        __obf_d5ee3308a7edaa80,
		reflect.Float32:       __obf_79d81581a625ef22,
		reflect.Float64:       __obf_2c87d512f404034f,
		reflect.Complex64:     __obf_cd47cad1f81437fb,
		reflect.Complex128:    __obf_cd47cad1f81437fb,
		reflect.Array:         __obf_a51ce3b2fcc836ae,
		reflect.Chan:          __obf_cd47cad1f81437fb,
		reflect.Func:          __obf_cd47cad1f81437fb,
		reflect.Interface:     __obf_ceed3ef343f3463a,
		reflect.Map:           __obf_d50c45158b108e4f,
		reflect.Pointer:       __obf_cd47cad1f81437fb,
		reflect.Slice:         __obf_8ba462222fa73e4f,
		reflect.String:        __obf_d48bbb6e52137026,
		reflect.Struct:        __obf_9651070b33634b16,
		reflect.UnsafePointer: __obf_cd47cad1f81437fb,
	}
}

func __obf_7581b02f483ae445(__obf_7bd64e8df8357cab reflect.Type) __obf_8e00c4b112264dba {
	if __obf_6d570581f4b60dbc, __obf_826ac2dbb957d7df := __obf_48ad70830e899de5.Load(__obf_7bd64e8df8357cab); __obf_826ac2dbb957d7df {
		return __obf_6d570581f4b60dbc.(__obf_8e00c4b112264dba)
	}
	__obf_36c78c696c47cfdb := _getDecoder(__obf_7bd64e8df8357cab)
	__obf_48ad70830e899de5.
		Store(__obf_7bd64e8df8357cab, __obf_36c78c696c47cfdb)
	return __obf_36c78c696c47cfdb
}

func _getDecoder(__obf_7bd64e8df8357cab reflect.Type) __obf_8e00c4b112264dba {
	__obf_38e9ce3a1f9f7821 := __obf_7bd64e8df8357cab.Kind()

	if __obf_38e9ce3a1f9f7821 == reflect.Ptr {
		if _, __obf_826ac2dbb957d7df := __obf_48ad70830e899de5.Load(__obf_7bd64e8df8357cab.Elem()); __obf_826ac2dbb957d7df {
			return __obf_4a88ac53d49eb0a7(__obf_7bd64e8df8357cab)
		}
	}

	if __obf_7bd64e8df8357cab.Implements(__obf_2160cf241ba36353) {
		return __obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_f1ed41da8f812d43)
	}
	if __obf_7bd64e8df8357cab.Implements(__obf_cf21162dfe1ffede) {
		return __obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_74eecddb150a9b98)
	}
	if __obf_7bd64e8df8357cab.Implements(__obf_26e34464a5da7cb3) {
		return __obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_0e4b8dfac7922af0)
	}
	if __obf_7bd64e8df8357cab.Implements(__obf_b02d3778ade2abbd) {
		return __obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_c6bad22bf39c11af)
	}

	// Addressable struct field value.
	if __obf_38e9ce3a1f9f7821 != reflect.Pointer {
		__obf_0d8a994785cda6df := reflect.PointerTo(__obf_7bd64e8df8357cab)
		if __obf_0d8a994785cda6df.Implements(__obf_2160cf241ba36353) {
			return __obf_9b7bcbc5f17a3202(__obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_f1ed41da8f812d43))
		}
		if __obf_0d8a994785cda6df.Implements(__obf_cf21162dfe1ffede) {
			return __obf_9b7bcbc5f17a3202(__obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_74eecddb150a9b98))
		}
		if __obf_0d8a994785cda6df.Implements(__obf_26e34464a5da7cb3) {
			return __obf_9b7bcbc5f17a3202(__obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_0e4b8dfac7922af0))
		}
		if __obf_0d8a994785cda6df.Implements(__obf_b02d3778ade2abbd) {
			return __obf_9b7bcbc5f17a3202(__obf_61112f45bafb05f6(__obf_7bd64e8df8357cab, __obf_c6bad22bf39c11af))
		}
	}

	switch __obf_38e9ce3a1f9f7821 {
	case reflect.Pointer:
		return __obf_4a88ac53d49eb0a7(__obf_7bd64e8df8357cab)
	case reflect.Slice:
		__obf_f50dd3cfb2ad5435 := __obf_7bd64e8df8357cab.Elem()
		if __obf_f50dd3cfb2ad5435.Kind() == reflect.Uint8 {
			return __obf_1c67f7750bd90b37
		}
		if __obf_f50dd3cfb2ad5435 == __obf_b128fe7c7d10a668 {
			return __obf_9b16848de3b7b97b
		}
	case reflect.Array:
		if __obf_7bd64e8df8357cab.Elem().Kind() == reflect.Uint8 {
			return __obf_499a52ef00987055
		}
	case reflect.Map:
		if __obf_7bd64e8df8357cab.Key() == __obf_b128fe7c7d10a668 {
			switch __obf_7bd64e8df8357cab.Elem() {
			case __obf_b128fe7c7d10a668:
				return __obf_aa35887bffb65a6c
			case __obf_24d1b614db17eb2b:
				return __obf_b269ad4ecc53ab26
			}
		}
	}

	return __obf_bc021b46631f5774[__obf_38e9ce3a1f9f7821]
}

func __obf_4a88ac53d49eb0a7(__obf_7bd64e8df8357cab reflect.Type) __obf_8e00c4b112264dba {
	__obf_a17f840f59d84bbc := __obf_7581b02f483ae445(__obf_7bd64e8df8357cab.Elem())
	return func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if __obf_a21885da2425f2b2.__obf_5e95865dadae4ab4() {
			if !__obf_6d570581f4b60dbc.IsNil() {
				__obf_6d570581f4b60dbc.
					Set(__obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_7bd64e8df8357cab).Elem())
			}
			return __obf_a21885da2425f2b2.DecodeNil()
		}
		if __obf_6d570581f4b60dbc.IsNil() {
			__obf_6d570581f4b60dbc.
				Set(__obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_7bd64e8df8357cab.Elem()))
		}
		return __obf_a17f840f59d84bbc(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc.Elem())
	}
}

func __obf_9b7bcbc5f17a3202(__obf_36c78c696c47cfdb __obf_8e00c4b112264dba) __obf_8e00c4b112264dba {
	return func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if !__obf_6d570581f4b60dbc.CanAddr() {
			return fmt.Errorf("msgpack: Decode(nonaddressable %T)", __obf_6d570581f4b60dbc.Interface())
		}
		return __obf_36c78c696c47cfdb(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc.Addr())
	}
}

func __obf_61112f45bafb05f6(__obf_7bd64e8df8357cab reflect.Type, __obf_36c78c696c47cfdb __obf_8e00c4b112264dba) __obf_8e00c4b112264dba {
	if __obf_8b51c57d6f65dcf5(__obf_7bd64e8df8357cab.Kind()) {
		return func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
			if __obf_a21885da2425f2b2.__obf_5e95865dadae4ab4() {
				return __obf_a21885da2425f2b2.__obf_30a79fe20d689305(__obf_6d570581f4b60dbc)
			}
			if __obf_6d570581f4b60dbc.IsNil() {
				__obf_6d570581f4b60dbc.
					Set(__obf_a21885da2425f2b2.__obf_1fc1190bc702b3d7(__obf_7bd64e8df8357cab.Elem()))
			}
			return __obf_36c78c696c47cfdb(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc)
		}
	}

	return func(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if __obf_a21885da2425f2b2.__obf_5e95865dadae4ab4() {
			return __obf_a21885da2425f2b2.__obf_30a79fe20d689305(__obf_6d570581f4b60dbc)
		}
		return __obf_36c78c696c47cfdb(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc)
	}
}

func __obf_f6915865a4c0b583(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_2cbebf6e64b5f24e, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeBool()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetBool(__obf_2cbebf6e64b5f24e)
	return nil
}

func __obf_ceed3ef343f3463a(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_a21885da2425f2b2.__obf_303fd81a8e2a088f(__obf_6d570581f4b60dbc)
	}
	return __obf_a21885da2425f2b2.DecodeValue(__obf_6d570581f4b60dbc.Elem())
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_303fd81a8e2a088f(__obf_6d570581f4b60dbc reflect.Value) error {
	__obf_6fee38a618ed3afa, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_6fee38a618ed3afa != nil {
		if __obf_6d570581f4b60dbc.Type() == __obf_4ab64edf794ae331 {
			if __obf_6fee38a618ed3afa, __obf_826ac2dbb957d7df := __obf_6fee38a618ed3afa.(string); __obf_826ac2dbb957d7df {
				__obf_6d570581f4b60dbc.
					Set(reflect.ValueOf(errors.New(__obf_6fee38a618ed3afa)))
				return nil
			}
		}
		__obf_6d570581f4b60dbc.
			Set(reflect.ValueOf(__obf_6fee38a618ed3afa))
	}

	return nil
}

func __obf_cd47cad1f81437fb(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return fmt.Errorf("msgpack: Decode(unsupported %s)", __obf_6d570581f4b60dbc.Type())
}

//------------------------------------------------------------------------------

func __obf_f1ed41da8f812d43(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_a17f840f59d84bbc := __obf_6d570581f4b60dbc.Interface().(CustomDecoder)
	return __obf_a17f840f59d84bbc.DecodeMsgpack(__obf_a21885da2425f2b2)
}

func __obf_74eecddb150a9b98(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	var __obf_f2ca794293605b73 []byte
	__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = make([]byte, 0, 64)
	if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_f2ca794293605b73 = __obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f
	__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = nil
	__obf_9f41f336ee20aa6f := __obf_6d570581f4b60dbc.Interface().(Unmarshaler)
	return __obf_9f41f336ee20aa6f.UnmarshalMsgpack(__obf_f2ca794293605b73)
}

func __obf_0e4b8dfac7922af0(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_652c67787b9c24c9, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeBytes()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_9f41f336ee20aa6f := __obf_6d570581f4b60dbc.Interface().(encoding.BinaryUnmarshaler)
	return __obf_9f41f336ee20aa6f.UnmarshalBinary(__obf_652c67787b9c24c9)
}

func __obf_c6bad22bf39c11af(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_652c67787b9c24c9, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeBytes()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_9f41f336ee20aa6f := __obf_6d570581f4b60dbc.Interface().(encoding.TextUnmarshaler)
	return __obf_9f41f336ee20aa6f.UnmarshalText(__obf_652c67787b9c24c9)
}
