package __obf_ae0bbdaffaa1f242

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
	__obf_7828a22179b13d85 *types.SlideOption
}

var _ types.Captcha = (*Captcha)(nil)

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_7828a22179b13d85 *types.SlideOption) *Captcha {
	if __obf_7828a22179b13d85 == nil {
		return nil
	}
	return &Captcha{__obf_7828a22179b13d85: __obf_7828a22179b13d85}
}

// func (c *Captcha) Type() types.CaptchaType {
// 	if c.opt.Type == Move {
// 		return types.MoveCaptchat
// 	}
// 	return types.DragCaptchat
// }

func (__obf_9e636908868342b9 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_63a4c78e3e074756,

		// slideType := randomSlideType(c.opt)
		__obf_d5b002b58d3eb1d0 := assets.PickTile()
	if __obf_d5b002b58d3eb1d0 != nil {
		return nil, __obf_d5b002b58d3eb1d0
	}
	__obf_46054e7567628a40 := util.RandInt(__obf_9e636908868342b9.__obf_7828a22179b13d85.Secondary.CountRange.Min, __obf_9e636908868342b9.__obf_7828a22179b13d85.Secondary.CountRange.Max)
	__obf_3ca57388aa6752a1, __obf_dab50866ad7f6c72 := __obf_9e636908868342b9.__obf_a09bc2563bda22ac(__obf_9e636908868342b9.__obf_7828a22179b13d85.Primary.Width, __obf_9e636908868342b9.__obf_7828a22179b13d85.Primary.Height, __obf_9e636908868342b9.__obf_7828a22179b13d85.Secondary.SizeRange, __obf_46054e7567628a40)
	var __obf_73053d5754e12223 *types.Block
	if len(__obf_3ca57388aa6752a1) > 1 {
		__obf_4291897e510adc66 := max(rand.IntN(len(__obf_3ca57388aa6752a1)), 0)
		__obf_73053d5754e12223 = __obf_3ca57388aa6752a1[__obf_4291897e510adc66]
	} else {
		__obf_73053d5754e12223 = __obf_3ca57388aa6752a1[0]
	}
	__obf_f37f373d0317b914, __obf_9c9372ff8bab5a13, __obf_d5b002b58d3eb1d0 := __obf_9e636908868342b9.__obf_c64986d404c5befe(__obf_9e636908868342b9.__obf_7828a22179b13d85.Primary.Width, __obf_9e636908868342b9.__obf_7828a22179b13d85.Primary.Height, __obf_63a4c78e3e074756.Shadow, __obf_3ca57388aa6752a1)
	if __obf_d5b002b58d3eb1d0 != nil {
		return nil, __obf_d5b002b58d3eb1d0
	}
	__obf_73053d5754e12223.
		Shape = __obf_63a4c78e3e074756.Overlay
	__obf_30c3995a9b03c58a, __obf_d5b002b58d3eb1d0 := __obf_9e636908868342b9.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_9c9372ff8bab5a13,
		Mask:       __obf_63a4c78e3e074756.Mask,
		Alpha:      __obf_9e636908868342b9.__obf_7828a22179b13d85.Primary.Alpha,
		Width:      __obf_73053d5754e12223.Width,
		Height:     __obf_73053d5754e12223.Height,
		Block:      __obf_73053d5754e12223,
	})

	if __obf_d5b002b58d3eb1d0 != nil {
		return nil, __obf_d5b002b58d3eb1d0
	}

	if __obf_9e636908868342b9.__obf_7828a22179b13d85.Type == types.MoveSlide {
		__obf_73053d5754e12223.
			DY = __obf_73053d5754e12223.Y
	} else {
		__obf_73053d5754e12223.
			DY = __obf_dab50866ad7f6c72.Y
	}
	__obf_73053d5754e12223.
		DX = __obf_dab50866ad7f6c72.X

	return &CaptchaData{__obf_a805d52fbb2d47f1: __obf_9e636908868342b9.__obf_7828a22179b13d85.Type, __obf_73053d5754e12223: __obf_73053d5754e12223, __obf_f37f373d0317b914: types.NewJPEGImage(__obf_f37f373d0317b914), __obf_30c3995a9b03c58a: types.NewPNGImage(__obf_30c3995a9b03c58a)}, nil
}

func (__obf_9e636908868342b9 *Captcha) DrawWithNRGBA(__obf_754113aaaca2be09 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_3ca57388aa6752a1 := __obf_754113aaaca2be09.DrawBlocks
	__obf_27eb9dcd4d7779d5 := helper.CreateNRGBACanvas(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, true)

	for _, __obf_890a806d4d4bd969 := range __obf_3ca57388aa6752a1 {
		__obf_32757d85b8868508, __obf_d5b002b58d3eb1d0 := __obf_9e636908868342b9.__obf_e49f106003e8ce29(__obf_890a806d4d4bd969.Width, __obf_890a806d4d4bd969.Height, __obf_890a806d4d4bd969.Shape)
		if __obf_d5b002b58d3eb1d0 != nil {
			return nil, nil, __obf_d5b002b58d3eb1d0
		}
		__obf_394a2b1d2399c804 := __obf_32757d85b8868508.Bounds()
		draw.Draw(__obf_27eb9dcd4d7779d5, image.Rect(__obf_890a806d4d4bd969.X, __obf_890a806d4d4bd969.Y, __obf_890a806d4d4bd969.X+__obf_394a2b1d2399c804.Dx(), __obf_890a806d4d4bd969.Y+__obf_394a2b1d2399c804.Dy()), __obf_32757d85b8868508, image.Point{}, draw.Over)
	}
	__obf_f1df2fdabd84879d := helper.CreateNRGBACanvas(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, true)
	if __obf_754113aaaca2be09.Background != nil {
		__obf_8b35ebf49d1f3dba := __obf_754113aaaca2be09.Background
		__obf_890a806d4d4bd969 := __obf_8b35ebf49d1f3dba.Bounds()
		__obf_e30c37bcfb8ba711 := helper.CreateNRGBACanvas(__obf_890a806d4d4bd969.Dx(), __obf_890a806d4d4bd969.Dy(), true)
		__obf_ef2304aec1631f1e := util.RangCutImagePos(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, __obf_8b35ebf49d1f3dba)
		draw.Draw(__obf_e30c37bcfb8ba711, __obf_890a806d4d4bd969, __obf_8b35ebf49d1f3dba, __obf_ef2304aec1631f1e, draw.Src)
		__obf_e30c37bcfb8ba711.
			SubImage(image.Rect(0, 0, __obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height))

		draw.Draw(__obf_f1df2fdabd84879d, __obf_f1df2fdabd84879d.Bounds(), __obf_e30c37bcfb8ba711, image.Point{}, draw.Over)
		draw.Draw(__obf_e30c37bcfb8ba711, __obf_27eb9dcd4d7779d5.Bounds(), __obf_27eb9dcd4d7779d5, image.Point{}, draw.Over)
		return __obf_e30c37bcfb8ba711, __obf_f1df2fdabd84879d, nil
	}

	return __obf_27eb9dcd4d7779d5, __obf_f1df2fdabd84879d, nil
}

func (__obf_9e636908868342b9 *Captcha) DrawWithTemplate(__obf_754113aaaca2be09 *DrawTplImageParam) (image.Image, error) {
	__obf_73053d5754e12223 := __obf_754113aaaca2be09.Block
	__obf_8b35ebf49d1f3dba := __obf_754113aaaca2be09.Background
	__obf_27eb9dcd4d7779d5 := helper.CreateNRGBACanvas(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, true)
	__obf_2861a25fb8b529ab := helper.CreateNRGBACanvas(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, true)
	__obf_319c3947d03b4363, __obf_d5b002b58d3eb1d0 := __obf_9e636908868342b9.__obf_e49f106003e8ce29(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, __obf_754113aaaca2be09.Mask)
	if __obf_d5b002b58d3eb1d0 != nil {
		return nil, __obf_d5b002b58d3eb1d0
	}

	draw.Draw(__obf_2861a25fb8b529ab, __obf_2861a25fb8b529ab.Bounds(), __obf_8b35ebf49d1f3dba, image.Pt(__obf_73053d5754e12223.X, __obf_73053d5754e12223.Y), draw.Src)
	draw.DrawMask(__obf_27eb9dcd4d7779d5, __obf_319c3947d03b4363.Bounds(), __obf_2861a25fb8b529ab, image.Point{}, __obf_319c3947d03b4363, image.Point{}, draw.Over)
	__obf_af5e2dd5e7d73fcf, __obf_d5b002b58d3eb1d0 := __obf_9e636908868342b9.__obf_e49f106003e8ce29(__obf_754113aaaca2be09.Width, __obf_754113aaaca2be09.Height, __obf_73053d5754e12223.Shape)
	if __obf_d5b002b58d3eb1d0 != nil {
		return nil, __obf_d5b002b58d3eb1d0
	}
	draw.Draw(__obf_27eb9dcd4d7779d5, __obf_af5e2dd5e7d73fcf.Bounds(), __obf_af5e2dd5e7d73fcf, image.Point{}, draw.Over)

	return __obf_27eb9dcd4d7779d5, nil
}

func (__obf_9e636908868342b9 *Captcha) __obf_e49f106003e8ce29(__obf_16ea4bf5cfce0020, __obf_5bd7ea7c03fc8b4e int, __obf_e3bccf51a84246a3 image.Image) (*types.NRGBA, error) {
	__obf_27eb9dcd4d7779d5 := helper.CreateNRGBACanvas(__obf_16ea4bf5cfce0020, __obf_5bd7ea7c03fc8b4e, true)
	draw.BiLinear.Scale(__obf_27eb9dcd4d7779d5, __obf_27eb9dcd4d7779d5.Bounds(), __obf_e3bccf51a84246a3, __obf_e3bccf51a84246a3.Bounds(), draw.Over, nil)
	return __obf_27eb9dcd4d7779d5, nil
}

func (__obf_9e636908868342b9 *Captcha) __obf_c64986d404c5befe(__obf_16ea4bf5cfce0020, __obf_5bd7ea7c03fc8b4e int, __obf_4e1eced9daaeef70 image.Image, __obf_3ca57388aa6752a1 []*types.Block) (image.Image, image.Image, error) {
	var __obf_a35bf79c2729d96b = make([]*types.Block, 0, len(__obf_3ca57388aa6752a1))
	for _, __obf_478e335647ce3491 := range __obf_3ca57388aa6752a1 {
		__obf_478e335647ce3491.
			Shape = __obf_4e1eced9daaeef70
		__obf_a35bf79c2729d96b = append(__obf_a35bf79c2729d96b, __obf_478e335647ce3491)
	}
	__obf_b1c170d65f45d862, __obf_d5b002b58d3eb1d0 := assets.PickBgImage()
	if __obf_d5b002b58d3eb1d0 != nil {
		return nil, nil, __obf_d5b002b58d3eb1d0
	}

	return __obf_9e636908868342b9.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_16ea4bf5cfce0020,
		Height:     __obf_5bd7ea7c03fc8b4e,
		Background: __obf_b1c170d65f45d862,
		Alpha:      __obf_9e636908868342b9.__obf_7828a22179b13d85.Primary.Alpha,
		DrawBlocks: __obf_a35bf79c2729d96b,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_9e636908868342b9 *Captcha) __obf_13a6d2d4d0877604() types.DeadZoneDirectionType {
	__obf_0da7cd602f709c31 := __obf_9e636908868342b9.__obf_7828a22179b13d85.Secondary.DeadZoneDirections
	__obf_4291897e510adc66 := rand.IntN(len(__obf_0da7cd602f709c31))
	if __obf_4291897e510adc66 < 0 {
		return 0
	}

	return __obf_0da7cd602f709c31[__obf_4291897e510adc66]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_9e636908868342b9 *Captcha) __obf_f755db8f55e02a07() int {
	__obf_bf07ec54d0084b67 := __obf_9e636908868342b9.__obf_7828a22179b13d85.Secondary.AnglePosRange
	__obf_4291897e510adc66 := rand.IntN(len(__obf_bf07ec54d0084b67))
	if __obf_4291897e510adc66 < 0 {
		return 0
	}
	__obf_534f4fd70c9dd997 := __obf_bf07ec54d0084b67[__obf_4291897e510adc66]
	return util.RandInt(__obf_534f4fd70c9dd997.Min, __obf_534f4fd70c9dd997.Max)
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
func (__obf_9e636908868342b9 *Captcha) __obf_a09bc2563bda22ac(__obf_16ea4bf5cfce0020, __obf_5bd7ea7c03fc8b4e int, __obf_6944d2509cf21e70 types.Range, __obf_ad99a4b02fbf4179 int) ([]*types.Block, types.Point) {
	var __obf_3ca57388aa6752a1 = make([]*types.Block, 0, __obf_ad99a4b02fbf4179)
	__obf_4e788626a2ee8b51 := // width := imageSize.Width
		// height := imageSize.Height

		__obf_9e636908868342b9.__obf_f755db8f55e02a07()
	__obf_af30fe2f5a49281a := util.RandInt(__obf_6944d2509cf21e70.Min, __obf_6944d2509cf21e70.Max)
	__obf_01a8f545653487b6 := __obf_af30fe2f5a49281a
	__obf_f447b668f073ddb5 := __obf_af30fe2f5a49281a
	__obf_2afaa0ec9a57cdc1 := __obf_9e636908868342b9.__obf_13a6d2d4d0877604()
	__obf_1e2d8deea1555dc9 := __obf_f447b668f073ddb5 / 2
	__obf_c73fd2e5dc7168df := (__obf_16ea4bf5cfce0020 - __obf_f447b668f073ddb5 - 20) / __obf_ad99a4b02fbf4179
	__obf_ae164905bb48d366 := __obf_9e636908868342b9.__obf_5b5452b6fe623482(5, __obf_5bd7ea7c03fc8b4e-__obf_01a8f545653487b6-5, __obf_01a8f545653487b6, __obf_2afaa0ec9a57cdc1)

	for __obf_ee2500bc5bdac157 := range __obf_ad99a4b02fbf4179 {
		var __obf_73053d5754e12223 = &types.Block{}
		__obf_e385351dbd56685d, __obf_888051b4afc7462c := __obf_9e636908868342b9.__obf_94723d121f22f3a0((__obf_ee2500bc5bdac157*__obf_c73fd2e5dc7168df)+__obf_1e2d8deea1555dc9+5, ((__obf_ee2500bc5bdac157+1)*__obf_c73fd2e5dc7168df)-__obf_1e2d8deea1555dc9, __obf_f447b668f073ddb5, __obf_2afaa0ec9a57cdc1)
		__obf_e385351dbd56685d = int(max(float64(__obf_e385351dbd56685d), float64(__obf_1e2d8deea1555dc9+5)))
		__obf_73053d5754e12223.
			X = util.RandInt(__obf_e385351dbd56685d+20, __obf_888051b4afc7462c+20) - __obf_1e2d8deea1555dc9

		if __obf_9e636908868342b9.__obf_7828a22179b13d85.Secondary.EnableVerticalRandom {
			__obf_ae164905bb48d366 = __obf_9e636908868342b9.__obf_5b5452b6fe623482(5, __obf_5bd7ea7c03fc8b4e-__obf_01a8f545653487b6-5, __obf_01a8f545653487b6, __obf_2afaa0ec9a57cdc1)
		}
		__obf_73053d5754e12223.
			Y = __obf_ae164905bb48d366
		__obf_73053d5754e12223.
			Width = __obf_f447b668f073ddb5
		__obf_73053d5754e12223.
			Height = __obf_01a8f545653487b6
		__obf_73053d5754e12223.
			Angle = __obf_4e788626a2ee8b51
		__obf_3ca57388aa6752a1 = append(__obf_3ca57388aa6752a1, __obf_73053d5754e12223)
	}
	__obf_ef2304aec1631f1e := types.Point{}
	if __obf_9e636908868342b9.__obf_7828a22179b13d85.Type == types.MoveSlide {
		__obf_ef2304aec1631f1e.
			X = util.RandInt(5, __obf_1e2d8deea1555dc9)
		__obf_ef2304aec1631f1e.
			Y = __obf_ae164905bb48d366
		return __obf_3ca57388aa6752a1, __obf_ef2304aec1631f1e
	}

	switch __obf_2afaa0ec9a57cdc1 {
	case types.DeadZoneDirectionTypeTop:
		__obf_ef2304aec1631f1e.
			X = util.RandInt(5, __obf_16ea4bf5cfce0020-__obf_f447b668f073ddb5-5)
		__obf_ef2304aec1631f1e.
			Y = 5
	case types.DeadZoneDirectionTypeBottom:
		__obf_ef2304aec1631f1e.
			X = util.RandInt(5, __obf_16ea4bf5cfce0020-__obf_f447b668f073ddb5-5)
		__obf_ef2304aec1631f1e.
			Y = __obf_5bd7ea7c03fc8b4e - __obf_01a8f545653487b6 - 5
	case types.DeadZoneDirectionTypeLeft:
		__obf_ef2304aec1631f1e.
			X = 5
		__obf_ef2304aec1631f1e.
			Y = util.RandInt(5, __obf_5bd7ea7c03fc8b4e-__obf_01a8f545653487b6-5)
	case types.DeadZoneDirectionTypeRight:
		__obf_ef2304aec1631f1e.
			X = __obf_16ea4bf5cfce0020 - __obf_f447b668f073ddb5 - 5
		__obf_ef2304aec1631f1e.
			Y = util.RandInt(5, __obf_5bd7ea7c03fc8b4e-__obf_01a8f545653487b6-5)
	}

	return __obf_3ca57388aa6752a1,

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
		__obf_ef2304aec1631f1e
}

func (__obf_9e636908868342b9 *Captcha) __obf_94723d121f22f3a0(__obf_e385351dbd56685d, __obf_888051b4afc7462c, __obf_9bd6ceb990ec86a4 int, __obf_2afaa0ec9a57cdc1 types.DeadZoneDirectionType) (int, int) {
	if __obf_2afaa0ec9a57cdc1 == types.DeadZoneDirectionTypeLeft {
		__obf_e385351dbd56685d += __obf_9bd6ceb990ec86a4
		__obf_888051b4afc7462c += __obf_9bd6ceb990ec86a4
	}
	return __obf_e385351dbd56685d,

		// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
		// params:
		//   - start: Start Y coordinate
		//   - end: End Y coordinate
		//   - value: types.Block height
		//   - dzdType: Dead zone direction
		//
		// return: Random Y coordinate
		__obf_888051b4afc7462c
}

func (__obf_9e636908868342b9 *Captcha) __obf_5b5452b6fe623482(__obf_e385351dbd56685d, __obf_888051b4afc7462c, __obf_9bd6ceb990ec86a4 int, __obf_2afaa0ec9a57cdc1 types.DeadZoneDirectionType) int {
	switch __obf_2afaa0ec9a57cdc1 {
	case types.DeadZoneDirectionTypeTop:
		__obf_e385351dbd56685d += __obf_9bd6ceb990ec86a4
	case types.DeadZoneDirectionTypeBottom:
		__obf_888051b4afc7462c -= __obf_9bd6ceb990ec86a4
	}
	return util.RandInt(__obf_e385351dbd56685d, __obf_888051b4afc7462c)
}
