package __obf_c338b4748c62107c

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_4417935ef63331f6 *Config
	__obf_d3392c2ce39da195 *internal.HttpClient
	// AccessToken 缓存相关
	__obf_0acaac2cee17a80e string
	__obf_8eec73473dd03f50 int64 // 过期时间戳 (秒)
	__obf_365c39b60a84586d sync.RWMutex
}

var (
	__obf_57bbf9ae88dbbfb2 *Program
	__obf_61a1d37d54df324e sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_c2e3665f98028d5c *Config) (*Program, error) {
	var __obf_67c1d84e8fb6ba1a error
	__obf_61a1d37d54df324e.Do(func() {
		__obf_57bbf9ae88dbbfb2 = &Program{
			__obf_4417935ef63331f6: __obf_c2e3665f98028d5c,
			__obf_d3392c2ce39da195: internal.NewClient(),
		}
	})
	return __obf_57bbf9ae88dbbfb2, __obf_67c1d84e8fb6ba1a
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
func (__obf_33502c4f6ca494c9 *Program) GetAccessToken() (string, error) {
	__obf_33502c4f6ca494c9.__obf_365c39b60a84586d.RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_33502c4f6ca494c9.__obf_0acaac2cee17a80e != "" && time.Now().Unix() < __obf_33502c4f6ca494c9.__obf_8eec73473dd03f50 {
		__obf_d2c6d2b9d060b08a := __obf_33502c4f6ca494c9.__obf_0acaac2cee17a80e
		__obf_33502c4f6ca494c9.__obf_365c39b60a84586d.RUnlock()
		return __obf_d2c6d2b9d060b08a, nil
	}
	__obf_33502c4f6ca494c9.__obf_365c39b60a84586d.RUnlock()

	// 如果过期或未获取，则加写锁进行获取
	__obf_33502c4f6ca494c9.__obf_365c39b60a84586d.Lock()
	defer __obf_33502c4f6ca494c9.__obf_365c39b60a84586d.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_33502c4f6ca494c9.__obf_0acaac2cee17a80e != "" && time.Now().Unix() < __obf_33502c4f6ca494c9.__obf_8eec73473dd03f50 {
		return __obf_33502c4f6ca494c9.__obf_0acaac2cee17a80e, nil
	}

	var (
		__obf_e2db4d8fe72b443b AccessTokenResponse
		__obf_67c1d84e8fb6ba1a error
	)
	if __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.UseStableAK {
		// 使用稳定版的 AccessToken 接口
		__obf_67c1d84e8fb6ba1a = __obf_33502c4f6ca494c9.__obf_d3392c2ce39da195.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppID,
			"secret":     __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppSecret,
		}, &__obf_e2db4d8fe72b443b)
	} else {
		// 使用普通的 AccessToken 接口
		__obf_67c1d84e8fb6ba1a = __obf_33502c4f6ca494c9.__obf_d3392c2ce39da195.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppID,
			"secret":     __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppSecret,
		}, &__obf_e2db4d8fe72b443b)
	}
	if __obf_67c1d84e8fb6ba1a != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_67c1d84e8fb6ba1a)
	}

	if __obf_e2db4d8fe72b443b.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_e2db4d8fe72b443b.ErrCode, __obf_e2db4d8fe72b443b.ErrMsg)
	}

	// 更新缓存
	__obf_33502c4f6ca494c9.__obf_0acaac2cee17a80e = __obf_e2db4d8fe72b443b.AccessToken
	// 提前10分钟过期，给刷新留出余地
	__obf_33502c4f6ca494c9.__obf_8eec73473dd03f50 = time.Now().Unix() + int64(__obf_e2db4d8fe72b443b.ExpiresIn) - 600

	return __obf_33502c4f6ca494c9.__obf_0acaac2cee17a80e, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_33502c4f6ca494c9 *Program) Code2Session(__obf_e9ed649484e00a1a string) (*Code2SessionResult, error) {
	__obf_5634c124704e842f := map[string]string{
		"appid":      __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppID,
		"secret":     __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppSecret,
		"js_code":    __obf_e9ed649484e00a1a,
		"grant_type": "authorization_code",
	}

	var __obf_e2db4d8fe72b443b Code2SessionResult
	__obf_67c1d84e8fb6ba1a := __obf_33502c4f6ca494c9.__obf_d3392c2ce39da195.Get(internal.MiniprogramCode2SessionURL, __obf_5634c124704e842f, &__obf_e2db4d8fe72b443b)
	if __obf_67c1d84e8fb6ba1a != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_67c1d84e8fb6ba1a)
	}

	if __obf_e2db4d8fe72b443b.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_e2db4d8fe72b443b.ErrCode, __obf_e2db4d8fe72b443b.ErrMsg)
	}
	return &__obf_e2db4d8fe72b443b, nil
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
func (__obf_33502c4f6ca494c9 *Program) GetPhoneNumber(__obf_2fc65b2ecae3373c string) (*PhoneNumberData, error) {
	__obf_0acaac2cee17a80e, __obf_67c1d84e8fb6ba1a := __obf_33502c4f6ca494c9.GetAccessToken()
	if __obf_67c1d84e8fb6ba1a != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_67c1d84e8fb6ba1a)
	}

	__obf_ea6d0dcfda90dae1 := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_0acaac2cee17a80e)

	var __obf_e2db4d8fe72b443b struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_67c1d84e8fb6ba1a = __obf_33502c4f6ca494c9.__obf_d3392c2ce39da195.PostJSON(
		__obf_ea6d0dcfda90dae1,
		map[string]any{
			"code": __obf_2fc65b2ecae3373c,
		},
		&__obf_e2db4d8fe72b443b,
	)
	if __obf_67c1d84e8fb6ba1a != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_67c1d84e8fb6ba1a)
	}

	if __obf_e2db4d8fe72b443b.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_e2db4d8fe72b443b.ErrCode, __obf_e2db4d8fe72b443b.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_e2db4d8fe72b443b.PhoneNumberData.Watermark.AppID != __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)",
			__obf_e2db4d8fe72b443b.PhoneNumberData.Watermark.AppID, __obf_33502c4f6ca494c9.__obf_4417935ef63331f6.AppID)
	}

	return &(__obf_e2db4d8fe72b443b.PhoneNumberData), nil
}
