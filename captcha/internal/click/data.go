package __obf_256ada120a510dc5

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	Dots                   []*types.Dot `json:"dots"`
	__obf_e89b4f6927401041 *types.JPEGImage
	__obf_b1611b851b2fe0d1 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_9e882512dfecdb44 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_9e882512dfecdb44.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_e89b4f6927401041
}

func (__obf_9e882512dfecdb44 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_9e882512dfecdb44.

		// GetData gets the target dots data
		__obf_b1611b851b2fe0d1
}

func (__obf_9e882512dfecdb44 *CaptchaData) GetData() any {
	return __obf_9e882512dfecdb44.Dots
}

func (__obf_9e882512dfecdb44 *CaptchaData) GetType() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_9e882512dfecdb44 *CaptchaData) Verify(__obf_b7604b64365f62cc any, __obf_4c8ec0fe01819f30 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	var __obf_cebd8c694afe3f70 []*types.Dot
	__obf_5f65ae527777a8b2 := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_cebd8c694afe3f70,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_932e808d6397fd78, __obf_975b35bfae5ea6a3 := mapstructure.NewDecoder(__obf_5f65ae527777a8b2)
	if __obf_975b35bfae5ea6a3 != nil {
		return false
	}

	if __obf_975b35bfae5ea6a3 := __obf_932e808d6397fd78.Decode(__obf_b7604b64365f62cc); __obf_975b35bfae5ea6a3 != nil {
		return false
	}

	if len(__obf_cebd8c694afe3f70) != len(__obf_9e882512dfecdb44.Dots) {
		return false
	}
	for __obf_a527076966eacc74, __obf_ad5e523e38020659 := range __obf_9e882512dfecdb44.Dots {
		__obf_49ff4265c0ca0123 := // clickX := clicks[i*2]
			// clickY := clicks[i*2+1]
			__obf_cebd8c694afe3f70[__obf_a527076966eacc74].X
		__obf_db5a4a30b5d421f9 := __obf_cebd8c694afe3f70[__obf_a527076966eacc74].Y
		__obf_6aeb7d7c2cdae676 := // Calculate the center point of the dot
			__obf_ad5e523e38020659.X + __obf_ad5e523e38020659.Size/2
		__obf_081978ecfe11b206 := __obf_ad5e523e38020659.Y + __obf_ad5e523e38020659.Size/2

		// Check if click is within tolerance range
		if __obf_49ff4265c0ca0123 < __obf_6aeb7d7c2cdae676-__obf_4c8ec0fe01819f30 || __obf_49ff4265c0ca0123 > __obf_6aeb7d7c2cdae676+__obf_4c8ec0fe01819f30 || __obf_db5a4a30b5d421f9 < __obf_081978ecfe11b206-__obf_4c8ec0fe01819f30 || __obf_db5a4a30b5d421f9 > __obf_081978ecfe11b206+__obf_4c8ec0fe01819f30 {
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
