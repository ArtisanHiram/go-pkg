package __obf_741aa76a5fc5f3d9

import types "github.com/ArtisanHiram/go-pkg/captcha/types"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_044d58016cdb062b SlideType
	__obf_9cac0044fa40322d *types.Block
	__obf_e80f5dd754e39a7c *types.JPEGImage
	__obf_580c5dd4e04d5aad *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[types.Block] = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_88d29df8b03b1a42 *CaptchaData) GetData() types.Block {
	return *__obf_88d29df8b03b1a42.__obf_9cac0044fa40322d
}

// GetPrimary gets the main CAPTCHA image
// return: Main image as CaptchaImage interface
func (__obf_88d29df8b03b1a42 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_88d29df8b03b1a42.__obf_e80f5dd754e39a7c
}

// GetSecondary gets the tile image
// return: Tile image as CaptchaImage interface
func (__obf_88d29df8b03b1a42 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_88d29df8b03b1a42.__obf_580c5dd4e04d5aad
}

func (__obf_88d29df8b03b1a42 *CaptchaData) Type() types.CaptchaType {
	if __obf_88d29df8b03b1a42.__obf_044d58016cdb062b == Move {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_88d29df8b03b1a42 *CaptchaData) Verify(__obf_85782e1dfbe3f009 types.Block, __obf_e2b066e0a69c563b int) bool {
	// Calculate the acceptable range for X and Y coordinates
	__obf_da510cbd651ea673 := __obf_88d29df8b03b1a42.__obf_9cac0044fa40322d.X - __obf_e2b066e0a69c563b
	__obf_e4ad64ecaa217088 := __obf_88d29df8b03b1a42.__obf_9cac0044fa40322d.X + __obf_e2b066e0a69c563b
	__obf_6adc3177b2496664 := __obf_88d29df8b03b1a42.__obf_9cac0044fa40322d.Y - __obf_e2b066e0a69c563b
	__obf_7e04e8546f0ecf0c := __obf_88d29df8b03b1a42.__obf_9cac0044fa40322d.Y + __obf_e2b066e0a69c563b

	return __obf_85782e1dfbe3f009.X >= __obf_da510cbd651ea673 && __obf_85782e1dfbe3f009.X <= __obf_e4ad64ecaa217088 &&
		__obf_85782e1dfbe3f009.Y >= __obf_6adc3177b2496664 && __obf_85782e1dfbe3f009.Y <= __obf_7e04e8546f0ecf0c
}

// VerifyOld is the original verify method (commented out)
// func (c CaptchaData) Verify(sx, sy, dx, dy, tolerance int) bool {
// 	newX := tolerance * 2
// 	newY := tolerance * 2
// 	newDx := dx - tolerance
// 	newDy := dy - tolerance
//
// 	return sx >= newDx &&
// 		sx <= newDx+newX &&
// 		sy >= newDy &&
// 		sy <= newDy+newY
// }
