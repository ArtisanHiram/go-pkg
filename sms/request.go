package __obf_a307862f84d54cc6

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

func (__obf_4c76c53a33a5db4d *Request) ComposeUrl(__obf_7c0d63a339764d75 string, __obf_0d83fb6d65187b1d string) string {
	__obf_db1bbfced8547939 := url.Values{}
	__obf_db1bbfced8547939.Add("AccessKeyId", __obf_4c76c53a33a5db4d.AccessKeyId)
	__obf_db1bbfced8547939.Add("Action", __obf_4c76c53a33a5db4d.Action)
	__obf_db1bbfced8547939.Add("Format", __obf_4c76c53a33a5db4d.Format)
	__obf_db1bbfced8547939.Add("PhoneNumbers", __obf_4c76c53a33a5db4d.PhoneNumbers)
	__obf_db1bbfced8547939.Add("RegionId", __obf_4c76c53a33a5db4d.RegionId)
	__obf_db1bbfced8547939.Add("SignName", __obf_4c76c53a33a5db4d.SignName)
	__obf_db1bbfced8547939.Add("SignatureMethod", __obf_4c76c53a33a5db4d.SignatureMethod)
	__obf_db1bbfced8547939.Add("SignatureNonce", __obf_4c76c53a33a5db4d.SignatureNonce)
	__obf_db1bbfced8547939.Add("SignatureVersion", __obf_4c76c53a33a5db4d.SignatureVersion)
	__obf_db1bbfced8547939.Add("TemplateCode", __obf_4c76c53a33a5db4d.TemplateCode)
	__obf_db1bbfced8547939.Add("TemplateParam", __obf_4c76c53a33a5db4d.TemplateParam)
	__obf_db1bbfced8547939.Add("Timestamp", __obf_4c76c53a33a5db4d.Timestamp)
	__obf_db1bbfced8547939.Add("Version", __obf_4c76c53a33a5db4d.Version)
	__obf_b6eb9d8a80bc2cb1 := __obf_47e75f2fdeec63af(__obf_db1bbfced8547939)
	Signature := __obf_03c24bcdceda96ba(__obf_cdacf59f3a4d15f3(__obf_b6eb9d8a80bc2cb1, __obf_0d83fb6d65187b1d))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_b6eb9d8a80bc2cb1

	return _url
}

func __obf_47e75f2fdeec63af(__obf_db1bbfced8547939 url.Values) string {
	var __obf_ed55315e422a93ba bytes.Buffer
	__obf_265fba0806a78870 := make([]string, 0, len(__obf_db1bbfced8547939))
	for __obf_35f1be3a8abc964a := range __obf_db1bbfced8547939 {
		__obf_265fba0806a78870 = append(__obf_265fba0806a78870, __obf_35f1be3a8abc964a)
	}
	sort.Strings(__obf_265fba0806a78870)
	for _, __obf_35f1be3a8abc964a := range __obf_265fba0806a78870 {
		if __obf_ed55315e422a93ba.Len() > 0 {
			__obf_ed55315e422a93ba.WriteString("&")
		}
		__obf_ed55315e422a93ba.WriteString(__obf_03c24bcdceda96ba(__obf_35f1be3a8abc964a))
		__obf_ed55315e422a93ba.WriteString("=")
		__obf_ed55315e422a93ba.WriteString(__obf_03c24bcdceda96ba(__obf_db1bbfced8547939.Get(__obf_35f1be3a8abc964a)))
	}
	return __obf_ed55315e422a93ba.String()
}

func __obf_03c24bcdceda96ba(__obf_fb1d74ba34ce4b7e string) string {
	__obf_fb1d74ba34ce4b7e = url.QueryEscape(__obf_fb1d74ba34ce4b7e)
	__obf_fb1d74ba34ce4b7e = strings.ReplaceAll(__obf_fb1d74ba34ce4b7e, "+", "%20")
	__obf_fb1d74ba34ce4b7e = strings.ReplaceAll(__obf_fb1d74ba34ce4b7e, "*", "%2A")
	__obf_fb1d74ba34ce4b7e = strings.ReplaceAll(__obf_fb1d74ba34ce4b7e, "%7E", "~")
	return __obf_fb1d74ba34ce4b7e
}
