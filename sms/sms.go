package __obf_18a392c58f8f4352

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

func (__obf_b254a8b35215b066 *Options) Send(__obf_8cff02b058fcb75f string, __obf_dc162ff8794e46e5 string) error {
	var __obf_ad1bf4caa20481c9 Request
	__obf_ad1bf4caa20481c9.
		Format = "JSON"
	__obf_ad1bf4caa20481c9.
		Version = "2017-05-25"
	__obf_ad1bf4caa20481c9.
		AccessKeyId = __obf_b254a8b35215b066.AccessKeyID
	__obf_ad1bf4caa20481c9.
		SignatureMethod = "HMAC-SHA1"
	__obf_ad1bf4caa20481c9.
		Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_ad1bf4caa20481c9.
		SignatureVersion = "1.0"
	__obf_ad1bf4caa20481c9.
		SignatureNonce = __obf_4a2251d2274fd09f()
	__obf_ad1bf4caa20481c9.
		Action = "SendSms"
	__obf_ad1bf4caa20481c9.
		SignName = __obf_b254a8b35215b066.SignName
	__obf_ad1bf4caa20481c9.
		TemplateCode = __obf_b254a8b35215b066.TemplateCode
	__obf_ad1bf4caa20481c9.
		PhoneNumbers = __obf_8cff02b058fcb75f
	__obf_ad1bf4caa20481c9.
		TemplateParam = __obf_dc162ff8794e46e5
	__obf_ad1bf4caa20481c9.
		RegionId = "cn-hangzhou"
	__obf_f89ac703e578e2a4 := __obf_ad1bf4caa20481c9.ComposeUrl("GET", __obf_b254a8b35215b066.AccessSecret)
	var __obf_14f77a9bd971525b *http.Response
	var __obf_beb566a4ca556e33 error
	__obf_14f77a9bd971525b, __obf_beb566a4ca556e33 = http.Get(__obf_f89ac703e578e2a4)
	if __obf_beb566a4ca556e33 != nil {
		return __obf_beb566a4ca556e33
	}
	var __obf_6eaeb0a8242c32cc []byte
	__obf_6eaeb0a8242c32cc, __obf_beb566a4ca556e33 = io.ReadAll(__obf_14f77a9bd971525b.Body)
	if __obf_beb566a4ca556e33 != nil {
		return __obf_beb566a4ca556e33
	}
	_m := make(map[string](string))
	__obf_beb566a4ca556e33 = json.Unmarshal(__obf_6eaeb0a8242c32cc, &_m)
	if __obf_beb566a4ca556e33 != nil {
		return __obf_beb566a4ca556e33
	}
	__obf_04e01c2cd94de512, __obf_1fca32696f843fc4 := _m["Message"]
	if __obf_1fca32696f843fc4 && strings.ToUpper(__obf_04e01c2cd94de512) == "OK" {
		return nil
	}
	if __obf_1fca32696f843fc4 {
		return errors.New(__obf_04e01c2cd94de512)
	}
	return errors.New("send sms error")
}
