package __obf_45140da9e3f6647e

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
	__obf_ca97d616fb24a1ba *Options
}

var _ types.Captcha = (*Captcha)(nil)

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_ca97d616fb24a1ba *Options) *Captcha {
	if __obf_ca97d616fb24a1ba == nil {
		return nil
	}
	return &Captcha{__obf_ca97d616fb24a1ba: __obf_ca97d616fb24a1ba}
}

// func (c *Captcha) Type() types.CaptchaType {
// 	if c.opts.Type == Move {
// 		return types.MoveCaptchat
// 	}
// 	return types.DragCaptchat
// }

func (__obf_012c5ff8d6b175b3 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_1cd7ae54a005ff5a,

		// slideType := randomSlideType(c.opts)
		__obf_7d580363a4d8cba8 := assets.PickTile()
	if __obf_7d580363a4d8cba8 != nil {
		return nil, __obf_7d580363a4d8cba8
	}
	__obf_fa8c209cec13f788 := util.RandInt(__obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Secondary.CountRange.Min, __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Secondary.CountRange.Max)
	__obf_0bf7208306d585e2, __obf_a1d9397dddfa2d15 := __obf_012c5ff8d6b175b3.__obf_bc3ab9db0b6cac2b(__obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Primary.Size, __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Secondary.SizeRange, __obf_fa8c209cec13f788)
	var __obf_3c2937c851cafbfc *types.Block
	if len(__obf_0bf7208306d585e2) > 1 {
		__obf_73ab4fa7cc5d5709 := max(rand.IntN(len(__obf_0bf7208306d585e2)), 0)
		__obf_3c2937c851cafbfc = __obf_0bf7208306d585e2[__obf_73ab4fa7cc5d5709]
	} else {
		__obf_3c2937c851cafbfc = __obf_0bf7208306d585e2[0]
	}
	__obf_8fcc6dfc88c19e7c, __obf_ae591bb0688bf8ff, __obf_7d580363a4d8cba8 := __obf_012c5ff8d6b175b3.__obf_064bd781c5456c53(__obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Primary.Size, __obf_1cd7ae54a005ff5a.Shadow, __obf_0bf7208306d585e2)
	if __obf_7d580363a4d8cba8 != nil {
		return nil, __obf_7d580363a4d8cba8
	}
	__obf_3c2937c851cafbfc.
		Shape = __obf_1cd7ae54a005ff5a.Overlay
	__obf_a99609c2d6fcf851, __obf_7d580363a4d8cba8 := __obf_012c5ff8d6b175b3.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_ae591bb0688bf8ff,
		Mask:       __obf_1cd7ae54a005ff5a.Mask,
		Alpha:      __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Primary.Alpha,
		Width:      __obf_3c2937c851cafbfc.Width,
		Height:     __obf_3c2937c851cafbfc.Height,
		Block:      __obf_3c2937c851cafbfc,
	})

	if __obf_7d580363a4d8cba8 != nil {
		return nil, __obf_7d580363a4d8cba8
	}

	if __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Type == Move {
		__obf_3c2937c851cafbfc.
			DY = __obf_3c2937c851cafbfc.Y
	} else {
		__obf_3c2937c851cafbfc.
			DY = __obf_a1d9397dddfa2d15.Y
	}
	__obf_3c2937c851cafbfc.
		DX = __obf_a1d9397dddfa2d15.X

	return &CaptchaData{__obf_f05be736ef82a7b3: __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Type, __obf_3c2937c851cafbfc: __obf_3c2937c851cafbfc, __obf_8fcc6dfc88c19e7c: types.NewJPEGImage(__obf_8fcc6dfc88c19e7c), __obf_a99609c2d6fcf851: types.NewPNGImage(__obf_a99609c2d6fcf851)}, nil
}

func (__obf_012c5ff8d6b175b3 *Captcha) DrawWithNRGBA(__obf_30604d6baa6da710 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_0bf7208306d585e2 := __obf_30604d6baa6da710.DrawBlocks
	__obf_d0f16e63db672115 := helper.CreateNRGBACanvas(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, true)

	for _, __obf_ea09cefce012b03c := range __obf_0bf7208306d585e2 {
		__obf_4043a5ead79aaac8, __obf_7d580363a4d8cba8 := __obf_012c5ff8d6b175b3.__obf_58cd966121e85902(__obf_ea09cefce012b03c.Width, __obf_ea09cefce012b03c.Height, __obf_ea09cefce012b03c.Shape)
		if __obf_7d580363a4d8cba8 != nil {
			return nil, nil, __obf_7d580363a4d8cba8
		}
		__obf_4d1b7ba3e8b2b8c0 := __obf_4043a5ead79aaac8.Bounds()
		draw.Draw(__obf_d0f16e63db672115, image.Rect(__obf_ea09cefce012b03c.X, __obf_ea09cefce012b03c.Y, __obf_ea09cefce012b03c.X+__obf_4d1b7ba3e8b2b8c0.Dx(), __obf_ea09cefce012b03c.Y+__obf_4d1b7ba3e8b2b8c0.Dy()), __obf_4043a5ead79aaac8, image.Point{}, draw.Over)
	}
	__obf_cc1cce4206b5ba84 := helper.CreateNRGBACanvas(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, true)
	if __obf_30604d6baa6da710.Background != nil {
		__obf_641236ae9938b8d9 := __obf_30604d6baa6da710.Background
		__obf_ea09cefce012b03c := __obf_641236ae9938b8d9.Bounds()
		__obf_dd35c75fb1f73ce1 := helper.CreateNRGBACanvas(__obf_ea09cefce012b03c.Dx(), __obf_ea09cefce012b03c.Dy(), true)
		__obf_6d7ca0f885852bd7 := util.RangCutImagePos(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, __obf_641236ae9938b8d9)
		draw.Draw(__obf_dd35c75fb1f73ce1, __obf_ea09cefce012b03c, __obf_641236ae9938b8d9, __obf_6d7ca0f885852bd7, draw.Src)
		__obf_dd35c75fb1f73ce1.
			SubImage(image.Rect(0, 0, __obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height))

		draw.Draw(__obf_cc1cce4206b5ba84, __obf_cc1cce4206b5ba84.Bounds(), __obf_dd35c75fb1f73ce1, image.Point{}, draw.Over)
		draw.Draw(__obf_dd35c75fb1f73ce1, __obf_d0f16e63db672115.Bounds(), __obf_d0f16e63db672115, image.Point{}, draw.Over)
		return __obf_dd35c75fb1f73ce1, __obf_cc1cce4206b5ba84, nil
	}

	return __obf_d0f16e63db672115, __obf_cc1cce4206b5ba84, nil
}

func (__obf_012c5ff8d6b175b3 *Captcha) DrawWithTemplate(__obf_30604d6baa6da710 *DrawTplImageParam) (image.Image, error) {
	__obf_3c2937c851cafbfc := __obf_30604d6baa6da710.Block
	__obf_641236ae9938b8d9 := __obf_30604d6baa6da710.Background
	__obf_d0f16e63db672115 := helper.CreateNRGBACanvas(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, true)
	__obf_145a72ad2a408141 := helper.CreateNRGBACanvas(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, true)
	__obf_be2eb1b3221b45e9, __obf_7d580363a4d8cba8 := __obf_012c5ff8d6b175b3.__obf_58cd966121e85902(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, __obf_30604d6baa6da710.Mask)
	if __obf_7d580363a4d8cba8 != nil {
		return nil, __obf_7d580363a4d8cba8
	}

	draw.Draw(__obf_145a72ad2a408141, __obf_145a72ad2a408141.Bounds(), __obf_641236ae9938b8d9, image.Pt(__obf_3c2937c851cafbfc.X, __obf_3c2937c851cafbfc.Y), draw.Src)
	draw.DrawMask(__obf_d0f16e63db672115, __obf_be2eb1b3221b45e9.Bounds(), __obf_145a72ad2a408141, image.Point{}, __obf_be2eb1b3221b45e9, image.Point{}, draw.Over)
	__obf_5ad4346860dac16a, __obf_7d580363a4d8cba8 := __obf_012c5ff8d6b175b3.__obf_58cd966121e85902(__obf_30604d6baa6da710.Width, __obf_30604d6baa6da710.Height, __obf_3c2937c851cafbfc.Shape)
	if __obf_7d580363a4d8cba8 != nil {
		return nil, __obf_7d580363a4d8cba8
	}
	draw.Draw(__obf_d0f16e63db672115, __obf_5ad4346860dac16a.Bounds(), __obf_5ad4346860dac16a, image.Point{}, draw.Over)

	return __obf_d0f16e63db672115, nil
}

func (__obf_012c5ff8d6b175b3 *Captcha) __obf_58cd966121e85902(__obf_60c492e3ed009962, __obf_53fb2f5b5c65aa9b int, __obf_1a59472774e4192a image.Image) (*types.NRGBA, error) {
	__obf_d0f16e63db672115 := helper.CreateNRGBACanvas(__obf_60c492e3ed009962, __obf_53fb2f5b5c65aa9b, true)
	draw.BiLinear.Scale(__obf_d0f16e63db672115, __obf_d0f16e63db672115.Bounds(), __obf_1a59472774e4192a, __obf_1a59472774e4192a.Bounds(), draw.Over, nil)
	return __obf_d0f16e63db672115, nil
}

func (__obf_012c5ff8d6b175b3 *Captcha) __obf_064bd781c5456c53(__obf_4f0a696b10c010cf types.Size, __obf_29ab119bf927059d image.Image, __obf_0bf7208306d585e2 []*types.Block) (image.Image, image.Image, error) {
	var __obf_59a6f2af54ba2291 = make([]*types.Block, 0, len(__obf_0bf7208306d585e2))
	for _, __obf_d3dd05ee587c53c9 := range __obf_0bf7208306d585e2 {
		__obf_d3dd05ee587c53c9.
			Shape = __obf_29ab119bf927059d
		__obf_59a6f2af54ba2291 = append(__obf_59a6f2af54ba2291, __obf_d3dd05ee587c53c9)
	}
	__obf_f711b2e3efa3257f, __obf_7d580363a4d8cba8 := assets.PickBgImage()
	if __obf_7d580363a4d8cba8 != nil {
		return nil, nil, __obf_7d580363a4d8cba8
	}

	return __obf_012c5ff8d6b175b3.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_4f0a696b10c010cf.Width,
		Height:     __obf_4f0a696b10c010cf.Height,
		Background: __obf_f711b2e3efa3257f,
		Alpha:      __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Primary.Alpha,
		DrawBlocks: __obf_59a6f2af54ba2291,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_012c5ff8d6b175b3 *Captcha) __obf_3e21558a4adc455a() DeadZoneDirectionType {
	__obf_ba8d81ee620db099 := __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Secondary.DeadZoneDirections
	__obf_73ab4fa7cc5d5709 := rand.IntN(len(__obf_ba8d81ee620db099))
	if __obf_73ab4fa7cc5d5709 < 0 {
		return 0
	}

	return __obf_ba8d81ee620db099[__obf_73ab4fa7cc5d5709]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_012c5ff8d6b175b3 *Captcha) __obf_1a43d9f375b1fe2c() int {
	__obf_53d81001afe84414 := __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Secondary.AnglePosRange
	__obf_73ab4fa7cc5d5709 := rand.IntN(len(__obf_53d81001afe84414))
	if __obf_73ab4fa7cc5d5709 < 0 {
		return 0
	}
	__obf_506f89269f462799 := __obf_53d81001afe84414[__obf_73ab4fa7cc5d5709]
	return util.RandInt(__obf_506f89269f462799.Min, __obf_506f89269f462799.Max)
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
func (__obf_012c5ff8d6b175b3 *Captcha) __obf_bc3ab9db0b6cac2b(__obf_1c28df5ce47fca4e types.Size, __obf_4f0a696b10c010cf types.Range, __obf_d9274f58c27835a9 int) ([]*types.Block, types.Point) {
	var __obf_0bf7208306d585e2 = make([]*types.Block, 0, __obf_d9274f58c27835a9)
	__obf_60c492e3ed009962 := __obf_1c28df5ce47fca4e.Width
	__obf_53fb2f5b5c65aa9b := __obf_1c28df5ce47fca4e.Height
	__obf_f70e7eb146183bda := __obf_012c5ff8d6b175b3.__obf_1a43d9f375b1fe2c()
	__obf_4ea536b4ccb1c135 := util.RandInt(__obf_4f0a696b10c010cf.Min, __obf_4f0a696b10c010cf.Max)
	__obf_5fad2568c7578b98 := __obf_4ea536b4ccb1c135
	__obf_207882fb6cc9e660 := __obf_4ea536b4ccb1c135
	__obf_c892f7b24842eef7 := __obf_012c5ff8d6b175b3.__obf_3e21558a4adc455a()
	__obf_b199f8bb9e9c131e := __obf_207882fb6cc9e660 / 2
	__obf_eafc8e48441a1650 := (__obf_60c492e3ed009962 - __obf_207882fb6cc9e660 - 20) / __obf_d9274f58c27835a9
	__obf_d3ee19a6518ffc6f := __obf_012c5ff8d6b175b3.__obf_310ae187fcfecfd0(5, __obf_53fb2f5b5c65aa9b-__obf_5fad2568c7578b98-5, __obf_5fad2568c7578b98, __obf_c892f7b24842eef7)

	for __obf_df1dff84289eec72 := range __obf_d9274f58c27835a9 {
		var __obf_3c2937c851cafbfc = &types.Block{}
		__obf_5c921d99474e394d, __obf_e4cd096715c137fc := __obf_012c5ff8d6b175b3.__obf_bde1690f5a5fd226((__obf_df1dff84289eec72*__obf_eafc8e48441a1650)+__obf_b199f8bb9e9c131e+5, ((__obf_df1dff84289eec72+1)*__obf_eafc8e48441a1650)-__obf_b199f8bb9e9c131e, __obf_207882fb6cc9e660, __obf_c892f7b24842eef7)
		__obf_5c921d99474e394d = int(max(float64(__obf_5c921d99474e394d), float64(__obf_b199f8bb9e9c131e+5)))
		__obf_3c2937c851cafbfc.
			X = util.RandInt(__obf_5c921d99474e394d+20, __obf_e4cd096715c137fc+20) - __obf_b199f8bb9e9c131e

		if __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Secondary.EnableVerticalRandom {
			__obf_d3ee19a6518ffc6f = __obf_012c5ff8d6b175b3.__obf_310ae187fcfecfd0(5, __obf_53fb2f5b5c65aa9b-__obf_5fad2568c7578b98-5, __obf_5fad2568c7578b98, __obf_c892f7b24842eef7)
		}
		__obf_3c2937c851cafbfc.
			Y = __obf_d3ee19a6518ffc6f
		__obf_3c2937c851cafbfc.
			Width = __obf_207882fb6cc9e660
		__obf_3c2937c851cafbfc.
			Height = __obf_5fad2568c7578b98
		__obf_3c2937c851cafbfc.
			Angle = __obf_f70e7eb146183bda
		__obf_0bf7208306d585e2 = append(__obf_0bf7208306d585e2, __obf_3c2937c851cafbfc)
	}
	__obf_6d7ca0f885852bd7 := types.Point{}
	if __obf_012c5ff8d6b175b3.__obf_ca97d616fb24a1ba.Type == Move {
		__obf_6d7ca0f885852bd7.
			X = util.RandInt(5, __obf_b199f8bb9e9c131e)
		__obf_6d7ca0f885852bd7.
			Y = __obf_d3ee19a6518ffc6f
		return __obf_0bf7208306d585e2, __obf_6d7ca0f885852bd7
	}

	switch __obf_c892f7b24842eef7 {
	case DeadZoneDirectionTypeTop:
		__obf_6d7ca0f885852bd7.
			X = util.RandInt(5, __obf_60c492e3ed009962-__obf_207882fb6cc9e660-5)
		__obf_6d7ca0f885852bd7.
			Y = 5
	case DeadZoneDirectionTypeBottom:
		__obf_6d7ca0f885852bd7.
			X = util.RandInt(5, __obf_60c492e3ed009962-__obf_207882fb6cc9e660-5)
		__obf_6d7ca0f885852bd7.
			Y = __obf_53fb2f5b5c65aa9b - __obf_5fad2568c7578b98 - 5
	case DeadZoneDirectionTypeLeft:
		__obf_6d7ca0f885852bd7.
			X = 5
		__obf_6d7ca0f885852bd7.
			Y = util.RandInt(5, __obf_53fb2f5b5c65aa9b-__obf_5fad2568c7578b98-5)
	case DeadZoneDirectionTypeRight:
		__obf_6d7ca0f885852bd7.
			X = __obf_60c492e3ed009962 - __obf_207882fb6cc9e660 - 5
		__obf_6d7ca0f885852bd7.
			Y = util.RandInt(5, __obf_53fb2f5b5c65aa9b-__obf_5fad2568c7578b98-5)
	}

	return __obf_0bf7208306d585e2,

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
		__obf_6d7ca0f885852bd7
}

func (__obf_012c5ff8d6b175b3 *Captcha) __obf_bde1690f5a5fd226(__obf_5c921d99474e394d, __obf_e4cd096715c137fc, __obf_0ffd2f16356a5178 int, __obf_c892f7b24842eef7 DeadZoneDirectionType) (int, int) {
	if __obf_c892f7b24842eef7 == DeadZoneDirectionTypeLeft {
		__obf_5c921d99474e394d += __obf_0ffd2f16356a5178
		__obf_e4cd096715c137fc += __obf_0ffd2f16356a5178
	}
	return __obf_5c921d99474e394d,

		// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
		// params:
		//   - start: Start Y coordinate
		//   - end: End Y coordinate
		//   - value: types.Block height
		//   - dzdType: Dead zone direction
		//
		// return: Random Y coordinate
		__obf_e4cd096715c137fc
}

func (__obf_012c5ff8d6b175b3 *Captcha) __obf_310ae187fcfecfd0(__obf_5c921d99474e394d, __obf_e4cd096715c137fc, __obf_0ffd2f16356a5178 int, __obf_c892f7b24842eef7 DeadZoneDirectionType) int {
	switch __obf_c892f7b24842eef7 {
	case DeadZoneDirectionTypeTop:
		__obf_5c921d99474e394d += __obf_0ffd2f16356a5178
	case DeadZoneDirectionTypeBottom:
		__obf_e4cd096715c137fc -= __obf_0ffd2f16356a5178
	}
	return util.RandInt(__obf_5c921d99474e394d, __obf_e4cd096715c137fc)
}
