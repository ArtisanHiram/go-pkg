package __obf_a4aad98aaf3178f4

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

func (__obf_66cc4b26e3c9cdbb RawMessage) EncodeMsgpack(__obf_dfba1ffd920163bc *Encoder) error {
	return __obf_dfba1ffd920163bc.__obf_ab9f7a04111e0451(__obf_66cc4b26e3c9cdbb)
}

func (__obf_66cc4b26e3c9cdbb *RawMessage) DecodeMsgpack(__obf_cc68ee77d25ca8f2 *Decoder) error {
	__obf_2130e988c0fe510a, __obf_4061ca0039150c39 := __obf_cc68ee77d25ca8f2.DecodeRaw()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	*__obf_66cc4b26e3c9cdbb = __obf_2130e988c0fe510a
	return nil
}

//------------------------------------------------------------------------------

type __obf_31d9eceb380e57ed struct {
	__obf_1ec1fd8709fc0afd string
	__obf_987b95dd4dcfc308 byte
}

func (__obf_4061ca0039150c39 __obf_31d9eceb380e57ed,) Error() string {
	return fmt.Sprintf("msgpack: unexpected code=%x decoding %s", __obf_4061ca0039150c39.__obf_987b95dd4dcfc308, __obf_4061ca0039150c39.__obf_1ec1fd8709fc0afd)
}
