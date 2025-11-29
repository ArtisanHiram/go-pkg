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

package __obf_3858dbfa2ccfdbe9

import (
	"image"
	"image/color"
)

// ycc is an in memory YCbCr image.  The Y, Cb and Cr samples are held in a
// single slice to increase resizing performance.
type __obf_1f271b19471a7add struct {
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
func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) PixOffset(__obf_15a5faef6b57be64, __obf_523cb3ec667c6a54 int) int {
	return (__obf_523cb3ec667c6a54-__obf_b6ea2276e9a14426.Rect.Min.Y)*__obf_b6ea2276e9a14426.Stride + (__obf_15a5faef6b57be64-__obf_b6ea2276e9a14426.Rect.Min.X)*3
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) Bounds() image.Rectangle {
	return __obf_b6ea2276e9a14426.Rect
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) ColorModel() color.Model {
	return color.YCbCrModel
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) At(__obf_15a5faef6b57be64, __obf_523cb3ec667c6a54 int) color.Color {
	if !(image.Point{__obf_15a5faef6b57be64, __obf_523cb3ec667c6a54}.In(__obf_b6ea2276e9a14426.Rect)) {
		return color.YCbCr{}
	}
	__obf_9d9681e93d5a567f := __obf_b6ea2276e9a14426.PixOffset(__obf_15a5faef6b57be64, __obf_523cb3ec667c6a54)
	return color.YCbCr{__obf_b6ea2276e9a14426.
		Pix[__obf_9d9681e93d5a567f+0], __obf_b6ea2276e9a14426.
		Pix[__obf_9d9681e93d5a567f+1], __obf_b6ea2276e9a14426.
		Pix[__obf_9d9681e93d5a567f+2],
	}
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) Opaque() bool {
	return true
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) SubImage(__obf_c3d2c2fc3dcef772 image.Rectangle) image.Image {
	__obf_c3d2c2fc3dcef772 = __obf_c3d2c2fc3dcef772.Intersect(__obf_b6ea2276e9a14426.Rect)
	if __obf_c3d2c2fc3dcef772.Empty() {
		return &__obf_1f271b19471a7add{SubsampleRatio: __obf_b6ea2276e9a14426.SubsampleRatio}
	}
	__obf_9d9681e93d5a567f := __obf_b6ea2276e9a14426.PixOffset(__obf_c3d2c2fc3dcef772.Min.X, __obf_c3d2c2fc3dcef772.Min.Y)
	return &__obf_1f271b19471a7add{
		Pix:            __obf_b6ea2276e9a14426.Pix[__obf_9d9681e93d5a567f:],
		Stride:         __obf_b6ea2276e9a14426.Stride,
		Rect:           __obf_c3d2c2fc3dcef772,
		SubsampleRatio: __obf_b6ea2276e9a14426.SubsampleRatio,
	}
}

// newYCC returns a new ycc with the given bounds and subsample ratio.
func __obf_8decfa8633084a6c(__obf_c3d2c2fc3dcef772 image.Rectangle, __obf_ee54ca98923ffa73 image.YCbCrSubsampleRatio) *__obf_1f271b19471a7add {
	__obf_c1bd60223d47dc21, __obf_5344f52927fd9577 := __obf_c3d2c2fc3dcef772.Dx(), __obf_c3d2c2fc3dcef772.Dy()
	__obf_a4ff4c1754a1421e := make([]uint8, 3*__obf_c1bd60223d47dc21*__obf_5344f52927fd9577)
	return &__obf_1f271b19471a7add{Pix: __obf_a4ff4c1754a1421e, Stride: 3 * __obf_c1bd60223d47dc21, Rect: __obf_c3d2c2fc3dcef772, SubsampleRatio: __obf_ee54ca98923ffa73}
}

// Copy of image.YCbCrSubsampleRatio constants - this allows us to support
// older versions of Go where these constants are not defined (i.e. Go 1.4)
const (
	__obf_a51eb1ac01c52766 image.YCbCrSubsampleRatio = iota
	__obf_70a35342ecf0f6e8
	__obf_1ae5ab1faf1fcf02
	__obf_4eb784d97cdc6886
	__obf_1370860b42b4d908
	__obf_d923197952916b8a
)

// YCbCr converts ycc to a YCbCr image with the same subsample ratio
// as the YCbCr image that ycc was generated from.
func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) YCbCr() *image.YCbCr {
	__obf_7edbb7a86f1dc3db := image.NewYCbCr(__obf_b6ea2276e9a14426.Rect, __obf_b6ea2276e9a14426.SubsampleRatio)
	switch __obf_7edbb7a86f1dc3db.SubsampleRatio {
	case __obf_70a35342ecf0f6e8:
		return __obf_b6ea2276e9a14426.__obf_139dbd1b460fabce(__obf_7edbb7a86f1dc3db)
	case __obf_1ae5ab1faf1fcf02:
		return __obf_b6ea2276e9a14426.__obf_c7bd024d129d660d(__obf_7edbb7a86f1dc3db)
	case __obf_4eb784d97cdc6886:
		return __obf_b6ea2276e9a14426.__obf_5f099b888855e0ca(__obf_7edbb7a86f1dc3db)
	case __obf_a51eb1ac01c52766:
		return __obf_b6ea2276e9a14426.__obf_db71ff9eab34d84d(__obf_7edbb7a86f1dc3db)
	case __obf_1370860b42b4d908:
		return __obf_b6ea2276e9a14426.__obf_00d0cae3b4c686d7(__obf_7edbb7a86f1dc3db)
	case __obf_d923197952916b8a:
		return __obf_b6ea2276e9a14426.__obf_f4a5a8664dd548bb(__obf_7edbb7a86f1dc3db)
	}
	return __obf_7edbb7a86f1dc3db
}

// imageYCbCrToYCC converts a YCbCr image to a ycc image for resizing.
func __obf_6235574c4a9792da(__obf_8da81bf3c5df7f0d *image.YCbCr) *__obf_1f271b19471a7add {
	__obf_c1bd60223d47dc21, __obf_5344f52927fd9577 := __obf_8da81bf3c5df7f0d.Rect.Dx(), __obf_8da81bf3c5df7f0d.Rect.Dy()
	__obf_b6ea2276e9a14426 := __obf_1f271b19471a7add{
		Pix:            make([]uint8, 3*__obf_c1bd60223d47dc21*__obf_5344f52927fd9577),
		Stride:         3 * __obf_c1bd60223d47dc21,
		Rect:           image.Rect(0, 0, __obf_c1bd60223d47dc21, __obf_5344f52927fd9577),
		SubsampleRatio: __obf_8da81bf3c5df7f0d.SubsampleRatio,
	}
	switch __obf_8da81bf3c5df7f0d.SubsampleRatio {
	case __obf_70a35342ecf0f6e8:
		return __obf_60f66a3f411e9acf(__obf_8da81bf3c5df7f0d, &__obf_b6ea2276e9a14426)
	case __obf_1ae5ab1faf1fcf02:
		return __obf_0b4dad5667322eb0(__obf_8da81bf3c5df7f0d, &__obf_b6ea2276e9a14426)
	case __obf_4eb784d97cdc6886:
		return __obf_f493a0cc26306797(__obf_8da81bf3c5df7f0d, &__obf_b6ea2276e9a14426)
	case __obf_a51eb1ac01c52766:
		return __obf_ac0ef8088e8e13fd(__obf_8da81bf3c5df7f0d, &__obf_b6ea2276e9a14426)
	case __obf_1370860b42b4d908:
		return __obf_5dfd3c059dc59274(__obf_8da81bf3c5df7f0d, &__obf_b6ea2276e9a14426)
	case __obf_d923197952916b8a:
		return __obf_02c2c13d39889465(__obf_8da81bf3c5df7f0d, &__obf_b6ea2276e9a14426)
	}
	return &__obf_b6ea2276e9a14426
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) __obf_139dbd1b460fabce(__obf_7edbb7a86f1dc3db *image.YCbCr) *image.YCbCr {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_7edbb7a86f1dc3db.Y
	Cb := __obf_7edbb7a86f1dc3db.Cb
	Cr := __obf_7edbb7a86f1dc3db.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_7edbb7a86f1dc3db.Rect.Max.Y-__obf_7edbb7a86f1dc3db.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.YStride
		__obf_317be6e3765a4817 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_7edbb7a86f1dc3db.Rect.Max.X-__obf_7edbb7a86f1dc3db.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/2
			Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64] = Pix[__obf_b6004fcf6d72df5a+0]
			Cb[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+1]
			Cr[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+2]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_7edbb7a86f1dc3db
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) __obf_c7bd024d129d660d(__obf_7edbb7a86f1dc3db *image.YCbCr) *image.YCbCr {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_7edbb7a86f1dc3db.Y
	Cb := __obf_7edbb7a86f1dc3db.Cb
	Cr := __obf_7edbb7a86f1dc3db.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_7edbb7a86f1dc3db.Rect.Max.Y-__obf_7edbb7a86f1dc3db.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.YStride
		__obf_317be6e3765a4817 := (__obf_523cb3ec667c6a54 / 2) * __obf_7edbb7a86f1dc3db.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_7edbb7a86f1dc3db.Rect.Max.X-__obf_7edbb7a86f1dc3db.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/2
			Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64] = Pix[__obf_b6004fcf6d72df5a+0]
			Cb[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+1]
			Cr[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+2]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_7edbb7a86f1dc3db
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) __obf_5f099b888855e0ca(__obf_7edbb7a86f1dc3db *image.YCbCr) *image.YCbCr {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_7edbb7a86f1dc3db.Y
	Cb := __obf_7edbb7a86f1dc3db.Cb
	Cr := __obf_7edbb7a86f1dc3db.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_7edbb7a86f1dc3db.Rect.Max.Y-__obf_7edbb7a86f1dc3db.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.YStride
		__obf_317be6e3765a4817 := (__obf_523cb3ec667c6a54 / 2) * __obf_7edbb7a86f1dc3db.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_7edbb7a86f1dc3db.Rect.Max.X-__obf_7edbb7a86f1dc3db.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64
			Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64] = Pix[__obf_b6004fcf6d72df5a+0]
			Cb[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+1]
			Cr[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+2]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_7edbb7a86f1dc3db
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) __obf_db71ff9eab34d84d(__obf_7edbb7a86f1dc3db *image.YCbCr) *image.YCbCr {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_7edbb7a86f1dc3db.Y
	Cb := __obf_7edbb7a86f1dc3db.Cb
	Cr := __obf_7edbb7a86f1dc3db.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_7edbb7a86f1dc3db.Rect.Max.Y-__obf_7edbb7a86f1dc3db.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.YStride
		__obf_317be6e3765a4817 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_7edbb7a86f1dc3db.Rect.Max.X-__obf_7edbb7a86f1dc3db.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64
			Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64] = Pix[__obf_b6004fcf6d72df5a+0]
			Cb[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+1]
			Cr[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+2]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_7edbb7a86f1dc3db
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) __obf_00d0cae3b4c686d7(__obf_7edbb7a86f1dc3db *image.YCbCr) *image.YCbCr {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_7edbb7a86f1dc3db.Y
	Cb := __obf_7edbb7a86f1dc3db.Cb
	Cr := __obf_7edbb7a86f1dc3db.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_7edbb7a86f1dc3db.Rect.Max.Y-__obf_7edbb7a86f1dc3db.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.YStride
		__obf_317be6e3765a4817 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_7edbb7a86f1dc3db.Rect.Max.X-__obf_7edbb7a86f1dc3db.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/4
			Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64] = Pix[__obf_b6004fcf6d72df5a+0]
			Cb[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+1]
			Cr[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+2]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_7edbb7a86f1dc3db
}

func (__obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) __obf_f4a5a8664dd548bb(__obf_7edbb7a86f1dc3db *image.YCbCr) *image.YCbCr {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_7edbb7a86f1dc3db.Y
	Cb := __obf_7edbb7a86f1dc3db.Cb
	Cr := __obf_7edbb7a86f1dc3db.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_7edbb7a86f1dc3db.Rect.Max.Y-__obf_7edbb7a86f1dc3db.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_7edbb7a86f1dc3db.YStride
		__obf_317be6e3765a4817 := (__obf_523cb3ec667c6a54 / 2) * __obf_7edbb7a86f1dc3db.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_7edbb7a86f1dc3db.Rect.Max.X-__obf_7edbb7a86f1dc3db.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/4
			Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64] = Pix[__obf_b6004fcf6d72df5a+0]
			Cb[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+1]
			Cr[__obf_69d5eccaa41f8a07] = Pix[__obf_b6004fcf6d72df5a+2]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_7edbb7a86f1dc3db
}

func __obf_60f66a3f411e9acf(__obf_8da81bf3c5df7f0d *image.YCbCr, __obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) *__obf_1f271b19471a7add {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_8da81bf3c5df7f0d.Y
	Cb := __obf_8da81bf3c5df7f0d.Cb
	Cr := __obf_8da81bf3c5df7f0d.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_8da81bf3c5df7f0d.Rect.Max.Y-__obf_8da81bf3c5df7f0d.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.YStride
		__obf_317be6e3765a4817 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_8da81bf3c5df7f0d.Rect.Max.X-__obf_8da81bf3c5df7f0d.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/2
			Pix[__obf_b6004fcf6d72df5a+0] = Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64]
			Pix[__obf_b6004fcf6d72df5a+1] = Cb[__obf_69d5eccaa41f8a07]
			Pix[__obf_b6004fcf6d72df5a+2] = Cr[__obf_69d5eccaa41f8a07]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_b6ea2276e9a14426
}

func __obf_0b4dad5667322eb0(__obf_8da81bf3c5df7f0d *image.YCbCr, __obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) *__obf_1f271b19471a7add {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_8da81bf3c5df7f0d.Y
	Cb := __obf_8da81bf3c5df7f0d.Cb
	Cr := __obf_8da81bf3c5df7f0d.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_8da81bf3c5df7f0d.Rect.Max.Y-__obf_8da81bf3c5df7f0d.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.YStride
		__obf_317be6e3765a4817 := (__obf_523cb3ec667c6a54 / 2) * __obf_8da81bf3c5df7f0d.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_8da81bf3c5df7f0d.Rect.Max.X-__obf_8da81bf3c5df7f0d.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/2
			Pix[__obf_b6004fcf6d72df5a+0] = Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64]
			Pix[__obf_b6004fcf6d72df5a+1] = Cb[__obf_69d5eccaa41f8a07]
			Pix[__obf_b6004fcf6d72df5a+2] = Cr[__obf_69d5eccaa41f8a07]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_b6ea2276e9a14426
}

func __obf_f493a0cc26306797(__obf_8da81bf3c5df7f0d *image.YCbCr, __obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) *__obf_1f271b19471a7add {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_8da81bf3c5df7f0d.Y
	Cb := __obf_8da81bf3c5df7f0d.Cb
	Cr := __obf_8da81bf3c5df7f0d.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_8da81bf3c5df7f0d.Rect.Max.Y-__obf_8da81bf3c5df7f0d.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.YStride
		__obf_317be6e3765a4817 := (__obf_523cb3ec667c6a54 / 2) * __obf_8da81bf3c5df7f0d.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_8da81bf3c5df7f0d.Rect.Max.X-__obf_8da81bf3c5df7f0d.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64
			Pix[__obf_b6004fcf6d72df5a+0] = Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64]
			Pix[__obf_b6004fcf6d72df5a+1] = Cb[__obf_69d5eccaa41f8a07]
			Pix[__obf_b6004fcf6d72df5a+2] = Cr[__obf_69d5eccaa41f8a07]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_b6ea2276e9a14426
}

func __obf_ac0ef8088e8e13fd(__obf_8da81bf3c5df7f0d *image.YCbCr, __obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) *__obf_1f271b19471a7add {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_8da81bf3c5df7f0d.Y
	Cb := __obf_8da81bf3c5df7f0d.Cb
	Cr := __obf_8da81bf3c5df7f0d.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_8da81bf3c5df7f0d.Rect.Max.Y-__obf_8da81bf3c5df7f0d.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.YStride
		__obf_317be6e3765a4817 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_8da81bf3c5df7f0d.Rect.Max.X-__obf_8da81bf3c5df7f0d.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64
			Pix[__obf_b6004fcf6d72df5a+0] = Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64]
			Pix[__obf_b6004fcf6d72df5a+1] = Cb[__obf_69d5eccaa41f8a07]
			Pix[__obf_b6004fcf6d72df5a+2] = Cr[__obf_69d5eccaa41f8a07]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_b6ea2276e9a14426
}

func __obf_5dfd3c059dc59274(__obf_8da81bf3c5df7f0d *image.YCbCr, __obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) *__obf_1f271b19471a7add {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_8da81bf3c5df7f0d.Y
	Cb := __obf_8da81bf3c5df7f0d.Cb
	Cr := __obf_8da81bf3c5df7f0d.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_8da81bf3c5df7f0d.Rect.Max.Y-__obf_8da81bf3c5df7f0d.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.YStride
		__obf_317be6e3765a4817 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_8da81bf3c5df7f0d.Rect.Max.X-__obf_8da81bf3c5df7f0d.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/4
			Pix[__obf_b6004fcf6d72df5a+0] = Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64]
			Pix[__obf_b6004fcf6d72df5a+1] = Cb[__obf_69d5eccaa41f8a07]
			Pix[__obf_b6004fcf6d72df5a+2] = Cr[__obf_69d5eccaa41f8a07]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_b6ea2276e9a14426
}

func __obf_02c2c13d39889465(__obf_8da81bf3c5df7f0d *image.YCbCr, __obf_b6ea2276e9a14426 *__obf_1f271b19471a7add) *__obf_1f271b19471a7add {
	var __obf_b6004fcf6d72df5a int
	Pix := __obf_b6ea2276e9a14426.Pix
	Y := __obf_8da81bf3c5df7f0d.Y
	Cb := __obf_8da81bf3c5df7f0d.Cb
	Cr := __obf_8da81bf3c5df7f0d.Cr
	for __obf_523cb3ec667c6a54 := 0; __obf_523cb3ec667c6a54 < __obf_8da81bf3c5df7f0d.Rect.Max.Y-__obf_8da81bf3c5df7f0d.Rect.Min.Y; __obf_523cb3ec667c6a54++ {
		__obf_f016c60b2ad12c64 := __obf_523cb3ec667c6a54 * __obf_8da81bf3c5df7f0d.YStride
		__obf_317be6e3765a4817 := (__obf_523cb3ec667c6a54 / 2) * __obf_8da81bf3c5df7f0d.CStride
		for __obf_15a5faef6b57be64 := 0; __obf_15a5faef6b57be64 < __obf_8da81bf3c5df7f0d.Rect.Max.X-__obf_8da81bf3c5df7f0d.Rect.Min.X; __obf_15a5faef6b57be64++ {
			__obf_69d5eccaa41f8a07 := __obf_317be6e3765a4817 + __obf_15a5faef6b57be64/4
			Pix[__obf_b6004fcf6d72df5a+0] = Y[__obf_f016c60b2ad12c64+__obf_15a5faef6b57be64]
			Pix[__obf_b6004fcf6d72df5a+1] = Cb[__obf_69d5eccaa41f8a07]
			Pix[__obf_b6004fcf6d72df5a+2] = Cr[__obf_69d5eccaa41f8a07]
			__obf_b6004fcf6d72df5a += 3
		}
	}
	return __obf_b6ea2276e9a14426
}
