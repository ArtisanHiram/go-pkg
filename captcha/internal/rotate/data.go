package __obf_220a18b254a8ad8c

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_1a66d3fac7d22235 int
	__obf_5a73f4d6e809a9e0 *types.PNGImage
	__obf_a37b058809dc2b58 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[int] = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_232060d34b1a5883 *CaptchaData) GetData() int {
	return __obf_232060d34b1a5883.__obf_1a66d3fac7d22235
}

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_232060d34b1a5883 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_232060d34b1a5883.__obf_5a73f4d6e809a9e0
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image as CaptchaImage interface
func (__obf_232060d34b1a5883 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_232060d34b1a5883.__obf_a37b058809dc2b58
}

func (__obf_232060d34b1a5883 *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_232060d34b1a5883 *CaptchaData) Verify(__obf_7e529b475c8965f3, __obf_265b0683a0bc4e8c int) bool {
	// Normalize angles to 0-360 range
	__obf_3c08d1e074684666 := (__obf_7e529b475c8965f3 + __obf_232060d34b1a5883.__obf_1a66d3fac7d22235) % 360
	if __obf_3c08d1e074684666 < 0 {
		__obf_3c08d1e074684666 += 360
	}

	// Check if total angle is close to 360 (or 0, which is equivalent)
	__obf_48d7cea7dc5d1aba := 360 - __obf_265b0683a0bc4e8c
	__obf_3324c2e1be921610 := __obf_265b0683a0bc4e8c

	return __obf_3c08d1e074684666 >= __obf_48d7cea7dc5d1aba || __obf_3c08d1e074684666 <= __obf_3324c2e1be921610
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
