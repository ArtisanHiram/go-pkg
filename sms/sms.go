package __obf_2457be4436ec6bd6

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

func (__obf_8e3ae9f43d7e8319 *Options) Send(__obf_570a43d1e6e4cdd0 string, __obf_667bc40f3af44edf string) error {
	var __obf_2f0ce35fd083a21e Request
	__obf_2f0ce35fd083a21e.Format = "JSON"
	__obf_2f0ce35fd083a21e.Version = "2017-05-25"
	__obf_2f0ce35fd083a21e.AccessKeyId = __obf_8e3ae9f43d7e8319.AccessKeyID
	__obf_2f0ce35fd083a21e.SignatureMethod = "HMAC-SHA1"
	__obf_2f0ce35fd083a21e.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_2f0ce35fd083a21e.SignatureVersion = "1.0"
	__obf_2f0ce35fd083a21e.SignatureNonce = __obf_7f5e91beda1d5df9()

	__obf_2f0ce35fd083a21e.Action = "SendSms"
	__obf_2f0ce35fd083a21e.SignName = __obf_8e3ae9f43d7e8319.SignName
	__obf_2f0ce35fd083a21e.TemplateCode = __obf_8e3ae9f43d7e8319.TemplateCode
	__obf_2f0ce35fd083a21e.PhoneNumbers = __obf_570a43d1e6e4cdd0
	__obf_2f0ce35fd083a21e.TemplateParam = __obf_667bc40f3af44edf
	__obf_2f0ce35fd083a21e.RegionId = "cn-hangzhou"

	__obf_01dd296b6ddb0b1e := __obf_2f0ce35fd083a21e.ComposeUrl("GET", __obf_8e3ae9f43d7e8319.AccessSecret)
	var __obf_d345221f9cc9fc0e *http.Response
	var __obf_06ffb474ec126f34 error
	__obf_d345221f9cc9fc0e, __obf_06ffb474ec126f34 = http.Get(__obf_01dd296b6ddb0b1e)
	if __obf_06ffb474ec126f34 != nil {
		return __obf_06ffb474ec126f34
	}
	var __obf_a0e2128bca223336 []byte
	__obf_a0e2128bca223336, __obf_06ffb474ec126f34 = io.ReadAll(__obf_d345221f9cc9fc0e.Body)
	if __obf_06ffb474ec126f34 != nil {
		return __obf_06ffb474ec126f34
	}
	_m := make(map[string](string))
	__obf_06ffb474ec126f34 = json.Unmarshal(__obf_a0e2128bca223336, &_m)
	if __obf_06ffb474ec126f34 != nil {
		return __obf_06ffb474ec126f34
	}
	__obf_d139527b84ee780b, __obf_9e3870b56d9e8320 := _m["Message"]
	if __obf_9e3870b56d9e8320 && strings.ToUpper(__obf_d139527b84ee780b) == "OK" {
		return nil
	}
	if __obf_9e3870b56d9e8320 {
		return errors.New(__obf_d139527b84ee780b)
	}
	return errors.New("send sms error")
}
