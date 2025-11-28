package __obf_9aafb22e4ca0e7fb

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
func RotatePoint(__obf_0b812693b84d04c2, __obf_03730d76b27c59ea, __obf_c83f156457284bb4, __obf_cae9bc17b9d9adb0 float64) (float64, float64) {
	return __obf_0b812693b84d04c2*__obf_cae9bc17b9d9adb0 - __obf_03730d76b27c59ea*__obf_c83f156457284bb4, __obf_0b812693b84d04c2*__obf_c83f156457284bb4 + __obf_03730d76b27c59ea*__obf_cae9bc17b9d9adb0
}

// RotatedSize calculates the size after rotation
func RotatedSize(__obf_0dd4e01c8068fdbf, __obf_6d2b3401cf8ebb02 int, __obf_d53fd2f383186464 float64) (int, int) {
	if __obf_0dd4e01c8068fdbf <= 0 || __obf_6d2b3401cf8ebb02 <= 0 {
		return 0, 0
	}

	__obf_c83f156457284bb4, __obf_cae9bc17b9d9adb0 := math.Sincos(math.Pi * __obf_d53fd2f383186464 / 180)
	__obf_6a668613f101a1ff, __obf_af1420bf6bba7b41 := RotatePoint(float64(__obf_0dd4e01c8068fdbf-1), 0, __obf_c83f156457284bb4, __obf_cae9bc17b9d9adb0)
	__obf_f1324b749a8c51ff, __obf_9139a374f0c46110 := RotatePoint(float64(__obf_0dd4e01c8068fdbf-1), float64(__obf_6d2b3401cf8ebb02-1), __obf_c83f156457284bb4, __obf_cae9bc17b9d9adb0)
	__obf_f278c1584da47cc7, __obf_c9e5a881297175fb := RotatePoint(0, float64(__obf_6d2b3401cf8ebb02-1), __obf_c83f156457284bb4, __obf_cae9bc17b9d9adb0)

	__obf_cde061c0214932f5 := min(__obf_6a668613f101a1ff, min(__obf_f1324b749a8c51ff, min(__obf_f278c1584da47cc7, 0)))
	__obf_bb4eaefbff63b95e := max(__obf_6a668613f101a1ff, max(__obf_f1324b749a8c51ff, max(__obf_f278c1584da47cc7, 0)))
	__obf_d5ff544c50c60cf5 := min(__obf_af1420bf6bba7b41, min(__obf_9139a374f0c46110, min(__obf_c9e5a881297175fb, 0)))
	__obf_94c3f51abe91523c := max(__obf_af1420bf6bba7b41, max(__obf_9139a374f0c46110, max(__obf_c9e5a881297175fb, 0)))

	__obf_1f0f1ee38d79568d := __obf_bb4eaefbff63b95e - __obf_cde061c0214932f5 + 1
	if __obf_1f0f1ee38d79568d-math.Floor(__obf_1f0f1ee38d79568d) > 0.1 {
		__obf_1f0f1ee38d79568d++
	}
	__obf_011652ee9d695352 := __obf_94c3f51abe91523c - __obf_d5ff544c50c60cf5 + 1
	if __obf_011652ee9d695352-math.Floor(__obf_011652ee9d695352) > 0.1 {
		__obf_011652ee9d695352++
	}

	return int(__obf_1f0f1ee38d79568d), int(__obf_011652ee9d695352)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_5aa5c6a1be9afbba image.Rectangle, __obf_1f0f1ee38d79568d int, __obf_011652ee9d695352 int, __obf_a3765594455b4b47 bool) image.Rectangle {
	var __obf_70a17d5d7e9e1495 image.Rectangle
	if __obf_1f0f1ee38d79568d*__obf_5aa5c6a1be9afbba.Dy() < __obf_011652ee9d695352*__obf_5aa5c6a1be9afbba.Dx() {
		__obf_a710b0b012848d24 := float64(__obf_1f0f1ee38d79568d) / float64(__obf_5aa5c6a1be9afbba.Dx())

		__obf_9862f26f093f39ea := int(float64(__obf_5aa5c6a1be9afbba.Dy()) * __obf_a710b0b012848d24)
		__obf_029f53e396c00469 := 0
		if __obf_a3765594455b4b47 {
			__obf_029f53e396c00469 = (__obf_011652ee9d695352 - __obf_9862f26f093f39ea) / 2
		}
		__obf_70a17d5d7e9e1495 = image.Rect(0, __obf_029f53e396c00469, __obf_1f0f1ee38d79568d, __obf_029f53e396c00469+__obf_9862f26f093f39ea)
	} else {
		__obf_a710b0b012848d24 := float64(__obf_011652ee9d695352) / float64(__obf_5aa5c6a1be9afbba.Dy())
		__obf_4750f0ec278675f6 := int(float64(__obf_5aa5c6a1be9afbba.Dx()) * __obf_a710b0b012848d24)
		__obf_029f53e396c00469 := 0
		if __obf_a3765594455b4b47 {
			__obf_029f53e396c00469 = (__obf_1f0f1ee38d79568d - __obf_4750f0ec278675f6) / 2
		}
		__obf_70a17d5d7e9e1495 = image.Rect(__obf_029f53e396c00469, 0, __obf_029f53e396c00469+__obf_4750f0ec278675f6, __obf_011652ee9d695352)
	}

	return __obf_70a17d5d7e9e1495
}

// t2x converts an integer to a hexadecimal string
func __obf_876c4952a0fda3d5(__obf_ac2cb4c7f23a9d33 int64) string {
	__obf_29fdafb5bf26169c := strconv.FormatInt(__obf_ac2cb4c7f23a9d33, 16)
	if len(__obf_29fdafb5bf26169c) == 1 {
		__obf_29fdafb5bf26169c = "0" + __obf_29fdafb5bf26169c
	}
	return __obf_29fdafb5bf26169c
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_c5e756f6dba19b40 float32) uint8 {
	__obf_d0d17c4446fa2398 := min(float64(__obf_c5e756f6dba19b40), 1)
	__obf_a1917eb6b4ef27bb := __obf_d0d17c4446fa2398 * 255
	return uint8(__obf_a1917eb6b4ef27bb)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_8a969f5ece4652ba int64, __obf_97e4694ec5eb8677 int64, __obf_bc58926e1bb7ec77 int64) string {
	__obf_a7a2e27dcfb5ddb3 := __obf_876c4952a0fda3d5(__obf_8a969f5ece4652ba)
	__obf_76a109c04a565057 := __obf_876c4952a0fda3d5(__obf_97e4694ec5eb8677)
	__obf_af1f8c99779118ac := __obf_876c4952a0fda3d5(__obf_bc58926e1bb7ec77)
	return __obf_a7a2e27dcfb5ddb3 + __obf_76a109c04a565057 + __obf_af1f8c99779118ac
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_edc030beb09845a7 string) (int64, int64, int64) {
	__obf_a7a2e27dcfb5ddb3, _ := strconv.ParseInt(__obf_edc030beb09845a7[:2], 16, 10)
	__obf_76a109c04a565057, _ := strconv.ParseInt(__obf_edc030beb09845a7[2:4], 16, 18)
	__obf_af1f8c99779118ac, _ := strconv.ParseInt(__obf_edc030beb09845a7[4:], 16, 10)
	return __obf_a7a2e27dcfb5ddb3, __obf_76a109c04a565057, __obf_af1f8c99779118ac
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_0fd189327138320e string) (__obf_6995714a45612c24 color.RGBA, __obf_c8d00ac01db0efb5 error) {
	__obf_6995714a45612c24.A = 0xff
	if __obf_0fd189327138320e[0] != '#' {
		return __obf_6995714a45612c24, ColorHexFormatErr
	}

	__obf_5c6902f2c3b4a66d := func(__obf_af1f8c99779118ac byte) byte {
		switch {
		case __obf_af1f8c99779118ac >= '0' && __obf_af1f8c99779118ac <= '9':
			return __obf_af1f8c99779118ac - '0'
		case __obf_af1f8c99779118ac >= 'a' && __obf_af1f8c99779118ac <= 'f':
			return __obf_af1f8c99779118ac - 'a' + 10
		case __obf_af1f8c99779118ac >= 'A' && __obf_af1f8c99779118ac <= 'F':
			return __obf_af1f8c99779118ac - 'A' + 10
		}
		__obf_c8d00ac01db0efb5 = ColorInvalidErr
		return 0
	}

	switch len(__obf_0fd189327138320e) {
	case 7:
		__obf_6995714a45612c24.R = __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[1])<<4 + __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[2])
		__obf_6995714a45612c24.G = __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[3])<<4 + __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[4])
		__obf_6995714a45612c24.B = __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[5])<<4 + __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[6])

	case 4:
		__obf_6995714a45612c24.R = __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[1]) * 17
		__obf_6995714a45612c24.G = __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[2]) * 17
		__obf_6995714a45612c24.B = __obf_5c6902f2c3b4a66d(__obf_0fd189327138320e[3]) * 17
	default:
		__obf_c8d00ac01db0efb5 = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_bfa0ac9754068b62 string) (bool, error) {
	_, __obf_c8d00ac01db0efb5 := os.Stat(__obf_bfa0ac9754068b62)
	if __obf_c8d00ac01db0efb5 == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_c8d00ac01db0efb5) {
		return false, nil
	}
	return false, __obf_c8d00ac01db0efb5
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_d2f6dc7eae80ad08 string) bool {
	for _, __obf_a7a2e27dcfb5ddb3 := range __obf_d2f6dc7eae80ad08 {
		if unicode.Is(unicode.Scripts["Han"], __obf_a7a2e27dcfb5ddb3) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_d2f6dc7eae80ad08 string) int {
	return utf8.RuneCountInString(__obf_d2f6dc7eae80ad08)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_5575f705b6f763bb []color.Color) color.RGBA {
	__obf_3f80174223929206 := PickN(__obf_5575f705b6f763bb, 1)[0]
	__obf_a7a2e27dcfb5ddb3, __obf_76a109c04a565057, __obf_af1f8c99779118ac, __obf_d0d17c4446fa2398 := __obf_3f80174223929206.RGBA()
	return color.RGBA{R: uint8(__obf_a7a2e27dcfb5ddb3), G: uint8(__obf_76a109c04a565057), B: uint8(__obf_af1f8c99779118ac), A: uint8(__obf_d0d17c4446fa2398)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_1f0f1ee38d79568d int, __obf_011652ee9d695352 int, __obf_b19838e146b51f6e image.Image) image.Point {
	__obf_af1f8c99779118ac := __obf_b19838e146b51f6e.Bounds()
	__obf_6d935acbd673374f := __obf_af1f8c99779118ac.Max.X
	__obf_e1fff36a1b54fbe4 := __obf_af1f8c99779118ac.Max.Y
	__obf_9859b0a842b73ec6 := 0
	__obf_355e280ae6a7fbfe := 0

	if __obf_6d935acbd673374f-__obf_1f0f1ee38d79568d > 0 {
		__obf_9859b0a842b73ec6 = rand.IntN(__obf_6d935acbd673374f - __obf_1f0f1ee38d79568d)
	}
	if __obf_e1fff36a1b54fbe4-__obf_011652ee9d695352 > 0 {
		__obf_355e280ae6a7fbfe = rand.IntN(__obf_e1fff36a1b54fbe4 - __obf_011652ee9d695352)
	}

	return image.Point{
		X: __obf_9859b0a842b73ec6,
		Y: __obf_355e280ae6a7fbfe,
	}
}

func RandString(__obf_24fb59573a7563e5 []string) string {
	__obf_fde7c1857dd05896 := rand.N(len(__obf_24fb59573a7563e5))
	return __obf_24fb59573a7563e5[__obf_fde7c1857dd05896]
}

func PickN[T any](__obf_601d2fcfa15aecc5 []T, __obf_11a540825a7c7a7e int) []T {
	if __obf_11a540825a7c7a7e <= 0 || len(__obf_601d2fcfa15aecc5) == 0 {
		return nil
	}
	if __obf_11a540825a7c7a7e >= len(__obf_601d2fcfa15aecc5) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_601d2fcfa15aecc5...)
	}

	// 拷贝一份，避免影响原 slice
	__obf_2f93bc7e73ababff := append([]T(nil), __obf_601d2fcfa15aecc5...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_2f93bc7e73ababff), func(__obf_0561496d0bd431c5, __obf_f516ab51ca117d18 int) {
		__obf_2f93bc7e73ababff[__obf_0561496d0bd431c5], __obf_2f93bc7e73ababff[__obf_f516ab51ca117d18] = __obf_2f93bc7e73ababff[__obf_f516ab51ca117d18], __obf_2f93bc7e73ababff[__obf_0561496d0bd431c5]
	})

	return __obf_2f93bc7e73ababff[:__obf_11a540825a7c7a7e]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
