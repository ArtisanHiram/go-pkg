package __obf_7621171b3eab50e9

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
	__obf_bdac2ecb07870a35 *types.SlideOption
}

var _ types.Captcha = (*Captcha)(nil)

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_bdac2ecb07870a35 *types.SlideOption) *Captcha {
	if __obf_bdac2ecb07870a35 == nil {
		return nil
	}
	return &Captcha{__obf_bdac2ecb07870a35: __obf_bdac2ecb07870a35}
}

// func (c *Captcha) Type() types.CaptchaType {
// 	if c.opt.Type == Move {
// 		return types.MoveCaptchat
// 	}
// 	return types.DragCaptchat
// }

func (__obf_4c68703ee0481a75 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_902526fcec0236d1,

		// slideType := randomSlideType(c.opt)
		__obf_a309197d4a7b9caa := assets.PickTile()
	if __obf_a309197d4a7b9caa != nil {
		return nil, __obf_a309197d4a7b9caa
	}
	__obf_ced1d3e713d5dc6d := util.RandInt(__obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Secondary.CountRange.Min, __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Secondary.CountRange.Max)
	__obf_a3ad588b3910ea0b, __obf_3b80c8f4e178fd22 := __obf_4c68703ee0481a75.__obf_5398960faf9a5da3(__obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Primary.Width, __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Primary.Height, __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Secondary.SizeRange, __obf_ced1d3e713d5dc6d)
	var __obf_fe556acd0600dc8f *types.Block
	if len(__obf_a3ad588b3910ea0b) > 1 {
		__obf_6079f6833c73400b := max(rand.IntN(len(__obf_a3ad588b3910ea0b)), 0)
		__obf_fe556acd0600dc8f = __obf_a3ad588b3910ea0b[__obf_6079f6833c73400b]
	} else {
		__obf_fe556acd0600dc8f = __obf_a3ad588b3910ea0b[0]
	}
	__obf_aa601e1953c0218c, __obf_d7fe24b9cb955d9f, __obf_a309197d4a7b9caa := __obf_4c68703ee0481a75.__obf_38c3abd336ee398f(__obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Primary.Width, __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Primary.Height, __obf_902526fcec0236d1.Shadow, __obf_a3ad588b3910ea0b)
	if __obf_a309197d4a7b9caa != nil {
		return nil, __obf_a309197d4a7b9caa
	}
	__obf_fe556acd0600dc8f.
		Shape = __obf_902526fcec0236d1.Overlay
	__obf_7aec6e82031fcbf1, __obf_a309197d4a7b9caa := __obf_4c68703ee0481a75.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_d7fe24b9cb955d9f,
		Mask:       __obf_902526fcec0236d1.Mask,
		Alpha:      __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Primary.Alpha,
		Width:      __obf_fe556acd0600dc8f.Width,
		Height:     __obf_fe556acd0600dc8f.Height,
		Block:      __obf_fe556acd0600dc8f,
	})

	if __obf_a309197d4a7b9caa != nil {
		return nil, __obf_a309197d4a7b9caa
	}

	if __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Type == types.MoveSlide {
		__obf_fe556acd0600dc8f.
			DY = __obf_fe556acd0600dc8f.Y
	} else {
		__obf_fe556acd0600dc8f.
			DY = __obf_3b80c8f4e178fd22.Y
	}
	__obf_fe556acd0600dc8f.
		DX = __obf_3b80c8f4e178fd22.X

	return &CaptchaData{
		Type:  __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Type,
		Block: __obf_fe556acd0600dc8f, __obf_aa601e1953c0218c: types.NewJPEGImage(__obf_aa601e1953c0218c), __obf_7aec6e82031fcbf1: types.NewPNGImage(__obf_7aec6e82031fcbf1),
	}, nil
}

func (__obf_4c68703ee0481a75 *Captcha) DrawWithNRGBA(__obf_517a9cf5c10dab97 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_a3ad588b3910ea0b := __obf_517a9cf5c10dab97.DrawBlocks
	__obf_21d04506c421a057 := helper.CreateNRGBACanvas(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, true)

	for _, __obf_7b306335dee6f50c := range __obf_a3ad588b3910ea0b {
		__obf_55b7dabab19f86c6, __obf_a309197d4a7b9caa := __obf_4c68703ee0481a75.__obf_bbf3e41264830eda(__obf_7b306335dee6f50c.Width, __obf_7b306335dee6f50c.Height, __obf_7b306335dee6f50c.Shape)
		if __obf_a309197d4a7b9caa != nil {
			return nil, nil, __obf_a309197d4a7b9caa
		}
		__obf_cccc043db02cc2a5 := __obf_55b7dabab19f86c6.Bounds()
		draw.Draw(__obf_21d04506c421a057, image.Rect(__obf_7b306335dee6f50c.X, __obf_7b306335dee6f50c.Y, __obf_7b306335dee6f50c.X+__obf_cccc043db02cc2a5.Dx(), __obf_7b306335dee6f50c.Y+__obf_cccc043db02cc2a5.Dy()), __obf_55b7dabab19f86c6, image.Point{}, draw.Over)
	}
	__obf_d02fe2bda0c79e40 := helper.CreateNRGBACanvas(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, true)
	if __obf_517a9cf5c10dab97.Background != nil {
		__obf_8033787aa30961f5 := __obf_517a9cf5c10dab97.Background
		__obf_7b306335dee6f50c := __obf_8033787aa30961f5.Bounds()
		__obf_a3cd9415fe3f47a0 := helper.CreateNRGBACanvas(__obf_7b306335dee6f50c.Dx(), __obf_7b306335dee6f50c.Dy(), true)
		__obf_4f7b186ab31ad6e6 := util.RangCutImagePos(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, __obf_8033787aa30961f5)
		draw.Draw(__obf_a3cd9415fe3f47a0, __obf_7b306335dee6f50c, __obf_8033787aa30961f5, __obf_4f7b186ab31ad6e6, draw.Src)
		__obf_a3cd9415fe3f47a0.
			SubImage(image.Rect(0, 0, __obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height))

		draw.Draw(__obf_d02fe2bda0c79e40, __obf_d02fe2bda0c79e40.Bounds(), __obf_a3cd9415fe3f47a0, image.Point{}, draw.Over)
		draw.Draw(__obf_a3cd9415fe3f47a0, __obf_21d04506c421a057.Bounds(), __obf_21d04506c421a057, image.Point{}, draw.Over)
		return __obf_a3cd9415fe3f47a0, __obf_d02fe2bda0c79e40, nil
	}

	return __obf_21d04506c421a057, __obf_d02fe2bda0c79e40, nil
}

func (__obf_4c68703ee0481a75 *Captcha) DrawWithTemplate(__obf_517a9cf5c10dab97 *DrawTplImageParam) (image.Image, error) {
	__obf_fe556acd0600dc8f := __obf_517a9cf5c10dab97.Block
	__obf_8033787aa30961f5 := __obf_517a9cf5c10dab97.Background
	__obf_21d04506c421a057 := helper.CreateNRGBACanvas(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, true)
	__obf_665d0de8b10a2151 := helper.CreateNRGBACanvas(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, true)
	__obf_4de63876aca69cc4, __obf_a309197d4a7b9caa := __obf_4c68703ee0481a75.__obf_bbf3e41264830eda(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, __obf_517a9cf5c10dab97.Mask)
	if __obf_a309197d4a7b9caa != nil {
		return nil, __obf_a309197d4a7b9caa
	}

	draw.Draw(__obf_665d0de8b10a2151, __obf_665d0de8b10a2151.Bounds(), __obf_8033787aa30961f5, image.Pt(__obf_fe556acd0600dc8f.X, __obf_fe556acd0600dc8f.Y), draw.Src)
	draw.DrawMask(__obf_21d04506c421a057, __obf_4de63876aca69cc4.Bounds(), __obf_665d0de8b10a2151, image.Point{}, __obf_4de63876aca69cc4, image.Point{}, draw.Over)
	__obf_00ab0af00d75493f, __obf_a309197d4a7b9caa := __obf_4c68703ee0481a75.__obf_bbf3e41264830eda(__obf_517a9cf5c10dab97.Width, __obf_517a9cf5c10dab97.Height, __obf_fe556acd0600dc8f.Shape)
	if __obf_a309197d4a7b9caa != nil {
		return nil, __obf_a309197d4a7b9caa
	}
	draw.Draw(__obf_21d04506c421a057, __obf_00ab0af00d75493f.Bounds(), __obf_00ab0af00d75493f, image.Point{}, draw.Over)

	return __obf_21d04506c421a057, nil
}

func (__obf_4c68703ee0481a75 *Captcha) __obf_bbf3e41264830eda(__obf_15297bf7371b5642, __obf_6342e9485709e3c8 int, __obf_805db6fe17b2384f image.Image) (*types.NRGBA, error) {
	__obf_21d04506c421a057 := helper.CreateNRGBACanvas(__obf_15297bf7371b5642, __obf_6342e9485709e3c8, true)
	draw.BiLinear.Scale(__obf_21d04506c421a057, __obf_21d04506c421a057.Bounds(), __obf_805db6fe17b2384f, __obf_805db6fe17b2384f.Bounds(), draw.Over, nil)
	return __obf_21d04506c421a057, nil
}

func (__obf_4c68703ee0481a75 *Captcha) __obf_38c3abd336ee398f(__obf_15297bf7371b5642, __obf_6342e9485709e3c8 int, __obf_57fc97dc1ab64b55 image.Image, __obf_a3ad588b3910ea0b []*types.Block) (image.Image, image.Image, error) {
	var __obf_74bc4a27c68c0be1 = make([]*types.Block, 0, len(__obf_a3ad588b3910ea0b))
	for _, __obf_127afded69753f40 := range __obf_a3ad588b3910ea0b {
		__obf_127afded69753f40.
			Shape = __obf_57fc97dc1ab64b55
		__obf_74bc4a27c68c0be1 = append(__obf_74bc4a27c68c0be1, __obf_127afded69753f40)
	}
	__obf_3e0becec2937672f, __obf_a309197d4a7b9caa := assets.PickBgImage()
	if __obf_a309197d4a7b9caa != nil {
		return nil, nil, __obf_a309197d4a7b9caa
	}

	return __obf_4c68703ee0481a75.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_15297bf7371b5642,
		Height:     __obf_6342e9485709e3c8,
		Background: __obf_3e0becec2937672f,
		Alpha:      __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Primary.Alpha,
		DrawBlocks: __obf_74bc4a27c68c0be1,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_4c68703ee0481a75 *Captcha) __obf_c395925c054b0d78() types.DeadZoneDirectionType {
	__obf_66e36c176b97bdef := __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Secondary.DeadZoneDirections
	__obf_6079f6833c73400b := rand.IntN(len(__obf_66e36c176b97bdef))
	if __obf_6079f6833c73400b < 0 {
		return 0
	}

	return __obf_66e36c176b97bdef[__obf_6079f6833c73400b]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_4c68703ee0481a75 *Captcha) __obf_09079f3cdb27f319() int {
	__obf_6dde67d9accff033 := __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Secondary.AnglePosRange
	__obf_6079f6833c73400b := rand.IntN(len(__obf_6dde67d9accff033))
	if __obf_6079f6833c73400b < 0 {
		return 0
	}
	__obf_da98df8a2d369e63 := __obf_6dde67d9accff033[__obf_6079f6833c73400b]
	return util.RandInt(__obf_da98df8a2d369e63.Min, __obf_da98df8a2d369e63.Max)
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
func (__obf_4c68703ee0481a75 *Captcha) __obf_5398960faf9a5da3(__obf_15297bf7371b5642, __obf_6342e9485709e3c8 int, __obf_d6f769e65c2e224b types.Range, __obf_9be91111e17045c3 int) ([]*types.Block, types.Point) {
	var __obf_a3ad588b3910ea0b = make([]*types.Block, 0, __obf_9be91111e17045c3)
	__obf_4c710309da48d413 := // width := imageSize.Width
		// height := imageSize.Height

		__obf_4c68703ee0481a75.__obf_09079f3cdb27f319()
	__obf_9590d379ed08c774 := util.RandInt(__obf_d6f769e65c2e224b.Min, __obf_d6f769e65c2e224b.Max)
	__obf_754a42fae967bd25 := __obf_9590d379ed08c774
	__obf_0900d13d2cc6a40b := __obf_9590d379ed08c774
	__obf_c287ce7debbccdfa := __obf_4c68703ee0481a75.__obf_c395925c054b0d78()
	__obf_bf550ea6026fa2ae := __obf_0900d13d2cc6a40b / 2
	__obf_d423f4bcf9d20919 := (__obf_15297bf7371b5642 - __obf_0900d13d2cc6a40b - 20) / __obf_9be91111e17045c3
	__obf_b15bed7a88bcf570 := __obf_4c68703ee0481a75.__obf_1b65216586c3a496(5, __obf_6342e9485709e3c8-__obf_754a42fae967bd25-5, __obf_754a42fae967bd25, __obf_c287ce7debbccdfa)

	for __obf_73157d9b2a1dafad := range __obf_9be91111e17045c3 {
		var __obf_fe556acd0600dc8f = &types.Block{}
		__obf_2e457f312b82e8cc, __obf_9f4cdb10a052dc0d := __obf_4c68703ee0481a75.__obf_2a0e9a54655b9559((__obf_73157d9b2a1dafad*__obf_d423f4bcf9d20919)+__obf_bf550ea6026fa2ae+5, ((__obf_73157d9b2a1dafad+1)*__obf_d423f4bcf9d20919)-__obf_bf550ea6026fa2ae, __obf_0900d13d2cc6a40b, __obf_c287ce7debbccdfa)
		__obf_2e457f312b82e8cc = int(max(float64(__obf_2e457f312b82e8cc), float64(__obf_bf550ea6026fa2ae+5)))
		__obf_fe556acd0600dc8f.
			X = util.RandInt(__obf_2e457f312b82e8cc+20, __obf_9f4cdb10a052dc0d+20) - __obf_bf550ea6026fa2ae

		if __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Secondary.EnableVerticalRandom {
			__obf_b15bed7a88bcf570 = __obf_4c68703ee0481a75.__obf_1b65216586c3a496(5, __obf_6342e9485709e3c8-__obf_754a42fae967bd25-5, __obf_754a42fae967bd25, __obf_c287ce7debbccdfa)
		}
		__obf_fe556acd0600dc8f.
			Y = __obf_b15bed7a88bcf570
		__obf_fe556acd0600dc8f.
			Width = __obf_0900d13d2cc6a40b
		__obf_fe556acd0600dc8f.
			Height = __obf_754a42fae967bd25
		__obf_fe556acd0600dc8f.
			Angle = __obf_4c710309da48d413
		__obf_a3ad588b3910ea0b = append(__obf_a3ad588b3910ea0b, __obf_fe556acd0600dc8f)
	}
	__obf_4f7b186ab31ad6e6 := types.Point{}
	if __obf_4c68703ee0481a75.__obf_bdac2ecb07870a35.Type == types.MoveSlide {
		__obf_4f7b186ab31ad6e6.
			X = util.RandInt(5, __obf_bf550ea6026fa2ae)
		__obf_4f7b186ab31ad6e6.
			Y = __obf_b15bed7a88bcf570
		return __obf_a3ad588b3910ea0b, __obf_4f7b186ab31ad6e6
	}

	switch __obf_c287ce7debbccdfa {
	case types.DeadZoneDirectionTypeTop:
		__obf_4f7b186ab31ad6e6.
			X = util.RandInt(5, __obf_15297bf7371b5642-__obf_0900d13d2cc6a40b-5)
		__obf_4f7b186ab31ad6e6.
			Y = 5
	case types.DeadZoneDirectionTypeBottom:
		__obf_4f7b186ab31ad6e6.
			X = util.RandInt(5, __obf_15297bf7371b5642-__obf_0900d13d2cc6a40b-5)
		__obf_4f7b186ab31ad6e6.
			Y = __obf_6342e9485709e3c8 - __obf_754a42fae967bd25 - 5
	case types.DeadZoneDirectionTypeLeft:
		__obf_4f7b186ab31ad6e6.
			X = 5
		__obf_4f7b186ab31ad6e6.
			Y = util.RandInt(5, __obf_6342e9485709e3c8-__obf_754a42fae967bd25-5)
	case types.DeadZoneDirectionTypeRight:
		__obf_4f7b186ab31ad6e6.
			X = __obf_15297bf7371b5642 - __obf_0900d13d2cc6a40b - 5
		__obf_4f7b186ab31ad6e6.
			Y = util.RandInt(5, __obf_6342e9485709e3c8-__obf_754a42fae967bd25-5)
	}

	return __obf_a3ad588b3910ea0b,

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
		__obf_4f7b186ab31ad6e6
}

func (__obf_4c68703ee0481a75 *Captcha) __obf_2a0e9a54655b9559(__obf_2e457f312b82e8cc, __obf_9f4cdb10a052dc0d, __obf_56b818b20b3b5113 int, __obf_c287ce7debbccdfa types.DeadZoneDirectionType) (int, int) {
	if __obf_c287ce7debbccdfa == types.DeadZoneDirectionTypeLeft {
		__obf_2e457f312b82e8cc += __obf_56b818b20b3b5113
		__obf_9f4cdb10a052dc0d += __obf_56b818b20b3b5113
	}
	return __obf_2e457f312b82e8cc,

		// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
		// params:
		//   - start: Start Y coordinate
		//   - end: End Y coordinate
		//   - value: types.Block height
		//   - dzdType: Dead zone direction
		//
		// return: Random Y coordinate
		__obf_9f4cdb10a052dc0d
}

func (__obf_4c68703ee0481a75 *Captcha) __obf_1b65216586c3a496(__obf_2e457f312b82e8cc, __obf_9f4cdb10a052dc0d, __obf_56b818b20b3b5113 int, __obf_c287ce7debbccdfa types.DeadZoneDirectionType) int {
	switch __obf_c287ce7debbccdfa {
	case types.DeadZoneDirectionTypeTop:
		__obf_2e457f312b82e8cc += __obf_56b818b20b3b5113
	case types.DeadZoneDirectionTypeBottom:
		__obf_9f4cdb10a052dc0d -= __obf_56b818b20b3b5113
	}
	return util.RandInt(__obf_2e457f312b82e8cc, __obf_9f4cdb10a052dc0d)
}
