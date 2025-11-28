package __obf_3a8eca76d39e021a

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
func RotatePoint(__obf_fb4ac65deb0abc24, __obf_43a5722247b5ac53, __obf_f397220a20088da5, __obf_e3ae7ca33641d111 float64) (float64, float64) {
	return __obf_fb4ac65deb0abc24*__obf_e3ae7ca33641d111 - __obf_43a5722247b5ac53*__obf_f397220a20088da5, __obf_fb4ac65deb0abc24*__obf_f397220a20088da5 + __obf_43a5722247b5ac53*__obf_e3ae7ca33641d111
}

// RotatedSize calculates the size after rotation
func RotatedSize(__obf_be6e8cae5b2c8541, __obf_495de6332fa3d4ec int, __obf_8043b5fbf7511810 float64) (int, int) {
	if __obf_be6e8cae5b2c8541 <= 0 || __obf_495de6332fa3d4ec <= 0 {
		return 0, 0
	}

	__obf_f397220a20088da5, __obf_e3ae7ca33641d111 := math.Sincos(math.Pi * __obf_8043b5fbf7511810 / 180)
	__obf_d62ecba1a82166a7, __obf_be8566bab050b773 := RotatePoint(float64(__obf_be6e8cae5b2c8541-1), 0, __obf_f397220a20088da5, __obf_e3ae7ca33641d111)
	__obf_6631499f243c6c07, __obf_66e9111c423e8fea := RotatePoint(float64(__obf_be6e8cae5b2c8541-1), float64(__obf_495de6332fa3d4ec-1), __obf_f397220a20088da5, __obf_e3ae7ca33641d111)
	__obf_bc1a6a85e91c30a6, __obf_85c83ce91e2732d2 := RotatePoint(0, float64(__obf_495de6332fa3d4ec-1), __obf_f397220a20088da5, __obf_e3ae7ca33641d111)

	__obf_eff6d7eb0ca772c3 := min(__obf_d62ecba1a82166a7, min(__obf_6631499f243c6c07, min(__obf_bc1a6a85e91c30a6, 0)))
	__obf_06b8ccafdf059dd8 := max(__obf_d62ecba1a82166a7, max(__obf_6631499f243c6c07, max(__obf_bc1a6a85e91c30a6, 0)))
	__obf_3892c7b2d6d1ae27 := min(__obf_be8566bab050b773, min(__obf_66e9111c423e8fea, min(__obf_85c83ce91e2732d2, 0)))
	__obf_223bfe6535e34aa2 := max(__obf_be8566bab050b773, max(__obf_66e9111c423e8fea, max(__obf_85c83ce91e2732d2, 0)))

	__obf_42da9ec45bf326d2 := __obf_06b8ccafdf059dd8 - __obf_eff6d7eb0ca772c3 + 1
	if __obf_42da9ec45bf326d2-math.Floor(__obf_42da9ec45bf326d2) > 0.1 {
		__obf_42da9ec45bf326d2++
	}
	__obf_cead67a459d05a97 := __obf_223bfe6535e34aa2 - __obf_3892c7b2d6d1ae27 + 1
	if __obf_cead67a459d05a97-math.Floor(__obf_cead67a459d05a97) > 0.1 {
		__obf_cead67a459d05a97++
	}

	return int(__obf_42da9ec45bf326d2), int(__obf_cead67a459d05a97)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_6ccf4baf567f8d0c image.Rectangle, __obf_42da9ec45bf326d2 int, __obf_cead67a459d05a97 int, __obf_5e022e832b2a4fb9 bool) image.Rectangle {
	var __obf_d192785a25411fc7 image.Rectangle
	if __obf_42da9ec45bf326d2*__obf_6ccf4baf567f8d0c.Dy() < __obf_cead67a459d05a97*__obf_6ccf4baf567f8d0c.Dx() {
		__obf_b448b899a9e2d5ce := float64(__obf_42da9ec45bf326d2) / float64(__obf_6ccf4baf567f8d0c.Dx())

		__obf_bb39def901809108 := int(float64(__obf_6ccf4baf567f8d0c.Dy()) * __obf_b448b899a9e2d5ce)
		__obf_444b964bdb2611d3 := 0
		if __obf_5e022e832b2a4fb9 {
			__obf_444b964bdb2611d3 = (__obf_cead67a459d05a97 - __obf_bb39def901809108) / 2
		}
		__obf_d192785a25411fc7 = image.Rect(0, __obf_444b964bdb2611d3, __obf_42da9ec45bf326d2, __obf_444b964bdb2611d3+__obf_bb39def901809108)
	} else {
		__obf_b448b899a9e2d5ce := float64(__obf_cead67a459d05a97) / float64(__obf_6ccf4baf567f8d0c.Dy())
		__obf_236a895c1dd25135 := int(float64(__obf_6ccf4baf567f8d0c.Dx()) * __obf_b448b899a9e2d5ce)
		__obf_444b964bdb2611d3 := 0
		if __obf_5e022e832b2a4fb9 {
			__obf_444b964bdb2611d3 = (__obf_42da9ec45bf326d2 - __obf_236a895c1dd25135) / 2
		}
		__obf_d192785a25411fc7 = image.Rect(__obf_444b964bdb2611d3, 0, __obf_444b964bdb2611d3+__obf_236a895c1dd25135, __obf_cead67a459d05a97)
	}

	return __obf_d192785a25411fc7
}

// t2x converts an integer to a hexadecimal string
func __obf_2c18d437357fe8b8(__obf_9c33f387697edfdc int64) string {
	__obf_27d732f068a82cb0 := strconv.FormatInt(__obf_9c33f387697edfdc, 16)
	if len(__obf_27d732f068a82cb0) == 1 {
		__obf_27d732f068a82cb0 = "0" + __obf_27d732f068a82cb0
	}
	return __obf_27d732f068a82cb0
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_050ece9bb28de5de float32) uint8 {
	__obf_005380c156fcec4f := min(float64(__obf_050ece9bb28de5de), 1)
	__obf_732c451c03a6d6db := __obf_005380c156fcec4f * 255
	return uint8(__obf_732c451c03a6d6db)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_e1b97aec89e9161a int64, __obf_2edfc7319cc59d85 int64, __obf_7712ca48122b6239 int64) string {
	__obf_9a9b1086ea6b6251 := __obf_2c18d437357fe8b8(__obf_e1b97aec89e9161a)
	__obf_e883c0d6f897aa60 := __obf_2c18d437357fe8b8(__obf_2edfc7319cc59d85)
	__obf_c689f43cdba733b4 := __obf_2c18d437357fe8b8(__obf_7712ca48122b6239)
	return __obf_9a9b1086ea6b6251 + __obf_e883c0d6f897aa60 + __obf_c689f43cdba733b4
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_c1a682c4746988d8 string) (int64, int64, int64) {
	__obf_9a9b1086ea6b6251, _ := strconv.ParseInt(__obf_c1a682c4746988d8[:2], 16, 10)
	__obf_e883c0d6f897aa60, _ := strconv.ParseInt(__obf_c1a682c4746988d8[2:4], 16, 18)
	__obf_c689f43cdba733b4, _ := strconv.ParseInt(__obf_c1a682c4746988d8[4:], 16, 10)
	return __obf_9a9b1086ea6b6251, __obf_e883c0d6f897aa60, __obf_c689f43cdba733b4
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_8706facc950fcb8e string) (__obf_d5f392407dca81db color.RGBA, __obf_d188595df3b06d83 error) {
	__obf_d5f392407dca81db.A = 0xff
	if __obf_8706facc950fcb8e[0] != '#' {
		return __obf_d5f392407dca81db, ColorHexFormatErr
	}

	__obf_231e5d468b478420 := func(__obf_c689f43cdba733b4 byte) byte {
		switch {
		case __obf_c689f43cdba733b4 >= '0' && __obf_c689f43cdba733b4 <= '9':
			return __obf_c689f43cdba733b4 - '0'
		case __obf_c689f43cdba733b4 >= 'a' && __obf_c689f43cdba733b4 <= 'f':
			return __obf_c689f43cdba733b4 - 'a' + 10
		case __obf_c689f43cdba733b4 >= 'A' && __obf_c689f43cdba733b4 <= 'F':
			return __obf_c689f43cdba733b4 - 'A' + 10
		}
		__obf_d188595df3b06d83 = ColorInvalidErr
		return 0
	}

	switch len(__obf_8706facc950fcb8e) {
	case 7:
		__obf_d5f392407dca81db.R = __obf_231e5d468b478420(__obf_8706facc950fcb8e[1])<<4 + __obf_231e5d468b478420(__obf_8706facc950fcb8e[2])
		__obf_d5f392407dca81db.G = __obf_231e5d468b478420(__obf_8706facc950fcb8e[3])<<4 + __obf_231e5d468b478420(__obf_8706facc950fcb8e[4])
		__obf_d5f392407dca81db.B = __obf_231e5d468b478420(__obf_8706facc950fcb8e[5])<<4 + __obf_231e5d468b478420(__obf_8706facc950fcb8e[6])

	case 4:
		__obf_d5f392407dca81db.R = __obf_231e5d468b478420(__obf_8706facc950fcb8e[1]) * 17
		__obf_d5f392407dca81db.G = __obf_231e5d468b478420(__obf_8706facc950fcb8e[2]) * 17
		__obf_d5f392407dca81db.B = __obf_231e5d468b478420(__obf_8706facc950fcb8e[3]) * 17
	default:
		__obf_d188595df3b06d83 = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_8d7955fde3834427 string) (bool, error) {
	_, __obf_d188595df3b06d83 := os.Stat(__obf_8d7955fde3834427)
	if __obf_d188595df3b06d83 == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_d188595df3b06d83) {
		return false, nil
	}
	return false, __obf_d188595df3b06d83
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_ca356c0954fe1007 string) bool {
	for _, __obf_9a9b1086ea6b6251 := range __obf_ca356c0954fe1007 {
		if unicode.Is(unicode.Scripts["Han"], __obf_9a9b1086ea6b6251) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_ca356c0954fe1007 string) int {
	return utf8.RuneCountInString(__obf_ca356c0954fe1007)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_3636d46a138c309e []color.Color) color.RGBA {
	__obf_1e33f1ffe3cb7e49 := PickN(__obf_3636d46a138c309e, 1)[0]
	__obf_9a9b1086ea6b6251, __obf_e883c0d6f897aa60, __obf_c689f43cdba733b4, __obf_005380c156fcec4f := __obf_1e33f1ffe3cb7e49.RGBA()
	return color.RGBA{R: uint8(__obf_9a9b1086ea6b6251), G: uint8(__obf_e883c0d6f897aa60), B: uint8(__obf_c689f43cdba733b4), A: uint8(__obf_005380c156fcec4f)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_42da9ec45bf326d2 int, __obf_cead67a459d05a97 int, __obf_d5358d2a73ea01f6 image.Image) image.Point {
	__obf_c689f43cdba733b4 := __obf_d5358d2a73ea01f6.Bounds()
	__obf_8650134c7c8b0d90 := __obf_c689f43cdba733b4.Max.X
	__obf_e5ffec76501c9a17 := __obf_c689f43cdba733b4.Max.Y
	__obf_26688f105e700728 := 0
	__obf_b388e42b1264a0ec := 0

	if __obf_8650134c7c8b0d90-__obf_42da9ec45bf326d2 > 0 {
		__obf_26688f105e700728 = rand.IntN(__obf_8650134c7c8b0d90 - __obf_42da9ec45bf326d2)
	}
	if __obf_e5ffec76501c9a17-__obf_cead67a459d05a97 > 0 {
		__obf_b388e42b1264a0ec = rand.IntN(__obf_e5ffec76501c9a17 - __obf_cead67a459d05a97)
	}

	return image.Point{
		X: __obf_26688f105e700728,
		Y: __obf_b388e42b1264a0ec,
	}
}

func RandString(__obf_8dddc52227bdce64 []string) string {
	__obf_5e66de797575e20d := rand.N(len(__obf_8dddc52227bdce64))
	return __obf_8dddc52227bdce64[__obf_5e66de797575e20d]
}

func PickN[T any](__obf_bab9f85e36b0ec02 []T, __obf_79aa29cd34bd57e9 int) []T {
	if __obf_79aa29cd34bd57e9 <= 0 || len(__obf_bab9f85e36b0ec02) == 0 {
		return nil
	}
	if __obf_79aa29cd34bd57e9 >= len(__obf_bab9f85e36b0ec02) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_bab9f85e36b0ec02...)
	}

	// 拷贝一份，避免影响原 slice
	__obf_c26879063a69f6c7 := append([]T(nil), __obf_bab9f85e36b0ec02...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_c26879063a69f6c7), func(__obf_bca8280e85c5f5d7, __obf_a53455ab6269ea40 int) {
		__obf_c26879063a69f6c7[__obf_bca8280e85c5f5d7], __obf_c26879063a69f6c7[__obf_a53455ab6269ea40] = __obf_c26879063a69f6c7[__obf_a53455ab6269ea40], __obf_c26879063a69f6c7[__obf_bca8280e85c5f5d7]
	})

	return __obf_c26879063a69f6c7[:__obf_79aa29cd34bd57e9]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
