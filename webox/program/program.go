package __obf_7142f8a9722c34c6

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_48b10066da36f8c0 *Config
	__obf_b3ae39fdbd9dd297 *internal.HttpClient
	__obf_43a53dca25dc0d56 string // AccessToken 缓存相关
	__obf_3a93b5f4263321ad int64
	__obf_39a9cd49b1237932 sync. // 过期时间戳 (秒)
				RWMutex
}

var (
	__obf_ba12ebb47621891a *Program
	__obf_b98e4e8f3c459fa7 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_bed71bdd09c60736 *Config) (*Program, error) {
	var __obf_793317582d39a3ba error
	__obf_b98e4e8f3c459fa7.
		Do(func() {
			__obf_ba12ebb47621891a = &Program{__obf_48b10066da36f8c0: __obf_bed71bdd09c60736, __obf_b3ae39fdbd9dd297: internal.NewClient()}
		})
	return __obf_ba12ebb47621891a, __obf_793317582d39a3ba
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
func (__obf_b35978ebc754f4ef *Program) GetAccessToken() (string, error) {
	__obf_b35978ebc754f4ef.__obf_39a9cd49b1237932.
		RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_b35978ebc754f4ef.__obf_43a53dca25dc0d56 != "" && time.Now().Unix() < __obf_b35978ebc754f4ef.__obf_3a93b5f4263321ad {
		__obf_c6f1e67df6e9756d := __obf_b35978ebc754f4ef.__obf_43a53dca25dc0d56
		__obf_b35978ebc754f4ef.__obf_39a9cd49b1237932.
			RUnlock()
		return __obf_c6f1e67df6e9756d, nil
	}
	__obf_b35978ebc754f4ef.__obf_39a9cd49b1237932.
		RUnlock()
	__obf_b35978ebc754f4ef.

		// 如果过期或未获取，则加写锁进行获取
		__obf_39a9cd49b1237932.
		Lock()
	defer __obf_b35978ebc754f4ef.__obf_39a9cd49b1237932.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_b35978ebc754f4ef.__obf_43a53dca25dc0d56 != "" && time.Now().Unix() < __obf_b35978ebc754f4ef.__obf_3a93b5f4263321ad {
		return __obf_b35978ebc754f4ef.__obf_43a53dca25dc0d56, nil
	}

	var (
		__obf_09f8f7dfe2965a9d AccessTokenResponse
		__obf_793317582d39a3ba error
	)
	if __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.UseStableAK {
		__obf_793317582d39a3ba = // 使用稳定版的 AccessToken 接口
			__obf_b35978ebc754f4ef.__obf_b3ae39fdbd9dd297.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppID,
				"secret":     __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppSecret,
			}, &__obf_09f8f7dfe2965a9d)
	} else {
		__obf_793317582d39a3ba = // 使用普通的 AccessToken 接口
			__obf_b35978ebc754f4ef.__obf_b3ae39fdbd9dd297.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppID,
				"secret":     __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppSecret,
			}, &__obf_09f8f7dfe2965a9d)
	}
	if __obf_793317582d39a3ba != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_793317582d39a3ba)
	}

	if __obf_09f8f7dfe2965a9d.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_09f8f7dfe2965a9d.ErrCode, __obf_09f8f7dfe2965a9d.ErrMsg)
	}
	__obf_b35978ebc754f4ef.

		// 更新缓存
		__obf_43a53dca25dc0d56 = __obf_09f8f7dfe2965a9d.AccessToken
	__obf_b35978ebc754f4ef.
		// 提前10分钟过期，给刷新留出余地
		__obf_3a93b5f4263321ad = time.Now().Unix() + int64(__obf_09f8f7dfe2965a9d.ExpiresIn) - 600

	return __obf_b35978ebc754f4ef.__obf_43a53dca25dc0d56, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_b35978ebc754f4ef *Program) Code2Session(__obf_0e3f8b6a44a4cbb4 string) (*Code2SessionResult, error) {
	__obf_3e76c7c43696e0d9 := map[string]string{
		"appid":      __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppID,
		"secret":     __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppSecret,
		"js_code":    __obf_0e3f8b6a44a4cbb4,
		"grant_type": "authorization_code",
	}

	var __obf_09f8f7dfe2965a9d Code2SessionResult
	__obf_793317582d39a3ba := __obf_b35978ebc754f4ef.__obf_b3ae39fdbd9dd297.Get(internal.MiniprogramCode2SessionURL, __obf_3e76c7c43696e0d9, &__obf_09f8f7dfe2965a9d)
	if __obf_793317582d39a3ba != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_793317582d39a3ba)
	}

	if __obf_09f8f7dfe2965a9d.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_09f8f7dfe2965a9d.ErrCode, __obf_09f8f7dfe2965a9d.ErrMsg)
	}
	return &__obf_09f8f7dfe2965a9d, nil
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
func (__obf_b35978ebc754f4ef *Program) GetPhoneNumber(__obf_23904445199ce698 string) (*PhoneNumberData, error) {
	__obf_43a53dca25dc0d56, __obf_793317582d39a3ba := __obf_b35978ebc754f4ef.GetAccessToken()
	if __obf_793317582d39a3ba != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_793317582d39a3ba)
	}
	__obf_6cb6d732f1212a34 := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_43a53dca25dc0d56)

	var __obf_09f8f7dfe2965a9d struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_793317582d39a3ba = __obf_b35978ebc754f4ef.__obf_b3ae39fdbd9dd297.PostJSON(__obf_6cb6d732f1212a34, map[string]any{
		"code": __obf_23904445199ce698,
	},
		&__obf_09f8f7dfe2965a9d,
	)
	if __obf_793317582d39a3ba != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_793317582d39a3ba)
	}

	if __obf_09f8f7dfe2965a9d.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_09f8f7dfe2965a9d.ErrCode, __obf_09f8f7dfe2965a9d.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_09f8f7dfe2965a9d.PhoneNumberData.Watermark.AppID != __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)", __obf_09f8f7dfe2965a9d.
			PhoneNumberData.Watermark.AppID, __obf_b35978ebc754f4ef.__obf_48b10066da36f8c0.AppID)
	}

	return &(__obf_09f8f7dfe2965a9d.PhoneNumberData), nil
}
