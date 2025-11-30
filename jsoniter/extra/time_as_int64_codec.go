package __obf_9a397ef96534ad45

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"time"
	"unsafe"
)

// RegisterTimeAsInt64Codec encode/decode time since number of unit since epoch. the precision is the unit.
func RegisterTimeAsInt64Codec(__obf_e5b1bb3f3d4c2831 time.Duration) {
	jsoniter.RegisterTypeEncoder("time.Time", &__obf_dca5baa5fd1c59c2{__obf_e5b1bb3f3d4c2831})
	jsoniter.RegisterTypeDecoder("time.Time", &__obf_dca5baa5fd1c59c2{__obf_e5b1bb3f3d4c2831})
}

type __obf_dca5baa5fd1c59c2 struct {
	__obf_e5b1bb3f3d4c2831 time.Duration
}

func (__obf_d92fa10e8df80d33 *__obf_dca5baa5fd1c59c2) Decode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_f2099f19d22a8175 *jsoniter.Iterator) {
	__obf_c7b6fd571be56d0d := __obf_f2099f19d22a8175.ReadInt64() * __obf_d92fa10e8df80d33.__obf_e5b1bb3f3d4c2831.Nanoseconds()
	*((*time.Time)(__obf_77e2593d34cc3a6a)) = time.Unix(0, __obf_c7b6fd571be56d0d)
}

func (__obf_d92fa10e8df80d33 *__obf_dca5baa5fd1c59c2) IsEmpty(__obf_77e2593d34cc3a6a unsafe.Pointer) bool {
	__obf_bec20bf2067d0605 := *((*time.Time)(__obf_77e2593d34cc3a6a))
	return __obf_bec20bf2067d0605.UnixNano() == 0
}
func (__obf_d92fa10e8df80d33 *__obf_dca5baa5fd1c59c2) Encode(__obf_77e2593d34cc3a6a unsafe.Pointer, __obf_3b68bc41f53ae503 *jsoniter.Stream) {
	__obf_bec20bf2067d0605 := *((*time.Time)(__obf_77e2593d34cc3a6a))
	__obf_3b68bc41f53ae503.
		WriteInt64(__obf_bec20bf2067d0605.UnixNano() / __obf_d92fa10e8df80d33.__obf_e5b1bb3f3d4c2831.Nanoseconds())
}
