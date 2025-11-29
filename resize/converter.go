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

import "image"

// Keep value in [0,255] range.
func __obf_b892a4c4c7d416e4(__obf_b13d16274eb954e2 int32) uint8 {
	// casting a negative int to an uint will result in an overflown
	// large uint. this behavior will be exploited here and in other functions
	// to achieve a higher performance.
	if uint32(__obf_b13d16274eb954e2) < 256 {
		return uint8(__obf_b13d16274eb954e2)
	}
	if __obf_b13d16274eb954e2 > 255 {
		return 255
	}
	return 0
}

// Keep value in [0,65535] range.
func __obf_ebd52bcd902e20ba(__obf_b13d16274eb954e2 int64) uint16 {
	if uint64(__obf_b13d16274eb954e2) < 65536 {
		return uint16(__obf_b13d16274eb954e2)
	}
	if __obf_b13d16274eb954e2 > 65535 {
		return 65535
	}
	return 0
}

func __obf_ca5fdfe67ed70e0d(__obf_b13d16274eb954e2 image.Image, __obf_6a55e60196cdd102 *image.RGBA64, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int32, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]int64
			var __obf_63bac68a07d23c2c int64
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case __obf_27fc29a27a3bf82f < 0:
						__obf_27fc29a27a3bf82f = 0
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = __obf_7d59a7861d2c3cb4
					}
					__obf_1f6def59ec86a578, __obf_ef636de778967ee2, __obf_4efacd728de97e8f, __obf_a4cd46772704603d := __obf_b13d16274eb954e2.At(__obf_27fc29a27a3bf82f+__obf_b13d16274eb954e2.Bounds().Min.X, __obf_c6d53ec5b4505f6a+__obf_b13d16274eb954e2.Bounds().Min.Y).RGBA()
					__obf_47e29cd4ec4bd91e[0] += int64(__obf_9166f977ce6b9e75) * int64(__obf_1f6def59ec86a578)
					__obf_47e29cd4ec4bd91e[1] += int64(__obf_9166f977ce6b9e75) * int64(__obf_ef636de778967ee2)
					__obf_47e29cd4ec4bd91e[2] += int64(__obf_9166f977ce6b9e75) * int64(__obf_4efacd728de97e8f)
					__obf_47e29cd4ec4bd91e[3] += int64(__obf_9166f977ce6b9e75) * int64(__obf_a4cd46772704603d)
					__obf_63bac68a07d23c2c += int64(__obf_9166f977ce6b9e75)
				}
			}
			__obf_87dd30f656544b12 := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*8
			__obf_1cf6d31eaacee463 := __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+1] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+2] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+3] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+4] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+5] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+6] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+7] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_63620e18c66f5397(__obf_b13d16274eb954e2 *image.RGBA, __obf_6a55e60196cdd102 *image.RGBA, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int16, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]int32
			var __obf_63bac68a07d23c2c int32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 4
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 4 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_47e29cd4ec4bd91e[0] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])
					__obf_47e29cd4ec4bd91e[1] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1])
					__obf_47e29cd4ec4bd91e[2] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])
					__obf_47e29cd4ec4bd91e[3] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3])
					__obf_63bac68a07d23c2c += int32(__obf_9166f977ce6b9e75)
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*4
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_b6a8709004d14349(__obf_b13d16274eb954e2 *image.NRGBA, __obf_6a55e60196cdd102 *image.RGBA, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int16, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]int32
			var __obf_63bac68a07d23c2c int32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 4
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 4 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_a4cd46772704603d :=// Forward alpha-premultiplication
					int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3])
					__obf_1f6def59ec86a578 := int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0]) * __obf_a4cd46772704603d
					__obf_1f6def59ec86a578 /= 0xff
					__obf_ef636de778967ee2 := int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1]) * __obf_a4cd46772704603d
					__obf_ef636de778967ee2 /= 0xff
					__obf_4efacd728de97e8f := int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2]) * __obf_a4cd46772704603d
					__obf_4efacd728de97e8f /= 0xff
					__obf_47e29cd4ec4bd91e[0] += int32(__obf_9166f977ce6b9e75) * __obf_1f6def59ec86a578
					__obf_47e29cd4ec4bd91e[1] += int32(__obf_9166f977ce6b9e75) * __obf_ef636de778967ee2
					__obf_47e29cd4ec4bd91e[2] += int32(__obf_9166f977ce6b9e75) * __obf_4efacd728de97e8f
					__obf_47e29cd4ec4bd91e[3] += int32(__obf_9166f977ce6b9e75) * __obf_a4cd46772704603d
					__obf_63bac68a07d23c2c += int32(__obf_9166f977ce6b9e75)
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*4
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = __obf_b892a4c4c7d416e4(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_b799c14ff2666c0a(__obf_b13d16274eb954e2 *image.RGBA64, __obf_6a55e60196cdd102 *image.RGBA64, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int32, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]int64
			var __obf_63bac68a07d23c2c int64
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 8
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 8 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_47e29cd4ec4bd91e[0] += int64(__obf_9166f977ce6b9e75) * (int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])<<8 | int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1]))
					__obf_47e29cd4ec4bd91e[1] += int64(__obf_9166f977ce6b9e75) * (int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])<<8 | int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3]))
					__obf_47e29cd4ec4bd91e[2] += int64(__obf_9166f977ce6b9e75) * (int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+4])<<8 | int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+5]))
					__obf_47e29cd4ec4bd91e[3] += int64(__obf_9166f977ce6b9e75) * (int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+6])<<8 | int64(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+7]))
					__obf_63bac68a07d23c2c += int64(__obf_9166f977ce6b9e75)
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*8
			__obf_1cf6d31eaacee463 := __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+4] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+5] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+6] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+7] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_1ab7584b5da9bac1(__obf_b13d16274eb954e2 *image.NRGBA64, __obf_6a55e60196cdd102 *image.RGBA64, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int32, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]int64
			var __obf_63bac68a07d23c2c int64
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 8
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 8 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_a4cd46772704603d :=// Forward alpha-premultiplication
					int64(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+6])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+7]))
					__obf_1f6def59ec86a578 := int64(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])<<8|uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1])) * __obf_a4cd46772704603d
					__obf_1f6def59ec86a578 /= 0xffff
					__obf_ef636de778967ee2 := int64(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])<<8|uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3])) * __obf_a4cd46772704603d
					__obf_ef636de778967ee2 /= 0xffff
					__obf_4efacd728de97e8f := int64(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+4])<<8|uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+5])) * __obf_a4cd46772704603d
					__obf_4efacd728de97e8f /= 0xffff
					__obf_47e29cd4ec4bd91e[0] += int64(__obf_9166f977ce6b9e75) * __obf_1f6def59ec86a578
					__obf_47e29cd4ec4bd91e[1] += int64(__obf_9166f977ce6b9e75) * __obf_ef636de778967ee2
					__obf_47e29cd4ec4bd91e[2] += int64(__obf_9166f977ce6b9e75) * __obf_4efacd728de97e8f
					__obf_47e29cd4ec4bd91e[3] += int64(__obf_9166f977ce6b9e75) * __obf_a4cd46772704603d
					__obf_63bac68a07d23c2c += int64(__obf_9166f977ce6b9e75)
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*8
			__obf_1cf6d31eaacee463 := __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+4] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+5] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_ebd52bcd902e20ba(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+6] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+7] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_5727e96fbc4064ea(__obf_b13d16274eb954e2 *image.Gray, __obf_6a55e60196cdd102 *image.Gray, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int16, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[(__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_3fe9f7c50ed9bc3c int32
			var __obf_63bac68a07d23c2c int32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case __obf_27fc29a27a3bf82f < 0:
						__obf_27fc29a27a3bf82f = 0
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = __obf_7d59a7861d2c3cb4
					}
					__obf_3fe9f7c50ed9bc3c += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f])
					__obf_63bac68a07d23c2c += int32(__obf_9166f977ce6b9e75)
				}
			}
			__obf_87dd30f656544b12 := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a - __obf_85e316700ddb907a.Min.X)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12] = __obf_b892a4c4c7d416e4(__obf_3fe9f7c50ed9bc3c / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_723b7c23ca44d438(__obf_b13d16274eb954e2 *image.Gray16, __obf_6a55e60196cdd102 *image.Gray16, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int32, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_3fe9f7c50ed9bc3c int64
			var __obf_63bac68a07d23c2c int64
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 2
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 2 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_3fe9f7c50ed9bc3c += int64(__obf_9166f977ce6b9e75) * int64(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])<<8|uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1]))
					__obf_63bac68a07d23c2c += int64(__obf_9166f977ce6b9e75)
				}
			}
			__obf_87dd30f656544b12 := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*2
			__obf_1cf6d31eaacee463 := __obf_ebd52bcd902e20ba(__obf_3fe9f7c50ed9bc3c / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+1] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_22b27a29fbaf62b4(__obf_b13d16274eb954e2 *__obf_786a8b420048cd65, __obf_6a55e60196cdd102 *__obf_786a8b420048cd65, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []int16, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_1c53ad4397d4b9ae [3]int32
			var __obf_63bac68a07d23c2c int32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				__obf_9166f977ce6b9e75 := __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49]
				if __obf_9166f977ce6b9e75 != 0 {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 3
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 3 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_1c53ad4397d4b9ae[0] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])
					__obf_1c53ad4397d4b9ae[1] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1])
					__obf_1c53ad4397d4b9ae[2] += int32(__obf_9166f977ce6b9e75) * int32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])
					__obf_63bac68a07d23c2c += int32(__obf_9166f977ce6b9e75)
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*3
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = __obf_b892a4c4c7d416e4(__obf_1c53ad4397d4b9ae[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = __obf_b892a4c4c7d416e4(__obf_1c53ad4397d4b9ae[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = __obf_b892a4c4c7d416e4(__obf_1c53ad4397d4b9ae[2] / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_61f66fd475da9f7c(__obf_b13d16274eb954e2 *__obf_786a8b420048cd65, __obf_6a55e60196cdd102 *__obf_786a8b420048cd65, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_1c53ad4397d4b9ae [3]float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 3
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 3 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_1c53ad4397d4b9ae[0] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])
					__obf_1c53ad4397d4b9ae[1] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1])
					__obf_1c53ad4397d4b9ae[2] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*3
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = __obf_d3cf617e522e2354(__obf_1c53ad4397d4b9ae[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = __obf_d3cf617e522e2354(__obf_1c53ad4397d4b9ae[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = __obf_d3cf617e522e2354(__obf_1c53ad4397d4b9ae[2] / __obf_63bac68a07d23c2c)
		}
	}
}
