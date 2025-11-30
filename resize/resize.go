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
package __obf_42b2ccbdafaee9c3

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
func (__obf_e39e622bfcb683d9 InterpolationFunction) __obf_9e806529abfb439f() (int, func(float64) float64) {
	switch __obf_e39e622bfcb683d9 {
	case Bilinear:
		return 2, __obf_51af381b36e05225
	case Bicubic:
		return 4, __obf_cfab9d784a3a0826
	case MitchellNetravali:
		return 4, __obf_dfd3b61eaeee906a
	case Lanczos2:
		return 4, __obf_33d6fa64374d1a50
	case Lanczos3:
		return 6, __obf_7f1d3df8ac799ca4
	default:
		// Default to NearestNeighbor.
		return 2, __obf_79c8d4ae69d6ba30
	}
}

// values <1 will sharpen the image
var __obf_eef0584a9ec97f3d = 1.0

// Resize scales an image to new width and height using the interpolation function interp.
// A new image with the given dimensions will be returned.
// If one of the parameters width or height is set to 0, its size will be calculated so that
// the aspect ratio is that of the originating image.
// The resizing algorithm uses channels for parallel computation.
// If the input image has width or height of 0, it is returned unchanged.
func Resize(__obf_10d8cb50cdca170a, __obf_a3be9e36fdc9e4e7 uint, __obf_99d54e4a80501ee0 image.Image, __obf_f1c487ecb00e0380 InterpolationFunction) image.Image {
	__obf_8eee20224e9cd53a, __obf_feb5ea3c2e40a3e0 := __obf_9586aad95779844e(__obf_10d8cb50cdca170a, __obf_a3be9e36fdc9e4e7, float64(__obf_99d54e4a80501ee0.Bounds().Dx()), float64(__obf_99d54e4a80501ee0.Bounds().Dy()))
	if __obf_10d8cb50cdca170a == 0 {
		__obf_10d8cb50cdca170a = uint(0.7 + float64(__obf_99d54e4a80501ee0.Bounds().Dx())/__obf_8eee20224e9cd53a)
	}
	if __obf_a3be9e36fdc9e4e7 == 0 {
		__obf_a3be9e36fdc9e4e7 = uint(0.7 + float64(__obf_99d54e4a80501ee0.Bounds().Dy())/__obf_feb5ea3c2e40a3e0)
	}

	// Trivial case: return input image
	if int(__obf_10d8cb50cdca170a) == __obf_99d54e4a80501ee0.Bounds().Dx() && int(__obf_a3be9e36fdc9e4e7) == __obf_99d54e4a80501ee0.Bounds().Dy() {
		return __obf_99d54e4a80501ee0
	}

	// Input image has no pixels
	if __obf_99d54e4a80501ee0.Bounds().Dx() <= 0 || __obf_99d54e4a80501ee0.Bounds().Dy() <= 0 {
		return __obf_99d54e4a80501ee0
	}

	if __obf_f1c487ecb00e0380 == NearestNeighbor {
		return __obf_600af590ff8d7a6f(__obf_10d8cb50cdca170a, __obf_a3be9e36fdc9e4e7, __obf_8eee20224e9cd53a, __obf_feb5ea3c2e40a3e0, __obf_99d54e4a80501ee0, __obf_f1c487ecb00e0380)
	}
	__obf_e69c5aa1970d4645, __obf_9e806529abfb439f := __obf_f1c487ecb00e0380.__obf_9e806529abfb439f()
	__obf_32909e41f434a230 := runtime.GOMAXPROCS(0)
	__obf_c591bc0a2a52ef77 := sync.WaitGroup{}

	// Generic access to image.Image is slow in tight loops.
	// The optimal access has to be determined from the concrete image type.
	switch __obf_305b51ae03473d18 := __obf_99d54e4a80501ee0.(type) {
	case *image.RGBA:
		__obf_d99ff2d491945013 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ea0d9b36f90839c7(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_efed9ef13011e083(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ea0d9b36f90839c7(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_efed9ef13011e083(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// 8-bit precision
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.NRGBA:
		__obf_d99ff2d491945013 := image.NewRGBA(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ea0d9b36f90839c7(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_32909e41f434a230; __obf_e39e622bfcb683d9++ {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_bd7f833c4584c433(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ea0d9b36f90839c7(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_efed9ef13011e083(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// 8-bit precision
					// accessing the YCbCr arrays in a tight loop is slow.
					// converting the image to ycc increases performance by 2x.
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05

	case *image.YCbCr:
		__obf_d99ff2d491945013 := __obf_78d696b65413d950(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)), __obf_305b51ae03473d18.SubsampleRatio)
		__obf_2c4de8da54695f05 := __obf_78d696b65413d950(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)), image.YCbCrSubsampleRatio444)
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ea0d9b36f90839c7(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c82b366489672bd7 := __obf_79e832e234bf2561(__obf_305b51ae03473d18)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_32909e41f434a230; __obf_e39e622bfcb683d9++ {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*__obf_e4d16584f024f650)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_4b9ffd27224f812a(__obf_c82b366489672bd7, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ea0d9b36f90839c7(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_32909e41f434a230; __obf_e39e622bfcb683d9++ {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*__obf_e4d16584f024f650)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_4b9ffd27224f812a(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05.YCbCr()
	case *image.RGBA64:
		__obf_d99ff2d491945013 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA64(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_4f01f1c9705b203c(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_59ce29d524ed837c(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_4f01f1c9705b203c(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_59ce29d524ed837c(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// 16-bit precision
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.NRGBA64:
		__obf_d99ff2d491945013 := image.NewRGBA64(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA64(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_4f01f1c9705b203c(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_9fa2c4083eecedb9(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_4f01f1c9705b203c(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_59ce29d524ed837c(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// 8-bit precision
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.Gray:
		__obf_d99ff2d491945013 := image.NewGray(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewGray(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ea0d9b36f90839c7(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_32909e41f434a230; __obf_e39e622bfcb683d9++ {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_a5d4dad1cf1a65b6(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ea0d9b36f90839c7(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_a5d4dad1cf1a65b6(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// 16-bit precision
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.Gray16:
		__obf_d99ff2d491945013 := image.NewGray16(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewGray16(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_4f01f1c9705b203c(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray16)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_1ff451e772d836c2(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_4f01f1c9705b203c(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray16)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_1ff451e772d836c2(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// 16-bit precision
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	default:
		__obf_d99ff2d491945013 := image.NewRGBA64(image.Rect(0, 0, __obf_99d54e4a80501ee0.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA64(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_4f01f1c9705b203c(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_17d6ef892ec1014e(__obf_99d54e4a80501ee0, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_4f01f1c9705b203c(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0, __obf_9e806529abfb439f)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_59ce29d524ed837c(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	}
}

func __obf_600af590ff8d7a6f(__obf_10d8cb50cdca170a, __obf_a3be9e36fdc9e4e7 uint, __obf_8eee20224e9cd53a, __obf_feb5ea3c2e40a3e0 float64, __obf_99d54e4a80501ee0 image.Image, __obf_f1c487ecb00e0380 InterpolationFunction) image.Image {
	__obf_e69c5aa1970d4645, _ := __obf_f1c487ecb00e0380.__obf_9e806529abfb439f()
	__obf_32909e41f434a230 := runtime.GOMAXPROCS(0)
	__obf_c591bc0a2a52ef77 := sync.WaitGroup{}

	switch __obf_305b51ae03473d18 := __obf_99d54e4a80501ee0.(type) {
	case *image.RGBA:
		__obf_d99ff2d491945013 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := 0; __obf_e39e622bfcb683d9 < __obf_32909e41f434a230; __obf_e39e622bfcb683d9++ {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_ad2a732b84ec2840(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// horizontal filter on transposed image, result is not transposed
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_ad2a732b84ec2840(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.NRGBA:
		__obf_d99ff2d491945013 := // 8-bit precision
			image.NewNRGBA(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewNRGBA(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.NRGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_9b903e60ecbdcdd4(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// horizontal filter on transposed image, result is not transposed
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.NRGBA)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_9b903e60ecbdcdd4(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.YCbCr:
		__obf_d99ff2d491945013 := // 8-bit precision
			// accessing the YCbCr arrays in a tight loop is slow.
			// converting the image to ycc increases performance by 2x.
			__obf_78d696b65413d950(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)), __obf_305b51ae03473d18.SubsampleRatio)
		__obf_2c4de8da54695f05 := __obf_78d696b65413d950(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)), image.YCbCrSubsampleRatio444)
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c82b366489672bd7 := __obf_79e832e234bf2561(__obf_305b51ae03473d18)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*__obf_e4d16584f024f650)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_a7249479cb109931(__obf_c82b366489672bd7, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*__obf_e4d16584f024f650)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_a7249479cb109931(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_feb5ea3c2e40a3e0, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05.YCbCr()
	case *image.RGBA64:
		__obf_d99ff2d491945013 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA64(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_e49384d8b15c2eb1(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// horizontal filter on transposed image, result is not transposed
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_e49384d8b15c2eb1(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.NRGBA64:
		__obf_d99ff2d491945013 := // 16-bit precision
			image.NewNRGBA64(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewNRGBA64(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.NRGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_c4d5850f6cc798c7(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// horizontal filter on transposed image, result is not transposed
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.NRGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_c4d5850f6cc798c7(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.Gray:
		__obf_d99ff2d491945013 := // 8-bit precision
			image.NewGray(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewGray(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_82962b2485b60364(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// horizontal filter on transposed image, result is not transposed
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_82962b2485b60364(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	case *image.Gray16:
		__obf_d99ff2d491945013 := // 16-bit precision
			image.NewGray16(image.Rect(0, 0, __obf_305b51ae03473d18.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewGray16(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray16)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_9a59e86fa91ba99d(__obf_305b51ae03473d18, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// horizontal filter on transposed image, result is not transposed
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.Gray16)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_9a59e86fa91ba99d(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	default:
		__obf_d99ff2d491945013 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_99d54e4a80501ee0.Bounds().Dy(), int(__obf_10d8cb50cdca170a)))
		__obf_2c4de8da54695f05 := image.NewRGBA64(image.Rect(0, 0, int(__obf_10d8cb50cdca170a), int(__obf_a3be9e36fdc9e4e7)))
		__obf_71732cf03a2f4bc1,

			// horizontal filter, results in transposed temporary image
			__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 := __obf_ed15c920a71519da(__obf_d99ff2d491945013.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_8eee20224e9cd53a)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_d99ff2d491945013, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_fed76ae352e86a98(__obf_99d54e4a80501ee0, __obf_27b21440dffba825, __obf_8eee20224e9cd53a, __obf_71732cf03a2f4bc1,

					// horizontal filter on transposed image, result is not transposed
					__obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		__obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32, __obf_b34b7524c4bfcd41 = __obf_ed15c920a71519da(__obf_2c4de8da54695f05.Bounds().Dy(), __obf_e69c5aa1970d4645, __obf_eef0584a9ec97f3d, __obf_feb5ea3c2e40a3e0)
		__obf_c591bc0a2a52ef77.
			Add(__obf_32909e41f434a230)
		for __obf_e39e622bfcb683d9 := range __obf_32909e41f434a230 {
			__obf_27b21440dffba825 := __obf_42c804e785a14918(__obf_2c4de8da54695f05, __obf_e39e622bfcb683d9, __obf_32909e41f434a230).(*image.RGBA64)
			go func() {
				defer __obf_c591bc0a2a52ef77.Done()
				__obf_e49384d8b15c2eb1(__obf_d99ff2d491945013, __obf_27b21440dffba825, __obf_71732cf03a2f4bc1, __obf_c22f1458e7509d32,

					// Calculates scaling factors using old and new image dimensions.
					__obf_b34b7524c4bfcd41)
			}()
		}
		__obf_c591bc0a2a52ef77.
			Wait()
		return __obf_2c4de8da54695f05
	}

}

func __obf_9586aad95779844e(__obf_10d8cb50cdca170a, __obf_a3be9e36fdc9e4e7 uint, __obf_d74e8faeb30cad68, __obf_c30f84decd3882e8 float64) (__obf_8eee20224e9cd53a, __obf_feb5ea3c2e40a3e0 float64) {
	if __obf_10d8cb50cdca170a == 0 {
		if __obf_a3be9e36fdc9e4e7 == 0 {
			__obf_8eee20224e9cd53a = 1.0
			__obf_feb5ea3c2e40a3e0 = 1.0
		} else {
			__obf_feb5ea3c2e40a3e0 = __obf_c30f84decd3882e8 / float64(__obf_a3be9e36fdc9e4e7)
			__obf_8eee20224e9cd53a = __obf_feb5ea3c2e40a3e0
		}
	} else {
		__obf_8eee20224e9cd53a = __obf_d74e8faeb30cad68 / float64(__obf_10d8cb50cdca170a)
		if __obf_a3be9e36fdc9e4e7 == 0 {
			__obf_feb5ea3c2e40a3e0 = __obf_8eee20224e9cd53a
		} else {
			__obf_feb5ea3c2e40a3e0 = __obf_c30f84decd3882e8 / float64(__obf_a3be9e36fdc9e4e7)
		}
	}
	return
}

type __obf_cbdfe5a00e73d422 interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func __obf_42c804e785a14918(__obf_99d54e4a80501ee0 __obf_cbdfe5a00e73d422, __obf_e39e622bfcb683d9, __obf_3a317d8c7f63320c int) image.Image {
	return __obf_99d54e4a80501ee0.SubImage(image.Rect(__obf_99d54e4a80501ee0.Bounds().Min.X, __obf_99d54e4a80501ee0.Bounds().Min.Y+__obf_e39e622bfcb683d9*__obf_99d54e4a80501ee0.Bounds().Dy()/__obf_3a317d8c7f63320c, __obf_99d54e4a80501ee0.Bounds().Max.X, __obf_99d54e4a80501ee0.Bounds().Min.Y+(__obf_e39e622bfcb683d9+1)*__obf_99d54e4a80501ee0.Bounds().Dy()/__obf_3a317d8c7f63320c))
}
