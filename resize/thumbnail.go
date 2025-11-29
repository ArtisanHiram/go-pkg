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

package __obf_71f4605f6400e2ce

import (
	"image"
)

// Thumbnail will downscale provided image to max width and height preserving
// original aspect ratio and using the interpolation function interp.
// It will return original image, without processing it, if original sizes
// are already smaller than provided constraints.
func Thumbnail(__obf_3cb09cfb0114d2fe, __obf_f54184a1c0412350 uint, __obf_0b888bcf5266913b image.Image, __obf_a66df7eedc4b118a InterpolationFunction) image.Image {
	__obf_5821b8f6a7a1d1b5 := __obf_0b888bcf5266913b.Bounds()
	__obf_1c6811b1af5b6261 := uint(__obf_5821b8f6a7a1d1b5.Dx())
	__obf_cd99f72f78c4c0f9 := uint(__obf_5821b8f6a7a1d1b5.Dy())
	__obf_8f40c57b6a11ca9c, __obf_8efbda41870e5544 := __obf_1c6811b1af5b6261, __obf_cd99f72f78c4c0f9

	// Return original image if it have same or smaller size as constraints
	if __obf_3cb09cfb0114d2fe >= __obf_1c6811b1af5b6261 && __obf_f54184a1c0412350 >= __obf_cd99f72f78c4c0f9 {
		return __obf_0b888bcf5266913b
	}

	// Preserve aspect ratio
	if __obf_1c6811b1af5b6261 > __obf_3cb09cfb0114d2fe {
		__obf_8efbda41870e5544 = max(uint(__obf_cd99f72f78c4c0f9*__obf_3cb09cfb0114d2fe/__obf_1c6811b1af5b6261), 1)
		__obf_8f40c57b6a11ca9c = __obf_3cb09cfb0114d2fe
	}

	if __obf_8efbda41870e5544 > __obf_f54184a1c0412350 {
		__obf_8f40c57b6a11ca9c = max(uint(__obf_8f40c57b6a11ca9c*__obf_f54184a1c0412350/__obf_8efbda41870e5544), 1)
		__obf_8efbda41870e5544 = __obf_f54184a1c0412350
	}
	return Resize(__obf_8f40c57b6a11ca9c, __obf_8efbda41870e5544, __obf_0b888bcf5266913b, __obf_a66df7eedc4b118a)
}
