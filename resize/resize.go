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
package __obf_4f71249d3f545aa0

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
func (__obf_b92089dde6a4a00b InterpolationFunction) __obf_25630250e4587e64() (int, func(float64) float64) {
	switch __obf_b92089dde6a4a00b {
	case Bilinear:
		return 2, __obf_62a9d12c047474d5
	case Bicubic:
		return 4, __obf_0a92c4954c46d5b5
	case MitchellNetravali:
		return 4, __obf_34caadcbfbe56d7c
	case Lanczos2:
		return 4, __obf_57e614f28bbd7ca3
	case Lanczos3:
		return 6, __obf_ed3f12ba8d7c6de5
	default:
		// Default to NearestNeighbor.
		return 2, __obf_3d46464e0d553bae
	}
}

// values <1 will sharpen the image
var __obf_ce79a531ca9b78ac = 1.0

// Resize scales an image to new width and height using the interpolation function interp.
// A new image with the given dimensions will be returned.
// If one of the parameters width or height is set to 0, its size will be calculated so that
// the aspect ratio is that of the originating image.
// The resizing algorithm uses channels for parallel computation.
// If the input image has width or height of 0, it is returned unchanged.
func Resize(__obf_8367d6c856df0960, __obf_f8c74f160b478006 uint, __obf_cb2c55730faeb3f0 image.Image, __obf_f617084106645ca0 InterpolationFunction) image.Image {
	__obf_b72b7f9b028a1c41, __obf_fd68fc877fd97a26 := __obf_ef7f92af091e9888(__obf_8367d6c856df0960, __obf_f8c74f160b478006, float64(__obf_cb2c55730faeb3f0.Bounds().Dx()), float64(__obf_cb2c55730faeb3f0.Bounds().Dy()))
	if __obf_8367d6c856df0960 == 0 {
		__obf_8367d6c856df0960 = uint(0.7 + float64(__obf_cb2c55730faeb3f0.Bounds().Dx())/__obf_b72b7f9b028a1c41)
	}
	if __obf_f8c74f160b478006 == 0 {
		__obf_f8c74f160b478006 = uint(0.7 + float64(__obf_cb2c55730faeb3f0.Bounds().Dy())/__obf_fd68fc877fd97a26)
	}

	// Trivial case: return input image
	if int(__obf_8367d6c856df0960) == __obf_cb2c55730faeb3f0.Bounds().Dx() && int(__obf_f8c74f160b478006) == __obf_cb2c55730faeb3f0.Bounds().Dy() {
		return __obf_cb2c55730faeb3f0
	}

	// Input image has no pixels
	if __obf_cb2c55730faeb3f0.Bounds().Dx() <= 0 || __obf_cb2c55730faeb3f0.Bounds().Dy() <= 0 {
		return __obf_cb2c55730faeb3f0
	}

	if __obf_f617084106645ca0 == NearestNeighbor {
		return __obf_65323ce6bf4e39dd(__obf_8367d6c856df0960, __obf_f8c74f160b478006, __obf_b72b7f9b028a1c41, __obf_fd68fc877fd97a26, __obf_cb2c55730faeb3f0, __obf_f617084106645ca0)
	}
	__obf_e7df7c4d1e822f47, __obf_25630250e4587e64 := __obf_f617084106645ca0.__obf_25630250e4587e64()
	__obf_bebf8ea905b0c76f := runtime.GOMAXPROCS(0)
	__obf_d51f58edd2167790 := sync.WaitGroup{}

	// Generic access to image.Image is slow in tight loops.
	// The optimal access has to be determined from the concrete image type.
	switch __obf_16d05d485a40050a := __obf_cb2c55730faeb3f0.(type) {
	case *image.RGBA:
		__obf_8dfc03881a81586c := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_bc71b3618bd8ea83(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_b2d4489701a105b7(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_bc71b3618bd8ea83(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_b2d4489701a105b7(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// 8-bit precision
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.NRGBA:
		__obf_8dfc03881a81586c := image.NewRGBA(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_bc71b3618bd8ea83(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_bebf8ea905b0c76f; __obf_b92089dde6a4a00b++ {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_e8e29c8f236bdc6f(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_bc71b3618bd8ea83(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_b2d4489701a105b7(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// 8-bit precision
					// accessing the YCbCr arrays in a tight loop is slow.
					// converting the image to ycc increases performance by 2x.
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad

	case *image.YCbCr:
		__obf_8dfc03881a81586c := __obf_f13836c808b60e84(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)), __obf_16d05d485a40050a.SubsampleRatio)
		__obf_ba974a8f7e5f9dad := __obf_f13836c808b60e84(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)), image.YCbCrSubsampleRatio444)
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_bc71b3618bd8ea83(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_a5eab9a5623e4978 := __obf_25f07534b3928897(__obf_16d05d485a40050a)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_bebf8ea905b0c76f; __obf_b92089dde6a4a00b++ {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*__obf_7b573c09d347a918)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_65afc53989013e4c(__obf_a5eab9a5623e4978, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_bc71b3618bd8ea83(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_bebf8ea905b0c76f; __obf_b92089dde6a4a00b++ {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*__obf_7b573c09d347a918)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_65afc53989013e4c(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad.YCbCr()
	case *image.RGBA64:
		__obf_8dfc03881a81586c := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA64(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_7e80e77b51d72f60(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_dd75986f2d7d6f93(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_7e80e77b51d72f60(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_dd75986f2d7d6f93(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// 16-bit precision
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.NRGBA64:
		__obf_8dfc03881a81586c := image.NewRGBA64(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA64(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_7e80e77b51d72f60(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_7dc22bccd4944b8d(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_7e80e77b51d72f60(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_dd75986f2d7d6f93(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// 8-bit precision
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.Gray:
		__obf_8dfc03881a81586c := image.NewGray(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewGray(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_bc71b3618bd8ea83(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_bebf8ea905b0c76f; __obf_b92089dde6a4a00b++ {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_e05bbe664e33033f(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_bc71b3618bd8ea83(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_e05bbe664e33033f(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// 16-bit precision
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.Gray16:
		__obf_8dfc03881a81586c := image.NewGray16(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewGray16(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_7e80e77b51d72f60(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray16)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_909dd2d677e37880(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_7e80e77b51d72f60(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray16)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_909dd2d677e37880(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// 16-bit precision
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	default:
		__obf_8dfc03881a81586c := image.NewRGBA64(image.Rect(0, 0, __obf_cb2c55730faeb3f0.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA64(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_7e80e77b51d72f60(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_ccde39b12bafba9e(__obf_cb2c55730faeb3f0, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_7e80e77b51d72f60(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26, __obf_25630250e4587e64)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_dd75986f2d7d6f93(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	}
}

func __obf_65323ce6bf4e39dd(__obf_8367d6c856df0960, __obf_f8c74f160b478006 uint, __obf_b72b7f9b028a1c41, __obf_fd68fc877fd97a26 float64, __obf_cb2c55730faeb3f0 image.Image, __obf_f617084106645ca0 InterpolationFunction) image.Image {
	__obf_e7df7c4d1e822f47, _ := __obf_f617084106645ca0.__obf_25630250e4587e64()
	__obf_bebf8ea905b0c76f := runtime.GOMAXPROCS(0)
	__obf_d51f58edd2167790 := sync.WaitGroup{}

	switch __obf_16d05d485a40050a := __obf_cb2c55730faeb3f0.(type) {
	case *image.RGBA:
		__obf_8dfc03881a81586c := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := 0; __obf_b92089dde6a4a00b < __obf_bebf8ea905b0c76f; __obf_b92089dde6a4a00b++ {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_6f78347896b0579d(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// horizontal filter on transposed image, result is not transposed
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_6f78347896b0579d(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.NRGBA:
		__obf_8dfc03881a81586c := // 8-bit precision
			image.NewNRGBA(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewNRGBA(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.NRGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_e973cc62a48a6a19(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// horizontal filter on transposed image, result is not transposed
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.NRGBA)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_e973cc62a48a6a19(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.YCbCr:
		__obf_8dfc03881a81586c := // 8-bit precision
			// accessing the YCbCr arrays in a tight loop is slow.
			// converting the image to ycc increases performance by 2x.
			__obf_f13836c808b60e84(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)), __obf_16d05d485a40050a.SubsampleRatio)
		__obf_ba974a8f7e5f9dad := __obf_f13836c808b60e84(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)), image.YCbCrSubsampleRatio444)
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_a5eab9a5623e4978 := __obf_25f07534b3928897(__obf_16d05d485a40050a)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*__obf_7b573c09d347a918)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_dee53dca87365a2a(__obf_a5eab9a5623e4978, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*__obf_7b573c09d347a918)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_dee53dca87365a2a(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_fd68fc877fd97a26, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad.YCbCr()
	case *image.RGBA64:
		__obf_8dfc03881a81586c := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA64(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_fc257f8bf8c94d34(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// horizontal filter on transposed image, result is not transposed
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_fc257f8bf8c94d34(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.NRGBA64:
		__obf_8dfc03881a81586c := // 16-bit precision
			image.NewNRGBA64(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewNRGBA64(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.NRGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_0598eeca293592f2(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// horizontal filter on transposed image, result is not transposed
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.NRGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_0598eeca293592f2(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.Gray:
		__obf_8dfc03881a81586c := // 8-bit precision
			image.NewGray(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewGray(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_271bfd0ed76a5978(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// horizontal filter on transposed image, result is not transposed
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_271bfd0ed76a5978(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	case *image.Gray16:
		__obf_8dfc03881a81586c := // 16-bit precision
			image.NewGray16(image.Rect(0, 0, __obf_16d05d485a40050a.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewGray16(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray16)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_1d39fd992b17e818(__obf_16d05d485a40050a, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// horizontal filter on transposed image, result is not transposed
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.Gray16)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_1d39fd992b17e818(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	default:
		__obf_8dfc03881a81586c := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_cb2c55730faeb3f0.Bounds().Dy(), int(__obf_8367d6c856df0960)))
		__obf_ba974a8f7e5f9dad := image.NewRGBA64(image.Rect(0, 0, int(__obf_8367d6c856df0960), int(__obf_f8c74f160b478006)))
		__obf_ed4a0c7a1913ca57,

			// horizontal filter, results in transposed temporary image
			__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 := __obf_b0d4c44c11725e77(__obf_8dfc03881a81586c.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_b72b7f9b028a1c41)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_8dfc03881a81586c, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_226995aa8d962794(__obf_cb2c55730faeb3f0, __obf_70c504e4578e8140, __obf_b72b7f9b028a1c41, __obf_ed4a0c7a1913ca57,

					// horizontal filter on transposed image, result is not transposed
					__obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		__obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b, __obf_743fac2b8ed0fc32 = __obf_b0d4c44c11725e77(__obf_ba974a8f7e5f9dad.Bounds().Dy(), __obf_e7df7c4d1e822f47, __obf_ce79a531ca9b78ac, __obf_fd68fc877fd97a26)
		__obf_d51f58edd2167790.
			Add(__obf_bebf8ea905b0c76f)
		for __obf_b92089dde6a4a00b := range __obf_bebf8ea905b0c76f {
			__obf_70c504e4578e8140 := __obf_663b78ff95859c03(__obf_ba974a8f7e5f9dad, __obf_b92089dde6a4a00b, __obf_bebf8ea905b0c76f).(*image.RGBA64)
			go func() {
				defer __obf_d51f58edd2167790.Done()
				__obf_fc257f8bf8c94d34(__obf_8dfc03881a81586c, __obf_70c504e4578e8140, __obf_ed4a0c7a1913ca57, __obf_3a86ac65e5f0ae4b,

					// Calculates scaling factors using old and new image dimensions.
					__obf_743fac2b8ed0fc32)
			}()
		}
		__obf_d51f58edd2167790.
			Wait()
		return __obf_ba974a8f7e5f9dad
	}

}

func __obf_ef7f92af091e9888(__obf_8367d6c856df0960, __obf_f8c74f160b478006 uint, __obf_e5960732385e1d00, __obf_25ee552957837d75 float64) (__obf_b72b7f9b028a1c41, __obf_fd68fc877fd97a26 float64) {
	if __obf_8367d6c856df0960 == 0 {
		if __obf_f8c74f160b478006 == 0 {
			__obf_b72b7f9b028a1c41 = 1.0
			__obf_fd68fc877fd97a26 = 1.0
		} else {
			__obf_fd68fc877fd97a26 = __obf_25ee552957837d75 / float64(__obf_f8c74f160b478006)
			__obf_b72b7f9b028a1c41 = __obf_fd68fc877fd97a26
		}
	} else {
		__obf_b72b7f9b028a1c41 = __obf_e5960732385e1d00 / float64(__obf_8367d6c856df0960)
		if __obf_f8c74f160b478006 == 0 {
			__obf_fd68fc877fd97a26 = __obf_b72b7f9b028a1c41
		} else {
			__obf_fd68fc877fd97a26 = __obf_25ee552957837d75 / float64(__obf_f8c74f160b478006)
		}
	}
	return
}

type __obf_96cdc03832570ed3 interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func __obf_663b78ff95859c03(__obf_cb2c55730faeb3f0 __obf_96cdc03832570ed3, __obf_b92089dde6a4a00b, __obf_fa6037c3a2eeaada int) image.Image {
	return __obf_cb2c55730faeb3f0.SubImage(image.Rect(__obf_cb2c55730faeb3f0.Bounds().Min.X, __obf_cb2c55730faeb3f0.Bounds().Min.Y+__obf_b92089dde6a4a00b*__obf_cb2c55730faeb3f0.Bounds().Dy()/__obf_fa6037c3a2eeaada, __obf_cb2c55730faeb3f0.Bounds().Max.X, __obf_cb2c55730faeb3f0.Bounds().Min.Y+(__obf_b92089dde6a4a00b+1)*__obf_cb2c55730faeb3f0.Bounds().Dy()/__obf_fa6037c3a2eeaada))
}
