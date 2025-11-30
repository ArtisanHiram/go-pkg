package __obf_dc51e1c30a41549a

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

func (__obf_7a3fcb2f4dad69f6 *Request) ComposeUrl(__obf_680d1069bf4e7d29 string, __obf_fbec4915c4cb08dc string) string {
	__obf_6bde3ce42574faed := url.Values{}
	__obf_6bde3ce42574faed.
		Add("AccessKeyId", __obf_7a3fcb2f4dad69f6.AccessKeyId)
	__obf_6bde3ce42574faed.
		Add("Action", __obf_7a3fcb2f4dad69f6.Action)
	__obf_6bde3ce42574faed.
		Add("Format", __obf_7a3fcb2f4dad69f6.Format)
	__obf_6bde3ce42574faed.
		Add("PhoneNumbers", __obf_7a3fcb2f4dad69f6.PhoneNumbers)
	__obf_6bde3ce42574faed.
		Add("RegionId", __obf_7a3fcb2f4dad69f6.RegionId)
	__obf_6bde3ce42574faed.
		Add("SignName", __obf_7a3fcb2f4dad69f6.SignName)
	__obf_6bde3ce42574faed.
		Add("SignatureMethod", __obf_7a3fcb2f4dad69f6.SignatureMethod)
	__obf_6bde3ce42574faed.
		Add("SignatureNonce", __obf_7a3fcb2f4dad69f6.SignatureNonce)
	__obf_6bde3ce42574faed.
		Add("SignatureVersion", __obf_7a3fcb2f4dad69f6.SignatureVersion)
	__obf_6bde3ce42574faed.
		Add("TemplateCode", __obf_7a3fcb2f4dad69f6.TemplateCode)
	__obf_6bde3ce42574faed.
		Add("TemplateParam", __obf_7a3fcb2f4dad69f6.TemplateParam)
	__obf_6bde3ce42574faed.
		Add("Timestamp", __obf_7a3fcb2f4dad69f6.Timestamp)
	__obf_6bde3ce42574faed.
		Add("Version", __obf_7a3fcb2f4dad69f6.Version)
	__obf_03e8fa9024bd1b4e := __obf_ab08003662ad475f(__obf_6bde3ce42574faed)
	Signature := __obf_11b188aaefdfc57a(__obf_2a1bcfd3101707aa(__obf_03e8fa9024bd1b4e, __obf_fbec4915c4cb08dc))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_03e8fa9024bd1b4e

	return _url
}

func __obf_ab08003662ad475f(__obf_6bde3ce42574faed url.Values) string {
	var __obf_e68723f9e199d96f bytes.Buffer
	__obf_087153070cadf448 := make([]string, 0, len(__obf_6bde3ce42574faed))
	for __obf_215a316a3bc29831 := range __obf_6bde3ce42574faed {
		__obf_087153070cadf448 = append(__obf_087153070cadf448, __obf_215a316a3bc29831)
	}
	sort.Strings(__obf_087153070cadf448)
	for _, __obf_215a316a3bc29831 := range __obf_087153070cadf448 {
		if __obf_e68723f9e199d96f.Len() > 0 {
			__obf_e68723f9e199d96f.
				WriteString("&")
		}
		__obf_e68723f9e199d96f.
			WriteString(__obf_11b188aaefdfc57a(__obf_215a316a3bc29831))
		__obf_e68723f9e199d96f.
			WriteString("=")
		__obf_e68723f9e199d96f.
			WriteString(__obf_11b188aaefdfc57a(__obf_6bde3ce42574faed.Get(__obf_215a316a3bc29831)))
	}
	return __obf_e68723f9e199d96f.String()
}

func __obf_11b188aaefdfc57a(__obf_8767044ef3d23e2c string) string {
	__obf_8767044ef3d23e2c = url.QueryEscape(__obf_8767044ef3d23e2c)
	__obf_8767044ef3d23e2c = strings.ReplaceAll(__obf_8767044ef3d23e2c, "+", "%20")
	__obf_8767044ef3d23e2c = strings.ReplaceAll(__obf_8767044ef3d23e2c, "*", "%2A")
	__obf_8767044ef3d23e2c = strings.ReplaceAll(__obf_8767044ef3d23e2c, "%7E", "~")
	return __obf_8767044ef3d23e2c
}
