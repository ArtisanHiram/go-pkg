package __obf_52bdfa18dc226ac6

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

func (__obf_b06715312c73ca7a *Request) ComposeUrl(__obf_1648f8242ea2cf90 string, __obf_2cda83fc4f45e711 string) string {
	__obf_0e04e3b86df07260 := url.Values{}
	__obf_0e04e3b86df07260.Add("AccessKeyId", __obf_b06715312c73ca7a.AccessKeyId)
	__obf_0e04e3b86df07260.Add("Action", __obf_b06715312c73ca7a.Action)
	__obf_0e04e3b86df07260.Add("Format", __obf_b06715312c73ca7a.Format)
	__obf_0e04e3b86df07260.Add("PhoneNumbers", __obf_b06715312c73ca7a.PhoneNumbers)
	__obf_0e04e3b86df07260.Add("RegionId", __obf_b06715312c73ca7a.RegionId)
	__obf_0e04e3b86df07260.Add("SignName", __obf_b06715312c73ca7a.SignName)
	__obf_0e04e3b86df07260.Add("SignatureMethod", __obf_b06715312c73ca7a.SignatureMethod)
	__obf_0e04e3b86df07260.Add("SignatureNonce", __obf_b06715312c73ca7a.SignatureNonce)
	__obf_0e04e3b86df07260.Add("SignatureVersion", __obf_b06715312c73ca7a.SignatureVersion)
	__obf_0e04e3b86df07260.Add("TemplateCode", __obf_b06715312c73ca7a.TemplateCode)
	__obf_0e04e3b86df07260.Add("TemplateParam", __obf_b06715312c73ca7a.TemplateParam)
	__obf_0e04e3b86df07260.Add("Timestamp", __obf_b06715312c73ca7a.Timestamp)
	__obf_0e04e3b86df07260.Add("Version", __obf_b06715312c73ca7a.Version)
	__obf_7b3d83ef9ab2c8f2 := __obf_4d1db89a54b812c0(__obf_0e04e3b86df07260)
	Signature := __obf_c6816d7254c31423(__obf_7c84f133c92d8b15(__obf_7b3d83ef9ab2c8f2, __obf_2cda83fc4f45e711))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_7b3d83ef9ab2c8f2

	return _url
}

func __obf_4d1db89a54b812c0(__obf_0e04e3b86df07260 url.Values) string {
	var __obf_fc411cca941abe58 bytes.Buffer
	__obf_c9d46c43a36bcf14 := make([]string, 0, len(__obf_0e04e3b86df07260))
	for __obf_3b30a59e7b0110b0 := range __obf_0e04e3b86df07260 {
		__obf_c9d46c43a36bcf14 = append(__obf_c9d46c43a36bcf14, __obf_3b30a59e7b0110b0)
	}
	sort.Strings(__obf_c9d46c43a36bcf14)
	for _, __obf_3b30a59e7b0110b0 := range __obf_c9d46c43a36bcf14 {
		if __obf_fc411cca941abe58.Len() > 0 {
			__obf_fc411cca941abe58.WriteString("&")
		}
		__obf_fc411cca941abe58.WriteString(__obf_c6816d7254c31423(__obf_3b30a59e7b0110b0))
		__obf_fc411cca941abe58.WriteString("=")
		__obf_fc411cca941abe58.WriteString(__obf_c6816d7254c31423(__obf_0e04e3b86df07260.Get(__obf_3b30a59e7b0110b0)))
	}
	return __obf_fc411cca941abe58.String()
}

func __obf_c6816d7254c31423(__obf_4a5450a7cf26ea8d string) string {
	__obf_4a5450a7cf26ea8d = url.QueryEscape(__obf_4a5450a7cf26ea8d)
	__obf_4a5450a7cf26ea8d = strings.ReplaceAll(__obf_4a5450a7cf26ea8d, "+", "%20")
	__obf_4a5450a7cf26ea8d = strings.ReplaceAll(__obf_4a5450a7cf26ea8d, "*", "%2A")
	__obf_4a5450a7cf26ea8d = strings.ReplaceAll(__obf_4a5450a7cf26ea8d, "%7E", "~")
	return __obf_4a5450a7cf26ea8d
}
