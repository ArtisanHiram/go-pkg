// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_13fa3c4d8cc2a5c1

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_087436cf3ba4af81 struct {
	__obf_90c5f727b336f89b []__obf_11b619b7d43c1a0d// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...

}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_b21c3e76a0cbcb0b(__obf_14fbbbd1ec149da3 *bitset.Bitset) __obf_087436cf3ba4af81 {
	__obf_def15252eb9409cf := __obf_14fbbbd1ec149da3.Len() / 8
	if __obf_14fbbbd1ec149da3.Len()%8 != 0 {
		__obf_def15252eb9409cf++
	}
	__obf_d8a1da44e57d6011 := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: make([]__obf_11b619b7d43c1a0d, __obf_def15252eb9409cf)}
	__obf_c4dd66f72f8a479f := __obf_def15252eb9409cf - 1
	for __obf_f0b271234bdec4ca := 0; __obf_f0b271234bdec4ca < __obf_14fbbbd1ec149da3.Len(); __obf_f0b271234bdec4ca += 8 {
		__obf_d8a1da44e57d6011.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] = __obf_11b619b7d43c1a0d(__obf_14fbbbd1ec149da3.ByteAt(__obf_f0b271234bdec4ca))
		__obf_c4dd66f72f8a479f--
	}

	return __obf_d8a1da44e57d6011
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_d07c34c23e218f0f(__obf_90c5f727b336f89b __obf_11b619b7d43c1a0d, __obf_b41724be7f1bc628 int) __obf_087436cf3ba4af81 {
	if __obf_90c5f727b336f89b == __obf_a890ef6f4a81fd07 {
		return __obf_087436cf3ba4af81{}
	}
	__obf_d8a1da44e57d6011 := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: make([]__obf_11b619b7d43c1a0d, __obf_b41724be7f1bc628+1)}
	__obf_d8a1da44e57d6011.__obf_90c5f727b336f89b[__obf_b41724be7f1bc628] = __obf_90c5f727b336f89b

	return __obf_d8a1da44e57d6011
}

func (__obf_f50a49059a731dcb __obf_087436cf3ba4af81) __obf_14fbbbd1ec149da3(__obf_03fc1f4c53891c31 int) []byte {
	__obf_d8a1da44e57d6011 := make([]byte, __obf_03fc1f4c53891c31)
	__obf_c4dd66f72f8a479f := __obf_03fc1f4c53891c31 - len(__obf_f50a49059a731dcb.__obf_90c5f727b336f89b)
	for __obf_f0b271234bdec4ca := len(__obf_f50a49059a731dcb.__obf_90c5f727b336f89b) - 1; __obf_f0b271234bdec4ca >= 0; __obf_f0b271234bdec4ca-- {
		__obf_d8a1da44e57d6011[__obf_c4dd66f72f8a479f] = byte(__obf_f50a49059a731dcb.__obf_90c5f727b336f89b[__obf_f0b271234bdec4ca])
		__obf_c4dd66f72f8a479f++
	}

	return __obf_d8a1da44e57d6011
}

// numTerms returns the number of
func (__obf_f50a49059a731dcb __obf_087436cf3ba4af81) __obf_03fc1f4c53891c31() int {
	return len(__obf_f50a49059a731dcb.

		// gfPolyMultiply returns a * b.
		__obf_90c5f727b336f89b)
}

func __obf_b680f3d50eea2b57(__obf_dd72989db1d736d6, __obf_d4d2022c93b1d338 __obf_087436cf3ba4af81) __obf_087436cf3ba4af81 {
	__obf_29afa6656efca9b8 := __obf_dd72989db1d736d6.__obf_03fc1f4c53891c31()
	__obf_c89e7464f62c5b94 := __obf_d4d2022c93b1d338.__obf_03fc1f4c53891c31()
	__obf_d8a1da44e57d6011 := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: make([]__obf_11b619b7d43c1a0d, __obf_29afa6656efca9b8+__obf_c89e7464f62c5b94)}

	for __obf_c4dd66f72f8a479f := 0; __obf_c4dd66f72f8a479f < __obf_29afa6656efca9b8; __obf_c4dd66f72f8a479f++ {
		for __obf_f0b271234bdec4ca := 0; __obf_f0b271234bdec4ca < __obf_c89e7464f62c5b94; __obf_f0b271234bdec4ca++ {
			if __obf_dd72989db1d736d6.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] != 0 && __obf_d4d2022c93b1d338.__obf_90c5f727b336f89b[__obf_f0b271234bdec4ca] != 0 {
				__obf_d6fc95dabf3aa764 := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: make([]__obf_11b619b7d43c1a0d, __obf_c4dd66f72f8a479f+__obf_f0b271234bdec4ca+1)}
				__obf_d6fc95dabf3aa764.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f+__obf_f0b271234bdec4ca] = __obf_00614e5ea538c8f7(__obf_dd72989db1d736d6.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f], __obf_d4d2022c93b1d338.__obf_90c5f727b336f89b[__obf_f0b271234bdec4ca])
				__obf_d8a1da44e57d6011 = __obf_7ad60108413aae90(__obf_d8a1da44e57d6011, __obf_d6fc95dabf3aa764)
			}
		}
	}

	return __obf_d8a1da44e57d6011.__obf_677413bf7f7b067a()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_b5870e3900b6d078(__obf_4fd11dac826620e5, __obf_afd3b3f3a3e4e985 __obf_087436cf3ba4af81) __obf_087436cf3ba4af81 {
	if __obf_afd3b3f3a3e4e985.__obf_7ee88c8d0109fc36(__obf_087436cf3ba4af81{}) {
		log.Panicln("Remainder by zero")
	}
	__obf_7f08354c0e804ad1 := __obf_4fd11dac826620e5

	for __obf_7f08354c0e804ad1.__obf_03fc1f4c53891c31() >= __obf_afd3b3f3a3e4e985.__obf_03fc1f4c53891c31() {
		__obf_b41724be7f1bc628 := __obf_7f08354c0e804ad1.__obf_03fc1f4c53891c31() - __obf_afd3b3f3a3e4e985.__obf_03fc1f4c53891c31()
		__obf_e093294829450480 := __obf_73faa4598e606f86(__obf_7f08354c0e804ad1.__obf_90c5f727b336f89b[__obf_7f08354c0e804ad1.__obf_03fc1f4c53891c31()-1], __obf_afd3b3f3a3e4e985.__obf_90c5f727b336f89b[__obf_afd3b3f3a3e4e985.__obf_03fc1f4c53891c31()-1])
		__obf_43a275e770bb6bb3 := __obf_b680f3d50eea2b57(__obf_afd3b3f3a3e4e985, __obf_d07c34c23e218f0f(__obf_e093294829450480, __obf_b41724be7f1bc628))
		__obf_7f08354c0e804ad1 = __obf_7ad60108413aae90(__obf_7f08354c0e804ad1, __obf_43a275e770bb6bb3)
	}

	return __obf_7f08354c0e804ad1.__obf_677413bf7f7b067a()
}

// gfPolyAdd returns a + b.
func __obf_7ad60108413aae90(__obf_dd72989db1d736d6, __obf_d4d2022c93b1d338 __obf_087436cf3ba4af81) __obf_087436cf3ba4af81 {
	__obf_29afa6656efca9b8 := __obf_dd72989db1d736d6.__obf_03fc1f4c53891c31()
	__obf_c89e7464f62c5b94 := __obf_d4d2022c93b1d338.__obf_03fc1f4c53891c31()
	__obf_03fc1f4c53891c31 := __obf_29afa6656efca9b8
	if __obf_c89e7464f62c5b94 > __obf_03fc1f4c53891c31 {
		__obf_03fc1f4c53891c31 = __obf_c89e7464f62c5b94
	}
	__obf_d8a1da44e57d6011 := __obf_087436cf3ba4af81{__obf_90c5f727b336f89b: make([]__obf_11b619b7d43c1a0d, __obf_03fc1f4c53891c31)}

	for __obf_c4dd66f72f8a479f := 0; __obf_c4dd66f72f8a479f < __obf_03fc1f4c53891c31; __obf_c4dd66f72f8a479f++ {
		switch {
		case __obf_29afa6656efca9b8 > __obf_c4dd66f72f8a479f && __obf_c89e7464f62c5b94 > __obf_c4dd66f72f8a479f:
			__obf_d8a1da44e57d6011.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] = __obf_f8850a0f4bacaf7c(__obf_dd72989db1d736d6.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f], __obf_d4d2022c93b1d338.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f])
		case __obf_29afa6656efca9b8 > __obf_c4dd66f72f8a479f:
			__obf_d8a1da44e57d6011.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] = __obf_dd72989db1d736d6.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f]
		default:
			__obf_d8a1da44e57d6011.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] = __obf_d4d2022c93b1d338.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f]
		}
	}

	return __obf_d8a1da44e57d6011.__obf_677413bf7f7b067a()
}

func (__obf_f50a49059a731dcb __obf_087436cf3ba4af81) __obf_677413bf7f7b067a() __obf_087436cf3ba4af81 {
	__obf_03fc1f4c53891c31 := __obf_f50a49059a731dcb.__obf_03fc1f4c53891c31()
	__obf_dda5948be69f243b := __obf_03fc1f4c53891c31 - 1

	for __obf_c4dd66f72f8a479f := __obf_03fc1f4c53891c31 - 1; __obf_c4dd66f72f8a479f >= 0; __obf_c4dd66f72f8a479f-- {
		if __obf_f50a49059a731dcb.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] != 0 {
			break
		}
		__obf_dda5948be69f243b = __obf_c4dd66f72f8a479f - 1
	}

	if __obf_dda5948be69f243b < 0 {
		return __obf_087436cf3ba4af81{}
	} else if __obf_dda5948be69f243b < __obf_03fc1f4c53891c31-1 {
		__obf_f50a49059a731dcb.__obf_90c5f727b336f89b = __obf_f50a49059a731dcb.__obf_90c5f727b336f89b[0 : __obf_dda5948be69f243b+1]
	}

	return __obf_f50a49059a731dcb
}

func (__obf_f50a49059a731dcb __obf_087436cf3ba4af81) string(__obf_db55f1337a506ed6 bool) string {
	var __obf_e62e4b92dcb1ccc7 string
	__obf_03fc1f4c53891c31 := __obf_f50a49059a731dcb.__obf_03fc1f4c53891c31()

	for __obf_c4dd66f72f8a479f := __obf_03fc1f4c53891c31 - 1; __obf_c4dd66f72f8a479f >= 0; __obf_c4dd66f72f8a479f-- {
		if __obf_f50a49059a731dcb.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] > 0 {
			if len(__obf_e62e4b92dcb1ccc7) > 0 {
				__obf_e62e4b92dcb1ccc7 += " + "
			}

			if !__obf_db55f1337a506ed6 {
				__obf_e62e4b92dcb1ccc7 += fmt.Sprintf("%dx^%d", __obf_f50a49059a731dcb.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f], __obf_c4dd66f72f8a479f)
			} else {
				__obf_e62e4b92dcb1ccc7 += fmt.Sprintf("a^%dx^%d", __obf_55b2ead7c4187484[__obf_f50a49059a731dcb.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f]], __obf_c4dd66f72f8a479f)
			}
		}
	}

	if len(__obf_e62e4b92dcb1ccc7) == 0 {
		__obf_e62e4b92dcb1ccc7 = "0"
	}

	return __obf_e62e4b92dcb1ccc7
}

// equals returns true if e == other.
func (__obf_f50a49059a731dcb __obf_087436cf3ba4af81) __obf_7ee88c8d0109fc36(__obf_45b84e82bcbcf6a8 __obf_087436cf3ba4af81) bool {
	var __obf_5d8d08f2beae24ec *__obf_087436cf3ba4af81
	var __obf_e5c0f56a12772d2c *__obf_087436cf3ba4af81

	if __obf_f50a49059a731dcb.__obf_03fc1f4c53891c31() > __obf_45b84e82bcbcf6a8.__obf_03fc1f4c53891c31() {
		__obf_5d8d08f2beae24ec = &__obf_45b84e82bcbcf6a8
		__obf_e5c0f56a12772d2c = &__obf_f50a49059a731dcb
	} else {
		__obf_5d8d08f2beae24ec = &__obf_f50a49059a731dcb
		__obf_e5c0f56a12772d2c = &__obf_45b84e82bcbcf6a8
	}
	__obf_e94cf12bfbd75513 := __obf_5d8d08f2beae24ec.__obf_03fc1f4c53891c31()
	__obf_e7fe2ac82b07df69 := __obf_e5c0f56a12772d2c.__obf_03fc1f4c53891c31()

	for __obf_c4dd66f72f8a479f := 0; __obf_c4dd66f72f8a479f < __obf_e94cf12bfbd75513; __obf_c4dd66f72f8a479f++ {
		if __obf_f50a49059a731dcb.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] != __obf_45b84e82bcbcf6a8.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] {
			return false
		}
	}

	for __obf_c4dd66f72f8a479f := __obf_e94cf12bfbd75513; __obf_c4dd66f72f8a479f < __obf_e7fe2ac82b07df69; __obf_c4dd66f72f8a479f++ {
		if __obf_e5c0f56a12772d2c.__obf_90c5f727b336f89b[__obf_c4dd66f72f8a479f] != 0 {
			return false
		}
	}

	return true
}
