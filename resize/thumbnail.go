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

package __obf_d24101647651f4c3

import (
	"image"
)

// Thumbnail will downscale provided image to max width and height preserving
// original aspect ratio and using the interpolation function interp.
// It will return original image, without processing it, if original sizes
// are already smaller than provided constraints.
func Thumbnail(__obf_4737a42c277b8cdf, __obf_1db04f5ed22585c8 uint, __obf_b9c5bf10bfae0f7a image.Image, __obf_0c2c83d39fca0ad5 InterpolationFunction) image.Image {
	__obf_9743b2780c056351 := __obf_b9c5bf10bfae0f7a.Bounds()
	__obf_13ad8a2e2dd50dbd := uint(__obf_9743b2780c056351.Dx())
	__obf_f354dcf7e4cbd7d8 := uint(__obf_9743b2780c056351.Dy())
	__obf_5b88654d8f85a1cb, __obf_b88faf03d4d948e9 := __obf_13ad8a2e2dd50dbd, __obf_f354dcf7e4cbd7d8

	// Return original image if it have same or smaller size as constraints
	if __obf_4737a42c277b8cdf >= __obf_13ad8a2e2dd50dbd && __obf_1db04f5ed22585c8 >= __obf_f354dcf7e4cbd7d8 {
		return __obf_b9c5bf10bfae0f7a
	}

	// Preserve aspect ratio
	if __obf_13ad8a2e2dd50dbd > __obf_4737a42c277b8cdf {
		__obf_b88faf03d4d948e9 = max(uint(__obf_f354dcf7e4cbd7d8*__obf_4737a42c277b8cdf/__obf_13ad8a2e2dd50dbd), 1)
		__obf_5b88654d8f85a1cb = __obf_4737a42c277b8cdf
	}

	if __obf_b88faf03d4d948e9 > __obf_1db04f5ed22585c8 {
		__obf_5b88654d8f85a1cb = max(uint(__obf_5b88654d8f85a1cb*__obf_1db04f5ed22585c8/__obf_b88faf03d4d948e9), 1)
		__obf_b88faf03d4d948e9 = __obf_1db04f5ed22585c8
	}
	return Resize(__obf_5b88654d8f85a1cb, __obf_b88faf03d4d948e9, __obf_b9c5bf10bfae0f7a, __obf_0c2c83d39fca0ad5)
}
