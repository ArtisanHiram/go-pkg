package __obf_db5b065626995a13

import (
	"encoding/xml" // 确保导入 encoding/xml
	"errors"
	"fmt"
	internal "github.com/ArtisanHiram/go-pkg/webox/internal"
	model "github.com/ArtisanHiram/go-pkg/webox/model"
	"sync"
	"time"
)

// Payment 支付模块实例
type Payment struct {
	__obf_1b760bc57b91a6a5 *Config
	__obf_afd32fe766822203 *internal.HttpClient
}

var (
	__obf_11dcb39763a29b0d *Payment
	__obf_09a4c0d83c7c95e1 sync.Once
)

// NewPayment 创建并返回支付模块的单例实例
// 根据配置中的 CertPath 和 KeyPath 决定是否创建带证书的客户端
func NewPayment(__obf_72917112532a715f *Config) (*Payment, error) {
	var __obf_48e3d1ffda112f63 error
	__obf_09a4c0d83c7c95e1.
		Do(func() {
			var __obf_afd32fe766822203 *internal.HttpClient
			if len(__obf_72917112532a715f.TLSCert.Certificate) == 0 || __obf_72917112532a715f.TLSCert.PrivateKey == nil {
				__obf_afd32fe766822203 = internal.NewClient()
			} else {
				__obf_afd32fe766822203, __obf_48e3d1ffda112f63 = internal.NewClientWithTLS(__obf_72917112532a715f.TLSCert)
				if __obf_48e3d1ffda112f63 != nil {
					__obf_48e3d1ffda112f63 = fmt.Errorf("create payment client with TLS failed: %w", __obf_48e3d1ffda112f63)
					return
				}
			}
			__obf_11dcb39763a29b0d = &Payment{__obf_1b760bc57b91a6a5: __obf_72917112532a715f, __obf_afd32fe766822203: __obf_afd32fe766822203}
		})
	return __obf_11dcb39763a29b0d, __obf_48e3d1ffda112f63
}

// UnifiedOrder 统一下单
func (__obf_c5ccd68480ddd84c *Payment) UnifiedOrder(__obf_74a7bee6f4ad7257 uint8, __obf_1d3053f9bb6f0284 model.UnifiedOrderRequest) (*model.UnifiedOrderResponse, error) {
	if __obf_1d3053f9bb6f0284.TradeType == "JSAPI" && __obf_1d3053f9bb6f0284.OpenID == "" {
		return nil, errors.New("trade_type is JSAPI but OpenID is not provided")
	}
	__obf_1d3053f9bb6f0284.
		XMLRequest = model.XMLRequest{
		AppID:    __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.AppID,
		MchID:    __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_1d3053f9bb6f0284.
		SignType = __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.SignType
	__obf_1d3053f9bb6f0284. // 使用配置中的签名类型
				NotifyURL = fmt.Sprintf("%s/%d", __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.NotifyURL, __obf_74a7bee6f4ad7257)
	__obf_1d3053f9bb6f0284. // 确保通知URL包含商户号
				XMLRequest.Sign = internal.CreateSign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_1d3053f9bb6f0284.ToMap())
	__obf_c485ef61f62c5ec0,

		// 发送XML请求，获取原始响应体
		__obf_48e3d1ffda112f63 := __obf_c5ccd68480ddd84c.__obf_afd32fe766822203.PostXML(internal.PaymentUnifiedOrderURL, &__obf_1d3053f9bb6f0284)
	if __obf_48e3d1ffda112f63 != nil {
		return nil, fmt.Errorf("unified order request failed: %w", __obf_48e3d1ffda112f63)
	}
	__obf_5c9e3b637dfae1f6,

		// 将原始响应体解析为map，用于验签
		__obf_48e3d1ffda112f63 := internal.XMLToMap(__obf_c485ef61f62c5ec0)
	if __obf_48e3d1ffda112f63 != nil {
		return nil, fmt.Errorf("parse unified order response XML to map failed: %w", __obf_48e3d1ffda112f63)
	}
	if __obf_5c9e3b637dfae1f6["return_code"] == model.FAIL {
		return nil, fmt.Errorf("支付失败: %s", __obf_5c9e3b637dfae1f6["return_msg"])
	}

	// 验签
	if !internal.VerifySign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_5c9e3b637dfae1f6) {
		return nil, errors.New("unified order response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体
	var __obf_f1ad3eec94891594 struct {
		model.XMLResponse
		model.UnifiedOrderResponse
	}
	__obf_48e3d1ffda112f63 = // 解析原始响应体为map，供验签使用
		xml.Unmarshal(__obf_c485ef61f62c5ec0, &__obf_f1ad3eec94891594)
	if __obf_48e3d1ffda112f63 != nil {
		return nil, fmt.Errorf("unmarshal unified order response XML to struct failed: %w", __obf_48e3d1ffda112f63)
	}

	// 检查返回码和业务结果码
	if __obf_f1ad3eec94891594.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API return error: %s - %s", __obf_f1ad3eec94891594.ReturnCode, __obf_f1ad3eec94891594.ReturnMsg)
	}
	if __obf_f1ad3eec94891594.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API business error: %s - %s (detail: %s)", __obf_f1ad3eec94891594.ErrCode, __obf_f1ad3eec94891594.ErrCodeDes, __obf_f1ad3eec94891594.ReturnMsg)
	}
	__obf_f1ad3eec94891594.
		Timestamp = fmt.Sprint(time.Now().Unix())
	__obf_f1ad3eec94891594.
		Package = fmt.Sprintf("prepay_id=%s", __obf_f1ad3eec94891594.PrepayID)
	__obf_f1ad3eec94891594.
		Sign = internal.CreateSign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, map[string]string{
		"appId":     __obf_f1ad3eec94891594.AppID,
		"nonceStr":  __obf_f1ad3eec94891594.NonceStr,
		"package":   __obf_f1ad3eec94891594.Package,
		"signType":  __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.SignType,
		"timeStamp": __obf_f1ad3eec94891594.Timestamp,
	})

	return &__obf_f1ad3eec94891594.UnifiedOrderResponse, nil
}

// Refund 发起退款请求 (需要API证书)
// transactionID: 微信订单号 (与outTradeNo二选一)
// outTradeNo: 商户订单号 (与transactionID二选一)
// outRefundNo: 商户退款单号 (唯一)
// totalFee: 订单总金额 (单位：分)
// refundFee: 退款金额 (单位：分)
// refundDesc: 退款原因
func (__obf_c5ccd68480ddd84c *Payment) Refund(__obf_1d3053f9bb6f0284 model.RefundRequest) (string, error) {

	if __obf_1d3053f9bb6f0284.TransactionID == "" && __obf_1d3053f9bb6f0284.OutTradeNo == "" {
		return "", errors.New("transaction_id or out_trade_no must be provided")
	}
	__obf_1d3053f9bb6f0284.
		XMLRequest = model.XMLRequest{
		AppID:    __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.AppID,
		MchID:    __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_1d3053f9bb6f0284.
		SignType = __obf_c5ccd68480ddd84c. // 使用配置中的签名类型
		__obf_1b760bc57b91a6a5.SignType
	__obf_1d3053f9bb6f0284.
		NotifyURL = __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.NotifyURL
	__obf_1d3053f9bb6f0284.
		XMLRequest.Sign = internal.CreateSign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_1d3053f9bb6f0284.ToMap())
	__obf_c485ef61f62c5ec0, // 签名放入请求结构体中
		__obf_48e3d1ffda112f63 := __obf_c5ccd68480ddd84c.__obf_afd32fe766822203.PostXML(internal.PaymentRefundURL, &__obf_1d3053f9bb6f0284)
	if __obf_48e3d1ffda112f63 != nil {
		return "", fmt.Errorf("refund request failed: %w", __obf_48e3d1ffda112f63)
	}
	__obf_5c9e3b637dfae1f6,

		// 将原始响应体解析为map，用于验签
		__obf_48e3d1ffda112f63 := internal.XMLToMap(__obf_c485ef61f62c5ec0)
	if __obf_48e3d1ffda112f63 != nil {
		return "", fmt.Errorf("parse refund response XML to map failed: %w", __obf_48e3d1ffda112f63)
	}

	// 验签
	if !internal.VerifySign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_5c9e3b637dfae1f6) {
		return "", errors.New("refund response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_f1ad3eec94891594 struct {
		model.XMLResponse
		model.RefundResponse
	}
	__obf_48e3d1ffda112f63 = xml.Unmarshal(__obf_c485ef61f62c5ec0, &__obf_f1ad3eec94891594)
	if __obf_48e3d1ffda112f63 != nil {
		return "", fmt.Errorf("unmarshal refund response XML to struct failed: %w", __obf_48e3d1ffda112f63)
	}

	if __obf_f1ad3eec94891594.ReturnCode != model.SUCCESS {
		return "", fmt.Errorf("refund API return error: %s - %s", __obf_f1ad3eec94891594.ReturnCode, __obf_f1ad3eec94891594.ReturnMsg)
	}
	if __obf_f1ad3eec94891594.ResultCode != model.SUCCESS {
		return "", fmt.Errorf("refund API business error: %s - %s (detail: %s)", __obf_f1ad3eec94891594.ErrCode, __obf_f1ad3eec94891594.ErrCodeDes, __obf_f1ad3eec94891594.ReturnMsg)
	}

	return string(__obf_c485ef61f62c5ec0), nil
}

func (__obf_c5ccd68480ddd84c *Payment) PayNotify(__obf_063d2cca25737686 []byte, __obf_14bf194e2a6301f8 func(model.PayNotifyRequest) error) error {

	// 6. 将原始XML解析为结构体 (用于更方便地访问字段)
	var __obf_1d3053f9bb6f0284 model.PayNotifyRequest
	__obf_48e3d1ffda112f63 := xml.Unmarshal(__obf_063d2cca25737686, &__obf_1d3053f9bb6f0284)
	if __obf_48e3d1ffda112f63 != nil {
		return fmt.Errorf("Payment Notify: Unmarshal XML to struct failed: %v, Body: %s", __obf_48e3d1ffda112f63, string(__obf_063d2cca25737686))
	}

	// 即使业务失败，如果通信成功 (ReturnCode是SUCCESS)，也应该返回 SUCCESS 给微信，
	// 告诉微信你已经收到通知，避免微信重复发送。
	// 只有在通信本身出错时 (如读取Body失败，验签失败等) 才返回 FAIL

	if __obf_1d3053f9bb6f0284.ReturnCode != model.SUCCESS {
		return fmt.Errorf("Payment Notify: ReturnCode is not SUCCESS: %s - %s", __obf_1d3053f9bb6f0284.ReturnCode, __obf_1d3053f9bb6f0284.ReturnMsg)
	}

	// 7. 检查业务结果 (ResultCode)
	if __obf_1d3053f9bb6f0284.ResultCode == model.SUCCESS {
		// 5. 验签
		// VerifySign 会在内部移除 reqMap 中的 "sign" 字段，所以传入前需要复制或注意。
		// 这里直接传入 reqMap 是可以的，因为验签后我们通常不再使用这个 map 进行后续操作。
		if !internal.VerifySign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_1d3053f9bb6f0284.ToMap()) {
			return errors.New("signature verification failed")
		}
		// 8. TODO: 在这里处理你的业务逻辑
		// 这一步是支付回调最核心的部分，务必保证以下几点：
		// a. **幂等性处理：**
		//    - 根据 `OutTradeNo` (商户订单号) 查询你自己的订单系统。
		//    - 检查订单当前状态，如果已经处理过（例如，已经设置为“已支付”），则直接返回 SUCCESS 给微信，避免重复处理。
		//    - 可以使用分布式锁或数据库事务来确保并发通知下的幂等性。
		// b. **数据校验：**
		//    - 验证 `notifyReq.AppID` 和 `notifyReq.MchID` 是否与你的配置一致，防止其他商户的通知。
		//    - **重点：** 验证 `notifyReq.TotalFee` (支付金额) 是否与你订单系统中的金额一致，防止金额被篡改。
		//    - 验证 `notifyReq.OutTradeNo` 是否是你系统发起的有效订单号。
		// c. **更新订单状态：**
		//    - 将你的订单状态更新为“已支付”或“交易完成”。
		//    - 记录微信支付的 `TransactionID`，用于后续的退款、对账等操作。
		//    - 如果有 `attach` 字段，解析并使用其中包含的业务数据。
		// d. **后续业务流程：**
		//    - 触发后续业务流程，例如：发货、充值、开通会员、发送短信通知用户等。

		// 模拟业务处理
		// 假设这里查询数据库并更新订单状态
		return __obf_14bf194e2a6301f8(__obf_1d3053f9bb6f0284)
		// orderID := req.OutTradeNo
		// actualAmount := req.TotalFee // 你的订单实际应收金额 (从数据库获取)
		// if actualAmount != req.TotalFee {
		// 	return ReturnMessage(model.FAIL, fmt.Sprintf("amount mismatch for order %s: expected %d, got %d", orderID, actualAmount, req.TotalFee))
		// }
	}
	return nil
}

// Notify 向微信返回响应
// data参数是需要转换为XML的map[string]string，如 {"return_code": model.SUCCESS, "return_msg": "OK"}
func ReturnMessage(__obf_26ab1e4baad69bd6, __obf_338c49da3e27dd83 string) []byte {
	__obf_5c9e3b637dfae1f6 := map[string]string{
		"return_code": __obf_26ab1e4baad69bd6,
		"return_msg":  __obf_338c49da3e27dd83,
	}
	return internal.MapToXML(__obf_5c9e3b637dfae1f6)
}

func (__obf_c5ccd68480ddd84c *Payment) TransBank(__obf_1d3053f9bb6f0284 model.TransBankRequest) (*model.TransBankResponse, error) {

	if __obf_1d3053f9bb6f0284.PartnerTradeNo == "" && __obf_1d3053f9bb6f0284.EncBankNo == "" && __obf_1d3053f9bb6f0284.EncTrueName == "" && __obf_1d3053f9bb6f0284.BankCode == "" {
		return nil, errors.New("请提供商户付款单号、收款方银行卡号、收款方用户名、收款方开户行")
	}
	__obf_1d3053f9bb6f0284.
		XMLRequest = model.XMLRequest{
		AppID:    __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.AppID,
		MchID:    __obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_1d3053f9bb6f0284.
		XMLRequest.Sign = internal.CreateSign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_1d3053f9bb6f0284.ToMap())
	__obf_c485ef61f62c5ec0, // 签名放入请求结构体中
		__obf_48e3d1ffda112f63 := __obf_c5ccd68480ddd84c.__obf_afd32fe766822203.PostXML(internal.PaymentTransURL, &__obf_1d3053f9bb6f0284)
	if __obf_48e3d1ffda112f63 != nil {
		return nil, fmt.Errorf("pay_bank request failed: %w", __obf_48e3d1ffda112f63)
	}
	__obf_5c9e3b637dfae1f6,

		// 将原始响应体解析为map，用于验签
		__obf_48e3d1ffda112f63 := internal.XMLToMap(__obf_c485ef61f62c5ec0)
	if __obf_48e3d1ffda112f63 != nil {
		return nil, fmt.Errorf("parse pay_bank response XML to map failed: %w", __obf_48e3d1ffda112f63)
	}

	// 验签
	if !internal.VerifySign(__obf_c5ccd68480ddd84c.__obf_1b760bc57b91a6a5.APIKey, __obf_5c9e3b637dfae1f6) {
		return nil, errors.New("pay_bank response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_f1ad3eec94891594 struct {
		model.XMLResponse
		model.TransBankResponse
	}
	__obf_48e3d1ffda112f63 = xml.Unmarshal(__obf_c485ef61f62c5ec0, &__obf_f1ad3eec94891594)
	if __obf_48e3d1ffda112f63 != nil {
		return nil, fmt.Errorf("unmarshal pay_bank response XML to struct failed: %w", __obf_48e3d1ffda112f63)
	}

	if __obf_f1ad3eec94891594.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API return error: %s - %s", __obf_f1ad3eec94891594.ReturnCode, __obf_f1ad3eec94891594.ReturnMsg)
	}
	if __obf_f1ad3eec94891594.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API business error: %s - %s (detail: %s)", __obf_f1ad3eec94891594.ErrCode, __obf_f1ad3eec94891594.ErrCodeDes, __obf_f1ad3eec94891594.ReturnMsg)
	}

	return &__obf_f1ad3eec94891594.TransBankResponse, nil
}
