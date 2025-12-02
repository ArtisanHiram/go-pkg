package __obf_a935eb7f548271a4

import (
	"encoding"
	"fmt"
	"reflect"
)

var __obf_2b6c97801680295d []__obf_4982458b85add976

//nolint:gochecknoinits
func init() {
	__obf_2b6c97801680295d = []__obf_4982458b85add976{
		reflect.Bool:          __obf_3e6a423b617bedc6,
		reflect.Int:           __obf_53328fa97f4370bb,
		reflect.Int8:          __obf_17e8f11c01632ce7,
		reflect.Int16:         __obf_aee8774824bbd873,
		reflect.Int32:         __obf_4c356e77af1b6a6b,
		reflect.Int64:         __obf_df3286b177337d63,
		reflect.Uint:          __obf_26b7f97739101864,
		reflect.Uint8:         __obf_2694089cc89b620b,
		reflect.Uint16:        __obf_64119742b2dcdc14,
		reflect.Uint32:        __obf_f89278ad27f637c0,
		reflect.Uint64:        __obf_2d4225df72c2ec23,
		reflect.Float32:       __obf_f67e8b412fa8016c,
		reflect.Float64:       __obf_8f648744c64b0585,
		reflect.Complex64:     __obf_7ee0487f979e6e76,
		reflect.Complex128:    __obf_7ee0487f979e6e76,
		reflect.Array:         __obf_8b865df2f8afbfbe,
		reflect.Chan:          __obf_7ee0487f979e6e76,
		reflect.Func:          __obf_7ee0487f979e6e76,
		reflect.Interface:     __obf_17c40b88235b8935,
		reflect.Map:           __obf_a743877260781a87,
		reflect.Ptr:           __obf_7ee0487f979e6e76,
		reflect.Slice:         __obf_e4145022d977f5f9,
		reflect.String:        __obf_96718da505bbd637,
		reflect.Struct:        __obf_1d06e69b22f89920,
		reflect.UnsafePointer: __obf_7ee0487f979e6e76,
	}
}

func __obf_d55d857e358cbb2a(__obf_7bd64e8df8357cab reflect.Type) __obf_4982458b85add976 {
	if __obf_6d570581f4b60dbc, __obf_826ac2dbb957d7df := __obf_a16fc4f2eeccec6d.Load(__obf_7bd64e8df8357cab); __obf_826ac2dbb957d7df {
		return __obf_6d570581f4b60dbc.(__obf_4982458b85add976)
	}
	__obf_36c78c696c47cfdb := _getEncoder(__obf_7bd64e8df8357cab)
	__obf_a16fc4f2eeccec6d.
		Store(__obf_7bd64e8df8357cab, __obf_36c78c696c47cfdb)
	return __obf_36c78c696c47cfdb
}

func _getEncoder(__obf_7bd64e8df8357cab reflect.Type) __obf_4982458b85add976 {
	__obf_38e9ce3a1f9f7821 := __obf_7bd64e8df8357cab.Kind()

	if __obf_38e9ce3a1f9f7821 == reflect.Pointer {
		if _, __obf_826ac2dbb957d7df := __obf_a16fc4f2eeccec6d.Load(__obf_7bd64e8df8357cab.Elem()); __obf_826ac2dbb957d7df {
			return __obf_275be5bf328de671(__obf_7bd64e8df8357cab)
		}
	}

	if __obf_7bd64e8df8357cab.Implements(__obf_0f9b3527fa1e5883) {
		return __obf_5eb05342a6dfc56d
	}
	if __obf_7bd64e8df8357cab.Implements(__obf_626f6da86a13b183) {
		return __obf_c553b24b5ba41d93
	}
	if __obf_7bd64e8df8357cab.Implements(__obf_17933e997378c10f) {
		return __obf_2602ac20859b1214
	}
	if __obf_7bd64e8df8357cab.Implements(__obf_499c27ef04771594) {
		return __obf_ca0d3dafae5f64b7
	}

	// Addressable struct field value.
	if __obf_38e9ce3a1f9f7821 != reflect.Pointer {
		__obf_0d8a994785cda6df := reflect.PointerTo(__obf_7bd64e8df8357cab)
		if __obf_0d8a994785cda6df.Implements(__obf_0f9b3527fa1e5883) {
			return __obf_0799b9a8d931de08
		}
		if __obf_0d8a994785cda6df.Implements(__obf_626f6da86a13b183) {
			return __obf_f85fe1712c68105b
		}
		if __obf_0d8a994785cda6df.Implements(__obf_17933e997378c10f) {
			return __obf_844dcf27bfd0f057
		}
		if __obf_0d8a994785cda6df.Implements(__obf_499c27ef04771594) {
			return __obf_f27f151adcf78d30
		}
	}

	if __obf_7bd64e8df8357cab == __obf_4ab64edf794ae331 {
		return __obf_ce9987bb8fab7197
	}

	switch __obf_38e9ce3a1f9f7821 {
	case reflect.Ptr:
		return __obf_275be5bf328de671(__obf_7bd64e8df8357cab)
	case reflect.Slice:
		__obf_f50dd3cfb2ad5435 := __obf_7bd64e8df8357cab.Elem()
		if __obf_f50dd3cfb2ad5435.Kind() == reflect.Uint8 {
			return __obf_892a8b5216486b44
		}
		if __obf_f50dd3cfb2ad5435 == __obf_b128fe7c7d10a668 {
			return __obf_cea67acf0689ac24
		}
	case reflect.Array:
		if __obf_7bd64e8df8357cab.Elem().Kind() == reflect.Uint8 {
			return __obf_ff83fd93faa9eb52
		}
	case reflect.Map:
		if __obf_7bd64e8df8357cab.Key() == __obf_b128fe7c7d10a668 {
			switch __obf_7bd64e8df8357cab.Elem() {
			case __obf_b128fe7c7d10a668:
				return __obf_e9d97c6d9e0e5b5f
			case __obf_8887b1926a35f0b7:
				return __obf_19525146631c9c00
			case __obf_24d1b614db17eb2b:
				return __obf_25ee786fca8b560b
			}
		}
	}

	return __obf_2b6c97801680295d[__obf_38e9ce3a1f9f7821]
}

func __obf_275be5bf328de671(__obf_7bd64e8df8357cab reflect.Type) __obf_4982458b85add976 {
	__obf_6a114719334e34d2 := __obf_d55d857e358cbb2a(__obf_7bd64e8df8357cab.Elem())
	return func(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
		if __obf_6d570581f4b60dbc.IsNil() {
			return __obf_ed37a34c347049f3.EncodeNil()
		}
		return __obf_6a114719334e34d2(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc.Elem())
	}
}

func __obf_0799b9a8d931de08(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if !__obf_6d570581f4b60dbc.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_6d570581f4b60dbc.Interface())
	}
	__obf_6a114719334e34d2 := __obf_6d570581f4b60dbc.Addr().Interface().(CustomEncoder)
	return __obf_6a114719334e34d2.EncodeMsgpack(__obf_ed37a34c347049f3)
}

func __obf_5eb05342a6dfc56d(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_8b51c57d6f65dcf5(__obf_6d570581f4b60dbc.Kind()) && __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	__obf_6a114719334e34d2 := __obf_6d570581f4b60dbc.Interface().(CustomEncoder)
	return __obf_6a114719334e34d2.EncodeMsgpack(__obf_ed37a34c347049f3)
}

func __obf_f85fe1712c68105b(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if !__obf_6d570581f4b60dbc.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_6d570581f4b60dbc.Interface())
	}
	return __obf_c553b24b5ba41d93(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc.Addr())
}

func __obf_c553b24b5ba41d93(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_8b51c57d6f65dcf5(__obf_6d570581f4b60dbc.Kind()) && __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	__obf_ccececcb28726e85 := __obf_6d570581f4b60dbc.Interface().(Marshaler)
	__obf_f2ca794293605b73, __obf_4d327e1cd40c2e21 := __obf_ccececcb28726e85.MarshalMsgpack()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	_, __obf_4d327e1cd40c2e21 = __obf_ed37a34c347049f3.__obf_a3adbf62c8801563.Write(__obf_f2ca794293605b73)
	return __obf_4d327e1cd40c2e21
}

func __obf_3e6a423b617bedc6(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return __obf_ed37a34c347049f3.EncodeBool(__obf_6d570581f4b60dbc.Bool())
}

func __obf_17c40b88235b8935(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	return __obf_ed37a34c347049f3.EncodeValue(__obf_6d570581f4b60dbc.Elem())
}

func __obf_ce9987bb8fab7197(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	return __obf_ed37a34c347049f3.EncodeString(__obf_6d570581f4b60dbc.Interface().(error).Error())
}

func __obf_7ee0487f979e6e76(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	return fmt.Errorf("msgpack: Encode(unsupported %s)", __obf_6d570581f4b60dbc.Type())
}

func __obf_8b51c57d6f65dcf5(__obf_38e9ce3a1f9f7821 reflect.Kind) bool {
	switch __obf_38e9ce3a1f9f7821 {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return true
	}
	return false
}

func __obf_2def3fc626a43dac(__obf_14c99752d39cbe18 reflect.Type) bool {
	if __obf_14c99752d39cbe18.Kind() == reflect.Ptr {
		__obf_14c99752d39cbe18 = __obf_14c99752d39cbe18.Elem()
	}
	return __obf_8b51c57d6f65dcf5(__obf_14c99752d39cbe18.Kind())
}

//------------------------------------------------------------------------------

func __obf_844dcf27bfd0f057(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if !__obf_6d570581f4b60dbc.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_6d570581f4b60dbc.Interface())
	}
	return __obf_2602ac20859b1214(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc.Addr())
}

func __obf_2602ac20859b1214(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_8b51c57d6f65dcf5(__obf_6d570581f4b60dbc.Kind()) && __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	__obf_ccececcb28726e85 := __obf_6d570581f4b60dbc.Interface().(encoding.BinaryMarshaler)
	__obf_652c67787b9c24c9, __obf_4d327e1cd40c2e21 := __obf_ccececcb28726e85.MarshalBinary()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	return __obf_ed37a34c347049f3.EncodeBytes(__obf_652c67787b9c24c9)
}

//------------------------------------------------------------------------------

func __obf_f27f151adcf78d30(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if !__obf_6d570581f4b60dbc.CanAddr() {
		return fmt.Errorf("msgpack: Encode(non-addressable %T)", __obf_6d570581f4b60dbc.Interface())
	}
	return __obf_ca0d3dafae5f64b7(__obf_ed37a34c347049f3, __obf_6d570581f4b60dbc.Addr())
}

func __obf_ca0d3dafae5f64b7(__obf_ed37a34c347049f3 *Encoder, __obf_6d570581f4b60dbc reflect.Value) error {
	if __obf_8b51c57d6f65dcf5(__obf_6d570581f4b60dbc.Kind()) && __obf_6d570581f4b60dbc.IsNil() {
		return __obf_ed37a34c347049f3.EncodeNil()
	}
	__obf_ccececcb28726e85 := __obf_6d570581f4b60dbc.Interface().(encoding.TextMarshaler)
	__obf_652c67787b9c24c9, __obf_4d327e1cd40c2e21 := __obf_ccececcb28726e85.MarshalText()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	return __obf_ed37a34c347049f3.EncodeBytes(__obf_652c67787b9c24c9)
}
