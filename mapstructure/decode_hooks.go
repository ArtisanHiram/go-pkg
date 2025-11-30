package __obf_4dc3483102e0d35a

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
func __obf_fd7b588b91edc0d0(__obf_07db53e7a011f227 DecodeHookFunc) DecodeHookFunc {
	// Create variables here so we can reference them with the reflect pkg
	var __obf_05d4af4162fa054e DecodeHookFuncType
	var __obf_fdae45937aef497a DecodeHookFuncKind
	var __obf_eda60f1ea25145ae DecodeHookFuncValue
	__obf_1884ced3d7ad7754 := // Fill in the variables into this interface and the rest is done
		// automatically using the reflect package.
		[]any{__obf_05d4af4162fa054e, __obf_fdae45937aef497a, __obf_eda60f1ea25145ae}
	__obf_3b55daf4bc34d8ef := reflect.ValueOf(__obf_07db53e7a011f227)
	__obf_0d71898b8d3e8a0b := __obf_3b55daf4bc34d8ef.Type()
	for _, __obf_fcf2168292487a60 := range __obf_1884ced3d7ad7754 {
		__obf_f52efc01595e8921 := reflect.ValueOf(__obf_fcf2168292487a60).Type()
		if __obf_0d71898b8d3e8a0b.ConvertibleTo(__obf_f52efc01595e8921) {
			return __obf_3b55daf4bc34d8ef.Convert(__obf_f52efc01595e8921).Interface()
		}
	}

	return nil
}

// DecodeHookExec executes the given decode hook. This should be used
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
// that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(__obf_fcf2168292487a60 DecodeHookFunc, __obf_73039542a851b768 reflect.Value, __obf_f41e135eac34c3c2 reflect.Value) (any, error) {

	switch __obf_4cd911a2d7099b07 := __obf_fd7b588b91edc0d0(__obf_fcf2168292487a60).(type) {
	case DecodeHookFuncType:
		return __obf_4cd911a2d7099b07(__obf_73039542a851b768.Type(), __obf_f41e135eac34c3c2.Type(), __obf_73039542a851b768.Interface())
	case DecodeHookFuncKind:
		return __obf_4cd911a2d7099b07(__obf_73039542a851b768.Kind(), __obf_f41e135eac34c3c2.Kind(), __obf_73039542a851b768.Interface())
	case DecodeHookFuncValue:
		return __obf_4cd911a2d7099b07(__obf_73039542a851b768, __obf_f41e135eac34c3c2)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
// automatically composes multiple DecodeHookFuncs.
//
// The composed funcs are called in order, with the result of the
// previous transformation.
func ComposeDecodeHookFunc(__obf_18432bd1e1b827f2 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Value, __obf_7af45612bc672ea8 reflect.Value) (any, error) {
		var __obf_6585aef4313e6005 error
		__obf_333b2acf09d806b3 := __obf_4cd911a2d7099b07.Interface()
		__obf_ced4fcb24c95d637 := __obf_4cd911a2d7099b07
		for _, __obf_05d4af4162fa054e := range __obf_18432bd1e1b827f2 {
			__obf_333b2acf09d806b3, __obf_6585aef4313e6005 = DecodeHookExec(__obf_05d4af4162fa054e, __obf_ced4fcb24c95d637, __obf_7af45612bc672ea8)
			if __obf_6585aef4313e6005 != nil {
				return nil, __obf_6585aef4313e6005
			}
			__obf_ced4fcb24c95d637 = reflect.ValueOf(__obf_333b2acf09d806b3)
		}

		return __obf_333b2acf09d806b3, nil
	}
}

// OrComposeDecodeHookFunc executes all input hook functions until one of them returns no error. In that case its value is returned.
// If all hooks return an error, OrComposeDecodeHookFunc returns an error concatenating all error messages.
func OrComposeDecodeHookFunc(__obf_0f737e41f3e2c006 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_2faf15f1900dd830, __obf_91f8fc20b44aae81 reflect.Value) (any, error) {
		var __obf_843fa9b12d9dfef0 string
		var __obf_a16028f63137427a any
		var __obf_6585aef4313e6005 error

		for _, __obf_4cd911a2d7099b07 := range __obf_0f737e41f3e2c006 {
			__obf_a16028f63137427a, __obf_6585aef4313e6005 = DecodeHookExec(__obf_4cd911a2d7099b07, __obf_2faf15f1900dd830, __obf_91f8fc20b44aae81)
			if __obf_6585aef4313e6005 != nil {
				__obf_843fa9b12d9dfef0 += __obf_6585aef4313e6005.Error() + "\n"
				continue
			}

			return __obf_a16028f63137427a, nil
		}

		return nil, errors.New(__obf_843fa9b12d9dfef0)
	}
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
// string to []string by splitting on the given sep.
func StringToSliceHookFunc(__obf_ff660de53bb79102 string) DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Kind, __obf_7af45612bc672ea8 reflect.Kind, __obf_333b2acf09d806b3 any) (any, error) {
		if __obf_4cd911a2d7099b07 != reflect.String || __obf_7af45612bc672ea8 != reflect.Slice {
			return __obf_333b2acf09d806b3, nil
		}
		__obf_fcf2168292487a60 := __obf_333b2acf09d806b3.(string)
		if __obf_fcf2168292487a60 == "" {
			return []string{}, nil
		}

		return strings.Split(__obf_fcf2168292487a60,

			// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
			// strings to time.Duration.
			__obf_ff660de53bb79102), nil
	}
}

func StringToTimeDurationHookFunc() DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Type, __obf_7af45612bc672ea8 reflect.Type, __obf_333b2acf09d806b3 any) (any, error) {
		if __obf_4cd911a2d7099b07.Kind() != reflect.String {
			return __obf_333b2acf09d806b3, nil
		}
		if __obf_7af45612bc672ea8 != reflect.TypeOf(time.Duration(5)) {
			return __obf_333b2acf09d806b3, nil
		}

		// Convert it by parsing
		return time.ParseDuration(__obf_333b2acf09d806b3.(string))
	}
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
// strings to net.IP
func StringToIPHookFunc() DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Type, __obf_7af45612bc672ea8 reflect.Type, __obf_333b2acf09d806b3 any) (any, error) {
		if __obf_4cd911a2d7099b07.Kind() != reflect.String {
			return __obf_333b2acf09d806b3, nil
		}
		if __obf_7af45612bc672ea8 != reflect.TypeOf(net.IP{}) {
			return __obf_333b2acf09d806b3, nil
		}
		__obf_4e9a339c860f41c9 := // Convert it by parsing
			net.ParseIP(__obf_333b2acf09d806b3.(string))
		if __obf_4e9a339c860f41c9 == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", __obf_333b2acf09d806b3)
		}

		return __obf_4e9a339c860f41c9, nil
	}
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
// strings to net.IPNet
func StringToIPNetHookFunc() DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Type, __obf_7af45612bc672ea8 reflect.Type, __obf_333b2acf09d806b3 any) (any, error) {
		if __obf_4cd911a2d7099b07.Kind() != reflect.String {
			return __obf_333b2acf09d806b3, nil
		}
		if __obf_7af45612bc672ea8 != reflect.TypeOf(net.IPNet{}) {
			return __obf_333b2acf09d806b3, nil
		}

		// Convert it by parsing
		_, net, __obf_6585aef4313e6005 := net.ParseCIDR(__obf_333b2acf09d806b3.(string))
		return net, __obf_6585aef4313e6005
	}
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func StringToTimeHookFunc(__obf_552b8b47a8fdc36c string) DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Type, __obf_7af45612bc672ea8 reflect.Type, __obf_333b2acf09d806b3 any) (any, error) {
		if __obf_4cd911a2d7099b07.Kind() != reflect.String {
			return __obf_333b2acf09d806b3, nil
		}
		if __obf_7af45612bc672ea8 != reflect.TypeOf(time.Time{}) {
			return __obf_333b2acf09d806b3, nil
		}

		// Convert it by parsing
		return time.Parse(__obf_552b8b47a8fdc36c, __obf_333b2acf09d806b3.(string))
	}
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
// the decoder.
//
// Note that this is significantly different from the WeaklyTypedInput option
// of the DecoderConfig.
func WeaklyTypedHook(__obf_4cd911a2d7099b07 reflect.Kind, __obf_7af45612bc672ea8 reflect.Kind, __obf_333b2acf09d806b3 any) (any, error) {
	__obf_a4fef5c9586c1966 := reflect.ValueOf(__obf_333b2acf09d806b3)
	switch __obf_7af45612bc672ea8 {
	case reflect.String:
		switch __obf_4cd911a2d7099b07 {
		case reflect.Bool:
			if __obf_a4fef5c9586c1966.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(__obf_a4fef5c9586c1966.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(__obf_a4fef5c9586c1966.Int(), 10), nil
		case reflect.Slice:
			__obf_f73efdd6543886c9 := __obf_a4fef5c9586c1966.Type()
			__obf_9bbf49c8c89fbd8a := __obf_f73efdd6543886c9.Elem().Kind()
			if __obf_9bbf49c8c89fbd8a == reflect.Uint8 {
				return string(__obf_a4fef5c9586c1966.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(__obf_a4fef5c9586c1966.Uint(), 10), nil
		}
	}

	return __obf_333b2acf09d806b3, nil
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
	return func(__obf_4cd911a2d7099b07 reflect.Value, __obf_7af45612bc672ea8 reflect.Value) (any, error) {
		if __obf_4cd911a2d7099b07.Kind() != reflect.Struct {
			return __obf_4cd911a2d7099b07.Interface(), nil
		}

		var __obf_99f0e7230007b81b any = struct{}{}
		if __obf_7af45612bc672ea8.Type() != reflect.TypeOf(&__obf_99f0e7230007b81b).Elem() {
			return __obf_4cd911a2d7099b07.Interface(), nil
		}
		__obf_36232451cab2dc03 := make(map[string]any)
		__obf_7af45612bc672ea8.
			Set(reflect.ValueOf(__obf_36232451cab2dc03))

		return __obf_4cd911a2d7099b07.Interface(), nil
	}
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshallerHookFunc() DecodeHookFuncType {
	return func(__obf_4cd911a2d7099b07 reflect.Type, __obf_7af45612bc672ea8 reflect.Type, __obf_333b2acf09d806b3 any) (any, error) {
		if __obf_4cd911a2d7099b07.Kind() != reflect.String {
			return __obf_333b2acf09d806b3, nil
		}
		__obf_382efc6537d96d52 := reflect.New(__obf_7af45612bc672ea8).Interface()
		__obf_d0977f2d6a589c81, __obf_c5c13e02010ecc6c := __obf_382efc6537d96d52.(encoding.TextUnmarshaler)
		if !__obf_c5c13e02010ecc6c {
			return __obf_333b2acf09d806b3, nil
		}
		__obf_98acb43641f668a3, __obf_c5c13e02010ecc6c := __obf_333b2acf09d806b3.(string)
		if !__obf_c5c13e02010ecc6c {
			__obf_98acb43641f668a3 = reflect.Indirect(reflect.ValueOf(&__obf_333b2acf09d806b3)).Elem().String()
		}
		if __obf_6585aef4313e6005 := __obf_d0977f2d6a589c81.UnmarshalText([]byte(__obf_98acb43641f668a3)); __obf_6585aef4313e6005 != nil {
			return nil, __obf_6585aef4313e6005
		}
		return __obf_382efc6537d96d52, nil
	}
}
