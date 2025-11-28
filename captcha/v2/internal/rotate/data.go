package __obf_f48dc973afdacd79

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_bbdf82463e548025 int
	__obf_3482409dfa579189 *types.PNGImage
	__obf_2ea6abf946c6d439 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[int] = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_36a78059fd35abb2 *CaptchaData) GetData() int {
	return __obf_36a78059fd35abb2.__obf_bbdf82463e548025
}

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_36a78059fd35abb2 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_36a78059fd35abb2.__obf_3482409dfa579189
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image as CaptchaImage interface
func (__obf_36a78059fd35abb2 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_36a78059fd35abb2.__obf_2ea6abf946c6d439
}

func (__obf_36a78059fd35abb2 *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_36a78059fd35abb2 *CaptchaData) Verify(__obf_1695d9ba7a62ed66, __obf_1f8290239f03ba33 int) bool {
	// Normalize angles to 0-360 range
	__obf_8b74e1d8a0d91f3b := (__obf_1695d9ba7a62ed66 + __obf_36a78059fd35abb2.__obf_bbdf82463e548025) % 360
	if __obf_8b74e1d8a0d91f3b < 0 {
		__obf_8b74e1d8a0d91f3b += 360
	}

	// Check if total angle is close to 360 (or 0, which is equivalent)
	__obf_fd64b5411842efde := 360 - __obf_1f8290239f03ba33
	__obf_793c5afcf3e45b46 := __obf_1f8290239f03ba33

	return __obf_8b74e1d8a0d91f3b >= __obf_fd64b5411842efde || __obf_8b74e1d8a0d91f3b <= __obf_793c5afcf3e45b46
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
