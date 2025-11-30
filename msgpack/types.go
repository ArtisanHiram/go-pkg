package __obf_fd770cb9675903b0

import (
	"encoding"
	"fmt"
	tagparser "github.com/ArtisanHiram/go-pkg/tagparser"
	"log"
	"reflect"
	"sync"
)

var __obf_f1f3ff99eed138e6 = reflect.TypeOf((*error)(nil)).Elem()

var (
	__obf_3dbe24e34d0a7693 = reflect.TypeOf((*CustomEncoder)(nil)).Elem()
	__obf_3370d035693fe504 = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
)

var (
	__obf_c3360c234e8a51cb = reflect.TypeOf((*Marshaler)(nil)).Elem()
	__obf_d8f141f48ff21664 = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var (
	__obf_6e4a841f6465176f = reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	__obf_4f3a83c066a65856 = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

var (
	__obf_2fa03b50688c3358 = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	__obf_5d3e3730f188b5b3 = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

type (
	__obf_dc120d370e164c91 func(*Encoder, reflect.Value) error
	__obf_fb10c6a4667e72a5 func(*Decoder, reflect.Value) error
)

var (
	__obf_56aff9039d970a50 sync.Map
	__obf_9ea14163dfc9e317 sync.Map
)

// Register registers encoder and decoder functions for a value.
// This is low level API and in most cases you should prefer implementing
// CustomEncoder/CustomDecoder or Marshaler/Unmarshaler interfaces.
func Register(__obf_28cbfc96ff0a56b6 any, __obf_a7b5a0ca650f48ee __obf_dc120d370e164c91, __obf_2f8f01810a02d049 __obf_fb10c6a4667e72a5) {
	__obf_8733059f76088ffc := reflect.TypeOf(__obf_28cbfc96ff0a56b6)
	if __obf_a7b5a0ca650f48ee != nil {
		__obf_56aff9039d970a50.
			Store(__obf_8733059f76088ffc, __obf_a7b5a0ca650f48ee)
	}
	if __obf_2f8f01810a02d049 != nil {
		__obf_9ea14163dfc9e317.
			Store(__obf_8733059f76088ffc,

				//------------------------------------------------------------------------------
				__obf_2f8f01810a02d049)
	}
}

const __obf_14ed3c13fed5c9eb = "msgpack"

var __obf_a93e575d0cc03783 = __obf_5e9ce64e8a4f202f()

type __obf_63a6a2ea3c258ea9 struct {
	__obf_777489ec8e6b2044 sync.Map
}

type __obf_c4ae3cab46f237c6 struct {
	__obf_8733059f76088ffc reflect.Type
	__obf_1a0f202964305730 string
}

func __obf_5e9ce64e8a4f202f() *__obf_63a6a2ea3c258ea9 {
	return new(__obf_63a6a2ea3c258ea9)
}

func (__obf_777489ec8e6b2044 *__obf_63a6a2ea3c258ea9) Fields(__obf_8733059f76088ffc reflect.Type, __obf_1a0f202964305730 string) *__obf_f92dfca778102077 {
	__obf_e155b5a95e20e61d := __obf_c4ae3cab46f237c6{__obf_1a0f202964305730: __obf_1a0f202964305730, __obf_8733059f76088ffc: __obf_8733059f76088ffc}

	if __obf_f328a048f76b7256, __obf_b00b3c6a10f90467 := __obf_777489ec8e6b2044.__obf_777489ec8e6b2044.Load(__obf_e155b5a95e20e61d); __obf_b00b3c6a10f90467 {
		return __obf_f328a048f76b7256.(*__obf_f92dfca778102077)
	}
	__obf_799ddc75e5087ecd := __obf_9ae346426935aa8f(__obf_8733059f76088ffc, __obf_1a0f202964305730)
	__obf_777489ec8e6b2044.__obf_777489ec8e6b2044.
		Store(__obf_e155b5a95e20e61d, __obf_799ddc75e5087ecd)

	return __obf_799ddc75e5087ecd
}

//------------------------------------------------------------------------------

type __obf_9b29490b522018e0 struct {
	__obf_ff15dfff9c06a391 __obf_dc120d370e164c91
	__obf_25e78266d5ba1944 __obf_fb10c6a4667e72a5
	__obf_2a5a18f2bc2906b2 string
	__obf_cf10c279d0cc414b []int
	__obf_d76ec37d0f70b6af bool
}

func (__obf_e25b9e550ea05af3 *__obf_9b29490b522018e0) Omit(__obf_e9038cda3b5cf711 *Encoder, __obf_475491b3031fb8b4 reflect.Value) bool {
	__obf_f328a048f76b7256, __obf_b00b3c6a10f90467 := __obf_f85a3cc5ca91e5d8(__obf_475491b3031fb8b4, __obf_e25b9e550ea05af3.__obf_cf10c279d0cc414b)
	if !__obf_b00b3c6a10f90467 {
		return true
	}
	__obf_5b67d1da8d6573b6 := __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_bb9969de1dccd587 != 0
	return (__obf_e25b9e550ea05af3.__obf_d76ec37d0f70b6af || __obf_5b67d1da8d6573b6) && __obf_e9038cda3b5cf711.__obf_b23e0f4a3ef9c482(__obf_f328a048f76b7256)
}

func (__obf_e25b9e550ea05af3 *__obf_9b29490b522018e0) EncodeValue(__obf_e9038cda3b5cf711 *Encoder, __obf_475491b3031fb8b4 reflect.Value) error {
	__obf_f328a048f76b7256, __obf_b00b3c6a10f90467 := __obf_f85a3cc5ca91e5d8(__obf_475491b3031fb8b4, __obf_e25b9e550ea05af3.__obf_cf10c279d0cc414b)
	if !__obf_b00b3c6a10f90467 {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	return __obf_e25b9e550ea05af3.__obf_ff15dfff9c06a391(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256)
}

func (__obf_e25b9e550ea05af3 *__obf_9b29490b522018e0) DecodeValue(__obf_5d389b660eefb08c *Decoder, __obf_475491b3031fb8b4 reflect.Value) error {
	__obf_f328a048f76b7256 := __obf_20464f3c855bf186(__obf_475491b3031fb8b4, __obf_e25b9e550ea05af3.

		//------------------------------------------------------------------------------
		__obf_cf10c279d0cc414b)
	return __obf_e25b9e550ea05af3.__obf_25e78266d5ba1944(__obf_5d389b660eefb08c, __obf_f328a048f76b7256)
}

type __obf_f92dfca778102077 struct {
	Type                   reflect.Type
	Map                    map[string]*__obf_9b29490b522018e0
	List                   []*__obf_9b29490b522018e0
	AsArray                bool
	__obf_d413c26526713c3c bool
}

func __obf_112077ed9ca8e25c(__obf_8733059f76088ffc reflect.Type) *__obf_f92dfca778102077 {
	return &__obf_f92dfca778102077{
		Type: __obf_8733059f76088ffc,
		Map:  make(map[string]*__obf_9b29490b522018e0, __obf_8733059f76088ffc.NumField()),
		List: make([]*__obf_9b29490b522018e0, 0, __obf_8733059f76088ffc.NumField()),
	}
}

func (__obf_799ddc75e5087ecd *__obf_f92dfca778102077) Add(__obf_9b29490b522018e0 *__obf_9b29490b522018e0) {
	__obf_799ddc75e5087ecd.__obf_70e7ac9ada6db324(__obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2)
	__obf_799ddc75e5087ecd.
		Map[__obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2] = __obf_9b29490b522018e0
	__obf_799ddc75e5087ecd.
		List = append(__obf_799ddc75e5087ecd.List, __obf_9b29490b522018e0)
	if __obf_9b29490b522018e0.__obf_d76ec37d0f70b6af {
		__obf_799ddc75e5087ecd.__obf_d413c26526713c3c = true
	}
}

func (__obf_799ddc75e5087ecd *__obf_f92dfca778102077) __obf_70e7ac9ada6db324(__obf_2a5a18f2bc2906b2 string) {
	if _, __obf_b00b3c6a10f90467 := __obf_799ddc75e5087ecd.Map[__obf_2a5a18f2bc2906b2]; __obf_b00b3c6a10f90467 {
		log.Printf("msgpack: %s already has field=%s", __obf_799ddc75e5087ecd.Type, __obf_2a5a18f2bc2906b2)
	}
}

func (__obf_799ddc75e5087ecd *__obf_f92dfca778102077) OmitEmpty(__obf_e9038cda3b5cf711 *Encoder, __obf_475491b3031fb8b4 reflect.Value) []*__obf_9b29490b522018e0 {
	__obf_5b67d1da8d6573b6 := __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_bb9969de1dccd587 != 0
	if !__obf_799ddc75e5087ecd.__obf_d413c26526713c3c && !__obf_5b67d1da8d6573b6 {
		return __obf_799ddc75e5087ecd.List
	}
	__obf_f92dfca778102077 := make([]*__obf_9b29490b522018e0, 0, len(__obf_799ddc75e5087ecd.List))

	for _, __obf_e25b9e550ea05af3 := range __obf_799ddc75e5087ecd.List {
		if !__obf_e25b9e550ea05af3.Omit(__obf_e9038cda3b5cf711, __obf_475491b3031fb8b4) {
			__obf_f92dfca778102077 = append(__obf_f92dfca778102077, __obf_e25b9e550ea05af3)
		}
	}

	return __obf_f92dfca778102077
}

func __obf_9ae346426935aa8f(__obf_8733059f76088ffc reflect.Type, __obf_05e0cc4d17f5f67f string) *__obf_f92dfca778102077 {
	__obf_799ddc75e5087ecd := __obf_112077ed9ca8e25c(__obf_8733059f76088ffc)

	var __obf_d76ec37d0f70b6af bool
	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_8733059f76088ffc.NumField(); __obf_4140b3ff04f75a36++ {
		__obf_e25b9e550ea05af3 := __obf_8733059f76088ffc.Field(__obf_4140b3ff04f75a36)
		__obf_cc400fa0937151e4 := __obf_e25b9e550ea05af3.Tag.Get(__obf_14ed3c13fed5c9eb)
		if __obf_cc400fa0937151e4 == "" && __obf_05e0cc4d17f5f67f != "" {
			__obf_cc400fa0937151e4 = __obf_e25b9e550ea05af3.Tag.Get(__obf_05e0cc4d17f5f67f)
		}
		__obf_1a0f202964305730 := tagparser.Parse(__obf_cc400fa0937151e4)
		if __obf_1a0f202964305730.Name == "-" {
			continue
		}

		if __obf_e25b9e550ea05af3.Name == "_msgpack" {
			__obf_799ddc75e5087ecd.
				AsArray = __obf_1a0f202964305730.HasOption("as_array") || __obf_1a0f202964305730.HasOption("asArray")
			if __obf_1a0f202964305730.HasOption("omitempty") {
				__obf_d76ec37d0f70b6af = true
			}
		}

		if __obf_e25b9e550ea05af3.PkgPath != "" && !__obf_e25b9e550ea05af3.Anonymous {
			continue
		}
		__obf_9b29490b522018e0 := &__obf_9b29490b522018e0{__obf_2a5a18f2bc2906b2: __obf_1a0f202964305730.Name, __obf_cf10c279d0cc414b: __obf_e25b9e550ea05af3.Index, __obf_d76ec37d0f70b6af: __obf_d76ec37d0f70b6af || __obf_1a0f202964305730.HasOption("omitempty")}

		if __obf_1a0f202964305730.HasOption("intern") {
			switch __obf_e25b9e550ea05af3.Type.Kind() {
			case reflect.Interface:
				__obf_9b29490b522018e0.__obf_ff15dfff9c06a391 = __obf_db4fdb7a2953f589
				__obf_9b29490b522018e0.__obf_25e78266d5ba1944 = __obf_2ae309e0ccdf900e
			case reflect.String:
				__obf_9b29490b522018e0.__obf_ff15dfff9c06a391 = __obf_b1eb723829a575f6
				__obf_9b29490b522018e0.__obf_25e78266d5ba1944 = __obf_7d869d41e2fef94d
			default:
				__obf_45342a3a754d12f5 := fmt.Errorf("msgpack: intern strings are not supported on %s", __obf_e25b9e550ea05af3.Type)
				panic(__obf_45342a3a754d12f5)
			}
		} else {
			__obf_9b29490b522018e0.__obf_ff15dfff9c06a391 = __obf_031e3a13e30d95dc(__obf_e25b9e550ea05af3.Type)
			__obf_9b29490b522018e0.__obf_25e78266d5ba1944 = __obf_538e4e278f06c9d6(__obf_e25b9e550ea05af3.Type)
		}

		if __obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2 == "" {
			__obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2 = __obf_e25b9e550ea05af3.Name
		}

		if __obf_e25b9e550ea05af3.Anonymous && !__obf_1a0f202964305730.HasOption("noinline") {
			__obf_1b4779a6f4ab4d79 := __obf_1a0f202964305730.HasOption("inline")
			if __obf_1b4779a6f4ab4d79 {
				__obf_39c2397fc5239421(__obf_799ddc75e5087ecd, __obf_e25b9e550ea05af3.Type, __obf_9b29490b522018e0, __obf_05e0cc4d17f5f67f)
			} else {
				__obf_1b4779a6f4ab4d79 = __obf_8b29f5c5ccc8e9b0(__obf_799ddc75e5087ecd, __obf_e25b9e550ea05af3.Type, __obf_9b29490b522018e0, __obf_05e0cc4d17f5f67f)
			}

			if __obf_1b4779a6f4ab4d79 {
				if _, __obf_b00b3c6a10f90467 := __obf_799ddc75e5087ecd.Map[__obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2]; __obf_b00b3c6a10f90467 {
					log.Printf("msgpack: %s already has field=%s", __obf_799ddc75e5087ecd.Type, __obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2)
				}
				__obf_799ddc75e5087ecd.
					Map[__obf_9b29490b522018e0.__obf_2a5a18f2bc2906b2] = __obf_9b29490b522018e0
				continue
			}
		}
		__obf_799ddc75e5087ecd.
			Add(__obf_9b29490b522018e0)

		if __obf_2603f980e982f0d4, __obf_b00b3c6a10f90467 := __obf_1a0f202964305730.Options["alias"]; __obf_b00b3c6a10f90467 {
			__obf_799ddc75e5087ecd.__obf_70e7ac9ada6db324(__obf_2603f980e982f0d4)
			__obf_799ddc75e5087ecd.
				Map[__obf_2603f980e982f0d4] = __obf_9b29490b522018e0
		}
	}
	return __obf_799ddc75e5087ecd
}

var (
	__obf_cf743c61808b5ca6 uintptr
	__obf_688bf8f319c824f9 uintptr
)

//nolint:gochecknoinits
func init() {
	__obf_cf743c61808b5ca6 = reflect.ValueOf(__obf_c52bd37d7024764c).Pointer()
	__obf_688bf8f319c824f9 = reflect.ValueOf(__obf_9fde65ded02d475f).Pointer()
}

func __obf_39c2397fc5239421(__obf_799ddc75e5087ecd *__obf_f92dfca778102077, __obf_8733059f76088ffc reflect.Type, __obf_e25b9e550ea05af3 *__obf_9b29490b522018e0, __obf_1a0f202964305730 string) {
	__obf_40c72be3f25e9213 := __obf_9ae346426935aa8f(__obf_8733059f76088ffc, __obf_1a0f202964305730).List
	for _, __obf_9b29490b522018e0 := range __obf_40c72be3f25e9213 {
		if _, __obf_b00b3c6a10f90467 := __obf_799ddc75e5087ecd.Map[__obf_9b29490b522018e0.
			// Don't inline shadowed fields.
			__obf_2a5a18f2bc2906b2]; __obf_b00b3c6a10f90467 {

			continue
		}
		__obf_9b29490b522018e0.__obf_cf10c279d0cc414b = append(__obf_e25b9e550ea05af3.__obf_cf10c279d0cc414b, __obf_9b29490b522018e0.__obf_cf10c279d0cc414b...)
		__obf_799ddc75e5087ecd.
			Add(__obf_9b29490b522018e0)
	}
}

func __obf_8b29f5c5ccc8e9b0(__obf_799ddc75e5087ecd *__obf_f92dfca778102077, __obf_8733059f76088ffc reflect.Type, __obf_e25b9e550ea05af3 *__obf_9b29490b522018e0, __obf_1a0f202964305730 string) bool {
	var __obf_ff15dfff9c06a391 __obf_dc120d370e164c91
	var __obf_25e78266d5ba1944 __obf_fb10c6a4667e72a5

	if __obf_8733059f76088ffc.Kind() == reflect.Struct {
		__obf_ff15dfff9c06a391 = __obf_e25b9e550ea05af3.__obf_ff15dfff9c06a391
		__obf_25e78266d5ba1944 = __obf_e25b9e550ea05af3.__obf_25e78266d5ba1944
	} else {
		for __obf_8733059f76088ffc.Kind() == reflect.Ptr {
			__obf_8733059f76088ffc = __obf_8733059f76088ffc.Elem()
			__obf_ff15dfff9c06a391 = __obf_031e3a13e30d95dc(__obf_8733059f76088ffc)
			__obf_25e78266d5ba1944 = __obf_538e4e278f06c9d6(__obf_8733059f76088ffc)
		}
		if __obf_8733059f76088ffc.Kind() != reflect.Struct {
			return false
		}
	}

	if reflect.ValueOf(__obf_ff15dfff9c06a391).Pointer() != __obf_cf743c61808b5ca6 {
		return false
	}
	if reflect.ValueOf(__obf_25e78266d5ba1944).Pointer() != __obf_688bf8f319c824f9 {
		return false
	}
	__obf_40c72be3f25e9213 := __obf_9ae346426935aa8f(__obf_8733059f76088ffc, __obf_1a0f202964305730).List
	for _, __obf_9b29490b522018e0 := range __obf_40c72be3f25e9213 {
		if _, __obf_b00b3c6a10f90467 := __obf_799ddc75e5087ecd.Map[__obf_9b29490b522018e0.
			// Don't auto inline if there are shadowed fields.
			__obf_2a5a18f2bc2906b2]; __obf_b00b3c6a10f90467 {

			return false
		}
	}

	for _, __obf_9b29490b522018e0 := range __obf_40c72be3f25e9213 {
		__obf_9b29490b522018e0.__obf_cf10c279d0cc414b = append(__obf_e25b9e550ea05af3.__obf_cf10c279d0cc414b, __obf_9b29490b522018e0.__obf_cf10c279d0cc414b...)
		__obf_799ddc75e5087ecd.
			Add(__obf_9b29490b522018e0)
	}
	return true
}

type __obf_6dab85375f599b5b interface {
	IsZero() bool
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_b23e0f4a3ef9c482(__obf_f328a048f76b7256 reflect.Value) bool {
	__obf_5cd51d276078fe9d := __obf_f328a048f76b7256.Kind()

	for __obf_5cd51d276078fe9d == reflect.Interface {
		if __obf_f328a048f76b7256.IsNil() {
			return true
		}
		__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Elem()
		__obf_5cd51d276078fe9d = __obf_f328a048f76b7256.Kind()
	}

	if __obf_590821fc429a58f8, __obf_b00b3c6a10f90467 := __obf_f328a048f76b7256.Interface().(__obf_6dab85375f599b5b); __obf_b00b3c6a10f90467 {
		return __obf_066825d3e1771287(__obf_5cd51d276078fe9d) && __obf_f328a048f76b7256.IsNil() || __obf_590821fc429a58f8.IsZero()
	}

	switch __obf_5cd51d276078fe9d {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_f328a048f76b7256.Len() == 0
	case reflect.Struct:
		__obf_7d081654985e1a56 := __obf_a93e575d0cc03783.Fields(__obf_f328a048f76b7256.Type(), __obf_e9038cda3b5cf711.__obf_6621ba3773b8696a)
		__obf_f92dfca778102077 := __obf_7d081654985e1a56.OmitEmpty(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256)
		return len(__obf_f92dfca778102077) == 0
	case reflect.Bool:
		return !__obf_f328a048f76b7256.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_f328a048f76b7256.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_f328a048f76b7256.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_f328a048f76b7256.Float() == 0
	case reflect.Ptr:
		return __obf_f328a048f76b7256.IsNil()
	default:
		return false
	}
}

func __obf_f85a3cc5ca91e5d8(__obf_f328a048f76b7256 reflect.Value, __obf_cf10c279d0cc414b []int) (_ reflect.Value, __obf_b00b3c6a10f90467 bool) {
	if len(__obf_cf10c279d0cc414b) == 1 {
		return __obf_f328a048f76b7256.Field(__obf_cf10c279d0cc414b[0]), true
	}

	for __obf_4140b3ff04f75a36, __obf_261dffcd6ce7f08f := range __obf_cf10c279d0cc414b {
		if __obf_4140b3ff04f75a36 > 0 {
			if __obf_f328a048f76b7256.Kind() == reflect.Ptr {
				if __obf_f328a048f76b7256.IsNil() {
					return __obf_f328a048f76b7256, false
				}
				__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Elem()
			}
		}
		__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Field(__obf_261dffcd6ce7f08f)
	}

	return __obf_f328a048f76b7256, true
}

func __obf_20464f3c855bf186(__obf_f328a048f76b7256 reflect.Value, __obf_cf10c279d0cc414b []int) reflect.Value {
	if len(__obf_cf10c279d0cc414b) == 1 {
		return __obf_f328a048f76b7256.Field(__obf_cf10c279d0cc414b[0])
	}

	for __obf_4140b3ff04f75a36, __obf_261dffcd6ce7f08f := range __obf_cf10c279d0cc414b {
		if __obf_4140b3ff04f75a36 > 0 {
			var __obf_b00b3c6a10f90467 bool
			__obf_f328a048f76b7256, __obf_b00b3c6a10f90467 = __obf_7af237286541ecbf(__obf_f328a048f76b7256)
			if !__obf_b00b3c6a10f90467 {
				return __obf_f328a048f76b7256
			}
		}
		__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Field(__obf_261dffcd6ce7f08f)
	}

	return __obf_f328a048f76b7256
}

func __obf_7af237286541ecbf(__obf_f328a048f76b7256 reflect.Value) (reflect.Value, bool) {
	if __obf_f328a048f76b7256.Kind() == reflect.Ptr {
		if __obf_f328a048f76b7256.IsNil() {
			if !__obf_f328a048f76b7256.CanSet() {
				return __obf_f328a048f76b7256, false
			}
			__obf_1a825469f5758f5f := __obf_f328a048f76b7256.Type().Elem()
			if __obf_1a825469f5758f5f.Kind() != reflect.Struct {
				return __obf_f328a048f76b7256, false
			}
			__obf_f328a048f76b7256.
				Set(__obf_ed787a5d7a207071(__obf_1a825469f5758f5f))
		}
		__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Elem()
	}
	return __obf_f328a048f76b7256, true
}
