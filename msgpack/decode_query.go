package __obf_a935eb7f548271a4

import (
	"fmt"
	"strconv"
	"strings"
)

type __obf_ad0119bb768458ef struct {
	__obf_7b829a391c4b2139 string
	__obf_5603ee7e7690c5f5 string
	__obf_7eb8c1f10fd4cab2 []any
	__obf_1121b055793cc371 bool
}

func (__obf_d8f054449f1d66ac *__obf_ad0119bb768458ef) __obf_a0a95e7fb07a89fa() {
	__obf_12c2fbb06793e18d := strings.IndexByte(__obf_d8f054449f1d66ac.__obf_7b829a391c4b2139, '.')
	if __obf_12c2fbb06793e18d == -1 {
		__obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5 = __obf_d8f054449f1d66ac.__obf_7b829a391c4b2139
		__obf_d8f054449f1d66ac.__obf_7b829a391c4b2139 = ""
		return
	}
	__obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5 = __obf_d8f054449f1d66ac.__obf_7b829a391c4b2139[:__obf_12c2fbb06793e18d]
	__obf_d8f054449f1d66ac.__obf_7b829a391c4b2139 = __obf_d8f054449f1d66ac.

		// Query extracts data specified by the query from the msgpack stream skipping
		// any other data. Query consists of map keys and array indexes separated with dot,
		// e.g. key1.0.key2.
		__obf_7b829a391c4b2139[__obf_12c2fbb06793e18d+1:]
}

func (__obf_a21885da2425f2b2 *Decoder) Query(__obf_7b829a391c4b2139 string) ([]any, error) {
	__obf_b32e29f5594f2124 := __obf_ad0119bb768458ef{__obf_7b829a391c4b2139: __obf_7b829a391c4b2139}
	if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_7b829a391c4b2139(&__obf_b32e29f5594f2124); __obf_4d327e1cd40c2e21 != nil {
		return nil, __obf_4d327e1cd40c2e21
	}
	return __obf_b32e29f5594f2124.__obf_7eb8c1f10fd4cab2, nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_7b829a391c4b2139(__obf_d8f054449f1d66ac *__obf_ad0119bb768458ef) error {
	__obf_d8f054449f1d66ac.__obf_a0a95e7fb07a89fa()
	if __obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5 == "" {
		__obf_6d570581f4b60dbc, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_c4227998727b27c6()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
		__obf_d8f054449f1d66ac.__obf_7eb8c1f10fd4cab2 = append(__obf_d8f054449f1d66ac.__obf_7eb8c1f10fd4cab2, __obf_6d570581f4b60dbc)
		return nil
	}
	__obf_b983039ef2a336eb, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.PeekCode()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	switch {
	case __obf_b983039ef2a336eb == Map16 || __obf_b983039ef2a336eb == Map32 || IsFixedMap(__obf_b983039ef2a336eb):
		__obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_204f5b8358d16348(__obf_d8f054449f1d66ac)
	case __obf_b983039ef2a336eb == Array16 || __obf_b983039ef2a336eb == Array32 || IsFixedArray(__obf_b983039ef2a336eb):
		__obf_4d327e1cd40c2e21 = __obf_a21885da2425f2b2.__obf_1179bf439a0336f7(__obf_d8f054449f1d66ac)
	default:
		__obf_4d327e1cd40c2e21 = fmt.Errorf("msgpack: unsupported code=%x decoding key=%q", __obf_b983039ef2a336eb, __obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5)
	}
	return __obf_4d327e1cd40c2e21
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_204f5b8358d16348(__obf_d8f054449f1d66ac *__obf_ad0119bb768458ef) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeMapLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return nil
	}

	for __obf_4a8e280ffaa52cf4 := range __obf_326af9bd942662ac {
		__obf_5603ee7e7690c5f5, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_4e3696e2f7b41e6e()
		if __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}

		if __obf_5603ee7e7690c5f5 == __obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5 {
			if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_7b829a391c4b2139(__obf_d8f054449f1d66ac); __obf_4d327e1cd40c2e21 != nil {
				return __obf_4d327e1cd40c2e21
			}
			if __obf_d8f054449f1d66ac.__obf_1121b055793cc371 {
				return __obf_a21885da2425f2b2.__obf_def69edd3090c83d((__obf_326af9bd942662ac - __obf_4a8e280ffaa52cf4 - 1) * 2)
			}
			return nil
		}

		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_1179bf439a0336f7(__obf_d8f054449f1d66ac *__obf_ad0119bb768458ef) error {
	__obf_326af9bd942662ac, __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.DecodeArrayLen()
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}
	if __obf_326af9bd942662ac == -1 {
		return nil
	}

	if __obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5 == "*" {
		__obf_d8f054449f1d66ac.__obf_1121b055793cc371 = true
		__obf_7b829a391c4b2139 := __obf_d8f054449f1d66ac.__obf_7b829a391c4b2139
		for range __obf_326af9bd942662ac {
			__obf_d8f054449f1d66ac.__obf_7b829a391c4b2139 = __obf_7b829a391c4b2139
			if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_7b829a391c4b2139(__obf_d8f054449f1d66ac); __obf_4d327e1cd40c2e21 != nil {
				return __obf_4d327e1cd40c2e21
			}
		}
		__obf_d8f054449f1d66ac.__obf_1121b055793cc371 = false
		return nil
	}
	__obf_12c2fbb06793e18d, __obf_4d327e1cd40c2e21 := strconv.Atoi(__obf_d8f054449f1d66ac.__obf_5603ee7e7690c5f5)
	if __obf_4d327e1cd40c2e21 != nil {
		return __obf_4d327e1cd40c2e21
	}

	for __obf_4a8e280ffaa52cf4 := range __obf_326af9bd942662ac {
		if __obf_4a8e280ffaa52cf4 == __obf_12c2fbb06793e18d {
			if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.__obf_7b829a391c4b2139(__obf_d8f054449f1d66ac); __obf_4d327e1cd40c2e21 != nil {
				return __obf_4d327e1cd40c2e21
			}
			if __obf_d8f054449f1d66ac.__obf_1121b055793cc371 {
				return __obf_a21885da2425f2b2.__obf_def69edd3090c83d(__obf_326af9bd942662ac - __obf_4a8e280ffaa52cf4 - 1)
			}
			return nil
		}

		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}

	return nil
}

func (__obf_a21885da2425f2b2 *Decoder) __obf_def69edd3090c83d(__obf_326af9bd942662ac int) error {
	for range __obf_326af9bd942662ac {
		if __obf_4d327e1cd40c2e21 := __obf_a21885da2425f2b2.Skip(); __obf_4d327e1cd40c2e21 != nil {
			return __obf_4d327e1cd40c2e21
		}
	}
	return nil
}
