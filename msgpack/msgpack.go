package __obf_fd770cb9675903b0

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

func (__obf_777489ec8e6b2044 RawMessage) EncodeMsgpack(__obf_a7b5a0ca650f48ee *Encoder) error {
	return __obf_a7b5a0ca650f48ee.__obf_af6d14a7babbd464(__obf_777489ec8e6b2044)
}

func (__obf_777489ec8e6b2044 *RawMessage) DecodeMsgpack(__obf_2f8f01810a02d049 *Decoder) error {
	__obf_21691f6d5038a3f0, __obf_45342a3a754d12f5 := __obf_2f8f01810a02d049.DecodeRaw()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	*__obf_777489ec8e6b2044 = __obf_21691f6d5038a3f0
	return nil
}

//------------------------------------------------------------------------------

type __obf_b1f4502463d63ed9 struct {
	__obf_d6c122c727b9e88f string
	__obf_cde82889ba8e4822 byte
}

func (__obf_45342a3a754d12f5 __obf_b1f4502463d63ed9,) Error() string {
	return fmt.Sprintf("msgpack: unexpected code=%x decoding %s", __obf_45342a3a754d12f5.__obf_cde82889ba8e4822, __obf_45342a3a754d12f5.__obf_d6c122c727b9e88f)
}
