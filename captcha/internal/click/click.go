package __obf_96c99f1db90b23fa

import (
	"image"
	"image/color"
	"math"
	"math/rand/v2"

	assets "github.com/ArtisanHiram/go-pkg/captcha/internal/assets"
	helper "github.com/ArtisanHiram/go-pkg/captcha/internal/helper"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/fixed"
)

type Captcha struct {
	__obf_36f645c70786640b *Options
}

func (__obf_8296aa0f7fc46a6d *Captcha) SetOptions(__obf_36f645c70786640b *Options) {
	__obf_8296aa0f7fc46a6d.__obf_36f645c70786640b = __obf_36f645c70786640b
}

func NewCaptcha(__obf_36f645c70786640b *Options) *Captcha {
	if __obf_36f645c70786640b == nil {
		return nil
	}
	return &Captcha{__obf_36f645c70786640b}
}

func __obf_85294fea53af3b72(__obf_36f645c70786640b *Options) DotType {
	__obf_ee1599da616aa57b := DotType(rand.IntN(2))
	if __obf_ee1599da616aa57b == Shape {
		__obf_36f645c70786640b.UseRGBA = false
		__obf_36f645c70786640b.Secondary.BgDistort = types.DistortLevel1
		__obf_36f645c70786640b.Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_36f645c70786640b.Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_36f645c70786640b.UseRGBA = true
		__obf_36f645c70786640b.Secondary.BgDistort = types.DistortLevel4
		__obf_36f645c70786640b.Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_36f645c70786640b.Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_ee1599da616aa57b
}

func (__obf_8296aa0f7fc46a6d *Captcha) DrawWithNRGBA(__obf_ceae0d6afea37bc5 DotType, __obf_4dea6ff3b5f0bff0 *DrawImageParams) (image.Image, error) {
	__obf_9ecfe5c0e9e57390 := __obf_4dea6ff3b5f0bff0.Dots
	__obf_f386d151dd57ff63 := helper.CreateNRGBACanvas(__obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height, true)

	for _, __obf_e60f690ddf09228f := range __obf_9ecfe5c0e9e57390 {

		__obf_358076043ee0523d, __obf_4ce27fed1d3f1222, __obf_9b38acf488d2d45c := __obf_8296aa0f7fc46a6d.DrawDotImage(__obf_ceae0d6afea37bc5, __obf_e60f690ddf09228f, __obf_4dea6ff3b5f0bff0)
		if __obf_9b38acf488d2d45c != nil {
			return nil, __obf_9b38acf488d2d45c
		}
		__obf_ae93aae8d908f1c5 := __obf_4ce27fed1d3f1222.MinX
		__obf_88b97d948b5b67e1 := __obf_4ce27fed1d3f1222.MaxX
		__obf_b9d51e6a824dd2ca := __obf_4ce27fed1d3f1222.MinY
		__obf_899e2170cfeb222e := __obf_4ce27fed1d3f1222.MaxY

		__obf_6414faaa8a2a1142 := __obf_88b97d948b5b67e1 - __obf_ae93aae8d908f1c5
		__obf_58307cb4018565e8 := __obf_899e2170cfeb222e - __obf_b9d51e6a824dd2ca

		draw.Draw(__obf_f386d151dd57ff63, image.Rect(__obf_e60f690ddf09228f.X, __obf_e60f690ddf09228f.Y, __obf_e60f690ddf09228f.X+__obf_6414faaa8a2a1142, __obf_e60f690ddf09228f.Y+__obf_58307cb4018565e8), __obf_358076043ee0523d, image.Point{X: __obf_ae93aae8d908f1c5, Y: __obf_b9d51e6a824dd2ca}, draw.Over)

	}

	__obf_385665ffaa74a7a1 := __obf_4dea6ff3b5f0bff0.Background
	__obf_111a10711e0739fb := __obf_f386d151dd57ff63.Bounds()
	__obf_fca4b324917dee14 := helper.CreateNRGBACanvas(__obf_111a10711e0739fb.Dx(), __obf_111a10711e0739fb.Dx(), true)
	__obf_7328d5ecb38719de := util.RangCutImagePos(__obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height, __obf_385665ffaa74a7a1)
	draw.Draw(__obf_fca4b324917dee14, __obf_111a10711e0739fb, __obf_385665ffaa74a7a1, __obf_7328d5ecb38719de, draw.Src)
	draw.Draw(__obf_fca4b324917dee14, __obf_f386d151dd57ff63.Bounds(), __obf_f386d151dd57ff63, image.Point{}, draw.Over)
	__obf_fca4b324917dee14.SubImage(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height))
	return __obf_fca4b324917dee14, nil
}

func (__obf_8296aa0f7fc46a6d *Captcha) DrawWithPalette(__obf_ceae0d6afea37bc5 DotType, __obf_4dea6ff3b5f0bff0 *DrawImageParams, __obf_e1c5607aed5d3a43 []color.Color, __obf_918c6cce09330df4 []color.Color) (image.Image, error) {
	__obf_9ecfe5c0e9e57390 := __obf_4dea6ff3b5f0bff0.Dots

	__obf_31b678e2deca0b0d := make([]color.Color, 0, len(__obf_918c6cce09330df4))
	for _, __obf_41de47e71026bafd := range __obf_918c6cce09330df4 {
		__obf_e6b9a3bd6ca91635, __obf_e4a76360da6828b4, __obf_111a10711e0739fb, _ := __obf_41de47e71026bafd.RGBA()
		__obf_57b10e3f5d414338 := util.FormatAlpha(__obf_4dea6ff3b5f0bff0.ThumbDisturbAlpha)
		__obf_31b678e2deca0b0d = append(__obf_31b678e2deca0b0d, color.RGBA{R: uint8(__obf_e6b9a3bd6ca91635), G: uint8(__obf_e4a76360da6828b4), B: uint8(__obf_111a10711e0739fb), A: __obf_57b10e3f5d414338})
	}

	var __obf_d10aad53eab7dcc3 = make([]color.Color, 0, len(__obf_e1c5607aed5d3a43)+len(__obf_31b678e2deca0b0d)+1)
	__obf_d10aad53eab7dcc3 = append(__obf_d10aad53eab7dcc3, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_d10aad53eab7dcc3 = append(__obf_d10aad53eab7dcc3, __obf_e1c5607aed5d3a43...)
	__obf_d10aad53eab7dcc3 = append(__obf_d10aad53eab7dcc3, __obf_31b678e2deca0b0d...)

	__obf_f386d151dd57ff63 := types.NewPalette(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height), __obf_d10aad53eab7dcc3)
	if __obf_4dea6ff3b5f0bff0.BgCircles > 0 {
		__obf_8296aa0f7fc46a6d.__obf_7e90abfc4fb4d1b0(__obf_f386d151dd57ff63, __obf_4dea6ff3b5f0bff0.BgCircles, 1, __obf_31b678e2deca0b0d)
	}
	if __obf_4dea6ff3b5f0bff0.BgSlimLines > 0 {
		__obf_8296aa0f7fc46a6d.__obf_973afbd585aa393b(__obf_f386d151dd57ff63, __obf_4dea6ff3b5f0bff0.BgSlimLines, __obf_31b678e2deca0b0d)
	}

	for _, __obf_e60f690ddf09228f := range __obf_9ecfe5c0e9e57390 {
		__obf_54836e4347be1c09, _ := util.ParseHexColor(__obf_e60f690ddf09228f.PrimaryColor)
		var __obf_9b38acf488d2d45c error
		if __obf_ceae0d6afea37bc5 == Shape {
			var __obf_358076043ee0523d *types.NRGBA
			__obf_358076043ee0523d, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawShapeImage(__obf_e60f690ddf09228f, __obf_54836e4347be1c09)
			if __obf_9b38acf488d2d45c != nil {
				return nil, __obf_9b38acf488d2d45c
			}

			__obf_358076043ee0523d.Rotate(__obf_e60f690ddf09228f.Angle, false)
			__obf_63bceb912bb8aee7 := __obf_f386d151dd57ff63.Bounds()
			__obf_5a4c51df03e3a09f := __obf_358076043ee0523d.Bounds()
			__obf_838b9a46ee7b10ba := image.Point{X: __obf_63bceb912bb8aee7.Dx() - __obf_5a4c51df03e3a09f.Dx(), Y: __obf_63bceb912bb8aee7.Dy() - __obf_5a4c51df03e3a09f.Dy()}
			draw.Draw(__obf_f386d151dd57ff63, image.Rect(__obf_e60f690ddf09228f.X, __obf_838b9a46ee7b10ba.Y, __obf_838b9a46ee7b10ba.X+__obf_5a4c51df03e3a09f.Dx(), __obf_838b9a46ee7b10ba.Y+__obf_5a4c51df03e3a09f.Dy()), __obf_358076043ee0523d, image.Point{}, draw.Over)
		} else {
			__obf_fb9342970f1fc89e := fixed.P(__obf_e60f690ddf09228f.X, __obf_e60f690ddf09228f.Y)
			__obf_4dea6ff3b5f0bff0 := &types.DrawStringParam{
				Color:       __obf_54836e4347be1c09,
				FontSize:    __obf_e60f690ddf09228f.FontSize,
				Size:        __obf_e60f690ddf09228f.Size,
				FontDPI:     __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.FontDPI,
				FontHinting: __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.FontHinting,
				Text:        __obf_e60f690ddf09228f.Text,
			}
			__obf_4dea6ff3b5f0bff0.Font, __obf_9b38acf488d2d45c = assets.PickFont()
			if __obf_9b38acf488d2d45c != nil {
				return nil, __obf_9b38acf488d2d45c
			}
			__obf_9b38acf488d2d45c = helper.DrawString(__obf_f386d151dd57ff63, __obf_4dea6ff3b5f0bff0, __obf_fb9342970f1fc89e)
		}

		if __obf_9b38acf488d2d45c != nil {
			return __obf_f386d151dd57ff63, __obf_9b38acf488d2d45c
		}
	}

	if __obf_4dea6ff3b5f0bff0.Background != nil {
		__obf_385665ffaa74a7a1 := __obf_4dea6ff3b5f0bff0.Background
		__obf_111a10711e0739fb := __obf_385665ffaa74a7a1.Bounds()
		__obf_fca4b324917dee14 := helper.CreateNRGBACanvas(__obf_111a10711e0739fb.Dx(), __obf_111a10711e0739fb.Dy(), true)
		__obf_7328d5ecb38719de := util.RangCutImagePos(__obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height, __obf_385665ffaa74a7a1)
		draw.Draw(__obf_fca4b324917dee14, __obf_111a10711e0739fb, __obf_385665ffaa74a7a1, __obf_7328d5ecb38719de, draw.Src)
		__obf_f386d151dd57ff63.Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_fca4b324917dee14, __obf_f386d151dd57ff63.Bounds(), __obf_f386d151dd57ff63, image.Point{}, draw.Over)
		__obf_3220e41737a017c7 := __obf_fca4b324917dee14.SubImage(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height)).(*image.NRGBA)
		return __obf_3220e41737a017c7, nil
	}

	if __obf_4dea6ff3b5f0bff0.BackgroundDistort > 0 {
		__obf_f386d151dd57ff63.Distort(float64(util.RandInt(5, 10)), float64(__obf_4dea6ff3b5f0bff0.BackgroundDistort))
	}

	return __obf_f386d151dd57ff63, nil

}

func (__obf_8296aa0f7fc46a6d *Captcha) DrawWithNRGBA2(__obf_ceae0d6afea37bc5 DotType, __obf_4dea6ff3b5f0bff0 *DrawImageParams, __obf_e1c5607aed5d3a43 []color.Color, __obf_918c6cce09330df4 []color.Color) (image.Image, error) {
	__obf_9ecfe5c0e9e57390 := __obf_4dea6ff3b5f0bff0.Dots

	__obf_31b678e2deca0b0d := make([]color.Color, 0, len(__obf_918c6cce09330df4))
	for _, __obf_41de47e71026bafd := range __obf_918c6cce09330df4 {
		__obf_e6b9a3bd6ca91635, __obf_e4a76360da6828b4, __obf_111a10711e0739fb, _ := __obf_41de47e71026bafd.RGBA()
		__obf_57b10e3f5d414338 := util.FormatAlpha(__obf_4dea6ff3b5f0bff0.ThumbDisturbAlpha)
		__obf_31b678e2deca0b0d = append(__obf_31b678e2deca0b0d, color.RGBA{R: uint8(__obf_e6b9a3bd6ca91635), G: uint8(__obf_e4a76360da6828b4), B: uint8(__obf_111a10711e0739fb), A: __obf_57b10e3f5d414338})
	}

	var __obf_d10aad53eab7dcc3 = make([]color.Color, 0, len(__obf_e1c5607aed5d3a43)+len(__obf_31b678e2deca0b0d)+1)
	__obf_d10aad53eab7dcc3 = append(__obf_d10aad53eab7dcc3, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_d10aad53eab7dcc3 = append(__obf_d10aad53eab7dcc3, __obf_e1c5607aed5d3a43...)
	__obf_d10aad53eab7dcc3 = append(__obf_d10aad53eab7dcc3, __obf_31b678e2deca0b0d...)

	__obf_4dbe921ef4af1067 := types.NewNRGBA(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height), true)
	if __obf_4dea6ff3b5f0bff0.Background != nil {
		__obf_385665ffaa74a7a1 := __obf_4dea6ff3b5f0bff0.Background
		__obf_111a10711e0739fb := __obf_385665ffaa74a7a1.Bounds()
		__obf_fca4b324917dee14 := helper.CreateNRGBACanvas(__obf_111a10711e0739fb.Dx(), __obf_111a10711e0739fb.Dy(), true)
		__obf_7328d5ecb38719de := util.RangCutImagePos(__obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height, __obf_385665ffaa74a7a1)
		draw.Draw(__obf_fca4b324917dee14, __obf_111a10711e0739fb, __obf_385665ffaa74a7a1, __obf_7328d5ecb38719de, draw.Src)
		__obf_3220e41737a017c7 := __obf_fca4b324917dee14.SubImage(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height)).(*image.NRGBA)
		draw.Draw(__obf_4dbe921ef4af1067, __obf_3220e41737a017c7.Bounds(), __obf_3220e41737a017c7, image.Point{}, draw.Over)
	}

	__obf_f386d151dd57ff63 := types.NewPalette(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height), __obf_d10aad53eab7dcc3)
	if __obf_4dea6ff3b5f0bff0.BgCircles > 0 {
		__obf_8296aa0f7fc46a6d.__obf_7e90abfc4fb4d1b0(__obf_f386d151dd57ff63, __obf_4dea6ff3b5f0bff0.BgCircles, 1, __obf_31b678e2deca0b0d)
	}
	if __obf_4dea6ff3b5f0bff0.BgSlimLines > 0 {
		__obf_8296aa0f7fc46a6d.__obf_973afbd585aa393b(__obf_f386d151dd57ff63, __obf_4dea6ff3b5f0bff0.BgSlimLines, __obf_31b678e2deca0b0d)
	}
	if __obf_4dea6ff3b5f0bff0.BackgroundDistort > 0 {
		__obf_f386d151dd57ff63.Distort(float64(util.RandInt(5, 10)), float64(__obf_4dea6ff3b5f0bff0.BackgroundDistort))
	}

	__obf_d7e1186773310cb4 := __obf_f386d151dd57ff63.Bounds()
	__obf_6414faaa8a2a1142 := __obf_d7e1186773310cb4.Dx() / len(__obf_9ecfe5c0e9e57390)
	for __obf_176527c945a10eb5, __obf_e60f690ddf09228f := range __obf_9ecfe5c0e9e57390 {
		__obf_54836e4347be1c09, _ := util.ParseHexColor(__obf_e60f690ddf09228f.PrimaryColor)
		var __obf_9b38acf488d2d45c error
		if __obf_ceae0d6afea37bc5 == Shape {
			var __obf_358076043ee0523d *types.NRGBA
			__obf_358076043ee0523d, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawShapeImage(__obf_e60f690ddf09228f, __obf_54836e4347be1c09)
			if __obf_9b38acf488d2d45c != nil {
				return nil, __obf_9b38acf488d2d45c
			}

			__obf_358076043ee0523d.Rotate(__obf_e60f690ddf09228f.Angle, false)

			__obf_63bceb912bb8aee7 := __obf_4dbe921ef4af1067.Bounds()
			__obf_5a4c51df03e3a09f := __obf_358076043ee0523d.Bounds()
			__obf_838b9a46ee7b10ba := image.Point{X: __obf_63bceb912bb8aee7.Dx() - __obf_5a4c51df03e3a09f.Dx(), Y: __obf_63bceb912bb8aee7.Dy() - __obf_5a4c51df03e3a09f.Dy()}
			draw.Draw(__obf_4dbe921ef4af1067, image.Rect(__obf_e60f690ddf09228f.X, __obf_838b9a46ee7b10ba.Y, __obf_838b9a46ee7b10ba.X+__obf_5a4c51df03e3a09f.Dx(), __obf_838b9a46ee7b10ba.Y+__obf_5a4c51df03e3a09f.Dy()), __obf_358076043ee0523d, image.Point{}, draw.Over)
		} else {
			var __obf_9b75a99bef87beb5 *types.NRGBA
			__obf_9b75a99bef87beb5, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawStringImage(__obf_e60f690ddf09228f, __obf_54836e4347be1c09)
			if __obf_9b38acf488d2d45c != nil {
				return nil, __obf_9b38acf488d2d45c
			}
			__obf_9b75a99bef87beb5.Rotate(__obf_e60f690ddf09228f.Angle, false)

			__obf_4ce27fed1d3f1222 := __obf_9b75a99bef87beb5.CalcMarginBlankArea()
			__obf_ae93aae8d908f1c5 := __obf_4ce27fed1d3f1222.MinX
			__obf_88b97d948b5b67e1 := __obf_4ce27fed1d3f1222.MaxX
			__obf_b9d51e6a824dd2ca := __obf_4ce27fed1d3f1222.MinY
			__obf_899e2170cfeb222e := __obf_4ce27fed1d3f1222.MaxY
			__obf_9b75a99bef87beb5.SubImage(image.Rect(__obf_ae93aae8d908f1c5, __obf_b9d51e6a824dd2ca, __obf_88b97d948b5b67e1, __obf_899e2170cfeb222e))
			__obf_a3633e3ad0c6d360 := __obf_9b75a99bef87beb5.Bounds()

			__obf_56e9608474b4264b := int(max(float64(__obf_6414faaa8a2a1142*__obf_176527c945a10eb5+__obf_6414faaa8a2a1142/__obf_a3633e3ad0c6d360.Dx()), 8))
			__obf_3a157c8a7156bc0e := util.RandInt(1, __obf_d7e1186773310cb4.Dy()-__obf_a3633e3ad0c6d360.Dy()-4)

			draw.Draw(__obf_4dbe921ef4af1067, image.Rect(__obf_56e9608474b4264b, __obf_3a157c8a7156bc0e, __obf_56e9608474b4264b+__obf_a3633e3ad0c6d360.Dx(), __obf_3a157c8a7156bc0e+__obf_a3633e3ad0c6d360.Dy()), __obf_9b75a99bef87beb5, image.Point{X: __obf_a3633e3ad0c6d360.Min.X, Y: __obf_a3633e3ad0c6d360.Min.Y}, draw.Over)
		}

		if __obf_9b38acf488d2d45c != nil {
			return __obf_4dbe921ef4af1067, __obf_9b38acf488d2d45c
		}
	}

	__obf_ba1d260d15504ee2 := types.NewNRGBA(image.Rect(0, 0, __obf_4dea6ff3b5f0bff0.Width, __obf_4dea6ff3b5f0bff0.Height), true)
	draw.Draw(__obf_ba1d260d15504ee2, __obf_f386d151dd57ff63.Bounds(), __obf_f386d151dd57ff63, image.Point{}, draw.Over)
	draw.Draw(__obf_4dbe921ef4af1067, __obf_ba1d260d15504ee2.Bounds(), __obf_ba1d260d15504ee2, image.Point{}, draw.Over)
	return __obf_4dbe921ef4af1067, nil
}

func (__obf_8296aa0f7fc46a6d *Captcha) __obf_7e90abfc4fb4d1b0(__obf_fca4b324917dee14 *types.Palette, __obf_74a323dfae8efb21, __obf_5a06391eed1c372a int, __obf_02966ab6f37d99d3 []color.Color) {
	__obf_508d62f1f7e68f43 := __obf_fca4b324917dee14.Bounds().Max.X
	__obf_7d11c7d1ccdd7ca4 := __obf_fca4b324917dee14.Bounds().Max.Y
	for range __obf_74a323dfae8efb21 {
		__obf_d8edae615c9199ed := util.RandColor(__obf_02966ab6f37d99d3)
		//co.A = uint8(0xee)
		__obf_e6b9a3bd6ca91635 := util.RandInt(1, __obf_5a06391eed1c372a)
		__obf_fca4b324917dee14.DrawCircle(util.RandInt(__obf_e6b9a3bd6ca91635, __obf_508d62f1f7e68f43-__obf_e6b9a3bd6ca91635), util.RandInt(__obf_e6b9a3bd6ca91635, __obf_7d11c7d1ccdd7ca4-__obf_e6b9a3bd6ca91635), __obf_e6b9a3bd6ca91635, __obf_d8edae615c9199ed)
	}
}

func (__obf_8296aa0f7fc46a6d *Captcha) __obf_973afbd585aa393b(__obf_fca4b324917dee14 *types.Palette, __obf_0efd317c38ba75ea int, __obf_02966ab6f37d99d3 []color.Color) {
	__obf_4f70e86bf5fa5f1a := __obf_fca4b324917dee14.Bounds().Max.X / 10
	__obf_77b2be7c02adf392 := __obf_4f70e86bf5fa5f1a * 9
	__obf_4466379f34011435 := __obf_fca4b324917dee14.Bounds().Max.Y / 3
	for __obf_176527c945a10eb5 := range __obf_0efd317c38ba75ea {
		__obf_c5facfd49b2a5a94 := image.Point{X: rand.IntN(__obf_4f70e86bf5fa5f1a), Y: rand.IntN(__obf_4466379f34011435)}
		__obf_46f54e18dae95128 := image.Point{X: rand.IntN(__obf_4f70e86bf5fa5f1a) + __obf_77b2be7c02adf392, Y: rand.IntN(__obf_4466379f34011435)}

		if __obf_176527c945a10eb5%2 == 0 {
			__obf_c5facfd49b2a5a94.Y = rand.IntN(__obf_4466379f34011435) + __obf_4466379f34011435*2
			__obf_46f54e18dae95128.Y = rand.IntN(__obf_4466379f34011435)
		} else {
			__obf_c5facfd49b2a5a94.Y = rand.IntN(__obf_4466379f34011435) + __obf_4466379f34011435*(__obf_176527c945a10eb5%2)
			__obf_46f54e18dae95128.Y = rand.IntN(__obf_4466379f34011435) + __obf_4466379f34011435*2
		}

		__obf_d8edae615c9199ed := util.RandColor(__obf_02966ab6f37d99d3)
		//co.A = uint8(0xee)
		__obf_fca4b324917dee14.DrawBeeline(__obf_c5facfd49b2a5a94, __obf_46f54e18dae95128, __obf_d8edae615c9199ed)
	}
}

func (__obf_8296aa0f7fc46a6d *Captcha) DrawDotImage(__obf_ceae0d6afea37bc5 DotType, __obf_68c032f360706780 *types.Dot, __obf_4dea6ff3b5f0bff0 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_54836e4347be1c09, _ := util.ParseHexColor(__obf_68c032f360706780.PrimaryColor)
	__obf_54836e4347be1c09.A = util.FormatAlpha(__obf_4dea6ff3b5f0bff0.Alpha)

	var __obf_9b75a99bef87beb5 image.Image
	var __obf_9b38acf488d2d45c error
	if __obf_ceae0d6afea37bc5 == Shape {
		__obf_9b75a99bef87beb5, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawShapeImage(__obf_68c032f360706780, __obf_54836e4347be1c09)
	} else {
		__obf_9b75a99bef87beb5, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawStringImage(__obf_68c032f360706780, __obf_54836e4347be1c09)
	}
	if __obf_9b38acf488d2d45c != nil {
		return nil, nil, __obf_9b38acf488d2d45c
	}

	__obf_66623e18e217fc23, _ := util.ParseHexColor(__obf_a72c755c3ed973fe)
	__obf_f386d151dd57ff63 := helper.CreateNRGBACanvas(__obf_68c032f360706780.Size+10, __obf_68c032f360706780.Size+10, true)
	if __obf_4dea6ff3b5f0bff0.ShowShadow {
		var __obf_14cc087fd3f7da40 *types.NRGBA
		if __obf_ceae0d6afea37bc5 == Shape {
			__obf_14cc087fd3f7da40, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawShapeImage(__obf_68c032f360706780, __obf_66623e18e217fc23)
		} else {
			__obf_14cc087fd3f7da40, __obf_9b38acf488d2d45c = __obf_8296aa0f7fc46a6d.DrawStringImage(__obf_68c032f360706780, __obf_66623e18e217fc23)
		}
		if __obf_9b38acf488d2d45c != nil {
			return nil, nil, __obf_9b38acf488d2d45c
		}

		__obf_660ae33e5203fce0 := __obf_4dea6ff3b5f0bff0.ShadowPoint.X
		__obf_90bcdaf1637e0419 := __obf_4dea6ff3b5f0bff0.ShadowPoint.Y
		draw.Draw(__obf_f386d151dd57ff63, __obf_14cc087fd3f7da40.Bounds(), __obf_14cc087fd3f7da40, image.Point{X: __obf_660ae33e5203fce0, Y: __obf_90bcdaf1637e0419}, draw.Over)
	}
	draw.Draw(__obf_f386d151dd57ff63, __obf_9b75a99bef87beb5.Bounds(), __obf_9b75a99bef87beb5, image.Point{}, draw.Over)
	__obf_f386d151dd57ff63.Rotate(__obf_68c032f360706780.Angle, false)

	__obf_cf567c456372b854 := __obf_f386d151dd57ff63.CalcMarginBlankArea()

	return __obf_f386d151dd57ff63, __obf_cf567c456372b854, nil
}

// DrawStringImage draws a text image
// params:
//   - dot: Draw dot
//   - textColor: Text color
//
// returns:
//   - types.NRGBA: Drawn text image
//   - error: Error information
func (__obf_8296aa0f7fc46a6d *Captcha) DrawStringImage(__obf_68c032f360706780 *types.Dot, __obf_fd96c3ae1d882fb2 color.Color) (*types.NRGBA, error) {
	__obf_f386d151dd57ff63 := helper.CreateNRGBACanvas(__obf_68c032f360706780.Size+10, __obf_68c032f360706780.Size+10, true)

	__obf_fb9342970f1fc89e := fixed.P(12, __obf_68c032f360706780.Size-5)
	if util.IsChineseChar(__obf_68c032f360706780.Text) {
		__obf_fb9342970f1fc89e = fixed.P(10, __obf_68c032f360706780.Size)
	}
	__obf_4dea6ff3b5f0bff0 := &types.DrawStringParam{
		Color:       __obf_fd96c3ae1d882fb2,
		FontSize:    __obf_68c032f360706780.FontSize,
		Size:        __obf_68c032f360706780.Size,
		FontDPI:     __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.FontDPI,
		FontHinting: __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.FontHinting,
		Text:        __obf_68c032f360706780.Text,
	}
	var __obf_9b38acf488d2d45c error
	__obf_4dea6ff3b5f0bff0.Font, __obf_9b38acf488d2d45c = assets.PickFont()
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}
	__obf_9b38acf488d2d45c = helper.DrawString(__obf_f386d151dd57ff63, __obf_4dea6ff3b5f0bff0, __obf_fb9342970f1fc89e)
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	return __obf_f386d151dd57ff63, nil
}

func (__obf_8296aa0f7fc46a6d *Captcha) DrawShapeImage(__obf_68c032f360706780 *types.Dot, __obf_54836e4347be1c09 color.Color) (*types.NRGBA, error) {
	__obf_95a1e79247d1ebe3, __obf_ce2ec0de807a1501, __obf_b02b86e42f0216b8, __obf_aadd898ac70ca91a := __obf_54836e4347be1c09.RGBA()

	var __obf_8c75ef3fa7871a78 = []color.RGBA{
		{R: uint8(__obf_95a1e79247d1ebe3), G: uint8(__obf_ce2ec0de807a1501), B: uint8(__obf_b02b86e42f0216b8), A: uint8(__obf_aadd898ac70ca91a)},
	}

	__obf_ba1d260d15504ee2 := helper.CreateNRGBACanvas(__obf_68c032f360706780.Size+10, __obf_68c032f360706780.Size+10, true)
	var __obf_a3633e3ad0c6d360 image.Rectangle
	var __obf_385665ffaa74a7a1 image.Image
	if __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.UseRGBA {
		__obf_f386d151dd57ff63 := helper.CreateNRGBACanvas(__obf_68c032f360706780.Size+10, __obf_68c032f360706780.Size+10, true)
		draw.BiLinear.Scale(__obf_f386d151dd57ff63, __obf_f386d151dd57ff63.Bounds(), __obf_68c032f360706780.Shape, __obf_68c032f360706780.Shape.Bounds(), draw.Over, nil)
		__obf_a3633e3ad0c6d360 = __obf_f386d151dd57ff63.Bounds()
		__obf_385665ffaa74a7a1 = __obf_f386d151dd57ff63
	} else {
		__obf_f386d151dd57ff63 := helper.CreatePaletteCanvas(__obf_68c032f360706780.Size+10, __obf_68c032f360706780.Size+10, __obf_8c75ef3fa7871a78)
		draw.BiLinear.Scale(__obf_f386d151dd57ff63, __obf_f386d151dd57ff63.Bounds(), __obf_68c032f360706780.Shape, __obf_68c032f360706780.Shape.Bounds(), draw.Over, nil)
		__obf_a3633e3ad0c6d360 = __obf_f386d151dd57ff63.Bounds()
		__obf_385665ffaa74a7a1 = __obf_f386d151dd57ff63
	}

	draw.Draw(__obf_ba1d260d15504ee2, __obf_a3633e3ad0c6d360, __obf_385665ffaa74a7a1, image.Point{}, draw.Over)

	return __obf_ba1d260d15504ee2, nil
}

func (__obf_3f5b5aab6d53fe90 *Captcha) Generate() (*CaptchaData, error) {
	__obf_ceae0d6afea37bc5 := __obf_85294fea53af3b72(__obf_3f5b5aab6d53fe90.__obf_36f645c70786640b)
	if __obf_ceae0d6afea37bc5 == Shape {
		return __obf_3f5b5aab6d53fe90.__obf_d4f4a170a857426d()
	}

	return __obf_3f5b5aab6d53fe90.__obf_6e0658a4f10117d6()
}

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_d4f4a170a857426d() (*CaptchaData, error) {

	var (
		__obf_7566428e4d7ddac2, __obf_004e91a5d6f9d1bf, __obf_d428b607232e523f []*types.Dot
		__obf_bc312d3608271e5b                                                 []image.Image
		__obf_d3a2a708aa6904bc                                                 []image.Image
		__obf_d1b1b184cf543db3, __obf_dedf9178334e9a06                         image.Image
		__obf_9b38acf488d2d45c                                                 error
	)
	__obf_d3a2a708aa6904bc, __obf_9b38acf488d2d45c = assets.GetShapes()
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	__obf_13e4ee348b84dfe5 := util.RandInt(__obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.LenRange.Min, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.LenRange.Max)
	__obf_d3a2a708aa6904bc = util.PickN(__obf_d3a2a708aa6904bc, __obf_13e4ee348b84dfe5)

	__obf_7566428e4d7ddac2 = __obf_3f5b5aab6d53fe90.__obf_06681d046ff271c8(__obf_d3a2a708aa6904bc, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.Size, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.SizeRange, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.DotPadding)
	__obf_d09772c60ee5fc61 := rand.Perm(len(__obf_7566428e4d7ddac2))
	__obf_1617367605964513 := util.RandInt(__obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.VerifyLenRange.Min, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.VerifyLenRange.Max)
	__obf_d428b607232e523f = make([]*types.Dot, __obf_1617367605964513)
	for __obf_176527c945a10eb5, __obf_327dc048f031fff8 := range __obf_d09772c60ee5fc61 {
		if __obf_176527c945a10eb5 >= __obf_1617367605964513 {
			break
		}

		__obf_68c032f360706780 := __obf_7566428e4d7ddac2[__obf_327dc048f031fff8]
		__obf_68c032f360706780.Index = __obf_176527c945a10eb5
		__obf_d428b607232e523f[__obf_176527c945a10eb5] = __obf_68c032f360706780
		__obf_bc312d3608271e5b = append(__obf_bc312d3608271e5b, __obf_d428b607232e523f[__obf_176527c945a10eb5].Shape)
	}
	__obf_004e91a5d6f9d1bf = __obf_3f5b5aab6d53fe90.__obf_06681d046ff271c8(__obf_bc312d3608271e5b, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.Size, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.SizeRange, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.DotPadding)

	__obf_d1b1b184cf543db3, __obf_9b38acf488d2d45c = __obf_3f5b5aab6d53fe90.__obf_a12aae315ad7245c(Shape, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.Size, __obf_7566428e4d7ddac2)
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	__obf_dedf9178334e9a06, __obf_9b38acf488d2d45c = __obf_3f5b5aab6d53fe90.__obf_7b8c0cd619ac3ac5(Shape, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.Size, __obf_004e91a5d6f9d1bf)
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	return &CaptchaData{
		__obf_9ecfe5c0e9e57390: __obf_d428b607232e523f,
		__obf_d1b1b184cf543db3: types.NewJPEGImage(__obf_d1b1b184cf543db3),
		__obf_dedf9178334e9a06: types.NewPNGImage(__obf_dedf9178334e9a06),
	}, nil
}

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_6e0658a4f10117d6() (*CaptchaData, error) {

	__obf_13e4ee348b84dfe5 := util.RandInt(__obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.LenRange.Min, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.LenRange.Max)
	__obf_03b9a9941787d422 := util.PickN(assets.GetChineseChars(), __obf_13e4ee348b84dfe5)

	var __obf_9ecfe5c0e9e57390, __obf_004e91a5d6f9d1bf, __obf_d428b607232e523f []*types.Dot
	var __obf_47feee294ac5f980 []string
	var __obf_d1b1b184cf543db3, __obf_dedf9178334e9a06 image.Image

	__obf_9ecfe5c0e9e57390 = __obf_3f5b5aab6d53fe90.__obf_d522ae8a287dcc38(__obf_03b9a9941787d422, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.Size, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.SizeRange, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.DotPadding)
	__obf_d09772c60ee5fc61 := rand.Perm(len(__obf_9ecfe5c0e9e57390))
	__obf_1617367605964513 := util.RandInt(__obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.VerifyLenRange.Min, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.VerifyLenRange.Max)
	__obf_d428b607232e523f = make([]*types.Dot, __obf_1617367605964513)
	for __obf_176527c945a10eb5, __obf_327dc048f031fff8 := range __obf_d09772c60ee5fc61 {
		if __obf_176527c945a10eb5 >= __obf_1617367605964513 {
			break
		}

		__obf_68c032f360706780 := __obf_9ecfe5c0e9e57390[__obf_327dc048f031fff8]
		__obf_68c032f360706780.Index = __obf_176527c945a10eb5
		__obf_d428b607232e523f[__obf_176527c945a10eb5] = __obf_68c032f360706780
		__obf_47feee294ac5f980 = append(__obf_47feee294ac5f980, __obf_d428b607232e523f[__obf_176527c945a10eb5].Text)
	}

	__obf_004e91a5d6f9d1bf = __obf_3f5b5aab6d53fe90.__obf_d522ae8a287dcc38(__obf_47feee294ac5f980, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.Size, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.SizeRange, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.DotPadding)
	var __obf_9b38acf488d2d45c error
	__obf_d1b1b184cf543db3, __obf_9b38acf488d2d45c = __obf_3f5b5aab6d53fe90.__obf_a12aae315ad7245c(Text, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.Size, __obf_9ecfe5c0e9e57390)
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	__obf_dedf9178334e9a06, __obf_9b38acf488d2d45c = __obf_3f5b5aab6d53fe90.__obf_7b8c0cd619ac3ac5(Text, __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Secondary.Size, __obf_004e91a5d6f9d1bf)
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	return &CaptchaData{
		__obf_9ecfe5c0e9e57390: __obf_d428b607232e523f,
		__obf_d1b1b184cf543db3: types.NewJPEGImage(__obf_d1b1b184cf543db3),
		__obf_dedf9178334e9a06: types.NewPNGImage(__obf_dedf9178334e9a06),
	}, nil
}

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_06681d046ff271c8(__obf_d3a2a708aa6904bc []image.Image, __obf_f701a915f13743a7 types.Size, __obf_5ac189d8fd9016c7 types.Range, __obf_edbad569857d009e int) []*types.Dot {

	__obf_13e4ee348b84dfe5 := len(__obf_d3a2a708aa6904bc)
	var __obf_9ecfe5c0e9e57390 = make([]*types.Dot, __obf_13e4ee348b84dfe5)
	__obf_6414faaa8a2a1142 := __obf_f701a915f13743a7.Width
	__obf_58307cb4018565e8 := __obf_f701a915f13743a7.Height
	if __obf_edbad569857d009e > 0 {
		__obf_6414faaa8a2a1142 -= __obf_edbad569857d009e
		__obf_58307cb4018565e8 -= __obf_edbad569857d009e
	}

	for __obf_176527c945a10eb5, __obf_e60f690ddf09228f := range __obf_d3a2a708aa6904bc {
		__obf_4c1f402977d1c0f2 := __obf_3f5b5aab6d53fe90.__obf_4c1f402977d1c0f2()

		__obf_f10ea87a9194ebc7 := util.PickN(__obf_e223a4630a6875c9, 1)[0]
		__obf_094d08352b51c952 := util.PickN(__obf_2ed7cd7bfa98e4d2, 1)[0]

		__obf_a71c0344a1c76b87 := util.RandInt(__obf_5ac189d8fd9016c7.Min, __obf_5ac189d8fd9016c7.Max)

		__obf_3a157c8a7156bc0e := 10
		__obf_20dd4e9d8ef73062 := __obf_6414faaa8a2a1142 / __obf_13e4ee348b84dfe5
		__obf_dadaaa434b80971f := math.Abs(float64(__obf_20dd4e9d8ef73062) - float64(__obf_a71c0344a1c76b87))
		__obf_66a70090799058d9 := (__obf_176527c945a10eb5 * __obf_20dd4e9d8ef73062) + util.RandInt(0, int(max(__obf_dadaaa434b80971f, 1)))
		__obf_2ec9f36eacc8816c := util.RandInt(__obf_3a157c8a7156bc0e, __obf_58307cb4018565e8+__obf_a71c0344a1c76b87)

		__obf_ad344a17adb4c284 := int(min(max(float64(__obf_66a70090799058d9), float64(__obf_3a157c8a7156bc0e)), float64(__obf_6414faaa8a2a1142-__obf_3a157c8a7156bc0e-(__obf_edbad569857d009e*2))))
		__obf_4466379f34011435 := int(min(max(float64(__obf_2ec9f36eacc8816c), float64(__obf_a71c0344a1c76b87+__obf_3a157c8a7156bc0e)), float64(__obf_58307cb4018565e8+(__obf_a71c0344a1c76b87/2)-(__obf_edbad569857d009e*2))))

		__obf_9ecfe5c0e9e57390[__obf_176527c945a10eb5] = &types.Dot{
			Index:          __obf_176527c945a10eb5,
			X:              __obf_ad344a17adb4c284,
			Y:              __obf_4466379f34011435 - __obf_a71c0344a1c76b87,
			FontSize:       __obf_a71c0344a1c76b87,
			Size:           __obf_a71c0344a1c76b87,
			Angle:          __obf_4c1f402977d1c0f2,
			PrimaryColor:   __obf_f10ea87a9194ebc7,
			SecondaryColor: __obf_094d08352b51c952,
			Shape:          __obf_e60f690ddf09228f,
		}
	}

	return __obf_9ecfe5c0e9e57390
}

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_d522ae8a287dcc38(__obf_03b9a9941787d422 []string, __obf_f701a915f13743a7 types.Size, __obf_fa99aa391bf1e3ff types.Range, __obf_edbad569857d009e int) []*types.Dot {
	__obf_13e4ee348b84dfe5 := len(__obf_03b9a9941787d422)
	var __obf_9ecfe5c0e9e57390 = make([]*types.Dot, __obf_13e4ee348b84dfe5)
	__obf_6414faaa8a2a1142 := __obf_f701a915f13743a7.Width
	__obf_58307cb4018565e8 := __obf_f701a915f13743a7.Height
	if __obf_edbad569857d009e > 0 {
		__obf_6414faaa8a2a1142 -= __obf_edbad569857d009e
		__obf_58307cb4018565e8 -= __obf_edbad569857d009e
	}

	for __obf_176527c945a10eb5, __obf_e60f690ddf09228f := range __obf_03b9a9941787d422 {
		__obf_4c1f402977d1c0f2 := __obf_3f5b5aab6d53fe90.__obf_4c1f402977d1c0f2()

		__obf_f10ea87a9194ebc7 := util.PickN(__obf_e223a4630a6875c9, 1)[0]
		__obf_094d08352b51c952 := util.PickN(__obf_2ed7cd7bfa98e4d2, 1)[0]

		__obf_a71c0344a1c76b87 := util.RandInt(__obf_fa99aa391bf1e3ff.Min, __obf_fa99aa391bf1e3ff.Max)
		__obf_bef35a6075aabff3 := __obf_a71c0344a1c76b87
		__obf_161fa89d6c484074 := __obf_a71c0344a1c76b87

		if util.LenChineseChar(__obf_e60f690ddf09228f) > 1 {
			__obf_161fa89d6c484074 = __obf_a71c0344a1c76b87 * util.LenChineseChar(__obf_e60f690ddf09228f)

			if __obf_4c1f402977d1c0f2 > 0 {
				__obf_2fc25ac41010c7b4 := __obf_161fa89d6c484074 - __obf_a71c0344a1c76b87
				__obf_c9706f4aadd02e85 := __obf_4c1f402977d1c0f2 % 90
				__obf_2df10cc44b83b280 := float64(__obf_2fc25ac41010c7b4) / 90
				__obf_e6b9a3bd6ca91635 := max(float64(__obf_c9706f4aadd02e85)*__obf_2df10cc44b83b280, 1)
				__obf_bef35a6075aabff3 = __obf_bef35a6075aabff3 + int(__obf_e6b9a3bd6ca91635)
				__obf_161fa89d6c484074 = __obf_161fa89d6c484074 + int(__obf_e6b9a3bd6ca91635)
			}
		}

		__obf_3a157c8a7156bc0e := 10
		__obf_20dd4e9d8ef73062 := __obf_6414faaa8a2a1142 / __obf_13e4ee348b84dfe5
		__obf_dadaaa434b80971f := math.Abs(float64(__obf_20dd4e9d8ef73062) - float64(__obf_161fa89d6c484074))
		__obf_66a70090799058d9 := (__obf_176527c945a10eb5 * __obf_20dd4e9d8ef73062) + util.RandInt(0, int(max(__obf_dadaaa434b80971f, 1)))
		__obf_2ec9f36eacc8816c := util.RandInt(__obf_3a157c8a7156bc0e, __obf_58307cb4018565e8+__obf_bef35a6075aabff3)

		__obf_ad344a17adb4c284 := int(min(max(float64(__obf_66a70090799058d9), float64(__obf_3a157c8a7156bc0e)), float64(__obf_6414faaa8a2a1142-__obf_3a157c8a7156bc0e-(__obf_edbad569857d009e*2))))
		__obf_4466379f34011435 := int(min(max(float64(__obf_2ec9f36eacc8816c), float64(__obf_bef35a6075aabff3+__obf_3a157c8a7156bc0e)), float64(__obf_58307cb4018565e8+(__obf_bef35a6075aabff3/2)-(__obf_edbad569857d009e*2))))

		__obf_9ecfe5c0e9e57390[__obf_176527c945a10eb5] = &types.Dot{
			Index:          __obf_176527c945a10eb5,
			X:              __obf_ad344a17adb4c284,
			Y:              __obf_4466379f34011435 - __obf_bef35a6075aabff3,
			FontSize:       __obf_a71c0344a1c76b87,
			Size:           max(__obf_161fa89d6c484074, __obf_bef35a6075aabff3),
			Angle:          __obf_4c1f402977d1c0f2,
			PrimaryColor:   __obf_f10ea87a9194ebc7,
			SecondaryColor: __obf_094d08352b51c952,
			Text:           __obf_e60f690ddf09228f,
		}
	}

	return __obf_9ecfe5c0e9e57390
}

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_a12aae315ad7245c(__obf_ceae0d6afea37bc5 DotType, __obf_fa99aa391bf1e3ff types.Size, __obf_9ecfe5c0e9e57390 []*types.Dot) (image.Image, error) {

	__obf_91ee00cd0ad38eb0 := &DrawImageParams{
		Width:       __obf_fa99aa391bf1e3ff.Width,
		Height:      __obf_fa99aa391bf1e3ff.Height,
		Alpha:       __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.Alpha,
		FontHinting: __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.FontHinting,
		Dots:        __obf_9ecfe5c0e9e57390,
		ShowShadow:  __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.ShowShadow,
		ShadowPoint: __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.ShadowPoint,
	}
	var __obf_9b38acf488d2d45c error
	__obf_91ee00cd0ad38eb0.Background, __obf_9b38acf488d2d45c = assets.PickBgImage()
	if __obf_9b38acf488d2d45c != nil {
		return nil, __obf_9b38acf488d2d45c
	}

	return __obf_3f5b5aab6d53fe90.DrawWithNRGBA(__obf_ceae0d6afea37bc5, __obf_91ee00cd0ad38eb0)
}

func (__obf_8296aa0f7fc46a6d *Captcha) __obf_7b8c0cd619ac3ac5(__obf_ceae0d6afea37bc5 DotType, __obf_fa99aa391bf1e3ff types.Size, __obf_9ecfe5c0e9e57390 []*types.Dot) (image.Image, error) {
	var (
		// err      error
		__obf_332def4d7dbc39f9 = make([]*types.Dot, 0, len(__obf_9ecfe5c0e9e57390))
	)

	__obf_6414faaa8a2a1142 := __obf_fa99aa391bf1e3ff.Width / len(__obf_9ecfe5c0e9e57390)

	for __obf_176527c945a10eb5, __obf_e60f690ddf09228f := range __obf_9ecfe5c0e9e57390 {
		__obf_13e4ee348b84dfe5 := 1
		if __obf_ceae0d6afea37bc5 == Text {
			__obf_13e4ee348b84dfe5 = len(__obf_e60f690ddf09228f.Text)
		}

		__obf_e60f690ddf09228f.X = int(max(float64(__obf_6414faaa8a2a1142*__obf_176527c945a10eb5+__obf_6414faaa8a2a1142/__obf_e60f690ddf09228f.Size), 8))
		__obf_6236739462c80905 := max(1, __obf_fa99aa391bf1e3ff.Height/16*__obf_13e4ee348b84dfe5)
		__obf_e60f690ddf09228f.Y = __obf_fa99aa391bf1e3ff.Height/2 + __obf_e60f690ddf09228f.FontSize/2 - rand.IntN(__obf_6236739462c80905)

		__obf_332def4d7dbc39f9 = append(__obf_332def4d7dbc39f9, __obf_e60f690ddf09228f)
	}

	__obf_4dea6ff3b5f0bff0 := &DrawImageParams{
		Width:             __obf_fa99aa391bf1e3ff.Width,
		Height:            __obf_fa99aa391bf1e3ff.Height,
		Dots:              __obf_332def4d7dbc39f9,
		BackgroundDistort: __obf_8296aa0f7fc46a6d.__obf_6211b3919eff8ed0(__obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.Secondary.BgDistort),
		BgCircles:         __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.Secondary.BgCircles,
		BgSlimLines:       __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_4a80a22ca5be8157 []color.Color
	for _, __obf_abc27996b841d6fd := range __obf_2ed7cd7bfa98e4d2 {
		__obf_d8edae615c9199ed, _ := util.ParseHexColor(__obf_abc27996b841d6fd)
		__obf_4a80a22ca5be8157 = append(__obf_4a80a22ca5be8157, __obf_d8edae615c9199ed)
	}

	var __obf_918c6cce09330df4 []color.Color
	for _, __obf_d8edae615c9199ed := range __obf_2ed7cd7bfa98e4d2 {
		__obf_3220e41737a017c7, _ := util.ParseHexColor(__obf_d8edae615c9199ed)
		__obf_918c6cce09330df4 = append(__obf_918c6cce09330df4, __obf_3220e41737a017c7)
	}

	if __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.UseRGBA || __obf_8296aa0f7fc46a6d.__obf_36f645c70786640b.Secondary.NonDeformAble {
		return __obf_8296aa0f7fc46a6d.DrawWithNRGBA2(__obf_ceae0d6afea37bc5, __obf_4dea6ff3b5f0bff0, __obf_4a80a22ca5be8157, __obf_918c6cce09330df4)
	}
	return __obf_8296aa0f7fc46a6d.DrawWithPalette(__obf_ceae0d6afea37bc5, __obf_4dea6ff3b5f0bff0, __obf_4a80a22ca5be8157, __obf_918c6cce09330df4)
}

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_6211b3919eff8ed0(__obf_e3b24f7328ccc938 int) int {
	switch __obf_e3b24f7328ccc938 {
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

func (__obf_3f5b5aab6d53fe90 *Captcha) __obf_4c1f402977d1c0f2() int {
	__obf_17860174bbcc2d10 := __obf_3f5b5aab6d53fe90.__obf_36f645c70786640b.Primary.AnglePosRange

	__obf_07cbddf49468e8cf := rand.IntN(len(__obf_17860174bbcc2d10))
	if __obf_07cbddf49468e8cf < 0 {
		return 0
	}

	__obf_fa91df9186451baa := __obf_17860174bbcc2d10[__obf_07cbddf49468e8cf]
	__obf_97dac6fcf0173677 := util.RandInt(__obf_fa91df9186451baa.Min, __obf_fa91df9186451baa.Max)

	return __obf_97dac6fcf0173677
}
