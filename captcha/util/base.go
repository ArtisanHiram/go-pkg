package __obf_ddeaac61ff7fc4bb

import (
	"errors"
	"image"
	"image/color"
	"math"
	"math/rand/v2"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

// RotatePoint rotates a point's coordinates
func RotatePoint(__obf_4dbc61c5494fe8ab, __obf_bd0ff70243ba6845, __obf_90fef108d805b4fc, __obf_c3a31c48ad1703ef float64) (float64, float64) {
	return __obf_4dbc61c5494fe8ab*__obf_c3a31c48ad1703ef - __obf_bd0ff70243ba6845*__obf_90fef108d805b4fc,

		// RotatedSize calculates the size after rotation
		__obf_4dbc61c5494fe8ab*__obf_90fef108d805b4fc + __obf_bd0ff70243ba6845*__obf_c3a31c48ad1703ef
}

func RotatedSize(__obf_1028a5fc459017b9, __obf_14f910d48cb4eb80 int, __obf_aff5632b23f86fea float64) (int, int) {
	if __obf_1028a5fc459017b9 <= 0 || __obf_14f910d48cb4eb80 <= 0 {
		return 0, 0
	}
	__obf_90fef108d805b4fc, __obf_c3a31c48ad1703ef := math.Sincos(math.Pi * __obf_aff5632b23f86fea / 180)
	__obf_4fb025352bacc555, __obf_4d09e79e1d6c28bc := RotatePoint(float64(__obf_1028a5fc459017b9-1), 0, __obf_90fef108d805b4fc, __obf_c3a31c48ad1703ef)
	__obf_9cc0ca2431118caa, __obf_e2c2a4a6c14b7a15 := RotatePoint(float64(__obf_1028a5fc459017b9-1), float64(__obf_14f910d48cb4eb80-1), __obf_90fef108d805b4fc, __obf_c3a31c48ad1703ef)
	__obf_c8f2785267b5d4ee, __obf_bc64e4ffc4507a31 := RotatePoint(0, float64(__obf_14f910d48cb4eb80-1), __obf_90fef108d805b4fc, __obf_c3a31c48ad1703ef)
	__obf_ed7cd93b80049be9 := min(__obf_4fb025352bacc555, min(__obf_9cc0ca2431118caa, min(__obf_c8f2785267b5d4ee, 0)))
	__obf_2129cf47cbc5541f := max(__obf_4fb025352bacc555, max(__obf_9cc0ca2431118caa, max(__obf_c8f2785267b5d4ee, 0)))
	__obf_5ba572cd94c1fbdc := min(__obf_4d09e79e1d6c28bc, min(__obf_e2c2a4a6c14b7a15, min(__obf_bc64e4ffc4507a31, 0)))
	__obf_d6d5f96d06009ddf := max(__obf_4d09e79e1d6c28bc, max(__obf_e2c2a4a6c14b7a15, max(__obf_bc64e4ffc4507a31, 0)))
	__obf_b152b4a52fadbc0b := __obf_2129cf47cbc5541f - __obf_ed7cd93b80049be9 + 1
	if __obf_b152b4a52fadbc0b-math.Floor(__obf_b152b4a52fadbc0b) > 0.1 {
		__obf_b152b4a52fadbc0b++
	}
	__obf_02c378c8626b024e := __obf_d6d5f96d06009ddf - __obf_5ba572cd94c1fbdc + 1
	if __obf_02c378c8626b024e-math.Floor(__obf_02c378c8626b024e) > 0.1 {
		__obf_02c378c8626b024e++
	}

	return int(__obf_b152b4a52fadbc0b), int(__obf_02c378c8626b024e)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_9140a877a8d2c499 image.Rectangle, __obf_b152b4a52fadbc0b int, __obf_02c378c8626b024e int, __obf_072cfcb2dbd5c768 bool) image.Rectangle {
	var __obf_775c82631478ee60 image.Rectangle
	if __obf_b152b4a52fadbc0b*__obf_9140a877a8d2c499.Dy() < __obf_02c378c8626b024e*__obf_9140a877a8d2c499.Dx() {
		__obf_4bfb2e0a77b80146 := float64(__obf_b152b4a52fadbc0b) / float64(__obf_9140a877a8d2c499.Dx())
		__obf_46a91062e8f9cc74 := int(float64(__obf_9140a877a8d2c499.Dy()) * __obf_4bfb2e0a77b80146)
		__obf_743f6ebb7681924e := 0
		if __obf_072cfcb2dbd5c768 {
			__obf_743f6ebb7681924e = (__obf_02c378c8626b024e - __obf_46a91062e8f9cc74) / 2
		}
		__obf_775c82631478ee60 = image.Rect(0, __obf_743f6ebb7681924e, __obf_b152b4a52fadbc0b, __obf_743f6ebb7681924e+__obf_46a91062e8f9cc74)
	} else {
		__obf_4bfb2e0a77b80146 := float64(__obf_02c378c8626b024e) / float64(__obf_9140a877a8d2c499.Dy())
		__obf_6a004954fd390211 := int(float64(__obf_9140a877a8d2c499.Dx()) * __obf_4bfb2e0a77b80146)
		__obf_743f6ebb7681924e := 0
		if __obf_072cfcb2dbd5c768 {
			__obf_743f6ebb7681924e = (__obf_b152b4a52fadbc0b - __obf_6a004954fd390211) / 2
		}
		__obf_775c82631478ee60 = image.Rect(__obf_743f6ebb7681924e, 0, __obf_743f6ebb7681924e+__obf_6a004954fd390211, __obf_02c378c8626b024e)
	}

	return __obf_775c82631478ee60
}

// t2x converts an integer to a hexadecimal string
func __obf_45103f49bb31315c(__obf_7bab9517d2a781a5 int64) string {
	__obf_ffc645f0851bfaf1 := strconv.FormatInt(__obf_7bab9517d2a781a5, 16)
	if len(__obf_ffc645f0851bfaf1) == 1 {
		__obf_ffc645f0851bfaf1 = "0" + __obf_ffc645f0851bfaf1
	}
	return __obf_ffc645f0851bfaf1
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_9610c58db7ca4318 float32) uint8 {
	__obf_02d34408cb67890e := min(float64(__obf_9610c58db7ca4318), 1)
	__obf_d33c7675d29bf228 := __obf_02d34408cb67890e * 255
	return uint8(__obf_d33c7675d29bf228)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_ee0e1f04d991f920 int64, __obf_02a9a17ef92a471d int64, __obf_8761840955d8173b int64) string {
	__obf_3057df8ae9885417 := __obf_45103f49bb31315c(__obf_ee0e1f04d991f920)
	__obf_6bb01effa05ab1ad := __obf_45103f49bb31315c(__obf_02a9a17ef92a471d)
	__obf_7723b6938513fa3a := __obf_45103f49bb31315c(__obf_8761840955d8173b)
	return __obf_3057df8ae9885417 + __obf_6bb01effa05ab1ad + __obf_7723b6938513fa3a
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_e8ad821f14f8fd33 string) (int64, int64, int64) {
	__obf_3057df8ae9885417, _ := strconv.ParseInt(__obf_e8ad821f14f8fd33[:2], 16, 10)
	__obf_6bb01effa05ab1ad, _ := strconv.ParseInt(__obf_e8ad821f14f8fd33[2:4], 16, 18)
	__obf_7723b6938513fa3a, _ := strconv.ParseInt(__obf_e8ad821f14f8fd33[4:], 16, 10)
	return __obf_3057df8ae9885417, __obf_6bb01effa05ab1ad, __obf_7723b6938513fa3a
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_4e82bdaa91dfc0b2 string) (__obf_8336b398b6a5d0e3 color.RGBA, __obf_630bf6ccf46ac8b5 error) {
	__obf_8336b398b6a5d0e3.
		A = 0xff
	if __obf_4e82bdaa91dfc0b2[0] != '#' {
		return __obf_8336b398b6a5d0e3, ColorHexFormatErr
	}
	__obf_094984e50b675eb7 := func(__obf_7723b6938513fa3a byte) byte {
		switch {
		case __obf_7723b6938513fa3a >= '0' && __obf_7723b6938513fa3a <= '9':
			return __obf_7723b6938513fa3a - '0'
		case __obf_7723b6938513fa3a >= 'a' && __obf_7723b6938513fa3a <= 'f':
			return __obf_7723b6938513fa3a - 'a' + 10
		case __obf_7723b6938513fa3a >= 'A' && __obf_7723b6938513fa3a <= 'F':
			return __obf_7723b6938513fa3a - 'A' + 10
		}
		__obf_630bf6ccf46ac8b5 = ColorInvalidErr
		return 0
	}

	switch len(__obf_4e82bdaa91dfc0b2) {
	case 7:
		__obf_8336b398b6a5d0e3.
			R = __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[1])<<4 + __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[2])
		__obf_8336b398b6a5d0e3.
			G = __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[3])<<4 + __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[4])
		__obf_8336b398b6a5d0e3.
			B = __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[5])<<4 + __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[6])

	case 4:
		__obf_8336b398b6a5d0e3.
			R = __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[1]) * 17
		__obf_8336b398b6a5d0e3.
			G = __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[2]) * 17
		__obf_8336b398b6a5d0e3.
			B = __obf_094984e50b675eb7(__obf_4e82bdaa91dfc0b2[3]) * 17
	default:
		__obf_630bf6ccf46ac8b5 = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_5619c94d1f6ee623 string) (bool, error) {
	_, __obf_630bf6ccf46ac8b5 := os.Stat(__obf_5619c94d1f6ee623)
	if __obf_630bf6ccf46ac8b5 == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_630bf6ccf46ac8b5) {
		return false, nil
	}
	return false, __obf_630bf6ccf46ac8b5
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_8602f038322df353 string) bool {
	for _, __obf_3057df8ae9885417 := range __obf_8602f038322df353 {
		if unicode.Is(unicode.Scripts["Han"], __obf_3057df8ae9885417) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_8602f038322df353 string) int {
	return utf8.RuneCountInString(__obf_8602f038322df353)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_0fe328e6db58dab9 []color.Color) color.RGBA {
	__obf_73697d2269cc4bb1 := PickN(__obf_0fe328e6db58dab9, 1)[0]
	__obf_3057df8ae9885417, __obf_6bb01effa05ab1ad, __obf_7723b6938513fa3a, __obf_02d34408cb67890e := __obf_73697d2269cc4bb1.RGBA()
	return color.RGBA{R: uint8(__obf_3057df8ae9885417), G: uint8(__obf_6bb01effa05ab1ad), B: uint8(__obf_7723b6938513fa3a), A: uint8(__obf_02d34408cb67890e)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_b152b4a52fadbc0b int, __obf_02c378c8626b024e int, __obf_e730a61923e268ec image.Image) image.Point {
	__obf_7723b6938513fa3a := __obf_e730a61923e268ec.Bounds()
	__obf_b970e3a1d8f15f94 := __obf_7723b6938513fa3a.Max.X
	__obf_995eff1f0a922179 := __obf_7723b6938513fa3a.Max.Y
	__obf_883c9fb44642e8ce := 0
	__obf_22c4a6e04dec76e8 := 0

	if __obf_b970e3a1d8f15f94-__obf_b152b4a52fadbc0b > 0 {
		__obf_883c9fb44642e8ce = rand.IntN(__obf_b970e3a1d8f15f94 - __obf_b152b4a52fadbc0b)
	}
	if __obf_995eff1f0a922179-__obf_02c378c8626b024e > 0 {
		__obf_22c4a6e04dec76e8 = rand.IntN(__obf_995eff1f0a922179 - __obf_02c378c8626b024e)
	}

	return image.Point{
		X: __obf_883c9fb44642e8ce,
		Y: __obf_22c4a6e04dec76e8,
	}
}

func RandString(__obf_70078f2e712efeb9 []string) string {
	__obf_d72065310913e60c := rand.N(len(__obf_70078f2e712efeb9))
	return __obf_70078f2e712efeb9[__obf_d72065310913e60c]
}

func PickN[T any](__obf_a408d5cdefdd195a []T, __obf_cb87bba412b1f569 int) []T {
	if __obf_cb87bba412b1f569 <= 0 || len(__obf_a408d5cdefdd195a) == 0 {
		return nil
	}
	if __obf_cb87bba412b1f569 >= len(__obf_a408d5cdefdd195a) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_a408d5cdefdd195a...)
	}
	__obf_95b481c0a0be89e5 := // 拷贝一份，避免影响原 slice
		append([]T(nil), __obf_a408d5cdefdd195a...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_95b481c0a0be89e5), func(__obf_d35ae4e1662820d6, __obf_1ad3573e6499f094 int) {
		__obf_95b481c0a0be89e5[__obf_d35ae4e1662820d6], __obf_95b481c0a0be89e5[__obf_1ad3573e6499f094] = __obf_95b481c0a0be89e5[__obf_1ad3573e6499f094], __obf_95b481c0a0be89e5[__obf_d35ae4e1662820d6]
	})

	return __obf_95b481c0a0be89e5[:__obf_cb87bba412b1f569]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
