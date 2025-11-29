package __obf_91620b895eeff9ed

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
	ToVal(__obf_bbfd2ac8618a6f0c any)
	Get(__obf_82c50fcbfc9ab432 ...any) Any
	Size() int
	Keys() []string
	GetInterface() any
	WriteTo(__obf_850a7457bb739a32 *Stream)
}

type __obf_58563642f42f4a04 struct{}

func (any *__obf_58563642f42f4a04) Get(__obf_82c50fcbfc9ab432 ...any) Any {
	return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, fmt.Errorf("GetIndex %v from simple value", __obf_82c50fcbfc9ab432)}
}

func (any *__obf_58563642f42f4a04) Size() int {
	return 0
}

func (any *__obf_58563642f42f4a04) Keys() []string {
	return []string{}
}

func (any *__obf_58563642f42f4a04) ToVal(__obf_35136ef2ac0f94e4 any) {
	panic("not implemented")
}

// WrapInt32 turn int32 into Any interface
func WrapInt32(__obf_bbfd2ac8618a6f0c int32) Any {
	return &__obf_aad087210be33f6a{__obf_58563642f42f4a04{}, __obf_bbfd2ac8618a6f0c}
}

// WrapInt64 turn int64 into Any interface
func WrapInt64(__obf_bbfd2ac8618a6f0c int64) Any {
	return &__obf_2a9d95e27a8da146{__obf_58563642f42f4a04{}, __obf_bbfd2ac8618a6f0c}
}

// WrapUint32 turn uint32 into Any interface
func WrapUint32(__obf_bbfd2ac8618a6f0c uint32) Any {
	return &__obf_935849676e112271{__obf_58563642f42f4a04{}, __obf_bbfd2ac8618a6f0c}
}

// WrapUint64 turn uint64 into Any interface
func WrapUint64(__obf_bbfd2ac8618a6f0c uint64) Any {
	return &__obf_d8d434196b35220b{__obf_58563642f42f4a04{}, __obf_bbfd2ac8618a6f0c}
}

// WrapFloat64 turn float64 into Any interface
func WrapFloat64(__obf_bbfd2ac8618a6f0c float64) Any {
	return &__obf_36d93f8a7e11e503{__obf_58563642f42f4a04{}, __obf_bbfd2ac8618a6f0c}
}

// WrapString turn string into Any interface
func WrapString(__obf_bbfd2ac8618a6f0c string) Any {
	return &__obf_e76c1c0259373083{__obf_58563642f42f4a04{}, __obf_bbfd2ac8618a6f0c}
}

// Wrap turn a go object into Any interface
func Wrap(__obf_bbfd2ac8618a6f0c any) Any {
	if __obf_bbfd2ac8618a6f0c == nil {
		return &__obf_314389891734b872{}
	}
	__obf_a988d27da18886a3, __obf_1f0451ad1ef5d54f := __obf_bbfd2ac8618a6f0c.(Any)
	if __obf_1f0451ad1ef5d54f {
		return __obf_a988d27da18886a3
	}
	__obf_29ebd0f2c324f5ea := reflect2.TypeOf(__obf_bbfd2ac8618a6f0c)
	switch __obf_29ebd0f2c324f5ea.Kind() {
	case reflect.Slice:
		return __obf_856f321acac25752(__obf_bbfd2ac8618a6f0c)
	case reflect.Struct:
		return __obf_1dcf4efa25cd71a9(__obf_bbfd2ac8618a6f0c)
	case reflect.Map:
		return __obf_200d6b32a72d70af(__obf_bbfd2ac8618a6f0c)
	case reflect.String:
		return WrapString(__obf_bbfd2ac8618a6f0c.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return WrapInt32(int32(__obf_bbfd2ac8618a6f0c.(int)))
		}
		return WrapInt64(int64(__obf_bbfd2ac8618a6f0c.(int)))
	case reflect.Int8:
		return WrapInt32(int32(__obf_bbfd2ac8618a6f0c.(int8)))
	case reflect.Int16:
		return WrapInt32(int32(__obf_bbfd2ac8618a6f0c.(int16)))
	case reflect.Int32:
		return WrapInt32(__obf_bbfd2ac8618a6f0c.(int32))
	case reflect.Int64:
		return WrapInt64(__obf_bbfd2ac8618a6f0c.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return WrapUint32(uint32(__obf_bbfd2ac8618a6f0c.(uint)))
		}
		return WrapUint64(uint64(__obf_bbfd2ac8618a6f0c.(uint)))
	case reflect.Uintptr:
		if __obf_848460bf3389f562 == 32 {
			return WrapUint32(uint32(__obf_bbfd2ac8618a6f0c.(uintptr)))
		}
		return WrapUint64(uint64(__obf_bbfd2ac8618a6f0c.(uintptr)))
	case reflect.Uint8:
		return WrapUint32(uint32(__obf_bbfd2ac8618a6f0c.(uint8)))
	case reflect.Uint16:
		return WrapUint32(uint32(__obf_bbfd2ac8618a6f0c.(uint16)))
	case reflect.Uint32:
		return WrapUint32(uint32(__obf_bbfd2ac8618a6f0c.(uint32)))
	case reflect.Uint64:
		return WrapUint64(__obf_bbfd2ac8618a6f0c.(uint64))
	case reflect.Float32:
		return WrapFloat64(float64(__obf_bbfd2ac8618a6f0c.(float32)))
	case reflect.Float64:
		return WrapFloat64(__obf_bbfd2ac8618a6f0c.(float64))
	case reflect.Bool:
		if __obf_bbfd2ac8618a6f0c.(bool) == true {
			return &__obf_b9f1ac879e5e5b7d{}
		}
		return &__obf_e7b6b7f71c2987c8{}
	}
	return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, fmt.Errorf("unsupported type: %v", __obf_29ebd0f2c324f5ea)}
}

// ReadAny read next JSON element as an Any object. It is a better json.RawMessage.
func (__obf_1bb30e8a74ed8233 *Iterator) ReadAny() Any {
	return __obf_1bb30e8a74ed8233.__obf_6ec0844a84418f37()
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_6ec0844a84418f37() Any {
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	switch __obf_f16b4157911bc9af {
	case '"':
		__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
		return &__obf_e76c1c0259373083{__obf_58563642f42f4a04{}, __obf_1bb30e8a74ed8233.ReadString()}
	case 'n':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('u', 'l', 'l') // null
		return &__obf_314389891734b872{}
	case 't':
		__obf_1bb30e8a74ed8233.__obf_3e1d2ad9a54f0d22('r', 'u', 'e') // true
		return &__obf_b9f1ac879e5e5b7d{}
	case 'f':
		__obf_1bb30e8a74ed8233.__obf_94a8e39928c8440c('a', 'l', 's', 'e') // false
		return &__obf_e7b6b7f71c2987c8{}
	case '{':
		return __obf_1bb30e8a74ed8233.__obf_f587169d3d2616cb()
	case '[':
		return __obf_1bb30e8a74ed8233.__obf_a3cd79cf641fd4d6()
	case '-':
		return __obf_1bb30e8a74ed8233.__obf_ab81d4dd0c0f7e05(false)
	case 0:
		return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, errors.New("input is empty")}
	default:
		return __obf_1bb30e8a74ed8233.__obf_ab81d4dd0c0f7e05(true)
	}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_ab81d4dd0c0f7e05(__obf_987c73e8313ce202 bool) Any {
	__obf_1bb30e8a74ed8233.__obf_c704d25028ca6789(__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 - 1)
	__obf_1bb30e8a74ed8233.__obf_d6a474bd319b770c()
	__obf_f9451d4138bf17d3 := __obf_1bb30e8a74ed8233.__obf_2141069796e7e215()
	return &__obf_37d3c66ca0437bb5{__obf_58563642f42f4a04{}, __obf_1bb30e8a74ed8233.__obf_892632c77e859037, __obf_f9451d4138bf17d3, nil}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_f587169d3d2616cb() Any {
	__obf_1bb30e8a74ed8233.__obf_c704d25028ca6789(__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 - 1)
	__obf_1bb30e8a74ed8233.__obf_8d12ebf64f0517d2()
	__obf_f9451d4138bf17d3 := __obf_1bb30e8a74ed8233.__obf_2141069796e7e215()
	return &__obf_9460519df17f70ab{__obf_58563642f42f4a04{}, __obf_1bb30e8a74ed8233.__obf_892632c77e859037, __obf_f9451d4138bf17d3, nil}
}

func (__obf_1bb30e8a74ed8233 *Iterator) __obf_a3cd79cf641fd4d6() Any {
	__obf_1bb30e8a74ed8233.__obf_c704d25028ca6789(__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 - 1)
	__obf_1bb30e8a74ed8233.__obf_531ac82b5de39aec()
	__obf_f9451d4138bf17d3 := __obf_1bb30e8a74ed8233.__obf_2141069796e7e215()
	return &__obf_8759b49e282d9de6{__obf_58563642f42f4a04{}, __obf_1bb30e8a74ed8233.__obf_892632c77e859037, __obf_f9451d4138bf17d3, nil}
}

func __obf_820a957cfe88153b(__obf_1bb30e8a74ed8233 *Iterator, __obf_0063094641d09a49 string) []byte {
	var __obf_93f189bffabaf5a4 []byte
	__obf_1bb30e8a74ed8233.
		ReadObjectCB(func(__obf_1bb30e8a74ed8233 *Iterator, __obf_7e01b5b4651074d4 string) bool {
			if __obf_7e01b5b4651074d4 == __obf_0063094641d09a49 {
				__obf_93f189bffabaf5a4 = __obf_1bb30e8a74ed8233.SkipAndReturnBytes()
				return false
			}
			__obf_1bb30e8a74ed8233.
				Skip()
			return true
		})
	return __obf_93f189bffabaf5a4
}

func __obf_3f3216ff423a5303(__obf_1bb30e8a74ed8233 *Iterator, __obf_0063094641d09a49 int) []byte {
	var __obf_93f189bffabaf5a4 []byte
	__obf_c80a670e818798d8 := 0
	__obf_1bb30e8a74ed8233.
		ReadArrayCB(func(__obf_1bb30e8a74ed8233 *Iterator) bool {
			if __obf_c80a670e818798d8 == __obf_0063094641d09a49 {
				__obf_93f189bffabaf5a4 = __obf_1bb30e8a74ed8233.SkipAndReturnBytes()
				return false
			}
			__obf_1bb30e8a74ed8233.
				Skip()
			__obf_c80a670e818798d8++
			return true
		})
	return __obf_93f189bffabaf5a4
}

func __obf_67ce74ddb71f8159(__obf_1bb30e8a74ed8233 *Iterator, __obf_82c50fcbfc9ab432 []any) Any {
	for __obf_5aa5c8829b97f182, __obf_decbb8bb80060e30 := range __obf_82c50fcbfc9ab432 {
		switch __obf_ab92e8bcb9e5b26a := __obf_decbb8bb80060e30.(type) {
		case string:
			__obf_930565f5477d9c6d := __obf_820a957cfe88153b(__obf_1bb30e8a74ed8233, __obf_ab92e8bcb9e5b26a)
			if __obf_930565f5477d9c6d == nil {
				return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432[__obf_5aa5c8829b97f182:])
			}
			__obf_1bb30e8a74ed8233.
				ResetBytes(__obf_930565f5477d9c6d)
		case int:
			__obf_930565f5477d9c6d := __obf_3f3216ff423a5303(__obf_1bb30e8a74ed8233, __obf_ab92e8bcb9e5b26a)
			if __obf_930565f5477d9c6d == nil {
				return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432[__obf_5aa5c8829b97f182:])
			}
			__obf_1bb30e8a74ed8233.
				ResetBytes(__obf_930565f5477d9c6d)
		case int32:
			if '*' == __obf_ab92e8bcb9e5b26a {
				return __obf_1bb30e8a74ed8233.__obf_6ec0844a84418f37().Get(__obf_82c50fcbfc9ab432[__obf_5aa5c8829b97f182:]...)
			}
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432[__obf_5aa5c8829b97f182:])
		default:
			return __obf_227499e666173bc3(__obf_82c50fcbfc9ab432[__obf_5aa5c8829b97f182:])
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		return &__obf_a4f671452051c765{__obf_58563642f42f4a04{}, __obf_1bb30e8a74ed8233.Error}
	}
	return __obf_1bb30e8a74ed8233.__obf_6ec0844a84418f37()
}

var __obf_7281516618f0c380 = reflect2.TypeOfPtr((*Any)(nil)).Elem()

func __obf_9aa737014fad08d9(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	if __obf_29ebd0f2c324f5ea == __obf_7281516618f0c380 {
		return &__obf_0e396b156f331243{}
	}
	if __obf_29ebd0f2c324f5ea.Implements(__obf_7281516618f0c380) {
		return &__obf_699352e366642e3f{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea}
	}
	return nil
}

func __obf_d3e2baf3cccfde91(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValEncoder {
	if __obf_29ebd0f2c324f5ea == __obf_7281516618f0c380 {
		return &__obf_0e396b156f331243{}
	}
	if __obf_29ebd0f2c324f5ea.Implements(__obf_7281516618f0c380) {
		return &__obf_699352e366642e3f{__obf_5ea8ee327a6f7e0d: __obf_29ebd0f2c324f5ea}
	}
	return nil
}

type __obf_699352e366642e3f struct {
	__obf_5ea8ee327a6f7e0d reflect2.Type
}

func (__obf_be93ede593e1d304 *__obf_699352e366642e3f) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	panic("not implemented")
}

func (__obf_be93ede593e1d304 *__obf_699352e366642e3f) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	__obf_35136ef2ac0f94e4 := __obf_be93ede593e1d304.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	any := __obf_35136ef2ac0f94e4.(Any)
	any.WriteTo(__obf_850a7457bb739a32)
}

func (__obf_be93ede593e1d304 *__obf_699352e366642e3f) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	__obf_35136ef2ac0f94e4 := __obf_be93ede593e1d304.__obf_5ea8ee327a6f7e0d.UnsafeIndirect(__obf_2a1474febb02279b)
	any := __obf_35136ef2ac0f94e4.(Any)
	return any.Size() == 0
}

type __obf_0e396b156f331243 struct {
}

func (__obf_be93ede593e1d304 *__obf_0e396b156f331243) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	*(*Any)(__obf_2a1474febb02279b) = __obf_1bb30e8a74ed8233.__obf_6ec0844a84418f37()
}

func (__obf_be93ede593e1d304 *__obf_0e396b156f331243) Encode(__obf_2a1474febb02279b unsafe.Pointer, __obf_850a7457bb739a32 *Stream) {
	any := *(*Any)(__obf_2a1474febb02279b)
	if any == nil {
		__obf_850a7457bb739a32.
			WriteNil()
		return
	}
	any.WriteTo(__obf_850a7457bb739a32)
}

func (__obf_be93ede593e1d304 *__obf_0e396b156f331243) IsEmpty(__obf_2a1474febb02279b unsafe.Pointer) bool {
	any := *(*Any)(__obf_2a1474febb02279b)
	return any.Size() == 0
}
