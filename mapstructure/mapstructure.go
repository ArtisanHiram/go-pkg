package __obf_7bd99df3562420c2

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
type DecodeHookFuncValue func(__obf_9d6dd945270ddf52 reflect.Value, __obf_c3010b7e6da97940 reflect.Value) (any, error)

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
	MatchName func(__obf_167d56764cf6a3e2, __obf_5e2e0c0f08a5e072 string) bool
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	__obf_a5c9f567f7eba485 *DecoderConfig
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
func Decode(__obf_69e39cfe531580d7 any, __obf_7aa7445a39948a67 any) error {
	__obf_a5c9f567f7eba485 := &DecoderConfig{
		Metadata: nil,
		Result:   __obf_7aa7445a39948a67,
	}
	__obf_583e8da456fa678a, __obf_bf6c24b903ce000b := NewDecoder(__obf_a5c9f567f7eba485)
	if __obf_bf6c24b903ce000b != nil {
		return __obf_bf6c24b903ce000b
	}

	return __obf_583e8da456fa678a.Decode(__obf_69e39cfe531580d7)
}

// WeakDecode is the same as Decode but is shorthand to enable
// WeaklyTypedInput. See DecoderConfig for more info.
func WeakDecode(__obf_69e39cfe531580d7, __obf_7aa7445a39948a67 any) error {
	__obf_a5c9f567f7eba485 := &DecoderConfig{
		Metadata:         nil,
		Result:           __obf_7aa7445a39948a67,
		WeaklyTypedInput: true,
	}
	__obf_583e8da456fa678a, __obf_bf6c24b903ce000b := NewDecoder(__obf_a5c9f567f7eba485)
	if __obf_bf6c24b903ce000b != nil {
		return __obf_bf6c24b903ce000b
	}

	return __obf_583e8da456fa678a.Decode(__obf_69e39cfe531580d7)
}

// DecodeMetadata is the same as Decode, but is shorthand to
// enable metadata collection. See DecoderConfig for more info.
func DecodeMetadata(__obf_69e39cfe531580d7 any, __obf_7aa7445a39948a67 any, __obf_27d22473ef4fdff4 *Metadata) error {
	__obf_a5c9f567f7eba485 := &DecoderConfig{
		Metadata: __obf_27d22473ef4fdff4,
		Result:   __obf_7aa7445a39948a67,
	}
	__obf_583e8da456fa678a, __obf_bf6c24b903ce000b := NewDecoder(__obf_a5c9f567f7eba485)
	if __obf_bf6c24b903ce000b != nil {
		return __obf_bf6c24b903ce000b
	}

	return __obf_583e8da456fa678a.Decode(__obf_69e39cfe531580d7)
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
// enable both WeaklyTypedInput and metadata collection. See
// DecoderConfig for more info.
func WeakDecodeMetadata(__obf_69e39cfe531580d7 any, __obf_7aa7445a39948a67 any, __obf_27d22473ef4fdff4 *Metadata) error {
	__obf_a5c9f567f7eba485 := &DecoderConfig{
		Metadata:         __obf_27d22473ef4fdff4,
		Result:           __obf_7aa7445a39948a67,
		WeaklyTypedInput: true,
	}
	__obf_583e8da456fa678a, __obf_bf6c24b903ce000b := NewDecoder(__obf_a5c9f567f7eba485)
	if __obf_bf6c24b903ce000b != nil {
		return __obf_bf6c24b903ce000b
	}

	return __obf_583e8da456fa678a.Decode(__obf_69e39cfe531580d7)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(__obf_a5c9f567f7eba485 *DecoderConfig) (*Decoder, error) {
	__obf_d3c0db8b5b8d28e6 := reflect.ValueOf(__obf_a5c9f567f7eba485.Result)
	if __obf_d3c0db8b5b8d28e6.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}
	__obf_d3c0db8b5b8d28e6 = __obf_d3c0db8b5b8d28e6.Elem()
	if !__obf_d3c0db8b5b8d28e6.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if __obf_a5c9f567f7eba485.Metadata != nil {
		if __obf_a5c9f567f7eba485.Metadata.Keys == nil {
			__obf_a5c9f567f7eba485.
				Metadata.Keys = make([]string, 0)
		}

		if __obf_a5c9f567f7eba485.Metadata.Unused == nil {
			__obf_a5c9f567f7eba485.
				Metadata.Unused = make([]string, 0)
		}

		if __obf_a5c9f567f7eba485.Metadata.Unset == nil {
			__obf_a5c9f567f7eba485.
				Metadata.Unset = make([]string, 0)
		}
	}

	if __obf_a5c9f567f7eba485.TagName == "" {
		__obf_a5c9f567f7eba485.
			TagName = "mapstructure"
	}

	if __obf_a5c9f567f7eba485.MatchName == nil {
		__obf_a5c9f567f7eba485.
			MatchName = strings.EqualFold
	}
	__obf_6ad3d371cac0efc2 := &Decoder{__obf_a5c9f567f7eba485: __obf_a5c9f567f7eba485}

	return __obf_6ad3d371cac0efc2, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (__obf_11aa7e4430cc378c *Decoder) Decode(__obf_69e39cfe531580d7 any) error {
	return __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82("", __obf_69e39cfe531580d7, reflect.ValueOf(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (__obf_11aa7e4430cc378c *Decoder) __obf_cc469f327ca4db82(__obf_040f4b76f7e163f8 string, __obf_69e39cfe531580d7 any, __obf_895ef9d3a39a9a8b reflect.Value) error {
	var __obf_f60ca8363b69fa1e reflect.Value
	if __obf_69e39cfe531580d7 != nil {
		__obf_f60ca8363b69fa1e = reflect.ValueOf(__obf_69e39cfe531580d7)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if __obf_f60ca8363b69fa1e.Kind() == reflect.Ptr && __obf_f60ca8363b69fa1e.IsNil() {
			__obf_69e39cfe531580d7 = nil
		}
	}

	if __obf_69e39cfe531580d7 == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ZeroFields {
			__obf_895ef9d3a39a9a8b.
				Set(reflect.Zero(__obf_895ef9d3a39a9a8b.Type()))

			if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata != nil && __obf_040f4b76f7e163f8 != "" {
				__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.
					Metadata.Keys = append(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata.Keys, __obf_040f4b76f7e163f8)
			}
		}
		return nil
	}

	if !__obf_f60ca8363b69fa1e.IsValid() {
		__obf_895ef9d3a39a9a8b.
			// If the input value is invalid, then we just set the value
			// to be the zero value.
			Set(reflect.Zero(__obf_895ef9d3a39a9a8b.Type()))
		if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata != nil && __obf_040f4b76f7e163f8 != "" {
			__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.
				Metadata.Keys = append(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata.Keys, __obf_040f4b76f7e163f8)
		}
		return nil
	}

	if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the input.
		var __obf_bf6c24b903ce000b error
		__obf_69e39cfe531580d7, __obf_bf6c24b903ce000b = DecodeHookExec(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.DecodeHook, __obf_f60ca8363b69fa1e, __obf_895ef9d3a39a9a8b)
		if __obf_bf6c24b903ce000b != nil {
			return fmt.Errorf("error decoding '%s': %w", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
	}

	var __obf_bf6c24b903ce000b error
	__obf_b075635e041e4fa6 := __obf_045d45fad1f5a0bc(__obf_895ef9d3a39a9a8b)
	__obf_06397827fc04d351 := true
	switch __obf_b075635e041e4fa6 {
	case reflect.Bool:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_73877d8afa80af89(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Interface:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_407e9d3e4930ce71(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.String:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_3f1bff41167e8ed2(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Int:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_375232ad60284928(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Uint:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_2fb724672b12b32a(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Float32:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_ec26544fb5023887(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Struct:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_3a979d7c2b59d928(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Map:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_5236cf9478516e08(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Ptr:
		__obf_06397827fc04d351, __obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_9ee9af437245a47f(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Slice:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_4358644c7db5e853(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Array:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_28cbf385ef2a4ccf(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7, __obf_895ef9d3a39a9a8b)
	case reflect.Func:
		__obf_bf6c24b903ce000b = __obf_11aa7e4430cc378c.__obf_b9d4339190794a2e(__obf_040f4b76f7e163f8, __obf_69e39cfe531580d7,

			// If we reached this point then we weren't able to decode it
			__obf_895ef9d3a39a9a8b)
	default:

		return fmt.Errorf("%s: unsupported type: %s", __obf_040f4b76f7e163f8, __obf_b075635e041e4fa6)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metainput.
	if __obf_06397827fc04d351 && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata != nil && __obf_040f4b76f7e163f8 != "" {
		__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.
			Metadata.Keys = append(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata.Keys, __obf_040f4b76f7e163f8)
	}

	return __obf_bf6c24b903ce000b
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (__obf_11aa7e4430cc378c *Decoder) __obf_407e9d3e4930ce71(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	if __obf_d3c0db8b5b8d28e6.IsValid() && __obf_d3c0db8b5b8d28e6.Elem().IsValid() {
		__obf_7f88955c8319c68a := __obf_d3c0db8b5b8d28e6.Elem()
		__obf_0f32854f7a386c8b := // If we can't address this element, then its not writable. Instead,
			// we make a copy of the value (which is a pointer and therefore
			// writable), decode into that, and replace the whole value.
			false
		if !__obf_7f88955c8319c68a.CanAddr() {
			__obf_0f32854f7a386c8b = true

			// Make *T
			copy := reflect.New(__obf_7f88955c8319c68a.Type())

			// *T = elem
			copy.Elem().Set(__obf_7f88955c8319c68a)
			__obf_7f88955c8319c68a = // Set elem so we decode into it
				copy
		}

		// Decode. If we have an error then return. We also return right
		// away if we're not a copy because that means we decoded directly.
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_040f4b76f7e163f8, __obf_57ff3a2a2bd1a861, __obf_7f88955c8319c68a); __obf_bf6c24b903ce000b != nil || !__obf_0f32854f7a386c8b {
			return __obf_bf6c24b903ce000b
		}
		__obf_d3c0db8b5b8d28e6.

			// If we're a copy, we need to set te final result
			Set(__obf_7f88955c8319c68a.Elem())
		return nil
	}
	__obf_d633e91150cb7889 := reflect.ValueOf(__obf_57ff3a2a2bd1a861)

	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if __obf_d633e91150cb7889.Kind() == reflect.Ptr && __obf_d633e91150cb7889.Type().Elem() == __obf_d3c0db8b5b8d28e6.Type() {
		__obf_d633e91150cb7889 = reflect.Indirect(__obf_d633e91150cb7889)
	}

	if !__obf_d633e91150cb7889.IsValid() {
		__obf_d633e91150cb7889 = reflect.Zero(__obf_d3c0db8b5b8d28e6.Type())
	}
	__obf_e58ea9c91a850d24 := __obf_d633e91150cb7889.Type()
	if !__obf_e58ea9c91a850d24.AssignableTo(__obf_d3c0db8b5b8d28e6.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_e58ea9c91a850d24)
	}
	__obf_d3c0db8b5b8d28e6.
		Set(__obf_d633e91150cb7889)
	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_3f1bff41167e8ed2(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_afb9302665f844b5 := __obf_045d45fad1f5a0bc(__obf_d633e91150cb7889)
	__obf_d370ef177c9bebd3 := true
	switch {
	case __obf_afb9302665f844b5 == reflect.String:
		__obf_d3c0db8b5b8d28e6.
			SetString(__obf_d633e91150cb7889.String())
	case __obf_afb9302665f844b5 == reflect.Bool && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		if __obf_d633e91150cb7889.Bool() {
			__obf_d3c0db8b5b8d28e6.
				SetString("1")
		} else {
			__obf_d3c0db8b5b8d28e6.
				SetString("0")
		}
	case __obf_afb9302665f844b5 == reflect.Int && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_d3c0db8b5b8d28e6.
			SetString(strconv.FormatInt(__obf_d633e91150cb7889.Int(), 10))
	case __obf_afb9302665f844b5 == reflect.Uint && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_d3c0db8b5b8d28e6.
			SetString(strconv.FormatUint(__obf_d633e91150cb7889.Uint(), 10))
	case __obf_afb9302665f844b5 == reflect.Float32 && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_d3c0db8b5b8d28e6.
			SetString(strconv.FormatFloat(__obf_d633e91150cb7889.Float(), 'f', -1, 64))
	case __obf_afb9302665f844b5 == reflect.Slice && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput, __obf_afb9302665f844b5 ==
		reflect.Array && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_4ab72a4f9eed38d8 := __obf_d633e91150cb7889.Type()
		__obf_9dd35fb0b81a6a3a := __obf_4ab72a4f9eed38d8.Elem().Kind()
		switch __obf_9dd35fb0b81a6a3a {
		case reflect.Uint8:
			var __obf_d82f6596033a1315 []uint8
			if __obf_afb9302665f844b5 == reflect.Array {
				__obf_d82f6596033a1315 = make([]uint8, __obf_d633e91150cb7889.Len())
				for __obf_8d0ae84c7a3f3f55 := range __obf_d82f6596033a1315 {
					__obf_d82f6596033a1315[__obf_8d0ae84c7a3f3f55] = __obf_d633e91150cb7889.Index(__obf_8d0ae84c7a3f3f55).Interface().(uint8)
				}
			} else {
				__obf_d82f6596033a1315 = __obf_d633e91150cb7889.Interface().([]uint8)
			}
			__obf_d3c0db8b5b8d28e6.
				SetString(string(__obf_d82f6596033a1315))
		default:
			__obf_d370ef177c9bebd3 = false
		}
	default:
		__obf_d370ef177c9bebd3 = false
	}

	if !__obf_d370ef177c9bebd3 {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_d633e91150cb7889.Type(), __obf_57ff3a2a2bd1a861)
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_375232ad60284928(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_afb9302665f844b5 := __obf_045d45fad1f5a0bc(__obf_d633e91150cb7889)
	__obf_4ab72a4f9eed38d8 := __obf_d633e91150cb7889.Type()

	switch {
	case __obf_afb9302665f844b5 == reflect.Int:
		__obf_d3c0db8b5b8d28e6.
			SetInt(__obf_d633e91150cb7889.Int())
	case __obf_afb9302665f844b5 == reflect.Uint:
		__obf_d3c0db8b5b8d28e6.
			SetInt(int64(__obf_d633e91150cb7889.Uint()))
	case __obf_afb9302665f844b5 == reflect.Float32:
		__obf_d3c0db8b5b8d28e6.
			SetInt(int64(__obf_d633e91150cb7889.Float()))
	case __obf_afb9302665f844b5 == reflect.Bool && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		if __obf_d633e91150cb7889.Bool() {
			__obf_d3c0db8b5b8d28e6.
				SetInt(1)
		} else {
			__obf_d3c0db8b5b8d28e6.
				SetInt(0)
		}
	case __obf_afb9302665f844b5 == reflect.String && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_1264b3f46bd409c7 := __obf_d633e91150cb7889.String()
		if __obf_1264b3f46bd409c7 == "" {
			__obf_1264b3f46bd409c7 = "0"
		}
		__obf_8d0ae84c7a3f3f55, __obf_bf6c24b903ce000b := strconv.ParseInt(__obf_1264b3f46bd409c7, 0, __obf_d3c0db8b5b8d28e6.Type().Bits())
		if __obf_bf6c24b903ce000b == nil {
			__obf_d3c0db8b5b8d28e6.
				SetInt(__obf_8d0ae84c7a3f3f55)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
	case __obf_4ab72a4f9eed38d8.PkgPath() == "encoding/json" && __obf_4ab72a4f9eed38d8.Name() == "Number":
		__obf_669980db468f7544 := __obf_57ff3a2a2bd1a861.(json.Number)
		__obf_8d0ae84c7a3f3f55, __obf_bf6c24b903ce000b := __obf_669980db468f7544.Int64()
		if __obf_bf6c24b903ce000b != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
		__obf_d3c0db8b5b8d28e6.
			SetInt(__obf_8d0ae84c7a3f3f55)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_d633e91150cb7889.Type(), __obf_57ff3a2a2bd1a861)
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_2fb724672b12b32a(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_afb9302665f844b5 := __obf_045d45fad1f5a0bc(__obf_d633e91150cb7889)
	__obf_4ab72a4f9eed38d8 := __obf_d633e91150cb7889.Type()

	switch {
	case __obf_afb9302665f844b5 == reflect.Int:
		__obf_8d0ae84c7a3f3f55 := __obf_d633e91150cb7889.Int()
		if __obf_8d0ae84c7a3f3f55 < 0 && !__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %d overflows uint", __obf_040f4b76f7e163f8, __obf_8d0ae84c7a3f3f55)
		}
		__obf_d3c0db8b5b8d28e6.
			SetUint(uint64(__obf_8d0ae84c7a3f3f55))
	case __obf_afb9302665f844b5 == reflect.Uint:
		__obf_d3c0db8b5b8d28e6.
			SetUint(__obf_d633e91150cb7889.Uint())
	case __obf_afb9302665f844b5 == reflect.Float32:
		__obf_2df189427d80bcc8 := __obf_d633e91150cb7889.Float()
		if __obf_2df189427d80bcc8 < 0 && !__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %f overflows uint", __obf_040f4b76f7e163f8, __obf_2df189427d80bcc8)
		}
		__obf_d3c0db8b5b8d28e6.
			SetUint(uint64(__obf_2df189427d80bcc8))
	case __obf_afb9302665f844b5 == reflect.Bool && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		if __obf_d633e91150cb7889.Bool() {
			__obf_d3c0db8b5b8d28e6.
				SetUint(1)
		} else {
			__obf_d3c0db8b5b8d28e6.
				SetUint(0)
		}
	case __obf_afb9302665f844b5 == reflect.String && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_1264b3f46bd409c7 := __obf_d633e91150cb7889.String()
		if __obf_1264b3f46bd409c7 == "" {
			__obf_1264b3f46bd409c7 = "0"
		}
		__obf_8d0ae84c7a3f3f55, __obf_bf6c24b903ce000b := strconv.ParseUint(__obf_1264b3f46bd409c7, 0, __obf_d3c0db8b5b8d28e6.Type().Bits())
		if __obf_bf6c24b903ce000b == nil {
			__obf_d3c0db8b5b8d28e6.
				SetUint(__obf_8d0ae84c7a3f3f55)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
	case __obf_4ab72a4f9eed38d8.PkgPath() == "encoding/json" && __obf_4ab72a4f9eed38d8.Name() == "Number":
		__obf_669980db468f7544 := __obf_57ff3a2a2bd1a861.(json.Number)
		__obf_8d0ae84c7a3f3f55, __obf_bf6c24b903ce000b := strconv.ParseUint(string(__obf_669980db468f7544), 0, 64)
		if __obf_bf6c24b903ce000b != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
		__obf_d3c0db8b5b8d28e6.
			SetUint(__obf_8d0ae84c7a3f3f55)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_d633e91150cb7889.Type(), __obf_57ff3a2a2bd1a861)
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_73877d8afa80af89(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_afb9302665f844b5 := __obf_045d45fad1f5a0bc(__obf_d633e91150cb7889)

	switch {
	case __obf_afb9302665f844b5 == reflect.Bool:
		__obf_d3c0db8b5b8d28e6.
			SetBool(__obf_d633e91150cb7889.Bool())
	case __obf_afb9302665f844b5 == reflect.Int && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_d3c0db8b5b8d28e6.
			SetBool(__obf_d633e91150cb7889.Int() != 0)
	case __obf_afb9302665f844b5 == reflect.Uint && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_d3c0db8b5b8d28e6.
			SetBool(__obf_d633e91150cb7889.Uint() != 0)
	case __obf_afb9302665f844b5 == reflect.Float32 && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_d3c0db8b5b8d28e6.
			SetBool(__obf_d633e91150cb7889.Float() != 0)
	case __obf_afb9302665f844b5 == reflect.String && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_257c76af251ccba5, __obf_bf6c24b903ce000b := strconv.ParseBool(__obf_d633e91150cb7889.String())
		if __obf_bf6c24b903ce000b == nil {
			__obf_d3c0db8b5b8d28e6.
				SetBool(__obf_257c76af251ccba5)
		} else if __obf_d633e91150cb7889.String() == "" {
			__obf_d3c0db8b5b8d28e6.
				SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_d633e91150cb7889.Type(), __obf_57ff3a2a2bd1a861)
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_ec26544fb5023887(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_afb9302665f844b5 := __obf_045d45fad1f5a0bc(__obf_d633e91150cb7889)
	__obf_4ab72a4f9eed38d8 := __obf_d633e91150cb7889.Type()

	switch {
	case __obf_afb9302665f844b5 == reflect.Int:
		__obf_d3c0db8b5b8d28e6.
			SetFloat(float64(__obf_d633e91150cb7889.Int()))
	case __obf_afb9302665f844b5 == reflect.Uint:
		__obf_d3c0db8b5b8d28e6.
			SetFloat(float64(__obf_d633e91150cb7889.Uint()))
	case __obf_afb9302665f844b5 == reflect.Float32:
		__obf_d3c0db8b5b8d28e6.
			SetFloat(__obf_d633e91150cb7889.Float())
	case __obf_afb9302665f844b5 == reflect.Bool && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		if __obf_d633e91150cb7889.Bool() {
			__obf_d3c0db8b5b8d28e6.
				SetFloat(1)
		} else {
			__obf_d3c0db8b5b8d28e6.
				SetFloat(0)
		}
	case __obf_afb9302665f844b5 == reflect.String && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput:
		__obf_1264b3f46bd409c7 := __obf_d633e91150cb7889.String()
		if __obf_1264b3f46bd409c7 == "" {
			__obf_1264b3f46bd409c7 = "0"
		}
		__obf_2df189427d80bcc8, __obf_bf6c24b903ce000b := strconv.ParseFloat(__obf_1264b3f46bd409c7, __obf_d3c0db8b5b8d28e6.Type().Bits())
		if __obf_bf6c24b903ce000b == nil {
			__obf_d3c0db8b5b8d28e6.
				SetFloat(__obf_2df189427d80bcc8)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
	case __obf_4ab72a4f9eed38d8.PkgPath() == "encoding/json" && __obf_4ab72a4f9eed38d8.Name() == "Number":
		__obf_669980db468f7544 := __obf_57ff3a2a2bd1a861.(json.Number)
		__obf_8d0ae84c7a3f3f55, __obf_bf6c24b903ce000b := __obf_669980db468f7544.Float64()
		if __obf_bf6c24b903ce000b != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_040f4b76f7e163f8, __obf_bf6c24b903ce000b)
		}
		__obf_d3c0db8b5b8d28e6.
			SetFloat(__obf_8d0ae84c7a3f3f55)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_d633e91150cb7889.Type(), __obf_57ff3a2a2bd1a861)
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_5236cf9478516e08(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_6a780683fe63de12 := __obf_d3c0db8b5b8d28e6.Type()
	__obf_0ee690cce0341eda := __obf_6a780683fe63de12.Key()
	__obf_79dae659ee7c89fe := __obf_6a780683fe63de12.Elem()
	__obf_9a52c1c1374b6cf7 := // By default we overwrite keys in the current map
		__obf_d3c0db8b5b8d28e6

	// If the map is nil or we're purposely zeroing fields, make a new map
	if __obf_9a52c1c1374b6cf7.IsNil() || __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ZeroFields {
		__obf_1e4199fdf4ebc826 := // Make a new map to hold our result
			reflect.MapOf(__obf_0ee690cce0341eda, __obf_79dae659ee7c89fe)
		__obf_9a52c1c1374b6cf7 = reflect.MakeMap(__obf_1e4199fdf4ebc826)
	}
	__obf_d633e91150cb7889 := // Check input type and based on the input type jump to the proper func
		reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	switch __obf_d633e91150cb7889.Kind() {
	case reflect.Map:
		return __obf_11aa7e4430cc378c.__obf_2c7df846624c5c17(__obf_040f4b76f7e163f8, __obf_d633e91150cb7889, __obf_d3c0db8b5b8d28e6, __obf_9a52c1c1374b6cf7)

	case reflect.Struct:
		return __obf_11aa7e4430cc378c.__obf_c7930a021d3f4c4a(__obf_d633e91150cb7889, __obf_d3c0db8b5b8d28e6, __obf_9a52c1c1374b6cf7)

	case reflect.Array, reflect.Slice:
		if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput {
			return __obf_11aa7e4430cc378c.__obf_643c287409bacf48(__obf_040f4b76f7e163f8, __obf_d633e91150cb7889, __obf_d3c0db8b5b8d28e6, __obf_9a52c1c1374b6cf7)
		}

		fallthrough

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_040f4b76f7e163f8, __obf_d633e91150cb7889.Kind())
	}
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_643c287409bacf48(__obf_040f4b76f7e163f8 string, __obf_d633e91150cb7889 reflect.Value, __obf_d3c0db8b5b8d28e6 reflect.Value, __obf_9a52c1c1374b6cf7 reflect.Value) error {
	// Special case for BC reasons (covered by tests)
	if __obf_d633e91150cb7889.Len() == 0 {
		__obf_d3c0db8b5b8d28e6.
			Set(__obf_9a52c1c1374b6cf7)
		return nil
	}

	for __obf_8d0ae84c7a3f3f55 := 0; __obf_8d0ae84c7a3f3f55 < __obf_d633e91150cb7889.Len(); __obf_8d0ae84c7a3f3f55++ {
		__obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_040f4b76f7e163f8+
			"["+strconv.Itoa(__obf_8d0ae84c7a3f3f55)+"]", __obf_d633e91150cb7889.
			Index(__obf_8d0ae84c7a3f3f55).Interface(), __obf_d3c0db8b5b8d28e6)
		if __obf_bf6c24b903ce000b != nil {
			return __obf_bf6c24b903ce000b
		}
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_2c7df846624c5c17(__obf_040f4b76f7e163f8 string, __obf_d633e91150cb7889 reflect.Value, __obf_d3c0db8b5b8d28e6 reflect.Value, __obf_9a52c1c1374b6cf7 reflect.Value) error {
	__obf_6a780683fe63de12 := __obf_d3c0db8b5b8d28e6.Type()
	__obf_0ee690cce0341eda := __obf_6a780683fe63de12.Key()
	__obf_79dae659ee7c89fe := __obf_6a780683fe63de12.Elem()

	// Accumulate errors
	errors := make([]string, 0)

	// If the input data is empty, then we just match what the input data is.
	if __obf_d633e91150cb7889.Len() == 0 {
		if __obf_d633e91150cb7889.IsNil() {
			if !__obf_d3c0db8b5b8d28e6.IsNil() {
				__obf_d3c0db8b5b8d28e6.
					Set(__obf_d633e91150cb7889)
			}
		} else {
			__obf_d3c0db8b5b8d28e6.
				// Set to empty allocated value
				Set(__obf_9a52c1c1374b6cf7)
		}

		return nil
	}

	for _, __obf_055d50a77e6a33af := range __obf_d633e91150cb7889.MapKeys() {
		__obf_5e2e0c0f08a5e072 := __obf_040f4b76f7e163f8 + "[" + __obf_055d50a77e6a33af.String() + "]"
		__obf_2a411099810dd2e3 := // First decode the key into the proper type
			reflect.Indirect(reflect.New(__obf_0ee690cce0341eda))
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_5e2e0c0f08a5e072, __obf_055d50a77e6a33af.Interface(), __obf_2a411099810dd2e3); __obf_bf6c24b903ce000b != nil {
			errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
			continue
		}
		__obf_443dc76a6d8afc5e := // Next decode the data into the proper type
			__obf_d633e91150cb7889.MapIndex(__obf_055d50a77e6a33af).Interface()
		__obf_97e269bbd9345015 := reflect.Indirect(reflect.New(__obf_79dae659ee7c89fe))
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_5e2e0c0f08a5e072, __obf_443dc76a6d8afc5e, __obf_97e269bbd9345015); __obf_bf6c24b903ce000b != nil {
			errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
			continue
		}
		__obf_9a52c1c1374b6cf7.
			SetMapIndex(__obf_2a411099810dd2e3, __obf_97e269bbd9345015)
	}
	__obf_d3c0db8b5b8d28e6.

		// Set the built up map to the value
		Set(__obf_9a52c1c1374b6cf7)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_c7930a021d3f4c4a(__obf_d633e91150cb7889 reflect.Value, __obf_d3c0db8b5b8d28e6 reflect.Value, __obf_9a52c1c1374b6cf7 reflect.Value) error {
	__obf_6c1f770c08d9f414 := __obf_d633e91150cb7889.Type()
	for __obf_8d0ae84c7a3f3f55 := 0; __obf_8d0ae84c7a3f3f55 < __obf_6c1f770c08d9f414.NumField(); __obf_8d0ae84c7a3f3f55++ {
		__obf_2df189427d80bcc8 := // Get the StructField first since this is a cheap operation. If the
			// field is unexported, then ignore it.
			__obf_6c1f770c08d9f414.Field(__obf_8d0ae84c7a3f3f55)
		if __obf_2df189427d80bcc8.PkgPath != "" {
			continue
		}
		__obf_443dc76a6d8afc5e := // Next get the actual value of this field and verify it is assignable
			// to the map value.
			__obf_d633e91150cb7889.Field(__obf_8d0ae84c7a3f3f55)
		if !__obf_443dc76a6d8afc5e.Type().AssignableTo(__obf_9a52c1c1374b6cf7.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", __obf_443dc76a6d8afc5e.Type(), __obf_9a52c1c1374b6cf7.Type().Elem())
		}
		__obf_29495edcd7b51735 := __obf_2df189427d80bcc8.Tag.Get(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.TagName)
		__obf_b36065769e24cfc6 := __obf_2df189427d80bcc8.Name

		if __obf_29495edcd7b51735 == "" && __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.IgnoreUntaggedFields {
			continue
		}
		__obf_9df3b81ebb729f03 := // If Squash is set in the config, we squash the field down.
			__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Squash && __obf_443dc76a6d8afc5e.Kind() == reflect.Struct && __obf_2df189427d80bcc8.Anonymous
		__obf_443dc76a6d8afc5e = __obf_abdd48efd257d1df(__obf_443dc76a6d8afc5e, __obf_11aa7e4430cc378c.

			// Determine the name of the key in the map
			__obf_a5c9f567f7eba485.TagName)

		if __obf_e54ac02377f657e0 := strings.Index(__obf_29495edcd7b51735, ","); __obf_e54ac02377f657e0 != -1 {
			if __obf_29495edcd7b51735[:__obf_e54ac02377f657e0] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Contains(__obf_29495edcd7b51735[__obf_e54ac02377f657e0+1:], "omitempty") && __obf_97e877b318ab6c80(__obf_443dc76a6d8afc5e) {
				continue
			}
			__obf_9df3b81ebb729f03 = // If "squash" is specified in the tag, we squash the field down.
				__obf_9df3b81ebb729f03 || strings.Contains(__obf_29495edcd7b51735[__obf_e54ac02377f657e0+1:], "squash")
			if __obf_9df3b81ebb729f03 {
				// When squashing, the embedded type can be a pointer to a struct.
				if __obf_443dc76a6d8afc5e.Kind() == reflect.Ptr && __obf_443dc76a6d8afc5e.Elem().Kind() == reflect.Struct {
					__obf_443dc76a6d8afc5e = __obf_443dc76a6d8afc5e.Elem()
				}

				// The final type must be a struct
				if __obf_443dc76a6d8afc5e.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", __obf_443dc76a6d8afc5e.Type())
				}
			}
			if __obf_8e7b602b9d4dab41 := __obf_29495edcd7b51735[:__obf_e54ac02377f657e0]; __obf_8e7b602b9d4dab41 != "" {
				__obf_b36065769e24cfc6 = __obf_8e7b602b9d4dab41
			}
		} else if len(__obf_29495edcd7b51735) > 0 {
			if __obf_29495edcd7b51735 == "-" {
				continue
			}
			__obf_b36065769e24cfc6 = __obf_29495edcd7b51735
		}

		switch __obf_443dc76a6d8afc5e.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			__obf_c67ae5c88bdd05e7 := reflect.New(__obf_443dc76a6d8afc5e.Type())
			__obf_c67ae5c88bdd05e7.
				Elem().Set(__obf_443dc76a6d8afc5e)
			__obf_2c255e96f1a894cd := __obf_9a52c1c1374b6cf7.Type()
			__obf_ead4397e9fae5a16 := __obf_2c255e96f1a894cd.Key()
			__obf_059226ab81d5e1e6 := __obf_2c255e96f1a894cd.Elem()
			__obf_74d477617e20e2a7 := reflect.MapOf(__obf_ead4397e9fae5a16, __obf_059226ab81d5e1e6)
			__obf_5530156746d00216 := reflect.MakeMap(__obf_74d477617e20e2a7)
			__obf_8cf73cdaa9ec1448 := // Creating a pointer to a map so that other methods can completely
				// overwrite the map if need be (looking at you decodeMapFromMap). The
				// indirection allows the underlying map to be settable (CanSet() == true)
				// where as reflect.MakeMap returns an unsettable map.
				reflect.New(__obf_5530156746d00216.Type())
			reflect.Indirect(__obf_8cf73cdaa9ec1448).Set(__obf_5530156746d00216)
			__obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_b36065769e24cfc6, __obf_c67ae5c88bdd05e7.Interface(), reflect.Indirect(__obf_8cf73cdaa9ec1448))
			if __obf_bf6c24b903ce000b != nil {
				return __obf_bf6c24b903ce000b
			}
			__obf_5530156746d00216 = // the underlying map may have been completely overwritten so pull
				// it indirectly out of the enclosing value.
				reflect.Indirect(__obf_8cf73cdaa9ec1448)

			if __obf_9df3b81ebb729f03 {
				for _, __obf_055d50a77e6a33af := range __obf_5530156746d00216.MapKeys() {
					__obf_9a52c1c1374b6cf7.
						SetMapIndex(__obf_055d50a77e6a33af, __obf_5530156746d00216.MapIndex(__obf_055d50a77e6a33af))
				}
			} else {
				__obf_9a52c1c1374b6cf7.
					SetMapIndex(reflect.ValueOf(__obf_b36065769e24cfc6), __obf_5530156746d00216)
			}

		default:
			__obf_9a52c1c1374b6cf7.
				SetMapIndex(reflect.ValueOf(__obf_b36065769e24cfc6), __obf_443dc76a6d8afc5e)
		}
	}

	if __obf_d3c0db8b5b8d28e6.CanAddr() {
		__obf_d3c0db8b5b8d28e6.
			Set(__obf_9a52c1c1374b6cf7)
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_9ee9af437245a47f(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) (bool, error) {
	__obf_bbd4bc91c95c66a6 := // If the input data is nil, then we want to just set the output
		// pointer to be nil as well.
		__obf_57ff3a2a2bd1a861 == nil
	if !__obf_bbd4bc91c95c66a6 {
		switch __obf_443dc76a6d8afc5e := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861)); __obf_443dc76a6d8afc5e.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			__obf_bbd4bc91c95c66a6 = __obf_443dc76a6d8afc5e.IsNil()
		}
	}
	if __obf_bbd4bc91c95c66a6 {
		if !__obf_d3c0db8b5b8d28e6.IsNil() && __obf_d3c0db8b5b8d28e6.CanSet() {
			__obf_a7413b11c0b28818 := reflect.New(__obf_d3c0db8b5b8d28e6.Type()).Elem()
			__obf_d3c0db8b5b8d28e6.
				Set(__obf_a7413b11c0b28818)
		}

		return true, nil
	}
	__obf_6a780683fe63de12 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		__obf_d3c0db8b5b8d28e6.Type()
	__obf_79dae659ee7c89fe := __obf_6a780683fe63de12.Elem()
	if __obf_d3c0db8b5b8d28e6.CanSet() {
		__obf_e9ed51e15119e20e := __obf_d3c0db8b5b8d28e6
		if __obf_e9ed51e15119e20e.IsNil() || __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ZeroFields {
			__obf_e9ed51e15119e20e = reflect.New(__obf_79dae659ee7c89fe)
		}

		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_040f4b76f7e163f8, __obf_57ff3a2a2bd1a861, reflect.Indirect(__obf_e9ed51e15119e20e)); __obf_bf6c24b903ce000b != nil {
			return false, __obf_bf6c24b903ce000b
		}
		__obf_d3c0db8b5b8d28e6.
			Set(__obf_e9ed51e15119e20e)
	} else {
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_040f4b76f7e163f8, __obf_57ff3a2a2bd1a861, reflect.Indirect(__obf_d3c0db8b5b8d28e6)); __obf_bf6c24b903ce000b != nil {
			return false, __obf_bf6c24b903ce000b
		}
	}
	return false, nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_b9d4339190794a2e(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	if __obf_d3c0db8b5b8d28e6.Type() != __obf_d633e91150cb7889.Type() {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_040f4b76f7e163f8, __obf_d3c0db8b5b8d28e6.
				Type(), __obf_d633e91150cb7889.Type(), __obf_57ff3a2a2bd1a861)
	}
	__obf_d3c0db8b5b8d28e6.
		Set(__obf_d633e91150cb7889)
	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_4358644c7db5e853(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_87e738f37d5f511d := __obf_d633e91150cb7889.Kind()
	__obf_6a780683fe63de12 := __obf_d3c0db8b5b8d28e6.Type()
	__obf_79dae659ee7c89fe := __obf_6a780683fe63de12.Elem()
	__obf_7aec840716eac50b := reflect.SliceOf(__obf_79dae659ee7c89fe)

	// If we have a non array/slice type then we first attempt to convert.
	if __obf_87e738f37d5f511d != reflect.Array && __obf_87e738f37d5f511d != reflect.Slice {
		if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case __obf_87e738f37d5f511d == reflect.Slice, __obf_87e738f37d5f511d == reflect.Array:
				break

			// Empty maps turn into empty slices
			case __obf_87e738f37d5f511d == reflect.Map:
				if __obf_d633e91150cb7889.Len() == 0 {
					__obf_d3c0db8b5b8d28e6.
						Set(reflect.MakeSlice(__obf_7aec840716eac50b, 0, 0))
					return nil
				}
				// Create slice of maps of other sizes
				return __obf_11aa7e4430cc378c.__obf_4358644c7db5e853(__obf_040f4b76f7e163f8, []any{__obf_57ff3a2a2bd1a861}, __obf_d3c0db8b5b8d28e6)

			case __obf_87e738f37d5f511d == reflect.String && __obf_79dae659ee7c89fe.Kind() == reflect.Uint8:
				return __obf_11aa7e4430cc378c.__obf_4358644c7db5e853(__obf_040f4b76f7e163f8, []byte(__obf_d633e91150cb7889.String()), __obf_d3c0db8b5b8d28e6)

			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				return __obf_11aa7e4430cc378c.__obf_4358644c7db5e853(__obf_040f4b76f7e163f8, []any{__obf_57ff3a2a2bd1a861}, __obf_d3c0db8b5b8d28e6)
			}
		}

		return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", __obf_040f4b76f7e163f8, __obf_87e738f37d5f511d)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if __obf_87e738f37d5f511d != reflect.Array && __obf_d633e91150cb7889.IsNil() {
		return nil
	}
	__obf_b471617427f366b1 := __obf_d3c0db8b5b8d28e6
	if __obf_b471617427f366b1.IsNil() || __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ZeroFields {
		__obf_b471617427f366b1 = // Make a new slice to hold our result, same size as the original data.
			reflect.MakeSlice(__obf_7aec840716eac50b, __obf_d633e91150cb7889.Len(), __obf_d633e91150cb7889.Len())
	} else if __obf_b471617427f366b1.Len() > __obf_d633e91150cb7889.Len() {
		__obf_b471617427f366b1 = __obf_b471617427f366b1.Slice(0, __obf_d633e91150cb7889.Len())
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_8d0ae84c7a3f3f55 := 0; __obf_8d0ae84c7a3f3f55 < __obf_d633e91150cb7889.Len(); __obf_8d0ae84c7a3f3f55++ {
		__obf_9bba010055e7e9f2 := __obf_d633e91150cb7889.Index(__obf_8d0ae84c7a3f3f55).Interface()
		for __obf_b471617427f366b1.Len() <= __obf_8d0ae84c7a3f3f55 {
			__obf_b471617427f366b1 = reflect.Append(__obf_b471617427f366b1, reflect.Zero(__obf_79dae659ee7c89fe))
		}
		__obf_d0ed975ddb99a844 := __obf_b471617427f366b1.Index(__obf_8d0ae84c7a3f3f55)
		__obf_5e2e0c0f08a5e072 := __obf_040f4b76f7e163f8 + "[" + strconv.Itoa(__obf_8d0ae84c7a3f3f55) + "]"
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_5e2e0c0f08a5e072, __obf_9bba010055e7e9f2, __obf_d0ed975ddb99a844); __obf_bf6c24b903ce000b != nil {
			errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
		}
	}
	__obf_d3c0db8b5b8d28e6.

		// Finally, set the value to the slice we built up
		Set(__obf_b471617427f366b1)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_28cbf385ef2a4ccf(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))
	__obf_87e738f37d5f511d := __obf_d633e91150cb7889.Kind()
	__obf_6a780683fe63de12 := __obf_d3c0db8b5b8d28e6.Type()
	__obf_79dae659ee7c89fe := __obf_6a780683fe63de12.Elem()
	__obf_838f266b7b54bacc := reflect.ArrayOf(__obf_6a780683fe63de12.Len(), __obf_79dae659ee7c89fe)
	__obf_415f2a132618b66a := __obf_d3c0db8b5b8d28e6

	if __obf_415f2a132618b66a.Interface() == reflect.Zero(__obf_415f2a132618b66a.Type()).Interface() || __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ZeroFields {
		// Check input type
		if __obf_87e738f37d5f511d != reflect.Array && __obf_87e738f37d5f511d != reflect.Slice {
			if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.WeaklyTypedInput {
				switch __obf_87e738f37d5f511d {
				// Empty maps turn into empty arrays
				case reflect.Map:
					if __obf_d633e91150cb7889.Len() == 0 {
						__obf_d3c0db8b5b8d28e6.
							Set(reflect.Zero(__obf_838f266b7b54bacc))
						return nil
					}

				// All other types we try to convert to the array type
				// and "lift" it into it. i.e. a string becomes a string array.
				default:
					// Just re-try this function with data as a slice.
					return __obf_11aa7e4430cc378c.__obf_28cbf385ef2a4ccf(__obf_040f4b76f7e163f8, []any{__obf_57ff3a2a2bd1a861}, __obf_d3c0db8b5b8d28e6)
				}
			}

			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", __obf_040f4b76f7e163f8, __obf_87e738f37d5f511d)

		}
		if __obf_d633e91150cb7889.Len() > __obf_838f266b7b54bacc.Len() {
			return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", __obf_040f4b76f7e163f8, __obf_838f266b7b54bacc.Len(), __obf_d633e91150cb7889.Len())

		}
		__obf_415f2a132618b66a = // Make a new array to hold our result, same size as the original data.
			reflect.New(__obf_838f266b7b54bacc).Elem()
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_8d0ae84c7a3f3f55 := 0; __obf_8d0ae84c7a3f3f55 < __obf_d633e91150cb7889.Len(); __obf_8d0ae84c7a3f3f55++ {
		__obf_9bba010055e7e9f2 := __obf_d633e91150cb7889.Index(__obf_8d0ae84c7a3f3f55).Interface()
		__obf_d0ed975ddb99a844 := __obf_415f2a132618b66a.Index(__obf_8d0ae84c7a3f3f55)
		__obf_5e2e0c0f08a5e072 := __obf_040f4b76f7e163f8 + "[" + strconv.Itoa(__obf_8d0ae84c7a3f3f55) + "]"
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_5e2e0c0f08a5e072, __obf_9bba010055e7e9f2, __obf_d0ed975ddb99a844); __obf_bf6c24b903ce000b != nil {
			errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
		}
	}
	__obf_d3c0db8b5b8d28e6.

		// Finally, set the value to the array we built up
		Set(__obf_415f2a132618b66a)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_3a979d7c2b59d928(__obf_040f4b76f7e163f8 string, __obf_57ff3a2a2bd1a861 any, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_d633e91150cb7889 := reflect.Indirect(reflect.ValueOf(__obf_57ff3a2a2bd1a861))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if __obf_d633e91150cb7889.Type() == __obf_d3c0db8b5b8d28e6.Type() {
		__obf_d3c0db8b5b8d28e6.
			Set(__obf_d633e91150cb7889)
		return nil
	}
	__obf_87e738f37d5f511d := __obf_d633e91150cb7889.Kind()
	switch __obf_87e738f37d5f511d {
	case reflect.Map:
		return __obf_11aa7e4430cc378c.__obf_7e04c36a8f95a514(__obf_040f4b76f7e163f8, __obf_d633e91150cb7889,

			// Not the most efficient way to do this but we can optimize later if
			// we want to. To convert from struct to struct we go to map first
			// as an intermediary.
			__obf_d3c0db8b5b8d28e6)

	case reflect.Struct:
		__obf_1e4199fdf4ebc826 := // Make a new map to hold our result
			reflect.TypeOf((map[string]any)(nil))
		__obf_7077a8e1795c77b2 := reflect.MakeMap(__obf_1e4199fdf4ebc826)
		__obf_8cf73cdaa9ec1448 := // Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			reflect.New(__obf_7077a8e1795c77b2.Type())

		reflect.Indirect(__obf_8cf73cdaa9ec1448).Set(__obf_7077a8e1795c77b2)
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_c7930a021d3f4c4a(__obf_d633e91150cb7889, reflect.Indirect(__obf_8cf73cdaa9ec1448), __obf_7077a8e1795c77b2); __obf_bf6c24b903ce000b != nil {
			return __obf_bf6c24b903ce000b
		}
		__obf_6ad3d371cac0efc2 := __obf_11aa7e4430cc378c.__obf_7e04c36a8f95a514(__obf_040f4b76f7e163f8, reflect.Indirect(__obf_8cf73cdaa9ec1448), __obf_d3c0db8b5b8d28e6)
		return __obf_6ad3d371cac0efc2

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_040f4b76f7e163f8, __obf_d633e91150cb7889.Kind())
	}
}

func (__obf_11aa7e4430cc378c *Decoder) __obf_7e04c36a8f95a514(__obf_040f4b76f7e163f8 string, __obf_d633e91150cb7889, __obf_d3c0db8b5b8d28e6 reflect.Value) error {
	__obf_e58ea9c91a850d24 := __obf_d633e91150cb7889.Type()
	if __obf_d1db9a7181ddead7 := __obf_e58ea9c91a850d24.Key().Kind(); __obf_d1db9a7181ddead7 != reflect.String && __obf_d1db9a7181ddead7 != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys", __obf_040f4b76f7e163f8, __obf_e58ea9c91a850d24.
				Key().Kind())
	}
	__obf_515769afd06bc114 := make(map[reflect.Value]struct{})
	__obf_b5f670b36a54818d := make(map[any]struct{})
	for _, __obf_297cd7ba9475c410 := range __obf_d633e91150cb7889.MapKeys() {
		__obf_515769afd06bc114[__obf_297cd7ba9475c410] = struct{}{}
		__obf_b5f670b36a54818d[__obf_297cd7ba9475c410.Interface()] = struct{}{}
	}
	__obf_570bb81189aeee87 := make(map[any]struct{})
	errors := make([]string, 0)
	__obf_250b9b73eefeb572 := // This slice will keep track of all the structs we'll be decoding.
		// There can be more than one struct if there are embedded structs
		// that are squashed.
		make([]reflect.Value, 1, 5)
	__obf_250b9b73eefeb572[0] = __obf_d3c0db8b5b8d28e6

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type __obf_147cb46c73e450de struct {
		__obf_147cb46c73e450de reflect.StructField
		__obf_d3c0db8b5b8d28e6 reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var __obf_51d4c650654e7108 *__obf_147cb46c73e450de
	__obf_5ea17afc83d89050 := []__obf_147cb46c73e450de{}
	for len(__obf_250b9b73eefeb572) > 0 {
		__obf_59e467228ec43379 := __obf_250b9b73eefeb572[0]
		__obf_250b9b73eefeb572 = __obf_250b9b73eefeb572[1:]
		__obf_87f97d32fcc1fd75 := __obf_59e467228ec43379.Type()

		for __obf_8d0ae84c7a3f3f55 := 0; __obf_8d0ae84c7a3f3f55 < __obf_87f97d32fcc1fd75.NumField(); __obf_8d0ae84c7a3f3f55++ {
			__obf_9b9f08835f5c11b3 := __obf_87f97d32fcc1fd75.Field(__obf_8d0ae84c7a3f3f55)
			__obf_5da22662169ba887 := __obf_59e467228ec43379.Field(__obf_8d0ae84c7a3f3f55)
			if __obf_5da22662169ba887.Kind() == reflect.Pointer && __obf_5da22662169ba887.Elem().Kind() == reflect.Struct {
				__obf_5da22662169ba887 = // Handle embedded struct pointers as embedded structs.
					__obf_5da22662169ba887.Elem()
			}
			__obf_9df3b81ebb729f03 := // If "squash" is specified in the tag, we squash the field down.
				__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Squash && __obf_5da22662169ba887.Kind() == reflect.Struct && __obf_9b9f08835f5c11b3.Anonymous
			__obf_4ce6272fff328ed9 := false
			__obf_813d529dde70dd29 := // We always parse the tags cause we're looking for other tags too
				strings.Split(__obf_9b9f08835f5c11b3.Tag.Get(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.TagName), ",")
			for _, __obf_1ca9f40ee9bb306e := range __obf_813d529dde70dd29[1:] {
				if __obf_1ca9f40ee9bb306e == "squash" {
					__obf_9df3b81ebb729f03 = true
					break
				}

				if __obf_1ca9f40ee9bb306e == "remain" {
					__obf_4ce6272fff328ed9 = true
					break
				}
			}

			if __obf_9df3b81ebb729f03 {
				if __obf_5da22662169ba887.Kind() != reflect.Struct {
					errors = __obf_8c0a23210109d8dd(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", __obf_9b9f08835f5c11b3.Name, __obf_5da22662169ba887.Kind()))
				} else {
					__obf_250b9b73eefeb572 = append(__obf_250b9b73eefeb572, __obf_5da22662169ba887)
				}
				continue
			}

			// Build our field
			if __obf_4ce6272fff328ed9 {
				__obf_51d4c650654e7108 = &__obf_147cb46c73e450de{__obf_9b9f08835f5c11b3, __obf_5da22662169ba887}
			} else {
				__obf_5ea17afc83d89050 = // Normal struct field, store it away
					append(__obf_5ea17afc83d89050, __obf_147cb46c73e450de{__obf_9b9f08835f5c11b3, __obf_5da22662169ba887})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, __obf_2df189427d80bcc8 := range __obf_5ea17afc83d89050 {
		__obf_147cb46c73e450de, __obf_7eb26dac7077aabe := __obf_2df189427d80bcc8.__obf_147cb46c73e450de, __obf_2df189427d80bcc8.__obf_d3c0db8b5b8d28e6
		__obf_5e2e0c0f08a5e072 := __obf_147cb46c73e450de.Name
		__obf_29495edcd7b51735 := __obf_147cb46c73e450de.Tag.Get(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.TagName)
		__obf_29495edcd7b51735 = strings.SplitN(__obf_29495edcd7b51735, ",", 2)[0]
		if __obf_29495edcd7b51735 != "" {
			__obf_5e2e0c0f08a5e072 = __obf_29495edcd7b51735
		}
		__obf_327e18e0e12403b3 := reflect.ValueOf(__obf_5e2e0c0f08a5e072)
		__obf_751e26b88fd5d7cd := __obf_d633e91150cb7889.MapIndex(__obf_327e18e0e12403b3)
		if !__obf_751e26b88fd5d7cd.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for __obf_297cd7ba9475c410 := range __obf_515769afd06bc114 {
				__obf_cd61bcc82dfb08d6, __obf_5da7816f07b20ebf := __obf_297cd7ba9475c410.Interface().(string)
				if !__obf_5da7816f07b20ebf {
					// Not a string key
					continue
				}

				if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.MatchName(__obf_cd61bcc82dfb08d6, __obf_5e2e0c0f08a5e072) {
					__obf_327e18e0e12403b3 = __obf_297cd7ba9475c410
					__obf_751e26b88fd5d7cd = __obf_d633e91150cb7889.MapIndex(__obf_297cd7ba9475c410)
					break
				}
			}

			if !__obf_751e26b88fd5d7cd.IsValid() {
				__obf_570bb81189aeee87[ // There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				__obf_5e2e0c0f08a5e072] = struct{}{}
				continue
			}
		}

		if !__obf_7eb26dac7077aabe.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !__obf_7eb26dac7077aabe.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(__obf_b5f670b36a54818d, __obf_327e18e0e12403b3.Interface())

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if __obf_040f4b76f7e163f8 != "" {
			__obf_5e2e0c0f08a5e072 = __obf_040f4b76f7e163f8 + "." + __obf_5e2e0c0f08a5e072
		}

		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_cc469f327ca4db82(__obf_5e2e0c0f08a5e072, __obf_751e26b88fd5d7cd.Interface(), __obf_7eb26dac7077aabe); __obf_bf6c24b903ce000b != nil {
			errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if __obf_51d4c650654e7108 != nil && len(__obf_b5f670b36a54818d) > 0 {
		__obf_4ce6272fff328ed9 := // Build a map of only the unused values
			map[any]any{}
		for __obf_71ec40664f41fdce := range __obf_b5f670b36a54818d {
			__obf_4ce6272fff328ed9[__obf_71ec40664f41fdce] = __obf_d633e91150cb7889.MapIndex(reflect.ValueOf(__obf_71ec40664f41fdce)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if __obf_bf6c24b903ce000b := __obf_11aa7e4430cc378c.__obf_5236cf9478516e08(__obf_040f4b76f7e163f8, __obf_4ce6272fff328ed9, __obf_51d4c650654e7108.__obf_d3c0db8b5b8d28e6); __obf_bf6c24b903ce000b != nil {
			errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
		}
		__obf_b5f670b36a54818d = // Set the map to nil so we have none so that the next check will
			// not error (ErrorUnused)
			nil
	}

	if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ErrorUnused && len(__obf_b5f670b36a54818d) > 0 {
		__obf_2a507a67e7a75327 := make([]string, 0, len(__obf_b5f670b36a54818d))
		for __obf_3ae65e0e5b9bad1b := range __obf_b5f670b36a54818d {
			__obf_2a507a67e7a75327 = append(__obf_2a507a67e7a75327, __obf_3ae65e0e5b9bad1b.(string))
		}
		sort.Strings(__obf_2a507a67e7a75327)
		__obf_bf6c24b903ce000b := fmt.Errorf("'%s' has invalid keys: %s", __obf_040f4b76f7e163f8, strings.Join(__obf_2a507a67e7a75327, ", "))
		errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
	}

	if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.ErrorUnset && len(__obf_570bb81189aeee87) > 0 {
		__obf_2a507a67e7a75327 := make([]string, 0, len(__obf_570bb81189aeee87))
		for __obf_3ae65e0e5b9bad1b := range __obf_570bb81189aeee87 {
			__obf_2a507a67e7a75327 = append(__obf_2a507a67e7a75327, __obf_3ae65e0e5b9bad1b.(string))
		}
		sort.Strings(__obf_2a507a67e7a75327)
		__obf_bf6c24b903ce000b := fmt.Errorf("'%s' has unset fields: %s", __obf_040f4b76f7e163f8, strings.Join(__obf_2a507a67e7a75327, ", "))
		errors = __obf_8c0a23210109d8dd(errors, __obf_bf6c24b903ce000b)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	if __obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata != nil {
		for __obf_3ae65e0e5b9bad1b := range __obf_b5f670b36a54818d {
			__obf_71ec40664f41fdce := __obf_3ae65e0e5b9bad1b.(string)
			if __obf_040f4b76f7e163f8 != "" {
				__obf_71ec40664f41fdce = __obf_040f4b76f7e163f8 + "." + __obf_71ec40664f41fdce
			}
			__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.
				Metadata.Unused = append(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata.Unused, __obf_71ec40664f41fdce)
		}
		for __obf_3ae65e0e5b9bad1b := range __obf_570bb81189aeee87 {
			__obf_71ec40664f41fdce := __obf_3ae65e0e5b9bad1b.(string)
			if __obf_040f4b76f7e163f8 != "" {
				__obf_71ec40664f41fdce = __obf_040f4b76f7e163f8 + "." + __obf_71ec40664f41fdce
			}
			__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.
				Metadata.Unset = append(__obf_11aa7e4430cc378c.__obf_a5c9f567f7eba485.Metadata.Unset, __obf_71ec40664f41fdce)
		}
	}

	return nil
}

func __obf_97e877b318ab6c80(__obf_443dc76a6d8afc5e reflect.Value) bool {
	switch __obf_045d45fad1f5a0bc(__obf_443dc76a6d8afc5e) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_443dc76a6d8afc5e.Len() == 0
	case reflect.Bool:
		return !__obf_443dc76a6d8afc5e.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_443dc76a6d8afc5e.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_443dc76a6d8afc5e.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_443dc76a6d8afc5e.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return __obf_443dc76a6d8afc5e.IsNil()
	}
	return false
}

func __obf_045d45fad1f5a0bc(__obf_d3c0db8b5b8d28e6 reflect.Value) reflect.Kind {
	__obf_d1db9a7181ddead7 := __obf_d3c0db8b5b8d28e6.Kind()

	switch {
	case __obf_d1db9a7181ddead7 >= reflect.Int && __obf_d1db9a7181ddead7 <= reflect.Int64:
		return reflect.Int
	case __obf_d1db9a7181ddead7 >= reflect.Uint && __obf_d1db9a7181ddead7 <= reflect.Uint64:
		return reflect.Uint
	case __obf_d1db9a7181ddead7 >= reflect.Float32 && __obf_d1db9a7181ddead7 <= reflect.Float64:
		return reflect.Float32
	default:
		return __obf_d1db9a7181ddead7
	}
}

func __obf_97d24ce452f8b875(__obf_6c1f770c08d9f414 reflect.Type, __obf_a7f48e186430592c bool, __obf_d3d5668c0053bdc7 string) bool {
	for __obf_8d0ae84c7a3f3f55 := 0; __obf_8d0ae84c7a3f3f55 < __obf_6c1f770c08d9f414.NumField(); __obf_8d0ae84c7a3f3f55++ {
		__obf_2df189427d80bcc8 := __obf_6c1f770c08d9f414.Field(__obf_8d0ae84c7a3f3f55)
		if __obf_2df189427d80bcc8.PkgPath == "" && !__obf_a7f48e186430592c { // check for unexported fields
			return true
		}
		if __obf_a7f48e186430592c && __obf_2df189427d80bcc8.Tag.Get(__obf_d3d5668c0053bdc7) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func __obf_abdd48efd257d1df(__obf_443dc76a6d8afc5e reflect.Value, __obf_d3d5668c0053bdc7 string) reflect.Value {
	if __obf_443dc76a6d8afc5e.Kind() != reflect.Ptr || __obf_443dc76a6d8afc5e.Elem().Kind() != reflect.Struct {
		return __obf_443dc76a6d8afc5e
	}
	__obf_9b04a152763b460a := __obf_443dc76a6d8afc5e.Elem()
	__obf_156811b2f8d54473 := __obf_9b04a152763b460a.Type()
	if __obf_97d24ce452f8b875(__obf_156811b2f8d54473, true, __obf_d3d5668c0053bdc7) {
		return __obf_9b04a152763b460a
	}
	return __obf_443dc76a6d8afc5e
}
