package __obf_a4aad98aaf3178f4

import (
	"bytes"
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	__obf_1f7890468d217f4e uint32 = 1 << iota
	__obf_3e711425f784a6a6
	__obf_e0996ed6dc481d61
	__obf_d003dec3186c9089
	__obf_ca379a8882c3af10
	__obf_33a8b849a29cb9a5
)

type __obf_a236ee016ac0ca58 interface {
	io.Writer
	WriteByte(byte) error
}

type __obf_0e1b27a9b594e937 struct {
	io.Writer
}

func __obf_f6e799aa08fc5d38(__obf_bc988362b9d818fa io.Writer) __obf_0e1b27a9b594e937 {
	return __obf_0e1b27a9b594e937{
		Writer: __obf_bc988362b9d818fa,
	}
}

func (__obf_018fe7e725574dca __obf_0e1b27a9b594e937) WriteByte(__obf_6dbe86b3d9d9b037 byte) error {
	_, __obf_4061ca0039150c39 := __obf_018fe7e725574dca.Write([]byte{__obf_6dbe86b3d9d9b037})
	return __obf_4061ca0039150c39
}

//------------------------------------------------------------------------------

var __obf_bcf8868e84f51678 = sync.Pool{
	New: func() any {
		return NewEncoder(nil)
	},
}

func GetEncoder() *Encoder {
	return __obf_bcf8868e84f51678.Get().(*Encoder)
}

func PutEncoder(__obf_dfba1ffd920163bc *Encoder) {
	__obf_dfba1ffd920163bc.__obf_bc988362b9d818fa = nil
	__obf_bcf8868e84f51678.
		Put(__obf_dfba1ffd920163bc)
}

// Marshal returns the MessagePack encoding of v.
func Marshal(__obf_a1f43267eeb48a1a any) ([]byte, error) {
	__obf_dfba1ffd920163bc := GetEncoder()

	var __obf_f5d391a424aa5a3a bytes.Buffer
	__obf_dfba1ffd920163bc.
		Reset(&__obf_f5d391a424aa5a3a)
	__obf_4061ca0039150c39 := __obf_dfba1ffd920163bc.Encode(__obf_a1f43267eeb48a1a)
	__obf_f57209cfda8e17cf := __obf_f5d391a424aa5a3a.Bytes()

	PutEncoder(__obf_dfba1ffd920163bc)

	if __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	return __obf_f57209cfda8e17cf, __obf_4061ca0039150c39
}

type Encoder struct {
	__obf_bc988362b9d818fa __obf_a236ee016ac0ca58
	__obf_57edbb6a6615f57c map[string]int
	__obf_72ba705ed504d210 string
	__obf_f5d391a424aa5a3a []byte
	__obf_ccee18a737f071fd []byte
	__obf_a8400e70a712e9fa uint32
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(__obf_bc988362b9d818fa io.Writer) *Encoder {
	__obf_2c8e97779108ab17 := &Encoder{__obf_f5d391a424aa5a3a: make([]byte, 9)}
	__obf_2c8e97779108ab17.
		Reset(__obf_bc988362b9d818fa)
	return __obf_2c8e97779108ab17
}

// Writer returns the Encoder's writer.
func (__obf_2c8e97779108ab17 *Encoder) Writer() io.Writer {
	return __obf_2c8e97779108ab17.

		// Reset discards any buffered data, resets all state, and switches the writer to write to w.
		__obf_bc988362b9d818fa
}

func (__obf_2c8e97779108ab17 *Encoder) Reset(__obf_bc988362b9d818fa io.Writer) {
	__obf_2c8e97779108ab17.
		ResetDict(__obf_bc988362b9d818fa, nil)
}

// ResetDict is like Reset, but also resets the dict.
func (__obf_2c8e97779108ab17 *Encoder) ResetDict(__obf_bc988362b9d818fa io.Writer, __obf_57edbb6a6615f57c map[string]int) {
	__obf_2c8e97779108ab17.
		ResetWriter(__obf_bc988362b9d818fa)
	__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa = 0
	__obf_2c8e97779108ab17.__obf_72ba705ed504d210 = ""
	__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c = __obf_57edbb6a6615f57c
}

func (__obf_2c8e97779108ab17 *Encoder) WithDict(__obf_57edbb6a6615f57c map[string]int, __obf_64fe495dbc0c543b func(*Encoder) error) error {
	__obf_187a5bedd2cd4c82 := __obf_2c8e97779108ab17.__obf_57edbb6a6615f57c
	__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c = __obf_57edbb6a6615f57c
	__obf_4061ca0039150c39 := __obf_64fe495dbc0c543b(__obf_2c8e97779108ab17)
	__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c = __obf_187a5bedd2cd4c82
	return __obf_4061ca0039150c39
}

func (__obf_2c8e97779108ab17 *Encoder) ResetWriter(__obf_bc988362b9d818fa io.Writer) {
	__obf_2c8e97779108ab17.__obf_57edbb6a6615f57c = nil
	if __obf_018fe7e725574dca, __obf_81b19f2a6159ab89 := __obf_bc988362b9d818fa.(__obf_a236ee016ac0ca58); __obf_81b19f2a6159ab89 {
		__obf_2c8e97779108ab17.__obf_bc988362b9d818fa = __obf_018fe7e725574dca
	} else if __obf_bc988362b9d818fa == nil {
		__obf_2c8e97779108ab17.__obf_bc988362b9d818fa = nil
	} else {
		__obf_2c8e97779108ab17.__obf_bc988362b9d818fa = __obf_f6e799aa08fc5d38(__obf_bc988362b9d818fa)
	}
}

// SetSortMapKeys causes the Encoder to encode map keys in increasing order.
// Supported map types are:
//   - map[string]string
//   - map[string]bool
//   - map[string]any
func (__obf_2c8e97779108ab17 *Encoder) SetSortMapKeys(__obf_2f731eb262a74383 bool) *Encoder {
	if __obf_2f731eb262a74383 {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa |= __obf_1f7890468d217f4e
	} else {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa &= ^__obf_1f7890468d217f4e
	}
	return __obf_2c8e97779108ab17
}

// SetCustomStructTag causes the Encoder to use a custom struct tag as
// fallback option if there is no msgpack tag.
func (__obf_2c8e97779108ab17 *Encoder) SetCustomStructTag(__obf_990470bfaf220103 string) {
	__obf_2c8e97779108ab17.__obf_72ba705ed504d210 = __obf_990470bfaf220103
}

// SetOmitEmpty causes the Encoder to omit empty values by default.
func (__obf_2c8e97779108ab17 *Encoder) SetOmitEmpty(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa |= __obf_33a8b849a29cb9a5
	} else {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa &= ^__obf_33a8b849a29cb9a5
	}
}

// UseArrayEncodedStructs causes the Encoder to encode Go structs as msgpack arrays.
func (__obf_2c8e97779108ab17 *Encoder) UseArrayEncodedStructs(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa |= __obf_3e711425f784a6a6
	} else {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa &= ^__obf_3e711425f784a6a6
	}
}

// UseCompactEncoding causes the Encoder to chose the most compact encoding.
// For example, it allows to encode small Go int64 as msgpack int8 saving 7 bytes.
func (__obf_2c8e97779108ab17 *Encoder) UseCompactInts(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa |= __obf_e0996ed6dc481d61
	} else {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa &= ^__obf_e0996ed6dc481d61
	}
}

// UseCompactFloats causes the Encoder to chose a compact integer encoding
// for floats that can be represented as integers.
func (__obf_2c8e97779108ab17 *Encoder) UseCompactFloats(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa |= __obf_d003dec3186c9089
	} else {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa &= ^__obf_d003dec3186c9089
	}
}

// UseInternedStrings causes the Encoder to intern strings.
func (__obf_2c8e97779108ab17 *Encoder) UseInternedStrings(__obf_2f731eb262a74383 bool) {
	if __obf_2f731eb262a74383 {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa |= __obf_ca379a8882c3af10
	} else {
		__obf_2c8e97779108ab17.__obf_a8400e70a712e9fa &= ^__obf_ca379a8882c3af10
	}
}

func (__obf_2c8e97779108ab17 *Encoder) Encode(__obf_a1f43267eeb48a1a any) error {
	switch __obf_a1f43267eeb48a1a := __obf_a1f43267eeb48a1a.(type) {
	case nil:
		return __obf_2c8e97779108ab17.EncodeNil()
	case string:
		return __obf_2c8e97779108ab17.EncodeString(__obf_a1f43267eeb48a1a)
	case []byte:
		return __obf_2c8e97779108ab17.EncodeBytes(__obf_a1f43267eeb48a1a)
	case int:
		return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_a1f43267eeb48a1a))
	case int64:
		return __obf_2c8e97779108ab17.__obf_2837531178551b8d(__obf_a1f43267eeb48a1a)
	case uint:
		return __obf_2c8e97779108ab17.EncodeUint(uint64(__obf_a1f43267eeb48a1a))
	case uint64:
		return __obf_2c8e97779108ab17.__obf_a5a421aeac90eb76(__obf_a1f43267eeb48a1a)
	case bool:
		return __obf_2c8e97779108ab17.EncodeBool(__obf_a1f43267eeb48a1a)
	case float32:
		return __obf_2c8e97779108ab17.EncodeFloat32(__obf_a1f43267eeb48a1a)
	case float64:
		return __obf_2c8e97779108ab17.EncodeFloat64(__obf_a1f43267eeb48a1a)
	case time.Duration:
		return __obf_2c8e97779108ab17.__obf_2837531178551b8d(int64(__obf_a1f43267eeb48a1a))
	case time.Time:
		return __obf_2c8e97779108ab17.EncodeTime(__obf_a1f43267eeb48a1a)
	}
	return __obf_2c8e97779108ab17.EncodeValue(reflect.ValueOf(__obf_a1f43267eeb48a1a))
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeMulti(__obf_a1f43267eeb48a1a ...any) error {
	for _, __obf_9f1e21befe556e87 := range __obf_a1f43267eeb48a1a {
		if __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.Encode(__obf_9f1e21befe556e87); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeValue(__obf_a1f43267eeb48a1a reflect.Value) error {
	__obf_64fe495dbc0c543b := __obf_f1f0aa6d8078637c(__obf_a1f43267eeb48a1a.Type())
	return __obf_64fe495dbc0c543b(__obf_2c8e97779108ab17, __obf_a1f43267eeb48a1a)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeNil() error {
	return __obf_2c8e97779108ab17.__obf_c087409347698442(Nil)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeBool(__obf_9055e1481edcd436 bool) error {
	if __obf_9055e1481edcd436 {
		return __obf_2c8e97779108ab17.__obf_c087409347698442(True)
	}
	return __obf_2c8e97779108ab17.__obf_c087409347698442(False)
}

func (__obf_2c8e97779108ab17 *Encoder) EncodeDuration(__obf_613397fefdec0ed0 time.Duration) error {
	return __obf_2c8e97779108ab17.EncodeInt(int64(__obf_613397fefdec0ed0))
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_c087409347698442(__obf_6dbe86b3d9d9b037 byte) error {
	return __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.WriteByte(__obf_6dbe86b3d9d9b037)
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_ab9f7a04111e0451(__obf_f57209cfda8e17cf []byte) error {
	_, __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.Write(__obf_f57209cfda8e17cf)
	return __obf_4061ca0039150c39
}

func (__obf_2c8e97779108ab17 *Encoder) __obf_176357c847e61809(__obf_759f458bfa19abba string) error {
	_, __obf_4061ca0039150c39 := __obf_2c8e97779108ab17.__obf_bc988362b9d818fa.Write(__obf_40df13eb8b76aa84(__obf_759f458bfa19abba))
	return __obf_4061ca0039150c39
}
