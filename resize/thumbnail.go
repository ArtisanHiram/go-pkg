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

package __obf_4f71249d3f545aa0

import (
	"image"
)

// Thumbnail will downscale provided image to max width and height preserving
// original aspect ratio and using the interpolation function interp.
// It will return original image, without processing it, if original sizes
// are already smaller than provided constraints.
func Thumbnail(__obf_934294700eb8e8a3, __obf_fa7c1ad284ee6499 uint, __obf_cb2c55730faeb3f0 image.Image, __obf_f617084106645ca0 InterpolationFunction) image.Image {
	__obf_af1af6eb6b5f32fa := __obf_cb2c55730faeb3f0.Bounds()
	__obf_b3cb0e4c438cc52f := uint(__obf_af1af6eb6b5f32fa.Dx())
	__obf_921c79a400c1db4e := uint(__obf_af1af6eb6b5f32fa.Dy())
	__obf_16857a4f655b360b, __obf_e95837b55e318ab2 := __obf_b3cb0e4c438cc52f, __obf_921c79a400c1db4e

	// Return original image if it have same or smaller size as constraints
	if __obf_934294700eb8e8a3 >= __obf_b3cb0e4c438cc52f && __obf_fa7c1ad284ee6499 >= __obf_921c79a400c1db4e {
		return __obf_cb2c55730faeb3f0
	}

	// Preserve aspect ratio
	if __obf_b3cb0e4c438cc52f > __obf_934294700eb8e8a3 {
		__obf_e95837b55e318ab2 = max(uint(__obf_921c79a400c1db4e*__obf_934294700eb8e8a3/__obf_b3cb0e4c438cc52f), 1)
		__obf_16857a4f655b360b = __obf_934294700eb8e8a3
	}

	if __obf_e95837b55e318ab2 > __obf_fa7c1ad284ee6499 {
		__obf_16857a4f655b360b = max(uint(__obf_16857a4f655b360b*__obf_fa7c1ad284ee6499/__obf_e95837b55e318ab2), 1)
		__obf_e95837b55e318ab2 = __obf_fa7c1ad284ee6499
	}
	return Resize(__obf_16857a4f655b360b, __obf_e95837b55e318ab2, __obf_cb2c55730faeb3f0, __obf_f617084106645ca0)
}
