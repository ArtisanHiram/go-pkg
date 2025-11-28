package __obf_96c99f1db90b23fa

import types "github.com/ArtisanHiram/go-pkg/captcha/types"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_9ecfe5c0e9e57390 []*types.Dot
	__obf_d1b1b184cf543db3 *types.JPEGImage
	__obf_dedf9178334e9a06 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[[]*types.Dot] = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_3f5b5aab6d53fe90 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_3f5b5aab6d53fe90.__obf_d1b1b184cf543db3
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image as CaptchaImage interface
func (__obf_3f5b5aab6d53fe90 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_3f5b5aab6d53fe90.__obf_dedf9178334e9a06
}

// GetData gets the target dots data
func (__obf_3f5b5aab6d53fe90 *CaptchaData) GetData() []*types.Dot {
	return __obf_3f5b5aab6d53fe90.__obf_9ecfe5c0e9e57390
}

func (__obf_3f5b5aab6d53fe90 *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_3f5b5aab6d53fe90 *CaptchaData) Verify(__obf_2d727629b6dc5dfb []*types.Dot, __obf_77fd0cfb3db60892 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	if len(__obf_2d727629b6dc5dfb) != len(__obf_3f5b5aab6d53fe90.__obf_9ecfe5c0e9e57390) {
		return false
	}

	for __obf_176527c945a10eb5, __obf_68c032f360706780 := range __obf_3f5b5aab6d53fe90.__obf_9ecfe5c0e9e57390 {
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_c5d039033cfb967a := __obf_2d727629b6dc5dfb[__obf_176527c945a10eb5].X
		__obf_825f010539c63826 := __obf_2d727629b6dc5dfb[__obf_176527c945a10eb5].Y

		// Calculate the center point of the dot
		__obf_4c0edcf9e88aee9f := __obf_68c032f360706780.X + __obf_68c032f360706780.Size/2
		__obf_25f07cf3768e1222 := __obf_68c032f360706780.Y + __obf_68c032f360706780.Size/2

		// Check if click is within tolerance range
		if __obf_c5d039033cfb967a < __obf_4c0edcf9e88aee9f-__obf_77fd0cfb3db60892 || __obf_c5d039033cfb967a > __obf_4c0edcf9e88aee9f+__obf_77fd0cfb3db60892 ||
			__obf_825f010539c63826 < __obf_25f07cf3768e1222-__obf_77fd0cfb3db60892 || __obf_825f010539c63826 > __obf_25f07cf3768e1222+__obf_77fd0cfb3db60892 {
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
