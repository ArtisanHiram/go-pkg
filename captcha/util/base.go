package __obf_ae73ba41a74c0e0a

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
func RotatePoint(__obf_1e7aca640fa8a127, __obf_9472e53e17085ed6, __obf_dd801466ad3aff43, __obf_b0bbb28eb3310eb6 float64) (float64, float64) {
	return __obf_1e7aca640fa8a127*__obf_b0bbb28eb3310eb6 - __obf_9472e53e17085ed6*__obf_dd801466ad3aff43,

		// RotatedSize calculates the size after rotation
		__obf_1e7aca640fa8a127*__obf_dd801466ad3aff43 + __obf_9472e53e17085ed6*__obf_b0bbb28eb3310eb6
}

func RotatedSize(__obf_daf87b9962d5dffe, __obf_d73d2d0575c092c2 int, __obf_44c0c71814f2321e float64) (int, int) {
	if __obf_daf87b9962d5dffe <= 0 || __obf_d73d2d0575c092c2 <= 0 {
		return 0, 0
	}
	__obf_dd801466ad3aff43, __obf_b0bbb28eb3310eb6 := math.Sincos(math.Pi * __obf_44c0c71814f2321e / 180)
	__obf_ceea6a074568e4e6, __obf_21269765b352fdb4 := RotatePoint(float64(__obf_daf87b9962d5dffe-1), 0, __obf_dd801466ad3aff43, __obf_b0bbb28eb3310eb6)
	__obf_2ab1729a3b8c7308, __obf_f7872d0ac564fdc3 := RotatePoint(float64(__obf_daf87b9962d5dffe-1), float64(__obf_d73d2d0575c092c2-1), __obf_dd801466ad3aff43, __obf_b0bbb28eb3310eb6)
	__obf_e63239a5d94b1f0f, __obf_9c50e550bcc95529 := RotatePoint(0, float64(__obf_d73d2d0575c092c2-1), __obf_dd801466ad3aff43, __obf_b0bbb28eb3310eb6)
	__obf_b4e10384860323ac := min(__obf_ceea6a074568e4e6, min(__obf_2ab1729a3b8c7308, min(__obf_e63239a5d94b1f0f, 0)))
	__obf_584ee39805f0d32a := max(__obf_ceea6a074568e4e6, max(__obf_2ab1729a3b8c7308, max(__obf_e63239a5d94b1f0f, 0)))
	__obf_f7375259a9d63d2d := min(__obf_21269765b352fdb4, min(__obf_f7872d0ac564fdc3, min(__obf_9c50e550bcc95529, 0)))
	__obf_7ebba0eefb58f35c := max(__obf_21269765b352fdb4, max(__obf_f7872d0ac564fdc3, max(__obf_9c50e550bcc95529, 0)))
	__obf_0e0e0b0f3e5d1b35 := __obf_584ee39805f0d32a - __obf_b4e10384860323ac + 1
	if __obf_0e0e0b0f3e5d1b35-math.Floor(__obf_0e0e0b0f3e5d1b35) > 0.1 {
		__obf_0e0e0b0f3e5d1b35++
	}
	__obf_8c0dbff2bad66fbf := __obf_7ebba0eefb58f35c - __obf_f7375259a9d63d2d + 1
	if __obf_8c0dbff2bad66fbf-math.Floor(__obf_8c0dbff2bad66fbf) > 0.1 {
		__obf_8c0dbff2bad66fbf++
	}

	return int(__obf_0e0e0b0f3e5d1b35), int(__obf_8c0dbff2bad66fbf)
}

// CalcResizedRect calculates the resized rectangle
func CalcResizedRect(__obf_6315b388e6d4ab7f image.Rectangle, __obf_0e0e0b0f3e5d1b35 int, __obf_8c0dbff2bad66fbf int, __obf_03daf4a1cc030a09 bool) image.Rectangle {
	var __obf_d39fe539870bb6e8 image.Rectangle
	if __obf_0e0e0b0f3e5d1b35*__obf_6315b388e6d4ab7f.Dy() < __obf_8c0dbff2bad66fbf*__obf_6315b388e6d4ab7f.Dx() {
		__obf_91dc3dced2f47e7d := float64(__obf_0e0e0b0f3e5d1b35) / float64(__obf_6315b388e6d4ab7f.Dx())
		__obf_1e9662a8dbd3b5cb := int(float64(__obf_6315b388e6d4ab7f.Dy()) * __obf_91dc3dced2f47e7d)
		__obf_82efb922b7076819 := 0
		if __obf_03daf4a1cc030a09 {
			__obf_82efb922b7076819 = (__obf_8c0dbff2bad66fbf - __obf_1e9662a8dbd3b5cb) / 2
		}
		__obf_d39fe539870bb6e8 = image.Rect(0, __obf_82efb922b7076819, __obf_0e0e0b0f3e5d1b35, __obf_82efb922b7076819+__obf_1e9662a8dbd3b5cb)
	} else {
		__obf_91dc3dced2f47e7d := float64(__obf_8c0dbff2bad66fbf) / float64(__obf_6315b388e6d4ab7f.Dy())
		__obf_f53762347353dec9 := int(float64(__obf_6315b388e6d4ab7f.Dx()) * __obf_91dc3dced2f47e7d)
		__obf_82efb922b7076819 := 0
		if __obf_03daf4a1cc030a09 {
			__obf_82efb922b7076819 = (__obf_0e0e0b0f3e5d1b35 - __obf_f53762347353dec9) / 2
		}
		__obf_d39fe539870bb6e8 = image.Rect(__obf_82efb922b7076819, 0, __obf_82efb922b7076819+__obf_f53762347353dec9, __obf_8c0dbff2bad66fbf)
	}

	return __obf_d39fe539870bb6e8
}

// t2x converts an integer to a hexadecimal string
func __obf_10aa416d7d78a1a8(__obf_ad3c9cf5dc8b42e7 int64) string {
	__obf_f2474f6642891165 := strconv.FormatInt(__obf_ad3c9cf5dc8b42e7, 16)
	if len(__obf_f2474f6642891165) == 1 {
		__obf_f2474f6642891165 = "0" + __obf_f2474f6642891165
	}
	return __obf_f2474f6642891165
}

// FormatAlpha formats the alpha value
func FormatAlpha(__obf_796b145e6f65ed1d float32) uint8 {
	__obf_543df975c05c4296 := min(float64(__obf_796b145e6f65ed1d), 1)
	__obf_1fdd29bee4214c67 := __obf_543df975c05c4296 * 255
	return uint8(__obf_1fdd29bee4214c67)
}

// RgbToHex converts RGB color to hexadecimal color
func RgbToHex(__obf_0294ad259006945e int64, __obf_0931d81e25aa4762 int64, __obf_7b94267b6d816e5d int64) string {
	__obf_11e154f2ed640f63 := __obf_10aa416d7d78a1a8(__obf_0294ad259006945e)
	__obf_e6f7330c29a28e29 := __obf_10aa416d7d78a1a8(__obf_0931d81e25aa4762)
	__obf_43ab9cb8031158e8 := __obf_10aa416d7d78a1a8(__obf_7b94267b6d816e5d)
	return __obf_11e154f2ed640f63 + __obf_e6f7330c29a28e29 + __obf_43ab9cb8031158e8
}

// HexToRgb converts hexadecimal color to RGB color
func HexToRgb(__obf_76b8994cc1472c1d string) (int64, int64, int64) {
	__obf_11e154f2ed640f63, _ := strconv.ParseInt(__obf_76b8994cc1472c1d[:2], 16, 10)
	__obf_e6f7330c29a28e29, _ := strconv.ParseInt(__obf_76b8994cc1472c1d[2:4], 16, 18)
	__obf_43ab9cb8031158e8, _ := strconv.ParseInt(__obf_76b8994cc1472c1d[4:], 16, 10)
	return __obf_11e154f2ed640f63, __obf_e6f7330c29a28e29, __obf_43ab9cb8031158e8
}

var (
	ColorHexFormatErr = errors.New("hex color must start with '#'")
	ColorInvalidErr   = errors.New("hexToByte component invalid")
)

// ParseHexColor converts a hex color to an RGBA color
func ParseHexColor(__obf_684540358834c785 string) (__obf_a6b4a34bf64f7486 color.RGBA, __obf_480cd6251cefe029 error) {
	__obf_a6b4a34bf64f7486.
		A = 0xff
	if __obf_684540358834c785[0] != '#' {
		return __obf_a6b4a34bf64f7486, ColorHexFormatErr
	}
	__obf_3a4f5fd72e328ff5 := func(__obf_43ab9cb8031158e8 byte) byte {
		switch {
		case __obf_43ab9cb8031158e8 >= '0' && __obf_43ab9cb8031158e8 <= '9':
			return __obf_43ab9cb8031158e8 - '0'
		case __obf_43ab9cb8031158e8 >= 'a' && __obf_43ab9cb8031158e8 <= 'f':
			return __obf_43ab9cb8031158e8 - 'a' + 10
		case __obf_43ab9cb8031158e8 >= 'A' && __obf_43ab9cb8031158e8 <= 'F':
			return __obf_43ab9cb8031158e8 - 'A' + 10
		}
		__obf_480cd6251cefe029 = ColorInvalidErr
		return 0
	}

	switch len(__obf_684540358834c785) {
	case 7:
		__obf_a6b4a34bf64f7486.
			R = __obf_3a4f5fd72e328ff5(__obf_684540358834c785[1])<<4 + __obf_3a4f5fd72e328ff5(__obf_684540358834c785[2])
		__obf_a6b4a34bf64f7486.
			G = __obf_3a4f5fd72e328ff5(__obf_684540358834c785[3])<<4 + __obf_3a4f5fd72e328ff5(__obf_684540358834c785[4])
		__obf_a6b4a34bf64f7486.
			B = __obf_3a4f5fd72e328ff5(__obf_684540358834c785[5])<<4 + __obf_3a4f5fd72e328ff5(__obf_684540358834c785[6])

	case 4:
		__obf_a6b4a34bf64f7486.
			R = __obf_3a4f5fd72e328ff5(__obf_684540358834c785[1]) * 17
		__obf_a6b4a34bf64f7486.
			G = __obf_3a4f5fd72e328ff5(__obf_684540358834c785[2]) * 17
		__obf_a6b4a34bf64f7486.
			B = __obf_3a4f5fd72e328ff5(__obf_684540358834c785[3]) * 17
	default:
		__obf_480cd6251cefe029 = ColorInvalidErr
	}
	return
}

// PathExists checks if a file path exists
func PathExists(__obf_b912d21f1390e44d string) (bool, error) {
	_, __obf_480cd6251cefe029 := os.Stat(__obf_b912d21f1390e44d)
	if __obf_480cd6251cefe029 == nil {
		return true, nil
	}
	if os.IsNotExist(__obf_480cd6251cefe029) {
		return false, nil
	}
	return false, __obf_480cd6251cefe029
}

// IsChineseChar checks if a string contains Chinese characters
func IsChineseChar(__obf_cc85b4c6ee121074 string) bool {
	for _, __obf_11e154f2ed640f63 := range __obf_cc85b4c6ee121074 {
		if unicode.Is(unicode.Scripts["Han"], __obf_11e154f2ed640f63) {
			return true
		}
	}
	return false
}

// LenChineseChar calculates the character length of a string (supports Chinese)
func LenChineseChar(__obf_cc85b4c6ee121074 string) int {
	return utf8.RuneCountInString(__obf_cc85b4c6ee121074)
}

// RandColor randomly selects an RGBA color
func RandColor(__obf_a4bca4b02e08712c []color.Color) color.RGBA {
	__obf_874d57d8db3d6b3f := PickN(__obf_a4bca4b02e08712c, 1)[0]
	__obf_11e154f2ed640f63, __obf_e6f7330c29a28e29, __obf_43ab9cb8031158e8, __obf_543df975c05c4296 := __obf_874d57d8db3d6b3f.RGBA()
	return color.RGBA{R: uint8(__obf_11e154f2ed640f63), G: uint8(__obf_e6f7330c29a28e29), B: uint8(__obf_43ab9cb8031158e8), A: uint8(__obf_543df975c05c4296)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(__obf_0e0e0b0f3e5d1b35 int, __obf_8c0dbff2bad66fbf int, __obf_49505236572de8ef image.Image) image.Point {
	__obf_43ab9cb8031158e8 := __obf_49505236572de8ef.Bounds()
	__obf_17b1442fceff22d3 := __obf_43ab9cb8031158e8.Max.X
	__obf_b560e06da7435991 := __obf_43ab9cb8031158e8.Max.Y
	__obf_ea2b5279750a41f3 := 0
	__obf_8bb593d666a097c7 := 0

	if __obf_17b1442fceff22d3-__obf_0e0e0b0f3e5d1b35 > 0 {
		__obf_ea2b5279750a41f3 = rand.IntN(__obf_17b1442fceff22d3 - __obf_0e0e0b0f3e5d1b35)
	}
	if __obf_b560e06da7435991-__obf_8c0dbff2bad66fbf > 0 {
		__obf_8bb593d666a097c7 = rand.IntN(__obf_b560e06da7435991 - __obf_8c0dbff2bad66fbf)
	}

	return image.Point{
		X: __obf_ea2b5279750a41f3,
		Y: __obf_8bb593d666a097c7,
	}
}

func RandString(__obf_df9818c4a1b529d9 []string) string {
	__obf_7cffc4a45ba51dca := rand.N(len(__obf_df9818c4a1b529d9))
	return __obf_df9818c4a1b529d9[__obf_7cffc4a45ba51dca]
}

func PickN[T any](__obf_62e55c636efc476a []T, __obf_3cc0ada1765d4af4 int) []T {
	if __obf_3cc0ada1765d4af4 <= 0 || len(__obf_62e55c636efc476a) == 0 {
		return nil
	}
	if __obf_3cc0ada1765d4af4 >= len(__obf_62e55c636efc476a) {
		// 返回一个副本，保证不修改原数组
		return append([]T(nil), __obf_62e55c636efc476a...)
	}
	__obf_a527d57362781e1b := // 拷贝一份，避免影响原 slice
		append([]T(nil), __obf_62e55c636efc476a...)

	// Fisher-Yates shuffle
	rand.Shuffle(len(__obf_a527d57362781e1b), func(__obf_43f455229137287d, __obf_7caa0e0609811d5b int) {
		__obf_a527d57362781e1b[__obf_43f455229137287d], __obf_a527d57362781e1b[__obf_7caa0e0609811d5b] = __obf_a527d57362781e1b[__obf_7caa0e0609811d5b], __obf_a527d57362781e1b[__obf_43f455229137287d]
	})

	return __obf_a527d57362781e1b[:__obf_3cc0ada1765d4af4]
}

// RandInt generates a safe random number in the interval [min, max]
func RandInt(min, max int) int {
	if min > max {
		return max
	}
	return min + rand.IntN(max-min+1)
}
