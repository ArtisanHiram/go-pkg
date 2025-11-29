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
package __obf_3858dbfa2ccfdbe9

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
func (__obf_9d9681e93d5a567f InterpolationFunction) __obf_e0303ccb3cff6d43() (int, func(float64) float64) {
	switch __obf_9d9681e93d5a567f {
	case Bilinear:
		return 2, __obf_29ef71619471455c
	case Bicubic:
		return 4, __obf_9b5fddd514be8a29
	case MitchellNetravali:
		return 4, __obf_0b97bf31dd3b88bb
	case Lanczos2:
		return 4, __obf_18ba59002a602c1c
	case Lanczos3:
		return 6, __obf_d3ba12b86735c961
	default:
		// Default to NearestNeighbor.
		return 2, __obf_ea38e93055f919d7
	}
}

// values <1 will sharpen the image
var __obf_74fbaffb42caf3d6 = 1.0

// Resize scales an image to new width and height using the interpolation function interp.
// A new image with the given dimensions will be returned.
// If one of the parameters width or height is set to 0, its size will be calculated so that
// the aspect ratio is that of the originating image.
// The resizing algorithm uses channels for parallel computation.
// If the input image has width or height of 0, it is returned unchanged.
func Resize(__obf_ba0d63dbead1ced9, __obf_c4381a5b61d233ea uint, __obf_93755cc657d3c1d7 image.Image, __obf_ebcc3d09dc09d13d InterpolationFunction) image.Image {
	__obf_bb203b2e26665a43, __obf_f464c7a8df992c8a := __obf_1da7c3c6d2fcdf43(__obf_ba0d63dbead1ced9, __obf_c4381a5b61d233ea, float64(__obf_93755cc657d3c1d7.Bounds().Dx()), float64(__obf_93755cc657d3c1d7.Bounds().Dy()))
	if __obf_ba0d63dbead1ced9 == 0 {
		__obf_ba0d63dbead1ced9 = uint(0.7 + float64(__obf_93755cc657d3c1d7.Bounds().Dx())/__obf_bb203b2e26665a43)
	}
	if __obf_c4381a5b61d233ea == 0 {
		__obf_c4381a5b61d233ea = uint(0.7 + float64(__obf_93755cc657d3c1d7.Bounds().Dy())/__obf_f464c7a8df992c8a)
	}

	// Trivial case: return input image
	if int(__obf_ba0d63dbead1ced9) == __obf_93755cc657d3c1d7.Bounds().Dx() && int(__obf_c4381a5b61d233ea) == __obf_93755cc657d3c1d7.Bounds().Dy() {
		return __obf_93755cc657d3c1d7
	}

	// Input image has no pixels
	if __obf_93755cc657d3c1d7.Bounds().Dx() <= 0 || __obf_93755cc657d3c1d7.Bounds().Dy() <= 0 {
		return __obf_93755cc657d3c1d7
	}

	if __obf_ebcc3d09dc09d13d == NearestNeighbor {
		return __obf_539928a5c1774b7b(__obf_ba0d63dbead1ced9, __obf_c4381a5b61d233ea, __obf_bb203b2e26665a43, __obf_f464c7a8df992c8a, __obf_93755cc657d3c1d7, __obf_ebcc3d09dc09d13d)
	}
	__obf_aacb5abca7c80175, __obf_e0303ccb3cff6d43 := __obf_ebcc3d09dc09d13d.__obf_e0303ccb3cff6d43()
	__obf_e37ee56dd8133d3f := runtime.GOMAXPROCS(0)
	__obf_c9026fcbe3b0f1fd := sync.WaitGroup{}

	// Generic access to image.Image is slow in tight loops.
	// The optimal access has to be determined from the concrete image type.
	switch __obf_b8394cc3da38f2d2 := __obf_93755cc657d3c1d7.(type) {
	case *image.RGBA:
		__obf_61b54db7046becb4 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_3a1983a2984d289a(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_9c2deddd470bf9c3(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_3a1983a2984d289a(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_9c2deddd470bf9c3(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// 8-bit precision
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.NRGBA:
		__obf_61b54db7046becb4 := image.NewRGBA(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_3a1983a2984d289a(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_e37ee56dd8133d3f; __obf_9d9681e93d5a567f++ {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_1ef8f3d1fec5a3c3(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_3a1983a2984d289a(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_9c2deddd470bf9c3(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// 8-bit precision
					// accessing the YCbCr arrays in a tight loop is slow.
					// converting the image to ycc increases performance by 2x.
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b

	case *image.YCbCr:
		__obf_61b54db7046becb4 := __obf_8decfa8633084a6c(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)), __obf_b8394cc3da38f2d2.SubsampleRatio)
		__obf_cb722e41a1b9d39b := __obf_8decfa8633084a6c(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)), image.YCbCrSubsampleRatio444)
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_3a1983a2984d289a(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_8da81bf3c5df7f0d := __obf_6235574c4a9792da(__obf_b8394cc3da38f2d2)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_e37ee56dd8133d3f; __obf_9d9681e93d5a567f++ {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*__obf_1f271b19471a7add)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_ab9eaf73a4199f9c(__obf_8da81bf3c5df7f0d, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_3a1983a2984d289a(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_e37ee56dd8133d3f; __obf_9d9681e93d5a567f++ {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*__obf_1f271b19471a7add)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_ab9eaf73a4199f9c(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b.YCbCr()
	case *image.RGBA64:
		__obf_61b54db7046becb4 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA64(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_4a688a78d5582d42(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_eb72b995d24df161(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_4a688a78d5582d42(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_eb72b995d24df161(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// 16-bit precision
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.NRGBA64:
		__obf_61b54db7046becb4 := image.NewRGBA64(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA64(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_4a688a78d5582d42(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_c1c846751d43d0f4(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_4a688a78d5582d42(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_eb72b995d24df161(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// 8-bit precision
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.Gray:
		__obf_61b54db7046becb4 := image.NewGray(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewGray(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_3a1983a2984d289a(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_e37ee56dd8133d3f; __obf_9d9681e93d5a567f++ {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_44079ea63696618c(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_3a1983a2984d289a(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_44079ea63696618c(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// 16-bit precision
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.Gray16:
		__obf_61b54db7046becb4 := image.NewGray16(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewGray16(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_4a688a78d5582d42(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray16)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_010f5d83554658c7(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_4a688a78d5582d42(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray16)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_010f5d83554658c7(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// 16-bit precision
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	default:
		__obf_61b54db7046becb4 := image.NewRGBA64(image.Rect(0, 0, __obf_93755cc657d3c1d7.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA64(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_4a688a78d5582d42(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_d42ce57ba503f1f8(__obf_93755cc657d3c1d7, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_4a688a78d5582d42(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a, __obf_e0303ccb3cff6d43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_eb72b995d24df161(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	}
}

func __obf_539928a5c1774b7b(__obf_ba0d63dbead1ced9, __obf_c4381a5b61d233ea uint, __obf_bb203b2e26665a43, __obf_f464c7a8df992c8a float64, __obf_93755cc657d3c1d7 image.Image, __obf_ebcc3d09dc09d13d InterpolationFunction) image.Image {
	__obf_aacb5abca7c80175, _ := __obf_ebcc3d09dc09d13d.__obf_e0303ccb3cff6d43()
	__obf_e37ee56dd8133d3f := runtime.GOMAXPROCS(0)
	__obf_c9026fcbe3b0f1fd := sync.WaitGroup{}

	switch __obf_b8394cc3da38f2d2 := __obf_93755cc657d3c1d7.(type) {
	case *image.RGBA:
		__obf_61b54db7046becb4 := // 8-bit precision
			image.NewRGBA(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := 0; __obf_9d9681e93d5a567f < __obf_e37ee56dd8133d3f; __obf_9d9681e93d5a567f++ {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_1d1b25a5bcd2a73c(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// horizontal filter on transposed image, result is not transposed
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_1d1b25a5bcd2a73c(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.NRGBA:
		__obf_61b54db7046becb4 := // 8-bit precision
			image.NewNRGBA(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewNRGBA(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.NRGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_bedef5bbf1790511(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// horizontal filter on transposed image, result is not transposed
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.NRGBA)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_bedef5bbf1790511(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.YCbCr:
		__obf_61b54db7046becb4 := // 8-bit precision
			// accessing the YCbCr arrays in a tight loop is slow.
			// converting the image to ycc increases performance by 2x.
			__obf_8decfa8633084a6c(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)), __obf_b8394cc3da38f2d2.SubsampleRatio)
		__obf_cb722e41a1b9d39b := __obf_8decfa8633084a6c(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)), image.YCbCrSubsampleRatio444)
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_8da81bf3c5df7f0d := __obf_6235574c4a9792da(__obf_b8394cc3da38f2d2)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*__obf_1f271b19471a7add)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_de165637ccf50364(__obf_8da81bf3c5df7f0d, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*__obf_1f271b19471a7add)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_de165637ccf50364(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_f464c7a8df992c8a, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b.YCbCr()
	case *image.RGBA64:
		__obf_61b54db7046becb4 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA64(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_7a7f84fcfd267669(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// horizontal filter on transposed image, result is not transposed
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_7a7f84fcfd267669(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.NRGBA64:
		__obf_61b54db7046becb4 := // 16-bit precision
			image.NewNRGBA64(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewNRGBA64(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.NRGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_5864fea3a63745f1(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// horizontal filter on transposed image, result is not transposed
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.NRGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_5864fea3a63745f1(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.Gray:
		__obf_61b54db7046becb4 := // 8-bit precision
			image.NewGray(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewGray(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_c0cb06ed5d8c6356(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// horizontal filter on transposed image, result is not transposed
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_c0cb06ed5d8c6356(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	case *image.Gray16:
		__obf_61b54db7046becb4 := // 16-bit precision
			image.NewGray16(image.Rect(0, 0, __obf_b8394cc3da38f2d2.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewGray16(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray16)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_9803aa82d6ad8fd8(__obf_b8394cc3da38f2d2, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// horizontal filter on transposed image, result is not transposed
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.Gray16)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_9803aa82d6ad8fd8(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	default:
		__obf_61b54db7046becb4 := // 16-bit precision
			image.NewRGBA64(image.Rect(0, 0, __obf_93755cc657d3c1d7.Bounds().Dy(), int(__obf_ba0d63dbead1ced9)))
		__obf_cb722e41a1b9d39b := image.NewRGBA64(image.Rect(0, 0, int(__obf_ba0d63dbead1ced9), int(__obf_c4381a5b61d233ea)))
		__obf_3fabfc59cdeeb8b4,

			// horizontal filter, results in transposed temporary image
			__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd := __obf_c218d3902b533b64(__obf_61b54db7046becb4.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_bb203b2e26665a43)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_61b54db7046becb4, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_9224993e7b70f1c7(__obf_93755cc657d3c1d7, __obf_731510c7cb336b26, __obf_bb203b2e26665a43, __obf_3fabfc59cdeeb8b4,

					// horizontal filter on transposed image, result is not transposed
					__obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		__obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8, __obf_52dfb4be4832ebbd = __obf_c218d3902b533b64(__obf_cb722e41a1b9d39b.Bounds().Dy(), __obf_aacb5abca7c80175, __obf_74fbaffb42caf3d6, __obf_f464c7a8df992c8a)
		__obf_c9026fcbe3b0f1fd.
			Add(__obf_e37ee56dd8133d3f)
		for __obf_9d9681e93d5a567f := range __obf_e37ee56dd8133d3f {
			__obf_731510c7cb336b26 := __obf_961eb980dfa40b9d(__obf_cb722e41a1b9d39b, __obf_9d9681e93d5a567f, __obf_e37ee56dd8133d3f).(*image.RGBA64)
			go func() {
				defer __obf_c9026fcbe3b0f1fd.Done()
				__obf_7a7f84fcfd267669(__obf_61b54db7046becb4, __obf_731510c7cb336b26, __obf_3fabfc59cdeeb8b4, __obf_c032dbb648d056a8,

					// Calculates scaling factors using old and new image dimensions.
					__obf_52dfb4be4832ebbd)
			}()
		}
		__obf_c9026fcbe3b0f1fd.
			Wait()
		return __obf_cb722e41a1b9d39b
	}

}

func __obf_1da7c3c6d2fcdf43(__obf_ba0d63dbead1ced9, __obf_c4381a5b61d233ea uint, __obf_a6306e37a4ae4714, __obf_aa60b9b8d8f67286 float64) (__obf_bb203b2e26665a43, __obf_f464c7a8df992c8a float64) {
	if __obf_ba0d63dbead1ced9 == 0 {
		if __obf_c4381a5b61d233ea == 0 {
			__obf_bb203b2e26665a43 = 1.0
			__obf_f464c7a8df992c8a = 1.0
		} else {
			__obf_f464c7a8df992c8a = __obf_aa60b9b8d8f67286 / float64(__obf_c4381a5b61d233ea)
			__obf_bb203b2e26665a43 = __obf_f464c7a8df992c8a
		}
	} else {
		__obf_bb203b2e26665a43 = __obf_a6306e37a4ae4714 / float64(__obf_ba0d63dbead1ced9)
		if __obf_c4381a5b61d233ea == 0 {
			__obf_f464c7a8df992c8a = __obf_bb203b2e26665a43
		} else {
			__obf_f464c7a8df992c8a = __obf_aa60b9b8d8f67286 / float64(__obf_c4381a5b61d233ea)
		}
	}
	return
}

type __obf_3a6bcf0735c1821e interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func __obf_961eb980dfa40b9d(__obf_93755cc657d3c1d7 __obf_3a6bcf0735c1821e, __obf_9d9681e93d5a567f, __obf_e65729ac49f2f338 int) image.Image {
	return __obf_93755cc657d3c1d7.SubImage(image.Rect(__obf_93755cc657d3c1d7.Bounds().Min.X, __obf_93755cc657d3c1d7.Bounds().Min.Y+__obf_9d9681e93d5a567f*__obf_93755cc657d3c1d7.Bounds().Dy()/__obf_e65729ac49f2f338, __obf_93755cc657d3c1d7.Bounds().Max.X, __obf_93755cc657d3c1d7.Bounds().Min.Y+(__obf_9d9681e93d5a567f+1)*__obf_93755cc657d3c1d7.Bounds().Dy()/__obf_e65729ac49f2f338))
}
