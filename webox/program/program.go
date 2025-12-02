package __obf_21262968b7c0ea9c

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_3428bab56bcb3cba *Config
	__obf_54745622f34c4bae *internal.HttpClient
	__obf_615e89148e58f557 string // AccessToken 缓存相关
	__obf_0bf9542cba33f703 int64
	__obf_cb0467894425a057 sync. // 过期时间戳 (秒)
				RWMutex
}

var (
	__obf_92670df82abce0bd *Program
	__obf_d8719c45345f0973 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_0c84c2fed374d121 *Config) (*Program, error) {
	var __obf_55f36149a3360a9d error
	__obf_d8719c45345f0973.
		Do(func() {
			__obf_92670df82abce0bd = &Program{__obf_3428bab56bcb3cba: __obf_0c84c2fed374d121, __obf_54745622f34c4bae: internal.NewClient()}
		})
	return __obf_92670df82abce0bd, __obf_55f36149a3360a9d
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
func (__obf_cd1a2145e771fe71 *Program) GetAccessToken() (string, error) {
	__obf_cd1a2145e771fe71.__obf_cb0467894425a057.
		RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_cd1a2145e771fe71.__obf_615e89148e58f557 != "" && time.Now().Unix() < __obf_cd1a2145e771fe71.__obf_0bf9542cba33f703 {
		__obf_dbad726022841dc1 := __obf_cd1a2145e771fe71.__obf_615e89148e58f557
		__obf_cd1a2145e771fe71.__obf_cb0467894425a057.
			RUnlock()
		return __obf_dbad726022841dc1, nil
	}
	__obf_cd1a2145e771fe71.__obf_cb0467894425a057.
		RUnlock()
	__obf_cd1a2145e771fe71.

		// 如果过期或未获取，则加写锁进行获取
		__obf_cb0467894425a057.
		Lock()
	defer __obf_cd1a2145e771fe71.__obf_cb0467894425a057.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_cd1a2145e771fe71.__obf_615e89148e58f557 != "" && time.Now().Unix() < __obf_cd1a2145e771fe71.__obf_0bf9542cba33f703 {
		return __obf_cd1a2145e771fe71.__obf_615e89148e58f557, nil
	}

	var (
		__obf_a603174c09833ac5 AccessTokenResponse
		__obf_55f36149a3360a9d error
	)
	if __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.UseStableAK {
		__obf_55f36149a3360a9d = // 使用稳定版的 AccessToken 接口
			__obf_cd1a2145e771fe71.__obf_54745622f34c4bae.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppID,
				"secret":     __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppSecret,
			}, &__obf_a603174c09833ac5)
	} else {
		__obf_55f36149a3360a9d = // 使用普通的 AccessToken 接口
			__obf_cd1a2145e771fe71.__obf_54745622f34c4bae.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
				"grant_type": "client_credential",
				"appid":      __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppID,
				"secret":     __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppSecret,
			}, &__obf_a603174c09833ac5)
	}
	if __obf_55f36149a3360a9d != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_55f36149a3360a9d)
	}

	if __obf_a603174c09833ac5.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_a603174c09833ac5.ErrCode, __obf_a603174c09833ac5.ErrMsg)
	}
	__obf_cd1a2145e771fe71.

		// 更新缓存
		__obf_615e89148e58f557 = __obf_a603174c09833ac5.AccessToken
	__obf_cd1a2145e771fe71.
		// 提前10分钟过期，给刷新留出余地
		__obf_0bf9542cba33f703 = time.Now().Unix() + int64(__obf_a603174c09833ac5.ExpiresIn) - 600

	return __obf_cd1a2145e771fe71.__obf_615e89148e58f557, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_cd1a2145e771fe71 *Program) Code2Session(__obf_348eeed95713d333 string) (*Code2SessionResult, error) {
	__obf_8277ec676c06caef := map[string]string{
		"appid":      __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppID,
		"secret":     __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppSecret,
		"js_code":    __obf_348eeed95713d333,
		"grant_type": "authorization_code",
	}

	var __obf_a603174c09833ac5 Code2SessionResult
	__obf_55f36149a3360a9d := __obf_cd1a2145e771fe71.__obf_54745622f34c4bae.Get(internal.MiniprogramCode2SessionURL, __obf_8277ec676c06caef, &__obf_a603174c09833ac5)
	if __obf_55f36149a3360a9d != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_55f36149a3360a9d)
	}

	if __obf_a603174c09833ac5.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_a603174c09833ac5.ErrCode, __obf_a603174c09833ac5.ErrMsg)
	}
	return &__obf_a603174c09833ac5, nil
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
func (__obf_cd1a2145e771fe71 *Program) GetPhoneNumber(__obf_521ea4c3e19f9bb3 string) (*PhoneNumberData, error) {
	__obf_615e89148e58f557, __obf_55f36149a3360a9d := __obf_cd1a2145e771fe71.GetAccessToken()
	if __obf_55f36149a3360a9d != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_55f36149a3360a9d)
	}
	__obf_63ad56627c0fdc77 := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_615e89148e58f557)

	var __obf_a603174c09833ac5 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_55f36149a3360a9d = __obf_cd1a2145e771fe71.__obf_54745622f34c4bae.PostJSON(__obf_63ad56627c0fdc77, map[string]any{
		"code": __obf_521ea4c3e19f9bb3,
	},
		&__obf_a603174c09833ac5,
	)
	if __obf_55f36149a3360a9d != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_55f36149a3360a9d)
	}

	if __obf_a603174c09833ac5.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_a603174c09833ac5.ErrCode, __obf_a603174c09833ac5.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_a603174c09833ac5.PhoneNumberData.Watermark.AppID != __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)", __obf_a603174c09833ac5.
			PhoneNumberData.Watermark.AppID, __obf_cd1a2145e771fe71.__obf_3428bab56bcb3cba.AppID)
	}

	return &(__obf_a603174c09833ac5.PhoneNumberData), nil
}
