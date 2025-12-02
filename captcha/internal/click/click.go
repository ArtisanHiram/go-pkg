package __obf_767326b1078782af

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
	__obf_8659ee70bdbf57ab *types.ClickOption
}

var _ types.Captcha = (*Captcha)(nil)

func NewCaptcha(__obf_8659ee70bdbf57ab *types.ClickOption) *Captcha {
	if __obf_8659ee70bdbf57ab == nil {
		return nil
	}
	return &Captcha{__obf_8659ee70bdbf57ab}
}

func (__obf_de65322a0d55e5bb *Captcha) Generate() (types.CaptchaData, error) {
	__obf_ac0f4faf58af21b0 := __obf_94f9de2fcea55dbd(__obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab)
	if __obf_ac0f4faf58af21b0 == Shape {
		return __obf_de65322a0d55e5bb.__obf_368e3993fbda0ca0()
	}

	return __obf_de65322a0d55e5bb.__obf_598c6b664136f15a()
}

func __obf_94f9de2fcea55dbd(__obf_8659ee70bdbf57ab *types.ClickOption) DotType {
	__obf_37e1f9e58b25ab60 := DotType(rand.IntN(2))
	if __obf_37e1f9e58b25ab60 == Shape {
		__obf_8659ee70bdbf57ab.
			UseRGBA = false
		__obf_8659ee70bdbf57ab.
			Secondary.BgDistort = types.DistortLevel1
		__obf_8659ee70bdbf57ab.
			Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_8659ee70bdbf57ab.
			Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_8659ee70bdbf57ab.
			UseRGBA = true
		__obf_8659ee70bdbf57ab.
			Secondary.BgDistort = types.DistortLevel4
		__obf_8659ee70bdbf57ab.
			Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_8659ee70bdbf57ab.
			Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_37e1f9e58b25ab60
}

func (__obf_0aaa9d42f2545c73 *Captcha) DrawWithNRGBA(__obf_ac0f4faf58af21b0 DotType, __obf_c31e22f60e6accb4 *DrawImageParams) (image.Image, error) {
	__obf_2028bb55834c38d8 := __obf_c31e22f60e6accb4.Dots
	__obf_60e5f15f77605018 := helper.CreateNRGBACanvas(__obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height, true)

	for _, __obf_c07c64e538582e45 := range __obf_2028bb55834c38d8 {
		__obf_37e551ef60991191, __obf_a85e728f16ab9692, __obf_4f137bf3a2d5ad2c := __obf_0aaa9d42f2545c73.DrawDotImage(__obf_ac0f4faf58af21b0, __obf_c07c64e538582e45, __obf_c31e22f60e6accb4)
		if __obf_4f137bf3a2d5ad2c != nil {
			return nil, __obf_4f137bf3a2d5ad2c
		}
		__obf_360de838b14aa620 := __obf_a85e728f16ab9692.MinX
		__obf_c79c157d36117f84 := __obf_a85e728f16ab9692.MaxX
		__obf_6e38eed823e9df44 := __obf_a85e728f16ab9692.MinY
		__obf_c5f3eeb93255b5bf := __obf_a85e728f16ab9692.MaxY
		__obf_a51318f96bb6cca4 := __obf_c79c157d36117f84 - __obf_360de838b14aa620
		__obf_0d08f75c1fe24ae6 := __obf_c5f3eeb93255b5bf - __obf_6e38eed823e9df44

		// draw.Draw(cvs, image.Rect(v.X, v.Y, v.X+width, v.Y+height), dotImage, image.Point{X: minX, Y: minY}, draw.Over)
		__obf_541b407a9301526e := __obf_c07c64e538582e45.X - __obf_a51318f96bb6cca4/2
		__obf_eaabf3d66fd4ca8e := __obf_c07c64e538582e45.Y - __obf_0d08f75c1fe24ae6/2
		draw.Draw(__obf_60e5f15f77605018, image.Rect(__obf_541b407a9301526e, __obf_eaabf3d66fd4ca8e, __obf_541b407a9301526e+__obf_a51318f96bb6cca4, __obf_eaabf3d66fd4ca8e+__obf_0d08f75c1fe24ae6), __obf_37e551ef60991191, image.Point{X: __obf_360de838b14aa620, Y: __obf_6e38eed823e9df44}, draw.Over)

	}
	__obf_a02b8d0e4a886ebd := __obf_c31e22f60e6accb4.Background
	__obf_73f49e5108181c83 := __obf_60e5f15f77605018.Bounds()
	__obf_8d555da44024984a := helper.CreateNRGBACanvas(__obf_73f49e5108181c83.Dx(), __obf_73f49e5108181c83.Dy(), true)
	__obf_9a5c9c174d53b2fe := util.RangCutImagePos(__obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height, __obf_a02b8d0e4a886ebd)
	draw.Draw(__obf_8d555da44024984a, __obf_73f49e5108181c83, __obf_a02b8d0e4a886ebd, __obf_9a5c9c174d53b2fe, draw.Src)
	draw.Draw(__obf_8d555da44024984a, __obf_60e5f15f77605018.Bounds(), __obf_60e5f15f77605018, image.Point{}, draw.Over)
	__obf_8d555da44024984a.
		SubImage(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height))
	return __obf_8d555da44024984a, nil
}

func (__obf_0aaa9d42f2545c73 *Captcha) DrawWithPalette(__obf_ac0f4faf58af21b0 DotType, __obf_c31e22f60e6accb4 *DrawImageParams, __obf_cbc5d1b0362a8418 []color.Color, __obf_ef2f557a68e49262 []color.Color) (image.Image, error) {
	__obf_2028bb55834c38d8 := __obf_c31e22f60e6accb4.Dots
	__obf_a8c210e69d86872b := make([]color.Color, 0, len(__obf_ef2f557a68e49262))
	for _, __obf_1f11f2398d10fb45 := range __obf_ef2f557a68e49262 {
		__obf_297d7fb14f68ca1c, __obf_d343b3f2b3fb0484, __obf_73f49e5108181c83, _ := __obf_1f11f2398d10fb45.RGBA()
		__obf_d5362208381caa28 := util.FormatAlpha(__obf_c31e22f60e6accb4.ThumbDisturbAlpha)
		__obf_a8c210e69d86872b = append(__obf_a8c210e69d86872b, color.RGBA{R: uint8(__obf_297d7fb14f68ca1c), G: uint8(__obf_d343b3f2b3fb0484), B: uint8(__obf_73f49e5108181c83), A: __obf_d5362208381caa28})
	}

	var __obf_e63a03cac0e30214 = make([]color.Color, 0, len(__obf_cbc5d1b0362a8418)+len(__obf_a8c210e69d86872b)+1)
	__obf_e63a03cac0e30214 = append(__obf_e63a03cac0e30214, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_e63a03cac0e30214 = append(__obf_e63a03cac0e30214, __obf_cbc5d1b0362a8418...)
	__obf_e63a03cac0e30214 = append(__obf_e63a03cac0e30214, __obf_a8c210e69d86872b...)
	__obf_60e5f15f77605018 := types.NewPalette(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height), __obf_e63a03cac0e30214)
	if __obf_c31e22f60e6accb4.BgCircles > 0 {
		__obf_0aaa9d42f2545c73.__obf_473a60a3b956775b(__obf_60e5f15f77605018, __obf_c31e22f60e6accb4.BgCircles, 1, __obf_a8c210e69d86872b)
	}
	if __obf_c31e22f60e6accb4.BgSlimLines > 0 {
		__obf_0aaa9d42f2545c73.__obf_e85e5e8dc615efa3(__obf_60e5f15f77605018, __obf_c31e22f60e6accb4.BgSlimLines, __obf_a8c210e69d86872b)
	}

	for _, __obf_c07c64e538582e45 := range __obf_2028bb55834c38d8 {
		__obf_8c4c2ec57f522dc0, _ := util.ParseHexColor(__obf_c07c64e538582e45.PrimaryColor)
		var __obf_4f137bf3a2d5ad2c error
		if __obf_ac0f4faf58af21b0 == Shape {
			var __obf_37e551ef60991191 *types.NRGBA
			__obf_37e551ef60991191, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawShapeImage(__obf_c07c64e538582e45, __obf_8c4c2ec57f522dc0)
			if __obf_4f137bf3a2d5ad2c != nil {
				return nil, __obf_4f137bf3a2d5ad2c
			}
			__obf_37e551ef60991191.
				Rotate(__obf_c07c64e538582e45.Angle, false)
			__obf_aa987fdae0fc62f8 := __obf_60e5f15f77605018.Bounds()
			__obf_0a9b267f8e3bc829 := __obf_37e551ef60991191.Bounds()
			__obf_c195c429cb6566c9 := image.Point{X: __obf_aa987fdae0fc62f8.Dx() - __obf_0a9b267f8e3bc829.Dx(), Y: __obf_aa987fdae0fc62f8.Dy() - __obf_0a9b267f8e3bc829.Dy()}
			draw.Draw(__obf_60e5f15f77605018, image.Rect(__obf_c07c64e538582e45.X, __obf_c195c429cb6566c9.Y, __obf_c195c429cb6566c9.X+__obf_0a9b267f8e3bc829.Dx(), __obf_c195c429cb6566c9.Y+__obf_0a9b267f8e3bc829.Dy()), __obf_37e551ef60991191, image.Point{}, draw.Over)
		} else {
			__obf_0bcb6aea1c516a3d := fixed.P(__obf_c07c64e538582e45.X, __obf_c07c64e538582e45.Y)
			__obf_c31e22f60e6accb4 := &types.DrawStringParam{
				Color:       __obf_8c4c2ec57f522dc0,
				FontSize:    __obf_c07c64e538582e45.FontSize,
				Size:        __obf_c07c64e538582e45.Size,
				FontDPI:     __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.FontDPI,
				FontHinting: __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.FontHinting,
				Text:        __obf_c07c64e538582e45.Text,
			}
			__obf_c31e22f60e6accb4.
				Font, __obf_4f137bf3a2d5ad2c = assets.PickFont()
			if __obf_4f137bf3a2d5ad2c != nil {
				return nil, __obf_4f137bf3a2d5ad2c
			}
			__obf_4f137bf3a2d5ad2c = helper.DrawString(__obf_60e5f15f77605018, __obf_c31e22f60e6accb4, __obf_0bcb6aea1c516a3d)
		}

		if __obf_4f137bf3a2d5ad2c != nil {
			return __obf_60e5f15f77605018, __obf_4f137bf3a2d5ad2c
		}
	}

	if __obf_c31e22f60e6accb4.Background != nil {
		__obf_a02b8d0e4a886ebd := __obf_c31e22f60e6accb4.Background
		__obf_73f49e5108181c83 := __obf_a02b8d0e4a886ebd.Bounds()
		__obf_8d555da44024984a := helper.CreateNRGBACanvas(__obf_73f49e5108181c83.Dx(), __obf_73f49e5108181c83.Dy(), true)
		__obf_9a5c9c174d53b2fe := util.RangCutImagePos(__obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height, __obf_a02b8d0e4a886ebd)
		draw.Draw(__obf_8d555da44024984a, __obf_73f49e5108181c83, __obf_a02b8d0e4a886ebd, __obf_9a5c9c174d53b2fe, draw.Src)
		__obf_60e5f15f77605018.
			Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_8d555da44024984a, __obf_60e5f15f77605018.Bounds(), __obf_60e5f15f77605018, image.Point{}, draw.Over)
		__obf_8b68dea5ea6d2497 := __obf_8d555da44024984a.SubImage(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height)).(*image.NRGBA)
		return __obf_8b68dea5ea6d2497, nil
	}

	if __obf_c31e22f60e6accb4.BackgroundDistort > 0 {
		__obf_60e5f15f77605018.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_c31e22f60e6accb4.BackgroundDistort))
	}

	return __obf_60e5f15f77605018, nil

}

func (__obf_0aaa9d42f2545c73 *Captcha) DrawWithNRGBA2(__obf_ac0f4faf58af21b0 DotType, __obf_c31e22f60e6accb4 *DrawImageParams, __obf_cbc5d1b0362a8418 []color.Color, __obf_ef2f557a68e49262 []color.Color) (image.Image, error) {
	__obf_2028bb55834c38d8 := __obf_c31e22f60e6accb4.Dots
	__obf_a8c210e69d86872b := make([]color.Color, 0, len(__obf_ef2f557a68e49262))
	for _, __obf_1f11f2398d10fb45 := range __obf_ef2f557a68e49262 {
		__obf_297d7fb14f68ca1c, __obf_d343b3f2b3fb0484, __obf_73f49e5108181c83, _ := __obf_1f11f2398d10fb45.RGBA()
		__obf_d5362208381caa28 := util.FormatAlpha(__obf_c31e22f60e6accb4.ThumbDisturbAlpha)
		__obf_a8c210e69d86872b = append(__obf_a8c210e69d86872b, color.RGBA{R: uint8(__obf_297d7fb14f68ca1c), G: uint8(__obf_d343b3f2b3fb0484), B: uint8(__obf_73f49e5108181c83), A: __obf_d5362208381caa28})
	}

	var __obf_e63a03cac0e30214 = make([]color.Color, 0, len(__obf_cbc5d1b0362a8418)+len(__obf_a8c210e69d86872b)+1)
	__obf_e63a03cac0e30214 = append(__obf_e63a03cac0e30214, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_e63a03cac0e30214 = append(__obf_e63a03cac0e30214, __obf_cbc5d1b0362a8418...)
	__obf_e63a03cac0e30214 = append(__obf_e63a03cac0e30214, __obf_a8c210e69d86872b...)
	__obf_7325ee775f3e5353 := types.NewNRGBA(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height), true)
	if __obf_c31e22f60e6accb4.Background != nil {
		__obf_a02b8d0e4a886ebd := __obf_c31e22f60e6accb4.Background
		__obf_73f49e5108181c83 := __obf_a02b8d0e4a886ebd.Bounds()
		__obf_8d555da44024984a := helper.CreateNRGBACanvas(__obf_73f49e5108181c83.Dx(), __obf_73f49e5108181c83.Dy(), true)
		__obf_9a5c9c174d53b2fe := util.RangCutImagePos(__obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height, __obf_a02b8d0e4a886ebd)
		draw.Draw(__obf_8d555da44024984a, __obf_73f49e5108181c83, __obf_a02b8d0e4a886ebd, __obf_9a5c9c174d53b2fe, draw.Src)
		__obf_8b68dea5ea6d2497 := __obf_8d555da44024984a.SubImage(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height)).(*image.NRGBA)
		draw.Draw(__obf_7325ee775f3e5353, __obf_8b68dea5ea6d2497.Bounds(), __obf_8b68dea5ea6d2497, image.Point{}, draw.Over)
	}
	__obf_60e5f15f77605018 := types.NewPalette(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height), __obf_e63a03cac0e30214)
	if __obf_c31e22f60e6accb4.BgCircles > 0 {
		__obf_0aaa9d42f2545c73.__obf_473a60a3b956775b(__obf_60e5f15f77605018, __obf_c31e22f60e6accb4.BgCircles, 1, __obf_a8c210e69d86872b)
	}
	if __obf_c31e22f60e6accb4.BgSlimLines > 0 {
		__obf_0aaa9d42f2545c73.__obf_e85e5e8dc615efa3(__obf_60e5f15f77605018, __obf_c31e22f60e6accb4.BgSlimLines, __obf_a8c210e69d86872b)
	}
	if __obf_c31e22f60e6accb4.BackgroundDistort > 0 {
		__obf_60e5f15f77605018.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_c31e22f60e6accb4.BackgroundDistort))
	}
	__obf_b76637c2f54617d7 := __obf_60e5f15f77605018.Bounds()
	__obf_a51318f96bb6cca4 := __obf_b76637c2f54617d7.Dx() / len(__obf_2028bb55834c38d8)
	for __obf_20aa71cdd9d88320, __obf_c07c64e538582e45 := range __obf_2028bb55834c38d8 {
		__obf_8c4c2ec57f522dc0, _ := util.ParseHexColor(__obf_c07c64e538582e45.PrimaryColor)
		var __obf_4f137bf3a2d5ad2c error
		if __obf_ac0f4faf58af21b0 == Shape {
			var __obf_37e551ef60991191 *types.NRGBA
			__obf_37e551ef60991191, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawShapeImage(__obf_c07c64e538582e45, __obf_8c4c2ec57f522dc0)
			if __obf_4f137bf3a2d5ad2c != nil {
				return nil, __obf_4f137bf3a2d5ad2c
			}
			__obf_37e551ef60991191.
				Rotate(__obf_c07c64e538582e45.Angle, false)
			__obf_aa987fdae0fc62f8 := __obf_7325ee775f3e5353.Bounds()
			__obf_0a9b267f8e3bc829 := __obf_37e551ef60991191.Bounds()
			__obf_c195c429cb6566c9 := image.Point{X: __obf_aa987fdae0fc62f8.Dx() - __obf_0a9b267f8e3bc829.Dx(), Y: __obf_aa987fdae0fc62f8.Dy() - __obf_0a9b267f8e3bc829.Dy()}
			draw.Draw(__obf_7325ee775f3e5353, image.Rect(__obf_c07c64e538582e45.X, __obf_c195c429cb6566c9.Y, __obf_c195c429cb6566c9.X+__obf_0a9b267f8e3bc829.Dx(), __obf_c195c429cb6566c9.Y+__obf_0a9b267f8e3bc829.Dy()), __obf_37e551ef60991191, image.Point{}, draw.Over)
		} else {
			var __obf_5675cd5d4389f309 *types.NRGBA
			__obf_5675cd5d4389f309, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawStringImage(__obf_c07c64e538582e45, __obf_8c4c2ec57f522dc0)
			if __obf_4f137bf3a2d5ad2c != nil {
				return nil, __obf_4f137bf3a2d5ad2c
			}
			__obf_5675cd5d4389f309.
				Rotate(__obf_c07c64e538582e45.Angle, false)
			__obf_a85e728f16ab9692 := __obf_5675cd5d4389f309.CalcMarginBlankArea()
			__obf_360de838b14aa620 := __obf_a85e728f16ab9692.MinX
			__obf_c79c157d36117f84 := __obf_a85e728f16ab9692.MaxX
			__obf_6e38eed823e9df44 := __obf_a85e728f16ab9692.MinY
			__obf_c5f3eeb93255b5bf := __obf_a85e728f16ab9692.MaxY
			__obf_5675cd5d4389f309.
				SubImage(image.Rect(__obf_360de838b14aa620, __obf_6e38eed823e9df44, __obf_c79c157d36117f84, __obf_c5f3eeb93255b5bf))
			__obf_516347dd0c9f207c := __obf_5675cd5d4389f309.Bounds()
			__obf_edb0b0da19f8060a := int(max(float64(__obf_a51318f96bb6cca4*__obf_20aa71cdd9d88320+__obf_a51318f96bb6cca4/__obf_516347dd0c9f207c.Dx()), 8))
			__obf_e5934f188ae696f6 := util.RandInt(1, __obf_b76637c2f54617d7.Dy()-__obf_516347dd0c9f207c.Dy()-4)

			draw.Draw(__obf_7325ee775f3e5353, image.Rect(__obf_edb0b0da19f8060a, __obf_e5934f188ae696f6, __obf_edb0b0da19f8060a+__obf_516347dd0c9f207c.Dx(), __obf_e5934f188ae696f6+__obf_516347dd0c9f207c.Dy()), __obf_5675cd5d4389f309, image.Point{X: __obf_516347dd0c9f207c.Min.X, Y: __obf_516347dd0c9f207c.Min.Y}, draw.Over)
		}

		if __obf_4f137bf3a2d5ad2c != nil {
			return __obf_7325ee775f3e5353, __obf_4f137bf3a2d5ad2c
		}
	}
	__obf_fb798136dde8d324 := types.NewNRGBA(image.Rect(0, 0, __obf_c31e22f60e6accb4.Width, __obf_c31e22f60e6accb4.Height), true)
	draw.Draw(__obf_fb798136dde8d324, __obf_60e5f15f77605018.Bounds(), __obf_60e5f15f77605018, image.Point{}, draw.Over)
	draw.Draw(__obf_7325ee775f3e5353, __obf_fb798136dde8d324.Bounds(), __obf_fb798136dde8d324, image.Point{}, draw.Over)
	return __obf_7325ee775f3e5353, nil
}

func (__obf_0aaa9d42f2545c73 *Captcha) __obf_473a60a3b956775b(__obf_8d555da44024984a *types.Palette, __obf_21f14df89d702b81, __obf_573b600af0f7d141 int, __obf_d3479f2f25f21887 []color.Color) {
	__obf_ccbac395a430a226 := __obf_8d555da44024984a.Bounds().Max.X
	__obf_91c886b479154c21 := __obf_8d555da44024984a.Bounds().Max.Y
	for range __obf_21f14df89d702b81 {
		__obf_d55625589b36c20c := util.RandColor(__obf_d3479f2f25f21887)
		__obf_297d7fb14f68ca1c := //co.A = uint8(0xee)
			util.RandInt(1, __obf_573b600af0f7d141)
		__obf_8d555da44024984a.
			DrawCircle(util.RandInt(__obf_297d7fb14f68ca1c, __obf_ccbac395a430a226-__obf_297d7fb14f68ca1c), util.RandInt(__obf_297d7fb14f68ca1c, __obf_91c886b479154c21-__obf_297d7fb14f68ca1c), __obf_297d7fb14f68ca1c, __obf_d55625589b36c20c)
	}
}

func (__obf_0aaa9d42f2545c73 *Captcha) __obf_e85e5e8dc615efa3(__obf_8d555da44024984a *types.Palette, __obf_4c0087e6903d2368 int, __obf_d3479f2f25f21887 []color.Color) {
	__obf_6a5714db27cbf648 := __obf_8d555da44024984a.Bounds().Max.X / 10
	__obf_d8f715a894682dfb := __obf_6a5714db27cbf648 * 9
	__obf_7f6203877254d5ef := __obf_8d555da44024984a.Bounds().Max.Y / 3
	for __obf_20aa71cdd9d88320 := range __obf_4c0087e6903d2368 {
		__obf_acaa0892c9b9426a := image.Point{X: rand.IntN(__obf_6a5714db27cbf648), Y: rand.IntN(__obf_7f6203877254d5ef)}
		__obf_5776bee544324174 := image.Point{X: rand.IntN(__obf_6a5714db27cbf648) + __obf_d8f715a894682dfb, Y: rand.IntN(__obf_7f6203877254d5ef)}

		if __obf_20aa71cdd9d88320%2 == 0 {
			__obf_acaa0892c9b9426a.
				Y = rand.IntN(__obf_7f6203877254d5ef) + __obf_7f6203877254d5ef*2
			__obf_5776bee544324174.
				Y = rand.IntN(__obf_7f6203877254d5ef)
		} else {
			__obf_acaa0892c9b9426a.
				Y = rand.IntN(__obf_7f6203877254d5ef) + __obf_7f6203877254d5ef*(__obf_20aa71cdd9d88320%2)
			__obf_5776bee544324174.
				Y = rand.IntN(__obf_7f6203877254d5ef) + __obf_7f6203877254d5ef*2
		}
		__obf_d55625589b36c20c := util.RandColor(__obf_d3479f2f25f21887)
		__obf_8d555da44024984a.
			//co.A = uint8(0xee)
			DrawBeeline(__obf_acaa0892c9b9426a, __obf_5776bee544324174, __obf_d55625589b36c20c)
	}
}

func (__obf_0aaa9d42f2545c73 *Captcha) DrawDotImage(__obf_ac0f4faf58af21b0 DotType, __obf_503314545b5d4fc1 *types.Dot, __obf_c31e22f60e6accb4 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_8c4c2ec57f522dc0, _ := util.ParseHexColor(__obf_503314545b5d4fc1.PrimaryColor)
	__obf_8c4c2ec57f522dc0.
		A = util.FormatAlpha(__obf_c31e22f60e6accb4.Alpha)

	var __obf_5675cd5d4389f309 image.Image
	var __obf_4f137bf3a2d5ad2c error
	if __obf_ac0f4faf58af21b0 == Shape {
		__obf_5675cd5d4389f309, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawShapeImage(__obf_503314545b5d4fc1, __obf_8c4c2ec57f522dc0)
	} else {
		__obf_5675cd5d4389f309, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawStringImage(__obf_503314545b5d4fc1, __obf_8c4c2ec57f522dc0)
	}
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, nil, __obf_4f137bf3a2d5ad2c
	}
	__obf_c3f8b29206ce82ea, _ := util.ParseHexColor(__obf_0a31bab85e53c8cf)
	__obf_60e5f15f77605018 := helper.CreateNRGBACanvas(__obf_503314545b5d4fc1.Size+10, __obf_503314545b5d4fc1.Size+10, true)
	if __obf_c31e22f60e6accb4.ShowShadow {
		var __obf_d1035e414719d44a *types.NRGBA
		if __obf_ac0f4faf58af21b0 == Shape {
			__obf_d1035e414719d44a, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawShapeImage(__obf_503314545b5d4fc1, __obf_c3f8b29206ce82ea)
		} else {
			__obf_d1035e414719d44a, __obf_4f137bf3a2d5ad2c = __obf_0aaa9d42f2545c73.DrawStringImage(__obf_503314545b5d4fc1, __obf_c3f8b29206ce82ea)
		}
		if __obf_4f137bf3a2d5ad2c != nil {
			return nil, nil, __obf_4f137bf3a2d5ad2c
		}
		__obf_117be5f0f8c79f6a := __obf_c31e22f60e6accb4.ShadowPoint.X
		__obf_8b8257406a82cb7d := __obf_c31e22f60e6accb4.ShadowPoint.Y
		draw.Draw(__obf_60e5f15f77605018, __obf_d1035e414719d44a.Bounds(), __obf_d1035e414719d44a, image.Point{X: __obf_117be5f0f8c79f6a, Y: __obf_8b8257406a82cb7d}, draw.Over)
	}
	draw.Draw(__obf_60e5f15f77605018, __obf_5675cd5d4389f309.Bounds(), __obf_5675cd5d4389f309, image.Point{}, draw.Over)
	__obf_60e5f15f77605018.
		Rotate(__obf_503314545b5d4fc1.Angle, false)
	__obf_f677063ff3cb55d0 := __obf_60e5f15f77605018.CalcMarginBlankArea()

	return __obf_60e5f15f77605018,

		// DrawStringImage draws a text image
		// params:
		//   - dot: Draw dot
		//   - textColor: Text color
		//
		// returns:
		//   - types.NRGBA: Drawn text image
		//   - error: Error information
		__obf_f677063ff3cb55d0, nil
}

func (__obf_0aaa9d42f2545c73 *Captcha) DrawStringImage(__obf_503314545b5d4fc1 *types.Dot, __obf_3d63529360f23b1c color.Color) (*types.NRGBA, error) {
	__obf_60e5f15f77605018 := helper.CreateNRGBACanvas(__obf_503314545b5d4fc1.Size+10, __obf_503314545b5d4fc1.Size+10, true)
	__obf_0bcb6aea1c516a3d := fixed.P(12, __obf_503314545b5d4fc1.Size-5)
	if util.IsChineseChar(__obf_503314545b5d4fc1.Text) {
		__obf_0bcb6aea1c516a3d = fixed.P(10, __obf_503314545b5d4fc1.Size)
	}
	__obf_c31e22f60e6accb4 := &types.DrawStringParam{
		Color:       __obf_3d63529360f23b1c,
		FontSize:    __obf_503314545b5d4fc1.FontSize,
		Size:        __obf_503314545b5d4fc1.Size,
		FontDPI:     __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.FontDPI,
		FontHinting: __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.FontHinting,
		Text:        __obf_503314545b5d4fc1.Text,
	}
	var __obf_4f137bf3a2d5ad2c error
	__obf_c31e22f60e6accb4.
		Font, __obf_4f137bf3a2d5ad2c = assets.PickFont()
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}
	__obf_4f137bf3a2d5ad2c = helper.DrawString(__obf_60e5f15f77605018, __obf_c31e22f60e6accb4, __obf_0bcb6aea1c516a3d)
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}

	return __obf_60e5f15f77605018, nil
}

func (__obf_0aaa9d42f2545c73 *Captcha) DrawShapeImage(__obf_503314545b5d4fc1 *types.Dot, __obf_8c4c2ec57f522dc0 color.Color) (*types.NRGBA, error) {
	__obf_ccea03d3ceaab864, __obf_975178e4722a81b9, __obf_478f9322cd7133d0, __obf_f14c055afef1c803 := __obf_8c4c2ec57f522dc0.RGBA()

	var __obf_4fd86ce8de3509ce = []color.RGBA{
		{R: uint8(__obf_ccea03d3ceaab864), G: uint8(__obf_975178e4722a81b9), B: uint8(__obf_478f9322cd7133d0), A: uint8(__obf_f14c055afef1c803)},
	}
	__obf_fb798136dde8d324 := helper.CreateNRGBACanvas(__obf_503314545b5d4fc1.Size+10, __obf_503314545b5d4fc1.Size+10, true)
	var __obf_516347dd0c9f207c image.Rectangle
	var __obf_a02b8d0e4a886ebd image.Image
	if __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.UseRGBA {
		__obf_60e5f15f77605018 := helper.CreateNRGBACanvas(__obf_503314545b5d4fc1.Size+10, __obf_503314545b5d4fc1.Size+10, true)
		draw.BiLinear.Scale(__obf_60e5f15f77605018, __obf_60e5f15f77605018.Bounds(), __obf_503314545b5d4fc1.Shape, __obf_503314545b5d4fc1.Shape.Bounds(), draw.Over, nil)
		__obf_516347dd0c9f207c = __obf_60e5f15f77605018.Bounds()
		__obf_a02b8d0e4a886ebd = __obf_60e5f15f77605018
	} else {
		__obf_60e5f15f77605018 := helper.CreatePaletteCanvas(__obf_503314545b5d4fc1.Size+10, __obf_503314545b5d4fc1.Size+10, __obf_4fd86ce8de3509ce)
		draw.BiLinear.Scale(__obf_60e5f15f77605018, __obf_60e5f15f77605018.Bounds(), __obf_503314545b5d4fc1.Shape, __obf_503314545b5d4fc1.Shape.Bounds(), draw.Over, nil)
		__obf_516347dd0c9f207c = __obf_60e5f15f77605018.Bounds()
		__obf_a02b8d0e4a886ebd = __obf_60e5f15f77605018
	}

	draw.Draw(__obf_fb798136dde8d324, __obf_516347dd0c9f207c, __obf_a02b8d0e4a886ebd, image.Point{}, draw.Over)

	return __obf_fb798136dde8d324, nil
}

func (__obf_de65322a0d55e5bb *Captcha) __obf_368e3993fbda0ca0() (*CaptchaData, error) {

	var (
		__obf_f7a058db0268595b, __obf_6fdc01aa4574c258, __obf_778b9733130a3fad []*types.Dot
		__obf_77626d69fa645acf                                                 []image.Image
		__obf_691773f5b8ed8b1a                                                 []image.Image
		__obf_4292b5547fe8df15, __obf_6871669979053980                         image.Image
		__obf_4f137bf3a2d5ad2c                                                 error
	)
	__obf_691773f5b8ed8b1a, __obf_4f137bf3a2d5ad2c = assets.GetShapes()
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}
	__obf_9d77e9eee553a2a3 := util.RandInt(__obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.LenRange.Min, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.LenRange.Max)
	__obf_691773f5b8ed8b1a = util.PickN(__obf_691773f5b8ed8b1a, __obf_9d77e9eee553a2a3)
	__obf_f7a058db0268595b = __obf_de65322a0d55e5bb.__obf_ecab942491810080(__obf_691773f5b8ed8b1a, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Height, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.SizeRange, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.DotPadding)
	__obf_8f7c947c8077f6aa := rand.Perm(len(__obf_f7a058db0268595b))
	__obf_c88ceb0cf69772aa := util.RandInt(__obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.VerifyLenRange.Min, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.VerifyLenRange.Max)
	__obf_778b9733130a3fad = make([]*types.Dot, min(__obf_c88ceb0cf69772aa, len(__obf_f7a058db0268595b)))
	for __obf_20aa71cdd9d88320, __obf_30e4d41cb98c70d0 := range __obf_8f7c947c8077f6aa {
		if __obf_20aa71cdd9d88320 >= __obf_c88ceb0cf69772aa {
			break
		}
		__obf_503314545b5d4fc1 := __obf_f7a058db0268595b[__obf_30e4d41cb98c70d0]
		__obf_503314545b5d4fc1.
			Index = __obf_20aa71cdd9d88320
		__obf_778b9733130a3fad[__obf_20aa71cdd9d88320] = __obf_503314545b5d4fc1
		__obf_77626d69fa645acf = append(__obf_77626d69fa645acf, __obf_778b9733130a3fad[__obf_20aa71cdd9d88320].Shape)
	}
	__obf_6fdc01aa4574c258 = __obf_de65322a0d55e5bb.__obf_ecab942491810080(__obf_77626d69fa645acf, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.SizeRange, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.DotPadding)
	__obf_4292b5547fe8df15, __obf_4f137bf3a2d5ad2c = __obf_de65322a0d55e5bb.__obf_62f552b76527b86a(Shape, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Height, __obf_f7a058db0268595b)
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}
	__obf_6871669979053980, __obf_4f137bf3a2d5ad2c = __obf_de65322a0d55e5bb.__obf_07751011a5f2078e(Shape, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Height, __obf_6fdc01aa4574c258)
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}

	return &CaptchaData{__obf_2028bb55834c38d8: __obf_778b9733130a3fad, __obf_4292b5547fe8df15: types.NewJPEGImage(__obf_4292b5547fe8df15), __obf_6871669979053980: types.NewPNGImage(__obf_6871669979053980)}, nil
}

func (__obf_de65322a0d55e5bb *Captcha) __obf_598c6b664136f15a() (*CaptchaData, error) {
	__obf_9d77e9eee553a2a3 := util.RandInt(__obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.LenRange.Min, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.LenRange.Max)
	__obf_484df0427125ab9b := util.PickN(assets.GetChineseChars(), __obf_9d77e9eee553a2a3)

	var __obf_2028bb55834c38d8, __obf_6fdc01aa4574c258, __obf_778b9733130a3fad []*types.Dot
	var __obf_102c4e877c932647 []string
	var __obf_4292b5547fe8df15, __obf_6871669979053980 image.Image
	__obf_2028bb55834c38d8 = __obf_de65322a0d55e5bb.__obf_c49a3d7710b15e8c(__obf_484df0427125ab9b, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Height, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.SizeRange, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.DotPadding)
	__obf_8f7c947c8077f6aa := rand.Perm(len(__obf_2028bb55834c38d8))
	__obf_c88ceb0cf69772aa := util.RandInt(__obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.VerifyLenRange.Min, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.VerifyLenRange.Max)
	__obf_778b9733130a3fad = make([]*types.Dot, min(__obf_c88ceb0cf69772aa, len(__obf_2028bb55834c38d8)))
	for __obf_20aa71cdd9d88320, __obf_30e4d41cb98c70d0 := range __obf_8f7c947c8077f6aa {
		if __obf_20aa71cdd9d88320 >= __obf_c88ceb0cf69772aa {
			break
		}
		__obf_503314545b5d4fc1 := __obf_2028bb55834c38d8[__obf_30e4d41cb98c70d0]
		__obf_503314545b5d4fc1.
			Index = __obf_20aa71cdd9d88320
		__obf_778b9733130a3fad[__obf_20aa71cdd9d88320] = __obf_503314545b5d4fc1
		__obf_102c4e877c932647 = append(__obf_102c4e877c932647, __obf_778b9733130a3fad[__obf_20aa71cdd9d88320].Text)
	}
	__obf_6fdc01aa4574c258 = __obf_de65322a0d55e5bb.__obf_c49a3d7710b15e8c(__obf_102c4e877c932647, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Height, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.SizeRange, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.DotPadding)
	var __obf_4f137bf3a2d5ad2c error
	__obf_4292b5547fe8df15, __obf_4f137bf3a2d5ad2c = __obf_de65322a0d55e5bb.__obf_62f552b76527b86a(Text, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Height, __obf_2028bb55834c38d8)
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}
	__obf_6871669979053980, __obf_4f137bf3a2d5ad2c = __obf_de65322a0d55e5bb.__obf_07751011a5f2078e(Text, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Width, __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Secondary.Height, __obf_6fdc01aa4574c258)
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}

	return &CaptchaData{__obf_2028bb55834c38d8: __obf_778b9733130a3fad, __obf_4292b5547fe8df15: types.NewJPEGImage(__obf_4292b5547fe8df15), __obf_6871669979053980: types.NewPNGImage(__obf_6871669979053980)}, nil
}

func (__obf_de65322a0d55e5bb *Captcha) __obf_ecab942491810080(__obf_691773f5b8ed8b1a []image.Image, __obf_a51318f96bb6cca4, __obf_0d08f75c1fe24ae6 int, __obf_75c525fff4331205 types.Range, __obf_5dd974549cda4ccc int) []*types.Dot {
	__obf_9d77e9eee553a2a3 := len(__obf_691773f5b8ed8b1a)
	var __obf_2028bb55834c38d8 = make([]*types.Dot, __obf_9d77e9eee553a2a3)
	// width := imageSize.Width
	// height := imageSize.Height
	if __obf_5dd974549cda4ccc > 0 {
		__obf_a51318f96bb6cca4 -= __obf_5dd974549cda4ccc
		__obf_0d08f75c1fe24ae6 -= __obf_5dd974549cda4ccc
	}

	for __obf_20aa71cdd9d88320, __obf_c07c64e538582e45 := range __obf_691773f5b8ed8b1a {
		__obf_16d35013830bac1b := __obf_de65322a0d55e5bb.__obf_16d35013830bac1b()
		__obf_65cee0e2ca621402 := util.PickN(__obf_1568d6d8c65a8135, 1)[0]
		__obf_7a1ffad1dcdf1c5a := util.PickN(__obf_8f2c333903421f7b, 1)[0]
		__obf_a47897ff2b371e95 := util.RandInt(__obf_75c525fff4331205.Min, __obf_75c525fff4331205.Max)
		__obf_e5934f188ae696f6 := 10
		__obf_ad6fd1603cff710a := __obf_a51318f96bb6cca4 / __obf_9d77e9eee553a2a3
		__obf_31f9f2358785a7e2 := math.Abs(float64(__obf_ad6fd1603cff710a) - float64(__obf_a47897ff2b371e95))
		__obf_7546adfbbc4f851c := (__obf_20aa71cdd9d88320 * __obf_ad6fd1603cff710a) + util.RandInt(0, int(max(__obf_31f9f2358785a7e2, 1)))
		__obf_a08987f5c7ed6b60 := util.RandInt(__obf_e5934f188ae696f6, __obf_0d08f75c1fe24ae6+__obf_a47897ff2b371e95)
		__obf_761f1461a0304967 := int(min(max(float64(__obf_7546adfbbc4f851c), float64(__obf_e5934f188ae696f6)), float64(__obf_a51318f96bb6cca4-__obf_e5934f188ae696f6-(__obf_5dd974549cda4ccc*2))))
		__obf_7f6203877254d5ef := int(min(max(float64(__obf_a08987f5c7ed6b60), float64(__obf_a47897ff2b371e95+__obf_e5934f188ae696f6)), float64(__obf_0d08f75c1fe24ae6+(__obf_a47897ff2b371e95/2)-(__obf_5dd974549cda4ccc*2))))
		__obf_2028bb55834c38d8[__obf_20aa71cdd9d88320] = &types.Dot{
			Index:          __obf_20aa71cdd9d88320,
			X:              __obf_761f1461a0304967,
			Y:              __obf_7f6203877254d5ef,
			FontSize:       __obf_a47897ff2b371e95,
			Size:           __obf_a47897ff2b371e95,
			Angle:          __obf_16d35013830bac1b,
			PrimaryColor:   __obf_65cee0e2ca621402,
			SecondaryColor: __obf_7a1ffad1dcdf1c5a,
			Shape:          __obf_c07c64e538582e45,
		}
	}

	return __obf_2028bb55834c38d8
}

func (__obf_de65322a0d55e5bb *Captcha) __obf_c49a3d7710b15e8c(__obf_484df0427125ab9b []string, __obf_a51318f96bb6cca4, __obf_0d08f75c1fe24ae6 int, __obf_5f5217239d239132 types.Range, __obf_5dd974549cda4ccc int) []*types.Dot {
	__obf_9d77e9eee553a2a3 := len(__obf_484df0427125ab9b)
	var __obf_2028bb55834c38d8 = make([]*types.Dot, __obf_9d77e9eee553a2a3)
	if __obf_5dd974549cda4ccc > 0 {
		__obf_a51318f96bb6cca4 -= __obf_5dd974549cda4ccc
		__obf_0d08f75c1fe24ae6 -= __obf_5dd974549cda4ccc
	}

	for __obf_20aa71cdd9d88320, __obf_c07c64e538582e45 := range __obf_484df0427125ab9b {
		__obf_16d35013830bac1b := __obf_de65322a0d55e5bb.__obf_16d35013830bac1b()
		__obf_65cee0e2ca621402 := util.PickN(__obf_1568d6d8c65a8135, 1)[0]
		__obf_7a1ffad1dcdf1c5a := util.PickN(__obf_8f2c333903421f7b, 1)[0]
		__obf_a47897ff2b371e95 := util.RandInt(__obf_5f5217239d239132.Min, __obf_5f5217239d239132.Max)
		__obf_571c8cbd80220ceb := __obf_a47897ff2b371e95
		__obf_433d3be771e52532 := __obf_a47897ff2b371e95

		if util.LenChineseChar(__obf_c07c64e538582e45) > 1 {
			__obf_433d3be771e52532 = __obf_a47897ff2b371e95 * util.LenChineseChar(__obf_c07c64e538582e45)

			if __obf_16d35013830bac1b > 0 {
				__obf_3d0ab02701f5c870 := __obf_433d3be771e52532 - __obf_a47897ff2b371e95
				__obf_1e8a8ce3801ffe08 := __obf_16d35013830bac1b % 90
				__obf_9908da40cdbef92e := float64(__obf_3d0ab02701f5c870) / 90
				__obf_297d7fb14f68ca1c := max(float64(__obf_1e8a8ce3801ffe08)*__obf_9908da40cdbef92e, 1)
				__obf_571c8cbd80220ceb = __obf_571c8cbd80220ceb + int(__obf_297d7fb14f68ca1c)
				__obf_433d3be771e52532 = __obf_433d3be771e52532 + int(__obf_297d7fb14f68ca1c)
			}
		}
		__obf_e5934f188ae696f6 := 10
		__obf_ad6fd1603cff710a := __obf_a51318f96bb6cca4 / __obf_9d77e9eee553a2a3
		__obf_31f9f2358785a7e2 := math.Abs(float64(__obf_ad6fd1603cff710a) - float64(__obf_433d3be771e52532))
		__obf_7546adfbbc4f851c := (__obf_20aa71cdd9d88320 * __obf_ad6fd1603cff710a) + util.RandInt(0, int(max(__obf_31f9f2358785a7e2, 1)))
		__obf_a08987f5c7ed6b60 := util.RandInt(__obf_e5934f188ae696f6, __obf_0d08f75c1fe24ae6+__obf_571c8cbd80220ceb)
		__obf_761f1461a0304967 := int(min(max(float64(__obf_7546adfbbc4f851c), float64(__obf_e5934f188ae696f6)), float64(__obf_a51318f96bb6cca4-__obf_e5934f188ae696f6-(__obf_5dd974549cda4ccc*2))))
		__obf_7f6203877254d5ef := int(min(max(float64(__obf_a08987f5c7ed6b60), float64(__obf_571c8cbd80220ceb+__obf_e5934f188ae696f6)), float64(__obf_0d08f75c1fe24ae6+(__obf_571c8cbd80220ceb/2)-(__obf_5dd974549cda4ccc*2))))
		__obf_2028bb55834c38d8[__obf_20aa71cdd9d88320] = &types.Dot{
			Index:          __obf_20aa71cdd9d88320,
			X:              __obf_761f1461a0304967,
			Y:              __obf_7f6203877254d5ef,
			FontSize:       __obf_a47897ff2b371e95,
			Size:           max(__obf_433d3be771e52532, __obf_571c8cbd80220ceb),
			Angle:          __obf_16d35013830bac1b,
			PrimaryColor:   __obf_65cee0e2ca621402,
			SecondaryColor: __obf_7a1ffad1dcdf1c5a,
			Text:           __obf_c07c64e538582e45,
		}
	}

	return __obf_2028bb55834c38d8
}

func (__obf_de65322a0d55e5bb *Captcha) __obf_62f552b76527b86a(__obf_ac0f4faf58af21b0 DotType, __obf_a51318f96bb6cca4, __obf_0d08f75c1fe24ae6 int, __obf_2028bb55834c38d8 []*types.Dot) (image.Image, error) {
	__obf_8aae7e2e801f3bf6 := &DrawImageParams{
		Width:       __obf_a51318f96bb6cca4,
		Height:      __obf_0d08f75c1fe24ae6,
		Alpha:       __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.Alpha,
		FontHinting: __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.FontHinting,
		Dots:        __obf_2028bb55834c38d8,
		ShowShadow:  __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.ShowShadow,
		ShadowPoint: __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.ShadowPoint,
	}
	var __obf_4f137bf3a2d5ad2c error
	__obf_8aae7e2e801f3bf6.
		Background, __obf_4f137bf3a2d5ad2c = assets.PickBgImage()
	if __obf_4f137bf3a2d5ad2c != nil {
		return nil, __obf_4f137bf3a2d5ad2c
	}

	return __obf_de65322a0d55e5bb.DrawWithNRGBA(__obf_ac0f4faf58af21b0, __obf_8aae7e2e801f3bf6)
}

func (__obf_0aaa9d42f2545c73 *Captcha) __obf_07751011a5f2078e(__obf_ac0f4faf58af21b0 DotType, __obf_a51318f96bb6cca4, __obf_0d08f75c1fe24ae6 int, __obf_2028bb55834c38d8 []*types.Dot) (image.Image, error) {
	var (
		__obf_f479913e373334ce =
		// err      error
		make([]*types.Dot, 0, len(__obf_2028bb55834c38d8))
	)
	__obf_1b1d6efe1ca4f4e7 := __obf_a51318f96bb6cca4 / len(__obf_2028bb55834c38d8)

	for __obf_20aa71cdd9d88320, __obf_c07c64e538582e45 := range __obf_2028bb55834c38d8 {
		__obf_9d77e9eee553a2a3 := 1
		if __obf_ac0f4faf58af21b0 == Text {
			__obf_9d77e9eee553a2a3 = len(__obf_c07c64e538582e45.Text)
		}
		__obf_c07c64e538582e45.
			X = int(max(float64(__obf_1b1d6efe1ca4f4e7*__obf_20aa71cdd9d88320+__obf_1b1d6efe1ca4f4e7/__obf_c07c64e538582e45.Size), 8))
		__obf_4cbeb18ca7dfde39 := max(1, __obf_0d08f75c1fe24ae6/16*__obf_9d77e9eee553a2a3)
		__obf_c07c64e538582e45.
			Y = __obf_0d08f75c1fe24ae6/2 + __obf_c07c64e538582e45.FontSize/2 - rand.IntN(__obf_4cbeb18ca7dfde39)
		__obf_f479913e373334ce = append(__obf_f479913e373334ce, __obf_c07c64e538582e45)
	}
	__obf_c31e22f60e6accb4 := &DrawImageParams{
		Width:             __obf_a51318f96bb6cca4,
		Height:            __obf_0d08f75c1fe24ae6,
		Dots:              __obf_f479913e373334ce,
		BackgroundDistort: __obf_0aaa9d42f2545c73.__obf_05dd88551428dd9b(__obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.Secondary.BgDistort),
		BgCircles:         __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.Secondary.BgCircles,
		BgSlimLines:       __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_95f7cfdd531243ee []color.Color
	for _, __obf_8c0a79a1abbb8020 := range __obf_8f2c333903421f7b {
		__obf_d55625589b36c20c, _ := util.ParseHexColor(__obf_8c0a79a1abbb8020)
		__obf_95f7cfdd531243ee = append(__obf_95f7cfdd531243ee, __obf_d55625589b36c20c)
	}

	var __obf_ef2f557a68e49262 []color.Color
	for _, __obf_d55625589b36c20c := range __obf_8f2c333903421f7b {
		__obf_8b68dea5ea6d2497, _ := util.ParseHexColor(__obf_d55625589b36c20c)
		__obf_ef2f557a68e49262 = append(__obf_ef2f557a68e49262, __obf_8b68dea5ea6d2497)
	}

	if __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.UseRGBA || __obf_0aaa9d42f2545c73.__obf_8659ee70bdbf57ab.Secondary.NonDeformAble {
		return __obf_0aaa9d42f2545c73.DrawWithNRGBA2(__obf_ac0f4faf58af21b0, __obf_c31e22f60e6accb4, __obf_95f7cfdd531243ee, __obf_ef2f557a68e49262)
	}
	return __obf_0aaa9d42f2545c73.DrawWithPalette(__obf_ac0f4faf58af21b0, __obf_c31e22f60e6accb4, __obf_95f7cfdd531243ee, __obf_ef2f557a68e49262)
}

func (__obf_de65322a0d55e5bb *Captcha) __obf_05dd88551428dd9b(__obf_1cd84af0ddc4873c int) int {
	switch __obf_1cd84af0ddc4873c {
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

func (__obf_de65322a0d55e5bb *Captcha) __obf_16d35013830bac1b() int {
	__obf_e0ca67c41c6eb046 := __obf_de65322a0d55e5bb.__obf_8659ee70bdbf57ab.Primary.AnglePosRange
	__obf_f2409b4fa6b679c3 := rand.IntN(len(__obf_e0ca67c41c6eb046))
	if __obf_f2409b4fa6b679c3 < 0 {
		return 0
	}
	__obf_c0d7e0e24eaa5d27 := __obf_e0ca67c41c6eb046[__obf_f2409b4fa6b679c3]
	__obf_60bced6ca856de2a := util.RandInt(__obf_c0d7e0e24eaa5d27.Min, __obf_c0d7e0e24eaa5d27.Max)

	return __obf_60bced6ca856de2a
}
