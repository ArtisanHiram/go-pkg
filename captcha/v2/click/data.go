package __obf_07af88af2ce1a9a4

import (
	core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"
)

// CaptData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_566e0cc4a480c2a3 []*Dot
	__obf_255287b28188217d *core.JPEGImage
	__obf_cc78efb9dd19cd20 *core.PNGImage
}

// var _ internal.Captcha = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image in JPEG format
func (__obf_660839016fa371e0 CaptchaData) GetPrimary() *core.JPEGImage {
	return __obf_660839016fa371e0.__obf_255287b28188217d
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image in PNG format
func (__obf_660839016fa371e0 CaptchaData) GetSecondary() *core.PNGImage {
	return __obf_660839016fa371e0.__obf_cc78efb9dd19cd20
}

func (__obf_660839016fa371e0 CaptchaData) GetData() []*Dot {
	return __obf_660839016fa371e0.__obf_566e0cc4a480c2a3
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates [x1, y1, x2, y2, ...]
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_660839016fa371e0 CaptchaData) Verify(__obf_d04b3b8027025626 []int, __obf_eab55a968cae7f80 int) bool {
	if len(__obf_d04b3b8027025626) != len(__obf_660839016fa371e0.__obf_566e0cc4a480c2a3)*2 {
		return false
	}

	for __obf_03cbab251fc402a6, __obf_95db2efe59be6d8d := range __obf_660839016fa371e0.__obf_566e0cc4a480c2a3 {
		__obf_9af655212eb1707e := __obf_d04b3b8027025626[__obf_03cbab251fc402a6*2]
		__obf_9010c5946a7a0ac6 := __obf_d04b3b8027025626[__obf_03cbab251fc402a6*2+1]

		// Calculate the center point of the dot
		__obf_0a4531ae2530777b := __obf_95db2efe59be6d8d.X + __obf_95db2efe59be6d8d.Size/2
		__obf_244f1a11676bb0e1 := __obf_95db2efe59be6d8d.Y + __obf_95db2efe59be6d8d.Size/2

		// Check if click is within tolerance range
		if __obf_9af655212eb1707e < __obf_0a4531ae2530777b-__obf_eab55a968cae7f80 || __obf_9af655212eb1707e > __obf_0a4531ae2530777b+__obf_eab55a968cae7f80 ||
			__obf_9010c5946a7a0ac6 < __obf_244f1a11676bb0e1-__obf_eab55a968cae7f80 || __obf_9010c5946a7a0ac6 > __obf_244f1a11676bb0e1+__obf_eab55a968cae7f80 {
			return false
		}
	}

	return true
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(sx, sy, dx, dy, width, height, tolerance int) bool {
// 	newWidth := width + (tolerance * 2)
// 	newHeight := height + (tolerance * 2)
// 	newDx := int(max(float64(dx), float64(dx-tolerance)))
// 	newDy := int(max(float64(dy), float64(dy-tolerance)))
//
// 	return sx >= newDx &&
// 		sx <= newDx+newWidth &&
// 		sy >= newDy &&
// 		sy <= newDy+newHeight
// }
