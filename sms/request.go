package __obf_cb62198a5f8c0e2c

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

func (__obf_2a85caaeffd32037 *Request) ComposeUrl(__obf_e13700a67eb28cbb string, __obf_af4b82e4375e005a string) string {
	__obf_26d21a557c57186e := url.Values{}
	__obf_26d21a557c57186e.Add("AccessKeyId", __obf_2a85caaeffd32037.AccessKeyId)
	__obf_26d21a557c57186e.Add("Action", __obf_2a85caaeffd32037.Action)
	__obf_26d21a557c57186e.Add("Format", __obf_2a85caaeffd32037.Format)
	__obf_26d21a557c57186e.Add("PhoneNumbers", __obf_2a85caaeffd32037.PhoneNumbers)
	__obf_26d21a557c57186e.Add("RegionId", __obf_2a85caaeffd32037.RegionId)
	__obf_26d21a557c57186e.Add("SignName", __obf_2a85caaeffd32037.SignName)
	__obf_26d21a557c57186e.Add("SignatureMethod", __obf_2a85caaeffd32037.SignatureMethod)
	__obf_26d21a557c57186e.Add("SignatureNonce", __obf_2a85caaeffd32037.SignatureNonce)
	__obf_26d21a557c57186e.Add("SignatureVersion", __obf_2a85caaeffd32037.SignatureVersion)
	__obf_26d21a557c57186e.Add("TemplateCode", __obf_2a85caaeffd32037.TemplateCode)
	__obf_26d21a557c57186e.Add("TemplateParam", __obf_2a85caaeffd32037.TemplateParam)
	__obf_26d21a557c57186e.Add("Timestamp", __obf_2a85caaeffd32037.Timestamp)
	__obf_26d21a557c57186e.Add("Version", __obf_2a85caaeffd32037.Version)
	__obf_3b756f6de76ea66f := __obf_2346704bb5e01c49(__obf_26d21a557c57186e)
	Signature := __obf_fdce8fc899639dcd(__obf_e993b87f82f2044b(__obf_3b756f6de76ea66f, __obf_af4b82e4375e005a))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_3b756f6de76ea66f

	return _url
}

func __obf_2346704bb5e01c49(__obf_26d21a557c57186e url.Values) string {
	var __obf_b80aea217554822a bytes.Buffer
	__obf_3a51df88210163a4 := make([]string, 0, len(__obf_26d21a557c57186e))
	for __obf_359296e4523d5970 := range __obf_26d21a557c57186e {
		__obf_3a51df88210163a4 = append(__obf_3a51df88210163a4, __obf_359296e4523d5970)
	}
	sort.Strings(__obf_3a51df88210163a4)
	for _, __obf_359296e4523d5970 := range __obf_3a51df88210163a4 {
		if __obf_b80aea217554822a.Len() > 0 {
			__obf_b80aea217554822a.WriteString("&")
		}
		__obf_b80aea217554822a.WriteString(__obf_fdce8fc899639dcd(__obf_359296e4523d5970))
		__obf_b80aea217554822a.WriteString("=")
		__obf_b80aea217554822a.WriteString(__obf_fdce8fc899639dcd(__obf_26d21a557c57186e.Get(__obf_359296e4523d5970)))
	}
	return __obf_b80aea217554822a.String()
}

func __obf_fdce8fc899639dcd(__obf_234ca448b6fcfb45 string) string {
	__obf_234ca448b6fcfb45 = url.QueryEscape(__obf_234ca448b6fcfb45)
	__obf_234ca448b6fcfb45 = strings.ReplaceAll(__obf_234ca448b6fcfb45, "+", "%20")
	__obf_234ca448b6fcfb45 = strings.ReplaceAll(__obf_234ca448b6fcfb45, "*", "%2A")
	__obf_234ca448b6fcfb45 = strings.ReplaceAll(__obf_234ca448b6fcfb45, "%7E", "~")
	return __obf_234ca448b6fcfb45
}
