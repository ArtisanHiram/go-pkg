package __obf_4612facabc6519c1

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
	__obf_00ec7f50dc60e397 *types.SlideOption
}

var _ types.Captcha = (*Captcha)(nil)

// NewCaptcha creates a new slide CAPTCHA instance in basic mode
func NewCaptcha(__obf_00ec7f50dc60e397 *types.SlideOption) *Captcha {
	if __obf_00ec7f50dc60e397 == nil {
		return nil
	}
	return &Captcha{__obf_00ec7f50dc60e397: __obf_00ec7f50dc60e397}
}

// func (c *Captcha) Type() types.CaptchaType {
// 	if c.opt.Type == Move {
// 		return types.MoveCaptchat
// 	}
// 	return types.DragCaptchat
// }

func (__obf_7961f70af277ccd5 *Captcha) Generate() (types.CaptchaData, error) {
	__obf_fb96d6fd062b4eee,

		// slideType := randomSlideType(c.opt)
		__obf_aaa460b2b1870a20 := assets.PickTile()
	if __obf_aaa460b2b1870a20 != nil {
		return nil, __obf_aaa460b2b1870a20
	}
	__obf_6a61285baaee9bf6 := util.RandInt(__obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Secondary.CountRange.Min, __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Secondary.CountRange.Max)
	__obf_0d0ca47d489f1890, _ := __obf_7961f70af277ccd5.__obf_2081b6d17250ae1c(__obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Primary.Width, __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Primary.Height, __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Secondary.SizeRange, __obf_6a61285baaee9bf6)
	var __obf_cde5616243b3a32c *types.Block
	if len(__obf_0d0ca47d489f1890) > 1 {
		__obf_84de90671371425f := rand.IntN(len(__obf_0d0ca47d489f1890))
		__obf_cde5616243b3a32c = __obf_0d0ca47d489f1890[__obf_84de90671371425f]
	} else {
		__obf_cde5616243b3a32c = __obf_0d0ca47d489f1890[0]
	}
	__obf_ee6bc09cdda2f9e2, __obf_3782b9531be5eeed, __obf_aaa460b2b1870a20 := __obf_7961f70af277ccd5.__obf_99c67c9210b1b5af(__obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Primary.Width, __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Primary.Height, __obf_fb96d6fd062b4eee.Shadow, __obf_0d0ca47d489f1890)
	if __obf_aaa460b2b1870a20 != nil {
		return nil, __obf_aaa460b2b1870a20
	}
	__obf_cde5616243b3a32c.
		Shape = __obf_fb96d6fd062b4eee.Overlay
	__obf_8c3d473aadb0f39a, __obf_aaa460b2b1870a20 := __obf_7961f70af277ccd5.DrawWithTemplate(&DrawTplImageParam{
		Background: __obf_3782b9531be5eeed,
		Mask:       __obf_fb96d6fd062b4eee.Mask,
		Alpha:      __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Primary.Alpha,
		Width:      __obf_cde5616243b3a32c.Width,
		Height:     __obf_cde5616243b3a32c.Height,
		Block:      __obf_cde5616243b3a32c,
	})

	return &CaptchaData{__obf_3535230d19dfb383: &types.Point{
		X: __obf_cde5616243b3a32c.X - __obf_cde5616243b3a32c.Width/2,
		Y: __obf_cde5616243b3a32c.Y - __obf_cde5616243b3a32c.Height/2,
	}, __obf_ee6bc09cdda2f9e2: types.NewJPEGImage(__obf_ee6bc09cdda2f9e2), __obf_8c3d473aadb0f39a: types.NewPNGImage(__obf_8c3d473aadb0f39a),
	}, __obf_aaa460b2b1870a20
}

func (__obf_7961f70af277ccd5 *Captcha) DrawWithNRGBA(__obf_a63fdf4d2a93ccd2 *DrawImageParam) (image.Image, image.Image, error) {
	__obf_0d0ca47d489f1890 := __obf_a63fdf4d2a93ccd2.DrawBlocks
	__obf_37076b2020878136 := helper.CreateNRGBACanvas(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, true)

	for _, __obf_4caf60b8bfc0f30c := range __obf_0d0ca47d489f1890 {
		__obf_ef49f6e24311dacd, __obf_aaa460b2b1870a20 := __obf_7961f70af277ccd5.__obf_fea289e91529b1bd(__obf_4caf60b8bfc0f30c.Width, __obf_4caf60b8bfc0f30c.Height, __obf_4caf60b8bfc0f30c.Shape)
		if __obf_aaa460b2b1870a20 != nil {
			return nil, nil, __obf_aaa460b2b1870a20
		}
		__obf_22c76c91f304b60c := __obf_ef49f6e24311dacd.Bounds()
		__obf_109b4cd29a66a7e2 := // draw.Draw(cvs, image.Rect(b.X, b.Y, b.X+graphBounds.Dx(), b.Y+graphBounds.Dy()), graphImage, image.Point{}, draw.Over)
			__obf_4caf60b8bfc0f30c.X - __obf_22c76c91f304b60c.Dx()/2
		__obf_c5c1ef6d8b4948a6 := __obf_4caf60b8bfc0f30c.Y - __obf_22c76c91f304b60c.Dy()/2
		draw.Draw(__obf_37076b2020878136, image.Rect(__obf_109b4cd29a66a7e2, __obf_c5c1ef6d8b4948a6, __obf_109b4cd29a66a7e2+__obf_22c76c91f304b60c.Dx(), __obf_c5c1ef6d8b4948a6+__obf_22c76c91f304b60c.Dy()), __obf_ef49f6e24311dacd, image.Point{}, draw.Over)
	}
	__obf_1348a10a87ccaff6 := helper.CreateNRGBACanvas(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, true)
	if __obf_a63fdf4d2a93ccd2.Background != nil {
		__obf_55c82494c02b5b24 := __obf_a63fdf4d2a93ccd2.Background
		__obf_4caf60b8bfc0f30c := __obf_55c82494c02b5b24.Bounds()
		__obf_8e10828d3b8dd960 := helper.CreateNRGBACanvas(__obf_4caf60b8bfc0f30c.Dx(), __obf_4caf60b8bfc0f30c.Dy(), true)
		__obf_d9915341ded0d1d8 := util.RangCutImagePos(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, __obf_55c82494c02b5b24)
		draw.Draw(__obf_8e10828d3b8dd960, __obf_4caf60b8bfc0f30c, __obf_55c82494c02b5b24, __obf_d9915341ded0d1d8, draw.Src)
		__obf_8e10828d3b8dd960.
			SubImage(image.Rect(0, 0, __obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height))

		draw.Draw(__obf_1348a10a87ccaff6, __obf_1348a10a87ccaff6.Bounds(), __obf_8e10828d3b8dd960, image.Point{}, draw.Over)
		draw.Draw(__obf_8e10828d3b8dd960, __obf_37076b2020878136.Bounds(), __obf_37076b2020878136, image.Point{}, draw.Over)
		return __obf_8e10828d3b8dd960, __obf_1348a10a87ccaff6, nil
	}

	return __obf_37076b2020878136, __obf_1348a10a87ccaff6, nil
}

func (__obf_7961f70af277ccd5 *Captcha) DrawWithTemplate(__obf_a63fdf4d2a93ccd2 *DrawTplImageParam) (image.Image, error) {
	__obf_cde5616243b3a32c := __obf_a63fdf4d2a93ccd2.Block
	__obf_55c82494c02b5b24 := __obf_a63fdf4d2a93ccd2.Background
	__obf_37076b2020878136 := helper.CreateNRGBACanvas(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, true)
	__obf_1dda5c3a5bcd30c6 := helper.CreateNRGBACanvas(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, true)
	__obf_f6c4785a365ebfa3, __obf_aaa460b2b1870a20 := __obf_7961f70af277ccd5.__obf_fea289e91529b1bd(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, __obf_a63fdf4d2a93ccd2.Mask)
	if __obf_aaa460b2b1870a20 != nil {
		return nil, __obf_aaa460b2b1870a20
	}
	__obf_109b4cd29a66a7e2 := // draw.Draw(bgCvs, bgCvs.Bounds(), bgImage, image.Pt(block.X, block.Y), draw.Src)
		__obf_cde5616243b3a32c.X - __obf_a63fdf4d2a93ccd2.Width/2
	__obf_c5c1ef6d8b4948a6 := __obf_cde5616243b3a32c.Y - __obf_a63fdf4d2a93ccd2.Height/2
	draw.Draw(__obf_1dda5c3a5bcd30c6, __obf_1dda5c3a5bcd30c6.Bounds(), __obf_55c82494c02b5b24, image.Pt(__obf_109b4cd29a66a7e2, __obf_c5c1ef6d8b4948a6), draw.Src)

	draw.DrawMask(__obf_37076b2020878136, __obf_f6c4785a365ebfa3.Bounds(), __obf_1dda5c3a5bcd30c6, image.Point{}, __obf_f6c4785a365ebfa3, image.Point{}, draw.Over)
	__obf_011a7ab29e07de77, __obf_aaa460b2b1870a20 := __obf_7961f70af277ccd5.__obf_fea289e91529b1bd(__obf_a63fdf4d2a93ccd2.Width, __obf_a63fdf4d2a93ccd2.Height, __obf_cde5616243b3a32c.Shape)
	if __obf_aaa460b2b1870a20 != nil {
		return nil, __obf_aaa460b2b1870a20
	}
	draw.Draw(__obf_37076b2020878136, __obf_011a7ab29e07de77.Bounds(), __obf_011a7ab29e07de77, image.Point{}, draw.Over)

	return __obf_37076b2020878136, nil
}

func (__obf_7961f70af277ccd5 *Captcha) __obf_fea289e91529b1bd(__obf_0d19bb9b424b014c, __obf_7110ecae7efc9ea7 int, __obf_bbdac8678c52b695 image.Image) (*types.NRGBA, error) {
	__obf_37076b2020878136 := helper.CreateNRGBACanvas(__obf_0d19bb9b424b014c, __obf_7110ecae7efc9ea7, true)
	draw.BiLinear.Scale(__obf_37076b2020878136, __obf_37076b2020878136.Bounds(), __obf_bbdac8678c52b695, __obf_bbdac8678c52b695.Bounds(), draw.Over, nil)
	return __obf_37076b2020878136, nil
}

func (__obf_7961f70af277ccd5 *Captcha) __obf_99c67c9210b1b5af(__obf_0d19bb9b424b014c, __obf_7110ecae7efc9ea7 int, __obf_bd0ed8a422ae6ff7 image.Image, __obf_0d0ca47d489f1890 []*types.Block) (image.Image, image.Image, error) {
	var __obf_224d1ff17adbc3e3 = make([]*types.Block, 0, len(__obf_0d0ca47d489f1890))
	for _, __obf_38a8d8b6e0092e16 := range __obf_0d0ca47d489f1890 {
		__obf_38a8d8b6e0092e16.
			Shape = __obf_bd0ed8a422ae6ff7
		__obf_224d1ff17adbc3e3 = append(__obf_224d1ff17adbc3e3, __obf_38a8d8b6e0092e16)
	}
	__obf_e0adeeafc79f24fd, __obf_aaa460b2b1870a20 := assets.PickBgImage()
	if __obf_aaa460b2b1870a20 != nil {
		return nil, nil, __obf_aaa460b2b1870a20
	}

	return __obf_7961f70af277ccd5.DrawWithNRGBA(&DrawImageParam{
		Width:      __obf_0d19bb9b424b014c,
		Height:     __obf_7110ecae7efc9ea7,
		Background: __obf_e0adeeafc79f24fd,
		Alpha:      __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Primary.Alpha,
		DrawBlocks: __obf_224d1ff17adbc3e3,
	})
}

// randDeadZoneDirection generates a random dead zone direction
// return: Dead zone direction
func (__obf_7961f70af277ccd5 *Captcha) __obf_b8b90d082dd8de4b() types.DeadZoneDirectionType {
	__obf_e06beb20952cc565 := __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Secondary.DeadZoneDirections
	__obf_84de90671371425f := rand.IntN(len(__obf_e06beb20952cc565))
	if __obf_84de90671371425f < 0 {
		return 0
	}

	return __obf_e06beb20952cc565[__obf_84de90671371425f]
}

// randGraphAngle generates a random graph angle
// return: Random angle value
func (__obf_7961f70af277ccd5 *Captcha) __obf_58ce8ddec257a0b9() int {
	__obf_16b4b17e095a5eab := __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Secondary.AnglePosRange
	__obf_84de90671371425f := rand.IntN(len(__obf_16b4b17e095a5eab))
	if __obf_84de90671371425f < 0 {
		return 0
	}
	__obf_40b9eb7043219cc4 := __obf_16b4b17e095a5eab[__obf_84de90671371425f]
	return util.RandInt(__obf_40b9eb7043219cc4.Min, __obf_40b9eb7043219cc4.Max)
}

// genGraphBlocks generates graph block data
// params:
//   - imageSize: Main image size
//   - size: Graph size range
//   - length: Number of graphs
//
// returns:
//   - []*types.Block: List of blocks
//   - *types.Point: 滑块的初始安全位置（起点），通常用于前端控制滑块按钮一开始出现在哪里，或者哪块区域是绝对不会出现拼图缺口的“死区”
func (__obf_7961f70af277ccd5 *Captcha) __obf_2081b6d17250ae1c(__obf_0d19bb9b424b014c, __obf_7110ecae7efc9ea7 int, __obf_3c91830d951bf67d types.Range, __obf_2c95cbd020cea5a6 int) ([]*types.Block, *types.Point) {
	var __obf_0d0ca47d489f1890 = make([]*types.Block, 0, __obf_2c95cbd020cea5a6)
	__obf_489257751067f547 := // width := imageSize.Width
		// height := imageSize.Height

		__obf_7961f70af277ccd5.__obf_58ce8ddec257a0b9()
	__obf_d7eb008eac8b7684 := util.RandInt(__obf_3c91830d951bf67d.Min, __obf_3c91830d951bf67d.Max)
	__obf_eabda8184036fb66 := __obf_d7eb008eac8b7684
	__obf_8bf1020c8e9fa499 := __obf_d7eb008eac8b7684
	__obf_5645cefdeb562e81 := __obf_7961f70af277ccd5.__obf_b8b90d082dd8de4b()
	__obf_b8bd28e9bc3e1dc9 := __obf_8bf1020c8e9fa499 / 2
	__obf_7e14f5ef0d3cbf0c := (__obf_0d19bb9b424b014c - __obf_8bf1020c8e9fa499 - 20) / __obf_2c95cbd020cea5a6
	__obf_d9195fd5c9431986 := __obf_7961f70af277ccd5.__obf_b3fc36523cb4a8c6(5, __obf_7110ecae7efc9ea7-__obf_eabda8184036fb66-5, __obf_eabda8184036fb66, __obf_5645cefdeb562e81)

	for __obf_efc85c789e3273c2 := range __obf_2c95cbd020cea5a6 {
		var __obf_cde5616243b3a32c = &types.Block{}
		__obf_9721d25a915d7363, __obf_73bd6469640b79c0 := __obf_7961f70af277ccd5.__obf_b6677998301555cd((__obf_efc85c789e3273c2*__obf_7e14f5ef0d3cbf0c)+__obf_8bf1020c8e9fa499/2+5, ((__obf_efc85c789e3273c2+1)*__obf_7e14f5ef0d3cbf0c)-__obf_8bf1020c8e9fa499/2, __obf_8bf1020c8e9fa499, __obf_5645cefdeb562e81)
		__obf_9721d25a915d7363 = int(max(float64(__obf_9721d25a915d7363), float64(__obf_b8bd28e9bc3e1dc9+5)))
		__obf_cde5616243b3a32c.
			X = util.RandInt(__obf_9721d25a915d7363+20, __obf_73bd6469640b79c0+20)

		if __obf_7961f70af277ccd5.__obf_00ec7f50dc60e397.Secondary.EnableVerticalRandom {
			__obf_d9195fd5c9431986 = __obf_7961f70af277ccd5.__obf_b3fc36523cb4a8c6(5, __obf_7110ecae7efc9ea7-__obf_eabda8184036fb66-5, __obf_eabda8184036fb66, __obf_5645cefdeb562e81)
		}
		__obf_cde5616243b3a32c.
			Y = __obf_d9195fd5c9431986 + __obf_eabda8184036fb66/2
		__obf_cde5616243b3a32c.
			Width = __obf_8bf1020c8e9fa499
		__obf_cde5616243b3a32c.
			Height = __obf_eabda8184036fb66
		__obf_cde5616243b3a32c.
			Angle = __obf_489257751067f547
		__obf_0d0ca47d489f1890 = append(__obf_0d0ca47d489f1890, __obf_cde5616243b3a32c)
	}
	__obf_d9915341ded0d1d8 := &types.Point{}

	switch __obf_5645cefdeb562e81 {
	case types.DeadZoneDirectionTypeTop:
		__obf_d9915341ded0d1d8.
			X = util.RandInt(__obf_8bf1020c8e9fa499/2+5, __obf_0d19bb9b424b014c-__obf_8bf1020c8e9fa499/2-5)
		__obf_d9915341ded0d1d8.
			Y = __obf_eabda8184036fb66/2 + 5
	case types.DeadZoneDirectionTypeBottom:
		__obf_d9915341ded0d1d8.
			X = util.RandInt(__obf_8bf1020c8e9fa499/2+5, __obf_0d19bb9b424b014c-__obf_8bf1020c8e9fa499/2-5)
		__obf_d9915341ded0d1d8.
			Y = __obf_7110ecae7efc9ea7 - __obf_eabda8184036fb66/2 - 5
	case types.DeadZoneDirectionTypeLeft:
		__obf_d9915341ded0d1d8.
			X = __obf_8bf1020c8e9fa499/2 + 5
		__obf_d9915341ded0d1d8.
			Y = util.RandInt(__obf_eabda8184036fb66/2+5, __obf_7110ecae7efc9ea7-__obf_eabda8184036fb66/2-5)
	case types.DeadZoneDirectionTypeRight:
		__obf_d9915341ded0d1d8.
			X = __obf_0d19bb9b424b014c - __obf_8bf1020c8e9fa499/2 - 5
		__obf_d9915341ded0d1d8.
			Y = util.RandInt(__obf_eabda8184036fb66/2+5, __obf_7110ecae7efc9ea7-__obf_eabda8184036fb66/2-5)
	}

	return __obf_0d0ca47d489f1890,

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
		__obf_d9915341ded0d1d8
}

func (__obf_7961f70af277ccd5 *Captcha) __obf_b6677998301555cd(__obf_9721d25a915d7363, __obf_73bd6469640b79c0, __obf_41a12f2d93300e96 int, __obf_5645cefdeb562e81 types.DeadZoneDirectionType) (int, int) {
	if __obf_5645cefdeb562e81 == types.DeadZoneDirectionTypeLeft {
		__obf_9721d25a915d7363 += __obf_41a12f2d93300e96
		__obf_73bd6469640b79c0 += __obf_41a12f2d93300e96
	}
	return __obf_9721d25a915d7363,

		// calcYWithDeadZone calculates the Y coordinate (considering dead zone)
		// params:
		//   - start: Start Y coordinate
		//   - end: End Y coordinate
		//   - value: types.Block height
		//   - dzdType: Dead zone direction
		//
		// return: Random Y coordinate
		__obf_73bd6469640b79c0
}

func (__obf_7961f70af277ccd5 *Captcha) __obf_b3fc36523cb4a8c6(__obf_9721d25a915d7363, __obf_73bd6469640b79c0, __obf_41a12f2d93300e96 int, __obf_5645cefdeb562e81 types.DeadZoneDirectionType) int {
	switch __obf_5645cefdeb562e81 {
	case types.DeadZoneDirectionTypeTop:
		__obf_9721d25a915d7363 += __obf_41a12f2d93300e96
	case types.DeadZoneDirectionTypeBottom:
		__obf_73bd6469640b79c0 -= __obf_41a12f2d93300e96
	}
	return util.RandInt(__obf_9721d25a915d7363, __obf_73bd6469640b79c0)
}
