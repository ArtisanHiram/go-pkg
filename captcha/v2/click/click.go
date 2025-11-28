package __obf_07af88af2ce1a9a4

import (
	"image"
	"image/color"
	"math"
	"math/rand/v2"

	core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Captcha struct {
	__obf_932c987f5e543587 *Options
}

func NewCaptcha(__obf_f6783bb9287a7bcb Option) *Captcha {

	__obf_932c987f5e543587 := &Options{
		FontDPI:     72, // Increased to 200 for high quality rendering
		FontHinting: font.HintingNone,
		Primary: Primary{
			Alpha:      1,
			DotPadding: 10,
			Size:       core.Size{Width: 400, Height: 340},
			LenRange:   core.Range{Min: 3, Max: 6},
			SizeRange:  core.Range{Min: 32, Max: 42},
			AnglePosRange: []core.Range{
				{Min: 20, Max: 35},
				{Min: 35, Max: 45},
				{Min: 45, Max: 60},
				{Min: 290, Max: 305},
				{Min: 305, Max: 325},
				{Min: 325, Max: 330},
			},
		},
		ShowShadow:  true,
		ShadowPoint: core.Point{X: -1, Y: -1},
		Secondary: Secondary{
			BgDistort:      core.DistortLevel4,
			BgCircles:      4,
			BgSlimLines:    2,
			DisturbAlpha:   1,
			DotPadding:     0,
			NonDeformAble:  true,
			VerifyLenRange: core.Range{Min: 2, Max: 4},
			Size:           core.Size{Width: 150, Height: 60},
			SizeRange:      core.Range{Min: 32, Max: 42},
		},
	}

	if __obf_f6783bb9287a7bcb != nil {
		__obf_f6783bb9287a7bcb(__obf_932c987f5e543587)
	}

	return &Captcha{__obf_932c987f5e543587}
}

func __obf_863848e15fd9b7ac(__obf_932c987f5e543587 *Options) DotType {
	__obf_bdf00fc7a64c81aa := DotType(rand.IntN(2))
	if __obf_bdf00fc7a64c81aa == Shape {
		__obf_932c987f5e543587.UseRGBA = false
		__obf_932c987f5e543587.Secondary.BgDistort = core.DistortLevel1
		__obf_932c987f5e543587.Primary.SizeRange = core.Range{Min: 24, Max: 30}
		__obf_932c987f5e543587.Secondary.SizeRange = core.Range{Min: 14, Max: 20}
	} else {
		__obf_932c987f5e543587.UseRGBA = true
		__obf_932c987f5e543587.Secondary.BgDistort = core.DistortLevel4
		__obf_932c987f5e543587.Primary.SizeRange = core.Range{Min: 32, Max: 42}
		__obf_932c987f5e543587.Secondary.SizeRange = core.Range{Min: 32, Max: 42}
	}
	return __obf_bdf00fc7a64c81aa
}

func (__obf_4d5cf3f253a7202e *Captcha) DrawWithNRGBA(__obf_933380cc9c059699 DotType, __obf_acb4c1626c979718 *DrawImageParams) (image.Image, error) {
	__obf_566e0cc4a480c2a3 := __obf_acb4c1626c979718.Dots
	__obf_9ef5e474db89525e := core.CreateNRGBACanvas(__obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height, true)

	for _, __obf_c7a6b4d6583596e4 := range __obf_566e0cc4a480c2a3 {

		__obf_31726915242cb473, __obf_f96137b43d5cfdd5, __obf_d34a61736c704de8 := __obf_4d5cf3f253a7202e.DrawDotImage(__obf_933380cc9c059699, __obf_c7a6b4d6583596e4, __obf_acb4c1626c979718)
		if __obf_d34a61736c704de8 != nil {
			return nil, __obf_d34a61736c704de8
		}
		__obf_a85d617772ec0b8f := __obf_f96137b43d5cfdd5.MinX
		__obf_9809a78ccc476a82 := __obf_f96137b43d5cfdd5.MaxX
		__obf_004ee94ef0dcaa42 := __obf_f96137b43d5cfdd5.MinY
		__obf_99c7f4d469876c55 := __obf_f96137b43d5cfdd5.MaxY

		__obf_99216d7bb578abfc := __obf_9809a78ccc476a82 - __obf_a85d617772ec0b8f
		__obf_f3770deb686e3822 := __obf_99c7f4d469876c55 - __obf_004ee94ef0dcaa42

		draw.Draw(__obf_9ef5e474db89525e, image.Rect(__obf_c7a6b4d6583596e4.X, __obf_c7a6b4d6583596e4.Y, __obf_c7a6b4d6583596e4.X+__obf_99216d7bb578abfc, __obf_c7a6b4d6583596e4.Y+__obf_f3770deb686e3822), __obf_31726915242cb473, image.Point{X: __obf_a85d617772ec0b8f, Y: __obf_004ee94ef0dcaa42}, draw.Over)

	}

	__obf_018d69ee7f20e62e := __obf_acb4c1626c979718.Background
	__obf_63ca5cced9fe80c5 := __obf_9ef5e474db89525e.Bounds()
	__obf_89fb4401e9bb9f7c := core.CreateNRGBACanvas(__obf_63ca5cced9fe80c5.Dx(), __obf_63ca5cced9fe80c5.Dx(), true)
	__obf_538d89a459bc2e72 := core.RangCutImagePos(__obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height, __obf_018d69ee7f20e62e)
	draw.Draw(__obf_89fb4401e9bb9f7c, __obf_63ca5cced9fe80c5, __obf_018d69ee7f20e62e, __obf_538d89a459bc2e72, draw.Src)
	draw.Draw(__obf_89fb4401e9bb9f7c, __obf_9ef5e474db89525e.Bounds(), __obf_9ef5e474db89525e, image.Point{}, draw.Over)
	__obf_89fb4401e9bb9f7c.SubImage(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height))
	return __obf_89fb4401e9bb9f7c, nil
}

func (__obf_4d5cf3f253a7202e *Captcha) DrawWithPalette(__obf_933380cc9c059699 DotType, __obf_acb4c1626c979718 *DrawImageParams, __obf_79c6546f25725f33 []color.Color, __obf_9c8a549620ddfbca []color.Color) (image.Image, error) {
	__obf_566e0cc4a480c2a3 := __obf_acb4c1626c979718.Dots

	__obf_56ee14f4650345b1 := make([]color.Color, 0, len(__obf_9c8a549620ddfbca))
	for _, __obf_0c478c6400ca3a54 := range __obf_9c8a549620ddfbca {
		__obf_0ae0700eb2a118a3, __obf_1b4e0e2e7449549a, __obf_63ca5cced9fe80c5, _ := __obf_0c478c6400ca3a54.RGBA()
		__obf_51b8d5bb767220f0 := core.FormatAlpha(__obf_acb4c1626c979718.ThumbDisturbAlpha)
		__obf_56ee14f4650345b1 = append(__obf_56ee14f4650345b1, color.RGBA{R: uint8(__obf_0ae0700eb2a118a3), G: uint8(__obf_1b4e0e2e7449549a), B: uint8(__obf_63ca5cced9fe80c5), A: __obf_51b8d5bb767220f0})
	}

	var __obf_2ebb67641a0cbce4 = make([]color.Color, 0, len(__obf_79c6546f25725f33)+len(__obf_56ee14f4650345b1)+1)
	__obf_2ebb67641a0cbce4 = append(__obf_2ebb67641a0cbce4, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_2ebb67641a0cbce4 = append(__obf_2ebb67641a0cbce4, __obf_79c6546f25725f33...)
	__obf_2ebb67641a0cbce4 = append(__obf_2ebb67641a0cbce4, __obf_56ee14f4650345b1...)

	__obf_9ef5e474db89525e := core.NewPalette(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height), __obf_2ebb67641a0cbce4)
	if __obf_acb4c1626c979718.BgCircles > 0 {
		__obf_4d5cf3f253a7202e.__obf_3e98723d510a37e5(__obf_9ef5e474db89525e, __obf_acb4c1626c979718.BgCircles, 1, __obf_56ee14f4650345b1)
	}
	if __obf_acb4c1626c979718.BgSlimLines > 0 {
		__obf_4d5cf3f253a7202e.__obf_4f1407dd268a672d(__obf_9ef5e474db89525e, __obf_acb4c1626c979718.BgSlimLines, __obf_56ee14f4650345b1)
	}

	for _, __obf_c7a6b4d6583596e4 := range __obf_566e0cc4a480c2a3 {
		__obf_4917f7bf0f3ce271, _ := core.ParseHexColor(__obf_c7a6b4d6583596e4.PrimaryColor)
		var __obf_d34a61736c704de8 error
		if __obf_933380cc9c059699 == Shape {
			var __obf_31726915242cb473 *core.NRGBA
			__obf_31726915242cb473, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawShapeImage(__obf_c7a6b4d6583596e4, __obf_4917f7bf0f3ce271)
			if __obf_d34a61736c704de8 != nil {
				return nil, __obf_d34a61736c704de8
			}

			__obf_31726915242cb473.Rotate(__obf_c7a6b4d6583596e4.Angle, false)
			__obf_8f8ff165b12ebe41 := __obf_9ef5e474db89525e.Bounds()
			__obf_193b496303a71c22 := __obf_31726915242cb473.Bounds()
			__obf_79dfd067dcc34684 := image.Point{X: __obf_8f8ff165b12ebe41.Dx() - __obf_193b496303a71c22.Dx(), Y: __obf_8f8ff165b12ebe41.Dy() - __obf_193b496303a71c22.Dy()}
			draw.Draw(__obf_9ef5e474db89525e, image.Rect(__obf_c7a6b4d6583596e4.X, __obf_79dfd067dcc34684.Y, __obf_79dfd067dcc34684.X+__obf_193b496303a71c22.Dx(), __obf_79dfd067dcc34684.Y+__obf_193b496303a71c22.Dy()), __obf_31726915242cb473, image.Point{}, draw.Over)
		} else {
			__obf_cd536c4c7711eb76 := fixed.P(__obf_c7a6b4d6583596e4.X, __obf_c7a6b4d6583596e4.Y)
			__obf_acb4c1626c979718 := &core.DrawStringParam{
				Color:       __obf_4917f7bf0f3ce271,
				FontSize:    __obf_c7a6b4d6583596e4.FontSize,
				Size:        __obf_c7a6b4d6583596e4.Size,
				FontDPI:     __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.FontDPI,
				FontHinting: __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.FontHinting,
				Text:        __obf_c7a6b4d6583596e4.Text,
			}
			__obf_acb4c1626c979718.Font, __obf_d34a61736c704de8 = core.PickFont()
			if __obf_d34a61736c704de8 != nil {
				return nil, __obf_d34a61736c704de8
			}
			__obf_d34a61736c704de8 = core.DrawString(__obf_9ef5e474db89525e, __obf_acb4c1626c979718, __obf_cd536c4c7711eb76)
		}

		if __obf_d34a61736c704de8 != nil {
			return __obf_9ef5e474db89525e, __obf_d34a61736c704de8
		}
	}

	if __obf_acb4c1626c979718.Background != nil {
		__obf_018d69ee7f20e62e := __obf_acb4c1626c979718.Background
		__obf_63ca5cced9fe80c5 := __obf_018d69ee7f20e62e.Bounds()
		__obf_89fb4401e9bb9f7c := core.CreateNRGBACanvas(__obf_63ca5cced9fe80c5.Dx(), __obf_63ca5cced9fe80c5.Dy(), true)
		__obf_538d89a459bc2e72 := core.RangCutImagePos(__obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height, __obf_018d69ee7f20e62e)
		draw.Draw(__obf_89fb4401e9bb9f7c, __obf_63ca5cced9fe80c5, __obf_018d69ee7f20e62e, __obf_538d89a459bc2e72, draw.Src)
		__obf_9ef5e474db89525e.Distort(float64(core.RandInt(5, 10)), float64(core.RandInt(120, 200)))
		draw.Draw(__obf_89fb4401e9bb9f7c, __obf_9ef5e474db89525e.Bounds(), __obf_9ef5e474db89525e, image.Point{}, draw.Over)
		__obf_8fbcfb8ce52f034a := __obf_89fb4401e9bb9f7c.SubImage(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height)).(*image.NRGBA)
		return __obf_8fbcfb8ce52f034a, nil
	}

	if __obf_acb4c1626c979718.BackgroundDistort > 0 {
		__obf_9ef5e474db89525e.Distort(float64(core.RandInt(5, 10)), float64(__obf_acb4c1626c979718.BackgroundDistort))
	}

	return __obf_9ef5e474db89525e, nil

}

func (__obf_4d5cf3f253a7202e *Captcha) DrawWithNRGBA2(__obf_933380cc9c059699 DotType, __obf_acb4c1626c979718 *DrawImageParams, __obf_79c6546f25725f33 []color.Color, __obf_9c8a549620ddfbca []color.Color) (image.Image, error) {
	__obf_566e0cc4a480c2a3 := __obf_acb4c1626c979718.Dots

	__obf_56ee14f4650345b1 := make([]color.Color, 0, len(__obf_9c8a549620ddfbca))
	for _, __obf_0c478c6400ca3a54 := range __obf_9c8a549620ddfbca {
		__obf_0ae0700eb2a118a3, __obf_1b4e0e2e7449549a, __obf_63ca5cced9fe80c5, _ := __obf_0c478c6400ca3a54.RGBA()
		__obf_51b8d5bb767220f0 := core.FormatAlpha(__obf_acb4c1626c979718.ThumbDisturbAlpha)
		__obf_56ee14f4650345b1 = append(__obf_56ee14f4650345b1, color.RGBA{R: uint8(__obf_0ae0700eb2a118a3), G: uint8(__obf_1b4e0e2e7449549a), B: uint8(__obf_63ca5cced9fe80c5), A: __obf_51b8d5bb767220f0})
	}

	var __obf_2ebb67641a0cbce4 = make([]color.Color, 0, len(__obf_79c6546f25725f33)+len(__obf_56ee14f4650345b1)+1)
	__obf_2ebb67641a0cbce4 = append(__obf_2ebb67641a0cbce4, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})
	__obf_2ebb67641a0cbce4 = append(__obf_2ebb67641a0cbce4, __obf_79c6546f25725f33...)
	__obf_2ebb67641a0cbce4 = append(__obf_2ebb67641a0cbce4, __obf_56ee14f4650345b1...)

	__obf_25f7a120c449bcd7 := core.NewNRGBA(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height), true)
	if __obf_acb4c1626c979718.Background != nil {
		__obf_018d69ee7f20e62e := __obf_acb4c1626c979718.Background
		__obf_63ca5cced9fe80c5 := __obf_018d69ee7f20e62e.Bounds()
		__obf_89fb4401e9bb9f7c := core.CreateNRGBACanvas(__obf_63ca5cced9fe80c5.Dx(), __obf_63ca5cced9fe80c5.Dy(), true)
		__obf_538d89a459bc2e72 := core.RangCutImagePos(__obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height, __obf_018d69ee7f20e62e)
		draw.Draw(__obf_89fb4401e9bb9f7c, __obf_63ca5cced9fe80c5, __obf_018d69ee7f20e62e, __obf_538d89a459bc2e72, draw.Src)
		__obf_8fbcfb8ce52f034a := __obf_89fb4401e9bb9f7c.SubImage(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height)).(*image.NRGBA)
		draw.Draw(__obf_25f7a120c449bcd7, __obf_8fbcfb8ce52f034a.Bounds(), __obf_8fbcfb8ce52f034a, image.Point{}, draw.Over)
	}

	__obf_9ef5e474db89525e := core.NewPalette(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height), __obf_2ebb67641a0cbce4)
	if __obf_acb4c1626c979718.BgCircles > 0 {
		__obf_4d5cf3f253a7202e.__obf_3e98723d510a37e5(__obf_9ef5e474db89525e, __obf_acb4c1626c979718.BgCircles, 1, __obf_56ee14f4650345b1)
	}
	if __obf_acb4c1626c979718.BgSlimLines > 0 {
		__obf_4d5cf3f253a7202e.__obf_4f1407dd268a672d(__obf_9ef5e474db89525e, __obf_acb4c1626c979718.BgSlimLines, __obf_56ee14f4650345b1)
	}
	if __obf_acb4c1626c979718.BackgroundDistort > 0 {
		__obf_9ef5e474db89525e.Distort(float64(core.RandInt(5, 10)), float64(__obf_acb4c1626c979718.BackgroundDistort))
	}

	__obf_fd2cb6d22fac0e97 := __obf_9ef5e474db89525e.Bounds()
	__obf_99216d7bb578abfc := __obf_fd2cb6d22fac0e97.Dx() / len(__obf_566e0cc4a480c2a3)
	for __obf_03cbab251fc402a6, __obf_c7a6b4d6583596e4 := range __obf_566e0cc4a480c2a3 {
		__obf_4917f7bf0f3ce271, _ := core.ParseHexColor(__obf_c7a6b4d6583596e4.PrimaryColor)
		var __obf_d34a61736c704de8 error
		if __obf_933380cc9c059699 == Shape {
			var __obf_31726915242cb473 *core.NRGBA
			__obf_31726915242cb473, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawShapeImage(__obf_c7a6b4d6583596e4, __obf_4917f7bf0f3ce271)
			if __obf_d34a61736c704de8 != nil {
				return nil, __obf_d34a61736c704de8
			}

			__obf_31726915242cb473.Rotate(__obf_c7a6b4d6583596e4.Angle, false)

			__obf_8f8ff165b12ebe41 := __obf_25f7a120c449bcd7.Bounds()
			__obf_193b496303a71c22 := __obf_31726915242cb473.Bounds()
			__obf_79dfd067dcc34684 := image.Point{X: __obf_8f8ff165b12ebe41.Dx() - __obf_193b496303a71c22.Dx(), Y: __obf_8f8ff165b12ebe41.Dy() - __obf_193b496303a71c22.Dy()}
			draw.Draw(__obf_25f7a120c449bcd7, image.Rect(__obf_c7a6b4d6583596e4.X, __obf_79dfd067dcc34684.Y, __obf_79dfd067dcc34684.X+__obf_193b496303a71c22.Dx(), __obf_79dfd067dcc34684.Y+__obf_193b496303a71c22.Dy()), __obf_31726915242cb473, image.Point{}, draw.Over)
		} else {
			var __obf_f1892f3c3f08afb4 *core.NRGBA
			__obf_f1892f3c3f08afb4, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawStringImage(__obf_c7a6b4d6583596e4, __obf_4917f7bf0f3ce271)
			if __obf_d34a61736c704de8 != nil {
				return nil, __obf_d34a61736c704de8
			}
			__obf_f1892f3c3f08afb4.Rotate(__obf_c7a6b4d6583596e4.Angle, false)

			__obf_f96137b43d5cfdd5 := __obf_f1892f3c3f08afb4.CalcMarginBlankArea()
			__obf_a85d617772ec0b8f := __obf_f96137b43d5cfdd5.MinX
			__obf_9809a78ccc476a82 := __obf_f96137b43d5cfdd5.MaxX
			__obf_004ee94ef0dcaa42 := __obf_f96137b43d5cfdd5.MinY
			__obf_99c7f4d469876c55 := __obf_f96137b43d5cfdd5.MaxY
			__obf_f1892f3c3f08afb4.SubImage(image.Rect(__obf_a85d617772ec0b8f, __obf_004ee94ef0dcaa42, __obf_9809a78ccc476a82, __obf_99c7f4d469876c55))
			__obf_ea4e658fdea3d9f2 := __obf_f1892f3c3f08afb4.Bounds()

			__obf_211580ddcf987dcd := int(max(float64(__obf_99216d7bb578abfc*__obf_03cbab251fc402a6+__obf_99216d7bb578abfc/__obf_ea4e658fdea3d9f2.Dx()), 8))
			__obf_75432551fa3148cb := core.RandInt(1, __obf_fd2cb6d22fac0e97.Dy()-__obf_ea4e658fdea3d9f2.Dy()-4)

			draw.Draw(__obf_25f7a120c449bcd7, image.Rect(__obf_211580ddcf987dcd, __obf_75432551fa3148cb, __obf_211580ddcf987dcd+__obf_ea4e658fdea3d9f2.Dx(), __obf_75432551fa3148cb+__obf_ea4e658fdea3d9f2.Dy()), __obf_f1892f3c3f08afb4, image.Point{X: __obf_ea4e658fdea3d9f2.Min.X, Y: __obf_ea4e658fdea3d9f2.Min.Y}, draw.Over)
		}

		if __obf_d34a61736c704de8 != nil {
			return __obf_25f7a120c449bcd7, __obf_d34a61736c704de8
		}
	}

	__obf_e70e5e919aa716e8 := core.NewNRGBA(image.Rect(0, 0, __obf_acb4c1626c979718.Width, __obf_acb4c1626c979718.Height), true)
	draw.Draw(__obf_e70e5e919aa716e8, __obf_9ef5e474db89525e.Bounds(), __obf_9ef5e474db89525e, image.Point{}, draw.Over)
	draw.Draw(__obf_25f7a120c449bcd7, __obf_e70e5e919aa716e8.Bounds(), __obf_e70e5e919aa716e8, image.Point{}, draw.Over)
	return __obf_25f7a120c449bcd7, nil
}

func (__obf_4d5cf3f253a7202e *Captcha) __obf_3e98723d510a37e5(__obf_89fb4401e9bb9f7c *core.Palette, __obf_fe4610b3cfa41c66, __obf_9c481bb6453259ec int, __obf_005598257b587227 []color.Color) {
	__obf_bfe4f4a5c2977d84 := __obf_89fb4401e9bb9f7c.Bounds().Max.X
	__obf_d72e409170ad8ac3 := __obf_89fb4401e9bb9f7c.Bounds().Max.Y
	for range __obf_fe4610b3cfa41c66 {
		__obf_94d30d5f9cee36ca := core.RandColor(__obf_005598257b587227)
		//co.A = uint8(0xee)
		__obf_0ae0700eb2a118a3 := core.RandInt(1, __obf_9c481bb6453259ec)
		__obf_89fb4401e9bb9f7c.DrawCircle(core.RandInt(__obf_0ae0700eb2a118a3, __obf_bfe4f4a5c2977d84-__obf_0ae0700eb2a118a3), core.RandInt(__obf_0ae0700eb2a118a3, __obf_d72e409170ad8ac3-__obf_0ae0700eb2a118a3), __obf_0ae0700eb2a118a3, __obf_94d30d5f9cee36ca)
	}
}

func (__obf_4d5cf3f253a7202e *Captcha) __obf_4f1407dd268a672d(__obf_89fb4401e9bb9f7c *core.Palette, __obf_ffea43e39a8543fd int, __obf_005598257b587227 []color.Color) {
	__obf_633ddc4d753aa5a3 := __obf_89fb4401e9bb9f7c.Bounds().Max.X / 10
	__obf_21bdbeb015644b3c := __obf_633ddc4d753aa5a3 * 9
	__obf_83181cc5462b3d2a := __obf_89fb4401e9bb9f7c.Bounds().Max.Y / 3
	for __obf_03cbab251fc402a6 := range __obf_ffea43e39a8543fd {
		__obf_6714aa97c4787155 := image.Point{X: rand.IntN(__obf_633ddc4d753aa5a3), Y: rand.IntN(__obf_83181cc5462b3d2a)}
		__obf_47e2bb3198a55acf := image.Point{X: rand.IntN(__obf_633ddc4d753aa5a3) + __obf_21bdbeb015644b3c, Y: rand.IntN(__obf_83181cc5462b3d2a)}

		if __obf_03cbab251fc402a6%2 == 0 {
			__obf_6714aa97c4787155.Y = rand.IntN(__obf_83181cc5462b3d2a) + __obf_83181cc5462b3d2a*2
			__obf_47e2bb3198a55acf.Y = rand.IntN(__obf_83181cc5462b3d2a)
		} else {
			__obf_6714aa97c4787155.Y = rand.IntN(__obf_83181cc5462b3d2a) + __obf_83181cc5462b3d2a*(__obf_03cbab251fc402a6%2)
			__obf_47e2bb3198a55acf.Y = rand.IntN(__obf_83181cc5462b3d2a) + __obf_83181cc5462b3d2a*2
		}

		__obf_94d30d5f9cee36ca := core.RandColor(__obf_005598257b587227)
		//co.A = uint8(0xee)
		__obf_89fb4401e9bb9f7c.DrawBeeline(__obf_6714aa97c4787155, __obf_47e2bb3198a55acf, __obf_94d30d5f9cee36ca)
	}
}

func (__obf_4d5cf3f253a7202e *Captcha) DrawDotImage(__obf_933380cc9c059699 DotType, __obf_95db2efe59be6d8d *Dot, __obf_acb4c1626c979718 *DrawImageParams) (*core.NRGBA, *core.AreaRect, error) {
	__obf_4917f7bf0f3ce271, _ := core.ParseHexColor(__obf_95db2efe59be6d8d.PrimaryColor)
	__obf_4917f7bf0f3ce271.A = core.FormatAlpha(__obf_acb4c1626c979718.Alpha)

	var __obf_f1892f3c3f08afb4 image.Image
	var __obf_d34a61736c704de8 error
	if __obf_933380cc9c059699 == Shape {
		__obf_f1892f3c3f08afb4, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawShapeImage(__obf_95db2efe59be6d8d, __obf_4917f7bf0f3ce271)
	} else {
		__obf_f1892f3c3f08afb4, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawStringImage(__obf_95db2efe59be6d8d, __obf_4917f7bf0f3ce271)
	}
	if __obf_d34a61736c704de8 != nil {
		return nil, nil, __obf_d34a61736c704de8
	}

	__obf_efb30116f914f4e5, _ := core.ParseHexColor(__obf_92ae0bd65abc7c17)
	__obf_9ef5e474db89525e := core.CreateNRGBACanvas(__obf_95db2efe59be6d8d.Size+10, __obf_95db2efe59be6d8d.Size+10, true)
	if __obf_acb4c1626c979718.ShowShadow {
		var __obf_7265ee89e9d57156 *core.NRGBA
		if __obf_933380cc9c059699 == Shape {
			__obf_7265ee89e9d57156, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawShapeImage(__obf_95db2efe59be6d8d, __obf_efb30116f914f4e5)
		} else {
			__obf_7265ee89e9d57156, __obf_d34a61736c704de8 = __obf_4d5cf3f253a7202e.DrawStringImage(__obf_95db2efe59be6d8d, __obf_efb30116f914f4e5)
		}
		if __obf_d34a61736c704de8 != nil {
			return nil, nil, __obf_d34a61736c704de8
		}

		__obf_a0c23644be617d63 := __obf_acb4c1626c979718.ShadowPoint.X
		__obf_bd696aebca2e6d73 := __obf_acb4c1626c979718.ShadowPoint.Y
		draw.Draw(__obf_9ef5e474db89525e, __obf_7265ee89e9d57156.Bounds(), __obf_7265ee89e9d57156, image.Point{X: __obf_a0c23644be617d63, Y: __obf_bd696aebca2e6d73}, draw.Over)
	}
	draw.Draw(__obf_9ef5e474db89525e, __obf_f1892f3c3f08afb4.Bounds(), __obf_f1892f3c3f08afb4, image.Point{}, draw.Over)
	__obf_9ef5e474db89525e.Rotate(__obf_95db2efe59be6d8d.Angle, false)

	__obf_24784eb862fa63e0 := __obf_9ef5e474db89525e.CalcMarginBlankArea()

	return __obf_9ef5e474db89525e, __obf_24784eb862fa63e0, nil
}

// DrawStringImage draws a text image
// params:
//   - dot: Draw dot
//   - textColor: Text color
//
// returns:
//   - core.NRGBA: Drawn text image
//   - error: Error information
func (__obf_4d5cf3f253a7202e *Captcha) DrawStringImage(__obf_95db2efe59be6d8d *Dot, __obf_1cd6dca184602806 color.Color) (*core.NRGBA, error) {
	__obf_9ef5e474db89525e := core.CreateNRGBACanvas(__obf_95db2efe59be6d8d.Size+10, __obf_95db2efe59be6d8d.Size+10, true)

	__obf_cd536c4c7711eb76 := fixed.P(12, __obf_95db2efe59be6d8d.Size-5)
	if core.IsChineseChar(__obf_95db2efe59be6d8d.Text) {
		__obf_cd536c4c7711eb76 = fixed.P(10, __obf_95db2efe59be6d8d.Size)
	}
	__obf_acb4c1626c979718 := &core.DrawStringParam{
		Color:       __obf_1cd6dca184602806,
		FontSize:    __obf_95db2efe59be6d8d.FontSize,
		Size:        __obf_95db2efe59be6d8d.Size,
		FontDPI:     __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.FontDPI,
		FontHinting: __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.FontHinting,
		Text:        __obf_95db2efe59be6d8d.Text,
	}
	var __obf_d34a61736c704de8 error
	__obf_acb4c1626c979718.Font, __obf_d34a61736c704de8 = core.PickFont()
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}
	__obf_d34a61736c704de8 = core.DrawString(__obf_9ef5e474db89525e, __obf_acb4c1626c979718, __obf_cd536c4c7711eb76)
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	return __obf_9ef5e474db89525e, nil
}

func (__obf_4d5cf3f253a7202e *Captcha) DrawShapeImage(__obf_95db2efe59be6d8d *Dot, __obf_4917f7bf0f3ce271 color.Color) (*core.NRGBA, error) {
	__obf_91550f102c90cbd0, __obf_9a898a97b075e306, __obf_88668d39fc040d66, __obf_dc3c71fb9bd3f432 := __obf_4917f7bf0f3ce271.RGBA()

	var __obf_dcdbecdd279519f0 = []color.RGBA{
		{R: uint8(__obf_91550f102c90cbd0), G: uint8(__obf_9a898a97b075e306), B: uint8(__obf_88668d39fc040d66), A: uint8(__obf_dc3c71fb9bd3f432)},
	}

	__obf_e70e5e919aa716e8 := core.CreateNRGBACanvas(__obf_95db2efe59be6d8d.Size+10, __obf_95db2efe59be6d8d.Size+10, true)
	var __obf_ea4e658fdea3d9f2 image.Rectangle
	var __obf_018d69ee7f20e62e image.Image
	if __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.UseRGBA {
		__obf_9ef5e474db89525e := core.CreateNRGBACanvas(__obf_95db2efe59be6d8d.Size+10, __obf_95db2efe59be6d8d.Size+10, true)
		draw.BiLinear.Scale(__obf_9ef5e474db89525e, __obf_9ef5e474db89525e.Bounds(), __obf_95db2efe59be6d8d.Shape, __obf_95db2efe59be6d8d.Shape.Bounds(), draw.Over, nil)
		__obf_ea4e658fdea3d9f2 = __obf_9ef5e474db89525e.Bounds()
		__obf_018d69ee7f20e62e = __obf_9ef5e474db89525e
	} else {
		__obf_9ef5e474db89525e := core.CreatePaletteCanvas(__obf_95db2efe59be6d8d.Size+10, __obf_95db2efe59be6d8d.Size+10, __obf_dcdbecdd279519f0)
		draw.BiLinear.Scale(__obf_9ef5e474db89525e, __obf_9ef5e474db89525e.Bounds(), __obf_95db2efe59be6d8d.Shape, __obf_95db2efe59be6d8d.Shape.Bounds(), draw.Over, nil)
		__obf_ea4e658fdea3d9f2 = __obf_9ef5e474db89525e.Bounds()
		__obf_018d69ee7f20e62e = __obf_9ef5e474db89525e
	}

	draw.Draw(__obf_e70e5e919aa716e8, __obf_ea4e658fdea3d9f2, __obf_018d69ee7f20e62e, image.Point{}, draw.Over)

	return __obf_e70e5e919aa716e8, nil
}

func (__obf_660839016fa371e0 *Captcha) Generate() (*CaptchaData, error) {
	__obf_933380cc9c059699 := __obf_863848e15fd9b7ac(__obf_660839016fa371e0.__obf_932c987f5e543587)
	if __obf_933380cc9c059699 == Shape {
		return __obf_660839016fa371e0.__obf_bdd36d2d988da5b1()
	}

	return __obf_660839016fa371e0.__obf_e72c970f70d3caac()
}

func (__obf_660839016fa371e0 *Captcha) __obf_bdd36d2d988da5b1() (*CaptchaData, error) {

	var (
		__obf_6c0e30d02bce5309, __obf_ba116b5f2d09f4a0, __obf_3016230bea327943 []*Dot
		__obf_d8dcaa08dbe1090f                                                 []image.Image
		__obf_53bd40d58e0ebf59                                                 []image.Image
		__obf_255287b28188217d, __obf_cc78efb9dd19cd20                         image.Image
		__obf_d34a61736c704de8                                                 error
	)
	__obf_53bd40d58e0ebf59, __obf_d34a61736c704de8 = core.GetShapes()
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	__obf_45135e7749d3529c := core.RandInt(__obf_660839016fa371e0.__obf_932c987f5e543587.Primary.LenRange.Min, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.LenRange.Max)
	__obf_53bd40d58e0ebf59 = core.PickN(__obf_53bd40d58e0ebf59, __obf_45135e7749d3529c)

	__obf_6c0e30d02bce5309 = __obf_660839016fa371e0.__obf_5016eeb67d8bcd14(__obf_53bd40d58e0ebf59, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.Size, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.SizeRange, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.DotPadding)
	__obf_37bc4ab880b9542e := rand.Perm(len(__obf_6c0e30d02bce5309))
	__obf_84ca1e2abdcb982d := core.RandInt(__obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.VerifyLenRange.Min, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.VerifyLenRange.Max)
	__obf_3016230bea327943 = make([]*Dot, __obf_84ca1e2abdcb982d)
	for __obf_03cbab251fc402a6, __obf_f4a4ada236209860 := range __obf_37bc4ab880b9542e {
		if __obf_03cbab251fc402a6 >= __obf_84ca1e2abdcb982d {
			break
		}

		__obf_95db2efe59be6d8d := __obf_6c0e30d02bce5309[__obf_f4a4ada236209860]
		__obf_95db2efe59be6d8d.Index = __obf_03cbab251fc402a6
		__obf_3016230bea327943[__obf_03cbab251fc402a6] = __obf_95db2efe59be6d8d
		__obf_d8dcaa08dbe1090f = append(__obf_d8dcaa08dbe1090f, __obf_3016230bea327943[__obf_03cbab251fc402a6].Shape)
	}
	__obf_ba116b5f2d09f4a0 = __obf_660839016fa371e0.__obf_5016eeb67d8bcd14(__obf_d8dcaa08dbe1090f, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.Size, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.SizeRange, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.DotPadding)

	__obf_255287b28188217d, __obf_d34a61736c704de8 = __obf_660839016fa371e0.__obf_7575de1fe35d06f1(Shape, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.Size, __obf_6c0e30d02bce5309)
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	__obf_cc78efb9dd19cd20, __obf_d34a61736c704de8 = __obf_660839016fa371e0.__obf_82b741c7da7e2910(Shape, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.Size, __obf_ba116b5f2d09f4a0)
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	return &CaptchaData{
		__obf_566e0cc4a480c2a3: __obf_3016230bea327943,
		__obf_255287b28188217d: core.NewJPEGImage(__obf_255287b28188217d),
		__obf_cc78efb9dd19cd20: core.NewPNGImage(__obf_cc78efb9dd19cd20),
	}, nil
}

func (__obf_660839016fa371e0 *Captcha) __obf_e72c970f70d3caac() (*CaptchaData, error) {

	__obf_45135e7749d3529c := core.RandInt(__obf_660839016fa371e0.__obf_932c987f5e543587.Primary.LenRange.Min, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.LenRange.Max)
	__obf_0286167621b2e2a4 := core.PickN(core.GetChineseChars(), __obf_45135e7749d3529c)

	var __obf_566e0cc4a480c2a3, __obf_ba116b5f2d09f4a0, __obf_3016230bea327943 []*Dot
	var __obf_345d8ceea9e75237 []string
	var __obf_255287b28188217d, __obf_cc78efb9dd19cd20 image.Image

	__obf_566e0cc4a480c2a3 = __obf_660839016fa371e0.__obf_b05d055448167fc5(__obf_0286167621b2e2a4, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.Size, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.SizeRange, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.DotPadding)
	__obf_37bc4ab880b9542e := rand.Perm(len(__obf_566e0cc4a480c2a3))
	__obf_84ca1e2abdcb982d := core.RandInt(__obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.VerifyLenRange.Min, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.VerifyLenRange.Max)
	__obf_3016230bea327943 = make([]*Dot, __obf_84ca1e2abdcb982d)
	for __obf_03cbab251fc402a6, __obf_f4a4ada236209860 := range __obf_37bc4ab880b9542e {
		if __obf_03cbab251fc402a6 >= __obf_84ca1e2abdcb982d {
			break
		}

		__obf_95db2efe59be6d8d := __obf_566e0cc4a480c2a3[__obf_f4a4ada236209860]
		__obf_95db2efe59be6d8d.Index = __obf_03cbab251fc402a6
		__obf_3016230bea327943[__obf_03cbab251fc402a6] = __obf_95db2efe59be6d8d
		__obf_345d8ceea9e75237 = append(__obf_345d8ceea9e75237, __obf_3016230bea327943[__obf_03cbab251fc402a6].Text)
	}

	__obf_ba116b5f2d09f4a0 = __obf_660839016fa371e0.__obf_b05d055448167fc5(__obf_345d8ceea9e75237, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.Size, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.SizeRange, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.DotPadding)
	var __obf_d34a61736c704de8 error
	__obf_255287b28188217d, __obf_d34a61736c704de8 = __obf_660839016fa371e0.__obf_7575de1fe35d06f1(Text, __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.Size, __obf_566e0cc4a480c2a3)
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	__obf_cc78efb9dd19cd20, __obf_d34a61736c704de8 = __obf_660839016fa371e0.__obf_82b741c7da7e2910(Text, __obf_660839016fa371e0.__obf_932c987f5e543587.Secondary.Size, __obf_ba116b5f2d09f4a0)
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	return &CaptchaData{
		__obf_566e0cc4a480c2a3: __obf_3016230bea327943,
		__obf_255287b28188217d: core.NewJPEGImage(__obf_255287b28188217d),
		__obf_cc78efb9dd19cd20: core.NewPNGImage(__obf_cc78efb9dd19cd20),
	}, nil
}

func (__obf_660839016fa371e0 *Captcha) __obf_5016eeb67d8bcd14(__obf_53bd40d58e0ebf59 []image.Image, __obf_1762bce44d66e7e4 core.Size, __obf_dc443379ca054cb8 core.Range, __obf_1678fc5c695835bf int) []*Dot {

	__obf_45135e7749d3529c := len(__obf_53bd40d58e0ebf59)
	var __obf_566e0cc4a480c2a3 = make([]*Dot, __obf_45135e7749d3529c)
	__obf_99216d7bb578abfc := __obf_1762bce44d66e7e4.Width
	__obf_f3770deb686e3822 := __obf_1762bce44d66e7e4.Height
	if __obf_1678fc5c695835bf > 0 {
		__obf_99216d7bb578abfc -= __obf_1678fc5c695835bf
		__obf_f3770deb686e3822 -= __obf_1678fc5c695835bf
	}

	for __obf_03cbab251fc402a6, __obf_c7a6b4d6583596e4 := range __obf_53bd40d58e0ebf59 {
		__obf_77015c656c3b620a := __obf_660839016fa371e0.__obf_77015c656c3b620a()

		__obf_207f442850e66e11 := core.PickN(__obf_054072fad5ce36d7, 1)[0]
		__obf_85051d5c8e9544e4 := core.PickN(__obf_4c1e9766248eece8, 1)[0]

		__obf_09170a70364d7719 := core.RandInt(__obf_dc443379ca054cb8.Min, __obf_dc443379ca054cb8.Max)

		__obf_75432551fa3148cb := 10
		__obf_2f15af799924b212 := __obf_99216d7bb578abfc / __obf_45135e7749d3529c
		__obf_f176800a81a10031 := math.Abs(float64(__obf_2f15af799924b212) - float64(__obf_09170a70364d7719))
		__obf_685ced7198e1a44b := (__obf_03cbab251fc402a6 * __obf_2f15af799924b212) + core.RandInt(0, int(max(__obf_f176800a81a10031, 1)))
		__obf_2acce876a29fdb64 := core.RandInt(__obf_75432551fa3148cb, __obf_f3770deb686e3822+__obf_09170a70364d7719)

		__obf_1abdfad8de03e3cc := int(min(max(float64(__obf_685ced7198e1a44b), float64(__obf_75432551fa3148cb)), float64(__obf_99216d7bb578abfc-__obf_75432551fa3148cb-(__obf_1678fc5c695835bf*2))))
		__obf_83181cc5462b3d2a := int(min(max(float64(__obf_2acce876a29fdb64), float64(__obf_09170a70364d7719+__obf_75432551fa3148cb)), float64(__obf_f3770deb686e3822+(__obf_09170a70364d7719/2)-(__obf_1678fc5c695835bf*2))))

		__obf_566e0cc4a480c2a3[__obf_03cbab251fc402a6] = &Dot{
			Index:          __obf_03cbab251fc402a6,
			X:              __obf_1abdfad8de03e3cc,
			Y:              __obf_83181cc5462b3d2a - __obf_09170a70364d7719,
			FontSize:       __obf_09170a70364d7719,
			Size:           __obf_09170a70364d7719,
			Angle:          __obf_77015c656c3b620a,
			PrimaryColor:   __obf_207f442850e66e11,
			SecondaryColor: __obf_85051d5c8e9544e4,
			Shape:          __obf_c7a6b4d6583596e4,
		}
	}

	return __obf_566e0cc4a480c2a3
}

func (__obf_660839016fa371e0 *Captcha) __obf_b05d055448167fc5(__obf_0286167621b2e2a4 []string, __obf_1762bce44d66e7e4 core.Size, __obf_e0f3c13244d13853 core.Range, __obf_1678fc5c695835bf int) []*Dot {
	__obf_45135e7749d3529c := len(__obf_0286167621b2e2a4)
	var __obf_566e0cc4a480c2a3 = make([]*Dot, __obf_45135e7749d3529c)
	__obf_99216d7bb578abfc := __obf_1762bce44d66e7e4.Width
	__obf_f3770deb686e3822 := __obf_1762bce44d66e7e4.Height
	if __obf_1678fc5c695835bf > 0 {
		__obf_99216d7bb578abfc -= __obf_1678fc5c695835bf
		__obf_f3770deb686e3822 -= __obf_1678fc5c695835bf
	}

	for __obf_03cbab251fc402a6, __obf_c7a6b4d6583596e4 := range __obf_0286167621b2e2a4 {
		__obf_77015c656c3b620a := __obf_660839016fa371e0.__obf_77015c656c3b620a()

		__obf_207f442850e66e11 := core.PickN(__obf_054072fad5ce36d7, 1)[0]
		__obf_85051d5c8e9544e4 := core.PickN(__obf_4c1e9766248eece8, 1)[0]

		__obf_09170a70364d7719 := core.RandInt(__obf_e0f3c13244d13853.Min, __obf_e0f3c13244d13853.Max)
		__obf_6cd5d2435ccf39bd := __obf_09170a70364d7719
		__obf_eb2d35b73bf5aa45 := __obf_09170a70364d7719

		if core.LenChineseChar(__obf_c7a6b4d6583596e4) > 1 {
			__obf_eb2d35b73bf5aa45 = __obf_09170a70364d7719 * core.LenChineseChar(__obf_c7a6b4d6583596e4)

			if __obf_77015c656c3b620a > 0 {
				__obf_5776e61453810e74 := __obf_eb2d35b73bf5aa45 - __obf_09170a70364d7719
				__obf_963e660ab9f7f69e := __obf_77015c656c3b620a % 90
				__obf_a0c5dff1ba6b88c8 := float64(__obf_5776e61453810e74) / 90
				__obf_0ae0700eb2a118a3 := max(float64(__obf_963e660ab9f7f69e)*__obf_a0c5dff1ba6b88c8, 1)
				__obf_6cd5d2435ccf39bd = __obf_6cd5d2435ccf39bd + int(__obf_0ae0700eb2a118a3)
				__obf_eb2d35b73bf5aa45 = __obf_eb2d35b73bf5aa45 + int(__obf_0ae0700eb2a118a3)
			}
		}

		__obf_75432551fa3148cb := 10
		__obf_2f15af799924b212 := __obf_99216d7bb578abfc / __obf_45135e7749d3529c
		__obf_f176800a81a10031 := math.Abs(float64(__obf_2f15af799924b212) - float64(__obf_eb2d35b73bf5aa45))
		__obf_685ced7198e1a44b := (__obf_03cbab251fc402a6 * __obf_2f15af799924b212) + core.RandInt(0, int(max(__obf_f176800a81a10031, 1)))
		__obf_2acce876a29fdb64 := core.RandInt(__obf_75432551fa3148cb, __obf_f3770deb686e3822+__obf_6cd5d2435ccf39bd)

		__obf_1abdfad8de03e3cc := int(min(max(float64(__obf_685ced7198e1a44b), float64(__obf_75432551fa3148cb)), float64(__obf_99216d7bb578abfc-__obf_75432551fa3148cb-(__obf_1678fc5c695835bf*2))))
		__obf_83181cc5462b3d2a := int(min(max(float64(__obf_2acce876a29fdb64), float64(__obf_6cd5d2435ccf39bd+__obf_75432551fa3148cb)), float64(__obf_f3770deb686e3822+(__obf_6cd5d2435ccf39bd/2)-(__obf_1678fc5c695835bf*2))))

		__obf_566e0cc4a480c2a3[__obf_03cbab251fc402a6] = &Dot{
			Index:          __obf_03cbab251fc402a6,
			X:              __obf_1abdfad8de03e3cc,
			Y:              __obf_83181cc5462b3d2a - __obf_6cd5d2435ccf39bd,
			FontSize:       __obf_09170a70364d7719,
			Size:           max(__obf_eb2d35b73bf5aa45, __obf_6cd5d2435ccf39bd),
			Angle:          __obf_77015c656c3b620a,
			PrimaryColor:   __obf_207f442850e66e11,
			SecondaryColor: __obf_85051d5c8e9544e4,
			Text:           __obf_c7a6b4d6583596e4,
		}
	}

	return __obf_566e0cc4a480c2a3
}

func (__obf_660839016fa371e0 *Captcha) __obf_7575de1fe35d06f1(__obf_933380cc9c059699 DotType, __obf_e0f3c13244d13853 core.Size, __obf_566e0cc4a480c2a3 []*Dot) (image.Image, error) {

	__obf_38a73a0a94d0be23 := &DrawImageParams{
		Width:       __obf_e0f3c13244d13853.Width,
		Height:      __obf_e0f3c13244d13853.Height,
		Alpha:       __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.Alpha,
		FontHinting: __obf_660839016fa371e0.__obf_932c987f5e543587.FontHinting,
		Dots:        __obf_566e0cc4a480c2a3,
		ShowShadow:  __obf_660839016fa371e0.__obf_932c987f5e543587.ShowShadow,
		ShadowPoint: __obf_660839016fa371e0.__obf_932c987f5e543587.ShadowPoint,
	}
	var __obf_d34a61736c704de8 error
	__obf_38a73a0a94d0be23.Background, __obf_d34a61736c704de8 = core.PickBgImage()
	if __obf_d34a61736c704de8 != nil {
		return nil, __obf_d34a61736c704de8
	}

	return __obf_660839016fa371e0.DrawWithNRGBA(__obf_933380cc9c059699, __obf_38a73a0a94d0be23)
}

func (__obf_4d5cf3f253a7202e *Captcha) __obf_82b741c7da7e2910(__obf_933380cc9c059699 DotType, __obf_e0f3c13244d13853 core.Size, __obf_566e0cc4a480c2a3 []*Dot) (image.Image, error) {
	var (
		// err      error
		__obf_416c0f47b234378a = make([]*Dot, 0, len(__obf_566e0cc4a480c2a3))
	)

	__obf_99216d7bb578abfc := __obf_e0f3c13244d13853.Width / len(__obf_566e0cc4a480c2a3)

	for __obf_03cbab251fc402a6, __obf_c7a6b4d6583596e4 := range __obf_566e0cc4a480c2a3 {
		__obf_45135e7749d3529c := 1
		if __obf_933380cc9c059699 == Text {
			__obf_45135e7749d3529c = len(__obf_c7a6b4d6583596e4.Text)
		}

		__obf_c7a6b4d6583596e4.X = int(max(float64(__obf_99216d7bb578abfc*__obf_03cbab251fc402a6+__obf_99216d7bb578abfc/__obf_c7a6b4d6583596e4.Size), 8))
		__obf_61692a9a83b78472 := max(1, __obf_e0f3c13244d13853.Height/16*__obf_45135e7749d3529c)
		__obf_c7a6b4d6583596e4.Y = __obf_e0f3c13244d13853.Height/2 + __obf_c7a6b4d6583596e4.FontSize/2 - rand.IntN(__obf_61692a9a83b78472)

		__obf_416c0f47b234378a = append(__obf_416c0f47b234378a, __obf_c7a6b4d6583596e4)
	}

	__obf_acb4c1626c979718 := &DrawImageParams{
		Width:             __obf_e0f3c13244d13853.Width,
		Height:            __obf_e0f3c13244d13853.Height,
		Dots:              __obf_416c0f47b234378a,
		BackgroundDistort: __obf_4d5cf3f253a7202e.__obf_a9bfdd8861059e32(__obf_4d5cf3f253a7202e.__obf_932c987f5e543587.Secondary.BgDistort),
		BgCircles:         __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.Secondary.BgCircles,
		BgSlimLines:       __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.Secondary.BgSlimLines,
		ThumbDisturbAlpha: __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.Secondary.DisturbAlpha,
	}

	// params.Background, err = core.PickThumbImage()
	// if err != nil {
	// 	return nil, err
	// }

	var __obf_6f06cb2de9b657f0 []color.Color
	for _, __obf_7eb33b2f30faa119 := range __obf_4c1e9766248eece8 {
		__obf_94d30d5f9cee36ca, _ := core.ParseHexColor(__obf_7eb33b2f30faa119)
		__obf_6f06cb2de9b657f0 = append(__obf_6f06cb2de9b657f0, __obf_94d30d5f9cee36ca)
	}

	var __obf_9c8a549620ddfbca []color.Color
	for _, __obf_94d30d5f9cee36ca := range __obf_4c1e9766248eece8 {
		__obf_8fbcfb8ce52f034a, _ := core.ParseHexColor(__obf_94d30d5f9cee36ca)
		__obf_9c8a549620ddfbca = append(__obf_9c8a549620ddfbca, __obf_8fbcfb8ce52f034a)
	}

	if __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.UseRGBA || __obf_4d5cf3f253a7202e.__obf_932c987f5e543587.Secondary.NonDeformAble {
		return __obf_4d5cf3f253a7202e.DrawWithNRGBA2(__obf_933380cc9c059699, __obf_acb4c1626c979718, __obf_6f06cb2de9b657f0, __obf_9c8a549620ddfbca)
	}
	return __obf_4d5cf3f253a7202e.DrawWithPalette(__obf_933380cc9c059699, __obf_acb4c1626c979718, __obf_6f06cb2de9b657f0, __obf_9c8a549620ddfbca)
}

func (__obf_660839016fa371e0 *Captcha) __obf_a9bfdd8861059e32(__obf_666a04afe72c9237 int) int {
	switch __obf_666a04afe72c9237 {
	case 1:
		return core.RandInt(240, 320)
	case 2:
		return core.RandInt(180, 240)
	case 3:
		return core.RandInt(120, 180)
	case 4:
		return core.RandInt(100, 160)
	case 5:
		return core.RandInt(80, 140)
	}
	return 0
}

func (__obf_660839016fa371e0 *Captcha) __obf_77015c656c3b620a() int {
	__obf_da8a8a2fedc89870 := __obf_660839016fa371e0.__obf_932c987f5e543587.Primary.AnglePosRange

	__obf_f4481368e8c8984c := rand.IntN(len(__obf_da8a8a2fedc89870))
	if __obf_f4481368e8c8984c < 0 {
		return 0
	}

	__obf_4a672f1292cd0aac := __obf_da8a8a2fedc89870[__obf_f4481368e8c8984c]
	__obf_9f376ce374f84518 := core.RandInt(__obf_4a672f1292cd0aac.Min, __obf_4a672f1292cd0aac.Max)

	return __obf_9f376ce374f84518
}
