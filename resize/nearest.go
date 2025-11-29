/*
Copyright (c) 2014, Charlie Vieth <charlie.vieth@gmail.com>

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

func __obf_d3cf617e522e2354(__obf_c6d53ec5b4505f6a float32) uint8 {
	// Nearest-neighbor values are always
	// positive no need to check lower-bound.
	if __obf_c6d53ec5b4505f6a > 0xfe {
		return 0xff
	}
	return uint8(__obf_c6d53ec5b4505f6a)
}

func __obf_c7dc9f45a410fcb6(__obf_c6d53ec5b4505f6a float32) uint16 {
	if __obf_c6d53ec5b4505f6a > 0xfffe {
		return 0xffff
	}
	return uint16(__obf_c6d53ec5b4505f6a)
}

func __obf_5d671bd01e4ff610(__obf_b13d16274eb954e2 image.Image, __obf_6a55e60196cdd102 *image.RGBA64, __obf_121d528711c040a7 float64, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := range __obf_6a27c875a42baeab {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case __obf_27fc29a27a3bf82f < 0:
						__obf_27fc29a27a3bf82f = 0
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = __obf_7d59a7861d2c3cb4
					}
					__obf_1f6def59ec86a578, __obf_ef636de778967ee2, __obf_4efacd728de97e8f, __obf_a4cd46772704603d := __obf_b13d16274eb954e2.At(__obf_27fc29a27a3bf82f+__obf_b13d16274eb954e2.Bounds().Min.X, __obf_c6d53ec5b4505f6a+__obf_b13d16274eb954e2.Bounds().Min.Y).RGBA()
					__obf_47e29cd4ec4bd91e[0] += float32(__obf_1f6def59ec86a578)
					__obf_47e29cd4ec4bd91e[1] += float32(__obf_ef636de778967ee2)
					__obf_47e29cd4ec4bd91e[2] += float32(__obf_4efacd728de97e8f)
					__obf_47e29cd4ec4bd91e[3] += float32(__obf_a4cd46772704603d)
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_87dd30f656544b12 := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*8
			__obf_1cf6d31eaacee463 := __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+1] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+2] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+3] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+4] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+5] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+6] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+7] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_66d1834bc9607e4b(__obf_b13d16274eb954e2, __obf_6a55e60196cdd102 *image.RGBA, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 4
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 4 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_47e29cd4ec4bd91e[0] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])
					__obf_47e29cd4ec4bd91e[1] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1])
					__obf_47e29cd4ec4bd91e[2] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])
					__obf_47e29cd4ec4bd91e[3] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3])
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*4
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_fe4127dc1da8788e(__obf_b13d16274eb954e2, __obf_6a55e60196cdd102 *image.NRGBA, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_6a27c875a42baeab; __obf_7d43b40e55127c49++ {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 4
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 4 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_47e29cd4ec4bd91e[0] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])
					__obf_47e29cd4ec4bd91e[1] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1])
					__obf_47e29cd4ec4bd91e[2] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])
					__obf_47e29cd4ec4bd91e[3] += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3])
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*4
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = __obf_d3cf617e522e2354(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_228913032261ceca(__obf_b13d16274eb954e2, __obf_6a55e60196cdd102 *image.RGBA64, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := range __obf_6a27c875a42baeab {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 8
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 8 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_47e29cd4ec4bd91e[0] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1]))
					__obf_47e29cd4ec4bd91e[1] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3]))
					__obf_47e29cd4ec4bd91e[2] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+4])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+5]))
					__obf_47e29cd4ec4bd91e[3] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+6])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+7]))
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*8
			__obf_1cf6d31eaacee463 := __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+4] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+5] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+6] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+7] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_dc6b83b67e049b85(__obf_b13d16274eb954e2, __obf_6a55e60196cdd102 *image.NRGBA64, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_47e29cd4ec4bd91e [4]float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := range __obf_6a27c875a42baeab {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 8
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 8 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_47e29cd4ec4bd91e[0] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1]))
					__obf_47e29cd4ec4bd91e[1] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+2])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+3]))
					__obf_47e29cd4ec4bd91e[2] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+4])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+5]))
					__obf_47e29cd4ec4bd91e[3] += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+6])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+7]))
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_ef43535e55679bdb := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*8
			__obf_1cf6d31eaacee463 := __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[0] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+1] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[1] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+2] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+3] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[2] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+4] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+5] = uint8(__obf_1cf6d31eaacee463)
			__obf_1cf6d31eaacee463 = __obf_c7dc9f45a410fcb6(__obf_47e29cd4ec4bd91e[3] / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+6] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_ef43535e55679bdb+7] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}

func __obf_0ca2f5762875ef96(__obf_b13d16274eb954e2, __obf_6a55e60196cdd102 *image.Gray, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_3fe9f7c50ed9bc3c float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := range __obf_6a27c875a42baeab {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case __obf_27fc29a27a3bf82f < 0:
						__obf_27fc29a27a3bf82f = 0
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = __obf_7d59a7861d2c3cb4
					}
					__obf_3fe9f7c50ed9bc3c += float32(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f])
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_87dd30f656544b12 := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a - __obf_85e316700ddb907a.Min.X)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12] = __obf_d3cf617e522e2354(__obf_3fe9f7c50ed9bc3c / __obf_63bac68a07d23c2c)
		}
	}
}

func __obf_525f68f0384d6897(__obf_b13d16274eb954e2, __obf_6a55e60196cdd102 *image.Gray16, __obf_73fc3de4b87d9301 []bool, __obf_87dd30f656544b12 []int, __obf_6a27c875a42baeab int) {
	__obf_85e316700ddb907a := __obf_6a55e60196cdd102.Bounds()
	__obf_7d59a7861d2c3cb4 := __obf_b13d16274eb954e2.Bounds().Dx() - 1

	for __obf_c6d53ec5b4505f6a := __obf_85e316700ddb907a.Min.X; __obf_c6d53ec5b4505f6a < __obf_85e316700ddb907a.Max.X; __obf_c6d53ec5b4505f6a++ {
		__obf_ccd90b77d368ba6c := __obf_b13d16274eb954e2.Pix[__obf_c6d53ec5b4505f6a*__obf_b13d16274eb954e2.Stride:]
		for __obf_eed2203158cb878d := __obf_85e316700ddb907a.Min.Y; __obf_eed2203158cb878d < __obf_85e316700ddb907a.Max.Y; __obf_eed2203158cb878d++ {
			var __obf_3fe9f7c50ed9bc3c float32
			var __obf_63bac68a07d23c2c float32
			__obf_45570ac86f60c691 := __obf_87dd30f656544b12[__obf_eed2203158cb878d]
			__obf_1986f534cb8126b3 := __obf_eed2203158cb878d * __obf_6a27c875a42baeab
			for __obf_7d43b40e55127c49 := range __obf_6a27c875a42baeab {
				if __obf_73fc3de4b87d9301[__obf_1986f534cb8126b3+__obf_7d43b40e55127c49] {
					__obf_27fc29a27a3bf82f := __obf_45570ac86f60c691 + __obf_7d43b40e55127c49
					switch {
					case uint(__obf_27fc29a27a3bf82f) < uint(__obf_7d59a7861d2c3cb4):
						__obf_27fc29a27a3bf82f *= 2
					case __obf_27fc29a27a3bf82f >= __obf_7d59a7861d2c3cb4:
						__obf_27fc29a27a3bf82f = 2 * __obf_7d59a7861d2c3cb4
					default:
						__obf_27fc29a27a3bf82f = 0
					}
					__obf_3fe9f7c50ed9bc3c += float32(uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+0])<<8 | uint16(__obf_ccd90b77d368ba6c[__obf_27fc29a27a3bf82f+1]))
					__obf_63bac68a07d23c2c++
				}
			}
			__obf_87dd30f656544b12 := (__obf_eed2203158cb878d-__obf_85e316700ddb907a.Min.Y)*__obf_6a55e60196cdd102.Stride + (__obf_c6d53ec5b4505f6a-__obf_85e316700ddb907a.Min.X)*2
			__obf_1cf6d31eaacee463 := __obf_c7dc9f45a410fcb6(__obf_3fe9f7c50ed9bc3c / __obf_63bac68a07d23c2c)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+0] = uint8(__obf_1cf6d31eaacee463 >> 8)
			__obf_6a55e60196cdd102.
				Pix[__obf_87dd30f656544b12+1] = uint8(__obf_1cf6d31eaacee463)
		}
	}
}
