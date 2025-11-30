package __obf_de86cdc8ae98b45b

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
)

var (
	__obf_f248948128de9dd0 = reflect.TypeOf((*any)(nil)).Elem()
	__obf_620340e4a4ed2e24 = reflect.TypeOf((*string)(nil)).Elem()
	__obf_20663e276851f318 = reflect.TypeOf((*bool)(nil)).Elem()
)

var __obf_e8915b214d32aa7b []__obf_b74c57495471e351

//nolint:gochecknoinits
func init() {
	__obf_e8915b214d32aa7b = []__obf_b74c57495471e351{
		reflect.Bool:          __obf_3be7812b01824ffd,
		reflect.Int:           __obf_58e1420db8e49e5b,
		reflect.Int8:          __obf_58e1420db8e49e5b,
		reflect.Int16:         __obf_58e1420db8e49e5b,
		reflect.Int32:         __obf_58e1420db8e49e5b,
		reflect.Int64:         __obf_58e1420db8e49e5b,
		reflect.Uint:          __obf_460b1f6481895b57,
		reflect.Uint8:         __obf_460b1f6481895b57,
		reflect.Uint16:        __obf_460b1f6481895b57,
		reflect.Uint32:        __obf_460b1f6481895b57,
		reflect.Uint64:        __obf_460b1f6481895b57,
		reflect.Float32:       __obf_42e61e6f8a4ef035,
		reflect.Float64:       __obf_57e51fbd6a8c8a03,
		reflect.Complex64:     __obf_a3d91a4eae13357e,
		reflect.Complex128:    __obf_a3d91a4eae13357e,
		reflect.Array:         __obf_28619ae619606d09,
		reflect.Chan:          __obf_a3d91a4eae13357e,
		reflect.Func:          __obf_a3d91a4eae13357e,
		reflect.Interface:     __obf_19f61f7db3d5aa0f,
		reflect.Map:           __obf_fc3c99dfcc57d122,
		reflect.Pointer:       __obf_a3d91a4eae13357e,
		reflect.Slice:         __obf_68ff6383b6d6d266,
		reflect.String:        __obf_47189e506d9624a7,
		reflect.Struct:        __obf_d6cbe6ba2b1343be,
		reflect.UnsafePointer: __obf_a3d91a4eae13357e,
	}
}

func __obf_25e3584cea7b6666(__obf_2dea2b2773c8afdf reflect.Type) __obf_b74c57495471e351 {
	if __obf_17732590722140e7, __obf_77cfff95beb3cc99 := __obf_9f46fbaef88684c0.Load(__obf_2dea2b2773c8afdf); __obf_77cfff95beb3cc99 {
		return __obf_17732590722140e7.(__obf_b74c57495471e351)
	}
	__obf_69331012fc3414ab := _getDecoder(__obf_2dea2b2773c8afdf)
	__obf_9f46fbaef88684c0.
		Store(__obf_2dea2b2773c8afdf, __obf_69331012fc3414ab)
	return __obf_69331012fc3414ab
}

func _getDecoder(__obf_2dea2b2773c8afdf reflect.Type) __obf_b74c57495471e351 {
	__obf_f02bd2095cf36557 := __obf_2dea2b2773c8afdf.Kind()

	if __obf_f02bd2095cf36557 == reflect.Ptr {
		if _, __obf_77cfff95beb3cc99 := __obf_9f46fbaef88684c0.Load(__obf_2dea2b2773c8afdf.Elem()); __obf_77cfff95beb3cc99 {
			return __obf_502bf223615fa6e8(__obf_2dea2b2773c8afdf)
		}
	}

	if __obf_2dea2b2773c8afdf.Implements(__obf_bb521c0d47017b87) {
		return __obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_1c78c492c5f31284)
	}
	if __obf_2dea2b2773c8afdf.Implements(__obf_46f8dad60f3aa967) {
		return __obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_61403dbfcd07f852)
	}
	if __obf_2dea2b2773c8afdf.Implements(__obf_cd3378d776026b55) {
		return __obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_fc0d8539c723b6ae)
	}
	if __obf_2dea2b2773c8afdf.Implements(__obf_0d5a1d97d27c066f) {
		return __obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_aa68827c91261ea1)
	}

	// Addressable struct field value.
	if __obf_f02bd2095cf36557 != reflect.Pointer {
		__obf_01f0154f1fce8e5a := reflect.PointerTo(__obf_2dea2b2773c8afdf)
		if __obf_01f0154f1fce8e5a.Implements(__obf_bb521c0d47017b87) {
			return __obf_a6fe33e518fc3321(__obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_1c78c492c5f31284))
		}
		if __obf_01f0154f1fce8e5a.Implements(__obf_46f8dad60f3aa967) {
			return __obf_a6fe33e518fc3321(__obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_61403dbfcd07f852))
		}
		if __obf_01f0154f1fce8e5a.Implements(__obf_cd3378d776026b55) {
			return __obf_a6fe33e518fc3321(__obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_fc0d8539c723b6ae))
		}
		if __obf_01f0154f1fce8e5a.Implements(__obf_0d5a1d97d27c066f) {
			return __obf_a6fe33e518fc3321(__obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, __obf_aa68827c91261ea1))
		}
	}

	switch __obf_f02bd2095cf36557 {
	case reflect.Pointer:
		return __obf_502bf223615fa6e8(__obf_2dea2b2773c8afdf)
	case reflect.Slice:
		__obf_dd3641741f7dbdf3 := __obf_2dea2b2773c8afdf.Elem()
		if __obf_dd3641741f7dbdf3.Kind() == reflect.Uint8 {
			return __obf_431a947ae748a61f
		}
		if __obf_dd3641741f7dbdf3 == __obf_620340e4a4ed2e24 {
			return __obf_8de5f6238f2c13bf
		}
	case reflect.Array:
		if __obf_2dea2b2773c8afdf.Elem().Kind() == reflect.Uint8 {
			return __obf_9dd8ad0f86a4d0ca
		}
	case reflect.Map:
		if __obf_2dea2b2773c8afdf.Key() == __obf_620340e4a4ed2e24 {
			switch __obf_2dea2b2773c8afdf.Elem() {
			case __obf_620340e4a4ed2e24:
				return __obf_62de82f75c0986e4
			case __obf_f248948128de9dd0:
				return __obf_b673e9c8a63629b4
			}
		}
	}

	return __obf_e8915b214d32aa7b[__obf_f02bd2095cf36557]
}

func __obf_502bf223615fa6e8(__obf_2dea2b2773c8afdf reflect.Type) __obf_b74c57495471e351 {
	__obf_0fe2a689830f8985 := __obf_25e3584cea7b6666(__obf_2dea2b2773c8afdf.Elem())
	return func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
		if __obf_dcaad1165bb07af9.__obf_61d7724cadbffcc4() {
			if !__obf_17732590722140e7.IsNil() {
				__obf_17732590722140e7.
					Set(__obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_2dea2b2773c8afdf).Elem())
			}
			return __obf_dcaad1165bb07af9.DecodeNil()
		}
		if __obf_17732590722140e7.IsNil() {
			__obf_17732590722140e7.
				Set(__obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_2dea2b2773c8afdf.Elem()))
		}
		return __obf_0fe2a689830f8985(__obf_dcaad1165bb07af9, __obf_17732590722140e7.Elem())
	}
}

func __obf_a6fe33e518fc3321(__obf_69331012fc3414ab __obf_b74c57495471e351) __obf_b74c57495471e351 {
	return func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
		if !__obf_17732590722140e7.CanAddr() {
			return fmt.Errorf("msgpack: Decode(nonaddressable %T)", __obf_17732590722140e7.Interface())
		}
		return __obf_69331012fc3414ab(__obf_dcaad1165bb07af9, __obf_17732590722140e7.Addr())
	}
}

func __obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf reflect.Type, __obf_69331012fc3414ab __obf_b74c57495471e351) __obf_b74c57495471e351 {
	if __obf_c3d2a512f69a483f(__obf_2dea2b2773c8afdf.Kind()) {
		return func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
			if __obf_dcaad1165bb07af9.__obf_61d7724cadbffcc4() {
				return __obf_dcaad1165bb07af9.__obf_0913827f98642042(__obf_17732590722140e7)
			}
			if __obf_17732590722140e7.IsNil() {
				__obf_17732590722140e7.
					Set(__obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_2dea2b2773c8afdf.Elem()))
			}
			return __obf_69331012fc3414ab(__obf_dcaad1165bb07af9, __obf_17732590722140e7)
		}
	}

	return func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
		if __obf_dcaad1165bb07af9.__obf_61d7724cadbffcc4() {
			return __obf_dcaad1165bb07af9.__obf_0913827f98642042(__obf_17732590722140e7)
		}
		return __obf_69331012fc3414ab(__obf_dcaad1165bb07af9, __obf_17732590722140e7)
	}
}

func __obf_3be7812b01824ffd(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_6b784159a101b716, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeBool()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_17732590722140e7.
		SetBool(__obf_6b784159a101b716)
	return nil
}

func __obf_19f61f7db3d5aa0f(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	if __obf_17732590722140e7.IsNil() {
		return __obf_dcaad1165bb07af9.__obf_489d5cb2b054d0c9(__obf_17732590722140e7)
	}
	return __obf_dcaad1165bb07af9.DecodeValue(__obf_17732590722140e7.Elem())
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_489d5cb2b054d0c9(__obf_17732590722140e7 reflect.Value) error {
	__obf_7c0d6551e0939afa, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	if __obf_7c0d6551e0939afa != nil {
		if __obf_17732590722140e7.Type() == __obf_2d0560f267da57d8 {
			if __obf_7c0d6551e0939afa, __obf_77cfff95beb3cc99 := __obf_7c0d6551e0939afa.(string); __obf_77cfff95beb3cc99 {
				__obf_17732590722140e7.
					Set(reflect.ValueOf(errors.New(__obf_7c0d6551e0939afa)))
				return nil
			}
		}
		__obf_17732590722140e7.
			Set(reflect.ValueOf(__obf_7c0d6551e0939afa))
	}

	return nil
}

func __obf_a3d91a4eae13357e(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	return fmt.Errorf("msgpack: Decode(unsupported %s)", __obf_17732590722140e7.Type())
}

//------------------------------------------------------------------------------

func __obf_1c78c492c5f31284(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_0fe2a689830f8985 := __obf_17732590722140e7.Interface().(CustomDecoder)
	return __obf_0fe2a689830f8985.DecodeMsgpack(__obf_dcaad1165bb07af9)
}

func __obf_61403dbfcd07f852(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	var __obf_a7fd7acf2bd4435f []byte
	__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = make([]byte, 0, 64)
	if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_a7fd7acf2bd4435f = __obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf
	__obf_dcaad1165bb07af9.__obf_fabeb54cc5377fbf = nil
	__obf_2c3746c7a528dc28 := __obf_17732590722140e7.Interface().(Unmarshaler)
	return __obf_2c3746c7a528dc28.UnmarshalMsgpack(__obf_a7fd7acf2bd4435f)
}

func __obf_fc0d8539c723b6ae(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_50d002bc8b306017, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeBytes()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2c3746c7a528dc28 := __obf_17732590722140e7.Interface().(encoding.BinaryUnmarshaler)
	return __obf_2c3746c7a528dc28.UnmarshalBinary(__obf_50d002bc8b306017)
}

func __obf_aa68827c91261ea1(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_50d002bc8b306017, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeBytes()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	__obf_2c3746c7a528dc28 := __obf_17732590722140e7.Interface().(encoding.TextUnmarshaler)
	return __obf_2c3746c7a528dc28.UnmarshalText(__obf_50d002bc8b306017)
}
