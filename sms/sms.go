package __obf_28053862c3d647b8

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

func (__obf_8f1c231dad50b2c0 *Options) Send(__obf_efc5054a806fd242 string, __obf_e76560b255687fd5 string) error {
	var __obf_89afef2b0cffdadf Request
	__obf_89afef2b0cffdadf.Format = "JSON"
	__obf_89afef2b0cffdadf.Version = "2017-05-25"
	__obf_89afef2b0cffdadf.AccessKeyId = __obf_8f1c231dad50b2c0.AccessKeyID
	__obf_89afef2b0cffdadf.SignatureMethod = "HMAC-SHA1"
	__obf_89afef2b0cffdadf.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05Z")
	__obf_89afef2b0cffdadf.SignatureVersion = "1.0"
	__obf_89afef2b0cffdadf.SignatureNonce = __obf_2a2a3aa08c34cdb2()

	__obf_89afef2b0cffdadf.Action = "SendSms"
	__obf_89afef2b0cffdadf.SignName = __obf_8f1c231dad50b2c0.SignName
	__obf_89afef2b0cffdadf.TemplateCode = __obf_8f1c231dad50b2c0.TemplateCode
	__obf_89afef2b0cffdadf.PhoneNumbers = __obf_efc5054a806fd242
	__obf_89afef2b0cffdadf.TemplateParam = __obf_e76560b255687fd5
	__obf_89afef2b0cffdadf.RegionId = "cn-hangzhou"

	__obf_de27e36b69203aea := __obf_89afef2b0cffdadf.ComposeUrl("GET", __obf_8f1c231dad50b2c0.AccessSecret)
	var __obf_4c9f47b636adabcf *http.Response
	var __obf_ae3b97aa1cdbe5ed error
	__obf_4c9f47b636adabcf, __obf_ae3b97aa1cdbe5ed = http.Get(__obf_de27e36b69203aea)
	if __obf_ae3b97aa1cdbe5ed != nil {
		return __obf_ae3b97aa1cdbe5ed
	}
	var __obf_30bbdf520e0c12af []byte
	__obf_30bbdf520e0c12af, __obf_ae3b97aa1cdbe5ed = io.ReadAll(__obf_4c9f47b636adabcf.Body)
	if __obf_ae3b97aa1cdbe5ed != nil {
		return __obf_ae3b97aa1cdbe5ed
	}
	_m := make(map[string](string))
	__obf_ae3b97aa1cdbe5ed = json.Unmarshal(__obf_30bbdf520e0c12af, &_m)
	if __obf_ae3b97aa1cdbe5ed != nil {
		return __obf_ae3b97aa1cdbe5ed
	}
	__obf_8a2ebed189fc955a, __obf_b0a6bc708b5ac22e := _m["Message"]
	if __obf_b0a6bc708b5ac22e && strings.ToUpper(__obf_8a2ebed189fc955a) == "OK" {
		return nil
	}
	if __obf_b0a6bc708b5ac22e {
		return errors.New(__obf_8a2ebed189fc955a)
	}
	return errors.New("send sms error")
}
