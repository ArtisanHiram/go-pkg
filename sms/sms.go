package __obf_dd37d42fbda9c938

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

func (__obf_a6318316fd517cea *Options) Send(__obf_d642df0ecda1085e string, __obf_3b5734d82f6ac84a string) error {
	var __obf_440e94bc0b827720 Request
	__obf_440e94bc0b827720.
		Format = "JSON"
	__obf_440e94bc0b827720.
		Version = "2017-05-25"
	__obf_440e94bc0b827720.
		AccessKeyId = __obf_a6318316fd517cea.AccessKeyID
	__obf_440e94bc0b827720.
		SignatureMethod = "HMAC-SHA1"
	__obf_440e94bc0b827720.
		Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_440e94bc0b827720.
		SignatureVersion = "1.0"
	__obf_440e94bc0b827720.
		SignatureNonce = __obf_4d1620aae5ae13dc()
	__obf_440e94bc0b827720.
		Action = "SendSms"
	__obf_440e94bc0b827720.
		SignName = __obf_a6318316fd517cea.SignName
	__obf_440e94bc0b827720.
		TemplateCode = __obf_a6318316fd517cea.TemplateCode
	__obf_440e94bc0b827720.
		PhoneNumbers = __obf_d642df0ecda1085e
	__obf_440e94bc0b827720.
		TemplateParam = __obf_3b5734d82f6ac84a
	__obf_440e94bc0b827720.
		RegionId = "cn-hangzhou"
	__obf_b68f062bcfc60953 := __obf_440e94bc0b827720.ComposeUrl("GET", __obf_a6318316fd517cea.AccessSecret)
	var __obf_40ae7f61384c9e23 *http.Response
	var __obf_882f5354c13d8836 error
	__obf_40ae7f61384c9e23, __obf_882f5354c13d8836 = http.Get(__obf_b68f062bcfc60953)
	if __obf_882f5354c13d8836 != nil {
		return __obf_882f5354c13d8836
	}
	var __obf_a1dfd1f4475c7fd7 []byte
	__obf_a1dfd1f4475c7fd7, __obf_882f5354c13d8836 = io.ReadAll(__obf_40ae7f61384c9e23.Body)
	if __obf_882f5354c13d8836 != nil {
		return __obf_882f5354c13d8836
	}
	_m := make(map[string](string))
	__obf_882f5354c13d8836 = json.Unmarshal(__obf_a1dfd1f4475c7fd7, &_m)
	if __obf_882f5354c13d8836 != nil {
		return __obf_882f5354c13d8836
	}
	__obf_d66c66365b01a605, __obf_6a3bc0c0444747cc := _m["Message"]
	if __obf_6a3bc0c0444747cc && strings.ToUpper(__obf_d66c66365b01a605) == "OK" {
		return nil
	}
	if __obf_6a3bc0c0444747cc {
		return errors.New(__obf_d66c66365b01a605)
	}
	return errors.New("send sms error")
}
