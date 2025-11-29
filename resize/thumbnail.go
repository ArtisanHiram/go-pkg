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

package __obf_3858dbfa2ccfdbe9

import (
	"image"
)

// Thumbnail will downscale provided image to max width and height preserving
// original aspect ratio and using the interpolation function interp.
// It will return original image, without processing it, if original sizes
// are already smaller than provided constraints.
func Thumbnail(__obf_0d6c2d73f3015ccf, __obf_ebc0b9db03408d03 uint, __obf_93755cc657d3c1d7 image.Image, __obf_ebcc3d09dc09d13d InterpolationFunction) image.Image {
	__obf_916dab67b5d569d6 := __obf_93755cc657d3c1d7.Bounds()
	__obf_3071cbeccdbebc6e := uint(__obf_916dab67b5d569d6.Dx())
	__obf_05307a08604bc33f := uint(__obf_916dab67b5d569d6.Dy())
	__obf_9f24a17228c693d2, __obf_a50553f2f446f0b2 := __obf_3071cbeccdbebc6e, __obf_05307a08604bc33f

	// Return original image if it have same or smaller size as constraints
	if __obf_0d6c2d73f3015ccf >= __obf_3071cbeccdbebc6e && __obf_ebc0b9db03408d03 >= __obf_05307a08604bc33f {
		return __obf_93755cc657d3c1d7
	}

	// Preserve aspect ratio
	if __obf_3071cbeccdbebc6e > __obf_0d6c2d73f3015ccf {
		__obf_a50553f2f446f0b2 = max(uint(__obf_05307a08604bc33f*__obf_0d6c2d73f3015ccf/__obf_3071cbeccdbebc6e), 1)
		__obf_9f24a17228c693d2 = __obf_0d6c2d73f3015ccf
	}

	if __obf_a50553f2f446f0b2 > __obf_ebc0b9db03408d03 {
		__obf_9f24a17228c693d2 = max(uint(__obf_9f24a17228c693d2*__obf_ebc0b9db03408d03/__obf_a50553f2f446f0b2), 1)
		__obf_a50553f2f446f0b2 = __obf_ebc0b9db03408d03
	}
	return Resize(__obf_9f24a17228c693d2, __obf_a50553f2f446f0b2, __obf_93755cc657d3c1d7, __obf_ebcc3d09dc09d13d)
}
