package __obf_c7b79b12b33d8238

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
func Unmarshal(__obf_1d34d01a9b83218a []byte, __obf_bc7196b82ffbf1d3 any) error {
	return ConfigDefault.Unmarshal(__obf_1d34d01a9b83218a,

		// UnmarshalFromString is a convenient method to read from string instead of []byte
		__obf_bc7196b82ffbf1d3)
}

func UnmarshalFromString(__obf_a3eaafc22faf1644 string, __obf_bc7196b82ffbf1d3 any) error {
	return ConfigDefault.UnmarshalFromString(__obf_a3eaafc22faf1644,

		// Get quick method to get value from deeply nested JSON structure
		__obf_bc7196b82ffbf1d3)
}

func Get(__obf_1d34d01a9b83218a []byte, __obf_5dde9fb4007294d4 ...any) Any {
	return ConfigDefault.Get(__obf_1d34d01a9b83218a,

		// Marshal adapts to json/encoding Marshal API
		//
		// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
		// Refer to https://godoc.org/encoding/json#Marshal for more information
		__obf_5dde9fb4007294d4...)
}

func Marshal(__obf_bc7196b82ffbf1d3 any) ([]byte, error) {
	return ConfigDefault.Marshal(__obf_bc7196b82ffbf1d3)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(__obf_bc7196b82ffbf1d3 any, __obf_5de9ece8fa3a16e3, __obf_be0b45c15e61f86b string) ([]byte, error) {
	return ConfigDefault.MarshalIndent(__obf_bc7196b82ffbf1d3,

		// MarshalToString convenient method to write as string instead of []byte
		__obf_5de9ece8fa3a16e3, __obf_be0b45c15e61f86b)
}

func MarshalToString(__obf_bc7196b82ffbf1d3 any) (string, error) {
	return ConfigDefault.MarshalToString(__obf_bc7196b82ffbf1d3)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(__obf_801a7eebc4303617 io.Reader) *Decoder {
	return ConfigDefault.NewDecoder(__obf_801a7eebc4303617)
}

// Decoder reads and decodes JSON values from an input stream.
// Decoder provides identical APIs with json/stream Decoder (Token() and UseNumber() are in progress)
type Decoder struct {
	__obf_0da8c843dad13139 *Iterator
}

// Decode decode JSON into any
func (__obf_2d6cfd588b796583 *Decoder) Decode(__obf_d6e2df4782353c64 any) error {
	if __obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2 == __obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_840246080559848c && __obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_801a7eebc4303617 != nil {
		if !__obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_baaf768e13842431() {
			return io.EOF
		}
	}
	__obf_2d6cfd588b796583.__obf_0da8c843dad13139.
		ReadVal(__obf_d6e2df4782353c64)
	__obf_ea0680f8b461a85b := __obf_2d6cfd588b796583.__obf_0da8c843dad13139.Error
	if __obf_ea0680f8b461a85b == io.EOF {
		return nil
	}
	return __obf_2d6cfd588b796583.__obf_0da8c843dad13139.Error
}

// More is there more?
func (__obf_2d6cfd588b796583 *Decoder) More() bool {
	__obf_0da8c843dad13139 := __obf_2d6cfd588b796583.__obf_0da8c843dad13139
	if __obf_0da8c843dad13139.Error != nil {
		return false
	}
	__obf_12577c03a12f0d2c := __obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	if __obf_12577c03a12f0d2c == 0 {
		return false
	}
	__obf_0da8c843dad13139.__obf_a675366c80290b83()
	return __obf_12577c03a12f0d2c != ']' && __obf_12577c03a12f0d2c != '}'
}

// Buffered remaining buffer
func (__obf_2d6cfd588b796583 *Decoder) Buffered() io.Reader {
	__obf_2565c9afb6012ebe := __obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_684d738bc3ac851a[__obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_6a63c9ac34fe84e2:__obf_2d6cfd588b796583.

		// UseNumber causes the Decoder to unmarshal a number into an any as a
		// Number instead of as a float64.
		__obf_0da8c843dad13139.__obf_840246080559848c]
	return bytes.NewReader(__obf_2565c9afb6012ebe)
}

func (__obf_2d6cfd588b796583 *Decoder) UseNumber() {
	__obf_c52e0bcfad4b8b71 := __obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc
	__obf_c52e0bcfad4b8b71.
		UseNumber = true
	__obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71 = __obf_c52e0bcfad4b8b71.__obf_ec5a9b5887aa3ec6(__obf_2d6cfd588b796583.__obf_0da8c843dad13139.

		// DisallowUnknownFields causes the Decoder to return an error when the destination
		// is a struct and the input contains object keys which do not match any
		// non-ignored, exported fields in the destination.
		__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb)
}

func (__obf_2d6cfd588b796583 *Decoder) DisallowUnknownFields() {
	__obf_c52e0bcfad4b8b71 := __obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc
	__obf_c52e0bcfad4b8b71.
		DisallowUnknownFields = true
	__obf_2d6cfd588b796583.__obf_0da8c843dad13139.__obf_c52e0bcfad4b8b71 = __obf_c52e0bcfad4b8b71.__obf_ec5a9b5887aa3ec6(__obf_2d6cfd588b796583.__obf_0da8c843dad13139.

		// NewEncoder same as json.NewEncoder
		__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb)
}

func NewEncoder(__obf_29b48e4cc46aebd6 io.Writer) *Encoder {
	return ConfigDefault.NewEncoder(__obf_29b48e4cc46aebd6)
}

// Encoder same as json.Encoder
type Encoder struct {
	__obf_d8f50bcfe87d8b47 *Stream
}

// Encode encode any as JSON to io.Writer
func (__obf_2d6cfd588b796583 *Encoder) Encode(__obf_35accd096612ebe4 any) error {
	__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.
		WriteVal(__obf_35accd096612ebe4)
	__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.
		WriteRaw("\n")
	__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.
		Flush()
	return __obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.Error
}

// SetIndent set the indention. Prefix is not supported
func (__obf_2d6cfd588b796583 *Encoder) SetIndent(__obf_5de9ece8fa3a16e3, __obf_be0b45c15e61f86b string) {
	__obf_ecb992e78814f9e4 := __obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc
	__obf_ecb992e78814f9e4.
		IndentionStep = len(__obf_be0b45c15e61f86b)
	__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71 = __obf_ecb992e78814f9e4.__obf_ec5a9b5887aa3ec6(__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.

		// SetEscapeHTML escape html by default, set to false to disable
		__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb)
}

func (__obf_2d6cfd588b796583 *Encoder) SetEscapeHTML(__obf_3f302bac9ebc5e76 bool) {
	__obf_ecb992e78814f9e4 := __obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71.__obf_790a57f66978b3bc
	__obf_ecb992e78814f9e4.
		EscapeHTML = __obf_3f302bac9ebc5e76
	__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.__obf_c52e0bcfad4b8b71 = __obf_ecb992e78814f9e4.__obf_ec5a9b5887aa3ec6(__obf_2d6cfd588b796583.__obf_d8f50bcfe87d8b47.

		// Valid reports whether data is a valid JSON encoding.
		__obf_c52e0bcfad4b8b71.__obf_3c82f4beb30882eb)
}

func Valid(__obf_1d34d01a9b83218a []byte) bool {
	return ConfigDefault.Valid(__obf_1d34d01a9b83218a)
}
