package __obf_4e766de6f7fc6549

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_720ba9ea19239003 int
	__obf_38820526d0b312f0 *types.PNGImage
	__obf_1e0e4b1d4069e3bb *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_598561cd40842fdf *CaptchaData) GetData() any {
	return __obf_598561cd40842fdf.

		// GetPrimary gets the main captcha image
		// return: Main image as CaptchaImage interface
		__obf_720ba9ea19239003
}

func (__obf_598561cd40842fdf *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_598561cd40842fdf.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_38820526d0b312f0
}

func (__obf_598561cd40842fdf *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_598561cd40842fdf.__obf_1e0e4b1d4069e3bb
}

func (__obf_598561cd40842fdf *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_598561cd40842fdf *CaptchaData) Verify(__obf_aafc95e91fb76432 any, __obf_a860fa928025ab8e int) bool {
	// if angle, ok := input.(int); ok {
	// 	// Normalize angles to 0-360 range
	// 	totalAngle := (angle + c.angle) % 360
	// 	if totalAngle < 0 {
	// 		totalAngle += 360
	// 	}

	// 	// Check if total angle is close to 360 (or 0, which is equivalent)
	// 	minAngle := 360 - tolerance
	// 	maxAngle := tolerance

	// 	return totalAngle >= minAngle || totalAngle <= maxAngle
	// }
	// return false

	var __obf_720ba9ea19239003 int
	// 使用类型选择来处理不同的输入类型
	switch __obf_0408e687767ef3e2 := __obf_aafc95e91fb76432.(type) {
	case int:
		__obf_720ba9ea19239003 = __obf_0408e687767ef3e2
	case float64:
		__obf_720ba9ea19239003 = // JSON 默认解析为 float64，这里将其转回 int
			int(__obf_0408e687767ef3e2)
	case int64:
		__obf_720ba9ea19239003 = int(__obf_0408e687767ef3e2)
	case float32:
		__obf_720ba9ea19239003 = int(__obf_0408e687767ef3e2)
	case string:
		// 如果前端传的是字符串 "90"，可能需要 strconv.Atoi 解析，视需求而定
		return false
	default:
		// 类型不匹配，直接验证失败
		return false
	}
	__obf_147f95e9ccb6e045 := // Normalize angles to 0-360 range
		(__obf_720ba9ea19239003 + __obf_598561cd40842fdf.__obf_720ba9ea19239003) % 360
	if __obf_147f95e9ccb6e045 < 0 {
		__obf_147f95e9ccb6e045 += 360
	}
	__obf_32ee2835ab7501cf := // Check if total angle is close to 360 (or 0, which is equivalent)
		360 - __obf_a860fa928025ab8e
	__obf_7836d7ea9062faff := __obf_a860fa928025ab8e

	return __obf_147f95e9ccb6e045 >= __obf_32ee2835ab7501cf || __obf_147f95e9ccb6e045 <= __obf_7836d7ea9062faff
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
