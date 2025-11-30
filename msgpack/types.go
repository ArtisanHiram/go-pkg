package __obf_3e0c303610a19bd4

import (
	"encoding"
	"fmt"
	tagparser "github.com/ArtisanHiram/go-pkg/tagparser"
	"log"
	"reflect"
	"sync"
)

var __obf_319e36d3ccef195a = reflect.TypeOf((*error)(nil)).Elem()

var (
	__obf_c08f5968e7619fb2 = reflect.TypeOf((*CustomEncoder)(nil)).Elem()
	__obf_1067551650cbff5c = reflect.TypeOf((*CustomDecoder)(nil)).Elem()
)

var (
	__obf_9e290b0c7c8448c5 = reflect.TypeOf((*Marshaler)(nil)).Elem()
	__obf_65915279442e2d73 = reflect.TypeOf((*Unmarshaler)(nil)).Elem()
)

var (
	__obf_c4304d6f5eef9ca6 = reflect.TypeOf((*encoding.BinaryMarshaler)(nil)).Elem()
	__obf_03e2d73a2d650e68 = reflect.TypeOf((*encoding.BinaryUnmarshaler)(nil)).Elem()
)

var (
	__obf_09aab39b0e9cc30f = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
	__obf_e682a601b05d7802 = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()
)

type (
	__obf_9cac677a20c1be1c func(*Encoder, reflect.Value) error
	__obf_299b7590962e0960 func(*Decoder, reflect.Value) error
)

var (
	__obf_b932a15e2d3972de sync.Map
	__obf_a46bdc105884db8b sync.Map
)

// Register registers encoder and decoder functions for a value.
// This is low level API and in most cases you should prefer implementing
// CustomEncoder/CustomDecoder or Marshaler/Unmarshaler interfaces.
func Register(__obf_752fc3666f4041f7 any, __obf_9721a7bc6b8fa008 __obf_9cac677a20c1be1c, __obf_0ed630deb24db696 __obf_299b7590962e0960) {
	__obf_616f98efc80197c6 := reflect.TypeOf(__obf_752fc3666f4041f7)
	if __obf_9721a7bc6b8fa008 != nil {
		__obf_b932a15e2d3972de.
			Store(__obf_616f98efc80197c6, __obf_9721a7bc6b8fa008)
	}
	if __obf_0ed630deb24db696 != nil {
		__obf_a46bdc105884db8b.
			Store(__obf_616f98efc80197c6,

				//------------------------------------------------------------------------------
				__obf_0ed630deb24db696)
	}
}

const __obf_572a68f3fcc6bd46 = "msgpack"

var __obf_3a8fae1624b5849e = __obf_d6ba4d72ebf51213()

type __obf_8884698946b71a57 struct {
	__obf_b416b6a4555be5c8 sync.Map
}

type __obf_d44484c33f5f827e struct {
	__obf_616f98efc80197c6 reflect.Type
	__obf_42ff6f144202733b string
}

func __obf_d6ba4d72ebf51213() *__obf_8884698946b71a57 {
	return new(__obf_8884698946b71a57)
}

func (__obf_b416b6a4555be5c8 *__obf_8884698946b71a57) Fields(__obf_616f98efc80197c6 reflect.Type, __obf_42ff6f144202733b string) *__obf_ce19c38e443cefa8 {
	__obf_ef8207c019b4cdb3 := __obf_d44484c33f5f827e{__obf_42ff6f144202733b: __obf_42ff6f144202733b, __obf_616f98efc80197c6: __obf_616f98efc80197c6}

	if __obf_63bbcee86d44fdde, __obf_ce8bef141eff8aab := __obf_b416b6a4555be5c8.__obf_b416b6a4555be5c8.Load(__obf_ef8207c019b4cdb3); __obf_ce8bef141eff8aab {
		return __obf_63bbcee86d44fdde.(*__obf_ce19c38e443cefa8)
	}
	__obf_e2aea562f6e32da0 := __obf_7fcccb205678bae5(__obf_616f98efc80197c6, __obf_42ff6f144202733b)
	__obf_b416b6a4555be5c8.__obf_b416b6a4555be5c8.
		Store(__obf_ef8207c019b4cdb3, __obf_e2aea562f6e32da0)

	return __obf_e2aea562f6e32da0
}

//------------------------------------------------------------------------------

type __obf_a5cb7d1c497d43ce struct {
	__obf_02e33ee3dfeba894 __obf_9cac677a20c1be1c
	__obf_91fd0d587e6af119 __obf_299b7590962e0960
	__obf_754359097736e0d5 string
	__obf_c94bb2895e92cb8b []int
	__obf_e3cfd1fc4e64d78e bool
}

func (__obf_4add40b0f5dc6b95 *__obf_a5cb7d1c497d43ce) Omit(__obf_77240dc7776b60b4 *Encoder, __obf_b8f9ff0ad69ee384 reflect.Value) bool {
	__obf_63bbcee86d44fdde, __obf_ce8bef141eff8aab := __obf_07a14ad4582e9fd8(__obf_b8f9ff0ad69ee384, __obf_4add40b0f5dc6b95.__obf_c94bb2895e92cb8b)
	if !__obf_ce8bef141eff8aab {
		return true
	}
	__obf_5d83c968250567c8 := __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_4536375463a63368 != 0
	return (__obf_4add40b0f5dc6b95.__obf_e3cfd1fc4e64d78e || __obf_5d83c968250567c8) && __obf_77240dc7776b60b4.__obf_2130807785985f00(__obf_63bbcee86d44fdde)
}

func (__obf_4add40b0f5dc6b95 *__obf_a5cb7d1c497d43ce) EncodeValue(__obf_77240dc7776b60b4 *Encoder, __obf_b8f9ff0ad69ee384 reflect.Value) error {
	__obf_63bbcee86d44fdde, __obf_ce8bef141eff8aab := __obf_07a14ad4582e9fd8(__obf_b8f9ff0ad69ee384, __obf_4add40b0f5dc6b95.__obf_c94bb2895e92cb8b)
	if !__obf_ce8bef141eff8aab {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	return __obf_4add40b0f5dc6b95.__obf_02e33ee3dfeba894(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde)
}

func (__obf_4add40b0f5dc6b95 *__obf_a5cb7d1c497d43ce) DecodeValue(__obf_dc35117108ba8439 *Decoder, __obf_b8f9ff0ad69ee384 reflect.Value) error {
	__obf_63bbcee86d44fdde := __obf_8a48edb29af845be(__obf_b8f9ff0ad69ee384, __obf_4add40b0f5dc6b95.

		//------------------------------------------------------------------------------
		__obf_c94bb2895e92cb8b)
	return __obf_4add40b0f5dc6b95.__obf_91fd0d587e6af119(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde)
}

type __obf_ce19c38e443cefa8 struct {
	Type                   reflect.Type
	Map                    map[string]*__obf_a5cb7d1c497d43ce
	List                   []*__obf_a5cb7d1c497d43ce
	AsArray                bool
	__obf_a3442d42a4d36166 bool
}

func __obf_2a0aaac9bc2b86d0(__obf_616f98efc80197c6 reflect.Type) *__obf_ce19c38e443cefa8 {
	return &__obf_ce19c38e443cefa8{
		Type: __obf_616f98efc80197c6,
		Map:  make(map[string]*__obf_a5cb7d1c497d43ce, __obf_616f98efc80197c6.NumField()),
		List: make([]*__obf_a5cb7d1c497d43ce, 0, __obf_616f98efc80197c6.NumField()),
	}
}

func (__obf_e2aea562f6e32da0 *__obf_ce19c38e443cefa8) Add(__obf_a5cb7d1c497d43ce *__obf_a5cb7d1c497d43ce) {
	__obf_e2aea562f6e32da0.__obf_97cab06946c1efc2(__obf_a5cb7d1c497d43ce.__obf_754359097736e0d5)
	__obf_e2aea562f6e32da0.
		Map[__obf_a5cb7d1c497d43ce.__obf_754359097736e0d5] = __obf_a5cb7d1c497d43ce
	__obf_e2aea562f6e32da0.
		List = append(__obf_e2aea562f6e32da0.List, __obf_a5cb7d1c497d43ce)
	if __obf_a5cb7d1c497d43ce.__obf_e3cfd1fc4e64d78e {
		__obf_e2aea562f6e32da0.__obf_a3442d42a4d36166 = true
	}
}

func (__obf_e2aea562f6e32da0 *__obf_ce19c38e443cefa8) __obf_97cab06946c1efc2(__obf_754359097736e0d5 string) {
	if _, __obf_ce8bef141eff8aab := __obf_e2aea562f6e32da0.Map[__obf_754359097736e0d5]; __obf_ce8bef141eff8aab {
		log.Printf("msgpack: %s already has field=%s", __obf_e2aea562f6e32da0.Type, __obf_754359097736e0d5)
	}
}

func (__obf_e2aea562f6e32da0 *__obf_ce19c38e443cefa8) OmitEmpty(__obf_77240dc7776b60b4 *Encoder, __obf_b8f9ff0ad69ee384 reflect.Value) []*__obf_a5cb7d1c497d43ce {
	__obf_5d83c968250567c8 := __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_4536375463a63368 != 0
	if !__obf_e2aea562f6e32da0.__obf_a3442d42a4d36166 && !__obf_5d83c968250567c8 {
		return __obf_e2aea562f6e32da0.List
	}
	__obf_ce19c38e443cefa8 := make([]*__obf_a5cb7d1c497d43ce, 0, len(__obf_e2aea562f6e32da0.List))

	for _, __obf_4add40b0f5dc6b95 := range __obf_e2aea562f6e32da0.List {
		if !__obf_4add40b0f5dc6b95.Omit(__obf_77240dc7776b60b4, __obf_b8f9ff0ad69ee384) {
			__obf_ce19c38e443cefa8 = append(__obf_ce19c38e443cefa8, __obf_4add40b0f5dc6b95)
		}
	}

	return __obf_ce19c38e443cefa8
}

func __obf_7fcccb205678bae5(__obf_616f98efc80197c6 reflect.Type, __obf_b9448f1adffe6c7e string) *__obf_ce19c38e443cefa8 {
	__obf_e2aea562f6e32da0 := __obf_2a0aaac9bc2b86d0(__obf_616f98efc80197c6)

	var __obf_e3cfd1fc4e64d78e bool
	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_616f98efc80197c6.NumField(); __obf_39714879601f9b69++ {
		__obf_4add40b0f5dc6b95 := __obf_616f98efc80197c6.Field(__obf_39714879601f9b69)
		__obf_b6799d24ec13a8d1 := __obf_4add40b0f5dc6b95.Tag.Get(__obf_572a68f3fcc6bd46)
		if __obf_b6799d24ec13a8d1 == "" && __obf_b9448f1adffe6c7e != "" {
			__obf_b6799d24ec13a8d1 = __obf_4add40b0f5dc6b95.Tag.Get(__obf_b9448f1adffe6c7e)
		}
		__obf_42ff6f144202733b := tagparser.Parse(__obf_b6799d24ec13a8d1)
		if __obf_42ff6f144202733b.Name == "-" {
			continue
		}

		if __obf_4add40b0f5dc6b95.Name == "_msgpack" {
			__obf_e2aea562f6e32da0.
				AsArray = __obf_42ff6f144202733b.HasOption("as_array") || __obf_42ff6f144202733b.HasOption("asArray")
			if __obf_42ff6f144202733b.HasOption("omitempty") {
				__obf_e3cfd1fc4e64d78e = true
			}
		}

		if __obf_4add40b0f5dc6b95.PkgPath != "" && !__obf_4add40b0f5dc6b95.Anonymous {
			continue
		}
		__obf_a5cb7d1c497d43ce := &__obf_a5cb7d1c497d43ce{__obf_754359097736e0d5: __obf_42ff6f144202733b.Name, __obf_c94bb2895e92cb8b: __obf_4add40b0f5dc6b95.Index, __obf_e3cfd1fc4e64d78e: __obf_e3cfd1fc4e64d78e || __obf_42ff6f144202733b.HasOption("omitempty")}

		if __obf_42ff6f144202733b.HasOption("intern") {
			switch __obf_4add40b0f5dc6b95.Type.Kind() {
			case reflect.Interface:
				__obf_a5cb7d1c497d43ce.__obf_02e33ee3dfeba894 = __obf_864478adcb603602
				__obf_a5cb7d1c497d43ce.__obf_91fd0d587e6af119 = __obf_68c203f3636fd603
			case reflect.String:
				__obf_a5cb7d1c497d43ce.__obf_02e33ee3dfeba894 = __obf_51341bc888154b6f
				__obf_a5cb7d1c497d43ce.__obf_91fd0d587e6af119 = __obf_113b700fc6a29ee1
			default:
				__obf_8882f6cf6e378ded := fmt.Errorf("msgpack: intern strings are not supported on %s", __obf_4add40b0f5dc6b95.Type)
				panic(__obf_8882f6cf6e378ded)
			}
		} else {
			__obf_a5cb7d1c497d43ce.__obf_02e33ee3dfeba894 = __obf_cc16be3b4933c617(__obf_4add40b0f5dc6b95.Type)
			__obf_a5cb7d1c497d43ce.__obf_91fd0d587e6af119 = __obf_9b0fb8e0beb3ab15(__obf_4add40b0f5dc6b95.Type)
		}

		if __obf_a5cb7d1c497d43ce.__obf_754359097736e0d5 == "" {
			__obf_a5cb7d1c497d43ce.__obf_754359097736e0d5 = __obf_4add40b0f5dc6b95.Name
		}

		if __obf_4add40b0f5dc6b95.Anonymous && !__obf_42ff6f144202733b.HasOption("noinline") {
			__obf_4d7261702ce5d2bf := __obf_42ff6f144202733b.HasOption("inline")
			if __obf_4d7261702ce5d2bf {
				__obf_776e11e1b5ddea73(__obf_e2aea562f6e32da0, __obf_4add40b0f5dc6b95.Type, __obf_a5cb7d1c497d43ce, __obf_b9448f1adffe6c7e)
			} else {
				__obf_4d7261702ce5d2bf = __obf_a60249c29c7c4dcf(__obf_e2aea562f6e32da0, __obf_4add40b0f5dc6b95.Type, __obf_a5cb7d1c497d43ce, __obf_b9448f1adffe6c7e)
			}

			if __obf_4d7261702ce5d2bf {
				if _, __obf_ce8bef141eff8aab := __obf_e2aea562f6e32da0.Map[__obf_a5cb7d1c497d43ce.__obf_754359097736e0d5]; __obf_ce8bef141eff8aab {
					log.Printf("msgpack: %s already has field=%s", __obf_e2aea562f6e32da0.Type, __obf_a5cb7d1c497d43ce.__obf_754359097736e0d5)
				}
				__obf_e2aea562f6e32da0.
					Map[__obf_a5cb7d1c497d43ce.__obf_754359097736e0d5] = __obf_a5cb7d1c497d43ce
				continue
			}
		}
		__obf_e2aea562f6e32da0.
			Add(__obf_a5cb7d1c497d43ce)

		if __obf_b63d180fe50ab281, __obf_ce8bef141eff8aab := __obf_42ff6f144202733b.Options["alias"]; __obf_ce8bef141eff8aab {
			__obf_e2aea562f6e32da0.__obf_97cab06946c1efc2(__obf_b63d180fe50ab281)
			__obf_e2aea562f6e32da0.
				Map[__obf_b63d180fe50ab281] = __obf_a5cb7d1c497d43ce
		}
	}
	return __obf_e2aea562f6e32da0
}

var (
	__obf_0c34bf045e7984d4 uintptr
	__obf_ec11f3e0c7fee7b7 uintptr
)

//nolint:gochecknoinits
func init() {
	__obf_0c34bf045e7984d4 = reflect.ValueOf(__obf_cbc3a9db722bf76c).Pointer()
	__obf_ec11f3e0c7fee7b7 = reflect.ValueOf(__obf_58fb4ae094054808).Pointer()
}

func __obf_776e11e1b5ddea73(__obf_e2aea562f6e32da0 *__obf_ce19c38e443cefa8, __obf_616f98efc80197c6 reflect.Type, __obf_4add40b0f5dc6b95 *__obf_a5cb7d1c497d43ce, __obf_42ff6f144202733b string) {
	__obf_67c3125e962868e5 := __obf_7fcccb205678bae5(__obf_616f98efc80197c6, __obf_42ff6f144202733b).List
	for _, __obf_a5cb7d1c497d43ce := range __obf_67c3125e962868e5 {
		if _, __obf_ce8bef141eff8aab := __obf_e2aea562f6e32da0.Map[__obf_a5cb7d1c497d43ce.
			// Don't inline shadowed fields.
			__obf_754359097736e0d5]; __obf_ce8bef141eff8aab {

			continue
		}
		__obf_a5cb7d1c497d43ce.__obf_c94bb2895e92cb8b = append(__obf_4add40b0f5dc6b95.__obf_c94bb2895e92cb8b, __obf_a5cb7d1c497d43ce.__obf_c94bb2895e92cb8b...)
		__obf_e2aea562f6e32da0.
			Add(__obf_a5cb7d1c497d43ce)
	}
}

func __obf_a60249c29c7c4dcf(__obf_e2aea562f6e32da0 *__obf_ce19c38e443cefa8, __obf_616f98efc80197c6 reflect.Type, __obf_4add40b0f5dc6b95 *__obf_a5cb7d1c497d43ce, __obf_42ff6f144202733b string) bool {
	var __obf_02e33ee3dfeba894 __obf_9cac677a20c1be1c
	var __obf_91fd0d587e6af119 __obf_299b7590962e0960

	if __obf_616f98efc80197c6.Kind() == reflect.Struct {
		__obf_02e33ee3dfeba894 = __obf_4add40b0f5dc6b95.__obf_02e33ee3dfeba894
		__obf_91fd0d587e6af119 = __obf_4add40b0f5dc6b95.__obf_91fd0d587e6af119
	} else {
		for __obf_616f98efc80197c6.Kind() == reflect.Ptr {
			__obf_616f98efc80197c6 = __obf_616f98efc80197c6.Elem()
			__obf_02e33ee3dfeba894 = __obf_cc16be3b4933c617(__obf_616f98efc80197c6)
			__obf_91fd0d587e6af119 = __obf_9b0fb8e0beb3ab15(__obf_616f98efc80197c6)
		}
		if __obf_616f98efc80197c6.Kind() != reflect.Struct {
			return false
		}
	}

	if reflect.ValueOf(__obf_02e33ee3dfeba894).Pointer() != __obf_0c34bf045e7984d4 {
		return false
	}
	if reflect.ValueOf(__obf_91fd0d587e6af119).Pointer() != __obf_ec11f3e0c7fee7b7 {
		return false
	}
	__obf_67c3125e962868e5 := __obf_7fcccb205678bae5(__obf_616f98efc80197c6, __obf_42ff6f144202733b).List
	for _, __obf_a5cb7d1c497d43ce := range __obf_67c3125e962868e5 {
		if _, __obf_ce8bef141eff8aab := __obf_e2aea562f6e32da0.Map[__obf_a5cb7d1c497d43ce.
			// Don't auto inline if there are shadowed fields.
			__obf_754359097736e0d5]; __obf_ce8bef141eff8aab {

			return false
		}
	}

	for _, __obf_a5cb7d1c497d43ce := range __obf_67c3125e962868e5 {
		__obf_a5cb7d1c497d43ce.__obf_c94bb2895e92cb8b = append(__obf_4add40b0f5dc6b95.__obf_c94bb2895e92cb8b, __obf_a5cb7d1c497d43ce.__obf_c94bb2895e92cb8b...)
		__obf_e2aea562f6e32da0.
			Add(__obf_a5cb7d1c497d43ce)
	}
	return true
}

type __obf_3e4700c471092a5c interface {
	IsZero() bool
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_2130807785985f00(__obf_63bbcee86d44fdde reflect.Value) bool {
	__obf_018b20441ecc0301 := __obf_63bbcee86d44fdde.Kind()

	for __obf_018b20441ecc0301 == reflect.Interface {
		if __obf_63bbcee86d44fdde.IsNil() {
			return true
		}
		__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Elem()
		__obf_018b20441ecc0301 = __obf_63bbcee86d44fdde.Kind()
	}

	if __obf_ef72885e579601c6, __obf_ce8bef141eff8aab := __obf_63bbcee86d44fdde.Interface().(__obf_3e4700c471092a5c); __obf_ce8bef141eff8aab {
		return __obf_ebd8855a2f3055ad(__obf_018b20441ecc0301) && __obf_63bbcee86d44fdde.IsNil() || __obf_ef72885e579601c6.IsZero()
	}

	switch __obf_018b20441ecc0301 {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_63bbcee86d44fdde.Len() == 0
	case reflect.Struct:
		__obf_ab0626d3c29fac63 := __obf_3a8fae1624b5849e.Fields(__obf_63bbcee86d44fdde.Type(), __obf_77240dc7776b60b4.__obf_6d0a0d8a79287f26)
		__obf_ce19c38e443cefa8 := __obf_ab0626d3c29fac63.OmitEmpty(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde)
		return len(__obf_ce19c38e443cefa8) == 0
	case reflect.Bool:
		return !__obf_63bbcee86d44fdde.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_63bbcee86d44fdde.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_63bbcee86d44fdde.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_63bbcee86d44fdde.Float() == 0
	case reflect.Ptr:
		return __obf_63bbcee86d44fdde.IsNil()
	default:
		return false
	}
}

func __obf_07a14ad4582e9fd8(__obf_63bbcee86d44fdde reflect.Value, __obf_c94bb2895e92cb8b []int) (_ reflect.Value, __obf_ce8bef141eff8aab bool) {
	if len(__obf_c94bb2895e92cb8b) == 1 {
		return __obf_63bbcee86d44fdde.Field(__obf_c94bb2895e92cb8b[0]), true
	}

	for __obf_39714879601f9b69, __obf_4040153f7700644e := range __obf_c94bb2895e92cb8b {
		if __obf_39714879601f9b69 > 0 {
			if __obf_63bbcee86d44fdde.Kind() == reflect.Ptr {
				if __obf_63bbcee86d44fdde.IsNil() {
					return __obf_63bbcee86d44fdde, false
				}
				__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Elem()
			}
		}
		__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Field(__obf_4040153f7700644e)
	}

	return __obf_63bbcee86d44fdde, true
}

func __obf_8a48edb29af845be(__obf_63bbcee86d44fdde reflect.Value, __obf_c94bb2895e92cb8b []int) reflect.Value {
	if len(__obf_c94bb2895e92cb8b) == 1 {
		return __obf_63bbcee86d44fdde.Field(__obf_c94bb2895e92cb8b[0])
	}

	for __obf_39714879601f9b69, __obf_4040153f7700644e := range __obf_c94bb2895e92cb8b {
		if __obf_39714879601f9b69 > 0 {
			var __obf_ce8bef141eff8aab bool
			__obf_63bbcee86d44fdde, __obf_ce8bef141eff8aab = __obf_4b8140a063da43be(__obf_63bbcee86d44fdde)
			if !__obf_ce8bef141eff8aab {
				return __obf_63bbcee86d44fdde
			}
		}
		__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Field(__obf_4040153f7700644e)
	}

	return __obf_63bbcee86d44fdde
}

func __obf_4b8140a063da43be(__obf_63bbcee86d44fdde reflect.Value) (reflect.Value, bool) {
	if __obf_63bbcee86d44fdde.Kind() == reflect.Ptr {
		if __obf_63bbcee86d44fdde.IsNil() {
			if !__obf_63bbcee86d44fdde.CanSet() {
				return __obf_63bbcee86d44fdde, false
			}
			__obf_3e615019e9372c2c := __obf_63bbcee86d44fdde.Type().Elem()
			if __obf_3e615019e9372c2c.Kind() != reflect.Struct {
				return __obf_63bbcee86d44fdde, false
			}
			__obf_63bbcee86d44fdde.
				Set(__obf_bcf79cb8d15fbf1f(__obf_3e615019e9372c2c))
		}
		__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Elem()
	}
	return __obf_63bbcee86d44fdde, true
}
