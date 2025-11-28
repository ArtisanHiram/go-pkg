package __obf_7b43310cbd758abe

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
	__obf_c8b2a472ca3619f7 *Options
}

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_c8b2a472ca3619f7 *Options) *Captcha {
	if __obf_c8b2a472ca3619f7 == nil {
		return nil
	}
	return &Captcha{
		__obf_c8b2a472ca3619f7: __obf_c8b2a472ca3619f7,
	}
}

// func randomSlideType(opts *Options) SlideType {
// 	slideType := SlideType(rand.IntN(2))
// 	if slideType == Move {
// 		opts.Secondary.DeadZoneDirections = []DeadZoneDirectionType{DeadZoneDirectionTypeLeft}
// 		opts.Secondary.EnableVerticalRandom = false
// 	} else {
// 		opts.Secondary.DeadZoneDirections = []DeadZoneDirectionType{
// 			DeadZoneDirectionTypeLeft,
// 			DeadZoneDirectionTypeRight,
// 			DeadZoneDirectionTypeBottom,
// 			DeadZoneDirectionTypeTop,
// 		}
// 		opts.Secondary.EnableVerticalRandom = true
// 	}
// 	return slideType
// }

// DrawWithNRGBA draws the main CAPTCHA image and background image using NRGBA format
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn CAPTCHA image
//   - image.Image: Drawn background image
//   - error: Error information
func (__obf_533af04904ee7f62 *Captcha) DrawWithNRGBA(__obf_48c134abb5c806c2 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_f511dde7f6376853 := __obf_48c134abb5c806c2.DrawBlocks
	__obf_029145c292029885 := helper.CreateNRGBACanvas(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, true)

	for _, __obf_8fe11b25bcff626c := range __obf_f511dde7f6376853 {
		__obf_0da34411b16aae43, __obf_815a32f6170a5e37 := __obf_533af04904ee7f62.__obf_77358da311ce81a5(__obf_8fe11b25bcff626c.Width, __obf_8fe11b25bcff626c.Height, __obf_8fe11b25bcff626c.Shape)
		if __obf_815a32f6170a5e37 != nil {
			return nil, nil, __obf_815a32f6170a5e37
		}

		__obf_cd8f4831e03f6f74 := __obf_0da34411b16aae43.Bounds()
		draw.Draw(__obf_029145c292029885, image.Rect(__obf_8fe11b25bcff626c.X, __obf_8fe11b25bcff626c.Y, __obf_8fe11b25bcff626c.X+__obf_cd8f4831e03f6f74.Dx(), __obf_8fe11b25bcff626c.Y+__obf_cd8f4831e03f6f74.Dy()), __obf_0da34411b16aae43, image.Point{}, draw.Over)
	}

	__obf_a7c0b3cec9f0ffae := helper.CreateNRGBACanvas(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, true)
	if __obf_48c134abb5c806c2.Background != nil {
		__obf_5a820796777ed357 := __obf_48c134abb5c806c2.Background
		__obf_8fe11b25bcff626c := __obf_5a820796777ed357.Bounds()
		__obf_999af3f415c9b80f := helper.CreateNRGBACanvas(__obf_8fe11b25bcff626c.Dx(), __obf_8fe11b25bcff626c.Dy(), true)
		__obf_d20461f6caa9aeb8 := util.RangCutImagePos(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, __obf_5a820796777ed357)
		draw.Draw(__obf_999af3f415c9b80f, __obf_8fe11b25bcff626c, __obf_5a820796777ed357, __obf_d20461f6caa9aeb8, draw.Src)
		__obf_999af3f415c9b80f.SubImage(image.Rect(0, 0, __obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height))

		draw.Draw(__obf_a7c0b3cec9f0ffae, __obf_a7c0b3cec9f0ffae.Bounds(), __obf_999af3f415c9b80f, image.Point{}, draw.Over)
		draw.Draw(__obf_999af3f415c9b80f, __obf_029145c292029885.Bounds(), __obf_029145c292029885, image.Point{}, draw.Over)
		return __obf_999af3f415c9b80f, __obf_a7c0b3cec9f0ffae, nil
	}

	return __obf_029145c292029885, __obf_a7c0b3cec9f0ffae, nil
}

// DrawWithTemplate draws the tile image using a template
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn tile image
//   - error: Error information
func (__obf_533af04904ee7f62 *Captcha) DrawWithTemplate(__obf_48c134abb5c806c2 *DrawTplImageParam) (image.Image, error) {
	__obf_48a33b6401bf2a46 := __obf_48c134abb5c806c2.Block
	__obf_5a820796777ed357 := __obf_48c134abb5c806c2.Background
	__obf_029145c292029885 := helper.CreateNRGBACanvas(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, true)
	__obf_035f1a3ab0bffc99 := helper.CreateNRGBACanvas(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, true)

	__obf_4eddc11b44c141f9, __obf_815a32f6170a5e37 := __obf_533af04904ee7f62.__obf_77358da311ce81a5(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, __obf_48c134abb5c806c2.Mask)
	if __obf_815a32f6170a5e37 != nil {
		return nil, __obf_815a32f6170a5e37
	}

	draw.Draw(__obf_035f1a3ab0bffc99, __obf_035f1a3ab0bffc99.Bounds(), __obf_5a820796777ed357, image.Pt(__obf_48a33b6401bf2a46.X, __obf_48a33b6401bf2a46.Y), draw.Src)
	draw.DrawMask(__obf_029145c292029885, __obf_4eddc11b44c141f9.Bounds(), __obf_035f1a3ab0bffc99, image.Point{}, __obf_4eddc11b44c141f9, image.Point{}, draw.Over)

	__obf_59aac214cadf1763, __obf_815a32f6170a5e37 := __obf_533af04904ee7f62.__obf_77358da311ce81a5(__obf_48c134abb5c806c2.Width, __obf_48c134abb5c806c2.Height, __obf_48a33b6401bf2a46.Shape)
	if __obf_815a32f6170a5e37 != nil {
		return nil, __obf_815a32f6170a5e37
	}
	draw.Draw(__obf_029145c292029885, __obf_59aac214cadf1763.Bounds(), __obf_59aac214cadf1763, image.Point{}, draw.Over)

	return __obf_029145c292029885, nil
}

// drawGraphImage draws a graph image
// params:
//   - width: Image width
//   - height: Image height
//   - img: Input image
//
// returns:
//   - types.NRGBA: Drawn graph canvas
//   - error: Error information
func (__obf_533af04904ee7f62 *Captcha) __obf_77358da311ce81a5(__obf_9f93f567a968ba47, __obf_ac7d7ce25ce27722 int, __obf_9a60feaf318233a4 image.Image) (*types.NRGBA, error) {
	__obf_029145c292029885 := helper.CreateNRGBACanvas(__obf_9f93f567a968ba47, __obf_ac7d7ce25ce27722, true)
	draw.BiLinear.Scale(__obf_029145c292029885, __obf_029145c292029885.Bounds(), __obf_9a60feaf318233a4, __obf_9a60feaf318233a4.Bounds(), draw.Over, nil)
	return __obf_029145c292029885, nil
}

// Generate generates slide CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_533af04904ee7f62 *Captcha) Generate() (*CaptchaData, error) {

	// slideType := randomSlideType(c.opts)
	__obf_0162179ca3c66d2d, __obf_815a32f6170a5e37 := assets.PickTile()
	if __obf_815a32f6170a5e37 != nil {
		return nil, __obf_815a32f6170a5e37
	}

	__obf_51e1d431c148d66c := util.RandInt(__obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Secondary.CountRange.Min, __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Secondary.CountRange.Max)
	__obf_f511dde7f6376853, __obf_bccc568563fb03d7 := __obf_533af04904ee7f62.__obf_3d4a564438f48053(__obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Primary.Size, __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Secondary.SizeRange, __obf_51e1d431c148d66c)
	var __obf_48a33b6401bf2a46 *types.Block
	if len(__obf_f511dde7f6376853) > 1 {
		__obf_fc4fc52d8df5628b := max(rand.IntN(len(__obf_f511dde7f6376853)), 0)
		__obf_48a33b6401bf2a46 = __obf_f511dde7f6376853[__obf_fc4fc52d8df5628b]
	} else {
		__obf_48a33b6401bf2a46 = __obf_f511dde7f6376853[0]
	}

	__obf_ae9d075a50d2eeb2, __obf_0f8b583ed46bc13c, __obf_815a32f6170a5e37 := __obf_533af04904ee7f62.__obf_744b5ecd3ca3903a(__obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Primary.Size, __obf_0162179ca3c66d2d.Shadow, __obf_f511dde7f6376853)
	if __obf_815a32f6170a5e37 != nil {
		return nil, __obf_815a32f6170a5e37
	}

	__obf_48a33b6401bf2a46.Shape = __obf_0162179ca3c66d2d.Overlay

	__obf_91551e891ceb13f1, __obf_815a32f6170a5e37 := __obf_533af04904ee7f62.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_0f8b583ed46bc13c,
		Mask:       __obf_0162179ca3c66d2d.Mask,
		Alpha:      __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Primary.Alpha,
		Width:      __obf_48a33b6401bf2a46.Width,
		Height:     __obf_48a33b6401bf2a46.Height,
		Block:      __obf_48a33b6401bf2a46,
	})

	if __obf_815a32f6170a5e37 != nil {
		return nil, __obf_815a32f6170a5e37
	}

	if __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Type == Move {
		__obf_48a33b6401bf2a46.DY = __obf_48a33b6401bf2a46.Y
	} else {
		__obf_48a33b6401bf2a46.DY = __obf_bccc568563fb03d7.Y
	}
	__obf_48a33b6401bf2a46.DX = __obf_bccc568563fb03d7.X

	return &CaptchaData{
		__obf_782e73e3bc2a2b90: __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Type,
		__obf_48a33b6401bf2a46: __obf_48a33b6401bf2a46,
		__obf_ae9d075a50d2eeb2: types.NewJPEGImage(__obf_ae9d075a50d2eeb2),
		__obf_91551e891ceb13f1: types.NewPNGImage(__obf_91551e891ceb13f1),
	}, nil
}

// genMainImage generates the main CAPTCHA image and background image
// params:
//   - size: Image size
//   - shadowImage: Shadow image
//   - blocks: List of blocks
//
// returns:
//   - image.Image: Main image
//   - image.Image: Background image
//   - error: Error information
func (__obf_533af04904ee7f62 *Captcha) __obf_744b5ecd3ca3903a(__obf_aad00d3034ed9067 types.Size, __obf_484e609f6e7cdb65 image.Image, __obf_f511dde7f6376853 []*types.Block) (image.Image, image.Image, error) {
	var __obf_b404d62ec25d559e = make([]*types.Block, 0, len(__obf_f511dde7f6376853))
	for _, __obf_3b8f88980dd18433 := range __obf_f511dde7f6376853 {
		__obf_3b8f88980dd18433.Shape = __obf_484e609f6e7cdb65
		__obf_b404d62ec25d559e = append(__obf_b404d62ec25d559e, __obf_3b8f88980dd18433)
	}

	__obf_1fee37d0587ba009, __obf_815a32f6170a5e37 := assets.PickBgImage()
	if __obf_815a32f6170a5e37 != nil {
		return nil, nil, __obf_815a32f6170a5e37
	}

	return __obf_533af04904ee7f62.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_aad00d3034ed9067.Width,
		Height:     __obf_aad00d3034ed9067.Height,
		Background: __obf_1fee37d0587ba009,
		Alpha:      __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Primary.Alpha,
		DrawBlocks: __obf_b404d62ec25d559e,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_533af04904ee7f62 *Captcha) __obf_80b7b3130efa2134() DeadZoneDirectionType {
	__obf_3682c63519ae1a3a := __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Secondary.DeadZoneDirections

	__obf_fc4fc52d8df5628b := rand.IntN(len(__obf_3682c63519ae1a3a))
	if __obf_fc4fc52d8df5628b < 0 {
		return 0
	}

	return __obf_3682c63519ae1a3a[__obf_fc4fc52d8df5628b]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_533af04904ee7f62 *Captcha) __obf_0f693a765ef8a07e() int {
	__obf_62c2b10d6cd85c0a := __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Secondary.AnglePosRange

	__obf_fc4fc52d8df5628b := rand.IntN(len(__obf_62c2b10d6cd85c0a))
	if __obf_fc4fc52d8df5628b < 0 {
		return 0
	}

	__obf_e88bd40baf6aeb5b := __obf_62c2b10d6cd85c0a[__obf_fc4fc52d8df5628b]
	return util.RandInt(__obf_e88bd40baf6aeb5b.Min, __obf_e88bd40baf6aeb5b.Max)
}

// genGraphBlocks generates graph block data
// params:
//   - imageSize: Main image size
//   - size: Graph size range
//   - length: Number of graphs
//
// returns:
//   - []*types.Block: List of blocks
//   - types.Point: Tile position
func (__obf_533af04904ee7f62 *Captcha) __obf_3d4a564438f48053(__obf_6cb0f928dfe5f5c6 types.Size, __obf_aad00d3034ed9067 types.Range, __obf_4a4100bf291d4edd int) ([]*types.Block, types.Point) {
	var __obf_f511dde7f6376853 = make([]*types.Block, 0, __obf_4a4100bf291d4edd)
	__obf_9f93f567a968ba47 := __obf_6cb0f928dfe5f5c6.Width
	__obf_ac7d7ce25ce27722 := __obf_6cb0f928dfe5f5c6.Height

	__obf_da585450f15dd6d9 := __obf_533af04904ee7f62.__obf_0f693a765ef8a07e()
	__obf_7f24b3201ab650d3 := util.RandInt(__obf_aad00d3034ed9067.Min, __obf_aad00d3034ed9067.Max)
	__obf_9715a8e4acb8e6bb := __obf_7f24b3201ab650d3
	__obf_b4f3e3681701efb0 := __obf_7f24b3201ab650d3

	__obf_8257a1eaf29026f3 := __obf_533af04904ee7f62.__obf_80b7b3130efa2134()
	__obf_4dbdd2b602f21ae0 := __obf_b4f3e3681701efb0 / 2
	__obf_b988f923318ff5d1 := (__obf_9f93f567a968ba47 - __obf_b4f3e3681701efb0 - 20) / __obf_4a4100bf291d4edd
	__obf_7ebb747a5fcb895a := __obf_533af04904ee7f62.__obf_8229eedd27fe0d0b(5, __obf_ac7d7ce25ce27722-__obf_9715a8e4acb8e6bb-5, __obf_9715a8e4acb8e6bb, __obf_8257a1eaf29026f3)

	for __obf_26220a125481b269 := range __obf_4a4100bf291d4edd {
		var __obf_48a33b6401bf2a46 = &types.Block{}
		__obf_afcfed03799ee8b2, __obf_b630e2c195832bb0 := __obf_533af04904ee7f62.__obf_fb9f97fea4c7c64d((__obf_26220a125481b269*__obf_b988f923318ff5d1)+__obf_4dbdd2b602f21ae0+5, ((__obf_26220a125481b269+1)*__obf_b988f923318ff5d1)-__obf_4dbdd2b602f21ae0, __obf_b4f3e3681701efb0, __obf_8257a1eaf29026f3)

		__obf_afcfed03799ee8b2 = int(max(float64(__obf_afcfed03799ee8b2), float64(__obf_4dbdd2b602f21ae0+5)))
		__obf_48a33b6401bf2a46.X = util.RandInt(__obf_afcfed03799ee8b2+20, __obf_b630e2c195832bb0+20) - __obf_4dbdd2b602f21ae0

		if __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Secondary.EnableVerticalRandom {
			__obf_7ebb747a5fcb895a = __obf_533af04904ee7f62.__obf_8229eedd27fe0d0b(5, __obf_ac7d7ce25ce27722-__obf_9715a8e4acb8e6bb-5, __obf_9715a8e4acb8e6bb, __obf_8257a1eaf29026f3)
		}

		__obf_48a33b6401bf2a46.Y = __obf_7ebb747a5fcb895a
		__obf_48a33b6401bf2a46.Width = __obf_b4f3e3681701efb0
		__obf_48a33b6401bf2a46.Height = __obf_9715a8e4acb8e6bb
		__obf_48a33b6401bf2a46.Angle = __obf_da585450f15dd6d9

		__obf_f511dde7f6376853 = append(__obf_f511dde7f6376853, __obf_48a33b6401bf2a46)
	}

	__obf_d20461f6caa9aeb8 := types.Point{}
	if __obf_533af04904ee7f62.__obf_c8b2a472ca3619f7.Type == Move {
		__obf_d20461f6caa9aeb8.X = util.RandInt(5, __obf_4dbdd2b602f21ae0)
		__obf_d20461f6caa9aeb8.Y = __obf_7ebb747a5fcb895a
		return __obf_f511dde7f6376853, __obf_d20461f6caa9aeb8
	}

	switch __obf_8257a1eaf29026f3 {
	case DeadZoneDirectionTypeTop:
		__obf_d20461f6caa9aeb8.X = util.RandInt(5, __obf_9f93f567a968ba47-__obf_b4f3e3681701efb0-5)
		__obf_d20461f6caa9aeb8.Y = 5
	case DeadZoneDirectionTypeBottom:
		__obf_d20461f6caa9aeb8.X = util.RandInt(5, __obf_9f93f567a968ba47-__obf_b4f3e3681701efb0-5)
		__obf_d20461f6caa9aeb8.Y = __obf_ac7d7ce25ce27722 - __obf_9715a8e4acb8e6bb - 5
	case DeadZoneDirectionTypeLeft:
		__obf_d20461f6caa9aeb8.X = 5
		__obf_d20461f6caa9aeb8.Y = util.RandInt(5, __obf_ac7d7ce25ce27722-__obf_9715a8e4acb8e6bb-5)
	case DeadZoneDirectionTypeRight:
		__obf_d20461f6caa9aeb8.X = __obf_9f93f567a968ba47 - __obf_b4f3e3681701efb0 - 5
		__obf_d20461f6caa9aeb8.Y = util.RandInt(5, __obf_ac7d7ce25ce27722-__obf_9715a8e4acb8e6bb-5)
	}

	return __obf_f511dde7f6376853, __obf_d20461f6caa9aeb8
}

// calcXWithDeadZone calculates the X coordinate range (considering dead zone)
// params:
//   - start: Start X coordinate
//   - end: End X coordinate
//   - value: types.Block width
//   - dzdType: Dead zone direction
//
// returns:
//   - int: Adjusted start X coordinate
//   - int: Adjusted end X coordinate
func (__obf_533af04904ee7f62 *Captcha) __obf_fb9f97fea4c7c64d(__obf_afcfed03799ee8b2, __obf_b630e2c195832bb0, __obf_72a033fd9701305a int, __obf_8257a1eaf29026f3 DeadZoneDirectionType) (int, int) {
	if __obf_8257a1eaf29026f3 == DeadZoneDirectionTypeLeft {
		__obf_afcfed03799ee8b2 += __obf_72a033fd9701305a
		__obf_b630e2c195832bb0 += __obf_72a033fd9701305a
	}
	return __obf_afcfed03799ee8b2, __obf_b630e2c195832bb0
}

// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
// params:
//   - start: Start Y coordinate
//   - end: End Y coordinate
//   - value: types.Block height
//   - dzdType: Dead zone direction
//
// return: Random Y coordinate
func (__obf_533af04904ee7f62 *Captcha) __obf_8229eedd27fe0d0b(__obf_afcfed03799ee8b2, __obf_b630e2c195832bb0, __obf_72a033fd9701305a int, __obf_8257a1eaf29026f3 DeadZoneDirectionType) int {
	switch __obf_8257a1eaf29026f3 {
	case DeadZoneDirectionTypeTop:
		__obf_afcfed03799ee8b2 += __obf_72a033fd9701305a
	case DeadZoneDirectionTypeBottom:
		__obf_b630e2c195832bb0 -= __obf_72a033fd9701305a
	}
	return util.RandInt(__obf_afcfed03799ee8b2, __obf_b630e2c195832bb0)
}
