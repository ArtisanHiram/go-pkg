package __obf_e017a70acec4eb68

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_f9c9a01651aaf15e int
	__obf_500547421fb3df93 *types.PNGImage
	__obf_a5b5318e5ebc7a2c *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_3e760783dbf49c5a *CaptchaData) GetData() any {
	return __obf_3e760783dbf49c5a.

		// GetPrimary gets the main captcha image
		// return: Main image as CaptchaImage interface
		__obf_f9c9a01651aaf15e
}

func (__obf_3e760783dbf49c5a *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_3e760783dbf49c5a.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_500547421fb3df93
}

func (__obf_3e760783dbf49c5a *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_3e760783dbf49c5a.__obf_a5b5318e5ebc7a2c
}

func (__obf_3e760783dbf49c5a *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_3e760783dbf49c5a *CaptchaData) Verify(__obf_32f39f706f0f9ccb any, __obf_239d5df3376ec43e int) bool {
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

	var __obf_f9c9a01651aaf15e int
	// 使用类型选择来处理不同的输入类型
	switch __obf_00c49d77ebb2cdc7 := __obf_32f39f706f0f9ccb.(type) {
	case int:
		__obf_f9c9a01651aaf15e = __obf_00c49d77ebb2cdc7
	case float64:
		__obf_f9c9a01651aaf15e = // JSON 默认解析为 float64，这里将其转回 int
			int(__obf_00c49d77ebb2cdc7)
	case int64:
		__obf_f9c9a01651aaf15e = int(__obf_00c49d77ebb2cdc7)
	case float32:
		__obf_f9c9a01651aaf15e = int(__obf_00c49d77ebb2cdc7)
	case string:
		// 如果前端传的是字符串 "90"，可能需要 strconv.Atoi 解析，视需求而定
		return false
	default:
		// 类型不匹配，直接验证失败
		return false
	}
	__obf_003799fa49b078a2 := // Normalize angles to 0-360 range
		(__obf_f9c9a01651aaf15e + __obf_3e760783dbf49c5a.__obf_f9c9a01651aaf15e) % 360
	if __obf_003799fa49b078a2 < 0 {
		__obf_003799fa49b078a2 += 360
	}
	__obf_69c906f734009166 := // Check if total angle is close to 360 (or 0, which is equivalent)
		360 - __obf_239d5df3376ec43e
	__obf_0970974a1725b5a4 := __obf_239d5df3376ec43e

	return __obf_003799fa49b078a2 >= __obf_69c906f734009166 || __obf_003799fa49b078a2 <= __obf_0970974a1725b5a4
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
