package __obf_cb62198a5f8c0e2c

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

func (__obf_db6cc726127a0d10 *Options) Send(__obf_5c23aa0719bb95d7 string, __obf_763d10aa72c42135 string) error {
	var __obf_7f4f808f8336e164 Request
	__obf_7f4f808f8336e164.Format = "JSON"
	__obf_7f4f808f8336e164.Version = "2017-05-25"
	__obf_7f4f808f8336e164.AccessKeyId = __obf_db6cc726127a0d10.AccessKeyID
	__obf_7f4f808f8336e164.SignatureMethod = "HMAC-SHA1"
	__obf_7f4f808f8336e164.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_7f4f808f8336e164.SignatureVersion = "1.0"
	__obf_7f4f808f8336e164.SignatureNonce = __obf_1f0e3362918c2912()

	__obf_7f4f808f8336e164.Action = "SendSms"
	__obf_7f4f808f8336e164.SignName = __obf_db6cc726127a0d10.SignName
	__obf_7f4f808f8336e164.TemplateCode = __obf_db6cc726127a0d10.TemplateCode
	__obf_7f4f808f8336e164.PhoneNumbers = __obf_5c23aa0719bb95d7
	__obf_7f4f808f8336e164.TemplateParam = __obf_763d10aa72c42135
	__obf_7f4f808f8336e164.RegionId = "cn-hangzhou"

	__obf_b71ea216ac9f4ad6 := __obf_7f4f808f8336e164.ComposeUrl("GET", __obf_db6cc726127a0d10.AccessSecret)
	var __obf_d2feb47c8fb4af9b *http.Response
	var __obf_e64b794f113acf71 error
	__obf_d2feb47c8fb4af9b, __obf_e64b794f113acf71 = http.Get(__obf_b71ea216ac9f4ad6)
	if __obf_e64b794f113acf71 != nil {
		return __obf_e64b794f113acf71
	}
	var __obf_bd05f014be491fb8 []byte
	__obf_bd05f014be491fb8, __obf_e64b794f113acf71 = io.ReadAll(__obf_d2feb47c8fb4af9b.Body)
	if __obf_e64b794f113acf71 != nil {
		return __obf_e64b794f113acf71
	}
	_m := make(map[string](string))
	__obf_e64b794f113acf71 = json.Unmarshal(__obf_bd05f014be491fb8, &_m)
	if __obf_e64b794f113acf71 != nil {
		return __obf_e64b794f113acf71
	}
	__obf_3a9dbb5f4d96de78, __obf_de3cb900cfcd35a3 := _m["Message"]
	if __obf_de3cb900cfcd35a3 && strings.ToUpper(__obf_3a9dbb5f4d96de78) == "OK" {
		return nil
	}
	if __obf_de3cb900cfcd35a3 {
		return errors.New(__obf_3a9dbb5f4d96de78)
	}
	return errors.New("send sms error")
}
