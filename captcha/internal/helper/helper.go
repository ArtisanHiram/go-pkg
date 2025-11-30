package __obf_b14d24b29d34b846

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
func DrawString(__obf_bbbfd6e29747a9f0 draw.Image, __obf_331a710a85c470da *types.DrawStringParam, __obf_3e1cf2f3315a835a fixed.Point26_6) error {
	__obf_ee3b8bfbba670767, __obf_801b060480f0cee4 := opentype.NewFace(__obf_331a710a85c470da.Font, &opentype.FaceOptions{
		Size:    float64(__obf_331a710a85c470da.FontSize),
		DPI:     float64(__obf_331a710a85c470da.FontDPI),
		Hinting: __obf_331a710a85c470da.FontHinting,
	})
	if __obf_801b060480f0cee4 != nil {
		return __obf_801b060480f0cee4
	}
	defer __obf_ee3b8bfbba670767.Close()
	__obf_c8bda2551f18ac9b := &font.Drawer{
		Dst:  __obf_bbbfd6e29747a9f0,
		Src:  image.NewUniform(__obf_331a710a85c470da.Color),
		Face: __obf_ee3b8bfbba670767,
		Dot:  __obf_3e1cf2f3315a835a,
	}
	__obf_c8bda2551f18ac9b.
		DrawString(__obf_331a710a85c470da.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_e56c5812359ee0cd, __obf_eb4a308a5061712b int, __obf_a0b45e9ce1fe28ea []color.RGBA) *types.Palette {
	__obf_bbbfd6e29747a9f0 := make([]color.Color, 0, len(__obf_a0b45e9ce1fe28ea)+1)
	__obf_bbbfd6e29747a9f0 = append(__obf_bbbfd6e29747a9f0, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_28d6381c666e0c17 := range __obf_a0b45e9ce1fe28ea {
		__obf_bbbfd6e29747a9f0 = append(__obf_bbbfd6e29747a9f0, __obf_28d6381c666e0c17)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_e56c5812359ee0cd, __obf_eb4a308a5061712b), __obf_bbbfd6e29747a9f0)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_e56c5812359ee0cd, __obf_eb4a308a5061712b int, __obf_b6a34424a958b17e bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_e56c5812359ee0cd, __obf_eb4a308a5061712b), __obf_b6a34424a958b17e)
}
