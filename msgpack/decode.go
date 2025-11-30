package __obf_fd770cb9675903b0

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
	__obf_6bc1f5097e067b3d = 1 << 20
	__obf_369a9acbbd312147 = // 1mb
	1e6
	__obf_d6d48ab082cf77ea = // 1m elements
	1e6                      // 1m elements
)

const (
	__obf_6bae37903e5f5d4a uint32 = 1 << iota
	__obf_5726d2683a07c9c6
	__obf_759b68f674456e1c
	__obf_d0b80752c6a68c61
)

type __obf_f445e62ad4df68a4 interface {
	io.Reader
	io.ByteScanner
}

//------------------------------------------------------------------------------

var __obf_ac0a5db99a70b7bd = sync.Pool{
	New: func() any {
		return NewDecoder(nil)
	},
}

func GetDecoder() *Decoder {
	return __obf_ac0a5db99a70b7bd.Get().(*Decoder)
}

func PutDecoder(__obf_2f8f01810a02d049 *Decoder) {
	__obf_2f8f01810a02d049.__obf_52c07f7b574f636a = nil
	__obf_2f8f01810a02d049.__obf_fe1ace7a2eb8bf9f = nil
	__obf_ac0a5db99a70b7bd.
		Put(__obf_2f8f01810a02d049)
}

//------------------------------------------------------------------------------

// Unmarshal decodes the MessagePack-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(__obf_cc76e8ed73142f28 []byte, __obf_f328a048f76b7256 any) error {
	__obf_2f8f01810a02d049 := GetDecoder()
	__obf_2f8f01810a02d049.
		UsePreallocateValues(true)
	__obf_2f8f01810a02d049.
		Reset(bytes.NewReader(__obf_cc76e8ed73142f28))
	__obf_45342a3a754d12f5 := __obf_2f8f01810a02d049.Decode(__obf_f328a048f76b7256)

	PutDecoder(__obf_2f8f01810a02d049)

	return __obf_45342a3a754d12f5
}

// A Decoder reads and decodes MessagePack values from an input stream.
type Decoder struct {
	__obf_52c07f7b574f636a io.Reader
	__obf_fe1ace7a2eb8bf9f io.ByteScanner
	__obf_a212983b973aa252 func(*Decoder) (any, error)
	__obf_6621ba3773b8696a string
	__obf_099b1f8c6882e66a []byte
	__obf_98388ae8c389d758 []byte
	__obf_ff96db22e12b6842 []string
	__obf_7185cb62de7af792 uint32
}

// NewDecoder returns a new decoder that reads from r.
//
// The decoder introduces its own buffering and may read data from r
// beyond the requested msgpack values. Buffering can be disabled
// by passing a reader that implements io.ByteScanner interface.
func NewDecoder(__obf_52c07f7b574f636a io.Reader) *Decoder {
	__obf_5d389b660eefb08c := new(Decoder)
	__obf_5d389b660eefb08c.
		Reset(__obf_52c07f7b574f636a)
	return __obf_5d389b660eefb08c
}

// Reset discards any buffered data, resets all state, and switches the buffered
// reader to read from r.
func (__obf_5d389b660eefb08c *Decoder) Reset(__obf_52c07f7b574f636a io.Reader) {
	__obf_5d389b660eefb08c.
		ResetDict(__obf_52c07f7b574f636a, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_5d389b660eefb08c *Decoder) ResetDict(__obf_52c07f7b574f636a io.Reader, __obf_ff96db22e12b6842 []string) {
	__obf_5d389b660eefb08c.
		ResetReader(__obf_52c07f7b574f636a)
	__obf_5d389b660eefb08c.__obf_7185cb62de7af792 = 0
	__obf_5d389b660eefb08c.__obf_6621ba3773b8696a = ""
	__obf_5d389b660eefb08c.__obf_ff96db22e12b6842 = __obf_ff96db22e12b6842
}

func (__obf_5d389b660eefb08c *Decoder) WithDict(__obf_ff96db22e12b6842 []string, __obf_1e4576e8508b04bc func(*Decoder) error) error {
	__obf_cf0262a7ff1775c5 := __obf_5d389b660eefb08c.__obf_ff96db22e12b6842
	__obf_5d389b660eefb08c.__obf_ff96db22e12b6842 = __obf_ff96db22e12b6842
	__obf_45342a3a754d12f5 := __obf_1e4576e8508b04bc(__obf_5d389b660eefb08c)
	__obf_5d389b660eefb08c.__obf_ff96db22e12b6842 = __obf_cf0262a7ff1775c5
	return __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) ResetReader(__obf_52c07f7b574f636a io.Reader) {
	__obf_5d389b660eefb08c.__obf_a212983b973aa252 = nil
	__obf_5d389b660eefb08c.__obf_ff96db22e12b6842 = nil

	if __obf_62830076477d32ce, __obf_b00b3c6a10f90467 := __obf_52c07f7b574f636a.(__obf_f445e62ad4df68a4); __obf_b00b3c6a10f90467 {
		__obf_5d389b660eefb08c.__obf_52c07f7b574f636a = __obf_62830076477d32ce
		__obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f = __obf_62830076477d32ce
	} else if __obf_52c07f7b574f636a == nil {
		__obf_5d389b660eefb08c.__obf_52c07f7b574f636a = nil
		__obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f = nil
	} else {
		__obf_62830076477d32ce := bufio.NewReader(__obf_52c07f7b574f636a)
		__obf_5d389b660eefb08c.__obf_52c07f7b574f636a = __obf_62830076477d32ce
		__obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f = __obf_62830076477d32ce
	}
}

func (__obf_5d389b660eefb08c *Decoder) SetMapDecoder(__obf_1e4576e8508b04bc func(*Decoder) (any, error)) {
	__obf_5d389b660eefb08c.__obf_a212983b973aa252 = __obf_1e4576e8508b04bc
}

// UseLooseInterfaceDecoding causes decoder to use DecodeInterfaceLoose
// to decode msgpack value into Go some.
func (__obf_5d389b660eefb08c *Decoder) UseLooseInterfaceDecoding(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 |= __obf_6bae37903e5f5d4a
	} else {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 &= ^__obf_6bae37903e5f5d4a
	}
}

// SetCustomStructTag causes the decoder to use the supplied tag as a fallback option
// if there is no msgpack tag.
func (__obf_5d389b660eefb08c *Decoder) SetCustomStructTag(__obf_1a0f202964305730 string) {
	__obf_5d389b660eefb08c.__obf_6621ba3773b8696a = __obf_1a0f202964305730
}

// DisallowUnknownFields causes the Decoder to return an error when the destination
// is a struct and the input contains object keys which do not match any
// non-ignored, exported fields in the destination.
func (__obf_5d389b660eefb08c *Decoder) DisallowUnknownFields(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 |= __obf_5726d2683a07c9c6
	} else {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 &= ^__obf_5726d2683a07c9c6
	}
}

// UseInternedStrings enables support for decoding interned strings.
func (__obf_5d389b660eefb08c *Decoder) UseInternedStrings(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 |= __obf_139c02350e62bdf7
	} else {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 &= ^__obf_139c02350e62bdf7
	}
}

// UsePreallocateValues enables preallocating values in chunks
func (__obf_5d389b660eefb08c *Decoder) UsePreallocateValues(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 |= __obf_759b68f674456e1c
	} else {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 &= ^__obf_759b68f674456e1c
	}
}

// DisableAllocLimit enables fully allocating slices/maps when the size is known
func (__obf_5d389b660eefb08c *Decoder) DisableAllocLimit(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 |= __obf_d0b80752c6a68c61
	} else {
		__obf_5d389b660eefb08c.__obf_7185cb62de7af792 &= ^__obf_d0b80752c6a68c61
	}
}

// Buffered returns a reader of the data remaining in the Decoder's buffer.
// The reader is valid until the next call to Decode.
func (__obf_5d389b660eefb08c *Decoder) Buffered() io.Reader {
	return __obf_5d389b660eefb08c.

		//nolint:gocyclo
		__obf_52c07f7b574f636a
}

func (__obf_5d389b660eefb08c *Decoder) Decode(__obf_f328a048f76b7256 any) error {
	var __obf_45342a3a754d12f5 error
	switch __obf_f328a048f76b7256 := __obf_f328a048f76b7256.(type) {
	case *string:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeString()
			return __obf_45342a3a754d12f5
		}
	case *[]byte:
		if __obf_f328a048f76b7256 != nil {
			return __obf_5d389b660eefb08c.__obf_f4b9f1a4ece9fd91(__obf_f328a048f76b7256)
		}
	case *int:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeInt()
			return __obf_45342a3a754d12f5
		}
	case *int8:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeInt8()
			return __obf_45342a3a754d12f5
		}
	case *int16:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeInt16()
			return __obf_45342a3a754d12f5
		}
	case *int32:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeInt32()
			return __obf_45342a3a754d12f5
		}
	case *int64:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeInt64()
			return __obf_45342a3a754d12f5
		}
	case *uint:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeUint()
			return __obf_45342a3a754d12f5
		}
	case *uint8:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeUint8()
			return __obf_45342a3a754d12f5
		}
	case *uint16:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeUint16()
			return __obf_45342a3a754d12f5
		}
	case *uint32:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeUint32()
			return __obf_45342a3a754d12f5
		}
	case *uint64:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeUint64()
			return __obf_45342a3a754d12f5
		}
	case *bool:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeBool()
			return __obf_45342a3a754d12f5
		}
	case *float32:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeFloat32()
			return __obf_45342a3a754d12f5
		}
	case *float64:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeFloat64()
			return __obf_45342a3a754d12f5
		}
	case *[]string:
		return __obf_5d389b660eefb08c.__obf_e4607e36ff9c01b4(__obf_f328a048f76b7256)
	case *map[string]string:
		return __obf_5d389b660eefb08c.__obf_4acef8852f5f3f0b(__obf_f328a048f76b7256)
	case *map[string]any:
		return __obf_5d389b660eefb08c.__obf_1bc77a41576c2d43(__obf_f328a048f76b7256)
	case *time.Duration:
		if __obf_f328a048f76b7256 != nil {
			__obf_e2c885d3da27fd46, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
			*__obf_f328a048f76b7256 = time.Duration(__obf_e2c885d3da27fd46)
			return __obf_45342a3a754d12f5
		}
	case *time.Time:
		if __obf_f328a048f76b7256 != nil {
			*__obf_f328a048f76b7256, __obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.DecodeTime()
			return __obf_45342a3a754d12f5
		}
	}
	__obf_e2c885d3da27fd46 := reflect.ValueOf(__obf_f328a048f76b7256)
	if !__obf_e2c885d3da27fd46.IsValid() {
		return errors.New("msgpack: Decode(nil)")
	}
	if __obf_e2c885d3da27fd46.Kind() != reflect.Ptr {
		return fmt.Errorf("msgpack: Decode(non-pointer %T)", __obf_f328a048f76b7256)
	}
	if __obf_e2c885d3da27fd46.IsNil() {
		return fmt.Errorf("msgpack: Decode(non-settable %T)", __obf_f328a048f76b7256)
	}
	__obf_e2c885d3da27fd46 = __obf_e2c885d3da27fd46.Elem()
	if __obf_e2c885d3da27fd46.Kind() == reflect.Interface {
		if !__obf_e2c885d3da27fd46.IsNil() {
			__obf_e2c885d3da27fd46 = __obf_e2c885d3da27fd46.Elem()
			if __obf_e2c885d3da27fd46.Kind() != reflect.Ptr {
				return fmt.Errorf("msgpack: Decode(non-pointer %s)", __obf_e2c885d3da27fd46.Type().String())
			}
		}
	}

	return __obf_5d389b660eefb08c.DecodeValue(__obf_e2c885d3da27fd46)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeMulti(__obf_f328a048f76b7256 ...any) error {
	for _, __obf_e2c885d3da27fd46 := range __obf_f328a048f76b7256 {
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Decode(__obf_e2c885d3da27fd46); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_d8bd92f8abc7c84f() (any, error) {
	if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_6bae37903e5f5d4a != 0 {
		return __obf_5d389b660eefb08c.DecodeInterfaceLoose()
	}
	return __obf_5d389b660eefb08c.DecodeInterface()
}

func (__obf_5d389b660eefb08c *Decoder) DecodeValue(__obf_f328a048f76b7256 reflect.Value) error {
	__obf_1b210fa2affbfd6d := __obf_538e4e278f06c9d6(__obf_f328a048f76b7256.Type())
	return __obf_1b210fa2affbfd6d(__obf_5d389b660eefb08c, __obf_f328a048f76b7256)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeNil() error {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_4148125b350dfea2 != Nil {
		return fmt.Errorf("msgpack: invalid code=%x decoding nil", __obf_4148125b350dfea2)
	}
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_4e408f1c6737e5d3(__obf_f328a048f76b7256 reflect.Value) error {
	__obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeNil()
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_45342a3a754d12f5
	}
	if __obf_f328a048f76b7256.Kind() == reflect.Ptr {
		__obf_f328a048f76b7256 = __obf_f328a048f76b7256.Elem()
	}
	__obf_f328a048f76b7256.
		Set(reflect.Zero(__obf_f328a048f76b7256.Type()))
	return __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) DecodeBool() (bool, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return false, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.bool(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) bool(__obf_4148125b350dfea2 byte) (bool, error) {
	if __obf_4148125b350dfea2 == Nil {
		return false, nil
	}
	if __obf_4148125b350dfea2 == False {
		return false, nil
	}
	if __obf_4148125b350dfea2 == True {
		return true, nil
	}
	return false, fmt.Errorf("msgpack: invalid code=%x decoding bool", __obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeDuration() (time.Duration, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeInt64()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return time.Duration(__obf_d774e4844c74bfe9), nil
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
func (__obf_5d389b660eefb08c *Decoder) DecodeInterface() (any, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}

	if IsFixedNum(__obf_4148125b350dfea2) {
		return int8(__obf_4148125b350dfea2), nil
	}
	if IsFixedMap(__obf_4148125b350dfea2) {
		__obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.UnreadByte()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_ee0b22a7a221cee0()
	}
	if IsFixedArray(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.__obf_e557eb76c54a6f47(__obf_4148125b350dfea2)
	}
	if IsFixedString(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.string(__obf_4148125b350dfea2)
	}

	switch __obf_4148125b350dfea2 {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_5d389b660eefb08c.bool(__obf_4148125b350dfea2)
	case Float:
		return __obf_5d389b660eefb08c.float32(__obf_4148125b350dfea2)
	case Double:
		return __obf_5d389b660eefb08c.float64(__obf_4148125b350dfea2)
	case Uint8:
		return __obf_5d389b660eefb08c.uint8()
	case Uint16:
		return __obf_5d389b660eefb08c.uint16()
	case Uint32:
		return __obf_5d389b660eefb08c.uint32()
	case Uint64:
		return __obf_5d389b660eefb08c.uint64()
	case Int8:
		return __obf_5d389b660eefb08c.int8()
	case Int16:
		return __obf_5d389b660eefb08c.int16()
	case Int32:
		return __obf_5d389b660eefb08c.int32()
	case Int64:
		return __obf_5d389b660eefb08c.int64()
	case Bin8, Bin16, Bin32:
		return __obf_5d389b660eefb08c.bytes(__obf_4148125b350dfea2, nil)
	case Str8, Str16, Str32:
		return __obf_5d389b660eefb08c.string(__obf_4148125b350dfea2)
	case Array16, Array32:
		return __obf_5d389b660eefb08c.__obf_e557eb76c54a6f47(__obf_4148125b350dfea2)
	case Map16, Map32:
		__obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.UnreadByte()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_ee0b22a7a221cee0()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_5d389b660eefb08c.__obf_fae8700b6c50321a(__obf_4148125b350dfea2)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_4148125b350dfea2)
}

// DecodeInterfaceLoose is like DecodeInterface except that:
//   - int8, int16, and int32 are converted to int64,
//   - uint8, uint16, and uint32 are converted to uint64,
//   - float32 is converted to float64.
//   - []byte is converted to string.
func (__obf_5d389b660eefb08c *Decoder) DecodeInterfaceLoose() (any, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}

	if IsFixedNum(__obf_4148125b350dfea2) {
		return int64(int8(__obf_4148125b350dfea2)), nil
	}
	if IsFixedMap(__obf_4148125b350dfea2) {
		__obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.UnreadByte()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_ee0b22a7a221cee0()
	}
	if IsFixedArray(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.__obf_e557eb76c54a6f47(__obf_4148125b350dfea2)
	}
	if IsFixedString(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.string(__obf_4148125b350dfea2)
	}

	switch __obf_4148125b350dfea2 {
	case Nil:
		return nil, nil
	case False, True:
		return __obf_5d389b660eefb08c.bool(__obf_4148125b350dfea2)
	case Float, Double:
		return __obf_5d389b660eefb08c.float64(__obf_4148125b350dfea2)
	case Uint8, Uint16, Uint32, Uint64:
		return __obf_5d389b660eefb08c.uint(__obf_4148125b350dfea2)
	case Int8, Int16, Int32, Int64:
		return __obf_5d389b660eefb08c.int(__obf_4148125b350dfea2)
	case Str8, Str16, Str32,
		Bin8, Bin16, Bin32:
		return __obf_5d389b660eefb08c.string(__obf_4148125b350dfea2)
	case Array16, Array32:
		return __obf_5d389b660eefb08c.__obf_e557eb76c54a6f47(__obf_4148125b350dfea2)
	case Map16, Map32:
		__obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.UnreadByte()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		return __obf_5d389b660eefb08c.__obf_ee0b22a7a221cee0()
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_5d389b660eefb08c.__obf_fae8700b6c50321a(__obf_4148125b350dfea2)
	}

	return 0, fmt.Errorf("msgpack: unknown code %x decoding any", __obf_4148125b350dfea2)
}

// Skip skips next value.
func (__obf_5d389b660eefb08c *Decoder) Skip() error {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	if IsFixedNum(__obf_4148125b350dfea2) {
		return nil
	}
	if IsFixedMap(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.__obf_f977b6e0dc35fc2e(__obf_4148125b350dfea2)
	}
	if IsFixedArray(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.__obf_8ed189638e6f0a73(__obf_4148125b350dfea2)
	}
	if IsFixedString(__obf_4148125b350dfea2) {
		return __obf_5d389b660eefb08c.__obf_870f749ef029eefb(__obf_4148125b350dfea2)
	}

	switch __obf_4148125b350dfea2 {
	case Nil, False, True:
		return nil
	case Uint8, Int8:
		return __obf_5d389b660eefb08c.__obf_d93e03df64d057a0(1)
	case Uint16, Int16:
		return __obf_5d389b660eefb08c.__obf_d93e03df64d057a0(2)
	case Uint32, Int32, Float:
		return __obf_5d389b660eefb08c.__obf_d93e03df64d057a0(4)
	case Uint64, Int64, Double:
		return __obf_5d389b660eefb08c.__obf_d93e03df64d057a0(8)
	case Bin8, Bin16, Bin32:
		return __obf_5d389b660eefb08c.__obf_870f749ef029eefb(__obf_4148125b350dfea2)
	case Str8, Str16, Str32:
		return __obf_5d389b660eefb08c.__obf_870f749ef029eefb(__obf_4148125b350dfea2)
	case Array16, Array32:
		return __obf_5d389b660eefb08c.__obf_8ed189638e6f0a73(__obf_4148125b350dfea2)
	case Map16, Map32:
		return __obf_5d389b660eefb08c.__obf_f977b6e0dc35fc2e(__obf_4148125b350dfea2)
	case FixExt1, FixExt2, FixExt4, FixExt8, FixExt16,
		Ext8, Ext16, Ext32:
		return __obf_5d389b660eefb08c.__obf_d955aa85bcfecb15(__obf_4148125b350dfea2)
	}

	return fmt.Errorf("msgpack: unknown code %x", __obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeRaw() (RawMessage, error) {
	__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = make([]byte, 0)
	if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	__obf_21691f6d5038a3f0 := RawMessage(__obf_5d389b660eefb08c.__obf_98388ae8c389d758)
	__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = nil
	return __obf_21691f6d5038a3f0, nil
}

// PeekCode returns the next MessagePack code without advancing the reader.
// Subpackage msgpack/codes defines the list of available
func (__obf_5d389b660eefb08c *Decoder) PeekCode() (byte, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.ReadByte()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_4148125b350dfea2, __obf_5d389b660eefb08c.

		// ReadFull reads exactly len(buf) bytes into the buf.
		__obf_fe1ace7a2eb8bf9f.UnreadByte()
}

func (__obf_5d389b660eefb08c *Decoder) ReadFull(__obf_099b1f8c6882e66a []byte) error {
	_, __obf_45342a3a754d12f5 := __obf_31c39140d4d66749(__obf_5d389b660eefb08c.__obf_52c07f7b574f636a, __obf_099b1f8c6882e66a, len(__obf_099b1f8c6882e66a))
	return __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) __obf_982d42aefd776ffb() bool {
	__obf_cde82889ba8e4822, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.PeekCode()
	return __obf_45342a3a754d12f5 == nil && __obf_cde82889ba8e4822 == Nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_4786096a37bd32ce() (byte, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_fe1ace7a2eb8bf9f.ReadByte()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	if __obf_5d389b660eefb08c.__obf_98388ae8c389d758 != nil {
		__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = append(__obf_5d389b660eefb08c.__obf_98388ae8c389d758, __obf_4148125b350dfea2)
	}
	return __obf_4148125b350dfea2, nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_e6a3e7cb981bc1c2(__obf_94333af0f0facd65 []byte) error {
	_, __obf_45342a3a754d12f5 := io.ReadFull(__obf_5d389b660eefb08c.__obf_52c07f7b574f636a, __obf_94333af0f0facd65)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_5d389b660eefb08c.__obf_98388ae8c389d758 != nil {
		__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = append(__obf_5d389b660eefb08c.__obf_98388ae8c389d758, __obf_94333af0f0facd65...)
	}
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_31c39140d4d66749(__obf_d774e4844c74bfe9 int) ([]byte, error) {
	var __obf_45342a3a754d12f5 error
	if __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_d0b80752c6a68c61 != 0 {
		__obf_5d389b660eefb08c.__obf_099b1f8c6882e66a, __obf_45342a3a754d12f5 = __obf_31c39140d4d66749(__obf_5d389b660eefb08c.__obf_52c07f7b574f636a, __obf_5d389b660eefb08c.__obf_099b1f8c6882e66a, __obf_d774e4844c74bfe9)
	} else {
		__obf_5d389b660eefb08c.__obf_099b1f8c6882e66a, __obf_45342a3a754d12f5 = __obf_e59aa74eb597ad9f(__obf_5d389b660eefb08c.__obf_52c07f7b574f636a, __obf_5d389b660eefb08c.__obf_099b1f8c6882e66a,

			// TODO: read directly into d.rec?
			__obf_d774e4844c74bfe9)
	}
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	if __obf_5d389b660eefb08c.__obf_98388ae8c389d758 != nil {
		__obf_5d389b660eefb08c.__obf_98388ae8c389d758 = append(__obf_5d389b660eefb08c.__obf_98388ae8c389d758, __obf_5d389b660eefb08c.__obf_099b1f8c6882e66a...)
	}
	return __obf_5d389b660eefb08c.__obf_099b1f8c6882e66a, nil
}

func __obf_31c39140d4d66749(__obf_52c07f7b574f636a io.Reader, __obf_94333af0f0facd65 []byte, __obf_d774e4844c74bfe9 int) ([]byte, error) {
	if __obf_94333af0f0facd65 == nil {
		if __obf_d774e4844c74bfe9 == 0 {
			return make([]byte, 0), nil
		}
		__obf_94333af0f0facd65 = make([]byte, 0, __obf_d774e4844c74bfe9)
	}

	if __obf_d774e4844c74bfe9 > cap(__obf_94333af0f0facd65) {
		__obf_94333af0f0facd65 = append(__obf_94333af0f0facd65, make([]byte, __obf_d774e4844c74bfe9-len(__obf_94333af0f0facd65))...)
	} else if __obf_d774e4844c74bfe9 <= cap(__obf_94333af0f0facd65) {
		__obf_94333af0f0facd65 = __obf_94333af0f0facd65[:__obf_d774e4844c74bfe9]
	}

	_, __obf_45342a3a754d12f5 := io.ReadFull(__obf_52c07f7b574f636a, __obf_94333af0f0facd65)
	return __obf_94333af0f0facd65, __obf_45342a3a754d12f5
}

func __obf_e59aa74eb597ad9f(__obf_52c07f7b574f636a io.Reader, __obf_94333af0f0facd65 []byte, __obf_d774e4844c74bfe9 int) ([]byte, error) {
	if __obf_94333af0f0facd65 == nil {
		if __obf_d774e4844c74bfe9 == 0 {
			return make([]byte, 0), nil
		}
		switch {
		case __obf_d774e4844c74bfe9 < 64:
			__obf_94333af0f0facd65 = make([]byte, 0, 64)
		case __obf_d774e4844c74bfe9 <= __obf_6bc1f5097e067b3d:
			__obf_94333af0f0facd65 = make([]byte, 0, __obf_d774e4844c74bfe9)
		default:
			__obf_94333af0f0facd65 = make([]byte, 0, __obf_6bc1f5097e067b3d)
		}
	}

	if __obf_d774e4844c74bfe9 <= cap(__obf_94333af0f0facd65) {
		__obf_94333af0f0facd65 = __obf_94333af0f0facd65[:__obf_d774e4844c74bfe9]
		_, __obf_45342a3a754d12f5 := io.ReadFull(__obf_52c07f7b574f636a, __obf_94333af0f0facd65)
		return __obf_94333af0f0facd65, __obf_45342a3a754d12f5
	}
	__obf_94333af0f0facd65 = __obf_94333af0f0facd65[:cap(__obf_94333af0f0facd65)]

	var __obf_37517a3103f26511 int
	for {
		__obf_47e568181aee74d9 := min(__obf_d774e4844c74bfe9-len(__obf_94333af0f0facd65), __obf_6bc1f5097e067b3d)
		__obf_94333af0f0facd65 = append(__obf_94333af0f0facd65, make([]byte, __obf_47e568181aee74d9)...)

		_, __obf_45342a3a754d12f5 := io.ReadFull(__obf_52c07f7b574f636a, __obf_94333af0f0facd65[__obf_37517a3103f26511:])
		if __obf_45342a3a754d12f5 != nil {
			return __obf_94333af0f0facd65, __obf_45342a3a754d12f5
		}

		if len(__obf_94333af0f0facd65) == __obf_d774e4844c74bfe9 {
			break
		}
		__obf_37517a3103f26511 = len(__obf_94333af0f0facd65)
	}

	return __obf_94333af0f0facd65, nil
}

func min(__obf_28a2fa10042399b1, //nolint:unparam
	__obf_94333af0f0facd65 int) int {
	if __obf_28a2fa10042399b1 <= __obf_94333af0f0facd65 {
		return __obf_28a2fa10042399b1
	}
	return __obf_94333af0f0facd65
}

func __obf_322f79ef2f26d64d(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.bytes(__obf_4148125b350dfea2, __obf_f328a048f76b7256.Bytes())
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetBytes(__obf_94333af0f0facd65)
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) DecodeBytesLen() (int, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeBytes() ([]byte, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.bytes(__obf_4148125b350dfea2, nil)
}

func (__obf_5d389b660eefb08c *Decoder) bytes(__obf_4148125b350dfea2 byte, __obf_94333af0f0facd65 []byte) ([]byte, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return nil, nil
	}
	return __obf_31c39140d4d66749(__obf_5d389b660eefb08c.__obf_52c07f7b574f636a, __obf_94333af0f0facd65, __obf_d774e4844c74bfe9)
}
