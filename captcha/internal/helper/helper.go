package __obf_46685d1b8357802f

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
func DrawString(__obf_5ecf496288e5f149 draw.Image, __obf_43b8e36b51bef8d5 *types.DrawStringParam, __obf_a8e1f8ffe0b12fa6 fixed.Point26_6) error {
	__obf_5953fb47dd765925, __obf_1e7c77fdd506b7f0 := opentype.NewFace(__obf_43b8e36b51bef8d5.Font, &opentype.FaceOptions{
		Size:    float64(__obf_43b8e36b51bef8d5.FontSize),
		DPI:     float64(__obf_43b8e36b51bef8d5.FontDPI),
		Hinting: __obf_43b8e36b51bef8d5.FontHinting,
	})
	if __obf_1e7c77fdd506b7f0 != nil {
		return __obf_1e7c77fdd506b7f0
	}
	defer __obf_5953fb47dd765925.Close()

	__obf_b7142c40f3499f83 := &font.Drawer{
		Dst:  __obf_5ecf496288e5f149,
		Src:  image.NewUniform(__obf_43b8e36b51bef8d5.Color),
		Face: __obf_5953fb47dd765925,
		Dot:  __obf_a8e1f8ffe0b12fa6,
	}
	__obf_b7142c40f3499f83.DrawString(__obf_43b8e36b51bef8d5.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_556b479bf2421137, __obf_d0c48e43cf612abf int, __obf_30ea65b03e929b79 []color.RGBA) *types.Palette {
	__obf_5ecf496288e5f149 := make([]color.Color, 0, len(__obf_30ea65b03e929b79)+1)
	__obf_5ecf496288e5f149 = append(__obf_5ecf496288e5f149, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_4fb580350581f30e := range __obf_30ea65b03e929b79 {
		__obf_5ecf496288e5f149 = append(__obf_5ecf496288e5f149, __obf_4fb580350581f30e)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_556b479bf2421137, __obf_d0c48e43cf612abf), __obf_5ecf496288e5f149)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_556b479bf2421137, __obf_d0c48e43cf612abf int, __obf_94e494f0dfc796a5 bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_556b479bf2421137, __obf_d0c48e43cf612abf), __obf_94e494f0dfc796a5)
}
