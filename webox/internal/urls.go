package __obf_dfacbdcb930f08e7

// 小程序相关URL
const (
	MiniprogramCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session"
	// 更多小程序URL可以在这里添加
	MiniprogramGetPhoneNumberURL       = "https://api.weixin.qq.com/wxa/business/getuserphonenumber" // 获取用户手机号新接口
	MiniprogramGetAccessTokenURL       = "https://api.weixin.qq.com/cgi-bin/token"                   // 获取 AccessToken 接口
	MiniprogramGetStableAccessTokenURL = "https://api.weixin.qq.com/cgi-bin/stable_token"
)

// 支付相关URL
const (
	PaymentUnifiedOrderURL = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	// 更多支付URL可以在这里添加，例如：退款、查询订单等
	PaymentRefundURL = "https://api.mch.weixin.qq.com/secapi/pay/refund"     // 退款接口需要证书
	PaymentTransURL  = "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank" // 付款到银行卡接口需要证书

)
