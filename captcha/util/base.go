package __obf_724f8a3b7b400b23

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
func RotatePoint(__obf_cc2ccdbbc1c378c0, __obf_0166e66ce84143f0, __obf_40109db957c86781, __obf_eb5e0214c10744ab float64) (float64, float64) {
	return __obf_cc2ccdbbc1c378c0*__obf_eb5e0214c10744ab - __obf_0166e66ce84143f0*__obf_40109db957c86781,

		// RotatedSize calculates the size after rotation
		__obf_cc2ccdbbc1c378c0*__obf_40109db957c86781 + __obf_0166e66ce84143f0*__obf_eb5e0214c10744ab
}

func RotatedSize(__obf_e90cbb46d1ef9e24, __obf_1023605e3c797c1b int, __obf_3e52e823b849d817 float64) (int, int) {
	if __obf_e90cbb46d1ef9e24 <= 0 || __obf_1023605e3c797c1b <= 0 {
		return 0, 0
	}
	__obf_40109db957c86781, __obf_eb5e0214c10744ab := math.Sincos(math.Pi * __obf_3e52e823b849d817 / 180)
	__obf_3880ad4ca2b9417a, __obf_292e8112802963da := RotatePoint(float64(__obf_e90cbb46d1ef9e24-1), 0, __obf_40109db957c86781, __obf_eb5e0214c10744ab)
	__obf_6925dc2a261f844d, __obf_05a8c0f2df8df8b2 := RotatePoint(float64(__obf_e90cbb46d1ef9e24-1), float64(__obf_1023605e3c797c1b-1), __obf_40109db957c86781, __obf_eb5e0214c10744ab)
	__obf_ee8e7ef9d6c470b0, __obf_a30530047ab7f825 := RotatePoint(0, float64(__obf_1023605e3c797c1b-1), __obf_40109db957c86781, __obf_eb5e0214c10744ab)
	__obf_2bac9a3ecfc10572 := min(__obf_3880ad4ca2b9417a, min(__obf_6925dc2a261f844d, min(__obf_ee8e7ef9d6c470b0, 0)))
	__obf_5630177e44bc5cc1 := max(__obf_3880ad4ca2b9417a, max(__obf_6925dc2a261f844d, max(__obf_ee8e7ef9d6c470b0, 0)))
	__obf_d4bc549de5756f9c := min(__obf_292e8112802963da, min(__obf_05a8c0f2df8df8b2, min(__obf_a30530047ab7f825, 0)))
	__obf_131c6bdb090f8d89 := max(__obf_292e8112802963da, max(__obf_05a8c0f2df8df8b2, max(__obf_a30530047ab7f825, 0)))
	__obf_1ccbf6d1fd3ae52f := __obf_5630177e44bc5cc1 - __obf_2bac9a3ecfc10572 + 1
	if __obf_1ccbf6d1fd3ae52f-math.Floor(__obf_1ccbf6d1fd3ae52f) > 0.1 {
		__obf_1ccbf6d1fd3ae52f++
	}
	__obf_2a0f692a5f482f64 := __obf_131c6bdb090f8d89 - __obf_d4bc549de5756f9c + 1
	if __obf_2a0f692a5f482f64-math.Floor(__obf_2a0f692a5f482f64) > 0.1 {
		__obf_2a0f692a5f482f64++
	}

	return int(__obf_1ccbf6d1fd3ae52f), int(__obf_2a0f692a5f482f64)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_d64ab16db533403a image.Rectangle, __obf_1ccbf6d1fd3ae52f int, __obf_2a0f692a5f482f64 int, __obf_a8895079967f037a bool) image.Rectangle {
	var __obf_eaa5e765edb30e27 image.Rectangle
	if __obf_1ccbf6d1fd3ae52f*__obf_d64ab16db533403a.Dy() < __obf_2a0f692a5f482f64*__obf_d64ab16db533403a.Dx() {
		__obf_aa5633c527cab361 := float64(__obf_1ccbf6d1fd3ae52f) / float64(__obf_d64ab16db533403a.Dx())
		__obf_8178e6b6db7777fb := int(float64(__obf_d64ab16db533403a.Dy()) * __obf_aa5633c527cab361)
		__obf_de732c080e9abae8 := 0
		if __obf_a8895079967f037a {
			__obf_de732c080e9abae8 = (__obf_2a0f692a5f482f64 - __obf_8178e6b6db7777fb) / 2
		}
		__obf_eaa5e765edb30e27 = image.Rect(0, __obf_de732c080e9abae8, __obf_1ccbf6d1fd3ae52f, __obf_de732c080e9abae8+__obf_8178e6b6db7777fb)
	} else {
		__obf_aa5633c527cab361 := float64(__obf_2a0f692a5f482f64) / float64(__obf_d64ab16db533403a.Dy())
		__obf_772e34b699cfe398 := int(float64(__obf_d64ab16db533403a.Dx()) * __obf_aa5633c527cab361)
		__obf_de732c080e9abae8 := 0
		if __obf_a8895079967f037a {
			__obf_de732c080e9abae8 = (__obf_1ccbf6d1fd3ae52f - __obf_772e34b699cfe398) / 2
		}
		__obf_eaa5e765edb30e27 = image.Rect(__obf_de732c080e9abae8, 0, __obf_de732c080e9abae8+__obf_772e34b699cfe398, __obf_2a0f692a5f482f64)
	}

	return __obf_eaa5e765edb30e27
}

// t2x converts an integer to a hexadecimal string
func __obf_994105b288bf3ece(__obf_2c3c2ddc7367ea85 int64) string {
	__obf_c19b1ca194015564 := strconv.FormatInt(__obf_2c3c2ddc7367ea85, 16)
	if len(__obf_c19b1ca194015564) == 1 {
		__obf_c19b1ca194015564 = "0" + __obf_c19b1ca194015564
	}
	return __obf_c19b1ca194015564
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_901a075c0cf28f0f float32) uint8 {
	__obf_f8b50512a9f18b8e := min(float64(__obf_901a075c0cf28f0f), 1)
	__obf_d5751398e4cf80a4 := __obf_f8b50512a9f18b8e * 255
	return uint8(__obf_d5751398e4cf80a4)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_a5ac0f15072e7e77 int64, __obf_2cb68715e08e525b int64, __obf_c150162d04fe6ee8 int64) string {
	__obf_4f124bf8f4dc3ed5 := __obf_994105b288bf3ece(__obf_a5ac0f15072e7e77)
	__obf_d6ca4ba6f26b40ac := __obf_994105b288bf3ece(__obf_2cb68715e08e525b)
	__obf_a0b806d33e63dadd := __obf_994105b288bf3ece(__obf_c150162d04fe6ee8)
	return __obf_4f124bf8f4dc3ed5 + __obf_d6ca4ba6f26b40ac + __obf_a0b806d33e63dadd
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_38f55d5f1f2e0549 string) (int64, int64, int64) {
	__obf_4f124bf8f4dc3ed5, _ := strconv.ParseInt(__obf_38f55d5f1f2e0549[:2], 16, 10)
	__obf_d6ca4ba6f26b40ac, _ := strconv.ParseInt(__obf_38f55d5f1f2e0549[2:4], 16, 18)
	__obf_a0b806d33e63dadd, _ := strconv.ParseInt(__obf_38f55d5f1f2e0549[4:], 16, 10)
	return __obf_4f124bf8f4dc3ed5, __obf_d6ca4ba6f26b40ac, __obf_a0b806d33e63dadd
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_dbcbb644d707da8b string) (__obf_b7e0c49b86197df8 color.RGBA, __obf_043e1e14636249eb error) {
	__obf_b7e0c49b86197df8.
		A = 0xff
	if __obf_dbcbb644d707da8b[0] != '#' {
		return __obf_b7e0c49b86197df8, ColorHexFormatErr
	}
	__obf_c10aa1a6d83ea9b0 := func(__obf_a0b806d33e63dadd byte) byte {
		switch {
		case __obf_a0b806d33e63dadd >= '0' && __obf_a0b806d33e63dadd <= '9':
			return __obf_a0b806d33e63dadd - '0'
		case __obf_a0b806d33e63dadd >= 'a' && __obf_a0b806d33e63dadd <= 'f':
			return __obf_a0b806d33e63dadd - 'a' + 10
		case __obf_a0b806d33e63dadd >= 'A' && __obf_a0b806d33e63dadd <= 'F':
			return __obf_a0b806d33e63dadd - 'A' + 10
		}
		__obf_043e1e14636249eb = ColorInvalidErr
		return 0
	}

	switch len(__obf_dbcbb644d707da8b) {
	case 7:
		__obf_b7e0c49b86197df8.
			R = __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[1])<<4 + __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[2])
		__obf_b7e0c49b86197df8.
			G = __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[3])<<4 + __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[4])
		__obf_b7e0c49b86197df8.
			B = __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[5])<<4 + __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[6])

	case 4:
		__obf_b7e0c49b86197df8.
			R = __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[1]) * 17
		__obf_b7e0c49b86197df8.
			G = __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[2]) * 17
		__obf_b7e0c49b86197df8.
			B = __obf_c10aa1a6d83ea9b0(__obf_dbcbb644d707da8b[3]) * 17
	default:
		__obf_043e1e14636249eb = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_8457ef135ea46003 string) (bool, error) {
	_, __obf_043e1e14636249eb := os.Stat(__obf_8457ef135ea46003)
	if __obf_043e1e14636249eb == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_043e1e14636249eb) {
		return false, nil
	}
	return false, __obf_043e1e14636249eb
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_309de8ada4c37842 string) bool {
	for _, __obf_4f124bf8f4dc3ed5 := range __obf_309de8ada4c37842 {
		if unicode.Is(unicode.Scripts["Han"], __obf_4f124bf8f4dc3ed5) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_309de8ada4c37842 string) int {
	return utf8.RuneCountInString(__obf_309de8ada4c37842)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_90a738feebd8551f []color.Color) color.RGBA {
	__obf_962190e4cbbf18f1 := PickN(__obf_90a738feebd8551f, 1)[0]
	__obf_4f124bf8f4dc3ed5, __obf_d6ca4ba6f26b40ac, __obf_a0b806d33e63dadd, __obf_f8b50512a9f18b8e := __obf_962190e4cbbf18f1.RGBA()
	return color.RGBA{R: uint8(__obf_4f124bf8f4dc3ed5), G: uint8(__obf_d6ca4ba6f26b40ac), B: uint8(__obf_a0b806d33e63dadd), A: uint8(__obf_f8b50512a9f18b8e)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_1ccbf6d1fd3ae52f int, __obf_2a0f692a5f482f64 int, __obf_4aeb51bb135c8905 image.Image) image.Point {
	__obf_a0b806d33e63dadd := __obf_4aeb51bb135c8905.Bounds()
	__obf_dcd1352f8aa92673 := __obf_a0b806d33e63dadd.Max.X
	__obf_76e395a97358be28 := __obf_a0b806d33e63dadd.Max.Y
	__obf_479bb12b29250d57 := 0
	__obf_979c3165d66dc7a5 := 0

	if __obf_dcd1352f8aa92673-__obf_1ccbf6d1fd3ae52f > 0 {
		__obf_479bb12b29250d57 = rand.IntN(__obf_dcd1352f8aa92673 - __obf_1ccbf6d1fd3ae52f)
	}
	if __obf_76e395a97358be28-__obf_2a0f692a5f482f64 > 0 {
		__obf_979c3165d66dc7a5 = rand.IntN(__obf_76e395a97358be28 - __obf_2a0f692a5f482f64)
	}

	return image.Point{
		X: __obf_479bb12b29250d57,
		Y: __obf_979c3165d66dc7a5,
	}
}

func RandString(__obf_28bf12fc72b10b84 []string) string {
	__obf_a073b205248e6e3a := rand.N(len(__obf_28bf12fc72b10b84))
	return __obf_28bf12fc72b10b84[__obf_a073b205248e6e3a]
}

func PickN[T any](__obf_5e328dbf9eb2baf5 []T, __obf_7540da85d53dc887 int) []T {
	if __obf_7540da85d53dc887 <= 0 || len(__obf_5e328dbf9eb2baf5) == 0 {
		return nil
	}
	if __obf_7540da85d53dc887 >= len(__obf_5e328dbf9eb2baf5) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_5e328dbf9eb2baf5...)
	}
	__obf_2bd2b7b11310a3bf := // 拷贝一份，避免影响原 slice
		append([]T(nil), __obf_5e328dbf9eb2baf5...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_2bd2b7b11310a3bf), func(__obf_5e663421ea0d9845, __obf_7fab1c72276f1db1 int) {
		__obf_2bd2b7b11310a3bf[__obf_5e663421ea0d9845], __obf_2bd2b7b11310a3bf[__obf_7fab1c72276f1db1] = __obf_2bd2b7b11310a3bf[__obf_7fab1c72276f1db1], __obf_2bd2b7b11310a3bf[__obf_5e663421ea0d9845]
	})

	return __obf_2bd2b7b11310a3bf[:__obf_7540da85d53dc887]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
