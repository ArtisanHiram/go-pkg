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

package __obf_42b2ccbdafaee9c3

import (
	"math"
)

func __obf_79c8d4ae69d6ba30(__obf_c82b366489672bd7 float64) float64 {
	if __obf_c82b366489672bd7 >= -0.5 && __obf_c82b366489672bd7 < 0.5 {
		return 1
	}
	return 0
}

func __obf_51af381b36e05225(__obf_c82b366489672bd7 float64) float64 {
	__obf_c82b366489672bd7 = math.Abs(__obf_c82b366489672bd7)
	if __obf_c82b366489672bd7 <= 1 {
		return 1 - __obf_c82b366489672bd7
	}
	return 0
}

func __obf_cfab9d784a3a0826(__obf_c82b366489672bd7 float64) float64 {
	__obf_c82b366489672bd7 = math.Abs(__obf_c82b366489672bd7)
	if __obf_c82b366489672bd7 <= 1 {
		return __obf_c82b366489672bd7*__obf_c82b366489672bd7*(1.5*__obf_c82b366489672bd7-2.5) + 1.0
	}
	if __obf_c82b366489672bd7 <= 2 {
		return __obf_c82b366489672bd7*(__obf_c82b366489672bd7*(2.5-0.5*__obf_c82b366489672bd7)-4.0) + 2.0
	}
	return 0
}

func __obf_dfd3b61eaeee906a(__obf_c82b366489672bd7 float64) float64 {
	__obf_c82b366489672bd7 = math.Abs(__obf_c82b366489672bd7)
	if __obf_c82b366489672bd7 <= 1 {
		return (7.0*__obf_c82b366489672bd7*__obf_c82b366489672bd7*__obf_c82b366489672bd7 - 12.0*__obf_c82b366489672bd7*__obf_c82b366489672bd7 + 5.33333333333) * 0.16666666666
	}
	if __obf_c82b366489672bd7 <= 2 {
		return (-2.33333333333*__obf_c82b366489672bd7*__obf_c82b366489672bd7*__obf_c82b366489672bd7 + 12.0*__obf_c82b366489672bd7*__obf_c82b366489672bd7 - 20.0*__obf_c82b366489672bd7 + 10.6666666667) * 0.16666666666
	}
	return 0
}

func __obf_542a6f2bb2fd0155(__obf_9ab3da411fe8abc3 float64) float64 {
	__obf_9ab3da411fe8abc3 = math.Abs(__obf_9ab3da411fe8abc3) * math.Pi
	if __obf_9ab3da411fe8abc3 >= 1.220703e-4 {
		return math.Sin(__obf_9ab3da411fe8abc3) / __obf_9ab3da411fe8abc3
	}
	return 1
}

func __obf_33d6fa64374d1a50(__obf_c82b366489672bd7 float64) float64 {
	if __obf_c82b366489672bd7 > -2 && __obf_c82b366489672bd7 < 2 {
		return __obf_542a6f2bb2fd0155(__obf_c82b366489672bd7) * __obf_542a6f2bb2fd0155(__obf_c82b366489672bd7*0.5)
	}
	return 0
}

func __obf_7f1d3df8ac799ca4(__obf_c82b366489672bd7 float64) float64 {
	if __obf_c82b366489672bd7 > -3 && __obf_c82b366489672bd7 < 3 {
		return __obf_542a6f2bb2fd0155(__obf_c82b366489672bd7) * __obf_542a6f2bb2fd0155(__obf_c82b366489672bd7*0.3333333333333333)
	}
	return 0
}

// range [-256,256]
func __obf_ea0d9b36f90839c7(__obf_c214bb9a0f38874f, __obf_b34b7524c4bfcd41 int, __obf_eef0584a9ec97f3d, __obf_dc6826d61e7dccec float64, __obf_9e806529abfb439f func(float64) float64) ([]int16, []int, int) {
	__obf_b34b7524c4bfcd41 = __obf_b34b7524c4bfcd41 * int(math.Max(math.Ceil(__obf_eef0584a9ec97f3d*__obf_dc6826d61e7dccec), 1))
	__obf_0c6c9d4dec5d5e6a := math.Min(1./(__obf_eef0584a9ec97f3d*__obf_dc6826d61e7dccec), 1)
	__obf_71732cf03a2f4bc1 := make([]int16, __obf_c214bb9a0f38874f*__obf_b34b7524c4bfcd41)
	__obf_3563c6352e74b67e := make([]int, __obf_c214bb9a0f38874f)
	for __obf_b876c8539198a4e8 := range __obf_c214bb9a0f38874f {
		__obf_6521a460917ac8a4 := __obf_dc6826d61e7dccec*(float64(__obf_b876c8539198a4e8)+0.5) - 0.5
		__obf_3563c6352e74b67e[__obf_b876c8539198a4e8] = int(__obf_6521a460917ac8a4) - __obf_b34b7524c4bfcd41/2 + 1
		__obf_6521a460917ac8a4 -= float64(__obf_3563c6352e74b67e[__obf_b876c8539198a4e8])
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
			__obf_c82b366489672bd7 := (__obf_6521a460917ac8a4 - float64(__obf_e39e622bfcb683d9)) * __obf_0c6c9d4dec5d5e6a
			__obf_71732cf03a2f4bc1[__obf_b876c8539198a4e8*__obf_b34b7524c4bfcd41+__obf_e39e622bfcb683d9] = int16(__obf_9e806529abfb439f(__obf_c82b366489672bd7) * 256)
		}
	}

	return __obf_71732cf03a2f4bc1, __obf_3563c6352e74b67e,

		// range [-65536,65536]
		__obf_b34b7524c4bfcd41
}

func __obf_4f01f1c9705b203c(__obf_c214bb9a0f38874f, __obf_b34b7524c4bfcd41 int, __obf_eef0584a9ec97f3d, __obf_dc6826d61e7dccec float64, __obf_9e806529abfb439f func(float64) float64) ([]int32, []int, int) {
	__obf_b34b7524c4bfcd41 = __obf_b34b7524c4bfcd41 * int(math.Max(math.Ceil(__obf_eef0584a9ec97f3d*__obf_dc6826d61e7dccec), 1))
	__obf_0c6c9d4dec5d5e6a := math.Min(1./(__obf_eef0584a9ec97f3d*__obf_dc6826d61e7dccec), 1)
	__obf_71732cf03a2f4bc1 := make([]int32, __obf_c214bb9a0f38874f*__obf_b34b7524c4bfcd41)
	__obf_3563c6352e74b67e := make([]int, __obf_c214bb9a0f38874f)
	for __obf_b876c8539198a4e8 := range __obf_c214bb9a0f38874f {
		__obf_6521a460917ac8a4 := __obf_dc6826d61e7dccec*(float64(__obf_b876c8539198a4e8)+0.5) - 0.5
		__obf_3563c6352e74b67e[__obf_b876c8539198a4e8] = int(__obf_6521a460917ac8a4) - __obf_b34b7524c4bfcd41/2 + 1
		__obf_6521a460917ac8a4 -= float64(__obf_3563c6352e74b67e[__obf_b876c8539198a4e8])
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
			__obf_c82b366489672bd7 := (__obf_6521a460917ac8a4 - float64(__obf_e39e622bfcb683d9)) * __obf_0c6c9d4dec5d5e6a
			__obf_71732cf03a2f4bc1[__obf_b876c8539198a4e8*__obf_b34b7524c4bfcd41+__obf_e39e622bfcb683d9] = int32(__obf_9e806529abfb439f(__obf_c82b366489672bd7) * 65536)
		}
	}

	return __obf_71732cf03a2f4bc1, __obf_3563c6352e74b67e, __obf_b34b7524c4bfcd41
}

func __obf_ed15c920a71519da(__obf_c214bb9a0f38874f, __obf_b34b7524c4bfcd41 int, __obf_eef0584a9ec97f3d, __obf_dc6826d61e7dccec float64) ([]bool, []int, int) {
	__obf_b34b7524c4bfcd41 = __obf_b34b7524c4bfcd41 * int(math.Max(math.Ceil(__obf_eef0584a9ec97f3d*__obf_dc6826d61e7dccec), 1))
	__obf_0c6c9d4dec5d5e6a := math.Min(1./(__obf_eef0584a9ec97f3d*__obf_dc6826d61e7dccec), 1)
	__obf_71732cf03a2f4bc1 := make([]bool, __obf_c214bb9a0f38874f*__obf_b34b7524c4bfcd41)
	__obf_3563c6352e74b67e := make([]int, __obf_c214bb9a0f38874f)
	for __obf_b876c8539198a4e8 := range __obf_c214bb9a0f38874f {
		__obf_6521a460917ac8a4 := __obf_dc6826d61e7dccec*(float64(__obf_b876c8539198a4e8)+0.5) - 0.5
		__obf_3563c6352e74b67e[__obf_b876c8539198a4e8] = int(__obf_6521a460917ac8a4) - __obf_b34b7524c4bfcd41/2 + 1
		__obf_6521a460917ac8a4 -= float64(__obf_3563c6352e74b67e[__obf_b876c8539198a4e8])
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_b34b7524c4bfcd41; __obf_e39e622bfcb683d9++ {
			__obf_c82b366489672bd7 := (__obf_6521a460917ac8a4 - float64(__obf_e39e622bfcb683d9)) * __obf_0c6c9d4dec5d5e6a
			if __obf_c82b366489672bd7 >= -0.5 && __obf_c82b366489672bd7 < 0.5 {
				__obf_71732cf03a2f4bc1[__obf_b876c8539198a4e8*__obf_b34b7524c4bfcd41+__obf_e39e622bfcb683d9] = true
			} else {
				__obf_71732cf03a2f4bc1[__obf_b876c8539198a4e8*__obf_b34b7524c4bfcd41+__obf_e39e622bfcb683d9] = false
			}
		}
	}

	return __obf_71732cf03a2f4bc1, __obf_3563c6352e74b67e, __obf_b34b7524c4bfcd41
}
