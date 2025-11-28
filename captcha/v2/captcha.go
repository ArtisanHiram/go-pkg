package __obf_f97d24288011c5b2

import (
	click "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/click"
	rotate "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/rotate"
	slide "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/slide"
	types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"

	"golang.org/x/image/font"
)

func NewClickCaptcha(__obf_2adc72f3691a20ab click.Option) *click.Captcha {
	__obf_958e808748aff430 := &click.Options{
		FontDPI:     72, // Increased to 200 for high quality rendering
		FontHinting: font.HintingNone,
		Primary: click.Primary{
			Alpha:      1,
			DotPadding: 10,
			Size:       types.Size{Width: 400, Height: 340},
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
		Secondary: click.Secondary{
			BgDistort:      types.DistortLevel4,
			BgCircles:      4,
			BgSlimLines:    2,
			DisturbAlpha:   1,
			DotPadding:     0,
			NonDeformAble:  true,
			VerifyLenRange: types.Range{Min: 2, Max: 4},
			Size:           types.Size{Width: 150, Height: 60},
			SizeRange:      types.Range{Min: 32, Max: 42},
		},
	}

	if __obf_2adc72f3691a20ab != nil {
		__obf_2adc72f3691a20ab(__obf_958e808748aff430)
	}

	return click.NewCaptcha(__obf_958e808748aff430)
}

func NewRotateCaptcha(__obf_2adc72f3691a20ab rotate.Option) *rotate.Captcha {
	__obf_958e808748aff430 := &rotate.Options{
		Primary: rotate.Primary{
			Size:  220,
			Alpha: 1,
			AnglePosRange: []types.Range{
				{Min: 30, Max: 330},
			},
		},
		Secondary: rotate.Secondary{
			Alpha:     1,
			SizeRange: []int{140, 150, 160, 170},
		},
	}

	if __obf_2adc72f3691a20ab != nil {
		__obf_2adc72f3691a20ab(__obf_958e808748aff430)
	}
	return rotate.NewCaptcha(__obf_958e808748aff430)
}

func NewMoveCaptcha(__obf_2adc72f3691a20ab slide.Option) *slide.Captcha {
	__obf_958e808748aff430 := &slide.Options{
		Type: slide.Move,
		Primary: slide.Primary{
			Size:  types.Size{Width: 400, Height: 340},
			Alpha: 1,
		},
		Secondary: slide.Secondary{
			EnableVerticalRandom: false,
			CountRange:           types.Range{Min: 1, Max: 5},
			SizeRange:            types.Range{Min: 60, Max: 70},
			AnglePosRange: []types.Range{
				{Min: 0, Max: 0},
			},
			DeadZoneDirections: []slide.DeadZoneDirectionType{slide.DeadZoneDirectionTypeLeft},
		},
	}

	if __obf_2adc72f3691a20ab != nil {
		__obf_2adc72f3691a20ab(__obf_958e808748aff430)
	}

	return slide.NewCaptcha(__obf_958e808748aff430)

}

func NewDragCaptcha(__obf_2adc72f3691a20ab slide.Option) *slide.Captcha {
	__obf_958e808748aff430 := &slide.Options{
		Type: slide.Drag,
		Primary: slide.Primary{
			Size:  types.Size{Width: 400, Height: 340},
			Alpha: 1,
		},
		Secondary: slide.Secondary{
			EnableVerticalRandom: true,
			CountRange:           types.Range{Min: 1, Max: 5},
			SizeRange:            types.Range{Min: 60, Max: 70},
			AnglePosRange: []types.Range{
				{Min: 0, Max: 0},
			},
			DeadZoneDirections: []slide.DeadZoneDirectionType{
				slide.DeadZoneDirectionTypeLeft,
				slide.DeadZoneDirectionTypeRight,
				slide.DeadZoneDirectionTypeBottom,
				slide.DeadZoneDirectionTypeTop,
			},
		},
	}

	if __obf_2adc72f3691a20ab != nil {
		__obf_2adc72f3691a20ab(__obf_958e808748aff430)
	}

	return slide.NewCaptcha(__obf_958e808748aff430)
}
