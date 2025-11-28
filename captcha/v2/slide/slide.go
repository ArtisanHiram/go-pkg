package __obf_40cc6e54bc08fb7a

import (
	"image"
	"math/rand/v2"

	core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"
	"golang.org/x/image/draw"
)

// Captcha is the concrete implementation of the Captcha interface
type Captcha struct {
	__obf_356cde2aed9a2fa6 *Options
}

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_a58a637c494be577 Option) *Captcha {
	__obf_356cde2aed9a2fa6 := &Options{
		Primary: Primary{
			Size:  core.Size{Width: 400, Height: 340},
			Alpha: 1,
		},
		Secondary: Secondary{
			CountRange: core.Range{Min: 1, Max: 5},
			SizeRange:  core.Range{Min: 60, Max: 70},
			AnglePosRange: []core.Range{
				{Min: 0, Max: 0},
			},
		},
	}

	if __obf_a58a637c494be577 != nil {
		__obf_a58a637c494be577(__obf_356cde2aed9a2fa6)
	}

	return &Captcha{
		__obf_356cde2aed9a2fa6: __obf_356cde2aed9a2fa6,
	}
}

func __obf_7a27a89ce9140e0c(__obf_356cde2aed9a2fa6 *Options) SlideType {
	__obf_932517ec293f769f := SlideType(rand.IntN(2))
	if __obf_932517ec293f769f == Move {
		__obf_356cde2aed9a2fa6.Secondary.DeadZoneDirections = []DeadZoneDirectionType{DeadZoneDirectionTypeLeft}
		__obf_356cde2aed9a2fa6.Secondary.EnableVerticalRandom = false
	} else {
		__obf_356cde2aed9a2fa6.Secondary.DeadZoneDirections = []DeadZoneDirectionType{
			DeadZoneDirectionTypeLeft,
			DeadZoneDirectionTypeRight,
			DeadZoneDirectionTypeBottom,
			DeadZoneDirectionTypeTop,
		}
		__obf_356cde2aed9a2fa6.Secondary.EnableVerticalRandom = true
	}
	return __obf_932517ec293f769f
}

// DrawWithNRGBA draws the main CAPTCHA image and background image using NRGBA format
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn CAPTCHA image
//   - image.Image: Drawn background image
//   - error: Error information
func (__obf_4504122ced09d7a1 *Captcha) DrawWithNRGBA(__obf_7cfd455ce2e65124 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_c51c8ea6a38306c5 := __obf_7cfd455ce2e65124.DrawBlocks
	__obf_ae84f678679a2e10 := core.CreateNRGBACanvas(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, true)

	for _, __obf_956a76774174edc5 := range __obf_c51c8ea6a38306c5 {
		__obf_6cbab8354a96222a, __obf_53d87ed04712e1d6 := __obf_4504122ced09d7a1.__obf_1df7722f401953ef(__obf_956a76774174edc5.Width, __obf_956a76774174edc5.Height, __obf_956a76774174edc5.Shape)
		if __obf_53d87ed04712e1d6 != nil {
			return nil, nil, __obf_53d87ed04712e1d6
		}

		__obf_e4eaad5d5c2b5684 := __obf_6cbab8354a96222a.Bounds()
		draw.Draw(__obf_ae84f678679a2e10, image.Rect(__obf_956a76774174edc5.X, __obf_956a76774174edc5.Y, __obf_956a76774174edc5.X+__obf_e4eaad5d5c2b5684.Dx(), __obf_956a76774174edc5.Y+__obf_e4eaad5d5c2b5684.Dy()), __obf_6cbab8354a96222a, image.Point{}, draw.Over)
	}

	__obf_51990625e392ddaa := core.CreateNRGBACanvas(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, true)
	if __obf_7cfd455ce2e65124.Background != nil {
		__obf_227dadaf2f81df76 := __obf_7cfd455ce2e65124.Background
		__obf_956a76774174edc5 := __obf_227dadaf2f81df76.Bounds()
		__obf_f4b9767290b8c8eb := core.CreateNRGBACanvas(__obf_956a76774174edc5.Dx(), __obf_956a76774174edc5.Dy(), true)
		__obf_dd8ccf8e1e49adba := core.RangCutImagePos(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, __obf_227dadaf2f81df76)
		draw.Draw(__obf_f4b9767290b8c8eb, __obf_956a76774174edc5, __obf_227dadaf2f81df76, __obf_dd8ccf8e1e49adba, draw.Src)
		__obf_f4b9767290b8c8eb.SubImage(image.Rect(0, 0, __obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height))

		draw.Draw(__obf_51990625e392ddaa, __obf_51990625e392ddaa.Bounds(), __obf_f4b9767290b8c8eb, image.Point{}, draw.Over)
		draw.Draw(__obf_f4b9767290b8c8eb, __obf_ae84f678679a2e10.Bounds(), __obf_ae84f678679a2e10, image.Point{}, draw.Over)
		return __obf_f4b9767290b8c8eb, __obf_51990625e392ddaa, nil
	}

	return __obf_ae84f678679a2e10, __obf_51990625e392ddaa, nil
}

// DrawWithTemplate draws the tile image using a template
// params:
//   - params: Drawing parameters
//
// returns:
//   - image.Image: Drawn tile image
//   - error: Error information
func (__obf_4504122ced09d7a1 *Captcha) DrawWithTemplate(__obf_7cfd455ce2e65124 *DrawTplImageParam) (image.Image, error) {
	__obf_d39062b0036fd87e := __obf_7cfd455ce2e65124.Block
	__obf_227dadaf2f81df76 := __obf_7cfd455ce2e65124.Background
	__obf_ae84f678679a2e10 := core.CreateNRGBACanvas(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, true)
	__obf_5410655c2a758da6 := core.CreateNRGBACanvas(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, true)

	__obf_6e8d638078485263, __obf_53d87ed04712e1d6 := __obf_4504122ced09d7a1.__obf_1df7722f401953ef(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, __obf_7cfd455ce2e65124.Mask)
	if __obf_53d87ed04712e1d6 != nil {
		return nil, __obf_53d87ed04712e1d6
	}

	draw.Draw(__obf_5410655c2a758da6, __obf_5410655c2a758da6.Bounds(), __obf_227dadaf2f81df76, image.Pt(__obf_d39062b0036fd87e.X, __obf_d39062b0036fd87e.Y), draw.Src)
	draw.DrawMask(__obf_ae84f678679a2e10, __obf_6e8d638078485263.Bounds(), __obf_5410655c2a758da6, image.Point{}, __obf_6e8d638078485263, image.Point{}, draw.Over)

	__obf_f7c67dbd2a945e6f, __obf_53d87ed04712e1d6 := __obf_4504122ced09d7a1.__obf_1df7722f401953ef(__obf_7cfd455ce2e65124.Width, __obf_7cfd455ce2e65124.Height, __obf_d39062b0036fd87e.Shape)
	if __obf_53d87ed04712e1d6 != nil {
		return nil, __obf_53d87ed04712e1d6
	}
	draw.Draw(__obf_ae84f678679a2e10, __obf_f7c67dbd2a945e6f.Bounds(), __obf_f7c67dbd2a945e6f, image.Point{}, draw.Over)

	return __obf_ae84f678679a2e10, nil
}

// drawGraphImage draws a graph image
// params:
//   - width: Image width
//   - height: Image height
//   - img: Input image
//
// returns:
//   - core.NRGBA: Drawn graph canvas
//   - error: Error information
func (__obf_4504122ced09d7a1 *Captcha) __obf_1df7722f401953ef(__obf_8b378c478e5667b0, __obf_329eb3efa9d84dd5 int, __obf_776c3bf51a3e4487 image.Image) (*core.NRGBA, error) {
	__obf_ae84f678679a2e10 := core.CreateNRGBACanvas(__obf_8b378c478e5667b0, __obf_329eb3efa9d84dd5, true)
	draw.BiLinear.Scale(__obf_ae84f678679a2e10, __obf_ae84f678679a2e10.Bounds(), __obf_776c3bf51a3e4487, __obf_776c3bf51a3e4487.Bounds(), draw.Over, nil)
	return __obf_ae84f678679a2e10, nil
}

// Generate generates slide CAPTCHA data
// returns:
//   - CaptchaData: Generated CAPTCHA data
//   - error: Error information
func (__obf_4504122ced09d7a1 *Captcha) Generate() (*CaptchaData, error) {

	__obf_932517ec293f769f := __obf_7a27a89ce9140e0c(__obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6)
	__obf_28ced6397300684b, __obf_53d87ed04712e1d6 := core.PickTile()
	if __obf_53d87ed04712e1d6 != nil {
		return nil, __obf_53d87ed04712e1d6
	}

	__obf_1bb8492670fec5b7 := core.RandInt(__obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Secondary.CountRange.Min, __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Secondary.CountRange.Max)
	__obf_c51c8ea6a38306c5, __obf_376cd763d261f6a8 := __obf_4504122ced09d7a1.__obf_712bdc7569ff3521(__obf_932517ec293f769f, __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Primary.Size, __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Secondary.SizeRange, __obf_1bb8492670fec5b7)
	var __obf_d39062b0036fd87e *Block
	if len(__obf_c51c8ea6a38306c5) > 1 {
		__obf_5e256b942443cb53 := max(rand.IntN(len(__obf_c51c8ea6a38306c5)), 0)
		__obf_d39062b0036fd87e = __obf_c51c8ea6a38306c5[__obf_5e256b942443cb53]
	} else {
		__obf_d39062b0036fd87e = __obf_c51c8ea6a38306c5[0]
	}

	__obf_162c7157de6fcfa9, __obf_9d0dd1b3c07a8234, __obf_53d87ed04712e1d6 := __obf_4504122ced09d7a1.__obf_4d1adf69bd0182ed(__obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Primary.Size, __obf_28ced6397300684b.Shadow, __obf_c51c8ea6a38306c5)
	if __obf_53d87ed04712e1d6 != nil {
		return nil, __obf_53d87ed04712e1d6
	}

	__obf_d39062b0036fd87e.Shape = __obf_28ced6397300684b.Overlay

	__obf_c4b0de4f376b141e, __obf_53d87ed04712e1d6 := __obf_4504122ced09d7a1.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_9d0dd1b3c07a8234,
		Mask:       __obf_28ced6397300684b.Mask,
		Alpha:      __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Primary.Alpha,
		Width:      __obf_d39062b0036fd87e.Width,
		Height:     __obf_d39062b0036fd87e.Height,
		Block:      __obf_d39062b0036fd87e,
	})

	if __obf_53d87ed04712e1d6 != nil {
		return nil, __obf_53d87ed04712e1d6
	}

	if __obf_932517ec293f769f == Move {
		__obf_d39062b0036fd87e.DY = __obf_d39062b0036fd87e.Y
	} else {
		__obf_d39062b0036fd87e.DY = __obf_376cd763d261f6a8.Y
	}
	__obf_d39062b0036fd87e.DX = __obf_376cd763d261f6a8.X

	return &CaptchaData{
		__obf_d39062b0036fd87e: __obf_d39062b0036fd87e,
		__obf_162c7157de6fcfa9: core.NewJPEGImage(__obf_162c7157de6fcfa9),
		__obf_c4b0de4f376b141e: core.NewPNGImage(__obf_c4b0de4f376b141e),
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
func (__obf_4504122ced09d7a1 *Captcha) __obf_4d1adf69bd0182ed(__obf_17ad03661e9c6c1a core.Size, __obf_ca9bc25a295adadb image.Image, __obf_c51c8ea6a38306c5 []*Block) (image.Image, image.Image, error) {
	var __obf_3abddbfb1dd8c94e = make([]*Block, 0, len(__obf_c51c8ea6a38306c5))
	for _, __obf_6980e5c85fbbb5bd := range __obf_c51c8ea6a38306c5 {
		__obf_6980e5c85fbbb5bd.Shape = __obf_ca9bc25a295adadb
		__obf_3abddbfb1dd8c94e = append(__obf_3abddbfb1dd8c94e, __obf_6980e5c85fbbb5bd)
	}

	__obf_e67c185992218b79, __obf_53d87ed04712e1d6 := core.PickBgImage()
	if __obf_53d87ed04712e1d6 != nil {
		return nil, nil, __obf_53d87ed04712e1d6
	}

	return __obf_4504122ced09d7a1.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_17ad03661e9c6c1a.Width,
		Height:     __obf_17ad03661e9c6c1a.Height,
		Background: __obf_e67c185992218b79,
		Alpha:      __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Primary.Alpha,
		DrawBlocks: __obf_3abddbfb1dd8c94e,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_4504122ced09d7a1 *Captcha) __obf_95c52f438bb3b12b() DeadZoneDirectionType {
	__obf_1ace0d1460d43ed2 := __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Secondary.DeadZoneDirections

	__obf_5e256b942443cb53 := rand.IntN(len(__obf_1ace0d1460d43ed2))
	if __obf_5e256b942443cb53 < 0 {
		return 0
	}

	return __obf_1ace0d1460d43ed2[__obf_5e256b942443cb53]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_4504122ced09d7a1 *Captcha) __obf_c20081778ac9c705() int {
	__obf_b760753968604b26 := __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Secondary.AnglePosRange

	__obf_5e256b942443cb53 := rand.IntN(len(__obf_b760753968604b26))
	if __obf_5e256b942443cb53 < 0 {
		return 0
	}

	__obf_6b02765c8f9d1b0d := __obf_b760753968604b26[__obf_5e256b942443cb53]
	return core.RandInt(__obf_6b02765c8f9d1b0d.Min, __obf_6b02765c8f9d1b0d.Max)
}

// genGraphBlocks generates graph block data
// params:
//   - imageSize: Main image size
//   - size: Graph size range
//   - length: Number of graphs
//
// returns:
//   - []*Block: List of blocks
//   - core.Point: Tile position
func (__obf_4504122ced09d7a1 *Captcha) __obf_712bdc7569ff3521(__obf_d9f7b533aa0ebdaf SlideType, __obf_e5590608d474a9b6 core.Size, __obf_17ad03661e9c6c1a core.Range, __obf_1e1ed6ea5b5777e8 int) ([]*Block, core.Point) {
	var __obf_c51c8ea6a38306c5 = make([]*Block, 0, __obf_1e1ed6ea5b5777e8)
	__obf_8b378c478e5667b0 := __obf_e5590608d474a9b6.Width
	__obf_329eb3efa9d84dd5 := __obf_e5590608d474a9b6.Height

	__obf_d7d3da2ab9671daa := __obf_4504122ced09d7a1.__obf_c20081778ac9c705()
	__obf_357ce54e79cd5ab0 := core.RandInt(__obf_17ad03661e9c6c1a.Min, __obf_17ad03661e9c6c1a.Max)
	__obf_f57654d33014d88c := __obf_357ce54e79cd5ab0
	__obf_0de822b413d99421 := __obf_357ce54e79cd5ab0

	__obf_2cee0485b72f98f2 := __obf_4504122ced09d7a1.__obf_95c52f438bb3b12b()
	__obf_60c627cf78296b1d := __obf_0de822b413d99421 / 2
	__obf_ece28bbadf1892e9 := (__obf_8b378c478e5667b0 - __obf_0de822b413d99421 - 20) / __obf_1e1ed6ea5b5777e8
	__obf_c81c8e0ea0df30f1 := __obf_4504122ced09d7a1.__obf_3ce956ee6f582bac(5, __obf_329eb3efa9d84dd5-__obf_f57654d33014d88c-5, __obf_f57654d33014d88c, __obf_2cee0485b72f98f2)

	for __obf_735c1c783930f9e3 := range __obf_1e1ed6ea5b5777e8 {
		var __obf_d39062b0036fd87e = &Block{}
		__obf_de9ed5d7c36fb14c, __obf_70c4811671338257 := __obf_4504122ced09d7a1.__obf_ba1958fa925f9d94((__obf_735c1c783930f9e3*__obf_ece28bbadf1892e9)+__obf_60c627cf78296b1d+5, ((__obf_735c1c783930f9e3+1)*__obf_ece28bbadf1892e9)-__obf_60c627cf78296b1d, __obf_0de822b413d99421, __obf_2cee0485b72f98f2)

		__obf_de9ed5d7c36fb14c = int(max(float64(__obf_de9ed5d7c36fb14c), float64(__obf_60c627cf78296b1d+5)))
		__obf_d39062b0036fd87e.X = core.RandInt(__obf_de9ed5d7c36fb14c+20, __obf_70c4811671338257+20) - __obf_60c627cf78296b1d

		if __obf_4504122ced09d7a1.__obf_356cde2aed9a2fa6.Secondary.EnableVerticalRandom {
			__obf_c81c8e0ea0df30f1 = __obf_4504122ced09d7a1.__obf_3ce956ee6f582bac(5, __obf_329eb3efa9d84dd5-__obf_f57654d33014d88c-5, __obf_f57654d33014d88c, __obf_2cee0485b72f98f2)
		}

		__obf_d39062b0036fd87e.Y = __obf_c81c8e0ea0df30f1
		__obf_d39062b0036fd87e.Width = __obf_0de822b413d99421
		__obf_d39062b0036fd87e.Height = __obf_f57654d33014d88c
		__obf_d39062b0036fd87e.Angle = __obf_d7d3da2ab9671daa

		__obf_c51c8ea6a38306c5 = append(__obf_c51c8ea6a38306c5, __obf_d39062b0036fd87e)
	}

	__obf_dd8ccf8e1e49adba := core.Point{}
	if __obf_d9f7b533aa0ebdaf == Move {
		__obf_dd8ccf8e1e49adba.X = core.RandInt(5, __obf_60c627cf78296b1d)
		__obf_dd8ccf8e1e49adba.Y = __obf_c81c8e0ea0df30f1
		return __obf_c51c8ea6a38306c5, __obf_dd8ccf8e1e49adba
	}

	switch __obf_2cee0485b72f98f2 {
	case DeadZoneDirectionTypeTop:
		__obf_dd8ccf8e1e49adba.X = core.RandInt(5, __obf_8b378c478e5667b0-__obf_0de822b413d99421-5)
		__obf_dd8ccf8e1e49adba.Y = 5
	case DeadZoneDirectionTypeBottom:
		__obf_dd8ccf8e1e49adba.X = core.RandInt(5, __obf_8b378c478e5667b0-__obf_0de822b413d99421-5)
		__obf_dd8ccf8e1e49adba.Y = __obf_329eb3efa9d84dd5 - __obf_f57654d33014d88c - 5
	case DeadZoneDirectionTypeLeft:
		__obf_dd8ccf8e1e49adba.X = 5
		__obf_dd8ccf8e1e49adba.Y = core.RandInt(5, __obf_329eb3efa9d84dd5-__obf_f57654d33014d88c-5)
	case DeadZoneDirectionTypeRight:
		__obf_dd8ccf8e1e49adba.X = __obf_8b378c478e5667b0 - __obf_0de822b413d99421 - 5
		__obf_dd8ccf8e1e49adba.Y = core.RandInt(5, __obf_329eb3efa9d84dd5-__obf_f57654d33014d88c-5)
	}

	return __obf_c51c8ea6a38306c5, __obf_dd8ccf8e1e49adba
}

// calcXWithDeadZone calculates the X coordinate range (considering dead zone)
// params:
//   - start: Start X coordinate
//   - end: End X coordinate
//   - value: Block width
//   - dzdType: Dead zone direction
//
// returns:
//   - int: Adjusted start X coordinate
//   - int: Adjusted end X coordinate
func (__obf_4504122ced09d7a1 *Captcha) __obf_ba1958fa925f9d94(__obf_de9ed5d7c36fb14c, __obf_70c4811671338257, __obf_9c8fee8af9e383c8 int, __obf_2cee0485b72f98f2 DeadZoneDirectionType) (int, int) {
	if __obf_2cee0485b72f98f2 == DeadZoneDirectionTypeLeft {
		__obf_de9ed5d7c36fb14c += __obf_9c8fee8af9e383c8
		__obf_70c4811671338257 += __obf_9c8fee8af9e383c8
	}
	return __obf_de9ed5d7c36fb14c, __obf_70c4811671338257
}

// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
// params:
//   - start: Start Y coordinate
//   - end: End Y coordinate
//   - value: Block height
//   - dzdType: Dead zone direction
//
// return: Random Y coordinate
func (__obf_4504122ced09d7a1 *Captcha) __obf_3ce956ee6f582bac(__obf_de9ed5d7c36fb14c, __obf_70c4811671338257, __obf_9c8fee8af9e383c8 int, __obf_2cee0485b72f98f2 DeadZoneDirectionType) int {
	switch __obf_2cee0485b72f98f2 {
	case DeadZoneDirectionTypeTop:
		__obf_de9ed5d7c36fb14c += __obf_9c8fee8af9e383c8
	case DeadZoneDirectionTypeBottom:
		__obf_70c4811671338257 -= __obf_9c8fee8af9e383c8
	}
	return core.RandInt(__obf_de9ed5d7c36fb14c, __obf_70c4811671338257)
}
