package __obf_28053862c3d647b8

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

func (__obf_921feb7315ea4ee0 *Request) ComposeUrl(__obf_2ec4c2cd43047f9d string, __obf_ead40cbd1507b913 string) string {
	__obf_cd28606aa87da3c9 := url.Values{}
	__obf_cd28606aa87da3c9.Add("AccessKeyId", __obf_921feb7315ea4ee0.AccessKeyId)
	__obf_cd28606aa87da3c9.Add("Action", __obf_921feb7315ea4ee0.Action)
	__obf_cd28606aa87da3c9.Add("Format", __obf_921feb7315ea4ee0.Format)
	__obf_cd28606aa87da3c9.Add("PhoneNumbers", __obf_921feb7315ea4ee0.PhoneNumbers)
	__obf_cd28606aa87da3c9.Add("RegionId", __obf_921feb7315ea4ee0.RegionId)
	__obf_cd28606aa87da3c9.Add("SignName", __obf_921feb7315ea4ee0.SignName)
	__obf_cd28606aa87da3c9.Add("SignatureMethod", __obf_921feb7315ea4ee0.SignatureMethod)
	__obf_cd28606aa87da3c9.Add("SignatureNonce", __obf_921feb7315ea4ee0.SignatureNonce)
	__obf_cd28606aa87da3c9.Add("SignatureVersion", __obf_921feb7315ea4ee0.SignatureVersion)
	__obf_cd28606aa87da3c9.Add("TemplateCode", __obf_921feb7315ea4ee0.TemplateCode)
	__obf_cd28606aa87da3c9.Add("TemplateParam", __obf_921feb7315ea4ee0.TemplateParam)
	__obf_cd28606aa87da3c9.Add("Timestamp", __obf_921feb7315ea4ee0.Timestamp)
	__obf_cd28606aa87da3c9.Add("Version", __obf_921feb7315ea4ee0.Version)
	__obf_39d1a0f71d2ffb91 := __obf_6bda852866144048(__obf_cd28606aa87da3c9)
	Signature := __obf_05b630c8ad786500(__obf_ae740a1c7efd2892(__obf_39d1a0f71d2ffb91, __obf_ead40cbd1507b913))

	_url := "http://dysmsapi.aliyuncs.com/?Signature=" + Signature + "&" + __obf_39d1a0f71d2ffb91

	return _url
}

func __obf_6bda852866144048(__obf_cd28606aa87da3c9 url.Values) string {
	var __obf_ad0e891925307581 bytes.Buffer
	__obf_64b4e8617885c055 := make([]string, 0, len(__obf_cd28606aa87da3c9))
	for __obf_9c510381bc5db23b := range __obf_cd28606aa87da3c9 {
		__obf_64b4e8617885c055 = append(__obf_64b4e8617885c055, __obf_9c510381bc5db23b)
	}
	sort.Strings(__obf_64b4e8617885c055)
	for _, __obf_9c510381bc5db23b := range __obf_64b4e8617885c055 {
		if __obf_ad0e891925307581.Len() > 0 {
			__obf_ad0e891925307581.WriteString("&")
		}
		__obf_ad0e891925307581.WriteString(__obf_05b630c8ad786500(__obf_9c510381bc5db23b))
		__obf_ad0e891925307581.WriteString("=")
		__obf_ad0e891925307581.WriteString(__obf_05b630c8ad786500(__obf_cd28606aa87da3c9.Get(__obf_9c510381bc5db23b)))
	}
	return __obf_ad0e891925307581.String()
}

func __obf_05b630c8ad786500(__obf_8c8b461b854d2212 string) string {
	__obf_8c8b461b854d2212 = url.QueryEscape(__obf_8c8b461b854d2212)
	__obf_8c8b461b854d2212 = strings.ReplaceAll(__obf_8c8b461b854d2212, "+", "%20")
	__obf_8c8b461b854d2212 = strings.ReplaceAll(__obf_8c8b461b854d2212, "*", "%2A")
	__obf_8c8b461b854d2212 = strings.ReplaceAll(__obf_8c8b461b854d2212, "%7E", "~")
	return __obf_8c8b461b854d2212
}
