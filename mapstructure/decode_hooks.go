package __obf_8873749254e9e83f

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
func __obf_944e59e3f526649f(__obf_0d5ec26e51d4b2f9 DecodeHookFunc) DecodeHookFunc {
	// Create variables here so we can reference them with the reflect pkg
	var __obf_4a8181d60dbb0080 DecodeHookFuncType
	var __obf_dc49611b40eeb0fc DecodeHookFuncKind
	var __obf_71066a22f0f215af DecodeHookFuncValue
	__obf_daaa5b9473e3396a := // Fill in the variables into this interface and the rest is done
		// automatically using the reflect package.
		[]any{__obf_4a8181d60dbb0080, __obf_dc49611b40eeb0fc, __obf_71066a22f0f215af}
	__obf_40a0aa2002a96053 := reflect.ValueOf(__obf_0d5ec26e51d4b2f9)
	__obf_3ca6ea1726c5f916 := __obf_40a0aa2002a96053.Type()
	for _, __obf_907fdbb25ddead55 := range __obf_daaa5b9473e3396a {
		__obf_2a619a21d532f91b := reflect.ValueOf(__obf_907fdbb25ddead55).Type()
		if __obf_3ca6ea1726c5f916.ConvertibleTo(__obf_2a619a21d532f91b) {
			return __obf_40a0aa2002a96053.Convert(__obf_2a619a21d532f91b).Interface()
		}
	}

	return nil
}

// DecodeHookExec executes the given decode hook. This should be used
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
// that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(__obf_907fdbb25ddead55 DecodeHookFunc, __obf_9498a1e77e30674f reflect.Value, __obf_aef128244080a96d reflect.Value) (any, error) {

	switch __obf_7bfc349763f12493 := __obf_944e59e3f526649f(__obf_907fdbb25ddead55).(type) {
	case DecodeHookFuncType:
		return __obf_7bfc349763f12493(__obf_9498a1e77e30674f.Type(), __obf_aef128244080a96d.Type(), __obf_9498a1e77e30674f.Interface())
	case DecodeHookFuncKind:
		return __obf_7bfc349763f12493(__obf_9498a1e77e30674f.Kind(), __obf_aef128244080a96d.Kind(), __obf_9498a1e77e30674f.Interface())
	case DecodeHookFuncValue:
		return __obf_7bfc349763f12493(__obf_9498a1e77e30674f, __obf_aef128244080a96d)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
// automatically composes multiple DecodeHookFuncs.
//
// The composed funcs are called in order, with the result of the
// previous transformation.
func ComposeDecodeHookFunc(__obf_dafb960d523ab157 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Value, __obf_c8c1434db3a4857c reflect.Value) (any, error) {
		var __obf_ce3ada6ea4dc41de error
		__obf_4bc7fc453056091b := __obf_7bfc349763f12493.Interface()
		__obf_58c95b76870f9962 := __obf_7bfc349763f12493
		for _, __obf_4a8181d60dbb0080 := range __obf_dafb960d523ab157 {
			__obf_4bc7fc453056091b, __obf_ce3ada6ea4dc41de = DecodeHookExec(__obf_4a8181d60dbb0080, __obf_58c95b76870f9962, __obf_c8c1434db3a4857c)
			if __obf_ce3ada6ea4dc41de != nil {
				return nil, __obf_ce3ada6ea4dc41de
			}
			__obf_58c95b76870f9962 = reflect.ValueOf(__obf_4bc7fc453056091b)
		}

		return __obf_4bc7fc453056091b, nil
	}
}

// OrComposeDecodeHookFunc executes all input hook functions until one of them returns no error. In that case its value is returned.
// If all hooks return an error, OrComposeDecodeHookFunc returns an error concatenating all error messages.
func OrComposeDecodeHookFunc(__obf_54befb0a2b1fc861 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_b4e301dccffcf27a, __obf_2509c56426afe825 reflect.Value) (any, error) {
		var __obf_e7d03e38cd36fe3d string
		var __obf_bb14be16ab94408f any
		var __obf_ce3ada6ea4dc41de error

		for _, __obf_7bfc349763f12493 := range __obf_54befb0a2b1fc861 {
			__obf_bb14be16ab94408f, __obf_ce3ada6ea4dc41de = DecodeHookExec(__obf_7bfc349763f12493, __obf_b4e301dccffcf27a, __obf_2509c56426afe825)
			if __obf_ce3ada6ea4dc41de != nil {
				__obf_e7d03e38cd36fe3d += __obf_ce3ada6ea4dc41de.Error() + "\n"
				continue
			}

			return __obf_bb14be16ab94408f, nil
		}

		return nil, errors.New(__obf_e7d03e38cd36fe3d)
	}
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
// string to []string by splitting on the given sep.
func StringToSliceHookFunc(__obf_4a3ca6f08bf443fd string) DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Kind, __obf_c8c1434db3a4857c reflect.Kind, __obf_4bc7fc453056091b any) (any, error) {
		if __obf_7bfc349763f12493 != reflect.String || __obf_c8c1434db3a4857c != reflect.Slice {
			return __obf_4bc7fc453056091b, nil
		}
		__obf_907fdbb25ddead55 := __obf_4bc7fc453056091b.(string)
		if __obf_907fdbb25ddead55 == "" {
			return []string{}, nil
		}

		return strings.Split(__obf_907fdbb25ddead55,

			// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
			// strings to time.Duration.
			__obf_4a3ca6f08bf443fd), nil
	}
}

func StringToTimeDurationHookFunc() DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Type, __obf_c8c1434db3a4857c reflect.Type, __obf_4bc7fc453056091b any) (any, error) {
		if __obf_7bfc349763f12493.Kind() != reflect.String {
			return __obf_4bc7fc453056091b, nil
		}
		if __obf_c8c1434db3a4857c != reflect.TypeOf(time.Duration(5)) {
			return __obf_4bc7fc453056091b, nil
		}

		// Convert it by parsing
		return time.ParseDuration(__obf_4bc7fc453056091b.(string))
	}
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
// strings to net.IP
func StringToIPHookFunc() DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Type, __obf_c8c1434db3a4857c reflect.Type, __obf_4bc7fc453056091b any) (any, error) {
		if __obf_7bfc349763f12493.Kind() != reflect.String {
			return __obf_4bc7fc453056091b, nil
		}
		if __obf_c8c1434db3a4857c != reflect.TypeOf(net.IP{}) {
			return __obf_4bc7fc453056091b, nil
		}
		__obf_615dbf65535b79c1 := // Convert it by parsing
			net.ParseIP(__obf_4bc7fc453056091b.(string))
		if __obf_615dbf65535b79c1 == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", __obf_4bc7fc453056091b)
		}

		return __obf_615dbf65535b79c1, nil
	}
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
// strings to net.IPNet
func StringToIPNetHookFunc() DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Type, __obf_c8c1434db3a4857c reflect.Type, __obf_4bc7fc453056091b any) (any, error) {
		if __obf_7bfc349763f12493.Kind() != reflect.String {
			return __obf_4bc7fc453056091b, nil
		}
		if __obf_c8c1434db3a4857c != reflect.TypeOf(net.IPNet{}) {
			return __obf_4bc7fc453056091b, nil
		}

		// Convert it by parsing
		_, net, __obf_ce3ada6ea4dc41de := net.ParseCIDR(__obf_4bc7fc453056091b.(string))
		return net, __obf_ce3ada6ea4dc41de
	}
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func StringToTimeHookFunc(__obf_6115f5857336fcd3 string) DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Type, __obf_c8c1434db3a4857c reflect.Type, __obf_4bc7fc453056091b any) (any, error) {
		if __obf_7bfc349763f12493.Kind() != reflect.String {
			return __obf_4bc7fc453056091b, nil
		}
		if __obf_c8c1434db3a4857c != reflect.TypeOf(time.Time{}) {
			return __obf_4bc7fc453056091b, nil
		}

		// Convert it by parsing
		return time.Parse(__obf_6115f5857336fcd3, __obf_4bc7fc453056091b.(string))
	}
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
// the decoder.
//
// Note that this is significantly different from the WeaklyTypedInput option
// of the DecoderConfig.
func WeaklyTypedHook(__obf_7bfc349763f12493 reflect.Kind, __obf_c8c1434db3a4857c reflect.Kind, __obf_4bc7fc453056091b any) (any, error) {
	__obf_0bfc5beb6a20600c := reflect.ValueOf(__obf_4bc7fc453056091b)
	switch __obf_c8c1434db3a4857c {
	case reflect.String:
		switch __obf_7bfc349763f12493 {
		case reflect.Bool:
			if __obf_0bfc5beb6a20600c.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(__obf_0bfc5beb6a20600c.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(__obf_0bfc5beb6a20600c.Int(), 10), nil
		case reflect.Slice:
			__obf_cb1be79795600114 := __obf_0bfc5beb6a20600c.Type()
			__obf_55286990c7ae804f := __obf_cb1be79795600114.Elem().Kind()
			if __obf_55286990c7ae804f == reflect.Uint8 {
				return string(__obf_0bfc5beb6a20600c.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(__obf_0bfc5beb6a20600c.Uint(), 10), nil
		}
	}

	return __obf_4bc7fc453056091b, nil
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
	return func(__obf_7bfc349763f12493 reflect.Value, __obf_c8c1434db3a4857c reflect.Value) (any, error) {
		if __obf_7bfc349763f12493.Kind() != reflect.Struct {
			return __obf_7bfc349763f12493.Interface(), nil
		}

		var __obf_4e5b08e68a148355 any = struct{}{}
		if __obf_c8c1434db3a4857c.Type() != reflect.TypeOf(&__obf_4e5b08e68a148355).Elem() {
			return __obf_7bfc349763f12493.Interface(), nil
		}
		__obf_2b6f8ee86c5211df := make(map[string]any)
		__obf_c8c1434db3a4857c.
			Set(reflect.ValueOf(__obf_2b6f8ee86c5211df))

		return __obf_7bfc349763f12493.Interface(), nil
	}
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshallerHookFunc() DecodeHookFuncType {
	return func(__obf_7bfc349763f12493 reflect.Type, __obf_c8c1434db3a4857c reflect.Type, __obf_4bc7fc453056091b any) (any, error) {
		if __obf_7bfc349763f12493.Kind() != reflect.String {
			return __obf_4bc7fc453056091b, nil
		}
		__obf_a9459758946d2eaa := reflect.New(__obf_c8c1434db3a4857c).Interface()
		__obf_b5c29378ff4fd89a, __obf_5006c1b1bc16f11d := __obf_a9459758946d2eaa.(encoding.TextUnmarshaler)
		if !__obf_5006c1b1bc16f11d {
			return __obf_4bc7fc453056091b, nil
		}
		__obf_288b09ee2b024465, __obf_5006c1b1bc16f11d := __obf_4bc7fc453056091b.(string)
		if !__obf_5006c1b1bc16f11d {
			__obf_288b09ee2b024465 = reflect.Indirect(reflect.ValueOf(&__obf_4bc7fc453056091b)).Elem().String()
		}
		if __obf_ce3ada6ea4dc41de := __obf_b5c29378ff4fd89a.UnmarshalText([]byte(__obf_288b09ee2b024465)); __obf_ce3ada6ea4dc41de != nil {
			return nil, __obf_ce3ada6ea4dc41de
		}
		return __obf_a9459758946d2eaa, nil
	}
}
