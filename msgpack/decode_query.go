package __obf_a4aad98aaf3178f4

import (
	"fmt"
	"strconv"
	"strings"
)

type __obf_cd6a6e10376e542b struct {
	__obf_13ff8a8f9c1b455b string
	__obf_992cb53d7b9cb024 string
	__obf_48794fabe30081af []any
	__obf_ce6fe7f0d4142031 bool
}

func (__obf_4f57b370dcf7608a *__obf_cd6a6e10376e542b) __obf_a94124b358f00e05() {
	__obf_f17afa4466d0cfda := strings.IndexByte(__obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b, '.')
	if __obf_f17afa4466d0cfda == -1 {
		__obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024 = __obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b
		__obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b = ""
		return
	}
	__obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024 = __obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b[:__obf_f17afa4466d0cfda]
	__obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b = __obf_4f57b370dcf7608a.

		// Query extracts data specified by the query from the msgpack stream skipping
		// any other data. Query consists of map keys and array indexes separated with dot,
		// e.g. key1.0.key2.
		__obf_13ff8a8f9c1b455b[__obf_f17afa4466d0cfda+1:]
}

func (__obf_613397fefdec0ed0 *Decoder) Query(__obf_13ff8a8f9c1b455b string) ([]any, error) {
	__obf_ca41977552def92e := __obf_cd6a6e10376e542b{__obf_13ff8a8f9c1b455b: __obf_13ff8a8f9c1b455b}
	if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_13ff8a8f9c1b455b(&__obf_ca41977552def92e); __obf_4061ca0039150c39 != nil {
		return nil, __obf_4061ca0039150c39
	}
	return __obf_ca41977552def92e.__obf_48794fabe30081af, nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_13ff8a8f9c1b455b(__obf_4f57b370dcf7608a *__obf_cd6a6e10376e542b) error {
	__obf_4f57b370dcf7608a.__obf_a94124b358f00e05()
	if __obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024 == "" {
		__obf_a1f43267eeb48a1a, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_ec86210fb048ddc6()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
		__obf_4f57b370dcf7608a.__obf_48794fabe30081af = append(__obf_4f57b370dcf7608a.__obf_48794fabe30081af, __obf_a1f43267eeb48a1a)
		return nil
	}
	__obf_987b95dd4dcfc308, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.PeekCode()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	switch {
	case __obf_987b95dd4dcfc308 == Map16 || __obf_987b95dd4dcfc308 == Map32 || IsFixedMap(__obf_987b95dd4dcfc308):
		__obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_3dacd0477c347ee1(__obf_4f57b370dcf7608a)
	case __obf_987b95dd4dcfc308 == Array16 || __obf_987b95dd4dcfc308 == Array32 || IsFixedArray(__obf_987b95dd4dcfc308):
		__obf_4061ca0039150c39 = __obf_613397fefdec0ed0.__obf_773d4c75c1ac34bc(__obf_4f57b370dcf7608a)
	default:
		__obf_4061ca0039150c39 = fmt.Errorf("msgpack: unsupported code=%x decoding key=%q", __obf_987b95dd4dcfc308, __obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024)
	}
	return __obf_4061ca0039150c39
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_3dacd0477c347ee1(__obf_4f57b370dcf7608a *__obf_cd6a6e10376e542b) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeMapLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return nil
	}

	for __obf_097d8b6061c9592b := range __obf_99a74e41c9c640c0 {
		__obf_992cb53d7b9cb024, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_dc853c0361533caa()
		if __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}

		if __obf_992cb53d7b9cb024 == __obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024 {
			if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_13ff8a8f9c1b455b(__obf_4f57b370dcf7608a); __obf_4061ca0039150c39 != nil {
				return __obf_4061ca0039150c39
			}
			if __obf_4f57b370dcf7608a.__obf_ce6fe7f0d4142031 {
				return __obf_613397fefdec0ed0.__obf_3a8113979d3561fa((__obf_99a74e41c9c640c0 - __obf_097d8b6061c9592b - 1) * 2)
			}
			return nil
		}

		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_773d4c75c1ac34bc(__obf_4f57b370dcf7608a *__obf_cd6a6e10376e542b) error {
	__obf_99a74e41c9c640c0, __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.DecodeArrayLen()
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}
	if __obf_99a74e41c9c640c0 == -1 {
		return nil
	}

	if __obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024 == "*" {
		__obf_4f57b370dcf7608a.__obf_ce6fe7f0d4142031 = true
		__obf_13ff8a8f9c1b455b := __obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b
		for range __obf_99a74e41c9c640c0 {
			__obf_4f57b370dcf7608a.__obf_13ff8a8f9c1b455b = __obf_13ff8a8f9c1b455b
			if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_13ff8a8f9c1b455b(__obf_4f57b370dcf7608a); __obf_4061ca0039150c39 != nil {
				return __obf_4061ca0039150c39
			}
		}
		__obf_4f57b370dcf7608a.__obf_ce6fe7f0d4142031 = false
		return nil
	}
	__obf_f17afa4466d0cfda, __obf_4061ca0039150c39 := strconv.Atoi(__obf_4f57b370dcf7608a.__obf_992cb53d7b9cb024)
	if __obf_4061ca0039150c39 != nil {
		return __obf_4061ca0039150c39
	}

	for __obf_097d8b6061c9592b := range __obf_99a74e41c9c640c0 {
		if __obf_097d8b6061c9592b == __obf_f17afa4466d0cfda {
			if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.__obf_13ff8a8f9c1b455b(__obf_4f57b370dcf7608a); __obf_4061ca0039150c39 != nil {
				return __obf_4061ca0039150c39
			}
			if __obf_4f57b370dcf7608a.__obf_ce6fe7f0d4142031 {
				return __obf_613397fefdec0ed0.__obf_3a8113979d3561fa(__obf_99a74e41c9c640c0 - __obf_097d8b6061c9592b - 1)
			}
			return nil
		}

		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}

	return nil
}

func (__obf_613397fefdec0ed0 *Decoder) __obf_3a8113979d3561fa(__obf_99a74e41c9c640c0 int) error {
	for range __obf_99a74e41c9c640c0 {
		if __obf_4061ca0039150c39 := __obf_613397fefdec0ed0.Skip(); __obf_4061ca0039150c39 != nil {
			return __obf_4061ca0039150c39
		}
	}
	return nil
}
