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

package __obf_4f71249d3f545aa0

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type __obf_7b573c09d347a918 struct {
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
func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) PixOffset(__obf_2a2c911b5786aa81, __obf_36c97b6a0fef1dad int) int {
	return (__obf_36c97b6a0fef1dad-__obf_6f27b861509320fe.Rect.Min.Y)*__obf_6f27b861509320fe.Stride + (__obf_2a2c911b5786aa81-__obf_6f27b861509320fe.Rect.Min.X)*3
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) Bounds() image.Rectangle {
	return __obf_6f27b861509320fe.Rect
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) ColorModel() color.Model {
	return color.YCbCrModel
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) At(__obf_2a2c911b5786aa81, __obf_36c97b6a0fef1dad int) color.Color {
	if !(image.Point{__obf_2a2c911b5786aa81, __obf_36c97b6a0fef1dad}.In(__obf_6f27b861509320fe.Rect)) {
		return color.YCbCr{}
	}
	__obf_b92089dde6a4a00b := __obf_6f27b861509320fe.PixOffset(__obf_2a2c911b5786aa81, __obf_36c97b6a0fef1dad)
	return color.YCbCr{__obf_6f27b861509320fe.
		Pix[__obf_b92089dde6a4a00b+0], __obf_6f27b861509320fe.
		Pix[__obf_b92089dde6a4a00b+1], __obf_6f27b861509320fe.
		Pix[__obf_b92089dde6a4a00b+2],
	}
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) SubImage(__obf_2be862e501ec49ab image.Rectangle) image.Image {
	__obf_2be862e501ec49ab = __obf_2be862e501ec49ab.Intersect(__obf_6f27b861509320fe.Rect)
	if __obf_2be862e501ec49ab.Empty() {
		return &__obf_7b573c09d347a918{SubsampleRatio: __obf_6f27b861509320fe.SubsampleRatio}
	}
	__obf_b92089dde6a4a00b := __obf_6f27b861509320fe.PixOffset(__obf_2be862e501ec49ab.Min.X, __obf_2be862e501ec49ab.Min.Y)
	return &__obf_7b573c09d347a918{
		Pix:            __obf_6f27b861509320fe.Pix[__obf_b92089dde6a4a00b:],
		Stride:         __obf_6f27b861509320fe.Stride,
		Rect:           __obf_2be862e501ec49ab,
		SubsampleRatio: __obf_6f27b861509320fe.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func __obf_f13836c808b60e84(__obf_2be862e501ec49ab image.Rectangle, __obf_5ce0d738033b048d image.YCbCrSubsampleRatio) *__obf_7b573c09d347a918 {
	__obf_b10e9e44ed77ef29, __obf_a6f6cf61df879a9c := __obf_2be862e501ec49ab.Dx(), __obf_2be862e501ec49ab.Dy()
	__obf_ea8b4489c04acf05 := make([]uint8, 3*__obf_b10e9e44ed77ef29*__obf_a6f6cf61df879a9c)
	return &__obf_7b573c09d347a918{Pix: __obf_ea8b4489c04acf05, Stride: 3 * __obf_b10e9e44ed77ef29, Rect: __obf_2be862e501ec49ab, SubsampleRatio: __obf_5ce0d738033b048d}
}

// Copy of image.YCbCrSubsampleRatio constants - this allows us to support
// older versions of Go where these constants are not defined (i.e. Go 1.4)
const (
	__obf_8ecf813d621bd093 image.YCbCrSubsampleRatio = iota
	__obf_7e14291d2f566e86
	__obf_7da62a15b5433be4
	__obf_b03bbc561d859443
	__obf_1a9d61624566710e
	__obf_3a10df0340313b23
)

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) YCbCr() *image.YCbCr {
	__obf_5dd6c2063d8db0bd := image.NewYCbCr(__obf_6f27b861509320fe.Rect, __obf_6f27b861509320fe.SubsampleRatio)
	switch __obf_5dd6c2063d8db0bd.SubsampleRatio {
	case __obf_7e14291d2f566e86:
		return __obf_6f27b861509320fe.__obf_3be7ae365b156287(__obf_5dd6c2063d8db0bd)
	case __obf_7da62a15b5433be4:
		return __obf_6f27b861509320fe.__obf_29ecaaa3e3f3b283(__obf_5dd6c2063d8db0bd)
	case __obf_b03bbc561d859443:
		return __obf_6f27b861509320fe.__obf_0fc8fb223a5983da(__obf_5dd6c2063d8db0bd)
	case __obf_8ecf813d621bd093:
		return __obf_6f27b861509320fe.__obf_8ac841612329a3fd(__obf_5dd6c2063d8db0bd)
	case __obf_1a9d61624566710e:
		return __obf_6f27b861509320fe.__obf_72375cc32edbf2ff(__obf_5dd6c2063d8db0bd)
	case __obf_3a10df0340313b23:
		return __obf_6f27b861509320fe.__obf_3a076dcddf2e5ad2(__obf_5dd6c2063d8db0bd)
	}
	return __obf_5dd6c2063d8db0bd
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func __obf_25f07534b3928897(__obf_a5eab9a5623e4978 *image.YCbCr) *__obf_7b573c09d347a918 {
	__obf_b10e9e44ed77ef29, __obf_a6f6cf61df879a9c := __obf_a5eab9a5623e4978.Rect.Dx(), __obf_a5eab9a5623e4978.Rect.Dy()
	__obf_6f27b861509320fe := __obf_7b573c09d347a918{
		Pix:            make([]uint8, 3*__obf_b10e9e44ed77ef29*__obf_a6f6cf61df879a9c),
		Stride:         3 * __obf_b10e9e44ed77ef29,
		Rect:           image.Rect(0, 0, __obf_b10e9e44ed77ef29, __obf_a6f6cf61df879a9c),
		SubsampleRatio: __obf_a5eab9a5623e4978.SubsampleRatio,
	}
	switch __obf_a5eab9a5623e4978.SubsampleRatio {
	case __obf_7e14291d2f566e86:
		return __obf_5df6a0d32a707459(__obf_a5eab9a5623e4978, &__obf_6f27b861509320fe)
	case __obf_7da62a15b5433be4:
		return __obf_9fefbaf8333d359a(__obf_a5eab9a5623e4978, &__obf_6f27b861509320fe)
	case __obf_b03bbc561d859443:
		return __obf_19e4e70813f08948(__obf_a5eab9a5623e4978, &__obf_6f27b861509320fe)
	case __obf_8ecf813d621bd093:
		return __obf_102c15cb84fe74d6(__obf_a5eab9a5623e4978, &__obf_6f27b861509320fe)
	case __obf_1a9d61624566710e:
		return __obf_b634bb4f76d72458(__obf_a5eab9a5623e4978, &__obf_6f27b861509320fe)
	case __obf_3a10df0340313b23:
		return __obf_377eec7f55caa367(__obf_a5eab9a5623e4978, &__obf_6f27b861509320fe)
	}
	return &__obf_6f27b861509320fe
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) __obf_3be7ae365b156287(__obf_5dd6c2063d8db0bd *image.YCbCr) *image.YCbCr {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_5dd6c2063d8db0bd.Y
	Cb := __obf_5dd6c2063d8db0bd.Cb
	Cr := __obf_5dd6c2063d8db0bd.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_5dd6c2063d8db0bd.Rect.Max.Y-__obf_5dd6c2063d8db0bd.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.YStride
		__obf_f57ac876f7ec649e := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_5dd6c2063d8db0bd.Rect.Max.X-__obf_5dd6c2063d8db0bd.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/2
			Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81] = Pix[__obf_f2e4a8d9588b5c7b+0]
			Cb[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+1]
			Cr[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+2]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_5dd6c2063d8db0bd
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) __obf_29ecaaa3e3f3b283(__obf_5dd6c2063d8db0bd *image.YCbCr) *image.YCbCr {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_5dd6c2063d8db0bd.Y
	Cb := __obf_5dd6c2063d8db0bd.Cb
	Cr := __obf_5dd6c2063d8db0bd.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_5dd6c2063d8db0bd.Rect.Max.Y-__obf_5dd6c2063d8db0bd.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.YStride
		__obf_f57ac876f7ec649e := (__obf_36c97b6a0fef1dad / 2) * __obf_5dd6c2063d8db0bd.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_5dd6c2063d8db0bd.Rect.Max.X-__obf_5dd6c2063d8db0bd.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/2
			Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81] = Pix[__obf_f2e4a8d9588b5c7b+0]
			Cb[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+1]
			Cr[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+2]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_5dd6c2063d8db0bd
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) __obf_0fc8fb223a5983da(__obf_5dd6c2063d8db0bd *image.YCbCr) *image.YCbCr {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_5dd6c2063d8db0bd.Y
	Cb := __obf_5dd6c2063d8db0bd.Cb
	Cr := __obf_5dd6c2063d8db0bd.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_5dd6c2063d8db0bd.Rect.Max.Y-__obf_5dd6c2063d8db0bd.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.YStride
		__obf_f57ac876f7ec649e := (__obf_36c97b6a0fef1dad / 2) * __obf_5dd6c2063d8db0bd.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_5dd6c2063d8db0bd.Rect.Max.X-__obf_5dd6c2063d8db0bd.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81
			Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81] = Pix[__obf_f2e4a8d9588b5c7b+0]
			Cb[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+1]
			Cr[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+2]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_5dd6c2063d8db0bd
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) __obf_8ac841612329a3fd(__obf_5dd6c2063d8db0bd *image.YCbCr) *image.YCbCr {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_5dd6c2063d8db0bd.Y
	Cb := __obf_5dd6c2063d8db0bd.Cb
	Cr := __obf_5dd6c2063d8db0bd.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_5dd6c2063d8db0bd.Rect.Max.Y-__obf_5dd6c2063d8db0bd.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.YStride
		__obf_f57ac876f7ec649e := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_5dd6c2063d8db0bd.Rect.Max.X-__obf_5dd6c2063d8db0bd.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81
			Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81] = Pix[__obf_f2e4a8d9588b5c7b+0]
			Cb[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+1]
			Cr[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+2]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_5dd6c2063d8db0bd
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) __obf_72375cc32edbf2ff(__obf_5dd6c2063d8db0bd *image.YCbCr) *image.YCbCr {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_5dd6c2063d8db0bd.Y
	Cb := __obf_5dd6c2063d8db0bd.Cb
	Cr := __obf_5dd6c2063d8db0bd.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_5dd6c2063d8db0bd.Rect.Max.Y-__obf_5dd6c2063d8db0bd.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.YStride
		__obf_f57ac876f7ec649e := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_5dd6c2063d8db0bd.Rect.Max.X-__obf_5dd6c2063d8db0bd.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/4
			Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81] = Pix[__obf_f2e4a8d9588b5c7b+0]
			Cb[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+1]
			Cr[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+2]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_5dd6c2063d8db0bd
}

func (__obf_6f27b861509320fe *__obf_7b573c09d347a918) __obf_3a076dcddf2e5ad2(__obf_5dd6c2063d8db0bd *image.YCbCr) *image.YCbCr {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_5dd6c2063d8db0bd.Y
	Cb := __obf_5dd6c2063d8db0bd.Cb
	Cr := __obf_5dd6c2063d8db0bd.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_5dd6c2063d8db0bd.Rect.Max.Y-__obf_5dd6c2063d8db0bd.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_5dd6c2063d8db0bd.YStride
		__obf_f57ac876f7ec649e := (__obf_36c97b6a0fef1dad / 2) * __obf_5dd6c2063d8db0bd.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_5dd6c2063d8db0bd.Rect.Max.X-__obf_5dd6c2063d8db0bd.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/4
			Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81] = Pix[__obf_f2e4a8d9588b5c7b+0]
			Cb[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+1]
			Cr[__obf_08a465ba68a06bdf] = Pix[__obf_f2e4a8d9588b5c7b+2]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_5dd6c2063d8db0bd
}

func __obf_5df6a0d32a707459(__obf_a5eab9a5623e4978 *image.YCbCr, __obf_6f27b861509320fe *__obf_7b573c09d347a918) *__obf_7b573c09d347a918 {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_a5eab9a5623e4978.Y
	Cb := __obf_a5eab9a5623e4978.Cb
	Cr := __obf_a5eab9a5623e4978.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_a5eab9a5623e4978.Rect.Max.Y-__obf_a5eab9a5623e4978.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.YStride
		__obf_f57ac876f7ec649e := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_a5eab9a5623e4978.Rect.Max.X-__obf_a5eab9a5623e4978.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/2
			Pix[__obf_f2e4a8d9588b5c7b+0] = Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81]
			Pix[__obf_f2e4a8d9588b5c7b+1] = Cb[__obf_08a465ba68a06bdf]
			Pix[__obf_f2e4a8d9588b5c7b+2] = Cr[__obf_08a465ba68a06bdf]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_6f27b861509320fe
}

func __obf_9fefbaf8333d359a(__obf_a5eab9a5623e4978 *image.YCbCr, __obf_6f27b861509320fe *__obf_7b573c09d347a918) *__obf_7b573c09d347a918 {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_a5eab9a5623e4978.Y
	Cb := __obf_a5eab9a5623e4978.Cb
	Cr := __obf_a5eab9a5623e4978.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_a5eab9a5623e4978.Rect.Max.Y-__obf_a5eab9a5623e4978.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.YStride
		__obf_f57ac876f7ec649e := (__obf_36c97b6a0fef1dad / 2) * __obf_a5eab9a5623e4978.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_a5eab9a5623e4978.Rect.Max.X-__obf_a5eab9a5623e4978.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/2
			Pix[__obf_f2e4a8d9588b5c7b+0] = Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81]
			Pix[__obf_f2e4a8d9588b5c7b+1] = Cb[__obf_08a465ba68a06bdf]
			Pix[__obf_f2e4a8d9588b5c7b+2] = Cr[__obf_08a465ba68a06bdf]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_6f27b861509320fe
}

func __obf_19e4e70813f08948(__obf_a5eab9a5623e4978 *image.YCbCr, __obf_6f27b861509320fe *__obf_7b573c09d347a918) *__obf_7b573c09d347a918 {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_a5eab9a5623e4978.Y
	Cb := __obf_a5eab9a5623e4978.Cb
	Cr := __obf_a5eab9a5623e4978.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_a5eab9a5623e4978.Rect.Max.Y-__obf_a5eab9a5623e4978.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.YStride
		__obf_f57ac876f7ec649e := (__obf_36c97b6a0fef1dad / 2) * __obf_a5eab9a5623e4978.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_a5eab9a5623e4978.Rect.Max.X-__obf_a5eab9a5623e4978.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81
			Pix[__obf_f2e4a8d9588b5c7b+0] = Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81]
			Pix[__obf_f2e4a8d9588b5c7b+1] = Cb[__obf_08a465ba68a06bdf]
			Pix[__obf_f2e4a8d9588b5c7b+2] = Cr[__obf_08a465ba68a06bdf]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_6f27b861509320fe
}

func __obf_102c15cb84fe74d6(__obf_a5eab9a5623e4978 *image.YCbCr, __obf_6f27b861509320fe *__obf_7b573c09d347a918) *__obf_7b573c09d347a918 {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_a5eab9a5623e4978.Y
	Cb := __obf_a5eab9a5623e4978.Cb
	Cr := __obf_a5eab9a5623e4978.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_a5eab9a5623e4978.Rect.Max.Y-__obf_a5eab9a5623e4978.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.YStride
		__obf_f57ac876f7ec649e := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_a5eab9a5623e4978.Rect.Max.X-__obf_a5eab9a5623e4978.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81
			Pix[__obf_f2e4a8d9588b5c7b+0] = Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81]
			Pix[__obf_f2e4a8d9588b5c7b+1] = Cb[__obf_08a465ba68a06bdf]
			Pix[__obf_f2e4a8d9588b5c7b+2] = Cr[__obf_08a465ba68a06bdf]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_6f27b861509320fe
}

func __obf_b634bb4f76d72458(__obf_a5eab9a5623e4978 *image.YCbCr, __obf_6f27b861509320fe *__obf_7b573c09d347a918) *__obf_7b573c09d347a918 {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_a5eab9a5623e4978.Y
	Cb := __obf_a5eab9a5623e4978.Cb
	Cr := __obf_a5eab9a5623e4978.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_a5eab9a5623e4978.Rect.Max.Y-__obf_a5eab9a5623e4978.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.YStride
		__obf_f57ac876f7ec649e := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_a5eab9a5623e4978.Rect.Max.X-__obf_a5eab9a5623e4978.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/4
			Pix[__obf_f2e4a8d9588b5c7b+0] = Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81]
			Pix[__obf_f2e4a8d9588b5c7b+1] = Cb[__obf_08a465ba68a06bdf]
			Pix[__obf_f2e4a8d9588b5c7b+2] = Cr[__obf_08a465ba68a06bdf]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_6f27b861509320fe
}

func __obf_377eec7f55caa367(__obf_a5eab9a5623e4978 *image.YCbCr, __obf_6f27b861509320fe *__obf_7b573c09d347a918) *__obf_7b573c09d347a918 {
	var __obf_f2e4a8d9588b5c7b int
	Pix := __obf_6f27b861509320fe.Pix
	Y := __obf_a5eab9a5623e4978.Y
	Cb := __obf_a5eab9a5623e4978.Cb
	Cr := __obf_a5eab9a5623e4978.Cr
	for __obf_36c97b6a0fef1dad := 0; __obf_36c97b6a0fef1dad < __obf_a5eab9a5623e4978.Rect.Max.Y-__obf_a5eab9a5623e4978.Rect.Min.Y; __obf_36c97b6a0fef1dad++ {
		__obf_91825b6af879bd30 := __obf_36c97b6a0fef1dad * __obf_a5eab9a5623e4978.YStride
		__obf_f57ac876f7ec649e := (__obf_36c97b6a0fef1dad / 2) * __obf_a5eab9a5623e4978.CStride
		for __obf_2a2c911b5786aa81 := 0; __obf_2a2c911b5786aa81 < __obf_a5eab9a5623e4978.Rect.Max.X-__obf_a5eab9a5623e4978.Rect.Min.X; __obf_2a2c911b5786aa81++ {
			__obf_08a465ba68a06bdf := __obf_f57ac876f7ec649e + __obf_2a2c911b5786aa81/4
			Pix[__obf_f2e4a8d9588b5c7b+0] = Y[__obf_91825b6af879bd30+__obf_2a2c911b5786aa81]
			Pix[__obf_f2e4a8d9588b5c7b+1] = Cb[__obf_08a465ba68a06bdf]
			Pix[__obf_f2e4a8d9588b5c7b+2] = Cr[__obf_08a465ba68a06bdf]
			__obf_f2e4a8d9588b5c7b += 3
		}
	}
	return __obf_6f27b861509320fe
}
