package __obf_1f22c6b8dfc77bff

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"time"
	"unsafe"
)

// RegisterTimeAsInt64Codec encode/decode time since number of unit since epoch. the precision is the unit.
func RegisterTimeAsInt64Codec(__obf_1203f39a948316e7 time.Duration) {
	jsoniter.RegisterTypeEncoder("time.Time", &__obf_7f0f82b84eb16c80{__obf_1203f39a948316e7})
	jsoniter.RegisterTypeDecoder("time.Time", &__obf_7f0f82b84eb16c80{__obf_1203f39a948316e7})
}

type __obf_7f0f82b84eb16c80 struct {
	__obf_1203f39a948316e7 time.Duration
}

func (__obf_fd6e92bcd33a0284 *__obf_7f0f82b84eb16c80) Decode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d021dab62946a708 *jsoniter.Iterator) {
	__obf_1541bc3b2d17c5da := __obf_d021dab62946a708.ReadInt64() * __obf_fd6e92bcd33a0284.__obf_1203f39a948316e7.Nanoseconds()
	*((*time.Time)(__obf_a5271c65f4ae17af)) = time.Unix(0, __obf_1541bc3b2d17c5da)
}

func (__obf_fd6e92bcd33a0284 *__obf_7f0f82b84eb16c80) IsEmpty(__obf_a5271c65f4ae17af unsafe.Pointer) bool {
	__obf_9f6949bedf709f9f := *((*time.Time)(__obf_a5271c65f4ae17af))
	return __obf_9f6949bedf709f9f.UnixNano() == 0
}
func (__obf_fd6e92bcd33a0284 *__obf_7f0f82b84eb16c80) Encode(__obf_a5271c65f4ae17af unsafe.Pointer, __obf_d178d558227696d6 *jsoniter.Stream) {
	__obf_9f6949bedf709f9f := *((*time.Time)(__obf_a5271c65f4ae17af))
	__obf_d178d558227696d6.
		WriteInt64(__obf_9f6949bedf709f9f.UnixNano() / __obf_fd6e92bcd33a0284.__obf_1203f39a948316e7.Nanoseconds())
}
