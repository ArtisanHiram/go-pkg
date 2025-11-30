package __obf_de86cdc8ae98b45b

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

func (__obf_2a26e7104b4c4373 RawMessage) EncodeMsgpack(__obf_6a6518f236eeec6e *Encoder) error {
	return __obf_6a6518f236eeec6e.__obf_b0daaf664cd5cdfb(__obf_2a26e7104b4c4373)
}

func (__obf_2a26e7104b4c4373 *RawMessage) DecodeMsgpack(__obf_9461956859459a44 *Decoder) error {
	__obf_72c36aa69a5d48f8, __obf_0feb3528b7b709ec := __obf_9461956859459a44.DecodeRaw()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	*__obf_2a26e7104b4c4373 = __obf_72c36aa69a5d48f8
	return nil
}

//------------------------------------------------------------------------------

type __obf_150b3812dfd829ef struct {
	__obf_ee374540697f97c0 string
	__obf_34e3ba264a6bb541 byte
}

func (__obf_0feb3528b7b709ec __obf_150b3812dfd829ef,) Error() string {
	return fmt.Sprintf("msgpack: unexpected code=%x decoding %s", __obf_0feb3528b7b709ec.__obf_34e3ba264a6bb541, __obf_0feb3528b7b709ec.__obf_ee374540697f97c0)
}
