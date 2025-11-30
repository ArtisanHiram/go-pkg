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

package __obf_d24101647651f4c3

import (
	"math"
)

func __obf_e958fd3b8bea8d9c(__obf_2a764779fc1401ce float64) float64 {
	if __obf_2a764779fc1401ce >= -0.5 && __obf_2a764779fc1401ce < 0.5 {
		return 1
	}
	return 0
}

func __obf_d66ee4de441cbd66(__obf_2a764779fc1401ce float64) float64 {
	__obf_2a764779fc1401ce = math.Abs(__obf_2a764779fc1401ce)
	if __obf_2a764779fc1401ce <= 1 {
		return 1 - __obf_2a764779fc1401ce
	}
	return 0
}

func __obf_a619d5b55f4394a4(__obf_2a764779fc1401ce float64) float64 {
	__obf_2a764779fc1401ce = math.Abs(__obf_2a764779fc1401ce)
	if __obf_2a764779fc1401ce <= 1 {
		return __obf_2a764779fc1401ce*__obf_2a764779fc1401ce*(1.5*__obf_2a764779fc1401ce-2.5) + 1.0
	}
	if __obf_2a764779fc1401ce <= 2 {
		return __obf_2a764779fc1401ce*(__obf_2a764779fc1401ce*(2.5-0.5*__obf_2a764779fc1401ce)-4.0) + 2.0
	}
	return 0
}

func __obf_7a44659276a985b3(__obf_2a764779fc1401ce float64) float64 {
	__obf_2a764779fc1401ce = math.Abs(__obf_2a764779fc1401ce)
	if __obf_2a764779fc1401ce <= 1 {
		return (7.0*__obf_2a764779fc1401ce*__obf_2a764779fc1401ce*__obf_2a764779fc1401ce - 12.0*__obf_2a764779fc1401ce*__obf_2a764779fc1401ce + 5.33333333333) * 0.16666666666
	}
	if __obf_2a764779fc1401ce <= 2 {
		return (-2.33333333333*__obf_2a764779fc1401ce*__obf_2a764779fc1401ce*__obf_2a764779fc1401ce + 12.0*__obf_2a764779fc1401ce*__obf_2a764779fc1401ce - 20.0*__obf_2a764779fc1401ce + 10.6666666667) * 0.16666666666
	}
	return 0
}

func __obf_b7312e848825c3d7(__obf_88d1a0c3dc734ee1 float64) float64 {
	__obf_88d1a0c3dc734ee1 = math.Abs(__obf_88d1a0c3dc734ee1) * math.Pi
	if __obf_88d1a0c3dc734ee1 >= 1.220703e-4 {
		return math.Sin(__obf_88d1a0c3dc734ee1) / __obf_88d1a0c3dc734ee1
	}
	return 1
}

func __obf_b3d8a34685baa99a(__obf_2a764779fc1401ce float64) float64 {
	if __obf_2a764779fc1401ce > -2 && __obf_2a764779fc1401ce < 2 {
		return __obf_b7312e848825c3d7(__obf_2a764779fc1401ce) * __obf_b7312e848825c3d7(__obf_2a764779fc1401ce*0.5)
	}
	return 0
}

func __obf_d668a1e333340f8d(__obf_2a764779fc1401ce float64) float64 {
	if __obf_2a764779fc1401ce > -3 && __obf_2a764779fc1401ce < 3 {
		return __obf_b7312e848825c3d7(__obf_2a764779fc1401ce) * __obf_b7312e848825c3d7(__obf_2a764779fc1401ce*0.3333333333333333)
	}
	return 0
}

// range [-256,256]
func __obf_b32cca196c808fe4(__obf_bbeb670f48a0fd60, __obf_277f8fd7946d0e2e int, __obf_b8f04e12843e5823, __obf_b65438e8a147153b float64, __obf_0a5c3b331d82d518 func(float64) float64) ([]int16, []int, int) {
	__obf_277f8fd7946d0e2e = __obf_277f8fd7946d0e2e * int(math.Max(math.Ceil(__obf_b8f04e12843e5823*__obf_b65438e8a147153b), 1))
	__obf_ff5daf2039973512 := math.Min(1./(__obf_b8f04e12843e5823*__obf_b65438e8a147153b), 1)
	__obf_779a7f81421e09ef := make([]int16, __obf_bbeb670f48a0fd60*__obf_277f8fd7946d0e2e)
	__obf_1eb9b282d993da7c := make([]int, __obf_bbeb670f48a0fd60)
	for __obf_097022f910960674 := range __obf_bbeb670f48a0fd60 {
		__obf_6dcd40695d737f95 := __obf_b65438e8a147153b*(float64(__obf_097022f910960674)+0.5) - 0.5
		__obf_1eb9b282d993da7c[__obf_097022f910960674] = int(__obf_6dcd40695d737f95) - __obf_277f8fd7946d0e2e/2 + 1
		__obf_6dcd40695d737f95 -= float64(__obf_1eb9b282d993da7c[__obf_097022f910960674])
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
			__obf_2a764779fc1401ce := (__obf_6dcd40695d737f95 - float64(__obf_41043e9a68df34f8)) * __obf_ff5daf2039973512
			__obf_779a7f81421e09ef[__obf_097022f910960674*__obf_277f8fd7946d0e2e+__obf_41043e9a68df34f8] = int16(__obf_0a5c3b331d82d518(__obf_2a764779fc1401ce) * 256)
		}
	}

	return __obf_779a7f81421e09ef, __obf_1eb9b282d993da7c,

		// range [-65536,65536]
		__obf_277f8fd7946d0e2e
}

func __obf_2a2dd68742b912d4(__obf_bbeb670f48a0fd60, __obf_277f8fd7946d0e2e int, __obf_b8f04e12843e5823, __obf_b65438e8a147153b float64, __obf_0a5c3b331d82d518 func(float64) float64) ([]int32, []int, int) {
	__obf_277f8fd7946d0e2e = __obf_277f8fd7946d0e2e * int(math.Max(math.Ceil(__obf_b8f04e12843e5823*__obf_b65438e8a147153b), 1))
	__obf_ff5daf2039973512 := math.Min(1./(__obf_b8f04e12843e5823*__obf_b65438e8a147153b), 1)
	__obf_779a7f81421e09ef := make([]int32, __obf_bbeb670f48a0fd60*__obf_277f8fd7946d0e2e)
	__obf_1eb9b282d993da7c := make([]int, __obf_bbeb670f48a0fd60)
	for __obf_097022f910960674 := range __obf_bbeb670f48a0fd60 {
		__obf_6dcd40695d737f95 := __obf_b65438e8a147153b*(float64(__obf_097022f910960674)+0.5) - 0.5
		__obf_1eb9b282d993da7c[__obf_097022f910960674] = int(__obf_6dcd40695d737f95) - __obf_277f8fd7946d0e2e/2 + 1
		__obf_6dcd40695d737f95 -= float64(__obf_1eb9b282d993da7c[__obf_097022f910960674])
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
			__obf_2a764779fc1401ce := (__obf_6dcd40695d737f95 - float64(__obf_41043e9a68df34f8)) * __obf_ff5daf2039973512
			__obf_779a7f81421e09ef[__obf_097022f910960674*__obf_277f8fd7946d0e2e+__obf_41043e9a68df34f8] = int32(__obf_0a5c3b331d82d518(__obf_2a764779fc1401ce) * 65536)
		}
	}

	return __obf_779a7f81421e09ef, __obf_1eb9b282d993da7c, __obf_277f8fd7946d0e2e
}

func __obf_073d620566900e2c(__obf_bbeb670f48a0fd60, __obf_277f8fd7946d0e2e int, __obf_b8f04e12843e5823, __obf_b65438e8a147153b float64) ([]bool, []int, int) {
	__obf_277f8fd7946d0e2e = __obf_277f8fd7946d0e2e * int(math.Max(math.Ceil(__obf_b8f04e12843e5823*__obf_b65438e8a147153b), 1))
	__obf_ff5daf2039973512 := math.Min(1./(__obf_b8f04e12843e5823*__obf_b65438e8a147153b), 1)
	__obf_779a7f81421e09ef := make([]bool, __obf_bbeb670f48a0fd60*__obf_277f8fd7946d0e2e)
	__obf_1eb9b282d993da7c := make([]int, __obf_bbeb670f48a0fd60)
	for __obf_097022f910960674 := range __obf_bbeb670f48a0fd60 {
		__obf_6dcd40695d737f95 := __obf_b65438e8a147153b*(float64(__obf_097022f910960674)+0.5) - 0.5
		__obf_1eb9b282d993da7c[__obf_097022f910960674] = int(__obf_6dcd40695d737f95) - __obf_277f8fd7946d0e2e/2 + 1
		__obf_6dcd40695d737f95 -= float64(__obf_1eb9b282d993da7c[__obf_097022f910960674])
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_277f8fd7946d0e2e; __obf_41043e9a68df34f8++ {
			__obf_2a764779fc1401ce := (__obf_6dcd40695d737f95 - float64(__obf_41043e9a68df34f8)) * __obf_ff5daf2039973512
			if __obf_2a764779fc1401ce >= -0.5 && __obf_2a764779fc1401ce < 0.5 {
				__obf_779a7f81421e09ef[__obf_097022f910960674*__obf_277f8fd7946d0e2e+__obf_41043e9a68df34f8] = true
			} else {
				__obf_779a7f81421e09ef[__obf_097022f910960674*__obf_277f8fd7946d0e2e+__obf_41043e9a68df34f8] = false
			}
		}
	}

	return __obf_779a7f81421e09ef, __obf_1eb9b282d993da7c, __obf_277f8fd7946d0e2e
}
