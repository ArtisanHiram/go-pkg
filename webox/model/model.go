package __obf_65905cbb9e2ee198

import (
	"encoding/xml"
	"fmt"
)

const (
	SUCCESS = "SUCCESS"
	FAIL    = "FAIL"
)

// 业务结果
type ErrResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// XMLRequest 支付请求的公共结构体，包含xml头部
type XMLRequest struct {
	XMLName  xml.Name `xml:"xml"`
	AppID    string   `xml:"appid"`     // 公众账号ID
	MchID    string   `xml:"mch_id"`    // 商户号
	NonceStr string   `xml:"nonce_str"` // 随机字符串
	Sign     string   `xml:"sign"`      // 签名
}

// XMLResponse 支付响应的公共结构体，包含xml头部和返回码
// 通信结果
type XMLResponse struct {
	XMLName xml.Name `xml:"xml"`
	// 通信结果
	ReturnCode string `xml:"return_code"` // 返回状态码
	ReturnMsg  string `xml:"return_msg"`  // 返回信息
	// 业务结果
	ResultCode string `xml:"result_code"`  // 业务结果
	ErrCode    string `xml:"err_code"`     // 错误代码
	ErrCodeDes string `xml:"err_code_des"` // 错误代码描述
}

// UnifiedOrderRequest 统一下单请求参数
type UnifiedOrderRequest struct {
	XMLRequest
	NotifyURL      string `xml:"notify_url"`                 // 通知地址
	Body           string `xml:"body"`                       // 商品描述
	OutTradeNo     string `xml:"out_trade_no"`               // 商户订单号
	TotalFee       int64  `xml:"total_fee"`                  // 总金额，单位分
	SpbillCreateIP string `xml:"spbill_create_ip,omitempty"` // 终端IP
	TradeType      string `xml:"trade_type"`                 // 交易类型，如 JSAPI、NATIVE、APP
	SignType       string `xml:"sign_type"`                  // 签名加密类型MD5\HMAC-SHA256
	OpenID         string `xml:"openid,omitempty"`           // 用户标识，trade_type=JSAPI 时必填
	ProductID      string `xml:"product_id,omitempty"`       // 商品ID trade_type=NATIVE时，此参数必传
	Attach         string `xml:"attach,omitempty"`           // 附加数据
	TimeExpire     string `xml:"time_expire,omitempty"`      // 订单失效时间

	// ... 其他可选参数，根据需要添加
}

func (__obf_6480c1d409e2c928 UnifiedOrderRequest) ToMap() map[string]string {
	__obf_3cc15a7a5c34e65b := map[string]string{
		"appid":            __obf_6480c1d409e2c928.AppID,
		"mch_id":           __obf_6480c1d409e2c928.MchID,
		"nonce_str":        __obf_6480c1d409e2c928.NonceStr,
		"openid":           __obf_6480c1d409e2c928.OpenID,
		"notify_url":       __obf_6480c1d409e2c928.NotifyURL,
		"body":             __obf_6480c1d409e2c928.Body,
		"out_trade_no":     __obf_6480c1d409e2c928.OutTradeNo,
		"total_fee":        fmt.Sprint(__obf_6480c1d409e2c928.TotalFee),
		"product_id":       __obf_6480c1d409e2c928.ProductID,
		"spbill_create_ip": __obf_6480c1d409e2c928.SpbillCreateIP,
		"attach":           __obf_6480c1d409e2c928.Attach,
		"sign_type":        __obf_6480c1d409e2c928.SignType,
		"trade_type":       __obf_6480c1d409e2c928.TradeType,
		"time_expire":      __obf_6480c1d409e2c928.TimeExpire,
		// "sign":      r.Sign,
	}

	return __obf_3cc15a7a5c34e65b

}

// UnifiedOrderResponse 统一下单响应参数
type UnifiedOrderResponse struct {
	TradeType string `xml:"trade_type"` // 交易类型
	PrepayID  string `xml:"prepay_id"`  // 预支付交易会话标识
	CodeURL   string `xml:"code_url"`   // 二维码链接，trade_type=NATIVE 时有返回
	MchID     string `xml:"mch_id"`
	AppID     string `xml:"appid"`
	Sign      string `xml:"sign"`
	NonceStr  string `xml:"nonce_str"`
	Timestamp string `xml:"timestamp,omitempty"`
	Package   string `xml:"package,omitempty"`
}

// RefundRequest 退款请求参数
type RefundRequest struct {
	XMLRequest
	TransactionID string `xml:"transaction_id,omitempty"` // 微信订单号
	OutTradeNo    string `xml:"out_trade_no,omitempty"`   // 商户订单号 (与transaction_id二选一)
	OutRefundNo   string `xml:"out_refund_no"`            // 商户退款单号
	TotalFee      int64  `xml:"total_fee"`                // 订单总金额，单位分
	RefundFee     int64  `xml:"refund_fee"`               // 退款总金额，单位分
	RefundDesc    string `xml:"refund_desc,omitempty"`    // 退款原因
	NotifyURL     string `xml:"notify_url"`               // 通知地址
	SignType      string `xml:"sign_type"`                // 签名加密类型MD5\HMAC-SHA256
	// ... 其他可选参数
}

func (__obf_6480c1d409e2c928 RefundRequest) ToMap() map[string]string {
	__obf_3cc15a7a5c34e65b := map[string]string{
		"appid":          __obf_6480c1d409e2c928.AppID,
		"mch_id":         __obf_6480c1d409e2c928.MchID,
		"nonce_str":      __obf_6480c1d409e2c928.NonceStr,
		"notify_url":     __obf_6480c1d409e2c928.NotifyURL,
		"out_refund_no":  __obf_6480c1d409e2c928.OutRefundNo,
		"total_fee":      fmt.Sprint(__obf_6480c1d409e2c928.TotalFee),
		"refund_fee":     fmt.Sprint(__obf_6480c1d409e2c928.RefundFee),
		"sign_type":      __obf_6480c1d409e2c928.SignType,
		"transaction_id": __obf_6480c1d409e2c928.TransactionID,
		"out_trade_no":   __obf_6480c1d409e2c928.OutTradeNo,
		// "sign":      r.Sign,
	}

	return __obf_3cc15a7a5c34e65b

}

// RefundResponse 退款响应参数
type RefundResponse struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	TransactionID string `xml:"transaction_id"`  // 微信订单号
	OutTradeNo    string `xml:"out_trade_no"`    // 商户订单号
	OutRefundNo   string `xml:"out_refund_no"`   // 商户退款单号
	RefundID      string `xml:"refund_id"`       // 微信退款单号
	RefundFee     int64  `xml:"refund_fee"`      // 退款总金额
	TotalFee      int64  `xml:"total_fee"`       // 订单总金额
	CashFee       int64  `xml:"cash_fee"`        // 现金支付金额
	CashRefundFee int64  `xml:"cash_refund_fee"` // 现金退款金额
	// ... 其他退款相关字段
}

// 付款到银行请求参数
type TransBankRequest struct {
	XMLRequest
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户付款单号
	EncBankNo      string `xml:"enc_bank_no"`      // 收款方银行卡号
	EncTrueName    string `xml:"enc_true_name"`    // 收款方用户名
	BankCode       string `xml:"bank_code"`        // 收款方开户行
	Amount         int64  `xml:"amount"`           // 付款金额
	Desc           string `xml:"desc,omitempty"`   // 付款说明
	// ... 其他可选参数
}

func (__obf_6480c1d409e2c928 TransBankRequest) ToMap() map[string]string {
	__obf_3cc15a7a5c34e65b := map[string]string{
		"appid":            __obf_6480c1d409e2c928.AppID,
		"mch_id":           __obf_6480c1d409e2c928.MchID,
		"nonce_str":        __obf_6480c1d409e2c928.NonceStr,
		"partner_trade_no": __obf_6480c1d409e2c928.PartnerTradeNo,
		"enc_bank_no":      __obf_6480c1d409e2c928.EncBankNo,
		"enc_true_name":    __obf_6480c1d409e2c928.EncTrueName,
		"bank_code":        __obf_6480c1d409e2c928.BankCode,
		"amount":           fmt.Sprint(__obf_6480c1d409e2c928.Amount),
	}

	return __obf_3cc15a7a5c34e65b
}

// TransBankResponse 付款到银行响应参数
type TransBankResponse struct {
	MchID          string `xml:"mch_id"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentNo      string `xml:"payment_no"`
	CmmsAmt        int    `xml:"cmms_amt"`
	Amount         int    `xml:"amount"` // 付款金额
}

// PayNotifyRequest 微信支付结果通知请求参数
// 这里的字段名需要与微信回调XML中的字段名一致 (驼峰命名或下划线命名取决于微信文档，Go的XML解析器会根据xml标签匹配)
// 微信支付回调通常使用下划线命名，如 return_code，我们使用xml标签来对应。
type PayNotifyRequest struct {
	XMLRequest
	XMLResponse
	// 业务成功时存在
	OpenID        string `xml:"openid"`             // 用户标识
	IsSubscribe   string `xml:"is_subscribe"`       // 是否关注公众账号 (Y/N)
	TradeType     string `xml:"trade_type"`         // 交易类型 (JSAPI, NATIVE, APP, etc.)
	BankType      string `xml:"bank_type"`          // 付款银行
	TotalFee      int64  `xml:"total_fee"`          // 订单总金额，单位分
	CashFee       int64  `xml:"cash_fee"`           // 现金支付金额，单位分
	FeeType       string `xml:"fee_type,omitempty"` // 现金支付货币类型
	TransactionID string `xml:"transaction_id"`     // 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no"`       // 商户系统内部订单号
	Attach        string `xml:"attach"`             // 商家数据包，原样返回
	TimeEnd       string `xml:"time_end"`           // 支付完成时间，格式为yyyyMMddHHmmss
	// Optional fields (could be omitted or have different types based on specific needs)
	CouponFee   int `xml:"coupon_fee,omitempty"`   // 代金券或立减优惠金额
	CouponCount int `xml:"coupon_count,omitempty"` // 代金券或立减优惠使用数量
	// Add more coupon details if needed (coupon_id_0, coupon_fee_0, etc.)
	// ... 其他可能返回的字段
}

func (__obf_6480c1d409e2c928 PayNotifyRequest) ToMap() map[string]string {
	return map[string]string{
		"appid":          __obf_6480c1d409e2c928.AppID,
		"mch_id":         __obf_6480c1d409e2c928.MchID,
		"nonce_str":      __obf_6480c1d409e2c928.NonceStr,
		"sign":           __obf_6480c1d409e2c928.Sign,
		"openid":         __obf_6480c1d409e2c928.OpenID,
		"is_subscribe":   __obf_6480c1d409e2c928.IsSubscribe,
		"trade_type":     __obf_6480c1d409e2c928.TradeType,
		"bank_type":      __obf_6480c1d409e2c928.BankType,
		"total_fee":      fmt.Sprint(__obf_6480c1d409e2c928.TotalFee),
		"cash_fee":       fmt.Sprint(__obf_6480c1d409e2c928.CashFee),
		"fee_type":       __obf_6480c1d409e2c928.FeeType,
		"transaction_id": __obf_6480c1d409e2c928.TransactionID,
		"out_trade_no":   __obf_6480c1d409e2c928.OutTradeNo,
		"attach":         __obf_6480c1d409e2c928.Attach,
		"time_end":       __obf_6480c1d409e2c928.TimeEnd,
		"result_code":    __obf_6480c1d409e2c928.ResultCode,
		"return_code":    __obf_6480c1d409e2c928.ReturnCode,
	}
}
