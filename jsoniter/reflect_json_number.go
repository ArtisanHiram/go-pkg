package __obf_703d23ba520c3295

import (
	"encoding/json"
	"strconv"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type Number string

// String returns the literal text of the number.
func (__obf_716729e0f8e419ed Number) String() string { return string(__obf_716729e0f8e419ed) }

// Float64 returns the number as a float64.
func (__obf_716729e0f8e419ed Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(__obf_716729e0f8e419ed), 64)
}

// Int64 returns the number as an int64.
func (__obf_716729e0f8e419ed Number) Int64() (int64, error) {
	return strconv.ParseInt(string(__obf_716729e0f8e419ed), 10, 64)
}

func CastJsonNumber(__obf_b924538fffe5fd64 any) (string, bool) {
	switch __obf_31335a1173d75d14 := __obf_b924538fffe5fd64.(type) {
	case json.Number:
		return string(__obf_31335a1173d75d14), true
	case Number:
		return string(__obf_31335a1173d75d14), true
	}
	return "", false
}

var __obf_f551a08aeabc9b95 = reflect2.TypeOfPtr((*json.Number)(nil)).Elem()
var __obf_6792c8460020b31d = reflect2.TypeOfPtr((*Number)(nil)).Elem()

func __obf_1cd9fe2978e7054f(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	if __obf_3b7f6abbae19451e.AssignableTo(__obf_f551a08aeabc9b95) {
		return &__obf_419b09df1a1c7c5e{}
	}
	if __obf_3b7f6abbae19451e.AssignableTo(__obf_6792c8460020b31d) {
		return &__obf_b697c6d4862bac42{}
	}
	return nil
}

func __obf_0c0040a8d86c5f59(__obf_2aaf7367de3ff86d *__obf_2aaf7367de3ff86d, __obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	if __obf_3b7f6abbae19451e.AssignableTo(__obf_f551a08aeabc9b95) {
		return &__obf_419b09df1a1c7c5e{}
	}
	if __obf_3b7f6abbae19451e.AssignableTo(__obf_6792c8460020b31d) {
		return &__obf_b697c6d4862bac42{}
	}
	return nil
}

type __obf_419b09df1a1c7c5e struct {
}

func (__obf_4acad06eb5535907 *__obf_419b09df1a1c7c5e) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	switch __obf_47edab4c16a2d88d.WhatIsNext() {
	case StringValue:
		*((*json.Number)(__obf_47fa53f3d161f45c)) = json.Number(__obf_47edab4c16a2d88d.ReadString())
	case NilValue:
		__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('n', 'u', 'l', 'l')
		*((*json.Number)(__obf_47fa53f3d161f45c)) = ""
	default:
		*((*json.Number)(__obf_47fa53f3d161f45c)) = json.Number([]byte(__obf_47edab4c16a2d88d.__obf_4c009361bc8be406()))
	}
}

func (__obf_4acad06eb5535907 *__obf_419b09df1a1c7c5e) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_11a8c2d8ba07392d := *((*json.Number)(__obf_47fa53f3d161f45c))
	if len(__obf_11a8c2d8ba07392d) == 0 {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('0')
	} else {
		__obf_9a34f62857fb1d1d.
			WriteRaw(string(__obf_11a8c2d8ba07392d))
	}
}

func (__obf_4acad06eb5535907 *__obf_419b09df1a1c7c5e) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return len(*((*json.Number)(__obf_47fa53f3d161f45c))) == 0
}

type __obf_b697c6d4862bac42 struct {
}

func (__obf_4acad06eb5535907 *__obf_b697c6d4862bac42) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	switch __obf_47edab4c16a2d88d.WhatIsNext() {
	case StringValue:
		*((*Number)(__obf_47fa53f3d161f45c)) = Number(__obf_47edab4c16a2d88d.ReadString())
	case NilValue:
		__obf_47edab4c16a2d88d.__obf_6f584222681dcca0('n', 'u', 'l', 'l')
		*((*Number)(__obf_47fa53f3d161f45c)) = ""
	default:
		*((*Number)(__obf_47fa53f3d161f45c)) = Number([]byte(__obf_47edab4c16a2d88d.__obf_4c009361bc8be406()))
	}
}

func (__obf_4acad06eb5535907 *__obf_b697c6d4862bac42) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	__obf_11a8c2d8ba07392d := *((*Number)(__obf_47fa53f3d161f45c))
	if len(__obf_11a8c2d8ba07392d) == 0 {
		__obf_9a34f62857fb1d1d.__obf_f7df2a5224462d19('0')
	} else {
		__obf_9a34f62857fb1d1d.
			WriteRaw(string(__obf_11a8c2d8ba07392d))
	}
}

func (__obf_4acad06eb5535907 *__obf_b697c6d4862bac42) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return len(*((*Number)(__obf_47fa53f3d161f45c))) == 0
}
