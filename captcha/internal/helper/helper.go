package __obf_339ea479ec4d549f

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
func DrawString(__obf_27e3280d911273cd draw.Image, __obf_6338fcba1f93e3ae *types.DrawStringParam, __obf_aa65fec8564bacce fixed.Point26_6) error {
	__obf_00e082898a79bcfc, __obf_cf21fc297e0f6e30 := opentype.NewFace(__obf_6338fcba1f93e3ae.Font, &opentype.FaceOptions{
		Size:    float64(__obf_6338fcba1f93e3ae.FontSize),
		DPI:     float64(__obf_6338fcba1f93e3ae.FontDPI),
		Hinting: __obf_6338fcba1f93e3ae.FontHinting,
	})
	if __obf_cf21fc297e0f6e30 != nil {
		return __obf_cf21fc297e0f6e30
	}
	defer __obf_00e082898a79bcfc.Close()

	__obf_281f48fc04c72ffe := &font.Drawer{
		Dst:  __obf_27e3280d911273cd,
		Src:  image.NewUniform(__obf_6338fcba1f93e3ae.Color),
		Face: __obf_00e082898a79bcfc,
		Dot:  __obf_aa65fec8564bacce,
	}
	__obf_281f48fc04c72ffe.DrawString(__obf_6338fcba1f93e3ae.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_c35c90c2e6ac7db9, __obf_d2b863d65591d09c int, __obf_3dd63487ec0219dd []color.RGBA) *types.Palette {
	__obf_27e3280d911273cd := make([]color.Color, 0, len(__obf_3dd63487ec0219dd)+1)
	__obf_27e3280d911273cd = append(__obf_27e3280d911273cd, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_a1ded02a6ce849ac := range __obf_3dd63487ec0219dd {
		__obf_27e3280d911273cd = append(__obf_27e3280d911273cd, __obf_a1ded02a6ce849ac)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_c35c90c2e6ac7db9, __obf_d2b863d65591d09c), __obf_27e3280d911273cd)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_c35c90c2e6ac7db9, __obf_d2b863d65591d09c int, __obf_6294b9693d93429f bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_c35c90c2e6ac7db9, __obf_d2b863d65591d09c), __obf_6294b9693d93429f)
}
