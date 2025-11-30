package __obf_3e0c303610a19bd4

import (
	"encoding"
	"errors"
	"fmt"
	"reflect"
)

var (
	__obf_164ae3a742d185b4 = reflect.TypeOf((*any)(nil)).Elem()
	__obf_24e46306ee5355ff = reflect.TypeOf((*string)(nil)).Elem()
	__obf_132fa72af3bf692d = reflect.TypeOf((*bool)(nil)).Elem()
)

var __obf_9fc30a23a7edeb3f []__obf_299b7590962e0960

//nolint:gochecknoinits
func init() {
	__obf_9fc30a23a7edeb3f = []__obf_299b7590962e0960{
		reflect.Bool:          __obf_7718d3de15e25e53,
		reflect.Int:           __obf_2e7a83f8fd124752,
		reflect.Int8:          __obf_2e7a83f8fd124752,
		reflect.Int16:         __obf_2e7a83f8fd124752,
		reflect.Int32:         __obf_2e7a83f8fd124752,
		reflect.Int64:         __obf_2e7a83f8fd124752,
		reflect.Uint:          __obf_f1c563c28223f805,
		reflect.Uint8:         __obf_f1c563c28223f805,
		reflect.Uint16:        __obf_f1c563c28223f805,
		reflect.Uint32:        __obf_f1c563c28223f805,
		reflect.Uint64:        __obf_f1c563c28223f805,
		reflect.Float32:       __obf_a0985953c5ea413f,
		reflect.Float64:       __obf_f008a12b21b1cd5f,
		reflect.Complex64:     __obf_182af43a792864b6,
		reflect.Complex128:    __obf_182af43a792864b6,
		reflect.Array:         __obf_99c138aaf5d3c77b,
		reflect.Chan:          __obf_182af43a792864b6,
		reflect.Func:          __obf_182af43a792864b6,
		reflect.Interface:     __obf_c0b16016f640c84c,
		reflect.Map:           __obf_c02d6c3c3de4ca55,
		reflect.Pointer:       __obf_182af43a792864b6,
		reflect.Slice:         __obf_8047bcf5dd14321e,
		reflect.String:        __obf_442ca6b4ba0753ac,
		reflect.Struct:        __obf_58fb4ae094054808,
		reflect.UnsafePointer: __obf_182af43a792864b6,
	}
}

func __obf_9b0fb8e0beb3ab15(__obf_616f98efc80197c6 reflect.Type) __obf_299b7590962e0960 {
	if __obf_63bbcee86d44fdde, __obf_ce8bef141eff8aab := __obf_a46bdc105884db8b.Load(__obf_616f98efc80197c6); __obf_ce8bef141eff8aab {
		return __obf_63bbcee86d44fdde.(__obf_299b7590962e0960)
	}
	__obf_6ff63d98a176dcfe := _getDecoder(__obf_616f98efc80197c6)
	__obf_a46bdc105884db8b.
		Store(__obf_616f98efc80197c6, __obf_6ff63d98a176dcfe)
	return __obf_6ff63d98a176dcfe
}

func _getDecoder(__obf_616f98efc80197c6 reflect.Type) __obf_299b7590962e0960 {
	__obf_018b20441ecc0301 := __obf_616f98efc80197c6.Kind()

	if __obf_018b20441ecc0301 == reflect.Ptr {
		if _, __obf_ce8bef141eff8aab := __obf_a46bdc105884db8b.Load(__obf_616f98efc80197c6.Elem()); __obf_ce8bef141eff8aab {
			return __obf_e39784ebd6f6c508(__obf_616f98efc80197c6)
		}
	}

	if __obf_616f98efc80197c6.Implements(__obf_1067551650cbff5c) {
		return __obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_39ecdcaeb0565bb6)
	}
	if __obf_616f98efc80197c6.Implements(__obf_65915279442e2d73) {
		return __obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_e7b149e85ccef2bf)
	}
	if __obf_616f98efc80197c6.Implements(__obf_03e2d73a2d650e68) {
		return __obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_14968680e5eec3a1)
	}
	if __obf_616f98efc80197c6.Implements(__obf_e682a601b05d7802) {
		return __obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_f21bbb6f7de15fab)
	}

	// Addressable struct field value.
	if __obf_018b20441ecc0301 != reflect.Pointer {
		__obf_b5a4664807537c0d := reflect.PointerTo(__obf_616f98efc80197c6)
		if __obf_b5a4664807537c0d.Implements(__obf_1067551650cbff5c) {
			return __obf_ebe2c4e32d0fd746(__obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_39ecdcaeb0565bb6))
		}
		if __obf_b5a4664807537c0d.Implements(__obf_65915279442e2d73) {
			return __obf_ebe2c4e32d0fd746(__obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_e7b149e85ccef2bf))
		}
		if __obf_b5a4664807537c0d.Implements(__obf_03e2d73a2d650e68) {
			return __obf_ebe2c4e32d0fd746(__obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_14968680e5eec3a1))
		}
		if __obf_b5a4664807537c0d.Implements(__obf_e682a601b05d7802) {
			return __obf_ebe2c4e32d0fd746(__obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6, __obf_f21bbb6f7de15fab))
		}
	}

	switch __obf_018b20441ecc0301 {
	case reflect.Pointer:
		return __obf_e39784ebd6f6c508(__obf_616f98efc80197c6)
	case reflect.Slice:
		__obf_265b8ddf380ff77a := __obf_616f98efc80197c6.Elem()
		if __obf_265b8ddf380ff77a.Kind() == reflect.Uint8 {
			return __obf_79dff4d84a62412e
		}
		if __obf_265b8ddf380ff77a == __obf_24e46306ee5355ff {
			return __obf_f8b216cbef75a778
		}
	case reflect.Array:
		if __obf_616f98efc80197c6.Elem().Kind() == reflect.Uint8 {
			return __obf_2866c43e524b97a1
		}
	case reflect.Map:
		if __obf_616f98efc80197c6.Key() == __obf_24e46306ee5355ff {
			switch __obf_616f98efc80197c6.Elem() {
			case __obf_24e46306ee5355ff:
				return __obf_3a3b656fd33bb5aa
			case __obf_164ae3a742d185b4:
				return __obf_ab027dc2684d6371
			}
		}
	}

	return __obf_9fc30a23a7edeb3f[__obf_018b20441ecc0301]
}

func __obf_e39784ebd6f6c508(__obf_616f98efc80197c6 reflect.Type) __obf_299b7590962e0960 {
	__obf_91fd0d587e6af119 := __obf_9b0fb8e0beb3ab15(__obf_616f98efc80197c6.Elem())
	return func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if __obf_dc35117108ba8439.__obf_56c3d1b3779316ba() {
			if !__obf_63bbcee86d44fdde.IsNil() {
				__obf_63bbcee86d44fdde.
					Set(__obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_616f98efc80197c6).Elem())
			}
			return __obf_dc35117108ba8439.DecodeNil()
		}
		if __obf_63bbcee86d44fdde.IsNil() {
			__obf_63bbcee86d44fdde.
				Set(__obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_616f98efc80197c6.Elem()))
		}
		return __obf_91fd0d587e6af119(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde.Elem())
	}
}

func __obf_ebe2c4e32d0fd746(__obf_6ff63d98a176dcfe __obf_299b7590962e0960) __obf_299b7590962e0960 {
	return func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if !__obf_63bbcee86d44fdde.CanAddr() {
			return fmt.Errorf("msgpack: Decode(nonaddressable %T)", __obf_63bbcee86d44fdde.Interface())
		}
		return __obf_6ff63d98a176dcfe(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde.Addr())
	}
}

func __obf_c4a9ac4b65666ad0(__obf_616f98efc80197c6 reflect.Type, __obf_6ff63d98a176dcfe __obf_299b7590962e0960) __obf_299b7590962e0960 {
	if __obf_ebd8855a2f3055ad(__obf_616f98efc80197c6.Kind()) {
		return func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
			if __obf_dc35117108ba8439.__obf_56c3d1b3779316ba() {
				return __obf_dc35117108ba8439.__obf_f8b6b5cb00b4ac2c(__obf_63bbcee86d44fdde)
			}
			if __obf_63bbcee86d44fdde.IsNil() {
				__obf_63bbcee86d44fdde.
					Set(__obf_dc35117108ba8439.__obf_94f521ecbd0b4afa(__obf_616f98efc80197c6.Elem()))
			}
			return __obf_6ff63d98a176dcfe(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde)
		}
	}

	return func(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
		if __obf_dc35117108ba8439.__obf_56c3d1b3779316ba() {
			return __obf_dc35117108ba8439.__obf_f8b6b5cb00b4ac2c(__obf_63bbcee86d44fdde)
		}
		return __obf_6ff63d98a176dcfe(__obf_dc35117108ba8439, __obf_63bbcee86d44fdde)
	}
}

func __obf_7718d3de15e25e53(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_3e30e51d97f3d287, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeBool()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_63bbcee86d44fdde.
		SetBool(__obf_3e30e51d97f3d287)
	return nil
}

func __obf_c0b16016f640c84c(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	if __obf_63bbcee86d44fdde.IsNil() {
		return __obf_dc35117108ba8439.__obf_fa3509ec2ed568f3(__obf_63bbcee86d44fdde)
	}
	return __obf_dc35117108ba8439.DecodeValue(__obf_63bbcee86d44fdde.Elem())
}

func (__obf_dc35117108ba8439 *Decoder) __obf_fa3509ec2ed568f3(__obf_63bbcee86d44fdde reflect.Value) error {
	__obf_785552c2ee57ae56, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	if __obf_785552c2ee57ae56 != nil {
		if __obf_63bbcee86d44fdde.Type() == __obf_319e36d3ccef195a {
			if __obf_785552c2ee57ae56, __obf_ce8bef141eff8aab := __obf_785552c2ee57ae56.(string); __obf_ce8bef141eff8aab {
				__obf_63bbcee86d44fdde.
					Set(reflect.ValueOf(errors.New(__obf_785552c2ee57ae56)))
				return nil
			}
		}
		__obf_63bbcee86d44fdde.
			Set(reflect.ValueOf(__obf_785552c2ee57ae56))
	}

	return nil
}

func __obf_182af43a792864b6(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	return fmt.Errorf("msgpack: Decode(unsupported %s)", __obf_63bbcee86d44fdde.Type())
}

//------------------------------------------------------------------------------

func __obf_39ecdcaeb0565bb6(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_91fd0d587e6af119 := __obf_63bbcee86d44fdde.Interface().(CustomDecoder)
	return __obf_91fd0d587e6af119.DecodeMsgpack(__obf_dc35117108ba8439)
}

func __obf_e7b149e85ccef2bf(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	var __obf_11bcc66cde095c11 []byte
	__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = make([]byte, 0, 64)
	if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_11bcc66cde095c11 = __obf_dc35117108ba8439.__obf_4903e7701e6854a2
	__obf_dc35117108ba8439.__obf_4903e7701e6854a2 = nil
	__obf_c458cfc89120e8f9 := __obf_63bbcee86d44fdde.Interface().(Unmarshaler)
	return __obf_c458cfc89120e8f9.UnmarshalMsgpack(__obf_11bcc66cde095c11)
}

func __obf_14968680e5eec3a1(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4d827abff69b9b40, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeBytes()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_c458cfc89120e8f9 := __obf_63bbcee86d44fdde.Interface().(encoding.BinaryUnmarshaler)
	return __obf_c458cfc89120e8f9.UnmarshalBinary(__obf_4d827abff69b9b40)
}

func __obf_f21bbb6f7de15fab(__obf_dc35117108ba8439 *Decoder, __obf_63bbcee86d44fdde reflect.Value) error {
	__obf_4d827abff69b9b40, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeBytes()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	__obf_c458cfc89120e8f9 := __obf_63bbcee86d44fdde.Interface().(encoding.TextUnmarshaler)
	return __obf_c458cfc89120e8f9.UnmarshalText(__obf_4d827abff69b9b40)
}
