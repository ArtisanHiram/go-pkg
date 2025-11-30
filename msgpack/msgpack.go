package __obf_3e0c303610a19bd4

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

func (__obf_b416b6a4555be5c8 RawMessage) EncodeMsgpack(__obf_9721a7bc6b8fa008 *Encoder) error {
	return __obf_9721a7bc6b8fa008.__obf_4432f415b3007804(__obf_b416b6a4555be5c8)
}

func (__obf_b416b6a4555be5c8 *RawMessage) DecodeMsgpack(__obf_0ed630deb24db696 *Decoder) error {
	__obf_ea75b6610bb8e3c5, __obf_8882f6cf6e378ded := __obf_0ed630deb24db696.DecodeRaw()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	*__obf_b416b6a4555be5c8 = __obf_ea75b6610bb8e3c5
	return nil
}

//------------------------------------------------------------------------------

type __obf_dc8806ddbc8f0bfa struct {
	__obf_41a3520de571ca37 string
	__obf_545372fefbb733e5 byte
}

func (__obf_8882f6cf6e378ded __obf_dc8806ddbc8f0bfa,) Error() string {
	return fmt.Sprintf("msgpack: unexpected code=%x decoding %s", __obf_8882f6cf6e378ded.__obf_545372fefbb733e5, __obf_8882f6cf6e378ded.__obf_41a3520de571ca37)
}
