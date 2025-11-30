package __obf_19fe0d6ff470ddc2

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_dbb4db6cc7ddc16b int
	__obf_efaba3b6c7084e36 *types.PNGImage
	__obf_0475d73f67e757fd *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_0ce9a2192c36c384 *CaptchaData) GetData() any {
	return __obf_0ce9a2192c36c384.

		// GetPrimary gets the main captcha image
		// return: Main image as CaptchaImage interface
		__obf_dbb4db6cc7ddc16b
}

func (__obf_0ce9a2192c36c384 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_0ce9a2192c36c384.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_efaba3b6c7084e36
}

func (__obf_0ce9a2192c36c384 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_0ce9a2192c36c384.__obf_0475d73f67e757fd
}

func (__obf_0ce9a2192c36c384 *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_0ce9a2192c36c384 *CaptchaData) Verify(__obf_eb9d11d9f51b9a1b any, __obf_d1cee4e7b0565378 int) bool {
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

	var __obf_dbb4db6cc7ddc16b int
	// 使用类型选择来处理不同的输入类型
	switch __obf_df50c2c529c8754b := __obf_eb9d11d9f51b9a1b.(type) {
	case int:
		__obf_dbb4db6cc7ddc16b = __obf_df50c2c529c8754b
	case float64:
		__obf_dbb4db6cc7ddc16b = // JSON 默认解析为 float64，这里将其转回 int
			int(__obf_df50c2c529c8754b)
	case int64:
		__obf_dbb4db6cc7ddc16b = int(__obf_df50c2c529c8754b)
	case float32:
		__obf_dbb4db6cc7ddc16b = int(__obf_df50c2c529c8754b)
	case string:
		// 如果前端传的是字符串 "90"，可能需要 strconv.Atoi 解析，视需求而定
		return false
	default:
		// 类型不匹配，直接验证失败
		return false
	}
	__obf_e93085a404bc8caf := // Normalize angles to 0-360 range
		(__obf_dbb4db6cc7ddc16b + __obf_0ce9a2192c36c384.__obf_dbb4db6cc7ddc16b) % 360
	if __obf_e93085a404bc8caf < 0 {
		__obf_e93085a404bc8caf += 360
	}
	__obf_2460bfa7f373db9f := // Check if total angle is close to 360 (or 0, which is equivalent)
		360 - __obf_d1cee4e7b0565378
	__obf_99a805d6ef0b2852 := __obf_d1cee4e7b0565378

	return __obf_e93085a404bc8caf >= __obf_2460bfa7f373db9f || __obf_e93085a404bc8caf <= __obf_99a805d6ef0b2852
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
