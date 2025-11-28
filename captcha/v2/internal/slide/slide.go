package __obf_eadd0fe376b13c04

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
	__obf_2143402c5416f344 *Options
}

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_2143402c5416f344 *Options) *Captcha {
	if __obf_2143402c5416f344 == nil {
		return nil
	}
	return &Captcha{
		__obf_2143402c5416f344: __obf_2143402c5416f344,
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
func (__obf_d0da3002c024252c *Captcha) DrawWithNRGBA(__obf_1c063f702632f045 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_4c02f8484594d441 := __obf_1c063f702632f045.DrawBlocks
	__obf_b228cad327975b6d := helper.CreateNRGBACanvas(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, true)

	for _, __obf_1891e3550eadd75c := range __obf_4c02f8484594d441 {
		__obf_9c359da243c0508f, __obf_85a8cf260a046d48 := __obf_d0da3002c024252c.__obf_4019adee3932076c(__obf_1891e3550eadd75c.Width, __obf_1891e3550eadd75c.Height, __obf_1891e3550eadd75c.Shape)
		if __obf_85a8cf260a046d48 != nil {
			return nil, nil, __obf_85a8cf260a046d48
		}

		__obf_4dca722ada15991d := __obf_9c359da243c0508f.Bounds()
		draw.Draw(__obf_b228cad327975b6d, image.Rect(__obf_1891e3550eadd75c.X, __obf_1891e3550eadd75c.Y, __obf_1891e3550eadd75c.X+__obf_4dca722ada15991d.Dx(), __obf_1891e3550eadd75c.Y+__obf_4dca722ada15991d.Dy()), __obf_9c359da243c0508f, image.Point{}, draw.Over)
	}

	__obf_34552c6081cd94a1 := helper.CreateNRGBACanvas(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, true)
	if __obf_1c063f702632f045.Background != nil {
		__obf_71a9d8e0b4d81510 := __obf_1c063f702632f045.Background
		__obf_1891e3550eadd75c := __obf_71a9d8e0b4d81510.Bounds()
		__obf_87c3c3ca9f5e9860 := helper.CreateNRGBACanvas(__obf_1891e3550eadd75c.Dx(), __obf_1891e3550eadd75c.Dy(), true)
		__obf_f0780e56f0d9d0e0 := util.RangCutImagePos(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, __obf_71a9d8e0b4d81510)
		draw.Draw(__obf_87c3c3ca9f5e9860, __obf_1891e3550eadd75c, __obf_71a9d8e0b4d81510, __obf_f0780e56f0d9d0e0, draw.Src)
		__obf_87c3c3ca9f5e9860.SubImage(image.Rect(0, 0, __obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height))

		draw.Draw(__obf_34552c6081cd94a1, __obf_34552c6081cd94a1.Bounds(), __obf_87c3c3ca9f5e9860, image.Point{}, draw.Over)
		draw.Draw(__obf_87c3c3ca9f5e9860, __obf_b228cad327975b6d.Bounds(), __obf_b228cad327975b6d, image.Point{}, draw.Over)
		return __obf_87c3c3ca9f5e9860, __obf_34552c6081cd94a1, nil
	}

	return __obf_b228cad327975b6d, __obf_34552c6081cd94a1, nil
}

// DrawWithTemplate draws the tile image using a template
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn tile image
//   - error: Error information
func (__obf_d0da3002c024252c *Captcha) DrawWithTemplate(__obf_1c063f702632f045 *DrawTplImageParam) (image.Image, error) {
	__obf_39adca10bea7dfb2 := __obf_1c063f702632f045.Block
	__obf_71a9d8e0b4d81510 := __obf_1c063f702632f045.Background
	__obf_b228cad327975b6d := helper.CreateNRGBACanvas(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, true)
	__obf_a319f2598b16885d := helper.CreateNRGBACanvas(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, true)

	__obf_1b0f5dd530dc7b6b, __obf_85a8cf260a046d48 := __obf_d0da3002c024252c.__obf_4019adee3932076c(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, __obf_1c063f702632f045.Mask)
	if __obf_85a8cf260a046d48 != nil {
		return nil, __obf_85a8cf260a046d48
	}

	draw.Draw(__obf_a319f2598b16885d, __obf_a319f2598b16885d.Bounds(), __obf_71a9d8e0b4d81510, image.Pt(__obf_39adca10bea7dfb2.X, __obf_39adca10bea7dfb2.Y), draw.Src)
	draw.DrawMask(__obf_b228cad327975b6d, __obf_1b0f5dd530dc7b6b.Bounds(), __obf_a319f2598b16885d, image.Point{}, __obf_1b0f5dd530dc7b6b, image.Point{}, draw.Over)

	__obf_bb290526c440ba2f, __obf_85a8cf260a046d48 := __obf_d0da3002c024252c.__obf_4019adee3932076c(__obf_1c063f702632f045.Width, __obf_1c063f702632f045.Height, __obf_39adca10bea7dfb2.Shape)
	if __obf_85a8cf260a046d48 != nil {
		return nil, __obf_85a8cf260a046d48
	}
	draw.Draw(__obf_b228cad327975b6d, __obf_bb290526c440ba2f.Bounds(), __obf_bb290526c440ba2f, image.Point{}, draw.Over)

	return __obf_b228cad327975b6d, nil
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
func (__obf_d0da3002c024252c *Captcha) __obf_4019adee3932076c(__obf_e7ccc17bcbe3de8f, __obf_4d1651cacb60a6a6 int, __obf_415c5f349bac7e10 image.Image) (*types.NRGBA, error) {
	__obf_b228cad327975b6d := helper.CreateNRGBACanvas(__obf_e7ccc17bcbe3de8f, __obf_4d1651cacb60a6a6, true)
	draw.BiLinear.Scale(__obf_b228cad327975b6d, __obf_b228cad327975b6d.Bounds(), __obf_415c5f349bac7e10, __obf_415c5f349bac7e10.Bounds(), draw.Over, nil)
	return __obf_b228cad327975b6d, nil
}

// Generate generates slide CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_d0da3002c024252c *Captcha) Generate() (*CaptchaData, error) {

	// slideType := randomSlideType(c.opts)
	__obf_61d5d22c626f1747, __obf_85a8cf260a046d48 := assets.PickTile()
	if __obf_85a8cf260a046d48 != nil {
		return nil, __obf_85a8cf260a046d48
	}

	__obf_8e2526151d55bc96 := util.RandInt(__obf_d0da3002c024252c.__obf_2143402c5416f344.Secondary.CountRange.Min, __obf_d0da3002c024252c.__obf_2143402c5416f344.Secondary.CountRange.Max)
	__obf_4c02f8484594d441, __obf_6cfd85515a59af87 := __obf_d0da3002c024252c.__obf_0e72984284722590(__obf_d0da3002c024252c.__obf_2143402c5416f344.Primary.Size, __obf_d0da3002c024252c.__obf_2143402c5416f344.Secondary.SizeRange, __obf_8e2526151d55bc96)
	var __obf_39adca10bea7dfb2 *types.Block
	if len(__obf_4c02f8484594d441) > 1 {
		__obf_f8cb0922fda16fde := max(rand.IntN(len(__obf_4c02f8484594d441)), 0)
		__obf_39adca10bea7dfb2 = __obf_4c02f8484594d441[__obf_f8cb0922fda16fde]
	} else {
		__obf_39adca10bea7dfb2 = __obf_4c02f8484594d441[0]
	}

	__obf_41c3b81b590a5d60, __obf_cca8c736007d1848, __obf_85a8cf260a046d48 := __obf_d0da3002c024252c.__obf_01fa68946f885017(__obf_d0da3002c024252c.__obf_2143402c5416f344.Primary.Size, __obf_61d5d22c626f1747.Shadow, __obf_4c02f8484594d441)
	if __obf_85a8cf260a046d48 != nil {
		return nil, __obf_85a8cf260a046d48
	}

	__obf_39adca10bea7dfb2.Shape = __obf_61d5d22c626f1747.Overlay

	__obf_60a3f278519f7fa6, __obf_85a8cf260a046d48 := __obf_d0da3002c024252c.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_cca8c736007d1848,
		Mask:       __obf_61d5d22c626f1747.Mask,
		Alpha:      __obf_d0da3002c024252c.__obf_2143402c5416f344.Primary.Alpha,
		Width:      __obf_39adca10bea7dfb2.Width,
		Height:     __obf_39adca10bea7dfb2.Height,
		Block:      __obf_39adca10bea7dfb2,
	})

	if __obf_85a8cf260a046d48 != nil {
		return nil, __obf_85a8cf260a046d48
	}

	if __obf_d0da3002c024252c.__obf_2143402c5416f344.Type == Move {
		__obf_39adca10bea7dfb2.DY = __obf_39adca10bea7dfb2.Y
	} else {
		__obf_39adca10bea7dfb2.DY = __obf_6cfd85515a59af87.Y
	}
	__obf_39adca10bea7dfb2.DX = __obf_6cfd85515a59af87.X

	return &CaptchaData{
		__obf_e83956f7a258747d: __obf_d0da3002c024252c.__obf_2143402c5416f344.Type,
		__obf_39adca10bea7dfb2: __obf_39adca10bea7dfb2,
		__obf_41c3b81b590a5d60: types.NewJPEGImage(__obf_41c3b81b590a5d60),
		__obf_60a3f278519f7fa6: types.NewPNGImage(__obf_60a3f278519f7fa6),
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
func (__obf_d0da3002c024252c *Captcha) __obf_01fa68946f885017(__obf_1014df2c1bd7d728 types.Size, __obf_4ea8f255cc61eabb image.Image, __obf_4c02f8484594d441 []*types.Block) (image.Image, image.Image, error) {
	var __obf_c1bd584796ba3c01 = make([]*types.Block, 0, len(__obf_4c02f8484594d441))
	for _, __obf_6c61e850039f794d := range __obf_4c02f8484594d441 {
		__obf_6c61e850039f794d.Shape = __obf_4ea8f255cc61eabb
		__obf_c1bd584796ba3c01 = append(__obf_c1bd584796ba3c01, __obf_6c61e850039f794d)
	}

	__obf_302ed1cd98621011, __obf_85a8cf260a046d48 := assets.PickBgImage()
	if __obf_85a8cf260a046d48 != nil {
		return nil, nil, __obf_85a8cf260a046d48
	}

	return __obf_d0da3002c024252c.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_1014df2c1bd7d728.Width,
		Height:     __obf_1014df2c1bd7d728.Height,
		Background: __obf_302ed1cd98621011,
		Alpha:      __obf_d0da3002c024252c.__obf_2143402c5416f344.Primary.Alpha,
		DrawBlocks: __obf_c1bd584796ba3c01,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_d0da3002c024252c *Captcha) __obf_9c03dca74c234212() DeadZoneDirectionType {
	__obf_23bc68a2916664a5 := __obf_d0da3002c024252c.__obf_2143402c5416f344.Secondary.DeadZoneDirections

	__obf_f8cb0922fda16fde := rand.IntN(len(__obf_23bc68a2916664a5))
	if __obf_f8cb0922fda16fde < 0 {
		return 0
	}

	return __obf_23bc68a2916664a5[__obf_f8cb0922fda16fde]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_d0da3002c024252c *Captcha) __obf_9636908124c4df13() int {
	__obf_5b90843209f73aab := __obf_d0da3002c024252c.__obf_2143402c5416f344.Secondary.AnglePosRange

	__obf_f8cb0922fda16fde := rand.IntN(len(__obf_5b90843209f73aab))
	if __obf_f8cb0922fda16fde < 0 {
		return 0
	}

	__obf_c8e198d2d6818a48 := __obf_5b90843209f73aab[__obf_f8cb0922fda16fde]
	return util.RandInt(__obf_c8e198d2d6818a48.Min, __obf_c8e198d2d6818a48.Max)
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
func (__obf_d0da3002c024252c *Captcha) __obf_0e72984284722590(__obf_93e511537f24135c types.Size, __obf_1014df2c1bd7d728 types.Range, __obf_b95dc959817bf014 int) ([]*types.Block, types.Point) {
	var __obf_4c02f8484594d441 = make([]*types.Block, 0, __obf_b95dc959817bf014)
	__obf_e7ccc17bcbe3de8f := __obf_93e511537f24135c.Width
	__obf_4d1651cacb60a6a6 := __obf_93e511537f24135c.Height

	__obf_82b362e926424580 := __obf_d0da3002c024252c.__obf_9636908124c4df13()
	__obf_8448f5d8efb0f799 := util.RandInt(__obf_1014df2c1bd7d728.Min, __obf_1014df2c1bd7d728.Max)
	__obf_36dc0e9eebb853ac := __obf_8448f5d8efb0f799
	__obf_f019f1692c863c03 := __obf_8448f5d8efb0f799

	__obf_d14feda271e24222 := __obf_d0da3002c024252c.__obf_9c03dca74c234212()
	__obf_0aa9dc10107889e1 := __obf_f019f1692c863c03 / 2
	__obf_5b2644321fd75462 := (__obf_e7ccc17bcbe3de8f - __obf_f019f1692c863c03 - 20) / __obf_b95dc959817bf014
	__obf_bcc22c54528fbf10 := __obf_d0da3002c024252c.__obf_90d9b2a68027961e(5, __obf_4d1651cacb60a6a6-__obf_36dc0e9eebb853ac-5, __obf_36dc0e9eebb853ac, __obf_d14feda271e24222)

	for __obf_6a9630456a34df8e := range __obf_b95dc959817bf014 {
		var __obf_39adca10bea7dfb2 = &types.Block{}
		__obf_8c13bef1705d8365, __obf_9ecdd1975f39bff6 := __obf_d0da3002c024252c.__obf_a8df72505ce07e7c((__obf_6a9630456a34df8e*__obf_5b2644321fd75462)+__obf_0aa9dc10107889e1+5, ((__obf_6a9630456a34df8e+1)*__obf_5b2644321fd75462)-__obf_0aa9dc10107889e1, __obf_f019f1692c863c03, __obf_d14feda271e24222)

		__obf_8c13bef1705d8365 = int(max(float64(__obf_8c13bef1705d8365), float64(__obf_0aa9dc10107889e1+5)))
		__obf_39adca10bea7dfb2.X = util.RandInt(__obf_8c13bef1705d8365+20, __obf_9ecdd1975f39bff6+20) - __obf_0aa9dc10107889e1

		if __obf_d0da3002c024252c.__obf_2143402c5416f344.Secondary.EnableVerticalRandom {
			__obf_bcc22c54528fbf10 = __obf_d0da3002c024252c.__obf_90d9b2a68027961e(5, __obf_4d1651cacb60a6a6-__obf_36dc0e9eebb853ac-5, __obf_36dc0e9eebb853ac, __obf_d14feda271e24222)
		}

		__obf_39adca10bea7dfb2.Y = __obf_bcc22c54528fbf10
		__obf_39adca10bea7dfb2.Width = __obf_f019f1692c863c03
		__obf_39adca10bea7dfb2.Height = __obf_36dc0e9eebb853ac
		__obf_39adca10bea7dfb2.Angle = __obf_82b362e926424580

		__obf_4c02f8484594d441 = append(__obf_4c02f8484594d441, __obf_39adca10bea7dfb2)
	}

	__obf_f0780e56f0d9d0e0 := types.Point{}
	if __obf_d0da3002c024252c.__obf_2143402c5416f344.Type == Move {
		__obf_f0780e56f0d9d0e0.X = util.RandInt(5, __obf_0aa9dc10107889e1)
		__obf_f0780e56f0d9d0e0.Y = __obf_bcc22c54528fbf10
		return __obf_4c02f8484594d441, __obf_f0780e56f0d9d0e0
	}

	switch __obf_d14feda271e24222 {
	case DeadZoneDirectionTypeTop:
		__obf_f0780e56f0d9d0e0.X = util.RandInt(5, __obf_e7ccc17bcbe3de8f-__obf_f019f1692c863c03-5)
		__obf_f0780e56f0d9d0e0.Y = 5
	case DeadZoneDirectionTypeBottom:
		__obf_f0780e56f0d9d0e0.X = util.RandInt(5, __obf_e7ccc17bcbe3de8f-__obf_f019f1692c863c03-5)
		__obf_f0780e56f0d9d0e0.Y = __obf_4d1651cacb60a6a6 - __obf_36dc0e9eebb853ac - 5
	case DeadZoneDirectionTypeLeft:
		__obf_f0780e56f0d9d0e0.X = 5
		__obf_f0780e56f0d9d0e0.Y = util.RandInt(5, __obf_4d1651cacb60a6a6-__obf_36dc0e9eebb853ac-5)
	case DeadZoneDirectionTypeRight:
		__obf_f0780e56f0d9d0e0.X = __obf_e7ccc17bcbe3de8f - __obf_f019f1692c863c03 - 5
		__obf_f0780e56f0d9d0e0.Y = util.RandInt(5, __obf_4d1651cacb60a6a6-__obf_36dc0e9eebb853ac-5)
	}

	return __obf_4c02f8484594d441, __obf_f0780e56f0d9d0e0
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
func (__obf_d0da3002c024252c *Captcha) __obf_a8df72505ce07e7c(__obf_8c13bef1705d8365, __obf_9ecdd1975f39bff6, __obf_38c7704fcc75f839 int, __obf_d14feda271e24222 DeadZoneDirectionType) (int, int) {
	if __obf_d14feda271e24222 == DeadZoneDirectionTypeLeft {
		__obf_8c13bef1705d8365 += __obf_38c7704fcc75f839
		__obf_9ecdd1975f39bff6 += __obf_38c7704fcc75f839
	}
	return __obf_8c13bef1705d8365, __obf_9ecdd1975f39bff6
}

// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
// params:
//   - start: Start Y coordinate
//   - end: End Y coordinate
//   - value: types.Block height
//   - dzdType: Dead zone direction
//
// return: Random Y coordinate
func (__obf_d0da3002c024252c *Captcha) __obf_90d9b2a68027961e(__obf_8c13bef1705d8365, __obf_9ecdd1975f39bff6, __obf_38c7704fcc75f839 int, __obf_d14feda271e24222 DeadZoneDirectionType) int {
	switch __obf_d14feda271e24222 {
	case DeadZoneDirectionTypeTop:
		__obf_8c13bef1705d8365 += __obf_38c7704fcc75f839
	case DeadZoneDirectionTypeBottom:
		__obf_9ecdd1975f39bff6 -= __obf_38c7704fcc75f839
	}
	return util.RandInt(__obf_8c13bef1705d8365, __obf_9ecdd1975f39bff6)
}
