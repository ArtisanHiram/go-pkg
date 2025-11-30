package __obf_59290109f8ce9c8f

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
	__obf_03afc9d99f21e175 *types.ClickOption
}

var _ types.Captcha = (*Captcha)(nil)

func NewCaptcha(__obf_03afc9d99f21e175 *types.ClickOption) *Captcha {
	if __obf_03afc9d99f21e175 == nil {
		return nil
	}
	return &Captcha{__obf_03afc9d99f21e175}
}

func (__obf_e0f9d59d598c10d8 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_41af80e859213b9c := __obf_0fd005ee4925dc74(__obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175)
	if __obf_41af80e859213b9c == Shape {
		return __obf_e0f9d59d598c10d8.__obf_4dfb9c5752c4ac93()
	}

	return __obf_e0f9d59d598c10d8.__obf_591266b39896d5bb()
}

func __obf_0fd005ee4925dc74(__obf_03afc9d99f21e175 *types.ClickOption) DotType {
	__obf_40caeaed1f746e6f := DotType(rand.IntN(2))
	if __obf_40caeaed1f746e6f == Shape {
		__obf_03afc9d99f21e175.
			UseRGBA = false
		__obf_03afc9d99f21e175.
			Secondary.BgDistort = types.DistortLevel1
		__obf_03afc9d99f21e175.
			Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_03afc9d99f21e175.
			Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_03afc9d99f21e175.
			UseRGBA = true
		__obf_03afc9d99f21e175.
			Secondary.BgDistort = types.DistortLevel4
		__obf_03afc9d99f21e175.
			Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_03afc9d99f21e175.
			Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_40caeaed1f746e6f
}

func (__obf_2a577d1bd2f12f6a *Captcha) DrawWithNRGBA(__obf_41af80e859213b9c DotType, __obf_b6321ad01abb7fde *DrawImageParams) (image.Image, error) {
	__obf_7f68597a0d34a028 := __obf_b6321ad01abb7fde.Dots
	__obf_9ba2ff5f86cfd090 := helper.CreateNRGBACanvas(__obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height, true)

	for _, __obf_1ed0b3f2e37769e4 := range __obf_7f68597a0d34a028 {
		__obf_26379a5758c87b01, __obf_62888636e933d9b2, __obf_573fb4e2cef8adee := __obf_2a577d1bd2f12f6a.DrawDotImage(__obf_41af80e859213b9c, __obf_1ed0b3f2e37769e4, __obf_b6321ad01abb7fde)
		if __obf_573fb4e2cef8adee != nil {
			return nil, __obf_573fb4e2cef8adee
		}
		__obf_34b7b90bfaf2c602 := __obf_62888636e933d9b2.MinX
		__obf_8f624cc219e56317 := __obf_62888636e933d9b2.MaxX
		__obf_4562cd4e5db453e1 := __obf_62888636e933d9b2.MinY
		__obf_0eb78f25a3d79b07 := __obf_62888636e933d9b2.MaxY
		__obf_a4543061eb0a2e0b := __obf_8f624cc219e56317 - __obf_34b7b90bfaf2c602
		__obf_c69f0a9feb9621f7 := __obf_0eb78f25a3d79b07 - __obf_4562cd4e5db453e1

		draw.Draw(__obf_9ba2ff5f86cfd090, image.Rect(__obf_1ed0b3f2e37769e4.X, __obf_1ed0b3f2e37769e4.Y, __obf_1ed0b3f2e37769e4.X+__obf_a4543061eb0a2e0b, __obf_1ed0b3f2e37769e4.Y+__obf_c69f0a9feb9621f7), __obf_26379a5758c87b01, image.Point{X: __obf_34b7b90bfaf2c602, Y: __obf_4562cd4e5db453e1}, draw.Over)

	}
	__obf_00aaa760d6f92456 := __obf_b6321ad01abb7fde.Background
	__obf_1f80226272fb8495 := __obf_9ba2ff5f86cfd090.Bounds()
	__obf_30bd90aa2dac79ac := helper.CreateNRGBACanvas(__obf_1f80226272fb8495.Dx(), __obf_1f80226272fb8495.Dx(), true)
	__obf_0c9a3599125b7cc1 := util.RangCutImagePos(__obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height, __obf_00aaa760d6f92456)
	draw.Draw(__obf_30bd90aa2dac79ac, __obf_1f80226272fb8495, __obf_00aaa760d6f92456, __obf_0c9a3599125b7cc1, draw.Src)
	draw.Draw(__obf_30bd90aa2dac79ac, __obf_9ba2ff5f86cfd090.Bounds(), __obf_9ba2ff5f86cfd090, image.Point{}, draw.Over)
	__obf_30bd90aa2dac79ac.
		SubImage(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height))
	return __obf_30bd90aa2dac79ac, nil
}

func (__obf_2a577d1bd2f12f6a *Captcha) DrawWithPalette(__obf_41af80e859213b9c DotType, __obf_b6321ad01abb7fde *DrawImageParams, __obf_fed3b4086390dac9 []color.Color, __obf_09570dc7552c46d3 []color.Color) (image.Image, error) {
	__obf_7f68597a0d34a028 := __obf_b6321ad01abb7fde.Dots
	__obf_f337fe71e636eecd := make([]color.Color, 0, len(__obf_09570dc7552c46d3))
	for _, __obf_5cc66552ad2bdba5 := range __obf_09570dc7552c46d3 {
		__obf_89aa6f66c78c925b, __obf_f90d03eae3c6b3fe, __obf_1f80226272fb8495, _ := __obf_5cc66552ad2bdba5.RGBA()
		__obf_7e1af467fe4f5fcf := util.FormatAlpha(__obf_b6321ad01abb7fde.ThumbDisturbAlpha)
		__obf_f337fe71e636eecd = append(__obf_f337fe71e636eecd, color.RGBA{R: uint8(__obf_89aa6f66c78c925b), G: uint8(__obf_f90d03eae3c6b3fe), B: uint8(__obf_1f80226272fb8495), A: __obf_7e1af467fe4f5fcf})
	}

	var __obf_b289c59fc077a945 = make([]color.Color, 0, len(__obf_fed3b4086390dac9)+len(__obf_f337fe71e636eecd)+1)
	__obf_b289c59fc077a945 = append(__obf_b289c59fc077a945, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_b289c59fc077a945 = append(__obf_b289c59fc077a945, __obf_fed3b4086390dac9...)
	__obf_b289c59fc077a945 = append(__obf_b289c59fc077a945, __obf_f337fe71e636eecd...)
	__obf_9ba2ff5f86cfd090 := types.NewPalette(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height), __obf_b289c59fc077a945)
	if __obf_b6321ad01abb7fde.BgCircles > 0 {
		__obf_2a577d1bd2f12f6a.__obf_ede1b10008196f44(__obf_9ba2ff5f86cfd090, __obf_b6321ad01abb7fde.BgCircles, 1, __obf_f337fe71e636eecd)
	}
	if __obf_b6321ad01abb7fde.BgSlimLines > 0 {
		__obf_2a577d1bd2f12f6a.__obf_41d4c5c3b8871cbd(__obf_9ba2ff5f86cfd090, __obf_b6321ad01abb7fde.BgSlimLines, __obf_f337fe71e636eecd)
	}

	for _, __obf_1ed0b3f2e37769e4 := range __obf_7f68597a0d34a028 {
		__obf_aa90d8e6b96cb78e, _ := util.ParseHexColor(__obf_1ed0b3f2e37769e4.PrimaryColor)
		var __obf_573fb4e2cef8adee error
		if __obf_41af80e859213b9c == Shape {
			var __obf_26379a5758c87b01 *types.NRGBA
			__obf_26379a5758c87b01, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawShapeImage(__obf_1ed0b3f2e37769e4, __obf_aa90d8e6b96cb78e)
			if __obf_573fb4e2cef8adee != nil {
				return nil, __obf_573fb4e2cef8adee
			}
			__obf_26379a5758c87b01.
				Rotate(__obf_1ed0b3f2e37769e4.Angle, false)
			__obf_67b77108295f3194 := __obf_9ba2ff5f86cfd090.Bounds()
			__obf_2617ea8f20c6d075 := __obf_26379a5758c87b01.Bounds()
			__obf_679bf68e264b3b93 := image.Point{X: __obf_67b77108295f3194.Dx() - __obf_2617ea8f20c6d075.Dx(), Y: __obf_67b77108295f3194.Dy() - __obf_2617ea8f20c6d075.Dy()}
			draw.Draw(__obf_9ba2ff5f86cfd090, image.Rect(__obf_1ed0b3f2e37769e4.X, __obf_679bf68e264b3b93.Y, __obf_679bf68e264b3b93.X+__obf_2617ea8f20c6d075.Dx(), __obf_679bf68e264b3b93.Y+__obf_2617ea8f20c6d075.Dy()), __obf_26379a5758c87b01, image.Point{}, draw.Over)
		} else {
			__obf_87e8685cf28e17a7 := fixed.P(__obf_1ed0b3f2e37769e4.X, __obf_1ed0b3f2e37769e4.Y)
			__obf_b6321ad01abb7fde := &types.DrawStringParam{
				Color:       __obf_aa90d8e6b96cb78e,
				FontSize:    __obf_1ed0b3f2e37769e4.FontSize,
				Size:        __obf_1ed0b3f2e37769e4.Size,
				FontDPI:     __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.FontDPI,
				FontHinting: __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.FontHinting,
				Text:        __obf_1ed0b3f2e37769e4.Text,
			}
			__obf_b6321ad01abb7fde.
				Font, __obf_573fb4e2cef8adee = assets.PickFont()
			if __obf_573fb4e2cef8adee != nil {
				return nil, __obf_573fb4e2cef8adee
			}
			__obf_573fb4e2cef8adee = helper.DrawString(__obf_9ba2ff5f86cfd090, __obf_b6321ad01abb7fde, __obf_87e8685cf28e17a7)
		}

		if __obf_573fb4e2cef8adee != nil {
			return __obf_9ba2ff5f86cfd090, __obf_573fb4e2cef8adee
		}
	}

	if __obf_b6321ad01abb7fde.Background != nil {
		__obf_00aaa760d6f92456 := __obf_b6321ad01abb7fde.Background
		__obf_1f80226272fb8495 := __obf_00aaa760d6f92456.Bounds()
		__obf_30bd90aa2dac79ac := helper.CreateNRGBACanvas(__obf_1f80226272fb8495.Dx(), __obf_1f80226272fb8495.Dy(), true)
		__obf_0c9a3599125b7cc1 := util.RangCutImagePos(__obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height, __obf_00aaa760d6f92456)
		draw.Draw(__obf_30bd90aa2dac79ac, __obf_1f80226272fb8495, __obf_00aaa760d6f92456, __obf_0c9a3599125b7cc1, draw.Src)
		__obf_9ba2ff5f86cfd090.
			Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_30bd90aa2dac79ac, __obf_9ba2ff5f86cfd090.Bounds(), __obf_9ba2ff5f86cfd090, image.Point{}, draw.Over)
		__obf_709b6d468bf64292 := __obf_30bd90aa2dac79ac.SubImage(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height)).(*image.NRGBA)
		return __obf_709b6d468bf64292, nil
	}

	if __obf_b6321ad01abb7fde.BackgroundDistort > 0 {
		__obf_9ba2ff5f86cfd090.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_b6321ad01abb7fde.BackgroundDistort))
	}

	return __obf_9ba2ff5f86cfd090, nil

}

func (__obf_2a577d1bd2f12f6a *Captcha) DrawWithNRGBA2(__obf_41af80e859213b9c DotType, __obf_b6321ad01abb7fde *DrawImageParams, __obf_fed3b4086390dac9 []color.Color, __obf_09570dc7552c46d3 []color.Color) (image.Image, error) {
	__obf_7f68597a0d34a028 := __obf_b6321ad01abb7fde.Dots
	__obf_f337fe71e636eecd := make([]color.Color, 0, len(__obf_09570dc7552c46d3))
	for _, __obf_5cc66552ad2bdba5 := range __obf_09570dc7552c46d3 {
		__obf_89aa6f66c78c925b, __obf_f90d03eae3c6b3fe, __obf_1f80226272fb8495, _ := __obf_5cc66552ad2bdba5.RGBA()
		__obf_7e1af467fe4f5fcf := util.FormatAlpha(__obf_b6321ad01abb7fde.ThumbDisturbAlpha)
		__obf_f337fe71e636eecd = append(__obf_f337fe71e636eecd, color.RGBA{R: uint8(__obf_89aa6f66c78c925b), G: uint8(__obf_f90d03eae3c6b3fe), B: uint8(__obf_1f80226272fb8495), A: __obf_7e1af467fe4f5fcf})
	}

	var __obf_b289c59fc077a945 = make([]color.Color, 0, len(__obf_fed3b4086390dac9)+len(__obf_f337fe71e636eecd)+1)
	__obf_b289c59fc077a945 = append(__obf_b289c59fc077a945, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_b289c59fc077a945 = append(__obf_b289c59fc077a945, __obf_fed3b4086390dac9...)
	__obf_b289c59fc077a945 = append(__obf_b289c59fc077a945, __obf_f337fe71e636eecd...)
	__obf_1349d00f1bc11da7 := types.NewNRGBA(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height), true)
	if __obf_b6321ad01abb7fde.Background != nil {
		__obf_00aaa760d6f92456 := __obf_b6321ad01abb7fde.Background
		__obf_1f80226272fb8495 := __obf_00aaa760d6f92456.Bounds()
		__obf_30bd90aa2dac79ac := helper.CreateNRGBACanvas(__obf_1f80226272fb8495.Dx(), __obf_1f80226272fb8495.Dy(), true)
		__obf_0c9a3599125b7cc1 := util.RangCutImagePos(__obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height, __obf_00aaa760d6f92456)
		draw.Draw(__obf_30bd90aa2dac79ac, __obf_1f80226272fb8495, __obf_00aaa760d6f92456, __obf_0c9a3599125b7cc1, draw.Src)
		__obf_709b6d468bf64292 := __obf_30bd90aa2dac79ac.SubImage(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height)).(*image.NRGBA)
		draw.Draw(__obf_1349d00f1bc11da7, __obf_709b6d468bf64292.Bounds(), __obf_709b6d468bf64292, image.Point{}, draw.Over)
	}
	__obf_9ba2ff5f86cfd090 := types.NewPalette(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height), __obf_b289c59fc077a945)
	if __obf_b6321ad01abb7fde.BgCircles > 0 {
		__obf_2a577d1bd2f12f6a.__obf_ede1b10008196f44(__obf_9ba2ff5f86cfd090, __obf_b6321ad01abb7fde.BgCircles, 1, __obf_f337fe71e636eecd)
	}
	if __obf_b6321ad01abb7fde.BgSlimLines > 0 {
		__obf_2a577d1bd2f12f6a.__obf_41d4c5c3b8871cbd(__obf_9ba2ff5f86cfd090, __obf_b6321ad01abb7fde.BgSlimLines, __obf_f337fe71e636eecd)
	}
	if __obf_b6321ad01abb7fde.BackgroundDistort > 0 {
		__obf_9ba2ff5f86cfd090.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_b6321ad01abb7fde.BackgroundDistort))
	}
	__obf_d26eb04885b3d13a := __obf_9ba2ff5f86cfd090.Bounds()
	__obf_a4543061eb0a2e0b := __obf_d26eb04885b3d13a.Dx() / len(__obf_7f68597a0d34a028)
	for __obf_c599f3c3680a56f5, __obf_1ed0b3f2e37769e4 := range __obf_7f68597a0d34a028 {
		__obf_aa90d8e6b96cb78e, _ := util.ParseHexColor(__obf_1ed0b3f2e37769e4.PrimaryColor)
		var __obf_573fb4e2cef8adee error
		if __obf_41af80e859213b9c == Shape {
			var __obf_26379a5758c87b01 *types.NRGBA
			__obf_26379a5758c87b01, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawShapeImage(__obf_1ed0b3f2e37769e4, __obf_aa90d8e6b96cb78e)
			if __obf_573fb4e2cef8adee != nil {
				return nil, __obf_573fb4e2cef8adee
			}
			__obf_26379a5758c87b01.
				Rotate(__obf_1ed0b3f2e37769e4.Angle, false)
			__obf_67b77108295f3194 := __obf_1349d00f1bc11da7.Bounds()
			__obf_2617ea8f20c6d075 := __obf_26379a5758c87b01.Bounds()
			__obf_679bf68e264b3b93 := image.Point{X: __obf_67b77108295f3194.Dx() - __obf_2617ea8f20c6d075.Dx(), Y: __obf_67b77108295f3194.Dy() - __obf_2617ea8f20c6d075.Dy()}
			draw.Draw(__obf_1349d00f1bc11da7, image.Rect(__obf_1ed0b3f2e37769e4.X, __obf_679bf68e264b3b93.Y, __obf_679bf68e264b3b93.X+__obf_2617ea8f20c6d075.Dx(), __obf_679bf68e264b3b93.Y+__obf_2617ea8f20c6d075.Dy()), __obf_26379a5758c87b01, image.Point{}, draw.Over)
		} else {
			var __obf_c8074fb0da8cfe92 *types.NRGBA
			__obf_c8074fb0da8cfe92, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawStringImage(__obf_1ed0b3f2e37769e4, __obf_aa90d8e6b96cb78e)
			if __obf_573fb4e2cef8adee != nil {
				return nil, __obf_573fb4e2cef8adee
			}
			__obf_c8074fb0da8cfe92.
				Rotate(__obf_1ed0b3f2e37769e4.Angle, false)
			__obf_62888636e933d9b2 := __obf_c8074fb0da8cfe92.CalcMarginBlankArea()
			__obf_34b7b90bfaf2c602 := __obf_62888636e933d9b2.MinX
			__obf_8f624cc219e56317 := __obf_62888636e933d9b2.MaxX
			__obf_4562cd4e5db453e1 := __obf_62888636e933d9b2.MinY
			__obf_0eb78f25a3d79b07 := __obf_62888636e933d9b2.MaxY
			__obf_c8074fb0da8cfe92.
				SubImage(image.Rect(__obf_34b7b90bfaf2c602, __obf_4562cd4e5db453e1, __obf_8f624cc219e56317, __obf_0eb78f25a3d79b07))
			__obf_e0a70f9b0e5e3710 := __obf_c8074fb0da8cfe92.Bounds()
			__obf_99a914e6be8dbe53 := int(max(float64(__obf_a4543061eb0a2e0b*__obf_c599f3c3680a56f5+__obf_a4543061eb0a2e0b/__obf_e0a70f9b0e5e3710.Dx()), 8))
			__obf_0b5eeb13b8b2d52f := util.RandInt(1, __obf_d26eb04885b3d13a.Dy()-__obf_e0a70f9b0e5e3710.Dy()-4)

			draw.Draw(__obf_1349d00f1bc11da7, image.Rect(__obf_99a914e6be8dbe53, __obf_0b5eeb13b8b2d52f, __obf_99a914e6be8dbe53+__obf_e0a70f9b0e5e3710.Dx(), __obf_0b5eeb13b8b2d52f+__obf_e0a70f9b0e5e3710.Dy()), __obf_c8074fb0da8cfe92, image.Point{X: __obf_e0a70f9b0e5e3710.Min.X, Y: __obf_e0a70f9b0e5e3710.Min.Y}, draw.Over)
		}

		if __obf_573fb4e2cef8adee != nil {
			return __obf_1349d00f1bc11da7, __obf_573fb4e2cef8adee
		}
	}
	__obf_02ceaa5d68e46dc9 := types.NewNRGBA(image.Rect(0, 0, __obf_b6321ad01abb7fde.Width, __obf_b6321ad01abb7fde.Height), true)
	draw.Draw(__obf_02ceaa5d68e46dc9, __obf_9ba2ff5f86cfd090.Bounds(), __obf_9ba2ff5f86cfd090, image.Point{}, draw.Over)
	draw.Draw(__obf_1349d00f1bc11da7, __obf_02ceaa5d68e46dc9.Bounds(), __obf_02ceaa5d68e46dc9, image.Point{}, draw.Over)
	return __obf_1349d00f1bc11da7, nil
}

func (__obf_2a577d1bd2f12f6a *Captcha) __obf_ede1b10008196f44(__obf_30bd90aa2dac79ac *types.Palette, __obf_6aa50d8e23daa683, __obf_9a9384dffbb933f7 int, __obf_7ef864bf74ca829e []color.Color) {
	__obf_4ad17217d34cd253 := __obf_30bd90aa2dac79ac.Bounds().Max.X
	__obf_4a9ac5110d1c4421 := __obf_30bd90aa2dac79ac.Bounds().Max.Y
	for range __obf_6aa50d8e23daa683 {
		__obf_f77412067062c4da := util.RandColor(__obf_7ef864bf74ca829e)
		__obf_89aa6f66c78c925b := //co.A = uint8(0xee)
			util.RandInt(1, __obf_9a9384dffbb933f7)
		__obf_30bd90aa2dac79ac.
			DrawCircle(util.RandInt(__obf_89aa6f66c78c925b, __obf_4ad17217d34cd253-__obf_89aa6f66c78c925b), util.RandInt(__obf_89aa6f66c78c925b, __obf_4a9ac5110d1c4421-__obf_89aa6f66c78c925b), __obf_89aa6f66c78c925b, __obf_f77412067062c4da)
	}
}

func (__obf_2a577d1bd2f12f6a *Captcha) __obf_41d4c5c3b8871cbd(__obf_30bd90aa2dac79ac *types.Palette, __obf_9df9b41f9f317bf5 int, __obf_7ef864bf74ca829e []color.Color) {
	__obf_345136d45eb8c039 := __obf_30bd90aa2dac79ac.Bounds().Max.X / 10
	__obf_9d198b0ed61e1439 := __obf_345136d45eb8c039 * 9
	__obf_d72c36055fc5f622 := __obf_30bd90aa2dac79ac.Bounds().Max.Y / 3
	for __obf_c599f3c3680a56f5 := range __obf_9df9b41f9f317bf5 {
		__obf_a6c7bec382cf9cd9 := image.Point{X: rand.IntN(__obf_345136d45eb8c039), Y: rand.IntN(__obf_d72c36055fc5f622)}
		__obf_298d993dc8e60914 := image.Point{X: rand.IntN(__obf_345136d45eb8c039) + __obf_9d198b0ed61e1439, Y: rand.IntN(__obf_d72c36055fc5f622)}

		if __obf_c599f3c3680a56f5%2 == 0 {
			__obf_a6c7bec382cf9cd9.
				Y = rand.IntN(__obf_d72c36055fc5f622) + __obf_d72c36055fc5f622*2
			__obf_298d993dc8e60914.
				Y = rand.IntN(__obf_d72c36055fc5f622)
		} else {
			__obf_a6c7bec382cf9cd9.
				Y = rand.IntN(__obf_d72c36055fc5f622) + __obf_d72c36055fc5f622*(__obf_c599f3c3680a56f5%2)
			__obf_298d993dc8e60914.
				Y = rand.IntN(__obf_d72c36055fc5f622) + __obf_d72c36055fc5f622*2
		}
		__obf_f77412067062c4da := util.RandColor(__obf_7ef864bf74ca829e)
		__obf_30bd90aa2dac79ac.
			//co.A = uint8(0xee)
			DrawBeeline(__obf_a6c7bec382cf9cd9, __obf_298d993dc8e60914, __obf_f77412067062c4da)
	}
}

func (__obf_2a577d1bd2f12f6a *Captcha) DrawDotImage(__obf_41af80e859213b9c DotType, __obf_d179dd23c6d2014b *types.Dot, __obf_b6321ad01abb7fde *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_aa90d8e6b96cb78e, _ := util.ParseHexColor(__obf_d179dd23c6d2014b.PrimaryColor)
	__obf_aa90d8e6b96cb78e.
		A = util.FormatAlpha(__obf_b6321ad01abb7fde.Alpha)

	var __obf_c8074fb0da8cfe92 image.Image
	var __obf_573fb4e2cef8adee error
	if __obf_41af80e859213b9c == Shape {
		__obf_c8074fb0da8cfe92, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawShapeImage(__obf_d179dd23c6d2014b, __obf_aa90d8e6b96cb78e)
	} else {
		__obf_c8074fb0da8cfe92, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawStringImage(__obf_d179dd23c6d2014b, __obf_aa90d8e6b96cb78e)
	}
	if __obf_573fb4e2cef8adee != nil {
		return nil, nil, __obf_573fb4e2cef8adee
	}
	__obf_b2f752595fa385ea, _ := util.ParseHexColor(__obf_aa810c905dac1551)
	__obf_9ba2ff5f86cfd090 := helper.CreateNRGBACanvas(__obf_d179dd23c6d2014b.Size+10, __obf_d179dd23c6d2014b.Size+10, true)
	if __obf_b6321ad01abb7fde.ShowShadow {
		var __obf_6cc9cbfc093bfd81 *types.NRGBA
		if __obf_41af80e859213b9c == Shape {
			__obf_6cc9cbfc093bfd81, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawShapeImage(__obf_d179dd23c6d2014b, __obf_b2f752595fa385ea)
		} else {
			__obf_6cc9cbfc093bfd81, __obf_573fb4e2cef8adee = __obf_2a577d1bd2f12f6a.DrawStringImage(__obf_d179dd23c6d2014b, __obf_b2f752595fa385ea)
		}
		if __obf_573fb4e2cef8adee != nil {
			return nil, nil, __obf_573fb4e2cef8adee
		}
		__obf_3b04eac8b24f91d4 := __obf_b6321ad01abb7fde.ShadowPoint.X
		__obf_43ce8acbc8280fbb := __obf_b6321ad01abb7fde.ShadowPoint.Y
		draw.Draw(__obf_9ba2ff5f86cfd090, __obf_6cc9cbfc093bfd81.Bounds(), __obf_6cc9cbfc093bfd81, image.Point{X: __obf_3b04eac8b24f91d4, Y: __obf_43ce8acbc8280fbb}, draw.Over)
	}
	draw.Draw(__obf_9ba2ff5f86cfd090, __obf_c8074fb0da8cfe92.Bounds(), __obf_c8074fb0da8cfe92, image.Point{}, draw.Over)
	__obf_9ba2ff5f86cfd090.
		Rotate(__obf_d179dd23c6d2014b.Angle, false)
	__obf_158d50f7190d9c0f := __obf_9ba2ff5f86cfd090.CalcMarginBlankArea()

	return __obf_9ba2ff5f86cfd090,

		// DrawStringImage draws a text image
		// params:
		//   - dot: Draw dot
		//   - textColor: Text color
		//
		// returns:
		//   - types.NRGBA: Drawn text image
		//   - error: Error information
		__obf_158d50f7190d9c0f, nil
}

func (__obf_2a577d1bd2f12f6a *Captcha) DrawStringImage(__obf_d179dd23c6d2014b *types.Dot, __obf_9bcd13c76f268b29 color.Color) (*types.NRGBA, error) {
	__obf_9ba2ff5f86cfd090 := helper.CreateNRGBACanvas(__obf_d179dd23c6d2014b.Size+10, __obf_d179dd23c6d2014b.Size+10, true)
	__obf_87e8685cf28e17a7 := fixed.P(12, __obf_d179dd23c6d2014b.Size-5)
	if util.IsChineseChar(__obf_d179dd23c6d2014b.Text) {
		__obf_87e8685cf28e17a7 = fixed.P(10, __obf_d179dd23c6d2014b.Size)
	}
	__obf_b6321ad01abb7fde := &types.DrawStringParam{
		Color:       __obf_9bcd13c76f268b29,
		FontSize:    __obf_d179dd23c6d2014b.FontSize,
		Size:        __obf_d179dd23c6d2014b.Size,
		FontDPI:     __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.FontDPI,
		FontHinting: __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.FontHinting,
		Text:        __obf_d179dd23c6d2014b.Text,
	}
	var __obf_573fb4e2cef8adee error
	__obf_b6321ad01abb7fde.
		Font, __obf_573fb4e2cef8adee = assets.PickFont()
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}
	__obf_573fb4e2cef8adee = helper.DrawString(__obf_9ba2ff5f86cfd090, __obf_b6321ad01abb7fde, __obf_87e8685cf28e17a7)
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}

	return __obf_9ba2ff5f86cfd090, nil
}

func (__obf_2a577d1bd2f12f6a *Captcha) DrawShapeImage(__obf_d179dd23c6d2014b *types.Dot, __obf_aa90d8e6b96cb78e color.Color) (*types.NRGBA, error) {
	__obf_1a790c0701615d48, __obf_74079feca9e7735d, __obf_22c70901828afd12, __obf_9a0006bb2706ddf0 := __obf_aa90d8e6b96cb78e.RGBA()

	var __obf_7ddf2f2f245fc198 = []color.RGBA{
		{R: uint8(__obf_1a790c0701615d48), G: uint8(__obf_74079feca9e7735d), B: uint8(__obf_22c70901828afd12), A: uint8(__obf_9a0006bb2706ddf0)},
	}
	__obf_02ceaa5d68e46dc9 := helper.CreateNRGBACanvas(__obf_d179dd23c6d2014b.Size+10, __obf_d179dd23c6d2014b.Size+10, true)
	var __obf_e0a70f9b0e5e3710 image.Rectangle
	var __obf_00aaa760d6f92456 image.Image
	if __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.UseRGBA {
		__obf_9ba2ff5f86cfd090 := helper.CreateNRGBACanvas(__obf_d179dd23c6d2014b.Size+10, __obf_d179dd23c6d2014b.Size+10, true)
		draw.BiLinear.Scale(__obf_9ba2ff5f86cfd090, __obf_9ba2ff5f86cfd090.Bounds(), __obf_d179dd23c6d2014b.Shape, __obf_d179dd23c6d2014b.Shape.Bounds(), draw.Over, nil)
		__obf_e0a70f9b0e5e3710 = __obf_9ba2ff5f86cfd090.Bounds()
		__obf_00aaa760d6f92456 = __obf_9ba2ff5f86cfd090
	} else {
		__obf_9ba2ff5f86cfd090 := helper.CreatePaletteCanvas(__obf_d179dd23c6d2014b.Size+10, __obf_d179dd23c6d2014b.Size+10, __obf_7ddf2f2f245fc198)
		draw.BiLinear.Scale(__obf_9ba2ff5f86cfd090, __obf_9ba2ff5f86cfd090.Bounds(), __obf_d179dd23c6d2014b.Shape, __obf_d179dd23c6d2014b.Shape.Bounds(), draw.Over, nil)
		__obf_e0a70f9b0e5e3710 = __obf_9ba2ff5f86cfd090.Bounds()
		__obf_00aaa760d6f92456 = __obf_9ba2ff5f86cfd090
	}

	draw.Draw(__obf_02ceaa5d68e46dc9, __obf_e0a70f9b0e5e3710, __obf_00aaa760d6f92456, image.Point{}, draw.Over)

	return __obf_02ceaa5d68e46dc9, nil
}

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_4dfb9c5752c4ac93() (*CaptchaData, error) {

	var (
		__obf_3430be58c123f6fe, __obf_39f51f543b42bf8c, __obf_308bb01dbe10110a []*types.Dot
		__obf_c1eb4ba2a2de5922                                                 []image.Image
		__obf_4897430c3c4dd92e                                                 []image.Image
		__obf_f7894dfc7807a63d, __obf_feb6cb098a8949dd                         image.Image
		__obf_573fb4e2cef8adee                                                 error
	)
	__obf_4897430c3c4dd92e, __obf_573fb4e2cef8adee = assets.GetShapes()
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}
	__obf_bed14be106dbdf63 := util.RandInt(__obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.LenRange.Min, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.LenRange.Max)
	__obf_4897430c3c4dd92e = util.PickN(__obf_4897430c3c4dd92e, __obf_bed14be106dbdf63)
	__obf_3430be58c123f6fe = __obf_e0f9d59d598c10d8.__obf_f27daf8a122156d1(__obf_4897430c3c4dd92e, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.Size, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.SizeRange, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.DotPadding)
	__obf_368909af54e7402d := rand.Perm(len(__obf_3430be58c123f6fe))
	__obf_2dd1cce6a15ec1ad := util.RandInt(__obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.VerifyLenRange.Min, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.VerifyLenRange.Max)
	__obf_308bb01dbe10110a = make([]*types.Dot, __obf_2dd1cce6a15ec1ad)
	for __obf_c599f3c3680a56f5, __obf_b651bf23cb9d28e6 := range __obf_368909af54e7402d {
		if __obf_c599f3c3680a56f5 >= __obf_2dd1cce6a15ec1ad {
			break
		}
		__obf_d179dd23c6d2014b := __obf_3430be58c123f6fe[__obf_b651bf23cb9d28e6]
		__obf_d179dd23c6d2014b.
			Index = __obf_c599f3c3680a56f5
		__obf_308bb01dbe10110a[__obf_c599f3c3680a56f5] = __obf_d179dd23c6d2014b
		__obf_c1eb4ba2a2de5922 = append(__obf_c1eb4ba2a2de5922, __obf_308bb01dbe10110a[__obf_c599f3c3680a56f5].Shape)
	}
	__obf_39f51f543b42bf8c = __obf_e0f9d59d598c10d8.__obf_f27daf8a122156d1(__obf_c1eb4ba2a2de5922, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.Size, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.SizeRange, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.DotPadding)
	__obf_f7894dfc7807a63d, __obf_573fb4e2cef8adee = __obf_e0f9d59d598c10d8.__obf_4845e52a61694361(Shape, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.Size, __obf_3430be58c123f6fe)
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}
	__obf_feb6cb098a8949dd, __obf_573fb4e2cef8adee = __obf_e0f9d59d598c10d8.__obf_b56493cc702de031(Shape, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.Size, __obf_39f51f543b42bf8c)
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}

	return &CaptchaData{__obf_7f68597a0d34a028: __obf_308bb01dbe10110a, __obf_f7894dfc7807a63d: types.NewJPEGImage(__obf_f7894dfc7807a63d), __obf_feb6cb098a8949dd: types.NewPNGImage(__obf_feb6cb098a8949dd)}, nil
}

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_591266b39896d5bb() (*CaptchaData, error) {
	__obf_bed14be106dbdf63 := util.RandInt(__obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.LenRange.Min, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.LenRange.Max)
	__obf_decf9a8056b8d5f1 := util.PickN(assets.GetChineseChars(), __obf_bed14be106dbdf63)

	var __obf_7f68597a0d34a028, __obf_39f51f543b42bf8c, __obf_308bb01dbe10110a []*types.Dot
	var __obf_e92d814917aefe2c []string
	var __obf_f7894dfc7807a63d, __obf_feb6cb098a8949dd image.Image
	__obf_7f68597a0d34a028 = __obf_e0f9d59d598c10d8.__obf_62c3b406ec8197bb(__obf_decf9a8056b8d5f1, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.Size, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.SizeRange, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.DotPadding)
	__obf_368909af54e7402d := rand.Perm(len(__obf_7f68597a0d34a028))
	__obf_2dd1cce6a15ec1ad := util.RandInt(__obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.VerifyLenRange.Min, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.VerifyLenRange.Max)
	__obf_308bb01dbe10110a = make([]*types.Dot, __obf_2dd1cce6a15ec1ad)
	for __obf_c599f3c3680a56f5, __obf_b651bf23cb9d28e6 := range __obf_368909af54e7402d {
		if __obf_c599f3c3680a56f5 >= __obf_2dd1cce6a15ec1ad {
			break
		}
		__obf_d179dd23c6d2014b := __obf_7f68597a0d34a028[__obf_b651bf23cb9d28e6]
		__obf_d179dd23c6d2014b.
			Index = __obf_c599f3c3680a56f5
		__obf_308bb01dbe10110a[__obf_c599f3c3680a56f5] = __obf_d179dd23c6d2014b
		__obf_e92d814917aefe2c = append(__obf_e92d814917aefe2c, __obf_308bb01dbe10110a[__obf_c599f3c3680a56f5].Text)
	}
	__obf_39f51f543b42bf8c = __obf_e0f9d59d598c10d8.__obf_62c3b406ec8197bb(__obf_e92d814917aefe2c, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.Size, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.SizeRange, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.DotPadding)
	var __obf_573fb4e2cef8adee error
	__obf_f7894dfc7807a63d, __obf_573fb4e2cef8adee = __obf_e0f9d59d598c10d8.__obf_4845e52a61694361(Text, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.Size, __obf_7f68597a0d34a028)
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}
	__obf_feb6cb098a8949dd, __obf_573fb4e2cef8adee = __obf_e0f9d59d598c10d8.__obf_b56493cc702de031(Text, __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Secondary.Size, __obf_39f51f543b42bf8c)
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}

	return &CaptchaData{__obf_7f68597a0d34a028: __obf_308bb01dbe10110a, __obf_f7894dfc7807a63d: types.NewJPEGImage(__obf_f7894dfc7807a63d), __obf_feb6cb098a8949dd: types.NewPNGImage(__obf_feb6cb098a8949dd)}, nil
}

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_f27daf8a122156d1(__obf_4897430c3c4dd92e []image.Image, __obf_da05e38f56e09ea3 types.Size, __obf_58cc87a0e6218150 types.Range, __obf_f0d5d97d0ec7f54d int) []*types.Dot {
	__obf_bed14be106dbdf63 := len(__obf_4897430c3c4dd92e)
	var __obf_7f68597a0d34a028 = make([]*types.Dot, __obf_bed14be106dbdf63)
	__obf_a4543061eb0a2e0b := __obf_da05e38f56e09ea3.Width
	__obf_c69f0a9feb9621f7 := __obf_da05e38f56e09ea3.Height
	if __obf_f0d5d97d0ec7f54d > 0 {
		__obf_a4543061eb0a2e0b -= __obf_f0d5d97d0ec7f54d
		__obf_c69f0a9feb9621f7 -= __obf_f0d5d97d0ec7f54d
	}

	for __obf_c599f3c3680a56f5, __obf_1ed0b3f2e37769e4 := range __obf_4897430c3c4dd92e {
		__obf_8a0dc659b560ec11 := __obf_e0f9d59d598c10d8.__obf_8a0dc659b560ec11()
		__obf_a8ea7a0c45d7ec2b := util.PickN(__obf_fa28e9463cf9775f, 1)[0]
		__obf_14c69cace08e2e33 := util.PickN(__obf_312459111d69d06c, 1)[0]
		__obf_3694f5a7bb9a8fc6 := util.RandInt(__obf_58cc87a0e6218150.Min, __obf_58cc87a0e6218150.Max)
		__obf_0b5eeb13b8b2d52f := 10
		__obf_06d89f453b8f83e3 := __obf_a4543061eb0a2e0b / __obf_bed14be106dbdf63
		__obf_3ee778a19c1f35b5 := math.Abs(float64(__obf_06d89f453b8f83e3) - float64(__obf_3694f5a7bb9a8fc6))
		__obf_321344ba59c9cbff := (__obf_c599f3c3680a56f5 * __obf_06d89f453b8f83e3) + util.RandInt(0, int(max(__obf_3ee778a19c1f35b5, 1)))
		__obf_b57e5984317689bf := util.RandInt(__obf_0b5eeb13b8b2d52f, __obf_c69f0a9feb9621f7+__obf_3694f5a7bb9a8fc6)
		__obf_9a77cb45e2ebdb2d := int(min(max(float64(__obf_321344ba59c9cbff), float64(__obf_0b5eeb13b8b2d52f)), float64(__obf_a4543061eb0a2e0b-__obf_0b5eeb13b8b2d52f-(__obf_f0d5d97d0ec7f54d*2))))
		__obf_d72c36055fc5f622 := int(min(max(float64(__obf_b57e5984317689bf), float64(__obf_3694f5a7bb9a8fc6+__obf_0b5eeb13b8b2d52f)), float64(__obf_c69f0a9feb9621f7+(__obf_3694f5a7bb9a8fc6/2)-(__obf_f0d5d97d0ec7f54d*2))))
		__obf_7f68597a0d34a028[__obf_c599f3c3680a56f5] = &types.Dot{
			Index:          __obf_c599f3c3680a56f5,
			X:              __obf_9a77cb45e2ebdb2d,
			Y:              __obf_d72c36055fc5f622 - __obf_3694f5a7bb9a8fc6,
			FontSize:       __obf_3694f5a7bb9a8fc6,
			Size:           __obf_3694f5a7bb9a8fc6,
			Angle:          __obf_8a0dc659b560ec11,
			PrimaryColor:   __obf_a8ea7a0c45d7ec2b,
			SecondaryColor: __obf_14c69cace08e2e33,
			Shape:          __obf_1ed0b3f2e37769e4,
		}
	}

	return __obf_7f68597a0d34a028
}

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_62c3b406ec8197bb(__obf_decf9a8056b8d5f1 []string, __obf_da05e38f56e09ea3 types.Size, __obf_55e4d1571152e66a types.Range, __obf_f0d5d97d0ec7f54d int) []*types.Dot {
	__obf_bed14be106dbdf63 := len(__obf_decf9a8056b8d5f1)
	var __obf_7f68597a0d34a028 = make([]*types.Dot, __obf_bed14be106dbdf63)
	__obf_a4543061eb0a2e0b := __obf_da05e38f56e09ea3.Width
	__obf_c69f0a9feb9621f7 := __obf_da05e38f56e09ea3.Height
	if __obf_f0d5d97d0ec7f54d > 0 {
		__obf_a4543061eb0a2e0b -= __obf_f0d5d97d0ec7f54d
		__obf_c69f0a9feb9621f7 -= __obf_f0d5d97d0ec7f54d
	}

	for __obf_c599f3c3680a56f5, __obf_1ed0b3f2e37769e4 := range __obf_decf9a8056b8d5f1 {
		__obf_8a0dc659b560ec11 := __obf_e0f9d59d598c10d8.__obf_8a0dc659b560ec11()
		__obf_a8ea7a0c45d7ec2b := util.PickN(__obf_fa28e9463cf9775f, 1)[0]
		__obf_14c69cace08e2e33 := util.PickN(__obf_312459111d69d06c, 1)[0]
		__obf_3694f5a7bb9a8fc6 := util.RandInt(__obf_55e4d1571152e66a.Min, __obf_55e4d1571152e66a.Max)
		__obf_2577d888331c3dcb := __obf_3694f5a7bb9a8fc6
		__obf_2754264a11a68656 := __obf_3694f5a7bb9a8fc6

		if util.LenChineseChar(__obf_1ed0b3f2e37769e4) > 1 {
			__obf_2754264a11a68656 = __obf_3694f5a7bb9a8fc6 * util.LenChineseChar(__obf_1ed0b3f2e37769e4)

			if __obf_8a0dc659b560ec11 > 0 {
				__obf_590069478bfbe405 := __obf_2754264a11a68656 - __obf_3694f5a7bb9a8fc6
				__obf_704cdb6bd9ed14a2 := __obf_8a0dc659b560ec11 % 90
				__obf_ee0817a6b06e56d5 := float64(__obf_590069478bfbe405) / 90
				__obf_89aa6f66c78c925b := max(float64(__obf_704cdb6bd9ed14a2)*__obf_ee0817a6b06e56d5, 1)
				__obf_2577d888331c3dcb = __obf_2577d888331c3dcb + int(__obf_89aa6f66c78c925b)
				__obf_2754264a11a68656 = __obf_2754264a11a68656 + int(__obf_89aa6f66c78c925b)
			}
		}
		__obf_0b5eeb13b8b2d52f := 10
		__obf_06d89f453b8f83e3 := __obf_a4543061eb0a2e0b / __obf_bed14be106dbdf63
		__obf_3ee778a19c1f35b5 := math.Abs(float64(__obf_06d89f453b8f83e3) - float64(__obf_2754264a11a68656))
		__obf_321344ba59c9cbff := (__obf_c599f3c3680a56f5 * __obf_06d89f453b8f83e3) + util.RandInt(0, int(max(__obf_3ee778a19c1f35b5, 1)))
		__obf_b57e5984317689bf := util.RandInt(__obf_0b5eeb13b8b2d52f, __obf_c69f0a9feb9621f7+__obf_2577d888331c3dcb)
		__obf_9a77cb45e2ebdb2d := int(min(max(float64(__obf_321344ba59c9cbff), float64(__obf_0b5eeb13b8b2d52f)), float64(__obf_a4543061eb0a2e0b-__obf_0b5eeb13b8b2d52f-(__obf_f0d5d97d0ec7f54d*2))))
		__obf_d72c36055fc5f622 := int(min(max(float64(__obf_b57e5984317689bf), float64(__obf_2577d888331c3dcb+__obf_0b5eeb13b8b2d52f)), float64(__obf_c69f0a9feb9621f7+(__obf_2577d888331c3dcb/2)-(__obf_f0d5d97d0ec7f54d*2))))
		__obf_7f68597a0d34a028[__obf_c599f3c3680a56f5] = &types.Dot{
			Index:          __obf_c599f3c3680a56f5,
			X:              __obf_9a77cb45e2ebdb2d,
			Y:              __obf_d72c36055fc5f622 - __obf_2577d888331c3dcb,
			FontSize:       __obf_3694f5a7bb9a8fc6,
			Size:           max(__obf_2754264a11a68656, __obf_2577d888331c3dcb),
			Angle:          __obf_8a0dc659b560ec11,
			PrimaryColor:   __obf_a8ea7a0c45d7ec2b,
			SecondaryColor: __obf_14c69cace08e2e33,
			Text:           __obf_1ed0b3f2e37769e4,
		}
	}

	return __obf_7f68597a0d34a028
}

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_4845e52a61694361(__obf_41af80e859213b9c DotType, __obf_55e4d1571152e66a types.Size, __obf_7f68597a0d34a028 []*types.Dot) (image.Image, error) {
	__obf_6429bd294169c216 := &DrawImageParams{
		Width:       __obf_55e4d1571152e66a.Width,
		Height:      __obf_55e4d1571152e66a.Height,
		Alpha:       __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.Alpha,
		FontHinting: __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.FontHinting,
		Dots:        __obf_7f68597a0d34a028,
		ShowShadow:  __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.ShowShadow,
		ShadowPoint: __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.ShadowPoint,
	}
	var __obf_573fb4e2cef8adee error
	__obf_6429bd294169c216.
		Background, __obf_573fb4e2cef8adee = assets.PickBgImage()
	if __obf_573fb4e2cef8adee != nil {
		return nil, __obf_573fb4e2cef8adee
	}

	return __obf_e0f9d59d598c10d8.DrawWithNRGBA(__obf_41af80e859213b9c, __obf_6429bd294169c216)
}

func (__obf_2a577d1bd2f12f6a *Captcha) __obf_b56493cc702de031(__obf_41af80e859213b9c DotType, __obf_55e4d1571152e66a types.Size, __obf_7f68597a0d34a028 []*types.Dot) (image.Image, error) {
	var (
		__obf_e045515eac38ac51 =
		// err      error
		make([]*types.Dot, 0, len(__obf_7f68597a0d34a028))
	)
	__obf_a4543061eb0a2e0b := __obf_55e4d1571152e66a.Width / len(__obf_7f68597a0d34a028)

	for __obf_c599f3c3680a56f5, __obf_1ed0b3f2e37769e4 := range __obf_7f68597a0d34a028 {
		__obf_bed14be106dbdf63 := 1
		if __obf_41af80e859213b9c == Text {
			__obf_bed14be106dbdf63 = len(__obf_1ed0b3f2e37769e4.Text)
		}
		__obf_1ed0b3f2e37769e4.
			X = int(max(float64(__obf_a4543061eb0a2e0b*__obf_c599f3c3680a56f5+__obf_a4543061eb0a2e0b/__obf_1ed0b3f2e37769e4.Size), 8))
		__obf_65bebd1310c67b45 := max(1, __obf_55e4d1571152e66a.Height/16*__obf_bed14be106dbdf63)
		__obf_1ed0b3f2e37769e4.
			Y = __obf_55e4d1571152e66a.Height/2 + __obf_1ed0b3f2e37769e4.FontSize/2 - rand.IntN(__obf_65bebd1310c67b45)
		__obf_e045515eac38ac51 = append(__obf_e045515eac38ac51, __obf_1ed0b3f2e37769e4)
	}
	__obf_b6321ad01abb7fde := &DrawImageParams{
		Width:             __obf_55e4d1571152e66a.Width,
		Height:            __obf_55e4d1571152e66a.Height,
		Dots:              __obf_e045515eac38ac51,
		BackgroundDistort: __obf_2a577d1bd2f12f6a.__obf_2cbeb47c02dcbfcd(__obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.Secondary.BgDistort),
		BgCircles:         __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.Secondary.BgCircles,
		BgSlimLines:       __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_2d2372316a3c806b []color.Color
	for _, __obf_19097d45d5541397 := range __obf_312459111d69d06c {
		__obf_f77412067062c4da, _ := util.ParseHexColor(__obf_19097d45d5541397)
		__obf_2d2372316a3c806b = append(__obf_2d2372316a3c806b, __obf_f77412067062c4da)
	}

	var __obf_09570dc7552c46d3 []color.Color
	for _, __obf_f77412067062c4da := range __obf_312459111d69d06c {
		__obf_709b6d468bf64292, _ := util.ParseHexColor(__obf_f77412067062c4da)
		__obf_09570dc7552c46d3 = append(__obf_09570dc7552c46d3, __obf_709b6d468bf64292)
	}

	if __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.UseRGBA || __obf_2a577d1bd2f12f6a.__obf_03afc9d99f21e175.Secondary.NonDeformAble {
		return __obf_2a577d1bd2f12f6a.DrawWithNRGBA2(__obf_41af80e859213b9c, __obf_b6321ad01abb7fde, __obf_2d2372316a3c806b, __obf_09570dc7552c46d3)
	}
	return __obf_2a577d1bd2f12f6a.DrawWithPalette(__obf_41af80e859213b9c, __obf_b6321ad01abb7fde, __obf_2d2372316a3c806b, __obf_09570dc7552c46d3)
}

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_2cbeb47c02dcbfcd(__obf_cce61d3578e85850 int) int {
	switch __obf_cce61d3578e85850 {
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

func (__obf_e0f9d59d598c10d8 *Captcha) __obf_8a0dc659b560ec11() int {
	__obf_0238d3d06d84d05f := __obf_e0f9d59d598c10d8.__obf_03afc9d99f21e175.Primary.AnglePosRange
	__obf_b6f8bb0163f50f7d := rand.IntN(len(__obf_0238d3d06d84d05f))
	if __obf_b6f8bb0163f50f7d < 0 {
		return 0
	}
	__obf_312e5fc99fcf145d := __obf_0238d3d06d84d05f[__obf_b6f8bb0163f50f7d]
	__obf_f70b3e3095e030e2 := util.RandInt(__obf_312e5fc99fcf145d.Min, __obf_312e5fc99fcf145d.Max)

	return __obf_f70b3e3095e030e2
}
