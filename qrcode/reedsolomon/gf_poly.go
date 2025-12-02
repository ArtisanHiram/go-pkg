// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_e1f3dcd8d719d52b

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_7f0a87df1252e0f1 struct {
	__obf_dc2d266987a80f84 []__obf_69935e5d788798dd// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...

}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_8c8e622b7e5d1808(__obf_e84c8ad5f9fe4eba *bitset.Bitset) __obf_7f0a87df1252e0f1 {
	__obf_e615b321e44f68ad := __obf_e84c8ad5f9fe4eba.Len() / 8
	if __obf_e84c8ad5f9fe4eba.Len()%8 != 0 {
		__obf_e615b321e44f68ad++
	}
	__obf_903beff2487e2326 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: make([]__obf_69935e5d788798dd, __obf_e615b321e44f68ad)}
	__obf_13b2631ac20e0422 := __obf_e615b321e44f68ad - 1
	for __obf_ddc14dcaed304a5d := 0; __obf_ddc14dcaed304a5d < __obf_e84c8ad5f9fe4eba.Len(); __obf_ddc14dcaed304a5d += 8 {
		__obf_903beff2487e2326.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] = __obf_69935e5d788798dd(__obf_e84c8ad5f9fe4eba.ByteAt(__obf_ddc14dcaed304a5d))
		__obf_13b2631ac20e0422--
	}

	return __obf_903beff2487e2326
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_52f952075352cc74(__obf_dc2d266987a80f84 __obf_69935e5d788798dd, __obf_eee26701a2608609 int) __obf_7f0a87df1252e0f1 {
	if __obf_dc2d266987a80f84 == __obf_9b11e26e4ca2739c {
		return __obf_7f0a87df1252e0f1{}
	}
	__obf_903beff2487e2326 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: make([]__obf_69935e5d788798dd, __obf_eee26701a2608609+1)}
	__obf_903beff2487e2326.__obf_dc2d266987a80f84[__obf_eee26701a2608609] = __obf_dc2d266987a80f84

	return __obf_903beff2487e2326
}

func (__obf_6a222ce425464583 __obf_7f0a87df1252e0f1) __obf_e84c8ad5f9fe4eba(__obf_90986b368b7bf87d int) []byte {
	__obf_903beff2487e2326 := make([]byte, __obf_90986b368b7bf87d)
	__obf_13b2631ac20e0422 := __obf_90986b368b7bf87d - len(__obf_6a222ce425464583.__obf_dc2d266987a80f84)
	for __obf_ddc14dcaed304a5d := len(__obf_6a222ce425464583.__obf_dc2d266987a80f84) - 1; __obf_ddc14dcaed304a5d >= 0; __obf_ddc14dcaed304a5d-- {
		__obf_903beff2487e2326[__obf_13b2631ac20e0422] = byte(__obf_6a222ce425464583.__obf_dc2d266987a80f84[__obf_ddc14dcaed304a5d])
		__obf_13b2631ac20e0422++
	}

	return __obf_903beff2487e2326
}

// numTerms returns the number of
func (__obf_6a222ce425464583 __obf_7f0a87df1252e0f1) __obf_90986b368b7bf87d() int {
	return len(__obf_6a222ce425464583.

		// gfPolyMultiply returns a * b.
		__obf_dc2d266987a80f84)
}

func __obf_64c6d4eebca914e2(__obf_dda00ebd89ffcf1f, __obf_718f5317ccb4d321 __obf_7f0a87df1252e0f1) __obf_7f0a87df1252e0f1 {
	__obf_d1e89ed61e7d23fa := __obf_dda00ebd89ffcf1f.__obf_90986b368b7bf87d()
	__obf_19b2386223438133 := __obf_718f5317ccb4d321.__obf_90986b368b7bf87d()
	__obf_903beff2487e2326 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: make([]__obf_69935e5d788798dd, __obf_d1e89ed61e7d23fa+__obf_19b2386223438133)}

	for __obf_13b2631ac20e0422 := 0; __obf_13b2631ac20e0422 < __obf_d1e89ed61e7d23fa; __obf_13b2631ac20e0422++ {
		for __obf_ddc14dcaed304a5d := 0; __obf_ddc14dcaed304a5d < __obf_19b2386223438133; __obf_ddc14dcaed304a5d++ {
			if __obf_dda00ebd89ffcf1f.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] != 0 && __obf_718f5317ccb4d321.__obf_dc2d266987a80f84[__obf_ddc14dcaed304a5d] != 0 {
				__obf_36c400076cb952f6 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: make([]__obf_69935e5d788798dd, __obf_13b2631ac20e0422+__obf_ddc14dcaed304a5d+1)}
				__obf_36c400076cb952f6.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422+__obf_ddc14dcaed304a5d] = __obf_30eb8b7025b8e913(__obf_dda00ebd89ffcf1f.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422], __obf_718f5317ccb4d321.__obf_dc2d266987a80f84[__obf_ddc14dcaed304a5d])
				__obf_903beff2487e2326 = __obf_1e7cf3ed2bd1b483(__obf_903beff2487e2326, __obf_36c400076cb952f6)
			}
		}
	}

	return __obf_903beff2487e2326.__obf_48c6b18751f9de7f()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_b38b77e4cdfbc62a(__obf_fdb22583663e49d9, __obf_8116d1ea03372de5 __obf_7f0a87df1252e0f1) __obf_7f0a87df1252e0f1 {
	if __obf_8116d1ea03372de5.__obf_94ba6c94c165b9b4(__obf_7f0a87df1252e0f1{}) {
		log.Panicln("Remainder by zero")
	}
	__obf_0fed35ba81544ce5 := __obf_fdb22583663e49d9

	for __obf_0fed35ba81544ce5.__obf_90986b368b7bf87d() >= __obf_8116d1ea03372de5.__obf_90986b368b7bf87d() {
		__obf_eee26701a2608609 := __obf_0fed35ba81544ce5.__obf_90986b368b7bf87d() - __obf_8116d1ea03372de5.__obf_90986b368b7bf87d()
		__obf_b242eb318f26341a := __obf_88bbe75c909f5125(__obf_0fed35ba81544ce5.__obf_dc2d266987a80f84[__obf_0fed35ba81544ce5.__obf_90986b368b7bf87d()-1], __obf_8116d1ea03372de5.__obf_dc2d266987a80f84[__obf_8116d1ea03372de5.__obf_90986b368b7bf87d()-1])
		__obf_279ccf1d8a18419e := __obf_64c6d4eebca914e2(__obf_8116d1ea03372de5, __obf_52f952075352cc74(__obf_b242eb318f26341a, __obf_eee26701a2608609))
		__obf_0fed35ba81544ce5 = __obf_1e7cf3ed2bd1b483(__obf_0fed35ba81544ce5, __obf_279ccf1d8a18419e)
	}

	return __obf_0fed35ba81544ce5.__obf_48c6b18751f9de7f()
}

// gfPolyAdd returns a + b.
func __obf_1e7cf3ed2bd1b483(__obf_dda00ebd89ffcf1f, __obf_718f5317ccb4d321 __obf_7f0a87df1252e0f1) __obf_7f0a87df1252e0f1 {
	__obf_d1e89ed61e7d23fa := __obf_dda00ebd89ffcf1f.__obf_90986b368b7bf87d()
	__obf_19b2386223438133 := __obf_718f5317ccb4d321.__obf_90986b368b7bf87d()
	__obf_90986b368b7bf87d := __obf_d1e89ed61e7d23fa
	if __obf_19b2386223438133 > __obf_90986b368b7bf87d {
		__obf_90986b368b7bf87d = __obf_19b2386223438133
	}
	__obf_903beff2487e2326 := __obf_7f0a87df1252e0f1{__obf_dc2d266987a80f84: make([]__obf_69935e5d788798dd, __obf_90986b368b7bf87d)}

	for __obf_13b2631ac20e0422 := 0; __obf_13b2631ac20e0422 < __obf_90986b368b7bf87d; __obf_13b2631ac20e0422++ {
		switch {
		case __obf_d1e89ed61e7d23fa > __obf_13b2631ac20e0422 && __obf_19b2386223438133 > __obf_13b2631ac20e0422:
			__obf_903beff2487e2326.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] = __obf_d7c3c7b13276edea(__obf_dda00ebd89ffcf1f.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422], __obf_718f5317ccb4d321.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422])
		case __obf_d1e89ed61e7d23fa > __obf_13b2631ac20e0422:
			__obf_903beff2487e2326.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] = __obf_dda00ebd89ffcf1f.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422]
		default:
			__obf_903beff2487e2326.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] = __obf_718f5317ccb4d321.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422]
		}
	}

	return __obf_903beff2487e2326.__obf_48c6b18751f9de7f()
}

func (__obf_6a222ce425464583 __obf_7f0a87df1252e0f1) __obf_48c6b18751f9de7f() __obf_7f0a87df1252e0f1 {
	__obf_90986b368b7bf87d := __obf_6a222ce425464583.__obf_90986b368b7bf87d()
	__obf_1b5d2df4fcb389df := __obf_90986b368b7bf87d - 1

	for __obf_13b2631ac20e0422 := __obf_90986b368b7bf87d - 1; __obf_13b2631ac20e0422 >= 0; __obf_13b2631ac20e0422-- {
		if __obf_6a222ce425464583.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] != 0 {
			break
		}
		__obf_1b5d2df4fcb389df = __obf_13b2631ac20e0422 - 1
	}

	if __obf_1b5d2df4fcb389df < 0 {
		return __obf_7f0a87df1252e0f1{}
	} else if __obf_1b5d2df4fcb389df < __obf_90986b368b7bf87d-1 {
		__obf_6a222ce425464583.__obf_dc2d266987a80f84 = __obf_6a222ce425464583.__obf_dc2d266987a80f84[0 : __obf_1b5d2df4fcb389df+1]
	}

	return __obf_6a222ce425464583
}

func (__obf_6a222ce425464583 __obf_7f0a87df1252e0f1) string(__obf_839cfeba5a59664d bool) string {
	var __obf_76f79d224d59bf86 string
	__obf_90986b368b7bf87d := __obf_6a222ce425464583.__obf_90986b368b7bf87d()

	for __obf_13b2631ac20e0422 := __obf_90986b368b7bf87d - 1; __obf_13b2631ac20e0422 >= 0; __obf_13b2631ac20e0422-- {
		if __obf_6a222ce425464583.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] > 0 {
			if len(__obf_76f79d224d59bf86) > 0 {
				__obf_76f79d224d59bf86 += " + "
			}

			if !__obf_839cfeba5a59664d {
				__obf_76f79d224d59bf86 += fmt.Sprintf("%dx^%d", __obf_6a222ce425464583.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422], __obf_13b2631ac20e0422)
			} else {
				__obf_76f79d224d59bf86 += fmt.Sprintf("a^%dx^%d", __obf_c770ec6df6f3be50[__obf_6a222ce425464583.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422]], __obf_13b2631ac20e0422)
			}
		}
	}

	if len(__obf_76f79d224d59bf86) == 0 {
		__obf_76f79d224d59bf86 = "0"
	}

	return __obf_76f79d224d59bf86
}

// equals returns true if e == other.
func (__obf_6a222ce425464583 __obf_7f0a87df1252e0f1) __obf_94ba6c94c165b9b4(__obf_193727952bb7eba7 __obf_7f0a87df1252e0f1) bool {
	var __obf_0b31cc30d3ea9604 *__obf_7f0a87df1252e0f1
	var __obf_40a03d5e48a75a3c *__obf_7f0a87df1252e0f1

	if __obf_6a222ce425464583.__obf_90986b368b7bf87d() > __obf_193727952bb7eba7.__obf_90986b368b7bf87d() {
		__obf_0b31cc30d3ea9604 = &__obf_193727952bb7eba7
		__obf_40a03d5e48a75a3c = &__obf_6a222ce425464583
	} else {
		__obf_0b31cc30d3ea9604 = &__obf_6a222ce425464583
		__obf_40a03d5e48a75a3c = &__obf_193727952bb7eba7
	}
	__obf_a49c9ae5b2fc0d52 := __obf_0b31cc30d3ea9604.__obf_90986b368b7bf87d()
	__obf_d358d20570d89b31 := __obf_40a03d5e48a75a3c.__obf_90986b368b7bf87d()

	for __obf_13b2631ac20e0422 := 0; __obf_13b2631ac20e0422 < __obf_a49c9ae5b2fc0d52; __obf_13b2631ac20e0422++ {
		if __obf_6a222ce425464583.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] != __obf_193727952bb7eba7.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] {
			return false
		}
	}

	for __obf_13b2631ac20e0422 := __obf_a49c9ae5b2fc0d52; __obf_13b2631ac20e0422 < __obf_d358d20570d89b31; __obf_13b2631ac20e0422++ {
		if __obf_40a03d5e48a75a3c.__obf_dc2d266987a80f84[__obf_13b2631ac20e0422] != 0 {
			return false
		}
	}

	return true
}
