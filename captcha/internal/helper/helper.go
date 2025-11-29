package __obf_de56821ae4303350

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
func DrawString(__obf_725679b759d2962d draw.Image, __obf_4a9c3195b2e46874 *types.DrawStringParam, __obf_0d67640c8a1d3953 fixed.Point26_6) error {
	__obf_fa19ccf549aefd83, __obf_7e0d30cbb9b6a603 := opentype.NewFace(__obf_4a9c3195b2e46874.Font, &opentype.FaceOptions{
		Size:    float64(__obf_4a9c3195b2e46874.FontSize),
		DPI:     float64(__obf_4a9c3195b2e46874.FontDPI),
		Hinting: __obf_4a9c3195b2e46874.FontHinting,
	})
	if __obf_7e0d30cbb9b6a603 != nil {
		return __obf_7e0d30cbb9b6a603
	}
	defer __obf_fa19ccf549aefd83.Close()
	__obf_b9a93fd801151b5c := &font.Drawer{
		Dst:  __obf_725679b759d2962d,
		Src:  image.NewUniform(__obf_4a9c3195b2e46874.Color),
		Face: __obf_fa19ccf549aefd83,
		Dot:  __obf_0d67640c8a1d3953,
	}
	__obf_b9a93fd801151b5c.
		DrawString(__obf_4a9c3195b2e46874.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_3fd53e6f70909f51, __obf_e865dea6cb4c1509 int, __obf_92aee474a2d8a1d0 []color.RGBA) *types.Palette {
	__obf_725679b759d2962d := make([]color.Color, 0, len(__obf_92aee474a2d8a1d0)+1)
	__obf_725679b759d2962d = append(__obf_725679b759d2962d, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_7ae6b70dd387bdc2 := range __obf_92aee474a2d8a1d0 {
		__obf_725679b759d2962d = append(__obf_725679b759d2962d, __obf_7ae6b70dd387bdc2)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_3fd53e6f70909f51, __obf_e865dea6cb4c1509), __obf_725679b759d2962d)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_3fd53e6f70909f51, __obf_e865dea6cb4c1509 int, __obf_bb546f5754f4fd92 bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_3fd53e6f70909f51, __obf_e865dea6cb4c1509), __obf_bb546f5754f4fd92)
}
