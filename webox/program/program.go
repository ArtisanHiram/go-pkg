package __obf_633b893d0a4821b9

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_439de1cf6bdf5991 *Config
	__obf_b583c040f28ea4e8 *internal.HttpClient
	// AccessToken 缓存相关
	__obf_9d464e96113d413d string
	__obf_55f5ff9c016635c0 int64 // 过期时间戳 (秒)
	__obf_28e08d37fd5c8951 sync.RWMutex
}

var (
	__obf_848671284868173e *Program
	__obf_4b115920c422c3f6 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_2f555632e8f720e7 *Config) (*Program, error) {
	var __obf_56287602be52ef04 error
	__obf_4b115920c422c3f6.Do(func() {
		__obf_848671284868173e = &Program{
			__obf_439de1cf6bdf5991: __obf_2f555632e8f720e7,
			__obf_b583c040f28ea4e8: internal.NewClient(),
		}
	})
	return __obf_848671284868173e, __obf_56287602be52ef04
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
func (__obf_7cef49d6192821d4 *Program) GetAccessToken() (string, error) {
	__obf_7cef49d6192821d4.__obf_28e08d37fd5c8951.RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_7cef49d6192821d4.__obf_9d464e96113d413d != "" && time.Now().Unix() < __obf_7cef49d6192821d4.__obf_55f5ff9c016635c0 {
		__obf_396c06cc8ad8810f := __obf_7cef49d6192821d4.__obf_9d464e96113d413d
		__obf_7cef49d6192821d4.__obf_28e08d37fd5c8951.RUnlock()
		return __obf_396c06cc8ad8810f, nil
	}
	__obf_7cef49d6192821d4.__obf_28e08d37fd5c8951.RUnlock()

	// 如果过期或未获取，则加写锁进行获取
	__obf_7cef49d6192821d4.__obf_28e08d37fd5c8951.Lock()
	defer __obf_7cef49d6192821d4.__obf_28e08d37fd5c8951.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_7cef49d6192821d4.__obf_9d464e96113d413d != "" && time.Now().Unix() < __obf_7cef49d6192821d4.__obf_55f5ff9c016635c0 {
		return __obf_7cef49d6192821d4.__obf_9d464e96113d413d, nil
	}

	var (
		__obf_2b64cff76a873a63 AccessTokenResponse
		__obf_56287602be52ef04 error
	)
	if __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.UseStableAK {
		// 使用稳定版的 AccessToken 接口
		__obf_56287602be52ef04 = __obf_7cef49d6192821d4.__obf_b583c040f28ea4e8.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppID,
			"secret":     __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppSecret,
		}, &__obf_2b64cff76a873a63)
	} else {
		// 使用普通的 AccessToken 接口
		__obf_56287602be52ef04 = __obf_7cef49d6192821d4.__obf_b583c040f28ea4e8.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppID,
			"secret":     __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppSecret,
		}, &__obf_2b64cff76a873a63)
	}
	if __obf_56287602be52ef04 != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_56287602be52ef04)
	}

	if __obf_2b64cff76a873a63.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_2b64cff76a873a63.ErrCode, __obf_2b64cff76a873a63.ErrMsg)
	}

	// 更新缓存
	__obf_7cef49d6192821d4.__obf_9d464e96113d413d = __obf_2b64cff76a873a63.AccessToken
	// 提前10分钟过期，给刷新留出余地
	__obf_7cef49d6192821d4.__obf_55f5ff9c016635c0 = time.Now().Unix() + int64(__obf_2b64cff76a873a63.ExpiresIn) - 600

	return __obf_7cef49d6192821d4.__obf_9d464e96113d413d, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_7cef49d6192821d4 *Program) Code2Session(__obf_edb6d8d201907bb4 string) (*Code2SessionResult, error) {
	__obf_819f43e0c408ab65 := map[string]string{
		"appid":      __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppID,
		"secret":     __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppSecret,
		"js_code":    __obf_edb6d8d201907bb4,
		"grant_type": "authorization_code",
	}

	var __obf_2b64cff76a873a63 Code2SessionResult
	__obf_56287602be52ef04 := __obf_7cef49d6192821d4.__obf_b583c040f28ea4e8.Get(internal.MiniprogramCode2SessionURL, __obf_819f43e0c408ab65, &__obf_2b64cff76a873a63)
	if __obf_56287602be52ef04 != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_56287602be52ef04)
	}

	if __obf_2b64cff76a873a63.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_2b64cff76a873a63.ErrCode, __obf_2b64cff76a873a63.ErrMsg)
	}
	return &__obf_2b64cff76a873a63, nil
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
func (__obf_7cef49d6192821d4 *Program) GetPhoneNumber(__obf_0fbda69bbda01fe6 string) (*PhoneNumberData, error) {
	__obf_9d464e96113d413d, __obf_56287602be52ef04 := __obf_7cef49d6192821d4.GetAccessToken()
	if __obf_56287602be52ef04 != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_56287602be52ef04)
	}

	__obf_cd56a266e13f1828 := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_9d464e96113d413d)

	var __obf_2b64cff76a873a63 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_56287602be52ef04 = __obf_7cef49d6192821d4.__obf_b583c040f28ea4e8.PostJSON(
		__obf_cd56a266e13f1828,
		map[string]any{
			"code": __obf_0fbda69bbda01fe6,
		},
		&__obf_2b64cff76a873a63,
	)
	if __obf_56287602be52ef04 != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_56287602be52ef04)
	}

	if __obf_2b64cff76a873a63.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_2b64cff76a873a63.ErrCode, __obf_2b64cff76a873a63.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_2b64cff76a873a63.PhoneNumberData.Watermark.AppID != __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)",
			__obf_2b64cff76a873a63.PhoneNumberData.Watermark.AppID, __obf_7cef49d6192821d4.__obf_439de1cf6bdf5991.AppID)
	}

	return &(__obf_2b64cff76a873a63.PhoneNumberData), nil
}
