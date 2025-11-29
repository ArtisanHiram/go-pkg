package __obf_a4aad98aaf3178f4

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
	__obf_a4710c2e9b57d831 = 1 << 20
	__obf_953f2fb1425dda7a = // 1mb
	1e6
	__obf_0dcaecf62133ed63 = // 1m elements
	1e6                      // 1m elements
)

const (
	__obf_8cb66a8d0819d7d2 uint32 = 1 << iota
	__obf_00d6d684a301507c
	__obf_97ba1004800fc0a8
	__obf_beecb71ff037aa1b
)

type __obf_e6e503ea75c6dcbb interface {
	io.Reader
	io.ByteScanner
}

//------------------------------------------------------------------------------

var __obf_a129852354c9a7a4 = sync.Pool{
	New: func() any {
		return NewDecoder(nil)
	},
}

func GetDecoder() *Decoder {
	return __obf_a129852354c9a7a4.Get().(*Decoder)
}

func PutDecoder(__obf_cc68ee77d25ca8f2 *Decoder) {
	__obf_cc68ee77d25ca8f2.__obf_f9ba010dbe3d722e = nil
	__obf_cc68ee77d25ca8f2.__obf_759f458bfa19abba = nil
	__obf_a129852354c9a7a4.
		Put(__obf_cc68ee77d25ca8f2)
}

//------------------------------------------------------------------------------

// Unmarshal decodes the MessagePack-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(__obf_69e4c283179273ce []byte, __obf_a1f43267eeb48a1a any) error {
	__obf_cc68ee77d25ca8f2 := GetDecoder()
	__obf_cc68ee77d25ca8f2.
		UsePreallocateValues(true)
	__obf_cc68ee77d25ca8f2.
		Reset(bytes.NewReader(__obf_69e4c283179273ce))
	__obf_4061ca0039150c39 := __obf_cc68ee77d25ca8f2.Decode(__obf_a1f43267eeb48a1a)

	PutDecoder(__obf_cc68ee77d25ca8f2)

	return __obf_4061ca0039150c39
}

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	__obf_f9ba010dbe3d722e io.Reader
	__obf_759f458bfa19abba io.ByteScanner
	__obf_d20effb6daf18186 func(*Decoder) (any, error)
	__obf_72ba705ed504d210 string
	__obf_f5d391a424aa5a3a []byte
	__obf_9af56e4ab518ffc1 []byte
	__obf_57edbb6a6615f57c []string
	__obf_a8400e70a712e9fa uint32
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r
// beyond the requested msgpack values. Buffering can be disabled
// by passing a reader that implements io.ByteScanner interface.
func NewDecoder(__obf_f9ba010dbe3d722e io.Reader) *Decoder {
	__obf_613397fefdec0ed0 := new(Decoder)
	__obf_613397fefdec0ed0.
		Reset(__obf_f9ba010dbe3d722e)
	return __obf_613397fefdec0ed0
}

// Reset discards any buffered data, resets all state, and switches the buffered
// reader to read from r.
func (__obf_613397fefdec0ed0 *Decoder) Reset(__obf_f9ba010dbe3d722e io.Reader) {
	__obf_613397fefdec0ed0.
		ResetDict(__obf_f9ba010dbe3d722e, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_613397fefdec0ed0 *Decoder) ResetDict(__obf_f9ba010dbe3d722e io.Reader, __obf_57edbb6a6615f57c []string) {
	__obf_613397fefdec0ed0.
		ResetReader(__obf_f9ba010dbe3d722e)
	__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa = 0
	__obf_613397fefdec0ed0.__obf_72ba705ed504d210 = ""
	__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c = __obf_57edbb6a6615f57c
}

func (__obf_613397fefdec0ed0 *Decoder) WithDict(__obf_57edbb6a6615f57c []string, __obf_64fe495dbc0c543b func(*Decoder) error) error {
	__obf_187a5bedd2cd4c82 := __obf_613397fefdec0ed0.__obf_57edbb6a6615f57c
	__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c = __obf_57edbb6a6615f57c
	__obf_4061ca0039150c39 := __obf_64fe495dbc0c543b(__obf_613397fefdec0ed0)
	__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c = __obf_187a5bedd2cd4c82
	return __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) ResetReader(__obf_f9ba010dbe3d722e io.Reader) {
	__obf_613397fefdec0ed0.__obf_d20effb6daf18186 = nil
	__obf_613397fefdec0ed0.__obf_57edbb6a6615f57c = nil

	if __obf_a16dd3b787484a29, __obf_81b19f2a6159ab89 := __obf_f9ba010dbe3d722e.(__obf_e6e503ea75c6dcbb); __obf_81b19f2a6159ab89 {
		__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e = __obf_a16dd3b787484a29
		__obf_613397fefdec0ed0.__obf_759f458bfa19abba = __obf_a16dd3b787484a29
	} else if __obf_f9ba010dbe3d722e == nil {
		__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e = nil
		__obf_613397fefdec0ed0.__obf_759f458bfa19abba = nil
	} else {
		__obf_a16dd3b787484a29 := bufio.NewReader(__obf_f9ba010dbe3d722e)
		__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e = __obf_a16dd3b787484a29
		__obf_613397fefdec0ed0.__obf_759f458bfa19abba = __obf_a16dd3b787484a29
	}
}

func (__obf_613397fefdec0ed0 *Decoder) SetMapDecoder(__obf_64fe495dbc0c543b func(*Decoder) (any, error)) {
	__obf_613397fefdec0ed0.__obf_d20effb6daf18186 = __obf_64fe495dbc0c543b
}

// UseLooseInterfaceDecoding causes decoder to use DecodeInterfaceLoose
// to decode msgpack value into Go some.
func (__obf_613397fefdec0ed0 *Decoder) UseLooseInterfaceDecoding(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa |= __obf_8cb66a8d0819d7d2
	} else {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa &= ^__obf_8cb66a8d0819d7d2
	}
}

// SetCustomStructTag causes the decoder to use the supplied tag as a fallback option
// if there is no msgpack tag.
func (__obf_613397fefdec0ed0 *Decoder) SetCustomStructTag(__obf_990470bfaf220103 string) {
	__obf_613397fefdec0ed0.__obf_72ba705ed504d210 = __obf_990470bfaf220103
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (__obf_613397fefdec0ed0 *Decoder) DisallowUnknownFields(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa |= __obf_00d6d684a301507c
	} else {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa &= ^__obf_00d6d684a301507c
	}
}

// UseInternedStrings enables support for decoding interned strings.
func (__obf_613397fefdec0ed0 *Decoder) UseInternedStrings(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa |= __obf_ca379a8882c3af10
	} else {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa &= ^__obf_ca379a8882c3af10
	}
}

// UsePreallocateValues enables preallocating values in chunks
func (__obf_613397fefdec0ed0 *Decoder) UsePreallocateValues(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa |= __obf_97ba1004800fc0a8
	} else {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa &= ^__obf_97ba1004800fc0a8
	}
}

// DisableAllocLimit enables fully allocating slices/maps when the size is known
func (__obf_613397fefdec0ed0 *Decoder) DisableAllocLimit(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa |= __obf_beecb71ff037aa1b
	} else {
		__obf_613397fefdec0ed0.__obf_a8400e70a712e9fa &= ^__obf_beecb71ff037aa1b
	}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (__obf_613397fefdec0ed0 *Decoder) Buffered() io.Reader {
	return __obf_613397fefdec0ed0.

		//nolint:gocyclo
		__obf_f9ba010dbe3d722e
}

func (__obf_613397fefdec0ed0 *Decoder) Decode(__obf_a1f43267eeb48a1a any) error {
	var __obf_4061ca0039150c39 error
	switch __obf_a1f43267eeb48a1a := __obf_a1f43267eeb48a1a.(type) {
	case *string:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeString()
			return __obf_4061ca0039150c39
		}
	case *[]byte:
		if __obf_a1f43267eeb48a1a != nil {
			return __obf_613397fefdec0ed0.__obf_16359546865c48fc(__obf_a1f43267eeb48a1a)
		}
	case *int:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeInt()
			return __obf_4061ca0039150c39
		}
	case *int8:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeInt8()
			return __obf_4061ca0039150c39
		}
	case *int16:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeInt16()
			return __obf_4061ca0039150c39
		}
	case *int32:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeInt32()
			return __obf_4061ca0039150c39
		}
	case *int64:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeInt64()
			return __obf_4061ca0039150c39
		}
	case *uint:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeUint()
			return __obf_4061ca0039150c39
		}
	case *uint8:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeUint8()
			return __obf_4061ca0039150c39
		}
	case *uint16:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeUint16()
			return __obf_4061ca0039150c39
		}
	case *uint32:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeUint32()
			return __obf_4061ca0039150c39
		}
	case *uint64:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeUint64()
			return __obf_4061ca0039150c39
		}
	case *bool:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeBool()
			return __obf_4061ca0039150c39
		}
	case *float32:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeFloat32()
			return __obf_4061ca0039150c39
		}
	case *float64:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeFloat64()
			return __obf_4061ca0039150c39
		}
	case *[]string:
		return __obf_613397fefdec0ed0.__obf_3f7e7f51ea420948(__obf_a1f43267eeb48a1a)
	case *map[string]string:
		return __obf_613397fefdec0ed0.__obf_b01b056f781c9c97(__obf_a1f43267eeb48a1a)
	case *map[string]any:
		return __obf_613397fefdec0ed0.__obf_27d72527258884d5(__obf_a1f43267eeb48a1a)
	case *time.Duration:
		if __obf_a1f43267eeb48a1a != nil {
			__obf_9f1e21befe556e87, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
			*__obf_a1f43267eeb48a1a = time.Duration(__obf_9f1e21befe556e87)
			return __obf_4061ca0039150c39
		}
	case *time.Time:
		if __obf_a1f43267eeb48a1a != nil {
			*__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 = __obf_613397fefdec0ed0.DecodeTime()
			return __obf_4061ca0039150c39
		}
	}
	__obf_9f1e21befe556e87 := reflect.ValueOf(__obf_a1f43267eeb48a1a)
	if !__obf_9f1e21befe556e87.IsValid() {
		return errors.New("msgpack: Decode(nil)")
	}
	if __obf_9f1e21befe556e87.Kind() != reflect.Ptr {
		return fmt.Errorf("msgpack: Decode(non-pointer %T)", __obf_a1f43267eeb48a1a)
	}
	if __obf_9f1e21befe556e87.IsNil() {
		return fmt.Errorf("msgpack: Decode(non-settable %T)", __obf_a1f43267eeb48a1a)
	}
	__obf_9f1e21befe556e87 = __obf_9f1e21befe556e87.Elem()
	if __obf_9f1e21befe556e87.Kind() == reflect.Interface {
		if !__obf_9f1e21befe556e87.IsNil() {
			__obf_9f1e21befe556e87 = __obf_9f1e21befe556e87.Elem()
			if __obf_9f1e21befe556e87.Kind() != reflect.Ptr {
				return fmt.Errorf("msgpack: Decode(non-pointer %s)", __obf_9f1e21befe556e87.Type().String())
			}
		}
	}

	return __obf_613397fefdec0ed0.DecodeValue(__obf_9f1e21befe556e87)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeMulti(__obf_a1f43267eeb48a1a ...any) error {
	for _, __obf_9f1e21befe556e87 := range __obf_a1f43267eeb48a1a {
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Decode(__obf_9f1e21befe556e87); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_ec86210fb048ddc6() (any, error) {
	if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_8cb66a8d0819d7d2 != 0 {
		return __obf_613397fefdec0ed0.DecodeInterfaceLoose()
	}
	return __obf_613397fefdec0ed0.DecodeInterface()
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeValue(__obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_5acdbf512a9b89fe := __obf_66268eb3a3deccf5(__obf_a1f43267eeb48a1a.Type())
	return __obf_5acdbf512a9b89fe(__obf_613397fefdec0ed0, __obf_a1f43267eeb48a1a)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeNil() error {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_6dbe86b3d9d9b037 != Nil {
		return fmt.Errorf("msgpack: invalid code=%x decoding nil", __obf_6dbe86b3d9d9b037)
	}
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_82c6f7a4267374a0(__obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeNil()
	if __obf_a1f43267eeb48a1a.IsNil() {
		return __obf_4061ca0039150c39
	}
	if __obf_a1f43267eeb48a1a.Kind() == reflect.Ptr {
		__obf_a1f43267eeb48a1a = __obf_a1f43267eeb48a1a.Elem()
	}
	__obf_a1f43267eeb48a1a.
		Set(reflect.Zero(__obf_a1f43267eeb48a1a.Type()))
	return __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeBool() (bool, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return false, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.bool(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) bool(__obf_6dbe86b3d9d9b037 byte) (bool, error) {
	if __obf_6dbe86b3d9d9b037 == Nil {
		return false, nil
	}
	if __obf_6dbe86b3d9d9b037 == False {
		return false, nil
	}
	if __obf_6dbe86b3d9d9b037 == True {
		return true, nil
	}
	return false, fmt.Errorf("msgpack: invalid code=%x decoding bool", __obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeDuration() (time.Duration, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeInt64()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return time.Duration(__obf_99a74e41c9c640c0), nil
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
func (__obf_613397fefdec0ed0 *Decoder) DecodeInterface() (any, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}

	if IsFixedNum(__obf_6dbe86b3d9d9b037) {
		return int8(__obf_6dbe86b3d9d9b037), nil
	}
	if IsFixedMap(__obf_6dbe86b3d9d9b037) {
		__obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_759f458bfa19abba.UnreadByte()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_01c9a38c34ed0945()
	}
	if IsFixedArray(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.__obf_1020cdfcb81f8fa6(__obf_6dbe86b3d9d9b037)
	}
	if IsFixedString(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.string(__obf_6dbe86b3d9d9b037)
	}

	switch __obf_6dbe86b3d9d9b037 {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_613397fefdec0ed0.bool(__obf_6dbe86b3d9d9b037)
	case Float:
		return __obf_613397fefdec0ed0.float32(__obf_6dbe86b3d9d9b037)
	case Double:
		return __obf_613397fefdec0ed0.float64(__obf_6dbe86b3d9d9b037)
	case Uint8:
		return __obf_613397fefdec0ed0.uint8()
	case Uint16:
		return __obf_613397fefdec0ed0.uint16()
	case Uint32:
		return __obf_613397fefdec0ed0.uint32()
	case Uint64:
		return __obf_613397fefdec0ed0.uint64()
	case Int8:
		return __obf_613397fefdec0ed0.int8()
	case Int16:
		return __obf_613397fefdec0ed0.int16()
	case Int32:
		return __obf_613397fefdec0ed0.int32()
	case Int64:
		return __obf_613397fefdec0ed0.int64()
	case Bin8, Bin16, Bin32:
		return __obf_613397fefdec0ed0.bytes(__obf_6dbe86b3d9d9b037, nil)
	case Str8, Str16, Str32:
		return __obf_613397fefdec0ed0.string(__obf_6dbe86b3d9d9b037)
	case Array16, Array32:
		return __obf_613397fefdec0ed0.__obf_1020cdfcb81f8fa6(__obf_6dbe86b3d9d9b037)
	case Map16, Map32:
		__obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_759f458bfa19abba.UnreadByte()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_01c9a38c34ed0945()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_613397fefdec0ed0.__obf_478bdbc02e8c61ec(__obf_6dbe86b3d9d9b037)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_6dbe86b3d9d9b037)
}

// DecodeInterfaceLoose is like DecodeInterface except that:
//   - int8, int16, and int32 are converted to int64,
//   - uint8, uint16, and uint32 are converted to uint64,
//   - float32 is converted to float64.
//   - []byte is converted to string.
func (__obf_613397fefdec0ed0 *Decoder) DecodeInterfaceLoose() (any, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}

	if IsFixedNum(__obf_6dbe86b3d9d9b037) {
		return int64(int8(__obf_6dbe86b3d9d9b037)), nil
	}
	if IsFixedMap(__obf_6dbe86b3d9d9b037) {
		__obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_759f458bfa19abba.UnreadByte()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_01c9a38c34ed0945()
	}
	if IsFixedArray(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.__obf_1020cdfcb81f8fa6(__obf_6dbe86b3d9d9b037)
	}
	if IsFixedString(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.string(__obf_6dbe86b3d9d9b037)
	}

	switch __obf_6dbe86b3d9d9b037 {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_613397fefdec0ed0.bool(__obf_6dbe86b3d9d9b037)
	case Float, Double:
		return __obf_613397fefdec0ed0.float64(__obf_6dbe86b3d9d9b037)
	case Uint8, Uint16, Uint32, Uint64:
		return __obf_613397fefdec0ed0.uint(__obf_6dbe86b3d9d9b037)
	case Int8, Int16, Int32, Int64:
		return __obf_613397fefdec0ed0.int(__obf_6dbe86b3d9d9b037)
	case Str8, Str16, Str32,
		Bin8, Bin16, Bin32:
		return __obf_613397fefdec0ed0.string(__obf_6dbe86b3d9d9b037)
	case Array16, Array32:
		return __obf_613397fefdec0ed0.__obf_1020cdfcb81f8fa6(__obf_6dbe86b3d9d9b037)
	case Map16, Map32:
		__obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_759f458bfa19abba.UnreadByte()
		if __obf_4061ca0039150c39 != nil {
			return nil, __obf_4061ca0039150c39
		}
		return __obf_613397fefdec0ed0.__obf_01c9a38c34ed0945()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_613397fefdec0ed0.__obf_478bdbc02e8c61ec(__obf_6dbe86b3d9d9b037)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_6dbe86b3d9d9b037)
}

// Skip skips next value.
func (__obf_613397fefdec0ed0 *Decoder) Skip() error {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	if IsFixedNum(__obf_6dbe86b3d9d9b037) {
		return nil
	}
	if IsFixedMap(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.__obf_a25930d7ee975bcc(__obf_6dbe86b3d9d9b037)
	}
	if IsFixedArray(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.__obf_b037cd157d532089(__obf_6dbe86b3d9d9b037)
	}
	if IsFixedString(__obf_6dbe86b3d9d9b037) {
		return __obf_613397fefdec0ed0.__obf_07d3af657b68439d(__obf_6dbe86b3d9d9b037)
	}

	switch __obf_6dbe86b3d9d9b037 {
	case Nil, False, True:
		return nil
	case Uint8, Int8:
		return __obf_613397fefdec0ed0.__obf_2bfec4fb91ed09f6(1)
	case Uint16, Int16:
		return __obf_613397fefdec0ed0.__obf_2bfec4fb91ed09f6(2)
	case Uint32, Int32, Float:
		return __obf_613397fefdec0ed0.__obf_2bfec4fb91ed09f6(4)
	case Uint64, Int64, Double:
		return __obf_613397fefdec0ed0.__obf_2bfec4fb91ed09f6(8)
	case Bin8, Bin16, Bin32:
		return __obf_613397fefdec0ed0.__obf_07d3af657b68439d(__obf_6dbe86b3d9d9b037)
	case Str8, Str16, Str32:
		return __obf_613397fefdec0ed0.__obf_07d3af657b68439d(__obf_6dbe86b3d9d9b037)
	case Array16, Array32:
		return __obf_613397fefdec0ed0.__obf_b037cd157d532089(__obf_6dbe86b3d9d9b037)
	case Map16, Map32:
		return __obf_613397fefdec0ed0.__obf_a25930d7ee975bcc(__obf_6dbe86b3d9d9b037)
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_613397fefdec0ed0.__obf_a07e0b14e7edfb58(__obf_6dbe86b3d9d9b037)
	}

	return fmt.Errorf("msgpack: unknown code %x", __obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeRaw() (RawMessage, error) {
	__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = make([]byte, 0)
	if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	__obf_2130e988c0fe510a := RawMessage(__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1)
	__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = nil
	return __obf_2130e988c0fe510a, nil
}

// PeekCode returns the next MessagePack code without advancing the reader.
// Subpackage msgpack/codes defines the list of available
func (__obf_613397fefdec0ed0 *Decoder) PeekCode() (byte, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_759f458bfa19abba.ReadByte()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_6dbe86b3d9d9b037, __obf_613397fefdec0ed0.

		// ReadFull reads exactly len(buf) bytes into the buf.
		__obf_759f458bfa19abba.UnreadByte()
}

func (__obf_613397fefdec0ed0 *Decoder) ReadFull(__obf_f5d391a424aa5a3a []byte) error {
	_, __obf_4061ca0039150c39 := __obf_4429146a43365de4(__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e, __obf_f5d391a424aa5a3a, len(__obf_f5d391a424aa5a3a))
	return __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_d6d4a347783e80b3() bool {
	__obf_987b95dd4dcfc308, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.PeekCode()
	return __obf_4061ca0039150c39 == nil && __obf_987b95dd4dcfc308 == Nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_46b95b57ed5617b5() (byte, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_759f458bfa19abba.ReadByte()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	if __obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 != nil {
		__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = append(__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1, __obf_6dbe86b3d9d9b037)
	}
	return __obf_6dbe86b3d9d9b037, nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_6814a41007039154(__obf_f57209cfda8e17cf []byte) error {
	_, __obf_4061ca0039150c39 := io.ReadFull(__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e, __obf_f57209cfda8e17cf)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 != nil {
		__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = append(__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1, __obf_f57209cfda8e17cf...)
	}
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_4429146a43365de4(__obf_99a74e41c9c640c0 int) ([]byte, error) {
	var __obf_4061ca0039150c39 error
	if __obf_613397fefdec0ed0.__obf_a8400e70a712e9fa&__obf_beecb71ff037aa1b != 0 {
		__obf_613397fefdec0ed0.__obf_f5d391a424aa5a3a, __obf_4061ca0039150c39 = __obf_4429146a43365de4(__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e, __obf_613397fefdec0ed0.__obf_f5d391a424aa5a3a, __obf_99a74e41c9c640c0)
	} else {
		__obf_613397fefdec0ed0.__obf_f5d391a424aa5a3a, __obf_4061ca0039150c39 = __obf_ad2683d8c108e7a0(__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e, __obf_613397fefdec0ed0.__obf_f5d391a424aa5a3a,

			// TODO: read directly into d.rec?
			__obf_99a74e41c9c640c0)
	}
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	if __obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 != nil {
		__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1 = append(__obf_613397fefdec0ed0.__obf_9af56e4ab518ffc1, __obf_613397fefdec0ed0.__obf_f5d391a424aa5a3a...)
	}
	return __obf_613397fefdec0ed0.__obf_f5d391a424aa5a3a, nil
}

func __obf_4429146a43365de4(__obf_f9ba010dbe3d722e io.Reader, __obf_f57209cfda8e17cf []byte, __obf_99a74e41c9c640c0 int) ([]byte, error) {
	if __obf_f57209cfda8e17cf == nil {
		if __obf_99a74e41c9c640c0 == 0 {
			return make([]byte, 0), nil
		}
		__obf_f57209cfda8e17cf = make([]byte, 0, __obf_99a74e41c9c640c0)
	}

	if __obf_99a74e41c9c640c0 > cap(__obf_f57209cfda8e17cf) {
		__obf_f57209cfda8e17cf = append(__obf_f57209cfda8e17cf, make([]byte, __obf_99a74e41c9c640c0-len(__obf_f57209cfda8e17cf))...)
	} else if __obf_99a74e41c9c640c0 <= cap(__obf_f57209cfda8e17cf) {
		__obf_f57209cfda8e17cf = __obf_f57209cfda8e17cf[:__obf_99a74e41c9c640c0]
	}

	_, __obf_4061ca0039150c39 := io.ReadFull(__obf_f9ba010dbe3d722e, __obf_f57209cfda8e17cf)
	return __obf_f57209cfda8e17cf, __obf_4061ca0039150c39
}

func __obf_ad2683d8c108e7a0(__obf_f9ba010dbe3d722e io.Reader, __obf_f57209cfda8e17cf []byte, __obf_99a74e41c9c640c0 int) ([]byte, error) {
	if __obf_f57209cfda8e17cf == nil {
		if __obf_99a74e41c9c640c0 == 0 {
			return make([]byte, 0), nil
		}
		switch {
		case __obf_99a74e41c9c640c0 < 64:
			__obf_f57209cfda8e17cf = make([]byte, 0, 64)
		case __obf_99a74e41c9c640c0 <= __obf_a4710c2e9b57d831:
			__obf_f57209cfda8e17cf = make([]byte, 0, __obf_99a74e41c9c640c0)
		default:
			__obf_f57209cfda8e17cf = make([]byte, 0, __obf_a4710c2e9b57d831)
		}
	}

	if __obf_99a74e41c9c640c0 <= cap(__obf_f57209cfda8e17cf) {
		__obf_f57209cfda8e17cf = __obf_f57209cfda8e17cf[:__obf_99a74e41c9c640c0]
		_, __obf_4061ca0039150c39 := io.ReadFull(__obf_f9ba010dbe3d722e, __obf_f57209cfda8e17cf)
		return __obf_f57209cfda8e17cf, __obf_4061ca0039150c39
	}
	__obf_f57209cfda8e17cf = __obf_f57209cfda8e17cf[:cap(__obf_f57209cfda8e17cf)]

	var __obf_aab3a57419803047 int
	for {
		__obf_5d330e1d7b0aa715 := min(__obf_99a74e41c9c640c0-len(__obf_f57209cfda8e17cf), __obf_a4710c2e9b57d831)
		__obf_f57209cfda8e17cf = append(__obf_f57209cfda8e17cf, make([]byte, __obf_5d330e1d7b0aa715)...)

		_, __obf_4061ca0039150c39 := io.ReadFull(__obf_f9ba010dbe3d722e, __obf_f57209cfda8e17cf[__obf_aab3a57419803047:])
		if __obf_4061ca0039150c39 != nil {
			return __obf_f57209cfda8e17cf, __obf_4061ca0039150c39
		}

		if len(__obf_f57209cfda8e17cf) == __obf_99a74e41c9c640c0 {
			break
		}
		__obf_aab3a57419803047 = len(__obf_f57209cfda8e17cf)
	}

	return __obf_f57209cfda8e17cf, nil
}

func min(__obf_9e17768eda0349c0, //nolint:unparam
	__obf_f57209cfda8e17cf int) int {
	if __obf_9e17768eda0349c0 <= __obf_f57209cfda8e17cf {
		return __obf_9e17768eda0349c0
	}
	return __obf_f57209cfda8e17cf
}

func __obf_9c86656b7d9d5cc1(__obf_613397fefdec0ed0 *Decoder, __obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_f57209cfda8e17cf, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.bytes(__obf_6dbe86b3d9d9b037, __obf_a1f43267eeb48a1a.Bytes())
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	__obf_a1f43267eeb48a1a.
		SetBytes(__obf_f57209cfda8e17cf)
	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeBytesLen() (int, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return 0, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
}

func (__obf_613397fefdec0ed0 *Decoder) DecodeBytes() ([]byte, error) {
	__obf_6dbe86b3d9d9b037, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_46b95b57ed5617b5()
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	return __obf_613397fefdec0ed0.bytes(__obf_6dbe86b3d9d9b037, nil)
}

func (__obf_613397fefdec0ed0 *Decoder) bytes(__obf_6dbe86b3d9d9b037 byte, __obf_f57209cfda8e17cf []byte) ([]byte, error) {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_0a2679fbfce7d32c(__obf_6dbe86b3d9d9b037)
	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return nil, nil
	}
	return __obf_4429146a43365de4(__obf_613397fefdec0ed0.__obf_f9ba010dbe3d722e, __obf_f57209cfda8e17cf, __obf_99a74e41c9c640c0)
}
