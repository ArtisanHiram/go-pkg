// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_6d5da6b2ca65c82e

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_40186e28e54c6d61 struct {
	// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...
	__obf_747aa4fb15a9f618 []__obf_29e803ad00f0d12b
}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_daafd0c6eef1f175(__obf_92f02e566ccfb5ff *bitset.Bitset) __obf_40186e28e54c6d61 {
	__obf_ffa3ff3a0e5c97f5 := __obf_92f02e566ccfb5ff.Len() / 8
	if __obf_92f02e566ccfb5ff.Len()%8 != 0 {
		__obf_ffa3ff3a0e5c97f5++
	}

	__obf_b8d4231c43bc645b := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: make([]__obf_29e803ad00f0d12b, __obf_ffa3ff3a0e5c97f5)}

	__obf_85aff0c19ced1f98 := __obf_ffa3ff3a0e5c97f5 - 1
	for __obf_57bc634e41a81dbd := 0; __obf_57bc634e41a81dbd < __obf_92f02e566ccfb5ff.Len(); __obf_57bc634e41a81dbd += 8 {
		__obf_b8d4231c43bc645b.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] = __obf_29e803ad00f0d12b(__obf_92f02e566ccfb5ff.ByteAt(__obf_57bc634e41a81dbd))
		__obf_85aff0c19ced1f98--
	}

	return __obf_b8d4231c43bc645b
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_3f8089776f332a11(__obf_747aa4fb15a9f618 __obf_29e803ad00f0d12b, __obf_d68be180ca1279c1 int) __obf_40186e28e54c6d61 {
	if __obf_747aa4fb15a9f618 == __obf_c3429c035d16c395 {
		return __obf_40186e28e54c6d61{}
	}

	__obf_b8d4231c43bc645b := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: make([]__obf_29e803ad00f0d12b, __obf_d68be180ca1279c1+1)}
	__obf_b8d4231c43bc645b.__obf_747aa4fb15a9f618[__obf_d68be180ca1279c1] = __obf_747aa4fb15a9f618

	return __obf_b8d4231c43bc645b
}

func (__obf_0374b33c8b7512ac __obf_40186e28e54c6d61) __obf_92f02e566ccfb5ff(__obf_93a440d3373e34e4 int) []byte {
	__obf_b8d4231c43bc645b := make([]byte, __obf_93a440d3373e34e4)

	__obf_85aff0c19ced1f98 := __obf_93a440d3373e34e4 - len(__obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618)
	for __obf_57bc634e41a81dbd := len(__obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618) - 1; __obf_57bc634e41a81dbd >= 0; __obf_57bc634e41a81dbd-- {
		__obf_b8d4231c43bc645b[__obf_85aff0c19ced1f98] = byte(__obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[__obf_57bc634e41a81dbd])
		__obf_85aff0c19ced1f98++
	}

	return __obf_b8d4231c43bc645b
}

// numTerms returns the number of
func (__obf_0374b33c8b7512ac __obf_40186e28e54c6d61) __obf_93a440d3373e34e4() int {
	return len(__obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618)
}

// gfPolyMultiply returns a * b.
func __obf_227da40f9c6c41f3(__obf_a5e0b8636672b1a1, __obf_d86b9025cb5ef594 __obf_40186e28e54c6d61) __obf_40186e28e54c6d61 {
	__obf_827228b07b687b6a := __obf_a5e0b8636672b1a1.__obf_93a440d3373e34e4()
	__obf_cfa95afdcf26bf90 := __obf_d86b9025cb5ef594.__obf_93a440d3373e34e4()

	__obf_b8d4231c43bc645b := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: make([]__obf_29e803ad00f0d12b, __obf_827228b07b687b6a+__obf_cfa95afdcf26bf90)}

	for __obf_85aff0c19ced1f98 := 0; __obf_85aff0c19ced1f98 < __obf_827228b07b687b6a; __obf_85aff0c19ced1f98++ {
		for __obf_57bc634e41a81dbd := 0; __obf_57bc634e41a81dbd < __obf_cfa95afdcf26bf90; __obf_57bc634e41a81dbd++ {
			if __obf_a5e0b8636672b1a1.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] != 0 && __obf_d86b9025cb5ef594.__obf_747aa4fb15a9f618[__obf_57bc634e41a81dbd] != 0 {
				__obf_f46a96e11f0703b8 := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: make([]__obf_29e803ad00f0d12b, __obf_85aff0c19ced1f98+__obf_57bc634e41a81dbd+1)}
				__obf_f46a96e11f0703b8.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98+__obf_57bc634e41a81dbd] = __obf_bd0138baa1f450bb(__obf_a5e0b8636672b1a1.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98], __obf_d86b9025cb5ef594.__obf_747aa4fb15a9f618[__obf_57bc634e41a81dbd])

				__obf_b8d4231c43bc645b = __obf_2db7af224534dec2(__obf_b8d4231c43bc645b, __obf_f46a96e11f0703b8)
			}
		}
	}

	return __obf_b8d4231c43bc645b.__obf_c66b1ac4af88d44a()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_888a3a462522e873(__obf_defc973ea49a9ed5, __obf_c4e9eee963aa65b3 __obf_40186e28e54c6d61) __obf_40186e28e54c6d61 {
	if __obf_c4e9eee963aa65b3.__obf_571f59554b81dbd1(__obf_40186e28e54c6d61{}) {
		log.Panicln("Remainder by zero")
	}

	__obf_8977186fbcc7b181 := __obf_defc973ea49a9ed5

	for __obf_8977186fbcc7b181.__obf_93a440d3373e34e4() >= __obf_c4e9eee963aa65b3.__obf_93a440d3373e34e4() {
		__obf_d68be180ca1279c1 := __obf_8977186fbcc7b181.__obf_93a440d3373e34e4() - __obf_c4e9eee963aa65b3.__obf_93a440d3373e34e4()
		__obf_47534168a9f6d5bf := __obf_3c69a8ec4ab6d915(__obf_8977186fbcc7b181.__obf_747aa4fb15a9f618[__obf_8977186fbcc7b181.__obf_93a440d3373e34e4()-1],
			__obf_c4e9eee963aa65b3.__obf_747aa4fb15a9f618[__obf_c4e9eee963aa65b3.__obf_93a440d3373e34e4()-1])

		__obf_43ab67d48f59044d := __obf_227da40f9c6c41f3(__obf_c4e9eee963aa65b3,
			__obf_3f8089776f332a11(__obf_47534168a9f6d5bf, __obf_d68be180ca1279c1))

		__obf_8977186fbcc7b181 = __obf_2db7af224534dec2(__obf_8977186fbcc7b181, __obf_43ab67d48f59044d)
	}

	return __obf_8977186fbcc7b181.__obf_c66b1ac4af88d44a()
}

// gfPolyAdd returns a + b.
func __obf_2db7af224534dec2(__obf_a5e0b8636672b1a1, __obf_d86b9025cb5ef594 __obf_40186e28e54c6d61) __obf_40186e28e54c6d61 {
	__obf_827228b07b687b6a := __obf_a5e0b8636672b1a1.__obf_93a440d3373e34e4()
	__obf_cfa95afdcf26bf90 := __obf_d86b9025cb5ef594.__obf_93a440d3373e34e4()

	__obf_93a440d3373e34e4 := __obf_827228b07b687b6a
	if __obf_cfa95afdcf26bf90 > __obf_93a440d3373e34e4 {
		__obf_93a440d3373e34e4 = __obf_cfa95afdcf26bf90
	}

	__obf_b8d4231c43bc645b := __obf_40186e28e54c6d61{__obf_747aa4fb15a9f618: make([]__obf_29e803ad00f0d12b, __obf_93a440d3373e34e4)}

	for __obf_85aff0c19ced1f98 := 0; __obf_85aff0c19ced1f98 < __obf_93a440d3373e34e4; __obf_85aff0c19ced1f98++ {
		switch {
		case __obf_827228b07b687b6a > __obf_85aff0c19ced1f98 && __obf_cfa95afdcf26bf90 > __obf_85aff0c19ced1f98:
			__obf_b8d4231c43bc645b.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] = __obf_1517ca40aa775c29(__obf_a5e0b8636672b1a1.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98], __obf_d86b9025cb5ef594.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98])
		case __obf_827228b07b687b6a > __obf_85aff0c19ced1f98:
			__obf_b8d4231c43bc645b.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] = __obf_a5e0b8636672b1a1.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98]
		default:
			__obf_b8d4231c43bc645b.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] = __obf_d86b9025cb5ef594.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98]
		}
	}

	return __obf_b8d4231c43bc645b.__obf_c66b1ac4af88d44a()
}

func (__obf_0374b33c8b7512ac __obf_40186e28e54c6d61) __obf_c66b1ac4af88d44a() __obf_40186e28e54c6d61 {
	__obf_93a440d3373e34e4 := __obf_0374b33c8b7512ac.__obf_93a440d3373e34e4()
	__obf_0092f7513469573b := __obf_93a440d3373e34e4 - 1

	for __obf_85aff0c19ced1f98 := __obf_93a440d3373e34e4 - 1; __obf_85aff0c19ced1f98 >= 0; __obf_85aff0c19ced1f98-- {
		if __obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] != 0 {
			break
		}

		__obf_0092f7513469573b = __obf_85aff0c19ced1f98 - 1
	}

	if __obf_0092f7513469573b < 0 {
		return __obf_40186e28e54c6d61{}
	} else if __obf_0092f7513469573b < __obf_93a440d3373e34e4-1 {
		__obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618 = __obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[0 : __obf_0092f7513469573b+1]
	}

	return __obf_0374b33c8b7512ac
}

func (__obf_0374b33c8b7512ac __obf_40186e28e54c6d61) string(__obf_e0f81d45b74cc0cc bool) string {
	var __obf_52ba994ea1ffee01 string
	__obf_93a440d3373e34e4 := __obf_0374b33c8b7512ac.__obf_93a440d3373e34e4()

	for __obf_85aff0c19ced1f98 := __obf_93a440d3373e34e4 - 1; __obf_85aff0c19ced1f98 >= 0; __obf_85aff0c19ced1f98-- {
		if __obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] > 0 {
			if len(__obf_52ba994ea1ffee01) > 0 {
				__obf_52ba994ea1ffee01 += " + "
			}

			if !__obf_e0f81d45b74cc0cc {
				__obf_52ba994ea1ffee01 += fmt.Sprintf("%dx^%d", __obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98], __obf_85aff0c19ced1f98)
			} else {
				__obf_52ba994ea1ffee01 += fmt.Sprintf("a^%dx^%d", __obf_922a3b54cf565bb2[__obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98]], __obf_85aff0c19ced1f98)
			}
		}
	}

	if len(__obf_52ba994ea1ffee01) == 0 {
		__obf_52ba994ea1ffee01 = "0"
	}

	return __obf_52ba994ea1ffee01
}

// equals returns true if e == other.
func (__obf_0374b33c8b7512ac __obf_40186e28e54c6d61) __obf_571f59554b81dbd1(__obf_f1480ca02be3ed9c __obf_40186e28e54c6d61) bool {
	var __obf_4a1d505f07383275 *__obf_40186e28e54c6d61
	var __obf_b822def49db881e0 *__obf_40186e28e54c6d61

	if __obf_0374b33c8b7512ac.__obf_93a440d3373e34e4() > __obf_f1480ca02be3ed9c.__obf_93a440d3373e34e4() {
		__obf_4a1d505f07383275 = &__obf_f1480ca02be3ed9c
		__obf_b822def49db881e0 = &__obf_0374b33c8b7512ac
	} else {
		__obf_4a1d505f07383275 = &__obf_0374b33c8b7512ac
		__obf_b822def49db881e0 = &__obf_f1480ca02be3ed9c
	}

	__obf_04d275aefc349455 := __obf_4a1d505f07383275.__obf_93a440d3373e34e4()
	__obf_be2a8d82e3f1a53b := __obf_b822def49db881e0.__obf_93a440d3373e34e4()

	for __obf_85aff0c19ced1f98 := 0; __obf_85aff0c19ced1f98 < __obf_04d275aefc349455; __obf_85aff0c19ced1f98++ {
		if __obf_0374b33c8b7512ac.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] != __obf_f1480ca02be3ed9c.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] {
			return false
		}
	}

	for __obf_85aff0c19ced1f98 := __obf_04d275aefc349455; __obf_85aff0c19ced1f98 < __obf_be2a8d82e3f1a53b; __obf_85aff0c19ced1f98++ {
		if __obf_b822def49db881e0.__obf_747aa4fb15a9f618[__obf_85aff0c19ced1f98] != 0 {
			return false
		}
	}

	return true
}
