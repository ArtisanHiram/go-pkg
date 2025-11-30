package __obf_c3cd04a15c56f32f

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
	ToVal(__obf_d59813f23ad740a8 any)
	Get(__obf_b90e4ca332607787 ...any) Any
	Size() int
	Keys() []string
	GetInterface() any
	WriteTo(__obf_2361f5214e84c60f *Stream)
}

type __obf_86279f81acb381f3 struct{}

func (any *__obf_86279f81acb381f3) Get(__obf_b90e4ca332607787 ...any) Any {
	return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, fmt.Errorf("GetIndex %v from simple value", __obf_b90e4ca332607787)}
}

func (any *__obf_86279f81acb381f3) Size() int {
	return 0
}

func (any *__obf_86279f81acb381f3) Keys() []string {
	return []string{}
}

func (any *__obf_86279f81acb381f3) ToVal(__obf_50acae8c0a871ba1 any) {
	panic("not implemented")
}

// WrapInt32 turn int32 into Any interface
func WrapInt32(__obf_d59813f23ad740a8 int32) Any {
	return &__obf_c335eb1f76d110d7{__obf_86279f81acb381f3{}, __obf_d59813f23ad740a8}
}

// WrapInt64 turn int64 into Any interface
func WrapInt64(__obf_d59813f23ad740a8 int64) Any {
	return &__obf_e42a057d4a279be0{__obf_86279f81acb381f3{}, __obf_d59813f23ad740a8}
}

// WrapUint32 turn uint32 into Any interface
func WrapUint32(__obf_d59813f23ad740a8 uint32) Any {
	return &__obf_286b4b716339f0b1{__obf_86279f81acb381f3{}, __obf_d59813f23ad740a8}
}

// WrapUint64 turn uint64 into Any interface
func WrapUint64(__obf_d59813f23ad740a8 uint64) Any {
	return &__obf_226272cc8e4cb56e{__obf_86279f81acb381f3{}, __obf_d59813f23ad740a8}
}

// WrapFloat64 turn float64 into Any interface
func WrapFloat64(__obf_d59813f23ad740a8 float64) Any {
	return &__obf_299c6b904e6b8eea{__obf_86279f81acb381f3{}, __obf_d59813f23ad740a8}
}

// WrapString turn string into Any interface
func WrapString(__obf_d59813f23ad740a8 string) Any {
	return &__obf_7402622ee5d0a489{__obf_86279f81acb381f3{}, __obf_d59813f23ad740a8}
}

// Wrap turn a go object into Any interface
func Wrap(__obf_d59813f23ad740a8 any) Any {
	if __obf_d59813f23ad740a8 == nil {
		return &__obf_d86254f4bc75f537{}
	}
	__obf_ff6932802afd73dc, __obf_b0aa92f22e2502fa := __obf_d59813f23ad740a8.(Any)
	if __obf_b0aa92f22e2502fa {
		return __obf_ff6932802afd73dc
	}
	__obf_8667d4fc2a95b0d7 := reflect2.TypeOf(__obf_d59813f23ad740a8)
	switch __obf_8667d4fc2a95b0d7.Kind() {
	case reflect.Slice:
		return __obf_90999f5d50faa290(__obf_d59813f23ad740a8)
	case reflect.Struct:
		return __obf_5ed91bdc7ddeb0d6(__obf_d59813f23ad740a8)
	case reflect.Map:
		return __obf_af999c25758b0ff1(__obf_d59813f23ad740a8)
	case reflect.String:
		return WrapString(__obf_d59813f23ad740a8.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return WrapInt32(int32(__obf_d59813f23ad740a8.(int)))
		}
		return WrapInt64(int64(__obf_d59813f23ad740a8.(int)))
	case reflect.Int8:
		return WrapInt32(int32(__obf_d59813f23ad740a8.(int8)))
	case reflect.Int16:
		return WrapInt32(int32(__obf_d59813f23ad740a8.(int16)))
	case reflect.Int32:
		return WrapInt32(__obf_d59813f23ad740a8.(int32))
	case reflect.Int64:
		return WrapInt64(__obf_d59813f23ad740a8.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return WrapUint32(uint32(__obf_d59813f23ad740a8.(uint)))
		}
		return WrapUint64(uint64(__obf_d59813f23ad740a8.(uint)))
	case reflect.Uintptr:
		if __obf_26c0da34d2f86b82 == 32 {
			return WrapUint32(uint32(__obf_d59813f23ad740a8.(uintptr)))
		}
		return WrapUint64(uint64(__obf_d59813f23ad740a8.(uintptr)))
	case reflect.Uint8:
		return WrapUint32(uint32(__obf_d59813f23ad740a8.(uint8)))
	case reflect.Uint16:
		return WrapUint32(uint32(__obf_d59813f23ad740a8.(uint16)))
	case reflect.Uint32:
		return WrapUint32(uint32(__obf_d59813f23ad740a8.(uint32)))
	case reflect.Uint64:
		return WrapUint64(__obf_d59813f23ad740a8.(uint64))
	case reflect.Float32:
		return WrapFloat64(float64(__obf_d59813f23ad740a8.(float32)))
	case reflect.Float64:
		return WrapFloat64(__obf_d59813f23ad740a8.(float64))
	case reflect.Bool:
		if __obf_d59813f23ad740a8.(bool) == true {
			return &__obf_a745e9b75f5e2436{}
		}
		return &__obf_ed8192b5fd8f56eb{}
	}
	return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, fmt.Errorf("unsupported type: %v", __obf_8667d4fc2a95b0d7)}
}

// ReadAny read next JSON element as an Any object. It is a better json.RawMessage.
func (__obf_edd9fe48d39445e4 *Iterator) ReadAny() Any {
	return __obf_edd9fe48d39445e4.__obf_1da8db8a1d168ded()
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_1da8db8a1d168ded() Any {
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	switch __obf_0c1bc1e511a43120 {
	case '"':
		__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
		return &__obf_7402622ee5d0a489{__obf_86279f81acb381f3{}, __obf_edd9fe48d39445e4.ReadString()}
	case 'n':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('u', 'l', 'l') // null
		return &__obf_d86254f4bc75f537{}
	case 't':
		__obf_edd9fe48d39445e4.__obf_8bc1f4b8d62f5247('r', 'u', 'e') // true
		return &__obf_a745e9b75f5e2436{}
	case 'f':
		__obf_edd9fe48d39445e4.__obf_f22f308da0537336('a', 'l', 's', 'e') // false
		return &__obf_ed8192b5fd8f56eb{}
	case '{':
		return __obf_edd9fe48d39445e4.__obf_1a80efb9e58ea245()
	case '[':
		return __obf_edd9fe48d39445e4.__obf_094ccb0852fef332()
	case '-':
		return __obf_edd9fe48d39445e4.__obf_5851726f84f723f4(false)
	case 0:
		return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, errors.New("input is empty")}
	default:
		return __obf_edd9fe48d39445e4.__obf_5851726f84f723f4(true)
	}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_5851726f84f723f4(__obf_9e089f46e065c417 bool) Any {
	__obf_edd9fe48d39445e4.__obf_e786284811b1ebdd(__obf_edd9fe48d39445e4.__obf_edd3c3885447229b - 1)
	__obf_edd9fe48d39445e4.__obf_fdeaf1c2cbe3aa15()
	__obf_b0c073e871d9f7f4 := __obf_edd9fe48d39445e4.__obf_e27ad1d64f9218e7()
	return &__obf_76146a3e4f4d12ac{__obf_86279f81acb381f3{}, __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab, __obf_b0c073e871d9f7f4, nil}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_1a80efb9e58ea245() Any {
	__obf_edd9fe48d39445e4.__obf_e786284811b1ebdd(__obf_edd9fe48d39445e4.__obf_edd3c3885447229b - 1)
	__obf_edd9fe48d39445e4.__obf_0704e073eacf05e6()
	__obf_b0c073e871d9f7f4 := __obf_edd9fe48d39445e4.__obf_e27ad1d64f9218e7()
	return &__obf_0caad688333e6f4f{__obf_86279f81acb381f3{}, __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab, __obf_b0c073e871d9f7f4, nil}
}

func (__obf_edd9fe48d39445e4 *Iterator) __obf_094ccb0852fef332() Any {
	__obf_edd9fe48d39445e4.__obf_e786284811b1ebdd(__obf_edd9fe48d39445e4.__obf_edd3c3885447229b - 1)
	__obf_edd9fe48d39445e4.__obf_5566b22788a0a705()
	__obf_b0c073e871d9f7f4 := __obf_edd9fe48d39445e4.__obf_e27ad1d64f9218e7()
	return &__obf_1a2e99c79ce5a488{__obf_86279f81acb381f3{}, __obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab, __obf_b0c073e871d9f7f4, nil}
}

func __obf_931a5a6194f47dee(__obf_edd9fe48d39445e4 *Iterator, __obf_2551686f851ee25d string) []byte {
	var __obf_da9e40840d5dbfdf []byte
	__obf_edd9fe48d39445e4.
		ReadObjectCB(func(__obf_edd9fe48d39445e4 *Iterator, __obf_f992c91036745ca0 string) bool {
			if __obf_f992c91036745ca0 == __obf_2551686f851ee25d {
				__obf_da9e40840d5dbfdf = __obf_edd9fe48d39445e4.SkipAndReturnBytes()
				return false
			}
			__obf_edd9fe48d39445e4.
				Skip()
			return true
		})
	return __obf_da9e40840d5dbfdf
}

func __obf_4c9d33490f37e941(__obf_edd9fe48d39445e4 *Iterator, __obf_2551686f851ee25d int) []byte {
	var __obf_da9e40840d5dbfdf []byte
	__obf_fd4464bb6b13cabd := 0
	__obf_edd9fe48d39445e4.
		ReadArrayCB(func(__obf_edd9fe48d39445e4 *Iterator) bool {
			if __obf_fd4464bb6b13cabd == __obf_2551686f851ee25d {
				__obf_da9e40840d5dbfdf = __obf_edd9fe48d39445e4.SkipAndReturnBytes()
				return false
			}
			__obf_edd9fe48d39445e4.
				Skip()
			__obf_fd4464bb6b13cabd++
			return true
		})
	return __obf_da9e40840d5dbfdf
}

func __obf_9979e7313b51aeca(__obf_edd9fe48d39445e4 *Iterator, __obf_b90e4ca332607787 []any) Any {
	for __obf_28d099df85f083a8, __obf_a9ccfd2902ffa688 := range __obf_b90e4ca332607787 {
		switch __obf_bb1024339a5c1253 := __obf_a9ccfd2902ffa688.(type) {
		case string:
			__obf_cd9b85b271fff465 := __obf_931a5a6194f47dee(__obf_edd9fe48d39445e4, __obf_bb1024339a5c1253)
			if __obf_cd9b85b271fff465 == nil {
				return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787[__obf_28d099df85f083a8:])
			}
			__obf_edd9fe48d39445e4.
				ResetBytes(__obf_cd9b85b271fff465)
		case int:
			__obf_cd9b85b271fff465 := __obf_4c9d33490f37e941(__obf_edd9fe48d39445e4, __obf_bb1024339a5c1253)
			if __obf_cd9b85b271fff465 == nil {
				return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787[__obf_28d099df85f083a8:])
			}
			__obf_edd9fe48d39445e4.
				ResetBytes(__obf_cd9b85b271fff465)
		case int32:
			if '*' == __obf_bb1024339a5c1253 {
				return __obf_edd9fe48d39445e4.__obf_1da8db8a1d168ded().Get(__obf_b90e4ca332607787[__obf_28d099df85f083a8:]...)
			}
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787[__obf_28d099df85f083a8:])
		default:
			return __obf_5a6677fd9236ae8d(__obf_b90e4ca332607787[__obf_28d099df85f083a8:])
		}
	}
	if __obf_edd9fe48d39445e4.Error != nil && __obf_edd9fe48d39445e4.Error != io.EOF {
		return &__obf_5183f6b11c973faf{__obf_86279f81acb381f3{}, __obf_edd9fe48d39445e4.Error}
	}
	return __obf_edd9fe48d39445e4.__obf_1da8db8a1d168ded()
}

var __obf_ffaf86ce45837916 = reflect2.TypeOfPtr((*Any)(nil)).Elem()

func __obf_79cadeaf3aef9cd1(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValDecoder {
	if __obf_8667d4fc2a95b0d7 == __obf_ffaf86ce45837916 {
		return &__obf_56fda69ec62d2ad1{}
	}
	if __obf_8667d4fc2a95b0d7.Implements(__obf_ffaf86ce45837916) {
		return &__obf_8b03a02ea10751c7{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7}
	}
	return nil
}

func __obf_257945e5ab511848(__obf_62435d89ebefd5aa *__obf_62435d89ebefd5aa, __obf_8667d4fc2a95b0d7 reflect2.Type) ValEncoder {
	if __obf_8667d4fc2a95b0d7 == __obf_ffaf86ce45837916 {
		return &__obf_56fda69ec62d2ad1{}
	}
	if __obf_8667d4fc2a95b0d7.Implements(__obf_ffaf86ce45837916) {
		return &__obf_8b03a02ea10751c7{__obf_797d4fe23b3556f8: __obf_8667d4fc2a95b0d7}
	}
	return nil
}

type __obf_8b03a02ea10751c7 struct {
	__obf_797d4fe23b3556f8 reflect2.Type
}

func (__obf_1569d098873d2228 *__obf_8b03a02ea10751c7) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	panic("not implemented")
}

func (__obf_1569d098873d2228 *__obf_8b03a02ea10751c7) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_50acae8c0a871ba1 := __obf_1569d098873d2228.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	any := __obf_50acae8c0a871ba1.(Any)
	any.WriteTo(__obf_2361f5214e84c60f)
}

func (__obf_1569d098873d2228 *__obf_8b03a02ea10751c7) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	__obf_50acae8c0a871ba1 := __obf_1569d098873d2228.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	any := __obf_50acae8c0a871ba1.(Any)
	return any.Size() == 0
}

type __obf_56fda69ec62d2ad1 struct {
}

func (__obf_1569d098873d2228 *__obf_56fda69ec62d2ad1) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	*(*Any)(__obf_753ab3fb72cdd659) = __obf_edd9fe48d39445e4.__obf_1da8db8a1d168ded()
}

func (__obf_1569d098873d2228 *__obf_56fda69ec62d2ad1) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	any := *(*Any)(__obf_753ab3fb72cdd659)
	if any == nil {
		__obf_2361f5214e84c60f.
			WriteNil()
		return
	}
	any.WriteTo(__obf_2361f5214e84c60f)
}

func (__obf_1569d098873d2228 *__obf_56fda69ec62d2ad1) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	any := *(*Any)(__obf_753ab3fb72cdd659)
	return any.Size() == 0
}
