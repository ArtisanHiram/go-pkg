package __obf_703d23ba520c3295

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"reflect"
	"unsafe"
)

func __obf_76e40d1501c65ef3(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	type __obf_7996e50e836419c3 struct {
		__obf_af9543e035a21def *Binding
		__obf_999dd3fc29c90048 string
		__obf_b5e6105aa17e42d8 bool
	}
	__obf_2dc4f25cfd406615 := []*__obf_7996e50e836419c3{}
	__obf_ab1ac896d4aa86ee := __obf_969b50c34ac12a33(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	for _, __obf_af9543e035a21def := range __obf_ab1ac896d4aa86ee.Fields {
		for _, __obf_999dd3fc29c90048 := range __obf_af9543e035a21def.ToNames {
			new := &__obf_7996e50e836419c3{__obf_af9543e035a21def: __obf_af9543e035a21def, __obf_999dd3fc29c90048: __obf_999dd3fc29c90048}
			for _, __obf_3dafc9b4a2f87f34 := range __obf_2dc4f25cfd406615 {
				if __obf_3dafc9b4a2f87f34.__obf_999dd3fc29c90048 != __obf_999dd3fc29c90048 {
					continue
				}
				__obf_3dafc9b4a2f87f34.__obf_b5e6105aa17e42d8, new.__obf_b5e6105aa17e42d8 = __obf_9de3265131b75a41(__obf_2aaf7367de3ff86d.__obf_ef74248d8ae9ba78, __obf_3dafc9b4a2f87f34.__obf_af9543e035a21def, new.__obf_af9543e035a21def)
			}
			__obf_2dc4f25cfd406615 = append(__obf_2dc4f25cfd406615, new)
		}
	}
	if len(__obf_2dc4f25cfd406615) == 0 {
		return &__obf_5b20ed0145344059{}
	}
	__obf_d0d5f7c76a24a313 := []__obf_9a884644ecf3529c{}
	for _, __obf_7996e50e836419c3 := range __obf_2dc4f25cfd406615 {
		if !__obf_7996e50e836419c3.__obf_b5e6105aa17e42d8 {
			__obf_d0d5f7c76a24a313 = append(__obf_d0d5f7c76a24a313, __obf_9a884644ecf3529c{__obf_cf840243a12ee308: __obf_7996e50e836419c3.__obf_af9543e035a21def.Encoder.(*__obf_2dd807eec6f01cb2), __obf_999dd3fc29c90048: __obf_7996e50e836419c3.__obf_999dd3fc29c90048})
		}
	}
	return &__obf_4bc81ce2168a7e53{__obf_3b7f6abbae19451e, __obf_d0d5f7c76a24a313}
}

func __obf_2ae551982b66bbee(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) __obf_8daa57d6b3a720c3 {
	__obf_cf840243a12ee308 := __obf_3774fbf920f93937(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	if __obf_cf840243a12ee308 != nil {
		return __obf_cf840243a12ee308
	}
	__obf_2bb4ea546eb4165c := __obf_3b7f6abbae19451e.Kind()
	switch __obf_2bb4ea546eb4165c {
	case reflect.Interface:
		return &__obf_7f31cdbc4d682d50{__obf_3b7f6abbae19451e}
	case reflect.Struct:
		return &__obf_4bc81ce2168a7e53{__obf_3b7f6abbae19451e: __obf_3b7f6abbae19451e}
	case reflect.Array:
		return &__obf_dfdeb193f963ce62{}
	case reflect.Slice:
		return &__obf_d936159a056aa5e5{}
	case reflect.Map:
		return __obf_e70914e98bbb608c(__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e)
	case reflect.Ptr:
		return &OptionalEncoder{}
	default:
		return &__obf_f80a54e26786ebf2{__obf_e56742d6e52f642e: fmt.Errorf("unsupported type: %v", __obf_3b7f6abbae19451e)}
	}
}

func __obf_9de3265131b75a41(__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78, __obf_3dafc9b4a2f87f34, new *Binding) (__obf_a63abce1e9e94c66, __obf_8f285a2701f37644 bool) {
	__obf_f8e4a96c2a16edce := new.Field.Tag().Get(__obf_25bd4f754a37b862.__obf_75bbb1cf9d1d76ab()) != ""
	__obf_913343f5ac2d066c := __obf_3dafc9b4a2f87f34.Field.Tag().Get(__obf_25bd4f754a37b862.__obf_75bbb1cf9d1d76ab()) != ""
	if __obf_f8e4a96c2a16edce {
		if __obf_913343f5ac2d066c {
			if len(__obf_3dafc9b4a2f87f34.__obf_6c2a2b1e7bd4d99b) > len(new.__obf_6c2a2b1e7bd4d99b) {
				return true, false
			} else if len(new.__obf_6c2a2b1e7bd4d99b) > len(__obf_3dafc9b4a2f87f34.__obf_6c2a2b1e7bd4d99b) {
				return false, true
			} else {
				return true, true
			}
		} else {
			return true, false
		}
	} else {
		if __obf_913343f5ac2d066c {
			return true, false
		}
		if len(__obf_3dafc9b4a2f87f34.__obf_6c2a2b1e7bd4d99b) > len(new.__obf_6c2a2b1e7bd4d99b) {
			return true, false
		} else if len(new.__obf_6c2a2b1e7bd4d99b) > len(__obf_3dafc9b4a2f87f34.__obf_6c2a2b1e7bd4d99b) {
			return false, true
		} else {
			return true, true
		}
	}
}

type __obf_2dd807eec6f01cb2 struct {
	__obf_13f6440419533990 reflect2.StructField
	__obf_e639aac8037c7c5c ValEncoder
	__obf_6744c81d52458bde bool
}

func (__obf_cf840243a12ee308 *__obf_2dd807eec6f01cb2) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_94be1fdf2ec0d598 := __obf_cf840243a12ee308.__obf_13f6440419533990.UnsafeGet(__obf_47fa53f3d161f45c)
	__obf_cf840243a12ee308.__obf_e639aac8037c7c5c.
		Encode(__obf_94be1fdf2ec0d598, __obf_9a34f62857fb1d1d)
	if __obf_9a34f62857fb1d1d.Error != nil && __obf_9a34f62857fb1d1d.Error != io.EOF {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("%s: %s", __obf_cf840243a12ee308.__obf_13f6440419533990.Name(), __obf_9a34f62857fb1d1d.Error.Error())
	}
}

func (__obf_cf840243a12ee308 *__obf_2dd807eec6f01cb2) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_94be1fdf2ec0d598 := __obf_cf840243a12ee308.__obf_13f6440419533990.UnsafeGet(__obf_47fa53f3d161f45c)
	return __obf_cf840243a12ee308.__obf_e639aac8037c7c5c.IsEmpty(__obf_94be1fdf2ec0d598)
}

func (__obf_cf840243a12ee308 *__obf_2dd807eec6f01cb2) IsEmbeddedPtrNil(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	__obf_0619e0ea83732b1b, __obf_22ffff425cd0177f := __obf_cf840243a12ee308.__obf_e639aac8037c7c5c.(IsEmbeddedPtrNil)
	if !__obf_22ffff425cd0177f {
		return false
	}
	__obf_94be1fdf2ec0d598 := __obf_cf840243a12ee308.__obf_13f6440419533990.UnsafeGet(__obf_47fa53f3d161f45c)
	return __obf_0619e0ea83732b1b.IsEmbeddedPtrNil(__obf_94be1fdf2ec0d598)
}

type IsEmbeddedPtrNil interface {
	IsEmbeddedPtrNil(__obf_47fa53f3d161f45c unsafe.Pointer) bool
}

type __obf_4bc81ce2168a7e53 struct {
	__obf_3b7f6abbae19451e reflect2.Type
	__obf_f4d348cda1281225 []__obf_9a884644ecf3529c
}

type __obf_9a884644ecf3529c struct {
	__obf_cf840243a12ee308 *__obf_2dd807eec6f01cb2
	__obf_999dd3fc29c90048 string
}

func (__obf_cf840243a12ee308 *__obf_4bc81ce2168a7e53) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteObjectStart()
	__obf_43562a3e5d77283b := false
	for _, __obf_13f6440419533990 := range __obf_cf840243a12ee308.__obf_f4d348cda1281225 {
		if __obf_13f6440419533990.__obf_cf840243a12ee308.__obf_6744c81d52458bde && __obf_13f6440419533990.__obf_cf840243a12ee308.IsEmpty(__obf_47fa53f3d161f45c) {
			continue
		}
		if __obf_13f6440419533990.__obf_cf840243a12ee308.IsEmbeddedPtrNil(__obf_47fa53f3d161f45c) {
			continue
		}
		if __obf_43562a3e5d77283b {
			__obf_9a34f62857fb1d1d.
				WriteMore()
		}
		__obf_9a34f62857fb1d1d.
			WriteObjectField(__obf_13f6440419533990.__obf_999dd3fc29c90048)
		__obf_13f6440419533990.__obf_cf840243a12ee308.
			Encode(__obf_47fa53f3d161f45c, __obf_9a34f62857fb1d1d)
		__obf_43562a3e5d77283b = true
	}
	__obf_9a34f62857fb1d1d.
		WriteObjectEnd()
	if __obf_9a34f62857fb1d1d.Error != nil && __obf_9a34f62857fb1d1d.Error != io.EOF {
		__obf_9a34f62857fb1d1d.
			Error = fmt.Errorf("%v.%s", __obf_cf840243a12ee308.__obf_3b7f6abbae19451e, __obf_9a34f62857fb1d1d.Error.Error())
	}
}

func (__obf_cf840243a12ee308 *__obf_4bc81ce2168a7e53) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return false
}

type __obf_5b20ed0145344059 struct {
}

func (__obf_cf840243a12ee308 *__obf_5b20ed0145344059) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.
		WriteEmptyObject()
}

func (__obf_cf840243a12ee308 *__obf_5b20ed0145344059) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return false
}

type __obf_d9cbb1c64198bf84 struct {
	__obf_8bb551383a0ba60c ValEncoder
}

func (__obf_cf840243a12ee308 *__obf_d9cbb1c64198bf84) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
	__obf_cf840243a12ee308.__obf_8bb551383a0ba60c.
		Encode(__obf_47fa53f3d161f45c, __obf_9a34f62857fb1d1d)
	__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('"')
}

func (__obf_cf840243a12ee308 *__obf_d9cbb1c64198bf84) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_8bb551383a0ba60c.IsEmpty(__obf_47fa53f3d161f45c)
}

type __obf_9f3b818ecbe63ff5 struct {
	__obf_8bb551383a0ba60c ValEncoder
	__obf_25bd4f754a37b862 *__obf_ef74248d8ae9ba78
}

func (__obf_cf840243a12ee308 *__obf_9f3b818ecbe63ff5) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_1775bfee3e546f98 := __obf_cf840243a12ee308.__obf_25bd4f754a37b862.BorrowStream(nil)
	__obf_1775bfee3e546f98.
		Attachment = __obf_9a34f62857fb1d1d.Attachment
	defer __obf_cf840243a12ee308.__obf_25bd4f754a37b862.ReturnStream(__obf_1775bfee3e546f98)
	__obf_cf840243a12ee308.__obf_8bb551383a0ba60c.
		Encode(__obf_47fa53f3d161f45c, __obf_1775bfee3e546f98)
	__obf_9a34f62857fb1d1d.
		WriteString(string(__obf_1775bfee3e546f98.Buffer()))
}

func (__obf_cf840243a12ee308 *__obf_9f3b818ecbe63ff5) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return __obf_cf840243a12ee308.__obf_8bb551383a0ba60c.IsEmpty(__obf_47fa53f3d161f45c)
}
