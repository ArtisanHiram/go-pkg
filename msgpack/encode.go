package __obf_fd770cb9675903b0

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_4010aa531727deba uint32 = 1 << iota
	__obf_b1ce96c64c5697ef
	__obf_a787b0ded9ae6f01
	__obf_19c02dcbfe4a74e9
	__obf_139c02350e62bdf7
	__obf_bb9969de1dccd587
)

type __obf_fe28721d9714bae7 interface {
	io.Writer
	WriteByte(byte) error
}

type __obf_05373cf17f9ff685 struct {
	io.Writer
}

func __obf_7787fc171aca764a(__obf_41c5519e4a6d21b0 io.Writer) __obf_05373cf17f9ff685 {
	return __obf_05373cf17f9ff685{
		Writer: __obf_41c5519e4a6d21b0,
	}
}

func (__obf_31d610987a73a02c __obf_05373cf17f9ff685) WriteByte(__obf_4148125b350dfea2 byte) error {
	_, __obf_45342a3a754d12f5 := __obf_31d610987a73a02c.Write([]byte{__obf_4148125b350dfea2})
	return __obf_45342a3a754d12f5
}

//------------------------------------------------------------------------------

var __obf_837ce45abf27ecab = sync.Pool{
	New: func() any {
		return NewEncoder(nil)
	},
}

func GetEncoder() *Encoder {
	return __obf_837ce45abf27ecab.Get().(*Encoder)
}

func PutEncoder(__obf_a7b5a0ca650f48ee *Encoder) {
	__obf_a7b5a0ca650f48ee.__obf_41c5519e4a6d21b0 = nil
	__obf_837ce45abf27ecab.
		Put(__obf_a7b5a0ca650f48ee)
}

// Marshal returns the MessagePack encoding of v.
func Marshal(__obf_f328a048f76b7256 any) ([]byte, error) {
	__obf_a7b5a0ca650f48ee := GetEncoder()

	var __obf_099b1f8c6882e66a bytes.Buffer
	__obf_a7b5a0ca650f48ee.
		Reset(&__obf_099b1f8c6882e66a)
	__obf_45342a3a754d12f5 := __obf_a7b5a0ca650f48ee.Encode(__obf_f328a048f76b7256)
	__obf_94333af0f0facd65 := __obf_099b1f8c6882e66a.Bytes()

	PutEncoder(__obf_a7b5a0ca650f48ee)

	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	return __obf_94333af0f0facd65, __obf_45342a3a754d12f5
}

type Encoder struct {
	__obf_41c5519e4a6d21b0 __obf_fe28721d9714bae7
	__obf_ff96db22e12b6842 map[string]int
	__obf_6621ba3773b8696a string
	__obf_099b1f8c6882e66a []byte
	__obf_07a810e976869435 []byte
	__obf_7185cb62de7af792 uint32
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(__obf_41c5519e4a6d21b0 io.Writer) *Encoder {
	__obf_e9038cda3b5cf711 := &Encoder{__obf_099b1f8c6882e66a: make([]byte, 9)}
	__obf_e9038cda3b5cf711.
		Reset(__obf_41c5519e4a6d21b0)
	return __obf_e9038cda3b5cf711
}

// Writer returns the Encoder's writer.
func (__obf_e9038cda3b5cf711 *Encoder) Writer() io.Writer {
	return __obf_e9038cda3b5cf711.

		// Reset discards any buffered data, resets all state, and switches the writer to write to w.
		__obf_41c5519e4a6d21b0
}

func (__obf_e9038cda3b5cf711 *Encoder) Reset(__obf_41c5519e4a6d21b0 io.Writer) {
	__obf_e9038cda3b5cf711.
		ResetDict(__obf_41c5519e4a6d21b0, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_e9038cda3b5cf711 *Encoder) ResetDict(__obf_41c5519e4a6d21b0 io.Writer, __obf_ff96db22e12b6842 map[string]int) {
	__obf_e9038cda3b5cf711.
		ResetWriter(__obf_41c5519e4a6d21b0)
	__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 = 0
	__obf_e9038cda3b5cf711.__obf_6621ba3773b8696a = ""
	__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842 = __obf_ff96db22e12b6842
}

func (__obf_e9038cda3b5cf711 *Encoder) WithDict(__obf_ff96db22e12b6842 map[string]int, __obf_1e4576e8508b04bc func(*Encoder) error) error {
	__obf_cf0262a7ff1775c5 := __obf_e9038cda3b5cf711.__obf_ff96db22e12b6842
	__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842 = __obf_ff96db22e12b6842
	__obf_45342a3a754d12f5 := __obf_1e4576e8508b04bc(__obf_e9038cda3b5cf711)
	__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842 = __obf_cf0262a7ff1775c5
	return __obf_45342a3a754d12f5
}

func (__obf_e9038cda3b5cf711 *Encoder) ResetWriter(__obf_41c5519e4a6d21b0 io.Writer) {
	__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842 = nil
	if __obf_31d610987a73a02c, __obf_b00b3c6a10f90467 := __obf_41c5519e4a6d21b0.(__obf_fe28721d9714bae7); __obf_b00b3c6a10f90467 {
		__obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0 = __obf_31d610987a73a02c
	} else if __obf_41c5519e4a6d21b0 == nil {
		__obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0 = nil
	} else {
		__obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0 = __obf_7787fc171aca764a(__obf_41c5519e4a6d21b0)
	}
}

// SetSortMapKeys causes the Encoder to encode map keys in increasing order.
// Supported map types are:
//   - map[string]string
//   - map[string]bool
//   - map[string]any
func (__obf_e9038cda3b5cf711 *Encoder) SetSortMapKeys(__obf_327e8d18e7c070c5 bool) *Encoder {
	if __obf_327e8d18e7c070c5 {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 |= __obf_4010aa531727deba
	} else {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 &= ^__obf_4010aa531727deba
	}
	return __obf_e9038cda3b5cf711
}

// SetCustomStructTag causes the Encoder to use a custom struct tag as
// fallback option if there is no msgpack tag.
func (__obf_e9038cda3b5cf711 *Encoder) SetCustomStructTag(__obf_1a0f202964305730 string) {
	__obf_e9038cda3b5cf711.__obf_6621ba3773b8696a = __obf_1a0f202964305730
}

// SetOmitEmpty causes the Encoder to omit empty values by default.
func (__obf_e9038cda3b5cf711 *Encoder) SetOmitEmpty(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 |= __obf_bb9969de1dccd587
	} else {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 &= ^__obf_bb9969de1dccd587
	}
}

// UseArrayEncodedStructs causes the Encoder to encode Go structs as msgpack arrays.
func (__obf_e9038cda3b5cf711 *Encoder) UseArrayEncodedStructs(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 |= __obf_b1ce96c64c5697ef
	} else {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 &= ^__obf_b1ce96c64c5697ef
	}
}

// UseCompactEncoding causes the Encoder to chose the most compact encoding.
// For example, it allows to encode small Go int64 as msgpack int8 saving 7 bytes.
func (__obf_e9038cda3b5cf711 *Encoder) UseCompactInts(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 |= __obf_a787b0ded9ae6f01
	} else {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 &= ^__obf_a787b0ded9ae6f01
	}
}

// UseCompactFloats causes the Encoder to chose a compact integer encoding
// for floats that can be represented as integers.
func (__obf_e9038cda3b5cf711 *Encoder) UseCompactFloats(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 |= __obf_19c02dcbfe4a74e9
	} else {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 &= ^__obf_19c02dcbfe4a74e9
	}
}

// UseInternedStrings causes the Encoder to intern strings.
func (__obf_e9038cda3b5cf711 *Encoder) UseInternedStrings(__obf_327e8d18e7c070c5 bool) {
	if __obf_327e8d18e7c070c5 {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 |= __obf_139c02350e62bdf7
	} else {
		__obf_e9038cda3b5cf711.__obf_7185cb62de7af792 &= ^__obf_139c02350e62bdf7
	}
}

func (__obf_e9038cda3b5cf711 *Encoder) Encode(__obf_f328a048f76b7256 any) error {
	switch __obf_f328a048f76b7256 := __obf_f328a048f76b7256.(type) {
	case nil:
		return __obf_e9038cda3b5cf711.EncodeNil()
	case string:
		return __obf_e9038cda3b5cf711.EncodeString(__obf_f328a048f76b7256)
	case []byte:
		return __obf_e9038cda3b5cf711.EncodeBytes(__obf_f328a048f76b7256)
	case int:
		return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_f328a048f76b7256))
	case int64:
		return __obf_e9038cda3b5cf711.__obf_eac777c6448d3606(__obf_f328a048f76b7256)
	case uint:
		return __obf_e9038cda3b5cf711.EncodeUint(uint64(__obf_f328a048f76b7256))
	case uint64:
		return __obf_e9038cda3b5cf711.__obf_c6c82d67d256dc52(__obf_f328a048f76b7256)
	case bool:
		return __obf_e9038cda3b5cf711.EncodeBool(__obf_f328a048f76b7256)
	case float32:
		return __obf_e9038cda3b5cf711.EncodeFloat32(__obf_f328a048f76b7256)
	case float64:
		return __obf_e9038cda3b5cf711.EncodeFloat64(__obf_f328a048f76b7256)
	case time.Duration:
		return __obf_e9038cda3b5cf711.__obf_eac777c6448d3606(int64(__obf_f328a048f76b7256))
	case time.Time:
		return __obf_e9038cda3b5cf711.EncodeTime(__obf_f328a048f76b7256)
	}
	return __obf_e9038cda3b5cf711.EncodeValue(reflect.ValueOf(__obf_f328a048f76b7256))
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeMulti(__obf_f328a048f76b7256 ...any) error {
	for _, __obf_e2c885d3da27fd46 := range __obf_f328a048f76b7256 {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.Encode(__obf_e2c885d3da27fd46); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeValue(__obf_f328a048f76b7256 reflect.Value) error {
	__obf_1e4576e8508b04bc := __obf_031e3a13e30d95dc(__obf_f328a048f76b7256.Type())
	return __obf_1e4576e8508b04bc(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeNil() error {
	return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(Nil)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeBool(__obf_28cbfc96ff0a56b6 bool) error {
	if __obf_28cbfc96ff0a56b6 {
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(True)
	}
	return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(False)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeDuration(__obf_5d389b660eefb08c time.Duration) error {
	return __obf_e9038cda3b5cf711.EncodeInt(int64(__obf_5d389b660eefb08c))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_a72c49227d01f9dd(__obf_4148125b350dfea2 byte) error {
	return __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.WriteByte(__obf_4148125b350dfea2)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_af6d14a7babbd464(__obf_94333af0f0facd65 []byte) error {
	_, __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.Write(__obf_94333af0f0facd65)
	return __obf_45342a3a754d12f5
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_016cc47c9f09546b(__obf_fe1ace7a2eb8bf9f string) error {
	_, __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_41c5519e4a6d21b0.Write(__obf_9b773961ed35818a(__obf_fe1ace7a2eb8bf9f))
	return __obf_45342a3a754d12f5
}
