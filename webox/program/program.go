package __obf_ca744dbefd94d505

import (
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Program 小程序模块实例
type Program struct {
	__obf_c59d811a179eb808 *Config
	__obf_f808897a13892981 *internal.HttpClient
	// AccessToken 缓存相关
	__obf_4f514c6095522745 string
	__obf_ec74de41b1c75906 int64 // 过期时间戳 (秒)
	__obf_b8ddfef49a767219 sync.RWMutex
}

var (
	__obf_3c8269a487db594d *Program
	__obf_acadea1ad83929f4 sync.Once
)

// NewMiniprogram 创建并返回小程序模块的单例实例
func NewProgram(__obf_2cd9f0faa4fb6a03 *Config) (*Program, error) {
	var __obf_0ed20f6602459e23 error
	__obf_acadea1ad83929f4.Do(func() {
		__obf_3c8269a487db594d = &Program{
			__obf_c59d811a179eb808: __obf_2cd9f0faa4fb6a03,
			__obf_f808897a13892981: internal.NewClient(),
		}
	})
	return __obf_3c8269a487db594d, __obf_0ed20f6602459e23
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
func (__obf_2abcb71132c3e4af *Program) GetAccessToken() (string, error) {
	__obf_2abcb71132c3e4af.__obf_b8ddfef49a767219.RLock()
	// 检查缓存，如果未过期，直接返回
	if __obf_2abcb71132c3e4af.__obf_4f514c6095522745 != "" && time.Now().Unix() < __obf_2abcb71132c3e4af.__obf_ec74de41b1c75906 {
		__obf_fad6a870d45473a1 := __obf_2abcb71132c3e4af.__obf_4f514c6095522745
		__obf_2abcb71132c3e4af.__obf_b8ddfef49a767219.RUnlock()
		return __obf_fad6a870d45473a1, nil
	}
	__obf_2abcb71132c3e4af.__obf_b8ddfef49a767219.RUnlock()

	// 如果过期或未获取，则加写锁进行获取
	__obf_2abcb71132c3e4af.__obf_b8ddfef49a767219.Lock()
	defer __obf_2abcb71132c3e4af.__obf_b8ddfef49a767219.Unlock()

	// 再次检查，防止多协程竞争时，在等待写锁期间，另一个协程已经更新了token
	if __obf_2abcb71132c3e4af.__obf_4f514c6095522745 != "" && time.Now().Unix() < __obf_2abcb71132c3e4af.__obf_ec74de41b1c75906 {
		return __obf_2abcb71132c3e4af.__obf_4f514c6095522745, nil
	}

	var (
		__obf_39538c6e772adb0b AccessTokenResponse
		__obf_0ed20f6602459e23 error
	)
	if __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.UseStableAK {
		// 使用稳定版的 AccessToken 接口
		__obf_0ed20f6602459e23 = __obf_2abcb71132c3e4af.__obf_f808897a13892981.PostJSON(internal.MiniprogramGetStableAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppID,
			"secret":     __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppSecret,
		}, &__obf_39538c6e772adb0b)
	} else {
		// 使用普通的 AccessToken 接口
		__obf_0ed20f6602459e23 = __obf_2abcb71132c3e4af.__obf_f808897a13892981.Get(internal.MiniprogramGetAccessTokenURL, map[string]string{
			"grant_type": "client_credential",
			"appid":      __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppID,
			"secret":     __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppSecret,
		}, &__obf_39538c6e772adb0b)
	}
	if __obf_0ed20f6602459e23 != nil {
		return "", fmt.Errorf("get access token request failed: %w", __obf_0ed20f6602459e23)
	}

	if __obf_39538c6e772adb0b.ErrCode != 0 {
		return "", fmt.Errorf("get access token API error: %d - %s", __obf_39538c6e772adb0b.ErrCode, __obf_39538c6e772adb0b.ErrMsg)
	}

	// 更新缓存
	__obf_2abcb71132c3e4af.__obf_4f514c6095522745 = __obf_39538c6e772adb0b.AccessToken
	// 提前10分钟过期，给刷新留出余地
	__obf_2abcb71132c3e4af.__obf_ec74de41b1c75906 = time.Now().Unix() + int64(__obf_39538c6e772adb0b.ExpiresIn) - 600

	return __obf_2abcb71132c3e4af.__obf_4f514c6095522745, nil
}

// Code2SessionResult code2Session 接口返回结构
type Code2SessionResult struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	model.ErrResponse
}

// Code2Session 使用 code 换取 openid 和 session_key
func (__obf_2abcb71132c3e4af *Program) Code2Session(__obf_9a56dbd603f07be4 string) (*Code2SessionResult, error) {
	__obf_ef1e6150ff76d26f := map[string]string{
		"appid":      __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppID,
		"secret":     __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppSecret,
		"js_code":    __obf_9a56dbd603f07be4,
		"grant_type": "authorization_code",
	}

	var __obf_39538c6e772adb0b Code2SessionResult
	__obf_0ed20f6602459e23 := __obf_2abcb71132c3e4af.__obf_f808897a13892981.Get(internal.MiniprogramCode2SessionURL, __obf_ef1e6150ff76d26f, &__obf_39538c6e772adb0b)
	if __obf_0ed20f6602459e23 != nil {
		return nil, fmt.Errorf("code2Session request failed: %w", __obf_0ed20f6602459e23)
	}

	if __obf_39538c6e772adb0b.ErrCode != 0 {
		return nil, fmt.Errorf("code2Session API error: %d - %s", __obf_39538c6e772adb0b.ErrCode, __obf_39538c6e772adb0b.ErrMsg)
	}
	return &__obf_39538c6e772adb0b, nil
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
func (__obf_2abcb71132c3e4af *Program) GetPhoneNumber(__obf_2595c4001db496a8 string) (*PhoneNumberData, error) {
	__obf_4f514c6095522745, __obf_0ed20f6602459e23 := __obf_2abcb71132c3e4af.GetAccessToken()
	if __obf_0ed20f6602459e23 != nil {
		return nil, fmt.Errorf("get access token failed: %w", __obf_0ed20f6602459e23)
	}

	__obf_733dd72721c692eb := fmt.Sprintf("%s?access_token=%s", internal.MiniprogramGetPhoneNumberURL, __obf_4f514c6095522745)

	var __obf_39538c6e772adb0b struct {
		model.ErrResponse
		PhoneNumberData `json:"phone_info"`
	}
	__obf_0ed20f6602459e23 = __obf_2abcb71132c3e4af.__obf_f808897a13892981.PostJSON(
		__obf_733dd72721c692eb,
		map[string]any{
			"code": __obf_2595c4001db496a8,
		},
		&__obf_39538c6e772adb0b,
	)
	if __obf_0ed20f6602459e23 != nil {
		return nil, fmt.Errorf("get phone number request failed: %w", __obf_0ed20f6602459e23)
	}

	if __obf_39538c6e772adb0b.ErrCode != 0 {
		return nil, fmt.Errorf("get phone number API error: %d - %s", __obf_39538c6e772adb0b.ErrCode, __obf_39538c6e772adb0b.ErrMsg)
	}

	// 验证AppID是否一致，增强安全性
	if __obf_39538c6e772adb0b.PhoneNumberData.Watermark.AppID != __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppID {
		return nil, fmt.Errorf("app_id in phone number watermark (%s) does not match configured app_id (%s)",
			__obf_39538c6e772adb0b.PhoneNumberData.Watermark.AppID, __obf_2abcb71132c3e4af.__obf_c59d811a179eb808.AppID)
	}

	return &(__obf_39538c6e772adb0b.PhoneNumberData), nil
}
