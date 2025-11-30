package __obf_c3cd04a15c56f32f

import (
	"reflect"
	"unsafe"

	"github.com/modern-go/reflect2"
)

type __obf_4fb4b91e56916c0b struct {
	__obf_797d4fe23b3556f8 reflect2.Type
}

func (__obf_c85b895560db528f *__obf_4fb4b91e56916c0b) Encode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_2361f5214e84c60f *Stream) {
	__obf_50acae8c0a871ba1 := __obf_c85b895560db528f.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	__obf_2361f5214e84c60f.
		WriteVal(__obf_50acae8c0a871ba1)
}

func (__obf_c85b895560db528f *__obf_4fb4b91e56916c0b) IsEmpty(__obf_753ab3fb72cdd659 unsafe.Pointer) bool {
	return __obf_c85b895560db528f.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659) == nil
}

type __obf_a3ca22e5223e9e8a struct {
}

func (__obf_924cc7ef585cdfb0 *__obf_a3ca22e5223e9e8a) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	__obf_3f4d6a194c2964c8 := (*any)(__obf_753ab3fb72cdd659)
	__obf_50acae8c0a871ba1 := *__obf_3f4d6a194c2964c8
	if __obf_50acae8c0a871ba1 == nil {
		*__obf_3f4d6a194c2964c8 = __obf_edd9fe48d39445e4.Read()
		return
	}
	__obf_8667d4fc2a95b0d7 := reflect2.TypeOf(__obf_50acae8c0a871ba1)
	if __obf_8667d4fc2a95b0d7.Kind() != reflect.Ptr {
		*__obf_3f4d6a194c2964c8 = __obf_edd9fe48d39445e4.Read()
		return
	}
	__obf_9fd3068ebc25d7f9 := __obf_8667d4fc2a95b0d7.(*reflect2.UnsafePtrType)
	__obf_f858051022f2edb3 := __obf_9fd3068ebc25d7f9.Elem()
	if __obf_edd9fe48d39445e4.WhatIsNext() == NilValue {
		if __obf_f858051022f2edb3.Kind() != reflect.Ptr {
			__obf_edd9fe48d39445e4.__obf_f22f308da0537336('n', 'u', 'l', 'l')
			*__obf_3f4d6a194c2964c8 = nil
			return
		}
	}
	if reflect2.IsNil(__obf_50acae8c0a871ba1) {
		__obf_50acae8c0a871ba1 := __obf_f858051022f2edb3.New()
		__obf_edd9fe48d39445e4.
			ReadVal(__obf_50acae8c0a871ba1)
		*__obf_3f4d6a194c2964c8 = __obf_50acae8c0a871ba1
		return
	}
	__obf_edd9fe48d39445e4.
		ReadVal(__obf_50acae8c0a871ba1)
}

type __obf_59cd180ea7ef10b5 struct {
	__obf_797d4fe23b3556f8 *reflect2.UnsafeIFaceType
}

func (__obf_924cc7ef585cdfb0 *__obf_59cd180ea7ef10b5) Decode(__obf_753ab3fb72cdd659 unsafe.Pointer, __obf_edd9fe48d39445e4 *Iterator) {
	if __obf_edd9fe48d39445e4.ReadNil() {
		__obf_924cc7ef585cdfb0.__obf_797d4fe23b3556f8.
			UnsafeSet(__obf_753ab3fb72cdd659, __obf_924cc7ef585cdfb0.__obf_797d4fe23b3556f8.UnsafeNew())
		return
	}
	__obf_50acae8c0a871ba1 := __obf_924cc7ef585cdfb0.__obf_797d4fe23b3556f8.UnsafeIndirect(__obf_753ab3fb72cdd659)
	if reflect2.IsNil(__obf_50acae8c0a871ba1) {
		__obf_edd9fe48d39445e4.
			ReportError("decode non empty interface", "can not unmarshal into nil")
		return
	}
	__obf_edd9fe48d39445e4.
		ReadVal(__obf_50acae8c0a871ba1)
}
