package __obf_37e82024588137a9

import (
	"bytes"
	"net/url"
	"sort"
	"strings"
)

type Request struct {
	//public
	Format      string //否	返回值的类型，支持JSON与XML。默认为XML
	Version     string //是	API版本号，为日期形式：YYYY-MM-DD，本版本对应为2016-09-27
	AccessKeyId string //是	阿里云颁发给用户的访问服务所用的密钥ID
	//Signature        string //是	签名结果串，关于签名的计算方法，请参见 签名机制。
	SignatureMethod  string //是	签名方式，目前支持HMAC-SHA1
	Timestamp        string //是	请求的时间戳。日期格式按照ISO8601标准表示，并需要使用UTC时间。格式为YYYY-MM-DDThh:mm:ssZ 例如，2015-11-23T04:00:00Z（为北京时间2015年11月23日12点0分0秒）
	SignatureVersion string //是	签名算法版本，目前版本是1.0
	SignatureNonce   string //是	唯一随机数，用于防止网络重放攻击。用户在不同请求间要使用不同的随机数值

	//sms
	Action        string //必须	操作接口名，系统规定参数，取值：SendSms
	SignName      string //必须	管理控制台中配置的短信签名（状态必须是验证通过）
	TemplateCode  string //必须	管理控制台中配置的审核通过的短信模板的模板CODE（状态必须是验证通过）
	PhoneNumbers  string //必须	目标手机号，多个手机号可以逗号分隔
	RegionId      string //必须 API支持的RegionID，如短信API的值为：cn-hangzhou
	TemplateParam string //必选	短信模板中的变量；数字需要转换为字符串；个人用户每个变量长度必须小于15个字符。 例如:短信模板为：“接受短信验证码${no}”,此参数传递{“no”:”123456”}，用户将接收到[短信签名]接受短信验证码123456
}

func (__obf_436884a01449419e *Request) ComposeUrl(__obf_9868ede177159f4a string, __obf_5e3467f51b91dd9a string) string {
	__obf_c67635bd1d068d37 := url.Values{}
	__obf_c67635bd1d068d37.
		Add("AccessKeyId", __obf_436884a01449419e.AccessKeyId)
	__obf_c67635bd1d068d37.
		Add("Action", __obf_436884a01449419e.Action)
	__obf_c67635bd1d068d37.
		Add("Format", __obf_436884a01449419e.Format)
	__obf_c67635bd1d068d37.
		Add("PhoneNumbers", __obf_436884a01449419e.PhoneNumbers)
	__obf_c67635bd1d068d37.
		Add("RegionId", __obf_436884a01449419e.RegionId)
	__obf_c67635bd1d068d37.
		Add("SignName", __obf_436884a01449419e.SignName)
	__obf_c67635bd1d068d37.
		Add("SignatureMethod", __obf_436884a01449419e.SignatureMethod)
	__obf_c67635bd1d068d37.
		Add("SignatureNonce", __obf_436884a01449419e.SignatureNonce)
	__obf_c67635bd1d068d37.
		Add("SignatureVersion", __obf_436884a01449419e.SignatureVersion)
	__obf_c67635bd1d068d37.
		Add("TemplateCode", __obf_436884a01449419e.TemplateCode)
	__obf_c67635bd1d068d37.
		Add("TemplateParam", __obf_436884a01449419e.TemplateParam)
	__obf_c67635bd1d068d37.
		Add("Timestamp", __obf_436884a01449419e.Timestamp)
	__obf_c67635bd1d068d37.
		Add("Version", __obf_436884a01449419e.Version)
	__obf_ed13a557bc5d42d8 := __obf_d1a5d2de53e2f453(__obf_c67635bd1d068d37)
	Signature := __obf_792a3f6a70b79377(__obf_e5ce4633ffe9b986(__obf_ed13a557bc5d42d8, __obf_5e3467f51b91dd9a))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_ed13a557bc5d42d8

	return _url
}

func __obf_d1a5d2de53e2f453(__obf_c67635bd1d068d37 url.Values) string {
	var __obf_ed0046641c04d48f bytes.Buffer
	__obf_8968c6927a2fdc3c := make([]string, 0, len(__obf_c67635bd1d068d37))
	for __obf_bc59592349c859ed := range __obf_c67635bd1d068d37 {
		__obf_8968c6927a2fdc3c = append(__obf_8968c6927a2fdc3c, __obf_bc59592349c859ed)
	}
	sort.Strings(__obf_8968c6927a2fdc3c)
	for _, __obf_bc59592349c859ed := range __obf_8968c6927a2fdc3c {
		if __obf_ed0046641c04d48f.Len() > 0 {
			__obf_ed0046641c04d48f.
				WriteString("&")
		}
		__obf_ed0046641c04d48f.
			WriteString(__obf_792a3f6a70b79377(__obf_bc59592349c859ed))
		__obf_ed0046641c04d48f.
			WriteString("=")
		__obf_ed0046641c04d48f.
			WriteString(__obf_792a3f6a70b79377(__obf_c67635bd1d068d37.Get(__obf_bc59592349c859ed)))
	}
	return __obf_ed0046641c04d48f.String()
}

func __obf_792a3f6a70b79377(__obf_f64fa67723b7e73c string) string {
	__obf_f64fa67723b7e73c = url.QueryEscape(__obf_f64fa67723b7e73c)
	__obf_f64fa67723b7e73c = strings.ReplaceAll(__obf_f64fa67723b7e73c, "+", "%20")
	__obf_f64fa67723b7e73c = strings.ReplaceAll(__obf_f64fa67723b7e73c, "*", "%2A")
	__obf_f64fa67723b7e73c = strings.ReplaceAll(__obf_f64fa67723b7e73c, "%7E", "~")
	return __obf_f64fa67723b7e73c
}
