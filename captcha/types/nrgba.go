package __obf_89363901d8df63bc

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_a89def74d931b834 image.Rectangle, __obf_7652d014b194c69b bool) *NRGBA {
	__obf_1572b176ff2c64b4 := image.NewNRGBA(__obf_a89def74d931b834)
	for __obf_e5fce044456d5b92 := 0; __obf_e5fce044456d5b92 < __obf_a89def74d931b834.Max.Y; __obf_e5fce044456d5b92++ {
		for __obf_4d548ce157fe2d7b := 0; __obf_4d548ce157fe2d7b < __obf_a89def74d931b834.Max.X; __obf_4d548ce157fe2d7b++ {
			if __obf_7652d014b194c69b {
				__obf_1572b176ff2c64b4.
					Set(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, color.Alpha{A: uint8(0)})
			} else {
				__obf_1572b176ff2c64b4.
					Set(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{__obf_1572b176ff2c64b4}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_3880035fe430c962 *NRGBA) DrawImage(__obf_7b74ec68f3d93f19 Palette, __obf_5f182ae83541ee1a *PositionRect, __obf_c8c4e84bdf0f9706 *AreaRect) {
	__obf_030657fbacf1ee00 := __obf_7b74ec68f3d93f19.Bounds().Max.X
	__obf_60d018dc204e7967 := __obf_7b74ec68f3d93f19.Bounds().Max.Y
	__obf_3876fa2304488c21 := __obf_5f182ae83541ee1a.X
	__obf_bb65587460878988 := __obf_5f182ae83541ee1a.Y
	__obf_98af21b948f83e91 := __obf_5f182ae83541ee1a.Height
	__obf_7a6598bfc4d9c3bd := __obf_c8c4e84bdf0f9706.MinX
	__obf_84fd4ac3ed71739e := __obf_c8c4e84bdf0f9706.MinY
	__obf_cfc8d3d9e1a8ca2e := __obf_c8c4e84bdf0f9706.MaxX
	__obf_f067434360bf931f := __obf_c8c4e84bdf0f9706.MaxY

	for __obf_4d548ce157fe2d7b := range __obf_030657fbacf1ee00 {
		for __obf_e5fce044456d5b92 := range __obf_60d018dc204e7967 {
			__obf_afee39acf6210256 := __obf_7b74ec68f3d93f19.At(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92)
			if _, _, _, __obf_dbe5c645f5262529 := __obf_afee39acf6210256.RGBA(); __obf_dbe5c645f5262529 > 0 {
				if __obf_4d548ce157fe2d7b >= __obf_7a6598bfc4d9c3bd && __obf_4d548ce157fe2d7b <= __obf_cfc8d3d9e1a8ca2e && __obf_e5fce044456d5b92 >= __obf_84fd4ac3ed71739e && __obf_e5fce044456d5b92 <= __obf_f067434360bf931f {
					__obf_3880035fe430c962.
						Set(__obf_3876fa2304488c21+(__obf_4d548ce157fe2d7b-__obf_7a6598bfc4d9c3bd), __obf_bb65587460878988-__obf_98af21b948f83e91+(__obf_e5fce044456d5b92-__obf_84fd4ac3ed71739e), __obf_7b74ec68f3d93f19.At(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_3880035fe430c962 *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_030657fbacf1ee00 := __obf_3880035fe430c962.Bounds().Max.X
	__obf_60d018dc204e7967 := __obf_3880035fe430c962.Bounds().Max.Y
	__obf_ef02f5aef30dfa3c := __obf_030657fbacf1ee00
	__obf_014a5112e2c2a925 := 0
	__obf_081a97e9258e5ee5 := __obf_60d018dc204e7967
	__obf_f14c1ac944ab90a0 := 0
	for __obf_4d548ce157fe2d7b := range __obf_030657fbacf1ee00 {
		for __obf_e5fce044456d5b92 := range __obf_60d018dc204e7967 {
			__obf_afee39acf6210256 := __obf_3880035fe430c962.At(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92)
			if _, _, _, __obf_dbe5c645f5262529 := __obf_afee39acf6210256.RGBA(); __obf_dbe5c645f5262529 > 0 {
				if __obf_4d548ce157fe2d7b < __obf_ef02f5aef30dfa3c {
					__obf_ef02f5aef30dfa3c = __obf_4d548ce157fe2d7b
				}
				if __obf_4d548ce157fe2d7b > __obf_014a5112e2c2a925 {
					__obf_014a5112e2c2a925 = __obf_4d548ce157fe2d7b
				}

				if __obf_e5fce044456d5b92 < __obf_081a97e9258e5ee5 {
					__obf_081a97e9258e5ee5 = __obf_e5fce044456d5b92
				}
				if __obf_e5fce044456d5b92 > __obf_f14c1ac944ab90a0 {
					__obf_f14c1ac944ab90a0 = __obf_e5fce044456d5b92
				}
			}
		}
	}
	__obf_ef02f5aef30dfa3c = int(max(0, float64(__obf_ef02f5aef30dfa3c-2)))
	__obf_014a5112e2c2a925 = int(min(float64(__obf_030657fbacf1ee00), float64(__obf_014a5112e2c2a925+2)))
	__obf_081a97e9258e5ee5 = int(max(0, float64(__obf_081a97e9258e5ee5-2)))
	__obf_f14c1ac944ab90a0 = int(min(float64(__obf_60d018dc204e7967), float64(__obf_f14c1ac944ab90a0+2)))

	return &AreaRect{__obf_ef02f5aef30dfa3c, __obf_014a5112e2c2a925,

		// Rotate rotates the canvas by any angle
		__obf_081a97e9258e5ee5, __obf_f14c1ac944ab90a0,
	}
}

func (__obf_3880035fe430c962 *NRGBA) Rotate(__obf_dbe5c645f5262529 int, __obf_8dba682b15bc5518 bool) {
	if __obf_dbe5c645f5262529 == 0 {
		return
	}
	__obf_1f7091190512fbb8 := float64(__obf_dbe5c645f5262529) * math.Pi / 180
	__obf_f7a9cd800b53cbec := __obf_3880035fe430c962.Bounds().Dx()
	__obf_eec5cd77836fab4d := __obf_3880035fe430c962.Bounds().Dy()
	__obf_672a1e7045b302e8, __obf_64febf9a6792142f := util.RotatedSize(__obf_f7a9cd800b53cbec, __obf_eec5cd77836fab4d, float64(__obf_dbe5c645f5262529))
	__obf_7b74ec68f3d93f19 := image.NewNRGBA(image.Rect(0, 0, __obf_672a1e7045b302e8, __obf_64febf9a6792142f))
	__obf_06a5f9c78594c262 := float64(__obf_672a1e7045b302e8) / 2
	__obf_08f9995043e9b3c8 := float64(__obf_64febf9a6792142f) / 2
	__obf_b70e2a73617caeb5 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_b70e2a73617caeb5 = __obf_b70e2a73617caeb5.Translate(__obf_06a5f9c78594c262, __obf_08f9995043e9b3c8)
	__obf_b70e2a73617caeb5 = __obf_b70e2a73617caeb5.Rotate(__obf_1f7091190512fbb8)
	__obf_b70e2a73617caeb5 = __obf_b70e2a73617caeb5.Translate(-__obf_06a5f9c78594c262, -__obf_08f9995043e9b3c8)
	__obf_4d548ce157fe2d7b := (__obf_672a1e7045b302e8 - __obf_3880035fe430c962.Bounds().Dx()) / 2
	__obf_e5fce044456d5b92 := (__obf_64febf9a6792142f - __obf_3880035fe430c962.Bounds().Dy()) / 2
	__obf_106240c091975608, __obf_831ca242f32d0b21 := float64(__obf_4d548ce157fe2d7b), float64(__obf_e5fce044456d5b92)
	__obf_0aa14e5826fd7e47 := __obf_b70e2a73617caeb5.Translate(__obf_106240c091975608, __obf_831ca242f32d0b21)
	__obf_0415988496f12b6c := f64.Aff3{__obf_0aa14e5826fd7e47.XX, __obf_0aa14e5826fd7e47.XY, __obf_0aa14e5826fd7e47.X0, __obf_0aa14e5826fd7e47.YX, __obf_0aa14e5826fd7e47.YY, __obf_0aa14e5826fd7e47.Y0}

	draw.BiLinear.Transform(__obf_7b74ec68f3d93f19, __obf_0415988496f12b6c, __obf_3880035fe430c962, __obf_3880035fe430c962.Bounds(), draw.Over, nil)
	__obf_3880035fe430c962.
		NRGBA = __obf_7b74ec68f3d93f19

	if __obf_8dba682b15bc5518 {
		__obf_bb54e9efc8da8422 := __obf_672a1e7045b302e8 - __obf_f7a9cd800b53cbec
		__obf_338dfc25ad90027d := __obf_64febf9a6792142f - __obf_eec5cd77836fab4d
		__obf_0996b35bb1f98ea1 := (__obf_bb54e9efc8da8422 / 2) + 1
		__obf_68d36fd00e636e53 := (__obf_338dfc25ad90027d / 2) + 1
		__obf_3880035fe430c962.
			SubImage(image.Rect(__obf_0996b35bb1f98ea1, __obf_68d36fd00e636e53,

				// CropCircle crops a circular area
				__obf_f7a9cd800b53cbec+__obf_0996b35bb1f98ea1, __obf_eec5cd77836fab4d+__obf_68d36fd00e636e53))
	}
}

func (__obf_3880035fe430c962 *NRGBA) CropCircle(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_126525ac437289c4 int) {
	__obf_7c33e88f7a9b618d := __obf_3880035fe430c962.Bounds()
	__obf_121f61fe2db52ca4 := image.NewNRGBA(__obf_7c33e88f7a9b618d)
	for __obf_13f11059123c69ac := __obf_7c33e88f7a9b618d.Min.Y; __obf_13f11059123c69ac < __obf_7c33e88f7a9b618d.Max.Y; __obf_13f11059123c69ac++ {
		for __obf_1881af43861cfd2c := __obf_7c33e88f7a9b618d.Min.X; __obf_1881af43861cfd2c < __obf_7c33e88f7a9b618d.Max.X; __obf_1881af43861cfd2c++ {
			__obf_99649129705a2db4 := math.Hypot(float64(__obf_1881af43861cfd2c-__obf_4d548ce157fe2d7b), float64(__obf_13f11059123c69ac-__obf_e5fce044456d5b92))
			if __obf_99649129705a2db4 <= float64(__obf_126525ac437289c4) {
				__obf_121f61fe2db52ca4.
					Set(__obf_1881af43861cfd2c, __obf_13f11059123c69ac, color.White)
			} else {
				__obf_121f61fe2db52ca4.
					Set(__obf_1881af43861cfd2c, __obf_13f11059123c69ac, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_121f61fe2db52ca4, __obf_121f61fe2db52ca4.Bounds(), __obf_3880035fe430c962, image.Point{X: 0, Y: 0}, __obf_121f61fe2db52ca4, image.Point{}, draw.Over)
	__obf_3880035fe430c962.
		NRGBA = __obf_121f61fe2db52ca4
}

// CropScaleCircle scales and crops a circular area
func (__obf_3880035fe430c962 *NRGBA) CropScaleCircle(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_126525ac437289c4 int, __obf_f3621a5c951e3bd7 int) {
	__obf_7c33e88f7a9b618d := __obf_3880035fe430c962.Bounds()
	__obf_121f61fe2db52ca4 := image.NewNRGBA(__obf_7c33e88f7a9b618d)

	for __obf_13f11059123c69ac := __obf_7c33e88f7a9b618d.Min.Y; __obf_13f11059123c69ac < __obf_7c33e88f7a9b618d.Max.Y; __obf_13f11059123c69ac++ {
		for __obf_1881af43861cfd2c := __obf_7c33e88f7a9b618d.Min.X; __obf_1881af43861cfd2c < __obf_7c33e88f7a9b618d.Max.X; __obf_1881af43861cfd2c++ {
			__obf_99649129705a2db4 := math.Hypot(float64(__obf_1881af43861cfd2c-__obf_4d548ce157fe2d7b), float64(__obf_13f11059123c69ac-__obf_e5fce044456d5b92))
			if __obf_99649129705a2db4 <= float64(__obf_126525ac437289c4) {
				__obf_121f61fe2db52ca4.
					Set(__obf_1881af43861cfd2c, __obf_13f11059123c69ac, color.White)
			} else {
				__obf_121f61fe2db52ca4.
					Set(__obf_1881af43861cfd2c, __obf_13f11059123c69ac, color.Transparent)
			}
		}
	}

	if __obf_f3621a5c951e3bd7 > 0 {
		__obf_219b6d907d7331ac := __obf_f3621a5c951e3bd7 * 2
		__obf_b53a4122fb6cbe37 := image.NewNRGBA(image.Rect(0, 0, __obf_3880035fe430c962.Bounds().Dx()-__obf_219b6d907d7331ac, __obf_3880035fe430c962.Bounds().Dy()-__obf_219b6d907d7331ac))
		draw.BiLinear.Scale(__obf_b53a4122fb6cbe37, __obf_b53a4122fb6cbe37.Bounds(), __obf_121f61fe2db52ca4, __obf_121f61fe2db52ca4.Bounds(), draw.Over, nil)
		__obf_121f61fe2db52ca4 = __obf_b53a4122fb6cbe37
	}

	draw.DrawMask(__obf_121f61fe2db52ca4, __obf_121f61fe2db52ca4.Bounds(), __obf_3880035fe430c962, image.Point{X: __obf_f3621a5c951e3bd7, Y: __obf_f3621a5c951e3bd7}, __obf_121f61fe2db52ca4, image.Point{}, draw.Over)
	__obf_3880035fe430c962.
		NRGBA = __obf_121f61fe2db52ca4
}

// Scale scales the canvas
func (__obf_3880035fe430c962 *NRGBA) Scale(__obf_f3621a5c951e3bd7 int, __obf_5ca4f6cafdd9fc70, __obf_c92e82cada653444 bool) {
	__obf_7b74ec68f3d93f19 := __obf_3880035fe430c962.NRGBA
	if __obf_f3621a5c951e3bd7 > 0 {
		__obf_219b6d907d7331ac := __obf_f3621a5c951e3bd7 * 2
		__obf_c18e022448568363 := __obf_3880035fe430c962.Bounds().Dx() - __obf_219b6d907d7331ac
		__obf_e0147a281b4ad2ef := __obf_3880035fe430c962.Bounds().Dy() - __obf_219b6d907d7331ac
		__obf_4477172c3a7d6449 := image.NewNRGBA(image.Rect(0, 0, __obf_c18e022448568363, __obf_e0147a281b4ad2ef))

		if !__obf_5ca4f6cafdd9fc70 {
			draw.BiLinear.Scale(__obf_4477172c3a7d6449, __obf_4477172c3a7d6449.Bounds(), __obf_3880035fe430c962, __obf_3880035fe430c962.Bounds(), draw.Over, nil)
		} else {
			__obf_07cf0ac50be1f0f9 := util.CalcResizedRect(__obf_3880035fe430c962.Bounds(), __obf_c18e022448568363, __obf_e0147a281b4ad2ef, __obf_c92e82cada653444)
			draw.ApproxBiLinear.Scale(__obf_4477172c3a7d6449, __obf_07cf0ac50be1f0f9.Bounds(), __obf_3880035fe430c962, __obf_3880035fe430c962.Bounds(), draw.Over, nil)
		}
		__obf_7b74ec68f3d93f19 = __obf_4477172c3a7d6449
	}
	__obf_3880035fe430c962.
		NRGBA = __obf_7b74ec68f3d93f19
}
