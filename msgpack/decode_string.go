package __obf_3e0c303610a19bd4

import (
	"fmt"
	"reflect"
)

func (__obf_dc35117108ba8439 *Decoder) __obf_81bdf94818b2687d(__obf_e46289218af709bf byte) (int, error) {
	if __obf_e46289218af709bf == Nil {
		return -1, nil
	}

	if IsFixedString(__obf_e46289218af709bf) {
		return int(__obf_e46289218af709bf & FixedStrMask), nil
	}

	switch __obf_e46289218af709bf {
	case Str8, Bin8:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint8()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Str16, Bin16:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint16()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	case Str32, Bin32:
		__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.uint32()
		return int(__obf_4909ae60ffbb8e53), __obf_8882f6cf6e378ded
	}

	return 0, fmt.Errorf("msgpack: invalid code=%x decoding string/bytes length", __obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) DecodeString() (string, error) {
	if __obf_14e539ac760532b2 := __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_a063db36bdbf4202 != 0; __obf_14e539ac760532b2 || len(__obf_dc35117108ba8439.__obf_a22a31b815544cae) > 0 {
		return __obf_dc35117108ba8439.__obf_7f3d352cfbc0638e(__obf_14e539ac760532b2)
	}
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.string(__obf_e46289218af709bf)
}

func (__obf_dc35117108ba8439 *Decoder) string(__obf_e46289218af709bf byte) (string, error) {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_9198a6182aacf2fb(__obf_4909ae60ffbb8e53)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_9198a6182aacf2fb(__obf_4909ae60ffbb8e53 int) (string, error) {
	if __obf_4909ae60ffbb8e53 <= 0 {
		return "", nil
	}
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(__obf_4909ae60ffbb8e53)
	return string(__obf_11bcc66cde095c11), __obf_8882f6cf6e378ded
}

func __obf_442ca6b4ba0753ac(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_61027e0491b6dd3d, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeString()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetString(__obf_61027e0491b6dd3d)
	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_5484baeee52d4c8a() (string, error) {
	if __obf_14e539ac760532b2 := __obf_dc35117108ba8439.__obf_3cf0882fa5a4cafb&__obf_a063db36bdbf4202 != 0; __obf_14e539ac760532b2 || len(__obf_dc35117108ba8439.__obf_a22a31b815544cae) > 0 {
		return __obf_dc35117108ba8439.__obf_7f3d352cfbc0638e(__obf_14e539ac760532b2)
	}
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return "", nil
	}
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_b06a4f273442ca29(__obf_4909ae60ffbb8e53)
	if __obf_8882f6cf6e378ded != nil {
		return "", __obf_8882f6cf6e378ded
	}

	return __obf_f682cfd58bdd5de3(__obf_11bcc66cde095c11), nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_85900a462858a70d(__obf_b5a4664807537c0d *[]byte) error {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return __obf_dc35117108ba8439.__obf_75bee3b95d953a89(__obf_e46289218af709bf, __obf_b5a4664807537c0d)
}

func (__obf_dc35117108ba8439 *Decoder) __obf_75bee3b95d953a89(__obf_e46289218af709bf byte, __obf_b5a4664807537c0d *[]byte) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		*__obf_b5a4664807537c0d = nil
		return nil
	}

	*__obf_b5a4664807537c0d, __obf_8882f6cf6e378ded = __obf_b06a4f273442ca29(__obf_dc35117108ba8439.__obf_b3e2bfc96bdb9204, *__obf_b5a4664807537c0d, __obf_4909ae60ffbb8e53)
	return __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) __obf_a8d3bdf080f01438(__obf_e46289218af709bf byte) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 <= 0 {
		return nil
	}
	return __obf_dc35117108ba8439.__obf_19f294acbba68c47(__obf_4909ae60ffbb8e53)
}

func __obf_2866c43e524b97a1(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_e46289218af709bf, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5b7d85f9093902c3()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_81bdf94818b2687d(__obf_e46289218af709bf)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return nil
	}
	if __obf_4909ae60ffbb8e53 > __obf_63bbcee86d44fdde.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_63bbcee86d44fdde.Type(), __obf_63bbcee86d44fdde.Len(), __obf_4909ae60ffbb8e53)
	}
	__obf_11bcc66cde095c11 := __obf_63bbcee86d44fdde.Slice(0, __obf_4909ae60ffbb8e53).Bytes()
	return __obf_dc35117108ba8439.__obf_ac68fb7f5c42aa1d(__obf_11bcc66cde095c11)
}
