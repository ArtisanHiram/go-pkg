package __obf_54406b1a1de84196

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_e73d974dafb71af7 image.Rectangle, __obf_ddd45669640656a2 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_e73d974dafb71af7,

			// palette struct for palette canvas
			__obf_ddd45669640656a2),
	}
}

type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_ddd45669640656a2 *Palette) Rotate(__obf_27c01b4675c92ab0 int) {
	if __obf_27c01b4675c92ab0 == 0 {
		return
	}
	__obf_f62553e9e9bb0107 := __obf_ddd45669640656a2
	__obf_58b88ba335f0f5cc := __obf_f62553e9e9bb0107.Bounds().Max.X
	__obf_a902e6c896fe3845 := __obf_f62553e9e9bb0107.Bounds().Max.Y
	__obf_e73d974dafb71af7 := __obf_58b88ba335f0f5cc / 2
	__obf_69442934a4ccd6bd := image.NewPaletted(image.Rect(0, 0, __obf_58b88ba335f0f5cc, __obf_a902e6c896fe3845), __obf_f62553e9e9bb0107.Palette)
	for __obf_48afb3005cd4a35c := 0; __obf_48afb3005cd4a35c <= __obf_69442934a4ccd6bd.Bounds().Max.X; __obf_48afb3005cd4a35c++ {
		for __obf_33a1f511b09ac2af := 0; __obf_33a1f511b09ac2af <= __obf_69442934a4ccd6bd.Bounds().Max.Y; __obf_33a1f511b09ac2af++ {
			__obf_c958ef268cf1b10c, __obf_264563f1045a41c1 := __obf_ddd45669640656a2.AngleSwapPoint(float64(__obf_48afb3005cd4a35c), float64(__obf_33a1f511b09ac2af), float64(__obf_e73d974dafb71af7), float64(__obf_27c01b4675c92ab0))
			__obf_69442934a4ccd6bd.
				SetColorIndex(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_f62553e9e9bb0107.ColorIndexAt(int(__obf_c958ef268cf1b10c), int(__obf_264563f1045a41c1)))
		}
	}
	__obf_5883e9285a516ece := __obf_69442934a4ccd6bd.Bounds().Max.X
	__obf_0f8e181d560d05a0 := __obf_69442934a4ccd6bd.Bounds().Max.Y
	for __obf_48afb3005cd4a35c := range __obf_5883e9285a516ece {
		for __obf_33a1f511b09ac2af := range __obf_0f8e181d560d05a0 {
			__obf_ddd45669640656a2.
				SetColorIndex(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af,

					// DrawCircle draws a circle
					__obf_69442934a4ccd6bd.ColorIndexAt(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af))
		}
	}
}

func (__obf_ddd45669640656a2 *Palette) DrawCircle(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_9dba054a785bce28 int, __obf_a26ee860b3cee35b color.RGBA) {
	__obf_e031ac4f64419f3f := 1 - __obf_9dba054a785bce28
	__obf_2b3a82cd54d0ddfe := 1
	__obf_54c369f529849e6d := -2 * __obf_9dba054a785bce28
	__obf_f2464594091f361d := 0
	__obf_178f6c958fb18eeb := __obf_9dba054a785bce28
	__obf_ddd45669640656a2.
		Set(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af+__obf_9dba054a785bce28, __obf_a26ee860b3cee35b)
	__obf_ddd45669640656a2.
		Set(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af-__obf_9dba054a785bce28, __obf_a26ee860b3cee35b)
	__obf_ddd45669640656a2.
		DrawHorizLine(__obf_48afb3005cd4a35c-__obf_9dba054a785bce28, __obf_48afb3005cd4a35c+__obf_9dba054a785bce28, __obf_33a1f511b09ac2af, __obf_a26ee860b3cee35b)

	for __obf_f2464594091f361d < __obf_178f6c958fb18eeb {
		if __obf_e031ac4f64419f3f >= 0 {
			__obf_178f6c958fb18eeb--
			__obf_54c369f529849e6d += 2
			__obf_e031ac4f64419f3f += __obf_54c369f529849e6d
		}
		__obf_f2464594091f361d++
		__obf_2b3a82cd54d0ddfe += 2
		__obf_e031ac4f64419f3f += __obf_2b3a82cd54d0ddfe
		__obf_ddd45669640656a2.
			DrawHorizLine(__obf_48afb3005cd4a35c-__obf_f2464594091f361d, __obf_48afb3005cd4a35c+__obf_f2464594091f361d, __obf_33a1f511b09ac2af+__obf_178f6c958fb18eeb, __obf_a26ee860b3cee35b)
		__obf_ddd45669640656a2.
			DrawHorizLine(__obf_48afb3005cd4a35c-__obf_f2464594091f361d, __obf_48afb3005cd4a35c+__obf_f2464594091f361d, __obf_33a1f511b09ac2af-__obf_178f6c958fb18eeb, __obf_a26ee860b3cee35b)
		__obf_ddd45669640656a2.
			DrawHorizLine(__obf_48afb3005cd4a35c-__obf_178f6c958fb18eeb, __obf_48afb3005cd4a35c+__obf_178f6c958fb18eeb, __obf_33a1f511b09ac2af+__obf_f2464594091f361d, __obf_a26ee860b3cee35b)
		__obf_ddd45669640656a2.
			DrawHorizLine(__obf_48afb3005cd4a35c-__obf_178f6c958fb18eeb, __obf_48afb3005cd4a35c+__obf_178f6c958fb18eeb,

				// DrawHorizLine draws a horizontal line
				__obf_33a1f511b09ac2af-__obf_f2464594091f361d, __obf_a26ee860b3cee35b)
	}
}

func (__obf_ddd45669640656a2 *Palette) DrawHorizLine(__obf_fea013f2d9472172, __obf_5acb14b57e85d18a, __obf_33a1f511b09ac2af int, __obf_a26ee860b3cee35b color.RGBA) {
	for __obf_48afb3005cd4a35c := __obf_fea013f2d9472172; __obf_48afb3005cd4a35c <= __obf_5acb14b57e85d18a; __obf_48afb3005cd4a35c++ {
		__obf_ddd45669640656a2.
			Set(__obf_48afb3005cd4a35c,

				// Distort distorts the canvas
				__obf_33a1f511b09ac2af, __obf_a26ee860b3cee35b)
	}
}

func (__obf_ddd45669640656a2 *Palette) Distort(__obf_3431916851685d17 float64, __obf_4bfb3fc0b135a4c9 float64) {
	__obf_2366a8ab85b2ae47 := __obf_ddd45669640656a2.Bounds().Max.X
	__obf_daaeaddeff0e74f7 := __obf_ddd45669640656a2.Bounds().Max.Y
	__obf_c256a6573815710e := NewPalette(image.Rect(0, 0, __obf_2366a8ab85b2ae47, __obf_daaeaddeff0e74f7), __obf_ddd45669640656a2.Palette)
	__obf_21594ae268a2eb66 := 2.0 * math.Pi / __obf_4bfb3fc0b135a4c9
	for __obf_48afb3005cd4a35c := range __obf_2366a8ab85b2ae47 {
		for __obf_33a1f511b09ac2af := 0; __obf_33a1f511b09ac2af < __obf_daaeaddeff0e74f7; __obf_33a1f511b09ac2af++ {
			__obf_f2464594091f361d := __obf_3431916851685d17 * math.Sin(float64(__obf_33a1f511b09ac2af)*__obf_21594ae268a2eb66)
			__obf_178f6c958fb18eeb := __obf_3431916851685d17 * math.Cos(float64(__obf_48afb3005cd4a35c)*__obf_21594ae268a2eb66)
			__obf_c256a6573815710e.
				SetColorIndex(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_ddd45669640656a2.ColorIndexAt(__obf_48afb3005cd4a35c+int(__obf_f2464594091f361d), __obf_33a1f511b09ac2af+int(__obf_178f6c958fb18eeb)))
		}
	}
	__obf_5883e9285a516ece := __obf_c256a6573815710e.Bounds().Max.X
	__obf_0f8e181d560d05a0 := __obf_c256a6573815710e.Bounds().Max.Y
	for __obf_48afb3005cd4a35c := range __obf_5883e9285a516ece {
		for __obf_33a1f511b09ac2af := 0; __obf_33a1f511b09ac2af < __obf_0f8e181d560d05a0; __obf_33a1f511b09ac2af++ {
			__obf_ddd45669640656a2.
				SetColorIndex(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_c256a6573815710e.ColorIndexAt(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af))
		}
	}
	__obf_c256a6573815710e.
		Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_ddd45669640656a2 *Palette) DrawBeeline(__obf_46d9564dfde8328b image.Point, __obf_a2da4d2ef85437d5 image.Point, __obf_60beead06dd7ac14 color.RGBA) {
	__obf_21594ae268a2eb66 := math.Abs(float64(__obf_46d9564dfde8328b.X - __obf_a2da4d2ef85437d5.X))
	__obf_8c43a22828ae34a2 := math.Abs(float64(__obf_a2da4d2ef85437d5.Y - __obf_46d9564dfde8328b.Y))
	__obf_8d9e93e2cec00e46, __obf_10fe1b32f925fedf := 1, 1
	if __obf_46d9564dfde8328b.X >= __obf_a2da4d2ef85437d5.X {
		__obf_8d9e93e2cec00e46 = -1
	}
	if __obf_46d9564dfde8328b.Y >= __obf_a2da4d2ef85437d5.Y {
		__obf_10fe1b32f925fedf = -1
	}
	__obf_d99742a1139768da := __obf_21594ae268a2eb66 - __obf_8c43a22828ae34a2
	for {
		__obf_ddd45669640656a2.
			Set(__obf_46d9564dfde8328b.X, __obf_46d9564dfde8328b.Y, __obf_60beead06dd7ac14)
		__obf_ddd45669640656a2.
			Set(__obf_46d9564dfde8328b.X+1, __obf_46d9564dfde8328b.Y, __obf_60beead06dd7ac14)
		__obf_ddd45669640656a2.
			Set(__obf_46d9564dfde8328b.X-1, __obf_46d9564dfde8328b.Y, __obf_60beead06dd7ac14)
		__obf_ddd45669640656a2.
			Set(__obf_46d9564dfde8328b.X+2, __obf_46d9564dfde8328b.Y, __obf_60beead06dd7ac14)
		__obf_ddd45669640656a2.
			Set(__obf_46d9564dfde8328b.X-2, __obf_46d9564dfde8328b.Y, __obf_60beead06dd7ac14)
		if __obf_46d9564dfde8328b.X == __obf_a2da4d2ef85437d5.X && __obf_46d9564dfde8328b.Y == __obf_a2da4d2ef85437d5.Y {
			return
		}
		__obf_6b750a766bb3fc9d := __obf_d99742a1139768da * 2
		if __obf_6b750a766bb3fc9d > -__obf_8c43a22828ae34a2 {
			__obf_d99742a1139768da -= __obf_8c43a22828ae34a2
			__obf_46d9564dfde8328b.
				X += __obf_8d9e93e2cec00e46
		}
		if __obf_6b750a766bb3fc9d < __obf_21594ae268a2eb66 {
			__obf_d99742a1139768da += __obf_21594ae268a2eb66
			__obf_46d9564dfde8328b.
				Y += __obf_10fe1b32f925fedf
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_ddd45669640656a2 *Palette) AngleSwapPoint(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af, __obf_e73d974dafb71af7, __obf_27c01b4675c92ab0 float64) (__obf_01607c6a0f660f66, __obf_b5dd1227e7098968 float64) {
	__obf_48afb3005cd4a35c -= __obf_e73d974dafb71af7
	__obf_33a1f511b09ac2af = __obf_e73d974dafb71af7 - __obf_33a1f511b09ac2af
	__obf_60f06896afc24770 := math.Sin(__obf_27c01b4675c92ab0 * (math.Pi / 180))
	__obf_7df3135163ec3e0a := math.Cos(__obf_27c01b4675c92ab0 * (math.Pi / 180))
	__obf_01607c6a0f660f66 = __obf_48afb3005cd4a35c*__obf_7df3135163ec3e0a + __obf_33a1f511b09ac2af*__obf_60f06896afc24770
	__obf_b5dd1227e7098968 = -__obf_48afb3005cd4a35c*__obf_60f06896afc24770 + __obf_33a1f511b09ac2af*__obf_7df3135163ec3e0a
	__obf_01607c6a0f660f66 += __obf_e73d974dafb71af7
	__obf_b5dd1227e7098968 = __obf_e73d974dafb71af7 - __obf_b5dd1227e7098968
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_ddd45669640656a2 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_5883e9285a516ece := __obf_ddd45669640656a2.Bounds().Max.X
	__obf_0f8e181d560d05a0 := __obf_ddd45669640656a2.Bounds().Max.Y
	__obf_907f96d0fe0868f4 := __obf_5883e9285a516ece
	__obf_f3f182f1c2fb0d54 := 0
	__obf_bf482b3ecc79f58c := __obf_0f8e181d560d05a0
	__obf_bcd72b5e7c0080de := 0
	for __obf_48afb3005cd4a35c := range __obf_5883e9285a516ece {
		for __obf_33a1f511b09ac2af := range __obf_0f8e181d560d05a0 {
			__obf_bbad58a05ae9b2bc := __obf_ddd45669640656a2.At(__obf_48afb3005cd4a35c, __obf_33a1f511b09ac2af)
			if _, _, _, __obf_23448ffa2f9aadbc := __obf_bbad58a05ae9b2bc.RGBA(); __obf_23448ffa2f9aadbc > 0 {
				if __obf_48afb3005cd4a35c < __obf_907f96d0fe0868f4 {
					__obf_907f96d0fe0868f4 = __obf_48afb3005cd4a35c
				}
				if __obf_48afb3005cd4a35c > __obf_f3f182f1c2fb0d54 {
					__obf_f3f182f1c2fb0d54 = __obf_48afb3005cd4a35c
				}

				if __obf_33a1f511b09ac2af < __obf_bf482b3ecc79f58c {
					__obf_bf482b3ecc79f58c = __obf_33a1f511b09ac2af
				}
				if __obf_33a1f511b09ac2af > __obf_bcd72b5e7c0080de {
					__obf_bcd72b5e7c0080de = __obf_33a1f511b09ac2af
				}
			}
		}
	}
	__obf_907f96d0fe0868f4 = int(max(0, float64(__obf_907f96d0fe0868f4-2)))
	__obf_f3f182f1c2fb0d54 = int(min(float64(__obf_5883e9285a516ece), float64(__obf_f3f182f1c2fb0d54+2)))
	__obf_bf482b3ecc79f58c = int(max(0, float64(__obf_bf482b3ecc79f58c-2)))
	__obf_bcd72b5e7c0080de = int(min(float64(__obf_0f8e181d560d05a0), float64(__obf_bcd72b5e7c0080de+2)))

	return &AreaRect{__obf_907f96d0fe0868f4, __obf_f3f182f1c2fb0d54, __obf_bf482b3ecc79f58c, __obf_bcd72b5e7c0080de}
}
