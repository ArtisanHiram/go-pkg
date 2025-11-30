package __obf_030d39f7a8de96c6

import (
	"bytes"
	"io"
)

// RawMessage to make replace json with jsoniter
type RawMessage []byte

// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information
func Unmarshal(__obf_c28f0e30eceb6d4a []byte, __obf_41b184145cfea25b any) error {
	return ConfigDefault.Unmarshal(__obf_c28f0e30eceb6d4a,

		// UnmarshalFromString is a convenient method to read from string instead of []byte
		__obf_41b184145cfea25b)
}

func UnmarshalFromString(__obf_428c3b4a95725c4a string, __obf_41b184145cfea25b any) error {
	return ConfigDefault.UnmarshalFromString(__obf_428c3b4a95725c4a,

		// Get quick method to get value from deeply nested JSON structure
		__obf_41b184145cfea25b)
}

func Get(__obf_c28f0e30eceb6d4a []byte, __obf_f77a9507782b919d ...any) Any {
	return ConfigDefault.Get(__obf_c28f0e30eceb6d4a,

		// Marshal adapts to json/encoding Marshal API
		//
		// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
		// Refer to https://godoc.org/encoding/json#Marshal for more information
		__obf_f77a9507782b919d...)
}

func Marshal(__obf_41b184145cfea25b any) ([]byte, error) {
	return ConfigDefault.Marshal(__obf_41b184145cfea25b)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(__obf_41b184145cfea25b any, __obf_834e1679b8081f46, __obf_2ee2bfac47d1636e string) ([]byte, error) {
	return ConfigDefault.MarshalIndent(__obf_41b184145cfea25b,

		// MarshalToString convenient method to write as string instead of []byte
		__obf_834e1679b8081f46, __obf_2ee2bfac47d1636e)
}

func MarshalToString(__obf_41b184145cfea25b any) (string, error) {
	return ConfigDefault.MarshalToString(__obf_41b184145cfea25b)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(__obf_7582c70e81d83895 io.Reader) *Decoder {
	return ConfigDefault.NewDecoder(__obf_7582c70e81d83895)
}

// Decoder reads and decodes JSON values from an input stream.
// Decoder provides identical APIs with json/stream Decoder (Token() and UseNumber() are in progress)
type Decoder struct {
	__obf_4ab56a99965952e7 *Iterator
}

// Decode decode JSON into any
func (__obf_95dec6f27d951053 *Decoder) Decode(__obf_eefa92392cc2442c any) error {
	if __obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_5e22d6270d27491f == __obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc && __obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_7582c70e81d83895 != nil {
		if !__obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_08a0b850abb3abbd() {
			return io.EOF
		}
	}
	__obf_95dec6f27d951053.__obf_4ab56a99965952e7.
		ReadVal(__obf_eefa92392cc2442c)
	__obf_fcc907dd69879566 := __obf_95dec6f27d951053.__obf_4ab56a99965952e7.Error
	if __obf_fcc907dd69879566 == io.EOF {
		return nil
	}
	return __obf_95dec6f27d951053.__obf_4ab56a99965952e7.Error
}

// More is there more?
func (__obf_95dec6f27d951053 *Decoder) More() bool {
	__obf_4ab56a99965952e7 := __obf_95dec6f27d951053.__obf_4ab56a99965952e7
	if __obf_4ab56a99965952e7.Error != nil {
		return false
	}
	__obf_24309b3b0ff9ba22 := __obf_4ab56a99965952e7.__obf_61df301d1f67ad73()
	if __obf_24309b3b0ff9ba22 == 0 {
		return false
	}
	__obf_4ab56a99965952e7.__obf_ba6d0be9a7ab6ae4()
	return __obf_24309b3b0ff9ba22 != ']' && __obf_24309b3b0ff9ba22 != '}'
}

// Buffered remaining buffer
func (__obf_95dec6f27d951053 *Decoder) Buffered() io.Reader {
	__obf_2257b1725dfd65c5 := __obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_0b1656d7ef4bd9df[__obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_5e22d6270d27491f:__obf_95dec6f27d951053.

		// UseNumber causes the Decoder to unmarshal a number into an any as a
		// Number instead of as a float64.
		__obf_4ab56a99965952e7.__obf_7c17042d9b4e73cc]
	return bytes.NewReader(__obf_2257b1725dfd65c5)
}

func (__obf_95dec6f27d951053 *Decoder) UseNumber() {
	__obf_679611dc56ff465b := __obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_fb97917857c6c8da
	__obf_679611dc56ff465b.
		UseNumber = true
	__obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_679611dc56ff465b = __obf_679611dc56ff465b.__obf_38802709ddf21ec6(__obf_95dec6f27d951053.__obf_4ab56a99965952e7.

		// DisallowUnknownFields causes the Decoder to return an error when the destination
		// is a struct and the input contains object keys which do not match any
		// non-ignored, exported fields in the destination.
		__obf_679611dc56ff465b.__obf_6621255bc1f80c8e)
}

func (__obf_95dec6f27d951053 *Decoder) DisallowUnknownFields() {
	__obf_679611dc56ff465b := __obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_679611dc56ff465b.__obf_fb97917857c6c8da
	__obf_679611dc56ff465b.
		DisallowUnknownFields = true
	__obf_95dec6f27d951053.__obf_4ab56a99965952e7.__obf_679611dc56ff465b = __obf_679611dc56ff465b.__obf_38802709ddf21ec6(__obf_95dec6f27d951053.__obf_4ab56a99965952e7.

		// NewEncoder same as json.NewEncoder
		__obf_679611dc56ff465b.__obf_6621255bc1f80c8e)
}

func NewEncoder(__obf_ae48508e054620bd io.Writer) *Encoder {
	return ConfigDefault.NewEncoder(__obf_ae48508e054620bd)
}

// Encoder same as json.Encoder
type Encoder struct {
	__obf_8a2c51fe22775655 *Stream
}

// Encode encode any as JSON to io.Writer
func (__obf_95dec6f27d951053 *Encoder) Encode(__obf_efaf2719b44ad8ac any) error {
	__obf_95dec6f27d951053.__obf_8a2c51fe22775655.
		WriteVal(__obf_efaf2719b44ad8ac)
	__obf_95dec6f27d951053.__obf_8a2c51fe22775655.
		WriteRaw("\n")
	__obf_95dec6f27d951053.__obf_8a2c51fe22775655.
		Flush()
	return __obf_95dec6f27d951053.__obf_8a2c51fe22775655.Error
}

// SetIndent set the indention. Prefix is not supported
func (__obf_95dec6f27d951053 *Encoder) SetIndent(__obf_834e1679b8081f46, __obf_2ee2bfac47d1636e string) {
	__obf_aefc68b53b020259 := __obf_95dec6f27d951053.__obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_fb97917857c6c8da
	__obf_aefc68b53b020259.
		IndentionStep = len(__obf_2ee2bfac47d1636e)
	__obf_95dec6f27d951053.__obf_8a2c51fe22775655.__obf_679611dc56ff465b = __obf_aefc68b53b020259.__obf_38802709ddf21ec6(__obf_95dec6f27d951053.__obf_8a2c51fe22775655.

		// SetEscapeHTML escape html by default, set to false to disable
		__obf_679611dc56ff465b.__obf_6621255bc1f80c8e)
}

func (__obf_95dec6f27d951053 *Encoder) SetEscapeHTML(__obf_ea893fc01ce71521 bool) {
	__obf_aefc68b53b020259 := __obf_95dec6f27d951053.__obf_8a2c51fe22775655.__obf_679611dc56ff465b.__obf_fb97917857c6c8da
	__obf_aefc68b53b020259.
		EscapeHTML = __obf_ea893fc01ce71521
	__obf_95dec6f27d951053.__obf_8a2c51fe22775655.__obf_679611dc56ff465b = __obf_aefc68b53b020259.__obf_38802709ddf21ec6(__obf_95dec6f27d951053.__obf_8a2c51fe22775655.

		// Valid reports whether data is a valid JSON encoding.
		__obf_679611dc56ff465b.__obf_6621255bc1f80c8e)
}

func Valid(__obf_c28f0e30eceb6d4a []byte) bool {
	return ConfigDefault.Valid(__obf_c28f0e30eceb6d4a)
}
