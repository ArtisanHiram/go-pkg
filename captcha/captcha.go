package __obf_236715d8b478ef4c

import (
	click "github.com/ArtisanHiram/go-pkg/captcha/internal/click"
	rotate "github.com/ArtisanHiram/go-pkg/captcha/internal/rotate"
	slide "github.com/ArtisanHiram/go-pkg/captcha/internal/slide"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"

	"golang.org/x/image/font"
)

func NewClickCaptcha(__obf_9e32f54e3dde6145 click.SetOption) *click.Captcha {
	__obf_1de1c802e51e7dd9 := &types.ClickOption{
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

	if __obf_9e32f54e3dde6145 != nil {
		__obf_9e32f54e3dde6145(__obf_1de1c802e51e7dd9)
	}

	return click.NewCaptcha(__obf_1de1c802e51e7dd9)
}

func NewRotateCaptcha(__obf_9e32f54e3dde6145 rotate.SetOption) *rotate.Captcha {
	__obf_1de1c802e51e7dd9 := &types.RotateOption{
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

	if __obf_9e32f54e3dde6145 != nil {
		__obf_9e32f54e3dde6145(__obf_1de1c802e51e7dd9)
	}
	return rotate.NewCaptcha(__obf_1de1c802e51e7dd9)
}

func NewMoveCaptcha(__obf_9e32f54e3dde6145 slide.SetOption) *slide.Captcha {
	__obf_1de1c802e51e7dd9 := &types.SlideOption{
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

	if __obf_9e32f54e3dde6145 != nil {
		__obf_9e32f54e3dde6145(__obf_1de1c802e51e7dd9)
	}

	return slide.NewCaptcha(__obf_1de1c802e51e7dd9)

}

func NewDragCaptcha(__obf_9e32f54e3dde6145 slide.SetOption) *slide.Captcha {
	__obf_1de1c802e51e7dd9 := &types.SlideOption{
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

	if __obf_9e32f54e3dde6145 != nil {
		__obf_9e32f54e3dde6145(__obf_1de1c802e51e7dd9)
	}

	return slide.NewCaptcha(__obf_1de1c802e51e7dd9)
}
