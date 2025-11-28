package __obf_f3613af2738989b9

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
)

// DrawString draws a string on the canvas with high-quality rendering
func DrawString(__obf_bae7d5b3fc18b943 draw.Image, __obf_0018bac08a4bb4b4 *types.DrawStringParam, __obf_49ec54258ff10ef6 fixed.Point26_6) error {
	__obf_f41933b64382c33f, __obf_183f5b20d9bc7819 := opentype.NewFace(__obf_0018bac08a4bb4b4.Font, &opentype.FaceOptions{
		Size:    float64(__obf_0018bac08a4bb4b4.FontSize),
		DPI:     float64(__obf_0018bac08a4bb4b4.FontDPI),
		Hinting: __obf_0018bac08a4bb4b4.FontHinting,
	})
	if __obf_183f5b20d9bc7819 != nil {
		return __obf_183f5b20d9bc7819
	}
	defer __obf_f41933b64382c33f.Close()

	__obf_ba8c1564ec9c1061 := &font.Drawer{
		Dst:  __obf_bae7d5b3fc18b943,
		Src:  image.NewUniform(__obf_0018bac08a4bb4b4.Color),
		Face: __obf_f41933b64382c33f,
		Dot:  __obf_49ec54258ff10ef6,
	}
	__obf_ba8c1564ec9c1061.DrawString(__obf_0018bac08a4bb4b4.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_7a8b82f3b1e44492, __obf_d26823b19a4221b0 int, __obf_c7c78332ea000074 []color.RGBA) *types.Palette {
	__obf_bae7d5b3fc18b943 := make([]color.Color, 0, len(__obf_c7c78332ea000074)+1)
	__obf_bae7d5b3fc18b943 = append(__obf_bae7d5b3fc18b943, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_564aad593551cc0d := range __obf_c7c78332ea000074 {
		__obf_bae7d5b3fc18b943 = append(__obf_bae7d5b3fc18b943, __obf_564aad593551cc0d)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_7a8b82f3b1e44492, __obf_d26823b19a4221b0), __obf_bae7d5b3fc18b943)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_7a8b82f3b1e44492, __obf_d26823b19a4221b0 int, __obf_f0bee8fb40bbdec3 bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_7a8b82f3b1e44492, __obf_d26823b19a4221b0), __obf_f0bee8fb40bbdec3)
}
