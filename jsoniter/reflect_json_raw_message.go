package __obf_703d23ba520c3295

import (
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_b000159c6e1a3f5c = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var __obf_54f479acf58fff71 = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func __obf_f3e98e16ceefbfa6(__obf_3b7f6abbae19451e reflect2.Type) ValEncoder {
	if __obf_3b7f6abbae19451e == __obf_b000159c6e1a3f5c {
		return &__obf_b24f337e8db41021{}
	}
	if __obf_3b7f6abbae19451e == __obf_54f479acf58fff71 {
		return &__obf_b1668a87826ed134{}
	}
	return nil
}

func __obf_8dbf41d7a8e2a43b(__obf_3b7f6abbae19451e reflect2.Type) ValDecoder {
	if __obf_3b7f6abbae19451e == __obf_b000159c6e1a3f5c {
		return &__obf_b24f337e8db41021{}
	}
	if __obf_3b7f6abbae19451e == __obf_54f479acf58fff71 {
		return &__obf_b1668a87826ed134{}
	}
	return nil
}

type __obf_b24f337e8db41021 struct {
}

func (__obf_4acad06eb5535907 *__obf_b24f337e8db41021) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.ReadNil() {
		*((*json.RawMessage)(__obf_47fa53f3d161f45c)) = nil
	} else {
		*((*json.RawMessage)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.SkipAndReturnBytes()
	}
}

func (__obf_4acad06eb5535907 *__obf_b24f337e8db41021) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if *((*json.RawMessage)(__obf_47fa53f3d161f45c)) == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
	} else {
		__obf_9a34f62857fb1d1d.
			WriteRaw(string(*((*json.RawMessage)(__obf_47fa53f3d161f45c))))
	}
}

func (__obf_4acad06eb5535907 *__obf_b24f337e8db41021) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(__obf_47fa53f3d161f45c))) == 0
}

type __obf_b1668a87826ed134 struct {
}

func (__obf_4acad06eb5535907 *__obf_b1668a87826ed134) Decode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_47edab4c16a2d88d *Iterator) {
	if __obf_47edab4c16a2d88d.ReadNil() {
		*((*RawMessage)(__obf_47fa53f3d161f45c)) = nil
	} else {
		*((*RawMessage)(__obf_47fa53f3d161f45c)) = __obf_47edab4c16a2d88d.SkipAndReturnBytes()
	}
}

func (__obf_4acad06eb5535907 *__obf_b1668a87826ed134) Encode(__obf_47fa53f3d161f45c unsafe.Pointer, __obf_9a34f62857fb1d1d *Stream) {
	if *((*RawMessage)(__obf_47fa53f3d161f45c)) == nil {
		__obf_9a34f62857fb1d1d.
			WriteNil()
	} else {
		__obf_9a34f62857fb1d1d.
			WriteRaw(string(*((*RawMessage)(__obf_47fa53f3d161f45c))))
	}
}

func (__obf_4acad06eb5535907 *__obf_b1668a87826ed134) IsEmpty(__obf_47fa53f3d161f45c unsafe.Pointer) bool {
	return len(*((*RawMessage)(__obf_47fa53f3d161f45c))) == 0
}
