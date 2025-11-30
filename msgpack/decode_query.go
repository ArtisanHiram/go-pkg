package __obf_3e0c303610a19bd4

import (
	"fmt"
	"strconv"
	"strings"
)

type __obf_4cae6a17c6678e0a struct {
	__obf_87f591514c99eda0 string
	__obf_ef8207c019b4cdb3 string
	__obf_28a7137f48d51bd1 []any
	__obf_61245d131d5d12d9 bool
}

func (__obf_dead1571ca5a6866 *__obf_4cae6a17c6678e0a) __obf_5b4e17a221110a53() {
	__obf_31cc89bd2a8a46e4 := strings.IndexByte(__obf_dead1571ca5a6866.__obf_87f591514c99eda0, '.')
	if __obf_31cc89bd2a8a46e4 == -1 {
		__obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3 = __obf_dead1571ca5a6866.__obf_87f591514c99eda0
		__obf_dead1571ca5a6866.__obf_87f591514c99eda0 = ""
		return
	}
	__obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3 = __obf_dead1571ca5a6866.__obf_87f591514c99eda0[:__obf_31cc89bd2a8a46e4]
	__obf_dead1571ca5a6866.__obf_87f591514c99eda0 = __obf_dead1571ca5a6866.

		// Query extracts data specified by the query from the msgpack stream skipping
		// any other data. Query consists of map keys and array indexes separated with dot,
		// e.g. key1.0.key2.
		__obf_87f591514c99eda0[__obf_31cc89bd2a8a46e4+1:]
}

func (__obf_dc35117108ba8439 *Decoder) Query(__obf_87f591514c99eda0 string) ([]any, error) {
	__obf_696a28d929afb221 := __obf_4cae6a17c6678e0a{__obf_87f591514c99eda0: __obf_87f591514c99eda0}
	if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_87f591514c99eda0(&__obf_696a28d929afb221); __obf_8882f6cf6e378ded != nil {
		return nil, __obf_8882f6cf6e378ded
	}
	return __obf_696a28d929afb221.__obf_28a7137f48d51bd1, nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_87f591514c99eda0(__obf_dead1571ca5a6866 *__obf_4cae6a17c6678e0a) error {
	__obf_dead1571ca5a6866.__obf_5b4e17a221110a53()
	if __obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3 == "" {
		__obf_63bbcee86d44fdde, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_93ff057c8a5c795c()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
		__obf_dead1571ca5a6866.__obf_28a7137f48d51bd1 = append(__obf_dead1571ca5a6866.__obf_28a7137f48d51bd1, __obf_63bbcee86d44fdde)
		return nil
	}
	__obf_545372fefbb733e5, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.PeekCode()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	switch {
	case __obf_545372fefbb733e5 == Map16 || __obf_545372fefbb733e5 == Map32 || IsFixedMap(__obf_545372fefbb733e5):
		__obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_17e2327b537b0503(__obf_dead1571ca5a6866)
	case __obf_545372fefbb733e5 == Array16 || __obf_545372fefbb733e5 == Array32 || IsFixedArray(__obf_545372fefbb733e5):
		__obf_8882f6cf6e378ded = __obf_dc35117108ba8439.__obf_f79285768626503f(__obf_dead1571ca5a6866)
	default:
		__obf_8882f6cf6e378ded = fmt.Errorf("msgpack: unsupported code=%x decoding key=%q", __obf_545372fefbb733e5, __obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3)
	}
	return __obf_8882f6cf6e378ded
}

func (__obf_dc35117108ba8439 *Decoder) __obf_17e2327b537b0503(__obf_dead1571ca5a6866 *__obf_4cae6a17c6678e0a) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeMapLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return nil
	}

	for __obf_39714879601f9b69 := range __obf_4909ae60ffbb8e53 {
		__obf_ef8207c019b4cdb3, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_5484baeee52d4c8a()
		if __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}

		if __obf_ef8207c019b4cdb3 == __obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3 {
			if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_87f591514c99eda0(__obf_dead1571ca5a6866); __obf_8882f6cf6e378ded != nil {
				return __obf_8882f6cf6e378ded
			}
			if __obf_dead1571ca5a6866.__obf_61245d131d5d12d9 {
				return __obf_dc35117108ba8439.__obf_28ab1120aa2d2cc0((__obf_4909ae60ffbb8e53 - __obf_39714879601f9b69 - 1) * 2)
			}
			return nil
		}

		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_f79285768626503f(__obf_dead1571ca5a6866 *__obf_4cae6a17c6678e0a) error {
	__obf_4909ae60ffbb8e53, __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.DecodeArrayLen()
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}
	if __obf_4909ae60ffbb8e53 == -1 {
		return nil
	}

	if __obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3 == "*" {
		__obf_dead1571ca5a6866.__obf_61245d131d5d12d9 = true
		__obf_87f591514c99eda0 := __obf_dead1571ca5a6866.__obf_87f591514c99eda0
		for range __obf_4909ae60ffbb8e53 {
			__obf_dead1571ca5a6866.__obf_87f591514c99eda0 = __obf_87f591514c99eda0
			if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_87f591514c99eda0(__obf_dead1571ca5a6866); __obf_8882f6cf6e378ded != nil {
				return __obf_8882f6cf6e378ded
			}
		}
		__obf_dead1571ca5a6866.__obf_61245d131d5d12d9 = false
		return nil
	}
	__obf_31cc89bd2a8a46e4, __obf_8882f6cf6e378ded := strconv.Atoi(__obf_dead1571ca5a6866.__obf_ef8207c019b4cdb3)
	if __obf_8882f6cf6e378ded != nil {
		return __obf_8882f6cf6e378ded
	}

	for __obf_39714879601f9b69 := range __obf_4909ae60ffbb8e53 {
		if __obf_39714879601f9b69 == __obf_31cc89bd2a8a46e4 {
			if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.__obf_87f591514c99eda0(__obf_dead1571ca5a6866); __obf_8882f6cf6e378ded != nil {
				return __obf_8882f6cf6e378ded
			}
			if __obf_dead1571ca5a6866.__obf_61245d131d5d12d9 {
				return __obf_dc35117108ba8439.__obf_28ab1120aa2d2cc0(__obf_4909ae60ffbb8e53 - __obf_39714879601f9b69 - 1)
			}
			return nil
		}

		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}

	return nil
}

func (__obf_dc35117108ba8439 *Decoder) __obf_28ab1120aa2d2cc0(__obf_4909ae60ffbb8e53 int) error {
	for range __obf_4909ae60ffbb8e53 {
		if __obf_8882f6cf6e378ded := __obf_dc35117108ba8439.Skip(); __obf_8882f6cf6e378ded != nil {
			return __obf_8882f6cf6e378ded
		}
	}
	return nil
}
