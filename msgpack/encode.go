package __obf_3edaa49e853afa16

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_67f6be0f2d2d4aa4 uint32 = 1 << iota
	__obf_a4be32a461c6088d
	__obf_5f5cd069c38159b5
	__obf_f553c7da4d60e61e
	__obf_48e006ba02b70b78
	__obf_344e8fa9adf9a114
)

type __obf_101c7dee423c31b9 interface {
	io.Writer
	WriteByte(byte) error
}

type __obf_609d711402a2acf0 struct {
	io.Writer
}

func __obf_1ef8882743fb843e(__obf_0f9492f42c73d099 io.Writer) __obf_609d711402a2acf0 {
	return __obf_609d711402a2acf0{
		Writer: __obf_0f9492f42c73d099,
	}
}

func (__obf_a89df47bb8a33ffe __obf_609d711402a2acf0) WriteByte(__obf_145c7a7d25eea2bd byte) error {
	_, __obf_3aa78ad28f50ed46 := __obf_a89df47bb8a33ffe.Write([]byte{__obf_145c7a7d25eea2bd})
	return __obf_3aa78ad28f50ed46
}

//------------------------------------------------------------------------------

var __obf_7d9ea374ab7a721b = sync.Pool{
	New: func() any {
		return NewEncoder(nil)
	},
}

func GetEncoder() *Encoder {
	return __obf_7d9ea374ab7a721b.Get().(*Encoder)
}

func PutEncoder(__obf_2fa4ef37f70ce1de *Encoder) {
	__obf_2fa4ef37f70ce1de.__obf_0f9492f42c73d099 = nil
	__obf_7d9ea374ab7a721b.
		Put(__obf_2fa4ef37f70ce1de)
}

// Marshal returns the MessagePack encoding of v.
func Marshal(__obf_85d270343a0dfe14 any) ([]byte, error) {
	__obf_2fa4ef37f70ce1de := GetEncoder()

	var __obf_8f3bed0c23610f75 bytes.Buffer
	__obf_2fa4ef37f70ce1de.
		Reset(&__obf_8f3bed0c23610f75)
	__obf_3aa78ad28f50ed46 := __obf_2fa4ef37f70ce1de.Encode(__obf_85d270343a0dfe14)
	__obf_9b4a5a04bdad2347 := __obf_8f3bed0c23610f75.Bytes()

	PutEncoder(__obf_2fa4ef37f70ce1de)

	if __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	return __obf_9b4a5a04bdad2347, __obf_3aa78ad28f50ed46
}

type Encoder struct {
	__obf_0f9492f42c73d099 __obf_101c7dee423c31b9
	__obf_b014355f64826d7b map[string]int
	__obf_ec5d9a4ce8dd41be string
	__obf_8f3bed0c23610f75 []byte
	__obf_c0eccc33acc16ac1 []byte
	__obf_6aa0efd537415192 uint32
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(__obf_0f9492f42c73d099 io.Writer) *Encoder {
	__obf_84d0d31a8288f191 := &Encoder{__obf_8f3bed0c23610f75: make([]byte, 9)}
	__obf_84d0d31a8288f191.
		Reset(__obf_0f9492f42c73d099)
	return __obf_84d0d31a8288f191
}

// Writer returns the Encoder's writer.
func (__obf_84d0d31a8288f191 *Encoder) Writer() io.Writer {
	return __obf_84d0d31a8288f191.

		// Reset discards any buffered data, resets all state, and switches the writer to write to w.
		__obf_0f9492f42c73d099
}

func (__obf_84d0d31a8288f191 *Encoder) Reset(__obf_0f9492f42c73d099 io.Writer) {
	__obf_84d0d31a8288f191.
		ResetDict(__obf_0f9492f42c73d099, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_84d0d31a8288f191 *Encoder) ResetDict(__obf_0f9492f42c73d099 io.Writer, __obf_b014355f64826d7b map[string]int) {
	__obf_84d0d31a8288f191.
		ResetWriter(__obf_0f9492f42c73d099)
	__obf_84d0d31a8288f191.__obf_6aa0efd537415192 = 0
	__obf_84d0d31a8288f191.__obf_ec5d9a4ce8dd41be = ""
	__obf_84d0d31a8288f191.__obf_b014355f64826d7b = __obf_b014355f64826d7b
}

func (__obf_84d0d31a8288f191 *Encoder) WithDict(__obf_b014355f64826d7b map[string]int, __obf_c8137d8e88d5b1dd func(*Encoder) error) error {
	__obf_119d5920d120ff61 := __obf_84d0d31a8288f191.__obf_b014355f64826d7b
	__obf_84d0d31a8288f191.__obf_b014355f64826d7b = __obf_b014355f64826d7b
	__obf_3aa78ad28f50ed46 := __obf_c8137d8e88d5b1dd(__obf_84d0d31a8288f191)
	__obf_84d0d31a8288f191.__obf_b014355f64826d7b = __obf_119d5920d120ff61
	return __obf_3aa78ad28f50ed46
}

func (__obf_84d0d31a8288f191 *Encoder) ResetWriter(__obf_0f9492f42c73d099 io.Writer) {
	__obf_84d0d31a8288f191.__obf_b014355f64826d7b = nil
	if __obf_a89df47bb8a33ffe, __obf_ccfb92cc26e4569f := __obf_0f9492f42c73d099.(__obf_101c7dee423c31b9); __obf_ccfb92cc26e4569f {
		__obf_84d0d31a8288f191.__obf_0f9492f42c73d099 = __obf_a89df47bb8a33ffe
	} else if __obf_0f9492f42c73d099 == nil {
		__obf_84d0d31a8288f191.__obf_0f9492f42c73d099 = nil
	} else {
		__obf_84d0d31a8288f191.__obf_0f9492f42c73d099 = __obf_1ef8882743fb843e(__obf_0f9492f42c73d099)
	}
}

// SetSortMapKeys causes the Encoder to encode map keys in increasing order.
// Supported map types are:
//   - map[string]string
//   - map[string]bool
//   - map[string]any
func (__obf_84d0d31a8288f191 *Encoder) SetSortMapKeys(__obf_6345dbff7f40032f bool) *Encoder {
	if __obf_6345dbff7f40032f {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 |= __obf_67f6be0f2d2d4aa4
	} else {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 &= ^__obf_67f6be0f2d2d4aa4
	}
	return __obf_84d0d31a8288f191
}

// SetCustomStructTag causes the Encoder to use a custom struct tag as
// fallback option if there is no msgpack tag.
func (__obf_84d0d31a8288f191 *Encoder) SetCustomStructTag(__obf_160a7dd2c812a6a1 string) {
	__obf_84d0d31a8288f191.__obf_ec5d9a4ce8dd41be = __obf_160a7dd2c812a6a1
}

// SetOmitEmpty causes the Encoder to omit empty values by default.
func (__obf_84d0d31a8288f191 *Encoder) SetOmitEmpty(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 |= __obf_344e8fa9adf9a114
	} else {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 &= ^__obf_344e8fa9adf9a114
	}
}

// UseArrayEncodedStructs causes the Encoder to encode Go structs as msgpack arrays.
func (__obf_84d0d31a8288f191 *Encoder) UseArrayEncodedStructs(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 |= __obf_a4be32a461c6088d
	} else {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 &= ^__obf_a4be32a461c6088d
	}
}

// UseCompactEncoding causes the Encoder to chose the most compact encoding.
// For example, it allows to encode small Go int64 as msgpack int8 saving 7 bytes.
func (__obf_84d0d31a8288f191 *Encoder) UseCompactInts(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 |= __obf_5f5cd069c38159b5
	} else {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 &= ^__obf_5f5cd069c38159b5
	}
}

// UseCompactFloats causes the Encoder to chose a compact integer encoding
// for floats that can be represented as integers.
func (__obf_84d0d31a8288f191 *Encoder) UseCompactFloats(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 |= __obf_f553c7da4d60e61e
	} else {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 &= ^__obf_f553c7da4d60e61e
	}
}

// UseInternedStrings causes the Encoder to intern strings.
func (__obf_84d0d31a8288f191 *Encoder) UseInternedStrings(__obf_6345dbff7f40032f bool) {
	if __obf_6345dbff7f40032f {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 |= __obf_48e006ba02b70b78
	} else {
		__obf_84d0d31a8288f191.__obf_6aa0efd537415192 &= ^__obf_48e006ba02b70b78
	}
}

func (__obf_84d0d31a8288f191 *Encoder) Encode(__obf_85d270343a0dfe14 any) error {
	switch __obf_85d270343a0dfe14 := __obf_85d270343a0dfe14.(type) {
	case nil:
		return __obf_84d0d31a8288f191.EncodeNil()
	case string:
		return __obf_84d0d31a8288f191.EncodeString(__obf_85d270343a0dfe14)
	case []byte:
		return __obf_84d0d31a8288f191.EncodeBytes(__obf_85d270343a0dfe14)
	case int:
		return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_85d270343a0dfe14))
	case int64:
		return __obf_84d0d31a8288f191.__obf_b746e88d88f1bc44(__obf_85d270343a0dfe14)
	case uint:
		return __obf_84d0d31a8288f191.EncodeUint(uint64(__obf_85d270343a0dfe14))
	case uint64:
		return __obf_84d0d31a8288f191.__obf_dde4b3164695e6f8(__obf_85d270343a0dfe14)
	case bool:
		return __obf_84d0d31a8288f191.EncodeBool(__obf_85d270343a0dfe14)
	case float32:
		return __obf_84d0d31a8288f191.EncodeFloat32(__obf_85d270343a0dfe14)
	case float64:
		return __obf_84d0d31a8288f191.EncodeFloat64(__obf_85d270343a0dfe14)
	case time.Duration:
		return __obf_84d0d31a8288f191.__obf_b746e88d88f1bc44(int64(__obf_85d270343a0dfe14))
	case time.Time:
		return __obf_84d0d31a8288f191.EncodeTime(__obf_85d270343a0dfe14)
	}
	return __obf_84d0d31a8288f191.EncodeValue(reflect.ValueOf(__obf_85d270343a0dfe14))
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeMulti(__obf_85d270343a0dfe14 ...any) error {
	for _, __obf_ab0d31723a444d49 := range __obf_85d270343a0dfe14 {
		if __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.Encode(__obf_ab0d31723a444d49); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeValue(__obf_85d270343a0dfe14 reflect.Value) error {
	__obf_c8137d8e88d5b1dd := __obf_1e49c68fa4194051(__obf_85d270343a0dfe14.Type())
	return __obf_c8137d8e88d5b1dd(__obf_84d0d31a8288f191, __obf_85d270343a0dfe14)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeNil() error {
	return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(Nil)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeBool(__obf_e6992bb5202a647c bool) error {
	if __obf_e6992bb5202a647c {
		return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(True)
	}
	return __obf_84d0d31a8288f191.__obf_4268b610398f5a51(False)
}

func (__obf_84d0d31a8288f191 *Encoder) EncodeDuration(__obf_015afbee33862a01 time.Duration) error {
	return __obf_84d0d31a8288f191.EncodeInt(int64(__obf_015afbee33862a01))
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_4268b610398f5a51(__obf_145c7a7d25eea2bd byte) error {
	return __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.WriteByte(__obf_145c7a7d25eea2bd)
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_fe140b946d6a369e(__obf_9b4a5a04bdad2347 []byte) error {
	_, __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.Write(__obf_9b4a5a04bdad2347)
	return __obf_3aa78ad28f50ed46
}

func (__obf_84d0d31a8288f191 *Encoder) __obf_938c3e39051aebef(__obf_6827ff1b59d7ccec string) error {
	_, __obf_3aa78ad28f50ed46 := __obf_84d0d31a8288f191.__obf_0f9492f42c73d099.Write(__obf_f0406d601699c6f7(__obf_6827ff1b59d7ccec))
	return __obf_3aa78ad28f50ed46
}
