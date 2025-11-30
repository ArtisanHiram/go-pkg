package __obf_37e82024588137a9

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

func (__obf_34c060da2e4546df *Options) Send(__obf_b3e3029dd70ecfce string, __obf_837dd689eeaa23e9 string) error {
	var __obf_717ede57932e9fd1 Request
	__obf_717ede57932e9fd1.
		Format = "JSON"
	__obf_717ede57932e9fd1.
		Version = "2017-05-25"
	__obf_717ede57932e9fd1.
		AccessKeyId = __obf_34c060da2e4546df.AccessKeyID
	__obf_717ede57932e9fd1.
		SignatureMethod = "HMAC-SHA1"
	__obf_717ede57932e9fd1.
		Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_717ede57932e9fd1.
		SignatureVersion = "1.0"
	__obf_717ede57932e9fd1.
		SignatureNonce = __obf_b74d37b56b7256e3()
	__obf_717ede57932e9fd1.
		Action = "SendSms"
	__obf_717ede57932e9fd1.
		SignName = __obf_34c060da2e4546df.SignName
	__obf_717ede57932e9fd1.
		TemplateCode = __obf_34c060da2e4546df.TemplateCode
	__obf_717ede57932e9fd1.
		PhoneNumbers = __obf_b3e3029dd70ecfce
	__obf_717ede57932e9fd1.
		TemplateParam = __obf_837dd689eeaa23e9
	__obf_717ede57932e9fd1.
		RegionId = "cn-hangzhou"
	__obf_3800c3b15eda8d3f := __obf_717ede57932e9fd1.ComposeUrl("GET", __obf_34c060da2e4546df.AccessSecret)
	var __obf_ff3e040bb6d30003 *http.Response
	var __obf_a89a7ef35adae91e error
	__obf_ff3e040bb6d30003, __obf_a89a7ef35adae91e = http.Get(__obf_3800c3b15eda8d3f)
	if __obf_a89a7ef35adae91e != nil {
		return __obf_a89a7ef35adae91e
	}
	var __obf_b24fa5c36f71482a []byte
	__obf_b24fa5c36f71482a, __obf_a89a7ef35adae91e = io.ReadAll(__obf_ff3e040bb6d30003.Body)
	if __obf_a89a7ef35adae91e != nil {
		return __obf_a89a7ef35adae91e
	}
	_m := make(map[string](string))
	__obf_a89a7ef35adae91e = json.Unmarshal(__obf_b24fa5c36f71482a, &_m)
	if __obf_a89a7ef35adae91e != nil {
		return __obf_a89a7ef35adae91e
	}
	__obf_7e5cecde9b7efba0, __obf_c147464bf21eb6e3 := _m["Message"]
	if __obf_c147464bf21eb6e3 && strings.ToUpper(__obf_7e5cecde9b7efba0) == "OK" {
		return nil
	}
	if __obf_c147464bf21eb6e3 {
		return errors.New(__obf_7e5cecde9b7efba0)
	}
	return errors.New("send sms error")
}
