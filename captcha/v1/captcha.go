package __obf_699038f0de0a4a2b

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
	__obf_cf6522dad577fc60 cache.Cache
	Width                  int     `yaml:"width"`
	Height                 int     `yaml:"height"`
	Length                 int     `yaml:"length"`
	Noises                 int     `yaml:"noises"`
	Expiration             int     `yaml:"expiration"` // minutes
	MaxSkew                float64 `yaml:"max-skew"`
	ColorPalette           color.Palette
}

func (__obf_b0677886b7817654 *Captcha) Init(__obf_974921de5dddc47a cache.Cache) {
	if __obf_b0677886b7817654.__obf_cf6522dad577fc60 == nil {
		__obf_b0677886b7817654.__obf_cf6522dad577fc60 = __obf_974921de5dddc47a
	}
}

// create a new captcha id
func (__obf_b0677886b7817654 *Captcha) Generate() (string, string) {
	__obf_d2a8777e100ff162 := StringUUID()
	__obf_34a3f0602d66ebb2 := RandomStr(__obf_b0677886b7817654.Length)
	__obf_b0677886b7817654.__obf_cf6522dad577fc60.Set(fmt.Sprintf(CachePrefix, __obf_d2a8777e100ff162), __obf_34a3f0602d66ebb2, time.Duration(__obf_b0677886b7817654.Expiration)*time.Minute)
	__obf_399e38ac31a4a0b4 := Image{
		Chars:   __obf_34a3f0602d66ebb2,
		Width:   __obf_b0677886b7817654.Width,
		Height:  __obf_b0677886b7817654.Height,
		Noises:  __obf_b0677886b7817654.Noises,
		MaxSkew: __obf_b0677886b7817654.MaxSkew,
	}
	return __obf_d2a8777e100ff162, __obf_399e38ac31a4a0b4.Base64()
}

// verify from a request
// func (c *Captcha) VerifyReq(req *http.Request) bool {
// 	_ = req.ParseForm()
// 	return c.Verify(req.Form.Get(c.FieldIdName), req.Form.Get(c.FieldCaptchaName))
// }

// direct verify id and challenge string
func (__obf_b0677886b7817654 *Captcha) Verify(__obf_d2a8777e100ff162 string, __obf_0bb9cbd58e182952 string) bool {
	if len(__obf_0bb9cbd58e182952) == 0 || len(__obf_d2a8777e100ff162) == 0 {
		return false
	}

	__obf_a125ba5acb715240, __obf_ccabaa39217573c9 := cache.Get[string](__obf_b0677886b7817654.__obf_cf6522dad577fc60, fmt.Sprintf(CachePrefix, __obf_d2a8777e100ff162))
	if __obf_ccabaa39217573c9 != nil {
		return false
	}
	return strings.EqualFold(__obf_0bb9cbd58e182952, __obf_a125ba5acb715240)
}
