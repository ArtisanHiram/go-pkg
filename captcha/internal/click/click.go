package __obf_256ada120a510dc5

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
	__obf_00f0aeca6467aaa1 *types.ClickOption
}

var _ types.Captcha = (*Captcha)(nil)

func NewCaptcha(__obf_00f0aeca6467aaa1 *types.ClickOption) *Captcha {
	if __obf_00f0aeca6467aaa1 == nil {
		return nil
	}
	return &Captcha{__obf_00f0aeca6467aaa1}
}

func (__obf_9e882512dfecdb44 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_9d6d1d4469cd05e9 := __obf_a1b5215d3a04c765(__obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1)
	if __obf_9d6d1d4469cd05e9 == Shape {
		return __obf_9e882512dfecdb44.__obf_171a337b123229b0()
	}

	return __obf_9e882512dfecdb44.__obf_2c4a976b918c046c()
}

func __obf_a1b5215d3a04c765(__obf_00f0aeca6467aaa1 *types.ClickOption) DotType {
	__obf_ccb5a885e88111ec := DotType(rand.IntN(2))
	if __obf_ccb5a885e88111ec == Shape {
		__obf_00f0aeca6467aaa1.
			UseRGBA = false
		__obf_00f0aeca6467aaa1.
			Secondary.BgDistort = types.DistortLevel1
		__obf_00f0aeca6467aaa1.
			Primary.SizeRange = types.Range{Min: 24, Max: 30}
		__obf_00f0aeca6467aaa1.
			Secondary.SizeRange = types.Range{Min: 14, Max: 20}
	} else {
		__obf_00f0aeca6467aaa1.
			UseRGBA = true
		__obf_00f0aeca6467aaa1.
			Secondary.BgDistort = types.DistortLevel4
		__obf_00f0aeca6467aaa1.
			Primary.SizeRange = types.Range{Min: 32, Max: 42}
		__obf_00f0aeca6467aaa1.
			Secondary.SizeRange = types.Range{Min: 32, Max: 42}
	}
	return __obf_ccb5a885e88111ec
}

func (__obf_f696d34b44f3fdd8 *Captcha) DrawWithNRGBA(__obf_9d6d1d4469cd05e9 DotType, __obf_51bc412a399dedcc *DrawImageParams) (image.Image, error) {
	__obf_cebd8c694afe3f70 := __obf_51bc412a399dedcc.Dots
	__obf_1bbc629d0c78aaf5 := helper.CreateNRGBACanvas(__obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height, true)

	for _, __obf_89154f718ade89cc := range __obf_cebd8c694afe3f70 {
		__obf_862bac5d12d0c8d0, __obf_64e0ac928b992b3f, __obf_975b35bfae5ea6a3 := __obf_f696d34b44f3fdd8.DrawDotImage(__obf_9d6d1d4469cd05e9, __obf_89154f718ade89cc, __obf_51bc412a399dedcc)
		if __obf_975b35bfae5ea6a3 != nil {
			return nil, __obf_975b35bfae5ea6a3
		}
		__obf_ca05692732cb95d7 := __obf_64e0ac928b992b3f.MinX
		__obf_d445e8e315dd163c := __obf_64e0ac928b992b3f.MaxX
		__obf_2fc2fa9f78a1a12c := __obf_64e0ac928b992b3f.MinY
		__obf_3aef9af65125ed7a := __obf_64e0ac928b992b3f.MaxY
		__obf_3d6d3ef86201f347 := __obf_d445e8e315dd163c - __obf_ca05692732cb95d7
		__obf_8bec6f113df201df := __obf_3aef9af65125ed7a - __obf_2fc2fa9f78a1a12c

		draw.Draw(__obf_1bbc629d0c78aaf5, image.Rect(__obf_89154f718ade89cc.X, __obf_89154f718ade89cc.Y, __obf_89154f718ade89cc.X+__obf_3d6d3ef86201f347, __obf_89154f718ade89cc.Y+__obf_8bec6f113df201df), __obf_862bac5d12d0c8d0, image.Point{X: __obf_ca05692732cb95d7, Y: __obf_2fc2fa9f78a1a12c}, draw.Over)

	}
	__obf_3838bd92fc50e2b1 := __obf_51bc412a399dedcc.Background
	__obf_8f1af84d783d1fdf := __obf_1bbc629d0c78aaf5.Bounds()
	__obf_65085240584a4f8d := helper.CreateNRGBACanvas(__obf_8f1af84d783d1fdf.Dx(), __obf_8f1af84d783d1fdf.Dy(), true)
	__obf_101e2e5bf4e001f8 := util.RangCutImagePos(__obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height, __obf_3838bd92fc50e2b1)
	draw.Draw(__obf_65085240584a4f8d, __obf_8f1af84d783d1fdf, __obf_3838bd92fc50e2b1, __obf_101e2e5bf4e001f8, draw.Src)
	draw.Draw(__obf_65085240584a4f8d, __obf_1bbc629d0c78aaf5.Bounds(), __obf_1bbc629d0c78aaf5, image.Point{}, draw.Over)
	__obf_65085240584a4f8d.
		SubImage(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height))
	return __obf_65085240584a4f8d, nil
}

func (__obf_f696d34b44f3fdd8 *Captcha) DrawWithPalette(__obf_9d6d1d4469cd05e9 DotType, __obf_51bc412a399dedcc *DrawImageParams, __obf_eeab5b534e33deeb []color.Color, __obf_acb9de7c156df7f1 []color.Color) (image.Image, error) {
	__obf_cebd8c694afe3f70 := __obf_51bc412a399dedcc.Dots
	__obf_f60047863ba91ca0 := make([]color.Color, 0, len(__obf_acb9de7c156df7f1))
	for _, __obf_7e827bd0d7783b41 := range __obf_acb9de7c156df7f1 {
		__obf_36d557cd25a0848b, __obf_07d0bdb67c95849e, __obf_8f1af84d783d1fdf, _ := __obf_7e827bd0d7783b41.RGBA()
		__obf_df2bdbfe63a6831d := util.FormatAlpha(__obf_51bc412a399dedcc.ThumbDisturbAlpha)
		__obf_f60047863ba91ca0 = append(__obf_f60047863ba91ca0, color.RGBA{R: uint8(__obf_36d557cd25a0848b), G: uint8(__obf_07d0bdb67c95849e), B: uint8(__obf_8f1af84d783d1fdf), A: __obf_df2bdbfe63a6831d})
	}

	var __obf_4e7b09a6621166ad = make([]color.Color, 0, len(__obf_eeab5b534e33deeb)+len(__obf_f60047863ba91ca0)+1)
	__obf_4e7b09a6621166ad = append(__obf_4e7b09a6621166ad, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_4e7b09a6621166ad = append(__obf_4e7b09a6621166ad, __obf_eeab5b534e33deeb...)
	__obf_4e7b09a6621166ad = append(__obf_4e7b09a6621166ad, __obf_f60047863ba91ca0...)
	__obf_1bbc629d0c78aaf5 := types.NewPalette(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height), __obf_4e7b09a6621166ad)
	if __obf_51bc412a399dedcc.BgCircles > 0 {
		__obf_f696d34b44f3fdd8.__obf_347a39507fde8e9c(__obf_1bbc629d0c78aaf5, __obf_51bc412a399dedcc.BgCircles, 1, __obf_f60047863ba91ca0)
	}
	if __obf_51bc412a399dedcc.BgSlimLines > 0 {
		__obf_f696d34b44f3fdd8.__obf_d5c31eba1645313d(__obf_1bbc629d0c78aaf5, __obf_51bc412a399dedcc.BgSlimLines, __obf_f60047863ba91ca0)
	}

	for _, __obf_89154f718ade89cc := range __obf_cebd8c694afe3f70 {
		__obf_50b8eba79108138e, _ := util.ParseHexColor(__obf_89154f718ade89cc.PrimaryColor)
		var __obf_975b35bfae5ea6a3 error
		if __obf_9d6d1d4469cd05e9 == Shape {
			var __obf_862bac5d12d0c8d0 *types.NRGBA
			__obf_862bac5d12d0c8d0, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawShapeImage(__obf_89154f718ade89cc, __obf_50b8eba79108138e)
			if __obf_975b35bfae5ea6a3 != nil {
				return nil, __obf_975b35bfae5ea6a3
			}
			__obf_862bac5d12d0c8d0.
				Rotate(__obf_89154f718ade89cc.Angle, false)
			__obf_6232a21a2cbdccfa := __obf_1bbc629d0c78aaf5.Bounds()
			__obf_cdc1ea3247bb855b := __obf_862bac5d12d0c8d0.Bounds()
			__obf_f71715d5d987b722 := image.Point{X: __obf_6232a21a2cbdccfa.Dx() - __obf_cdc1ea3247bb855b.Dx(), Y: __obf_6232a21a2cbdccfa.Dy() - __obf_cdc1ea3247bb855b.Dy()}
			draw.Draw(__obf_1bbc629d0c78aaf5, image.Rect(__obf_89154f718ade89cc.X, __obf_f71715d5d987b722.Y, __obf_f71715d5d987b722.X+__obf_cdc1ea3247bb855b.Dx(), __obf_f71715d5d987b722.Y+__obf_cdc1ea3247bb855b.Dy()), __obf_862bac5d12d0c8d0, image.Point{}, draw.Over)
		} else {
			__obf_9075da6187c5fa4b := fixed.P(__obf_89154f718ade89cc.X, __obf_89154f718ade89cc.Y)
			__obf_51bc412a399dedcc := &types.DrawStringParam{
				Color:       __obf_50b8eba79108138e,
				FontSize:    __obf_89154f718ade89cc.FontSize,
				Size:        __obf_89154f718ade89cc.Size,
				FontDPI:     __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.FontDPI,
				FontHinting: __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.FontHinting,
				Text:        __obf_89154f718ade89cc.Text,
			}
			__obf_51bc412a399dedcc.
				Font, __obf_975b35bfae5ea6a3 = assets.PickFont()
			if __obf_975b35bfae5ea6a3 != nil {
				return nil, __obf_975b35bfae5ea6a3
			}
			__obf_975b35bfae5ea6a3 = helper.DrawString(__obf_1bbc629d0c78aaf5, __obf_51bc412a399dedcc, __obf_9075da6187c5fa4b)
		}

		if __obf_975b35bfae5ea6a3 != nil {
			return __obf_1bbc629d0c78aaf5, __obf_975b35bfae5ea6a3
		}
	}

	if __obf_51bc412a399dedcc.Background != nil {
		__obf_3838bd92fc50e2b1 := __obf_51bc412a399dedcc.Background
		__obf_8f1af84d783d1fdf := __obf_3838bd92fc50e2b1.Bounds()
		__obf_65085240584a4f8d := helper.CreateNRGBACanvas(__obf_8f1af84d783d1fdf.Dx(), __obf_8f1af84d783d1fdf.Dy(), true)
		__obf_101e2e5bf4e001f8 := util.RangCutImagePos(__obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height, __obf_3838bd92fc50e2b1)
		draw.Draw(__obf_65085240584a4f8d, __obf_8f1af84d783d1fdf, __obf_3838bd92fc50e2b1, __obf_101e2e5bf4e001f8, draw.Src)
		__obf_1bbc629d0c78aaf5.
			Distort(float64(util.RandInt(5, 10)), float64(util.RandInt(120, 200)))
		draw.Draw(__obf_65085240584a4f8d, __obf_1bbc629d0c78aaf5.Bounds(), __obf_1bbc629d0c78aaf5, image.Point{}, draw.Over)
		__obf_377d97e65bff019a := __obf_65085240584a4f8d.SubImage(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height)).(*image.NRGBA)
		return __obf_377d97e65bff019a, nil
	}

	if __obf_51bc412a399dedcc.BackgroundDistort > 0 {
		__obf_1bbc629d0c78aaf5.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_51bc412a399dedcc.BackgroundDistort))
	}

	return __obf_1bbc629d0c78aaf5, nil

}

func (__obf_f696d34b44f3fdd8 *Captcha) DrawWithNRGBA2(__obf_9d6d1d4469cd05e9 DotType, __obf_51bc412a399dedcc *DrawImageParams, __obf_eeab5b534e33deeb []color.Color, __obf_acb9de7c156df7f1 []color.Color) (image.Image, error) {
	__obf_cebd8c694afe3f70 := __obf_51bc412a399dedcc.Dots
	__obf_f60047863ba91ca0 := make([]color.Color, 0, len(__obf_acb9de7c156df7f1))
	for _, __obf_7e827bd0d7783b41 := range __obf_acb9de7c156df7f1 {
		__obf_36d557cd25a0848b, __obf_07d0bdb67c95849e, __obf_8f1af84d783d1fdf, _ := __obf_7e827bd0d7783b41.RGBA()
		__obf_df2bdbfe63a6831d := util.FormatAlpha(__obf_51bc412a399dedcc.ThumbDisturbAlpha)
		__obf_f60047863ba91ca0 = append(__obf_f60047863ba91ca0, color.RGBA{R: uint8(__obf_36d557cd25a0848b), G: uint8(__obf_07d0bdb67c95849e), B: uint8(__obf_8f1af84d783d1fdf), A: __obf_df2bdbfe63a6831d})
	}

	var __obf_4e7b09a6621166ad = make([]color.Color, 0, len(__obf_eeab5b534e33deeb)+len(__obf_f60047863ba91ca0)+1)
	__obf_4e7b09a6621166ad = append(__obf_4e7b09a6621166ad, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_4e7b09a6621166ad = append(__obf_4e7b09a6621166ad, __obf_eeab5b534e33deeb...)
	__obf_4e7b09a6621166ad = append(__obf_4e7b09a6621166ad, __obf_f60047863ba91ca0...)
	__obf_861a8149461b3af2 := types.NewNRGBA(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height), true)
	if __obf_51bc412a399dedcc.Background != nil {
		__obf_3838bd92fc50e2b1 := __obf_51bc412a399dedcc.Background
		__obf_8f1af84d783d1fdf := __obf_3838bd92fc50e2b1.Bounds()
		__obf_65085240584a4f8d := helper.CreateNRGBACanvas(__obf_8f1af84d783d1fdf.Dx(), __obf_8f1af84d783d1fdf.Dy(), true)
		__obf_101e2e5bf4e001f8 := util.RangCutImagePos(__obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height, __obf_3838bd92fc50e2b1)
		draw.Draw(__obf_65085240584a4f8d, __obf_8f1af84d783d1fdf, __obf_3838bd92fc50e2b1, __obf_101e2e5bf4e001f8, draw.Src)
		__obf_377d97e65bff019a := __obf_65085240584a4f8d.SubImage(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height)).(*image.NRGBA)
		draw.Draw(__obf_861a8149461b3af2, __obf_377d97e65bff019a.Bounds(), __obf_377d97e65bff019a, image.Point{}, draw.Over)
	}
	__obf_1bbc629d0c78aaf5 := types.NewPalette(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height), __obf_4e7b09a6621166ad)
	if __obf_51bc412a399dedcc.BgCircles > 0 {
		__obf_f696d34b44f3fdd8.__obf_347a39507fde8e9c(__obf_1bbc629d0c78aaf5, __obf_51bc412a399dedcc.BgCircles, 1, __obf_f60047863ba91ca0)
	}
	if __obf_51bc412a399dedcc.BgSlimLines > 0 {
		__obf_f696d34b44f3fdd8.__obf_d5c31eba1645313d(__obf_1bbc629d0c78aaf5, __obf_51bc412a399dedcc.BgSlimLines, __obf_f60047863ba91ca0)
	}
	if __obf_51bc412a399dedcc.BackgroundDistort > 0 {
		__obf_1bbc629d0c78aaf5.
			Distort(float64(util.RandInt(5, 10)), float64(__obf_51bc412a399dedcc.BackgroundDistort))
	}
	__obf_3941dfce73b39612 := __obf_1bbc629d0c78aaf5.Bounds()
	__obf_3d6d3ef86201f347 := __obf_3941dfce73b39612.Dx() / len(__obf_cebd8c694afe3f70)
	for __obf_a527076966eacc74, __obf_89154f718ade89cc := range __obf_cebd8c694afe3f70 {
		__obf_50b8eba79108138e, _ := util.ParseHexColor(__obf_89154f718ade89cc.PrimaryColor)
		var __obf_975b35bfae5ea6a3 error
		if __obf_9d6d1d4469cd05e9 == Shape {
			var __obf_862bac5d12d0c8d0 *types.NRGBA
			__obf_862bac5d12d0c8d0, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawShapeImage(__obf_89154f718ade89cc, __obf_50b8eba79108138e)
			if __obf_975b35bfae5ea6a3 != nil {
				return nil, __obf_975b35bfae5ea6a3
			}
			__obf_862bac5d12d0c8d0.
				Rotate(__obf_89154f718ade89cc.Angle, false)
			__obf_6232a21a2cbdccfa := __obf_861a8149461b3af2.Bounds()
			__obf_cdc1ea3247bb855b := __obf_862bac5d12d0c8d0.Bounds()
			__obf_f71715d5d987b722 := image.Point{X: __obf_6232a21a2cbdccfa.Dx() - __obf_cdc1ea3247bb855b.Dx(), Y: __obf_6232a21a2cbdccfa.Dy() - __obf_cdc1ea3247bb855b.Dy()}
			draw.Draw(__obf_861a8149461b3af2, image.Rect(__obf_89154f718ade89cc.X, __obf_f71715d5d987b722.Y, __obf_f71715d5d987b722.X+__obf_cdc1ea3247bb855b.Dx(), __obf_f71715d5d987b722.Y+__obf_cdc1ea3247bb855b.Dy()), __obf_862bac5d12d0c8d0, image.Point{}, draw.Over)
		} else {
			var __obf_4bc08140ebd19704 *types.NRGBA
			__obf_4bc08140ebd19704, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawStringImage(__obf_89154f718ade89cc, __obf_50b8eba79108138e)
			if __obf_975b35bfae5ea6a3 != nil {
				return nil, __obf_975b35bfae5ea6a3
			}
			__obf_4bc08140ebd19704.
				Rotate(__obf_89154f718ade89cc.Angle, false)
			__obf_64e0ac928b992b3f := __obf_4bc08140ebd19704.CalcMarginBlankArea()
			__obf_ca05692732cb95d7 := __obf_64e0ac928b992b3f.MinX
			__obf_d445e8e315dd163c := __obf_64e0ac928b992b3f.MaxX
			__obf_2fc2fa9f78a1a12c := __obf_64e0ac928b992b3f.MinY
			__obf_3aef9af65125ed7a := __obf_64e0ac928b992b3f.MaxY
			__obf_4bc08140ebd19704.
				SubImage(image.Rect(__obf_ca05692732cb95d7, __obf_2fc2fa9f78a1a12c, __obf_d445e8e315dd163c, __obf_3aef9af65125ed7a))
			__obf_b110e856c679dd61 := __obf_4bc08140ebd19704.Bounds()
			__obf_6383c52d0f3987a4 := int(max(float64(__obf_3d6d3ef86201f347*__obf_a527076966eacc74+__obf_3d6d3ef86201f347/__obf_b110e856c679dd61.Dx()), 8))
			__obf_7b03186ae310ef3b := util.RandInt(1, __obf_3941dfce73b39612.Dy()-__obf_b110e856c679dd61.Dy()-4)

			draw.Draw(__obf_861a8149461b3af2, image.Rect(__obf_6383c52d0f3987a4, __obf_7b03186ae310ef3b, __obf_6383c52d0f3987a4+__obf_b110e856c679dd61.Dx(), __obf_7b03186ae310ef3b+__obf_b110e856c679dd61.Dy()), __obf_4bc08140ebd19704, image.Point{X: __obf_b110e856c679dd61.Min.X, Y: __obf_b110e856c679dd61.Min.Y}, draw.Over)
		}

		if __obf_975b35bfae5ea6a3 != nil {
			return __obf_861a8149461b3af2, __obf_975b35bfae5ea6a3
		}
	}
	__obf_f802facf53e98d2c := types.NewNRGBA(image.Rect(0, 0, __obf_51bc412a399dedcc.Width, __obf_51bc412a399dedcc.Height), true)
	draw.Draw(__obf_f802facf53e98d2c, __obf_1bbc629d0c78aaf5.Bounds(), __obf_1bbc629d0c78aaf5, image.Point{}, draw.Over)
	draw.Draw(__obf_861a8149461b3af2, __obf_f802facf53e98d2c.Bounds(), __obf_f802facf53e98d2c, image.Point{}, draw.Over)
	return __obf_861a8149461b3af2, nil
}

func (__obf_f696d34b44f3fdd8 *Captcha) __obf_347a39507fde8e9c(__obf_65085240584a4f8d *types.Palette, __obf_9bdbc2f35d3d2360, __obf_a6cecdab13d393b5 int, __obf_f4db8cf3fd61de02 []color.Color) {
	__obf_504a5a4b233acf61 := __obf_65085240584a4f8d.Bounds().Max.X
	__obf_88014894bd793dcb := __obf_65085240584a4f8d.Bounds().Max.Y
	for range __obf_9bdbc2f35d3d2360 {
		__obf_7d5b4edf9d2da96c := util.RandColor(__obf_f4db8cf3fd61de02)
		__obf_36d557cd25a0848b := //co.A = uint8(0xee)
			util.RandInt(1, __obf_a6cecdab13d393b5)
		__obf_65085240584a4f8d.
			DrawCircle(util.RandInt(__obf_36d557cd25a0848b, __obf_504a5a4b233acf61-__obf_36d557cd25a0848b), util.RandInt(__obf_36d557cd25a0848b, __obf_88014894bd793dcb-__obf_36d557cd25a0848b), __obf_36d557cd25a0848b, __obf_7d5b4edf9d2da96c)
	}
}

func (__obf_f696d34b44f3fdd8 *Captcha) __obf_d5c31eba1645313d(__obf_65085240584a4f8d *types.Palette, __obf_5b2ec9e80cd6a121 int, __obf_f4db8cf3fd61de02 []color.Color) {
	__obf_ab40fd1f2df2073f := __obf_65085240584a4f8d.Bounds().Max.X / 10
	__obf_2a0c9feef996d1bf := __obf_ab40fd1f2df2073f * 9
	__obf_dc5b3b7a66d9687c := __obf_65085240584a4f8d.Bounds().Max.Y / 3
	for __obf_a527076966eacc74 := range __obf_5b2ec9e80cd6a121 {
		__obf_f5a246814d160544 := image.Point{X: rand.IntN(__obf_ab40fd1f2df2073f), Y: rand.IntN(__obf_dc5b3b7a66d9687c)}
		__obf_e9260082e1e529cf := image.Point{X: rand.IntN(__obf_ab40fd1f2df2073f) + __obf_2a0c9feef996d1bf, Y: rand.IntN(__obf_dc5b3b7a66d9687c)}

		if __obf_a527076966eacc74%2 == 0 {
			__obf_f5a246814d160544.
				Y = rand.IntN(__obf_dc5b3b7a66d9687c) + __obf_dc5b3b7a66d9687c*2
			__obf_e9260082e1e529cf.
				Y = rand.IntN(__obf_dc5b3b7a66d9687c)
		} else {
			__obf_f5a246814d160544.
				Y = rand.IntN(__obf_dc5b3b7a66d9687c) + __obf_dc5b3b7a66d9687c*(__obf_a527076966eacc74%2)
			__obf_e9260082e1e529cf.
				Y = rand.IntN(__obf_dc5b3b7a66d9687c) + __obf_dc5b3b7a66d9687c*2
		}
		__obf_7d5b4edf9d2da96c := util.RandColor(__obf_f4db8cf3fd61de02)
		__obf_65085240584a4f8d.
			//co.A = uint8(0xee)
			DrawBeeline(__obf_f5a246814d160544, __obf_e9260082e1e529cf, __obf_7d5b4edf9d2da96c)
	}
}

func (__obf_f696d34b44f3fdd8 *Captcha) DrawDotImage(__obf_9d6d1d4469cd05e9 DotType, __obf_ad5e523e38020659 *types.Dot, __obf_51bc412a399dedcc *DrawImageParams) (*types.NRGBA, *types.AreaRect, error) {
	__obf_50b8eba79108138e, _ := util.ParseHexColor(__obf_ad5e523e38020659.PrimaryColor)
	__obf_50b8eba79108138e.
		A = util.FormatAlpha(__obf_51bc412a399dedcc.Alpha)

	var __obf_4bc08140ebd19704 image.Image
	var __obf_975b35bfae5ea6a3 error
	if __obf_9d6d1d4469cd05e9 == Shape {
		__obf_4bc08140ebd19704, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawShapeImage(__obf_ad5e523e38020659, __obf_50b8eba79108138e)
	} else {
		__obf_4bc08140ebd19704, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawStringImage(__obf_ad5e523e38020659, __obf_50b8eba79108138e)
	}
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, nil, __obf_975b35bfae5ea6a3
	}
	__obf_7fcf98583d75cf11, _ := util.ParseHexColor(__obf_a71a33d1fd6e3674)
	__obf_1bbc629d0c78aaf5 := helper.CreateNRGBACanvas(__obf_ad5e523e38020659.Size+10, __obf_ad5e523e38020659.Size+10, true)
	if __obf_51bc412a399dedcc.ShowShadow {
		var __obf_7c202747fa44c947 *types.NRGBA
		if __obf_9d6d1d4469cd05e9 == Shape {
			__obf_7c202747fa44c947, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawShapeImage(__obf_ad5e523e38020659, __obf_7fcf98583d75cf11)
		} else {
			__obf_7c202747fa44c947, __obf_975b35bfae5ea6a3 = __obf_f696d34b44f3fdd8.DrawStringImage(__obf_ad5e523e38020659, __obf_7fcf98583d75cf11)
		}
		if __obf_975b35bfae5ea6a3 != nil {
			return nil, nil, __obf_975b35bfae5ea6a3
		}
		__obf_e1ec515e927aee55 := __obf_51bc412a399dedcc.ShadowPoint.X
		__obf_594d668997d7d3d6 := __obf_51bc412a399dedcc.ShadowPoint.Y
		draw.Draw(__obf_1bbc629d0c78aaf5, __obf_7c202747fa44c947.Bounds(), __obf_7c202747fa44c947, image.Point{X: __obf_e1ec515e927aee55, Y: __obf_594d668997d7d3d6}, draw.Over)
	}
	draw.Draw(__obf_1bbc629d0c78aaf5, __obf_4bc08140ebd19704.Bounds(), __obf_4bc08140ebd19704, image.Point{}, draw.Over)
	__obf_1bbc629d0c78aaf5.
		Rotate(__obf_ad5e523e38020659.Angle, false)
	__obf_4698ac8d747b79f0 := __obf_1bbc629d0c78aaf5.CalcMarginBlankArea()

	return __obf_1bbc629d0c78aaf5,

		// DrawStringImage draws a text image
		// params:
		//   - dot: Draw dot
		//   - textColor: Text color
		//
		// returns:
		//   - types.NRGBA: Drawn text image
		//   - error: Error information
		__obf_4698ac8d747b79f0, nil
}

func (__obf_f696d34b44f3fdd8 *Captcha) DrawStringImage(__obf_ad5e523e38020659 *types.Dot, __obf_f9174f2501d9dda1 color.Color) (*types.NRGBA, error) {
	__obf_1bbc629d0c78aaf5 := helper.CreateNRGBACanvas(__obf_ad5e523e38020659.Size+10, __obf_ad5e523e38020659.Size+10, true)
	__obf_9075da6187c5fa4b := fixed.P(12, __obf_ad5e523e38020659.Size-5)
	if util.IsChineseChar(__obf_ad5e523e38020659.Text) {
		__obf_9075da6187c5fa4b = fixed.P(10, __obf_ad5e523e38020659.Size)
	}
	__obf_51bc412a399dedcc := &types.DrawStringParam{
		Color:       __obf_f9174f2501d9dda1,
		FontSize:    __obf_ad5e523e38020659.FontSize,
		Size:        __obf_ad5e523e38020659.Size,
		FontDPI:     __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.FontDPI,
		FontHinting: __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.FontHinting,
		Text:        __obf_ad5e523e38020659.Text,
	}
	var __obf_975b35bfae5ea6a3 error
	__obf_51bc412a399dedcc.
		Font, __obf_975b35bfae5ea6a3 = assets.PickFont()
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}
	__obf_975b35bfae5ea6a3 = helper.DrawString(__obf_1bbc629d0c78aaf5, __obf_51bc412a399dedcc, __obf_9075da6187c5fa4b)
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}

	return __obf_1bbc629d0c78aaf5, nil
}

func (__obf_f696d34b44f3fdd8 *Captcha) DrawShapeImage(__obf_ad5e523e38020659 *types.Dot, __obf_50b8eba79108138e color.Color) (*types.NRGBA, error) {
	__obf_05868c29181bddef, __obf_ede921056c44b149, __obf_abe2caa2f2407ed3, __obf_05b651b3d5c6b110 := __obf_50b8eba79108138e.RGBA()

	var __obf_3e32eb0efa5a7794 = []color.RGBA{
		{R: uint8(__obf_05868c29181bddef), G: uint8(__obf_ede921056c44b149), B: uint8(__obf_abe2caa2f2407ed3), A: uint8(__obf_05b651b3d5c6b110)},
	}
	__obf_f802facf53e98d2c := helper.CreateNRGBACanvas(__obf_ad5e523e38020659.Size+10, __obf_ad5e523e38020659.Size+10, true)
	var __obf_b110e856c679dd61 image.Rectangle
	var __obf_3838bd92fc50e2b1 image.Image
	if __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.UseRGBA {
		__obf_1bbc629d0c78aaf5 := helper.CreateNRGBACanvas(__obf_ad5e523e38020659.Size+10, __obf_ad5e523e38020659.Size+10, true)
		draw.BiLinear.Scale(__obf_1bbc629d0c78aaf5, __obf_1bbc629d0c78aaf5.Bounds(), __obf_ad5e523e38020659.Shape, __obf_ad5e523e38020659.Shape.Bounds(), draw.Over, nil)
		__obf_b110e856c679dd61 = __obf_1bbc629d0c78aaf5.Bounds()
		__obf_3838bd92fc50e2b1 = __obf_1bbc629d0c78aaf5
	} else {
		__obf_1bbc629d0c78aaf5 := helper.CreatePaletteCanvas(__obf_ad5e523e38020659.Size+10, __obf_ad5e523e38020659.Size+10, __obf_3e32eb0efa5a7794)
		draw.BiLinear.Scale(__obf_1bbc629d0c78aaf5, __obf_1bbc629d0c78aaf5.Bounds(), __obf_ad5e523e38020659.Shape, __obf_ad5e523e38020659.Shape.Bounds(), draw.Over, nil)
		__obf_b110e856c679dd61 = __obf_1bbc629d0c78aaf5.Bounds()
		__obf_3838bd92fc50e2b1 = __obf_1bbc629d0c78aaf5
	}

	draw.Draw(__obf_f802facf53e98d2c, __obf_b110e856c679dd61, __obf_3838bd92fc50e2b1, image.Point{}, draw.Over)

	return __obf_f802facf53e98d2c, nil
}

func (__obf_9e882512dfecdb44 *Captcha) __obf_171a337b123229b0() (*CaptchaData, error) {

	var (
		__obf_ac3663934602b2c7, __obf_c8767c1b455e6aa4, __obf_4c05c302db60e06e []*types.Dot
		__obf_698385d6a52b104c                                                 []image.Image
		__obf_e29b3794bade7654                                                 []image.Image
		__obf_e89b4f6927401041, __obf_b1611b851b2fe0d1                         image.Image
		__obf_975b35bfae5ea6a3                                                 error
	)
	__obf_e29b3794bade7654, __obf_975b35bfae5ea6a3 = assets.GetShapes()
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}
	__obf_6b4c767333694c76 := util.RandInt(__obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.LenRange.Min, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.LenRange.Max)
	__obf_e29b3794bade7654 = util.PickN(__obf_e29b3794bade7654, __obf_6b4c767333694c76)
	__obf_ac3663934602b2c7 = __obf_9e882512dfecdb44.__obf_ad6019e21ced55d1(__obf_e29b3794bade7654, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Height, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.SizeRange, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.DotPadding)
	__obf_797dee5f8342b940 := rand.Perm(len(__obf_ac3663934602b2c7))
	__obf_b29b5a6897702b54 := util.RandInt(__obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.VerifyLenRange.Min, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.VerifyLenRange.Max)
	__obf_4c05c302db60e06e = make([]*types.Dot, __obf_b29b5a6897702b54)
	for __obf_a527076966eacc74, __obf_3ddb3435f741c42b := range __obf_797dee5f8342b940 {
		if __obf_a527076966eacc74 >= __obf_b29b5a6897702b54 {
			break
		}
		__obf_ad5e523e38020659 := __obf_ac3663934602b2c7[__obf_3ddb3435f741c42b]
		__obf_ad5e523e38020659.
			Index = __obf_a527076966eacc74
		__obf_4c05c302db60e06e[__obf_a527076966eacc74] = __obf_ad5e523e38020659
		__obf_698385d6a52b104c = append(__obf_698385d6a52b104c, __obf_4c05c302db60e06e[__obf_a527076966eacc74].Shape)
	}
	__obf_c8767c1b455e6aa4 = __obf_9e882512dfecdb44.__obf_ad6019e21ced55d1(__obf_698385d6a52b104c, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.SizeRange, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.DotPadding)
	__obf_e89b4f6927401041, __obf_975b35bfae5ea6a3 = __obf_9e882512dfecdb44.__obf_71713e28a870267d(Shape, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Height, __obf_ac3663934602b2c7)
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}
	__obf_b1611b851b2fe0d1, __obf_975b35bfae5ea6a3 = __obf_9e882512dfecdb44.__obf_7e87738e474a8615(Shape, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Height, __obf_c8767c1b455e6aa4)
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}

	return &CaptchaData{
		Dots: __obf_4c05c302db60e06e, __obf_e89b4f6927401041: types.NewJPEGImage(__obf_e89b4f6927401041), __obf_b1611b851b2fe0d1: types.NewPNGImage(__obf_b1611b851b2fe0d1),
	}, nil
}

func (__obf_9e882512dfecdb44 *Captcha) __obf_2c4a976b918c046c() (*CaptchaData, error) {
	__obf_6b4c767333694c76 := util.RandInt(__obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.LenRange.Min, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.LenRange.Max)
	__obf_a3a819f93423efbb := util.PickN(assets.GetChineseChars(), __obf_6b4c767333694c76)

	var __obf_cebd8c694afe3f70, __obf_c8767c1b455e6aa4, __obf_4c05c302db60e06e []*types.Dot
	var __obf_1a08c8ca4854256a []string
	var __obf_e89b4f6927401041, __obf_b1611b851b2fe0d1 image.Image
	__obf_cebd8c694afe3f70 = __obf_9e882512dfecdb44.__obf_9c1d59cfff99054b(__obf_a3a819f93423efbb, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Height, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.SizeRange, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.DotPadding)
	__obf_797dee5f8342b940 := rand.Perm(len(__obf_cebd8c694afe3f70))
	__obf_b29b5a6897702b54 := util.RandInt(__obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.VerifyLenRange.Min, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.VerifyLenRange.Max)
	__obf_4c05c302db60e06e = make([]*types.Dot, __obf_b29b5a6897702b54)
	for __obf_a527076966eacc74, __obf_3ddb3435f741c42b := range __obf_797dee5f8342b940 {
		if __obf_a527076966eacc74 >= __obf_b29b5a6897702b54 {
			break
		}
		__obf_ad5e523e38020659 := __obf_cebd8c694afe3f70[__obf_3ddb3435f741c42b]
		__obf_ad5e523e38020659.
			Index = __obf_a527076966eacc74
		__obf_4c05c302db60e06e[__obf_a527076966eacc74] = __obf_ad5e523e38020659
		__obf_1a08c8ca4854256a = append(__obf_1a08c8ca4854256a, __obf_4c05c302db60e06e[__obf_a527076966eacc74].Text)
	}
	__obf_c8767c1b455e6aa4 = __obf_9e882512dfecdb44.__obf_9c1d59cfff99054b(__obf_1a08c8ca4854256a, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.Height, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.SizeRange, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.DotPadding)
	var __obf_975b35bfae5ea6a3 error
	__obf_e89b4f6927401041, __obf_975b35bfae5ea6a3 = __obf_9e882512dfecdb44.__obf_71713e28a870267d(Text, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Height, __obf_cebd8c694afe3f70)
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}
	__obf_b1611b851b2fe0d1, __obf_975b35bfae5ea6a3 = __obf_9e882512dfecdb44.__obf_7e87738e474a8615(Text, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.Width, __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Secondary.Height, __obf_c8767c1b455e6aa4)
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}

	return &CaptchaData{
		Dots: __obf_4c05c302db60e06e, __obf_e89b4f6927401041: types.NewJPEGImage(__obf_e89b4f6927401041), __obf_b1611b851b2fe0d1: types.NewPNGImage(__obf_b1611b851b2fe0d1),
	}, nil
}

func (__obf_9e882512dfecdb44 *Captcha) __obf_ad6019e21ced55d1(__obf_e29b3794bade7654 []image.Image, __obf_3d6d3ef86201f347, __obf_8bec6f113df201df int, __obf_4c03e36dc2fe7cf0 types.Range, __obf_44fb5bf2c98f10ff int) []*types.Dot {
	__obf_6b4c767333694c76 := len(__obf_e29b3794bade7654)
	var __obf_cebd8c694afe3f70 = make([]*types.Dot, __obf_6b4c767333694c76)
	// width := imageSize.Width
	// height := imageSize.Height
	if __obf_44fb5bf2c98f10ff > 0 {
		__obf_3d6d3ef86201f347 -= __obf_44fb5bf2c98f10ff
		__obf_8bec6f113df201df -= __obf_44fb5bf2c98f10ff
	}

	for __obf_a527076966eacc74, __obf_89154f718ade89cc := range __obf_e29b3794bade7654 {
		__obf_69f1b04c526691cb := __obf_9e882512dfecdb44.__obf_69f1b04c526691cb()
		__obf_1959683d77f5e429 := util.PickN(__obf_1cbdb786e99d3d81, 1)[0]
		__obf_ec64aef6b14fdf93 := util.PickN(__obf_931d77fa45b76043, 1)[0]
		__obf_dca141655af4b78a := util.RandInt(__obf_4c03e36dc2fe7cf0.Min, __obf_4c03e36dc2fe7cf0.Max)
		__obf_7b03186ae310ef3b := 10
		__obf_c1c78361496f1532 := __obf_3d6d3ef86201f347 / __obf_6b4c767333694c76
		__obf_c594dfb83d7d5f5d := math.Abs(float64(__obf_c1c78361496f1532) - float64(__obf_dca141655af4b78a))
		__obf_3613cc5b848fd107 := (__obf_a527076966eacc74 * __obf_c1c78361496f1532) + util.RandInt(0, int(max(__obf_c594dfb83d7d5f5d, 1)))
		__obf_4d84899fe91cc79e := util.RandInt(__obf_7b03186ae310ef3b, __obf_8bec6f113df201df+__obf_dca141655af4b78a)
		__obf_3af9ae76a99db8fc := int(min(max(float64(__obf_3613cc5b848fd107), float64(__obf_7b03186ae310ef3b)), float64(__obf_3d6d3ef86201f347-__obf_7b03186ae310ef3b-(__obf_44fb5bf2c98f10ff*2))))
		__obf_dc5b3b7a66d9687c := int(min(max(float64(__obf_4d84899fe91cc79e), float64(__obf_dca141655af4b78a+__obf_7b03186ae310ef3b)), float64(__obf_8bec6f113df201df+(__obf_dca141655af4b78a/2)-(__obf_44fb5bf2c98f10ff*2))))
		__obf_cebd8c694afe3f70[__obf_a527076966eacc74] = &types.Dot{
			Index:          __obf_a527076966eacc74,
			X:              __obf_3af9ae76a99db8fc,
			Y:              __obf_dc5b3b7a66d9687c - __obf_dca141655af4b78a,
			FontSize:       __obf_dca141655af4b78a,
			Size:           __obf_dca141655af4b78a,
			Angle:          __obf_69f1b04c526691cb,
			PrimaryColor:   __obf_1959683d77f5e429,
			SecondaryColor: __obf_ec64aef6b14fdf93,
			Shape:          __obf_89154f718ade89cc,
		}
	}

	return __obf_cebd8c694afe3f70
}

func (__obf_9e882512dfecdb44 *Captcha) __obf_9c1d59cfff99054b(__obf_a3a819f93423efbb []string, __obf_3d6d3ef86201f347, __obf_8bec6f113df201df int, __obf_0c9ec72a02ae49c9 types.Range, __obf_44fb5bf2c98f10ff int) []*types.Dot {
	__obf_6b4c767333694c76 := len(__obf_a3a819f93423efbb)
	var __obf_cebd8c694afe3f70 = make([]*types.Dot, __obf_6b4c767333694c76)
	if __obf_44fb5bf2c98f10ff > 0 {
		__obf_3d6d3ef86201f347 -= __obf_44fb5bf2c98f10ff
		__obf_8bec6f113df201df -= __obf_44fb5bf2c98f10ff
	}

	for __obf_a527076966eacc74, __obf_89154f718ade89cc := range __obf_a3a819f93423efbb {
		__obf_69f1b04c526691cb := __obf_9e882512dfecdb44.__obf_69f1b04c526691cb()
		__obf_1959683d77f5e429 := util.PickN(__obf_1cbdb786e99d3d81, 1)[0]
		__obf_ec64aef6b14fdf93 := util.PickN(__obf_931d77fa45b76043, 1)[0]
		__obf_dca141655af4b78a := util.RandInt(__obf_0c9ec72a02ae49c9.Min, __obf_0c9ec72a02ae49c9.Max)
		__obf_410a31c9e5a9a4e4 := __obf_dca141655af4b78a
		__obf_97f70ce9d1ef16b0 := __obf_dca141655af4b78a

		if util.LenChineseChar(__obf_89154f718ade89cc) > 1 {
			__obf_97f70ce9d1ef16b0 = __obf_dca141655af4b78a * util.LenChineseChar(__obf_89154f718ade89cc)

			if __obf_69f1b04c526691cb > 0 {
				__obf_732f1e914c210596 := __obf_97f70ce9d1ef16b0 - __obf_dca141655af4b78a
				__obf_71c5072741e928fe := __obf_69f1b04c526691cb % 90
				__obf_d14a1262804c942d := float64(__obf_732f1e914c210596) / 90
				__obf_36d557cd25a0848b := max(float64(__obf_71c5072741e928fe)*__obf_d14a1262804c942d, 1)
				__obf_410a31c9e5a9a4e4 = __obf_410a31c9e5a9a4e4 + int(__obf_36d557cd25a0848b)
				__obf_97f70ce9d1ef16b0 = __obf_97f70ce9d1ef16b0 + int(__obf_36d557cd25a0848b)
			}
		}
		__obf_7b03186ae310ef3b := 10
		__obf_c1c78361496f1532 := __obf_3d6d3ef86201f347 / __obf_6b4c767333694c76
		__obf_c594dfb83d7d5f5d := math.Abs(float64(__obf_c1c78361496f1532) - float64(__obf_97f70ce9d1ef16b0))
		__obf_3613cc5b848fd107 := (__obf_a527076966eacc74 * __obf_c1c78361496f1532) + util.RandInt(0, int(max(__obf_c594dfb83d7d5f5d, 1)))
		__obf_4d84899fe91cc79e := util.RandInt(__obf_7b03186ae310ef3b, __obf_8bec6f113df201df+__obf_410a31c9e5a9a4e4)
		__obf_3af9ae76a99db8fc := int(min(max(float64(__obf_3613cc5b848fd107), float64(__obf_7b03186ae310ef3b)), float64(__obf_3d6d3ef86201f347-__obf_7b03186ae310ef3b-(__obf_44fb5bf2c98f10ff*2))))
		__obf_dc5b3b7a66d9687c := int(min(max(float64(__obf_4d84899fe91cc79e), float64(__obf_410a31c9e5a9a4e4+__obf_7b03186ae310ef3b)), float64(__obf_8bec6f113df201df+(__obf_410a31c9e5a9a4e4/2)-(__obf_44fb5bf2c98f10ff*2))))
		__obf_cebd8c694afe3f70[__obf_a527076966eacc74] = &types.Dot{
			Index:          __obf_a527076966eacc74,
			X:              __obf_3af9ae76a99db8fc,
			Y:              __obf_dc5b3b7a66d9687c - __obf_410a31c9e5a9a4e4,
			FontSize:       __obf_dca141655af4b78a,
			Size:           max(__obf_97f70ce9d1ef16b0, __obf_410a31c9e5a9a4e4),
			Angle:          __obf_69f1b04c526691cb,
			PrimaryColor:   __obf_1959683d77f5e429,
			SecondaryColor: __obf_ec64aef6b14fdf93,
			Text:           __obf_89154f718ade89cc,
		}
	}

	return __obf_cebd8c694afe3f70
}

func (__obf_9e882512dfecdb44 *Captcha) __obf_71713e28a870267d(__obf_9d6d1d4469cd05e9 DotType, __obf_3d6d3ef86201f347, __obf_8bec6f113df201df int, __obf_cebd8c694afe3f70 []*types.Dot) (image.Image, error) {
	__obf_6f57a5a696d562a4 := &DrawImageParams{
		Width:       __obf_3d6d3ef86201f347,
		Height:      __obf_8bec6f113df201df,
		Alpha:       __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.Alpha,
		FontHinting: __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.FontHinting,
		Dots:        __obf_cebd8c694afe3f70,
		ShowShadow:  __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.ShowShadow,
		ShadowPoint: __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.ShadowPoint,
	}
	var __obf_975b35bfae5ea6a3 error
	__obf_6f57a5a696d562a4.
		Background, __obf_975b35bfae5ea6a3 = assets.PickBgImage()
	if __obf_975b35bfae5ea6a3 != nil {
		return nil, __obf_975b35bfae5ea6a3
	}

	return __obf_9e882512dfecdb44.DrawWithNRGBA(__obf_9d6d1d4469cd05e9, __obf_6f57a5a696d562a4)
}

func (__obf_f696d34b44f3fdd8 *Captcha) __obf_7e87738e474a8615(__obf_9d6d1d4469cd05e9 DotType, __obf_3d6d3ef86201f347, __obf_8bec6f113df201df int, __obf_cebd8c694afe3f70 []*types.Dot) (image.Image, error) {
	var (
		__obf_f98f2a7993cec733 =
		// err      error
		make([]*types.Dot, 0, len(__obf_cebd8c694afe3f70))
	)
	__obf_51a0eed938905b54 := __obf_3d6d3ef86201f347 / len(__obf_cebd8c694afe3f70)

	for __obf_a527076966eacc74, __obf_89154f718ade89cc := range __obf_cebd8c694afe3f70 {
		__obf_6b4c767333694c76 := 1
		if __obf_9d6d1d4469cd05e9 == Text {
			__obf_6b4c767333694c76 = len(__obf_89154f718ade89cc.Text)
		}
		__obf_89154f718ade89cc.
			X = int(max(float64(__obf_51a0eed938905b54*__obf_a527076966eacc74+__obf_51a0eed938905b54/__obf_89154f718ade89cc.Size), 8))
		__obf_0df37759ad6dead3 := max(1, __obf_8bec6f113df201df/16*__obf_6b4c767333694c76)
		__obf_89154f718ade89cc.
			Y = __obf_8bec6f113df201df/2 + __obf_89154f718ade89cc.FontSize/2 - rand.IntN(__obf_0df37759ad6dead3)
		__obf_f98f2a7993cec733 = append(__obf_f98f2a7993cec733, __obf_89154f718ade89cc)
	}
	__obf_51bc412a399dedcc := &DrawImageParams{
		Width:             __obf_3d6d3ef86201f347,
		Height:            __obf_8bec6f113df201df,
		Dots:              __obf_f98f2a7993cec733,
		BackgroundDistort: __obf_f696d34b44f3fdd8.__obf_b21eeb99b564ef20(__obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.Secondary.BgDistort),
		BgCircles:         __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.Secondary.BgCircles,
		BgSlimLines:       __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.Secondary.DisturbAlpha,
	}

	// params.Background, err = util.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_63245a2f31f3f75b []color.Color
	for _, __obf_9104a94d3697189b := range __obf_931d77fa45b76043 {
		__obf_7d5b4edf9d2da96c, _ := util.ParseHexColor(__obf_9104a94d3697189b)
		__obf_63245a2f31f3f75b = append(__obf_63245a2f31f3f75b, __obf_7d5b4edf9d2da96c)
	}

	var __obf_acb9de7c156df7f1 []color.Color
	for _, __obf_7d5b4edf9d2da96c := range __obf_931d77fa45b76043 {
		__obf_377d97e65bff019a, _ := util.ParseHexColor(__obf_7d5b4edf9d2da96c)
		__obf_acb9de7c156df7f1 = append(__obf_acb9de7c156df7f1, __obf_377d97e65bff019a)
	}

	if __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.UseRGBA || __obf_f696d34b44f3fdd8.__obf_00f0aeca6467aaa1.Secondary.NonDeformAble {
		return __obf_f696d34b44f3fdd8.DrawWithNRGBA2(__obf_9d6d1d4469cd05e9, __obf_51bc412a399dedcc, __obf_63245a2f31f3f75b, __obf_acb9de7c156df7f1)
	}
	return __obf_f696d34b44f3fdd8.DrawWithPalette(__obf_9d6d1d4469cd05e9, __obf_51bc412a399dedcc, __obf_63245a2f31f3f75b, __obf_acb9de7c156df7f1)
}

func (__obf_9e882512dfecdb44 *Captcha) __obf_b21eeb99b564ef20(__obf_501c02e831db0d29 int) int {
	switch __obf_501c02e831db0d29 {
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

func (__obf_9e882512dfecdb44 *Captcha) __obf_69f1b04c526691cb() int {
	__obf_5ade83b8e176093d := __obf_9e882512dfecdb44.__obf_00f0aeca6467aaa1.Primary.AnglePosRange
	__obf_a24b41b71d53701c := rand.IntN(len(__obf_5ade83b8e176093d))
	if __obf_a24b41b71d53701c < 0 {
		return 0
	}
	__obf_16156461caf706d2 := __obf_5ade83b8e176093d[__obf_a24b41b71d53701c]
	__obf_94b130bc9086e749 := util.RandInt(__obf_16156461caf706d2.Min, __obf_16156461caf706d2.Max)

	return __obf_94b130bc9086e749
}
