package __obf_a935eb7f548271a4

import (
	"math"
	"reflect"
)

var __obf_c02e26ace76ac379 = reflect.TypeOf(([]string)(nil))

func __obf_96718da505bbd637(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeString(__obf_6d570581f4b60dbc.String())
}

func __obf_892a8b5216486b44(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeBytes(__obf_6d570581f4b60dbc.Bytes())
}

func __obf_ff83fd93faa9eb52(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeBytesLen(__obf_6d570581f4b60dbc.Len()); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	if __obf_6d570581f4b60dbc.CanAddr() {
		__obf_f2ca794293605b73 := __obf_6d570581f4b60dbc.Slice(0, __obf_6d570581f4b60dbc.Len()).Bytes()
		return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_f2ca794293605b73)
	}
	__obf_ed37a34c347049f3.__obf_a2e16952f247f218 = __obf_8b2bbc80a0defb4e(__obf_ed37a34c347049f3.__obf_a2e16952f247f218, __obf_6d570581f4b60dbc.Len())
	reflect.Copy(reflect.ValueOf(__obf_ed37a34c347049f3.__obf_a2e16952f247f218), __obf_6d570581f4b60dbc)
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_ed37a34c347049f3.__obf_a2e16952f247f218)
}

func __obf_8b2bbc80a0defb4e(__obf_f2ca794293605b73 []byte, __obf_326af9bd942662ac int) []byte {
	if cap(__obf_f2ca794293605b73) >= __obf_326af9bd942662ac {
		return __obf_f2ca794293605b73[:__obf_326af9bd942662ac]
	}
	__obf_f2ca794293605b73 = __obf_f2ca794293605b73[:cap(__obf_f2ca794293605b73)]
	__obf_f2ca794293605b73 = append(__obf_f2ca794293605b73, make([]byte, __obf_326af9bd942662ac-len(__obf_f2ca794293605b73))...)
	return __obf_f2ca794293605b73
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeBytesLen(__obf_3c213c92f0597e4d int) error {
	if __obf_3c213c92f0597e4d < 256 {
		return __obf_ed37a34c347049f3.__obf_327ee6f3f671bfd3(Bin8, uint8(__obf_3c213c92f0597e4d))
	}
	if __obf_3c213c92f0597e4d <= math.MaxUint16 {
		return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Bin16, uint16(__obf_3c213c92f0597e4d))
	}
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Bin32, uint32(__obf_3c213c92f0597e4d))
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_d98ff735b05cc87d(__obf_3c213c92f0597e4d int) error {
	if __obf_3c213c92f0597e4d < 32 {
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixedStrLow | byte(__obf_3c213c92f0597e4d))
	}
	if __obf_3c213c92f0597e4d < 256 {
		return __obf_ed37a34c347049f3.__obf_327ee6f3f671bfd3(Str8, uint8(__obf_3c213c92f0597e4d))
	}
	if __obf_3c213c92f0597e4d <= math.MaxUint16 {
		return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Str16, uint16(__obf_3c213c92f0597e4d))
	}
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Str32, uint32(__obf_3c213c92f0597e4d))
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeString(__obf_6d570581f4b60dbc string) error {
	if __obf_65a759dbf3ace040 := __obf_ed37a34c347049f3.__obf_d617f3769ce16c47&__obf_bf24fd2d64b91255 != 0; __obf_65a759dbf3ace040 || len(__obf_ed37a34c347049f3.__obf_2cdd7fad308aef1d) > 0 {
		return __obf_ed37a34c347049f3.__obf_f3306d3e9feeada5(__obf_6d570581f4b60dbc, __obf_65a759dbf3ace040)
	}
	return __obf_ed37a34c347049f3.__obf_dd784acb29c85421(__obf_6d570581f4b60dbc)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_dd784acb29c85421(__obf_6d570581f4b60dbc string) error {
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.__obf_d98ff735b05cc87d(len(__obf_6d570581f4b60dbc)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return __obf_ed37a34c347049f3.__obf_6515630c78622bba(__obf_6d570581f4b60dbc)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeBytes(__obf_6d570581f4b60dbc []byte) error {
	if __obf_6d570581f4b60dbc == nil {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeBytesLen(len(__obf_6d570581f4b60dbc)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	return __obf_ed37a34c347049f3.__obf_654725fc7d18ae48(__obf_6d570581f4b60dbc)
}

func (__obf_ed37a34c347049f3 *Encoder) EncodeArrayLen(__obf_3c213c92f0597e4d int) error {
	if __obf_3c213c92f0597e4d < 16 {
		return __obf_ed37a34c347049f3.__obf_2252dbfd2473b643(FixedArrayLow | byte(__obf_3c213c92f0597e4d))
	}
	if __obf_3c213c92f0597e4d <= math.MaxUint16 {
		return __obf_ed37a34c347049f3.__obf_dd92b251146e86fb(Array16, uint16(__obf_3c213c92f0597e4d))
	}
	return __obf_ed37a34c347049f3.__obf_21945a55177e9fbd(Array32, uint32(__obf_3c213c92f0597e4d))
}

func __obf_cea67acf0689ac24(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_cc04a11b8213c893 := __obf_6d570581f4b60dbc.Convert(__obf_c02e26ace76ac379).Interface().([]string)
	return __obf_ed37a34c347049f3.__obf_4c4703eb1a8e28f5(__obf_cc04a11b8213c893)
}

func (__obf_ed37a34c347049f3 *Encoder) __obf_4c4703eb1a8e28f5(__obf_b62c60fba2fd9788 []string) error {
	if __obf_b62c60fba2fd9788 == nil {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeArrayLen(len(__obf_b62c60fba2fd9788)); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	for _, __obf_6d570581f4b60dbc := range __obf_b62c60fba2fd9788 {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeString(__obf_6d570581f4b60dbc); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}

func __obf_e4145022d977f5f9(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	return __obf_8b865df2f8afbfbe(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc)
}

func __obf_8b865df2f8afbfbe(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	__obf_3c213c92f0597e4d := __obf_6d570581f4b60dbc.Len()
	if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeArrayLen(__obf_3c213c92f0597e4d); __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	for __obf_4a8e280ffaa52cf4 := 0; __obf_4a8e280ffaa52cf4 < __obf_3c213c92f0597e4d; __obf_4a8e280ffaa52cf4++ {
		if __obf_4d327e1cd40c2e21 := __obf_ed37a34c347049f3.EncodeValue(__obf_6d570581f4b60dbc.Index(__obf_4a8e280ffaa52cf4)); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}
