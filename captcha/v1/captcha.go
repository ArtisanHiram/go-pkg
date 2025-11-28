package __obf_4574f91399c0b6ef

import (
	"fmt"
	cache "github.com/ArtisanHiram/go-pkg/cache"
	"image/color"
	"strings"
	"time"
)

const (
	CachePrefix = "captcha_%s"
)

// Captcha represents a captcha service.
type Captcha struct {
	__obf_e9b428b374c6d26d cache.Cache
	Width                  int     `yaml:"width"`
	Height                 int     `yaml:"height"`
	Length                 int     `yaml:"length"`
	Noises                 int     `yaml:"noises"`
	Expiration             int     `yaml:"expiration"` // minutes
	MaxSkew                float64 `yaml:"max-skew"`
	ColorPalette           color.Palette
}

func (__obf_77c8e6f59f12faaf *Captcha) Init(__obf_d02a48564af7ad31 cache.Cache) {
	if __obf_77c8e6f59f12faaf.__obf_e9b428b374c6d26d == nil {
		__obf_77c8e6f59f12faaf.__obf_e9b428b374c6d26d = __obf_d02a48564af7ad31
	}
}

// create a new captcha id
func (__obf_77c8e6f59f12faaf *Captcha) Generate() (string, string) {
	__obf_4612d75530c7f5f4 := StringUUID()
	__obf_dc067a42154bf95f := RandomStr(__obf_77c8e6f59f12faaf.Length)
	__obf_77c8e6f59f12faaf.__obf_e9b428b374c6d26d.Set(fmt.Sprintf(CachePrefix, __obf_4612d75530c7f5f4), __obf_dc067a42154bf95f, time.Duration(__obf_77c8e6f59f12faaf.Expiration)*time.Minute)
	__obf_0a8d6fcbe8ee1bdf := Image{
		Chars:   __obf_dc067a42154bf95f,
		Width:   __obf_77c8e6f59f12faaf.Width,
		Height:  __obf_77c8e6f59f12faaf.Height,
		Noises:  __obf_77c8e6f59f12faaf.Noises,
		MaxSkew: __obf_77c8e6f59f12faaf.MaxSkew,
	}
	return __obf_4612d75530c7f5f4, __obf_0a8d6fcbe8ee1bdf.Base64()
}

// verify from a request
// func (c *Captcha) VerifyReq(req *http.Request) bool {
// 	_ = req.ParseForm()
// 	return c.Verify(req.Form.Get(c.FieldIdName), req.Form.Get(c.FieldCaptchaName))
// }

// direct verify id and challenge string
func (__obf_77c8e6f59f12faaf *Captcha) Verify(__obf_4612d75530c7f5f4 string, __obf_0ecaede6f5f6c1f7 string) bool {
	if len(__obf_0ecaede6f5f6c1f7) == 0 || len(__obf_4612d75530c7f5f4) == 0 {
		return false
	}

	__obf_6c624db67cd15db8, __obf_f03d347b1ec97670 := cache.Get[string](__obf_77c8e6f59f12faaf.__obf_e9b428b374c6d26d, fmt.Sprintf(CachePrefix, __obf_4612d75530c7f5f4))
	if __obf_f03d347b1ec97670 != nil {
		return false
	}
	return strings.EqualFold(__obf_0ecaede6f5f6c1f7, __obf_6c624db67cd15db8)
}
