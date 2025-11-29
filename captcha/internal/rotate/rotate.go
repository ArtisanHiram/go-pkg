package __obf_3a7793861edae94b

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
	__obf_d35d2f1a410377c3 *Options
}

var _ types.Captcha = (*Captcha)(nil)

func (__obf_801f980d9ae99a74 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_41e3caff498f445c := __obf_801f980d9ae99a74.__obf_ddcdf91a572590f5()
	__obf_860f866a1a7270fa, __obf_d71636ed60e91d3f := __obf_801f980d9ae99a74.__obf_6bbe5b5e248647e3(__obf_801f980d9ae99a74.__obf_d35d2f1a410377c3.Primary.Size)
	if __obf_d71636ed60e91d3f != nil {
		return nil, __obf_d71636ed60e91d3f
	}
	__obf_50419e50fd91db67 := __obf_801f980d9ae99a74.__obf_a6e822c77f9924bd()
	__obf_ee543d5bbdc5b150, __obf_d71636ed60e91d3f := __obf_801f980d9ae99a74.__obf_f505850048b36206(__obf_860f866a1a7270fa, __obf_50419e50fd91db67, __obf_41e3caff498f445c)
	if __obf_d71636ed60e91d3f != nil {
		return nil, __obf_d71636ed60e91d3f
	}

	return &CaptchaData{__obf_50419e50fd91db67: __obf_50419e50fd91db67, __obf_860f866a1a7270fa: types.NewPNGImage(__obf_860f866a1a7270fa), __obf_ee543d5bbdc5b150: types.NewPNGImage(__obf_ee543d5bbdc5b150)}, nil
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_d35d2f1a410377c3 *Options) *Captcha {
	if __obf_d35d2f1a410377c3 == nil {
		return nil
	}
	return &Captcha{__obf_d35d2f1a410377c3: __obf_d35d2f1a410377c3}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_801f980d9ae99a74 *Captcha) DrawWithNRGBA(__obf_d1a387086403bc39 *DrawImageParams) (image.Image, error) {
	__obf_168a5ff438448bcb := helper.CreateNRGBACanvas(__obf_d1a387086403bc39.Size, __obf_d1a387086403bc39.Size, true)
	if __obf_d1a387086403bc39.Background != nil {
		__obf_510969c9880bfd12 := __obf_d1a387086403bc39.Background
		__obf_512f4b0a07d9710b := __obf_510969c9880bfd12.Bounds()
		__obf_9e60d122e4c157a5 := helper.CreateNRGBACanvas(__obf_512f4b0a07d9710b.Dx(), __obf_512f4b0a07d9710b.Dy(), true)
		__obf_f0acdc5539ba2fa8 := util.RangCutImagePos(__obf_d1a387086403bc39.Size, __obf_d1a387086403bc39.Size, __obf_510969c9880bfd12)
		draw.Draw(__obf_9e60d122e4c157a5, __obf_512f4b0a07d9710b, __obf_510969c9880bfd12, __obf_f0acdc5539ba2fa8, draw.Over)
		__obf_9e60d122e4c157a5.
			SubImage(image.Rect(0, 0, __obf_d1a387086403bc39.Size, __obf_d1a387086403bc39.Size))
		draw.Draw(__obf_168a5ff438448bcb, __obf_168a5ff438448bcb.Bounds(), __obf_9e60d122e4c157a5, image.Point{}, draw.Over)
	}
	__obf_168a5ff438448bcb.
		CropCircle(__obf_168a5ff438448bcb.Bounds().Dx()/2, __obf_168a5ff438448bcb.Bounds().Dy()/2, __obf_168a5ff438448bcb.Bounds().Dy()/2)
	return __obf_168a5ff438448bcb, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_801f980d9ae99a74 *Captcha) DrawWithCropCircle(__obf_d1a387086403bc39 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_510969c9880bfd12 := __obf_d1a387086403bc39.Background
	__obf_4f0ce77e9ef3a064 := __obf_510969c9880bfd12.Bounds()
	__obf_ade1412963ae9aa9 := helper.CreateNRGBACanvas(__obf_510969c9880bfd12.Bounds().Dx(), __obf_510969c9880bfd12.Bounds().Dy(), true)
	draw.Draw(__obf_ade1412963ae9aa9, __obf_510969c9880bfd12.Bounds(), __obf_510969c9880bfd12, image.Point{}, draw.Over)
	__obf_ade1412963ae9aa9.
		CropScaleCircle(__obf_510969c9880bfd12.Bounds().Dx()/2, __obf_510969c9880bfd12.Bounds().Dy()/2, __obf_510969c9880bfd12.Bounds().Dy()/2, __obf_d1a387086403bc39.ScaleRatioSize)
	__obf_ade1412963ae9aa9.
		Rotate(__obf_d1a387086403bc39.Angle, true)
	__obf_9d100fcdc52304e9 := __obf_ade1412963ae9aa9.Bounds()
	if __obf_9d100fcdc52304e9.Dy() > __obf_4f0ce77e9ef3a064.Dy() || __obf_9d100fcdc52304e9.Dx() > __obf_4f0ce77e9ef3a064.Dx() {
		__obf_cd6614f739b15f99 := helper.CreateNRGBACanvas(__obf_510969c9880bfd12.Bounds().Dx(), __obf_510969c9880bfd12.Bounds().Dy(), true)
		draw.Draw(__obf_cd6614f739b15f99, __obf_cd6614f739b15f99.Bounds(), __obf_ade1412963ae9aa9, image.Point{X: (__obf_9d100fcdc52304e9.Dy() - __obf_4f0ce77e9ef3a064.Dy()) / 2, Y: (__obf_9d100fcdc52304e9.Dx() - __obf_4f0ce77e9ef3a064.Dx()) / 2}, draw.Over)
		__obf_ade1412963ae9aa9 = __obf_cd6614f739b15f99
	}

	return __obf_ade1412963ae9aa9, nil
}

func (__obf_801f980d9ae99a74 *Captcha) __obf_6bbe5b5e248647e3(__obf_9f86da20d0967676 int) (image.Image, error) {
	__obf_0b4c2439edcbedeb, __obf_d71636ed60e91d3f := assets.PickBgImage()
	if __obf_d71636ed60e91d3f != nil {
		return nil, __obf_d71636ed60e91d3f
	}

	return __obf_801f980d9ae99a74.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_9f86da20d0967676,
		Background: __obf_0b4c2439edcbedeb,
		Alpha:      __obf_801f980d9ae99a74.__obf_d35d2f1a410377c3.Primary.Alpha,
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
func (__obf_801f980d9ae99a74 *Captcha) __obf_f505850048b36206(__obf_510969c9880bfd12 image.Image, __obf_50419e50fd91db67 int, __obf_41e3caff498f445c int) (image.Image, error) {
	return __obf_801f980d9ae99a74.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_510969c9880bfd12,
		Alpha:          __obf_801f980d9ae99a74.__obf_d35d2f1a410377c3.Secondary.Alpha,
		Size:           __obf_41e3caff498f445c,
		Angle:          __obf_50419e50fd91db67,
		ScaleRatioSize: (__obf_801f980d9ae99a74.__obf_d35d2f1a410377c3.Primary.Size - __obf_41e3caff498f445c) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_801f980d9ae99a74 *Captcha) __obf_a6e822c77f9924bd() int {
	__obf_cc7571a637886fe6 := __obf_801f980d9ae99a74.__obf_d35d2f1a410377c3.Primary.AnglePosRange
	__obf_7afac9f011ec4bf0 := rand.IntN(len(__obf_cc7571a637886fe6))
	if __obf_7afac9f011ec4bf0 < 0 {
		return 0
	}
	__obf_50419e50fd91db67 := __obf_cc7571a637886fe6[__obf_7afac9f011ec4bf0]
	return util.RandInt(__obf_50419e50fd91db67.Min, __obf_50419e50fd91db67.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_801f980d9ae99a74 *Captcha) __obf_ddcdf91a572590f5() int {
	__obf_afee98069b387f17 := __obf_801f980d9ae99a74.__obf_d35d2f1a410377c3.Secondary.SizeRange
	__obf_7afac9f011ec4bf0 := rand.IntN(len(__obf_afee98069b387f17))
	if __obf_7afac9f011ec4bf0 < 0 {
		return __obf_afee98069b387f17[0]
	}

	return __obf_afee98069b387f17[__obf_7afac9f011ec4bf0]
}
