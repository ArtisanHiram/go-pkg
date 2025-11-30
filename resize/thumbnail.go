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

package __obf_ac510735d4d13cdd

import (
	"image"
)

// Thumbnail will downscale provided image to max width and height preserving
// original aspect ratio and using the interpolation function interp.
// It will return original image, without processing it, if original sizes
// are already smaller than provided constraints.
func Thumbnail(__obf_752da2cb27b84607, __obf_8436c2544e9c6842 uint, __obf_4f5a8d0b34527afc image.Image, __obf_00264b26e01a01f0 InterpolationFunction) image.Image {
	__obf_4c9ed49e96316785 := __obf_4f5a8d0b34527afc.Bounds()
	__obf_bf6f3ec7e08e3983 := uint(__obf_4c9ed49e96316785.Dx())
	__obf_897dbc09c7559dae := uint(__obf_4c9ed49e96316785.Dy())
	__obf_4443d16d9e84a6da, __obf_8f872695c2eb2849 := __obf_bf6f3ec7e08e3983, __obf_897dbc09c7559dae

	// Return original image if it have same or smaller size as constraints
	if __obf_752da2cb27b84607 >= __obf_bf6f3ec7e08e3983 && __obf_8436c2544e9c6842 >= __obf_897dbc09c7559dae {
		return __obf_4f5a8d0b34527afc
	}

	// Preserve aspect ratio
	if __obf_bf6f3ec7e08e3983 > __obf_752da2cb27b84607 {
		__obf_8f872695c2eb2849 = max(uint(__obf_897dbc09c7559dae*__obf_752da2cb27b84607/__obf_bf6f3ec7e08e3983), 1)
		__obf_4443d16d9e84a6da = __obf_752da2cb27b84607
	}

	if __obf_8f872695c2eb2849 > __obf_8436c2544e9c6842 {
		__obf_4443d16d9e84a6da = max(uint(__obf_4443d16d9e84a6da*__obf_8436c2544e9c6842/__obf_8f872695c2eb2849), 1)
		__obf_8f872695c2eb2849 = __obf_8436c2544e9c6842
	}
	return Resize(__obf_4443d16d9e84a6da, __obf_8f872695c2eb2849, __obf_4f5a8d0b34527afc, __obf_00264b26e01a01f0)
}
