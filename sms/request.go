package __obf_7147dbcba87e70eb

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

func (__obf_11951d3b3f0af32a *Request) ComposeUrl(__obf_86f377abe3973f7f string, __obf_2d230803be94b4a9 string) string {
	__obf_8d1a2986736459e9 := url.Values{}
	__obf_8d1a2986736459e9.
		Add("AccessKeyId", __obf_11951d3b3f0af32a.AccessKeyId)
	__obf_8d1a2986736459e9.
		Add("Action", __obf_11951d3b3f0af32a.Action)
	__obf_8d1a2986736459e9.
		Add("Format", __obf_11951d3b3f0af32a.Format)
	__obf_8d1a2986736459e9.
		Add("PhoneNumbers", __obf_11951d3b3f0af32a.PhoneNumbers)
	__obf_8d1a2986736459e9.
		Add("RegionId", __obf_11951d3b3f0af32a.RegionId)
	__obf_8d1a2986736459e9.
		Add("SignName", __obf_11951d3b3f0af32a.SignName)
	__obf_8d1a2986736459e9.
		Add("SignatureMethod", __obf_11951d3b3f0af32a.SignatureMethod)
	__obf_8d1a2986736459e9.
		Add("SignatureNonce", __obf_11951d3b3f0af32a.SignatureNonce)
	__obf_8d1a2986736459e9.
		Add("SignatureVersion", __obf_11951d3b3f0af32a.SignatureVersion)
	__obf_8d1a2986736459e9.
		Add("TemplateCode", __obf_11951d3b3f0af32a.TemplateCode)
	__obf_8d1a2986736459e9.
		Add("TemplateParam", __obf_11951d3b3f0af32a.TemplateParam)
	__obf_8d1a2986736459e9.
		Add("Timestamp", __obf_11951d3b3f0af32a.Timestamp)
	__obf_8d1a2986736459e9.
		Add("Version", __obf_11951d3b3f0af32a.Version)
	__obf_7b4c32517341fd2d := __obf_70535f4bc6e4828c(__obf_8d1a2986736459e9)
	Signature := __obf_b3ef92f29440b2e9(__obf_63582ab28b26163a(__obf_7b4c32517341fd2d, __obf_2d230803be94b4a9))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_7b4c32517341fd2d

	return _url
}

func __obf_70535f4bc6e4828c(__obf_8d1a2986736459e9 url.Values) string {
	var __obf_86662e90ccc51e4a bytes.Buffer
	__obf_5ed8c170a1b8f4a4 := make([]string, 0, len(__obf_8d1a2986736459e9))
	for __obf_0c5b1f67a8a5cdde := range __obf_8d1a2986736459e9 {
		__obf_5ed8c170a1b8f4a4 = append(__obf_5ed8c170a1b8f4a4, __obf_0c5b1f67a8a5cdde)
	}
	sort.Strings(__obf_5ed8c170a1b8f4a4)
	for _, __obf_0c5b1f67a8a5cdde := range __obf_5ed8c170a1b8f4a4 {
		if __obf_86662e90ccc51e4a.Len() > 0 {
			__obf_86662e90ccc51e4a.
				WriteString("&")
		}
		__obf_86662e90ccc51e4a.
			WriteString(__obf_b3ef92f29440b2e9(__obf_0c5b1f67a8a5cdde))
		__obf_86662e90ccc51e4a.
			WriteString("=")
		__obf_86662e90ccc51e4a.
			WriteString(__obf_b3ef92f29440b2e9(__obf_8d1a2986736459e9.Get(__obf_0c5b1f67a8a5cdde)))
	}
	return __obf_86662e90ccc51e4a.String()
}

func __obf_b3ef92f29440b2e9(__obf_36f63639e4a81c02 string) string {
	__obf_36f63639e4a81c02 = url.QueryEscape(__obf_36f63639e4a81c02)
	__obf_36f63639e4a81c02 = strings.ReplaceAll(__obf_36f63639e4a81c02, "+", "%20")
	__obf_36f63639e4a81c02 = strings.ReplaceAll(__obf_36f63639e4a81c02, "*", "%2A")
	__obf_36f63639e4a81c02 = strings.ReplaceAll(__obf_36f63639e4a81c02, "%7E", "~")
	return __obf_36f63639e4a81c02
}
