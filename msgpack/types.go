package __obf_de86cdc8ae98b45b

import (
	"encoding"
	"fmt"
	tagparser "github.com/ArtisanHiram/go-pkg/tagparser"
	"log"
	"reflect"
	"sync"
)

var __obf_2d0560f267da57d8 = reflect.TypeOf((*error)(nil)).Elem()

var (
	__obf_97b6aca480193051 = reflect.TypeOf((*CustomEncoder)(nil)).Elem()
	__obf_bb521c0d47017b87 = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
)

var (
	__obf_38acb3bdfb404f10 = reflect.TypeOf((*Marshaler)(nil)).Elem()
	__obf_46f8dad60f3aa967 = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var (
	__obf_6ff2001fa5096e39 = reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	__obf_cd3378d776026b55 = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

var (
	__obf_e9715cc721daa9f1 = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	__obf_0d5a1d97d27c066f = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

type (
	__obf_9399ab762e15c2a7 func(*Encoder, reflect.Value) error
	__obf_b74c57495471e351 func(*Decoder, reflect.Value) error
)

var (
	__obf_1d13df182269d8e8 sync.Map
	__obf_9f46fbaef88684c0 sync.Map
)

// Register registers encoder and decoder functions for a value.
// This is low level API and in most cases you should prefer implementing
// CustomEncoder/CustomDecoder or Marshaler/Unmarshaler interfaces.
func Register(__obf_1926a1e373f9e647 any, __obf_6a6518f236eeec6e __obf_9399ab762e15c2a7, __obf_9461956859459a44 __obf_b74c57495471e351) {
	__obf_2dea2b2773c8afdf := reflect.TypeOf(__obf_1926a1e373f9e647)
	if __obf_6a6518f236eeec6e != nil {
		__obf_1d13df182269d8e8.
			Store(__obf_2dea2b2773c8afdf, __obf_6a6518f236eeec6e)
	}
	if __obf_9461956859459a44 != nil {
		__obf_9f46fbaef88684c0.
			Store(__obf_2dea2b2773c8afdf,

				//------------------------------------------------------------------------------
				__obf_9461956859459a44)
	}
}

const __obf_52ee8cc3b64f8557 = "msgpack"

var __obf_4642342e308ea07b = __obf_1d70c78b4fe4542a()

type __obf_96e0d1b4897c9420 struct {
	__obf_2a26e7104b4c4373 sync.Map
}

type __obf_b688c906327e2fb3 struct {
	__obf_2dea2b2773c8afdf reflect.Type
	__obf_ee8ab0888686da67 string
}

func __obf_1d70c78b4fe4542a() *__obf_96e0d1b4897c9420 {
	return new(__obf_96e0d1b4897c9420)
}

func (__obf_2a26e7104b4c4373 *__obf_96e0d1b4897c9420) Fields(__obf_2dea2b2773c8afdf reflect.Type, __obf_ee8ab0888686da67 string) *__obf_7553cd24ebfd7da5 {
	__obf_9cab2959bb95a876 := __obf_b688c906327e2fb3{__obf_ee8ab0888686da67: __obf_ee8ab0888686da67, __obf_2dea2b2773c8afdf: __obf_2dea2b2773c8afdf}

	if __obf_17732590722140e7, __obf_77cfff95beb3cc99 := __obf_2a26e7104b4c4373.__obf_2a26e7104b4c4373.Load(__obf_9cab2959bb95a876); __obf_77cfff95beb3cc99 {
		return __obf_17732590722140e7.(*__obf_7553cd24ebfd7da5)
	}
	__obf_512c4f2f38a27036 := __obf_a1afea8ea1650215(__obf_2dea2b2773c8afdf, __obf_ee8ab0888686da67)
	__obf_2a26e7104b4c4373.__obf_2a26e7104b4c4373.
		Store(__obf_9cab2959bb95a876, __obf_512c4f2f38a27036)

	return __obf_512c4f2f38a27036
}

//------------------------------------------------------------------------------

type __obf_4ee6afc37d3486aa struct {
	__obf_2981aff99465546b __obf_9399ab762e15c2a7
	__obf_0fe2a689830f8985 __obf_b74c57495471e351
	__obf_d7fb5470f7183af5 string
	__obf_8671df0a9132a879 []int
	__obf_2886bcbe3c68d1ad bool
}

func (__obf_cf8245c3f9570af5 *__obf_4ee6afc37d3486aa) Omit(__obf_7bae0b613da60dd3 *Encoder, __obf_84ee659cd19e5239 reflect.Value) bool {
	__obf_17732590722140e7, __obf_77cfff95beb3cc99 := __obf_3000a20be09182df(__obf_84ee659cd19e5239, __obf_cf8245c3f9570af5.__obf_8671df0a9132a879)
	if !__obf_77cfff95beb3cc99 {
		return true
	}
	__obf_6688e73b147a3578 := __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_eadbff0f95d7376e != 0
	return (__obf_cf8245c3f9570af5.__obf_2886bcbe3c68d1ad || __obf_6688e73b147a3578) && __obf_7bae0b613da60dd3.__obf_2474e494c3416afb(__obf_17732590722140e7)
}

func (__obf_cf8245c3f9570af5 *__obf_4ee6afc37d3486aa) EncodeValue(__obf_7bae0b613da60dd3 *Encoder, __obf_84ee659cd19e5239 reflect.Value) error {
	__obf_17732590722140e7, __obf_77cfff95beb3cc99 := __obf_3000a20be09182df(__obf_84ee659cd19e5239, __obf_cf8245c3f9570af5.__obf_8671df0a9132a879)
	if !__obf_77cfff95beb3cc99 {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	return __obf_cf8245c3f9570af5.__obf_2981aff99465546b(__obf_7bae0b613da60dd3, __obf_17732590722140e7)
}

func (__obf_cf8245c3f9570af5 *__obf_4ee6afc37d3486aa) DecodeValue(__obf_dcaad1165bb07af9 *Decoder, __obf_84ee659cd19e5239 reflect.Value) error {
	__obf_17732590722140e7 := __obf_3af645acd1a85372(__obf_84ee659cd19e5239, __obf_cf8245c3f9570af5.

		//------------------------------------------------------------------------------
		__obf_8671df0a9132a879)
	return __obf_cf8245c3f9570af5.__obf_0fe2a689830f8985(__obf_dcaad1165bb07af9, __obf_17732590722140e7)
}

type __obf_7553cd24ebfd7da5 struct {
	Type                   reflect.Type
	Map                    map[string]*__obf_4ee6afc37d3486aa
	List                   []*__obf_4ee6afc37d3486aa
	AsArray                bool
	__obf_28bf675512406c00 bool
}

func __obf_c8552f1ad413965e(__obf_2dea2b2773c8afdf reflect.Type) *__obf_7553cd24ebfd7da5 {
	return &__obf_7553cd24ebfd7da5{
		Type: __obf_2dea2b2773c8afdf,
		Map:  make(map[string]*__obf_4ee6afc37d3486aa, __obf_2dea2b2773c8afdf.NumField()),
		List: make([]*__obf_4ee6afc37d3486aa, 0, __obf_2dea2b2773c8afdf.NumField()),
	}
}

func (__obf_512c4f2f38a27036 *__obf_7553cd24ebfd7da5) Add(__obf_4ee6afc37d3486aa *__obf_4ee6afc37d3486aa) {
	__obf_512c4f2f38a27036.__obf_34e09f7112e593d4(__obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5)
	__obf_512c4f2f38a27036.
		Map[__obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5] = __obf_4ee6afc37d3486aa
	__obf_512c4f2f38a27036.
		List = append(__obf_512c4f2f38a27036.List, __obf_4ee6afc37d3486aa)
	if __obf_4ee6afc37d3486aa.__obf_2886bcbe3c68d1ad {
		__obf_512c4f2f38a27036.__obf_28bf675512406c00 = true
	}
}

func (__obf_512c4f2f38a27036 *__obf_7553cd24ebfd7da5) __obf_34e09f7112e593d4(__obf_d7fb5470f7183af5 string) {
	if _, __obf_77cfff95beb3cc99 := __obf_512c4f2f38a27036.Map[__obf_d7fb5470f7183af5]; __obf_77cfff95beb3cc99 {
		log.Printf("msgpack: %s already has field=%s", __obf_512c4f2f38a27036.Type, __obf_d7fb5470f7183af5)
	}
}

func (__obf_512c4f2f38a27036 *__obf_7553cd24ebfd7da5) OmitEmpty(__obf_7bae0b613da60dd3 *Encoder, __obf_84ee659cd19e5239 reflect.Value) []*__obf_4ee6afc37d3486aa {
	__obf_6688e73b147a3578 := __obf_7bae0b613da60dd3.__obf_a15a642720e39c3e&__obf_eadbff0f95d7376e != 0
	if !__obf_512c4f2f38a27036.__obf_28bf675512406c00 && !__obf_6688e73b147a3578 {
		return __obf_512c4f2f38a27036.List
	}
	__obf_7553cd24ebfd7da5 := make([]*__obf_4ee6afc37d3486aa, 0, len(__obf_512c4f2f38a27036.List))

	for _, __obf_cf8245c3f9570af5 := range __obf_512c4f2f38a27036.List {
		if !__obf_cf8245c3f9570af5.Omit(__obf_7bae0b613da60dd3, __obf_84ee659cd19e5239) {
			__obf_7553cd24ebfd7da5 = append(__obf_7553cd24ebfd7da5, __obf_cf8245c3f9570af5)
		}
	}

	return __obf_7553cd24ebfd7da5
}

func __obf_a1afea8ea1650215(__obf_2dea2b2773c8afdf reflect.Type, __obf_1baa78958a35813b string) *__obf_7553cd24ebfd7da5 {
	__obf_512c4f2f38a27036 := __obf_c8552f1ad413965e(__obf_2dea2b2773c8afdf)

	var __obf_2886bcbe3c68d1ad bool
	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_2dea2b2773c8afdf.NumField(); __obf_101eec84c8338296++ {
		__obf_cf8245c3f9570af5 := __obf_2dea2b2773c8afdf.Field(__obf_101eec84c8338296)
		__obf_c5846980969247b0 := __obf_cf8245c3f9570af5.Tag.Get(__obf_52ee8cc3b64f8557)
		if __obf_c5846980969247b0 == "" && __obf_1baa78958a35813b != "" {
			__obf_c5846980969247b0 = __obf_cf8245c3f9570af5.Tag.Get(__obf_1baa78958a35813b)
		}
		__obf_ee8ab0888686da67 := tagparser.Parse(__obf_c5846980969247b0)
		if __obf_ee8ab0888686da67.Name == "-" {
			continue
		}

		if __obf_cf8245c3f9570af5.Name == "_msgpack" {
			__obf_512c4f2f38a27036.
				AsArray = __obf_ee8ab0888686da67.HasOption("as_array") || __obf_ee8ab0888686da67.HasOption("asArray")
			if __obf_ee8ab0888686da67.HasOption("omitempty") {
				__obf_2886bcbe3c68d1ad = true
			}
		}

		if __obf_cf8245c3f9570af5.PkgPath != "" && !__obf_cf8245c3f9570af5.Anonymous {
			continue
		}
		__obf_4ee6afc37d3486aa := &__obf_4ee6afc37d3486aa{__obf_d7fb5470f7183af5: __obf_ee8ab0888686da67.Name, __obf_8671df0a9132a879: __obf_cf8245c3f9570af5.Index, __obf_2886bcbe3c68d1ad: __obf_2886bcbe3c68d1ad || __obf_ee8ab0888686da67.HasOption("omitempty")}

		if __obf_ee8ab0888686da67.HasOption("intern") {
			switch __obf_cf8245c3f9570af5.Type.Kind() {
			case reflect.Interface:
				__obf_4ee6afc37d3486aa.__obf_2981aff99465546b = __obf_c4de622e12ad86a8
				__obf_4ee6afc37d3486aa.__obf_0fe2a689830f8985 = __obf_06166b79d4faa976
			case reflect.String:
				__obf_4ee6afc37d3486aa.__obf_2981aff99465546b = __obf_f85540ca8b34fdcf
				__obf_4ee6afc37d3486aa.__obf_0fe2a689830f8985 = __obf_0e872b4337eae851
			default:
				__obf_0feb3528b7b709ec := fmt.Errorf("msgpack: intern strings are not supported on %s", __obf_cf8245c3f9570af5.Type)
				panic(__obf_0feb3528b7b709ec)
			}
		} else {
			__obf_4ee6afc37d3486aa.__obf_2981aff99465546b = __obf_d4a1e60c459c24e4(__obf_cf8245c3f9570af5.Type)
			__obf_4ee6afc37d3486aa.__obf_0fe2a689830f8985 = __obf_25e3584cea7b6666(__obf_cf8245c3f9570af5.Type)
		}

		if __obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5 == "" {
			__obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5 = __obf_cf8245c3f9570af5.Name
		}

		if __obf_cf8245c3f9570af5.Anonymous && !__obf_ee8ab0888686da67.HasOption("noinline") {
			__obf_28b90aa7c659546c := __obf_ee8ab0888686da67.HasOption("inline")
			if __obf_28b90aa7c659546c {
				__obf_92e165aeca28f6dc(__obf_512c4f2f38a27036, __obf_cf8245c3f9570af5.Type, __obf_4ee6afc37d3486aa, __obf_1baa78958a35813b)
			} else {
				__obf_28b90aa7c659546c = __obf_040a352a98474ebd(__obf_512c4f2f38a27036, __obf_cf8245c3f9570af5.Type, __obf_4ee6afc37d3486aa, __obf_1baa78958a35813b)
			}

			if __obf_28b90aa7c659546c {
				if _, __obf_77cfff95beb3cc99 := __obf_512c4f2f38a27036.Map[__obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5]; __obf_77cfff95beb3cc99 {
					log.Printf("msgpack: %s already has field=%s", __obf_512c4f2f38a27036.Type, __obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5)
				}
				__obf_512c4f2f38a27036.
					Map[__obf_4ee6afc37d3486aa.__obf_d7fb5470f7183af5] = __obf_4ee6afc37d3486aa
				continue
			}
		}
		__obf_512c4f2f38a27036.
			Add(__obf_4ee6afc37d3486aa)

		if __obf_0f1eb6b9a3261d1a, __obf_77cfff95beb3cc99 := __obf_ee8ab0888686da67.Options["alias"]; __obf_77cfff95beb3cc99 {
			__obf_512c4f2f38a27036.__obf_34e09f7112e593d4(__obf_0f1eb6b9a3261d1a)
			__obf_512c4f2f38a27036.
				Map[__obf_0f1eb6b9a3261d1a] = __obf_4ee6afc37d3486aa
		}
	}
	return __obf_512c4f2f38a27036
}

var (
	__obf_813f2feab27667bd uintptr
	__obf_5a4deacae7a97a44 uintptr
)

//nolint:gochecknoinits
func init() {
	__obf_813f2feab27667bd = reflect.ValueOf(__obf_95de04fb2a8b0af3).Pointer()
	__obf_5a4deacae7a97a44 = reflect.ValueOf(__obf_d6cbe6ba2b1343be).Pointer()
}

func __obf_92e165aeca28f6dc(__obf_512c4f2f38a27036 *__obf_7553cd24ebfd7da5, __obf_2dea2b2773c8afdf reflect.Type, __obf_cf8245c3f9570af5 *__obf_4ee6afc37d3486aa, __obf_ee8ab0888686da67 string) {
	__obf_d89d6fd0b094b410 := __obf_a1afea8ea1650215(__obf_2dea2b2773c8afdf, __obf_ee8ab0888686da67).List
	for _, __obf_4ee6afc37d3486aa := range __obf_d89d6fd0b094b410 {
		if _, __obf_77cfff95beb3cc99 := __obf_512c4f2f38a27036.Map[__obf_4ee6afc37d3486aa.
			// Don't inline shadowed fields.
			__obf_d7fb5470f7183af5]; __obf_77cfff95beb3cc99 {

			continue
		}
		__obf_4ee6afc37d3486aa.__obf_8671df0a9132a879 = append(__obf_cf8245c3f9570af5.__obf_8671df0a9132a879, __obf_4ee6afc37d3486aa.__obf_8671df0a9132a879...)
		__obf_512c4f2f38a27036.
			Add(__obf_4ee6afc37d3486aa)
	}
}

func __obf_040a352a98474ebd(__obf_512c4f2f38a27036 *__obf_7553cd24ebfd7da5, __obf_2dea2b2773c8afdf reflect.Type, __obf_cf8245c3f9570af5 *__obf_4ee6afc37d3486aa, __obf_ee8ab0888686da67 string) bool {
	var __obf_2981aff99465546b __obf_9399ab762e15c2a7
	var __obf_0fe2a689830f8985 __obf_b74c57495471e351

	if __obf_2dea2b2773c8afdf.Kind() == reflect.Struct {
		__obf_2981aff99465546b = __obf_cf8245c3f9570af5.__obf_2981aff99465546b
		__obf_0fe2a689830f8985 = __obf_cf8245c3f9570af5.__obf_0fe2a689830f8985
	} else {
		for __obf_2dea2b2773c8afdf.Kind() == reflect.Ptr {
			__obf_2dea2b2773c8afdf = __obf_2dea2b2773c8afdf.Elem()
			__obf_2981aff99465546b = __obf_d4a1e60c459c24e4(__obf_2dea2b2773c8afdf)
			__obf_0fe2a689830f8985 = __obf_25e3584cea7b6666(__obf_2dea2b2773c8afdf)
		}
		if __obf_2dea2b2773c8afdf.Kind() != reflect.Struct {
			return false
		}
	}

	if reflect.ValueOf(__obf_2981aff99465546b).Pointer() != __obf_813f2feab27667bd {
		return false
	}
	if reflect.ValueOf(__obf_0fe2a689830f8985).Pointer() != __obf_5a4deacae7a97a44 {
		return false
	}
	__obf_d89d6fd0b094b410 := __obf_a1afea8ea1650215(__obf_2dea2b2773c8afdf, __obf_ee8ab0888686da67).List
	for _, __obf_4ee6afc37d3486aa := range __obf_d89d6fd0b094b410 {
		if _, __obf_77cfff95beb3cc99 := __obf_512c4f2f38a27036.Map[__obf_4ee6afc37d3486aa.
			// Don't auto inline if there are shadowed fields.
			__obf_d7fb5470f7183af5]; __obf_77cfff95beb3cc99 {

			return false
		}
	}

	for _, __obf_4ee6afc37d3486aa := range __obf_d89d6fd0b094b410 {
		__obf_4ee6afc37d3486aa.__obf_8671df0a9132a879 = append(__obf_cf8245c3f9570af5.__obf_8671df0a9132a879, __obf_4ee6afc37d3486aa.__obf_8671df0a9132a879...)
		__obf_512c4f2f38a27036.
			Add(__obf_4ee6afc37d3486aa)
	}
	return true
}

type __obf_9b96ada439691d07 interface {
	IsZero() bool
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_2474e494c3416afb(__obf_17732590722140e7 reflect.Value) bool {
	__obf_f02bd2095cf36557 := __obf_17732590722140e7.Kind()

	for __obf_f02bd2095cf36557 == reflect.Interface {
		if __obf_17732590722140e7.IsNil() {
			return true
		}
		__obf_17732590722140e7 = __obf_17732590722140e7.Elem()
		__obf_f02bd2095cf36557 = __obf_17732590722140e7.Kind()
	}

	if __obf_04c2f3dd4a3cca11, __obf_77cfff95beb3cc99 := __obf_17732590722140e7.Interface().(__obf_9b96ada439691d07); __obf_77cfff95beb3cc99 {
		return __obf_c3d2a512f69a483f(__obf_f02bd2095cf36557) && __obf_17732590722140e7.IsNil() || __obf_04c2f3dd4a3cca11.IsZero()
	}

	switch __obf_f02bd2095cf36557 {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_17732590722140e7.Len() == 0
	case reflect.Struct:
		__obf_fa8592fc07f7c881 := __obf_4642342e308ea07b.Fields(__obf_17732590722140e7.Type(), __obf_7bae0b613da60dd3.__obf_603dff1ad8d49539)
		__obf_7553cd24ebfd7da5 := __obf_fa8592fc07f7c881.OmitEmpty(__obf_7bae0b613da60dd3, __obf_17732590722140e7)
		return len(__obf_7553cd24ebfd7da5) == 0
	case reflect.Bool:
		return !__obf_17732590722140e7.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_17732590722140e7.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_17732590722140e7.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_17732590722140e7.Float() == 0
	case reflect.Ptr:
		return __obf_17732590722140e7.IsNil()
	default:
		return false
	}
}

func __obf_3000a20be09182df(__obf_17732590722140e7 reflect.Value, __obf_8671df0a9132a879 []int) (_ reflect.Value, __obf_77cfff95beb3cc99 bool) {
	if len(__obf_8671df0a9132a879) == 1 {
		return __obf_17732590722140e7.Field(__obf_8671df0a9132a879[0]), true
	}

	for __obf_101eec84c8338296, __obf_a6d03241f2c8bded := range __obf_8671df0a9132a879 {
		if __obf_101eec84c8338296 > 0 {
			if __obf_17732590722140e7.Kind() == reflect.Ptr {
				if __obf_17732590722140e7.IsNil() {
					return __obf_17732590722140e7, false
				}
				__obf_17732590722140e7 = __obf_17732590722140e7.Elem()
			}
		}
		__obf_17732590722140e7 = __obf_17732590722140e7.Field(__obf_a6d03241f2c8bded)
	}

	return __obf_17732590722140e7, true
}

func __obf_3af645acd1a85372(__obf_17732590722140e7 reflect.Value, __obf_8671df0a9132a879 []int) reflect.Value {
	if len(__obf_8671df0a9132a879) == 1 {
		return __obf_17732590722140e7.Field(__obf_8671df0a9132a879[0])
	}

	for __obf_101eec84c8338296, __obf_a6d03241f2c8bded := range __obf_8671df0a9132a879 {
		if __obf_101eec84c8338296 > 0 {
			var __obf_77cfff95beb3cc99 bool
			__obf_17732590722140e7, __obf_77cfff95beb3cc99 = __obf_1e9708c0d7199214(__obf_17732590722140e7)
			if !__obf_77cfff95beb3cc99 {
				return __obf_17732590722140e7
			}
		}
		__obf_17732590722140e7 = __obf_17732590722140e7.Field(__obf_a6d03241f2c8bded)
	}

	return __obf_17732590722140e7
}

func __obf_1e9708c0d7199214(__obf_17732590722140e7 reflect.Value) (reflect.Value, bool) {
	if __obf_17732590722140e7.Kind() == reflect.Ptr {
		if __obf_17732590722140e7.IsNil() {
			if !__obf_17732590722140e7.CanSet() {
				return __obf_17732590722140e7, false
			}
			__obf_dd0314b1f4b16bdd := __obf_17732590722140e7.Type().Elem()
			if __obf_dd0314b1f4b16bdd.Kind() != reflect.Struct {
				return __obf_17732590722140e7, false
			}
			__obf_17732590722140e7.
				Set(__obf_cd14bf657d2f8fbe(__obf_dd0314b1f4b16bdd))
		}
		__obf_17732590722140e7 = __obf_17732590722140e7.Elem()
	}
	return __obf_17732590722140e7, true
}
