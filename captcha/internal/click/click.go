package __obf_a334e8eb146672ab

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
	__obf_3d43bbdf2fa5cdea *types.ClickOption
}

var _ types.Captcha = (*Captcha)(nil)

func NewCaptcha(__obf_3d43bbdf2fa5cdea *types.ClickOption) *Captcha {
	if __obf_3d43bbdf2fa5cdea == nil {
		return nil
	}
	return &Captcha{__obf_3d43bbdf2fa5cdea}
}

func (__obf_6fc320ffbfc77017 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_c112f65c651fcf04 := __obf_22348947b5a83ed0(__obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea)
	if __obf_c112f65c651fcf04 == Shape {
		return __obf_6fc320ffbfc77017.__obf_624ec29df8fa8452()
	}

	return __obf_6fc320ffbfc77017.__obf_b52fbb7a70540196()
}

func __obf_22348947b5a83ed0(__obf_3d43bbdf2fa5cdea *types.ClickOption) DotType {
	__obf_a6730603b7baa33b := DotType(rand.IntN(2))
	if __obf_a6730603b7baa33b == Shape {
		__obf_3d43bbdf2fa5cdea.
			UseRGBA = false
		__obf_3d43bbdf2fa5cdea.
			Secondary.BgDistort = types.DistortLevel1
		__obf_3d43bbdf2fa5cdea.
			Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_3d43bbdf2fa5cdea.
			Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_3d43bbdf2fa5cdea.
			UseRGBA = true
		__obf_3d43bbdf2fa5cdea.
			Secondary.BgDistort = types.DistortLevel4
		__obf_3d43bbdf2fa5cdea.
			Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_3d43bbdf2fa5cdea.
			Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_a6730603b7baa33b
}

func (__obf_0f1aca653a04363a *Captcha) DrawWithNRGBA(__obf_c112f65c651fcf04 DotType, __obf_ece4016def2d6db0 *DrawImageParams) (image.Image, error) {
	__obf_f7494f398db65194 := __obf_ece4016def2d6db0.Dots
	__obf_806b828cee04f014 := helper.CreateNRGBACanvas(__obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height, true)

	for _, __obf_014ab0047358450e := range __obf_f7494f398db65194 {
		__obf_0243467e06c1e7af, __obf_dc9373f1c7ca11cd, __obf_48243aced6a1b5c5 := __obf_0f1aca653a04363a.DrawDotImage(__obf_c112f65c651fcf04, __obf_014ab0047358450e, __obf_ece4016def2d6db0)
		if __obf_48243aced6a1b5c5 != nil {
			return nil, __obf_48243aced6a1b5c5
		}
		__obf_b87c20462956b39a := __obf_dc9373f1c7ca11cd.MinX
		__obf_7ee02c32fac258ea := __obf_dc9373f1c7ca11cd.MaxX
		__obf_19ef218cafaba4ca := __obf_dc9373f1c7ca11cd.MinY
		__obf_d3199103df576771 := __obf_dc9373f1c7ca11cd.MaxY
		__obf_f246da57f4cb3862 := __obf_7ee02c32fac258ea - __obf_b87c20462956b39a
		__obf_9091d3f6b7268dcd := __obf_d3199103df576771 - __obf_19ef218cafaba4ca

		draw.Draw(__obf_806b828cee04f014, image.Rect(__obf_014ab0047358450e.X, __obf_014ab0047358450e.Y, __obf_014ab0047358450e.X+__obf_f246da57f4cb3862, __obf_014ab0047358450e.Y+__obf_9091d3f6b7268dcd), __obf_0243467e06c1e7af, image.Point{X: __obf_b87c20462956b39a, Y: __obf_19ef218cafaba4ca}, draw.Over)

	}
	__obf_07790da0efb07292 := __obf_ece4016def2d6db0.Background
	__obf_7ce3c8b6df694560 := __obf_806b828cee04f014.Bounds()
	__obf_91f3dcf2cb98d330 := helper.CreateNRGBACanvas(__obf_7ce3c8b6df694560.Dx(), __obf_7ce3c8b6df694560.Dy(), true)
	__obf_0c9d74577036e6b5 := util.RangCutImagePos(__obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height, __obf_07790da0efb07292)
	draw.Draw(__obf_91f3dcf2cb98d330, __obf_7ce3c8b6df694560, __obf_07790da0efb07292, __obf_0c9d74577036e6b5, draw.Src)
	draw.Draw(__obf_91f3dcf2cb98d330, __obf_806b828cee04f014.Bounds(), __obf_806b828cee04f014, image.Point{}, draw.Over)
	__obf_91f3dcf2cb98d330.
		SubImage(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height))
	return __obf_91f3dcf2cb98d330, nil
}

func (__obf_0f1aca653a04363a *Captcha) DrawWithPalette(__obf_c112f65c651fcf04 DotType, __obf_ece4016def2d6db0 *DrawImageParams, __obf_76da2cb11cb229d6 []color.Color, __obf_dbca0f5a47406936 []color.Color) (image.Image, error) {
	__obf_f7494f398db65194 := __obf_ece4016def2d6db0.Dots
	__obf_478952763442d7a0 := make([]color.Color, 0, len(__obf_dbca0f5a47406936))
	for _, __obf_41e10907260cbbe7 := range __obf_dbca0f5a47406936 {
		__obf_a52cd80f96bb60c8, __obf_379d806d340fa807, __obf_7ce3c8b6df694560, _ := __obf_41e10907260cbbe7.RGBA()
		__obf_56affa2bb097feaa := util.FormatAlpha(__obf_ece4016def2d6db0.ThumbDisturbAlpha)
		__obf_478952763442d7a0 = append(__obf_478952763442d7a0, color.RGBA{R: uint8(__obf_a52cd80f96bb60c8), G: uint8(__obf_379d806d340fa807), B: uint8(__obf_7ce3c8b6df694560), A: __obf_56affa2bb097feaa})
	}

	var __obf_3882b974f07c2437 = make([]color.Color, 0, len(__obf_76da2cb11cb229d6)+len(__obf_478952763442d7a0)+1)
	__obf_3882b974f07c2437 = append(__obf_3882b974f07c2437, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_3882b974f07c2437 = append(__obf_3882b974f07c2437, __obf_76da2cb11cb229d6...)
	__obf_3882b974f07c2437 = append(__obf_3882b974f07c2437, __obf_478952763442d7a0...)
	__obf_806b828cee04f014 := types.NewPalette(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height), __obf_3882b974f07c2437)
	if __obf_ece4016def2d6db0.BgCircles > 0 {
		__obf_0f1aca653a04363a.__obf_4dfc37217a033611(__obf_806b828cee04f014, __obf_ece4016def2d6db0.BgCircles, 1, __obf_478952763442d7a0)
	}
	if __obf_ece4016def2d6db0.BgSlimLines > 0 {
		__obf_0f1aca653a04363a.__obf_d167968c22e0fd2b(__obf_806b828cee04f014, __obf_ece4016def2d6db0.BgSlimLines, __obf_478952763442d7a0)
	}

	for _, __obf_014ab0047358450e := range __obf_f7494f398db65194 {
		__obf_f09080f4322096ff, _ := util.ParseHexColor(__obf_014ab0047358450e.PrimaryColor)
		var __obf_48243aced6a1b5c5 error
		if __obf_c112f65c651fcf04 == Shape {
			var __obf_0243467e06c1e7af *types.NRGBA
			__obf_0243467e06c1e7af, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawShapeImage(__obf_014ab0047358450e, __obf_f09080f4322096ff)
			if __obf_48243aced6a1b5c5 != nil {
				return nil, __obf_48243aced6a1b5c5
			}
			__obf_0243467e06c1e7af.
				Rotate(__obf_014ab0047358450e.Angle, false)
			__obf_3866f8eef4d73f02 := __obf_806b828cee04f014.Bounds()
			__obf_93c1d1656503215d := __obf_0243467e06c1e7af.Bounds()
			__obf_5be6366012432e63 := image.Point{X: __obf_3866f8eef4d73f02.Dx() - __obf_93c1d1656503215d.Dx(), Y: __obf_3866f8eef4d73f02.Dy() - __obf_93c1d1656503215d.Dy()}
			draw.Draw(__obf_806b828cee04f014, image.Rect(__obf_014ab0047358450e.X, __obf_5be6366012432e63.Y, __obf_5be6366012432e63.X+__obf_93c1d1656503215d.Dx(), __obf_5be6366012432e63.Y+__obf_93c1d1656503215d.Dy()), __obf_0243467e06c1e7af, image.Point{}, draw.Over)
		} else {
			__obf_8c7728b444595a48 := fixed.P(__obf_014ab0047358450e.X, __obf_014ab0047358450e.Y)
			__obf_ece4016def2d6db0 := &types.DrawStringParam{
				Color:       __obf_f09080f4322096ff,
				FontSize:    __obf_014ab0047358450e.FontSize,
				Size:        __obf_014ab0047358450e.Size,
				FontDPI:     __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.FontDPI,
				FontHinting: __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.FontHinting,
				Text:        __obf_014ab0047358450e.Text,
			}
			__obf_ece4016def2d6db0.
				Font, __obf_48243aced6a1b5c5 = assets.PickFont()
			if __obf_48243aced6a1b5c5 != nil {
				return nil, __obf_48243aced6a1b5c5
			}
			__obf_48243aced6a1b5c5 = helper.DrawString(__obf_806b828cee04f014, __obf_ece4016def2d6db0, __obf_8c7728b444595a48)
		}

		if __obf_48243aced6a1b5c5 != nil {
			return __obf_806b828cee04f014, __obf_48243aced6a1b5c5
		}
	}

	if __obf_ece4016def2d6db0.Background != nil {
		__obf_07790da0efb07292 := __obf_ece4016def2d6db0.Background
		__obf_7ce3c8b6df694560 := __obf_07790da0efb07292.Bounds()
		__obf_91f3dcf2cb98d330 := helper.CreateNRGBACanvas(__obf_7ce3c8b6df694560.Dx(), __obf_7ce3c8b6df694560.Dy(), true)
		__obf_0c9d74577036e6b5 := util.RangCutImagePos(__obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height, __obf_07790da0efb07292)
		draw.Draw(__obf_91f3dcf2cb98d330, __obf_7ce3c8b6df694560, __obf_07790da0efb07292, __obf_0c9d74577036e6b5, draw.Src)
		__obf_806b828cee04f014.
			Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_91f3dcf2cb98d330, __obf_806b828cee04f014.Bounds(), __obf_806b828cee04f014, image.Point{}, draw.Over)
		__obf_7778291a332582b1 := __obf_91f3dcf2cb98d330.SubImage(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height)).(*image.NRGBA)
		return __obf_7778291a332582b1, nil
	}

	if __obf_ece4016def2d6db0.BackgroundDistort > 0 {
		__obf_806b828cee04f014.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_ece4016def2d6db0.BackgroundDistort))
	}

	return __obf_806b828cee04f014, nil

}

func (__obf_0f1aca653a04363a *Captcha) DrawWithNRGBA2(__obf_c112f65c651fcf04 DotType, __obf_ece4016def2d6db0 *DrawImageParams, __obf_76da2cb11cb229d6 []color.Color, __obf_dbca0f5a47406936 []color.Color) (image.Image, error) {
	__obf_f7494f398db65194 := __obf_ece4016def2d6db0.Dots
	__obf_478952763442d7a0 := make([]color.Color, 0, len(__obf_dbca0f5a47406936))
	for _, __obf_41e10907260cbbe7 := range __obf_dbca0f5a47406936 {
		__obf_a52cd80f96bb60c8, __obf_379d806d340fa807, __obf_7ce3c8b6df694560, _ := __obf_41e10907260cbbe7.RGBA()
		__obf_56affa2bb097feaa := util.FormatAlpha(__obf_ece4016def2d6db0.ThumbDisturbAlpha)
		__obf_478952763442d7a0 = append(__obf_478952763442d7a0, color.RGBA{R: uint8(__obf_a52cd80f96bb60c8), G: uint8(__obf_379d806d340fa807), B: uint8(__obf_7ce3c8b6df694560), A: __obf_56affa2bb097feaa})
	}

	var __obf_3882b974f07c2437 = make([]color.Color, 0, len(__obf_76da2cb11cb229d6)+len(__obf_478952763442d7a0)+1)
	__obf_3882b974f07c2437 = append(__obf_3882b974f07c2437, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_3882b974f07c2437 = append(__obf_3882b974f07c2437, __obf_76da2cb11cb229d6...)
	__obf_3882b974f07c2437 = append(__obf_3882b974f07c2437, __obf_478952763442d7a0...)
	__obf_66005b290575b0e6 := types.NewNRGBA(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height), true)
	if __obf_ece4016def2d6db0.Background != nil {
		__obf_07790da0efb07292 := __obf_ece4016def2d6db0.Background
		__obf_7ce3c8b6df694560 := __obf_07790da0efb07292.Bounds()
		__obf_91f3dcf2cb98d330 := helper.CreateNRGBACanvas(__obf_7ce3c8b6df694560.Dx(), __obf_7ce3c8b6df694560.Dy(), true)
		__obf_0c9d74577036e6b5 := util.RangCutImagePos(__obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height, __obf_07790da0efb07292)
		draw.Draw(__obf_91f3dcf2cb98d330, __obf_7ce3c8b6df694560, __obf_07790da0efb07292, __obf_0c9d74577036e6b5, draw.Src)
		__obf_7778291a332582b1 := __obf_91f3dcf2cb98d330.SubImage(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height)).(*image.NRGBA)
		draw.Draw(__obf_66005b290575b0e6, __obf_7778291a332582b1.Bounds(), __obf_7778291a332582b1, image.Point{}, draw.Over)
	}
	__obf_806b828cee04f014 := types.NewPalette(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height), __obf_3882b974f07c2437)
	if __obf_ece4016def2d6db0.BgCircles > 0 {
		__obf_0f1aca653a04363a.__obf_4dfc37217a033611(__obf_806b828cee04f014, __obf_ece4016def2d6db0.BgCircles, 1, __obf_478952763442d7a0)
	}
	if __obf_ece4016def2d6db0.BgSlimLines > 0 {
		__obf_0f1aca653a04363a.__obf_d167968c22e0fd2b(__obf_806b828cee04f014, __obf_ece4016def2d6db0.BgSlimLines, __obf_478952763442d7a0)
	}
	if __obf_ece4016def2d6db0.BackgroundDistort > 0 {
		__obf_806b828cee04f014.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_ece4016def2d6db0.BackgroundDistort))
	}
	__obf_3a1cc835cf885d39 := __obf_806b828cee04f014.Bounds()
	__obf_f246da57f4cb3862 := __obf_3a1cc835cf885d39.Dx() / len(__obf_f7494f398db65194)
	for __obf_1f520fc8bd1d0fba, __obf_014ab0047358450e := range __obf_f7494f398db65194 {
		__obf_f09080f4322096ff, _ := util.ParseHexColor(__obf_014ab0047358450e.PrimaryColor)
		var __obf_48243aced6a1b5c5 error
		if __obf_c112f65c651fcf04 == Shape {
			var __obf_0243467e06c1e7af *types.NRGBA
			__obf_0243467e06c1e7af, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawShapeImage(__obf_014ab0047358450e, __obf_f09080f4322096ff)
			if __obf_48243aced6a1b5c5 != nil {
				return nil, __obf_48243aced6a1b5c5
			}
			__obf_0243467e06c1e7af.
				Rotate(__obf_014ab0047358450e.Angle, false)
			__obf_3866f8eef4d73f02 := __obf_66005b290575b0e6.Bounds()
			__obf_93c1d1656503215d := __obf_0243467e06c1e7af.Bounds()
			__obf_5be6366012432e63 := image.Point{X: __obf_3866f8eef4d73f02.Dx() - __obf_93c1d1656503215d.Dx(), Y: __obf_3866f8eef4d73f02.Dy() - __obf_93c1d1656503215d.Dy()}
			draw.Draw(__obf_66005b290575b0e6, image.Rect(__obf_014ab0047358450e.X, __obf_5be6366012432e63.Y, __obf_5be6366012432e63.X+__obf_93c1d1656503215d.Dx(), __obf_5be6366012432e63.Y+__obf_93c1d1656503215d.Dy()), __obf_0243467e06c1e7af, image.Point{}, draw.Over)
		} else {
			var __obf_b7ede88b595de28d *types.NRGBA
			__obf_b7ede88b595de28d, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawStringImage(__obf_014ab0047358450e, __obf_f09080f4322096ff)
			if __obf_48243aced6a1b5c5 != nil {
				return nil, __obf_48243aced6a1b5c5
			}
			__obf_b7ede88b595de28d.
				Rotate(__obf_014ab0047358450e.Angle, false)
			__obf_dc9373f1c7ca11cd := __obf_b7ede88b595de28d.CalcMarginBlankArea()
			__obf_b87c20462956b39a := __obf_dc9373f1c7ca11cd.MinX
			__obf_7ee02c32fac258ea := __obf_dc9373f1c7ca11cd.MaxX
			__obf_19ef218cafaba4ca := __obf_dc9373f1c7ca11cd.MinY
			__obf_d3199103df576771 := __obf_dc9373f1c7ca11cd.MaxY
			__obf_b7ede88b595de28d.
				SubImage(image.Rect(__obf_b87c20462956b39a, __obf_19ef218cafaba4ca, __obf_7ee02c32fac258ea, __obf_d3199103df576771))
			__obf_036d52f5e09fd82a := __obf_b7ede88b595de28d.Bounds()
			__obf_597f0c6bfbc4a84a := int(max(float64(__obf_f246da57f4cb3862*__obf_1f520fc8bd1d0fba+__obf_f246da57f4cb3862/__obf_036d52f5e09fd82a.Dx()), 8))
			__obf_cc9401c17a2e2f22 := util.RandInt(1, __obf_3a1cc835cf885d39.Dy()-__obf_036d52f5e09fd82a.Dy()-4)

			draw.Draw(__obf_66005b290575b0e6, image.Rect(__obf_597f0c6bfbc4a84a, __obf_cc9401c17a2e2f22, __obf_597f0c6bfbc4a84a+__obf_036d52f5e09fd82a.Dx(), __obf_cc9401c17a2e2f22+__obf_036d52f5e09fd82a.Dy()), __obf_b7ede88b595de28d, image.Point{X: __obf_036d52f5e09fd82a.Min.X, Y: __obf_036d52f5e09fd82a.Min.Y}, draw.Over)
		}

		if __obf_48243aced6a1b5c5 != nil {
			return __obf_66005b290575b0e6, __obf_48243aced6a1b5c5
		}
	}
	__obf_628ca96b9f843b5d := types.NewNRGBA(image.Rect(0, 0, __obf_ece4016def2d6db0.Width, __obf_ece4016def2d6db0.Height), true)
	draw.Draw(__obf_628ca96b9f843b5d, __obf_806b828cee04f014.Bounds(), __obf_806b828cee04f014, image.Point{}, draw.Over)
	draw.Draw(__obf_66005b290575b0e6, __obf_628ca96b9f843b5d.Bounds(), __obf_628ca96b9f843b5d, image.Point{}, draw.Over)
	return __obf_66005b290575b0e6, nil
}

func (__obf_0f1aca653a04363a *Captcha) __obf_4dfc37217a033611(__obf_91f3dcf2cb98d330 *types.Palette, __obf_94ee5c5974c1fdf3, __obf_e624e4e742a788b0 int, __obf_9de6997a6ba0c29f []color.Color) {
	__obf_07f58b888f8bd4ec := __obf_91f3dcf2cb98d330.Bounds().Max.X
	__obf_482a2a6f426b9961 := __obf_91f3dcf2cb98d330.Bounds().Max.Y
	for range __obf_94ee5c5974c1fdf3 {
		__obf_10213221b9f3427b := util.RandColor(__obf_9de6997a6ba0c29f)
		__obf_a52cd80f96bb60c8 := //co.A = uint8(0xee)
			util.RandInt(1, __obf_e624e4e742a788b0)
		__obf_91f3dcf2cb98d330.
			DrawCircle(util.RandInt(__obf_a52cd80f96bb60c8, __obf_07f58b888f8bd4ec-__obf_a52cd80f96bb60c8), util.RandInt(__obf_a52cd80f96bb60c8, __obf_482a2a6f426b9961-__obf_a52cd80f96bb60c8), __obf_a52cd80f96bb60c8, __obf_10213221b9f3427b)
	}
}

func (__obf_0f1aca653a04363a *Captcha) __obf_d167968c22e0fd2b(__obf_91f3dcf2cb98d330 *types.Palette, __obf_11310637faef2a4b int, __obf_9de6997a6ba0c29f []color.Color) {
	__obf_5d48adee8878cf92 := __obf_91f3dcf2cb98d330.Bounds().Max.X / 10
	__obf_10d3aa83380337e6 := __obf_5d48adee8878cf92 * 9
	__obf_e7f217bb23bb3389 := __obf_91f3dcf2cb98d330.Bounds().Max.Y / 3
	for __obf_1f520fc8bd1d0fba := range __obf_11310637faef2a4b {
		__obf_5374c0702932b5fb := image.Point{X: rand.IntN(__obf_5d48adee8878cf92), Y: rand.IntN(__obf_e7f217bb23bb3389)}
		__obf_3fb433e2756dfe23 := image.Point{X: rand.IntN(__obf_5d48adee8878cf92) + __obf_10d3aa83380337e6, Y: rand.IntN(__obf_e7f217bb23bb3389)}

		if __obf_1f520fc8bd1d0fba%2 == 0 {
			__obf_5374c0702932b5fb.
				Y = rand.IntN(__obf_e7f217bb23bb3389) + __obf_e7f217bb23bb3389*2
			__obf_3fb433e2756dfe23.
				Y = rand.IntN(__obf_e7f217bb23bb3389)
		} else {
			__obf_5374c0702932b5fb.
				Y = rand.IntN(__obf_e7f217bb23bb3389) + __obf_e7f217bb23bb3389*(__obf_1f520fc8bd1d0fba%2)
			__obf_3fb433e2756dfe23.
				Y = rand.IntN(__obf_e7f217bb23bb3389) + __obf_e7f217bb23bb3389*2
		}
		__obf_10213221b9f3427b := util.RandColor(__obf_9de6997a6ba0c29f)
		__obf_91f3dcf2cb98d330.
			//co.A = uint8(0xee)
			DrawBeeline(__obf_5374c0702932b5fb, __obf_3fb433e2756dfe23, __obf_10213221b9f3427b)
	}
}

func (__obf_0f1aca653a04363a *Captcha) DrawDotImage(__obf_c112f65c651fcf04 DotType, __obf_9bab71876408a0b2 *types.Dot, __obf_ece4016def2d6db0 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_f09080f4322096ff, _ := util.ParseHexColor(__obf_9bab71876408a0b2.PrimaryColor)
	__obf_f09080f4322096ff.
		A = util.FormatAlpha(__obf_ece4016def2d6db0.Alpha)

	var __obf_b7ede88b595de28d image.Image
	var __obf_48243aced6a1b5c5 error
	if __obf_c112f65c651fcf04 == Shape {
		__obf_b7ede88b595de28d, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawShapeImage(__obf_9bab71876408a0b2, __obf_f09080f4322096ff)
	} else {
		__obf_b7ede88b595de28d, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawStringImage(__obf_9bab71876408a0b2, __obf_f09080f4322096ff)
	}
	if __obf_48243aced6a1b5c5 != nil {
		return nil, nil, __obf_48243aced6a1b5c5
	}
	__obf_d1df1441764df5fe, _ := util.ParseHexColor(__obf_8325939935425490)
	__obf_806b828cee04f014 := helper.CreateNRGBACanvas(__obf_9bab71876408a0b2.Size+10, __obf_9bab71876408a0b2.Size+10, true)
	if __obf_ece4016def2d6db0.ShowShadow {
		var __obf_bf64971dc2165fa8 *types.NRGBA
		if __obf_c112f65c651fcf04 == Shape {
			__obf_bf64971dc2165fa8, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawShapeImage(__obf_9bab71876408a0b2, __obf_d1df1441764df5fe)
		} else {
			__obf_bf64971dc2165fa8, __obf_48243aced6a1b5c5 = __obf_0f1aca653a04363a.DrawStringImage(__obf_9bab71876408a0b2, __obf_d1df1441764df5fe)
		}
		if __obf_48243aced6a1b5c5 != nil {
			return nil, nil, __obf_48243aced6a1b5c5
		}
		__obf_f9ccbd125e487849 := __obf_ece4016def2d6db0.ShadowPoint.X
		__obf_dae6ed737a1e9b5f := __obf_ece4016def2d6db0.ShadowPoint.Y
		draw.Draw(__obf_806b828cee04f014, __obf_bf64971dc2165fa8.Bounds(), __obf_bf64971dc2165fa8, image.Point{X: __obf_f9ccbd125e487849, Y: __obf_dae6ed737a1e9b5f}, draw.Over)
	}
	draw.Draw(__obf_806b828cee04f014, __obf_b7ede88b595de28d.Bounds(), __obf_b7ede88b595de28d, image.Point{}, draw.Over)
	__obf_806b828cee04f014.
		Rotate(__obf_9bab71876408a0b2.Angle, false)
	__obf_a5343b646e5c7254 := __obf_806b828cee04f014.CalcMarginBlankArea()

	return __obf_806b828cee04f014,

		// DrawStringImage draws a text image
		// params:
		//   - dot: Draw dot
		//   - textColor: Text color
		//
		// returns:
		//   - types.NRGBA: Drawn text image
		//   - error: Error information
		__obf_a5343b646e5c7254, nil
}

func (__obf_0f1aca653a04363a *Captcha) DrawStringImage(__obf_9bab71876408a0b2 *types.Dot, __obf_f8edc2d8b7705068 color.Color) (*types.NRGBA, error) {
	__obf_806b828cee04f014 := helper.CreateNRGBACanvas(__obf_9bab71876408a0b2.Size+10, __obf_9bab71876408a0b2.Size+10, true)
	__obf_8c7728b444595a48 := fixed.P(12, __obf_9bab71876408a0b2.Size-5)
	if util.IsChineseChar(__obf_9bab71876408a0b2.Text) {
		__obf_8c7728b444595a48 = fixed.P(10, __obf_9bab71876408a0b2.Size)
	}
	__obf_ece4016def2d6db0 := &types.DrawStringParam{
		Color:       __obf_f8edc2d8b7705068,
		FontSize:    __obf_9bab71876408a0b2.FontSize,
		Size:        __obf_9bab71876408a0b2.Size,
		FontDPI:     __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.FontDPI,
		FontHinting: __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.FontHinting,
		Text:        __obf_9bab71876408a0b2.Text,
	}
	var __obf_48243aced6a1b5c5 error
	__obf_ece4016def2d6db0.
		Font, __obf_48243aced6a1b5c5 = assets.PickFont()
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}
	__obf_48243aced6a1b5c5 = helper.DrawString(__obf_806b828cee04f014, __obf_ece4016def2d6db0, __obf_8c7728b444595a48)
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}

	return __obf_806b828cee04f014, nil
}

func (__obf_0f1aca653a04363a *Captcha) DrawShapeImage(__obf_9bab71876408a0b2 *types.Dot, __obf_f09080f4322096ff color.Color) (*types.NRGBA, error) {
	__obf_fc0708805ae85111, __obf_54d2ec88adc9e9ec, __obf_2c930c003e4a234b, __obf_40fa2d7e306c67bf := __obf_f09080f4322096ff.RGBA()

	var __obf_3c7fe566d0291378 = []color.RGBA{
		{R: uint8(__obf_fc0708805ae85111), G: uint8(__obf_54d2ec88adc9e9ec), B: uint8(__obf_2c930c003e4a234b), A: uint8(__obf_40fa2d7e306c67bf)},
	}
	__obf_628ca96b9f843b5d := helper.CreateNRGBACanvas(__obf_9bab71876408a0b2.Size+10, __obf_9bab71876408a0b2.Size+10, true)
	var __obf_036d52f5e09fd82a image.Rectangle
	var __obf_07790da0efb07292 image.Image
	if __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.UseRGBA {
		__obf_806b828cee04f014 := helper.CreateNRGBACanvas(__obf_9bab71876408a0b2.Size+10, __obf_9bab71876408a0b2.Size+10, true)
		draw.BiLinear.Scale(__obf_806b828cee04f014, __obf_806b828cee04f014.Bounds(), __obf_9bab71876408a0b2.Shape, __obf_9bab71876408a0b2.Shape.Bounds(), draw.Over, nil)
		__obf_036d52f5e09fd82a = __obf_806b828cee04f014.Bounds()
		__obf_07790da0efb07292 = __obf_806b828cee04f014
	} else {
		__obf_806b828cee04f014 := helper.CreatePaletteCanvas(__obf_9bab71876408a0b2.Size+10, __obf_9bab71876408a0b2.Size+10, __obf_3c7fe566d0291378)
		draw.BiLinear.Scale(__obf_806b828cee04f014, __obf_806b828cee04f014.Bounds(), __obf_9bab71876408a0b2.Shape, __obf_9bab71876408a0b2.Shape.Bounds(), draw.Over, nil)
		__obf_036d52f5e09fd82a = __obf_806b828cee04f014.Bounds()
		__obf_07790da0efb07292 = __obf_806b828cee04f014
	}

	draw.Draw(__obf_628ca96b9f843b5d, __obf_036d52f5e09fd82a, __obf_07790da0efb07292, image.Point{}, draw.Over)

	return __obf_628ca96b9f843b5d, nil
}

func (__obf_6fc320ffbfc77017 *Captcha) __obf_624ec29df8fa8452() (*CaptchaData, error) {

	var (
		__obf_54709c91224bea89, __obf_4b25c080de0589a1, __obf_ed6d947c65ad7eda []*types.Dot
		__obf_7986c1e4d6594623                                                 []image.Image
		__obf_f39bc0f22cbf62e8                                                 []image.Image
		__obf_e6925b4ca56caa93, __obf_131375f4a4b92e57                         image.Image
		__obf_48243aced6a1b5c5                                                 error
	)
	__obf_f39bc0f22cbf62e8, __obf_48243aced6a1b5c5 = assets.GetShapes()
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}
	__obf_e13651e48ac7b871 := util.RandInt(__obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.LenRange.Min, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.LenRange.Max)
	__obf_f39bc0f22cbf62e8 = util.PickN(__obf_f39bc0f22cbf62e8, __obf_e13651e48ac7b871)
	__obf_54709c91224bea89 = __obf_6fc320ffbfc77017.__obf_8c1071779654d808(__obf_f39bc0f22cbf62e8, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Height, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.SizeRange, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.DotPadding)
	__obf_225f64bd13d4f74d := rand.Perm(len(__obf_54709c91224bea89))
	__obf_979020f5df4d4bfd := util.RandInt(__obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.VerifyLenRange.Min, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.VerifyLenRange.Max)
	__obf_ed6d947c65ad7eda = make([]*types.Dot, __obf_979020f5df4d4bfd)
	for __obf_1f520fc8bd1d0fba, __obf_d8fd56811adf1973 := range __obf_225f64bd13d4f74d {
		if __obf_1f520fc8bd1d0fba >= __obf_979020f5df4d4bfd {
			break
		}
		__obf_9bab71876408a0b2 := __obf_54709c91224bea89[__obf_d8fd56811adf1973]
		__obf_9bab71876408a0b2.
			Index = __obf_1f520fc8bd1d0fba
		__obf_ed6d947c65ad7eda[__obf_1f520fc8bd1d0fba] = __obf_9bab71876408a0b2
		__obf_7986c1e4d6594623 = append(__obf_7986c1e4d6594623, __obf_ed6d947c65ad7eda[__obf_1f520fc8bd1d0fba].Shape)
	}
	__obf_4b25c080de0589a1 = __obf_6fc320ffbfc77017.__obf_8c1071779654d808(__obf_7986c1e4d6594623, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.SizeRange, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.DotPadding)
	__obf_e6925b4ca56caa93, __obf_48243aced6a1b5c5 = __obf_6fc320ffbfc77017.__obf_d8ec7d13971e4e3c(Shape, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Height, __obf_54709c91224bea89)
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}
	__obf_131375f4a4b92e57, __obf_48243aced6a1b5c5 = __obf_6fc320ffbfc77017.__obf_21f55b5dacea5e86(Shape, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Height, __obf_4b25c080de0589a1)
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}

	return &CaptchaData{__obf_f7494f398db65194: __obf_ed6d947c65ad7eda, __obf_e6925b4ca56caa93: types.NewJPEGImage(__obf_e6925b4ca56caa93), __obf_131375f4a4b92e57: types.NewPNGImage(__obf_131375f4a4b92e57)}, nil
}

func (__obf_6fc320ffbfc77017 *Captcha) __obf_b52fbb7a70540196() (*CaptchaData, error) {
	__obf_e13651e48ac7b871 := util.RandInt(__obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.LenRange.Min, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.LenRange.Max)
	__obf_6af6592ca4175d46 := util.PickN(assets.GetChineseChars(), __obf_e13651e48ac7b871)

	var __obf_f7494f398db65194, __obf_4b25c080de0589a1, __obf_ed6d947c65ad7eda []*types.Dot
	var __obf_de917dfc5c8da78e []string
	var __obf_e6925b4ca56caa93, __obf_131375f4a4b92e57 image.Image
	__obf_f7494f398db65194 = __obf_6fc320ffbfc77017.__obf_f1e3d0b53bb225e9(__obf_6af6592ca4175d46, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Height, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.SizeRange, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.DotPadding)
	__obf_225f64bd13d4f74d := rand.Perm(len(__obf_f7494f398db65194))
	__obf_979020f5df4d4bfd := util.RandInt(__obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.VerifyLenRange.Min, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.VerifyLenRange.Max)
	__obf_ed6d947c65ad7eda = make([]*types.Dot, __obf_979020f5df4d4bfd)
	for __obf_1f520fc8bd1d0fba, __obf_d8fd56811adf1973 := range __obf_225f64bd13d4f74d {
		if __obf_1f520fc8bd1d0fba >= __obf_979020f5df4d4bfd {
			break
		}
		__obf_9bab71876408a0b2 := __obf_f7494f398db65194[__obf_d8fd56811adf1973]
		__obf_9bab71876408a0b2.
			Index = __obf_1f520fc8bd1d0fba
		__obf_ed6d947c65ad7eda[__obf_1f520fc8bd1d0fba] = __obf_9bab71876408a0b2
		__obf_de917dfc5c8da78e = append(__obf_de917dfc5c8da78e, __obf_ed6d947c65ad7eda[__obf_1f520fc8bd1d0fba].Text)
	}
	__obf_4b25c080de0589a1 = __obf_6fc320ffbfc77017.__obf_f1e3d0b53bb225e9(__obf_de917dfc5c8da78e, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.Height, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.SizeRange, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.DotPadding)
	var __obf_48243aced6a1b5c5 error
	__obf_e6925b4ca56caa93, __obf_48243aced6a1b5c5 = __obf_6fc320ffbfc77017.__obf_d8ec7d13971e4e3c(Text, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Height, __obf_f7494f398db65194)
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}
	__obf_131375f4a4b92e57, __obf_48243aced6a1b5c5 = __obf_6fc320ffbfc77017.__obf_21f55b5dacea5e86(Text, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.Width, __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Secondary.Height, __obf_4b25c080de0589a1)
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}

	return &CaptchaData{__obf_f7494f398db65194: __obf_ed6d947c65ad7eda, __obf_e6925b4ca56caa93: types.NewJPEGImage(__obf_e6925b4ca56caa93), __obf_131375f4a4b92e57: types.NewPNGImage(__obf_131375f4a4b92e57)}, nil
}

func (__obf_6fc320ffbfc77017 *Captcha) __obf_8c1071779654d808(__obf_f39bc0f22cbf62e8 []image.Image, __obf_f246da57f4cb3862, __obf_9091d3f6b7268dcd int, __obf_72003c52c8f32232 types.Range, __obf_5489edadc0738797 int) []*types.Dot {
	__obf_e13651e48ac7b871 := len(__obf_f39bc0f22cbf62e8)
	var __obf_f7494f398db65194 = make([]*types.Dot, __obf_e13651e48ac7b871)
	// width := imageSize.Width
	// height := imageSize.Height
	if __obf_5489edadc0738797 > 0 {
		__obf_f246da57f4cb3862 -= __obf_5489edadc0738797
		__obf_9091d3f6b7268dcd -= __obf_5489edadc0738797
	}

	for __obf_1f520fc8bd1d0fba, __obf_014ab0047358450e := range __obf_f39bc0f22cbf62e8 {
		__obf_68f9e8b08ba6297f := __obf_6fc320ffbfc77017.__obf_68f9e8b08ba6297f()
		__obf_6a63c1f04b96ade4 := util.PickN(__obf_83c572abf9337148, 1)[0]
		__obf_4c2adbc4a6d9fe5f := util.PickN(__obf_69bfeeba36ce7750, 1)[0]
		__obf_ad013388b3318c79 := util.RandInt(__obf_72003c52c8f32232.Min, __obf_72003c52c8f32232.Max)
		__obf_cc9401c17a2e2f22 := 10
		__obf_4274d61c9b35bdc5 := __obf_f246da57f4cb3862 / __obf_e13651e48ac7b871
		__obf_046c8dd6cf293fe4 := math.Abs(float64(__obf_4274d61c9b35bdc5) - float64(__obf_ad013388b3318c79))
		__obf_34134814a7b1abfe := (__obf_1f520fc8bd1d0fba * __obf_4274d61c9b35bdc5) + util.RandInt(0, int(max(__obf_046c8dd6cf293fe4, 1)))
		__obf_e665f4bbf1b847f4 := util.RandInt(__obf_cc9401c17a2e2f22, __obf_9091d3f6b7268dcd+__obf_ad013388b3318c79)
		__obf_46be962c13268f5a := int(min(max(float64(__obf_34134814a7b1abfe), float64(__obf_cc9401c17a2e2f22)), float64(__obf_f246da57f4cb3862-__obf_cc9401c17a2e2f22-(__obf_5489edadc0738797*2))))
		__obf_e7f217bb23bb3389 := int(min(max(float64(__obf_e665f4bbf1b847f4), float64(__obf_ad013388b3318c79+__obf_cc9401c17a2e2f22)), float64(__obf_9091d3f6b7268dcd+(__obf_ad013388b3318c79/2)-(__obf_5489edadc0738797*2))))
		__obf_f7494f398db65194[__obf_1f520fc8bd1d0fba] = &types.Dot{
			Index:          __obf_1f520fc8bd1d0fba,
			X:              __obf_46be962c13268f5a,
			Y:              __obf_e7f217bb23bb3389 - __obf_ad013388b3318c79,
			FontSize:       __obf_ad013388b3318c79,
			Size:           __obf_ad013388b3318c79,
			Angle:          __obf_68f9e8b08ba6297f,
			PrimaryColor:   __obf_6a63c1f04b96ade4,
			SecondaryColor: __obf_4c2adbc4a6d9fe5f,
			Shape:          __obf_014ab0047358450e,
		}
	}

	return __obf_f7494f398db65194
}

func (__obf_6fc320ffbfc77017 *Captcha) __obf_f1e3d0b53bb225e9(__obf_6af6592ca4175d46 []string, __obf_f246da57f4cb3862, __obf_9091d3f6b7268dcd int, __obf_3b13434c5fd4293b types.Range, __obf_5489edadc0738797 int) []*types.Dot {
	__obf_e13651e48ac7b871 := len(__obf_6af6592ca4175d46)
	var __obf_f7494f398db65194 = make([]*types.Dot, __obf_e13651e48ac7b871)
	if __obf_5489edadc0738797 > 0 {
		__obf_f246da57f4cb3862 -= __obf_5489edadc0738797
		__obf_9091d3f6b7268dcd -= __obf_5489edadc0738797
	}

	for __obf_1f520fc8bd1d0fba, __obf_014ab0047358450e := range __obf_6af6592ca4175d46 {
		__obf_68f9e8b08ba6297f := __obf_6fc320ffbfc77017.__obf_68f9e8b08ba6297f()
		__obf_6a63c1f04b96ade4 := util.PickN(__obf_83c572abf9337148, 1)[0]
		__obf_4c2adbc4a6d9fe5f := util.PickN(__obf_69bfeeba36ce7750, 1)[0]
		__obf_ad013388b3318c79 := util.RandInt(__obf_3b13434c5fd4293b.Min, __obf_3b13434c5fd4293b.Max)
		__obf_b3af34e64e8775e9 := __obf_ad013388b3318c79
		__obf_10e905d3d91870a9 := __obf_ad013388b3318c79

		if util.LenChineseChar(__obf_014ab0047358450e) > 1 {
			__obf_10e905d3d91870a9 = __obf_ad013388b3318c79 * util.LenChineseChar(__obf_014ab0047358450e)

			if __obf_68f9e8b08ba6297f > 0 {
				__obf_893e3508f2f9208a := __obf_10e905d3d91870a9 - __obf_ad013388b3318c79
				__obf_07e457bc9e2c5b93 := __obf_68f9e8b08ba6297f % 90
				__obf_384d39546fa8b21e := float64(__obf_893e3508f2f9208a) / 90
				__obf_a52cd80f96bb60c8 := max(float64(__obf_07e457bc9e2c5b93)*__obf_384d39546fa8b21e, 1)
				__obf_b3af34e64e8775e9 = __obf_b3af34e64e8775e9 + int(__obf_a52cd80f96bb60c8)
				__obf_10e905d3d91870a9 = __obf_10e905d3d91870a9 + int(__obf_a52cd80f96bb60c8)
			}
		}
		__obf_cc9401c17a2e2f22 := 10
		__obf_4274d61c9b35bdc5 := __obf_f246da57f4cb3862 / __obf_e13651e48ac7b871
		__obf_046c8dd6cf293fe4 := math.Abs(float64(__obf_4274d61c9b35bdc5) - float64(__obf_10e905d3d91870a9))
		__obf_34134814a7b1abfe := (__obf_1f520fc8bd1d0fba * __obf_4274d61c9b35bdc5) + util.RandInt(0, int(max(__obf_046c8dd6cf293fe4, 1)))
		__obf_e665f4bbf1b847f4 := util.RandInt(__obf_cc9401c17a2e2f22, __obf_9091d3f6b7268dcd+__obf_b3af34e64e8775e9)
		__obf_46be962c13268f5a := int(min(max(float64(__obf_34134814a7b1abfe), float64(__obf_cc9401c17a2e2f22)), float64(__obf_f246da57f4cb3862-__obf_cc9401c17a2e2f22-(__obf_5489edadc0738797*2))))
		__obf_e7f217bb23bb3389 := int(min(max(float64(__obf_e665f4bbf1b847f4), float64(__obf_b3af34e64e8775e9+__obf_cc9401c17a2e2f22)), float64(__obf_9091d3f6b7268dcd+(__obf_b3af34e64e8775e9/2)-(__obf_5489edadc0738797*2))))
		__obf_f7494f398db65194[__obf_1f520fc8bd1d0fba] = &types.Dot{
			Index:          __obf_1f520fc8bd1d0fba,
			X:              __obf_46be962c13268f5a,
			Y:              __obf_e7f217bb23bb3389 - __obf_b3af34e64e8775e9,
			FontSize:       __obf_ad013388b3318c79,
			Size:           max(__obf_10e905d3d91870a9, __obf_b3af34e64e8775e9),
			Angle:          __obf_68f9e8b08ba6297f,
			PrimaryColor:   __obf_6a63c1f04b96ade4,
			SecondaryColor: __obf_4c2adbc4a6d9fe5f,
			Text:           __obf_014ab0047358450e,
		}
	}

	return __obf_f7494f398db65194
}

func (__obf_6fc320ffbfc77017 *Captcha) __obf_d8ec7d13971e4e3c(__obf_c112f65c651fcf04 DotType, __obf_f246da57f4cb3862, __obf_9091d3f6b7268dcd int, __obf_f7494f398db65194 []*types.Dot) (image.Image, error) {
	__obf_84ded1ae7847cb47 := &DrawImageParams{
		Width:       __obf_f246da57f4cb3862,
		Height:      __obf_9091d3f6b7268dcd,
		Alpha:       __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.Alpha,
		FontHinting: __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.FontHinting,
		Dots:        __obf_f7494f398db65194,
		ShowShadow:  __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.ShowShadow,
		ShadowPoint: __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.ShadowPoint,
	}
	var __obf_48243aced6a1b5c5 error
	__obf_84ded1ae7847cb47.
		Background, __obf_48243aced6a1b5c5 = assets.PickBgImage()
	if __obf_48243aced6a1b5c5 != nil {
		return nil, __obf_48243aced6a1b5c5
	}

	return __obf_6fc320ffbfc77017.DrawWithNRGBA(__obf_c112f65c651fcf04, __obf_84ded1ae7847cb47)
}

func (__obf_0f1aca653a04363a *Captcha) __obf_21f55b5dacea5e86(__obf_c112f65c651fcf04 DotType, __obf_f246da57f4cb3862, __obf_9091d3f6b7268dcd int, __obf_f7494f398db65194 []*types.Dot) (image.Image, error) {
	var (
		__obf_e27312faaf82e65a =
		// err      error
		make([]*types.Dot, 0, len(__obf_f7494f398db65194))
	)
	__obf_61978efdff4ff941 := __obf_f246da57f4cb3862 / len(__obf_f7494f398db65194)

	for __obf_1f520fc8bd1d0fba, __obf_014ab0047358450e := range __obf_f7494f398db65194 {
		__obf_e13651e48ac7b871 := 1
		if __obf_c112f65c651fcf04 == Text {
			__obf_e13651e48ac7b871 = len(__obf_014ab0047358450e.Text)
		}
		__obf_014ab0047358450e.
			X = int(max(float64(__obf_61978efdff4ff941*__obf_1f520fc8bd1d0fba+__obf_61978efdff4ff941/__obf_014ab0047358450e.Size), 8))
		__obf_b986e8831ab9762a := max(1, __obf_9091d3f6b7268dcd/16*__obf_e13651e48ac7b871)
		__obf_014ab0047358450e.
			Y = __obf_9091d3f6b7268dcd/2 + __obf_014ab0047358450e.FontSize/2 - rand.IntN(__obf_b986e8831ab9762a)
		__obf_e27312faaf82e65a = append(__obf_e27312faaf82e65a, __obf_014ab0047358450e)
	}
	__obf_ece4016def2d6db0 := &DrawImageParams{
		Width:             __obf_f246da57f4cb3862,
		Height:            __obf_9091d3f6b7268dcd,
		Dots:              __obf_e27312faaf82e65a,
		BackgroundDistort: __obf_0f1aca653a04363a.__obf_5ecf3765e23ebeff(__obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.Secondary.BgDistort),
		BgCircles:         __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.Secondary.BgCircles,
		BgSlimLines:       __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_aaf0b708fdf2a15c []color.Color
	for _, __obf_a098b356d94dc03f := range __obf_69bfeeba36ce7750 {
		__obf_10213221b9f3427b, _ := util.ParseHexColor(__obf_a098b356d94dc03f)
		__obf_aaf0b708fdf2a15c = append(__obf_aaf0b708fdf2a15c, __obf_10213221b9f3427b)
	}

	var __obf_dbca0f5a47406936 []color.Color
	for _, __obf_10213221b9f3427b := range __obf_69bfeeba36ce7750 {
		__obf_7778291a332582b1, _ := util.ParseHexColor(__obf_10213221b9f3427b)
		__obf_dbca0f5a47406936 = append(__obf_dbca0f5a47406936, __obf_7778291a332582b1)
	}

	if __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.UseRGBA || __obf_0f1aca653a04363a.__obf_3d43bbdf2fa5cdea.Secondary.NonDeformAble {
		return __obf_0f1aca653a04363a.DrawWithNRGBA2(__obf_c112f65c651fcf04, __obf_ece4016def2d6db0, __obf_aaf0b708fdf2a15c, __obf_dbca0f5a47406936)
	}
	return __obf_0f1aca653a04363a.DrawWithPalette(__obf_c112f65c651fcf04, __obf_ece4016def2d6db0, __obf_aaf0b708fdf2a15c, __obf_dbca0f5a47406936)
}

func (__obf_6fc320ffbfc77017 *Captcha) __obf_5ecf3765e23ebeff(__obf_200ebdf601d7c5a0 int) int {
	switch __obf_200ebdf601d7c5a0 {
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

func (__obf_6fc320ffbfc77017 *Captcha) __obf_68f9e8b08ba6297f() int {
	__obf_ca9ed8c095b54bd6 := __obf_6fc320ffbfc77017.__obf_3d43bbdf2fa5cdea.Primary.AnglePosRange
	__obf_f5871bae1215aa4e := rand.IntN(len(__obf_ca9ed8c095b54bd6))
	if __obf_f5871bae1215aa4e < 0 {
		return 0
	}
	__obf_943f35295a72bc5f := __obf_ca9ed8c095b54bd6[__obf_f5871bae1215aa4e]
	__obf_87d2f2f8702e09cb := util.RandInt(__obf_943f35295a72bc5f.Min, __obf_943f35295a72bc5f.Max)

	return __obf_87d2f2f8702e09cb
}
