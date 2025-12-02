package __obf_a935eb7f548271a4

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
	__obf_3976b13eb0640f4f = 1 << 20
	__obf_69adbcf4aa0f33cb = // 1mb
	1e6
	__obf_c02b114384c3720b = // 1m elements
	1e6                      // 1m elements
)

const (
	__obf_66edf43e9dc81701 uint32 = 1 << iota
	__obf_5c467062833ca44f
	__obf_4cf2a1ef1c20b040
	__obf_20f89e0d9adcdf29
)

type __obf_4a72d494de5347e1 interface {
	io.Reader
	io.ByteScanner
}

//------------------------------------------------------------------------------

var __obf_7d5f1b706c87c1b6 = sync.Pool{
	New: func() any {
		return NewDecoder(nil)
	},
}

func GetDecoder() *Decoder {
	return __obf_7d5f1b706c87c1b6.Get().(*Decoder)
}

func PutDecoder(__obf_fbe570faa7707c02 *Decoder) {
	__obf_fbe570faa7707c02.__obf_4f94e5d98c7e4bb3 = nil
	__obf_fbe570faa7707c02.__obf_b62c60fba2fd9788 = nil
	__obf_7d5f1b706c87c1b6.
		Put(__obf_fbe570faa7707c02)
}

//------------------------------------------------------------------------------

// Unmarshal decodes the MessagePack-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(__obf_652c67787b9c24c9 []byte, __obf_6d570581f4b60dbc any) error {
	__obf_fbe570faa7707c02 := GetDecoder()
	__obf_fbe570faa7707c02.
		UsePreallocateValues(true)
	__obf_fbe570faa7707c02.
		Reset(bytes.NewReader(__obf_652c67787b9c24c9))
	__obf_4d327e1cd40c2e21 := __obf_fbe570faa7707c02.Decode(__obf_6d570581f4b60dbc)

	PutDecoder(__obf_fbe570faa7707c02)

	return __obf_4d327e1cd40c2e21
}

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	__obf_4f94e5d98c7e4bb3 io.Reader
	__obf_b62c60fba2fd9788 io.ByteScanner
	__obf_6ef8e6548ebf1dd5 func(*Decoder) (any, error)
	__obf_90e41d275c699c20 string
	__obf_a2e16952f247f218 []byte
	__obf_20ef30c3130f4f7f []byte
	__obf_2cdd7fad308aef1d []string
	__obf_d617f3769ce16c47 uint32
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r
// beyond the requested msgpack values. Buffering can be disabled
// by passing a reader that implements io.ByteScanner interface.
func NewDecoder(__obf_4f94e5d98c7e4bb3 io.Reader) *Decoder {
	__obf_a21885da2425f2b2 := new(Decoder)
	__obf_a21885da2425f2b2.
		Reset(__obf_4f94e5d98c7e4bb3)
	return __obf_a21885da2425f2b2
}

// Reset discards any buffered data, resets all state, and switches the buffered
// reader to read from r.
func (__obf_a21885da2425f2b2 *Decoder) Reset(__obf_4f94e5d98c7e4bb3 io.Reader) {
	__obf_a21885da2425f2b2.
		ResetDict(__obf_4f94e5d98c7e4bb3, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_a21885da2425f2b2 *Decoder) ResetDict(__obf_4f94e5d98c7e4bb3 io.Reader, __obf_2cdd7fad308aef1d []string) {
	__obf_a21885da2425f2b2.
		ResetReader(__obf_4f94e5d98c7e4bb3)
	__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 = 0
	__obf_a21885da2425f2b2.__obf_90e41d275c699c20 = ""
	__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d = __obf_2cdd7fad308aef1d
}

func (__obf_a21885da2425f2b2 *Decoder) WithDict(__obf_2cdd7fad308aef1d []string, __obf_36c78c696c47cfdb func(*Decoder) error) error {
	__obf_f893ccda83eefbf7 := __obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d
	__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d = __obf_2cdd7fad308aef1d
	__obf_4d327e1cd40c2e21 := __obf_36c78c696c47cfdb(__obf_a21885da2425f2b2)
	__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d = __obf_f893ccda83eefbf7
	return __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) ResetReader(__obf_4f94e5d98c7e4bb3 io.Reader) {
	__obf_a21885da2425f2b2.__obf_6ef8e6548ebf1dd5 = nil
	__obf_a21885da2425f2b2.__obf_2cdd7fad308aef1d = nil

	if __obf_0c0ec58080e43270, __obf_826ac2dbb957d7df := __obf_4f94e5d98c7e4bb3.(__obf_4a72d494de5347e1); __obf_826ac2dbb957d7df {
		__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3 = __obf_0c0ec58080e43270
		__obf_a21885da2425f2b2.__obf_b62c60fba2fd9788 = __obf_0c0ec58080e43270
	} else if __obf_4f94e5d98c7e4bb3 == nil {
		__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3 = nil
		__obf_a21885da2425f2b2.__obf_b62c60fba2fd9788 = nil
	} else {
		__obf_0c0ec58080e43270 := bufio.NewReader(__obf_4f94e5d98c7e4bb3)
		__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3 = __obf_0c0ec58080e43270
		__obf_a21885da2425f2b2.__obf_b62c60fba2fd9788 = __obf_0c0ec58080e43270
	}
}

func (__obf_a21885da2425f2b2 *Decoder) SetMapDecoder(__obf_36c78c696c47cfdb func(*Decoder) (any, error)) {
	__obf_a21885da2425f2b2.__obf_6ef8e6548ebf1dd5 = __obf_36c78c696c47cfdb
}

// UseLooseInterfaceDecoding causes decoder to use DecodeInterfaceLoose
// to decode msgpack value into Go some.
func (__obf_a21885da2425f2b2 *Decoder) UseLooseInterfaceDecoding(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 |= __obf_66edf43e9dc81701
	} else {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 &= ^__obf_66edf43e9dc81701
	}
}

// SetCustomStructTag causes the decoder to use the supplied tag as a fallback option
// if there is no msgpack tag.
func (__obf_a21885da2425f2b2 *Decoder) SetCustomStructTag(__obf_f03106ae8627dc7d string) {
	__obf_a21885da2425f2b2.__obf_90e41d275c699c20 = __obf_f03106ae8627dc7d
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (__obf_a21885da2425f2b2 *Decoder) DisallowUnknownFields(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 |= __obf_5c467062833ca44f
	} else {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 &= ^__obf_5c467062833ca44f
	}
}

// UseInternedStrings enables support for decoding interned strings.
func (__obf_a21885da2425f2b2 *Decoder) UseInternedStrings(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 |= __obf_bf24fd2d64b91255
	} else {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 &= ^__obf_bf24fd2d64b91255
	}
}

// UsePreallocateValues enables preallocating values in chunks
func (__obf_a21885da2425f2b2 *Decoder) UsePreallocateValues(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 |= __obf_4cf2a1ef1c20b040
	} else {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 &= ^__obf_4cf2a1ef1c20b040
	}
}

// DisableAllocLimit enables fully allocating slices/maps when the size is known
func (__obf_a21885da2425f2b2 *Decoder) DisableAllocLimit(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 |= __obf_20f89e0d9adcdf29
	} else {
		__obf_a21885da2425f2b2.__obf_d617f3769ce16c47 &= ^__obf_20f89e0d9adcdf29
	}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (__obf_a21885da2425f2b2 *Decoder) Buffered() io.Reader {
	return __obf_a21885da2425f2b2.

		//nolint:gocyclo
		__obf_4f94e5d98c7e4bb3
}

func (__obf_a21885da2425f2b2 *Decoder) Decode(__obf_6d570581f4b60dbc any) error {
	var __obf_4d327e1cd40c2e21 error
	switch __obf_6d570581f4b60dbc := __obf_6d570581f4b60dbc.(type) {
	case *string:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeString()
			return __obf_4d327e1cd40c2e21
		}
	case *[]byte:
		if __obf_6d570581f4b60dbc != nil {
			return __obf_a21885da2425f2b2.__obf_24d3dffd03973291(__obf_6d570581f4b60dbc)
		}
	case *int:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeInt()
			return __obf_4d327e1cd40c2e21
		}
	case *int8:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeInt8()
			return __obf_4d327e1cd40c2e21
		}
	case *int16:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeInt16()
			return __obf_4d327e1cd40c2e21
		}
	case *int32:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeInt32()
			return __obf_4d327e1cd40c2e21
		}
	case *int64:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeInt64()
			return __obf_4d327e1cd40c2e21
		}
	case *uint:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeUint()
			return __obf_4d327e1cd40c2e21
		}
	case *uint8:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeUint8()
			return __obf_4d327e1cd40c2e21
		}
	case *uint16:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeUint16()
			return __obf_4d327e1cd40c2e21
		}
	case *uint32:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeUint32()
			return __obf_4d327e1cd40c2e21
		}
	case *uint64:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeUint64()
			return __obf_4d327e1cd40c2e21
		}
	case *bool:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeBool()
			return __obf_4d327e1cd40c2e21
		}
	case *float32:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeFloat32()
			return __obf_4d327e1cd40c2e21
		}
	case *float64:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeFloat64()
			return __obf_4d327e1cd40c2e21
		}
	case *[]string:
		return __obf_a21885da2425f2b2.__obf_95cef988e387ad44(__obf_6d570581f4b60dbc)
	case *map[string]string:
		return __obf_a21885da2425f2b2.__obf_17bf0c8c3e4bc6de(__obf_6d570581f4b60dbc)
	case *map[string]any:
		return __obf_a21885da2425f2b2.__obf_7b03c6c7ec65752f(__obf_6d570581f4b60dbc)
	case *time.Duration:
		if __obf_6d570581f4b60dbc != nil {
			__obf_6fee38a618ed3afa, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
			*__obf_6d570581f4b60dbc = time.Duration(__obf_6fee38a618ed3afa)
			return __obf_4d327e1cd40c2e21
		}
	case *time.Time:
		if __obf_6d570581f4b60dbc != nil {
			*__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.DecodeTime()
			return __obf_4d327e1cd40c2e21
		}
	}
	__obf_6fee38a618ed3afa := reflect.ValueOf(__obf_6d570581f4b60dbc)
	if !__obf_6fee38a618ed3afa.IsValid() {
		return errors.New("msgpack: Decode(nil)")
	}
	if __obf_6fee38a618ed3afa.Kind() != reflect.Ptr {
		return fmt.Errorf("msgpack: Decode(non-pointer %T)", __obf_6d570581f4b60dbc)
	}
	if __obf_6fee38a618ed3afa.IsNil() {
		return fmt.Errorf("msgpack: Decode(non-settable %T)", __obf_6d570581f4b60dbc)
	}
	__obf_6fee38a618ed3afa = __obf_6fee38a618ed3afa.Elem()
	if __obf_6fee38a618ed3afa.Kind() == reflect.Interface {
		if !__obf_6fee38a618ed3afa.IsNil() {
			__obf_6fee38a618ed3afa = __obf_6fee38a618ed3afa.Elem()
			if __obf_6fee38a618ed3afa.Kind() != reflect.Ptr {
				return fmt.Errorf("msgpack: Decode(non-pointer %s)", __obf_6fee38a618ed3afa.Type().String())
			}
		}
	}

	return __obf_a21885da2425f2b2.DecodeValue(__obf_6fee38a618ed3afa)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeMulti(__obf_6d570581f4b60dbc ...any) error {
	for _, __obf_6fee38a618ed3afa := range __obf_6d570581f4b60dbc {
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Decode(__obf_6fee38a618ed3afa); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_c4227998727b27c6() (any, error) {
	if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_66edf43e9dc81701 != 0 {
		return __obf_a21885da2425f2b2.DecodeInterfaceLoose()
	}
	return __obf_a21885da2425f2b2.DecodeInterface()
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeValue(__obf_6d570581f4b60dbc reflect.Value) error {
	__obf_e408804459f261f2 := __obf_7581b02f483ae445(__obf_6d570581f4b60dbc.Type())
	return __obf_e408804459f261f2(__obf_a21885da2425f2b2, __obf_6d570581f4b60dbc)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeNil() error {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_f5df560f4d67421b != Nil {
		return fmt.Errorf("msgpack: invalid code=%x decoding nil", __obf_f5df560f4d67421b)
	}
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_30a79fe20d689305(__obf_6d570581f4b60dbc reflect.Value) error {
	__obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeNil()
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_6d570581f4b60dbc.Kind() == reflect.Ptr {
		__obf_6d570581f4b60dbc = __obf_6d570581f4b60dbc.Elem()
	}
	__obf_6d570581f4b60dbc.
		Set(reflect.Zero(__obf_6d570581f4b60dbc.Type()))
	return __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeBool() (bool, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return false, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.bool(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) bool(__obf_f5df560f4d67421b byte) (bool, error) {
	if __obf_f5df560f4d67421b == Nil {
		return false, nil
	}
	if __obf_f5df560f4d67421b == False {
		return false, nil
	}
	if __obf_f5df560f4d67421b == True {
		return true, nil
	}
	return false, fmt.Errorf("msgpack: invalid code=%x decoding bool", __obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeDuration() (time.Duration, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeInt64()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return time.Duration(__obf_326af9bd942662ac), nil
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
func (__obf_a21885da2425f2b2 *Decoder) DecodeInterface() (any, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}

	if IsFixedNum(__obf_f5df560f4d67421b) {
		return int8(__obf_f5df560f4d67421b), nil
	}
	if IsFixedMap(__obf_f5df560f4d67421b) {
		__obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.UnreadByte()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_1d0c0ae9a1e78fcd()
	}
	if IsFixedArray(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.__obf_9d4786f56b1c142d(__obf_f5df560f4d67421b)
	}
	if IsFixedString(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.string(__obf_f5df560f4d67421b)
	}

	switch __obf_f5df560f4d67421b {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_a21885da2425f2b2.bool(__obf_f5df560f4d67421b)
	case Float:
		return __obf_a21885da2425f2b2.float32(__obf_f5df560f4d67421b)
	case Double:
		return __obf_a21885da2425f2b2.float64(__obf_f5df560f4d67421b)
	case Uint8:
		return __obf_a21885da2425f2b2.uint8()
	case Uint16:
		return __obf_a21885da2425f2b2.uint16()
	case Uint32:
		return __obf_a21885da2425f2b2.uint32()
	case Uint64:
		return __obf_a21885da2425f2b2.uint64()
	case Int8:
		return __obf_a21885da2425f2b2.int8()
	case Int16:
		return __obf_a21885da2425f2b2.int16()
	case Int32:
		return __obf_a21885da2425f2b2.int32()
	case Int64:
		return __obf_a21885da2425f2b2.int64()
	case Bin8, Bin16, Bin32:
		return __obf_a21885da2425f2b2.bytes(__obf_f5df560f4d67421b, nil)
	case Str8, Str16, Str32:
		return __obf_a21885da2425f2b2.string(__obf_f5df560f4d67421b)
	case Array16, Array32:
		return __obf_a21885da2425f2b2.__obf_9d4786f56b1c142d(__obf_f5df560f4d67421b)
	case Map16, Map32:
		__obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.UnreadByte()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_1d0c0ae9a1e78fcd()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_a21885da2425f2b2.__obf_f93129cdacf36826(__obf_f5df560f4d67421b)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_f5df560f4d67421b)
}

// DecodeInterfaceLoose is like DecodeInterface except that:
//   - int8, int16, and int32 are converted to int64,
//   - uint8, uint16, and uint32 are converted to uint64,
//   - float32 is converted to float64.
//   - []byte is converted to string.
func (__obf_a21885da2425f2b2 *Decoder) DecodeInterfaceLoose() (any, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}

	if IsFixedNum(__obf_f5df560f4d67421b) {
		return int64(int8(__obf_f5df560f4d67421b)), nil
	}
	if IsFixedMap(__obf_f5df560f4d67421b) {
		__obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.UnreadByte()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_1d0c0ae9a1e78fcd()
	}
	if IsFixedArray(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.__obf_9d4786f56b1c142d(__obf_f5df560f4d67421b)
	}
	if IsFixedString(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.string(__obf_f5df560f4d67421b)
	}

	switch __obf_f5df560f4d67421b {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_a21885da2425f2b2.bool(__obf_f5df560f4d67421b)
	case Float, Double:
		return __obf_a21885da2425f2b2.float64(__obf_f5df560f4d67421b)
	case Uint8, Uint16, Uint32, Uint64:
		return __obf_a21885da2425f2b2.uint(__obf_f5df560f4d67421b)
	case Int8, Int16, Int32, Int64:
		return __obf_a21885da2425f2b2.int(__obf_f5df560f4d67421b)
	case Str8, Str16, Str32,
		Bin8, Bin16, Bin32:
		return __obf_a21885da2425f2b2.string(__obf_f5df560f4d67421b)
	case Array16, Array32:
		return __obf_a21885da2425f2b2.__obf_9d4786f56b1c142d(__obf_f5df560f4d67421b)
	case Map16, Map32:
		__obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.UnreadByte()
		if __obf_4d327e1cd40c2e21 != nil {
			return nil, __obf_4d327e1cd40c2e21
		}
		return __obf_a21885da2425f2b2.__obf_1d0c0ae9a1e78fcd()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_a21885da2425f2b2.__obf_f93129cdacf36826(__obf_f5df560f4d67421b)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_f5df560f4d67421b)
}

// Skip skips next value.
func (__obf_a21885da2425f2b2 *Decoder) Skip() error {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if IsFixedNum(__obf_f5df560f4d67421b) {
		return nil
	}
	if IsFixedMap(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.__obf_ad6b3dbd72d54464(__obf_f5df560f4d67421b)
	}
	if IsFixedArray(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.__obf_533ee38d6a6706df(__obf_f5df560f4d67421b)
	}
	if IsFixedString(__obf_f5df560f4d67421b) {
		return __obf_a21885da2425f2b2.__obf_4b7348bc9d246008(__obf_f5df560f4d67421b)
	}

	switch __obf_f5df560f4d67421b {
	case Nil, False, True:
		return nil
	case Uint8, Int8:
		return __obf_a21885da2425f2b2.__obf_b2541a0cb78c8e1f(1)
	case Uint16, Int16:
		return __obf_a21885da2425f2b2.__obf_b2541a0cb78c8e1f(2)
	case Uint32, Int32, Float:
		return __obf_a21885da2425f2b2.__obf_b2541a0cb78c8e1f(4)
	case Uint64, Int64, Double:
		return __obf_a21885da2425f2b2.__obf_b2541a0cb78c8e1f(8)
	case Bin8, Bin16, Bin32:
		return __obf_a21885da2425f2b2.__obf_4b7348bc9d246008(__obf_f5df560f4d67421b)
	case Str8, Str16, Str32:
		return __obf_a21885da2425f2b2.__obf_4b7348bc9d246008(__obf_f5df560f4d67421b)
	case Array16, Array32:
		return __obf_a21885da2425f2b2.__obf_533ee38d6a6706df(__obf_f5df560f4d67421b)
	case Map16, Map32:
		return __obf_a21885da2425f2b2.__obf_ad6b3dbd72d54464(__obf_f5df560f4d67421b)
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_a21885da2425f2b2.__obf_fa04fccd68154ccb(__obf_f5df560f4d67421b)
	}

	return fmt.Errorf("msgpack: unknown code %x", __obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeRaw() (RawMessage, error) {
	__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = make([]byte, 0)
	if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	__obf_8d91290aff37bb4e := RawMessage(__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f)
	__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = nil
	return __obf_8d91290aff37bb4e, nil
}

// PeekCode returns the next MessagePack code without advancing the reader.
// Subpackage msgpack/codes defines the list of available
func (__obf_a21885da2425f2b2 *Decoder) PeekCode() (byte, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.ReadByte()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_f5df560f4d67421b, __obf_a21885da2425f2b2.

		// ReadFull reads exactly len(buf) bytes into the buf.
		__obf_b62c60fba2fd9788.UnreadByte()
}

func (__obf_a21885da2425f2b2 *Decoder) ReadFull(__obf_a2e16952f247f218 []byte) error {
	_, __obf_4d327e1cd40c2e21 := __obf_595cbf51b6653ebf(__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3, __obf_a2e16952f247f218, len(__obf_a2e16952f247f218))
	return __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_5e95865dadae4ab4() bool {
	__obf_b983039ef2a336eb, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.PeekCode()
	return __obf_4d327e1cd40c2e21 == nil && __obf_b983039ef2a336eb == Nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_61e5879fcf7b35ce() (byte, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_b62c60fba2fd9788.ReadByte()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	if __obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f != nil {
		__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = append(__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f, __obf_f5df560f4d67421b)
	}
	return __obf_f5df560f4d67421b, nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_fd322376412b2809(__obf_f2ca794293605b73 []byte) error {
	_, __obf_4d327e1cd40c2e21 := io.ReadFull(__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3, __obf_f2ca794293605b73)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f != nil {
		__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = append(__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f, __obf_f2ca794293605b73...)
	}
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_595cbf51b6653ebf(__obf_326af9bd942662ac int) ([]byte, error) {
	var __obf_4d327e1cd40c2e21 error
	if __obf_a21885da2425f2b2.__obf_d617f3769ce16c47&__obf_20f89e0d9adcdf29 != 0 {
		__obf_a21885da2425f2b2.__obf_a2e16952f247f218, __obf_4d327e1cd40c2e21 = __obf_595cbf51b6653ebf(__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3, __obf_a21885da2425f2b2.__obf_a2e16952f247f218, __obf_326af9bd942662ac)
	} else {
		__obf_a21885da2425f2b2.__obf_a2e16952f247f218, __obf_4d327e1cd40c2e21 = __obf_f7a37502c6f621d8(__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3, __obf_a21885da2425f2b2.__obf_a2e16952f247f218,

			// TODO: read directly into d.rec?
			__obf_326af9bd942662ac)
	}
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	if __obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f != nil {
		__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f = append(__obf_a21885da2425f2b2.__obf_20ef30c3130f4f7f, __obf_a21885da2425f2b2.__obf_a2e16952f247f218...)
	}
	return __obf_a21885da2425f2b2.__obf_a2e16952f247f218, nil
}

func __obf_595cbf51b6653ebf(__obf_4f94e5d98c7e4bb3 io.Reader, __obf_f2ca794293605b73 []byte, __obf_326af9bd942662ac int) ([]byte, error) {
	if __obf_f2ca794293605b73 == nil {
		if __obf_326af9bd942662ac == 0 {
			return make([]byte, 0), nil
		}
		__obf_f2ca794293605b73 = make([]byte, 0, __obf_326af9bd942662ac)
	}

	if __obf_326af9bd942662ac > cap(__obf_f2ca794293605b73) {
		__obf_f2ca794293605b73 = append(__obf_f2ca794293605b73, make([]byte, __obf_326af9bd942662ac-len(__obf_f2ca794293605b73))...)
	} else if __obf_326af9bd942662ac <= cap(__obf_f2ca794293605b73) {
		__obf_f2ca794293605b73 = __obf_f2ca794293605b73[:__obf_326af9bd942662ac]
	}

	_, __obf_4d327e1cd40c2e21 := io.ReadFull(__obf_4f94e5d98c7e4bb3, __obf_f2ca794293605b73)
	return __obf_f2ca794293605b73, __obf_4d327e1cd40c2e21
}

func __obf_f7a37502c6f621d8(__obf_4f94e5d98c7e4bb3 io.Reader, __obf_f2ca794293605b73 []byte, __obf_326af9bd942662ac int) ([]byte, error) {
	if __obf_f2ca794293605b73 == nil {
		if __obf_326af9bd942662ac == 0 {
			return make([]byte, 0), nil
		}
		switch {
		case __obf_326af9bd942662ac < 64:
			__obf_f2ca794293605b73 = make([]byte, 0, 64)
		case __obf_326af9bd942662ac <= __obf_3976b13eb0640f4f:
			__obf_f2ca794293605b73 = make([]byte, 0, __obf_326af9bd942662ac)
		default:
			__obf_f2ca794293605b73 = make([]byte, 0, __obf_3976b13eb0640f4f)
		}
	}

	if __obf_326af9bd942662ac <= cap(__obf_f2ca794293605b73) {
		__obf_f2ca794293605b73 = __obf_f2ca794293605b73[:__obf_326af9bd942662ac]
		_, __obf_4d327e1cd40c2e21 := io.ReadFull(__obf_4f94e5d98c7e4bb3, __obf_f2ca794293605b73)
		return __obf_f2ca794293605b73, __obf_4d327e1cd40c2e21
	}
	__obf_f2ca794293605b73 = __obf_f2ca794293605b73[:cap(__obf_f2ca794293605b73)]

	var __obf_8088e8f5046df03f int
	for {
		__obf_0df6e434dcab8fe0 := min(__obf_326af9bd942662ac-len(__obf_f2ca794293605b73), __obf_3976b13eb0640f4f)
		__obf_f2ca794293605b73 = append(__obf_f2ca794293605b73, make([]byte, __obf_0df6e434dcab8fe0)...)

		_, __obf_4d327e1cd40c2e21 := io.ReadFull(__obf_4f94e5d98c7e4bb3, __obf_f2ca794293605b73[__obf_8088e8f5046df03f:])
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_f2ca794293605b73, __obf_4d327e1cd40c2e21
		}

		if len(__obf_f2ca794293605b73) == __obf_326af9bd942662ac {
			break
		}
		__obf_8088e8f5046df03f = len(__obf_f2ca794293605b73)
	}

	return __obf_f2ca794293605b73, nil
}

func min(__obf_8374f4d53f54dc7c, //nolint:unparam
	__obf_f2ca794293605b73 int) int {
	if __obf_8374f4d53f54dc7c <= __obf_f2ca794293605b73 {
		return __obf_8374f4d53f54dc7c
	}
	return __obf_f2ca794293605b73
}

func __obf_1c67f7750bd90b37(__obf_a21885da2425f2b2 *Decoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.bytes(__obf_f5df560f4d67421b, __obf_6d570581f4b60dbc.Bytes())
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	__obf_6d570581f4b60dbc.
		SetBytes(__obf_f2ca794293605b73)
	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeBytesLen() (int, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return 0, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
}

func (__obf_a21885da2425f2b2 *Decoder) DecodeBytes() ([]byte, error) {
	__obf_f5df560f4d67421b, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_61e5879fcf7b35ce()
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	return __obf_a21885da2425f2b2.bytes(__obf_f5df560f4d67421b, nil)
}

func (__obf_a21885da2425f2b2 *Decoder) bytes(__obf_f5df560f4d67421b byte, __obf_f2ca794293605b73 []byte) ([]byte, error) {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c0a95532c414ce37(__obf_f5df560f4d67421b)
	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return nil, nil
	}
	return __obf_595cbf51b6653ebf(__obf_a21885da2425f2b2.__obf_4f94e5d98c7e4bb3, __obf_f2ca794293605b73, __obf_326af9bd942662ac)
}
