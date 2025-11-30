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

package __obf_ac510735d4d13cdd

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type __obf_9474021cedba6d1a struct {
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
func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) PixOffset(__obf_21a3146223eb514f, __obf_4811266080d7795e int) int {
	return (__obf_4811266080d7795e-__obf_5fd02db579f190b4.Rect.Min.Y)*__obf_5fd02db579f190b4.Stride + (__obf_21a3146223eb514f-__obf_5fd02db579f190b4.Rect.Min.X)*3
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) Bounds() image.Rectangle {
	return __obf_5fd02db579f190b4.Rect
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) ColorModel() color.Model {
	return color.YCbCrModel
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) At(__obf_21a3146223eb514f, __obf_4811266080d7795e int) color.Color {
	if !(image.Point{__obf_21a3146223eb514f, __obf_4811266080d7795e}.In(__obf_5fd02db579f190b4.Rect)) {
		return color.YCbCr{}
	}
	__obf_7326873273fab46e := __obf_5fd02db579f190b4.PixOffset(__obf_21a3146223eb514f, __obf_4811266080d7795e)
	return color.YCbCr{__obf_5fd02db579f190b4.
		Pix[__obf_7326873273fab46e+0], __obf_5fd02db579f190b4.
		Pix[__obf_7326873273fab46e+1], __obf_5fd02db579f190b4.
		Pix[__obf_7326873273fab46e+2],
	}
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) SubImage(__obf_b4b27f4d455c3be9 image.Rectangle) image.Image {
	__obf_b4b27f4d455c3be9 = __obf_b4b27f4d455c3be9.Intersect(__obf_5fd02db579f190b4.Rect)
	if __obf_b4b27f4d455c3be9.Empty() {
		return &__obf_9474021cedba6d1a{SubsampleRatio: __obf_5fd02db579f190b4.SubsampleRatio}
	}
	__obf_7326873273fab46e := __obf_5fd02db579f190b4.PixOffset(__obf_b4b27f4d455c3be9.Min.X, __obf_b4b27f4d455c3be9.Min.Y)
	return &__obf_9474021cedba6d1a{
		Pix:            __obf_5fd02db579f190b4.Pix[__obf_7326873273fab46e:],
		Stride:         __obf_5fd02db579f190b4.Stride,
		Rect:           __obf_b4b27f4d455c3be9,
		SubsampleRatio: __obf_5fd02db579f190b4.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func __obf_57395a4603262f46(__obf_b4b27f4d455c3be9 image.Rectangle, __obf_8a6fed9748779416 image.YCbCrSubsampleRatio) *__obf_9474021cedba6d1a {
	__obf_9c9ff2b43e3e5453, __obf_a7e77eb8cb2d21b0 := __obf_b4b27f4d455c3be9.Dx(), __obf_b4b27f4d455c3be9.Dy()
	__obf_7f6707955213aad2 := make([]uint8, 3*__obf_9c9ff2b43e3e5453*__obf_a7e77eb8cb2d21b0)
	return &__obf_9474021cedba6d1a{Pix: __obf_7f6707955213aad2, Stride: 3 * __obf_9c9ff2b43e3e5453, Rect: __obf_b4b27f4d455c3be9, SubsampleRatio: __obf_8a6fed9748779416}
}

// Copy of image.YCbCrSubsampleRatio constants - this allows us to support
// older versions of Go where these constants are not defined (i.e. Go 1.4)
const (
	__obf_6f63bcdacd885acf image.YCbCrSubsampleRatio = iota
	__obf_63d58e2f12921d95
	__obf_80b1638b17a4367f
	__obf_b26af6c8edadbd66
	__obf_a064de5c95154858
	__obf_32cbe6a343e5a18d
)

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) YCbCr() *image.YCbCr {
	__obf_a05823a28621664e := image.NewYCbCr(__obf_5fd02db579f190b4.Rect, __obf_5fd02db579f190b4.SubsampleRatio)
	switch __obf_a05823a28621664e.SubsampleRatio {
	case __obf_63d58e2f12921d95:
		return __obf_5fd02db579f190b4.__obf_1f48a964f1c7e61b(__obf_a05823a28621664e)
	case __obf_80b1638b17a4367f:
		return __obf_5fd02db579f190b4.__obf_06121f84d52b569d(__obf_a05823a28621664e)
	case __obf_b26af6c8edadbd66:
		return __obf_5fd02db579f190b4.__obf_761da9a7b862baa5(__obf_a05823a28621664e)
	case __obf_6f63bcdacd885acf:
		return __obf_5fd02db579f190b4.__obf_ccc0969b5347d9e6(__obf_a05823a28621664e)
	case __obf_a064de5c95154858:
		return __obf_5fd02db579f190b4.__obf_77b13e0cc5fb1d3c(__obf_a05823a28621664e)
	case __obf_32cbe6a343e5a18d:
		return __obf_5fd02db579f190b4.__obf_ade402fcb4f250cc(__obf_a05823a28621664e)
	}
	return __obf_a05823a28621664e
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func __obf_76286d70dcc66ea1(__obf_3ccd19e51488be10 *image.YCbCr) *__obf_9474021cedba6d1a {
	__obf_9c9ff2b43e3e5453, __obf_a7e77eb8cb2d21b0 := __obf_3ccd19e51488be10.Rect.Dx(), __obf_3ccd19e51488be10.Rect.Dy()
	__obf_5fd02db579f190b4 := __obf_9474021cedba6d1a{
		Pix:            make([]uint8, 3*__obf_9c9ff2b43e3e5453*__obf_a7e77eb8cb2d21b0),
		Stride:         3 * __obf_9c9ff2b43e3e5453,
		Rect:           image.Rect(0, 0, __obf_9c9ff2b43e3e5453, __obf_a7e77eb8cb2d21b0),
		SubsampleRatio: __obf_3ccd19e51488be10.SubsampleRatio,
	}
	switch __obf_3ccd19e51488be10.SubsampleRatio {
	case __obf_63d58e2f12921d95:
		return __obf_050e2d825f58f6b4(__obf_3ccd19e51488be10, &__obf_5fd02db579f190b4)
	case __obf_80b1638b17a4367f:
		return __obf_5cd768069157852e(__obf_3ccd19e51488be10, &__obf_5fd02db579f190b4)
	case __obf_b26af6c8edadbd66:
		return __obf_b2d8e081f70af140(__obf_3ccd19e51488be10, &__obf_5fd02db579f190b4)
	case __obf_6f63bcdacd885acf:
		return __obf_2adb0deba9d6ec31(__obf_3ccd19e51488be10, &__obf_5fd02db579f190b4)
	case __obf_a064de5c95154858:
		return __obf_53760723335bb13c(__obf_3ccd19e51488be10, &__obf_5fd02db579f190b4)
	case __obf_32cbe6a343e5a18d:
		return __obf_1ea11da1a27ac5db(__obf_3ccd19e51488be10, &__obf_5fd02db579f190b4)
	}
	return &__obf_5fd02db579f190b4
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) __obf_1f48a964f1c7e61b(__obf_a05823a28621664e *image.YCbCr) *image.YCbCr {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_a05823a28621664e.Y
	Cb := __obf_a05823a28621664e.Cb
	Cr := __obf_a05823a28621664e.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_a05823a28621664e.Rect.Max.Y-__obf_a05823a28621664e.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_a05823a28621664e.YStride
		__obf_5ad115631709a2a3 := __obf_4811266080d7795e * __obf_a05823a28621664e.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_a05823a28621664e.Rect.Max.X-__obf_a05823a28621664e.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/2
			Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f] = Pix[__obf_dfa50e8a4ab3e6d5+0]
			Cb[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+1]
			Cr[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+2]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_a05823a28621664e
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) __obf_06121f84d52b569d(__obf_a05823a28621664e *image.YCbCr) *image.YCbCr {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_a05823a28621664e.Y
	Cb := __obf_a05823a28621664e.Cb
	Cr := __obf_a05823a28621664e.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_a05823a28621664e.Rect.Max.Y-__obf_a05823a28621664e.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_a05823a28621664e.YStride
		__obf_5ad115631709a2a3 := (__obf_4811266080d7795e / 2) * __obf_a05823a28621664e.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_a05823a28621664e.Rect.Max.X-__obf_a05823a28621664e.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/2
			Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f] = Pix[__obf_dfa50e8a4ab3e6d5+0]
			Cb[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+1]
			Cr[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+2]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_a05823a28621664e
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) __obf_761da9a7b862baa5(__obf_a05823a28621664e *image.YCbCr) *image.YCbCr {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_a05823a28621664e.Y
	Cb := __obf_a05823a28621664e.Cb
	Cr := __obf_a05823a28621664e.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_a05823a28621664e.Rect.Max.Y-__obf_a05823a28621664e.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_a05823a28621664e.YStride
		__obf_5ad115631709a2a3 := (__obf_4811266080d7795e / 2) * __obf_a05823a28621664e.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_a05823a28621664e.Rect.Max.X-__obf_a05823a28621664e.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f
			Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f] = Pix[__obf_dfa50e8a4ab3e6d5+0]
			Cb[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+1]
			Cr[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+2]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_a05823a28621664e
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) __obf_ccc0969b5347d9e6(__obf_a05823a28621664e *image.YCbCr) *image.YCbCr {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_a05823a28621664e.Y
	Cb := __obf_a05823a28621664e.Cb
	Cr := __obf_a05823a28621664e.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_a05823a28621664e.Rect.Max.Y-__obf_a05823a28621664e.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_a05823a28621664e.YStride
		__obf_5ad115631709a2a3 := __obf_4811266080d7795e * __obf_a05823a28621664e.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_a05823a28621664e.Rect.Max.X-__obf_a05823a28621664e.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f
			Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f] = Pix[__obf_dfa50e8a4ab3e6d5+0]
			Cb[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+1]
			Cr[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+2]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_a05823a28621664e
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) __obf_77b13e0cc5fb1d3c(__obf_a05823a28621664e *image.YCbCr) *image.YCbCr {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_a05823a28621664e.Y
	Cb := __obf_a05823a28621664e.Cb
	Cr := __obf_a05823a28621664e.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_a05823a28621664e.Rect.Max.Y-__obf_a05823a28621664e.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_a05823a28621664e.YStride
		__obf_5ad115631709a2a3 := __obf_4811266080d7795e * __obf_a05823a28621664e.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_a05823a28621664e.Rect.Max.X-__obf_a05823a28621664e.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/4
			Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f] = Pix[__obf_dfa50e8a4ab3e6d5+0]
			Cb[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+1]
			Cr[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+2]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_a05823a28621664e
}

func (__obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) __obf_ade402fcb4f250cc(__obf_a05823a28621664e *image.YCbCr) *image.YCbCr {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_a05823a28621664e.Y
	Cb := __obf_a05823a28621664e.Cb
	Cr := __obf_a05823a28621664e.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_a05823a28621664e.Rect.Max.Y-__obf_a05823a28621664e.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_a05823a28621664e.YStride
		__obf_5ad115631709a2a3 := (__obf_4811266080d7795e / 2) * __obf_a05823a28621664e.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_a05823a28621664e.Rect.Max.X-__obf_a05823a28621664e.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/4
			Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f] = Pix[__obf_dfa50e8a4ab3e6d5+0]
			Cb[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+1]
			Cr[__obf_5a46f1764400ebbd] = Pix[__obf_dfa50e8a4ab3e6d5+2]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_a05823a28621664e
}

func __obf_050e2d825f58f6b4(__obf_3ccd19e51488be10 *image.YCbCr, __obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) *__obf_9474021cedba6d1a {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_3ccd19e51488be10.Y
	Cb := __obf_3ccd19e51488be10.Cb
	Cr := __obf_3ccd19e51488be10.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_3ccd19e51488be10.Rect.Max.Y-__obf_3ccd19e51488be10.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_3ccd19e51488be10.YStride
		__obf_5ad115631709a2a3 := __obf_4811266080d7795e * __obf_3ccd19e51488be10.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_3ccd19e51488be10.Rect.Max.X-__obf_3ccd19e51488be10.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/2
			Pix[__obf_dfa50e8a4ab3e6d5+0] = Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f]
			Pix[__obf_dfa50e8a4ab3e6d5+1] = Cb[__obf_5a46f1764400ebbd]
			Pix[__obf_dfa50e8a4ab3e6d5+2] = Cr[__obf_5a46f1764400ebbd]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_5fd02db579f190b4
}

func __obf_5cd768069157852e(__obf_3ccd19e51488be10 *image.YCbCr, __obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) *__obf_9474021cedba6d1a {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_3ccd19e51488be10.Y
	Cb := __obf_3ccd19e51488be10.Cb
	Cr := __obf_3ccd19e51488be10.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_3ccd19e51488be10.Rect.Max.Y-__obf_3ccd19e51488be10.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_3ccd19e51488be10.YStride
		__obf_5ad115631709a2a3 := (__obf_4811266080d7795e / 2) * __obf_3ccd19e51488be10.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_3ccd19e51488be10.Rect.Max.X-__obf_3ccd19e51488be10.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/2
			Pix[__obf_dfa50e8a4ab3e6d5+0] = Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f]
			Pix[__obf_dfa50e8a4ab3e6d5+1] = Cb[__obf_5a46f1764400ebbd]
			Pix[__obf_dfa50e8a4ab3e6d5+2] = Cr[__obf_5a46f1764400ebbd]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_5fd02db579f190b4
}

func __obf_b2d8e081f70af140(__obf_3ccd19e51488be10 *image.YCbCr, __obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) *__obf_9474021cedba6d1a {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_3ccd19e51488be10.Y
	Cb := __obf_3ccd19e51488be10.Cb
	Cr := __obf_3ccd19e51488be10.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_3ccd19e51488be10.Rect.Max.Y-__obf_3ccd19e51488be10.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_3ccd19e51488be10.YStride
		__obf_5ad115631709a2a3 := (__obf_4811266080d7795e / 2) * __obf_3ccd19e51488be10.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_3ccd19e51488be10.Rect.Max.X-__obf_3ccd19e51488be10.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f
			Pix[__obf_dfa50e8a4ab3e6d5+0] = Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f]
			Pix[__obf_dfa50e8a4ab3e6d5+1] = Cb[__obf_5a46f1764400ebbd]
			Pix[__obf_dfa50e8a4ab3e6d5+2] = Cr[__obf_5a46f1764400ebbd]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_5fd02db579f190b4
}

func __obf_2adb0deba9d6ec31(__obf_3ccd19e51488be10 *image.YCbCr, __obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) *__obf_9474021cedba6d1a {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_3ccd19e51488be10.Y
	Cb := __obf_3ccd19e51488be10.Cb
	Cr := __obf_3ccd19e51488be10.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_3ccd19e51488be10.Rect.Max.Y-__obf_3ccd19e51488be10.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_3ccd19e51488be10.YStride
		__obf_5ad115631709a2a3 := __obf_4811266080d7795e * __obf_3ccd19e51488be10.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_3ccd19e51488be10.Rect.Max.X-__obf_3ccd19e51488be10.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f
			Pix[__obf_dfa50e8a4ab3e6d5+0] = Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f]
			Pix[__obf_dfa50e8a4ab3e6d5+1] = Cb[__obf_5a46f1764400ebbd]
			Pix[__obf_dfa50e8a4ab3e6d5+2] = Cr[__obf_5a46f1764400ebbd]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_5fd02db579f190b4
}

func __obf_53760723335bb13c(__obf_3ccd19e51488be10 *image.YCbCr, __obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) *__obf_9474021cedba6d1a {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_3ccd19e51488be10.Y
	Cb := __obf_3ccd19e51488be10.Cb
	Cr := __obf_3ccd19e51488be10.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_3ccd19e51488be10.Rect.Max.Y-__obf_3ccd19e51488be10.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_3ccd19e51488be10.YStride
		__obf_5ad115631709a2a3 := __obf_4811266080d7795e * __obf_3ccd19e51488be10.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_3ccd19e51488be10.Rect.Max.X-__obf_3ccd19e51488be10.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/4
			Pix[__obf_dfa50e8a4ab3e6d5+0] = Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f]
			Pix[__obf_dfa50e8a4ab3e6d5+1] = Cb[__obf_5a46f1764400ebbd]
			Pix[__obf_dfa50e8a4ab3e6d5+2] = Cr[__obf_5a46f1764400ebbd]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_5fd02db579f190b4
}

func __obf_1ea11da1a27ac5db(__obf_3ccd19e51488be10 *image.YCbCr, __obf_5fd02db579f190b4 *__obf_9474021cedba6d1a) *__obf_9474021cedba6d1a {
	var __obf_dfa50e8a4ab3e6d5 int
	Pix := __obf_5fd02db579f190b4.Pix
	Y := __obf_3ccd19e51488be10.Y
	Cb := __obf_3ccd19e51488be10.Cb
	Cr := __obf_3ccd19e51488be10.Cr
	for __obf_4811266080d7795e := 0; __obf_4811266080d7795e < __obf_3ccd19e51488be10.Rect.Max.Y-__obf_3ccd19e51488be10.Rect.Min.Y; __obf_4811266080d7795e++ {
		__obf_6680b431eeccfaeb := __obf_4811266080d7795e * __obf_3ccd19e51488be10.YStride
		__obf_5ad115631709a2a3 := (__obf_4811266080d7795e / 2) * __obf_3ccd19e51488be10.CStride
		for __obf_21a3146223eb514f := 0; __obf_21a3146223eb514f < __obf_3ccd19e51488be10.Rect.Max.X-__obf_3ccd19e51488be10.Rect.Min.X; __obf_21a3146223eb514f++ {
			__obf_5a46f1764400ebbd := __obf_5ad115631709a2a3 + __obf_21a3146223eb514f/4
			Pix[__obf_dfa50e8a4ab3e6d5+0] = Y[__obf_6680b431eeccfaeb+__obf_21a3146223eb514f]
			Pix[__obf_dfa50e8a4ab3e6d5+1] = Cb[__obf_5a46f1764400ebbd]
			Pix[__obf_dfa50e8a4ab3e6d5+2] = Cr[__obf_5a46f1764400ebbd]
			__obf_dfa50e8a4ab3e6d5 += 3
		}
	}
	return __obf_5fd02db579f190b4
}
