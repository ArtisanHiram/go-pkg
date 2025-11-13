package __obf_721a4aff228e6809

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

func (__obf_65f0e20120194aca *Options) Send(__obf_06391ab73b490165 string, __obf_9b827ac758d5ff26 string) error {
	var __obf_0b5f1501b5928b44 Request
	__obf_0b5f1501b5928b44.Format = "JSON"
	__obf_0b5f1501b5928b44.Version = "2017-05-25"
	__obf_0b5f1501b5928b44.AccessKeyId = __obf_65f0e20120194aca.AccessKeyID
	__obf_0b5f1501b5928b44.SignatureMethod = "HMAC-SHA1"
	__obf_0b5f1501b5928b44.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_0b5f1501b5928b44.SignatureVersion = "1.0"
	__obf_0b5f1501b5928b44.SignatureNonce = __obf_5fc7e9a6b9ce2950()

	__obf_0b5f1501b5928b44.Action = "SendSms"
	__obf_0b5f1501b5928b44.SignName = __obf_65f0e20120194aca.SignName
	__obf_0b5f1501b5928b44.TemplateCode = __obf_65f0e20120194aca.TemplateCode
	__obf_0b5f1501b5928b44.PhoneNumbers = __obf_06391ab73b490165
	__obf_0b5f1501b5928b44.TemplateParam = __obf_9b827ac758d5ff26
	__obf_0b5f1501b5928b44.RegionId = "cn-hangzhou"

	__obf_4bda816a93627fd9 := __obf_0b5f1501b5928b44.ComposeUrl("GET", __obf_65f0e20120194aca.AccessSecret)
	var __obf_4f91759e18b6e040 *http.Response
	var __obf_609bc6214c29016a error
	__obf_4f91759e18b6e040, __obf_609bc6214c29016a = http.Get(__obf_4bda816a93627fd9)
	if __obf_609bc6214c29016a != nil {
		return __obf_609bc6214c29016a
	}
	var __obf_50a119c22c210e18 []byte
	__obf_50a119c22c210e18, __obf_609bc6214c29016a = io.ReadAll(__obf_4f91759e18b6e040.Body)
	if __obf_609bc6214c29016a != nil {
		return __obf_609bc6214c29016a
	}
	_m := make(map[string](string))
	__obf_609bc6214c29016a = json.Unmarshal(__obf_50a119c22c210e18, &_m)
	if __obf_609bc6214c29016a != nil {
		return __obf_609bc6214c29016a
	}
	__obf_8dfe15ffce61aa7d, __obf_a6662bcf8093e2ff := _m["Message"]
	if __obf_a6662bcf8093e2ff && strings.ToUpper(__obf_8dfe15ffce61aa7d) == "OK" {
		return nil
	}
	if __obf_a6662bcf8093e2ff {
		return errors.New(__obf_8dfe15ffce61aa7d)
	}
	return errors.New("send sms error")
}
