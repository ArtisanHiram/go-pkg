package __obf_38f1d2f091ad74e0

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"time"
	"unsafe"
)

// RegisterTimeAsInt64Codec encode/decode time since number of unit since epoch. the precision is the unit.
func RegisterTimeAsInt64Codec(__obf_1840c122265d7096 time.Duration) {
	jsoniter.RegisterTypeEncoder("time.Time", &__obf_1108cec4560439e4{__obf_1840c122265d7096})
	jsoniter.RegisterTypeDecoder("time.Time", &__obf_1108cec4560439e4{__obf_1840c122265d7096})
}

type __obf_1108cec4560439e4 struct {
	__obf_1840c122265d7096 time.Duration
}

func (__obf_af0a31f95254d8e7 *__obf_1108cec4560439e4) Decode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_113f80f39cc94185 *jsoniter.Iterator) {
	__obf_5edfd6d3446170ab := __obf_113f80f39cc94185.ReadInt64() * __obf_af0a31f95254d8e7.__obf_1840c122265d7096.Nanoseconds()
	*((*time.Time)(__obf_35567cf7daf6e12d)) = time.Unix(0, __obf_5edfd6d3446170ab)
}

func (__obf_af0a31f95254d8e7 *__obf_1108cec4560439e4) IsEmpty(__obf_35567cf7daf6e12d unsafe.Pointer) bool {
	__obf_0dcdf726eaf380c7 := *((*time.Time)(__obf_35567cf7daf6e12d))
	return __obf_0dcdf726eaf380c7.UnixNano() == 0
}
func (__obf_af0a31f95254d8e7 *__obf_1108cec4560439e4) Encode(__obf_35567cf7daf6e12d unsafe.Pointer, __obf_27f0100519cc4eeb *jsoniter.Stream) {
	__obf_0dcdf726eaf380c7 := *((*time.Time)(__obf_35567cf7daf6e12d))
	__obf_27f0100519cc4eeb.
		WriteInt64(__obf_0dcdf726eaf380c7.UnixNano() / __obf_af0a31f95254d8e7.__obf_1840c122265d7096.Nanoseconds())
}
