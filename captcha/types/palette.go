package __obf_6dcb1d06bd949756

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_22f909fc49d8ec44 image.Rectangle, __obf_285d17dd3dece463 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_22f909fc49d8ec44,

			// palette struct for palette canvas
			__obf_285d17dd3dece463),
	}
}

type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_285d17dd3dece463 *Palette) Rotate(__obf_dadbdf6b4f60b255 int) {
	if __obf_dadbdf6b4f60b255 == 0 {
		return
	}
	__obf_b0d9330f2201d023 := __obf_285d17dd3dece463
	__obf_aa6eacf895705c67 := __obf_b0d9330f2201d023.Bounds().Max.X
	__obf_b1f64ca00d21b5ce := __obf_b0d9330f2201d023.Bounds().Max.Y
	__obf_22f909fc49d8ec44 := __obf_aa6eacf895705c67 / 2
	__obf_ccecbfc5a8b45ab2 := image.NewPaletted(image.Rect(0, 0, __obf_aa6eacf895705c67, __obf_b1f64ca00d21b5ce), __obf_b0d9330f2201d023.Palette)
	for __obf_57d99e8bd0a3444e := 0; __obf_57d99e8bd0a3444e <= __obf_ccecbfc5a8b45ab2.Bounds().Max.X; __obf_57d99e8bd0a3444e++ {
		for __obf_3376ebf4b35235fa := 0; __obf_3376ebf4b35235fa <= __obf_ccecbfc5a8b45ab2.Bounds().Max.Y; __obf_3376ebf4b35235fa++ {
			__obf_55ce7a8f0436e128, __obf_8b9ee5dbffbff0e7 := __obf_285d17dd3dece463.AngleSwapPoint(float64(__obf_57d99e8bd0a3444e), float64(__obf_3376ebf4b35235fa), float64(__obf_22f909fc49d8ec44), float64(__obf_dadbdf6b4f60b255))
			__obf_ccecbfc5a8b45ab2.
				SetColorIndex(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_b0d9330f2201d023.ColorIndexAt(int(__obf_55ce7a8f0436e128), int(__obf_8b9ee5dbffbff0e7)))
		}
	}
	__obf_4d379dc6f44f3aa1 := __obf_ccecbfc5a8b45ab2.Bounds().Max.X
	__obf_d334e96f1dfe592d := __obf_ccecbfc5a8b45ab2.Bounds().Max.Y
	for __obf_57d99e8bd0a3444e := range __obf_4d379dc6f44f3aa1 {
		for __obf_3376ebf4b35235fa := range __obf_d334e96f1dfe592d {
			__obf_285d17dd3dece463.
				SetColorIndex(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa,

					// DrawCircle draws a circle
					__obf_ccecbfc5a8b45ab2.ColorIndexAt(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa))
		}
	}
}

func (__obf_285d17dd3dece463 *Palette) DrawCircle(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_99fb896ae843f26c int, __obf_6b14c339de9bea34 color.RGBA) {
	__obf_e06cc2e37ea49439 := 1 - __obf_99fb896ae843f26c
	__obf_80c9703017ab091f := 1
	__obf_157b9f2f8bafc210 := -2 * __obf_99fb896ae843f26c
	__obf_b1a049ae887c59f4 := 0
	__obf_931e87dc1858f367 := __obf_99fb896ae843f26c
	__obf_285d17dd3dece463.
		Set(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa+__obf_99fb896ae843f26c, __obf_6b14c339de9bea34)
	__obf_285d17dd3dece463.
		Set(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa-__obf_99fb896ae843f26c, __obf_6b14c339de9bea34)
	__obf_285d17dd3dece463.
		DrawHorizLine(__obf_57d99e8bd0a3444e-__obf_99fb896ae843f26c, __obf_57d99e8bd0a3444e+__obf_99fb896ae843f26c, __obf_3376ebf4b35235fa, __obf_6b14c339de9bea34)

	for __obf_b1a049ae887c59f4 < __obf_931e87dc1858f367 {
		if __obf_e06cc2e37ea49439 >= 0 {
			__obf_931e87dc1858f367--
			__obf_157b9f2f8bafc210 += 2
			__obf_e06cc2e37ea49439 += __obf_157b9f2f8bafc210
		}
		__obf_b1a049ae887c59f4++
		__obf_80c9703017ab091f += 2
		__obf_e06cc2e37ea49439 += __obf_80c9703017ab091f
		__obf_285d17dd3dece463.
			DrawHorizLine(__obf_57d99e8bd0a3444e-__obf_b1a049ae887c59f4, __obf_57d99e8bd0a3444e+__obf_b1a049ae887c59f4, __obf_3376ebf4b35235fa+__obf_931e87dc1858f367, __obf_6b14c339de9bea34)
		__obf_285d17dd3dece463.
			DrawHorizLine(__obf_57d99e8bd0a3444e-__obf_b1a049ae887c59f4, __obf_57d99e8bd0a3444e+__obf_b1a049ae887c59f4, __obf_3376ebf4b35235fa-__obf_931e87dc1858f367, __obf_6b14c339de9bea34)
		__obf_285d17dd3dece463.
			DrawHorizLine(__obf_57d99e8bd0a3444e-__obf_931e87dc1858f367, __obf_57d99e8bd0a3444e+__obf_931e87dc1858f367, __obf_3376ebf4b35235fa+__obf_b1a049ae887c59f4, __obf_6b14c339de9bea34)
		__obf_285d17dd3dece463.
			DrawHorizLine(__obf_57d99e8bd0a3444e-__obf_931e87dc1858f367, __obf_57d99e8bd0a3444e+__obf_931e87dc1858f367,

				// DrawHorizLine draws a horizontal line
				__obf_3376ebf4b35235fa-__obf_b1a049ae887c59f4, __obf_6b14c339de9bea34)
	}
}

func (__obf_285d17dd3dece463 *Palette) DrawHorizLine(__obf_18f6397e0daeb2c5, __obf_9c2e9bae92fc2001, __obf_3376ebf4b35235fa int, __obf_6b14c339de9bea34 color.RGBA) {
	for __obf_57d99e8bd0a3444e := __obf_18f6397e0daeb2c5; __obf_57d99e8bd0a3444e <= __obf_9c2e9bae92fc2001; __obf_57d99e8bd0a3444e++ {
		__obf_285d17dd3dece463.
			Set(__obf_57d99e8bd0a3444e,

				// Distort distorts the canvas
				__obf_3376ebf4b35235fa, __obf_6b14c339de9bea34)
	}
}

func (__obf_285d17dd3dece463 *Palette) Distort(__obf_cc1000eb1d86a6f4 float64, __obf_859a83c0bbaa9f68 float64) {
	__obf_1f037513843595bc := __obf_285d17dd3dece463.Bounds().Max.X
	__obf_ac11522e37274a9b := __obf_285d17dd3dece463.Bounds().Max.Y
	__obf_f0325dce455ca156 := NewPalette(image.Rect(0, 0, __obf_1f037513843595bc, __obf_ac11522e37274a9b), __obf_285d17dd3dece463.Palette)
	__obf_e56a0193d7dbf602 := 2.0 * math.Pi / __obf_859a83c0bbaa9f68
	for __obf_57d99e8bd0a3444e := range __obf_1f037513843595bc {
		for __obf_3376ebf4b35235fa := 0; __obf_3376ebf4b35235fa < __obf_ac11522e37274a9b; __obf_3376ebf4b35235fa++ {
			__obf_b1a049ae887c59f4 := __obf_cc1000eb1d86a6f4 * math.Sin(float64(__obf_3376ebf4b35235fa)*__obf_e56a0193d7dbf602)
			__obf_931e87dc1858f367 := __obf_cc1000eb1d86a6f4 * math.Cos(float64(__obf_57d99e8bd0a3444e)*__obf_e56a0193d7dbf602)
			__obf_f0325dce455ca156.
				SetColorIndex(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_285d17dd3dece463.ColorIndexAt(__obf_57d99e8bd0a3444e+int(__obf_b1a049ae887c59f4), __obf_3376ebf4b35235fa+int(__obf_931e87dc1858f367)))
		}
	}
	__obf_4d379dc6f44f3aa1 := __obf_f0325dce455ca156.Bounds().Max.X
	__obf_d334e96f1dfe592d := __obf_f0325dce455ca156.Bounds().Max.Y
	for __obf_57d99e8bd0a3444e := range __obf_4d379dc6f44f3aa1 {
		for __obf_3376ebf4b35235fa := 0; __obf_3376ebf4b35235fa < __obf_d334e96f1dfe592d; __obf_3376ebf4b35235fa++ {
			__obf_285d17dd3dece463.
				SetColorIndex(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_f0325dce455ca156.ColorIndexAt(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa))
		}
	}
	__obf_f0325dce455ca156.
		Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_285d17dd3dece463 *Palette) DrawBeeline(__obf_9e0c3e490ad8f2d5 image.Point, __obf_7370cc7161f48ae9 image.Point, __obf_67447fed68871f04 color.RGBA) {
	__obf_e56a0193d7dbf602 := math.Abs(float64(__obf_9e0c3e490ad8f2d5.X - __obf_7370cc7161f48ae9.X))
	__obf_f23a92c76d13cc24 := math.Abs(float64(__obf_7370cc7161f48ae9.Y - __obf_9e0c3e490ad8f2d5.Y))
	__obf_a4ececaaee4cdbbe, __obf_e00ef5c3b0a1efb9 := 1, 1
	if __obf_9e0c3e490ad8f2d5.X >= __obf_7370cc7161f48ae9.X {
		__obf_a4ececaaee4cdbbe = -1
	}
	if __obf_9e0c3e490ad8f2d5.Y >= __obf_7370cc7161f48ae9.Y {
		__obf_e00ef5c3b0a1efb9 = -1
	}
	__obf_d349b3521516d4ce := __obf_e56a0193d7dbf602 - __obf_f23a92c76d13cc24
	for {
		__obf_285d17dd3dece463.
			Set(__obf_9e0c3e490ad8f2d5.X, __obf_9e0c3e490ad8f2d5.Y, __obf_67447fed68871f04)
		__obf_285d17dd3dece463.
			Set(__obf_9e0c3e490ad8f2d5.X+1, __obf_9e0c3e490ad8f2d5.Y, __obf_67447fed68871f04)
		__obf_285d17dd3dece463.
			Set(__obf_9e0c3e490ad8f2d5.X-1, __obf_9e0c3e490ad8f2d5.Y, __obf_67447fed68871f04)
		__obf_285d17dd3dece463.
			Set(__obf_9e0c3e490ad8f2d5.X+2, __obf_9e0c3e490ad8f2d5.Y, __obf_67447fed68871f04)
		__obf_285d17dd3dece463.
			Set(__obf_9e0c3e490ad8f2d5.X-2, __obf_9e0c3e490ad8f2d5.Y, __obf_67447fed68871f04)
		if __obf_9e0c3e490ad8f2d5.X == __obf_7370cc7161f48ae9.X && __obf_9e0c3e490ad8f2d5.Y == __obf_7370cc7161f48ae9.Y {
			return
		}
		__obf_488d04945f49f309 := __obf_d349b3521516d4ce * 2
		if __obf_488d04945f49f309 > -__obf_f23a92c76d13cc24 {
			__obf_d349b3521516d4ce -= __obf_f23a92c76d13cc24
			__obf_9e0c3e490ad8f2d5.
				X += __obf_a4ececaaee4cdbbe
		}
		if __obf_488d04945f49f309 < __obf_e56a0193d7dbf602 {
			__obf_d349b3521516d4ce += __obf_e56a0193d7dbf602
			__obf_9e0c3e490ad8f2d5.
				Y += __obf_e00ef5c3b0a1efb9
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_285d17dd3dece463 *Palette) AngleSwapPoint(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa, __obf_22f909fc49d8ec44, __obf_dadbdf6b4f60b255 float64) (__obf_22568cf301d6f9f7, __obf_0c631971ef2432dc float64) {
	__obf_57d99e8bd0a3444e -= __obf_22f909fc49d8ec44
	__obf_3376ebf4b35235fa = __obf_22f909fc49d8ec44 - __obf_3376ebf4b35235fa
	__obf_e272b3fae3b09432 := math.Sin(__obf_dadbdf6b4f60b255 * (math.Pi / 180))
	__obf_b35b7c9644fc6fb5 := math.Cos(__obf_dadbdf6b4f60b255 * (math.Pi / 180))
	__obf_22568cf301d6f9f7 = __obf_57d99e8bd0a3444e*__obf_b35b7c9644fc6fb5 + __obf_3376ebf4b35235fa*__obf_e272b3fae3b09432
	__obf_0c631971ef2432dc = -__obf_57d99e8bd0a3444e*__obf_e272b3fae3b09432 + __obf_3376ebf4b35235fa*__obf_b35b7c9644fc6fb5
	__obf_22568cf301d6f9f7 += __obf_22f909fc49d8ec44
	__obf_0c631971ef2432dc = __obf_22f909fc49d8ec44 - __obf_0c631971ef2432dc
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_285d17dd3dece463 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_4d379dc6f44f3aa1 := __obf_285d17dd3dece463.Bounds().Max.X
	__obf_d334e96f1dfe592d := __obf_285d17dd3dece463.Bounds().Max.Y
	__obf_872716a227557da0 := __obf_4d379dc6f44f3aa1
	__obf_04cdedc6038048f9 := 0
	__obf_3bbf802af4ff2cdd := __obf_d334e96f1dfe592d
	__obf_97e0ab00d4cde441 := 0
	for __obf_57d99e8bd0a3444e := range __obf_4d379dc6f44f3aa1 {
		for __obf_3376ebf4b35235fa := range __obf_d334e96f1dfe592d {
			__obf_2039314a03b948f3 := __obf_285d17dd3dece463.At(__obf_57d99e8bd0a3444e, __obf_3376ebf4b35235fa)
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

	return &AreaRect{__obf_872716a227557da0, __obf_04cdedc6038048f9, __obf_3bbf802af4ff2cdd, __obf_97e0ab00d4cde441}
}
