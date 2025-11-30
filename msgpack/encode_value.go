package __obf_de86cdc8ae98b45b

import (
	"encoding"
	"fmt"
	"reflect"
)

var __obf_000c9134f42f238b []__obf_9399ab762e15c2a7

//nolint:gochecknoinits
func init() {
	__obf_000c9134f42f238b = []__obf_9399ab762e15c2a7{
		reflect.Bool:          __obf_be4417c89afe5ce2,
		reflect.Int:           __obf_0f70427f0c5e77bf,
		reflect.Int8:          __obf_8858b34e9db12038,
		reflect.Int16:         __obf_43983a1b215551a5,
		reflect.Int32:         __obf_1878d2a59a353162,
		reflect.Int64:         __obf_b998d2da39e2780c,
		reflect.Uint:          __obf_35394d5f1dd6a3c5,
		reflect.Uint8:         __obf_d187dcb37efcfe6d,
		reflect.Uint16:        __obf_6148c4456411af75,
		reflect.Uint32:        __obf_baeb98a0e3d5e172,
		reflect.Uint64:        __obf_4722bc9436e891f9,
		reflect.Float32:       __obf_ccb4e0a7821f8534,
		reflect.Float64:       __obf_30345a7c18493f02,
		reflect.Complex64:     __obf_3b3084541220549e,
		reflect.Complex128:    __obf_3b3084541220549e,
		reflect.Array:         __obf_743bdcb3d7776bf9,
		reflect.Chan:          __obf_3b3084541220549e,
		reflect.Func:          __obf_3b3084541220549e,
		reflect.Interface:     __obf_9b27d42a37edd5fa,
		reflect.Map:           __obf_90655a0691bba685,
		reflect.Ptr:           __obf_3b3084541220549e,
		reflect.Slice:         __obf_0b25b51b78e0d671,
		reflect.String:        __obf_439390bb02ad86d8,
		reflect.Struct:        __obf_95de04fb2a8b0af3,
		reflect.UnsafePointer: __obf_3b3084541220549e,
	}
}

func __obf_d4a1e60c459c24e4(__obf_2dea2b2773c8afdf reflect.Type) __obf_9399ab762e15c2a7 {
	if __obf_17732590722140e7, __obf_77cfff95beb3cc99 := __obf_1d13df182269d8e8.Load(__obf_2dea2b2773c8afdf); __obf_77cfff95beb3cc99 {
		return __obf_17732590722140e7.(__obf_9399ab762e15c2a7)
	}
	__obf_69331012fc3414ab := _getEncoder(__obf_2dea2b2773c8afdf)
	__obf_1d13df182269d8e8.
		Store(__obf_2dea2b2773c8afdf, __obf_69331012fc3414ab)
	return __obf_69331012fc3414ab
}

func _getEncoder(__obf_2dea2b2773c8afdf reflect.Type) __obf_9399ab762e15c2a7 {
	__obf_f02bd2095cf36557 := __obf_2dea2b2773c8afdf.Kind()

	if __obf_f02bd2095cf36557 == reflect.Pointer {
		if _, __obf_77cfff95beb3cc99 := __obf_1d13df182269d8e8.Load(__obf_2dea2b2773c8afdf.Elem()); __obf_77cfff95beb3cc99 {
			return __obf_48e1ad39f2749ae8(__obf_2dea2b2773c8afdf)
		}
	}

	if __obf_2dea2b2773c8afdf.Implements(__obf_97b6aca480193051) {
		return __obf_db68bceccf09f9aa
	}
	if __obf_2dea2b2773c8afdf.Implements(__obf_38acb3bdfb404f10) {
		return __obf_a1489e97b6bcc224
	}
	if __obf_2dea2b2773c8afdf.Implements(__obf_6ff2001fa5096e39) {
		return __obf_c95a25011a653211
	}
	if __obf_2dea2b2773c8afdf.Implements(__obf_e9715cc721daa9f1) {
		return __obf_04fe8b9e4d5079a0
	}

	// Addressable struct field value.
	if __obf_f02bd2095cf36557 != reflect.Pointer {
		__obf_01f0154f1fce8e5a := reflect.PointerTo(__obf_2dea2b2773c8afdf)
		if __obf_01f0154f1fce8e5a.Implements(__obf_97b6aca480193051) {
			return __obf_8e4a1d666153a9a4
		}
		if __obf_01f0154f1fce8e5a.Implements(__obf_38acb3bdfb404f10) {
			return __obf_e7a1f1805989c67d
		}
		if __obf_01f0154f1fce8e5a.Implements(__obf_6ff2001fa5096e39) {
			return __obf_9f4e73bf9bbe3482
		}
		if __obf_01f0154f1fce8e5a.Implements(__obf_e9715cc721daa9f1) {
			return __obf_bf9b1bb778d10bee
		}
	}

	if __obf_2dea2b2773c8afdf == __obf_2d0560f267da57d8 {
		return __obf_96dfce3ecc743624
	}

	switch __obf_f02bd2095cf36557 {
	case reflect.Ptr:
		return __obf_48e1ad39f2749ae8(__obf_2dea2b2773c8afdf)
	case reflect.Slice:
		__obf_dd3641741f7dbdf3 := __obf_2dea2b2773c8afdf.Elem()
		if __obf_dd3641741f7dbdf3.Kind() == reflect.Uint8 {
			return __obf_efad3cfa905a24ec
		}
		if __obf_dd3641741f7dbdf3 == __obf_620340e4a4ed2e24 {
			return __obf_c225cd0483ee2104
		}
	case reflect.Array:
		if __obf_2dea2b2773c8afdf.Elem().Kind() == reflect.Uint8 {
			return __obf_2c6376330bc8c07a
		}
	case reflect.Map:
		if __obf_2dea2b2773c8afdf.Key() == __obf_620340e4a4ed2e24 {
			switch __obf_2dea2b2773c8afdf.Elem() {
			case __obf_620340e4a4ed2e24:
				return __obf_9805b50aa1981a7e
			case __obf_20663e276851f318:
				return __obf_59e60c775a39bdfb
			case __obf_f248948128de9dd0:
				return __obf_62ccef293845516e
			}
		}
	}

	return __obf_000c9134f42f238b[__obf_f02bd2095cf36557]
}

func __obf_48e1ad39f2749ae8(__obf_2dea2b2773c8afdf reflect.Type) __obf_9399ab762e15c2a7 {
	__obf_2981aff99465546b := __obf_d4a1e60c459c24e4(__obf_2dea2b2773c8afdf.Elem())
	return func(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
		if __obf_17732590722140e7.IsNil() {
			return __obf_7bae0b613da60dd3.EncodeNil()
		}
		return __obf_2981aff99465546b(__obf_7bae0b613da60dd3, __obf_17732590722140e7.Elem())
	}
}

func __obf_8e4a1d666153a9a4(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if !__obf_17732590722140e7.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_17732590722140e7.Interface())
	}
	__obf_2981aff99465546b := __obf_17732590722140e7.Addr().Interface().(CustomEncoder)
	return __obf_2981aff99465546b.EncodeMsgpack(__obf_7bae0b613da60dd3)
}

func __obf_db68bceccf09f9aa(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_c3d2a512f69a483f(__obf_17732590722140e7.Kind()) && __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	__obf_2981aff99465546b := __obf_17732590722140e7.Interface().(CustomEncoder)
	return __obf_2981aff99465546b.EncodeMsgpack(__obf_7bae0b613da60dd3)
}

func __obf_e7a1f1805989c67d(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if !__obf_17732590722140e7.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_17732590722140e7.Interface())
	}
	return __obf_a1489e97b6bcc224(__obf_7bae0b613da60dd3, __obf_17732590722140e7.Addr())
}

func __obf_a1489e97b6bcc224(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_c3d2a512f69a483f(__obf_17732590722140e7.Kind()) && __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	__obf_0b9677586c447323 := __obf_17732590722140e7.Interface().(Marshaler)
	__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_0b9677586c447323.MarshalMsgpack()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	_, __obf_0feb3528b7b709ec = __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.Write(__obf_a7fd7acf2bd4435f)
	return __obf_0feb3528b7b709ec
}

func __obf_be4417c89afe5ce2(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return __obf_7bae0b613da60dd3.EncodeBool(__obf_17732590722140e7.Bool())
}

func __obf_9b27d42a37edd5fa(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	return __obf_7bae0b613da60dd3.EncodeValue(__obf_17732590722140e7.Elem())
}

func __obf_96dfce3ecc743624(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	return __obf_7bae0b613da60dd3.EncodeString(__obf_17732590722140e7.Interface().(error).Error())
}

func __obf_3b3084541220549e(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	return fmt.Errorf("msgpack: Encode(unsupported %s)", __obf_17732590722140e7.Type())
}

func __obf_c3d2a512f69a483f(__obf_f02bd2095cf36557 reflect.Kind) bool {
	switch __obf_f02bd2095cf36557 {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func __obf_f51597fee1c3be88(__obf_b5900a65e8e7d9b5 reflect.Type) bool {
	if __obf_b5900a65e8e7d9b5.Kind() == reflect.Ptr {
		__obf_b5900a65e8e7d9b5 = __obf_b5900a65e8e7d9b5.Elem()
	}
	return __obf_c3d2a512f69a483f(__obf_b5900a65e8e7d9b5.Kind())
}

//------------------------------------------------------------------------------

func __obf_9f4e73bf9bbe3482(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if !__obf_17732590722140e7.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_17732590722140e7.Interface())
	}
	return __obf_c95a25011a653211(__obf_7bae0b613da60dd3, __obf_17732590722140e7.Addr())
}

func __obf_c95a25011a653211(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_c3d2a512f69a483f(__obf_17732590722140e7.Kind()) && __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	__obf_0b9677586c447323 := __obf_17732590722140e7.Interface().(encoding.BinaryMarshaler)
	__obf_50d002bc8b306017, __obf_0feb3528b7b709ec := __obf_0b9677586c447323.MarshalBinary()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	return __obf_7bae0b613da60dd3.EncodeBytes(__obf_50d002bc8b306017)
}

//------------------------------------------------------------------------------

func __obf_bf9b1bb778d10bee(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if !__obf_17732590722140e7.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_17732590722140e7.Interface())
	}
	return __obf_04fe8b9e4d5079a0(__obf_7bae0b613da60dd3, __obf_17732590722140e7.Addr())
}

func __obf_04fe8b9e4d5079a0(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_c3d2a512f69a483f(__obf_17732590722140e7.Kind()) && __obf_17732590722140e7.IsNil() {
		return __obf_7bae0b613da60dd3.EncodeNil()
	}
	__obf_0b9677586c447323 := __obf_17732590722140e7.Interface().(encoding.TextMarshaler)
	__obf_50d002bc8b306017, __obf_0feb3528b7b709ec := __obf_0b9677586c447323.MarshalText()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	return __obf_7bae0b613da60dd3.EncodeBytes(__obf_50d002bc8b306017)
}
