package __obf_3edaa49e853afa16

import "fmt"

type Marshaler interface {
	MarshalMsgpack() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalMsgpack([]byte) error
}

type CustomEncoder interface {
	EncodeMsgpack(*Encoder) error
}

type CustomDecoder interface {
	DecodeMsgpack(*Decoder) error
}

//------------------------------------------------------------------------------

type RawMessage []byte

var (
	_ CustomEncoder = (RawMessage)(nil)
	_ CustomDecoder = (*RawMessage)(nil)
)

func (__obf_c26f5adfb3152475 RawMessage) EncodeMsgpack(__obf_2fa4ef37f70ce1de *Encoder) error {
	return __obf_2fa4ef37f70ce1de.__obf_fe140b946d6a369e(__obf_c26f5adfb3152475)
}

func (__obf_c26f5adfb3152475 *RawMessage) DecodeMsgpack(__obf_20b87cf22b03338d *Decoder) error {
	__obf_771cde5bcb38d993, __obf_3aa78ad28f50ed46 := __obf_20b87cf22b03338d.DecodeRaw()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	*__obf_c26f5adfb3152475 = __obf_771cde5bcb38d993
	return nil
}

//------------------------------------------------------------------------------

type __obf_6b688bac335393d7 struct {
	__obf_6e0dac6e0c9d5f66 string
	__obf_508e2d6ff53655c0 byte
}

func (__obf_3aa78ad28f50ed46 __obf_6b688bac335393d7,) Error() string {
	return fmt.Sprintf("msgpack: unexpected code=%x decoding %s", __obf_3aa78ad28f50ed46.__obf_508e2d6ff53655c0, __obf_3aa78ad28f50ed46.__obf_6e0dac6e0c9d5f66)
}
