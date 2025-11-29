/*
Copyright (c) 2012, Jan Schlicht <jan.schlicht@gmail.com>

Permission to use, copy, modify, and/or distribute this software for any purpose
with or without fee is hereby granted, provided that the above copyright notice
and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
THIS SOFTWARE.
*/

package __obf_71f4605f6400e2ce

import (
	"math"
)

func __obf_d8b0075bb092e33d(__obf_b13d16274eb954e2 float64) float64 {
	if __obf_b13d16274eb954e2 >= -0.5 && __obf_b13d16274eb954e2 < 0.5 {
		return 1
	}
	return 0
}

func __obf_995190fb5ae77946(__obf_b13d16274eb954e2 float64) float64 {
	__obf_b13d16274eb954e2 = math.Abs(__obf_b13d16274eb954e2)
	if __obf_b13d16274eb954e2 <= 1 {
		return 1 - __obf_b13d16274eb954e2
	}
	return 0
}

func __obf_c0b2553e410838b3(__obf_b13d16274eb954e2 float64) float64 {
	__obf_b13d16274eb954e2 = math.Abs(__obf_b13d16274eb954e2)
	if __obf_b13d16274eb954e2 <= 1 {
		return __obf_b13d16274eb954e2*__obf_b13d16274eb954e2*(1.5*__obf_b13d16274eb954e2-2.5) + 1.0
	}
	if __obf_b13d16274eb954e2 <= 2 {
		return __obf_b13d16274eb954e2*(__obf_b13d16274eb954e2*(2.5-0.5*__obf_b13d16274eb954e2)-4.0) + 2.0
	}
	return 0
}

func __obf_52c4a4763619b896(__obf_b13d16274eb954e2 float64) float64 {
	__obf_b13d16274eb954e2 = math.Abs(__obf_b13d16274eb954e2)
	if __obf_b13d16274eb954e2 <= 1 {
		return (7.0*__obf_b13d16274eb954e2*__obf_b13d16274eb954e2*__obf_b13d16274eb954e2 - 12.0*__obf_b13d16274eb954e2*__obf_b13d16274eb954e2 + 5.33333333333) * 0.16666666666
	}
	if __obf_b13d16274eb954e2 <= 2 {
		return (-2.33333333333*__obf_b13d16274eb954e2*__obf_b13d16274eb954e2*__obf_b13d16274eb954e2 + 12.0*__obf_b13d16274eb954e2*__obf_b13d16274eb954e2 - 20.0*__obf_b13d16274eb954e2 + 10.6666666667) * 0.16666666666
	}
	return 0
}

func __obf_284d7a302d591cdd(__obf_c6d53ec5b4505f6a float64) float64 {
	__obf_c6d53ec5b4505f6a = math.Abs(__obf_c6d53ec5b4505f6a) * math.Pi
	if __obf_c6d53ec5b4505f6a >= 1.220703e-4 {
		return math.Sin(__obf_c6d53ec5b4505f6a) / __obf_c6d53ec5b4505f6a
	}
	return 1
}

func __obf_f2fccd76139d1121(__obf_b13d16274eb954e2 float64) float64 {
	if __obf_b13d16274eb954e2 > -2 && __obf_b13d16274eb954e2 < 2 {
		return __obf_284d7a302d591cdd(__obf_b13d16274eb954e2) * __obf_284d7a302d591cdd(__obf_b13d16274eb954e2*0.5)
	}
	return 0
}

func __obf_e214463bbc8b0b56(__obf_b13d16274eb954e2 float64) float64 {
	if __obf_b13d16274eb954e2 > -3 && __obf_b13d16274eb954e2 < 3 {
		return __obf_284d7a302d591cdd(__obf_b13d16274eb954e2) * __obf_284d7a302d591cdd(__obf_b13d16274eb954e2*0.3333333333333333)
	}
	return 0
}

// range [-256,256]
func __obf_fc245964fca32498(__obf_34a5a372eef91e93, __obf_6a27c875a42baeab int, __obf_940ae86fccad7783, __obf_121d528711c040a7 float64, __obf_2fde8fe3fe7a0697 func(float64) float64) ([]int16, []int, int) {
	__obf_6a27c875a42baeab = __obf_6a27c875a42baeab * int(math.Max(math.Ceil(__obf_940ae86fccad7783*__obf_121d528711c040a7), 1))
	__obf_8baacb2094c1deed := math.Min(1./(__obf_940ae86fccad7783*__obf_121d528711c040a7), 1)
	__obf_73fc3de4b87d9301 := make([]int16, __obf_34a5a372eef91e93*__obf_6a27c875a42baeab)
	__obf_45570ac86f60c691 := make([]int, __obf_34a5a372eef91e93)
	for __obf_eed2203158cb878d := range __obf_34a5a372eef91e93 {
		__obf_a5fbd52b49136b3d := __obf_121d528711c040a7*(float64(__obf_eed2203158cb878d)+0.5) - 0.5
		__obf_45570ac86f60c691[__obf_eed2203158cb878d] = int(__obf_a5fbd52b49136b3d) - __obf_6a27c875a42baeab/2 + 1
		__obf_a5fbd52b49136b3d -= float64(__obf_45570ac86f60c691[__obf_eed2203158cb878d])
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
			__obf_b13d16274eb954e2 := (__obf_a5fbd52b49136b3d - float64(__obf_7d43b40e55127c49)) * __obf_8baacb2094c1deed
			__obf_73fc3de4b87d9301[__obf_eed2203158cb878d*__obf_6a27c875a42baeab+__obf_7d43b40e55127c49] = int16(__obf_2fde8fe3fe7a0697(__obf_b13d16274eb954e2) * 256)
		}
	}

	return __obf_73fc3de4b87d9301, __obf_45570ac86f60c691,

		// range [-65536,65536]
		__obf_6a27c875a42baeab
}

func __obf_0bc5ed636cd197d7(__obf_34a5a372eef91e93, __obf_6a27c875a42baeab int, __obf_940ae86fccad7783, __obf_121d528711c040a7 float64, __obf_2fde8fe3fe7a0697 func(float64) float64) ([]int32, []int, int) {
	__obf_6a27c875a42baeab = __obf_6a27c875a42baeab * int(math.Max(math.Ceil(__obf_940ae86fccad7783*__obf_121d528711c040a7), 1))
	__obf_8baacb2094c1deed := math.Min(1./(__obf_940ae86fccad7783*__obf_121d528711c040a7), 1)
	__obf_73fc3de4b87d9301 := make([]int32, __obf_34a5a372eef91e93*__obf_6a27c875a42baeab)
	__obf_45570ac86f60c691 := make([]int, __obf_34a5a372eef91e93)
	for __obf_eed2203158cb878d := range __obf_34a5a372eef91e93 {
		__obf_a5fbd52b49136b3d := __obf_121d528711c040a7*(float64(__obf_eed2203158cb878d)+0.5) - 0.5
		__obf_45570ac86f60c691[__obf_eed2203158cb878d] = int(__obf_a5fbd52b49136b3d) - __obf_6a27c875a42baeab/2 + 1
		__obf_a5fbd52b49136b3d -= float64(__obf_45570ac86f60c691[__obf_eed2203158cb878d])
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
			__obf_b13d16274eb954e2 := (__obf_a5fbd52b49136b3d - float64(__obf_7d43b40e55127c49)) * __obf_8baacb2094c1deed
			__obf_73fc3de4b87d9301[__obf_eed2203158cb878d*__obf_6a27c875a42baeab+__obf_7d43b40e55127c49] = int32(__obf_2fde8fe3fe7a0697(__obf_b13d16274eb954e2) * 65536)
		}
	}

	return __obf_73fc3de4b87d9301, __obf_45570ac86f60c691, __obf_6a27c875a42baeab
}

func __obf_0b05d823d55af64a(__obf_34a5a372eef91e93, __obf_6a27c875a42baeab int, __obf_940ae86fccad7783, __obf_121d528711c040a7 float64) ([]bool, []int, int) {
	__obf_6a27c875a42baeab = __obf_6a27c875a42baeab * int(math.Max(math.Ceil(__obf_940ae86fccad7783*__obf_121d528711c040a7), 1))
	__obf_8baacb2094c1deed := math.Min(1./(__obf_940ae86fccad7783*__obf_121d528711c040a7), 1)
	__obf_73fc3de4b87d9301 := make([]bool, __obf_34a5a372eef91e93*__obf_6a27c875a42baeab)
	__obf_45570ac86f60c691 := make([]int, __obf_34a5a372eef91e93)
	for __obf_eed2203158cb878d := range __obf_34a5a372eef91e93 {
		__obf_a5fbd52b49136b3d := __obf_121d528711c040a7*(float64(__obf_eed2203158cb878d)+0.5) - 0.5
		__obf_45570ac86f60c691[__obf_eed2203158cb878d] = int(__obf_a5fbd52b49136b3d) - __obf_6a27c875a42baeab/2 + 1
		__obf_a5fbd52b49136b3d -= float64(__obf_45570ac86f60c691[__obf_eed2203158cb878d])
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
			__obf_b13d16274eb954e2 := (__obf_a5fbd52b49136b3d - float64(__obf_7d43b40e55127c49)) * __obf_8baacb2094c1deed
			if __obf_b13d16274eb954e2 >= -0.5 && __obf_b13d16274eb954e2 < 0.5 {
				__obf_73fc3de4b87d9301[__obf_eed2203158cb878d*__obf_6a27c875a42baeab+__obf_7d43b40e55127c49] = true
			} else {
				__obf_73fc3de4b87d9301[__obf_eed2203158cb878d*__obf_6a27c875a42baeab+__obf_7d43b40e55127c49] = false
			}
		}
	}

	return __obf_73fc3de4b87d9301, __obf_45570ac86f60c691, __obf_6a27c875a42baeab
}
