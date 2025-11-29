package __obf_060749efdc6ad522

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"time"
	"unsafe"
)

// RegisterTimeAsInt64Codec encode/decode time since number of unit since epoch. the precision is the unit.
func RegisterTimeAsInt64Codec(__obf_84f237ca9b11406a time.Duration) {
	jsoniter.RegisterTypeEncoder("time.Time", &__obf_34cd72723da42fe2{__obf_84f237ca9b11406a})
	jsoniter.RegisterTypeDecoder("time.Time", &__obf_34cd72723da42fe2{__obf_84f237ca9b11406a})
}

type __obf_34cd72723da42fe2 struct {
	__obf_84f237ca9b11406a time.Duration
}

func (__obf_5b4176d65d17beeb *__obf_34cd72723da42fe2) Decode(__obf_deff95ef54922509 unsafe.Pointer, __obf_7eafecc31dc47100 *jsoniter.Iterator) {
	__obf_13905604eb38164e := __obf_7eafecc31dc47100.ReadInt64() * __obf_5b4176d65d17beeb.__obf_84f237ca9b11406a.Nanoseconds()
	*((*time.Time)(__obf_deff95ef54922509)) = time.Unix(0, __obf_13905604eb38164e)
}

func (__obf_5b4176d65d17beeb *__obf_34cd72723da42fe2) IsEmpty(__obf_deff95ef54922509 unsafe.Pointer) bool {
	__obf_2854ca25952b60e0 := *((*time.Time)(__obf_deff95ef54922509))
	return __obf_2854ca25952b60e0.UnixNano() == 0
}
func (__obf_5b4176d65d17beeb *__obf_34cd72723da42fe2) Encode(__obf_deff95ef54922509 unsafe.Pointer, __obf_d9afe3b0234a705f *jsoniter.Stream) {
	__obf_2854ca25952b60e0 := *((*time.Time)(__obf_deff95ef54922509))
	__obf_d9afe3b0234a705f.
		WriteInt64(__obf_2854ca25952b60e0.UnixNano() / __obf_5b4176d65d17beeb.__obf_84f237ca9b11406a.Nanoseconds())
}
