package __obf_7621171b3eab50e9

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	Type                   types.SlideType `json:"type"`
	Block                  *types.Block    `json:"block"`
	__obf_aa601e1953c0218c *types.JPEGImage
	__obf_7aec6e82031fcbf1 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_4c68703ee0481a75 *CaptchaData) GetData() any {
	return *__obf_4c68703ee0481a75.Block
}

// GetPrimary gets the main CAPTCHA image
// return: Main image as CaptchaImage interface
func (__obf_4c68703ee0481a75 *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_4c68703ee0481a75.

		// GetSecondary gets the tile image
		// return: Tile image as CaptchaImage interface
		__obf_aa601e1953c0218c
}

func (__obf_4c68703ee0481a75 *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_4c68703ee0481a75.__obf_7aec6e82031fcbf1
}

func (__obf_4c68703ee0481a75 *CaptchaData) GetType() types.CaptchaType {
	if __obf_4c68703ee0481a75.Type == types.MoveSlide {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_4c68703ee0481a75 *CaptchaData) Verify(__obf_56fd94495dd4d900 any, __obf_d16cda94d2b35b97 int) bool {

	var __obf_fe556acd0600dc8f types.Block
	__obf_ba5b503de4e5a4ea := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_fe556acd0600dc8f,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_931f22a7197912d3, __obf_a309197d4a7b9caa := mapstructure.NewDecoder(__obf_ba5b503de4e5a4ea)
	if __obf_a309197d4a7b9caa != nil {
		return false
	}

	if __obf_a309197d4a7b9caa := __obf_931f22a7197912d3.Decode(__obf_56fd94495dd4d900); __obf_a309197d4a7b9caa != nil {
		return false
	}
	__obf_1fdf51284bbb4a0e := __obf_4c68703ee0481a75.Block.X - __obf_d16cda94d2b35b97
	__obf_49118c907a90b8c9 := __obf_4c68703ee0481a75.Block.X + __obf_d16cda94d2b35b97
	__obf_15eb179482b482bc := __obf_4c68703ee0481a75.Block.Y - __obf_d16cda94d2b35b97
	__obf_afb4b90a134c0b9f := __obf_4c68703ee0481a75.Block.Y + __obf_d16cda94d2b35b97

	return __obf_fe556acd0600dc8f.X >= __obf_1fdf51284bbb4a0e && __obf_fe556acd0600dc8f.X <= __obf_49118c907a90b8c9 && __obf_fe556acd0600dc8f.
		Y >= __obf_15eb179482b482bc && __obf_fe556acd0600dc8f.Y <= __obf_afb4b90a134c0b9f

}
