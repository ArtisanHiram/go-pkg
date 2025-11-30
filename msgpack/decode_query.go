package __obf_fd770cb9675903b0

import (
	"fmt"
	"strconv"
	"strings"
)

type __obf_bcf8fce0287c6d64 struct {
	__obf_3ea486c8a3af8da5 string
	__obf_e155b5a95e20e61d string
	__obf_585aa8d0d6128b71 []any
	__obf_22c41cd496630034 bool
}

func (__obf_9f310029af8b652b *__obf_bcf8fce0287c6d64) __obf_963e6e482fe878bd() {
	__obf_262845baa82482a2 := strings.IndexByte(__obf_9f310029af8b652b.__obf_3ea486c8a3af8da5, '.')
	if __obf_262845baa82482a2 == -1 {
		__obf_9f310029af8b652b.__obf_e155b5a95e20e61d = __obf_9f310029af8b652b.__obf_3ea486c8a3af8da5
		__obf_9f310029af8b652b.__obf_3ea486c8a3af8da5 = ""
		return
	}
	__obf_9f310029af8b652b.__obf_e155b5a95e20e61d = __obf_9f310029af8b652b.__obf_3ea486c8a3af8da5[:__obf_262845baa82482a2]
	__obf_9f310029af8b652b.__obf_3ea486c8a3af8da5 = __obf_9f310029af8b652b.

		// Query extracts data specified by the query from the msgpack stream skipping
		// any other data. Query consists of map keys and array indexes separated with dot,
		// e.g. key1.0.key2.
		__obf_3ea486c8a3af8da5[__obf_262845baa82482a2+1:]
}

func (__obf_5d389b660eefb08c *Decoder) Query(__obf_3ea486c8a3af8da5 string) ([]any, error) {
	__obf_f041ea83fdf1a5a6 := __obf_bcf8fce0287c6d64{__obf_3ea486c8a3af8da5: __obf_3ea486c8a3af8da5}
	if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_3ea486c8a3af8da5(&__obf_f041ea83fdf1a5a6); __obf_45342a3a754d12f5 != nil {
		return nil, __obf_45342a3a754d12f5
	}
	return __obf_f041ea83fdf1a5a6.__obf_585aa8d0d6128b71, nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_3ea486c8a3af8da5(__obf_9f310029af8b652b *__obf_bcf8fce0287c6d64) error {
	__obf_9f310029af8b652b.__obf_963e6e482fe878bd()
	if __obf_9f310029af8b652b.__obf_e155b5a95e20e61d == "" {
		__obf_f328a048f76b7256, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_d8bd92f8abc7c84f()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
		__obf_9f310029af8b652b.__obf_585aa8d0d6128b71 = append(__obf_9f310029af8b652b.__obf_585aa8d0d6128b71, __obf_f328a048f76b7256)
		return nil
	}
	__obf_cde82889ba8e4822, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.PeekCode()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	switch {
	case __obf_cde82889ba8e4822 == Map16 || __obf_cde82889ba8e4822 == Map32 || IsFixedMap(__obf_cde82889ba8e4822):
		__obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_c7f084fbafe9bca4(__obf_9f310029af8b652b)
	case __obf_cde82889ba8e4822 == Array16 || __obf_cde82889ba8e4822 == Array32 || IsFixedArray(__obf_cde82889ba8e4822):
		__obf_45342a3a754d12f5 = __obf_5d389b660eefb08c.__obf_e25f46c15b2c307d(__obf_9f310029af8b652b)
	default:
		__obf_45342a3a754d12f5 = fmt.Errorf("msgpack: unsupported code=%x decoding key=%q", __obf_cde82889ba8e4822, __obf_9f310029af8b652b.__obf_e155b5a95e20e61d)
	}
	return __obf_45342a3a754d12f5
}

func (__obf_5d389b660eefb08c *Decoder) __obf_c7f084fbafe9bca4(__obf_9f310029af8b652b *__obf_bcf8fce0287c6d64) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeMapLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return nil
	}

	for __obf_4140b3ff04f75a36 := range __obf_d774e4844c74bfe9 {
		__obf_e155b5a95e20e61d, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_479ae9aab7b05d68()
		if __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}

		if __obf_e155b5a95e20e61d == __obf_9f310029af8b652b.__obf_e155b5a95e20e61d {
			if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_3ea486c8a3af8da5(__obf_9f310029af8b652b); __obf_45342a3a754d12f5 != nil {
				return __obf_45342a3a754d12f5
			}
			if __obf_9f310029af8b652b.__obf_22c41cd496630034 {
				return __obf_5d389b660eefb08c.__obf_823a604433f8b593((__obf_d774e4844c74bfe9 - __obf_4140b3ff04f75a36 - 1) * 2)
			}
			return nil
		}

		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_e25f46c15b2c307d(__obf_9f310029af8b652b *__obf_bcf8fce0287c6d64) error {
	__obf_d774e4844c74bfe9, __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.DecodeArrayLen()
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}
	if __obf_d774e4844c74bfe9 == -1 {
		return nil
	}

	if __obf_9f310029af8b652b.__obf_e155b5a95e20e61d == "*" {
		__obf_9f310029af8b652b.__obf_22c41cd496630034 = true
		__obf_3ea486c8a3af8da5 := __obf_9f310029af8b652b.__obf_3ea486c8a3af8da5
		for range __obf_d774e4844c74bfe9 {
			__obf_9f310029af8b652b.__obf_3ea486c8a3af8da5 = __obf_3ea486c8a3af8da5
			if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_3ea486c8a3af8da5(__obf_9f310029af8b652b); __obf_45342a3a754d12f5 != nil {
				return __obf_45342a3a754d12f5
			}
		}
		__obf_9f310029af8b652b.__obf_22c41cd496630034 = false
		return nil
	}
	__obf_262845baa82482a2, __obf_45342a3a754d12f5 := strconv.Atoi(__obf_9f310029af8b652b.__obf_e155b5a95e20e61d)
	if __obf_45342a3a754d12f5 != nil {
		return __obf_45342a3a754d12f5
	}

	for __obf_4140b3ff04f75a36 := range __obf_d774e4844c74bfe9 {
		if __obf_4140b3ff04f75a36 == __obf_262845baa82482a2 {
			if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.__obf_3ea486c8a3af8da5(__obf_9f310029af8b652b); __obf_45342a3a754d12f5 != nil {
				return __obf_45342a3a754d12f5
			}
			if __obf_9f310029af8b652b.__obf_22c41cd496630034 {
				return __obf_5d389b660eefb08c.__obf_823a604433f8b593(__obf_d774e4844c74bfe9 - __obf_4140b3ff04f75a36 - 1)
			}
			return nil
		}

		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}

	return nil
}

func (__obf_5d389b660eefb08c *Decoder) __obf_823a604433f8b593(__obf_d774e4844c74bfe9 int) error {
	for range __obf_d774e4844c74bfe9 {
		if __obf_45342a3a754d12f5 := __obf_5d389b660eefb08c.Skip(); __obf_45342a3a754d12f5 != nil {
			return __obf_45342a3a754d12f5
		}
	}
	return nil
}
