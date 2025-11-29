package __obf_9833fbfb1ccf920c

import (
	click "github.com/ArtisanHiram/go-pkg/captcha/internal/click"
	rotate "github.com/ArtisanHiram/go-pkg/captcha/internal/rotate"
	slide "github.com/ArtisanHiram/go-pkg/captcha/internal/slide"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"

	"golang.org/x/image/font"
)

func NewClickCaptcha(__obf_712cabc91cd7b883 click.Option) *click.Captcha {
	__obf_1fde67f2789fcf6a := &click.Options{
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

	if __obf_712cabc91cd7b883 != nil {
		__obf_712cabc91cd7b883(__obf_1fde67f2789fcf6a)
	}

	return click.NewCaptcha(__obf_1fde67f2789fcf6a)
}

func NewRotateCaptcha(__obf_712cabc91cd7b883 rotate.Option) *rotate.Captcha {
	__obf_1fde67f2789fcf6a := &rotate.Options{
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

	if __obf_712cabc91cd7b883 != nil {
		__obf_712cabc91cd7b883(__obf_1fde67f2789fcf6a)
	}
	return rotate.NewCaptcha(__obf_1fde67f2789fcf6a)
}

func NewMoveCaptcha(__obf_712cabc91cd7b883 slide.Option) *slide.Captcha {
	__obf_1fde67f2789fcf6a := &slide.Options{
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

	if __obf_712cabc91cd7b883 != nil {
		__obf_712cabc91cd7b883(__obf_1fde67f2789fcf6a)
	}

	return slide.NewCaptcha(__obf_1fde67f2789fcf6a)

}

func NewDragCaptcha(__obf_712cabc91cd7b883 slide.Option) *slide.Captcha {
	__obf_1fde67f2789fcf6a := &slide.Options{
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

	if __obf_712cabc91cd7b883 != nil {
		__obf_712cabc91cd7b883(__obf_1fde67f2789fcf6a)
	}

	return slide.NewCaptcha(__obf_1fde67f2789fcf6a)
}
