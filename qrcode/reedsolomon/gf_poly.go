// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_d5be46fdf9a047aa

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_681065f079dec5b9 struct {
	__obf_f7ff1cffece1f4f8 []__obf_94bafdc0d5ce6bf6// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...

}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_b72e69aa5faca050(__obf_6a71e09d65c25d89 *bitset.Bitset) __obf_681065f079dec5b9 {
	__obf_841d39c7fa11d7ff := __obf_6a71e09d65c25d89.Len() / 8
	if __obf_6a71e09d65c25d89.Len()%8 != 0 {
		__obf_841d39c7fa11d7ff++
	}
	__obf_2efda911f3d2d3c1 := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: make([]__obf_94bafdc0d5ce6bf6, __obf_841d39c7fa11d7ff)}
	__obf_183af01b71412ca5 := __obf_841d39c7fa11d7ff - 1
	for __obf_62f9d824e6456ab1 := 0; __obf_62f9d824e6456ab1 < __obf_6a71e09d65c25d89.Len(); __obf_62f9d824e6456ab1 += 8 {
		__obf_2efda911f3d2d3c1.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] = __obf_94bafdc0d5ce6bf6(__obf_6a71e09d65c25d89.ByteAt(__obf_62f9d824e6456ab1))
		__obf_183af01b71412ca5--
	}

	return __obf_2efda911f3d2d3c1
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_0498c51b59c427a4(__obf_f7ff1cffece1f4f8 __obf_94bafdc0d5ce6bf6, __obf_bce1f4d9d31e3eda int) __obf_681065f079dec5b9 {
	if __obf_f7ff1cffece1f4f8 == __obf_b79d07b388290551 {
		return __obf_681065f079dec5b9{}
	}
	__obf_2efda911f3d2d3c1 := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: make([]__obf_94bafdc0d5ce6bf6, __obf_bce1f4d9d31e3eda+1)}
	__obf_2efda911f3d2d3c1.__obf_f7ff1cffece1f4f8[__obf_bce1f4d9d31e3eda] = __obf_f7ff1cffece1f4f8

	return __obf_2efda911f3d2d3c1
}

func (__obf_ef7ed160d78f24f5 __obf_681065f079dec5b9) __obf_6a71e09d65c25d89(__obf_3ed35d872bf70b8b int) []byte {
	__obf_2efda911f3d2d3c1 := make([]byte, __obf_3ed35d872bf70b8b)
	__obf_183af01b71412ca5 := __obf_3ed35d872bf70b8b - len(__obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8)
	for __obf_62f9d824e6456ab1 := len(__obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8) - 1; __obf_62f9d824e6456ab1 >= 0; __obf_62f9d824e6456ab1-- {
		__obf_2efda911f3d2d3c1[__obf_183af01b71412ca5] = byte(__obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[__obf_62f9d824e6456ab1])
		__obf_183af01b71412ca5++
	}

	return __obf_2efda911f3d2d3c1
}

// numTerms returns the number of
func (__obf_ef7ed160d78f24f5 __obf_681065f079dec5b9) __obf_3ed35d872bf70b8b() int {
	return len(__obf_ef7ed160d78f24f5.

		// gfPolyMultiply returns a * b.
		__obf_f7ff1cffece1f4f8)
}

func __obf_83e2645b1ae7eb30(__obf_98a34ba73ba6a166, __obf_0e533b075e8992c6 __obf_681065f079dec5b9) __obf_681065f079dec5b9 {
	__obf_93c7ffcabd7e446e := __obf_98a34ba73ba6a166.__obf_3ed35d872bf70b8b()
	__obf_90a1d81be045d14c := __obf_0e533b075e8992c6.__obf_3ed35d872bf70b8b()
	__obf_2efda911f3d2d3c1 := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: make([]__obf_94bafdc0d5ce6bf6, __obf_93c7ffcabd7e446e+__obf_90a1d81be045d14c)}

	for __obf_183af01b71412ca5 := 0; __obf_183af01b71412ca5 < __obf_93c7ffcabd7e446e; __obf_183af01b71412ca5++ {
		for __obf_62f9d824e6456ab1 := 0; __obf_62f9d824e6456ab1 < __obf_90a1d81be045d14c; __obf_62f9d824e6456ab1++ {
			if __obf_98a34ba73ba6a166.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] != 0 && __obf_0e533b075e8992c6.__obf_f7ff1cffece1f4f8[__obf_62f9d824e6456ab1] != 0 {
				__obf_6f1e002b99f49337 := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: make([]__obf_94bafdc0d5ce6bf6, __obf_183af01b71412ca5+__obf_62f9d824e6456ab1+1)}
				__obf_6f1e002b99f49337.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5+__obf_62f9d824e6456ab1] = __obf_9969f7a117b39136(__obf_98a34ba73ba6a166.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5], __obf_0e533b075e8992c6.__obf_f7ff1cffece1f4f8[__obf_62f9d824e6456ab1])
				__obf_2efda911f3d2d3c1 = __obf_55165c4592812ee1(__obf_2efda911f3d2d3c1, __obf_6f1e002b99f49337)
			}
		}
	}

	return __obf_2efda911f3d2d3c1.__obf_492624977e9d05aa()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_c170f4171b722bb2(__obf_bfb6a81589546aff, __obf_82bd180cbbe64356 __obf_681065f079dec5b9) __obf_681065f079dec5b9 {
	if __obf_82bd180cbbe64356.__obf_b731279436d7a849(__obf_681065f079dec5b9{}) {
		log.Panicln("Remainder by zero")
	}
	__obf_d87cecb5fbe51cf0 := __obf_bfb6a81589546aff

	for __obf_d87cecb5fbe51cf0.__obf_3ed35d872bf70b8b() >= __obf_82bd180cbbe64356.__obf_3ed35d872bf70b8b() {
		__obf_bce1f4d9d31e3eda := __obf_d87cecb5fbe51cf0.__obf_3ed35d872bf70b8b() - __obf_82bd180cbbe64356.__obf_3ed35d872bf70b8b()
		__obf_44c3451e29b312cc := __obf_04fa20e14d374dcc(__obf_d87cecb5fbe51cf0.__obf_f7ff1cffece1f4f8[__obf_d87cecb5fbe51cf0.__obf_3ed35d872bf70b8b()-1], __obf_82bd180cbbe64356.__obf_f7ff1cffece1f4f8[__obf_82bd180cbbe64356.__obf_3ed35d872bf70b8b()-1])
		__obf_66291eb52d672276 := __obf_83e2645b1ae7eb30(__obf_82bd180cbbe64356, __obf_0498c51b59c427a4(__obf_44c3451e29b312cc, __obf_bce1f4d9d31e3eda))
		__obf_d87cecb5fbe51cf0 = __obf_55165c4592812ee1(__obf_d87cecb5fbe51cf0, __obf_66291eb52d672276)
	}

	return __obf_d87cecb5fbe51cf0.__obf_492624977e9d05aa()
}

// gfPolyAdd returns a + b.
func __obf_55165c4592812ee1(__obf_98a34ba73ba6a166, __obf_0e533b075e8992c6 __obf_681065f079dec5b9) __obf_681065f079dec5b9 {
	__obf_93c7ffcabd7e446e := __obf_98a34ba73ba6a166.__obf_3ed35d872bf70b8b()
	__obf_90a1d81be045d14c := __obf_0e533b075e8992c6.__obf_3ed35d872bf70b8b()
	__obf_3ed35d872bf70b8b := __obf_93c7ffcabd7e446e
	if __obf_90a1d81be045d14c > __obf_3ed35d872bf70b8b {
		__obf_3ed35d872bf70b8b = __obf_90a1d81be045d14c
	}
	__obf_2efda911f3d2d3c1 := __obf_681065f079dec5b9{__obf_f7ff1cffece1f4f8: make([]__obf_94bafdc0d5ce6bf6, __obf_3ed35d872bf70b8b)}

	for __obf_183af01b71412ca5 := 0; __obf_183af01b71412ca5 < __obf_3ed35d872bf70b8b; __obf_183af01b71412ca5++ {
		switch {
		case __obf_93c7ffcabd7e446e > __obf_183af01b71412ca5 && __obf_90a1d81be045d14c > __obf_183af01b71412ca5:
			__obf_2efda911f3d2d3c1.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] = __obf_c0193b490a988aa4(__obf_98a34ba73ba6a166.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5], __obf_0e533b075e8992c6.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5])
		case __obf_93c7ffcabd7e446e > __obf_183af01b71412ca5:
			__obf_2efda911f3d2d3c1.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] = __obf_98a34ba73ba6a166.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5]
		default:
			__obf_2efda911f3d2d3c1.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] = __obf_0e533b075e8992c6.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5]
		}
	}

	return __obf_2efda911f3d2d3c1.__obf_492624977e9d05aa()
}

func (__obf_ef7ed160d78f24f5 __obf_681065f079dec5b9) __obf_492624977e9d05aa() __obf_681065f079dec5b9 {
	__obf_3ed35d872bf70b8b := __obf_ef7ed160d78f24f5.__obf_3ed35d872bf70b8b()
	__obf_25a577ef7de103ef := __obf_3ed35d872bf70b8b - 1

	for __obf_183af01b71412ca5 := __obf_3ed35d872bf70b8b - 1; __obf_183af01b71412ca5 >= 0; __obf_183af01b71412ca5-- {
		if __obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] != 0 {
			break
		}
		__obf_25a577ef7de103ef = __obf_183af01b71412ca5 - 1
	}

	if __obf_25a577ef7de103ef < 0 {
		return __obf_681065f079dec5b9{}
	} else if __obf_25a577ef7de103ef < __obf_3ed35d872bf70b8b-1 {
		__obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8 = __obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[0 : __obf_25a577ef7de103ef+1]
	}

	return __obf_ef7ed160d78f24f5
}

func (__obf_ef7ed160d78f24f5 __obf_681065f079dec5b9) string(__obf_78e26e618504cba8 bool) string {
	var __obf_159527e1fe1264c3 string
	__obf_3ed35d872bf70b8b := __obf_ef7ed160d78f24f5.__obf_3ed35d872bf70b8b()

	for __obf_183af01b71412ca5 := __obf_3ed35d872bf70b8b - 1; __obf_183af01b71412ca5 >= 0; __obf_183af01b71412ca5-- {
		if __obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] > 0 {
			if len(__obf_159527e1fe1264c3) > 0 {
				__obf_159527e1fe1264c3 += " + "
			}

			if !__obf_78e26e618504cba8 {
				__obf_159527e1fe1264c3 += fmt.Sprintf("%dx^%d", __obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5], __obf_183af01b71412ca5)
			} else {
				__obf_159527e1fe1264c3 += fmt.Sprintf("a^%dx^%d", __obf_c02e00c89ccc9280[__obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5]], __obf_183af01b71412ca5)
			}
		}
	}

	if len(__obf_159527e1fe1264c3) == 0 {
		__obf_159527e1fe1264c3 = "0"
	}

	return __obf_159527e1fe1264c3
}

// equals returns true if e == other.
func (__obf_ef7ed160d78f24f5 __obf_681065f079dec5b9) __obf_b731279436d7a849(__obf_e13a87e6f406dafa __obf_681065f079dec5b9) bool {
	var __obf_bd5c48d133cfb0a1 *__obf_681065f079dec5b9
	var __obf_a0a61b17074ea4f2 *__obf_681065f079dec5b9

	if __obf_ef7ed160d78f24f5.__obf_3ed35d872bf70b8b() > __obf_e13a87e6f406dafa.__obf_3ed35d872bf70b8b() {
		__obf_bd5c48d133cfb0a1 = &__obf_e13a87e6f406dafa
		__obf_a0a61b17074ea4f2 = &__obf_ef7ed160d78f24f5
	} else {
		__obf_bd5c48d133cfb0a1 = &__obf_ef7ed160d78f24f5
		__obf_a0a61b17074ea4f2 = &__obf_e13a87e6f406dafa
	}
	__obf_4c9d2a735970abb5 := __obf_bd5c48d133cfb0a1.__obf_3ed35d872bf70b8b()
	__obf_3d43dde623d1178e := __obf_a0a61b17074ea4f2.__obf_3ed35d872bf70b8b()

	for __obf_183af01b71412ca5 := 0; __obf_183af01b71412ca5 < __obf_4c9d2a735970abb5; __obf_183af01b71412ca5++ {
		if __obf_ef7ed160d78f24f5.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] != __obf_e13a87e6f406dafa.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] {
			return false
		}
	}

	for __obf_183af01b71412ca5 := __obf_4c9d2a735970abb5; __obf_183af01b71412ca5 < __obf_3d43dde623d1178e; __obf_183af01b71412ca5++ {
		if __obf_a0a61b17074ea4f2.__obf_f7ff1cffece1f4f8[__obf_183af01b71412ca5] != 0 {
			return false
		}
	}

	return true
}
