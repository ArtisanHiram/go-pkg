package __obf_de86cdc8ae98b45b

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_3da2e5262e6dbb1b uint32 = 1 << iota
	__obf_1724e39b1a63df3b
	__obf_968e60dffbce5433
	__obf_7a9e24d108f17fe5
	__obf_1fcd090a6da01656
	__obf_eadbff0f95d7376e
)

type __obf_663c3e6006fa918b interface {
	io.Writer
	WriteByte(byte) error
}

type __obf_5bf0b32cf8244e29 struct {
	io.Writer
}

func __obf_871d8e8a2b975593(__obf_007a7bbeb11f001c io.Writer) __obf_5bf0b32cf8244e29 {
	return __obf_5bf0b32cf8244e29{
		Writer: __obf_007a7bbeb11f001c,
	}
}

func (__obf_56974dd8aa648567 __obf_5bf0b32cf8244e29) WriteByte(__obf_ecf6d2d3253a058d byte) error {
	_, __obf_0feb3528b7b709ec := __obf_56974dd8aa648567.Write([]byte{__obf_ecf6d2d3253a058d})
	return __obf_0feb3528b7b709ec
}

//------------------------------------------------------------------------------

var __obf_ac08f7698d5d4853 = sync.Pool{
	New: func() any {
		return NewEncoder(nil)
	},
}

func GetEncoder() *Encoder {
	return __obf_ac08f7698d5d4853.Get().(*Encoder)
}

func PutEncoder(__obf_6a6518f236eeec6e *Encoder) {
	__obf_6a6518f236eeec6e.__obf_007a7bbeb11f001c = nil
	__obf_ac08f7698d5d4853.
		Put(__obf_6a6518f236eeec6e)
}

// Marshal returns the MessagePack encoding of v.
func Marshal(__obf_17732590722140e7 any) ([]byte, error) {
	__obf_6a6518f236eeec6e := GetEncoder()

	var __obf_9c8c6a5e4cc8822c bytes.Buffer
	__obf_6a6518f236eeec6e.
		Reset(&__obf_9c8c6a5e4cc8822c)
	__obf_0feb3528b7b709ec := __obf_6a6518f236eeec6e.Encode(__obf_17732590722140e7)
	__obf_a7fd7acf2bd4435f := __obf_9c8c6a5e4cc8822c.Bytes()

	PutEncoder(__obf_6a6518f236eeec6e)

	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	return __obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec
}

type Encoder struct {
	__obf_007a7bbeb11f001c __obf_663c3e6006fa918b
	__obf_25128eea0d506b65 map[string]int
	__obf_603dff1ad8d49539 string
	__obf_9c8c6a5e4cc8822c []byte
	__obf_985719a68b5833e6 []byte
	__obf_a15a642720e39c3e uint32
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(__obf_007a7bbeb11f001c io.Writer) *Encoder {
	__obf_7bae0b613da60dd3 := &Encoder{__obf_9c8c6a5e4cc8822c: make([]byte, 9)}
	__obf_7bae0b613da60dd3.
		Reset(__obf_007a7bbeb11f001c)
	return __obf_7bae0b613da60dd3
}

// Writer returns the Encoder's writer.
func (__obf_7bae0b613da60dd3 *Encoder) Writer() io.Writer {
	return __obf_7bae0b613da60dd3.

		// Reset discards any buffered data, resets all state, and switches the writer to write to w.
		__obf_007a7bbeb11f001c
}

func (__obf_7bae0b613da60dd3 *Encoder) Reset(__obf_007a7bbeb11f001c io.Writer) {
	__obf_7bae0b613da60dd3.
		ResetDict(__obf_007a7bbeb11f001c, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_7bae0b613da60dd3 *Encoder) ResetDict(__obf_007a7bbeb11f001c io.Writer, __obf_25128eea0d506b65 map[string]int) {
	__obf_7bae0b613da60dd3.
		ResetWriter(__obf_007a7bbeb11f001c)
	__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e = 0
	__obf_7bae0b613da60dd3.__obf_603dff1ad8d49539 = ""
	__obf_7bae0b613da60dd3.__obf_25128eea0d506b65 = __obf_25128eea0d506b65
}

func (__obf_7bae0b613da60dd3 *Encoder) WithDict(__obf_25128eea0d506b65 map[string]int, __obf_69331012fc3414ab func(*Encoder) error) error {
	__obf_3ae10cdb339e5623 := __obf_7bae0b613da60dd3.__obf_25128eea0d506b65
	__obf_7bae0b613da60dd3.__obf_25128eea0d506b65 = __obf_25128eea0d506b65
	__obf_0feb3528b7b709ec := __obf_69331012fc3414ab(__obf_7bae0b613da60dd3)
	__obf_7bae0b613da60dd3.__obf_25128eea0d506b65 = __obf_3ae10cdb339e5623
	return __obf_0feb3528b7b709ec
}

func (__obf_7bae0b613da60dd3 *Encoder) ResetWriter(__obf_007a7bbeb11f001c io.Writer) {
	__obf_7bae0b613da60dd3.__obf_25128eea0d506b65 = nil
	if __obf_56974dd8aa648567, __obf_77cfff95beb3cc99 := __obf_007a7bbeb11f001c.(__obf_663c3e6006fa918b); __obf_77cfff95beb3cc99 {
		__obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c = __obf_56974dd8aa648567
	} else if __obf_007a7bbeb11f001c == nil {
		__obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c = nil
	} else {
		__obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c = __obf_871d8e8a2b975593(__obf_007a7bbeb11f001c)
	}
}

// SetSortMapKeys causes the Encoder to encode map keys in increasing order.
// Supported map types are:
//   - map[string]string
//   - map[string]bool
//   - map[string]any
func (__obf_7bae0b613da60dd3 *Encoder) SetSortMapKeys(__obf_a8129bb947347fed bool) *Encoder {
	if __obf_a8129bb947347fed {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e |= __obf_3da2e5262e6dbb1b
	} else {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e &= ^__obf_3da2e5262e6dbb1b
	}
	return __obf_7bae0b613da60dd3
}

// SetCustomStructTag causes the Encoder to use a custom struct tag as
// fallback option if there is no msgpack tag.
func (__obf_7bae0b613da60dd3 *Encoder) SetCustomStructTag(__obf_ee8ab0888686da67 string) {
	__obf_7bae0b613da60dd3.__obf_603dff1ad8d49539 = __obf_ee8ab0888686da67
}

// SetOmitEmpty causes the Encoder to omit empty values by default.
func (__obf_7bae0b613da60dd3 *Encoder) SetOmitEmpty(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e |= __obf_eadbff0f95d7376e
	} else {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e &= ^__obf_eadbff0f95d7376e
	}
}

// UseArrayEncodedStructs causes the Encoder to encode Go structs as msgpack arrays.
func (__obf_7bae0b613da60dd3 *Encoder) UseArrayEncodedStructs(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e |= __obf_1724e39b1a63df3b
	} else {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e &= ^__obf_1724e39b1a63df3b
	}
}

// UseCompactEncoding causes the Encoder to chose the most compact encoding.
// For example, it allows to encode small Go int64 as msgpack int8 saving 7 bytes.
func (__obf_7bae0b613da60dd3 *Encoder) UseCompactInts(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e |= __obf_968e60dffbce5433
	} else {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e &= ^__obf_968e60dffbce5433
	}
}

// UseCompactFloats causes the Encoder to chose a compact integer encoding
// for floats that can be represented as integers.
func (__obf_7bae0b613da60dd3 *Encoder) UseCompactFloats(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e |= __obf_7a9e24d108f17fe5
	} else {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e &= ^__obf_7a9e24d108f17fe5
	}
}

// UseInternedStrings causes the Encoder to intern strings.
func (__obf_7bae0b613da60dd3 *Encoder) UseInternedStrings(__obf_a8129bb947347fed bool) {
	if __obf_a8129bb947347fed {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e |= __obf_1fcd090a6da01656
	} else {
		__obf_7bae0b613da60dd3.__obf_a15a642720e39c3e &= ^__obf_1fcd090a6da01656
	}
}

func (__obf_7bae0b613da60dd3 *Encoder) Encode(__obf_17732590722140e7 any) error {
	switch __obf_17732590722140e7 := __obf_17732590722140e7.(type) {
	case nil:
		return __obf_7bae0b613da60dd3.EncodeNil()
	case string:
		return __obf_7bae0b613da60dd3.EncodeString(__obf_17732590722140e7)
	case []byte:
		return __obf_7bae0b613da60dd3.EncodeBytes(__obf_17732590722140e7)
	case int:
		return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_17732590722140e7))
	case int64:
		return __obf_7bae0b613da60dd3.__obf_545a172b23c8efda(__obf_17732590722140e7)
	case uint:
		return __obf_7bae0b613da60dd3.EncodeUint(uint64(__obf_17732590722140e7))
	case uint64:
		return __obf_7bae0b613da60dd3.__obf_70cf4747f5711e55(__obf_17732590722140e7)
	case bool:
		return __obf_7bae0b613da60dd3.EncodeBool(__obf_17732590722140e7)
	case float32:
		return __obf_7bae0b613da60dd3.EncodeFloat32(__obf_17732590722140e7)
	case float64:
		return __obf_7bae0b613da60dd3.EncodeFloat64(__obf_17732590722140e7)
	case time.Duration:
		return __obf_7bae0b613da60dd3.__obf_545a172b23c8efda(int64(__obf_17732590722140e7))
	case time.Time:
		return __obf_7bae0b613da60dd3.EncodeTime(__obf_17732590722140e7)
	}
	return __obf_7bae0b613da60dd3.EncodeValue(reflect.ValueOf(__obf_17732590722140e7))
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeMulti(__obf_17732590722140e7 ...any) error {
	for _, __obf_7c0d6551e0939afa := range __obf_17732590722140e7 {
		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.Encode(__obf_7c0d6551e0939afa); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeValue(__obf_17732590722140e7 reflect.Value) error {
	__obf_69331012fc3414ab := __obf_d4a1e60c459c24e4(__obf_17732590722140e7.Type())
	return __obf_69331012fc3414ab(__obf_7bae0b613da60dd3, __obf_17732590722140e7)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeNil() error {
	return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(Nil)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeBool(__obf_1926a1e373f9e647 bool) error {
	if __obf_1926a1e373f9e647 {
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(True)
	}
	return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(False)
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeDuration(__obf_dcaad1165bb07af9 time.Duration) error {
	return __obf_7bae0b613da60dd3.EncodeInt(int64(__obf_dcaad1165bb07af9))
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_47b86f684323fc33(__obf_ecf6d2d3253a058d byte) error {
	return __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.WriteByte(__obf_ecf6d2d3253a058d)
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_b0daaf664cd5cdfb(__obf_a7fd7acf2bd4435f []byte) error {
	_, __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.Write(__obf_a7fd7acf2bd4435f)
	return __obf_0feb3528b7b709ec
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_ff0da04198a8b5e2(__obf_a93d004caac96500 string) error {
	_, __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.Write(__obf_26e7c5987cb4f459(__obf_a93d004caac96500))
	return __obf_0feb3528b7b709ec
}
