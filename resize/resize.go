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
package __obf_71f4605f6400e2ce

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
func (__obf_7d43b40e55127c49 InterpolationFunction) __obf_2fde8fe3fe7a0697() (int, func(float64) float64) {
	switch __obf_7d43b40e55127c49 {
	case Bilinear:
		return 2, __obf_995190fb5ae77946
	case Bicubic:
		return 4, __obf_c0b2553e410838b3
	case MitchellNetravali:
		return 4, __obf_52c4a4763619b896
	case Lanczos2:
		return 4, __obf_f2fccd76139d1121
	case Lanczos3:
		return 6, __obf_e214463bbc8b0b56
	default:
		// Default to NearestNeighbor.
		return 2, __obf_d8b0075bb092e33d
	}
}

// values <1 will sharpen the image
var __obf_940ae86fccad7783 = 1.0

// Resize scales an image to new width and height using the interpolation function interp.
// A new image with the given dimensions will be returned.
// If one of the parameters width or height is set to 0, its size will be calculated so that
// the aspect ratio is that of the originating image.
// The resizing algorithm uses channels for parallel computation.
// If the input image has width or height of 0, it is returned unchanged.
func Resize(__obf_2270f271c29c803d, __obf_c0a7e6ac00747c51 uint, __obf_0b888bcf5266913b image.Image, __obf_a66df7eedc4b118a InterpolationFunction) image.Image {
	__obf_4fcccfd3fdc0059d, __obf_364ce72b6402e431 := __obf_b043fb79a6823327(__obf_2270f271c29c803d, __obf_c0a7e6ac00747c51, float64(__obf_0b888bcf5266913b.Bounds().Dx()), float64(__obf_0b888bcf5266913b.Bounds().Dy()))
	if __obf_2270f271c29c803d == 0 {
		__obf_2270f271c29c803d = uint(0.7 + float64(__obf_0b888bcf5266913b.Bounds().Dx())/__obf_4fcccfd3fdc0059d)
	}
	if __obf_c0a7e6ac00747c51 == 0 {
		__obf_c0a7e6ac00747c51 = uint(0.7 + float64(__obf_0b888bcf5266913b.Bounds().Dy())/__obf_364ce72b6402e431)
	}

	// Trivial case: return input image
	if int(__obf_2270f271c29c803d) == __obf_0b888bcf5266913b.Bounds().Dx() && int(__obf_c0a7e6ac00747c51) == __obf_0b888bcf5266913b.Bounds().Dy() {
		return __obf_0b888bcf5266913b
	}

	// Input image has no pixels
	if __obf_0b888bcf5266913b.Bounds().Dx() <= 0 || __obf_0b888bcf5266913b.Bounds().Dy() <= 0 {
		return __obf_0b888bcf5266913b
	}

	if __obf_a66df7eedc4b118a == NearestNeighbor {
		return __obf_7afb035bf6a7a01a(__obf_2270f271c29c803d, __obf_c0a7e6ac00747c51, __obf_4fcccfd3fdc0059d, __obf_364ce72b6402e431, __obf_0b888bcf5266913b, __obf_a66df7eedc4b118a)
	}
	__obf_cbece4382a004042, __obf_2fde8fe3fe7a0697 := __obf_a66df7eedc4b118a.__obf_2fde8fe3fe7a0697()
	__obf_5cd92652c9267242 := runtime.GOMAXPROCS(0)
	__obf_4e676684903f7719 := sync.WaitGroup{}

	// Generic access to image.Image is slow in tight loops.
	// The optimal access has to be determined from the concrete image type.
	switch __obf_25c0ecf52c2bfea4 := __obf_0b888bcf5266913b.(type) {
	case *image.RGBA:
		__obf_86d7893cee7d50d6 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_fc245964fca32498(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_63620e18c66f5397(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_fc245964fca32498(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_63620e18c66f5397(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// 8-bit precision
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.NRGBA:
		__obf_86d7893cee7d50d6 := image.NewRGBA(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_fc245964fca32498(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_5cd92652c9267242; __obf_7d43b40e55127c49++ {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_b6a8709004d14349(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_fc245964fca32498(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_63620e18c66f5397(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// 8-bit precision
					// accessing the YCbCr arrays in a tight loop is slow.
					// converting the image to ycc increases performance by 2x.
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b

	case *image.YCbCr:
		__obf_86d7893cee7d50d6 := __obf_1b258dd421287100(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)), __obf_25c0ecf52c2bfea4.SubsampleRatio)
		__obf_6ff3da8be26cb26b := __obf_1b258dd421287100(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)), image.YCbCrSubsampleRatio444)
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_fc245964fca32498(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_b13d16274eb954e2 := __obf_fcb9eb75f0392749(__obf_25c0ecf52c2bfea4)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_5cd92652c9267242; __obf_7d43b40e55127c49++ {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*__obf_786a8b420048cd65)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_22b27a29fbaf62b4(__obf_b13d16274eb954e2, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_fc245964fca32498(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_5cd92652c9267242; __obf_7d43b40e55127c49++ {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*__obf_786a8b420048cd65)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_22b27a29fbaf62b4(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b.YCbCr()
	case *image.RGBA64:
		__obf_86d7893cee7d50d6 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA64(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0bc5ed636cd197d7(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_b799c14ff2666c0a(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0bc5ed636cd197d7(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_b799c14ff2666c0a(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// 16-bit precision
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.NRGBA64:
		__obf_86d7893cee7d50d6 := image.NewRGBA64(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA64(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0bc5ed636cd197d7(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_1ab7584b5da9bac1(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0bc5ed636cd197d7(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_b799c14ff2666c0a(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// 8-bit precision
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.Gray:
		__obf_86d7893cee7d50d6 := image.NewGray(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewGray(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_fc245964fca32498(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_5cd92652c9267242; __obf_7d43b40e55127c49++ {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_5727e96fbc4064ea(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_fc245964fca32498(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_5727e96fbc4064ea(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// 16-bit precision
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.Gray16:
		__obf_86d7893cee7d50d6 := image.NewGray16(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewGray16(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0bc5ed636cd197d7(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray16)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_723b7c23ca44d438(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0bc5ed636cd197d7(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray16)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_723b7c23ca44d438(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// 16-bit precision
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	default:
		__obf_86d7893cee7d50d6 := image.NewRGBA64(image.Rect(0, 0, __obf_0b888bcf5266913b.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA64(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0bc5ed636cd197d7(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_ca5fdfe67ed70e0d(__obf_0b888bcf5266913b, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0bc5ed636cd197d7(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431, __obf_2fde8fe3fe7a0697)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_b799c14ff2666c0a(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	}
}

func __obf_7afb035bf6a7a01a(__obf_2270f271c29c803d, __obf_c0a7e6ac00747c51 uint, __obf_4fcccfd3fdc0059d, __obf_364ce72b6402e431 float64, __obf_0b888bcf5266913b image.Image, __obf_a66df7eedc4b118a InterpolationFunction) image.Image {
	__obf_cbece4382a004042, _ := __obf_a66df7eedc4b118a.__obf_2fde8fe3fe7a0697()
	__obf_5cd92652c9267242 := runtime.GOMAXPROCS(0)
	__obf_4e676684903f7719 := sync.WaitGroup{}

	switch __obf_25c0ecf52c2bfea4 := __obf_0b888bcf5266913b.(type) {
	case *image.RGBA:
		__obf_86d7893cee7d50d6 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := 0; __obf_7d43b40e55127c49 < __obf_5cd92652c9267242; __obf_7d43b40e55127c49++ {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_66d1834bc9607e4b(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// horizontal filter on transposed image, result is not transposed
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_66d1834bc9607e4b(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.NRGBA:
		__obf_86d7893cee7d50d6 := // 8-bit precision
			image.NewNRGBA(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewNRGBA(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.NRGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_fe4127dc1da8788e(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// horizontal filter on transposed image, result is not transposed
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.NRGBA)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_fe4127dc1da8788e(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.YCbCr:
		__obf_86d7893cee7d50d6 := // 8-bit precision
			// accessing the YCbCr arrays in a tight loop is slow.
			// converting the image to ycc increases performance by 2x.
			__obf_1b258dd421287100(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)), __obf_25c0ecf52c2bfea4.SubsampleRatio)
		__obf_6ff3da8be26cb26b := __obf_1b258dd421287100(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)), image.YCbCrSubsampleRatio444)
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_b13d16274eb954e2 := __obf_fcb9eb75f0392749(__obf_25c0ecf52c2bfea4)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*__obf_786a8b420048cd65)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_61f66fd475da9f7c(__obf_b13d16274eb954e2, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*__obf_786a8b420048cd65)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_61f66fd475da9f7c(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_364ce72b6402e431, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b.YCbCr()
	case *image.RGBA64:
		__obf_86d7893cee7d50d6 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA64(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_228913032261ceca(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// horizontal filter on transposed image, result is not transposed
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_228913032261ceca(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.NRGBA64:
		__obf_86d7893cee7d50d6 := // 16-bit precision
			image.NewNRGBA64(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewNRGBA64(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.NRGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_dc6b83b67e049b85(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// horizontal filter on transposed image, result is not transposed
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.NRGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_dc6b83b67e049b85(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.Gray:
		__obf_86d7893cee7d50d6 := // 8-bit precision
			image.NewGray(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewGray(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_0ca2f5762875ef96(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// horizontal filter on transposed image, result is not transposed
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_0ca2f5762875ef96(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	case *image.Gray16:
		__obf_86d7893cee7d50d6 := // 16-bit precision
			image.NewGray16(image.Rect(0, 0, __obf_25c0ecf52c2bfea4.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewGray16(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray16)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_525f68f0384d6897(__obf_25c0ecf52c2bfea4, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// horizontal filter on transposed image, result is not transposed
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.Gray16)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_525f68f0384d6897(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	default:
		__obf_86d7893cee7d50d6 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_0b888bcf5266913b.Bounds().Dy(), int(__obf_2270f271c29c803d)))
		__obf_6ff3da8be26cb26b := image.NewRGBA64(image.Rect(0, 0, int(__obf_2270f271c29c803d), int(__obf_c0a7e6ac00747c51)))
		__obf_73fc3de4b87d9301,

			// horizontal filter, results in transposed temporary image
			__obf_87dd30f656544b12, __obf_6a27c875a42baeab := __obf_0b05d823d55af64a(__obf_86d7893cee7d50d6.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_4fcccfd3fdc0059d)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_86d7893cee7d50d6, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_5d671bd01e4ff610(__obf_0b888bcf5266913b, __obf_6bdba559ac42ebb6, __obf_4fcccfd3fdc0059d, __obf_73fc3de4b87d9301,

					// horizontal filter on transposed image, result is not transposed
					__obf_87dd30f656544b12, __obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		__obf_73fc3de4b87d9301, __obf_87dd30f656544b12, __obf_6a27c875a42baeab = __obf_0b05d823d55af64a(__obf_6ff3da8be26cb26b.Bounds().Dy(), __obf_cbece4382a004042, __obf_940ae86fccad7783, __obf_364ce72b6402e431)
		__obf_4e676684903f7719.
			Add(__obf_5cd92652c9267242)
		for __obf_7d43b40e55127c49 := range __obf_5cd92652c9267242 {
			__obf_6bdba559ac42ebb6 := __obf_fd35f8fae09df176(__obf_6ff3da8be26cb26b, __obf_7d43b40e55127c49, __obf_5cd92652c9267242).(*image.RGBA64)
			go func() {
				defer __obf_4e676684903f7719.Done()
				__obf_228913032261ceca(__obf_86d7893cee7d50d6, __obf_6bdba559ac42ebb6, __obf_73fc3de4b87d9301, __obf_87dd30f656544b12,

					// Calculates scaling factors using old and new image dimensions.
					__obf_6a27c875a42baeab)
			}()
		}
		__obf_4e676684903f7719.
			Wait()
		return __obf_6ff3da8be26cb26b
	}

}

func __obf_b043fb79a6823327(__obf_2270f271c29c803d, __obf_c0a7e6ac00747c51 uint, __obf_7c6dcadd6029cf9d, __obf_444c79c43a6493ef float64) (__obf_4fcccfd3fdc0059d, __obf_364ce72b6402e431 float64) {
	if __obf_2270f271c29c803d == 0 {
		if __obf_c0a7e6ac00747c51 == 0 {
			__obf_4fcccfd3fdc0059d = 1.0
			__obf_364ce72b6402e431 = 1.0
		} else {
			__obf_364ce72b6402e431 = __obf_444c79c43a6493ef / float64(__obf_c0a7e6ac00747c51)
			__obf_4fcccfd3fdc0059d = __obf_364ce72b6402e431
		}
	} else {
		__obf_4fcccfd3fdc0059d = __obf_7c6dcadd6029cf9d / float64(__obf_2270f271c29c803d)
		if __obf_c0a7e6ac00747c51 == 0 {
			__obf_364ce72b6402e431 = __obf_4fcccfd3fdc0059d
		} else {
			__obf_364ce72b6402e431 = __obf_444c79c43a6493ef / float64(__obf_c0a7e6ac00747c51)
		}
	}
	return
}

type __obf_ec4e0af1084a4a4e interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func __obf_fd35f8fae09df176(__obf_0b888bcf5266913b __obf_ec4e0af1084a4a4e, __obf_7d43b40e55127c49, __obf_9ff0618a89aa686a int) image.Image {
	return __obf_0b888bcf5266913b.SubImage(image.Rect(__obf_0b888bcf5266913b.Bounds().Min.X, __obf_0b888bcf5266913b.Bounds().Min.Y+__obf_7d43b40e55127c49*__obf_0b888bcf5266913b.Bounds().Dy()/__obf_9ff0618a89aa686a, __obf_0b888bcf5266913b.Bounds().Max.X, __obf_0b888bcf5266913b.Bounds().Min.Y+(__obf_7d43b40e55127c49+1)*__obf_0b888bcf5266913b.Bounds().Dy()/__obf_9ff0618a89aa686a))
}
