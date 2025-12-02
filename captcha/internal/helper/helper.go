package __obf_fa149b22d3814e30

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

// DrawString draws a string on the canvas with high-quality rendering
func DrawString(__obf_9f9a936eb2e94626 draw.Image, __obf_79f4eb76e5bd22c2 *types.DrawStringParam, __obf_53a220da81685ca9 fixed.Point26_6) error {
	__obf_44dfa64663c32c4e, __obf_cf1c648003a311b3 := opentype.NewFace(__obf_79f4eb76e5bd22c2.Font, &opentype.FaceOptions{
		Size:    float64(__obf_79f4eb76e5bd22c2.FontSize),
		DPI:     float64(__obf_79f4eb76e5bd22c2.FontDPI),
		Hinting: __obf_79f4eb76e5bd22c2.FontHinting,
	})
	if __obf_cf1c648003a311b3 != nil {
		return __obf_cf1c648003a311b3
	}
	defer __obf_44dfa64663c32c4e.Close()
	__obf_6c0ef3e0b2d144be := &font.Drawer{
		Dst:  __obf_9f9a936eb2e94626,
		Src:  image.NewUniform(__obf_79f4eb76e5bd22c2.Color),
		Face: __obf_44dfa64663c32c4e,
		Dot:  __obf_53a220da81685ca9,
	}
	__obf_6c0ef3e0b2d144be.
		DrawString(__obf_79f4eb76e5bd22c2.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_16f63ad182dbb409, __obf_0154ab49120701ff int, __obf_b8b29feaf1c7d62e []color.RGBA) *types.Palette {
	__obf_9f9a936eb2e94626 := make([]color.Color, 0, len(__obf_b8b29feaf1c7d62e)+1)
	__obf_9f9a936eb2e94626 = append(__obf_9f9a936eb2e94626, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_8f68814d18077660 := range __obf_b8b29feaf1c7d62e {
		__obf_9f9a936eb2e94626 = append(__obf_9f9a936eb2e94626, __obf_8f68814d18077660)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_16f63ad182dbb409, __obf_0154ab49120701ff), __obf_9f9a936eb2e94626)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_16f63ad182dbb409, __obf_0154ab49120701ff int, __obf_ea6f382a3d861a0a bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_16f63ad182dbb409, __obf_0154ab49120701ff), __obf_ea6f382a3d861a0a)
}
