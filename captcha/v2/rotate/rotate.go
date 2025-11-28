package __obf_48d1d3ff050bec1d

import (
	"image"
	"math/rand/v2"

	core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"
	"golang.org/x/image/draw"
)

// Captcha is the concrete implementation of the Captcha interface
type Captcha struct {
	__obf_6e97b0b1f3cdd756 *Options
}

// NewCaptcha creates a new rotate CAPTCHA instance
func NewCaptcha(__obf_5a4bb3acf7fdb663 Option) *Captcha {
	__obf_6e97b0b1f3cdd756 := &Options{
		Primary: Primary{
			Size:  220,
			Alpha: 1,
			AnglePosRange: []core.Range{
				{Min: 30, Max: 330},
			},
		},
		Secondary: Secondary{
			Alpha:     1,
			SizeRange: []int{140, 150, 160, 170},
		},
	}

	if __obf_5a4bb3acf7fdb663 != nil {
		__obf_5a4bb3acf7fdb663(__obf_6e97b0b1f3cdd756)
	}

	return &Captcha{
		__obf_6e97b0b1f3cdd756: __obf_6e97b0b1f3cdd756,
	}
}

// DrawWithNRGBA draws the main CAPTCHA image using NRGBA format
// params:
//   - params: Drawing parameters
//
// return:
//   - image.Image: Drawn image
//   - error: Error information
func (__obf_1eb094def04182e9 *Captcha) DrawWithNRGBA(__obf_b0a37a603ca39816 *DrawImageParams) (image.Image, error) {
	__obf_8c113a13017a3594 := core.CreateNRGBACanvas(__obf_b0a37a603ca39816.Size, __obf_b0a37a603ca39816.Size, true)
	if __obf_b0a37a603ca39816.Background != nil {
		__obf_bce8952e92a3770a := __obf_b0a37a603ca39816.Background
		__obf_fed681f10f27966d := __obf_bce8952e92a3770a.Bounds()
		__obf_f5172f0f413cdf5b := core.CreateNRGBACanvas(__obf_fed681f10f27966d.Dx(), __obf_fed681f10f27966d.Dy(), true)
		__obf_3b88b8d598fdcdc4 := core.RangCutImagePos(__obf_b0a37a603ca39816.Size, __obf_b0a37a603ca39816.Size, __obf_bce8952e92a3770a)
		draw.Draw(__obf_f5172f0f413cdf5b, __obf_fed681f10f27966d, __obf_bce8952e92a3770a, __obf_3b88b8d598fdcdc4, draw.Over)
		__obf_f5172f0f413cdf5b.SubImage(image.Rect(0, 0, __obf_b0a37a603ca39816.Size, __obf_b0a37a603ca39816.Size))
		draw.Draw(__obf_8c113a13017a3594, __obf_8c113a13017a3594.Bounds(), __obf_f5172f0f413cdf5b, image.Point{}, draw.Over)
	}

	__obf_8c113a13017a3594.CropCircle(__obf_8c113a13017a3594.Bounds().Dx()/2, __obf_8c113a13017a3594.Bounds().Dy()/2, __obf_8c113a13017a3594.Bounds().Dy()/2)
	return __obf_8c113a13017a3594, nil
}

// DrawWithCropCircle draws a cropped circle image (thumbnail)
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn thumbnail image
//   - error: Error information
func (__obf_1eb094def04182e9 *Captcha) DrawWithCropCircle(__obf_b0a37a603ca39816 *DrawCropCircleImageParams) (image.Image, error) {
	__obf_bce8952e92a3770a := __obf_b0a37a603ca39816.Background

	__obf_dde5c8d21b438959 := __obf_bce8952e92a3770a.Bounds()
	__obf_ea8af86131c045da := core.CreateNRGBACanvas(__obf_bce8952e92a3770a.Bounds().Dx(), __obf_bce8952e92a3770a.Bounds().Dy(), true)
	draw.Draw(__obf_ea8af86131c045da, __obf_bce8952e92a3770a.Bounds(), __obf_bce8952e92a3770a, image.Point{}, draw.Over)
	__obf_ea8af86131c045da.CropScaleCircle(__obf_bce8952e92a3770a.Bounds().Dx()/2, __obf_bce8952e92a3770a.Bounds().Dy()/2, __obf_bce8952e92a3770a.Bounds().Dy()/2, __obf_b0a37a603ca39816.ScaleRatioSize)
	__obf_ea8af86131c045da.Rotate(__obf_b0a37a603ca39816.Angle, true)

	__obf_5132077bf224c2e0 := __obf_ea8af86131c045da.Bounds()
	if __obf_5132077bf224c2e0.Dy() > __obf_dde5c8d21b438959.Dy() || __obf_5132077bf224c2e0.Dx() > __obf_dde5c8d21b438959.Dx() {
		__obf_caf1f9b26131c881 := core.CreateNRGBACanvas(__obf_bce8952e92a3770a.Bounds().Dx(), __obf_bce8952e92a3770a.Bounds().Dy(), true)
		draw.Draw(__obf_caf1f9b26131c881, __obf_caf1f9b26131c881.Bounds(), __obf_ea8af86131c045da, image.Point{X: (__obf_5132077bf224c2e0.Dy() - __obf_dde5c8d21b438959.Dy()) / 2, Y: (__obf_5132077bf224c2e0.Dx() - __obf_dde5c8d21b438959.Dx()) / 2}, draw.Over)
		__obf_ea8af86131c045da = __obf_caf1f9b26131c881
	}

	return __obf_ea8af86131c045da, nil
}

// Generate generates rotate CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_1eb094def04182e9 *Captcha) Generate() (*CaptchaData, error) {
	__obf_4b064127547c217f := __obf_1eb094def04182e9.__obf_5100b381091dd701()

	__obf_1f9603774b9e6d25, __obf_a2a9e3a81885fdd2 := __obf_1eb094def04182e9.__obf_56324bb5762710df(__obf_1eb094def04182e9.__obf_6e97b0b1f3cdd756.Primary.Size)
	if __obf_a2a9e3a81885fdd2 != nil {
		return nil, __obf_a2a9e3a81885fdd2
	}
	__obf_168883ae03b7841e := __obf_1eb094def04182e9.__obf_5483e9b9ae3d98f2()
	__obf_ffe34c16f5089ec9, __obf_a2a9e3a81885fdd2 := __obf_1eb094def04182e9.__obf_7b2eb253e4ea5ef9(__obf_1f9603774b9e6d25, __obf_168883ae03b7841e, __obf_4b064127547c217f)
	if __obf_a2a9e3a81885fdd2 != nil {
		return nil, __obf_a2a9e3a81885fdd2
	}

	return &CaptchaData{
		__obf_168883ae03b7841e: __obf_168883ae03b7841e,
		__obf_1f9603774b9e6d25: core.NewPNGImage(__obf_1f9603774b9e6d25),
		__obf_ffe34c16f5089ec9: core.NewPNGImage(__obf_ffe34c16f5089ec9),
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
func (__obf_1eb094def04182e9 *Captcha) __obf_56324bb5762710df(__obf_2699386ee5d9e413 int) (image.Image, error) {

	__obf_9905389c2f821c13, __obf_a2a9e3a81885fdd2 := core.PickBgImage()
	if __obf_a2a9e3a81885fdd2 != nil {
		return nil, __obf_a2a9e3a81885fdd2
	}

	return __obf_1eb094def04182e9.DrawWithNRGBA(&DrawImageParams{
		Size:       __obf_2699386ee5d9e413,
		Background: __obf_9905389c2f821c13,
		Alpha:      __obf_1eb094def04182e9.__obf_6e97b0b1f3cdd756.Primary.Alpha,
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
func (__obf_1eb094def04182e9 *Captcha) __obf_7b2eb253e4ea5ef9(__obf_bce8952e92a3770a image.Image, __obf_168883ae03b7841e int, __obf_4b064127547c217f int) (image.Image, error) {
	return __obf_1eb094def04182e9.DrawWithCropCircle(&DrawCropCircleImageParams{
		Background:     __obf_bce8952e92a3770a,
		Alpha:          __obf_1eb094def04182e9.__obf_6e97b0b1f3cdd756.Secondary.Alpha,
		Size:           __obf_4b064127547c217f,
		Angle:          __obf_168883ae03b7841e,
		ScaleRatioSize: (__obf_1eb094def04182e9.__obf_6e97b0b1f3cdd756.Primary.Size - __obf_4b064127547c217f) / 2,
	})
}

// randAngle generates a random angle
// returns: Random angle value
func (__obf_1eb094def04182e9 *Captcha) __obf_5483e9b9ae3d98f2() int {
	__obf_2f76020e4a23ca68 := __obf_1eb094def04182e9.__obf_6e97b0b1f3cdd756.Primary.AnglePosRange

	__obf_818c41ab8bcdfdea := rand.IntN(len(__obf_2f76020e4a23ca68))
	if __obf_818c41ab8bcdfdea < 0 {
		return 0
	}

	__obf_168883ae03b7841e := __obf_2f76020e4a23ca68[__obf_818c41ab8bcdfdea]
	return core.RandInt(__obf_168883ae03b7841e.Min, __obf_168883ae03b7841e.Max)
}

// randThumbSize generates a random thumbnail size
// return: Random thumbnail size
func (__obf_1eb094def04182e9 *Captcha) __obf_5100b381091dd701() int {
	__obf_b5eb6ca8b4512705 := __obf_1eb094def04182e9.__obf_6e97b0b1f3cdd756.Secondary.SizeRange

	__obf_818c41ab8bcdfdea := rand.IntN(len(__obf_b5eb6ca8b4512705))
	if __obf_818c41ab8bcdfdea < 0 {
		return __obf_b5eb6ca8b4512705[0]
	}

	return __obf_b5eb6ca8b4512705[__obf_818c41ab8bcdfdea]
}
