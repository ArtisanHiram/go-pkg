package __obf_f075990603c6fd45

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_e0fb28c1ab249b67 []*types.Dot
	__obf_8eb85ea5aa7dca32 *types.JPEGImage
	__obf_5e5bc037f5f349e0 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetPrimary gets the main captcha image
// return: Main image as CaptchaImage interface
func (__obf_b3aecd711f830444 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_b3aecd711f830444.

		// GetSecondary gets the thumbnail image
		// return: Thumbnail image as CaptchaImage interface
		__obf_8eb85ea5aa7dca32
}

func (__obf_b3aecd711f830444 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_b3aecd711f830444.

		// GetData gets the target dots data
		__obf_5e5bc037f5f349e0
}

func (__obf_b3aecd711f830444 *CaptchaData) GetData() any {
	return __obf_b3aecd711f830444.__obf_e0fb28c1ab249b67
}

func (__obf_b3aecd711f830444 *CaptchaData) Type() types.CaptchaType {
	return types.ClickCaptchat
}

// Verify checks if the clicked coordinates match any of the target dots
// clicks: slice of clicked coordinates
// tolerance: allowed pixel deviation for matching
// Returns true if all target dots are clicked correctly in order
func (__obf_b3aecd711f830444 *CaptchaData) Verify(__obf_d32cfa3665d90729 any, __obf_ccce17755f09ae37 int) bool {
	// if len(clicks) != len(c.dots)*2 {
	// 	return false
	// }

	var __obf_e0fb28c1ab249b67 []*types.Dot
	__obf_65224995171d0c31 := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_e0fb28c1ab249b67,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_255ef17b714ee970, __obf_6090e69b7a7b925f := mapstructure.NewDecoder(__obf_65224995171d0c31)
	if __obf_6090e69b7a7b925f != nil {
		return false
	}

	if __obf_6090e69b7a7b925f := __obf_255ef17b714ee970.Decode(__obf_d32cfa3665d90729); __obf_6090e69b7a7b925f != nil {
		return false
	}

	if len(__obf_e0fb28c1ab249b67) != len(__obf_b3aecd711f830444.__obf_e0fb28c1ab249b67) {
		return false
	}
	for __obf_339c01a0f79fa1a4, __obf_c6ea6b2418e24fc3 := range __obf_b3aecd711f830444.
		// clickX := clicks[i*2]
		// clickY := clicks[i*2+1]
		__obf_e0fb28c1ab249b67 {
		__obf_bab260193d5ee7d1 := __obf_e0fb28c1ab249b67[__obf_339c01a0f79fa1a4].X
		__obf_f47f093400bade31 := __obf_e0fb28c1ab249b67[__obf_339c01a0f79fa1a4].Y
		__obf_2da36a0db21f0582 := // Calculate the center point of the dot
			__obf_c6ea6b2418e24fc3.X + __obf_c6ea6b2418e24fc3.Size/2
		__obf_ee833d82d1c436b2 := __obf_c6ea6b2418e24fc3.Y + __obf_c6ea6b2418e24fc3.Size/2

		// Check if click is within tolerance range
		if __obf_bab260193d5ee7d1 < __obf_2da36a0db21f0582-__obf_ccce17755f09ae37 || __obf_bab260193d5ee7d1 > __obf_2da36a0db21f0582+__obf_ccce17755f09ae37 || __obf_f47f093400bade31 < __obf_ee833d82d1c436b2-__obf_ccce17755f09ae37 || __obf_f47f093400bade31 > __obf_ee833d82d1c436b2+__obf_ccce17755f09ae37 {
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
