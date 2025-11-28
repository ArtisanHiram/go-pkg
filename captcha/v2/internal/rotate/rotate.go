package __obf_f48dc973afdacd79

import (
	"image"
	"math/rand/v2"

	assets "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/assets"
	helper "github.com/ArtisanHiram/go-pkg/captcha/v2/internal/helper"
	types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"
	util "github.com/ArtisanHiram/go-pkg/captcha/v2/util"
	"golang.org/x/image/draw"
)

// Captcha is the concrete implementation of the Captcha interface
type Captcha struct {
	__obf_7b75bc0b46b45bda *Options
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_7b75bc0b46b45bda *Options) *Captcha {
	if __obf_7b75bc0b46b45bda == nil {
		return nil
	}
	return &Captcha{
		__obf_7b75bc0b46b45bda: __obf_7b75bc0b46b45bda,
	}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_36a78059fd35abb2 *Captcha) DrawWithNRGBA(__obf_f6aa2935a958f523 *DrawImageParams) (image.Image, error) {
	__obf_212559b1d2632067 := helper.CreateNRGBACanvas(__obf_f6aa2935a958f523.Size, __obf_f6aa2935a958f523.Size, true)
	if __obf_f6aa2935a958f523.Background != nil {
		__obf_d5323e8fcd35f8a0 := __obf_f6aa2935a958f523.Background
		__obf_a48d8f82e77e3d09 := __obf_d5323e8fcd35f8a0.Bounds()
		__obf_13eab3a32e517075 := helper.CreateNRGBACanvas(__obf_a48d8f82e77e3d09.Dx(), __obf_a48d8f82e77e3d09.Dy(), true)
		__obf_6b1b5829051f869e := util.RangCutImagePos(__obf_f6aa2935a958f523.Size, __obf_f6aa2935a958f523.Size, __obf_d5323e8fcd35f8a0)
		draw.Draw(__obf_13eab3a32e517075, __obf_a48d8f82e77e3d09, __obf_d5323e8fcd35f8a0, __obf_6b1b5829051f869e, draw.Over)
		__obf_13eab3a32e517075.SubImage(image.Rect(0, 0, __obf_f6aa2935a958f523.Size, __obf_f6aa2935a958f523.Size))
		draw.Draw(__obf_212559b1d2632067, __obf_212559b1d2632067.Bounds(), __obf_13eab3a32e517075, image.Point{}, draw.Over)
	}

	__obf_212559b1d2632067.CropCircle(__obf_212559b1d2632067.Bounds().Dx()/2, __obf_212559b1d2632067.Bounds().Dy()/2, __obf_212559b1d2632067.Bounds().Dy()/2)
	return __obf_212559b1d2632067, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_36a78059fd35abb2 *Captcha) DrawWithCropCircle(__obf_f6aa2935a958f523 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_d5323e8fcd35f8a0 := __obf_f6aa2935a958f523.Background

	__obf_6276493423a49e2a := __obf_d5323e8fcd35f8a0.Bounds()
	__obf_2e18e50356f484a2 := helper.CreateNRGBACanvas(__obf_d5323e8fcd35f8a0.Bounds().Dx(), __obf_d5323e8fcd35f8a0.Bounds().Dy(), true)
	draw.Draw(__obf_2e18e50356f484a2, __obf_d5323e8fcd35f8a0.Bounds(), __obf_d5323e8fcd35f8a0, image.Point{}, draw.Over)
	__obf_2e18e50356f484a2.CropScaleCircle(__obf_d5323e8fcd35f8a0.Bounds().Dx()/2, __obf_d5323e8fcd35f8a0.Bounds().Dy()/2, __obf_d5323e8fcd35f8a0.Bounds().Dy()/2, __obf_f6aa2935a958f523.ScaleRatioSize)
	__obf_2e18e50356f484a2.Rotate(__obf_f6aa2935a958f523.Angle, true)

	__obf_45cecd33fa7abc45 := __obf_2e18e50356f484a2.Bounds()
	if __obf_45cecd33fa7abc45.Dy() > __obf_6276493423a49e2a.Dy() || __obf_45cecd33fa7abc45.Dx() > __obf_6276493423a49e2a.Dx() {
		__obf_1e7b08c98edc0f76 := helper.CreateNRGBACanvas(__obf_d5323e8fcd35f8a0.Bounds().Dx(), __obf_d5323e8fcd35f8a0.Bounds().Dy(), true)
		draw.Draw(__obf_1e7b08c98edc0f76, __obf_1e7b08c98edc0f76.Bounds(), __obf_2e18e50356f484a2, image.Point{X: (__obf_45cecd33fa7abc45.Dy() - __obf_6276493423a49e2a.Dy()) / 2, Y: (__obf_45cecd33fa7abc45.Dx() - __obf_6276493423a49e2a.Dx()) / 2}, draw.Over)
		__obf_2e18e50356f484a2 = __obf_1e7b08c98edc0f76
	}

	return __obf_2e18e50356f484a2, nil
}

// Generate generates rotate CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_36a78059fd35abb2 *Captcha) Generate() (*CaptchaData, error) {
	__obf_15da80a8f53facaf := __obf_36a78059fd35abb2.__obf_545b9037d6f907f3()

	__obf_3482409dfa579189, __obf_a9a016df08796ca8 := __obf_36a78059fd35abb2.__obf_f46a13809279b44f(__obf_36a78059fd35abb2.__obf_7b75bc0b46b45bda.Primary.Size)
	if __obf_a9a016df08796ca8 != nil {
		return nil, __obf_a9a016df08796ca8
	}
	__obf_bbdf82463e548025 := __obf_36a78059fd35abb2.__obf_d815f3aee6b94b28()
	__obf_2ea6abf946c6d439, __obf_a9a016df08796ca8 := __obf_36a78059fd35abb2.__obf_7fe3fbf5657bd3e5(__obf_3482409dfa579189, __obf_bbdf82463e548025, __obf_15da80a8f53facaf)
	if __obf_a9a016df08796ca8 != nil {
		return nil, __obf_a9a016df08796ca8
	}

	return &CaptchaData{
		__obf_bbdf82463e548025: __obf_bbdf82463e548025,
		__obf_3482409dfa579189: types.NewPNGImage(__obf_3482409dfa579189),
		__obf_2ea6abf946c6d439: types.NewPNGImage(__obf_2ea6abf946c6d439),
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
func (__obf_36a78059fd35abb2 *Captcha) __obf_f46a13809279b44f(__obf_f97667d40f445301 int) (image.Image, error) {

	__obf_786e659f1622d9cf, __obf_a9a016df08796ca8 := assets.PickBgImage()
	if __obf_a9a016df08796ca8 != nil {
		return nil, __obf_a9a016df08796ca8
	}

	return __obf_36a78059fd35abb2.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_f97667d40f445301,
		Background: __obf_786e659f1622d9cf,
		Alpha:      __obf_36a78059fd35abb2.__obf_7b75bc0b46b45bda.Primary.Alpha,
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
func (__obf_36a78059fd35abb2 *Captcha) __obf_7fe3fbf5657bd3e5(__obf_d5323e8fcd35f8a0 image.Image, __obf_bbdf82463e548025 int, __obf_15da80a8f53facaf int) (image.Image, error) {
	return __obf_36a78059fd35abb2.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_d5323e8fcd35f8a0,
		Alpha:          __obf_36a78059fd35abb2.__obf_7b75bc0b46b45bda.Secondary.Alpha,
		Size:           __obf_15da80a8f53facaf,
		Angle:          __obf_bbdf82463e548025,
		ScaleRatioSize: (__obf_36a78059fd35abb2.__obf_7b75bc0b46b45bda.Primary.Size - __obf_15da80a8f53facaf) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_36a78059fd35abb2 *Captcha) __obf_d815f3aee6b94b28() int {
	__obf_ce9ad534417bb5b0 := __obf_36a78059fd35abb2.__obf_7b75bc0b46b45bda.Primary.AnglePosRange

	__obf_64fc1b369aa271d2 := rand.IntN(len(__obf_ce9ad534417bb5b0))
	if __obf_64fc1b369aa271d2 < 0 {
		return 0
	}

	__obf_bbdf82463e548025 := __obf_ce9ad534417bb5b0[__obf_64fc1b369aa271d2]
	return util.RandInt(__obf_bbdf82463e548025.Min, __obf_bbdf82463e548025.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_36a78059fd35abb2 *Captcha) __obf_545b9037d6f907f3() int {
	__obf_37f544db16db9b4b := __obf_36a78059fd35abb2.__obf_7b75bc0b46b45bda.Secondary.SizeRange

	__obf_64fc1b369aa271d2 := rand.IntN(len(__obf_37f544db16db9b4b))
	if __obf_64fc1b369aa271d2 < 0 {
		return __obf_37f544db16db9b4b[0]
	}

	return __obf_37f544db16db9b4b[__obf_64fc1b369aa271d2]
}
