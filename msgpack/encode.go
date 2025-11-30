package __obf_3e0c303610a19bd4

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_8dd54bc25ec67c4c uint32 = 1 << iota
	__obf_7680454fe6a3ccd7
	__obf_506e82775380d9f7
	__obf_7ad13a8ba1bf6d52
	__obf_a063db36bdbf4202
	__obf_4536375463a63368
)

type __obf_1b6707f665313db3 interface {
	io.Writer
	WriteByte(byte) error
}

type __obf_f2c0d5e5b1149641 struct {
	io.Writer
}

func __obf_b29f1eb51c0e90b1(__obf_279ea11cfd0ea98d io.Writer) __obf_f2c0d5e5b1149641 {
	return __obf_f2c0d5e5b1149641{
		Writer: __obf_279ea11cfd0ea98d,
	}
}

func (__obf_95096a8fd2042e78 __obf_f2c0d5e5b1149641) WriteByte(__obf_e46289218af709bf byte) error {
	_, __obf_8882f6cf6e378ded := __obf_95096a8fd2042e78.Write([]byte{__obf_e46289218af709bf})
	return __obf_8882f6cf6e378ded
}

//------------------------------------------------------------------------------

var __obf_fba837edc792aacf = sync.Pool{
	New: func() any {
		return NewEncoder(nil)
	},
}

func GetEncoder() *Encoder {
	return __obf_fba837edc792aacf.Get().(*Encoder)
}

func PutEncoder(__obf_9721a7bc6b8fa008 *Encoder) {
	__obf_9721a7bc6b8fa008.__obf_279ea11cfd0ea98d = nil
	__obf_fba837edc792aacf.
		Put(__obf_9721a7bc6b8fa008)
}

// Marshal returns the MessagePack encoding of v.
func Marshal(__obf_63bbcee86d44fdde any) ([]byte, error) {
	__obf_9721a7bc6b8fa008 := GetEncoder()

	var __obf_27e4a1a33a31e4a7 bytes.Buffer
	__obf_9721a7bc6b8fa008.
		Reset(&__obf_27e4a1a33a31e4a7)
	__obf_8882f6cf6e378ded := __obf_9721a7bc6b8fa008.Encode(__obf_63bbcee86d44fdde)
	__obf_11bcc66cde095c11 := __obf_27e4a1a33a31e4a7.Bytes()

	PutEncoder(__obf_9721a7bc6b8fa008)

	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	return __obf_11bcc66cde095c11, __obf_8882f6cf6e378ded
}

type Encoder struct {
	__obf_279ea11cfd0ea98d __obf_1b6707f665313db3
	__obf_a22a31b815544cae map[string]int
	__obf_6d0a0d8a79287f26 string
	__obf_27e4a1a33a31e4a7 []byte
	__obf_42dab97af35e666e []byte
	__obf_3cf0882fa5a4cafb uint32
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(__obf_279ea11cfd0ea98d io.Writer) *Encoder {
	__obf_77240dc7776b60b4 := &Encoder{__obf_27e4a1a33a31e4a7: make([]byte, 9)}
	__obf_77240dc7776b60b4.
		Reset(__obf_279ea11cfd0ea98d)
	return __obf_77240dc7776b60b4
}

// Writer returns the Encoder's writer.
func (__obf_77240dc7776b60b4 *Encoder) Writer() io.Writer {
	return __obf_77240dc7776b60b4.

		// Reset discards any buffered data, resets all state, and switches the writer to write to w.
		__obf_279ea11cfd0ea98d
}

func (__obf_77240dc7776b60b4 *Encoder) Reset(__obf_279ea11cfd0ea98d io.Writer) {
	__obf_77240dc7776b60b4.
		ResetDict(__obf_279ea11cfd0ea98d, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_77240dc7776b60b4 *Encoder) ResetDict(__obf_279ea11cfd0ea98d io.Writer, __obf_a22a31b815544cae map[string]int) {
	__obf_77240dc7776b60b4.
		ResetWriter(__obf_279ea11cfd0ea98d)
	__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb = 0
	__obf_77240dc7776b60b4.__obf_6d0a0d8a79287f26 = ""
	__obf_77240dc7776b60b4.__obf_a22a31b815544cae = __obf_a22a31b815544cae
}

func (__obf_77240dc7776b60b4 *Encoder) WithDict(__obf_a22a31b815544cae map[string]int, __obf_6ff63d98a176dcfe func(*Encoder) error) error {
	__obf_d07f84e9e0444e28 := __obf_77240dc7776b60b4.__obf_a22a31b815544cae
	__obf_77240dc7776b60b4.__obf_a22a31b815544cae = __obf_a22a31b815544cae
	__obf_8882f6cf6e378ded := __obf_6ff63d98a176dcfe(__obf_77240dc7776b60b4)
	__obf_77240dc7776b60b4.__obf_a22a31b815544cae = __obf_d07f84e9e0444e28
	return __obf_8882f6cf6e378ded
}

func (__obf_77240dc7776b60b4 *Encoder) ResetWriter(__obf_279ea11cfd0ea98d io.Writer) {
	__obf_77240dc7776b60b4.__obf_a22a31b815544cae = nil
	if __obf_95096a8fd2042e78, __obf_ce8bef141eff8aab := __obf_279ea11cfd0ea98d.(__obf_1b6707f665313db3); __obf_ce8bef141eff8aab {
		__obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d = __obf_95096a8fd2042e78
	} else if __obf_279ea11cfd0ea98d == nil {
		__obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d = nil
	} else {
		__obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d = __obf_b29f1eb51c0e90b1(__obf_279ea11cfd0ea98d)
	}
}

// SetSortMapKeys causes the Encoder to encode map keys in increasing order.
// Supported map types are:
//   - map[string]string
//   - map[string]bool
//   - map[string]any
func (__obf_77240dc7776b60b4 *Encoder) SetSortMapKeys(__obf_2142d863a5bde925 bool) *Encoder {
	if __obf_2142d863a5bde925 {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb |= __obf_8dd54bc25ec67c4c
	} else {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb &= ^__obf_8dd54bc25ec67c4c
	}
	return __obf_77240dc7776b60b4
}

// SetCustomStructTag causes the Encoder to use a custom struct tag as
// fallback option if there is no msgpack tag.
func (__obf_77240dc7776b60b4 *Encoder) SetCustomStructTag(__obf_42ff6f144202733b string) {
	__obf_77240dc7776b60b4.__obf_6d0a0d8a79287f26 = __obf_42ff6f144202733b
}

// SetOmitEmpty causes the Encoder to omit empty values by default.
func (__obf_77240dc7776b60b4 *Encoder) SetOmitEmpty(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb |= __obf_4536375463a63368
	} else {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb &= ^__obf_4536375463a63368
	}
}

// UseArrayEncodedStructs causes the Encoder to encode Go structs as msgpack arrays.
func (__obf_77240dc7776b60b4 *Encoder) UseArrayEncodedStructs(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb |= __obf_7680454fe6a3ccd7
	} else {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb &= ^__obf_7680454fe6a3ccd7
	}
}

// UseCompactEncoding causes the Encoder to chose the most compact encoding.
// For example, it allows to encode small Go int64 as msgpack int8 saving 7 bytes.
func (__obf_77240dc7776b60b4 *Encoder) UseCompactInts(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb |= __obf_506e82775380d9f7
	} else {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb &= ^__obf_506e82775380d9f7
	}
}

// UseCompactFloats causes the Encoder to chose a compact integer encoding
// for floats that can be represented as integers.
func (__obf_77240dc7776b60b4 *Encoder) UseCompactFloats(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb |= __obf_7ad13a8ba1bf6d52
	} else {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb &= ^__obf_7ad13a8ba1bf6d52
	}
}

// UseInternedStrings causes the Encoder to intern strings.
func (__obf_77240dc7776b60b4 *Encoder) UseInternedStrings(__obf_2142d863a5bde925 bool) {
	if __obf_2142d863a5bde925 {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb |= __obf_a063db36bdbf4202
	} else {
		__obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb &= ^__obf_a063db36bdbf4202
	}
}

func (__obf_77240dc7776b60b4 *Encoder) Encode(__obf_63bbcee86d44fdde any) error {
	switch __obf_63bbcee86d44fdde := __obf_63bbcee86d44fdde.(type) {
	case nil:
		return __obf_77240dc7776b60b4.EncodeNil()
	case string:
		return __obf_77240dc7776b60b4.EncodeString(__obf_63bbcee86d44fdde)
	case []byte:
		return __obf_77240dc7776b60b4.EncodeBytes(__obf_63bbcee86d44fdde)
	case int:
		return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_63bbcee86d44fdde))
	case int64:
		return __obf_77240dc7776b60b4.__obf_c919f124e6c2fb8b(__obf_63bbcee86d44fdde)
	case uint:
		return __obf_77240dc7776b60b4.EncodeUint(uint64(__obf_63bbcee86d44fdde))
	case uint64:
		return __obf_77240dc7776b60b4.__obf_e32b04a6bb97fb7e(__obf_63bbcee86d44fdde)
	case bool:
		return __obf_77240dc7776b60b4.EncodeBool(__obf_63bbcee86d44fdde)
	case float32:
		return __obf_77240dc7776b60b4.EncodeFloat32(__obf_63bbcee86d44fdde)
	case float64:
		return __obf_77240dc7776b60b4.EncodeFloat64(__obf_63bbcee86d44fdde)
	case time.Duration:
		return __obf_77240dc7776b60b4.__obf_c919f124e6c2fb8b(int64(__obf_63bbcee86d44fdde))
	case time.Time:
		return __obf_77240dc7776b60b4.EncodeTime(__obf_63bbcee86d44fdde)
	}
	return __obf_77240dc7776b60b4.EncodeValue(reflect.ValueOf(__obf_63bbcee86d44fdde))
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeMulti(__obf_63bbcee86d44fdde ...any) error {
	for _, __obf_785552c2ee57ae56 := range __obf_63bbcee86d44fdde {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.Encode(__obf_785552c2ee57ae56); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeValue(__obf_63bbcee86d44fdde reflect.Value) error {
	__obf_6ff63d98a176dcfe := __obf_cc16be3b4933c617(__obf_63bbcee86d44fdde.Type())
	return __obf_6ff63d98a176dcfe(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeNil() error {
	return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(Nil)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeBool(__obf_752fc3666f4041f7 bool) error {
	if __obf_752fc3666f4041f7 {
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(True)
	}
	return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(False)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeDuration(__obf_dc35117108ba8439 time.Duration) error {
	return __obf_77240dc7776b60b4.EncodeInt(int64(__obf_dc35117108ba8439))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_e8d9aa5d4f5eb974(__obf_e46289218af709bf byte) error {
	return __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.WriteByte(__obf_e46289218af709bf)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_4432f415b3007804(__obf_11bcc66cde095c11 []byte) error {
	_, __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.Write(__obf_11bcc66cde095c11)
	return __obf_8882f6cf6e378ded
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_21c7568bc3d83316(__obf_61027e0491b6dd3d string) error {
	_, __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.Write(__obf_31018e102d310148(__obf_61027e0491b6dd3d))
	return __obf_8882f6cf6e378ded
}
