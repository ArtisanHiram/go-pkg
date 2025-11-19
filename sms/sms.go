package __obf_a307862f84d54cc6

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

func (__obf_ac9dc3b2c75524d9 *Options) Send(__obf_2802e83c7f93f1cb string, __obf_b1ccde5dd72f74a4 string) error {
	var __obf_a730795941f2c552 Request
	__obf_a730795941f2c552.Format = "JSON"
	__obf_a730795941f2c552.Version = "2017-05-25"
	__obf_a730795941f2c552.AccessKeyId = __obf_ac9dc3b2c75524d9.AccessKeyID
	__obf_a730795941f2c552.SignatureMethod = "HMAC-SHA1"
	__obf_a730795941f2c552.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_a730795941f2c552.SignatureVersion = "1.0"
	__obf_a730795941f2c552.SignatureNonce = __obf_a49f1d27524bcbdb()

	__obf_a730795941f2c552.Action = "SendSms"
	__obf_a730795941f2c552.SignName = __obf_ac9dc3b2c75524d9.SignName
	__obf_a730795941f2c552.TemplateCode = __obf_ac9dc3b2c75524d9.TemplateCode
	__obf_a730795941f2c552.PhoneNumbers = __obf_2802e83c7f93f1cb
	__obf_a730795941f2c552.TemplateParam = __obf_b1ccde5dd72f74a4
	__obf_a730795941f2c552.RegionId = "cn-hangzhou"

	__obf_af535edd25f00420 := __obf_a730795941f2c552.ComposeUrl("GET", __obf_ac9dc3b2c75524d9.AccessSecret)
	var __obf_81ab536977c44e1c *http.Response
	var __obf_013de50b06d3048a error
	__obf_81ab536977c44e1c, __obf_013de50b06d3048a = http.Get(__obf_af535edd25f00420)
	if __obf_013de50b06d3048a != nil {
		return __obf_013de50b06d3048a
	}
	var __obf_856053960379f1d3 []byte
	__obf_856053960379f1d3, __obf_013de50b06d3048a = io.ReadAll(__obf_81ab536977c44e1c.Body)
	if __obf_013de50b06d3048a != nil {
		return __obf_013de50b06d3048a
	}
	_m := make(map[string](string))
	__obf_013de50b06d3048a = json.Unmarshal(__obf_856053960379f1d3, &_m)
	if __obf_013de50b06d3048a != nil {
		return __obf_013de50b06d3048a
	}
	__obf_66ccbb1864c5d3b3, __obf_d2962737e2d012b1 := _m["Message"]
	if __obf_d2962737e2d012b1 && strings.ToUpper(__obf_66ccbb1864c5d3b3) == "OK" {
		return nil
	}
	if __obf_d2962737e2d012b1 {
		return errors.New(__obf_66ccbb1864c5d3b3)
	}
	return errors.New("send sms error")
}
