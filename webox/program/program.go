package __obf_bffb4474e614115e

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_a34d462e5478990f *Config
	__obf_de36a32b671e188b *internal.HttpClient
	// AccessToken 缓存相关
	__obf_bb27cdbd9182c9da string
	__obf_e14f14fd02ff8d3e int64 // 过期时间戳 (秒)
	__obf_47a647bba59d8e45 sync.RWMutex
}

var (
	__obf_edd449872ed9500d *Program
	__obf_4c92f5f130582e3b sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_113f2e098f779e06 *Config) (*Program, error) {
	var __obf_17cbe2dddea3277e error
	__obf_4c92f5f130582e3b.Do(func() {
		__obf_edd449872ed9500d = &Program{
			__obf_a34d462e5478990f: __obf_113f2e098f779e06,
			__obf_de36a32b671e188b: internal.NewClient(),
		}
	})
	return __obf_edd449872ed9500d, __obf_17cbe2dddea3277e
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
func (__obf_84e2f74053fbdd48 *Program) GetAccessToken() (string, error) {
	__obf_84e2f74053fbdd48.__obf_47a647bba59d8e45.RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_84e2f74053fbdd48.__obf_bb27cdbd9182c9da != "" && time.Now().Unix() < __obf_84e2f74053fbdd48.__obf_e14f14fd02ff8d3e {
		__obf_5d7fad2f2a8d0c59 := __obf_84e2f74053fbdd48.__obf_bb27cdbd9182c9da
		__obf_84e2f74053fbdd48.__obf_47a647bba59d8e45.RUnlock()
		return __obf_5d7fad2f2a8d0c59, nil
	}
	__obf_84e2f74053fbdd48.__obf_47a647bba59d8e45.RUnlock()

	// 如果过期或未获取，则加写锁进行获取
	__obf_84e2f74053fbdd48.__obf_47a647bba59d8e45.Lock()
	defer __obf_84e2f74053fbdd48.__obf_47a647bba59d8e45.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_84e2f74053fbdd48.__obf_bb27cdbd9182c9da != "" && time.Now().Unix() < __obf_84e2f74053fbdd48.__obf_e14f14fd02ff8d3e {
		return __obf_84e2f74053fbdd48.__obf_bb27cdbd9182c9da, nil
	}

	var (
		__obf_613f810b3a9b4ee7 AccessTokenResponse
		__obf_17cbe2dddea3277e error
	)
	if __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.UseStableAK {
		// 使用稳定版的 AccessToken 接口
		__obf_17cbe2dddea3277e = __obf_84e2f74053fbdd48.__obf_de36a32b671e188b.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppID,
			"secret":     __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppSecret,
		}, &__obf_613f810b3a9b4ee7)
	} else {
		// 使用普通的 AccessToken 接口
		__obf_17cbe2dddea3277e = __obf_84e2f74053fbdd48.__obf_de36a32b671e188b.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppID,
			"secret":     __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppSecret,
		}, &__obf_613f810b3a9b4ee7)
	}
	if __obf_17cbe2dddea3277e != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_17cbe2dddea3277e)
	}

	if __obf_613f810b3a9b4ee7.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_613f810b3a9b4ee7.ErrCode, __obf_613f810b3a9b4ee7.ErrMsg)
	}

	// 更新缓存
	__obf_84e2f74053fbdd48.__obf_bb27cdbd9182c9da = __obf_613f810b3a9b4ee7.AccessToken
	// 提前10分钟过期，给刷新留出余地
	__obf_84e2f74053fbdd48.__obf_e14f14fd02ff8d3e = time.Now().Unix() + int64(__obf_613f810b3a9b4ee7.ExpiresIn) - 600

	return __obf_84e2f74053fbdd48.__obf_bb27cdbd9182c9da, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_84e2f74053fbdd48 *Program) Code2Session(__obf_5c62e91c6747fd74 string) (*Code2SessionResult, error) {
	__obf_a095030dbc26722d := map[string]string{
		"appid":      __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppID,
		"secret":     __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppSecret,
		"js_code":    __obf_5c62e91c6747fd74,
		"grant_type": "authorization_code",
	}

	var __obf_613f810b3a9b4ee7 Code2SessionResult
	__obf_17cbe2dddea3277e := __obf_84e2f74053fbdd48.__obf_de36a32b671e188b.Get(internal.MiniprogramCode2SessionURL, __obf_a095030dbc26722d, &__obf_613f810b3a9b4ee7)
	if __obf_17cbe2dddea3277e != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_17cbe2dddea3277e)
	}

	if __obf_613f810b3a9b4ee7.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_613f810b3a9b4ee7.ErrCode, __obf_613f810b3a9b4ee7.ErrMsg)
	}
	return &__obf_613f810b3a9b4ee7, nil
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
func (__obf_84e2f74053fbdd48 *Program) GetPhoneNumber(__obf_e88448cb75240bd6 string) (*PhoneNumberData, error) {
	__obf_bb27cdbd9182c9da, __obf_17cbe2dddea3277e := __obf_84e2f74053fbdd48.GetAccessToken()
	if __obf_17cbe2dddea3277e != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_17cbe2dddea3277e)
	}

	__obf_086910e3ec2f989c := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_bb27cdbd9182c9da)

	var __obf_613f810b3a9b4ee7 struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_17cbe2dddea3277e = __obf_84e2f74053fbdd48.__obf_de36a32b671e188b.PostJSON(
		__obf_086910e3ec2f989c,
		map[string]any{
			"code": __obf_e88448cb75240bd6,
		},
		&__obf_613f810b3a9b4ee7,
	)
	if __obf_17cbe2dddea3277e != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_17cbe2dddea3277e)
	}

	if __obf_613f810b3a9b4ee7.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_613f810b3a9b4ee7.ErrCode, __obf_613f810b3a9b4ee7.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_613f810b3a9b4ee7.PhoneNumberData.Watermark.AppID != __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)",
			__obf_613f810b3a9b4ee7.PhoneNumberData.Watermark.AppID, __obf_84e2f74053fbdd48.__obf_a34d462e5478990f.AppID)
	}

	return &(__obf_613f810b3a9b4ee7.PhoneNumberData), nil
}
