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

// Package resize implements various image resizing methods.
//
// The package works with the Image interface described in the image package.
// Various interpolation methods are provided and multiple processors may be
// utilized in the computations.
//
// Example:
//
//	imgResized := resize.Resize(1000, 0, imgOld, resize.MitchellNetravali)
package __obf_d24101647651f4c3

import (
	"image"
	"runtime"
	"sync"
)

// An InterpolationFunction provides the parameters that describe an
// interpolation kernel. It returns the number of samples to take
// and the kernel function to use for sampling.
type InterpolationFunction int

// InterpolationFunction constants
const (
	// Nearest-neighbor interpolation
	NearestNeighbor InterpolationFunction = iota
	// Bilinear interpolation
	Bilinear
	// Bicubic interpolation (with cubic hermite spline)
	Bicubic
	// Mitchell-Netravali interpolation
	MitchellNetravali
	// Lanczos interpolation (a=2)
	Lanczos2
	// Lanczos interpolation (a=3)
	Lanczos3
)

// kernal, returns an InterpolationFunctions taps and kernel.
func (__obf_41043e9a68df34f8 InterpolationFunction) __obf_0a5c3b331d82d518() (int, func(float64) float64) {
	switch __obf_41043e9a68df34f8 {
	case Bilinear:
		return 2, __obf_d66ee4de441cbd66
	case Bicubic:
		return 4, __obf_a619d5b55f4394a4
	case MitchellNetravali:
		return 4, __obf_7a44659276a985b3
	case Lanczos2:
		return 4, __obf_b3d8a34685baa99a
	case Lanczos3:
		return 6, __obf_d668a1e333340f8d
	default:
		// Default to NearestNeighbor.
		return 2, __obf_e958fd3b8bea8d9c
	}
}

// values <1 will sharpen the image
var __obf_b8f04e12843e5823 = 1.0

// Resize scales an image to new width and height using the interpolation function interp.
// A new image with the given dimensions will be returned.
// If one of the parameters width or height is set to 0, its size will be calculated so that
// the aspect ratio is that of the originating image.
// The resizing algorithm uses channels for parallel computation.
// If the input image has width or height of 0, it is returned unchanged.
func Resize(__obf_a22ab8344dedd519, __obf_2667e7551de50ad4 uint, __obf_b9c5bf10bfae0f7a image.Image, __obf_0c2c83d39fca0ad5 InterpolationFunction) image.Image {
	__obf_5ba5939ea89ea0de, __obf_e469e0e79b6f8761 := __obf_0c8f992b82cb1c93(__obf_a22ab8344dedd519, __obf_2667e7551de50ad4, float64(__obf_b9c5bf10bfae0f7a.Bounds().Dx()), float64(__obf_b9c5bf10bfae0f7a.Bounds().Dy()))
	if __obf_a22ab8344dedd519 == 0 {
		__obf_a22ab8344dedd519 = uint(0.7 + float64(__obf_b9c5bf10bfae0f7a.Bounds().Dx())/__obf_5ba5939ea89ea0de)
	}
	if __obf_2667e7551de50ad4 == 0 {
		__obf_2667e7551de50ad4 = uint(0.7 + float64(__obf_b9c5bf10bfae0f7a.Bounds().Dy())/__obf_e469e0e79b6f8761)
	}

	// Trivial case: return input image
	if int(__obf_a22ab8344dedd519) == __obf_b9c5bf10bfae0f7a.Bounds().Dx() && int(__obf_2667e7551de50ad4) == __obf_b9c5bf10bfae0f7a.Bounds().Dy() {
		return __obf_b9c5bf10bfae0f7a
	}

	// Input image has no pixels
	if __obf_b9c5bf10bfae0f7a.Bounds().Dx() <= 0 || __obf_b9c5bf10bfae0f7a.Bounds().Dy() <= 0 {
		return __obf_b9c5bf10bfae0f7a
	}

	if __obf_0c2c83d39fca0ad5 == NearestNeighbor {
		return __obf_13e95d81671719c7(__obf_a22ab8344dedd519, __obf_2667e7551de50ad4, __obf_5ba5939ea89ea0de, __obf_e469e0e79b6f8761, __obf_b9c5bf10bfae0f7a, __obf_0c2c83d39fca0ad5)
	}
	__obf_3a126acc9d6ba2c1, __obf_0a5c3b331d82d518 := __obf_0c2c83d39fca0ad5.__obf_0a5c3b331d82d518()
	__obf_ca7e6a5f6d539d14 := runtime.GOMAXPROCS(0)
	__obf_e9b9d04aa0aee136 := sync.WaitGroup{}

	// Generic access to image.Image is slow in tight loops.
	// The optimal access has to be determined from the concrete image type.
	switch __obf_2f9019f1e31b6623 := __obf_b9c5bf10bfae0f7a.(type) {
	case *image.RGBA:
		__obf_0c62244a1c8127bf := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_b32cca196c808fe4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_8ec8d60aa6453477(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_b32cca196c808fe4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_8ec8d60aa6453477(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// 8-bit precision
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.NRGBA:
		__obf_0c62244a1c8127bf := image.NewRGBA(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_b32cca196c808fe4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_ca7e6a5f6d539d14; __obf_41043e9a68df34f8++ {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_f38f8e4e927a9b1f(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_b32cca196c808fe4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_8ec8d60aa6453477(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// 8-bit precision
					// accessing the YCbCr arrays in a tight loop is slow.
					// converting the image to ycc increases performance by 2x.
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826

	case *image.YCbCr:
		__obf_0c62244a1c8127bf := __obf_526e91fd49ed301d(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)), __obf_2f9019f1e31b6623.SubsampleRatio)
		__obf_050e6c4f8fdd3826 := __obf_526e91fd49ed301d(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)), image.YCbCrSubsampleRatio444)
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_b32cca196c808fe4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_2a764779fc1401ce := __obf_47343abcf82d0c31(__obf_2f9019f1e31b6623)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_ca7e6a5f6d539d14; __obf_41043e9a68df34f8++ {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*__obf_489a9befc4898050)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_5df6e1d64c485b9f(__obf_2a764779fc1401ce, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_b32cca196c808fe4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_ca7e6a5f6d539d14; __obf_41043e9a68df34f8++ {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*__obf_489a9befc4898050)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_5df6e1d64c485b9f(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826.YCbCr()
	case *image.RGBA64:
		__obf_0c62244a1c8127bf := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA64(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_2a2dd68742b912d4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_454a0ae091c9479b(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_2a2dd68742b912d4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_454a0ae091c9479b(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// 16-bit precision
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.NRGBA64:
		__obf_0c62244a1c8127bf := image.NewRGBA64(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA64(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_2a2dd68742b912d4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_a658dbd4fd6775a7(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_2a2dd68742b912d4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_454a0ae091c9479b(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// 8-bit precision
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.Gray:
		__obf_0c62244a1c8127bf := image.NewGray(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewGray(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_b32cca196c808fe4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_ca7e6a5f6d539d14; __obf_41043e9a68df34f8++ {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_1eb0358183feb59b(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_b32cca196c808fe4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_1eb0358183feb59b(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// 16-bit precision
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.Gray16:
		__obf_0c62244a1c8127bf := image.NewGray16(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewGray16(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_2a2dd68742b912d4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray16)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_a51435ed23b2f46f(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_2a2dd68742b912d4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray16)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_a51435ed23b2f46f(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// 16-bit precision
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	default:
		__obf_0c62244a1c8127bf := image.NewRGBA64(image.Rect(0, 0, __obf_b9c5bf10bfae0f7a.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA64(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_2a2dd68742b912d4(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_989e6431922ef0ad(__obf_b9c5bf10bfae0f7a, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_2a2dd68742b912d4(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761, __obf_0a5c3b331d82d518)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_454a0ae091c9479b(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	}
}

func __obf_13e95d81671719c7(__obf_a22ab8344dedd519, __obf_2667e7551de50ad4 uint, __obf_5ba5939ea89ea0de, __obf_e469e0e79b6f8761 float64, __obf_b9c5bf10bfae0f7a image.Image, __obf_0c2c83d39fca0ad5 InterpolationFunction) image.Image {
	__obf_3a126acc9d6ba2c1, _ := __obf_0c2c83d39fca0ad5.__obf_0a5c3b331d82d518()
	__obf_ca7e6a5f6d539d14 := runtime.GOMAXPROCS(0)
	__obf_e9b9d04aa0aee136 := sync.WaitGroup{}

	switch __obf_2f9019f1e31b6623 := __obf_b9c5bf10bfae0f7a.(type) {
	case *image.RGBA:
		__obf_0c62244a1c8127bf := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := 0; __obf_41043e9a68df34f8 < __obf_ca7e6a5f6d539d14; __obf_41043e9a68df34f8++ {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_f6fedca77d332791(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// horizontal filter on transposed image, result is not transposed
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_f6fedca77d332791(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.NRGBA:
		__obf_0c62244a1c8127bf := // 8-bit precision
			image.NewNRGBA(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewNRGBA(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.NRGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_bdf16a224669ff32(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// horizontal filter on transposed image, result is not transposed
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.NRGBA)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_bdf16a224669ff32(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.YCbCr:
		__obf_0c62244a1c8127bf := // 8-bit precision
			// accessing the YCbCr arrays in a tight loop is slow.
			// converting the image to ycc increases performance by 2x.
			__obf_526e91fd49ed301d(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)), __obf_2f9019f1e31b6623.SubsampleRatio)
		__obf_050e6c4f8fdd3826 := __obf_526e91fd49ed301d(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)), image.YCbCrSubsampleRatio444)
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_2a764779fc1401ce := __obf_47343abcf82d0c31(__obf_2f9019f1e31b6623)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*__obf_489a9befc4898050)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_8b9375bb53904eb4(__obf_2a764779fc1401ce, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*__obf_489a9befc4898050)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_8b9375bb53904eb4(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_e469e0e79b6f8761, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826.YCbCr()
	case *image.RGBA64:
		__obf_0c62244a1c8127bf := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA64(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_bb156615cc9077bb(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// horizontal filter on transposed image, result is not transposed
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_bb156615cc9077bb(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.NRGBA64:
		__obf_0c62244a1c8127bf := // 16-bit precision
			image.NewNRGBA64(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewNRGBA64(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.NRGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_cd535ce510158a2c(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// horizontal filter on transposed image, result is not transposed
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.NRGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_cd535ce510158a2c(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.Gray:
		__obf_0c62244a1c8127bf := // 8-bit precision
			image.NewGray(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewGray(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_5fc478f60b2eec87(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// horizontal filter on transposed image, result is not transposed
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_5fc478f60b2eec87(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	case *image.Gray16:
		__obf_0c62244a1c8127bf := // 16-bit precision
			image.NewGray16(image.Rect(0, 0, __obf_2f9019f1e31b6623.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewGray16(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray16)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_a8f9385dd16a2f12(__obf_2f9019f1e31b6623, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// horizontal filter on transposed image, result is not transposed
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.Gray16)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_a8f9385dd16a2f12(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	default:
		__obf_0c62244a1c8127bf := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_b9c5bf10bfae0f7a.Bounds().Dy(), int(__obf_a22ab8344dedd519)))
		__obf_050e6c4f8fdd3826 := image.NewRGBA64(image.Rect(0, 0, int(__obf_a22ab8344dedd519), int(__obf_2667e7551de50ad4)))
		__obf_779a7f81421e09ef,

			// horizontal filter, results in transposed temporary image
			__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e := __obf_073d620566900e2c(__obf_0c62244a1c8127bf.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_5ba5939ea89ea0de)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_0c62244a1c8127bf, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_92a9ee3518f7cb1f(__obf_b9c5bf10bfae0f7a, __obf_a51405ad505373a9, __obf_5ba5939ea89ea0de, __obf_779a7f81421e09ef,

					// horizontal filter on transposed image, result is not transposed
					__obf_04772264944a3f8f, __obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		__obf_779a7f81421e09ef, __obf_04772264944a3f8f, __obf_277f8fd7946d0e2e = __obf_073d620566900e2c(__obf_050e6c4f8fdd3826.Bounds().Dy(), __obf_3a126acc9d6ba2c1, __obf_b8f04e12843e5823, __obf_e469e0e79b6f8761)
		__obf_e9b9d04aa0aee136.
			Add(__obf_ca7e6a5f6d539d14)
		for __obf_41043e9a68df34f8 := range __obf_ca7e6a5f6d539d14 {
			__obf_a51405ad505373a9 := __obf_bd75b5bfd7ac8518(__obf_050e6c4f8fdd3826, __obf_41043e9a68df34f8, __obf_ca7e6a5f6d539d14).(*image.RGBA64)
			go func() {
				defer __obf_e9b9d04aa0aee136.Done()
				__obf_bb156615cc9077bb(__obf_0c62244a1c8127bf, __obf_a51405ad505373a9, __obf_779a7f81421e09ef, __obf_04772264944a3f8f,

					// Calculates scaling factors using old and new image dimensions.
					__obf_277f8fd7946d0e2e)
			}()
		}
		__obf_e9b9d04aa0aee136.
			Wait()
		return __obf_050e6c4f8fdd3826
	}

}

func __obf_0c8f992b82cb1c93(__obf_a22ab8344dedd519, __obf_2667e7551de50ad4 uint, __obf_c3f789899b3b05f4, __obf_cd01440d64e39479 float64) (__obf_5ba5939ea89ea0de, __obf_e469e0e79b6f8761 float64) {
	if __obf_a22ab8344dedd519 == 0 {
		if __obf_2667e7551de50ad4 == 0 {
			__obf_5ba5939ea89ea0de = 1.0
			__obf_e469e0e79b6f8761 = 1.0
		} else {
			__obf_e469e0e79b6f8761 = __obf_cd01440d64e39479 / float64(__obf_2667e7551de50ad4)
			__obf_5ba5939ea89ea0de = __obf_e469e0e79b6f8761
		}
	} else {
		__obf_5ba5939ea89ea0de = __obf_c3f789899b3b05f4 / float64(__obf_a22ab8344dedd519)
		if __obf_2667e7551de50ad4 == 0 {
			__obf_e469e0e79b6f8761 = __obf_5ba5939ea89ea0de
		} else {
			__obf_e469e0e79b6f8761 = __obf_cd01440d64e39479 / float64(__obf_2667e7551de50ad4)
		}
	}
	return
}

type __obf_ebabb0c464b3cb0c interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func __obf_bd75b5bfd7ac8518(__obf_b9c5bf10bfae0f7a __obf_ebabb0c464b3cb0c, __obf_41043e9a68df34f8, __obf_ad1856af42b6f59d int) image.Image {
	return __obf_b9c5bf10bfae0f7a.SubImage(image.Rect(__obf_b9c5bf10bfae0f7a.Bounds().Min.X, __obf_b9c5bf10bfae0f7a.Bounds().Min.Y+__obf_41043e9a68df34f8*__obf_b9c5bf10bfae0f7a.Bounds().Dy()/__obf_ad1856af42b6f59d, __obf_b9c5bf10bfae0f7a.Bounds().Max.X, __obf_b9c5bf10bfae0f7a.Bounds().Min.Y+(__obf_41043e9a68df34f8+1)*__obf_b9c5bf10bfae0f7a.Bounds().Dy()/__obf_ad1856af42b6f59d))
}
