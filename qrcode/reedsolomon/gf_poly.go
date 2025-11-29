// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_79572f72dbe37a0e

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_c9418fcf8fc08b00 struct {
	__obf_887f3c7bd44a27c9 []__obf_21ecc2aea3abd05b// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...

}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_1c8448dc38dc8598(__obf_a8518aca69865c96 *bitset.Bitset) __obf_c9418fcf8fc08b00 {
	__obf_dbd87879bee2c6b2 := __obf_a8518aca69865c96.Len() / 8
	if __obf_a8518aca69865c96.Len()%8 != 0 {
		__obf_dbd87879bee2c6b2++
	}
	__obf_8214529cf7250f23 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: make([]__obf_21ecc2aea3abd05b, __obf_dbd87879bee2c6b2)}
	__obf_d9d01cdc0f623246 := __obf_dbd87879bee2c6b2 - 1
	for __obf_3a706b6ca0fee8f4 := 0; __obf_3a706b6ca0fee8f4 < __obf_a8518aca69865c96.Len(); __obf_3a706b6ca0fee8f4 += 8 {
		__obf_8214529cf7250f23.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] = __obf_21ecc2aea3abd05b(__obf_a8518aca69865c96.ByteAt(__obf_3a706b6ca0fee8f4))
		__obf_d9d01cdc0f623246--
	}

	return __obf_8214529cf7250f23
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_ff98f1989ce99555(__obf_887f3c7bd44a27c9 __obf_21ecc2aea3abd05b, __obf_a8bb449e9bfb0f73 int) __obf_c9418fcf8fc08b00 {
	if __obf_887f3c7bd44a27c9 == __obf_d03cb957801a9f11 {
		return __obf_c9418fcf8fc08b00{}
	}
	__obf_8214529cf7250f23 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: make([]__obf_21ecc2aea3abd05b, __obf_a8bb449e9bfb0f73+1)}
	__obf_8214529cf7250f23.__obf_887f3c7bd44a27c9[__obf_a8bb449e9bfb0f73] = __obf_887f3c7bd44a27c9

	return __obf_8214529cf7250f23
}

func (__obf_26971b6a91e72a5f __obf_c9418fcf8fc08b00) __obf_a8518aca69865c96(__obf_369de7d395a78ca1 int) []byte {
	__obf_8214529cf7250f23 := make([]byte, __obf_369de7d395a78ca1)
	__obf_d9d01cdc0f623246 := __obf_369de7d395a78ca1 - len(__obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9)
	for __obf_3a706b6ca0fee8f4 := len(__obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9) - 1; __obf_3a706b6ca0fee8f4 >= 0; __obf_3a706b6ca0fee8f4-- {
		__obf_8214529cf7250f23[__obf_d9d01cdc0f623246] = byte(__obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[__obf_3a706b6ca0fee8f4])
		__obf_d9d01cdc0f623246++
	}

	return __obf_8214529cf7250f23
}

// numTerms returns the number of
func (__obf_26971b6a91e72a5f __obf_c9418fcf8fc08b00) __obf_369de7d395a78ca1() int {
	return len(__obf_26971b6a91e72a5f.

		// gfPolyMultiply returns a * b.
		__obf_887f3c7bd44a27c9)
}

func __obf_5bf6d069e50e9834(__obf_6da1e839dbadb4c1, __obf_7961952f8342be7e __obf_c9418fcf8fc08b00) __obf_c9418fcf8fc08b00 {
	__obf_171cc6187d955573 := __obf_6da1e839dbadb4c1.__obf_369de7d395a78ca1()
	__obf_98813834e5a45e52 := __obf_7961952f8342be7e.__obf_369de7d395a78ca1()
	__obf_8214529cf7250f23 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: make([]__obf_21ecc2aea3abd05b, __obf_171cc6187d955573+__obf_98813834e5a45e52)}

	for __obf_d9d01cdc0f623246 := 0; __obf_d9d01cdc0f623246 < __obf_171cc6187d955573; __obf_d9d01cdc0f623246++ {
		for __obf_3a706b6ca0fee8f4 := 0; __obf_3a706b6ca0fee8f4 < __obf_98813834e5a45e52; __obf_3a706b6ca0fee8f4++ {
			if __obf_6da1e839dbadb4c1.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] != 0 && __obf_7961952f8342be7e.__obf_887f3c7bd44a27c9[__obf_3a706b6ca0fee8f4] != 0 {
				__obf_602e708fb4b170e5 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: make([]__obf_21ecc2aea3abd05b, __obf_d9d01cdc0f623246+__obf_3a706b6ca0fee8f4+1)}
				__obf_602e708fb4b170e5.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246+__obf_3a706b6ca0fee8f4] = __obf_564e8c977a044222(__obf_6da1e839dbadb4c1.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246], __obf_7961952f8342be7e.__obf_887f3c7bd44a27c9[__obf_3a706b6ca0fee8f4])
				__obf_8214529cf7250f23 = __obf_267a7e744db90984(__obf_8214529cf7250f23, __obf_602e708fb4b170e5)
			}
		}
	}

	return __obf_8214529cf7250f23.__obf_357b917a4818946a()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_27b66f0b936c98aa(__obf_33e3e4ecb9129526, __obf_270ff0f946185a75 __obf_c9418fcf8fc08b00) __obf_c9418fcf8fc08b00 {
	if __obf_270ff0f946185a75.__obf_4fd4f4c5b918dde7(__obf_c9418fcf8fc08b00{}) {
		log.Panicln("Remainder by zero")
	}
	__obf_f435504415a4dac9 := __obf_33e3e4ecb9129526

	for __obf_f435504415a4dac9.__obf_369de7d395a78ca1() >= __obf_270ff0f946185a75.__obf_369de7d395a78ca1() {
		__obf_a8bb449e9bfb0f73 := __obf_f435504415a4dac9.__obf_369de7d395a78ca1() - __obf_270ff0f946185a75.__obf_369de7d395a78ca1()
		__obf_dc8858df484475fd := __obf_e7852300ea89ef86(__obf_f435504415a4dac9.__obf_887f3c7bd44a27c9[__obf_f435504415a4dac9.__obf_369de7d395a78ca1()-1], __obf_270ff0f946185a75.__obf_887f3c7bd44a27c9[__obf_270ff0f946185a75.__obf_369de7d395a78ca1()-1])
		__obf_7d44ffc17fd93131 := __obf_5bf6d069e50e9834(__obf_270ff0f946185a75, __obf_ff98f1989ce99555(__obf_dc8858df484475fd, __obf_a8bb449e9bfb0f73))
		__obf_f435504415a4dac9 = __obf_267a7e744db90984(__obf_f435504415a4dac9, __obf_7d44ffc17fd93131)
	}

	return __obf_f435504415a4dac9.__obf_357b917a4818946a()
}

// gfPolyAdd returns a + b.
func __obf_267a7e744db90984(__obf_6da1e839dbadb4c1, __obf_7961952f8342be7e __obf_c9418fcf8fc08b00) __obf_c9418fcf8fc08b00 {
	__obf_171cc6187d955573 := __obf_6da1e839dbadb4c1.__obf_369de7d395a78ca1()
	__obf_98813834e5a45e52 := __obf_7961952f8342be7e.__obf_369de7d395a78ca1()
	__obf_369de7d395a78ca1 := __obf_171cc6187d955573
	if __obf_98813834e5a45e52 > __obf_369de7d395a78ca1 {
		__obf_369de7d395a78ca1 = __obf_98813834e5a45e52
	}
	__obf_8214529cf7250f23 := __obf_c9418fcf8fc08b00{__obf_887f3c7bd44a27c9: make([]__obf_21ecc2aea3abd05b, __obf_369de7d395a78ca1)}

	for __obf_d9d01cdc0f623246 := 0; __obf_d9d01cdc0f623246 < __obf_369de7d395a78ca1; __obf_d9d01cdc0f623246++ {
		switch {
		case __obf_171cc6187d955573 > __obf_d9d01cdc0f623246 && __obf_98813834e5a45e52 > __obf_d9d01cdc0f623246:
			__obf_8214529cf7250f23.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] = __obf_63638239f9b31a62(__obf_6da1e839dbadb4c1.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246], __obf_7961952f8342be7e.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246])
		case __obf_171cc6187d955573 > __obf_d9d01cdc0f623246:
			__obf_8214529cf7250f23.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] = __obf_6da1e839dbadb4c1.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246]
		default:
			__obf_8214529cf7250f23.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] = __obf_7961952f8342be7e.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246]
		}
	}

	return __obf_8214529cf7250f23.__obf_357b917a4818946a()
}

func (__obf_26971b6a91e72a5f __obf_c9418fcf8fc08b00) __obf_357b917a4818946a() __obf_c9418fcf8fc08b00 {
	__obf_369de7d395a78ca1 := __obf_26971b6a91e72a5f.__obf_369de7d395a78ca1()
	__obf_29c8f9004c1b0676 := __obf_369de7d395a78ca1 - 1

	for __obf_d9d01cdc0f623246 := __obf_369de7d395a78ca1 - 1; __obf_d9d01cdc0f623246 >= 0; __obf_d9d01cdc0f623246-- {
		if __obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] != 0 {
			break
		}
		__obf_29c8f9004c1b0676 = __obf_d9d01cdc0f623246 - 1
	}

	if __obf_29c8f9004c1b0676 < 0 {
		return __obf_c9418fcf8fc08b00{}
	} else if __obf_29c8f9004c1b0676 < __obf_369de7d395a78ca1-1 {
		__obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9 = __obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[0 : __obf_29c8f9004c1b0676+1]
	}

	return __obf_26971b6a91e72a5f
}

func (__obf_26971b6a91e72a5f __obf_c9418fcf8fc08b00) string(__obf_fbfc12fa38413553 bool) string {
	var __obf_2cdc81be242376c9 string
	__obf_369de7d395a78ca1 := __obf_26971b6a91e72a5f.__obf_369de7d395a78ca1()

	for __obf_d9d01cdc0f623246 := __obf_369de7d395a78ca1 - 1; __obf_d9d01cdc0f623246 >= 0; __obf_d9d01cdc0f623246-- {
		if __obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] > 0 {
			if len(__obf_2cdc81be242376c9) > 0 {
				__obf_2cdc81be242376c9 += " + "
			}

			if !__obf_fbfc12fa38413553 {
				__obf_2cdc81be242376c9 += fmt.Sprintf("%dx^%d", __obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246], __obf_d9d01cdc0f623246)
			} else {
				__obf_2cdc81be242376c9 += fmt.Sprintf("a^%dx^%d", __obf_8ba5fa7cb9892d7f[__obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246]], __obf_d9d01cdc0f623246)
			}
		}
	}

	if len(__obf_2cdc81be242376c9) == 0 {
		__obf_2cdc81be242376c9 = "0"
	}

	return __obf_2cdc81be242376c9
}

// equals returns true if e == other.
func (__obf_26971b6a91e72a5f __obf_c9418fcf8fc08b00) __obf_4fd4f4c5b918dde7(__obf_cf8d5ace41daf6c6 __obf_c9418fcf8fc08b00) bool {
	var __obf_ca41d4020f8cf3e4 *__obf_c9418fcf8fc08b00
	var __obf_d20d1fbbd2f84cdd *__obf_c9418fcf8fc08b00

	if __obf_26971b6a91e72a5f.__obf_369de7d395a78ca1() > __obf_cf8d5ace41daf6c6.__obf_369de7d395a78ca1() {
		__obf_ca41d4020f8cf3e4 = &__obf_cf8d5ace41daf6c6
		__obf_d20d1fbbd2f84cdd = &__obf_26971b6a91e72a5f
	} else {
		__obf_ca41d4020f8cf3e4 = &__obf_26971b6a91e72a5f
		__obf_d20d1fbbd2f84cdd = &__obf_cf8d5ace41daf6c6
	}
	__obf_5ddac53ecf610330 := __obf_ca41d4020f8cf3e4.__obf_369de7d395a78ca1()
	__obf_91666f77f664bad5 := __obf_d20d1fbbd2f84cdd.__obf_369de7d395a78ca1()

	for __obf_d9d01cdc0f623246 := 0; __obf_d9d01cdc0f623246 < __obf_5ddac53ecf610330; __obf_d9d01cdc0f623246++ {
		if __obf_26971b6a91e72a5f.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] != __obf_cf8d5ace41daf6c6.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] {
			return false
		}
	}

	for __obf_d9d01cdc0f623246 := __obf_5ddac53ecf610330; __obf_d9d01cdc0f623246 < __obf_91666f77f664bad5; __obf_d9d01cdc0f623246++ {
		if __obf_d20d1fbbd2f84cdd.__obf_887f3c7bd44a27c9[__obf_d9d01cdc0f623246] != 0 {
			return false
		}
	}

	return true
}
