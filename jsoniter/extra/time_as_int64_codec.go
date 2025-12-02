package __obf_eed9c5a643743c33

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"time"
	"unsafe"
)

// RegisterTimeAsInt64Codec encode/decode time since number of unit since epoch. the precision is the unit.
func RegisterTimeAsInt64Codec(__obf_ef8b6d1d2137b069 time.Duration) {
	jsoniter.RegisterTypeEncoder("time.Time", &__obf_4a792a261e45d925{__obf_ef8b6d1d2137b069})
	jsoniter.RegisterTypeDecoder("time.Time", &__obf_4a792a261e45d925{__obf_ef8b6d1d2137b069})
}

type __obf_4a792a261e45d925 struct {
	__obf_ef8b6d1d2137b069 time.Duration
}

func (__obf_ccd1a335201fc2ad *__obf_4a792a261e45d925) Decode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_338eca60ccc15d82 *jsoniter.Iterator) {
	__obf_2fbe99b69404672f := __obf_338eca60ccc15d82.ReadInt64() * __obf_ccd1a335201fc2ad.__obf_ef8b6d1d2137b069.Nanoseconds()
	*((*time.Time)(__obf_e88eec03e62b77ec)) = time.Unix(0, __obf_2fbe99b69404672f)
}

func (__obf_ccd1a335201fc2ad *__obf_4a792a261e45d925) IsEmpty(__obf_e88eec03e62b77ec unsafe.Pointer) bool {
	__obf_a5f0b01ad4ee5d7c := *((*time.Time)(__obf_e88eec03e62b77ec))
	return __obf_a5f0b01ad4ee5d7c.UnixNano() == 0
}
func (__obf_ccd1a335201fc2ad *__obf_4a792a261e45d925) Encode(__obf_e88eec03e62b77ec unsafe.Pointer, __obf_837be8d71267e11f *jsoniter.Stream) {
	__obf_a5f0b01ad4ee5d7c := *((*time.Time)(__obf_e88eec03e62b77ec))
	__obf_837be8d71267e11f.
		WriteInt64(__obf_a5f0b01ad4ee5d7c.UnixNano() / __obf_ccd1a335201fc2ad.__obf_ef8b6d1d2137b069.Nanoseconds())
}
