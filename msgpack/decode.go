package __obf_3edaa49e853afa16

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
	__obf_a5e732b826466c45 = 1 << 20
	__obf_dc3dfe27c3ffba99 = // 1mb
	1e6
	__obf_5d6aaa093307f6f6 = // 1m elements
	1e6                      // 1m elements
)

const (
	__obf_c776950ec0cb354a uint32 = 1 << iota
	__obf_2b5f039e8d42091e
	__obf_5157ad164f309c0c
	__obf_9d8a15eb3a073456
)

type __obf_5ceeb46d66c5edfe interface {
	io.Reader
	io.ByteScanner
}

//------------------------------------------------------------------------------

var __obf_eba8abeefb498c7b = sync.Pool{
	New: func() any {
		return NewDecoder(nil)
	},
}

func GetDecoder() *Decoder {
	return __obf_eba8abeefb498c7b.Get().(*Decoder)
}

func PutDecoder(__obf_20b87cf22b03338d *Decoder) {
	__obf_20b87cf22b03338d.__obf_9b78235c82fd11c0 = nil
	__obf_20b87cf22b03338d.__obf_6827ff1b59d7ccec = nil
	__obf_eba8abeefb498c7b.
		Put(__obf_20b87cf22b03338d)
}

//------------------------------------------------------------------------------

// Unmarshal decodes the MessagePack-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(__obf_2171bf5dcbdd2ac8 []byte, __obf_85d270343a0dfe14 any) error {
	__obf_20b87cf22b03338d := GetDecoder()
	__obf_20b87cf22b03338d.
		UsePreallocateValues(true)
	__obf_20b87cf22b03338d.
		Reset(bytes.NewReader(__obf_2171bf5dcbdd2ac8))
	__obf_3aa78ad28f50ed46 := __obf_20b87cf22b03338d.Decode(__obf_85d270343a0dfe14)

	PutDecoder(__obf_20b87cf22b03338d)

	return __obf_3aa78ad28f50ed46
}

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	__obf_9b78235c82fd11c0 io.Reader
	__obf_6827ff1b59d7ccec io.ByteScanner
	__obf_e63e8a485c325197 func(*Decoder) (any, error)
	__obf_ec5d9a4ce8dd41be string
	__obf_8f3bed0c23610f75 []byte
	__obf_e11087a315c4c191 []byte
	__obf_b014355f64826d7b []string
	__obf_6aa0efd537415192 uint32
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r
// beyond the requested msgpack values. Buffering can be disabled
// by passing a reader that implements io.ByteScanner interface.
func NewDecoder(__obf_9b78235c82fd11c0 io.Reader) *Decoder {
	__obf_015afbee33862a01 := new(Decoder)
	__obf_015afbee33862a01.
		Reset(__obf_9b78235c82fd11c0)
	return __obf_015afbee33862a01
}

// Reset discards any buffered data, resets all state, and switches the buffered
// reader to read from r.
func (__obf_015afbee33862a01 *Decoder) Reset(__obf_9b78235c82fd11c0 io.Reader) {
	__obf_015afbee33862a01.
		ResetDict(__obf_9b78235c82fd11c0, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_015afbee33862a01 *Decoder) ResetDict(__obf_9b78235c82fd11c0 io.Reader, __obf_b014355f64826d7b []string) {
	__obf_015afbee33862a01.
		ResetReader(__obf_9b78235c82fd11c0)
	__obf_015afbee33862a01.__obf_6aa0efd537415192 = 0
	__obf_015afbee33862a01.__obf_ec5d9a4ce8dd41be = ""
	__obf_015afbee33862a01.__obf_b014355f64826d7b = __obf_b014355f64826d7b
}

func (__obf_015afbee33862a01 *Decoder) WithDict(__obf_b014355f64826d7b []string, __obf_c8137d8e88d5b1dd func(*Decoder) error) error {
	__obf_119d5920d120ff61 := __obf_015afbee33862a01.__obf_b014355f64826d7b
	__obf_015afbee33862a01.__obf_b014355f64826d7b = __obf_b014355f64826d7b
	__obf_3aa78ad28f50ed46 := __obf_c8137d8e88d5b1dd(__obf_015afbee33862a01)
	__obf_015afbee33862a01.__obf_b014355f64826d7b = __obf_119d5920d120ff61
	return __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) ResetReader(__obf_9b78235c82fd11c0 io.Reader) {
	__obf_015afbee33862a01.__obf_e63e8a485c325197 = nil
	__obf_015afbee33862a01.__obf_b014355f64826d7b = nil

	if __obf_777a304a139de399, __obf_ccfb92cc26e4569f := __obf_9b78235c82fd11c0.(__obf_5ceeb46d66c5edfe); __obf_ccfb92cc26e4569f {
		__obf_015afbee33862a01.__obf_9b78235c82fd11c0 = __obf_777a304a139de399
		__obf_015afbee33862a01.__obf_6827ff1b59d7ccec = __obf_777a304a139de399
	} else if __obf_9b78235c82fd11c0 == nil {
		__obf_015afbee33862a01.__obf_9b78235c82fd11c0 = nil
		__obf_015afbee33862a01.__obf_6827ff1b59d7ccec = nil
	} else {
		__obf_777a304a139de399 := bufio.NewReader(__obf_9b78235c82fd11c0)
		__obf_015afbee33862a01.__obf_9b78235c82fd11c0 = __obf_777a304a139de399
		__obf_015afbee33862a01.__obf_6827ff1b59d7ccec = __obf_777a304a139de399
	}
}

func (__obf_015afbee33862a01 *Decoder) SetMapDecoder(__obf_c8137d8e88d5b1dd func(*Decoder) (any, error)) {
	__obf_015afbee33862a01.__obf_e63e8a485c325197 = __obf_c8137d8e88d5b1dd
}

// UseLooseInterfaceDecoding causes decoder to use DecodeInterfaceLoose
// to decode msgpack value into Go some.
func (__obf_015afbee33862a01 *Decoder) UseLooseInterfaceDecoding(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 |= __obf_c776950ec0cb354a
	} else {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 &= ^__obf_c776950ec0cb354a
	}
}

// SetCustomStructTag causes the decoder to use the supplied tag as a fallback option
// if there is no msgpack tag.
func (__obf_015afbee33862a01 *Decoder) SetCustomStructTag(__obf_160a7dd2c812a6a1 string) {
	__obf_015afbee33862a01.__obf_ec5d9a4ce8dd41be = __obf_160a7dd2c812a6a1
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (__obf_015afbee33862a01 *Decoder) DisallowUnknownFields(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 |= __obf_2b5f039e8d42091e
	} else {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 &= ^__obf_2b5f039e8d42091e
	}
}

// UseInternedStrings enables support for decoding interned strings.
func (__obf_015afbee33862a01 *Decoder) UseInternedStrings(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 |= __obf_48e006ba02b70b78
	} else {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 &= ^__obf_48e006ba02b70b78
	}
}

// UsePreallocateValues enables preallocating values in chunks
func (__obf_015afbee33862a01 *Decoder) UsePreallocateValues(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 |= __obf_5157ad164f309c0c
	} else {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 &= ^__obf_5157ad164f309c0c
	}
}

// DisableAllocLimit enables fully allocating slices/maps when the size is known
func (__obf_015afbee33862a01 *Decoder) DisableAllocLimit(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 |= __obf_9d8a15eb3a073456
	} else {
		__obf_015afbee33862a01.__obf_6aa0efd537415192 &= ^__obf_9d8a15eb3a073456
	}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (__obf_015afbee33862a01 *Decoder) Buffered() io.Reader {
	return __obf_015afbee33862a01.

		//nolint:gocyclo
		__obf_9b78235c82fd11c0
}

func (__obf_015afbee33862a01 *Decoder) Decode(__obf_85d270343a0dfe14 any) error {
	var __obf_3aa78ad28f50ed46 error
	switch __obf_85d270343a0dfe14 := __obf_85d270343a0dfe14.(type) {
	case *string:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeString()
			return __obf_3aa78ad28f50ed46
		}
	case *[]byte:
		if __obf_85d270343a0dfe14 != nil {
			return __obf_015afbee33862a01.__obf_2b47e19d74adf37a(__obf_85d270343a0dfe14)
		}
	case *int:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeInt()
			return __obf_3aa78ad28f50ed46
		}
	case *int8:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeInt8()
			return __obf_3aa78ad28f50ed46
		}
	case *int16:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeInt16()
			return __obf_3aa78ad28f50ed46
		}
	case *int32:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeInt32()
			return __obf_3aa78ad28f50ed46
		}
	case *int64:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeInt64()
			return __obf_3aa78ad28f50ed46
		}
	case *uint:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeUint()
			return __obf_3aa78ad28f50ed46
		}
	case *uint8:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeUint8()
			return __obf_3aa78ad28f50ed46
		}
	case *uint16:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeUint16()
			return __obf_3aa78ad28f50ed46
		}
	case *uint32:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeUint32()
			return __obf_3aa78ad28f50ed46
		}
	case *uint64:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeUint64()
			return __obf_3aa78ad28f50ed46
		}
	case *bool:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeBool()
			return __obf_3aa78ad28f50ed46
		}
	case *float32:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeFloat32()
			return __obf_3aa78ad28f50ed46
		}
	case *float64:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeFloat64()
			return __obf_3aa78ad28f50ed46
		}
	case *[]string:
		return __obf_015afbee33862a01.__obf_a75b379e5bfbda38(__obf_85d270343a0dfe14)
	case *map[string]string:
		return __obf_015afbee33862a01.__obf_7e971a691672ab09(__obf_85d270343a0dfe14)
	case *map[string]any:
		return __obf_015afbee33862a01.__obf_c5ba7e7660e774e8(__obf_85d270343a0dfe14)
	case *time.Duration:
		if __obf_85d270343a0dfe14 != nil {
			__obf_ab0d31723a444d49, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
			*__obf_85d270343a0dfe14 = time.Duration(__obf_ab0d31723a444d49)
			return __obf_3aa78ad28f50ed46
		}
	case *time.Time:
		if __obf_85d270343a0dfe14 != nil {
			*__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.DecodeTime()
			return __obf_3aa78ad28f50ed46
		}
	}
	__obf_ab0d31723a444d49 := reflect.ValueOf(__obf_85d270343a0dfe14)
	if !__obf_ab0d31723a444d49.IsValid() {
		return errors.New("msgpack: Decode(nil)")
	}
	if __obf_ab0d31723a444d49.Kind() != reflect.Ptr {
		return fmt.Errorf("msgpack: Decode(non-pointer %T)", __obf_85d270343a0dfe14)
	}
	if __obf_ab0d31723a444d49.IsNil() {
		return fmt.Errorf("msgpack: Decode(non-settable %T)", __obf_85d270343a0dfe14)
	}
	__obf_ab0d31723a444d49 = __obf_ab0d31723a444d49.Elem()
	if __obf_ab0d31723a444d49.Kind() == reflect.Interface {
		if !__obf_ab0d31723a444d49.IsNil() {
			__obf_ab0d31723a444d49 = __obf_ab0d31723a444d49.Elem()
			if __obf_ab0d31723a444d49.Kind() != reflect.Ptr {
				return fmt.Errorf("msgpack: Decode(non-pointer %s)", __obf_ab0d31723a444d49.Type().String())
			}
		}
	}

	return __obf_015afbee33862a01.DecodeValue(__obf_ab0d31723a444d49)
}

func (__obf_015afbee33862a01 *Decoder) DecodeMulti(__obf_85d270343a0dfe14 ...any) error {
	for _, __obf_ab0d31723a444d49 := range __obf_85d270343a0dfe14 {
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Decode(__obf_ab0d31723a444d49); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_ed59dd67f1d39cb1() (any, error) {
	if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_c776950ec0cb354a != 0 {
		return __obf_015afbee33862a01.DecodeInterfaceLoose()
	}
	return __obf_015afbee33862a01.DecodeInterface()
}

func (__obf_015afbee33862a01 *Decoder) DecodeValue(__obf_85d270343a0dfe14 reflect.Value) error {
	__obf_263376b30f868da0 := __obf_6253a63b14aba59e(__obf_85d270343a0dfe14.Type())
	return __obf_263376b30f868da0(__obf_015afbee33862a01, __obf_85d270343a0dfe14)
}

func (__obf_015afbee33862a01 *Decoder) DecodeNil() error {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_145c7a7d25eea2bd != Nil {
		return fmt.Errorf("msgpack: invalid code=%x decoding nil", __obf_145c7a7d25eea2bd)
	}
	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_38b8c150d8f5b61b(__obf_85d270343a0dfe14 reflect.Value) error {
	__obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeNil()
	if __obf_85d270343a0dfe14.IsNil() {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_85d270343a0dfe14.Kind() == reflect.Ptr {
		__obf_85d270343a0dfe14 = __obf_85d270343a0dfe14.Elem()
	}
	__obf_85d270343a0dfe14.
		Set(reflect.Zero(__obf_85d270343a0dfe14.Type()))
	return __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) DecodeBool() (bool, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return false, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.bool(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) bool(__obf_145c7a7d25eea2bd byte) (bool, error) {
	if __obf_145c7a7d25eea2bd == Nil {
		return false, nil
	}
	if __obf_145c7a7d25eea2bd == False {
		return false, nil
	}
	if __obf_145c7a7d25eea2bd == True {
		return true, nil
	}
	return false, fmt.Errorf("msgpack: invalid code=%x decoding bool", __obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) DecodeDuration() (time.Duration, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeInt64()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return time.Duration(__obf_56127cd370854f0b), nil
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
func (__obf_015afbee33862a01 *Decoder) DecodeInterface() (any, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}

	if IsFixedNum(__obf_145c7a7d25eea2bd) {
		return int8(__obf_145c7a7d25eea2bd), nil
	}
	if IsFixedMap(__obf_145c7a7d25eea2bd) {
		__obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.UnreadByte()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_987a7825732f4656()
	}
	if IsFixedArray(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.__obf_977775eeaf115ac9(__obf_145c7a7d25eea2bd)
	}
	if IsFixedString(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.string(__obf_145c7a7d25eea2bd)
	}

	switch __obf_145c7a7d25eea2bd {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_015afbee33862a01.bool(__obf_145c7a7d25eea2bd)
	case Float:
		return __obf_015afbee33862a01.float32(__obf_145c7a7d25eea2bd)
	case Double:
		return __obf_015afbee33862a01.float64(__obf_145c7a7d25eea2bd)
	case Uint8:
		return __obf_015afbee33862a01.uint8()
	case Uint16:
		return __obf_015afbee33862a01.uint16()
	case Uint32:
		return __obf_015afbee33862a01.uint32()
	case Uint64:
		return __obf_015afbee33862a01.uint64()
	case Int8:
		return __obf_015afbee33862a01.int8()
	case Int16:
		return __obf_015afbee33862a01.int16()
	case Int32:
		return __obf_015afbee33862a01.int32()
	case Int64:
		return __obf_015afbee33862a01.int64()
	case Bin8, Bin16, Bin32:
		return __obf_015afbee33862a01.bytes(__obf_145c7a7d25eea2bd, nil)
	case Str8, Str16, Str32:
		return __obf_015afbee33862a01.string(__obf_145c7a7d25eea2bd)
	case Array16, Array32:
		return __obf_015afbee33862a01.__obf_977775eeaf115ac9(__obf_145c7a7d25eea2bd)
	case Map16, Map32:
		__obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.UnreadByte()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_987a7825732f4656()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_015afbee33862a01.__obf_8795dc4d75fc75c5(__obf_145c7a7d25eea2bd)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_145c7a7d25eea2bd)
}

// DecodeInterfaceLoose is like DecodeInterface except that:
//   - int8, int16, and int32 are converted to int64,
//   - uint8, uint16, and uint32 are converted to uint64,
//   - float32 is converted to float64.
//   - []byte is converted to string.
func (__obf_015afbee33862a01 *Decoder) DecodeInterfaceLoose() (any, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}

	if IsFixedNum(__obf_145c7a7d25eea2bd) {
		return int64(int8(__obf_145c7a7d25eea2bd)), nil
	}
	if IsFixedMap(__obf_145c7a7d25eea2bd) {
		__obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.UnreadByte()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_987a7825732f4656()
	}
	if IsFixedArray(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.__obf_977775eeaf115ac9(__obf_145c7a7d25eea2bd)
	}
	if IsFixedString(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.string(__obf_145c7a7d25eea2bd)
	}

	switch __obf_145c7a7d25eea2bd {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_015afbee33862a01.bool(__obf_145c7a7d25eea2bd)
	case Float, Double:
		return __obf_015afbee33862a01.float64(__obf_145c7a7d25eea2bd)
	case Uint8, Uint16, Uint32, Uint64:
		return __obf_015afbee33862a01.uint(__obf_145c7a7d25eea2bd)
	case Int8, Int16, Int32, Int64:
		return __obf_015afbee33862a01.int(__obf_145c7a7d25eea2bd)
	case Str8, Str16, Str32,
		Bin8, Bin16, Bin32:
		return __obf_015afbee33862a01.string(__obf_145c7a7d25eea2bd)
	case Array16, Array32:
		return __obf_015afbee33862a01.__obf_977775eeaf115ac9(__obf_145c7a7d25eea2bd)
	case Map16, Map32:
		__obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.UnreadByte()
		if __obf_3aa78ad28f50ed46 != nil {
			return nil, __obf_3aa78ad28f50ed46
		}
		return __obf_015afbee33862a01.__obf_987a7825732f4656()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_015afbee33862a01.__obf_8795dc4d75fc75c5(__obf_145c7a7d25eea2bd)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_145c7a7d25eea2bd)
}

// Skip skips next value.
func (__obf_015afbee33862a01 *Decoder) Skip() error {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	if IsFixedNum(__obf_145c7a7d25eea2bd) {
		return nil
	}
	if IsFixedMap(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.__obf_609d01762403129e(__obf_145c7a7d25eea2bd)
	}
	if IsFixedArray(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.__obf_39e3b2fed086f4e4(__obf_145c7a7d25eea2bd)
	}
	if IsFixedString(__obf_145c7a7d25eea2bd) {
		return __obf_015afbee33862a01.__obf_6d8f264a5e217b4a(__obf_145c7a7d25eea2bd)
	}

	switch __obf_145c7a7d25eea2bd {
	case Nil, False, True:
		return nil
	case Uint8, Int8:
		return __obf_015afbee33862a01.__obf_1a431295b897a1b7(1)
	case Uint16, Int16:
		return __obf_015afbee33862a01.__obf_1a431295b897a1b7(2)
	case Uint32, Int32, Float:
		return __obf_015afbee33862a01.__obf_1a431295b897a1b7(4)
	case Uint64, Int64, Double:
		return __obf_015afbee33862a01.__obf_1a431295b897a1b7(8)
	case Bin8, Bin16, Bin32:
		return __obf_015afbee33862a01.__obf_6d8f264a5e217b4a(__obf_145c7a7d25eea2bd)
	case Str8, Str16, Str32:
		return __obf_015afbee33862a01.__obf_6d8f264a5e217b4a(__obf_145c7a7d25eea2bd)
	case Array16, Array32:
		return __obf_015afbee33862a01.__obf_39e3b2fed086f4e4(__obf_145c7a7d25eea2bd)
	case Map16, Map32:
		return __obf_015afbee33862a01.__obf_609d01762403129e(__obf_145c7a7d25eea2bd)
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_015afbee33862a01.__obf_31f60c6085a2cba9(__obf_145c7a7d25eea2bd)
	}

	return fmt.Errorf("msgpack: unknown code %x", __obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) DecodeRaw() (RawMessage, error) {
	__obf_015afbee33862a01.__obf_e11087a315c4c191 = make([]byte, 0)
	if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	__obf_771cde5bcb38d993 := RawMessage(__obf_015afbee33862a01.__obf_e11087a315c4c191)
	__obf_015afbee33862a01.__obf_e11087a315c4c191 = nil
	return __obf_771cde5bcb38d993, nil
}

// PeekCode returns the next MessagePack code without advancing the reader.
// Subpackage msgpack/codes defines the list of available
func (__obf_015afbee33862a01 *Decoder) PeekCode() (byte, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.ReadByte()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_145c7a7d25eea2bd, __obf_015afbee33862a01.

		// ReadFull reads exactly len(buf) bytes into the buf.
		__obf_6827ff1b59d7ccec.UnreadByte()
}

func (__obf_015afbee33862a01 *Decoder) ReadFull(__obf_8f3bed0c23610f75 []byte) error {
	_, __obf_3aa78ad28f50ed46 := __obf_2c3d4aa0a72cea84(__obf_015afbee33862a01.__obf_9b78235c82fd11c0, __obf_8f3bed0c23610f75, len(__obf_8f3bed0c23610f75))
	return __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) __obf_7773819fc7c589cb() bool {
	__obf_508e2d6ff53655c0, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.PeekCode()
	return __obf_3aa78ad28f50ed46 == nil && __obf_508e2d6ff53655c0 == Nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_4c5f8d154c9027ea() (byte, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_6827ff1b59d7ccec.ReadByte()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	if __obf_015afbee33862a01.__obf_e11087a315c4c191 != nil {
		__obf_015afbee33862a01.__obf_e11087a315c4c191 = append(__obf_015afbee33862a01.__obf_e11087a315c4c191, __obf_145c7a7d25eea2bd)
	}
	return __obf_145c7a7d25eea2bd, nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_8ba8963a8ad56c5c(__obf_9b4a5a04bdad2347 []byte) error {
	_, __obf_3aa78ad28f50ed46 := io.ReadFull(__obf_015afbee33862a01.__obf_9b78235c82fd11c0, __obf_9b4a5a04bdad2347)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_015afbee33862a01.__obf_e11087a315c4c191 != nil {
		__obf_015afbee33862a01.__obf_e11087a315c4c191 = append(__obf_015afbee33862a01.__obf_e11087a315c4c191, __obf_9b4a5a04bdad2347...)
	}
	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_2c3d4aa0a72cea84(__obf_56127cd370854f0b int) ([]byte, error) {
	var __obf_3aa78ad28f50ed46 error
	if __obf_015afbee33862a01.__obf_6aa0efd537415192&__obf_9d8a15eb3a073456 != 0 {
		__obf_015afbee33862a01.__obf_8f3bed0c23610f75, __obf_3aa78ad28f50ed46 = __obf_2c3d4aa0a72cea84(__obf_015afbee33862a01.__obf_9b78235c82fd11c0, __obf_015afbee33862a01.__obf_8f3bed0c23610f75, __obf_56127cd370854f0b)
	} else {
		__obf_015afbee33862a01.__obf_8f3bed0c23610f75, __obf_3aa78ad28f50ed46 = __obf_470c46bee8a8fb03(__obf_015afbee33862a01.__obf_9b78235c82fd11c0, __obf_015afbee33862a01.__obf_8f3bed0c23610f75,

			// TODO: read directly into d.rec?
			__obf_56127cd370854f0b)
	}
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	if __obf_015afbee33862a01.__obf_e11087a315c4c191 != nil {
		__obf_015afbee33862a01.__obf_e11087a315c4c191 = append(__obf_015afbee33862a01.__obf_e11087a315c4c191, __obf_015afbee33862a01.__obf_8f3bed0c23610f75...)
	}
	return __obf_015afbee33862a01.__obf_8f3bed0c23610f75, nil
}

func __obf_2c3d4aa0a72cea84(__obf_9b78235c82fd11c0 io.Reader, __obf_9b4a5a04bdad2347 []byte, __obf_56127cd370854f0b int) ([]byte, error) {
	if __obf_9b4a5a04bdad2347 == nil {
		if __obf_56127cd370854f0b == 0 {
			return make([]byte, 0), nil
		}
		__obf_9b4a5a04bdad2347 = make([]byte, 0, __obf_56127cd370854f0b)
	}

	if __obf_56127cd370854f0b > cap(__obf_9b4a5a04bdad2347) {
		__obf_9b4a5a04bdad2347 = append(__obf_9b4a5a04bdad2347, make([]byte, __obf_56127cd370854f0b-len(__obf_9b4a5a04bdad2347))...)
	} else if __obf_56127cd370854f0b <= cap(__obf_9b4a5a04bdad2347) {
		__obf_9b4a5a04bdad2347 = __obf_9b4a5a04bdad2347[:__obf_56127cd370854f0b]
	}

	_, __obf_3aa78ad28f50ed46 := io.ReadFull(__obf_9b78235c82fd11c0, __obf_9b4a5a04bdad2347)
	return __obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46
}

func __obf_470c46bee8a8fb03(__obf_9b78235c82fd11c0 io.Reader, __obf_9b4a5a04bdad2347 []byte, __obf_56127cd370854f0b int) ([]byte, error) {
	if __obf_9b4a5a04bdad2347 == nil {
		if __obf_56127cd370854f0b == 0 {
			return make([]byte, 0), nil
		}
		switch {
		case __obf_56127cd370854f0b < 64:
			__obf_9b4a5a04bdad2347 = make([]byte, 0, 64)
		case __obf_56127cd370854f0b <= __obf_a5e732b826466c45:
			__obf_9b4a5a04bdad2347 = make([]byte, 0, __obf_56127cd370854f0b)
		default:
			__obf_9b4a5a04bdad2347 = make([]byte, 0, __obf_a5e732b826466c45)
		}
	}

	if __obf_56127cd370854f0b <= cap(__obf_9b4a5a04bdad2347) {
		__obf_9b4a5a04bdad2347 = __obf_9b4a5a04bdad2347[:__obf_56127cd370854f0b]
		_, __obf_3aa78ad28f50ed46 := io.ReadFull(__obf_9b78235c82fd11c0, __obf_9b4a5a04bdad2347)
		return __obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46
	}
	__obf_9b4a5a04bdad2347 = __obf_9b4a5a04bdad2347[:cap(__obf_9b4a5a04bdad2347)]

	var __obf_7a0c029b3c407913 int
	for {
		__obf_9dc94311e0067baf := min(__obf_56127cd370854f0b-len(__obf_9b4a5a04bdad2347), __obf_a5e732b826466c45)
		__obf_9b4a5a04bdad2347 = append(__obf_9b4a5a04bdad2347, make([]byte, __obf_9dc94311e0067baf)...)

		_, __obf_3aa78ad28f50ed46 := io.ReadFull(__obf_9b78235c82fd11c0, __obf_9b4a5a04bdad2347[__obf_7a0c029b3c407913:])
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46
		}

		if len(__obf_9b4a5a04bdad2347) == __obf_56127cd370854f0b {
			break
		}
		__obf_7a0c029b3c407913 = len(__obf_9b4a5a04bdad2347)
	}

	return __obf_9b4a5a04bdad2347, nil
}

func min(__obf_c06e3d080de08409, //nolint:unparam
	__obf_9b4a5a04bdad2347 int) int {
	if __obf_c06e3d080de08409 <= __obf_9b4a5a04bdad2347 {
		return __obf_c06e3d080de08409
	}
	return __obf_9b4a5a04bdad2347
}

func __obf_9d816678b5cb9529(__obf_015afbee33862a01 *Decoder, __obf_85d270343a0dfe14 reflect.Value) error {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.bytes(__obf_145c7a7d25eea2bd, __obf_85d270343a0dfe14.Bytes())
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	__obf_85d270343a0dfe14.
		SetBytes(__obf_9b4a5a04bdad2347)
	return nil
}

func (__obf_015afbee33862a01 *Decoder) DecodeBytesLen() (int, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return 0, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
}

func (__obf_015afbee33862a01 *Decoder) DecodeBytes() ([]byte, error) {
	__obf_145c7a7d25eea2bd, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_4c5f8d154c9027ea()
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	return __obf_015afbee33862a01.bytes(__obf_145c7a7d25eea2bd, nil)
}

func (__obf_015afbee33862a01 *Decoder) bytes(__obf_145c7a7d25eea2bd byte, __obf_9b4a5a04bdad2347 []byte) ([]byte, error) {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_196dd6b18db329d3(__obf_145c7a7d25eea2bd)
	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return nil, nil
	}
	return __obf_2c3d4aa0a72cea84(__obf_015afbee33862a01.__obf_9b78235c82fd11c0, __obf_9b4a5a04bdad2347, __obf_56127cd370854f0b)
}
