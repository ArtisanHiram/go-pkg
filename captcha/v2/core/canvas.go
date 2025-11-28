package __obf_0e73be9e4c3ea3fd

import (
	"image"
	"image/color"
)

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_0107a6cb34e1b800, __obf_1f3235dfd61ebd0b int, __obf_f47867622d10be63 []color.RGBA) *Palette {
	__obf_c0d27cca35ba9554 := make([]color.Color, 0, len(__obf_f47867622d10be63)+1)
	__obf_c0d27cca35ba9554 = append(__obf_c0d27cca35ba9554, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_f8353e66f8032833 := range __obf_f47867622d10be63 {
		__obf_c0d27cca35ba9554 = append(__obf_c0d27cca35ba9554, __obf_f8353e66f8032833)
	}

	return NewPalette(image.Rect(0, 0, __obf_0107a6cb34e1b800, __obf_1f3235dfd61ebd0b), __obf_c0d27cca35ba9554)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_0107a6cb34e1b800, __obf_1f3235dfd61ebd0b int, __obf_64e9fc49b08a5410 bool) *NRGBA {
	return NewNRGBA(image.Rect(0, 0, __obf_0107a6cb34e1b800, __obf_1f3235dfd61ebd0b), __obf_64e9fc49b08a5410)
}
