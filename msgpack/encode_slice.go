package __obf_fd770cb9675903b0

import (
	"math"
	"reflect"
)

var __obf_525e1bfa898162e1 = reflect.TypeOf(([]string)(nil))

func __obf_f3b7f586caae0bd6(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeString(__obf_f328a048f76b7256.String())
}

func __obf_898cd0032343e4d6(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	return __obf_e9038cda3b5cf711.EncodeBytes(__obf_f328a048f76b7256.Bytes())
}

func __obf_66038bf3c325450f(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeBytesLen(__obf_f328a048f76b7256.Len()); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	if __obf_f328a048f76b7256.CanAddr() {
		__obf_94333af0f0facd65 := __obf_f328a048f76b7256.Slice(0, __obf_f328a048f76b7256.Len()).Bytes()
		return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_94333af0f0facd65)
	}
	__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a = __obf_efd927a4bbaeb0ca(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a, __obf_f328a048f76b7256.Len())
	reflect.Copy(reflect.ValueOf(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a), __obf_f328a048f76b7256)
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_e9038cda3b5cf711.__obf_099b1f8c6882e66a)
}

func __obf_efd927a4bbaeb0ca(__obf_94333af0f0facd65 []byte, __obf_d774e4844c74bfe9 int) []byte {
	if cap(__obf_94333af0f0facd65) >= __obf_d774e4844c74bfe9 {
		return __obf_94333af0f0facd65[:__obf_d774e4844c74bfe9]
	}
	__obf_94333af0f0facd65 = __obf_94333af0f0facd65[:cap(__obf_94333af0f0facd65)]
	__obf_94333af0f0facd65 = append(__obf_94333af0f0facd65, make([]byte, __obf_d774e4844c74bfe9-len(__obf_94333af0f0facd65))...)
	return __obf_94333af0f0facd65
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeBytesLen(__obf_bddea2836c583aa2 int) error {
	if __obf_bddea2836c583aa2 < 256 {
		return __obf_e9038cda3b5cf711.__obf_c58302e6ea7043c4(Bin8, uint8(__obf_bddea2836c583aa2))
	}
	if __obf_bddea2836c583aa2 <= math.MaxUint16 {
		return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Bin16, uint16(__obf_bddea2836c583aa2))
	}
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Bin32, uint32(__obf_bddea2836c583aa2))
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_c15b1f5a41dd93d4(__obf_bddea2836c583aa2 int) error {
	if __obf_bddea2836c583aa2 < 32 {
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixedStrLow | byte(__obf_bddea2836c583aa2))
	}
	if __obf_bddea2836c583aa2 < 256 {
		return __obf_e9038cda3b5cf711.__obf_c58302e6ea7043c4(Str8, uint8(__obf_bddea2836c583aa2))
	}
	if __obf_bddea2836c583aa2 <= math.MaxUint16 {
		return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Str16, uint16(__obf_bddea2836c583aa2))
	}
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Str32, uint32(__obf_bddea2836c583aa2))
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeString(__obf_f328a048f76b7256 string) error {
	if __obf_4968ec805e7c5822 := __obf_e9038cda3b5cf711.__obf_7185cb62de7af792&__obf_139c02350e62bdf7 != 0; __obf_4968ec805e7c5822 || len(__obf_e9038cda3b5cf711.__obf_ff96db22e12b6842) > 0 {
		return __obf_e9038cda3b5cf711.__obf_27c3beed6bd703be(__obf_f328a048f76b7256, __obf_4968ec805e7c5822)
	}
	return __obf_e9038cda3b5cf711.__obf_8d015d0182bfd5d0(__obf_f328a048f76b7256)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_8d015d0182bfd5d0(__obf_f328a048f76b7256 string) error {
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.__obf_c15b1f5a41dd93d4(len(__obf_f328a048f76b7256)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return __obf_e9038cda3b5cf711.__obf_016cc47c9f09546b(__obf_f328a048f76b7256)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeBytes(__obf_f328a048f76b7256 []byte) error {
	if __obf_f328a048f76b7256 == nil {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeBytesLen(len(__obf_f328a048f76b7256)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	return __obf_e9038cda3b5cf711.__obf_af6d14a7babbd464(__obf_f328a048f76b7256)
}

func (__obf_e9038cda3b5cf711 *Encoder) EncodeArrayLen(__obf_bddea2836c583aa2 int) error {
	if __obf_bddea2836c583aa2 < 16 {
		return __obf_e9038cda3b5cf711.__obf_a72c49227d01f9dd(FixedArrayLow | byte(__obf_bddea2836c583aa2))
	}
	if __obf_bddea2836c583aa2 <= math.MaxUint16 {
		return __obf_e9038cda3b5cf711.__obf_a98b5ca0c9f03534(Array16, uint16(__obf_bddea2836c583aa2))
	}
	return __obf_e9038cda3b5cf711.__obf_e0e3a3f6b942651e(Array32, uint32(__obf_bddea2836c583aa2))
}

func __obf_1d2042271514bd6b(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_31c0f1d5a751574c := __obf_f328a048f76b7256.Convert(__obf_525e1bfa898162e1).Interface().([]string)
	return __obf_e9038cda3b5cf711.__obf_ecab5da9a1ec3de7(__obf_31c0f1d5a751574c)
}

func (__obf_e9038cda3b5cf711 *Encoder) __obf_ecab5da9a1ec3de7(__obf_fe1ace7a2eb8bf9f []string) error {
	if __obf_fe1ace7a2eb8bf9f == nil {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeArrayLen(len(__obf_fe1ace7a2eb8bf9f)); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	for _, __obf_f328a048f76b7256 := range __obf_fe1ace7a2eb8bf9f {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeString(__obf_f328a048f76b7256); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}

func __obf_352d72be21a6a0d7(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	if __obf_f328a048f76b7256.IsNil() {
		return __obf_e9038cda3b5cf711.EncodeNil()
	}
	return __obf_2e9719daceaf70e4(__obf_e9038cda3b5cf711, __obf_f328a048f76b7256)
}

func __obf_2e9719daceaf70e4(__obf_e9038cda3b5cf711 *Encoder, __obf_f328a048f76b7256 reflect.Value) error {
	__obf_bddea2836c583aa2 := __obf_f328a048f76b7256.Len()
	if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeArrayLen(__obf_bddea2836c583aa2); __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	for __obf_4140b3ff04f75a36 := 0; __obf_4140b3ff04f75a36 < __obf_bddea2836c583aa2; __obf_4140b3ff04f75a36++ {
		if __obf_45342a3a754d12f5 := __obf_e9038cda3b5cf711.EncodeValue(__obf_f328a048f76b7256.Index(__obf_4140b3ff04f75a36)); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}
