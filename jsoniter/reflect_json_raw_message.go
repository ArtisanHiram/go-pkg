package __obf_5b802ce8d9ba56d6

import (
	"encoding/json"
	"unsafe"

	"github.com/modern-go/reflect2"
)

var __obf_018c7f794e00198f = reflect2.TypeOfPtr((*json.RawMessage)(nil)).Elem()
var __obf_a4955bf424b9dc55 = reflect2.TypeOfPtr((*RawMessage)(nil)).Elem()

func __obf_248bb80e8824b081(__obf_5efc66d8c338c133 reflect2.Type) ValEncoder {
	if __obf_5efc66d8c338c133 == __obf_018c7f794e00198f {
		return &__obf_9a9138574410794d{}
	}
	if __obf_5efc66d8c338c133 == __obf_a4955bf424b9dc55 {
		return &__obf_6c7be86530d92b73{}
	}
	return nil
}

func __obf_e31e45c2077f2473(__obf_5efc66d8c338c133 reflect2.Type) ValDecoder {
	if __obf_5efc66d8c338c133 == __obf_018c7f794e00198f {
		return &__obf_9a9138574410794d{}
	}
	if __obf_5efc66d8c338c133 == __obf_a4955bf424b9dc55 {
		return &__obf_6c7be86530d92b73{}
	}
	return nil
}

type __obf_9a9138574410794d struct {
}

func (__obf_0099bc81efe356aa *__obf_9a9138574410794d) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.ReadNil() {
		*((*json.RawMessage)(__obf_d3c919547bf7e47a)) = nil
	} else {
		*((*json.RawMessage)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.SkipAndReturnBytes()
	}
}

func (__obf_0099bc81efe356aa *__obf_9a9138574410794d) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if *((*json.RawMessage)(__obf_d3c919547bf7e47a)) == nil {
		__obf_00fc491caa967a74.
			WriteNil()
	} else {
		__obf_00fc491caa967a74.
			WriteRaw(string(*((*json.RawMessage)(__obf_d3c919547bf7e47a))))
	}
}

func (__obf_0099bc81efe356aa *__obf_9a9138574410794d) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return len(*((*json.RawMessage)(__obf_d3c919547bf7e47a))) == 0
}

type __obf_6c7be86530d92b73 struct {
}

func (__obf_0099bc81efe356aa *__obf_6c7be86530d92b73) Decode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_67008a6a9e5ba828 *Iterator) {
	if __obf_67008a6a9e5ba828.ReadNil() {
		*((*RawMessage)(__obf_d3c919547bf7e47a)) = nil
	} else {
		*((*RawMessage)(__obf_d3c919547bf7e47a)) = __obf_67008a6a9e5ba828.SkipAndReturnBytes()
	}
}

func (__obf_0099bc81efe356aa *__obf_6c7be86530d92b73) Encode(__obf_d3c919547bf7e47a unsafe.Pointer, __obf_00fc491caa967a74 *Stream) {
	if *((*RawMessage)(__obf_d3c919547bf7e47a)) == nil {
		__obf_00fc491caa967a74.
			WriteNil()
	} else {
		__obf_00fc491caa967a74.
			WriteRaw(string(*((*RawMessage)(__obf_d3c919547bf7e47a))))
	}
}

func (__obf_0099bc81efe356aa *__obf_6c7be86530d92b73) IsEmpty(__obf_d3c919547bf7e47a unsafe.Pointer) bool {
	return len(*((*RawMessage)(__obf_d3c919547bf7e47a))) == 0
}
