package __obf_ec2f65f16fa88470

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
type DecodeHookFuncValue func(__obf_3a8aa8877b053f0b reflect.Value, __obf_6dd945bb01fb3b32 reflect.Value) (any, error)

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
	MatchName func(__obf_667ff40abd218181, __obf_1cc7b5d4481b301c string) bool
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	__obf_da116ef5edc24797 *DecoderConfig
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
func Decode(__obf_d8521da8ec1eda3b any, __obf_79635ef79c313f29 any) error {
	__obf_da116ef5edc24797 := &DecoderConfig{
		Metadata: nil,
		Result:   __obf_79635ef79c313f29,
	}
	__obf_d18b2e6d691ba699, __obf_6d8ef670b0943b3c := NewDecoder(__obf_da116ef5edc24797)
	if __obf_6d8ef670b0943b3c != nil {
		return __obf_6d8ef670b0943b3c
	}

	return __obf_d18b2e6d691ba699.Decode(__obf_d8521da8ec1eda3b)
}

// WeakDecode is the same as Decode but is shorthand to enable
// WeaklyTypedInput. See DecoderConfig for more info.
func WeakDecode(__obf_d8521da8ec1eda3b, __obf_79635ef79c313f29 any) error {
	__obf_da116ef5edc24797 := &DecoderConfig{
		Metadata:         nil,
		Result:           __obf_79635ef79c313f29,
		WeaklyTypedInput: true,
	}
	__obf_d18b2e6d691ba699, __obf_6d8ef670b0943b3c := NewDecoder(__obf_da116ef5edc24797)
	if __obf_6d8ef670b0943b3c != nil {
		return __obf_6d8ef670b0943b3c
	}

	return __obf_d18b2e6d691ba699.Decode(__obf_d8521da8ec1eda3b)
}

// DecodeMetadata is the same as Decode, but is shorthand to
// enable metadata collection. See DecoderConfig for more info.
func DecodeMetadata(__obf_d8521da8ec1eda3b any, __obf_79635ef79c313f29 any, __obf_075bd489d9865302 *Metadata) error {
	__obf_da116ef5edc24797 := &DecoderConfig{
		Metadata: __obf_075bd489d9865302,
		Result:   __obf_79635ef79c313f29,
	}
	__obf_d18b2e6d691ba699, __obf_6d8ef670b0943b3c := NewDecoder(__obf_da116ef5edc24797)
	if __obf_6d8ef670b0943b3c != nil {
		return __obf_6d8ef670b0943b3c
	}

	return __obf_d18b2e6d691ba699.Decode(__obf_d8521da8ec1eda3b)
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
// enable both WeaklyTypedInput and metadata collection. See
// DecoderConfig for more info.
func WeakDecodeMetadata(__obf_d8521da8ec1eda3b any, __obf_79635ef79c313f29 any, __obf_075bd489d9865302 *Metadata) error {
	__obf_da116ef5edc24797 := &DecoderConfig{
		Metadata:         __obf_075bd489d9865302,
		Result:           __obf_79635ef79c313f29,
		WeaklyTypedInput: true,
	}
	__obf_d18b2e6d691ba699, __obf_6d8ef670b0943b3c := NewDecoder(__obf_da116ef5edc24797)
	if __obf_6d8ef670b0943b3c != nil {
		return __obf_6d8ef670b0943b3c
	}

	return __obf_d18b2e6d691ba699.Decode(__obf_d8521da8ec1eda3b)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(__obf_da116ef5edc24797 *DecoderConfig) (*Decoder, error) {
	__obf_cd65ed68b120e383 := reflect.ValueOf(__obf_da116ef5edc24797.Result)
	if __obf_cd65ed68b120e383.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}
	__obf_cd65ed68b120e383 = __obf_cd65ed68b120e383.Elem()
	if !__obf_cd65ed68b120e383.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if __obf_da116ef5edc24797.Metadata != nil {
		if __obf_da116ef5edc24797.Metadata.Keys == nil {
			__obf_da116ef5edc24797.
				Metadata.Keys = make([]string, 0)
		}

		if __obf_da116ef5edc24797.Metadata.Unused == nil {
			__obf_da116ef5edc24797.
				Metadata.Unused = make([]string, 0)
		}

		if __obf_da116ef5edc24797.Metadata.Unset == nil {
			__obf_da116ef5edc24797.
				Metadata.Unset = make([]string, 0)
		}
	}

	if __obf_da116ef5edc24797.TagName == "" {
		__obf_da116ef5edc24797.
			TagName = "mapstructure"
	}

	if __obf_da116ef5edc24797.MatchName == nil {
		__obf_da116ef5edc24797.
			MatchName = strings.EqualFold
	}
	__obf_bb9d1909c7b1c125 := &Decoder{__obf_da116ef5edc24797: __obf_da116ef5edc24797}

	return __obf_bb9d1909c7b1c125, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (__obf_15070727b144cc26 *Decoder) Decode(__obf_d8521da8ec1eda3b any) error {
	return __obf_15070727b144cc26.__obf_54e09d10fb5e308c("", __obf_d8521da8ec1eda3b, reflect.ValueOf(__obf_15070727b144cc26.__obf_da116ef5edc24797.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (__obf_15070727b144cc26 *Decoder) __obf_54e09d10fb5e308c(__obf_97d979f40f648bde string, __obf_d8521da8ec1eda3b any, __obf_e679a50f9ef83ba5 reflect.Value) error {
	var __obf_45e310d47988c135 reflect.Value
	if __obf_d8521da8ec1eda3b != nil {
		__obf_45e310d47988c135 = reflect.ValueOf(__obf_d8521da8ec1eda3b)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if __obf_45e310d47988c135.Kind() == reflect.Ptr && __obf_45e310d47988c135.IsNil() {
			__obf_d8521da8ec1eda3b = nil
		}
	}

	if __obf_d8521da8ec1eda3b == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if __obf_15070727b144cc26.__obf_da116ef5edc24797.ZeroFields {
			__obf_e679a50f9ef83ba5.
				Set(reflect.Zero(__obf_e679a50f9ef83ba5.Type()))

			if __obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata != nil && __obf_97d979f40f648bde != "" {
				__obf_15070727b144cc26.__obf_da116ef5edc24797.
					Metadata.Keys = append(__obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata.Keys, __obf_97d979f40f648bde)
			}
		}
		return nil
	}

	if !__obf_45e310d47988c135.IsValid() {
		__obf_e679a50f9ef83ba5.
			// If the input value is invalid, then we just set the value
			// to be the zero value.
			Set(reflect.Zero(__obf_e679a50f9ef83ba5.Type()))
		if __obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata != nil && __obf_97d979f40f648bde != "" {
			__obf_15070727b144cc26.__obf_da116ef5edc24797.
				Metadata.Keys = append(__obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata.Keys, __obf_97d979f40f648bde)
		}
		return nil
	}

	if __obf_15070727b144cc26.__obf_da116ef5edc24797.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the input.
		var __obf_6d8ef670b0943b3c error
		__obf_d8521da8ec1eda3b, __obf_6d8ef670b0943b3c = DecodeHookExec(__obf_15070727b144cc26.__obf_da116ef5edc24797.DecodeHook, __obf_45e310d47988c135, __obf_e679a50f9ef83ba5)
		if __obf_6d8ef670b0943b3c != nil {
			return fmt.Errorf("error decoding '%s': %w", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
	}

	var __obf_6d8ef670b0943b3c error
	__obf_17bcee0eac4391c4 := __obf_ed31af980b14987e(__obf_e679a50f9ef83ba5)
	__obf_7c38737190903e34 := true
	switch __obf_17bcee0eac4391c4 {
	case reflect.Bool:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_bd1426afee26ea4f(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Interface:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_8e5d3c59d3711c73(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.String:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_09e78b174dc93b15(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Int:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_3bb51fd5d1a2bffd(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Uint:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_7c63467fdcb0d6ec(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Float32:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_02660e9348752825(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Struct:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_577e93fa4008dcd6(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Map:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_d20c19b1598726f2(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Ptr:
		__obf_7c38737190903e34, __obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_6b92db91c5345245(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Slice:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_a1d958c6ee649900(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Array:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_912e64085a8c4c83(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b, __obf_e679a50f9ef83ba5)
	case reflect.Func:
		__obf_6d8ef670b0943b3c = __obf_15070727b144cc26.__obf_d4a936768e419fa1(__obf_97d979f40f648bde, __obf_d8521da8ec1eda3b,

			// If we reached this point then we weren't able to decode it
			__obf_e679a50f9ef83ba5)
	default:

		return fmt.Errorf("%s: unsupported type: %s", __obf_97d979f40f648bde, __obf_17bcee0eac4391c4)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metainput.
	if __obf_7c38737190903e34 && __obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata != nil && __obf_97d979f40f648bde != "" {
		__obf_15070727b144cc26.__obf_da116ef5edc24797.
			Metadata.Keys = append(__obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata.Keys, __obf_97d979f40f648bde)
	}

	return __obf_6d8ef670b0943b3c
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (__obf_15070727b144cc26 *Decoder) __obf_8e5d3c59d3711c73(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	if __obf_cd65ed68b120e383.IsValid() && __obf_cd65ed68b120e383.Elem().IsValid() {
		__obf_a985093c55e986f9 := __obf_cd65ed68b120e383.Elem()
		__obf_2199e12b729261a2 := // If we can't address this element, then its not writable. Instead,
			// we make a copy of the value (which is a pointer and therefore
			// writable), decode into that, and replace the whole value.
			false
		if !__obf_a985093c55e986f9.CanAddr() {
			__obf_2199e12b729261a2 = true

			// Make *T
			copy := reflect.New(__obf_a985093c55e986f9.Type())

			// *T = elem
			copy.Elem().Set(__obf_a985093c55e986f9)
			__obf_a985093c55e986f9 = // Set elem so we decode into it
				copy
		}

		// Decode. If we have an error then return. We also return right
		// away if we're not a copy because that means we decoded directly.
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_97d979f40f648bde, __obf_e428e8b18837f120, __obf_a985093c55e986f9); __obf_6d8ef670b0943b3c != nil || !__obf_2199e12b729261a2 {
			return __obf_6d8ef670b0943b3c
		}
		__obf_cd65ed68b120e383.

			// If we're a copy, we need to set te final result
			Set(__obf_a985093c55e986f9.Elem())
		return nil
	}
	__obf_a30e8e069ebf8d56 := reflect.ValueOf(__obf_e428e8b18837f120)

	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if __obf_a30e8e069ebf8d56.Kind() == reflect.Ptr && __obf_a30e8e069ebf8d56.Type().Elem() == __obf_cd65ed68b120e383.Type() {
		__obf_a30e8e069ebf8d56 = reflect.Indirect(__obf_a30e8e069ebf8d56)
	}

	if !__obf_a30e8e069ebf8d56.IsValid() {
		__obf_a30e8e069ebf8d56 = reflect.Zero(__obf_cd65ed68b120e383.Type())
	}
	__obf_8d4acaad13f3fa10 := __obf_a30e8e069ebf8d56.Type()
	if !__obf_8d4acaad13f3fa10.AssignableTo(__obf_cd65ed68b120e383.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_8d4acaad13f3fa10)
	}
	__obf_cd65ed68b120e383.
		Set(__obf_a30e8e069ebf8d56)
	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_09e78b174dc93b15(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8c8f659ba1d013f1 := __obf_ed31af980b14987e(__obf_a30e8e069ebf8d56)
	__obf_c90541f4a19cc78a := true
	switch {
	case __obf_8c8f659ba1d013f1 == reflect.String:
		__obf_cd65ed68b120e383.
			SetString(__obf_a30e8e069ebf8d56.String())
	case __obf_8c8f659ba1d013f1 == reflect.Bool && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		if __obf_a30e8e069ebf8d56.Bool() {
			__obf_cd65ed68b120e383.
				SetString("1")
		} else {
			__obf_cd65ed68b120e383.
				SetString("0")
		}
	case __obf_8c8f659ba1d013f1 == reflect.Int && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_cd65ed68b120e383.
			SetString(strconv.FormatInt(__obf_a30e8e069ebf8d56.Int(), 10))
	case __obf_8c8f659ba1d013f1 == reflect.Uint && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_cd65ed68b120e383.
			SetString(strconv.FormatUint(__obf_a30e8e069ebf8d56.Uint(), 10))
	case __obf_8c8f659ba1d013f1 == reflect.Float32 && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_cd65ed68b120e383.
			SetString(strconv.FormatFloat(__obf_a30e8e069ebf8d56.Float(), 'f', -1, 64))
	case __obf_8c8f659ba1d013f1 == reflect.Slice && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput, __obf_8c8f659ba1d013f1 ==
		reflect.Array && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_9dd014252313d6cd := __obf_a30e8e069ebf8d56.Type()
		__obf_0673273cce64670d := __obf_9dd014252313d6cd.Elem().Kind()
		switch __obf_0673273cce64670d {
		case reflect.Uint8:
			var __obf_90be5f4e9a63d34c []uint8
			if __obf_8c8f659ba1d013f1 == reflect.Array {
				__obf_90be5f4e9a63d34c = make([]uint8, __obf_a30e8e069ebf8d56.Len())
				for __obf_6cbf6487a01daf31 := range __obf_90be5f4e9a63d34c {
					__obf_90be5f4e9a63d34c[__obf_6cbf6487a01daf31] = __obf_a30e8e069ebf8d56.Index(__obf_6cbf6487a01daf31).Interface().(uint8)
				}
			} else {
				__obf_90be5f4e9a63d34c = __obf_a30e8e069ebf8d56.Interface().([]uint8)
			}
			__obf_cd65ed68b120e383.
				SetString(string(__obf_90be5f4e9a63d34c))
		default:
			__obf_c90541f4a19cc78a = false
		}
	default:
		__obf_c90541f4a19cc78a = false
	}

	if !__obf_c90541f4a19cc78a {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_a30e8e069ebf8d56.Type(), __obf_e428e8b18837f120)
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_3bb51fd5d1a2bffd(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8c8f659ba1d013f1 := __obf_ed31af980b14987e(__obf_a30e8e069ebf8d56)
	__obf_9dd014252313d6cd := __obf_a30e8e069ebf8d56.Type()

	switch {
	case __obf_8c8f659ba1d013f1 == reflect.Int:
		__obf_cd65ed68b120e383.
			SetInt(__obf_a30e8e069ebf8d56.Int())
	case __obf_8c8f659ba1d013f1 == reflect.Uint:
		__obf_cd65ed68b120e383.
			SetInt(int64(__obf_a30e8e069ebf8d56.Uint()))
	case __obf_8c8f659ba1d013f1 == reflect.Float32:
		__obf_cd65ed68b120e383.
			SetInt(int64(__obf_a30e8e069ebf8d56.Float()))
	case __obf_8c8f659ba1d013f1 == reflect.Bool && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		if __obf_a30e8e069ebf8d56.Bool() {
			__obf_cd65ed68b120e383.
				SetInt(1)
		} else {
			__obf_cd65ed68b120e383.
				SetInt(0)
		}
	case __obf_8c8f659ba1d013f1 == reflect.String && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_ae7e7e9fd3e1b096 := __obf_a30e8e069ebf8d56.String()
		if __obf_ae7e7e9fd3e1b096 == "" {
			__obf_ae7e7e9fd3e1b096 = "0"
		}
		__obf_6cbf6487a01daf31, __obf_6d8ef670b0943b3c := strconv.ParseInt(__obf_ae7e7e9fd3e1b096, 0, __obf_cd65ed68b120e383.Type().Bits())
		if __obf_6d8ef670b0943b3c == nil {
			__obf_cd65ed68b120e383.
				SetInt(__obf_6cbf6487a01daf31)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
	case __obf_9dd014252313d6cd.PkgPath() == "encoding/json" && __obf_9dd014252313d6cd.Name() == "Number":
		__obf_f52aeda9f67e10ec := __obf_e428e8b18837f120.(json.Number)
		__obf_6cbf6487a01daf31, __obf_6d8ef670b0943b3c := __obf_f52aeda9f67e10ec.Int64()
		if __obf_6d8ef670b0943b3c != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
		__obf_cd65ed68b120e383.
			SetInt(__obf_6cbf6487a01daf31)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_a30e8e069ebf8d56.Type(), __obf_e428e8b18837f120)
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_7c63467fdcb0d6ec(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8c8f659ba1d013f1 := __obf_ed31af980b14987e(__obf_a30e8e069ebf8d56)
	__obf_9dd014252313d6cd := __obf_a30e8e069ebf8d56.Type()

	switch {
	case __obf_8c8f659ba1d013f1 == reflect.Int:
		__obf_6cbf6487a01daf31 := __obf_a30e8e069ebf8d56.Int()
		if __obf_6cbf6487a01daf31 < 0 && !__obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %d overflows uint", __obf_97d979f40f648bde, __obf_6cbf6487a01daf31)
		}
		__obf_cd65ed68b120e383.
			SetUint(uint64(__obf_6cbf6487a01daf31))
	case __obf_8c8f659ba1d013f1 == reflect.Uint:
		__obf_cd65ed68b120e383.
			SetUint(__obf_a30e8e069ebf8d56.Uint())
	case __obf_8c8f659ba1d013f1 == reflect.Float32:
		__obf_502bda9c3cc67a99 := __obf_a30e8e069ebf8d56.Float()
		if __obf_502bda9c3cc67a99 < 0 && !__obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %f overflows uint", __obf_97d979f40f648bde, __obf_502bda9c3cc67a99)
		}
		__obf_cd65ed68b120e383.
			SetUint(uint64(__obf_502bda9c3cc67a99))
	case __obf_8c8f659ba1d013f1 == reflect.Bool && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		if __obf_a30e8e069ebf8d56.Bool() {
			__obf_cd65ed68b120e383.
				SetUint(1)
		} else {
			__obf_cd65ed68b120e383.
				SetUint(0)
		}
	case __obf_8c8f659ba1d013f1 == reflect.String && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_ae7e7e9fd3e1b096 := __obf_a30e8e069ebf8d56.String()
		if __obf_ae7e7e9fd3e1b096 == "" {
			__obf_ae7e7e9fd3e1b096 = "0"
		}
		__obf_6cbf6487a01daf31, __obf_6d8ef670b0943b3c := strconv.ParseUint(__obf_ae7e7e9fd3e1b096, 0, __obf_cd65ed68b120e383.Type().Bits())
		if __obf_6d8ef670b0943b3c == nil {
			__obf_cd65ed68b120e383.
				SetUint(__obf_6cbf6487a01daf31)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
	case __obf_9dd014252313d6cd.PkgPath() == "encoding/json" && __obf_9dd014252313d6cd.Name() == "Number":
		__obf_f52aeda9f67e10ec := __obf_e428e8b18837f120.(json.Number)
		__obf_6cbf6487a01daf31, __obf_6d8ef670b0943b3c := strconv.ParseUint(string(__obf_f52aeda9f67e10ec), 0, 64)
		if __obf_6d8ef670b0943b3c != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
		__obf_cd65ed68b120e383.
			SetUint(__obf_6cbf6487a01daf31)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_a30e8e069ebf8d56.Type(), __obf_e428e8b18837f120)
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_bd1426afee26ea4f(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8c8f659ba1d013f1 := __obf_ed31af980b14987e(__obf_a30e8e069ebf8d56)

	switch {
	case __obf_8c8f659ba1d013f1 == reflect.Bool:
		__obf_cd65ed68b120e383.
			SetBool(__obf_a30e8e069ebf8d56.Bool())
	case __obf_8c8f659ba1d013f1 == reflect.Int && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_cd65ed68b120e383.
			SetBool(__obf_a30e8e069ebf8d56.Int() != 0)
	case __obf_8c8f659ba1d013f1 == reflect.Uint && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_cd65ed68b120e383.
			SetBool(__obf_a30e8e069ebf8d56.Uint() != 0)
	case __obf_8c8f659ba1d013f1 == reflect.Float32 && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_cd65ed68b120e383.
			SetBool(__obf_a30e8e069ebf8d56.Float() != 0)
	case __obf_8c8f659ba1d013f1 == reflect.String && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_ed3004c91f084575, __obf_6d8ef670b0943b3c := strconv.ParseBool(__obf_a30e8e069ebf8d56.String())
		if __obf_6d8ef670b0943b3c == nil {
			__obf_cd65ed68b120e383.
				SetBool(__obf_ed3004c91f084575)
		} else if __obf_a30e8e069ebf8d56.String() == "" {
			__obf_cd65ed68b120e383.
				SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_a30e8e069ebf8d56.Type(), __obf_e428e8b18837f120)
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_02660e9348752825(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8c8f659ba1d013f1 := __obf_ed31af980b14987e(__obf_a30e8e069ebf8d56)
	__obf_9dd014252313d6cd := __obf_a30e8e069ebf8d56.Type()

	switch {
	case __obf_8c8f659ba1d013f1 == reflect.Int:
		__obf_cd65ed68b120e383.
			SetFloat(float64(__obf_a30e8e069ebf8d56.Int()))
	case __obf_8c8f659ba1d013f1 == reflect.Uint:
		__obf_cd65ed68b120e383.
			SetFloat(float64(__obf_a30e8e069ebf8d56.Uint()))
	case __obf_8c8f659ba1d013f1 == reflect.Float32:
		__obf_cd65ed68b120e383.
			SetFloat(__obf_a30e8e069ebf8d56.Float())
	case __obf_8c8f659ba1d013f1 == reflect.Bool && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		if __obf_a30e8e069ebf8d56.Bool() {
			__obf_cd65ed68b120e383.
				SetFloat(1)
		} else {
			__obf_cd65ed68b120e383.
				SetFloat(0)
		}
	case __obf_8c8f659ba1d013f1 == reflect.String && __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput:
		__obf_ae7e7e9fd3e1b096 := __obf_a30e8e069ebf8d56.String()
		if __obf_ae7e7e9fd3e1b096 == "" {
			__obf_ae7e7e9fd3e1b096 = "0"
		}
		__obf_502bda9c3cc67a99, __obf_6d8ef670b0943b3c := strconv.ParseFloat(__obf_ae7e7e9fd3e1b096, __obf_cd65ed68b120e383.Type().Bits())
		if __obf_6d8ef670b0943b3c == nil {
			__obf_cd65ed68b120e383.
				SetFloat(__obf_502bda9c3cc67a99)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
	case __obf_9dd014252313d6cd.PkgPath() == "encoding/json" && __obf_9dd014252313d6cd.Name() == "Number":
		__obf_f52aeda9f67e10ec := __obf_e428e8b18837f120.(json.Number)
		__obf_6cbf6487a01daf31, __obf_6d8ef670b0943b3c := __obf_f52aeda9f67e10ec.Float64()
		if __obf_6d8ef670b0943b3c != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_97d979f40f648bde, __obf_6d8ef670b0943b3c)
		}
		__obf_cd65ed68b120e383.
			SetFloat(__obf_6cbf6487a01daf31)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_a30e8e069ebf8d56.Type(), __obf_e428e8b18837f120)
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_d20c19b1598726f2(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_d438fc1373de555c := __obf_cd65ed68b120e383.Type()
	__obf_b4220ef79b7777b4 := __obf_d438fc1373de555c.Key()
	__obf_672842fe2f3a868a := __obf_d438fc1373de555c.Elem()
	__obf_8cf6c124f55656c2 := // By default we overwrite keys in the current map
		__obf_cd65ed68b120e383

	// If the map is nil or we're purposely zeroing fields, make a new map
	if __obf_8cf6c124f55656c2.IsNil() || __obf_15070727b144cc26.__obf_da116ef5edc24797.ZeroFields {
		__obf_00862f74a2b44518 := // Make a new map to hold our result
			reflect.MapOf(__obf_b4220ef79b7777b4, __obf_672842fe2f3a868a)
		__obf_8cf6c124f55656c2 = reflect.MakeMap(__obf_00862f74a2b44518)
	}
	__obf_a30e8e069ebf8d56 := // Check input type and based on the input type jump to the proper func
		reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	switch __obf_a30e8e069ebf8d56.Kind() {
	case reflect.Map:
		return __obf_15070727b144cc26.__obf_2c810e860cf3aead(__obf_97d979f40f648bde, __obf_a30e8e069ebf8d56, __obf_cd65ed68b120e383, __obf_8cf6c124f55656c2)

	case reflect.Struct:
		return __obf_15070727b144cc26.__obf_612b7a47c762a7ba(__obf_a30e8e069ebf8d56, __obf_cd65ed68b120e383, __obf_8cf6c124f55656c2)

	case reflect.Array, reflect.Slice:
		if __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput {
			return __obf_15070727b144cc26.__obf_1c4af7f1d55174d8(__obf_97d979f40f648bde, __obf_a30e8e069ebf8d56, __obf_cd65ed68b120e383, __obf_8cf6c124f55656c2)
		}

		fallthrough

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_97d979f40f648bde, __obf_a30e8e069ebf8d56.Kind())
	}
}

func (__obf_15070727b144cc26 *Decoder) __obf_1c4af7f1d55174d8(__obf_97d979f40f648bde string, __obf_a30e8e069ebf8d56 reflect.Value, __obf_cd65ed68b120e383 reflect.Value, __obf_8cf6c124f55656c2 reflect.Value) error {
	// Special case for BC reasons (covered by tests)
	if __obf_a30e8e069ebf8d56.Len() == 0 {
		__obf_cd65ed68b120e383.
			Set(__obf_8cf6c124f55656c2)
		return nil
	}

	for __obf_6cbf6487a01daf31 := 0; __obf_6cbf6487a01daf31 < __obf_a30e8e069ebf8d56.Len(); __obf_6cbf6487a01daf31++ {
		__obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_97d979f40f648bde+
			"["+strconv.Itoa(__obf_6cbf6487a01daf31)+"]", __obf_a30e8e069ebf8d56.
			Index(__obf_6cbf6487a01daf31).Interface(), __obf_cd65ed68b120e383)
		if __obf_6d8ef670b0943b3c != nil {
			return __obf_6d8ef670b0943b3c
		}
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_2c810e860cf3aead(__obf_97d979f40f648bde string, __obf_a30e8e069ebf8d56 reflect.Value, __obf_cd65ed68b120e383 reflect.Value, __obf_8cf6c124f55656c2 reflect.Value) error {
	__obf_d438fc1373de555c := __obf_cd65ed68b120e383.Type()
	__obf_b4220ef79b7777b4 := __obf_d438fc1373de555c.Key()
	__obf_672842fe2f3a868a := __obf_d438fc1373de555c.Elem()

	// Accumulate errors
	errors := make([]string, 0)

	// If the input data is empty, then we just match what the input data is.
	if __obf_a30e8e069ebf8d56.Len() == 0 {
		if __obf_a30e8e069ebf8d56.IsNil() {
			if !__obf_cd65ed68b120e383.IsNil() {
				__obf_cd65ed68b120e383.
					Set(__obf_a30e8e069ebf8d56)
			}
		} else {
			__obf_cd65ed68b120e383.
				// Set to empty allocated value
				Set(__obf_8cf6c124f55656c2)
		}

		return nil
	}

	for _, __obf_4386f256ac305461 := range __obf_a30e8e069ebf8d56.MapKeys() {
		__obf_1cc7b5d4481b301c := __obf_97d979f40f648bde + "[" + __obf_4386f256ac305461.String() + "]"
		__obf_ec6027ba208eccaa := // First decode the key into the proper type
			reflect.Indirect(reflect.New(__obf_b4220ef79b7777b4))
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_1cc7b5d4481b301c, __obf_4386f256ac305461.Interface(), __obf_ec6027ba208eccaa); __obf_6d8ef670b0943b3c != nil {
			errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
			continue
		}
		__obf_7646bc0f4c5c46a0 := // Next decode the data into the proper type
			__obf_a30e8e069ebf8d56.MapIndex(__obf_4386f256ac305461).Interface()
		__obf_856bd6b000743bde := reflect.Indirect(reflect.New(__obf_672842fe2f3a868a))
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_1cc7b5d4481b301c, __obf_7646bc0f4c5c46a0, __obf_856bd6b000743bde); __obf_6d8ef670b0943b3c != nil {
			errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
			continue
		}
		__obf_8cf6c124f55656c2.
			SetMapIndex(__obf_ec6027ba208eccaa, __obf_856bd6b000743bde)
	}
	__obf_cd65ed68b120e383.

		// Set the built up map to the value
		Set(__obf_8cf6c124f55656c2)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_612b7a47c762a7ba(__obf_a30e8e069ebf8d56 reflect.Value, __obf_cd65ed68b120e383 reflect.Value, __obf_8cf6c124f55656c2 reflect.Value) error {
	__obf_db92856142f658a1 := __obf_a30e8e069ebf8d56.Type()
	for __obf_6cbf6487a01daf31 := 0; __obf_6cbf6487a01daf31 < __obf_db92856142f658a1.NumField(); __obf_6cbf6487a01daf31++ {
		__obf_502bda9c3cc67a99 := // Get the StructField first since this is a cheap operation. If the
			// field is unexported, then ignore it.
			__obf_db92856142f658a1.Field(__obf_6cbf6487a01daf31)
		if __obf_502bda9c3cc67a99.PkgPath != "" {
			continue
		}
		__obf_7646bc0f4c5c46a0 := // Next get the actual value of this field and verify it is assignable
			// to the map value.
			__obf_a30e8e069ebf8d56.Field(__obf_6cbf6487a01daf31)
		if !__obf_7646bc0f4c5c46a0.Type().AssignableTo(__obf_8cf6c124f55656c2.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", __obf_7646bc0f4c5c46a0.Type(), __obf_8cf6c124f55656c2.Type().Elem())
		}
		__obf_e17241ba1062a03c := __obf_502bda9c3cc67a99.Tag.Get(__obf_15070727b144cc26.__obf_da116ef5edc24797.TagName)
		__obf_5d6685223a3ba1df := __obf_502bda9c3cc67a99.Name

		if __obf_e17241ba1062a03c == "" && __obf_15070727b144cc26.__obf_da116ef5edc24797.IgnoreUntaggedFields {
			continue
		}
		__obf_d1f207c6142a9a90 := // If Squash is set in the config, we squash the field down.
			__obf_15070727b144cc26.__obf_da116ef5edc24797.Squash && __obf_7646bc0f4c5c46a0.Kind() == reflect.Struct && __obf_502bda9c3cc67a99.Anonymous
		__obf_7646bc0f4c5c46a0 = __obf_dfbab38f747ff8d8(__obf_7646bc0f4c5c46a0, __obf_15070727b144cc26.

			// Determine the name of the key in the map
			__obf_da116ef5edc24797.TagName)

		if __obf_e10c6f5a852ef775 := strings.Index(__obf_e17241ba1062a03c, ","); __obf_e10c6f5a852ef775 != -1 {
			if __obf_e17241ba1062a03c[:__obf_e10c6f5a852ef775] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Contains(__obf_e17241ba1062a03c[__obf_e10c6f5a852ef775+1:], "omitempty") && __obf_84bf21308b86c1c2(__obf_7646bc0f4c5c46a0) {
				continue
			}
			__obf_d1f207c6142a9a90 = // If "squash" is specified in the tag, we squash the field down.
				__obf_d1f207c6142a9a90 || strings.Contains(__obf_e17241ba1062a03c[__obf_e10c6f5a852ef775+1:], "squash")
			if __obf_d1f207c6142a9a90 {
				// When squashing, the embedded type can be a pointer to a struct.
				if __obf_7646bc0f4c5c46a0.Kind() == reflect.Ptr && __obf_7646bc0f4c5c46a0.Elem().Kind() == reflect.Struct {
					__obf_7646bc0f4c5c46a0 = __obf_7646bc0f4c5c46a0.Elem()
				}

				// The final type must be a struct
				if __obf_7646bc0f4c5c46a0.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", __obf_7646bc0f4c5c46a0.Type())
				}
			}
			if __obf_6611791189124fa7 := __obf_e17241ba1062a03c[:__obf_e10c6f5a852ef775]; __obf_6611791189124fa7 != "" {
				__obf_5d6685223a3ba1df = __obf_6611791189124fa7
			}
		} else if len(__obf_e17241ba1062a03c) > 0 {
			if __obf_e17241ba1062a03c == "-" {
				continue
			}
			__obf_5d6685223a3ba1df = __obf_e17241ba1062a03c
		}

		switch __obf_7646bc0f4c5c46a0.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			__obf_9ae550bfb49def15 := reflect.New(__obf_7646bc0f4c5c46a0.Type())
			__obf_9ae550bfb49def15.
				Elem().Set(__obf_7646bc0f4c5c46a0)
			__obf_9dac2155a9eeb452 := __obf_8cf6c124f55656c2.Type()
			__obf_a90c3542f1f338cc := __obf_9dac2155a9eeb452.Key()
			__obf_a697434ed18a5975 := __obf_9dac2155a9eeb452.Elem()
			__obf_78ccc4952eb5b2b7 := reflect.MapOf(__obf_a90c3542f1f338cc, __obf_a697434ed18a5975)
			__obf_fb5b6a3d64627b10 := reflect.MakeMap(__obf_78ccc4952eb5b2b7)
			__obf_3776f89115922937 := // Creating a pointer to a map so that other methods can completely
				// overwrite the map if need be (looking at you decodeMapFromMap). The
				// indirection allows the underlying map to be settable (CanSet() == true)
				// where as reflect.MakeMap returns an unsettable map.
				reflect.New(__obf_fb5b6a3d64627b10.Type())
			reflect.Indirect(__obf_3776f89115922937).Set(__obf_fb5b6a3d64627b10)
			__obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_5d6685223a3ba1df, __obf_9ae550bfb49def15.Interface(), reflect.Indirect(__obf_3776f89115922937))
			if __obf_6d8ef670b0943b3c != nil {
				return __obf_6d8ef670b0943b3c
			}
			__obf_fb5b6a3d64627b10 = // the underlying map may have been completely overwritten so pull
				// it indirectly out of the enclosing value.
				reflect.Indirect(__obf_3776f89115922937)

			if __obf_d1f207c6142a9a90 {
				for _, __obf_4386f256ac305461 := range __obf_fb5b6a3d64627b10.MapKeys() {
					__obf_8cf6c124f55656c2.
						SetMapIndex(__obf_4386f256ac305461, __obf_fb5b6a3d64627b10.MapIndex(__obf_4386f256ac305461))
				}
			} else {
				__obf_8cf6c124f55656c2.
					SetMapIndex(reflect.ValueOf(__obf_5d6685223a3ba1df), __obf_fb5b6a3d64627b10)
			}

		default:
			__obf_8cf6c124f55656c2.
				SetMapIndex(reflect.ValueOf(__obf_5d6685223a3ba1df), __obf_7646bc0f4c5c46a0)
		}
	}

	if __obf_cd65ed68b120e383.CanAddr() {
		__obf_cd65ed68b120e383.
			Set(__obf_8cf6c124f55656c2)
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_6b92db91c5345245(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) (bool, error) {
	__obf_97e25d22577459ba := // If the input data is nil, then we want to just set the output
		// pointer to be nil as well.
		__obf_e428e8b18837f120 == nil
	if !__obf_97e25d22577459ba {
		switch __obf_7646bc0f4c5c46a0 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120)); __obf_7646bc0f4c5c46a0.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			__obf_97e25d22577459ba = __obf_7646bc0f4c5c46a0.IsNil()
		}
	}
	if __obf_97e25d22577459ba {
		if !__obf_cd65ed68b120e383.IsNil() && __obf_cd65ed68b120e383.CanSet() {
			__obf_7058c0df19dd7f12 := reflect.New(__obf_cd65ed68b120e383.Type()).Elem()
			__obf_cd65ed68b120e383.
				Set(__obf_7058c0df19dd7f12)
		}

		return true, nil
	}
	__obf_d438fc1373de555c := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		__obf_cd65ed68b120e383.Type()
	__obf_672842fe2f3a868a := __obf_d438fc1373de555c.Elem()
	if __obf_cd65ed68b120e383.CanSet() {
		__obf_fc37f2cc1e72fc7c := __obf_cd65ed68b120e383
		if __obf_fc37f2cc1e72fc7c.IsNil() || __obf_15070727b144cc26.__obf_da116ef5edc24797.ZeroFields {
			__obf_fc37f2cc1e72fc7c = reflect.New(__obf_672842fe2f3a868a)
		}

		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_97d979f40f648bde, __obf_e428e8b18837f120, reflect.Indirect(__obf_fc37f2cc1e72fc7c)); __obf_6d8ef670b0943b3c != nil {
			return false, __obf_6d8ef670b0943b3c
		}
		__obf_cd65ed68b120e383.
			Set(__obf_fc37f2cc1e72fc7c)
	} else {
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_97d979f40f648bde, __obf_e428e8b18837f120, reflect.Indirect(__obf_cd65ed68b120e383)); __obf_6d8ef670b0943b3c != nil {
			return false, __obf_6d8ef670b0943b3c
		}
	}
	return false, nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_d4a936768e419fa1(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	if __obf_cd65ed68b120e383.Type() != __obf_a30e8e069ebf8d56.Type() {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_97d979f40f648bde, __obf_cd65ed68b120e383.
				Type(), __obf_a30e8e069ebf8d56.Type(), __obf_e428e8b18837f120)
	}
	__obf_cd65ed68b120e383.
		Set(__obf_a30e8e069ebf8d56)
	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_a1d958c6ee649900(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8579ec0b420909df := __obf_a30e8e069ebf8d56.Kind()
	__obf_d438fc1373de555c := __obf_cd65ed68b120e383.Type()
	__obf_672842fe2f3a868a := __obf_d438fc1373de555c.Elem()
	__obf_34ea561b47bbb0dc := reflect.SliceOf(__obf_672842fe2f3a868a)

	// If we have a non array/slice type then we first attempt to convert.
	if __obf_8579ec0b420909df != reflect.Array && __obf_8579ec0b420909df != reflect.Slice {
		if __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case __obf_8579ec0b420909df == reflect.Slice, __obf_8579ec0b420909df == reflect.Array:
				break

			// Empty maps turn into empty slices
			case __obf_8579ec0b420909df == reflect.Map:
				if __obf_a30e8e069ebf8d56.Len() == 0 {
					__obf_cd65ed68b120e383.
						Set(reflect.MakeSlice(__obf_34ea561b47bbb0dc, 0, 0))
					return nil
				}
				// Create slice of maps of other sizes
				return __obf_15070727b144cc26.__obf_a1d958c6ee649900(__obf_97d979f40f648bde, []any{__obf_e428e8b18837f120}, __obf_cd65ed68b120e383)

			case __obf_8579ec0b420909df == reflect.String && __obf_672842fe2f3a868a.Kind() == reflect.Uint8:
				return __obf_15070727b144cc26.__obf_a1d958c6ee649900(__obf_97d979f40f648bde, []byte(__obf_a30e8e069ebf8d56.String()), __obf_cd65ed68b120e383)

			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				return __obf_15070727b144cc26.__obf_a1d958c6ee649900(__obf_97d979f40f648bde, []any{__obf_e428e8b18837f120}, __obf_cd65ed68b120e383)
			}
		}

		return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", __obf_97d979f40f648bde, __obf_8579ec0b420909df)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if __obf_8579ec0b420909df != reflect.Array && __obf_a30e8e069ebf8d56.IsNil() {
		return nil
	}
	__obf_5bec2e3cdc1f320e := __obf_cd65ed68b120e383
	if __obf_5bec2e3cdc1f320e.IsNil() || __obf_15070727b144cc26.__obf_da116ef5edc24797.ZeroFields {
		__obf_5bec2e3cdc1f320e = // Make a new slice to hold our result, same size as the original data.
			reflect.MakeSlice(__obf_34ea561b47bbb0dc, __obf_a30e8e069ebf8d56.Len(), __obf_a30e8e069ebf8d56.Len())
	} else if __obf_5bec2e3cdc1f320e.Len() > __obf_a30e8e069ebf8d56.Len() {
		__obf_5bec2e3cdc1f320e = __obf_5bec2e3cdc1f320e.Slice(0, __obf_a30e8e069ebf8d56.Len())
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_6cbf6487a01daf31 := 0; __obf_6cbf6487a01daf31 < __obf_a30e8e069ebf8d56.Len(); __obf_6cbf6487a01daf31++ {
		__obf_5c777374effb5543 := __obf_a30e8e069ebf8d56.Index(__obf_6cbf6487a01daf31).Interface()
		for __obf_5bec2e3cdc1f320e.Len() <= __obf_6cbf6487a01daf31 {
			__obf_5bec2e3cdc1f320e = reflect.Append(__obf_5bec2e3cdc1f320e, reflect.Zero(__obf_672842fe2f3a868a))
		}
		__obf_f88502ed8d5261eb := __obf_5bec2e3cdc1f320e.Index(__obf_6cbf6487a01daf31)
		__obf_1cc7b5d4481b301c := __obf_97d979f40f648bde + "[" + strconv.Itoa(__obf_6cbf6487a01daf31) + "]"
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_1cc7b5d4481b301c, __obf_5c777374effb5543, __obf_f88502ed8d5261eb); __obf_6d8ef670b0943b3c != nil {
			errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
		}
	}
	__obf_cd65ed68b120e383.

		// Finally, set the value to the slice we built up
		Set(__obf_5bec2e3cdc1f320e)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_912e64085a8c4c83(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))
	__obf_8579ec0b420909df := __obf_a30e8e069ebf8d56.Kind()
	__obf_d438fc1373de555c := __obf_cd65ed68b120e383.Type()
	__obf_672842fe2f3a868a := __obf_d438fc1373de555c.Elem()
	__obf_ecc57b2debdd7e29 := reflect.ArrayOf(__obf_d438fc1373de555c.Len(), __obf_672842fe2f3a868a)
	__obf_17a36d0e350f9611 := __obf_cd65ed68b120e383

	if __obf_17a36d0e350f9611.Interface() == reflect.Zero(__obf_17a36d0e350f9611.Type()).Interface() || __obf_15070727b144cc26.__obf_da116ef5edc24797.ZeroFields {
		// Check input type
		if __obf_8579ec0b420909df != reflect.Array && __obf_8579ec0b420909df != reflect.Slice {
			if __obf_15070727b144cc26.__obf_da116ef5edc24797.WeaklyTypedInput {
				switch __obf_8579ec0b420909df {
				// Empty maps turn into empty arrays
				case reflect.Map:
					if __obf_a30e8e069ebf8d56.Len() == 0 {
						__obf_cd65ed68b120e383.
							Set(reflect.Zero(__obf_ecc57b2debdd7e29))
						return nil
					}

				// All other types we try to convert to the array type
				// and "lift" it into it. i.e. a string becomes a string array.
				default:
					// Just re-try this function with data as a slice.
					return __obf_15070727b144cc26.__obf_912e64085a8c4c83(__obf_97d979f40f648bde, []any{__obf_e428e8b18837f120}, __obf_cd65ed68b120e383)
				}
			}

			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", __obf_97d979f40f648bde, __obf_8579ec0b420909df)

		}
		if __obf_a30e8e069ebf8d56.Len() > __obf_ecc57b2debdd7e29.Len() {
			return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", __obf_97d979f40f648bde, __obf_ecc57b2debdd7e29.Len(), __obf_a30e8e069ebf8d56.Len())

		}
		__obf_17a36d0e350f9611 = // Make a new array to hold our result, same size as the original data.
			reflect.New(__obf_ecc57b2debdd7e29).Elem()
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_6cbf6487a01daf31 := 0; __obf_6cbf6487a01daf31 < __obf_a30e8e069ebf8d56.Len(); __obf_6cbf6487a01daf31++ {
		__obf_5c777374effb5543 := __obf_a30e8e069ebf8d56.Index(__obf_6cbf6487a01daf31).Interface()
		__obf_f88502ed8d5261eb := __obf_17a36d0e350f9611.Index(__obf_6cbf6487a01daf31)
		__obf_1cc7b5d4481b301c := __obf_97d979f40f648bde + "[" + strconv.Itoa(__obf_6cbf6487a01daf31) + "]"
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_1cc7b5d4481b301c, __obf_5c777374effb5543, __obf_f88502ed8d5261eb); __obf_6d8ef670b0943b3c != nil {
			errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
		}
	}
	__obf_cd65ed68b120e383.

		// Finally, set the value to the array we built up
		Set(__obf_17a36d0e350f9611)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_15070727b144cc26 *Decoder) __obf_577e93fa4008dcd6(__obf_97d979f40f648bde string, __obf_e428e8b18837f120 any, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_a30e8e069ebf8d56 := reflect.Indirect(reflect.ValueOf(__obf_e428e8b18837f120))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if __obf_a30e8e069ebf8d56.Type() == __obf_cd65ed68b120e383.Type() {
		__obf_cd65ed68b120e383.
			Set(__obf_a30e8e069ebf8d56)
		return nil
	}
	__obf_8579ec0b420909df := __obf_a30e8e069ebf8d56.Kind()
	switch __obf_8579ec0b420909df {
	case reflect.Map:
		return __obf_15070727b144cc26.__obf_d53ca9cc93fd7474(__obf_97d979f40f648bde, __obf_a30e8e069ebf8d56,

			// Not the most efficient way to do this but we can optimize later if
			// we want to. To convert from struct to struct we go to map first
			// as an intermediary.
			__obf_cd65ed68b120e383)

	case reflect.Struct:
		__obf_00862f74a2b44518 := // Make a new map to hold our result
			reflect.TypeOf((map[string]any)(nil))
		__obf_f9cf6792d44638d4 := reflect.MakeMap(__obf_00862f74a2b44518)
		__obf_3776f89115922937 := // Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			reflect.New(__obf_f9cf6792d44638d4.Type())

		reflect.Indirect(__obf_3776f89115922937).Set(__obf_f9cf6792d44638d4)
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_612b7a47c762a7ba(__obf_a30e8e069ebf8d56, reflect.Indirect(__obf_3776f89115922937), __obf_f9cf6792d44638d4); __obf_6d8ef670b0943b3c != nil {
			return __obf_6d8ef670b0943b3c
		}
		__obf_bb9d1909c7b1c125 := __obf_15070727b144cc26.__obf_d53ca9cc93fd7474(__obf_97d979f40f648bde, reflect.Indirect(__obf_3776f89115922937), __obf_cd65ed68b120e383)
		return __obf_bb9d1909c7b1c125

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_97d979f40f648bde, __obf_a30e8e069ebf8d56.Kind())
	}
}

func (__obf_15070727b144cc26 *Decoder) __obf_d53ca9cc93fd7474(__obf_97d979f40f648bde string, __obf_a30e8e069ebf8d56, __obf_cd65ed68b120e383 reflect.Value) error {
	__obf_8d4acaad13f3fa10 := __obf_a30e8e069ebf8d56.Type()
	if __obf_165b367257420d62 := __obf_8d4acaad13f3fa10.Key().Kind(); __obf_165b367257420d62 != reflect.String && __obf_165b367257420d62 != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys", __obf_97d979f40f648bde, __obf_8d4acaad13f3fa10.
				Key().Kind())
	}
	__obf_22b6d54822b48b0a := make(map[reflect.Value]struct{})
	__obf_f0ba9b1aa666a7d3 := make(map[any]struct{})
	for _, __obf_f84959325360dc86 := range __obf_a30e8e069ebf8d56.MapKeys() {
		__obf_22b6d54822b48b0a[__obf_f84959325360dc86] = struct{}{}
		__obf_f0ba9b1aa666a7d3[__obf_f84959325360dc86.Interface()] = struct{}{}
	}
	__obf_53a61ba0f0b5bc0e := make(map[any]struct{})
	errors := make([]string, 0)
	__obf_a25f0c212c30d176 := // This slice will keep track of all the structs we'll be decoding.
		// There can be more than one struct if there are embedded structs
		// that are squashed.
		make([]reflect.Value, 1, 5)
	__obf_a25f0c212c30d176[0] = __obf_cd65ed68b120e383

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type __obf_ae6c7d28571e7800 struct {
		__obf_ae6c7d28571e7800 reflect.StructField
		__obf_cd65ed68b120e383 reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var __obf_7a5fdf08eefe7903 *__obf_ae6c7d28571e7800
	__obf_854397262555beb2 := []__obf_ae6c7d28571e7800{}
	for len(__obf_a25f0c212c30d176) > 0 {
		__obf_a4e23eee89aea029 := __obf_a25f0c212c30d176[0]
		__obf_a25f0c212c30d176 = __obf_a25f0c212c30d176[1:]
		__obf_fb57b7d8bda927e0 := __obf_a4e23eee89aea029.Type()

		for __obf_6cbf6487a01daf31 := 0; __obf_6cbf6487a01daf31 < __obf_fb57b7d8bda927e0.NumField(); __obf_6cbf6487a01daf31++ {
			__obf_01a55dda7e04907b := __obf_fb57b7d8bda927e0.Field(__obf_6cbf6487a01daf31)
			__obf_3a8d69c90600142d := __obf_a4e23eee89aea029.Field(__obf_6cbf6487a01daf31)
			if __obf_3a8d69c90600142d.Kind() == reflect.Pointer && __obf_3a8d69c90600142d.Elem().Kind() == reflect.Struct {
				__obf_3a8d69c90600142d = // Handle embedded struct pointers as embedded structs.
					__obf_3a8d69c90600142d.Elem()
			}
			__obf_d1f207c6142a9a90 := // If "squash" is specified in the tag, we squash the field down.
				__obf_15070727b144cc26.__obf_da116ef5edc24797.Squash && __obf_3a8d69c90600142d.Kind() == reflect.Struct && __obf_01a55dda7e04907b.Anonymous
			__obf_a4e48e694db32ca8 := false
			__obf_5692049e9329b25a := // We always parse the tags cause we're looking for other tags too
				strings.Split(__obf_01a55dda7e04907b.Tag.Get(__obf_15070727b144cc26.__obf_da116ef5edc24797.TagName), ",")
			for _, __obf_33a47078edbe9a33 := range __obf_5692049e9329b25a[1:] {
				if __obf_33a47078edbe9a33 == "squash" {
					__obf_d1f207c6142a9a90 = true
					break
				}

				if __obf_33a47078edbe9a33 == "remain" {
					__obf_a4e48e694db32ca8 = true
					break
				}
			}

			if __obf_d1f207c6142a9a90 {
				if __obf_3a8d69c90600142d.Kind() != reflect.Struct {
					errors = __obf_9ddf59feabf88c65(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", __obf_01a55dda7e04907b.Name, __obf_3a8d69c90600142d.Kind()))
				} else {
					__obf_a25f0c212c30d176 = append(__obf_a25f0c212c30d176, __obf_3a8d69c90600142d)
				}
				continue
			}

			// Build our field
			if __obf_a4e48e694db32ca8 {
				__obf_7a5fdf08eefe7903 = &__obf_ae6c7d28571e7800{__obf_01a55dda7e04907b, __obf_3a8d69c90600142d}
			} else {
				__obf_854397262555beb2 = // Normal struct field, store it away
					append(__obf_854397262555beb2, __obf_ae6c7d28571e7800{__obf_01a55dda7e04907b, __obf_3a8d69c90600142d})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, __obf_502bda9c3cc67a99 := range __obf_854397262555beb2 {
		__obf_ae6c7d28571e7800, __obf_c86b228d22ad2b75 := __obf_502bda9c3cc67a99.__obf_ae6c7d28571e7800, __obf_502bda9c3cc67a99.__obf_cd65ed68b120e383
		__obf_1cc7b5d4481b301c := __obf_ae6c7d28571e7800.Name
		__obf_e17241ba1062a03c := __obf_ae6c7d28571e7800.Tag.Get(__obf_15070727b144cc26.__obf_da116ef5edc24797.TagName)
		__obf_e17241ba1062a03c = strings.SplitN(__obf_e17241ba1062a03c, ",", 2)[0]
		if __obf_e17241ba1062a03c != "" {
			__obf_1cc7b5d4481b301c = __obf_e17241ba1062a03c
		}
		__obf_d7967926eb7245f6 := reflect.ValueOf(__obf_1cc7b5d4481b301c)
		__obf_78765ef1a15ffab1 := __obf_a30e8e069ebf8d56.MapIndex(__obf_d7967926eb7245f6)
		if !__obf_78765ef1a15ffab1.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for __obf_f84959325360dc86 := range __obf_22b6d54822b48b0a {
				__obf_4ec00d5b7d6ad4cc, __obf_8a9bb3ba25fa94bd := __obf_f84959325360dc86.Interface().(string)
				if !__obf_8a9bb3ba25fa94bd {
					// Not a string key
					continue
				}

				if __obf_15070727b144cc26.__obf_da116ef5edc24797.MatchName(__obf_4ec00d5b7d6ad4cc, __obf_1cc7b5d4481b301c) {
					__obf_d7967926eb7245f6 = __obf_f84959325360dc86
					__obf_78765ef1a15ffab1 = __obf_a30e8e069ebf8d56.MapIndex(__obf_f84959325360dc86)
					break
				}
			}

			if !__obf_78765ef1a15ffab1.IsValid() {
				__obf_53a61ba0f0b5bc0e[ // There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				__obf_1cc7b5d4481b301c] = struct{}{}
				continue
			}
		}

		if !__obf_c86b228d22ad2b75.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !__obf_c86b228d22ad2b75.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(__obf_f0ba9b1aa666a7d3, __obf_d7967926eb7245f6.Interface())

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if __obf_97d979f40f648bde != "" {
			__obf_1cc7b5d4481b301c = __obf_97d979f40f648bde + "." + __obf_1cc7b5d4481b301c
		}

		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_54e09d10fb5e308c(__obf_1cc7b5d4481b301c, __obf_78765ef1a15ffab1.Interface(), __obf_c86b228d22ad2b75); __obf_6d8ef670b0943b3c != nil {
			errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if __obf_7a5fdf08eefe7903 != nil && len(__obf_f0ba9b1aa666a7d3) > 0 {
		__obf_a4e48e694db32ca8 := // Build a map of only the unused values
			map[any]any{}
		for __obf_a3663ecd8f9727ce := range __obf_f0ba9b1aa666a7d3 {
			__obf_a4e48e694db32ca8[__obf_a3663ecd8f9727ce] = __obf_a30e8e069ebf8d56.MapIndex(reflect.ValueOf(__obf_a3663ecd8f9727ce)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if __obf_6d8ef670b0943b3c := __obf_15070727b144cc26.__obf_d20c19b1598726f2(__obf_97d979f40f648bde, __obf_a4e48e694db32ca8, __obf_7a5fdf08eefe7903.__obf_cd65ed68b120e383); __obf_6d8ef670b0943b3c != nil {
			errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
		}
		__obf_f0ba9b1aa666a7d3 = // Set the map to nil so we have none so that the next check will
			// not error (ErrorUnused)
			nil
	}

	if __obf_15070727b144cc26.__obf_da116ef5edc24797.ErrorUnused && len(__obf_f0ba9b1aa666a7d3) > 0 {
		__obf_b6647af23b047f3b := make([]string, 0, len(__obf_f0ba9b1aa666a7d3))
		for __obf_c1eaee865d5276cb := range __obf_f0ba9b1aa666a7d3 {
			__obf_b6647af23b047f3b = append(__obf_b6647af23b047f3b, __obf_c1eaee865d5276cb.(string))
		}
		sort.Strings(__obf_b6647af23b047f3b)
		__obf_6d8ef670b0943b3c := fmt.Errorf("'%s' has invalid keys: %s", __obf_97d979f40f648bde, strings.Join(__obf_b6647af23b047f3b, ", "))
		errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
	}

	if __obf_15070727b144cc26.__obf_da116ef5edc24797.ErrorUnset && len(__obf_53a61ba0f0b5bc0e) > 0 {
		__obf_b6647af23b047f3b := make([]string, 0, len(__obf_53a61ba0f0b5bc0e))
		for __obf_c1eaee865d5276cb := range __obf_53a61ba0f0b5bc0e {
			__obf_b6647af23b047f3b = append(__obf_b6647af23b047f3b, __obf_c1eaee865d5276cb.(string))
		}
		sort.Strings(__obf_b6647af23b047f3b)
		__obf_6d8ef670b0943b3c := fmt.Errorf("'%s' has unset fields: %s", __obf_97d979f40f648bde, strings.Join(__obf_b6647af23b047f3b, ", "))
		errors = __obf_9ddf59feabf88c65(errors, __obf_6d8ef670b0943b3c)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	if __obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata != nil {
		for __obf_c1eaee865d5276cb := range __obf_f0ba9b1aa666a7d3 {
			__obf_a3663ecd8f9727ce := __obf_c1eaee865d5276cb.(string)
			if __obf_97d979f40f648bde != "" {
				__obf_a3663ecd8f9727ce = __obf_97d979f40f648bde + "." + __obf_a3663ecd8f9727ce
			}
			__obf_15070727b144cc26.__obf_da116ef5edc24797.
				Metadata.Unused = append(__obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata.Unused, __obf_a3663ecd8f9727ce)
		}
		for __obf_c1eaee865d5276cb := range __obf_53a61ba0f0b5bc0e {
			__obf_a3663ecd8f9727ce := __obf_c1eaee865d5276cb.(string)
			if __obf_97d979f40f648bde != "" {
				__obf_a3663ecd8f9727ce = __obf_97d979f40f648bde + "." + __obf_a3663ecd8f9727ce
			}
			__obf_15070727b144cc26.__obf_da116ef5edc24797.
				Metadata.Unset = append(__obf_15070727b144cc26.__obf_da116ef5edc24797.Metadata.Unset, __obf_a3663ecd8f9727ce)
		}
	}

	return nil
}

func __obf_84bf21308b86c1c2(__obf_7646bc0f4c5c46a0 reflect.Value) bool {
	switch __obf_ed31af980b14987e(__obf_7646bc0f4c5c46a0) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_7646bc0f4c5c46a0.Len() == 0
	case reflect.Bool:
		return !__obf_7646bc0f4c5c46a0.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_7646bc0f4c5c46a0.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_7646bc0f4c5c46a0.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_7646bc0f4c5c46a0.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return __obf_7646bc0f4c5c46a0.IsNil()
	}
	return false
}

func __obf_ed31af980b14987e(__obf_cd65ed68b120e383 reflect.Value) reflect.Kind {
	__obf_165b367257420d62 := __obf_cd65ed68b120e383.Kind()

	switch {
	case __obf_165b367257420d62 >= reflect.Int && __obf_165b367257420d62 <= reflect.Int64:
		return reflect.Int
	case __obf_165b367257420d62 >= reflect.Uint && __obf_165b367257420d62 <= reflect.Uint64:
		return reflect.Uint
	case __obf_165b367257420d62 >= reflect.Float32 && __obf_165b367257420d62 <= reflect.Float64:
		return reflect.Float32
	default:
		return __obf_165b367257420d62
	}
}

func __obf_d896a2373c16e434(__obf_db92856142f658a1 reflect.Type, __obf_dba5bc1b87e57f30 bool, __obf_8c4c3caad8c4e1b2 string) bool {
	for __obf_6cbf6487a01daf31 := 0; __obf_6cbf6487a01daf31 < __obf_db92856142f658a1.NumField(); __obf_6cbf6487a01daf31++ {
		__obf_502bda9c3cc67a99 := __obf_db92856142f658a1.Field(__obf_6cbf6487a01daf31)
		if __obf_502bda9c3cc67a99.PkgPath == "" && !__obf_dba5bc1b87e57f30 { // check for unexported fields
			return true
		}
		if __obf_dba5bc1b87e57f30 && __obf_502bda9c3cc67a99.Tag.Get(__obf_8c4c3caad8c4e1b2) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func __obf_dfbab38f747ff8d8(__obf_7646bc0f4c5c46a0 reflect.Value, __obf_8c4c3caad8c4e1b2 string) reflect.Value {
	if __obf_7646bc0f4c5c46a0.Kind() != reflect.Ptr || __obf_7646bc0f4c5c46a0.Elem().Kind() != reflect.Struct {
		return __obf_7646bc0f4c5c46a0
	}
	__obf_98a6f478e2260f40 := __obf_7646bc0f4c5c46a0.Elem()
	__obf_e6f41069f0fe28a7 := __obf_98a6f478e2260f40.Type()
	if __obf_d896a2373c16e434(__obf_e6f41069f0fe28a7, true, __obf_8c4c3caad8c4e1b2) {
		return __obf_98a6f478e2260f40
	}
	return __obf_7646bc0f4c5c46a0
}
