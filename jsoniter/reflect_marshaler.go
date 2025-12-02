package __obf_c7b79b12b33d8238

import (
	"encoding"
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_369a01837b5e3e42 = reflect2.TypeOfPtr((*json.Marshaler)(nil)).Elem()
var __obf_f5c4044aeb40e807 = reflect2.TypeOfPtr((*json.Unmarshaler)(nil)).Elem()
var __obf_c3f942f0f1ba7a21 = reflect2.TypeOfPtr((*encoding.TextMarshaler)(nil)).Elem()
var __obf_1233738a1f65c955 = reflect2.TypeOfPtr((*encoding.TextUnmarshaler)(nil)).Elem()

func __obf_2c53588571656d99(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValDecoder {
	__obf_e2840a6d1d1a664b := reflect2.PtrTo(__obf_edcded11af6ebc4c)
	if __obf_e2840a6d1d1a664b.Implements(__obf_f5c4044aeb40e807) {
		return &__obf_488d5d64f097f8dc{
			&__obf_7829690424eb3cfa{__obf_e2840a6d1d1a664b},
		}
	}
	if __obf_e2840a6d1d1a664b.Implements(__obf_1233738a1f65c955) {
		return &__obf_488d5d64f097f8dc{
			&__obf_1ed943b6487d22ca{__obf_e2840a6d1d1a664b},
		}
	}
	return nil
}

func __obf_5ddc38919c4ef5f0(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	if __obf_edcded11af6ebc4c == __obf_369a01837b5e3e42 {
		__obf_07e25a5cafedd0df := __obf_347e3fc852eb458f(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
		var __obf_c091c27480fae550 ValEncoder = &__obf_2cc8be7fb1a1b955{__obf_07e25a5cafedd0df: __obf_07e25a5cafedd0df}
		return __obf_c091c27480fae550
	}
	if __obf_edcded11af6ebc4c.Implements(__obf_369a01837b5e3e42) {
		__obf_07e25a5cafedd0df := __obf_347e3fc852eb458f(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
		var __obf_c091c27480fae550 ValEncoder = &__obf_db1cc37ca0e93fce{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c, __obf_07e25a5cafedd0df: __obf_07e25a5cafedd0df}
		return __obf_c091c27480fae550
	}
	__obf_e2840a6d1d1a664b := reflect2.PtrTo(__obf_edcded11af6ebc4c)
	if __obf_99ec45908bceabdb.__obf_5de9ece8fa3a16e3 != "" && __obf_e2840a6d1d1a664b.Implements(__obf_369a01837b5e3e42) {
		__obf_07e25a5cafedd0df := __obf_347e3fc852eb458f(__obf_99ec45908bceabdb, __obf_e2840a6d1d1a664b)
		var __obf_c091c27480fae550 ValEncoder = &__obf_db1cc37ca0e93fce{__obf_e0ba73c3f13f4455: __obf_e2840a6d1d1a664b, __obf_07e25a5cafedd0df: __obf_07e25a5cafedd0df}
		return &__obf_f992d744271af46a{__obf_c091c27480fae550}
	}
	if __obf_edcded11af6ebc4c == __obf_c3f942f0f1ba7a21 {
		__obf_07e25a5cafedd0df := __obf_347e3fc852eb458f(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
		var __obf_c091c27480fae550 ValEncoder = &__obf_848eb91cb8d800f4{__obf_07e25a5cafedd0df: __obf_07e25a5cafedd0df, __obf_7ae2c2c658aa1916: __obf_99ec45908bceabdb.EncoderOf(reflect2.TypeOf(""))}
		return __obf_c091c27480fae550
	}
	if __obf_edcded11af6ebc4c.Implements(__obf_c3f942f0f1ba7a21) {
		__obf_07e25a5cafedd0df := __obf_347e3fc852eb458f(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
		var __obf_c091c27480fae550 ValEncoder = &__obf_c2c5222d98675825{__obf_e0ba73c3f13f4455: __obf_edcded11af6ebc4c, __obf_7ae2c2c658aa1916: __obf_99ec45908bceabdb.EncoderOf(reflect2.TypeOf("")), __obf_07e25a5cafedd0df: __obf_07e25a5cafedd0df}
		return __obf_c091c27480fae550
	}
	// if prefix is empty, the type is the root type
	if __obf_99ec45908bceabdb.__obf_5de9ece8fa3a16e3 != "" && __obf_e2840a6d1d1a664b.Implements(__obf_c3f942f0f1ba7a21) {
		__obf_07e25a5cafedd0df := __obf_347e3fc852eb458f(__obf_99ec45908bceabdb, __obf_e2840a6d1d1a664b)
		var __obf_c091c27480fae550 ValEncoder = &__obf_c2c5222d98675825{__obf_e0ba73c3f13f4455: __obf_e2840a6d1d1a664b, __obf_7ae2c2c658aa1916: __obf_99ec45908bceabdb.EncoderOf(reflect2.TypeOf("")), __obf_07e25a5cafedd0df: __obf_07e25a5cafedd0df}
		return &__obf_f992d744271af46a{__obf_c091c27480fae550}
	}
	return nil
}

type __obf_db1cc37ca0e93fce struct {
	__obf_07e25a5cafedd0df __obf_07e25a5cafedd0df
	__obf_e0ba73c3f13f4455 reflect2.Type
}

func (__obf_c091c27480fae550 *__obf_db1cc37ca0e93fce) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d6e2df4782353c64 := __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	if __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.IsNullable() && reflect2.IsNil(__obf_d6e2df4782353c64) {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_9948ab09ee797b6e := __obf_d6e2df4782353c64.(json.Marshaler)
	__obf_290a519ecf5a657d, __obf_ea0680f8b461a85b := __obf_9948ab09ee797b6e.MarshalJSON()
	if __obf_ea0680f8b461a85b != nil {
		__obf_d8f50bcfe87d8b47.
			Error = __obf_ea0680f8b461a85b
	} else {
		__obf_081eec0c3bb95313 := // html escape was already done by jsoniter
			// but the extra '\n' should be trimed
			len(__obf_290a519ecf5a657d)
		if __obf_081eec0c3bb95313 > 0 && __obf_290a519ecf5a657d[__obf_081eec0c3bb95313-1] == '\n' {
			__obf_290a519ecf5a657d = __obf_290a519ecf5a657d[:__obf_081eec0c3bb95313-1]
		}
		__obf_d8f50bcfe87d8b47.
			Write(__obf_290a519ecf5a657d)
	}
}

func (__obf_c091c27480fae550 *__obf_db1cc37ca0e93fce) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_07e25a5cafedd0df.IsEmpty(__obf_575c04f2b9d91315)
}

type __obf_2cc8be7fb1a1b955 struct {
	__obf_07e25a5cafedd0df __obf_07e25a5cafedd0df
}

func (__obf_c091c27480fae550 *__obf_2cc8be7fb1a1b955) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_9948ab09ee797b6e := *(*json.Marshaler)(__obf_575c04f2b9d91315)
	if __obf_9948ab09ee797b6e == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_290a519ecf5a657d, __obf_ea0680f8b461a85b := __obf_9948ab09ee797b6e.MarshalJSON()
	if __obf_ea0680f8b461a85b != nil {
		__obf_d8f50bcfe87d8b47.
			Error = __obf_ea0680f8b461a85b
	} else {
		__obf_d8f50bcfe87d8b47.
			Write(__obf_290a519ecf5a657d)
	}
}

func (__obf_c091c27480fae550 *__obf_2cc8be7fb1a1b955) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_07e25a5cafedd0df.IsEmpty(__obf_575c04f2b9d91315)
}

type __obf_c2c5222d98675825 struct {
	__obf_e0ba73c3f13f4455 reflect2.Type
	__obf_7ae2c2c658aa1916 ValEncoder
	__obf_07e25a5cafedd0df __obf_07e25a5cafedd0df
}

func (__obf_c091c27480fae550 *__obf_c2c5222d98675825) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d6e2df4782353c64 := __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	if __obf_c091c27480fae550.__obf_e0ba73c3f13f4455.IsNullable() && reflect2.IsNil(__obf_d6e2df4782353c64) {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_9948ab09ee797b6e := (__obf_d6e2df4782353c64).(encoding.TextMarshaler)
	__obf_290a519ecf5a657d, __obf_ea0680f8b461a85b := __obf_9948ab09ee797b6e.MarshalText()
	if __obf_ea0680f8b461a85b != nil {
		__obf_d8f50bcfe87d8b47.
			Error = __obf_ea0680f8b461a85b
	} else {
		__obf_a3eaafc22faf1644 := string(__obf_290a519ecf5a657d)
		__obf_c091c27480fae550.__obf_7ae2c2c658aa1916.
			Encode(unsafe.Pointer(&__obf_a3eaafc22faf1644), __obf_d8f50bcfe87d8b47)
	}
}

func (__obf_c091c27480fae550 *__obf_c2c5222d98675825) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_07e25a5cafedd0df.IsEmpty(__obf_575c04f2b9d91315)
}

type __obf_848eb91cb8d800f4 struct {
	__obf_7ae2c2c658aa1916 ValEncoder
	__obf_07e25a5cafedd0df __obf_07e25a5cafedd0df
}

func (__obf_c091c27480fae550 *__obf_848eb91cb8d800f4) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_9948ab09ee797b6e := *(*encoding.TextMarshaler)(__obf_575c04f2b9d91315)
	if __obf_9948ab09ee797b6e == nil {
		__obf_d8f50bcfe87d8b47.
			WriteNil()
		return
	}
	__obf_290a519ecf5a657d, __obf_ea0680f8b461a85b := __obf_9948ab09ee797b6e.MarshalText()
	if __obf_ea0680f8b461a85b != nil {
		__obf_d8f50bcfe87d8b47.
			Error = __obf_ea0680f8b461a85b
	} else {
		__obf_a3eaafc22faf1644 := string(__obf_290a519ecf5a657d)
		__obf_c091c27480fae550.__obf_7ae2c2c658aa1916.
			Encode(unsafe.Pointer(&__obf_a3eaafc22faf1644), __obf_d8f50bcfe87d8b47)
	}
}

func (__obf_c091c27480fae550 *__obf_848eb91cb8d800f4) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_07e25a5cafedd0df.IsEmpty(__obf_575c04f2b9d91315)
}

type __obf_7829690424eb3cfa struct {
	__obf_e0ba73c3f13f4455 reflect2.Type
}

func (__obf_801f19702638809c *__obf_7829690424eb3cfa) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_e0ba73c3f13f4455 := __obf_801f19702638809c.__obf_e0ba73c3f13f4455
	__obf_d6e2df4782353c64 := __obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	__obf_f123f58974b18049 := __obf_d6e2df4782353c64.(json.Unmarshaler)
	__obf_0da8c843dad13139.__obf_2b436fcb1c0ca36d()
	__obf_0da8c843dad13139.
		// skip spaces
		__obf_a675366c80290b83()
	__obf_290a519ecf5a657d := __obf_0da8c843dad13139.SkipAndReturnBytes()
	__obf_ea0680f8b461a85b := __obf_f123f58974b18049.UnmarshalJSON(__obf_290a519ecf5a657d)
	if __obf_ea0680f8b461a85b != nil {
		__obf_0da8c843dad13139.
			ReportError("unmarshalerDecoder", __obf_ea0680f8b461a85b.Error())
	}
}

type __obf_1ed943b6487d22ca struct {
	__obf_e0ba73c3f13f4455 reflect2.Type
}

func (__obf_801f19702638809c *__obf_1ed943b6487d22ca) Decode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_0da8c843dad13139 *Iterator) {
	__obf_e0ba73c3f13f4455 := __obf_801f19702638809c.__obf_e0ba73c3f13f4455
	__obf_d6e2df4782353c64 := __obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	if reflect2.IsNil(__obf_d6e2df4782353c64) {
		__obf_e2840a6d1d1a664b := __obf_e0ba73c3f13f4455.(*reflect2.UnsafePtrType)
		__obf_f5ccfed08a3cf863 := __obf_e2840a6d1d1a664b.Elem()
		__obf_43cefbbf0648e3be := __obf_f5ccfed08a3cf863.UnsafeNew()
		__obf_e2840a6d1d1a664b.
			UnsafeSet(__obf_575c04f2b9d91315, unsafe.Pointer(&__obf_43cefbbf0648e3be))
		__obf_d6e2df4782353c64 = __obf_e0ba73c3f13f4455.UnsafeIndirect(__obf_575c04f2b9d91315)
	}
	__obf_f123f58974b18049 := (__obf_d6e2df4782353c64).(encoding.TextUnmarshaler)
	__obf_a3eaafc22faf1644 := __obf_0da8c843dad13139.ReadString()
	__obf_ea0680f8b461a85b := __obf_f123f58974b18049.UnmarshalText([]byte(__obf_a3eaafc22faf1644))
	if __obf_ea0680f8b461a85b != nil {
		__obf_0da8c843dad13139.
			ReportError("textUnmarshalerDecoder", __obf_ea0680f8b461a85b.Error())
	}
}
