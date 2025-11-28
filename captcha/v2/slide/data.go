package __obf_40cc6e54bc08fb7a

import core "github.com/ArtisanHiram/go-pkg/captcha/v2/core"

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_d39062b0036fd87e *Block
	__obf_162c7157de6fcfa9 *core.JPEGImage
	__obf_c4b0de4f376b141e *core.PNGImage
}

// var _ internal.Captcha = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: Pointer to block data
func (__obf_4504122ced09d7a1 CaptchaData) GetData() *Block {
	return __obf_4504122ced09d7a1.__obf_d39062b0036fd87e
}

// GetPrimary gets the main CAPTCHA image
// return: Main image in JPEG format
func (__obf_4504122ced09d7a1 CaptchaData) GetPrimary() *core.JPEGImage {
	return __obf_4504122ced09d7a1.__obf_162c7157de6fcfa9
}

// GetSecondary gets the tile image
// return: Tile image in PNG format
func (__obf_4504122ced09d7a1 CaptchaData) GetSecondary() *core.PNGImage {
	return __obf_4504122ced09d7a1.__obf_c4b0de4f376b141e
}

// Verify checks if the slide position matches the block position
// userX: the X coordinate where user slid the block to
// userY: the Y coordinate where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_4504122ced09d7a1 CaptchaData) Verify(__obf_dd8ccf8e1e49adba core.Point, __obf_d9bfc8a08a9afb13 int) bool {
	// Calculate the acceptable range for X and Y coordinates
	__obf_722aa94ec7ba1000 := __obf_4504122ced09d7a1.__obf_d39062b0036fd87e.X - __obf_d9bfc8a08a9afb13
	__obf_2cc00eea4019cb5f := __obf_4504122ced09d7a1.__obf_d39062b0036fd87e.X + __obf_d9bfc8a08a9afb13
	__obf_b9d5f0fcf2bc49a3 := __obf_4504122ced09d7a1.__obf_d39062b0036fd87e.Y - __obf_d9bfc8a08a9afb13
	__obf_10fe878eb0a3e8ba := __obf_4504122ced09d7a1.__obf_d39062b0036fd87e.Y + __obf_d9bfc8a08a9afb13

	return __obf_dd8ccf8e1e49adba.X >= __obf_722aa94ec7ba1000 && __obf_dd8ccf8e1e49adba.X <= __obf_2cc00eea4019cb5f &&
		__obf_dd8ccf8e1e49adba.Y >= __obf_b9d5f0fcf2bc49a3 && __obf_dd8ccf8e1e49adba.Y <= __obf_10fe878eb0a3e8ba
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
