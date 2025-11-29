package __obf_3edaa49e853afa16

import (
	"encoding"
	"fmt"
	tagparser "github.com/ArtisanHiram/go-pkg/tagparser"
	"log"
	"reflect"
	"sync"
)

var __obf_6d051195c59d18d2 = reflect.TypeOf((*error)(nil)).Elem()

var (
	__obf_3dd0ef54e702a457 = reflect.TypeOf((*CustomEncoder)(nil)).Elem()
	__obf_3fcf1c5b6db37260 = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
)

var (
	__obf_6ff28f42337db34f = reflect.TypeOf((*Marshaler)(nil)).Elem()
	__obf_b2238d76404cb5bd = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var (
	__obf_a091a0cc441d3614 = reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	__obf_217405f1e132acef = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

var (
	__obf_4acc62fe1992007c = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	__obf_9d24d27869bcc179 = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

type (
	__obf_b2b0b98cf9f599e0 func(*Encoder, reflect.Value) error
	__obf_47bea90bc2ca8162 func(*Decoder, reflect.Value) error
)

var (
	__obf_126a303617e63b80 sync.Map
	__obf_8d7a14362d6f6fc4 sync.Map
)

// Register registers encoder and decoder functions for a value.
// This is low level API and in most cases you should prefer implementing
// CustomEncoder/CustomDecoder or Marshaler/Unmarshaler interfaces.
func Register(__obf_e6992bb5202a647c any, __obf_2fa4ef37f70ce1de __obf_b2b0b98cf9f599e0, __obf_20b87cf22b03338d __obf_47bea90bc2ca8162) {
	__obf_af250e9784b6a92c := reflect.TypeOf(__obf_e6992bb5202a647c)
	if __obf_2fa4ef37f70ce1de != nil {
		__obf_126a303617e63b80.
			Store(__obf_af250e9784b6a92c, __obf_2fa4ef37f70ce1de)
	}
	if __obf_20b87cf22b03338d != nil {
		__obf_8d7a14362d6f6fc4.
			Store(__obf_af250e9784b6a92c,

				//------------------------------------------------------------------------------
				__obf_20b87cf22b03338d)
	}
}

const __obf_76b9147e50457d35 = "msgpack"

var __obf_be0ec04e52f483d9 = __obf_bf8b9ce9be0ec131()

type __obf_f46a1541d31cf935 struct {
	__obf_c26f5adfb3152475 sync.Map
}

type __obf_f38cd1467f900b37 struct {
	__obf_af250e9784b6a92c reflect.Type
	__obf_160a7dd2c812a6a1 string
}

func __obf_bf8b9ce9be0ec131() *__obf_f46a1541d31cf935 {
	return new(__obf_f46a1541d31cf935)
}

func (__obf_c26f5adfb3152475 *__obf_f46a1541d31cf935) Fields(__obf_af250e9784b6a92c reflect.Type, __obf_160a7dd2c812a6a1 string) *__obf_fdeb38537c10bbcb {
	__obf_813ce87902031fb7 := __obf_f38cd1467f900b37{__obf_160a7dd2c812a6a1: __obf_160a7dd2c812a6a1, __obf_af250e9784b6a92c: __obf_af250e9784b6a92c}

	if __obf_85d270343a0dfe14, __obf_ccfb92cc26e4569f := __obf_c26f5adfb3152475.__obf_c26f5adfb3152475.Load(__obf_813ce87902031fb7); __obf_ccfb92cc26e4569f {
		return __obf_85d270343a0dfe14.(*__obf_fdeb38537c10bbcb)
	}
	__obf_343ca3a5e5fa96b2 := __obf_a4027e51a5c3a654(__obf_af250e9784b6a92c, __obf_160a7dd2c812a6a1)
	__obf_c26f5adfb3152475.__obf_c26f5adfb3152475.
		Store(__obf_813ce87902031fb7, __obf_343ca3a5e5fa96b2)

	return __obf_343ca3a5e5fa96b2
}

//------------------------------------------------------------------------------

type __obf_90e4cc2c82c33de0 struct {
	__obf_88d8471374f73dd1 __obf_b2b0b98cf9f599e0
	__obf_2cb4408eb815d713 __obf_47bea90bc2ca8162
	__obf_90a8fc4880887fcb string
	__obf_f688c6362d0017ac []int
	__obf_718fa6c1e694bddf bool
}

func (__obf_40a54345a49c3cd5 *__obf_90e4cc2c82c33de0) Omit(__obf_84d0d31a8288f191 *Encoder, __obf_14a1657cb1cdf60b reflect.Value) bool {
	__obf_85d270343a0dfe14, __obf_ccfb92cc26e4569f := __obf_6a12a270b7a6fb5f(__obf_14a1657cb1cdf60b, __obf_40a54345a49c3cd5.__obf_f688c6362d0017ac)
	if !__obf_ccfb92cc26e4569f {
		return true
	}
	__obf_eb0df5206cf3ccb1 := __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_344e8fa9adf9a114 != 0
	return (__obf_40a54345a49c3cd5.__obf_718fa6c1e694bddf || __obf_eb0df5206cf3ccb1) && __obf_84d0d31a8288f191.__obf_3ec23a13a2fbfd34(__obf_85d270343a0dfe14)
}

func (__obf_40a54345a49c3cd5 *__obf_90e4cc2c82c33de0) EncodeValue(__obf_84d0d31a8288f191 *Encoder, __obf_14a1657cb1cdf60b reflect.Value) error {
	__obf_85d270343a0dfe14, __obf_ccfb92cc26e4569f := __obf_6a12a270b7a6fb5f(__obf_14a1657cb1cdf60b, __obf_40a54345a49c3cd5.__obf_f688c6362d0017ac)
	if !__obf_ccfb92cc26e4569f {
		return __obf_84d0d31a8288f191.EncodeNil()
	}
	return __obf_40a54345a49c3cd5.__obf_88d8471374f73dd1(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14)
}

func (__obf_40a54345a49c3cd5 *__obf_90e4cc2c82c33de0) DecodeValue(__obf_015afbee33862a01 *Decoder, __obf_14a1657cb1cdf60b reflect.Value) error {
	__obf_85d270343a0dfe14 := __obf_7e91e0ea1f83caac(__obf_14a1657cb1cdf60b, __obf_40a54345a49c3cd5.

		//------------------------------------------------------------------------------
		__obf_f688c6362d0017ac)
	return __obf_40a54345a49c3cd5.__obf_2cb4408eb815d713(__obf_015afbee33862a01, __obf_85d270343a0dfe14)
}

type __obf_fdeb38537c10bbcb struct {
	Type                   reflect.Type
	Map                    map[string]*__obf_90e4cc2c82c33de0
	List                   []*__obf_90e4cc2c82c33de0
	AsArray                bool
	__obf_e8e6b5aaa4ecece0 bool
}

func __obf_1f3cbda8d66ed4db(__obf_af250e9784b6a92c reflect.Type) *__obf_fdeb38537c10bbcb {
	return &__obf_fdeb38537c10bbcb{
		Type: __obf_af250e9784b6a92c,
		Map:  make(map[string]*__obf_90e4cc2c82c33de0, __obf_af250e9784b6a92c.NumField()),
		List: make([]*__obf_90e4cc2c82c33de0, 0, __obf_af250e9784b6a92c.NumField()),
	}
}

func (__obf_343ca3a5e5fa96b2 *__obf_fdeb38537c10bbcb) Add(__obf_90e4cc2c82c33de0 *__obf_90e4cc2c82c33de0) {
	__obf_343ca3a5e5fa96b2.__obf_da6fde7a6536b2dc(__obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb)
	__obf_343ca3a5e5fa96b2.
		Map[__obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb] = __obf_90e4cc2c82c33de0
	__obf_343ca3a5e5fa96b2.
		List = append(__obf_343ca3a5e5fa96b2.List, __obf_90e4cc2c82c33de0)
	if __obf_90e4cc2c82c33de0.__obf_718fa6c1e694bddf {
		__obf_343ca3a5e5fa96b2.__obf_e8e6b5aaa4ecece0 = true
	}
}

func (__obf_343ca3a5e5fa96b2 *__obf_fdeb38537c10bbcb) __obf_da6fde7a6536b2dc(__obf_90a8fc4880887fcb string) {
	if _, __obf_ccfb92cc26e4569f := __obf_343ca3a5e5fa96b2.Map[__obf_90a8fc4880887fcb]; __obf_ccfb92cc26e4569f {
		log.Printf("msgpack: %s already has field=%s", __obf_343ca3a5e5fa96b2.Type, __obf_90a8fc4880887fcb)
	}
}

func (__obf_343ca3a5e5fa96b2 *__obf_fdeb38537c10bbcb) OmitEmpty(__obf_84d0d31a8288f191 *Encoder, __obf_14a1657cb1cdf60b reflect.Value) []*__obf_90e4cc2c82c33de0 {
	__obf_eb0df5206cf3ccb1 := __obf_84d0d31a8288f191.__obf_6aa0efd537415192&__obf_344e8fa9adf9a114 != 0
	if !__obf_343ca3a5e5fa96b2.__obf_e8e6b5aaa4ecece0 && !__obf_eb0df5206cf3ccb1 {
		return __obf_343ca3a5e5fa96b2.List
	}
	__obf_fdeb38537c10bbcb := make([]*__obf_90e4cc2c82c33de0, 0, len(__obf_343ca3a5e5fa96b2.List))

	for _, __obf_40a54345a49c3cd5 := range __obf_343ca3a5e5fa96b2.List {
		if !__obf_40a54345a49c3cd5.Omit(__obf_84d0d31a8288f191, __obf_14a1657cb1cdf60b) {
			__obf_fdeb38537c10bbcb = append(__obf_fdeb38537c10bbcb, __obf_40a54345a49c3cd5)
		}
	}

	return __obf_fdeb38537c10bbcb
}

func __obf_a4027e51a5c3a654(__obf_af250e9784b6a92c reflect.Type, __obf_afa480657f2940a5 string) *__obf_fdeb38537c10bbcb {
	__obf_343ca3a5e5fa96b2 := __obf_1f3cbda8d66ed4db(__obf_af250e9784b6a92c)

	var __obf_718fa6c1e694bddf bool
	for __obf_bd2da3a1d4616916 := 0; __obf_bd2da3a1d4616916 < __obf_af250e9784b6a92c.NumField(); __obf_bd2da3a1d4616916++ {
		__obf_40a54345a49c3cd5 := __obf_af250e9784b6a92c.Field(__obf_bd2da3a1d4616916)
		__obf_c5fb3761254d6cd7 := __obf_40a54345a49c3cd5.Tag.Get(__obf_76b9147e50457d35)
		if __obf_c5fb3761254d6cd7 == "" && __obf_afa480657f2940a5 != "" {
			__obf_c5fb3761254d6cd7 = __obf_40a54345a49c3cd5.Tag.Get(__obf_afa480657f2940a5)
		}
		__obf_160a7dd2c812a6a1 := tagparser.Parse(__obf_c5fb3761254d6cd7)
		if __obf_160a7dd2c812a6a1.Name == "-" {
			continue
		}

		if __obf_40a54345a49c3cd5.Name == "_msgpack" {
			__obf_343ca3a5e5fa96b2.
				AsArray = __obf_160a7dd2c812a6a1.HasOption("as_array") || __obf_160a7dd2c812a6a1.HasOption("asArray")
			if __obf_160a7dd2c812a6a1.HasOption("omitempty") {
				__obf_718fa6c1e694bddf = true
			}
		}

		if __obf_40a54345a49c3cd5.PkgPath != "" && !__obf_40a54345a49c3cd5.Anonymous {
			continue
		}
		__obf_90e4cc2c82c33de0 := &__obf_90e4cc2c82c33de0{__obf_90a8fc4880887fcb: __obf_160a7dd2c812a6a1.Name, __obf_f688c6362d0017ac: __obf_40a54345a49c3cd5.Index, __obf_718fa6c1e694bddf: __obf_718fa6c1e694bddf || __obf_160a7dd2c812a6a1.HasOption("omitempty")}

		if __obf_160a7dd2c812a6a1.HasOption("intern") {
			switch __obf_40a54345a49c3cd5.Type.Kind() {
			case reflect.Interface:
				__obf_90e4cc2c82c33de0.__obf_88d8471374f73dd1 = __obf_7ad73e7d72971622
				__obf_90e4cc2c82c33de0.__obf_2cb4408eb815d713 = __obf_867c9a7f5c7de360
			case reflect.String:
				__obf_90e4cc2c82c33de0.__obf_88d8471374f73dd1 = __obf_25ae4057eed7711b
				__obf_90e4cc2c82c33de0.__obf_2cb4408eb815d713 = __obf_0156ce491a786988
			default:
				__obf_3aa78ad28f50ed46 := fmt.Errorf("msgpack: intern strings are not supported on %s", __obf_40a54345a49c3cd5.Type)
				panic(__obf_3aa78ad28f50ed46)
			}
		} else {
			__obf_90e4cc2c82c33de0.__obf_88d8471374f73dd1 = __obf_1e49c68fa4194051(__obf_40a54345a49c3cd5.Type)
			__obf_90e4cc2c82c33de0.__obf_2cb4408eb815d713 = __obf_6253a63b14aba59e(__obf_40a54345a49c3cd5.Type)
		}

		if __obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb == "" {
			__obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb = __obf_40a54345a49c3cd5.Name
		}

		if __obf_40a54345a49c3cd5.Anonymous && !__obf_160a7dd2c812a6a1.HasOption("noinline") {
			__obf_ec76d0da66306448 := __obf_160a7dd2c812a6a1.HasOption("inline")
			if __obf_ec76d0da66306448 {
				__obf_4840c79d70930133(__obf_343ca3a5e5fa96b2, __obf_40a54345a49c3cd5.Type, __obf_90e4cc2c82c33de0, __obf_afa480657f2940a5)
			} else {
				__obf_ec76d0da66306448 = __obf_1eecb67382c1d018(__obf_343ca3a5e5fa96b2, __obf_40a54345a49c3cd5.Type, __obf_90e4cc2c82c33de0, __obf_afa480657f2940a5)
			}

			if __obf_ec76d0da66306448 {
				if _, __obf_ccfb92cc26e4569f := __obf_343ca3a5e5fa96b2.Map[__obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb]; __obf_ccfb92cc26e4569f {
					log.Printf("msgpack: %s already has field=%s", __obf_343ca3a5e5fa96b2.Type, __obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb)
				}
				__obf_343ca3a5e5fa96b2.
					Map[__obf_90e4cc2c82c33de0.__obf_90a8fc4880887fcb] = __obf_90e4cc2c82c33de0
				continue
			}
		}
		__obf_343ca3a5e5fa96b2.
			Add(__obf_90e4cc2c82c33de0)

		if __obf_76a51e87fec3b03b, __obf_ccfb92cc26e4569f := __obf_160a7dd2c812a6a1.Options["alias"]; __obf_ccfb92cc26e4569f {
			__obf_343ca3a5e5fa96b2.__obf_da6fde7a6536b2dc(__obf_76a51e87fec3b03b)
			__obf_343ca3a5e5fa96b2.
				Map[__obf_76a51e87fec3b03b] = __obf_90e4cc2c82c33de0
		}
	}
	return __obf_343ca3a5e5fa96b2
}

var (
	__obf_92e9325f6a76cfca uintptr
	__obf_ff3c6a2acf133796 uintptr
)

//nolint:gochecknoinits
func init() {
	__obf_92e9325f6a76cfca = reflect.ValueOf(__obf_bb1447402cee3139).Pointer()
	__obf_ff3c6a2acf133796 = reflect.ValueOf(__obf_4f268d405fba2c41).Pointer()
}

func __obf_4840c79d70930133(__obf_343ca3a5e5fa96b2 *__obf_fdeb38537c10bbcb, __obf_af250e9784b6a92c reflect.Type, __obf_40a54345a49c3cd5 *__obf_90e4cc2c82c33de0, __obf_160a7dd2c812a6a1 string) {
	__obf_c17de34cbd034a5a := __obf_a4027e51a5c3a654(__obf_af250e9784b6a92c, __obf_160a7dd2c812a6a1).List
	for _, __obf_90e4cc2c82c33de0 := range __obf_c17de34cbd034a5a {
		if _, __obf_ccfb92cc26e4569f := __obf_343ca3a5e5fa96b2.Map[__obf_90e4cc2c82c33de0.
			// Don't inline shadowed fields.
			__obf_90a8fc4880887fcb]; __obf_ccfb92cc26e4569f {

			continue
		}
		__obf_90e4cc2c82c33de0.__obf_f688c6362d0017ac = append(__obf_40a54345a49c3cd5.__obf_f688c6362d0017ac, __obf_90e4cc2c82c33de0.__obf_f688c6362d0017ac...)
		__obf_343ca3a5e5fa96b2.
			Add(__obf_90e4cc2c82c33de0)
	}
}

func __obf_1eecb67382c1d018(__obf_343ca3a5e5fa96b2 *__obf_fdeb38537c10bbcb, __obf_af250e9784b6a92c reflect.Type, __obf_40a54345a49c3cd5 *__obf_90e4cc2c82c33de0, __obf_160a7dd2c812a6a1 string) bool {
	var __obf_88d8471374f73dd1 __obf_b2b0b98cf9f599e0
	var __obf_2cb4408eb815d713 __obf_47bea90bc2ca8162

	if __obf_af250e9784b6a92c.Kind() == reflect.Struct {
		__obf_88d8471374f73dd1 = __obf_40a54345a49c3cd5.__obf_88d8471374f73dd1
		__obf_2cb4408eb815d713 = __obf_40a54345a49c3cd5.__obf_2cb4408eb815d713
	} else {
		for __obf_af250e9784b6a92c.Kind() == reflect.Ptr {
			__obf_af250e9784b6a92c = __obf_af250e9784b6a92c.Elem()
			__obf_88d8471374f73dd1 = __obf_1e49c68fa4194051(__obf_af250e9784b6a92c)
			__obf_2cb4408eb815d713 = __obf_6253a63b14aba59e(__obf_af250e9784b6a92c)
		}
		if __obf_af250e9784b6a92c.Kind() != reflect.Struct {
			return false
		}
	}

	if reflect.ValueOf(__obf_88d8471374f73dd1).Pointer() != __obf_92e9325f6a76cfca {
		return false
	}
	if reflect.ValueOf(__obf_2cb4408eb815d713).Pointer() != __obf_ff3c6a2acf133796 {
		return false
	}
	__obf_c17de34cbd034a5a := __obf_a4027e51a5c3a654(__obf_af250e9784b6a92c, __obf_160a7dd2c812a6a1).List
	for _, __obf_90e4cc2c82c33de0 := range __obf_c17de34cbd034a5a {
		if _, __obf_ccfb92cc26e4569f := __obf_343ca3a5e5fa96b2.Map[__obf_90e4cc2c82c33de0.
			// Don't auto inline if there are shadowed fields.
			__obf_90a8fc4880887fcb]; __obf_ccfb92cc26e4569f {

			return false
		}
	}

	for _, __obf_90e4cc2c82c33de0 := range __obf_c17de34cbd034a5a {
		__obf_90e4cc2c82c33de0.__obf_f688c6362d0017ac = append(__obf_40a54345a49c3cd5.__obf_f688c6362d0017ac, __obf_90e4cc2c82c33de0.__obf_f688c6362d0017ac...)
		__obf_343ca3a5e5fa96b2.
			Add(__obf_90e4cc2c82c33de0)
	}
	return true
}

type __obf_a87f1759e98e562e interface {
	IsZero() bool
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_3ec23a13a2fbfd34(__obf_85d270343a0dfe14 reflect.Value) bool {
	__obf_5819e74754f40cb2 := __obf_85d270343a0dfe14.Kind()

	for __obf_5819e74754f40cb2 == reflect.Interface {
		if __obf_85d270343a0dfe14.IsNil() {
			return true
		}
		__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Elem()
		__obf_5819e74754f40cb2 = __obf_85d270343a0dfe14.Kind()
	}

	if __obf_89316e9145798849, __obf_ccfb92cc26e4569f := __obf_85d270343a0dfe14.Interface().(__obf_a87f1759e98e562e); __obf_ccfb92cc26e4569f {
		return __obf_6238d8f3c0669e3b(__obf_5819e74754f40cb2) && __obf_85d270343a0dfe14.IsNil() || __obf_89316e9145798849.IsZero()
	}

	switch __obf_5819e74754f40cb2 {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_85d270343a0dfe14.Len() == 0
	case reflect.Struct:
		__obf_94f9ea99a9fa147b := __obf_be0ec04e52f483d9.Fields(__obf_85d270343a0dfe14.Type(), __obf_84d0d31a8288f191.__obf_ec5d9a4ce8dd41be)
		__obf_fdeb38537c10bbcb := __obf_94f9ea99a9fa147b.OmitEmpty(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14)
		return len(__obf_fdeb38537c10bbcb) == 0
	case reflect.Bool:
		return !__obf_85d270343a0dfe14.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_85d270343a0dfe14.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_85d270343a0dfe14.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_85d270343a0dfe14.Float() == 0
	case reflect.Ptr:
		return __obf_85d270343a0dfe14.IsNil()
	default:
		return false
	}
}

func __obf_6a12a270b7a6fb5f(__obf_85d270343a0dfe14 reflect.Value, __obf_f688c6362d0017ac []int) (_ reflect.Value, __obf_ccfb92cc26e4569f bool) {
	if len(__obf_f688c6362d0017ac) == 1 {
		return __obf_85d270343a0dfe14.Field(__obf_f688c6362d0017ac[0]), true
	}

	for __obf_bd2da3a1d4616916, __obf_eff265c728cf9640 := range __obf_f688c6362d0017ac {
		if __obf_bd2da3a1d4616916 > 0 {
			if __obf_85d270343a0dfe14.Kind() == reflect.Ptr {
				if __obf_85d270343a0dfe14.IsNil() {
					return __obf_85d270343a0dfe14, false
				}
				__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Elem()
			}
		}
		__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Field(__obf_eff265c728cf9640)
	}

	return __obf_85d270343a0dfe14, true
}

func __obf_7e91e0ea1f83caac(__obf_85d270343a0dfe14 reflect.Value, __obf_f688c6362d0017ac []int) reflect.Value {
	if len(__obf_f688c6362d0017ac) == 1 {
		return __obf_85d270343a0dfe14.Field(__obf_f688c6362d0017ac[0])
	}

	for __obf_bd2da3a1d4616916, __obf_eff265c728cf9640 := range __obf_f688c6362d0017ac {
		if __obf_bd2da3a1d4616916 > 0 {
			var __obf_ccfb92cc26e4569f bool
			__obf_85d270343a0dfe14, __obf_ccfb92cc26e4569f = __obf_a2e2f31cff9a8c39(__obf_85d270343a0dfe14)
			if !__obf_ccfb92cc26e4569f {
				return __obf_85d270343a0dfe14
			}
		}
		__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Field(__obf_eff265c728cf9640)
	}

	return __obf_85d270343a0dfe14
}

func __obf_a2e2f31cff9a8c39(__obf_85d270343a0dfe14 reflect.Value) (reflect.Value, bool) {
	if __obf_85d270343a0dfe14.Kind() == reflect.Ptr {
		if __obf_85d270343a0dfe14.IsNil() {
			if !__obf_85d270343a0dfe14.CanSet() {
				return __obf_85d270343a0dfe14, false
			}
			__obf_bce1b192b5df8d1e := __obf_85d270343a0dfe14.Type().Elem()
			if __obf_bce1b192b5df8d1e.Kind() != reflect.Struct {
				return __obf_85d270343a0dfe14, false
			}
			__obf_85d270343a0dfe14.
				Set(__obf_c3a387a7995b1678(__obf_bce1b192b5df8d1e))
		}
		__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Elem()
	}
	return __obf_85d270343a0dfe14, true
}
