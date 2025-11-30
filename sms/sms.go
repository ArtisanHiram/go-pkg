package __obf_dc51e1c30a41549a

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

func (__obf_a57bc87d82b489a7 *Options) Send(__obf_088d7899d07eceed string, __obf_6fb58df1c2665ac7 string) error {
	var __obf_dd7fa50e6b28960a Request
	__obf_dd7fa50e6b28960a.
		Format = "JSON"
	__obf_dd7fa50e6b28960a.
		Version = "2017-05-25"
	__obf_dd7fa50e6b28960a.
		AccessKeyId = __obf_a57bc87d82b489a7.AccessKeyID
	__obf_dd7fa50e6b28960a.
		SignatureMethod = "HMAC-SHA1"
	__obf_dd7fa50e6b28960a.
		Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_dd7fa50e6b28960a.
		SignatureVersion = "1.0"
	__obf_dd7fa50e6b28960a.
		SignatureNonce = __obf_97dd4bb2856c2c3b()
	__obf_dd7fa50e6b28960a.
		Action = "SendSms"
	__obf_dd7fa50e6b28960a.
		SignName = __obf_a57bc87d82b489a7.SignName
	__obf_dd7fa50e6b28960a.
		TemplateCode = __obf_a57bc87d82b489a7.TemplateCode
	__obf_dd7fa50e6b28960a.
		PhoneNumbers = __obf_088d7899d07eceed
	__obf_dd7fa50e6b28960a.
		TemplateParam = __obf_6fb58df1c2665ac7
	__obf_dd7fa50e6b28960a.
		RegionId = "cn-hangzhou"
	__obf_d9ce63f5964c1edc := __obf_dd7fa50e6b28960a.ComposeUrl("GET", __obf_a57bc87d82b489a7.AccessSecret)
	var __obf_0609dc70193616d6 *http.Response
	var __obf_107183985243bfb9 error
	__obf_0609dc70193616d6, __obf_107183985243bfb9 = http.Get(__obf_d9ce63f5964c1edc)
	if __obf_107183985243bfb9 != nil {
		return __obf_107183985243bfb9
	}
	var __obf_199e588fe0fa9f9c []byte
	__obf_199e588fe0fa9f9c, __obf_107183985243bfb9 = io.ReadAll(__obf_0609dc70193616d6.Body)
	if __obf_107183985243bfb9 != nil {
		return __obf_107183985243bfb9
	}
	_m := make(map[string](string))
	__obf_107183985243bfb9 = json.Unmarshal(__obf_199e588fe0fa9f9c, &_m)
	if __obf_107183985243bfb9 != nil {
		return __obf_107183985243bfb9
	}
	__obf_92a00fb6d8299853, __obf_385848b9437ded4a := _m["Message"]
	if __obf_385848b9437ded4a && strings.ToUpper(__obf_92a00fb6d8299853) == "OK" {
		return nil
	}
	if __obf_385848b9437ded4a {
		return errors.New(__obf_92a00fb6d8299853)
	}
	return errors.New("send sms error")
}
