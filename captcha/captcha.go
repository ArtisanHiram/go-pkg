package __obf_d44072021ce2eea3

import (
	click "github.com/ArtisanHiram/go-pkg/captcha/internal/click"
	rotate "github.com/ArtisanHiram/go-pkg/captcha/internal/rotate"
	slide "github.com/ArtisanHiram/go-pkg/captcha/internal/slide"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"

	"golang.org/x/image/font"
)

func NewClickCaptcha(__obf_ee83a7adf3a0a96f click.SetOption) *click.Captcha {
	__obf_0bf0438c9dd559a3 := &types.ClickOption{
		FontDPI:     72, // Increased to 200 for high quality rendering
		FontHinting: font.HintingNone,
		Primary: types.ClickPrimary{
			Alpha:      1,
			DotPadding: 10,
			Width:      400,
			Height:     340,
			LenRange:   types.Range{Min: 3, Max: 6},
			SizeRange:  types.Range{Min: 32, Max: 42},
			AnglePosRange: []types.Range{
				{Min: 20, Max: 35},
				{Min: 35, Max: 45},
				{Min: 45, Max: 60},
				{Min: 290, Max: 305},
				{Min: 305, Max: 325},
				{Min: 325, Max: 330},
			},
		},
		ShowShadow:  true,
		ShadowPoint: types.Point{X: -1, Y: -1},
		Secondary: types.ClickSecondary{
			BgDistort:      types.DistortLevel4,
			BgCircles:      4,
			BgSlimLines:    2,
			DisturbAlpha:   1,
			DotPadding:     0,
			NonDeformAble:  true,
			VerifyLenRange: types.Range{Min: 2, Max: 4},
			Width:          150,
			Height:         60,
			SizeRange:      types.Range{Min: 32, Max: 42},
		},
	}

	if __obf_ee83a7adf3a0a96f != nil {
		__obf_ee83a7adf3a0a96f(__obf_0bf0438c9dd559a3)
	}

	return click.NewCaptcha(__obf_0bf0438c9dd559a3)
}

func NewRotateCaptcha(__obf_ee83a7adf3a0a96f rotate.SetOption) *rotate.Captcha {
	__obf_0bf0438c9dd559a3 := &types.RotateOption{
		Primary: types.RotatePrimary{
			Size:  220,
			Alpha: 1,
			AnglePosRange: []types.Range{
				{Min: 30, Max: 330},
			},
		},
		Secondary: types.RotateSecondary{
			Alpha:     1,
			SizeRange: []int{140, 150, 160, 170},
		},
	}

	if __obf_ee83a7adf3a0a96f != nil {
		__obf_ee83a7adf3a0a96f(__obf_0bf0438c9dd559a3)
	}
	return rotate.NewCaptcha(__obf_0bf0438c9dd559a3)
}

func NewMoveCaptcha(__obf_ee83a7adf3a0a96f slide.SetOption) *slide.Captcha {
	__obf_0bf0438c9dd559a3 := &types.SlideOption{
		Type: types.MoveSlide,
		Primary: types.SlidePrimary{
			Width:  400,
			Height: 340,
			Alpha:  1,
		},
		Secondary: types.SlideSecondary{
			EnableVerticalRandom: false,
			CountRange:           types.Range{Min: 1, Max: 5},
			SizeRange:            types.Range{Min: 60, Max: 70},
			AnglePosRange: []types.Range{
				{Min: 0, Max: 0},
			},
			DeadZoneDirections: []types.DeadZoneDirectionType{types.DeadZoneDirectionTypeLeft},
		},
	}

	if __obf_ee83a7adf3a0a96f != nil {
		__obf_ee83a7adf3a0a96f(__obf_0bf0438c9dd559a3)
	}

	return slide.NewCaptcha(__obf_0bf0438c9dd559a3)

}

func NewDragCaptcha(__obf_ee83a7adf3a0a96f slide.SetOption) *slide.Captcha {
	__obf_0bf0438c9dd559a3 := &types.SlideOption{
		Type: types.DragSlide,
		Primary: types.SlidePrimary{
			Width:  400,
			Height: 340,
			Alpha:  1,
		},
		Secondary: types.SlideSecondary{
			EnableVerticalRandom: true,
			CountRange:           types.Range{Min: 1, Max: 5},
			SizeRange:            types.Range{Min: 60, Max: 70},
			AnglePosRange: []types.Range{
				{Min: 0, Max: 0},
			},
			DeadZoneDirections: []types.DeadZoneDirectionType{
				types.DeadZoneDirectionTypeLeft,
				types.DeadZoneDirectionTypeRight,
				types.DeadZoneDirectionTypeBottom,
				types.DeadZoneDirectionTypeTop,
			},
		},
	}

	if __obf_ee83a7adf3a0a96f != nil {
		__obf_ee83a7adf3a0a96f(__obf_0bf0438c9dd559a3)
	}

	return slide.NewCaptcha(__obf_0bf0438c9dd559a3)
}
