package __obf_fd770cb9675903b0

import (
	"fmt"
	"reflect"
)

func (__obf_5d389b660eefb08c *Decoder) __obf_32880f8a880f1644(__obf_4148125b350dfea2 byte) (int, error) {
	if __obf_4148125b350dfea2 == Nil {
		return -1, nil
	}

	if IsFixedString(__obf_4148125b350dfea2) {
		return int(__obf_4148125b350dfea2 & FixedStrMask), nil
	}

	switch __obf_4148125b350dfea2 {
	case Str8, Bin8:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint8()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Str16, Bin16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Str32, Bin32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	}

	return 0, fmt.Errorf("msgpack: invalid code=%x decoding string/bytes length", __obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) DecodeString() (string, error) {
	if __obf_4968ec805e7c5822 := __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_139c02350e62bdf7 != 0; __obf_4968ec805e7c5822 || len(__obf_5d389b660eefb08c.__obf_ff96db22e12b6842) > 0 {
		return __obf_5d389b660eefb08c.__obf_06f8691b61014235(__obf_4968ec805e7c5822)
	}
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.string(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) string(__obf_4148125b350dfea2 byte) (string, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_6b53a4200bd06df8(__obf_d774e4844c74bfe9)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_6b53a4200bd06df8(__obf_d774e4844c74bfe9 int) (string, error) {
	if __obf_d774e4844c74bfe9 <= 0 {
		return "", nil
	}
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(__obf_d774e4844c74bfe9)
	return string(__obf_94333af0f0facd65), __obf_45342a3a754d12f5
}

func __obf_5ac5b748094f0aee(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeString()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_f328a048f76b7256.
		SetString(__obf_fe1ace7a2eb8bf9f)
	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_479ae9aab7b05d68() (string, error) {
	if __obf_4968ec805e7c5822 := __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_139c02350e62bdf7 != 0; __obf_4968ec805e7c5822 || len(__obf_5d389b660eefb08c.__obf_ff96db22e12b6842) > 0 {
		return __obf_5d389b660eefb08c.__obf_06f8691b61014235(__obf_4968ec805e7c5822)
	}
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return "", nil
	}
	__obf_94333af0f0facd65, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_31c39140d4d66749(__obf_d774e4844c74bfe9)
	if __obf_45342a3a754d12f5 != nil {
		return "", __obf_45342a3a754d12f5
	}

	return __obf_8c86f40d3a80c733(__obf_94333af0f0facd65), nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_f4b9f1a4ece9fd91(__obf_2c49f9d2007cfaf6 *[]byte) error {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_6545e5e821451815(__obf_4148125b350dfea2, __obf_2c49f9d2007cfaf6)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_6545e5e821451815(__obf_4148125b350dfea2 byte, __obf_2c49f9d2007cfaf6 *[]byte) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		*__obf_2c49f9d2007cfaf6 = nil
		return nil
	}

	*__obf_2c49f9d2007cfaf6, __obf_45342a3a754d12f5 = __obf_31c39140d4d66749(__obf_5d389b660eefb08c.__obf_52c07f7b574f636a, *__obf_2c49f9d2007cfaf6, __obf_d774e4844c74bfe9)
	return __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) __obf_870f749ef029eefb(__obf_4148125b350dfea2 byte) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 <= 0 {
		return nil
	}
	return __obf_5d389b660eefb08c.__obf_d93e03df64d057a0(__obf_d774e4844c74bfe9)
}

func __obf_bcddc3fae95e8e05(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_32880f8a880f1644(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return nil
	}
	if __obf_d774e4844c74bfe9 > __obf_f328a048f76b7256.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_f328a048f76b7256.Type(), __obf_f328a048f76b7256.Len(), __obf_d774e4844c74bfe9)
	}
	__obf_94333af0f0facd65 := __obf_f328a048f76b7256.Slice(0, __obf_d774e4844c74bfe9).Bytes()
	return __obf_5d389b660eefb08c.__obf_e6a3e7cb981bc1c2(__obf_94333af0f0facd65)
}
