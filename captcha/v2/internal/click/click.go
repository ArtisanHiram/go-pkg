package __obf_25fa2a83bb367991

import (
	"image"
	"image/color"
	"math"
	"math/rand/v2"

	assets "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/assets"
	helper "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/helper"
	types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"
	util "github.com/ArtisanHiram/go-pkg/captcha/v2/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/fixed"
)

type Captcha struct {
	__obf_0165d5ca2affa4f2 *Options
}

func (__obf_e26e4e8685ed069a *Captcha) SetOptions(__obf_0165d5ca2affa4f2 *Options) {
	__obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2 = __obf_0165d5ca2affa4f2
}

func NewCaptcha(__obf_0165d5ca2affa4f2 *Options) *Captcha {
	if __obf_0165d5ca2affa4f2 == nil {
		return nil
	}
	return &Captcha{__obf_0165d5ca2affa4f2}
}

func __obf_903df99384192099(__obf_0165d5ca2affa4f2 *Options) DotType {
	__obf_bf0eedb256519fdb := DotType(rand.IntN(2))
	if __obf_bf0eedb256519fdb == Shape {
		__obf_0165d5ca2affa4f2.UseRGBA = false
		__obf_0165d5ca2affa4f2.Secondary.BgDistort = types.DistortLevel1
		__obf_0165d5ca2affa4f2.Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_0165d5ca2affa4f2.Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_0165d5ca2affa4f2.UseRGBA = true
		__obf_0165d5ca2affa4f2.Secondary.BgDistort = types.DistortLevel4
		__obf_0165d5ca2affa4f2.Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_0165d5ca2affa4f2.Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_bf0eedb256519fdb
}

func (__obf_e26e4e8685ed069a *Captcha) DrawWithNRGBA(__obf_b33cdabadfa1c77a DotType, __obf_6517b7b0b8d95035 *DrawImageParams) (image.Image, error) {
	__obf_8ad847a1c4d1bb62 := __obf_6517b7b0b8d95035.Dots
	__obf_e2fec375dbeba962 := helper.CreateNRGBACanvas(__obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height, true)

	for _, __obf_bf17c6ed91bf6d7a := range __obf_8ad847a1c4d1bb62 {

		__obf_182890480c435aae, __obf_7b625d1c61e39e61, __obf_97caac8d824f79e8 := __obf_e26e4e8685ed069a.DrawDotImage(__obf_b33cdabadfa1c77a, __obf_bf17c6ed91bf6d7a, __obf_6517b7b0b8d95035)
		if __obf_97caac8d824f79e8 != nil {
			return nil, __obf_97caac8d824f79e8
		}
		__obf_af4f4068e4a97eff := __obf_7b625d1c61e39e61.MinX
		__obf_65580fba8ff2766b := __obf_7b625d1c61e39e61.MaxX
		__obf_66323e76789badf6 := __obf_7b625d1c61e39e61.MinY
		__obf_95fcd5619036722a := __obf_7b625d1c61e39e61.MaxY

		__obf_73ceb9f368768ce6 := __obf_65580fba8ff2766b - __obf_af4f4068e4a97eff
		__obf_b4863a4af23b9182 := __obf_95fcd5619036722a - __obf_66323e76789badf6

		draw.Draw(__obf_e2fec375dbeba962, image.Rect(__obf_bf17c6ed91bf6d7a.X, __obf_bf17c6ed91bf6d7a.Y, __obf_bf17c6ed91bf6d7a.X+__obf_73ceb9f368768ce6, __obf_bf17c6ed91bf6d7a.Y+__obf_b4863a4af23b9182), __obf_182890480c435aae, image.Point{X: __obf_af4f4068e4a97eff, Y: __obf_66323e76789badf6}, draw.Over)

	}

	__obf_c370552153fc41b0 := __obf_6517b7b0b8d95035.Background
	__obf_ffdf79b2752bb9ab := __obf_e2fec375dbeba962.Bounds()
	__obf_248367172d2f6835 := helper.CreateNRGBACanvas(__obf_ffdf79b2752bb9ab.Dx(), __obf_ffdf79b2752bb9ab.Dx(), true)
	__obf_258946cf86904834 := util.RangCutImagePos(__obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height, __obf_c370552153fc41b0)
	draw.Draw(__obf_248367172d2f6835, __obf_ffdf79b2752bb9ab, __obf_c370552153fc41b0, __obf_258946cf86904834, draw.Src)
	draw.Draw(__obf_248367172d2f6835, __obf_e2fec375dbeba962.Bounds(), __obf_e2fec375dbeba962, image.Point{}, draw.Over)
	__obf_248367172d2f6835.SubImage(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height))
	return __obf_248367172d2f6835, nil
}

func (__obf_e26e4e8685ed069a *Captcha) DrawWithPalette(__obf_b33cdabadfa1c77a DotType, __obf_6517b7b0b8d95035 *DrawImageParams, __obf_ac9475d8561e9726 []color.Color, __obf_8dc5f028b0705463 []color.Color) (image.Image, error) {
	__obf_8ad847a1c4d1bb62 := __obf_6517b7b0b8d95035.Dots

	__obf_6132503b25126e14 := make([]color.Color, 0, len(__obf_8dc5f028b0705463))
	for _, __obf_5a015837a17da16f := range __obf_8dc5f028b0705463 {
		__obf_25ee4fc1cbeefec8, __obf_1c5b1cb1cb8a7be6, __obf_ffdf79b2752bb9ab, _ := __obf_5a015837a17da16f.RGBA()
		__obf_6d1764a6f15e1db4 := util.FormatAlpha(__obf_6517b7b0b8d95035.ThumbDisturbAlpha)
		__obf_6132503b25126e14 = append(__obf_6132503b25126e14, color.RGBA{R: uint8(__obf_25ee4fc1cbeefec8), G: uint8(__obf_1c5b1cb1cb8a7be6), B: uint8(__obf_ffdf79b2752bb9ab), A: __obf_6d1764a6f15e1db4})
	}

	var __obf_e62845a6f6576fa2 = make([]color.Color, 0, len(__obf_ac9475d8561e9726)+len(__obf_6132503b25126e14)+1)
	__obf_e62845a6f6576fa2 = append(__obf_e62845a6f6576fa2, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_e62845a6f6576fa2 = append(__obf_e62845a6f6576fa2, __obf_ac9475d8561e9726...)
	__obf_e62845a6f6576fa2 = append(__obf_e62845a6f6576fa2, __obf_6132503b25126e14...)

	__obf_e2fec375dbeba962 := types.NewPalette(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height), __obf_e62845a6f6576fa2)
	if __obf_6517b7b0b8d95035.BgCircles > 0 {
		__obf_e26e4e8685ed069a.__obf_0f9e689fe35ba26a(__obf_e2fec375dbeba962, __obf_6517b7b0b8d95035.BgCircles, 1, __obf_6132503b25126e14)
	}
	if __obf_6517b7b0b8d95035.BgSlimLines > 0 {
		__obf_e26e4e8685ed069a.__obf_55fb27372d462639(__obf_e2fec375dbeba962, __obf_6517b7b0b8d95035.BgSlimLines, __obf_6132503b25126e14)
	}

	for _, __obf_bf17c6ed91bf6d7a := range __obf_8ad847a1c4d1bb62 {
		__obf_da3674b816a64466, _ := util.ParseHexColor(__obf_bf17c6ed91bf6d7a.PrimaryColor)
		var __obf_97caac8d824f79e8 error
		if __obf_b33cdabadfa1c77a == Shape {
			var __obf_182890480c435aae *types.NRGBA
			__obf_182890480c435aae, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawShapeImage(__obf_bf17c6ed91bf6d7a, __obf_da3674b816a64466)
			if __obf_97caac8d824f79e8 != nil {
				return nil, __obf_97caac8d824f79e8
			}

			__obf_182890480c435aae.Rotate(__obf_bf17c6ed91bf6d7a.Angle, false)
			__obf_cb5fd991b08eb851 := __obf_e2fec375dbeba962.Bounds()
			__obf_68e8c08550054ef9 := __obf_182890480c435aae.Bounds()
			__obf_d3e15d72125846f4 := image.Point{X: __obf_cb5fd991b08eb851.Dx() - __obf_68e8c08550054ef9.Dx(), Y: __obf_cb5fd991b08eb851.Dy() - __obf_68e8c08550054ef9.Dy()}
			draw.Draw(__obf_e2fec375dbeba962, image.Rect(__obf_bf17c6ed91bf6d7a.X, __obf_d3e15d72125846f4.Y, __obf_d3e15d72125846f4.X+__obf_68e8c08550054ef9.Dx(), __obf_d3e15d72125846f4.Y+__obf_68e8c08550054ef9.Dy()), __obf_182890480c435aae, image.Point{}, draw.Over)
		} else {
			__obf_9145bf829faa1ecc := fixed.P(__obf_bf17c6ed91bf6d7a.X, __obf_bf17c6ed91bf6d7a.Y)
			__obf_6517b7b0b8d95035 := &types.DrawStringParam{
				Color:       __obf_da3674b816a64466,
				FontSize:    __obf_bf17c6ed91bf6d7a.FontSize,
				Size:        __obf_bf17c6ed91bf6d7a.Size,
				FontDPI:     __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.FontDPI,
				FontHinting: __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.FontHinting,
				Text:        __obf_bf17c6ed91bf6d7a.Text,
			}
			__obf_6517b7b0b8d95035.Font, __obf_97caac8d824f79e8 = assets.PickFont()
			if __obf_97caac8d824f79e8 != nil {
				return nil, __obf_97caac8d824f79e8
			}
			__obf_97caac8d824f79e8 = helper.DrawString(__obf_e2fec375dbeba962, __obf_6517b7b0b8d95035, __obf_9145bf829faa1ecc)
		}

		if __obf_97caac8d824f79e8 != nil {
			return __obf_e2fec375dbeba962, __obf_97caac8d824f79e8
		}
	}

	if __obf_6517b7b0b8d95035.Background != nil {
		__obf_c370552153fc41b0 := __obf_6517b7b0b8d95035.Background
		__obf_ffdf79b2752bb9ab := __obf_c370552153fc41b0.Bounds()
		__obf_248367172d2f6835 := helper.CreateNRGBACanvas(__obf_ffdf79b2752bb9ab.Dx(), __obf_ffdf79b2752bb9ab.Dy(), true)
		__obf_258946cf86904834 := util.RangCutImagePos(__obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height, __obf_c370552153fc41b0)
		draw.Draw(__obf_248367172d2f6835, __obf_ffdf79b2752bb9ab, __obf_c370552153fc41b0, __obf_258946cf86904834, draw.Src)
		__obf_e2fec375dbeba962.Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_248367172d2f6835, __obf_e2fec375dbeba962.Bounds(), __obf_e2fec375dbeba962, image.Point{}, draw.Over)
		__obf_950b8de3fa1d5c0c := __obf_248367172d2f6835.SubImage(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height)).(*image.NRGBA)
		return __obf_950b8de3fa1d5c0c, nil
	}

	if __obf_6517b7b0b8d95035.BackgroundDistort > 0 {
		__obf_e2fec375dbeba962.Distort(float64(util.RandInt(5, 10)), float64(__obf_6517b7b0b8d95035.BackgroundDistort))
	}

	return __obf_e2fec375dbeba962, nil

}

func (__obf_e26e4e8685ed069a *Captcha) DrawWithNRGBA2(__obf_b33cdabadfa1c77a DotType, __obf_6517b7b0b8d95035 *DrawImageParams, __obf_ac9475d8561e9726 []color.Color, __obf_8dc5f028b0705463 []color.Color) (image.Image, error) {
	__obf_8ad847a1c4d1bb62 := __obf_6517b7b0b8d95035.Dots

	__obf_6132503b25126e14 := make([]color.Color, 0, len(__obf_8dc5f028b0705463))
	for _, __obf_5a015837a17da16f := range __obf_8dc5f028b0705463 {
		__obf_25ee4fc1cbeefec8, __obf_1c5b1cb1cb8a7be6, __obf_ffdf79b2752bb9ab, _ := __obf_5a015837a17da16f.RGBA()
		__obf_6d1764a6f15e1db4 := util.FormatAlpha(__obf_6517b7b0b8d95035.ThumbDisturbAlpha)
		__obf_6132503b25126e14 = append(__obf_6132503b25126e14, color.RGBA{R: uint8(__obf_25ee4fc1cbeefec8), G: uint8(__obf_1c5b1cb1cb8a7be6), B: uint8(__obf_ffdf79b2752bb9ab), A: __obf_6d1764a6f15e1db4})
	}

	var __obf_e62845a6f6576fa2 = make([]color.Color, 0, len(__obf_ac9475d8561e9726)+len(__obf_6132503b25126e14)+1)
	__obf_e62845a6f6576fa2 = append(__obf_e62845a6f6576fa2, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_e62845a6f6576fa2 = append(__obf_e62845a6f6576fa2, __obf_ac9475d8561e9726...)
	__obf_e62845a6f6576fa2 = append(__obf_e62845a6f6576fa2, __obf_6132503b25126e14...)

	__obf_ce8a52714a4e6e92 := types.NewNRGBA(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height), true)
	if __obf_6517b7b0b8d95035.Background != nil {
		__obf_c370552153fc41b0 := __obf_6517b7b0b8d95035.Background
		__obf_ffdf79b2752bb9ab := __obf_c370552153fc41b0.Bounds()
		__obf_248367172d2f6835 := helper.CreateNRGBACanvas(__obf_ffdf79b2752bb9ab.Dx(), __obf_ffdf79b2752bb9ab.Dy(), true)
		__obf_258946cf86904834 := util.RangCutImagePos(__obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height, __obf_c370552153fc41b0)
		draw.Draw(__obf_248367172d2f6835, __obf_ffdf79b2752bb9ab, __obf_c370552153fc41b0, __obf_258946cf86904834, draw.Src)
		__obf_950b8de3fa1d5c0c := __obf_248367172d2f6835.SubImage(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height)).(*image.NRGBA)
		draw.Draw(__obf_ce8a52714a4e6e92, __obf_950b8de3fa1d5c0c.Bounds(), __obf_950b8de3fa1d5c0c, image.Point{}, draw.Over)
	}

	__obf_e2fec375dbeba962 := types.NewPalette(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height), __obf_e62845a6f6576fa2)
	if __obf_6517b7b0b8d95035.BgCircles > 0 {
		__obf_e26e4e8685ed069a.__obf_0f9e689fe35ba26a(__obf_e2fec375dbeba962, __obf_6517b7b0b8d95035.BgCircles, 1, __obf_6132503b25126e14)
	}
	if __obf_6517b7b0b8d95035.BgSlimLines > 0 {
		__obf_e26e4e8685ed069a.__obf_55fb27372d462639(__obf_e2fec375dbeba962, __obf_6517b7b0b8d95035.BgSlimLines, __obf_6132503b25126e14)
	}
	if __obf_6517b7b0b8d95035.BackgroundDistort > 0 {
		__obf_e2fec375dbeba962.Distort(float64(util.RandInt(5, 10)), float64(__obf_6517b7b0b8d95035.BackgroundDistort))
	}

	__obf_8936fc7d4e104a59 := __obf_e2fec375dbeba962.Bounds()
	__obf_73ceb9f368768ce6 := __obf_8936fc7d4e104a59.Dx() / len(__obf_8ad847a1c4d1bb62)
	for __obf_0063e9b84dd66db9, __obf_bf17c6ed91bf6d7a := range __obf_8ad847a1c4d1bb62 {
		__obf_da3674b816a64466, _ := util.ParseHexColor(__obf_bf17c6ed91bf6d7a.PrimaryColor)
		var __obf_97caac8d824f79e8 error
		if __obf_b33cdabadfa1c77a == Shape {
			var __obf_182890480c435aae *types.NRGBA
			__obf_182890480c435aae, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawShapeImage(__obf_bf17c6ed91bf6d7a, __obf_da3674b816a64466)
			if __obf_97caac8d824f79e8 != nil {
				return nil, __obf_97caac8d824f79e8
			}

			__obf_182890480c435aae.Rotate(__obf_bf17c6ed91bf6d7a.Angle, false)

			__obf_cb5fd991b08eb851 := __obf_ce8a52714a4e6e92.Bounds()
			__obf_68e8c08550054ef9 := __obf_182890480c435aae.Bounds()
			__obf_d3e15d72125846f4 := image.Point{X: __obf_cb5fd991b08eb851.Dx() - __obf_68e8c08550054ef9.Dx(), Y: __obf_cb5fd991b08eb851.Dy() - __obf_68e8c08550054ef9.Dy()}
			draw.Draw(__obf_ce8a52714a4e6e92, image.Rect(__obf_bf17c6ed91bf6d7a.X, __obf_d3e15d72125846f4.Y, __obf_d3e15d72125846f4.X+__obf_68e8c08550054ef9.Dx(), __obf_d3e15d72125846f4.Y+__obf_68e8c08550054ef9.Dy()), __obf_182890480c435aae, image.Point{}, draw.Over)
		} else {
			var __obf_5205fe16f7755e4b *types.NRGBA
			__obf_5205fe16f7755e4b, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawStringImage(__obf_bf17c6ed91bf6d7a, __obf_da3674b816a64466)
			if __obf_97caac8d824f79e8 != nil {
				return nil, __obf_97caac8d824f79e8
			}
			__obf_5205fe16f7755e4b.Rotate(__obf_bf17c6ed91bf6d7a.Angle, false)

			__obf_7b625d1c61e39e61 := __obf_5205fe16f7755e4b.CalcMarginBlankArea()
			__obf_af4f4068e4a97eff := __obf_7b625d1c61e39e61.MinX
			__obf_65580fba8ff2766b := __obf_7b625d1c61e39e61.MaxX
			__obf_66323e76789badf6 := __obf_7b625d1c61e39e61.MinY
			__obf_95fcd5619036722a := __obf_7b625d1c61e39e61.MaxY
			__obf_5205fe16f7755e4b.SubImage(image.Rect(__obf_af4f4068e4a97eff, __obf_66323e76789badf6, __obf_65580fba8ff2766b, __obf_95fcd5619036722a))
			__obf_cf156f78b910adf6 := __obf_5205fe16f7755e4b.Bounds()

			__obf_a807555efb3ee019 := int(max(float64(__obf_73ceb9f368768ce6*__obf_0063e9b84dd66db9+__obf_73ceb9f368768ce6/__obf_cf156f78b910adf6.Dx()), 8))
			__obf_ca28de2ad27ecc00 := util.RandInt(1, __obf_8936fc7d4e104a59.Dy()-__obf_cf156f78b910adf6.Dy()-4)

			draw.Draw(__obf_ce8a52714a4e6e92, image.Rect(__obf_a807555efb3ee019, __obf_ca28de2ad27ecc00, __obf_a807555efb3ee019+__obf_cf156f78b910adf6.Dx(), __obf_ca28de2ad27ecc00+__obf_cf156f78b910adf6.Dy()), __obf_5205fe16f7755e4b, image.Point{X: __obf_cf156f78b910adf6.Min.X, Y: __obf_cf156f78b910adf6.Min.Y}, draw.Over)
		}

		if __obf_97caac8d824f79e8 != nil {
			return __obf_ce8a52714a4e6e92, __obf_97caac8d824f79e8
		}
	}

	__obf_e64f371b1e0bfe7f := types.NewNRGBA(image.Rect(0, 0, __obf_6517b7b0b8d95035.Width, __obf_6517b7b0b8d95035.Height), true)
	draw.Draw(__obf_e64f371b1e0bfe7f, __obf_e2fec375dbeba962.Bounds(), __obf_e2fec375dbeba962, image.Point{}, draw.Over)
	draw.Draw(__obf_ce8a52714a4e6e92, __obf_e64f371b1e0bfe7f.Bounds(), __obf_e64f371b1e0bfe7f, image.Point{}, draw.Over)
	return __obf_ce8a52714a4e6e92, nil
}

func (__obf_e26e4e8685ed069a *Captcha) __obf_0f9e689fe35ba26a(__obf_248367172d2f6835 *types.Palette, __obf_0a49b71ca176cf05, __obf_26f3debbbae4578b int, __obf_dace03535993bee3 []color.Color) {
	__obf_148c43a6b576b700 := __obf_248367172d2f6835.Bounds().Max.X
	__obf_e404238ddb5774c4 := __obf_248367172d2f6835.Bounds().Max.Y
	for range __obf_0a49b71ca176cf05 {
		__obf_2b9d9188865fc0c5 := util.RandColor(__obf_dace03535993bee3)
		//co.A = uint8(0xee)
		__obf_25ee4fc1cbeefec8 := util.RandInt(1, __obf_26f3debbbae4578b)
		__obf_248367172d2f6835.DrawCircle(util.RandInt(__obf_25ee4fc1cbeefec8, __obf_148c43a6b576b700-__obf_25ee4fc1cbeefec8), util.RandInt(__obf_25ee4fc1cbeefec8, __obf_e404238ddb5774c4-__obf_25ee4fc1cbeefec8), __obf_25ee4fc1cbeefec8, __obf_2b9d9188865fc0c5)
	}
}

func (__obf_e26e4e8685ed069a *Captcha) __obf_55fb27372d462639(__obf_248367172d2f6835 *types.Palette, __obf_adc15de44ac2ad76 int, __obf_dace03535993bee3 []color.Color) {
	__obf_683b1bdee3d25723 := __obf_248367172d2f6835.Bounds().Max.X / 10
	__obf_3696cb276a99da14 := __obf_683b1bdee3d25723 * 9
	__obf_a1d042285b822244 := __obf_248367172d2f6835.Bounds().Max.Y / 3
	for __obf_0063e9b84dd66db9 := range __obf_adc15de44ac2ad76 {
		__obf_917e440ee012bcfb := image.Point{X: rand.IntN(__obf_683b1bdee3d25723), Y: rand.IntN(__obf_a1d042285b822244)}
		__obf_e04659cf9eedf40d := image.Point{X: rand.IntN(__obf_683b1bdee3d25723) + __obf_3696cb276a99da14, Y: rand.IntN(__obf_a1d042285b822244)}

		if __obf_0063e9b84dd66db9%2 == 0 {
			__obf_917e440ee012bcfb.Y = rand.IntN(__obf_a1d042285b822244) + __obf_a1d042285b822244*2
			__obf_e04659cf9eedf40d.Y = rand.IntN(__obf_a1d042285b822244)
		} else {
			__obf_917e440ee012bcfb.Y = rand.IntN(__obf_a1d042285b822244) + __obf_a1d042285b822244*(__obf_0063e9b84dd66db9%2)
			__obf_e04659cf9eedf40d.Y = rand.IntN(__obf_a1d042285b822244) + __obf_a1d042285b822244*2
		}

		__obf_2b9d9188865fc0c5 := util.RandColor(__obf_dace03535993bee3)
		//co.A = uint8(0xee)
		__obf_248367172d2f6835.DrawBeeline(__obf_917e440ee012bcfb, __obf_e04659cf9eedf40d, __obf_2b9d9188865fc0c5)
	}
}

func (__obf_e26e4e8685ed069a *Captcha) DrawDotImage(__obf_b33cdabadfa1c77a DotType, __obf_c8b7e75042fa4c1c *types.Dot, __obf_6517b7b0b8d95035 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_da3674b816a64466, _ := util.ParseHexColor(__obf_c8b7e75042fa4c1c.PrimaryColor)
	__obf_da3674b816a64466.A = util.FormatAlpha(__obf_6517b7b0b8d95035.Alpha)

	var __obf_5205fe16f7755e4b image.Image
	var __obf_97caac8d824f79e8 error
	if __obf_b33cdabadfa1c77a == Shape {
		__obf_5205fe16f7755e4b, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawShapeImage(__obf_c8b7e75042fa4c1c, __obf_da3674b816a64466)
	} else {
		__obf_5205fe16f7755e4b, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawStringImage(__obf_c8b7e75042fa4c1c, __obf_da3674b816a64466)
	}
	if __obf_97caac8d824f79e8 != nil {
		return nil, nil, __obf_97caac8d824f79e8
	}

	__obf_c50f3fb11c197198, _ := util.ParseHexColor(__obf_5c62e8d2accc90a0)
	__obf_e2fec375dbeba962 := helper.CreateNRGBACanvas(__obf_c8b7e75042fa4c1c.Size+10, __obf_c8b7e75042fa4c1c.Size+10, true)
	if __obf_6517b7b0b8d95035.ShowShadow {
		var __obf_b445d2316d8692dd *types.NRGBA
		if __obf_b33cdabadfa1c77a == Shape {
			__obf_b445d2316d8692dd, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawShapeImage(__obf_c8b7e75042fa4c1c, __obf_c50f3fb11c197198)
		} else {
			__obf_b445d2316d8692dd, __obf_97caac8d824f79e8 = __obf_e26e4e8685ed069a.DrawStringImage(__obf_c8b7e75042fa4c1c, __obf_c50f3fb11c197198)
		}
		if __obf_97caac8d824f79e8 != nil {
			return nil, nil, __obf_97caac8d824f79e8
		}

		__obf_e1e08ac65d69ce23 := __obf_6517b7b0b8d95035.ShadowPoint.X
		__obf_d9b0a8395ac1499b := __obf_6517b7b0b8d95035.ShadowPoint.Y
		draw.Draw(__obf_e2fec375dbeba962, __obf_b445d2316d8692dd.Bounds(), __obf_b445d2316d8692dd, image.Point{X: __obf_e1e08ac65d69ce23, Y: __obf_d9b0a8395ac1499b}, draw.Over)
	}
	draw.Draw(__obf_e2fec375dbeba962, __obf_5205fe16f7755e4b.Bounds(), __obf_5205fe16f7755e4b, image.Point{}, draw.Over)
	__obf_e2fec375dbeba962.Rotate(__obf_c8b7e75042fa4c1c.Angle, false)

	__obf_7a45d3c603b6076d := __obf_e2fec375dbeba962.CalcMarginBlankArea()

	return __obf_e2fec375dbeba962, __obf_7a45d3c603b6076d, nil
}

// DrawStringImage draws a text image
// params:
//   - dot: Draw dot
//   - textColor: Text color
//
// returns:
//   - types.NRGBA: Drawn text image
//   - error: Error information
func (__obf_e26e4e8685ed069a *Captcha) DrawStringImage(__obf_c8b7e75042fa4c1c *types.Dot, __obf_9c193f97a88b536c color.Color) (*types.NRGBA, error) {
	__obf_e2fec375dbeba962 := helper.CreateNRGBACanvas(__obf_c8b7e75042fa4c1c.Size+10, __obf_c8b7e75042fa4c1c.Size+10, true)

	__obf_9145bf829faa1ecc := fixed.P(12, __obf_c8b7e75042fa4c1c.Size-5)
	if util.IsChineseChar(__obf_c8b7e75042fa4c1c.Text) {
		__obf_9145bf829faa1ecc = fixed.P(10, __obf_c8b7e75042fa4c1c.Size)
	}
	__obf_6517b7b0b8d95035 := &types.DrawStringParam{
		Color:       __obf_9c193f97a88b536c,
		FontSize:    __obf_c8b7e75042fa4c1c.FontSize,
		Size:        __obf_c8b7e75042fa4c1c.Size,
		FontDPI:     __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.FontDPI,
		FontHinting: __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.FontHinting,
		Text:        __obf_c8b7e75042fa4c1c.Text,
	}
	var __obf_97caac8d824f79e8 error
	__obf_6517b7b0b8d95035.Font, __obf_97caac8d824f79e8 = assets.PickFont()
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}
	__obf_97caac8d824f79e8 = helper.DrawString(__obf_e2fec375dbeba962, __obf_6517b7b0b8d95035, __obf_9145bf829faa1ecc)
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	return __obf_e2fec375dbeba962, nil
}

func (__obf_e26e4e8685ed069a *Captcha) DrawShapeImage(__obf_c8b7e75042fa4c1c *types.Dot, __obf_da3674b816a64466 color.Color) (*types.NRGBA, error) {
	__obf_c0047b46cf83cd33, __obf_d2bc20c52ce3cccd, __obf_2557f4c9faa93f85, __obf_87544532058a9698 := __obf_da3674b816a64466.RGBA()

	var __obf_73a8e8d6710a9ce6 = []color.RGBA{
		{R: uint8(__obf_c0047b46cf83cd33), G: uint8(__obf_d2bc20c52ce3cccd), B: uint8(__obf_2557f4c9faa93f85), A: uint8(__obf_87544532058a9698)},
	}

	__obf_e64f371b1e0bfe7f := helper.CreateNRGBACanvas(__obf_c8b7e75042fa4c1c.Size+10, __obf_c8b7e75042fa4c1c.Size+10, true)
	var __obf_cf156f78b910adf6 image.Rectangle
	var __obf_c370552153fc41b0 image.Image
	if __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.UseRGBA {
		__obf_e2fec375dbeba962 := helper.CreateNRGBACanvas(__obf_c8b7e75042fa4c1c.Size+10, __obf_c8b7e75042fa4c1c.Size+10, true)
		draw.BiLinear.Scale(__obf_e2fec375dbeba962, __obf_e2fec375dbeba962.Bounds(), __obf_c8b7e75042fa4c1c.Shape, __obf_c8b7e75042fa4c1c.Shape.Bounds(), draw.Over, nil)
		__obf_cf156f78b910adf6 = __obf_e2fec375dbeba962.Bounds()
		__obf_c370552153fc41b0 = __obf_e2fec375dbeba962
	} else {
		__obf_e2fec375dbeba962 := helper.CreatePaletteCanvas(__obf_c8b7e75042fa4c1c.Size+10, __obf_c8b7e75042fa4c1c.Size+10, __obf_73a8e8d6710a9ce6)
		draw.BiLinear.Scale(__obf_e2fec375dbeba962, __obf_e2fec375dbeba962.Bounds(), __obf_c8b7e75042fa4c1c.Shape, __obf_c8b7e75042fa4c1c.Shape.Bounds(), draw.Over, nil)
		__obf_cf156f78b910adf6 = __obf_e2fec375dbeba962.Bounds()
		__obf_c370552153fc41b0 = __obf_e2fec375dbeba962
	}

	draw.Draw(__obf_e64f371b1e0bfe7f, __obf_cf156f78b910adf6, __obf_c370552153fc41b0, image.Point{}, draw.Over)

	return __obf_e64f371b1e0bfe7f, nil
}

func (__obf_d5c9aa819977350f *Captcha) Generate() (*CaptchaData, error) {
	__obf_b33cdabadfa1c77a := __obf_903df99384192099(__obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2)
	if __obf_b33cdabadfa1c77a == Shape {
		return __obf_d5c9aa819977350f.__obf_7da4e8cc090632ee()
	}

	return __obf_d5c9aa819977350f.__obf_9518176a81c6d5fc()
}

func (__obf_d5c9aa819977350f *Captcha) __obf_7da4e8cc090632ee() (*CaptchaData, error) {

	var (
		__obf_1ff6f3ef40c11cfa, __obf_b63f242380353235, __obf_3606e81ddd32d149 []*types.Dot
		__obf_962d2354a2fcfc66                                                 []image.Image
		__obf_54cb61a7c5697e74                                                 []image.Image
		__obf_ff8c367564ca56d3, __obf_4a1046d8157ce073                         image.Image
		__obf_97caac8d824f79e8                                                 error
	)
	__obf_54cb61a7c5697e74, __obf_97caac8d824f79e8 = assets.GetShapes()
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	__obf_205eaf6d696fc78f := util.RandInt(__obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.LenRange.Min, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.LenRange.Max)
	__obf_54cb61a7c5697e74 = util.PickN(__obf_54cb61a7c5697e74, __obf_205eaf6d696fc78f)

	__obf_1ff6f3ef40c11cfa = __obf_d5c9aa819977350f.__obf_4c5e684f8eeb8924(__obf_54cb61a7c5697e74, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.Size, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.SizeRange, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.DotPadding)
	__obf_2da8e31f95b50039 := rand.Perm(len(__obf_1ff6f3ef40c11cfa))
	__obf_81a9e9a183981b3a := util.RandInt(__obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.VerifyLenRange.Min, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.VerifyLenRange.Max)
	__obf_3606e81ddd32d149 = make([]*types.Dot, __obf_81a9e9a183981b3a)
	for __obf_0063e9b84dd66db9, __obf_4861179109ebd094 := range __obf_2da8e31f95b50039 {
		if __obf_0063e9b84dd66db9 >= __obf_81a9e9a183981b3a {
			break
		}

		__obf_c8b7e75042fa4c1c := __obf_1ff6f3ef40c11cfa[__obf_4861179109ebd094]
		__obf_c8b7e75042fa4c1c.Index = __obf_0063e9b84dd66db9
		__obf_3606e81ddd32d149[__obf_0063e9b84dd66db9] = __obf_c8b7e75042fa4c1c
		__obf_962d2354a2fcfc66 = append(__obf_962d2354a2fcfc66, __obf_3606e81ddd32d149[__obf_0063e9b84dd66db9].Shape)
	}
	__obf_b63f242380353235 = __obf_d5c9aa819977350f.__obf_4c5e684f8eeb8924(__obf_962d2354a2fcfc66, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.Size, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.SizeRange, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.DotPadding)

	__obf_ff8c367564ca56d3, __obf_97caac8d824f79e8 = __obf_d5c9aa819977350f.__obf_e43bd65826440e4b(Shape, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.Size, __obf_1ff6f3ef40c11cfa)
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	__obf_4a1046d8157ce073, __obf_97caac8d824f79e8 = __obf_d5c9aa819977350f.__obf_0b1872de8df51e91(Shape, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.Size, __obf_b63f242380353235)
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	return &CaptchaData{
		__obf_8ad847a1c4d1bb62: __obf_3606e81ddd32d149,
		__obf_ff8c367564ca56d3: types.NewJPEGImage(__obf_ff8c367564ca56d3),
		__obf_4a1046d8157ce073: types.NewPNGImage(__obf_4a1046d8157ce073),
	}, nil
}

func (__obf_d5c9aa819977350f *Captcha) __obf_9518176a81c6d5fc() (*CaptchaData, error) {

	__obf_205eaf6d696fc78f := util.RandInt(__obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.LenRange.Min, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.LenRange.Max)
	__obf_352ac1b3e3e10061 := util.PickN(assets.GetChineseChars(), __obf_205eaf6d696fc78f)

	var __obf_8ad847a1c4d1bb62, __obf_b63f242380353235, __obf_3606e81ddd32d149 []*types.Dot
	var __obf_0f3a57a3a2d60896 []string
	var __obf_ff8c367564ca56d3, __obf_4a1046d8157ce073 image.Image

	__obf_8ad847a1c4d1bb62 = __obf_d5c9aa819977350f.__obf_567d8d9c088d167c(__obf_352ac1b3e3e10061, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.Size, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.SizeRange, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.DotPadding)
	__obf_2da8e31f95b50039 := rand.Perm(len(__obf_8ad847a1c4d1bb62))
	__obf_81a9e9a183981b3a := util.RandInt(__obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.VerifyLenRange.Min, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.VerifyLenRange.Max)
	__obf_3606e81ddd32d149 = make([]*types.Dot, __obf_81a9e9a183981b3a)
	for __obf_0063e9b84dd66db9, __obf_4861179109ebd094 := range __obf_2da8e31f95b50039 {
		if __obf_0063e9b84dd66db9 >= __obf_81a9e9a183981b3a {
			break
		}

		__obf_c8b7e75042fa4c1c := __obf_8ad847a1c4d1bb62[__obf_4861179109ebd094]
		__obf_c8b7e75042fa4c1c.Index = __obf_0063e9b84dd66db9
		__obf_3606e81ddd32d149[__obf_0063e9b84dd66db9] = __obf_c8b7e75042fa4c1c
		__obf_0f3a57a3a2d60896 = append(__obf_0f3a57a3a2d60896, __obf_3606e81ddd32d149[__obf_0063e9b84dd66db9].Text)
	}

	__obf_b63f242380353235 = __obf_d5c9aa819977350f.__obf_567d8d9c088d167c(__obf_0f3a57a3a2d60896, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.Size, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.SizeRange, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.DotPadding)
	var __obf_97caac8d824f79e8 error
	__obf_ff8c367564ca56d3, __obf_97caac8d824f79e8 = __obf_d5c9aa819977350f.__obf_e43bd65826440e4b(Text, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.Size, __obf_8ad847a1c4d1bb62)
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	__obf_4a1046d8157ce073, __obf_97caac8d824f79e8 = __obf_d5c9aa819977350f.__obf_0b1872de8df51e91(Text, __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Secondary.Size, __obf_b63f242380353235)
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	return &CaptchaData{
		__obf_8ad847a1c4d1bb62: __obf_3606e81ddd32d149,
		__obf_ff8c367564ca56d3: types.NewJPEGImage(__obf_ff8c367564ca56d3),
		__obf_4a1046d8157ce073: types.NewPNGImage(__obf_4a1046d8157ce073),
	}, nil
}

func (__obf_d5c9aa819977350f *Captcha) __obf_4c5e684f8eeb8924(__obf_54cb61a7c5697e74 []image.Image, __obf_3b8e052ebfe84975 types.Size, __obf_d263cfb0750e00ea types.Range, __obf_e3eb7eebc5334755 int) []*types.Dot {

	__obf_205eaf6d696fc78f := len(__obf_54cb61a7c5697e74)
	var __obf_8ad847a1c4d1bb62 = make([]*types.Dot, __obf_205eaf6d696fc78f)
	__obf_73ceb9f368768ce6 := __obf_3b8e052ebfe84975.Width
	__obf_b4863a4af23b9182 := __obf_3b8e052ebfe84975.Height
	if __obf_e3eb7eebc5334755 > 0 {
		__obf_73ceb9f368768ce6 -= __obf_e3eb7eebc5334755
		__obf_b4863a4af23b9182 -= __obf_e3eb7eebc5334755
	}

	for __obf_0063e9b84dd66db9, __obf_bf17c6ed91bf6d7a := range __obf_54cb61a7c5697e74 {
		__obf_7982009cbf503fb8 := __obf_d5c9aa819977350f.__obf_7982009cbf503fb8()

		__obf_b080d9c82ceddab4 := util.PickN(__obf_44bd2e60e057d134, 1)[0]
		__obf_51cc5ee50cad4055 := util.PickN(__obf_648216acf80325e3, 1)[0]

		__obf_d855acd1f6ce085b := util.RandInt(__obf_d263cfb0750e00ea.Min, __obf_d263cfb0750e00ea.Max)

		__obf_ca28de2ad27ecc00 := 10
		__obf_e41f9902d04a6a24 := __obf_73ceb9f368768ce6 / __obf_205eaf6d696fc78f
		__obf_34e57093f8f772d2 := math.Abs(float64(__obf_e41f9902d04a6a24) - float64(__obf_d855acd1f6ce085b))
		__obf_2c3925b3e30b9fff := (__obf_0063e9b84dd66db9 * __obf_e41f9902d04a6a24) + util.RandInt(0, int(max(__obf_34e57093f8f772d2, 1)))
		__obf_7f9fab24ed20aad9 := util.RandInt(__obf_ca28de2ad27ecc00, __obf_b4863a4af23b9182+__obf_d855acd1f6ce085b)

		__obf_3c9bd1196287d6df := int(min(max(float64(__obf_2c3925b3e30b9fff), float64(__obf_ca28de2ad27ecc00)), float64(__obf_73ceb9f368768ce6-__obf_ca28de2ad27ecc00-(__obf_e3eb7eebc5334755*2))))
		__obf_a1d042285b822244 := int(min(max(float64(__obf_7f9fab24ed20aad9), float64(__obf_d855acd1f6ce085b+__obf_ca28de2ad27ecc00)), float64(__obf_b4863a4af23b9182+(__obf_d855acd1f6ce085b/2)-(__obf_e3eb7eebc5334755*2))))

		__obf_8ad847a1c4d1bb62[__obf_0063e9b84dd66db9] = &types.Dot{
			Index:          __obf_0063e9b84dd66db9,
			X:              __obf_3c9bd1196287d6df,
			Y:              __obf_a1d042285b822244 - __obf_d855acd1f6ce085b,
			FontSize:       __obf_d855acd1f6ce085b,
			Size:           __obf_d855acd1f6ce085b,
			Angle:          __obf_7982009cbf503fb8,
			PrimaryColor:   __obf_b080d9c82ceddab4,
			SecondaryColor: __obf_51cc5ee50cad4055,
			Shape:          __obf_bf17c6ed91bf6d7a,
		}
	}

	return __obf_8ad847a1c4d1bb62
}

func (__obf_d5c9aa819977350f *Captcha) __obf_567d8d9c088d167c(__obf_352ac1b3e3e10061 []string, __obf_3b8e052ebfe84975 types.Size, __obf_87ad0c5c74a355db types.Range, __obf_e3eb7eebc5334755 int) []*types.Dot {
	__obf_205eaf6d696fc78f := len(__obf_352ac1b3e3e10061)
	var __obf_8ad847a1c4d1bb62 = make([]*types.Dot, __obf_205eaf6d696fc78f)
	__obf_73ceb9f368768ce6 := __obf_3b8e052ebfe84975.Width
	__obf_b4863a4af23b9182 := __obf_3b8e052ebfe84975.Height
	if __obf_e3eb7eebc5334755 > 0 {
		__obf_73ceb9f368768ce6 -= __obf_e3eb7eebc5334755
		__obf_b4863a4af23b9182 -= __obf_e3eb7eebc5334755
	}

	for __obf_0063e9b84dd66db9, __obf_bf17c6ed91bf6d7a := range __obf_352ac1b3e3e10061 {
		__obf_7982009cbf503fb8 := __obf_d5c9aa819977350f.__obf_7982009cbf503fb8()

		__obf_b080d9c82ceddab4 := util.PickN(__obf_44bd2e60e057d134, 1)[0]
		__obf_51cc5ee50cad4055 := util.PickN(__obf_648216acf80325e3, 1)[0]

		__obf_d855acd1f6ce085b := util.RandInt(__obf_87ad0c5c74a355db.Min, __obf_87ad0c5c74a355db.Max)
		__obf_4519e34dc4c1aa24 := __obf_d855acd1f6ce085b
		__obf_42338929f65546e0 := __obf_d855acd1f6ce085b

		if util.LenChineseChar(__obf_bf17c6ed91bf6d7a) > 1 {
			__obf_42338929f65546e0 = __obf_d855acd1f6ce085b * util.LenChineseChar(__obf_bf17c6ed91bf6d7a)

			if __obf_7982009cbf503fb8 > 0 {
				__obf_87e78523fb3f7118 := __obf_42338929f65546e0 - __obf_d855acd1f6ce085b
				__obf_7cf623a140694e68 := __obf_7982009cbf503fb8 % 90
				__obf_7ff12e9a9ee3479c := float64(__obf_87e78523fb3f7118) / 90
				__obf_25ee4fc1cbeefec8 := max(float64(__obf_7cf623a140694e68)*__obf_7ff12e9a9ee3479c, 1)
				__obf_4519e34dc4c1aa24 = __obf_4519e34dc4c1aa24 + int(__obf_25ee4fc1cbeefec8)
				__obf_42338929f65546e0 = __obf_42338929f65546e0 + int(__obf_25ee4fc1cbeefec8)
			}
		}

		__obf_ca28de2ad27ecc00 := 10
		__obf_e41f9902d04a6a24 := __obf_73ceb9f368768ce6 / __obf_205eaf6d696fc78f
		__obf_34e57093f8f772d2 := math.Abs(float64(__obf_e41f9902d04a6a24) - float64(__obf_42338929f65546e0))
		__obf_2c3925b3e30b9fff := (__obf_0063e9b84dd66db9 * __obf_e41f9902d04a6a24) + util.RandInt(0, int(max(__obf_34e57093f8f772d2, 1)))
		__obf_7f9fab24ed20aad9 := util.RandInt(__obf_ca28de2ad27ecc00, __obf_b4863a4af23b9182+__obf_4519e34dc4c1aa24)

		__obf_3c9bd1196287d6df := int(min(max(float64(__obf_2c3925b3e30b9fff), float64(__obf_ca28de2ad27ecc00)), float64(__obf_73ceb9f368768ce6-__obf_ca28de2ad27ecc00-(__obf_e3eb7eebc5334755*2))))
		__obf_a1d042285b822244 := int(min(max(float64(__obf_7f9fab24ed20aad9), float64(__obf_4519e34dc4c1aa24+__obf_ca28de2ad27ecc00)), float64(__obf_b4863a4af23b9182+(__obf_4519e34dc4c1aa24/2)-(__obf_e3eb7eebc5334755*2))))

		__obf_8ad847a1c4d1bb62[__obf_0063e9b84dd66db9] = &types.Dot{
			Index:          __obf_0063e9b84dd66db9,
			X:              __obf_3c9bd1196287d6df,
			Y:              __obf_a1d042285b822244 - __obf_4519e34dc4c1aa24,
			FontSize:       __obf_d855acd1f6ce085b,
			Size:           max(__obf_42338929f65546e0, __obf_4519e34dc4c1aa24),
			Angle:          __obf_7982009cbf503fb8,
			PrimaryColor:   __obf_b080d9c82ceddab4,
			SecondaryColor: __obf_51cc5ee50cad4055,
			Text:           __obf_bf17c6ed91bf6d7a,
		}
	}

	return __obf_8ad847a1c4d1bb62
}

func (__obf_d5c9aa819977350f *Captcha) __obf_e43bd65826440e4b(__obf_b33cdabadfa1c77a DotType, __obf_87ad0c5c74a355db types.Size, __obf_8ad847a1c4d1bb62 []*types.Dot) (image.Image, error) {

	__obf_e505995db7277fff := &DrawImageParams{
		Width:       __obf_87ad0c5c74a355db.Width,
		Height:      __obf_87ad0c5c74a355db.Height,
		Alpha:       __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.Alpha,
		FontHinting: __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.FontHinting,
		Dots:        __obf_8ad847a1c4d1bb62,
		ShowShadow:  __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.ShowShadow,
		ShadowPoint: __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.ShadowPoint,
	}
	var __obf_97caac8d824f79e8 error
	__obf_e505995db7277fff.Background, __obf_97caac8d824f79e8 = assets.PickBgImage()
	if __obf_97caac8d824f79e8 != nil {
		return nil, __obf_97caac8d824f79e8
	}

	return __obf_d5c9aa819977350f.DrawWithNRGBA(__obf_b33cdabadfa1c77a, __obf_e505995db7277fff)
}

func (__obf_e26e4e8685ed069a *Captcha) __obf_0b1872de8df51e91(__obf_b33cdabadfa1c77a DotType, __obf_87ad0c5c74a355db types.Size, __obf_8ad847a1c4d1bb62 []*types.Dot) (image.Image, error) {
	var (
		// err      error
		__obf_f5c4a5a899beb559 = make([]*types.Dot, 0, len(__obf_8ad847a1c4d1bb62))
	)

	__obf_73ceb9f368768ce6 := __obf_87ad0c5c74a355db.Width / len(__obf_8ad847a1c4d1bb62)

	for __obf_0063e9b84dd66db9, __obf_bf17c6ed91bf6d7a := range __obf_8ad847a1c4d1bb62 {
		__obf_205eaf6d696fc78f := 1
		if __obf_b33cdabadfa1c77a == Text {
			__obf_205eaf6d696fc78f = len(__obf_bf17c6ed91bf6d7a.Text)
		}

		__obf_bf17c6ed91bf6d7a.X = int(max(float64(__obf_73ceb9f368768ce6*__obf_0063e9b84dd66db9+__obf_73ceb9f368768ce6/__obf_bf17c6ed91bf6d7a.Size), 8))
		__obf_519593997e16eba9 := max(1, __obf_87ad0c5c74a355db.Height/16*__obf_205eaf6d696fc78f)
		__obf_bf17c6ed91bf6d7a.Y = __obf_87ad0c5c74a355db.Height/2 + __obf_bf17c6ed91bf6d7a.FontSize/2 - rand.IntN(__obf_519593997e16eba9)

		__obf_f5c4a5a899beb559 = append(__obf_f5c4a5a899beb559, __obf_bf17c6ed91bf6d7a)
	}

	__obf_6517b7b0b8d95035 := &DrawImageParams{
		Width:             __obf_87ad0c5c74a355db.Width,
		Height:            __obf_87ad0c5c74a355db.Height,
		Dots:              __obf_f5c4a5a899beb559,
		BackgroundDistort: __obf_e26e4e8685ed069a.__obf_68810d1c7245a7cf(__obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.Secondary.BgDistort),
		BgCircles:         __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.Secondary.BgCircles,
		BgSlimLines:       __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_cff244a672f53262 []color.Color
	for _, __obf_af95c93e8f1b75da := range __obf_648216acf80325e3 {
		__obf_2b9d9188865fc0c5, _ := util.ParseHexColor(__obf_af95c93e8f1b75da)
		__obf_cff244a672f53262 = append(__obf_cff244a672f53262, __obf_2b9d9188865fc0c5)
	}

	var __obf_8dc5f028b0705463 []color.Color
	for _, __obf_2b9d9188865fc0c5 := range __obf_648216acf80325e3 {
		__obf_950b8de3fa1d5c0c, _ := util.ParseHexColor(__obf_2b9d9188865fc0c5)
		__obf_8dc5f028b0705463 = append(__obf_8dc5f028b0705463, __obf_950b8de3fa1d5c0c)
	}

	if __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.UseRGBA || __obf_e26e4e8685ed069a.__obf_0165d5ca2affa4f2.Secondary.NonDeformAble {
		return __obf_e26e4e8685ed069a.DrawWithNRGBA2(__obf_b33cdabadfa1c77a, __obf_6517b7b0b8d95035, __obf_cff244a672f53262, __obf_8dc5f028b0705463)
	}
	return __obf_e26e4e8685ed069a.DrawWithPalette(__obf_b33cdabadfa1c77a, __obf_6517b7b0b8d95035, __obf_cff244a672f53262, __obf_8dc5f028b0705463)
}

func (__obf_d5c9aa819977350f *Captcha) __obf_68810d1c7245a7cf(__obf_3255262f4f327763 int) int {
	switch __obf_3255262f4f327763 {
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

func (__obf_d5c9aa819977350f *Captcha) __obf_7982009cbf503fb8() int {
	__obf_8c60de4e1f54e46e := __obf_d5c9aa819977350f.__obf_0165d5ca2affa4f2.Primary.AnglePosRange

	__obf_62d0d51b968485cc := rand.IntN(len(__obf_8c60de4e1f54e46e))
	if __obf_62d0d51b968485cc < 0 {
		return 0
	}

	__obf_d5f6db75e73a7080 := __obf_8c60de4e1f54e46e[__obf_62d0d51b968485cc]
	__obf_b788719733557f80 := util.RandInt(__obf_d5f6db75e73a7080.Min, __obf_d5f6db75e73a7080.Max)

	return __obf_b788719733557f80
}
