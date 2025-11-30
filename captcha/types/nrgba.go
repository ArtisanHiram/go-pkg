package __obf_6dcb1d06bd949756

import (
	util "github.com/ArtisanHiram/go-pkg/captcha/util"
	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
	"image"
	"image/color"
	"math"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_22f909fc49d8ec44 image.Rectangle, __obf_72e01c2e6121887b bool) *NRGBA {
	__obf_e14c2d5d262dcb36 := image.NewNRGBA(__obf_22f909fc49d8ec44)
	for __obf_3376ebf4b35235fa := 0; __obf_3376ebf4b35235fa < __obf_22f909fc49d8ec44.Max.Y; __obf_3376ebf4b35235fa++ {
		for __obf_57d99e8bd0a3444e := 0; __obf_57d99e8bd0a3444e < __obf_22f909fc49d8ec44.Max.X; __obf_57d99e8bd0a3444e++ {
			if __obf_72e01c2e6121887b {
				__obf_e14c2d5d262dcb36.
					Set(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, color.Alpha{A: uint8(0)})
			} else {
				__obf_e14c2d5d262dcb36.
					Set(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{__obf_e14c2d5d262dcb36}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_9a07ffd3335deb22 *NRGBA) DrawImage(__obf_587078e38cdb4111 Palette, __obf_08751c1bf1a6a3f5 *PositionRect, __obf_f8e40e9c14ad5e31 *AreaRect) {
	__obf_4d379dc6f44f3aa1 := __obf_587078e38cdb4111.Bounds().Max.X
	__obf_d334e96f1dfe592d := __obf_587078e38cdb4111.Bounds().Max.Y
	__obf_6a7a87d2232f59fe := __obf_08751c1bf1a6a3f5.X
	__obf_ad8e8af201bbf2bf := __obf_08751c1bf1a6a3f5.Y
	__obf_72fb5d6b686b92ff := __obf_08751c1bf1a6a3f5.Height
	__obf_ed1a220719aad352 := __obf_f8e40e9c14ad5e31.MinX
	__obf_2533de6d920ae75c := __obf_f8e40e9c14ad5e31.MinY
	__obf_a2420db71565a410 := __obf_f8e40e9c14ad5e31.MaxX
	__obf_c952301e80b65985 := __obf_f8e40e9c14ad5e31.MaxY

	for __obf_57d99e8bd0a3444e := range __obf_4d379dc6f44f3aa1 {
		for __obf_3376ebf4b35235fa := range __obf_d334e96f1dfe592d {
			__obf_2039314a03b948f3 := __obf_587078e38cdb4111.At(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa)
			if _, _, _, __obf_59e7114cb8b11964 := __obf_2039314a03b948f3.RGBA(); __obf_59e7114cb8b11964 > 0 {
				if __obf_57d99e8bd0a3444e >= __obf_ed1a220719aad352 && __obf_57d99e8bd0a3444e <= __obf_a2420db71565a410 && __obf_3376ebf4b35235fa >= __obf_2533de6d920ae75c && __obf_3376ebf4b35235fa <= __obf_c952301e80b65985 {
					__obf_9a07ffd3335deb22.
						Set(__obf_6a7a87d2232f59fe+(__obf_57d99e8bd0a3444e-__obf_ed1a220719aad352), __obf_ad8e8af201bbf2bf-__obf_72fb5d6b686b92ff+(__obf_3376ebf4b35235fa-__obf_2533de6d920ae75c), __obf_587078e38cdb4111.At(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_9a07ffd3335deb22 *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_4d379dc6f44f3aa1 := __obf_9a07ffd3335deb22.Bounds().Max.X
	__obf_d334e96f1dfe592d := __obf_9a07ffd3335deb22.Bounds().Max.Y
	__obf_872716a227557da0 := __obf_4d379dc6f44f3aa1
	__obf_04cdedc6038048f9 := 0
	__obf_3bbf802af4ff2cdd := __obf_d334e96f1dfe592d
	__obf_97e0ab00d4cde441 := 0
	for __obf_57d99e8bd0a3444e := range __obf_4d379dc6f44f3aa1 {
		for __obf_3376ebf4b35235fa := range __obf_d334e96f1dfe592d {
			__obf_2039314a03b948f3 := __obf_9a07ffd3335deb22.At(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa)
			if _, _, _, __obf_59e7114cb8b11964 := __obf_2039314a03b948f3.RGBA(); __obf_59e7114cb8b11964 > 0 {
				if __obf_57d99e8bd0a3444e < __obf_872716a227557da0 {
					__obf_872716a227557da0 = __obf_57d99e8bd0a3444e
				}
				if __obf_57d99e8bd0a3444e > __obf_04cdedc6038048f9 {
					__obf_04cdedc6038048f9 = __obf_57d99e8bd0a3444e
				}

				if __obf_3376ebf4b35235fa < __obf_3bbf802af4ff2cdd {
					__obf_3bbf802af4ff2cdd = __obf_3376ebf4b35235fa
				}
				if __obf_3376ebf4b35235fa > __obf_97e0ab00d4cde441 {
					__obf_97e0ab00d4cde441 = __obf_3376ebf4b35235fa
				}
			}
		}
	}
	__obf_872716a227557da0 = int(max(0, float64(__obf_872716a227557da0-2)))
	__obf_04cdedc6038048f9 = int(min(float64(__obf_4d379dc6f44f3aa1), float64(__obf_04cdedc6038048f9+2)))
	__obf_3bbf802af4ff2cdd = int(max(0, float64(__obf_3bbf802af4ff2cdd-2)))
	__obf_97e0ab00d4cde441 = int(min(float64(__obf_d334e96f1dfe592d), float64(__obf_97e0ab00d4cde441+2)))

	return &AreaRect{__obf_872716a227557da0, __obf_04cdedc6038048f9,

		// Rotate rotates the canvas by any angle
		__obf_3bbf802af4ff2cdd, __obf_97e0ab00d4cde441,
	}
}

func (__obf_9a07ffd3335deb22 *NRGBA) Rotate(__obf_59e7114cb8b11964 int, __obf_fe8e65a07ec5ff38 bool) {
	if __obf_59e7114cb8b11964 == 0 {
		return
	}
	__obf_dadbdf6b4f60b255 := float64(__obf_59e7114cb8b11964) * math.Pi / 180
	__obf_b3bf2bb53e2d7586 := __obf_9a07ffd3335deb22.Bounds().Dx()
	__obf_29b765c8d33ae80b := __obf_9a07ffd3335deb22.Bounds().Dy()
	__obf_1f037513843595bc, __obf_ac11522e37274a9b := util.RotatedSize(__obf_b3bf2bb53e2d7586, __obf_29b765c8d33ae80b, float64(__obf_59e7114cb8b11964))
	__obf_587078e38cdb4111 := image.NewNRGBA(image.Rect(0, 0, __obf_1f037513843595bc, __obf_ac11522e37274a9b))
	__obf_975aeb53386cbd26 := float64(__obf_1f037513843595bc) / 2
	__obf_40fe6b1e0582cce4 := float64(__obf_ac11522e37274a9b) / 2
	__obf_3efd0805fa5cfe7c := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_3efd0805fa5cfe7c = __obf_3efd0805fa5cfe7c.Translate(__obf_975aeb53386cbd26, __obf_40fe6b1e0582cce4)
	__obf_3efd0805fa5cfe7c = __obf_3efd0805fa5cfe7c.Rotate(__obf_dadbdf6b4f60b255)
	__obf_3efd0805fa5cfe7c = __obf_3efd0805fa5cfe7c.Translate(-__obf_975aeb53386cbd26, -__obf_40fe6b1e0582cce4)
	__obf_57d99e8bd0a3444e := (__obf_1f037513843595bc - __obf_9a07ffd3335deb22.Bounds().Dx()) / 2
	__obf_3376ebf4b35235fa := (__obf_ac11522e37274a9b - __obf_9a07ffd3335deb22.Bounds().Dy()) / 2
	__obf_bfccf54a230f5965, __obf_251e9d6faf6e6399 := float64(__obf_57d99e8bd0a3444e), float64(__obf_3376ebf4b35235fa)
	__obf_d73d886accf933e7 := __obf_3efd0805fa5cfe7c.Translate(__obf_bfccf54a230f5965, __obf_251e9d6faf6e6399)
	__obf_824941f23a08cd1f := f64.Aff3{__obf_d73d886accf933e7.XX, __obf_d73d886accf933e7.XY, __obf_d73d886accf933e7.X0, __obf_d73d886accf933e7.YX, __obf_d73d886accf933e7.YY, __obf_d73d886accf933e7.Y0}

	draw.BiLinear.Transform(__obf_587078e38cdb4111, __obf_824941f23a08cd1f, __obf_9a07ffd3335deb22, __obf_9a07ffd3335deb22.Bounds(), draw.Over, nil)
	__obf_9a07ffd3335deb22.
		NRGBA = __obf_587078e38cdb4111

	if __obf_fe8e65a07ec5ff38 {
		__obf_abdcfc06e8dd4373 := __obf_1f037513843595bc - __obf_b3bf2bb53e2d7586
		__obf_aa1fcef8d75afe89 := __obf_ac11522e37274a9b - __obf_29b765c8d33ae80b
		__obf_e56a0193d7dbf602 := (__obf_abdcfc06e8dd4373 / 2) + 1
		__obf_f23a92c76d13cc24 := (__obf_aa1fcef8d75afe89 / 2) + 1
		__obf_9a07ffd3335deb22.
			SubImage(image.Rect(__obf_e56a0193d7dbf602, __obf_f23a92c76d13cc24,

				// CropCircle crops a circular area
				__obf_b3bf2bb53e2d7586+__obf_e56a0193d7dbf602, __obf_29b765c8d33ae80b+__obf_f23a92c76d13cc24))
	}
}

func (__obf_9a07ffd3335deb22 *NRGBA) CropCircle(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_99fb896ae843f26c int) {
	__obf_a982933c2d6742a9 := __obf_9a07ffd3335deb22.Bounds()
	__obf_65806c398ff76ef1 := image.NewNRGBA(__obf_a982933c2d6742a9)
	for __obf_740911224ae50f3f := __obf_a982933c2d6742a9.Min.Y; __obf_740911224ae50f3f < __obf_a982933c2d6742a9.Max.Y; __obf_740911224ae50f3f++ {
		for __obf_35332c239115ba75 := __obf_a982933c2d6742a9.Min.X; __obf_35332c239115ba75 < __obf_a982933c2d6742a9.Max.X; __obf_35332c239115ba75++ {
			__obf_da6de6a13fe84822 := math.Hypot(float64(__obf_35332c239115ba75-__obf_57d99e8bd0a3444e), float64(__obf_740911224ae50f3f-__obf_3376ebf4b35235fa))
			if __obf_da6de6a13fe84822 <= float64(__obf_99fb896ae843f26c) {
				__obf_65806c398ff76ef1.
					Set(__obf_35332c239115ba75, __obf_740911224ae50f3f, color.White)
			} else {
				__obf_65806c398ff76ef1.
					Set(__obf_35332c239115ba75, __obf_740911224ae50f3f, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_65806c398ff76ef1, __obf_65806c398ff76ef1.Bounds(), __obf_9a07ffd3335deb22, image.Point{X: 0, Y: 0}, __obf_65806c398ff76ef1, image.Point{}, draw.Over)
	__obf_9a07ffd3335deb22.
		NRGBA = __obf_65806c398ff76ef1
}

// CropScaleCircle scales and crops a circular area
func (__obf_9a07ffd3335deb22 *NRGBA) CropScaleCircle(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_99fb896ae843f26c int, __obf_96e46bb9447fe5fc int) {
	__obf_a982933c2d6742a9 := __obf_9a07ffd3335deb22.Bounds()
	__obf_65806c398ff76ef1 := image.NewNRGBA(__obf_a982933c2d6742a9)

	for __obf_740911224ae50f3f := __obf_a982933c2d6742a9.Min.Y; __obf_740911224ae50f3f < __obf_a982933c2d6742a9.Max.Y; __obf_740911224ae50f3f++ {
		for __obf_35332c239115ba75 := __obf_a982933c2d6742a9.Min.X; __obf_35332c239115ba75 < __obf_a982933c2d6742a9.Max.X; __obf_35332c239115ba75++ {
			__obf_da6de6a13fe84822 := math.Hypot(float64(__obf_35332c239115ba75-__obf_57d99e8bd0a3444e), float64(__obf_740911224ae50f3f-__obf_3376ebf4b35235fa))
			if __obf_da6de6a13fe84822 <= float64(__obf_99fb896ae843f26c) {
				__obf_65806c398ff76ef1.
					Set(__obf_35332c239115ba75, __obf_740911224ae50f3f, color.White)
			} else {
				__obf_65806c398ff76ef1.
					Set(__obf_35332c239115ba75, __obf_740911224ae50f3f, color.Transparent)
			}
		}
	}

	if __obf_96e46bb9447fe5fc > 0 {
		__obf_021fb399db8407ec := __obf_96e46bb9447fe5fc * 2
		__obf_af2d722dd8b19e6a := image.NewNRGBA(image.Rect(0, 0, __obf_9a07ffd3335deb22.Bounds().Dx()-__obf_021fb399db8407ec, __obf_9a07ffd3335deb22.Bounds().Dy()-__obf_021fb399db8407ec))
		draw.BiLinear.Scale(__obf_af2d722dd8b19e6a, __obf_af2d722dd8b19e6a.Bounds(), __obf_65806c398ff76ef1, __obf_65806c398ff76ef1.Bounds(), draw.Over, nil)
		__obf_65806c398ff76ef1 = __obf_af2d722dd8b19e6a
	}

	draw.DrawMask(__obf_65806c398ff76ef1, __obf_65806c398ff76ef1.Bounds(), __obf_9a07ffd3335deb22, image.Point{X: __obf_96e46bb9447fe5fc, Y: __obf_96e46bb9447fe5fc}, __obf_65806c398ff76ef1, image.Point{}, draw.Over)
	__obf_9a07ffd3335deb22.
		NRGBA = __obf_65806c398ff76ef1
}

// Scale scales the canvas
func (__obf_9a07ffd3335deb22 *NRGBA) Scale(__obf_96e46bb9447fe5fc int, __obf_7b88c89d431ba915, __obf_1ae362e0d907afdb bool) {
	__obf_587078e38cdb4111 := __obf_9a07ffd3335deb22.NRGBA
	if __obf_96e46bb9447fe5fc > 0 {
		__obf_021fb399db8407ec := __obf_96e46bb9447fe5fc * 2
		__obf_ed796b16648298f7 := __obf_9a07ffd3335deb22.Bounds().Dx() - __obf_021fb399db8407ec
		__obf_1f546ebacf9ec396 := __obf_9a07ffd3335deb22.Bounds().Dy() - __obf_021fb399db8407ec
		__obf_c135bec9cde0b445 := image.NewNRGBA(image.Rect(0, 0, __obf_ed796b16648298f7, __obf_1f546ebacf9ec396))

		if !__obf_7b88c89d431ba915 {
			draw.BiLinear.Scale(__obf_c135bec9cde0b445, __obf_c135bec9cde0b445.Bounds(), __obf_9a07ffd3335deb22, __obf_9a07ffd3335deb22.Bounds(), draw.Over, nil)
		} else {
			__obf_86d6f848574fe427 := util.CalcResizedRect(__obf_9a07ffd3335deb22.Bounds(), __obf_ed796b16648298f7, __obf_1f546ebacf9ec396, __obf_1ae362e0d907afdb)
			draw.ApproxBiLinear.Scale(__obf_c135bec9cde0b445, __obf_86d6f848574fe427.Bounds(), __obf_9a07ffd3335deb22, __obf_9a07ffd3335deb22.Bounds(), draw.Over, nil)
		}
		__obf_587078e38cdb4111 = __obf_c135bec9cde0b445
	}
	__obf_9a07ffd3335deb22.
		NRGBA = __obf_587078e38cdb4111
}
