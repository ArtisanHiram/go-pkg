package __obf_4e766de6f7fc6549

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
	__obf_5c1359684eb79b70 *types.RotateOption
}

var _ types.Captcha = (*Captcha)(nil)

func (__obf_598561cd40842fdf *Captcha) Generate() (types.CaptchaData, error) {
	__obf_5dfaeb3c1bd2a70a := __obf_598561cd40842fdf.__obf_858466f65ad88c3e()
	__obf_38820526d0b312f0, __obf_8e9665c84006346f := __obf_598561cd40842fdf.__obf_71ef3916f37c74fa(__obf_598561cd40842fdf.__obf_5c1359684eb79b70.Primary.Size)
	if __obf_8e9665c84006346f != nil {
		return nil, __obf_8e9665c84006346f
	}
	__obf_720ba9ea19239003 := __obf_598561cd40842fdf.__obf_c3950e0945f974a4()
	__obf_1e0e4b1d4069e3bb, __obf_8e9665c84006346f := __obf_598561cd40842fdf.__obf_a56199890c419472(__obf_38820526d0b312f0, __obf_720ba9ea19239003, __obf_5dfaeb3c1bd2a70a)
	if __obf_8e9665c84006346f != nil {
		return nil, __obf_8e9665c84006346f
	}

	return &CaptchaData{__obf_720ba9ea19239003: __obf_720ba9ea19239003, __obf_38820526d0b312f0: types.NewPNGImage(__obf_38820526d0b312f0), __obf_1e0e4b1d4069e3bb: types.NewPNGImage(__obf_1e0e4b1d4069e3bb)}, nil
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_5c1359684eb79b70 *types.RotateOption) *Captcha {
	if __obf_5c1359684eb79b70 == nil {
		return nil
	}
	return &Captcha{__obf_5c1359684eb79b70: __obf_5c1359684eb79b70}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_598561cd40842fdf *Captcha) DrawWithNRGBA(__obf_8491b090f9e1121f *DrawImageParams) (image.Image, error) {
	__obf_2d0ece5bcf1fe46a := helper.CreateNRGBACanvas(__obf_8491b090f9e1121f.Size, __obf_8491b090f9e1121f.Size, true)
	if __obf_8491b090f9e1121f.Background != nil {
		__obf_8503f0efd7228247 := __obf_8491b090f9e1121f.Background
		__obf_50644a3aac3d3944 := __obf_8503f0efd7228247.Bounds()
		__obf_02e7169cdc485ac6 := helper.CreateNRGBACanvas(__obf_50644a3aac3d3944.Dx(), __obf_50644a3aac3d3944.Dy(), true)
		__obf_7d525f2b0cc78797 := util.RangCutImagePos(__obf_8491b090f9e1121f.Size, __obf_8491b090f9e1121f.Size, __obf_8503f0efd7228247)
		draw.Draw(__obf_02e7169cdc485ac6, __obf_50644a3aac3d3944, __obf_8503f0efd7228247, __obf_7d525f2b0cc78797, draw.Over)
		__obf_02e7169cdc485ac6.
			SubImage(image.Rect(0, 0, __obf_8491b090f9e1121f.Size, __obf_8491b090f9e1121f.Size))
		draw.Draw(__obf_2d0ece5bcf1fe46a, __obf_2d0ece5bcf1fe46a.Bounds(), __obf_02e7169cdc485ac6, image.Point{}, draw.Over)
	}
	__obf_2d0ece5bcf1fe46a.
		CropCircle(__obf_2d0ece5bcf1fe46a.Bounds().Dx()/2, __obf_2d0ece5bcf1fe46a.Bounds().Dy()/2, __obf_2d0ece5bcf1fe46a.Bounds().Dy()/2)
	return __obf_2d0ece5bcf1fe46a, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_598561cd40842fdf *Captcha) DrawWithCropCircle(__obf_8491b090f9e1121f *DrawCropCircleImageParams) (image.Image, error) {
	__obf_8503f0efd7228247 := __obf_8491b090f9e1121f.Background
	__obf_42e7e0d9c8ed06ae := __obf_8503f0efd7228247.Bounds()
	__obf_55a2e60902f1c1d6 := helper.CreateNRGBACanvas(__obf_8503f0efd7228247.Bounds().Dx(), __obf_8503f0efd7228247.Bounds().Dy(), true)
	draw.Draw(__obf_55a2e60902f1c1d6, __obf_8503f0efd7228247.Bounds(), __obf_8503f0efd7228247, image.Point{}, draw.Over)
	__obf_55a2e60902f1c1d6.
		CropScaleCircle(__obf_8503f0efd7228247.Bounds().Dx()/2, __obf_8503f0efd7228247.Bounds().Dy()/2, __obf_8503f0efd7228247.Bounds().Dy()/2, __obf_8491b090f9e1121f.ScaleRatioSize)
	__obf_55a2e60902f1c1d6.
		Rotate(__obf_8491b090f9e1121f.Angle, true)
	__obf_a3aa3fc4cb6e09f5 := __obf_55a2e60902f1c1d6.Bounds()
	if __obf_a3aa3fc4cb6e09f5.Dy() > __obf_42e7e0d9c8ed06ae.Dy() || __obf_a3aa3fc4cb6e09f5.Dx() > __obf_42e7e0d9c8ed06ae.Dx() {
		__obf_0e55f449019f083b := helper.CreateNRGBACanvas(__obf_8503f0efd7228247.Bounds().Dx(), __obf_8503f0efd7228247.Bounds().Dy(), true)
		draw.Draw(__obf_0e55f449019f083b, __obf_0e55f449019f083b.Bounds(), __obf_55a2e60902f1c1d6, image.Point{X: (__obf_a3aa3fc4cb6e09f5.Dy() - __obf_42e7e0d9c8ed06ae.Dy()) / 2, Y: (__obf_a3aa3fc4cb6e09f5.Dx() - __obf_42e7e0d9c8ed06ae.Dx()) / 2}, draw.Over)
		__obf_55a2e60902f1c1d6 = __obf_0e55f449019f083b
	}

	return __obf_55a2e60902f1c1d6, nil
}

func (__obf_598561cd40842fdf *Captcha) __obf_71ef3916f37c74fa(__obf_3acfd648cd265204 int) (image.Image, error) {
	__obf_c57e727dde257304, __obf_8e9665c84006346f := assets.PickBgImage()
	if __obf_8e9665c84006346f != nil {
		return nil, __obf_8e9665c84006346f
	}

	return __obf_598561cd40842fdf.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_3acfd648cd265204,
		Background: __obf_c57e727dde257304,
		Alpha:      __obf_598561cd40842fdf.__obf_5c1359684eb79b70.Primary.Alpha,
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
func (__obf_598561cd40842fdf *Captcha) __obf_a56199890c419472(__obf_8503f0efd7228247 image.Image, __obf_720ba9ea19239003 int, __obf_5dfaeb3c1bd2a70a int) (image.Image, error) {
	return __obf_598561cd40842fdf.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_8503f0efd7228247,
		Alpha:          __obf_598561cd40842fdf.__obf_5c1359684eb79b70.Secondary.Alpha,
		Size:           __obf_5dfaeb3c1bd2a70a,
		Angle:          __obf_720ba9ea19239003,
		ScaleRatioSize: (__obf_598561cd40842fdf.__obf_5c1359684eb79b70.Primary.Size - __obf_5dfaeb3c1bd2a70a) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_598561cd40842fdf *Captcha) __obf_c3950e0945f974a4() int {
	__obf_a239c34dc66cea7d := __obf_598561cd40842fdf.__obf_5c1359684eb79b70.Primary.AnglePosRange
	__obf_08bcf3a7a56fe25b := rand.IntN(len(__obf_a239c34dc66cea7d))
	if __obf_08bcf3a7a56fe25b < 0 {
		return 0
	}
	__obf_720ba9ea19239003 := __obf_a239c34dc66cea7d[__obf_08bcf3a7a56fe25b]
	return util.RandInt(__obf_720ba9ea19239003.Min, __obf_720ba9ea19239003.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_598561cd40842fdf *Captcha) __obf_858466f65ad88c3e() int {
	__obf_ec8184ff5a96c4b5 := __obf_598561cd40842fdf.__obf_5c1359684eb79b70.Secondary.SizeRange
	__obf_08bcf3a7a56fe25b := rand.IntN(len(__obf_ec8184ff5a96c4b5))
	if __obf_08bcf3a7a56fe25b < 0 {
		return __obf_ec8184ff5a96c4b5[0]
	}

	return __obf_ec8184ff5a96c4b5[__obf_08bcf3a7a56fe25b]
}
