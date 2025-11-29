package __obf_7b8e1f9e7332b319

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
func DrawString(__obf_232f42e4fdb825d2 draw.Image, __obf_0c4d2b7cfa206579 *types.DrawStringParam, __obf_b08be2dc2761a504 fixed.Point26_6) error {
	__obf_15786880837521ab, __obf_2d6a99a815edf3a4 := opentype.NewFace(__obf_0c4d2b7cfa206579.Font, &opentype.FaceOptions{
		Size:    float64(__obf_0c4d2b7cfa206579.FontSize),
		DPI:     float64(__obf_0c4d2b7cfa206579.FontDPI),
		Hinting: __obf_0c4d2b7cfa206579.FontHinting,
	})
	if __obf_2d6a99a815edf3a4 != nil {
		return __obf_2d6a99a815edf3a4
	}
	defer __obf_15786880837521ab.Close()
	__obf_f922abc78153c150 := &font.Drawer{
		Dst:  __obf_232f42e4fdb825d2,
		Src:  image.NewUniform(__obf_0c4d2b7cfa206579.Color),
		Face: __obf_15786880837521ab,
		Dot:  __obf_b08be2dc2761a504,
	}
	__obf_f922abc78153c150.
		DrawString(__obf_0c4d2b7cfa206579.Text)

	return nil
}

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(__obf_ad8f5349890184e7, __obf_fdd2437fdf307821 int, __obf_51169af6a5e1d9a0 []color.RGBA) *types.Palette {
	__obf_232f42e4fdb825d2 := make([]color.Color, 0, len(__obf_51169af6a5e1d9a0)+1)
	__obf_232f42e4fdb825d2 = append(__obf_232f42e4fdb825d2, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, __obf_deaf34ad3ed01b57 := range __obf_51169af6a5e1d9a0 {
		__obf_232f42e4fdb825d2 = append(__obf_232f42e4fdb825d2, __obf_deaf34ad3ed01b57)
	}

	return types.NewPalette(image.Rect(0, 0, __obf_ad8f5349890184e7, __obf_fdd2437fdf307821), __obf_232f42e4fdb825d2)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(__obf_ad8f5349890184e7, __obf_fdd2437fdf307821 int, __obf_c419fc582f7de9be bool) *types.NRGBA {
	return types.NewNRGBA(image.Rect(0, 0, __obf_ad8f5349890184e7, __obf_fdd2437fdf307821), __obf_c419fc582f7de9be)
}
