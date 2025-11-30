package __obf_de86cdc8ae98b45b

import (
	"fmt"
	"math"
	"reflect"
)

type __obf_55911355a3eaccfd struct {
	Type    reflect.Type
	Decoder func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value, __obf_eef44dfac5af7de1 int) error
}

var __obf_11268e7b8aa6c195 = make(map[int8]*__obf_55911355a3eaccfd)

type MarshalerUnmarshaler interface {
	Marshaler
	Unmarshaler
}

func RegisterExt(__obf_da58875072df271b int8, __obf_1926a1e373f9e647 MarshalerUnmarshaler) {
	RegisterExtEncoder(__obf_da58875072df271b, __obf_1926a1e373f9e647, func(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) ([]byte, error) {
		__obf_0b9677586c447323 := __obf_17732590722140e7.Interface().(Marshaler)
		return __obf_0b9677586c447323.MarshalMsgpack()
	})
	RegisterExtDecoder(__obf_da58875072df271b, __obf_1926a1e373f9e647, func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value, __obf_eef44dfac5af7de1 int) error {
		__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_8f5a813341411779(__obf_eef44dfac5af7de1)
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		return __obf_17732590722140e7.Interface().(Unmarshaler).UnmarshalMsgpack(__obf_a7fd7acf2bd4435f)
	})
}

func UnregisterExt(__obf_da58875072df271b int8) {
	__obf_68da35c2f34b709c(__obf_da58875072df271b)
	__obf_1108802bcba1cc7b(__obf_da58875072df271b)
}

func RegisterExtEncoder(__obf_da58875072df271b int8, __obf_1926a1e373f9e647 any, __obf_2981aff99465546b func(__obf_6a6518f236eeec6e *Encoder, __obf_17732590722140e7 reflect.Value) ([]byte, error),
) {
	__obf_68da35c2f34b709c(__obf_da58875072df271b)
	__obf_2dea2b2773c8afdf := reflect.TypeOf(__obf_1926a1e373f9e647)
	__obf_f2df1f96c0828a97 := __obf_d01b412f493cff3f(__obf_da58875072df271b, __obf_2dea2b2773c8afdf, __obf_2981aff99465546b)
	__obf_1d13df182269d8e8.
		Store(__obf_da58875072df271b, __obf_2dea2b2773c8afdf)
	__obf_1d13df182269d8e8.
		Store(__obf_2dea2b2773c8afdf, __obf_f2df1f96c0828a97)
	if __obf_2dea2b2773c8afdf.Kind() == reflect.Ptr {
		__obf_1d13df182269d8e8.
			Store(__obf_2dea2b2773c8afdf.Elem(), __obf_3706b5236b1084b3(__obf_f2df1f96c0828a97))
	}
}

func __obf_68da35c2f34b709c(__obf_da58875072df271b int8) {
	__obf_b5900a65e8e7d9b5, __obf_77cfff95beb3cc99 := __obf_1d13df182269d8e8.Load(__obf_da58875072df271b)
	if !__obf_77cfff95beb3cc99 {
		return
	}
	__obf_1d13df182269d8e8.
		Delete(__obf_da58875072df271b)
	__obf_2dea2b2773c8afdf := __obf_b5900a65e8e7d9b5.(reflect.Type)
	__obf_1d13df182269d8e8.
		Delete(__obf_2dea2b2773c8afdf)
	if __obf_2dea2b2773c8afdf.Kind() == reflect.Ptr {
		__obf_1d13df182269d8e8.
			Delete(__obf_2dea2b2773c8afdf.Elem())
	}
}

func __obf_d01b412f493cff3f(__obf_da58875072df271b int8, __obf_2dea2b2773c8afdf reflect.Type, __obf_2981aff99465546b func(__obf_6a6518f236eeec6e *Encoder, __obf_17732590722140e7 reflect.Value) ([]byte, error),
) __obf_9399ab762e15c2a7 {
	__obf_c3d2a512f69a483f := __obf_2dea2b2773c8afdf.Kind() == reflect.Ptr

	return func(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
		if __obf_c3d2a512f69a483f && __obf_17732590722140e7.IsNil() {
			return __obf_7bae0b613da60dd3.EncodeNil()
		}
		__obf_a7fd7acf2bd4435f, __obf_0feb3528b7b709ec := __obf_2981aff99465546b(__obf_7bae0b613da60dd3, __obf_17732590722140e7)
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}

		if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.EncodeExtHeader(__obf_da58875072df271b, len(__obf_a7fd7acf2bd4435f)); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}

		return __obf_7bae0b613da60dd3.__obf_b0daaf664cd5cdfb(__obf_a7fd7acf2bd4435f)
	}
}

func __obf_3706b5236b1084b3(__obf_f2df1f96c0828a97 __obf_9399ab762e15c2a7) __obf_9399ab762e15c2a7 {
	return func(__obf_7bae0b613da60dd3 *Encoder, __obf_17732590722140e7 reflect.Value) error {
		if !__obf_17732590722140e7.CanAddr() {
			return fmt.Errorf("msgpack: EncodeExt(nonaddressable %T)", __obf_17732590722140e7.Interface())
		}
		return __obf_f2df1f96c0828a97(__obf_7bae0b613da60dd3, __obf_17732590722140e7.Addr())
	}
}

func RegisterExtDecoder(__obf_da58875072df271b int8, __obf_1926a1e373f9e647 any, __obf_0fe2a689830f8985 func(__obf_9461956859459a44 *Decoder, __obf_17732590722140e7 reflect.Value, __obf_eef44dfac5af7de1 int) error,
) {
	__obf_1108802bcba1cc7b(__obf_da58875072df271b)
	__obf_2dea2b2773c8afdf := reflect.TypeOf(__obf_1926a1e373f9e647)
	__obf_9f4f5185b1a0da0a := __obf_347bcbc38377084d(__obf_da58875072df271b, __obf_2dea2b2773c8afdf, __obf_0fe2a689830f8985)
	__obf_11268e7b8aa6c195[__obf_da58875072df271b] = &__obf_55911355a3eaccfd{
		Type:    __obf_2dea2b2773c8afdf,
		Decoder: __obf_0fe2a689830f8985,
	}
	__obf_9f46fbaef88684c0.
		Store(__obf_da58875072df271b, __obf_2dea2b2773c8afdf)
	__obf_9f46fbaef88684c0.
		Store(__obf_2dea2b2773c8afdf, __obf_9f4f5185b1a0da0a)
	if __obf_2dea2b2773c8afdf.Kind() == reflect.Ptr {
		__obf_9f46fbaef88684c0.
			Store(__obf_2dea2b2773c8afdf.Elem(), __obf_c65194f969dc6855(__obf_9f4f5185b1a0da0a))
	}
}

func __obf_1108802bcba1cc7b(__obf_da58875072df271b int8) {
	__obf_b5900a65e8e7d9b5, __obf_77cfff95beb3cc99 := __obf_9f46fbaef88684c0.Load(__obf_da58875072df271b)
	if !__obf_77cfff95beb3cc99 {
		return
	}
	__obf_9f46fbaef88684c0.
		Delete(__obf_da58875072df271b)
	delete(__obf_11268e7b8aa6c195, __obf_da58875072df271b)
	__obf_2dea2b2773c8afdf := __obf_b5900a65e8e7d9b5.(reflect.Type)
	__obf_9f46fbaef88684c0.
		Delete(__obf_2dea2b2773c8afdf)
	if __obf_2dea2b2773c8afdf.Kind() == reflect.Ptr {
		__obf_9f46fbaef88684c0.
			Delete(__obf_2dea2b2773c8afdf.Elem())
	}
}

func __obf_347bcbc38377084d(__obf_aec0415c29846112 int8, __obf_2dea2b2773c8afdf reflect.Type, __obf_0fe2a689830f8985 func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value, __obf_eef44dfac5af7de1 int) error,
) __obf_b74c57495471e351 {
	return __obf_6016fab0ae04eaa0(__obf_2dea2b2773c8afdf, func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
		__obf_da58875072df271b, __obf_eef44dfac5af7de1, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeExtHeader()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		if __obf_da58875072df271b != __obf_aec0415c29846112 {
			return fmt.Errorf("msgpack: got ext type=%d, wanted %d", __obf_da58875072df271b, __obf_aec0415c29846112)
		}
		return __obf_0fe2a689830f8985(__obf_dcaad1165bb07af9, __obf_17732590722140e7, __obf_eef44dfac5af7de1)
	})
}

func __obf_c65194f969dc6855(__obf_9f4f5185b1a0da0a __obf_b74c57495471e351) __obf_b74c57495471e351 {
	return func(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
		if !__obf_17732590722140e7.CanAddr() {
			return fmt.Errorf("msgpack: DecodeExt(nonaddressable %T)", __obf_17732590722140e7.Interface())
		}
		return __obf_9f4f5185b1a0da0a(__obf_dcaad1165bb07af9, __obf_17732590722140e7.Addr())
	}
}

func (__obf_7bae0b613da60dd3 *Encoder) EncodeExtHeader(__obf_da58875072df271b int8, __obf_eef44dfac5af7de1 int) error {
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_0fe969e97a7f7c13(__obf_eef44dfac5af7de1); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_0feb3528b7b709ec := __obf_7bae0b613da60dd3.__obf_007a7bbeb11f001c.WriteByte(byte(__obf_da58875072df271b)); __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return nil
}

func (__obf_7bae0b613da60dd3 *Encoder) __obf_0fe969e97a7f7c13(__obf_6dbc1eb636e18f94 int) error {
	switch __obf_6dbc1eb636e18f94 {
	case 1:
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt1)
	case 2:
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt2)
	case 4:
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt4)
	case 8:
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt8)
	case 16:
		return __obf_7bae0b613da60dd3.__obf_47b86f684323fc33(FixExt16)
	}
	if __obf_6dbc1eb636e18f94 <= math.MaxUint8 {
		return __obf_7bae0b613da60dd3.__obf_fa2b4086fff3f28d(Ext8, uint8(__obf_6dbc1eb636e18f94))
	}
	if __obf_6dbc1eb636e18f94 <= math.MaxUint16 {
		return __obf_7bae0b613da60dd3.__obf_bd6e685495964ed5(Ext16, uint16(__obf_6dbc1eb636e18f94))
	}
	return __obf_7bae0b613da60dd3.__obf_a95f338c34ae207a(Ext32, uint32(__obf_6dbc1eb636e18f94))
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeExtHeader() (__obf_da58875072df271b int8, __obf_eef44dfac5af7de1 int, __obf_0feb3528b7b709ec error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return
	}
	return __obf_dcaad1165bb07af9.__obf_73f4503729af0402(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_73f4503729af0402(__obf_ecf6d2d3253a058d byte) (int8, int, error) {
	__obf_eef44dfac5af7de1, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_a08846e5aacf8d35(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return 0, 0, __obf_0feb3528b7b709ec
	}
	__obf_da58875072df271b, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, 0, __obf_0feb3528b7b709ec
	}

	return int8(__obf_da58875072df271b), __obf_eef44dfac5af7de1, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_a08846e5aacf8d35(__obf_ecf6d2d3253a058d byte) (int, error) {
	switch __obf_ecf6d2d3253a058d {
	case FixExt1:
		return 1, nil
	case FixExt2:
		return 2, nil
	case FixExt4:
		return 4, nil
	case FixExt8:
		return 8, nil
	case FixExt16:
		return 16, nil
	case Ext8:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint8()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Ext16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Ext32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	default:
		return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", __obf_ecf6d2d3253a058d)
	}
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_7b329929aba3053d(__obf_ecf6d2d3253a058d byte) (any, error) {
	__obf_da58875072df271b, __obf_eef44dfac5af7de1, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_73f4503729af0402(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	__obf_0ed0bb469e5eebe7, __obf_77cfff95beb3cc99 := __obf_11268e7b8aa6c195[__obf_da58875072df271b]
	if !__obf_77cfff95beb3cc99 {
		return nil, fmt.Errorf("msgpack: unknown ext id=%d", __obf_da58875072df271b)
	}
	__obf_17732590722140e7 := __obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_0ed0bb469e5eebe7.Type).Elem()
	if __obf_c3d2a512f69a483f(__obf_17732590722140e7.Kind()) && __obf_17732590722140e7.IsNil() {
		__obf_17732590722140e7.
			Set(__obf_dcaad1165bb07af9.__obf_598c5a4b42040fb1(__obf_0ed0bb469e5eebe7.Type.Elem()))
	}

	if __obf_0feb3528b7b709ec := __obf_0ed0bb469e5eebe7.Decoder(__obf_dcaad1165bb07af9, __obf_17732590722140e7, __obf_eef44dfac5af7de1); __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}

	return __obf_17732590722140e7.Interface(), nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_ebb31dcc8229349b(__obf_ecf6d2d3253a058d byte) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_a08846e5aacf8d35(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_6f540045c1336cac(__obf_2b0247e819b1bf4a + 1)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_9019231c4f64c0e5(__obf_ecf6d2d3253a058d byte) error {
	// Read ext type.
	_, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	// Read ext body len.
	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_35d5b7a3d0f72655(__obf_ecf6d2d3253a058d); __obf_101eec84c8338296++ {
		_, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}

func __obf_35d5b7a3d0f72655(__obf_ecf6d2d3253a058d byte) int {
	switch __obf_ecf6d2d3253a058d {
	case Ext8:
		return 1
	case Ext16:
		return 2
	case Ext32:
		return 4
	}
	return 0
}
