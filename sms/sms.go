package __obf_0f0912cb4961947f

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type Options struct {
	AccessKeyID  string `yaml:"accesskey-id"`
	AccessSecret string `yaml:"accesskey-secret"`
	SignName     string `yaml:"signname"`      //签名名称
	TemplateCode string `yaml:"template-code"` //模板code
}

func (__obf_f9ec1d2e15ad9f86 *Options) Send(__obf_ffb1946695f25c0b string, __obf_afc1a670727b87e7 string) error {
	var __obf_a66f32824d1c69ab Request
	__obf_a66f32824d1c69ab.
		Format = "JSON"
	__obf_a66f32824d1c69ab.
		Version = "2017-05-25"
	__obf_a66f32824d1c69ab.
		AccessKeyId = __obf_f9ec1d2e15ad9f86.AccessKeyID
	__obf_a66f32824d1c69ab.
		SignatureMethod = "HMAC-SHA1"
	__obf_a66f32824d1c69ab.
		Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_a66f32824d1c69ab.
		SignatureVersion = "1.0"
	__obf_a66f32824d1c69ab.
		SignatureNonce = __obf_d01954fe5f7f19ca()
	__obf_a66f32824d1c69ab.
		Action = "SendSms"
	__obf_a66f32824d1c69ab.
		SignName = __obf_f9ec1d2e15ad9f86.SignName
	__obf_a66f32824d1c69ab.
		TemplateCode = __obf_f9ec1d2e15ad9f86.TemplateCode
	__obf_a66f32824d1c69ab.
		PhoneNumbers = __obf_ffb1946695f25c0b
	__obf_a66f32824d1c69ab.
		TemplateParam = __obf_afc1a670727b87e7
	__obf_a66f32824d1c69ab.
		RegionId = "cn-hangzhou"
	__obf_ea276745d48dc284 := __obf_a66f32824d1c69ab.ComposeUrl("GET", __obf_f9ec1d2e15ad9f86.AccessSecret)
	var __obf_e2ddee3f46207c32 *http.Response
	var __obf_cf598b1a83f71121 error
	__obf_e2ddee3f46207c32, __obf_cf598b1a83f71121 = http.Get(__obf_ea276745d48dc284)
	if __obf_cf598b1a83f71121 != nil {
		return __obf_cf598b1a83f71121
	}
	var __obf_2e053c5802479b30 []byte
	__obf_2e053c5802479b30, __obf_cf598b1a83f71121 = io.ReadAll(__obf_e2ddee3f46207c32.Body)
	if __obf_cf598b1a83f71121 != nil {
		return __obf_cf598b1a83f71121
	}
	_m := make(map[string](string))
	__obf_cf598b1a83f71121 = json.Unmarshal(__obf_2e053c5802479b30, &_m)
	if __obf_cf598b1a83f71121 != nil {
		return __obf_cf598b1a83f71121
	}
	__obf_3529138d6ef9a416, __obf_700450931b113267 := _m["Message"]
	if __obf_700450931b113267 && strings.ToUpper(__obf_3529138d6ef9a416) == "OK" {
		return nil
	}
	if __obf_700450931b113267 {
		return errors.New(__obf_3529138d6ef9a416)
	}
	return errors.New("send sms error")
}
