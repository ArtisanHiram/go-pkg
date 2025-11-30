package __obf_738b46210fdb4199

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_c09045abdc5d3619 image.Rectangle, __obf_24dca18af6feb178 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_c09045abdc5d3619,

			// palette struct for palette canvas
			__obf_24dca18af6feb178),
	}
}

type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_24dca18af6feb178 *Palette) Rotate(__obf_3a2dad55c354b1d6 int) {
	if __obf_3a2dad55c354b1d6 == 0 {
		return
	}
	__obf_f311aa0c4c36a5d7 := __obf_24dca18af6feb178
	__obf_da28b3bb7dbcda2b := __obf_f311aa0c4c36a5d7.Bounds().Max.X
	__obf_fa37a5ce5b7fe70f := __obf_f311aa0c4c36a5d7.Bounds().Max.Y
	__obf_c09045abdc5d3619 := __obf_da28b3bb7dbcda2b / 2
	__obf_cc502ad9d0d21dd8 := image.NewPaletted(image.Rect(0, 0, __obf_da28b3bb7dbcda2b, __obf_fa37a5ce5b7fe70f), __obf_f311aa0c4c36a5d7.Palette)
	for __obf_bd372f9dc9e16d75 := 0; __obf_bd372f9dc9e16d75 <= __obf_cc502ad9d0d21dd8.Bounds().Max.X; __obf_bd372f9dc9e16d75++ {
		for __obf_a69b4d01937bfbfc := 0; __obf_a69b4d01937bfbfc <= __obf_cc502ad9d0d21dd8.Bounds().Max.Y; __obf_a69b4d01937bfbfc++ {
			__obf_aaefdfdd53ea9958, __obf_9e4dc95d0b70d96a := __obf_24dca18af6feb178.AngleSwapPoint(float64(__obf_bd372f9dc9e16d75), float64(__obf_a69b4d01937bfbfc), float64(__obf_c09045abdc5d3619), float64(__obf_3a2dad55c354b1d6))
			__obf_cc502ad9d0d21dd8.
				SetColorIndex(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_f311aa0c4c36a5d7.ColorIndexAt(int(__obf_aaefdfdd53ea9958), int(__obf_9e4dc95d0b70d96a)))
		}
	}
	__obf_d17e3b2fb9b48db3 := __obf_cc502ad9d0d21dd8.Bounds().Max.X
	__obf_e0ad449835d092e9 := __obf_cc502ad9d0d21dd8.Bounds().Max.Y
	for __obf_bd372f9dc9e16d75 := range __obf_d17e3b2fb9b48db3 {
		for __obf_a69b4d01937bfbfc := range __obf_e0ad449835d092e9 {
			__obf_24dca18af6feb178.
				SetColorIndex(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc,

					// DrawCircle draws a circle
					__obf_cc502ad9d0d21dd8.ColorIndexAt(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc))
		}
	}
}

func (__obf_24dca18af6feb178 *Palette) DrawCircle(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_6b93bc026df65fff int, __obf_cf060e7b58cbd8c5 color.RGBA) {
	__obf_73086ebe1aad30bd := 1 - __obf_6b93bc026df65fff
	__obf_b17acf1df0c592bf := 1
	__obf_ce47915aa5adb906 := -2 * __obf_6b93bc026df65fff
	__obf_5951423922edb364 := 0
	__obf_95281919c4d87261 := __obf_6b93bc026df65fff
	__obf_24dca18af6feb178.
		Set(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc+__obf_6b93bc026df65fff, __obf_cf060e7b58cbd8c5)
	__obf_24dca18af6feb178.
		Set(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc-__obf_6b93bc026df65fff, __obf_cf060e7b58cbd8c5)
	__obf_24dca18af6feb178.
		DrawHorizLine(__obf_bd372f9dc9e16d75-__obf_6b93bc026df65fff, __obf_bd372f9dc9e16d75+__obf_6b93bc026df65fff, __obf_a69b4d01937bfbfc, __obf_cf060e7b58cbd8c5)

	for __obf_5951423922edb364 < __obf_95281919c4d87261 {
		if __obf_73086ebe1aad30bd >= 0 {
			__obf_95281919c4d87261--
			__obf_ce47915aa5adb906 += 2
			__obf_73086ebe1aad30bd += __obf_ce47915aa5adb906
		}
		__obf_5951423922edb364++
		__obf_b17acf1df0c592bf += 2
		__obf_73086ebe1aad30bd += __obf_b17acf1df0c592bf
		__obf_24dca18af6feb178.
			DrawHorizLine(__obf_bd372f9dc9e16d75-__obf_5951423922edb364, __obf_bd372f9dc9e16d75+__obf_5951423922edb364, __obf_a69b4d01937bfbfc+__obf_95281919c4d87261, __obf_cf060e7b58cbd8c5)
		__obf_24dca18af6feb178.
			DrawHorizLine(__obf_bd372f9dc9e16d75-__obf_5951423922edb364, __obf_bd372f9dc9e16d75+__obf_5951423922edb364, __obf_a69b4d01937bfbfc-__obf_95281919c4d87261, __obf_cf060e7b58cbd8c5)
		__obf_24dca18af6feb178.
			DrawHorizLine(__obf_bd372f9dc9e16d75-__obf_95281919c4d87261, __obf_bd372f9dc9e16d75+__obf_95281919c4d87261, __obf_a69b4d01937bfbfc+__obf_5951423922edb364, __obf_cf060e7b58cbd8c5)
		__obf_24dca18af6feb178.
			DrawHorizLine(__obf_bd372f9dc9e16d75-__obf_95281919c4d87261, __obf_bd372f9dc9e16d75+__obf_95281919c4d87261,

				// DrawHorizLine draws a horizontal line
				__obf_a69b4d01937bfbfc-__obf_5951423922edb364, __obf_cf060e7b58cbd8c5)
	}
}

func (__obf_24dca18af6feb178 *Palette) DrawHorizLine(__obf_1fc3777fd512c47a, __obf_b51b4e665b0fad85, __obf_a69b4d01937bfbfc int, __obf_cf060e7b58cbd8c5 color.RGBA) {
	for __obf_bd372f9dc9e16d75 := __obf_1fc3777fd512c47a; __obf_bd372f9dc9e16d75 <= __obf_b51b4e665b0fad85; __obf_bd372f9dc9e16d75++ {
		__obf_24dca18af6feb178.
			Set(__obf_bd372f9dc9e16d75,

				// Distort distorts the canvas
				__obf_a69b4d01937bfbfc, __obf_cf060e7b58cbd8c5)
	}
}

func (__obf_24dca18af6feb178 *Palette) Distort(__obf_3edcdb7c63c485fe float64, __obf_d7226a471a8a0a36 float64) {
	__obf_ccd0c8ac1f16e3bd := __obf_24dca18af6feb178.Bounds().Max.X
	__obf_90a966c355c00010 := __obf_24dca18af6feb178.Bounds().Max.Y
	__obf_e9eccf4e599e5cbc := NewPalette(image.Rect(0, 0, __obf_ccd0c8ac1f16e3bd, __obf_90a966c355c00010), __obf_24dca18af6feb178.Palette)
	__obf_936088d1b4a69899 := 2.0 * math.Pi / __obf_d7226a471a8a0a36
	for __obf_bd372f9dc9e16d75 := range __obf_ccd0c8ac1f16e3bd {
		for __obf_a69b4d01937bfbfc := 0; __obf_a69b4d01937bfbfc < __obf_90a966c355c00010; __obf_a69b4d01937bfbfc++ {
			__obf_5951423922edb364 := __obf_3edcdb7c63c485fe * math.Sin(float64(__obf_a69b4d01937bfbfc)*__obf_936088d1b4a69899)
			__obf_95281919c4d87261 := __obf_3edcdb7c63c485fe * math.Cos(float64(__obf_bd372f9dc9e16d75)*__obf_936088d1b4a69899)
			__obf_e9eccf4e599e5cbc.
				SetColorIndex(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_24dca18af6feb178.ColorIndexAt(__obf_bd372f9dc9e16d75+int(__obf_5951423922edb364), __obf_a69b4d01937bfbfc+int(__obf_95281919c4d87261)))
		}
	}
	__obf_d17e3b2fb9b48db3 := __obf_e9eccf4e599e5cbc.Bounds().Max.X
	__obf_e0ad449835d092e9 := __obf_e9eccf4e599e5cbc.Bounds().Max.Y
	for __obf_bd372f9dc9e16d75 := range __obf_d17e3b2fb9b48db3 {
		for __obf_a69b4d01937bfbfc := 0; __obf_a69b4d01937bfbfc < __obf_e0ad449835d092e9; __obf_a69b4d01937bfbfc++ {
			__obf_24dca18af6feb178.
				SetColorIndex(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_e9eccf4e599e5cbc.ColorIndexAt(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc))
		}
	}
	__obf_e9eccf4e599e5cbc.
		Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_24dca18af6feb178 *Palette) DrawBeeline(__obf_23281d9c04e41cb6 image.Point, __obf_f72b82b64fc402eb image.Point, __obf_ff10c1bbb8d2fdda color.RGBA) {
	__obf_936088d1b4a69899 := math.Abs(float64(__obf_23281d9c04e41cb6.X - __obf_f72b82b64fc402eb.X))
	__obf_141144958ec0757c := math.Abs(float64(__obf_f72b82b64fc402eb.Y - __obf_23281d9c04e41cb6.Y))
	__obf_53be3572988aceca, __obf_ed348ff651e7f60e := 1, 1
	if __obf_23281d9c04e41cb6.X >= __obf_f72b82b64fc402eb.X {
		__obf_53be3572988aceca = -1
	}
	if __obf_23281d9c04e41cb6.Y >= __obf_f72b82b64fc402eb.Y {
		__obf_ed348ff651e7f60e = -1
	}
	__obf_95d84a130cd584e5 := __obf_936088d1b4a69899 - __obf_141144958ec0757c
	for {
		__obf_24dca18af6feb178.
			Set(__obf_23281d9c04e41cb6.X, __obf_23281d9c04e41cb6.Y, __obf_ff10c1bbb8d2fdda)
		__obf_24dca18af6feb178.
			Set(__obf_23281d9c04e41cb6.X+1, __obf_23281d9c04e41cb6.Y, __obf_ff10c1bbb8d2fdda)
		__obf_24dca18af6feb178.
			Set(__obf_23281d9c04e41cb6.X-1, __obf_23281d9c04e41cb6.Y, __obf_ff10c1bbb8d2fdda)
		__obf_24dca18af6feb178.
			Set(__obf_23281d9c04e41cb6.X+2, __obf_23281d9c04e41cb6.Y, __obf_ff10c1bbb8d2fdda)
		__obf_24dca18af6feb178.
			Set(__obf_23281d9c04e41cb6.X-2, __obf_23281d9c04e41cb6.Y, __obf_ff10c1bbb8d2fdda)
		if __obf_23281d9c04e41cb6.X == __obf_f72b82b64fc402eb.X && __obf_23281d9c04e41cb6.Y == __obf_f72b82b64fc402eb.Y {
			return
		}
		__obf_258c750098c7731d := __obf_95d84a130cd584e5 * 2
		if __obf_258c750098c7731d > -__obf_141144958ec0757c {
			__obf_95d84a130cd584e5 -= __obf_141144958ec0757c
			__obf_23281d9c04e41cb6.
				X += __obf_53be3572988aceca
		}
		if __obf_258c750098c7731d < __obf_936088d1b4a69899 {
			__obf_95d84a130cd584e5 += __obf_936088d1b4a69899
			__obf_23281d9c04e41cb6.
				Y += __obf_ed348ff651e7f60e
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_24dca18af6feb178 *Palette) AngleSwapPoint(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc, __obf_c09045abdc5d3619, __obf_3a2dad55c354b1d6 float64) (__obf_d4a0b02e6b76c897, __obf_2c96eb9d5bf6b705 float64) {
	__obf_bd372f9dc9e16d75 -= __obf_c09045abdc5d3619
	__obf_a69b4d01937bfbfc = __obf_c09045abdc5d3619 - __obf_a69b4d01937bfbfc
	__obf_00eeb677fb04234b := math.Sin(__obf_3a2dad55c354b1d6 * (math.Pi / 180))
	__obf_cd19e570ab31e612 := math.Cos(__obf_3a2dad55c354b1d6 * (math.Pi / 180))
	__obf_d4a0b02e6b76c897 = __obf_bd372f9dc9e16d75*__obf_cd19e570ab31e612 + __obf_a69b4d01937bfbfc*__obf_00eeb677fb04234b
	__obf_2c96eb9d5bf6b705 = -__obf_bd372f9dc9e16d75*__obf_00eeb677fb04234b + __obf_a69b4d01937bfbfc*__obf_cd19e570ab31e612
	__obf_d4a0b02e6b76c897 += __obf_c09045abdc5d3619
	__obf_2c96eb9d5bf6b705 = __obf_c09045abdc5d3619 - __obf_2c96eb9d5bf6b705
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_24dca18af6feb178 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_d17e3b2fb9b48db3 := __obf_24dca18af6feb178.Bounds().Max.X
	__obf_e0ad449835d092e9 := __obf_24dca18af6feb178.Bounds().Max.Y
	__obf_ba97c5a5606ddd7e := __obf_d17e3b2fb9b48db3
	__obf_c6b746bbb5d1f238 := 0
	__obf_36f26cd4c2083cc6 := __obf_e0ad449835d092e9
	__obf_7d40f40942b754c0 := 0
	for __obf_bd372f9dc9e16d75 := range __obf_d17e3b2fb9b48db3 {
		for __obf_a69b4d01937bfbfc := range __obf_e0ad449835d092e9 {
			__obf_e4c726f405e92dc2 := __obf_24dca18af6feb178.At(__obf_bd372f9dc9e16d75, __obf_a69b4d01937bfbfc)
			if _, _, _, __obf_5eeb26a04aed8320 := __obf_e4c726f405e92dc2.RGBA(); __obf_5eeb26a04aed8320 > 0 {
				if __obf_bd372f9dc9e16d75 < __obf_ba97c5a5606ddd7e {
					__obf_ba97c5a5606ddd7e = __obf_bd372f9dc9e16d75
				}
				if __obf_bd372f9dc9e16d75 > __obf_c6b746bbb5d1f238 {
					__obf_c6b746bbb5d1f238 = __obf_bd372f9dc9e16d75
				}

				if __obf_a69b4d01937bfbfc < __obf_36f26cd4c2083cc6 {
					__obf_36f26cd4c2083cc6 = __obf_a69b4d01937bfbfc
				}
				if __obf_a69b4d01937bfbfc > __obf_7d40f40942b754c0 {
					__obf_7d40f40942b754c0 = __obf_a69b4d01937bfbfc
				}
			}
		}
	}
	__obf_ba97c5a5606ddd7e = int(max(0, float64(__obf_ba97c5a5606ddd7e-2)))
	__obf_c6b746bbb5d1f238 = int(min(float64(__obf_d17e3b2fb9b48db3), float64(__obf_c6b746bbb5d1f238+2)))
	__obf_36f26cd4c2083cc6 = int(max(0, float64(__obf_36f26cd4c2083cc6-2)))
	__obf_7d40f40942b754c0 = int(min(float64(__obf_e0ad449835d092e9), float64(__obf_7d40f40942b754c0+2)))

	return &AreaRect{__obf_ba97c5a5606ddd7e, __obf_c6b746bbb5d1f238, __obf_36f26cd4c2083cc6, __obf_7d40f40942b754c0}
}
