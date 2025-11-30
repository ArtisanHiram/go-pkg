package __obf_703d23ba520c3295

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
	ToVal(__obf_b924538fffe5fd64 any)
	Get(__obf_267bdf615cb1b310 ...any) Any
	Size() int
	Keys() []string
	GetInterface() any
	WriteTo(__obf_9a34f62857fb1d1d *Stream)
}

type __obf_b2a2e31336581ab8 struct{}

func (any *__obf_b2a2e31336581ab8) Get(__obf_267bdf615cb1b310 ...any) Any {
	return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, fmt.Errorf("GetIndex %v from simple value", __obf_267bdf615cb1b310)}
}

func (any *__obf_b2a2e31336581ab8) Size() int {
	return 0
}

func (any *__obf_b2a2e31336581ab8) Keys() []string {
	return []string{}
}

func (any *__obf_b2a2e31336581ab8) ToVal(__obf_02ba706f4f104d71 any) {
	panic("not implemented")
}

// WrapInt32 turn int32 into Any interface
func WrapInt32(__obf_b924538fffe5fd64 int32) Any {
	return &__obf_592e58127a384a3b{__obf_b2a2e31336581ab8{}, __obf_b924538fffe5fd64}
}

// WrapInt64 turn int64 into Any interface
func WrapInt64(__obf_b924538fffe5fd64 int64) Any {
	return &__obf_1d6d72beae0ea33d{__obf_b2a2e31336581ab8{}, __obf_b924538fffe5fd64}
}

// WrapUint32 turn uint32 into Any interface
func WrapUint32(__obf_b924538fffe5fd64 uint32) Any {
	return &__obf_71f4ab36a81c97f3{__obf_b2a2e31336581ab8{}, __obf_b924538fffe5fd64}
}

// WrapUint64 turn uint64 into Any interface
func WrapUint64(__obf_b924538fffe5fd64 uint64) Any {
	return &__obf_c0a5fd87615531ba{__obf_b2a2e31336581ab8{}, __obf_b924538fffe5fd64}
}

// WrapFloat64 turn float64 into Any interface
func WrapFloat64(__obf_b924538fffe5fd64 float64) Any {
	return &__obf_ba7fcab5df6006b8{__obf_b2a2e31336581ab8{}, __obf_b924538fffe5fd64}
}

// WrapString turn string into Any interface
func WrapString(__obf_b924538fffe5fd64 string) Any {
	return &__obf_223737ac6afcb5d9{__obf_b2a2e31336581ab8{}, __obf_b924538fffe5fd64}
}

// Wrap turn a go object into Any interface
func Wrap(__obf_b924538fffe5fd64 any) Any {
	if __obf_b924538fffe5fd64 == nil {
		return &__obf_dcb3a6cf344e0d42{}
	}
	__obf_f89c36489b1a3604, __obf_181be11c2857d042 := __obf_b924538fffe5fd64.(Any)
	if __obf_181be11c2857d042 {
		return __obf_f89c36489b1a3604
	}
	__obf_3b7f6abbae19451e := reflect2.TypeOf(__obf_b924538fffe5fd64)
	switch __obf_3b7f6abbae19451e.Kind() {
	case reflect.Slice:
		return __obf_7484f9ce272651b2(__obf_b924538fffe5fd64)
	case reflect.Struct:
		return __obf_70c50c635ea444d2(__obf_b924538fffe5fd64)
	case reflect.Map:
		return __obf_495ed748a54702d1(__obf_b924538fffe5fd64)
	case reflect.String:
		return WrapString(__obf_b924538fffe5fd64.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return WrapInt32(int32(__obf_b924538fffe5fd64.(int)))
		}
		return WrapInt64(int64(__obf_b924538fffe5fd64.(int)))
	case reflect.Int8:
		return WrapInt32(int32(__obf_b924538fffe5fd64.(int8)))
	case reflect.Int16:
		return WrapInt32(int32(__obf_b924538fffe5fd64.(int16)))
	case reflect.Int32:
		return WrapInt32(__obf_b924538fffe5fd64.(int32))
	case reflect.Int64:
		return WrapInt64(__obf_b924538fffe5fd64.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return WrapUint32(uint32(__obf_b924538fffe5fd64.(uint)))
		}
		return WrapUint64(uint64(__obf_b924538fffe5fd64.(uint)))
	case reflect.Uintptr:
		if __obf_6927eb49141afe0a == 32 {
			return WrapUint32(uint32(__obf_b924538fffe5fd64.(uintptr)))
		}
		return WrapUint64(uint64(__obf_b924538fffe5fd64.(uintptr)))
	case reflect.Uint8:
		return WrapUint32(uint32(__obf_b924538fffe5fd64.(uint8)))
	case reflect.Uint16:
		return WrapUint32(uint32(__obf_b924538fffe5fd64.(uint16)))
	case reflect.Uint32:
		return WrapUint32(uint32(__obf_b924538fffe5fd64.(uint32)))
	case reflect.Uint64:
		return WrapUint64(__obf_b924538fffe5fd64.(uint64))
	case reflect.Float32:
		return WrapFloat64(float64(__obf_b924538fffe5fd64.(float32)))
	case reflect.Float64:
		return WrapFloat64(__obf_b924538fffe5fd64.(float64))
	case reflect.Bool:
		if __obf_b924538fffe5fd64.(bool) == true {
			return &__obf_201570b71d830df8{}
		}
		return &__obf_856b120117cd1f8f{}
	}
	return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, fmt.Errorf("unsupported type: %v", __obf_3b7f6abbae19451e)}
}

// ReadAny read next JSON element as an Any object. It is a better json.RawMessage.
func (__obf_47edab4c16a2d88d *Iterator) ReadAny() Any {
	return __obf_47edab4c16a2d88d.__obf_721b69294316727a()
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_721b69294316727a() Any {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	switch __obf_bd08f5161fff294a {
	case '"':
		__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
		return &__obf_223737ac6afcb5d9{__obf_b2a2e31336581ab8{}, __obf_47edab4c16a2d88d.ReadString()}
	case 'n':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l') // null
		return &__obf_dcb3a6cf344e0d42{}
	case 't':
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('r', 'u', 'e') // true
		return &__obf_201570b71d830df8{}
	case 'f':
		__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('a', 'l', 's', 'e') // false
		return &__obf_856b120117cd1f8f{}
	case '{':
		return __obf_47edab4c16a2d88d.__obf_1fb1e1d7faa1d407()
	case '[':
		return __obf_47edab4c16a2d88d.__obf_483ab794560566b4()
	case '-':
		return __obf_47edab4c16a2d88d.__obf_a429a786a1a5fb0a(false)
	case 0:
		return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, errors.New("input is empty")}
	default:
		return __obf_47edab4c16a2d88d.__obf_a429a786a1a5fb0a(true)
	}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_a429a786a1a5fb0a(__obf_08103b5fb6b6760e bool) Any {
	__obf_47edab4c16a2d88d.__obf_57875b3e9496553e(__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 - 1)
	__obf_47edab4c16a2d88d.__obf_91b8504ad35d5b75()
	__obf_29da2d468f52e915 := __obf_47edab4c16a2d88d.__obf_842667b1c2db6627()
	return &__obf_2444f00b77eded26{__obf_b2a2e31336581ab8{}, __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862, __obf_29da2d468f52e915, nil}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_1fb1e1d7faa1d407() Any {
	__obf_47edab4c16a2d88d.__obf_57875b3e9496553e(__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 - 1)
	__obf_47edab4c16a2d88d.__obf_e53a692fb1024439()
	__obf_29da2d468f52e915 := __obf_47edab4c16a2d88d.__obf_842667b1c2db6627()
	return &__obf_c7e0be96be0cd81d{__obf_b2a2e31336581ab8{}, __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862, __obf_29da2d468f52e915, nil}
}

func (__obf_47edab4c16a2d88d *Iterator) __obf_483ab794560566b4() Any {
	__obf_47edab4c16a2d88d.__obf_57875b3e9496553e(__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 - 1)
	__obf_47edab4c16a2d88d.__obf_93a7c6915e88f56c()
	__obf_29da2d468f52e915 := __obf_47edab4c16a2d88d.__obf_842667b1c2db6627()
	return &__obf_909c8add24236c9b{__obf_b2a2e31336581ab8{}, __obf_47edab4c16a2d88d.__obf_25bd4f754a37b862, __obf_29da2d468f52e915, nil}
}

func __obf_29b231eb99e7ebaa(__obf_47edab4c16a2d88d *Iterator, __obf_2d2eb82817864725 string) []byte {
	var __obf_87a836f73d31233b []byte
	__obf_47edab4c16a2d88d.
		ReadObjectCB(func(__obf_47edab4c16a2d88d *Iterator, __obf_13f6440419533990 string) bool {
			if __obf_13f6440419533990 == __obf_2d2eb82817864725 {
				__obf_87a836f73d31233b = __obf_47edab4c16a2d88d.SkipAndReturnBytes()
				return false
			}
			__obf_47edab4c16a2d88d.
				Skip()
			return true
		})
	return __obf_87a836f73d31233b
}

func __obf_275bd1a9c11e5d3d(__obf_47edab4c16a2d88d *Iterator, __obf_2d2eb82817864725 int) []byte {
	var __obf_87a836f73d31233b []byte
	__obf_716729e0f8e419ed := 0
	__obf_47edab4c16a2d88d.
		ReadArrayCB(func(__obf_47edab4c16a2d88d *Iterator) bool {
			if __obf_716729e0f8e419ed == __obf_2d2eb82817864725 {
				__obf_87a836f73d31233b = __obf_47edab4c16a2d88d.SkipAndReturnBytes()
				return false
			}
			__obf_47edab4c16a2d88d.
				Skip()
			__obf_716729e0f8e419ed++
			return true
		})
	return __obf_87a836f73d31233b
}

func __obf_da70fd4b038d543a(__obf_47edab4c16a2d88d *Iterator, __obf_267bdf615cb1b310 []any) Any {
	for __obf_b0a5d2bd48690f1d, __obf_645d73566afc38ac := range __obf_267bdf615cb1b310 {
		switch __obf_7e6488d24bb81e3b := __obf_645d73566afc38ac.(type) {
		case string:
			__obf_24e18de62be92d01 := __obf_29b231eb99e7ebaa(__obf_47edab4c16a2d88d, __obf_7e6488d24bb81e3b)
			if __obf_24e18de62be92d01 == nil {
				return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310[__obf_b0a5d2bd48690f1d:])
			}
			__obf_47edab4c16a2d88d.
				ResetBytes(__obf_24e18de62be92d01)
		case int:
			__obf_24e18de62be92d01 := __obf_275bd1a9c11e5d3d(__obf_47edab4c16a2d88d, __obf_7e6488d24bb81e3b)
			if __obf_24e18de62be92d01 == nil {
				return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310[__obf_b0a5d2bd48690f1d:])
			}
			__obf_47edab4c16a2d88d.
				ResetBytes(__obf_24e18de62be92d01)
		case int32:
			if '*' == __obf_7e6488d24bb81e3b {
				return __obf_47edab4c16a2d88d.__obf_721b69294316727a().Get(__obf_267bdf615cb1b310[__obf_b0a5d2bd48690f1d:]...)
			}
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310[__obf_b0a5d2bd48690f1d:])
		default:
			return __obf_b82811e6c38b9e10(__obf_267bdf615cb1b310[__obf_b0a5d2bd48690f1d:])
		}
	}
	if __obf_47edab4c16a2d88d.Error != nil && __obf_47edab4c16a2d88d.Error != io.EOF {
		return &__obf_5d650408c7f7f591{__obf_b2a2e31336581ab8{}, __obf_47edab4c16a2d88d.Error}
	}
	return __obf_47edab4c16a2d88d.__obf_721b69294316727a()
}

var __obf_290a7d6480533c63 = reflect2.TypeOfPtr((*Any)(nil)).Elem()

func __obf_9e92fe737ca75ec6(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	if __obf_3b7f6abbae19451e == __obf_290a7d6480533c63 {
		return &__obf_43c2be15e32892e0{}
	}
	if __obf_3b7f6abbae19451e.Implements(__obf_290a7d6480533c63) {
		return &__obf_845bac373c0b0da4{__obf_15da62f311934a45: __obf_3b7f6abbae19451e}
	}
	return nil
}

func __obf_733be021e2b2eaf7(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	if __obf_3b7f6abbae19451e == __obf_290a7d6480533c63 {
		return &__obf_43c2be15e32892e0{}
	}
	if __obf_3b7f6abbae19451e.Implements(__obf_290a7d6480533c63) {
		return &__obf_845bac373c0b0da4{__obf_15da62f311934a45: __obf_3b7f6abbae19451e}
	}
	return nil
}

type __obf_845bac373c0b0da4 struct {
	__obf_15da62f311934a45 reflect2.Type
}

func (__obf_4acad06eb5535907 *__obf_845bac373c0b0da4) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	panic("not implemented")
}

func (__obf_4acad06eb5535907 *__obf_845bac373c0b0da4) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_02ba706f4f104d71 := __obf_4acad06eb5535907.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	any := __obf_02ba706f4f104d71.(Any)
	any.WriteTo(__obf_9a34f62857fb1d1d)
}

func (__obf_4acad06eb5535907 *__obf_845bac373c0b0da4) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_02ba706f4f104d71 := __obf_4acad06eb5535907.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	any := __obf_02ba706f4f104d71.(Any)
	return any.Size() == 0
}

type __obf_43c2be15e32892e0 struct {
}

func (__obf_4acad06eb5535907 *__obf_43c2be15e32892e0) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	*(*Any)(__obf_47fa53f3d161f45c) = __obf_47edab4c16a2d88d.__obf_721b69294316727a()
}

func (__obf_4acad06eb5535907 *__obf_43c2be15e32892e0) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	any := *(*Any)(__obf_47fa53f3d161f45c)
	if any == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	any.WriteTo(__obf_9a34f62857fb1d1d)
}

func (__obf_4acad06eb5535907 *__obf_43c2be15e32892e0) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	any := *(*Any)(__obf_47fa53f3d161f45c)
	return any.Size() == 0
}
