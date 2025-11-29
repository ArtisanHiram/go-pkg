package __obf_5b802ce8d9ba56d6

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
func Unmarshal(__obf_5d3dff745e2b805b []byte, __obf_91fc460a138dcae7 any) error {
	return ConfigDefault.Unmarshal(__obf_5d3dff745e2b805b,

		// UnmarshalFromString is a convenient method to read from string instead of []byte
		__obf_91fc460a138dcae7)
}

func UnmarshalFromString(__obf_12c21b79fa86dcba string, __obf_91fc460a138dcae7 any) error {
	return ConfigDefault.UnmarshalFromString(__obf_12c21b79fa86dcba,

		// Get quick method to get value from deeply nested JSON structure
		__obf_91fc460a138dcae7)
}

func Get(__obf_5d3dff745e2b805b []byte, __obf_c5d71353f4393b3c ...any) Any {
	return ConfigDefault.Get(__obf_5d3dff745e2b805b,

		// Marshal adapts to json/encoding Marshal API
		//
		// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
		// Refer to https://godoc.org/encoding/json#Marshal for more information
		__obf_c5d71353f4393b3c...)
}

func Marshal(__obf_91fc460a138dcae7 any) ([]byte, error) {
	return ConfigDefault.Marshal(__obf_91fc460a138dcae7)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(__obf_91fc460a138dcae7 any, __obf_f3dd783c02964acf, __obf_8e554c17d6bbd080 string) ([]byte, error) {
	return ConfigDefault.MarshalIndent(__obf_91fc460a138dcae7,

		// MarshalToString convenient method to write as string instead of []byte
		__obf_f3dd783c02964acf, __obf_8e554c17d6bbd080)
}

func MarshalToString(__obf_91fc460a138dcae7 any) (string, error) {
	return ConfigDefault.MarshalToString(__obf_91fc460a138dcae7)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(__obf_1fcab394164d9b41 io.Reader) *Decoder {
	return ConfigDefault.NewDecoder(__obf_1fcab394164d9b41)
}

// Decoder reads and decodes JSON values from an input stream.
// Decoder provides identical APIs with json/stream Decoder (Token() and UseNumber() are in progress)
type Decoder struct {
	__obf_67008a6a9e5ba828 *Iterator
}

// Decode decode JSON into any
func (__obf_b78f5c04b6b5bb1a *Decoder) Decode(__obf_f9b1526531f3c024 any) error {
	if __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36 == __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_3a36550914545c79 && __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_1fcab394164d9b41 != nil {
		if !__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_538f3d3337e0395f() {
			return io.EOF
		}
	}
	__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.
		ReadVal(__obf_f9b1526531f3c024)
	__obf_6d071d48f3b321ad := __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.Error
	if __obf_6d071d48f3b321ad == io.EOF {
		return nil
	}
	return __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.Error
}

// More is there more?
func (__obf_b78f5c04b6b5bb1a *Decoder) More() bool {
	__obf_67008a6a9e5ba828 := __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828
	if __obf_67008a6a9e5ba828.Error != nil {
		return false
	}
	__obf_dab9baaadfa7c8c2 := __obf_67008a6a9e5ba828.__obf_b781a59d5a0d2490()
	if __obf_dab9baaadfa7c8c2 == 0 {
		return false
	}
	__obf_67008a6a9e5ba828.__obf_3284a1eaa2a0abb6()
	return __obf_dab9baaadfa7c8c2 != ']' && __obf_dab9baaadfa7c8c2 != '}'
}

// Buffered remaining buffer
func (__obf_b78f5c04b6b5bb1a *Decoder) Buffered() io.Reader {
	__obf_1a7ec2cfcbfcf807 := __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_9fc06d9180f0daca[__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_14babd6f9a55bd36:__obf_b78f5c04b6b5bb1a.

		// UseNumber causes the Decoder to unmarshal a number into an any as a
		// Number instead of as a float64.
		__obf_67008a6a9e5ba828.__obf_3a36550914545c79]
	return bytes.NewReader(__obf_1a7ec2cfcbfcf807)
}

func (__obf_b78f5c04b6b5bb1a *Decoder) UseNumber() {
	__obf_dca106e1cdf85392 := __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_97c6754e53034071
	__obf_dca106e1cdf85392.
		UseNumber = true
	__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392 = __obf_dca106e1cdf85392.__obf_25a90ed48fcf6327(__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.

		// DisallowUnknownFields causes the Decoder to return an error when the destination
		// is a struct and the input contains object keys which do not match any
		// non-ignored, exported fields in the destination.
		__obf_dca106e1cdf85392.__obf_251d7593467966e4)
}

func (__obf_b78f5c04b6b5bb1a *Decoder) DisallowUnknownFields() {
	__obf_dca106e1cdf85392 := __obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392.__obf_97c6754e53034071
	__obf_dca106e1cdf85392.
		DisallowUnknownFields = true
	__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.__obf_dca106e1cdf85392 = __obf_dca106e1cdf85392.__obf_25a90ed48fcf6327(__obf_b78f5c04b6b5bb1a.__obf_67008a6a9e5ba828.

		// NewEncoder same as json.NewEncoder
		__obf_dca106e1cdf85392.__obf_251d7593467966e4)
}

func NewEncoder(__obf_3dbedd483fe5a38a io.Writer) *Encoder {
	return ConfigDefault.NewEncoder(__obf_3dbedd483fe5a38a)
}

// Encoder same as json.Encoder
type Encoder struct {
	__obf_00fc491caa967a74 *Stream
}

// Encode encode any as JSON to io.Writer
func (__obf_b78f5c04b6b5bb1a *Encoder) Encode(__obf_5406b11e33b9d1d3 any) error {
	__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.
		WriteVal(__obf_5406b11e33b9d1d3)
	__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.
		WriteRaw("\n")
	__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.
		Flush()
	return __obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.Error
}

// SetIndent set the indention. Prefix is not supported
func (__obf_b78f5c04b6b5bb1a *Encoder) SetIndent(__obf_f3dd783c02964acf, __obf_8e554c17d6bbd080 string) {
	__obf_02f4c89f127e623f := __obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_97c6754e53034071
	__obf_02f4c89f127e623f.
		IndentionStep = len(__obf_8e554c17d6bbd080)
	__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.__obf_dca106e1cdf85392 = __obf_02f4c89f127e623f.__obf_25a90ed48fcf6327(__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.

		// SetEscapeHTML escape html by default, set to false to disable
		__obf_dca106e1cdf85392.__obf_251d7593467966e4)
}

func (__obf_b78f5c04b6b5bb1a *Encoder) SetEscapeHTML(__obf_0e412ce07ebed81c bool) {
	__obf_02f4c89f127e623f := __obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.__obf_dca106e1cdf85392.__obf_97c6754e53034071
	__obf_02f4c89f127e623f.
		EscapeHTML = __obf_0e412ce07ebed81c
	__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.__obf_dca106e1cdf85392 = __obf_02f4c89f127e623f.__obf_25a90ed48fcf6327(__obf_b78f5c04b6b5bb1a.__obf_00fc491caa967a74.

		// Valid reports whether data is a valid JSON encoding.
		__obf_dca106e1cdf85392.__obf_251d7593467966e4)
}

func Valid(__obf_5d3dff745e2b805b []byte) bool {
	return ConfigDefault.Valid(__obf_5d3dff745e2b805b)
}
