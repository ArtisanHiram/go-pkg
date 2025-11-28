package __obf_deb7e65d40f46bd0

import (
	"image"
	"image/color"
	"math"
)

// NewPalette creates a palette canvas
func NewPalette(__obf_941de68f67190911 image.Rectangle, __obf_3fa3e6b785f3e768 color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(__obf_941de68f67190911, __obf_3fa3e6b785f3e768),
	}
}

// palette struct for palette canvas
type Palette struct {
	*image.Paletted
}

// Rotate rotates the canvas by any angle
func (__obf_3fa3e6b785f3e768 *Palette) Rotate(__obf_0a4f6c9b020fe9b9 int) {
	if __obf_0a4f6c9b020fe9b9 == 0 {
		return
	}

	__obf_8ea40b6dbb37b06f := __obf_3fa3e6b785f3e768
	__obf_b418adbaa6995e8e := __obf_8ea40b6dbb37b06f.Bounds().Max.X
	__obf_0f52812d1cb99ebd := __obf_8ea40b6dbb37b06f.Bounds().Max.Y
	__obf_941de68f67190911 := __obf_b418adbaa6995e8e / 2
	__obf_ffecc9d882361ef9 := image.NewPaletted(image.Rect(0, 0, __obf_b418adbaa6995e8e, __obf_0f52812d1cb99ebd), __obf_8ea40b6dbb37b06f.Palette)
	for __obf_ce61f234cbe42808 := 0; __obf_ce61f234cbe42808 <= __obf_ffecc9d882361ef9.Bounds().Max.X; __obf_ce61f234cbe42808++ {
		for __obf_7a88a48a3d7db1aa := 0; __obf_7a88a48a3d7db1aa <= __obf_ffecc9d882361ef9.Bounds().Max.Y; __obf_7a88a48a3d7db1aa++ {
			__obf_657f5a2a9c55c188, __obf_3130ae025ccdf88a := __obf_3fa3e6b785f3e768.AngleSwapPoint(float64(__obf_ce61f234cbe42808), float64(__obf_7a88a48a3d7db1aa), float64(__obf_941de68f67190911), float64(__obf_0a4f6c9b020fe9b9))
			__obf_ffecc9d882361ef9.SetColorIndex(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_8ea40b6dbb37b06f.ColorIndexAt(int(__obf_657f5a2a9c55c188), int(__obf_3130ae025ccdf88a)))
		}
	}

	__obf_c04c6105d2f272af := __obf_ffecc9d882361ef9.Bounds().Max.X
	__obf_7cd3f0787f1e8091 := __obf_ffecc9d882361ef9.Bounds().Max.Y
	for __obf_ce61f234cbe42808 := range __obf_c04c6105d2f272af {
		for __obf_7a88a48a3d7db1aa := range __obf_7cd3f0787f1e8091 {
			__obf_3fa3e6b785f3e768.SetColorIndex(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_ffecc9d882361ef9.ColorIndexAt(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa))
		}
	}
}

// DrawCircle draws a circle
func (__obf_3fa3e6b785f3e768 *Palette) DrawCircle(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_77b9044f3fa7cc43 int, __obf_f012bb226e8e4a2d color.RGBA) {
	__obf_71778c093f38dff0 := 1 - __obf_77b9044f3fa7cc43
	__obf_5b232bff284b3c36 := 1
	__obf_0c7c695daa9fa249 := -2 * __obf_77b9044f3fa7cc43
	__obf_c048d7e761d12429 := 0
	__obf_a1fd94d901fd046a := __obf_77b9044f3fa7cc43

	__obf_3fa3e6b785f3e768.Set(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa+__obf_77b9044f3fa7cc43, __obf_f012bb226e8e4a2d)
	__obf_3fa3e6b785f3e768.Set(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa-__obf_77b9044f3fa7cc43, __obf_f012bb226e8e4a2d)
	__obf_3fa3e6b785f3e768.DrawHorizLine(__obf_ce61f234cbe42808-__obf_77b9044f3fa7cc43, __obf_ce61f234cbe42808+__obf_77b9044f3fa7cc43, __obf_7a88a48a3d7db1aa, __obf_f012bb226e8e4a2d)

	for __obf_c048d7e761d12429 < __obf_a1fd94d901fd046a {
		if __obf_71778c093f38dff0 >= 0 {
			__obf_a1fd94d901fd046a--
			__obf_0c7c695daa9fa249 += 2
			__obf_71778c093f38dff0 += __obf_0c7c695daa9fa249
		}
		__obf_c048d7e761d12429++
		__obf_5b232bff284b3c36 += 2
		__obf_71778c093f38dff0 += __obf_5b232bff284b3c36
		__obf_3fa3e6b785f3e768.DrawHorizLine(__obf_ce61f234cbe42808-__obf_c048d7e761d12429, __obf_ce61f234cbe42808+__obf_c048d7e761d12429, __obf_7a88a48a3d7db1aa+__obf_a1fd94d901fd046a, __obf_f012bb226e8e4a2d)
		__obf_3fa3e6b785f3e768.DrawHorizLine(__obf_ce61f234cbe42808-__obf_c048d7e761d12429, __obf_ce61f234cbe42808+__obf_c048d7e761d12429, __obf_7a88a48a3d7db1aa-__obf_a1fd94d901fd046a, __obf_f012bb226e8e4a2d)
		__obf_3fa3e6b785f3e768.DrawHorizLine(__obf_ce61f234cbe42808-__obf_a1fd94d901fd046a, __obf_ce61f234cbe42808+__obf_a1fd94d901fd046a, __obf_7a88a48a3d7db1aa+__obf_c048d7e761d12429, __obf_f012bb226e8e4a2d)
		__obf_3fa3e6b785f3e768.DrawHorizLine(__obf_ce61f234cbe42808-__obf_a1fd94d901fd046a, __obf_ce61f234cbe42808+__obf_a1fd94d901fd046a, __obf_7a88a48a3d7db1aa-__obf_c048d7e761d12429, __obf_f012bb226e8e4a2d)
	}
}

// DrawHorizLine draws a horizontal line
func (__obf_3fa3e6b785f3e768 *Palette) DrawHorizLine(__obf_c7967199aceaa200, __obf_c7de3502aaaf5ec4, __obf_7a88a48a3d7db1aa int, __obf_f012bb226e8e4a2d color.RGBA) {
	for __obf_ce61f234cbe42808 := __obf_c7967199aceaa200; __obf_ce61f234cbe42808 <= __obf_c7de3502aaaf5ec4; __obf_ce61f234cbe42808++ {
		__obf_3fa3e6b785f3e768.Set(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_f012bb226e8e4a2d)
	}
}

// Distort distorts the canvas
func (__obf_3fa3e6b785f3e768 *Palette) Distort(__obf_15815fa32d23a301 float64, __obf_06f41b3abd23ff33 float64) {
	__obf_3fbfb4aba2bc08f3 := __obf_3fa3e6b785f3e768.Bounds().Max.X
	__obf_7113bad0317e187f := __obf_3fa3e6b785f3e768.Bounds().Max.Y
	__obf_34bffb9cded95529 := NewPalette(image.Rect(0, 0, __obf_3fbfb4aba2bc08f3, __obf_7113bad0317e187f), __obf_3fa3e6b785f3e768.Palette)
	__obf_a0ccd38eacb0134e := 2.0 * math.Pi / __obf_06f41b3abd23ff33
	for __obf_ce61f234cbe42808 := range __obf_3fbfb4aba2bc08f3 {
		for __obf_7a88a48a3d7db1aa := 0; __obf_7a88a48a3d7db1aa < __obf_7113bad0317e187f; __obf_7a88a48a3d7db1aa++ {
			__obf_c048d7e761d12429 := __obf_15815fa32d23a301 * math.Sin(float64(__obf_7a88a48a3d7db1aa)*__obf_a0ccd38eacb0134e)
			__obf_a1fd94d901fd046a := __obf_15815fa32d23a301 * math.Cos(float64(__obf_ce61f234cbe42808)*__obf_a0ccd38eacb0134e)
			__obf_34bffb9cded95529.SetColorIndex(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_3fa3e6b785f3e768.ColorIndexAt(__obf_ce61f234cbe42808+int(__obf_c048d7e761d12429), __obf_7a88a48a3d7db1aa+int(__obf_a1fd94d901fd046a)))
		}
	}

	__obf_c04c6105d2f272af := __obf_34bffb9cded95529.Bounds().Max.X
	__obf_7cd3f0787f1e8091 := __obf_34bffb9cded95529.Bounds().Max.Y
	for __obf_ce61f234cbe42808 := range __obf_c04c6105d2f272af {
		for __obf_7a88a48a3d7db1aa := 0; __obf_7a88a48a3d7db1aa < __obf_7cd3f0787f1e8091; __obf_7a88a48a3d7db1aa++ {
			__obf_3fa3e6b785f3e768.SetColorIndex(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_34bffb9cded95529.ColorIndexAt(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa))
		}
	}

	__obf_34bffb9cded95529.Palette = nil
}

// DrawBeeline draws a straight line
func (__obf_3fa3e6b785f3e768 *Palette) DrawBeeline(__obf_b4717c71b33a7b07 image.Point, __obf_2427229f0405eb35 image.Point, __obf_4ffa90e060d53745 color.RGBA) {
	__obf_a0ccd38eacb0134e := math.Abs(float64(__obf_b4717c71b33a7b07.X - __obf_2427229f0405eb35.X))
	__obf_4199cd4b4fbd0e10 := math.Abs(float64(__obf_2427229f0405eb35.Y - __obf_b4717c71b33a7b07.Y))
	__obf_8961456c8cdcdfb4, __obf_d42d1707feb2ec3a := 1, 1
	if __obf_b4717c71b33a7b07.X >= __obf_2427229f0405eb35.X {
		__obf_8961456c8cdcdfb4 = -1
	}
	if __obf_b4717c71b33a7b07.Y >= __obf_2427229f0405eb35.Y {
		__obf_d42d1707feb2ec3a = -1
	}
	__obf_cdc2a8664de7ba0a := __obf_a0ccd38eacb0134e - __obf_4199cd4b4fbd0e10
	for {
		__obf_3fa3e6b785f3e768.Set(__obf_b4717c71b33a7b07.X, __obf_b4717c71b33a7b07.Y, __obf_4ffa90e060d53745)
		__obf_3fa3e6b785f3e768.Set(__obf_b4717c71b33a7b07.X+1, __obf_b4717c71b33a7b07.Y, __obf_4ffa90e060d53745)
		__obf_3fa3e6b785f3e768.Set(__obf_b4717c71b33a7b07.X-1, __obf_b4717c71b33a7b07.Y, __obf_4ffa90e060d53745)
		__obf_3fa3e6b785f3e768.Set(__obf_b4717c71b33a7b07.X+2, __obf_b4717c71b33a7b07.Y, __obf_4ffa90e060d53745)
		__obf_3fa3e6b785f3e768.Set(__obf_b4717c71b33a7b07.X-2, __obf_b4717c71b33a7b07.Y, __obf_4ffa90e060d53745)
		if __obf_b4717c71b33a7b07.X == __obf_2427229f0405eb35.X && __obf_b4717c71b33a7b07.Y == __obf_2427229f0405eb35.Y {
			return
		}
		__obf_e6803a22a847b029 := __obf_cdc2a8664de7ba0a * 2
		if __obf_e6803a22a847b029 > -__obf_4199cd4b4fbd0e10 {
			__obf_cdc2a8664de7ba0a -= __obf_4199cd4b4fbd0e10
			__obf_b4717c71b33a7b07.X += __obf_8961456c8cdcdfb4
		}
		if __obf_e6803a22a847b029 < __obf_a0ccd38eacb0134e {
			__obf_cdc2a8664de7ba0a += __obf_a0ccd38eacb0134e
			__obf_b4717c71b33a7b07.Y += __obf_d42d1707feb2ec3a
		}
	}
}

// AngleSwapPoint converts point coordinates based on angle
func (__obf_3fa3e6b785f3e768 *Palette) AngleSwapPoint(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa, __obf_941de68f67190911, __obf_0a4f6c9b020fe9b9 float64) (__obf_9778367ade944aea, __obf_e2de0a93f479457b float64) {
	__obf_ce61f234cbe42808 -= __obf_941de68f67190911
	__obf_7a88a48a3d7db1aa = __obf_941de68f67190911 - __obf_7a88a48a3d7db1aa
	__obf_d80dcee345236159 := math.Sin(__obf_0a4f6c9b020fe9b9 * (math.Pi / 180))
	__obf_8ce704f7782816f5 := math.Cos(__obf_0a4f6c9b020fe9b9 * (math.Pi / 180))
	__obf_9778367ade944aea = __obf_ce61f234cbe42808*__obf_8ce704f7782816f5 + __obf_7a88a48a3d7db1aa*__obf_d80dcee345236159
	__obf_e2de0a93f479457b = -__obf_ce61f234cbe42808*__obf_d80dcee345236159 + __obf_7a88a48a3d7db1aa*__obf_8ce704f7782816f5
	__obf_9778367ade944aea += __obf_941de68f67190911
	__obf_e2de0a93f479457b = __obf_941de68f67190911 - __obf_e2de0a93f479457b
	return
}

// CalcMarginBlankArea calculates the blank area of the canvas
func (__obf_3fa3e6b785f3e768 *Palette) CalcMarginBlankArea() *AreaRect {
	__obf_c04c6105d2f272af := __obf_3fa3e6b785f3e768.Bounds().Max.X
	__obf_7cd3f0787f1e8091 := __obf_3fa3e6b785f3e768.Bounds().Max.Y
	__obf_8f8dd92499417004 := __obf_c04c6105d2f272af
	__obf_ac672b1f936c0c6e := 0
	__obf_9ce451081920ce95 := __obf_7cd3f0787f1e8091
	__obf_71ac0e772fc668a4 := 0
	for __obf_ce61f234cbe42808 := range __obf_c04c6105d2f272af {
		for __obf_7a88a48a3d7db1aa := range __obf_7cd3f0787f1e8091 {
			__obf_ceed5a3454040e74 := __obf_3fa3e6b785f3e768.At(__obf_ce61f234cbe42808, __obf_7a88a48a3d7db1aa)
			if _, _, _, __obf_a66951353e16056a := __obf_ceed5a3454040e74.RGBA(); __obf_a66951353e16056a > 0 {
				if __obf_ce61f234cbe42808 < __obf_8f8dd92499417004 {
					__obf_8f8dd92499417004 = __obf_ce61f234cbe42808
				}
				if __obf_ce61f234cbe42808 > __obf_ac672b1f936c0c6e {
					__obf_ac672b1f936c0c6e = __obf_ce61f234cbe42808
				}

				if __obf_7a88a48a3d7db1aa < __obf_9ce451081920ce95 {
					__obf_9ce451081920ce95 = __obf_7a88a48a3d7db1aa
				}
				if __obf_7a88a48a3d7db1aa > __obf_71ac0e772fc668a4 {
					__obf_71ac0e772fc668a4 = __obf_7a88a48a3d7db1aa
				}
			}
		}
	}

	__obf_8f8dd92499417004 = int(max(0, float64(__obf_8f8dd92499417004-2)))
	__obf_ac672b1f936c0c6e = int(min(float64(__obf_c04c6105d2f272af), float64(__obf_ac672b1f936c0c6e+2)))
	__obf_9ce451081920ce95 = int(max(0, float64(__obf_9ce451081920ce95-2)))
	__obf_71ac0e772fc668a4 = int(min(float64(__obf_7cd3f0787f1e8091), float64(__obf_71ac0e772fc668a4+2)))

	return &AreaRect{
		__obf_8f8dd92499417004,
		__obf_ac672b1f936c0c6e,
		__obf_9ce451081920ce95,
		__obf_71ac0e772fc668a4,
	}
}
