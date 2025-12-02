package __obf_a935eb7f548271a4

import (
	"encoding"
	"fmt"
	tagparser "github.com/ArtisanHiram/go-pkg/tagparser"
	"log"
	"reflect"
	"sync"
)

var __obf_4ab64edf794ae331 = reflect.TypeOf((*error)(nil)).Elem()

var (
	__obf_0f9b3527fa1e5883 = reflect.TypeOf((*CustomEncoder)(nil)).Elem()
	__obf_2160cf241ba36353 = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
)

var (
	__obf_626f6da86a13b183 = reflect.TypeOf((*Marshaler)(nil)).Elem()
	__obf_cf21162dfe1ffede = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var (
	__obf_17933e997378c10f = reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	__obf_26e34464a5da7cb3 = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

var (
	__obf_499c27ef04771594 = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	__obf_b02d3778ade2abbd = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

type (
	__obf_4982458b85add976 func(*Encoder, reflect.Value) error
	__obf_8e00c4b112264dba func(*Decoder, reflect.Value) error
)

var (
	__obf_a16fc4f2eeccec6d sync.Map
	__obf_48ad70830e899de5 sync.Map
)

// Register registers encoder and decoder functions for a value.
// This is low level API and in most cases you should prefer implementing
// CustomEncoder/CustomDecoder or Marshaler/Unmarshaler interfaces.
func Register(__obf_205ee5d820f118f1 any, __obf_6ca0309463018ed3 __obf_4982458b85add976, __obf_fbe570faa7707c02 __obf_8e00c4b112264dba) {
	__obf_7bd64e8df8357cab := reflect.TypeOf(__obf_205ee5d820f118f1)
	if __obf_6ca0309463018ed3 != nil {
		__obf_a16fc4f2eeccec6d.
			Store(__obf_7bd64e8df8357cab, __obf_6ca0309463018ed3)
	}
	if __obf_fbe570faa7707c02 != nil {
		__obf_48ad70830e899de5.
			Store(__obf_7bd64e8df8357cab,

				//------------------------------------------------------------------------------
				__obf_fbe570faa7707c02)
	}
}

const __obf_46536fb00199c1c1 = "msgpack"

var __obf_b7d9de4837c900e0 = __obf_6baff3bdecae317e()

type __obf_07eea9df2b0e1214 struct {
	__obf_26d12ef0df36a324 sync.Map
}

type __obf_335b15246df7cf16 struct {
	__obf_7bd64e8df8357cab reflect.Type
	__obf_f03106ae8627dc7d string
}

func __obf_6baff3bdecae317e() *__obf_07eea9df2b0e1214 {
	return new(__obf_07eea9df2b0e1214)
}

func (__obf_26d12ef0df36a324 *__obf_07eea9df2b0e1214) Fields(__obf_7bd64e8df8357cab reflect.Type, __obf_f03106ae8627dc7d string) *__obf_02acefac40101465 {
	__obf_5603ee7e7690c5f5 := __obf_335b15246df7cf16{__obf_f03106ae8627dc7d: __obf_f03106ae8627dc7d, __obf_7bd64e8df8357cab: __obf_7bd64e8df8357cab}

	if __obf_6d570581f4b60dbc, __obf_826ac2dbb957d7df := __obf_26d12ef0df36a324.__obf_26d12ef0df36a324.Load(__obf_5603ee7e7690c5f5); __obf_826ac2dbb957d7df {
		return __obf_6d570581f4b60dbc.(*__obf_02acefac40101465)
	}
	__obf_9c2b4d6b1671402e := __obf_491d5f7fd00638ce(__obf_7bd64e8df8357cab, __obf_f03106ae8627dc7d)
	__obf_26d12ef0df36a324.__obf_26d12ef0df36a324.
		Store(__obf_5603ee7e7690c5f5, __obf_9c2b4d6b1671402e)

	return __obf_9c2b4d6b1671402e
}

//------------------------------------------------------------------------------

type __obf_3b4b6de10b0b5ac0 struct {
	__obf_6a114719334e34d2 __obf_4982458b85add976
	__obf_a17f840f59d84bbc __obf_8e00c4b112264dba
	__obf_e80b5f2d9b7252c4 string
	__obf_103c16dc90a9ea84 []int
	__obf_bc86f05139d04ea8 bool
}

func (__obf_f6dd34b5068d19f3 *__obf_3b4b6de10b0b5ac0) Omit(__obf_ed37a34c347049f3 *Encoder, __obf_5e97d46983ad5951 reflect.Value) bool {
	__obf_6d570581f4b60dbc, __obf_826ac2dbb957d7df := __obf_32504e8d6a3ea0f7(__obf_5e97d46983ad5951, __obf_f6dd34b5068d19f3.__obf_103c16dc90a9ea84)
	if !__obf_826ac2dbb957d7df {
		return true
	}
	__obf_86e74603e42d86ef := __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_48749b6f1055fd1f != 0
	return (__obf_f6dd34b5068d19f3.__obf_bc86f05139d04ea8 || __obf_86e74603e42d86ef) && __obf_ed37a34c347049f3.__obf_63b36a1bf7d91690(__obf_6d570581f4b60dbc)
}

func (__obf_f6dd34b5068d19f3 *__obf_3b4b6de10b0b5ac0) EncodeValue(__obf_ed37a34c347049f3 *Encoder, __obf_5e97d46983ad5951 reflect.Value) error {
	__obf_6d570581f4b60dbc, __obf_826ac2dbb957d7df := __obf_32504e8d6a3ea0f7(__obf_5e97d46983ad5951, __obf_f6dd34b5068d19f3.__obf_103c16dc90a9ea84)
	if !__obf_826ac2dbb957d7df {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	return __obf_f6dd34b5068d19f3.__obf_6a114719334e34d2(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc)
}

func (__obf_f6dd34b5068d19f3 *__obf_3b4b6de10b0b5ac0) DecodeValue(__obf_a21885da2425f2b2 *Decoder, __obf_5e97d46983ad5951 reflect.Value) error {
	__obf_6d570581f4b60dbc := __obf_cc0e964b1b683764(__obf_5e97d46983ad5951, __obf_f6dd34b5068d19f3.

		//------------------------------------------------------------------------------
		__obf_103c16dc90a9ea84)
	return __obf_f6dd34b5068d19f3.__obf_a17f840f59d84bbc(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc)
}

type __obf_02acefac40101465 struct {
	Type                   reflect.Type
	Map                    map[string]*__obf_3b4b6de10b0b5ac0
	List                   []*__obf_3b4b6de10b0b5ac0
	AsArray                bool
	__obf_08f4807129d222bf bool
}

func __obf_1d483e0eb80c7f96(__obf_7bd64e8df8357cab reflect.Type) *__obf_02acefac40101465 {
	return &__obf_02acefac40101465{
		Type: __obf_7bd64e8df8357cab,
		Map:  make(map[string]*__obf_3b4b6de10b0b5ac0, __obf_7bd64e8df8357cab.NumField()),
		List: make([]*__obf_3b4b6de10b0b5ac0, 0, __obf_7bd64e8df8357cab.NumField()),
	}
}

func (__obf_9c2b4d6b1671402e *__obf_02acefac40101465) Add(__obf_3b4b6de10b0b5ac0 *__obf_3b4b6de10b0b5ac0) {
	__obf_9c2b4d6b1671402e.__obf_ab4f902703c45623(__obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4)
	__obf_9c2b4d6b1671402e.
		Map[__obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4] = __obf_3b4b6de10b0b5ac0
	__obf_9c2b4d6b1671402e.
		List = append(__obf_9c2b4d6b1671402e.List, __obf_3b4b6de10b0b5ac0)
	if __obf_3b4b6de10b0b5ac0.__obf_bc86f05139d04ea8 {
		__obf_9c2b4d6b1671402e.__obf_08f4807129d222bf = true
	}
}

func (__obf_9c2b4d6b1671402e *__obf_02acefac40101465) __obf_ab4f902703c45623(__obf_e80b5f2d9b7252c4 string) {
	if _, __obf_826ac2dbb957d7df := __obf_9c2b4d6b1671402e.Map[__obf_e80b5f2d9b7252c4]; __obf_826ac2dbb957d7df {
		log.Printf("msgpack: %s already has field=%s", __obf_9c2b4d6b1671402e.Type, __obf_e80b5f2d9b7252c4)
	}
}

func (__obf_9c2b4d6b1671402e *__obf_02acefac40101465) OmitEmpty(__obf_ed37a34c347049f3 *Encoder, __obf_5e97d46983ad5951 reflect.Value) []*__obf_3b4b6de10b0b5ac0 {
	__obf_86e74603e42d86ef := __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_48749b6f1055fd1f != 0
	if !__obf_9c2b4d6b1671402e.__obf_08f4807129d222bf && !__obf_86e74603e42d86ef {
		return __obf_9c2b4d6b1671402e.List
	}
	__obf_02acefac40101465 := make([]*__obf_3b4b6de10b0b5ac0, 0, len(__obf_9c2b4d6b1671402e.List))

	for _, __obf_f6dd34b5068d19f3 := range __obf_9c2b4d6b1671402e.List {
		if !__obf_f6dd34b5068d19f3.Omit(__obf_ed37a34c347049f3, __obf_5e97d46983ad5951) {
			__obf_02acefac40101465 = append(__obf_02acefac40101465, __obf_f6dd34b5068d19f3)
		}
	}

	return __obf_02acefac40101465
}

func __obf_491d5f7fd00638ce(__obf_7bd64e8df8357cab reflect.Type, __obf_b2b5c8cb4ebffae9 string) *__obf_02acefac40101465 {
	__obf_9c2b4d6b1671402e := __obf_1d483e0eb80c7f96(__obf_7bd64e8df8357cab)

	var __obf_bc86f05139d04ea8 bool
	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_7bd64e8df8357cab.NumField(); __obf_4a8e280ffaa52cf4++ {
		__obf_f6dd34b5068d19f3 := __obf_7bd64e8df8357cab.Field(__obf_4a8e280ffaa52cf4)
		__obf_10c365d04fd4e906 := __obf_f6dd34b5068d19f3.Tag.Get(__obf_46536fb00199c1c1)
		if __obf_10c365d04fd4e906 == "" && __obf_b2b5c8cb4ebffae9 != "" {
			__obf_10c365d04fd4e906 = __obf_f6dd34b5068d19f3.Tag.Get(__obf_b2b5c8cb4ebffae9)
		}
		__obf_f03106ae8627dc7d := tagparser.Parse(__obf_10c365d04fd4e906)
		if __obf_f03106ae8627dc7d.Name == "-" {
			continue
		}

		if __obf_f6dd34b5068d19f3.Name == "_msgpack" {
			__obf_9c2b4d6b1671402e.
				AsArray = __obf_f03106ae8627dc7d.HasOption("as_array") || __obf_f03106ae8627dc7d.HasOption("asArray")
			if __obf_f03106ae8627dc7d.HasOption("omitempty") {
				__obf_bc86f05139d04ea8 = true
			}
		}

		if __obf_f6dd34b5068d19f3.PkgPath != "" && !__obf_f6dd34b5068d19f3.Anonymous {
			continue
		}
		__obf_3b4b6de10b0b5ac0 := &__obf_3b4b6de10b0b5ac0{__obf_e80b5f2d9b7252c4: __obf_f03106ae8627dc7d.Name, __obf_103c16dc90a9ea84: __obf_f6dd34b5068d19f3.Index, __obf_bc86f05139d04ea8: __obf_bc86f05139d04ea8 || __obf_f03106ae8627dc7d.HasOption("omitempty")}

		if __obf_f03106ae8627dc7d.HasOption("intern") {
			switch __obf_f6dd34b5068d19f3.Type.Kind() {
			case reflect.Interface:
				__obf_3b4b6de10b0b5ac0.__obf_6a114719334e34d2 = __obf_14ce8aee2699d8c1
				__obf_3b4b6de10b0b5ac0.__obf_a17f840f59d84bbc = __obf_9d0cdf74c5f70d62
			case reflect.String:
				__obf_3b4b6de10b0b5ac0.__obf_6a114719334e34d2 = __obf_f1fd7fac8f632e34
				__obf_3b4b6de10b0b5ac0.__obf_a17f840f59d84bbc = __obf_f81a4aeb4252941f
			default:
				__obf_4d327e1cd40c2e21 := fmt.Errorf("msgpack: intern strings are not supported on %s", __obf_f6dd34b5068d19f3.Type)
				panic(__obf_4d327e1cd40c2e21)
			}
		} else {
			__obf_3b4b6de10b0b5ac0.__obf_6a114719334e34d2 = __obf_d55d857e358cbb2a(__obf_f6dd34b5068d19f3.Type)
			__obf_3b4b6de10b0b5ac0.__obf_a17f840f59d84bbc = __obf_7581b02f483ae445(__obf_f6dd34b5068d19f3.Type)
		}

		if __obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4 == "" {
			__obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4 = __obf_f6dd34b5068d19f3.Name
		}

		if __obf_f6dd34b5068d19f3.Anonymous && !__obf_f03106ae8627dc7d.HasOption("noinline") {
			__obf_fd1392efb70dce72 := __obf_f03106ae8627dc7d.HasOption("inline")
			if __obf_fd1392efb70dce72 {
				__obf_228047e1a27a86d9(__obf_9c2b4d6b1671402e, __obf_f6dd34b5068d19f3.Type, __obf_3b4b6de10b0b5ac0, __obf_b2b5c8cb4ebffae9)
			} else {
				__obf_fd1392efb70dce72 = __obf_89ac8721a35f4fbc(__obf_9c2b4d6b1671402e, __obf_f6dd34b5068d19f3.Type, __obf_3b4b6de10b0b5ac0, __obf_b2b5c8cb4ebffae9)
			}

			if __obf_fd1392efb70dce72 {
				if _, __obf_826ac2dbb957d7df := __obf_9c2b4d6b1671402e.Map[__obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4]; __obf_826ac2dbb957d7df {
					log.Printf("msgpack: %s already has field=%s", __obf_9c2b4d6b1671402e.Type, __obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4)
				}
				__obf_9c2b4d6b1671402e.
					Map[__obf_3b4b6de10b0b5ac0.__obf_e80b5f2d9b7252c4] = __obf_3b4b6de10b0b5ac0
				continue
			}
		}
		__obf_9c2b4d6b1671402e.
			Add(__obf_3b4b6de10b0b5ac0)

		if __obf_3cb2272e310520d9, __obf_826ac2dbb957d7df := __obf_f03106ae8627dc7d.Options["alias"]; __obf_826ac2dbb957d7df {
			__obf_9c2b4d6b1671402e.__obf_ab4f902703c45623(__obf_3cb2272e310520d9)
			__obf_9c2b4d6b1671402e.
				Map[__obf_3cb2272e310520d9] = __obf_3b4b6de10b0b5ac0
		}
	}
	return __obf_9c2b4d6b1671402e
}

var (
	__obf_7a6f5eda6eb6eb9a uintptr
	__obf_39b9436645424756 uintptr
)

//nolint:gochecknoinits
func init() {
	__obf_7a6f5eda6eb6eb9a = reflect.ValueOf(__obf_1d06e69b22f89920).Pointer()
	__obf_39b9436645424756 = reflect.ValueOf(__obf_9651070b33634b16).Pointer()
}

func __obf_228047e1a27a86d9(__obf_9c2b4d6b1671402e *__obf_02acefac40101465, __obf_7bd64e8df8357cab reflect.Type, __obf_f6dd34b5068d19f3 *__obf_3b4b6de10b0b5ac0, __obf_f03106ae8627dc7d string) {
	__obf_05a9c97bec248086 := __obf_491d5f7fd00638ce(__obf_7bd64e8df8357cab, __obf_f03106ae8627dc7d).List
	for _, __obf_3b4b6de10b0b5ac0 := range __obf_05a9c97bec248086 {
		if _, __obf_826ac2dbb957d7df := __obf_9c2b4d6b1671402e.Map[__obf_3b4b6de10b0b5ac0.
			// Don't inline shadowed fields.
			__obf_e80b5f2d9b7252c4]; __obf_826ac2dbb957d7df {

			continue
		}
		__obf_3b4b6de10b0b5ac0.__obf_103c16dc90a9ea84 = append(__obf_f6dd34b5068d19f3.__obf_103c16dc90a9ea84, __obf_3b4b6de10b0b5ac0.__obf_103c16dc90a9ea84...)
		__obf_9c2b4d6b1671402e.
			Add(__obf_3b4b6de10b0b5ac0)
	}
}

func __obf_89ac8721a35f4fbc(__obf_9c2b4d6b1671402e *__obf_02acefac40101465, __obf_7bd64e8df8357cab reflect.Type, __obf_f6dd34b5068d19f3 *__obf_3b4b6de10b0b5ac0, __obf_f03106ae8627dc7d string) bool {
	var __obf_6a114719334e34d2 __obf_4982458b85add976
	var __obf_a17f840f59d84bbc __obf_8e00c4b112264dba

	if __obf_7bd64e8df8357cab.Kind() == reflect.Struct {
		__obf_6a114719334e34d2 = __obf_f6dd34b5068d19f3.__obf_6a114719334e34d2
		__obf_a17f840f59d84bbc = __obf_f6dd34b5068d19f3.__obf_a17f840f59d84bbc
	} else {
		for __obf_7bd64e8df8357cab.Kind() == reflect.Ptr {
			__obf_7bd64e8df8357cab = __obf_7bd64e8df8357cab.Elem()
			__obf_6a114719334e34d2 = __obf_d55d857e358cbb2a(__obf_7bd64e8df8357cab)
			__obf_a17f840f59d84bbc = __obf_7581b02f483ae445(__obf_7bd64e8df8357cab)
		}
		if __obf_7bd64e8df8357cab.Kind() != reflect.Struct {
			return false
		}
	}

	if reflect.ValueOf(__obf_6a114719334e34d2).Pointer() != __obf_7a6f5eda6eb6eb9a {
		return false
	}
	if reflect.ValueOf(__obf_a17f840f59d84bbc).Pointer() != __obf_39b9436645424756 {
		return false
	}
	__obf_05a9c97bec248086 := __obf_491d5f7fd00638ce(__obf_7bd64e8df8357cab, __obf_f03106ae8627dc7d).List
	for _, __obf_3b4b6de10b0b5ac0 := range __obf_05a9c97bec248086 {
		if _, __obf_826ac2dbb957d7df := __obf_9c2b4d6b1671402e.Map[__obf_3b4b6de10b0b5ac0.
			// Don't auto inline if there are shadowed fields.
			__obf_e80b5f2d9b7252c4]; __obf_826ac2dbb957d7df {

			return false
		}
	}

	for _, __obf_3b4b6de10b0b5ac0 := range __obf_05a9c97bec248086 {
		__obf_3b4b6de10b0b5ac0.__obf_103c16dc90a9ea84 = append(__obf_f6dd34b5068d19f3.__obf_103c16dc90a9ea84, __obf_3b4b6de10b0b5ac0.__obf_103c16dc90a9ea84...)
		__obf_9c2b4d6b1671402e.
			Add(__obf_3b4b6de10b0b5ac0)
	}
	return true
}

type __obf_85d9957d1058979f interface {
	IsZero() bool
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_63b36a1bf7d91690(__obf_6d570581f4b60dbc reflect.Value) bool {
	__obf_38e9ce3a1f9f7821 := __obf_6d570581f4b60dbc.Kind()

	for __obf_38e9ce3a1f9f7821 == reflect.Interface {
		if __obf_6d570581f4b60dbc.IsNil() {
			return true
		}
		__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Elem()
		__obf_38e9ce3a1f9f7821 = __obf_6d570581f4b60dbc.Kind()
	}

	if __obf_b172314ef57dca6f, __obf_826ac2dbb957d7df := __obf_6d570581f4b60dbc.Interface().(__obf_85d9957d1058979f); __obf_826ac2dbb957d7df {
		return __obf_8b51c57d6f65dcf5(__obf_38e9ce3a1f9f7821) && __obf_6d570581f4b60dbc.IsNil() || __obf_b172314ef57dca6f.IsZero()
	}

	switch __obf_38e9ce3a1f9f7821 {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_6d570581f4b60dbc.Len() == 0
	case reflect.Struct:
		__obf_bf5291b060bc4afb := __obf_b7d9de4837c900e0.Fields(__obf_6d570581f4b60dbc.Type(), __obf_ed37a34c347049f3.__obf_90e41d275c699c20)
		__obf_02acefac40101465 := __obf_bf5291b060bc4afb.OmitEmpty(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc)
		return len(__obf_02acefac40101465) == 0
	case reflect.Bool:
		return !__obf_6d570581f4b60dbc.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_6d570581f4b60dbc.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_6d570581f4b60dbc.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_6d570581f4b60dbc.Float() == 0
	case reflect.Ptr:
		return __obf_6d570581f4b60dbc.IsNil()
	default:
		return false
	}
}

func __obf_32504e8d6a3ea0f7(__obf_6d570581f4b60dbc reflect.Value, __obf_103c16dc90a9ea84 []int) (_ reflect.Value, __obf_826ac2dbb957d7df bool) {
	if len(__obf_103c16dc90a9ea84) == 1 {
		return __obf_6d570581f4b60dbc.Field(__obf_103c16dc90a9ea84[0]), true
	}

	for __obf_4a8e280ffaa52cf4, __obf_b8b71f09b098e0ae := range __obf_103c16dc90a9ea84 {
		if __obf_4a8e280ffaa52cf4 > 0 {
			if __obf_6d570581f4b60dbc.Kind() == reflect.Ptr {
				if __obf_6d570581f4b60dbc.IsNil() {
					return __obf_6d570581f4b60dbc, false
				}
				__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Elem()
			}
		}
		__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Field(__obf_b8b71f09b098e0ae)
	}

	return __obf_6d570581f4b60dbc, true
}

func __obf_cc0e964b1b683764(__obf_6d570581f4b60dbc reflect.Value, __obf_103c16dc90a9ea84 []int) reflect.Value {
	if len(__obf_103c16dc90a9ea84) == 1 {
		return __obf_6d570581f4b60dbc.Field(__obf_103c16dc90a9ea84[0])
	}

	for __obf_4a8e280ffaa52cf4, __obf_b8b71f09b098e0ae := range __obf_103c16dc90a9ea84 {
		if __obf_4a8e280ffaa52cf4 > 0 {
			var __obf_826ac2dbb957d7df bool
			__obf_6d570581f4b60dbc, __obf_826ac2dbb957d7df = __obf_899c5d2fcfae136e(__obf_6d570581f4b60dbc)
			if !__obf_826ac2dbb957d7df {
				return __obf_6d570581f4b60dbc
			}
		}
		__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Field(__obf_b8b71f09b098e0ae)
	}

	return __obf_6d570581f4b60dbc
}

func __obf_899c5d2fcfae136e(__obf_6d570581f4b60dbc reflect.Value) (reflect.Value, bool) {
	if __obf_6d570581f4b60dbc.Kind() == reflect.Ptr {
		if __obf_6d570581f4b60dbc.IsNil() {
			if !__obf_6d570581f4b60dbc.CanSet() {
				return __obf_6d570581f4b60dbc, false
			}
			__obf_065eadb9468b7d59 := __obf_6d570581f4b60dbc.Type().Elem()
			if __obf_065eadb9468b7d59.Kind() != reflect.Struct {
				return __obf_6d570581f4b60dbc, false
			}
			__obf_6d570581f4b60dbc.
				Set(__obf_14360e03a3b5da30(__obf_065eadb9468b7d59))
		}
		__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Elem()
	}
	return __obf_6d570581f4b60dbc, true
}
