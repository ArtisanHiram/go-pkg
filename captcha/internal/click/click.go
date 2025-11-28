package __obf_b2344bb9128f17e6

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
	__obf_126d1b2696ab5fed *Options
}

func (__obf_0e570eb08dbdd435 *Captcha) SetOptions(__obf_126d1b2696ab5fed *Options) {
	__obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed = __obf_126d1b2696ab5fed
}

func NewCaptcha(__obf_126d1b2696ab5fed *Options) *Captcha {
	if __obf_126d1b2696ab5fed == nil {
		return nil
	}
	return &Captcha{__obf_126d1b2696ab5fed}
}

func __obf_524140d8ce6def5b(__obf_126d1b2696ab5fed *Options) DotType {
	__obf_13b49cad6b38c878 := DotType(rand.IntN(2))
	if __obf_13b49cad6b38c878 == Shape {
		__obf_126d1b2696ab5fed.UseRGBA = false
		__obf_126d1b2696ab5fed.Secondary.BgDistort = types.DistortLevel1
		__obf_126d1b2696ab5fed.Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_126d1b2696ab5fed.Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_126d1b2696ab5fed.UseRGBA = true
		__obf_126d1b2696ab5fed.Secondary.BgDistort = types.DistortLevel4
		__obf_126d1b2696ab5fed.Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_126d1b2696ab5fed.Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_13b49cad6b38c878
}

func (__obf_0e570eb08dbdd435 *Captcha) DrawWithNRGBA(__obf_894388cdf1835fdb DotType, __obf_b409aea2e7dfafc6 *DrawImageParams) (image.Image, error) {
	__obf_8c02fb9797f8ce7a := __obf_b409aea2e7dfafc6.Dots
	__obf_06b112bbb1f58b48 := helper.CreateNRGBACanvas(__obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height, true)

	for _, __obf_a73552353b93ea58 := range __obf_8c02fb9797f8ce7a {

		__obf_903f882c7fdefd44, __obf_a88b3f75204ef50b, __obf_3afd7555fbd2aaf7 := __obf_0e570eb08dbdd435.DrawDotImage(__obf_894388cdf1835fdb, __obf_a73552353b93ea58, __obf_b409aea2e7dfafc6)
		if __obf_3afd7555fbd2aaf7 != nil {
			return nil, __obf_3afd7555fbd2aaf7
		}
		__obf_c2f9726a226c8aee := __obf_a88b3f75204ef50b.MinX
		__obf_7632b95a9e2bd000 := __obf_a88b3f75204ef50b.MaxX
		__obf_1bf05473b01b5f8f := __obf_a88b3f75204ef50b.MinY
		__obf_b609cbb811a15da4 := __obf_a88b3f75204ef50b.MaxY

		__obf_ba3c808c322fb14d := __obf_7632b95a9e2bd000 - __obf_c2f9726a226c8aee
		__obf_d355b65d00dab8b0 := __obf_b609cbb811a15da4 - __obf_1bf05473b01b5f8f

		draw.Draw(__obf_06b112bbb1f58b48, image.Rect(__obf_a73552353b93ea58.X, __obf_a73552353b93ea58.Y, __obf_a73552353b93ea58.X+__obf_ba3c808c322fb14d, __obf_a73552353b93ea58.Y+__obf_d355b65d00dab8b0), __obf_903f882c7fdefd44, image.Point{X: __obf_c2f9726a226c8aee, Y: __obf_1bf05473b01b5f8f}, draw.Over)

	}

	__obf_65fc1d4866918f94 := __obf_b409aea2e7dfafc6.Background
	__obf_8de2475006027da3 := __obf_06b112bbb1f58b48.Bounds()
	__obf_c0329d72b1b272fd := helper.CreateNRGBACanvas(__obf_8de2475006027da3.Dx(), __obf_8de2475006027da3.Dx(), true)
	__obf_4766956ca8f27e48 := util.RangCutImagePos(__obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height, __obf_65fc1d4866918f94)
	draw.Draw(__obf_c0329d72b1b272fd, __obf_8de2475006027da3, __obf_65fc1d4866918f94, __obf_4766956ca8f27e48, draw.Src)
	draw.Draw(__obf_c0329d72b1b272fd, __obf_06b112bbb1f58b48.Bounds(), __obf_06b112bbb1f58b48, image.Point{}, draw.Over)
	__obf_c0329d72b1b272fd.SubImage(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height))
	return __obf_c0329d72b1b272fd, nil
}

func (__obf_0e570eb08dbdd435 *Captcha) DrawWithPalette(__obf_894388cdf1835fdb DotType, __obf_b409aea2e7dfafc6 *DrawImageParams, __obf_5c2e38b0fea02c9b []color.Color, __obf_0a7e2714f6db07cc []color.Color) (image.Image, error) {
	__obf_8c02fb9797f8ce7a := __obf_b409aea2e7dfafc6.Dots

	__obf_4f89b68e570d1e7c := make([]color.Color, 0, len(__obf_0a7e2714f6db07cc))
	for _, __obf_3a29ccd9a5e023ad := range __obf_0a7e2714f6db07cc {
		__obf_2166b3491b2ba98f, __obf_3add3ddb64c8a0a2, __obf_8de2475006027da3, _ := __obf_3a29ccd9a5e023ad.RGBA()
		__obf_b5494ff4aae8a404 := util.FormatAlpha(__obf_b409aea2e7dfafc6.ThumbDisturbAlpha)
		__obf_4f89b68e570d1e7c = append(__obf_4f89b68e570d1e7c, color.RGBA{R: uint8(__obf_2166b3491b2ba98f), G: uint8(__obf_3add3ddb64c8a0a2), B: uint8(__obf_8de2475006027da3), A: __obf_b5494ff4aae8a404})
	}

	var __obf_a2f4a25139ddb3d0 = make([]color.Color, 0, len(__obf_5c2e38b0fea02c9b)+len(__obf_4f89b68e570d1e7c)+1)
	__obf_a2f4a25139ddb3d0 = append(__obf_a2f4a25139ddb3d0, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_a2f4a25139ddb3d0 = append(__obf_a2f4a25139ddb3d0, __obf_5c2e38b0fea02c9b...)
	__obf_a2f4a25139ddb3d0 = append(__obf_a2f4a25139ddb3d0, __obf_4f89b68e570d1e7c...)

	__obf_06b112bbb1f58b48 := types.NewPalette(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height), __obf_a2f4a25139ddb3d0)
	if __obf_b409aea2e7dfafc6.BgCircles > 0 {
		__obf_0e570eb08dbdd435.__obf_d001f08fb63696f8(__obf_06b112bbb1f58b48, __obf_b409aea2e7dfafc6.BgCircles, 1, __obf_4f89b68e570d1e7c)
	}
	if __obf_b409aea2e7dfafc6.BgSlimLines > 0 {
		__obf_0e570eb08dbdd435.__obf_0f2c353dd575b9af(__obf_06b112bbb1f58b48, __obf_b409aea2e7dfafc6.BgSlimLines, __obf_4f89b68e570d1e7c)
	}

	for _, __obf_a73552353b93ea58 := range __obf_8c02fb9797f8ce7a {
		__obf_ad611a46c4a5252c, _ := util.ParseHexColor(__obf_a73552353b93ea58.PrimaryColor)
		var __obf_3afd7555fbd2aaf7 error
		if __obf_894388cdf1835fdb == Shape {
			var __obf_903f882c7fdefd44 *types.NRGBA
			__obf_903f882c7fdefd44, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawShapeImage(__obf_a73552353b93ea58, __obf_ad611a46c4a5252c)
			if __obf_3afd7555fbd2aaf7 != nil {
				return nil, __obf_3afd7555fbd2aaf7
			}

			__obf_903f882c7fdefd44.Rotate(__obf_a73552353b93ea58.Angle, false)
			__obf_59d5b7cda2cc03aa := __obf_06b112bbb1f58b48.Bounds()
			__obf_9f343eb325c91ccb := __obf_903f882c7fdefd44.Bounds()
			__obf_943605d96f200e05 := image.Point{X: __obf_59d5b7cda2cc03aa.Dx() - __obf_9f343eb325c91ccb.Dx(), Y: __obf_59d5b7cda2cc03aa.Dy() - __obf_9f343eb325c91ccb.Dy()}
			draw.Draw(__obf_06b112bbb1f58b48, image.Rect(__obf_a73552353b93ea58.X, __obf_943605d96f200e05.Y, __obf_943605d96f200e05.X+__obf_9f343eb325c91ccb.Dx(), __obf_943605d96f200e05.Y+__obf_9f343eb325c91ccb.Dy()), __obf_903f882c7fdefd44, image.Point{}, draw.Over)
		} else {
			__obf_04555369bb47c7d0 := fixed.P(__obf_a73552353b93ea58.X, __obf_a73552353b93ea58.Y)
			__obf_b409aea2e7dfafc6 := &types.DrawStringParam{
				Color:       __obf_ad611a46c4a5252c,
				FontSize:    __obf_a73552353b93ea58.FontSize,
				Size:        __obf_a73552353b93ea58.Size,
				FontDPI:     __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.FontDPI,
				FontHinting: __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.FontHinting,
				Text:        __obf_a73552353b93ea58.Text,
			}
			__obf_b409aea2e7dfafc6.Font, __obf_3afd7555fbd2aaf7 = assets.PickFont()
			if __obf_3afd7555fbd2aaf7 != nil {
				return nil, __obf_3afd7555fbd2aaf7
			}
			__obf_3afd7555fbd2aaf7 = helper.DrawString(__obf_06b112bbb1f58b48, __obf_b409aea2e7dfafc6, __obf_04555369bb47c7d0)
		}

		if __obf_3afd7555fbd2aaf7 != nil {
			return __obf_06b112bbb1f58b48, __obf_3afd7555fbd2aaf7
		}
	}

	if __obf_b409aea2e7dfafc6.Background != nil {
		__obf_65fc1d4866918f94 := __obf_b409aea2e7dfafc6.Background
		__obf_8de2475006027da3 := __obf_65fc1d4866918f94.Bounds()
		__obf_c0329d72b1b272fd := helper.CreateNRGBACanvas(__obf_8de2475006027da3.Dx(), __obf_8de2475006027da3.Dy(), true)
		__obf_4766956ca8f27e48 := util.RangCutImagePos(__obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height, __obf_65fc1d4866918f94)
		draw.Draw(__obf_c0329d72b1b272fd, __obf_8de2475006027da3, __obf_65fc1d4866918f94, __obf_4766956ca8f27e48, draw.Src)
		__obf_06b112bbb1f58b48.Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_c0329d72b1b272fd, __obf_06b112bbb1f58b48.Bounds(), __obf_06b112bbb1f58b48, image.Point{}, draw.Over)
		__obf_9352bdf4be21c20e := __obf_c0329d72b1b272fd.SubImage(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height)).(*image.NRGBA)
		return __obf_9352bdf4be21c20e, nil
	}

	if __obf_b409aea2e7dfafc6.BackgroundDistort > 0 {
		__obf_06b112bbb1f58b48.Distort(float64(util.RandInt(5, 10)), float64(__obf_b409aea2e7dfafc6.BackgroundDistort))
	}

	return __obf_06b112bbb1f58b48, nil

}

func (__obf_0e570eb08dbdd435 *Captcha) DrawWithNRGBA2(__obf_894388cdf1835fdb DotType, __obf_b409aea2e7dfafc6 *DrawImageParams, __obf_5c2e38b0fea02c9b []color.Color, __obf_0a7e2714f6db07cc []color.Color) (image.Image, error) {
	__obf_8c02fb9797f8ce7a := __obf_b409aea2e7dfafc6.Dots

	__obf_4f89b68e570d1e7c := make([]color.Color, 0, len(__obf_0a7e2714f6db07cc))
	for _, __obf_3a29ccd9a5e023ad := range __obf_0a7e2714f6db07cc {
		__obf_2166b3491b2ba98f, __obf_3add3ddb64c8a0a2, __obf_8de2475006027da3, _ := __obf_3a29ccd9a5e023ad.RGBA()
		__obf_b5494ff4aae8a404 := util.FormatAlpha(__obf_b409aea2e7dfafc6.ThumbDisturbAlpha)
		__obf_4f89b68e570d1e7c = append(__obf_4f89b68e570d1e7c, color.RGBA{R: uint8(__obf_2166b3491b2ba98f), G: uint8(__obf_3add3ddb64c8a0a2), B: uint8(__obf_8de2475006027da3), A: __obf_b5494ff4aae8a404})
	}

	var __obf_a2f4a25139ddb3d0 = make([]color.Color, 0, len(__obf_5c2e38b0fea02c9b)+len(__obf_4f89b68e570d1e7c)+1)
	__obf_a2f4a25139ddb3d0 = append(__obf_a2f4a25139ddb3d0, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_a2f4a25139ddb3d0 = append(__obf_a2f4a25139ddb3d0, __obf_5c2e38b0fea02c9b...)
	__obf_a2f4a25139ddb3d0 = append(__obf_a2f4a25139ddb3d0, __obf_4f89b68e570d1e7c...)

	__obf_92c01502f79d8aa1 := types.NewNRGBA(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height), true)
	if __obf_b409aea2e7dfafc6.Background != nil {
		__obf_65fc1d4866918f94 := __obf_b409aea2e7dfafc6.Background
		__obf_8de2475006027da3 := __obf_65fc1d4866918f94.Bounds()
		__obf_c0329d72b1b272fd := helper.CreateNRGBACanvas(__obf_8de2475006027da3.Dx(), __obf_8de2475006027da3.Dy(), true)
		__obf_4766956ca8f27e48 := util.RangCutImagePos(__obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height, __obf_65fc1d4866918f94)
		draw.Draw(__obf_c0329d72b1b272fd, __obf_8de2475006027da3, __obf_65fc1d4866918f94, __obf_4766956ca8f27e48, draw.Src)
		__obf_9352bdf4be21c20e := __obf_c0329d72b1b272fd.SubImage(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height)).(*image.NRGBA)
		draw.Draw(__obf_92c01502f79d8aa1, __obf_9352bdf4be21c20e.Bounds(), __obf_9352bdf4be21c20e, image.Point{}, draw.Over)
	}

	__obf_06b112bbb1f58b48 := types.NewPalette(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height), __obf_a2f4a25139ddb3d0)
	if __obf_b409aea2e7dfafc6.BgCircles > 0 {
		__obf_0e570eb08dbdd435.__obf_d001f08fb63696f8(__obf_06b112bbb1f58b48, __obf_b409aea2e7dfafc6.BgCircles, 1, __obf_4f89b68e570d1e7c)
	}
	if __obf_b409aea2e7dfafc6.BgSlimLines > 0 {
		__obf_0e570eb08dbdd435.__obf_0f2c353dd575b9af(__obf_06b112bbb1f58b48, __obf_b409aea2e7dfafc6.BgSlimLines, __obf_4f89b68e570d1e7c)
	}
	if __obf_b409aea2e7dfafc6.BackgroundDistort > 0 {
		__obf_06b112bbb1f58b48.Distort(float64(util.RandInt(5, 10)), float64(__obf_b409aea2e7dfafc6.BackgroundDistort))
	}

	__obf_36f8058c5ddbb62f := __obf_06b112bbb1f58b48.Bounds()
	__obf_ba3c808c322fb14d := __obf_36f8058c5ddbb62f.Dx() / len(__obf_8c02fb9797f8ce7a)
	for __obf_76ec865d65fb400e, __obf_a73552353b93ea58 := range __obf_8c02fb9797f8ce7a {
		__obf_ad611a46c4a5252c, _ := util.ParseHexColor(__obf_a73552353b93ea58.PrimaryColor)
		var __obf_3afd7555fbd2aaf7 error
		if __obf_894388cdf1835fdb == Shape {
			var __obf_903f882c7fdefd44 *types.NRGBA
			__obf_903f882c7fdefd44, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawShapeImage(__obf_a73552353b93ea58, __obf_ad611a46c4a5252c)
			if __obf_3afd7555fbd2aaf7 != nil {
				return nil, __obf_3afd7555fbd2aaf7
			}

			__obf_903f882c7fdefd44.Rotate(__obf_a73552353b93ea58.Angle, false)

			__obf_59d5b7cda2cc03aa := __obf_92c01502f79d8aa1.Bounds()
			__obf_9f343eb325c91ccb := __obf_903f882c7fdefd44.Bounds()
			__obf_943605d96f200e05 := image.Point{X: __obf_59d5b7cda2cc03aa.Dx() - __obf_9f343eb325c91ccb.Dx(), Y: __obf_59d5b7cda2cc03aa.Dy() - __obf_9f343eb325c91ccb.Dy()}
			draw.Draw(__obf_92c01502f79d8aa1, image.Rect(__obf_a73552353b93ea58.X, __obf_943605d96f200e05.Y, __obf_943605d96f200e05.X+__obf_9f343eb325c91ccb.Dx(), __obf_943605d96f200e05.Y+__obf_9f343eb325c91ccb.Dy()), __obf_903f882c7fdefd44, image.Point{}, draw.Over)
		} else {
			var __obf_4c029424cf0a53ef *types.NRGBA
			__obf_4c029424cf0a53ef, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawStringImage(__obf_a73552353b93ea58, __obf_ad611a46c4a5252c)
			if __obf_3afd7555fbd2aaf7 != nil {
				return nil, __obf_3afd7555fbd2aaf7
			}
			__obf_4c029424cf0a53ef.Rotate(__obf_a73552353b93ea58.Angle, false)

			__obf_a88b3f75204ef50b := __obf_4c029424cf0a53ef.CalcMarginBlankArea()
			__obf_c2f9726a226c8aee := __obf_a88b3f75204ef50b.MinX
			__obf_7632b95a9e2bd000 := __obf_a88b3f75204ef50b.MaxX
			__obf_1bf05473b01b5f8f := __obf_a88b3f75204ef50b.MinY
			__obf_b609cbb811a15da4 := __obf_a88b3f75204ef50b.MaxY
			__obf_4c029424cf0a53ef.SubImage(image.Rect(__obf_c2f9726a226c8aee, __obf_1bf05473b01b5f8f, __obf_7632b95a9e2bd000, __obf_b609cbb811a15da4))
			__obf_b5e82a98533f8553 := __obf_4c029424cf0a53ef.Bounds()

			__obf_7dc04ddca0db4183 := int(max(float64(__obf_ba3c808c322fb14d*__obf_76ec865d65fb400e+__obf_ba3c808c322fb14d/__obf_b5e82a98533f8553.Dx()), 8))
			__obf_4efaefa63fa6e566 := util.RandInt(1, __obf_36f8058c5ddbb62f.Dy()-__obf_b5e82a98533f8553.Dy()-4)

			draw.Draw(__obf_92c01502f79d8aa1, image.Rect(__obf_7dc04ddca0db4183, __obf_4efaefa63fa6e566, __obf_7dc04ddca0db4183+__obf_b5e82a98533f8553.Dx(), __obf_4efaefa63fa6e566+__obf_b5e82a98533f8553.Dy()), __obf_4c029424cf0a53ef, image.Point{X: __obf_b5e82a98533f8553.Min.X, Y: __obf_b5e82a98533f8553.Min.Y}, draw.Over)
		}

		if __obf_3afd7555fbd2aaf7 != nil {
			return __obf_92c01502f79d8aa1, __obf_3afd7555fbd2aaf7
		}
	}

	__obf_7dad0498e0b3b97d := types.NewNRGBA(image.Rect(0, 0, __obf_b409aea2e7dfafc6.Width, __obf_b409aea2e7dfafc6.Height), true)
	draw.Draw(__obf_7dad0498e0b3b97d, __obf_06b112bbb1f58b48.Bounds(), __obf_06b112bbb1f58b48, image.Point{}, draw.Over)
	draw.Draw(__obf_92c01502f79d8aa1, __obf_7dad0498e0b3b97d.Bounds(), __obf_7dad0498e0b3b97d, image.Point{}, draw.Over)
	return __obf_92c01502f79d8aa1, nil
}

func (__obf_0e570eb08dbdd435 *Captcha) __obf_d001f08fb63696f8(__obf_c0329d72b1b272fd *types.Palette, __obf_008bae6553da7b95, __obf_1ecd4dc8a508fb4c int, __obf_fd80fb07f7829c9c []color.Color) {
	__obf_3ebc662524920ea5 := __obf_c0329d72b1b272fd.Bounds().Max.X
	__obf_2fd045e9321f7b15 := __obf_c0329d72b1b272fd.Bounds().Max.Y
	for range __obf_008bae6553da7b95 {
		__obf_646dbac6c9d3b39a := util.RandColor(__obf_fd80fb07f7829c9c)
		//co.A = uint8(0xee)
		__obf_2166b3491b2ba98f := util.RandInt(1, __obf_1ecd4dc8a508fb4c)
		__obf_c0329d72b1b272fd.DrawCircle(util.RandInt(__obf_2166b3491b2ba98f, __obf_3ebc662524920ea5-__obf_2166b3491b2ba98f), util.RandInt(__obf_2166b3491b2ba98f, __obf_2fd045e9321f7b15-__obf_2166b3491b2ba98f), __obf_2166b3491b2ba98f, __obf_646dbac6c9d3b39a)
	}
}

func (__obf_0e570eb08dbdd435 *Captcha) __obf_0f2c353dd575b9af(__obf_c0329d72b1b272fd *types.Palette, __obf_702481c071f63cdd int, __obf_fd80fb07f7829c9c []color.Color) {
	__obf_8cff1b162870f5a8 := __obf_c0329d72b1b272fd.Bounds().Max.X / 10
	__obf_1c0b3e980dbbd1e1 := __obf_8cff1b162870f5a8 * 9
	__obf_bc574d6b33b94515 := __obf_c0329d72b1b272fd.Bounds().Max.Y / 3
	for __obf_76ec865d65fb400e := range __obf_702481c071f63cdd {
		__obf_f5cb084092927ec6 := image.Point{X: rand.IntN(__obf_8cff1b162870f5a8), Y: rand.IntN(__obf_bc574d6b33b94515)}
		__obf_672171c323f2a0f8 := image.Point{X: rand.IntN(__obf_8cff1b162870f5a8) + __obf_1c0b3e980dbbd1e1, Y: rand.IntN(__obf_bc574d6b33b94515)}

		if __obf_76ec865d65fb400e%2 == 0 {
			__obf_f5cb084092927ec6.Y = rand.IntN(__obf_bc574d6b33b94515) + __obf_bc574d6b33b94515*2
			__obf_672171c323f2a0f8.Y = rand.IntN(__obf_bc574d6b33b94515)
		} else {
			__obf_f5cb084092927ec6.Y = rand.IntN(__obf_bc574d6b33b94515) + __obf_bc574d6b33b94515*(__obf_76ec865d65fb400e%2)
			__obf_672171c323f2a0f8.Y = rand.IntN(__obf_bc574d6b33b94515) + __obf_bc574d6b33b94515*2
		}

		__obf_646dbac6c9d3b39a := util.RandColor(__obf_fd80fb07f7829c9c)
		//co.A = uint8(0xee)
		__obf_c0329d72b1b272fd.DrawBeeline(__obf_f5cb084092927ec6, __obf_672171c323f2a0f8, __obf_646dbac6c9d3b39a)
	}
}

func (__obf_0e570eb08dbdd435 *Captcha) DrawDotImage(__obf_894388cdf1835fdb DotType, __obf_24103e8f3612481d *types.Dot, __obf_b409aea2e7dfafc6 *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_ad611a46c4a5252c, _ := util.ParseHexColor(__obf_24103e8f3612481d.PrimaryColor)
	__obf_ad611a46c4a5252c.A = util.FormatAlpha(__obf_b409aea2e7dfafc6.Alpha)

	var __obf_4c029424cf0a53ef image.Image
	var __obf_3afd7555fbd2aaf7 error
	if __obf_894388cdf1835fdb == Shape {
		__obf_4c029424cf0a53ef, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawShapeImage(__obf_24103e8f3612481d, __obf_ad611a46c4a5252c)
	} else {
		__obf_4c029424cf0a53ef, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawStringImage(__obf_24103e8f3612481d, __obf_ad611a46c4a5252c)
	}
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, nil, __obf_3afd7555fbd2aaf7
	}

	__obf_3a5fbe8b5bdcc927, _ := util.ParseHexColor(__obf_2d753efdaf369fc2)
	__obf_06b112bbb1f58b48 := helper.CreateNRGBACanvas(__obf_24103e8f3612481d.Size+10, __obf_24103e8f3612481d.Size+10, true)
	if __obf_b409aea2e7dfafc6.ShowShadow {
		var __obf_879fb97c14eea021 *types.NRGBA
		if __obf_894388cdf1835fdb == Shape {
			__obf_879fb97c14eea021, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawShapeImage(__obf_24103e8f3612481d, __obf_3a5fbe8b5bdcc927)
		} else {
			__obf_879fb97c14eea021, __obf_3afd7555fbd2aaf7 = __obf_0e570eb08dbdd435.DrawStringImage(__obf_24103e8f3612481d, __obf_3a5fbe8b5bdcc927)
		}
		if __obf_3afd7555fbd2aaf7 != nil {
			return nil, nil, __obf_3afd7555fbd2aaf7
		}

		__obf_ba4dbd415a06d4ee := __obf_b409aea2e7dfafc6.ShadowPoint.X
		__obf_cefc24b67196ab53 := __obf_b409aea2e7dfafc6.ShadowPoint.Y
		draw.Draw(__obf_06b112bbb1f58b48, __obf_879fb97c14eea021.Bounds(), __obf_879fb97c14eea021, image.Point{X: __obf_ba4dbd415a06d4ee, Y: __obf_cefc24b67196ab53}, draw.Over)
	}
	draw.Draw(__obf_06b112bbb1f58b48, __obf_4c029424cf0a53ef.Bounds(), __obf_4c029424cf0a53ef, image.Point{}, draw.Over)
	__obf_06b112bbb1f58b48.Rotate(__obf_24103e8f3612481d.Angle, false)

	__obf_fd3c32f905c1a25e := __obf_06b112bbb1f58b48.CalcMarginBlankArea()

	return __obf_06b112bbb1f58b48, __obf_fd3c32f905c1a25e, nil
}

// DrawStringImage draws a text image
// params:
//   - dot: Draw dot
//   - textColor: Text color
//
// returns:
//   - types.NRGBA: Drawn text image
//   - error: Error information
func (__obf_0e570eb08dbdd435 *Captcha) DrawStringImage(__obf_24103e8f3612481d *types.Dot, __obf_58a809110a9cff2c color.Color) (*types.NRGBA, error) {
	__obf_06b112bbb1f58b48 := helper.CreateNRGBACanvas(__obf_24103e8f3612481d.Size+10, __obf_24103e8f3612481d.Size+10, true)

	__obf_04555369bb47c7d0 := fixed.P(12, __obf_24103e8f3612481d.Size-5)
	if util.IsChineseChar(__obf_24103e8f3612481d.Text) {
		__obf_04555369bb47c7d0 = fixed.P(10, __obf_24103e8f3612481d.Size)
	}
	__obf_b409aea2e7dfafc6 := &types.DrawStringParam{
		Color:       __obf_58a809110a9cff2c,
		FontSize:    __obf_24103e8f3612481d.FontSize,
		Size:        __obf_24103e8f3612481d.Size,
		FontDPI:     __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.FontDPI,
		FontHinting: __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.FontHinting,
		Text:        __obf_24103e8f3612481d.Text,
	}
	var __obf_3afd7555fbd2aaf7 error
	__obf_b409aea2e7dfafc6.Font, __obf_3afd7555fbd2aaf7 = assets.PickFont()
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}
	__obf_3afd7555fbd2aaf7 = helper.DrawString(__obf_06b112bbb1f58b48, __obf_b409aea2e7dfafc6, __obf_04555369bb47c7d0)
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	return __obf_06b112bbb1f58b48, nil
}

func (__obf_0e570eb08dbdd435 *Captcha) DrawShapeImage(__obf_24103e8f3612481d *types.Dot, __obf_ad611a46c4a5252c color.Color) (*types.NRGBA, error) {
	__obf_1b1872d59b0cb47d, __obf_57f4dab14dd725ca, __obf_e53e217b1e8b3e8c, __obf_28d80fec96925685 := __obf_ad611a46c4a5252c.RGBA()

	var __obf_79973909653069b4 = []color.RGBA{
		{R: uint8(__obf_1b1872d59b0cb47d), G: uint8(__obf_57f4dab14dd725ca), B: uint8(__obf_e53e217b1e8b3e8c), A: uint8(__obf_28d80fec96925685)},
	}

	__obf_7dad0498e0b3b97d := helper.CreateNRGBACanvas(__obf_24103e8f3612481d.Size+10, __obf_24103e8f3612481d.Size+10, true)
	var __obf_b5e82a98533f8553 image.Rectangle
	var __obf_65fc1d4866918f94 image.Image
	if __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.UseRGBA {
		__obf_06b112bbb1f58b48 := helper.CreateNRGBACanvas(__obf_24103e8f3612481d.Size+10, __obf_24103e8f3612481d.Size+10, true)
		draw.BiLinear.Scale(__obf_06b112bbb1f58b48, __obf_06b112bbb1f58b48.Bounds(), __obf_24103e8f3612481d.Shape, __obf_24103e8f3612481d.Shape.Bounds(), draw.Over, nil)
		__obf_b5e82a98533f8553 = __obf_06b112bbb1f58b48.Bounds()
		__obf_65fc1d4866918f94 = __obf_06b112bbb1f58b48
	} else {
		__obf_06b112bbb1f58b48 := helper.CreatePaletteCanvas(__obf_24103e8f3612481d.Size+10, __obf_24103e8f3612481d.Size+10, __obf_79973909653069b4)
		draw.BiLinear.Scale(__obf_06b112bbb1f58b48, __obf_06b112bbb1f58b48.Bounds(), __obf_24103e8f3612481d.Shape, __obf_24103e8f3612481d.Shape.Bounds(), draw.Over, nil)
		__obf_b5e82a98533f8553 = __obf_06b112bbb1f58b48.Bounds()
		__obf_65fc1d4866918f94 = __obf_06b112bbb1f58b48
	}

	draw.Draw(__obf_7dad0498e0b3b97d, __obf_b5e82a98533f8553, __obf_65fc1d4866918f94, image.Point{}, draw.Over)

	return __obf_7dad0498e0b3b97d, nil
}

func (__obf_a38f2b46ae5b8a61 *Captcha) Generate() (*CaptchaData, error) {
	__obf_894388cdf1835fdb := __obf_524140d8ce6def5b(__obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed)
	if __obf_894388cdf1835fdb == Shape {
		return __obf_a38f2b46ae5b8a61.__obf_eb88aca381048520()
	}

	return __obf_a38f2b46ae5b8a61.__obf_f4833661c773b37f()
}

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_eb88aca381048520() (*CaptchaData, error) {

	var (
		__obf_d597451ffe0651f9, __obf_45ab857fc357fe54, __obf_c4780ba3c86af6d8 []*types.Dot
		__obf_3adb7890adeaf5ea                                                 []image.Image
		__obf_9dbee94bde9bf7c2                                                 []image.Image
		__obf_7e08e5ae21de2666, __obf_ca3de5ebdc66e8cf                         image.Image
		__obf_3afd7555fbd2aaf7                                                 error
	)
	__obf_9dbee94bde9bf7c2, __obf_3afd7555fbd2aaf7 = assets.GetShapes()
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	__obf_e6619f60efb925b4 := util.RandInt(__obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.LenRange.Min, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.LenRange.Max)
	__obf_9dbee94bde9bf7c2 = util.PickN(__obf_9dbee94bde9bf7c2, __obf_e6619f60efb925b4)

	__obf_d597451ffe0651f9 = __obf_a38f2b46ae5b8a61.__obf_f8464ea648ee0a46(__obf_9dbee94bde9bf7c2, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.Size, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.SizeRange, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.DotPadding)
	__obf_4f512648ca5ff60e := rand.Perm(len(__obf_d597451ffe0651f9))
	__obf_1b16744fd175ce29 := util.RandInt(__obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.VerifyLenRange.Min, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.VerifyLenRange.Max)
	__obf_c4780ba3c86af6d8 = make([]*types.Dot, __obf_1b16744fd175ce29)
	for __obf_76ec865d65fb400e, __obf_c4d9e6f5811029c3 := range __obf_4f512648ca5ff60e {
		if __obf_76ec865d65fb400e >= __obf_1b16744fd175ce29 {
			break
		}

		__obf_24103e8f3612481d := __obf_d597451ffe0651f9[__obf_c4d9e6f5811029c3]
		__obf_24103e8f3612481d.Index = __obf_76ec865d65fb400e
		__obf_c4780ba3c86af6d8[__obf_76ec865d65fb400e] = __obf_24103e8f3612481d
		__obf_3adb7890adeaf5ea = append(__obf_3adb7890adeaf5ea, __obf_c4780ba3c86af6d8[__obf_76ec865d65fb400e].Shape)
	}
	__obf_45ab857fc357fe54 = __obf_a38f2b46ae5b8a61.__obf_f8464ea648ee0a46(__obf_3adb7890adeaf5ea, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.Size, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.SizeRange, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.DotPadding)

	__obf_7e08e5ae21de2666, __obf_3afd7555fbd2aaf7 = __obf_a38f2b46ae5b8a61.__obf_ecd9130aa1680f17(Shape, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.Size, __obf_d597451ffe0651f9)
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	__obf_ca3de5ebdc66e8cf, __obf_3afd7555fbd2aaf7 = __obf_a38f2b46ae5b8a61.__obf_8e00d42c5345ca01(Shape, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.Size, __obf_45ab857fc357fe54)
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	return &CaptchaData{
		__obf_8c02fb9797f8ce7a: __obf_c4780ba3c86af6d8,
		__obf_7e08e5ae21de2666: types.NewJPEGImage(__obf_7e08e5ae21de2666),
		__obf_ca3de5ebdc66e8cf: types.NewPNGImage(__obf_ca3de5ebdc66e8cf),
	}, nil
}

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_f4833661c773b37f() (*CaptchaData, error) {

	__obf_e6619f60efb925b4 := util.RandInt(__obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.LenRange.Min, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.LenRange.Max)
	__obf_d76f3c1a10a32b1d := util.PickN(assets.GetChineseChars(), __obf_e6619f60efb925b4)

	var __obf_8c02fb9797f8ce7a, __obf_45ab857fc357fe54, __obf_c4780ba3c86af6d8 []*types.Dot
	var __obf_206eb5fac3f21019 []string
	var __obf_7e08e5ae21de2666, __obf_ca3de5ebdc66e8cf image.Image

	__obf_8c02fb9797f8ce7a = __obf_a38f2b46ae5b8a61.__obf_e916b0796f76cdb3(__obf_d76f3c1a10a32b1d, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.Size, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.SizeRange, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.DotPadding)
	__obf_4f512648ca5ff60e := rand.Perm(len(__obf_8c02fb9797f8ce7a))
	__obf_1b16744fd175ce29 := util.RandInt(__obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.VerifyLenRange.Min, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.VerifyLenRange.Max)
	__obf_c4780ba3c86af6d8 = make([]*types.Dot, __obf_1b16744fd175ce29)
	for __obf_76ec865d65fb400e, __obf_c4d9e6f5811029c3 := range __obf_4f512648ca5ff60e {
		if __obf_76ec865d65fb400e >= __obf_1b16744fd175ce29 {
			break
		}

		__obf_24103e8f3612481d := __obf_8c02fb9797f8ce7a[__obf_c4d9e6f5811029c3]
		__obf_24103e8f3612481d.Index = __obf_76ec865d65fb400e
		__obf_c4780ba3c86af6d8[__obf_76ec865d65fb400e] = __obf_24103e8f3612481d
		__obf_206eb5fac3f21019 = append(__obf_206eb5fac3f21019, __obf_c4780ba3c86af6d8[__obf_76ec865d65fb400e].Text)
	}

	__obf_45ab857fc357fe54 = __obf_a38f2b46ae5b8a61.__obf_e916b0796f76cdb3(__obf_206eb5fac3f21019, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.Size, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.SizeRange, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.DotPadding)
	var __obf_3afd7555fbd2aaf7 error
	__obf_7e08e5ae21de2666, __obf_3afd7555fbd2aaf7 = __obf_a38f2b46ae5b8a61.__obf_ecd9130aa1680f17(Text, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.Size, __obf_8c02fb9797f8ce7a)
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	__obf_ca3de5ebdc66e8cf, __obf_3afd7555fbd2aaf7 = __obf_a38f2b46ae5b8a61.__obf_8e00d42c5345ca01(Text, __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Secondary.Size, __obf_45ab857fc357fe54)
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	return &CaptchaData{
		__obf_8c02fb9797f8ce7a: __obf_c4780ba3c86af6d8,
		__obf_7e08e5ae21de2666: types.NewJPEGImage(__obf_7e08e5ae21de2666),
		__obf_ca3de5ebdc66e8cf: types.NewPNGImage(__obf_ca3de5ebdc66e8cf),
	}, nil
}

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_f8464ea648ee0a46(__obf_9dbee94bde9bf7c2 []image.Image, __obf_07e4883c00829111 types.Size, __obf_22cfa231d727caba types.Range, __obf_6c3787e4e90738e6 int) []*types.Dot {

	__obf_e6619f60efb925b4 := len(__obf_9dbee94bde9bf7c2)
	var __obf_8c02fb9797f8ce7a = make([]*types.Dot, __obf_e6619f60efb925b4)
	__obf_ba3c808c322fb14d := __obf_07e4883c00829111.Width
	__obf_d355b65d00dab8b0 := __obf_07e4883c00829111.Height
	if __obf_6c3787e4e90738e6 > 0 {
		__obf_ba3c808c322fb14d -= __obf_6c3787e4e90738e6
		__obf_d355b65d00dab8b0 -= __obf_6c3787e4e90738e6
	}

	for __obf_76ec865d65fb400e, __obf_a73552353b93ea58 := range __obf_9dbee94bde9bf7c2 {
		__obf_2eb119eabf4bb2b0 := __obf_a38f2b46ae5b8a61.__obf_2eb119eabf4bb2b0()

		__obf_ab1c810ffb18ac92 := util.PickN(__obf_1994734142d06f37, 1)[0]
		__obf_54ed21beabc94e46 := util.PickN(__obf_4e6d924f540378ee, 1)[0]

		__obf_4416ac5e37b142cb := util.RandInt(__obf_22cfa231d727caba.Min, __obf_22cfa231d727caba.Max)

		__obf_4efaefa63fa6e566 := 10
		__obf_05077f2592ac7fef := __obf_ba3c808c322fb14d / __obf_e6619f60efb925b4
		__obf_b289b16d8d39babc := math.Abs(float64(__obf_05077f2592ac7fef) - float64(__obf_4416ac5e37b142cb))
		__obf_938434875a9b0e05 := (__obf_76ec865d65fb400e * __obf_05077f2592ac7fef) + util.RandInt(0, int(max(__obf_b289b16d8d39babc, 1)))
		__obf_1398e12288680db0 := util.RandInt(__obf_4efaefa63fa6e566, __obf_d355b65d00dab8b0+__obf_4416ac5e37b142cb)

		__obf_3eb1093f9d168527 := int(min(max(float64(__obf_938434875a9b0e05), float64(__obf_4efaefa63fa6e566)), float64(__obf_ba3c808c322fb14d-__obf_4efaefa63fa6e566-(__obf_6c3787e4e90738e6*2))))
		__obf_bc574d6b33b94515 := int(min(max(float64(__obf_1398e12288680db0), float64(__obf_4416ac5e37b142cb+__obf_4efaefa63fa6e566)), float64(__obf_d355b65d00dab8b0+(__obf_4416ac5e37b142cb/2)-(__obf_6c3787e4e90738e6*2))))

		__obf_8c02fb9797f8ce7a[__obf_76ec865d65fb400e] = &types.Dot{
			Index:          __obf_76ec865d65fb400e,
			X:              __obf_3eb1093f9d168527,
			Y:              __obf_bc574d6b33b94515 - __obf_4416ac5e37b142cb,
			FontSize:       __obf_4416ac5e37b142cb,
			Size:           __obf_4416ac5e37b142cb,
			Angle:          __obf_2eb119eabf4bb2b0,
			PrimaryColor:   __obf_ab1c810ffb18ac92,
			SecondaryColor: __obf_54ed21beabc94e46,
			Shape:          __obf_a73552353b93ea58,
		}
	}

	return __obf_8c02fb9797f8ce7a
}

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_e916b0796f76cdb3(__obf_d76f3c1a10a32b1d []string, __obf_07e4883c00829111 types.Size, __obf_5cdfffd27c3cb227 types.Range, __obf_6c3787e4e90738e6 int) []*types.Dot {
	__obf_e6619f60efb925b4 := len(__obf_d76f3c1a10a32b1d)
	var __obf_8c02fb9797f8ce7a = make([]*types.Dot, __obf_e6619f60efb925b4)
	__obf_ba3c808c322fb14d := __obf_07e4883c00829111.Width
	__obf_d355b65d00dab8b0 := __obf_07e4883c00829111.Height
	if __obf_6c3787e4e90738e6 > 0 {
		__obf_ba3c808c322fb14d -= __obf_6c3787e4e90738e6
		__obf_d355b65d00dab8b0 -= __obf_6c3787e4e90738e6
	}

	for __obf_76ec865d65fb400e, __obf_a73552353b93ea58 := range __obf_d76f3c1a10a32b1d {
		__obf_2eb119eabf4bb2b0 := __obf_a38f2b46ae5b8a61.__obf_2eb119eabf4bb2b0()

		__obf_ab1c810ffb18ac92 := util.PickN(__obf_1994734142d06f37, 1)[0]
		__obf_54ed21beabc94e46 := util.PickN(__obf_4e6d924f540378ee, 1)[0]

		__obf_4416ac5e37b142cb := util.RandInt(__obf_5cdfffd27c3cb227.Min, __obf_5cdfffd27c3cb227.Max)
		__obf_f5f5a1731ea1dcd3 := __obf_4416ac5e37b142cb
		__obf_e9c83f18f53dc6b7 := __obf_4416ac5e37b142cb

		if util.LenChineseChar(__obf_a73552353b93ea58) > 1 {
			__obf_e9c83f18f53dc6b7 = __obf_4416ac5e37b142cb * util.LenChineseChar(__obf_a73552353b93ea58)

			if __obf_2eb119eabf4bb2b0 > 0 {
				__obf_bef4c0cbf5b1d3b1 := __obf_e9c83f18f53dc6b7 - __obf_4416ac5e37b142cb
				__obf_45632d5fafdc7083 := __obf_2eb119eabf4bb2b0 % 90
				__obf_f86e8f6bb25a4e59 := float64(__obf_bef4c0cbf5b1d3b1) / 90
				__obf_2166b3491b2ba98f := max(float64(__obf_45632d5fafdc7083)*__obf_f86e8f6bb25a4e59, 1)
				__obf_f5f5a1731ea1dcd3 = __obf_f5f5a1731ea1dcd3 + int(__obf_2166b3491b2ba98f)
				__obf_e9c83f18f53dc6b7 = __obf_e9c83f18f53dc6b7 + int(__obf_2166b3491b2ba98f)
			}
		}

		__obf_4efaefa63fa6e566 := 10
		__obf_05077f2592ac7fef := __obf_ba3c808c322fb14d / __obf_e6619f60efb925b4
		__obf_b289b16d8d39babc := math.Abs(float64(__obf_05077f2592ac7fef) - float64(__obf_e9c83f18f53dc6b7))
		__obf_938434875a9b0e05 := (__obf_76ec865d65fb400e * __obf_05077f2592ac7fef) + util.RandInt(0, int(max(__obf_b289b16d8d39babc, 1)))
		__obf_1398e12288680db0 := util.RandInt(__obf_4efaefa63fa6e566, __obf_d355b65d00dab8b0+__obf_f5f5a1731ea1dcd3)

		__obf_3eb1093f9d168527 := int(min(max(float64(__obf_938434875a9b0e05), float64(__obf_4efaefa63fa6e566)), float64(__obf_ba3c808c322fb14d-__obf_4efaefa63fa6e566-(__obf_6c3787e4e90738e6*2))))
		__obf_bc574d6b33b94515 := int(min(max(float64(__obf_1398e12288680db0), float64(__obf_f5f5a1731ea1dcd3+__obf_4efaefa63fa6e566)), float64(__obf_d355b65d00dab8b0+(__obf_f5f5a1731ea1dcd3/2)-(__obf_6c3787e4e90738e6*2))))

		__obf_8c02fb9797f8ce7a[__obf_76ec865d65fb400e] = &types.Dot{
			Index:          __obf_76ec865d65fb400e,
			X:              __obf_3eb1093f9d168527,
			Y:              __obf_bc574d6b33b94515 - __obf_f5f5a1731ea1dcd3,
			FontSize:       __obf_4416ac5e37b142cb,
			Size:           max(__obf_e9c83f18f53dc6b7, __obf_f5f5a1731ea1dcd3),
			Angle:          __obf_2eb119eabf4bb2b0,
			PrimaryColor:   __obf_ab1c810ffb18ac92,
			SecondaryColor: __obf_54ed21beabc94e46,
			Text:           __obf_a73552353b93ea58,
		}
	}

	return __obf_8c02fb9797f8ce7a
}

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_ecd9130aa1680f17(__obf_894388cdf1835fdb DotType, __obf_5cdfffd27c3cb227 types.Size, __obf_8c02fb9797f8ce7a []*types.Dot) (image.Image, error) {

	__obf_b8d6c1c930938f37 := &DrawImageParams{
		Width:       __obf_5cdfffd27c3cb227.Width,
		Height:      __obf_5cdfffd27c3cb227.Height,
		Alpha:       __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.Alpha,
		FontHinting: __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.FontHinting,
		Dots:        __obf_8c02fb9797f8ce7a,
		ShowShadow:  __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.ShowShadow,
		ShadowPoint: __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.ShadowPoint,
	}
	var __obf_3afd7555fbd2aaf7 error
	__obf_b8d6c1c930938f37.Background, __obf_3afd7555fbd2aaf7 = assets.PickBgImage()
	if __obf_3afd7555fbd2aaf7 != nil {
		return nil, __obf_3afd7555fbd2aaf7
	}

	return __obf_a38f2b46ae5b8a61.DrawWithNRGBA(__obf_894388cdf1835fdb, __obf_b8d6c1c930938f37)
}

func (__obf_0e570eb08dbdd435 *Captcha) __obf_8e00d42c5345ca01(__obf_894388cdf1835fdb DotType, __obf_5cdfffd27c3cb227 types.Size, __obf_8c02fb9797f8ce7a []*types.Dot) (image.Image, error) {
	var (
		// err      error
		__obf_334a335c066fc24d = make([]*types.Dot, 0, len(__obf_8c02fb9797f8ce7a))
	)

	__obf_ba3c808c322fb14d := __obf_5cdfffd27c3cb227.Width / len(__obf_8c02fb9797f8ce7a)

	for __obf_76ec865d65fb400e, __obf_a73552353b93ea58 := range __obf_8c02fb9797f8ce7a {
		__obf_e6619f60efb925b4 := 1
		if __obf_894388cdf1835fdb == Text {
			__obf_e6619f60efb925b4 = len(__obf_a73552353b93ea58.Text)
		}

		__obf_a73552353b93ea58.X = int(max(float64(__obf_ba3c808c322fb14d*__obf_76ec865d65fb400e+__obf_ba3c808c322fb14d/__obf_a73552353b93ea58.Size), 8))
		__obf_4ab74b1e595eb38c := max(1, __obf_5cdfffd27c3cb227.Height/16*__obf_e6619f60efb925b4)
		__obf_a73552353b93ea58.Y = __obf_5cdfffd27c3cb227.Height/2 + __obf_a73552353b93ea58.FontSize/2 - rand.IntN(__obf_4ab74b1e595eb38c)

		__obf_334a335c066fc24d = append(__obf_334a335c066fc24d, __obf_a73552353b93ea58)
	}

	__obf_b409aea2e7dfafc6 := &DrawImageParams{
		Width:             __obf_5cdfffd27c3cb227.Width,
		Height:            __obf_5cdfffd27c3cb227.Height,
		Dots:              __obf_334a335c066fc24d,
		BackgroundDistort: __obf_0e570eb08dbdd435.__obf_d8a0e5d116da0b7e(__obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.Secondary.BgDistort),
		BgCircles:         __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.Secondary.BgCircles,
		BgSlimLines:       __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_c4f006c1cc80cc9e []color.Color
	for _, __obf_e3f81f39ec4f78f0 := range __obf_4e6d924f540378ee {
		__obf_646dbac6c9d3b39a, _ := util.ParseHexColor(__obf_e3f81f39ec4f78f0)
		__obf_c4f006c1cc80cc9e = append(__obf_c4f006c1cc80cc9e, __obf_646dbac6c9d3b39a)
	}

	var __obf_0a7e2714f6db07cc []color.Color
	for _, __obf_646dbac6c9d3b39a := range __obf_4e6d924f540378ee {
		__obf_9352bdf4be21c20e, _ := util.ParseHexColor(__obf_646dbac6c9d3b39a)
		__obf_0a7e2714f6db07cc = append(__obf_0a7e2714f6db07cc, __obf_9352bdf4be21c20e)
	}

	if __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.UseRGBA || __obf_0e570eb08dbdd435.__obf_126d1b2696ab5fed.Secondary.NonDeformAble {
		return __obf_0e570eb08dbdd435.DrawWithNRGBA2(__obf_894388cdf1835fdb, __obf_b409aea2e7dfafc6, __obf_c4f006c1cc80cc9e, __obf_0a7e2714f6db07cc)
	}
	return __obf_0e570eb08dbdd435.DrawWithPalette(__obf_894388cdf1835fdb, __obf_b409aea2e7dfafc6, __obf_c4f006c1cc80cc9e, __obf_0a7e2714f6db07cc)
}

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_d8a0e5d116da0b7e(__obf_7021dcbe75a0c4ba int) int {
	switch __obf_7021dcbe75a0c4ba {
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

func (__obf_a38f2b46ae5b8a61 *Captcha) __obf_2eb119eabf4bb2b0() int {
	__obf_1ebc3bf66674d15e := __obf_a38f2b46ae5b8a61.__obf_126d1b2696ab5fed.Primary.AnglePosRange

	__obf_08bb31725c829f7a := rand.IntN(len(__obf_1ebc3bf66674d15e))
	if __obf_08bb31725c829f7a < 0 {
		return 0
	}

	__obf_ae7ba421a6b63a0d := __obf_1ebc3bf66674d15e[__obf_08bb31725c829f7a]
	__obf_244c18b144065473 := util.RandInt(__obf_ae7ba421a6b63a0d.Min, __obf_ae7ba421a6b63a0d.Max)

	return __obf_244c18b144065473
}
