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

package __obf_ac510735d4d13cdd

import (
	"math"
)

func __obf_6a3a721569bbf375(__obf_3ccd19e51488be10 float64) float64 {
	if __obf_3ccd19e51488be10 >= -0.5 && __obf_3ccd19e51488be10 < 0.5 {
		return 1
	}
	return 0
}

func __obf_75f6718d0c68ce45(__obf_3ccd19e51488be10 float64) float64 {
	__obf_3ccd19e51488be10 = math.Abs(__obf_3ccd19e51488be10)
	if __obf_3ccd19e51488be10 <= 1 {
		return 1 - __obf_3ccd19e51488be10
	}
	return 0
}

func __obf_6dbe12e595f20b30(__obf_3ccd19e51488be10 float64) float64 {
	__obf_3ccd19e51488be10 = math.Abs(__obf_3ccd19e51488be10)
	if __obf_3ccd19e51488be10 <= 1 {
		return __obf_3ccd19e51488be10*__obf_3ccd19e51488be10*(1.5*__obf_3ccd19e51488be10-2.5) + 1.0
	}
	if __obf_3ccd19e51488be10 <= 2 {
		return __obf_3ccd19e51488be10*(__obf_3ccd19e51488be10*(2.5-0.5*__obf_3ccd19e51488be10)-4.0) + 2.0
	}
	return 0
}

func __obf_4f5e93c03ec56195(__obf_3ccd19e51488be10 float64) float64 {
	__obf_3ccd19e51488be10 = math.Abs(__obf_3ccd19e51488be10)
	if __obf_3ccd19e51488be10 <= 1 {
		return (7.0*__obf_3ccd19e51488be10*__obf_3ccd19e51488be10*__obf_3ccd19e51488be10 - 12.0*__obf_3ccd19e51488be10*__obf_3ccd19e51488be10 + 5.33333333333) * 0.16666666666
	}
	if __obf_3ccd19e51488be10 <= 2 {
		return (-2.33333333333*__obf_3ccd19e51488be10*__obf_3ccd19e51488be10*__obf_3ccd19e51488be10 + 12.0*__obf_3ccd19e51488be10*__obf_3ccd19e51488be10 - 20.0*__obf_3ccd19e51488be10 + 10.6666666667) * 0.16666666666
	}
	return 0
}

func __obf_8ed673ce6beb86c3(__obf_21a3146223eb514f float64) float64 {
	__obf_21a3146223eb514f = math.Abs(__obf_21a3146223eb514f) * math.Pi
	if __obf_21a3146223eb514f >= 1.220703e-4 {
		return math.Sin(__obf_21a3146223eb514f) / __obf_21a3146223eb514f
	}
	return 1
}

func __obf_55ba6e33e6203680(__obf_3ccd19e51488be10 float64) float64 {
	if __obf_3ccd19e51488be10 > -2 && __obf_3ccd19e51488be10 < 2 {
		return __obf_8ed673ce6beb86c3(__obf_3ccd19e51488be10) * __obf_8ed673ce6beb86c3(__obf_3ccd19e51488be10*0.5)
	}
	return 0
}

func __obf_bc625d1a66a5919f(__obf_3ccd19e51488be10 float64) float64 {
	if __obf_3ccd19e51488be10 > -3 && __obf_3ccd19e51488be10 < 3 {
		return __obf_8ed673ce6beb86c3(__obf_3ccd19e51488be10) * __obf_8ed673ce6beb86c3(__obf_3ccd19e51488be10*0.3333333333333333)
	}
	return 0
}

// range [-256,256]
func __obf_ce2a9e53779a9624(__obf_a4bb31388097f2ff, __obf_052b80a7daaa449d int, __obf_b5313b280a820ad0, __obf_6f7dfba2f58bb527 float64, __obf_0a93ef33b92d5d34 func(float64) float64) ([]int16, []int, int) {
	__obf_052b80a7daaa449d = __obf_052b80a7daaa449d * int(math.Max(math.Ceil(__obf_b5313b280a820ad0*__obf_6f7dfba2f58bb527), 1))
	__obf_622fcccdc60bc763 := math.Min(1./(__obf_b5313b280a820ad0*__obf_6f7dfba2f58bb527), 1)
	__obf_46e0cb0caad4034a := make([]int16, __obf_a4bb31388097f2ff*__obf_052b80a7daaa449d)
	__obf_807c674c91e487a0 := make([]int, __obf_a4bb31388097f2ff)
	for __obf_4811266080d7795e := range __obf_a4bb31388097f2ff {
		__obf_7475603fed39c165 := __obf_6f7dfba2f58bb527*(float64(__obf_4811266080d7795e)+0.5) - 0.5
		__obf_807c674c91e487a0[__obf_4811266080d7795e] = int(__obf_7475603fed39c165) - __obf_052b80a7daaa449d/2 + 1
		__obf_7475603fed39c165 -= float64(__obf_807c674c91e487a0[__obf_4811266080d7795e])
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
			__obf_3ccd19e51488be10 := (__obf_7475603fed39c165 - float64(__obf_7326873273fab46e)) * __obf_622fcccdc60bc763
			__obf_46e0cb0caad4034a[__obf_4811266080d7795e*__obf_052b80a7daaa449d+__obf_7326873273fab46e] = int16(__obf_0a93ef33b92d5d34(__obf_3ccd19e51488be10) * 256)
		}
	}

	return __obf_46e0cb0caad4034a, __obf_807c674c91e487a0,

		// range [-65536,65536]
		__obf_052b80a7daaa449d
}

func __obf_9e73f2386e0e3872(__obf_a4bb31388097f2ff, __obf_052b80a7daaa449d int, __obf_b5313b280a820ad0, __obf_6f7dfba2f58bb527 float64, __obf_0a93ef33b92d5d34 func(float64) float64) ([]int32, []int, int) {
	__obf_052b80a7daaa449d = __obf_052b80a7daaa449d * int(math.Max(math.Ceil(__obf_b5313b280a820ad0*__obf_6f7dfba2f58bb527), 1))
	__obf_622fcccdc60bc763 := math.Min(1./(__obf_b5313b280a820ad0*__obf_6f7dfba2f58bb527), 1)
	__obf_46e0cb0caad4034a := make([]int32, __obf_a4bb31388097f2ff*__obf_052b80a7daaa449d)
	__obf_807c674c91e487a0 := make([]int, __obf_a4bb31388097f2ff)
	for __obf_4811266080d7795e := range __obf_a4bb31388097f2ff {
		__obf_7475603fed39c165 := __obf_6f7dfba2f58bb527*(float64(__obf_4811266080d7795e)+0.5) - 0.5
		__obf_807c674c91e487a0[__obf_4811266080d7795e] = int(__obf_7475603fed39c165) - __obf_052b80a7daaa449d/2 + 1
		__obf_7475603fed39c165 -= float64(__obf_807c674c91e487a0[__obf_4811266080d7795e])
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
			__obf_3ccd19e51488be10 := (__obf_7475603fed39c165 - float64(__obf_7326873273fab46e)) * __obf_622fcccdc60bc763
			__obf_46e0cb0caad4034a[__obf_4811266080d7795e*__obf_052b80a7daaa449d+__obf_7326873273fab46e] = int32(__obf_0a93ef33b92d5d34(__obf_3ccd19e51488be10) * 65536)
		}
	}

	return __obf_46e0cb0caad4034a, __obf_807c674c91e487a0, __obf_052b80a7daaa449d
}

func __obf_b8bb0d275d8b4e0d(__obf_a4bb31388097f2ff, __obf_052b80a7daaa449d int, __obf_b5313b280a820ad0, __obf_6f7dfba2f58bb527 float64) ([]bool, []int, int) {
	__obf_052b80a7daaa449d = __obf_052b80a7daaa449d * int(math.Max(math.Ceil(__obf_b5313b280a820ad0*__obf_6f7dfba2f58bb527), 1))
	__obf_622fcccdc60bc763 := math.Min(1./(__obf_b5313b280a820ad0*__obf_6f7dfba2f58bb527), 1)
	__obf_46e0cb0caad4034a := make([]bool, __obf_a4bb31388097f2ff*__obf_052b80a7daaa449d)
	__obf_807c674c91e487a0 := make([]int, __obf_a4bb31388097f2ff)
	for __obf_4811266080d7795e := range __obf_a4bb31388097f2ff {
		__obf_7475603fed39c165 := __obf_6f7dfba2f58bb527*(float64(__obf_4811266080d7795e)+0.5) - 0.5
		__obf_807c674c91e487a0[__obf_4811266080d7795e] = int(__obf_7475603fed39c165) - __obf_052b80a7daaa449d/2 + 1
		__obf_7475603fed39c165 -= float64(__obf_807c674c91e487a0[__obf_4811266080d7795e])
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_052b80a7daaa449d; __obf_7326873273fab46e++ {
			__obf_3ccd19e51488be10 := (__obf_7475603fed39c165 - float64(__obf_7326873273fab46e)) * __obf_622fcccdc60bc763
			if __obf_3ccd19e51488be10 >= -0.5 && __obf_3ccd19e51488be10 < 0.5 {
				__obf_46e0cb0caad4034a[__obf_4811266080d7795e*__obf_052b80a7daaa449d+__obf_7326873273fab46e] = true
			} else {
				__obf_46e0cb0caad4034a[__obf_4811266080d7795e*__obf_052b80a7daaa449d+__obf_7326873273fab46e] = false
			}
		}
	}

	return __obf_46e0cb0caad4034a, __obf_807c674c91e487a0, __obf_052b80a7daaa449d
}
