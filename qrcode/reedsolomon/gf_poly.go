// go-qrcode
// Copyright 2014 Tom Harwood

package __obf_0a6b19e606a79505

import (
	"fmt"
	"log"

	bitset "github.com/ArtisanHiram/go-pkg/qrcode/bitset"
)

// gfPoly is a polynomial over GF(2^8).
type __obf_c6021ed06d3a96ab struct {
	__obf_930e137791b175f5 []__obf_1005b52d0b467f2a// The ith value is the coefficient of the ith degree of x.
	// term[0]*(x^0) + term[1]*(x^1) + term[2]*(x^2) ...

}

// newGFPolyFromData returns |data| as a polynomial over GF(2^8).
//
// Each data byte becomes the coefficient of an x term.
//
// For an n byte input the polynomial is:
// data[n-1]*(x^n-1) + data[n-2]*(x^n-2) ... + data[0]*(x^0).
func __obf_c64ac0e17e6e6c02(__obf_28ed8d33917ceb06 *bitset.Bitset) __obf_c6021ed06d3a96ab {
	__obf_56d03b442eed1298 := __obf_28ed8d33917ceb06.Len() / 8
	if __obf_28ed8d33917ceb06.Len()%8 != 0 {
		__obf_56d03b442eed1298++
	}
	__obf_1d5fad7dc8e9e1a7 := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: make([]__obf_1005b52d0b467f2a, __obf_56d03b442eed1298)}
	__obf_18c84b23c99ce50d := __obf_56d03b442eed1298 - 1
	for __obf_05d969dc3f41e520 := 0; __obf_05d969dc3f41e520 < __obf_28ed8d33917ceb06.Len(); __obf_05d969dc3f41e520 += 8 {
		__obf_1d5fad7dc8e9e1a7.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] = __obf_1005b52d0b467f2a(__obf_28ed8d33917ceb06.ByteAt(__obf_05d969dc3f41e520))
		__obf_18c84b23c99ce50d--
	}

	return __obf_1d5fad7dc8e9e1a7
}

// newGFPolyMonomial returns term*(x^degree).
func __obf_e9fb1cfcff10030e(__obf_930e137791b175f5 __obf_1005b52d0b467f2a, __obf_60b1039746a15fe5 int) __obf_c6021ed06d3a96ab {
	if __obf_930e137791b175f5 == __obf_74388a1f073e5a33 {
		return __obf_c6021ed06d3a96ab{}
	}
	__obf_1d5fad7dc8e9e1a7 := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: make([]__obf_1005b52d0b467f2a, __obf_60b1039746a15fe5+1)}
	__obf_1d5fad7dc8e9e1a7.__obf_930e137791b175f5[__obf_60b1039746a15fe5] = __obf_930e137791b175f5

	return __obf_1d5fad7dc8e9e1a7
}

func (__obf_8ca4d5a5552f8746 __obf_c6021ed06d3a96ab) __obf_28ed8d33917ceb06(__obf_78c3d913df4a5580 int) []byte {
	__obf_1d5fad7dc8e9e1a7 := make([]byte, __obf_78c3d913df4a5580)
	__obf_18c84b23c99ce50d := __obf_78c3d913df4a5580 - len(__obf_8ca4d5a5552f8746.__obf_930e137791b175f5)
	for __obf_05d969dc3f41e520 := len(__obf_8ca4d5a5552f8746.__obf_930e137791b175f5) - 1; __obf_05d969dc3f41e520 >= 0; __obf_05d969dc3f41e520-- {
		__obf_1d5fad7dc8e9e1a7[__obf_18c84b23c99ce50d] = byte(__obf_8ca4d5a5552f8746.__obf_930e137791b175f5[__obf_05d969dc3f41e520])
		__obf_18c84b23c99ce50d++
	}

	return __obf_1d5fad7dc8e9e1a7
}

// numTerms returns the number of
func (__obf_8ca4d5a5552f8746 __obf_c6021ed06d3a96ab) __obf_78c3d913df4a5580() int {
	return len(__obf_8ca4d5a5552f8746.

		// gfPolyMultiply returns a * b.
		__obf_930e137791b175f5)
}

func __obf_19e74ed5348a919f(__obf_44c83566874b3d58, __obf_54089db7aa10cd4f __obf_c6021ed06d3a96ab) __obf_c6021ed06d3a96ab {
	__obf_d8e42017c1642603 := __obf_44c83566874b3d58.__obf_78c3d913df4a5580()
	__obf_b9d3f1f774abec52 := __obf_54089db7aa10cd4f.__obf_78c3d913df4a5580()
	__obf_1d5fad7dc8e9e1a7 := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: make([]__obf_1005b52d0b467f2a, __obf_d8e42017c1642603+__obf_b9d3f1f774abec52)}

	for __obf_18c84b23c99ce50d := 0; __obf_18c84b23c99ce50d < __obf_d8e42017c1642603; __obf_18c84b23c99ce50d++ {
		for __obf_05d969dc3f41e520 := 0; __obf_05d969dc3f41e520 < __obf_b9d3f1f774abec52; __obf_05d969dc3f41e520++ {
			if __obf_44c83566874b3d58.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] != 0 && __obf_54089db7aa10cd4f.__obf_930e137791b175f5[__obf_05d969dc3f41e520] != 0 {
				__obf_0eb5e9f957fdd41f := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: make([]__obf_1005b52d0b467f2a, __obf_18c84b23c99ce50d+__obf_05d969dc3f41e520+1)}
				__obf_0eb5e9f957fdd41f.__obf_930e137791b175f5[__obf_18c84b23c99ce50d+__obf_05d969dc3f41e520] = __obf_6e5e2fdc814867da(__obf_44c83566874b3d58.__obf_930e137791b175f5[__obf_18c84b23c99ce50d], __obf_54089db7aa10cd4f.__obf_930e137791b175f5[__obf_05d969dc3f41e520])
				__obf_1d5fad7dc8e9e1a7 = __obf_04a3bdcce93c2117(__obf_1d5fad7dc8e9e1a7, __obf_0eb5e9f957fdd41f)
			}
		}
	}

	return __obf_1d5fad7dc8e9e1a7.__obf_388d8f4ab6eefc6a()
}

// gfPolyRemainder return the remainder of numerator / denominator.
func __obf_8e9e781c42fba024(__obf_599e93e336246fb2, __obf_0c258b70906ce13b __obf_c6021ed06d3a96ab) __obf_c6021ed06d3a96ab {
	if __obf_0c258b70906ce13b.__obf_718521931934eaef(__obf_c6021ed06d3a96ab{}) {
		log.Panicln("Remainder by zero")
	}
	__obf_9d5e81d223c58052 := __obf_599e93e336246fb2

	for __obf_9d5e81d223c58052.__obf_78c3d913df4a5580() >= __obf_0c258b70906ce13b.__obf_78c3d913df4a5580() {
		__obf_60b1039746a15fe5 := __obf_9d5e81d223c58052.__obf_78c3d913df4a5580() - __obf_0c258b70906ce13b.__obf_78c3d913df4a5580()
		__obf_da684f9c3bc548b4 := __obf_72b9ba8e4962a920(__obf_9d5e81d223c58052.__obf_930e137791b175f5[__obf_9d5e81d223c58052.__obf_78c3d913df4a5580()-1], __obf_0c258b70906ce13b.__obf_930e137791b175f5[__obf_0c258b70906ce13b.__obf_78c3d913df4a5580()-1])
		__obf_eb61b7a36a8fb9da := __obf_19e74ed5348a919f(__obf_0c258b70906ce13b, __obf_e9fb1cfcff10030e(__obf_da684f9c3bc548b4, __obf_60b1039746a15fe5))
		__obf_9d5e81d223c58052 = __obf_04a3bdcce93c2117(__obf_9d5e81d223c58052, __obf_eb61b7a36a8fb9da)
	}

	return __obf_9d5e81d223c58052.__obf_388d8f4ab6eefc6a()
}

// gfPolyAdd returns a + b.
func __obf_04a3bdcce93c2117(__obf_44c83566874b3d58, __obf_54089db7aa10cd4f __obf_c6021ed06d3a96ab) __obf_c6021ed06d3a96ab {
	__obf_d8e42017c1642603 := __obf_44c83566874b3d58.__obf_78c3d913df4a5580()
	__obf_b9d3f1f774abec52 := __obf_54089db7aa10cd4f.__obf_78c3d913df4a5580()
	__obf_78c3d913df4a5580 := __obf_d8e42017c1642603
	if __obf_b9d3f1f774abec52 > __obf_78c3d913df4a5580 {
		__obf_78c3d913df4a5580 = __obf_b9d3f1f774abec52
	}
	__obf_1d5fad7dc8e9e1a7 := __obf_c6021ed06d3a96ab{__obf_930e137791b175f5: make([]__obf_1005b52d0b467f2a, __obf_78c3d913df4a5580)}

	for __obf_18c84b23c99ce50d := 0; __obf_18c84b23c99ce50d < __obf_78c3d913df4a5580; __obf_18c84b23c99ce50d++ {
		switch {
		case __obf_d8e42017c1642603 > __obf_18c84b23c99ce50d && __obf_b9d3f1f774abec52 > __obf_18c84b23c99ce50d:
			__obf_1d5fad7dc8e9e1a7.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] = __obf_c897ad4cbfe3790d(__obf_44c83566874b3d58.__obf_930e137791b175f5[__obf_18c84b23c99ce50d], __obf_54089db7aa10cd4f.__obf_930e137791b175f5[__obf_18c84b23c99ce50d])
		case __obf_d8e42017c1642603 > __obf_18c84b23c99ce50d:
			__obf_1d5fad7dc8e9e1a7.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] = __obf_44c83566874b3d58.__obf_930e137791b175f5[__obf_18c84b23c99ce50d]
		default:
			__obf_1d5fad7dc8e9e1a7.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] = __obf_54089db7aa10cd4f.__obf_930e137791b175f5[__obf_18c84b23c99ce50d]
		}
	}

	return __obf_1d5fad7dc8e9e1a7.__obf_388d8f4ab6eefc6a()
}

func (__obf_8ca4d5a5552f8746 __obf_c6021ed06d3a96ab) __obf_388d8f4ab6eefc6a() __obf_c6021ed06d3a96ab {
	__obf_78c3d913df4a5580 := __obf_8ca4d5a5552f8746.__obf_78c3d913df4a5580()
	__obf_a13ec5daa97e877c := __obf_78c3d913df4a5580 - 1

	for __obf_18c84b23c99ce50d := __obf_78c3d913df4a5580 - 1; __obf_18c84b23c99ce50d >= 0; __obf_18c84b23c99ce50d-- {
		if __obf_8ca4d5a5552f8746.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] != 0 {
			break
		}
		__obf_a13ec5daa97e877c = __obf_18c84b23c99ce50d - 1
	}

	if __obf_a13ec5daa97e877c < 0 {
		return __obf_c6021ed06d3a96ab{}
	} else if __obf_a13ec5daa97e877c < __obf_78c3d913df4a5580-1 {
		__obf_8ca4d5a5552f8746.__obf_930e137791b175f5 = __obf_8ca4d5a5552f8746.__obf_930e137791b175f5[0 : __obf_a13ec5daa97e877c+1]
	}

	return __obf_8ca4d5a5552f8746
}

func (__obf_8ca4d5a5552f8746 __obf_c6021ed06d3a96ab) string(__obf_9fe7ce81011f6f3e bool) string {
	var __obf_fe62cd7c0020f2c1 string
	__obf_78c3d913df4a5580 := __obf_8ca4d5a5552f8746.__obf_78c3d913df4a5580()

	for __obf_18c84b23c99ce50d := __obf_78c3d913df4a5580 - 1; __obf_18c84b23c99ce50d >= 0; __obf_18c84b23c99ce50d-- {
		if __obf_8ca4d5a5552f8746.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] > 0 {
			if len(__obf_fe62cd7c0020f2c1) > 0 {
				__obf_fe62cd7c0020f2c1 += " + "
			}

			if !__obf_9fe7ce81011f6f3e {
				__obf_fe62cd7c0020f2c1 += fmt.Sprintf("%dx^%d", __obf_8ca4d5a5552f8746.__obf_930e137791b175f5[__obf_18c84b23c99ce50d], __obf_18c84b23c99ce50d)
			} else {
				__obf_fe62cd7c0020f2c1 += fmt.Sprintf("a^%dx^%d", __obf_95c9dee3ea1fd467[__obf_8ca4d5a5552f8746.__obf_930e137791b175f5[__obf_18c84b23c99ce50d]], __obf_18c84b23c99ce50d)
			}
		}
	}

	if len(__obf_fe62cd7c0020f2c1) == 0 {
		__obf_fe62cd7c0020f2c1 = "0"
	}

	return __obf_fe62cd7c0020f2c1
}

// equals returns true if e == other.
func (__obf_8ca4d5a5552f8746 __obf_c6021ed06d3a96ab) __obf_718521931934eaef(__obf_33d2921cc5d7682e __obf_c6021ed06d3a96ab) bool {
	var __obf_462999eafd13530a *__obf_c6021ed06d3a96ab
	var __obf_552afc8c2c175141 *__obf_c6021ed06d3a96ab

	if __obf_8ca4d5a5552f8746.__obf_78c3d913df4a5580() > __obf_33d2921cc5d7682e.__obf_78c3d913df4a5580() {
		__obf_462999eafd13530a = &__obf_33d2921cc5d7682e
		__obf_552afc8c2c175141 = &__obf_8ca4d5a5552f8746
	} else {
		__obf_462999eafd13530a = &__obf_8ca4d5a5552f8746
		__obf_552afc8c2c175141 = &__obf_33d2921cc5d7682e
	}
	__obf_4579da18ba5eec9a := __obf_462999eafd13530a.__obf_78c3d913df4a5580()
	__obf_4eb15779eafed19b := __obf_552afc8c2c175141.__obf_78c3d913df4a5580()

	for __obf_18c84b23c99ce50d := 0; __obf_18c84b23c99ce50d < __obf_4579da18ba5eec9a; __obf_18c84b23c99ce50d++ {
		if __obf_8ca4d5a5552f8746.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] != __obf_33d2921cc5d7682e.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] {
			return false
		}
	}

	for __obf_18c84b23c99ce50d := __obf_4579da18ba5eec9a; __obf_18c84b23c99ce50d < __obf_4eb15779eafed19b; __obf_18c84b23c99ce50d++ {
		if __obf_552afc8c2c175141.__obf_930e137791b175f5[__obf_18c84b23c99ce50d] != 0 {
			return false
		}
	}

	return true
}
