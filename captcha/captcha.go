package __obf_e49f90b5aaf58063

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
	__obf_57cfccd67b13108e cache.Cache
	Width                  int     `yaml:"width"`
	Height                 int     `yaml:"height"`
	Length                 int     `yaml:"length"`
	Noises                 int     `yaml:"noises"`
	Expiration             int     `yaml:"expiration"` // minutes
	MaxSkew                float64 `yaml:"max-skew"`
	ColorPalette           color.Palette
}

func (__obf_98f8d0c445e06488 *Captcha) Init(__obf_bdd8cc607d070139 cache.Cache) {
	if __obf_98f8d0c445e06488.__obf_57cfccd67b13108e == nil {
		__obf_98f8d0c445e06488.__obf_57cfccd67b13108e = __obf_bdd8cc607d070139
	}
}

// create a new captcha id
func (__obf_98f8d0c445e06488 *Captcha) Generate() (string, string) {
	__obf_55fe7235cf866948 := StringUUID()
	__obf_10eafabe741c335e := RandomStr(__obf_98f8d0c445e06488.Length)
	__obf_98f8d0c445e06488.__obf_57cfccd67b13108e.Set(fmt.Sprintf(CachePrefix, __obf_55fe7235cf866948), __obf_10eafabe741c335e, time.Duration(__obf_98f8d0c445e06488.Expiration)*time.Minute)
	__obf_5ef5f278f950e167 := Image{
		Chars:   __obf_10eafabe741c335e,
		Width:   __obf_98f8d0c445e06488.Width,
		Height:  __obf_98f8d0c445e06488.Height,
		Noises:  __obf_98f8d0c445e06488.Noises,
		MaxSkew: __obf_98f8d0c445e06488.MaxSkew,
	}
	return __obf_55fe7235cf866948, __obf_5ef5f278f950e167.Base64()
}

// verify from a request
// func (c *Captcha) VerifyReq(req *http.Request) bool {
// 	_ = req.ParseForm()
// 	return c.Verify(req.Form.Get(c.FieldIdName), req.Form.Get(c.FieldCaptchaName))
// }

// direct verify id and challenge string
func (__obf_98f8d0c445e06488 *Captcha) Verify(__obf_55fe7235cf866948 string, __obf_21e720384978b98d string) bool {
	if len(__obf_21e720384978b98d) == 0 || len(__obf_55fe7235cf866948) == 0 {
		return false
	}

	__obf_9bae8b63c43448cc, __obf_fea6cef273213c03 := cache.Get[string](__obf_98f8d0c445e06488.__obf_57cfccd67b13108e, fmt.Sprintf(CachePrefix, __obf_55fe7235cf866948))
	if __obf_fea6cef273213c03 != nil {
		return false
	}
	return strings.EqualFold(__obf_21e720384978b98d, __obf_9bae8b63c43448cc)
}
