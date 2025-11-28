package __obf_b2344bb9128f17e6

import types "github.com/ArtisanHiram/go-pkg/captcha/types"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_8c02fb9797f8ce7a []*types.Dot
	__obf_7e08e5ae21de2666 *types.JPEGImage
	__obf_ca3de5ebdc66e8cf *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[[]*types.Dot] = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_a38f2b46ae5b8a61 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_a38f2b46ae5b8a61.__obf_7e08e5ae21de2666
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image as CaptchaImage interface
func (__obf_a38f2b46ae5b8a61 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_a38f2b46ae5b8a61.__obf_ca3de5ebdc66e8cf
}

// GetData gets the target dots data
func (__obf_a38f2b46ae5b8a61 *CaptchaData) GetData() []*types.Dot {
	return __obf_a38f2b46ae5b8a61.__obf_8c02fb9797f8ce7a
}

func (__obf_a38f2b46ae5b8a61 *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_a38f2b46ae5b8a61 *CaptchaData) Verify(__obf_eda1d32b18f59544 []*types.Dot, __obf_f8405c4e3a20c061 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	if len(__obf_eda1d32b18f59544) != len(__obf_a38f2b46ae5b8a61.__obf_8c02fb9797f8ce7a) {
		return false
	}

	for __obf_76ec865d65fb400e, __obf_24103e8f3612481d := range __obf_a38f2b46ae5b8a61.__obf_8c02fb9797f8ce7a {
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_9538715a5b840083 := __obf_eda1d32b18f59544[__obf_76ec865d65fb400e].X
		__obf_26db953e3ce28edb := __obf_eda1d32b18f59544[__obf_76ec865d65fb400e].Y

		// Calculate the center point of the dot
		__obf_7239f696c39b2834 := __obf_24103e8f3612481d.X + __obf_24103e8f3612481d.Size/2
		__obf_79e79031243e2d1f := __obf_24103e8f3612481d.Y + __obf_24103e8f3612481d.Size/2

		// Check if click is within tolerance range
		if __obf_9538715a5b840083 < __obf_7239f696c39b2834-__obf_f8405c4e3a20c061 || __obf_9538715a5b840083 > __obf_7239f696c39b2834+__obf_f8405c4e3a20c061 ||
			__obf_26db953e3ce28edb < __obf_79e79031243e2d1f-__obf_f8405c4e3a20c061 || __obf_26db953e3ce28edb > __obf_79e79031243e2d1f+__obf_f8405c4e3a20c061 {
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
