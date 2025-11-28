package __obf_44ec487f80286103

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
func RotatePoint(__obf_71b50dc7165aeef7, __obf_1fbb9386b3b19015, __obf_678953a9e5d671a8, __obf_1aa33f1bd19908a2 float64) (float64, float64) {
	return __obf_71b50dc7165aeef7*__obf_1aa33f1bd19908a2 - __obf_1fbb9386b3b19015*__obf_678953a9e5d671a8, __obf_71b50dc7165aeef7*__obf_678953a9e5d671a8 + __obf_1fbb9386b3b19015*__obf_1aa33f1bd19908a2
}

// RotatedSize calculates the size after rotation
func RotatedSize(__obf_bb2d168d5baaf841, __obf_579cb645c388f1e1 int, __obf_e0d34617afaa7ea3 float64) (int, int) {
	if __obf_bb2d168d5baaf841 <= 0 || __obf_579cb645c388f1e1 <= 0 {
		return 0, 0
	}

	__obf_678953a9e5d671a8, __obf_1aa33f1bd19908a2 := math.Sincos(math.Pi * __obf_e0d34617afaa7ea3 / 180)
	__obf_2b3ab14b2ed9f57d, __obf_b63a12fef6c0cd41 := RotatePoint(float64(__obf_bb2d168d5baaf841-1), 0, __obf_678953a9e5d671a8, __obf_1aa33f1bd19908a2)
	__obf_068c9dac91d825f9, __obf_ebaf5d62fb506598 := RotatePoint(float64(__obf_bb2d168d5baaf841-1), float64(__obf_579cb645c388f1e1-1), __obf_678953a9e5d671a8, __obf_1aa33f1bd19908a2)
	__obf_3b6d53e9bdf97178, __obf_89dc62c5077bdc8f := RotatePoint(0, float64(__obf_579cb645c388f1e1-1), __obf_678953a9e5d671a8, __obf_1aa33f1bd19908a2)

	__obf_7ea12a2ea92b8484 := min(__obf_2b3ab14b2ed9f57d, min(__obf_068c9dac91d825f9, min(__obf_3b6d53e9bdf97178, 0)))
	__obf_0255c5404d1f7beb := max(__obf_2b3ab14b2ed9f57d, max(__obf_068c9dac91d825f9, max(__obf_3b6d53e9bdf97178, 0)))
	__obf_0961eb35d5dfd893 := min(__obf_b63a12fef6c0cd41, min(__obf_ebaf5d62fb506598, min(__obf_89dc62c5077bdc8f, 0)))
	__obf_6477abbe780198bd := max(__obf_b63a12fef6c0cd41, max(__obf_ebaf5d62fb506598, max(__obf_89dc62c5077bdc8f, 0)))

	__obf_f0e9d3ab8f09178c := __obf_0255c5404d1f7beb - __obf_7ea12a2ea92b8484 + 1
	if __obf_f0e9d3ab8f09178c-math.Floor(__obf_f0e9d3ab8f09178c) > 0.1 {
		__obf_f0e9d3ab8f09178c++
	}
	__obf_9358f08fbbfdcbde := __obf_6477abbe780198bd - __obf_0961eb35d5dfd893 + 1
	if __obf_9358f08fbbfdcbde-math.Floor(__obf_9358f08fbbfdcbde) > 0.1 {
		__obf_9358f08fbbfdcbde++
	}

	return int(__obf_f0e9d3ab8f09178c), int(__obf_9358f08fbbfdcbde)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_2e62c1e9f0727ee9 image.Rectangle, __obf_f0e9d3ab8f09178c int, __obf_9358f08fbbfdcbde int, __obf_039ed7d4f5b039cd bool) image.Rectangle {
	var __obf_e1420155fc6e8fe7 image.Rectangle
	if __obf_f0e9d3ab8f09178c*__obf_2e62c1e9f0727ee9.Dy() < __obf_9358f08fbbfdcbde*__obf_2e62c1e9f0727ee9.Dx() {
		__obf_d442ebfd73d403b6 := float64(__obf_f0e9d3ab8f09178c) / float64(__obf_2e62c1e9f0727ee9.Dx())

		__obf_ed211a6cf648ec15 := int(float64(__obf_2e62c1e9f0727ee9.Dy()) * __obf_d442ebfd73d403b6)
		__obf_005e92c0c6264dfe := 0
		if __obf_039ed7d4f5b039cd {
			__obf_005e92c0c6264dfe = (__obf_9358f08fbbfdcbde - __obf_ed211a6cf648ec15) / 2
		}
		__obf_e1420155fc6e8fe7 = image.Rect(0, __obf_005e92c0c6264dfe, __obf_f0e9d3ab8f09178c, __obf_005e92c0c6264dfe+__obf_ed211a6cf648ec15)
	} else {
		__obf_d442ebfd73d403b6 := float64(__obf_9358f08fbbfdcbde) / float64(__obf_2e62c1e9f0727ee9.Dy())
		__obf_827ae6efa58dee16 := int(float64(__obf_2e62c1e9f0727ee9.Dx()) * __obf_d442ebfd73d403b6)
		__obf_005e92c0c6264dfe := 0
		if __obf_039ed7d4f5b039cd {
			__obf_005e92c0c6264dfe = (__obf_f0e9d3ab8f09178c - __obf_827ae6efa58dee16) / 2
		}
		__obf_e1420155fc6e8fe7 = image.Rect(__obf_005e92c0c6264dfe, 0, __obf_005e92c0c6264dfe+__obf_827ae6efa58dee16, __obf_9358f08fbbfdcbde)
	}

	return __obf_e1420155fc6e8fe7
}

// t2x converts an integer to a hexadecimal string
func __obf_8cc20242778f0329(__obf_27c3efacfa6993bb int64) string {
	__obf_4614d4c1b5851516 := strconv.FormatInt(__obf_27c3efacfa6993bb, 16)
	if len(__obf_4614d4c1b5851516) == 1 {
		__obf_4614d4c1b5851516 = "0" + __obf_4614d4c1b5851516
	}
	return __obf_4614d4c1b5851516
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_9b2d203d855ddba0 float32) uint8 {
	__obf_3d088e1f71bd34c5 := min(float64(__obf_9b2d203d855ddba0), 1)
	__obf_4e09a2fd5d46392b := __obf_3d088e1f71bd34c5 * 255
	return uint8(__obf_4e09a2fd5d46392b)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_ab593b9181be8eed int64, __obf_ed8af04fa4240065 int64, __obf_abbabc113743e334 int64) string {
	__obf_955452a21a98fbde := __obf_8cc20242778f0329(__obf_ab593b9181be8eed)
	__obf_f5731989548bb78d := __obf_8cc20242778f0329(__obf_ed8af04fa4240065)
	__obf_7571c697f623cc50 := __obf_8cc20242778f0329(__obf_abbabc113743e334)
	return __obf_955452a21a98fbde + __obf_f5731989548bb78d + __obf_7571c697f623cc50
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_8ec12996ce30b731 string) (int64, int64, int64) {
	__obf_955452a21a98fbde, _ := strconv.ParseInt(__obf_8ec12996ce30b731[:2], 16, 10)
	__obf_f5731989548bb78d, _ := strconv.ParseInt(__obf_8ec12996ce30b731[2:4], 16, 18)
	__obf_7571c697f623cc50, _ := strconv.ParseInt(__obf_8ec12996ce30b731[4:], 16, 10)
	return __obf_955452a21a98fbde, __obf_f5731989548bb78d, __obf_7571c697f623cc50
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_805f291ff4031cd3 string) (__obf_f0046e0a9a7ebae3 color.RGBA, __obf_9ad8a31ce70de66b error) {
	__obf_f0046e0a9a7ebae3.A = 0xff
	if __obf_805f291ff4031cd3[0] != '#' {
		return __obf_f0046e0a9a7ebae3, ColorHexFormatErr
	}

	__obf_b6c7181a7ecb905e := func(__obf_7571c697f623cc50 byte) byte {
		switch {
		case __obf_7571c697f623cc50 >= '0' && __obf_7571c697f623cc50 <= '9':
			return __obf_7571c697f623cc50 - '0'
		case __obf_7571c697f623cc50 >= 'a' && __obf_7571c697f623cc50 <= 'f':
			return __obf_7571c697f623cc50 - 'a' + 10
		case __obf_7571c697f623cc50 >= 'A' && __obf_7571c697f623cc50 <= 'F':
			return __obf_7571c697f623cc50 - 'A' + 10
		}
		__obf_9ad8a31ce70de66b = ColorInvalidErr
		return 0
	}

	switch len(__obf_805f291ff4031cd3) {
	case 7:
		__obf_f0046e0a9a7ebae3.R = __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[1])<<4 + __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[2])
		__obf_f0046e0a9a7ebae3.G = __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[3])<<4 + __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[4])
		__obf_f0046e0a9a7ebae3.B = __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[5])<<4 + __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[6])

	case 4:
		__obf_f0046e0a9a7ebae3.R = __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[1]) * 17
		__obf_f0046e0a9a7ebae3.G = __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[2]) * 17
		__obf_f0046e0a9a7ebae3.B = __obf_b6c7181a7ecb905e(__obf_805f291ff4031cd3[3]) * 17
	default:
		__obf_9ad8a31ce70de66b = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_742bfbfd48e11cab string) (bool, error) {
	_, __obf_9ad8a31ce70de66b := os.Stat(__obf_742bfbfd48e11cab)
	if __obf_9ad8a31ce70de66b == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_9ad8a31ce70de66b) {
		return false, nil
	}
	return false, __obf_9ad8a31ce70de66b
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_53513a471cd435e3 string) bool {
	for _, __obf_955452a21a98fbde := range __obf_53513a471cd435e3 {
		if unicode.Is(unicode.Scripts["Han"], __obf_955452a21a98fbde) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_53513a471cd435e3 string) int {
	return utf8.RuneCountInString(__obf_53513a471cd435e3)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_7db7629a9bff103c []color.Color) color.RGBA {
	__obf_46f7a01b9fd32467 := PickN(__obf_7db7629a9bff103c, 1)[0]
	__obf_955452a21a98fbde, __obf_f5731989548bb78d, __obf_7571c697f623cc50, __obf_3d088e1f71bd34c5 := __obf_46f7a01b9fd32467.RGBA()
	return color.RGBA{R: uint8(__obf_955452a21a98fbde), G: uint8(__obf_f5731989548bb78d), B: uint8(__obf_7571c697f623cc50), A: uint8(__obf_3d088e1f71bd34c5)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_f0e9d3ab8f09178c int, __obf_9358f08fbbfdcbde int, __obf_1392ec6be47e851e image.Image) image.Point {
	__obf_7571c697f623cc50 := __obf_1392ec6be47e851e.Bounds()
	__obf_9e489be1b02937c1 := __obf_7571c697f623cc50.Max.X
	__obf_668f3d00f0ed0f24 := __obf_7571c697f623cc50.Max.Y
	__obf_0802e4d345c4cfcf := 0
	__obf_ceeadb001faf1e5c := 0

	if __obf_9e489be1b02937c1-__obf_f0e9d3ab8f09178c > 0 {
		__obf_0802e4d345c4cfcf = rand.IntN(__obf_9e489be1b02937c1 - __obf_f0e9d3ab8f09178c)
	}
	if __obf_668f3d00f0ed0f24-__obf_9358f08fbbfdcbde > 0 {
		__obf_ceeadb001faf1e5c = rand.IntN(__obf_668f3d00f0ed0f24 - __obf_9358f08fbbfdcbde)
	}

	return image.Point{
		X: __obf_0802e4d345c4cfcf,
		Y: __obf_ceeadb001faf1e5c,
	}
}

func RandString(__obf_a0c84fecf91ae419 []string) string {
	__obf_ecb3654302d3cf01 := rand.N(len(__obf_a0c84fecf91ae419))
	return __obf_a0c84fecf91ae419[__obf_ecb3654302d3cf01]
}

func PickN[T any](__obf_1a314b60fff2f072 []T, __obf_81122c450fe7ae3c int) []T {
	if __obf_81122c450fe7ae3c <= 0 || len(__obf_1a314b60fff2f072) == 0 {
		return nil
	}
	if __obf_81122c450fe7ae3c >= len(__obf_1a314b60fff2f072) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_1a314b60fff2f072...)
	}

	// 拷贝一份，避免影响原 slice
	__obf_da327e1280ee443b := append([]T(nil), __obf_1a314b60fff2f072...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_da327e1280ee443b), func(__obf_bc2f225dbb10e874, __obf_aa9c40ba1eb3052b int) {
		__obf_da327e1280ee443b[__obf_bc2f225dbb10e874], __obf_da327e1280ee443b[__obf_aa9c40ba1eb3052b] = __obf_da327e1280ee443b[__obf_aa9c40ba1eb3052b], __obf_da327e1280ee443b[__obf_bc2f225dbb10e874]
	})

	return __obf_da327e1280ee443b[:__obf_81122c450fe7ae3c]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
