package __obf_a334e8eb146672ab

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_f7494f398db65194 []*types.Dot
	__obf_e6925b4ca56caa93 *types.JPEGImage
	__obf_131375f4a4b92e57 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_6fc320ffbfc77017 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_6fc320ffbfc77017.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_e6925b4ca56caa93
}

func (__obf_6fc320ffbfc77017 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_6fc320ffbfc77017.

		// GetData gets the target dots data
		__obf_131375f4a4b92e57
}

func (__obf_6fc320ffbfc77017 *CaptchaData) GetData() any {
	return __obf_6fc320ffbfc77017.__obf_f7494f398db65194
}

func (__obf_6fc320ffbfc77017 *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_6fc320ffbfc77017 *CaptchaData) Verify(__obf_ac8081ae068efefa any, __obf_fe79cc216ffe2700 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	var __obf_f7494f398db65194 []*types.Dot
	__obf_6d8ad1a6172069ed := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_f7494f398db65194,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_4132958ad98950b7, __obf_48243aced6a1b5c5 := mapstructure.NewDecoder(__obf_6d8ad1a6172069ed)
	if __obf_48243aced6a1b5c5 != nil {
		return false
	}

	if __obf_48243aced6a1b5c5 := __obf_4132958ad98950b7.Decode(__obf_ac8081ae068efefa); __obf_48243aced6a1b5c5 != nil {
		return false
	}

	if len(__obf_f7494f398db65194) != len(__obf_6fc320ffbfc77017.__obf_f7494f398db65194) {
		return false
	}
	for __obf_1f520fc8bd1d0fba, __obf_9bab71876408a0b2 := range __obf_6fc320ffbfc77017.
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_f7494f398db65194 {
		__obf_d6a12cb10e528bba := __obf_f7494f398db65194[__obf_1f520fc8bd1d0fba].X
		__obf_69c6d1d8613dc664 := __obf_f7494f398db65194[__obf_1f520fc8bd1d0fba].Y
		__obf_e114242f5fc5f206 := // Calculate the center point of the dot
			__obf_9bab71876408a0b2.X + __obf_9bab71876408a0b2.Size/2
		__obf_dd709b15b01bdcd5 := __obf_9bab71876408a0b2.Y + __obf_9bab71876408a0b2.Size/2

		// Check if click is within tolerance range
		if __obf_d6a12cb10e528bba < __obf_e114242f5fc5f206-__obf_fe79cc216ffe2700 || __obf_d6a12cb10e528bba > __obf_e114242f5fc5f206+__obf_fe79cc216ffe2700 || __obf_69c6d1d8613dc664 < __obf_dd709b15b01bdcd5-__obf_fe79cc216ffe2700 || __obf_69c6d1d8613dc664 > __obf_dd709b15b01bdcd5+__obf_fe79cc216ffe2700 {
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
