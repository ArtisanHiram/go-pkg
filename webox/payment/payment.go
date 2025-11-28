package __obf_6ccf18d8f5aab98f

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
	__obf_1607cb15609a424b *Config
	__obf_443dbb7ff0b9cfde *internal.HttpClient
}

var (
	__obf_ffae0a51b010aa7e *Payment
	__obf_eb84427a119dae32 sync.Once
)

// NewPayment 创建并返回支付模块的单例实例
// 根据配置中的 CertPath 和 KeyPath 决定是否创建带证书的客户端
func NewPayment(__obf_5a123da13f1a88bc *Config) (*Payment, error) {
	var __obf_f4364d86368fdd92 error
	__obf_eb84427a119dae32.Do(func() {
		var __obf_443dbb7ff0b9cfde *internal.HttpClient
		if len(__obf_5a123da13f1a88bc.TLSCert.Certificate) == 0 || __obf_5a123da13f1a88bc.TLSCert.PrivateKey == nil {
			__obf_443dbb7ff0b9cfde = internal.NewClient()
		} else {
			__obf_443dbb7ff0b9cfde, __obf_f4364d86368fdd92 = internal.NewClientWithTLS(__obf_5a123da13f1a88bc.TLSCert)
			if __obf_f4364d86368fdd92 != nil {
				__obf_f4364d86368fdd92 = fmt.Errorf("create payment client with TLS failed: %w", __obf_f4364d86368fdd92)
				return
			}
		}
		__obf_ffae0a51b010aa7e = &Payment{
			__obf_1607cb15609a424b: __obf_5a123da13f1a88bc,
			__obf_443dbb7ff0b9cfde: __obf_443dbb7ff0b9cfde,
		}
	})
	return __obf_ffae0a51b010aa7e, __obf_f4364d86368fdd92
}

// UnifiedOrder 统一下单
func (__obf_9b75ae39f500dd91 *Payment) UnifiedOrder(__obf_752a8380dbe117d0 uint8, __obf_7cb39cd1da0603a4 model.UnifiedOrderRequest) (*model.UnifiedOrderResponse, error) {
	if __obf_7cb39cd1da0603a4.TradeType == "JSAPI" && __obf_7cb39cd1da0603a4.OpenID == "" {
		return nil, errors.New("trade_type is JSAPI but OpenID is not provided")
	}
	__obf_7cb39cd1da0603a4.XMLRequest = model.XMLRequest{
		AppID:    __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.AppID,
		MchID:    __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_7cb39cd1da0603a4.SignType = __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.SignType                                                 // 使用配置中的签名类型
	__obf_7cb39cd1da0603a4.NotifyURL = fmt.Sprintf("%s/%d", __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.NotifyURL, __obf_752a8380dbe117d0) // 确保通知URL包含商户号
	__obf_7cb39cd1da0603a4.XMLRequest.Sign = internal.CreateSign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_7cb39cd1da0603a4.ToMap())

	// 发送XML请求，获取原始响应体
	__obf_4a6c52a6165e7781, __obf_f4364d86368fdd92 := __obf_9b75ae39f500dd91.__obf_443dbb7ff0b9cfde.PostXML(internal.PaymentUnifiedOrderURL, &__obf_7cb39cd1da0603a4)
	if __obf_f4364d86368fdd92 != nil {
		return nil, fmt.Errorf("unified order request failed: %w", __obf_f4364d86368fdd92)
	}

	// 将原始响应体解析为map，用于验签
	__obf_d08e01b9e49e8a67, __obf_f4364d86368fdd92 := internal.XMLToMap(__obf_4a6c52a6165e7781)
	if __obf_f4364d86368fdd92 != nil {
		return nil, fmt.Errorf("parse unified order response XML to map failed: %w", __obf_f4364d86368fdd92)
	}
	if __obf_d08e01b9e49e8a67["return_code"] == model.FAIL {
		return nil, fmt.Errorf("支付失败: %s", __obf_d08e01b9e49e8a67["return_msg"])
	}

	// 验签
	if !internal.VerifySign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_d08e01b9e49e8a67) {
		return nil, errors.New("unified order response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体
	var __obf_e9d852bd4884fa82 struct {
		model.XMLResponse
		model.UnifiedOrderResponse
	}
	// 解析原始响应体为map，供验签使用
	__obf_f4364d86368fdd92 = xml.Unmarshal(__obf_4a6c52a6165e7781, &__obf_e9d852bd4884fa82)
	if __obf_f4364d86368fdd92 != nil {
		return nil, fmt.Errorf("unmarshal unified order response XML to struct failed: %w", __obf_f4364d86368fdd92)
	}

	// 检查返回码和业务结果码
	if __obf_e9d852bd4884fa82.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API return error: %s - %s", __obf_e9d852bd4884fa82.ReturnCode, __obf_e9d852bd4884fa82.ReturnMsg)
	}
	if __obf_e9d852bd4884fa82.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API business error: %s - %s (detail: %s)", __obf_e9d852bd4884fa82.ErrCode, __obf_e9d852bd4884fa82.ErrCodeDes, __obf_e9d852bd4884fa82.ReturnMsg)
	}
	__obf_e9d852bd4884fa82.Timestamp = fmt.Sprint(time.Now().Unix())
	__obf_e9d852bd4884fa82.Package = fmt.Sprintf("prepay_id=%s", __obf_e9d852bd4884fa82.PrepayID)
	__obf_e9d852bd4884fa82.Sign = internal.CreateSign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, map[string]string{
		"appId":     __obf_e9d852bd4884fa82.AppID,
		"nonceStr":  __obf_e9d852bd4884fa82.NonceStr,
		"package":   __obf_e9d852bd4884fa82.Package,
		"signType":  __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.SignType,
		"timeStamp": __obf_e9d852bd4884fa82.Timestamp,
	})

	return &__obf_e9d852bd4884fa82.UnifiedOrderResponse, nil
}

// Refund 发起退款请求 (需要API证书)
// transactionID: 微信订单号 (与outTradeNo二选一)
// outTradeNo: 商户订单号 (与transactionID二选一)
// outRefundNo: 商户退款单号 (唯一)
// totalFee: 订单总金额 (单位：分)
// refundFee: 退款金额 (单位：分)
// refundDesc: 退款原因
func (__obf_9b75ae39f500dd91 *Payment) Refund(__obf_7cb39cd1da0603a4 model.RefundRequest) (string, error) {

	if __obf_7cb39cd1da0603a4.TransactionID == "" && __obf_7cb39cd1da0603a4.OutTradeNo == "" {
		return "", errors.New("transaction_id or out_trade_no must be provided")
	}

	__obf_7cb39cd1da0603a4.XMLRequest = model.XMLRequest{
		AppID:    __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.AppID,
		MchID:    __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_7cb39cd1da0603a4.SignType = __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.SignType // 使用配置中的签名类型
	__obf_7cb39cd1da0603a4.NotifyURL = __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.NotifyURL
	__obf_7cb39cd1da0603a4.XMLRequest.Sign = internal.CreateSign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_7cb39cd1da0603a4.ToMap()) // 签名放入请求结构体中

	__obf_4a6c52a6165e7781, __obf_f4364d86368fdd92 := __obf_9b75ae39f500dd91.__obf_443dbb7ff0b9cfde.PostXML(internal.PaymentRefundURL, &__obf_7cb39cd1da0603a4)
	if __obf_f4364d86368fdd92 != nil {
		return "", fmt.Errorf("refund request failed: %w", __obf_f4364d86368fdd92)
	}

	// 将原始响应体解析为map，用于验签
	__obf_d08e01b9e49e8a67, __obf_f4364d86368fdd92 := internal.XMLToMap(__obf_4a6c52a6165e7781)
	if __obf_f4364d86368fdd92 != nil {
		return "", fmt.Errorf("parse refund response XML to map failed: %w", __obf_f4364d86368fdd92)
	}

	// 验签
	if !internal.VerifySign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_d08e01b9e49e8a67) {
		return "", errors.New("refund response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_e9d852bd4884fa82 struct {
		model.XMLResponse
		model.RefundResponse
	}
	__obf_f4364d86368fdd92 = xml.Unmarshal(__obf_4a6c52a6165e7781, &__obf_e9d852bd4884fa82)
	if __obf_f4364d86368fdd92 != nil {
		return "", fmt.Errorf("unmarshal refund response XML to struct failed: %w", __obf_f4364d86368fdd92)
	}

	if __obf_e9d852bd4884fa82.ReturnCode != model.SUCCESS {
		return "", fmt.Errorf("refund API return error: %s - %s", __obf_e9d852bd4884fa82.ReturnCode, __obf_e9d852bd4884fa82.ReturnMsg)
	}
	if __obf_e9d852bd4884fa82.ResultCode != model.SUCCESS {
		return "", fmt.Errorf("refund API business error: %s - %s (detail: %s)", __obf_e9d852bd4884fa82.ErrCode, __obf_e9d852bd4884fa82.ErrCodeDes, __obf_e9d852bd4884fa82.ReturnMsg)
	}

	return string(__obf_4a6c52a6165e7781), nil
}

func (__obf_9b75ae39f500dd91 *Payment) PayNotify(__obf_1ff37d1426bfdde7 []byte, __obf_b8bac2e81d270319 func(model.PayNotifyRequest) error) error {

	// 6. 将原始XML解析为结构体 (用于更方便地访问字段)
	var __obf_7cb39cd1da0603a4 model.PayNotifyRequest
	__obf_f4364d86368fdd92 := xml.Unmarshal(__obf_1ff37d1426bfdde7, &__obf_7cb39cd1da0603a4)
	if __obf_f4364d86368fdd92 != nil {
		return fmt.Errorf("Payment Notify: Unmarshal XML to struct failed: %v, Body: %s", __obf_f4364d86368fdd92, string(__obf_1ff37d1426bfdde7))
	}

	// 即使业务失败，如果通信成功 (ReturnCode是SUCCESS)，也应该返回 SUCCESS 给微信，
	// 告诉微信你已经收到通知，避免微信重复发送。
	// 只有在通信本身出错时 (如读取Body失败，验签失败等) 才返回 FAIL

	if __obf_7cb39cd1da0603a4.ReturnCode != model.SUCCESS {
		return fmt.Errorf("Payment Notify: ReturnCode is not SUCCESS: %s - %s", __obf_7cb39cd1da0603a4.ReturnCode, __obf_7cb39cd1da0603a4.ReturnMsg)
	}

	// 7. 检查业务结果 (ResultCode)
	if __obf_7cb39cd1da0603a4.ResultCode == model.SUCCESS {
		// 5. 验签
		// VerifySign 会在内部移除 reqMap 中的 "sign" 字段，所以传入前需要复制或注意。
		// 这里直接传入 reqMap 是可以的，因为验签后我们通常不再使用这个 map 进行后续操作。
		if !internal.VerifySign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_7cb39cd1da0603a4.ToMap()) {
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
		return __obf_b8bac2e81d270319(__obf_7cb39cd1da0603a4)
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
func ReturnMessage(__obf_75803e62173152a6, __obf_2c01b95e20f3e19e string) []byte {
	__obf_d08e01b9e49e8a67 := map[string]string{
		"return_code": __obf_75803e62173152a6,
		"return_msg":  __obf_2c01b95e20f3e19e,
	}
	return internal.MapToXML(__obf_d08e01b9e49e8a67)
}

func (__obf_9b75ae39f500dd91 *Payment) TransBank(__obf_7cb39cd1da0603a4 model.TransBankRequest) (*model.TransBankResponse, error) {

	if __obf_7cb39cd1da0603a4.PartnerTradeNo == "" && __obf_7cb39cd1da0603a4.EncBankNo == "" && __obf_7cb39cd1da0603a4.EncTrueName == "" && __obf_7cb39cd1da0603a4.BankCode == "" {
		return nil, errors.New("请提供商户付款单号、收款方银行卡号、收款方用户名、收款方开户行")
	}

	__obf_7cb39cd1da0603a4.XMLRequest = model.XMLRequest{
		AppID:    __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.AppID,
		MchID:    __obf_9b75ae39f500dd91.__obf_1607cb15609a424b.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_7cb39cd1da0603a4.XMLRequest.Sign = internal.CreateSign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_7cb39cd1da0603a4.ToMap()) // 签名放入请求结构体中

	__obf_4a6c52a6165e7781, __obf_f4364d86368fdd92 := __obf_9b75ae39f500dd91.__obf_443dbb7ff0b9cfde.PostXML(internal.PaymentTransURL, &__obf_7cb39cd1da0603a4)
	if __obf_f4364d86368fdd92 != nil {
		return nil, fmt.Errorf("pay_bank request failed: %w", __obf_f4364d86368fdd92)
	}

	// 将原始响应体解析为map，用于验签
	__obf_d08e01b9e49e8a67, __obf_f4364d86368fdd92 := internal.XMLToMap(__obf_4a6c52a6165e7781)
	if __obf_f4364d86368fdd92 != nil {
		return nil, fmt.Errorf("parse pay_bank response XML to map failed: %w", __obf_f4364d86368fdd92)
	}

	// 验签
	if !internal.VerifySign(__obf_9b75ae39f500dd91.__obf_1607cb15609a424b.APIKey, __obf_d08e01b9e49e8a67) {
		return nil, errors.New("pay_bank response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_e9d852bd4884fa82 struct {
		model.XMLResponse
		model.TransBankResponse
	}
	__obf_f4364d86368fdd92 = xml.Unmarshal(__obf_4a6c52a6165e7781, &__obf_e9d852bd4884fa82)
	if __obf_f4364d86368fdd92 != nil {
		return nil, fmt.Errorf("unmarshal pay_bank response XML to struct failed: %w", __obf_f4364d86368fdd92)
	}

	if __obf_e9d852bd4884fa82.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API return error: %s - %s", __obf_e9d852bd4884fa82.ReturnCode, __obf_e9d852bd4884fa82.ReturnMsg)
	}
	if __obf_e9d852bd4884fa82.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API business error: %s - %s (detail: %s)", __obf_e9d852bd4884fa82.ErrCode, __obf_e9d852bd4884fa82.ErrCodeDes, __obf_e9d852bd4884fa82.ReturnMsg)
	}

	return &__obf_e9d852bd4884fa82.TransBankResponse, nil
}
