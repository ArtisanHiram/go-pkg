// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_eba3d68b9605015e

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_71eb0882383cb0ef struct {
	__obf_ea46b5761247c020 []__obf_dbf34f1740ee9966// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...

}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_f3a458a9c566d56a(__obf_4e04dfa10559e27c *bitset.Bitset) __obf_71eb0882383cb0ef {
	__obf_aa12da5761ad5324 := __obf_4e04dfa10559e27c.Len() / 8
	if __obf_4e04dfa10559e27c.Len()%8 != 0 {
		__obf_aa12da5761ad5324++
	}
	__obf_7fca30a86b604ba4 := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: make([]__obf_dbf34f1740ee9966, __obf_aa12da5761ad5324)}
	__obf_a2ccc63c916c0db9 := __obf_aa12da5761ad5324 - 1
	for __obf_d6d3881236d45cd7 := 0; __obf_d6d3881236d45cd7 < __obf_4e04dfa10559e27c.Len(); __obf_d6d3881236d45cd7 += 8 {
		__obf_7fca30a86b604ba4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] = __obf_dbf34f1740ee9966(__obf_4e04dfa10559e27c.ByteAt(__obf_d6d3881236d45cd7))
		__obf_a2ccc63c916c0db9--
	}

	return __obf_7fca30a86b604ba4
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_a2a64cfa09775cd0(__obf_ea46b5761247c020 __obf_dbf34f1740ee9966, __obf_bf07404dd6aeed84 int) __obf_71eb0882383cb0ef {
	if __obf_ea46b5761247c020 == __obf_fbee904bb0eb5f92 {
		return __obf_71eb0882383cb0ef{}
	}
	__obf_7fca30a86b604ba4 := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: make([]__obf_dbf34f1740ee9966, __obf_bf07404dd6aeed84+1)}
	__obf_7fca30a86b604ba4.__obf_ea46b5761247c020[__obf_bf07404dd6aeed84] = __obf_ea46b5761247c020

	return __obf_7fca30a86b604ba4
}

func (__obf_d27433cd224562e4 __obf_71eb0882383cb0ef) __obf_4e04dfa10559e27c(__obf_d94642eeb977270c int) []byte {
	__obf_7fca30a86b604ba4 := make([]byte, __obf_d94642eeb977270c)
	__obf_a2ccc63c916c0db9 := __obf_d94642eeb977270c - len(__obf_d27433cd224562e4.__obf_ea46b5761247c020)
	for __obf_d6d3881236d45cd7 := len(__obf_d27433cd224562e4.__obf_ea46b5761247c020) - 1; __obf_d6d3881236d45cd7 >= 0; __obf_d6d3881236d45cd7-- {
		__obf_7fca30a86b604ba4[__obf_a2ccc63c916c0db9] = byte(__obf_d27433cd224562e4.__obf_ea46b5761247c020[__obf_d6d3881236d45cd7])
		__obf_a2ccc63c916c0db9++
	}

	return __obf_7fca30a86b604ba4
}

// numTerms returns the number of
func (__obf_d27433cd224562e4 __obf_71eb0882383cb0ef) __obf_d94642eeb977270c() int {
	return len(__obf_d27433cd224562e4.

		// gfPolyMultiply returns a * b.
		__obf_ea46b5761247c020)
}

func __obf_613fea38d6c97b2d(__obf_65fb4057f12dc33b, __obf_25d82b826f8df37a __obf_71eb0882383cb0ef) __obf_71eb0882383cb0ef {
	__obf_7099447d4eced32b := __obf_65fb4057f12dc33b.__obf_d94642eeb977270c()
	__obf_eddb1d1596c8e4de := __obf_25d82b826f8df37a.__obf_d94642eeb977270c()
	__obf_7fca30a86b604ba4 := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: make([]__obf_dbf34f1740ee9966, __obf_7099447d4eced32b+__obf_eddb1d1596c8e4de)}

	for __obf_a2ccc63c916c0db9 := 0; __obf_a2ccc63c916c0db9 < __obf_7099447d4eced32b; __obf_a2ccc63c916c0db9++ {
		for __obf_d6d3881236d45cd7 := 0; __obf_d6d3881236d45cd7 < __obf_eddb1d1596c8e4de; __obf_d6d3881236d45cd7++ {
			if __obf_65fb4057f12dc33b.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] != 0 && __obf_25d82b826f8df37a.__obf_ea46b5761247c020[__obf_d6d3881236d45cd7] != 0 {
				__obf_fd7e7fbc1c12a327 := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: make([]__obf_dbf34f1740ee9966, __obf_a2ccc63c916c0db9+__obf_d6d3881236d45cd7+1)}
				__obf_fd7e7fbc1c12a327.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9+__obf_d6d3881236d45cd7] = __obf_4cc39f544c1bed5c(__obf_65fb4057f12dc33b.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9], __obf_25d82b826f8df37a.__obf_ea46b5761247c020[__obf_d6d3881236d45cd7])
				__obf_7fca30a86b604ba4 = __obf_6842a3e45295a35a(__obf_7fca30a86b604ba4, __obf_fd7e7fbc1c12a327)
			}
		}
	}

	return __obf_7fca30a86b604ba4.__obf_59bfd375ff8d2ee7()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_9b723898ea675104(__obf_eea83fe6ab4aefe7, __obf_9f41a97e09144d25 __obf_71eb0882383cb0ef) __obf_71eb0882383cb0ef {
	if __obf_9f41a97e09144d25.__obf_f8b91d99fb7263a4(__obf_71eb0882383cb0ef{}) {
		log.Panicln("Remainder by zero")
	}
	__obf_9c353bc1b664ae8f := __obf_eea83fe6ab4aefe7

	for __obf_9c353bc1b664ae8f.__obf_d94642eeb977270c() >= __obf_9f41a97e09144d25.__obf_d94642eeb977270c() {
		__obf_bf07404dd6aeed84 := __obf_9c353bc1b664ae8f.__obf_d94642eeb977270c() - __obf_9f41a97e09144d25.__obf_d94642eeb977270c()
		__obf_269b30fdd3eadc2f := __obf_9badb47652259621(__obf_9c353bc1b664ae8f.__obf_ea46b5761247c020[__obf_9c353bc1b664ae8f.__obf_d94642eeb977270c()-1], __obf_9f41a97e09144d25.__obf_ea46b5761247c020[__obf_9f41a97e09144d25.__obf_d94642eeb977270c()-1])
		__obf_4b8301ead053993d := __obf_613fea38d6c97b2d(__obf_9f41a97e09144d25, __obf_a2a64cfa09775cd0(__obf_269b30fdd3eadc2f, __obf_bf07404dd6aeed84))
		__obf_9c353bc1b664ae8f = __obf_6842a3e45295a35a(__obf_9c353bc1b664ae8f, __obf_4b8301ead053993d)
	}

	return __obf_9c353bc1b664ae8f.__obf_59bfd375ff8d2ee7()
}

// gfPolyAdd returns a + b.
func __obf_6842a3e45295a35a(__obf_65fb4057f12dc33b, __obf_25d82b826f8df37a __obf_71eb0882383cb0ef) __obf_71eb0882383cb0ef {
	__obf_7099447d4eced32b := __obf_65fb4057f12dc33b.__obf_d94642eeb977270c()
	__obf_eddb1d1596c8e4de := __obf_25d82b826f8df37a.__obf_d94642eeb977270c()
	__obf_d94642eeb977270c := __obf_7099447d4eced32b
	if __obf_eddb1d1596c8e4de > __obf_d94642eeb977270c {
		__obf_d94642eeb977270c = __obf_eddb1d1596c8e4de
	}
	__obf_7fca30a86b604ba4 := __obf_71eb0882383cb0ef{__obf_ea46b5761247c020: make([]__obf_dbf34f1740ee9966, __obf_d94642eeb977270c)}

	for __obf_a2ccc63c916c0db9 := 0; __obf_a2ccc63c916c0db9 < __obf_d94642eeb977270c; __obf_a2ccc63c916c0db9++ {
		switch {
		case __obf_7099447d4eced32b > __obf_a2ccc63c916c0db9 && __obf_eddb1d1596c8e4de > __obf_a2ccc63c916c0db9:
			__obf_7fca30a86b604ba4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] = __obf_394177fab40ac561(__obf_65fb4057f12dc33b.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9], __obf_25d82b826f8df37a.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9])
		case __obf_7099447d4eced32b > __obf_a2ccc63c916c0db9:
			__obf_7fca30a86b604ba4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] = __obf_65fb4057f12dc33b.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9]
		default:
			__obf_7fca30a86b604ba4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] = __obf_25d82b826f8df37a.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9]
		}
	}

	return __obf_7fca30a86b604ba4.__obf_59bfd375ff8d2ee7()
}

func (__obf_d27433cd224562e4 __obf_71eb0882383cb0ef) __obf_59bfd375ff8d2ee7() __obf_71eb0882383cb0ef {
	__obf_d94642eeb977270c := __obf_d27433cd224562e4.__obf_d94642eeb977270c()
	__obf_d781958a0c55d323 := __obf_d94642eeb977270c - 1

	for __obf_a2ccc63c916c0db9 := __obf_d94642eeb977270c - 1; __obf_a2ccc63c916c0db9 >= 0; __obf_a2ccc63c916c0db9-- {
		if __obf_d27433cd224562e4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] != 0 {
			break
		}
		__obf_d781958a0c55d323 = __obf_a2ccc63c916c0db9 - 1
	}

	if __obf_d781958a0c55d323 < 0 {
		return __obf_71eb0882383cb0ef{}
	} else if __obf_d781958a0c55d323 < __obf_d94642eeb977270c-1 {
		__obf_d27433cd224562e4.__obf_ea46b5761247c020 = __obf_d27433cd224562e4.__obf_ea46b5761247c020[0 : __obf_d781958a0c55d323+1]
	}

	return __obf_d27433cd224562e4
}

func (__obf_d27433cd224562e4 __obf_71eb0882383cb0ef) string(__obf_2096d3eeb0e726d1 bool) string {
	var __obf_b9bf1b15d04a361b string
	__obf_d94642eeb977270c := __obf_d27433cd224562e4.__obf_d94642eeb977270c()

	for __obf_a2ccc63c916c0db9 := __obf_d94642eeb977270c - 1; __obf_a2ccc63c916c0db9 >= 0; __obf_a2ccc63c916c0db9-- {
		if __obf_d27433cd224562e4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] > 0 {
			if len(__obf_b9bf1b15d04a361b) > 0 {
				__obf_b9bf1b15d04a361b += " + "
			}

			if !__obf_2096d3eeb0e726d1 {
				__obf_b9bf1b15d04a361b += fmt.Sprintf("%dx^%d", __obf_d27433cd224562e4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9], __obf_a2ccc63c916c0db9)
			} else {
				__obf_b9bf1b15d04a361b += fmt.Sprintf("a^%dx^%d", __obf_a48718d8a46b86cc[__obf_d27433cd224562e4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9]], __obf_a2ccc63c916c0db9)
			}
		}
	}

	if len(__obf_b9bf1b15d04a361b) == 0 {
		__obf_b9bf1b15d04a361b = "0"
	}

	return __obf_b9bf1b15d04a361b
}

// equals returns true if e == other.
func (__obf_d27433cd224562e4 __obf_71eb0882383cb0ef) __obf_f8b91d99fb7263a4(__obf_8450914f6fe8235d __obf_71eb0882383cb0ef) bool {
	var __obf_be914efbbee9ca05 *__obf_71eb0882383cb0ef
	var __obf_f26de2a48bdf505d *__obf_71eb0882383cb0ef

	if __obf_d27433cd224562e4.__obf_d94642eeb977270c() > __obf_8450914f6fe8235d.__obf_d94642eeb977270c() {
		__obf_be914efbbee9ca05 = &__obf_8450914f6fe8235d
		__obf_f26de2a48bdf505d = &__obf_d27433cd224562e4
	} else {
		__obf_be914efbbee9ca05 = &__obf_d27433cd224562e4
		__obf_f26de2a48bdf505d = &__obf_8450914f6fe8235d
	}
	__obf_1d7170b02f2c64c7 := __obf_be914efbbee9ca05.__obf_d94642eeb977270c()
	__obf_b71e5b9e25bdb3d3 := __obf_f26de2a48bdf505d.__obf_d94642eeb977270c()

	for __obf_a2ccc63c916c0db9 := 0; __obf_a2ccc63c916c0db9 < __obf_1d7170b02f2c64c7; __obf_a2ccc63c916c0db9++ {
		if __obf_d27433cd224562e4.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] != __obf_8450914f6fe8235d.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] {
			return false
		}
	}

	for __obf_a2ccc63c916c0db9 := __obf_1d7170b02f2c64c7; __obf_a2ccc63c916c0db9 < __obf_b71e5b9e25bdb3d3; __obf_a2ccc63c916c0db9++ {
		if __obf_f26de2a48bdf505d.__obf_ea46b5761247c020[__obf_a2ccc63c916c0db9] != 0 {
			return false
		}
	}

	return true
}
