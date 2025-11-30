package __obf_8f6d6401b1ecc0c1

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_d60f5f321ab8c4a9 *Config
	__obf_3f0b496c0fdc6ff5 *internal.HttpClient
	__obf_5e7e9299e26e93b9 string // AccessToken 缓存相关
	__obf_aa202438f2f5e687 int64
	__obf_a8003d3122360db0 sync. // 过期时间戳 (秒)
				RWMutex
}

var (
	__obf_6848ad629d1c0fec *Program
	__obf_3ce9aee4805f4fa1 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_016721780360a479 *Config) (*Program, error) {
	var __obf_48701698d5850dd0 error
	__obf_3ce9aee4805f4fa1.
		Do(func() {
			__obf_6848ad629d1c0fec = &Program{__obf_d60f5f321ab8c4a9: __obf_016721780360a479, __obf_3f0b496c0fdc6ff5: internal.NewClient()}
		})
	return __obf_6848ad629d1c0fec, __obf_48701698d5850dd0
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
func (__obf_ea3957ac67b0e48d *Program) GetAccessToken() (string, error) {
	__obf_ea3957ac67b0e48d.__obf_a8003d3122360db0.
		RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_ea3957ac67b0e48d.__obf_5e7e9299e26e93b9 != "" && time.Now().Unix() < __obf_ea3957ac67b0e48d.__obf_aa202438f2f5e687 {
		__obf_405df278cea1466a := __obf_ea3957ac67b0e48d.__obf_5e7e9299e26e93b9
		__obf_ea3957ac67b0e48d.__obf_a8003d3122360db0.
			RUnlock()
		return __obf_405df278cea1466a, nil
	}
	__obf_ea3957ac67b0e48d.__obf_a8003d3122360db0.
		RUnlock()
	__obf_ea3957ac67b0e48d.

		// 如果过期或未获取，则加写锁进行获取
		__obf_a8003d3122360db0.
		Lock()
	defer __obf_ea3957ac67b0e48d.__obf_a8003d3122360db0.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_ea3957ac67b0e48d.__obf_5e7e9299e26e93b9 != "" && time.Now().Unix() < __obf_ea3957ac67b0e48d.__obf_aa202438f2f5e687 {
		return __obf_ea3957ac67b0e48d.__obf_5e7e9299e26e93b9, nil
	}

	var (
		__obf_4fc01b7a9d375fa2 AccessTokenResponse
		__obf_48701698d5850dd0 error
	)
	if __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.UseStableAK {
		__obf_48701698d5850dd0 = // 使用稳定版的 AccessToken 接口
			__obf_ea3957ac67b0e48d.__obf_3f0b496c0fdc6ff5.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppID,
				"secret":     __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppSecret,
			}, &__obf_4fc01b7a9d375fa2)
	} else {
		__obf_48701698d5850dd0 = // 使用普通的 AccessToken 接口
			__obf_ea3957ac67b0e48d.__obf_3f0b496c0fdc6ff5.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppID,
				"secret":     __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppSecret,
			}, &__obf_4fc01b7a9d375fa2)
	}
	if __obf_48701698d5850dd0 != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_48701698d5850dd0)
	}

	if __obf_4fc01b7a9d375fa2.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_4fc01b7a9d375fa2.ErrCode, __obf_4fc01b7a9d375fa2.ErrMsg)
	}
	__obf_ea3957ac67b0e48d.

		// 更新缓存
		__obf_5e7e9299e26e93b9 = __obf_4fc01b7a9d375fa2.AccessToken
	__obf_ea3957ac67b0e48d.
		// 提前10分钟过期，给刷新留出余地
		__obf_aa202438f2f5e687 = time.Now().Unix() + int64(__obf_4fc01b7a9d375fa2.ExpiresIn) - 600

	return __obf_ea3957ac67b0e48d.__obf_5e7e9299e26e93b9, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_ea3957ac67b0e48d *Program) Code2Session(__obf_0836a704144ea1f9 string) (*Code2SessionResult, error) {
	__obf_53fa25be78082cb5 := map[string]string{
		"appid":      __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppID,
		"secret":     __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppSecret,
		"js_code":    __obf_0836a704144ea1f9,
		"grant_type": "authorization_code",
	}

	var __obf_4fc01b7a9d375fa2 Code2SessionResult
	__obf_48701698d5850dd0 := __obf_ea3957ac67b0e48d.__obf_3f0b496c0fdc6ff5.Get(internal.MiniprogramCode2SessionURL, __obf_53fa25be78082cb5, &__obf_4fc01b7a9d375fa2)
	if __obf_48701698d5850dd0 != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_48701698d5850dd0)
	}

	if __obf_4fc01b7a9d375fa2.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_4fc01b7a9d375fa2.ErrCode, __obf_4fc01b7a9d375fa2.ErrMsg)
	}
	return &__obf_4fc01b7a9d375fa2, nil
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
func (__obf_ea3957ac67b0e48d *Program) GetPhoneNumber(__obf_ac96291eebc7ba34 string) (*PhoneNumberData, error) {
	__obf_5e7e9299e26e93b9, __obf_48701698d5850dd0 := __obf_ea3957ac67b0e48d.GetAccessToken()
	if __obf_48701698d5850dd0 != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_48701698d5850dd0)
	}
	__obf_ec44de140db7b76b := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_5e7e9299e26e93b9)

	var __obf_4fc01b7a9d375fa2 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_48701698d5850dd0 = __obf_ea3957ac67b0e48d.__obf_3f0b496c0fdc6ff5.PostJSON(__obf_ec44de140db7b76b, map[string]any{
		"code": __obf_ac96291eebc7ba34,
	},
		&__obf_4fc01b7a9d375fa2,
	)
	if __obf_48701698d5850dd0 != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_48701698d5850dd0)
	}

	if __obf_4fc01b7a9d375fa2.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_4fc01b7a9d375fa2.ErrCode, __obf_4fc01b7a9d375fa2.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_4fc01b7a9d375fa2.PhoneNumberData.Watermark.AppID != __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)", __obf_4fc01b7a9d375fa2.
			PhoneNumberData.Watermark.AppID, __obf_ea3957ac67b0e48d.__obf_d60f5f321ab8c4a9.AppID)
	}

	return &(__obf_4fc01b7a9d375fa2.PhoneNumberData), nil
}
