package __obf_c3cd04a15c56f32f

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
func Unmarshal(__obf_905295f6bd29aafe []byte, __obf_f967fc9e700d5e33 any) error {
	return ConfigDefault.Unmarshal(__obf_905295f6bd29aafe,

		// UnmarshalFromString is a convenient method to read from string instead of []byte
		__obf_f967fc9e700d5e33)
}

func UnmarshalFromString(__obf_2d944450d48e5810 string, __obf_f967fc9e700d5e33 any) error {
	return ConfigDefault.UnmarshalFromString(__obf_2d944450d48e5810,

		// Get quick method to get value from deeply nested JSON structure
		__obf_f967fc9e700d5e33)
}

func Get(__obf_905295f6bd29aafe []byte, __obf_b90e4ca332607787 ...any) Any {
	return ConfigDefault.Get(__obf_905295f6bd29aafe,

		// Marshal adapts to json/encoding Marshal API
		//
		// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
		// Refer to https://godoc.org/encoding/json#Marshal for more information
		__obf_b90e4ca332607787...)
}

func Marshal(__obf_f967fc9e700d5e33 any) ([]byte, error) {
	return ConfigDefault.Marshal(__obf_f967fc9e700d5e33)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(__obf_f967fc9e700d5e33 any, __obf_f517983aa5897f07, __obf_30653456d76688b8 string) ([]byte, error) {
	return ConfigDefault.MarshalIndent(__obf_f967fc9e700d5e33,

		// MarshalToString convenient method to write as string instead of []byte
		__obf_f517983aa5897f07, __obf_30653456d76688b8)
}

func MarshalToString(__obf_f967fc9e700d5e33 any) (string, error) {
	return ConfigDefault.MarshalToString(__obf_f967fc9e700d5e33)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(__obf_a539369e10727e42 io.Reader) *Decoder {
	return ConfigDefault.NewDecoder(__obf_a539369e10727e42)
}

// Decoder reads and decodes JSON values from an input stream.
// Decoder provides identical APIs with json/stream Decoder (Token() and UseNumber() are in progress)
type Decoder struct {
	__obf_edd9fe48d39445e4 *Iterator
}

// Decode decode JSON into any
func (__obf_41873dcb77e49290 *Decoder) Decode(__obf_50acae8c0a871ba1 any) error {
	if __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_edd3c3885447229b == __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57 && __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_a539369e10727e42 != nil {
		if !__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_35e704c169a88a74() {
			return io.EOF
		}
	}
	__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.
		ReadVal(__obf_50acae8c0a871ba1)
	__obf_5966ecc5edb9b17e := __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.Error
	if __obf_5966ecc5edb9b17e == io.EOF {
		return nil
	}
	return __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.Error
}

// More is there more?
func (__obf_41873dcb77e49290 *Decoder) More() bool {
	__obf_edd9fe48d39445e4 := __obf_41873dcb77e49290.__obf_edd9fe48d39445e4
	if __obf_edd9fe48d39445e4.Error != nil {
		return false
	}
	__obf_0c1bc1e511a43120 := __obf_edd9fe48d39445e4.__obf_a26eeeac1d6f5a11()
	if __obf_0c1bc1e511a43120 == 0 {
		return false
	}
	__obf_edd9fe48d39445e4.__obf_a0622aded2d86ded()
	return __obf_0c1bc1e511a43120 != ']' && __obf_0c1bc1e511a43120 != '}'
}

// Buffered remaining buffer
func (__obf_41873dcb77e49290 *Decoder) Buffered() io.Reader {
	__obf_17fc21b94a0188cf := __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_ace979f6622823f3[__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_edd3c3885447229b:__obf_41873dcb77e49290.

		// UseNumber causes the Decoder to unmarshal a number into an any as a
		// Number instead of as a float64.
		__obf_edd9fe48d39445e4.__obf_3a475d1c1ae2cd57]
	return bytes.NewReader(__obf_17fc21b94a0188cf)
}

func (__obf_41873dcb77e49290 *Decoder) UseNumber() {
	__obf_0c8065c219ccf0ab := __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab
	__obf_0c8065c219ccf0ab.
		UseNumber = true
	__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab = __obf_0c8065c219ccf0ab.__obf_c3344b4ada41a413(__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.

		// DisallowUnknownFields causes the Decoder to return an error when the destination
		// is a struct and the input contains object keys which do not match any
		// non-ignored, exported fields in the destination.
		__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6)
}

func (__obf_41873dcb77e49290 *Decoder) DisallowUnknownFields() {
	__obf_0c8065c219ccf0ab := __obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab
	__obf_0c8065c219ccf0ab.
		DisallowUnknownFields = true
	__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.__obf_0c8065c219ccf0ab = __obf_0c8065c219ccf0ab.__obf_c3344b4ada41a413(__obf_41873dcb77e49290.__obf_edd9fe48d39445e4.

		// NewEncoder same as json.NewEncoder
		__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6)
}

func NewEncoder(__obf_b1699a146b20b28e io.Writer) *Encoder {
	return ConfigDefault.NewEncoder(__obf_b1699a146b20b28e)
}

// Encoder same as json.Encoder
type Encoder struct {
	__obf_2361f5214e84c60f *Stream
}

// Encode encode any as JSON to io.Writer
func (__obf_41873dcb77e49290 *Encoder) Encode(__obf_d59813f23ad740a8 any) error {
	__obf_41873dcb77e49290.__obf_2361f5214e84c60f.
		WriteVal(__obf_d59813f23ad740a8)
	__obf_41873dcb77e49290.__obf_2361f5214e84c60f.
		WriteRaw("\n")
	__obf_41873dcb77e49290.__obf_2361f5214e84c60f.
		Flush()
	return __obf_41873dcb77e49290.__obf_2361f5214e84c60f.Error
}

// SetIndent set the indention. Prefix is not supported
func (__obf_41873dcb77e49290 *Encoder) SetIndent(__obf_f517983aa5897f07, __obf_30653456d76688b8 string) {
	__obf_e80db87a4554a9f5 := __obf_41873dcb77e49290.__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab
	__obf_e80db87a4554a9f5.
		IndentionStep = len(__obf_30653456d76688b8)
	__obf_41873dcb77e49290.__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab = __obf_e80db87a4554a9f5.__obf_c3344b4ada41a413(__obf_41873dcb77e49290.__obf_2361f5214e84c60f.

		// SetEscapeHTML escape html by default, set to false to disable
		__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6)
}

func (__obf_41873dcb77e49290 *Encoder) SetEscapeHTML(__obf_98a068a278faf6e8 bool) {
	__obf_e80db87a4554a9f5 := __obf_41873dcb77e49290.__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab.__obf_6f7054bdfbca17ab
	__obf_e80db87a4554a9f5.
		EscapeHTML = __obf_98a068a278faf6e8
	__obf_41873dcb77e49290.__obf_2361f5214e84c60f.__obf_0c8065c219ccf0ab = __obf_e80db87a4554a9f5.__obf_c3344b4ada41a413(__obf_41873dcb77e49290.__obf_2361f5214e84c60f.

		// Valid reports whether data is a valid JSON encoding.
		__obf_0c8065c219ccf0ab.__obf_a1eac24bcc56f0a6)
}

func Valid(__obf_905295f6bd29aafe []byte) bool {
	return ConfigDefault.Valid(__obf_905295f6bd29aafe)
}
