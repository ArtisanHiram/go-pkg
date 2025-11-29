package __obf_b1febe121b2dfc2d

import (
	types "github.com/ArtisanHiram/go-pkg/captcha/types"
	mapstructure "github.com/ArtisanHiram/go-pkg/mapstructure"
)

// CaptchaData is the concrete implementation of the CaptchaData interface
type CaptchaData struct {
	__obf_b78c570ef70ccb05 SlideType
	__obf_fc1b3b29182d412e *types.Block
	__obf_ece3f1a292371bc7 *types.JPEGImage
	__obf_765e5abee3631ca2 *types.PNGImage
}

// Ensure CaptchaData implements util.CaptchaData interface
var _ types.CaptchaData = (*CaptchaData)(nil)

// GetData gets the block data of the CAPTCHA
// return: types.Block data
func (__obf_6c40c4c2143a89cf *CaptchaData) GetData() any {
	return *__obf_6c40c4c2143a89cf.

		// GetPrimary gets the main CAPTCHA image
		// return: Main image as CaptchaImage interface
		__obf_fc1b3b29182d412e
}

func (__obf_6c40c4c2143a89cf *CaptchaData) GetPrimary() types.CaptchaImage {
	return __obf_6c40c4c2143a89cf.

		// GetSecondary gets the tile image
		// return: Tile image as CaptchaImage interface
		__obf_ece3f1a292371bc7
}

func (__obf_6c40c4c2143a89cf *CaptchaData) GetSecondary() types.CaptchaImage {
	return __obf_6c40c4c2143a89cf.__obf_765e5abee3631ca2
}

func (__obf_6c40c4c2143a89cf *CaptchaData) Type() types.CaptchaType {
	if __obf_6c40c4c2143a89cf.__obf_b78c570ef70ccb05 == Move {
		return types.MoveCaptchat
	}
	return types.DragCaptchat
}

// Verify checks if the slide position matches the block position
// point: the coordinates where user slid the block to
// tolerance: allowed pixel deviation for matching
// Returns true if the user's position matches the block's actual position within tolerance
func (__obf_6c40c4c2143a89cf *CaptchaData) Verify(__obf_5ac08fb1f87f15e8 any, __obf_d1cc8c6836b68cd9 int) bool {

	var __obf_fc1b3b29182d412e types.Block
	__obf_24f57cd7cc70ce80 := // 配置 mapstructure
		&mapstructure.DecoderConfig{
			Result:           &__obf_fc1b3b29182d412e,
			TagName:          "json",
			WeaklyTypedInput: true, // 关键：允许弱类型转换（如 float64 -> int）
		}
	__obf_ba95b631f9317985, __obf_5aebd7c35b6603fd := mapstructure.NewDecoder(__obf_24f57cd7cc70ce80)
	if __obf_5aebd7c35b6603fd != nil {
		return false
	}

	if __obf_5aebd7c35b6603fd := __obf_ba95b631f9317985.Decode(__obf_5ac08fb1f87f15e8); __obf_5aebd7c35b6603fd != nil {
		return false
	}
	__obf_dc8b8f1229f0df5a := __obf_6c40c4c2143a89cf.__obf_fc1b3b29182d412e.X - __obf_d1cc8c6836b68cd9
	__obf_bd9fc31f6f66f15e := __obf_6c40c4c2143a89cf.__obf_fc1b3b29182d412e.X + __obf_d1cc8c6836b68cd9
	__obf_2ffee3a51bbd65fe := __obf_6c40c4c2143a89cf.__obf_fc1b3b29182d412e.Y - __obf_d1cc8c6836b68cd9
	__obf_c6d7240ffc73ff0b := __obf_6c40c4c2143a89cf.__obf_fc1b3b29182d412e.Y + __obf_d1cc8c6836b68cd9

	return __obf_fc1b3b29182d412e.X >= __obf_dc8b8f1229f0df5a && __obf_fc1b3b29182d412e.X <= __obf_bd9fc31f6f66f15e && __obf_fc1b3b29182d412e.
		Y >= __obf_2ffee3a51bbd65fe && __obf_fc1b3b29182d412e.Y <= __obf_c6d7240ffc73ff0b

}
