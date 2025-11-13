package __obf_85f036759f76ec38

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
	__obf_f9ab07d9fa6ca3a9 cache.Cache
	Width                  int     `yaml:"width"`
	Height                 int     `yaml:"height"`
	Length                 int     `yaml:"length"`
	Noises                 int     `yaml:"noises"`
	Expiration             int     `yaml:"expiration"` // minutes
	MaxSkew                float64 `yaml:"max-skew"`
	ColorPalette           color.Palette
}

func (__obf_b28fa363ae71e592 *Captcha) Init(__obf_92a4a9bad6cbb69a cache.Cache) {
	if __obf_b28fa363ae71e592.__obf_f9ab07d9fa6ca3a9 == nil {
		__obf_b28fa363ae71e592.__obf_f9ab07d9fa6ca3a9 = __obf_92a4a9bad6cbb69a
	}
}

// create a new captcha id
func (__obf_b28fa363ae71e592 *Captcha) Generate() (string, string) {
	__obf_ee79d6821f7987e1 := StringUUID()
	__obf_c5376c2e84331d8b := RandomStr(__obf_b28fa363ae71e592.Length)
	__obf_b28fa363ae71e592.__obf_f9ab07d9fa6ca3a9.Set(fmt.Sprintf(CachePrefix, __obf_ee79d6821f7987e1), __obf_c5376c2e84331d8b, time.Duration(__obf_b28fa363ae71e592.Expiration)*time.Minute)
	__obf_26cee9c7b1194c15 := Image{
		Chars:   __obf_c5376c2e84331d8b,
		Width:   __obf_b28fa363ae71e592.Width,
		Height:  __obf_b28fa363ae71e592.Height,
		Noises:  __obf_b28fa363ae71e592.Noises,
		MaxSkew: __obf_b28fa363ae71e592.MaxSkew,
	}
	return __obf_ee79d6821f7987e1, __obf_26cee9c7b1194c15.Base64()
}

// verify from a request
// func (c *Captcha) VerifyReq(req *http.Request) bool {
// 	_ = req.ParseForm()
// 	return c.Verify(req.Form.Get(c.FieldIdName), req.Form.Get(c.FieldCaptchaName))
// }

// direct verify id and challenge string
func (__obf_b28fa363ae71e592 *Captcha) Verify(__obf_ee79d6821f7987e1 string, __obf_4f2ce9f47dbddb37 string) bool {
	if len(__obf_4f2ce9f47dbddb37) == 0 || len(__obf_ee79d6821f7987e1) == 0 {
		return false
	}

	__obf_e161bbda129a05cf, __obf_bb2771e9267479a3 := cache.Get[string](__obf_b28fa363ae71e592.__obf_f9ab07d9fa6ca3a9, fmt.Sprintf(CachePrefix, __obf_ee79d6821f7987e1))
	if __obf_bb2771e9267479a3 != nil {
		return false
	}
	return strings.EqualFold(__obf_4f2ce9f47dbddb37, __obf_e161bbda129a05cf)
}
