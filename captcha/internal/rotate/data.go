package __obf_f99ff5685a4c2bb7

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	msgpack "github.com/ArtisanHiram/go-pkg/msgpack"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_7deb1b8e84658d4f int
	__obf_a9d22809133c29eb *types.PNGImage
	__obf_97e26f7838d6d222 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the rotation angle
func (__obf_1982ac2195a4989f *CaptchaData) GetPayload() types.CaptchaPayload {
	__obf_bf30f1c4a17a706b := types.CaptchaPayload{
		Type: types.RotateCaptchat,
	}
	__obf_bf30f1c4a17a706b.
		Payload, _ = msgpack.Marshal(__obf_1982ac2195a4989f.__obf_7deb1b8e84658d4f)
	return __obf_bf30f1c4a17a706b
}

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_1982ac2195a4989f *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_1982ac2195a4989f.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_a9d22809133c29eb
}

func (__obf_1982ac2195a4989f *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_1982ac2195a4989f.__obf_97e26f7838d6d222
}
