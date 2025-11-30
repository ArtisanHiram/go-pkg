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

package __obf_42b2ccbdafaee9c3

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type __obf_e4d16584f024f650 struct {
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
func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) PixOffset(__obf_9ab3da411fe8abc3, __obf_b876c8539198a4e8 int) int {
	return (__obf_b876c8539198a4e8-__obf_8e9fad7c873cabdb.Rect.Min.Y)*__obf_8e9fad7c873cabdb.Stride + (__obf_9ab3da411fe8abc3-__obf_8e9fad7c873cabdb.Rect.Min.X)*3
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) Bounds() image.Rectangle {
	return __obf_8e9fad7c873cabdb.Rect
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) ColorModel() color.Model {
	return color.YCbCrModel
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) At(__obf_9ab3da411fe8abc3, __obf_b876c8539198a4e8 int) color.Color {
	if !(image.Point{__obf_9ab3da411fe8abc3, __obf_b876c8539198a4e8}.In(__obf_8e9fad7c873cabdb.Rect)) {
		return color.YCbCr{}
	}
	__obf_e39e622bfcb683d9 := __obf_8e9fad7c873cabdb.PixOffset(__obf_9ab3da411fe8abc3, __obf_b876c8539198a4e8)
	return color.YCbCr{__obf_8e9fad7c873cabdb.
		Pix[__obf_e39e622bfcb683d9+0], __obf_8e9fad7c873cabdb.
		Pix[__obf_e39e622bfcb683d9+1], __obf_8e9fad7c873cabdb.
		Pix[__obf_e39e622bfcb683d9+2],
	}
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) SubImage(__obf_e02b66492c64e775 image.Rectangle) image.Image {
	__obf_e02b66492c64e775 = __obf_e02b66492c64e775.Intersect(__obf_8e9fad7c873cabdb.Rect)
	if __obf_e02b66492c64e775.Empty() {
		return &__obf_e4d16584f024f650{SubsampleRatio: __obf_8e9fad7c873cabdb.SubsampleRatio}
	}
	__obf_e39e622bfcb683d9 := __obf_8e9fad7c873cabdb.PixOffset(__obf_e02b66492c64e775.Min.X, __obf_e02b66492c64e775.Min.Y)
	return &__obf_e4d16584f024f650{
		Pix:            __obf_8e9fad7c873cabdb.Pix[__obf_e39e622bfcb683d9:],
		Stride:         __obf_8e9fad7c873cabdb.Stride,
		Rect:           __obf_e02b66492c64e775,
		SubsampleRatio: __obf_8e9fad7c873cabdb.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func __obf_78d696b65413d950(__obf_e02b66492c64e775 image.Rectangle, __obf_7ea54fb9f5eeddd2 image.YCbCrSubsampleRatio) *__obf_e4d16584f024f650 {
	__obf_4790b99cfc76f48b, __obf_7ba9b70e0dc09474 := __obf_e02b66492c64e775.Dx(), __obf_e02b66492c64e775.Dy()
	__obf_6b92d447370720f4 := make([]uint8, 3*__obf_4790b99cfc76f48b*__obf_7ba9b70e0dc09474)
	return &__obf_e4d16584f024f650{Pix: __obf_6b92d447370720f4, Stride: 3 * __obf_4790b99cfc76f48b, Rect: __obf_e02b66492c64e775, SubsampleRatio: __obf_7ea54fb9f5eeddd2}
}

// Copy of image.YCbCrSubsampleRatio constants - this allows us to support
// older versions of Go where these constants are not defined (i.e. Go 1.4)
const (
	__obf_af25db7bf08054cf image.YCbCrSubsampleRatio = iota
	__obf_6b3c6fbec8c856c9
	__obf_c2cc1315469ede2a
	__obf_ff94fc0d6d363b01
	__obf_acef428cdbb60271
	__obf_a9c3ef08c18bc4fe
)

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) YCbCr() *image.YCbCr {
	__obf_72f30faa2f2e254a := image.NewYCbCr(__obf_8e9fad7c873cabdb.Rect, __obf_8e9fad7c873cabdb.SubsampleRatio)
	switch __obf_72f30faa2f2e254a.SubsampleRatio {
	case __obf_6b3c6fbec8c856c9:
		return __obf_8e9fad7c873cabdb.__obf_fe7bc245ba623e4a(__obf_72f30faa2f2e254a)
	case __obf_c2cc1315469ede2a:
		return __obf_8e9fad7c873cabdb.__obf_811761512a5e5bad(__obf_72f30faa2f2e254a)
	case __obf_ff94fc0d6d363b01:
		return __obf_8e9fad7c873cabdb.__obf_bbe2d08c246297fb(__obf_72f30faa2f2e254a)
	case __obf_af25db7bf08054cf:
		return __obf_8e9fad7c873cabdb.__obf_5614cab7dafe9397(__obf_72f30faa2f2e254a)
	case __obf_acef428cdbb60271:
		return __obf_8e9fad7c873cabdb.__obf_a3f151fdb2be368a(__obf_72f30faa2f2e254a)
	case __obf_a9c3ef08c18bc4fe:
		return __obf_8e9fad7c873cabdb.__obf_e12b60260dd60890(__obf_72f30faa2f2e254a)
	}
	return __obf_72f30faa2f2e254a
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func __obf_79e832e234bf2561(__obf_c82b366489672bd7 *image.YCbCr) *__obf_e4d16584f024f650 {
	__obf_4790b99cfc76f48b, __obf_7ba9b70e0dc09474 := __obf_c82b366489672bd7.Rect.Dx(), __obf_c82b366489672bd7.Rect.Dy()
	__obf_8e9fad7c873cabdb := __obf_e4d16584f024f650{
		Pix:            make([]uint8, 3*__obf_4790b99cfc76f48b*__obf_7ba9b70e0dc09474),
		Stride:         3 * __obf_4790b99cfc76f48b,
		Rect:           image.Rect(0, 0, __obf_4790b99cfc76f48b, __obf_7ba9b70e0dc09474),
		SubsampleRatio: __obf_c82b366489672bd7.SubsampleRatio,
	}
	switch __obf_c82b366489672bd7.SubsampleRatio {
	case __obf_6b3c6fbec8c856c9:
		return __obf_86132b669b39c04c(__obf_c82b366489672bd7, &__obf_8e9fad7c873cabdb)
	case __obf_c2cc1315469ede2a:
		return __obf_1c9a03d1327a1ec2(__obf_c82b366489672bd7, &__obf_8e9fad7c873cabdb)
	case __obf_ff94fc0d6d363b01:
		return __obf_6db73c219703cb6d(__obf_c82b366489672bd7, &__obf_8e9fad7c873cabdb)
	case __obf_af25db7bf08054cf:
		return __obf_635c3add97b843f6(__obf_c82b366489672bd7, &__obf_8e9fad7c873cabdb)
	case __obf_acef428cdbb60271:
		return __obf_0fa58eb035eaf649(__obf_c82b366489672bd7, &__obf_8e9fad7c873cabdb)
	case __obf_a9c3ef08c18bc4fe:
		return __obf_aa61694b20d105aa(__obf_c82b366489672bd7, &__obf_8e9fad7c873cabdb)
	}
	return &__obf_8e9fad7c873cabdb
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) __obf_fe7bc245ba623e4a(__obf_72f30faa2f2e254a *image.YCbCr) *image.YCbCr {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_72f30faa2f2e254a.Y
	Cb := __obf_72f30faa2f2e254a.Cb
	Cr := __obf_72f30faa2f2e254a.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_72f30faa2f2e254a.Rect.Max.Y-__obf_72f30faa2f2e254a.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.YStride
		__obf_c3eb2babe638e016 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_72f30faa2f2e254a.Rect.Max.X-__obf_72f30faa2f2e254a.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/2
			Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3] = Pix[__obf_e3767e7171bac185+0]
			Cb[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+1]
			Cr[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+2]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_72f30faa2f2e254a
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) __obf_811761512a5e5bad(__obf_72f30faa2f2e254a *image.YCbCr) *image.YCbCr {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_72f30faa2f2e254a.Y
	Cb := __obf_72f30faa2f2e254a.Cb
	Cr := __obf_72f30faa2f2e254a.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_72f30faa2f2e254a.Rect.Max.Y-__obf_72f30faa2f2e254a.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.YStride
		__obf_c3eb2babe638e016 := (__obf_b876c8539198a4e8 / 2) * __obf_72f30faa2f2e254a.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_72f30faa2f2e254a.Rect.Max.X-__obf_72f30faa2f2e254a.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/2
			Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3] = Pix[__obf_e3767e7171bac185+0]
			Cb[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+1]
			Cr[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+2]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_72f30faa2f2e254a
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) __obf_bbe2d08c246297fb(__obf_72f30faa2f2e254a *image.YCbCr) *image.YCbCr {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_72f30faa2f2e254a.Y
	Cb := __obf_72f30faa2f2e254a.Cb
	Cr := __obf_72f30faa2f2e254a.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_72f30faa2f2e254a.Rect.Max.Y-__obf_72f30faa2f2e254a.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.YStride
		__obf_c3eb2babe638e016 := (__obf_b876c8539198a4e8 / 2) * __obf_72f30faa2f2e254a.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_72f30faa2f2e254a.Rect.Max.X-__obf_72f30faa2f2e254a.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3
			Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3] = Pix[__obf_e3767e7171bac185+0]
			Cb[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+1]
			Cr[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+2]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_72f30faa2f2e254a
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) __obf_5614cab7dafe9397(__obf_72f30faa2f2e254a *image.YCbCr) *image.YCbCr {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_72f30faa2f2e254a.Y
	Cb := __obf_72f30faa2f2e254a.Cb
	Cr := __obf_72f30faa2f2e254a.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_72f30faa2f2e254a.Rect.Max.Y-__obf_72f30faa2f2e254a.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.YStride
		__obf_c3eb2babe638e016 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_72f30faa2f2e254a.Rect.Max.X-__obf_72f30faa2f2e254a.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3
			Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3] = Pix[__obf_e3767e7171bac185+0]
			Cb[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+1]
			Cr[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+2]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_72f30faa2f2e254a
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) __obf_a3f151fdb2be368a(__obf_72f30faa2f2e254a *image.YCbCr) *image.YCbCr {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_72f30faa2f2e254a.Y
	Cb := __obf_72f30faa2f2e254a.Cb
	Cr := __obf_72f30faa2f2e254a.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_72f30faa2f2e254a.Rect.Max.Y-__obf_72f30faa2f2e254a.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.YStride
		__obf_c3eb2babe638e016 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_72f30faa2f2e254a.Rect.Max.X-__obf_72f30faa2f2e254a.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/4
			Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3] = Pix[__obf_e3767e7171bac185+0]
			Cb[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+1]
			Cr[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+2]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_72f30faa2f2e254a
}

func (__obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) __obf_e12b60260dd60890(__obf_72f30faa2f2e254a *image.YCbCr) *image.YCbCr {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_72f30faa2f2e254a.Y
	Cb := __obf_72f30faa2f2e254a.Cb
	Cr := __obf_72f30faa2f2e254a.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_72f30faa2f2e254a.Rect.Max.Y-__obf_72f30faa2f2e254a.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_72f30faa2f2e254a.YStride
		__obf_c3eb2babe638e016 := (__obf_b876c8539198a4e8 / 2) * __obf_72f30faa2f2e254a.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_72f30faa2f2e254a.Rect.Max.X-__obf_72f30faa2f2e254a.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/4
			Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3] = Pix[__obf_e3767e7171bac185+0]
			Cb[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+1]
			Cr[__obf_acde913b01c6d744] = Pix[__obf_e3767e7171bac185+2]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_72f30faa2f2e254a
}

func __obf_86132b669b39c04c(__obf_c82b366489672bd7 *image.YCbCr, __obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) *__obf_e4d16584f024f650 {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_c82b366489672bd7.Y
	Cb := __obf_c82b366489672bd7.Cb
	Cr := __obf_c82b366489672bd7.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_c82b366489672bd7.Rect.Max.Y-__obf_c82b366489672bd7.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.YStride
		__obf_c3eb2babe638e016 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_c82b366489672bd7.Rect.Max.X-__obf_c82b366489672bd7.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/2
			Pix[__obf_e3767e7171bac185+0] = Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3]
			Pix[__obf_e3767e7171bac185+1] = Cb[__obf_acde913b01c6d744]
			Pix[__obf_e3767e7171bac185+2] = Cr[__obf_acde913b01c6d744]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_8e9fad7c873cabdb
}

func __obf_1c9a03d1327a1ec2(__obf_c82b366489672bd7 *image.YCbCr, __obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) *__obf_e4d16584f024f650 {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_c82b366489672bd7.Y
	Cb := __obf_c82b366489672bd7.Cb
	Cr := __obf_c82b366489672bd7.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_c82b366489672bd7.Rect.Max.Y-__obf_c82b366489672bd7.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.YStride
		__obf_c3eb2babe638e016 := (__obf_b876c8539198a4e8 / 2) * __obf_c82b366489672bd7.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_c82b366489672bd7.Rect.Max.X-__obf_c82b366489672bd7.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/2
			Pix[__obf_e3767e7171bac185+0] = Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3]
			Pix[__obf_e3767e7171bac185+1] = Cb[__obf_acde913b01c6d744]
			Pix[__obf_e3767e7171bac185+2] = Cr[__obf_acde913b01c6d744]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_8e9fad7c873cabdb
}

func __obf_6db73c219703cb6d(__obf_c82b366489672bd7 *image.YCbCr, __obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) *__obf_e4d16584f024f650 {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_c82b366489672bd7.Y
	Cb := __obf_c82b366489672bd7.Cb
	Cr := __obf_c82b366489672bd7.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_c82b366489672bd7.Rect.Max.Y-__obf_c82b366489672bd7.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.YStride
		__obf_c3eb2babe638e016 := (__obf_b876c8539198a4e8 / 2) * __obf_c82b366489672bd7.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_c82b366489672bd7.Rect.Max.X-__obf_c82b366489672bd7.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3
			Pix[__obf_e3767e7171bac185+0] = Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3]
			Pix[__obf_e3767e7171bac185+1] = Cb[__obf_acde913b01c6d744]
			Pix[__obf_e3767e7171bac185+2] = Cr[__obf_acde913b01c6d744]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_8e9fad7c873cabdb
}

func __obf_635c3add97b843f6(__obf_c82b366489672bd7 *image.YCbCr, __obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) *__obf_e4d16584f024f650 {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_c82b366489672bd7.Y
	Cb := __obf_c82b366489672bd7.Cb
	Cr := __obf_c82b366489672bd7.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_c82b366489672bd7.Rect.Max.Y-__obf_c82b366489672bd7.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.YStride
		__obf_c3eb2babe638e016 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_c82b366489672bd7.Rect.Max.X-__obf_c82b366489672bd7.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3
			Pix[__obf_e3767e7171bac185+0] = Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3]
			Pix[__obf_e3767e7171bac185+1] = Cb[__obf_acde913b01c6d744]
			Pix[__obf_e3767e7171bac185+2] = Cr[__obf_acde913b01c6d744]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_8e9fad7c873cabdb
}

func __obf_0fa58eb035eaf649(__obf_c82b366489672bd7 *image.YCbCr, __obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) *__obf_e4d16584f024f650 {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_c82b366489672bd7.Y
	Cb := __obf_c82b366489672bd7.Cb
	Cr := __obf_c82b366489672bd7.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_c82b366489672bd7.Rect.Max.Y-__obf_c82b366489672bd7.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.YStride
		__obf_c3eb2babe638e016 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_c82b366489672bd7.Rect.Max.X-__obf_c82b366489672bd7.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/4
			Pix[__obf_e3767e7171bac185+0] = Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3]
			Pix[__obf_e3767e7171bac185+1] = Cb[__obf_acde913b01c6d744]
			Pix[__obf_e3767e7171bac185+2] = Cr[__obf_acde913b01c6d744]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_8e9fad7c873cabdb
}

func __obf_aa61694b20d105aa(__obf_c82b366489672bd7 *image.YCbCr, __obf_8e9fad7c873cabdb *__obf_e4d16584f024f650) *__obf_e4d16584f024f650 {
	var __obf_e3767e7171bac185 int
	Pix := __obf_8e9fad7c873cabdb.Pix
	Y := __obf_c82b366489672bd7.Y
	Cb := __obf_c82b366489672bd7.Cb
	Cr := __obf_c82b366489672bd7.Cr
	for __obf_b876c8539198a4e8 := 0; __obf_b876c8539198a4e8 < __obf_c82b366489672bd7.Rect.Max.Y-__obf_c82b366489672bd7.Rect.Min.Y; __obf_b876c8539198a4e8++ {
		__obf_510a349327618f91 := __obf_b876c8539198a4e8 * __obf_c82b366489672bd7.YStride
		__obf_c3eb2babe638e016 := (__obf_b876c8539198a4e8 / 2) * __obf_c82b366489672bd7.CStride
		for __obf_9ab3da411fe8abc3 := 0; __obf_9ab3da411fe8abc3 < __obf_c82b366489672bd7.Rect.Max.X-__obf_c82b366489672bd7.Rect.Min.X; __obf_9ab3da411fe8abc3++ {
			__obf_acde913b01c6d744 := __obf_c3eb2babe638e016 + __obf_9ab3da411fe8abc3/4
			Pix[__obf_e3767e7171bac185+0] = Y[__obf_510a349327618f91+__obf_9ab3da411fe8abc3]
			Pix[__obf_e3767e7171bac185+1] = Cb[__obf_acde913b01c6d744]
			Pix[__obf_e3767e7171bac185+2] = Cr[__obf_acde913b01c6d744]
			__obf_e3767e7171bac185 += 3
		}
	}
	return __obf_8e9fad7c873cabdb
}
