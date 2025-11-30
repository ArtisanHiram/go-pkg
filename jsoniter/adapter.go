package __obf_703d23ba520c3295

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
func Unmarshal(__obf_c180417ee204f8c5 []byte, __obf_9cfb140618a077d4 any) error {
	return ConfigDefault.Unmarshal(__obf_c180417ee204f8c5,

		// UnmarshalFromString is a convenient method to read from string instead of []byte
		__obf_9cfb140618a077d4)
}

func UnmarshalFromString(__obf_44b48c26051f8033 string, __obf_9cfb140618a077d4 any) error {
	return ConfigDefault.UnmarshalFromString(__obf_44b48c26051f8033,

		// Get quick method to get value from deeply nested JSON structure
		__obf_9cfb140618a077d4)
}

func Get(__obf_c180417ee204f8c5 []byte, __obf_267bdf615cb1b310 ...any) Any {
	return ConfigDefault.Get(__obf_c180417ee204f8c5,

		// Marshal adapts to json/encoding Marshal API
		//
		// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
		// Refer to https://godoc.org/encoding/json#Marshal for more information
		__obf_267bdf615cb1b310...)
}

func Marshal(__obf_9cfb140618a077d4 any) ([]byte, error) {
	return ConfigDefault.Marshal(__obf_9cfb140618a077d4)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(__obf_9cfb140618a077d4 any, __obf_c5fd316f3c4457de, __obf_0d119fc055fc9f67 string) ([]byte, error) {
	return ConfigDefault.MarshalIndent(__obf_9cfb140618a077d4,

		// MarshalToString convenient method to write as string instead of []byte
		__obf_c5fd316f3c4457de, __obf_0d119fc055fc9f67)
}

func MarshalToString(__obf_9cfb140618a077d4 any) (string, error) {
	return ConfigDefault.MarshalToString(__obf_9cfb140618a077d4)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(__obf_3943589abf9d9fb6 io.Reader) *Decoder {
	return ConfigDefault.NewDecoder(__obf_3943589abf9d9fb6)
}

// Decoder reads and decodes JSON values from an input stream.
// Decoder provides identical APIs with json/stream Decoder (Token() and UseNumber() are in progress)
type Decoder struct {
	__obf_47edab4c16a2d88d *Iterator
}

// Decode decode JSON into any
func (__obf_ecca231db1e027ea *Decoder) Decode(__obf_02ba706f4f104d71 any) error {
	if __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11 == __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5 && __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_3943589abf9d9fb6 != nil {
		if !__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_e147dd47ce22b671() {
			return io.EOF
		}
	}
	__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.
		ReadVal(__obf_02ba706f4f104d71)
	__obf_e56742d6e52f642e := __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.Error
	if __obf_e56742d6e52f642e == io.EOF {
		return nil
	}
	return __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.Error
}

// More is there more?
func (__obf_ecca231db1e027ea *Decoder) More() bool {
	__obf_47edab4c16a2d88d := __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d
	if __obf_47edab4c16a2d88d.Error != nil {
		return false
	}
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == 0 {
		return false
	}
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	return __obf_bd08f5161fff294a != ']' && __obf_bd08f5161fff294a != '}'
}

// Buffered remaining buffer
func (__obf_ecca231db1e027ea *Decoder) Buffered() io.Reader {
	__obf_87dfc2abf099a36b := __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_a065f8e0da5f5952[__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_da6aa1cfbd772c11:__obf_ecca231db1e027ea.

		// UseNumber causes the Decoder to unmarshal a number into an any as a
		// Number instead of as a float64.
		__obf_47edab4c16a2d88d.__obf_687f71f0b8c00ef5]
	return bytes.NewReader(__obf_87dfc2abf099a36b)
}

func (__obf_ecca231db1e027ea *Decoder) UseNumber() {
	__obf_25bd4f754a37b862 := __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea
	__obf_25bd4f754a37b862.
		UseNumber = true
	__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862 = __obf_25bd4f754a37b862.__obf_06ac4a8116bedc14(__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.

		// DisallowUnknownFields causes the Decoder to return an error when the destination
		// is a struct and the input contains object keys which do not match any
		// non-ignored, exported fields in the destination.
		__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b)
}

func (__obf_ecca231db1e027ea *Decoder) DisallowUnknownFields() {
	__obf_25bd4f754a37b862 := __obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea
	__obf_25bd4f754a37b862.
		DisallowUnknownFields = true
	__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.__obf_25bd4f754a37b862 = __obf_25bd4f754a37b862.__obf_06ac4a8116bedc14(__obf_ecca231db1e027ea.__obf_47edab4c16a2d88d.

		// NewEncoder same as json.NewEncoder
		__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b)
}

func NewEncoder(__obf_8d0bd0f6d68b0c59 io.Writer) *Encoder {
	return ConfigDefault.NewEncoder(__obf_8d0bd0f6d68b0c59)
}

// Encoder same as json.Encoder
type Encoder struct {
	__obf_9a34f62857fb1d1d *Stream
}

// Encode encode any as JSON to io.Writer
func (__obf_ecca231db1e027ea *Encoder) Encode(__obf_b924538fffe5fd64 any) error {
	__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.
		WriteVal(__obf_b924538fffe5fd64)
	__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.
		WriteRaw("\n")
	__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.
		Flush()
	return __obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.Error
}

// SetIndent set the indention. Prefix is not supported
func (__obf_ecca231db1e027ea *Encoder) SetIndent(__obf_c5fd316f3c4457de, __obf_0d119fc055fc9f67 string) {
	__obf_e24a8f8aa618a582 := __obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea
	__obf_e24a8f8aa618a582.
		IndentionStep = len(__obf_0d119fc055fc9f67)
	__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862 = __obf_e24a8f8aa618a582.__obf_06ac4a8116bedc14(__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.

		// SetEscapeHTML escape html by default, set to false to disable
		__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b)
}

func (__obf_ecca231db1e027ea *Encoder) SetEscapeHTML(__obf_3ea234a937346d00 bool) {
	__obf_e24a8f8aa618a582 := __obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.__obf_4d2d9c0287da7aea
	__obf_e24a8f8aa618a582.
		EscapeHTML = __obf_3ea234a937346d00
	__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862 = __obf_e24a8f8aa618a582.__obf_06ac4a8116bedc14(__obf_ecca231db1e027ea.__obf_9a34f62857fb1d1d.

		// Valid reports whether data is a valid JSON encoding.
		__obf_25bd4f754a37b862.__obf_324ee50b8c0b2f3b)
}

func Valid(__obf_c180417ee204f8c5 []byte) bool {
	return ConfigDefault.Valid(__obf_c180417ee204f8c5)
}
