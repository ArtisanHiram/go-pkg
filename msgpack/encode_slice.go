package __obf_3e0c303610a19bd4

import (
	"math"
	"reflect"
)

var __obf_e8f2be337b3fb6a9 = reflect.TypeOf(([]string)(nil))

func __obf_4013add33227fac1(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeString(__obf_63bbcee86d44fdde.String())
}

func __obf_3242ef092432362b(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeBytes(__obf_63bbcee86d44fdde.Bytes())
}

func __obf_39d5205b9d649841(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeBytesLen(__obf_63bbcee86d44fdde.Len()); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	if __obf_63bbcee86d44fdde.CanAddr() {
		__obf_11bcc66cde095c11 := __obf_63bbcee86d44fdde.Slice(0, __obf_63bbcee86d44fdde.Len()).Bytes()
		return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_11bcc66cde095c11)
	}
	__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7 = __obf_b1e02f9df5acfb33(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7, __obf_63bbcee86d44fdde.Len())
	reflect.Copy(reflect.ValueOf(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7), __obf_63bbcee86d44fdde)
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_77240dc7776b60b4.__obf_27e4a1a33a31e4a7)
}

func __obf_b1e02f9df5acfb33(__obf_11bcc66cde095c11 []byte, __obf_4909ae60ffbb8e53 int) []byte {
	if cap(__obf_11bcc66cde095c11) >= __obf_4909ae60ffbb8e53 {
		return __obf_11bcc66cde095c11[:__obf_4909ae60ffbb8e53]
	}
	__obf_11bcc66cde095c11 = __obf_11bcc66cde095c11[:cap(__obf_11bcc66cde095c11)]
	__obf_11bcc66cde095c11 = append(__obf_11bcc66cde095c11, make([]byte, __obf_4909ae60ffbb8e53-len(__obf_11bcc66cde095c11))...)
	return __obf_11bcc66cde095c11
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeBytesLen(__obf_d62f82ed950927da int) error {
	if __obf_d62f82ed950927da < 256 {
		return __obf_77240dc7776b60b4.__obf_a5ea8b1a318bf9eb(Bin8, uint8(__obf_d62f82ed950927da))
	}
	if __obf_d62f82ed950927da <= math.MaxUint16 {
		return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Bin16, uint16(__obf_d62f82ed950927da))
	}
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Bin32, uint32(__obf_d62f82ed950927da))
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_07e8b2f917ce99e7(__obf_d62f82ed950927da int) error {
	if __obf_d62f82ed950927da < 32 {
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixedStrLow | byte(__obf_d62f82ed950927da))
	}
	if __obf_d62f82ed950927da < 256 {
		return __obf_77240dc7776b60b4.__obf_a5ea8b1a318bf9eb(Str8, uint8(__obf_d62f82ed950927da))
	}
	if __obf_d62f82ed950927da <= math.MaxUint16 {
		return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Str16, uint16(__obf_d62f82ed950927da))
	}
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Str32, uint32(__obf_d62f82ed950927da))
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeString(__obf_63bbcee86d44fdde string) error {
	if __obf_14e539ac760532b2 := __obf_77240dc7776b60b4.__obf_3cf0882fa5a4cafb&__obf_a063db36bdbf4202 != 0; __obf_14e539ac760532b2 || len(__obf_77240dc7776b60b4.__obf_a22a31b815544cae) > 0 {
		return __obf_77240dc7776b60b4.__obf_553232b05cafcb2b(__obf_63bbcee86d44fdde, __obf_14e539ac760532b2)
	}
	return __obf_77240dc7776b60b4.__obf_a30f8acd02da8779(__obf_63bbcee86d44fdde)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_a30f8acd02da8779(__obf_63bbcee86d44fdde string) error {
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.__obf_07e8b2f917ce99e7(len(__obf_63bbcee86d44fdde)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return __obf_77240dc7776b60b4.__obf_21c7568bc3d83316(__obf_63bbcee86d44fdde)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeBytes(__obf_63bbcee86d44fdde []byte) error {
	if __obf_63bbcee86d44fdde == nil {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeBytesLen(len(__obf_63bbcee86d44fdde)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	return __obf_77240dc7776b60b4.__obf_4432f415b3007804(__obf_63bbcee86d44fdde)
}

func (__obf_77240dc7776b60b4 *Encoder) EncodeArrayLen(__obf_d62f82ed950927da int) error {
	if __obf_d62f82ed950927da < 16 {
		return __obf_77240dc7776b60b4.__obf_e8d9aa5d4f5eb974(FixedArrayLow | byte(__obf_d62f82ed950927da))
	}
	if __obf_d62f82ed950927da <= math.MaxUint16 {
		return __obf_77240dc7776b60b4.__obf_9ecc220d5defc1d9(Array16, uint16(__obf_d62f82ed950927da))
	}
	return __obf_77240dc7776b60b4.__obf_70cbe6e3e07b6347(Array32, uint32(__obf_d62f82ed950927da))
}

func __obf_f5ae91e031add6e1(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_66e8d01dfb8a3535 := __obf_63bbcee86d44fdde.Convert(__obf_e8f2be337b3fb6a9).Interface().([]string)
	return __obf_77240dc7776b60b4.__obf_826b9874ef1e0720(__obf_66e8d01dfb8a3535)
}

func (__obf_77240dc7776b60b4 *Encoder) __obf_826b9874ef1e0720(__obf_61027e0491b6dd3d []string) error {
	if __obf_61027e0491b6dd3d == nil {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeArrayLen(len(__obf_61027e0491b6dd3d)); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	for _, __obf_63bbcee86d44fdde := range __obf_61027e0491b6dd3d {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeString(__obf_63bbcee86d44fdde); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}

func __obf_ddcfd0255d36e7b4(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	return __obf_fb2bd25c8bdcdcf8(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde)
}

func __obf_fb2bd25c8bdcdcf8(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_d62f82ed950927da := __obf_63bbcee86d44fdde.Len()
	if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeArrayLen(__obf_d62f82ed950927da); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	for __obf_39714879601f9b69 := 0; __obf_39714879601f9b69 < __obf_d62f82ed950927da; __obf_39714879601f9b69++ {
		if __obf_8882f6cf6e378ded := __obf_77240dc7776b60b4.EncodeValue(__obf_63bbcee86d44fdde.Index(__obf_39714879601f9b69)); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}
