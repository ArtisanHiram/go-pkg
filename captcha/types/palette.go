package __obf_154ce31cd9492d61

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_c3f6882d3032b52f image.Rectangle, __obf_91a96cde82d83f16 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_c3f6882d3032b52f, __obf_91a96cde82d83f16),
	}
}

// palette struct for palette canvas
type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_91a96cde82d83f16 *Palette) Rotate(__obf_2c1097e2e9ce1f24 int) {
	if __obf_2c1097e2e9ce1f24 == 0 {
		return
	}

	__obf_b389498684e2f4ab := __obf_91a96cde82d83f16
	__obf_a7f53fc21ff4cb69 := __obf_b389498684e2f4ab.Bounds().Max.X
	__obf_70ad39a3b8bc7be4 := __obf_b389498684e2f4ab.Bounds().Max.Y
	__obf_c3f6882d3032b52f := __obf_a7f53fc21ff4cb69 / 2
	__obf_6e78b58a0254a98f := image.NewPaletted(image.Rect(0, 0, __obf_a7f53fc21ff4cb69, __obf_70ad39a3b8bc7be4), __obf_b389498684e2f4ab.Palette)
	for __obf_c87d6c7ba4a30d82 := 0; __obf_c87d6c7ba4a30d82 <= __obf_6e78b58a0254a98f.Bounds().Max.X; __obf_c87d6c7ba4a30d82++ {
		for __obf_f8fead3d682d659f := 0; __obf_f8fead3d682d659f <= __obf_6e78b58a0254a98f.Bounds().Max.Y; __obf_f8fead3d682d659f++ {
			__obf_fb77e68506ce52e0, __obf_f3808609fa19850f := __obf_91a96cde82d83f16.AngleSwapPoint(float64(__obf_c87d6c7ba4a30d82), float64(__obf_f8fead3d682d659f), float64(__obf_c3f6882d3032b52f), float64(__obf_2c1097e2e9ce1f24))
			__obf_6e78b58a0254a98f.SetColorIndex(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_b389498684e2f4ab.ColorIndexAt(int(__obf_fb77e68506ce52e0), int(__obf_f3808609fa19850f)))
		}
	}

	__obf_d14caaad13625683 := __obf_6e78b58a0254a98f.Bounds().Max.X
	__obf_ff9981d45d8deb98 := __obf_6e78b58a0254a98f.Bounds().Max.Y
	for __obf_c87d6c7ba4a30d82 := range __obf_d14caaad13625683 {
		for __obf_f8fead3d682d659f := range __obf_ff9981d45d8deb98 {
			__obf_91a96cde82d83f16.SetColorIndex(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_6e78b58a0254a98f.ColorIndexAt(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f))
		}
	}
}

// DrawCircle draws a circle
func (__obf_91a96cde82d83f16 *Palette) DrawCircle(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_ac266a6cc46c4953 int, __obf_ad3a8dc916d9ec04 color.RGBA) {
	__obf_00976d4b76af639d := 1 - __obf_ac266a6cc46c4953
	__obf_5e86c11e01b92b16 := 1
	__obf_a305e991ad9f6dc3 := -2 * __obf_ac266a6cc46c4953
	__obf_990d2e0d5849ff80 := 0
	__obf_79ab9e3d1bd1d7f4 := __obf_ac266a6cc46c4953

	__obf_91a96cde82d83f16.Set(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f+__obf_ac266a6cc46c4953, __obf_ad3a8dc916d9ec04)
	__obf_91a96cde82d83f16.Set(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f-__obf_ac266a6cc46c4953, __obf_ad3a8dc916d9ec04)
	__obf_91a96cde82d83f16.DrawHorizLine(__obf_c87d6c7ba4a30d82-__obf_ac266a6cc46c4953, __obf_c87d6c7ba4a30d82+__obf_ac266a6cc46c4953, __obf_f8fead3d682d659f, __obf_ad3a8dc916d9ec04)

	for __obf_990d2e0d5849ff80 < __obf_79ab9e3d1bd1d7f4 {
		if __obf_00976d4b76af639d >= 0 {
			__obf_79ab9e3d1bd1d7f4--
			__obf_a305e991ad9f6dc3 += 2
			__obf_00976d4b76af639d += __obf_a305e991ad9f6dc3
		}
		__obf_990d2e0d5849ff80++
		__obf_5e86c11e01b92b16 += 2
		__obf_00976d4b76af639d += __obf_5e86c11e01b92b16
		__obf_91a96cde82d83f16.DrawHorizLine(__obf_c87d6c7ba4a30d82-__obf_990d2e0d5849ff80, __obf_c87d6c7ba4a30d82+__obf_990d2e0d5849ff80, __obf_f8fead3d682d659f+__obf_79ab9e3d1bd1d7f4, __obf_ad3a8dc916d9ec04)
		__obf_91a96cde82d83f16.DrawHorizLine(__obf_c87d6c7ba4a30d82-__obf_990d2e0d5849ff80, __obf_c87d6c7ba4a30d82+__obf_990d2e0d5849ff80, __obf_f8fead3d682d659f-__obf_79ab9e3d1bd1d7f4, __obf_ad3a8dc916d9ec04)
		__obf_91a96cde82d83f16.DrawHorizLine(__obf_c87d6c7ba4a30d82-__obf_79ab9e3d1bd1d7f4, __obf_c87d6c7ba4a30d82+__obf_79ab9e3d1bd1d7f4, __obf_f8fead3d682d659f+__obf_990d2e0d5849ff80, __obf_ad3a8dc916d9ec04)
		__obf_91a96cde82d83f16.DrawHorizLine(__obf_c87d6c7ba4a30d82-__obf_79ab9e3d1bd1d7f4, __obf_c87d6c7ba4a30d82+__obf_79ab9e3d1bd1d7f4, __obf_f8fead3d682d659f-__obf_990d2e0d5849ff80, __obf_ad3a8dc916d9ec04)
	}
}

// DrawHorizLine draws a horizontal line
func (__obf_91a96cde82d83f16 *Palette) DrawHorizLine(__obf_ed7370a4f40eac65, __obf_5b15e5cee53069ea, __obf_f8fead3d682d659f int, __obf_ad3a8dc916d9ec04 color.RGBA) {
	for __obf_c87d6c7ba4a30d82 := __obf_ed7370a4f40eac65; __obf_c87d6c7ba4a30d82 <= __obf_5b15e5cee53069ea; __obf_c87d6c7ba4a30d82++ {
		__obf_91a96cde82d83f16.Set(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_ad3a8dc916d9ec04)
	}
}

// Distort distorts the canvas
func (__obf_91a96cde82d83f16 *Palette) Distort(__obf_e120b3693d6edc37 float64, __obf_56e46a699dd409fa float64) {
	__obf_5c507cf47bf5267e := __obf_91a96cde82d83f16.Bounds().Max.X
	__obf_f03bf0f36c5833f9 := __obf_91a96cde82d83f16.Bounds().Max.Y
	__obf_dc52efb8af25b8a0 := NewPalette(image.Rect(0, 0, __obf_5c507cf47bf5267e, __obf_f03bf0f36c5833f9), __obf_91a96cde82d83f16.Palette)
	__obf_2ab04c1711998195 := 2.0 * math.Pi / __obf_56e46a699dd409fa
	for __obf_c87d6c7ba4a30d82 := range __obf_5c507cf47bf5267e {
		for __obf_f8fead3d682d659f := 0; __obf_f8fead3d682d659f < __obf_f03bf0f36c5833f9; __obf_f8fead3d682d659f++ {
			__obf_990d2e0d5849ff80 := __obf_e120b3693d6edc37 * math.Sin(float64(__obf_f8fead3d682d659f)*__obf_2ab04c1711998195)
			__obf_79ab9e3d1bd1d7f4 := __obf_e120b3693d6edc37 * math.Cos(float64(__obf_c87d6c7ba4a30d82)*__obf_2ab04c1711998195)
			__obf_dc52efb8af25b8a0.SetColorIndex(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_91a96cde82d83f16.ColorIndexAt(__obf_c87d6c7ba4a30d82+int(__obf_990d2e0d5849ff80), __obf_f8fead3d682d659f+int(__obf_79ab9e3d1bd1d7f4)))
		}
	}

	__obf_d14caaad13625683 := __obf_dc52efb8af25b8a0.Bounds().Max.X
	__obf_ff9981d45d8deb98 := __obf_dc52efb8af25b8a0.Bounds().Max.Y
	for __obf_c87d6c7ba4a30d82 := range __obf_d14caaad13625683 {
		for __obf_f8fead3d682d659f := 0; __obf_f8fead3d682d659f < __obf_ff9981d45d8deb98; __obf_f8fead3d682d659f++ {
			__obf_91a96cde82d83f16.SetColorIndex(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_dc52efb8af25b8a0.ColorIndexAt(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f))
		}
	}

	__obf_dc52efb8af25b8a0.Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_91a96cde82d83f16 *Palette) DrawBeeline(__obf_fa6a97d0e7da39fb image.Point, __obf_dfa162ecb02c063a image.Point, __obf_20b69069a99344ee color.RGBA) {
	__obf_2ab04c1711998195 := math.Abs(float64(__obf_fa6a97d0e7da39fb.X - __obf_dfa162ecb02c063a.X))
	__obf_45d338c45fe3468d := math.Abs(float64(__obf_dfa162ecb02c063a.Y - __obf_fa6a97d0e7da39fb.Y))
	__obf_963d969f08b1d119, __obf_8ff9861776e145d0 := 1, 1
	if __obf_fa6a97d0e7da39fb.X >= __obf_dfa162ecb02c063a.X {
		__obf_963d969f08b1d119 = -1
	}
	if __obf_fa6a97d0e7da39fb.Y >= __obf_dfa162ecb02c063a.Y {
		__obf_8ff9861776e145d0 = -1
	}
	__obf_ffc118ace13b322b := __obf_2ab04c1711998195 - __obf_45d338c45fe3468d
	for {
		__obf_91a96cde82d83f16.Set(__obf_fa6a97d0e7da39fb.X, __obf_fa6a97d0e7da39fb.Y, __obf_20b69069a99344ee)
		__obf_91a96cde82d83f16.Set(__obf_fa6a97d0e7da39fb.X+1, __obf_fa6a97d0e7da39fb.Y, __obf_20b69069a99344ee)
		__obf_91a96cde82d83f16.Set(__obf_fa6a97d0e7da39fb.X-1, __obf_fa6a97d0e7da39fb.Y, __obf_20b69069a99344ee)
		__obf_91a96cde82d83f16.Set(__obf_fa6a97d0e7da39fb.X+2, __obf_fa6a97d0e7da39fb.Y, __obf_20b69069a99344ee)
		__obf_91a96cde82d83f16.Set(__obf_fa6a97d0e7da39fb.X-2, __obf_fa6a97d0e7da39fb.Y, __obf_20b69069a99344ee)
		if __obf_fa6a97d0e7da39fb.X == __obf_dfa162ecb02c063a.X && __obf_fa6a97d0e7da39fb.Y == __obf_dfa162ecb02c063a.Y {
			return
		}
		__obf_ee57c21abc6a701c := __obf_ffc118ace13b322b * 2
		if __obf_ee57c21abc6a701c > -__obf_45d338c45fe3468d {
			__obf_ffc118ace13b322b -= __obf_45d338c45fe3468d
			__obf_fa6a97d0e7da39fb.X += __obf_963d969f08b1d119
		}
		if __obf_ee57c21abc6a701c < __obf_2ab04c1711998195 {
			__obf_ffc118ace13b322b += __obf_2ab04c1711998195
			__obf_fa6a97d0e7da39fb.Y += __obf_8ff9861776e145d0
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_91a96cde82d83f16 *Palette) AngleSwapPoint(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f, __obf_c3f6882d3032b52f, __obf_2c1097e2e9ce1f24 float64) (__obf_381a3275012ca2b3, __obf_889390e07de55833 float64) {
	__obf_c87d6c7ba4a30d82 -= __obf_c3f6882d3032b52f
	__obf_f8fead3d682d659f = __obf_c3f6882d3032b52f - __obf_f8fead3d682d659f
	__obf_84e9f6f27c5212a1 := math.Sin(__obf_2c1097e2e9ce1f24 * (math.Pi / 180))
	__obf_7c031f062d5b4f62 := math.Cos(__obf_2c1097e2e9ce1f24 * (math.Pi / 180))
	__obf_381a3275012ca2b3 = __obf_c87d6c7ba4a30d82*__obf_7c031f062d5b4f62 + __obf_f8fead3d682d659f*__obf_84e9f6f27c5212a1
	__obf_889390e07de55833 = -__obf_c87d6c7ba4a30d82*__obf_84e9f6f27c5212a1 + __obf_f8fead3d682d659f*__obf_7c031f062d5b4f62
	__obf_381a3275012ca2b3 += __obf_c3f6882d3032b52f
	__obf_889390e07de55833 = __obf_c3f6882d3032b52f - __obf_889390e07de55833
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_91a96cde82d83f16 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_d14caaad13625683 := __obf_91a96cde82d83f16.Bounds().Max.X
	__obf_ff9981d45d8deb98 := __obf_91a96cde82d83f16.Bounds().Max.Y
	__obf_2112cc534df9dd1a := __obf_d14caaad13625683
	__obf_5bcc151bb94dd6aa := 0
	__obf_f49db9de609cf423 := __obf_ff9981d45d8deb98
	__obf_e6c5d6e5feaf61f3 := 0
	for __obf_c87d6c7ba4a30d82 := range __obf_d14caaad13625683 {
		for __obf_f8fead3d682d659f := range __obf_ff9981d45d8deb98 {
			__obf_12c04581a749b5b3 := __obf_91a96cde82d83f16.At(__obf_c87d6c7ba4a30d82, __obf_f8fead3d682d659f)
			if _, _, _, __obf_91f28fffbdd83b09 := __obf_12c04581a749b5b3.RGBA(); __obf_91f28fffbdd83b09 > 0 {
				if __obf_c87d6c7ba4a30d82 < __obf_2112cc534df9dd1a {
					__obf_2112cc534df9dd1a = __obf_c87d6c7ba4a30d82
				}
				if __obf_c87d6c7ba4a30d82 > __obf_5bcc151bb94dd6aa {
					__obf_5bcc151bb94dd6aa = __obf_c87d6c7ba4a30d82
				}

				if __obf_f8fead3d682d659f < __obf_f49db9de609cf423 {
					__obf_f49db9de609cf423 = __obf_f8fead3d682d659f
				}
				if __obf_f8fead3d682d659f > __obf_e6c5d6e5feaf61f3 {
					__obf_e6c5d6e5feaf61f3 = __obf_f8fead3d682d659f
				}
			}
		}
	}

	__obf_2112cc534df9dd1a = int(max(0, float64(__obf_2112cc534df9dd1a-2)))
	__obf_5bcc151bb94dd6aa = int(min(float64(__obf_d14caaad13625683), float64(__obf_5bcc151bb94dd6aa+2)))
	__obf_f49db9de609cf423 = int(max(0, float64(__obf_f49db9de609cf423-2)))
	__obf_e6c5d6e5feaf61f3 = int(min(float64(__obf_ff9981d45d8deb98), float64(__obf_e6c5d6e5feaf61f3+2)))

	return &AreaRect{
		__obf_2112cc534df9dd1a,
		__obf_5bcc151bb94dd6aa,
		__obf_f49db9de609cf423,
		__obf_e6c5d6e5feaf61f3,
	}
}
