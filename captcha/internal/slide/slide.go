package __obf_741aa76a5fc5f3d9

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
	__obf_a96d72e19f373063 *Options
}

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_a96d72e19f373063 *Options) *Captcha {
	if __obf_a96d72e19f373063 == nil {
		return nil
	}
	return &Captcha{
		__obf_a96d72e19f373063: __obf_a96d72e19f373063,
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
func (__obf_88d29df8b03b1a42 *Captcha) DrawWithNRGBA(__obf_97771edf55c9150e *DrawImageParam) (image.Image, image.Image, error) {
	__obf_52a829ebf5856498 := __obf_97771edf55c9150e.DrawBlocks
	__obf_392875deaa5d9ad6 := helper.CreateNRGBACanvas(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, true)

	for _, __obf_2d4441a8bdca1c7f := range __obf_52a829ebf5856498 {
		__obf_63ea0ae3ce3b5061, __obf_85d126d2e91a5e83 := __obf_88d29df8b03b1a42.__obf_c397241adb840cbb(__obf_2d4441a8bdca1c7f.Width, __obf_2d4441a8bdca1c7f.Height, __obf_2d4441a8bdca1c7f.Shape)
		if __obf_85d126d2e91a5e83 != nil {
			return nil, nil, __obf_85d126d2e91a5e83
		}

		__obf_b8a354b47f500b42 := __obf_63ea0ae3ce3b5061.Bounds()
		draw.Draw(__obf_392875deaa5d9ad6, image.Rect(__obf_2d4441a8bdca1c7f.X, __obf_2d4441a8bdca1c7f.Y, __obf_2d4441a8bdca1c7f.X+__obf_b8a354b47f500b42.Dx(), __obf_2d4441a8bdca1c7f.Y+__obf_b8a354b47f500b42.Dy()), __obf_63ea0ae3ce3b5061, image.Point{}, draw.Over)
	}

	__obf_98ec1a56f54c3d03 := helper.CreateNRGBACanvas(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, true)
	if __obf_97771edf55c9150e.Background != nil {
		__obf_1d3eb3dac61c8939 := __obf_97771edf55c9150e.Background
		__obf_2d4441a8bdca1c7f := __obf_1d3eb3dac61c8939.Bounds()
		__obf_61bca9c721e29bae := helper.CreateNRGBACanvas(__obf_2d4441a8bdca1c7f.Dx(), __obf_2d4441a8bdca1c7f.Dy(), true)
		__obf_082b1de38cb8ddab := util.RangCutImagePos(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, __obf_1d3eb3dac61c8939)
		draw.Draw(__obf_61bca9c721e29bae, __obf_2d4441a8bdca1c7f, __obf_1d3eb3dac61c8939, __obf_082b1de38cb8ddab, draw.Src)
		__obf_61bca9c721e29bae.SubImage(image.Rect(0, 0, __obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height))

		draw.Draw(__obf_98ec1a56f54c3d03, __obf_98ec1a56f54c3d03.Bounds(), __obf_61bca9c721e29bae, image.Point{}, draw.Over)
		draw.Draw(__obf_61bca9c721e29bae, __obf_392875deaa5d9ad6.Bounds(), __obf_392875deaa5d9ad6, image.Point{}, draw.Over)
		return __obf_61bca9c721e29bae, __obf_98ec1a56f54c3d03, nil
	}

	return __obf_392875deaa5d9ad6, __obf_98ec1a56f54c3d03, nil
}

// DrawWithTemplate draws the tile image using a template
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn tile image
//   - error: Error information
func (__obf_88d29df8b03b1a42 *Captcha) DrawWithTemplate(__obf_97771edf55c9150e *DrawTplImageParam) (image.Image, error) {
	__obf_9cac0044fa40322d := __obf_97771edf55c9150e.Block
	__obf_1d3eb3dac61c8939 := __obf_97771edf55c9150e.Background
	__obf_392875deaa5d9ad6 := helper.CreateNRGBACanvas(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, true)
	__obf_f61f85fa1a65e999 := helper.CreateNRGBACanvas(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, true)

	__obf_ed05f38e220fcaff, __obf_85d126d2e91a5e83 := __obf_88d29df8b03b1a42.__obf_c397241adb840cbb(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, __obf_97771edf55c9150e.Mask)
	if __obf_85d126d2e91a5e83 != nil {
		return nil, __obf_85d126d2e91a5e83
	}

	draw.Draw(__obf_f61f85fa1a65e999, __obf_f61f85fa1a65e999.Bounds(), __obf_1d3eb3dac61c8939, image.Pt(__obf_9cac0044fa40322d.X, __obf_9cac0044fa40322d.Y), draw.Src)
	draw.DrawMask(__obf_392875deaa5d9ad6, __obf_ed05f38e220fcaff.Bounds(), __obf_f61f85fa1a65e999, image.Point{}, __obf_ed05f38e220fcaff, image.Point{}, draw.Over)

	__obf_0862dafcca3eb8e9, __obf_85d126d2e91a5e83 := __obf_88d29df8b03b1a42.__obf_c397241adb840cbb(__obf_97771edf55c9150e.Width, __obf_97771edf55c9150e.Height, __obf_9cac0044fa40322d.Shape)
	if __obf_85d126d2e91a5e83 != nil {
		return nil, __obf_85d126d2e91a5e83
	}
	draw.Draw(__obf_392875deaa5d9ad6, __obf_0862dafcca3eb8e9.Bounds(), __obf_0862dafcca3eb8e9, image.Point{}, draw.Over)

	return __obf_392875deaa5d9ad6, nil
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
func (__obf_88d29df8b03b1a42 *Captcha) __obf_c397241adb840cbb(__obf_b460933246891894, __obf_17863730397ae92c int, __obf_1620f85a61f850fb image.Image) (*types.NRGBA, error) {
	__obf_392875deaa5d9ad6 := helper.CreateNRGBACanvas(__obf_b460933246891894, __obf_17863730397ae92c, true)
	draw.BiLinear.Scale(__obf_392875deaa5d9ad6, __obf_392875deaa5d9ad6.Bounds(), __obf_1620f85a61f850fb, __obf_1620f85a61f850fb.Bounds(), draw.Over, nil)
	return __obf_392875deaa5d9ad6, nil
}

// Generate generates slide CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_88d29df8b03b1a42 *Captcha) Generate() (*CaptchaData, error) {

	// slideType := randomSlideType(c.opts)
	__obf_68817fcabb375001, __obf_85d126d2e91a5e83 := assets.PickTile()
	if __obf_85d126d2e91a5e83 != nil {
		return nil, __obf_85d126d2e91a5e83
	}

	__obf_ece83b7548236782 := util.RandInt(__obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Secondary.CountRange.Min, __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Secondary.CountRange.Max)
	__obf_52a829ebf5856498, __obf_815a328a4a78c79c := __obf_88d29df8b03b1a42.__obf_64da4b79b6fd7344(__obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Primary.Size, __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Secondary.SizeRange, __obf_ece83b7548236782)
	var __obf_9cac0044fa40322d *types.Block
	if len(__obf_52a829ebf5856498) > 1 {
		__obf_6a98a025611c552e := max(rand.IntN(len(__obf_52a829ebf5856498)), 0)
		__obf_9cac0044fa40322d = __obf_52a829ebf5856498[__obf_6a98a025611c552e]
	} else {
		__obf_9cac0044fa40322d = __obf_52a829ebf5856498[0]
	}

	__obf_e80f5dd754e39a7c, __obf_726360c5815c51dc, __obf_85d126d2e91a5e83 := __obf_88d29df8b03b1a42.__obf_658505deea819366(__obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Primary.Size, __obf_68817fcabb375001.Shadow, __obf_52a829ebf5856498)
	if __obf_85d126d2e91a5e83 != nil {
		return nil, __obf_85d126d2e91a5e83
	}

	__obf_9cac0044fa40322d.Shape = __obf_68817fcabb375001.Overlay

	__obf_580c5dd4e04d5aad, __obf_85d126d2e91a5e83 := __obf_88d29df8b03b1a42.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_726360c5815c51dc,
		Mask:       __obf_68817fcabb375001.Mask,
		Alpha:      __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Primary.Alpha,
		Width:      __obf_9cac0044fa40322d.Width,
		Height:     __obf_9cac0044fa40322d.Height,
		Block:      __obf_9cac0044fa40322d,
	})

	if __obf_85d126d2e91a5e83 != nil {
		return nil, __obf_85d126d2e91a5e83
	}

	if __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Type == Move {
		__obf_9cac0044fa40322d.DY = __obf_9cac0044fa40322d.Y
	} else {
		__obf_9cac0044fa40322d.DY = __obf_815a328a4a78c79c.Y
	}
	__obf_9cac0044fa40322d.DX = __obf_815a328a4a78c79c.X

	return &CaptchaData{
		__obf_044d58016cdb062b: __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Type,
		__obf_9cac0044fa40322d: __obf_9cac0044fa40322d,
		__obf_e80f5dd754e39a7c: types.NewJPEGImage(__obf_e80f5dd754e39a7c),
		__obf_580c5dd4e04d5aad: types.NewPNGImage(__obf_580c5dd4e04d5aad),
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
func (__obf_88d29df8b03b1a42 *Captcha) __obf_658505deea819366(__obf_5dbe57b866017c43 types.Size, __obf_8e9d3c0dec98ef9a image.Image, __obf_52a829ebf5856498 []*types.Block) (image.Image, image.Image, error) {
	var __obf_2cae0cafa55cd190 = make([]*types.Block, 0, len(__obf_52a829ebf5856498))
	for _, __obf_8ac92e8cba6e9fc6 := range __obf_52a829ebf5856498 {
		__obf_8ac92e8cba6e9fc6.Shape = __obf_8e9d3c0dec98ef9a
		__obf_2cae0cafa55cd190 = append(__obf_2cae0cafa55cd190, __obf_8ac92e8cba6e9fc6)
	}

	__obf_e057eea84b67fa7f, __obf_85d126d2e91a5e83 := assets.PickBgImage()
	if __obf_85d126d2e91a5e83 != nil {
		return nil, nil, __obf_85d126d2e91a5e83
	}

	return __obf_88d29df8b03b1a42.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_5dbe57b866017c43.Width,
		Height:     __obf_5dbe57b866017c43.Height,
		Background: __obf_e057eea84b67fa7f,
		Alpha:      __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Primary.Alpha,
		DrawBlocks: __obf_2cae0cafa55cd190,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_88d29df8b03b1a42 *Captcha) __obf_42e0a76645ce69cc() DeadZoneDirectionType {
	__obf_91fd38692c2e653b := __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Secondary.DeadZoneDirections

	__obf_6a98a025611c552e := rand.IntN(len(__obf_91fd38692c2e653b))
	if __obf_6a98a025611c552e < 0 {
		return 0
	}

	return __obf_91fd38692c2e653b[__obf_6a98a025611c552e]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_88d29df8b03b1a42 *Captcha) __obf_a0cc8e3e67043fcc() int {
	__obf_cbb4502f7a9185c3 := __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Secondary.AnglePosRange

	__obf_6a98a025611c552e := rand.IntN(len(__obf_cbb4502f7a9185c3))
	if __obf_6a98a025611c552e < 0 {
		return 0
	}

	__obf_b2e8c4ef97898533 := __obf_cbb4502f7a9185c3[__obf_6a98a025611c552e]
	return util.RandInt(__obf_b2e8c4ef97898533.Min, __obf_b2e8c4ef97898533.Max)
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
func (__obf_88d29df8b03b1a42 *Captcha) __obf_64da4b79b6fd7344(__obf_89b63b2b2a9bfa91 types.Size, __obf_5dbe57b866017c43 types.Range, __obf_5c0ede00b6b25eec int) ([]*types.Block, types.Point) {
	var __obf_52a829ebf5856498 = make([]*types.Block, 0, __obf_5c0ede00b6b25eec)
	__obf_b460933246891894 := __obf_89b63b2b2a9bfa91.Width
	__obf_17863730397ae92c := __obf_89b63b2b2a9bfa91.Height

	__obf_aa07f6f33824bb23 := __obf_88d29df8b03b1a42.__obf_a0cc8e3e67043fcc()
	__obf_9074397a278c70a4 := util.RandInt(__obf_5dbe57b866017c43.Min, __obf_5dbe57b866017c43.Max)
	__obf_72115085392950ec := __obf_9074397a278c70a4
	__obf_44bafe160b6ae19f := __obf_9074397a278c70a4

	__obf_3b4b1220c9f28718 := __obf_88d29df8b03b1a42.__obf_42e0a76645ce69cc()
	__obf_3edfa0a4145c7c93 := __obf_44bafe160b6ae19f / 2
	__obf_84062d338f605504 := (__obf_b460933246891894 - __obf_44bafe160b6ae19f - 20) / __obf_5c0ede00b6b25eec
	__obf_08baf68fada01c6f := __obf_88d29df8b03b1a42.__obf_b36687f2d5cc98a1(5, __obf_17863730397ae92c-__obf_72115085392950ec-5, __obf_72115085392950ec, __obf_3b4b1220c9f28718)

	for __obf_8c6d46030da0f7e6 := range __obf_5c0ede00b6b25eec {
		var __obf_9cac0044fa40322d = &types.Block{}
		__obf_22fd05148ebd2927, __obf_f0a902301ae7e19e := __obf_88d29df8b03b1a42.__obf_1b8f25be0d25f2d8((__obf_8c6d46030da0f7e6*__obf_84062d338f605504)+__obf_3edfa0a4145c7c93+5, ((__obf_8c6d46030da0f7e6+1)*__obf_84062d338f605504)-__obf_3edfa0a4145c7c93, __obf_44bafe160b6ae19f, __obf_3b4b1220c9f28718)

		__obf_22fd05148ebd2927 = int(max(float64(__obf_22fd05148ebd2927), float64(__obf_3edfa0a4145c7c93+5)))
		__obf_9cac0044fa40322d.X = util.RandInt(__obf_22fd05148ebd2927+20, __obf_f0a902301ae7e19e+20) - __obf_3edfa0a4145c7c93

		if __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Secondary.EnableVerticalRandom {
			__obf_08baf68fada01c6f = __obf_88d29df8b03b1a42.__obf_b36687f2d5cc98a1(5, __obf_17863730397ae92c-__obf_72115085392950ec-5, __obf_72115085392950ec, __obf_3b4b1220c9f28718)
		}

		__obf_9cac0044fa40322d.Y = __obf_08baf68fada01c6f
		__obf_9cac0044fa40322d.Width = __obf_44bafe160b6ae19f
		__obf_9cac0044fa40322d.Height = __obf_72115085392950ec
		__obf_9cac0044fa40322d.Angle = __obf_aa07f6f33824bb23

		__obf_52a829ebf5856498 = append(__obf_52a829ebf5856498, __obf_9cac0044fa40322d)
	}

	__obf_082b1de38cb8ddab := types.Point{}
	if __obf_88d29df8b03b1a42.__obf_a96d72e19f373063.Type == Move {
		__obf_082b1de38cb8ddab.X = util.RandInt(5, __obf_3edfa0a4145c7c93)
		__obf_082b1de38cb8ddab.Y = __obf_08baf68fada01c6f
		return __obf_52a829ebf5856498, __obf_082b1de38cb8ddab
	}

	switch __obf_3b4b1220c9f28718 {
	case DeadZoneDirectionTypeTop:
		__obf_082b1de38cb8ddab.X = util.RandInt(5, __obf_b460933246891894-__obf_44bafe160b6ae19f-5)
		__obf_082b1de38cb8ddab.Y = 5
	case DeadZoneDirectionTypeBottom:
		__obf_082b1de38cb8ddab.X = util.RandInt(5, __obf_b460933246891894-__obf_44bafe160b6ae19f-5)
		__obf_082b1de38cb8ddab.Y = __obf_17863730397ae92c - __obf_72115085392950ec - 5
	case DeadZoneDirectionTypeLeft:
		__obf_082b1de38cb8ddab.X = 5
		__obf_082b1de38cb8ddab.Y = util.RandInt(5, __obf_17863730397ae92c-__obf_72115085392950ec-5)
	case DeadZoneDirectionTypeRight:
		__obf_082b1de38cb8ddab.X = __obf_b460933246891894 - __obf_44bafe160b6ae19f - 5
		__obf_082b1de38cb8ddab.Y = util.RandInt(5, __obf_17863730397ae92c-__obf_72115085392950ec-5)
	}

	return __obf_52a829ebf5856498, __obf_082b1de38cb8ddab
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
func (__obf_88d29df8b03b1a42 *Captcha) __obf_1b8f25be0d25f2d8(__obf_22fd05148ebd2927, __obf_f0a902301ae7e19e, __obf_c6581ce561d53c81 int, __obf_3b4b1220c9f28718 DeadZoneDirectionType) (int, int) {
	if __obf_3b4b1220c9f28718 == DeadZoneDirectionTypeLeft {
		__obf_22fd05148ebd2927 += __obf_c6581ce561d53c81
		__obf_f0a902301ae7e19e += __obf_c6581ce561d53c81
	}
	return __obf_22fd05148ebd2927, __obf_f0a902301ae7e19e
}

// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
// params:
//   - start: Start Y coordinate
//   - end: End Y coordinate
//   - value: types.Block height
//   - dzdType: Dead zone direction
//
// return: Random Y coordinate
func (__obf_88d29df8b03b1a42 *Captcha) __obf_b36687f2d5cc98a1(__obf_22fd05148ebd2927, __obf_f0a902301ae7e19e, __obf_c6581ce561d53c81 int, __obf_3b4b1220c9f28718 DeadZoneDirectionType) int {
	switch __obf_3b4b1220c9f28718 {
	case DeadZoneDirectionTypeTop:
		__obf_22fd05148ebd2927 += __obf_c6581ce561d53c81
	case DeadZoneDirectionTypeBottom:
		__obf_f0a902301ae7e19e -= __obf_c6581ce561d53c81
	}
	return util.RandInt(__obf_22fd05148ebd2927, __obf_f0a902301ae7e19e)
}
