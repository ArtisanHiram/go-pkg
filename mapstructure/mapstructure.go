package __obf_c953b7a5114a5dbe

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
type DecodeHookFuncValue func(__obf_14810ac6c1a1ff4b reflect.Value, __obf_bb1596a458011647 reflect.Value) (any, error)

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
	MatchName func(__obf_aaad2cb91e02253b, __obf_f2a21f20827fddb2 string) bool
}

// A Decoder takes a raw interface value and turns it into structured
// data, keeping track of rich error information along the way in case
// anything goes wrong. Unlike the basic top-level Decode method, you can
// more finely control how the Decoder behaves using the DecoderConfig
// structure. The top-level Decode method is just a convenience that sets
// up the most basic Decoder.
type Decoder struct {
	__obf_f7db361728796f85 *DecoderConfig
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
func Decode(__obf_a32fd44226d0eaf4 any, __obf_0c5869ab7b689e34 any) error {
	__obf_f7db361728796f85 := &DecoderConfig{
		Metadata: nil,
		Result:   __obf_0c5869ab7b689e34,
	}
	__obf_fd96af80d8b50f34, __obf_fb7fbd58154fa1dd := NewDecoder(__obf_f7db361728796f85)
	if __obf_fb7fbd58154fa1dd != nil {
		return __obf_fb7fbd58154fa1dd
	}

	return __obf_fd96af80d8b50f34.Decode(__obf_a32fd44226d0eaf4)
}

// WeakDecode is the same as Decode but is shorthand to enable
// WeaklyTypedInput. See DecoderConfig for more info.
func WeakDecode(__obf_a32fd44226d0eaf4, __obf_0c5869ab7b689e34 any) error {
	__obf_f7db361728796f85 := &DecoderConfig{
		Metadata:         nil,
		Result:           __obf_0c5869ab7b689e34,
		WeaklyTypedInput: true,
	}
	__obf_fd96af80d8b50f34, __obf_fb7fbd58154fa1dd := NewDecoder(__obf_f7db361728796f85)
	if __obf_fb7fbd58154fa1dd != nil {
		return __obf_fb7fbd58154fa1dd
	}

	return __obf_fd96af80d8b50f34.Decode(__obf_a32fd44226d0eaf4)
}

// DecodeMetadata is the same as Decode, but is shorthand to
// enable metadata collection. See DecoderConfig for more info.
func DecodeMetadata(__obf_a32fd44226d0eaf4 any, __obf_0c5869ab7b689e34 any, __obf_8322a0badc99400f *Metadata) error {
	__obf_f7db361728796f85 := &DecoderConfig{
		Metadata: __obf_8322a0badc99400f,
		Result:   __obf_0c5869ab7b689e34,
	}
	__obf_fd96af80d8b50f34, __obf_fb7fbd58154fa1dd := NewDecoder(__obf_f7db361728796f85)
	if __obf_fb7fbd58154fa1dd != nil {
		return __obf_fb7fbd58154fa1dd
	}

	return __obf_fd96af80d8b50f34.Decode(__obf_a32fd44226d0eaf4)
}

// WeakDecodeMetadata is the same as Decode, but is shorthand to
// enable both WeaklyTypedInput and metadata collection. See
// DecoderConfig for more info.
func WeakDecodeMetadata(__obf_a32fd44226d0eaf4 any, __obf_0c5869ab7b689e34 any, __obf_8322a0badc99400f *Metadata) error {
	__obf_f7db361728796f85 := &DecoderConfig{
		Metadata:         __obf_8322a0badc99400f,
		Result:           __obf_0c5869ab7b689e34,
		WeaklyTypedInput: true,
	}
	__obf_fd96af80d8b50f34, __obf_fb7fbd58154fa1dd := NewDecoder(__obf_f7db361728796f85)
	if __obf_fb7fbd58154fa1dd != nil {
		return __obf_fb7fbd58154fa1dd
	}

	return __obf_fd96af80d8b50f34.Decode(__obf_a32fd44226d0eaf4)
}

// NewDecoder returns a new decoder for the given configuration. Once
// a decoder has been returned, the same configuration must not be used
// again.
func NewDecoder(__obf_f7db361728796f85 *DecoderConfig) (*Decoder, error) {
	__obf_55480fa872ab8d98 := reflect.ValueOf(__obf_f7db361728796f85.Result)
	if __obf_55480fa872ab8d98.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}
	__obf_55480fa872ab8d98 = __obf_55480fa872ab8d98.Elem()
	if !__obf_55480fa872ab8d98.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if __obf_f7db361728796f85.Metadata != nil {
		if __obf_f7db361728796f85.Metadata.Keys == nil {
			__obf_f7db361728796f85.
				Metadata.Keys = make([]string, 0)
		}

		if __obf_f7db361728796f85.Metadata.Unused == nil {
			__obf_f7db361728796f85.
				Metadata.Unused = make([]string, 0)
		}

		if __obf_f7db361728796f85.Metadata.Unset == nil {
			__obf_f7db361728796f85.
				Metadata.Unset = make([]string, 0)
		}
	}

	if __obf_f7db361728796f85.TagName == "" {
		__obf_f7db361728796f85.
			TagName = "mapstructure"
	}

	if __obf_f7db361728796f85.MatchName == nil {
		__obf_f7db361728796f85.
			MatchName = strings.EqualFold
	}
	__obf_47a3e13e334b0913 := &Decoder{__obf_f7db361728796f85: __obf_f7db361728796f85}

	return __obf_47a3e13e334b0913, nil
}

// Decode decodes the given raw interface to the target pointer specified
// by the configuration.
func (__obf_1b8775533a1609ca *Decoder) Decode(__obf_a32fd44226d0eaf4 any) error {
	return __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238("", __obf_a32fd44226d0eaf4, reflect.ValueOf(__obf_1b8775533a1609ca.__obf_f7db361728796f85.Result).Elem())
}

// Decodes an unknown data type into a specific reflection value.
func (__obf_1b8775533a1609ca *Decoder) __obf_0473c83ce7e0b238(__obf_f36d6422fc6a0190 string, __obf_a32fd44226d0eaf4 any, __obf_ecf492ff0db025f4 reflect.Value) error {
	var __obf_df7e120891da2d5e reflect.Value
	if __obf_a32fd44226d0eaf4 != nil {
		__obf_df7e120891da2d5e = reflect.ValueOf(__obf_a32fd44226d0eaf4)

		// We need to check here if input is a typed nil. Typed nils won't
		// match the "input == nil" below so we check that here.
		if __obf_df7e120891da2d5e.Kind() == reflect.Ptr && __obf_df7e120891da2d5e.IsNil() {
			__obf_a32fd44226d0eaf4 = nil
		}
	}

	if __obf_a32fd44226d0eaf4 == nil {
		// If the data is nil, then we don't set anything, unless ZeroFields is set
		// to true.
		if __obf_1b8775533a1609ca.__obf_f7db361728796f85.ZeroFields {
			__obf_ecf492ff0db025f4.
				Set(reflect.Zero(__obf_ecf492ff0db025f4.Type()))

			if __obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata != nil && __obf_f36d6422fc6a0190 != "" {
				__obf_1b8775533a1609ca.__obf_f7db361728796f85.
					Metadata.Keys = append(__obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata.Keys, __obf_f36d6422fc6a0190)
			}
		}
		return nil
	}

	if !__obf_df7e120891da2d5e.IsValid() {
		__obf_ecf492ff0db025f4.
			// If the input value is invalid, then we just set the value
			// to be the zero value.
			Set(reflect.Zero(__obf_ecf492ff0db025f4.Type()))
		if __obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata != nil && __obf_f36d6422fc6a0190 != "" {
			__obf_1b8775533a1609ca.__obf_f7db361728796f85.
				Metadata.Keys = append(__obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata.Keys, __obf_f36d6422fc6a0190)
		}
		return nil
	}

	if __obf_1b8775533a1609ca.__obf_f7db361728796f85.DecodeHook != nil {
		// We have a DecodeHook, so let's pre-process the input.
		var __obf_fb7fbd58154fa1dd error
		__obf_a32fd44226d0eaf4, __obf_fb7fbd58154fa1dd = DecodeHookExec(__obf_1b8775533a1609ca.__obf_f7db361728796f85.DecodeHook, __obf_df7e120891da2d5e, __obf_ecf492ff0db025f4)
		if __obf_fb7fbd58154fa1dd != nil {
			return fmt.Errorf("error decoding '%s': %w", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
	}

	var __obf_fb7fbd58154fa1dd error
	__obf_756f5959a9b63a19 := __obf_c57e419440fe363b(__obf_ecf492ff0db025f4)
	__obf_014977bef48dafea := true
	switch __obf_756f5959a9b63a19 {
	case reflect.Bool:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_b32c0c2e0abc7678(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Interface:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_7e7fef9733f2775b(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.String:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_67478b6409640cb0(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Int:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_b89fb2c2f196ef6a(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Uint:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_04fc611047ec2b16(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Float32:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_56cc256ab94979c3(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Struct:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_4b5ce3ffb9b54ebf(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Map:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_bdc14c5e7a664ecd(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Ptr:
		__obf_014977bef48dafea, __obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_759141d66f396441(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Slice:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_ec5d8403f901afa0(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Array:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_3d3439b3dca75bf7(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4, __obf_ecf492ff0db025f4)
	case reflect.Func:
		__obf_fb7fbd58154fa1dd = __obf_1b8775533a1609ca.__obf_53abac02713a0541(__obf_f36d6422fc6a0190, __obf_a32fd44226d0eaf4,

			// If we reached this point then we weren't able to decode it
			__obf_ecf492ff0db025f4)
	default:

		return fmt.Errorf("%s: unsupported type: %s", __obf_f36d6422fc6a0190, __obf_756f5959a9b63a19)
	}

	// If we reached here, then we successfully decoded SOMETHING, so
	// mark the key as used if we're tracking metainput.
	if __obf_014977bef48dafea && __obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata != nil && __obf_f36d6422fc6a0190 != "" {
		__obf_1b8775533a1609ca.__obf_f7db361728796f85.
			Metadata.Keys = append(__obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata.Keys, __obf_f36d6422fc6a0190)
	}

	return __obf_fb7fbd58154fa1dd
}

// This decodes a basic type (bool, int, string, etc.) and sets the
// value to "data" of that type.
func (__obf_1b8775533a1609ca *Decoder) __obf_7e7fef9733f2775b(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	if __obf_55480fa872ab8d98.IsValid() && __obf_55480fa872ab8d98.Elem().IsValid() {
		__obf_7d8fb87b5ab9610e := __obf_55480fa872ab8d98.Elem()
		__obf_e0abfccd8e64fcb5 := // If we can't address this element, then its not writable. Instead,
			// we make a copy of the value (which is a pointer and therefore
			// writable), decode into that, and replace the whole value.
			false
		if !__obf_7d8fb87b5ab9610e.CanAddr() {
			__obf_e0abfccd8e64fcb5 = true

			// Make *T
			copy := reflect.New(__obf_7d8fb87b5ab9610e.Type())

			// *T = elem
			copy.Elem().Set(__obf_7d8fb87b5ab9610e)
			__obf_7d8fb87b5ab9610e = // Set elem so we decode into it
				copy
		}

		// Decode. If we have an error then return. We also return right
		// away if we're not a copy because that means we decoded directly.
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f36d6422fc6a0190, __obf_50c6df487a55f30e, __obf_7d8fb87b5ab9610e); __obf_fb7fbd58154fa1dd != nil || !__obf_e0abfccd8e64fcb5 {
			return __obf_fb7fbd58154fa1dd
		}
		__obf_55480fa872ab8d98.

			// If we're a copy, we need to set te final result
			Set(__obf_7d8fb87b5ab9610e.Elem())
		return nil
	}
	__obf_a361258b272167f7 := reflect.ValueOf(__obf_50c6df487a55f30e)

	// If the input data is a pointer, and the assigned type is the dereference
	// of that exact pointer, then indirect it so that we can assign it.
	// Example: *string to string
	if __obf_a361258b272167f7.Kind() == reflect.Ptr && __obf_a361258b272167f7.Type().Elem() == __obf_55480fa872ab8d98.Type() {
		__obf_a361258b272167f7 = reflect.Indirect(__obf_a361258b272167f7)
	}

	if !__obf_a361258b272167f7.IsValid() {
		__obf_a361258b272167f7 = reflect.Zero(__obf_55480fa872ab8d98.Type())
	}
	__obf_97e84555c565f923 := __obf_a361258b272167f7.Type()
	if !__obf_97e84555c565f923.AssignableTo(__obf_55480fa872ab8d98.Type()) {
		return fmt.Errorf(
			"'%s' expected type '%s', got '%s'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_97e84555c565f923)
	}
	__obf_55480fa872ab8d98.
		Set(__obf_a361258b272167f7)
	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_67478b6409640cb0(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_f405571f43f30f46 := __obf_c57e419440fe363b(__obf_a361258b272167f7)
	__obf_503ff9cc2619ec73 := true
	switch {
	case __obf_f405571f43f30f46 == reflect.String:
		__obf_55480fa872ab8d98.
			SetString(__obf_a361258b272167f7.String())
	case __obf_f405571f43f30f46 == reflect.Bool && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		if __obf_a361258b272167f7.Bool() {
			__obf_55480fa872ab8d98.
				SetString("1")
		} else {
			__obf_55480fa872ab8d98.
				SetString("0")
		}
	case __obf_f405571f43f30f46 == reflect.Int && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_55480fa872ab8d98.
			SetString(strconv.FormatInt(__obf_a361258b272167f7.Int(), 10))
	case __obf_f405571f43f30f46 == reflect.Uint && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_55480fa872ab8d98.
			SetString(strconv.FormatUint(__obf_a361258b272167f7.Uint(), 10))
	case __obf_f405571f43f30f46 == reflect.Float32 && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_55480fa872ab8d98.
			SetString(strconv.FormatFloat(__obf_a361258b272167f7.Float(), 'f', -1, 64))
	case __obf_f405571f43f30f46 == reflect.Slice && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput, __obf_f405571f43f30f46 ==
		reflect.Array && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_b42cbbd1bcdb3a20 := __obf_a361258b272167f7.Type()
		__obf_021d31b2022f7cf6 := __obf_b42cbbd1bcdb3a20.Elem().Kind()
		switch __obf_021d31b2022f7cf6 {
		case reflect.Uint8:
			var __obf_4f6cfb618761bc62 []uint8
			if __obf_f405571f43f30f46 == reflect.Array {
				__obf_4f6cfb618761bc62 = make([]uint8, __obf_a361258b272167f7.Len())
				for __obf_507d21ee6c71f742 := range __obf_4f6cfb618761bc62 {
					__obf_4f6cfb618761bc62[__obf_507d21ee6c71f742] = __obf_a361258b272167f7.Index(__obf_507d21ee6c71f742).Interface().(uint8)
				}
			} else {
				__obf_4f6cfb618761bc62 = __obf_a361258b272167f7.Interface().([]uint8)
			}
			__obf_55480fa872ab8d98.
				SetString(string(__obf_4f6cfb618761bc62))
		default:
			__obf_503ff9cc2619ec73 = false
		}
	default:
		__obf_503ff9cc2619ec73 = false
	}

	if !__obf_503ff9cc2619ec73 {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_a361258b272167f7.Type(), __obf_50c6df487a55f30e)
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_b89fb2c2f196ef6a(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_f405571f43f30f46 := __obf_c57e419440fe363b(__obf_a361258b272167f7)
	__obf_b42cbbd1bcdb3a20 := __obf_a361258b272167f7.Type()

	switch {
	case __obf_f405571f43f30f46 == reflect.Int:
		__obf_55480fa872ab8d98.
			SetInt(__obf_a361258b272167f7.Int())
	case __obf_f405571f43f30f46 == reflect.Uint:
		__obf_55480fa872ab8d98.
			SetInt(int64(__obf_a361258b272167f7.Uint()))
	case __obf_f405571f43f30f46 == reflect.Float32:
		__obf_55480fa872ab8d98.
			SetInt(int64(__obf_a361258b272167f7.Float()))
	case __obf_f405571f43f30f46 == reflect.Bool && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		if __obf_a361258b272167f7.Bool() {
			__obf_55480fa872ab8d98.
				SetInt(1)
		} else {
			__obf_55480fa872ab8d98.
				SetInt(0)
		}
	case __obf_f405571f43f30f46 == reflect.String && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_58fce06456cade02 := __obf_a361258b272167f7.String()
		if __obf_58fce06456cade02 == "" {
			__obf_58fce06456cade02 = "0"
		}
		__obf_507d21ee6c71f742, __obf_fb7fbd58154fa1dd := strconv.ParseInt(__obf_58fce06456cade02, 0, __obf_55480fa872ab8d98.Type().Bits())
		if __obf_fb7fbd58154fa1dd == nil {
			__obf_55480fa872ab8d98.
				SetInt(__obf_507d21ee6c71f742)
		} else {
			return fmt.Errorf("cannot parse '%s' as int: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
	case __obf_b42cbbd1bcdb3a20.PkgPath() == "encoding/json" && __obf_b42cbbd1bcdb3a20.Name() == "Number":
		__obf_7eb1e2d7b2c84418 := __obf_50c6df487a55f30e.(json.Number)
		__obf_507d21ee6c71f742, __obf_fb7fbd58154fa1dd := __obf_7eb1e2d7b2c84418.Int64()
		if __obf_fb7fbd58154fa1dd != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
		__obf_55480fa872ab8d98.
			SetInt(__obf_507d21ee6c71f742)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_a361258b272167f7.Type(), __obf_50c6df487a55f30e)
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_04fc611047ec2b16(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_f405571f43f30f46 := __obf_c57e419440fe363b(__obf_a361258b272167f7)
	__obf_b42cbbd1bcdb3a20 := __obf_a361258b272167f7.Type()

	switch {
	case __obf_f405571f43f30f46 == reflect.Int:
		__obf_507d21ee6c71f742 := __obf_a361258b272167f7.Int()
		if __obf_507d21ee6c71f742 < 0 && !__obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %d overflows uint", __obf_f36d6422fc6a0190, __obf_507d21ee6c71f742)
		}
		__obf_55480fa872ab8d98.
			SetUint(uint64(__obf_507d21ee6c71f742))
	case __obf_f405571f43f30f46 == reflect.Uint:
		__obf_55480fa872ab8d98.
			SetUint(__obf_a361258b272167f7.Uint())
	case __obf_f405571f43f30f46 == reflect.Float32:
		__obf_f500d725634477d0 := __obf_a361258b272167f7.Float()
		if __obf_f500d725634477d0 < 0 && !__obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput {
			return fmt.Errorf("cannot parse '%s', %f overflows uint", __obf_f36d6422fc6a0190, __obf_f500d725634477d0)
		}
		__obf_55480fa872ab8d98.
			SetUint(uint64(__obf_f500d725634477d0))
	case __obf_f405571f43f30f46 == reflect.Bool && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		if __obf_a361258b272167f7.Bool() {
			__obf_55480fa872ab8d98.
				SetUint(1)
		} else {
			__obf_55480fa872ab8d98.
				SetUint(0)
		}
	case __obf_f405571f43f30f46 == reflect.String && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_58fce06456cade02 := __obf_a361258b272167f7.String()
		if __obf_58fce06456cade02 == "" {
			__obf_58fce06456cade02 = "0"
		}
		__obf_507d21ee6c71f742, __obf_fb7fbd58154fa1dd := strconv.ParseUint(__obf_58fce06456cade02, 0, __obf_55480fa872ab8d98.Type().Bits())
		if __obf_fb7fbd58154fa1dd == nil {
			__obf_55480fa872ab8d98.
				SetUint(__obf_507d21ee6c71f742)
		} else {
			return fmt.Errorf("cannot parse '%s' as uint: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
	case __obf_b42cbbd1bcdb3a20.PkgPath() == "encoding/json" && __obf_b42cbbd1bcdb3a20.Name() == "Number":
		__obf_7eb1e2d7b2c84418 := __obf_50c6df487a55f30e.(json.Number)
		__obf_507d21ee6c71f742, __obf_fb7fbd58154fa1dd := strconv.ParseUint(string(__obf_7eb1e2d7b2c84418), 0, 64)
		if __obf_fb7fbd58154fa1dd != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
		__obf_55480fa872ab8d98.
			SetUint(__obf_507d21ee6c71f742)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_a361258b272167f7.Type(), __obf_50c6df487a55f30e)
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_b32c0c2e0abc7678(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_f405571f43f30f46 := __obf_c57e419440fe363b(__obf_a361258b272167f7)

	switch {
	case __obf_f405571f43f30f46 == reflect.Bool:
		__obf_55480fa872ab8d98.
			SetBool(__obf_a361258b272167f7.Bool())
	case __obf_f405571f43f30f46 == reflect.Int && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_55480fa872ab8d98.
			SetBool(__obf_a361258b272167f7.Int() != 0)
	case __obf_f405571f43f30f46 == reflect.Uint && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_55480fa872ab8d98.
			SetBool(__obf_a361258b272167f7.Uint() != 0)
	case __obf_f405571f43f30f46 == reflect.Float32 && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_55480fa872ab8d98.
			SetBool(__obf_a361258b272167f7.Float() != 0)
	case __obf_f405571f43f30f46 == reflect.String && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_fde44f2ea6c186ca, __obf_fb7fbd58154fa1dd := strconv.ParseBool(__obf_a361258b272167f7.String())
		if __obf_fb7fbd58154fa1dd == nil {
			__obf_55480fa872ab8d98.
				SetBool(__obf_fde44f2ea6c186ca)
		} else if __obf_a361258b272167f7.String() == "" {
			__obf_55480fa872ab8d98.
				SetBool(false)
		} else {
			return fmt.Errorf("cannot parse '%s' as bool: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_a361258b272167f7.Type(), __obf_50c6df487a55f30e)
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_56cc256ab94979c3(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_f405571f43f30f46 := __obf_c57e419440fe363b(__obf_a361258b272167f7)
	__obf_b42cbbd1bcdb3a20 := __obf_a361258b272167f7.Type()

	switch {
	case __obf_f405571f43f30f46 == reflect.Int:
		__obf_55480fa872ab8d98.
			SetFloat(float64(__obf_a361258b272167f7.Int()))
	case __obf_f405571f43f30f46 == reflect.Uint:
		__obf_55480fa872ab8d98.
			SetFloat(float64(__obf_a361258b272167f7.Uint()))
	case __obf_f405571f43f30f46 == reflect.Float32:
		__obf_55480fa872ab8d98.
			SetFloat(__obf_a361258b272167f7.Float())
	case __obf_f405571f43f30f46 == reflect.Bool && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		if __obf_a361258b272167f7.Bool() {
			__obf_55480fa872ab8d98.
				SetFloat(1)
		} else {
			__obf_55480fa872ab8d98.
				SetFloat(0)
		}
	case __obf_f405571f43f30f46 == reflect.String && __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput:
		__obf_58fce06456cade02 := __obf_a361258b272167f7.String()
		if __obf_58fce06456cade02 == "" {
			__obf_58fce06456cade02 = "0"
		}
		__obf_f500d725634477d0, __obf_fb7fbd58154fa1dd := strconv.ParseFloat(__obf_58fce06456cade02, __obf_55480fa872ab8d98.Type().Bits())
		if __obf_fb7fbd58154fa1dd == nil {
			__obf_55480fa872ab8d98.
				SetFloat(__obf_f500d725634477d0)
		} else {
			return fmt.Errorf("cannot parse '%s' as float: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
	case __obf_b42cbbd1bcdb3a20.PkgPath() == "encoding/json" && __obf_b42cbbd1bcdb3a20.Name() == "Number":
		__obf_7eb1e2d7b2c84418 := __obf_50c6df487a55f30e.(json.Number)
		__obf_507d21ee6c71f742, __obf_fb7fbd58154fa1dd := __obf_7eb1e2d7b2c84418.Float64()
		if __obf_fb7fbd58154fa1dd != nil {
			return fmt.Errorf(
				"error decoding json.Number into %s: %s", __obf_f36d6422fc6a0190, __obf_fb7fbd58154fa1dd)
		}
		__obf_55480fa872ab8d98.
			SetFloat(__obf_507d21ee6c71f742)
	default:
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_a361258b272167f7.Type(), __obf_50c6df487a55f30e)
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_bdc14c5e7a664ecd(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a14f0657cebf3b51 := __obf_55480fa872ab8d98.Type()
	__obf_4587c71a1da9dc23 := __obf_a14f0657cebf3b51.Key()
	__obf_ba31417ed62a42fd := __obf_a14f0657cebf3b51.Elem()
	__obf_3b26fac56296e3c2 := // By default we overwrite keys in the current map
		__obf_55480fa872ab8d98

	// If the map is nil or we're purposely zeroing fields, make a new map
	if __obf_3b26fac56296e3c2.IsNil() || __obf_1b8775533a1609ca.__obf_f7db361728796f85.ZeroFields {
		__obf_666225348aa31443 := // Make a new map to hold our result
			reflect.MapOf(__obf_4587c71a1da9dc23, __obf_ba31417ed62a42fd)
		__obf_3b26fac56296e3c2 = reflect.MakeMap(__obf_666225348aa31443)
	}
	__obf_a361258b272167f7 := // Check input type and based on the input type jump to the proper func
		reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	switch __obf_a361258b272167f7.Kind() {
	case reflect.Map:
		return __obf_1b8775533a1609ca.__obf_8257fc9915f62e56(__obf_f36d6422fc6a0190, __obf_a361258b272167f7, __obf_55480fa872ab8d98, __obf_3b26fac56296e3c2)

	case reflect.Struct:
		return __obf_1b8775533a1609ca.__obf_ae1bcf6b7b7324df(__obf_a361258b272167f7, __obf_55480fa872ab8d98, __obf_3b26fac56296e3c2)

	case reflect.Array, reflect.Slice:
		if __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput {
			return __obf_1b8775533a1609ca.__obf_67c730f6953bf6e2(__obf_f36d6422fc6a0190, __obf_a361258b272167f7, __obf_55480fa872ab8d98, __obf_3b26fac56296e3c2)
		}

		fallthrough

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_f36d6422fc6a0190, __obf_a361258b272167f7.Kind())
	}
}

func (__obf_1b8775533a1609ca *Decoder) __obf_67c730f6953bf6e2(__obf_f36d6422fc6a0190 string, __obf_a361258b272167f7 reflect.Value, __obf_55480fa872ab8d98 reflect.Value, __obf_3b26fac56296e3c2 reflect.Value) error {
	// Special case for BC reasons (covered by tests)
	if __obf_a361258b272167f7.Len() == 0 {
		__obf_55480fa872ab8d98.
			Set(__obf_3b26fac56296e3c2)
		return nil
	}

	for __obf_507d21ee6c71f742 := 0; __obf_507d21ee6c71f742 < __obf_a361258b272167f7.Len(); __obf_507d21ee6c71f742++ {
		__obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f36d6422fc6a0190+
			"["+strconv.Itoa(__obf_507d21ee6c71f742)+"]", __obf_a361258b272167f7.
			Index(__obf_507d21ee6c71f742).Interface(), __obf_55480fa872ab8d98)
		if __obf_fb7fbd58154fa1dd != nil {
			return __obf_fb7fbd58154fa1dd
		}
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_8257fc9915f62e56(__obf_f36d6422fc6a0190 string, __obf_a361258b272167f7 reflect.Value, __obf_55480fa872ab8d98 reflect.Value, __obf_3b26fac56296e3c2 reflect.Value) error {
	__obf_a14f0657cebf3b51 := __obf_55480fa872ab8d98.Type()
	__obf_4587c71a1da9dc23 := __obf_a14f0657cebf3b51.Key()
	__obf_ba31417ed62a42fd := __obf_a14f0657cebf3b51.Elem()

	// Accumulate errors
	errors := make([]string, 0)

	// If the input data is empty, then we just match what the input data is.
	if __obf_a361258b272167f7.Len() == 0 {
		if __obf_a361258b272167f7.IsNil() {
			if !__obf_55480fa872ab8d98.IsNil() {
				__obf_55480fa872ab8d98.
					Set(__obf_a361258b272167f7)
			}
		} else {
			__obf_55480fa872ab8d98.
				// Set to empty allocated value
				Set(__obf_3b26fac56296e3c2)
		}

		return nil
	}

	for _, __obf_e8ff2af7f1d900b1 := range __obf_a361258b272167f7.MapKeys() {
		__obf_f2a21f20827fddb2 := __obf_f36d6422fc6a0190 + "[" + __obf_e8ff2af7f1d900b1.String() + "]"
		__obf_2ae455e0e25917c5 := // First decode the key into the proper type
			reflect.Indirect(reflect.New(__obf_4587c71a1da9dc23))
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f2a21f20827fddb2, __obf_e8ff2af7f1d900b1.Interface(), __obf_2ae455e0e25917c5); __obf_fb7fbd58154fa1dd != nil {
			errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
			continue
		}
		__obf_71e738e554663e8c := // Next decode the data into the proper type
			__obf_a361258b272167f7.MapIndex(__obf_e8ff2af7f1d900b1).Interface()
		__obf_c439a15553595188 := reflect.Indirect(reflect.New(__obf_ba31417ed62a42fd))
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f2a21f20827fddb2, __obf_71e738e554663e8c, __obf_c439a15553595188); __obf_fb7fbd58154fa1dd != nil {
			errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
			continue
		}
		__obf_3b26fac56296e3c2.
			SetMapIndex(__obf_2ae455e0e25917c5, __obf_c439a15553595188)
	}
	__obf_55480fa872ab8d98.

		// Set the built up map to the value
		Set(__obf_3b26fac56296e3c2)

	// If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_ae1bcf6b7b7324df(__obf_a361258b272167f7 reflect.Value, __obf_55480fa872ab8d98 reflect.Value, __obf_3b26fac56296e3c2 reflect.Value) error {
	__obf_a7bb948c0a474582 := __obf_a361258b272167f7.Type()
	for __obf_507d21ee6c71f742 := 0; __obf_507d21ee6c71f742 < __obf_a7bb948c0a474582.NumField(); __obf_507d21ee6c71f742++ {
		__obf_f500d725634477d0 := // Get the StructField first since this is a cheap operation. If the
			// field is unexported, then ignore it.
			__obf_a7bb948c0a474582.Field(__obf_507d21ee6c71f742)
		if __obf_f500d725634477d0.PkgPath != "" {
			continue
		}
		__obf_71e738e554663e8c := // Next get the actual value of this field and verify it is assignable
			// to the map value.
			__obf_a361258b272167f7.Field(__obf_507d21ee6c71f742)
		if !__obf_71e738e554663e8c.Type().AssignableTo(__obf_3b26fac56296e3c2.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", __obf_71e738e554663e8c.Type(), __obf_3b26fac56296e3c2.Type().Elem())
		}
		__obf_6bd3aaffdf3c9c69 := __obf_f500d725634477d0.Tag.Get(__obf_1b8775533a1609ca.__obf_f7db361728796f85.TagName)
		__obf_033882014ae458f9 := __obf_f500d725634477d0.Name

		if __obf_6bd3aaffdf3c9c69 == "" && __obf_1b8775533a1609ca.__obf_f7db361728796f85.IgnoreUntaggedFields {
			continue
		}
		__obf_9844facec8f54006 := // If Squash is set in the config, we squash the field down.
			__obf_1b8775533a1609ca.__obf_f7db361728796f85.Squash && __obf_71e738e554663e8c.Kind() == reflect.Struct && __obf_f500d725634477d0.Anonymous
		__obf_71e738e554663e8c = __obf_fddd3ccd9d135e72(__obf_71e738e554663e8c, __obf_1b8775533a1609ca.

			// Determine the name of the key in the map
			__obf_f7db361728796f85.TagName)

		if __obf_871007343f567a0e := strings.Index(__obf_6bd3aaffdf3c9c69, ","); __obf_871007343f567a0e != -1 {
			if __obf_6bd3aaffdf3c9c69[:__obf_871007343f567a0e] == "-" {
				continue
			}
			// If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Contains(__obf_6bd3aaffdf3c9c69[__obf_871007343f567a0e+1:], "omitempty") && __obf_fc9d762a25588019(__obf_71e738e554663e8c) {
				continue
			}
			__obf_9844facec8f54006 = // If "squash" is specified in the tag, we squash the field down.
				__obf_9844facec8f54006 || strings.Contains(__obf_6bd3aaffdf3c9c69[__obf_871007343f567a0e+1:], "squash")
			if __obf_9844facec8f54006 {
				// When squashing, the embedded type can be a pointer to a struct.
				if __obf_71e738e554663e8c.Kind() == reflect.Ptr && __obf_71e738e554663e8c.Elem().Kind() == reflect.Struct {
					__obf_71e738e554663e8c = __obf_71e738e554663e8c.Elem()
				}

				// The final type must be a struct
				if __obf_71e738e554663e8c.Kind() != reflect.Struct {
					return fmt.Errorf("cannot squash non-struct type '%s'", __obf_71e738e554663e8c.Type())
				}
			}
			if __obf_3eb16aec0c925f4f := __obf_6bd3aaffdf3c9c69[:__obf_871007343f567a0e]; __obf_3eb16aec0c925f4f != "" {
				__obf_033882014ae458f9 = __obf_3eb16aec0c925f4f
			}
		} else if len(__obf_6bd3aaffdf3c9c69) > 0 {
			if __obf_6bd3aaffdf3c9c69 == "-" {
				continue
			}
			__obf_033882014ae458f9 = __obf_6bd3aaffdf3c9c69
		}

		switch __obf_71e738e554663e8c.Kind() {
		// this is an embedded struct, so handle it differently
		case reflect.Struct:
			__obf_4168ef5529c55bfb := reflect.New(__obf_71e738e554663e8c.Type())
			__obf_4168ef5529c55bfb.
				Elem().Set(__obf_71e738e554663e8c)
			__obf_16f09a2b1183057d := __obf_3b26fac56296e3c2.Type()
			__obf_6cb560e8d789c337 := __obf_16f09a2b1183057d.Key()
			__obf_8b943ddde27f1df2 := __obf_16f09a2b1183057d.Elem()
			__obf_b98d42d5029a9cb8 := reflect.MapOf(__obf_6cb560e8d789c337, __obf_8b943ddde27f1df2)
			__obf_2bf985401da8f609 := reflect.MakeMap(__obf_b98d42d5029a9cb8)
			__obf_a433332f5bcd4d2c := // Creating a pointer to a map so that other methods can completely
				// overwrite the map if need be (looking at you decodeMapFromMap). The
				// indirection allows the underlying map to be settable (CanSet() == true)
				// where as reflect.MakeMap returns an unsettable map.
				reflect.New(__obf_2bf985401da8f609.Type())
			reflect.Indirect(__obf_a433332f5bcd4d2c).Set(__obf_2bf985401da8f609)
			__obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_033882014ae458f9, __obf_4168ef5529c55bfb.Interface(), reflect.Indirect(__obf_a433332f5bcd4d2c))
			if __obf_fb7fbd58154fa1dd != nil {
				return __obf_fb7fbd58154fa1dd
			}
			__obf_2bf985401da8f609 = // the underlying map may have been completely overwritten so pull
				// it indirectly out of the enclosing value.
				reflect.Indirect(__obf_a433332f5bcd4d2c)

			if __obf_9844facec8f54006 {
				for _, __obf_e8ff2af7f1d900b1 := range __obf_2bf985401da8f609.MapKeys() {
					__obf_3b26fac56296e3c2.
						SetMapIndex(__obf_e8ff2af7f1d900b1, __obf_2bf985401da8f609.MapIndex(__obf_e8ff2af7f1d900b1))
				}
			} else {
				__obf_3b26fac56296e3c2.
					SetMapIndex(reflect.ValueOf(__obf_033882014ae458f9), __obf_2bf985401da8f609)
			}

		default:
			__obf_3b26fac56296e3c2.
				SetMapIndex(reflect.ValueOf(__obf_033882014ae458f9), __obf_71e738e554663e8c)
		}
	}

	if __obf_55480fa872ab8d98.CanAddr() {
		__obf_55480fa872ab8d98.
			Set(__obf_3b26fac56296e3c2)
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_759141d66f396441(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) (bool, error) {
	__obf_7158f164cee6cc04 := // If the input data is nil, then we want to just set the output
		// pointer to be nil as well.
		__obf_50c6df487a55f30e == nil
	if !__obf_7158f164cee6cc04 {
		switch __obf_71e738e554663e8c := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e)); __obf_71e738e554663e8c.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			__obf_7158f164cee6cc04 = __obf_71e738e554663e8c.IsNil()
		}
	}
	if __obf_7158f164cee6cc04 {
		if !__obf_55480fa872ab8d98.IsNil() && __obf_55480fa872ab8d98.CanSet() {
			__obf_6239557d33ccf5e8 := reflect.New(__obf_55480fa872ab8d98.Type()).Elem()
			__obf_55480fa872ab8d98.
				Set(__obf_6239557d33ccf5e8)
		}

		return true, nil
	}
	__obf_a14f0657cebf3b51 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		__obf_55480fa872ab8d98.Type()
	__obf_ba31417ed62a42fd := __obf_a14f0657cebf3b51.Elem()
	if __obf_55480fa872ab8d98.CanSet() {
		__obf_83034572c9c6d1b5 := __obf_55480fa872ab8d98
		if __obf_83034572c9c6d1b5.IsNil() || __obf_1b8775533a1609ca.__obf_f7db361728796f85.ZeroFields {
			__obf_83034572c9c6d1b5 = reflect.New(__obf_ba31417ed62a42fd)
		}

		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f36d6422fc6a0190, __obf_50c6df487a55f30e, reflect.Indirect(__obf_83034572c9c6d1b5)); __obf_fb7fbd58154fa1dd != nil {
			return false, __obf_fb7fbd58154fa1dd
		}
		__obf_55480fa872ab8d98.
			Set(__obf_83034572c9c6d1b5)
	} else {
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f36d6422fc6a0190, __obf_50c6df487a55f30e, reflect.Indirect(__obf_55480fa872ab8d98)); __obf_fb7fbd58154fa1dd != nil {
			return false, __obf_fb7fbd58154fa1dd
		}
	}
	return false, nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_53abac02713a0541(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := // Create an element of the concrete (non pointer) type and decode
		// into that. Then set the value of the pointer to this type.
		reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	if __obf_55480fa872ab8d98.Type() != __obf_a361258b272167f7.Type() {
		return fmt.Errorf(
			"'%s' expected type '%s', got unconvertible type '%s', value: '%v'", __obf_f36d6422fc6a0190, __obf_55480fa872ab8d98.
				Type(), __obf_a361258b272167f7.Type(), __obf_50c6df487a55f30e)
	}
	__obf_55480fa872ab8d98.
		Set(__obf_a361258b272167f7)
	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_ec5d8403f901afa0(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_3fcc5d2c7621606f := __obf_a361258b272167f7.Kind()
	__obf_a14f0657cebf3b51 := __obf_55480fa872ab8d98.Type()
	__obf_ba31417ed62a42fd := __obf_a14f0657cebf3b51.Elem()
	__obf_05e9b549e95fa26d := reflect.SliceOf(__obf_ba31417ed62a42fd)

	// If we have a non array/slice type then we first attempt to convert.
	if __obf_3fcc5d2c7621606f != reflect.Array && __obf_3fcc5d2c7621606f != reflect.Slice {
		if __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput {
			switch {
			// Slice and array we use the normal logic
			case __obf_3fcc5d2c7621606f == reflect.Slice, __obf_3fcc5d2c7621606f == reflect.Array:
				break

			// Empty maps turn into empty slices
			case __obf_3fcc5d2c7621606f == reflect.Map:
				if __obf_a361258b272167f7.Len() == 0 {
					__obf_55480fa872ab8d98.
						Set(reflect.MakeSlice(__obf_05e9b549e95fa26d, 0, 0))
					return nil
				}
				// Create slice of maps of other sizes
				return __obf_1b8775533a1609ca.__obf_ec5d8403f901afa0(__obf_f36d6422fc6a0190, []any{__obf_50c6df487a55f30e}, __obf_55480fa872ab8d98)

			case __obf_3fcc5d2c7621606f == reflect.String && __obf_ba31417ed62a42fd.Kind() == reflect.Uint8:
				return __obf_1b8775533a1609ca.__obf_ec5d8403f901afa0(__obf_f36d6422fc6a0190, []byte(__obf_a361258b272167f7.String()), __obf_55480fa872ab8d98)

			// All other types we try to convert to the slice type
			// and "lift" it into it. i.e. a string becomes a string slice.
			default:
				// Just re-try this function with data as a slice.
				return __obf_1b8775533a1609ca.__obf_ec5d8403f901afa0(__obf_f36d6422fc6a0190, []any{__obf_50c6df487a55f30e}, __obf_55480fa872ab8d98)
			}
		}

		return fmt.Errorf(
			"'%s': source data must be an array or slice, got %s", __obf_f36d6422fc6a0190, __obf_3fcc5d2c7621606f)
	}

	// If the input value is nil, then don't allocate since empty != nil
	if __obf_3fcc5d2c7621606f != reflect.Array && __obf_a361258b272167f7.IsNil() {
		return nil
	}
	__obf_ea240fefa8d3bcd5 := __obf_55480fa872ab8d98
	if __obf_ea240fefa8d3bcd5.IsNil() || __obf_1b8775533a1609ca.__obf_f7db361728796f85.ZeroFields {
		__obf_ea240fefa8d3bcd5 = // Make a new slice to hold our result, same size as the original data.
			reflect.MakeSlice(__obf_05e9b549e95fa26d, __obf_a361258b272167f7.Len(), __obf_a361258b272167f7.Len())
	} else if __obf_ea240fefa8d3bcd5.Len() > __obf_a361258b272167f7.Len() {
		__obf_ea240fefa8d3bcd5 = __obf_ea240fefa8d3bcd5.Slice(0, __obf_a361258b272167f7.Len())
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_507d21ee6c71f742 := 0; __obf_507d21ee6c71f742 < __obf_a361258b272167f7.Len(); __obf_507d21ee6c71f742++ {
		__obf_1d489860ddb2adf1 := __obf_a361258b272167f7.Index(__obf_507d21ee6c71f742).Interface()
		for __obf_ea240fefa8d3bcd5.Len() <= __obf_507d21ee6c71f742 {
			__obf_ea240fefa8d3bcd5 = reflect.Append(__obf_ea240fefa8d3bcd5, reflect.Zero(__obf_ba31417ed62a42fd))
		}
		__obf_d8a976e99ea9067c := __obf_ea240fefa8d3bcd5.Index(__obf_507d21ee6c71f742)
		__obf_f2a21f20827fddb2 := __obf_f36d6422fc6a0190 + "[" + strconv.Itoa(__obf_507d21ee6c71f742) + "]"
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f2a21f20827fddb2, __obf_1d489860ddb2adf1, __obf_d8a976e99ea9067c); __obf_fb7fbd58154fa1dd != nil {
			errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
		}
	}
	__obf_55480fa872ab8d98.

		// Finally, set the value to the slice we built up
		Set(__obf_ea240fefa8d3bcd5)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_3d3439b3dca75bf7(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))
	__obf_3fcc5d2c7621606f := __obf_a361258b272167f7.Kind()
	__obf_a14f0657cebf3b51 := __obf_55480fa872ab8d98.Type()
	__obf_ba31417ed62a42fd := __obf_a14f0657cebf3b51.Elem()
	__obf_dbed2d7cc0191ddd := reflect.ArrayOf(__obf_a14f0657cebf3b51.Len(), __obf_ba31417ed62a42fd)
	__obf_62c1958251c6dacc := __obf_55480fa872ab8d98

	if __obf_62c1958251c6dacc.Interface() == reflect.Zero(__obf_62c1958251c6dacc.Type()).Interface() || __obf_1b8775533a1609ca.__obf_f7db361728796f85.ZeroFields {
		// Check input type
		if __obf_3fcc5d2c7621606f != reflect.Array && __obf_3fcc5d2c7621606f != reflect.Slice {
			if __obf_1b8775533a1609ca.__obf_f7db361728796f85.WeaklyTypedInput {
				switch __obf_3fcc5d2c7621606f {
				// Empty maps turn into empty arrays
				case reflect.Map:
					if __obf_a361258b272167f7.Len() == 0 {
						__obf_55480fa872ab8d98.
							Set(reflect.Zero(__obf_dbed2d7cc0191ddd))
						return nil
					}

				// All other types we try to convert to the array type
				// and "lift" it into it. i.e. a string becomes a string array.
				default:
					// Just re-try this function with data as a slice.
					return __obf_1b8775533a1609ca.__obf_3d3439b3dca75bf7(__obf_f36d6422fc6a0190, []any{__obf_50c6df487a55f30e}, __obf_55480fa872ab8d98)
				}
			}

			return fmt.Errorf(
				"'%s': source data must be an array or slice, got %s", __obf_f36d6422fc6a0190, __obf_3fcc5d2c7621606f)

		}
		if __obf_a361258b272167f7.Len() > __obf_dbed2d7cc0191ddd.Len() {
			return fmt.Errorf(
				"'%s': expected source data to have length less or equal to %d, got %d", __obf_f36d6422fc6a0190, __obf_dbed2d7cc0191ddd.Len(), __obf_a361258b272167f7.Len())

		}
		__obf_62c1958251c6dacc = // Make a new array to hold our result, same size as the original data.
			reflect.New(__obf_dbed2d7cc0191ddd).Elem()
	}

	// Accumulate any errors
	errors := make([]string, 0)

	for __obf_507d21ee6c71f742 := 0; __obf_507d21ee6c71f742 < __obf_a361258b272167f7.Len(); __obf_507d21ee6c71f742++ {
		__obf_1d489860ddb2adf1 := __obf_a361258b272167f7.Index(__obf_507d21ee6c71f742).Interface()
		__obf_d8a976e99ea9067c := __obf_62c1958251c6dacc.Index(__obf_507d21ee6c71f742)
		__obf_f2a21f20827fddb2 := __obf_f36d6422fc6a0190 + "[" + strconv.Itoa(__obf_507d21ee6c71f742) + "]"
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f2a21f20827fddb2, __obf_1d489860ddb2adf1, __obf_d8a976e99ea9067c); __obf_fb7fbd58154fa1dd != nil {
			errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
		}
	}
	__obf_55480fa872ab8d98.

		// Finally, set the value to the array we built up
		Set(__obf_62c1958251c6dacc)

	// If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}

	return nil
}

func (__obf_1b8775533a1609ca *Decoder) __obf_4b5ce3ffb9b54ebf(__obf_f36d6422fc6a0190 string, __obf_50c6df487a55f30e any, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_a361258b272167f7 := reflect.Indirect(reflect.ValueOf(__obf_50c6df487a55f30e))

	// If the type of the value to write to and the data match directly,
	// then we just set it directly instead of recursing into the structure.
	if __obf_a361258b272167f7.Type() == __obf_55480fa872ab8d98.Type() {
		__obf_55480fa872ab8d98.
			Set(__obf_a361258b272167f7)
		return nil
	}
	__obf_3fcc5d2c7621606f := __obf_a361258b272167f7.Kind()
	switch __obf_3fcc5d2c7621606f {
	case reflect.Map:
		return __obf_1b8775533a1609ca.__obf_ca39dbb39c31c92e(__obf_f36d6422fc6a0190, __obf_a361258b272167f7,

			// Not the most efficient way to do this but we can optimize later if
			// we want to. To convert from struct to struct we go to map first
			// as an intermediary.
			__obf_55480fa872ab8d98)

	case reflect.Struct:
		__obf_666225348aa31443 := // Make a new map to hold our result
			reflect.TypeOf((map[string]any)(nil))
		__obf_7428fd24cd6afa71 := reflect.MakeMap(__obf_666225348aa31443)
		__obf_a433332f5bcd4d2c := // Creating a pointer to a map so that other methods can completely
			// overwrite the map if need be (looking at you decodeMapFromMap). The
			// indirection allows the underlying map to be settable (CanSet() == true)
			// where as reflect.MakeMap returns an unsettable map.
			reflect.New(__obf_7428fd24cd6afa71.Type())

		reflect.Indirect(__obf_a433332f5bcd4d2c).Set(__obf_7428fd24cd6afa71)
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_ae1bcf6b7b7324df(__obf_a361258b272167f7, reflect.Indirect(__obf_a433332f5bcd4d2c), __obf_7428fd24cd6afa71); __obf_fb7fbd58154fa1dd != nil {
			return __obf_fb7fbd58154fa1dd
		}
		__obf_47a3e13e334b0913 := __obf_1b8775533a1609ca.__obf_ca39dbb39c31c92e(__obf_f36d6422fc6a0190, reflect.Indirect(__obf_a433332f5bcd4d2c), __obf_55480fa872ab8d98)
		return __obf_47a3e13e334b0913

	default:
		return fmt.Errorf("'%s' expected a map, got '%s'", __obf_f36d6422fc6a0190, __obf_a361258b272167f7.Kind())
	}
}

func (__obf_1b8775533a1609ca *Decoder) __obf_ca39dbb39c31c92e(__obf_f36d6422fc6a0190 string, __obf_a361258b272167f7, __obf_55480fa872ab8d98 reflect.Value) error {
	__obf_97e84555c565f923 := __obf_a361258b272167f7.Type()
	if __obf_ac9697769d164007 := __obf_97e84555c565f923.Key().Kind(); __obf_ac9697769d164007 != reflect.String && __obf_ac9697769d164007 != reflect.Interface {
		return fmt.Errorf(
			"'%s' needs a map with string keys, has '%s' keys", __obf_f36d6422fc6a0190, __obf_97e84555c565f923.
				Key().Kind())
	}
	__obf_5952d8f8adb6fa34 := make(map[reflect.Value]struct{})
	__obf_6d6b00bdb5f004b7 := make(map[any]struct{})
	for _, __obf_6980dd69b7b949ff := range __obf_a361258b272167f7.MapKeys() {
		__obf_5952d8f8adb6fa34[__obf_6980dd69b7b949ff] = struct{}{}
		__obf_6d6b00bdb5f004b7[__obf_6980dd69b7b949ff.Interface()] = struct{}{}
	}
	__obf_cb86440747d83af5 := make(map[any]struct{})
	errors := make([]string, 0)
	__obf_4f2659454b8f5b55 := // This slice will keep track of all the structs we'll be decoding.
		// There can be more than one struct if there are embedded structs
		// that are squashed.
		make([]reflect.Value, 1, 5)
	__obf_4f2659454b8f5b55[0] = __obf_55480fa872ab8d98

	// Compile the list of all the fields that we're going to be decoding
	// from all the structs.
	type __obf_707f936ecb9b9556 struct {
		__obf_707f936ecb9b9556 reflect.StructField
		__obf_55480fa872ab8d98 reflect.Value
	}

	// remainField is set to a valid field set with the "remain" tag if
	// we are keeping track of remaining values.
	var __obf_ddc785350dc99652 *__obf_707f936ecb9b9556
	__obf_df402ff7a9d25ca1 := []__obf_707f936ecb9b9556{}
	for len(__obf_4f2659454b8f5b55) > 0 {
		__obf_ff24917db7b9ed9f := __obf_4f2659454b8f5b55[0]
		__obf_4f2659454b8f5b55 = __obf_4f2659454b8f5b55[1:]
		__obf_9064aa6102683e7d := __obf_ff24917db7b9ed9f.Type()

		for __obf_507d21ee6c71f742 := 0; __obf_507d21ee6c71f742 < __obf_9064aa6102683e7d.NumField(); __obf_507d21ee6c71f742++ {
			__obf_dcf6536e63dcd81c := __obf_9064aa6102683e7d.Field(__obf_507d21ee6c71f742)
			__obf_c4dacb5b95c46e31 := __obf_ff24917db7b9ed9f.Field(__obf_507d21ee6c71f742)
			if __obf_c4dacb5b95c46e31.Kind() == reflect.Pointer && __obf_c4dacb5b95c46e31.Elem().Kind() == reflect.Struct {
				__obf_c4dacb5b95c46e31 = // Handle embedded struct pointers as embedded structs.
					__obf_c4dacb5b95c46e31.Elem()
			}
			__obf_9844facec8f54006 := // If "squash" is specified in the tag, we squash the field down.
				__obf_1b8775533a1609ca.__obf_f7db361728796f85.Squash && __obf_c4dacb5b95c46e31.Kind() == reflect.Struct && __obf_dcf6536e63dcd81c.Anonymous
			__obf_d4cb6286c0e8d3c5 := false
			__obf_88aad3fb7503adc6 := // We always parse the tags cause we're looking for other tags too
				strings.Split(__obf_dcf6536e63dcd81c.Tag.Get(__obf_1b8775533a1609ca.__obf_f7db361728796f85.TagName), ",")
			for _, __obf_277942b1615c54a6 := range __obf_88aad3fb7503adc6[1:] {
				if __obf_277942b1615c54a6 == "squash" {
					__obf_9844facec8f54006 = true
					break
				}

				if __obf_277942b1615c54a6 == "remain" {
					__obf_d4cb6286c0e8d3c5 = true
					break
				}
			}

			if __obf_9844facec8f54006 {
				if __obf_c4dacb5b95c46e31.Kind() != reflect.Struct {
					errors = __obf_b59ff522c1d88f2e(errors,
						fmt.Errorf("%s: unsupported type for squash: %s", __obf_dcf6536e63dcd81c.Name, __obf_c4dacb5b95c46e31.Kind()))
				} else {
					__obf_4f2659454b8f5b55 = append(__obf_4f2659454b8f5b55, __obf_c4dacb5b95c46e31)
				}
				continue
			}

			// Build our field
			if __obf_d4cb6286c0e8d3c5 {
				__obf_ddc785350dc99652 = &__obf_707f936ecb9b9556{__obf_dcf6536e63dcd81c, __obf_c4dacb5b95c46e31}
			} else {
				__obf_df402ff7a9d25ca1 = // Normal struct field, store it away
					append(__obf_df402ff7a9d25ca1, __obf_707f936ecb9b9556{__obf_dcf6536e63dcd81c, __obf_c4dacb5b95c46e31})
			}
		}
	}

	// for fieldType, field := range fields {
	for _, __obf_f500d725634477d0 := range __obf_df402ff7a9d25ca1 {
		__obf_707f936ecb9b9556, __obf_9dd52fb3270fe2e1 := __obf_f500d725634477d0.__obf_707f936ecb9b9556, __obf_f500d725634477d0.__obf_55480fa872ab8d98
		__obf_f2a21f20827fddb2 := __obf_707f936ecb9b9556.Name
		__obf_6bd3aaffdf3c9c69 := __obf_707f936ecb9b9556.Tag.Get(__obf_1b8775533a1609ca.__obf_f7db361728796f85.TagName)
		__obf_6bd3aaffdf3c9c69 = strings.SplitN(__obf_6bd3aaffdf3c9c69, ",", 2)[0]
		if __obf_6bd3aaffdf3c9c69 != "" {
			__obf_f2a21f20827fddb2 = __obf_6bd3aaffdf3c9c69
		}
		__obf_05fb3bae393b7778 := reflect.ValueOf(__obf_f2a21f20827fddb2)
		__obf_91aff6f050f27964 := __obf_a361258b272167f7.MapIndex(__obf_05fb3bae393b7778)
		if !__obf_91aff6f050f27964.IsValid() {
			// Do a slower search by iterating over each key and
			// doing case-insensitive search.
			for __obf_6980dd69b7b949ff := range __obf_5952d8f8adb6fa34 {
				__obf_07a5fc3488fd554f, __obf_f67b17ac0536943b := __obf_6980dd69b7b949ff.Interface().(string)
				if !__obf_f67b17ac0536943b {
					// Not a string key
					continue
				}

				if __obf_1b8775533a1609ca.__obf_f7db361728796f85.MatchName(__obf_07a5fc3488fd554f, __obf_f2a21f20827fddb2) {
					__obf_05fb3bae393b7778 = __obf_6980dd69b7b949ff
					__obf_91aff6f050f27964 = __obf_a361258b272167f7.MapIndex(__obf_6980dd69b7b949ff)
					break
				}
			}

			if !__obf_91aff6f050f27964.IsValid() {
				__obf_cb86440747d83af5[ // There was no matching key in the map for the value in
				// the struct. Remember it for potential errors and metadata.
				__obf_f2a21f20827fddb2] = struct{}{}
				continue
			}
		}

		if !__obf_9dd52fb3270fe2e1.IsValid() {
			// This should never happen
			panic("field is not valid")
		}

		// If we can't set the field, then it is unexported or something,
		// and we just continue onwards.
		if !__obf_9dd52fb3270fe2e1.CanSet() {
			continue
		}

		// Delete the key we're using from the unused map so we stop tracking
		delete(__obf_6d6b00bdb5f004b7, __obf_05fb3bae393b7778.Interface())

		// If the name is empty string, then we're at the root, and we
		// don't dot-join the fields.
		if __obf_f36d6422fc6a0190 != "" {
			__obf_f2a21f20827fddb2 = __obf_f36d6422fc6a0190 + "." + __obf_f2a21f20827fddb2
		}

		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_0473c83ce7e0b238(__obf_f2a21f20827fddb2, __obf_91aff6f050f27964.Interface(), __obf_9dd52fb3270fe2e1); __obf_fb7fbd58154fa1dd != nil {
			errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
		}
	}

	// If we have a "remain"-tagged field and we have unused keys then
	// we put the unused keys directly into the remain field.
	if __obf_ddc785350dc99652 != nil && len(__obf_6d6b00bdb5f004b7) > 0 {
		__obf_d4cb6286c0e8d3c5 := // Build a map of only the unused values
			map[any]any{}
		for __obf_144bebeeff5d6ce1 := range __obf_6d6b00bdb5f004b7 {
			__obf_d4cb6286c0e8d3c5[__obf_144bebeeff5d6ce1] = __obf_a361258b272167f7.MapIndex(reflect.ValueOf(__obf_144bebeeff5d6ce1)).Interface()
		}

		// Decode it as-if we were just decoding this map onto our map.
		if __obf_fb7fbd58154fa1dd := __obf_1b8775533a1609ca.__obf_bdc14c5e7a664ecd(__obf_f36d6422fc6a0190, __obf_d4cb6286c0e8d3c5, __obf_ddc785350dc99652.__obf_55480fa872ab8d98); __obf_fb7fbd58154fa1dd != nil {
			errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
		}
		__obf_6d6b00bdb5f004b7 = // Set the map to nil so we have none so that the next check will
			// not error (ErrorUnused)
			nil
	}

	if __obf_1b8775533a1609ca.__obf_f7db361728796f85.ErrorUnused && len(__obf_6d6b00bdb5f004b7) > 0 {
		__obf_aa872e70b2c2d77d := make([]string, 0, len(__obf_6d6b00bdb5f004b7))
		for __obf_62862f30276befb9 := range __obf_6d6b00bdb5f004b7 {
			__obf_aa872e70b2c2d77d = append(__obf_aa872e70b2c2d77d, __obf_62862f30276befb9.(string))
		}
		sort.Strings(__obf_aa872e70b2c2d77d)
		__obf_fb7fbd58154fa1dd := fmt.Errorf("'%s' has invalid keys: %s", __obf_f36d6422fc6a0190, strings.Join(__obf_aa872e70b2c2d77d, ", "))
		errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
	}

	if __obf_1b8775533a1609ca.__obf_f7db361728796f85.ErrorUnset && len(__obf_cb86440747d83af5) > 0 {
		__obf_aa872e70b2c2d77d := make([]string, 0, len(__obf_cb86440747d83af5))
		for __obf_62862f30276befb9 := range __obf_cb86440747d83af5 {
			__obf_aa872e70b2c2d77d = append(__obf_aa872e70b2c2d77d, __obf_62862f30276befb9.(string))
		}
		sort.Strings(__obf_aa872e70b2c2d77d)
		__obf_fb7fbd58154fa1dd := fmt.Errorf("'%s' has unset fields: %s", __obf_f36d6422fc6a0190, strings.Join(__obf_aa872e70b2c2d77d, ", "))
		errors = __obf_b59ff522c1d88f2e(errors, __obf_fb7fbd58154fa1dd)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	// Add the unused keys to the list of unused keys if we're tracking metadata
	if __obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata != nil {
		for __obf_62862f30276befb9 := range __obf_6d6b00bdb5f004b7 {
			__obf_144bebeeff5d6ce1 := __obf_62862f30276befb9.(string)
			if __obf_f36d6422fc6a0190 != "" {
				__obf_144bebeeff5d6ce1 = __obf_f36d6422fc6a0190 + "." + __obf_144bebeeff5d6ce1
			}
			__obf_1b8775533a1609ca.__obf_f7db361728796f85.
				Metadata.Unused = append(__obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata.Unused, __obf_144bebeeff5d6ce1)
		}
		for __obf_62862f30276befb9 := range __obf_cb86440747d83af5 {
			__obf_144bebeeff5d6ce1 := __obf_62862f30276befb9.(string)
			if __obf_f36d6422fc6a0190 != "" {
				__obf_144bebeeff5d6ce1 = __obf_f36d6422fc6a0190 + "." + __obf_144bebeeff5d6ce1
			}
			__obf_1b8775533a1609ca.__obf_f7db361728796f85.
				Metadata.Unset = append(__obf_1b8775533a1609ca.__obf_f7db361728796f85.Metadata.Unset, __obf_144bebeeff5d6ce1)
		}
	}

	return nil
}

func __obf_fc9d762a25588019(__obf_71e738e554663e8c reflect.Value) bool {
	switch __obf_c57e419440fe363b(__obf_71e738e554663e8c) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return __obf_71e738e554663e8c.Len() == 0
	case reflect.Bool:
		return !__obf_71e738e554663e8c.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return __obf_71e738e554663e8c.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return __obf_71e738e554663e8c.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return __obf_71e738e554663e8c.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return __obf_71e738e554663e8c.IsNil()
	}
	return false
}

func __obf_c57e419440fe363b(__obf_55480fa872ab8d98 reflect.Value) reflect.Kind {
	__obf_ac9697769d164007 := __obf_55480fa872ab8d98.Kind()

	switch {
	case __obf_ac9697769d164007 >= reflect.Int && __obf_ac9697769d164007 <= reflect.Int64:
		return reflect.Int
	case __obf_ac9697769d164007 >= reflect.Uint && __obf_ac9697769d164007 <= reflect.Uint64:
		return reflect.Uint
	case __obf_ac9697769d164007 >= reflect.Float32 && __obf_ac9697769d164007 <= reflect.Float64:
		return reflect.Float32
	default:
		return __obf_ac9697769d164007
	}
}

func __obf_19d3c3518f9d3fd3(__obf_a7bb948c0a474582 reflect.Type, __obf_9095b82010c0295e bool, __obf_82f020389dfc6ec8 string) bool {
	for __obf_507d21ee6c71f742 := 0; __obf_507d21ee6c71f742 < __obf_a7bb948c0a474582.NumField(); __obf_507d21ee6c71f742++ {
		__obf_f500d725634477d0 := __obf_a7bb948c0a474582.Field(__obf_507d21ee6c71f742)
		if __obf_f500d725634477d0.PkgPath == "" && !__obf_9095b82010c0295e { // check for unexported fields
			return true
		}
		if __obf_9095b82010c0295e && __obf_f500d725634477d0.Tag.Get(__obf_82f020389dfc6ec8) != "" { // check for mapstructure tags inside
			return true
		}
	}
	return false
}

func __obf_fddd3ccd9d135e72(__obf_71e738e554663e8c reflect.Value, __obf_82f020389dfc6ec8 string) reflect.Value {
	if __obf_71e738e554663e8c.Kind() != reflect.Ptr || __obf_71e738e554663e8c.Elem().Kind() != reflect.Struct {
		return __obf_71e738e554663e8c
	}
	__obf_4ae05966d2a2621b := __obf_71e738e554663e8c.Elem()
	__obf_4d0409c55ef213a4 := __obf_4ae05966d2a2621b.Type()
	if __obf_19d3c3518f9d3fd3(__obf_4d0409c55ef213a4, true, __obf_82f020389dfc6ec8) {
		return __obf_4ae05966d2a2621b
	}
	return __obf_71e738e554663e8c
}
