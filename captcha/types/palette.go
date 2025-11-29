package __obf_244ef50d49151021

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_662a4428339c8e5b image.Rectangle, __obf_2e70b49a4a643481 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_662a4428339c8e5b,

			// palette struct for palette canvas
			__obf_2e70b49a4a643481),
	}
}

type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_2e70b49a4a643481 *Palette) Rotate(__obf_2b5072b762a61186 int) {
	if __obf_2b5072b762a61186 == 0 {
		return
	}
	__obf_e20d08f7af3e7323 := __obf_2e70b49a4a643481
	__obf_c9319b7e133ffbf3 := __obf_e20d08f7af3e7323.Bounds().Max.X
	__obf_15734090882caa7f := __obf_e20d08f7af3e7323.Bounds().Max.Y
	__obf_662a4428339c8e5b := __obf_c9319b7e133ffbf3 / 2
	__obf_51b09dba8f91c6f5 := image.NewPaletted(image.Rect(0, 0, __obf_c9319b7e133ffbf3, __obf_15734090882caa7f), __obf_e20d08f7af3e7323.Palette)
	for __obf_e487b3618d36e102 := 0; __obf_e487b3618d36e102 <= __obf_51b09dba8f91c6f5.Bounds().Max.X; __obf_e487b3618d36e102++ {
		for __obf_8cbd87ef15403513 := 0; __obf_8cbd87ef15403513 <= __obf_51b09dba8f91c6f5.Bounds().Max.Y; __obf_8cbd87ef15403513++ {
			__obf_4e172b662bf3e6a8, __obf_82a1080185d9e2b7 := __obf_2e70b49a4a643481.AngleSwapPoint(float64(__obf_e487b3618d36e102), float64(__obf_8cbd87ef15403513), float64(__obf_662a4428339c8e5b), float64(__obf_2b5072b762a61186))
			__obf_51b09dba8f91c6f5.
				SetColorIndex(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_e20d08f7af3e7323.ColorIndexAt(int(__obf_4e172b662bf3e6a8), int(__obf_82a1080185d9e2b7)))
		}
	}
	__obf_a7ac3c0c472a6787 := __obf_51b09dba8f91c6f5.Bounds().Max.X
	__obf_e6eb35f8a1f9e292 := __obf_51b09dba8f91c6f5.Bounds().Max.Y
	for __obf_e487b3618d36e102 := range __obf_a7ac3c0c472a6787 {
		for __obf_8cbd87ef15403513 := range __obf_e6eb35f8a1f9e292 {
			__obf_2e70b49a4a643481.
				SetColorIndex(__obf_e487b3618d36e102, __obf_8cbd87ef15403513,

					// DrawCircle draws a circle
					__obf_51b09dba8f91c6f5.ColorIndexAt(__obf_e487b3618d36e102, __obf_8cbd87ef15403513))
		}
	}
}

func (__obf_2e70b49a4a643481 *Palette) DrawCircle(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_2ff57e13a5c6715c int, __obf_dbbef660ec878016 color.RGBA) {
	__obf_4826f2d76e2f158f := 1 - __obf_2ff57e13a5c6715c
	__obf_9018859d4affdead := 1
	__obf_156e7e2332409226 := -2 * __obf_2ff57e13a5c6715c
	__obf_21d568c8af477352 := 0
	__obf_15ed10e49b097b9b := __obf_2ff57e13a5c6715c
	__obf_2e70b49a4a643481.
		Set(__obf_e487b3618d36e102, __obf_8cbd87ef15403513+__obf_2ff57e13a5c6715c, __obf_dbbef660ec878016)
	__obf_2e70b49a4a643481.
		Set(__obf_e487b3618d36e102, __obf_8cbd87ef15403513-__obf_2ff57e13a5c6715c, __obf_dbbef660ec878016)
	__obf_2e70b49a4a643481.
		DrawHorizLine(__obf_e487b3618d36e102-__obf_2ff57e13a5c6715c, __obf_e487b3618d36e102+__obf_2ff57e13a5c6715c, __obf_8cbd87ef15403513, __obf_dbbef660ec878016)

	for __obf_21d568c8af477352 < __obf_15ed10e49b097b9b {
		if __obf_4826f2d76e2f158f >= 0 {
			__obf_15ed10e49b097b9b--
			__obf_156e7e2332409226 += 2
			__obf_4826f2d76e2f158f += __obf_156e7e2332409226
		}
		__obf_21d568c8af477352++
		__obf_9018859d4affdead += 2
		__obf_4826f2d76e2f158f += __obf_9018859d4affdead
		__obf_2e70b49a4a643481.
			DrawHorizLine(__obf_e487b3618d36e102-__obf_21d568c8af477352, __obf_e487b3618d36e102+__obf_21d568c8af477352, __obf_8cbd87ef15403513+__obf_15ed10e49b097b9b, __obf_dbbef660ec878016)
		__obf_2e70b49a4a643481.
			DrawHorizLine(__obf_e487b3618d36e102-__obf_21d568c8af477352, __obf_e487b3618d36e102+__obf_21d568c8af477352, __obf_8cbd87ef15403513-__obf_15ed10e49b097b9b, __obf_dbbef660ec878016)
		__obf_2e70b49a4a643481.
			DrawHorizLine(__obf_e487b3618d36e102-__obf_15ed10e49b097b9b, __obf_e487b3618d36e102+__obf_15ed10e49b097b9b, __obf_8cbd87ef15403513+__obf_21d568c8af477352, __obf_dbbef660ec878016)
		__obf_2e70b49a4a643481.
			DrawHorizLine(__obf_e487b3618d36e102-__obf_15ed10e49b097b9b, __obf_e487b3618d36e102+__obf_15ed10e49b097b9b,

				// DrawHorizLine draws a horizontal line
				__obf_8cbd87ef15403513-__obf_21d568c8af477352, __obf_dbbef660ec878016)
	}
}

func (__obf_2e70b49a4a643481 *Palette) DrawHorizLine(__obf_16602036e0d7dca7, __obf_19f8fccc6b52b371, __obf_8cbd87ef15403513 int, __obf_dbbef660ec878016 color.RGBA) {
	for __obf_e487b3618d36e102 := __obf_16602036e0d7dca7; __obf_e487b3618d36e102 <= __obf_19f8fccc6b52b371; __obf_e487b3618d36e102++ {
		__obf_2e70b49a4a643481.
			Set(__obf_e487b3618d36e102,

				// Distort distorts the canvas
				__obf_8cbd87ef15403513, __obf_dbbef660ec878016)
	}
}

func (__obf_2e70b49a4a643481 *Palette) Distort(__obf_efaf0cb33c8a76e0 float64, __obf_24a6403ac72cd3e9 float64) {
	__obf_b0bdb05fb01e5fd8 := __obf_2e70b49a4a643481.Bounds().Max.X
	__obf_e0d45be06f8e29d9 := __obf_2e70b49a4a643481.Bounds().Max.Y
	__obf_807f781f433abaa5 := NewPalette(image.Rect(0, 0, __obf_b0bdb05fb01e5fd8, __obf_e0d45be06f8e29d9), __obf_2e70b49a4a643481.Palette)
	__obf_ed508a8173048d4d := 2.0 * math.Pi / __obf_24a6403ac72cd3e9
	for __obf_e487b3618d36e102 := range __obf_b0bdb05fb01e5fd8 {
		for __obf_8cbd87ef15403513 := 0; __obf_8cbd87ef15403513 < __obf_e0d45be06f8e29d9; __obf_8cbd87ef15403513++ {
			__obf_21d568c8af477352 := __obf_efaf0cb33c8a76e0 * math.Sin(float64(__obf_8cbd87ef15403513)*__obf_ed508a8173048d4d)
			__obf_15ed10e49b097b9b := __obf_efaf0cb33c8a76e0 * math.Cos(float64(__obf_e487b3618d36e102)*__obf_ed508a8173048d4d)
			__obf_807f781f433abaa5.
				SetColorIndex(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_2e70b49a4a643481.ColorIndexAt(__obf_e487b3618d36e102+int(__obf_21d568c8af477352), __obf_8cbd87ef15403513+int(__obf_15ed10e49b097b9b)))
		}
	}
	__obf_a7ac3c0c472a6787 := __obf_807f781f433abaa5.Bounds().Max.X
	__obf_e6eb35f8a1f9e292 := __obf_807f781f433abaa5.Bounds().Max.Y
	for __obf_e487b3618d36e102 := range __obf_a7ac3c0c472a6787 {
		for __obf_8cbd87ef15403513 := 0; __obf_8cbd87ef15403513 < __obf_e6eb35f8a1f9e292; __obf_8cbd87ef15403513++ {
			__obf_2e70b49a4a643481.
				SetColorIndex(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_807f781f433abaa5.ColorIndexAt(__obf_e487b3618d36e102, __obf_8cbd87ef15403513))
		}
	}
	__obf_807f781f433abaa5.
		Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_2e70b49a4a643481 *Palette) DrawBeeline(__obf_b5ba889f6a8e65f5 image.Point, __obf_c353584c268e7cd3 image.Point, __obf_e017af0fa030338c color.RGBA) {
	__obf_ed508a8173048d4d := math.Abs(float64(__obf_b5ba889f6a8e65f5.X - __obf_c353584c268e7cd3.X))
	__obf_c8ea7f83d06e4eb7 := math.Abs(float64(__obf_c353584c268e7cd3.Y - __obf_b5ba889f6a8e65f5.Y))
	__obf_f5ab01d88f3e3c64, __obf_0014b3ba25267477 := 1, 1
	if __obf_b5ba889f6a8e65f5.X >= __obf_c353584c268e7cd3.X {
		__obf_f5ab01d88f3e3c64 = -1
	}
	if __obf_b5ba889f6a8e65f5.Y >= __obf_c353584c268e7cd3.Y {
		__obf_0014b3ba25267477 = -1
	}
	__obf_45f5b0f1f05cc095 := __obf_ed508a8173048d4d - __obf_c8ea7f83d06e4eb7
	for {
		__obf_2e70b49a4a643481.
			Set(__obf_b5ba889f6a8e65f5.X, __obf_b5ba889f6a8e65f5.Y, __obf_e017af0fa030338c)
		__obf_2e70b49a4a643481.
			Set(__obf_b5ba889f6a8e65f5.X+1, __obf_b5ba889f6a8e65f5.Y, __obf_e017af0fa030338c)
		__obf_2e70b49a4a643481.
			Set(__obf_b5ba889f6a8e65f5.X-1, __obf_b5ba889f6a8e65f5.Y, __obf_e017af0fa030338c)
		__obf_2e70b49a4a643481.
			Set(__obf_b5ba889f6a8e65f5.X+2, __obf_b5ba889f6a8e65f5.Y, __obf_e017af0fa030338c)
		__obf_2e70b49a4a643481.
			Set(__obf_b5ba889f6a8e65f5.X-2, __obf_b5ba889f6a8e65f5.Y, __obf_e017af0fa030338c)
		if __obf_b5ba889f6a8e65f5.X == __obf_c353584c268e7cd3.X && __obf_b5ba889f6a8e65f5.Y == __obf_c353584c268e7cd3.Y {
			return
		}
		__obf_2296757c6c4a29c6 := __obf_45f5b0f1f05cc095 * 2
		if __obf_2296757c6c4a29c6 > -__obf_c8ea7f83d06e4eb7 {
			__obf_45f5b0f1f05cc095 -= __obf_c8ea7f83d06e4eb7
			__obf_b5ba889f6a8e65f5.
				X += __obf_f5ab01d88f3e3c64
		}
		if __obf_2296757c6c4a29c6 < __obf_ed508a8173048d4d {
			__obf_45f5b0f1f05cc095 += __obf_ed508a8173048d4d
			__obf_b5ba889f6a8e65f5.
				Y += __obf_0014b3ba25267477
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_2e70b49a4a643481 *Palette) AngleSwapPoint(__obf_e487b3618d36e102, __obf_8cbd87ef15403513, __obf_662a4428339c8e5b, __obf_2b5072b762a61186 float64) (__obf_26d03c91bdc3bec7, __obf_557816fb587a2d2e float64) {
	__obf_e487b3618d36e102 -= __obf_662a4428339c8e5b
	__obf_8cbd87ef15403513 = __obf_662a4428339c8e5b - __obf_8cbd87ef15403513
	__obf_abb431c9dd0acd15 := math.Sin(__obf_2b5072b762a61186 * (math.Pi / 180))
	__obf_0754bf4ae3c4ae6c := math.Cos(__obf_2b5072b762a61186 * (math.Pi / 180))
	__obf_26d03c91bdc3bec7 = __obf_e487b3618d36e102*__obf_0754bf4ae3c4ae6c + __obf_8cbd87ef15403513*__obf_abb431c9dd0acd15
	__obf_557816fb587a2d2e = -__obf_e487b3618d36e102*__obf_abb431c9dd0acd15 + __obf_8cbd87ef15403513*__obf_0754bf4ae3c4ae6c
	__obf_26d03c91bdc3bec7 += __obf_662a4428339c8e5b
	__obf_557816fb587a2d2e = __obf_662a4428339c8e5b - __obf_557816fb587a2d2e
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_2e70b49a4a643481 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_a7ac3c0c472a6787 := __obf_2e70b49a4a643481.Bounds().Max.X
	__obf_e6eb35f8a1f9e292 := __obf_2e70b49a4a643481.Bounds().Max.Y
	__obf_4a5b8d8ec177e616 := __obf_a7ac3c0c472a6787
	__obf_37b7b15db85179cd := 0
	__obf_f5116c14efec20ac := __obf_e6eb35f8a1f9e292
	__obf_436aa8c9e2895e2a := 0
	for __obf_e487b3618d36e102 := range __obf_a7ac3c0c472a6787 {
		for __obf_8cbd87ef15403513 := range __obf_e6eb35f8a1f9e292 {
			__obf_15fd876741357eb2 := __obf_2e70b49a4a643481.At(__obf_e487b3618d36e102, __obf_8cbd87ef15403513)
			if _, _, _, __obf_2b033369da82f821 := __obf_15fd876741357eb2.RGBA(); __obf_2b033369da82f821 > 0 {
				if __obf_e487b3618d36e102 < __obf_4a5b8d8ec177e616 {
					__obf_4a5b8d8ec177e616 = __obf_e487b3618d36e102
				}
				if __obf_e487b3618d36e102 > __obf_37b7b15db85179cd {
					__obf_37b7b15db85179cd = __obf_e487b3618d36e102
				}

				if __obf_8cbd87ef15403513 < __obf_f5116c14efec20ac {
					__obf_f5116c14efec20ac = __obf_8cbd87ef15403513
				}
				if __obf_8cbd87ef15403513 > __obf_436aa8c9e2895e2a {
					__obf_436aa8c9e2895e2a = __obf_8cbd87ef15403513
				}
			}
		}
	}
	__obf_4a5b8d8ec177e616 = int(max(0, float64(__obf_4a5b8d8ec177e616-2)))
	__obf_37b7b15db85179cd = int(min(float64(__obf_a7ac3c0c472a6787), float64(__obf_37b7b15db85179cd+2)))
	__obf_f5116c14efec20ac = int(max(0, float64(__obf_f5116c14efec20ac-2)))
	__obf_436aa8c9e2895e2a = int(min(float64(__obf_e6eb35f8a1f9e292), float64(__obf_436aa8c9e2895e2a+2)))

	return &AreaRect{__obf_4a5b8d8ec177e616, __obf_37b7b15db85179cd, __obf_f5116c14efec20ac, __obf_436aa8c9e2895e2a}
}
