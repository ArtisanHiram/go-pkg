package __obf_3a7793861edae94b

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_50419e50fd91db67 int
	__obf_860f866a1a7270fa *types.PNGImage
	__obf_ee543d5bbdc5b150 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_801f980d9ae99a74 *CaptchaData) GetData() any {
	return __obf_801f980d9ae99a74.

		// GetPrimary gets the main captcha image
		// return: Main image as CaptchaImage interface
		__obf_50419e50fd91db67
}

func (__obf_801f980d9ae99a74 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_801f980d9ae99a74.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_860f866a1a7270fa
}

func (__obf_801f980d9ae99a74 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_801f980d9ae99a74.__obf_ee543d5bbdc5b150
}

func (__obf_801f980d9ae99a74 *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_801f980d9ae99a74 *CaptchaData) Verify(__obf_990fe709b877dd34 any, __obf_dccbdb9cb4de4b38 int) bool {
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

	var __obf_50419e50fd91db67 int
	// 使用类型选择来处理不同的输入类型
	switch __obf_d809e0aa6d4ba555 := __obf_990fe709b877dd34.(type) {
	case int:
		__obf_50419e50fd91db67 = __obf_d809e0aa6d4ba555
	case float64:
		__obf_50419e50fd91db67 = // JSON 默认解析为 float64，这里将其转回 int
			int(__obf_d809e0aa6d4ba555)
	case int64:
		__obf_50419e50fd91db67 = int(__obf_d809e0aa6d4ba555)
	case float32:
		__obf_50419e50fd91db67 = int(__obf_d809e0aa6d4ba555)
	case string:
		// 如果前端传的是字符串 "90"，可能需要 strconv.Atoi 解析，视需求而定
		return false
	default:
		// 类型不匹配，直接验证失败
		return false
	}
	__obf_cdc510cb00fa21f6 := // Normalize angles to 0-360 range
		(__obf_50419e50fd91db67 + __obf_801f980d9ae99a74.__obf_50419e50fd91db67) % 360
	if __obf_cdc510cb00fa21f6 < 0 {
		__obf_cdc510cb00fa21f6 += 360
	}
	__obf_2dd5c353c43bd67f := // Check if total angle is close to 360 (or 0, which is equivalent)
		360 - __obf_dccbdb9cb4de4b38
	__obf_aa8160c337c70f26 := __obf_dccbdb9cb4de4b38

	return __obf_cdc510cb00fa21f6 >= __obf_2dd5c353c43bd67f || __obf_cdc510cb00fa21f6 <= __obf_aa8160c337c70f26
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
