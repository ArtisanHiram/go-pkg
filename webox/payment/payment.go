package __obf_7d0a8d04d1ebfa9c

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
	__obf_65d5ccebfefd642b *Config
	__obf_3581d3195030f5eb *internal.HttpClient
}

var (
	__obf_7ff5a5e62179ce89 *Payment
	__obf_2b07bc76a8904441 sync.Once
)

// NewPayment 创建并返回支付模块的单例实例
// 根据配置中的 CertPath 和 KeyPath 决定是否创建带证书的客户端
func NewPayment(__obf_249b289248ec4537 *Config) (*Payment, error) {
	var __obf_679831abd2fe70b9 error
	__obf_2b07bc76a8904441.Do(func() {
		var __obf_3581d3195030f5eb *internal.HttpClient
		if len(__obf_249b289248ec4537.TLSCert.Certificate) == 0 || __obf_249b289248ec4537.TLSCert.PrivateKey == nil {
			__obf_3581d3195030f5eb = internal.NewClient()
		} else {
			__obf_3581d3195030f5eb, __obf_679831abd2fe70b9 = internal.NewClientWithTLS(__obf_249b289248ec4537.TLSCert)
			if __obf_679831abd2fe70b9 != nil {
				__obf_679831abd2fe70b9 = fmt.Errorf("create payment client with TLS failed: %w", __obf_679831abd2fe70b9)
				return
			}
		}
		__obf_7ff5a5e62179ce89 = &Payment{
			__obf_65d5ccebfefd642b: __obf_249b289248ec4537,
			__obf_3581d3195030f5eb: __obf_3581d3195030f5eb,
		}
	})
	return __obf_7ff5a5e62179ce89, __obf_679831abd2fe70b9
}

// UnifiedOrder 统一下单
func (__obf_ee2a8f52dc9fd2bf *Payment) UnifiedOrder(__obf_552f2e255e245f16 uint8, __obf_5ebddbe4ce24fce3 model.UnifiedOrderRequest) (*model.UnifiedOrderResponse, error) {
	if __obf_5ebddbe4ce24fce3.TradeType == "JSAPI" && __obf_5ebddbe4ce24fce3.OpenID == "" {
		return nil, errors.New("trade_type is JSAPI but OpenID is not provided")
	}
	__obf_5ebddbe4ce24fce3.XMLRequest = model.XMLRequest{
		AppID:    __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.AppID,
		MchID:    __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_5ebddbe4ce24fce3.SignType = __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.SignType                                                 // 使用配置中的签名类型
	__obf_5ebddbe4ce24fce3.NotifyURL = fmt.Sprintf("%s/%d", __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.NotifyURL, __obf_552f2e255e245f16) // 确保通知URL包含商户号
	__obf_5ebddbe4ce24fce3.XMLRequest.Sign = internal.CreateSign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_5ebddbe4ce24fce3.ToMap())

	// 发送XML请求，获取原始响应体
	__obf_2e6463f93fb1448a, __obf_679831abd2fe70b9 := __obf_ee2a8f52dc9fd2bf.__obf_3581d3195030f5eb.PostXML(internal.PaymentUnifiedOrderURL, &__obf_5ebddbe4ce24fce3)
	if __obf_679831abd2fe70b9 != nil {
		return nil, fmt.Errorf("unified order request failed: %w", __obf_679831abd2fe70b9)
	}

	// 将原始响应体解析为map，用于验签
	__obf_38271f28ca66b2d3, __obf_679831abd2fe70b9 := internal.XMLToMap(__obf_2e6463f93fb1448a)
	if __obf_679831abd2fe70b9 != nil {
		return nil, fmt.Errorf("parse unified order response XML to map failed: %w", __obf_679831abd2fe70b9)
	}
	if __obf_38271f28ca66b2d3["return_code"] == model.FAIL {
		return nil, fmt.Errorf("支付失败: %s", __obf_38271f28ca66b2d3["return_msg"])
	}

	// 验签
	if !internal.VerifySign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_38271f28ca66b2d3) {
		return nil, errors.New("unified order response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体
	var __obf_a1bfdb946a2657c6 struct {
		model.XMLResponse
		model.UnifiedOrderResponse
	}
	// 解析原始响应体为map，供验签使用
	__obf_679831abd2fe70b9 = xml.Unmarshal(__obf_2e6463f93fb1448a, &__obf_a1bfdb946a2657c6)
	if __obf_679831abd2fe70b9 != nil {
		return nil, fmt.Errorf("unmarshal unified order response XML to struct failed: %w", __obf_679831abd2fe70b9)
	}

	// 检查返回码和业务结果码
	if __obf_a1bfdb946a2657c6.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API return error: %s - %s", __obf_a1bfdb946a2657c6.ReturnCode, __obf_a1bfdb946a2657c6.ReturnMsg)
	}
	if __obf_a1bfdb946a2657c6.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API business error: %s - %s (detail: %s)", __obf_a1bfdb946a2657c6.ErrCode, __obf_a1bfdb946a2657c6.ErrCodeDes, __obf_a1bfdb946a2657c6.ReturnMsg)
	}
	__obf_a1bfdb946a2657c6.Timestamp = fmt.Sprint(time.Now().Unix())
	__obf_a1bfdb946a2657c6.Package = fmt.Sprintf("prepay_id=%s", __obf_a1bfdb946a2657c6.PrepayID)
	__obf_a1bfdb946a2657c6.Sign = internal.CreateSign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, map[string]string{
		"appId":     __obf_a1bfdb946a2657c6.AppID,
		"nonceStr":  __obf_a1bfdb946a2657c6.NonceStr,
		"package":   __obf_a1bfdb946a2657c6.Package,
		"signType":  __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.SignType,
		"timeStamp": __obf_a1bfdb946a2657c6.Timestamp,
	})

	return &__obf_a1bfdb946a2657c6.UnifiedOrderResponse, nil
}

// Refund 发起退款请求 (需要API证书)
// transactionID: 微信订单号 (与outTradeNo二选一)
// outTradeNo: 商户订单号 (与transactionID二选一)
// outRefundNo: 商户退款单号 (唯一)
// totalFee: 订单总金额 (单位：分)
// refundFee: 退款金额 (单位：分)
// refundDesc: 退款原因
func (__obf_ee2a8f52dc9fd2bf *Payment) Refund(__obf_5ebddbe4ce24fce3 model.RefundRequest) (string, error) {

	if __obf_5ebddbe4ce24fce3.TransactionID == "" && __obf_5ebddbe4ce24fce3.OutTradeNo == "" {
		return "", errors.New("transaction_id or out_trade_no must be provided")
	}

	__obf_5ebddbe4ce24fce3.XMLRequest = model.XMLRequest{
		AppID:    __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.AppID,
		MchID:    __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_5ebddbe4ce24fce3.SignType = __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.SignType // 使用配置中的签名类型
	__obf_5ebddbe4ce24fce3.NotifyURL = __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.NotifyURL
	__obf_5ebddbe4ce24fce3.XMLRequest.Sign = internal.CreateSign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_5ebddbe4ce24fce3.ToMap()) // 签名放入请求结构体中

	__obf_2e6463f93fb1448a, __obf_679831abd2fe70b9 := __obf_ee2a8f52dc9fd2bf.__obf_3581d3195030f5eb.PostXML(internal.PaymentRefundURL, &__obf_5ebddbe4ce24fce3)
	if __obf_679831abd2fe70b9 != nil {
		return "", fmt.Errorf("refund request failed: %w", __obf_679831abd2fe70b9)
	}

	// 将原始响应体解析为map，用于验签
	__obf_38271f28ca66b2d3, __obf_679831abd2fe70b9 := internal.XMLToMap(__obf_2e6463f93fb1448a)
	if __obf_679831abd2fe70b9 != nil {
		return "", fmt.Errorf("parse refund response XML to map failed: %w", __obf_679831abd2fe70b9)
	}

	// 验签
	if !internal.VerifySign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_38271f28ca66b2d3) {
		return "", errors.New("refund response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_a1bfdb946a2657c6 struct {
		model.XMLResponse
		model.RefundResponse
	}
	__obf_679831abd2fe70b9 = xml.Unmarshal(__obf_2e6463f93fb1448a, &__obf_a1bfdb946a2657c6)
	if __obf_679831abd2fe70b9 != nil {
		return "", fmt.Errorf("unmarshal refund response XML to struct failed: %w", __obf_679831abd2fe70b9)
	}

	if __obf_a1bfdb946a2657c6.ReturnCode != model.SUCCESS {
		return "", fmt.Errorf("refund API return error: %s - %s", __obf_a1bfdb946a2657c6.ReturnCode, __obf_a1bfdb946a2657c6.ReturnMsg)
	}
	if __obf_a1bfdb946a2657c6.ResultCode != model.SUCCESS {
		return "", fmt.Errorf("refund API business error: %s - %s (detail: %s)", __obf_a1bfdb946a2657c6.ErrCode, __obf_a1bfdb946a2657c6.ErrCodeDes, __obf_a1bfdb946a2657c6.ReturnMsg)
	}

	return string(__obf_2e6463f93fb1448a), nil
}

func (__obf_ee2a8f52dc9fd2bf *Payment) PayNotify(__obf_0dfc8259fb0e5a8b []byte, __obf_c202fce4e9609352 func(model.PayNotifyRequest) error) error {

	// 6. 将原始XML解析为结构体 (用于更方便地访问字段)
	var __obf_5ebddbe4ce24fce3 model.PayNotifyRequest
	__obf_679831abd2fe70b9 := xml.Unmarshal(__obf_0dfc8259fb0e5a8b, &__obf_5ebddbe4ce24fce3)
	if __obf_679831abd2fe70b9 != nil {
		return fmt.Errorf("Payment Notify: Unmarshal XML to struct failed: %v, Body: %s", __obf_679831abd2fe70b9, string(__obf_0dfc8259fb0e5a8b))
	}

	// 即使业务失败，如果通信成功 (ReturnCode是SUCCESS)，也应该返回 SUCCESS 给微信，
	// 告诉微信你已经收到通知，避免微信重复发送。
	// 只有在通信本身出错时 (如读取Body失败，验签失败等) 才返回 FAIL

	if __obf_5ebddbe4ce24fce3.ReturnCode != model.SUCCESS {
		return fmt.Errorf("Payment Notify: ReturnCode is not SUCCESS: %s - %s", __obf_5ebddbe4ce24fce3.ReturnCode, __obf_5ebddbe4ce24fce3.ReturnMsg)
	}

	// 7. 检查业务结果 (ResultCode)
	if __obf_5ebddbe4ce24fce3.ResultCode == model.SUCCESS {
		// 5. 验签
		// VerifySign 会在内部移除 reqMap 中的 "sign" 字段，所以传入前需要复制或注意。
		// 这里直接传入 reqMap 是可以的，因为验签后我们通常不再使用这个 map 进行后续操作。
		if !internal.VerifySign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_5ebddbe4ce24fce3.ToMap()) {
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
		return __obf_c202fce4e9609352(__obf_5ebddbe4ce24fce3)
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
func ReturnMessage(__obf_78b97fad45ca7753, __obf_82f198350262c448 string) []byte {
	__obf_38271f28ca66b2d3 := map[string]string{
		"return_code": __obf_78b97fad45ca7753,
		"return_msg":  __obf_82f198350262c448,
	}
	return internal.MapToXML(__obf_38271f28ca66b2d3)
}

func (__obf_ee2a8f52dc9fd2bf *Payment) TransBank(__obf_5ebddbe4ce24fce3 model.TransBankRequest) (*model.TransBankResponse, error) {

	if __obf_5ebddbe4ce24fce3.PartnerTradeNo == "" && __obf_5ebddbe4ce24fce3.EncBankNo == "" && __obf_5ebddbe4ce24fce3.EncTrueName == "" && __obf_5ebddbe4ce24fce3.BankCode == "" {
		return nil, errors.New("请提供商户付款单号、收款方银行卡号、收款方用户名、收款方开户行")
	}

	__obf_5ebddbe4ce24fce3.XMLRequest = model.XMLRequest{
		AppID:    __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.AppID,
		MchID:    __obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_5ebddbe4ce24fce3.XMLRequest.Sign = internal.CreateSign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_5ebddbe4ce24fce3.ToMap()) // 签名放入请求结构体中

	__obf_2e6463f93fb1448a, __obf_679831abd2fe70b9 := __obf_ee2a8f52dc9fd2bf.__obf_3581d3195030f5eb.PostXML(internal.PaymentTransURL, &__obf_5ebddbe4ce24fce3)
	if __obf_679831abd2fe70b9 != nil {
		return nil, fmt.Errorf("pay_bank request failed: %w", __obf_679831abd2fe70b9)
	}

	// 将原始响应体解析为map，用于验签
	__obf_38271f28ca66b2d3, __obf_679831abd2fe70b9 := internal.XMLToMap(__obf_2e6463f93fb1448a)
	if __obf_679831abd2fe70b9 != nil {
		return nil, fmt.Errorf("parse pay_bank response XML to map failed: %w", __obf_679831abd2fe70b9)
	}

	// 验签
	if !internal.VerifySign(__obf_ee2a8f52dc9fd2bf.__obf_65d5ccebfefd642b.APIKey, __obf_38271f28ca66b2d3) {
		return nil, errors.New("pay_bank response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_a1bfdb946a2657c6 struct {
		model.XMLResponse
		model.TransBankResponse
	}
	__obf_679831abd2fe70b9 = xml.Unmarshal(__obf_2e6463f93fb1448a, &__obf_a1bfdb946a2657c6)
	if __obf_679831abd2fe70b9 != nil {
		return nil, fmt.Errorf("unmarshal pay_bank response XML to struct failed: %w", __obf_679831abd2fe70b9)
	}

	if __obf_a1bfdb946a2657c6.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API return error: %s - %s", __obf_a1bfdb946a2657c6.ReturnCode, __obf_a1bfdb946a2657c6.ReturnMsg)
	}
	if __obf_a1bfdb946a2657c6.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API business error: %s - %s (detail: %s)", __obf_a1bfdb946a2657c6.ErrCode, __obf_a1bfdb946a2657c6.ErrCodeDes, __obf_a1bfdb946a2657c6.ReturnMsg)
	}

	return &__obf_a1bfdb946a2657c6.TransBankResponse, nil
}
