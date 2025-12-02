package __obf_8873749254e9e83f

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// DecodeHookFunc is the callback function that can be used for
// data transformations. See "DecodeHook" in the DecoderConfig
// struct.
//
// The type must be one of DecodeHookFuncType, DecodeHookFuncKind, or
// DecodeHookFuncValue.
// Values are a superset of Types (Values can return types), and Types are a
// superset of Kinds (Types can return Kinds) and are generally a richer thing
// to use, but Kinds are simpler if you only need those.
//
// The reason DecodeHookFunc is multi-typed is for backwards compatibility:
// we started with Kinds and then realized Types were the better solution,
// but have a promise to not break backwards compat so we now support
// both.
type DecodeHookFunc any

// DecodeHookFuncType is a DecodeHookFunc which has complete information about
// the source and target types.
type DecodeHookFuncType func(reflect.Type, reflect.Type, any) (any, error)

// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
// source and target types.
type DecodeHookFuncKind func(reflect.Kind, reflect.Kind, any) (any, error)

// DecodeHookFuncValue is a DecodeHookFunc which has complete access to both the source and target
// values.
type DecodeHookFuncValue func(__obf_9498a1e77e30674f reflect.Value, __obf_aef128244080a96d reflect.Value) (any, error)

// DecoderConfig is the configuration that is used to create a new decoder
// and allows customization of various aspects of decoding.
type DecoderConfig struct {
	// DecodeHook, if set, will be called before any decoding and any
	// type conversion (if WeaklyTypedInput is on). This lets you modify
	// the values before they're set down onto the resulting struct. The
	// DecodeHook is called for every map and value in the input. This means
	// that if a struct has embedded fields with squash tags the decode hook
	// is called only once with all of the input data, not once for each
	// embedded struct.
	//
	// If an error is returned, the entire decode will fail with that error.
	DecodeHook DecodeHookFunc

	// If ErrorUnused is true, then it is an error for there to exist
	// keys in the original map that were unused in the decoding process
	// (extra keys).
	ErrorUnused bool

	// If ErrorUnset is true, then it is an error for there to exist
	// fields in the result that were not set in the decoding process
	// (extra fields). This only applies to decoding to a struct. This
	// will affect all nested structs as well.
	ErrorUnset bool

	// ZeroFields, if set to true, will zero fields before writing them.
	// For example, a map will be emptied before decoded values are put in
	// it. If this is false, a map will be merged.
	ZeroFields bool

	// If WeaklyTypedInput is true, the decoder will make the following
	// "weak" conversions:
	//
	//   - bools to string (true = "1", false = "0")
	//   - numbers to string (base 10)
	//   - bools to int/uint (true = 1, false = 0)
	//   - strings to int/uint (base implied by prefix)
	//   - int to bool (true if value != 0)
	//   - string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
	//     FALSE, false, False. Anything else is an error)
	//   - empty array = empty map and vice versa
	//   - negative numbers to overflowed uint values (base 10)
	//   - slice of maps to a merged map
	//   - single values are converted to slices if required. Each
	//     element is weakly decoded. For example: "4" can become []int{4}
	//     if the target type is an int slice.
	//
	WeaklyTypedInput bool

	// Squash will squash embedded structs.  A squash tag may also be
	// added to an individual struct field using a tag.  For example:
	//
	//  type Parent struct {
	//      Child `mapstructure:",squash"`
	//  }
	Squash bool

	// Metadata is the struct that will contain extra metadata about
	// the decoding. If this is nil, then no metadata will be tracked.
	Metadata *Metadata

	// Result is a pointer to the struct that will contain the decoded
	// value.
	Result any

	// The tag name that mapstructure reads for field names. This
	// defaults to "mapstructure"
	TagName string

	// IgnoreUntaggedFields ignores all struct fields without explicit
	// TagName, comparable to `mapstructure:"-"` as default behaviour.
	IgnoreUntaggedFields bool

	// MatchName is the function used to match the map key to the struct
	// field name or tag. Defaults to `strings.EqualFold`. This can be used
	// to implement case-sensitive tag values, support snake casing, etc.
	MatchName func(__obf_2cca0ec99f95efa8, __obf_87f0d44838486dbb string) bool
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	__obf_9413fb8c4aa2a800 *DecoderConfig
}

// Metadata contains information about decoding a structure that
// is tedious or difficult to get otherwise.
type Metadata struct {
	// Keys are the keys of the structure which were successfully decoded
	Keys []string

	// Unused is a slice of keys that were found in the raw value but
	// weren't decoded since there was no matching field in the result interface
	Unused []string

	// Unset is a slice of field names that were found in the result interface
	// but weren't set in the decoding process since there was no matching value
	// in the input
	Unset []string
}

// Decode takes an input structure and uses reflection to translate it to
// the output structure. output must be a pointer to a map or struct.
func Decode(__obf_4d1394874cd54c98 any, __obf_8150817c27b926a9 any) error {
	__obf_9413fb8c4aa2a800 := &DecoderConfig{
		Metadata: nil,
		Result:   __obf_8150817c27b926a9,
	}
	__obf_95de370ca4bb3299, __obf_ce3ada6ea4dc41de := NewDecoder(__obf_9413fb8c4aa2a800)
	if __obf_ce3ada6ea4dc41de != nil {
		return __obf_ce3ada6ea4dc41de
	}

	return __obf_95de370ca4bb3299.Decode(__obf_4d1394874cd54c98)
}

// WeakDecode is the same as Decode but is shorthand to enable
// WeaklyTypedInput. See DecoderConfig for more info.
func WeakDecode(__obf_4d1394874cd54c98, __obf_8150817c27b926a9 any) error {
	__obf_9413fb8c4aa2a800 := &DecoderConfig{
		Metadata:         nil,
		Result:           __obf_8150817c27b926a9,
		WeaklyTypedInput: true,
	}
	__obf_95de370ca4bb3299, __obf_ce3ada6ea4dc41de := NewDecoder(__obf_9413fb8c4aa2a800)
	if __obf_ce3ada6ea4dc41de != nil {
		return __obf_ce3ada6ea4dc41de
	}

	return __obf_95de370ca4bb3299.Decode(__obf_4d1394874cd54c98)
}

// DecodeMetadata is the same as Decode, but is shorthand to
// enable metadata collection. See DecoderConfig for more info.
func DecodeMetadata(__obf_4d1394874cd54c98 any, __obf_8150817c27b926a9 any, __obf_76639a058ba13e2c *Metadata) error {
	__obf_9413fb8c4aa2a800 := &DecoderConfig{
		Metadata: __obf_76639a058ba13e2c,
		Result:   __obf_8150817c27b926a9,
	}
	__obf_95de370ca4bb3299, __obf_ce3ada6ea4dc41de := NewDecoder(__obf_9413fb8c4aa2a800)
	if __obf_ce3ada6ea4dc41de != nil {
		return __obf_ce3ada6ea4dc41de
	}

	return __obf_95de370ca4bb3299.Decode(__obf_4d1394874cd54c98)
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
// enable both WeaklyTypedInput and metadata collection. See
// DecoderConfig for more info.
func WeakDecodeMetadata(__obf_4d1394874cd54c98 any, __obf_8150817c27b926a9 any, __obf_76639a058ba13e2c *Metadata) error {
	__obf_9413fb8c4aa2a800 := &DecoderConfig{
		Metadata:         __obf_76639a058ba13e2c,
		Result:           __obf_8150817c27b926a9,
		WeaklyTypedInput: true,
	}
	__obf_95de370ca4bb3299, __obf_ce3ada6ea4dc41de := NewDecoder(__obf_9413fb8c4aa2a800)
	if __obf_ce3ada6ea4dc41de != nil {
		return __obf_ce3ada6ea4dc41de
	}

	return __obf_95de370ca4bb3299.Decode(__obf_4d1394874cd54c98)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(__obf_9413fb8c4aa2a800 *DecoderConfig) (*Decoder, error) {
	__obf_b7dec3cc8a6aba3c := reflect.ValueOf(__obf_9413fb8c4aa2a800.Result)
	if __obf_b7dec3cc8a6aba3c.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}
	__obf_b7dec3cc8a6aba3c = __obf_b7dec3cc8a6aba3c.Elem()
	if !__obf_b7dec3cc8a6aba3c.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if __obf_9413fb8c4aa2a800.Metadata != nil {
		if __obf_9413fb8c4aa2a800.Metadata.Keys == nil {
			__obf_9413fb8c4aa2a800.
				Metadata.Keys = make([]string, 0)
		}

		if __obf_9413fb8c4aa2a800.Metadata.Unused == nil {
			__obf_9413fb8c4aa2a800.
				Metadata.Unused = make([]string, 0)
		}

		if __obf_9413fb8c4aa2a800.Metadata.Unset == nil {
			__obf_9413fb8c4aa2a800.
				Metadata.Unset = make([]string, 0)
		}
	}

	if __obf_9413fb8c4aa2a800.TagName == "" {
		__obf_9413fb8c4aa2a800.
			TagName = "mapstructure"
	}

	if __obf_9413fb8c4aa2a800.MatchName == nil {
		__obf_9413fb8c4aa2a800.
			MatchName = strings.EqualFold
	}
	__obf_a9459758946d2eaa := &Decoder{__obf_9413fb8c4aa2a800: __obf_9413fb8c4aa2a800}

	return __obf_a9459758946d2eaa, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (__obf_b6d13e44c74c2560 *Decoder) Decode(__obf_4d1394874cd54c98 any) error {
	return __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35("", __obf_4d1394874cd54c98, reflect.ValueOf(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (__obf_b6d13e44c74c2560 *Decoder) __obf_eea77241dc8adb35(__obf_91e38c6481387fc0 string, __obf_4d1394874cd54c98 any, __obf_c1d382d8ffd17367 reflect.Value) error {
	var __obf_a1045a48a94db2e1 reflect.Value
	if __obf_4d1394874cd54c98 != nil {
		__obf_a1045a48a94db2e1 = reflect.ValueOf(__obf_4d1394874cd54c98)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if __obf_a1045a48a94db2e1.Kind() == reflect.Ptr && __obf_a1045a48a94db2e1.IsNil() {
			__obf_4d1394874cd54c98 = nil
		}
	}

	if __obf_4d1394874cd54c98 == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ZeroFields {
			__obf_c1d382d8ffd17367.
				Set(reflect.Zero(__obf_c1d382d8ffd17367.Type()))

			if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata != nil && __obf_91e38c6481387fc0 != "" {
				__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.
					Metadata.Keys = append(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata.Keys, __obf_91e38c6481387fc0)
			}
		}
		return nil
	}

	if !__obf_a1045a48a94db2e1.IsValid() {
		__obf_c1d382d8ffd17367.
			// If the input value is invalid, then we just set the value
			// to be the zero value.
			Set(reflect.Zero(__obf_c1d382d8ffd17367.Type()))
		if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata != nil && __obf_91e38c6481387fc0 != "" {
			__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.
				Metadata.Keys = append(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata.Keys, __obf_91e38c6481387fc0)
		}
		return nil
	}

	if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the input.
		var __obf_ce3ada6ea4dc41de error
		__obf_4d1394874cd54c98, __obf_ce3ada6ea4dc41de = DecodeHookExec(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.DecodeHook, __obf_a1045a48a94db2e1, __obf_c1d382d8ffd17367)
		if __obf_ce3ada6ea4dc41de != nil {
			return fmt.Errorf("error decoding '%s': %w", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
	}

	var __obf_ce3ada6ea4dc41de error
	__obf_e0d6c8e6f9160067 := __obf_32f26d0d12491636(__obf_c1d382d8ffd17367)
	__obf_124c651114f017c7 := true
	switch __obf_e0d6c8e6f9160067 {
	case reflect.Bool:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_aceb27a842250df2(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Interface:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_c9179d101ce685cf(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.String:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_aa0f8fa85e2621fc(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Int:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_7e3deca40624a63e(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Uint:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_c7f3188f58bad066(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Float32:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_139c2209e3bd3273(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Struct:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_0ea7f7e628e7a170(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Map:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_e0b8cf654a653244(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Ptr:
		__obf_124c651114f017c7, __obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_5e29a4de20dbe76f(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Slice:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_5830e3d709c533b9(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Array:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_4efb141f47e352d3(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98, __obf_c1d382d8ffd17367)
	case reflect.Func:
		__obf_ce3ada6ea4dc41de = __obf_b6d13e44c74c2560.__obf_03a933e715babe51(__obf_91e38c6481387fc0, __obf_4d1394874cd54c98,

			// If we reached this point then we weren't able to decode it
			__obf_c1d382d8ffd17367)
	default:

		return fmt.Errorf("%s: unsupported type: %s", __obf_91e38c6481387fc0, __obf_e0d6c8e6f9160067)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metainput.
	if __obf_124c651114f017c7 && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata != nil && __obf_91e38c6481387fc0 != "" {
		__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.
			Metadata.Keys = append(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata.Keys, __obf_91e38c6481387fc0)
	}

	return __obf_ce3ada6ea4dc41de
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (__obf_b6d13e44c74c2560 *Decoder) __obf_c9179d101ce685cf(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	if __obf_b7dec3cc8a6aba3c.IsValid() && __obf_b7dec3cc8a6aba3c.Elem().IsValid() {
		__obf_340beb890b8df74f := __obf_b7dec3cc8a6aba3c.Elem()
		__obf_9969d7b0333bcb09 := // If we can't address this element, then its not writable. Instead,
			// we make a copy of the value (which is a pointer and therefore
			// writable), decode into that, and replace the whole value.
			false
		if !__obf_340beb890b8df74f.CanAddr() {
			__obf_9969d7b0333bcb09 = true

			// Make *T
			copy := reflect.New(__obf_340beb890b8df74f.Type())

			// *T = elem
			copy.Elem().Set(__obf_340beb890b8df74f)
			__obf_340beb890b8df74f = // Set elem so we decode into it
				copy
		}

		// Decode. If we have an error then return. We also return right
		// away if we're not a copy because that means we decoded directly.
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_91e38c6481387fc0, __obf_4bc7fc453056091b, __obf_340beb890b8df74f); __obf_ce3ada6ea4dc41de != nil || !__obf_9969d7b0333bcb09 {
			return __obf_ce3ada6ea4dc41de
		}
		__obf_b7dec3cc8a6aba3c.

			// If we're a copy, we need to set te final result
			Set(__obf_340beb890b8df74f.Elem())
		return nil
	}
	__obf_0bfc5beb6a20600c := reflect.ValueOf(__obf_4bc7fc453056091b)

	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if __obf_0bfc5beb6a20600c.Kind() == reflect.Ptr && __obf_0bfc5beb6a20600c.Type().Elem() == __obf_b7dec3cc8a6aba3c.Type() {
		__obf_0bfc5beb6a20600c = reflect.Indirect(__obf_0bfc5beb6a20600c)
	}

	if !__obf_0bfc5beb6a20600c.IsValid() {
		__obf_0bfc5beb6a20600c = reflect.Zero(__obf_b7dec3cc8a6aba3c.Type())
	}
	__obf_ccfb61cfb66ec1f2 := __obf_0bfc5beb6a20600c.Type()
	if !__obf_ccfb61cfb66ec1f2.AssignableTo(__obf_b7dec3cc8a6aba3c.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_ccfb61cfb66ec1f2)
	}
	__obf_b7dec3cc8a6aba3c.
		Set(__obf_0bfc5beb6a20600c)
	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_aa0f8fa85e2621fc(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_4681297d649eea44 := __obf_32f26d0d12491636(__obf_0bfc5beb6a20600c)
	__obf_5e0accca5a88f1fa := true
	switch {
	case __obf_4681297d649eea44 == reflect.String:
		__obf_b7dec3cc8a6aba3c.
			SetString(__obf_0bfc5beb6a20600c.String())
	case __obf_4681297d649eea44 == reflect.Bool && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		if __obf_0bfc5beb6a20600c.Bool() {
			__obf_b7dec3cc8a6aba3c.
				SetString("1")
		} else {
			__obf_b7dec3cc8a6aba3c.
				SetString("0")
		}
	case __obf_4681297d649eea44 == reflect.Int && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_b7dec3cc8a6aba3c.
			SetString(strconv.FormatInt(__obf_0bfc5beb6a20600c.Int(), 10))
	case __obf_4681297d649eea44 == reflect.Uint && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_b7dec3cc8a6aba3c.
			SetString(strconv.FormatUint(__obf_0bfc5beb6a20600c.Uint(), 10))
	case __obf_4681297d649eea44 == reflect.Float32 && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_b7dec3cc8a6aba3c.
			SetString(strconv.FormatFloat(__obf_0bfc5beb6a20600c.Float(), 'f', -1, 64))
	case __obf_4681297d649eea44 == reflect.Slice && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput, __obf_4681297d649eea44 ==
		reflect.Array && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_cb1be79795600114 := __obf_0bfc5beb6a20600c.Type()
		__obf_55286990c7ae804f := __obf_cb1be79795600114.Elem().Kind()
		switch __obf_55286990c7ae804f {
		case reflect.Uint8:
			var __obf_689b9a8c4a647489 []uint8
			if __obf_4681297d649eea44 == reflect.Array {
				__obf_689b9a8c4a647489 = make([]uint8, __obf_0bfc5beb6a20600c.Len())
				for __obf_4e5b08e68a148355 := range __obf_689b9a8c4a647489 {
					__obf_689b9a8c4a647489[__obf_4e5b08e68a148355] = __obf_0bfc5beb6a20600c.Index(__obf_4e5b08e68a148355).Interface().(uint8)
				}
			} else {
				__obf_689b9a8c4a647489 = __obf_0bfc5beb6a20600c.Interface().([]uint8)
			}
			__obf_b7dec3cc8a6aba3c.
				SetString(string(__obf_689b9a8c4a647489))
		default:
			__obf_5e0accca5a88f1fa = false
		}
	default:
		__obf_5e0accca5a88f1fa = false
	}

	if !__obf_5e0accca5a88f1fa {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_0bfc5beb6a20600c.Type(), __obf_4bc7fc453056091b)
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_7e3deca40624a63e(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_4681297d649eea44 := __obf_32f26d0d12491636(__obf_0bfc5beb6a20600c)
	__obf_cb1be79795600114 := __obf_0bfc5beb6a20600c.Type()

	switch {
	case __obf_4681297d649eea44 == reflect.Int:
		__obf_b7dec3cc8a6aba3c.
			SetInt(__obf_0bfc5beb6a20600c.Int())
	case __obf_4681297d649eea44 == reflect.Uint:
		__obf_b7dec3cc8a6aba3c.
			SetInt(int64(__obf_0bfc5beb6a20600c.Uint()))
	case __obf_4681297d649eea44 == reflect.Float32:
		__obf_b7dec3cc8a6aba3c.
			SetInt(int64(__obf_0bfc5beb6a20600c.Float()))
	case __obf_4681297d649eea44 == reflect.Bool && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		if __obf_0bfc5beb6a20600c.Bool() {
			__obf_b7dec3cc8a6aba3c.
				SetInt(1)
		} else {
			__obf_b7dec3cc8a6aba3c.
				SetInt(0)
		}
	case __obf_4681297d649eea44 == reflect.String && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_288b09ee2b024465 := __obf_0bfc5beb6a20600c.String()
		if __obf_288b09ee2b024465 == "" {
			__obf_288b09ee2b024465 = "0"
		}
		__obf_4e5b08e68a148355, __obf_ce3ada6ea4dc41de := strconv.ParseInt(__obf_288b09ee2b024465, 0, __obf_b7dec3cc8a6aba3c.Type().Bits())
		if __obf_ce3ada6ea4dc41de == nil {
			__obf_b7dec3cc8a6aba3c.
				SetInt(__obf_4e5b08e68a148355)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
	case __obf_cb1be79795600114.PkgPath() == "encoding/json" && __obf_cb1be79795600114.Name() == "Number":
		__obf_5e82faecd5a1fa26 := __obf_4bc7fc453056091b.(json.Number)
		__obf_4e5b08e68a148355, __obf_ce3ada6ea4dc41de := __obf_5e82faecd5a1fa26.Int64()
		if __obf_ce3ada6ea4dc41de != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
		__obf_b7dec3cc8a6aba3c.
			SetInt(__obf_4e5b08e68a148355)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_0bfc5beb6a20600c.Type(), __obf_4bc7fc453056091b)
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_c7f3188f58bad066(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_4681297d649eea44 := __obf_32f26d0d12491636(__obf_0bfc5beb6a20600c)
	__obf_cb1be79795600114 := __obf_0bfc5beb6a20600c.Type()

	switch {
	case __obf_4681297d649eea44 == reflect.Int:
		__obf_4e5b08e68a148355 := __obf_0bfc5beb6a20600c.Int()
		if __obf_4e5b08e68a148355 < 0 && !__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %d overflows uint", __obf_91e38c6481387fc0, __obf_4e5b08e68a148355)
		}
		__obf_b7dec3cc8a6aba3c.
			SetUint(uint64(__obf_4e5b08e68a148355))
	case __obf_4681297d649eea44 == reflect.Uint:
		__obf_b7dec3cc8a6aba3c.
			SetUint(__obf_0bfc5beb6a20600c.Uint())
	case __obf_4681297d649eea44 == reflect.Float32:
		__obf_7bfc349763f12493 := __obf_0bfc5beb6a20600c.Float()
		if __obf_7bfc349763f12493 < 0 && !__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %f overflows uint", __obf_91e38c6481387fc0, __obf_7bfc349763f12493)
		}
		__obf_b7dec3cc8a6aba3c.
			SetUint(uint64(__obf_7bfc349763f12493))
	case __obf_4681297d649eea44 == reflect.Bool && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		if __obf_0bfc5beb6a20600c.Bool() {
			__obf_b7dec3cc8a6aba3c.
				SetUint(1)
		} else {
			__obf_b7dec3cc8a6aba3c.
				SetUint(0)
		}
	case __obf_4681297d649eea44 == reflect.String && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_288b09ee2b024465 := __obf_0bfc5beb6a20600c.String()
		if __obf_288b09ee2b024465 == "" {
			__obf_288b09ee2b024465 = "0"
		}
		__obf_4e5b08e68a148355, __obf_ce3ada6ea4dc41de := strconv.ParseUint(__obf_288b09ee2b024465, 0, __obf_b7dec3cc8a6aba3c.Type().Bits())
		if __obf_ce3ada6ea4dc41de == nil {
			__obf_b7dec3cc8a6aba3c.
				SetUint(__obf_4e5b08e68a148355)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
	case __obf_cb1be79795600114.PkgPath() == "encoding/json" && __obf_cb1be79795600114.Name() == "Number":
		__obf_5e82faecd5a1fa26 := __obf_4bc7fc453056091b.(json.Number)
		__obf_4e5b08e68a148355, __obf_ce3ada6ea4dc41de := strconv.ParseUint(string(__obf_5e82faecd5a1fa26), 0, 64)
		if __obf_ce3ada6ea4dc41de != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
		__obf_b7dec3cc8a6aba3c.
			SetUint(__obf_4e5b08e68a148355)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_0bfc5beb6a20600c.Type(), __obf_4bc7fc453056091b)
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_aceb27a842250df2(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_4681297d649eea44 := __obf_32f26d0d12491636(__obf_0bfc5beb6a20600c)

	switch {
	case __obf_4681297d649eea44 == reflect.Bool:
		__obf_b7dec3cc8a6aba3c.
			SetBool(__obf_0bfc5beb6a20600c.Bool())
	case __obf_4681297d649eea44 == reflect.Int && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_b7dec3cc8a6aba3c.
			SetBool(__obf_0bfc5beb6a20600c.Int() != 0)
	case __obf_4681297d649eea44 == reflect.Uint && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_b7dec3cc8a6aba3c.
			SetBool(__obf_0bfc5beb6a20600c.Uint() != 0)
	case __obf_4681297d649eea44 == reflect.Float32 && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_b7dec3cc8a6aba3c.
			SetBool(__obf_0bfc5beb6a20600c.Float() != 0)
	case __obf_4681297d649eea44 == reflect.String && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_2509c56426afe825, __obf_ce3ada6ea4dc41de := strconv.ParseBool(__obf_0bfc5beb6a20600c.String())
		if __obf_ce3ada6ea4dc41de == nil {
			__obf_b7dec3cc8a6aba3c.
				SetBool(__obf_2509c56426afe825)
		} else if __obf_0bfc5beb6a20600c.String() == "" {
			__obf_b7dec3cc8a6aba3c.
				SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_0bfc5beb6a20600c.Type(), __obf_4bc7fc453056091b)
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_139c2209e3bd3273(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_4681297d649eea44 := __obf_32f26d0d12491636(__obf_0bfc5beb6a20600c)
	__obf_cb1be79795600114 := __obf_0bfc5beb6a20600c.Type()

	switch {
	case __obf_4681297d649eea44 == reflect.Int:
		__obf_b7dec3cc8a6aba3c.
			SetFloat(float64(__obf_0bfc5beb6a20600c.Int()))
	case __obf_4681297d649eea44 == reflect.Uint:
		__obf_b7dec3cc8a6aba3c.
			SetFloat(float64(__obf_0bfc5beb6a20600c.Uint()))
	case __obf_4681297d649eea44 == reflect.Float32:
		__obf_b7dec3cc8a6aba3c.
			SetFloat(__obf_0bfc5beb6a20600c.Float())
	case __obf_4681297d649eea44 == reflect.Bool && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		if __obf_0bfc5beb6a20600c.Bool() {
			__obf_b7dec3cc8a6aba3c.
				SetFloat(1)
		} else {
			__obf_b7dec3cc8a6aba3c.
				SetFloat(0)
		}
	case __obf_4681297d649eea44 == reflect.String && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput:
		__obf_288b09ee2b024465 := __obf_0bfc5beb6a20600c.String()
		if __obf_288b09ee2b024465 == "" {
			__obf_288b09ee2b024465 = "0"
		}
		__obf_7bfc349763f12493, __obf_ce3ada6ea4dc41de := strconv.ParseFloat(__obf_288b09ee2b024465, __obf_b7dec3cc8a6aba3c.Type().Bits())
		if __obf_ce3ada6ea4dc41de == nil {
			__obf_b7dec3cc8a6aba3c.
				SetFloat(__obf_7bfc349763f12493)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
	case __obf_cb1be79795600114.PkgPath() == "encoding/json" && __obf_cb1be79795600114.Name() == "Number":
		__obf_5e82faecd5a1fa26 := __obf_4bc7fc453056091b.(json.Number)
		__obf_4e5b08e68a148355, __obf_ce3ada6ea4dc41de := __obf_5e82faecd5a1fa26.Float64()
		if __obf_ce3ada6ea4dc41de != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_91e38c6481387fc0, __obf_ce3ada6ea4dc41de)
		}
		__obf_b7dec3cc8a6aba3c.
			SetFloat(__obf_4e5b08e68a148355)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_0bfc5beb6a20600c.Type(), __obf_4bc7fc453056091b)
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_e0b8cf654a653244(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_ff3a378dc2f0e275 := __obf_b7dec3cc8a6aba3c.Type()
	__obf_342201c2a86c923f := __obf_ff3a378dc2f0e275.Key()
	__obf_6117515577845688 := __obf_ff3a378dc2f0e275.Elem()
	__obf_0f1a01daa73e7a35 := // By default we overwrite keys in the current map
		__obf_b7dec3cc8a6aba3c

	// If the map is nil or we're purposely zeroing fields, make a new map
	if __obf_0f1a01daa73e7a35.IsNil() || __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ZeroFields {
		__obf_c134280a495c4576 := // Make a new map to hold our result
			reflect.MapOf(__obf_342201c2a86c923f, __obf_6117515577845688)
		__obf_0f1a01daa73e7a35 = reflect.MakeMap(__obf_c134280a495c4576)
	}
	__obf_0bfc5beb6a20600c := // Check input type and based on the input type jump to the proper func
		reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	switch __obf_0bfc5beb6a20600c.Kind() {
	case reflect.Map:
		return __obf_b6d13e44c74c2560.__obf_48991e54d94b3725(__obf_91e38c6481387fc0, __obf_0bfc5beb6a20600c, __obf_b7dec3cc8a6aba3c, __obf_0f1a01daa73e7a35)

	case reflect.Struct:
		return __obf_b6d13e44c74c2560.__obf_9b1e964eafe7b32e(__obf_0bfc5beb6a20600c, __obf_b7dec3cc8a6aba3c, __obf_0f1a01daa73e7a35)

	case reflect.Array, reflect.Slice:
		if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput {
			return __obf_b6d13e44c74c2560.__obf_03fe305b5722e018(__obf_91e38c6481387fc0, __obf_0bfc5beb6a20600c, __obf_b7dec3cc8a6aba3c, __obf_0f1a01daa73e7a35)
		}

		fallthrough

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_91e38c6481387fc0, __obf_0bfc5beb6a20600c.Kind())
	}
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_03fe305b5722e018(__obf_91e38c6481387fc0 string, __obf_0bfc5beb6a20600c reflect.Value, __obf_b7dec3cc8a6aba3c reflect.Value, __obf_0f1a01daa73e7a35 reflect.Value) error {
	// Special case for BC reasons (covered by tests)
	if __obf_0bfc5beb6a20600c.Len() == 0 {
		__obf_b7dec3cc8a6aba3c.
			Set(__obf_0f1a01daa73e7a35)
		return nil
	}

	for __obf_4e5b08e68a148355 := 0; __obf_4e5b08e68a148355 < __obf_0bfc5beb6a20600c.Len(); __obf_4e5b08e68a148355++ {
		__obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_91e38c6481387fc0+
			"["+strconv.Itoa(__obf_4e5b08e68a148355)+"]", __obf_0bfc5beb6a20600c.
			Index(__obf_4e5b08e68a148355).Interface(), __obf_b7dec3cc8a6aba3c)
		if __obf_ce3ada6ea4dc41de != nil {
			return __obf_ce3ada6ea4dc41de
		}
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_48991e54d94b3725(__obf_91e38c6481387fc0 string, __obf_0bfc5beb6a20600c reflect.Value, __obf_b7dec3cc8a6aba3c reflect.Value, __obf_0f1a01daa73e7a35 reflect.Value) error {
	__obf_ff3a378dc2f0e275 := __obf_b7dec3cc8a6aba3c.Type()
	__obf_342201c2a86c923f := __obf_ff3a378dc2f0e275.Key()
	__obf_6117515577845688 := __obf_ff3a378dc2f0e275.Elem()

	// Accumulate errors
	errors := make([]string, 0)

	// If the input data is empty, then we just match what the input data is.
	if __obf_0bfc5beb6a20600c.Len() == 0 {
		if __obf_0bfc5beb6a20600c.IsNil() {
			if !__obf_b7dec3cc8a6aba3c.IsNil() {
				__obf_b7dec3cc8a6aba3c.
					Set(__obf_0bfc5beb6a20600c)
			}
		} else {
			__obf_b7dec3cc8a6aba3c.
				// Set to empty allocated value
				Set(__obf_0f1a01daa73e7a35)
		}

		return nil
	}

	for _, __obf_cbac51807b02f04b := range __obf_0bfc5beb6a20600c.MapKeys() {
		__obf_87f0d44838486dbb := __obf_91e38c6481387fc0 + "[" + __obf_cbac51807b02f04b.String() + "]"
		__obf_75be26611dfc5936 := // First decode the key into the proper type
			reflect.Indirect(reflect.New(__obf_342201c2a86c923f))
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_87f0d44838486dbb, __obf_cbac51807b02f04b.Interface(), __obf_75be26611dfc5936); __obf_ce3ada6ea4dc41de != nil {
			errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
			continue
		}
		__obf_40a0aa2002a96053 := // Next decode the data into the proper type
			__obf_0bfc5beb6a20600c.MapIndex(__obf_cbac51807b02f04b).Interface()
		__obf_2d3a329b8189a532 := reflect.Indirect(reflect.New(__obf_6117515577845688))
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_87f0d44838486dbb, __obf_40a0aa2002a96053, __obf_2d3a329b8189a532); __obf_ce3ada6ea4dc41de != nil {
			errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
			continue
		}
		__obf_0f1a01daa73e7a35.
			SetMapIndex(__obf_75be26611dfc5936, __obf_2d3a329b8189a532)
	}
	__obf_b7dec3cc8a6aba3c.

		// Set the built up map to the value
		Set(__obf_0f1a01daa73e7a35)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_9b1e964eafe7b32e(__obf_0bfc5beb6a20600c reflect.Value, __obf_b7dec3cc8a6aba3c reflect.Value, __obf_0f1a01daa73e7a35 reflect.Value) error {
	__obf_b923d4b5912f1d6f := __obf_0bfc5beb6a20600c.Type()
	for __obf_4e5b08e68a148355 := 0; __obf_4e5b08e68a148355 < __obf_b923d4b5912f1d6f.NumField(); __obf_4e5b08e68a148355++ {
		__obf_7bfc349763f12493 := // Get the StructField first since this is a cheap operation. If the
			// field is unexported, then ignore it.
			__obf_b923d4b5912f1d6f.Field(__obf_4e5b08e68a148355)
		if __obf_7bfc349763f12493.PkgPath != "" {
			continue
		}
		__obf_40a0aa2002a96053 := // Next get the actual value of this field and verify it is assignable
			// to the map value.
			__obf_0bfc5beb6a20600c.Field(__obf_4e5b08e68a148355)
		if !__obf_40a0aa2002a96053.Type().AssignableTo(__obf_0f1a01daa73e7a35.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", __obf_40a0aa2002a96053.Type(), __obf_0f1a01daa73e7a35.Type().Elem())
		}
		__obf_6f0cf29468f87f16 := __obf_7bfc349763f12493.Tag.Get(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.TagName)
		__obf_a22eb71eb7839c93 := __obf_7bfc349763f12493.Name

		if __obf_6f0cf29468f87f16 == "" && __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.IgnoreUntaggedFields {
			continue
		}
		__obf_b55c7cfdc0f1b9d0 := // If Squash is set in the config, we squash the field down.
			__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Squash && __obf_40a0aa2002a96053.Kind() == reflect.Struct && __obf_7bfc349763f12493.Anonymous
		__obf_40a0aa2002a96053 = __obf_6a72a1d5381f9484(__obf_40a0aa2002a96053, __obf_b6d13e44c74c2560.

			// Determine the name of the key in the map
			__obf_9413fb8c4aa2a800.TagName)

		if __obf_49c2d697afa0d5e9 := strings.Index(__obf_6f0cf29468f87f16, ","); __obf_49c2d697afa0d5e9 != -1 {
			if __obf_6f0cf29468f87f16[:__obf_49c2d697afa0d5e9] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Contains(__obf_6f0cf29468f87f16[__obf_49c2d697afa0d5e9+1:], "omitempty") && __obf_ab306ff92a59bd1a(__obf_40a0aa2002a96053) {
				continue
			}
			__obf_b55c7cfdc0f1b9d0 = // If "squash" is specified in the tag, we squash the field down.
				__obf_b55c7cfdc0f1b9d0 || strings.Contains(__obf_6f0cf29468f87f16[__obf_49c2d697afa0d5e9+1:], "squash")
			if __obf_b55c7cfdc0f1b9d0 {
				// When squashing, the embedded type can be a pointer to a struct.
				if __obf_40a0aa2002a96053.Kind() == reflect.Ptr && __obf_40a0aa2002a96053.Elem().Kind() == reflect.Struct {
					__obf_40a0aa2002a96053 = __obf_40a0aa2002a96053.Elem()
				}

				// The final type must be a struct
				if __obf_40a0aa2002a96053.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", __obf_40a0aa2002a96053.Type())
				}
			}
			if __obf_e45a30a3c52938ef := __obf_6f0cf29468f87f16[:__obf_49c2d697afa0d5e9]; __obf_e45a30a3c52938ef != "" {
				__obf_a22eb71eb7839c93 = __obf_e45a30a3c52938ef
			}
		} else if len(__obf_6f0cf29468f87f16) > 0 {
			if __obf_6f0cf29468f87f16 == "-" {
				continue
			}
			__obf_a22eb71eb7839c93 = __obf_6f0cf29468f87f16
		}

		switch __obf_40a0aa2002a96053.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			__obf_b142638dfac3d013 := reflect.New(__obf_40a0aa2002a96053.Type())
			__obf_b142638dfac3d013.
				Elem().Set(__obf_40a0aa2002a96053)
			__obf_863562a74f402b0f := __obf_0f1a01daa73e7a35.Type()
			__obf_6c86852a8e09de7c := __obf_863562a74f402b0f.Key()
			__obf_1fe948a83100544b := __obf_863562a74f402b0f.Elem()
			__obf_3f58ce45e5ea278a := reflect.MapOf(__obf_6c86852a8e09de7c, __obf_1fe948a83100544b)
			__obf_cc76fec0ce35db9f := reflect.MakeMap(__obf_3f58ce45e5ea278a)
			__obf_3cc2790bca22b671 := // Creating a pointer to a map so that other methods can completely
				// overwrite the map if need be (looking at you decodeMapFromMap). The
				// indirection allows the underlying map to be settable (CanSet() == true)
				// where as reflect.MakeMap returns an unsettable map.
				reflect.New(__obf_cc76fec0ce35db9f.Type())
			reflect.Indirect(__obf_3cc2790bca22b671).Set(__obf_cc76fec0ce35db9f)
			__obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_a22eb71eb7839c93, __obf_b142638dfac3d013.Interface(), reflect.Indirect(__obf_3cc2790bca22b671))
			if __obf_ce3ada6ea4dc41de != nil {
				return __obf_ce3ada6ea4dc41de
			}
			__obf_cc76fec0ce35db9f = // the underlying map may have been completely overwritten so pull
				// it indirectly out of the enclosing value.
				reflect.Indirect(__obf_3cc2790bca22b671)

			if __obf_b55c7cfdc0f1b9d0 {
				for _, __obf_cbac51807b02f04b := range __obf_cc76fec0ce35db9f.MapKeys() {
					__obf_0f1a01daa73e7a35.
						SetMapIndex(__obf_cbac51807b02f04b, __obf_cc76fec0ce35db9f.MapIndex(__obf_cbac51807b02f04b))
				}
			} else {
				__obf_0f1a01daa73e7a35.
					SetMapIndex(reflect.ValueOf(__obf_a22eb71eb7839c93), __obf_cc76fec0ce35db9f)
			}

		default:
			__obf_0f1a01daa73e7a35.
				SetMapIndex(reflect.ValueOf(__obf_a22eb71eb7839c93), __obf_40a0aa2002a96053)
		}
	}

	if __obf_b7dec3cc8a6aba3c.CanAddr() {
		__obf_b7dec3cc8a6aba3c.
			Set(__obf_0f1a01daa73e7a35)
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_5e29a4de20dbe76f(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) (bool, error) {
	__obf_0bb843803aa02291 := // If the input data is nil, then we want to just set the output
		// pointer to be nil as well.
		__obf_4bc7fc453056091b == nil
	if !__obf_0bb843803aa02291 {
		switch __obf_40a0aa2002a96053 := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b)); __obf_40a0aa2002a96053.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			__obf_0bb843803aa02291 = __obf_40a0aa2002a96053.IsNil()
		}
	}
	if __obf_0bb843803aa02291 {
		if !__obf_b7dec3cc8a6aba3c.IsNil() && __obf_b7dec3cc8a6aba3c.CanSet() {
			__obf_d6ce16d271686428 := reflect.New(__obf_b7dec3cc8a6aba3c.Type()).Elem()
			__obf_b7dec3cc8a6aba3c.
				Set(__obf_d6ce16d271686428)
		}

		return true, nil
	}
	__obf_ff3a378dc2f0e275 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		__obf_b7dec3cc8a6aba3c.Type()
	__obf_6117515577845688 := __obf_ff3a378dc2f0e275.Elem()
	if __obf_b7dec3cc8a6aba3c.CanSet() {
		__obf_5b1513df52ac39f3 := __obf_b7dec3cc8a6aba3c
		if __obf_5b1513df52ac39f3.IsNil() || __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ZeroFields {
			__obf_5b1513df52ac39f3 = reflect.New(__obf_6117515577845688)
		}

		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_91e38c6481387fc0, __obf_4bc7fc453056091b, reflect.Indirect(__obf_5b1513df52ac39f3)); __obf_ce3ada6ea4dc41de != nil {
			return false, __obf_ce3ada6ea4dc41de
		}
		__obf_b7dec3cc8a6aba3c.
			Set(__obf_5b1513df52ac39f3)
	} else {
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_91e38c6481387fc0, __obf_4bc7fc453056091b, reflect.Indirect(__obf_b7dec3cc8a6aba3c)); __obf_ce3ada6ea4dc41de != nil {
			return false, __obf_ce3ada6ea4dc41de
		}
	}
	return false, nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_03a933e715babe51(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	if __obf_b7dec3cc8a6aba3c.Type() != __obf_0bfc5beb6a20600c.Type() {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_91e38c6481387fc0, __obf_b7dec3cc8a6aba3c.
				Type(), __obf_0bfc5beb6a20600c.Type(), __obf_4bc7fc453056091b)
	}
	__obf_b7dec3cc8a6aba3c.
		Set(__obf_0bfc5beb6a20600c)
	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_5830e3d709c533b9(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_22e7b1e2803215a9 := __obf_0bfc5beb6a20600c.Kind()
	__obf_ff3a378dc2f0e275 := __obf_b7dec3cc8a6aba3c.Type()
	__obf_6117515577845688 := __obf_ff3a378dc2f0e275.Elem()
	__obf_8a207f060f0969bb := reflect.SliceOf(__obf_6117515577845688)

	// If we have a non array/slice type then we first attempt to convert.
	if __obf_22e7b1e2803215a9 != reflect.Array && __obf_22e7b1e2803215a9 != reflect.Slice {
		if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case __obf_22e7b1e2803215a9 == reflect.Slice, __obf_22e7b1e2803215a9 == reflect.Array:
				break

			// Empty maps turn into empty slices
			case __obf_22e7b1e2803215a9 == reflect.Map:
				if __obf_0bfc5beb6a20600c.Len() == 0 {
					__obf_b7dec3cc8a6aba3c.
						Set(reflect.MakeSlice(__obf_8a207f060f0969bb, 0, 0))
					return nil
				}
				// Create slice of maps of other sizes
				return __obf_b6d13e44c74c2560.__obf_5830e3d709c533b9(__obf_91e38c6481387fc0, []any{__obf_4bc7fc453056091b}, __obf_b7dec3cc8a6aba3c)

			case __obf_22e7b1e2803215a9 == reflect.String && __obf_6117515577845688.Kind() == reflect.Uint8:
				return __obf_b6d13e44c74c2560.__obf_5830e3d709c533b9(__obf_91e38c6481387fc0, []byte(__obf_0bfc5beb6a20600c.String()), __obf_b7dec3cc8a6aba3c)

			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				return __obf_b6d13e44c74c2560.__obf_5830e3d709c533b9(__obf_91e38c6481387fc0, []any{__obf_4bc7fc453056091b}, __obf_b7dec3cc8a6aba3c)
			}
		}

		return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", __obf_91e38c6481387fc0, __obf_22e7b1e2803215a9)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if __obf_22e7b1e2803215a9 != reflect.Array && __obf_0bfc5beb6a20600c.IsNil() {
		return nil
	}
	__obf_8db1509f1a548a01 := __obf_b7dec3cc8a6aba3c
	if __obf_8db1509f1a548a01.IsNil() || __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ZeroFields {
		__obf_8db1509f1a548a01 = // Make a new slice to hold our result, same size as the original data.
			reflect.MakeSlice(__obf_8a207f060f0969bb, __obf_0bfc5beb6a20600c.Len(), __obf_0bfc5beb6a20600c.Len())
	} else if __obf_8db1509f1a548a01.Len() > __obf_0bfc5beb6a20600c.Len() {
		__obf_8db1509f1a548a01 = __obf_8db1509f1a548a01.Slice(0, __obf_0bfc5beb6a20600c.Len())
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_4e5b08e68a148355 := 0; __obf_4e5b08e68a148355 < __obf_0bfc5beb6a20600c.Len(); __obf_4e5b08e68a148355++ {
		__obf_c1d57c0ecd71e4da := __obf_0bfc5beb6a20600c.Index(__obf_4e5b08e68a148355).Interface()
		for __obf_8db1509f1a548a01.Len() <= __obf_4e5b08e68a148355 {
			__obf_8db1509f1a548a01 = reflect.Append(__obf_8db1509f1a548a01, reflect.Zero(__obf_6117515577845688))
		}
		__obf_7c7580d5bba4812d := __obf_8db1509f1a548a01.Index(__obf_4e5b08e68a148355)
		__obf_87f0d44838486dbb := __obf_91e38c6481387fc0 + "[" + strconv.Itoa(__obf_4e5b08e68a148355) + "]"
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_87f0d44838486dbb, __obf_c1d57c0ecd71e4da, __obf_7c7580d5bba4812d); __obf_ce3ada6ea4dc41de != nil {
			errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
		}
	}
	__obf_b7dec3cc8a6aba3c.

		// Finally, set the value to the slice we built up
		Set(__obf_8db1509f1a548a01)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_4efb141f47e352d3(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))
	__obf_22e7b1e2803215a9 := __obf_0bfc5beb6a20600c.Kind()
	__obf_ff3a378dc2f0e275 := __obf_b7dec3cc8a6aba3c.Type()
	__obf_6117515577845688 := __obf_ff3a378dc2f0e275.Elem()
	__obf_a53ae4bc6e17a0f1 := reflect.ArrayOf(__obf_ff3a378dc2f0e275.Len(), __obf_6117515577845688)
	__obf_c586c435150030a0 := __obf_b7dec3cc8a6aba3c

	if __obf_c586c435150030a0.Interface() == reflect.Zero(__obf_c586c435150030a0.Type()).Interface() || __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ZeroFields {
		// Check input type
		if __obf_22e7b1e2803215a9 != reflect.Array && __obf_22e7b1e2803215a9 != reflect.Slice {
			if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.WeaklyTypedInput {
				switch __obf_22e7b1e2803215a9 {
				// Empty maps turn into empty arrays
				case reflect.Map:
					if __obf_0bfc5beb6a20600c.Len() == 0 {
						__obf_b7dec3cc8a6aba3c.
							Set(reflect.Zero(__obf_a53ae4bc6e17a0f1))
						return nil
					}

				// All other types we try to convert to the array type
				// and "lift" it into it. i.e. a string becomes a string array.
				default:
					// Just re-try this function with data as a slice.
					return __obf_b6d13e44c74c2560.__obf_4efb141f47e352d3(__obf_91e38c6481387fc0, []any{__obf_4bc7fc453056091b}, __obf_b7dec3cc8a6aba3c)
				}
			}

			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", __obf_91e38c6481387fc0, __obf_22e7b1e2803215a9)

		}
		if __obf_0bfc5beb6a20600c.Len() > __obf_a53ae4bc6e17a0f1.Len() {
			return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", __obf_91e38c6481387fc0, __obf_a53ae4bc6e17a0f1.Len(), __obf_0bfc5beb6a20600c.Len())

		}
		__obf_c586c435150030a0 = // Make a new array to hold our result, same size as the original data.
			reflect.New(__obf_a53ae4bc6e17a0f1).Elem()
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_4e5b08e68a148355 := 0; __obf_4e5b08e68a148355 < __obf_0bfc5beb6a20600c.Len(); __obf_4e5b08e68a148355++ {
		__obf_c1d57c0ecd71e4da := __obf_0bfc5beb6a20600c.Index(__obf_4e5b08e68a148355).Interface()
		__obf_7c7580d5bba4812d := __obf_c586c435150030a0.Index(__obf_4e5b08e68a148355)
		__obf_87f0d44838486dbb := __obf_91e38c6481387fc0 + "[" + strconv.Itoa(__obf_4e5b08e68a148355) + "]"
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_87f0d44838486dbb, __obf_c1d57c0ecd71e4da, __obf_7c7580d5bba4812d); __obf_ce3ada6ea4dc41de != nil {
			errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
		}
	}
	__obf_b7dec3cc8a6aba3c.

		// Finally, set the value to the array we built up
		Set(__obf_c586c435150030a0)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_0ea7f7e628e7a170(__obf_91e38c6481387fc0 string, __obf_4bc7fc453056091b any, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_0bfc5beb6a20600c := reflect.Indirect(reflect.ValueOf(__obf_4bc7fc453056091b))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if __obf_0bfc5beb6a20600c.Type() == __obf_b7dec3cc8a6aba3c.Type() {
		__obf_b7dec3cc8a6aba3c.
			Set(__obf_0bfc5beb6a20600c)
		return nil
	}
	__obf_22e7b1e2803215a9 := __obf_0bfc5beb6a20600c.Kind()
	switch __obf_22e7b1e2803215a9 {
	case reflect.Map:
		return __obf_b6d13e44c74c2560.__obf_b5507488e5ee3429(__obf_91e38c6481387fc0, __obf_0bfc5beb6a20600c,

			// Not the most efficient way to do this but we can optimize later if
			// we want to. To convert from struct to struct we go to map first
			// as an intermediary.
			__obf_b7dec3cc8a6aba3c)

	case reflect.Struct:
		__obf_c134280a495c4576 := // Make a new map to hold our result
			reflect.TypeOf((map[string]any)(nil))
		__obf_24b2139adb5ba068 := reflect.MakeMap(__obf_c134280a495c4576)
		__obf_3cc2790bca22b671 := // Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			reflect.New(__obf_24b2139adb5ba068.Type())

		reflect.Indirect(__obf_3cc2790bca22b671).Set(__obf_24b2139adb5ba068)
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_9b1e964eafe7b32e(__obf_0bfc5beb6a20600c, reflect.Indirect(__obf_3cc2790bca22b671), __obf_24b2139adb5ba068); __obf_ce3ada6ea4dc41de != nil {
			return __obf_ce3ada6ea4dc41de
		}
		__obf_a9459758946d2eaa := __obf_b6d13e44c74c2560.__obf_b5507488e5ee3429(__obf_91e38c6481387fc0, reflect.Indirect(__obf_3cc2790bca22b671), __obf_b7dec3cc8a6aba3c)
		return __obf_a9459758946d2eaa

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_91e38c6481387fc0, __obf_0bfc5beb6a20600c.Kind())
	}
}

func (__obf_b6d13e44c74c2560 *Decoder) __obf_b5507488e5ee3429(__obf_91e38c6481387fc0 string, __obf_0bfc5beb6a20600c, __obf_b7dec3cc8a6aba3c reflect.Value) error {
	__obf_ccfb61cfb66ec1f2 := __obf_0bfc5beb6a20600c.Type()
	if __obf_ab0df9454f2e3d74 := __obf_ccfb61cfb66ec1f2.Key().Kind(); __obf_ab0df9454f2e3d74 != reflect.String && __obf_ab0df9454f2e3d74 != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys", __obf_91e38c6481387fc0, __obf_ccfb61cfb66ec1f2.
				Key().Kind())
	}
	__obf_7a4440f446a4589f := make(map[reflect.Value]struct{})
	__obf_80839d986b68289f := make(map[any]struct{})
	for _, __obf_3bf9ca1dbf187d0a := range __obf_0bfc5beb6a20600c.MapKeys() {
		__obf_7a4440f446a4589f[__obf_3bf9ca1dbf187d0a] = struct{}{}
		__obf_80839d986b68289f[__obf_3bf9ca1dbf187d0a.Interface()] = struct{}{}
	}
	__obf_8b65c5d114b8b9cf := make(map[any]struct{})
	errors := make([]string, 0)
	__obf_be7fa424c964dadc := // This slice will keep track of all the structs we'll be decoding.
		// There can be more than one struct if there are embedded structs
		// that are squashed.
		make([]reflect.Value, 1, 5)
	__obf_be7fa424c964dadc[0] = __obf_b7dec3cc8a6aba3c

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type __obf_7dde8e9af4dbb643 struct {
		__obf_7dde8e9af4dbb643 reflect.StructField
		__obf_b7dec3cc8a6aba3c reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var __obf_b6f11cd0fb030f21 *__obf_7dde8e9af4dbb643
	__obf_66381feacf472cc9 := []__obf_7dde8e9af4dbb643{}
	for len(__obf_be7fa424c964dadc) > 0 {
		__obf_3a79a889b3305ae3 := __obf_be7fa424c964dadc[0]
		__obf_be7fa424c964dadc = __obf_be7fa424c964dadc[1:]
		__obf_27bdb95bef79cd84 := __obf_3a79a889b3305ae3.Type()

		for __obf_4e5b08e68a148355 := 0; __obf_4e5b08e68a148355 < __obf_27bdb95bef79cd84.NumField(); __obf_4e5b08e68a148355++ {
			__obf_66820d1d17760105 := __obf_27bdb95bef79cd84.Field(__obf_4e5b08e68a148355)
			__obf_eed1d158ce0191d7 := __obf_3a79a889b3305ae3.Field(__obf_4e5b08e68a148355)
			if __obf_eed1d158ce0191d7.Kind() == reflect.Pointer && __obf_eed1d158ce0191d7.Elem().Kind() == reflect.Struct {
				__obf_eed1d158ce0191d7 = // Handle embedded struct pointers as embedded structs.
					__obf_eed1d158ce0191d7.Elem()
			}
			__obf_b55c7cfdc0f1b9d0 := // If "squash" is specified in the tag, we squash the field down.
				__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Squash && __obf_eed1d158ce0191d7.Kind() == reflect.Struct && __obf_66820d1d17760105.Anonymous
			__obf_5edc804f1fa565ef := false
			__obf_7a700357c560dc78 := // We always parse the tags cause we're looking for other tags too
				strings.Split(__obf_66820d1d17760105.Tag.Get(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.TagName), ",")
			for _, __obf_84f128febf9cad68 := range __obf_7a700357c560dc78[1:] {
				if __obf_84f128febf9cad68 == "squash" {
					__obf_b55c7cfdc0f1b9d0 = true
					break
				}

				if __obf_84f128febf9cad68 == "remain" {
					__obf_5edc804f1fa565ef = true
					break
				}
			}

			if __obf_b55c7cfdc0f1b9d0 {
				if __obf_eed1d158ce0191d7.Kind() != reflect.Struct {
					errors = __obf_8cd4c4c9f4ceb7f3(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", __obf_66820d1d17760105.Name, __obf_eed1d158ce0191d7.Kind()))
				} else {
					__obf_be7fa424c964dadc = append(__obf_be7fa424c964dadc, __obf_eed1d158ce0191d7)
				}
				continue
			}

			// Build our field
			if __obf_5edc804f1fa565ef {
				__obf_b6f11cd0fb030f21 = &__obf_7dde8e9af4dbb643{__obf_66820d1d17760105, __obf_eed1d158ce0191d7}
			} else {
				__obf_66381feacf472cc9 = // Normal struct field, store it away
					append(__obf_66381feacf472cc9, __obf_7dde8e9af4dbb643{__obf_66820d1d17760105, __obf_eed1d158ce0191d7})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, __obf_7bfc349763f12493 := range __obf_66381feacf472cc9 {
		__obf_7dde8e9af4dbb643, __obf_a31a3d4797680790 := __obf_7bfc349763f12493.__obf_7dde8e9af4dbb643, __obf_7bfc349763f12493.__obf_b7dec3cc8a6aba3c
		__obf_87f0d44838486dbb := __obf_7dde8e9af4dbb643.Name
		__obf_6f0cf29468f87f16 := __obf_7dde8e9af4dbb643.Tag.Get(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.TagName)
		__obf_6f0cf29468f87f16 = strings.SplitN(__obf_6f0cf29468f87f16, ",", 2)[0]
		if __obf_6f0cf29468f87f16 != "" {
			__obf_87f0d44838486dbb = __obf_6f0cf29468f87f16
		}
		__obf_d3e8630dee20bb46 := reflect.ValueOf(__obf_87f0d44838486dbb)
		__obf_2cabc602fbcf7a2c := __obf_0bfc5beb6a20600c.MapIndex(__obf_d3e8630dee20bb46)
		if !__obf_2cabc602fbcf7a2c.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for __obf_3bf9ca1dbf187d0a := range __obf_7a4440f446a4589f {
				__obf_a16878cc1b6ba613, __obf_5006c1b1bc16f11d := __obf_3bf9ca1dbf187d0a.Interface().(string)
				if !__obf_5006c1b1bc16f11d {
					// Not a string key
					continue
				}

				if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.MatchName(__obf_a16878cc1b6ba613, __obf_87f0d44838486dbb) {
					__obf_d3e8630dee20bb46 = __obf_3bf9ca1dbf187d0a
					__obf_2cabc602fbcf7a2c = __obf_0bfc5beb6a20600c.MapIndex(__obf_3bf9ca1dbf187d0a)
					break
				}
			}

			if !__obf_2cabc602fbcf7a2c.IsValid() {
				__obf_8b65c5d114b8b9cf[ // There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				__obf_87f0d44838486dbb] = struct{}{}
				continue
			}
		}

		if !__obf_a31a3d4797680790.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !__obf_a31a3d4797680790.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(__obf_80839d986b68289f, __obf_d3e8630dee20bb46.Interface())

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if __obf_91e38c6481387fc0 != "" {
			__obf_87f0d44838486dbb = __obf_91e38c6481387fc0 + "." + __obf_87f0d44838486dbb
		}

		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_eea77241dc8adb35(__obf_87f0d44838486dbb, __obf_2cabc602fbcf7a2c.Interface(), __obf_a31a3d4797680790); __obf_ce3ada6ea4dc41de != nil {
			errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if __obf_b6f11cd0fb030f21 != nil && len(__obf_80839d986b68289f) > 0 {
		__obf_5edc804f1fa565ef := // Build a map of only the unused values
			map[any]any{}
		for __obf_517fee4efca8b611 := range __obf_80839d986b68289f {
			__obf_5edc804f1fa565ef[__obf_517fee4efca8b611] = __obf_0bfc5beb6a20600c.MapIndex(reflect.ValueOf(__obf_517fee4efca8b611)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if __obf_ce3ada6ea4dc41de := __obf_b6d13e44c74c2560.__obf_e0b8cf654a653244(__obf_91e38c6481387fc0, __obf_5edc804f1fa565ef, __obf_b6f11cd0fb030f21.__obf_b7dec3cc8a6aba3c); __obf_ce3ada6ea4dc41de != nil {
			errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
		}
		__obf_80839d986b68289f = // Set the map to nil so we have none so that the next check will
			// not error (ErrorUnused)
			nil
	}

	if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ErrorUnused && len(__obf_80839d986b68289f) > 0 {
		__obf_942a68f95456b784 := make([]string, 0, len(__obf_80839d986b68289f))
		for __obf_76bafb611bbadaac := range __obf_80839d986b68289f {
			__obf_942a68f95456b784 = append(__obf_942a68f95456b784, __obf_76bafb611bbadaac.(string))
		}
		sort.Strings(__obf_942a68f95456b784)
		__obf_ce3ada6ea4dc41de := fmt.Errorf("'%s' has invalid keys: %s", __obf_91e38c6481387fc0, strings.Join(__obf_942a68f95456b784, ", "))
		errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
	}

	if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.ErrorUnset && len(__obf_8b65c5d114b8b9cf) > 0 {
		__obf_942a68f95456b784 := make([]string, 0, len(__obf_8b65c5d114b8b9cf))
		for __obf_76bafb611bbadaac := range __obf_8b65c5d114b8b9cf {
			__obf_942a68f95456b784 = append(__obf_942a68f95456b784, __obf_76bafb611bbadaac.(string))
		}
		sort.Strings(__obf_942a68f95456b784)
		__obf_ce3ada6ea4dc41de := fmt.Errorf("'%s' has unset fields: %s", __obf_91e38c6481387fc0, strings.Join(__obf_942a68f95456b784, ", "))
		errors = __obf_8cd4c4c9f4ceb7f3(errors, __obf_ce3ada6ea4dc41de)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	if __obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata != nil {
		for __obf_76bafb611bbadaac := range __obf_80839d986b68289f {
			__obf_517fee4efca8b611 := __obf_76bafb611bbadaac.(string)
			if __obf_91e38c6481387fc0 != "" {
				__obf_517fee4efca8b611 = __obf_91e38c6481387fc0 + "." + __obf_517fee4efca8b611
			}
			__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.
				Metadata.Unused = append(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata.Unused, __obf_517fee4efca8b611)
		}
		for __obf_76bafb611bbadaac := range __obf_8b65c5d114b8b9cf {
			__obf_517fee4efca8b611 := __obf_76bafb611bbadaac.(string)
			if __obf_91e38c6481387fc0 != "" {
				__obf_517fee4efca8b611 = __obf_91e38c6481387fc0 + "." + __obf_517fee4efca8b611
			}
			__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.
				Metadata.Unset = append(__obf_b6d13e44c74c2560.__obf_9413fb8c4aa2a800.Metadata.Unset, __obf_517fee4efca8b611)
		}
	}

	return nil
}

func __obf_ab306ff92a59bd1a(__obf_40a0aa2002a96053 reflect.Value) bool {
	switch __obf_32f26d0d12491636(__obf_40a0aa2002a96053) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_40a0aa2002a96053.Len() == 0
	case reflect.Bool:
		return !__obf_40a0aa2002a96053.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_40a0aa2002a96053.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_40a0aa2002a96053.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_40a0aa2002a96053.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return __obf_40a0aa2002a96053.IsNil()
	}
	return false
}

func __obf_32f26d0d12491636(__obf_b7dec3cc8a6aba3c reflect.Value) reflect.Kind {
	__obf_ab0df9454f2e3d74 := __obf_b7dec3cc8a6aba3c.Kind()

	switch {
	case __obf_ab0df9454f2e3d74 >= reflect.Int && __obf_ab0df9454f2e3d74 <= reflect.Int64:
		return reflect.Int
	case __obf_ab0df9454f2e3d74 >= reflect.Uint && __obf_ab0df9454f2e3d74 <= reflect.Uint64:
		return reflect.Uint
	case __obf_ab0df9454f2e3d74 >= reflect.Float32 && __obf_ab0df9454f2e3d74 <= reflect.Float64:
		return reflect.Float32
	default:
		return __obf_ab0df9454f2e3d74
	}
}

func __obf_86db9ef32de43776(__obf_b923d4b5912f1d6f reflect.Type, __obf_091363c566e0ae44 bool, __obf_5b1733e4a1118e05 string) bool {
	for __obf_4e5b08e68a148355 := 0; __obf_4e5b08e68a148355 < __obf_b923d4b5912f1d6f.NumField(); __obf_4e5b08e68a148355++ {
		__obf_7bfc349763f12493 := __obf_b923d4b5912f1d6f.Field(__obf_4e5b08e68a148355)
		if __obf_7bfc349763f12493.PkgPath == "" && !__obf_091363c566e0ae44 { // check for unexported fields
			return true
		}
		if __obf_091363c566e0ae44 && __obf_7bfc349763f12493.Tag.Get(__obf_5b1733e4a1118e05) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func __obf_6a72a1d5381f9484(__obf_40a0aa2002a96053 reflect.Value, __obf_5b1733e4a1118e05 string) reflect.Value {
	if __obf_40a0aa2002a96053.Kind() != reflect.Ptr || __obf_40a0aa2002a96053.Elem().Kind() != reflect.Struct {
		return __obf_40a0aa2002a96053
	}
	__obf_e20aec6ec3723ceb := __obf_40a0aa2002a96053.Elem()
	__obf_05a81edf07e8c513 := __obf_e20aec6ec3723ceb.Type()
	if __obf_86db9ef32de43776(__obf_05a81edf07e8c513, true, __obf_5b1733e4a1118e05) {
		return __obf_e20aec6ec3723ceb
	}
	return __obf_40a0aa2002a96053
}
