package __obf_eadd0fe376b13c04

import types "github.com/ArtisanHiram/go-pkg/captcha/v2/types"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_e83956f7a258747d SlideType
	__obf_39adca10bea7dfb2 *types.Block
	__obf_41c3b81b590a5d60 *types.JPEGImage
	__obf_60a3f278519f7fa6 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[types.Block] = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_d0da3002c024252c *CaptchaData) GetData() types.Block {
	return *__obf_d0da3002c024252c.__obf_39adca10bea7dfb2
}

// GetPrimary gets the main CAPTCHA image
// return: Main image as CaptchaImage interface
func (__obf_d0da3002c024252c *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_d0da3002c024252c.__obf_41c3b81b590a5d60
}

// GetSecondary gets the tile image
// return: Tile image as CaptchaImage interface
func (__obf_d0da3002c024252c *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_d0da3002c024252c.__obf_60a3f278519f7fa6
}

func (__obf_d0da3002c024252c *CaptchaData) Type() types.CaptchaType {
	if __obf_d0da3002c024252c.__obf_e83956f7a258747d == Move {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_d0da3002c024252c *CaptchaData) Verify(__obf_be648116785a6003 types.Block, __obf_bf4032f1e7500366 int) bool {
	// Calculate the acceptable range for X and Y coordinates
	__obf_f3aa965cf3b0f0e2 := __obf_d0da3002c024252c.__obf_39adca10bea7dfb2.X - __obf_bf4032f1e7500366
	__obf_f72c03a4e4e3df29 := __obf_d0da3002c024252c.__obf_39adca10bea7dfb2.X + __obf_bf4032f1e7500366
	__obf_aa8a61c6357e5c74 := __obf_d0da3002c024252c.__obf_39adca10bea7dfb2.Y - __obf_bf4032f1e7500366
	__obf_3c872e94461e2a59 := __obf_d0da3002c024252c.__obf_39adca10bea7dfb2.Y + __obf_bf4032f1e7500366

	return __obf_be648116785a6003.X >= __obf_f3aa965cf3b0f0e2 && __obf_be648116785a6003.X <= __obf_f72c03a4e4e3df29 &&
		__obf_be648116785a6003.Y >= __obf_aa8a61c6357e5c74 && __obf_be648116785a6003.Y <= __obf_3c872e94461e2a59
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
