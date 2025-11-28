package __obf_48d1d3ff050bec1d

import core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_168883ae03b7841e int
	__obf_1f9603774b9e6d25 *core.PNGImage
	__obf_ffe34c16f5089ec9 *core.PNGImage
}

// var _ internal.Captcha = (*CaptchaData)(nil)

// GetData gets the block data
func (__obf_1eb094def04182e9 CaptchaData) GetData() int {
	return __obf_1eb094def04182e9.__obf_168883ae03b7841e
}

// GetPrimary gets the main captcha image
// return: Main image in PNG format
func (__obf_1eb094def04182e9 CaptchaData) GetPrimary() *core.PNGImage {
	return __obf_1eb094def04182e9.__obf_1f9603774b9e6d25
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image in PNG format
func (__obf_1eb094def04182e9 CaptchaData) GetSecondary() *core.PNGImage {
	return __obf_1eb094def04182e9.__obf_ffe34c16f5089ec9
}

// Verify checks if the rotated angle is correct
// userAngle: the angle rotated by user (0-360 degrees)
// tolerance: allowed angle deviation in degrees
// Returns true if the combined angle (userAngle + original angle) is close to 360 degrees
func (__obf_1eb094def04182e9 CaptchaData) Verify(__obf_e65bfabada7a37be, __obf_a9f4fec4548ab2ce int) bool {
	// Normalize angles to 0-360 range
	__obf_36f40956216a7e24 := (__obf_e65bfabada7a37be + __obf_1eb094def04182e9.__obf_168883ae03b7841e) % 360
	if __obf_36f40956216a7e24 < 0 {
		__obf_36f40956216a7e24 += 360
	}

	// Check if total angle is close to 360 (or 0, which is equivalent)
	__obf_28e6ddec682ce3c8 := 360 - __obf_a9f4fec4548ab2ce
	__obf_eca3b6a47bc02493 := __obf_a9f4fec4548ab2ce

	return __obf_36f40956216a7e24 >= __obf_28e6ddec682ce3c8 || __obf_36f40956216a7e24 <= __obf_eca3b6a47bc02493
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(angle, tolerance int) bool {
// 	minAngle := 360 - tolerance
// 	maxAngle := 360 + tolerance
// 	angle += c.angle
//
// 	return angle >= minAngle && angle <= maxAngle
// }
