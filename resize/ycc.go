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

package __obf_d24101647651f4c3

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type __obf_489a9befc4898050 struct {
	// Pix holds the image's pixels, in Y, Cb, Cr order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
	// SubsampleRatio is the subsample ratio of the original YCbCr image.
	SubsampleRatio image.YCbCrSubsampleRatio
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) PixOffset(__obf_88d1a0c3dc734ee1, __obf_097022f910960674 int) int {
	return (__obf_097022f910960674-__obf_a78e262cc2a46e75.Rect.Min.Y)*__obf_a78e262cc2a46e75.Stride + (__obf_88d1a0c3dc734ee1-__obf_a78e262cc2a46e75.Rect.Min.X)*3
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) Bounds() image.Rectangle {
	return __obf_a78e262cc2a46e75.Rect
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) ColorModel() color.Model {
	return color.YCbCrModel
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) At(__obf_88d1a0c3dc734ee1, __obf_097022f910960674 int) color.Color {
	if !(image.Point{__obf_88d1a0c3dc734ee1, __obf_097022f910960674}.In(__obf_a78e262cc2a46e75.Rect)) {
		return color.YCbCr{}
	}
	__obf_41043e9a68df34f8 := __obf_a78e262cc2a46e75.PixOffset(__obf_88d1a0c3dc734ee1, __obf_097022f910960674)
	return color.YCbCr{__obf_a78e262cc2a46e75.
		Pix[__obf_41043e9a68df34f8+0], __obf_a78e262cc2a46e75.
		Pix[__obf_41043e9a68df34f8+1], __obf_a78e262cc2a46e75.
		Pix[__obf_41043e9a68df34f8+2],
	}
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) SubImage(__obf_13d76cd467248ddd image.Rectangle) image.Image {
	__obf_13d76cd467248ddd = __obf_13d76cd467248ddd.Intersect(__obf_a78e262cc2a46e75.Rect)
	if __obf_13d76cd467248ddd.Empty() {
		return &__obf_489a9befc4898050{SubsampleRatio: __obf_a78e262cc2a46e75.SubsampleRatio}
	}
	__obf_41043e9a68df34f8 := __obf_a78e262cc2a46e75.PixOffset(__obf_13d76cd467248ddd.Min.X, __obf_13d76cd467248ddd.Min.Y)
	return &__obf_489a9befc4898050{
		Pix:            __obf_a78e262cc2a46e75.Pix[__obf_41043e9a68df34f8:],
		Stride:         __obf_a78e262cc2a46e75.Stride,
		Rect:           __obf_13d76cd467248ddd,
		SubsampleRatio: __obf_a78e262cc2a46e75.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func __obf_526e91fd49ed301d(__obf_13d76cd467248ddd image.Rectangle, __obf_9efacb8361bfa66b image.YCbCrSubsampleRatio) *__obf_489a9befc4898050 {
	__obf_ca1efc628c04c544, __obf_ffe12cc399cae12c := __obf_13d76cd467248ddd.Dx(), __obf_13d76cd467248ddd.Dy()
	__obf_e2d3181f4cbc4f17 := make([]uint8, 3*__obf_ca1efc628c04c544*__obf_ffe12cc399cae12c)
	return &__obf_489a9befc4898050{Pix: __obf_e2d3181f4cbc4f17, Stride: 3 * __obf_ca1efc628c04c544, Rect: __obf_13d76cd467248ddd, SubsampleRatio: __obf_9efacb8361bfa66b}
}

// Copy of image.YCbCrSubsampleRatio constants - this allows us to support
// older versions of Go where these constants are not defined (i.e. Go 1.4)
const (
	__obf_3b6ebb64431af374 image.YCbCrSubsampleRatio = iota
	__obf_a48a314d29039d42
	__obf_253ffc1f955c8ecc
	__obf_95c58aef23f99370
	__obf_3f164eae143eca12
	__obf_3f3c3a54e4918d8a
)

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) YCbCr() *image.YCbCr {
	__obf_cfdb4775d1615df9 := image.NewYCbCr(__obf_a78e262cc2a46e75.Rect, __obf_a78e262cc2a46e75.SubsampleRatio)
	switch __obf_cfdb4775d1615df9.SubsampleRatio {
	case __obf_a48a314d29039d42:
		return __obf_a78e262cc2a46e75.__obf_81c47f62df5bc07c(__obf_cfdb4775d1615df9)
	case __obf_253ffc1f955c8ecc:
		return __obf_a78e262cc2a46e75.__obf_5e916b32cf366194(__obf_cfdb4775d1615df9)
	case __obf_95c58aef23f99370:
		return __obf_a78e262cc2a46e75.__obf_93077fdb79d9cab3(__obf_cfdb4775d1615df9)
	case __obf_3b6ebb64431af374:
		return __obf_a78e262cc2a46e75.__obf_bfbd244dd8fab9d1(__obf_cfdb4775d1615df9)
	case __obf_3f164eae143eca12:
		return __obf_a78e262cc2a46e75.__obf_b206122605909b20(__obf_cfdb4775d1615df9)
	case __obf_3f3c3a54e4918d8a:
		return __obf_a78e262cc2a46e75.__obf_3410bdd7129ba83d(__obf_cfdb4775d1615df9)
	}
	return __obf_cfdb4775d1615df9
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func __obf_47343abcf82d0c31(__obf_2a764779fc1401ce *image.YCbCr) *__obf_489a9befc4898050 {
	__obf_ca1efc628c04c544, __obf_ffe12cc399cae12c := __obf_2a764779fc1401ce.Rect.Dx(), __obf_2a764779fc1401ce.Rect.Dy()
	__obf_a78e262cc2a46e75 := __obf_489a9befc4898050{
		Pix:            make([]uint8, 3*__obf_ca1efc628c04c544*__obf_ffe12cc399cae12c),
		Stride:         3 * __obf_ca1efc628c04c544,
		Rect:           image.Rect(0, 0, __obf_ca1efc628c04c544, __obf_ffe12cc399cae12c),
		SubsampleRatio: __obf_2a764779fc1401ce.SubsampleRatio,
	}
	switch __obf_2a764779fc1401ce.SubsampleRatio {
	case __obf_a48a314d29039d42:
		return __obf_b90b6a0a6de0b679(__obf_2a764779fc1401ce, &__obf_a78e262cc2a46e75)
	case __obf_253ffc1f955c8ecc:
		return __obf_3a93c813f2ba455a(__obf_2a764779fc1401ce, &__obf_a78e262cc2a46e75)
	case __obf_95c58aef23f99370:
		return __obf_c185f9af89f56e49(__obf_2a764779fc1401ce, &__obf_a78e262cc2a46e75)
	case __obf_3b6ebb64431af374:
		return __obf_afdc0c3c0ec918d5(__obf_2a764779fc1401ce, &__obf_a78e262cc2a46e75)
	case __obf_3f164eae143eca12:
		return __obf_bf7f786f18f8b9b4(__obf_2a764779fc1401ce, &__obf_a78e262cc2a46e75)
	case __obf_3f3c3a54e4918d8a:
		return __obf_9d88ae4f3df9b073(__obf_2a764779fc1401ce, &__obf_a78e262cc2a46e75)
	}
	return &__obf_a78e262cc2a46e75
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) __obf_81c47f62df5bc07c(__obf_cfdb4775d1615df9 *image.YCbCr) *image.YCbCr {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_cfdb4775d1615df9.Y
	Cb := __obf_cfdb4775d1615df9.Cb
	Cr := __obf_cfdb4775d1615df9.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_cfdb4775d1615df9.Rect.Max.Y-__obf_cfdb4775d1615df9.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_cfdb4775d1615df9.YStride
		__obf_a33c8f680552a9ab := __obf_097022f910960674 * __obf_cfdb4775d1615df9.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_cfdb4775d1615df9.Rect.Max.X-__obf_cfdb4775d1615df9.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/2
			Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1] = Pix[__obf_8cf7887eba79c0fd+0]
			Cb[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+1]
			Cr[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+2]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_cfdb4775d1615df9
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) __obf_5e916b32cf366194(__obf_cfdb4775d1615df9 *image.YCbCr) *image.YCbCr {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_cfdb4775d1615df9.Y
	Cb := __obf_cfdb4775d1615df9.Cb
	Cr := __obf_cfdb4775d1615df9.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_cfdb4775d1615df9.Rect.Max.Y-__obf_cfdb4775d1615df9.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_cfdb4775d1615df9.YStride
		__obf_a33c8f680552a9ab := (__obf_097022f910960674 / 2) * __obf_cfdb4775d1615df9.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_cfdb4775d1615df9.Rect.Max.X-__obf_cfdb4775d1615df9.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/2
			Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1] = Pix[__obf_8cf7887eba79c0fd+0]
			Cb[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+1]
			Cr[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+2]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_cfdb4775d1615df9
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) __obf_93077fdb79d9cab3(__obf_cfdb4775d1615df9 *image.YCbCr) *image.YCbCr {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_cfdb4775d1615df9.Y
	Cb := __obf_cfdb4775d1615df9.Cb
	Cr := __obf_cfdb4775d1615df9.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_cfdb4775d1615df9.Rect.Max.Y-__obf_cfdb4775d1615df9.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_cfdb4775d1615df9.YStride
		__obf_a33c8f680552a9ab := (__obf_097022f910960674 / 2) * __obf_cfdb4775d1615df9.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_cfdb4775d1615df9.Rect.Max.X-__obf_cfdb4775d1615df9.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1
			Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1] = Pix[__obf_8cf7887eba79c0fd+0]
			Cb[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+1]
			Cr[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+2]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_cfdb4775d1615df9
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) __obf_bfbd244dd8fab9d1(__obf_cfdb4775d1615df9 *image.YCbCr) *image.YCbCr {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_cfdb4775d1615df9.Y
	Cb := __obf_cfdb4775d1615df9.Cb
	Cr := __obf_cfdb4775d1615df9.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_cfdb4775d1615df9.Rect.Max.Y-__obf_cfdb4775d1615df9.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_cfdb4775d1615df9.YStride
		__obf_a33c8f680552a9ab := __obf_097022f910960674 * __obf_cfdb4775d1615df9.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_cfdb4775d1615df9.Rect.Max.X-__obf_cfdb4775d1615df9.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1
			Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1] = Pix[__obf_8cf7887eba79c0fd+0]
			Cb[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+1]
			Cr[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+2]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_cfdb4775d1615df9
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) __obf_b206122605909b20(__obf_cfdb4775d1615df9 *image.YCbCr) *image.YCbCr {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_cfdb4775d1615df9.Y
	Cb := __obf_cfdb4775d1615df9.Cb
	Cr := __obf_cfdb4775d1615df9.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_cfdb4775d1615df9.Rect.Max.Y-__obf_cfdb4775d1615df9.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_cfdb4775d1615df9.YStride
		__obf_a33c8f680552a9ab := __obf_097022f910960674 * __obf_cfdb4775d1615df9.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_cfdb4775d1615df9.Rect.Max.X-__obf_cfdb4775d1615df9.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/4
			Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1] = Pix[__obf_8cf7887eba79c0fd+0]
			Cb[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+1]
			Cr[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+2]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_cfdb4775d1615df9
}

func (__obf_a78e262cc2a46e75 *__obf_489a9befc4898050) __obf_3410bdd7129ba83d(__obf_cfdb4775d1615df9 *image.YCbCr) *image.YCbCr {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_cfdb4775d1615df9.Y
	Cb := __obf_cfdb4775d1615df9.Cb
	Cr := __obf_cfdb4775d1615df9.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_cfdb4775d1615df9.Rect.Max.Y-__obf_cfdb4775d1615df9.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_cfdb4775d1615df9.YStride
		__obf_a33c8f680552a9ab := (__obf_097022f910960674 / 2) * __obf_cfdb4775d1615df9.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_cfdb4775d1615df9.Rect.Max.X-__obf_cfdb4775d1615df9.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/4
			Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1] = Pix[__obf_8cf7887eba79c0fd+0]
			Cb[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+1]
			Cr[__obf_c3764290f5437417] = Pix[__obf_8cf7887eba79c0fd+2]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_cfdb4775d1615df9
}

func __obf_b90b6a0a6de0b679(__obf_2a764779fc1401ce *image.YCbCr, __obf_a78e262cc2a46e75 *__obf_489a9befc4898050) *__obf_489a9befc4898050 {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_2a764779fc1401ce.Y
	Cb := __obf_2a764779fc1401ce.Cb
	Cr := __obf_2a764779fc1401ce.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_2a764779fc1401ce.Rect.Max.Y-__obf_2a764779fc1401ce.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_2a764779fc1401ce.YStride
		__obf_a33c8f680552a9ab := __obf_097022f910960674 * __obf_2a764779fc1401ce.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_2a764779fc1401ce.Rect.Max.X-__obf_2a764779fc1401ce.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/2
			Pix[__obf_8cf7887eba79c0fd+0] = Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1]
			Pix[__obf_8cf7887eba79c0fd+1] = Cb[__obf_c3764290f5437417]
			Pix[__obf_8cf7887eba79c0fd+2] = Cr[__obf_c3764290f5437417]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_a78e262cc2a46e75
}

func __obf_3a93c813f2ba455a(__obf_2a764779fc1401ce *image.YCbCr, __obf_a78e262cc2a46e75 *__obf_489a9befc4898050) *__obf_489a9befc4898050 {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_2a764779fc1401ce.Y
	Cb := __obf_2a764779fc1401ce.Cb
	Cr := __obf_2a764779fc1401ce.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_2a764779fc1401ce.Rect.Max.Y-__obf_2a764779fc1401ce.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_2a764779fc1401ce.YStride
		__obf_a33c8f680552a9ab := (__obf_097022f910960674 / 2) * __obf_2a764779fc1401ce.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_2a764779fc1401ce.Rect.Max.X-__obf_2a764779fc1401ce.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/2
			Pix[__obf_8cf7887eba79c0fd+0] = Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1]
			Pix[__obf_8cf7887eba79c0fd+1] = Cb[__obf_c3764290f5437417]
			Pix[__obf_8cf7887eba79c0fd+2] = Cr[__obf_c3764290f5437417]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_a78e262cc2a46e75
}

func __obf_c185f9af89f56e49(__obf_2a764779fc1401ce *image.YCbCr, __obf_a78e262cc2a46e75 *__obf_489a9befc4898050) *__obf_489a9befc4898050 {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_2a764779fc1401ce.Y
	Cb := __obf_2a764779fc1401ce.Cb
	Cr := __obf_2a764779fc1401ce.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_2a764779fc1401ce.Rect.Max.Y-__obf_2a764779fc1401ce.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_2a764779fc1401ce.YStride
		__obf_a33c8f680552a9ab := (__obf_097022f910960674 / 2) * __obf_2a764779fc1401ce.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_2a764779fc1401ce.Rect.Max.X-__obf_2a764779fc1401ce.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1
			Pix[__obf_8cf7887eba79c0fd+0] = Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1]
			Pix[__obf_8cf7887eba79c0fd+1] = Cb[__obf_c3764290f5437417]
			Pix[__obf_8cf7887eba79c0fd+2] = Cr[__obf_c3764290f5437417]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_a78e262cc2a46e75
}

func __obf_afdc0c3c0ec918d5(__obf_2a764779fc1401ce *image.YCbCr, __obf_a78e262cc2a46e75 *__obf_489a9befc4898050) *__obf_489a9befc4898050 {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_2a764779fc1401ce.Y
	Cb := __obf_2a764779fc1401ce.Cb
	Cr := __obf_2a764779fc1401ce.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_2a764779fc1401ce.Rect.Max.Y-__obf_2a764779fc1401ce.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_2a764779fc1401ce.YStride
		__obf_a33c8f680552a9ab := __obf_097022f910960674 * __obf_2a764779fc1401ce.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_2a764779fc1401ce.Rect.Max.X-__obf_2a764779fc1401ce.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1
			Pix[__obf_8cf7887eba79c0fd+0] = Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1]
			Pix[__obf_8cf7887eba79c0fd+1] = Cb[__obf_c3764290f5437417]
			Pix[__obf_8cf7887eba79c0fd+2] = Cr[__obf_c3764290f5437417]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_a78e262cc2a46e75
}

func __obf_bf7f786f18f8b9b4(__obf_2a764779fc1401ce *image.YCbCr, __obf_a78e262cc2a46e75 *__obf_489a9befc4898050) *__obf_489a9befc4898050 {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_2a764779fc1401ce.Y
	Cb := __obf_2a764779fc1401ce.Cb
	Cr := __obf_2a764779fc1401ce.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_2a764779fc1401ce.Rect.Max.Y-__obf_2a764779fc1401ce.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_2a764779fc1401ce.YStride
		__obf_a33c8f680552a9ab := __obf_097022f910960674 * __obf_2a764779fc1401ce.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_2a764779fc1401ce.Rect.Max.X-__obf_2a764779fc1401ce.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/4
			Pix[__obf_8cf7887eba79c0fd+0] = Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1]
			Pix[__obf_8cf7887eba79c0fd+1] = Cb[__obf_c3764290f5437417]
			Pix[__obf_8cf7887eba79c0fd+2] = Cr[__obf_c3764290f5437417]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_a78e262cc2a46e75
}

func __obf_9d88ae4f3df9b073(__obf_2a764779fc1401ce *image.YCbCr, __obf_a78e262cc2a46e75 *__obf_489a9befc4898050) *__obf_489a9befc4898050 {
	var __obf_8cf7887eba79c0fd int
	Pix := __obf_a78e262cc2a46e75.Pix
	Y := __obf_2a764779fc1401ce.Y
	Cb := __obf_2a764779fc1401ce.Cb
	Cr := __obf_2a764779fc1401ce.Cr
	for __obf_097022f910960674 := 0; __obf_097022f910960674 < __obf_2a764779fc1401ce.Rect.Max.Y-__obf_2a764779fc1401ce.Rect.Min.Y; __obf_097022f910960674++ {
		__obf_658d180b67713d2d := __obf_097022f910960674 * __obf_2a764779fc1401ce.YStride
		__obf_a33c8f680552a9ab := (__obf_097022f910960674 / 2) * __obf_2a764779fc1401ce.CStride
		for __obf_88d1a0c3dc734ee1 := 0; __obf_88d1a0c3dc734ee1 < __obf_2a764779fc1401ce.Rect.Max.X-__obf_2a764779fc1401ce.Rect.Min.X; __obf_88d1a0c3dc734ee1++ {
			__obf_c3764290f5437417 := __obf_a33c8f680552a9ab + __obf_88d1a0c3dc734ee1/4
			Pix[__obf_8cf7887eba79c0fd+0] = Y[__obf_658d180b67713d2d+__obf_88d1a0c3dc734ee1]
			Pix[__obf_8cf7887eba79c0fd+1] = Cb[__obf_c3764290f5437417]
			Pix[__obf_8cf7887eba79c0fd+2] = Cr[__obf_c3764290f5437417]
			__obf_8cf7887eba79c0fd += 3
		}
	}
	return __obf_a78e262cc2a46e75
}
