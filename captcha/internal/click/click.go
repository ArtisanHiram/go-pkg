package __obf_a0c8110fd085454e

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
	__obf_f8f07b39b37290ac *Options
}

var _ types.Captcha = (*Captcha)(nil)

func NewCaptcha(__obf_f8f07b39b37290ac *Options) *Captcha {
	if __obf_f8f07b39b37290ac == nil {
		return nil
	}
	return &Captcha{__obf_f8f07b39b37290ac}
}

func (__obf_761b3f820674d2e3 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_ac0f57b9f94f3dbe := __obf_014f711943baebe8(__obf_761b3f820674d2e3.__obf_f8f07b39b37290ac)
	if __obf_ac0f57b9f94f3dbe == Shape {
		return __obf_761b3f820674d2e3.__obf_a3d185177a709dd1()
	}

	return __obf_761b3f820674d2e3.__obf_3a74189fe74ae6c9()
}

func __obf_014f711943baebe8(__obf_f8f07b39b37290ac *Options) DotType {
	__obf_72e889006391b3bf := DotType(rand.IntN(2))
	if __obf_72e889006391b3bf == Shape {
		__obf_f8f07b39b37290ac.
			UseRGBA = false
		__obf_f8f07b39b37290ac.
			Secondary.BgDistort = types.DistortLevel1
		__obf_f8f07b39b37290ac.
			Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_f8f07b39b37290ac.
			Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_f8f07b39b37290ac.
			UseRGBA = true
		__obf_f8f07b39b37290ac.
			Secondary.BgDistort = types.DistortLevel4
		__obf_f8f07b39b37290ac.
			Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_f8f07b39b37290ac.
			Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_72e889006391b3bf
}

func (__obf_3d944245c503e38e *Captcha) DrawWithNRGBA(__obf_ac0f57b9f94f3dbe DotType, __obf_d1d4e3f72f8cd606 *DrawImageParams) (image.Image, error) {
	__obf_cea1e242f930c716 := __obf_d1d4e3f72f8cd606.Dots
	__obf_5e90e389aeb91533 := helper.CreateNRGBACanvas(__obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height, true)

	for _, __obf_31d51e7dec1e6919 := range __obf_cea1e242f930c716 {
		__obf_fb4bf526fd5a3942, __obf_966b5bdceb42309f, __obf_43790225c3b3432f := __obf_3d944245c503e38e.DrawDotImage(__obf_ac0f57b9f94f3dbe, __obf_31d51e7dec1e6919, __obf_d1d4e3f72f8cd606)
		if __obf_43790225c3b3432f != nil {
			return nil, __obf_43790225c3b3432f
		}
		__obf_65e3654278303959 := __obf_966b5bdceb42309f.MinX
		__obf_e4340f973964335d := __obf_966b5bdceb42309f.MaxX
		__obf_139754cbaa70e25e := __obf_966b5bdceb42309f.MinY
		__obf_6cd98c2ef722c4e1 := __obf_966b5bdceb42309f.MaxY
		__obf_1181c8127158d833 := __obf_e4340f973964335d - __obf_65e3654278303959
		__obf_62c56517f1f23b8f := __obf_6cd98c2ef722c4e1 - __obf_139754cbaa70e25e

		draw.Draw(__obf_5e90e389aeb91533, image.Rect(__obf_31d51e7dec1e6919.X, __obf_31d51e7dec1e6919.Y, __obf_31d51e7dec1e6919.X+__obf_1181c8127158d833, __obf_31d51e7dec1e6919.Y+__obf_62c56517f1f23b8f), __obf_fb4bf526fd5a3942, image.Point{X: __obf_65e3654278303959, Y: __obf_139754cbaa70e25e}, draw.Over)

	}
	__obf_e3bc7b3b39cb093b := __obf_d1d4e3f72f8cd606.Background
	__obf_a7f7bd1f608f58bf := __obf_5e90e389aeb91533.Bounds()
	__obf_84e592da746d7c09 := helper.CreateNRGBACanvas(__obf_a7f7bd1f608f58bf.Dx(), __obf_a7f7bd1f608f58bf.Dx(), true)
	__obf_19f97e8a0e1b18b3 := util.RangCutImagePos(__obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height, __obf_e3bc7b3b39cb093b)
	draw.Draw(__obf_84e592da746d7c09, __obf_a7f7bd1f608f58bf, __obf_e3bc7b3b39cb093b, __obf_19f97e8a0e1b18b3, draw.Src)
	draw.Draw(__obf_84e592da746d7c09, __obf_5e90e389aeb91533.Bounds(), __obf_5e90e389aeb91533, image.Point{}, draw.Over)
	__obf_84e592da746d7c09.
		SubImage(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height))
	return __obf_84e592da746d7c09, nil
}

func (__obf_3d944245c503e38e *Captcha) DrawWithPalette(__obf_ac0f57b9f94f3dbe DotType, __obf_d1d4e3f72f8cd606 *DrawImageParams, __obf_91b15260b589b417 []color.Color, __obf_c66d486b139ce9bc []color.Color) (image.Image, error) {
	__obf_cea1e242f930c716 := __obf_d1d4e3f72f8cd606.Dots
	__obf_b26e92f04896321e := make([]color.Color, 0, len(__obf_c66d486b139ce9bc))
	for _, __obf_6f0d77b9defe3274 := range __obf_c66d486b139ce9bc {
		__obf_ef5914876e5fc840, __obf_c912c01bd881536e, __obf_a7f7bd1f608f58bf, _ := __obf_6f0d77b9defe3274.RGBA()
		__obf_118a1df592364653 := util.FormatAlpha(__obf_d1d4e3f72f8cd606.ThumbDisturbAlpha)
		__obf_b26e92f04896321e = append(__obf_b26e92f04896321e, color.RGBA{R: uint8(__obf_ef5914876e5fc840), G: uint8(__obf_c912c01bd881536e), B: uint8(__obf_a7f7bd1f608f58bf), A: __obf_118a1df592364653})
	}

	var __obf_645a3e8fe26a5ce4 = make([]color.Color, 0, len(__obf_91b15260b589b417)+len(__obf_b26e92f04896321e)+1)
	__obf_645a3e8fe26a5ce4 = append(__obf_645a3e8fe26a5ce4, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_645a3e8fe26a5ce4 = append(__obf_645a3e8fe26a5ce4, __obf_91b15260b589b417...)
	__obf_645a3e8fe26a5ce4 = append(__obf_645a3e8fe26a5ce4, __obf_b26e92f04896321e...)
	__obf_5e90e389aeb91533 := types.NewPalette(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height), __obf_645a3e8fe26a5ce4)
	if __obf_d1d4e3f72f8cd606.BgCircles > 0 {
		__obf_3d944245c503e38e.__obf_39c36569bc24721c(__obf_5e90e389aeb91533, __obf_d1d4e3f72f8cd606.BgCircles, 1, __obf_b26e92f04896321e)
	}
	if __obf_d1d4e3f72f8cd606.BgSlimLines > 0 {
		__obf_3d944245c503e38e.__obf_9941afd1d2367d55(__obf_5e90e389aeb91533, __obf_d1d4e3f72f8cd606.BgSlimLines, __obf_b26e92f04896321e)
	}

	for _, __obf_31d51e7dec1e6919 := range __obf_cea1e242f930c716 {
		__obf_b8a694fd5c1c8919, _ := util.ParseHexColor(__obf_31d51e7dec1e6919.PrimaryColor)
		var __obf_43790225c3b3432f error
		if __obf_ac0f57b9f94f3dbe == Shape {
			var __obf_fb4bf526fd5a3942 *types.NRGBA
			__obf_fb4bf526fd5a3942, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawShapeImage(__obf_31d51e7dec1e6919, __obf_b8a694fd5c1c8919)
			if __obf_43790225c3b3432f != nil {
				return nil, __obf_43790225c3b3432f
			}
			__obf_fb4bf526fd5a3942.
				Rotate(__obf_31d51e7dec1e6919.Angle, false)
			__obf_b4619e5a13562e3a := __obf_5e90e389aeb91533.Bounds()
			__obf_cc56fef1cbc2ac92 := __obf_fb4bf526fd5a3942.Bounds()
			__obf_225bef0fe9ec0b5b := image.Point{X: __obf_b4619e5a13562e3a.Dx() - __obf_cc56fef1cbc2ac92.Dx(), Y: __obf_b4619e5a13562e3a.Dy() - __obf_cc56fef1cbc2ac92.Dy()}
			draw.Draw(__obf_5e90e389aeb91533, image.Rect(__obf_31d51e7dec1e6919.X, __obf_225bef0fe9ec0b5b.Y, __obf_225bef0fe9ec0b5b.X+__obf_cc56fef1cbc2ac92.Dx(), __obf_225bef0fe9ec0b5b.Y+__obf_cc56fef1cbc2ac92.Dy()), __obf_fb4bf526fd5a3942, image.Point{}, draw.Over)
		} else {
			__obf_f198f5dc9244d650 := fixed.P(__obf_31d51e7dec1e6919.X, __obf_31d51e7dec1e6919.Y)
			__obf_d1d4e3f72f8cd606 := &types.DrawStringParam{
				Color:       __obf_b8a694fd5c1c8919,
				FontSize:    __obf_31d51e7dec1e6919.FontSize,
				Size:        __obf_31d51e7dec1e6919.Size,
				FontDPI:     __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.FontDPI,
				FontHinting: __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.FontHinting,
				Text:        __obf_31d51e7dec1e6919.Text,
			}
			__obf_d1d4e3f72f8cd606.
				Font, __obf_43790225c3b3432f = assets.PickFont()
			if __obf_43790225c3b3432f != nil {
				return nil, __obf_43790225c3b3432f
			}
			__obf_43790225c3b3432f = helper.DrawString(__obf_5e90e389aeb91533, __obf_d1d4e3f72f8cd606, __obf_f198f5dc9244d650)
		}

		if __obf_43790225c3b3432f != nil {
			return __obf_5e90e389aeb91533, __obf_43790225c3b3432f
		}
	}

	if __obf_d1d4e3f72f8cd606.Background != nil {
		__obf_e3bc7b3b39cb093b := __obf_d1d4e3f72f8cd606.Background
		__obf_a7f7bd1f608f58bf := __obf_e3bc7b3b39cb093b.Bounds()
		__obf_84e592da746d7c09 := helper.CreateNRGBACanvas(__obf_a7f7bd1f608f58bf.Dx(), __obf_a7f7bd1f608f58bf.Dy(), true)
		__obf_19f97e8a0e1b18b3 := util.RangCutImagePos(__obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height, __obf_e3bc7b3b39cb093b)
		draw.Draw(__obf_84e592da746d7c09, __obf_a7f7bd1f608f58bf, __obf_e3bc7b3b39cb093b, __obf_19f97e8a0e1b18b3, draw.Src)
		__obf_5e90e389aeb91533.
			Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_84e592da746d7c09, __obf_5e90e389aeb91533.Bounds(), __obf_5e90e389aeb91533, image.Point{}, draw.Over)
		__obf_e7aaf7358dd58db4 := __obf_84e592da746d7c09.SubImage(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height)).(*image.NRGBA)
		return __obf_e7aaf7358dd58db4, nil
	}

	if __obf_d1d4e3f72f8cd606.BackgroundDistort > 0 {
		__obf_5e90e389aeb91533.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_d1d4e3f72f8cd606.BackgroundDistort))
	}

	return __obf_5e90e389aeb91533, nil

}

func (__obf_3d944245c503e38e *Captcha) DrawWithNRGBA2(__obf_ac0f57b9f94f3dbe DotType, __obf_d1d4e3f72f8cd606 *DrawImageParams, __obf_91b15260b589b417 []color.Color, __obf_c66d486b139ce9bc []color.Color) (image.Image, error) {
	__obf_cea1e242f930c716 := __obf_d1d4e3f72f8cd606.Dots
	__obf_b26e92f04896321e := make([]color.Color, 0, len(__obf_c66d486b139ce9bc))
	for _, __obf_6f0d77b9defe3274 := range __obf_c66d486b139ce9bc {
		__obf_ef5914876e5fc840, __obf_c912c01bd881536e, __obf_a7f7bd1f608f58bf, _ := __obf_6f0d77b9defe3274.RGBA()
		__obf_118a1df592364653 := util.FormatAlpha(__obf_d1d4e3f72f8cd606.ThumbDisturbAlpha)
		__obf_b26e92f04896321e = append(__obf_b26e92f04896321e, color.RGBA{R: uint8(__obf_ef5914876e5fc840), G: uint8(__obf_c912c01bd881536e), B: uint8(__obf_a7f7bd1f608f58bf), A: __obf_118a1df592364653})
	}

	var __obf_645a3e8fe26a5ce4 = make([]color.Color, 0, len(__obf_91b15260b589b417)+len(__obf_b26e92f04896321e)+1)
	__obf_645a3e8fe26a5ce4 = append(__obf_645a3e8fe26a5ce4, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_645a3e8fe26a5ce4 = append(__obf_645a3e8fe26a5ce4, __obf_91b15260b589b417...)
	__obf_645a3e8fe26a5ce4 = append(__obf_645a3e8fe26a5ce4, __obf_b26e92f04896321e...)
	__obf_f93a97fa9672a24c := types.NewNRGBA(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height), true)
	if __obf_d1d4e3f72f8cd606.Background != nil {
		__obf_e3bc7b3b39cb093b := __obf_d1d4e3f72f8cd606.Background
		__obf_a7f7bd1f608f58bf := __obf_e3bc7b3b39cb093b.Bounds()
		__obf_84e592da746d7c09 := helper.CreateNRGBACanvas(__obf_a7f7bd1f608f58bf.Dx(), __obf_a7f7bd1f608f58bf.Dy(), true)
		__obf_19f97e8a0e1b18b3 := util.RangCutImagePos(__obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height, __obf_e3bc7b3b39cb093b)
		draw.Draw(__obf_84e592da746d7c09, __obf_a7f7bd1f608f58bf, __obf_e3bc7b3b39cb093b, __obf_19f97e8a0e1b18b3, draw.Src)
		__obf_e7aaf7358dd58db4 := __obf_84e592da746d7c09.SubImage(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height)).(*image.NRGBA)
		draw.Draw(__obf_f93a97fa9672a24c, __obf_e7aaf7358dd58db4.Bounds(), __obf_e7aaf7358dd58db4, image.Point{}, draw.Over)
	}
	__obf_5e90e389aeb91533 := types.NewPalette(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height), __obf_645a3e8fe26a5ce4)
	if __obf_d1d4e3f72f8cd606.BgCircles > 0 {
		__obf_3d944245c503e38e.__obf_39c36569bc24721c(__obf_5e90e389aeb91533, __obf_d1d4e3f72f8cd606.BgCircles, 1, __obf_b26e92f04896321e)
	}
	if __obf_d1d4e3f72f8cd606.BgSlimLines > 0 {
		__obf_3d944245c503e38e.__obf_9941afd1d2367d55(__obf_5e90e389aeb91533, __obf_d1d4e3f72f8cd606.BgSlimLines, __obf_b26e92f04896321e)
	}
	if __obf_d1d4e3f72f8cd606.BackgroundDistort > 0 {
		__obf_5e90e389aeb91533.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_d1d4e3f72f8cd606.BackgroundDistort))
	}
	__obf_d27774be9b372382 := __obf_5e90e389aeb91533.Bounds()
	__obf_1181c8127158d833 := __obf_d27774be9b372382.Dx() / len(__obf_cea1e242f930c716)
	for __obf_d9fed1663089e979, __obf_31d51e7dec1e6919 := range __obf_cea1e242f930c716 {
		__obf_b8a694fd5c1c8919, _ := util.ParseHexColor(__obf_31d51e7dec1e6919.PrimaryColor)
		var __obf_43790225c3b3432f error
		if __obf_ac0f57b9f94f3dbe == Shape {
			var __obf_fb4bf526fd5a3942 *types.NRGBA
			__obf_fb4bf526fd5a3942, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawShapeImage(__obf_31d51e7dec1e6919, __obf_b8a694fd5c1c8919)
			if __obf_43790225c3b3432f != nil {
				return nil, __obf_43790225c3b3432f
			}
			__obf_fb4bf526fd5a3942.
				Rotate(__obf_31d51e7dec1e6919.Angle, false)
			__obf_b4619e5a13562e3a := __obf_f93a97fa9672a24c.Bounds()
			__obf_cc56fef1cbc2ac92 := __obf_fb4bf526fd5a3942.Bounds()
			__obf_225bef0fe9ec0b5b := image.Point{X: __obf_b4619e5a13562e3a.Dx() - __obf_cc56fef1cbc2ac92.Dx(), Y: __obf_b4619e5a13562e3a.Dy() - __obf_cc56fef1cbc2ac92.Dy()}
			draw.Draw(__obf_f93a97fa9672a24c, image.Rect(__obf_31d51e7dec1e6919.X, __obf_225bef0fe9ec0b5b.Y, __obf_225bef0fe9ec0b5b.X+__obf_cc56fef1cbc2ac92.Dx(), __obf_225bef0fe9ec0b5b.Y+__obf_cc56fef1cbc2ac92.Dy()), __obf_fb4bf526fd5a3942, image.Point{}, draw.Over)
		} else {
			var __obf_b63d9efda3ea5cd7 *types.NRGBA
			__obf_b63d9efda3ea5cd7, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawStringImage(__obf_31d51e7dec1e6919, __obf_b8a694fd5c1c8919)
			if __obf_43790225c3b3432f != nil {
				return nil, __obf_43790225c3b3432f
			}
			__obf_b63d9efda3ea5cd7.
				Rotate(__obf_31d51e7dec1e6919.Angle, false)
			__obf_966b5bdceb42309f := __obf_b63d9efda3ea5cd7.CalcMarginBlankArea()
			__obf_65e3654278303959 := __obf_966b5bdceb42309f.MinX
			__obf_e4340f973964335d := __obf_966b5bdceb42309f.MaxX
			__obf_139754cbaa70e25e := __obf_966b5bdceb42309f.MinY
			__obf_6cd98c2ef722c4e1 := __obf_966b5bdceb42309f.MaxY
			__obf_b63d9efda3ea5cd7.
				SubImage(image.Rect(__obf_65e3654278303959, __obf_139754cbaa70e25e, __obf_e4340f973964335d, __obf_6cd98c2ef722c4e1))
			__obf_897cea6b749cef85 := __obf_b63d9efda3ea5cd7.Bounds()
			__obf_9d30264ecb5561cd := int(max(float64(__obf_1181c8127158d833*__obf_d9fed1663089e979+__obf_1181c8127158d833/__obf_897cea6b749cef85.Dx()), 8))
			__obf_701b7b1349750044 := util.RandInt(1, __obf_d27774be9b372382.Dy()-__obf_897cea6b749cef85.Dy()-4)

			draw.Draw(__obf_f93a97fa9672a24c, image.Rect(__obf_9d30264ecb5561cd, __obf_701b7b1349750044, __obf_9d30264ecb5561cd+__obf_897cea6b749cef85.Dx(), __obf_701b7b1349750044+__obf_897cea6b749cef85.Dy()), __obf_b63d9efda3ea5cd7, image.Point{X: __obf_897cea6b749cef85.Min.X, Y: __obf_897cea6b749cef85.Min.Y}, draw.Over)
		}

		if __obf_43790225c3b3432f != nil {
			return __obf_f93a97fa9672a24c, __obf_43790225c3b3432f
		}
	}
	__obf_9530b4ea7904a975 := types.NewNRGBA(image.Rect(0, 0, __obf_d1d4e3f72f8cd606.Width, __obf_d1d4e3f72f8cd606.Height), true)
	draw.Draw(__obf_9530b4ea7904a975, __obf_5e90e389aeb91533.Bounds(), __obf_5e90e389aeb91533, image.Point{}, draw.Over)
	draw.Draw(__obf_f93a97fa9672a24c, __obf_9530b4ea7904a975.Bounds(), __obf_9530b4ea7904a975, image.Point{}, draw.Over)
	return __obf_f93a97fa9672a24c, nil
}

func (__obf_3d944245c503e38e *Captcha) __obf_39c36569bc24721c(__obf_84e592da746d7c09 *types.Palette, __obf_89590d2dc9eab630, __obf_d87ba36fcde00384 int, __obf_f9117be9f2b81ec9 []color.Color) {
	__obf_e43ae88db9f80ce7 := __obf_84e592da746d7c09.Bounds().Max.X
	__obf_e5f38270079b0292 := __obf_84e592da746d7c09.Bounds().Max.Y
	for range __obf_89590d2dc9eab630 {
		__obf_5aecf180c5bc1e97 := util.RandColor(__obf_f9117be9f2b81ec9)
		__obf_ef5914876e5fc840 := //co.A = uint8(0xee)
			util.RandInt(1, __obf_d87ba36fcde00384)
		__obf_84e592da746d7c09.
			DrawCircle(util.RandInt(__obf_ef5914876e5fc840, __obf_e43ae88db9f80ce7-__obf_ef5914876e5fc840), util.RandInt(__obf_ef5914876e5fc840, __obf_e5f38270079b0292-__obf_ef5914876e5fc840), __obf_ef5914876e5fc840, __obf_5aecf180c5bc1e97)
	}
}

func (__obf_3d944245c503e38e *Captcha) __obf_9941afd1d2367d55(__obf_84e592da746d7c09 *types.Palette, __obf_864813e13cbd37a7 int, __obf_f9117be9f2b81ec9 []color.Color) {
	__obf_69206f82702a003b := __obf_84e592da746d7c09.Bounds().Max.X / 10
	__obf_9257b2dc4713d823 := __obf_69206f82702a003b * 9
	__obf_6f3efd4586e6dcf4 := __obf_84e592da746d7c09.Bounds().Max.Y / 3
	for __obf_d9fed1663089e979 := range __obf_864813e13cbd37a7 {
		__obf_b8829da13bb68821 := image.Point{X: rand.IntN(__obf_69206f82702a003b), Y: rand.IntN(__obf_6f3efd4586e6dcf4)}
		__obf_45c0df5312e5200e := image.Point{X: rand.IntN(__obf_69206f82702a003b) + __obf_9257b2dc4713d823, Y: rand.IntN(__obf_6f3efd4586e6dcf4)}

		if __obf_d9fed1663089e979%2 == 0 {
			__obf_b8829da13bb68821.
				Y = rand.IntN(__obf_6f3efd4586e6dcf4) + __obf_6f3efd4586e6dcf4*2
			__obf_45c0df5312e5200e.
				Y = rand.IntN(__obf_6f3efd4586e6dcf4)
		} else {
			__obf_b8829da13bb68821.
				Y = rand.IntN(__obf_6f3efd4586e6dcf4) + __obf_6f3efd4586e6dcf4*(__obf_d9fed1663089e979%2)
			__obf_45c0df5312e5200e.
				Y = rand.IntN(__obf_6f3efd4586e6dcf4) + __obf_6f3efd4586e6dcf4*2
		}
		__obf_5aecf180c5bc1e97 := util.RandColor(__obf_f9117be9f2b81ec9)
		__obf_84e592da746d7c09.
			//co.A = uint8(0xee)
			DrawBeeline(__obf_b8829da13bb68821, __obf_45c0df5312e5200e, __obf_5aecf180c5bc1e97)
	}
}

func (__obf_3d944245c503e38e *Captcha) DrawDotImage(__obf_ac0f57b9f94f3dbe DotType, __obf_b4922ee002e08e21 *types.Dot, __obf_d1d4e3f72f8cd606 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_b8a694fd5c1c8919, _ := util.ParseHexColor(__obf_b4922ee002e08e21.PrimaryColor)
	__obf_b8a694fd5c1c8919.
		A = util.FormatAlpha(__obf_d1d4e3f72f8cd606.Alpha)

	var __obf_b63d9efda3ea5cd7 image.Image
	var __obf_43790225c3b3432f error
	if __obf_ac0f57b9f94f3dbe == Shape {
		__obf_b63d9efda3ea5cd7, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawShapeImage(__obf_b4922ee002e08e21, __obf_b8a694fd5c1c8919)
	} else {
		__obf_b63d9efda3ea5cd7, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawStringImage(__obf_b4922ee002e08e21, __obf_b8a694fd5c1c8919)
	}
	if __obf_43790225c3b3432f != nil {
		return nil, nil, __obf_43790225c3b3432f
	}
	__obf_c8dbc8e646036d2e, _ := util.ParseHexColor(__obf_34f2d4e86729b26c)
	__obf_5e90e389aeb91533 := helper.CreateNRGBACanvas(__obf_b4922ee002e08e21.Size+10, __obf_b4922ee002e08e21.Size+10, true)
	if __obf_d1d4e3f72f8cd606.ShowShadow {
		var __obf_05f5bf1e737ecb5d *types.NRGBA
		if __obf_ac0f57b9f94f3dbe == Shape {
			__obf_05f5bf1e737ecb5d, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawShapeImage(__obf_b4922ee002e08e21, __obf_c8dbc8e646036d2e)
		} else {
			__obf_05f5bf1e737ecb5d, __obf_43790225c3b3432f = __obf_3d944245c503e38e.DrawStringImage(__obf_b4922ee002e08e21, __obf_c8dbc8e646036d2e)
		}
		if __obf_43790225c3b3432f != nil {
			return nil, nil, __obf_43790225c3b3432f
		}
		__obf_ac9a634221b8f7fd := __obf_d1d4e3f72f8cd606.ShadowPoint.X
		__obf_20e40d36ba084565 := __obf_d1d4e3f72f8cd606.ShadowPoint.Y
		draw.Draw(__obf_5e90e389aeb91533, __obf_05f5bf1e737ecb5d.Bounds(), __obf_05f5bf1e737ecb5d, image.Point{X: __obf_ac9a634221b8f7fd, Y: __obf_20e40d36ba084565}, draw.Over)
	}
	draw.Draw(__obf_5e90e389aeb91533, __obf_b63d9efda3ea5cd7.Bounds(), __obf_b63d9efda3ea5cd7, image.Point{}, draw.Over)
	__obf_5e90e389aeb91533.
		Rotate(__obf_b4922ee002e08e21.Angle, false)
	__obf_458c82c2d61f50fd := __obf_5e90e389aeb91533.CalcMarginBlankArea()

	return __obf_5e90e389aeb91533,

		// DrawStringImage draws a text image
		// params:
		//   - dot: Draw dot
		//   - textColor: Text color
		//
		// returns:
		//   - types.NRGBA: Drawn text image
		//   - error: Error information
		__obf_458c82c2d61f50fd, nil
}

func (__obf_3d944245c503e38e *Captcha) DrawStringImage(__obf_b4922ee002e08e21 *types.Dot, __obf_60467093f8b93df8 color.Color) (*types.NRGBA, error) {
	__obf_5e90e389aeb91533 := helper.CreateNRGBACanvas(__obf_b4922ee002e08e21.Size+10, __obf_b4922ee002e08e21.Size+10, true)
	__obf_f198f5dc9244d650 := fixed.P(12, __obf_b4922ee002e08e21.Size-5)
	if util.IsChineseChar(__obf_b4922ee002e08e21.Text) {
		__obf_f198f5dc9244d650 = fixed.P(10, __obf_b4922ee002e08e21.Size)
	}
	__obf_d1d4e3f72f8cd606 := &types.DrawStringParam{
		Color:       __obf_60467093f8b93df8,
		FontSize:    __obf_b4922ee002e08e21.FontSize,
		Size:        __obf_b4922ee002e08e21.Size,
		FontDPI:     __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.FontDPI,
		FontHinting: __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.FontHinting,
		Text:        __obf_b4922ee002e08e21.Text,
	}
	var __obf_43790225c3b3432f error
	__obf_d1d4e3f72f8cd606.
		Font, __obf_43790225c3b3432f = assets.PickFont()
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}
	__obf_43790225c3b3432f = helper.DrawString(__obf_5e90e389aeb91533, __obf_d1d4e3f72f8cd606, __obf_f198f5dc9244d650)
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}

	return __obf_5e90e389aeb91533, nil
}

func (__obf_3d944245c503e38e *Captcha) DrawShapeImage(__obf_b4922ee002e08e21 *types.Dot, __obf_b8a694fd5c1c8919 color.Color) (*types.NRGBA, error) {
	__obf_8b95010f2c8009e6, __obf_d8ecd84fd0499a08, __obf_676c47750b6f1420, __obf_cb98306611433dd9 := __obf_b8a694fd5c1c8919.RGBA()

	var __obf_36c7af87e517b9c9 = []color.RGBA{
		{R: uint8(__obf_8b95010f2c8009e6), G: uint8(__obf_d8ecd84fd0499a08), B: uint8(__obf_676c47750b6f1420), A: uint8(__obf_cb98306611433dd9)},
	}
	__obf_9530b4ea7904a975 := helper.CreateNRGBACanvas(__obf_b4922ee002e08e21.Size+10, __obf_b4922ee002e08e21.Size+10, true)
	var __obf_897cea6b749cef85 image.Rectangle
	var __obf_e3bc7b3b39cb093b image.Image
	if __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.UseRGBA {
		__obf_5e90e389aeb91533 := helper.CreateNRGBACanvas(__obf_b4922ee002e08e21.Size+10, __obf_b4922ee002e08e21.Size+10, true)
		draw.BiLinear.Scale(__obf_5e90e389aeb91533, __obf_5e90e389aeb91533.Bounds(), __obf_b4922ee002e08e21.Shape, __obf_b4922ee002e08e21.Shape.Bounds(), draw.Over, nil)
		__obf_897cea6b749cef85 = __obf_5e90e389aeb91533.Bounds()
		__obf_e3bc7b3b39cb093b = __obf_5e90e389aeb91533
	} else {
		__obf_5e90e389aeb91533 := helper.CreatePaletteCanvas(__obf_b4922ee002e08e21.Size+10, __obf_b4922ee002e08e21.Size+10, __obf_36c7af87e517b9c9)
		draw.BiLinear.Scale(__obf_5e90e389aeb91533, __obf_5e90e389aeb91533.Bounds(), __obf_b4922ee002e08e21.Shape, __obf_b4922ee002e08e21.Shape.Bounds(), draw.Over, nil)
		__obf_897cea6b749cef85 = __obf_5e90e389aeb91533.Bounds()
		__obf_e3bc7b3b39cb093b = __obf_5e90e389aeb91533
	}

	draw.Draw(__obf_9530b4ea7904a975, __obf_897cea6b749cef85, __obf_e3bc7b3b39cb093b, image.Point{}, draw.Over)

	return __obf_9530b4ea7904a975, nil
}

func (__obf_761b3f820674d2e3 *Captcha) __obf_a3d185177a709dd1() (*CaptchaData, error) {

	var (
		__obf_4f1f316b94a446c6, __obf_aff3092bc0f2e76c, __obf_0800445bbca5209b []*types.Dot
		__obf_f800ea4adb319059                                                 []image.Image
		__obf_58d0687063ccf22c                                                 []image.Image
		__obf_d88cc15f0bafbf24, __obf_30d550091f68f974                         image.Image
		__obf_43790225c3b3432f                                                 error
	)
	__obf_58d0687063ccf22c, __obf_43790225c3b3432f = assets.GetShapes()
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}
	__obf_474592b71dcad790 := util.RandInt(__obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.LenRange.Min, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.LenRange.Max)
	__obf_58d0687063ccf22c = util.PickN(__obf_58d0687063ccf22c, __obf_474592b71dcad790)
	__obf_4f1f316b94a446c6 = __obf_761b3f820674d2e3.__obf_a5e86731db6fb612(__obf_58d0687063ccf22c, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.Size, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.SizeRange, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.DotPadding)
	__obf_7a057c77eb19a593 := rand.Perm(len(__obf_4f1f316b94a446c6))
	__obf_8ab4ed134ac216cf := util.RandInt(__obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.VerifyLenRange.Min, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.VerifyLenRange.Max)
	__obf_0800445bbca5209b = make([]*types.Dot, __obf_8ab4ed134ac216cf)
	for __obf_d9fed1663089e979, __obf_ab64545afc28626d := range __obf_7a057c77eb19a593 {
		if __obf_d9fed1663089e979 >= __obf_8ab4ed134ac216cf {
			break
		}
		__obf_b4922ee002e08e21 := __obf_4f1f316b94a446c6[__obf_ab64545afc28626d]
		__obf_b4922ee002e08e21.
			Index = __obf_d9fed1663089e979
		__obf_0800445bbca5209b[__obf_d9fed1663089e979] = __obf_b4922ee002e08e21
		__obf_f800ea4adb319059 = append(__obf_f800ea4adb319059, __obf_0800445bbca5209b[__obf_d9fed1663089e979].Shape)
	}
	__obf_aff3092bc0f2e76c = __obf_761b3f820674d2e3.__obf_a5e86731db6fb612(__obf_f800ea4adb319059, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.Size, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.SizeRange, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.DotPadding)
	__obf_d88cc15f0bafbf24, __obf_43790225c3b3432f = __obf_761b3f820674d2e3.__obf_dc16a15893b5d39a(Shape, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.Size, __obf_4f1f316b94a446c6)
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}
	__obf_30d550091f68f974, __obf_43790225c3b3432f = __obf_761b3f820674d2e3.__obf_4bbe1721f6d33f0d(Shape, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.Size, __obf_aff3092bc0f2e76c)
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}

	return &CaptchaData{__obf_cea1e242f930c716: __obf_0800445bbca5209b, __obf_d88cc15f0bafbf24: types.NewJPEGImage(__obf_d88cc15f0bafbf24), __obf_30d550091f68f974: types.NewPNGImage(__obf_30d550091f68f974)}, nil
}

func (__obf_761b3f820674d2e3 *Captcha) __obf_3a74189fe74ae6c9() (*CaptchaData, error) {
	__obf_474592b71dcad790 := util.RandInt(__obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.LenRange.Min, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.LenRange.Max)
	__obf_672b56148a957cb6 := util.PickN(assets.GetChineseChars(), __obf_474592b71dcad790)

	var __obf_cea1e242f930c716, __obf_aff3092bc0f2e76c, __obf_0800445bbca5209b []*types.Dot
	var __obf_3869ae6f0465f6e2 []string
	var __obf_d88cc15f0bafbf24, __obf_30d550091f68f974 image.Image
	__obf_cea1e242f930c716 = __obf_761b3f820674d2e3.__obf_64ac234ec964fe52(__obf_672b56148a957cb6, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.Size, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.SizeRange, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.DotPadding)
	__obf_7a057c77eb19a593 := rand.Perm(len(__obf_cea1e242f930c716))
	__obf_8ab4ed134ac216cf := util.RandInt(__obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.VerifyLenRange.Min, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.VerifyLenRange.Max)
	__obf_0800445bbca5209b = make([]*types.Dot, __obf_8ab4ed134ac216cf)
	for __obf_d9fed1663089e979, __obf_ab64545afc28626d := range __obf_7a057c77eb19a593 {
		if __obf_d9fed1663089e979 >= __obf_8ab4ed134ac216cf {
			break
		}
		__obf_b4922ee002e08e21 := __obf_cea1e242f930c716[__obf_ab64545afc28626d]
		__obf_b4922ee002e08e21.
			Index = __obf_d9fed1663089e979
		__obf_0800445bbca5209b[__obf_d9fed1663089e979] = __obf_b4922ee002e08e21
		__obf_3869ae6f0465f6e2 = append(__obf_3869ae6f0465f6e2, __obf_0800445bbca5209b[__obf_d9fed1663089e979].Text)
	}
	__obf_aff3092bc0f2e76c = __obf_761b3f820674d2e3.__obf_64ac234ec964fe52(__obf_3869ae6f0465f6e2, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.Size, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.SizeRange, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.DotPadding)
	var __obf_43790225c3b3432f error
	__obf_d88cc15f0bafbf24, __obf_43790225c3b3432f = __obf_761b3f820674d2e3.__obf_dc16a15893b5d39a(Text, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.Size, __obf_cea1e242f930c716)
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}
	__obf_30d550091f68f974, __obf_43790225c3b3432f = __obf_761b3f820674d2e3.__obf_4bbe1721f6d33f0d(Text, __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Secondary.Size, __obf_aff3092bc0f2e76c)
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}

	return &CaptchaData{__obf_cea1e242f930c716: __obf_0800445bbca5209b, __obf_d88cc15f0bafbf24: types.NewJPEGImage(__obf_d88cc15f0bafbf24), __obf_30d550091f68f974: types.NewPNGImage(__obf_30d550091f68f974)}, nil
}

func (__obf_761b3f820674d2e3 *Captcha) __obf_a5e86731db6fb612(__obf_58d0687063ccf22c []image.Image, __obf_846e5dc6f87ca803 types.Size, __obf_4443458cdf3df12c types.Range, __obf_ed27aeedcda6e704 int) []*types.Dot {
	__obf_474592b71dcad790 := len(__obf_58d0687063ccf22c)
	var __obf_cea1e242f930c716 = make([]*types.Dot, __obf_474592b71dcad790)
	__obf_1181c8127158d833 := __obf_846e5dc6f87ca803.Width
	__obf_62c56517f1f23b8f := __obf_846e5dc6f87ca803.Height
	if __obf_ed27aeedcda6e704 > 0 {
		__obf_1181c8127158d833 -= __obf_ed27aeedcda6e704
		__obf_62c56517f1f23b8f -= __obf_ed27aeedcda6e704
	}

	for __obf_d9fed1663089e979, __obf_31d51e7dec1e6919 := range __obf_58d0687063ccf22c {
		__obf_5f7d9856f2b51411 := __obf_761b3f820674d2e3.__obf_5f7d9856f2b51411()
		__obf_548c98e1fb85dc5f := util.PickN(__obf_3d160061e64e17b8, 1)[0]
		__obf_0a4669e7f098d36d := util.PickN(__obf_ec0d6d3545d013bd, 1)[0]
		__obf_d6a715c82e2cceb8 := util.RandInt(__obf_4443458cdf3df12c.Min, __obf_4443458cdf3df12c.Max)
		__obf_701b7b1349750044 := 10
		__obf_35fade7994c3002d := __obf_1181c8127158d833 / __obf_474592b71dcad790
		__obf_03732f538dd903bc := math.Abs(float64(__obf_35fade7994c3002d) - float64(__obf_d6a715c82e2cceb8))
		__obf_8500b38ef204669f := (__obf_d9fed1663089e979 * __obf_35fade7994c3002d) + util.RandInt(0, int(max(__obf_03732f538dd903bc, 1)))
		__obf_54d8d4d8729bb28d := util.RandInt(__obf_701b7b1349750044, __obf_62c56517f1f23b8f+__obf_d6a715c82e2cceb8)
		__obf_a8eab991d4d46530 := int(min(max(float64(__obf_8500b38ef204669f), float64(__obf_701b7b1349750044)), float64(__obf_1181c8127158d833-__obf_701b7b1349750044-(__obf_ed27aeedcda6e704*2))))
		__obf_6f3efd4586e6dcf4 := int(min(max(float64(__obf_54d8d4d8729bb28d), float64(__obf_d6a715c82e2cceb8+__obf_701b7b1349750044)), float64(__obf_62c56517f1f23b8f+(__obf_d6a715c82e2cceb8/2)-(__obf_ed27aeedcda6e704*2))))
		__obf_cea1e242f930c716[__obf_d9fed1663089e979] = &types.Dot{
			Index:          __obf_d9fed1663089e979,
			X:              __obf_a8eab991d4d46530,
			Y:              __obf_6f3efd4586e6dcf4 - __obf_d6a715c82e2cceb8,
			FontSize:       __obf_d6a715c82e2cceb8,
			Size:           __obf_d6a715c82e2cceb8,
			Angle:          __obf_5f7d9856f2b51411,
			PrimaryColor:   __obf_548c98e1fb85dc5f,
			SecondaryColor: __obf_0a4669e7f098d36d,
			Shape:          __obf_31d51e7dec1e6919,
		}
	}

	return __obf_cea1e242f930c716
}

func (__obf_761b3f820674d2e3 *Captcha) __obf_64ac234ec964fe52(__obf_672b56148a957cb6 []string, __obf_846e5dc6f87ca803 types.Size, __obf_bf5d3e3096f04d26 types.Range, __obf_ed27aeedcda6e704 int) []*types.Dot {
	__obf_474592b71dcad790 := len(__obf_672b56148a957cb6)
	var __obf_cea1e242f930c716 = make([]*types.Dot, __obf_474592b71dcad790)
	__obf_1181c8127158d833 := __obf_846e5dc6f87ca803.Width
	__obf_62c56517f1f23b8f := __obf_846e5dc6f87ca803.Height
	if __obf_ed27aeedcda6e704 > 0 {
		__obf_1181c8127158d833 -= __obf_ed27aeedcda6e704
		__obf_62c56517f1f23b8f -= __obf_ed27aeedcda6e704
	}

	for __obf_d9fed1663089e979, __obf_31d51e7dec1e6919 := range __obf_672b56148a957cb6 {
		__obf_5f7d9856f2b51411 := __obf_761b3f820674d2e3.__obf_5f7d9856f2b51411()
		__obf_548c98e1fb85dc5f := util.PickN(__obf_3d160061e64e17b8, 1)[0]
		__obf_0a4669e7f098d36d := util.PickN(__obf_ec0d6d3545d013bd, 1)[0]
		__obf_d6a715c82e2cceb8 := util.RandInt(__obf_bf5d3e3096f04d26.Min, __obf_bf5d3e3096f04d26.Max)
		__obf_7bd8316fbe7be95e := __obf_d6a715c82e2cceb8
		__obf_95536cc6642308e2 := __obf_d6a715c82e2cceb8

		if util.LenChineseChar(__obf_31d51e7dec1e6919) > 1 {
			__obf_95536cc6642308e2 = __obf_d6a715c82e2cceb8 * util.LenChineseChar(__obf_31d51e7dec1e6919)

			if __obf_5f7d9856f2b51411 > 0 {
				__obf_d4594700d52ae5f8 := __obf_95536cc6642308e2 - __obf_d6a715c82e2cceb8
				__obf_e4c6b2d23c5a026c := __obf_5f7d9856f2b51411 % 90
				__obf_a9e4858b013f5fbe := float64(__obf_d4594700d52ae5f8) / 90
				__obf_ef5914876e5fc840 := max(float64(__obf_e4c6b2d23c5a026c)*__obf_a9e4858b013f5fbe, 1)
				__obf_7bd8316fbe7be95e = __obf_7bd8316fbe7be95e + int(__obf_ef5914876e5fc840)
				__obf_95536cc6642308e2 = __obf_95536cc6642308e2 + int(__obf_ef5914876e5fc840)
			}
		}
		__obf_701b7b1349750044 := 10
		__obf_35fade7994c3002d := __obf_1181c8127158d833 / __obf_474592b71dcad790
		__obf_03732f538dd903bc := math.Abs(float64(__obf_35fade7994c3002d) - float64(__obf_95536cc6642308e2))
		__obf_8500b38ef204669f := (__obf_d9fed1663089e979 * __obf_35fade7994c3002d) + util.RandInt(0, int(max(__obf_03732f538dd903bc, 1)))
		__obf_54d8d4d8729bb28d := util.RandInt(__obf_701b7b1349750044, __obf_62c56517f1f23b8f+__obf_7bd8316fbe7be95e)
		__obf_a8eab991d4d46530 := int(min(max(float64(__obf_8500b38ef204669f), float64(__obf_701b7b1349750044)), float64(__obf_1181c8127158d833-__obf_701b7b1349750044-(__obf_ed27aeedcda6e704*2))))
		__obf_6f3efd4586e6dcf4 := int(min(max(float64(__obf_54d8d4d8729bb28d), float64(__obf_7bd8316fbe7be95e+__obf_701b7b1349750044)), float64(__obf_62c56517f1f23b8f+(__obf_7bd8316fbe7be95e/2)-(__obf_ed27aeedcda6e704*2))))
		__obf_cea1e242f930c716[__obf_d9fed1663089e979] = &types.Dot{
			Index:          __obf_d9fed1663089e979,
			X:              __obf_a8eab991d4d46530,
			Y:              __obf_6f3efd4586e6dcf4 - __obf_7bd8316fbe7be95e,
			FontSize:       __obf_d6a715c82e2cceb8,
			Size:           max(__obf_95536cc6642308e2, __obf_7bd8316fbe7be95e),
			Angle:          __obf_5f7d9856f2b51411,
			PrimaryColor:   __obf_548c98e1fb85dc5f,
			SecondaryColor: __obf_0a4669e7f098d36d,
			Text:           __obf_31d51e7dec1e6919,
		}
	}

	return __obf_cea1e242f930c716
}

func (__obf_761b3f820674d2e3 *Captcha) __obf_dc16a15893b5d39a(__obf_ac0f57b9f94f3dbe DotType, __obf_bf5d3e3096f04d26 types.Size, __obf_cea1e242f930c716 []*types.Dot) (image.Image, error) {
	__obf_0f02b4bf4768e772 := &DrawImageParams{
		Width:       __obf_bf5d3e3096f04d26.Width,
		Height:      __obf_bf5d3e3096f04d26.Height,
		Alpha:       __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.Alpha,
		FontHinting: __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.FontHinting,
		Dots:        __obf_cea1e242f930c716,
		ShowShadow:  __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.ShowShadow,
		ShadowPoint: __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.ShadowPoint,
	}
	var __obf_43790225c3b3432f error
	__obf_0f02b4bf4768e772.
		Background, __obf_43790225c3b3432f = assets.PickBgImage()
	if __obf_43790225c3b3432f != nil {
		return nil, __obf_43790225c3b3432f
	}

	return __obf_761b3f820674d2e3.DrawWithNRGBA(__obf_ac0f57b9f94f3dbe, __obf_0f02b4bf4768e772)
}

func (__obf_3d944245c503e38e *Captcha) __obf_4bbe1721f6d33f0d(__obf_ac0f57b9f94f3dbe DotType, __obf_bf5d3e3096f04d26 types.Size, __obf_cea1e242f930c716 []*types.Dot) (image.Image, error) {
	var (
		__obf_5e36be47ec52fd6f =
		// err      error
		make([]*types.Dot, 0, len(__obf_cea1e242f930c716))
	)
	__obf_1181c8127158d833 := __obf_bf5d3e3096f04d26.Width / len(__obf_cea1e242f930c716)

	for __obf_d9fed1663089e979, __obf_31d51e7dec1e6919 := range __obf_cea1e242f930c716 {
		__obf_474592b71dcad790 := 1
		if __obf_ac0f57b9f94f3dbe == Text {
			__obf_474592b71dcad790 = len(__obf_31d51e7dec1e6919.Text)
		}
		__obf_31d51e7dec1e6919.
			X = int(max(float64(__obf_1181c8127158d833*__obf_d9fed1663089e979+__obf_1181c8127158d833/__obf_31d51e7dec1e6919.Size), 8))
		__obf_70bff2c5e643ed2e := max(1, __obf_bf5d3e3096f04d26.Height/16*__obf_474592b71dcad790)
		__obf_31d51e7dec1e6919.
			Y = __obf_bf5d3e3096f04d26.Height/2 + __obf_31d51e7dec1e6919.FontSize/2 - rand.IntN(__obf_70bff2c5e643ed2e)
		__obf_5e36be47ec52fd6f = append(__obf_5e36be47ec52fd6f, __obf_31d51e7dec1e6919)
	}
	__obf_d1d4e3f72f8cd606 := &DrawImageParams{
		Width:             __obf_bf5d3e3096f04d26.Width,
		Height:            __obf_bf5d3e3096f04d26.Height,
		Dots:              __obf_5e36be47ec52fd6f,
		BackgroundDistort: __obf_3d944245c503e38e.__obf_ee2b0bad4455d17c(__obf_3d944245c503e38e.__obf_f8f07b39b37290ac.Secondary.BgDistort),
		BgCircles:         __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.Secondary.BgCircles,
		BgSlimLines:       __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_630f70feb2a4afa3 []color.Color
	for _, __obf_38bbd6b03dca9647 := range __obf_ec0d6d3545d013bd {
		__obf_5aecf180c5bc1e97, _ := util.ParseHexColor(__obf_38bbd6b03dca9647)
		__obf_630f70feb2a4afa3 = append(__obf_630f70feb2a4afa3, __obf_5aecf180c5bc1e97)
	}

	var __obf_c66d486b139ce9bc []color.Color
	for _, __obf_5aecf180c5bc1e97 := range __obf_ec0d6d3545d013bd {
		__obf_e7aaf7358dd58db4, _ := util.ParseHexColor(__obf_5aecf180c5bc1e97)
		__obf_c66d486b139ce9bc = append(__obf_c66d486b139ce9bc, __obf_e7aaf7358dd58db4)
	}

	if __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.UseRGBA || __obf_3d944245c503e38e.__obf_f8f07b39b37290ac.Secondary.NonDeformAble {
		return __obf_3d944245c503e38e.DrawWithNRGBA2(__obf_ac0f57b9f94f3dbe, __obf_d1d4e3f72f8cd606, __obf_630f70feb2a4afa3, __obf_c66d486b139ce9bc)
	}
	return __obf_3d944245c503e38e.DrawWithPalette(__obf_ac0f57b9f94f3dbe, __obf_d1d4e3f72f8cd606, __obf_630f70feb2a4afa3, __obf_c66d486b139ce9bc)
}

func (__obf_761b3f820674d2e3 *Captcha) __obf_ee2b0bad4455d17c(__obf_22fd172e788f7739 int) int {
	switch __obf_22fd172e788f7739 {
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

func (__obf_761b3f820674d2e3 *Captcha) __obf_5f7d9856f2b51411() int {
	__obf_0b821f6f30016270 := __obf_761b3f820674d2e3.__obf_f8f07b39b37290ac.Primary.AnglePosRange
	__obf_9c770e9f829d8df6 := rand.IntN(len(__obf_0b821f6f30016270))
	if __obf_9c770e9f829d8df6 < 0 {
		return 0
	}
	__obf_a7441b41d64c16b3 := __obf_0b821f6f30016270[__obf_9c770e9f829d8df6]
	__obf_457c793f4f785200 := util.RandInt(__obf_a7441b41d64c16b3.Min, __obf_a7441b41d64c16b3.Max)

	return __obf_457c793f4f785200
}
