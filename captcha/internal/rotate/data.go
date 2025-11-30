package __obf_4b873d45958741b9

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	Angle                  int `json:"angle"`
	__obf_ee7efadff09b92f2 *types.PNGImage
	__obf_c9f1e5ba3e028129 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_4f5780d7172901a0 *CaptchaData) GetData() any {
	return __obf_4f5780d7172901a0.Angle
}

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_4f5780d7172901a0 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_4f5780d7172901a0.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_ee7efadff09b92f2
}

func (__obf_4f5780d7172901a0 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_4f5780d7172901a0.__obf_c9f1e5ba3e028129
}

func (__obf_4f5780d7172901a0 *CaptchaData) GetType() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_4f5780d7172901a0 *CaptchaData) Verify(__obf_aa63c6273cec2b48 any, __obf_5f9ad8c424fc4a13 int) bool {
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

	var __obf_fde4510d2c9d2319 int
	// 使用类型选择来处理不同的输入类型
	switch __obf_c66ece1cbb82ef54 := __obf_aa63c6273cec2b48.(type) {
	case int:
		__obf_fde4510d2c9d2319 = __obf_c66ece1cbb82ef54
	case float64:
		__obf_fde4510d2c9d2319 = // JSON 默认解析为 float64，这里将其转回 int
			int(__obf_c66ece1cbb82ef54)
	case int64:
		__obf_fde4510d2c9d2319 = int(__obf_c66ece1cbb82ef54)
	case float32:
		__obf_fde4510d2c9d2319 = int(__obf_c66ece1cbb82ef54)
	case string:
		// 如果前端传的是字符串 "90"，可能需要 strconv.Atoi 解析，视需求而定
		return false
	default:
		// 类型不匹配，直接验证失败
		return false
	}
	__obf_9c97ab06692460d7 := // Normalize angles to 0-360 range
		(__obf_fde4510d2c9d2319 + __obf_4f5780d7172901a0.Angle) % 360
	if __obf_9c97ab06692460d7 < 0 {
		__obf_9c97ab06692460d7 += 360
	}
	__obf_139c26df7cd97d9b := // Check if total angle is close to 360 (or 0, which is equivalent)
		360 - __obf_5f9ad8c424fc4a13
	__obf_3f79f7f9901e438f := __obf_5f9ad8c424fc4a13

	return __obf_9c97ab06692460d7 >= __obf_139c26df7cd97d9b || __obf_9c97ab06692460d7 <= __obf_3f79f7f9901e438f
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
