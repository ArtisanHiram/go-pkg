package __obf_68a92d5117d8d921

import (
	"encoding"
	"errors"
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// typedDecodeHook takes a raw DecodeHookFunc (an any) and turns
// it into the proper DecodeHookFunc type, such as DecodeHookFuncType.
func __obf_375cb77df355cd0b(__obf_4c2820bceec8e875 DecodeHookFunc) DecodeHookFunc {
	// Create variables here so we can reference them with the reflect pkg
	var __obf_7e431bd0fa30b679 DecodeHookFuncType
	var __obf_25179ae1c113550a DecodeHookFuncKind
	var __obf_60ceded5841bc8b7 DecodeHookFuncValue
	__obf_e9fc43f4eabcd5e9 := // Fill in the variables into this interface and the rest is done
		// automatically using the reflect package.
		[]any{__obf_7e431bd0fa30b679, __obf_25179ae1c113550a, __obf_60ceded5841bc8b7}
	__obf_7dd7d18cd783c84c := reflect.ValueOf(__obf_4c2820bceec8e875)
	__obf_8d73fdf8222cf919 := __obf_7dd7d18cd783c84c.Type()
	for _, __obf_0bdbba89bdca2257 := range __obf_e9fc43f4eabcd5e9 {
		__obf_5d5e9895bceeb853 := reflect.ValueOf(__obf_0bdbba89bdca2257).Type()
		if __obf_8d73fdf8222cf919.ConvertibleTo(__obf_5d5e9895bceeb853) {
			return __obf_7dd7d18cd783c84c.Convert(__obf_5d5e9895bceeb853).Interface()
		}
	}

	return nil
}

// DecodeHookExec executes the given decode hook. This should be used
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
// that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(__obf_0bdbba89bdca2257 DecodeHookFunc, __obf_793e64bff3015f1d reflect.Value, __obf_f4678c4e7943a8ce reflect.Value) (any, error) {

	switch __obf_9037ed2181de9473 := __obf_375cb77df355cd0b(__obf_0bdbba89bdca2257).(type) {
	case DecodeHookFuncType:
		return __obf_9037ed2181de9473(__obf_793e64bff3015f1d.Type(), __obf_f4678c4e7943a8ce.Type(), __obf_793e64bff3015f1d.Interface())
	case DecodeHookFuncKind:
		return __obf_9037ed2181de9473(__obf_793e64bff3015f1d.Kind(), __obf_f4678c4e7943a8ce.Kind(), __obf_793e64bff3015f1d.Interface())
	case DecodeHookFuncValue:
		return __obf_9037ed2181de9473(__obf_793e64bff3015f1d, __obf_f4678c4e7943a8ce)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
// automatically composes multiple DecodeHookFuncs.
//
// The composed funcs are called in order, with the result of the
// previous transformation.
func ComposeDecodeHookFunc(__obf_bafaa6b868d2ab47 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Value, __obf_7d7cb737fce4bf63 reflect.Value) (any, error) {
		var __obf_ef25cdf0d96b47a8 error
		__obf_d6e7451ec5237c6d := __obf_9037ed2181de9473.Interface()
		__obf_683d2eeb129dbc1e := __obf_9037ed2181de9473
		for _, __obf_7e431bd0fa30b679 := range __obf_bafaa6b868d2ab47 {
			__obf_d6e7451ec5237c6d, __obf_ef25cdf0d96b47a8 = DecodeHookExec(__obf_7e431bd0fa30b679, __obf_683d2eeb129dbc1e, __obf_7d7cb737fce4bf63)
			if __obf_ef25cdf0d96b47a8 != nil {
				return nil, __obf_ef25cdf0d96b47a8
			}
			__obf_683d2eeb129dbc1e = reflect.ValueOf(__obf_d6e7451ec5237c6d)
		}

		return __obf_d6e7451ec5237c6d, nil
	}
}

// OrComposeDecodeHookFunc executes all input hook functions until one of them returns no error. In that case its value is returned.
// If all hooks return an error, OrComposeDecodeHookFunc returns an error concatenating all error messages.
func OrComposeDecodeHookFunc(__obf_13acdfed7f1930fb ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_56c43809620ffe14, __obf_dac936e7449e698a reflect.Value) (any, error) {
		var __obf_e51a59a55ab88993 string
		var __obf_920089fb9446f466 any
		var __obf_ef25cdf0d96b47a8 error

		for _, __obf_9037ed2181de9473 := range __obf_13acdfed7f1930fb {
			__obf_920089fb9446f466, __obf_ef25cdf0d96b47a8 = DecodeHookExec(__obf_9037ed2181de9473, __obf_56c43809620ffe14, __obf_dac936e7449e698a)
			if __obf_ef25cdf0d96b47a8 != nil {
				__obf_e51a59a55ab88993 += __obf_ef25cdf0d96b47a8.Error() + "\n"
				continue
			}

			return __obf_920089fb9446f466, nil
		}

		return nil, errors.New(__obf_e51a59a55ab88993)
	}
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
// string to []string by splitting on the given sep.
func StringToSliceHookFunc(__obf_b6231a8fbd759924 string) DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Kind, __obf_7d7cb737fce4bf63 reflect.Kind, __obf_d6e7451ec5237c6d any) (any, error) {
		if __obf_9037ed2181de9473 != reflect.String || __obf_7d7cb737fce4bf63 != reflect.Slice {
			return __obf_d6e7451ec5237c6d, nil
		}
		__obf_0bdbba89bdca2257 := __obf_d6e7451ec5237c6d.(string)
		if __obf_0bdbba89bdca2257 == "" {
			return []string{}, nil
		}

		return strings.Split(__obf_0bdbba89bdca2257,

			// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
			// strings to time.Duration.
			__obf_b6231a8fbd759924), nil
	}
}

func StringToTimeDurationHookFunc() DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Type, __obf_7d7cb737fce4bf63 reflect.Type, __obf_d6e7451ec5237c6d any) (any, error) {
		if __obf_9037ed2181de9473.Kind() != reflect.String {
			return __obf_d6e7451ec5237c6d, nil
		}
		if __obf_7d7cb737fce4bf63 != reflect.TypeOf(time.Duration(5)) {
			return __obf_d6e7451ec5237c6d, nil
		}

		// Convert it by parsing
		return time.ParseDuration(__obf_d6e7451ec5237c6d.(string))
	}
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
// strings to net.IP
func StringToIPHookFunc() DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Type, __obf_7d7cb737fce4bf63 reflect.Type, __obf_d6e7451ec5237c6d any) (any, error) {
		if __obf_9037ed2181de9473.Kind() != reflect.String {
			return __obf_d6e7451ec5237c6d, nil
		}
		if __obf_7d7cb737fce4bf63 != reflect.TypeOf(net.IP{}) {
			return __obf_d6e7451ec5237c6d, nil
		}
		__obf_c596c67338d39956 := // Convert it by parsing
			net.ParseIP(__obf_d6e7451ec5237c6d.(string))
		if __obf_c596c67338d39956 == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", __obf_d6e7451ec5237c6d)
		}

		return __obf_c596c67338d39956, nil
	}
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
// strings to net.IPNet
func StringToIPNetHookFunc() DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Type, __obf_7d7cb737fce4bf63 reflect.Type, __obf_d6e7451ec5237c6d any) (any, error) {
		if __obf_9037ed2181de9473.Kind() != reflect.String {
			return __obf_d6e7451ec5237c6d, nil
		}
		if __obf_7d7cb737fce4bf63 != reflect.TypeOf(net.IPNet{}) {
			return __obf_d6e7451ec5237c6d, nil
		}

		// Convert it by parsing
		_, net, __obf_ef25cdf0d96b47a8 := net.ParseCIDR(__obf_d6e7451ec5237c6d.(string))
		return net, __obf_ef25cdf0d96b47a8
	}
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func StringToTimeHookFunc(__obf_8f266809bcd4f884 string) DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Type, __obf_7d7cb737fce4bf63 reflect.Type, __obf_d6e7451ec5237c6d any) (any, error) {
		if __obf_9037ed2181de9473.Kind() != reflect.String {
			return __obf_d6e7451ec5237c6d, nil
		}
		if __obf_7d7cb737fce4bf63 != reflect.TypeOf(time.Time{}) {
			return __obf_d6e7451ec5237c6d, nil
		}

		// Convert it by parsing
		return time.Parse(__obf_8f266809bcd4f884, __obf_d6e7451ec5237c6d.(string))
	}
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
// the decoder.
//
// Note that this is significantly different from the WeaklyTypedInput option
// of the DecoderConfig.
func WeaklyTypedHook(__obf_9037ed2181de9473 reflect.Kind, __obf_7d7cb737fce4bf63 reflect.Kind, __obf_d6e7451ec5237c6d any) (any, error) {
	__obf_6001932d600cd7ec := reflect.ValueOf(__obf_d6e7451ec5237c6d)
	switch __obf_7d7cb737fce4bf63 {
	case reflect.String:
		switch __obf_9037ed2181de9473 {
		case reflect.Bool:
			if __obf_6001932d600cd7ec.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(__obf_6001932d600cd7ec.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(__obf_6001932d600cd7ec.Int(), 10), nil
		case reflect.Slice:
			__obf_2fec4dc52e9242d2 := __obf_6001932d600cd7ec.Type()
			__obf_f1bf6756e05f60e0 := __obf_2fec4dc52e9242d2.Elem().Kind()
			if __obf_f1bf6756e05f60e0 == reflect.Uint8 {
				return string(__obf_6001932d600cd7ec.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(__obf_6001932d600cd7ec.Uint(), 10), nil
		}
	}

	return __obf_d6e7451ec5237c6d, nil
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
	return func(__obf_9037ed2181de9473 reflect.Value, __obf_7d7cb737fce4bf63 reflect.Value) (any, error) {
		if __obf_9037ed2181de9473.Kind() != reflect.Struct {
			return __obf_9037ed2181de9473.Interface(), nil
		}

		var __obf_a6629621259bc2bc any = struct{}{}
		if __obf_7d7cb737fce4bf63.Type() != reflect.TypeOf(&__obf_a6629621259bc2bc).Elem() {
			return __obf_9037ed2181de9473.Interface(), nil
		}
		__obf_74f8c805bc7061a1 := make(map[string]any)
		__obf_7d7cb737fce4bf63.
			Set(reflect.ValueOf(__obf_74f8c805bc7061a1))

		return __obf_9037ed2181de9473.Interface(), nil
	}
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshallerHookFunc() DecodeHookFuncType {
	return func(__obf_9037ed2181de9473 reflect.Type, __obf_7d7cb737fce4bf63 reflect.Type, __obf_d6e7451ec5237c6d any) (any, error) {
		if __obf_9037ed2181de9473.Kind() != reflect.String {
			return __obf_d6e7451ec5237c6d, nil
		}
		__obf_84b0717583bd55a4 := reflect.New(__obf_7d7cb737fce4bf63).Interface()
		__obf_502c7394280c73e1, __obf_b45c7b8adb1c2d97 := __obf_84b0717583bd55a4.(encoding.TextUnmarshaler)
		if !__obf_b45c7b8adb1c2d97 {
			return __obf_d6e7451ec5237c6d, nil
		}
		__obf_415fa9df6fadc02e, __obf_b45c7b8adb1c2d97 := __obf_d6e7451ec5237c6d.(string)
		if !__obf_b45c7b8adb1c2d97 {
			__obf_415fa9df6fadc02e = reflect.Indirect(reflect.ValueOf(&__obf_d6e7451ec5237c6d)).Elem().String()
		}
		if __obf_ef25cdf0d96b47a8 := __obf_502c7394280c73e1.UnmarshalText([]byte(__obf_415fa9df6fadc02e)); __obf_ef25cdf0d96b47a8 != nil {
			return nil, __obf_ef25cdf0d96b47a8
		}
		return __obf_84b0717583bd55a4, nil
	}
}
