package __obf_3e0c303610a19bd4

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
	__obf_0eaede5e05901d4f = 1 << 20
	__obf_180374296ff44439 = // 1mb
	1e6
	__obf_94b1e1f52a0a7c90 = // 1m elements
	1e6                      // 1m elements
)

const (
	__obf_fa3605fe60bfee1d uint32 = 1 << iota
	__obf_8807906d7618037e
	__obf_875c65698843f8e6
	__obf_4b9eaa01d630fce4
)

type __obf_19298dd8c5eeed42 interface {
	io.Reader
	io.ByteScanner
}

//------------------------------------------------------------------------------

var __obf_1b7ac9162ebb7e71 = sync.Pool{
	New: func() any {
		return NewDecoder(nil)
	},
}

func GetDecoder() *Decoder {
	return __obf_1b7ac9162ebb7e71.Get().(*Decoder)
}

func PutDecoder(__obf_0ed630deb24db696 *Decoder) {
	__obf_0ed630deb24db696.__obf_b3e2bfc96bdb9204 = nil
	__obf_0ed630deb24db696.__obf_61027e0491b6dd3d = nil
	__obf_1b7ac9162ebb7e71.
		Put(__obf_0ed630deb24db696)
}

//------------------------------------------------------------------------------

// Unmarshal decodes the MessagePack-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(__obf_4d827abff69b9b40 []byte, __obf_63bbcee86d44fdde any) error {
	__obf_0ed630deb24db696 := GetDecoder()
	__obf_0ed630deb24db696.
		UsePreallocateValues(true)
	__obf_0ed630deb24db696.
		Reset(bytes.NewReader(__obf_4d827abff69b9b40))
	__obf_8882f6cf6e378ded := __obf_0ed630deb24db696.Decode(__obf_63bbcee86d44fdde)

	PutDecoder(__obf_0ed630deb24db696)

	return __obf_8882f6cf6e378ded
}

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	__obf_b3e2bfc96bdb9204 io.Reader
	__obf_61027e0491b6dd3d io.ByteScanner
	__obf_7162a9ba642f3ac5 func(*Decoder) (any, error)
	__obf_6d0a0d8a79287f26 string
	__obf_27e4a1a33a31e4a7 []byte
	__obf_4903e7701e6854a2 []byte
	__obf_a22a31b815544cae []string
	__obf_3cf0882fa5a4cafb uint32
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r
// beyond the requested msgpack values. Buffering can be disabled
// by passing a reader that implements io.ByteScanner interface.
func NewDecoder(__obf_b3e2bfc96bdb9204 io.Reader) *Decoder {
	__obf_dc35117108ba8439 := new(Decoder)
	__obf_dc35117108ba8439.
		Reset(__obf_b3e2bfc96bdb9204)
	return __obf_dc35117108ba8439
}

// Reset discards any buffered data, resets all state, and switches the buffered
// reader to read from r.
func (__obf_dc35117108ba8439 *Decoder) Reset(__obf_b3e2bfc96bdb9204 io.Reader) {
	__obf_dc35117108ba8439.
		ResetDict(__obf_b3e2bfc96bdb9204, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_dc35117108ba8439 *Decoder) ResetDict(__obf_b3e2bfc96bdb9204 io.Reader, __obf_a22a31b815544cae []string) {
	__obf_dc35117108ba8439.
		ResetReader(__obf_b3e2bfc96bdb9204)
	__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb = 0
	__obf_dc35117108ba8439.__obf_6d0a0d8a79287f26 = ""
	__obf_dc35117108ba8439.__obf_a22a31b815544cae = __obf_a22a31b815544cae
}

func (__obf_dc35117108ba8439 *Decoder) WithDict(__obf_a22a31b815544cae []string, __obf_6ff63d98a176dcfe func(*Decoder) error) error {
	__obf_d07f84e9e0444e28 := __obf_dc35117108ba8439.__obf_a22a31b815544cae
	__obf_dc35117108ba8439.__obf_a22a31b815544cae = __obf_a22a31b815544cae
	__obf_8882f6cf6e378ded := __obf_6ff63d98a176dcfe(__obf_dc35117108ba8439)
	__obf_dc35117108ba8439.__obf_a22a31b815544cae = __obf_d07f84e9e0444e28
	return __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) ResetReader(__obf_b3e2bfc96bdb9204 io.Reader) {
	__obf_dc35117108ba8439.__obf_7162a9ba642f3ac5 = nil
	__obf_dc35117108ba8439.__obf_a22a31b815544cae = nil

	if __obf_a06a89d585706b15, __obf_ce8bef141eff8aab := __obf_b3e2bfc96bdb9204.(__obf_19298dd8c5eeed42); __obf_ce8bef141eff8aab {
		__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204 = __obf_a06a89d585706b15
		__obf_dc35117108ba8439.__obf_61027e0491b6dd3d = __obf_a06a89d585706b15
	} else if __obf_b3e2bfc96bdb9204 == nil {
		__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204 = nil
		__obf_dc35117108ba8439.__obf_61027e0491b6dd3d = nil
	} else {
		__obf_a06a89d585706b15 := bufio.NewReader(__obf_b3e2bfc96bdb9204)
		__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204 = __obf_a06a89d585706b15
		__obf_dc35117108ba8439.__obf_61027e0491b6dd3d = __obf_a06a89d585706b15
	}
}

func (__obf_dc35117108ba8439 *Decoder) SetMapDecoder(__obf_6ff63d98a176dcfe func(*Decoder) (any, error)) {
	__obf_dc35117108ba8439.__obf_7162a9ba642f3ac5 = __obf_6ff63d98a176dcfe
}

// UseLooseInterfaceDecoding causes decoder to use DecodeInterfaceLoose
// to decode msgpack value into Go some.
func (__obf_dc35117108ba8439 *Decoder) UseLooseInterfaceDecoding(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb |= __obf_fa3605fe60bfee1d
	} else {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb &= ^__obf_fa3605fe60bfee1d
	}
}

// SetCustomStructTag causes the decoder to use the supplied tag as a fallback option
// if there is no msgpack tag.
func (__obf_dc35117108ba8439 *Decoder) SetCustomStructTag(__obf_42ff6f144202733b string) {
	__obf_dc35117108ba8439.__obf_6d0a0d8a79287f26 = __obf_42ff6f144202733b
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (__obf_dc35117108ba8439 *Decoder) DisallowUnknownFields(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb |= __obf_8807906d7618037e
	} else {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb &= ^__obf_8807906d7618037e
	}
}

// UseInternedStrings enables support for decoding interned strings.
func (__obf_dc35117108ba8439 *Decoder) UseInternedStrings(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb |= __obf_a063db36bdbf4202
	} else {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb &= ^__obf_a063db36bdbf4202
	}
}

// UsePreallocateValues enables preallocating values in chunks
func (__obf_dc35117108ba8439 *Decoder) UsePreallocateValues(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb |= __obf_875c65698843f8e6
	} else {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb &= ^__obf_875c65698843f8e6
	}
}

// DisableAllocLimit enables fully allocating slices/maps when the size is known
func (__obf_dc35117108ba8439 *Decoder) DisableAllocLimit(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb |= __obf_4b9eaa01d630fce4
	} else {
		__obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb &= ^__obf_4b9eaa01d630fce4
	}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (__obf_dc35117108ba8439 *Decoder) Buffered() io.Reader {
	return __obf_dc35117108ba8439.

		//nolint:gocyclo
		__obf_b3e2bfc96bdb9204
}

func (__obf_dc35117108ba8439 *Decoder) Decode(__obf_63bbcee86d44fdde any) error {
	var __obf_8882f6cf6e378ded error
	switch __obf_63bbcee86d44fdde := __obf_63bbcee86d44fdde.(type) {
	case *string:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeString()
			return __obf_8882f6cf6e378ded
		}
	case *[]byte:
		if __obf_63bbcee86d44fdde != nil {
			return __obf_dc35117108ba8439.__obf_85900a462858a70d(__obf_63bbcee86d44fdde)
		}
	case *int:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeInt()
			return __obf_8882f6cf6e378ded
		}
	case *int8:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeInt8()
			return __obf_8882f6cf6e378ded
		}
	case *int16:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeInt16()
			return __obf_8882f6cf6e378ded
		}
	case *int32:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeInt32()
			return __obf_8882f6cf6e378ded
		}
	case *int64:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeInt64()
			return __obf_8882f6cf6e378ded
		}
	case *uint:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeUint()
			return __obf_8882f6cf6e378ded
		}
	case *uint8:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeUint8()
			return __obf_8882f6cf6e378ded
		}
	case *uint16:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeUint16()
			return __obf_8882f6cf6e378ded
		}
	case *uint32:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeUint32()
			return __obf_8882f6cf6e378ded
		}
	case *uint64:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeUint64()
			return __obf_8882f6cf6e378ded
		}
	case *bool:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeBool()
			return __obf_8882f6cf6e378ded
		}
	case *float32:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeFloat32()
			return __obf_8882f6cf6e378ded
		}
	case *float64:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeFloat64()
			return __obf_8882f6cf6e378ded
		}
	case *[]string:
		return __obf_dc35117108ba8439.__obf_bdc4d45b04bc0799(__obf_63bbcee86d44fdde)
	case *map[string]string:
		return __obf_dc35117108ba8439.__obf_58af590e339dfa38(__obf_63bbcee86d44fdde)
	case *map[string]any:
		return __obf_dc35117108ba8439.__obf_85426e39f3ed0878(__obf_63bbcee86d44fdde)
	case *time.Duration:
		if __obf_63bbcee86d44fdde != nil {
			__obf_785552c2ee57ae56, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
			*__obf_63bbcee86d44fdde = time.Duration(__obf_785552c2ee57ae56)
			return __obf_8882f6cf6e378ded
		}
	case *time.Time:
		if __obf_63bbcee86d44fdde != nil {
			*__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded = __obf_dc35117108ba8439.DecodeTime()
			return __obf_8882f6cf6e378ded
		}
	}
	__obf_785552c2ee57ae56 := reflect.ValueOf(__obf_63bbcee86d44fdde)
	if !__obf_785552c2ee57ae56.IsValid() {
		return errors.New("msgpack: Decode(nil)")
	}
	if __obf_785552c2ee57ae56.Kind() != reflect.Ptr {
		return fmt.Errorf("msgpack: Decode(non-pointer %T)", __obf_63bbcee86d44fdde)
	}
	if __obf_785552c2ee57ae56.IsNil() {
		return fmt.Errorf("msgpack: Decode(non-settable %T)", __obf_63bbcee86d44fdde)
	}
	__obf_785552c2ee57ae56 = __obf_785552c2ee57ae56.Elem()
	if __obf_785552c2ee57ae56.Kind() == reflect.Interface {
		if !__obf_785552c2ee57ae56.IsNil() {
			__obf_785552c2ee57ae56 = __obf_785552c2ee57ae56.Elem()
			if __obf_785552c2ee57ae56.Kind() != reflect.Ptr {
				return fmt.Errorf("msgpack: Decode(non-pointer %s)", __obf_785552c2ee57ae56.Type().String())
			}
		}
	}

	return __obf_dc35117108ba8439.DecodeValue(__obf_785552c2ee57ae56)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeMulti(__obf_63bbcee86d44fdde ...any) error {
	for _, __obf_785552c2ee57ae56 := range __obf_63bbcee86d44fdde {
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Decode(__obf_785552c2ee57ae56); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_93ff057c8a5c795c() (any, error) {
	if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_fa3605fe60bfee1d != 0 {
		return __obf_dc35117108ba8439.DecodeInterfaceLoose()
	}
	return __obf_dc35117108ba8439.DecodeInterface()
}

func (__obf_dc35117108ba8439 *Decoder) DecodeValue(__obf_63bbcee86d44fdde reflect.Value) error {
	__obf_bcda9de57dfcca05 := __obf_9b0fb8e0beb3ab15(__obf_63bbcee86d44fdde.Type())
	return __obf_bcda9de57dfcca05(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeNil() error {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_e46289218af709bf != Nil {
		return fmt.Errorf("msgpack: invalid code=%x decoding nil", __obf_e46289218af709bf)
	}
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_f8b6b5cb00b4ac2c(__obf_63bbcee86d44fdde reflect.Value) error {
	__obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeNil()
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_8882f6cf6e378ded
	}
	if __obf_63bbcee86d44fdde.Kind() == reflect.Ptr {
		__obf_63bbcee86d44fdde = __obf_63bbcee86d44fdde.Elem()
	}
	__obf_63bbcee86d44fdde.
		Set(reflect.Zero(__obf_63bbcee86d44fdde.Type()))
	return __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) DecodeBool() (bool, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return false, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.bool(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) bool(__obf_e46289218af709bf byte) (bool, error) {
	if __obf_e46289218af709bf == Nil {
		return false, nil
	}
	if __obf_e46289218af709bf == False {
		return false, nil
	}
	if __obf_e46289218af709bf == True {
		return true, nil
	}
	return false, fmt.Errorf("msgpack: invalid code=%x decoding bool", __obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeDuration() (time.Duration, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeInt64()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return time.Duration(__obf_4909ae60ffbb8e53), nil
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
func (__obf_dc35117108ba8439 *Decoder) DecodeInterface() (any, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}

	if IsFixedNum(__obf_e46289218af709bf) {
		return int8(__obf_e46289218af709bf), nil
	}
	if IsFixedMap(__obf_e46289218af709bf) {
		__obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.UnreadByte()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_a2dd6e1a3453d2c3()
	}
	if IsFixedArray(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.__obf_562eb479f437208e(__obf_e46289218af709bf)
	}
	if IsFixedString(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.string(__obf_e46289218af709bf)
	}

	switch __obf_e46289218af709bf {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_dc35117108ba8439.bool(__obf_e46289218af709bf)
	case Float:
		return __obf_dc35117108ba8439.float32(__obf_e46289218af709bf)
	case Double:
		return __obf_dc35117108ba8439.float64(__obf_e46289218af709bf)
	case Uint8:
		return __obf_dc35117108ba8439.uint8()
	case Uint16:
		return __obf_dc35117108ba8439.uint16()
	case Uint32:
		return __obf_dc35117108ba8439.uint32()
	case Uint64:
		return __obf_dc35117108ba8439.uint64()
	case Int8:
		return __obf_dc35117108ba8439.int8()
	case Int16:
		return __obf_dc35117108ba8439.int16()
	case Int32:
		return __obf_dc35117108ba8439.int32()
	case Int64:
		return __obf_dc35117108ba8439.int64()
	case Bin8, Bin16, Bin32:
		return __obf_dc35117108ba8439.bytes(__obf_e46289218af709bf, nil)
	case Str8, Str16, Str32:
		return __obf_dc35117108ba8439.string(__obf_e46289218af709bf)
	case Array16, Array32:
		return __obf_dc35117108ba8439.__obf_562eb479f437208e(__obf_e46289218af709bf)
	case Map16, Map32:
		__obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.UnreadByte()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_a2dd6e1a3453d2c3()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_dc35117108ba8439.__obf_5b9b53ab12d1dfc9(__obf_e46289218af709bf)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_e46289218af709bf)
}

// DecodeInterfaceLoose is like DecodeInterface except that:
//   - int8, int16, and int32 are converted to int64,
//   - uint8, uint16, and uint32 are converted to uint64,
//   - float32 is converted to float64.
//   - []byte is converted to string.
func (__obf_dc35117108ba8439 *Decoder) DecodeInterfaceLoose() (any, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}

	if IsFixedNum(__obf_e46289218af709bf) {
		return int64(int8(__obf_e46289218af709bf)), nil
	}
	if IsFixedMap(__obf_e46289218af709bf) {
		__obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.UnreadByte()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_a2dd6e1a3453d2c3()
	}
	if IsFixedArray(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.__obf_562eb479f437208e(__obf_e46289218af709bf)
	}
	if IsFixedString(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.string(__obf_e46289218af709bf)
	}

	switch __obf_e46289218af709bf {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_dc35117108ba8439.bool(__obf_e46289218af709bf)
	case Float, Double:
		return __obf_dc35117108ba8439.float64(__obf_e46289218af709bf)
	case Uint8, Uint16, Uint32, Uint64:
		return __obf_dc35117108ba8439.uint(__obf_e46289218af709bf)
	case Int8, Int16, Int32, Int64:
		return __obf_dc35117108ba8439.int(__obf_e46289218af709bf)
	case Str8, Str16, Str32,
		Bin8, Bin16, Bin32:
		return __obf_dc35117108ba8439.string(__obf_e46289218af709bf)
	case Array16, Array32:
		return __obf_dc35117108ba8439.__obf_562eb479f437208e(__obf_e46289218af709bf)
	case Map16, Map32:
		__obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.UnreadByte()
		if __obf_8882f6cf6e378ded != nil {
			return nil, __obf_8882f6cf6e378ded
		}
		return __obf_dc35117108ba8439.__obf_a2dd6e1a3453d2c3()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_dc35117108ba8439.__obf_5b9b53ab12d1dfc9(__obf_e46289218af709bf)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_e46289218af709bf)
}

// Skip skips next value.
func (__obf_dc35117108ba8439 *Decoder) Skip() error {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	if IsFixedNum(__obf_e46289218af709bf) {
		return nil
	}
	if IsFixedMap(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.__obf_ab578c2c2c260fb6(__obf_e46289218af709bf)
	}
	if IsFixedArray(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.__obf_3d61f79527fa3520(__obf_e46289218af709bf)
	}
	if IsFixedString(__obf_e46289218af709bf) {
		return __obf_dc35117108ba8439.__obf_a8d3bdf080f01438(__obf_e46289218af709bf)
	}

	switch __obf_e46289218af709bf {
	case Nil, False, True:
		return nil
	case Uint8, Int8:
		return __obf_dc35117108ba8439.__obf_19f294acbba68c47(1)
	case Uint16, Int16:
		return __obf_dc35117108ba8439.__obf_19f294acbba68c47(2)
	case Uint32, Int32, Float:
		return __obf_dc35117108ba8439.__obf_19f294acbba68c47(4)
	case Uint64, Int64, Double:
		return __obf_dc35117108ba8439.__obf_19f294acbba68c47(8)
	case Bin8, Bin16, Bin32:
		return __obf_dc35117108ba8439.__obf_a8d3bdf080f01438(__obf_e46289218af709bf)
	case Str8, Str16, Str32:
		return __obf_dc35117108ba8439.__obf_a8d3bdf080f01438(__obf_e46289218af709bf)
	case Array16, Array32:
		return __obf_dc35117108ba8439.__obf_3d61f79527fa3520(__obf_e46289218af709bf)
	case Map16, Map32:
		return __obf_dc35117108ba8439.__obf_ab578c2c2c260fb6(__obf_e46289218af709bf)
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_dc35117108ba8439.__obf_f260dda48c6603bb(__obf_e46289218af709bf)
	}

	return fmt.Errorf("msgpack: unknown code %x", __obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeRaw() (RawMessage, error) {
	__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = make([]byte, 0)
	if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	__obf_ea75b6610bb8e3c5 := RawMessage(__obf_dc35117108ba8439.__obf_4903e7701e6854a2)
	__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = nil
	return __obf_ea75b6610bb8e3c5, nil
}

// PeekCode returns the next MessagePack code without advancing the reader.
// Subpackage msgpack/codes defines the list of available
func (__obf_dc35117108ba8439 *Decoder) PeekCode() (byte, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.ReadByte()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_e46289218af709bf, __obf_dc35117108ba8439.

		// ReadFull reads exactly len(buf) bytes into the buf.
		__obf_61027e0491b6dd3d.UnreadByte()
}

func (__obf_dc35117108ba8439 *Decoder) ReadFull(__obf_27e4a1a33a31e4a7 []byte) error {
	_, __obf_8882f6cf6e378ded := __obf_b06a4f273442ca29(__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204, __obf_27e4a1a33a31e4a7, len(__obf_27e4a1a33a31e4a7))
	return __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) __obf_56c3d1b3779316ba() bool {
	__obf_545372fefbb733e5, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.PeekCode()
	return __obf_8882f6cf6e378ded == nil && __obf_545372fefbb733e5 == Nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_5b7d85f9093902c3() (byte, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_61027e0491b6dd3d.ReadByte()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	if __obf_dc35117108ba8439.__obf_4903e7701e6854a2 != nil {
		__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = append(__obf_dc35117108ba8439.__obf_4903e7701e6854a2, __obf_e46289218af709bf)
	}
	return __obf_e46289218af709bf, nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_ac68fb7f5c42aa1d(__obf_11bcc66cde095c11 []byte) error {
	_, __obf_8882f6cf6e378ded := io.ReadFull(__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204, __obf_11bcc66cde095c11)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_dc35117108ba8439.__obf_4903e7701e6854a2 != nil {
		__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = append(__obf_dc35117108ba8439.__obf_4903e7701e6854a2, __obf_11bcc66cde095c11...)
	}
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_b06a4f273442ca29(__obf_4909ae60ffbb8e53 int) ([]byte, error) {
	var __obf_8882f6cf6e378ded error
	if __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_4b9eaa01d630fce4 != 0 {
		__obf_dc35117108ba8439.__obf_27e4a1a33a31e4a7, __obf_8882f6cf6e378ded = __obf_b06a4f273442ca29(__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204, __obf_dc35117108ba8439.__obf_27e4a1a33a31e4a7, __obf_4909ae60ffbb8e53)
	} else {
		__obf_dc35117108ba8439.__obf_27e4a1a33a31e4a7, __obf_8882f6cf6e378ded = __obf_577041cfd106e1f3(__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204, __obf_dc35117108ba8439.__obf_27e4a1a33a31e4a7,

			// TODO: read directly into d.rec?
			__obf_4909ae60ffbb8e53)
	}
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	if __obf_dc35117108ba8439.__obf_4903e7701e6854a2 != nil {
		__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = append(__obf_dc35117108ba8439.__obf_4903e7701e6854a2, __obf_dc35117108ba8439.__obf_27e4a1a33a31e4a7...)
	}
	return __obf_dc35117108ba8439.__obf_27e4a1a33a31e4a7, nil
}

func __obf_b06a4f273442ca29(__obf_b3e2bfc96bdb9204 io.Reader, __obf_11bcc66cde095c11 []byte, __obf_4909ae60ffbb8e53 int) ([]byte, error) {
	if __obf_11bcc66cde095c11 == nil {
		if __obf_4909ae60ffbb8e53 == 0 {
			return make([]byte, 0), nil
		}
		__obf_11bcc66cde095c11 = make([]byte, 0, __obf_4909ae60ffbb8e53)
	}

	if __obf_4909ae60ffbb8e53 > cap(__obf_11bcc66cde095c11) {
		__obf_11bcc66cde095c11 = append(__obf_11bcc66cde095c11, make([]byte, __obf_4909ae60ffbb8e53-len(__obf_11bcc66cde095c11))...)
	} else if __obf_4909ae60ffbb8e53 <= cap(__obf_11bcc66cde095c11) {
		__obf_11bcc66cde095c11 = __obf_11bcc66cde095c11[:__obf_4909ae60ffbb8e53]
	}

	_, __obf_8882f6cf6e378ded := io.ReadFull(__obf_b3e2bfc96bdb9204, __obf_11bcc66cde095c11)
	return __obf_11bcc66cde095c11, __obf_8882f6cf6e378ded
}

func __obf_577041cfd106e1f3(__obf_b3e2bfc96bdb9204 io.Reader, __obf_11bcc66cde095c11 []byte, __obf_4909ae60ffbb8e53 int) ([]byte, error) {
	if __obf_11bcc66cde095c11 == nil {
		if __obf_4909ae60ffbb8e53 == 0 {
			return make([]byte, 0), nil
		}
		switch {
		case __obf_4909ae60ffbb8e53 < 64:
			__obf_11bcc66cde095c11 = make([]byte, 0, 64)
		case __obf_4909ae60ffbb8e53 <= __obf_0eaede5e05901d4f:
			__obf_11bcc66cde095c11 = make([]byte, 0, __obf_4909ae60ffbb8e53)
		default:
			__obf_11bcc66cde095c11 = make([]byte, 0, __obf_0eaede5e05901d4f)
		}
	}

	if __obf_4909ae60ffbb8e53 <= cap(__obf_11bcc66cde095c11) {
		__obf_11bcc66cde095c11 = __obf_11bcc66cde095c11[:__obf_4909ae60ffbb8e53]
		_, __obf_8882f6cf6e378ded := io.ReadFull(__obf_b3e2bfc96bdb9204, __obf_11bcc66cde095c11)
		return __obf_11bcc66cde095c11, __obf_8882f6cf6e378ded
	}
	__obf_11bcc66cde095c11 = __obf_11bcc66cde095c11[:cap(__obf_11bcc66cde095c11)]

	var __obf_956969735bba2920 int
	for {
		__obf_c7276befb1a7616c := min(__obf_4909ae60ffbb8e53-len(__obf_11bcc66cde095c11), __obf_0eaede5e05901d4f)
		__obf_11bcc66cde095c11 = append(__obf_11bcc66cde095c11, make([]byte, __obf_c7276befb1a7616c)...)

		_, __obf_8882f6cf6e378ded := io.ReadFull(__obf_b3e2bfc96bdb9204, __obf_11bcc66cde095c11[__obf_956969735bba2920:])
		if __obf_8882f6cf6e378ded != nil {
			return __obf_11bcc66cde095c11, __obf_8882f6cf6e378ded
		}

		if len(__obf_11bcc66cde095c11) == __obf_4909ae60ffbb8e53 {
			break
		}
		__obf_956969735bba2920 = len(__obf_11bcc66cde095c11)
	}

	return __obf_11bcc66cde095c11, nil
}

func min(__obf_e3fb07f0dd6df716, //nolint:unparam
	__obf_11bcc66cde095c11 int) int {
	if __obf_e3fb07f0dd6df716 <= __obf_11bcc66cde095c11 {
		return __obf_e3fb07f0dd6df716
	}
	return __obf_11bcc66cde095c11
}

func __obf_79dff4d84a62412e(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.bytes(__obf_e46289218af709bf, __obf_63bbcee86d44fdde.Bytes())
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetBytes(__obf_11bcc66cde095c11)
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) DecodeBytesLen() (int, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeBytes() ([]byte, error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.bytes(__obf_e46289218af709bf, nil)
}

func (__obf_dc35117108ba8439 *Decoder) bytes(__obf_e46289218af709bf byte, __obf_11bcc66cde095c11 []byte) ([]byte, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return nil, nil
	}
	return __obf_b06a4f273442ca29(__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204, __obf_11bcc66cde095c11, __obf_4909ae60ffbb8e53)
}
