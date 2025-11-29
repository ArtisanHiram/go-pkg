package __obf_91620b895eeff9ed

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
func Unmarshal(__obf_2894589095c323b4 []byte, __obf_97f02ddd0770463f any) error {
	return ConfigDefault.Unmarshal(__obf_2894589095c323b4,

		// UnmarshalFromString is a convenient method to read from string instead of []byte
		__obf_97f02ddd0770463f)
}

func UnmarshalFromString(__obf_e91bd2feb751e4f1 string, __obf_97f02ddd0770463f any) error {
	return ConfigDefault.UnmarshalFromString(__obf_e91bd2feb751e4f1,

		// Get quick method to get value from deeply nested JSON structure
		__obf_97f02ddd0770463f)
}

func Get(__obf_2894589095c323b4 []byte, __obf_82c50fcbfc9ab432 ...any) Any {
	return ConfigDefault.Get(__obf_2894589095c323b4,

		// Marshal adapts to json/encoding Marshal API
		//
		// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
		// Refer to https://godoc.org/encoding/json#Marshal for more information
		__obf_82c50fcbfc9ab432...)
}

func Marshal(__obf_97f02ddd0770463f any) ([]byte, error) {
	return ConfigDefault.Marshal(__obf_97f02ddd0770463f)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(__obf_97f02ddd0770463f any, __obf_6e20c91e70903582, __obf_6ab410c037fcd950 string) ([]byte, error) {
	return ConfigDefault.MarshalIndent(__obf_97f02ddd0770463f,

		// MarshalToString convenient method to write as string instead of []byte
		__obf_6e20c91e70903582, __obf_6ab410c037fcd950)
}

func MarshalToString(__obf_97f02ddd0770463f any) (string, error) {
	return ConfigDefault.MarshalToString(__obf_97f02ddd0770463f)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(__obf_fe90e88a3ba69cf3 io.Reader) *Decoder {
	return ConfigDefault.NewDecoder(__obf_fe90e88a3ba69cf3)
}

// Decoder reads and decodes JSON values from an input stream.
// Decoder provides identical APIs with json/stream Decoder (Token() and UseNumber() are in progress)
type Decoder struct {
	__obf_1bb30e8a74ed8233 *Iterator
}

// Decode decode JSON into any
func (__obf_c2d286e997a4a112 *Decoder) Decode(__obf_35136ef2ac0f94e4 any) error {
	if __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21 == __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae && __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_fe90e88a3ba69cf3 != nil {
		if !__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_e927802c539fc09d() {
			return io.EOF
		}
	}
	__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.
		ReadVal(__obf_35136ef2ac0f94e4)
	__obf_62967739bca1d11d := __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.Error
	if __obf_62967739bca1d11d == io.EOF {
		return nil
	}
	return __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.Error
}

// More is there more?
func (__obf_c2d286e997a4a112 *Decoder) More() bool {
	__obf_1bb30e8a74ed8233 := __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233
	if __obf_1bb30e8a74ed8233.Error != nil {
		return false
	}
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af == 0 {
		return false
	}
	__obf_1bb30e8a74ed8233.__obf_a163df67f9bb1c4b()
	return __obf_f16b4157911bc9af != ']' && __obf_f16b4157911bc9af != '}'
}

// Buffered remaining buffer
func (__obf_c2d286e997a4a112 *Decoder) Buffered() io.Reader {
	__obf_661d5d31d1b381f5 := __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_184433571fa55237[__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_a657fb48fcb34e21:__obf_c2d286e997a4a112.

		// UseNumber causes the Decoder to unmarshal a number into an any as a
		// Number instead of as a float64.
		__obf_1bb30e8a74ed8233.__obf_15d837671d2809ae]
	return bytes.NewReader(__obf_661d5d31d1b381f5)
}

func (__obf_c2d286e997a4a112 *Decoder) UseNumber() {
	__obf_892632c77e859037 := __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_e05f98c9e7cc9089
	__obf_892632c77e859037.
		UseNumber = true
	__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_892632c77e859037 = __obf_892632c77e859037.__obf_6f0de5d9837befa9(__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.

		// DisallowUnknownFields causes the Decoder to return an error when the destination
		// is a struct and the input contains object keys which do not match any
		// non-ignored, exported fields in the destination.
		__obf_892632c77e859037.__obf_b4dfae156c7ac26c)
}

func (__obf_c2d286e997a4a112 *Decoder) DisallowUnknownFields() {
	__obf_892632c77e859037 := __obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_e05f98c9e7cc9089
	__obf_892632c77e859037.
		DisallowUnknownFields = true
	__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.__obf_892632c77e859037 = __obf_892632c77e859037.__obf_6f0de5d9837befa9(__obf_c2d286e997a4a112.__obf_1bb30e8a74ed8233.

		// NewEncoder same as json.NewEncoder
		__obf_892632c77e859037.__obf_b4dfae156c7ac26c)
}

func NewEncoder(__obf_e4132b2c3ef9e14e io.Writer) *Encoder {
	return ConfigDefault.NewEncoder(__obf_e4132b2c3ef9e14e)
}

// Encoder same as json.Encoder
type Encoder struct {
	__obf_850a7457bb739a32 *Stream
}

// Encode encode any as JSON to io.Writer
func (__obf_c2d286e997a4a112 *Encoder) Encode(__obf_bbfd2ac8618a6f0c any) error {
	__obf_c2d286e997a4a112.__obf_850a7457bb739a32.
		WriteVal(__obf_bbfd2ac8618a6f0c)
	__obf_c2d286e997a4a112.__obf_850a7457bb739a32.
		WriteRaw("\n")
	__obf_c2d286e997a4a112.__obf_850a7457bb739a32.
		Flush()
	return __obf_c2d286e997a4a112.__obf_850a7457bb739a32.Error
}

// SetIndent set the indention. Prefix is not supported
func (__obf_c2d286e997a4a112 *Encoder) SetIndent(__obf_6e20c91e70903582, __obf_6ab410c037fcd950 string) {
	__obf_6e9f1b1176b96c20 := __obf_c2d286e997a4a112.__obf_850a7457bb739a32.__obf_892632c77e859037.__obf_e05f98c9e7cc9089
	__obf_6e9f1b1176b96c20.
		IndentionStep = len(__obf_6ab410c037fcd950)
	__obf_c2d286e997a4a112.__obf_850a7457bb739a32.__obf_892632c77e859037 = __obf_6e9f1b1176b96c20.__obf_6f0de5d9837befa9(__obf_c2d286e997a4a112.__obf_850a7457bb739a32.

		// SetEscapeHTML escape html by default, set to false to disable
		__obf_892632c77e859037.__obf_b4dfae156c7ac26c)
}

func (__obf_c2d286e997a4a112 *Encoder) SetEscapeHTML(__obf_6ee6f0f6c0425676 bool) {
	__obf_6e9f1b1176b96c20 := __obf_c2d286e997a4a112.__obf_850a7457bb739a32.__obf_892632c77e859037.__obf_e05f98c9e7cc9089
	__obf_6e9f1b1176b96c20.
		EscapeHTML = __obf_6ee6f0f6c0425676
	__obf_c2d286e997a4a112.__obf_850a7457bb739a32.__obf_892632c77e859037 = __obf_6e9f1b1176b96c20.__obf_6f0de5d9837befa9(__obf_c2d286e997a4a112.__obf_850a7457bb739a32.

		// Valid reports whether data is a valid JSON encoding.
		__obf_892632c77e859037.__obf_b4dfae156c7ac26c)
}

func Valid(__obf_2894589095c323b4 []byte) bool {
	return ConfigDefault.Valid(__obf_2894589095c323b4)
}
