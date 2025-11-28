package __obf_25fa2a83bb367991

import types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_8ad847a1c4d1bb62 []*types.Dot
	__obf_ff8c367564ca56d3 *types.JPEGImage
	__obf_4a1046d8157ce073 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[[]*types.Dot] = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_d5c9aa819977350f *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_d5c9aa819977350f.__obf_ff8c367564ca56d3
}

// GetSecondary gets the thumbnail image
// return: Thumbnail image as CaptchaImage interface
func (__obf_d5c9aa819977350f *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_d5c9aa819977350f.__obf_4a1046d8157ce073
}

// GetData gets the target dots data
func (__obf_d5c9aa819977350f *CaptchaData) GetData() []*types.Dot {
	return __obf_d5c9aa819977350f.__obf_8ad847a1c4d1bb62
}

func (__obf_d5c9aa819977350f *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_d5c9aa819977350f *CaptchaData) Verify(__obf_63471fc8f384d517 []*types.Dot, __obf_74182e2a43aebef3 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	if len(__obf_63471fc8f384d517) != len(__obf_d5c9aa819977350f.__obf_8ad847a1c4d1bb62) {
		return false
	}

	for __obf_0063e9b84dd66db9, __obf_c8b7e75042fa4c1c := range __obf_d5c9aa819977350f.__obf_8ad847a1c4d1bb62 {
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_6ab9d251c0804ebd := __obf_63471fc8f384d517[__obf_0063e9b84dd66db9].X
		__obf_2cb59821e9f8487d := __obf_63471fc8f384d517[__obf_0063e9b84dd66db9].Y

		// Calculate the center point of the dot
		__obf_93694fe3b23ab0d2 := __obf_c8b7e75042fa4c1c.X + __obf_c8b7e75042fa4c1c.Size/2
		__obf_e761cf06984abbfc := __obf_c8b7e75042fa4c1c.Y + __obf_c8b7e75042fa4c1c.Size/2

		// Check if click is within tolerance range
		if __obf_6ab9d251c0804ebd < __obf_93694fe3b23ab0d2-__obf_74182e2a43aebef3 || __obf_6ab9d251c0804ebd > __obf_93694fe3b23ab0d2+__obf_74182e2a43aebef3 ||
			__obf_2cb59821e9f8487d < __obf_e761cf06984abbfc-__obf_74182e2a43aebef3 || __obf_2cb59821e9f8487d > __obf_e761cf06984abbfc+__obf_74182e2a43aebef3 {
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
