package __obf_2457be4436ec6bd6

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

func (__obf_bc87e00d0f971da9 *Request) ComposeUrl(__obf_99cc18ae3cad05ab string, __obf_84b05fd3511d0033 string) string {
	__obf_1b1a42ee5d06be4f := url.Values{}
	__obf_1b1a42ee5d06be4f.Add("AccessKeyId", __obf_bc87e00d0f971da9.AccessKeyId)
	__obf_1b1a42ee5d06be4f.Add("Action", __obf_bc87e00d0f971da9.Action)
	__obf_1b1a42ee5d06be4f.Add("Format", __obf_bc87e00d0f971da9.Format)
	__obf_1b1a42ee5d06be4f.Add("PhoneNumbers", __obf_bc87e00d0f971da9.PhoneNumbers)
	__obf_1b1a42ee5d06be4f.Add("RegionId", __obf_bc87e00d0f971da9.RegionId)
	__obf_1b1a42ee5d06be4f.Add("SignName", __obf_bc87e00d0f971da9.SignName)
	__obf_1b1a42ee5d06be4f.Add("SignatureMethod", __obf_bc87e00d0f971da9.SignatureMethod)
	__obf_1b1a42ee5d06be4f.Add("SignatureNonce", __obf_bc87e00d0f971da9.SignatureNonce)
	__obf_1b1a42ee5d06be4f.Add("SignatureVersion", __obf_bc87e00d0f971da9.SignatureVersion)
	__obf_1b1a42ee5d06be4f.Add("TemplateCode", __obf_bc87e00d0f971da9.TemplateCode)
	__obf_1b1a42ee5d06be4f.Add("TemplateParam", __obf_bc87e00d0f971da9.TemplateParam)
	__obf_1b1a42ee5d06be4f.Add("Timestamp", __obf_bc87e00d0f971da9.Timestamp)
	__obf_1b1a42ee5d06be4f.Add("Version", __obf_bc87e00d0f971da9.Version)
	__obf_381c728d8bda2181 := __obf_4b8f7bed5d437541(__obf_1b1a42ee5d06be4f)
	Signature := __obf_ccb162ee6cf7f68b(__obf_08f1451672c192b9(__obf_381c728d8bda2181, __obf_84b05fd3511d0033))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_381c728d8bda2181

	return _url
}

func __obf_4b8f7bed5d437541(__obf_1b1a42ee5d06be4f url.Values) string {
	var __obf_e6b758df5fde9a75 bytes.Buffer
	__obf_878ff4833d03155c := make([]string, 0, len(__obf_1b1a42ee5d06be4f))
	for __obf_d329e32dea151b4c := range __obf_1b1a42ee5d06be4f {
		__obf_878ff4833d03155c = append(__obf_878ff4833d03155c, __obf_d329e32dea151b4c)
	}
	sort.Strings(__obf_878ff4833d03155c)
	for _, __obf_d329e32dea151b4c := range __obf_878ff4833d03155c {
		if __obf_e6b758df5fde9a75.Len() > 0 {
			__obf_e6b758df5fde9a75.WriteString("&")
		}
		__obf_e6b758df5fde9a75.WriteString(__obf_ccb162ee6cf7f68b(__obf_d329e32dea151b4c))
		__obf_e6b758df5fde9a75.WriteString("=")
		__obf_e6b758df5fde9a75.WriteString(__obf_ccb162ee6cf7f68b(__obf_1b1a42ee5d06be4f.Get(__obf_d329e32dea151b4c)))
	}
	return __obf_e6b758df5fde9a75.String()
}

func __obf_ccb162ee6cf7f68b(__obf_75dd9dc6bd8f2f19 string) string {
	__obf_75dd9dc6bd8f2f19 = url.QueryEscape(__obf_75dd9dc6bd8f2f19)
	__obf_75dd9dc6bd8f2f19 = strings.ReplaceAll(__obf_75dd9dc6bd8f2f19, "+", "%20")
	__obf_75dd9dc6bd8f2f19 = strings.ReplaceAll(__obf_75dd9dc6bd8f2f19, "*", "%2A")
	__obf_75dd9dc6bd8f2f19 = strings.ReplaceAll(__obf_75dd9dc6bd8f2f19, "%7E", "~")
	return __obf_75dd9dc6bd8f2f19
}
