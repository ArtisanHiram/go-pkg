package __obf_0e73be9e4c3ea3fd

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
	"golang.org/x/image/math/f64"
)

// NewNRGBA creates an NRGBA canvas
func NewNRGBA(__obf_a6fdf04202ab9ad1 image.Rectangle, __obf_64e9fc49b08a5410 bool) *NRGBA {
	__obf_260fd86be4634262 := image.NewNRGBA(__obf_a6fdf04202ab9ad1)
	for __obf_44ef04afc5c8241a := 0; __obf_44ef04afc5c8241a < __obf_a6fdf04202ab9ad1.Max.Y; __obf_44ef04afc5c8241a++ {
		for __obf_fe086773828b5535 := 0; __obf_fe086773828b5535 < __obf_a6fdf04202ab9ad1.Max.X; __obf_fe086773828b5535++ {
			if __obf_64e9fc49b08a5410 {
				__obf_260fd86be4634262.Set(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, color.Alpha{A: uint8(0)})
			} else {
				__obf_260fd86be4634262.Set(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, color.RGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}
	}

	return &NRGBA{
		__obf_260fd86be4634262,
	}
}

// nRGBA struct for NRGBA canvas
type NRGBA struct {
	*image.NRGBA
}

// DrawImage draws an image on the canvas
func (__obf_24d6334a9384367b *NRGBA) DrawImage(__obf_d1b36837c286905e Palette, __obf_4875dbc244b598fc *PositionRect, __obf_30c144e9b5707611 *AreaRect) {
	__obf_08f63af009e2059f := __obf_d1b36837c286905e.Bounds().Max.X
	__obf_71d18dd207e86c6a := __obf_d1b36837c286905e.Bounds().Max.Y

	__obf_d7efc27e652f6243 := __obf_4875dbc244b598fc.X
	__obf_a1ccb99605ac3bbf := __obf_4875dbc244b598fc.Y
	__obf_1a657db69177b979 := __obf_4875dbc244b598fc.Height

	__obf_917f47eaf281bb00 := __obf_30c144e9b5707611.MinX
	__obf_f0e4826c210ff934 := __obf_30c144e9b5707611.MinY
	__obf_0664ce4553a3fa24 := __obf_30c144e9b5707611.MaxX
	__obf_e9590880df2b216f := __obf_30c144e9b5707611.MaxY

	for __obf_fe086773828b5535 := range __obf_08f63af009e2059f {
		for __obf_44ef04afc5c8241a := range __obf_71d18dd207e86c6a {
			__obf_f8353e66f8032833 := __obf_d1b36837c286905e.At(__obf_fe086773828b5535, __obf_44ef04afc5c8241a)
			if _, _, _, __obf_51d4fa10e4721296 := __obf_f8353e66f8032833.RGBA(); __obf_51d4fa10e4721296 > 0 {
				if __obf_fe086773828b5535 >= __obf_917f47eaf281bb00 && __obf_fe086773828b5535 <= __obf_0664ce4553a3fa24 && __obf_44ef04afc5c8241a >= __obf_f0e4826c210ff934 && __obf_44ef04afc5c8241a <= __obf_e9590880df2b216f {
					__obf_24d6334a9384367b.Set(__obf_d7efc27e652f6243+(__obf_fe086773828b5535-__obf_917f47eaf281bb00), __obf_a1ccb99605ac3bbf-__obf_1a657db69177b979+(__obf_44ef04afc5c8241a-__obf_f0e4826c210ff934), __obf_d1b36837c286905e.At(__obf_fe086773828b5535, __obf_44ef04afc5c8241a))
				}
			}
		}
	}
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_24d6334a9384367b *NRGBA) CalcMarginBlankArea() *AreaRect {
	__obf_08f63af009e2059f := __obf_24d6334a9384367b.Bounds().Max.X
	__obf_71d18dd207e86c6a := __obf_24d6334a9384367b.Bounds().Max.Y
	__obf_8a297e9a65a6b461 := __obf_08f63af009e2059f
	__obf_6dbed9412bb93f30 := 0
	__obf_a6d1e2b517fb3ca3 := __obf_71d18dd207e86c6a
	__obf_23d5edea6d8dce9a := 0
	for __obf_fe086773828b5535 := range __obf_08f63af009e2059f {
		for __obf_44ef04afc5c8241a := range __obf_71d18dd207e86c6a {
			__obf_f8353e66f8032833 := __obf_24d6334a9384367b.At(__obf_fe086773828b5535, __obf_44ef04afc5c8241a)
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

// Rotate rotates the canvas by any angle
func (__obf_24d6334a9384367b *NRGBA) Rotate(__obf_51d4fa10e4721296 int, __obf_c3f67c1a57f9e07f bool) {
	if __obf_51d4fa10e4721296 == 0 {
		return
	}

	__obf_b51ed5d77ee8eb95 := float64(__obf_51d4fa10e4721296) * math.Pi / 180

	__obf_069fe07debfd09df := __obf_24d6334a9384367b.Bounds().Dx()
	__obf_7ac90e8abf5f3d4c := __obf_24d6334a9384367b.Bounds().Dy()
	__obf_6e1fff01dcc62820, __obf_462a1d8e0f4b0353 := RotatedSize(__obf_069fe07debfd09df, __obf_7ac90e8abf5f3d4c, float64(__obf_51d4fa10e4721296))
	__obf_d1b36837c286905e := image.NewNRGBA(image.Rect(0, 0, __obf_6e1fff01dcc62820, __obf_462a1d8e0f4b0353))

	__obf_5cfc55f28eddaecd := float64(__obf_6e1fff01dcc62820) / 2
	__obf_af34d97c504e1d5e := float64(__obf_462a1d8e0f4b0353) / 2

	__obf_3b835f3910392644 := Matrix{
		1, 0,
		0, 1,
		0, 0,
	}
	__obf_3b835f3910392644 = __obf_3b835f3910392644.Translate(__obf_5cfc55f28eddaecd, __obf_af34d97c504e1d5e)
	__obf_3b835f3910392644 = __obf_3b835f3910392644.Rotate(__obf_b51ed5d77ee8eb95)
	__obf_3b835f3910392644 = __obf_3b835f3910392644.Translate(-__obf_5cfc55f28eddaecd, -__obf_af34d97c504e1d5e)

	__obf_fe086773828b5535 := (__obf_6e1fff01dcc62820 - __obf_24d6334a9384367b.Bounds().Dx()) / 2
	__obf_44ef04afc5c8241a := (__obf_462a1d8e0f4b0353 - __obf_24d6334a9384367b.Bounds().Dy()) / 2
	__obf_dc6ef29159caa0ed, __obf_7a441b86be805694 := float64(__obf_fe086773828b5535), float64(__obf_44ef04afc5c8241a)

	__obf_7180bb73bdb90a3c := __obf_3b835f3910392644.Translate(__obf_dc6ef29159caa0ed, __obf_7a441b86be805694)
	__obf_d488d5b62a46f0d0 := f64.Aff3{__obf_7180bb73bdb90a3c.XX, __obf_7180bb73bdb90a3c.XY, __obf_7180bb73bdb90a3c.X0, __obf_7180bb73bdb90a3c.YX, __obf_7180bb73bdb90a3c.YY, __obf_7180bb73bdb90a3c.Y0}

	draw.BiLinear.Transform(__obf_d1b36837c286905e, __obf_d488d5b62a46f0d0, __obf_24d6334a9384367b, __obf_24d6334a9384367b.Bounds(), draw.Over, nil)
	__obf_24d6334a9384367b.NRGBA = __obf_d1b36837c286905e

	if __obf_c3f67c1a57f9e07f {
		__obf_343c8a89667c9131 := __obf_6e1fff01dcc62820 - __obf_069fe07debfd09df
		__obf_8b939fd5c6502686 := __obf_462a1d8e0f4b0353 - __obf_7ac90e8abf5f3d4c
		__obf_6acf1e629d444cbc := (__obf_343c8a89667c9131 / 2) + 1
		__obf_9694d55e6354e1b8 := (__obf_8b939fd5c6502686 / 2) + 1
		__obf_24d6334a9384367b.SubImage(image.Rect(__obf_6acf1e629d444cbc, __obf_9694d55e6354e1b8, __obf_069fe07debfd09df+__obf_6acf1e629d444cbc, __obf_7ac90e8abf5f3d4c+__obf_9694d55e6354e1b8))
	}
}

// CropCircle crops a circular area
func (__obf_24d6334a9384367b *NRGBA) CropCircle(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_920a1f23584de193 int) {
	__obf_5482ad0442b35da9 := __obf_24d6334a9384367b.Bounds()
	__obf_71bb5cb69fd330b9 := image.NewNRGBA(__obf_5482ad0442b35da9)
	for __obf_2c7bc2201a249217 := __obf_5482ad0442b35da9.Min.Y; __obf_2c7bc2201a249217 < __obf_5482ad0442b35da9.Max.Y; __obf_2c7bc2201a249217++ {
		for __obf_154f941903283413 := __obf_5482ad0442b35da9.Min.X; __obf_154f941903283413 < __obf_5482ad0442b35da9.Max.X; __obf_154f941903283413++ {
			__obf_cd8fa911a1304bdf := math.Hypot(float64(__obf_154f941903283413-__obf_fe086773828b5535), float64(__obf_2c7bc2201a249217-__obf_44ef04afc5c8241a))
			if __obf_cd8fa911a1304bdf <= float64(__obf_920a1f23584de193) {
				__obf_71bb5cb69fd330b9.Set(__obf_154f941903283413, __obf_2c7bc2201a249217, color.White)
			} else {
				__obf_71bb5cb69fd330b9.Set(__obf_154f941903283413, __obf_2c7bc2201a249217, color.Transparent)
			}
		}
	}

	draw.DrawMask(__obf_71bb5cb69fd330b9, __obf_71bb5cb69fd330b9.Bounds(), __obf_24d6334a9384367b, image.Point{X: 0, Y: 0}, __obf_71bb5cb69fd330b9, image.Point{}, draw.Over)
	__obf_24d6334a9384367b.NRGBA = __obf_71bb5cb69fd330b9
}

// CropScaleCircle scales and crops a circular area
func (__obf_24d6334a9384367b *NRGBA) CropScaleCircle(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_920a1f23584de193 int, __obf_8644d9caddfda387 int) {
	__obf_5482ad0442b35da9 := __obf_24d6334a9384367b.Bounds()
	__obf_71bb5cb69fd330b9 := image.NewNRGBA(__obf_5482ad0442b35da9)

	for __obf_2c7bc2201a249217 := __obf_5482ad0442b35da9.Min.Y; __obf_2c7bc2201a249217 < __obf_5482ad0442b35da9.Max.Y; __obf_2c7bc2201a249217++ {
		for __obf_154f941903283413 := __obf_5482ad0442b35da9.Min.X; __obf_154f941903283413 < __obf_5482ad0442b35da9.Max.X; __obf_154f941903283413++ {
			__obf_cd8fa911a1304bdf := math.Hypot(float64(__obf_154f941903283413-__obf_fe086773828b5535), float64(__obf_2c7bc2201a249217-__obf_44ef04afc5c8241a))
			if __obf_cd8fa911a1304bdf <= float64(__obf_920a1f23584de193) {
				__obf_71bb5cb69fd330b9.Set(__obf_154f941903283413, __obf_2c7bc2201a249217, color.White)
			} else {
				__obf_71bb5cb69fd330b9.Set(__obf_154f941903283413, __obf_2c7bc2201a249217, color.Transparent)
			}
		}
	}

	if __obf_8644d9caddfda387 > 0 {
		__obf_be5e0d14960a8f42 := __obf_8644d9caddfda387 * 2
		__obf_a748dcd67e4755b9 := image.NewNRGBA(image.Rect(0, 0, __obf_24d6334a9384367b.Bounds().Dx()-__obf_be5e0d14960a8f42, __obf_24d6334a9384367b.Bounds().Dy()-__obf_be5e0d14960a8f42))
		draw.BiLinear.Scale(__obf_a748dcd67e4755b9, __obf_a748dcd67e4755b9.Bounds(), __obf_71bb5cb69fd330b9, __obf_71bb5cb69fd330b9.Bounds(), draw.Over, nil)
		__obf_71bb5cb69fd330b9 = __obf_a748dcd67e4755b9
	}

	draw.DrawMask(__obf_71bb5cb69fd330b9, __obf_71bb5cb69fd330b9.Bounds(), __obf_24d6334a9384367b, image.Point{X: __obf_8644d9caddfda387, Y: __obf_8644d9caddfda387}, __obf_71bb5cb69fd330b9, image.Point{}, draw.Over)
	__obf_24d6334a9384367b.NRGBA = __obf_71bb5cb69fd330b9
}

// Scale scales the canvas
func (__obf_24d6334a9384367b *NRGBA) Scale(__obf_8644d9caddfda387 int, __obf_830a32ac8136502e, __obf_a03a45f487b7672e bool) {
	__obf_d1b36837c286905e := __obf_24d6334a9384367b.NRGBA
	if __obf_8644d9caddfda387 > 0 {
		__obf_be5e0d14960a8f42 := __obf_8644d9caddfda387 * 2
		__obf_f03181654d0054ef := __obf_24d6334a9384367b.Bounds().Dx() - __obf_be5e0d14960a8f42
		__obf_38df7d193682cf0d := __obf_24d6334a9384367b.Bounds().Dy() - __obf_be5e0d14960a8f42
		__obf_b9f005244f8771fc := image.NewNRGBA(image.Rect(0, 0, __obf_f03181654d0054ef, __obf_38df7d193682cf0d))

		if !__obf_830a32ac8136502e {
			draw.BiLinear.Scale(__obf_b9f005244f8771fc, __obf_b9f005244f8771fc.Bounds(), __obf_24d6334a9384367b, __obf_24d6334a9384367b.Bounds(), draw.Over, nil)
		} else {
			__obf_e5bbe575453b9248 := CalcResizedRect(__obf_24d6334a9384367b.Bounds(), __obf_f03181654d0054ef, __obf_38df7d193682cf0d, __obf_a03a45f487b7672e)
			draw.ApproxBiLinear.Scale(__obf_b9f005244f8771fc, __obf_e5bbe575453b9248.Bounds(), __obf_24d6334a9384367b, __obf_24d6334a9384367b.Bounds(), draw.Over, nil)
		}

		__obf_d1b36837c286905e = __obf_b9f005244f8771fc
	}

	__obf_24d6334a9384367b.NRGBA = __obf_d1b36837c286905e
}
