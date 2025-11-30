package __obf_acc972acee86b366

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
func DrawString(__obf_8e4fbb467ecd3e76 draw.Image, __obf_e4ce4cb0f09249fc *types.DrawStringParam, __obf_276862b5c0dc8b20 fixed.Point26_6) error {
	__obf_c322751dc3dfe379, __obf_5ac9200fd05f8b71 := opentype.NewFace(__obf_e4ce4cb0f09249fc.Font, &opentype.FaceOptions{
		Size:    float64(__obf_e4ce4cb0f09249fc.FontSize),
		DPI:     float64(__obf_e4ce4cb0f09249fc.FontDPI),
		Hinting: __obf_e4ce4cb0f09249fc.FontHinting,
	})
	if __obf_5ac9200fd05f8b71 != nil {
		return __obf_5ac9200fd05f8b71
	}
	defer __obf_c322751dc3dfe379.Close()
	__obf_47e171db5303ac65 := &font.Drawer{
		Dst:  __obf_8e4fbb467ecd3e76,
		Src:  image.NewUniform(__obf_e4ce4cb0f09249fc.Color),
		Face: __obf_c322751dc3dfe379,
		Dot:  __obf_276862b5c0dc8b20,
	}
	__obf_47e171db5303ac65.
		DrawString(__obf_e4ce4cb0f09249fc.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_6fcba06a1228910e, __obf_2dd32e405269226e int, __obf_417ad15ad1451cfe []color.RGBA) *types.Palette {
	__obf_8e4fbb467ecd3e76 := make([]color.Color, 0, len(__obf_417ad15ad1451cfe)+1)
	__obf_8e4fbb467ecd3e76 = append(__obf_8e4fbb467ecd3e76, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_0fd3b7721858bd6f := range __obf_417ad15ad1451cfe {
		__obf_8e4fbb467ecd3e76 = append(__obf_8e4fbb467ecd3e76, __obf_0fd3b7721858bd6f)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_6fcba06a1228910e, __obf_2dd32e405269226e), __obf_8e4fbb467ecd3e76)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_6fcba06a1228910e, __obf_2dd32e405269226e int, __obf_d599ecd65e21f757 bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_6fcba06a1228910e, __obf_2dd32e405269226e), __obf_d599ecd65e21f757)
}
