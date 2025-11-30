package __obf_de86cdc8ae98b45b

import (
	"fmt"
	"strconv"
	"strings"
)

type __obf_796ab3aa8b298d7c struct {
	__obf_24c28199fdd9c899 string
	__obf_9cab2959bb95a876 string
	__obf_8f87b29abc21c78e []any
	__obf_65759fc87fc8b2a1 bool
}

func (__obf_8603db7064a72dd8 *__obf_796ab3aa8b298d7c) __obf_5c39e58d3d8e1337() {
	__obf_2ae61e2652c8efc2 := strings.IndexByte(__obf_8603db7064a72dd8.__obf_24c28199fdd9c899, '.')
	if __obf_2ae61e2652c8efc2 == -1 {
		__obf_8603db7064a72dd8.__obf_9cab2959bb95a876 = __obf_8603db7064a72dd8.__obf_24c28199fdd9c899
		__obf_8603db7064a72dd8.__obf_24c28199fdd9c899 = ""
		return
	}
	__obf_8603db7064a72dd8.__obf_9cab2959bb95a876 = __obf_8603db7064a72dd8.__obf_24c28199fdd9c899[:__obf_2ae61e2652c8efc2]
	__obf_8603db7064a72dd8.__obf_24c28199fdd9c899 = __obf_8603db7064a72dd8.

		// Query extracts data specified by the query from the msgpack stream skipping
		// any other data. Query consists of map keys and array indexes separated with dot,
		// e.g. key1.0.key2.
		__obf_24c28199fdd9c899[__obf_2ae61e2652c8efc2+1:]
}

func (__obf_dcaad1165bb07af9 *Decoder) Query(__obf_24c28199fdd9c899 string) ([]any, error) {
	__obf_d6385083ce8048de := __obf_796ab3aa8b298d7c{__obf_24c28199fdd9c899: __obf_24c28199fdd9c899}
	if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_24c28199fdd9c899(&__obf_d6385083ce8048de); __obf_0feb3528b7b709ec != nil {
		return nil, __obf_0feb3528b7b709ec
	}
	return __obf_d6385083ce8048de.__obf_8f87b29abc21c78e, nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_24c28199fdd9c899(__obf_8603db7064a72dd8 *__obf_796ab3aa8b298d7c) error {
	__obf_8603db7064a72dd8.__obf_5c39e58d3d8e1337()
	if __obf_8603db7064a72dd8.__obf_9cab2959bb95a876 == "" {
		__obf_17732590722140e7, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_7d3129f334945713()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
		__obf_8603db7064a72dd8.__obf_8f87b29abc21c78e = append(__obf_8603db7064a72dd8.__obf_8f87b29abc21c78e, __obf_17732590722140e7)
		return nil
	}
	__obf_34e3ba264a6bb541, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.PeekCode()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	switch {
	case __obf_34e3ba264a6bb541 == Map16 || __obf_34e3ba264a6bb541 == Map32 || IsFixedMap(__obf_34e3ba264a6bb541):
		__obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_42d13e65cc243bc9(__obf_8603db7064a72dd8)
	case __obf_34e3ba264a6bb541 == Array16 || __obf_34e3ba264a6bb541 == Array32 || IsFixedArray(__obf_34e3ba264a6bb541):
		__obf_0feb3528b7b709ec = __obf_dcaad1165bb07af9.__obf_4a7aefd0dacb7012(__obf_8603db7064a72dd8)
	default:
		__obf_0feb3528b7b709ec = fmt.Errorf("msgpack: unsupported code=%x decoding key=%q", __obf_34e3ba264a6bb541, __obf_8603db7064a72dd8.__obf_9cab2959bb95a876)
	}
	return __obf_0feb3528b7b709ec
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_42d13e65cc243bc9(__obf_8603db7064a72dd8 *__obf_796ab3aa8b298d7c) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeMapLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return nil
	}

	for __obf_101eec84c8338296 := range __obf_2b0247e819b1bf4a {
		__obf_9cab2959bb95a876, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_0605127b8a406c17()
		if __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}

		if __obf_9cab2959bb95a876 == __obf_8603db7064a72dd8.__obf_9cab2959bb95a876 {
			if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_24c28199fdd9c899(__obf_8603db7064a72dd8); __obf_0feb3528b7b709ec != nil {
				return __obf_0feb3528b7b709ec
			}
			if __obf_8603db7064a72dd8.__obf_65759fc87fc8b2a1 {
				return __obf_dcaad1165bb07af9.__obf_222ef3e26cad437d((__obf_2b0247e819b1bf4a - __obf_101eec84c8338296 - 1) * 2)
			}
			return nil
		}

		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_4a7aefd0dacb7012(__obf_8603db7064a72dd8 *__obf_796ab3aa8b298d7c) error {
	__obf_2b0247e819b1bf4a, __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.DecodeArrayLen()
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}
	if __obf_2b0247e819b1bf4a == -1 {
		return nil
	}

	if __obf_8603db7064a72dd8.__obf_9cab2959bb95a876 == "*" {
		__obf_8603db7064a72dd8.__obf_65759fc87fc8b2a1 = true
		__obf_24c28199fdd9c899 := __obf_8603db7064a72dd8.__obf_24c28199fdd9c899
		for range __obf_2b0247e819b1bf4a {
			__obf_8603db7064a72dd8.__obf_24c28199fdd9c899 = __obf_24c28199fdd9c899
			if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_24c28199fdd9c899(__obf_8603db7064a72dd8); __obf_0feb3528b7b709ec != nil {
				return __obf_0feb3528b7b709ec
			}
		}
		__obf_8603db7064a72dd8.__obf_65759fc87fc8b2a1 = false
		return nil
	}
	__obf_2ae61e2652c8efc2, __obf_0feb3528b7b709ec := strconv.Atoi(__obf_8603db7064a72dd8.__obf_9cab2959bb95a876)
	if __obf_0feb3528b7b709ec != nil {
		return __obf_0feb3528b7b709ec
	}

	for __obf_101eec84c8338296 := range __obf_2b0247e819b1bf4a {
		if __obf_101eec84c8338296 == __obf_2ae61e2652c8efc2 {
			if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.__obf_24c28199fdd9c899(__obf_8603db7064a72dd8); __obf_0feb3528b7b709ec != nil {
				return __obf_0feb3528b7b709ec
			}
			if __obf_8603db7064a72dd8.__obf_65759fc87fc8b2a1 {
				return __obf_dcaad1165bb07af9.__obf_222ef3e26cad437d(__obf_2b0247e819b1bf4a - __obf_101eec84c8338296 - 1)
			}
			return nil
		}

		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}

	return nil
}

func (__obf_dcaad1165bb07af9 *Decoder) __obf_222ef3e26cad437d(__obf_2b0247e819b1bf4a int) error {
	for range __obf_2b0247e819b1bf4a {
		if __obf_0feb3528b7b709ec := __obf_dcaad1165bb07af9.Skip(); __obf_0feb3528b7b709ec != nil {
			return __obf_0feb3528b7b709ec
		}
	}
	return nil
}
