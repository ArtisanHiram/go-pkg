package __obf_4b873d45958741b9

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
	__obf_333de8b3f9d27a1d *types.RotateOption
}

var _ types.Captcha = (*Captcha)(nil)

func (__obf_4f5780d7172901a0 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_c37258d7fa06b5bb := __obf_4f5780d7172901a0.__obf_30adf6650a0c698c()
	__obf_ee7efadff09b92f2, __obf_8fcb878c1da6d243 := __obf_4f5780d7172901a0.__obf_4f2b4340af99c402(__obf_4f5780d7172901a0.__obf_333de8b3f9d27a1d.Primary.Size)
	if __obf_8fcb878c1da6d243 != nil {
		return nil, __obf_8fcb878c1da6d243
	}
	__obf_fde4510d2c9d2319 := __obf_4f5780d7172901a0.__obf_9924a55ead5306bd()
	__obf_c9f1e5ba3e028129, __obf_8fcb878c1da6d243 := __obf_4f5780d7172901a0.__obf_2a465d70a75e68e5(__obf_ee7efadff09b92f2, __obf_fde4510d2c9d2319, __obf_c37258d7fa06b5bb)
	if __obf_8fcb878c1da6d243 != nil {
		return nil, __obf_8fcb878c1da6d243
	}

	return &CaptchaData{
		Angle: __obf_fde4510d2c9d2319, __obf_ee7efadff09b92f2: types.NewPNGImage(__obf_ee7efadff09b92f2), __obf_c9f1e5ba3e028129: types.NewPNGImage(__obf_c9f1e5ba3e028129),
	}, nil
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_333de8b3f9d27a1d *types.RotateOption) *Captcha {
	if __obf_333de8b3f9d27a1d == nil {
		return nil
	}
	return &Captcha{__obf_333de8b3f9d27a1d: __obf_333de8b3f9d27a1d}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_4f5780d7172901a0 *Captcha) DrawWithNRGBA(__obf_21a687a82dd51c49 *DrawImageParams) (image.Image, error) {
	__obf_eb8abc51732738af := helper.CreateNRGBACanvas(__obf_21a687a82dd51c49.Size, __obf_21a687a82dd51c49.Size, true)
	if __obf_21a687a82dd51c49.Background != nil {
		__obf_b76781d1a11e0e5b := __obf_21a687a82dd51c49.Background
		__obf_e87e8f526074cc8d := __obf_b76781d1a11e0e5b.Bounds()
		__obf_3b5cec9ed4597776 := helper.CreateNRGBACanvas(__obf_e87e8f526074cc8d.Dx(), __obf_e87e8f526074cc8d.Dy(), true)
		__obf_e7d5608d51f7f31f := util.RangCutImagePos(__obf_21a687a82dd51c49.Size, __obf_21a687a82dd51c49.Size, __obf_b76781d1a11e0e5b)
		draw.Draw(__obf_3b5cec9ed4597776, __obf_e87e8f526074cc8d, __obf_b76781d1a11e0e5b, __obf_e7d5608d51f7f31f, draw.Over)
		__obf_3b5cec9ed4597776.
			SubImage(image.Rect(0, 0, __obf_21a687a82dd51c49.Size, __obf_21a687a82dd51c49.Size))
		draw.Draw(__obf_eb8abc51732738af, __obf_eb8abc51732738af.Bounds(), __obf_3b5cec9ed4597776, image.Point{}, draw.Over)
	}
	__obf_eb8abc51732738af.
		CropCircle(__obf_eb8abc51732738af.Bounds().Dx()/2, __obf_eb8abc51732738af.Bounds().Dy()/2, __obf_eb8abc51732738af.Bounds().Dy()/2)
	return __obf_eb8abc51732738af, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_4f5780d7172901a0 *Captcha) DrawWithCropCircle(__obf_21a687a82dd51c49 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_b76781d1a11e0e5b := __obf_21a687a82dd51c49.Background
	__obf_7bbd6659bfdd9779 := __obf_b76781d1a11e0e5b.Bounds()
	__obf_a27d337cabf64be4 := helper.CreateNRGBACanvas(__obf_b76781d1a11e0e5b.Bounds().Dx(), __obf_b76781d1a11e0e5b.Bounds().Dy(), true)
	draw.Draw(__obf_a27d337cabf64be4, __obf_b76781d1a11e0e5b.Bounds(), __obf_b76781d1a11e0e5b, image.Point{}, draw.Over)
	__obf_a27d337cabf64be4.
		CropScaleCircle(__obf_b76781d1a11e0e5b.Bounds().Dx()/2, __obf_b76781d1a11e0e5b.Bounds().Dy()/2, __obf_b76781d1a11e0e5b.Bounds().Dy()/2, __obf_21a687a82dd51c49.ScaleRatioSize)
	__obf_a27d337cabf64be4.
		Rotate(__obf_21a687a82dd51c49.Angle, true)
	__obf_618977667df267ca := __obf_a27d337cabf64be4.Bounds()
	if __obf_618977667df267ca.Dy() > __obf_7bbd6659bfdd9779.Dy() || __obf_618977667df267ca.Dx() > __obf_7bbd6659bfdd9779.Dx() {
		__obf_79ed4b361066f253 := helper.CreateNRGBACanvas(__obf_b76781d1a11e0e5b.Bounds().Dx(), __obf_b76781d1a11e0e5b.Bounds().Dy(), true)
		draw.Draw(__obf_79ed4b361066f253, __obf_79ed4b361066f253.Bounds(), __obf_a27d337cabf64be4, image.Point{X: (__obf_618977667df267ca.Dy() - __obf_7bbd6659bfdd9779.Dy()) / 2, Y: (__obf_618977667df267ca.Dx() - __obf_7bbd6659bfdd9779.Dx()) / 2}, draw.Over)
		__obf_a27d337cabf64be4 = __obf_79ed4b361066f253
	}

	return __obf_a27d337cabf64be4, nil
}

func (__obf_4f5780d7172901a0 *Captcha) __obf_4f2b4340af99c402(__obf_8c75943f0e04d4e7 int) (image.Image, error) {
	__obf_0433c2fa8058d317, __obf_8fcb878c1da6d243 := assets.PickBgImage()
	if __obf_8fcb878c1da6d243 != nil {
		return nil, __obf_8fcb878c1da6d243
	}

	return __obf_4f5780d7172901a0.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_8c75943f0e04d4e7,
		Background: __obf_0433c2fa8058d317,
		Alpha:      __obf_4f5780d7172901a0.__obf_333de8b3f9d27a1d.Primary.Alpha,
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
func (__obf_4f5780d7172901a0 *Captcha) __obf_2a465d70a75e68e5(__obf_b76781d1a11e0e5b image.Image, __obf_fde4510d2c9d2319 int, __obf_c37258d7fa06b5bb int) (image.Image, error) {
	return __obf_4f5780d7172901a0.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_b76781d1a11e0e5b,
		Alpha:          __obf_4f5780d7172901a0.__obf_333de8b3f9d27a1d.Secondary.Alpha,
		Size:           __obf_c37258d7fa06b5bb,
		Angle:          __obf_fde4510d2c9d2319,
		ScaleRatioSize: (__obf_4f5780d7172901a0.__obf_333de8b3f9d27a1d.Primary.Size - __obf_c37258d7fa06b5bb) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_4f5780d7172901a0 *Captcha) __obf_9924a55ead5306bd() int {
	__obf_2e67025aca90bf9f := __obf_4f5780d7172901a0.__obf_333de8b3f9d27a1d.Primary.AnglePosRange
	__obf_a10d7ffbeb286e4c := rand.IntN(len(__obf_2e67025aca90bf9f))
	if __obf_a10d7ffbeb286e4c < 0 {
		return 0
	}
	__obf_fde4510d2c9d2319 := __obf_2e67025aca90bf9f[__obf_a10d7ffbeb286e4c]
	return util.RandInt(__obf_fde4510d2c9d2319.Min, __obf_fde4510d2c9d2319.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_4f5780d7172901a0 *Captcha) __obf_30adf6650a0c698c() int {
	__obf_72bf943bdba25131 := __obf_4f5780d7172901a0.__obf_333de8b3f9d27a1d.Secondary.SizeRange
	__obf_a10d7ffbeb286e4c := rand.IntN(len(__obf_72bf943bdba25131))
	if __obf_a10d7ffbeb286e4c < 0 {
		return __obf_72bf943bdba25131[0]
	}

	return __obf_72bf943bdba25131[__obf_a10d7ffbeb286e4c]
}
