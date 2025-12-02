package __obf_c7b79b12b33d8238

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
	ToVal(__obf_35accd096612ebe4 any)
	Get(__obf_5dde9fb4007294d4 ...any) Any
	Size() int
	Keys() []string
	GetInterface() any
	WriteTo(__obf_d8f50bcfe87d8b47 *Stream)
}

type __obf_3a44d9c5204c2e12 struct{}

func (any *__obf_3a44d9c5204c2e12) Get(__obf_5dde9fb4007294d4 ...any) Any {
	return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, fmt.Errorf("GetIndex %v from simple value", __obf_5dde9fb4007294d4)}
}

func (any *__obf_3a44d9c5204c2e12) Size() int {
	return 0
}

func (any *__obf_3a44d9c5204c2e12) Keys() []string {
	return []string{}
}

func (any *__obf_3a44d9c5204c2e12) ToVal(__obf_d6e2df4782353c64 any) {
	panic("not implemented")
}

// WrapInt32 turn int32 into Any interface
func WrapInt32(__obf_35accd096612ebe4 int32) Any {
	return &__obf_67b34d104eef23fb{__obf_3a44d9c5204c2e12{}, __obf_35accd096612ebe4}
}

// WrapInt64 turn int64 into Any interface
func WrapInt64(__obf_35accd096612ebe4 int64) Any {
	return &__obf_064ca4e3fc05241a{__obf_3a44d9c5204c2e12{}, __obf_35accd096612ebe4}
}

// WrapUint32 turn uint32 into Any interface
func WrapUint32(__obf_35accd096612ebe4 uint32) Any {
	return &__obf_e407dbf923c34658{__obf_3a44d9c5204c2e12{}, __obf_35accd096612ebe4}
}

// WrapUint64 turn uint64 into Any interface
func WrapUint64(__obf_35accd096612ebe4 uint64) Any {
	return &__obf_aeafd238348bf351{__obf_3a44d9c5204c2e12{}, __obf_35accd096612ebe4}
}

// WrapFloat64 turn float64 into Any interface
func WrapFloat64(__obf_35accd096612ebe4 float64) Any {
	return &__obf_fc7eb09dc621f051{__obf_3a44d9c5204c2e12{}, __obf_35accd096612ebe4}
}

// WrapString turn string into Any interface
func WrapString(__obf_35accd096612ebe4 string) Any {
	return &__obf_755c14a10251d761{__obf_3a44d9c5204c2e12{}, __obf_35accd096612ebe4}
}

// Wrap turn a go object into Any interface
func Wrap(__obf_35accd096612ebe4 any) Any {
	if __obf_35accd096612ebe4 == nil {
		return &__obf_3296d866fc4daf2b{}
	}
	__obf_e9556f5ffa497986, __obf_d071260ea7e72a00 := __obf_35accd096612ebe4.(Any)
	if __obf_d071260ea7e72a00 {
		return __obf_e9556f5ffa497986
	}
	__obf_edcded11af6ebc4c := reflect2.TypeOf(__obf_35accd096612ebe4)
	switch __obf_edcded11af6ebc4c.Kind() {
	case reflect.Slice:
		return __obf_a040b37b3279e567(__obf_35accd096612ebe4)
	case reflect.Struct:
		return __obf_e7c64fed69971093(__obf_35accd096612ebe4)
	case reflect.Map:
		return __obf_2259c4eb0c25660e(__obf_35accd096612ebe4)
	case reflect.String:
		return WrapString(__obf_35accd096612ebe4.(string))
	case reflect.Int:
		if strconv.IntSize == 32 {
			return WrapInt32(int32(__obf_35accd096612ebe4.(int)))
		}
		return WrapInt64(int64(__obf_35accd096612ebe4.(int)))
	case reflect.Int8:
		return WrapInt32(int32(__obf_35accd096612ebe4.(int8)))
	case reflect.Int16:
		return WrapInt32(int32(__obf_35accd096612ebe4.(int16)))
	case reflect.Int32:
		return WrapInt32(__obf_35accd096612ebe4.(int32))
	case reflect.Int64:
		return WrapInt64(__obf_35accd096612ebe4.(int64))
	case reflect.Uint:
		if strconv.IntSize == 32 {
			return WrapUint32(uint32(__obf_35accd096612ebe4.(uint)))
		}
		return WrapUint64(uint64(__obf_35accd096612ebe4.(uint)))
	case reflect.Uintptr:
		if __obf_95a95e14d0f9c485 == 32 {
			return WrapUint32(uint32(__obf_35accd096612ebe4.(uintptr)))
		}
		return WrapUint64(uint64(__obf_35accd096612ebe4.(uintptr)))
	case reflect.Uint8:
		return WrapUint32(uint32(__obf_35accd096612ebe4.(uint8)))
	case reflect.Uint16:
		return WrapUint32(uint32(__obf_35accd096612ebe4.(uint16)))
	case reflect.Uint32:
		return WrapUint32(uint32(__obf_35accd096612ebe4.(uint32)))
	case reflect.Uint64:
		return WrapUint64(__obf_35accd096612ebe4.(uint64))
	case reflect.Float32:
		return WrapFloat64(float64(__obf_35accd096612ebe4.(float32)))
	case reflect.Float64:
		return WrapFloat64(__obf_35accd096612ebe4.(float64))
	case reflect.Bool:
		if __obf_35accd096612ebe4.(bool) == true {
			return &__obf_2709e0e538cbe0e2{}
		}
		return &__obf_6fd7e945c46568e9{}
	}
	return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, fmt.Errorf("unsupported type: %v", __obf_edcded11af6ebc4c)}
}

// ReadAny read next JSON element as an Any object. It is a better json.RawMessage.
func (__obf_0da8c843dad13139 *Iterator) ReadAny() Any {
	return __obf_0da8c843dad13139.__obf_772f435663149392()
}

func (__obf_0da8c843dad13139 *Iterator) __obf_772f435663149392() Any {
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	switch __obf_12577c03a12f0d2c {
	case '"':
		__obf_0da8c843dad13139.__obf_a675366c80290b83()
		return &__obf_755c14a10251d761{__obf_3a44d9c5204c2e12{}, __obf_0da8c843dad13139.ReadString()}
	case 'n':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('u', 'l', 'l') // null
		return &__obf_3296d866fc4daf2b{}
	case 't':
		__obf_0da8c843dad13139.__obf_5da8d54e7a1e782c('r', 'u', 'e') // true
		return &__obf_2709e0e538cbe0e2{}
	case 'f':
		__obf_0da8c843dad13139.__obf_62908d9424a8486b('a', 'l', 's', 'e') // false
		return &__obf_6fd7e945c46568e9{}
	case '{':
		return __obf_0da8c843dad13139.__obf_8873a5b089cd3d82()
	case '[':
		return __obf_0da8c843dad13139.__obf_13ba3ce8b43586ca()
	case '-':
		return __obf_0da8c843dad13139.__obf_c4f1f40458187665(false)
	case 0:
		return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, errors.New("input is empty")}
	default:
		return __obf_0da8c843dad13139.__obf_c4f1f40458187665(true)
	}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_c4f1f40458187665(__obf_3da6f88be8d72d88 bool) Any {
	__obf_0da8c843dad13139.__obf_611b08915abc95f3(__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 - 1)
	__obf_0da8c843dad13139.__obf_9dd6f9cece1dabe0()
	__obf_daf67e76d85bc55f := __obf_0da8c843dad13139.__obf_d96514238c6d6bd1()
	return &__obf_7ab06fd17433b644{__obf_3a44d9c5204c2e12{}, __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71, __obf_daf67e76d85bc55f, nil}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_8873a5b089cd3d82() Any {
	__obf_0da8c843dad13139.__obf_611b08915abc95f3(__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 - 1)
	__obf_0da8c843dad13139.__obf_f42b686c47b3acb5()
	__obf_daf67e76d85bc55f := __obf_0da8c843dad13139.__obf_d96514238c6d6bd1()
	return &__obf_293b886b2fccdeac{__obf_3a44d9c5204c2e12{}, __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71, __obf_daf67e76d85bc55f, nil}
}

func (__obf_0da8c843dad13139 *Iterator) __obf_13ba3ce8b43586ca() Any {
	__obf_0da8c843dad13139.__obf_611b08915abc95f3(__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 - 1)
	__obf_0da8c843dad13139.__obf_a12f980c2c849354()
	__obf_daf67e76d85bc55f := __obf_0da8c843dad13139.__obf_d96514238c6d6bd1()
	return &__obf_94d970ccdadcfb91{__obf_3a44d9c5204c2e12{}, __obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71, __obf_daf67e76d85bc55f, nil}
}

func __obf_17f2269292836639(__obf_0da8c843dad13139 *Iterator, __obf_91175e052ee012bd string) []byte {
	var __obf_22405d874667f998 []byte
	__obf_0da8c843dad13139.
		ReadObjectCB(func(__obf_0da8c843dad13139 *Iterator, __obf_ad34f8a6de2b01cb string) bool {
			if __obf_ad34f8a6de2b01cb == __obf_91175e052ee012bd {
				__obf_22405d874667f998 = __obf_0da8c843dad13139.SkipAndReturnBytes()
				return false
			}
			__obf_0da8c843dad13139.
				Skip()
			return true
		})
	return __obf_22405d874667f998
}

func __obf_7a195008f4d0d118(__obf_0da8c843dad13139 *Iterator, __obf_91175e052ee012bd int) []byte {
	var __obf_22405d874667f998 []byte
	__obf_ff4330e73e137d5c := 0
	__obf_0da8c843dad13139.
		ReadArrayCB(func(__obf_0da8c843dad13139 *Iterator) bool {
			if __obf_ff4330e73e137d5c == __obf_91175e052ee012bd {
				__obf_22405d874667f998 = __obf_0da8c843dad13139.SkipAndReturnBytes()
				return false
			}
			__obf_0da8c843dad13139.
				Skip()
			__obf_ff4330e73e137d5c++
			return true
		})
	return __obf_22405d874667f998
}

func __obf_fedb6216061761ae(__obf_0da8c843dad13139 *Iterator, __obf_5dde9fb4007294d4 []any) Any {
	for __obf_a051269af3a541bb, __obf_d2dcac1798fccd0f := range __obf_5dde9fb4007294d4 {
		switch __obf_31b942107db5d323 := __obf_d2dcac1798fccd0f.(type) {
		case string:
			__obf_39f3961bda254565 := __obf_17f2269292836639(__obf_0da8c843dad13139, __obf_31b942107db5d323)
			if __obf_39f3961bda254565 == nil {
				return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4[__obf_a051269af3a541bb:])
			}
			__obf_0da8c843dad13139.
				ResetBytes(__obf_39f3961bda254565)
		case int:
			__obf_39f3961bda254565 := __obf_7a195008f4d0d118(__obf_0da8c843dad13139, __obf_31b942107db5d323)
			if __obf_39f3961bda254565 == nil {
				return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4[__obf_a051269af3a541bb:])
			}
			__obf_0da8c843dad13139.
				ResetBytes(__obf_39f3961bda254565)
		case int32:
			if '*' == __obf_31b942107db5d323 {
				return __obf_0da8c843dad13139.__obf_772f435663149392().Get(__obf_5dde9fb4007294d4[__obf_a051269af3a541bb:]...)
			}
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4[__obf_a051269af3a541bb:])
		default:
			return __obf_5b62b1b9d47ae70b(__obf_5dde9fb4007294d4[__obf_a051269af3a541bb:])
		}
	}
	if __obf_0da8c843dad13139.Error != nil && __obf_0da8c843dad13139.Error != io.EOF {
		return &__obf_cb30187c64f44cee{__obf_3a44d9c5204c2e12{}, __obf_0da8c843dad13139.Error}
	}
	return __obf_0da8c843dad13139.__obf_772f435663149392()
}

var __obf_50a9ff9d46f983fb = reflect2.TypeOfPtr((*Any)(nil)).Elem()

func __obf_1a6dc69a4bb1ccaa(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	if __obf_edcded11af6ebc4c == __obf_50a9ff9d46f983fb {
		return &__obf_1e13e14b8f5bf867{}
	}
	if __obf_edcded11af6ebc4c.Implements(__obf_50a9ff9d46f983fb) {
		return &__obf_5f1f8838c162879d{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c}
	}
	return nil
}

func __obf_eab6c90d596283e9(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	if __obf_edcded11af6ebc4c == __obf_50a9ff9d46f983fb {
		return &__obf_1e13e14b8f5bf867{}
	}
	if __obf_edcded11af6ebc4c.Implements(__obf_50a9ff9d46f983fb) {
		return &__obf_5f1f8838c162879d{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c}
	}
	return nil
}

type __obf_5f1f8838c162879d struct {
	__obf_e0ba73c3f13f4455 reflect2.Type
}

func (__obf_26e7ab1d7cef55d9 *__obf_5f1f8838c162879d) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	panic("not implemented")
}

func (__obf_26e7ab1d7cef55d9 *__obf_5f1f8838c162879d) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d6e2df4782353c64 := __obf_26e7ab1d7cef55d9.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	any := __obf_d6e2df4782353c64.(Any)
	any.WriteTo(__obf_d8f50bcfe87d8b47)
}

func (__obf_26e7ab1d7cef55d9 *__obf_5f1f8838c162879d) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_d6e2df4782353c64 := __obf_26e7ab1d7cef55d9.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	any := __obf_d6e2df4782353c64.(Any)
	return any.Size() == 0
}

type __obf_1e13e14b8f5bf867 struct {
}

func (__obf_26e7ab1d7cef55d9 *__obf_1e13e14b8f5bf867) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	*(*Any)(__obf_575c04f2b9d91315) = __obf_0da8c843dad13139.__obf_772f435663149392()
}

func (__obf_26e7ab1d7cef55d9 *__obf_1e13e14b8f5bf867) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	any := *(*Any)(__obf_575c04f2b9d91315)
	if any == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	any.WriteTo(__obf_d8f50bcfe87d8b47)
}

func (__obf_26e7ab1d7cef55d9 *__obf_1e13e14b8f5bf867) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	any := *(*Any)(__obf_575c04f2b9d91315)
	return any.Size() == 0
}
