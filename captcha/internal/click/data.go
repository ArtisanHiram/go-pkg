package __obf_767326b1078782af

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_2028bb55834c38d8 []*types.Dot
	__obf_4292b5547fe8df15 *types.JPEGImage
	__obf_6871669979053980 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_de65322a0d55e5bb *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_de65322a0d55e5bb.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_4292b5547fe8df15
}

func (__obf_de65322a0d55e5bb *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_de65322a0d55e5bb.__obf_6871669979053980
}

func (__obf_de65322a0d55e5bb *CaptchaData) GetPayload() types.CaptchaPayload {
	__obf_97b2e25741090381 := types.CaptchaPayload{
		Type: types.ClickCaptchat,
	}
	__obf_97b2e25741090381.
		Payload, _ = msgpack.Marshal(__obf_de65322a0d55e5bb.__obf_2028bb55834c38d8)
	return __obf_97b2e25741090381
}
