package __obf_4612facabc6519c1

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_3535230d19dfb383 *types.Point
	__obf_ee6bc09cdda2f9e2 *types.JPEGImage
	__obf_8c3d473aadb0f39a *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_7961f70af277ccd5 *CaptchaData) GetPayload() types.CaptchaPayload {
	__obf_6e23655b04c52e4c := types.CaptchaPayload{Type: types.SlideCaptchat}
	__obf_6e23655b04c52e4c.
		Payload, _ = msgpack.Marshal(__obf_7961f70af277ccd5.__obf_3535230d19dfb383)
	return __obf_6e23655b04c52e4c
}

// GetPrimary gets the main CAPTCHA image
// return: Main image as CaptchaImage interface
func (__obf_7961f70af277ccd5 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_7961f70af277ccd5.

		// GetSecondary gets the tile image
		// return: Tile image as CaptchaImage interface
		__obf_ee6bc09cdda2f9e2
}

func (__obf_7961f70af277ccd5 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_7961f70af277ccd5.__obf_8c3d473aadb0f39a
}
