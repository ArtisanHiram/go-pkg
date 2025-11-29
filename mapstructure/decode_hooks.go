package __obf_7bd99df3562420c2

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
func __obf_deec805bab5f5d70(__obf_41b5e5e021bb78e7 DecodeHookFunc) DecodeHookFunc {
	// Create variables here so we can reference them with the reflect pkg
	var __obf_4b5ad8a28cd2d262 DecodeHookFuncType
	var __obf_40a13ab9f231410f DecodeHookFuncKind
	var __obf_91de2cf87ebc1d69 DecodeHookFuncValue
	__obf_7d62412f3f087393 := // Fill in the variables into this interface and the rest is done
		// automatically using the reflect package.
		[]any{__obf_4b5ad8a28cd2d262, __obf_40a13ab9f231410f, __obf_91de2cf87ebc1d69}
	__obf_443dc76a6d8afc5e := reflect.ValueOf(__obf_41b5e5e021bb78e7)
	__obf_012d246ff6850eb5 := __obf_443dc76a6d8afc5e.Type()
	for _, __obf_11bc8366b4d80e72 := range __obf_7d62412f3f087393 {
		__obf_f5063245b901f787 := reflect.ValueOf(__obf_11bc8366b4d80e72).Type()
		if __obf_012d246ff6850eb5.ConvertibleTo(__obf_f5063245b901f787) {
			return __obf_443dc76a6d8afc5e.Convert(__obf_f5063245b901f787).Interface()
		}
	}

	return nil
}

// DecodeHookExec executes the given decode hook. This should be used
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
// that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(__obf_11bc8366b4d80e72 DecodeHookFunc, __obf_9d6dd945270ddf52 reflect.Value, __obf_c3010b7e6da97940 reflect.Value) (any, error) {

	switch __obf_2df189427d80bcc8 := __obf_deec805bab5f5d70(__obf_11bc8366b4d80e72).(type) {
	case DecodeHookFuncType:
		return __obf_2df189427d80bcc8(__obf_9d6dd945270ddf52.Type(), __obf_c3010b7e6da97940.Type(), __obf_9d6dd945270ddf52.Interface())
	case DecodeHookFuncKind:
		return __obf_2df189427d80bcc8(__obf_9d6dd945270ddf52.Kind(), __obf_c3010b7e6da97940.Kind(), __obf_9d6dd945270ddf52.Interface())
	case DecodeHookFuncValue:
		return __obf_2df189427d80bcc8(__obf_9d6dd945270ddf52, __obf_c3010b7e6da97940)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
// automatically composes multiple DecodeHookFuncs.
//
// The composed funcs are called in order, with the result of the
// previous transformation.
func ComposeDecodeHookFunc(__obf_a1246ae8097e3cf6 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Value, __obf_0b1634bf673d507e reflect.Value) (any, error) {
		var __obf_bf6c24b903ce000b error
		__obf_57ff3a2a2bd1a861 := __obf_2df189427d80bcc8.Interface()
		__obf_6330a2262b4ca785 := __obf_2df189427d80bcc8
		for _, __obf_4b5ad8a28cd2d262 := range __obf_a1246ae8097e3cf6 {
			__obf_57ff3a2a2bd1a861, __obf_bf6c24b903ce000b = DecodeHookExec(__obf_4b5ad8a28cd2d262, __obf_6330a2262b4ca785, __obf_0b1634bf673d507e)
			if __obf_bf6c24b903ce000b != nil {
				return nil, __obf_bf6c24b903ce000b
			}
			__obf_6330a2262b4ca785 = reflect.ValueOf(__obf_57ff3a2a2bd1a861)
		}

		return __obf_57ff3a2a2bd1a861, nil
	}
}

// OrComposeDecodeHookFunc executes all input hook functions until one of them returns no error. In that case its value is returned.
// If all hooks return an error, OrComposeDecodeHookFunc returns an error concatenating all error messages.
func OrComposeDecodeHookFunc(__obf_37fb11ea80130d5b ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_f801aba5725c4c71, __obf_257c76af251ccba5 reflect.Value) (any, error) {
		var __obf_da5075712b0caca4 string
		var __obf_5472f4e9ba66a825 any
		var __obf_bf6c24b903ce000b error

		for _, __obf_2df189427d80bcc8 := range __obf_37fb11ea80130d5b {
			__obf_5472f4e9ba66a825, __obf_bf6c24b903ce000b = DecodeHookExec(__obf_2df189427d80bcc8, __obf_f801aba5725c4c71, __obf_257c76af251ccba5)
			if __obf_bf6c24b903ce000b != nil {
				__obf_da5075712b0caca4 += __obf_bf6c24b903ce000b.Error() + "\n"
				continue
			}

			return __obf_5472f4e9ba66a825, nil
		}

		return nil, errors.New(__obf_da5075712b0caca4)
	}
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
// string to []string by splitting on the given sep.
func StringToSliceHookFunc(__obf_240f7a1cb0fdb818 string) DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Kind, __obf_0b1634bf673d507e reflect.Kind, __obf_57ff3a2a2bd1a861 any) (any, error) {
		if __obf_2df189427d80bcc8 != reflect.String || __obf_0b1634bf673d507e != reflect.Slice {
			return __obf_57ff3a2a2bd1a861, nil
		}
		__obf_11bc8366b4d80e72 := __obf_57ff3a2a2bd1a861.(string)
		if __obf_11bc8366b4d80e72 == "" {
			return []string{}, nil
		}

		return strings.Split(__obf_11bc8366b4d80e72,

			// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
			// strings to time.Duration.
			__obf_240f7a1cb0fdb818), nil
	}
}

func StringToTimeDurationHookFunc() DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Type, __obf_0b1634bf673d507e reflect.Type, __obf_57ff3a2a2bd1a861 any) (any, error) {
		if __obf_2df189427d80bcc8.Kind() != reflect.String {
			return __obf_57ff3a2a2bd1a861, nil
		}
		if __obf_0b1634bf673d507e != reflect.TypeOf(time.Duration(5)) {
			return __obf_57ff3a2a2bd1a861, nil
		}

		// Convert it by parsing
		return time.ParseDuration(__obf_57ff3a2a2bd1a861.(string))
	}
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
// strings to net.IP
func StringToIPHookFunc() DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Type, __obf_0b1634bf673d507e reflect.Type, __obf_57ff3a2a2bd1a861 any) (any, error) {
		if __obf_2df189427d80bcc8.Kind() != reflect.String {
			return __obf_57ff3a2a2bd1a861, nil
		}
		if __obf_0b1634bf673d507e != reflect.TypeOf(net.IP{}) {
			return __obf_57ff3a2a2bd1a861, nil
		}
		__obf_7afef972d9dbb211 := // Convert it by parsing
			net.ParseIP(__obf_57ff3a2a2bd1a861.(string))
		if __obf_7afef972d9dbb211 == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", __obf_57ff3a2a2bd1a861)
		}

		return __obf_7afef972d9dbb211, nil
	}
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
// strings to net.IPNet
func StringToIPNetHookFunc() DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Type, __obf_0b1634bf673d507e reflect.Type, __obf_57ff3a2a2bd1a861 any) (any, error) {
		if __obf_2df189427d80bcc8.Kind() != reflect.String {
			return __obf_57ff3a2a2bd1a861, nil
		}
		if __obf_0b1634bf673d507e != reflect.TypeOf(net.IPNet{}) {
			return __obf_57ff3a2a2bd1a861, nil
		}

		// Convert it by parsing
		_, net, __obf_bf6c24b903ce000b := net.ParseCIDR(__obf_57ff3a2a2bd1a861.(string))
		return net, __obf_bf6c24b903ce000b
	}
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func StringToTimeHookFunc(__obf_c43914ef62e70538 string) DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Type, __obf_0b1634bf673d507e reflect.Type, __obf_57ff3a2a2bd1a861 any) (any, error) {
		if __obf_2df189427d80bcc8.Kind() != reflect.String {
			return __obf_57ff3a2a2bd1a861, nil
		}
		if __obf_0b1634bf673d507e != reflect.TypeOf(time.Time{}) {
			return __obf_57ff3a2a2bd1a861, nil
		}

		// Convert it by parsing
		return time.Parse(__obf_c43914ef62e70538, __obf_57ff3a2a2bd1a861.(string))
	}
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
// the decoder.
//
// Note that this is significantly different from the WeaklyTypedInput option
// of the DecoderConfig.
func WeaklyTypedHook(__obf_2df189427d80bcc8 reflect.Kind, __obf_0b1634bf673d507e reflect.Kind, __obf_57ff3a2a2bd1a861 any) (any, error) {
	__obf_d633e91150cb7889 := reflect.ValueOf(__obf_57ff3a2a2bd1a861)
	switch __obf_0b1634bf673d507e {
	case reflect.String:
		switch __obf_2df189427d80bcc8 {
		case reflect.Bool:
			if __obf_d633e91150cb7889.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(__obf_d633e91150cb7889.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(__obf_d633e91150cb7889.Int(), 10), nil
		case reflect.Slice:
			__obf_4ab72a4f9eed38d8 := __obf_d633e91150cb7889.Type()
			__obf_9dd35fb0b81a6a3a := __obf_4ab72a4f9eed38d8.Elem().Kind()
			if __obf_9dd35fb0b81a6a3a == reflect.Uint8 {
				return string(__obf_d633e91150cb7889.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(__obf_d633e91150cb7889.Uint(), 10), nil
		}
	}

	return __obf_57ff3a2a2bd1a861, nil
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
	return func(__obf_2df189427d80bcc8 reflect.Value, __obf_0b1634bf673d507e reflect.Value) (any, error) {
		if __obf_2df189427d80bcc8.Kind() != reflect.Struct {
			return __obf_2df189427d80bcc8.Interface(), nil
		}

		var __obf_8d0ae84c7a3f3f55 any = struct{}{}
		if __obf_0b1634bf673d507e.Type() != reflect.TypeOf(&__obf_8d0ae84c7a3f3f55).Elem() {
			return __obf_2df189427d80bcc8.Interface(), nil
		}
		__obf_d8bef46104947b25 := make(map[string]any)
		__obf_0b1634bf673d507e.
			Set(reflect.ValueOf(__obf_d8bef46104947b25))

		return __obf_2df189427d80bcc8.Interface(), nil
	}
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshallerHookFunc() DecodeHookFuncType {
	return func(__obf_2df189427d80bcc8 reflect.Type, __obf_0b1634bf673d507e reflect.Type, __obf_57ff3a2a2bd1a861 any) (any, error) {
		if __obf_2df189427d80bcc8.Kind() != reflect.String {
			return __obf_57ff3a2a2bd1a861, nil
		}
		__obf_6ad3d371cac0efc2 := reflect.New(__obf_0b1634bf673d507e).Interface()
		__obf_5b504138dd3ef47c, __obf_5da7816f07b20ebf := __obf_6ad3d371cac0efc2.(encoding.TextUnmarshaler)
		if !__obf_5da7816f07b20ebf {
			return __obf_57ff3a2a2bd1a861, nil
		}
		__obf_1264b3f46bd409c7, __obf_5da7816f07b20ebf := __obf_57ff3a2a2bd1a861.(string)
		if !__obf_5da7816f07b20ebf {
			__obf_1264b3f46bd409c7 = reflect.Indirect(reflect.ValueOf(&__obf_57ff3a2a2bd1a861)).Elem().String()
		}
		if __obf_bf6c24b903ce000b := __obf_5b504138dd3ef47c.UnmarshalText([]byte(__obf_1264b3f46bd409c7)); __obf_bf6c24b903ce000b != nil {
			return nil, __obf_bf6c24b903ce000b
		}
		return __obf_6ad3d371cac0efc2, nil
	}
}
