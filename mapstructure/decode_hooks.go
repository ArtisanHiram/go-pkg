package __obf_c953b7a5114a5dbe

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
func __obf_3f0502bf47cc5bb4(__obf_5282ccb3b20b97cb DecodeHookFunc) DecodeHookFunc {
	// Create variables here so we can reference them with the reflect pkg
	var __obf_dc218385319b42f5 DecodeHookFuncType
	var __obf_696cd2aa2a225e0d DecodeHookFuncKind
	var __obf_c8b3e84533d385d1 DecodeHookFuncValue
	__obf_bd7fa4561867c57c := // Fill in the variables into this interface and the rest is done
		// automatically using the reflect package.
		[]any{__obf_dc218385319b42f5, __obf_696cd2aa2a225e0d, __obf_c8b3e84533d385d1}
	__obf_71e738e554663e8c := reflect.ValueOf(__obf_5282ccb3b20b97cb)
	__obf_fa6eab0139f5e064 := __obf_71e738e554663e8c.Type()
	for _, __obf_923a3d5018b1dd17 := range __obf_bd7fa4561867c57c {
		__obf_58eb82edda18275a := reflect.ValueOf(__obf_923a3d5018b1dd17).Type()
		if __obf_fa6eab0139f5e064.ConvertibleTo(__obf_58eb82edda18275a) {
			return __obf_71e738e554663e8c.Convert(__obf_58eb82edda18275a).Interface()
		}
	}

	return nil
}

// DecodeHookExec executes the given decode hook. This should be used
// since it'll naturally degrade to the older backwards compatible DecodeHookFunc
// that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(__obf_923a3d5018b1dd17 DecodeHookFunc, __obf_14810ac6c1a1ff4b reflect.Value, __obf_bb1596a458011647 reflect.Value) (any, error) {

	switch __obf_f500d725634477d0 := __obf_3f0502bf47cc5bb4(__obf_923a3d5018b1dd17).(type) {
	case DecodeHookFuncType:
		return __obf_f500d725634477d0(__obf_14810ac6c1a1ff4b.Type(), __obf_bb1596a458011647.Type(), __obf_14810ac6c1a1ff4b.Interface())
	case DecodeHookFuncKind:
		return __obf_f500d725634477d0(__obf_14810ac6c1a1ff4b.Kind(), __obf_bb1596a458011647.Kind(), __obf_14810ac6c1a1ff4b.Interface())
	case DecodeHookFuncValue:
		return __obf_f500d725634477d0(__obf_14810ac6c1a1ff4b, __obf_bb1596a458011647)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

// ComposeDecodeHookFunc creates a single DecodeHookFunc that
// automatically composes multiple DecodeHookFuncs.
//
// The composed funcs are called in order, with the result of the
// previous transformation.
func ComposeDecodeHookFunc(__obf_a81f214e35488842 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Value, __obf_a4317c2d3eff175b reflect.Value) (any, error) {
		var __obf_fb7fbd58154fa1dd error
		__obf_50c6df487a55f30e := __obf_f500d725634477d0.Interface()
		__obf_5c6af0d0c65b0161 := __obf_f500d725634477d0
		for _, __obf_dc218385319b42f5 := range __obf_a81f214e35488842 {
			__obf_50c6df487a55f30e, __obf_fb7fbd58154fa1dd = DecodeHookExec(__obf_dc218385319b42f5, __obf_5c6af0d0c65b0161, __obf_a4317c2d3eff175b)
			if __obf_fb7fbd58154fa1dd != nil {
				return nil, __obf_fb7fbd58154fa1dd
			}
			__obf_5c6af0d0c65b0161 = reflect.ValueOf(__obf_50c6df487a55f30e)
		}

		return __obf_50c6df487a55f30e, nil
	}
}

// OrComposeDecodeHookFunc executes all input hook functions until one of them returns no error. In that case its value is returned.
// If all hooks return an error, OrComposeDecodeHookFunc returns an error concatenating all error messages.
func OrComposeDecodeHookFunc(__obf_7819447ae8d1be76 ...DecodeHookFunc) DecodeHookFunc {
	return func(__obf_4a31f88a30a6d70f, __obf_fde44f2ea6c186ca reflect.Value) (any, error) {
		var __obf_567f9f5f27924b01 string
		var __obf_f535e2e35c6cc7e1 any
		var __obf_fb7fbd58154fa1dd error

		for _, __obf_f500d725634477d0 := range __obf_7819447ae8d1be76 {
			__obf_f535e2e35c6cc7e1, __obf_fb7fbd58154fa1dd = DecodeHookExec(__obf_f500d725634477d0, __obf_4a31f88a30a6d70f, __obf_fde44f2ea6c186ca)
			if __obf_fb7fbd58154fa1dd != nil {
				__obf_567f9f5f27924b01 += __obf_fb7fbd58154fa1dd.Error() + "\n"
				continue
			}

			return __obf_f535e2e35c6cc7e1, nil
		}

		return nil, errors.New(__obf_567f9f5f27924b01)
	}
}

// StringToSliceHookFunc returns a DecodeHookFunc that converts
// string to []string by splitting on the given sep.
func StringToSliceHookFunc(__obf_04b9e06755411223 string) DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Kind, __obf_a4317c2d3eff175b reflect.Kind, __obf_50c6df487a55f30e any) (any, error) {
		if __obf_f500d725634477d0 != reflect.String || __obf_a4317c2d3eff175b != reflect.Slice {
			return __obf_50c6df487a55f30e, nil
		}
		__obf_923a3d5018b1dd17 := __obf_50c6df487a55f30e.(string)
		if __obf_923a3d5018b1dd17 == "" {
			return []string{}, nil
		}

		return strings.Split(__obf_923a3d5018b1dd17,

			// StringToTimeDurationHookFunc returns a DecodeHookFunc that converts
			// strings to time.Duration.
			__obf_04b9e06755411223), nil
	}
}

func StringToTimeDurationHookFunc() DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Type, __obf_a4317c2d3eff175b reflect.Type, __obf_50c6df487a55f30e any) (any, error) {
		if __obf_f500d725634477d0.Kind() != reflect.String {
			return __obf_50c6df487a55f30e, nil
		}
		if __obf_a4317c2d3eff175b != reflect.TypeOf(time.Duration(5)) {
			return __obf_50c6df487a55f30e, nil
		}

		// Convert it by parsing
		return time.ParseDuration(__obf_50c6df487a55f30e.(string))
	}
}

// StringToIPHookFunc returns a DecodeHookFunc that converts
// strings to net.IP
func StringToIPHookFunc() DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Type, __obf_a4317c2d3eff175b reflect.Type, __obf_50c6df487a55f30e any) (any, error) {
		if __obf_f500d725634477d0.Kind() != reflect.String {
			return __obf_50c6df487a55f30e, nil
		}
		if __obf_a4317c2d3eff175b != reflect.TypeOf(net.IP{}) {
			return __obf_50c6df487a55f30e, nil
		}
		__obf_f0e28abcec4c3f59 := // Convert it by parsing
			net.ParseIP(__obf_50c6df487a55f30e.(string))
		if __obf_f0e28abcec4c3f59 == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", __obf_50c6df487a55f30e)
		}

		return __obf_f0e28abcec4c3f59, nil
	}
}

// StringToIPNetHookFunc returns a DecodeHookFunc that converts
// strings to net.IPNet
func StringToIPNetHookFunc() DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Type, __obf_a4317c2d3eff175b reflect.Type, __obf_50c6df487a55f30e any) (any, error) {
		if __obf_f500d725634477d0.Kind() != reflect.String {
			return __obf_50c6df487a55f30e, nil
		}
		if __obf_a4317c2d3eff175b != reflect.TypeOf(net.IPNet{}) {
			return __obf_50c6df487a55f30e, nil
		}

		// Convert it by parsing
		_, net, __obf_fb7fbd58154fa1dd := net.ParseCIDR(__obf_50c6df487a55f30e.(string))
		return net, __obf_fb7fbd58154fa1dd
	}
}

// StringToTimeHookFunc returns a DecodeHookFunc that converts
// strings to time.Time.
func StringToTimeHookFunc(__obf_f1a6faec27b1b900 string) DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Type, __obf_a4317c2d3eff175b reflect.Type, __obf_50c6df487a55f30e any) (any, error) {
		if __obf_f500d725634477d0.Kind() != reflect.String {
			return __obf_50c6df487a55f30e, nil
		}
		if __obf_a4317c2d3eff175b != reflect.TypeOf(time.Time{}) {
			return __obf_50c6df487a55f30e, nil
		}

		// Convert it by parsing
		return time.Parse(__obf_f1a6faec27b1b900, __obf_50c6df487a55f30e.(string))
	}
}

// WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
// the decoder.
//
// Note that this is significantly different from the WeaklyTypedInput option
// of the DecoderConfig.
func WeaklyTypedHook(__obf_f500d725634477d0 reflect.Kind, __obf_a4317c2d3eff175b reflect.Kind, __obf_50c6df487a55f30e any) (any, error) {
	__obf_a361258b272167f7 := reflect.ValueOf(__obf_50c6df487a55f30e)
	switch __obf_a4317c2d3eff175b {
	case reflect.String:
		switch __obf_f500d725634477d0 {
		case reflect.Bool:
			if __obf_a361258b272167f7.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(__obf_a361258b272167f7.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(__obf_a361258b272167f7.Int(), 10), nil
		case reflect.Slice:
			__obf_b42cbbd1bcdb3a20 := __obf_a361258b272167f7.Type()
			__obf_021d31b2022f7cf6 := __obf_b42cbbd1bcdb3a20.Elem().Kind()
			if __obf_021d31b2022f7cf6 == reflect.Uint8 {
				return string(__obf_a361258b272167f7.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(__obf_a361258b272167f7.Uint(), 10), nil
		}
	}

	return __obf_50c6df487a55f30e, nil
}

func RecursiveStructToMapHookFunc() DecodeHookFunc {
	return func(__obf_f500d725634477d0 reflect.Value, __obf_a4317c2d3eff175b reflect.Value) (any, error) {
		if __obf_f500d725634477d0.Kind() != reflect.Struct {
			return __obf_f500d725634477d0.Interface(), nil
		}

		var __obf_507d21ee6c71f742 any = struct{}{}
		if __obf_a4317c2d3eff175b.Type() != reflect.TypeOf(&__obf_507d21ee6c71f742).Elem() {
			return __obf_f500d725634477d0.Interface(), nil
		}
		__obf_9fff3e6ecc8a1671 := make(map[string]any)
		__obf_a4317c2d3eff175b.
			Set(reflect.ValueOf(__obf_9fff3e6ecc8a1671))

		return __obf_f500d725634477d0.Interface(), nil
	}
}

// TextUnmarshallerHookFunc returns a DecodeHookFunc that applies
// strings to the UnmarshalText function, when the target type
// implements the encoding.TextUnmarshaler interface
func TextUnmarshallerHookFunc() DecodeHookFuncType {
	return func(__obf_f500d725634477d0 reflect.Type, __obf_a4317c2d3eff175b reflect.Type, __obf_50c6df487a55f30e any) (any, error) {
		if __obf_f500d725634477d0.Kind() != reflect.String {
			return __obf_50c6df487a55f30e, nil
		}
		__obf_47a3e13e334b0913 := reflect.New(__obf_a4317c2d3eff175b).Interface()
		__obf_c02331217c95e7f2, __obf_f67b17ac0536943b := __obf_47a3e13e334b0913.(encoding.TextUnmarshaler)
		if !__obf_f67b17ac0536943b {
			return __obf_50c6df487a55f30e, nil
		}
		__obf_58fce06456cade02, __obf_f67b17ac0536943b := __obf_50c6df487a55f30e.(string)
		if !__obf_f67b17ac0536943b {
			__obf_58fce06456cade02 = reflect.Indirect(reflect.ValueOf(&__obf_50c6df487a55f30e)).Elem().String()
		}
		if __obf_fb7fbd58154fa1dd := __obf_c02331217c95e7f2.UnmarshalText([]byte(__obf_58fce06456cade02)); __obf_fb7fbd58154fa1dd != nil {
			return nil, __obf_fb7fbd58154fa1dd
		}
		return __obf_47a3e13e334b0913, nil
	}
}
