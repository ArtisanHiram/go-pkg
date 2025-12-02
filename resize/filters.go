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

package __obf_4f71249d3f545aa0

import (
	"math"
)

func __obf_3d46464e0d553bae(__obf_a5eab9a5623e4978 float64) float64 {
	if __obf_a5eab9a5623e4978 >= -0.5 && __obf_a5eab9a5623e4978 < 0.5 {
		return 1
	}
	return 0
}

func __obf_62a9d12c047474d5(__obf_a5eab9a5623e4978 float64) float64 {
	__obf_a5eab9a5623e4978 = math.Abs(__obf_a5eab9a5623e4978)
	if __obf_a5eab9a5623e4978 <= 1 {
		return 1 - __obf_a5eab9a5623e4978
	}
	return 0
}

func __obf_0a92c4954c46d5b5(__obf_a5eab9a5623e4978 float64) float64 {
	__obf_a5eab9a5623e4978 = math.Abs(__obf_a5eab9a5623e4978)
	if __obf_a5eab9a5623e4978 <= 1 {
		return __obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978*(1.5*__obf_a5eab9a5623e4978-2.5) + 1.0
	}
	if __obf_a5eab9a5623e4978 <= 2 {
		return __obf_a5eab9a5623e4978*(__obf_a5eab9a5623e4978*(2.5-0.5*__obf_a5eab9a5623e4978)-4.0) + 2.0
	}
	return 0
}

func __obf_34caadcbfbe56d7c(__obf_a5eab9a5623e4978 float64) float64 {
	__obf_a5eab9a5623e4978 = math.Abs(__obf_a5eab9a5623e4978)
	if __obf_a5eab9a5623e4978 <= 1 {
		return (7.0*__obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978 - 12.0*__obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978 + 5.33333333333) * 0.16666666666
	}
	if __obf_a5eab9a5623e4978 <= 2 {
		return (-2.33333333333*__obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978 + 12.0*__obf_a5eab9a5623e4978*__obf_a5eab9a5623e4978 - 20.0*__obf_a5eab9a5623e4978 + 10.6666666667) * 0.16666666666
	}
	return 0
}

func __obf_d8291ab49cf7349e(__obf_2a2c911b5786aa81 float64) float64 {
	__obf_2a2c911b5786aa81 = math.Abs(__obf_2a2c911b5786aa81) * math.Pi
	if __obf_2a2c911b5786aa81 >= 1.220703e-4 {
		return math.Sin(__obf_2a2c911b5786aa81) / __obf_2a2c911b5786aa81
	}
	return 1
}

func __obf_57e614f28bbd7ca3(__obf_a5eab9a5623e4978 float64) float64 {
	if __obf_a5eab9a5623e4978 > -2 && __obf_a5eab9a5623e4978 < 2 {
		return __obf_d8291ab49cf7349e(__obf_a5eab9a5623e4978) * __obf_d8291ab49cf7349e(__obf_a5eab9a5623e4978*0.5)
	}
	return 0
}

func __obf_ed3f12ba8d7c6de5(__obf_a5eab9a5623e4978 float64) float64 {
	if __obf_a5eab9a5623e4978 > -3 && __obf_a5eab9a5623e4978 < 3 {
		return __obf_d8291ab49cf7349e(__obf_a5eab9a5623e4978) * __obf_d8291ab49cf7349e(__obf_a5eab9a5623e4978*0.3333333333333333)
	}
	return 0
}

// range [-256,256]
func __obf_bc71b3618bd8ea83(__obf_18cb70a86b1e19d7, __obf_743fac2b8ed0fc32 int, __obf_ce79a531ca9b78ac, __obf_eff7440668961f2c float64, __obf_25630250e4587e64 func(float64) float64) ([]int16, []int, int) {
	__obf_743fac2b8ed0fc32 = __obf_743fac2b8ed0fc32 * int(math.Max(math.Ceil(__obf_ce79a531ca9b78ac*__obf_eff7440668961f2c), 1))
	__obf_f7238f4f448814ea := math.Min(1./(__obf_ce79a531ca9b78ac*__obf_eff7440668961f2c), 1)
	__obf_ed4a0c7a1913ca57 := make([]int16, __obf_18cb70a86b1e19d7*__obf_743fac2b8ed0fc32)
	__obf_de3222b659094173 := make([]int, __obf_18cb70a86b1e19d7)
	for __obf_36c97b6a0fef1dad := range __obf_18cb70a86b1e19d7 {
		__obf_756148fef84f73ba := __obf_eff7440668961f2c*(float64(__obf_36c97b6a0fef1dad)+0.5) - 0.5
		__obf_de3222b659094173[__obf_36c97b6a0fef1dad] = int(__obf_756148fef84f73ba) - __obf_743fac2b8ed0fc32/2 + 1
		__obf_756148fef84f73ba -= float64(__obf_de3222b659094173[__obf_36c97b6a0fef1dad])
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
			__obf_a5eab9a5623e4978 := (__obf_756148fef84f73ba - float64(__obf_b92089dde6a4a00b)) * __obf_f7238f4f448814ea
			__obf_ed4a0c7a1913ca57[__obf_36c97b6a0fef1dad*__obf_743fac2b8ed0fc32+__obf_b92089dde6a4a00b] = int16(__obf_25630250e4587e64(__obf_a5eab9a5623e4978) * 256)
		}
	}

	return __obf_ed4a0c7a1913ca57, __obf_de3222b659094173,

		// range [-65536,65536]
		__obf_743fac2b8ed0fc32
}

func __obf_7e80e77b51d72f60(__obf_18cb70a86b1e19d7, __obf_743fac2b8ed0fc32 int, __obf_ce79a531ca9b78ac, __obf_eff7440668961f2c float64, __obf_25630250e4587e64 func(float64) float64) ([]int32, []int, int) {
	__obf_743fac2b8ed0fc32 = __obf_743fac2b8ed0fc32 * int(math.Max(math.Ceil(__obf_ce79a531ca9b78ac*__obf_eff7440668961f2c), 1))
	__obf_f7238f4f448814ea := math.Min(1./(__obf_ce79a531ca9b78ac*__obf_eff7440668961f2c), 1)
	__obf_ed4a0c7a1913ca57 := make([]int32, __obf_18cb70a86b1e19d7*__obf_743fac2b8ed0fc32)
	__obf_de3222b659094173 := make([]int, __obf_18cb70a86b1e19d7)
	for __obf_36c97b6a0fef1dad := range __obf_18cb70a86b1e19d7 {
		__obf_756148fef84f73ba := __obf_eff7440668961f2c*(float64(__obf_36c97b6a0fef1dad)+0.5) - 0.5
		__obf_de3222b659094173[__obf_36c97b6a0fef1dad] = int(__obf_756148fef84f73ba) - __obf_743fac2b8ed0fc32/2 + 1
		__obf_756148fef84f73ba -= float64(__obf_de3222b659094173[__obf_36c97b6a0fef1dad])
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
			__obf_a5eab9a5623e4978 := (__obf_756148fef84f73ba - float64(__obf_b92089dde6a4a00b)) * __obf_f7238f4f448814ea
			__obf_ed4a0c7a1913ca57[__obf_36c97b6a0fef1dad*__obf_743fac2b8ed0fc32+__obf_b92089dde6a4a00b] = int32(__obf_25630250e4587e64(__obf_a5eab9a5623e4978) * 65536)
		}
	}

	return __obf_ed4a0c7a1913ca57, __obf_de3222b659094173, __obf_743fac2b8ed0fc32
}

func __obf_b0d4c44c11725e77(__obf_18cb70a86b1e19d7, __obf_743fac2b8ed0fc32 int, __obf_ce79a531ca9b78ac, __obf_eff7440668961f2c float64) ([]bool, []int, int) {
	__obf_743fac2b8ed0fc32 = __obf_743fac2b8ed0fc32 * int(math.Max(math.Ceil(__obf_ce79a531ca9b78ac*__obf_eff7440668961f2c), 1))
	__obf_f7238f4f448814ea := math.Min(1./(__obf_ce79a531ca9b78ac*__obf_eff7440668961f2c), 1)
	__obf_ed4a0c7a1913ca57 := make([]bool, __obf_18cb70a86b1e19d7*__obf_743fac2b8ed0fc32)
	__obf_de3222b659094173 := make([]int, __obf_18cb70a86b1e19d7)
	for __obf_36c97b6a0fef1dad := range __obf_18cb70a86b1e19d7 {
		__obf_756148fef84f73ba := __obf_eff7440668961f2c*(float64(__obf_36c97b6a0fef1dad)+0.5) - 0.5
		__obf_de3222b659094173[__obf_36c97b6a0fef1dad] = int(__obf_756148fef84f73ba) - __obf_743fac2b8ed0fc32/2 + 1
		__obf_756148fef84f73ba -= float64(__obf_de3222b659094173[__obf_36c97b6a0fef1dad])
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_743fac2b8ed0fc32; __obf_b92089dde6a4a00b++ {
			__obf_a5eab9a5623e4978 := (__obf_756148fef84f73ba - float64(__obf_b92089dde6a4a00b)) * __obf_f7238f4f448814ea
			if __obf_a5eab9a5623e4978 >= -0.5 && __obf_a5eab9a5623e4978 < 0.5 {
				__obf_ed4a0c7a1913ca57[__obf_36c97b6a0fef1dad*__obf_743fac2b8ed0fc32+__obf_b92089dde6a4a00b] = true
			} else {
				__obf_ed4a0c7a1913ca57[__obf_36c97b6a0fef1dad*__obf_743fac2b8ed0fc32+__obf_b92089dde6a4a00b] = false
			}
		}
	}

	return __obf_ed4a0c7a1913ca57, __obf_de3222b659094173, __obf_743fac2b8ed0fc32
}
