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

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type __obf_786a8b420048cd65 struct {
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
func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) PixOffset(__obf_c6d53ec5b4505f6a, __obf_eed2203158cb878d int) int {
	return (__obf_eed2203158cb878d-__obf_1c53ad4397d4b9ae.Rect.Min.Y)*__obf_1c53ad4397d4b9ae.Stride + (__obf_c6d53ec5b4505f6a-__obf_1c53ad4397d4b9ae.Rect.Min.X)*3
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) Bounds() image.Rectangle {
	return __obf_1c53ad4397d4b9ae.Rect
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) ColorModel() color.Model {
	return color.YCbCrModel
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) At(__obf_c6d53ec5b4505f6a, __obf_eed2203158cb878d int) color.Color {
	if !(image.Point{__obf_c6d53ec5b4505f6a, __obf_eed2203158cb878d}.In(__obf_1c53ad4397d4b9ae.Rect)) {
		return color.YCbCr{}
	}
	__obf_7d43b40e55127c49 := __obf_1c53ad4397d4b9ae.PixOffset(__obf_c6d53ec5b4505f6a, __obf_eed2203158cb878d)
	return color.YCbCr{__obf_1c53ad4397d4b9ae.
		Pix[__obf_7d43b40e55127c49+0], __obf_1c53ad4397d4b9ae.
		Pix[__obf_7d43b40e55127c49+1], __obf_1c53ad4397d4b9ae.
		Pix[__obf_7d43b40e55127c49+2],
	}
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) SubImage(__obf_1f6def59ec86a578 image.Rectangle) image.Image {
	__obf_1f6def59ec86a578 = __obf_1f6def59ec86a578.Intersect(__obf_1c53ad4397d4b9ae.Rect)
	if __obf_1f6def59ec86a578.Empty() {
		return &__obf_786a8b420048cd65{SubsampleRatio: __obf_1c53ad4397d4b9ae.SubsampleRatio}
	}
	__obf_7d43b40e55127c49 := __obf_1c53ad4397d4b9ae.PixOffset(__obf_1f6def59ec86a578.Min.X, __obf_1f6def59ec86a578.Min.Y)
	return &__obf_786a8b420048cd65{
		Pix:            __obf_1c53ad4397d4b9ae.Pix[__obf_7d43b40e55127c49:],
		Stride:         __obf_1c53ad4397d4b9ae.Stride,
		Rect:           __obf_1f6def59ec86a578,
		SubsampleRatio: __obf_1c53ad4397d4b9ae.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func __obf_1b258dd421287100(__obf_1f6def59ec86a578 image.Rectangle, __obf_e8883f0649f4817e image.YCbCrSubsampleRatio) *__obf_786a8b420048cd65 {
	__obf_5358e7a9db19986b, __obf_256b49a4ea358eeb := __obf_1f6def59ec86a578.Dx(), __obf_1f6def59ec86a578.Dy()
	__obf_42bdc1cc6fe080d3 := make([]uint8, 3*__obf_5358e7a9db19986b*__obf_256b49a4ea358eeb)
	return &__obf_786a8b420048cd65{Pix: __obf_42bdc1cc6fe080d3, Stride: 3 * __obf_5358e7a9db19986b, Rect: __obf_1f6def59ec86a578, SubsampleRatio: __obf_e8883f0649f4817e}
}

// Copy of image.YCbCrSubsampleRatio constants - this allows us to support
// older versions of Go where these constants are not defined (i.e. Go 1.4)
const (
	__obf_02ac68fc9cbc8441 image.YCbCrSubsampleRatio = iota
	__obf_c331ec71612cf606
	__obf_fd1adb463076babb
	__obf_6efc98209b79485d
	__obf_ec8fa2b0adacb378
	__obf_e4ffc790bf709051
)

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) YCbCr() *image.YCbCr {
	__obf_fcc28d0eb472022e := image.NewYCbCr(__obf_1c53ad4397d4b9ae.Rect, __obf_1c53ad4397d4b9ae.SubsampleRatio)
	switch __obf_fcc28d0eb472022e.SubsampleRatio {
	case __obf_c331ec71612cf606:
		return __obf_1c53ad4397d4b9ae.__obf_860222aeb0ce640b(__obf_fcc28d0eb472022e)
	case __obf_fd1adb463076babb:
		return __obf_1c53ad4397d4b9ae.__obf_bda6a895907a47e6(__obf_fcc28d0eb472022e)
	case __obf_6efc98209b79485d:
		return __obf_1c53ad4397d4b9ae.__obf_31cdd1f60e4e7e22(__obf_fcc28d0eb472022e)
	case __obf_02ac68fc9cbc8441:
		return __obf_1c53ad4397d4b9ae.__obf_0f2b33a269f54728(__obf_fcc28d0eb472022e)
	case __obf_ec8fa2b0adacb378:
		return __obf_1c53ad4397d4b9ae.__obf_2967d8f3c994be5e(__obf_fcc28d0eb472022e)
	case __obf_e4ffc790bf709051:
		return __obf_1c53ad4397d4b9ae.__obf_08210d27d9b963ce(__obf_fcc28d0eb472022e)
	}
	return __obf_fcc28d0eb472022e
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func __obf_fcb9eb75f0392749(__obf_b13d16274eb954e2 *image.YCbCr) *__obf_786a8b420048cd65 {
	__obf_5358e7a9db19986b, __obf_256b49a4ea358eeb := __obf_b13d16274eb954e2.Rect.Dx(), __obf_b13d16274eb954e2.Rect.Dy()
	__obf_1c53ad4397d4b9ae := __obf_786a8b420048cd65{
		Pix:            make([]uint8, 3*__obf_5358e7a9db19986b*__obf_256b49a4ea358eeb),
		Stride:         3 * __obf_5358e7a9db19986b,
		Rect:           image.Rect(0, 0, __obf_5358e7a9db19986b, __obf_256b49a4ea358eeb),
		SubsampleRatio: __obf_b13d16274eb954e2.SubsampleRatio,
	}
	switch __obf_b13d16274eb954e2.SubsampleRatio {
	case __obf_c331ec71612cf606:
		return __obf_0e7d4b4fb799af25(__obf_b13d16274eb954e2, &__obf_1c53ad4397d4b9ae)
	case __obf_fd1adb463076babb:
		return __obf_d30d7b6fce66c77a(__obf_b13d16274eb954e2, &__obf_1c53ad4397d4b9ae)
	case __obf_6efc98209b79485d:
		return __obf_178a094cb5780e34(__obf_b13d16274eb954e2, &__obf_1c53ad4397d4b9ae)
	case __obf_02ac68fc9cbc8441:
		return __obf_b5ea799a718ce045(__obf_b13d16274eb954e2, &__obf_1c53ad4397d4b9ae)
	case __obf_ec8fa2b0adacb378:
		return __obf_27cdd7ee674816d4(__obf_b13d16274eb954e2, &__obf_1c53ad4397d4b9ae)
	case __obf_e4ffc790bf709051:
		return __obf_87bd739f45afea1f(__obf_b13d16274eb954e2, &__obf_1c53ad4397d4b9ae)
	}
	return &__obf_1c53ad4397d4b9ae
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) __obf_860222aeb0ce640b(__obf_fcc28d0eb472022e *image.YCbCr) *image.YCbCr {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_fcc28d0eb472022e.Y
	Cb := __obf_fcc28d0eb472022e.Cb
	Cr := __obf_fcc28d0eb472022e.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_fcc28d0eb472022e.Rect.Max.Y-__obf_fcc28d0eb472022e.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.YStride
		__obf_ef36a6f27c0b18dd := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_fcc28d0eb472022e.Rect.Max.X-__obf_fcc28d0eb472022e.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/2
			Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a] = Pix[__obf_7ec798c4884b80db+0]
			Cb[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+1]
			Cr[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+2]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_fcc28d0eb472022e
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) __obf_bda6a895907a47e6(__obf_fcc28d0eb472022e *image.YCbCr) *image.YCbCr {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_fcc28d0eb472022e.Y
	Cb := __obf_fcc28d0eb472022e.Cb
	Cr := __obf_fcc28d0eb472022e.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_fcc28d0eb472022e.Rect.Max.Y-__obf_fcc28d0eb472022e.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.YStride
		__obf_ef36a6f27c0b18dd := (__obf_eed2203158cb878d / 2) * __obf_fcc28d0eb472022e.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_fcc28d0eb472022e.Rect.Max.X-__obf_fcc28d0eb472022e.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/2
			Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a] = Pix[__obf_7ec798c4884b80db+0]
			Cb[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+1]
			Cr[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+2]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_fcc28d0eb472022e
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) __obf_31cdd1f60e4e7e22(__obf_fcc28d0eb472022e *image.YCbCr) *image.YCbCr {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_fcc28d0eb472022e.Y
	Cb := __obf_fcc28d0eb472022e.Cb
	Cr := __obf_fcc28d0eb472022e.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_fcc28d0eb472022e.Rect.Max.Y-__obf_fcc28d0eb472022e.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.YStride
		__obf_ef36a6f27c0b18dd := (__obf_eed2203158cb878d / 2) * __obf_fcc28d0eb472022e.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_fcc28d0eb472022e.Rect.Max.X-__obf_fcc28d0eb472022e.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a
			Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a] = Pix[__obf_7ec798c4884b80db+0]
			Cb[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+1]
			Cr[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+2]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_fcc28d0eb472022e
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) __obf_0f2b33a269f54728(__obf_fcc28d0eb472022e *image.YCbCr) *image.YCbCr {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_fcc28d0eb472022e.Y
	Cb := __obf_fcc28d0eb472022e.Cb
	Cr := __obf_fcc28d0eb472022e.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_fcc28d0eb472022e.Rect.Max.Y-__obf_fcc28d0eb472022e.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.YStride
		__obf_ef36a6f27c0b18dd := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_fcc28d0eb472022e.Rect.Max.X-__obf_fcc28d0eb472022e.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a
			Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a] = Pix[__obf_7ec798c4884b80db+0]
			Cb[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+1]
			Cr[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+2]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_fcc28d0eb472022e
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) __obf_2967d8f3c994be5e(__obf_fcc28d0eb472022e *image.YCbCr) *image.YCbCr {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_fcc28d0eb472022e.Y
	Cb := __obf_fcc28d0eb472022e.Cb
	Cr := __obf_fcc28d0eb472022e.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_fcc28d0eb472022e.Rect.Max.Y-__obf_fcc28d0eb472022e.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.YStride
		__obf_ef36a6f27c0b18dd := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_fcc28d0eb472022e.Rect.Max.X-__obf_fcc28d0eb472022e.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/4
			Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a] = Pix[__obf_7ec798c4884b80db+0]
			Cb[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+1]
			Cr[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+2]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_fcc28d0eb472022e
}

func (__obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) __obf_08210d27d9b963ce(__obf_fcc28d0eb472022e *image.YCbCr) *image.YCbCr {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_fcc28d0eb472022e.Y
	Cb := __obf_fcc28d0eb472022e.Cb
	Cr := __obf_fcc28d0eb472022e.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_fcc28d0eb472022e.Rect.Max.Y-__obf_fcc28d0eb472022e.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_fcc28d0eb472022e.YStride
		__obf_ef36a6f27c0b18dd := (__obf_eed2203158cb878d / 2) * __obf_fcc28d0eb472022e.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_fcc28d0eb472022e.Rect.Max.X-__obf_fcc28d0eb472022e.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/4
			Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a] = Pix[__obf_7ec798c4884b80db+0]
			Cb[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+1]
			Cr[__obf_1986f534cb8126b3] = Pix[__obf_7ec798c4884b80db+2]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_fcc28d0eb472022e
}

func __obf_0e7d4b4fb799af25(__obf_b13d16274eb954e2 *image.YCbCr, __obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) *__obf_786a8b420048cd65 {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_b13d16274eb954e2.Y
	Cb := __obf_b13d16274eb954e2.Cb
	Cr := __obf_b13d16274eb954e2.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_b13d16274eb954e2.Rect.Max.Y-__obf_b13d16274eb954e2.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.YStride
		__obf_ef36a6f27c0b18dd := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_b13d16274eb954e2.Rect.Max.X-__obf_b13d16274eb954e2.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/2
			Pix[__obf_7ec798c4884b80db+0] = Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a]
			Pix[__obf_7ec798c4884b80db+1] = Cb[__obf_1986f534cb8126b3]
			Pix[__obf_7ec798c4884b80db+2] = Cr[__obf_1986f534cb8126b3]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_1c53ad4397d4b9ae
}

func __obf_d30d7b6fce66c77a(__obf_b13d16274eb954e2 *image.YCbCr, __obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) *__obf_786a8b420048cd65 {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_b13d16274eb954e2.Y
	Cb := __obf_b13d16274eb954e2.Cb
	Cr := __obf_b13d16274eb954e2.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_b13d16274eb954e2.Rect.Max.Y-__obf_b13d16274eb954e2.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.YStride
		__obf_ef36a6f27c0b18dd := (__obf_eed2203158cb878d / 2) * __obf_b13d16274eb954e2.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_b13d16274eb954e2.Rect.Max.X-__obf_b13d16274eb954e2.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/2
			Pix[__obf_7ec798c4884b80db+0] = Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a]
			Pix[__obf_7ec798c4884b80db+1] = Cb[__obf_1986f534cb8126b3]
			Pix[__obf_7ec798c4884b80db+2] = Cr[__obf_1986f534cb8126b3]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_1c53ad4397d4b9ae
}

func __obf_178a094cb5780e34(__obf_b13d16274eb954e2 *image.YCbCr, __obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) *__obf_786a8b420048cd65 {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_b13d16274eb954e2.Y
	Cb := __obf_b13d16274eb954e2.Cb
	Cr := __obf_b13d16274eb954e2.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_b13d16274eb954e2.Rect.Max.Y-__obf_b13d16274eb954e2.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.YStride
		__obf_ef36a6f27c0b18dd := (__obf_eed2203158cb878d / 2) * __obf_b13d16274eb954e2.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_b13d16274eb954e2.Rect.Max.X-__obf_b13d16274eb954e2.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a
			Pix[__obf_7ec798c4884b80db+0] = Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a]
			Pix[__obf_7ec798c4884b80db+1] = Cb[__obf_1986f534cb8126b3]
			Pix[__obf_7ec798c4884b80db+2] = Cr[__obf_1986f534cb8126b3]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_1c53ad4397d4b9ae
}

func __obf_b5ea799a718ce045(__obf_b13d16274eb954e2 *image.YCbCr, __obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) *__obf_786a8b420048cd65 {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_b13d16274eb954e2.Y
	Cb := __obf_b13d16274eb954e2.Cb
	Cr := __obf_b13d16274eb954e2.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_b13d16274eb954e2.Rect.Max.Y-__obf_b13d16274eb954e2.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.YStride
		__obf_ef36a6f27c0b18dd := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_b13d16274eb954e2.Rect.Max.X-__obf_b13d16274eb954e2.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a
			Pix[__obf_7ec798c4884b80db+0] = Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a]
			Pix[__obf_7ec798c4884b80db+1] = Cb[__obf_1986f534cb8126b3]
			Pix[__obf_7ec798c4884b80db+2] = Cr[__obf_1986f534cb8126b3]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_1c53ad4397d4b9ae
}

func __obf_27cdd7ee674816d4(__obf_b13d16274eb954e2 *image.YCbCr, __obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) *__obf_786a8b420048cd65 {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_b13d16274eb954e2.Y
	Cb := __obf_b13d16274eb954e2.Cb
	Cr := __obf_b13d16274eb954e2.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_b13d16274eb954e2.Rect.Max.Y-__obf_b13d16274eb954e2.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.YStride
		__obf_ef36a6f27c0b18dd := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_b13d16274eb954e2.Rect.Max.X-__obf_b13d16274eb954e2.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/4
			Pix[__obf_7ec798c4884b80db+0] = Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a]
			Pix[__obf_7ec798c4884b80db+1] = Cb[__obf_1986f534cb8126b3]
			Pix[__obf_7ec798c4884b80db+2] = Cr[__obf_1986f534cb8126b3]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_1c53ad4397d4b9ae
}

func __obf_87bd739f45afea1f(__obf_b13d16274eb954e2 *image.YCbCr, __obf_1c53ad4397d4b9ae *__obf_786a8b420048cd65) *__obf_786a8b420048cd65 {
	var __obf_7ec798c4884b80db int
	Pix := __obf_1c53ad4397d4b9ae.Pix
	Y := __obf_b13d16274eb954e2.Y
	Cb := __obf_b13d16274eb954e2.Cb
	Cr := __obf_b13d16274eb954e2.Cr
	for __obf_eed2203158cb878d := 0; __obf_eed2203158cb878d < __obf_b13d16274eb954e2.Rect.Max.Y-__obf_b13d16274eb954e2.Rect.Min.Y; __obf_eed2203158cb878d++ {
		__obf_55d093c1aa60a23c := __obf_eed2203158cb878d * __obf_b13d16274eb954e2.YStride
		__obf_ef36a6f27c0b18dd := (__obf_eed2203158cb878d / 2) * __obf_b13d16274eb954e2.CStride
		for __obf_c6d53ec5b4505f6a := 0; __obf_c6d53ec5b4505f6a < __obf_b13d16274eb954e2.Rect.Max.X-__obf_b13d16274eb954e2.Rect.Min.X; __obf_c6d53ec5b4505f6a++ {
			__obf_1986f534cb8126b3 := __obf_ef36a6f27c0b18dd + __obf_c6d53ec5b4505f6a/4
			Pix[__obf_7ec798c4884b80db+0] = Y[__obf_55d093c1aa60a23c+__obf_c6d53ec5b4505f6a]
			Pix[__obf_7ec798c4884b80db+1] = Cb[__obf_1986f534cb8126b3]
			Pix[__obf_7ec798c4884b80db+2] = Cr[__obf_1986f534cb8126b3]
			__obf_7ec798c4884b80db += 3
		}
	}
	return __obf_1c53ad4397d4b9ae
}
