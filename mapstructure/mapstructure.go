package __obf_68a92d5117d8d921

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
type DecodeHookFuncValue func(__obf_793e64bff3015f1d reflect.Value, __obf_f4678c4e7943a8ce reflect.Value) (any, error)

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
	MatchName func(__obf_7dacff2cc7465865, __obf_465154ef45471268 string) bool
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	__obf_85fcbb7e5f1202f3 *DecoderConfig
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
func Decode(__obf_b4c78a01be9e73e2 any, __obf_c6589c87a4ce4950 any) error {
	__obf_85fcbb7e5f1202f3 := &DecoderConfig{
		Metadata: nil,
		Result:   __obf_c6589c87a4ce4950,
	}
	__obf_01ffd7dc0557f650, __obf_ef25cdf0d96b47a8 := NewDecoder(__obf_85fcbb7e5f1202f3)
	if __obf_ef25cdf0d96b47a8 != nil {
		return __obf_ef25cdf0d96b47a8
	}

	return __obf_01ffd7dc0557f650.Decode(__obf_b4c78a01be9e73e2)
}

// WeakDecode is the same as Decode but is shorthand to enable
// WeaklyTypedInput. See DecoderConfig for more info.
func WeakDecode(__obf_b4c78a01be9e73e2, __obf_c6589c87a4ce4950 any) error {
	__obf_85fcbb7e5f1202f3 := &DecoderConfig{
		Metadata:         nil,
		Result:           __obf_c6589c87a4ce4950,
		WeaklyTypedInput: true,
	}
	__obf_01ffd7dc0557f650, __obf_ef25cdf0d96b47a8 := NewDecoder(__obf_85fcbb7e5f1202f3)
	if __obf_ef25cdf0d96b47a8 != nil {
		return __obf_ef25cdf0d96b47a8
	}

	return __obf_01ffd7dc0557f650.Decode(__obf_b4c78a01be9e73e2)
}

// DecodeMetadata is the same as Decode, but is shorthand to
// enable metadata collection. See DecoderConfig for more info.
func DecodeMetadata(__obf_b4c78a01be9e73e2 any, __obf_c6589c87a4ce4950 any, __obf_e267105a7708636e *Metadata) error {
	__obf_85fcbb7e5f1202f3 := &DecoderConfig{
		Metadata: __obf_e267105a7708636e,
		Result:   __obf_c6589c87a4ce4950,
	}
	__obf_01ffd7dc0557f650, __obf_ef25cdf0d96b47a8 := NewDecoder(__obf_85fcbb7e5f1202f3)
	if __obf_ef25cdf0d96b47a8 != nil {
		return __obf_ef25cdf0d96b47a8
	}

	return __obf_01ffd7dc0557f650.Decode(__obf_b4c78a01be9e73e2)
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
// enable both WeaklyTypedInput and metadata collection. See
// DecoderConfig for more info.
func WeakDecodeMetadata(__obf_b4c78a01be9e73e2 any, __obf_c6589c87a4ce4950 any, __obf_e267105a7708636e *Metadata) error {
	__obf_85fcbb7e5f1202f3 := &DecoderConfig{
		Metadata:         __obf_e267105a7708636e,
		Result:           __obf_c6589c87a4ce4950,
		WeaklyTypedInput: true,
	}
	__obf_01ffd7dc0557f650, __obf_ef25cdf0d96b47a8 := NewDecoder(__obf_85fcbb7e5f1202f3)
	if __obf_ef25cdf0d96b47a8 != nil {
		return __obf_ef25cdf0d96b47a8
	}

	return __obf_01ffd7dc0557f650.Decode(__obf_b4c78a01be9e73e2)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(__obf_85fcbb7e5f1202f3 *DecoderConfig) (*Decoder, error) {
	__obf_306c5ad91737b908 := reflect.ValueOf(__obf_85fcbb7e5f1202f3.Result)
	if __obf_306c5ad91737b908.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}
	__obf_306c5ad91737b908 = __obf_306c5ad91737b908.Elem()
	if !__obf_306c5ad91737b908.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if __obf_85fcbb7e5f1202f3.Metadata != nil {
		if __obf_85fcbb7e5f1202f3.Metadata.Keys == nil {
			__obf_85fcbb7e5f1202f3.
				Metadata.Keys = make([]string, 0)
		}

		if __obf_85fcbb7e5f1202f3.Metadata.Unused == nil {
			__obf_85fcbb7e5f1202f3.
				Metadata.Unused = make([]string, 0)
		}

		if __obf_85fcbb7e5f1202f3.Metadata.Unset == nil {
			__obf_85fcbb7e5f1202f3.
				Metadata.Unset = make([]string, 0)
		}
	}

	if __obf_85fcbb7e5f1202f3.TagName == "" {
		__obf_85fcbb7e5f1202f3.
			TagName = "mapstructure"
	}

	if __obf_85fcbb7e5f1202f3.MatchName == nil {
		__obf_85fcbb7e5f1202f3.
			MatchName = strings.EqualFold
	}
	__obf_84b0717583bd55a4 := &Decoder{__obf_85fcbb7e5f1202f3: __obf_85fcbb7e5f1202f3}

	return __obf_84b0717583bd55a4, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (__obf_f16772980031b91a *Decoder) Decode(__obf_b4c78a01be9e73e2 any) error {
	return __obf_f16772980031b91a.__obf_26d8f56953be1262("", __obf_b4c78a01be9e73e2, reflect.ValueOf(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (__obf_f16772980031b91a *Decoder) __obf_26d8f56953be1262(__obf_bee876148a124a99 string, __obf_b4c78a01be9e73e2 any, __obf_54bc158726d6a1ec reflect.Value) error {
	var __obf_4ac7979984d0548b reflect.Value
	if __obf_b4c78a01be9e73e2 != nil {
		__obf_4ac7979984d0548b = reflect.ValueOf(__obf_b4c78a01be9e73e2)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if __obf_4ac7979984d0548b.Kind() == reflect.Ptr && __obf_4ac7979984d0548b.IsNil() {
			__obf_b4c78a01be9e73e2 = nil
		}
	}

	if __obf_b4c78a01be9e73e2 == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ZeroFields {
			__obf_54bc158726d6a1ec.
				Set(reflect.Zero(__obf_54bc158726d6a1ec.Type()))

			if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata != nil && __obf_bee876148a124a99 != "" {
				__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.
					Metadata.Keys = append(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata.Keys, __obf_bee876148a124a99)
			}
		}
		return nil
	}

	if !__obf_4ac7979984d0548b.IsValid() {
		__obf_54bc158726d6a1ec.
			// If the input value is invalid, then we just set the value
			// to be the zero value.
			Set(reflect.Zero(__obf_54bc158726d6a1ec.Type()))
		if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata != nil && __obf_bee876148a124a99 != "" {
			__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.
				Metadata.Keys = append(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata.Keys, __obf_bee876148a124a99)
		}
		return nil
	}

	if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the input.
		var __obf_ef25cdf0d96b47a8 error
		__obf_b4c78a01be9e73e2, __obf_ef25cdf0d96b47a8 = DecodeHookExec(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.DecodeHook, __obf_4ac7979984d0548b, __obf_54bc158726d6a1ec)
		if __obf_ef25cdf0d96b47a8 != nil {
			return fmt.Errorf("error decoding '%s': %w", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
	}

	var __obf_ef25cdf0d96b47a8 error
	__obf_facb1ff77232f7b4 := __obf_c3efe5973b2597dc(__obf_54bc158726d6a1ec)
	__obf_cb5d6ebe6208ab7a := true
	switch __obf_facb1ff77232f7b4 {
	case reflect.Bool:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_5d5ccd2a0c9e8d4a(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Interface:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_852fc4a5371db776(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.String:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_e9dc33c137088a85(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Int:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_b786b884f850dbd0(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Uint:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_88ea86f88374c87d(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Float32:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_12cad62884df7e1e(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Struct:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_195c05f649453920(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Map:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_f1681715c2229fb2(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Ptr:
		__obf_cb5d6ebe6208ab7a, __obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_3aa25401044d6204(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Slice:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_c102bb670129f9e6(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Array:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_978e1dbf26790dcd(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2, __obf_54bc158726d6a1ec)
	case reflect.Func:
		__obf_ef25cdf0d96b47a8 = __obf_f16772980031b91a.__obf_5db3819bb8f0f28f(__obf_bee876148a124a99, __obf_b4c78a01be9e73e2,

			// If we reached this point then we weren't able to decode it
			__obf_54bc158726d6a1ec)
	default:

		return fmt.Errorf("%s: unsupported type: %s", __obf_bee876148a124a99, __obf_facb1ff77232f7b4)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metainput.
	if __obf_cb5d6ebe6208ab7a && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata != nil && __obf_bee876148a124a99 != "" {
		__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.
			Metadata.Keys = append(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata.Keys, __obf_bee876148a124a99)
	}

	return __obf_ef25cdf0d96b47a8
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (__obf_f16772980031b91a *Decoder) __obf_852fc4a5371db776(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	if __obf_306c5ad91737b908.IsValid() && __obf_306c5ad91737b908.Elem().IsValid() {
		__obf_a4d888b3f7a359f5 := __obf_306c5ad91737b908.Elem()
		__obf_45571b163497001a := // If we can't address this element, then its not writable. Instead,
			// we make a copy of the value (which is a pointer and therefore
			// writable), decode into that, and replace the whole value.
			false
		if !__obf_a4d888b3f7a359f5.CanAddr() {
			__obf_45571b163497001a = true

			// Make *T
			copy := reflect.New(__obf_a4d888b3f7a359f5.Type())

			// *T = elem
			copy.Elem().Set(__obf_a4d888b3f7a359f5)
			__obf_a4d888b3f7a359f5 = // Set elem so we decode into it
				copy
		}

		// Decode. If we have an error then return. We also return right
		// away if we're not a copy because that means we decoded directly.
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_bee876148a124a99, __obf_d6e7451ec5237c6d, __obf_a4d888b3f7a359f5); __obf_ef25cdf0d96b47a8 != nil || !__obf_45571b163497001a {
			return __obf_ef25cdf0d96b47a8
		}
		__obf_306c5ad91737b908.

			// If we're a copy, we need to set te final result
			Set(__obf_a4d888b3f7a359f5.Elem())
		return nil
	}
	__obf_6001932d600cd7ec := reflect.ValueOf(__obf_d6e7451ec5237c6d)

	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if __obf_6001932d600cd7ec.Kind() == reflect.Ptr && __obf_6001932d600cd7ec.Type().Elem() == __obf_306c5ad91737b908.Type() {
		__obf_6001932d600cd7ec = reflect.Indirect(__obf_6001932d600cd7ec)
	}

	if !__obf_6001932d600cd7ec.IsValid() {
		__obf_6001932d600cd7ec = reflect.Zero(__obf_306c5ad91737b908.Type())
	}
	__obf_40f50f46d0d164df := __obf_6001932d600cd7ec.Type()
	if !__obf_40f50f46d0d164df.AssignableTo(__obf_306c5ad91737b908.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_40f50f46d0d164df)
	}
	__obf_306c5ad91737b908.
		Set(__obf_6001932d600cd7ec)
	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_e9dc33c137088a85(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_fd3ba775e1c89bcc := __obf_c3efe5973b2597dc(__obf_6001932d600cd7ec)
	__obf_f1c414f7fd62cfd4 := true
	switch {
	case __obf_fd3ba775e1c89bcc == reflect.String:
		__obf_306c5ad91737b908.
			SetString(__obf_6001932d600cd7ec.String())
	case __obf_fd3ba775e1c89bcc == reflect.Bool && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		if __obf_6001932d600cd7ec.Bool() {
			__obf_306c5ad91737b908.
				SetString("1")
		} else {
			__obf_306c5ad91737b908.
				SetString("0")
		}
	case __obf_fd3ba775e1c89bcc == reflect.Int && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_306c5ad91737b908.
			SetString(strconv.FormatInt(__obf_6001932d600cd7ec.Int(), 10))
	case __obf_fd3ba775e1c89bcc == reflect.Uint && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_306c5ad91737b908.
			SetString(strconv.FormatUint(__obf_6001932d600cd7ec.Uint(), 10))
	case __obf_fd3ba775e1c89bcc == reflect.Float32 && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_306c5ad91737b908.
			SetString(strconv.FormatFloat(__obf_6001932d600cd7ec.Float(), 'f', -1, 64))
	case __obf_fd3ba775e1c89bcc == reflect.Slice && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput, __obf_fd3ba775e1c89bcc ==
		reflect.Array && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_2fec4dc52e9242d2 := __obf_6001932d600cd7ec.Type()
		__obf_f1bf6756e05f60e0 := __obf_2fec4dc52e9242d2.Elem().Kind()
		switch __obf_f1bf6756e05f60e0 {
		case reflect.Uint8:
			var __obf_a189f68b6198ef81 []uint8
			if __obf_fd3ba775e1c89bcc == reflect.Array {
				__obf_a189f68b6198ef81 = make([]uint8, __obf_6001932d600cd7ec.Len())
				for __obf_a6629621259bc2bc := range __obf_a189f68b6198ef81 {
					__obf_a189f68b6198ef81[__obf_a6629621259bc2bc] = __obf_6001932d600cd7ec.Index(__obf_a6629621259bc2bc).Interface().(uint8)
				}
			} else {
				__obf_a189f68b6198ef81 = __obf_6001932d600cd7ec.Interface().([]uint8)
			}
			__obf_306c5ad91737b908.
				SetString(string(__obf_a189f68b6198ef81))
		default:
			__obf_f1c414f7fd62cfd4 = false
		}
	default:
		__obf_f1c414f7fd62cfd4 = false
	}

	if !__obf_f1c414f7fd62cfd4 {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_6001932d600cd7ec.Type(), __obf_d6e7451ec5237c6d)
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_b786b884f850dbd0(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_fd3ba775e1c89bcc := __obf_c3efe5973b2597dc(__obf_6001932d600cd7ec)
	__obf_2fec4dc52e9242d2 := __obf_6001932d600cd7ec.Type()

	switch {
	case __obf_fd3ba775e1c89bcc == reflect.Int:
		__obf_306c5ad91737b908.
			SetInt(__obf_6001932d600cd7ec.Int())
	case __obf_fd3ba775e1c89bcc == reflect.Uint:
		__obf_306c5ad91737b908.
			SetInt(int64(__obf_6001932d600cd7ec.Uint()))
	case __obf_fd3ba775e1c89bcc == reflect.Float32:
		__obf_306c5ad91737b908.
			SetInt(int64(__obf_6001932d600cd7ec.Float()))
	case __obf_fd3ba775e1c89bcc == reflect.Bool && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		if __obf_6001932d600cd7ec.Bool() {
			__obf_306c5ad91737b908.
				SetInt(1)
		} else {
			__obf_306c5ad91737b908.
				SetInt(0)
		}
	case __obf_fd3ba775e1c89bcc == reflect.String && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_415fa9df6fadc02e := __obf_6001932d600cd7ec.String()
		if __obf_415fa9df6fadc02e == "" {
			__obf_415fa9df6fadc02e = "0"
		}
		__obf_a6629621259bc2bc, __obf_ef25cdf0d96b47a8 := strconv.ParseInt(__obf_415fa9df6fadc02e, 0, __obf_306c5ad91737b908.Type().Bits())
		if __obf_ef25cdf0d96b47a8 == nil {
			__obf_306c5ad91737b908.
				SetInt(__obf_a6629621259bc2bc)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
	case __obf_2fec4dc52e9242d2.PkgPath() == "encoding/json" && __obf_2fec4dc52e9242d2.Name() == "Number":
		__obf_9817ff96c0db1b7c := __obf_d6e7451ec5237c6d.(json.Number)
		__obf_a6629621259bc2bc, __obf_ef25cdf0d96b47a8 := __obf_9817ff96c0db1b7c.Int64()
		if __obf_ef25cdf0d96b47a8 != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
		__obf_306c5ad91737b908.
			SetInt(__obf_a6629621259bc2bc)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_6001932d600cd7ec.Type(), __obf_d6e7451ec5237c6d)
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_88ea86f88374c87d(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_fd3ba775e1c89bcc := __obf_c3efe5973b2597dc(__obf_6001932d600cd7ec)
	__obf_2fec4dc52e9242d2 := __obf_6001932d600cd7ec.Type()

	switch {
	case __obf_fd3ba775e1c89bcc == reflect.Int:
		__obf_a6629621259bc2bc := __obf_6001932d600cd7ec.Int()
		if __obf_a6629621259bc2bc < 0 && !__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %d overflows uint", __obf_bee876148a124a99, __obf_a6629621259bc2bc)
		}
		__obf_306c5ad91737b908.
			SetUint(uint64(__obf_a6629621259bc2bc))
	case __obf_fd3ba775e1c89bcc == reflect.Uint:
		__obf_306c5ad91737b908.
			SetUint(__obf_6001932d600cd7ec.Uint())
	case __obf_fd3ba775e1c89bcc == reflect.Float32:
		__obf_9037ed2181de9473 := __obf_6001932d600cd7ec.Float()
		if __obf_9037ed2181de9473 < 0 && !__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %f overflows uint", __obf_bee876148a124a99, __obf_9037ed2181de9473)
		}
		__obf_306c5ad91737b908.
			SetUint(uint64(__obf_9037ed2181de9473))
	case __obf_fd3ba775e1c89bcc == reflect.Bool && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		if __obf_6001932d600cd7ec.Bool() {
			__obf_306c5ad91737b908.
				SetUint(1)
		} else {
			__obf_306c5ad91737b908.
				SetUint(0)
		}
	case __obf_fd3ba775e1c89bcc == reflect.String && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_415fa9df6fadc02e := __obf_6001932d600cd7ec.String()
		if __obf_415fa9df6fadc02e == "" {
			__obf_415fa9df6fadc02e = "0"
		}
		__obf_a6629621259bc2bc, __obf_ef25cdf0d96b47a8 := strconv.ParseUint(__obf_415fa9df6fadc02e, 0, __obf_306c5ad91737b908.Type().Bits())
		if __obf_ef25cdf0d96b47a8 == nil {
			__obf_306c5ad91737b908.
				SetUint(__obf_a6629621259bc2bc)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
	case __obf_2fec4dc52e9242d2.PkgPath() == "encoding/json" && __obf_2fec4dc52e9242d2.Name() == "Number":
		__obf_9817ff96c0db1b7c := __obf_d6e7451ec5237c6d.(json.Number)
		__obf_a6629621259bc2bc, __obf_ef25cdf0d96b47a8 := strconv.ParseUint(string(__obf_9817ff96c0db1b7c), 0, 64)
		if __obf_ef25cdf0d96b47a8 != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
		__obf_306c5ad91737b908.
			SetUint(__obf_a6629621259bc2bc)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_6001932d600cd7ec.Type(), __obf_d6e7451ec5237c6d)
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_5d5ccd2a0c9e8d4a(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_fd3ba775e1c89bcc := __obf_c3efe5973b2597dc(__obf_6001932d600cd7ec)

	switch {
	case __obf_fd3ba775e1c89bcc == reflect.Bool:
		__obf_306c5ad91737b908.
			SetBool(__obf_6001932d600cd7ec.Bool())
	case __obf_fd3ba775e1c89bcc == reflect.Int && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_306c5ad91737b908.
			SetBool(__obf_6001932d600cd7ec.Int() != 0)
	case __obf_fd3ba775e1c89bcc == reflect.Uint && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_306c5ad91737b908.
			SetBool(__obf_6001932d600cd7ec.Uint() != 0)
	case __obf_fd3ba775e1c89bcc == reflect.Float32 && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_306c5ad91737b908.
			SetBool(__obf_6001932d600cd7ec.Float() != 0)
	case __obf_fd3ba775e1c89bcc == reflect.String && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_dac936e7449e698a, __obf_ef25cdf0d96b47a8 := strconv.ParseBool(__obf_6001932d600cd7ec.String())
		if __obf_ef25cdf0d96b47a8 == nil {
			__obf_306c5ad91737b908.
				SetBool(__obf_dac936e7449e698a)
		} else if __obf_6001932d600cd7ec.String() == "" {
			__obf_306c5ad91737b908.
				SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_6001932d600cd7ec.Type(), __obf_d6e7451ec5237c6d)
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_12cad62884df7e1e(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_fd3ba775e1c89bcc := __obf_c3efe5973b2597dc(__obf_6001932d600cd7ec)
	__obf_2fec4dc52e9242d2 := __obf_6001932d600cd7ec.Type()

	switch {
	case __obf_fd3ba775e1c89bcc == reflect.Int:
		__obf_306c5ad91737b908.
			SetFloat(float64(__obf_6001932d600cd7ec.Int()))
	case __obf_fd3ba775e1c89bcc == reflect.Uint:
		__obf_306c5ad91737b908.
			SetFloat(float64(__obf_6001932d600cd7ec.Uint()))
	case __obf_fd3ba775e1c89bcc == reflect.Float32:
		__obf_306c5ad91737b908.
			SetFloat(__obf_6001932d600cd7ec.Float())
	case __obf_fd3ba775e1c89bcc == reflect.Bool && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		if __obf_6001932d600cd7ec.Bool() {
			__obf_306c5ad91737b908.
				SetFloat(1)
		} else {
			__obf_306c5ad91737b908.
				SetFloat(0)
		}
	case __obf_fd3ba775e1c89bcc == reflect.String && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput:
		__obf_415fa9df6fadc02e := __obf_6001932d600cd7ec.String()
		if __obf_415fa9df6fadc02e == "" {
			__obf_415fa9df6fadc02e = "0"
		}
		__obf_9037ed2181de9473, __obf_ef25cdf0d96b47a8 := strconv.ParseFloat(__obf_415fa9df6fadc02e, __obf_306c5ad91737b908.Type().Bits())
		if __obf_ef25cdf0d96b47a8 == nil {
			__obf_306c5ad91737b908.
				SetFloat(__obf_9037ed2181de9473)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
	case __obf_2fec4dc52e9242d2.PkgPath() == "encoding/json" && __obf_2fec4dc52e9242d2.Name() == "Number":
		__obf_9817ff96c0db1b7c := __obf_d6e7451ec5237c6d.(json.Number)
		__obf_a6629621259bc2bc, __obf_ef25cdf0d96b47a8 := __obf_9817ff96c0db1b7c.Float64()
		if __obf_ef25cdf0d96b47a8 != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_bee876148a124a99, __obf_ef25cdf0d96b47a8)
		}
		__obf_306c5ad91737b908.
			SetFloat(__obf_a6629621259bc2bc)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_6001932d600cd7ec.Type(), __obf_d6e7451ec5237c6d)
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_f1681715c2229fb2(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_3b1960ca374cf483 := __obf_306c5ad91737b908.Type()
	__obf_7e30add1bc94e75d := __obf_3b1960ca374cf483.Key()
	__obf_edabb506f80bbde6 := __obf_3b1960ca374cf483.Elem()
	__obf_05b9d32ed8d7b2b9 := // By default we overwrite keys in the current map
		__obf_306c5ad91737b908

	// If the map is nil or we're purposely zeroing fields, make a new map
	if __obf_05b9d32ed8d7b2b9.IsNil() || __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ZeroFields {
		__obf_e5263e6aa1d47e9b := // Make a new map to hold our result
			reflect.MapOf(__obf_7e30add1bc94e75d, __obf_edabb506f80bbde6)
		__obf_05b9d32ed8d7b2b9 = reflect.MakeMap(__obf_e5263e6aa1d47e9b)
	}
	__obf_6001932d600cd7ec := // Check input type and based on the input type jump to the proper func
		reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	switch __obf_6001932d600cd7ec.Kind() {
	case reflect.Map:
		return __obf_f16772980031b91a.__obf_445ac64a58a00d8f(__obf_bee876148a124a99, __obf_6001932d600cd7ec, __obf_306c5ad91737b908, __obf_05b9d32ed8d7b2b9)

	case reflect.Struct:
		return __obf_f16772980031b91a.__obf_fb9dce73b9f2e1c0(__obf_6001932d600cd7ec, __obf_306c5ad91737b908, __obf_05b9d32ed8d7b2b9)

	case reflect.Array, reflect.Slice:
		if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput {
			return __obf_f16772980031b91a.__obf_725c3a28f13e62cd(__obf_bee876148a124a99, __obf_6001932d600cd7ec, __obf_306c5ad91737b908, __obf_05b9d32ed8d7b2b9)
		}

		fallthrough

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_bee876148a124a99, __obf_6001932d600cd7ec.Kind())
	}
}

func (__obf_f16772980031b91a *Decoder) __obf_725c3a28f13e62cd(__obf_bee876148a124a99 string, __obf_6001932d600cd7ec reflect.Value, __obf_306c5ad91737b908 reflect.Value, __obf_05b9d32ed8d7b2b9 reflect.Value) error {
	// Special case for BC reasons (covered by tests)
	if __obf_6001932d600cd7ec.Len() == 0 {
		__obf_306c5ad91737b908.
			Set(__obf_05b9d32ed8d7b2b9)
		return nil
	}

	for __obf_a6629621259bc2bc := 0; __obf_a6629621259bc2bc < __obf_6001932d600cd7ec.Len(); __obf_a6629621259bc2bc++ {
		__obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_bee876148a124a99+
			"["+strconv.Itoa(__obf_a6629621259bc2bc)+"]", __obf_6001932d600cd7ec.
			Index(__obf_a6629621259bc2bc).Interface(), __obf_306c5ad91737b908)
		if __obf_ef25cdf0d96b47a8 != nil {
			return __obf_ef25cdf0d96b47a8
		}
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_445ac64a58a00d8f(__obf_bee876148a124a99 string, __obf_6001932d600cd7ec reflect.Value, __obf_306c5ad91737b908 reflect.Value, __obf_05b9d32ed8d7b2b9 reflect.Value) error {
	__obf_3b1960ca374cf483 := __obf_306c5ad91737b908.Type()
	__obf_7e30add1bc94e75d := __obf_3b1960ca374cf483.Key()
	__obf_edabb506f80bbde6 := __obf_3b1960ca374cf483.Elem()

	// Accumulate errors
	errors := make([]string, 0)

	// If the input data is empty, then we just match what the input data is.
	if __obf_6001932d600cd7ec.Len() == 0 {
		if __obf_6001932d600cd7ec.IsNil() {
			if !__obf_306c5ad91737b908.IsNil() {
				__obf_306c5ad91737b908.
					Set(__obf_6001932d600cd7ec)
			}
		} else {
			__obf_306c5ad91737b908.
				// Set to empty allocated value
				Set(__obf_05b9d32ed8d7b2b9)
		}

		return nil
	}

	for _, __obf_838ecb491af7c884 := range __obf_6001932d600cd7ec.MapKeys() {
		__obf_465154ef45471268 := __obf_bee876148a124a99 + "[" + __obf_838ecb491af7c884.String() + "]"
		__obf_222bba34c3d126eb := // First decode the key into the proper type
			reflect.Indirect(reflect.New(__obf_7e30add1bc94e75d))
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_465154ef45471268, __obf_838ecb491af7c884.Interface(), __obf_222bba34c3d126eb); __obf_ef25cdf0d96b47a8 != nil {
			errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
			continue
		}
		__obf_7dd7d18cd783c84c := // Next decode the data into the proper type
			__obf_6001932d600cd7ec.MapIndex(__obf_838ecb491af7c884).Interface()
		__obf_b1dc6e12bbeabeb9 := reflect.Indirect(reflect.New(__obf_edabb506f80bbde6))
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_465154ef45471268, __obf_7dd7d18cd783c84c, __obf_b1dc6e12bbeabeb9); __obf_ef25cdf0d96b47a8 != nil {
			errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
			continue
		}
		__obf_05b9d32ed8d7b2b9.
			SetMapIndex(__obf_222bba34c3d126eb, __obf_b1dc6e12bbeabeb9)
	}
	__obf_306c5ad91737b908.

		// Set the built up map to the value
		Set(__obf_05b9d32ed8d7b2b9)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_fb9dce73b9f2e1c0(__obf_6001932d600cd7ec reflect.Value, __obf_306c5ad91737b908 reflect.Value, __obf_05b9d32ed8d7b2b9 reflect.Value) error {
	__obf_9b5a61c7349f4a82 := __obf_6001932d600cd7ec.Type()
	for __obf_a6629621259bc2bc := 0; __obf_a6629621259bc2bc < __obf_9b5a61c7349f4a82.NumField(); __obf_a6629621259bc2bc++ {
		__obf_9037ed2181de9473 := // Get the StructField first since this is a cheap operation. If the
			// field is unexported, then ignore it.
			__obf_9b5a61c7349f4a82.Field(__obf_a6629621259bc2bc)
		if __obf_9037ed2181de9473.PkgPath != "" {
			continue
		}
		__obf_7dd7d18cd783c84c := // Next get the actual value of this field and verify it is assignable
			// to the map value.
			__obf_6001932d600cd7ec.Field(__obf_a6629621259bc2bc)
		if !__obf_7dd7d18cd783c84c.Type().AssignableTo(__obf_05b9d32ed8d7b2b9.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", __obf_7dd7d18cd783c84c.Type(), __obf_05b9d32ed8d7b2b9.Type().Elem())
		}
		__obf_46b922b23776eb91 := __obf_9037ed2181de9473.Tag.Get(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.TagName)
		__obf_77c915cdae748ece := __obf_9037ed2181de9473.Name

		if __obf_46b922b23776eb91 == "" && __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.IgnoreUntaggedFields {
			continue
		}
		__obf_c862d25e221f0df5 := // If Squash is set in the config, we squash the field down.
			__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Squash && __obf_7dd7d18cd783c84c.Kind() == reflect.Struct && __obf_9037ed2181de9473.Anonymous
		__obf_7dd7d18cd783c84c = __obf_154eba4f55763995(__obf_7dd7d18cd783c84c, __obf_f16772980031b91a.

			// Determine the name of the key in the map
			__obf_85fcbb7e5f1202f3.TagName)

		if __obf_23149fa2b37539b1 := strings.Index(__obf_46b922b23776eb91, ","); __obf_23149fa2b37539b1 != -1 {
			if __obf_46b922b23776eb91[:__obf_23149fa2b37539b1] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Contains(__obf_46b922b23776eb91[__obf_23149fa2b37539b1+1:], "omitempty") && __obf_d91d3de48b9b8cc4(__obf_7dd7d18cd783c84c) {
				continue
			}
			__obf_c862d25e221f0df5 = // If "squash" is specified in the tag, we squash the field down.
				__obf_c862d25e221f0df5 || strings.Contains(__obf_46b922b23776eb91[__obf_23149fa2b37539b1+1:], "squash")
			if __obf_c862d25e221f0df5 {
				// When squashing, the embedded type can be a pointer to a struct.
				if __obf_7dd7d18cd783c84c.Kind() == reflect.Ptr && __obf_7dd7d18cd783c84c.Elem().Kind() == reflect.Struct {
					__obf_7dd7d18cd783c84c = __obf_7dd7d18cd783c84c.Elem()
				}

				// The final type must be a struct
				if __obf_7dd7d18cd783c84c.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", __obf_7dd7d18cd783c84c.Type())
				}
			}
			if __obf_0cde5625a0249089 := __obf_46b922b23776eb91[:__obf_23149fa2b37539b1]; __obf_0cde5625a0249089 != "" {
				__obf_77c915cdae748ece = __obf_0cde5625a0249089
			}
		} else if len(__obf_46b922b23776eb91) > 0 {
			if __obf_46b922b23776eb91 == "-" {
				continue
			}
			__obf_77c915cdae748ece = __obf_46b922b23776eb91
		}

		switch __obf_7dd7d18cd783c84c.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			__obf_f5ef369df0a75c93 := reflect.New(__obf_7dd7d18cd783c84c.Type())
			__obf_f5ef369df0a75c93.
				Elem().Set(__obf_7dd7d18cd783c84c)
			__obf_5f81d6a762682c86 := __obf_05b9d32ed8d7b2b9.Type()
			__obf_add28b4ebd092d8d := __obf_5f81d6a762682c86.Key()
			__obf_d62cb04ed61cadcf := __obf_5f81d6a762682c86.Elem()
			__obf_dde1fb81cbfb43eb := reflect.MapOf(__obf_add28b4ebd092d8d, __obf_d62cb04ed61cadcf)
			__obf_dc75a03f34b98395 := reflect.MakeMap(__obf_dde1fb81cbfb43eb)
			__obf_cf82d0ca15ffa8c6 := // Creating a pointer to a map so that other methods can completely
				// overwrite the map if need be (looking at you decodeMapFromMap). The
				// indirection allows the underlying map to be settable (CanSet() == true)
				// where as reflect.MakeMap returns an unsettable map.
				reflect.New(__obf_dc75a03f34b98395.Type())
			reflect.Indirect(__obf_cf82d0ca15ffa8c6).Set(__obf_dc75a03f34b98395)
			__obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_77c915cdae748ece, __obf_f5ef369df0a75c93.Interface(), reflect.Indirect(__obf_cf82d0ca15ffa8c6))
			if __obf_ef25cdf0d96b47a8 != nil {
				return __obf_ef25cdf0d96b47a8
			}
			__obf_dc75a03f34b98395 = // the underlying map may have been completely overwritten so pull
				// it indirectly out of the enclosing value.
				reflect.Indirect(__obf_cf82d0ca15ffa8c6)

			if __obf_c862d25e221f0df5 {
				for _, __obf_838ecb491af7c884 := range __obf_dc75a03f34b98395.MapKeys() {
					__obf_05b9d32ed8d7b2b9.
						SetMapIndex(__obf_838ecb491af7c884, __obf_dc75a03f34b98395.MapIndex(__obf_838ecb491af7c884))
				}
			} else {
				__obf_05b9d32ed8d7b2b9.
					SetMapIndex(reflect.ValueOf(__obf_77c915cdae748ece), __obf_dc75a03f34b98395)
			}

		default:
			__obf_05b9d32ed8d7b2b9.
				SetMapIndex(reflect.ValueOf(__obf_77c915cdae748ece), __obf_7dd7d18cd783c84c)
		}
	}

	if __obf_306c5ad91737b908.CanAddr() {
		__obf_306c5ad91737b908.
			Set(__obf_05b9d32ed8d7b2b9)
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_3aa25401044d6204(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) (bool, error) {
	__obf_66eeaa4ec80995d9 := // If the input data is nil, then we want to just set the output
		// pointer to be nil as well.
		__obf_d6e7451ec5237c6d == nil
	if !__obf_66eeaa4ec80995d9 {
		switch __obf_7dd7d18cd783c84c := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d)); __obf_7dd7d18cd783c84c.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			__obf_66eeaa4ec80995d9 = __obf_7dd7d18cd783c84c.IsNil()
		}
	}
	if __obf_66eeaa4ec80995d9 {
		if !__obf_306c5ad91737b908.IsNil() && __obf_306c5ad91737b908.CanSet() {
			__obf_b78628dc30284c82 := reflect.New(__obf_306c5ad91737b908.Type()).Elem()
			__obf_306c5ad91737b908.
				Set(__obf_b78628dc30284c82)
		}

		return true, nil
	}
	__obf_3b1960ca374cf483 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		__obf_306c5ad91737b908.Type()
	__obf_edabb506f80bbde6 := __obf_3b1960ca374cf483.Elem()
	if __obf_306c5ad91737b908.CanSet() {
		__obf_4f45012efb5f451d := __obf_306c5ad91737b908
		if __obf_4f45012efb5f451d.IsNil() || __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ZeroFields {
			__obf_4f45012efb5f451d = reflect.New(__obf_edabb506f80bbde6)
		}

		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_bee876148a124a99, __obf_d6e7451ec5237c6d, reflect.Indirect(__obf_4f45012efb5f451d)); __obf_ef25cdf0d96b47a8 != nil {
			return false, __obf_ef25cdf0d96b47a8
		}
		__obf_306c5ad91737b908.
			Set(__obf_4f45012efb5f451d)
	} else {
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_bee876148a124a99, __obf_d6e7451ec5237c6d, reflect.Indirect(__obf_306c5ad91737b908)); __obf_ef25cdf0d96b47a8 != nil {
			return false, __obf_ef25cdf0d96b47a8
		}
	}
	return false, nil
}

func (__obf_f16772980031b91a *Decoder) __obf_5db3819bb8f0f28f(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	if __obf_306c5ad91737b908.Type() != __obf_6001932d600cd7ec.Type() {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_bee876148a124a99, __obf_306c5ad91737b908.
				Type(), __obf_6001932d600cd7ec.Type(), __obf_d6e7451ec5237c6d)
	}
	__obf_306c5ad91737b908.
		Set(__obf_6001932d600cd7ec)
	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_c102bb670129f9e6(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_ebbf06248293d294 := __obf_6001932d600cd7ec.Kind()
	__obf_3b1960ca374cf483 := __obf_306c5ad91737b908.Type()
	__obf_edabb506f80bbde6 := __obf_3b1960ca374cf483.Elem()
	__obf_e60ec3bc56a7ecbe := reflect.SliceOf(__obf_edabb506f80bbde6)

	// If we have a non array/slice type then we first attempt to convert.
	if __obf_ebbf06248293d294 != reflect.Array && __obf_ebbf06248293d294 != reflect.Slice {
		if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case __obf_ebbf06248293d294 == reflect.Slice, __obf_ebbf06248293d294 == reflect.Array:
				break

			// Empty maps turn into empty slices
			case __obf_ebbf06248293d294 == reflect.Map:
				if __obf_6001932d600cd7ec.Len() == 0 {
					__obf_306c5ad91737b908.
						Set(reflect.MakeSlice(__obf_e60ec3bc56a7ecbe, 0, 0))
					return nil
				}
				// Create slice of maps of other sizes
				return __obf_f16772980031b91a.__obf_c102bb670129f9e6(__obf_bee876148a124a99, []any{__obf_d6e7451ec5237c6d}, __obf_306c5ad91737b908)

			case __obf_ebbf06248293d294 == reflect.String && __obf_edabb506f80bbde6.Kind() == reflect.Uint8:
				return __obf_f16772980031b91a.__obf_c102bb670129f9e6(__obf_bee876148a124a99, []byte(__obf_6001932d600cd7ec.String()), __obf_306c5ad91737b908)

			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				return __obf_f16772980031b91a.__obf_c102bb670129f9e6(__obf_bee876148a124a99, []any{__obf_d6e7451ec5237c6d}, __obf_306c5ad91737b908)
			}
		}

		return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", __obf_bee876148a124a99, __obf_ebbf06248293d294)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if __obf_ebbf06248293d294 != reflect.Array && __obf_6001932d600cd7ec.IsNil() {
		return nil
	}
	__obf_b51605cdeedbad29 := __obf_306c5ad91737b908
	if __obf_b51605cdeedbad29.IsNil() || __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ZeroFields {
		__obf_b51605cdeedbad29 = // Make a new slice to hold our result, same size as the original data.
			reflect.MakeSlice(__obf_e60ec3bc56a7ecbe, __obf_6001932d600cd7ec.Len(), __obf_6001932d600cd7ec.Len())
	} else if __obf_b51605cdeedbad29.Len() > __obf_6001932d600cd7ec.Len() {
		__obf_b51605cdeedbad29 = __obf_b51605cdeedbad29.Slice(0, __obf_6001932d600cd7ec.Len())
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_a6629621259bc2bc := 0; __obf_a6629621259bc2bc < __obf_6001932d600cd7ec.Len(); __obf_a6629621259bc2bc++ {
		__obf_c1838709e8af33d5 := __obf_6001932d600cd7ec.Index(__obf_a6629621259bc2bc).Interface()
		for __obf_b51605cdeedbad29.Len() <= __obf_a6629621259bc2bc {
			__obf_b51605cdeedbad29 = reflect.Append(__obf_b51605cdeedbad29, reflect.Zero(__obf_edabb506f80bbde6))
		}
		__obf_7fb0280c081d03e5 := __obf_b51605cdeedbad29.Index(__obf_a6629621259bc2bc)
		__obf_465154ef45471268 := __obf_bee876148a124a99 + "[" + strconv.Itoa(__obf_a6629621259bc2bc) + "]"
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_465154ef45471268, __obf_c1838709e8af33d5, __obf_7fb0280c081d03e5); __obf_ef25cdf0d96b47a8 != nil {
			errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
		}
	}
	__obf_306c5ad91737b908.

		// Finally, set the value to the slice we built up
		Set(__obf_b51605cdeedbad29)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_978e1dbf26790dcd(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))
	__obf_ebbf06248293d294 := __obf_6001932d600cd7ec.Kind()
	__obf_3b1960ca374cf483 := __obf_306c5ad91737b908.Type()
	__obf_edabb506f80bbde6 := __obf_3b1960ca374cf483.Elem()
	__obf_d7bdcbc5a36a4482 := reflect.ArrayOf(__obf_3b1960ca374cf483.Len(), __obf_edabb506f80bbde6)
	__obf_a5f768b9a5a60bf0 := __obf_306c5ad91737b908

	if __obf_a5f768b9a5a60bf0.Interface() == reflect.Zero(__obf_a5f768b9a5a60bf0.Type()).Interface() || __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ZeroFields {
		// Check input type
		if __obf_ebbf06248293d294 != reflect.Array && __obf_ebbf06248293d294 != reflect.Slice {
			if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.WeaklyTypedInput {
				switch __obf_ebbf06248293d294 {
				// Empty maps turn into empty arrays
				case reflect.Map:
					if __obf_6001932d600cd7ec.Len() == 0 {
						__obf_306c5ad91737b908.
							Set(reflect.Zero(__obf_d7bdcbc5a36a4482))
						return nil
					}

				// All other types we try to convert to the array type
				// and "lift" it into it. i.e. a string becomes a string array.
				default:
					// Just re-try this function with data as a slice.
					return __obf_f16772980031b91a.__obf_978e1dbf26790dcd(__obf_bee876148a124a99, []any{__obf_d6e7451ec5237c6d}, __obf_306c5ad91737b908)
				}
			}

			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", __obf_bee876148a124a99, __obf_ebbf06248293d294)

		}
		if __obf_6001932d600cd7ec.Len() > __obf_d7bdcbc5a36a4482.Len() {
			return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", __obf_bee876148a124a99, __obf_d7bdcbc5a36a4482.Len(), __obf_6001932d600cd7ec.Len())

		}
		__obf_a5f768b9a5a60bf0 = // Make a new array to hold our result, same size as the original data.
			reflect.New(__obf_d7bdcbc5a36a4482).Elem()
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_a6629621259bc2bc := 0; __obf_a6629621259bc2bc < __obf_6001932d600cd7ec.Len(); __obf_a6629621259bc2bc++ {
		__obf_c1838709e8af33d5 := __obf_6001932d600cd7ec.Index(__obf_a6629621259bc2bc).Interface()
		__obf_7fb0280c081d03e5 := __obf_a5f768b9a5a60bf0.Index(__obf_a6629621259bc2bc)
		__obf_465154ef45471268 := __obf_bee876148a124a99 + "[" + strconv.Itoa(__obf_a6629621259bc2bc) + "]"
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_465154ef45471268, __obf_c1838709e8af33d5, __obf_7fb0280c081d03e5); __obf_ef25cdf0d96b47a8 != nil {
			errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
		}
	}
	__obf_306c5ad91737b908.

		// Finally, set the value to the array we built up
		Set(__obf_a5f768b9a5a60bf0)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_f16772980031b91a *Decoder) __obf_195c05f649453920(__obf_bee876148a124a99 string, __obf_d6e7451ec5237c6d any, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_6001932d600cd7ec := reflect.Indirect(reflect.ValueOf(__obf_d6e7451ec5237c6d))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if __obf_6001932d600cd7ec.Type() == __obf_306c5ad91737b908.Type() {
		__obf_306c5ad91737b908.
			Set(__obf_6001932d600cd7ec)
		return nil
	}
	__obf_ebbf06248293d294 := __obf_6001932d600cd7ec.Kind()
	switch __obf_ebbf06248293d294 {
	case reflect.Map:
		return __obf_f16772980031b91a.__obf_34c6ff0847f9718b(__obf_bee876148a124a99, __obf_6001932d600cd7ec,

			// Not the most efficient way to do this but we can optimize later if
			// we want to. To convert from struct to struct we go to map first
			// as an intermediary.
			__obf_306c5ad91737b908)

	case reflect.Struct:
		__obf_e5263e6aa1d47e9b := // Make a new map to hold our result
			reflect.TypeOf((map[string]any)(nil))
		__obf_7ccf46266e1e3329 := reflect.MakeMap(__obf_e5263e6aa1d47e9b)
		__obf_cf82d0ca15ffa8c6 := // Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			reflect.New(__obf_7ccf46266e1e3329.Type())

		reflect.Indirect(__obf_cf82d0ca15ffa8c6).Set(__obf_7ccf46266e1e3329)
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_fb9dce73b9f2e1c0(__obf_6001932d600cd7ec, reflect.Indirect(__obf_cf82d0ca15ffa8c6), __obf_7ccf46266e1e3329); __obf_ef25cdf0d96b47a8 != nil {
			return __obf_ef25cdf0d96b47a8
		}
		__obf_84b0717583bd55a4 := __obf_f16772980031b91a.__obf_34c6ff0847f9718b(__obf_bee876148a124a99, reflect.Indirect(__obf_cf82d0ca15ffa8c6), __obf_306c5ad91737b908)
		return __obf_84b0717583bd55a4

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_bee876148a124a99, __obf_6001932d600cd7ec.Kind())
	}
}

func (__obf_f16772980031b91a *Decoder) __obf_34c6ff0847f9718b(__obf_bee876148a124a99 string, __obf_6001932d600cd7ec, __obf_306c5ad91737b908 reflect.Value) error {
	__obf_40f50f46d0d164df := __obf_6001932d600cd7ec.Type()
	if __obf_4ead81dbe29048b5 := __obf_40f50f46d0d164df.Key().Kind(); __obf_4ead81dbe29048b5 != reflect.String && __obf_4ead81dbe29048b5 != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys", __obf_bee876148a124a99, __obf_40f50f46d0d164df.
				Key().Kind())
	}
	__obf_28b81627c50b740e := make(map[reflect.Value]struct{})
	__obf_e5b98008e5d6c01a := make(map[any]struct{})
	for _, __obf_72f206bc61dbd854 := range __obf_6001932d600cd7ec.MapKeys() {
		__obf_28b81627c50b740e[__obf_72f206bc61dbd854] = struct{}{}
		__obf_e5b98008e5d6c01a[__obf_72f206bc61dbd854.Interface()] = struct{}{}
	}
	__obf_199e05aaf9b5fb14 := make(map[any]struct{})
	errors := make([]string, 0)
	__obf_f15ec178202ca641 := // This slice will keep track of all the structs we'll be decoding.
		// There can be more than one struct if there are embedded structs
		// that are squashed.
		make([]reflect.Value, 1, 5)
	__obf_f15ec178202ca641[0] = __obf_306c5ad91737b908

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type __obf_1b89b504cb8578ce struct {
		__obf_1b89b504cb8578ce reflect.StructField
		__obf_306c5ad91737b908 reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var __obf_d07eed2f6ff0c92e *__obf_1b89b504cb8578ce
	__obf_eb4611098d3f5537 := []__obf_1b89b504cb8578ce{}
	for len(__obf_f15ec178202ca641) > 0 {
		__obf_a8c665105f213a1b := __obf_f15ec178202ca641[0]
		__obf_f15ec178202ca641 = __obf_f15ec178202ca641[1:]
		__obf_02c179a80230ff8d := __obf_a8c665105f213a1b.Type()

		for __obf_a6629621259bc2bc := 0; __obf_a6629621259bc2bc < __obf_02c179a80230ff8d.NumField(); __obf_a6629621259bc2bc++ {
			__obf_1882933cea3860c5 := __obf_02c179a80230ff8d.Field(__obf_a6629621259bc2bc)
			__obf_782b41b018a033ad := __obf_a8c665105f213a1b.Field(__obf_a6629621259bc2bc)
			if __obf_782b41b018a033ad.Kind() == reflect.Pointer && __obf_782b41b018a033ad.Elem().Kind() == reflect.Struct {
				__obf_782b41b018a033ad = // Handle embedded struct pointers as embedded structs.
					__obf_782b41b018a033ad.Elem()
			}
			__obf_c862d25e221f0df5 := // If "squash" is specified in the tag, we squash the field down.
				__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Squash && __obf_782b41b018a033ad.Kind() == reflect.Struct && __obf_1882933cea3860c5.Anonymous
			__obf_538c43ae6b59e5c8 := false
			__obf_989a443537596a4b := // We always parse the tags cause we're looking for other tags too
				strings.Split(__obf_1882933cea3860c5.Tag.Get(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.TagName), ",")
			for _, __obf_666c45c08e09a0b6 := range __obf_989a443537596a4b[1:] {
				if __obf_666c45c08e09a0b6 == "squash" {
					__obf_c862d25e221f0df5 = true
					break
				}

				if __obf_666c45c08e09a0b6 == "remain" {
					__obf_538c43ae6b59e5c8 = true
					break
				}
			}

			if __obf_c862d25e221f0df5 {
				if __obf_782b41b018a033ad.Kind() != reflect.Struct {
					errors = __obf_39ea766e30fe264e(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", __obf_1882933cea3860c5.Name, __obf_782b41b018a033ad.Kind()))
				} else {
					__obf_f15ec178202ca641 = append(__obf_f15ec178202ca641, __obf_782b41b018a033ad)
				}
				continue
			}

			// Build our field
			if __obf_538c43ae6b59e5c8 {
				__obf_d07eed2f6ff0c92e = &__obf_1b89b504cb8578ce{__obf_1882933cea3860c5, __obf_782b41b018a033ad}
			} else {
				__obf_eb4611098d3f5537 = // Normal struct field, store it away
					append(__obf_eb4611098d3f5537, __obf_1b89b504cb8578ce{__obf_1882933cea3860c5, __obf_782b41b018a033ad})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, __obf_9037ed2181de9473 := range __obf_eb4611098d3f5537 {
		__obf_1b89b504cb8578ce, __obf_48c67fdcea42ea0f := __obf_9037ed2181de9473.__obf_1b89b504cb8578ce, __obf_9037ed2181de9473.__obf_306c5ad91737b908
		__obf_465154ef45471268 := __obf_1b89b504cb8578ce.Name
		__obf_46b922b23776eb91 := __obf_1b89b504cb8578ce.Tag.Get(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.TagName)
		__obf_46b922b23776eb91 = strings.SplitN(__obf_46b922b23776eb91, ",", 2)[0]
		if __obf_46b922b23776eb91 != "" {
			__obf_465154ef45471268 = __obf_46b922b23776eb91
		}
		__obf_ca284b1b24df7546 := reflect.ValueOf(__obf_465154ef45471268)
		__obf_43b22450d0648c13 := __obf_6001932d600cd7ec.MapIndex(__obf_ca284b1b24df7546)
		if !__obf_43b22450d0648c13.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for __obf_72f206bc61dbd854 := range __obf_28b81627c50b740e {
				__obf_c1e4cac3ee94b366, __obf_b45c7b8adb1c2d97 := __obf_72f206bc61dbd854.Interface().(string)
				if !__obf_b45c7b8adb1c2d97 {
					// Not a string key
					continue
				}

				if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.MatchName(__obf_c1e4cac3ee94b366, __obf_465154ef45471268) {
					__obf_ca284b1b24df7546 = __obf_72f206bc61dbd854
					__obf_43b22450d0648c13 = __obf_6001932d600cd7ec.MapIndex(__obf_72f206bc61dbd854)
					break
				}
			}

			if !__obf_43b22450d0648c13.IsValid() {
				__obf_199e05aaf9b5fb14[ // There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				__obf_465154ef45471268] = struct{}{}
				continue
			}
		}

		if !__obf_48c67fdcea42ea0f.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !__obf_48c67fdcea42ea0f.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(__obf_e5b98008e5d6c01a, __obf_ca284b1b24df7546.Interface())

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if __obf_bee876148a124a99 != "" {
			__obf_465154ef45471268 = __obf_bee876148a124a99 + "." + __obf_465154ef45471268
		}

		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_26d8f56953be1262(__obf_465154ef45471268, __obf_43b22450d0648c13.Interface(), __obf_48c67fdcea42ea0f); __obf_ef25cdf0d96b47a8 != nil {
			errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if __obf_d07eed2f6ff0c92e != nil && len(__obf_e5b98008e5d6c01a) > 0 {
		__obf_538c43ae6b59e5c8 := // Build a map of only the unused values
			map[any]any{}
		for __obf_0247a888fdcb96a7 := range __obf_e5b98008e5d6c01a {
			__obf_538c43ae6b59e5c8[__obf_0247a888fdcb96a7] = __obf_6001932d600cd7ec.MapIndex(reflect.ValueOf(__obf_0247a888fdcb96a7)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if __obf_ef25cdf0d96b47a8 := __obf_f16772980031b91a.__obf_f1681715c2229fb2(__obf_bee876148a124a99, __obf_538c43ae6b59e5c8, __obf_d07eed2f6ff0c92e.__obf_306c5ad91737b908); __obf_ef25cdf0d96b47a8 != nil {
			errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
		}
		__obf_e5b98008e5d6c01a = // Set the map to nil so we have none so that the next check will
			// not error (ErrorUnused)
			nil
	}

	if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ErrorUnused && len(__obf_e5b98008e5d6c01a) > 0 {
		__obf_eb9fda10511d860f := make([]string, 0, len(__obf_e5b98008e5d6c01a))
		for __obf_d320e23c5f3d8925 := range __obf_e5b98008e5d6c01a {
			__obf_eb9fda10511d860f = append(__obf_eb9fda10511d860f, __obf_d320e23c5f3d8925.(string))
		}
		sort.Strings(__obf_eb9fda10511d860f)
		__obf_ef25cdf0d96b47a8 := fmt.Errorf("'%s' has invalid keys: %s", __obf_bee876148a124a99, strings.Join(__obf_eb9fda10511d860f, ", "))
		errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
	}

	if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.ErrorUnset && len(__obf_199e05aaf9b5fb14) > 0 {
		__obf_eb9fda10511d860f := make([]string, 0, len(__obf_199e05aaf9b5fb14))
		for __obf_d320e23c5f3d8925 := range __obf_199e05aaf9b5fb14 {
			__obf_eb9fda10511d860f = append(__obf_eb9fda10511d860f, __obf_d320e23c5f3d8925.(string))
		}
		sort.Strings(__obf_eb9fda10511d860f)
		__obf_ef25cdf0d96b47a8 := fmt.Errorf("'%s' has unset fields: %s", __obf_bee876148a124a99, strings.Join(__obf_eb9fda10511d860f, ", "))
		errors = __obf_39ea766e30fe264e(errors, __obf_ef25cdf0d96b47a8)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	if __obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata != nil {
		for __obf_d320e23c5f3d8925 := range __obf_e5b98008e5d6c01a {
			__obf_0247a888fdcb96a7 := __obf_d320e23c5f3d8925.(string)
			if __obf_bee876148a124a99 != "" {
				__obf_0247a888fdcb96a7 = __obf_bee876148a124a99 + "." + __obf_0247a888fdcb96a7
			}
			__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.
				Metadata.Unused = append(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata.Unused, __obf_0247a888fdcb96a7)
		}
		for __obf_d320e23c5f3d8925 := range __obf_199e05aaf9b5fb14 {
			__obf_0247a888fdcb96a7 := __obf_d320e23c5f3d8925.(string)
			if __obf_bee876148a124a99 != "" {
				__obf_0247a888fdcb96a7 = __obf_bee876148a124a99 + "." + __obf_0247a888fdcb96a7
			}
			__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.
				Metadata.Unset = append(__obf_f16772980031b91a.__obf_85fcbb7e5f1202f3.Metadata.Unset, __obf_0247a888fdcb96a7)
		}
	}

	return nil
}

func __obf_d91d3de48b9b8cc4(__obf_7dd7d18cd783c84c reflect.Value) bool {
	switch __obf_c3efe5973b2597dc(__obf_7dd7d18cd783c84c) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_7dd7d18cd783c84c.Len() == 0
	case reflect.Bool:
		return !__obf_7dd7d18cd783c84c.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_7dd7d18cd783c84c.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_7dd7d18cd783c84c.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_7dd7d18cd783c84c.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return __obf_7dd7d18cd783c84c.IsNil()
	}
	return false
}

func __obf_c3efe5973b2597dc(__obf_306c5ad91737b908 reflect.Value) reflect.Kind {
	__obf_4ead81dbe29048b5 := __obf_306c5ad91737b908.Kind()

	switch {
	case __obf_4ead81dbe29048b5 >= reflect.Int && __obf_4ead81dbe29048b5 <= reflect.Int64:
		return reflect.Int
	case __obf_4ead81dbe29048b5 >= reflect.Uint && __obf_4ead81dbe29048b5 <= reflect.Uint64:
		return reflect.Uint
	case __obf_4ead81dbe29048b5 >= reflect.Float32 && __obf_4ead81dbe29048b5 <= reflect.Float64:
		return reflect.Float32
	default:
		return __obf_4ead81dbe29048b5
	}
}

func __obf_91144e344ea2df31(__obf_9b5a61c7349f4a82 reflect.Type, __obf_1254101d2f73a9ac bool, __obf_72f1c20f1f7a9652 string) bool {
	for __obf_a6629621259bc2bc := 0; __obf_a6629621259bc2bc < __obf_9b5a61c7349f4a82.NumField(); __obf_a6629621259bc2bc++ {
		__obf_9037ed2181de9473 := __obf_9b5a61c7349f4a82.Field(__obf_a6629621259bc2bc)
		if __obf_9037ed2181de9473.PkgPath == "" && !__obf_1254101d2f73a9ac { // check for unexported fields
			return true
		}
		if __obf_1254101d2f73a9ac && __obf_9037ed2181de9473.Tag.Get(__obf_72f1c20f1f7a9652) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func __obf_154eba4f55763995(__obf_7dd7d18cd783c84c reflect.Value, __obf_72f1c20f1f7a9652 string) reflect.Value {
	if __obf_7dd7d18cd783c84c.Kind() != reflect.Ptr || __obf_7dd7d18cd783c84c.Elem().Kind() != reflect.Struct {
		return __obf_7dd7d18cd783c84c
	}
	__obf_f9ffb0182b2339f3 := __obf_7dd7d18cd783c84c.Elem()
	__obf_17853f62c84da088 := __obf_f9ffb0182b2339f3.Type()
	if __obf_91144e344ea2df31(__obf_17853f62c84da088, true, __obf_72f1c20f1f7a9652) {
		return __obf_f9ffb0182b2339f3
	}
	return __obf_7dd7d18cd783c84c
}
