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

package __obf_42b2ccbdafaee9c3

import (
	"image"
)

// Thumbnail will downscale provided image to max width and height preserving
// original aspect ratio and using the interpolation function interp.
// It will return original image, without processing it, if original sizes
// are already smaller than provided constraints.
func Thumbnail(__obf_56d06c6cd95c30a4, __obf_340e826815fe3163 uint, __obf_99d54e4a80501ee0 image.Image, __obf_f1c487ecb00e0380 InterpolationFunction) image.Image {
	__obf_5db009bf616c12da := __obf_99d54e4a80501ee0.Bounds()
	__obf_2afa3344815e639f := uint(__obf_5db009bf616c12da.Dx())
	__obf_fe92c904738f23c0 := uint(__obf_5db009bf616c12da.Dy())
	__obf_d3d37d3ab1a814bd, __obf_5270a1e11aa3b663 := __obf_2afa3344815e639f, __obf_fe92c904738f23c0

	// Return original image if it have same or smaller size as constraints
	if __obf_56d06c6cd95c30a4 >= __obf_2afa3344815e639f && __obf_340e826815fe3163 >= __obf_fe92c904738f23c0 {
		return __obf_99d54e4a80501ee0
	}

	// Preserve aspect ratio
	if __obf_2afa3344815e639f > __obf_56d06c6cd95c30a4 {
		__obf_5270a1e11aa3b663 = max(uint(__obf_fe92c904738f23c0*__obf_56d06c6cd95c30a4/__obf_2afa3344815e639f), 1)
		__obf_d3d37d3ab1a814bd = __obf_56d06c6cd95c30a4
	}

	if __obf_5270a1e11aa3b663 > __obf_340e826815fe3163 {
		__obf_d3d37d3ab1a814bd = max(uint(__obf_d3d37d3ab1a814bd*__obf_340e826815fe3163/__obf_5270a1e11aa3b663), 1)
		__obf_5270a1e11aa3b663 = __obf_340e826815fe3163
	}
	return Resize(__obf_d3d37d3ab1a814bd, __obf_5270a1e11aa3b663, __obf_99d54e4a80501ee0, __obf_f1c487ecb00e0380)
}
