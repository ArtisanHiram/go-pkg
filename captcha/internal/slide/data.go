package __obf_ae0bbdaffaa1f242

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_a805d52fbb2d47f1 types.SlideType
	__obf_73053d5754e12223 *types.Block
	__obf_f37f373d0317b914 *types.JPEGImage
	__obf_30c3995a9b03c58a *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_9e636908868342b9 *CaptchaData) GetData() any {
	return *__obf_9e636908868342b9.

		// GetPrimary gets the main CAPTCHA image
		// return: Main image as CaptchaImage interface
		__obf_73053d5754e12223
}

func (__obf_9e636908868342b9 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_9e636908868342b9.

		// GetSecondary gets the tile image
		// return: Tile image as CaptchaImage interface
		__obf_f37f373d0317b914
}

func (__obf_9e636908868342b9 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_9e636908868342b9.__obf_30c3995a9b03c58a
}

func (__obf_9e636908868342b9 *CaptchaData) Type() types.CaptchaType {
	if __obf_9e636908868342b9.__obf_a805d52fbb2d47f1 == types.MoveSlide {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_9e636908868342b9 *CaptchaData) Verify(__obf_b366776e1c1deb15 any, __obf_efb130efcadb715e int) bool {

	var __obf_73053d5754e12223 types.Block
	__obf_69eae12501a775da := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_73053d5754e12223,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_95ab8ca85c1acd59, __obf_d5b002b58d3eb1d0 := mapstructure.NewDecoder(__obf_69eae12501a775da)
	if __obf_d5b002b58d3eb1d0 != nil {
		return false
	}

	if __obf_d5b002b58d3eb1d0 := __obf_95ab8ca85c1acd59.Decode(__obf_b366776e1c1deb15); __obf_d5b002b58d3eb1d0 != nil {
		return false
	}
	__obf_e8faaf1b825651d2 := __obf_9e636908868342b9.__obf_73053d5754e12223.X - __obf_efb130efcadb715e
	__obf_9e66dbcb2219802b := __obf_9e636908868342b9.__obf_73053d5754e12223.X + __obf_efb130efcadb715e
	__obf_2b463ad7b0a75756 := __obf_9e636908868342b9.__obf_73053d5754e12223.Y - __obf_efb130efcadb715e
	__obf_2ae8b791cef7d78a := __obf_9e636908868342b9.__obf_73053d5754e12223.Y + __obf_efb130efcadb715e

	return __obf_73053d5754e12223.X >= __obf_e8faaf1b825651d2 && __obf_73053d5754e12223.X <= __obf_9e66dbcb2219802b && __obf_73053d5754e12223.
		Y >= __obf_2b463ad7b0a75756 && __obf_73053d5754e12223.Y <= __obf_2ae8b791cef7d78a

}
