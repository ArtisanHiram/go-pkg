package __obf_df8b7004e69b594a

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_1fad1ea30d01716d *Config
	__obf_329d7348b9c3e7a2 *internal.HttpClient
	__obf_80109ad1c37b6a16 string // AccessToken 缓存相关
	__obf_ca102c28189ef9e6 int64
	__obf_03d279f1e3021a8c sync. // 过期时间戳 (秒)
				RWMutex
}

var (
	__obf_20a1e3c0d972f997 *Program
	__obf_1d00f89e685b1cee sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_c92fa05b1d109847 *Config) (*Program, error) {
	var __obf_c76bb9af4ce678f1 error
	__obf_1d00f89e685b1cee.
		Do(func() {
			__obf_20a1e3c0d972f997 = &Program{__obf_1fad1ea30d01716d: __obf_c92fa05b1d109847, __obf_329d7348b9c3e7a2: internal.NewClient()}
		})
	return __obf_20a1e3c0d972f997, __obf_c76bb9af4ce678f1
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
func (__obf_74d04cb532fe3038 *Program) GetAccessToken() (string, error) {
	__obf_74d04cb532fe3038.__obf_03d279f1e3021a8c.
		RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_74d04cb532fe3038.__obf_80109ad1c37b6a16 != "" && time.Now().Unix() < __obf_74d04cb532fe3038.__obf_ca102c28189ef9e6 {
		__obf_589c8814fd87cde5 := __obf_74d04cb532fe3038.__obf_80109ad1c37b6a16
		__obf_74d04cb532fe3038.__obf_03d279f1e3021a8c.
			RUnlock()
		return __obf_589c8814fd87cde5, nil
	}
	__obf_74d04cb532fe3038.__obf_03d279f1e3021a8c.
		RUnlock()
	__obf_74d04cb532fe3038.

		// 如果过期或未获取，则加写锁进行获取
		__obf_03d279f1e3021a8c.
		Lock()
	defer __obf_74d04cb532fe3038.__obf_03d279f1e3021a8c.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_74d04cb532fe3038.__obf_80109ad1c37b6a16 != "" && time.Now().Unix() < __obf_74d04cb532fe3038.__obf_ca102c28189ef9e6 {
		return __obf_74d04cb532fe3038.__obf_80109ad1c37b6a16, nil
	}

	var (
		__obf_19818ccc6cca2339 AccessTokenResponse
		__obf_c76bb9af4ce678f1 error
	)
	if __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.UseStableAK {
		__obf_c76bb9af4ce678f1 = // 使用稳定版的 AccessToken 接口
			__obf_74d04cb532fe3038.__obf_329d7348b9c3e7a2.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppID,
				"secret":     __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppSecret,
			}, &__obf_19818ccc6cca2339)
	} else {
		__obf_c76bb9af4ce678f1 = // 使用普通的 AccessToken 接口
			__obf_74d04cb532fe3038.__obf_329d7348b9c3e7a2.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppID,
				"secret":     __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppSecret,
			}, &__obf_19818ccc6cca2339)
	}
	if __obf_c76bb9af4ce678f1 != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_c76bb9af4ce678f1)
	}

	if __obf_19818ccc6cca2339.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_19818ccc6cca2339.ErrCode, __obf_19818ccc6cca2339.ErrMsg)
	}
	__obf_74d04cb532fe3038.

		// 更新缓存
		__obf_80109ad1c37b6a16 = __obf_19818ccc6cca2339.AccessToken
	__obf_74d04cb532fe3038.
		// 提前10分钟过期，给刷新留出余地
		__obf_ca102c28189ef9e6 = time.Now().Unix() + int64(__obf_19818ccc6cca2339.ExpiresIn) - 600

	return __obf_74d04cb532fe3038.__obf_80109ad1c37b6a16, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_74d04cb532fe3038 *Program) Code2Session(__obf_67ce045a806a5052 string) (*Code2SessionResult, error) {
	__obf_4d6391cf227a558e := map[string]string{
		"appid":      __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppID,
		"secret":     __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppSecret,
		"js_code":    __obf_67ce045a806a5052,
		"grant_type": "authorization_code",
	}

	var __obf_19818ccc6cca2339 Code2SessionResult
	__obf_c76bb9af4ce678f1 := __obf_74d04cb532fe3038.__obf_329d7348b9c3e7a2.Get(internal.MiniprogramCode2SessionURL, __obf_4d6391cf227a558e, &__obf_19818ccc6cca2339)
	if __obf_c76bb9af4ce678f1 != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_c76bb9af4ce678f1)
	}

	if __obf_19818ccc6cca2339.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_19818ccc6cca2339.ErrCode, __obf_19818ccc6cca2339.ErrMsg)
	}
	return &__obf_19818ccc6cca2339, nil
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
func (__obf_74d04cb532fe3038 *Program) GetPhoneNumber(__obf_83f16234e3b7afda string) (*PhoneNumberData, error) {
	__obf_80109ad1c37b6a16, __obf_c76bb9af4ce678f1 := __obf_74d04cb532fe3038.GetAccessToken()
	if __obf_c76bb9af4ce678f1 != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_c76bb9af4ce678f1)
	}
	__obf_7babf847d79a6167 := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_80109ad1c37b6a16)

	var __obf_19818ccc6cca2339 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_c76bb9af4ce678f1 = __obf_74d04cb532fe3038.__obf_329d7348b9c3e7a2.PostJSON(__obf_7babf847d79a6167, map[string]any{
		"code": __obf_83f16234e3b7afda,
	},
		&__obf_19818ccc6cca2339,
	)
	if __obf_c76bb9af4ce678f1 != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_c76bb9af4ce678f1)
	}

	if __obf_19818ccc6cca2339.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_19818ccc6cca2339.ErrCode, __obf_19818ccc6cca2339.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_19818ccc6cca2339.PhoneNumberData.Watermark.AppID != __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)", __obf_19818ccc6cca2339.
			PhoneNumberData.Watermark.AppID, __obf_74d04cb532fe3038.__obf_1fad1ea30d01716d.AppID)
	}

	return &(__obf_19818ccc6cca2339.PhoneNumberData), nil
}
