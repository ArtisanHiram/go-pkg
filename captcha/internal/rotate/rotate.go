package __obf_f99ff5685a4c2bb7

import (
	assets "github.com/ArtisanHiram/go-pkg/captcha/internal/assets"
	helper "github.com/ArtisanHiram/go-pkg/captcha/internal/helper"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"image"
	"math/rand/v2"
)

// Captcha is the concrete implementation of the Captcha interface
type Captcha struct {
	__obf_8a6631ca9e4be277 *types.RotateOption
}

var _ types.Captcha = (*Captcha)(nil)

func (__obf_1982ac2195a4989f *Captcha) Generate() (types.CaptchaData, error) {
	__obf_c872f6e59e693aa9 := __obf_1982ac2195a4989f.__obf_c7f0c078b553c35b()
	__obf_a9d22809133c29eb, __obf_fa33f2c62ad8ba58 := __obf_1982ac2195a4989f.__obf_45253d4d37d091a6(__obf_1982ac2195a4989f.__obf_8a6631ca9e4be277.Primary.Size)
	if __obf_fa33f2c62ad8ba58 != nil {
		return nil, __obf_fa33f2c62ad8ba58
	}
	__obf_7deb1b8e84658d4f := __obf_1982ac2195a4989f.__obf_8d2920431b6e1c97()
	__obf_97e26f7838d6d222, __obf_fa33f2c62ad8ba58 := __obf_1982ac2195a4989f.__obf_d183f45751e501d7(__obf_a9d22809133c29eb, __obf_7deb1b8e84658d4f, __obf_c872f6e59e693aa9)
	if __obf_fa33f2c62ad8ba58 != nil {
		return nil, __obf_fa33f2c62ad8ba58
	}

	return &CaptchaData{__obf_7deb1b8e84658d4f: __obf_7deb1b8e84658d4f, __obf_a9d22809133c29eb: types.NewPNGImage(__obf_a9d22809133c29eb), __obf_97e26f7838d6d222: types.NewPNGImage(__obf_97e26f7838d6d222)}, nil
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_8a6631ca9e4be277 *types.RotateOption) *Captcha {
	if __obf_8a6631ca9e4be277 == nil {
		return nil
	}
	return &Captcha{__obf_8a6631ca9e4be277: __obf_8a6631ca9e4be277}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_1982ac2195a4989f *Captcha) DrawWithNRGBA(__obf_48dc9af23d56b87a *DrawImageParams) (image.Image, error) {
	__obf_1d9a6a33ebc13707 := helper.CreateNRGBACanvas(__obf_48dc9af23d56b87a.Size, __obf_48dc9af23d56b87a.Size, true)
	if __obf_48dc9af23d56b87a.Background != nil {
		__obf_44238cbe3d4ed6dc := __obf_48dc9af23d56b87a.Background
		__obf_1f2a0ee2d2231451 := __obf_44238cbe3d4ed6dc.Bounds()
		__obf_4eaab3870182c6aa := helper.CreateNRGBACanvas(__obf_1f2a0ee2d2231451.Dx(), __obf_1f2a0ee2d2231451.Dy(), true)
		__obf_f738e9f567ba7464 := util.RangCutImagePos(__obf_48dc9af23d56b87a.Size, __obf_48dc9af23d56b87a.Size, __obf_44238cbe3d4ed6dc)
		draw.Draw(__obf_4eaab3870182c6aa, __obf_1f2a0ee2d2231451, __obf_44238cbe3d4ed6dc, __obf_f738e9f567ba7464, draw.Over)
		__obf_4eaab3870182c6aa.
			SubImage(image.Rect(0, 0, __obf_48dc9af23d56b87a.Size, __obf_48dc9af23d56b87a.Size))
		draw.Draw(__obf_1d9a6a33ebc13707, __obf_1d9a6a33ebc13707.Bounds(), __obf_4eaab3870182c6aa, image.Point{}, draw.Over)
	}
	__obf_1d9a6a33ebc13707.
		CropCircle(__obf_1d9a6a33ebc13707.Bounds().Dx()/2, __obf_1d9a6a33ebc13707.Bounds().Dy()/2, __obf_1d9a6a33ebc13707.Bounds().Dy()/2)
	return __obf_1d9a6a33ebc13707, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_1982ac2195a4989f *Captcha) DrawWithCropCircle(__obf_48dc9af23d56b87a *DrawCropCircleImageParams) (image.Image, error) {
	__obf_44238cbe3d4ed6dc := __obf_48dc9af23d56b87a.Background
	__obf_58edfc8625ff78e8 := __obf_44238cbe3d4ed6dc.Bounds()
	__obf_2f4d8b34559c5ab6 := helper.CreateNRGBACanvas(__obf_44238cbe3d4ed6dc.Bounds().Dx(), __obf_44238cbe3d4ed6dc.Bounds().Dy(), true)
	draw.Draw(__obf_2f4d8b34559c5ab6, __obf_44238cbe3d4ed6dc.Bounds(), __obf_44238cbe3d4ed6dc, image.Point{}, draw.Over)
	__obf_2f4d8b34559c5ab6.
		CropScaleCircle(__obf_44238cbe3d4ed6dc.Bounds().Dx()/2, __obf_44238cbe3d4ed6dc.Bounds().Dy()/2, __obf_44238cbe3d4ed6dc.Bounds().Dy()/2, __obf_48dc9af23d56b87a.ScaleRatioSize)
	__obf_2f4d8b34559c5ab6.
		Rotate(__obf_48dc9af23d56b87a.Angle, true)
	__obf_32705c3035d022a7 := __obf_2f4d8b34559c5ab6.Bounds()
	if __obf_32705c3035d022a7.Dy() > __obf_58edfc8625ff78e8.Dy() || __obf_32705c3035d022a7.Dx() > __obf_58edfc8625ff78e8.Dx() {
		__obf_26ec2bb59a06cbfe := helper.CreateNRGBACanvas(__obf_44238cbe3d4ed6dc.Bounds().Dx(), __obf_44238cbe3d4ed6dc.Bounds().Dy(), true)
		draw.Draw(__obf_26ec2bb59a06cbfe, __obf_26ec2bb59a06cbfe.Bounds(), __obf_2f4d8b34559c5ab6, image.Point{X: (__obf_32705c3035d022a7.Dy() - __obf_58edfc8625ff78e8.Dy()) / 2, Y: (__obf_32705c3035d022a7.Dx() - __obf_58edfc8625ff78e8.Dx()) / 2}, draw.Over)
		__obf_2f4d8b34559c5ab6 = __obf_26ec2bb59a06cbfe
	}

	return __obf_2f4d8b34559c5ab6, nil
}

func (__obf_1982ac2195a4989f *Captcha) __obf_45253d4d37d091a6(__obf_5ea40ef879982da6 int) (image.Image, error) {
	__obf_e2a7b05e21479368, __obf_fa33f2c62ad8ba58 := assets.PickBgImage()
	if __obf_fa33f2c62ad8ba58 != nil {
		return nil, __obf_fa33f2c62ad8ba58
	}

	return __obf_1982ac2195a4989f.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_5ea40ef879982da6,
		Background: __obf_e2a7b05e21479368,
		Alpha:      __obf_1982ac2195a4989f.__obf_8a6631ca9e4be277.Primary.Alpha,
	})
}

// genThumbImage generates a thumbnail image
// params:
//   - bgImage: Background image
//   - block: Block data
//   - thumbSize: Thumbnail size
//
// returns:
//   - image.Image: Generated thumbnail image
//   - error: Error information
func (__obf_1982ac2195a4989f *Captcha) __obf_d183f45751e501d7(__obf_44238cbe3d4ed6dc image.Image, __obf_7deb1b8e84658d4f int, __obf_c872f6e59e693aa9 int) (image.Image, error) {
	return __obf_1982ac2195a4989f.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_44238cbe3d4ed6dc,
		Alpha:          __obf_1982ac2195a4989f.__obf_8a6631ca9e4be277.Secondary.Alpha,
		Size:           __obf_c872f6e59e693aa9,
		Angle:          __obf_7deb1b8e84658d4f,
		ScaleRatioSize: (__obf_1982ac2195a4989f.__obf_8a6631ca9e4be277.Primary.Size - __obf_c872f6e59e693aa9) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_1982ac2195a4989f *Captcha) __obf_8d2920431b6e1c97() int {
	__obf_c5a2173c083c24d0 := __obf_1982ac2195a4989f.__obf_8a6631ca9e4be277.Primary.AnglePosRange
	__obf_ee59ce717a7f7fde := rand.IntN(len(__obf_c5a2173c083c24d0))
	if __obf_ee59ce717a7f7fde < 0 {
		return 0
	}
	__obf_7deb1b8e84658d4f := __obf_c5a2173c083c24d0[__obf_ee59ce717a7f7fde]
	return util.RandInt(__obf_7deb1b8e84658d4f.Min, __obf_7deb1b8e84658d4f.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_1982ac2195a4989f *Captcha) __obf_c7f0c078b553c35b() int {
	__obf_eb76db729bce9451 := __obf_1982ac2195a4989f.__obf_8a6631ca9e4be277.Secondary.SizeRange
	__obf_ee59ce717a7f7fde := rand.IntN(len(__obf_eb76db729bce9451))
	if __obf_ee59ce717a7f7fde < 0 {
		return __obf_eb76db729bce9451[0]
	}

	return __obf_eb76db729bce9451[__obf_ee59ce717a7f7fde]
}
