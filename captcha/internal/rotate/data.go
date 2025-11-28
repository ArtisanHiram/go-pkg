package __obf_771d454808f56d40

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_b4468f4140748fac int
	__obf_67ce4ac668bc5d62 *types.PNGImage
	__obf_15e7c5e530990c10 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[int] = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_19428d8b430755cc *CaptchaData) GetData() int {
	return __obf_19428d8b430755cc.__obf_b4468f4140748fac
}

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_19428d8b430755cc *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_19428d8b430755cc.__obf_67ce4ac668bc5d62
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image as CaptchaImage interface
func (__obf_19428d8b430755cc *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_19428d8b430755cc.__obf_15e7c5e530990c10
}

func (__obf_19428d8b430755cc *CaptchaData) Type() types.CaptchaType {
	return types.RotateCaptchat
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_19428d8b430755cc *CaptchaData) Verify(__obf_27368830a60ccff1, __obf_4b1dea3f619379c2 int) bool {
	// Normalize angles to 0-360 range
	__obf_afa0927c0c9af521 := (__obf_27368830a60ccff1 + __obf_19428d8b430755cc.__obf_b4468f4140748fac) % 360
	if __obf_afa0927c0c9af521 < 0 {
		__obf_afa0927c0c9af521 += 360
	}

	// Check if total angle is close to 360 (or 0, which is equivalent)
	__obf_bd5393bb338ba5fe := 360 - __obf_4b1dea3f619379c2
	__obf_e0ffa5a903269019 := __obf_4b1dea3f619379c2

	return __obf_afa0927c0c9af521 >= __obf_bd5393bb338ba5fe || __obf_afa0927c0c9af521 <= __obf_e0ffa5a903269019
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
