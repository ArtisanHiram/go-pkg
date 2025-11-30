package __obf_d9d9d49338ae6cf8

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_e442cb7089667e2e *Config
	__obf_eccdc89abbff2ca2 *internal.HttpClient
	__obf_4be0ed09ab6a7f70 string // AccessToken 缓存相关
	__obf_f46e5e1ef5f44321 int64
	__obf_5d374ac6aee7450d sync. // 过期时间戳 (秒)
				RWMutex
}

var (
	__obf_3027412d2122a955 *Program
	__obf_bf401f0a3a9fedb7 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_054715d0b1090cdb *Config) (*Program, error) {
	var __obf_22775626b7719eb3 error
	__obf_bf401f0a3a9fedb7.
		Do(func() {
			__obf_3027412d2122a955 = &Program{__obf_e442cb7089667e2e: __obf_054715d0b1090cdb, __obf_eccdc89abbff2ca2: internal.NewClient()}
		})
	return __obf_3027412d2122a955, __obf_22775626b7719eb3
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
func (__obf_3f809229dd13e416 *Program) GetAccessToken() (string, error) {
	__obf_3f809229dd13e416.__obf_5d374ac6aee7450d.
		RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_3f809229dd13e416.__obf_4be0ed09ab6a7f70 != "" && time.Now().Unix() < __obf_3f809229dd13e416.__obf_f46e5e1ef5f44321 {
		__obf_83fcebb007793f38 := __obf_3f809229dd13e416.__obf_4be0ed09ab6a7f70
		__obf_3f809229dd13e416.__obf_5d374ac6aee7450d.
			RUnlock()
		return __obf_83fcebb007793f38, nil
	}
	__obf_3f809229dd13e416.__obf_5d374ac6aee7450d.
		RUnlock()
	__obf_3f809229dd13e416.

		// 如果过期或未获取，则加写锁进行获取
		__obf_5d374ac6aee7450d.
		Lock()
	defer __obf_3f809229dd13e416.__obf_5d374ac6aee7450d.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_3f809229dd13e416.__obf_4be0ed09ab6a7f70 != "" && time.Now().Unix() < __obf_3f809229dd13e416.__obf_f46e5e1ef5f44321 {
		return __obf_3f809229dd13e416.__obf_4be0ed09ab6a7f70, nil
	}

	var (
		__obf_c706947367eee9c7 AccessTokenResponse
		__obf_22775626b7719eb3 error
	)
	if __obf_3f809229dd13e416.__obf_e442cb7089667e2e.UseStableAK {
		__obf_22775626b7719eb3 = // 使用稳定版的 AccessToken 接口
			__obf_3f809229dd13e416.__obf_eccdc89abbff2ca2.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppID,
				"secret":     __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppSecret,
			}, &__obf_c706947367eee9c7)
	} else {
		__obf_22775626b7719eb3 = // 使用普通的 AccessToken 接口
			__obf_3f809229dd13e416.__obf_eccdc89abbff2ca2.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppID,
				"secret":     __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppSecret,
			}, &__obf_c706947367eee9c7)
	}
	if __obf_22775626b7719eb3 != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_22775626b7719eb3)
	}

	if __obf_c706947367eee9c7.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_c706947367eee9c7.ErrCode, __obf_c706947367eee9c7.ErrMsg)
	}
	__obf_3f809229dd13e416.

		// 更新缓存
		__obf_4be0ed09ab6a7f70 = __obf_c706947367eee9c7.AccessToken
	__obf_3f809229dd13e416.
		// 提前10分钟过期，给刷新留出余地
		__obf_f46e5e1ef5f44321 = time.Now().Unix() + int64(__obf_c706947367eee9c7.ExpiresIn) - 600

	return __obf_3f809229dd13e416.__obf_4be0ed09ab6a7f70, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_3f809229dd13e416 *Program) Code2Session(__obf_a3a02ebc634995d6 string) (*Code2SessionResult, error) {
	__obf_52239c6899081efd := map[string]string{
		"appid":      __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppID,
		"secret":     __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppSecret,
		"js_code":    __obf_a3a02ebc634995d6,
		"grant_type": "authorization_code",
	}

	var __obf_c706947367eee9c7 Code2SessionResult
	__obf_22775626b7719eb3 := __obf_3f809229dd13e416.__obf_eccdc89abbff2ca2.Get(internal.MiniprogramCode2SessionURL, __obf_52239c6899081efd, &__obf_c706947367eee9c7)
	if __obf_22775626b7719eb3 != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_22775626b7719eb3)
	}

	if __obf_c706947367eee9c7.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_c706947367eee9c7.ErrCode, __obf_c706947367eee9c7.ErrMsg)
	}
	return &__obf_c706947367eee9c7, nil
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
func (__obf_3f809229dd13e416 *Program) GetPhoneNumber(__obf_873000c212f08074 string) (*PhoneNumberData, error) {
	__obf_4be0ed09ab6a7f70, __obf_22775626b7719eb3 := __obf_3f809229dd13e416.GetAccessToken()
	if __obf_22775626b7719eb3 != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_22775626b7719eb3)
	}
	__obf_37414721d290b9c9 := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_4be0ed09ab6a7f70)

	var __obf_c706947367eee9c7 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_22775626b7719eb3 = __obf_3f809229dd13e416.__obf_eccdc89abbff2ca2.PostJSON(__obf_37414721d290b9c9, map[string]any{
		"code": __obf_873000c212f08074,
	},
		&__obf_c706947367eee9c7,
	)
	if __obf_22775626b7719eb3 != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_22775626b7719eb3)
	}

	if __obf_c706947367eee9c7.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_c706947367eee9c7.ErrCode, __obf_c706947367eee9c7.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_c706947367eee9c7.PhoneNumberData.Watermark.AppID != __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)", __obf_c706947367eee9c7.
			PhoneNumberData.Watermark.AppID, __obf_3f809229dd13e416.__obf_e442cb7089667e2e.AppID)
	}

	return &(__obf_c706947367eee9c7.PhoneNumberData), nil
}
