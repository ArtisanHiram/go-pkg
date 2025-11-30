package __obf_ec2f65f16fa88470

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
func __obf_1ed4291921c84c37(__obf_ce1cd81a1e481cca DecodeHookFunc) DecodeHookFunc {
	// Create variables here so we can reference them with the reflect pkg
	var __obf_d0fe1b34a9b256b5 DecodeHookFuncType
	var __obf_b10b51c89072eaf2 DecodeHookFuncKind
	var __obf_d9b36f89a227c094 DecodeHookFuncValue
	__obf_e82d6c6df12a6f88 := // Fill in the variables into this interface and the rest is done
		// automatically using the reflect package.
		[]any{__obf_d0fe1b34a9b256b5, __obf_b10b51c89072eaf2, __obf_d9b36f89a227c094}
	__obf_7646bc0f4c5c46a0 := reflect.ValueOf(__obf_ce1cd81a1e481cca)
	__obf_700f0af0d255eb29 := __obf_7646bc0f4c5c46a0.Type()
	for _, __obf_ea4e291789bf028a := range __obf_e82d6c6df12a6f88 {
		__obf_299dd768b6b52b49 := reflect.ValueOf(__obf_ea4e291789bf028a).Type()
		if __obf_700f0af0d255eb29.ConvertibleTo(__obf_299dd768b6b52b49) {
			return __obf_7646bc0f4c5c46a0.Convert(__obf_299dd768b6b52b49).Interface()
		}
	}

	return nil
}

// DecodeHookExec executes the given decode hook. This should be used
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
// that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(__obf_ea4e291789bf028a DecodeHookFunc, __obf_3a8aa8877b053f0b reflect.Value, __obf_6dd945bb01fb3b32 reflect.Value) (any, error) {

	switch __obf_502bda9c3cc67a99 := __obf_1ed4291921c84c37(__obf_ea4e291789bf028a).(type) {
	case DecodeHookFuncType:
		return __obf_502bda9c3cc67a99(__obf_3a8aa8877b053f0b.Type(), __obf_6dd945bb01fb3b32.Type(), __obf_3a8aa8877b053f0b.Interface())
	case DecodeHookFuncKind:
		return __obf_502bda9c3cc67a99(__obf_3a8aa8877b053f0b.Kind(), __obf_6dd945bb01fb3b32.Kind(), __obf_3a8aa8877b053f0b.Interface())
	case DecodeHookFuncValue:
		return __obf_502bda9c3cc67a99(__obf_3a8aa8877b053f0b, __obf_6dd945bb01fb3b32)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
// automatically composes multiple DecodeHookFuncs.
//
// The composed funcs are called in order, with the result of the
// previous transformation.
func ComposeDecodeHookFunc(__obf_ee8b78d564b5e19a ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Value, __obf_5b64473fd5ec5ddf reflect.Value) (any, error) {
		var __obf_6d8ef670b0943b3c error
		__obf_e428e8b18837f120 := __obf_502bda9c3cc67a99.Interface()
		__obf_77425076d6ad3a6a := __obf_502bda9c3cc67a99
		for _, __obf_d0fe1b34a9b256b5 := range __obf_ee8b78d564b5e19a {
			__obf_e428e8b18837f120, __obf_6d8ef670b0943b3c = DecodeHookExec(__obf_d0fe1b34a9b256b5, __obf_77425076d6ad3a6a, __obf_5b64473fd5ec5ddf)
			if __obf_6d8ef670b0943b3c != nil {
				return nil, __obf_6d8ef670b0943b3c
			}
			__obf_77425076d6ad3a6a = reflect.ValueOf(__obf_e428e8b18837f120)
		}

		return __obf_e428e8b18837f120, nil
	}
}

// OrComposeDecodeHookFunc executes all input hook functions until one of them returns no error. In that case its value is returned.
// If all hooks return an error, OrComposeDecodeHookFunc returns an error concatenating all error messages.
func OrComposeDecodeHookFunc(__obf_00b2eef44b113463 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_1d31a60b34fd6107, __obf_ed3004c91f084575 reflect.Value) (any, error) {
		var __obf_d5d44c42ab9034f3 string
		var __obf_3852e370232dcc1f any
		var __obf_6d8ef670b0943b3c error

		for _, __obf_502bda9c3cc67a99 := range __obf_00b2eef44b113463 {
			__obf_3852e370232dcc1f, __obf_6d8ef670b0943b3c = DecodeHookExec(__obf_502bda9c3cc67a99, __obf_1d31a60b34fd6107, __obf_ed3004c91f084575)
			if __obf_6d8ef670b0943b3c != nil {
				__obf_d5d44c42ab9034f3 += __obf_6d8ef670b0943b3c.Error() + "\n"
				continue
			}

			return __obf_3852e370232dcc1f, nil
		}

		return nil, errors.New(__obf_d5d44c42ab9034f3)
	}
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
// string to []string by splitting on the given sep.
func StringToSliceHookFunc(__obf_89f223de73c3e039 string) DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Kind, __obf_5b64473fd5ec5ddf reflect.Kind, __obf_e428e8b18837f120 any) (any, error) {
		if __obf_502bda9c3cc67a99 != reflect.String || __obf_5b64473fd5ec5ddf != reflect.Slice {
			return __obf_e428e8b18837f120, nil
		}
		__obf_ea4e291789bf028a := __obf_e428e8b18837f120.(string)
		if __obf_ea4e291789bf028a == "" {
			return []string{}, nil
		}

		return strings.Split(__obf_ea4e291789bf028a,

			// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
			// strings to time.Duration.
			__obf_89f223de73c3e039), nil
	}
}

func StringToTimeDurationHookFunc() DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Type, __obf_5b64473fd5ec5ddf reflect.Type, __obf_e428e8b18837f120 any) (any, error) {
		if __obf_502bda9c3cc67a99.Kind() != reflect.String {
			return __obf_e428e8b18837f120, nil
		}
		if __obf_5b64473fd5ec5ddf != reflect.TypeOf(time.Duration(5)) {
			return __obf_e428e8b18837f120, nil
		}

		// Convert it by parsing
		return time.ParseDuration(__obf_e428e8b18837f120.(string))
	}
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
// strings to net.IP
func StringToIPHookFunc() DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Type, __obf_5b64473fd5ec5ddf reflect.Type, __obf_e428e8b18837f120 any) (any, error) {
		if __obf_502bda9c3cc67a99.Kind() != reflect.String {
			return __obf_e428e8b18837f120, nil
		}
		if __obf_5b64473fd5ec5ddf != reflect.TypeOf(net.IP{}) {
			return __obf_e428e8b18837f120, nil
		}
		__obf_ea4fccc6b378a7cf := // Convert it by parsing
			net.ParseIP(__obf_e428e8b18837f120.(string))
		if __obf_ea4fccc6b378a7cf == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", __obf_e428e8b18837f120)
		}

		return __obf_ea4fccc6b378a7cf, nil
	}
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
// strings to net.IPNet
func StringToIPNetHookFunc() DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Type, __obf_5b64473fd5ec5ddf reflect.Type, __obf_e428e8b18837f120 any) (any, error) {
		if __obf_502bda9c3cc67a99.Kind() != reflect.String {
			return __obf_e428e8b18837f120, nil
		}
		if __obf_5b64473fd5ec5ddf != reflect.TypeOf(net.IPNet{}) {
			return __obf_e428e8b18837f120, nil
		}

		// Convert it by parsing
		_, net, __obf_6d8ef670b0943b3c := net.ParseCIDR(__obf_e428e8b18837f120.(string))
		return net, __obf_6d8ef670b0943b3c
	}
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func StringToTimeHookFunc(__obf_3d9eae9c3ad3bbd5 string) DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Type, __obf_5b64473fd5ec5ddf reflect.Type, __obf_e428e8b18837f120 any) (any, error) {
		if __obf_502bda9c3cc67a99.Kind() != reflect.String {
			return __obf_e428e8b18837f120, nil
		}
		if __obf_5b64473fd5ec5ddf != reflect.TypeOf(time.Time{}) {
			return __obf_e428e8b18837f120, nil
		}

		// Convert it by parsing
		return time.Parse(__obf_3d9eae9c3ad3bbd5, __obf_e428e8b18837f120.(string))
	}
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
// the decoder.
//
// Note that this is significantly different from the WeaklyTypedInput option
// of the DecoderConfig.
func WeaklyTypedHook(__obf_502bda9c3cc67a99 reflect.Kind, __obf_5b64473fd5ec5ddf reflect.Kind, __obf_e428e8b18837f120 any) (any, error) {
	__obf_a30e8e069ebf8d56 := reflect.ValueOf(__obf_e428e8b18837f120)
	switch __obf_5b64473fd5ec5ddf {
	case reflect.String:
		switch __obf_502bda9c3cc67a99 {
		case reflect.Bool:
			if __obf_a30e8e069ebf8d56.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(__obf_a30e8e069ebf8d56.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(__obf_a30e8e069ebf8d56.Int(), 10), nil
		case reflect.Slice:
			__obf_9dd014252313d6cd := __obf_a30e8e069ebf8d56.Type()
			__obf_0673273cce64670d := __obf_9dd014252313d6cd.Elem().Kind()
			if __obf_0673273cce64670d == reflect.Uint8 {
				return string(__obf_a30e8e069ebf8d56.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(__obf_a30e8e069ebf8d56.Uint(), 10), nil
		}
	}

	return __obf_e428e8b18837f120, nil
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
	return func(__obf_502bda9c3cc67a99 reflect.Value, __obf_5b64473fd5ec5ddf reflect.Value) (any, error) {
		if __obf_502bda9c3cc67a99.Kind() != reflect.Struct {
			return __obf_502bda9c3cc67a99.Interface(), nil
		}

		var __obf_6cbf6487a01daf31 any = struct{}{}
		if __obf_5b64473fd5ec5ddf.Type() != reflect.TypeOf(&__obf_6cbf6487a01daf31).Elem() {
			return __obf_502bda9c3cc67a99.Interface(), nil
		}
		__obf_aeebc03738ee7adc := make(map[string]any)
		__obf_5b64473fd5ec5ddf.
			Set(reflect.ValueOf(__obf_aeebc03738ee7adc))

		return __obf_502bda9c3cc67a99.Interface(), nil
	}
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshallerHookFunc() DecodeHookFuncType {
	return func(__obf_502bda9c3cc67a99 reflect.Type, __obf_5b64473fd5ec5ddf reflect.Type, __obf_e428e8b18837f120 any) (any, error) {
		if __obf_502bda9c3cc67a99.Kind() != reflect.String {
			return __obf_e428e8b18837f120, nil
		}
		__obf_bb9d1909c7b1c125 := reflect.New(__obf_5b64473fd5ec5ddf).Interface()
		__obf_b295189023a904e0, __obf_8a9bb3ba25fa94bd := __obf_bb9d1909c7b1c125.(encoding.TextUnmarshaler)
		if !__obf_8a9bb3ba25fa94bd {
			return __obf_e428e8b18837f120, nil
		}
		__obf_ae7e7e9fd3e1b096, __obf_8a9bb3ba25fa94bd := __obf_e428e8b18837f120.(string)
		if !__obf_8a9bb3ba25fa94bd {
			__obf_ae7e7e9fd3e1b096 = reflect.Indirect(reflect.ValueOf(&__obf_e428e8b18837f120)).Elem().String()
		}
		if __obf_6d8ef670b0943b3c := __obf_b295189023a904e0.UnmarshalText([]byte(__obf_ae7e7e9fd3e1b096)); __obf_6d8ef670b0943b3c != nil {
			return nil, __obf_6d8ef670b0943b3c
		}
		return __obf_bb9d1909c7b1c125, nil
	}
}
