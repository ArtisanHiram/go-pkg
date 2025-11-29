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

package __obf_3858dbfa2ccfdbe9

import (
	"math"
)

func __obf_ea38e93055f919d7(__obf_8da81bf3c5df7f0d float64) float64 {
	if __obf_8da81bf3c5df7f0d >= -0.5 && __obf_8da81bf3c5df7f0d < 0.5 {
		return 1
	}
	return 0
}

func __obf_29ef71619471455c(__obf_8da81bf3c5df7f0d float64) float64 {
	__obf_8da81bf3c5df7f0d = math.Abs(__obf_8da81bf3c5df7f0d)
	if __obf_8da81bf3c5df7f0d <= 1 {
		return 1 - __obf_8da81bf3c5df7f0d
	}
	return 0
}

func __obf_9b5fddd514be8a29(__obf_8da81bf3c5df7f0d float64) float64 {
	__obf_8da81bf3c5df7f0d = math.Abs(__obf_8da81bf3c5df7f0d)
	if __obf_8da81bf3c5df7f0d <= 1 {
		return __obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d*(1.5*__obf_8da81bf3c5df7f0d-2.5) + 1.0
	}
	if __obf_8da81bf3c5df7f0d <= 2 {
		return __obf_8da81bf3c5df7f0d*(__obf_8da81bf3c5df7f0d*(2.5-0.5*__obf_8da81bf3c5df7f0d)-4.0) + 2.0
	}
	return 0
}

func __obf_0b97bf31dd3b88bb(__obf_8da81bf3c5df7f0d float64) float64 {
	__obf_8da81bf3c5df7f0d = math.Abs(__obf_8da81bf3c5df7f0d)
	if __obf_8da81bf3c5df7f0d <= 1 {
		return (7.0*__obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d - 12.0*__obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d + 5.33333333333) * 0.16666666666
	}
	if __obf_8da81bf3c5df7f0d <= 2 {
		return (-2.33333333333*__obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d + 12.0*__obf_8da81bf3c5df7f0d*__obf_8da81bf3c5df7f0d - 20.0*__obf_8da81bf3c5df7f0d + 10.6666666667) * 0.16666666666
	}
	return 0
}

func __obf_9cb4f15c8c762b0b(__obf_15a5faef6b57be64 float64) float64 {
	__obf_15a5faef6b57be64 = math.Abs(__obf_15a5faef6b57be64) * math.Pi
	if __obf_15a5faef6b57be64 >= 1.220703e-4 {
		return math.Sin(__obf_15a5faef6b57be64) / __obf_15a5faef6b57be64
	}
	return 1
}

func __obf_18ba59002a602c1c(__obf_8da81bf3c5df7f0d float64) float64 {
	if __obf_8da81bf3c5df7f0d > -2 && __obf_8da81bf3c5df7f0d < 2 {
		return __obf_9cb4f15c8c762b0b(__obf_8da81bf3c5df7f0d) * __obf_9cb4f15c8c762b0b(__obf_8da81bf3c5df7f0d*0.5)
	}
	return 0
}

func __obf_d3ba12b86735c961(__obf_8da81bf3c5df7f0d float64) float64 {
	if __obf_8da81bf3c5df7f0d > -3 && __obf_8da81bf3c5df7f0d < 3 {
		return __obf_9cb4f15c8c762b0b(__obf_8da81bf3c5df7f0d) * __obf_9cb4f15c8c762b0b(__obf_8da81bf3c5df7f0d*0.3333333333333333)
	}
	return 0
}

// range [-256,256]
func __obf_3a1983a2984d289a(__obf_b3f8b84ce96f95b6, __obf_52dfb4be4832ebbd int, __obf_74fbaffb42caf3d6, __obf_0d0fd6056291c76b float64, __obf_e0303ccb3cff6d43 func(float64) float64) ([]int16, []int, int) {
	__obf_52dfb4be4832ebbd = __obf_52dfb4be4832ebbd * int(math.Max(math.Ceil(__obf_74fbaffb42caf3d6*__obf_0d0fd6056291c76b), 1))
	__obf_55e77894ce721218 := math.Min(1./(__obf_74fbaffb42caf3d6*__obf_0d0fd6056291c76b), 1)
	__obf_3fabfc59cdeeb8b4 := make([]int16, __obf_b3f8b84ce96f95b6*__obf_52dfb4be4832ebbd)
	__obf_5f2d417327929024 := make([]int, __obf_b3f8b84ce96f95b6)
	for __obf_523cb3ec667c6a54 := range __obf_b3f8b84ce96f95b6 {
		__obf_299b71e49ebe66d8 := __obf_0d0fd6056291c76b*(float64(__obf_523cb3ec667c6a54)+0.5) - 0.5
		__obf_5f2d417327929024[__obf_523cb3ec667c6a54] = int(__obf_299b71e49ebe66d8) - __obf_52dfb4be4832ebbd/2 + 1
		__obf_299b71e49ebe66d8 -= float64(__obf_5f2d417327929024[__obf_523cb3ec667c6a54])
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
			__obf_8da81bf3c5df7f0d := (__obf_299b71e49ebe66d8 - float64(__obf_9d9681e93d5a567f)) * __obf_55e77894ce721218
			__obf_3fabfc59cdeeb8b4[__obf_523cb3ec667c6a54*__obf_52dfb4be4832ebbd+__obf_9d9681e93d5a567f] = int16(__obf_e0303ccb3cff6d43(__obf_8da81bf3c5df7f0d) * 256)
		}
	}

	return __obf_3fabfc59cdeeb8b4, __obf_5f2d417327929024,

		// range [-65536,65536]
		__obf_52dfb4be4832ebbd
}

func __obf_4a688a78d5582d42(__obf_b3f8b84ce96f95b6, __obf_52dfb4be4832ebbd int, __obf_74fbaffb42caf3d6, __obf_0d0fd6056291c76b float64, __obf_e0303ccb3cff6d43 func(float64) float64) ([]int32, []int, int) {
	__obf_52dfb4be4832ebbd = __obf_52dfb4be4832ebbd * int(math.Max(math.Ceil(__obf_74fbaffb42caf3d6*__obf_0d0fd6056291c76b), 1))
	__obf_55e77894ce721218 := math.Min(1./(__obf_74fbaffb42caf3d6*__obf_0d0fd6056291c76b), 1)
	__obf_3fabfc59cdeeb8b4 := make([]int32, __obf_b3f8b84ce96f95b6*__obf_52dfb4be4832ebbd)
	__obf_5f2d417327929024 := make([]int, __obf_b3f8b84ce96f95b6)
	for __obf_523cb3ec667c6a54 := range __obf_b3f8b84ce96f95b6 {
		__obf_299b71e49ebe66d8 := __obf_0d0fd6056291c76b*(float64(__obf_523cb3ec667c6a54)+0.5) - 0.5
		__obf_5f2d417327929024[__obf_523cb3ec667c6a54] = int(__obf_299b71e49ebe66d8) - __obf_52dfb4be4832ebbd/2 + 1
		__obf_299b71e49ebe66d8 -= float64(__obf_5f2d417327929024[__obf_523cb3ec667c6a54])
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
			__obf_8da81bf3c5df7f0d := (__obf_299b71e49ebe66d8 - float64(__obf_9d9681e93d5a567f)) * __obf_55e77894ce721218
			__obf_3fabfc59cdeeb8b4[__obf_523cb3ec667c6a54*__obf_52dfb4be4832ebbd+__obf_9d9681e93d5a567f] = int32(__obf_e0303ccb3cff6d43(__obf_8da81bf3c5df7f0d) * 65536)
		}
	}

	return __obf_3fabfc59cdeeb8b4, __obf_5f2d417327929024, __obf_52dfb4be4832ebbd
}

func __obf_c218d3902b533b64(__obf_b3f8b84ce96f95b6, __obf_52dfb4be4832ebbd int, __obf_74fbaffb42caf3d6, __obf_0d0fd6056291c76b float64) ([]bool, []int, int) {
	__obf_52dfb4be4832ebbd = __obf_52dfb4be4832ebbd * int(math.Max(math.Ceil(__obf_74fbaffb42caf3d6*__obf_0d0fd6056291c76b), 1))
	__obf_55e77894ce721218 := math.Min(1./(__obf_74fbaffb42caf3d6*__obf_0d0fd6056291c76b), 1)
	__obf_3fabfc59cdeeb8b4 := make([]bool, __obf_b3f8b84ce96f95b6*__obf_52dfb4be4832ebbd)
	__obf_5f2d417327929024 := make([]int, __obf_b3f8b84ce96f95b6)
	for __obf_523cb3ec667c6a54 := range __obf_b3f8b84ce96f95b6 {
		__obf_299b71e49ebe66d8 := __obf_0d0fd6056291c76b*(float64(__obf_523cb3ec667c6a54)+0.5) - 0.5
		__obf_5f2d417327929024[__obf_523cb3ec667c6a54] = int(__obf_299b71e49ebe66d8) - __obf_52dfb4be4832ebbd/2 + 1
		__obf_299b71e49ebe66d8 -= float64(__obf_5f2d417327929024[__obf_523cb3ec667c6a54])
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_52dfb4be4832ebbd; __obf_9d9681e93d5a567f++ {
			__obf_8da81bf3c5df7f0d := (__obf_299b71e49ebe66d8 - float64(__obf_9d9681e93d5a567f)) * __obf_55e77894ce721218
			if __obf_8da81bf3c5df7f0d >= -0.5 && __obf_8da81bf3c5df7f0d < 0.5 {
				__obf_3fabfc59cdeeb8b4[__obf_523cb3ec667c6a54*__obf_52dfb4be4832ebbd+__obf_9d9681e93d5a567f] = true
			} else {
				__obf_3fabfc59cdeeb8b4[__obf_523cb3ec667c6a54*__obf_52dfb4be4832ebbd+__obf_9d9681e93d5a567f] = false
			}
		}
	}

	return __obf_3fabfc59cdeeb8b4, __obf_5f2d417327929024, __obf_52dfb4be4832ebbd
}
