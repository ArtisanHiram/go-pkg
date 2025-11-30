package __obf_030d39f7a8de96c6

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"reflect"
	"unsafe"
)

func __obf_9070536cbef1ec38(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) ValEncoder {
	type __obf_8e5e2942d0e6a6e1 struct {
		__obf_04ea35bce3978e8f *Binding
		__obf_15e61366046a3c44 string
		__obf_9d83fb66a13b8998 bool
	}
	__obf_bab8bf2dc6310235 := []*__obf_8e5e2942d0e6a6e1{}
	__obf_ef1b847f404bee13 := __obf_1219e3aba7e47ed0(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	for _, __obf_04ea35bce3978e8f := range __obf_ef1b847f404bee13.Fields {
		for _, __obf_15e61366046a3c44 := range __obf_04ea35bce3978e8f.ToNames {
			new := &__obf_8e5e2942d0e6a6e1{__obf_04ea35bce3978e8f: __obf_04ea35bce3978e8f, __obf_15e61366046a3c44: __obf_15e61366046a3c44}
			for _, __obf_09fa488c9c17c6a9 := range __obf_bab8bf2dc6310235 {
				if __obf_09fa488c9c17c6a9.__obf_15e61366046a3c44 != __obf_15e61366046a3c44 {
					continue
				}
				__obf_09fa488c9c17c6a9.__obf_9d83fb66a13b8998, new.__obf_9d83fb66a13b8998 = __obf_accc8def568f71d8(__obf_a565fbaffca944f0.__obf_81639538813814ff, __obf_09fa488c9c17c6a9.__obf_04ea35bce3978e8f, new.__obf_04ea35bce3978e8f)
			}
			__obf_bab8bf2dc6310235 = append(__obf_bab8bf2dc6310235, new)
		}
	}
	if len(__obf_bab8bf2dc6310235) == 0 {
		return &__obf_04007c530d356eb8{}
	}
	__obf_535ba859b2a93314 := []__obf_e1683c18ed80c005{}
	for _, __obf_8e5e2942d0e6a6e1 := range __obf_bab8bf2dc6310235 {
		if !__obf_8e5e2942d0e6a6e1.__obf_9d83fb66a13b8998 {
			__obf_535ba859b2a93314 = append(__obf_535ba859b2a93314, __obf_e1683c18ed80c005{__obf_008f61596d7e9523: __obf_8e5e2942d0e6a6e1.__obf_04ea35bce3978e8f.Encoder.(*__obf_06e12c234657ff27), __obf_15e61366046a3c44: __obf_8e5e2942d0e6a6e1.__obf_15e61366046a3c44})
		}
	}
	return &__obf_60b76762a9b33529{__obf_8744d0b8c80ccdc1, __obf_535ba859b2a93314}
}

func __obf_adf48ba0507caf28(__obf_a565fbaffca944f0 *__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1 reflect2.Type) __obf_b5475e6763577f46 {
	__obf_008f61596d7e9523 := __obf_438eebf6331ef15b(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	if __obf_008f61596d7e9523 != nil {
		return __obf_008f61596d7e9523
	}
	__obf_30e494a4f2832c2f := __obf_8744d0b8c80ccdc1.Kind()
	switch __obf_30e494a4f2832c2f {
	case reflect.Interface:
		return &__obf_3691828d7304f473{__obf_8744d0b8c80ccdc1}
	case reflect.Struct:
		return &__obf_60b76762a9b33529{__obf_8744d0b8c80ccdc1: __obf_8744d0b8c80ccdc1}
	case reflect.Array:
		return &__obf_31e9b5ebcd4b6405{}
	case reflect.Slice:
		return &__obf_a69908ce40c0d051{}
	case reflect.Map:
		return __obf_5cc5784cdf0ac33a(__obf_a565fbaffca944f0, __obf_8744d0b8c80ccdc1)
	case reflect.Ptr:
		return &OptionalEncoder{}
	default:
		return &__obf_f14cb8ddffae87b3{__obf_fcc907dd69879566: fmt.Errorf("unsupported type: %v", __obf_8744d0b8c80ccdc1)}
	}
}

func __obf_accc8def568f71d8(__obf_679611dc56ff465b *__obf_81639538813814ff, __obf_09fa488c9c17c6a9, new *Binding) (__obf_4229457075af2115, __obf_78181e83750ef80b bool) {
	__obf_78531b87e53b141f := new.Field.Tag().Get(__obf_679611dc56ff465b.__obf_79fd0a7f8863c590()) != ""
	__obf_3f7cd25e650ea074 := __obf_09fa488c9c17c6a9.Field.Tag().Get(__obf_679611dc56ff465b.__obf_79fd0a7f8863c590()) != ""
	if __obf_78531b87e53b141f {
		if __obf_3f7cd25e650ea074 {
			if len(__obf_09fa488c9c17c6a9.__obf_9cdf9b6aa25d4f68) > len(new.__obf_9cdf9b6aa25d4f68) {
				return true, false
			} else if len(new.__obf_9cdf9b6aa25d4f68) > len(__obf_09fa488c9c17c6a9.__obf_9cdf9b6aa25d4f68) {
				return false, true
			} else {
				return true, true
			}
		} else {
			return true, false
		}
	} else {
		if __obf_3f7cd25e650ea074 {
			return true, false
		}
		if len(__obf_09fa488c9c17c6a9.__obf_9cdf9b6aa25d4f68) > len(new.__obf_9cdf9b6aa25d4f68) {
			return true, false
		} else if len(new.__obf_9cdf9b6aa25d4f68) > len(__obf_09fa488c9c17c6a9.__obf_9cdf9b6aa25d4f68) {
			return false, true
		} else {
			return true, true
		}
	}
}

type __obf_06e12c234657ff27 struct {
	__obf_cd4d02f264b18d55 reflect2.StructField
	__obf_b7243a4c5b39bcf9 ValEncoder
	__obf_ef6aa90a5fa198f2 bool
}

func (__obf_008f61596d7e9523 *__obf_06e12c234657ff27) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_ba13c65f6e3a796c := __obf_008f61596d7e9523.__obf_cd4d02f264b18d55.UnsafeGet(__obf_dbbf371b91136494)
	__obf_008f61596d7e9523.__obf_b7243a4c5b39bcf9.
		Encode(__obf_ba13c65f6e3a796c, __obf_8a2c51fe22775655)
	if __obf_8a2c51fe22775655.Error != nil && __obf_8a2c51fe22775655.Error != io.EOF {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("%s: %s", __obf_008f61596d7e9523.__obf_cd4d02f264b18d55.Name(), __obf_8a2c51fe22775655.Error.Error())
	}
}

func (__obf_008f61596d7e9523 *__obf_06e12c234657ff27) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_ba13c65f6e3a796c := __obf_008f61596d7e9523.__obf_cd4d02f264b18d55.UnsafeGet(__obf_dbbf371b91136494)
	return __obf_008f61596d7e9523.__obf_b7243a4c5b39bcf9.IsEmpty(__obf_ba13c65f6e3a796c)
}

func (__obf_008f61596d7e9523 *__obf_06e12c234657ff27) IsEmbeddedPtrNil(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	__obf_f22b71bcff862810, __obf_ae94b7511a43a3e7 := __obf_008f61596d7e9523.__obf_b7243a4c5b39bcf9.(IsEmbeddedPtrNil)
	if !__obf_ae94b7511a43a3e7 {
		return false
	}
	__obf_ba13c65f6e3a796c := __obf_008f61596d7e9523.__obf_cd4d02f264b18d55.UnsafeGet(__obf_dbbf371b91136494)
	return __obf_f22b71bcff862810.IsEmbeddedPtrNil(__obf_ba13c65f6e3a796c)
}

type IsEmbeddedPtrNil interface {
	IsEmbeddedPtrNil(__obf_dbbf371b91136494 unsafe.Pointer) bool
}

type __obf_60b76762a9b33529 struct {
	__obf_8744d0b8c80ccdc1 reflect2.Type
	__obf_dadaf38c30ea4eba []__obf_e1683c18ed80c005
}

type __obf_e1683c18ed80c005 struct {
	__obf_008f61596d7e9523 *__obf_06e12c234657ff27
	__obf_15e61366046a3c44 string
}

func (__obf_008f61596d7e9523 *__obf_60b76762a9b33529) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteObjectStart()
	__obf_be4a6ebd4fcf4cbf := false
	for _, __obf_cd4d02f264b18d55 := range __obf_008f61596d7e9523.__obf_dadaf38c30ea4eba {
		if __obf_cd4d02f264b18d55.__obf_008f61596d7e9523.__obf_ef6aa90a5fa198f2 && __obf_cd4d02f264b18d55.__obf_008f61596d7e9523.IsEmpty(__obf_dbbf371b91136494) {
			continue
		}
		if __obf_cd4d02f264b18d55.__obf_008f61596d7e9523.IsEmbeddedPtrNil(__obf_dbbf371b91136494) {
			continue
		}
		if __obf_be4a6ebd4fcf4cbf {
			__obf_8a2c51fe22775655.
				WriteMore()
		}
		__obf_8a2c51fe22775655.
			WriteObjectField(__obf_cd4d02f264b18d55.__obf_15e61366046a3c44)
		__obf_cd4d02f264b18d55.__obf_008f61596d7e9523.
			Encode(__obf_dbbf371b91136494, __obf_8a2c51fe22775655)
		__obf_be4a6ebd4fcf4cbf = true
	}
	__obf_8a2c51fe22775655.
		WriteObjectEnd()
	if __obf_8a2c51fe22775655.Error != nil && __obf_8a2c51fe22775655.Error != io.EOF {
		__obf_8a2c51fe22775655.
			Error = fmt.Errorf("%v.%s", __obf_008f61596d7e9523.__obf_8744d0b8c80ccdc1, __obf_8a2c51fe22775655.Error.Error())
	}
}

func (__obf_008f61596d7e9523 *__obf_60b76762a9b33529) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return false
}

type __obf_04007c530d356eb8 struct {
}

func (__obf_008f61596d7e9523 *__obf_04007c530d356eb8) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.
		WriteEmptyObject()
}

func (__obf_008f61596d7e9523 *__obf_04007c530d356eb8) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return false
}

type __obf_e8dab09dfce32843 struct {
	__obf_73c590fa207e5962 ValEncoder
}

func (__obf_008f61596d7e9523 *__obf_e8dab09dfce32843) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
	__obf_008f61596d7e9523.__obf_73c590fa207e5962.
		Encode(__obf_dbbf371b91136494, __obf_8a2c51fe22775655)
	__obf_8a2c51fe22775655.__obf_41130daa346c0e53('"')
}

func (__obf_008f61596d7e9523 *__obf_e8dab09dfce32843) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_73c590fa207e5962.IsEmpty(__obf_dbbf371b91136494)
}

type __obf_d418a6a275dc78e1 struct {
	__obf_73c590fa207e5962 ValEncoder
	__obf_679611dc56ff465b *__obf_81639538813814ff
}

func (__obf_008f61596d7e9523 *__obf_d418a6a275dc78e1) Encode(__obf_dbbf371b91136494 unsafe.Pointer, __obf_8a2c51fe22775655 *Stream) {
	__obf_46057b7c36e00232 := __obf_008f61596d7e9523.__obf_679611dc56ff465b.BorrowStream(nil)
	__obf_46057b7c36e00232.
		Attachment = __obf_8a2c51fe22775655.Attachment
	defer __obf_008f61596d7e9523.__obf_679611dc56ff465b.ReturnStream(__obf_46057b7c36e00232)
	__obf_008f61596d7e9523.__obf_73c590fa207e5962.
		Encode(__obf_dbbf371b91136494, __obf_46057b7c36e00232)
	__obf_8a2c51fe22775655.
		WriteString(string(__obf_46057b7c36e00232.Buffer()))
}

func (__obf_008f61596d7e9523 *__obf_d418a6a275dc78e1) IsEmpty(__obf_dbbf371b91136494 unsafe.Pointer) bool {
	return __obf_008f61596d7e9523.__obf_73c590fa207e5962.IsEmpty(__obf_dbbf371b91136494)
}
