package __obf_a935eb7f548271a4

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_51d2317856ab3853 uint32 = 1 << iota
	__obf_d6a65589ccd01600
	__obf_123cc756bbdcb1da
	__obf_9c276895701d3946
	__obf_bf24fd2d64b91255
	__obf_48749b6f1055fd1f
)

type __obf_5cc4ea1df41d6297 interface {
	io.Writer
	WriteByte(byte) error
}

type __obf_d22969359cb255ff struct {
	io.Writer
}

func __obf_b6c9e834d863298e(__obf_a3adbf62c8801563 io.Writer) __obf_d22969359cb255ff {
	return __obf_d22969359cb255ff{
		Writer: __obf_a3adbf62c8801563,
	}
}

func (__obf_d6463c1892be423c __obf_d22969359cb255ff) WriteByte(__obf_f5df560f4d67421b byte) error {
	_, __obf_4d327e1cd40c2e21 := __obf_d6463c1892be423c.Write([]byte{__obf_f5df560f4d67421b})
	return __obf_4d327e1cd40c2e21
}

//------------------------------------------------------------------------------

var __obf_be9724e324463ba4 = sync.Pool{
	New: func() any {
		return NewEncoder(nil)
	},
}

func GetEncoder() *Encoder {
	return __obf_be9724e324463ba4.Get().(*Encoder)
}

func PutEncoder(__obf_6ca0309463018ed3 *Encoder) {
	__obf_6ca0309463018ed3.__obf_a3adbf62c8801563 = nil
	__obf_be9724e324463ba4.
		Put(__obf_6ca0309463018ed3)
}

// Marshal returns the MessagePack encoding of v.
func Marshal(__obf_6d570581f4b60dbc any) ([]byte, error) {
	__obf_6ca0309463018ed3 := GetEncoder()

	var __obf_a2e16952f247f218 bytes.Buffer
	__obf_6ca0309463018ed3.
		Reset(&__obf_a2e16952f247f218)
	__obf_4d327e1cd40c2e21 := __obf_6ca0309463018ed3.Encode(__obf_6d570581f4b60dbc)
	__obf_f2ca794293605b73 := __obf_a2e16952f247f218.Bytes()

	PutEncoder(__obf_6ca0309463018ed3)

	if __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	return __obf_f2ca794293605b73, __obf_4d327e1cd40c2e21
}

type Encoder struct {
	__obf_a3adbf62c8801563 __obf_5cc4ea1df41d6297
	__obf_2cdd7fad308aef1d map[string]int
	__obf_90e41d275c699c20 string
	__obf_a2e16952f247f218 []byte
	__obf_f1da64bc48a2ea7f []byte
	__obf_d617f3769ce16c47 uint32
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(__obf_a3adbf62c8801563 io.Writer) *Encoder {
	__obf_ed37a34c347049f3 := &Encoder{__obf_a2e16952f247f218: make([]byte, 9)}
	__obf_ed37a34c347049f3.
		Reset(__obf_a3adbf62c8801563)
	return __obf_ed37a34c347049f3
}

// Writer returns the Encoder's writer.
func (__obf_ed37a34c347049f3 *Encoder) Writer() io.Writer {
	return __obf_ed37a34c347049f3.

		// Reset discards any buffered data, resets all state, and switches the writer to write to w.
		__obf_a3adbf62c8801563
}

func (__obf_ed37a34c347049f3 *Encoder) Reset(__obf_a3adbf62c8801563 io.Writer) {
	__obf_ed37a34c347049f3.
		ResetDict(__obf_a3adbf62c8801563, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_ed37a34c347049f3 *Encoder) ResetDict(__obf_a3adbf62c8801563 io.Writer, __obf_2cdd7fad308aef1d map[string]int) {
	__obf_ed37a34c347049f3.
		ResetWriter(__obf_a3adbf62c8801563)
	__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 = 0
	__obf_ed37a34c347049f3.__obf_90e41d275c699c20 = ""
	__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d = __obf_2cdd7fad308aef1d
}

func (__obf_ed37a34c347049f3 *Encoder) WithDict(__obf_2cdd7fad308aef1d map[string]int, __obf_36c78c696c47cfdb func(*Encoder) error) error {
	__obf_f893ccda83eefbf7 := __obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d
	__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d = __obf_2cdd7fad308aef1d
	__obf_4d327e1cd40c2e21 := __obf_36c78c696c47cfdb(__obf_ed37a34c347049f3)
	__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d = __obf_f893ccda83eefbf7
	return __obf_4d327e1cd40c2e21
}

func (__obf_ed37a34c347049f3 *Encoder) ResetWriter(__obf_a3adbf62c8801563 io.Writer) {
	__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d = nil
	if __obf_d6463c1892be423c, __obf_826ac2dbb957d7df := __obf_a3adbf62c8801563.(__obf_5cc4ea1df41d6297); __obf_826ac2dbb957d7df {
		__obf_ed37a34c347049f3.__obf_a3adbf62c8801563 = __obf_d6463c1892be423c
	} else if __obf_a3adbf62c8801563 == nil {
		__obf_ed37a34c347049f3.__obf_a3adbf62c8801563 = nil
	} else {
		__obf_ed37a34c347049f3.__obf_a3adbf62c8801563 = __obf_b6c9e834d863298e(__obf_a3adbf62c8801563)
	}
}

// SetSortMapKeys causes the Encoder to encode map keys in increasing order.
// Supported map types are:
//   - map[string]string
//   - map[string]bool
//   - map[string]any
func (__obf_ed37a34c347049f3 *Encoder) SetSortMapKeys(__obf_9632803b7d3e9b4e bool) *Encoder {
	if __obf_9632803b7d3e9b4e {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 |= __obf_51d2317856ab3853
	} else {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 &= ^__obf_51d2317856ab3853
	}
	return __obf_ed37a34c347049f3
}

// SetCustomStructTag causes the Encoder to use a custom struct tag as
// fallback option if there is no msgpack tag.
func (__obf_ed37a34c347049f3 *Encoder) SetCustomStructTag(__obf_f03106ae8627dc7d string) {
	__obf_ed37a34c347049f3.__obf_90e41d275c699c20 = __obf_f03106ae8627dc7d
}

// SetOmitEmpty causes the Encoder to omit empty values by default.
func (__obf_ed37a34c347049f3 *Encoder) SetOmitEmpty(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 |= __obf_48749b6f1055fd1f
	} else {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 &= ^__obf_48749b6f1055fd1f
	}
}

// UseArrayEncodedStructs causes the Encoder to encode Go structs as msgpack arrays.
func (__obf_ed37a34c347049f3 *Encoder) UseArrayEncodedStructs(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 |= __obf_d6a65589ccd01600
	} else {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 &= ^__obf_d6a65589ccd01600
	}
}

// UseCompactEncoding causes the Encoder to chose the most compact encoding.
// For example, it allows to encode small Go int64 as msgpack int8 saving 7 bytes.
func (__obf_ed37a34c347049f3 *Encoder) UseCompactInts(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 |= __obf_123cc756bbdcb1da
	} else {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 &= ^__obf_123cc756bbdcb1da
	}
}

// UseCompactFloats causes the Encoder to chose a compact integer encoding
// for floats that can be represented as integers.
func (__obf_ed37a34c347049f3 *Encoder) UseCompactFloats(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 |= __obf_9c276895701d3946
	} else {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 &= ^__obf_9c276895701d3946
	}
}

// UseInternedStrings causes the Encoder to intern strings.
func (__obf_ed37a34c347049f3 *Encoder) UseInternedStrings(__obf_9632803b7d3e9b4e bool) {
	if __obf_9632803b7d3e9b4e {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 |= __obf_bf24fd2d64b91255
	} else {
		__obf_ed37a34c347049f3.__obf_d617f3769ce16c47 &= ^__obf_bf24fd2d64b91255
	}
}

func (__obf_ed37a34c347049f3 *Encoder) Encode(__obf_6d570581f4b60dbc any) error {
	switch __obf_6d570581f4b60dbc := __obf_6d570581f4b60dbc.(type) {
	case nil:
		return __obf_ed37a34c347049f3.EncodeNil()
	case string:
		return __obf_ed37a34c347049f3.EncodeString(__obf_6d570581f4b60dbc)
	case []byte:
		return __obf_ed37a34c347049f3.EncodeBytes(__obf_6d570581f4b60dbc)
	case int:
		return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_6d570581f4b60dbc))
	case int64:
		return __obf_ed37a34c347049f3.__obf_35627de976adab1d(__obf_6d570581f4b60dbc)
	case uint:
		return __obf_ed37a34c347049f3.EncodeUint(uint64(__obf_6d570581f4b60dbc))
	case uint64:
		return __obf_ed37a34c347049f3.__obf_71513c11118b96bf(__obf_6d570581f4b60dbc)
	case bool:
		return __obf_ed37a34c347049f3.EncodeBool(__obf_6d570581f4b60dbc)
	case float32:
		return __obf_ed37a34c347049f3.EncodeFloat32(__obf_6d570581f4b60dbc)
	case float64:
		return __obf_ed37a34c347049f3.EncodeFloat64(__obf_6d570581f4b60dbc)
	case time.Duration:
		return __obf_ed37a34c347049f3.__obf_35627de976adab1d(int64(__obf_6d570581f4b60dbc))
	case time.Time:
		return __obf_ed37a34c347049f3.EncodeTime(__obf_6d570581f4b60dbc)
	}
	return __obf_ed37a34c347049f3.EncodeValue(reflect.ValueOf(__obf_6d570581f4b60dbc))
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeMulti(__obf_6d570581f4b60dbc ...any) error {
	for _, __obf_6fee38a618ed3afa := range __obf_6d570581f4b60dbc {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.Encode(__obf_6fee38a618ed3afa); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeValue(__obf_6d570581f4b60dbc reflect.Value) error {
	__obf_36c78c696c47cfdb := __obf_d55d857e358cbb2a(__obf_6d570581f4b60dbc.Type())
	return __obf_36c78c696c47cfdb(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeNil() error {
	return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(Nil)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeBool(__obf_205ee5d820f118f1 bool) error {
	if __obf_205ee5d820f118f1 {
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(True)
	}
	return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(False)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeDuration(__obf_a21885da2425f2b2 time.Duration) error {
	return __obf_ed37a34c347049f3.EncodeInt(int64(__obf_a21885da2425f2b2))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_2252dbfd2473b643(__obf_f5df560f4d67421b byte) error {
	return __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.WriteByte(__obf_f5df560f4d67421b)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_654725fc7d18ae48(__obf_f2ca794293605b73 []byte) error {
	_, __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.Write(__obf_f2ca794293605b73)
	return __obf_4d327e1cd40c2e21
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_6515630c78622bba(__obf_b62c60fba2fd9788 string) error {
	_, __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.Write(__obf_2204f386aa52f210(__obf_b62c60fba2fd9788))
	return __obf_4d327e1cd40c2e21
}
