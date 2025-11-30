package __obf_de86cdc8ae98b45b

import (
	"fmt"
	"reflect"
)

var __obf_e1e749f42ef9d0e6 = reflect.TypeOf((*[]string)(nil))

// DecodeArrayLen decodes array length. Length is -1 when array is nil.
func (__obf_dcaad1165bb07af9 *Decoder) DecodeArrayLen() (int, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return 0, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_5f964cf181cfcbf3(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_5f964cf181cfcbf3(__obf_ecf6d2d3253a058d byte) (int, error) {
	if __obf_ecf6d2d3253a058d == Nil {
		return -1, nil
	} else if __obf_ecf6d2d3253a058d >= FixedArrayLow && __obf_ecf6d2d3253a058d <= FixedArrayHigh {
		return int(__obf_ecf6d2d3253a058d & FixedArrayMask), nil
	}
	switch __obf_ecf6d2d3253a058d {
	case Array16:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint16()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	case Array32:
		__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.uint32()
		return int(__obf_2b0247e819b1bf4a), __obf_0feb3528b7b709ec
	}
	return 0, fmt.Errorf("msgpack: invalid code=%x decoding array length", __obf_ecf6d2d3253a058d)
}

func __obf_8de5f6238f2c13bf(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_01f0154f1fce8e5a := __obf_17732590722140e7.Addr().Convert(__obf_e1e749f42ef9d0e6).Interface().(*[]string)
	return __obf_dcaad1165bb07af9.__obf_6a43c6fd68849758(__obf_01f0154f1fce8e5a)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_6a43c6fd68849758(__obf_01f0154f1fce8e5a *[]string) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeArrayLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return nil
	}
	__obf_02601a79721c2043 := __obf_fa528de2519f24de(*__obf_01f0154f1fce8e5a, __obf_2b0247e819b1bf4a, __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_c646a26ee823abc3 != 0)
	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_2b0247e819b1bf4a; __obf_101eec84c8338296++ {
		__obf_a93d004caac96500, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeString()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		__obf_02601a79721c2043 = append(__obf_02601a79721c2043, __obf_a93d004caac96500)
	}
	*__obf_01f0154f1fce8e5a = __obf_02601a79721c2043

	return nil
}

func __obf_fa528de2519f24de(__obf_a93d004caac96500 []string, __obf_2b0247e819b1bf4a int, __obf_a78133bc1cc5d2fb bool) []string {
	if !__obf_a78133bc1cc5d2fb && __obf_2b0247e819b1bf4a > __obf_bb4b009fa46ad65f {
		__obf_2b0247e819b1bf4a = __obf_bb4b009fa46ad65f
	}

	if __obf_a93d004caac96500 == nil {
		return make([]string, 0, __obf_2b0247e819b1bf4a)
	}

	if cap(__obf_a93d004caac96500) >= __obf_2b0247e819b1bf4a {
		return __obf_a93d004caac96500[:0]
	}
	__obf_a93d004caac96500 = __obf_a93d004caac96500[:cap(__obf_a93d004caac96500)]
	__obf_a93d004caac96500 = append(__obf_a93d004caac96500, make([]string, __obf_2b0247e819b1bf4a-len(__obf_a93d004caac96500))...)
	return __obf_a93d004caac96500[:0]
}

func __obf_68ff6383b6d6d266(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeArrayLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	if __obf_2b0247e819b1bf4a == -1 {
		__obf_17732590722140e7.
			Set(reflect.Zero(__obf_17732590722140e7.Type()))
		return nil
	}
	if __obf_2b0247e819b1bf4a == 0 && __obf_17732590722140e7.IsNil() {
		__obf_17732590722140e7.
			Set(reflect.MakeSlice(__obf_17732590722140e7.Type(), 0, 0))
		return nil
	}

	if __obf_17732590722140e7.Cap() >= __obf_2b0247e819b1bf4a {
		__obf_17732590722140e7.
			Set(__obf_17732590722140e7.Slice(0, __obf_2b0247e819b1bf4a))
	} else if __obf_17732590722140e7.Len() < __obf_17732590722140e7.Cap() {
		__obf_17732590722140e7.
			Set(__obf_17732590722140e7.Slice(0, __obf_17732590722140e7.Cap()))
	}
	__obf_a78133bc1cc5d2fb := __obf_dcaad1165bb07af9.__obf_a15a642720e39c3e&__obf_c646a26ee823abc3 != 1

	if __obf_a78133bc1cc5d2fb && __obf_2b0247e819b1bf4a > __obf_17732590722140e7.Len() {
		__obf_17732590722140e7.
			Set(__obf_d1be133c460b0e52(__obf_17732590722140e7, __obf_2b0247e819b1bf4a, __obf_a78133bc1cc5d2fb))
	}

	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_2b0247e819b1bf4a; __obf_101eec84c8338296++ {
		if !__obf_a78133bc1cc5d2fb && __obf_101eec84c8338296 >= __obf_17732590722140e7.Len() {
			__obf_17732590722140e7.
				Set(__obf_d1be133c460b0e52(__obf_17732590722140e7, __obf_2b0247e819b1bf4a, __obf_a78133bc1cc5d2fb))
		}
		__obf_dd3641741f7dbdf3 := __obf_17732590722140e7.Index(__obf_101eec84c8338296)
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeValue(__obf_dd3641741f7dbdf3); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func __obf_d1be133c460b0e52(__obf_17732590722140e7 reflect.Value, __obf_2b0247e819b1bf4a int, __obf_a78133bc1cc5d2fb bool) reflect.Value {
	__obf_e9a0ff5f428c359e := __obf_2b0247e819b1bf4a - __obf_17732590722140e7.Len()
	if !__obf_a78133bc1cc5d2fb && __obf_e9a0ff5f428c359e > __obf_bb4b009fa46ad65f {
		__obf_e9a0ff5f428c359e = __obf_bb4b009fa46ad65f
	}
	__obf_17732590722140e7 = reflect.AppendSlice(__obf_17732590722140e7, reflect.MakeSlice(__obf_17732590722140e7.Type(), __obf_e9a0ff5f428c359e, __obf_e9a0ff5f428c359e))
	return __obf_17732590722140e7
}

func __obf_28619ae619606d09(__obf_dcaad1165bb07af9 *Decoder, __obf_17732590722140e7 reflect.Value) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeArrayLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	if __obf_2b0247e819b1bf4a == -1 {
		return nil
	}
	if __obf_2b0247e819b1bf4a > __obf_17732590722140e7.Len() {
		return fmt.Errorf("%s len is %d, but msgpack has %d elements", __obf_17732590722140e7.Type(), __obf_17732590722140e7.Len(), __obf_2b0247e819b1bf4a)
	}

	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_2b0247e819b1bf4a; __obf_101eec84c8338296++ {
		__obf_9d65368d74c58cbf := __obf_17732590722140e7.Index(__obf_101eec84c8338296)
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeValue(__obf_9d65368d74c58cbf); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) DecodeSlice() ([]any, error) {
	__obf_ecf6d2d3253a058d, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_74ac05d9a9da01f4()
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	return __obf_dcaad1165bb07af9.__obf_9d79cf59b9daca03(__obf_ecf6d2d3253a058d)
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_9d79cf59b9daca03(__obf_ecf6d2d3253a058d byte) ([]any, error) {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_5f964cf181cfcbf3(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return nil, nil
	}
	__obf_a93d004caac96500 := make([]any, 0, __obf_2b0247e819b1bf4a)
	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_2b0247e819b1bf4a; __obf_101eec84c8338296++ {
		__obf_17732590722140e7, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
		if __obf_0feb3528b7b709ec != nil {
			return nil, __obf_0feb3528b7b709ec
		}
		__obf_a93d004caac96500 = append(__obf_a93d004caac96500, __obf_17732590722140e7)
	}

	return __obf_a93d004caac96500, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_f7821d6a7eae9517(__obf_ecf6d2d3253a058d byte) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_5f964cf181cfcbf3(__obf_ecf6d2d3253a058d)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	for __obf_101eec84c8338296 := 0; __obf_101eec84c8338296 < __obf_2b0247e819b1bf4a; __obf_101eec84c8338296++ {
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}
