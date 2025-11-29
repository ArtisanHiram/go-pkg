package __obf_f075990603c6fd45

import (
	assets "github.com/ArtisanHiram/go-pkg/captcha/internal/assets"
	helper "github.com/ArtisanHiram/go-pkg/captcha/internal/helper"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"math"
	"math/rand/v2"
)

type Captcha struct {
	__obf_bd5a0fd2a490b814 *Options
}

var _ types.Captcha = (*Captcha)(nil)

func NewCaptcha(__obf_bd5a0fd2a490b814 *Options) *Captcha {
	if __obf_bd5a0fd2a490b814 == nil {
		return nil
	}
	return &Captcha{__obf_bd5a0fd2a490b814}
}

func (__obf_b3aecd711f830444 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_0eedbbe6c84cee22 := __obf_5400ac0dc0a0e004(__obf_b3aecd711f830444.__obf_bd5a0fd2a490b814)
	if __obf_0eedbbe6c84cee22 == Shape {
		return __obf_b3aecd711f830444.__obf_513845481c55c482()
	}

	return __obf_b3aecd711f830444.__obf_a3822403d0118f05()
}

func __obf_5400ac0dc0a0e004(__obf_bd5a0fd2a490b814 *Options) DotType {
	__obf_65ba21db855a9127 := DotType(rand.IntN(2))
	if __obf_65ba21db855a9127 == Shape {
		__obf_bd5a0fd2a490b814.
			UseRGBA = false
		__obf_bd5a0fd2a490b814.
			Secondary.BgDistort = types.DistortLevel1
		__obf_bd5a0fd2a490b814.
			Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_bd5a0fd2a490b814.
			Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_bd5a0fd2a490b814.
			UseRGBA = true
		__obf_bd5a0fd2a490b814.
			Secondary.BgDistort = types.DistortLevel4
		__obf_bd5a0fd2a490b814.
			Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_bd5a0fd2a490b814.
			Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_65ba21db855a9127
}

func (__obf_fb2f9bdfdc07de19 *Captcha) DrawWithNRGBA(__obf_0eedbbe6c84cee22 DotType, __obf_f68ed5589ee6f148 *DrawImageParams) (image.Image, error) {
	__obf_e0fb28c1ab249b67 := __obf_f68ed5589ee6f148.Dots
	__obf_f81bc3cb9ca665c8 := helper.CreateNRGBACanvas(__obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height, true)

	for _, __obf_a63fac8b033b6a09 := range __obf_e0fb28c1ab249b67 {
		__obf_6e09dad06375607e, __obf_aad283b20a07afeb, __obf_6090e69b7a7b925f := __obf_fb2f9bdfdc07de19.DrawDotImage(__obf_0eedbbe6c84cee22, __obf_a63fac8b033b6a09, __obf_f68ed5589ee6f148)
		if __obf_6090e69b7a7b925f != nil {
			return nil, __obf_6090e69b7a7b925f
		}
		__obf_5c6f1e1cfa61d6ef := __obf_aad283b20a07afeb.MinX
		__obf_e7e5ce13f22d1462 := __obf_aad283b20a07afeb.MaxX
		__obf_440cb6e116206d32 := __obf_aad283b20a07afeb.MinY
		__obf_5417d0e695d2e4de := __obf_aad283b20a07afeb.MaxY
		__obf_06628ab133c4f92c := __obf_e7e5ce13f22d1462 - __obf_5c6f1e1cfa61d6ef
		__obf_baee50d2555425be := __obf_5417d0e695d2e4de - __obf_440cb6e116206d32

		draw.Draw(__obf_f81bc3cb9ca665c8, image.Rect(__obf_a63fac8b033b6a09.X, __obf_a63fac8b033b6a09.Y, __obf_a63fac8b033b6a09.X+__obf_06628ab133c4f92c, __obf_a63fac8b033b6a09.Y+__obf_baee50d2555425be), __obf_6e09dad06375607e, image.Point{X: __obf_5c6f1e1cfa61d6ef, Y: __obf_440cb6e116206d32}, draw.Over)

	}
	__obf_533983578d64921e := __obf_f68ed5589ee6f148.Background
	__obf_a06164ee30eab17c := __obf_f81bc3cb9ca665c8.Bounds()
	__obf_6336b1f510798ac5 := helper.CreateNRGBACanvas(__obf_a06164ee30eab17c.Dx(), __obf_a06164ee30eab17c.Dx(), true)
	__obf_2d9dfcddd995cb20 := util.RangCutImagePos(__obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height, __obf_533983578d64921e)
	draw.Draw(__obf_6336b1f510798ac5, __obf_a06164ee30eab17c, __obf_533983578d64921e, __obf_2d9dfcddd995cb20, draw.Src)
	draw.Draw(__obf_6336b1f510798ac5, __obf_f81bc3cb9ca665c8.Bounds(), __obf_f81bc3cb9ca665c8, image.Point{}, draw.Over)
	__obf_6336b1f510798ac5.
		SubImage(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height))
	return __obf_6336b1f510798ac5, nil
}

func (__obf_fb2f9bdfdc07de19 *Captcha) DrawWithPalette(__obf_0eedbbe6c84cee22 DotType, __obf_f68ed5589ee6f148 *DrawImageParams, __obf_3093ed4cccfa7613 []color.Color, __obf_b6802dbfd4e5d0aa []color.Color) (image.Image, error) {
	__obf_e0fb28c1ab249b67 := __obf_f68ed5589ee6f148.Dots
	__obf_cb85e20ced5ecd3c := make([]color.Color, 0, len(__obf_b6802dbfd4e5d0aa))
	for _, __obf_aac71c89c4ff3276 := range __obf_b6802dbfd4e5d0aa {
		__obf_a96f563dbf6bc266, __obf_6f0a24f140a66d23, __obf_a06164ee30eab17c, _ := __obf_aac71c89c4ff3276.RGBA()
		__obf_30da8afc5cf528cf := util.FormatAlpha(__obf_f68ed5589ee6f148.ThumbDisturbAlpha)
		__obf_cb85e20ced5ecd3c = append(__obf_cb85e20ced5ecd3c, color.RGBA{R: uint8(__obf_a96f563dbf6bc266), G: uint8(__obf_6f0a24f140a66d23), B: uint8(__obf_a06164ee30eab17c), A: __obf_30da8afc5cf528cf})
	}

	var __obf_6dca8b8d3f8eea5b = make([]color.Color, 0, len(__obf_3093ed4cccfa7613)+len(__obf_cb85e20ced5ecd3c)+1)
	__obf_6dca8b8d3f8eea5b = append(__obf_6dca8b8d3f8eea5b, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_6dca8b8d3f8eea5b = append(__obf_6dca8b8d3f8eea5b, __obf_3093ed4cccfa7613...)
	__obf_6dca8b8d3f8eea5b = append(__obf_6dca8b8d3f8eea5b, __obf_cb85e20ced5ecd3c...)
	__obf_f81bc3cb9ca665c8 := types.NewPalette(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height), __obf_6dca8b8d3f8eea5b)
	if __obf_f68ed5589ee6f148.BgCircles > 0 {
		__obf_fb2f9bdfdc07de19.__obf_2063ad9b308c2d42(__obf_f81bc3cb9ca665c8, __obf_f68ed5589ee6f148.BgCircles, 1, __obf_cb85e20ced5ecd3c)
	}
	if __obf_f68ed5589ee6f148.BgSlimLines > 0 {
		__obf_fb2f9bdfdc07de19.__obf_83c0ddefd3976efb(__obf_f81bc3cb9ca665c8, __obf_f68ed5589ee6f148.BgSlimLines, __obf_cb85e20ced5ecd3c)
	}

	for _, __obf_a63fac8b033b6a09 := range __obf_e0fb28c1ab249b67 {
		__obf_ef9930496338d2fc, _ := util.ParseHexColor(__obf_a63fac8b033b6a09.PrimaryColor)
		var __obf_6090e69b7a7b925f error
		if __obf_0eedbbe6c84cee22 == Shape {
			var __obf_6e09dad06375607e *types.NRGBA
			__obf_6e09dad06375607e, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawShapeImage(__obf_a63fac8b033b6a09, __obf_ef9930496338d2fc)
			if __obf_6090e69b7a7b925f != nil {
				return nil, __obf_6090e69b7a7b925f
			}
			__obf_6e09dad06375607e.
				Rotate(__obf_a63fac8b033b6a09.Angle, false)
			__obf_8785a0c1cfcb08ba := __obf_f81bc3cb9ca665c8.Bounds()
			__obf_ef670a20257d6bb0 := __obf_6e09dad06375607e.Bounds()
			__obf_e0b8abf91c44ed15 := image.Point{X: __obf_8785a0c1cfcb08ba.Dx() - __obf_ef670a20257d6bb0.Dx(), Y: __obf_8785a0c1cfcb08ba.Dy() - __obf_ef670a20257d6bb0.Dy()}
			draw.Draw(__obf_f81bc3cb9ca665c8, image.Rect(__obf_a63fac8b033b6a09.X, __obf_e0b8abf91c44ed15.Y, __obf_e0b8abf91c44ed15.X+__obf_ef670a20257d6bb0.Dx(), __obf_e0b8abf91c44ed15.Y+__obf_ef670a20257d6bb0.Dy()), __obf_6e09dad06375607e, image.Point{}, draw.Over)
		} else {
			__obf_6c051c6050cf1745 := fixed.P(__obf_a63fac8b033b6a09.X, __obf_a63fac8b033b6a09.Y)
			__obf_f68ed5589ee6f148 := &types.DrawStringParam{
				Color:       __obf_ef9930496338d2fc,
				FontSize:    __obf_a63fac8b033b6a09.FontSize,
				Size:        __obf_a63fac8b033b6a09.Size,
				FontDPI:     __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.FontDPI,
				FontHinting: __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.FontHinting,
				Text:        __obf_a63fac8b033b6a09.Text,
			}
			__obf_f68ed5589ee6f148.
				Font, __obf_6090e69b7a7b925f = assets.PickFont()
			if __obf_6090e69b7a7b925f != nil {
				return nil, __obf_6090e69b7a7b925f
			}
			__obf_6090e69b7a7b925f = helper.DrawString(__obf_f81bc3cb9ca665c8, __obf_f68ed5589ee6f148, __obf_6c051c6050cf1745)
		}

		if __obf_6090e69b7a7b925f != nil {
			return __obf_f81bc3cb9ca665c8, __obf_6090e69b7a7b925f
		}
	}

	if __obf_f68ed5589ee6f148.Background != nil {
		__obf_533983578d64921e := __obf_f68ed5589ee6f148.Background
		__obf_a06164ee30eab17c := __obf_533983578d64921e.Bounds()
		__obf_6336b1f510798ac5 := helper.CreateNRGBACanvas(__obf_a06164ee30eab17c.Dx(), __obf_a06164ee30eab17c.Dy(), true)
		__obf_2d9dfcddd995cb20 := util.RangCutImagePos(__obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height, __obf_533983578d64921e)
		draw.Draw(__obf_6336b1f510798ac5, __obf_a06164ee30eab17c, __obf_533983578d64921e, __obf_2d9dfcddd995cb20, draw.Src)
		__obf_f81bc3cb9ca665c8.
			Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_6336b1f510798ac5, __obf_f81bc3cb9ca665c8.Bounds(), __obf_f81bc3cb9ca665c8, image.Point{}, draw.Over)
		__obf_14a37f210c305f6c := __obf_6336b1f510798ac5.SubImage(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height)).(*image.NRGBA)
		return __obf_14a37f210c305f6c, nil
	}

	if __obf_f68ed5589ee6f148.BackgroundDistort > 0 {
		__obf_f81bc3cb9ca665c8.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_f68ed5589ee6f148.BackgroundDistort))
	}

	return __obf_f81bc3cb9ca665c8, nil

}

func (__obf_fb2f9bdfdc07de19 *Captcha) DrawWithNRGBA2(__obf_0eedbbe6c84cee22 DotType, __obf_f68ed5589ee6f148 *DrawImageParams, __obf_3093ed4cccfa7613 []color.Color, __obf_b6802dbfd4e5d0aa []color.Color) (image.Image, error) {
	__obf_e0fb28c1ab249b67 := __obf_f68ed5589ee6f148.Dots
	__obf_cb85e20ced5ecd3c := make([]color.Color, 0, len(__obf_b6802dbfd4e5d0aa))
	for _, __obf_aac71c89c4ff3276 := range __obf_b6802dbfd4e5d0aa {
		__obf_a96f563dbf6bc266, __obf_6f0a24f140a66d23, __obf_a06164ee30eab17c, _ := __obf_aac71c89c4ff3276.RGBA()
		__obf_30da8afc5cf528cf := util.FormatAlpha(__obf_f68ed5589ee6f148.ThumbDisturbAlpha)
		__obf_cb85e20ced5ecd3c = append(__obf_cb85e20ced5ecd3c, color.RGBA{R: uint8(__obf_a96f563dbf6bc266), G: uint8(__obf_6f0a24f140a66d23), B: uint8(__obf_a06164ee30eab17c), A: __obf_30da8afc5cf528cf})
	}

	var __obf_6dca8b8d3f8eea5b = make([]color.Color, 0, len(__obf_3093ed4cccfa7613)+len(__obf_cb85e20ced5ecd3c)+1)
	__obf_6dca8b8d3f8eea5b = append(__obf_6dca8b8d3f8eea5b, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_6dca8b8d3f8eea5b = append(__obf_6dca8b8d3f8eea5b, __obf_3093ed4cccfa7613...)
	__obf_6dca8b8d3f8eea5b = append(__obf_6dca8b8d3f8eea5b, __obf_cb85e20ced5ecd3c...)
	__obf_bffa9ae2ae857705 := types.NewNRGBA(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height), true)
	if __obf_f68ed5589ee6f148.Background != nil {
		__obf_533983578d64921e := __obf_f68ed5589ee6f148.Background
		__obf_a06164ee30eab17c := __obf_533983578d64921e.Bounds()
		__obf_6336b1f510798ac5 := helper.CreateNRGBACanvas(__obf_a06164ee30eab17c.Dx(), __obf_a06164ee30eab17c.Dy(), true)
		__obf_2d9dfcddd995cb20 := util.RangCutImagePos(__obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height, __obf_533983578d64921e)
		draw.Draw(__obf_6336b1f510798ac5, __obf_a06164ee30eab17c, __obf_533983578d64921e, __obf_2d9dfcddd995cb20, draw.Src)
		__obf_14a37f210c305f6c := __obf_6336b1f510798ac5.SubImage(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height)).(*image.NRGBA)
		draw.Draw(__obf_bffa9ae2ae857705, __obf_14a37f210c305f6c.Bounds(), __obf_14a37f210c305f6c, image.Point{}, draw.Over)
	}
	__obf_f81bc3cb9ca665c8 := types.NewPalette(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height), __obf_6dca8b8d3f8eea5b)
	if __obf_f68ed5589ee6f148.BgCircles > 0 {
		__obf_fb2f9bdfdc07de19.__obf_2063ad9b308c2d42(__obf_f81bc3cb9ca665c8, __obf_f68ed5589ee6f148.BgCircles, 1, __obf_cb85e20ced5ecd3c)
	}
	if __obf_f68ed5589ee6f148.BgSlimLines > 0 {
		__obf_fb2f9bdfdc07de19.__obf_83c0ddefd3976efb(__obf_f81bc3cb9ca665c8, __obf_f68ed5589ee6f148.BgSlimLines, __obf_cb85e20ced5ecd3c)
	}
	if __obf_f68ed5589ee6f148.BackgroundDistort > 0 {
		__obf_f81bc3cb9ca665c8.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_f68ed5589ee6f148.BackgroundDistort))
	}
	__obf_204e9d4e93dd7d17 := __obf_f81bc3cb9ca665c8.Bounds()
	__obf_06628ab133c4f92c := __obf_204e9d4e93dd7d17.Dx() / len(__obf_e0fb28c1ab249b67)
	for __obf_339c01a0f79fa1a4, __obf_a63fac8b033b6a09 := range __obf_e0fb28c1ab249b67 {
		__obf_ef9930496338d2fc, _ := util.ParseHexColor(__obf_a63fac8b033b6a09.PrimaryColor)
		var __obf_6090e69b7a7b925f error
		if __obf_0eedbbe6c84cee22 == Shape {
			var __obf_6e09dad06375607e *types.NRGBA
			__obf_6e09dad06375607e, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawShapeImage(__obf_a63fac8b033b6a09, __obf_ef9930496338d2fc)
			if __obf_6090e69b7a7b925f != nil {
				return nil, __obf_6090e69b7a7b925f
			}
			__obf_6e09dad06375607e.
				Rotate(__obf_a63fac8b033b6a09.Angle, false)
			__obf_8785a0c1cfcb08ba := __obf_bffa9ae2ae857705.Bounds()
			__obf_ef670a20257d6bb0 := __obf_6e09dad06375607e.Bounds()
			__obf_e0b8abf91c44ed15 := image.Point{X: __obf_8785a0c1cfcb08ba.Dx() - __obf_ef670a20257d6bb0.Dx(), Y: __obf_8785a0c1cfcb08ba.Dy() - __obf_ef670a20257d6bb0.Dy()}
			draw.Draw(__obf_bffa9ae2ae857705, image.Rect(__obf_a63fac8b033b6a09.X, __obf_e0b8abf91c44ed15.Y, __obf_e0b8abf91c44ed15.X+__obf_ef670a20257d6bb0.Dx(), __obf_e0b8abf91c44ed15.Y+__obf_ef670a20257d6bb0.Dy()), __obf_6e09dad06375607e, image.Point{}, draw.Over)
		} else {
			var __obf_76d1fc08a8e11e82 *types.NRGBA
			__obf_76d1fc08a8e11e82, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawStringImage(__obf_a63fac8b033b6a09, __obf_ef9930496338d2fc)
			if __obf_6090e69b7a7b925f != nil {
				return nil, __obf_6090e69b7a7b925f
			}
			__obf_76d1fc08a8e11e82.
				Rotate(__obf_a63fac8b033b6a09.Angle, false)
			__obf_aad283b20a07afeb := __obf_76d1fc08a8e11e82.CalcMarginBlankArea()
			__obf_5c6f1e1cfa61d6ef := __obf_aad283b20a07afeb.MinX
			__obf_e7e5ce13f22d1462 := __obf_aad283b20a07afeb.MaxX
			__obf_440cb6e116206d32 := __obf_aad283b20a07afeb.MinY
			__obf_5417d0e695d2e4de := __obf_aad283b20a07afeb.MaxY
			__obf_76d1fc08a8e11e82.
				SubImage(image.Rect(__obf_5c6f1e1cfa61d6ef, __obf_440cb6e116206d32, __obf_e7e5ce13f22d1462, __obf_5417d0e695d2e4de))
			__obf_d6be6a387b246965 := __obf_76d1fc08a8e11e82.Bounds()
			__obf_6a7713303fd0b751 := int(max(float64(__obf_06628ab133c4f92c*__obf_339c01a0f79fa1a4+__obf_06628ab133c4f92c/__obf_d6be6a387b246965.Dx()), 8))
			__obf_6c666fe1e4aefe76 := util.RandInt(1, __obf_204e9d4e93dd7d17.Dy()-__obf_d6be6a387b246965.Dy()-4)

			draw.Draw(__obf_bffa9ae2ae857705, image.Rect(__obf_6a7713303fd0b751, __obf_6c666fe1e4aefe76, __obf_6a7713303fd0b751+__obf_d6be6a387b246965.Dx(), __obf_6c666fe1e4aefe76+__obf_d6be6a387b246965.Dy()), __obf_76d1fc08a8e11e82, image.Point{X: __obf_d6be6a387b246965.Min.X, Y: __obf_d6be6a387b246965.Min.Y}, draw.Over)
		}

		if __obf_6090e69b7a7b925f != nil {
			return __obf_bffa9ae2ae857705, __obf_6090e69b7a7b925f
		}
	}
	__obf_e3ccf6f4ae246502 := types.NewNRGBA(image.Rect(0, 0, __obf_f68ed5589ee6f148.Width, __obf_f68ed5589ee6f148.Height), true)
	draw.Draw(__obf_e3ccf6f4ae246502, __obf_f81bc3cb9ca665c8.Bounds(), __obf_f81bc3cb9ca665c8, image.Point{}, draw.Over)
	draw.Draw(__obf_bffa9ae2ae857705, __obf_e3ccf6f4ae246502.Bounds(), __obf_e3ccf6f4ae246502, image.Point{}, draw.Over)
	return __obf_bffa9ae2ae857705, nil
}

func (__obf_fb2f9bdfdc07de19 *Captcha) __obf_2063ad9b308c2d42(__obf_6336b1f510798ac5 *types.Palette, __obf_4eef14fe3715bd2b, __obf_f39506479fb2e530 int, __obf_dcb830e5c18849ed []color.Color) {
	__obf_79e34f28716734ff := __obf_6336b1f510798ac5.Bounds().Max.X
	__obf_9381141c635259a8 := __obf_6336b1f510798ac5.Bounds().Max.Y
	for range __obf_4eef14fe3715bd2b {
		__obf_5dc2bba1e365f992 := util.RandColor(__obf_dcb830e5c18849ed)
		__obf_a96f563dbf6bc266 := //co.A = uint8(0xee)
			util.RandInt(1, __obf_f39506479fb2e530)
		__obf_6336b1f510798ac5.
			DrawCircle(util.RandInt(__obf_a96f563dbf6bc266, __obf_79e34f28716734ff-__obf_a96f563dbf6bc266), util.RandInt(__obf_a96f563dbf6bc266, __obf_9381141c635259a8-__obf_a96f563dbf6bc266), __obf_a96f563dbf6bc266, __obf_5dc2bba1e365f992)
	}
}

func (__obf_fb2f9bdfdc07de19 *Captcha) __obf_83c0ddefd3976efb(__obf_6336b1f510798ac5 *types.Palette, __obf_3482c2f0caa937bb int, __obf_dcb830e5c18849ed []color.Color) {
	__obf_e7bf24d8bd22f728 := __obf_6336b1f510798ac5.Bounds().Max.X / 10
	__obf_d31ff4e6f2fccd0b := __obf_e7bf24d8bd22f728 * 9
	__obf_9fababd947259d6e := __obf_6336b1f510798ac5.Bounds().Max.Y / 3
	for __obf_339c01a0f79fa1a4 := range __obf_3482c2f0caa937bb {
		__obf_d8f8205355e7b3f7 := image.Point{X: rand.IntN(__obf_e7bf24d8bd22f728), Y: rand.IntN(__obf_9fababd947259d6e)}
		__obf_1576662b93a39f44 := image.Point{X: rand.IntN(__obf_e7bf24d8bd22f728) + __obf_d31ff4e6f2fccd0b, Y: rand.IntN(__obf_9fababd947259d6e)}

		if __obf_339c01a0f79fa1a4%2 == 0 {
			__obf_d8f8205355e7b3f7.
				Y = rand.IntN(__obf_9fababd947259d6e) + __obf_9fababd947259d6e*2
			__obf_1576662b93a39f44.
				Y = rand.IntN(__obf_9fababd947259d6e)
		} else {
			__obf_d8f8205355e7b3f7.
				Y = rand.IntN(__obf_9fababd947259d6e) + __obf_9fababd947259d6e*(__obf_339c01a0f79fa1a4%2)
			__obf_1576662b93a39f44.
				Y = rand.IntN(__obf_9fababd947259d6e) + __obf_9fababd947259d6e*2
		}
		__obf_5dc2bba1e365f992 := util.RandColor(__obf_dcb830e5c18849ed)
		__obf_6336b1f510798ac5.
			//co.A = uint8(0xee)
			DrawBeeline(__obf_d8f8205355e7b3f7, __obf_1576662b93a39f44, __obf_5dc2bba1e365f992)
	}
}

func (__obf_fb2f9bdfdc07de19 *Captcha) DrawDotImage(__obf_0eedbbe6c84cee22 DotType, __obf_c6ea6b2418e24fc3 *types.Dot, __obf_f68ed5589ee6f148 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_ef9930496338d2fc, _ := util.ParseHexColor(__obf_c6ea6b2418e24fc3.PrimaryColor)
	__obf_ef9930496338d2fc.
		A = util.FormatAlpha(__obf_f68ed5589ee6f148.Alpha)

	var __obf_76d1fc08a8e11e82 image.Image
	var __obf_6090e69b7a7b925f error
	if __obf_0eedbbe6c84cee22 == Shape {
		__obf_76d1fc08a8e11e82, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawShapeImage(__obf_c6ea6b2418e24fc3, __obf_ef9930496338d2fc)
	} else {
		__obf_76d1fc08a8e11e82, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawStringImage(__obf_c6ea6b2418e24fc3, __obf_ef9930496338d2fc)
	}
	if __obf_6090e69b7a7b925f != nil {
		return nil, nil, __obf_6090e69b7a7b925f
	}
	__obf_480abdca04e39ee2, _ := util.ParseHexColor(__obf_e5fb09c4c592d9f0)
	__obf_f81bc3cb9ca665c8 := helper.CreateNRGBACanvas(__obf_c6ea6b2418e24fc3.Size+10, __obf_c6ea6b2418e24fc3.Size+10, true)
	if __obf_f68ed5589ee6f148.ShowShadow {
		var __obf_566468cd811cbeb2 *types.NRGBA
		if __obf_0eedbbe6c84cee22 == Shape {
			__obf_566468cd811cbeb2, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawShapeImage(__obf_c6ea6b2418e24fc3, __obf_480abdca04e39ee2)
		} else {
			__obf_566468cd811cbeb2, __obf_6090e69b7a7b925f = __obf_fb2f9bdfdc07de19.DrawStringImage(__obf_c6ea6b2418e24fc3, __obf_480abdca04e39ee2)
		}
		if __obf_6090e69b7a7b925f != nil {
			return nil, nil, __obf_6090e69b7a7b925f
		}
		__obf_2f8a735f1dd97eef := __obf_f68ed5589ee6f148.ShadowPoint.X
		__obf_9b0fe539c88969e9 := __obf_f68ed5589ee6f148.ShadowPoint.Y
		draw.Draw(__obf_f81bc3cb9ca665c8, __obf_566468cd811cbeb2.Bounds(), __obf_566468cd811cbeb2, image.Point{X: __obf_2f8a735f1dd97eef, Y: __obf_9b0fe539c88969e9}, draw.Over)
	}
	draw.Draw(__obf_f81bc3cb9ca665c8, __obf_76d1fc08a8e11e82.Bounds(), __obf_76d1fc08a8e11e82, image.Point{}, draw.Over)
	__obf_f81bc3cb9ca665c8.
		Rotate(__obf_c6ea6b2418e24fc3.Angle, false)
	__obf_84b56b17821bde93 := __obf_f81bc3cb9ca665c8.CalcMarginBlankArea()

	return __obf_f81bc3cb9ca665c8,

		// DrawStringImage draws a text image
		// params:
		//   - dot: Draw dot
		//   - textColor: Text color
		//
		// returns:
		//   - types.NRGBA: Drawn text image
		//   - error: Error information
		__obf_84b56b17821bde93, nil
}

func (__obf_fb2f9bdfdc07de19 *Captcha) DrawStringImage(__obf_c6ea6b2418e24fc3 *types.Dot, __obf_c3e514899de8d5b6 color.Color) (*types.NRGBA, error) {
	__obf_f81bc3cb9ca665c8 := helper.CreateNRGBACanvas(__obf_c6ea6b2418e24fc3.Size+10, __obf_c6ea6b2418e24fc3.Size+10, true)
	__obf_6c051c6050cf1745 := fixed.P(12, __obf_c6ea6b2418e24fc3.Size-5)
	if util.IsChineseChar(__obf_c6ea6b2418e24fc3.Text) {
		__obf_6c051c6050cf1745 = fixed.P(10, __obf_c6ea6b2418e24fc3.Size)
	}
	__obf_f68ed5589ee6f148 := &types.DrawStringParam{
		Color:       __obf_c3e514899de8d5b6,
		FontSize:    __obf_c6ea6b2418e24fc3.FontSize,
		Size:        __obf_c6ea6b2418e24fc3.Size,
		FontDPI:     __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.FontDPI,
		FontHinting: __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.FontHinting,
		Text:        __obf_c6ea6b2418e24fc3.Text,
	}
	var __obf_6090e69b7a7b925f error
	__obf_f68ed5589ee6f148.
		Font, __obf_6090e69b7a7b925f = assets.PickFont()
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}
	__obf_6090e69b7a7b925f = helper.DrawString(__obf_f81bc3cb9ca665c8, __obf_f68ed5589ee6f148, __obf_6c051c6050cf1745)
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}

	return __obf_f81bc3cb9ca665c8, nil
}

func (__obf_fb2f9bdfdc07de19 *Captcha) DrawShapeImage(__obf_c6ea6b2418e24fc3 *types.Dot, __obf_ef9930496338d2fc color.Color) (*types.NRGBA, error) {
	__obf_ce8b0e2b227ae6a8, __obf_ebf9db201ca295cf, __obf_180a4af6d05d02fc, __obf_9dab6c8b645ae3ac := __obf_ef9930496338d2fc.RGBA()

	var __obf_27d8968b8907fef2 = []color.RGBA{
		{R: uint8(__obf_ce8b0e2b227ae6a8), G: uint8(__obf_ebf9db201ca295cf), B: uint8(__obf_180a4af6d05d02fc), A: uint8(__obf_9dab6c8b645ae3ac)},
	}
	__obf_e3ccf6f4ae246502 := helper.CreateNRGBACanvas(__obf_c6ea6b2418e24fc3.Size+10, __obf_c6ea6b2418e24fc3.Size+10, true)
	var __obf_d6be6a387b246965 image.Rectangle
	var __obf_533983578d64921e image.Image
	if __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.UseRGBA {
		__obf_f81bc3cb9ca665c8 := helper.CreateNRGBACanvas(__obf_c6ea6b2418e24fc3.Size+10, __obf_c6ea6b2418e24fc3.Size+10, true)
		draw.BiLinear.Scale(__obf_f81bc3cb9ca665c8, __obf_f81bc3cb9ca665c8.Bounds(), __obf_c6ea6b2418e24fc3.Shape, __obf_c6ea6b2418e24fc3.Shape.Bounds(), draw.Over, nil)
		__obf_d6be6a387b246965 = __obf_f81bc3cb9ca665c8.Bounds()
		__obf_533983578d64921e = __obf_f81bc3cb9ca665c8
	} else {
		__obf_f81bc3cb9ca665c8 := helper.CreatePaletteCanvas(__obf_c6ea6b2418e24fc3.Size+10, __obf_c6ea6b2418e24fc3.Size+10, __obf_27d8968b8907fef2)
		draw.BiLinear.Scale(__obf_f81bc3cb9ca665c8, __obf_f81bc3cb9ca665c8.Bounds(), __obf_c6ea6b2418e24fc3.Shape, __obf_c6ea6b2418e24fc3.Shape.Bounds(), draw.Over, nil)
		__obf_d6be6a387b246965 = __obf_f81bc3cb9ca665c8.Bounds()
		__obf_533983578d64921e = __obf_f81bc3cb9ca665c8
	}

	draw.Draw(__obf_e3ccf6f4ae246502, __obf_d6be6a387b246965, __obf_533983578d64921e, image.Point{}, draw.Over)

	return __obf_e3ccf6f4ae246502, nil
}

func (__obf_b3aecd711f830444 *Captcha) __obf_513845481c55c482() (*CaptchaData, error) {

	var (
		__obf_12d21cebce924097, __obf_8d67160faddf1ee5, __obf_db19ab5183be17c0 []*types.Dot
		__obf_4622ff46274809c8                                                 []image.Image
		__obf_dcbc86a5a13adc88                                                 []image.Image
		__obf_8eb85ea5aa7dca32, __obf_5e5bc037f5f349e0                         image.Image
		__obf_6090e69b7a7b925f                                                 error
	)
	__obf_dcbc86a5a13adc88, __obf_6090e69b7a7b925f = assets.GetShapes()
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}
	__obf_fd906028639fa50b := util.RandInt(__obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.LenRange.Min, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.LenRange.Max)
	__obf_dcbc86a5a13adc88 = util.PickN(__obf_dcbc86a5a13adc88, __obf_fd906028639fa50b)
	__obf_12d21cebce924097 = __obf_b3aecd711f830444.__obf_9da4914226e857e5(__obf_dcbc86a5a13adc88, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.Size, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.SizeRange, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.DotPadding)
	__obf_7751d18d174c7957 := rand.Perm(len(__obf_12d21cebce924097))
	__obf_212bb7a675794b3e := util.RandInt(__obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.VerifyLenRange.Min, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.VerifyLenRange.Max)
	__obf_db19ab5183be17c0 = make([]*types.Dot, __obf_212bb7a675794b3e)
	for __obf_339c01a0f79fa1a4, __obf_02d2a2ce4d3b1dd7 := range __obf_7751d18d174c7957 {
		if __obf_339c01a0f79fa1a4 >= __obf_212bb7a675794b3e {
			break
		}
		__obf_c6ea6b2418e24fc3 := __obf_12d21cebce924097[__obf_02d2a2ce4d3b1dd7]
		__obf_c6ea6b2418e24fc3.
			Index = __obf_339c01a0f79fa1a4
		__obf_db19ab5183be17c0[__obf_339c01a0f79fa1a4] = __obf_c6ea6b2418e24fc3
		__obf_4622ff46274809c8 = append(__obf_4622ff46274809c8, __obf_db19ab5183be17c0[__obf_339c01a0f79fa1a4].Shape)
	}
	__obf_8d67160faddf1ee5 = __obf_b3aecd711f830444.__obf_9da4914226e857e5(__obf_4622ff46274809c8, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.Size, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.SizeRange, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.DotPadding)
	__obf_8eb85ea5aa7dca32, __obf_6090e69b7a7b925f = __obf_b3aecd711f830444.__obf_f30dfd6a7024c455(Shape, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.Size, __obf_12d21cebce924097)
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}
	__obf_5e5bc037f5f349e0, __obf_6090e69b7a7b925f = __obf_b3aecd711f830444.__obf_11ca91c9d4fc9998(Shape, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.Size, __obf_8d67160faddf1ee5)
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}

	return &CaptchaData{__obf_e0fb28c1ab249b67: __obf_db19ab5183be17c0, __obf_8eb85ea5aa7dca32: types.NewJPEGImage(__obf_8eb85ea5aa7dca32), __obf_5e5bc037f5f349e0: types.NewPNGImage(__obf_5e5bc037f5f349e0)}, nil
}

func (__obf_b3aecd711f830444 *Captcha) __obf_a3822403d0118f05() (*CaptchaData, error) {
	__obf_fd906028639fa50b := util.RandInt(__obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.LenRange.Min, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.LenRange.Max)
	__obf_d928ca389b478fa4 := util.PickN(assets.GetChineseChars(), __obf_fd906028639fa50b)

	var __obf_e0fb28c1ab249b67, __obf_8d67160faddf1ee5, __obf_db19ab5183be17c0 []*types.Dot
	var __obf_3659f39c3de82724 []string
	var __obf_8eb85ea5aa7dca32, __obf_5e5bc037f5f349e0 image.Image
	__obf_e0fb28c1ab249b67 = __obf_b3aecd711f830444.__obf_672f06a4b832c1f9(__obf_d928ca389b478fa4, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.Size, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.SizeRange, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.DotPadding)
	__obf_7751d18d174c7957 := rand.Perm(len(__obf_e0fb28c1ab249b67))
	__obf_212bb7a675794b3e := util.RandInt(__obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.VerifyLenRange.Min, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.VerifyLenRange.Max)
	__obf_db19ab5183be17c0 = make([]*types.Dot, __obf_212bb7a675794b3e)
	for __obf_339c01a0f79fa1a4, __obf_02d2a2ce4d3b1dd7 := range __obf_7751d18d174c7957 {
		if __obf_339c01a0f79fa1a4 >= __obf_212bb7a675794b3e {
			break
		}
		__obf_c6ea6b2418e24fc3 := __obf_e0fb28c1ab249b67[__obf_02d2a2ce4d3b1dd7]
		__obf_c6ea6b2418e24fc3.
			Index = __obf_339c01a0f79fa1a4
		__obf_db19ab5183be17c0[__obf_339c01a0f79fa1a4] = __obf_c6ea6b2418e24fc3
		__obf_3659f39c3de82724 = append(__obf_3659f39c3de82724, __obf_db19ab5183be17c0[__obf_339c01a0f79fa1a4].Text)
	}
	__obf_8d67160faddf1ee5 = __obf_b3aecd711f830444.__obf_672f06a4b832c1f9(__obf_3659f39c3de82724, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.Size, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.SizeRange, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.DotPadding)
	var __obf_6090e69b7a7b925f error
	__obf_8eb85ea5aa7dca32, __obf_6090e69b7a7b925f = __obf_b3aecd711f830444.__obf_f30dfd6a7024c455(Text, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.Size, __obf_e0fb28c1ab249b67)
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}
	__obf_5e5bc037f5f349e0, __obf_6090e69b7a7b925f = __obf_b3aecd711f830444.__obf_11ca91c9d4fc9998(Text, __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Secondary.Size, __obf_8d67160faddf1ee5)
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}

	return &CaptchaData{__obf_e0fb28c1ab249b67: __obf_db19ab5183be17c0, __obf_8eb85ea5aa7dca32: types.NewJPEGImage(__obf_8eb85ea5aa7dca32), __obf_5e5bc037f5f349e0: types.NewPNGImage(__obf_5e5bc037f5f349e0)}, nil
}

func (__obf_b3aecd711f830444 *Captcha) __obf_9da4914226e857e5(__obf_dcbc86a5a13adc88 []image.Image, __obf_04fb7cc8dc836c79 types.Size, __obf_3b41180d64875a8a types.Range, __obf_dda733cd16ccde62 int) []*types.Dot {
	__obf_fd906028639fa50b := len(__obf_dcbc86a5a13adc88)
	var __obf_e0fb28c1ab249b67 = make([]*types.Dot, __obf_fd906028639fa50b)
	__obf_06628ab133c4f92c := __obf_04fb7cc8dc836c79.Width
	__obf_baee50d2555425be := __obf_04fb7cc8dc836c79.Height
	if __obf_dda733cd16ccde62 > 0 {
		__obf_06628ab133c4f92c -= __obf_dda733cd16ccde62
		__obf_baee50d2555425be -= __obf_dda733cd16ccde62
	}

	for __obf_339c01a0f79fa1a4, __obf_a63fac8b033b6a09 := range __obf_dcbc86a5a13adc88 {
		__obf_f16b2da5958b1356 := __obf_b3aecd711f830444.__obf_f16b2da5958b1356()
		__obf_20c935ac6e8be4f2 := util.PickN(__obf_21256905e3e9ffec, 1)[0]
		__obf_37bb1e22996a5498 := util.PickN(__obf_c16ce6571c50b663, 1)[0]
		__obf_32cb3f5973ffed8e := util.RandInt(__obf_3b41180d64875a8a.Min, __obf_3b41180d64875a8a.Max)
		__obf_6c666fe1e4aefe76 := 10
		__obf_949c14a2357771fb := __obf_06628ab133c4f92c / __obf_fd906028639fa50b
		__obf_ac3c61739c9f5d1e := math.Abs(float64(__obf_949c14a2357771fb) - float64(__obf_32cb3f5973ffed8e))
		__obf_387b9d24c46825a7 := (__obf_339c01a0f79fa1a4 * __obf_949c14a2357771fb) + util.RandInt(0, int(max(__obf_ac3c61739c9f5d1e, 1)))
		__obf_c2d8443d1c97cbfe := util.RandInt(__obf_6c666fe1e4aefe76, __obf_baee50d2555425be+__obf_32cb3f5973ffed8e)
		__obf_7e2d9fc1bb1543c9 := int(min(max(float64(__obf_387b9d24c46825a7), float64(__obf_6c666fe1e4aefe76)), float64(__obf_06628ab133c4f92c-__obf_6c666fe1e4aefe76-(__obf_dda733cd16ccde62*2))))
		__obf_9fababd947259d6e := int(min(max(float64(__obf_c2d8443d1c97cbfe), float64(__obf_32cb3f5973ffed8e+__obf_6c666fe1e4aefe76)), float64(__obf_baee50d2555425be+(__obf_32cb3f5973ffed8e/2)-(__obf_dda733cd16ccde62*2))))
		__obf_e0fb28c1ab249b67[__obf_339c01a0f79fa1a4] = &types.Dot{
			Index:          __obf_339c01a0f79fa1a4,
			X:              __obf_7e2d9fc1bb1543c9,
			Y:              __obf_9fababd947259d6e - __obf_32cb3f5973ffed8e,
			FontSize:       __obf_32cb3f5973ffed8e,
			Size:           __obf_32cb3f5973ffed8e,
			Angle:          __obf_f16b2da5958b1356,
			PrimaryColor:   __obf_20c935ac6e8be4f2,
			SecondaryColor: __obf_37bb1e22996a5498,
			Shape:          __obf_a63fac8b033b6a09,
		}
	}

	return __obf_e0fb28c1ab249b67
}

func (__obf_b3aecd711f830444 *Captcha) __obf_672f06a4b832c1f9(__obf_d928ca389b478fa4 []string, __obf_04fb7cc8dc836c79 types.Size, __obf_5cc8cae6d500bcb3 types.Range, __obf_dda733cd16ccde62 int) []*types.Dot {
	__obf_fd906028639fa50b := len(__obf_d928ca389b478fa4)
	var __obf_e0fb28c1ab249b67 = make([]*types.Dot, __obf_fd906028639fa50b)
	__obf_06628ab133c4f92c := __obf_04fb7cc8dc836c79.Width
	__obf_baee50d2555425be := __obf_04fb7cc8dc836c79.Height
	if __obf_dda733cd16ccde62 > 0 {
		__obf_06628ab133c4f92c -= __obf_dda733cd16ccde62
		__obf_baee50d2555425be -= __obf_dda733cd16ccde62
	}

	for __obf_339c01a0f79fa1a4, __obf_a63fac8b033b6a09 := range __obf_d928ca389b478fa4 {
		__obf_f16b2da5958b1356 := __obf_b3aecd711f830444.__obf_f16b2da5958b1356()
		__obf_20c935ac6e8be4f2 := util.PickN(__obf_21256905e3e9ffec, 1)[0]
		__obf_37bb1e22996a5498 := util.PickN(__obf_c16ce6571c50b663, 1)[0]
		__obf_32cb3f5973ffed8e := util.RandInt(__obf_5cc8cae6d500bcb3.Min, __obf_5cc8cae6d500bcb3.Max)
		__obf_2246b324dc6504b4 := __obf_32cb3f5973ffed8e
		__obf_5a1bdb8d2b3ded3e := __obf_32cb3f5973ffed8e

		if util.LenChineseChar(__obf_a63fac8b033b6a09) > 1 {
			__obf_5a1bdb8d2b3ded3e = __obf_32cb3f5973ffed8e * util.LenChineseChar(__obf_a63fac8b033b6a09)

			if __obf_f16b2da5958b1356 > 0 {
				__obf_168f3b3292da0bda := __obf_5a1bdb8d2b3ded3e - __obf_32cb3f5973ffed8e
				__obf_9a581592e0fb88ea := __obf_f16b2da5958b1356 % 90
				__obf_1a862ef8e8e7167f := float64(__obf_168f3b3292da0bda) / 90
				__obf_a96f563dbf6bc266 := max(float64(__obf_9a581592e0fb88ea)*__obf_1a862ef8e8e7167f, 1)
				__obf_2246b324dc6504b4 = __obf_2246b324dc6504b4 + int(__obf_a96f563dbf6bc266)
				__obf_5a1bdb8d2b3ded3e = __obf_5a1bdb8d2b3ded3e + int(__obf_a96f563dbf6bc266)
			}
		}
		__obf_6c666fe1e4aefe76 := 10
		__obf_949c14a2357771fb := __obf_06628ab133c4f92c / __obf_fd906028639fa50b
		__obf_ac3c61739c9f5d1e := math.Abs(float64(__obf_949c14a2357771fb) - float64(__obf_5a1bdb8d2b3ded3e))
		__obf_387b9d24c46825a7 := (__obf_339c01a0f79fa1a4 * __obf_949c14a2357771fb) + util.RandInt(0, int(max(__obf_ac3c61739c9f5d1e, 1)))
		__obf_c2d8443d1c97cbfe := util.RandInt(__obf_6c666fe1e4aefe76, __obf_baee50d2555425be+__obf_2246b324dc6504b4)
		__obf_7e2d9fc1bb1543c9 := int(min(max(float64(__obf_387b9d24c46825a7), float64(__obf_6c666fe1e4aefe76)), float64(__obf_06628ab133c4f92c-__obf_6c666fe1e4aefe76-(__obf_dda733cd16ccde62*2))))
		__obf_9fababd947259d6e := int(min(max(float64(__obf_c2d8443d1c97cbfe), float64(__obf_2246b324dc6504b4+__obf_6c666fe1e4aefe76)), float64(__obf_baee50d2555425be+(__obf_2246b324dc6504b4/2)-(__obf_dda733cd16ccde62*2))))
		__obf_e0fb28c1ab249b67[__obf_339c01a0f79fa1a4] = &types.Dot{
			Index:          __obf_339c01a0f79fa1a4,
			X:              __obf_7e2d9fc1bb1543c9,
			Y:              __obf_9fababd947259d6e - __obf_2246b324dc6504b4,
			FontSize:       __obf_32cb3f5973ffed8e,
			Size:           max(__obf_5a1bdb8d2b3ded3e, __obf_2246b324dc6504b4),
			Angle:          __obf_f16b2da5958b1356,
			PrimaryColor:   __obf_20c935ac6e8be4f2,
			SecondaryColor: __obf_37bb1e22996a5498,
			Text:           __obf_a63fac8b033b6a09,
		}
	}

	return __obf_e0fb28c1ab249b67
}

func (__obf_b3aecd711f830444 *Captcha) __obf_f30dfd6a7024c455(__obf_0eedbbe6c84cee22 DotType, __obf_5cc8cae6d500bcb3 types.Size, __obf_e0fb28c1ab249b67 []*types.Dot) (image.Image, error) {
	__obf_4c0832c4a893def1 := &DrawImageParams{
		Width:       __obf_5cc8cae6d500bcb3.Width,
		Height:      __obf_5cc8cae6d500bcb3.Height,
		Alpha:       __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.Alpha,
		FontHinting: __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.FontHinting,
		Dots:        __obf_e0fb28c1ab249b67,
		ShowShadow:  __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.ShowShadow,
		ShadowPoint: __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.ShadowPoint,
	}
	var __obf_6090e69b7a7b925f error
	__obf_4c0832c4a893def1.
		Background, __obf_6090e69b7a7b925f = assets.PickBgImage()
	if __obf_6090e69b7a7b925f != nil {
		return nil, __obf_6090e69b7a7b925f
	}

	return __obf_b3aecd711f830444.DrawWithNRGBA(__obf_0eedbbe6c84cee22, __obf_4c0832c4a893def1)
}

func (__obf_fb2f9bdfdc07de19 *Captcha) __obf_11ca91c9d4fc9998(__obf_0eedbbe6c84cee22 DotType, __obf_5cc8cae6d500bcb3 types.Size, __obf_e0fb28c1ab249b67 []*types.Dot) (image.Image, error) {
	var (
		__obf_e8e3e1dc446132b7 =
		// err      error
		make([]*types.Dot, 0, len(__obf_e0fb28c1ab249b67))
	)
	__obf_06628ab133c4f92c := __obf_5cc8cae6d500bcb3.Width / len(__obf_e0fb28c1ab249b67)

	for __obf_339c01a0f79fa1a4, __obf_a63fac8b033b6a09 := range __obf_e0fb28c1ab249b67 {
		__obf_fd906028639fa50b := 1
		if __obf_0eedbbe6c84cee22 == Text {
			__obf_fd906028639fa50b = len(__obf_a63fac8b033b6a09.Text)
		}
		__obf_a63fac8b033b6a09.
			X = int(max(float64(__obf_06628ab133c4f92c*__obf_339c01a0f79fa1a4+__obf_06628ab133c4f92c/__obf_a63fac8b033b6a09.Size), 8))
		__obf_47cff75448b44691 := max(1, __obf_5cc8cae6d500bcb3.Height/16*__obf_fd906028639fa50b)
		__obf_a63fac8b033b6a09.
			Y = __obf_5cc8cae6d500bcb3.Height/2 + __obf_a63fac8b033b6a09.FontSize/2 - rand.IntN(__obf_47cff75448b44691)
		__obf_e8e3e1dc446132b7 = append(__obf_e8e3e1dc446132b7, __obf_a63fac8b033b6a09)
	}
	__obf_f68ed5589ee6f148 := &DrawImageParams{
		Width:             __obf_5cc8cae6d500bcb3.Width,
		Height:            __obf_5cc8cae6d500bcb3.Height,
		Dots:              __obf_e8e3e1dc446132b7,
		BackgroundDistort: __obf_fb2f9bdfdc07de19.__obf_45d12119f89f3fb7(__obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.Secondary.BgDistort),
		BgCircles:         __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.Secondary.BgCircles,
		BgSlimLines:       __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_de0296ed64a1ee9d []color.Color
	for _, __obf_731838253fee91bd := range __obf_c16ce6571c50b663 {
		__obf_5dc2bba1e365f992, _ := util.ParseHexColor(__obf_731838253fee91bd)
		__obf_de0296ed64a1ee9d = append(__obf_de0296ed64a1ee9d, __obf_5dc2bba1e365f992)
	}

	var __obf_b6802dbfd4e5d0aa []color.Color
	for _, __obf_5dc2bba1e365f992 := range __obf_c16ce6571c50b663 {
		__obf_14a37f210c305f6c, _ := util.ParseHexColor(__obf_5dc2bba1e365f992)
		__obf_b6802dbfd4e5d0aa = append(__obf_b6802dbfd4e5d0aa, __obf_14a37f210c305f6c)
	}

	if __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.UseRGBA || __obf_fb2f9bdfdc07de19.__obf_bd5a0fd2a490b814.Secondary.NonDeformAble {
		return __obf_fb2f9bdfdc07de19.DrawWithNRGBA2(__obf_0eedbbe6c84cee22, __obf_f68ed5589ee6f148, __obf_de0296ed64a1ee9d, __obf_b6802dbfd4e5d0aa)
	}
	return __obf_fb2f9bdfdc07de19.DrawWithPalette(__obf_0eedbbe6c84cee22, __obf_f68ed5589ee6f148, __obf_de0296ed64a1ee9d, __obf_b6802dbfd4e5d0aa)
}

func (__obf_b3aecd711f830444 *Captcha) __obf_45d12119f89f3fb7(__obf_26b86685c5bb8762 int) int {
	switch __obf_26b86685c5bb8762 {
	case 1:
		return util.RandInt(240, 320)
	case 2:
		return util.RandInt(180, 240)
	case 3:
		return util.RandInt(120, 180)
	case 4:
		return util.RandInt(100, 160)
	case 5:
		return util.RandInt(80, 140)
	}
	return 0
}

func (__obf_b3aecd711f830444 *Captcha) __obf_f16b2da5958b1356() int {
	__obf_ac1d74f6bdfe9083 := __obf_b3aecd711f830444.__obf_bd5a0fd2a490b814.Primary.AnglePosRange
	__obf_e576f99dcb5bed5a := rand.IntN(len(__obf_ac1d74f6bdfe9083))
	if __obf_e576f99dcb5bed5a < 0 {
		return 0
	}
	__obf_feef62c0167fcfaf := __obf_ac1d74f6bdfe9083[__obf_e576f99dcb5bed5a]
	__obf_43f37316ce10b491 := util.RandInt(__obf_feef62c0167fcfaf.Min, __obf_feef62c0167fcfaf.Max)

	return __obf_43f37316ce10b491
}
