package __obf_6d6e76c2b0d22280

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
func DrawString(__obf_af7513e722bb880c draw.Image, __obf_87ca9edd0a2da98a *types.DrawStringParam, __obf_67176c3b4c86ceb5 fixed.Point26_6) error {
	__obf_39b2522ea2b1a592, __obf_3f459bfa7b3a9a32 := opentype.NewFace(__obf_87ca9edd0a2da98a.Font, &opentype.FaceOptions{
		Size:    float64(__obf_87ca9edd0a2da98a.FontSize),
		DPI:     float64(__obf_87ca9edd0a2da98a.FontDPI),
		Hinting: __obf_87ca9edd0a2da98a.FontHinting,
	})
	if __obf_3f459bfa7b3a9a32 != nil {
		return __obf_3f459bfa7b3a9a32
	}
	defer __obf_39b2522ea2b1a592.Close()
	__obf_8d31483bba874543 := &font.Drawer{
		Dst:  __obf_af7513e722bb880c,
		Src:  image.NewUniform(__obf_87ca9edd0a2da98a.Color),
		Face: __obf_39b2522ea2b1a592,
		Dot:  __obf_67176c3b4c86ceb5,
	}
	__obf_8d31483bba874543.
		DrawString(__obf_87ca9edd0a2da98a.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_76b86dd6b407f86a, __obf_a31176ef0b86bf9a int, __obf_c7e80a2e38fdf658 []color.RGBA) *types.Palette {
	__obf_af7513e722bb880c := make([]color.Color, 0, len(__obf_c7e80a2e38fdf658)+1)
	__obf_af7513e722bb880c = append(__obf_af7513e722bb880c, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_d243d925002bc204 := range __obf_c7e80a2e38fdf658 {
		__obf_af7513e722bb880c = append(__obf_af7513e722bb880c, __obf_d243d925002bc204)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_76b86dd6b407f86a, __obf_a31176ef0b86bf9a), __obf_af7513e722bb880c)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_76b86dd6b407f86a, __obf_a31176ef0b86bf9a int, __obf_05899d056ee79bfa bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_76b86dd6b407f86a, __obf_a31176ef0b86bf9a), __obf_05899d056ee79bfa)
}
