package __obf_703d23ba520c3295

import (
	"encoding"
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_77e9f4c46ca05c12 = reflect2.TypeOfPtr((*json.Marshaler)(nil)).Elem()
var __obf_5be0c81d693cd68b = reflect2.TypeOfPtr((*json.Unmarshaler)(nil)).Elem()
var __obf_04d34c6784ab4d74 = reflect2.TypeOfPtr((*encoding.TextMarshaler)(nil)).Elem()
var __obf_9b7bc9dc1ee32b52 = reflect2.TypeOfPtr((*encoding.TextUnmarshaler)(nil)).Elem()

func __obf_8983286d3baa6ceb(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_95093efb9321684e := reflect2.PtrTo(__obf_3b7f6abbae19451e)
	if __obf_95093efb9321684e.Implements(__obf_5be0c81d693cd68b) {
		return &__obf_8eed859ba47bd05c{
			&__obf_71c95f6b15d8fce6{__obf_95093efb9321684e},
		}
	}
	if __obf_95093efb9321684e.Implements(__obf_9b7bc9dc1ee32b52) {
		return &__obf_8eed859ba47bd05c{
			&__obf_87fa515d0cb34be8{__obf_95093efb9321684e},
		}
	}
	return nil
}

func __obf_888be1be2b861d32(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	if __obf_3b7f6abbae19451e == __obf_77e9f4c46ca05c12 {
		__obf_8daa57d6b3a720c3 := __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
		var __obf_cf840243a12ee308 ValEncoder = &__obf_cd7998eaf02f70f4{__obf_8daa57d6b3a720c3: __obf_8daa57d6b3a720c3}
		return __obf_cf840243a12ee308
	}
	if __obf_3b7f6abbae19451e.Implements(__obf_77e9f4c46ca05c12) {
		__obf_8daa57d6b3a720c3 := __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
		var __obf_cf840243a12ee308 ValEncoder = &__obf_c9232fe8e58e0add{__obf_15da62f311934a45: __obf_3b7f6abbae19451e, __obf_8daa57d6b3a720c3: __obf_8daa57d6b3a720c3}
		return __obf_cf840243a12ee308
	}
	__obf_95093efb9321684e := reflect2.PtrTo(__obf_3b7f6abbae19451e)
	if __obf_2aaf7367de3ff86d.__obf_c5fd316f3c4457de != "" && __obf_95093efb9321684e.Implements(__obf_77e9f4c46ca05c12) {
		__obf_8daa57d6b3a720c3 := __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d, __obf_95093efb9321684e)
		var __obf_cf840243a12ee308 ValEncoder = &__obf_c9232fe8e58e0add{__obf_15da62f311934a45: __obf_95093efb9321684e, __obf_8daa57d6b3a720c3: __obf_8daa57d6b3a720c3}
		return &__obf_1de20804cb7c71ec{__obf_cf840243a12ee308}
	}
	if __obf_3b7f6abbae19451e == __obf_04d34c6784ab4d74 {
		__obf_8daa57d6b3a720c3 := __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
		var __obf_cf840243a12ee308 ValEncoder = &__obf_0a029e0e5aaa3fdf{__obf_8daa57d6b3a720c3: __obf_8daa57d6b3a720c3, __obf_cf5770f0ffaef5fe: __obf_2aaf7367de3ff86d.EncoderOf(reflect2.TypeOf(""))}
		return __obf_cf840243a12ee308
	}
	if __obf_3b7f6abbae19451e.Implements(__obf_04d34c6784ab4d74) {
		__obf_8daa57d6b3a720c3 := __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
		var __obf_cf840243a12ee308 ValEncoder = &__obf_43aa2e0944cf758f{__obf_15da62f311934a45: __obf_3b7f6abbae19451e, __obf_cf5770f0ffaef5fe: __obf_2aaf7367de3ff86d.EncoderOf(reflect2.TypeOf("")), __obf_8daa57d6b3a720c3: __obf_8daa57d6b3a720c3}
		return __obf_cf840243a12ee308
	}
	// if prefix is empty, the type is the root type
	if __obf_2aaf7367de3ff86d.__obf_c5fd316f3c4457de != "" && __obf_95093efb9321684e.Implements(__obf_04d34c6784ab4d74) {
		__obf_8daa57d6b3a720c3 := __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d, __obf_95093efb9321684e)
		var __obf_cf840243a12ee308 ValEncoder = &__obf_43aa2e0944cf758f{__obf_15da62f311934a45: __obf_95093efb9321684e, __obf_cf5770f0ffaef5fe: __obf_2aaf7367de3ff86d.EncoderOf(reflect2.TypeOf("")), __obf_8daa57d6b3a720c3: __obf_8daa57d6b3a720c3}
		return &__obf_1de20804cb7c71ec{__obf_cf840243a12ee308}
	}
	return nil
}

type __obf_c9232fe8e58e0add struct {
	__obf_8daa57d6b3a720c3 __obf_8daa57d6b3a720c3
	__obf_15da62f311934a45 reflect2.Type
}

func (__obf_cf840243a12ee308 *__obf_c9232fe8e58e0add) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_02ba706f4f104d71 := __obf_cf840243a12ee308.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	if __obf_cf840243a12ee308.__obf_15da62f311934a45.IsNullable() && reflect2.IsNil(__obf_02ba706f4f104d71) {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_efc30cdfc2bbd703 := __obf_02ba706f4f104d71.(json.Marshaler)
	__obf_2dce8d733190f3f1, __obf_e56742d6e52f642e := __obf_efc30cdfc2bbd703.MarshalJSON()
	if __obf_e56742d6e52f642e != nil {
		__obf_9a34f62857fb1d1d.
			Error = __obf_e56742d6e52f642e
	} else {
		__obf_60df492a94bf0073 := // html escape was already done by jsoniter
			// but the extra '\n' should be trimed
			len(__obf_2dce8d733190f3f1)
		if __obf_60df492a94bf0073 > 0 && __obf_2dce8d733190f3f1[__obf_60df492a94bf0073-1] == '\n' {
			__obf_2dce8d733190f3f1 = __obf_2dce8d733190f3f1[:__obf_60df492a94bf0073-1]
		}
		__obf_9a34f62857fb1d1d.
			Write(__obf_2dce8d733190f3f1)
	}
}

func (__obf_cf840243a12ee308 *__obf_c9232fe8e58e0add) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_8daa57d6b3a720c3.IsEmpty(__obf_47fa53f3d161f45c)
}

type __obf_cd7998eaf02f70f4 struct {
	__obf_8daa57d6b3a720c3 __obf_8daa57d6b3a720c3
}

func (__obf_cf840243a12ee308 *__obf_cd7998eaf02f70f4) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_efc30cdfc2bbd703 := *(*json.Marshaler)(__obf_47fa53f3d161f45c)
	if __obf_efc30cdfc2bbd703 == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_2dce8d733190f3f1, __obf_e56742d6e52f642e := __obf_efc30cdfc2bbd703.MarshalJSON()
	if __obf_e56742d6e52f642e != nil {
		__obf_9a34f62857fb1d1d.
			Error = __obf_e56742d6e52f642e
	} else {
		__obf_9a34f62857fb1d1d.
			Write(__obf_2dce8d733190f3f1)
	}
}

func (__obf_cf840243a12ee308 *__obf_cd7998eaf02f70f4) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_8daa57d6b3a720c3.IsEmpty(__obf_47fa53f3d161f45c)
}

type __obf_43aa2e0944cf758f struct {
	__obf_15da62f311934a45 reflect2.Type
	__obf_cf5770f0ffaef5fe ValEncoder
	__obf_8daa57d6b3a720c3 __obf_8daa57d6b3a720c3
}

func (__obf_cf840243a12ee308 *__obf_43aa2e0944cf758f) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_02ba706f4f104d71 := __obf_cf840243a12ee308.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	if __obf_cf840243a12ee308.__obf_15da62f311934a45.IsNullable() && reflect2.IsNil(__obf_02ba706f4f104d71) {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_efc30cdfc2bbd703 := (__obf_02ba706f4f104d71).(encoding.TextMarshaler)
	__obf_2dce8d733190f3f1, __obf_e56742d6e52f642e := __obf_efc30cdfc2bbd703.MarshalText()
	if __obf_e56742d6e52f642e != nil {
		__obf_9a34f62857fb1d1d.
			Error = __obf_e56742d6e52f642e
	} else {
		__obf_44b48c26051f8033 := string(__obf_2dce8d733190f3f1)
		__obf_cf840243a12ee308.__obf_cf5770f0ffaef5fe.
			Encode(unsafe.Pointer(&__obf_44b48c26051f8033), __obf_9a34f62857fb1d1d)
	}
}

func (__obf_cf840243a12ee308 *__obf_43aa2e0944cf758f) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_8daa57d6b3a720c3.IsEmpty(__obf_47fa53f3d161f45c)
}

type __obf_0a029e0e5aaa3fdf struct {
	__obf_cf5770f0ffaef5fe ValEncoder
	__obf_8daa57d6b3a720c3 __obf_8daa57d6b3a720c3
}

func (__obf_cf840243a12ee308 *__obf_0a029e0e5aaa3fdf) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_efc30cdfc2bbd703 := *(*encoding.TextMarshaler)(__obf_47fa53f3d161f45c)
	if __obf_efc30cdfc2bbd703 == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_2dce8d733190f3f1, __obf_e56742d6e52f642e := __obf_efc30cdfc2bbd703.MarshalText()
	if __obf_e56742d6e52f642e != nil {
		__obf_9a34f62857fb1d1d.
			Error = __obf_e56742d6e52f642e
	} else {
		__obf_44b48c26051f8033 := string(__obf_2dce8d733190f3f1)
		__obf_cf840243a12ee308.__obf_cf5770f0ffaef5fe.
			Encode(unsafe.Pointer(&__obf_44b48c26051f8033), __obf_9a34f62857fb1d1d)
	}
}

func (__obf_cf840243a12ee308 *__obf_0a029e0e5aaa3fdf) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_8daa57d6b3a720c3.IsEmpty(__obf_47fa53f3d161f45c)
}

type __obf_71c95f6b15d8fce6 struct {
	__obf_15da62f311934a45 reflect2.Type
}

func (__obf_c9719e68325f7d44 *__obf_71c95f6b15d8fce6) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_15da62f311934a45 := __obf_c9719e68325f7d44.__obf_15da62f311934a45
	__obf_02ba706f4f104d71 := __obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	__obf_8e93113c2a45b2c0 := __obf_02ba706f4f104d71.(json.Unmarshaler)
	__obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	__obf_47edab4c16a2d88d.
		// skip spaces
		__obf_743b1a47c346c8a3()
	__obf_2dce8d733190f3f1 := __obf_47edab4c16a2d88d.SkipAndReturnBytes()
	__obf_e56742d6e52f642e := __obf_8e93113c2a45b2c0.UnmarshalJSON(__obf_2dce8d733190f3f1)
	if __obf_e56742d6e52f642e != nil {
		__obf_47edab4c16a2d88d.
			ReportError("unmarshalerDecoder", __obf_e56742d6e52f642e.Error())
	}
}

type __obf_87fa515d0cb34be8 struct {
	__obf_15da62f311934a45 reflect2.Type
}

func (__obf_c9719e68325f7d44 *__obf_87fa515d0cb34be8) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_15da62f311934a45 := __obf_c9719e68325f7d44.__obf_15da62f311934a45
	__obf_02ba706f4f104d71 := __obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	if reflect2.IsNil(__obf_02ba706f4f104d71) {
		__obf_95093efb9321684e := __obf_15da62f311934a45.(*reflect2.UnsafePtrType)
		__obf_8b28c1d23e9d3881 := __obf_95093efb9321684e.Elem()
		__obf_146c3283557f8c5c := __obf_8b28c1d23e9d3881.UnsafeNew()
		__obf_95093efb9321684e.
			UnsafeSet(__obf_47fa53f3d161f45c, unsafe.Pointer(&__obf_146c3283557f8c5c))
		__obf_02ba706f4f104d71 = __obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	}
	__obf_8e93113c2a45b2c0 := (__obf_02ba706f4f104d71).(encoding.TextUnmarshaler)
	__obf_44b48c26051f8033 := __obf_47edab4c16a2d88d.ReadString()
	__obf_e56742d6e52f642e := __obf_8e93113c2a45b2c0.UnmarshalText([]byte(__obf_44b48c26051f8033))
	if __obf_e56742d6e52f642e != nil {
		__obf_47edab4c16a2d88d.
			ReportError("textUnmarshalerDecoder", __obf_e56742d6e52f642e.Error())
	}
}
