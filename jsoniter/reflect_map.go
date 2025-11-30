package __obf_703d23ba520c3295

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_ef246e5fc788f28b(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_669b914f4eb06d1d := __obf_3b7f6abbae19451e.(*reflect2.UnsafeMapType)
	__obf_5e74eb253110edcf := __obf_76df447803cf1211(__obf_2aaf7367de3ff86d.append("[mapKey]"), __obf_669b914f4eb06d1d.Key())
	__obf_0db873ef26515049 := __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d.append("[mapElem]"), __obf_669b914f4eb06d1d.Elem())
	return &__obf_302717eb88d89600{__obf_669b914f4eb06d1d: __obf_669b914f4eb06d1d, __obf_fa79d56c907a638d: __obf_669b914f4eb06d1d.Key(), __obf_8b28c1d23e9d3881: __obf_669b914f4eb06d1d.Elem(), __obf_5e74eb253110edcf: __obf_5e74eb253110edcf, __obf_0db873ef26515049: __obf_0db873ef26515049}
}

func __obf_e70914e98bbb608c(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_669b914f4eb06d1d := __obf_3b7f6abbae19451e.(*reflect2.UnsafeMapType)
	if __obf_2aaf7367de3ff86d.__obf_e5ee8066c2c3cabf {
		return &__obf_b9c48409b0fbb821{__obf_669b914f4eb06d1d: __obf_669b914f4eb06d1d, __obf_d0b1b5a3176525c6: __obf_e663d3cce0cbb7bc(__obf_2aaf7367de3ff86d.append("[mapKey]"), __obf_669b914f4eb06d1d.Key()), __obf_8bb551383a0ba60c: __obf_906496c658b349cf(__obf_2aaf7367de3ff86d.append("[mapElem]"), __obf_669b914f4eb06d1d.Elem())}
	}
	return &__obf_3a14c5ac4effa256{__obf_669b914f4eb06d1d: __obf_669b914f4eb06d1d, __obf_d0b1b5a3176525c6: __obf_e663d3cce0cbb7bc(__obf_2aaf7367de3ff86d.append("[mapKey]"), __obf_669b914f4eb06d1d.Key()), __obf_8bb551383a0ba60c: __obf_906496c658b349cf(__obf_2aaf7367de3ff86d.append("[mapElem]"), __obf_669b914f4eb06d1d.Elem())}
}

func __obf_76df447803cf1211(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	__obf_c9719e68325f7d44 := __obf_2aaf7367de3ff86d.__obf_7c9972bb545a908d.CreateMapKeyDecoder(__obf_3b7f6abbae19451e)
	if __obf_c9719e68325f7d44 != nil {
		return __obf_c9719e68325f7d44
	}
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_c9719e68325f7d44 := __obf_dfc6ab1ee0bfd79e.CreateMapKeyDecoder(__obf_3b7f6abbae19451e)
		if __obf_c9719e68325f7d44 != nil {
			return __obf_c9719e68325f7d44
		}
	}
	__obf_95093efb9321684e := reflect2.PtrTo(__obf_3b7f6abbae19451e)
	if __obf_95093efb9321684e.Implements(__obf_5be0c81d693cd68b) {
		return &__obf_8eed859ba47bd05c{
			&__obf_71c95f6b15d8fce6{__obf_15da62f311934a45: __obf_95093efb9321684e},
		}
	}
	if __obf_3b7f6abbae19451e.Implements(__obf_5be0c81d693cd68b) {
		return &__obf_71c95f6b15d8fce6{__obf_15da62f311934a45: __obf_3b7f6abbae19451e}
	}
	if __obf_95093efb9321684e.Implements(__obf_9b7bc9dc1ee32b52) {
		return &__obf_8eed859ba47bd05c{
			&__obf_87fa515d0cb34be8{__obf_15da62f311934a45: __obf_95093efb9321684e},
		}
	}
	if __obf_3b7f6abbae19451e.Implements(__obf_9b7bc9dc1ee32b52) {
		return &__obf_87fa515d0cb34be8{__obf_15da62f311934a45: __obf_3b7f6abbae19451e}
	}

	switch __obf_3b7f6abbae19451e.Kind() {
	case reflect.String:
		return __obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_3b7f6abbae19451e = reflect2.DefaultTypeOfKind(__obf_3b7f6abbae19451e.Kind())
		return &__obf_b08cc5e6e2acbb91{__obf_b27fe2bc17d94621(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)}
	default:
		return &__obf_63b2b5a85529d72d{__obf_e56742d6e52f642e: fmt.Errorf("unsupported map key type: %v", __obf_3b7f6abbae19451e)}
	}
}

func __obf_e663d3cce0cbb7bc(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	__obf_cf840243a12ee308 := __obf_2aaf7367de3ff86d.__obf_9289f2028bb085f2.CreateMapKeyEncoder(__obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	for _, __obf_dfc6ab1ee0bfd79e := range __obf_2aaf7367de3ff86d.__obf_324ee50b8c0b2f3b {
		__obf_cf840243a12ee308 := __obf_dfc6ab1ee0bfd79e.CreateMapKeyEncoder(__obf_3b7f6abbae19451e)
		if __obf_cf840243a12ee308 != nil {
			return __obf_cf840243a12ee308
		}
	}

	if __obf_3b7f6abbae19451e.Kind() != reflect.String {
		if __obf_3b7f6abbae19451e == __obf_04d34c6784ab4d74 {
			return &__obf_0a029e0e5aaa3fdf{__obf_cf5770f0ffaef5fe: __obf_2aaf7367de3ff86d.EncoderOf(reflect2.TypeOf(""))}
		}
		if __obf_3b7f6abbae19451e.Implements(__obf_04d34c6784ab4d74) {
			return &__obf_43aa2e0944cf758f{__obf_15da62f311934a45: __obf_3b7f6abbae19451e, __obf_cf5770f0ffaef5fe: __obf_2aaf7367de3ff86d.EncoderOf(reflect2.TypeOf(""))}
		}
	}

	switch __obf_3b7f6abbae19451e.Kind() {
	case reflect.String:
		return __obf_906496c658b349cf(__obf_2aaf7367de3ff86d, reflect2.DefaultTypeOfKind(reflect.String))
	case reflect.Bool,
		reflect.Uint8, reflect.Int8,
		reflect.Uint16, reflect.Int16,
		reflect.Uint32, reflect.Int32,
		reflect.Uint64, reflect.Int64,
		reflect.Uint, reflect.Int,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr:
		__obf_3b7f6abbae19451e = reflect2.DefaultTypeOfKind(__obf_3b7f6abbae19451e.Kind())
		return &__obf_83ff42a6a7e96b5d{__obf_906496c658b349cf(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)}
	default:
		if __obf_3b7f6abbae19451e.Kind() == reflect.Interface {
			return &__obf_f7781b13dcd43ea4{__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e}
		}
		return &__obf_f80a54e26786ebf2{__obf_e56742d6e52f642e: fmt.Errorf("unsupported map key type: %v", __obf_3b7f6abbae19451e)}
	}
}

type __obf_302717eb88d89600 struct {
	__obf_669b914f4eb06d1d *reflect2.UnsafeMapType
	__obf_fa79d56c907a638d reflect2.Type
	__obf_8b28c1d23e9d3881 reflect2.Type
	__obf_5e74eb253110edcf ValDecoder
	__obf_0db873ef26515049 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_302717eb88d89600) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_669b914f4eb06d1d := __obf_c9719e68325f7d44.__obf_669b914f4eb06d1d
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == 'n' {
		__obf_47edab4c16a2d88d.__obf_0704d4be12be5e96('u', 'l', 'l')
		*(*unsafe.Pointer)(__obf_47fa53f3d161f45c) = nil
		__obf_669b914f4eb06d1d.
			UnsafeSet(__obf_47fa53f3d161f45c, __obf_669b914f4eb06d1d.UnsafeNew())
		return
	}
	if __obf_669b914f4eb06d1d.UnsafeIsNil(__obf_47fa53f3d161f45c) {
		__obf_669b914f4eb06d1d.
			UnsafeSet(__obf_47fa53f3d161f45c, __obf_669b914f4eb06d1d.UnsafeMakeMap(0))
	}
	if __obf_bd08f5161fff294a != '{' {
		__obf_47edab4c16a2d88d.
			ReportError("ReadMapCB", `expect { or n, but found `+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
	__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a == '}' {
		return
	}
	__obf_47edab4c16a2d88d.__obf_743b1a47c346c8a3()
	__obf_4580e6a9e5b1ee92 := __obf_c9719e68325f7d44.__obf_fa79d56c907a638d.UnsafeNew()
	__obf_c9719e68325f7d44.__obf_5e74eb253110edcf.
		Decode(__obf_4580e6a9e5b1ee92, __obf_47edab4c16a2d88d)
	__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a != ':' {
		__obf_47edab4c16a2d88d.
			ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
	__obf_146c3283557f8c5c := __obf_c9719e68325f7d44.__obf_8b28c1d23e9d3881.UnsafeNew()
	__obf_c9719e68325f7d44.__obf_0db873ef26515049.
		Decode(__obf_146c3283557f8c5c, __obf_47edab4c16a2d88d)
	__obf_c9719e68325f7d44.__obf_669b914f4eb06d1d.
		UnsafeSetIndex(__obf_47fa53f3d161f45c, __obf_4580e6a9e5b1ee92, __obf_146c3283557f8c5c)
	for __obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec(); __obf_bd08f5161fff294a == ','; __obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec() {
		__obf_4580e6a9e5b1ee92 := __obf_c9719e68325f7d44.__obf_fa79d56c907a638d.UnsafeNew()
		__obf_c9719e68325f7d44.__obf_5e74eb253110edcf.
			Decode(__obf_4580e6a9e5b1ee92, __obf_47edab4c16a2d88d)
		__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
		if __obf_bd08f5161fff294a != ':' {
			__obf_47edab4c16a2d88d.
				ReportError("ReadMapCB", "expect : after object field, but found "+string([]byte{__obf_bd08f5161fff294a}))
			return
		}
		__obf_146c3283557f8c5c := __obf_c9719e68325f7d44.__obf_8b28c1d23e9d3881.UnsafeNew()
		__obf_c9719e68325f7d44.__obf_0db873ef26515049.
			Decode(__obf_146c3283557f8c5c, __obf_47edab4c16a2d88d)
		__obf_c9719e68325f7d44.__obf_669b914f4eb06d1d.
			UnsafeSetIndex(__obf_47fa53f3d161f45c, __obf_4580e6a9e5b1ee92, __obf_146c3283557f8c5c)
	}
	if __obf_bd08f5161fff294a != '}' {
		__obf_47edab4c16a2d88d.
			ReportError("ReadMapCB", `expect }, but found `+string([]byte{__obf_bd08f5161fff294a}))
	}
}

type __obf_b08cc5e6e2acbb91 struct {
	__obf_c9719e68325f7d44 ValDecoder
}

func (__obf_c9719e68325f7d44 *__obf_b08cc5e6e2acbb91) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	__obf_bd08f5161fff294a := __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a != '"' {
		__obf_47edab4c16a2d88d.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
	__obf_c9719e68325f7d44.__obf_c9719e68325f7d44.
		Decode(__obf_47fa53f3d161f45c, __obf_47edab4c16a2d88d)
	__obf_bd08f5161fff294a = __obf_47edab4c16a2d88d.__obf_6f7a9d6f8e846eec()
	if __obf_bd08f5161fff294a != '"' {
		__obf_47edab4c16a2d88d.
			ReportError("ReadMapCB", `expect ", but found `+string([]byte{__obf_bd08f5161fff294a}))
		return
	}
}

type __obf_83ff42a6a7e96b5d struct {
	__obf_cf840243a12ee308 ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_83ff42a6a7e96b5d) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
	__obf_cf840243a12ee308.__obf_cf840243a12ee308.
		Encode(__obf_47fa53f3d161f45c, __obf_9a34f62857fb1d1d)
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
}

func (__obf_cf840243a12ee308 *__obf_83ff42a6a7e96b5d) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return false
}

type __obf_f7781b13dcd43ea4 struct {
	__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d
	__obf_15da62f311934a45 reflect2.Type
}

func (__obf_cf840243a12ee308 *__obf_f7781b13dcd43ea4) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_02ba706f4f104d71 := __obf_cf840243a12ee308.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	__obf_e663d3cce0cbb7bc(__obf_cf840243a12ee308.__obf_2aaf7367de3ff86d, reflect2.TypeOf(__obf_02ba706f4f104d71)).Encode(reflect2.PtrOf(__obf_02ba706f4f104d71), __obf_9a34f62857fb1d1d)
}

func (__obf_cf840243a12ee308 *__obf_f7781b13dcd43ea4) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_02ba706f4f104d71 := __obf_cf840243a12ee308.__obf_15da62f311934a45.UnsafeIndirect(__obf_47fa53f3d161f45c)
	return __obf_e663d3cce0cbb7bc(__obf_cf840243a12ee308.__obf_2aaf7367de3ff86d, reflect2.TypeOf(__obf_02ba706f4f104d71)).IsEmpty(reflect2.PtrOf(__obf_02ba706f4f104d71))
}

type __obf_3a14c5ac4effa256 struct {
	__obf_669b914f4eb06d1d *reflect2.UnsafeMapType
	__obf_d0b1b5a3176525c6 ValEncoder
	__obf_8bb551383a0ba60c ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_3a14c5ac4effa256) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if *(*unsafe.Pointer)(__obf_47fa53f3d161f45c) == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_9a34f62857fb1d1d.
		WriteObjectStart()
	__obf_47edab4c16a2d88d := __obf_cf840243a12ee308.__obf_669b914f4eb06d1d.UnsafeIterate(__obf_47fa53f3d161f45c)
	for __obf_b0a5d2bd48690f1d := 0; __obf_47edab4c16a2d88d.HasNext(); __obf_b0a5d2bd48690f1d++ {
		if __obf_b0a5d2bd48690f1d != 0 {
			__obf_9a34f62857fb1d1d.
				WriteMore()
		}
		__obf_4580e6a9e5b1ee92, __obf_146c3283557f8c5c := __obf_47edab4c16a2d88d.UnsafeNext()
		__obf_cf840243a12ee308.__obf_d0b1b5a3176525c6.
			Encode(__obf_4580e6a9e5b1ee92, __obf_9a34f62857fb1d1d)
		if __obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c > 0 {
			__obf_9a34f62857fb1d1d.__obf_e3076a8aafc9c7f0(byte(':'), byte(' '))
		} else {
			__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19(':')
		}
		__obf_cf840243a12ee308.__obf_8bb551383a0ba60c.
			Encode(__obf_146c3283557f8c5c, __obf_9a34f62857fb1d1d)
	}
	__obf_9a34f62857fb1d1d.
		WriteObjectEnd()
}

func (__obf_cf840243a12ee308 *__obf_3a14c5ac4effa256) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_47edab4c16a2d88d := __obf_cf840243a12ee308.__obf_669b914f4eb06d1d.UnsafeIterate(__obf_47fa53f3d161f45c)
	return !__obf_47edab4c16a2d88d.HasNext()
}

type __obf_b9c48409b0fbb821 struct {
	__obf_669b914f4eb06d1d *reflect2.UnsafeMapType
	__obf_d0b1b5a3176525c6 ValEncoder
	__obf_8bb551383a0ba60c ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_b9c48409b0fbb821) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if *(*unsafe.Pointer)(__obf_47fa53f3d161f45c) == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
		return
	}
	__obf_9a34f62857fb1d1d.
		WriteObjectStart()
	__obf_4c648aef3413de9b := __obf_cf840243a12ee308.__obf_669b914f4eb06d1d.UnsafeIterate(__obf_47fa53f3d161f45c)
	__obf_5efcbf34c1361e9f := __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.BorrowStream(nil)
	__obf_5efcbf34c1361e9f.
		Attachment = __obf_9a34f62857fb1d1d.Attachment
	__obf_100e87163e3c8c47 := __obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.BorrowIterator(nil)
	__obf_b379a983d84ca102 := __obf_5f88096e58b17b0d{}
	for __obf_4c648aef3413de9b.HasNext() {
		__obf_4580e6a9e5b1ee92, __obf_146c3283557f8c5c := __obf_4c648aef3413de9b.UnsafeNext()
		__obf_ee493db6f6dcf9e7 := __obf_5efcbf34c1361e9f.Buffered()
		__obf_cf840243a12ee308.__obf_d0b1b5a3176525c6.
			Encode(__obf_4580e6a9e5b1ee92, __obf_5efcbf34c1361e9f)
		if __obf_5efcbf34c1361e9f.Error != nil && __obf_5efcbf34c1361e9f.Error != io.EOF && __obf_9a34f62857fb1d1d.Error == nil {
			__obf_9a34f62857fb1d1d.
				Error = __obf_5efcbf34c1361e9f.Error
		}
		__obf_e6c2846a5577a06e := __obf_5efcbf34c1361e9f.Buffer()[__obf_ee493db6f6dcf9e7:]
		__obf_100e87163e3c8c47.
			ResetBytes(__obf_e6c2846a5577a06e)
		__obf_e4c2dc96ff229e95 := __obf_100e87163e3c8c47.ReadString()
		if __obf_9a34f62857fb1d1d.__obf_1dd0566c8f02217c > 0 {
			__obf_5efcbf34c1361e9f.__obf_e3076a8aafc9c7f0(byte(':'), byte(' '))
		} else {
			__obf_5efcbf34c1361e9f.__obf_f7df2a5224462d19(':')
		}
		__obf_cf840243a12ee308.__obf_8bb551383a0ba60c.
			Encode(__obf_146c3283557f8c5c, __obf_5efcbf34c1361e9f)
		__obf_b379a983d84ca102 = append(__obf_b379a983d84ca102, __obf_8da408df8ae090a0{__obf_4580e6a9e5b1ee92: __obf_e4c2dc96ff229e95, __obf_8fa8a42329d6cea8: __obf_5efcbf34c1361e9f.Buffer()[__obf_ee493db6f6dcf9e7:]})
	}
	sort.Sort(__obf_b379a983d84ca102)
	for __obf_b0a5d2bd48690f1d, __obf_8fa8a42329d6cea8 := range __obf_b379a983d84ca102 {
		if __obf_b0a5d2bd48690f1d != 0 {
			__obf_9a34f62857fb1d1d.
				WriteMore()
		}
		__obf_9a34f62857fb1d1d.
			Write(__obf_8fa8a42329d6cea8.__obf_8fa8a42329d6cea8)
	}
	if __obf_5efcbf34c1361e9f.Error != nil && __obf_9a34f62857fb1d1d.Error == nil {
		__obf_9a34f62857fb1d1d.
			Error = __obf_5efcbf34c1361e9f.Error
	}
	__obf_9a34f62857fb1d1d.
		WriteObjectEnd()
	__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.
		ReturnStream(__obf_5efcbf34c1361e9f)
	__obf_9a34f62857fb1d1d.__obf_25bd4f754a37b862.
		ReturnIterator(__obf_100e87163e3c8c47)
}

func (__obf_cf840243a12ee308 *__obf_b9c48409b0fbb821) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_47edab4c16a2d88d := __obf_cf840243a12ee308.__obf_669b914f4eb06d1d.UnsafeIterate(__obf_47fa53f3d161f45c)
	return !__obf_47edab4c16a2d88d.HasNext()
}

type __obf_5f88096e58b17b0d []__obf_8da408df8ae090a0

type __obf_8da408df8ae090a0 struct {
	__obf_4580e6a9e5b1ee92 string
	__obf_8fa8a42329d6cea8 []byte
}

func (__obf_ab0c6efb6a61e30a __obf_5f88096e58b17b0d) Len() int { return len(__obf_ab0c6efb6a61e30a) }
func (__obf_ab0c6efb6a61e30a __obf_5f88096e58b17b0d) Swap(__obf_b0a5d2bd48690f1d, __obf_0e04c22ffdf339de int) {
	__obf_ab0c6efb6a61e30a[__obf_b0a5d2bd48690f1d], __obf_ab0c6efb6a61e30a[__obf_0e04c22ffdf339de] = __obf_ab0c6efb6a61e30a[__obf_0e04c22ffdf339de], __obf_ab0c6efb6a61e30a[__obf_b0a5d2bd48690f1d]
}
func (__obf_ab0c6efb6a61e30a __obf_5f88096e58b17b0d) Less(__obf_b0a5d2bd48690f1d, __obf_0e04c22ffdf339de int) bool {
	return __obf_ab0c6efb6a61e30a[__obf_b0a5d2bd48690f1d].__obf_4580e6a9e5b1ee92 < __obf_ab0c6efb6a61e30a[__obf_0e04c22ffdf339de].__obf_4580e6a9e5b1ee92
}
