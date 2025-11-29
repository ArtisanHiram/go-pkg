package __obf_3860403e44e18318

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
func RotatePoint(__obf_ad550545eb5e3492, __obf_aaa70bd6c198f22a, __obf_07e36924cf5fbe65, __obf_dcec6a3343cdd4fa float64) (float64, float64) {
	return __obf_ad550545eb5e3492*__obf_dcec6a3343cdd4fa - __obf_aaa70bd6c198f22a*__obf_07e36924cf5fbe65,

		// RotatedSize calculates the size after rotation
		__obf_ad550545eb5e3492*__obf_07e36924cf5fbe65 + __obf_aaa70bd6c198f22a*__obf_dcec6a3343cdd4fa
}

func RotatedSize(__obf_a285d9c6db7a6da4, __obf_a56fe594d81995d7 int, __obf_cd98204ea33509ad float64) (int, int) {
	if __obf_a285d9c6db7a6da4 <= 0 || __obf_a56fe594d81995d7 <= 0 {
		return 0, 0
	}
	__obf_07e36924cf5fbe65, __obf_dcec6a3343cdd4fa := math.Sincos(math.Pi * __obf_cd98204ea33509ad / 180)
	__obf_16b55da1edd70743, __obf_818ed87707ea7cf8 := RotatePoint(float64(__obf_a285d9c6db7a6da4-1), 0, __obf_07e36924cf5fbe65, __obf_dcec6a3343cdd4fa)
	__obf_8942f8cd2b8a1d28, __obf_bd593e928cd4f120 := RotatePoint(float64(__obf_a285d9c6db7a6da4-1), float64(__obf_a56fe594d81995d7-1), __obf_07e36924cf5fbe65, __obf_dcec6a3343cdd4fa)
	__obf_50c2e75c68eedef2, __obf_49dea2fe32387661 := RotatePoint(0, float64(__obf_a56fe594d81995d7-1), __obf_07e36924cf5fbe65, __obf_dcec6a3343cdd4fa)
	__obf_50e0a08a80eed377 := min(__obf_16b55da1edd70743, min(__obf_8942f8cd2b8a1d28, min(__obf_50c2e75c68eedef2, 0)))
	__obf_e4498efde7ac644e := max(__obf_16b55da1edd70743, max(__obf_8942f8cd2b8a1d28, max(__obf_50c2e75c68eedef2, 0)))
	__obf_6962b9622d75675c := min(__obf_818ed87707ea7cf8, min(__obf_bd593e928cd4f120, min(__obf_49dea2fe32387661, 0)))
	__obf_995e2f98dc5e0cdb := max(__obf_818ed87707ea7cf8, max(__obf_bd593e928cd4f120, max(__obf_49dea2fe32387661, 0)))
	__obf_9b0052aae736215a := __obf_e4498efde7ac644e - __obf_50e0a08a80eed377 + 1
	if __obf_9b0052aae736215a-math.Floor(__obf_9b0052aae736215a) > 0.1 {
		__obf_9b0052aae736215a++
	}
	__obf_f857aea34adf2b31 := __obf_995e2f98dc5e0cdb - __obf_6962b9622d75675c + 1
	if __obf_f857aea34adf2b31-math.Floor(__obf_f857aea34adf2b31) > 0.1 {
		__obf_f857aea34adf2b31++
	}

	return int(__obf_9b0052aae736215a), int(__obf_f857aea34adf2b31)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_3166d2a36ec273c9 image.Rectangle, __obf_9b0052aae736215a int, __obf_f857aea34adf2b31 int, __obf_7bce4b5fb83b2de6 bool) image.Rectangle {
	var __obf_d80306e25cc9472b image.Rectangle
	if __obf_9b0052aae736215a*__obf_3166d2a36ec273c9.Dy() < __obf_f857aea34adf2b31*__obf_3166d2a36ec273c9.Dx() {
		__obf_3598b3e9a331b37a := float64(__obf_9b0052aae736215a) / float64(__obf_3166d2a36ec273c9.Dx())
		__obf_440940d65fc9fda4 := int(float64(__obf_3166d2a36ec273c9.Dy()) * __obf_3598b3e9a331b37a)
		__obf_fc00d1e342a1ac98 := 0
		if __obf_7bce4b5fb83b2de6 {
			__obf_fc00d1e342a1ac98 = (__obf_f857aea34adf2b31 - __obf_440940d65fc9fda4) / 2
		}
		__obf_d80306e25cc9472b = image.Rect(0, __obf_fc00d1e342a1ac98, __obf_9b0052aae736215a, __obf_fc00d1e342a1ac98+__obf_440940d65fc9fda4)
	} else {
		__obf_3598b3e9a331b37a := float64(__obf_f857aea34adf2b31) / float64(__obf_3166d2a36ec273c9.Dy())
		__obf_1797961fb2a4cd9b := int(float64(__obf_3166d2a36ec273c9.Dx()) * __obf_3598b3e9a331b37a)
		__obf_fc00d1e342a1ac98 := 0
		if __obf_7bce4b5fb83b2de6 {
			__obf_fc00d1e342a1ac98 = (__obf_9b0052aae736215a - __obf_1797961fb2a4cd9b) / 2
		}
		__obf_d80306e25cc9472b = image.Rect(__obf_fc00d1e342a1ac98, 0, __obf_fc00d1e342a1ac98+__obf_1797961fb2a4cd9b, __obf_f857aea34adf2b31)
	}

	return __obf_d80306e25cc9472b
}

// t2x converts an integer to a hexadecimal string
func __obf_5114fbe63f281a88(__obf_b52da4bcf1d25e87 int64) string {
	__obf_af9c74186ab19d60 := strconv.FormatInt(__obf_b52da4bcf1d25e87, 16)
	if len(__obf_af9c74186ab19d60) == 1 {
		__obf_af9c74186ab19d60 = "0" + __obf_af9c74186ab19d60
	}
	return __obf_af9c74186ab19d60
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_7cd816d005df8ffe float32) uint8 {
	__obf_2cfffe814d05be73 := min(float64(__obf_7cd816d005df8ffe), 1)
	__obf_8848ccfc7aa82c4d := __obf_2cfffe814d05be73 * 255
	return uint8(__obf_8848ccfc7aa82c4d)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_2061b72d263e5dc6 int64, __obf_560cafdc16698210 int64, __obf_e85c1a722e28507a int64) string {
	__obf_3b0471d953e55548 := __obf_5114fbe63f281a88(__obf_2061b72d263e5dc6)
	__obf_c09dade8b2b5ab0f := __obf_5114fbe63f281a88(__obf_560cafdc16698210)
	__obf_3536db1209733836 := __obf_5114fbe63f281a88(__obf_e85c1a722e28507a)
	return __obf_3b0471d953e55548 + __obf_c09dade8b2b5ab0f + __obf_3536db1209733836
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_61e7fd6d844c52bc string) (int64, int64, int64) {
	__obf_3b0471d953e55548, _ := strconv.ParseInt(__obf_61e7fd6d844c52bc[:2], 16, 10)
	__obf_c09dade8b2b5ab0f, _ := strconv.ParseInt(__obf_61e7fd6d844c52bc[2:4], 16, 18)
	__obf_3536db1209733836, _ := strconv.ParseInt(__obf_61e7fd6d844c52bc[4:], 16, 10)
	return __obf_3b0471d953e55548, __obf_c09dade8b2b5ab0f, __obf_3536db1209733836
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_041320c70a794a0c string) (__obf_0c3c354bd419abd5 color.RGBA, __obf_7834a057dbd32804 error) {
	__obf_0c3c354bd419abd5.
		A = 0xff
	if __obf_041320c70a794a0c[0] != '#' {
		return __obf_0c3c354bd419abd5, ColorHexFormatErr
	}
	__obf_3b4a2837c49e697a := func(__obf_3536db1209733836 byte) byte {
		switch {
		case __obf_3536db1209733836 >= '0' && __obf_3536db1209733836 <= '9':
			return __obf_3536db1209733836 - '0'
		case __obf_3536db1209733836 >= 'a' && __obf_3536db1209733836 <= 'f':
			return __obf_3536db1209733836 - 'a' + 10
		case __obf_3536db1209733836 >= 'A' && __obf_3536db1209733836 <= 'F':
			return __obf_3536db1209733836 - 'A' + 10
		}
		__obf_7834a057dbd32804 = ColorInvalidErr
		return 0
	}

	switch len(__obf_041320c70a794a0c) {
	case 7:
		__obf_0c3c354bd419abd5.
			R = __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[1])<<4 + __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[2])
		__obf_0c3c354bd419abd5.
			G = __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[3])<<4 + __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[4])
		__obf_0c3c354bd419abd5.
			B = __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[5])<<4 + __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[6])

	case 4:
		__obf_0c3c354bd419abd5.
			R = __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[1]) * 17
		__obf_0c3c354bd419abd5.
			G = __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[2]) * 17
		__obf_0c3c354bd419abd5.
			B = __obf_3b4a2837c49e697a(__obf_041320c70a794a0c[3]) * 17
	default:
		__obf_7834a057dbd32804 = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_e33026515c109ace string) (bool, error) {
	_, __obf_7834a057dbd32804 := os.Stat(__obf_e33026515c109ace)
	if __obf_7834a057dbd32804 == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_7834a057dbd32804) {
		return false, nil
	}
	return false, __obf_7834a057dbd32804
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_13028d2043d852d7 string) bool {
	for _, __obf_3b0471d953e55548 := range __obf_13028d2043d852d7 {
		if unicode.Is(unicode.Scripts["Han"], __obf_3b0471d953e55548) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_13028d2043d852d7 string) int {
	return utf8.RuneCountInString(__obf_13028d2043d852d7)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_e4a997e3bbea7348 []color.Color) color.RGBA {
	__obf_cd08b2bfb29b0c32 := PickN(__obf_e4a997e3bbea7348, 1)[0]
	__obf_3b0471d953e55548, __obf_c09dade8b2b5ab0f, __obf_3536db1209733836, __obf_2cfffe814d05be73 := __obf_cd08b2bfb29b0c32.RGBA()
	return color.RGBA{R: uint8(__obf_3b0471d953e55548), G: uint8(__obf_c09dade8b2b5ab0f), B: uint8(__obf_3536db1209733836), A: uint8(__obf_2cfffe814d05be73)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_9b0052aae736215a int, __obf_f857aea34adf2b31 int, __obf_29c6dafe17296cb0 image.Image) image.Point {
	__obf_3536db1209733836 := __obf_29c6dafe17296cb0.Bounds()
	__obf_7bb33930d21eacf3 := __obf_3536db1209733836.Max.X
	__obf_9a34940961e5b500 := __obf_3536db1209733836.Max.Y
	__obf_1519b886f64cb4a7 := 0
	__obf_1835f23b20a3e5dd := 0

	if __obf_7bb33930d21eacf3-__obf_9b0052aae736215a > 0 {
		__obf_1519b886f64cb4a7 = rand.IntN(__obf_7bb33930d21eacf3 - __obf_9b0052aae736215a)
	}
	if __obf_9a34940961e5b500-__obf_f857aea34adf2b31 > 0 {
		__obf_1835f23b20a3e5dd = rand.IntN(__obf_9a34940961e5b500 - __obf_f857aea34adf2b31)
	}

	return image.Point{
		X: __obf_1519b886f64cb4a7,
		Y: __obf_1835f23b20a3e5dd,
	}
}

func RandString(__obf_997845342b507d6c []string) string {
	__obf_49105b702a8ce36a := rand.N(len(__obf_997845342b507d6c))
	return __obf_997845342b507d6c[__obf_49105b702a8ce36a]
}

func PickN[T any](__obf_d2d1cf37b013667b []T, __obf_b506e48722ffcbf3 int) []T {
	if __obf_b506e48722ffcbf3 <= 0 || len(__obf_d2d1cf37b013667b) == 0 {
		return nil
	}
	if __obf_b506e48722ffcbf3 >= len(__obf_d2d1cf37b013667b) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_d2d1cf37b013667b...)
	}
	__obf_187e224478c75a75 := // 拷贝一份，避免影响原 slice
		append([]T(nil), __obf_d2d1cf37b013667b...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_187e224478c75a75), func(__obf_539cce309b069944, __obf_65f87cc9bf783ef8 int) {
		__obf_187e224478c75a75[__obf_539cce309b069944], __obf_187e224478c75a75[__obf_65f87cc9bf783ef8] = __obf_187e224478c75a75[__obf_65f87cc9bf783ef8], __obf_187e224478c75a75[__obf_539cce309b069944]
	})

	return __obf_187e224478c75a75[:__obf_b506e48722ffcbf3]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
