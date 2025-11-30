package __obf_89d565d7ca4115dd

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
	__obf_bd64ea1c9065ee48 *types.SlideOption
}

var _ types.Captcha = (*Captcha)(nil)

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_bd64ea1c9065ee48 *types.SlideOption) *Captcha {
	if __obf_bd64ea1c9065ee48 == nil {
		return nil
	}
	return &Captcha{__obf_bd64ea1c9065ee48: __obf_bd64ea1c9065ee48}
}

// func (c *Captcha) Type() types.CaptchaType {
// 	if c.opt.Type == Move {
// 		return types.MoveCaptchat
// 	}
// 	return types.DragCaptchat
// }

func (__obf_3ad6454bef56219b *Captcha) Generate() (types.CaptchaData, error) {
	__obf_b55188ee496e244c,

		// slideType := randomSlideType(c.opt)
		__obf_8bf8c3d5ac659533 := assets.PickTile()
	if __obf_8bf8c3d5ac659533 != nil {
		return nil, __obf_8bf8c3d5ac659533
	}
	__obf_c37db36e2e4ed424 := util.RandInt(__obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Secondary.CountRange.Min, __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Secondary.CountRange.Max)
	__obf_3bd54e1702c0918d, __obf_81a47782a097029b := __obf_3ad6454bef56219b.__obf_5c9e0a9974feb5a7(__obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Primary.Size, __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Secondary.SizeRange, __obf_c37db36e2e4ed424)
	var __obf_3fc2dcf01757c488 *types.Block
	if len(__obf_3bd54e1702c0918d) > 1 {
		__obf_a6de0153f62b6db7 := max(rand.IntN(len(__obf_3bd54e1702c0918d)), 0)
		__obf_3fc2dcf01757c488 = __obf_3bd54e1702c0918d[__obf_a6de0153f62b6db7]
	} else {
		__obf_3fc2dcf01757c488 = __obf_3bd54e1702c0918d[0]
	}
	__obf_d2af1e67e2974419, __obf_989a0582eaeaba68, __obf_8bf8c3d5ac659533 := __obf_3ad6454bef56219b.__obf_799261db2346c17d(__obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Primary.Size, __obf_b55188ee496e244c.Shadow, __obf_3bd54e1702c0918d)
	if __obf_8bf8c3d5ac659533 != nil {
		return nil, __obf_8bf8c3d5ac659533
	}
	__obf_3fc2dcf01757c488.
		Shape = __obf_b55188ee496e244c.Overlay
	__obf_cfcc1a610d000093, __obf_8bf8c3d5ac659533 := __obf_3ad6454bef56219b.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_989a0582eaeaba68,
		Mask:       __obf_b55188ee496e244c.Mask,
		Alpha:      __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Primary.Alpha,
		Width:      __obf_3fc2dcf01757c488.Width,
		Height:     __obf_3fc2dcf01757c488.Height,
		Block:      __obf_3fc2dcf01757c488,
	})

	if __obf_8bf8c3d5ac659533 != nil {
		return nil, __obf_8bf8c3d5ac659533
	}

	if __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Type == types.MoveSlide {
		__obf_3fc2dcf01757c488.
			DY = __obf_3fc2dcf01757c488.Y
	} else {
		__obf_3fc2dcf01757c488.
			DY = __obf_81a47782a097029b.Y
	}
	__obf_3fc2dcf01757c488.
		DX = __obf_81a47782a097029b.X

	return &CaptchaData{__obf_8a104016d2195942: __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Type, __obf_3fc2dcf01757c488: __obf_3fc2dcf01757c488, __obf_d2af1e67e2974419: types.NewJPEGImage(__obf_d2af1e67e2974419), __obf_cfcc1a610d000093: types.NewPNGImage(__obf_cfcc1a610d000093)}, nil
}

func (__obf_3ad6454bef56219b *Captcha) DrawWithNRGBA(__obf_af01869838299192 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_3bd54e1702c0918d := __obf_af01869838299192.DrawBlocks
	__obf_79877608ee217751 := helper.CreateNRGBACanvas(__obf_af01869838299192.Width, __obf_af01869838299192.Height, true)

	for _, __obf_6b80eafe917f343b := range __obf_3bd54e1702c0918d {
		__obf_beb2393bee6e00d2, __obf_8bf8c3d5ac659533 := __obf_3ad6454bef56219b.__obf_f7f5616854b6e397(__obf_6b80eafe917f343b.Width, __obf_6b80eafe917f343b.Height, __obf_6b80eafe917f343b.Shape)
		if __obf_8bf8c3d5ac659533 != nil {
			return nil, nil, __obf_8bf8c3d5ac659533
		}
		__obf_77c56e517ef67219 := __obf_beb2393bee6e00d2.Bounds()
		draw.Draw(__obf_79877608ee217751, image.Rect(__obf_6b80eafe917f343b.X, __obf_6b80eafe917f343b.Y, __obf_6b80eafe917f343b.X+__obf_77c56e517ef67219.Dx(), __obf_6b80eafe917f343b.Y+__obf_77c56e517ef67219.Dy()), __obf_beb2393bee6e00d2, image.Point{}, draw.Over)
	}
	__obf_825db275a192e572 := helper.CreateNRGBACanvas(__obf_af01869838299192.Width, __obf_af01869838299192.Height, true)
	if __obf_af01869838299192.Background != nil {
		__obf_1ebdfd40b539a79b := __obf_af01869838299192.Background
		__obf_6b80eafe917f343b := __obf_1ebdfd40b539a79b.Bounds()
		__obf_e19fed2de697a6ce := helper.CreateNRGBACanvas(__obf_6b80eafe917f343b.Dx(), __obf_6b80eafe917f343b.Dy(), true)
		__obf_a7754654d7448f50 := util.RangCutImagePos(__obf_af01869838299192.Width, __obf_af01869838299192.Height, __obf_1ebdfd40b539a79b)
		draw.Draw(__obf_e19fed2de697a6ce, __obf_6b80eafe917f343b, __obf_1ebdfd40b539a79b, __obf_a7754654d7448f50, draw.Src)
		__obf_e19fed2de697a6ce.
			SubImage(image.Rect(0, 0, __obf_af01869838299192.Width, __obf_af01869838299192.Height))

		draw.Draw(__obf_825db275a192e572, __obf_825db275a192e572.Bounds(), __obf_e19fed2de697a6ce, image.Point{}, draw.Over)
		draw.Draw(__obf_e19fed2de697a6ce, __obf_79877608ee217751.Bounds(), __obf_79877608ee217751, image.Point{}, draw.Over)
		return __obf_e19fed2de697a6ce, __obf_825db275a192e572, nil
	}

	return __obf_79877608ee217751, __obf_825db275a192e572, nil
}

func (__obf_3ad6454bef56219b *Captcha) DrawWithTemplate(__obf_af01869838299192 *DrawTplImageParam) (image.Image, error) {
	__obf_3fc2dcf01757c488 := __obf_af01869838299192.Block
	__obf_1ebdfd40b539a79b := __obf_af01869838299192.Background
	__obf_79877608ee217751 := helper.CreateNRGBACanvas(__obf_af01869838299192.Width, __obf_af01869838299192.Height, true)
	__obf_eebccac9fb89a53b := helper.CreateNRGBACanvas(__obf_af01869838299192.Width, __obf_af01869838299192.Height, true)
	__obf_759e8ed8f5b5308d, __obf_8bf8c3d5ac659533 := __obf_3ad6454bef56219b.__obf_f7f5616854b6e397(__obf_af01869838299192.Width, __obf_af01869838299192.Height, __obf_af01869838299192.Mask)
	if __obf_8bf8c3d5ac659533 != nil {
		return nil, __obf_8bf8c3d5ac659533
	}

	draw.Draw(__obf_eebccac9fb89a53b, __obf_eebccac9fb89a53b.Bounds(), __obf_1ebdfd40b539a79b, image.Pt(__obf_3fc2dcf01757c488.X, __obf_3fc2dcf01757c488.Y), draw.Src)
	draw.DrawMask(__obf_79877608ee217751, __obf_759e8ed8f5b5308d.Bounds(), __obf_eebccac9fb89a53b, image.Point{}, __obf_759e8ed8f5b5308d, image.Point{}, draw.Over)
	__obf_622c3767192bf405, __obf_8bf8c3d5ac659533 := __obf_3ad6454bef56219b.__obf_f7f5616854b6e397(__obf_af01869838299192.Width, __obf_af01869838299192.Height, __obf_3fc2dcf01757c488.Shape)
	if __obf_8bf8c3d5ac659533 != nil {
		return nil, __obf_8bf8c3d5ac659533
	}
	draw.Draw(__obf_79877608ee217751, __obf_622c3767192bf405.Bounds(), __obf_622c3767192bf405, image.Point{}, draw.Over)

	return __obf_79877608ee217751, nil
}

func (__obf_3ad6454bef56219b *Captcha) __obf_f7f5616854b6e397(__obf_de1a129bd2a85889, __obf_ed1cf8aa04289204 int, __obf_be9048380398bd9f image.Image) (*types.NRGBA, error) {
	__obf_79877608ee217751 := helper.CreateNRGBACanvas(__obf_de1a129bd2a85889, __obf_ed1cf8aa04289204, true)
	draw.BiLinear.Scale(__obf_79877608ee217751, __obf_79877608ee217751.Bounds(), __obf_be9048380398bd9f, __obf_be9048380398bd9f.Bounds(), draw.Over, nil)
	return __obf_79877608ee217751, nil
}

func (__obf_3ad6454bef56219b *Captcha) __obf_799261db2346c17d(__obf_9fa4a2d8c7b00e86 types.Size, __obf_2a06c7cc4f6b4515 image.Image, __obf_3bd54e1702c0918d []*types.Block) (image.Image, image.Image, error) {
	var __obf_60a8e45128806291 = make([]*types.Block, 0, len(__obf_3bd54e1702c0918d))
	for _, __obf_6aeb45f0a725d990 := range __obf_3bd54e1702c0918d {
		__obf_6aeb45f0a725d990.
			Shape = __obf_2a06c7cc4f6b4515
		__obf_60a8e45128806291 = append(__obf_60a8e45128806291, __obf_6aeb45f0a725d990)
	}
	__obf_5e79753bdb2b7c80, __obf_8bf8c3d5ac659533 := assets.PickBgImage()
	if __obf_8bf8c3d5ac659533 != nil {
		return nil, nil, __obf_8bf8c3d5ac659533
	}

	return __obf_3ad6454bef56219b.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_9fa4a2d8c7b00e86.Width,
		Height:     __obf_9fa4a2d8c7b00e86.Height,
		Background: __obf_5e79753bdb2b7c80,
		Alpha:      __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Primary.Alpha,
		DrawBlocks: __obf_60a8e45128806291,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_3ad6454bef56219b *Captcha) __obf_ae69f32244686609() types.DeadZoneDirectionType {
	__obf_51fae6e74d2e26ac := __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Secondary.DeadZoneDirections
	__obf_a6de0153f62b6db7 := rand.IntN(len(__obf_51fae6e74d2e26ac))
	if __obf_a6de0153f62b6db7 < 0 {
		return 0
	}

	return __obf_51fae6e74d2e26ac[__obf_a6de0153f62b6db7]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_3ad6454bef56219b *Captcha) __obf_2f6f2725e23d8f3b() int {
	__obf_2f3c6a5191a1faf6 := __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Secondary.AnglePosRange
	__obf_a6de0153f62b6db7 := rand.IntN(len(__obf_2f3c6a5191a1faf6))
	if __obf_a6de0153f62b6db7 < 0 {
		return 0
	}
	__obf_dca837aa5d2acb1f := __obf_2f3c6a5191a1faf6[__obf_a6de0153f62b6db7]
	return util.RandInt(__obf_dca837aa5d2acb1f.Min, __obf_dca837aa5d2acb1f.Max)
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
func (__obf_3ad6454bef56219b *Captcha) __obf_5c9e0a9974feb5a7(__obf_52c6d059b07f3091 types.Size, __obf_9fa4a2d8c7b00e86 types.Range, __obf_e1837aa7d1aa0d65 int) ([]*types.Block, types.Point) {
	var __obf_3bd54e1702c0918d = make([]*types.Block, 0, __obf_e1837aa7d1aa0d65)
	__obf_de1a129bd2a85889 := __obf_52c6d059b07f3091.Width
	__obf_ed1cf8aa04289204 := __obf_52c6d059b07f3091.Height
	__obf_194c86d5ca2d848b := __obf_3ad6454bef56219b.__obf_2f6f2725e23d8f3b()
	__obf_ef455dbaffab3100 := util.RandInt(__obf_9fa4a2d8c7b00e86.Min, __obf_9fa4a2d8c7b00e86.Max)
	__obf_4b4de4602a117417 := __obf_ef455dbaffab3100
	__obf_fc0965a938577507 := __obf_ef455dbaffab3100
	__obf_478bdbe0799d92f1 := __obf_3ad6454bef56219b.__obf_ae69f32244686609()
	__obf_0780325d22665b2c := __obf_fc0965a938577507 / 2
	__obf_31953ff7b6490ee6 := (__obf_de1a129bd2a85889 - __obf_fc0965a938577507 - 20) / __obf_e1837aa7d1aa0d65
	__obf_8821dbe5b0603c39 := __obf_3ad6454bef56219b.__obf_dfa11664a1cb11c6(5, __obf_ed1cf8aa04289204-__obf_4b4de4602a117417-5, __obf_4b4de4602a117417, __obf_478bdbe0799d92f1)

	for __obf_086c2204b0227fe5 := range __obf_e1837aa7d1aa0d65 {
		var __obf_3fc2dcf01757c488 = &types.Block{}
		__obf_9841624fea8fb15d, __obf_f71520008c3e637e := __obf_3ad6454bef56219b.__obf_c01200571fc317f3((__obf_086c2204b0227fe5*__obf_31953ff7b6490ee6)+__obf_0780325d22665b2c+5, ((__obf_086c2204b0227fe5+1)*__obf_31953ff7b6490ee6)-__obf_0780325d22665b2c, __obf_fc0965a938577507, __obf_478bdbe0799d92f1)
		__obf_9841624fea8fb15d = int(max(float64(__obf_9841624fea8fb15d), float64(__obf_0780325d22665b2c+5)))
		__obf_3fc2dcf01757c488.
			X = util.RandInt(__obf_9841624fea8fb15d+20, __obf_f71520008c3e637e+20) - __obf_0780325d22665b2c

		if __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Secondary.EnableVerticalRandom {
			__obf_8821dbe5b0603c39 = __obf_3ad6454bef56219b.__obf_dfa11664a1cb11c6(5, __obf_ed1cf8aa04289204-__obf_4b4de4602a117417-5, __obf_4b4de4602a117417, __obf_478bdbe0799d92f1)
		}
		__obf_3fc2dcf01757c488.
			Y = __obf_8821dbe5b0603c39
		__obf_3fc2dcf01757c488.
			Width = __obf_fc0965a938577507
		__obf_3fc2dcf01757c488.
			Height = __obf_4b4de4602a117417
		__obf_3fc2dcf01757c488.
			Angle = __obf_194c86d5ca2d848b
		__obf_3bd54e1702c0918d = append(__obf_3bd54e1702c0918d, __obf_3fc2dcf01757c488)
	}
	__obf_a7754654d7448f50 := types.Point{}
	if __obf_3ad6454bef56219b.__obf_bd64ea1c9065ee48.Type == types.MoveSlide {
		__obf_a7754654d7448f50.
			X = util.RandInt(5, __obf_0780325d22665b2c)
		__obf_a7754654d7448f50.
			Y = __obf_8821dbe5b0603c39
		return __obf_3bd54e1702c0918d, __obf_a7754654d7448f50
	}

	switch __obf_478bdbe0799d92f1 {
	case types.DeadZoneDirectionTypeTop:
		__obf_a7754654d7448f50.
			X = util.RandInt(5, __obf_de1a129bd2a85889-__obf_fc0965a938577507-5)
		__obf_a7754654d7448f50.
			Y = 5
	case types.DeadZoneDirectionTypeBottom:
		__obf_a7754654d7448f50.
			X = util.RandInt(5, __obf_de1a129bd2a85889-__obf_fc0965a938577507-5)
		__obf_a7754654d7448f50.
			Y = __obf_ed1cf8aa04289204 - __obf_4b4de4602a117417 - 5
	case types.DeadZoneDirectionTypeLeft:
		__obf_a7754654d7448f50.
			X = 5
		__obf_a7754654d7448f50.
			Y = util.RandInt(5, __obf_ed1cf8aa04289204-__obf_4b4de4602a117417-5)
	case types.DeadZoneDirectionTypeRight:
		__obf_a7754654d7448f50.
			X = __obf_de1a129bd2a85889 - __obf_fc0965a938577507 - 5
		__obf_a7754654d7448f50.
			Y = util.RandInt(5, __obf_ed1cf8aa04289204-__obf_4b4de4602a117417-5)
	}

	return __obf_3bd54e1702c0918d,

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
		__obf_a7754654d7448f50
}

func (__obf_3ad6454bef56219b *Captcha) __obf_c01200571fc317f3(__obf_9841624fea8fb15d, __obf_f71520008c3e637e, __obf_e83c09f3234f027e int, __obf_478bdbe0799d92f1 types.DeadZoneDirectionType) (int, int) {
	if __obf_478bdbe0799d92f1 == types.DeadZoneDirectionTypeLeft {
		__obf_9841624fea8fb15d += __obf_e83c09f3234f027e
		__obf_f71520008c3e637e += __obf_e83c09f3234f027e
	}
	return __obf_9841624fea8fb15d,

		// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
		// params:
		//   - start: Start Y coordinate
		//   - end: End Y coordinate
		//   - value: types.Block height
		//   - dzdType: Dead zone direction
		//
		// return: Random Y coordinate
		__obf_f71520008c3e637e
}

func (__obf_3ad6454bef56219b *Captcha) __obf_dfa11664a1cb11c6(__obf_9841624fea8fb15d, __obf_f71520008c3e637e, __obf_e83c09f3234f027e int, __obf_478bdbe0799d92f1 types.DeadZoneDirectionType) int {
	switch __obf_478bdbe0799d92f1 {
	case types.DeadZoneDirectionTypeTop:
		__obf_9841624fea8fb15d += __obf_e83c09f3234f027e
	case types.DeadZoneDirectionTypeBottom:
		__obf_f71520008c3e637e -= __obf_e83c09f3234f027e
	}
	return util.RandInt(__obf_9841624fea8fb15d, __obf_f71520008c3e637e)
}
