package __obf_0e73be9e4c3ea3fd

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand/v2"
	"os"
	"path"
	"strconv"
	"unicode"
	"unicode/utf8"

	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

var (
	ImageEmptyErr       = errors.New("image is empty")
	ImageMissingDataErr = errors.New("missing image data")
)

// saveToFile saves an image to a file
func __obf_fe2051ae37fdf99e(__obf_d1b36837c286905e image.Image, __obf_674636158f7d2792 string, __obf_22687f349d736f7d bool, __obf_01f2bb93c2c9908f int) error {
	var __obf_54253066e2fe26c4 *os.File
	var __obf_335fb0d6019c95ed error

	__obf_335fb0d6019c95ed = os.MkdirAll(path.Dir(__obf_674636158f7d2792), os.ModePerm)
	if __obf_335fb0d6019c95ed != nil {
		return __obf_335fb0d6019c95ed
	}

	if _, __obf_335fb0d6019c95ed = os.Stat(__obf_674636158f7d2792); os.IsNotExist(__obf_335fb0d6019c95ed) {
		__obf_54253066e2fe26c4, __obf_335fb0d6019c95ed = os.Create(__obf_674636158f7d2792)
	} else {
		__obf_54253066e2fe26c4, __obf_335fb0d6019c95ed = os.OpenFile(__obf_674636158f7d2792, os.O_RDWR, 0666)
	}
	if __obf_335fb0d6019c95ed != nil {
		return __obf_335fb0d6019c95ed
	}
	defer __obf_54253066e2fe26c4.Close()

	if __obf_22687f349d736f7d {
		__obf_335fb0d6019c95ed = png.Encode(__obf_54253066e2fe26c4, __obf_d1b36837c286905e)
	} else {
		__obf_335fb0d6019c95ed = jpeg.Encode(__obf_54253066e2fe26c4, __obf_d1b36837c286905e, &jpeg.Options{Quality: __obf_01f2bb93c2c9908f})
	}

	return __obf_335fb0d6019c95ed
}

// RotatePoint rotates a point's coordinates
func RotatePoint(__obf_fe086773828b5535, __obf_44ef04afc5c8241a, __obf_0f4ac1279453b377, __obf_282c0e1502614485 float64) (float64, float64) {
	return __obf_fe086773828b5535*__obf_282c0e1502614485 - __obf_44ef04afc5c8241a*__obf_0f4ac1279453b377, __obf_fe086773828b5535*__obf_0f4ac1279453b377 + __obf_44ef04afc5c8241a*__obf_282c0e1502614485
}

// RotatedSize calculates the size after rotation
func RotatedSize(__obf_6e1fff01dcc62820, __obf_462a1d8e0f4b0353 int, __obf_b51ed5d77ee8eb95 float64) (int, int) {
	if __obf_6e1fff01dcc62820 <= 0 || __obf_462a1d8e0f4b0353 <= 0 {
		return 0, 0
	}

	__obf_0f4ac1279453b377, __obf_282c0e1502614485 := math.Sincos(math.Pi * __obf_b51ed5d77ee8eb95 / 180)
	__obf_b78ac2d2e1cf7363, __obf_cdc91d61a35557f1 := RotatePoint(float64(__obf_6e1fff01dcc62820-1), 0, __obf_0f4ac1279453b377, __obf_282c0e1502614485)
	__obf_578532d304e3a7c8, __obf_c1bf9607c29dde23 := RotatePoint(float64(__obf_6e1fff01dcc62820-1), float64(__obf_462a1d8e0f4b0353-1), __obf_0f4ac1279453b377, __obf_282c0e1502614485)
	__obf_9fc14d8abfaed74a, __obf_958d17d9d30b5155 := RotatePoint(0, float64(__obf_462a1d8e0f4b0353-1), __obf_0f4ac1279453b377, __obf_282c0e1502614485)

	__obf_8a297e9a65a6b461 := min(__obf_b78ac2d2e1cf7363, min(__obf_578532d304e3a7c8, min(__obf_9fc14d8abfaed74a, 0)))
	__obf_6dbed9412bb93f30 := max(__obf_b78ac2d2e1cf7363, max(__obf_578532d304e3a7c8, max(__obf_9fc14d8abfaed74a, 0)))
	__obf_a6d1e2b517fb3ca3 := min(__obf_cdc91d61a35557f1, min(__obf_c1bf9607c29dde23, min(__obf_958d17d9d30b5155, 0)))
	__obf_23d5edea6d8dce9a := max(__obf_cdc91d61a35557f1, max(__obf_c1bf9607c29dde23, max(__obf_958d17d9d30b5155, 0)))

	__obf_0107a6cb34e1b800 := __obf_6dbed9412bb93f30 - __obf_8a297e9a65a6b461 + 1
	if __obf_0107a6cb34e1b800-math.Floor(__obf_0107a6cb34e1b800) > 0.1 {
		__obf_0107a6cb34e1b800++
	}
	__obf_1f3235dfd61ebd0b := __obf_23d5edea6d8dce9a - __obf_a6d1e2b517fb3ca3 + 1
	if __obf_1f3235dfd61ebd0b-math.Floor(__obf_1f3235dfd61ebd0b) > 0.1 {
		__obf_1f3235dfd61ebd0b++
	}

	return int(__obf_0107a6cb34e1b800), int(__obf_1f3235dfd61ebd0b)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_3d84420b7c1bf00e image.Rectangle, __obf_0107a6cb34e1b800 int, __obf_1f3235dfd61ebd0b int, __obf_a03a45f487b7672e bool) image.Rectangle {
	var __obf_e5bbe575453b9248 image.Rectangle
	if __obf_0107a6cb34e1b800*__obf_3d84420b7c1bf00e.Dy() < __obf_1f3235dfd61ebd0b*__obf_3d84420b7c1bf00e.Dx() {
		__obf_ce59ab8babef59ae := float64(__obf_0107a6cb34e1b800) / float64(__obf_3d84420b7c1bf00e.Dx())

		__obf_8490a5c406d4303a := int(float64(__obf_3d84420b7c1bf00e.Dy()) * __obf_ce59ab8babef59ae)
		__obf_d6c32691ec6a01ec := 0
		if __obf_a03a45f487b7672e {
			__obf_d6c32691ec6a01ec = (__obf_1f3235dfd61ebd0b - __obf_8490a5c406d4303a) / 2
		}
		__obf_e5bbe575453b9248 = image.Rect(0, __obf_d6c32691ec6a01ec, __obf_0107a6cb34e1b800, __obf_d6c32691ec6a01ec+__obf_8490a5c406d4303a)
	} else {
		__obf_ce59ab8babef59ae := float64(__obf_1f3235dfd61ebd0b) / float64(__obf_3d84420b7c1bf00e.Dy())
		__obf_f9050723b571b0fb := int(float64(__obf_3d84420b7c1bf00e.Dx()) * __obf_ce59ab8babef59ae)
		__obf_d6c32691ec6a01ec := 0
		if __obf_a03a45f487b7672e {
			__obf_d6c32691ec6a01ec = (__obf_0107a6cb34e1b800 - __obf_f9050723b571b0fb) / 2
		}
		__obf_e5bbe575453b9248 = image.Rect(__obf_d6c32691ec6a01ec, 0, __obf_d6c32691ec6a01ec+__obf_f9050723b571b0fb, __obf_1f3235dfd61ebd0b)
	}

	return __obf_e5bbe575453b9248
}

// t2x converts an integer to a hexadecimal string
func __obf_5ede71559354c72a(__obf_ef01b999744a29da int64) string {
	__obf_57d21e60bb197233 := strconv.FormatInt(__obf_ef01b999744a29da, 16)
	if len(__obf_57d21e60bb197233) == 1 {
		__obf_57d21e60bb197233 = "0" + __obf_57d21e60bb197233
	}
	return __obf_57d21e60bb197233
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_7b6c329fad835ae3 float32) uint8 {
	__obf_51d4fa10e4721296 := min(float64(__obf_7b6c329fad835ae3), 1)
	__obf_73dc767a16a750c2 := __obf_51d4fa10e4721296 * 255
	return uint8(__obf_73dc767a16a750c2)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_61bf14a4668572e2 int64, __obf_4e99a7022b6bcf19 int64, __obf_48205805993a948a int64) string {
	__obf_a6fdf04202ab9ad1 := __obf_5ede71559354c72a(__obf_61bf14a4668572e2)
	__obf_453cdb4a628cd8c2 := __obf_5ede71559354c72a(__obf_4e99a7022b6bcf19)
	__obf_e3b1e50668ab4d43 := __obf_5ede71559354c72a(__obf_48205805993a948a)
	return __obf_a6fdf04202ab9ad1 + __obf_453cdb4a628cd8c2 + __obf_e3b1e50668ab4d43
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_ac29e23fd5e7b424 string) (int64, int64, int64) {
	__obf_a6fdf04202ab9ad1, _ := strconv.ParseInt(__obf_ac29e23fd5e7b424[:2], 16, 10)
	__obf_453cdb4a628cd8c2, _ := strconv.ParseInt(__obf_ac29e23fd5e7b424[2:4], 16, 18)
	__obf_e3b1e50668ab4d43, _ := strconv.ParseInt(__obf_ac29e23fd5e7b424[4:], 16, 10)
	return __obf_a6fdf04202ab9ad1, __obf_453cdb4a628cd8c2, __obf_e3b1e50668ab4d43
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_f2c63d2da6f1cab9 string) (__obf_a5de954533461107 color.RGBA, __obf_335fb0d6019c95ed error) {
	__obf_a5de954533461107.A = 0xff
	if __obf_f2c63d2da6f1cab9[0] != '#' {
		return __obf_a5de954533461107, ColorHexFormatErr
	}

	__obf_99cb4681946dfe3c := func(__obf_e3b1e50668ab4d43 byte) byte {
		switch {
		case __obf_e3b1e50668ab4d43 >= '0' && __obf_e3b1e50668ab4d43 <= '9':
			return __obf_e3b1e50668ab4d43 - '0'
		case __obf_e3b1e50668ab4d43 >= 'a' && __obf_e3b1e50668ab4d43 <= 'f':
			return __obf_e3b1e50668ab4d43 - 'a' + 10
		case __obf_e3b1e50668ab4d43 >= 'A' && __obf_e3b1e50668ab4d43 <= 'F':
			return __obf_e3b1e50668ab4d43 - 'A' + 10
		}
		__obf_335fb0d6019c95ed = ColorInvalidErr
		return 0
	}

	switch len(__obf_f2c63d2da6f1cab9) {
	case 7:
		__obf_a5de954533461107.R = __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[1])<<4 + __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[2])
		__obf_a5de954533461107.G = __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[3])<<4 + __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[4])
		__obf_a5de954533461107.B = __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[5])<<4 + __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[6])

	case 4:
		__obf_a5de954533461107.R = __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[1]) * 17
		__obf_a5de954533461107.G = __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[2]) * 17
		__obf_a5de954533461107.B = __obf_99cb4681946dfe3c(__obf_f2c63d2da6f1cab9[3]) * 17
	default:
		__obf_335fb0d6019c95ed = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(path string) (bool, error) {
	_, __obf_335fb0d6019c95ed := os.Stat(path)
	if __obf_335fb0d6019c95ed == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_335fb0d6019c95ed) {
		return false, nil
	}
	return false, __obf_335fb0d6019c95ed
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_fcdfbbb8ecccdaee string) bool {
	for _, __obf_a6fdf04202ab9ad1 := range __obf_fcdfbbb8ecccdaee {
		if unicode.Is(unicode.Scripts["Han"], __obf_a6fdf04202ab9ad1) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_fcdfbbb8ecccdaee string) int {
	return utf8.RuneCountInString(__obf_fcdfbbb8ecccdaee)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_f8353e66f8032833 []color.Color) color.RGBA {
	__obf_512af10b41253004 := PickN(__obf_f8353e66f8032833, 1)[0]
	__obf_a6fdf04202ab9ad1, __obf_453cdb4a628cd8c2, __obf_e3b1e50668ab4d43, __obf_51d4fa10e4721296 := __obf_512af10b41253004.RGBA()
	return color.RGBA{R: uint8(__obf_a6fdf04202ab9ad1), G: uint8(__obf_453cdb4a628cd8c2), B: uint8(__obf_e3b1e50668ab4d43), A: uint8(__obf_51d4fa10e4721296)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_0107a6cb34e1b800 int, __obf_1f3235dfd61ebd0b int, __obf_d1b36837c286905e image.Image) image.Point {
	__obf_e3b1e50668ab4d43 := __obf_d1b36837c286905e.Bounds()
	__obf_8d52e892278b0cd0 := __obf_e3b1e50668ab4d43.Max.X
	__obf_439af1c582c925b5 := __obf_e3b1e50668ab4d43.Max.Y
	__obf_2c414250ee2fdc99 := 0
	__obf_d4add37130b0ff70 := 0

	if __obf_8d52e892278b0cd0-__obf_0107a6cb34e1b800 > 0 {
		__obf_2c414250ee2fdc99 = rand.IntN(__obf_8d52e892278b0cd0 - __obf_0107a6cb34e1b800)
	}
	if __obf_439af1c582c925b5-__obf_1f3235dfd61ebd0b > 0 {
		__obf_d4add37130b0ff70 = rand.IntN(__obf_439af1c582c925b5 - __obf_1f3235dfd61ebd0b)
	}

	return image.Point{
		X: __obf_2c414250ee2fdc99,
		Y: __obf_d4add37130b0ff70,
	}
}

// DrawString draws a string on the canvas with high-quality rendering
func DrawString(__obf_c0d27cca35ba9554 draw.Image, __obf_eba91f5957ab6eb6 *DrawStringParam, __obf_339dacedf4b35ce3 fixed.Point26_6) error {
	__obf_eda6f58c871c600b, __obf_335fb0d6019c95ed := opentype.NewFace(__obf_eba91f5957ab6eb6.Font, &opentype.FaceOptions{
		Size:    float64(__obf_eba91f5957ab6eb6.FontSize),
		DPI:     float64(__obf_eba91f5957ab6eb6.FontDPI),
		Hinting: __obf_eba91f5957ab6eb6.FontHinting,
	})
	if __obf_335fb0d6019c95ed != nil {
		return __obf_335fb0d6019c95ed
	}
	defer __obf_eda6f58c871c600b.Close()

	__obf_94c407ab26ed3d9c := &font.Drawer{
		Dst:  __obf_c0d27cca35ba9554,
		Src:  image.NewUniform(__obf_eba91f5957ab6eb6.Color),
		Face: __obf_eda6f58c871c600b,
		Dot:  __obf_339dacedf4b35ce3,
	}
	__obf_94c407ab26ed3d9c.DrawString(__obf_eba91f5957ab6eb6.Text)

	return nil
}

func RandString(__obf_88fa6b99677dd6d4 []string) string {
	__obf_a55706027c208a67 := rand.N(len(__obf_88fa6b99677dd6d4))
	return __obf_88fa6b99677dd6d4[__obf_a55706027c208a67]
}

func PickN[T any](__obf_23c7bdced1de6276 []T, __obf_24d6334a9384367b int) []T {
	if __obf_24d6334a9384367b <= 0 || len(__obf_23c7bdced1de6276) == 0 {
		return nil
	}
	if __obf_24d6334a9384367b >= len(__obf_23c7bdced1de6276) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_23c7bdced1de6276...)
	}

	// 拷贝一份，避免影响原 slice
	__obf_fd8d5c796a4e1884 := append([]T(nil), __obf_23c7bdced1de6276...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_fd8d5c796a4e1884), func(__obf_c09c9af0804efac3, __obf_f95bc0d48c52ee44 int) {
		__obf_fd8d5c796a4e1884[__obf_c09c9af0804efac3], __obf_fd8d5c796a4e1884[__obf_f95bc0d48c52ee44] = __obf_fd8d5c796a4e1884[__obf_f95bc0d48c52ee44], __obf_fd8d5c796a4e1884[__obf_c09c9af0804efac3]
	})

	return __obf_fd8d5c796a4e1884[:__obf_24d6334a9384367b]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
