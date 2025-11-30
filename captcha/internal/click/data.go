package __obf_59290109f8ce9c8f

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_7f68597a0d34a028 []*types.Dot
	__obf_f7894dfc7807a63d *types.JPEGImage
	__obf_feb6cb098a8949dd *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_e0f9d59d598c10d8 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_e0f9d59d598c10d8.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_f7894dfc7807a63d
}

func (__obf_e0f9d59d598c10d8 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_e0f9d59d598c10d8.

		// GetData gets the target dots data
		__obf_feb6cb098a8949dd
}

func (__obf_e0f9d59d598c10d8 *CaptchaData) GetData() any {
	return __obf_e0f9d59d598c10d8.__obf_7f68597a0d34a028
}

func (__obf_e0f9d59d598c10d8 *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_e0f9d59d598c10d8 *CaptchaData) Verify(__obf_b9964f514fedda4c any, __obf_96791cdf6f0abeec int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	var __obf_7f68597a0d34a028 []*types.Dot
	__obf_592faf099fbe3ec3 := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_7f68597a0d34a028,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_558a72f4577d6119, __obf_573fb4e2cef8adee := mapstructure.NewDecoder(__obf_592faf099fbe3ec3)
	if __obf_573fb4e2cef8adee != nil {
		return false
	}

	if __obf_573fb4e2cef8adee := __obf_558a72f4577d6119.Decode(__obf_b9964f514fedda4c); __obf_573fb4e2cef8adee != nil {
		return false
	}

	if len(__obf_7f68597a0d34a028) != len(__obf_e0f9d59d598c10d8.__obf_7f68597a0d34a028) {
		return false
	}
	for __obf_c599f3c3680a56f5, __obf_d179dd23c6d2014b := range __obf_e0f9d59d598c10d8.
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_7f68597a0d34a028 {
		__obf_2a33393e72d7cf17 := __obf_7f68597a0d34a028[__obf_c599f3c3680a56f5].X
		__obf_a8fc99032f1de08b := __obf_7f68597a0d34a028[__obf_c599f3c3680a56f5].Y
		__obf_147cc71c328277eb := // Calculate the center point of the dot
			__obf_d179dd23c6d2014b.X + __obf_d179dd23c6d2014b.Size/2
		__obf_a6736922d7fa7a7b := __obf_d179dd23c6d2014b.Y + __obf_d179dd23c6d2014b.Size/2

		// Check if click is within tolerance range
		if __obf_2a33393e72d7cf17 < __obf_147cc71c328277eb-__obf_96791cdf6f0abeec || __obf_2a33393e72d7cf17 > __obf_147cc71c328277eb+__obf_96791cdf6f0abeec || __obf_a8fc99032f1de08b < __obf_a6736922d7fa7a7b-__obf_96791cdf6f0abeec || __obf_a8fc99032f1de08b > __obf_a6736922d7fa7a7b+__obf_96791cdf6f0abeec {
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
