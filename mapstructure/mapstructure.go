package __obf_4dc3483102e0d35a

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
type DecodeHookFuncValue func(__obf_73039542a851b768 reflect.Value, __obf_f41e135eac34c3c2 reflect.Value) (any, error)

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
	MatchName func(__obf_0e4366fda7989295, __obf_7dc427653c0378b5 string) bool
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	__obf_9cbc86dfb44f54f3 *DecoderConfig
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
func Decode(__obf_4d9c12d37915eb62 any, __obf_232de58736b63a83 any) error {
	__obf_9cbc86dfb44f54f3 := &DecoderConfig{
		Metadata: nil,
		Result:   __obf_232de58736b63a83,
	}
	__obf_87da35307221dfe4, __obf_6585aef4313e6005 := NewDecoder(__obf_9cbc86dfb44f54f3)
	if __obf_6585aef4313e6005 != nil {
		return __obf_6585aef4313e6005
	}

	return __obf_87da35307221dfe4.Decode(__obf_4d9c12d37915eb62)
}

// WeakDecode is the same as Decode but is shorthand to enable
// WeaklyTypedInput. See DecoderConfig for more info.
func WeakDecode(__obf_4d9c12d37915eb62, __obf_232de58736b63a83 any) error {
	__obf_9cbc86dfb44f54f3 := &DecoderConfig{
		Metadata:         nil,
		Result:           __obf_232de58736b63a83,
		WeaklyTypedInput: true,
	}
	__obf_87da35307221dfe4, __obf_6585aef4313e6005 := NewDecoder(__obf_9cbc86dfb44f54f3)
	if __obf_6585aef4313e6005 != nil {
		return __obf_6585aef4313e6005
	}

	return __obf_87da35307221dfe4.Decode(__obf_4d9c12d37915eb62)
}

// DecodeMetadata is the same as Decode, but is shorthand to
// enable metadata collection. See DecoderConfig for more info.
func DecodeMetadata(__obf_4d9c12d37915eb62 any, __obf_232de58736b63a83 any, __obf_7a6e81a80e21d471 *Metadata) error {
	__obf_9cbc86dfb44f54f3 := &DecoderConfig{
		Metadata: __obf_7a6e81a80e21d471,
		Result:   __obf_232de58736b63a83,
	}
	__obf_87da35307221dfe4, __obf_6585aef4313e6005 := NewDecoder(__obf_9cbc86dfb44f54f3)
	if __obf_6585aef4313e6005 != nil {
		return __obf_6585aef4313e6005
	}

	return __obf_87da35307221dfe4.Decode(__obf_4d9c12d37915eb62)
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
// enable both WeaklyTypedInput and metadata collection. See
// DecoderConfig for more info.
func WeakDecodeMetadata(__obf_4d9c12d37915eb62 any, __obf_232de58736b63a83 any, __obf_7a6e81a80e21d471 *Metadata) error {
	__obf_9cbc86dfb44f54f3 := &DecoderConfig{
		Metadata:         __obf_7a6e81a80e21d471,
		Result:           __obf_232de58736b63a83,
		WeaklyTypedInput: true,
	}
	__obf_87da35307221dfe4, __obf_6585aef4313e6005 := NewDecoder(__obf_9cbc86dfb44f54f3)
	if __obf_6585aef4313e6005 != nil {
		return __obf_6585aef4313e6005
	}

	return __obf_87da35307221dfe4.Decode(__obf_4d9c12d37915eb62)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(__obf_9cbc86dfb44f54f3 *DecoderConfig) (*Decoder, error) {
	__obf_b422869c3c403707 := reflect.ValueOf(__obf_9cbc86dfb44f54f3.Result)
	if __obf_b422869c3c403707.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}
	__obf_b422869c3c403707 = __obf_b422869c3c403707.Elem()
	if !__obf_b422869c3c403707.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if __obf_9cbc86dfb44f54f3.Metadata != nil {
		if __obf_9cbc86dfb44f54f3.Metadata.Keys == nil {
			__obf_9cbc86dfb44f54f3.
				Metadata.Keys = make([]string, 0)
		}

		if __obf_9cbc86dfb44f54f3.Metadata.Unused == nil {
			__obf_9cbc86dfb44f54f3.
				Metadata.Unused = make([]string, 0)
		}

		if __obf_9cbc86dfb44f54f3.Metadata.Unset == nil {
			__obf_9cbc86dfb44f54f3.
				Metadata.Unset = make([]string, 0)
		}
	}

	if __obf_9cbc86dfb44f54f3.TagName == "" {
		__obf_9cbc86dfb44f54f3.
			TagName = "mapstructure"
	}

	if __obf_9cbc86dfb44f54f3.MatchName == nil {
		__obf_9cbc86dfb44f54f3.
			MatchName = strings.EqualFold
	}
	__obf_382efc6537d96d52 := &Decoder{__obf_9cbc86dfb44f54f3: __obf_9cbc86dfb44f54f3}

	return __obf_382efc6537d96d52, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (__obf_e47c7e0847f6ec2f *Decoder) Decode(__obf_4d9c12d37915eb62 any) error {
	return __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4("", __obf_4d9c12d37915eb62, reflect.ValueOf(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (__obf_e47c7e0847f6ec2f *Decoder) __obf_f7c6980d711168d4(__obf_255580bc58cf586d string, __obf_4d9c12d37915eb62 any, __obf_292b1f30ecd72b8b reflect.Value) error {
	var __obf_e2a4b9fc8d88b6d7 reflect.Value
	if __obf_4d9c12d37915eb62 != nil {
		__obf_e2a4b9fc8d88b6d7 = reflect.ValueOf(__obf_4d9c12d37915eb62)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if __obf_e2a4b9fc8d88b6d7.Kind() == reflect.Ptr && __obf_e2a4b9fc8d88b6d7.IsNil() {
			__obf_4d9c12d37915eb62 = nil
		}
	}

	if __obf_4d9c12d37915eb62 == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ZeroFields {
			__obf_292b1f30ecd72b8b.
				Set(reflect.Zero(__obf_292b1f30ecd72b8b.Type()))

			if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata != nil && __obf_255580bc58cf586d != "" {
				__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.
					Metadata.Keys = append(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata.Keys, __obf_255580bc58cf586d)
			}
		}
		return nil
	}

	if !__obf_e2a4b9fc8d88b6d7.IsValid() {
		__obf_292b1f30ecd72b8b.
			// If the input value is invalid, then we just set the value
			// to be the zero value.
			Set(reflect.Zero(__obf_292b1f30ecd72b8b.Type()))
		if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata != nil && __obf_255580bc58cf586d != "" {
			__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.
				Metadata.Keys = append(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata.Keys, __obf_255580bc58cf586d)
		}
		return nil
	}

	if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the input.
		var __obf_6585aef4313e6005 error
		__obf_4d9c12d37915eb62, __obf_6585aef4313e6005 = DecodeHookExec(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.DecodeHook, __obf_e2a4b9fc8d88b6d7, __obf_292b1f30ecd72b8b)
		if __obf_6585aef4313e6005 != nil {
			return fmt.Errorf("error decoding '%s': %w", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
	}

	var __obf_6585aef4313e6005 error
	__obf_453a19697d1a7aeb := __obf_2d4c05a1f6a0cc52(__obf_292b1f30ecd72b8b)
	__obf_5a5743e19dd535f9 := true
	switch __obf_453a19697d1a7aeb {
	case reflect.Bool:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_510fd3ba9d14ce71(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Interface:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_b924fb75edfd8142(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.String:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_66afe0d51bc4a44a(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Int:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_aa7ae73ac8306ef3(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Uint:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_21eef3aa233920c7(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Float32:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_051f2692931eb0df(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Struct:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_70a1673aa1d5007b(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Map:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_34e83c4c958ecc46(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Ptr:
		__obf_5a5743e19dd535f9, __obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_b2031a7d08ca7bbc(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Slice:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_6cc4d6625c3961b3(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Array:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_d961bcc337062438(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62, __obf_292b1f30ecd72b8b)
	case reflect.Func:
		__obf_6585aef4313e6005 = __obf_e47c7e0847f6ec2f.__obf_4c760466d764236a(__obf_255580bc58cf586d, __obf_4d9c12d37915eb62,

			// If we reached this point then we weren't able to decode it
			__obf_292b1f30ecd72b8b)
	default:

		return fmt.Errorf("%s: unsupported type: %s", __obf_255580bc58cf586d, __obf_453a19697d1a7aeb)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metainput.
	if __obf_5a5743e19dd535f9 && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata != nil && __obf_255580bc58cf586d != "" {
		__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.
			Metadata.Keys = append(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata.Keys, __obf_255580bc58cf586d)
	}

	return __obf_6585aef4313e6005
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (__obf_e47c7e0847f6ec2f *Decoder) __obf_b924fb75edfd8142(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	if __obf_b422869c3c403707.IsValid() && __obf_b422869c3c403707.Elem().IsValid() {
		__obf_65309c2aa575550c := __obf_b422869c3c403707.Elem()
		__obf_5445eec759e92217 := // If we can't address this element, then its not writable. Instead,
			// we make a copy of the value (which is a pointer and therefore
			// writable), decode into that, and replace the whole value.
			false
		if !__obf_65309c2aa575550c.CanAddr() {
			__obf_5445eec759e92217 = true

			// Make *T
			copy := reflect.New(__obf_65309c2aa575550c.Type())

			// *T = elem
			copy.Elem().Set(__obf_65309c2aa575550c)
			__obf_65309c2aa575550c = // Set elem so we decode into it
				copy
		}

		// Decode. If we have an error then return. We also return right
		// away if we're not a copy because that means we decoded directly.
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_255580bc58cf586d, __obf_333b2acf09d806b3, __obf_65309c2aa575550c); __obf_6585aef4313e6005 != nil || !__obf_5445eec759e92217 {
			return __obf_6585aef4313e6005
		}
		__obf_b422869c3c403707.

			// If we're a copy, we need to set te final result
			Set(__obf_65309c2aa575550c.Elem())
		return nil
	}
	__obf_a4fef5c9586c1966 := reflect.ValueOf(__obf_333b2acf09d806b3)

	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if __obf_a4fef5c9586c1966.Kind() == reflect.Ptr && __obf_a4fef5c9586c1966.Type().Elem() == __obf_b422869c3c403707.Type() {
		__obf_a4fef5c9586c1966 = reflect.Indirect(__obf_a4fef5c9586c1966)
	}

	if !__obf_a4fef5c9586c1966.IsValid() {
		__obf_a4fef5c9586c1966 = reflect.Zero(__obf_b422869c3c403707.Type())
	}
	__obf_796baf2c45c67595 := __obf_a4fef5c9586c1966.Type()
	if !__obf_796baf2c45c67595.AssignableTo(__obf_b422869c3c403707.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_796baf2c45c67595)
	}
	__obf_b422869c3c403707.
		Set(__obf_a4fef5c9586c1966)
	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_66afe0d51bc4a44a(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_7a924d47297a474f := __obf_2d4c05a1f6a0cc52(__obf_a4fef5c9586c1966)
	__obf_ca5e382d2b75cd20 := true
	switch {
	case __obf_7a924d47297a474f == reflect.String:
		__obf_b422869c3c403707.
			SetString(__obf_a4fef5c9586c1966.String())
	case __obf_7a924d47297a474f == reflect.Bool && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		if __obf_a4fef5c9586c1966.Bool() {
			__obf_b422869c3c403707.
				SetString("1")
		} else {
			__obf_b422869c3c403707.
				SetString("0")
		}
	case __obf_7a924d47297a474f == reflect.Int && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_b422869c3c403707.
			SetString(strconv.FormatInt(__obf_a4fef5c9586c1966.Int(), 10))
	case __obf_7a924d47297a474f == reflect.Uint && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_b422869c3c403707.
			SetString(strconv.FormatUint(__obf_a4fef5c9586c1966.Uint(), 10))
	case __obf_7a924d47297a474f == reflect.Float32 && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_b422869c3c403707.
			SetString(strconv.FormatFloat(__obf_a4fef5c9586c1966.Float(), 'f', -1, 64))
	case __obf_7a924d47297a474f == reflect.Slice && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput, __obf_7a924d47297a474f ==
		reflect.Array && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_f73efdd6543886c9 := __obf_a4fef5c9586c1966.Type()
		__obf_9bbf49c8c89fbd8a := __obf_f73efdd6543886c9.Elem().Kind()
		switch __obf_9bbf49c8c89fbd8a {
		case reflect.Uint8:
			var __obf_070d49286e69fa1c []uint8
			if __obf_7a924d47297a474f == reflect.Array {
				__obf_070d49286e69fa1c = make([]uint8, __obf_a4fef5c9586c1966.Len())
				for __obf_99f0e7230007b81b := range __obf_070d49286e69fa1c {
					__obf_070d49286e69fa1c[__obf_99f0e7230007b81b] = __obf_a4fef5c9586c1966.Index(__obf_99f0e7230007b81b).Interface().(uint8)
				}
			} else {
				__obf_070d49286e69fa1c = __obf_a4fef5c9586c1966.Interface().([]uint8)
			}
			__obf_b422869c3c403707.
				SetString(string(__obf_070d49286e69fa1c))
		default:
			__obf_ca5e382d2b75cd20 = false
		}
	default:
		__obf_ca5e382d2b75cd20 = false
	}

	if !__obf_ca5e382d2b75cd20 {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_a4fef5c9586c1966.Type(), __obf_333b2acf09d806b3)
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_aa7ae73ac8306ef3(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_7a924d47297a474f := __obf_2d4c05a1f6a0cc52(__obf_a4fef5c9586c1966)
	__obf_f73efdd6543886c9 := __obf_a4fef5c9586c1966.Type()

	switch {
	case __obf_7a924d47297a474f == reflect.Int:
		__obf_b422869c3c403707.
			SetInt(__obf_a4fef5c9586c1966.Int())
	case __obf_7a924d47297a474f == reflect.Uint:
		__obf_b422869c3c403707.
			SetInt(int64(__obf_a4fef5c9586c1966.Uint()))
	case __obf_7a924d47297a474f == reflect.Float32:
		__obf_b422869c3c403707.
			SetInt(int64(__obf_a4fef5c9586c1966.Float()))
	case __obf_7a924d47297a474f == reflect.Bool && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		if __obf_a4fef5c9586c1966.Bool() {
			__obf_b422869c3c403707.
				SetInt(1)
		} else {
			__obf_b422869c3c403707.
				SetInt(0)
		}
	case __obf_7a924d47297a474f == reflect.String && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_98acb43641f668a3 := __obf_a4fef5c9586c1966.String()
		if __obf_98acb43641f668a3 == "" {
			__obf_98acb43641f668a3 = "0"
		}
		__obf_99f0e7230007b81b, __obf_6585aef4313e6005 := strconv.ParseInt(__obf_98acb43641f668a3, 0, __obf_b422869c3c403707.Type().Bits())
		if __obf_6585aef4313e6005 == nil {
			__obf_b422869c3c403707.
				SetInt(__obf_99f0e7230007b81b)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
	case __obf_f73efdd6543886c9.PkgPath() == "encoding/json" && __obf_f73efdd6543886c9.Name() == "Number":
		__obf_20127d5b36bcdbc4 := __obf_333b2acf09d806b3.(json.Number)
		__obf_99f0e7230007b81b, __obf_6585aef4313e6005 := __obf_20127d5b36bcdbc4.Int64()
		if __obf_6585aef4313e6005 != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
		__obf_b422869c3c403707.
			SetInt(__obf_99f0e7230007b81b)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_a4fef5c9586c1966.Type(), __obf_333b2acf09d806b3)
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_21eef3aa233920c7(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_7a924d47297a474f := __obf_2d4c05a1f6a0cc52(__obf_a4fef5c9586c1966)
	__obf_f73efdd6543886c9 := __obf_a4fef5c9586c1966.Type()

	switch {
	case __obf_7a924d47297a474f == reflect.Int:
		__obf_99f0e7230007b81b := __obf_a4fef5c9586c1966.Int()
		if __obf_99f0e7230007b81b < 0 && !__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %d overflows uint", __obf_255580bc58cf586d, __obf_99f0e7230007b81b)
		}
		__obf_b422869c3c403707.
			SetUint(uint64(__obf_99f0e7230007b81b))
	case __obf_7a924d47297a474f == reflect.Uint:
		__obf_b422869c3c403707.
			SetUint(__obf_a4fef5c9586c1966.Uint())
	case __obf_7a924d47297a474f == reflect.Float32:
		__obf_4cd911a2d7099b07 := __obf_a4fef5c9586c1966.Float()
		if __obf_4cd911a2d7099b07 < 0 && !__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %f overflows uint", __obf_255580bc58cf586d, __obf_4cd911a2d7099b07)
		}
		__obf_b422869c3c403707.
			SetUint(uint64(__obf_4cd911a2d7099b07))
	case __obf_7a924d47297a474f == reflect.Bool && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		if __obf_a4fef5c9586c1966.Bool() {
			__obf_b422869c3c403707.
				SetUint(1)
		} else {
			__obf_b422869c3c403707.
				SetUint(0)
		}
	case __obf_7a924d47297a474f == reflect.String && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_98acb43641f668a3 := __obf_a4fef5c9586c1966.String()
		if __obf_98acb43641f668a3 == "" {
			__obf_98acb43641f668a3 = "0"
		}
		__obf_99f0e7230007b81b, __obf_6585aef4313e6005 := strconv.ParseUint(__obf_98acb43641f668a3, 0, __obf_b422869c3c403707.Type().Bits())
		if __obf_6585aef4313e6005 == nil {
			__obf_b422869c3c403707.
				SetUint(__obf_99f0e7230007b81b)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
	case __obf_f73efdd6543886c9.PkgPath() == "encoding/json" && __obf_f73efdd6543886c9.Name() == "Number":
		__obf_20127d5b36bcdbc4 := __obf_333b2acf09d806b3.(json.Number)
		__obf_99f0e7230007b81b, __obf_6585aef4313e6005 := strconv.ParseUint(string(__obf_20127d5b36bcdbc4), 0, 64)
		if __obf_6585aef4313e6005 != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
		__obf_b422869c3c403707.
			SetUint(__obf_99f0e7230007b81b)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_a4fef5c9586c1966.Type(), __obf_333b2acf09d806b3)
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_510fd3ba9d14ce71(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_7a924d47297a474f := __obf_2d4c05a1f6a0cc52(__obf_a4fef5c9586c1966)

	switch {
	case __obf_7a924d47297a474f == reflect.Bool:
		__obf_b422869c3c403707.
			SetBool(__obf_a4fef5c9586c1966.Bool())
	case __obf_7a924d47297a474f == reflect.Int && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_b422869c3c403707.
			SetBool(__obf_a4fef5c9586c1966.Int() != 0)
	case __obf_7a924d47297a474f == reflect.Uint && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_b422869c3c403707.
			SetBool(__obf_a4fef5c9586c1966.Uint() != 0)
	case __obf_7a924d47297a474f == reflect.Float32 && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_b422869c3c403707.
			SetBool(__obf_a4fef5c9586c1966.Float() != 0)
	case __obf_7a924d47297a474f == reflect.String && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_91f8fc20b44aae81, __obf_6585aef4313e6005 := strconv.ParseBool(__obf_a4fef5c9586c1966.String())
		if __obf_6585aef4313e6005 == nil {
			__obf_b422869c3c403707.
				SetBool(__obf_91f8fc20b44aae81)
		} else if __obf_a4fef5c9586c1966.String() == "" {
			__obf_b422869c3c403707.
				SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_a4fef5c9586c1966.Type(), __obf_333b2acf09d806b3)
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_051f2692931eb0df(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_7a924d47297a474f := __obf_2d4c05a1f6a0cc52(__obf_a4fef5c9586c1966)
	__obf_f73efdd6543886c9 := __obf_a4fef5c9586c1966.Type()

	switch {
	case __obf_7a924d47297a474f == reflect.Int:
		__obf_b422869c3c403707.
			SetFloat(float64(__obf_a4fef5c9586c1966.Int()))
	case __obf_7a924d47297a474f == reflect.Uint:
		__obf_b422869c3c403707.
			SetFloat(float64(__obf_a4fef5c9586c1966.Uint()))
	case __obf_7a924d47297a474f == reflect.Float32:
		__obf_b422869c3c403707.
			SetFloat(__obf_a4fef5c9586c1966.Float())
	case __obf_7a924d47297a474f == reflect.Bool && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		if __obf_a4fef5c9586c1966.Bool() {
			__obf_b422869c3c403707.
				SetFloat(1)
		} else {
			__obf_b422869c3c403707.
				SetFloat(0)
		}
	case __obf_7a924d47297a474f == reflect.String && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput:
		__obf_98acb43641f668a3 := __obf_a4fef5c9586c1966.String()
		if __obf_98acb43641f668a3 == "" {
			__obf_98acb43641f668a3 = "0"
		}
		__obf_4cd911a2d7099b07, __obf_6585aef4313e6005 := strconv.ParseFloat(__obf_98acb43641f668a3, __obf_b422869c3c403707.Type().Bits())
		if __obf_6585aef4313e6005 == nil {
			__obf_b422869c3c403707.
				SetFloat(__obf_4cd911a2d7099b07)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
	case __obf_f73efdd6543886c9.PkgPath() == "encoding/json" && __obf_f73efdd6543886c9.Name() == "Number":
		__obf_20127d5b36bcdbc4 := __obf_333b2acf09d806b3.(json.Number)
		__obf_99f0e7230007b81b, __obf_6585aef4313e6005 := __obf_20127d5b36bcdbc4.Float64()
		if __obf_6585aef4313e6005 != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_255580bc58cf586d, __obf_6585aef4313e6005)
		}
		__obf_b422869c3c403707.
			SetFloat(__obf_99f0e7230007b81b)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_a4fef5c9586c1966.Type(), __obf_333b2acf09d806b3)
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_34e83c4c958ecc46(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_0f0f0dc47b4c11f6 := __obf_b422869c3c403707.Type()
	__obf_9d4ff9a48374d663 := __obf_0f0f0dc47b4c11f6.Key()
	__obf_1ad97e308f12ff97 := __obf_0f0f0dc47b4c11f6.Elem()
	__obf_b318e1e0f10bf319 := // By default we overwrite keys in the current map
		__obf_b422869c3c403707

	// If the map is nil or we're purposely zeroing fields, make a new map
	if __obf_b318e1e0f10bf319.IsNil() || __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ZeroFields {
		__obf_9b9561b099ed0878 := // Make a new map to hold our result
			reflect.MapOf(__obf_9d4ff9a48374d663, __obf_1ad97e308f12ff97)
		__obf_b318e1e0f10bf319 = reflect.MakeMap(__obf_9b9561b099ed0878)
	}
	__obf_a4fef5c9586c1966 := // Check input type and based on the input type jump to the proper func
		reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	switch __obf_a4fef5c9586c1966.Kind() {
	case reflect.Map:
		return __obf_e47c7e0847f6ec2f.__obf_625cea809fa477ca(__obf_255580bc58cf586d, __obf_a4fef5c9586c1966, __obf_b422869c3c403707, __obf_b318e1e0f10bf319)

	case reflect.Struct:
		return __obf_e47c7e0847f6ec2f.__obf_97555ac6b6159b2f(__obf_a4fef5c9586c1966, __obf_b422869c3c403707, __obf_b318e1e0f10bf319)

	case reflect.Array, reflect.Slice:
		if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput {
			return __obf_e47c7e0847f6ec2f.__obf_cd9e333795d70208(__obf_255580bc58cf586d, __obf_a4fef5c9586c1966, __obf_b422869c3c403707, __obf_b318e1e0f10bf319)
		}

		fallthrough

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_255580bc58cf586d, __obf_a4fef5c9586c1966.Kind())
	}
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_cd9e333795d70208(__obf_255580bc58cf586d string, __obf_a4fef5c9586c1966 reflect.Value, __obf_b422869c3c403707 reflect.Value, __obf_b318e1e0f10bf319 reflect.Value) error {
	// Special case for BC reasons (covered by tests)
	if __obf_a4fef5c9586c1966.Len() == 0 {
		__obf_b422869c3c403707.
			Set(__obf_b318e1e0f10bf319)
		return nil
	}

	for __obf_99f0e7230007b81b := 0; __obf_99f0e7230007b81b < __obf_a4fef5c9586c1966.Len(); __obf_99f0e7230007b81b++ {
		__obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_255580bc58cf586d+
			"["+strconv.Itoa(__obf_99f0e7230007b81b)+"]", __obf_a4fef5c9586c1966.
			Index(__obf_99f0e7230007b81b).Interface(), __obf_b422869c3c403707)
		if __obf_6585aef4313e6005 != nil {
			return __obf_6585aef4313e6005
		}
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_625cea809fa477ca(__obf_255580bc58cf586d string, __obf_a4fef5c9586c1966 reflect.Value, __obf_b422869c3c403707 reflect.Value, __obf_b318e1e0f10bf319 reflect.Value) error {
	__obf_0f0f0dc47b4c11f6 := __obf_b422869c3c403707.Type()
	__obf_9d4ff9a48374d663 := __obf_0f0f0dc47b4c11f6.Key()
	__obf_1ad97e308f12ff97 := __obf_0f0f0dc47b4c11f6.Elem()

	// Accumulate errors
	errors := make([]string, 0)

	// If the input data is empty, then we just match what the input data is.
	if __obf_a4fef5c9586c1966.Len() == 0 {
		if __obf_a4fef5c9586c1966.IsNil() {
			if !__obf_b422869c3c403707.IsNil() {
				__obf_b422869c3c403707.
					Set(__obf_a4fef5c9586c1966)
			}
		} else {
			__obf_b422869c3c403707.
				// Set to empty allocated value
				Set(__obf_b318e1e0f10bf319)
		}

		return nil
	}

	for _, __obf_286aaa57afc54e13 := range __obf_a4fef5c9586c1966.MapKeys() {
		__obf_7dc427653c0378b5 := __obf_255580bc58cf586d + "[" + __obf_286aaa57afc54e13.String() + "]"
		__obf_224038468b299df9 := // First decode the key into the proper type
			reflect.Indirect(reflect.New(__obf_9d4ff9a48374d663))
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_7dc427653c0378b5, __obf_286aaa57afc54e13.Interface(), __obf_224038468b299df9); __obf_6585aef4313e6005 != nil {
			errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
			continue
		}
		__obf_3b55daf4bc34d8ef := // Next decode the data into the proper type
			__obf_a4fef5c9586c1966.MapIndex(__obf_286aaa57afc54e13).Interface()
		__obf_052c04c68f6b81ce := reflect.Indirect(reflect.New(__obf_1ad97e308f12ff97))
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_7dc427653c0378b5, __obf_3b55daf4bc34d8ef, __obf_052c04c68f6b81ce); __obf_6585aef4313e6005 != nil {
			errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
			continue
		}
		__obf_b318e1e0f10bf319.
			SetMapIndex(__obf_224038468b299df9, __obf_052c04c68f6b81ce)
	}
	__obf_b422869c3c403707.

		// Set the built up map to the value
		Set(__obf_b318e1e0f10bf319)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_97555ac6b6159b2f(__obf_a4fef5c9586c1966 reflect.Value, __obf_b422869c3c403707 reflect.Value, __obf_b318e1e0f10bf319 reflect.Value) error {
	__obf_0ce39965f55d3635 := __obf_a4fef5c9586c1966.Type()
	for __obf_99f0e7230007b81b := 0; __obf_99f0e7230007b81b < __obf_0ce39965f55d3635.NumField(); __obf_99f0e7230007b81b++ {
		__obf_4cd911a2d7099b07 := // Get the StructField first since this is a cheap operation. If the
			// field is unexported, then ignore it.
			__obf_0ce39965f55d3635.Field(__obf_99f0e7230007b81b)
		if __obf_4cd911a2d7099b07.PkgPath != "" {
			continue
		}
		__obf_3b55daf4bc34d8ef := // Next get the actual value of this field and verify it is assignable
			// to the map value.
			__obf_a4fef5c9586c1966.Field(__obf_99f0e7230007b81b)
		if !__obf_3b55daf4bc34d8ef.Type().AssignableTo(__obf_b318e1e0f10bf319.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", __obf_3b55daf4bc34d8ef.Type(), __obf_b318e1e0f10bf319.Type().Elem())
		}
		__obf_a7ffad665b9b475a := __obf_4cd911a2d7099b07.Tag.Get(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.TagName)
		__obf_2652134bfeb12996 := __obf_4cd911a2d7099b07.Name

		if __obf_a7ffad665b9b475a == "" && __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.IgnoreUntaggedFields {
			continue
		}
		__obf_1d2967a4fbf89d16 := // If Squash is set in the config, we squash the field down.
			__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Squash && __obf_3b55daf4bc34d8ef.Kind() == reflect.Struct && __obf_4cd911a2d7099b07.Anonymous
		__obf_3b55daf4bc34d8ef = __obf_97f4f196f76ac28c(__obf_3b55daf4bc34d8ef, __obf_e47c7e0847f6ec2f.

			// Determine the name of the key in the map
			__obf_9cbc86dfb44f54f3.TagName)

		if __obf_009b06b3de20e276 := strings.Index(__obf_a7ffad665b9b475a, ","); __obf_009b06b3de20e276 != -1 {
			if __obf_a7ffad665b9b475a[:__obf_009b06b3de20e276] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Contains(__obf_a7ffad665b9b475a[__obf_009b06b3de20e276+1:], "omitempty") && __obf_6231b046585c6bd7(__obf_3b55daf4bc34d8ef) {
				continue
			}
			__obf_1d2967a4fbf89d16 = // If "squash" is specified in the tag, we squash the field down.
				__obf_1d2967a4fbf89d16 || strings.Contains(__obf_a7ffad665b9b475a[__obf_009b06b3de20e276+1:], "squash")
			if __obf_1d2967a4fbf89d16 {
				// When squashing, the embedded type can be a pointer to a struct.
				if __obf_3b55daf4bc34d8ef.Kind() == reflect.Ptr && __obf_3b55daf4bc34d8ef.Elem().Kind() == reflect.Struct {
					__obf_3b55daf4bc34d8ef = __obf_3b55daf4bc34d8ef.Elem()
				}

				// The final type must be a struct
				if __obf_3b55daf4bc34d8ef.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", __obf_3b55daf4bc34d8ef.Type())
				}
			}
			if __obf_e4e0ee0855e988c2 := __obf_a7ffad665b9b475a[:__obf_009b06b3de20e276]; __obf_e4e0ee0855e988c2 != "" {
				__obf_2652134bfeb12996 = __obf_e4e0ee0855e988c2
			}
		} else if len(__obf_a7ffad665b9b475a) > 0 {
			if __obf_a7ffad665b9b475a == "-" {
				continue
			}
			__obf_2652134bfeb12996 = __obf_a7ffad665b9b475a
		}

		switch __obf_3b55daf4bc34d8ef.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			__obf_a0c38d95073a6b51 := reflect.New(__obf_3b55daf4bc34d8ef.Type())
			__obf_a0c38d95073a6b51.
				Elem().Set(__obf_3b55daf4bc34d8ef)
			__obf_589a2f1779a4e34a := __obf_b318e1e0f10bf319.Type()
			__obf_c2a2b926d2085909 := __obf_589a2f1779a4e34a.Key()
			__obf_4875bac7bb2e9a96 := __obf_589a2f1779a4e34a.Elem()
			__obf_b436f44dcad206da := reflect.MapOf(__obf_c2a2b926d2085909, __obf_4875bac7bb2e9a96)
			__obf_e79b1d656da1b72d := reflect.MakeMap(__obf_b436f44dcad206da)
			__obf_59d16ef2ee7b8e9a := // Creating a pointer to a map so that other methods can completely
				// overwrite the map if need be (looking at you decodeMapFromMap). The
				// indirection allows the underlying map to be settable (CanSet() == true)
				// where as reflect.MakeMap returns an unsettable map.
				reflect.New(__obf_e79b1d656da1b72d.Type())
			reflect.Indirect(__obf_59d16ef2ee7b8e9a).Set(__obf_e79b1d656da1b72d)
			__obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_2652134bfeb12996, __obf_a0c38d95073a6b51.Interface(), reflect.Indirect(__obf_59d16ef2ee7b8e9a))
			if __obf_6585aef4313e6005 != nil {
				return __obf_6585aef4313e6005
			}
			__obf_e79b1d656da1b72d = // the underlying map may have been completely overwritten so pull
				// it indirectly out of the enclosing value.
				reflect.Indirect(__obf_59d16ef2ee7b8e9a)

			if __obf_1d2967a4fbf89d16 {
				for _, __obf_286aaa57afc54e13 := range __obf_e79b1d656da1b72d.MapKeys() {
					__obf_b318e1e0f10bf319.
						SetMapIndex(__obf_286aaa57afc54e13, __obf_e79b1d656da1b72d.MapIndex(__obf_286aaa57afc54e13))
				}
			} else {
				__obf_b318e1e0f10bf319.
					SetMapIndex(reflect.ValueOf(__obf_2652134bfeb12996), __obf_e79b1d656da1b72d)
			}

		default:
			__obf_b318e1e0f10bf319.
				SetMapIndex(reflect.ValueOf(__obf_2652134bfeb12996), __obf_3b55daf4bc34d8ef)
		}
	}

	if __obf_b422869c3c403707.CanAddr() {
		__obf_b422869c3c403707.
			Set(__obf_b318e1e0f10bf319)
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_b2031a7d08ca7bbc(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) (bool, error) {
	__obf_d7eaf0936a366655 := // If the input data is nil, then we want to just set the output
		// pointer to be nil as well.
		__obf_333b2acf09d806b3 == nil
	if !__obf_d7eaf0936a366655 {
		switch __obf_3b55daf4bc34d8ef := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3)); __obf_3b55daf4bc34d8ef.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			__obf_d7eaf0936a366655 = __obf_3b55daf4bc34d8ef.IsNil()
		}
	}
	if __obf_d7eaf0936a366655 {
		if !__obf_b422869c3c403707.IsNil() && __obf_b422869c3c403707.CanSet() {
			__obf_026674d80634df49 := reflect.New(__obf_b422869c3c403707.Type()).Elem()
			__obf_b422869c3c403707.
				Set(__obf_026674d80634df49)
		}

		return true, nil
	}
	__obf_0f0f0dc47b4c11f6 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		__obf_b422869c3c403707.Type()
	__obf_1ad97e308f12ff97 := __obf_0f0f0dc47b4c11f6.Elem()
	if __obf_b422869c3c403707.CanSet() {
		__obf_f1f4167cbd1d2241 := __obf_b422869c3c403707
		if __obf_f1f4167cbd1d2241.IsNil() || __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ZeroFields {
			__obf_f1f4167cbd1d2241 = reflect.New(__obf_1ad97e308f12ff97)
		}

		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_255580bc58cf586d, __obf_333b2acf09d806b3, reflect.Indirect(__obf_f1f4167cbd1d2241)); __obf_6585aef4313e6005 != nil {
			return false, __obf_6585aef4313e6005
		}
		__obf_b422869c3c403707.
			Set(__obf_f1f4167cbd1d2241)
	} else {
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_255580bc58cf586d, __obf_333b2acf09d806b3, reflect.Indirect(__obf_b422869c3c403707)); __obf_6585aef4313e6005 != nil {
			return false, __obf_6585aef4313e6005
		}
	}
	return false, nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_4c760466d764236a(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	if __obf_b422869c3c403707.Type() != __obf_a4fef5c9586c1966.Type() {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_255580bc58cf586d, __obf_b422869c3c403707.
				Type(), __obf_a4fef5c9586c1966.Type(), __obf_333b2acf09d806b3)
	}
	__obf_b422869c3c403707.
		Set(__obf_a4fef5c9586c1966)
	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_6cc4d6625c3961b3(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_49cf32094ca76abc := __obf_a4fef5c9586c1966.Kind()
	__obf_0f0f0dc47b4c11f6 := __obf_b422869c3c403707.Type()
	__obf_1ad97e308f12ff97 := __obf_0f0f0dc47b4c11f6.Elem()
	__obf_52a8a94ef473454c := reflect.SliceOf(__obf_1ad97e308f12ff97)

	// If we have a non array/slice type then we first attempt to convert.
	if __obf_49cf32094ca76abc != reflect.Array && __obf_49cf32094ca76abc != reflect.Slice {
		if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case __obf_49cf32094ca76abc == reflect.Slice, __obf_49cf32094ca76abc == reflect.Array:
				break

			// Empty maps turn into empty slices
			case __obf_49cf32094ca76abc == reflect.Map:
				if __obf_a4fef5c9586c1966.Len() == 0 {
					__obf_b422869c3c403707.
						Set(reflect.MakeSlice(__obf_52a8a94ef473454c, 0, 0))
					return nil
				}
				// Create slice of maps of other sizes
				return __obf_e47c7e0847f6ec2f.__obf_6cc4d6625c3961b3(__obf_255580bc58cf586d, []any{__obf_333b2acf09d806b3}, __obf_b422869c3c403707)

			case __obf_49cf32094ca76abc == reflect.String && __obf_1ad97e308f12ff97.Kind() == reflect.Uint8:
				return __obf_e47c7e0847f6ec2f.__obf_6cc4d6625c3961b3(__obf_255580bc58cf586d, []byte(__obf_a4fef5c9586c1966.String()), __obf_b422869c3c403707)

			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				return __obf_e47c7e0847f6ec2f.__obf_6cc4d6625c3961b3(__obf_255580bc58cf586d, []any{__obf_333b2acf09d806b3}, __obf_b422869c3c403707)
			}
		}

		return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", __obf_255580bc58cf586d, __obf_49cf32094ca76abc)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if __obf_49cf32094ca76abc != reflect.Array && __obf_a4fef5c9586c1966.IsNil() {
		return nil
	}
	__obf_3f4615e5a25ac96a := __obf_b422869c3c403707
	if __obf_3f4615e5a25ac96a.IsNil() || __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ZeroFields {
		__obf_3f4615e5a25ac96a = // Make a new slice to hold our result, same size as the original data.
			reflect.MakeSlice(__obf_52a8a94ef473454c, __obf_a4fef5c9586c1966.Len(), __obf_a4fef5c9586c1966.Len())
	} else if __obf_3f4615e5a25ac96a.Len() > __obf_a4fef5c9586c1966.Len() {
		__obf_3f4615e5a25ac96a = __obf_3f4615e5a25ac96a.Slice(0, __obf_a4fef5c9586c1966.Len())
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_99f0e7230007b81b := 0; __obf_99f0e7230007b81b < __obf_a4fef5c9586c1966.Len(); __obf_99f0e7230007b81b++ {
		__obf_71775d454b314261 := __obf_a4fef5c9586c1966.Index(__obf_99f0e7230007b81b).Interface()
		for __obf_3f4615e5a25ac96a.Len() <= __obf_99f0e7230007b81b {
			__obf_3f4615e5a25ac96a = reflect.Append(__obf_3f4615e5a25ac96a, reflect.Zero(__obf_1ad97e308f12ff97))
		}
		__obf_c0d0ff0291a611e2 := __obf_3f4615e5a25ac96a.Index(__obf_99f0e7230007b81b)
		__obf_7dc427653c0378b5 := __obf_255580bc58cf586d + "[" + strconv.Itoa(__obf_99f0e7230007b81b) + "]"
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_7dc427653c0378b5, __obf_71775d454b314261, __obf_c0d0ff0291a611e2); __obf_6585aef4313e6005 != nil {
			errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
		}
	}
	__obf_b422869c3c403707.

		// Finally, set the value to the slice we built up
		Set(__obf_3f4615e5a25ac96a)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_d961bcc337062438(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))
	__obf_49cf32094ca76abc := __obf_a4fef5c9586c1966.Kind()
	__obf_0f0f0dc47b4c11f6 := __obf_b422869c3c403707.Type()
	__obf_1ad97e308f12ff97 := __obf_0f0f0dc47b4c11f6.Elem()
	__obf_da439654c7d9df87 := reflect.ArrayOf(__obf_0f0f0dc47b4c11f6.Len(), __obf_1ad97e308f12ff97)
	__obf_7d1e33974f3f30ad := __obf_b422869c3c403707

	if __obf_7d1e33974f3f30ad.Interface() == reflect.Zero(__obf_7d1e33974f3f30ad.Type()).Interface() || __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ZeroFields {
		// Check input type
		if __obf_49cf32094ca76abc != reflect.Array && __obf_49cf32094ca76abc != reflect.Slice {
			if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.WeaklyTypedInput {
				switch __obf_49cf32094ca76abc {
				// Empty maps turn into empty arrays
				case reflect.Map:
					if __obf_a4fef5c9586c1966.Len() == 0 {
						__obf_b422869c3c403707.
							Set(reflect.Zero(__obf_da439654c7d9df87))
						return nil
					}

				// All other types we try to convert to the array type
				// and "lift" it into it. i.e. a string becomes a string array.
				default:
					// Just re-try this function with data as a slice.
					return __obf_e47c7e0847f6ec2f.__obf_d961bcc337062438(__obf_255580bc58cf586d, []any{__obf_333b2acf09d806b3}, __obf_b422869c3c403707)
				}
			}

			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", __obf_255580bc58cf586d, __obf_49cf32094ca76abc)

		}
		if __obf_a4fef5c9586c1966.Len() > __obf_da439654c7d9df87.Len() {
			return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", __obf_255580bc58cf586d, __obf_da439654c7d9df87.Len(), __obf_a4fef5c9586c1966.Len())

		}
		__obf_7d1e33974f3f30ad = // Make a new array to hold our result, same size as the original data.
			reflect.New(__obf_da439654c7d9df87).Elem()
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_99f0e7230007b81b := 0; __obf_99f0e7230007b81b < __obf_a4fef5c9586c1966.Len(); __obf_99f0e7230007b81b++ {
		__obf_71775d454b314261 := __obf_a4fef5c9586c1966.Index(__obf_99f0e7230007b81b).Interface()
		__obf_c0d0ff0291a611e2 := __obf_7d1e33974f3f30ad.Index(__obf_99f0e7230007b81b)
		__obf_7dc427653c0378b5 := __obf_255580bc58cf586d + "[" + strconv.Itoa(__obf_99f0e7230007b81b) + "]"
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_7dc427653c0378b5, __obf_71775d454b314261, __obf_c0d0ff0291a611e2); __obf_6585aef4313e6005 != nil {
			errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
		}
	}
	__obf_b422869c3c403707.

		// Finally, set the value to the array we built up
		Set(__obf_7d1e33974f3f30ad)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_70a1673aa1d5007b(__obf_255580bc58cf586d string, __obf_333b2acf09d806b3 any, __obf_b422869c3c403707 reflect.Value) error {
	__obf_a4fef5c9586c1966 := reflect.Indirect(reflect.ValueOf(__obf_333b2acf09d806b3))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if __obf_a4fef5c9586c1966.Type() == __obf_b422869c3c403707.Type() {
		__obf_b422869c3c403707.
			Set(__obf_a4fef5c9586c1966)
		return nil
	}
	__obf_49cf32094ca76abc := __obf_a4fef5c9586c1966.Kind()
	switch __obf_49cf32094ca76abc {
	case reflect.Map:
		return __obf_e47c7e0847f6ec2f.__obf_b9a94569a06dfb7f(__obf_255580bc58cf586d, __obf_a4fef5c9586c1966,

			// Not the most efficient way to do this but we can optimize later if
			// we want to. To convert from struct to struct we go to map first
			// as an intermediary.
			__obf_b422869c3c403707)

	case reflect.Struct:
		__obf_9b9561b099ed0878 := // Make a new map to hold our result
			reflect.TypeOf((map[string]any)(nil))
		__obf_0752c184c4fefebe := reflect.MakeMap(__obf_9b9561b099ed0878)
		__obf_59d16ef2ee7b8e9a := // Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			reflect.New(__obf_0752c184c4fefebe.Type())

		reflect.Indirect(__obf_59d16ef2ee7b8e9a).Set(__obf_0752c184c4fefebe)
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_97555ac6b6159b2f(__obf_a4fef5c9586c1966, reflect.Indirect(__obf_59d16ef2ee7b8e9a), __obf_0752c184c4fefebe); __obf_6585aef4313e6005 != nil {
			return __obf_6585aef4313e6005
		}
		__obf_382efc6537d96d52 := __obf_e47c7e0847f6ec2f.__obf_b9a94569a06dfb7f(__obf_255580bc58cf586d, reflect.Indirect(__obf_59d16ef2ee7b8e9a), __obf_b422869c3c403707)
		return __obf_382efc6537d96d52

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_255580bc58cf586d, __obf_a4fef5c9586c1966.Kind())
	}
}

func (__obf_e47c7e0847f6ec2f *Decoder) __obf_b9a94569a06dfb7f(__obf_255580bc58cf586d string, __obf_a4fef5c9586c1966, __obf_b422869c3c403707 reflect.Value) error {
	__obf_796baf2c45c67595 := __obf_a4fef5c9586c1966.Type()
	if __obf_5b98b7b2a0e7f235 := __obf_796baf2c45c67595.Key().Kind(); __obf_5b98b7b2a0e7f235 != reflect.String && __obf_5b98b7b2a0e7f235 != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys", __obf_255580bc58cf586d, __obf_796baf2c45c67595.
				Key().Kind())
	}
	__obf_45708d6e50b5932a := make(map[reflect.Value]struct{})
	__obf_5cbf32022761eee0 := make(map[any]struct{})
	for _, __obf_99de953620bf1b40 := range __obf_a4fef5c9586c1966.MapKeys() {
		__obf_45708d6e50b5932a[__obf_99de953620bf1b40] = struct{}{}
		__obf_5cbf32022761eee0[__obf_99de953620bf1b40.Interface()] = struct{}{}
	}
	__obf_b3fca7be3200e7ee := make(map[any]struct{})
	errors := make([]string, 0)
	__obf_8fb3a5479269cc35 := // This slice will keep track of all the structs we'll be decoding.
		// There can be more than one struct if there are embedded structs
		// that are squashed.
		make([]reflect.Value, 1, 5)
	__obf_8fb3a5479269cc35[0] = __obf_b422869c3c403707

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type __obf_3388a9eeff976d58 struct {
		__obf_3388a9eeff976d58 reflect.StructField
		__obf_b422869c3c403707 reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var __obf_28352443449d0d26 *__obf_3388a9eeff976d58
	__obf_5e668f1153ddd53d := []__obf_3388a9eeff976d58{}
	for len(__obf_8fb3a5479269cc35) > 0 {
		__obf_901145da80e3e4a0 := __obf_8fb3a5479269cc35[0]
		__obf_8fb3a5479269cc35 = __obf_8fb3a5479269cc35[1:]
		__obf_72d2762bcea406f0 := __obf_901145da80e3e4a0.Type()

		for __obf_99f0e7230007b81b := 0; __obf_99f0e7230007b81b < __obf_72d2762bcea406f0.NumField(); __obf_99f0e7230007b81b++ {
			__obf_812ad5948c1198d0 := __obf_72d2762bcea406f0.Field(__obf_99f0e7230007b81b)
			__obf_3005c79ff2cf5d14 := __obf_901145da80e3e4a0.Field(__obf_99f0e7230007b81b)
			if __obf_3005c79ff2cf5d14.Kind() == reflect.Pointer && __obf_3005c79ff2cf5d14.Elem().Kind() == reflect.Struct {
				__obf_3005c79ff2cf5d14 = // Handle embedded struct pointers as embedded structs.
					__obf_3005c79ff2cf5d14.Elem()
			}
			__obf_1d2967a4fbf89d16 := // If "squash" is specified in the tag, we squash the field down.
				__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Squash && __obf_3005c79ff2cf5d14.Kind() == reflect.Struct && __obf_812ad5948c1198d0.Anonymous
			__obf_61a8b9b31a9d6283 := false
			__obf_6466ede07f7eb97b := // We always parse the tags cause we're looking for other tags too
				strings.Split(__obf_812ad5948c1198d0.Tag.Get(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.TagName), ",")
			for _, __obf_d2af23c6ba65ab9e := range __obf_6466ede07f7eb97b[1:] {
				if __obf_d2af23c6ba65ab9e == "squash" {
					__obf_1d2967a4fbf89d16 = true
					break
				}

				if __obf_d2af23c6ba65ab9e == "remain" {
					__obf_61a8b9b31a9d6283 = true
					break
				}
			}

			if __obf_1d2967a4fbf89d16 {
				if __obf_3005c79ff2cf5d14.Kind() != reflect.Struct {
					errors = __obf_513ecd3219b0baa9(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", __obf_812ad5948c1198d0.Name, __obf_3005c79ff2cf5d14.Kind()))
				} else {
					__obf_8fb3a5479269cc35 = append(__obf_8fb3a5479269cc35, __obf_3005c79ff2cf5d14)
				}
				continue
			}

			// Build our field
			if __obf_61a8b9b31a9d6283 {
				__obf_28352443449d0d26 = &__obf_3388a9eeff976d58{__obf_812ad5948c1198d0, __obf_3005c79ff2cf5d14}
			} else {
				__obf_5e668f1153ddd53d = // Normal struct field, store it away
					append(__obf_5e668f1153ddd53d, __obf_3388a9eeff976d58{__obf_812ad5948c1198d0, __obf_3005c79ff2cf5d14})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, __obf_4cd911a2d7099b07 := range __obf_5e668f1153ddd53d {
		__obf_3388a9eeff976d58, __obf_521ce878bf94986e := __obf_4cd911a2d7099b07.__obf_3388a9eeff976d58, __obf_4cd911a2d7099b07.__obf_b422869c3c403707
		__obf_7dc427653c0378b5 := __obf_3388a9eeff976d58.Name
		__obf_a7ffad665b9b475a := __obf_3388a9eeff976d58.Tag.Get(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.TagName)
		__obf_a7ffad665b9b475a = strings.SplitN(__obf_a7ffad665b9b475a, ",", 2)[0]
		if __obf_a7ffad665b9b475a != "" {
			__obf_7dc427653c0378b5 = __obf_a7ffad665b9b475a
		}
		__obf_a081e9eb5dde0e7b := reflect.ValueOf(__obf_7dc427653c0378b5)
		__obf_0a741326bd62ec68 := __obf_a4fef5c9586c1966.MapIndex(__obf_a081e9eb5dde0e7b)
		if !__obf_0a741326bd62ec68.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for __obf_99de953620bf1b40 := range __obf_45708d6e50b5932a {
				__obf_c60c70442772ce87, __obf_c5c13e02010ecc6c := __obf_99de953620bf1b40.Interface().(string)
				if !__obf_c5c13e02010ecc6c {
					// Not a string key
					continue
				}

				if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.MatchName(__obf_c60c70442772ce87, __obf_7dc427653c0378b5) {
					__obf_a081e9eb5dde0e7b = __obf_99de953620bf1b40
					__obf_0a741326bd62ec68 = __obf_a4fef5c9586c1966.MapIndex(__obf_99de953620bf1b40)
					break
				}
			}

			if !__obf_0a741326bd62ec68.IsValid() {
				__obf_b3fca7be3200e7ee[ // There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				__obf_7dc427653c0378b5] = struct{}{}
				continue
			}
		}

		if !__obf_521ce878bf94986e.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !__obf_521ce878bf94986e.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(__obf_5cbf32022761eee0, __obf_a081e9eb5dde0e7b.Interface())

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if __obf_255580bc58cf586d != "" {
			__obf_7dc427653c0378b5 = __obf_255580bc58cf586d + "." + __obf_7dc427653c0378b5
		}

		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_f7c6980d711168d4(__obf_7dc427653c0378b5, __obf_0a741326bd62ec68.Interface(), __obf_521ce878bf94986e); __obf_6585aef4313e6005 != nil {
			errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if __obf_28352443449d0d26 != nil && len(__obf_5cbf32022761eee0) > 0 {
		__obf_61a8b9b31a9d6283 := // Build a map of only the unused values
			map[any]any{}
		for __obf_a79812f955505d02 := range __obf_5cbf32022761eee0 {
			__obf_61a8b9b31a9d6283[__obf_a79812f955505d02] = __obf_a4fef5c9586c1966.MapIndex(reflect.ValueOf(__obf_a79812f955505d02)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if __obf_6585aef4313e6005 := __obf_e47c7e0847f6ec2f.__obf_34e83c4c958ecc46(__obf_255580bc58cf586d, __obf_61a8b9b31a9d6283, __obf_28352443449d0d26.__obf_b422869c3c403707); __obf_6585aef4313e6005 != nil {
			errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
		}
		__obf_5cbf32022761eee0 = // Set the map to nil so we have none so that the next check will
			// not error (ErrorUnused)
			nil
	}

	if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ErrorUnused && len(__obf_5cbf32022761eee0) > 0 {
		__obf_0122b16769bf1e29 := make([]string, 0, len(__obf_5cbf32022761eee0))
		for __obf_c895a6f31c8bfed1 := range __obf_5cbf32022761eee0 {
			__obf_0122b16769bf1e29 = append(__obf_0122b16769bf1e29, __obf_c895a6f31c8bfed1.(string))
		}
		sort.Strings(__obf_0122b16769bf1e29)
		__obf_6585aef4313e6005 := fmt.Errorf("'%s' has invalid keys: %s", __obf_255580bc58cf586d, strings.Join(__obf_0122b16769bf1e29, ", "))
		errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
	}

	if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.ErrorUnset && len(__obf_b3fca7be3200e7ee) > 0 {
		__obf_0122b16769bf1e29 := make([]string, 0, len(__obf_b3fca7be3200e7ee))
		for __obf_c895a6f31c8bfed1 := range __obf_b3fca7be3200e7ee {
			__obf_0122b16769bf1e29 = append(__obf_0122b16769bf1e29, __obf_c895a6f31c8bfed1.(string))
		}
		sort.Strings(__obf_0122b16769bf1e29)
		__obf_6585aef4313e6005 := fmt.Errorf("'%s' has unset fields: %s", __obf_255580bc58cf586d, strings.Join(__obf_0122b16769bf1e29, ", "))
		errors = __obf_513ecd3219b0baa9(errors, __obf_6585aef4313e6005)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	if __obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata != nil {
		for __obf_c895a6f31c8bfed1 := range __obf_5cbf32022761eee0 {
			__obf_a79812f955505d02 := __obf_c895a6f31c8bfed1.(string)
			if __obf_255580bc58cf586d != "" {
				__obf_a79812f955505d02 = __obf_255580bc58cf586d + "." + __obf_a79812f955505d02
			}
			__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.
				Metadata.Unused = append(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata.Unused, __obf_a79812f955505d02)
		}
		for __obf_c895a6f31c8bfed1 := range __obf_b3fca7be3200e7ee {
			__obf_a79812f955505d02 := __obf_c895a6f31c8bfed1.(string)
			if __obf_255580bc58cf586d != "" {
				__obf_a79812f955505d02 = __obf_255580bc58cf586d + "." + __obf_a79812f955505d02
			}
			__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.
				Metadata.Unset = append(__obf_e47c7e0847f6ec2f.__obf_9cbc86dfb44f54f3.Metadata.Unset, __obf_a79812f955505d02)
		}
	}

	return nil
}

func __obf_6231b046585c6bd7(__obf_3b55daf4bc34d8ef reflect.Value) bool {
	switch __obf_2d4c05a1f6a0cc52(__obf_3b55daf4bc34d8ef) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_3b55daf4bc34d8ef.Len() == 0
	case reflect.Bool:
		return !__obf_3b55daf4bc34d8ef.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_3b55daf4bc34d8ef.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_3b55daf4bc34d8ef.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_3b55daf4bc34d8ef.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return __obf_3b55daf4bc34d8ef.IsNil()
	}
	return false
}

func __obf_2d4c05a1f6a0cc52(__obf_b422869c3c403707 reflect.Value) reflect.Kind {
	__obf_5b98b7b2a0e7f235 := __obf_b422869c3c403707.Kind()

	switch {
	case __obf_5b98b7b2a0e7f235 >= reflect.Int && __obf_5b98b7b2a0e7f235 <= reflect.Int64:
		return reflect.Int
	case __obf_5b98b7b2a0e7f235 >= reflect.Uint && __obf_5b98b7b2a0e7f235 <= reflect.Uint64:
		return reflect.Uint
	case __obf_5b98b7b2a0e7f235 >= reflect.Float32 && __obf_5b98b7b2a0e7f235 <= reflect.Float64:
		return reflect.Float32
	default:
		return __obf_5b98b7b2a0e7f235
	}
}

func __obf_e470c6f0f4c4fb2a(__obf_0ce39965f55d3635 reflect.Type, __obf_b833b88590e3c92c bool, __obf_fa55af5342352ddc string) bool {
	for __obf_99f0e7230007b81b := 0; __obf_99f0e7230007b81b < __obf_0ce39965f55d3635.NumField(); __obf_99f0e7230007b81b++ {
		__obf_4cd911a2d7099b07 := __obf_0ce39965f55d3635.Field(__obf_99f0e7230007b81b)
		if __obf_4cd911a2d7099b07.PkgPath == "" && !__obf_b833b88590e3c92c { // check for unexported fields
			return true
		}
		if __obf_b833b88590e3c92c && __obf_4cd911a2d7099b07.Tag.Get(__obf_fa55af5342352ddc) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func __obf_97f4f196f76ac28c(__obf_3b55daf4bc34d8ef reflect.Value, __obf_fa55af5342352ddc string) reflect.Value {
	if __obf_3b55daf4bc34d8ef.Kind() != reflect.Ptr || __obf_3b55daf4bc34d8ef.Elem().Kind() != reflect.Struct {
		return __obf_3b55daf4bc34d8ef
	}
	__obf_3fca72dd952a32ae := __obf_3b55daf4bc34d8ef.Elem()
	__obf_72ee56c58fb7b372 := __obf_3fca72dd952a32ae.Type()
	if __obf_e470c6f0f4c4fb2a(__obf_72ee56c58fb7b372, true, __obf_fa55af5342352ddc) {
		return __obf_3fca72dd952a32ae
	}
	return __obf_3b55daf4bc34d8ef
}
