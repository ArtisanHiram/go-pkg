package __obf_c7b79b12b33d8238

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"io"
	"reflect"
	"unsafe"
)

func __obf_5191b4476a0b2b2f(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) ValEncoder {
	type __obf_72c924a23efc3585 struct {
		__obf_56b28257c2a1cc8d *Binding
		__obf_a3d812fac643004b string
		__obf_b020b7b9f5f7dacc bool
	}
	__obf_44aea3b3531d146b := []*__obf_72c924a23efc3585{}
	__obf_31531899b0d36123 := __obf_0572835446c88a68(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	for _, __obf_56b28257c2a1cc8d := range __obf_31531899b0d36123.Fields {
		for _, __obf_a3d812fac643004b := range __obf_56b28257c2a1cc8d.ToNames {
			new := &__obf_72c924a23efc3585{__obf_56b28257c2a1cc8d: __obf_56b28257c2a1cc8d, __obf_a3d812fac643004b: __obf_a3d812fac643004b}
			for _, __obf_ffbea95e31df6856 := range __obf_44aea3b3531d146b {
				if __obf_ffbea95e31df6856.__obf_a3d812fac643004b != __obf_a3d812fac643004b {
					continue
				}
				__obf_ffbea95e31df6856.__obf_b020b7b9f5f7dacc, new.__obf_b020b7b9f5f7dacc = __obf_169b1277fbacdc3f(__obf_99ec45908bceabdb.__obf_30fe5c95cabd69c0, __obf_ffbea95e31df6856.__obf_56b28257c2a1cc8d, new.__obf_56b28257c2a1cc8d)
			}
			__obf_44aea3b3531d146b = append(__obf_44aea3b3531d146b, new)
		}
	}
	if len(__obf_44aea3b3531d146b) == 0 {
		return &__obf_538622eb3c5000ab{}
	}
	__obf_d3fe11085fea0736 := []__obf_617b2253eaac7f7a{}
	for _, __obf_72c924a23efc3585 := range __obf_44aea3b3531d146b {
		if !__obf_72c924a23efc3585.__obf_b020b7b9f5f7dacc {
			__obf_d3fe11085fea0736 = append(__obf_d3fe11085fea0736, __obf_617b2253eaac7f7a{__obf_c091c27480fae550: __obf_72c924a23efc3585.__obf_56b28257c2a1cc8d.Encoder.(*__obf_f750ce163e10b96f), __obf_a3d812fac643004b: __obf_72c924a23efc3585.__obf_a3d812fac643004b})
		}
	}
	return &__obf_142f562362b041d8{__obf_edcded11af6ebc4c, __obf_d3fe11085fea0736}
}

func __obf_347e3fc852eb458f(__obf_99ec45908bceabdb *__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c reflect2.Type) __obf_07e25a5cafedd0df {
	__obf_c091c27480fae550 := __obf_1f22c50e61ed6284(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	if __obf_c091c27480fae550 != nil {
		return __obf_c091c27480fae550
	}
	__obf_105d1cd1a59219cd := __obf_edcded11af6ebc4c.Kind()
	switch __obf_105d1cd1a59219cd {
	case reflect.Interface:
		return &__obf_16a4e7713c271876{__obf_edcded11af6ebc4c}
	case reflect.Struct:
		return &__obf_142f562362b041d8{__obf_edcded11af6ebc4c: __obf_edcded11af6ebc4c}
	case reflect.Array:
		return &__obf_d81c3f14e6ac6dc4{}
	case reflect.Slice:
		return &__obf_577577703e34ae45{}
	case reflect.Map:
		return __obf_8be1ac6303cc573e(__obf_99ec45908bceabdb, __obf_edcded11af6ebc4c)
	case reflect.Ptr:
		return &OptionalEncoder{}
	default:
		return &__obf_531b9739373c7313{__obf_ea0680f8b461a85b: fmt.Errorf("unsupported type: %v", __obf_edcded11af6ebc4c)}
	}
}

func __obf_169b1277fbacdc3f(__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0, __obf_ffbea95e31df6856, new *Binding) (__obf_81ed0382471e779e, __obf_5019fe5e38747cb9 bool) {
	__obf_c72b689093b7f83f := new.Field.Tag().Get(__obf_c52e0bcfad4b8b71.__obf_2b09ab97bcd3be4f()) != ""
	__obf_7aa1a31c8a30e7ce := __obf_ffbea95e31df6856.Field.Tag().Get(__obf_c52e0bcfad4b8b71.__obf_2b09ab97bcd3be4f()) != ""
	if __obf_c72b689093b7f83f {
		if __obf_7aa1a31c8a30e7ce {
			if len(__obf_ffbea95e31df6856.__obf_82fe242d978d1bf2) > len(new.__obf_82fe242d978d1bf2) {
				return true, false
			} else if len(new.__obf_82fe242d978d1bf2) > len(__obf_ffbea95e31df6856.__obf_82fe242d978d1bf2) {
				return false, true
			} else {
				return true, true
			}
		} else {
			return true, false
		}
	} else {
		if __obf_7aa1a31c8a30e7ce {
			return true, false
		}
		if len(__obf_ffbea95e31df6856.__obf_82fe242d978d1bf2) > len(new.__obf_82fe242d978d1bf2) {
			return true, false
		} else if len(new.__obf_82fe242d978d1bf2) > len(__obf_ffbea95e31df6856.__obf_82fe242d978d1bf2) {
			return false, true
		} else {
			return true, true
		}
	}
}

type __obf_f750ce163e10b96f struct {
	__obf_ad34f8a6de2b01cb reflect2.StructField
	__obf_0ae3f6628b391580 ValEncoder
	__obf_1f6282d06d1e12e1 bool
}

func (__obf_c091c27480fae550 *__obf_f750ce163e10b96f) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_c34985751cfe0592 := __obf_c091c27480fae550.__obf_ad34f8a6de2b01cb.UnsafeGet(__obf_575c04f2b9d91315)
	__obf_c091c27480fae550.__obf_0ae3f6628b391580.
		Encode(__obf_c34985751cfe0592, __obf_d8f50bcfe87d8b47)
	if __obf_d8f50bcfe87d8b47.Error != nil && __obf_d8f50bcfe87d8b47.Error != io.EOF {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("%s: %s", __obf_c091c27480fae550.__obf_ad34f8a6de2b01cb.Name(), __obf_d8f50bcfe87d8b47.Error.Error())
	}
}

func (__obf_c091c27480fae550 *__obf_f750ce163e10b96f) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_c34985751cfe0592 := __obf_c091c27480fae550.__obf_ad34f8a6de2b01cb.UnsafeGet(__obf_575c04f2b9d91315)
	return __obf_c091c27480fae550.__obf_0ae3f6628b391580.IsEmpty(__obf_c34985751cfe0592)
}

func (__obf_c091c27480fae550 *__obf_f750ce163e10b96f) IsEmbeddedPtrNil(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	__obf_bdbbea1ea376ff9b, __obf_a7ca508d9e901782 := __obf_c091c27480fae550.__obf_0ae3f6628b391580.(IsEmbeddedPtrNil)
	if !__obf_a7ca508d9e901782 {
		return false
	}
	__obf_c34985751cfe0592 := __obf_c091c27480fae550.__obf_ad34f8a6de2b01cb.UnsafeGet(__obf_575c04f2b9d91315)
	return __obf_bdbbea1ea376ff9b.IsEmbeddedPtrNil(__obf_c34985751cfe0592)
}

type IsEmbeddedPtrNil interface {
	IsEmbeddedPtrNil(__obf_575c04f2b9d91315 unsafe.Pointer) bool
}

type __obf_142f562362b041d8 struct {
	__obf_edcded11af6ebc4c reflect2.Type
	__obf_c397376f03d6d1bb []__obf_617b2253eaac7f7a
}

type __obf_617b2253eaac7f7a struct {
	__obf_c091c27480fae550 *__obf_f750ce163e10b96f
	__obf_a3d812fac643004b string
}

func (__obf_c091c27480fae550 *__obf_142f562362b041d8) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteObjectStart()
	__obf_3e3349d77786069e := false
	for _, __obf_ad34f8a6de2b01cb := range __obf_c091c27480fae550.__obf_c397376f03d6d1bb {
		if __obf_ad34f8a6de2b01cb.__obf_c091c27480fae550.__obf_1f6282d06d1e12e1 && __obf_ad34f8a6de2b01cb.__obf_c091c27480fae550.IsEmpty(__obf_575c04f2b9d91315) {
			continue
		}
		if __obf_ad34f8a6de2b01cb.__obf_c091c27480fae550.IsEmbeddedPtrNil(__obf_575c04f2b9d91315) {
			continue
		}
		if __obf_3e3349d77786069e {
			__obf_d8f50bcfe87d8b47.
				WriteMore()
		}
		__obf_d8f50bcfe87d8b47.
			WriteObjectField(__obf_ad34f8a6de2b01cb.__obf_a3d812fac643004b)
		__obf_ad34f8a6de2b01cb.__obf_c091c27480fae550.
			Encode(__obf_575c04f2b9d91315, __obf_d8f50bcfe87d8b47)
		__obf_3e3349d77786069e = true
	}
	__obf_d8f50bcfe87d8b47.
		WriteObjectEnd()
	if __obf_d8f50bcfe87d8b47.Error != nil && __obf_d8f50bcfe87d8b47.Error != io.EOF {
		__obf_d8f50bcfe87d8b47.
			Error = fmt.Errorf("%v.%s", __obf_c091c27480fae550.__obf_edcded11af6ebc4c, __obf_d8f50bcfe87d8b47.Error.Error())
	}
}

func (__obf_c091c27480fae550 *__obf_142f562362b041d8) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return false
}

type __obf_538622eb3c5000ab struct {
}

func (__obf_c091c27480fae550 *__obf_538622eb3c5000ab) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.
		WriteEmptyObject()
}

func (__obf_c091c27480fae550 *__obf_538622eb3c5000ab) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return false
}

type __obf_5e5315ac220d9e9b struct {
	__obf_e6894379459b04e8 ValEncoder
}

func (__obf_c091c27480fae550 *__obf_5e5315ac220d9e9b) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
	__obf_c091c27480fae550.__obf_e6894379459b04e8.
		Encode(__obf_575c04f2b9d91315, __obf_d8f50bcfe87d8b47)
	__obf_d8f50bcfe87d8b47.__obf_7340077d41996822('"')
}

func (__obf_c091c27480fae550 *__obf_5e5315ac220d9e9b) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_e6894379459b04e8.IsEmpty(__obf_575c04f2b9d91315)
}

type __obf_7de5e4e3e891d771 struct {
	__obf_e6894379459b04e8 ValEncoder
	__obf_c52e0bcfad4b8b71 *__obf_30fe5c95cabd69c0
}

func (__obf_c091c27480fae550 *__obf_7de5e4e3e891d771) Encode(__obf_575c04f2b9d91315 unsafe.Pointer, __obf_d8f50bcfe87d8b47 *Stream) {
	__obf_21ff191365e47676 := __obf_c091c27480fae550.__obf_c52e0bcfad4b8b71.BorrowStream(nil)
	__obf_21ff191365e47676.
		Attachment = __obf_d8f50bcfe87d8b47.Attachment
	defer __obf_c091c27480fae550.__obf_c52e0bcfad4b8b71.ReturnStream(__obf_21ff191365e47676)
	__obf_c091c27480fae550.__obf_e6894379459b04e8.
		Encode(__obf_575c04f2b9d91315, __obf_21ff191365e47676)
	__obf_d8f50bcfe87d8b47.
		WriteString(string(__obf_21ff191365e47676.Buffer()))
}

func (__obf_c091c27480fae550 *__obf_7de5e4e3e891d771) IsEmpty(__obf_575c04f2b9d91315 unsafe.Pointer) bool {
	return __obf_c091c27480fae550.__obf_e6894379459b04e8.IsEmpty(__obf_575c04f2b9d91315)
}
