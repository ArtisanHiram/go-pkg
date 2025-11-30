package __obf_3e0c303610a19bd4

import (
	"encoding"
	"fmt"
	"reflect"
)

var __obf_82b8ea2cef466c5b []__obf_9cac677a20c1be1c

//nolint:gochecknoinits
func init() {
	__obf_82b8ea2cef466c5b = []__obf_9cac677a20c1be1c{
		reflect.Bool:          __obf_d2869c9a3200bea2,
		reflect.Int:           __obf_da00a84d1175f148,
		reflect.Int8:          __obf_38a64e9c8c30348f,
		reflect.Int16:         __obf_192fd05a8152c0cf,
		reflect.Int32:         __obf_60d6bd6b5e468734,
		reflect.Int64:         __obf_f9c867edad184432,
		reflect.Uint:          __obf_c079cfc8e2339c96,
		reflect.Uint8:         __obf_bfc025a7f29928a9,
		reflect.Uint16:        __obf_81f7a2d6252d3484,
		reflect.Uint32:        __obf_37b3e9bbb21b1a1b,
		reflect.Uint64:        __obf_f2dc55a2f1f1d57e,
		reflect.Float32:       __obf_7de64ad2273f4b55,
		reflect.Float64:       __obf_239ea56f7cf6dcb6,
		reflect.Complex64:     __obf_84d1348de9dc6d4b,
		reflect.Complex128:    __obf_84d1348de9dc6d4b,
		reflect.Array:         __obf_fb2bd25c8bdcdcf8,
		reflect.Chan:          __obf_84d1348de9dc6d4b,
		reflect.Func:          __obf_84d1348de9dc6d4b,
		reflect.Interface:     __obf_fe0a14e8a8e85b87,
		reflect.Map:           __obf_10a19aa600413a97,
		reflect.Ptr:           __obf_84d1348de9dc6d4b,
		reflect.Slice:         __obf_ddcfd0255d36e7b4,
		reflect.String:        __obf_4013add33227fac1,
		reflect.Struct:        __obf_cbc3a9db722bf76c,
		reflect.UnsafePointer: __obf_84d1348de9dc6d4b,
	}
}

func __obf_cc16be3b4933c617(__obf_616f98efc80197c6 reflect.Type) __obf_9cac677a20c1be1c {
	if __obf_63bbcee86d44fdde, __obf_ce8bef141eff8aab := __obf_b932a15e2d3972de.Load(__obf_616f98efc80197c6); __obf_ce8bef141eff8aab {
		return __obf_63bbcee86d44fdde.(__obf_9cac677a20c1be1c)
	}
	__obf_6ff63d98a176dcfe := _getEncoder(__obf_616f98efc80197c6)
	__obf_b932a15e2d3972de.
		Store(__obf_616f98efc80197c6, __obf_6ff63d98a176dcfe)
	return __obf_6ff63d98a176dcfe
}

func _getEncoder(__obf_616f98efc80197c6 reflect.Type) __obf_9cac677a20c1be1c {
	__obf_018b20441ecc0301 := __obf_616f98efc80197c6.Kind()

	if __obf_018b20441ecc0301 == reflect.Pointer {
		if _, __obf_ce8bef141eff8aab := __obf_b932a15e2d3972de.Load(__obf_616f98efc80197c6.Elem()); __obf_ce8bef141eff8aab {
			return __obf_3a315178032b651a(__obf_616f98efc80197c6)
		}
	}

	if __obf_616f98efc80197c6.Implements(__obf_c08f5968e7619fb2) {
		return __obf_e790239966ac9649
	}
	if __obf_616f98efc80197c6.Implements(__obf_9e290b0c7c8448c5) {
		return __obf_272c459bd7da7f67
	}
	if __obf_616f98efc80197c6.Implements(__obf_c4304d6f5eef9ca6) {
		return __obf_4ec29cc81f341361
	}
	if __obf_616f98efc80197c6.Implements(__obf_09aab39b0e9cc30f) {
		return __obf_ea73e6847b71cfa3
	}

	// Addressable struct field value.
	if __obf_018b20441ecc0301 != reflect.Pointer {
		__obf_b5a4664807537c0d := reflect.PointerTo(__obf_616f98efc80197c6)
		if __obf_b5a4664807537c0d.Implements(__obf_c08f5968e7619fb2) {
			return __obf_590a6fbccc45ea35
		}
		if __obf_b5a4664807537c0d.Implements(__obf_9e290b0c7c8448c5) {
			return __obf_b0989fe698a253b0
		}
		if __obf_b5a4664807537c0d.Implements(__obf_c4304d6f5eef9ca6) {
			return __obf_96fd2715c424b835
		}
		if __obf_b5a4664807537c0d.Implements(__obf_09aab39b0e9cc30f) {
			return __obf_1a07e316711e9952
		}
	}

	if __obf_616f98efc80197c6 == __obf_319e36d3ccef195a {
		return __obf_34c0ebbd2cc6db0b
	}

	switch __obf_018b20441ecc0301 {
	case reflect.Ptr:
		return __obf_3a315178032b651a(__obf_616f98efc80197c6)
	case reflect.Slice:
		__obf_265b8ddf380ff77a := __obf_616f98efc80197c6.Elem()
		if __obf_265b8ddf380ff77a.Kind() == reflect.Uint8 {
			return __obf_3242ef092432362b
		}
		if __obf_265b8ddf380ff77a == __obf_24e46306ee5355ff {
			return __obf_f5ae91e031add6e1
		}
	case reflect.Array:
		if __obf_616f98efc80197c6.Elem().Kind() == reflect.Uint8 {
			return __obf_39d5205b9d649841
		}
	case reflect.Map:
		if __obf_616f98efc80197c6.Key() == __obf_24e46306ee5355ff {
			switch __obf_616f98efc80197c6.Elem() {
			case __obf_24e46306ee5355ff:
				return __obf_af298ee09c591531
			case __obf_132fa72af3bf692d:
				return __obf_6666d9a2247b6097
			case __obf_164ae3a742d185b4:
				return __obf_026b0bc834cd7f40
			}
		}
	}

	return __obf_82b8ea2cef466c5b[__obf_018b20441ecc0301]
}

func __obf_3a315178032b651a(__obf_616f98efc80197c6 reflect.Type) __obf_9cac677a20c1be1c {
	__obf_02e33ee3dfeba894 := __obf_cc16be3b4933c617(__obf_616f98efc80197c6.Elem())
	return func(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if __obf_63bbcee86d44fdde.IsNil() {
			return __obf_77240dc7776b60b4.EncodeNil()
		}
		return __obf_02e33ee3dfeba894(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde.Elem())
	}
}

func __obf_590a6fbccc45ea35(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if !__obf_63bbcee86d44fdde.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_63bbcee86d44fdde.Interface())
	}
	__obf_02e33ee3dfeba894 := __obf_63bbcee86d44fdde.Addr().Interface().(CustomEncoder)
	return __obf_02e33ee3dfeba894.EncodeMsgpack(__obf_77240dc7776b60b4)
}

func __obf_e790239966ac9649(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_ebd8855a2f3055ad(__obf_63bbcee86d44fdde.Kind()) && __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	__obf_02e33ee3dfeba894 := __obf_63bbcee86d44fdde.Interface().(CustomEncoder)
	return __obf_02e33ee3dfeba894.EncodeMsgpack(__obf_77240dc7776b60b4)
}

func __obf_b0989fe698a253b0(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if !__obf_63bbcee86d44fdde.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_63bbcee86d44fdde.Interface())
	}
	return __obf_272c459bd7da7f67(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde.Addr())
}

func __obf_272c459bd7da7f67(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_ebd8855a2f3055ad(__obf_63bbcee86d44fdde.Kind()) && __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	__obf_9a68ec3a8d69fdbf := __obf_63bbcee86d44fdde.Interface().(Marshaler)
	__obf_11bcc66cde095c11, __obf_8882f6cf6e378ded := __obf_9a68ec3a8d69fdbf.MarshalMsgpack()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	_, __obf_8882f6cf6e378ded = __obf_77240dc7776b60b4.__obf_279ea11cfd0ea98d.Write(__obf_11bcc66cde095c11)
	return __obf_8882f6cf6e378ded
}

func __obf_d2869c9a3200bea2(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return __obf_77240dc7776b60b4.EncodeBool(__obf_63bbcee86d44fdde.Bool())
}

func __obf_fe0a14e8a8e85b87(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	return __obf_77240dc7776b60b4.EncodeValue(__obf_63bbcee86d44fdde.Elem())
}

func __obf_34c0ebbd2cc6db0b(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	return __obf_77240dc7776b60b4.EncodeString(__obf_63bbcee86d44fdde.Interface().(error).Error())
}

func __obf_84d1348de9dc6d4b(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return fmt.Errorf("msgpack: Encode(unsupported %s)", __obf_63bbcee86d44fdde.Type())
}

func __obf_ebd8855a2f3055ad(__obf_018b20441ecc0301 reflect.Kind) bool {
	switch __obf_018b20441ecc0301 {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func __obf_0ecf01dfd95d226f(__obf_cb646455b7fd3ba7 reflect.Type) bool {
	if __obf_cb646455b7fd3ba7.Kind() == reflect.Ptr {
		__obf_cb646455b7fd3ba7 = __obf_cb646455b7fd3ba7.Elem()
	}
	return __obf_ebd8855a2f3055ad(__obf_cb646455b7fd3ba7.Kind())
}

//------------------------------------------------------------------------------

func __obf_96fd2715c424b835(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if !__obf_63bbcee86d44fdde.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_63bbcee86d44fdde.Interface())
	}
	return __obf_4ec29cc81f341361(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde.Addr())
}

func __obf_4ec29cc81f341361(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_ebd8855a2f3055ad(__obf_63bbcee86d44fdde.Kind()) && __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	__obf_9a68ec3a8d69fdbf := __obf_63bbcee86d44fdde.Interface().(encoding.BinaryMarshaler)
	__obf_4d827abff69b9b40, __obf_8882f6cf6e378ded := __obf_9a68ec3a8d69fdbf.MarshalBinary()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	return __obf_77240dc7776b60b4.EncodeBytes(__obf_4d827abff69b9b40)
}

//------------------------------------------------------------------------------

func __obf_1a07e316711e9952(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if !__obf_63bbcee86d44fdde.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_63bbcee86d44fdde.Interface())
	}
	return __obf_ea73e6847b71cfa3(__obf_77240dc7776b60b4, __obf_63bbcee86d44fdde.Addr())
}

func __obf_ea73e6847b71cfa3(__obf_77240dc7776b60b4 *Encoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_ebd8855a2f3055ad(__obf_63bbcee86d44fdde.Kind()) && __obf_63bbcee86d44fdde.IsNil() {
		return __obf_77240dc7776b60b4.EncodeNil()
	}
	__obf_9a68ec3a8d69fdbf := __obf_63bbcee86d44fdde.Interface().(encoding.TextMarshaler)
	__obf_4d827abff69b9b40, __obf_8882f6cf6e378ded := __obf_9a68ec3a8d69fdbf.MarshalText()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	return __obf_77240dc7776b60b4.EncodeBytes(__obf_4d827abff69b9b40)
}
