package __obf_771d454808f56d40

import (
	"image"
	"math/rand/v2"

	assets "github.com/ArtisanHiram/go-pkg/captcha/internal/assets"
	helper "github.com/ArtisanHiram/go-pkg/captcha/internal/helper"
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
)

// Captcha is the concrete implementation of the Captcha interface
type Captcha struct {
	__obf_b38a8327e4592b39 *Options
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_b38a8327e4592b39 *Options) *Captcha {
	if __obf_b38a8327e4592b39 == nil {
		return nil
	}
	return &Captcha{
		__obf_b38a8327e4592b39: __obf_b38a8327e4592b39,
	}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_19428d8b430755cc *Captcha) DrawWithNRGBA(__obf_555dc0c8dcbf92e5 *DrawImageParams) (image.Image, error) {
	__obf_932eac91fc9ec491 := helper.CreateNRGBACanvas(__obf_555dc0c8dcbf92e5.Size, __obf_555dc0c8dcbf92e5.Size, true)
	if __obf_555dc0c8dcbf92e5.Background != nil {
		__obf_7ef55de814b2b183 := __obf_555dc0c8dcbf92e5.Background
		__obf_538c6908d1f234af := __obf_7ef55de814b2b183.Bounds()
		__obf_78b61695d6ae2975 := helper.CreateNRGBACanvas(__obf_538c6908d1f234af.Dx(), __obf_538c6908d1f234af.Dy(), true)
		__obf_6af623ffab4f9874 := util.RangCutImagePos(__obf_555dc0c8dcbf92e5.Size, __obf_555dc0c8dcbf92e5.Size, __obf_7ef55de814b2b183)
		draw.Draw(__obf_78b61695d6ae2975, __obf_538c6908d1f234af, __obf_7ef55de814b2b183, __obf_6af623ffab4f9874, draw.Over)
		__obf_78b61695d6ae2975.SubImage(image.Rect(0, 0, __obf_555dc0c8dcbf92e5.Size, __obf_555dc0c8dcbf92e5.Size))
		draw.Draw(__obf_932eac91fc9ec491, __obf_932eac91fc9ec491.Bounds(), __obf_78b61695d6ae2975, image.Point{}, draw.Over)
	}

	__obf_932eac91fc9ec491.CropCircle(__obf_932eac91fc9ec491.Bounds().Dx()/2, __obf_932eac91fc9ec491.Bounds().Dy()/2, __obf_932eac91fc9ec491.Bounds().Dy()/2)
	return __obf_932eac91fc9ec491, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_19428d8b430755cc *Captcha) DrawWithCropCircle(__obf_555dc0c8dcbf92e5 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_7ef55de814b2b183 := __obf_555dc0c8dcbf92e5.Background

	__obf_71d2af6a925b4186 := __obf_7ef55de814b2b183.Bounds()
	__obf_325fde44bcf14803 := helper.CreateNRGBACanvas(__obf_7ef55de814b2b183.Bounds().Dx(), __obf_7ef55de814b2b183.Bounds().Dy(), true)
	draw.Draw(__obf_325fde44bcf14803, __obf_7ef55de814b2b183.Bounds(), __obf_7ef55de814b2b183, image.Point{}, draw.Over)
	__obf_325fde44bcf14803.CropScaleCircle(__obf_7ef55de814b2b183.Bounds().Dx()/2, __obf_7ef55de814b2b183.Bounds().Dy()/2, __obf_7ef55de814b2b183.Bounds().Dy()/2, __obf_555dc0c8dcbf92e5.ScaleRatioSize)
	__obf_325fde44bcf14803.Rotate(__obf_555dc0c8dcbf92e5.Angle, true)

	__obf_2c59054e289fdeae := __obf_325fde44bcf14803.Bounds()
	if __obf_2c59054e289fdeae.Dy() > __obf_71d2af6a925b4186.Dy() || __obf_2c59054e289fdeae.Dx() > __obf_71d2af6a925b4186.Dx() {
		__obf_6043ea1904b6aa95 := helper.CreateNRGBACanvas(__obf_7ef55de814b2b183.Bounds().Dx(), __obf_7ef55de814b2b183.Bounds().Dy(), true)
		draw.Draw(__obf_6043ea1904b6aa95, __obf_6043ea1904b6aa95.Bounds(), __obf_325fde44bcf14803, image.Point{X: (__obf_2c59054e289fdeae.Dy() - __obf_71d2af6a925b4186.Dy()) / 2, Y: (__obf_2c59054e289fdeae.Dx() - __obf_71d2af6a925b4186.Dx()) / 2}, draw.Over)
		__obf_325fde44bcf14803 = __obf_6043ea1904b6aa95
	}

	return __obf_325fde44bcf14803, nil
}

// Generate generates rotate CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_19428d8b430755cc *Captcha) Generate() (*CaptchaData, error) {
	__obf_72094832975d8395 := __obf_19428d8b430755cc.__obf_69a8f7dcc86285cf()

	__obf_67ce4ac668bc5d62, __obf_b92ece9497e7c4fd := __obf_19428d8b430755cc.__obf_ef1c4945120932c4(__obf_19428d8b430755cc.__obf_b38a8327e4592b39.Primary.Size)
	if __obf_b92ece9497e7c4fd != nil {
		return nil, __obf_b92ece9497e7c4fd
	}
	__obf_b4468f4140748fac := __obf_19428d8b430755cc.__obf_b8cdac8a8fc4f3f6()
	__obf_15e7c5e530990c10, __obf_b92ece9497e7c4fd := __obf_19428d8b430755cc.__obf_8aae30f7c8445765(__obf_67ce4ac668bc5d62, __obf_b4468f4140748fac, __obf_72094832975d8395)
	if __obf_b92ece9497e7c4fd != nil {
		return nil, __obf_b92ece9497e7c4fd
	}

	return &CaptchaData{
		__obf_b4468f4140748fac: __obf_b4468f4140748fac,
		__obf_67ce4ac668bc5d62: types.NewPNGImage(__obf_67ce4ac668bc5d62),
		__obf_15e7c5e530990c10: types.NewPNGImage(__obf_15e7c5e530990c10),
	}, nil
}

// genMainImage generates the main CAPTCHA image
// params:
//   - size: Image size
//   - block: Block data
//
// returns:
//   - image.Image: Generated main image
//   - error: Error information
func (__obf_19428d8b430755cc *Captcha) __obf_ef1c4945120932c4(__obf_1c22f398199704ef int) (image.Image, error) {

	__obf_a8d99c9ce0de0da2, __obf_b92ece9497e7c4fd := assets.PickBgImage()
	if __obf_b92ece9497e7c4fd != nil {
		return nil, __obf_b92ece9497e7c4fd
	}

	return __obf_19428d8b430755cc.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_1c22f398199704ef,
		Background: __obf_a8d99c9ce0de0da2,
		Alpha:      __obf_19428d8b430755cc.__obf_b38a8327e4592b39.Primary.Alpha,
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
func (__obf_19428d8b430755cc *Captcha) __obf_8aae30f7c8445765(__obf_7ef55de814b2b183 image.Image, __obf_b4468f4140748fac int, __obf_72094832975d8395 int) (image.Image, error) {
	return __obf_19428d8b430755cc.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_7ef55de814b2b183,
		Alpha:          __obf_19428d8b430755cc.__obf_b38a8327e4592b39.Secondary.Alpha,
		Size:           __obf_72094832975d8395,
		Angle:          __obf_b4468f4140748fac,
		ScaleRatioSize: (__obf_19428d8b430755cc.__obf_b38a8327e4592b39.Primary.Size - __obf_72094832975d8395) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_19428d8b430755cc *Captcha) __obf_b8cdac8a8fc4f3f6() int {
	__obf_bda166943ac4345b := __obf_19428d8b430755cc.__obf_b38a8327e4592b39.Primary.AnglePosRange

	__obf_7d4b554e40b84152 := rand.IntN(len(__obf_bda166943ac4345b))
	if __obf_7d4b554e40b84152 < 0 {
		return 0
	}

	__obf_b4468f4140748fac := __obf_bda166943ac4345b[__obf_7d4b554e40b84152]
	return util.RandInt(__obf_b4468f4140748fac.Min, __obf_b4468f4140748fac.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_19428d8b430755cc *Captcha) __obf_69a8f7dcc86285cf() int {
	__obf_c99f4e9d48b02248 := __obf_19428d8b430755cc.__obf_b38a8327e4592b39.Secondary.SizeRange

	__obf_7d4b554e40b84152 := rand.IntN(len(__obf_c99f4e9d48b02248))
	if __obf_7d4b554e40b84152 < 0 {
		return __obf_c99f4e9d48b02248[0]
	}

	return __obf_c99f4e9d48b02248[__obf_7d4b554e40b84152]
}
