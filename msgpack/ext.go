package __obf_3e0c303610a19bd4

import (
	"fmt"
	"math"
	"reflect"
)

type __obf_6be15d2128e93f39 struct {
	Type    reflect.Type
	Decoder func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value, __obf_7e99dd6aa706c57b int) error
}

var __obf_d7c9956a45f4419c = make(map[int8]*__obf_6be15d2128e93f39)

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

func RegisterExt(__obf_927f1edb79ba3be1 int8, __obf_752fc3666f4041f7 MarshalerUnmarshaler) {
	RegisterExtEncoder(__obf_927f1edb79ba3be1, __obf_752fc3666f4041f7, func(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) ([]byte, error) {
		__obf_9a68ec3a8d69fdbf := __obf_63bbcee86d44fdde.Interface().(Marshaler)
		return __obf_9a68ec3a8d69fdbf.MarshalMsgpack()
	})
	RegisterExtDecoder(__obf_927f1edb79ba3be1, __obf_752fc3666f4041f7, func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value, __obf_7e99dd6aa706c57b int) error {
		__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(__obf_7e99dd6aa706c57b)
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		return __obf_63bbcee86d44fdde.Interface().(Unmarshaler).UnmarshalMsgpack(__obf_11bcc66cde095c11)
	})
}

func UnregisterExt(__obf_927f1edb79ba3be1 int8) {
	__obf_3e7b6ef91fb6f5aa(__obf_927f1edb79ba3be1)
	__obf_48f9e4889360bcdb(__obf_927f1edb79ba3be1)
}

func RegisterExtEncoder(__obf_927f1edb79ba3be1 int8, __obf_752fc3666f4041f7 any, __obf_02e33ee3dfeba894 func(__obf_9721a7bc6b8fa008 *Encoder, __obf_63bbcee86d44fdde reflect.Value) ([]byte, error),
) {
	__obf_3e7b6ef91fb6f5aa(__obf_927f1edb79ba3be1)
	__obf_616f98efc80197c6 := reflect.TypeOf(__obf_752fc3666f4041f7)
	__obf_ee46be0dc82c314f := __obf_8e5df4266de74335(__obf_927f1edb79ba3be1, __obf_616f98efc80197c6, __obf_02e33ee3dfeba894)
	__obf_b932a15e2d3972de.
		Store(__obf_927f1edb79ba3be1, __obf_616f98efc80197c6)
	__obf_b932a15e2d3972de.
		Store(__obf_616f98efc80197c6, __obf_ee46be0dc82c314f)
	if __obf_616f98efc80197c6.Kind() == reflect.Ptr {
		__obf_b932a15e2d3972de.
			Store(__obf_616f98efc80197c6.Elem(), __obf_efbae37935976036(__obf_ee46be0dc82c314f))
	}
}

func __obf_3e7b6ef91fb6f5aa(__obf_927f1edb79ba3be1 int8) {
	__obf_cb646455b7fd3ba7, __obf_ce8bef141eff8aab := __obf_b932a15e2d3972de.Load(__obf_927f1edb79ba3be1)
	if !__obf_ce8bef141eff8aab {
		return
	}
	__obf_b932a15e2d3972de.
		Delete(__obf_927f1edb79ba3be1)
	__obf_616f98efc80197c6 := __obf_cb646455b7fd3ba7.(reflect.Type)
	__obf_b932a15e2d3972de.
		Delete(__obf_616f98efc80197c6)
	if __obf_616f98efc80197c6.Kind() == reflect.Ptr {
		__obf_b932a15e2d3972de.
			Delete(__obf_616f98efc80197c6.Elem())
	}
}

func __obf_8e5df4266de74335(__obf_927f1edb79ba3be1 int8, __obf_616f98efc80197c6 reflect.Type, __obf_02e33ee3dfeba894 func(__obf_9721a7bc6b8fa008 *Encoder, __obf_63bbcee86d44fdde reflect.Value) ([]byte, error),
) __obf_9cac677a20c1be1c {
	__obf_ebd8855a2f3055ad := __obf_616f98efc80197c6.Kind() == reflect.Ptr

	return func(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if __obf_ebd8855a2f3055ad && __obf_63bbcee86d44fdde.IsNil() {
			return __obf_77240dc7776b60b4.EncodeNil()
		}
		__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_02e33ee3dfeba894(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde)
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}

		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeExtHeader(__obf_927f1edb79ba3be1, len(__obf_11bcc66cde095c11)); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}

		return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_11bcc66cde095c11)
	}
}

func __obf_efbae37935976036(__obf_ee46be0dc82c314f __obf_9cac677a20c1be1c) __obf_9cac677a20c1be1c {
	return func(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if !__obf_63bbcee86d44fdde.CanAddr() {
			return fmt.Errorf("msgpack: EncodeExt(nonaddressable %T)", __obf_63bbcee86d44fdde.Interface())
		}
		return __obf_ee46be0dc82c314f(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde.Addr())
	}
}

func RegisterExtDecoder(__obf_927f1edb79ba3be1 int8, __obf_752fc3666f4041f7 any, __obf_91fd0d587e6af119 func(__obf_0ed630deb24db696 *Decoder, __obf_63bbcee86d44fdde reflect.Value, __obf_7e99dd6aa706c57b int) error,
) {
	__obf_48f9e4889360bcdb(__obf_927f1edb79ba3be1)
	__obf_616f98efc80197c6 := reflect.TypeOf(__obf_752fc3666f4041f7)
	__obf_65623278443ddfbc := __obf_8333f62a59b544eb(__obf_927f1edb79ba3be1, __obf_616f98efc80197c6, __obf_91fd0d587e6af119)
	__obf_d7c9956a45f4419c[__obf_927f1edb79ba3be1] = &__obf_6be15d2128e93f39{
		Type:    __obf_616f98efc80197c6,
		Decoder: __obf_91fd0d587e6af119,
	}
	__obf_a46bdc105884db8b.
		Store(__obf_927f1edb79ba3be1, __obf_616f98efc80197c6)
	__obf_a46bdc105884db8b.
		Store(__obf_616f98efc80197c6, __obf_65623278443ddfbc)
	if __obf_616f98efc80197c6.Kind() == reflect.Ptr {
		__obf_a46bdc105884db8b.
			Store(__obf_616f98efc80197c6.Elem(), __obf_8de8d939775d0c44(__obf_65623278443ddfbc))
	}
}

func __obf_48f9e4889360bcdb(__obf_927f1edb79ba3be1 int8) {
	__obf_cb646455b7fd3ba7, __obf_ce8bef141eff8aab := __obf_a46bdc105884db8b.Load(__obf_927f1edb79ba3be1)
	if !__obf_ce8bef141eff8aab {
		return
	}
	__obf_a46bdc105884db8b.
		Delete(__obf_927f1edb79ba3be1)
	delete(__obf_d7c9956a45f4419c, __obf_927f1edb79ba3be1)
	__obf_616f98efc80197c6 := __obf_cb646455b7fd3ba7.(reflect.Type)
	__obf_a46bdc105884db8b.
		Delete(__obf_616f98efc80197c6)
	if __obf_616f98efc80197c6.Kind() == reflect.Ptr {
		__obf_a46bdc105884db8b.
			Delete(__obf_616f98efc80197c6.Elem())
	}
}

func __obf_8333f62a59b544eb(__obf_5a36b7d07d181a8b int8, __obf_616f98efc80197c6 reflect.Type, __obf_91fd0d587e6af119 func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value, __obf_7e99dd6aa706c57b int) error,
) __obf_299b7590962e0960 {
	return __obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
		__obf_927f1edb79ba3be1, __obf_7e99dd6aa706c57b, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeExtHeader()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		if __obf_927f1edb79ba3be1 != __obf_5a36b7d07d181a8b {
			return fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_927f1edb79ba3be1, __obf_5a36b7d07d181a8b)
		}
		return __obf_91fd0d587e6af119(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde, __obf_7e99dd6aa706c57b)
	})
}

func __obf_8de8d939775d0c44(__obf_65623278443ddfbc __obf_299b7590962e0960) __obf_299b7590962e0960 {
	return func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if !__obf_63bbcee86d44fdde.CanAddr() {
			return fmt.Errorf("msgpack: DecodeExt(nonaddressable %T)", __obf_63bbcee86d44fdde.Interface())
		}
		return __obf_65623278443ddfbc(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde.Addr())
	}
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeExtHeader(__obf_927f1edb79ba3be1 int8, __obf_7e99dd6aa706c57b int) error {
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_7a9a2753ca4f49af(__obf_7e99dd6aa706c57b); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.WriteByte(byte(__obf_927f1edb79ba3be1)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return nil
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_7a9a2753ca4f49af(__obf_d62f82ed950927da int) error {
	switch __obf_d62f82ed950927da {
	case 1:
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt1)
	case 2:
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt2)
	case 4:
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt4)
	case 8:
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt8)
	case 16:
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixExt16)
	}
	if __obf_d62f82ed950927da <= math.MaxUint8 {
		return __obf_77240dc7776b60b4.__obf_a5ea8b1a318bf9eb(Ext8, uint8(__obf_d62f82ed950927da))
	}
	if __obf_d62f82ed950927da <= math.MaxUint16 {
		return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Ext16, uint16(__obf_d62f82ed950927da))
	}
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Ext32, uint32(__obf_d62f82ed950927da))
}

func (__obf_dc35117108ba8439 *Decoder) DecodeExtHeader() (__obf_927f1edb79ba3be1 int8, __obf_7e99dd6aa706c57b int, __obf_8882f6cf6e378ded error) {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return
	}
	return __obf_dc35117108ba8439.__obf_3ac0957b06c080b5(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_3ac0957b06c080b5(__obf_e46289218af709bf byte) (int8, int, error) {
	__obf_7e99dd6aa706c57b, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_cd81d070fbb203b9(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return 0, 0, __obf_8882f6cf6e378ded
	}
	__obf_927f1edb79ba3be1, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return 0, 0, __obf_8882f6cf6e378ded
	}

	return int8(__obf_927f1edb79ba3be1), __obf_7e99dd6aa706c57b, nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_cd81d070fbb203b9(__obf_e46289218af709bf byte) (int, error) {
	switch __obf_e46289218af709bf {
	case FixExt1:
		return 1, nil
	case FixExt2:
		return 2, nil
	case FixExt4:
		return 4, nil
	case FixExt8:
		return 8, nil
	case FixExt16:
		return 16, nil
	case Ext8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Ext16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Ext32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	default:
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", __obf_e46289218af709bf)
	}
}

func (__obf_dc35117108ba8439 *Decoder) __obf_5b9b53ab12d1dfc9(__obf_e46289218af709bf byte) (any, error) {
	__obf_927f1edb79ba3be1, __obf_7e99dd6aa706c57b, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_3ac0957b06c080b5(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	__obf_49b9e0bac7f04521, __obf_ce8bef141eff8aab := __obf_d7c9956a45f4419c[__obf_927f1edb79ba3be1]
	if !__obf_ce8bef141eff8aab {
		return nil, fmt.Errorf("msgpack: unknown ext id=%d", __obf_927f1edb79ba3be1)
	}
	__obf_63bbcee86d44fdde := __obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_49b9e0bac7f04521.Type).Elem()
	if __obf_ebd8855a2f3055ad(__obf_63bbcee86d44fdde.Kind()) && __obf_63bbcee86d44fdde.IsNil() {
		__obf_63bbcee86d44fdde.
			Set(__obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_49b9e0bac7f04521.Type.Elem()))
	}

	if __obf_8882f6cf6e378ded := __obf_49b9e0bac7f04521.Decoder(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde, __obf_7e99dd6aa706c57b); __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}

	return __obf_63bbcee86d44fdde.Interface(), nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_f260dda48c6603bb(__obf_e46289218af709bf byte) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_cd81d070fbb203b9(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_19f294acbba68c47(__obf_4909ae60ffbb8e53 + 1)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_2ae9d023bf150409(__obf_e46289218af709bf byte) error {
	// Read ext type.
	_, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	// Read ext body len.
	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_f97b03557973cf47(__obf_e46289218af709bf); __obf_39714879601f9b69++ {
		_, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}

func __obf_f97b03557973cf47(__obf_e46289218af709bf byte) int {
	switch __obf_e46289218af709bf {
	case Ext8:
		return 1
	case Ext16:
		return 2
	case Ext32:
		return 4
	}
	return 0
}
