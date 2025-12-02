package __obf_0f0912cb4961947f

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

func (__obf_3671ea45d3ea0053 *Request) ComposeUrl(__obf_5b71f846896019ba string, __obf_dd94e101d195a7c7 string) string {
	__obf_d4d4d122eddfd9ea := url.Values{}
	__obf_d4d4d122eddfd9ea.
		Add("AccessKeyId", __obf_3671ea45d3ea0053.AccessKeyId)
	__obf_d4d4d122eddfd9ea.
		Add("Action", __obf_3671ea45d3ea0053.Action)
	__obf_d4d4d122eddfd9ea.
		Add("Format", __obf_3671ea45d3ea0053.Format)
	__obf_d4d4d122eddfd9ea.
		Add("PhoneNumbers", __obf_3671ea45d3ea0053.PhoneNumbers)
	__obf_d4d4d122eddfd9ea.
		Add("RegionId", __obf_3671ea45d3ea0053.RegionId)
	__obf_d4d4d122eddfd9ea.
		Add("SignName", __obf_3671ea45d3ea0053.SignName)
	__obf_d4d4d122eddfd9ea.
		Add("SignatureMethod", __obf_3671ea45d3ea0053.SignatureMethod)
	__obf_d4d4d122eddfd9ea.
		Add("SignatureNonce", __obf_3671ea45d3ea0053.SignatureNonce)
	__obf_d4d4d122eddfd9ea.
		Add("SignatureVersion", __obf_3671ea45d3ea0053.SignatureVersion)
	__obf_d4d4d122eddfd9ea.
		Add("TemplateCode", __obf_3671ea45d3ea0053.TemplateCode)
	__obf_d4d4d122eddfd9ea.
		Add("TemplateParam", __obf_3671ea45d3ea0053.TemplateParam)
	__obf_d4d4d122eddfd9ea.
		Add("Timestamp", __obf_3671ea45d3ea0053.Timestamp)
	__obf_d4d4d122eddfd9ea.
		Add("Version", __obf_3671ea45d3ea0053.Version)
	__obf_bb95acccf25e05f1 := __obf_c2b27722a4309195(__obf_d4d4d122eddfd9ea)
	Signature := __obf_4e6893075e8516f5(__obf_911e53cade1302f7(__obf_bb95acccf25e05f1, __obf_dd94e101d195a7c7))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_bb95acccf25e05f1

	return _url
}

func __obf_c2b27722a4309195(__obf_d4d4d122eddfd9ea url.Values) string {
	var __obf_8df63b5ae6bde37d bytes.Buffer
	__obf_aed34b90dffe778b := make([]string, 0, len(__obf_d4d4d122eddfd9ea))
	for __obf_579aa8570e0ceba0 := range __obf_d4d4d122eddfd9ea {
		__obf_aed34b90dffe778b = append(__obf_aed34b90dffe778b, __obf_579aa8570e0ceba0)
	}
	sort.Strings(__obf_aed34b90dffe778b)
	for _, __obf_579aa8570e0ceba0 := range __obf_aed34b90dffe778b {
		if __obf_8df63b5ae6bde37d.Len() > 0 {
			__obf_8df63b5ae6bde37d.
				WriteString("&")
		}
		__obf_8df63b5ae6bde37d.
			WriteString(__obf_4e6893075e8516f5(__obf_579aa8570e0ceba0))
		__obf_8df63b5ae6bde37d.
			WriteString("=")
		__obf_8df63b5ae6bde37d.
			WriteString(__obf_4e6893075e8516f5(__obf_d4d4d122eddfd9ea.Get(__obf_579aa8570e0ceba0)))
	}
	return __obf_8df63b5ae6bde37d.String()
}

func __obf_4e6893075e8516f5(__obf_fc75162f1b2ea3b8 string) string {
	__obf_fc75162f1b2ea3b8 = url.QueryEscape(__obf_fc75162f1b2ea3b8)
	__obf_fc75162f1b2ea3b8 = strings.ReplaceAll(__obf_fc75162f1b2ea3b8, "+", "%20")
	__obf_fc75162f1b2ea3b8 = strings.ReplaceAll(__obf_fc75162f1b2ea3b8, "*", "%2A")
	__obf_fc75162f1b2ea3b8 = strings.ReplaceAll(__obf_fc75162f1b2ea3b8, "%7E", "~")
	return __obf_fc75162f1b2ea3b8
}
