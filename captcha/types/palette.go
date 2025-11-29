package __obf_bda21a78cb74016a

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_49fbaf091ac38b90 image.Rectangle, __obf_f9811d5ad180ca6c color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_49fbaf091ac38b90,

			// palette struct for palette canvas
			__obf_f9811d5ad180ca6c),
	}
}

type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_f9811d5ad180ca6c *Palette) Rotate(__obf_1ed21c16a991a4c4 int) {
	if __obf_1ed21c16a991a4c4 == 0 {
		return
	}
	__obf_2318571e51c1cdc0 := __obf_f9811d5ad180ca6c
	__obf_846dcd5dff3fdc4a := __obf_2318571e51c1cdc0.Bounds().Max.X
	__obf_a9ee4ebbff00c852 := __obf_2318571e51c1cdc0.Bounds().Max.Y
	__obf_49fbaf091ac38b90 := __obf_846dcd5dff3fdc4a / 2
	__obf_6b14c6b6da40b10d := image.NewPaletted(image.Rect(0, 0, __obf_846dcd5dff3fdc4a, __obf_a9ee4ebbff00c852), __obf_2318571e51c1cdc0.Palette)
	for __obf_e22e7e43df48995e := 0; __obf_e22e7e43df48995e <= __obf_6b14c6b6da40b10d.Bounds().Max.X; __obf_e22e7e43df48995e++ {
		for __obf_a4dd1bd05990217f := 0; __obf_a4dd1bd05990217f <= __obf_6b14c6b6da40b10d.Bounds().Max.Y; __obf_a4dd1bd05990217f++ {
			__obf_110de2b4951ac74b, __obf_bdd6993b2dbacb26 := __obf_f9811d5ad180ca6c.AngleSwapPoint(float64(__obf_e22e7e43df48995e), float64(__obf_a4dd1bd05990217f), float64(__obf_49fbaf091ac38b90), float64(__obf_1ed21c16a991a4c4))
			__obf_6b14c6b6da40b10d.
				SetColorIndex(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_2318571e51c1cdc0.ColorIndexAt(int(__obf_110de2b4951ac74b), int(__obf_bdd6993b2dbacb26)))
		}
	}
	__obf_86fb1adf8186600b := __obf_6b14c6b6da40b10d.Bounds().Max.X
	__obf_151c7628f33db4f3 := __obf_6b14c6b6da40b10d.Bounds().Max.Y
	for __obf_e22e7e43df48995e := range __obf_86fb1adf8186600b {
		for __obf_a4dd1bd05990217f := range __obf_151c7628f33db4f3 {
			__obf_f9811d5ad180ca6c.
				SetColorIndex(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f,

					// DrawCircle draws a circle
					__obf_6b14c6b6da40b10d.ColorIndexAt(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f))
		}
	}
}

func (__obf_f9811d5ad180ca6c *Palette) DrawCircle(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_f0b1976bf7767947 int, __obf_6ac27ae0c7f24e19 color.RGBA) {
	__obf_e410250f1f5b821d := 1 - __obf_f0b1976bf7767947
	__obf_1202ba30fd9e6380 := 1
	__obf_e9091bdad9b1a062 := -2 * __obf_f0b1976bf7767947
	__obf_5dc723194163a108 := 0
	__obf_d9fa4a49c471dcfa := __obf_f0b1976bf7767947
	__obf_f9811d5ad180ca6c.
		Set(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f+__obf_f0b1976bf7767947, __obf_6ac27ae0c7f24e19)
	__obf_f9811d5ad180ca6c.
		Set(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f-__obf_f0b1976bf7767947, __obf_6ac27ae0c7f24e19)
	__obf_f9811d5ad180ca6c.
		DrawHorizLine(__obf_e22e7e43df48995e-__obf_f0b1976bf7767947, __obf_e22e7e43df48995e+__obf_f0b1976bf7767947, __obf_a4dd1bd05990217f, __obf_6ac27ae0c7f24e19)

	for __obf_5dc723194163a108 < __obf_d9fa4a49c471dcfa {
		if __obf_e410250f1f5b821d >= 0 {
			__obf_d9fa4a49c471dcfa--
			__obf_e9091bdad9b1a062 += 2
			__obf_e410250f1f5b821d += __obf_e9091bdad9b1a062
		}
		__obf_5dc723194163a108++
		__obf_1202ba30fd9e6380 += 2
		__obf_e410250f1f5b821d += __obf_1202ba30fd9e6380
		__obf_f9811d5ad180ca6c.
			DrawHorizLine(__obf_e22e7e43df48995e-__obf_5dc723194163a108, __obf_e22e7e43df48995e+__obf_5dc723194163a108, __obf_a4dd1bd05990217f+__obf_d9fa4a49c471dcfa, __obf_6ac27ae0c7f24e19)
		__obf_f9811d5ad180ca6c.
			DrawHorizLine(__obf_e22e7e43df48995e-__obf_5dc723194163a108, __obf_e22e7e43df48995e+__obf_5dc723194163a108, __obf_a4dd1bd05990217f-__obf_d9fa4a49c471dcfa, __obf_6ac27ae0c7f24e19)
		__obf_f9811d5ad180ca6c.
			DrawHorizLine(__obf_e22e7e43df48995e-__obf_d9fa4a49c471dcfa, __obf_e22e7e43df48995e+__obf_d9fa4a49c471dcfa, __obf_a4dd1bd05990217f+__obf_5dc723194163a108, __obf_6ac27ae0c7f24e19)
		__obf_f9811d5ad180ca6c.
			DrawHorizLine(__obf_e22e7e43df48995e-__obf_d9fa4a49c471dcfa, __obf_e22e7e43df48995e+__obf_d9fa4a49c471dcfa,

				// DrawHorizLine draws a horizontal line
				__obf_a4dd1bd05990217f-__obf_5dc723194163a108, __obf_6ac27ae0c7f24e19)
	}
}

func (__obf_f9811d5ad180ca6c *Palette) DrawHorizLine(__obf_8104e7d203ef43c3, __obf_37ba45c9003334cb, __obf_a4dd1bd05990217f int, __obf_6ac27ae0c7f24e19 color.RGBA) {
	for __obf_e22e7e43df48995e := __obf_8104e7d203ef43c3; __obf_e22e7e43df48995e <= __obf_37ba45c9003334cb; __obf_e22e7e43df48995e++ {
		__obf_f9811d5ad180ca6c.
			Set(__obf_e22e7e43df48995e,

				// Distort distorts the canvas
				__obf_a4dd1bd05990217f, __obf_6ac27ae0c7f24e19)
	}
}

func (__obf_f9811d5ad180ca6c *Palette) Distort(__obf_aecd379e5c5c7398 float64, __obf_ce5d928ef179cbc8 float64) {
	__obf_6f8e99209c36dd76 := __obf_f9811d5ad180ca6c.Bounds().Max.X
	__obf_b02fcc956b08e1ae := __obf_f9811d5ad180ca6c.Bounds().Max.Y
	__obf_ce09ee653fddfbd3 := NewPalette(image.Rect(0, 0, __obf_6f8e99209c36dd76, __obf_b02fcc956b08e1ae), __obf_f9811d5ad180ca6c.Palette)
	__obf_2d0838a79d72d9c2 := 2.0 * math.Pi / __obf_ce5d928ef179cbc8
	for __obf_e22e7e43df48995e := range __obf_6f8e99209c36dd76 {
		for __obf_a4dd1bd05990217f := 0; __obf_a4dd1bd05990217f < __obf_b02fcc956b08e1ae; __obf_a4dd1bd05990217f++ {
			__obf_5dc723194163a108 := __obf_aecd379e5c5c7398 * math.Sin(float64(__obf_a4dd1bd05990217f)*__obf_2d0838a79d72d9c2)
			__obf_d9fa4a49c471dcfa := __obf_aecd379e5c5c7398 * math.Cos(float64(__obf_e22e7e43df48995e)*__obf_2d0838a79d72d9c2)
			__obf_ce09ee653fddfbd3.
				SetColorIndex(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_f9811d5ad180ca6c.ColorIndexAt(__obf_e22e7e43df48995e+int(__obf_5dc723194163a108), __obf_a4dd1bd05990217f+int(__obf_d9fa4a49c471dcfa)))
		}
	}
	__obf_86fb1adf8186600b := __obf_ce09ee653fddfbd3.Bounds().Max.X
	__obf_151c7628f33db4f3 := __obf_ce09ee653fddfbd3.Bounds().Max.Y
	for __obf_e22e7e43df48995e := range __obf_86fb1adf8186600b {
		for __obf_a4dd1bd05990217f := 0; __obf_a4dd1bd05990217f < __obf_151c7628f33db4f3; __obf_a4dd1bd05990217f++ {
			__obf_f9811d5ad180ca6c.
				SetColorIndex(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_ce09ee653fddfbd3.ColorIndexAt(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f))
		}
	}
	__obf_ce09ee653fddfbd3.
		Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_f9811d5ad180ca6c *Palette) DrawBeeline(__obf_69d434f6b7593a82 image.Point, __obf_8fba081dfc158f10 image.Point, __obf_d93b5e0b638aa5f6 color.RGBA) {
	__obf_2d0838a79d72d9c2 := math.Abs(float64(__obf_69d434f6b7593a82.X - __obf_8fba081dfc158f10.X))
	__obf_68c9cbdb41f80cb3 := math.Abs(float64(__obf_8fba081dfc158f10.Y - __obf_69d434f6b7593a82.Y))
	__obf_dee267c29a69ea79, __obf_da9428382656c25e := 1, 1
	if __obf_69d434f6b7593a82.X >= __obf_8fba081dfc158f10.X {
		__obf_dee267c29a69ea79 = -1
	}
	if __obf_69d434f6b7593a82.Y >= __obf_8fba081dfc158f10.Y {
		__obf_da9428382656c25e = -1
	}
	__obf_eb4c912f039240f3 := __obf_2d0838a79d72d9c2 - __obf_68c9cbdb41f80cb3
	for {
		__obf_f9811d5ad180ca6c.
			Set(__obf_69d434f6b7593a82.X, __obf_69d434f6b7593a82.Y, __obf_d93b5e0b638aa5f6)
		__obf_f9811d5ad180ca6c.
			Set(__obf_69d434f6b7593a82.X+1, __obf_69d434f6b7593a82.Y, __obf_d93b5e0b638aa5f6)
		__obf_f9811d5ad180ca6c.
			Set(__obf_69d434f6b7593a82.X-1, __obf_69d434f6b7593a82.Y, __obf_d93b5e0b638aa5f6)
		__obf_f9811d5ad180ca6c.
			Set(__obf_69d434f6b7593a82.X+2, __obf_69d434f6b7593a82.Y, __obf_d93b5e0b638aa5f6)
		__obf_f9811d5ad180ca6c.
			Set(__obf_69d434f6b7593a82.X-2, __obf_69d434f6b7593a82.Y, __obf_d93b5e0b638aa5f6)
		if __obf_69d434f6b7593a82.X == __obf_8fba081dfc158f10.X && __obf_69d434f6b7593a82.Y == __obf_8fba081dfc158f10.Y {
			return
		}
		__obf_b73a17096071a6f0 := __obf_eb4c912f039240f3 * 2
		if __obf_b73a17096071a6f0 > -__obf_68c9cbdb41f80cb3 {
			__obf_eb4c912f039240f3 -= __obf_68c9cbdb41f80cb3
			__obf_69d434f6b7593a82.
				X += __obf_dee267c29a69ea79
		}
		if __obf_b73a17096071a6f0 < __obf_2d0838a79d72d9c2 {
			__obf_eb4c912f039240f3 += __obf_2d0838a79d72d9c2
			__obf_69d434f6b7593a82.
				Y += __obf_da9428382656c25e
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_f9811d5ad180ca6c *Palette) AngleSwapPoint(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f, __obf_49fbaf091ac38b90, __obf_1ed21c16a991a4c4 float64) (__obf_dd1b30e5888eb2ca, __obf_1c14fcfae6b3f36e float64) {
	__obf_e22e7e43df48995e -= __obf_49fbaf091ac38b90
	__obf_a4dd1bd05990217f = __obf_49fbaf091ac38b90 - __obf_a4dd1bd05990217f
	__obf_100a8b00d63fd25b := math.Sin(__obf_1ed21c16a991a4c4 * (math.Pi / 180))
	__obf_2c437cf77dd33361 := math.Cos(__obf_1ed21c16a991a4c4 * (math.Pi / 180))
	__obf_dd1b30e5888eb2ca = __obf_e22e7e43df48995e*__obf_2c437cf77dd33361 + __obf_a4dd1bd05990217f*__obf_100a8b00d63fd25b
	__obf_1c14fcfae6b3f36e = -__obf_e22e7e43df48995e*__obf_100a8b00d63fd25b + __obf_a4dd1bd05990217f*__obf_2c437cf77dd33361
	__obf_dd1b30e5888eb2ca += __obf_49fbaf091ac38b90
	__obf_1c14fcfae6b3f36e = __obf_49fbaf091ac38b90 - __obf_1c14fcfae6b3f36e
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_f9811d5ad180ca6c *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_86fb1adf8186600b := __obf_f9811d5ad180ca6c.Bounds().Max.X
	__obf_151c7628f33db4f3 := __obf_f9811d5ad180ca6c.Bounds().Max.Y
	__obf_cd41514d006d0679 := __obf_86fb1adf8186600b
	__obf_a8629efeda1a9547 := 0
	__obf_2d89945f425dcbd1 := __obf_151c7628f33db4f3
	__obf_3d1794aad27ceed1 := 0
	for __obf_e22e7e43df48995e := range __obf_86fb1adf8186600b {
		for __obf_a4dd1bd05990217f := range __obf_151c7628f33db4f3 {
			__obf_e7796a53fdcf1f25 := __obf_f9811d5ad180ca6c.At(__obf_e22e7e43df48995e, __obf_a4dd1bd05990217f)
			if _, _, _, __obf_6ed92d728f70797c := __obf_e7796a53fdcf1f25.RGBA(); __obf_6ed92d728f70797c > 0 {
				if __obf_e22e7e43df48995e < __obf_cd41514d006d0679 {
					__obf_cd41514d006d0679 = __obf_e22e7e43df48995e
				}
				if __obf_e22e7e43df48995e > __obf_a8629efeda1a9547 {
					__obf_a8629efeda1a9547 = __obf_e22e7e43df48995e
				}

				if __obf_a4dd1bd05990217f < __obf_2d89945f425dcbd1 {
					__obf_2d89945f425dcbd1 = __obf_a4dd1bd05990217f
				}
				if __obf_a4dd1bd05990217f > __obf_3d1794aad27ceed1 {
					__obf_3d1794aad27ceed1 = __obf_a4dd1bd05990217f
				}
			}
		}
	}
	__obf_cd41514d006d0679 = int(max(0, float64(__obf_cd41514d006d0679-2)))
	__obf_a8629efeda1a9547 = int(min(float64(__obf_86fb1adf8186600b), float64(__obf_a8629efeda1a9547+2)))
	__obf_2d89945f425dcbd1 = int(max(0, float64(__obf_2d89945f425dcbd1-2)))
	__obf_3d1794aad27ceed1 = int(min(float64(__obf_151c7628f33db4f3), float64(__obf_3d1794aad27ceed1+2)))

	return &AreaRect{__obf_cd41514d006d0679, __obf_a8629efeda1a9547, __obf_2d89945f425dcbd1, __obf_3d1794aad27ceed1}
}
