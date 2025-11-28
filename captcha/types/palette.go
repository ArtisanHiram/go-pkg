package __obf_b0f3614e51931450

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_698deef70980b293 image.Rectangle, __obf_f34639443602c065 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_698deef70980b293, __obf_f34639443602c065),
	}
}

// palette struct for palette canvas
type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_f34639443602c065 *Palette) Rotate(__obf_31d599a870027db7 int) {
	if __obf_31d599a870027db7 == 0 {
		return
	}

	__obf_9e3644967771077a := __obf_f34639443602c065
	__obf_6174b6789f9528b5 := __obf_9e3644967771077a.Bounds().Max.X
	__obf_bea3c706ffff9ef0 := __obf_9e3644967771077a.Bounds().Max.Y
	__obf_698deef70980b293 := __obf_6174b6789f9528b5 / 2
	__obf_b692a67d50ef7f3c := image.NewPaletted(image.Rect(0, 0, __obf_6174b6789f9528b5, __obf_bea3c706ffff9ef0), __obf_9e3644967771077a.Palette)
	for __obf_80fae96c5947b2de := 0; __obf_80fae96c5947b2de <= __obf_b692a67d50ef7f3c.Bounds().Max.X; __obf_80fae96c5947b2de++ {
		for __obf_431a48e7317ece6d := 0; __obf_431a48e7317ece6d <= __obf_b692a67d50ef7f3c.Bounds().Max.Y; __obf_431a48e7317ece6d++ {
			__obf_f977f99c59c221c3, __obf_ef7d1e31bd77d21f := __obf_f34639443602c065.AngleSwapPoint(float64(__obf_80fae96c5947b2de), float64(__obf_431a48e7317ece6d), float64(__obf_698deef70980b293), float64(__obf_31d599a870027db7))
			__obf_b692a67d50ef7f3c.SetColorIndex(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_9e3644967771077a.ColorIndexAt(int(__obf_f977f99c59c221c3), int(__obf_ef7d1e31bd77d21f)))
		}
	}

	__obf_e61ea68b462e25a1 := __obf_b692a67d50ef7f3c.Bounds().Max.X
	__obf_5cc37636829791d6 := __obf_b692a67d50ef7f3c.Bounds().Max.Y
	for __obf_80fae96c5947b2de := range __obf_e61ea68b462e25a1 {
		for __obf_431a48e7317ece6d := range __obf_5cc37636829791d6 {
			__obf_f34639443602c065.SetColorIndex(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_b692a67d50ef7f3c.ColorIndexAt(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d))
		}
	}
}

// DrawCircle draws a circle
func (__obf_f34639443602c065 *Palette) DrawCircle(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_b290e37944a6ece3 int, __obf_587b9fceec62883e color.RGBA) {
	__obf_ad69fb6ba990618d := 1 - __obf_b290e37944a6ece3
	__obf_5f954fb0e3effca3 := 1
	__obf_4553ea39f5ff68fa := -2 * __obf_b290e37944a6ece3
	__obf_ecb441a4ce96b36a := 0
	__obf_1a7ab87e567bf319 := __obf_b290e37944a6ece3

	__obf_f34639443602c065.Set(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d+__obf_b290e37944a6ece3, __obf_587b9fceec62883e)
	__obf_f34639443602c065.Set(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d-__obf_b290e37944a6ece3, __obf_587b9fceec62883e)
	__obf_f34639443602c065.DrawHorizLine(__obf_80fae96c5947b2de-__obf_b290e37944a6ece3, __obf_80fae96c5947b2de+__obf_b290e37944a6ece3, __obf_431a48e7317ece6d, __obf_587b9fceec62883e)

	for __obf_ecb441a4ce96b36a < __obf_1a7ab87e567bf319 {
		if __obf_ad69fb6ba990618d >= 0 {
			__obf_1a7ab87e567bf319--
			__obf_4553ea39f5ff68fa += 2
			__obf_ad69fb6ba990618d += __obf_4553ea39f5ff68fa
		}
		__obf_ecb441a4ce96b36a++
		__obf_5f954fb0e3effca3 += 2
		__obf_ad69fb6ba990618d += __obf_5f954fb0e3effca3
		__obf_f34639443602c065.DrawHorizLine(__obf_80fae96c5947b2de-__obf_ecb441a4ce96b36a, __obf_80fae96c5947b2de+__obf_ecb441a4ce96b36a, __obf_431a48e7317ece6d+__obf_1a7ab87e567bf319, __obf_587b9fceec62883e)
		__obf_f34639443602c065.DrawHorizLine(__obf_80fae96c5947b2de-__obf_ecb441a4ce96b36a, __obf_80fae96c5947b2de+__obf_ecb441a4ce96b36a, __obf_431a48e7317ece6d-__obf_1a7ab87e567bf319, __obf_587b9fceec62883e)
		__obf_f34639443602c065.DrawHorizLine(__obf_80fae96c5947b2de-__obf_1a7ab87e567bf319, __obf_80fae96c5947b2de+__obf_1a7ab87e567bf319, __obf_431a48e7317ece6d+__obf_ecb441a4ce96b36a, __obf_587b9fceec62883e)
		__obf_f34639443602c065.DrawHorizLine(__obf_80fae96c5947b2de-__obf_1a7ab87e567bf319, __obf_80fae96c5947b2de+__obf_1a7ab87e567bf319, __obf_431a48e7317ece6d-__obf_ecb441a4ce96b36a, __obf_587b9fceec62883e)
	}
}

// DrawHorizLine draws a horizontal line
func (__obf_f34639443602c065 *Palette) DrawHorizLine(__obf_c65a33fa97e57197, __obf_99724eb1801a48a0, __obf_431a48e7317ece6d int, __obf_587b9fceec62883e color.RGBA) {
	for __obf_80fae96c5947b2de := __obf_c65a33fa97e57197; __obf_80fae96c5947b2de <= __obf_99724eb1801a48a0; __obf_80fae96c5947b2de++ {
		__obf_f34639443602c065.Set(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_587b9fceec62883e)
	}
}

// Distort distorts the canvas
func (__obf_f34639443602c065 *Palette) Distort(__obf_a8f9c4009c1ab0e3 float64, __obf_e6491469052abb71 float64) {
	__obf_ac67bef5710e786a := __obf_f34639443602c065.Bounds().Max.X
	__obf_e455662263d832db := __obf_f34639443602c065.Bounds().Max.Y
	__obf_cac96793afab52b6 := NewPalette(image.Rect(0, 0, __obf_ac67bef5710e786a, __obf_e455662263d832db), __obf_f34639443602c065.Palette)
	__obf_5a4d5259564716ed := 2.0 * math.Pi / __obf_e6491469052abb71
	for __obf_80fae96c5947b2de := range __obf_ac67bef5710e786a {
		for __obf_431a48e7317ece6d := 0; __obf_431a48e7317ece6d < __obf_e455662263d832db; __obf_431a48e7317ece6d++ {
			__obf_ecb441a4ce96b36a := __obf_a8f9c4009c1ab0e3 * math.Sin(float64(__obf_431a48e7317ece6d)*__obf_5a4d5259564716ed)
			__obf_1a7ab87e567bf319 := __obf_a8f9c4009c1ab0e3 * math.Cos(float64(__obf_80fae96c5947b2de)*__obf_5a4d5259564716ed)
			__obf_cac96793afab52b6.SetColorIndex(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_f34639443602c065.ColorIndexAt(__obf_80fae96c5947b2de+int(__obf_ecb441a4ce96b36a), __obf_431a48e7317ece6d+int(__obf_1a7ab87e567bf319)))
		}
	}

	__obf_e61ea68b462e25a1 := __obf_cac96793afab52b6.Bounds().Max.X
	__obf_5cc37636829791d6 := __obf_cac96793afab52b6.Bounds().Max.Y
	for __obf_80fae96c5947b2de := range __obf_e61ea68b462e25a1 {
		for __obf_431a48e7317ece6d := 0; __obf_431a48e7317ece6d < __obf_5cc37636829791d6; __obf_431a48e7317ece6d++ {
			__obf_f34639443602c065.SetColorIndex(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_cac96793afab52b6.ColorIndexAt(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d))
		}
	}

	__obf_cac96793afab52b6.Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_f34639443602c065 *Palette) DrawBeeline(__obf_3e45008d4ba3c73b image.Point, __obf_0e4ae2ff6523e1c9 image.Point, __obf_f0abb089426179e6 color.RGBA) {
	__obf_5a4d5259564716ed := math.Abs(float64(__obf_3e45008d4ba3c73b.X - __obf_0e4ae2ff6523e1c9.X))
	__obf_1a34e4081dd5892b := math.Abs(float64(__obf_0e4ae2ff6523e1c9.Y - __obf_3e45008d4ba3c73b.Y))
	__obf_26a84b4d1bfe95d3, __obf_9caf26d4d5e92e0a := 1, 1
	if __obf_3e45008d4ba3c73b.X >= __obf_0e4ae2ff6523e1c9.X {
		__obf_26a84b4d1bfe95d3 = -1
	}
	if __obf_3e45008d4ba3c73b.Y >= __obf_0e4ae2ff6523e1c9.Y {
		__obf_9caf26d4d5e92e0a = -1
	}
	__obf_369d77920507ae91 := __obf_5a4d5259564716ed - __obf_1a34e4081dd5892b
	for {
		__obf_f34639443602c065.Set(__obf_3e45008d4ba3c73b.X, __obf_3e45008d4ba3c73b.Y, __obf_f0abb089426179e6)
		__obf_f34639443602c065.Set(__obf_3e45008d4ba3c73b.X+1, __obf_3e45008d4ba3c73b.Y, __obf_f0abb089426179e6)
		__obf_f34639443602c065.Set(__obf_3e45008d4ba3c73b.X-1, __obf_3e45008d4ba3c73b.Y, __obf_f0abb089426179e6)
		__obf_f34639443602c065.Set(__obf_3e45008d4ba3c73b.X+2, __obf_3e45008d4ba3c73b.Y, __obf_f0abb089426179e6)
		__obf_f34639443602c065.Set(__obf_3e45008d4ba3c73b.X-2, __obf_3e45008d4ba3c73b.Y, __obf_f0abb089426179e6)
		if __obf_3e45008d4ba3c73b.X == __obf_0e4ae2ff6523e1c9.X && __obf_3e45008d4ba3c73b.Y == __obf_0e4ae2ff6523e1c9.Y {
			return
		}
		__obf_3896326ad4836117 := __obf_369d77920507ae91 * 2
		if __obf_3896326ad4836117 > -__obf_1a34e4081dd5892b {
			__obf_369d77920507ae91 -= __obf_1a34e4081dd5892b
			__obf_3e45008d4ba3c73b.X += __obf_26a84b4d1bfe95d3
		}
		if __obf_3896326ad4836117 < __obf_5a4d5259564716ed {
			__obf_369d77920507ae91 += __obf_5a4d5259564716ed
			__obf_3e45008d4ba3c73b.Y += __obf_9caf26d4d5e92e0a
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_f34639443602c065 *Palette) AngleSwapPoint(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d, __obf_698deef70980b293, __obf_31d599a870027db7 float64) (__obf_3e6059cda13ac423, __obf_373f98fdead6a9af float64) {
	__obf_80fae96c5947b2de -= __obf_698deef70980b293
	__obf_431a48e7317ece6d = __obf_698deef70980b293 - __obf_431a48e7317ece6d
	__obf_d18496b73aeeba2b := math.Sin(__obf_31d599a870027db7 * (math.Pi / 180))
	__obf_6120b9fd5b92019f := math.Cos(__obf_31d599a870027db7 * (math.Pi / 180))
	__obf_3e6059cda13ac423 = __obf_80fae96c5947b2de*__obf_6120b9fd5b92019f + __obf_431a48e7317ece6d*__obf_d18496b73aeeba2b
	__obf_373f98fdead6a9af = -__obf_80fae96c5947b2de*__obf_d18496b73aeeba2b + __obf_431a48e7317ece6d*__obf_6120b9fd5b92019f
	__obf_3e6059cda13ac423 += __obf_698deef70980b293
	__obf_373f98fdead6a9af = __obf_698deef70980b293 - __obf_373f98fdead6a9af
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_f34639443602c065 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_e61ea68b462e25a1 := __obf_f34639443602c065.Bounds().Max.X
	__obf_5cc37636829791d6 := __obf_f34639443602c065.Bounds().Max.Y
	__obf_00491d31e8fc60c8 := __obf_e61ea68b462e25a1
	__obf_7c3283c789f60bf3 := 0
	__obf_4995721af19641b6 := __obf_5cc37636829791d6
	__obf_c0a3fea089ed7c65 := 0
	for __obf_80fae96c5947b2de := range __obf_e61ea68b462e25a1 {
		for __obf_431a48e7317ece6d := range __obf_5cc37636829791d6 {
			__obf_ddc50ca41a79bca6 := __obf_f34639443602c065.At(__obf_80fae96c5947b2de, __obf_431a48e7317ece6d)
			if _, _, _, __obf_bd896093c58f47ac := __obf_ddc50ca41a79bca6.RGBA(); __obf_bd896093c58f47ac > 0 {
				if __obf_80fae96c5947b2de < __obf_00491d31e8fc60c8 {
					__obf_00491d31e8fc60c8 = __obf_80fae96c5947b2de
				}
				if __obf_80fae96c5947b2de > __obf_7c3283c789f60bf3 {
					__obf_7c3283c789f60bf3 = __obf_80fae96c5947b2de
				}

				if __obf_431a48e7317ece6d < __obf_4995721af19641b6 {
					__obf_4995721af19641b6 = __obf_431a48e7317ece6d
				}
				if __obf_431a48e7317ece6d > __obf_c0a3fea089ed7c65 {
					__obf_c0a3fea089ed7c65 = __obf_431a48e7317ece6d
				}
			}
		}
	}

	__obf_00491d31e8fc60c8 = int(max(0, float64(__obf_00491d31e8fc60c8-2)))
	__obf_7c3283c789f60bf3 = int(min(float64(__obf_e61ea68b462e25a1), float64(__obf_7c3283c789f60bf3+2)))
	__obf_4995721af19641b6 = int(max(0, float64(__obf_4995721af19641b6-2)))
	__obf_c0a3fea089ed7c65 = int(min(float64(__obf_5cc37636829791d6), float64(__obf_c0a3fea089ed7c65+2)))

	return &AreaRect{
		__obf_00491d31e8fc60c8,
		__obf_7c3283c789f60bf3,
		__obf_4995721af19641b6,
		__obf_c0a3fea089ed7c65,
	}
}
