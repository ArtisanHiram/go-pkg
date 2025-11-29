package __obf_789dc33d47f4ab2c

import (
	jsoniter "github.com/ArtisanHiram/go-pkg/jsoniter"
	"time"
	"unsafe"
)

// RegisterTimeAsInt64Codec encode/decode time since number of unit since epoch. the precision is the unit.
func RegisterTimeAsInt64Codec(__obf_d8b9e40b7f731d13 time.Duration) {
	jsoniter.RegisterTypeEncoder("time.Time", &__obf_3b825ce76f454dd4{__obf_d8b9e40b7f731d13})
	jsoniter.RegisterTypeDecoder("time.Time", &__obf_3b825ce76f454dd4{__obf_d8b9e40b7f731d13})
}

type __obf_3b825ce76f454dd4 struct {
	__obf_d8b9e40b7f731d13 time.Duration
}

func (__obf_d02baddfcf2572b4 *__obf_3b825ce76f454dd4) Decode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_e97e2ea2e3d1b0a2 *jsoniter.Iterator) {
	__obf_17f4f084caef3691 := __obf_e97e2ea2e3d1b0a2.ReadInt64() * __obf_d02baddfcf2572b4.__obf_d8b9e40b7f731d13.Nanoseconds()
	*((*time.Time)(__obf_4280d1a4ac42e464)) = time.Unix(0, __obf_17f4f084caef3691)
}

func (__obf_d02baddfcf2572b4 *__obf_3b825ce76f454dd4) IsEmpty(__obf_4280d1a4ac42e464 unsafe.Pointer) bool {
	__obf_eeac1575cde01ea6 := *((*time.Time)(__obf_4280d1a4ac42e464))
	return __obf_eeac1575cde01ea6.UnixNano() == 0
}
func (__obf_d02baddfcf2572b4 *__obf_3b825ce76f454dd4) Encode(__obf_4280d1a4ac42e464 unsafe.Pointer, __obf_fe6de34d139b90e2 *jsoniter.Stream) {
	__obf_eeac1575cde01ea6 := *((*time.Time)(__obf_4280d1a4ac42e464))
	__obf_fe6de34d139b90e2.
		WriteInt64(__obf_eeac1575cde01ea6.UnixNano() / __obf_d02baddfcf2572b4.__obf_d8b9e40b7f731d13.Nanoseconds())
}
