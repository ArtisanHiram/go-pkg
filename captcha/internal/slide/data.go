package __obf_45140da9e3f6647e

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_f05be736ef82a7b3 SlideType
	__obf_3c2937c851cafbfc *types.Block
	__obf_8fcc6dfc88c19e7c *types.JPEGImage
	__obf_a99609c2d6fcf851 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_012c5ff8d6b175b3 *CaptchaData) GetData() any {
	return *__obf_012c5ff8d6b175b3.

		// GetPrimary gets the main CAPTCHA image
		// return: Main image as CaptchaImage interface
		__obf_3c2937c851cafbfc
}

func (__obf_012c5ff8d6b175b3 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_012c5ff8d6b175b3.

		// GetSecondary gets the tile image
		// return: Tile image as CaptchaImage interface
		__obf_8fcc6dfc88c19e7c
}

func (__obf_012c5ff8d6b175b3 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_012c5ff8d6b175b3.__obf_a99609c2d6fcf851
}

func (__obf_012c5ff8d6b175b3 *CaptchaData) Type() types.CaptchaType {
	if __obf_012c5ff8d6b175b3.__obf_f05be736ef82a7b3 == Move {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_012c5ff8d6b175b3 *CaptchaData) Verify(__obf_6b4196184ed131c7 any, __obf_d301885bbd1a637f int) bool {

	var __obf_3c2937c851cafbfc types.Block
	__obf_6f2e94dc3d985a7d := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_3c2937c851cafbfc,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_1dc98a9d6c300853, __obf_7d580363a4d8cba8 := mapstructure.NewDecoder(__obf_6f2e94dc3d985a7d)
	if __obf_7d580363a4d8cba8 != nil {
		return false
	}

	if __obf_7d580363a4d8cba8 := __obf_1dc98a9d6c300853.Decode(__obf_6b4196184ed131c7); __obf_7d580363a4d8cba8 != nil {
		return false
	}
	__obf_70d868ca4463fe69 := __obf_012c5ff8d6b175b3.__obf_3c2937c851cafbfc.X - __obf_d301885bbd1a637f
	__obf_edd9e4e33a31a840 := __obf_012c5ff8d6b175b3.__obf_3c2937c851cafbfc.X + __obf_d301885bbd1a637f
	__obf_65d8b2b2d92f60b7 := __obf_012c5ff8d6b175b3.__obf_3c2937c851cafbfc.Y - __obf_d301885bbd1a637f
	__obf_4b9ec9e4d4baecac := __obf_012c5ff8d6b175b3.__obf_3c2937c851cafbfc.Y + __obf_d301885bbd1a637f

	return __obf_3c2937c851cafbfc.X >= __obf_70d868ca4463fe69 && __obf_3c2937c851cafbfc.X <= __obf_edd9e4e33a31a840 && __obf_3c2937c851cafbfc.
		Y >= __obf_65d8b2b2d92f60b7 && __obf_3c2937c851cafbfc.Y <= __obf_4b9ec9e4d4baecac

}
