package __obf_fd770cb9675903b0

import (
	"fmt"
	"reflect"
)

var __obf_1b4372ec12e7e260 = reflect.TypeOf((*[]string)(nil))

// DecodeArrayLen decodes array length. Length is -1 when array is nil.
func (__obf_5d389b660eefb08c *Decoder) DecodeArrayLen() (int, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return 0, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_a07a0c18e5d282be(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_a07a0c18e5d282be(__obf_4148125b350dfea2 byte) (int, error) {
	if __obf_4148125b350dfea2 == Nil {
		return -1, nil
	} else if __obf_4148125b350dfea2 >= FixedArrayLow && __obf_4148125b350dfea2 <= FixedArrayHigh {
		return int(__obf_4148125b350dfea2 & FixedArrayMask), nil
	}
	switch __obf_4148125b350dfea2 {
	case Array16:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint16()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	case Array32:
		__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.uint32()
		return int(__obf_d774e4844c74bfe9), __obf_45342a3a754d12f5
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding array length", __obf_4148125b350dfea2)
}

func __obf_d77a8d5f087ab98c(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_2c49f9d2007cfaf6 := __obf_f328a048f76b7256.Addr().Convert(__obf_1b4372ec12e7e260).Interface().(*[]string)
	return __obf_5d389b660eefb08c.__obf_e4607e36ff9c01b4(__obf_2c49f9d2007cfaf6)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_e4607e36ff9c01b4(__obf_2c49f9d2007cfaf6 *[]string) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeArrayLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return nil
	}
	__obf_31c0f1d5a751574c := __obf_d483ee8dd5025272(*__obf_2c49f9d2007cfaf6, __obf_d774e4844c74bfe9, __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_d0b80752c6a68c61 != 0)
	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_d774e4844c74bfe9; __obf_4140b3ff04f75a36++ {
		__obf_fe1ace7a2eb8bf9f, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeString()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		__obf_31c0f1d5a751574c = append(__obf_31c0f1d5a751574c, __obf_fe1ace7a2eb8bf9f)
	}
	*__obf_2c49f9d2007cfaf6 = __obf_31c0f1d5a751574c

	return nil
}

func __obf_d483ee8dd5025272(__obf_fe1ace7a2eb8bf9f []string, __obf_d774e4844c74bfe9 int, __obf_9253e0099c4f88c4 bool) []string {
	if !__obf_9253e0099c4f88c4 && __obf_d774e4844c74bfe9 > __obf_369a9acbbd312147 {
		__obf_d774e4844c74bfe9 = __obf_369a9acbbd312147
	}

	if __obf_fe1ace7a2eb8bf9f == nil {
		return make([]string, 0, __obf_d774e4844c74bfe9)
	}

	if cap(__obf_fe1ace7a2eb8bf9f) >= __obf_d774e4844c74bfe9 {
		return __obf_fe1ace7a2eb8bf9f[:0]
	}
	__obf_fe1ace7a2eb8bf9f = __obf_fe1ace7a2eb8bf9f[:cap(__obf_fe1ace7a2eb8bf9f)]
	__obf_fe1ace7a2eb8bf9f = append(__obf_fe1ace7a2eb8bf9f, make([]string, __obf_d774e4844c74bfe9-len(__obf_fe1ace7a2eb8bf9f))...)
	return __obf_fe1ace7a2eb8bf9f[:0]
}

func __obf_bee2295884206fce(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeArrayLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	if __obf_d774e4844c74bfe9 == -1 {
		__obf_f328a048f76b7256.
			Set(reflect.Zero(__obf_f328a048f76b7256.Type()))
		return nil
	}
	if __obf_d774e4844c74bfe9 == 0 && __obf_f328a048f76b7256.IsNil() {
		__obf_f328a048f76b7256.
			Set(reflect.MakeSlice(__obf_f328a048f76b7256.Type(), 0, 0))
		return nil
	}

	if __obf_f328a048f76b7256.Cap() >= __obf_d774e4844c74bfe9 {
		__obf_f328a048f76b7256.
			Set(__obf_f328a048f76b7256.Slice(0, __obf_d774e4844c74bfe9))
	} else if __obf_f328a048f76b7256.Len() < __obf_f328a048f76b7256.Cap() {
		__obf_f328a048f76b7256.
			Set(__obf_f328a048f76b7256.Slice(0, __obf_f328a048f76b7256.Cap()))
	}
	__obf_9253e0099c4f88c4 := __obf_5d389b660eefb08c.__obf_7185cb62de7af792&__obf_d0b80752c6a68c61 != 1

	if __obf_9253e0099c4f88c4 && __obf_d774e4844c74bfe9 > __obf_f328a048f76b7256.Len() {
		__obf_f328a048f76b7256.
			Set(__obf_7d5892cec89dc78f(__obf_f328a048f76b7256, __obf_d774e4844c74bfe9, __obf_9253e0099c4f88c4))
	}

	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_d774e4844c74bfe9; __obf_4140b3ff04f75a36++ {
		if !__obf_9253e0099c4f88c4 && __obf_4140b3ff04f75a36 >= __obf_f328a048f76b7256.Len() {
			__obf_f328a048f76b7256.
				Set(__obf_7d5892cec89dc78f(__obf_f328a048f76b7256, __obf_d774e4844c74bfe9, __obf_9253e0099c4f88c4))
		}
		__obf_d6cb414a13d9d0f7 := __obf_f328a048f76b7256.Index(__obf_4140b3ff04f75a36)
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeValue(__obf_d6cb414a13d9d0f7); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func __obf_7d5892cec89dc78f(__obf_f328a048f76b7256 reflect.Value, __obf_d774e4844c74bfe9 int, __obf_9253e0099c4f88c4 bool) reflect.Value {
	__obf_aaef2229df52bffd := __obf_d774e4844c74bfe9 - __obf_f328a048f76b7256.Len()
	if !__obf_9253e0099c4f88c4 && __obf_aaef2229df52bffd > __obf_369a9acbbd312147 {
		__obf_aaef2229df52bffd = __obf_369a9acbbd312147
	}
	__obf_f328a048f76b7256 = reflect.AppendSlice(__obf_f328a048f76b7256, reflect.MakeSlice(__obf_f328a048f76b7256.Type(), __obf_aaef2229df52bffd, __obf_aaef2229df52bffd))
	return __obf_f328a048f76b7256
}

func __obf_b694339ec1e870f3(__obf_5d389b660eefb08c *Decoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeArrayLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	if __obf_d774e4844c74bfe9 == -1 {
		return nil
	}
	if __obf_d774e4844c74bfe9 > __obf_f328a048f76b7256.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_f328a048f76b7256.Type(), __obf_f328a048f76b7256.Len(), __obf_d774e4844c74bfe9)
	}

	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_d774e4844c74bfe9; __obf_4140b3ff04f75a36++ {
		__obf_a2cca5fb70136667 := __obf_f328a048f76b7256.Index(__obf_4140b3ff04f75a36)
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeValue(__obf_a2cca5fb70136667); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_5d389b660eefb08c *Decoder) DecodeSlice() ([]any, error) {
	__obf_4148125b350dfea2, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_4786096a37bd32ce()
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	return __obf_5d389b660eefb08c.__obf_e557eb76c54a6f47(__obf_4148125b350dfea2)
}

func (__obf_5d389b660eefb08c *Decoder) __obf_e557eb76c54a6f47(__obf_4148125b350dfea2 byte) ([]any, error) {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_a07a0c18e5d282be(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return nil, nil
	}
	__obf_fe1ace7a2eb8bf9f := make([]any, 0, __obf_d774e4844c74bfe9)
	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_d774e4844c74bfe9; __obf_4140b3ff04f75a36++ {
		__obf_f328a048f76b7256, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
		if __obf_45342a3a754d12f5 != nil {
			return nil, __obf_45342a3a754d12f5
		}
		__obf_fe1ace7a2eb8bf9f = append(__obf_fe1ace7a2eb8bf9f, __obf_f328a048f76b7256)
	}

	return __obf_fe1ace7a2eb8bf9f, nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_8ed189638e6f0a73(__obf_4148125b350dfea2 byte) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_a07a0c18e5d282be(__obf_4148125b350dfea2)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_d774e4844c74bfe9; __obf_4140b3ff04f75a36++ {
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}
