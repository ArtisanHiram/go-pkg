package __obf_b208cbe38c7e3325

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
	__obf_2d2cf9566604eb22 *Config
	__obf_9786d25234c49e0b *internal.HttpClient
}

var (
	__obf_5fe4c1650c09a964 *Payment
	__obf_314204671e77fe7d sync.Once
)

// NewPayment 创建并返回支付模块的单例实例
// 根据配置中的 CertPath 和 KeyPath 决定是否创建带证书的客户端
func NewPayment(__obf_4955881acf1d99f7 *Config) (*Payment, error) {
	var __obf_4a2f7560afdd36ef error
	__obf_314204671e77fe7d.
		Do(func() {
			var __obf_9786d25234c49e0b *internal.HttpClient
			if len(__obf_4955881acf1d99f7.TLSCert.Certificate) == 0 || __obf_4955881acf1d99f7.TLSCert.PrivateKey == nil {
				__obf_9786d25234c49e0b = internal.NewClient()
			} else {
				__obf_9786d25234c49e0b, __obf_4a2f7560afdd36ef = internal.NewClientWithTLS(__obf_4955881acf1d99f7.TLSCert)
				if __obf_4a2f7560afdd36ef != nil {
					__obf_4a2f7560afdd36ef = fmt.Errorf("create payment client with TLS failed: %w", __obf_4a2f7560afdd36ef)
					return
				}
			}
			__obf_5fe4c1650c09a964 = &Payment{__obf_2d2cf9566604eb22: __obf_4955881acf1d99f7, __obf_9786d25234c49e0b: __obf_9786d25234c49e0b}
		})
	return __obf_5fe4c1650c09a964, __obf_4a2f7560afdd36ef
}

// UnifiedOrder 统一下单
func (__obf_f17b0ca43e47f997 *Payment) UnifiedOrder(__obf_34c0ad174aab2ae8 uint8, __obf_ec98f6d456695344 model.UnifiedOrderRequest) (*model.UnifiedOrderResponse, error) {
	if __obf_ec98f6d456695344.TradeType == "JSAPI" && __obf_ec98f6d456695344.OpenID == "" {
		return nil, errors.New("trade_type is JSAPI but OpenID is not provided")
	}
	__obf_ec98f6d456695344.
		XMLRequest = model.XMLRequest{
		AppID:    __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.AppID,
		MchID:    __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_ec98f6d456695344.
		SignType = __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.SignType
	__obf_ec98f6d456695344. // 使用配置中的签名类型
				NotifyURL = fmt.Sprintf("%s/%d", __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.NotifyURL, __obf_34c0ad174aab2ae8)
	__obf_ec98f6d456695344. // 确保通知URL包含商户号
				XMLRequest.Sign = internal.CreateSign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_ec98f6d456695344.ToMap())
	__obf_2700ed360b8b55fe,

		// 发送XML请求，获取原始响应体
		__obf_4a2f7560afdd36ef := __obf_f17b0ca43e47f997.__obf_9786d25234c49e0b.PostXML(internal.PaymentUnifiedOrderURL, &__obf_ec98f6d456695344)
	if __obf_4a2f7560afdd36ef != nil {
		return nil, fmt.Errorf("unified order request failed: %w", __obf_4a2f7560afdd36ef)
	}
	__obf_15077b2fd5e8d293,

		// 将原始响应体解析为map，用于验签
		__obf_4a2f7560afdd36ef := internal.XMLToMap(__obf_2700ed360b8b55fe)
	if __obf_4a2f7560afdd36ef != nil {
		return nil, fmt.Errorf("parse unified order response XML to map failed: %w", __obf_4a2f7560afdd36ef)
	}
	if __obf_15077b2fd5e8d293["return_code"] == model.FAIL {
		return nil, fmt.Errorf("支付失败: %s", __obf_15077b2fd5e8d293["return_msg"])
	}

	// 验签
	if !internal.VerifySign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_15077b2fd5e8d293) {
		return nil, errors.New("unified order response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体
	var __obf_d976f33745eefcb6 struct {
		model.XMLResponse
		model.UnifiedOrderResponse
	}
	__obf_4a2f7560afdd36ef = // 解析原始响应体为map，供验签使用
		xml.Unmarshal(__obf_2700ed360b8b55fe, &__obf_d976f33745eefcb6)
	if __obf_4a2f7560afdd36ef != nil {
		return nil, fmt.Errorf("unmarshal unified order response XML to struct failed: %w", __obf_4a2f7560afdd36ef)
	}

	// 检查返回码和业务结果码
	if __obf_d976f33745eefcb6.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API return error: %s - %s", __obf_d976f33745eefcb6.ReturnCode, __obf_d976f33745eefcb6.ReturnMsg)
	}
	if __obf_d976f33745eefcb6.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API business error: %s - %s (detail: %s)", __obf_d976f33745eefcb6.ErrCode, __obf_d976f33745eefcb6.ErrCodeDes, __obf_d976f33745eefcb6.ReturnMsg)
	}
	__obf_d976f33745eefcb6.
		Timestamp = fmt.Sprint(time.Now().Unix())
	__obf_d976f33745eefcb6.
		Package = fmt.Sprintf("prepay_id=%s", __obf_d976f33745eefcb6.PrepayID)
	__obf_d976f33745eefcb6.
		Sign = internal.CreateSign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, map[string]string{
		"appId":     __obf_d976f33745eefcb6.AppID,
		"nonceStr":  __obf_d976f33745eefcb6.NonceStr,
		"package":   __obf_d976f33745eefcb6.Package,
		"signType":  __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.SignType,
		"timeStamp": __obf_d976f33745eefcb6.Timestamp,
	})

	return &__obf_d976f33745eefcb6.UnifiedOrderResponse, nil
}

// Refund 发起退款请求 (需要API证书)
// transactionID: 微信订单号 (与outTradeNo二选一)
// outTradeNo: 商户订单号 (与transactionID二选一)
// outRefundNo: 商户退款单号 (唯一)
// totalFee: 订单总金额 (单位：分)
// refundFee: 退款金额 (单位：分)
// refundDesc: 退款原因
func (__obf_f17b0ca43e47f997 *Payment) Refund(__obf_ec98f6d456695344 model.RefundRequest) (string, error) {

	if __obf_ec98f6d456695344.TransactionID == "" && __obf_ec98f6d456695344.OutTradeNo == "" {
		return "", errors.New("transaction_id or out_trade_no must be provided")
	}
	__obf_ec98f6d456695344.
		XMLRequest = model.XMLRequest{
		AppID:    __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.AppID,
		MchID:    __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_ec98f6d456695344.
		SignType = __obf_f17b0ca43e47f997. // 使用配置中的签名类型
		__obf_2d2cf9566604eb22.SignType
	__obf_ec98f6d456695344.
		NotifyURL = __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.NotifyURL
	__obf_ec98f6d456695344.
		XMLRequest.Sign = internal.CreateSign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_ec98f6d456695344.ToMap())
	__obf_2700ed360b8b55fe, // 签名放入请求结构体中
		__obf_4a2f7560afdd36ef := __obf_f17b0ca43e47f997.__obf_9786d25234c49e0b.PostXML(internal.PaymentRefundURL, &__obf_ec98f6d456695344)
	if __obf_4a2f7560afdd36ef != nil {
		return "", fmt.Errorf("refund request failed: %w", __obf_4a2f7560afdd36ef)
	}
	__obf_15077b2fd5e8d293,

		// 将原始响应体解析为map，用于验签
		__obf_4a2f7560afdd36ef := internal.XMLToMap(__obf_2700ed360b8b55fe)
	if __obf_4a2f7560afdd36ef != nil {
		return "", fmt.Errorf("parse refund response XML to map failed: %w", __obf_4a2f7560afdd36ef)
	}

	// 验签
	if !internal.VerifySign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_15077b2fd5e8d293) {
		return "", errors.New("refund response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_d976f33745eefcb6 struct {
		model.XMLResponse
		model.RefundResponse
	}
	__obf_4a2f7560afdd36ef = xml.Unmarshal(__obf_2700ed360b8b55fe, &__obf_d976f33745eefcb6)
	if __obf_4a2f7560afdd36ef != nil {
		return "", fmt.Errorf("unmarshal refund response XML to struct failed: %w", __obf_4a2f7560afdd36ef)
	}

	if __obf_d976f33745eefcb6.ReturnCode != model.SUCCESS {
		return "", fmt.Errorf("refund API return error: %s - %s", __obf_d976f33745eefcb6.ReturnCode, __obf_d976f33745eefcb6.ReturnMsg)
	}
	if __obf_d976f33745eefcb6.ResultCode != model.SUCCESS {
		return "", fmt.Errorf("refund API business error: %s - %s (detail: %s)", __obf_d976f33745eefcb6.ErrCode, __obf_d976f33745eefcb6.ErrCodeDes, __obf_d976f33745eefcb6.ReturnMsg)
	}

	return string(__obf_2700ed360b8b55fe), nil
}

func (__obf_f17b0ca43e47f997 *Payment) PayNotify(__obf_212866f03f511325 []byte, __obf_698b8d8a0386d7f7 func(model.PayNotifyRequest) error) error {

	// 6. 将原始XML解析为结构体 (用于更方便地访问字段)
	var __obf_ec98f6d456695344 model.PayNotifyRequest
	__obf_4a2f7560afdd36ef := xml.Unmarshal(__obf_212866f03f511325, &__obf_ec98f6d456695344)
	if __obf_4a2f7560afdd36ef != nil {
		return fmt.Errorf("Payment Notify: Unmarshal XML to struct failed: %v, Body: %s", __obf_4a2f7560afdd36ef, string(__obf_212866f03f511325))
	}

	// 即使业务失败，如果通信成功 (ReturnCode是SUCCESS)，也应该返回 SUCCESS 给微信，
	// 告诉微信你已经收到通知，避免微信重复发送。
	// 只有在通信本身出错时 (如读取Body失败，验签失败等) 才返回 FAIL

	if __obf_ec98f6d456695344.ReturnCode != model.SUCCESS {
		return fmt.Errorf("Payment Notify: ReturnCode is not SUCCESS: %s - %s", __obf_ec98f6d456695344.ReturnCode, __obf_ec98f6d456695344.ReturnMsg)
	}

	// 7. 检查业务结果 (ResultCode)
	if __obf_ec98f6d456695344.ResultCode == model.SUCCESS {
		// 5. 验签
		// VerifySign 会在内部移除 reqMap 中的 "sign" 字段，所以传入前需要复制或注意。
		// 这里直接传入 reqMap 是可以的，因为验签后我们通常不再使用这个 map 进行后续操作。
		if !internal.VerifySign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_ec98f6d456695344.ToMap()) {
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
		return __obf_698b8d8a0386d7f7(__obf_ec98f6d456695344)
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
func ReturnMessage(__obf_c1982d7a12982645, __obf_006572066fc766f0 string) []byte {
	__obf_15077b2fd5e8d293 := map[string]string{
		"return_code": __obf_c1982d7a12982645,
		"return_msg":  __obf_006572066fc766f0,
	}
	return internal.MapToXML(__obf_15077b2fd5e8d293)
}

func (__obf_f17b0ca43e47f997 *Payment) TransBank(__obf_ec98f6d456695344 model.TransBankRequest) (*model.TransBankResponse, error) {

	if __obf_ec98f6d456695344.PartnerTradeNo == "" && __obf_ec98f6d456695344.EncBankNo == "" && __obf_ec98f6d456695344.EncTrueName == "" && __obf_ec98f6d456695344.BankCode == "" {
		return nil, errors.New("请提供商户付款单号、收款方银行卡号、收款方用户名、收款方开户行")
	}
	__obf_ec98f6d456695344.
		XMLRequest = model.XMLRequest{
		AppID:    __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.AppID,
		MchID:    __obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_ec98f6d456695344.
		XMLRequest.Sign = internal.CreateSign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_ec98f6d456695344.ToMap())
	__obf_2700ed360b8b55fe, // 签名放入请求结构体中
		__obf_4a2f7560afdd36ef := __obf_f17b0ca43e47f997.__obf_9786d25234c49e0b.PostXML(internal.PaymentTransURL, &__obf_ec98f6d456695344)
	if __obf_4a2f7560afdd36ef != nil {
		return nil, fmt.Errorf("pay_bank request failed: %w", __obf_4a2f7560afdd36ef)
	}
	__obf_15077b2fd5e8d293,

		// 将原始响应体解析为map，用于验签
		__obf_4a2f7560afdd36ef := internal.XMLToMap(__obf_2700ed360b8b55fe)
	if __obf_4a2f7560afdd36ef != nil {
		return nil, fmt.Errorf("parse pay_bank response XML to map failed: %w", __obf_4a2f7560afdd36ef)
	}

	// 验签
	if !internal.VerifySign(__obf_f17b0ca43e47f997.__obf_2d2cf9566604eb22.APIKey, __obf_15077b2fd5e8d293) {
		return nil, errors.New("pay_bank response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_d976f33745eefcb6 struct {
		model.XMLResponse
		model.TransBankResponse
	}
	__obf_4a2f7560afdd36ef = xml.Unmarshal(__obf_2700ed360b8b55fe, &__obf_d976f33745eefcb6)
	if __obf_4a2f7560afdd36ef != nil {
		return nil, fmt.Errorf("unmarshal pay_bank response XML to struct failed: %w", __obf_4a2f7560afdd36ef)
	}

	if __obf_d976f33745eefcb6.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API return error: %s - %s", __obf_d976f33745eefcb6.ReturnCode, __obf_d976f33745eefcb6.ReturnMsg)
	}
	if __obf_d976f33745eefcb6.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API business error: %s - %s (detail: %s)", __obf_d976f33745eefcb6.ErrCode, __obf_d976f33745eefcb6.ErrCodeDes, __obf_d976f33745eefcb6.ReturnMsg)
	}

	return &__obf_d976f33745eefcb6.TransBankResponse, nil
}
