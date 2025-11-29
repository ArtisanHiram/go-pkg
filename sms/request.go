package __obf_18a392c58f8f4352

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

func (__obf_fdb43822b535c9dd *Request) ComposeUrl(__obf_4966c7259579709a string, __obf_723833b2544b27f1 string) string {
	__obf_8c57794a4c0b9350 := url.Values{}
	__obf_8c57794a4c0b9350.
		Add("AccessKeyId", __obf_fdb43822b535c9dd.AccessKeyId)
	__obf_8c57794a4c0b9350.
		Add("Action", __obf_fdb43822b535c9dd.Action)
	__obf_8c57794a4c0b9350.
		Add("Format", __obf_fdb43822b535c9dd.Format)
	__obf_8c57794a4c0b9350.
		Add("PhoneNumbers", __obf_fdb43822b535c9dd.PhoneNumbers)
	__obf_8c57794a4c0b9350.
		Add("RegionId", __obf_fdb43822b535c9dd.RegionId)
	__obf_8c57794a4c0b9350.
		Add("SignName", __obf_fdb43822b535c9dd.SignName)
	__obf_8c57794a4c0b9350.
		Add("SignatureMethod", __obf_fdb43822b535c9dd.SignatureMethod)
	__obf_8c57794a4c0b9350.
		Add("SignatureNonce", __obf_fdb43822b535c9dd.SignatureNonce)
	__obf_8c57794a4c0b9350.
		Add("SignatureVersion", __obf_fdb43822b535c9dd.SignatureVersion)
	__obf_8c57794a4c0b9350.
		Add("TemplateCode", __obf_fdb43822b535c9dd.TemplateCode)
	__obf_8c57794a4c0b9350.
		Add("TemplateParam", __obf_fdb43822b535c9dd.TemplateParam)
	__obf_8c57794a4c0b9350.
		Add("Timestamp", __obf_fdb43822b535c9dd.Timestamp)
	__obf_8c57794a4c0b9350.
		Add("Version", __obf_fdb43822b535c9dd.Version)
	__obf_b12e879e3a35e51f := __obf_6121932bb4480036(__obf_8c57794a4c0b9350)
	Signature := __obf_786b19c67f51c958(__obf_1212d7c248dd52c5(__obf_b12e879e3a35e51f, __obf_723833b2544b27f1))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_b12e879e3a35e51f

	return _url
}

func __obf_6121932bb4480036(__obf_8c57794a4c0b9350 url.Values) string {
	var __obf_08ab4bee36857eb4 bytes.Buffer
	__obf_f0d9b158ee944df0 := make([]string, 0, len(__obf_8c57794a4c0b9350))
	for __obf_3cec27c9aab50c6a := range __obf_8c57794a4c0b9350 {
		__obf_f0d9b158ee944df0 = append(__obf_f0d9b158ee944df0, __obf_3cec27c9aab50c6a)
	}
	sort.Strings(__obf_f0d9b158ee944df0)
	for _, __obf_3cec27c9aab50c6a := range __obf_f0d9b158ee944df0 {
		if __obf_08ab4bee36857eb4.Len() > 0 {
			__obf_08ab4bee36857eb4.
				WriteString("&")
		}
		__obf_08ab4bee36857eb4.
			WriteString(__obf_786b19c67f51c958(__obf_3cec27c9aab50c6a))
		__obf_08ab4bee36857eb4.
			WriteString("=")
		__obf_08ab4bee36857eb4.
			WriteString(__obf_786b19c67f51c958(__obf_8c57794a4c0b9350.Get(__obf_3cec27c9aab50c6a)))
	}
	return __obf_08ab4bee36857eb4.String()
}

func __obf_786b19c67f51c958(__obf_9f3ee8f5d6ba3984 string) string {
	__obf_9f3ee8f5d6ba3984 = url.QueryEscape(__obf_9f3ee8f5d6ba3984)
	__obf_9f3ee8f5d6ba3984 = strings.ReplaceAll(__obf_9f3ee8f5d6ba3984, "+", "%20")
	__obf_9f3ee8f5d6ba3984 = strings.ReplaceAll(__obf_9f3ee8f5d6ba3984, "*", "%2A")
	__obf_9f3ee8f5d6ba3984 = strings.ReplaceAll(__obf_9f3ee8f5d6ba3984, "%7E", "~")
	return __obf_9f3ee8f5d6ba3984
}
