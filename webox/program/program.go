package __obf_abf0107273afed37

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_66fee4aa5605d761 *Config
	__obf_39e54c9148149a72 *internal.HttpClient
	// AccessToken 缓存相关
	__obf_dd3a0e30c5e2a583 string
	__obf_6c5f537c9b7fd338 int64 // 过期时间戳 (秒)
	__obf_523201c92a85006f sync.RWMutex
}

var (
	__obf_5b872999f9666ac3 *Program
	__obf_3c7bb163e704e8d1 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_d6e1b30072c85cfb *Config) (*Program, error) {
	var __obf_144318b43cb60222 error
	__obf_3c7bb163e704e8d1.Do(func() {
		__obf_5b872999f9666ac3 = &Program{
			__obf_66fee4aa5605d761: __obf_d6e1b30072c85cfb,
			__obf_39e54c9148149a72: internal.NewClient(),
		}
	})
	return __obf_5b872999f9666ac3, __obf_144318b43cb60222
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
func (__obf_a6a4e7eeb91f7d43 *Program) GetAccessToken() (string, error) {
	__obf_a6a4e7eeb91f7d43.__obf_523201c92a85006f.RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_a6a4e7eeb91f7d43.__obf_dd3a0e30c5e2a583 != "" && time.Now().Unix() < __obf_a6a4e7eeb91f7d43.__obf_6c5f537c9b7fd338 {
		__obf_884db4d31946b5b7 := __obf_a6a4e7eeb91f7d43.__obf_dd3a0e30c5e2a583
		__obf_a6a4e7eeb91f7d43.__obf_523201c92a85006f.RUnlock()
		return __obf_884db4d31946b5b7, nil
	}
	__obf_a6a4e7eeb91f7d43.__obf_523201c92a85006f.RUnlock()

	// 如果过期或未获取，则加写锁进行获取
	__obf_a6a4e7eeb91f7d43.__obf_523201c92a85006f.Lock()
	defer __obf_a6a4e7eeb91f7d43.__obf_523201c92a85006f.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_a6a4e7eeb91f7d43.__obf_dd3a0e30c5e2a583 != "" && time.Now().Unix() < __obf_a6a4e7eeb91f7d43.__obf_6c5f537c9b7fd338 {
		return __obf_a6a4e7eeb91f7d43.__obf_dd3a0e30c5e2a583, nil
	}

	var (
		__obf_e1d0a193414068a4 AccessTokenResponse
		__obf_144318b43cb60222 error
	)
	if __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.UseStableAK {
		// 使用稳定版的 AccessToken 接口
		__obf_144318b43cb60222 = __obf_a6a4e7eeb91f7d43.__obf_39e54c9148149a72.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppID,
			"secret":     __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppSecret,
		}, &__obf_e1d0a193414068a4)
	} else {
		// 使用普通的 AccessToken 接口
		__obf_144318b43cb60222 = __obf_a6a4e7eeb91f7d43.__obf_39e54c9148149a72.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppID,
			"secret":     __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppSecret,
		}, &__obf_e1d0a193414068a4)
	}
	if __obf_144318b43cb60222 != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_144318b43cb60222)
	}

	if __obf_e1d0a193414068a4.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_e1d0a193414068a4.ErrCode, __obf_e1d0a193414068a4.ErrMsg)
	}

	// 更新缓存
	__obf_a6a4e7eeb91f7d43.__obf_dd3a0e30c5e2a583 = __obf_e1d0a193414068a4.AccessToken
	// 提前10分钟过期，给刷新留出余地
	__obf_a6a4e7eeb91f7d43.__obf_6c5f537c9b7fd338 = time.Now().Unix() + int64(__obf_e1d0a193414068a4.ExpiresIn) - 600

	return __obf_a6a4e7eeb91f7d43.__obf_dd3a0e30c5e2a583, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_a6a4e7eeb91f7d43 *Program) Code2Session(__obf_db9c5dc62d37ccfb string) (*Code2SessionResult, error) {
	__obf_a94b839080418930 := map[string]string{
		"appid":      __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppID,
		"secret":     __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppSecret,
		"js_code":    __obf_db9c5dc62d37ccfb,
		"grant_type": "authorization_code",
	}

	var __obf_e1d0a193414068a4 Code2SessionResult
	__obf_144318b43cb60222 := __obf_a6a4e7eeb91f7d43.__obf_39e54c9148149a72.Get(internal.MiniprogramCode2SessionURL, __obf_a94b839080418930, &__obf_e1d0a193414068a4)
	if __obf_144318b43cb60222 != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_144318b43cb60222)
	}

	if __obf_e1d0a193414068a4.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_e1d0a193414068a4.ErrCode, __obf_e1d0a193414068a4.ErrMsg)
	}
	return &__obf_e1d0a193414068a4, nil
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
func (__obf_a6a4e7eeb91f7d43 *Program) GetPhoneNumber(__obf_ad542397d244d054 string) (*PhoneNumberData, error) {
	__obf_dd3a0e30c5e2a583, __obf_144318b43cb60222 := __obf_a6a4e7eeb91f7d43.GetAccessToken()
	if __obf_144318b43cb60222 != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_144318b43cb60222)
	}

	__obf_613b0bdd4011e41e := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_dd3a0e30c5e2a583)

	var __obf_e1d0a193414068a4 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_144318b43cb60222 = __obf_a6a4e7eeb91f7d43.__obf_39e54c9148149a72.PostJSON(
		__obf_613b0bdd4011e41e,
		map[string]any{
			"code": __obf_ad542397d244d054,
		},
		&__obf_e1d0a193414068a4,
	)
	if __obf_144318b43cb60222 != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_144318b43cb60222)
	}

	if __obf_e1d0a193414068a4.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_e1d0a193414068a4.ErrCode, __obf_e1d0a193414068a4.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_e1d0a193414068a4.PhoneNumberData.Watermark.AppID != __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)",
			__obf_e1d0a193414068a4.PhoneNumberData.Watermark.AppID, __obf_a6a4e7eeb91f7d43.__obf_66fee4aa5605d761.AppID)
	}

	return &(__obf_e1d0a193414068a4.PhoneNumberData), nil
}
