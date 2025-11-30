package __obf_7147dbcba87e70eb

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

func (__obf_6af19692977a124b *Options) Send(__obf_e76b29731d1409bc string, __obf_a77eb2d09bf7d5e1 string) error {
	var __obf_275a1d33558e050c Request
	__obf_275a1d33558e050c.
		Format = "JSON"
	__obf_275a1d33558e050c.
		Version = "2017-05-25"
	__obf_275a1d33558e050c.
		AccessKeyId = __obf_6af19692977a124b.AccessKeyID
	__obf_275a1d33558e050c.
		SignatureMethod = "HMAC-SHA1"
	__obf_275a1d33558e050c.
		Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_275a1d33558e050c.
		SignatureVersion = "1.0"
	__obf_275a1d33558e050c.
		SignatureNonce = __obf_c1c5408641580711()
	__obf_275a1d33558e050c.
		Action = "SendSms"
	__obf_275a1d33558e050c.
		SignName = __obf_6af19692977a124b.SignName
	__obf_275a1d33558e050c.
		TemplateCode = __obf_6af19692977a124b.TemplateCode
	__obf_275a1d33558e050c.
		PhoneNumbers = __obf_e76b29731d1409bc
	__obf_275a1d33558e050c.
		TemplateParam = __obf_a77eb2d09bf7d5e1
	__obf_275a1d33558e050c.
		RegionId = "cn-hangzhou"
	__obf_4f03c776a5028810 := __obf_275a1d33558e050c.ComposeUrl("GET", __obf_6af19692977a124b.AccessSecret)
	var __obf_006119e71a159d30 *http.Response
	var __obf_e34bd7cedf05ef8e error
	__obf_006119e71a159d30, __obf_e34bd7cedf05ef8e = http.Get(__obf_4f03c776a5028810)
	if __obf_e34bd7cedf05ef8e != nil {
		return __obf_e34bd7cedf05ef8e
	}
	var __obf_14fc605f8bf651d8 []byte
	__obf_14fc605f8bf651d8, __obf_e34bd7cedf05ef8e = io.ReadAll(__obf_006119e71a159d30.Body)
	if __obf_e34bd7cedf05ef8e != nil {
		return __obf_e34bd7cedf05ef8e
	}
	_m := make(map[string](string))
	__obf_e34bd7cedf05ef8e = json.Unmarshal(__obf_14fc605f8bf651d8, &_m)
	if __obf_e34bd7cedf05ef8e != nil {
		return __obf_e34bd7cedf05ef8e
	}
	__obf_c6cb43e25b5ccbd6, __obf_8a9ce927b6f9991b := _m["Message"]
	if __obf_8a9ce927b6f9991b && strings.ToUpper(__obf_c6cb43e25b5ccbd6) == "OK" {
		return nil
	}
	if __obf_8a9ce927b6f9991b {
		return errors.New(__obf_c6cb43e25b5ccbd6)
	}
	return errors.New("send sms error")
}
