package __obf_030d39f7a8de96c6

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

// Any generic object representation.
// The lazy json implementation holds []byte and parse lazily.
type Any interface {
	LastError() error
	ValueType() ValueType
	MustBeValid() Any
	ToBool() bool
	ToInt() int
	ToInt32() int32
	ToInt64() int64
	ToUint() uint
	ToUint32() uint32
	ToUint64() uint64
	ToFloat32() float32
	ToFloat64() float64
	ToString() string
	ToVal(__obf_efaf2719b44ad8ac any)
	Get(__obf_f77a9507782b919d ...any) Any
	Size() int
	Keys() []string
	GetInterface() any
	WriteTo(__obf_8a2c51fe22775655 *Stream)
}

type __obf_7834a13a259af712 struct{}

func (any *__obf_7834a13a259af712) Get(__obf_f77a9507782b919d ...any) Any {
	return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, fmt.Errorf("GetIndex %v from simple value", __obf_f77a9507782b919d)}
}

func (any *__obf_7834a13a259af712) Size() int {
	return 0
}

func (any *__obf_7834a13a259af712) Keys() []string {
	return []string{}
}

func (any *__obf_7834a13a259af712) ToVal(__obf_eefa92392cc2442c any) {
	panic("not implemented")
}

// WrapInt32 turn int32 into Any interface
func WrapInt32(__obf_efaf2719b44ad8ac int32) Any {
	return &__obf_66e4549b4c63a1c1{__obf_7834a13a259af712{}, __obf_efaf2719b44ad8ac}
}

// WrapInt64 turn int64 into Any interface
func WrapInt64(__obf_efaf2719b44ad8ac int64) Any {
	return &__obf_9108c7af265fa584{__obf_7834a13a259af712{}, __obf_efaf2719b44ad8ac}
}

// WrapUint32 turn uint32 into Any interface
func WrapUint32(__obf_efaf2719b44ad8ac uint32) Any {
	return &__obf_99e4188e77cbabaa{__obf_7834a13a259af712{}, __obf_efaf2719b44ad8ac}
}

// WrapUint64 turn uint64 into Any interface
func WrapUint64(__obf_efaf2719b44ad8ac uint64) Any {
	return &__obf_4bd84d2e4678d344{__obf_7834a13a259af712{}, __obf_efaf2719b44ad8ac}
}

// WrapFloat64 turn float64 into Any interface
func WrapFloat64(__obf_efaf2719b44ad8ac float64) Any {
	return &__obf_07c9a175e2e50f0d{__obf_7834a13a259af712{}, __obf_efaf2719b44ad8ac}
}

// WrapString turn string into Any interface
func WrapString(__obf_efaf2719b44ad8ac string) Any {
	return &__obf_b2998b74708f8d59{__obf_7834a13a259af712{}, __obf_efaf2719b44ad8ac}
}

// Wrap turn a go object into Any interface
func Wrap(__obf_efaf2719b44ad8ac any) Any {
	if __obf_efaf2719b44ad8ac == nil {
		return &__obf_bc31f79e0dcee722{}
	}
	__obf_ac95e79e9d5acab0, __obf_bc7073bb09c5a3ae := __obf_efaf2719b44ad8ac.(Any)
	if __obf_bc7073bb09c5a3ae {
		return __obf_ac95e79e9d5acab0
	}
	__obf_8744d0b8c80ccdc1 := reflect2.TypeOf(__obf_efaf2719b44ad8ac)
	switch __obf_8744d0b8c80ccdc1.Kind() {
	case reflect.Slice:
		return __obf_eb0024afd144815a(__obf_efaf2719b44ad8ac)
	case reflect.Struct:
		return __obf_0cb32cddab75abc0(__obf_efaf2719b44ad8ac)
	case reflect.Map:
		return __obf_1dc61d22b3defb62(__obf_efaf2719b44ad8ac)
	case reflect.String:
		return WrapString(__obf_efaf2719b44ad8ac.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return WrapInt32(int32(__obf_efaf2719b44ad8ac.(int)))
		}
		return WrapInt64(int64(__obf_efaf2719b44ad8ac.(int)))
	case reflect.Int8:
		return WrapInt32(int32(__obf_efaf2719b44ad8ac.(int8)))
	case reflect.Int16:
		return WrapInt32(int32(__obf_efaf2719b44ad8ac.(int16)))
	case reflect.Int32:
		return WrapInt32(__obf_efaf2719b44ad8ac.(int32))
	case reflect.Int64:
		return WrapInt64(__obf_efaf2719b44ad8ac.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return WrapUint32(uint32(__obf_efaf2719b44ad8ac.(uint)))
		}
		return WrapUint64(uint64(__obf_efaf2719b44ad8ac.(uint)))
	case reflect.Uintptr:
		if __obf_eb8e424502869741 == 32 {
			return WrapUint32(uint32(__obf_efaf2719b44ad8ac.(uintptr)))
		}
		return WrapUint64(uint64(__obf_efaf2719b44ad8ac.(uintptr)))
	case reflect.Uint8:
		return WrapUint32(uint32(__obf_efaf2719b44ad8ac.(uint8)))
	case reflect.Uint16:
		return WrapUint32(uint32(__obf_efaf2719b44ad8ac.(uint16)))
	case reflect.Uint32:
		return WrapUint32(uint32(__obf_efaf2719b44ad8ac.(uint32)))
	case reflect.Uint64:
		return WrapUint64(__obf_efaf2719b44ad8ac.(uint64))
	case reflect.Float32:
		return WrapFloat64(float64(__obf_efaf2719b44ad8ac.(float32)))
	case reflect.Float64:
		return WrapFloat64(__obf_efaf2719b44ad8ac.(float64))
	case reflect.Bool:
		if __obf_efaf2719b44ad8ac.(bool) == true {
			return &__obf_c1e44b1851e558dd{}
		}
		return &__obf_391e15815a33c392{}
	}
	return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, fmt.Errorf("unsupported type: %v", __obf_8744d0b8c80ccdc1)}
}

// ReadAny read next JSON element as an Any object. It is a better json.RawMessage.
func (__obf_4ab56a99965952e7 *Iterator) ReadAny() Any {
	return __obf_4ab56a99965952e7.__obf_d0ec77416ba1caaa()
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_d0ec77416ba1caaa() Any {
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	switch __obf_24309b3b0ff9ba22 {
	case '"':
		__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
		return &__obf_b2998b74708f8d59{__obf_7834a13a259af712{}, __obf_4ab56a99965952e7.ReadString()}
	case 'n':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('u', 'l', 'l') // null
		return &__obf_bc31f79e0dcee722{}
	case 't':
		__obf_4ab56a99965952e7.__obf_01038f7a5ba9ca56('r', 'u', 'e') // true
		return &__obf_c1e44b1851e558dd{}
	case 'f':
		__obf_4ab56a99965952e7.__obf_aaf95589108acb4b('a', 'l', 's', 'e') // false
		return &__obf_391e15815a33c392{}
	case '{':
		return __obf_4ab56a99965952e7.__obf_2cf622e65c9b3e05()
	case '[':
		return __obf_4ab56a99965952e7.__obf_38ad096faefce2bf()
	case '-':
		return __obf_4ab56a99965952e7.__obf_09326eb045640050(false)
	case 0:
		return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, errors.New("input is empty")}
	default:
		return __obf_4ab56a99965952e7.__obf_09326eb045640050(true)
	}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_09326eb045640050(__obf_bb7955e825ad5b14 bool) Any {
	__obf_4ab56a99965952e7.__obf_f1e0f594888555ba(__obf_4ab56a99965952e7.__obf_5e22d6270d27491f - 1)
	__obf_4ab56a99965952e7.__obf_38c33066e8f2b292()
	__obf_b2b1f67a8936496a := __obf_4ab56a99965952e7.__obf_4403a37119edc6b3()
	return &__obf_0796d07d4885e149{__obf_7834a13a259af712{}, __obf_4ab56a99965952e7.__obf_679611dc56ff465b, __obf_b2b1f67a8936496a, nil}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_2cf622e65c9b3e05() Any {
	__obf_4ab56a99965952e7.__obf_f1e0f594888555ba(__obf_4ab56a99965952e7.__obf_5e22d6270d27491f - 1)
	__obf_4ab56a99965952e7.__obf_56cafb6b2d847244()
	__obf_b2b1f67a8936496a := __obf_4ab56a99965952e7.__obf_4403a37119edc6b3()
	return &__obf_93b2c1e9efd325de{__obf_7834a13a259af712{}, __obf_4ab56a99965952e7.__obf_679611dc56ff465b, __obf_b2b1f67a8936496a, nil}
}

func (__obf_4ab56a99965952e7 *Iterator) __obf_38ad096faefce2bf() Any {
	__obf_4ab56a99965952e7.__obf_f1e0f594888555ba(__obf_4ab56a99965952e7.__obf_5e22d6270d27491f - 1)
	__obf_4ab56a99965952e7.__obf_e8e6ee2ef6c27f13()
	__obf_b2b1f67a8936496a := __obf_4ab56a99965952e7.__obf_4403a37119edc6b3()
	return &__obf_6861b65fa3b114d9{__obf_7834a13a259af712{}, __obf_4ab56a99965952e7.__obf_679611dc56ff465b, __obf_b2b1f67a8936496a, nil}
}

func __obf_d17dcb541db3da34(__obf_4ab56a99965952e7 *Iterator, __obf_a6ddbe69238cec9a string) []byte {
	var __obf_ccdc4ee8b9d2f52f []byte
	__obf_4ab56a99965952e7.
		ReadObjectCB(func(__obf_4ab56a99965952e7 *Iterator, __obf_cd4d02f264b18d55 string) bool {
			if __obf_cd4d02f264b18d55 == __obf_a6ddbe69238cec9a {
				__obf_ccdc4ee8b9d2f52f = __obf_4ab56a99965952e7.SkipAndReturnBytes()
				return false
			}
			__obf_4ab56a99965952e7.
				Skip()
			return true
		})
	return __obf_ccdc4ee8b9d2f52f
}

func __obf_6f7138d64db3c315(__obf_4ab56a99965952e7 *Iterator, __obf_a6ddbe69238cec9a int) []byte {
	var __obf_ccdc4ee8b9d2f52f []byte
	__obf_e40b3963a92ca4a0 := 0
	__obf_4ab56a99965952e7.
		ReadArrayCB(func(__obf_4ab56a99965952e7 *Iterator) bool {
			if __obf_e40b3963a92ca4a0 == __obf_a6ddbe69238cec9a {
				__obf_ccdc4ee8b9d2f52f = __obf_4ab56a99965952e7.SkipAndReturnBytes()
				return false
			}
			__obf_4ab56a99965952e7.
				Skip()
			__obf_e40b3963a92ca4a0++
			return true
		})
	return __obf_ccdc4ee8b9d2f52f
}

func __obf_4cd15630eeb40546(__obf_4ab56a99965952e7 *Iterator, __obf_f77a9507782b919d []any) Any {
	for __obf_82c6e05b9d226c58, __obf_d1d1ccc388cccd57 := range __obf_f77a9507782b919d {
		switch __obf_a29eed7785a10cc5 := __obf_d1d1ccc388cccd57.(type) {
		case string:
			__obf_d2de1fa55acc5ff0 := __obf_d17dcb541db3da34(__obf_4ab56a99965952e7, __obf_a29eed7785a10cc5)
			if __obf_d2de1fa55acc5ff0 == nil {
				return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d[__obf_82c6e05b9d226c58:])
			}
			__obf_4ab56a99965952e7.
				ResetBytes(__obf_d2de1fa55acc5ff0)
		case int:
			__obf_d2de1fa55acc5ff0 := __obf_6f7138d64db3c315(__obf_4ab56a99965952e7, __obf_a29eed7785a10cc5)
			if __obf_d2de1fa55acc5ff0 == nil {
				return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d[__obf_82c6e05b9d226c58:])
			}
			__obf_4ab56a99965952e7.
				ResetBytes(__obf_d2de1fa55acc5ff0)
		case int32:
			if '*' == __obf_a29eed7785a10cc5 {
				return __obf_4ab56a99965952e7.__obf_d0ec77416ba1caaa().Get(__obf_f77a9507782b919d[__obf_82c6e05b9d226c58:]...)
			}
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d[__obf_82c6e05b9d226c58:])
		default:
			return __obf_0adb97f96f4a14bd(__obf_f77a9507782b919d[__obf_82c6e05b9d226c58:])
		}
	}
	if __obf_4ab56a99965952e7.Error != nil && __obf_4ab56a99965952e7.Error != io.EOF {
		return &__obf_aa2610a72837656e{__obf_7834a13a259af712{}, __obf_4ab56a99965952e7.Error}
	}
	return __obf_4ab56a99965952e7.__obf_d0ec77416ba1caaa()
}

var __obf_22a394224baf1954 = reflect2.TypeOfPtr((*Any)(nil)).Elem()

func __obf_b457e3c1bf18f46a(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValDecoder {
	if __obf_8744d0b8c80ccdc1 == __obf_22a394224baf1954 {
		return &__obf_1a52d94f786e39ea{}
	}
	if __obf_8744d0b8c80ccdc1.Implements(__obf_22a394224baf1954) {
		return &__obf_9b08aab04913187d{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1}
	}
	return nil
}

func __obf_9ed4e77098f589a9(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	if __obf_8744d0b8c80ccdc1 == __obf_22a394224baf1954 {
		return &__obf_1a52d94f786e39ea{}
	}
	if __obf_8744d0b8c80ccdc1.Implements(__obf_22a394224baf1954) {
		return &__obf_9b08aab04913187d{__obf_a7610e23a63622fd: __obf_8744d0b8c80ccdc1}
	}
	return nil
}

type __obf_9b08aab04913187d struct {
	__obf_a7610e23a63622fd reflect2.Type
}

func (__obf_7bfaa2f467491ab9 *__obf_9b08aab04913187d) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	panic("not implemented")
}

func (__obf_7bfaa2f467491ab9 *__obf_9b08aab04913187d) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_eefa92392cc2442c := __obf_7bfaa2f467491ab9.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	any := __obf_eefa92392cc2442c.(Any)
	any.WriteTo(__obf_8a2c51fe22775655)
}

func (__obf_7bfaa2f467491ab9 *__obf_9b08aab04913187d) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_eefa92392cc2442c := __obf_7bfaa2f467491ab9.__obf_a7610e23a63622fd.UnsafeIndirect(__obf_dbbf371b91136494)
	any := __obf_eefa92392cc2442c.(Any)
	return any.Size() == 0
}

type __obf_1a52d94f786e39ea struct {
}

func (__obf_7bfaa2f467491ab9 *__obf_1a52d94f786e39ea) Decode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_4ab56a99965952e7 *Iterator) {
	*(*Any)(__obf_dbbf371b91136494) = __obf_4ab56a99965952e7.__obf_d0ec77416ba1caaa()
}

func (__obf_7bfaa2f467491ab9 *__obf_1a52d94f786e39ea) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	any := *(*Any)(__obf_dbbf371b91136494)
	if any == nil {
		__obf_8a2c51fe22775655.
			WriteNil()
		return
	}
	any.WriteTo(__obf_8a2c51fe22775655)
}

func (__obf_7bfaa2f467491ab9 *__obf_1a52d94f786e39ea) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	any := *(*Any)(__obf_dbbf371b91136494)
	return any.Size() == 0
}
