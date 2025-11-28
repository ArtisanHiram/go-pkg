package __obf_0e73be9e4c3ea3fd

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_a6fdf04202ab9ad1 image.Rectangle, __obf_c0d27cca35ba9554 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_a6fdf04202ab9ad1, __obf_c0d27cca35ba9554),
	}
}

// palette struct for palette canvas
type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_c0d27cca35ba9554 *Palette) Rotate(__obf_b51ed5d77ee8eb95 int) {
	if __obf_b51ed5d77ee8eb95 == 0 {
		return
	}

	__obf_3be7f6087dd3785d := __obf_c0d27cca35ba9554
	__obf_0107a6cb34e1b800 := __obf_3be7f6087dd3785d.Bounds().Max.X
	__obf_1f3235dfd61ebd0b := __obf_3be7f6087dd3785d.Bounds().Max.Y
	__obf_a6fdf04202ab9ad1 := __obf_0107a6cb34e1b800 / 2
	__obf_94a08be74f10361b := image.NewPaletted(image.Rect(0, 0, __obf_0107a6cb34e1b800, __obf_1f3235dfd61ebd0b), __obf_3be7f6087dd3785d.Palette)
	for __obf_fe086773828b5535 := 0; __obf_fe086773828b5535 <= __obf_94a08be74f10361b.Bounds().Max.X; __obf_fe086773828b5535++ {
		for __obf_44ef04afc5c8241a := 0; __obf_44ef04afc5c8241a <= __obf_94a08be74f10361b.Bounds().Max.Y; __obf_44ef04afc5c8241a++ {
			__obf_f4f274a081c54b08, __obf_504a67153a2adc23 := __obf_c0d27cca35ba9554.AngleSwapPoint(float64(__obf_fe086773828b5535), float64(__obf_44ef04afc5c8241a), float64(__obf_a6fdf04202ab9ad1), float64(__obf_b51ed5d77ee8eb95))
			__obf_94a08be74f10361b.SetColorIndex(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_3be7f6087dd3785d.ColorIndexAt(int(__obf_f4f274a081c54b08), int(__obf_504a67153a2adc23)))
		}
	}

	__obf_08f63af009e2059f := __obf_94a08be74f10361b.Bounds().Max.X
	__obf_71d18dd207e86c6a := __obf_94a08be74f10361b.Bounds().Max.Y
	for __obf_fe086773828b5535 := range __obf_08f63af009e2059f {
		for __obf_44ef04afc5c8241a := range __obf_71d18dd207e86c6a {
			__obf_c0d27cca35ba9554.SetColorIndex(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_94a08be74f10361b.ColorIndexAt(__obf_fe086773828b5535, __obf_44ef04afc5c8241a))
		}
	}
}

// DrawCircle draws a circle
func (__obf_c0d27cca35ba9554 *Palette) DrawCircle(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_920a1f23584de193 int, __obf_a5de954533461107 color.RGBA) {
	__obf_d91011513f63ebf0 := 1 - __obf_920a1f23584de193
	__obf_1a8c49893d231a8e := 1
	__obf_7329470948b4e284 := -2 * __obf_920a1f23584de193
	__obf_c032edef9f7867e0 := 0
	__obf_89bdc2e8062574b0 := __obf_920a1f23584de193

	__obf_c0d27cca35ba9554.Set(__obf_fe086773828b5535, __obf_44ef04afc5c8241a+__obf_920a1f23584de193, __obf_a5de954533461107)
	__obf_c0d27cca35ba9554.Set(__obf_fe086773828b5535, __obf_44ef04afc5c8241a-__obf_920a1f23584de193, __obf_a5de954533461107)
	__obf_c0d27cca35ba9554.DrawHorizLine(__obf_fe086773828b5535-__obf_920a1f23584de193, __obf_fe086773828b5535+__obf_920a1f23584de193, __obf_44ef04afc5c8241a, __obf_a5de954533461107)

	for __obf_c032edef9f7867e0 < __obf_89bdc2e8062574b0 {
		if __obf_d91011513f63ebf0 >= 0 {
			__obf_89bdc2e8062574b0--
			__obf_7329470948b4e284 += 2
			__obf_d91011513f63ebf0 += __obf_7329470948b4e284
		}
		__obf_c032edef9f7867e0++
		__obf_1a8c49893d231a8e += 2
		__obf_d91011513f63ebf0 += __obf_1a8c49893d231a8e
		__obf_c0d27cca35ba9554.DrawHorizLine(__obf_fe086773828b5535-__obf_c032edef9f7867e0, __obf_fe086773828b5535+__obf_c032edef9f7867e0, __obf_44ef04afc5c8241a+__obf_89bdc2e8062574b0, __obf_a5de954533461107)
		__obf_c0d27cca35ba9554.DrawHorizLine(__obf_fe086773828b5535-__obf_c032edef9f7867e0, __obf_fe086773828b5535+__obf_c032edef9f7867e0, __obf_44ef04afc5c8241a-__obf_89bdc2e8062574b0, __obf_a5de954533461107)
		__obf_c0d27cca35ba9554.DrawHorizLine(__obf_fe086773828b5535-__obf_89bdc2e8062574b0, __obf_fe086773828b5535+__obf_89bdc2e8062574b0, __obf_44ef04afc5c8241a+__obf_c032edef9f7867e0, __obf_a5de954533461107)
		__obf_c0d27cca35ba9554.DrawHorizLine(__obf_fe086773828b5535-__obf_89bdc2e8062574b0, __obf_fe086773828b5535+__obf_89bdc2e8062574b0, __obf_44ef04afc5c8241a-__obf_c032edef9f7867e0, __obf_a5de954533461107)
	}
}

// DrawHorizLine draws a horizontal line
func (__obf_c0d27cca35ba9554 *Palette) DrawHorizLine(__obf_a96898f0ed50ce1c, __obf_e7bf44b03f0c66de, __obf_44ef04afc5c8241a int, __obf_a5de954533461107 color.RGBA) {
	for __obf_fe086773828b5535 := __obf_a96898f0ed50ce1c; __obf_fe086773828b5535 <= __obf_e7bf44b03f0c66de; __obf_fe086773828b5535++ {
		__obf_c0d27cca35ba9554.Set(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_a5de954533461107)
	}
}

// Distort distorts the canvas
func (__obf_c0d27cca35ba9554 *Palette) Distort(__obf_ac9e8844ac797c1c float64, __obf_87586e4e4e684862 float64) {
	__obf_6e1fff01dcc62820 := __obf_c0d27cca35ba9554.Bounds().Max.X
	__obf_462a1d8e0f4b0353 := __obf_c0d27cca35ba9554.Bounds().Max.Y
	__obf_d05761f4b131e510 := NewPalette(image.Rect(0, 0, __obf_6e1fff01dcc62820, __obf_462a1d8e0f4b0353), __obf_c0d27cca35ba9554.Palette)
	__obf_6acf1e629d444cbc := 2.0 * math.Pi / __obf_87586e4e4e684862
	for __obf_fe086773828b5535 := range __obf_6e1fff01dcc62820 {
		for __obf_44ef04afc5c8241a := 0; __obf_44ef04afc5c8241a < __obf_462a1d8e0f4b0353; __obf_44ef04afc5c8241a++ {
			__obf_c032edef9f7867e0 := __obf_ac9e8844ac797c1c * math.Sin(float64(__obf_44ef04afc5c8241a)*__obf_6acf1e629d444cbc)
			__obf_89bdc2e8062574b0 := __obf_ac9e8844ac797c1c * math.Cos(float64(__obf_fe086773828b5535)*__obf_6acf1e629d444cbc)
			__obf_d05761f4b131e510.SetColorIndex(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_c0d27cca35ba9554.ColorIndexAt(__obf_fe086773828b5535+int(__obf_c032edef9f7867e0), __obf_44ef04afc5c8241a+int(__obf_89bdc2e8062574b0)))
		}
	}

	__obf_08f63af009e2059f := __obf_d05761f4b131e510.Bounds().Max.X
	__obf_71d18dd207e86c6a := __obf_d05761f4b131e510.Bounds().Max.Y
	for __obf_fe086773828b5535 := range __obf_08f63af009e2059f {
		for __obf_44ef04afc5c8241a := 0; __obf_44ef04afc5c8241a < __obf_71d18dd207e86c6a; __obf_44ef04afc5c8241a++ {
			__obf_c0d27cca35ba9554.SetColorIndex(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_d05761f4b131e510.ColorIndexAt(__obf_fe086773828b5535, __obf_44ef04afc5c8241a))
		}
	}

	__obf_d05761f4b131e510.Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_c0d27cca35ba9554 *Palette) DrawBeeline(__obf_ede46109bf6beec5 image.Point, __obf_3c51a056c8e30711 image.Point, __obf_0204f3b5c8924f76 color.RGBA) {
	__obf_6acf1e629d444cbc := math.Abs(float64(__obf_ede46109bf6beec5.X - __obf_3c51a056c8e30711.X))
	__obf_9694d55e6354e1b8 := math.Abs(float64(__obf_3c51a056c8e30711.Y - __obf_ede46109bf6beec5.Y))
	__obf_dec4fd248192db69, __obf_f914f80811c9c522 := 1, 1
	if __obf_ede46109bf6beec5.X >= __obf_3c51a056c8e30711.X {
		__obf_dec4fd248192db69 = -1
	}
	if __obf_ede46109bf6beec5.Y >= __obf_3c51a056c8e30711.Y {
		__obf_f914f80811c9c522 = -1
	}
	__obf_335fb0d6019c95ed := __obf_6acf1e629d444cbc - __obf_9694d55e6354e1b8
	for {
		__obf_c0d27cca35ba9554.Set(__obf_ede46109bf6beec5.X, __obf_ede46109bf6beec5.Y, __obf_0204f3b5c8924f76)
		__obf_c0d27cca35ba9554.Set(__obf_ede46109bf6beec5.X+1, __obf_ede46109bf6beec5.Y, __obf_0204f3b5c8924f76)
		__obf_c0d27cca35ba9554.Set(__obf_ede46109bf6beec5.X-1, __obf_ede46109bf6beec5.Y, __obf_0204f3b5c8924f76)
		__obf_c0d27cca35ba9554.Set(__obf_ede46109bf6beec5.X+2, __obf_ede46109bf6beec5.Y, __obf_0204f3b5c8924f76)
		__obf_c0d27cca35ba9554.Set(__obf_ede46109bf6beec5.X-2, __obf_ede46109bf6beec5.Y, __obf_0204f3b5c8924f76)
		if __obf_ede46109bf6beec5.X == __obf_3c51a056c8e30711.X && __obf_ede46109bf6beec5.Y == __obf_3c51a056c8e30711.Y {
			return
		}
		__obf_e8990a4814150863 := __obf_335fb0d6019c95ed * 2
		if __obf_e8990a4814150863 > -__obf_9694d55e6354e1b8 {
			__obf_335fb0d6019c95ed -= __obf_9694d55e6354e1b8
			__obf_ede46109bf6beec5.X += __obf_dec4fd248192db69
		}
		if __obf_e8990a4814150863 < __obf_6acf1e629d444cbc {
			__obf_335fb0d6019c95ed += __obf_6acf1e629d444cbc
			__obf_ede46109bf6beec5.Y += __obf_f914f80811c9c522
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_c0d27cca35ba9554 *Palette) AngleSwapPoint(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_a6fdf04202ab9ad1, __obf_b51ed5d77ee8eb95 float64) (__obf_4596e278855415ed, __obf_a6f9e41f931023cd float64) {
	__obf_fe086773828b5535 -= __obf_a6fdf04202ab9ad1
	__obf_44ef04afc5c8241a = __obf_a6fdf04202ab9ad1 - __obf_44ef04afc5c8241a
	__obf_25f35019dbad29f4 := math.Sin(__obf_b51ed5d77ee8eb95 * (math.Pi / 180))
	__obf_5bb4dd56de35a027 := math.Cos(__obf_b51ed5d77ee8eb95 * (math.Pi / 180))
	__obf_4596e278855415ed = __obf_fe086773828b5535*__obf_5bb4dd56de35a027 + __obf_44ef04afc5c8241a*__obf_25f35019dbad29f4
	__obf_a6f9e41f931023cd = -__obf_fe086773828b5535*__obf_25f35019dbad29f4 + __obf_44ef04afc5c8241a*__obf_5bb4dd56de35a027
	__obf_4596e278855415ed += __obf_a6fdf04202ab9ad1
	__obf_a6f9e41f931023cd = __obf_a6fdf04202ab9ad1 - __obf_a6f9e41f931023cd
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_c0d27cca35ba9554 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_08f63af009e2059f := __obf_c0d27cca35ba9554.Bounds().Max.X
	__obf_71d18dd207e86c6a := __obf_c0d27cca35ba9554.Bounds().Max.Y
	__obf_8a297e9a65a6b461 := __obf_08f63af009e2059f
	__obf_6dbed9412bb93f30 := 0
	__obf_a6d1e2b517fb3ca3 := __obf_71d18dd207e86c6a
	__obf_23d5edea6d8dce9a := 0
	for __obf_fe086773828b5535 := range __obf_08f63af009e2059f {
		for __obf_44ef04afc5c8241a := range __obf_71d18dd207e86c6a {
			__obf_f8353e66f8032833 := __obf_c0d27cca35ba9554.At(__obf_fe086773828b5535, __obf_44ef04afc5c8241a)
			if _, _, _, __obf_51d4fa10e4721296 := __obf_f8353e66f8032833.RGBA(); __obf_51d4fa10e4721296 > 0 {
				if __obf_fe086773828b5535 < __obf_8a297e9a65a6b461 {
					__obf_8a297e9a65a6b461 = __obf_fe086773828b5535
				}
				if __obf_fe086773828b5535 > __obf_6dbed9412bb93f30 {
					__obf_6dbed9412bb93f30 = __obf_fe086773828b5535
				}

				if __obf_44ef04afc5c8241a < __obf_a6d1e2b517fb3ca3 {
					__obf_a6d1e2b517fb3ca3 = __obf_44ef04afc5c8241a
				}
				if __obf_44ef04afc5c8241a > __obf_23d5edea6d8dce9a {
					__obf_23d5edea6d8dce9a = __obf_44ef04afc5c8241a
				}
			}
		}
	}

	__obf_8a297e9a65a6b461 = int(max(0, float64(__obf_8a297e9a65a6b461-2)))
	__obf_6dbed9412bb93f30 = int(min(float64(__obf_08f63af009e2059f), float64(__obf_6dbed9412bb93f30+2)))
	__obf_a6d1e2b517fb3ca3 = int(max(0, float64(__obf_a6d1e2b517fb3ca3-2)))
	__obf_23d5edea6d8dce9a = int(min(float64(__obf_71d18dd207e86c6a), float64(__obf_23d5edea6d8dce9a+2)))

	return &AreaRect{
		__obf_8a297e9a65a6b461,
		__obf_6dbed9412bb93f30,
		__obf_a6d1e2b517fb3ca3,
		__obf_23d5edea6d8dce9a,
	}
}
