package __obf_de86cdc8ae98b45b

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_d570f482b4bbddc6 = 1 << 20
	__obf_bb4b009fa46ad65f = // 1mb
	1e6
	__obf_5e8134fa61b47535 = // 1m elements
	1e6                      // 1m elements
)

const (
	__obf_fa610f2bf51f21de uint32 = 1 << iota
	__obf_64c53d062f7fdfd5
	__obf_079ef68f50d79bcd
	__obf_c646a26ee823abc3
)

type __obf_3ea42149e533ba5d interface {
	io.Reader
	io.ByteScanner
}

//------------------------------------------------------------------------------

var __obf_ede081eee0bedbe1 = sync.Pool{
	New: func() any {
		return NewDecoder(nil)
	},
}

func GetDecoder() *Decoder {
	return __obf_ede081eee0bedbe1.Get().(*Decoder)
}

func PutDecoder(__obf_9461956859459a44 *Decoder) {
	__obf_9461956859459a44.__obf_4609cadeaf50917c = nil
	__obf_9461956859459a44.__obf_a93d004caac96500 = nil
	__obf_ede081eee0bedbe1.
		Put(__obf_9461956859459a44)
}

//------------------------------------------------------------------------------

// Unmarshal decodes the MessagePack-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(__obf_50d002bc8b306017 []byte, __obf_17732590722140e7 any) error {
	__obf_9461956859459a44 := GetDecoder()
	__obf_9461956859459a44.
		UsePreallocateValues(true)
	__obf_9461956859459a44.
		Reset(bytes.NewReader(__obf_50d002bc8b306017))
	__obf_0feb3528b7b709ec := __obf_9461956859459a44.Decode(__obf_17732590722140e7)

	PutDecoder(__obf_9461956859459a44)

	return __obf_0feb3528b7b709ec
}

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	__obf_4609cadeaf50917c io.Reader
	__obf_a93d004caac96500 io.ByteScanner
	__obf_c1135a4c578558d5 func(*Decoder) (any, error)
	__obf_603dff1ad8d49539 string
	__obf_9c8c6a5e4cc8822c []byte
	__obf_fabeb54cc5377fbf []byte
	__obf_25128eea0d506b65 []string
	__obf_a15a642720e39c3e uint32
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r
// beyond the requested msgpack values. Buffering can be disabled
// by passing a reader that implements io.ByteScanner interface.
func NewDecoder(__obf_4609cadeaf50917c io.Reader) *Decoder {
	__obf_dcaad1165bb07af9 := new(Decoder)
	__obf_dcaad1165bb07af9.
		Reset(__obf_4609cadeaf50917c)
	return __obf_dcaad1165bb07af9
}

// Reset discards any buffered data, resets all state, and switches the buffered
// reader to read from r.
func (__obf_dcaad1165bb07af9 *Decoder) Reset(__obf_4609cadeaf50917c io.Reader) {
	__obf_dcaad1165bb07af9.
		ResetDict(__obf_4609cadeaf50917c, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_dcaad1165bb07af9 *Decoder) ResetDict(__obf_4609cadeaf50917c io.Reader, __obf_25128eea0d506b65 []string) {
	__obf_dcaad1165bb07af9.
		ResetReader(__obf_4609cadeaf50917c)
	__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e = 0
	__obf_dcaad1165bb07af9.__obf_603dff1ad8d49539 = ""
	__obf_dcaad1165bb07af9.__obf_25128eea0d506b65 = __obf_25128eea0d506b65
}

func (__obf_dcaad1165bb07af9 *Decoder) WithDict(__obf_25128eea0d506b65 []string, __obf_69331012fc3414ab func(*Decoder) error) error {
	__obf_3ae10cdb339e5623 := __obf_dcaad1165bb07af9.__obf_25128eea0d506b65
	__obf_dcaad1165bb07af9.__obf_25128eea0d506b65 = __obf_25128eea0d506b65
	__obf_0feb3528b7b709ec := __obf_69331012fc3414ab(__obf_dcaad1165bb07af9)
	__obf_dcaad1165bb07af9.__obf_25128eea0d506b65 = __obf_3ae10cdb339e5623
	return __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) ResetReader(__obf_4609cadeaf50917c io.Reader) {
	__obf_dcaad1165bb07af9.__obf_c1135a4c578558d5 = nil
	__obf_dcaad1165bb07af9.__obf_25128eea0d506b65 = nil

	if __obf_037c1bd9b013e3f5, __obf_77cfff95beb3cc99 := __obf_4609cadeaf50917c.(__obf_3ea42149e533ba5d); __obf_77cfff95beb3cc99 {
		__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c = __obf_037c1bd9b013e3f5
		__obf_dcaad1165bb07af9.__obf_a93d004caac96500 = __obf_037c1bd9b013e3f5
	} else if __obf_4609cadeaf50917c == nil {
		__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c = nil
		__obf_dcaad1165bb07af9.__obf_a93d004caac96500 = nil
	} else {
		__obf_037c1bd9b013e3f5 := bufio.NewReader(__obf_4609cadeaf50917c)
		__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c = __obf_037c1bd9b013e3f5
		__obf_dcaad1165bb07af9.__obf_a93d004caac96500 = __obf_037c1bd9b013e3f5
	}
}

func (__obf_dcaad1165bb07af9 *Decoder) SetMapDecoder(__obf_69331012fc3414ab func(*Decoder) (any, error)) {
	__obf_dcaad1165bb07af9.__obf_c1135a4c578558d5 = __obf_69331012fc3414ab
}

// UseLooseInterfaceDecoding causes decoder to use DecodeInterfaceLoose
// to decode msgpack value into Go some.
func (__obf_dcaad1165bb07af9 *Decoder) UseLooseInterfaceDecoding(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e |= __obf_fa610f2bf51f21de
	} else {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e &= ^__obf_fa610f2bf51f21de
	}
}

// SetCustomStructTag causes the decoder to use the supplied tag as a fallback option
// if there is no msgpack tag.
func (__obf_dcaad1165bb07af9 *Decoder) SetCustomStructTag(__obf_ee8ab0888686da67 string) {
	__obf_dcaad1165bb07af9.__obf_603dff1ad8d49539 = __obf_ee8ab0888686da67
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (__obf_dcaad1165bb07af9 *Decoder) DisallowUnknownFields(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e |= __obf_64c53d062f7fdfd5
	} else {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e &= ^__obf_64c53d062f7fdfd5
	}
}

// UseInternedStrings enables support for decoding interned strings.
func (__obf_dcaad1165bb07af9 *Decoder) UseInternedStrings(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e |= __obf_1fcd090a6da01656
	} else {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e &= ^__obf_1fcd090a6da01656
	}
}

// UsePreallocateValues enables preallocating values in chunks
func (__obf_dcaad1165bb07af9 *Decoder) UsePreallocateValues(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e |= __obf_079ef68f50d79bcd
	} else {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e &= ^__obf_079ef68f50d79bcd
	}
}

// DisableAllocLimit enables fully allocating slices/maps when the size is known
func (__obf_dcaad1165bb07af9 *Decoder) DisableAllocLimit(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e |= __obf_c646a26ee823abc3
	} else {
		__obf_dcaad1165bb07af9.__obf_a15a642720e39c3e &= ^__obf_c646a26ee823abc3
	}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (__obf_dcaad1165bb07af9 *Decoder) Buffered() io.Reader {
	return __obf_dcaad1165bb07af9.

		//nolint:gocyclo
		__obf_4609cadeaf50917c
}

func (__obf_dcaad1165bb07af9 *Decoder) Decode(__obf_17732590722140e7 any) error {
	var __obf_0feb3528b7b709ec error
	switch __obf_17732590722140e7 := __obf_17732590722140e7.(type) {
	case *string:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeString()
			return __obf_0feb3528b7b709ec
		}
	case *[]byte:
		if __obf_17732590722140e7 != nil {
			return __obf_dcaad1165bb07af9.__obf_a7cf3730c2f8bbd8(__obf_17732590722140e7)
		}
	case *int:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeInt()
			return __obf_0feb3528b7b709ec
		}
	case *int8:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeInt8()
			return __obf_0feb3528b7b709ec
		}
	case *int16:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeInt16()
			return __obf_0feb3528b7b709ec
		}
	case *int32:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeInt32()
			return __obf_0feb3528b7b709ec
		}
	case *int64:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeInt64()
			return __obf_0feb3528b7b709ec
		}
	case *uint:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeUint()
			return __obf_0feb3528b7b709ec
		}
	case *uint8:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeUint8()
			return __obf_0feb3528b7b709ec
		}
	case *uint16:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeUint16()
			return __obf_0feb3528b7b709ec
		}
	case *uint32:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeUint32()
			return __obf_0feb3528b7b709ec
		}
	case *uint64:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeUint64()
			return __obf_0feb3528b7b709ec
		}
	case *bool:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeBool()
			return __obf_0feb3528b7b709ec
		}
	case *float32:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeFloat32()
			return __obf_0feb3528b7b709ec
		}
	case *float64:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeFloat64()
			return __obf_0feb3528b7b709ec
		}
	case *[]string:
		return __obf_dcaad1165bb07af9.__obf_6a43c6fd68849758(__obf_17732590722140e7)
	case *map[string]string:
		return __obf_dcaad1165bb07af9.__obf_234bb66c0ec0a03d(__obf_17732590722140e7)
	case *map[string]any:
		return __obf_dcaad1165bb07af9.__obf_0886d84e85a12f76(__obf_17732590722140e7)
	case *time.Duration:
		if __obf_17732590722140e7 != nil {
			__obf_7c0d6551e0939afa, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
			*__obf_17732590722140e7 = time.Duration(__obf_7c0d6551e0939afa)
			return __obf_0feb3528b7b709ec
		}
	case *time.Time:
		if __obf_17732590722140e7 != nil {
			*__obf_17732590722140e7, __obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.DecodeTime()
			return __obf_0feb3528b7b709ec
		}
	}
	__obf_7c0d6551e0939afa := reflect.ValueOf(__obf_17732590722140e7)
	if !__obf_7c0d6551e0939afa.IsValid() {
		return errors.New("msgpack: Decode(nil)")
	}
	if __obf_7c0d6551e0939afa.Kind() != reflect.Ptr {
		return fmt.Errorf("msgpack: Decode(non-pointer %T)", __obf_17732590722140e7)
	}
	if __obf_7c0d6551e0939afa.IsNil() {
		return fmt.Errorf("msgpack: Decode(non-settable %T)", __obf_17732590722140e7)
	}
	__obf_7c0d6551e0939afa = __obf_7c0d6551e0939afa.Elem()
	if __obf_7c0d6551e0939afa.Kind() == reflect.Interface {
		if !__obf_7c0d6551e0939afa.IsNil() {
			__obf_7c0d6551e0939afa = __obf_7c0d6551e0939afa.Elem()
			if __obf_7c0d6551e0939afa.Kind() != reflect.Ptr {
				return fmt.Errorf("msgpack: Decode(non-pointer %s)", __obf_7c0d6551e0939afa.Type().String())
			}
		}
	}

	return __obf_dcaad1165bb07af9.DecodeValue(__obf_7c0d6551e0939afa)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeMulti(__obf_17732590722140e7 ...any) error {
	for _, __obf_7c0d6551e0939afa := range __obf_17732590722140e7 {
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Decode(__obf_7c0d6551e0939afa); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_7d3129f334945713() (any, error) {
	if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_fa610f2bf51f21de != 0 {
		return __obf_dcaad1165bb07af9.DecodeInterfaceLoose()
	}
	return __obf_dcaad1165bb07af9.DecodeInterface()
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeValue(__obf_17732590722140e7 reflect.Value) error {
	__obf_3a1865a66e130d7b := __obf_25e3584cea7b6666(__obf_17732590722140e7.Type())
	return __obf_3a1865a66e130d7b(__obf_dcaad1165bb07af9, __obf_17732590722140e7)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeNil() error {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_ecf6d2d3253a058d != Nil {
		return fmt.Errorf("msgpack: invalid code=%x decoding nil", __obf_ecf6d2d3253a058d)
	}
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_0913827f98642042(__obf_17732590722140e7 reflect.Value) error {
	__obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeNil()
	if __obf_17732590722140e7.IsNil() {
		return __obf_0feb3528b7b709ec
	}
	if __obf_17732590722140e7.Kind() == reflect.Ptr {
		__obf_17732590722140e7 = __obf_17732590722140e7.Elem()
	}
	__obf_17732590722140e7.
		Set(reflect.Zero(__obf_17732590722140e7.Type()))
	return __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeBool() (bool, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return false, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.bool(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) bool(__obf_ecf6d2d3253a058d byte) (bool, error) {
	if __obf_ecf6d2d3253a058d == Nil {
		return false, nil
	}
	if __obf_ecf6d2d3253a058d == False {
		return false, nil
	}
	if __obf_ecf6d2d3253a058d == True {
		return true, nil
	}
	return false, fmt.Errorf("msgpack: invalid code=%x decoding bool", __obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeDuration() (time.Duration, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeInt64()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return time.Duration(__obf_2b0247e819b1bf4a), nil
}

// DecodeInterface decodes value into interface. It returns following types:
//   - nil,
//   - bool,
//   - int8, int16, int32, int64,
//   - uint8, uint16, uint32, uint64,
//   - float32 and float64,
//   - string,
//   - []byte,
//   - slices of any of the above,
//   - maps of any of the above.
//
// DecodeInterface should be used only when you don't know the type of value
// you are decoding. For example, if you are decoding number it is better to use
// DecodeInt64 for negative numbers and DecodeUint64 for positive numbers.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeInterface() (any, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}

	if IsFixedNum(__obf_ecf6d2d3253a058d) {
		return int8(__obf_ecf6d2d3253a058d), nil
	}
	if IsFixedMap(__obf_ecf6d2d3253a058d) {
		__obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_a93d004caac96500.UnreadByte()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_8468f219671f8d5a()
	}
	if IsFixedArray(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.__obf_9d79cf59b9daca03(__obf_ecf6d2d3253a058d)
	}
	if IsFixedString(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.string(__obf_ecf6d2d3253a058d)
	}

	switch __obf_ecf6d2d3253a058d {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_dcaad1165bb07af9.bool(__obf_ecf6d2d3253a058d)
	case Float:
		return __obf_dcaad1165bb07af9.float32(__obf_ecf6d2d3253a058d)
	case Double:
		return __obf_dcaad1165bb07af9.float64(__obf_ecf6d2d3253a058d)
	case Uint8:
		return __obf_dcaad1165bb07af9.uint8()
	case Uint16:
		return __obf_dcaad1165bb07af9.uint16()
	case Uint32:
		return __obf_dcaad1165bb07af9.uint32()
	case Uint64:
		return __obf_dcaad1165bb07af9.uint64()
	case Int8:
		return __obf_dcaad1165bb07af9.int8()
	case Int16:
		return __obf_dcaad1165bb07af9.int16()
	case Int32:
		return __obf_dcaad1165bb07af9.int32()
	case Int64:
		return __obf_dcaad1165bb07af9.int64()
	case Bin8, Bin16, Bin32:
		return __obf_dcaad1165bb07af9.bytes(__obf_ecf6d2d3253a058d, nil)
	case Str8, Str16, Str32:
		return __obf_dcaad1165bb07af9.string(__obf_ecf6d2d3253a058d)
	case Array16, Array32:
		return __obf_dcaad1165bb07af9.__obf_9d79cf59b9daca03(__obf_ecf6d2d3253a058d)
	case Map16, Map32:
		__obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_a93d004caac96500.UnreadByte()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_8468f219671f8d5a()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_dcaad1165bb07af9.__obf_7b329929aba3053d(__obf_ecf6d2d3253a058d)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_ecf6d2d3253a058d)
}

// DecodeInterfaceLoose is like DecodeInterface except that:
//   - int8, int16, and int32 are converted to int64,
//   - uint8, uint16, and uint32 are converted to uint64,
//   - float32 is converted to float64.
//   - []byte is converted to string.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeInterfaceLoose() (any, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}

	if IsFixedNum(__obf_ecf6d2d3253a058d) {
		return int64(int8(__obf_ecf6d2d3253a058d)), nil
	}
	if IsFixedMap(__obf_ecf6d2d3253a058d) {
		__obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_a93d004caac96500.UnreadByte()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_8468f219671f8d5a()
	}
	if IsFixedArray(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.__obf_9d79cf59b9daca03(__obf_ecf6d2d3253a058d)
	}
	if IsFixedString(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.string(__obf_ecf6d2d3253a058d)
	}

	switch __obf_ecf6d2d3253a058d {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_dcaad1165bb07af9.bool(__obf_ecf6d2d3253a058d)
	case Float, Double:
		return __obf_dcaad1165bb07af9.float64(__obf_ecf6d2d3253a058d)
	case Uint8, Uint16, Uint32, Uint64:
		return __obf_dcaad1165bb07af9.uint(__obf_ecf6d2d3253a058d)
	case Int8, Int16, Int32, Int64:
		return __obf_dcaad1165bb07af9.int(__obf_ecf6d2d3253a058d)
	case Str8, Str16, Str32,
		Bin8, Bin16, Bin32:
		return __obf_dcaad1165bb07af9.string(__obf_ecf6d2d3253a058d)
	case Array16, Array32:
		return __obf_dcaad1165bb07af9.__obf_9d79cf59b9daca03(__obf_ecf6d2d3253a058d)
	case Map16, Map32:
		__obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_a93d004caac96500.UnreadByte()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		return __obf_dcaad1165bb07af9.__obf_8468f219671f8d5a()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_dcaad1165bb07af9.__obf_7b329929aba3053d(__obf_ecf6d2d3253a058d)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_ecf6d2d3253a058d)
}

// Skip skips next value.
func (__obf_dcaad1165bb07af9 *Decoder) Skip() error {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	if IsFixedNum(__obf_ecf6d2d3253a058d) {
		return nil
	}
	if IsFixedMap(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.__obf_293bf0315964dccb(__obf_ecf6d2d3253a058d)
	}
	if IsFixedArray(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.__obf_f7821d6a7eae9517(__obf_ecf6d2d3253a058d)
	}
	if IsFixedString(__obf_ecf6d2d3253a058d) {
		return __obf_dcaad1165bb07af9.__obf_d22198f61bd105ee(__obf_ecf6d2d3253a058d)
	}

	switch __obf_ecf6d2d3253a058d {
	case Nil, False, True:
		return nil
	case Uint8, Int8:
		return __obf_dcaad1165bb07af9.__obf_6f540045c1336cac(1)
	case Uint16, Int16:
		return __obf_dcaad1165bb07af9.__obf_6f540045c1336cac(2)
	case Uint32, Int32, Float:
		return __obf_dcaad1165bb07af9.__obf_6f540045c1336cac(4)
	case Uint64, Int64, Double:
		return __obf_dcaad1165bb07af9.__obf_6f540045c1336cac(8)
	case Bin8, Bin16, Bin32:
		return __obf_dcaad1165bb07af9.__obf_d22198f61bd105ee(__obf_ecf6d2d3253a058d)
	case Str8, Str16, Str32:
		return __obf_dcaad1165bb07af9.__obf_d22198f61bd105ee(__obf_ecf6d2d3253a058d)
	case Array16, Array32:
		return __obf_dcaad1165bb07af9.__obf_f7821d6a7eae9517(__obf_ecf6d2d3253a058d)
	case Map16, Map32:
		return __obf_dcaad1165bb07af9.__obf_293bf0315964dccb(__obf_ecf6d2d3253a058d)
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_dcaad1165bb07af9.__obf_ebb31dcc8229349b(__obf_ecf6d2d3253a058d)
	}

	return fmt.Errorf("msgpack: unknown code %x", __obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeRaw() (RawMessage, error) {
	__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = make([]byte, 0)
	if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	__obf_72c36aa69a5d48f8 := RawMessage(__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf)
	__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = nil
	return __obf_72c36aa69a5d48f8, nil
}

// PeekCode returns the next MessagePack code without advancing the reader.
// Subpackage msgpack/codes defines the list of available
func (__obf_dcaad1165bb07af9 *Decoder) PeekCode() (byte, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_a93d004caac96500.ReadByte()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_ecf6d2d3253a058d, __obf_dcaad1165bb07af9.

		// ReadFull reads exactly len(buf) bytes into the buf.
		__obf_a93d004caac96500.UnreadByte()
}

func (__obf_dcaad1165bb07af9 *Decoder) ReadFull(__obf_9c8c6a5e4cc8822c []byte) error {
	_, __obf_0feb3528b7b709ec := __obf_8f5a813341411779(__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c, __obf_9c8c6a5e4cc8822c, len(__obf_9c8c6a5e4cc8822c))
	return __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_61d7724cadbffcc4() bool {
	__obf_34e3ba264a6bb541, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.PeekCode()
	return __obf_0feb3528b7b709ec == nil && __obf_34e3ba264a6bb541 == Nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_74ac05d9a9da01f4() (byte, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_a93d004caac96500.ReadByte()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	if __obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf != nil {
		__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = append(__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf, __obf_ecf6d2d3253a058d)
	}
	return __obf_ecf6d2d3253a058d, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_73059ce4dea16cf3(__obf_a7fd7acf2bd4435f []byte) error {
	_, __obf_0feb3528b7b709ec := io.ReadFull(__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c, __obf_a7fd7acf2bd4435f)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf != nil {
		__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = append(__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf, __obf_a7fd7acf2bd4435f...)
	}
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_8f5a813341411779(__obf_2b0247e819b1bf4a int) ([]byte, error) {
	var __obf_0feb3528b7b709ec error
	if __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_c646a26ee823abc3 != 0 {
		__obf_dcaad1165bb07af9.__obf_9c8c6a5e4cc8822c, __obf_0feb3528b7b709ec = __obf_8f5a813341411779(__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c, __obf_dcaad1165bb07af9.__obf_9c8c6a5e4cc8822c, __obf_2b0247e819b1bf4a)
	} else {
		__obf_dcaad1165bb07af9.__obf_9c8c6a5e4cc8822c, __obf_0feb3528b7b709ec = __obf_0f60ebc434d83ecd(__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c, __obf_dcaad1165bb07af9.__obf_9c8c6a5e4cc8822c,

			// TODO: read directly into d.rec?
			__obf_2b0247e819b1bf4a)
	}
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	if __obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf != nil {
		__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = append(__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf, __obf_dcaad1165bb07af9.__obf_9c8c6a5e4cc8822c...)
	}
	return __obf_dcaad1165bb07af9.__obf_9c8c6a5e4cc8822c, nil
}

func __obf_8f5a813341411779(__obf_4609cadeaf50917c io.Reader, __obf_a7fd7acf2bd4435f []byte, __obf_2b0247e819b1bf4a int) ([]byte, error) {
	if __obf_a7fd7acf2bd4435f == nil {
		if __obf_2b0247e819b1bf4a == 0 {
			return make([]byte, 0), nil
		}
		__obf_a7fd7acf2bd4435f = make([]byte, 0, __obf_2b0247e819b1bf4a)
	}

	if __obf_2b0247e819b1bf4a > cap(__obf_a7fd7acf2bd4435f) {
		__obf_a7fd7acf2bd4435f = append(__obf_a7fd7acf2bd4435f, make([]byte, __obf_2b0247e819b1bf4a-len(__obf_a7fd7acf2bd4435f))...)
	} else if __obf_2b0247e819b1bf4a <= cap(__obf_a7fd7acf2bd4435f) {
		__obf_a7fd7acf2bd4435f = __obf_a7fd7acf2bd4435f[:__obf_2b0247e819b1bf4a]
	}

	_, __obf_0feb3528b7b709ec := io.ReadFull(__obf_4609cadeaf50917c, __obf_a7fd7acf2bd4435f)
	return __obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec
}

func __obf_0f60ebc434d83ecd(__obf_4609cadeaf50917c io.Reader, __obf_a7fd7acf2bd4435f []byte, __obf_2b0247e819b1bf4a int) ([]byte, error) {
	if __obf_a7fd7acf2bd4435f == nil {
		if __obf_2b0247e819b1bf4a == 0 {
			return make([]byte, 0), nil
		}
		switch {
		case __obf_2b0247e819b1bf4a < 64:
			__obf_a7fd7acf2bd4435f = make([]byte, 0, 64)
		case __obf_2b0247e819b1bf4a <= __obf_d570f482b4bbddc6:
			__obf_a7fd7acf2bd4435f = make([]byte, 0, __obf_2b0247e819b1bf4a)
		default:
			__obf_a7fd7acf2bd4435f = make([]byte, 0, __obf_d570f482b4bbddc6)
		}
	}

	if __obf_2b0247e819b1bf4a <= cap(__obf_a7fd7acf2bd4435f) {
		__obf_a7fd7acf2bd4435f = __obf_a7fd7acf2bd4435f[:__obf_2b0247e819b1bf4a]
		_, __obf_0feb3528b7b709ec := io.ReadFull(__obf_4609cadeaf50917c, __obf_a7fd7acf2bd4435f)
		return __obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec
	}
	__obf_a7fd7acf2bd4435f = __obf_a7fd7acf2bd4435f[:cap(__obf_a7fd7acf2bd4435f)]

	var __obf_69e9263181462409 int
	for {
		__obf_2a22e1acdde91e55 := min(__obf_2b0247e819b1bf4a-len(__obf_a7fd7acf2bd4435f), __obf_d570f482b4bbddc6)
		__obf_a7fd7acf2bd4435f = append(__obf_a7fd7acf2bd4435f, make([]byte, __obf_2a22e1acdde91e55)...)

		_, __obf_0feb3528b7b709ec := io.ReadFull(__obf_4609cadeaf50917c, __obf_a7fd7acf2bd4435f[__obf_69e9263181462409:])
		if __obf_0feb3528b7b709ec != nil {
			return __obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec
		}

		if len(__obf_a7fd7acf2bd4435f) == __obf_2b0247e819b1bf4a {
			break
		}
		__obf_69e9263181462409 = len(__obf_a7fd7acf2bd4435f)
	}

	return __obf_a7fd7acf2bd4435f, nil
}

func min(__obf_00d82cdd95284094, //nolint:unparam
	__obf_a7fd7acf2bd4435f int) int {
	if __obf_00d82cdd95284094 <= __obf_a7fd7acf2bd4435f {
		return __obf_00d82cdd95284094
	}
	return __obf_a7fd7acf2bd4435f
}

func __obf_431a947ae748a61f(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.bytes(__obf_ecf6d2d3253a058d, __obf_17732590722140e7.Bytes())
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetBytes(__obf_a7fd7acf2bd4435f)
	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeBytesLen() (int, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeBytes() ([]byte, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.bytes(__obf_ecf6d2d3253a058d, nil)
}

func (__obf_dcaad1165bb07af9 *Decoder) bytes(__obf_ecf6d2d3253a058d byte, __obf_a7fd7acf2bd4435f []byte) ([]byte, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_869c18dc6d397629(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return nil, nil
	}
	return __obf_8f5a813341411779(__obf_dcaad1165bb07af9.__obf_4609cadeaf50917c, __obf_a7fd7acf2bd4435f, __obf_2b0247e819b1bf4a)
}
