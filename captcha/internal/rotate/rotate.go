package __obf_220a18b254a8ad8c

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
	__obf_69d320ed4fef6931 *Options
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_69d320ed4fef6931 *Options) *Captcha {
	if __obf_69d320ed4fef6931 == nil {
		return nil
	}
	return &Captcha{
		__obf_69d320ed4fef6931: __obf_69d320ed4fef6931,
	}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_232060d34b1a5883 *Captcha) DrawWithNRGBA(__obf_763012e71a7d0a31 *DrawImageParams) (image.Image, error) {
	__obf_9463655ef915ce8b := helper.CreateNRGBACanvas(__obf_763012e71a7d0a31.Size, __obf_763012e71a7d0a31.Size, true)
	if __obf_763012e71a7d0a31.Background != nil {
		__obf_7f615559558cf3f7 := __obf_763012e71a7d0a31.Background
		__obf_6e02fa725d419b97 := __obf_7f615559558cf3f7.Bounds()
		__obf_39a0194b1128ffcc := helper.CreateNRGBACanvas(__obf_6e02fa725d419b97.Dx(), __obf_6e02fa725d419b97.Dy(), true)
		__obf_f3992358334023b8 := util.RangCutImagePos(__obf_763012e71a7d0a31.Size, __obf_763012e71a7d0a31.Size, __obf_7f615559558cf3f7)
		draw.Draw(__obf_39a0194b1128ffcc, __obf_6e02fa725d419b97, __obf_7f615559558cf3f7, __obf_f3992358334023b8, draw.Over)
		__obf_39a0194b1128ffcc.SubImage(image.Rect(0, 0, __obf_763012e71a7d0a31.Size, __obf_763012e71a7d0a31.Size))
		draw.Draw(__obf_9463655ef915ce8b, __obf_9463655ef915ce8b.Bounds(), __obf_39a0194b1128ffcc, image.Point{}, draw.Over)
	}

	__obf_9463655ef915ce8b.CropCircle(__obf_9463655ef915ce8b.Bounds().Dx()/2, __obf_9463655ef915ce8b.Bounds().Dy()/2, __obf_9463655ef915ce8b.Bounds().Dy()/2)
	return __obf_9463655ef915ce8b, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_232060d34b1a5883 *Captcha) DrawWithCropCircle(__obf_763012e71a7d0a31 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_7f615559558cf3f7 := __obf_763012e71a7d0a31.Background

	__obf_89388062de99aa26 := __obf_7f615559558cf3f7.Bounds()
	__obf_d1a4fbce9d5a4dc7 := helper.CreateNRGBACanvas(__obf_7f615559558cf3f7.Bounds().Dx(), __obf_7f615559558cf3f7.Bounds().Dy(), true)
	draw.Draw(__obf_d1a4fbce9d5a4dc7, __obf_7f615559558cf3f7.Bounds(), __obf_7f615559558cf3f7, image.Point{}, draw.Over)
	__obf_d1a4fbce9d5a4dc7.CropScaleCircle(__obf_7f615559558cf3f7.Bounds().Dx()/2, __obf_7f615559558cf3f7.Bounds().Dy()/2, __obf_7f615559558cf3f7.Bounds().Dy()/2, __obf_763012e71a7d0a31.ScaleRatioSize)
	__obf_d1a4fbce9d5a4dc7.Rotate(__obf_763012e71a7d0a31.Angle, true)

	__obf_999e33ed618209d4 := __obf_d1a4fbce9d5a4dc7.Bounds()
	if __obf_999e33ed618209d4.Dy() > __obf_89388062de99aa26.Dy() || __obf_999e33ed618209d4.Dx() > __obf_89388062de99aa26.Dx() {
		__obf_3b8afdaa8f4960c9 := helper.CreateNRGBACanvas(__obf_7f615559558cf3f7.Bounds().Dx(), __obf_7f615559558cf3f7.Bounds().Dy(), true)
		draw.Draw(__obf_3b8afdaa8f4960c9, __obf_3b8afdaa8f4960c9.Bounds(), __obf_d1a4fbce9d5a4dc7, image.Point{X: (__obf_999e33ed618209d4.Dy() - __obf_89388062de99aa26.Dy()) / 2, Y: (__obf_999e33ed618209d4.Dx() - __obf_89388062de99aa26.Dx()) / 2}, draw.Over)
		__obf_d1a4fbce9d5a4dc7 = __obf_3b8afdaa8f4960c9
	}

	return __obf_d1a4fbce9d5a4dc7, nil
}

// Generate generates rotate CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_232060d34b1a5883 *Captcha) Generate() (*CaptchaData, error) {
	__obf_6e050a12e37b7f26 := __obf_232060d34b1a5883.__obf_59442a8f938e86ef()

	__obf_5a73f4d6e809a9e0, __obf_c6d3abe7331184ce := __obf_232060d34b1a5883.__obf_5c56008ea36ca54a(__obf_232060d34b1a5883.__obf_69d320ed4fef6931.Primary.Size)
	if __obf_c6d3abe7331184ce != nil {
		return nil, __obf_c6d3abe7331184ce
	}
	__obf_1a66d3fac7d22235 := __obf_232060d34b1a5883.__obf_a2fa37f4b327dfa4()
	__obf_a37b058809dc2b58, __obf_c6d3abe7331184ce := __obf_232060d34b1a5883.__obf_aefc9cda9ad43446(__obf_5a73f4d6e809a9e0, __obf_1a66d3fac7d22235, __obf_6e050a12e37b7f26)
	if __obf_c6d3abe7331184ce != nil {
		return nil, __obf_c6d3abe7331184ce
	}

	return &CaptchaData{
		__obf_1a66d3fac7d22235: __obf_1a66d3fac7d22235,
		__obf_5a73f4d6e809a9e0: types.NewPNGImage(__obf_5a73f4d6e809a9e0),
		__obf_a37b058809dc2b58: types.NewPNGImage(__obf_a37b058809dc2b58),
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
func (__obf_232060d34b1a5883 *Captcha) __obf_5c56008ea36ca54a(__obf_3e1c3a430ef2f994 int) (image.Image, error) {

	__obf_00086e3f5271db98, __obf_c6d3abe7331184ce := assets.PickBgImage()
	if __obf_c6d3abe7331184ce != nil {
		return nil, __obf_c6d3abe7331184ce
	}

	return __obf_232060d34b1a5883.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_3e1c3a430ef2f994,
		Background: __obf_00086e3f5271db98,
		Alpha:      __obf_232060d34b1a5883.__obf_69d320ed4fef6931.Primary.Alpha,
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
func (__obf_232060d34b1a5883 *Captcha) __obf_aefc9cda9ad43446(__obf_7f615559558cf3f7 image.Image, __obf_1a66d3fac7d22235 int, __obf_6e050a12e37b7f26 int) (image.Image, error) {
	return __obf_232060d34b1a5883.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_7f615559558cf3f7,
		Alpha:          __obf_232060d34b1a5883.__obf_69d320ed4fef6931.Secondary.Alpha,
		Size:           __obf_6e050a12e37b7f26,
		Angle:          __obf_1a66d3fac7d22235,
		ScaleRatioSize: (__obf_232060d34b1a5883.__obf_69d320ed4fef6931.Primary.Size - __obf_6e050a12e37b7f26) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_232060d34b1a5883 *Captcha) __obf_a2fa37f4b327dfa4() int {
	__obf_ba5883aaee7f9634 := __obf_232060d34b1a5883.__obf_69d320ed4fef6931.Primary.AnglePosRange

	__obf_af3524a9d6b3286e := rand.IntN(len(__obf_ba5883aaee7f9634))
	if __obf_af3524a9d6b3286e < 0 {
		return 0
	}

	__obf_1a66d3fac7d22235 := __obf_ba5883aaee7f9634[__obf_af3524a9d6b3286e]
	return util.RandInt(__obf_1a66d3fac7d22235.Min, __obf_1a66d3fac7d22235.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_232060d34b1a5883 *Captcha) __obf_59442a8f938e86ef() int {
	__obf_0536bc68dfda7073 := __obf_232060d34b1a5883.__obf_69d320ed4fef6931.Secondary.SizeRange

	__obf_af3524a9d6b3286e := rand.IntN(len(__obf_0536bc68dfda7073))
	if __obf_af3524a9d6b3286e < 0 {
		return __obf_0536bc68dfda7073[0]
	}

	return __obf_0536bc68dfda7073[__obf_af3524a9d6b3286e]
}
