package __obf_dd37d42fbda9c938

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

func (__obf_3b271c014009f4ce *Request) ComposeUrl(__obf_dea4ca8e2bc3b5e8 string, __obf_9796f2cbd157ab60 string) string {
	__obf_36c1048e83ed061d := url.Values{}
	__obf_36c1048e83ed061d.
		Add("AccessKeyId", __obf_3b271c014009f4ce.AccessKeyId)
	__obf_36c1048e83ed061d.
		Add("Action", __obf_3b271c014009f4ce.Action)
	__obf_36c1048e83ed061d.
		Add("Format", __obf_3b271c014009f4ce.Format)
	__obf_36c1048e83ed061d.
		Add("PhoneNumbers", __obf_3b271c014009f4ce.PhoneNumbers)
	__obf_36c1048e83ed061d.
		Add("RegionId", __obf_3b271c014009f4ce.RegionId)
	__obf_36c1048e83ed061d.
		Add("SignName", __obf_3b271c014009f4ce.SignName)
	__obf_36c1048e83ed061d.
		Add("SignatureMethod", __obf_3b271c014009f4ce.SignatureMethod)
	__obf_36c1048e83ed061d.
		Add("SignatureNonce", __obf_3b271c014009f4ce.SignatureNonce)
	__obf_36c1048e83ed061d.
		Add("SignatureVersion", __obf_3b271c014009f4ce.SignatureVersion)
	__obf_36c1048e83ed061d.
		Add("TemplateCode", __obf_3b271c014009f4ce.TemplateCode)
	__obf_36c1048e83ed061d.
		Add("TemplateParam", __obf_3b271c014009f4ce.TemplateParam)
	__obf_36c1048e83ed061d.
		Add("Timestamp", __obf_3b271c014009f4ce.Timestamp)
	__obf_36c1048e83ed061d.
		Add("Version", __obf_3b271c014009f4ce.Version)
	__obf_92a1890c51da5c0a := __obf_3d4a1793e73d0a22(__obf_36c1048e83ed061d)
	Signature := __obf_be8d53df8820d385(__obf_3072ced250ac589d(__obf_92a1890c51da5c0a, __obf_9796f2cbd157ab60))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_92a1890c51da5c0a

	return _url
}

func __obf_3d4a1793e73d0a22(__obf_36c1048e83ed061d url.Values) string {
	var __obf_3ad276f752c7af0b bytes.Buffer
	__obf_c575bf2f27f522c0 := make([]string, 0, len(__obf_36c1048e83ed061d))
	for __obf_b60a5975c1af7e63 := range __obf_36c1048e83ed061d {
		__obf_c575bf2f27f522c0 = append(__obf_c575bf2f27f522c0, __obf_b60a5975c1af7e63)
	}
	sort.Strings(__obf_c575bf2f27f522c0)
	for _, __obf_b60a5975c1af7e63 := range __obf_c575bf2f27f522c0 {
		if __obf_3ad276f752c7af0b.Len() > 0 {
			__obf_3ad276f752c7af0b.
				WriteString("&")
		}
		__obf_3ad276f752c7af0b.
			WriteString(__obf_be8d53df8820d385(__obf_b60a5975c1af7e63))
		__obf_3ad276f752c7af0b.
			WriteString("=")
		__obf_3ad276f752c7af0b.
			WriteString(__obf_be8d53df8820d385(__obf_36c1048e83ed061d.Get(__obf_b60a5975c1af7e63)))
	}
	return __obf_3ad276f752c7af0b.String()
}

func __obf_be8d53df8820d385(__obf_b23fd47ea99e8acb string) string {
	__obf_b23fd47ea99e8acb = url.QueryEscape(__obf_b23fd47ea99e8acb)
	__obf_b23fd47ea99e8acb = strings.ReplaceAll(__obf_b23fd47ea99e8acb, "+", "%20")
	__obf_b23fd47ea99e8acb = strings.ReplaceAll(__obf_b23fd47ea99e8acb, "*", "%2A")
	__obf_b23fd47ea99e8acb = strings.ReplaceAll(__obf_b23fd47ea99e8acb, "%7E", "~")
	return __obf_b23fd47ea99e8acb
}
