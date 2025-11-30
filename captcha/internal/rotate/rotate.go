package __obf_19fe0d6ff470ddc2

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
	__obf_69fafa088d6b76c6 *types.RotateOption
}

var _ types.Captcha = (*Captcha)(nil)

func (__obf_0ce9a2192c36c384 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_173caf9a3a4c691b := __obf_0ce9a2192c36c384.__obf_c20277a03e2e0236()
	__obf_efaba3b6c7084e36, __obf_3f94c08dae590e1d := __obf_0ce9a2192c36c384.__obf_437ed618440850cc(__obf_0ce9a2192c36c384.__obf_69fafa088d6b76c6.Primary.Size)
	if __obf_3f94c08dae590e1d != nil {
		return nil, __obf_3f94c08dae590e1d
	}
	__obf_dbb4db6cc7ddc16b := __obf_0ce9a2192c36c384.__obf_36999f1318c69923()
	__obf_0475d73f67e757fd, __obf_3f94c08dae590e1d := __obf_0ce9a2192c36c384.__obf_ebda6d9145c62d9b(__obf_efaba3b6c7084e36, __obf_dbb4db6cc7ddc16b, __obf_173caf9a3a4c691b)
	if __obf_3f94c08dae590e1d != nil {
		return nil, __obf_3f94c08dae590e1d
	}

	return &CaptchaData{__obf_dbb4db6cc7ddc16b: __obf_dbb4db6cc7ddc16b, __obf_efaba3b6c7084e36: types.NewPNGImage(__obf_efaba3b6c7084e36), __obf_0475d73f67e757fd: types.NewPNGImage(__obf_0475d73f67e757fd)}, nil
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_69fafa088d6b76c6 *types.RotateOption) *Captcha {
	if __obf_69fafa088d6b76c6 == nil {
		return nil
	}
	return &Captcha{__obf_69fafa088d6b76c6: __obf_69fafa088d6b76c6}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_0ce9a2192c36c384 *Captcha) DrawWithNRGBA(__obf_3e3c366c4d7a1dd5 *DrawImageParams) (image.Image, error) {
	__obf_d450d646fa9c7743 := helper.CreateNRGBACanvas(__obf_3e3c366c4d7a1dd5.Size, __obf_3e3c366c4d7a1dd5.Size, true)
	if __obf_3e3c366c4d7a1dd5.Background != nil {
		__obf_b35998bec5459f6d := __obf_3e3c366c4d7a1dd5.Background
		__obf_7981020e70e0ad52 := __obf_b35998bec5459f6d.Bounds()
		__obf_53ebfa315e6792de := helper.CreateNRGBACanvas(__obf_7981020e70e0ad52.Dx(), __obf_7981020e70e0ad52.Dy(), true)
		__obf_578876a27c38f119 := util.RangCutImagePos(__obf_3e3c366c4d7a1dd5.Size, __obf_3e3c366c4d7a1dd5.Size, __obf_b35998bec5459f6d)
		draw.Draw(__obf_53ebfa315e6792de, __obf_7981020e70e0ad52, __obf_b35998bec5459f6d, __obf_578876a27c38f119, draw.Over)
		__obf_53ebfa315e6792de.
			SubImage(image.Rect(0, 0, __obf_3e3c366c4d7a1dd5.Size, __obf_3e3c366c4d7a1dd5.Size))
		draw.Draw(__obf_d450d646fa9c7743, __obf_d450d646fa9c7743.Bounds(), __obf_53ebfa315e6792de, image.Point{}, draw.Over)
	}
	__obf_d450d646fa9c7743.
		CropCircle(__obf_d450d646fa9c7743.Bounds().Dx()/2, __obf_d450d646fa9c7743.Bounds().Dy()/2, __obf_d450d646fa9c7743.Bounds().Dy()/2)
	return __obf_d450d646fa9c7743, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_0ce9a2192c36c384 *Captcha) DrawWithCropCircle(__obf_3e3c366c4d7a1dd5 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_b35998bec5459f6d := __obf_3e3c366c4d7a1dd5.Background
	__obf_738538f6affdb46d := __obf_b35998bec5459f6d.Bounds()
	__obf_e19ea62b89a50606 := helper.CreateNRGBACanvas(__obf_b35998bec5459f6d.Bounds().Dx(), __obf_b35998bec5459f6d.Bounds().Dy(), true)
	draw.Draw(__obf_e19ea62b89a50606, __obf_b35998bec5459f6d.Bounds(), __obf_b35998bec5459f6d, image.Point{}, draw.Over)
	__obf_e19ea62b89a50606.
		CropScaleCircle(__obf_b35998bec5459f6d.Bounds().Dx()/2, __obf_b35998bec5459f6d.Bounds().Dy()/2, __obf_b35998bec5459f6d.Bounds().Dy()/2, __obf_3e3c366c4d7a1dd5.ScaleRatioSize)
	__obf_e19ea62b89a50606.
		Rotate(__obf_3e3c366c4d7a1dd5.Angle, true)
	__obf_b0e86f5d10e3fc69 := __obf_e19ea62b89a50606.Bounds()
	if __obf_b0e86f5d10e3fc69.Dy() > __obf_738538f6affdb46d.Dy() || __obf_b0e86f5d10e3fc69.Dx() > __obf_738538f6affdb46d.Dx() {
		__obf_9fd8212b9d0f8772 := helper.CreateNRGBACanvas(__obf_b35998bec5459f6d.Bounds().Dx(), __obf_b35998bec5459f6d.Bounds().Dy(), true)
		draw.Draw(__obf_9fd8212b9d0f8772, __obf_9fd8212b9d0f8772.Bounds(), __obf_e19ea62b89a50606, image.Point{X: (__obf_b0e86f5d10e3fc69.Dy() - __obf_738538f6affdb46d.Dy()) / 2, Y: (__obf_b0e86f5d10e3fc69.Dx() - __obf_738538f6affdb46d.Dx()) / 2}, draw.Over)
		__obf_e19ea62b89a50606 = __obf_9fd8212b9d0f8772
	}

	return __obf_e19ea62b89a50606, nil
}

func (__obf_0ce9a2192c36c384 *Captcha) __obf_437ed618440850cc(__obf_0786667aab0a12d2 int) (image.Image, error) {
	__obf_edba1daf4e9d0e73, __obf_3f94c08dae590e1d := assets.PickBgImage()
	if __obf_3f94c08dae590e1d != nil {
		return nil, __obf_3f94c08dae590e1d
	}

	return __obf_0ce9a2192c36c384.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_0786667aab0a12d2,
		Background: __obf_edba1daf4e9d0e73,
		Alpha:      __obf_0ce9a2192c36c384.__obf_69fafa088d6b76c6.Primary.Alpha,
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
func (__obf_0ce9a2192c36c384 *Captcha) __obf_ebda6d9145c62d9b(__obf_b35998bec5459f6d image.Image, __obf_dbb4db6cc7ddc16b int, __obf_173caf9a3a4c691b int) (image.Image, error) {
	return __obf_0ce9a2192c36c384.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_b35998bec5459f6d,
		Alpha:          __obf_0ce9a2192c36c384.__obf_69fafa088d6b76c6.Secondary.Alpha,
		Size:           __obf_173caf9a3a4c691b,
		Angle:          __obf_dbb4db6cc7ddc16b,
		ScaleRatioSize: (__obf_0ce9a2192c36c384.__obf_69fafa088d6b76c6.Primary.Size - __obf_173caf9a3a4c691b) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_0ce9a2192c36c384 *Captcha) __obf_36999f1318c69923() int {
	__obf_d651893848ecc4fb := __obf_0ce9a2192c36c384.__obf_69fafa088d6b76c6.Primary.AnglePosRange
	__obf_afb5af04f27c7c95 := rand.IntN(len(__obf_d651893848ecc4fb))
	if __obf_afb5af04f27c7c95 < 0 {
		return 0
	}
	__obf_dbb4db6cc7ddc16b := __obf_d651893848ecc4fb[__obf_afb5af04f27c7c95]
	return util.RandInt(__obf_dbb4db6cc7ddc16b.Min, __obf_dbb4db6cc7ddc16b.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_0ce9a2192c36c384 *Captcha) __obf_c20277a03e2e0236() int {
	__obf_ad1dd59230a8fcb5 := __obf_0ce9a2192c36c384.__obf_69fafa088d6b76c6.Secondary.SizeRange
	__obf_afb5af04f27c7c95 := rand.IntN(len(__obf_ad1dd59230a8fcb5))
	if __obf_afb5af04f27c7c95 < 0 {
		return __obf_ad1dd59230a8fcb5[0]
	}

	return __obf_ad1dd59230a8fcb5[__obf_afb5af04f27c7c95]
}
