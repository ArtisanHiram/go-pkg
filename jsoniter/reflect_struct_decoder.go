package __obf_91620b895eeff9ed

import (
	"fmt"
	"io"
	"strings"
	"unsafe"

	"github.com/modern-go/reflect2"
)

func __obf_5e7b461732099c43(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type) ValDecoder {
	__obf_884ece42b376f6b2 := map[string]*Binding{}
	__obf_e9a2f5e72a0f0289 := __obf_d2a025550d51a745(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea)
	for _, __obf_9a21bdc8399ba183 := range __obf_e9a2f5e72a0f0289.Fields {
		for _, __obf_29f4cf19cdaefa05 := range __obf_9a21bdc8399ba183.FromNames {
			__obf_6d85186cba056552 := __obf_884ece42b376f6b2[__obf_29f4cf19cdaefa05]
			if __obf_6d85186cba056552 == nil {
				__obf_884ece42b376f6b2[__obf_29f4cf19cdaefa05] = __obf_9a21bdc8399ba183
				continue
			}
			__obf_f771f31cc06d5af5, __obf_685e2d4ce26180b4 := __obf_7edb6a45370750ef(__obf_2f9c5aed866cce75.__obf_972c2293804d141c, __obf_6d85186cba056552, __obf_9a21bdc8399ba183)
			if __obf_f771f31cc06d5af5 {
				delete(__obf_884ece42b376f6b2, __obf_29f4cf19cdaefa05)
			}
			if !__obf_685e2d4ce26180b4 {
				__obf_884ece42b376f6b2[__obf_29f4cf19cdaefa05] = __obf_9a21bdc8399ba183
			}
		}
	}
	__obf_26c5a8c732658174 := map[string]*__obf_e06930e4e8a0db9e{}
	for __obf_4b23cd971318cca4, __obf_9a21bdc8399ba183 := range __obf_884ece42b376f6b2 {
		__obf_26c5a8c732658174[__obf_4b23cd971318cca4] = __obf_9a21bdc8399ba183.Decoder.(*__obf_e06930e4e8a0db9e)
	}

	if !__obf_2f9c5aed866cce75.__obf_f087a7a47617f72a() {
		for __obf_4b23cd971318cca4, __obf_9a21bdc8399ba183 := range __obf_884ece42b376f6b2 {
			if _, __obf_93f189bffabaf5a4 := __obf_26c5a8c732658174[strings.ToLower(__obf_4b23cd971318cca4)]; !__obf_93f189bffabaf5a4 {
				__obf_26c5a8c732658174[strings.ToLower(__obf_4b23cd971318cca4)] = __obf_9a21bdc8399ba183.Decoder.(*__obf_e06930e4e8a0db9e)
			}
		}
	}

	return __obf_575896a2aa387ab1(__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174)
}

func __obf_575896a2aa387ab1(__obf_2f9c5aed866cce75 *__obf_2f9c5aed866cce75, __obf_29ebd0f2c324f5ea reflect2.Type, __obf_26c5a8c732658174 map[string]*__obf_e06930e4e8a0db9e) ValDecoder {
	if __obf_2f9c5aed866cce75.__obf_8d8c301f6a725e61 {
		return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea: __obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174: __obf_26c5a8c732658174, __obf_8d8c301f6a725e61: true}
	}
	__obf_6c453ca5a12fc6a5 := map[int64]struct{}{
		0: {},
	}

	switch len(__obf_26c5a8c732658174) {
	case 0:
		return &__obf_42572ed5e5be3c21{__obf_29ebd0f2c324f5ea}
	case 1:
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			return &__obf_4695e674b81abc4a{__obf_29ebd0f2c324f5ea, __obf_0af4d01ff4db9010, __obf_6cd5b2ea3acb38a5}
		}
	case 2:
		var __obf_e0875af5717c60d7 int64
		var __obf_10666ea1c0dcd334 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_e0875af5717c60d7 == 0 {
				__obf_e0875af5717c60d7 = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_10666ea1c0dcd334 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_a07ad08177b01d0c{__obf_29ebd0f2c324f5ea, __obf_e0875af5717c60d7, __obf_6b07ae4fb77efb32, __obf_10666ea1c0dcd334, __obf_f7b6365c11e4d197}
	case 3:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_307bbd0e08d44978{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28}
	case 4:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_e89e44302550e0fa{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e}
	case 5:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_77c10ea86cb185b8 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		var __obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else if __obf_2885e2628b11fe1c == 0 {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_77c10ea86cb185b8 = __obf_0af4d01ff4db9010
				__obf_250d1e4bb4eab1eb = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_3c022c3f11df0e49{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e, __obf_77c10ea86cb185b8, __obf_250d1e4bb4eab1eb}
	case 6:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_77c10ea86cb185b8 int64
		var __obf_eb9ec0b6ba2c1ad8 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		var __obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
		var __obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else if __obf_2885e2628b11fe1c == 0 {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			} else if __obf_77c10ea86cb185b8 == 0 {
				__obf_77c10ea86cb185b8 = __obf_0af4d01ff4db9010
				__obf_250d1e4bb4eab1eb = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_eb9ec0b6ba2c1ad8 = __obf_0af4d01ff4db9010
				__obf_75ee90d83729278b = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_57593e55083ce530{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e, __obf_77c10ea86cb185b8, __obf_250d1e4bb4eab1eb, __obf_eb9ec0b6ba2c1ad8, __obf_75ee90d83729278b}
	case 7:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_77c10ea86cb185b8 int64
		var __obf_eb9ec0b6ba2c1ad8 int64
		var __obf_156b7ae78b272dd8 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		var __obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
		var __obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
		var __obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else if __obf_2885e2628b11fe1c == 0 {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			} else if __obf_77c10ea86cb185b8 == 0 {
				__obf_77c10ea86cb185b8 = __obf_0af4d01ff4db9010
				__obf_250d1e4bb4eab1eb = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb9ec0b6ba2c1ad8 == 0 {
				__obf_eb9ec0b6ba2c1ad8 = __obf_0af4d01ff4db9010
				__obf_75ee90d83729278b = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_156b7ae78b272dd8 = __obf_0af4d01ff4db9010
				__obf_60444e029f463ce4 = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_2f0db7f8abbb6c8a{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e, __obf_77c10ea86cb185b8, __obf_250d1e4bb4eab1eb, __obf_eb9ec0b6ba2c1ad8, __obf_75ee90d83729278b, __obf_156b7ae78b272dd8, __obf_60444e029f463ce4}
	case 8:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_77c10ea86cb185b8 int64
		var __obf_eb9ec0b6ba2c1ad8 int64
		var __obf_156b7ae78b272dd8 int64
		var __obf_c44454301d9798a7 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		var __obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
		var __obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
		var __obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
		var __obf_fea513943ca1a17c *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else if __obf_2885e2628b11fe1c == 0 {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			} else if __obf_77c10ea86cb185b8 == 0 {
				__obf_77c10ea86cb185b8 = __obf_0af4d01ff4db9010
				__obf_250d1e4bb4eab1eb = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb9ec0b6ba2c1ad8 == 0 {
				__obf_eb9ec0b6ba2c1ad8 = __obf_0af4d01ff4db9010
				__obf_75ee90d83729278b = __obf_6cd5b2ea3acb38a5
			} else if __obf_156b7ae78b272dd8 == 0 {
				__obf_156b7ae78b272dd8 = __obf_0af4d01ff4db9010
				__obf_60444e029f463ce4 = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_c44454301d9798a7 = __obf_0af4d01ff4db9010
				__obf_fea513943ca1a17c = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_b6d4c9de59cb77ad{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e, __obf_77c10ea86cb185b8, __obf_250d1e4bb4eab1eb, __obf_eb9ec0b6ba2c1ad8, __obf_75ee90d83729278b, __obf_156b7ae78b272dd8, __obf_60444e029f463ce4, __obf_c44454301d9798a7, __obf_fea513943ca1a17c}
	case 9:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_77c10ea86cb185b8 int64
		var __obf_eb9ec0b6ba2c1ad8 int64
		var __obf_156b7ae78b272dd8 int64
		var __obf_c44454301d9798a7 int64
		var __obf_0b3f0647d28282e9 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		var __obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
		var __obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
		var __obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
		var __obf_fea513943ca1a17c *__obf_e06930e4e8a0db9e
		var __obf_22621062d10dbb68 *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else if __obf_2885e2628b11fe1c == 0 {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			} else if __obf_77c10ea86cb185b8 == 0 {
				__obf_77c10ea86cb185b8 = __obf_0af4d01ff4db9010
				__obf_250d1e4bb4eab1eb = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb9ec0b6ba2c1ad8 == 0 {
				__obf_eb9ec0b6ba2c1ad8 = __obf_0af4d01ff4db9010
				__obf_75ee90d83729278b = __obf_6cd5b2ea3acb38a5
			} else if __obf_156b7ae78b272dd8 == 0 {
				__obf_156b7ae78b272dd8 = __obf_0af4d01ff4db9010
				__obf_60444e029f463ce4 = __obf_6cd5b2ea3acb38a5
			} else if __obf_c44454301d9798a7 == 0 {
				__obf_c44454301d9798a7 = __obf_0af4d01ff4db9010
				__obf_fea513943ca1a17c = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_0b3f0647d28282e9 = __obf_0af4d01ff4db9010
				__obf_22621062d10dbb68 = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_31919df64b2f584d{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e, __obf_77c10ea86cb185b8, __obf_250d1e4bb4eab1eb, __obf_eb9ec0b6ba2c1ad8, __obf_75ee90d83729278b, __obf_156b7ae78b272dd8, __obf_60444e029f463ce4, __obf_c44454301d9798a7, __obf_fea513943ca1a17c, __obf_0b3f0647d28282e9, __obf_22621062d10dbb68}
	case 10:
		var __obf_809a40fccabfb8ec int64
		var __obf_eb33942c43de96d9 int64
		var __obf_0b47c440287a56f5 int64
		var __obf_2885e2628b11fe1c int64
		var __obf_77c10ea86cb185b8 int64
		var __obf_eb9ec0b6ba2c1ad8 int64
		var __obf_156b7ae78b272dd8 int64
		var __obf_c44454301d9798a7 int64
		var __obf_0b3f0647d28282e9 int64
		var __obf_0474620b886dcc85 int64
		var __obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
		var __obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
		var __obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
		var __obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
		var __obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
		var __obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
		var __obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
		var __obf_fea513943ca1a17c *__obf_e06930e4e8a0db9e
		var __obf_22621062d10dbb68 *__obf_e06930e4e8a0db9e
		var __obf_bc4e59499e3bf5bb *__obf_e06930e4e8a0db9e
		for __obf_04fb01d7218c6db3, __obf_6cd5b2ea3acb38a5 := range __obf_26c5a8c732658174 {
			__obf_0af4d01ff4db9010 := __obf_a34bad55ba1d026b(__obf_04fb01d7218c6db3, __obf_2f9c5aed866cce75.__obf_f087a7a47617f72a())
			_, __obf_b6a1e6d3e95b1e15 := __obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010]
			if __obf_b6a1e6d3e95b1e15 {
				return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
			}
			__obf_6c453ca5a12fc6a5[__obf_0af4d01ff4db9010] = struct{}{}
			if __obf_809a40fccabfb8ec == 0 {
				__obf_809a40fccabfb8ec = __obf_0af4d01ff4db9010
				__obf_6b07ae4fb77efb32 = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb33942c43de96d9 == 0 {
				__obf_eb33942c43de96d9 = __obf_0af4d01ff4db9010
				__obf_f7b6365c11e4d197 = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b47c440287a56f5 == 0 {
				__obf_0b47c440287a56f5 = __obf_0af4d01ff4db9010
				__obf_400b51e13f592f28 = __obf_6cd5b2ea3acb38a5
			} else if __obf_2885e2628b11fe1c == 0 {
				__obf_2885e2628b11fe1c = __obf_0af4d01ff4db9010
				__obf_0e1ecaedc845f11e = __obf_6cd5b2ea3acb38a5
			} else if __obf_77c10ea86cb185b8 == 0 {
				__obf_77c10ea86cb185b8 = __obf_0af4d01ff4db9010
				__obf_250d1e4bb4eab1eb = __obf_6cd5b2ea3acb38a5
			} else if __obf_eb9ec0b6ba2c1ad8 == 0 {
				__obf_eb9ec0b6ba2c1ad8 = __obf_0af4d01ff4db9010
				__obf_75ee90d83729278b = __obf_6cd5b2ea3acb38a5
			} else if __obf_156b7ae78b272dd8 == 0 {
				__obf_156b7ae78b272dd8 = __obf_0af4d01ff4db9010
				__obf_60444e029f463ce4 = __obf_6cd5b2ea3acb38a5
			} else if __obf_c44454301d9798a7 == 0 {
				__obf_c44454301d9798a7 = __obf_0af4d01ff4db9010
				__obf_fea513943ca1a17c = __obf_6cd5b2ea3acb38a5
			} else if __obf_0b3f0647d28282e9 == 0 {
				__obf_0b3f0647d28282e9 = __obf_0af4d01ff4db9010
				__obf_22621062d10dbb68 = __obf_6cd5b2ea3acb38a5
			} else {
				__obf_0474620b886dcc85 = __obf_0af4d01ff4db9010
				__obf_bc4e59499e3bf5bb = __obf_6cd5b2ea3acb38a5
			}
		}
		return &__obf_d1f166b2cc45a239{__obf_29ebd0f2c324f5ea, __obf_809a40fccabfb8ec, __obf_6b07ae4fb77efb32, __obf_eb33942c43de96d9, __obf_f7b6365c11e4d197, __obf_0b47c440287a56f5, __obf_400b51e13f592f28, __obf_2885e2628b11fe1c, __obf_0e1ecaedc845f11e, __obf_77c10ea86cb185b8, __obf_250d1e4bb4eab1eb, __obf_eb9ec0b6ba2c1ad8, __obf_75ee90d83729278b, __obf_156b7ae78b272dd8, __obf_60444e029f463ce4, __obf_c44454301d9798a7, __obf_fea513943ca1a17c, __obf_0b3f0647d28282e9, __obf_22621062d10dbb68, __obf_0474620b886dcc85, __obf_bc4e59499e3bf5bb}
	}
	return &__obf_f1e70b086792a05d{__obf_29ebd0f2c324f5ea, __obf_26c5a8c732658174, false}
}

type __obf_f1e70b086792a05d struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_26c5a8c732658174 map[string]*__obf_e06930e4e8a0db9e
	__obf_8d8c301f6a725e61 bool
}

func (__obf_6fd3f72eb9b5574c *__obf_f1e70b086792a05d) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	var __obf_f16b4157911bc9af byte
	for __obf_f16b4157911bc9af = ','; __obf_f16b4157911bc9af == ','; __obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049() {
		__obf_6fd3f72eb9b5574c.__obf_490c181ea5f2a995(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	if __obf_f16b4157911bc9af != '}' {
		__obf_1bb30e8a74ed8233.
			ReportError("struct Decode", `expect }, but found `+string([]byte{__obf_f16b4157911bc9af}))
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

func (__obf_6fd3f72eb9b5574c *__obf_f1e70b086792a05d) __obf_490c181ea5f2a995(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	var __obf_7e01b5b4651074d4 string
	var __obf_6cd5b2ea3acb38a5 *__obf_e06930e4e8a0db9e
	if __obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_065d67ebe135fdba {
		__obf_add7536e8e3ffe0e := __obf_1bb30e8a74ed8233.ReadStringAsSlice()
		__obf_7e01b5b4651074d4 = *(*string)(unsafe.Pointer(&__obf_add7536e8e3ffe0e))
		__obf_6cd5b2ea3acb38a5 = __obf_6fd3f72eb9b5574c.__obf_26c5a8c732658174[__obf_7e01b5b4651074d4]
		if __obf_6cd5b2ea3acb38a5 == nil && !__obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_f087a7a47617f72a {
			__obf_6cd5b2ea3acb38a5 = __obf_6fd3f72eb9b5574c.__obf_26c5a8c732658174[strings.ToLower(__obf_7e01b5b4651074d4)]
		}
	} else {
		__obf_7e01b5b4651074d4 = __obf_1bb30e8a74ed8233.ReadString()
		__obf_6cd5b2ea3acb38a5 = __obf_6fd3f72eb9b5574c.__obf_26c5a8c732658174[__obf_7e01b5b4651074d4]
		if __obf_6cd5b2ea3acb38a5 == nil && !__obf_1bb30e8a74ed8233.__obf_892632c77e859037.__obf_f087a7a47617f72a {
			__obf_6cd5b2ea3acb38a5 = __obf_6fd3f72eb9b5574c.__obf_26c5a8c732658174[strings.ToLower(__obf_7e01b5b4651074d4)]
		}
	}
	if __obf_6cd5b2ea3acb38a5 == nil {
		if __obf_6fd3f72eb9b5574c.__obf_8d8c301f6a725e61 {
			__obf_6457f9c7d28f805d := "found unknown field: " + __obf_7e01b5b4651074d4
			__obf_1bb30e8a74ed8233.
				ReportError("ReadObject", __obf_6457f9c7d28f805d)
		}
		__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
		if __obf_f16b4157911bc9af != ':' {
			__obf_1bb30e8a74ed8233.
				ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
		}
		__obf_1bb30e8a74ed8233.
			Skip()
		return
	}
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af != ':' {
		__obf_1bb30e8a74ed8233.
			ReportError("ReadObject", "expect : after object field, but found "+string([]byte{__obf_f16b4157911bc9af}))
	}
	__obf_6cd5b2ea3acb38a5.
		Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
}

type __obf_42572ed5e5be3c21 struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
}

func (__obf_6fd3f72eb9b5574c *__obf_42572ed5e5be3c21) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_f47701e28c5effac := __obf_1bb30e8a74ed8233.WhatIsNext()
	if __obf_f47701e28c5effac != ObjectValue && __obf_f47701e28c5effac != NilValue {
		__obf_1bb30e8a74ed8233.
			ReportError("skipObjectDecoder", "expect object or null")
		return
	}
	__obf_1bb30e8a74ed8233.
		Skip()
}

type __obf_4695e674b81abc4a struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_0af4d01ff4db9010 int64
	__obf_6cd5b2ea3acb38a5 *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_4695e674b81abc4a) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		if __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() == __obf_6fd3f72eb9b5574c.__obf_0af4d01ff4db9010 {
			__obf_6fd3f72eb9b5574c.__obf_6cd5b2ea3acb38a5.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		} else {
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_a07ad08177b01d0c struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_a07ad08177b01d0c) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_307bbd0e08d44978 struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_307bbd0e08d44978) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_e89e44302550e0fa struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_e89e44302550e0fa) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_3c022c3f11df0e49 struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
	__obf_a04f8c1ddec4b285 int64
	__obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_3c022c3f11df0e49) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_a04f8c1ddec4b285:
			__obf_6fd3f72eb9b5574c.__obf_250d1e4bb4eab1eb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_57593e55083ce530 struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
	__obf_a04f8c1ddec4b285 int64
	__obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
	__obf_b0baf9cb361044f0 int64
	__obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_57593e55083ce530) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_a04f8c1ddec4b285:
			__obf_6fd3f72eb9b5574c.__obf_250d1e4bb4eab1eb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_b0baf9cb361044f0:
			__obf_6fd3f72eb9b5574c.__obf_75ee90d83729278b.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_2f0db7f8abbb6c8a struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
	__obf_a04f8c1ddec4b285 int64
	__obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
	__obf_b0baf9cb361044f0 int64
	__obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
	__obf_5a3ba66978ce0652 int64
	__obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_2f0db7f8abbb6c8a) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_a04f8c1ddec4b285:
			__obf_6fd3f72eb9b5574c.__obf_250d1e4bb4eab1eb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_b0baf9cb361044f0:
			__obf_6fd3f72eb9b5574c.__obf_75ee90d83729278b.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_5a3ba66978ce0652:
			__obf_6fd3f72eb9b5574c.__obf_60444e029f463ce4.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_b6d4c9de59cb77ad struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
	__obf_a04f8c1ddec4b285 int64
	__obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
	__obf_b0baf9cb361044f0 int64
	__obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
	__obf_5a3ba66978ce0652 int64
	__obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
	__obf_6be0b208c8d8ee33 int64
	__obf_fea513943ca1a17c *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_b6d4c9de59cb77ad) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_a04f8c1ddec4b285:
			__obf_6fd3f72eb9b5574c.__obf_250d1e4bb4eab1eb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_b0baf9cb361044f0:
			__obf_6fd3f72eb9b5574c.__obf_75ee90d83729278b.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_5a3ba66978ce0652:
			__obf_6fd3f72eb9b5574c.__obf_60444e029f463ce4.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_6be0b208c8d8ee33:
			__obf_6fd3f72eb9b5574c.__obf_fea513943ca1a17c.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_31919df64b2f584d struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
	__obf_a04f8c1ddec4b285 int64
	__obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
	__obf_b0baf9cb361044f0 int64
	__obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
	__obf_5a3ba66978ce0652 int64
	__obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
	__obf_6be0b208c8d8ee33 int64
	__obf_fea513943ca1a17c *__obf_e06930e4e8a0db9e
	__obf_9513c678a7f56a64 int64
	__obf_22621062d10dbb68 *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_31919df64b2f584d) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_a04f8c1ddec4b285:
			__obf_6fd3f72eb9b5574c.__obf_250d1e4bb4eab1eb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_b0baf9cb361044f0:
			__obf_6fd3f72eb9b5574c.__obf_75ee90d83729278b.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_5a3ba66978ce0652:
			__obf_6fd3f72eb9b5574c.__obf_60444e029f463ce4.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_6be0b208c8d8ee33:
			__obf_6fd3f72eb9b5574c.__obf_fea513943ca1a17c.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_9513c678a7f56a64:
			__obf_6fd3f72eb9b5574c.__obf_22621062d10dbb68.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_d1f166b2cc45a239 struct {
	__obf_29ebd0f2c324f5ea reflect2.Type
	__obf_e0875af5717c60d7 int64
	__obf_6b07ae4fb77efb32 *__obf_e06930e4e8a0db9e
	__obf_10666ea1c0dcd334 int64
	__obf_f7b6365c11e4d197 *__obf_e06930e4e8a0db9e
	__obf_584ca0e5ae1b90c6 int64
	__obf_400b51e13f592f28 *__obf_e06930e4e8a0db9e
	__obf_adc4855b094e0184 int64
	__obf_0e1ecaedc845f11e *__obf_e06930e4e8a0db9e
	__obf_a04f8c1ddec4b285 int64
	__obf_250d1e4bb4eab1eb *__obf_e06930e4e8a0db9e
	__obf_b0baf9cb361044f0 int64
	__obf_75ee90d83729278b *__obf_e06930e4e8a0db9e
	__obf_5a3ba66978ce0652 int64
	__obf_60444e029f463ce4 *__obf_e06930e4e8a0db9e
	__obf_6be0b208c8d8ee33 int64
	__obf_fea513943ca1a17c *__obf_e06930e4e8a0db9e
	__obf_9513c678a7f56a64 int64
	__obf_22621062d10dbb68 *__obf_e06930e4e8a0db9e
	__obf_8d7c2a74d36313ad int64
	__obf_bc4e59499e3bf5bb *__obf_e06930e4e8a0db9e
}

func (__obf_6fd3f72eb9b5574c *__obf_d1f166b2cc45a239) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if !__obf_1bb30e8a74ed8233.__obf_613c71685c28c50b() {
		return
	}
	if !__obf_1bb30e8a74ed8233.__obf_4361ed1341bd481b() {
		return
	}
	for {
		switch __obf_1bb30e8a74ed8233.__obf_d848eeb634118020() {
		case __obf_6fd3f72eb9b5574c.__obf_e0875af5717c60d7:
			__obf_6fd3f72eb9b5574c.__obf_6b07ae4fb77efb32.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_10666ea1c0dcd334:
			__obf_6fd3f72eb9b5574c.__obf_f7b6365c11e4d197.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_584ca0e5ae1b90c6:
			__obf_6fd3f72eb9b5574c.__obf_400b51e13f592f28.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_adc4855b094e0184:
			__obf_6fd3f72eb9b5574c.__obf_0e1ecaedc845f11e.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_a04f8c1ddec4b285:
			__obf_6fd3f72eb9b5574c.__obf_250d1e4bb4eab1eb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_b0baf9cb361044f0:
			__obf_6fd3f72eb9b5574c.__obf_75ee90d83729278b.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_5a3ba66978ce0652:
			__obf_6fd3f72eb9b5574c.__obf_60444e029f463ce4.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_6be0b208c8d8ee33:
			__obf_6fd3f72eb9b5574c.__obf_fea513943ca1a17c.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_9513c678a7f56a64:
			__obf_6fd3f72eb9b5574c.__obf_22621062d10dbb68.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		case __obf_6fd3f72eb9b5574c.__obf_8d7c2a74d36313ad:
			__obf_6fd3f72eb9b5574c.__obf_bc4e59499e3bf5bb.
				Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		default:
			__obf_1bb30e8a74ed8233.
				Skip()
		}
		if __obf_1bb30e8a74ed8233.__obf_40d85895d419ee4a() {
			break
		}
	}
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF && len(__obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea.Type1().Name()) != 0 {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%v.%s", __obf_6fd3f72eb9b5574c.__obf_29ebd0f2c324f5ea, __obf_1bb30e8a74ed8233.Error.Error())
	}
	__obf_1bb30e8a74ed8233.__obf_e792aa2f6324acea()
}

type __obf_e06930e4e8a0db9e struct {
	__obf_7e01b5b4651074d4 reflect2.StructField
	__obf_6cd5b2ea3acb38a5 ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_e06930e4e8a0db9e) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_43dac9f05df8fcc6 := __obf_6fd3f72eb9b5574c.__obf_7e01b5b4651074d4.UnsafeGet(__obf_2a1474febb02279b)
	__obf_6fd3f72eb9b5574c.__obf_6cd5b2ea3acb38a5.
		Decode(__obf_43dac9f05df8fcc6, __obf_1bb30e8a74ed8233)
	if __obf_1bb30e8a74ed8233.Error != nil && __obf_1bb30e8a74ed8233.Error != io.EOF {
		__obf_1bb30e8a74ed8233.
			Error = fmt.Errorf("%s: %s", __obf_6fd3f72eb9b5574c.__obf_7e01b5b4651074d4.Name(), __obf_1bb30e8a74ed8233.Error.Error())
	}
}

type __obf_cd171cb12faf3492 struct {
	__obf_0110b2c027051777 ValDecoder
	__obf_892632c77e859037 *__obf_972c2293804d141c
}

func (__obf_6fd3f72eb9b5574c *__obf_cd171cb12faf3492) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
		Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	__obf_e91bd2feb751e4f1 := *((*string)(__obf_2a1474febb02279b))
	__obf_8e209b818aa7e088 := __obf_6fd3f72eb9b5574c.__obf_892632c77e859037.BorrowIterator([]byte(__obf_e91bd2feb751e4f1))
	defer __obf_6fd3f72eb9b5574c.__obf_892632c77e859037.ReturnIterator(__obf_8e209b818aa7e088)
	*((*string)(__obf_2a1474febb02279b)) = __obf_8e209b818aa7e088.ReadString()
}

type __obf_246bb5830b114872 struct {
	__obf_0110b2c027051777 ValDecoder
}

func (__obf_6fd3f72eb9b5574c *__obf_246bb5830b114872) Decode(__obf_2a1474febb02279b unsafe.Pointer, __obf_1bb30e8a74ed8233 *Iterator) {
	if __obf_1bb30e8a74ed8233.WhatIsNext() == NilValue {
		__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
			Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
		return
	}
	__obf_f16b4157911bc9af := __obf_1bb30e8a74ed8233.__obf_684faa48ae8c5049()
	if __obf_f16b4157911bc9af != '"' {
		__obf_1bb30e8a74ed8233.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
	__obf_6fd3f72eb9b5574c.__obf_0110b2c027051777.
		Decode(__obf_2a1474febb02279b, __obf_1bb30e8a74ed8233)
	if __obf_1bb30e8a74ed8233.Error != nil {
		return
	}
	__obf_f16b4157911bc9af = __obf_1bb30e8a74ed8233.__obf_9617ab9cc89bcddc()
	if __obf_f16b4157911bc9af != '"' {
		__obf_1bb30e8a74ed8233.
			ReportError("stringModeNumberDecoder", `expect ", but found `+string([]byte{__obf_f16b4157911bc9af}))
		return
	}
}
