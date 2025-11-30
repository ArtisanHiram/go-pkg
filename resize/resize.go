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
package __obf_ac510735d4d13cdd

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
func (__obf_7326873273fab46e InterpolationFunction) __obf_0a93ef33b92d5d34() (int, func(float64) float64) {
	switch __obf_7326873273fab46e {
	case Bilinear:
		return 2, __obf_75f6718d0c68ce45
	case Bicubic:
		return 4, __obf_6dbe12e595f20b30
	case MitchellNetravali:
		return 4, __obf_4f5e93c03ec56195
	case Lanczos2:
		return 4, __obf_55ba6e33e6203680
	case Lanczos3:
		return 6, __obf_bc625d1a66a5919f
	default:
		// Default to NearestNeighbor.
		return 2, __obf_6a3a721569bbf375
	}
}

// values <1 will sharpen the image
var __obf_b5313b280a820ad0 = 1.0

// Resize scales an image to new width and height using the interpolation function interp.
// A new image with the given dimensions will be returned.
// If one of the parameters width or height is set to 0, its size will be calculated so that
// the aspect ratio is that of the originating image.
// The resizing algorithm uses channels for parallel computation.
// If the input image has width or height of 0, it is returned unchanged.
func Resize(__obf_bc571956a33c8293, __obf_5a5a28ca143a2f3f uint, __obf_4f5a8d0b34527afc image.Image, __obf_00264b26e01a01f0 InterpolationFunction) image.Image {
	__obf_d5b7dea7a78fd48e, __obf_cbfadbd3718a054b := __obf_46b5846e696e2e9f(__obf_bc571956a33c8293, __obf_5a5a28ca143a2f3f, float64(__obf_4f5a8d0b34527afc.Bounds().Dx()), float64(__obf_4f5a8d0b34527afc.Bounds().Dy()))
	if __obf_bc571956a33c8293 == 0 {
		__obf_bc571956a33c8293 = uint(0.7 + float64(__obf_4f5a8d0b34527afc.Bounds().Dx())/__obf_d5b7dea7a78fd48e)
	}
	if __obf_5a5a28ca143a2f3f == 0 {
		__obf_5a5a28ca143a2f3f = uint(0.7 + float64(__obf_4f5a8d0b34527afc.Bounds().Dy())/__obf_cbfadbd3718a054b)
	}

	// Trivial case: return input image
	if int(__obf_bc571956a33c8293) == __obf_4f5a8d0b34527afc.Bounds().Dx() && int(__obf_5a5a28ca143a2f3f) == __obf_4f5a8d0b34527afc.Bounds().Dy() {
		return __obf_4f5a8d0b34527afc
	}

	// Input image has no pixels
	if __obf_4f5a8d0b34527afc.Bounds().Dx() <= 0 || __obf_4f5a8d0b34527afc.Bounds().Dy() <= 0 {
		return __obf_4f5a8d0b34527afc
	}

	if __obf_00264b26e01a01f0 == NearestNeighbor {
		return __obf_fb6a667917811534(__obf_bc571956a33c8293, __obf_5a5a28ca143a2f3f, __obf_d5b7dea7a78fd48e, __obf_cbfadbd3718a054b, __obf_4f5a8d0b34527afc, __obf_00264b26e01a01f0)
	}
	__obf_d028ed8a4cd5fa47, __obf_0a93ef33b92d5d34 := __obf_00264b26e01a01f0.__obf_0a93ef33b92d5d34()
	__obf_47cc8a0a92b92912 := runtime.GOMAXPROCS(0)
	__obf_cd504da1ae38dec4 := sync.WaitGroup{}

	// Generic access to image.Image is slow in tight loops.
	// The optimal access has to be determined from the concrete image type.
	switch __obf_e40229e8249ca178 := __obf_4f5a8d0b34527afc.(type) {
	case *image.RGBA:
		__obf_30b4a81a5645f273 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_ce2a9e53779a9624(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_a81ef026a03961f8(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_ce2a9e53779a9624(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_a81ef026a03961f8(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// 8-bit precision
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.NRGBA:
		__obf_30b4a81a5645f273 := image.NewRGBA(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_ce2a9e53779a9624(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_47cc8a0a92b92912; __obf_7326873273fab46e++ {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_e01bcd08df027a03(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_ce2a9e53779a9624(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_a81ef026a03961f8(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// 8-bit precision
					// accessing the YCbCr arrays in a tight loop is slow.
					// converting the image to ycc increases performance by 2x.
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7

	case *image.YCbCr:
		__obf_30b4a81a5645f273 := __obf_57395a4603262f46(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)), __obf_e40229e8249ca178.SubsampleRatio)
		__obf_c27c5a7dd11abfd7 := __obf_57395a4603262f46(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)), image.YCbCrSubsampleRatio444)
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_ce2a9e53779a9624(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_3ccd19e51488be10 := __obf_76286d70dcc66ea1(__obf_e40229e8249ca178)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_47cc8a0a92b92912; __obf_7326873273fab46e++ {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*__obf_9474021cedba6d1a)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_76c3d32722dd59f4(__obf_3ccd19e51488be10, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_ce2a9e53779a9624(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_47cc8a0a92b92912; __obf_7326873273fab46e++ {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*__obf_9474021cedba6d1a)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_76c3d32722dd59f4(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7.YCbCr()
	case *image.RGBA64:
		__obf_30b4a81a5645f273 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA64(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_9e73f2386e0e3872(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_fd9e808f1958058e(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_9e73f2386e0e3872(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_fd9e808f1958058e(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// 16-bit precision
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.NRGBA64:
		__obf_30b4a81a5645f273 := image.NewRGBA64(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA64(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_9e73f2386e0e3872(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_3d8976013120b419(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_9e73f2386e0e3872(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_fd9e808f1958058e(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// 8-bit precision
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.Gray:
		__obf_30b4a81a5645f273 := image.NewGray(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewGray(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_ce2a9e53779a9624(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_47cc8a0a92b92912; __obf_7326873273fab46e++ {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_e7b82bf93376b1f1(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_ce2a9e53779a9624(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_e7b82bf93376b1f1(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// 16-bit precision
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.Gray16:
		__obf_30b4a81a5645f273 := image.NewGray16(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewGray16(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_9e73f2386e0e3872(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray16)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_88dc9a250ceb45c4(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_9e73f2386e0e3872(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray16)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_88dc9a250ceb45c4(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// 16-bit precision
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	default:
		__obf_30b4a81a5645f273 := image.NewRGBA64(image.Rect(0, 0, __obf_4f5a8d0b34527afc.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA64(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_9e73f2386e0e3872(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_1e8dec540911dfe5(__obf_4f5a8d0b34527afc, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_9e73f2386e0e3872(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b, __obf_0a93ef33b92d5d34)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_fd9e808f1958058e(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	}
}

func __obf_fb6a667917811534(__obf_bc571956a33c8293, __obf_5a5a28ca143a2f3f uint, __obf_d5b7dea7a78fd48e, __obf_cbfadbd3718a054b float64, __obf_4f5a8d0b34527afc image.Image, __obf_00264b26e01a01f0 InterpolationFunction) image.Image {
	__obf_d028ed8a4cd5fa47, _ := __obf_00264b26e01a01f0.__obf_0a93ef33b92d5d34()
	__obf_47cc8a0a92b92912 := runtime.GOMAXPROCS(0)
	__obf_cd504da1ae38dec4 := sync.WaitGroup{}

	switch __obf_e40229e8249ca178 := __obf_4f5a8d0b34527afc.(type) {
	case *image.RGBA:
		__obf_30b4a81a5645f273 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := 0; __obf_7326873273fab46e < __obf_47cc8a0a92b92912; __obf_7326873273fab46e++ {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_9dfa6bb18b3e659e(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// horizontal filter on transposed image, result is not transposed
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_9dfa6bb18b3e659e(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.NRGBA:
		__obf_30b4a81a5645f273 := // 8-bit precision
			image.NewNRGBA(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewNRGBA(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.NRGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_d3929eb1ff5d9a81(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// horizontal filter on transposed image, result is not transposed
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.NRGBA)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_d3929eb1ff5d9a81(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.YCbCr:
		__obf_30b4a81a5645f273 := // 8-bit precision
			// accessing the YCbCr arrays in a tight loop is slow.
			// converting the image to ycc increases performance by 2x.
			__obf_57395a4603262f46(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)), __obf_e40229e8249ca178.SubsampleRatio)
		__obf_c27c5a7dd11abfd7 := __obf_57395a4603262f46(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)), image.YCbCrSubsampleRatio444)
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_3ccd19e51488be10 := __obf_76286d70dcc66ea1(__obf_e40229e8249ca178)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*__obf_9474021cedba6d1a)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_cb273764563e26fc(__obf_3ccd19e51488be10, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*__obf_9474021cedba6d1a)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_cb273764563e26fc(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_cbfadbd3718a054b, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7.YCbCr()
	case *image.RGBA64:
		__obf_30b4a81a5645f273 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA64(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_e7f8b39bc0d4199a(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// horizontal filter on transposed image, result is not transposed
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_e7f8b39bc0d4199a(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.NRGBA64:
		__obf_30b4a81a5645f273 := // 16-bit precision
			image.NewNRGBA64(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewNRGBA64(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.NRGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_2cf1e988bdf22f40(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// horizontal filter on transposed image, result is not transposed
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.NRGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_2cf1e988bdf22f40(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.Gray:
		__obf_30b4a81a5645f273 := // 8-bit precision
			image.NewGray(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewGray(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_220524a8a738e132(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// horizontal filter on transposed image, result is not transposed
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_220524a8a738e132(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	case *image.Gray16:
		__obf_30b4a81a5645f273 := // 16-bit precision
			image.NewGray16(image.Rect(0, 0, __obf_e40229e8249ca178.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewGray16(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray16)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_b46a659aaf06ae03(__obf_e40229e8249ca178, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// horizontal filter on transposed image, result is not transposed
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.Gray16)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_b46a659aaf06ae03(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	default:
		__obf_30b4a81a5645f273 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_4f5a8d0b34527afc.Bounds().Dy(), int(__obf_bc571956a33c8293)))
		__obf_c27c5a7dd11abfd7 := image.NewRGBA64(image.Rect(0, 0, int(__obf_bc571956a33c8293), int(__obf_5a5a28ca143a2f3f)))
		__obf_46e0cb0caad4034a,

			// horizontal filter, results in transposed temporary image
			__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d := __obf_b8bb0d275d8b4e0d(__obf_30b4a81a5645f273.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_d5b7dea7a78fd48e)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_30b4a81a5645f273, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_d4c3da6972dfb241(__obf_4f5a8d0b34527afc, __obf_dc36ea2ab56fb609, __obf_d5b7dea7a78fd48e, __obf_46e0cb0caad4034a,

					// horizontal filter on transposed image, result is not transposed
					__obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		__obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b, __obf_052b80a7daaa449d = __obf_b8bb0d275d8b4e0d(__obf_c27c5a7dd11abfd7.Bounds().Dy(), __obf_d028ed8a4cd5fa47, __obf_b5313b280a820ad0, __obf_cbfadbd3718a054b)
		__obf_cd504da1ae38dec4.
			Add(__obf_47cc8a0a92b92912)
		for __obf_7326873273fab46e := range __obf_47cc8a0a92b92912 {
			__obf_dc36ea2ab56fb609 := __obf_0a10803f6145f736(__obf_c27c5a7dd11abfd7, __obf_7326873273fab46e, __obf_47cc8a0a92b92912).(*image.RGBA64)
			go func() {
				defer __obf_cd504da1ae38dec4.Done()
				__obf_e7f8b39bc0d4199a(__obf_30b4a81a5645f273, __obf_dc36ea2ab56fb609, __obf_46e0cb0caad4034a, __obf_cf2c82255bf7ee2b,

					// Calculates scaling factors using old and new image dimensions.
					__obf_052b80a7daaa449d)
			}()
		}
		__obf_cd504da1ae38dec4.
			Wait()
		return __obf_c27c5a7dd11abfd7
	}

}

func __obf_46b5846e696e2e9f(__obf_bc571956a33c8293, __obf_5a5a28ca143a2f3f uint, __obf_8877583a5ecc7d86, __obf_6c3c4bd8256403ef float64) (__obf_d5b7dea7a78fd48e, __obf_cbfadbd3718a054b float64) {
	if __obf_bc571956a33c8293 == 0 {
		if __obf_5a5a28ca143a2f3f == 0 {
			__obf_d5b7dea7a78fd48e = 1.0
			__obf_cbfadbd3718a054b = 1.0
		} else {
			__obf_cbfadbd3718a054b = __obf_6c3c4bd8256403ef / float64(__obf_5a5a28ca143a2f3f)
			__obf_d5b7dea7a78fd48e = __obf_cbfadbd3718a054b
		}
	} else {
		__obf_d5b7dea7a78fd48e = __obf_8877583a5ecc7d86 / float64(__obf_bc571956a33c8293)
		if __obf_5a5a28ca143a2f3f == 0 {
			__obf_cbfadbd3718a054b = __obf_d5b7dea7a78fd48e
		} else {
			__obf_cbfadbd3718a054b = __obf_6c3c4bd8256403ef / float64(__obf_5a5a28ca143a2f3f)
		}
	}
	return
}

type __obf_400bbffddada8b3f interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func __obf_0a10803f6145f736(__obf_4f5a8d0b34527afc __obf_400bbffddada8b3f, __obf_7326873273fab46e, __obf_ba0c9781114791b5 int) image.Image {
	return __obf_4f5a8d0b34527afc.SubImage(image.Rect(__obf_4f5a8d0b34527afc.Bounds().Min.X, __obf_4f5a8d0b34527afc.Bounds().Min.Y+__obf_7326873273fab46e*__obf_4f5a8d0b34527afc.Bounds().Dy()/__obf_ba0c9781114791b5, __obf_4f5a8d0b34527afc.Bounds().Max.X, __obf_4f5a8d0b34527afc.Bounds().Min.Y+(__obf_7326873273fab46e+1)*__obf_4f5a8d0b34527afc.Bounds().Dy()/__obf_ba0c9781114791b5))
}
