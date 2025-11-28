package __obf_52bdfa18dc226ac6

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

func (__obf_db4804dbf7dcf319 *Options) Send(__obf_a18109f6351b8d69 string, __obf_848f42953c385e91 string) error {
	var __obf_f86a6142763355f6 Request
	__obf_f86a6142763355f6.Format = "JSON"
	__obf_f86a6142763355f6.Version = "2017-05-25"
	__obf_f86a6142763355f6.AccessKeyId = __obf_db4804dbf7dcf319.AccessKeyID
	__obf_f86a6142763355f6.SignatureMethod = "HMAC-SHA1"
	__obf_f86a6142763355f6.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_f86a6142763355f6.SignatureVersion = "1.0"
	__obf_f86a6142763355f6.SignatureNonce = __obf_126da17986d8b304()

	__obf_f86a6142763355f6.Action = "SendSms"
	__obf_f86a6142763355f6.SignName = __obf_db4804dbf7dcf319.SignName
	__obf_f86a6142763355f6.TemplateCode = __obf_db4804dbf7dcf319.TemplateCode
	__obf_f86a6142763355f6.PhoneNumbers = __obf_a18109f6351b8d69
	__obf_f86a6142763355f6.TemplateParam = __obf_848f42953c385e91
	__obf_f86a6142763355f6.RegionId = "cn-hangzhou"

	__obf_b868c106ee96eb5a := __obf_f86a6142763355f6.ComposeUrl("GET", __obf_db4804dbf7dcf319.AccessSecret)
	var __obf_e32d8a9eb5361d17 *http.Response
	var __obf_57475ef7e317f848 error
	__obf_e32d8a9eb5361d17, __obf_57475ef7e317f848 = http.Get(__obf_b868c106ee96eb5a)
	if __obf_57475ef7e317f848 != nil {
		return __obf_57475ef7e317f848
	}
	var __obf_43f093a485273158 []byte
	__obf_43f093a485273158, __obf_57475ef7e317f848 = io.ReadAll(__obf_e32d8a9eb5361d17.Body)
	if __obf_57475ef7e317f848 != nil {
		return __obf_57475ef7e317f848
	}
	_m := make(map[string](string))
	__obf_57475ef7e317f848 = json.Unmarshal(__obf_43f093a485273158, &_m)
	if __obf_57475ef7e317f848 != nil {
		return __obf_57475ef7e317f848
	}
	__obf_5882d5fe519356f2, __obf_10359abf3efe2975 := _m["Message"]
	if __obf_10359abf3efe2975 && strings.ToUpper(__obf_5882d5fe519356f2) == "OK" {
		return nil
	}
	if __obf_10359abf3efe2975 {
		return errors.New(__obf_5882d5fe519356f2)
	}
	return errors.New("send sms error")
}
