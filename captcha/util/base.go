package __obf_1c2965545693a835

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
func RotatePoint(__obf_a63cad945e972fca, __obf_021ff6a91d01de34, __obf_6a9f0ab17a071c9a, __obf_bfe52e62046db6eb float64) (float64, float64) {
	return __obf_a63cad945e972fca*__obf_bfe52e62046db6eb - __obf_021ff6a91d01de34*__obf_6a9f0ab17a071c9a,

		// RotatedSize calculates the size after rotation
		__obf_a63cad945e972fca*__obf_6a9f0ab17a071c9a + __obf_021ff6a91d01de34*__obf_bfe52e62046db6eb
}

func RotatedSize(__obf_41b2dda21c56ae5b, __obf_732f9cfa854e40b8 int, __obf_eb4b86389b424112 float64) (int, int) {
	if __obf_41b2dda21c56ae5b <= 0 || __obf_732f9cfa854e40b8 <= 0 {
		return 0, 0
	}
	__obf_6a9f0ab17a071c9a, __obf_bfe52e62046db6eb := math.Sincos(math.Pi * __obf_eb4b86389b424112 / 180)
	__obf_58c9d87eca424963, __obf_381f3567196407f3 := RotatePoint(float64(__obf_41b2dda21c56ae5b-1), 0, __obf_6a9f0ab17a071c9a, __obf_bfe52e62046db6eb)
	__obf_6e9faff7fc1b26bc, __obf_969af149658cf800 := RotatePoint(float64(__obf_41b2dda21c56ae5b-1), float64(__obf_732f9cfa854e40b8-1), __obf_6a9f0ab17a071c9a, __obf_bfe52e62046db6eb)
	__obf_76d5e8339e68a1eb, __obf_d9bcf3367c0fe766 := RotatePoint(0, float64(__obf_732f9cfa854e40b8-1), __obf_6a9f0ab17a071c9a, __obf_bfe52e62046db6eb)
	__obf_9568495e64fecb9a := min(__obf_58c9d87eca424963, min(__obf_6e9faff7fc1b26bc, min(__obf_76d5e8339e68a1eb, 0)))
	__obf_3a3c24a586919612 := max(__obf_58c9d87eca424963, max(__obf_6e9faff7fc1b26bc, max(__obf_76d5e8339e68a1eb, 0)))
	__obf_7dcb0c90f2108495 := min(__obf_381f3567196407f3, min(__obf_969af149658cf800, min(__obf_d9bcf3367c0fe766, 0)))
	__obf_07770457a258d9ca := max(__obf_381f3567196407f3, max(__obf_969af149658cf800, max(__obf_d9bcf3367c0fe766, 0)))
	__obf_10992b4ba77e618e := __obf_3a3c24a586919612 - __obf_9568495e64fecb9a + 1
	if __obf_10992b4ba77e618e-math.Floor(__obf_10992b4ba77e618e) > 0.1 {
		__obf_10992b4ba77e618e++
	}
	__obf_adf312b75aaf0c4e := __obf_07770457a258d9ca - __obf_7dcb0c90f2108495 + 1
	if __obf_adf312b75aaf0c4e-math.Floor(__obf_adf312b75aaf0c4e) > 0.1 {
		__obf_adf312b75aaf0c4e++
	}

	return int(__obf_10992b4ba77e618e), int(__obf_adf312b75aaf0c4e)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_4671ef3b288bfa58 image.Rectangle, __obf_10992b4ba77e618e int, __obf_adf312b75aaf0c4e int, __obf_a649049b5a3072a5 bool) image.Rectangle {
	var __obf_31a4d7a3133e3d72 image.Rectangle
	if __obf_10992b4ba77e618e*__obf_4671ef3b288bfa58.Dy() < __obf_adf312b75aaf0c4e*__obf_4671ef3b288bfa58.Dx() {
		__obf_d3d8b7cf93438a10 := float64(__obf_10992b4ba77e618e) / float64(__obf_4671ef3b288bfa58.Dx())
		__obf_d83a3f89e1b7f449 := int(float64(__obf_4671ef3b288bfa58.Dy()) * __obf_d3d8b7cf93438a10)
		__obf_6727410c0077ddc7 := 0
		if __obf_a649049b5a3072a5 {
			__obf_6727410c0077ddc7 = (__obf_adf312b75aaf0c4e - __obf_d83a3f89e1b7f449) / 2
		}
		__obf_31a4d7a3133e3d72 = image.Rect(0, __obf_6727410c0077ddc7, __obf_10992b4ba77e618e, __obf_6727410c0077ddc7+__obf_d83a3f89e1b7f449)
	} else {
		__obf_d3d8b7cf93438a10 := float64(__obf_adf312b75aaf0c4e) / float64(__obf_4671ef3b288bfa58.Dy())
		__obf_6efec5daf220a986 := int(float64(__obf_4671ef3b288bfa58.Dx()) * __obf_d3d8b7cf93438a10)
		__obf_6727410c0077ddc7 := 0
		if __obf_a649049b5a3072a5 {
			__obf_6727410c0077ddc7 = (__obf_10992b4ba77e618e - __obf_6efec5daf220a986) / 2
		}
		__obf_31a4d7a3133e3d72 = image.Rect(__obf_6727410c0077ddc7, 0, __obf_6727410c0077ddc7+__obf_6efec5daf220a986, __obf_adf312b75aaf0c4e)
	}

	return __obf_31a4d7a3133e3d72
}

// t2x converts an integer to a hexadecimal string
func __obf_c5f154f9eea9c774(__obf_5c4f633d83d5171c int64) string {
	__obf_ffe10bbfa0be85de := strconv.FormatInt(__obf_5c4f633d83d5171c, 16)
	if len(__obf_ffe10bbfa0be85de) == 1 {
		__obf_ffe10bbfa0be85de = "0" + __obf_ffe10bbfa0be85de
	}
	return __obf_ffe10bbfa0be85de
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_dadb5bae15938891 float32) uint8 {
	__obf_b9a92f4e2c80ee3c := min(float64(__obf_dadb5bae15938891), 1)
	__obf_c42c4618284a204a := __obf_b9a92f4e2c80ee3c * 255
	return uint8(__obf_c42c4618284a204a)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_af4af93c3e797ce7 int64, __obf_32a3d57ef91fb356 int64, __obf_3060f4edc961d1a9 int64) string {
	__obf_7a0b47927c3bd3b0 := __obf_c5f154f9eea9c774(__obf_af4af93c3e797ce7)
	__obf_2bfc60892655e239 := __obf_c5f154f9eea9c774(__obf_32a3d57ef91fb356)
	__obf_d935b3348e7c0cef := __obf_c5f154f9eea9c774(__obf_3060f4edc961d1a9)
	return __obf_7a0b47927c3bd3b0 + __obf_2bfc60892655e239 + __obf_d935b3348e7c0cef
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_776fe5eb3ae78185 string) (int64, int64, int64) {
	__obf_7a0b47927c3bd3b0, _ := strconv.ParseInt(__obf_776fe5eb3ae78185[:2], 16, 10)
	__obf_2bfc60892655e239, _ := strconv.ParseInt(__obf_776fe5eb3ae78185[2:4], 16, 18)
	__obf_d935b3348e7c0cef, _ := strconv.ParseInt(__obf_776fe5eb3ae78185[4:], 16, 10)
	return __obf_7a0b47927c3bd3b0, __obf_2bfc60892655e239, __obf_d935b3348e7c0cef
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_e7ca3654bc7cd9ec string) (__obf_ff426c14f40340c1 color.RGBA, __obf_675a3caf504b632f error) {
	__obf_ff426c14f40340c1.
		A = 0xff
	if __obf_e7ca3654bc7cd9ec[0] != '#' {
		return __obf_ff426c14f40340c1, ColorHexFormatErr
	}
	__obf_145e87283d5c61cf := func(__obf_d935b3348e7c0cef byte) byte {
		switch {
		case __obf_d935b3348e7c0cef >= '0' && __obf_d935b3348e7c0cef <= '9':
			return __obf_d935b3348e7c0cef - '0'
		case __obf_d935b3348e7c0cef >= 'a' && __obf_d935b3348e7c0cef <= 'f':
			return __obf_d935b3348e7c0cef - 'a' + 10
		case __obf_d935b3348e7c0cef >= 'A' && __obf_d935b3348e7c0cef <= 'F':
			return __obf_d935b3348e7c0cef - 'A' + 10
		}
		__obf_675a3caf504b632f = ColorInvalidErr
		return 0
	}

	switch len(__obf_e7ca3654bc7cd9ec) {
	case 7:
		__obf_ff426c14f40340c1.
			R = __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[1])<<4 + __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[2])
		__obf_ff426c14f40340c1.
			G = __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[3])<<4 + __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[4])
		__obf_ff426c14f40340c1.
			B = __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[5])<<4 + __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[6])

	case 4:
		__obf_ff426c14f40340c1.
			R = __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[1]) * 17
		__obf_ff426c14f40340c1.
			G = __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[2]) * 17
		__obf_ff426c14f40340c1.
			B = __obf_145e87283d5c61cf(__obf_e7ca3654bc7cd9ec[3]) * 17
	default:
		__obf_675a3caf504b632f = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_5c264627f7ac7788 string) (bool, error) {
	_, __obf_675a3caf504b632f := os.Stat(__obf_5c264627f7ac7788)
	if __obf_675a3caf504b632f == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_675a3caf504b632f) {
		return false, nil
	}
	return false, __obf_675a3caf504b632f
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_d953bfdd181022ab string) bool {
	for _, __obf_7a0b47927c3bd3b0 := range __obf_d953bfdd181022ab {
		if unicode.Is(unicode.Scripts["Han"], __obf_7a0b47927c3bd3b0) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_d953bfdd181022ab string) int {
	return utf8.RuneCountInString(__obf_d953bfdd181022ab)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_6b26b669d3176a85 []color.Color) color.RGBA {
	__obf_c429d4e852f9aca2 := PickN(__obf_6b26b669d3176a85, 1)[0]
	__obf_7a0b47927c3bd3b0, __obf_2bfc60892655e239, __obf_d935b3348e7c0cef, __obf_b9a92f4e2c80ee3c := __obf_c429d4e852f9aca2.RGBA()
	return color.RGBA{R: uint8(__obf_7a0b47927c3bd3b0), G: uint8(__obf_2bfc60892655e239), B: uint8(__obf_d935b3348e7c0cef), A: uint8(__obf_b9a92f4e2c80ee3c)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_10992b4ba77e618e int, __obf_adf312b75aaf0c4e int, __obf_887e29ba67702921 image.Image) image.Point {
	__obf_d935b3348e7c0cef := __obf_887e29ba67702921.Bounds()
	__obf_2f14e3da9cff6d8b := __obf_d935b3348e7c0cef.Max.X
	__obf_2a5d46f1de50f6bc := __obf_d935b3348e7c0cef.Max.Y
	__obf_7ec94c3c601d1fcc := 0
	__obf_9948708df2478d45 := 0

	if __obf_2f14e3da9cff6d8b-__obf_10992b4ba77e618e > 0 {
		__obf_7ec94c3c601d1fcc = rand.IntN(__obf_2f14e3da9cff6d8b - __obf_10992b4ba77e618e)
	}
	if __obf_2a5d46f1de50f6bc-__obf_adf312b75aaf0c4e > 0 {
		__obf_9948708df2478d45 = rand.IntN(__obf_2a5d46f1de50f6bc - __obf_adf312b75aaf0c4e)
	}

	return image.Point{
		X: __obf_7ec94c3c601d1fcc,
		Y: __obf_9948708df2478d45,
	}
}

func RandString(__obf_a9363ce1ea8e03ee []string) string {
	__obf_6e0409791c98a836 := rand.N(len(__obf_a9363ce1ea8e03ee))
	return __obf_a9363ce1ea8e03ee[__obf_6e0409791c98a836]
}

func PickN[T any](__obf_2b31d7de3ecdf5d6 []T, __obf_cbba97241a66b5ac int) []T {
	if __obf_cbba97241a66b5ac <= 0 || len(__obf_2b31d7de3ecdf5d6) == 0 {
		return nil
	}
	if __obf_cbba97241a66b5ac >= len(__obf_2b31d7de3ecdf5d6) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_2b31d7de3ecdf5d6...)
	}
	__obf_9682734013a951b6 := // 拷贝一份，避免影响原 slice
		append([]T(nil), __obf_2b31d7de3ecdf5d6...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_9682734013a951b6), func(__obf_f8f33cc0f340b365, __obf_79f85012b8eff081 int) {
		__obf_9682734013a951b6[__obf_f8f33cc0f340b365], __obf_9682734013a951b6[__obf_79f85012b8eff081] = __obf_9682734013a951b6[__obf_79f85012b8eff081], __obf_9682734013a951b6[__obf_f8f33cc0f340b365]
	})

	return __obf_9682734013a951b6[:__obf_cbba97241a66b5ac]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
