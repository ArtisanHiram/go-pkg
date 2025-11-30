package __obf_611d3abd867098c0

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
func RotatePoint(__obf_30abeb13ff9c4cb6, __obf_be66a418f9efe81c, __obf_1cfc73210bd90dee, __obf_ec1a0e58e51ac1c5 float64) (float64, float64) {
	return __obf_30abeb13ff9c4cb6*__obf_ec1a0e58e51ac1c5 - __obf_be66a418f9efe81c*__obf_1cfc73210bd90dee,

		// RotatedSize calculates the size after rotation
		__obf_30abeb13ff9c4cb6*__obf_1cfc73210bd90dee + __obf_be66a418f9efe81c*__obf_ec1a0e58e51ac1c5
}

func RotatedSize(__obf_f76f450cbdec6b3c, __obf_9b73d8da36219bd6 int, __obf_5904b7d2827b8b6f float64) (int, int) {
	if __obf_f76f450cbdec6b3c <= 0 || __obf_9b73d8da36219bd6 <= 0 {
		return 0, 0
	}
	__obf_1cfc73210bd90dee, __obf_ec1a0e58e51ac1c5 := math.Sincos(math.Pi * __obf_5904b7d2827b8b6f / 180)
	__obf_3eee80ccbca0204a, __obf_81069e6cd48371a9 := RotatePoint(float64(__obf_f76f450cbdec6b3c-1), 0, __obf_1cfc73210bd90dee, __obf_ec1a0e58e51ac1c5)
	__obf_a9b687df02403629, __obf_e0a7a55b3e66706f := RotatePoint(float64(__obf_f76f450cbdec6b3c-1), float64(__obf_9b73d8da36219bd6-1), __obf_1cfc73210bd90dee, __obf_ec1a0e58e51ac1c5)
	__obf_a2279bf91302ce1c, __obf_4b1675e0db67d4ba := RotatePoint(0, float64(__obf_9b73d8da36219bd6-1), __obf_1cfc73210bd90dee, __obf_ec1a0e58e51ac1c5)
	__obf_d3417ff6dcbc49c9 := min(__obf_3eee80ccbca0204a, min(__obf_a9b687df02403629, min(__obf_a2279bf91302ce1c, 0)))
	__obf_5bac3be9e4f9389c := max(__obf_3eee80ccbca0204a, max(__obf_a9b687df02403629, max(__obf_a2279bf91302ce1c, 0)))
	__obf_620161024513b289 := min(__obf_81069e6cd48371a9, min(__obf_e0a7a55b3e66706f, min(__obf_4b1675e0db67d4ba, 0)))
	__obf_ed710250d2109971 := max(__obf_81069e6cd48371a9, max(__obf_e0a7a55b3e66706f, max(__obf_4b1675e0db67d4ba, 0)))
	__obf_648fefc6b7e6cf06 := __obf_5bac3be9e4f9389c - __obf_d3417ff6dcbc49c9 + 1
	if __obf_648fefc6b7e6cf06-math.Floor(__obf_648fefc6b7e6cf06) > 0.1 {
		__obf_648fefc6b7e6cf06++
	}
	__obf_1b3aef7994082943 := __obf_ed710250d2109971 - __obf_620161024513b289 + 1
	if __obf_1b3aef7994082943-math.Floor(__obf_1b3aef7994082943) > 0.1 {
		__obf_1b3aef7994082943++
	}

	return int(__obf_648fefc6b7e6cf06), int(__obf_1b3aef7994082943)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_97286c05f1a190ac image.Rectangle, __obf_648fefc6b7e6cf06 int, __obf_1b3aef7994082943 int, __obf_1896a412032f1768 bool) image.Rectangle {
	var __obf_59ef46f3d921a96f image.Rectangle
	if __obf_648fefc6b7e6cf06*__obf_97286c05f1a190ac.Dy() < __obf_1b3aef7994082943*__obf_97286c05f1a190ac.Dx() {
		__obf_f3dce07c6ef04a42 := float64(__obf_648fefc6b7e6cf06) / float64(__obf_97286c05f1a190ac.Dx())
		__obf_292c621f3352bcec := int(float64(__obf_97286c05f1a190ac.Dy()) * __obf_f3dce07c6ef04a42)
		__obf_fe985d46cfa5a12f := 0
		if __obf_1896a412032f1768 {
			__obf_fe985d46cfa5a12f = (__obf_1b3aef7994082943 - __obf_292c621f3352bcec) / 2
		}
		__obf_59ef46f3d921a96f = image.Rect(0, __obf_fe985d46cfa5a12f, __obf_648fefc6b7e6cf06, __obf_fe985d46cfa5a12f+__obf_292c621f3352bcec)
	} else {
		__obf_f3dce07c6ef04a42 := float64(__obf_1b3aef7994082943) / float64(__obf_97286c05f1a190ac.Dy())
		__obf_87827455cc20b8ba := int(float64(__obf_97286c05f1a190ac.Dx()) * __obf_f3dce07c6ef04a42)
		__obf_fe985d46cfa5a12f := 0
		if __obf_1896a412032f1768 {
			__obf_fe985d46cfa5a12f = (__obf_648fefc6b7e6cf06 - __obf_87827455cc20b8ba) / 2
		}
		__obf_59ef46f3d921a96f = image.Rect(__obf_fe985d46cfa5a12f, 0, __obf_fe985d46cfa5a12f+__obf_87827455cc20b8ba, __obf_1b3aef7994082943)
	}

	return __obf_59ef46f3d921a96f
}

// t2x converts an integer to a hexadecimal string
func __obf_aaccf2019481b17f(__obf_8e3734bbf55f3c20 int64) string {
	__obf_1879d3401478787a := strconv.FormatInt(__obf_8e3734bbf55f3c20, 16)
	if len(__obf_1879d3401478787a) == 1 {
		__obf_1879d3401478787a = "0" + __obf_1879d3401478787a
	}
	return __obf_1879d3401478787a
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_423866825060aa77 float32) uint8 {
	__obf_394daac6f239348e := min(float64(__obf_423866825060aa77), 1)
	__obf_ac171d653f050888 := __obf_394daac6f239348e * 255
	return uint8(__obf_ac171d653f050888)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_00c266cd6774664f int64, __obf_c5f9ba2a834792fd int64, __obf_9a01d1c8e81008d2 int64) string {
	__obf_5112dca4c8ba2fb4 := __obf_aaccf2019481b17f(__obf_00c266cd6774664f)
	__obf_7b384eb7f1b94a50 := __obf_aaccf2019481b17f(__obf_c5f9ba2a834792fd)
	__obf_39012e53e716527b := __obf_aaccf2019481b17f(__obf_9a01d1c8e81008d2)
	return __obf_5112dca4c8ba2fb4 + __obf_7b384eb7f1b94a50 + __obf_39012e53e716527b
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_d1eb37dc7c34ee60 string) (int64, int64, int64) {
	__obf_5112dca4c8ba2fb4, _ := strconv.ParseInt(__obf_d1eb37dc7c34ee60[:2], 16, 10)
	__obf_7b384eb7f1b94a50, _ := strconv.ParseInt(__obf_d1eb37dc7c34ee60[2:4], 16, 18)
	__obf_39012e53e716527b, _ := strconv.ParseInt(__obf_d1eb37dc7c34ee60[4:], 16, 10)
	return __obf_5112dca4c8ba2fb4, __obf_7b384eb7f1b94a50, __obf_39012e53e716527b
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_987b805c256d3274 string) (__obf_be38999717495b5d color.RGBA, __obf_e9ef0dd78ff22cdd error) {
	__obf_be38999717495b5d.
		A = 0xff
	if __obf_987b805c256d3274[0] != '#' {
		return __obf_be38999717495b5d, ColorHexFormatErr
	}
	__obf_7c35a08a27297fc5 := func(__obf_39012e53e716527b byte) byte {
		switch {
		case __obf_39012e53e716527b >= '0' && __obf_39012e53e716527b <= '9':
			return __obf_39012e53e716527b - '0'
		case __obf_39012e53e716527b >= 'a' && __obf_39012e53e716527b <= 'f':
			return __obf_39012e53e716527b - 'a' + 10
		case __obf_39012e53e716527b >= 'A' && __obf_39012e53e716527b <= 'F':
			return __obf_39012e53e716527b - 'A' + 10
		}
		__obf_e9ef0dd78ff22cdd = ColorInvalidErr
		return 0
	}

	switch len(__obf_987b805c256d3274) {
	case 7:
		__obf_be38999717495b5d.
			R = __obf_7c35a08a27297fc5(__obf_987b805c256d3274[1])<<4 + __obf_7c35a08a27297fc5(__obf_987b805c256d3274[2])
		__obf_be38999717495b5d.
			G = __obf_7c35a08a27297fc5(__obf_987b805c256d3274[3])<<4 + __obf_7c35a08a27297fc5(__obf_987b805c256d3274[4])
		__obf_be38999717495b5d.
			B = __obf_7c35a08a27297fc5(__obf_987b805c256d3274[5])<<4 + __obf_7c35a08a27297fc5(__obf_987b805c256d3274[6])

	case 4:
		__obf_be38999717495b5d.
			R = __obf_7c35a08a27297fc5(__obf_987b805c256d3274[1]) * 17
		__obf_be38999717495b5d.
			G = __obf_7c35a08a27297fc5(__obf_987b805c256d3274[2]) * 17
		__obf_be38999717495b5d.
			B = __obf_7c35a08a27297fc5(__obf_987b805c256d3274[3]) * 17
	default:
		__obf_e9ef0dd78ff22cdd = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_5868458a2a441067 string) (bool, error) {
	_, __obf_e9ef0dd78ff22cdd := os.Stat(__obf_5868458a2a441067)
	if __obf_e9ef0dd78ff22cdd == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_e9ef0dd78ff22cdd) {
		return false, nil
	}
	return false, __obf_e9ef0dd78ff22cdd
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_9cf65c20c0a6a4d9 string) bool {
	for _, __obf_5112dca4c8ba2fb4 := range __obf_9cf65c20c0a6a4d9 {
		if unicode.Is(unicode.Scripts["Han"], __obf_5112dca4c8ba2fb4) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_9cf65c20c0a6a4d9 string) int {
	return utf8.RuneCountInString(__obf_9cf65c20c0a6a4d9)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_154371eb1fbf0626 []color.Color) color.RGBA {
	__obf_b9a8aeb31e5bf396 := PickN(__obf_154371eb1fbf0626, 1)[0]
	__obf_5112dca4c8ba2fb4, __obf_7b384eb7f1b94a50, __obf_39012e53e716527b, __obf_394daac6f239348e := __obf_b9a8aeb31e5bf396.RGBA()
	return color.RGBA{R: uint8(__obf_5112dca4c8ba2fb4), G: uint8(__obf_7b384eb7f1b94a50), B: uint8(__obf_39012e53e716527b), A: uint8(__obf_394daac6f239348e)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_648fefc6b7e6cf06 int, __obf_1b3aef7994082943 int, __obf_6922282326d724ae image.Image) image.Point {
	__obf_39012e53e716527b := __obf_6922282326d724ae.Bounds()
	__obf_a97e35597f244ad7 := __obf_39012e53e716527b.Max.X
	__obf_90e312edaa4d2138 := __obf_39012e53e716527b.Max.Y
	__obf_57a631362164bb1d := 0
	__obf_c9d814f564b525fc := 0

	if __obf_a97e35597f244ad7-__obf_648fefc6b7e6cf06 > 0 {
		__obf_57a631362164bb1d = rand.IntN(__obf_a97e35597f244ad7 - __obf_648fefc6b7e6cf06)
	}
	if __obf_90e312edaa4d2138-__obf_1b3aef7994082943 > 0 {
		__obf_c9d814f564b525fc = rand.IntN(__obf_90e312edaa4d2138 - __obf_1b3aef7994082943)
	}

	return image.Point{
		X: __obf_57a631362164bb1d,
		Y: __obf_c9d814f564b525fc,
	}
}

func RandString(__obf_1ba8eb62a808ff3e []string) string {
	__obf_f7f59e8f8c2a912c := rand.N(len(__obf_1ba8eb62a808ff3e))
	return __obf_1ba8eb62a808ff3e[__obf_f7f59e8f8c2a912c]
}

func PickN[T any](__obf_6769c7ea99f38f12 []T, __obf_35f22da933056d2c int) []T {
	if __obf_35f22da933056d2c <= 0 || len(__obf_6769c7ea99f38f12) == 0 {
		return nil
	}
	if __obf_35f22da933056d2c >= len(__obf_6769c7ea99f38f12) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_6769c7ea99f38f12...)
	}
	__obf_5153d3cc059e9d1e := // 拷贝一份，避免影响原 slice
		append([]T(nil), __obf_6769c7ea99f38f12...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_5153d3cc059e9d1e), func(__obf_e0d7d095e7ef8836, __obf_3dea8073d26c0563 int) {
		__obf_5153d3cc059e9d1e[__obf_e0d7d095e7ef8836], __obf_5153d3cc059e9d1e[__obf_3dea8073d26c0563] = __obf_5153d3cc059e9d1e[__obf_3dea8073d26c0563], __obf_5153d3cc059e9d1e[__obf_e0d7d095e7ef8836]
	})

	return __obf_5153d3cc059e9d1e[:__obf_35f22da933056d2c]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
