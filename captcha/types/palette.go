package __obf_89363901d8df63bc

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_a89def74d931b834 image.Rectangle, __obf_ce0f98807b85f11c color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_a89def74d931b834,

			// palette struct for palette canvas
			__obf_ce0f98807b85f11c),
	}
}

type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_ce0f98807b85f11c *Palette) Rotate(__obf_1f7091190512fbb8 int) {
	if __obf_1f7091190512fbb8 == 0 {
		return
	}
	__obf_228e6b18ef94d799 := __obf_ce0f98807b85f11c
	__obf_f10777ec7e9bb6bc := __obf_228e6b18ef94d799.Bounds().Max.X
	__obf_e16e2348abf5243e := __obf_228e6b18ef94d799.Bounds().Max.Y
	__obf_a89def74d931b834 := __obf_f10777ec7e9bb6bc / 2
	__obf_7f7ec84291a34c02 := image.NewPaletted(image.Rect(0, 0, __obf_f10777ec7e9bb6bc, __obf_e16e2348abf5243e), __obf_228e6b18ef94d799.Palette)
	for __obf_4d548ce157fe2d7b := 0; __obf_4d548ce157fe2d7b <= __obf_7f7ec84291a34c02.Bounds().Max.X; __obf_4d548ce157fe2d7b++ {
		for __obf_e5fce044456d5b92 := 0; __obf_e5fce044456d5b92 <= __obf_7f7ec84291a34c02.Bounds().Max.Y; __obf_e5fce044456d5b92++ {
			__obf_84283d311c9ec184, __obf_6aaef088895dd60b := __obf_ce0f98807b85f11c.AngleSwapPoint(float64(__obf_4d548ce157fe2d7b), float64(__obf_e5fce044456d5b92), float64(__obf_a89def74d931b834), float64(__obf_1f7091190512fbb8))
			__obf_7f7ec84291a34c02.
				SetColorIndex(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_228e6b18ef94d799.ColorIndexAt(int(__obf_84283d311c9ec184), int(__obf_6aaef088895dd60b)))
		}
	}
	__obf_030657fbacf1ee00 := __obf_7f7ec84291a34c02.Bounds().Max.X
	__obf_60d018dc204e7967 := __obf_7f7ec84291a34c02.Bounds().Max.Y
	for __obf_4d548ce157fe2d7b := range __obf_030657fbacf1ee00 {
		for __obf_e5fce044456d5b92 := range __obf_60d018dc204e7967 {
			__obf_ce0f98807b85f11c.
				SetColorIndex(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92,

					// DrawCircle draws a circle
					__obf_7f7ec84291a34c02.ColorIndexAt(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92))
		}
	}
}

func (__obf_ce0f98807b85f11c *Palette) DrawCircle(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_126525ac437289c4 int, __obf_8f1cd35f635f319d color.RGBA) {
	__obf_9ab78ba63bb78d83 := 1 - __obf_126525ac437289c4
	__obf_3768018153c7018a := 1
	__obf_ad178bdc83492b1a := -2 * __obf_126525ac437289c4
	__obf_f538188d1686f81b := 0
	__obf_73b8f8029c024ce8 := __obf_126525ac437289c4
	__obf_ce0f98807b85f11c.
		Set(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92+__obf_126525ac437289c4, __obf_8f1cd35f635f319d)
	__obf_ce0f98807b85f11c.
		Set(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92-__obf_126525ac437289c4, __obf_8f1cd35f635f319d)
	__obf_ce0f98807b85f11c.
		DrawHorizLine(__obf_4d548ce157fe2d7b-__obf_126525ac437289c4, __obf_4d548ce157fe2d7b+__obf_126525ac437289c4, __obf_e5fce044456d5b92, __obf_8f1cd35f635f319d)

	for __obf_f538188d1686f81b < __obf_73b8f8029c024ce8 {
		if __obf_9ab78ba63bb78d83 >= 0 {
			__obf_73b8f8029c024ce8--
			__obf_ad178bdc83492b1a += 2
			__obf_9ab78ba63bb78d83 += __obf_ad178bdc83492b1a
		}
		__obf_f538188d1686f81b++
		__obf_3768018153c7018a += 2
		__obf_9ab78ba63bb78d83 += __obf_3768018153c7018a
		__obf_ce0f98807b85f11c.
			DrawHorizLine(__obf_4d548ce157fe2d7b-__obf_f538188d1686f81b, __obf_4d548ce157fe2d7b+__obf_f538188d1686f81b, __obf_e5fce044456d5b92+__obf_73b8f8029c024ce8, __obf_8f1cd35f635f319d)
		__obf_ce0f98807b85f11c.
			DrawHorizLine(__obf_4d548ce157fe2d7b-__obf_f538188d1686f81b, __obf_4d548ce157fe2d7b+__obf_f538188d1686f81b, __obf_e5fce044456d5b92-__obf_73b8f8029c024ce8, __obf_8f1cd35f635f319d)
		__obf_ce0f98807b85f11c.
			DrawHorizLine(__obf_4d548ce157fe2d7b-__obf_73b8f8029c024ce8, __obf_4d548ce157fe2d7b+__obf_73b8f8029c024ce8, __obf_e5fce044456d5b92+__obf_f538188d1686f81b, __obf_8f1cd35f635f319d)
		__obf_ce0f98807b85f11c.
			DrawHorizLine(__obf_4d548ce157fe2d7b-__obf_73b8f8029c024ce8, __obf_4d548ce157fe2d7b+__obf_73b8f8029c024ce8,

				// DrawHorizLine draws a horizontal line
				__obf_e5fce044456d5b92-__obf_f538188d1686f81b, __obf_8f1cd35f635f319d)
	}
}

func (__obf_ce0f98807b85f11c *Palette) DrawHorizLine(__obf_9b4076142363436e, __obf_874e344711b134fa, __obf_e5fce044456d5b92 int, __obf_8f1cd35f635f319d color.RGBA) {
	for __obf_4d548ce157fe2d7b := __obf_9b4076142363436e; __obf_4d548ce157fe2d7b <= __obf_874e344711b134fa; __obf_4d548ce157fe2d7b++ {
		__obf_ce0f98807b85f11c.
			Set(__obf_4d548ce157fe2d7b,

				// Distort distorts the canvas
				__obf_e5fce044456d5b92, __obf_8f1cd35f635f319d)
	}
}

func (__obf_ce0f98807b85f11c *Palette) Distort(__obf_e2aea071bbcfa3fe float64, __obf_7a8a26d359333515 float64) {
	__obf_672a1e7045b302e8 := __obf_ce0f98807b85f11c.Bounds().Max.X
	__obf_64febf9a6792142f := __obf_ce0f98807b85f11c.Bounds().Max.Y
	__obf_2715763b241b5450 := NewPalette(image.Rect(0, 0, __obf_672a1e7045b302e8, __obf_64febf9a6792142f), __obf_ce0f98807b85f11c.Palette)
	__obf_0996b35bb1f98ea1 := 2.0 * math.Pi / __obf_7a8a26d359333515
	for __obf_4d548ce157fe2d7b := range __obf_672a1e7045b302e8 {
		for __obf_e5fce044456d5b92 := 0; __obf_e5fce044456d5b92 < __obf_64febf9a6792142f; __obf_e5fce044456d5b92++ {
			__obf_f538188d1686f81b := __obf_e2aea071bbcfa3fe * math.Sin(float64(__obf_e5fce044456d5b92)*__obf_0996b35bb1f98ea1)
			__obf_73b8f8029c024ce8 := __obf_e2aea071bbcfa3fe * math.Cos(float64(__obf_4d548ce157fe2d7b)*__obf_0996b35bb1f98ea1)
			__obf_2715763b241b5450.
				SetColorIndex(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_ce0f98807b85f11c.ColorIndexAt(__obf_4d548ce157fe2d7b+int(__obf_f538188d1686f81b), __obf_e5fce044456d5b92+int(__obf_73b8f8029c024ce8)))
		}
	}
	__obf_030657fbacf1ee00 := __obf_2715763b241b5450.Bounds().Max.X
	__obf_60d018dc204e7967 := __obf_2715763b241b5450.Bounds().Max.Y
	for __obf_4d548ce157fe2d7b := range __obf_030657fbacf1ee00 {
		for __obf_e5fce044456d5b92 := 0; __obf_e5fce044456d5b92 < __obf_60d018dc204e7967; __obf_e5fce044456d5b92++ {
			__obf_ce0f98807b85f11c.
				SetColorIndex(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_2715763b241b5450.ColorIndexAt(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92))
		}
	}
	__obf_2715763b241b5450.
		Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_ce0f98807b85f11c *Palette) DrawBeeline(__obf_5d939491edc59ad6 image.Point, __obf_7aa2afcd49769389 image.Point, __obf_405c1cd87657693e color.RGBA) {
	__obf_0996b35bb1f98ea1 := math.Abs(float64(__obf_5d939491edc59ad6.X - __obf_7aa2afcd49769389.X))
	__obf_68d36fd00e636e53 := math.Abs(float64(__obf_7aa2afcd49769389.Y - __obf_5d939491edc59ad6.Y))
	__obf_8d57671244f1c5aa, __obf_8b053c23a483a4da := 1, 1
	if __obf_5d939491edc59ad6.X >= __obf_7aa2afcd49769389.X {
		__obf_8d57671244f1c5aa = -1
	}
	if __obf_5d939491edc59ad6.Y >= __obf_7aa2afcd49769389.Y {
		__obf_8b053c23a483a4da = -1
	}
	__obf_7b78cb8578d9f1c2 := __obf_0996b35bb1f98ea1 - __obf_68d36fd00e636e53
	for {
		__obf_ce0f98807b85f11c.
			Set(__obf_5d939491edc59ad6.X, __obf_5d939491edc59ad6.Y, __obf_405c1cd87657693e)
		__obf_ce0f98807b85f11c.
			Set(__obf_5d939491edc59ad6.X+1, __obf_5d939491edc59ad6.Y, __obf_405c1cd87657693e)
		__obf_ce0f98807b85f11c.
			Set(__obf_5d939491edc59ad6.X-1, __obf_5d939491edc59ad6.Y, __obf_405c1cd87657693e)
		__obf_ce0f98807b85f11c.
			Set(__obf_5d939491edc59ad6.X+2, __obf_5d939491edc59ad6.Y, __obf_405c1cd87657693e)
		__obf_ce0f98807b85f11c.
			Set(__obf_5d939491edc59ad6.X-2, __obf_5d939491edc59ad6.Y, __obf_405c1cd87657693e)
		if __obf_5d939491edc59ad6.X == __obf_7aa2afcd49769389.X && __obf_5d939491edc59ad6.Y == __obf_7aa2afcd49769389.Y {
			return
		}
		__obf_b857a4c078334e92 := __obf_7b78cb8578d9f1c2 * 2
		if __obf_b857a4c078334e92 > -__obf_68d36fd00e636e53 {
			__obf_7b78cb8578d9f1c2 -= __obf_68d36fd00e636e53
			__obf_5d939491edc59ad6.
				X += __obf_8d57671244f1c5aa
		}
		if __obf_b857a4c078334e92 < __obf_0996b35bb1f98ea1 {
			__obf_7b78cb8578d9f1c2 += __obf_0996b35bb1f98ea1
			__obf_5d939491edc59ad6.
				Y += __obf_8b053c23a483a4da
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_ce0f98807b85f11c *Palette) AngleSwapPoint(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92, __obf_a89def74d931b834, __obf_1f7091190512fbb8 float64) (__obf_c12017298cd50a12, __obf_34b6ecb300df7fce float64) {
	__obf_4d548ce157fe2d7b -= __obf_a89def74d931b834
	__obf_e5fce044456d5b92 = __obf_a89def74d931b834 - __obf_e5fce044456d5b92
	__obf_98c76658d423994d := math.Sin(__obf_1f7091190512fbb8 * (math.Pi / 180))
	__obf_99aa9bf9b9a405f7 := math.Cos(__obf_1f7091190512fbb8 * (math.Pi / 180))
	__obf_c12017298cd50a12 = __obf_4d548ce157fe2d7b*__obf_99aa9bf9b9a405f7 + __obf_e5fce044456d5b92*__obf_98c76658d423994d
	__obf_34b6ecb300df7fce = -__obf_4d548ce157fe2d7b*__obf_98c76658d423994d + __obf_e5fce044456d5b92*__obf_99aa9bf9b9a405f7
	__obf_c12017298cd50a12 += __obf_a89def74d931b834
	__obf_34b6ecb300df7fce = __obf_a89def74d931b834 - __obf_34b6ecb300df7fce
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_ce0f98807b85f11c *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_030657fbacf1ee00 := __obf_ce0f98807b85f11c.Bounds().Max.X
	__obf_60d018dc204e7967 := __obf_ce0f98807b85f11c.Bounds().Max.Y
	__obf_ef02f5aef30dfa3c := __obf_030657fbacf1ee00
	__obf_014a5112e2c2a925 := 0
	__obf_081a97e9258e5ee5 := __obf_60d018dc204e7967
	__obf_f14c1ac944ab90a0 := 0
	for __obf_4d548ce157fe2d7b := range __obf_030657fbacf1ee00 {
		for __obf_e5fce044456d5b92 := range __obf_60d018dc204e7967 {
			__obf_afee39acf6210256 := __obf_ce0f98807b85f11c.At(__obf_4d548ce157fe2d7b, __obf_e5fce044456d5b92)
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

	return &AreaRect{__obf_ef02f5aef30dfa3c, __obf_014a5112e2c2a925, __obf_081a97e9258e5ee5, __obf_f14c1ac944ab90a0}
}
