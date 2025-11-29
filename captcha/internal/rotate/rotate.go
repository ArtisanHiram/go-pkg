package __obf_e017a70acec4eb68

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
	__obf_5ace98d4b75065f4 *Options
}

var _ types.Captcha = (*Captcha)(nil)

func (__obf_3e760783dbf49c5a *Captcha) Generate() (types.CaptchaData, error) {
	__obf_c6d635faa67c7ae6 := __obf_3e760783dbf49c5a.__obf_d4115b9deb1ede70()
	__obf_500547421fb3df93, __obf_b4c4cf14feeaf592 := __obf_3e760783dbf49c5a.__obf_88d5fed3bc1bff4f(__obf_3e760783dbf49c5a.__obf_5ace98d4b75065f4.Primary.Size)
	if __obf_b4c4cf14feeaf592 != nil {
		return nil, __obf_b4c4cf14feeaf592
	}
	__obf_f9c9a01651aaf15e := __obf_3e760783dbf49c5a.__obf_0ceadbaf6b084be6()
	__obf_a5b5318e5ebc7a2c, __obf_b4c4cf14feeaf592 := __obf_3e760783dbf49c5a.__obf_d3ecd9ba7934c34b(__obf_500547421fb3df93, __obf_f9c9a01651aaf15e, __obf_c6d635faa67c7ae6)
	if __obf_b4c4cf14feeaf592 != nil {
		return nil, __obf_b4c4cf14feeaf592
	}

	return &CaptchaData{__obf_f9c9a01651aaf15e: __obf_f9c9a01651aaf15e, __obf_500547421fb3df93: types.NewPNGImage(__obf_500547421fb3df93), __obf_a5b5318e5ebc7a2c: types.NewPNGImage(__obf_a5b5318e5ebc7a2c)}, nil
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_5ace98d4b75065f4 *Options) *Captcha {
	if __obf_5ace98d4b75065f4 == nil {
		return nil
	}
	return &Captcha{__obf_5ace98d4b75065f4: __obf_5ace98d4b75065f4}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_3e760783dbf49c5a *Captcha) DrawWithNRGBA(__obf_fe92b35fa565dcef *DrawImageParams) (image.Image, error) {
	__obf_a2cfee6de1551b17 := helper.CreateNRGBACanvas(__obf_fe92b35fa565dcef.Size, __obf_fe92b35fa565dcef.Size, true)
	if __obf_fe92b35fa565dcef.Background != nil {
		__obf_0594f054c57944cf := __obf_fe92b35fa565dcef.Background
		__obf_997fc9c4bb601eca := __obf_0594f054c57944cf.Bounds()
		__obf_acd960368e219487 := helper.CreateNRGBACanvas(__obf_997fc9c4bb601eca.Dx(), __obf_997fc9c4bb601eca.Dy(), true)
		__obf_2e40e8e2dcab5b83 := util.RangCutImagePos(__obf_fe92b35fa565dcef.Size, __obf_fe92b35fa565dcef.Size, __obf_0594f054c57944cf)
		draw.Draw(__obf_acd960368e219487, __obf_997fc9c4bb601eca, __obf_0594f054c57944cf, __obf_2e40e8e2dcab5b83, draw.Over)
		__obf_acd960368e219487.
			SubImage(image.Rect(0, 0, __obf_fe92b35fa565dcef.Size, __obf_fe92b35fa565dcef.Size))
		draw.Draw(__obf_a2cfee6de1551b17, __obf_a2cfee6de1551b17.Bounds(), __obf_acd960368e219487, image.Point{}, draw.Over)
	}
	__obf_a2cfee6de1551b17.
		CropCircle(__obf_a2cfee6de1551b17.Bounds().Dx()/2, __obf_a2cfee6de1551b17.Bounds().Dy()/2, __obf_a2cfee6de1551b17.Bounds().Dy()/2)
	return __obf_a2cfee6de1551b17, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_3e760783dbf49c5a *Captcha) DrawWithCropCircle(__obf_fe92b35fa565dcef *DrawCropCircleImageParams) (image.Image, error) {
	__obf_0594f054c57944cf := __obf_fe92b35fa565dcef.Background
	__obf_31d7cb2a7f8f7e87 := __obf_0594f054c57944cf.Bounds()
	__obf_a96f03e430996c71 := helper.CreateNRGBACanvas(__obf_0594f054c57944cf.Bounds().Dx(), __obf_0594f054c57944cf.Bounds().Dy(), true)
	draw.Draw(__obf_a96f03e430996c71, __obf_0594f054c57944cf.Bounds(), __obf_0594f054c57944cf, image.Point{}, draw.Over)
	__obf_a96f03e430996c71.
		CropScaleCircle(__obf_0594f054c57944cf.Bounds().Dx()/2, __obf_0594f054c57944cf.Bounds().Dy()/2, __obf_0594f054c57944cf.Bounds().Dy()/2, __obf_fe92b35fa565dcef.ScaleRatioSize)
	__obf_a96f03e430996c71.
		Rotate(__obf_fe92b35fa565dcef.Angle, true)
	__obf_e1095470eb4fefe4 := __obf_a96f03e430996c71.Bounds()
	if __obf_e1095470eb4fefe4.Dy() > __obf_31d7cb2a7f8f7e87.Dy() || __obf_e1095470eb4fefe4.Dx() > __obf_31d7cb2a7f8f7e87.Dx() {
		__obf_84bc6da9d796ec0d := helper.CreateNRGBACanvas(__obf_0594f054c57944cf.Bounds().Dx(), __obf_0594f054c57944cf.Bounds().Dy(), true)
		draw.Draw(__obf_84bc6da9d796ec0d, __obf_84bc6da9d796ec0d.Bounds(), __obf_a96f03e430996c71, image.Point{X: (__obf_e1095470eb4fefe4.Dy() - __obf_31d7cb2a7f8f7e87.Dy()) / 2, Y: (__obf_e1095470eb4fefe4.Dx() - __obf_31d7cb2a7f8f7e87.Dx()) / 2}, draw.Over)
		__obf_a96f03e430996c71 = __obf_84bc6da9d796ec0d
	}

	return __obf_a96f03e430996c71, nil
}

func (__obf_3e760783dbf49c5a *Captcha) __obf_88d5fed3bc1bff4f(__obf_9680c355ab6a7d55 int) (image.Image, error) {
	__obf_c7ccd6117b2a26e2, __obf_b4c4cf14feeaf592 := assets.PickBgImage()
	if __obf_b4c4cf14feeaf592 != nil {
		return nil, __obf_b4c4cf14feeaf592
	}

	return __obf_3e760783dbf49c5a.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_9680c355ab6a7d55,
		Background: __obf_c7ccd6117b2a26e2,
		Alpha:      __obf_3e760783dbf49c5a.__obf_5ace98d4b75065f4.Primary.Alpha,
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
func (__obf_3e760783dbf49c5a *Captcha) __obf_d3ecd9ba7934c34b(__obf_0594f054c57944cf image.Image, __obf_f9c9a01651aaf15e int, __obf_c6d635faa67c7ae6 int) (image.Image, error) {
	return __obf_3e760783dbf49c5a.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_0594f054c57944cf,
		Alpha:          __obf_3e760783dbf49c5a.__obf_5ace98d4b75065f4.Secondary.Alpha,
		Size:           __obf_c6d635faa67c7ae6,
		Angle:          __obf_f9c9a01651aaf15e,
		ScaleRatioSize: (__obf_3e760783dbf49c5a.__obf_5ace98d4b75065f4.Primary.Size - __obf_c6d635faa67c7ae6) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_3e760783dbf49c5a *Captcha) __obf_0ceadbaf6b084be6() int {
	__obf_ad4e5eafeb1f7787 := __obf_3e760783dbf49c5a.__obf_5ace98d4b75065f4.Primary.AnglePosRange
	__obf_1f963b117c8e0e1c := rand.IntN(len(__obf_ad4e5eafeb1f7787))
	if __obf_1f963b117c8e0e1c < 0 {
		return 0
	}
	__obf_f9c9a01651aaf15e := __obf_ad4e5eafeb1f7787[__obf_1f963b117c8e0e1c]
	return util.RandInt(__obf_f9c9a01651aaf15e.Min, __obf_f9c9a01651aaf15e.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_3e760783dbf49c5a *Captcha) __obf_d4115b9deb1ede70() int {
	__obf_369078b6e4f4b120 := __obf_3e760783dbf49c5a.__obf_5ace98d4b75065f4.Secondary.SizeRange
	__obf_1f963b117c8e0e1c := rand.IntN(len(__obf_369078b6e4f4b120))
	if __obf_1f963b117c8e0e1c < 0 {
		return __obf_369078b6e4f4b120[0]
	}

	return __obf_369078b6e4f4b120[__obf_1f963b117c8e0e1c]
}
