package __obf_a935eb7f548271a4

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

func (__obf_26d12ef0df36a324 RawMessage) EncodeMsgpack(__obf_6ca0309463018ed3 *Encoder) error {
	return __obf_6ca0309463018ed3.__obf_654725fc7d18ae48(__obf_26d12ef0df36a324)
}

func (__obf_26d12ef0df36a324 *RawMessage) DecodeMsgpack(__obf_fbe570faa7707c02 *Decoder) error {
	__obf_8d91290aff37bb4e, __obf_4d327e1cd40c2e21 := __obf_fbe570faa7707c02.DecodeRaw()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	*__obf_26d12ef0df36a324 = __obf_8d91290aff37bb4e
	return nil
}

//------------------------------------------------------------------------------

type __obf_769655cd52f45f6d struct {
	__obf_feeba0b68911eadc string
	__obf_b983039ef2a336eb byte
}

func (__obf_4d327e1cd40c2e21 __obf_769655cd52f45f6d,) Error() string {
	return fmt.Sprintf("msgpack: unexpected code=%x decoding %s", __obf_4d327e1cd40c2e21.__obf_b983039ef2a336eb, __obf_4d327e1cd40c2e21.__obf_feeba0b68911eadc)
}
