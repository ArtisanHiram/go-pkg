package __obf_3edaa49e853afa16

import (
	"fmt"
	"strconv"
	"strings"
)

type __obf_6c9debbbd4e8b232 struct {
	__obf_a0ed3c5c019a34c5 string
	__obf_813ce87902031fb7 string
	__obf_d174adee2c60f1db []any
	__obf_38a3e5d51882430c bool
}

func (__obf_0d0d96d041e4f664 *__obf_6c9debbbd4e8b232) __obf_496d2b7d22f72742() {
	__obf_996abb886f808883 := strings.IndexByte(__obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5, '.')
	if __obf_996abb886f808883 == -1 {
		__obf_0d0d96d041e4f664.__obf_813ce87902031fb7 = __obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5
		__obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5 = ""
		return
	}
	__obf_0d0d96d041e4f664.__obf_813ce87902031fb7 = __obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5[:__obf_996abb886f808883]
	__obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5 = __obf_0d0d96d041e4f664.

		// Query extracts data specified by the query from the msgpack stream skipping
		// any other data. Query consists of map keys and array indexes separated with dot,
		// e.g. key1.0.key2.
		__obf_a0ed3c5c019a34c5[__obf_996abb886f808883+1:]
}

func (__obf_015afbee33862a01 *Decoder) Query(__obf_a0ed3c5c019a34c5 string) ([]any, error) {
	__obf_6c298130e577790e := __obf_6c9debbbd4e8b232{__obf_a0ed3c5c019a34c5: __obf_a0ed3c5c019a34c5}
	if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_a0ed3c5c019a34c5(&__obf_6c298130e577790e); __obf_3aa78ad28f50ed46 != nil {
		return nil, __obf_3aa78ad28f50ed46
	}
	return __obf_6c298130e577790e.__obf_d174adee2c60f1db, nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_a0ed3c5c019a34c5(__obf_0d0d96d041e4f664 *__obf_6c9debbbd4e8b232) error {
	__obf_0d0d96d041e4f664.__obf_496d2b7d22f72742()
	if __obf_0d0d96d041e4f664.__obf_813ce87902031fb7 == "" {
		__obf_85d270343a0dfe14, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_ed59dd67f1d39cb1()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
		__obf_0d0d96d041e4f664.__obf_d174adee2c60f1db = append(__obf_0d0d96d041e4f664.__obf_d174adee2c60f1db, __obf_85d270343a0dfe14)
		return nil
	}
	__obf_508e2d6ff53655c0, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.PeekCode()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	switch {
	case __obf_508e2d6ff53655c0 == Map16 || __obf_508e2d6ff53655c0 == Map32 || IsFixedMap(__obf_508e2d6ff53655c0):
		__obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_958065cacc99a299(__obf_0d0d96d041e4f664)
	case __obf_508e2d6ff53655c0 == Array16 || __obf_508e2d6ff53655c0 == Array32 || IsFixedArray(__obf_508e2d6ff53655c0):
		__obf_3aa78ad28f50ed46 = __obf_015afbee33862a01.__obf_00c17d833418027f(__obf_0d0d96d041e4f664)
	default:
		__obf_3aa78ad28f50ed46 = fmt.Errorf("msgpack: unsupported code=%x decoding key=%q", __obf_508e2d6ff53655c0, __obf_0d0d96d041e4f664.__obf_813ce87902031fb7)
	}
	return __obf_3aa78ad28f50ed46
}

func (__obf_015afbee33862a01 *Decoder) __obf_958065cacc99a299(__obf_0d0d96d041e4f664 *__obf_6c9debbbd4e8b232) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeMapLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return nil
	}

	for __obf_bd2da3a1d4616916 := range __obf_56127cd370854f0b {
		__obf_813ce87902031fb7, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_aeaabfb6cb02451a()
		if __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}

		if __obf_813ce87902031fb7 == __obf_0d0d96d041e4f664.__obf_813ce87902031fb7 {
			if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_a0ed3c5c019a34c5(__obf_0d0d96d041e4f664); __obf_3aa78ad28f50ed46 != nil {
				return __obf_3aa78ad28f50ed46
			}
			if __obf_0d0d96d041e4f664.__obf_38a3e5d51882430c {
				return __obf_015afbee33862a01.__obf_c6d590f40802ea74((__obf_56127cd370854f0b - __obf_bd2da3a1d4616916 - 1) * 2)
			}
			return nil
		}

		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_00c17d833418027f(__obf_0d0d96d041e4f664 *__obf_6c9debbbd4e8b232) error {
	__obf_56127cd370854f0b, __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.DecodeArrayLen()
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}
	if __obf_56127cd370854f0b == -1 {
		return nil
	}

	if __obf_0d0d96d041e4f664.__obf_813ce87902031fb7 == "*" {
		__obf_0d0d96d041e4f664.__obf_38a3e5d51882430c = true
		__obf_a0ed3c5c019a34c5 := __obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5
		for range __obf_56127cd370854f0b {
			__obf_0d0d96d041e4f664.__obf_a0ed3c5c019a34c5 = __obf_a0ed3c5c019a34c5
			if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_a0ed3c5c019a34c5(__obf_0d0d96d041e4f664); __obf_3aa78ad28f50ed46 != nil {
				return __obf_3aa78ad28f50ed46
			}
		}
		__obf_0d0d96d041e4f664.__obf_38a3e5d51882430c = false
		return nil
	}
	__obf_996abb886f808883, __obf_3aa78ad28f50ed46 := strconv.Atoi(__obf_0d0d96d041e4f664.__obf_813ce87902031fb7)
	if __obf_3aa78ad28f50ed46 != nil {
		return __obf_3aa78ad28f50ed46
	}

	for __obf_bd2da3a1d4616916 := range __obf_56127cd370854f0b {
		if __obf_bd2da3a1d4616916 == __obf_996abb886f808883 {
			if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.__obf_a0ed3c5c019a34c5(__obf_0d0d96d041e4f664); __obf_3aa78ad28f50ed46 != nil {
				return __obf_3aa78ad28f50ed46
			}
			if __obf_0d0d96d041e4f664.__obf_38a3e5d51882430c {
				return __obf_015afbee33862a01.__obf_c6d590f40802ea74(__obf_56127cd370854f0b - __obf_bd2da3a1d4616916 - 1)
			}
			return nil
		}

		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}

	return nil
}

func (__obf_015afbee33862a01 *Decoder) __obf_c6d590f40802ea74(__obf_56127cd370854f0b int) error {
	for range __obf_56127cd370854f0b {
		if __obf_3aa78ad28f50ed46 := __obf_015afbee33862a01.Skip(); __obf_3aa78ad28f50ed46 != nil {
			return __obf_3aa78ad28f50ed46
		}
	}
	return nil
}
