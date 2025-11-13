package __obf_721a4aff228e6809

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

func (__obf_9df80c204293a127 *Request) ComposeUrl(__obf_2cb8e9a41b848248 string, __obf_3481be3dcb25fb5d string) string {
	__obf_29ad5f87103cd17c := url.Values{}
	__obf_29ad5f87103cd17c.Add("AccessKeyId", __obf_9df80c204293a127.AccessKeyId)
	__obf_29ad5f87103cd17c.Add("Action", __obf_9df80c204293a127.Action)
	__obf_29ad5f87103cd17c.Add("Format", __obf_9df80c204293a127.Format)
	__obf_29ad5f87103cd17c.Add("PhoneNumbers", __obf_9df80c204293a127.PhoneNumbers)
	__obf_29ad5f87103cd17c.Add("RegionId", __obf_9df80c204293a127.RegionId)
	__obf_29ad5f87103cd17c.Add("SignName", __obf_9df80c204293a127.SignName)
	__obf_29ad5f87103cd17c.Add("SignatureMethod", __obf_9df80c204293a127.SignatureMethod)
	__obf_29ad5f87103cd17c.Add("SignatureNonce", __obf_9df80c204293a127.SignatureNonce)
	__obf_29ad5f87103cd17c.Add("SignatureVersion", __obf_9df80c204293a127.SignatureVersion)
	__obf_29ad5f87103cd17c.Add("TemplateCode", __obf_9df80c204293a127.TemplateCode)
	__obf_29ad5f87103cd17c.Add("TemplateParam", __obf_9df80c204293a127.TemplateParam)
	__obf_29ad5f87103cd17c.Add("Timestamp", __obf_9df80c204293a127.Timestamp)
	__obf_29ad5f87103cd17c.Add("Version", __obf_9df80c204293a127.Version)
	__obf_445dc85641e3552e := __obf_14d9c902120ed57d(__obf_29ad5f87103cd17c)
	Signature := __obf_ec014832ed2718dc(__obf_7312697087077e2b(__obf_445dc85641e3552e, __obf_3481be3dcb25fb5d))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_445dc85641e3552e

	return _url
}

func __obf_14d9c902120ed57d(__obf_29ad5f87103cd17c url.Values) string {
	var __obf_a4e1cb5211d93fe3 bytes.Buffer
	__obf_496f8b8c5e614a25 := make([]string, 0, len(__obf_29ad5f87103cd17c))
	for __obf_bc3b7cf0498b0789 := range __obf_29ad5f87103cd17c {
		__obf_496f8b8c5e614a25 = append(__obf_496f8b8c5e614a25, __obf_bc3b7cf0498b0789)
	}
	sort.Strings(__obf_496f8b8c5e614a25)
	for _, __obf_bc3b7cf0498b0789 := range __obf_496f8b8c5e614a25 {
		if __obf_a4e1cb5211d93fe3.Len() > 0 {
			__obf_a4e1cb5211d93fe3.WriteString("&")
		}
		__obf_a4e1cb5211d93fe3.WriteString(__obf_ec014832ed2718dc(__obf_bc3b7cf0498b0789))
		__obf_a4e1cb5211d93fe3.WriteString("=")
		__obf_a4e1cb5211d93fe3.WriteString(__obf_ec014832ed2718dc(__obf_29ad5f87103cd17c.Get(__obf_bc3b7cf0498b0789)))
	}
	return __obf_a4e1cb5211d93fe3.String()
}

func __obf_ec014832ed2718dc(__obf_bbc97f6f2cd5931a string) string {
	__obf_bbc97f6f2cd5931a = url.QueryEscape(__obf_bbc97f6f2cd5931a)
	__obf_bbc97f6f2cd5931a = strings.ReplaceAll(__obf_bbc97f6f2cd5931a, "+", "%20")
	__obf_bbc97f6f2cd5931a = strings.ReplaceAll(__obf_bbc97f6f2cd5931a, "*", "%2A")
	__obf_bbc97f6f2cd5931a = strings.ReplaceAll(__obf_bbc97f6f2cd5931a, "%7E", "~")
	return __obf_bbc97f6f2cd5931a
}
