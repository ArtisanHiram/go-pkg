package __obf_89d565d7ca4115dd

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_8a104016d2195942 types.SlideType
	__obf_3fc2dcf01757c488 *types.Block
	__obf_d2af1e67e2974419 *types.JPEGImage
	__obf_cfcc1a610d000093 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_3ad6454bef56219b *CaptchaData) GetData() any {
	return *__obf_3ad6454bef56219b.

		// GetPrimary gets the main CAPTCHA image
		// return: Main image as CaptchaImage interface
		__obf_3fc2dcf01757c488
}

func (__obf_3ad6454bef56219b *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_3ad6454bef56219b.

		// GetSecondary gets the tile image
		// return: Tile image as CaptchaImage interface
		__obf_d2af1e67e2974419
}

func (__obf_3ad6454bef56219b *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_3ad6454bef56219b.__obf_cfcc1a610d000093
}

func (__obf_3ad6454bef56219b *CaptchaData) Type() types.CaptchaType {
	if __obf_3ad6454bef56219b.__obf_8a104016d2195942 == types.MoveSlide {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_3ad6454bef56219b *CaptchaData) Verify(__obf_1cb93459ee8ada1f any, __obf_d9d254d69c98d049 int) bool {

	var __obf_3fc2dcf01757c488 types.Block
	__obf_e1599f725de2f261 := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_3fc2dcf01757c488,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_4984481493b5de2b, __obf_8bf8c3d5ac659533 := mapstructure.NewDecoder(__obf_e1599f725de2f261)
	if __obf_8bf8c3d5ac659533 != nil {
		return false
	}

	if __obf_8bf8c3d5ac659533 := __obf_4984481493b5de2b.Decode(__obf_1cb93459ee8ada1f); __obf_8bf8c3d5ac659533 != nil {
		return false
	}
	__obf_c48c36eb93c97ddf := __obf_3ad6454bef56219b.__obf_3fc2dcf01757c488.X - __obf_d9d254d69c98d049
	__obf_2900f78c6e2e5922 := __obf_3ad6454bef56219b.__obf_3fc2dcf01757c488.X + __obf_d9d254d69c98d049
	__obf_d533c57c9510979a := __obf_3ad6454bef56219b.__obf_3fc2dcf01757c488.Y - __obf_d9d254d69c98d049
	__obf_252c6cd797d4b02f := __obf_3ad6454bef56219b.__obf_3fc2dcf01757c488.Y + __obf_d9d254d69c98d049

	return __obf_3fc2dcf01757c488.X >= __obf_c48c36eb93c97ddf && __obf_3fc2dcf01757c488.X <= __obf_2900f78c6e2e5922 && __obf_3fc2dcf01757c488.
		Y >= __obf_d533c57c9510979a && __obf_3fc2dcf01757c488.Y <= __obf_252c6cd797d4b02f

}
