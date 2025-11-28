package __obf_deb7e65d40f46bd0

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/v2/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_941de68f67190911 image.Rectangle, __obf_e177f4c31b84dee8 bool) *NRGBA {
	__obf_8d077f7cd507d18d := image.NewNRGBA(__obf_941de68f67190911)
	for __obf_7a88a48a3d7db1aa := 0; __obf_7a88a48a3d7db1aa < __obf_941de68f67190911.Max.Y; __obf_7a88a48a3d7db1aa++ {
		for __obf_ce61f234cbe42808 := 0; __obf_ce61f234cbe42808 < __obf_941de68f67190911.Max.X; __obf_ce61f234cbe42808++ {
			if __obf_e177f4c31b84dee8 {
				__obf_8d077f7cd507d18d.Set(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, color.Alpha{A: uint8(0)})
			} else {
				__obf_8d077f7cd507d18d.Set(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{
		__obf_8d077f7cd507d18d,
	}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_12d48a2caf2868b5 *NRGBA) DrawImage(__obf_7e8297a5326972ed Palette, __obf_ec0ea5efc4ffd83f *PositionRect, __obf_552ded48e25434e1 *AreaRect) {
	__obf_c04c6105d2f272af := __obf_7e8297a5326972ed.Bounds().Max.X
	__obf_7cd3f0787f1e8091 := __obf_7e8297a5326972ed.Bounds().Max.Y

	__obf_282efdbfbdcdaded := __obf_ec0ea5efc4ffd83f.X
	__obf_78fc32d66e9ea5f1 := __obf_ec0ea5efc4ffd83f.Y
	__obf_fcbcfed86f846cc2 := __obf_ec0ea5efc4ffd83f.Height

	__obf_dcc1e427c3177912 := __obf_552ded48e25434e1.MinX
	__obf_5b930653000c5548 := __obf_552ded48e25434e1.MinY
	__obf_872f1db30809dbe2 := __obf_552ded48e25434e1.MaxX
	__obf_2044f05a9a44f819 := __obf_552ded48e25434e1.MaxY

	for __obf_ce61f234cbe42808 := range __obf_c04c6105d2f272af {
		for __obf_7a88a48a3d7db1aa := range __obf_7cd3f0787f1e8091 {
			__obf_ceed5a3454040e74 := __obf_7e8297a5326972ed.At(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa)
			if _, _, _, __obf_a66951353e16056a := __obf_ceed5a3454040e74.RGBA(); __obf_a66951353e16056a > 0 {
				if __obf_ce61f234cbe42808 >= __obf_dcc1e427c3177912 && __obf_ce61f234cbe42808 <= __obf_872f1db30809dbe2 && __obf_7a88a48a3d7db1aa >= __obf_5b930653000c5548 && __obf_7a88a48a3d7db1aa <= __obf_2044f05a9a44f819 {
					__obf_12d48a2caf2868b5.Set(__obf_282efdbfbdcdaded+(__obf_ce61f234cbe42808-__obf_dcc1e427c3177912), __obf_78fc32d66e9ea5f1-__obf_fcbcfed86f846cc2+(__obf_7a88a48a3d7db1aa-__obf_5b930653000c5548), __obf_7e8297a5326972ed.At(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_12d48a2caf2868b5 *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_c04c6105d2f272af := __obf_12d48a2caf2868b5.Bounds().Max.X
	__obf_7cd3f0787f1e8091 := __obf_12d48a2caf2868b5.Bounds().Max.Y
	__obf_8f8dd92499417004 := __obf_c04c6105d2f272af
	__obf_ac672b1f936c0c6e := 0
	__obf_9ce451081920ce95 := __obf_7cd3f0787f1e8091
	__obf_71ac0e772fc668a4 := 0
	for __obf_ce61f234cbe42808 := range __obf_c04c6105d2f272af {
		for __obf_7a88a48a3d7db1aa := range __obf_7cd3f0787f1e8091 {
			__obf_ceed5a3454040e74 := __obf_12d48a2caf2868b5.At(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa)
			if _, _, _, __obf_a66951353e16056a := __obf_ceed5a3454040e74.RGBA(); __obf_a66951353e16056a > 0 {
				if __obf_ce61f234cbe42808 < __obf_8f8dd92499417004 {
					__obf_8f8dd92499417004 = __obf_ce61f234cbe42808
				}
				if __obf_ce61f234cbe42808 > __obf_ac672b1f936c0c6e {
					__obf_ac672b1f936c0c6e = __obf_ce61f234cbe42808
				}

				if __obf_7a88a48a3d7db1aa < __obf_9ce451081920ce95 {
					__obf_9ce451081920ce95 = __obf_7a88a48a3d7db1aa
				}
				if __obf_7a88a48a3d7db1aa > __obf_71ac0e772fc668a4 {
					__obf_71ac0e772fc668a4 = __obf_7a88a48a3d7db1aa
				}
			}
		}
	}

	__obf_8f8dd92499417004 = int(max(0, float64(__obf_8f8dd92499417004-2)))
	__obf_ac672b1f936c0c6e = int(min(float64(__obf_c04c6105d2f272af), float64(__obf_ac672b1f936c0c6e+2)))
	__obf_9ce451081920ce95 = int(max(0, float64(__obf_9ce451081920ce95-2)))
	__obf_71ac0e772fc668a4 = int(min(float64(__obf_7cd3f0787f1e8091), float64(__obf_71ac0e772fc668a4+2)))

	return &AreaRect{
		__obf_8f8dd92499417004,
		__obf_ac672b1f936c0c6e,
		__obf_9ce451081920ce95,
		__obf_71ac0e772fc668a4,
	}
}

// Rotate rotates the canvas by any angle
func (__obf_12d48a2caf2868b5 *NRGBA) Rotate(__obf_a66951353e16056a int, __obf_e733c2dd705f7cf7 bool) {
	if __obf_a66951353e16056a == 0 {
		return
	}

	__obf_0a4f6c9b020fe9b9 := float64(__obf_a66951353e16056a) * math.Pi / 180

	__obf_23afce5199f5c83f := __obf_12d48a2caf2868b5.Bounds().Dx()
	__obf_d7684240f5ae06f9 := __obf_12d48a2caf2868b5.Bounds().Dy()
	__obf_3fbfb4aba2bc08f3, __obf_7113bad0317e187f := util.RotatedSize(__obf_23afce5199f5c83f, __obf_d7684240f5ae06f9, float64(__obf_a66951353e16056a))
	__obf_7e8297a5326972ed := image.NewNRGBA(image.Rect(0, 0, __obf_3fbfb4aba2bc08f3, __obf_7113bad0317e187f))

	__obf_b8628fbd80ea90ce := float64(__obf_3fbfb4aba2bc08f3) / 2
	__obf_26c040f6dd63aff4 := float64(__obf_7113bad0317e187f) / 2

	__obf_efaa56bf425d8a12 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_efaa56bf425d8a12 = __obf_efaa56bf425d8a12.Translate(__obf_b8628fbd80ea90ce, __obf_26c040f6dd63aff4)
	__obf_efaa56bf425d8a12 = __obf_efaa56bf425d8a12.Rotate(__obf_0a4f6c9b020fe9b9)
	__obf_efaa56bf425d8a12 = __obf_efaa56bf425d8a12.Translate(-__obf_b8628fbd80ea90ce, -__obf_26c040f6dd63aff4)

	__obf_ce61f234cbe42808 := (__obf_3fbfb4aba2bc08f3 - __obf_12d48a2caf2868b5.Bounds().Dx()) / 2
	__obf_7a88a48a3d7db1aa := (__obf_7113bad0317e187f - __obf_12d48a2caf2868b5.Bounds().Dy()) / 2
	__obf_94051787adea968d, __obf_7f4da368579ca3d9 := float64(__obf_ce61f234cbe42808), float64(__obf_7a88a48a3d7db1aa)

	__obf_05ad685cfa0359cc := __obf_efaa56bf425d8a12.Translate(__obf_94051787adea968d, __obf_7f4da368579ca3d9)
	__obf_162bfdf1f5eb40d0 := f64.Aff3{__obf_05ad685cfa0359cc.XX, __obf_05ad685cfa0359cc.XY, __obf_05ad685cfa0359cc.X0, __obf_05ad685cfa0359cc.YX, __obf_05ad685cfa0359cc.YY, __obf_05ad685cfa0359cc.Y0}

	draw.BiLinear.Transform(__obf_7e8297a5326972ed, __obf_162bfdf1f5eb40d0, __obf_12d48a2caf2868b5, __obf_12d48a2caf2868b5.Bounds(), draw.Over, nil)
	__obf_12d48a2caf2868b5.NRGBA = __obf_7e8297a5326972ed

	if __obf_e733c2dd705f7cf7 {
		__obf_047c233a057f77f8 := __obf_3fbfb4aba2bc08f3 - __obf_23afce5199f5c83f
		__obf_7e8fbf2f9d609e5b := __obf_7113bad0317e187f - __obf_d7684240f5ae06f9
		__obf_a0ccd38eacb0134e := (__obf_047c233a057f77f8 / 2) + 1
		__obf_4199cd4b4fbd0e10 := (__obf_7e8fbf2f9d609e5b / 2) + 1
		__obf_12d48a2caf2868b5.SubImage(image.Rect(__obf_a0ccd38eacb0134e, __obf_4199cd4b4fbd0e10, __obf_23afce5199f5c83f+__obf_a0ccd38eacb0134e, __obf_d7684240f5ae06f9+__obf_4199cd4b4fbd0e10))
	}
}

// CropCircle crops a circular area
func (__obf_12d48a2caf2868b5 *NRGBA) CropCircle(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_77b9044f3fa7cc43 int) {
	__obf_46f6f14eeeff546d := __obf_12d48a2caf2868b5.Bounds()
	__obf_1437ed4c79ba21fe := image.NewNRGBA(__obf_46f6f14eeeff546d)
	for __obf_42578c6ea60810ad := __obf_46f6f14eeeff546d.Min.Y; __obf_42578c6ea60810ad < __obf_46f6f14eeeff546d.Max.Y; __obf_42578c6ea60810ad++ {
		for __obf_1d2323197587f30c := __obf_46f6f14eeeff546d.Min.X; __obf_1d2323197587f30c < __obf_46f6f14eeeff546d.Max.X; __obf_1d2323197587f30c++ {
			__obf_4b02a812a7145793 := math.Hypot(float64(__obf_1d2323197587f30c-__obf_ce61f234cbe42808), float64(__obf_42578c6ea60810ad-__obf_7a88a48a3d7db1aa))
			if __obf_4b02a812a7145793 <= float64(__obf_77b9044f3fa7cc43) {
				__obf_1437ed4c79ba21fe.Set(__obf_1d2323197587f30c, __obf_42578c6ea60810ad, color.White)
			} else {
				__obf_1437ed4c79ba21fe.Set(__obf_1d2323197587f30c, __obf_42578c6ea60810ad, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_1437ed4c79ba21fe, __obf_1437ed4c79ba21fe.Bounds(), __obf_12d48a2caf2868b5, image.Point{X: 0, Y: 0}, __obf_1437ed4c79ba21fe, image.Point{}, draw.Over)
	__obf_12d48a2caf2868b5.NRGBA = __obf_1437ed4c79ba21fe
}

// CropScaleCircle scales and crops a circular area
func (__obf_12d48a2caf2868b5 *NRGBA) CropScaleCircle(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_77b9044f3fa7cc43 int, __obf_c2d392cd67e967d0 int) {
	__obf_46f6f14eeeff546d := __obf_12d48a2caf2868b5.Bounds()
	__obf_1437ed4c79ba21fe := image.NewNRGBA(__obf_46f6f14eeeff546d)

	for __obf_42578c6ea60810ad := __obf_46f6f14eeeff546d.Min.Y; __obf_42578c6ea60810ad < __obf_46f6f14eeeff546d.Max.Y; __obf_42578c6ea60810ad++ {
		for __obf_1d2323197587f30c := __obf_46f6f14eeeff546d.Min.X; __obf_1d2323197587f30c < __obf_46f6f14eeeff546d.Max.X; __obf_1d2323197587f30c++ {
			__obf_4b02a812a7145793 := math.Hypot(float64(__obf_1d2323197587f30c-__obf_ce61f234cbe42808), float64(__obf_42578c6ea60810ad-__obf_7a88a48a3d7db1aa))
			if __obf_4b02a812a7145793 <= float64(__obf_77b9044f3fa7cc43) {
				__obf_1437ed4c79ba21fe.Set(__obf_1d2323197587f30c, __obf_42578c6ea60810ad, color.White)
			} else {
				__obf_1437ed4c79ba21fe.Set(__obf_1d2323197587f30c, __obf_42578c6ea60810ad, color.Transparent)
			}
		}
	}

	if __obf_c2d392cd67e967d0 > 0 {
		__obf_21c0bce55b023756 := __obf_c2d392cd67e967d0 * 2
		__obf_c36baf52719c2e39 := image.NewNRGBA(image.Rect(0, 0, __obf_12d48a2caf2868b5.Bounds().Dx()-__obf_21c0bce55b023756, __obf_12d48a2caf2868b5.Bounds().Dy()-__obf_21c0bce55b023756))
		draw.BiLinear.Scale(__obf_c36baf52719c2e39, __obf_c36baf52719c2e39.Bounds(), __obf_1437ed4c79ba21fe, __obf_1437ed4c79ba21fe.Bounds(), draw.Over, nil)
		__obf_1437ed4c79ba21fe = __obf_c36baf52719c2e39
	}

	draw.DrawMask(__obf_1437ed4c79ba21fe, __obf_1437ed4c79ba21fe.Bounds(), __obf_12d48a2caf2868b5, image.Point{X: __obf_c2d392cd67e967d0, Y: __obf_c2d392cd67e967d0}, __obf_1437ed4c79ba21fe, image.Point{}, draw.Over)
	__obf_12d48a2caf2868b5.NRGBA = __obf_1437ed4c79ba21fe
}

// Scale scales the canvas
func (__obf_12d48a2caf2868b5 *NRGBA) Scale(__obf_c2d392cd67e967d0 int, __obf_6f918d00a2acfa41, __obf_b7da1752264579a8 bool) {
	__obf_7e8297a5326972ed := __obf_12d48a2caf2868b5.NRGBA
	if __obf_c2d392cd67e967d0 > 0 {
		__obf_21c0bce55b023756 := __obf_c2d392cd67e967d0 * 2
		__obf_07b60d2f28984390 := __obf_12d48a2caf2868b5.Bounds().Dx() - __obf_21c0bce55b023756
		__obf_171461f156c2ff06 := __obf_12d48a2caf2868b5.Bounds().Dy() - __obf_21c0bce55b023756
		__obf_4721c91d1fccee38 := image.NewNRGBA(image.Rect(0, 0, __obf_07b60d2f28984390, __obf_171461f156c2ff06))

		if !__obf_6f918d00a2acfa41 {
			draw.BiLinear.Scale(__obf_4721c91d1fccee38, __obf_4721c91d1fccee38.Bounds(), __obf_12d48a2caf2868b5, __obf_12d48a2caf2868b5.Bounds(), draw.Over, nil)
		} else {
			__obf_1407b37a85f597f4 := util.CalcResizedRect(__obf_12d48a2caf2868b5.Bounds(), __obf_07b60d2f28984390, __obf_171461f156c2ff06, __obf_b7da1752264579a8)
			draw.ApproxBiLinear.Scale(__obf_4721c91d1fccee38, __obf_1407b37a85f597f4.Bounds(), __obf_12d48a2caf2868b5, __obf_12d48a2caf2868b5.Bounds(), draw.Over, nil)
		}

		__obf_7e8297a5326972ed = __obf_4721c91d1fccee38
	}

	__obf_12d48a2caf2868b5.NRGBA = __obf_7e8297a5326972ed
}
