package __obf_a0c8110fd085454e

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_cea1e242f930c716 []*types.Dot
	__obf_d88cc15f0bafbf24 *types.JPEGImage
	__obf_30d550091f68f974 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_761b3f820674d2e3 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_761b3f820674d2e3.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_d88cc15f0bafbf24
}

func (__obf_761b3f820674d2e3 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_761b3f820674d2e3.

		// GetData gets the target dots data
		__obf_30d550091f68f974
}

func (__obf_761b3f820674d2e3 *CaptchaData) GetData() any {
	return __obf_761b3f820674d2e3.__obf_cea1e242f930c716
}

func (__obf_761b3f820674d2e3 *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_761b3f820674d2e3 *CaptchaData) Verify(__obf_e32e2dedaa5096ca any, __obf_7161ea62e78e23e3 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	var __obf_cea1e242f930c716 []*types.Dot
	__obf_e1cc2008a5f56072 := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_cea1e242f930c716,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_b0c425150960b60d, __obf_43790225c3b3432f := mapstructure.NewDecoder(__obf_e1cc2008a5f56072)
	if __obf_43790225c3b3432f != nil {
		return false
	}

	if __obf_43790225c3b3432f := __obf_b0c425150960b60d.Decode(__obf_e32e2dedaa5096ca); __obf_43790225c3b3432f != nil {
		return false
	}

	if len(__obf_cea1e242f930c716) != len(__obf_761b3f820674d2e3.__obf_cea1e242f930c716) {
		return false
	}
	for __obf_d9fed1663089e979, __obf_b4922ee002e08e21 := range __obf_761b3f820674d2e3.
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_cea1e242f930c716 {
		__obf_3544ec82ee713232 := __obf_cea1e242f930c716[__obf_d9fed1663089e979].X
		__obf_ec1ce5a82d0b9af3 := __obf_cea1e242f930c716[__obf_d9fed1663089e979].Y
		__obf_91d919a815d9189c := // Calculate the center point of the dot
			__obf_b4922ee002e08e21.X + __obf_b4922ee002e08e21.Size/2
		__obf_a25d5405a298af9b := __obf_b4922ee002e08e21.Y + __obf_b4922ee002e08e21.Size/2

		// Check if click is within tolerance range
		if __obf_3544ec82ee713232 < __obf_91d919a815d9189c-__obf_7161ea62e78e23e3 || __obf_3544ec82ee713232 > __obf_91d919a815d9189c+__obf_7161ea62e78e23e3 || __obf_ec1ce5a82d0b9af3 < __obf_a25d5405a298af9b-__obf_7161ea62e78e23e3 || __obf_ec1ce5a82d0b9af3 > __obf_a25d5405a298af9b+__obf_7161ea62e78e23e3 {
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
