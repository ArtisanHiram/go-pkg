package __obf_7b43310cbd758abe

import types "github.com/ArtisanHiram/go-pkg/captcha/types"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_782e73e3bc2a2b90 SlideType
	__obf_48a33b6401bf2a46 *types.Block
	__obf_ae9d075a50d2eeb2 *types.JPEGImage
	__obf_91551e891ceb13f1 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.Captcha[types.Block] = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_533af04904ee7f62 *CaptchaData) GetData() types.Block {
	return *__obf_533af04904ee7f62.__obf_48a33b6401bf2a46
}

// GetPrimary gets the main CAPTCHA image
// return: Main image as CaptchaImage interface
func (__obf_533af04904ee7f62 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_533af04904ee7f62.__obf_ae9d075a50d2eeb2
}

// GetSecondary gets the tile image
// return: Tile image as CaptchaImage interface
func (__obf_533af04904ee7f62 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_533af04904ee7f62.__obf_91551e891ceb13f1
}

func (__obf_533af04904ee7f62 *CaptchaData) Type() types.CaptchaType {
	if __obf_533af04904ee7f62.__obf_782e73e3bc2a2b90 == Move {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_533af04904ee7f62 *CaptchaData) Verify(__obf_c0c774c4f0a79b4e types.Block, __obf_455fc6d01040cbfe int) bool {
	// Calculate the acceptable range for X and Y coordinates
	__obf_7a4e7ae5e4939fe0 := __obf_533af04904ee7f62.__obf_48a33b6401bf2a46.X - __obf_455fc6d01040cbfe
	__obf_d999daee922c940c := __obf_533af04904ee7f62.__obf_48a33b6401bf2a46.X + __obf_455fc6d01040cbfe
	__obf_30ca7e04288286c9 := __obf_533af04904ee7f62.__obf_48a33b6401bf2a46.Y - __obf_455fc6d01040cbfe
	__obf_360445ad54bb9f58 := __obf_533af04904ee7f62.__obf_48a33b6401bf2a46.Y + __obf_455fc6d01040cbfe

	return __obf_c0c774c4f0a79b4e.X >= __obf_7a4e7ae5e4939fe0 && __obf_c0c774c4f0a79b4e.X <= __obf_d999daee922c940c &&
		__obf_c0c774c4f0a79b4e.Y >= __obf_30ca7e04288286c9 && __obf_c0c774c4f0a79b4e.Y <= __obf_360445ad54bb9f58
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
