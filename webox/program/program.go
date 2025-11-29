package __obf_b1435a5edddb39ec

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_91d3db89668021be *Config
	__obf_2a10f2aa7b1082ea *internal.HttpClient
	__obf_d8ec236f8faec2b5 string // AccessToken 缓存相关
	__obf_6138d4476d2f2737 int64
	__obf_c2b8f2ebe52c70d3 sync. // 过期时间戳 (秒)
				RWMutex
}

var (
	__obf_265ca6a8d6ba388c *Program
	__obf_76925092da95a0ce sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_7569f50dc52618f7 *Config) (*Program, error) {
	var __obf_0a261d29aa4c028c error
	__obf_76925092da95a0ce.
		Do(func() {
			__obf_265ca6a8d6ba388c = &Program{__obf_91d3db89668021be: __obf_7569f50dc52618f7, __obf_2a10f2aa7b1082ea: internal.NewClient()}
		})
	return __obf_265ca6a8d6ba388c, __obf_0a261d29aa4c028c
}

// AccessTokenResponse 获取 AccessToken 接口返回结构
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"` // 有效期，单位：秒
	model.ErrResponse
}

// GetAccessToken 获取小程序全局唯一调用凭据 access_token
// 注意：access_token 有效期为2小时，重复获取会导致上一次获取的失效。
// 建议在生产环境中使用缓存机制，避免频繁获取。
func (__obf_f63b979e580b020d *Program) GetAccessToken() (string, error) {
	__obf_f63b979e580b020d.__obf_c2b8f2ebe52c70d3.
		RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_f63b979e580b020d.__obf_d8ec236f8faec2b5 != "" && time.Now().Unix() < __obf_f63b979e580b020d.__obf_6138d4476d2f2737 {
		__obf_17703496eac078ab := __obf_f63b979e580b020d.__obf_d8ec236f8faec2b5
		__obf_f63b979e580b020d.__obf_c2b8f2ebe52c70d3.
			RUnlock()
		return __obf_17703496eac078ab, nil
	}
	__obf_f63b979e580b020d.__obf_c2b8f2ebe52c70d3.
		RUnlock()
	__obf_f63b979e580b020d.

		// 如果过期或未获取，则加写锁进行获取
		__obf_c2b8f2ebe52c70d3.
		Lock()
	defer __obf_f63b979e580b020d.__obf_c2b8f2ebe52c70d3.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_f63b979e580b020d.__obf_d8ec236f8faec2b5 != "" && time.Now().Unix() < __obf_f63b979e580b020d.__obf_6138d4476d2f2737 {
		return __obf_f63b979e580b020d.__obf_d8ec236f8faec2b5, nil
	}

	var (
		__obf_dc05770778b0cf5b AccessTokenResponse
		__obf_0a261d29aa4c028c error
	)
	if __obf_f63b979e580b020d.__obf_91d3db89668021be.UseStableAK {
		__obf_0a261d29aa4c028c = // 使用稳定版的 AccessToken 接口
			__obf_f63b979e580b020d.__obf_2a10f2aa7b1082ea.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_f63b979e580b020d.__obf_91d3db89668021be.AppID,
				"secret":     __obf_f63b979e580b020d.__obf_91d3db89668021be.AppSecret,
			}, &__obf_dc05770778b0cf5b)
	} else {
		__obf_0a261d29aa4c028c = // 使用普通的 AccessToken 接口
			__obf_f63b979e580b020d.__obf_2a10f2aa7b1082ea.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_f63b979e580b020d.__obf_91d3db89668021be.AppID,
				"secret":     __obf_f63b979e580b020d.__obf_91d3db89668021be.AppSecret,
			}, &__obf_dc05770778b0cf5b)
	}
	if __obf_0a261d29aa4c028c != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_0a261d29aa4c028c)
	}

	if __obf_dc05770778b0cf5b.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_dc05770778b0cf5b.ErrCode, __obf_dc05770778b0cf5b.ErrMsg)
	}
	__obf_f63b979e580b020d.

		// 更新缓存
		__obf_d8ec236f8faec2b5 = __obf_dc05770778b0cf5b.AccessToken
	__obf_f63b979e580b020d.
		// 提前10分钟过期，给刷新留出余地
		__obf_6138d4476d2f2737 = time.Now().Unix() + int64(__obf_dc05770778b0cf5b.ExpiresIn) - 600

	return __obf_f63b979e580b020d.__obf_d8ec236f8faec2b5, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_f63b979e580b020d *Program) Code2Session(__obf_16fb43d74c78c144 string) (*Code2SessionResult, error) {
	__obf_a002a13413f96256 := map[string]string{
		"appid":      __obf_f63b979e580b020d.__obf_91d3db89668021be.AppID,
		"secret":     __obf_f63b979e580b020d.__obf_91d3db89668021be.AppSecret,
		"js_code":    __obf_16fb43d74c78c144,
		"grant_type": "authorization_code",
	}

	var __obf_dc05770778b0cf5b Code2SessionResult
	__obf_0a261d29aa4c028c := __obf_f63b979e580b020d.__obf_2a10f2aa7b1082ea.Get(internal.MiniprogramCode2SessionURL, __obf_a002a13413f96256, &__obf_dc05770778b0cf5b)
	if __obf_0a261d29aa4c028c != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_0a261d29aa4c028c)
	}

	if __obf_dc05770778b0cf5b.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_dc05770778b0cf5b.ErrCode, __obf_dc05770778b0cf5b.ErrMsg)
	}
	return &__obf_dc05770778b0cf5b, nil
}

type PhoneNumberData struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Watermark       struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

// GetPhoneNumber 通过小程序前端传入的code获取用户手机号 (新版本方式)
// phoneCode: 小程序前端通过 `wx.getPhoneNumber` 接口获取的动态令牌 `e.detail.code`
func (__obf_f63b979e580b020d *Program) GetPhoneNumber(__obf_cd14bc7f56ddf9c6 string) (*PhoneNumberData, error) {
	__obf_d8ec236f8faec2b5, __obf_0a261d29aa4c028c := __obf_f63b979e580b020d.GetAccessToken()
	if __obf_0a261d29aa4c028c != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_0a261d29aa4c028c)
	}
	__obf_8f66d712f6af48dd := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_d8ec236f8faec2b5)

	var __obf_dc05770778b0cf5b struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_0a261d29aa4c028c = __obf_f63b979e580b020d.__obf_2a10f2aa7b1082ea.PostJSON(__obf_8f66d712f6af48dd, map[string]any{
		"code": __obf_cd14bc7f56ddf9c6,
	},
		&__obf_dc05770778b0cf5b,
	)
	if __obf_0a261d29aa4c028c != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_0a261d29aa4c028c)
	}

	if __obf_dc05770778b0cf5b.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_dc05770778b0cf5b.ErrCode, __obf_dc05770778b0cf5b.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_dc05770778b0cf5b.PhoneNumberData.Watermark.AppID != __obf_f63b979e580b020d.__obf_91d3db89668021be.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)", __obf_dc05770778b0cf5b.
			PhoneNumberData.Watermark.AppID, __obf_f63b979e580b020d.__obf_91d3db89668021be.AppID)
	}

	return &(__obf_dc05770778b0cf5b.PhoneNumberData), nil
}
