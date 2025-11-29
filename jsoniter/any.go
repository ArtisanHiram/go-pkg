package __obf_5b802ce8d9ba56d6

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
	ToVal(__obf_5406b11e33b9d1d3 any)
	Get(__obf_c5d71353f4393b3c ...any) Any
	Size() int
	Keys() []string
	GetInterface() any
	WriteTo(__obf_00fc491caa967a74 *Stream)
}

type __obf_fb75d4e4562132ae struct{}

func (any *__obf_fb75d4e4562132ae) Get(__obf_c5d71353f4393b3c ...any) Any {
	return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, fmt.Errorf("GetIndex %v from simple value", __obf_c5d71353f4393b3c)}
}

func (any *__obf_fb75d4e4562132ae) Size() int {
	return 0
}

func (any *__obf_fb75d4e4562132ae) Keys() []string {
	return []string{}
}

func (any *__obf_fb75d4e4562132ae) ToVal(__obf_f9b1526531f3c024 any) {
	panic("not implemented")
}

// WrapInt32 turn int32 into Any interface
func WrapInt32(__obf_5406b11e33b9d1d3 int32) Any {
	return &__obf_3b63a102cb2fa2bb{__obf_fb75d4e4562132ae{}, __obf_5406b11e33b9d1d3}
}

// WrapInt64 turn int64 into Any interface
func WrapInt64(__obf_5406b11e33b9d1d3 int64) Any {
	return &__obf_7d29c083a07482f6{__obf_fb75d4e4562132ae{}, __obf_5406b11e33b9d1d3}
}

// WrapUint32 turn uint32 into Any interface
func WrapUint32(__obf_5406b11e33b9d1d3 uint32) Any {
	return &__obf_dbd8d5f30edec5df{__obf_fb75d4e4562132ae{}, __obf_5406b11e33b9d1d3}
}

// WrapUint64 turn uint64 into Any interface
func WrapUint64(__obf_5406b11e33b9d1d3 uint64) Any {
	return &__obf_1eb46b63f55e7ee7{__obf_fb75d4e4562132ae{}, __obf_5406b11e33b9d1d3}
}

// WrapFloat64 turn float64 into Any interface
func WrapFloat64(__obf_5406b11e33b9d1d3 float64) Any {
	return &__obf_3073bf0ff9855ed4{__obf_fb75d4e4562132ae{}, __obf_5406b11e33b9d1d3}
}

// WrapString turn string into Any interface
func WrapString(__obf_5406b11e33b9d1d3 string) Any {
	return &__obf_6ff1a29b3162ef37{__obf_fb75d4e4562132ae{}, __obf_5406b11e33b9d1d3}
}

// Wrap turn a go object into Any interface
func Wrap(__obf_5406b11e33b9d1d3 any) Any {
	if __obf_5406b11e33b9d1d3 == nil {
		return &__obf_d46f328cec1c970b{}
	}
	__obf_fd169c69359ba9fe, __obf_050c6c300d7642b6 := __obf_5406b11e33b9d1d3.(Any)
	if __obf_050c6c300d7642b6 {
		return __obf_fd169c69359ba9fe
	}
	__obf_5efc66d8c338c133 := reflect2.TypeOf(__obf_5406b11e33b9d1d3)
	switch __obf_5efc66d8c338c133.Kind() {
	case reflect.Slice:
		return __obf_a6d72607e427694a(__obf_5406b11e33b9d1d3)
	case reflect.Struct:
		return __obf_4b9ae4bb2bbc6708(__obf_5406b11e33b9d1d3)
	case reflect.Map:
		return __obf_03d594f9555cce37(__obf_5406b11e33b9d1d3)
	case reflect.String:
		return WrapString(__obf_5406b11e33b9d1d3.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return WrapInt32(int32(__obf_5406b11e33b9d1d3.(int)))
		}
		return WrapInt64(int64(__obf_5406b11e33b9d1d3.(int)))
	case reflect.Int8:
		return WrapInt32(int32(__obf_5406b11e33b9d1d3.(int8)))
	case reflect.Int16:
		return WrapInt32(int32(__obf_5406b11e33b9d1d3.(int16)))
	case reflect.Int32:
		return WrapInt32(__obf_5406b11e33b9d1d3.(int32))
	case reflect.Int64:
		return WrapInt64(__obf_5406b11e33b9d1d3.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return WrapUint32(uint32(__obf_5406b11e33b9d1d3.(uint)))
		}
		return WrapUint64(uint64(__obf_5406b11e33b9d1d3.(uint)))
	case reflect.Uintptr:
		if __obf_1b422ac0cf965f7a == 32 {
			return WrapUint32(uint32(__obf_5406b11e33b9d1d3.(uintptr)))
		}
		return WrapUint64(uint64(__obf_5406b11e33b9d1d3.(uintptr)))
	case reflect.Uint8:
		return WrapUint32(uint32(__obf_5406b11e33b9d1d3.(uint8)))
	case reflect.Uint16:
		return WrapUint32(uint32(__obf_5406b11e33b9d1d3.(uint16)))
	case reflect.Uint32:
		return WrapUint32(uint32(__obf_5406b11e33b9d1d3.(uint32)))
	case reflect.Uint64:
		return WrapUint64(__obf_5406b11e33b9d1d3.(uint64))
	case reflect.Float32:
		return WrapFloat64(float64(__obf_5406b11e33b9d1d3.(float32)))
	case reflect.Float64:
		return WrapFloat64(__obf_5406b11e33b9d1d3.(float64))
	case reflect.Bool:
		if __obf_5406b11e33b9d1d3.(bool) == true {
			return &__obf_09e7a4adb16a3462{}
		}
		return &__obf_7a5bd240952f3b1b{}
	}
	return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, fmt.Errorf("unsupported type: %v", __obf_5efc66d8c338c133)}
}

// ReadAny read next JSON element as an Any object. It is a better json.RawMessage.
func (__obf_67008a6a9e5ba828 *Iterator) ReadAny() Any {
	return __obf_67008a6a9e5ba828.__obf_f0a5bf7546ffe762()
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_f0a5bf7546ffe762() Any {
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	switch __obf_dab9baaadfa7c8c2 {
	case '"':
		__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
		return &__obf_6ff1a29b3162ef37{__obf_fb75d4e4562132ae{}, __obf_67008a6a9e5ba828.ReadString()}
	case 'n':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('u', 'l', 'l') // null
		return &__obf_d46f328cec1c970b{}
	case 't':
		__obf_67008a6a9e5ba828.__obf_4aeb767e0be7277a('r', 'u', 'e') // true
		return &__obf_09e7a4adb16a3462{}
	case 'f':
		__obf_67008a6a9e5ba828.__obf_acc74c95f4492ff8('a', 'l', 's', 'e') // false
		return &__obf_7a5bd240952f3b1b{}
	case '{':
		return __obf_67008a6a9e5ba828.__obf_c733fe33063da150()
	case '[':
		return __obf_67008a6a9e5ba828.__obf_6f5c3c1fec473554()
	case '-':
		return __obf_67008a6a9e5ba828.__obf_206ffb25cffa8fd5(false)
	case 0:
		return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, errors.New("input is empty")}
	default:
		return __obf_67008a6a9e5ba828.__obf_206ffb25cffa8fd5(true)
	}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_206ffb25cffa8fd5(__obf_0c5a5ed463269bab bool) Any {
	__obf_67008a6a9e5ba828.__obf_1b03992be8504c10(__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 - 1)
	__obf_67008a6a9e5ba828.__obf_a64603a02c650185()
	__obf_8d30cc3198671840 := __obf_67008a6a9e5ba828.__obf_3998584530a209bb()
	return &__obf_dff8f3718df12fe4{__obf_fb75d4e4562132ae{}, __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392, __obf_8d30cc3198671840, nil}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_c733fe33063da150() Any {
	__obf_67008a6a9e5ba828.__obf_1b03992be8504c10(__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 - 1)
	__obf_67008a6a9e5ba828.__obf_7e3d8ea4ab2a8258()
	__obf_8d30cc3198671840 := __obf_67008a6a9e5ba828.__obf_3998584530a209bb()
	return &__obf_92721e3ae32bb50a{__obf_fb75d4e4562132ae{}, __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392, __obf_8d30cc3198671840, nil}
}

func (__obf_67008a6a9e5ba828 *Iterator) __obf_6f5c3c1fec473554() Any {
	__obf_67008a6a9e5ba828.__obf_1b03992be8504c10(__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 - 1)
	__obf_67008a6a9e5ba828.__obf_1d506622058b77ef()
	__obf_8d30cc3198671840 := __obf_67008a6a9e5ba828.__obf_3998584530a209bb()
	return &__obf_41614ecb47fbcff9{__obf_fb75d4e4562132ae{}, __obf_67008a6a9e5ba828.__obf_dca106e1cdf85392, __obf_8d30cc3198671840, nil}
}

func __obf_51a7c2c2318325e4(__obf_67008a6a9e5ba828 *Iterator, __obf_6175131f514bcffc string) []byte {
	var __obf_98a3eefc4c187186 []byte
	__obf_67008a6a9e5ba828.
		ReadObjectCB(func(__obf_67008a6a9e5ba828 *Iterator, __obf_61998fb083387c3c string) bool {
			if __obf_61998fb083387c3c == __obf_6175131f514bcffc {
				__obf_98a3eefc4c187186 = __obf_67008a6a9e5ba828.SkipAndReturnBytes()
				return false
			}
			__obf_67008a6a9e5ba828.
				Skip()
			return true
		})
	return __obf_98a3eefc4c187186
}

func __obf_e8d51e45c7dc6fc5(__obf_67008a6a9e5ba828 *Iterator, __obf_6175131f514bcffc int) []byte {
	var __obf_98a3eefc4c187186 []byte
	__obf_2c0ce853cb81f014 := 0
	__obf_67008a6a9e5ba828.
		ReadArrayCB(func(__obf_67008a6a9e5ba828 *Iterator) bool {
			if __obf_2c0ce853cb81f014 == __obf_6175131f514bcffc {
				__obf_98a3eefc4c187186 = __obf_67008a6a9e5ba828.SkipAndReturnBytes()
				return false
			}
			__obf_67008a6a9e5ba828.
				Skip()
			__obf_2c0ce853cb81f014++
			return true
		})
	return __obf_98a3eefc4c187186
}

func __obf_9c645421e8d71b4a(__obf_67008a6a9e5ba828 *Iterator, __obf_c5d71353f4393b3c []any) Any {
	for __obf_2deec7c38ffb6ae3, __obf_d9a277b87563d6fc := range __obf_c5d71353f4393b3c {
		switch __obf_20bde2be80d52297 := __obf_d9a277b87563d6fc.(type) {
		case string:
			__obf_611f49298438d349 := __obf_51a7c2c2318325e4(__obf_67008a6a9e5ba828, __obf_20bde2be80d52297)
			if __obf_611f49298438d349 == nil {
				return __obf_e990980207770dd2(__obf_c5d71353f4393b3c[__obf_2deec7c38ffb6ae3:])
			}
			__obf_67008a6a9e5ba828.
				ResetBytes(__obf_611f49298438d349)
		case int:
			__obf_611f49298438d349 := __obf_e8d51e45c7dc6fc5(__obf_67008a6a9e5ba828, __obf_20bde2be80d52297)
			if __obf_611f49298438d349 == nil {
				return __obf_e990980207770dd2(__obf_c5d71353f4393b3c[__obf_2deec7c38ffb6ae3:])
			}
			__obf_67008a6a9e5ba828.
				ResetBytes(__obf_611f49298438d349)
		case int32:
			if '*' == __obf_20bde2be80d52297 {
				return __obf_67008a6a9e5ba828.__obf_f0a5bf7546ffe762().Get(__obf_c5d71353f4393b3c[__obf_2deec7c38ffb6ae3:]...)
			}
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c[__obf_2deec7c38ffb6ae3:])
		default:
			return __obf_e990980207770dd2(__obf_c5d71353f4393b3c[__obf_2deec7c38ffb6ae3:])
		}
	}
	if __obf_67008a6a9e5ba828.Error != nil && __obf_67008a6a9e5ba828.Error != io.EOF {
		return &__obf_00698f76feb93563{__obf_fb75d4e4562132ae{}, __obf_67008a6a9e5ba828.Error}
	}
	return __obf_67008a6a9e5ba828.__obf_f0a5bf7546ffe762()
}

var __obf_f4cd8d8ddfa6a4ed = reflect2.TypeOfPtr((*Any)(nil)).Elem()

func __obf_d7f8199541699318(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	if __obf_5efc66d8c338c133 == __obf_f4cd8d8ddfa6a4ed {
		return &__obf_8f6563971e7e9e59{}
	}
	if __obf_5efc66d8c338c133.Implements(__obf_f4cd8d8ddfa6a4ed) {
		return &__obf_e4e7a000eb35c64d{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133}
	}
	return nil
}

func __obf_3f8d707915d65625(__obf_08da24be66d0d13c *__obf_08da24be66d0d13c, __obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	if __obf_5efc66d8c338c133 == __obf_f4cd8d8ddfa6a4ed {
		return &__obf_8f6563971e7e9e59{}
	}
	if __obf_5efc66d8c338c133.Implements(__obf_f4cd8d8ddfa6a4ed) {
		return &__obf_e4e7a000eb35c64d{__obf_5e777ac7034ab220: __obf_5efc66d8c338c133}
	}
	return nil
}

type __obf_e4e7a000eb35c64d struct {
	__obf_5e777ac7034ab220 reflect2.Type
}

func (__obf_0099bc81efe356aa *__obf_e4e7a000eb35c64d) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	panic("not implemented")
}

func (__obf_0099bc81efe356aa *__obf_e4e7a000eb35c64d) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	__obf_f9b1526531f3c024 := __obf_0099bc81efe356aa.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	any := __obf_f9b1526531f3c024.(Any)
	any.WriteTo(__obf_00fc491caa967a74)
}

func (__obf_0099bc81efe356aa *__obf_e4e7a000eb35c64d) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	__obf_f9b1526531f3c024 := __obf_0099bc81efe356aa.__obf_5e777ac7034ab220.UnsafeIndirect(__obf_d3c919547bf7e47a)
	any := __obf_f9b1526531f3c024.(Any)
	return any.Size() == 0
}

type __obf_8f6563971e7e9e59 struct {
}

func (__obf_0099bc81efe356aa *__obf_8f6563971e7e9e59) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	*(*Any)(__obf_d3c919547bf7e47a) = __obf_67008a6a9e5ba828.__obf_f0a5bf7546ffe762()
}

func (__obf_0099bc81efe356aa *__obf_8f6563971e7e9e59) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	any := *(*Any)(__obf_d3c919547bf7e47a)
	if any == nil {
		__obf_00fc491caa967a74.
			WriteNil()
		return
	}
	any.WriteTo(__obf_00fc491caa967a74)
}

func (__obf_0099bc81efe356aa *__obf_8f6563971e7e9e59) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	any := *(*Any)(__obf_d3c919547bf7e47a)
	return any.Size() == 0
}
