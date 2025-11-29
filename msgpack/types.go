package __obf_a4aad98aaf3178f4

import (
	"encoding"
	"fmt"
	tagparser "github.com/ArtisanHiram/go-pkg/tagparser"
	"log"
	"reflect"
	"sync"
)

var __obf_d66851dd29a0f414 = reflect.TypeOf((*error)(nil)).Elem()

var (
	__obf_a644aad6e1b3ff55 = reflect.TypeOf((*CustomEncoder)(nil)).Elem()
	__obf_283fd634becb1665 = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
)

var (
	__obf_1b905c6b55da13d6 = reflect.TypeOf((*Marshaler)(nil)).Elem()
	__obf_26665c3f2693777c = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var (
	__obf_8b41e5fad6d95fea = reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	__obf_4da005c7e58bcab6 = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

var (
	__obf_efdcdb0d1ee35009 = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	__obf_1e6abd2dc26ca3b3 = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

type (
	__obf_9d8023233a199a4f func(*Encoder, reflect.Value) error
	__obf_c1c8351491cde4cc func(*Decoder, reflect.Value) error
)

var (
	__obf_a25a843d85821e5c sync.Map
	__obf_6a2362b607537b3d sync.Map
)

// Register registers encoder and decoder functions for a value.
// This is low level API and in most cases you should prefer implementing
// CustomEncoder/CustomDecoder or Marshaler/Unmarshaler interfaces.
func Register(__obf_9055e1481edcd436 any, __obf_dfba1ffd920163bc __obf_9d8023233a199a4f, __obf_cc68ee77d25ca8f2 __obf_c1c8351491cde4cc) {
	__obf_bbfd30fcc08dc1bc := reflect.TypeOf(__obf_9055e1481edcd436)
	if __obf_dfba1ffd920163bc != nil {
		__obf_a25a843d85821e5c.
			Store(__obf_bbfd30fcc08dc1bc, __obf_dfba1ffd920163bc)
	}
	if __obf_cc68ee77d25ca8f2 != nil {
		__obf_6a2362b607537b3d.
			Store(__obf_bbfd30fcc08dc1bc,

				//------------------------------------------------------------------------------
				__obf_cc68ee77d25ca8f2)
	}
}

const __obf_d38dc40ea03e2c2b = "msgpack"

var __obf_bd827009b4ede0e2 = __obf_e98475c600625ab2()

type __obf_9a95d7e378720b6b struct {
	__obf_66cc4b26e3c9cdbb sync.Map
}

type __obf_ddf7e55a89b61f3f struct {
	__obf_bbfd30fcc08dc1bc reflect.Type
	__obf_990470bfaf220103 string
}

func __obf_e98475c600625ab2() *__obf_9a95d7e378720b6b {
	return new(__obf_9a95d7e378720b6b)
}

func (__obf_66cc4b26e3c9cdbb *__obf_9a95d7e378720b6b) Fields(__obf_bbfd30fcc08dc1bc reflect.Type, __obf_990470bfaf220103 string) *__obf_0cf794214f02df4d {
	__obf_992cb53d7b9cb024 := __obf_ddf7e55a89b61f3f{__obf_990470bfaf220103: __obf_990470bfaf220103, __obf_bbfd30fcc08dc1bc: __obf_bbfd30fcc08dc1bc}

	if __obf_a1f43267eeb48a1a, __obf_81b19f2a6159ab89 := __obf_66cc4b26e3c9cdbb.__obf_66cc4b26e3c9cdbb.Load(__obf_992cb53d7b9cb024); __obf_81b19f2a6159ab89 {
		return __obf_a1f43267eeb48a1a.(*__obf_0cf794214f02df4d)
	}
	__obf_f82968c6203f3e5f := __obf_19dc6e60fa248a31(__obf_bbfd30fcc08dc1bc, __obf_990470bfaf220103)
	__obf_66cc4b26e3c9cdbb.__obf_66cc4b26e3c9cdbb.
		Store(__obf_992cb53d7b9cb024, __obf_f82968c6203f3e5f)

	return __obf_f82968c6203f3e5f
}

//------------------------------------------------------------------------------

type __obf_efecfc090178574b struct {
	__obf_c4f748de4fd652e4 __obf_9d8023233a199a4f
	__obf_815edc9179203b9a __obf_c1c8351491cde4cc
	__obf_071b55c16b16fe35 string
	__obf_eb6d7a2127ac443c []int
	__obf_54872002c509f9af bool
}

func (__obf_da4a2296298d318f *__obf_efecfc090178574b) Omit(__obf_2c8e97779108ab17 *Encoder, __obf_7500ff0c7bba82eb reflect.Value) bool {
	__obf_a1f43267eeb48a1a, __obf_81b19f2a6159ab89 := __obf_01236718d8fc377c(__obf_7500ff0c7bba82eb, __obf_da4a2296298d318f.__obf_eb6d7a2127ac443c)
	if !__obf_81b19f2a6159ab89 {
		return true
	}
	__obf_77fbdc238afe14e3 := __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_33a8b849a29cb9a5 != 0
	return (__obf_da4a2296298d318f.__obf_54872002c509f9af || __obf_77fbdc238afe14e3) && __obf_2c8e97779108ab17.__obf_9c83a9dd41bc3039(__obf_a1f43267eeb48a1a)
}

func (__obf_da4a2296298d318f *__obf_efecfc090178574b) EncodeValue(__obf_2c8e97779108ab17 *Encoder, __obf_7500ff0c7bba82eb reflect.Value) error {
	__obf_a1f43267eeb48a1a, __obf_81b19f2a6159ab89 := __obf_01236718d8fc377c(__obf_7500ff0c7bba82eb, __obf_da4a2296298d318f.__obf_eb6d7a2127ac443c)
	if !__obf_81b19f2a6159ab89 {
		return __obf_2c8e97779108ab17.EncodeNil()
	}
	return __obf_da4a2296298d318f.__obf_c4f748de4fd652e4(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a)
}

func (__obf_da4a2296298d318f *__obf_efecfc090178574b) DecodeValue(__obf_613397fefdec0ed0 *Decoder, __obf_7500ff0c7bba82eb reflect.Value) error {
	__obf_a1f43267eeb48a1a := __obf_0790f3ff4aeb54d2(__obf_7500ff0c7bba82eb, __obf_da4a2296298d318f.

		//------------------------------------------------------------------------------
		__obf_eb6d7a2127ac443c)
	return __obf_da4a2296298d318f.__obf_815edc9179203b9a(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a)
}

type __obf_0cf794214f02df4d struct {
	Type                   reflect.Type
	Map                    map[string]*__obf_efecfc090178574b
	List                   []*__obf_efecfc090178574b
	AsArray                bool
	__obf_c3743ce6002f3764 bool
}

func __obf_14675d85cdb560af(__obf_bbfd30fcc08dc1bc reflect.Type) *__obf_0cf794214f02df4d {
	return &__obf_0cf794214f02df4d{
		Type: __obf_bbfd30fcc08dc1bc,
		Map:  make(map[string]*__obf_efecfc090178574b, __obf_bbfd30fcc08dc1bc.NumField()),
		List: make([]*__obf_efecfc090178574b, 0, __obf_bbfd30fcc08dc1bc.NumField()),
	}
}

func (__obf_f82968c6203f3e5f *__obf_0cf794214f02df4d) Add(__obf_efecfc090178574b *__obf_efecfc090178574b) {
	__obf_f82968c6203f3e5f.__obf_c0679a8a79b0d990(__obf_efecfc090178574b.__obf_071b55c16b16fe35)
	__obf_f82968c6203f3e5f.
		Map[__obf_efecfc090178574b.__obf_071b55c16b16fe35] = __obf_efecfc090178574b
	__obf_f82968c6203f3e5f.
		List = append(__obf_f82968c6203f3e5f.List, __obf_efecfc090178574b)
	if __obf_efecfc090178574b.__obf_54872002c509f9af {
		__obf_f82968c6203f3e5f.__obf_c3743ce6002f3764 = true
	}
}

func (__obf_f82968c6203f3e5f *__obf_0cf794214f02df4d) __obf_c0679a8a79b0d990(__obf_071b55c16b16fe35 string) {
	if _, __obf_81b19f2a6159ab89 := __obf_f82968c6203f3e5f.Map[__obf_071b55c16b16fe35]; __obf_81b19f2a6159ab89 {
		log.Printf("msgpack: %s already has field=%s", __obf_f82968c6203f3e5f.Type, __obf_071b55c16b16fe35)
	}
}

func (__obf_f82968c6203f3e5f *__obf_0cf794214f02df4d) OmitEmpty(__obf_2c8e97779108ab17 *Encoder, __obf_7500ff0c7bba82eb reflect.Value) []*__obf_efecfc090178574b {
	__obf_77fbdc238afe14e3 := __obf_2c8e97779108ab17.__obf_a8400e70a712e9fa&__obf_33a8b849a29cb9a5 != 0
	if !__obf_f82968c6203f3e5f.__obf_c3743ce6002f3764 && !__obf_77fbdc238afe14e3 {
		return __obf_f82968c6203f3e5f.List
	}
	__obf_0cf794214f02df4d := make([]*__obf_efecfc090178574b, 0, len(__obf_f82968c6203f3e5f.List))

	for _, __obf_da4a2296298d318f := range __obf_f82968c6203f3e5f.List {
		if !__obf_da4a2296298d318f.Omit(__obf_2c8e97779108ab17, __obf_7500ff0c7bba82eb) {
			__obf_0cf794214f02df4d = append(__obf_0cf794214f02df4d, __obf_da4a2296298d318f)
		}
	}

	return __obf_0cf794214f02df4d
}

func __obf_19dc6e60fa248a31(__obf_bbfd30fcc08dc1bc reflect.Type, __obf_ef05ce7047f8fbc1 string) *__obf_0cf794214f02df4d {
	__obf_f82968c6203f3e5f := __obf_14675d85cdb560af(__obf_bbfd30fcc08dc1bc)

	var __obf_54872002c509f9af bool
	for __obf_097d8b6061c9592b := 0; __obf_097d8b6061c9592b < __obf_bbfd30fcc08dc1bc.NumField(); __obf_097d8b6061c9592b++ {
		__obf_da4a2296298d318f := __obf_bbfd30fcc08dc1bc.Field(__obf_097d8b6061c9592b)
		__obf_fb47d8610b56cb66 := __obf_da4a2296298d318f.Tag.Get(__obf_d38dc40ea03e2c2b)
		if __obf_fb47d8610b56cb66 == "" && __obf_ef05ce7047f8fbc1 != "" {
			__obf_fb47d8610b56cb66 = __obf_da4a2296298d318f.Tag.Get(__obf_ef05ce7047f8fbc1)
		}
		__obf_990470bfaf220103 := tagparser.Parse(__obf_fb47d8610b56cb66)
		if __obf_990470bfaf220103.Name == "-" {
			continue
		}

		if __obf_da4a2296298d318f.Name == "_msgpack" {
			__obf_f82968c6203f3e5f.
				AsArray = __obf_990470bfaf220103.HasOption("as_array") || __obf_990470bfaf220103.HasOption("asArray")
			if __obf_990470bfaf220103.HasOption("omitempty") {
				__obf_54872002c509f9af = true
			}
		}

		if __obf_da4a2296298d318f.PkgPath != "" && !__obf_da4a2296298d318f.Anonymous {
			continue
		}
		__obf_efecfc090178574b := &__obf_efecfc090178574b{__obf_071b55c16b16fe35: __obf_990470bfaf220103.Name, __obf_eb6d7a2127ac443c: __obf_da4a2296298d318f.Index, __obf_54872002c509f9af: __obf_54872002c509f9af || __obf_990470bfaf220103.HasOption("omitempty")}

		if __obf_990470bfaf220103.HasOption("intern") {
			switch __obf_da4a2296298d318f.Type.Kind() {
			case reflect.Interface:
				__obf_efecfc090178574b.__obf_c4f748de4fd652e4 = __obf_39ec8e83b1a4666e
				__obf_efecfc090178574b.__obf_815edc9179203b9a = __obf_17a6f68227288f35
			case reflect.String:
				__obf_efecfc090178574b.__obf_c4f748de4fd652e4 = __obf_825879324202941a
				__obf_efecfc090178574b.__obf_815edc9179203b9a = __obf_28e12bbdc2b0bf28
			default:
				__obf_4061ca0039150c39 := fmt.Errorf("msgpack: intern strings are not supported on %s", __obf_da4a2296298d318f.Type)
				panic(__obf_4061ca0039150c39)
			}
		} else {
			__obf_efecfc090178574b.__obf_c4f748de4fd652e4 = __obf_f1f0aa6d8078637c(__obf_da4a2296298d318f.Type)
			__obf_efecfc090178574b.__obf_815edc9179203b9a = __obf_66268eb3a3deccf5(__obf_da4a2296298d318f.Type)
		}

		if __obf_efecfc090178574b.__obf_071b55c16b16fe35 == "" {
			__obf_efecfc090178574b.__obf_071b55c16b16fe35 = __obf_da4a2296298d318f.Name
		}

		if __obf_da4a2296298d318f.Anonymous && !__obf_990470bfaf220103.HasOption("noinline") {
			__obf_f3708040158ecb15 := __obf_990470bfaf220103.HasOption("inline")
			if __obf_f3708040158ecb15 {
				__obf_62f4d3a46c971c36(__obf_f82968c6203f3e5f, __obf_da4a2296298d318f.Type, __obf_efecfc090178574b, __obf_ef05ce7047f8fbc1)
			} else {
				__obf_f3708040158ecb15 = __obf_1ffee7d719941f64(__obf_f82968c6203f3e5f, __obf_da4a2296298d318f.Type, __obf_efecfc090178574b, __obf_ef05ce7047f8fbc1)
			}

			if __obf_f3708040158ecb15 {
				if _, __obf_81b19f2a6159ab89 := __obf_f82968c6203f3e5f.Map[__obf_efecfc090178574b.__obf_071b55c16b16fe35]; __obf_81b19f2a6159ab89 {
					log.Printf("msgpack: %s already has field=%s", __obf_f82968c6203f3e5f.Type, __obf_efecfc090178574b.__obf_071b55c16b16fe35)
				}
				__obf_f82968c6203f3e5f.
					Map[__obf_efecfc090178574b.__obf_071b55c16b16fe35] = __obf_efecfc090178574b
				continue
			}
		}
		__obf_f82968c6203f3e5f.
			Add(__obf_efecfc090178574b)

		if __obf_047f072f57fa9778, __obf_81b19f2a6159ab89 := __obf_990470bfaf220103.Options["alias"]; __obf_81b19f2a6159ab89 {
			__obf_f82968c6203f3e5f.__obf_c0679a8a79b0d990(__obf_047f072f57fa9778)
			__obf_f82968c6203f3e5f.
				Map[__obf_047f072f57fa9778] = __obf_efecfc090178574b
		}
	}
	return __obf_f82968c6203f3e5f
}

var (
	__obf_b287ee13551609b5 uintptr
	__obf_3bd50ef9ce0e4794 uintptr
)

//nolint:gochecknoinits
func init() {
	__obf_b287ee13551609b5 = reflect.ValueOf(__obf_8fcfe77f763e1de9).Pointer()
	__obf_3bd50ef9ce0e4794 = reflect.ValueOf(__obf_cf4307adf926d203).Pointer()
}

func __obf_62f4d3a46c971c36(__obf_f82968c6203f3e5f *__obf_0cf794214f02df4d, __obf_bbfd30fcc08dc1bc reflect.Type, __obf_da4a2296298d318f *__obf_efecfc090178574b, __obf_990470bfaf220103 string) {
	__obf_c3467fac7b886c06 := __obf_19dc6e60fa248a31(__obf_bbfd30fcc08dc1bc, __obf_990470bfaf220103).List
	for _, __obf_efecfc090178574b := range __obf_c3467fac7b886c06 {
		if _, __obf_81b19f2a6159ab89 := __obf_f82968c6203f3e5f.Map[__obf_efecfc090178574b.
			// Don't inline shadowed fields.
			__obf_071b55c16b16fe35]; __obf_81b19f2a6159ab89 {

			continue
		}
		__obf_efecfc090178574b.__obf_eb6d7a2127ac443c = append(__obf_da4a2296298d318f.__obf_eb6d7a2127ac443c, __obf_efecfc090178574b.__obf_eb6d7a2127ac443c...)
		__obf_f82968c6203f3e5f.
			Add(__obf_efecfc090178574b)
	}
}

func __obf_1ffee7d719941f64(__obf_f82968c6203f3e5f *__obf_0cf794214f02df4d, __obf_bbfd30fcc08dc1bc reflect.Type, __obf_da4a2296298d318f *__obf_efecfc090178574b, __obf_990470bfaf220103 string) bool {
	var __obf_c4f748de4fd652e4 __obf_9d8023233a199a4f
	var __obf_815edc9179203b9a __obf_c1c8351491cde4cc

	if __obf_bbfd30fcc08dc1bc.Kind() == reflect.Struct {
		__obf_c4f748de4fd652e4 = __obf_da4a2296298d318f.__obf_c4f748de4fd652e4
		__obf_815edc9179203b9a = __obf_da4a2296298d318f.__obf_815edc9179203b9a
	} else {
		for __obf_bbfd30fcc08dc1bc.Kind() == reflect.Ptr {
			__obf_bbfd30fcc08dc1bc = __obf_bbfd30fcc08dc1bc.Elem()
			__obf_c4f748de4fd652e4 = __obf_f1f0aa6d8078637c(__obf_bbfd30fcc08dc1bc)
			__obf_815edc9179203b9a = __obf_66268eb3a3deccf5(__obf_bbfd30fcc08dc1bc)
		}
		if __obf_bbfd30fcc08dc1bc.Kind() != reflect.Struct {
			return false
		}
	}

	if reflect.ValueOf(__obf_c4f748de4fd652e4).Pointer() != __obf_b287ee13551609b5 {
		return false
	}
	if reflect.ValueOf(__obf_815edc9179203b9a).Pointer() != __obf_3bd50ef9ce0e4794 {
		return false
	}
	__obf_c3467fac7b886c06 := __obf_19dc6e60fa248a31(__obf_bbfd30fcc08dc1bc, __obf_990470bfaf220103).List
	for _, __obf_efecfc090178574b := range __obf_c3467fac7b886c06 {
		if _, __obf_81b19f2a6159ab89 := __obf_f82968c6203f3e5f.Map[__obf_efecfc090178574b.
			// Don't auto inline if there are shadowed fields.
			__obf_071b55c16b16fe35]; __obf_81b19f2a6159ab89 {

			return false
		}
	}

	for _, __obf_efecfc090178574b := range __obf_c3467fac7b886c06 {
		__obf_efecfc090178574b.__obf_eb6d7a2127ac443c = append(__obf_da4a2296298d318f.__obf_eb6d7a2127ac443c, __obf_efecfc090178574b.__obf_eb6d7a2127ac443c...)
		__obf_f82968c6203f3e5f.
			Add(__obf_efecfc090178574b)
	}
	return true
}

type __obf_2251d07325c68807 interface {
	IsZero() bool
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_9c83a9dd41bc3039(__obf_a1f43267eeb48a1a reflect.Value) bool {
	__obf_91ae1c24bc584c3a := __obf_a1f43267eeb48a1a.Kind()

	for __obf_91ae1c24bc584c3a == reflect.Interface {
		if __obf_a1f43267eeb48a1a.IsNil() {
			return true
		}
		__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Elem()
		__obf_91ae1c24bc584c3a = __obf_a1f43267eeb48a1a.Kind()
	}

	if __obf_a722bccc68b1b463, __obf_81b19f2a6159ab89 := __obf_a1f43267eeb48a1a.Interface().(__obf_2251d07325c68807); __obf_81b19f2a6159ab89 {
		return __obf_178efdbad2797673(__obf_91ae1c24bc584c3a) && __obf_a1f43267eeb48a1a.IsNil() || __obf_a722bccc68b1b463.IsZero()
	}

	switch __obf_91ae1c24bc584c3a {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_a1f43267eeb48a1a.Len() == 0
	case reflect.Struct:
		__obf_eb37ef6b7fc3d8f3 := __obf_bd827009b4ede0e2.Fields(__obf_a1f43267eeb48a1a.Type(), __obf_2c8e97779108ab17.__obf_72ba705ed504d210)
		__obf_0cf794214f02df4d := __obf_eb37ef6b7fc3d8f3.OmitEmpty(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a)
		return len(__obf_0cf794214f02df4d) == 0
	case reflect.Bool:
		return !__obf_a1f43267eeb48a1a.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_a1f43267eeb48a1a.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_a1f43267eeb48a1a.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_a1f43267eeb48a1a.Float() == 0
	case reflect.Ptr:
		return __obf_a1f43267eeb48a1a.IsNil()
	default:
		return false
	}
}

func __obf_01236718d8fc377c(__obf_a1f43267eeb48a1a reflect.Value, __obf_eb6d7a2127ac443c []int) (_ reflect.Value, __obf_81b19f2a6159ab89 bool) {
	if len(__obf_eb6d7a2127ac443c) == 1 {
		return __obf_a1f43267eeb48a1a.Field(__obf_eb6d7a2127ac443c[0]), true
	}

	for __obf_097d8b6061c9592b, __obf_82aee0190d456a88 := range __obf_eb6d7a2127ac443c {
		if __obf_097d8b6061c9592b > 0 {
			if __obf_a1f43267eeb48a1a.Kind() == reflect.Ptr {
				if __obf_a1f43267eeb48a1a.IsNil() {
					return __obf_a1f43267eeb48a1a, false
				}
				__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Elem()
			}
		}
		__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Field(__obf_82aee0190d456a88)
	}

	return __obf_a1f43267eeb48a1a, true
}

func __obf_0790f3ff4aeb54d2(__obf_a1f43267eeb48a1a reflect.Value, __obf_eb6d7a2127ac443c []int) reflect.Value {
	if len(__obf_eb6d7a2127ac443c) == 1 {
		return __obf_a1f43267eeb48a1a.Field(__obf_eb6d7a2127ac443c[0])
	}

	for __obf_097d8b6061c9592b, __obf_82aee0190d456a88 := range __obf_eb6d7a2127ac443c {
		if __obf_097d8b6061c9592b > 0 {
			var __obf_81b19f2a6159ab89 bool
			__obf_a1f43267eeb48a1a, __obf_81b19f2a6159ab89 = __obf_58bb2ae8a0702afc(__obf_a1f43267eeb48a1a)
			if !__obf_81b19f2a6159ab89 {
				return __obf_a1f43267eeb48a1a
			}
		}
		__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Field(__obf_82aee0190d456a88)
	}

	return __obf_a1f43267eeb48a1a
}

func __obf_58bb2ae8a0702afc(__obf_a1f43267eeb48a1a reflect.Value) (reflect.Value, bool) {
	if __obf_a1f43267eeb48a1a.Kind() == reflect.Ptr {
		if __obf_a1f43267eeb48a1a.IsNil() {
			if !__obf_a1f43267eeb48a1a.CanSet() {
				return __obf_a1f43267eeb48a1a, false
			}
			__obf_e0bd53df0007d029 := __obf_a1f43267eeb48a1a.Type().Elem()
			if __obf_e0bd53df0007d029.Kind() != reflect.Struct {
				return __obf_a1f43267eeb48a1a, false
			}
			__obf_a1f43267eeb48a1a.
				Set(__obf_5f0c2ccecff4d570(__obf_e0bd53df0007d029))
		}
		__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Elem()
	}
	return __obf_a1f43267eeb48a1a, true
}
