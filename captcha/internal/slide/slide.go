package __obf_b1febe121b2dfc2d

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
	__obf_195f5263d9c40fa7 *Options
}

var _ types.Captcha = (*Captcha)(nil)

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_195f5263d9c40fa7 *Options) *Captcha {
	if __obf_195f5263d9c40fa7 == nil {
		return nil
	}
	return &Captcha{__obf_195f5263d9c40fa7: __obf_195f5263d9c40fa7}
}

func (__obf_6c40c4c2143a89cf *Captcha) Generate() (types.CaptchaData, error) {
	__obf_bfc5712f7defe7ee,

		// slideType := randomSlideType(c.opts)
		__obf_5aebd7c35b6603fd := assets.PickTile()
	if __obf_5aebd7c35b6603fd != nil {
		return nil, __obf_5aebd7c35b6603fd
	}
	__obf_7a30ad9a5e348667 := util.RandInt(__obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Secondary.CountRange.Min, __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Secondary.CountRange.Max)
	__obf_1edab5b16168422c, __obf_25680b49bee7b8cd := __obf_6c40c4c2143a89cf.__obf_c6431c3f8222f8a9(__obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Primary.Size, __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Secondary.SizeRange, __obf_7a30ad9a5e348667)
	var __obf_fc1b3b29182d412e *types.Block
	if len(__obf_1edab5b16168422c) > 1 {
		__obf_d0c55d7e94de4ab3 := max(rand.IntN(len(__obf_1edab5b16168422c)), 0)
		__obf_fc1b3b29182d412e = __obf_1edab5b16168422c[__obf_d0c55d7e94de4ab3]
	} else {
		__obf_fc1b3b29182d412e = __obf_1edab5b16168422c[0]
	}
	__obf_ece3f1a292371bc7, __obf_ce29229ace617897, __obf_5aebd7c35b6603fd := __obf_6c40c4c2143a89cf.__obf_1a6614a8b01291c4(__obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Primary.Size, __obf_bfc5712f7defe7ee.Shadow, __obf_1edab5b16168422c)
	if __obf_5aebd7c35b6603fd != nil {
		return nil, __obf_5aebd7c35b6603fd
	}
	__obf_fc1b3b29182d412e.
		Shape = __obf_bfc5712f7defe7ee.Overlay
	__obf_765e5abee3631ca2, __obf_5aebd7c35b6603fd := __obf_6c40c4c2143a89cf.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_ce29229ace617897,
		Mask:       __obf_bfc5712f7defe7ee.Mask,
		Alpha:      __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Primary.Alpha,
		Width:      __obf_fc1b3b29182d412e.Width,
		Height:     __obf_fc1b3b29182d412e.Height,
		Block:      __obf_fc1b3b29182d412e,
	})

	if __obf_5aebd7c35b6603fd != nil {
		return nil, __obf_5aebd7c35b6603fd
	}

	if __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Type == Move {
		__obf_fc1b3b29182d412e.
			DY = __obf_fc1b3b29182d412e.Y
	} else {
		__obf_fc1b3b29182d412e.
			DY = __obf_25680b49bee7b8cd.Y
	}
	__obf_fc1b3b29182d412e.
		DX = __obf_25680b49bee7b8cd.X

	return &CaptchaData{__obf_b78c570ef70ccb05: __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Type, __obf_fc1b3b29182d412e: __obf_fc1b3b29182d412e, __obf_ece3f1a292371bc7: types.NewJPEGImage(__obf_ece3f1a292371bc7), __obf_765e5abee3631ca2: types.NewPNGImage(__obf_765e5abee3631ca2)}, nil
}

func (__obf_6c40c4c2143a89cf *Captcha) DrawWithNRGBA(__obf_71f255c3ea29422c *DrawImageParam) (image.Image, image.Image, error) {
	__obf_1edab5b16168422c := __obf_71f255c3ea29422c.DrawBlocks
	__obf_7b3584bd111a3fe6 := helper.CreateNRGBACanvas(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, true)

	for _, __obf_7c35927d0adde051 := range __obf_1edab5b16168422c {
		__obf_ceb1794341487dd6, __obf_5aebd7c35b6603fd := __obf_6c40c4c2143a89cf.__obf_bd6aaa08c23860f2(__obf_7c35927d0adde051.Width, __obf_7c35927d0adde051.Height, __obf_7c35927d0adde051.Shape)
		if __obf_5aebd7c35b6603fd != nil {
			return nil, nil, __obf_5aebd7c35b6603fd
		}
		__obf_a2fbd4af53d4c076 := __obf_ceb1794341487dd6.Bounds()
		draw.Draw(__obf_7b3584bd111a3fe6, image.Rect(__obf_7c35927d0adde051.X, __obf_7c35927d0adde051.Y, __obf_7c35927d0adde051.X+__obf_a2fbd4af53d4c076.Dx(), __obf_7c35927d0adde051.Y+__obf_a2fbd4af53d4c076.Dy()), __obf_ceb1794341487dd6, image.Point{}, draw.Over)
	}
	__obf_c2ba679c4cff2eeb := helper.CreateNRGBACanvas(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, true)
	if __obf_71f255c3ea29422c.Background != nil {
		__obf_1bbb3f6179a38655 := __obf_71f255c3ea29422c.Background
		__obf_7c35927d0adde051 := __obf_1bbb3f6179a38655.Bounds()
		__obf_d05e89d7f2a100f4 := helper.CreateNRGBACanvas(__obf_7c35927d0adde051.Dx(), __obf_7c35927d0adde051.Dy(), true)
		__obf_7c26d84cefb213cc := util.RangCutImagePos(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, __obf_1bbb3f6179a38655)
		draw.Draw(__obf_d05e89d7f2a100f4, __obf_7c35927d0adde051, __obf_1bbb3f6179a38655, __obf_7c26d84cefb213cc, draw.Src)
		__obf_d05e89d7f2a100f4.
			SubImage(image.Rect(0, 0, __obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height))

		draw.Draw(__obf_c2ba679c4cff2eeb, __obf_c2ba679c4cff2eeb.Bounds(), __obf_d05e89d7f2a100f4, image.Point{}, draw.Over)
		draw.Draw(__obf_d05e89d7f2a100f4, __obf_7b3584bd111a3fe6.Bounds(), __obf_7b3584bd111a3fe6, image.Point{}, draw.Over)
		return __obf_d05e89d7f2a100f4, __obf_c2ba679c4cff2eeb, nil
	}

	return __obf_7b3584bd111a3fe6, __obf_c2ba679c4cff2eeb, nil
}

func (__obf_6c40c4c2143a89cf *Captcha) DrawWithTemplate(__obf_71f255c3ea29422c *DrawTplImageParam) (image.Image, error) {
	__obf_fc1b3b29182d412e := __obf_71f255c3ea29422c.Block
	__obf_1bbb3f6179a38655 := __obf_71f255c3ea29422c.Background
	__obf_7b3584bd111a3fe6 := helper.CreateNRGBACanvas(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, true)
	__obf_9721af8b1b9a5873 := helper.CreateNRGBACanvas(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, true)
	__obf_e07647d6ed81856e, __obf_5aebd7c35b6603fd := __obf_6c40c4c2143a89cf.__obf_bd6aaa08c23860f2(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, __obf_71f255c3ea29422c.Mask)
	if __obf_5aebd7c35b6603fd != nil {
		return nil, __obf_5aebd7c35b6603fd
	}

	draw.Draw(__obf_9721af8b1b9a5873, __obf_9721af8b1b9a5873.Bounds(), __obf_1bbb3f6179a38655, image.Pt(__obf_fc1b3b29182d412e.X, __obf_fc1b3b29182d412e.Y), draw.Src)
	draw.DrawMask(__obf_7b3584bd111a3fe6, __obf_e07647d6ed81856e.Bounds(), __obf_9721af8b1b9a5873, image.Point{}, __obf_e07647d6ed81856e, image.Point{}, draw.Over)
	__obf_f4163ce5cad7788b, __obf_5aebd7c35b6603fd := __obf_6c40c4c2143a89cf.__obf_bd6aaa08c23860f2(__obf_71f255c3ea29422c.Width, __obf_71f255c3ea29422c.Height, __obf_fc1b3b29182d412e.Shape)
	if __obf_5aebd7c35b6603fd != nil {
		return nil, __obf_5aebd7c35b6603fd
	}
	draw.Draw(__obf_7b3584bd111a3fe6, __obf_f4163ce5cad7788b.Bounds(), __obf_f4163ce5cad7788b, image.Point{}, draw.Over)

	return __obf_7b3584bd111a3fe6, nil
}

func (__obf_6c40c4c2143a89cf *Captcha) __obf_bd6aaa08c23860f2(__obf_b5007f75c87ca886, __obf_f39217514bf96315 int, __obf_20f34af22b1cb93e image.Image) (*types.NRGBA, error) {
	__obf_7b3584bd111a3fe6 := helper.CreateNRGBACanvas(__obf_b5007f75c87ca886, __obf_f39217514bf96315, true)
	draw.BiLinear.Scale(__obf_7b3584bd111a3fe6, __obf_7b3584bd111a3fe6.Bounds(), __obf_20f34af22b1cb93e, __obf_20f34af22b1cb93e.Bounds(), draw.Over, nil)
	return __obf_7b3584bd111a3fe6, nil
}

func (__obf_6c40c4c2143a89cf *Captcha) __obf_1a6614a8b01291c4(__obf_067c37c603bffc8d types.Size, __obf_53c8d07bb7f86fd7 image.Image, __obf_1edab5b16168422c []*types.Block) (image.Image, image.Image, error) {
	var __obf_f4cc6f28902600c6 = make([]*types.Block, 0, len(__obf_1edab5b16168422c))
	for _, __obf_3a08cc16aa144792 := range __obf_1edab5b16168422c {
		__obf_3a08cc16aa144792.
			Shape = __obf_53c8d07bb7f86fd7
		__obf_f4cc6f28902600c6 = append(__obf_f4cc6f28902600c6, __obf_3a08cc16aa144792)
	}
	__obf_2652eded0aaeff0b, __obf_5aebd7c35b6603fd := assets.PickBgImage()
	if __obf_5aebd7c35b6603fd != nil {
		return nil, nil, __obf_5aebd7c35b6603fd
	}

	return __obf_6c40c4c2143a89cf.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_067c37c603bffc8d.Width,
		Height:     __obf_067c37c603bffc8d.Height,
		Background: __obf_2652eded0aaeff0b,
		Alpha:      __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Primary.Alpha,
		DrawBlocks: __obf_f4cc6f28902600c6,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_6c40c4c2143a89cf *Captcha) __obf_9e8e6adb3493bd16() DeadZoneDirectionType {
	__obf_e5865ac8cb20d95d := __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Secondary.DeadZoneDirections
	__obf_d0c55d7e94de4ab3 := rand.IntN(len(__obf_e5865ac8cb20d95d))
	if __obf_d0c55d7e94de4ab3 < 0 {
		return 0
	}

	return __obf_e5865ac8cb20d95d[__obf_d0c55d7e94de4ab3]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_6c40c4c2143a89cf *Captcha) __obf_1e2b8a0c7d68a9f4() int {
	__obf_13b6cc0d1d2c6ef9 := __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Secondary.AnglePosRange
	__obf_d0c55d7e94de4ab3 := rand.IntN(len(__obf_13b6cc0d1d2c6ef9))
	if __obf_d0c55d7e94de4ab3 < 0 {
		return 0
	}
	__obf_0076db988688d2a8 := __obf_13b6cc0d1d2c6ef9[__obf_d0c55d7e94de4ab3]
	return util.RandInt(__obf_0076db988688d2a8.Min, __obf_0076db988688d2a8.Max)
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
func (__obf_6c40c4c2143a89cf *Captcha) __obf_c6431c3f8222f8a9(__obf_76308d3d64d47828 types.Size, __obf_067c37c603bffc8d types.Range, __obf_86b43f9d7c7452cd int) ([]*types.Block, types.Point) {
	var __obf_1edab5b16168422c = make([]*types.Block, 0, __obf_86b43f9d7c7452cd)
	__obf_b5007f75c87ca886 := __obf_76308d3d64d47828.Width
	__obf_f39217514bf96315 := __obf_76308d3d64d47828.Height
	__obf_37e0971ef3dd4b08 := __obf_6c40c4c2143a89cf.__obf_1e2b8a0c7d68a9f4()
	__obf_a3689b7c0e0a9a9c := util.RandInt(__obf_067c37c603bffc8d.Min, __obf_067c37c603bffc8d.Max)
	__obf_99a2b3d0bd1ca1a5 := __obf_a3689b7c0e0a9a9c
	__obf_7326020f8165ae8d := __obf_a3689b7c0e0a9a9c
	__obf_c882f20d3c1eb5f3 := __obf_6c40c4c2143a89cf.__obf_9e8e6adb3493bd16()
	__obf_62a130791db26e35 := __obf_7326020f8165ae8d / 2
	__obf_01fbb5115ff257fb := (__obf_b5007f75c87ca886 - __obf_7326020f8165ae8d - 20) / __obf_86b43f9d7c7452cd
	__obf_a365fcf4931c088a := __obf_6c40c4c2143a89cf.__obf_966f2e37d66f0b3a(5, __obf_f39217514bf96315-__obf_99a2b3d0bd1ca1a5-5, __obf_99a2b3d0bd1ca1a5, __obf_c882f20d3c1eb5f3)

	for __obf_b30879033ba2ce97 := range __obf_86b43f9d7c7452cd {
		var __obf_fc1b3b29182d412e = &types.Block{}
		__obf_c31ad7db39476e96, __obf_b420abdd9570f880 := __obf_6c40c4c2143a89cf.__obf_255dec706da37109((__obf_b30879033ba2ce97*__obf_01fbb5115ff257fb)+__obf_62a130791db26e35+5, ((__obf_b30879033ba2ce97+1)*__obf_01fbb5115ff257fb)-__obf_62a130791db26e35, __obf_7326020f8165ae8d, __obf_c882f20d3c1eb5f3)
		__obf_c31ad7db39476e96 = int(max(float64(__obf_c31ad7db39476e96), float64(__obf_62a130791db26e35+5)))
		__obf_fc1b3b29182d412e.
			X = util.RandInt(__obf_c31ad7db39476e96+20, __obf_b420abdd9570f880+20) - __obf_62a130791db26e35

		if __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Secondary.EnableVerticalRandom {
			__obf_a365fcf4931c088a = __obf_6c40c4c2143a89cf.__obf_966f2e37d66f0b3a(5, __obf_f39217514bf96315-__obf_99a2b3d0bd1ca1a5-5, __obf_99a2b3d0bd1ca1a5, __obf_c882f20d3c1eb5f3)
		}
		__obf_fc1b3b29182d412e.
			Y = __obf_a365fcf4931c088a
		__obf_fc1b3b29182d412e.
			Width = __obf_7326020f8165ae8d
		__obf_fc1b3b29182d412e.
			Height = __obf_99a2b3d0bd1ca1a5
		__obf_fc1b3b29182d412e.
			Angle = __obf_37e0971ef3dd4b08
		__obf_1edab5b16168422c = append(__obf_1edab5b16168422c, __obf_fc1b3b29182d412e)
	}
	__obf_7c26d84cefb213cc := types.Point{}
	if __obf_6c40c4c2143a89cf.__obf_195f5263d9c40fa7.Type == Move {
		__obf_7c26d84cefb213cc.
			X = util.RandInt(5, __obf_62a130791db26e35)
		__obf_7c26d84cefb213cc.
			Y = __obf_a365fcf4931c088a
		return __obf_1edab5b16168422c, __obf_7c26d84cefb213cc
	}

	switch __obf_c882f20d3c1eb5f3 {
	case DeadZoneDirectionTypeTop:
		__obf_7c26d84cefb213cc.
			X = util.RandInt(5, __obf_b5007f75c87ca886-__obf_7326020f8165ae8d-5)
		__obf_7c26d84cefb213cc.
			Y = 5
	case DeadZoneDirectionTypeBottom:
		__obf_7c26d84cefb213cc.
			X = util.RandInt(5, __obf_b5007f75c87ca886-__obf_7326020f8165ae8d-5)
		__obf_7c26d84cefb213cc.
			Y = __obf_f39217514bf96315 - __obf_99a2b3d0bd1ca1a5 - 5
	case DeadZoneDirectionTypeLeft:
		__obf_7c26d84cefb213cc.
			X = 5
		__obf_7c26d84cefb213cc.
			Y = util.RandInt(5, __obf_f39217514bf96315-__obf_99a2b3d0bd1ca1a5-5)
	case DeadZoneDirectionTypeRight:
		__obf_7c26d84cefb213cc.
			X = __obf_b5007f75c87ca886 - __obf_7326020f8165ae8d - 5
		__obf_7c26d84cefb213cc.
			Y = util.RandInt(5, __obf_f39217514bf96315-__obf_99a2b3d0bd1ca1a5-5)
	}

	return __obf_1edab5b16168422c,

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
		__obf_7c26d84cefb213cc
}

func (__obf_6c40c4c2143a89cf *Captcha) __obf_255dec706da37109(__obf_c31ad7db39476e96, __obf_b420abdd9570f880, __obf_a2de9a4ccbb03daf int, __obf_c882f20d3c1eb5f3 DeadZoneDirectionType) (int, int) {
	if __obf_c882f20d3c1eb5f3 == DeadZoneDirectionTypeLeft {
		__obf_c31ad7db39476e96 += __obf_a2de9a4ccbb03daf
		__obf_b420abdd9570f880 += __obf_a2de9a4ccbb03daf
	}
	return __obf_c31ad7db39476e96,

		// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
		// params:
		//   - start: Start Y coordinate
		//   - end: End Y coordinate
		//   - value: types.Block height
		//   - dzdType: Dead zone direction
		//
		// return: Random Y coordinate
		__obf_b420abdd9570f880
}

func (__obf_6c40c4c2143a89cf *Captcha) __obf_966f2e37d66f0b3a(__obf_c31ad7db39476e96, __obf_b420abdd9570f880, __obf_a2de9a4ccbb03daf int, __obf_c882f20d3c1eb5f3 DeadZoneDirectionType) int {
	switch __obf_c882f20d3c1eb5f3 {
	case DeadZoneDirectionTypeTop:
		__obf_c31ad7db39476e96 += __obf_a2de9a4ccbb03daf
	case DeadZoneDirectionTypeBottom:
		__obf_b420abdd9570f880 -= __obf_a2de9a4ccbb03daf
	}
	return util.RandInt(__obf_c31ad7db39476e96, __obf_b420abdd9570f880)
}
