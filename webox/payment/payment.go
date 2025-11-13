package __obf_819ce2f3b265eaaa

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
	__obf_ce7faf1fa8664d06 *Config
	__obf_b663a22cdfdd82bd *internal.HttpClient
}

var (
	__obf_b287ae543c205e6c *Payment
	__obf_9f10ca794f047568 sync.Once
)

// NewPayment 创建并返回支付模块的单例实例
// 根据配置中的 CertPath 和 KeyPath 决定是否创建带证书的客户端
func NewPayment(__obf_1037853c0c74ede5 *Config) (*Payment, error) {
	var __obf_b554e824306c499a error
	__obf_9f10ca794f047568.Do(func() {
		var __obf_b663a22cdfdd82bd *internal.HttpClient
		if len(__obf_1037853c0c74ede5.TLSCert.Certificate) == 0 || __obf_1037853c0c74ede5.TLSCert.PrivateKey == nil {
			__obf_b663a22cdfdd82bd = internal.NewClient()
		} else {
			__obf_b663a22cdfdd82bd, __obf_b554e824306c499a = internal.NewClientWithTLS(__obf_1037853c0c74ede5.TLSCert)
			if __obf_b554e824306c499a != nil {
				__obf_b554e824306c499a = fmt.Errorf("create payment client with TLS failed: %w", __obf_b554e824306c499a)
				return
			}
		}
		__obf_b287ae543c205e6c = &Payment{
			__obf_ce7faf1fa8664d06: __obf_1037853c0c74ede5,
			__obf_b663a22cdfdd82bd: __obf_b663a22cdfdd82bd,
		}
	})
	return __obf_b287ae543c205e6c, __obf_b554e824306c499a
}

// UnifiedOrder 统一下单
func (__obf_c41ee8d1f15e7935 *Payment) UnifiedOrder(__obf_756f1007caf8bcfb uint8, __obf_3aa821703573a064 model.UnifiedOrderRequest) (*model.UnifiedOrderResponse, error) {
	if __obf_3aa821703573a064.TradeType == "JSAPI" && __obf_3aa821703573a064.OpenID == "" {
		return nil, errors.New("trade_type is JSAPI but OpenID is not provided")
	}
	__obf_3aa821703573a064.XMLRequest = model.XMLRequest{
		AppID:    __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.AppID,
		MchID:    __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_3aa821703573a064.SignType = __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.SignType                                                 // 使用配置中的签名类型
	__obf_3aa821703573a064.NotifyURL = fmt.Sprintf("%s/%d", __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.NotifyURL, __obf_756f1007caf8bcfb) // 确保通知URL包含商户号
	__obf_3aa821703573a064.XMLRequest.Sign = internal.CreateSign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_3aa821703573a064.ToMap())

	// 发送XML请求，获取原始响应体
	__obf_fcd98faaf4c4f1c5, __obf_b554e824306c499a := __obf_c41ee8d1f15e7935.__obf_b663a22cdfdd82bd.PostXML(internal.PaymentUnifiedOrderURL, &__obf_3aa821703573a064)
	if __obf_b554e824306c499a != nil {
		return nil, fmt.Errorf("unified order request failed: %w", __obf_b554e824306c499a)
	}

	// 将原始响应体解析为map，用于验签
	__obf_b5df4e88652c4f07, __obf_b554e824306c499a := internal.XMLToMap(__obf_fcd98faaf4c4f1c5)
	if __obf_b554e824306c499a != nil {
		return nil, fmt.Errorf("parse unified order response XML to map failed: %w", __obf_b554e824306c499a)
	}
	if __obf_b5df4e88652c4f07["return_code"] == model.FAIL {
		return nil, fmt.Errorf("支付失败: %s", __obf_b5df4e88652c4f07["return_msg"])
	}

	// 验签
	if !internal.VerifySign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_b5df4e88652c4f07) {
		return nil, errors.New("unified order response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体
	var __obf_962b232fe731106e struct {
		model.XMLResponse
		model.UnifiedOrderResponse
	}
	// 解析原始响应体为map，供验签使用
	__obf_b554e824306c499a = xml.Unmarshal(__obf_fcd98faaf4c4f1c5, &__obf_962b232fe731106e)
	if __obf_b554e824306c499a != nil {
		return nil, fmt.Errorf("unmarshal unified order response XML to struct failed: %w", __obf_b554e824306c499a)
	}

	// 检查返回码和业务结果码
	if __obf_962b232fe731106e.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API return error: %s - %s", __obf_962b232fe731106e.ReturnCode, __obf_962b232fe731106e.ReturnMsg)
	}
	if __obf_962b232fe731106e.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("unified order API business error: %s - %s (detail: %s)", __obf_962b232fe731106e.ErrCode, __obf_962b232fe731106e.ErrCodeDes, __obf_962b232fe731106e.ReturnMsg)
	}
	__obf_962b232fe731106e.Timestamp = fmt.Sprint(time.Now().Unix())
	__obf_962b232fe731106e.Package = fmt.Sprintf("prepay_id=%s", __obf_962b232fe731106e.PrepayID)
	__obf_962b232fe731106e.Sign = internal.CreateSign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, map[string]string{
		"appId":     __obf_962b232fe731106e.AppID,
		"nonceStr":  __obf_962b232fe731106e.NonceStr,
		"package":   __obf_962b232fe731106e.Package,
		"signType":  __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.SignType,
		"timeStamp": __obf_962b232fe731106e.Timestamp,
	})

	return &__obf_962b232fe731106e.UnifiedOrderResponse, nil
}

// Refund 发起退款请求 (需要API证书)
// transactionID: 微信订单号 (与outTradeNo二选一)
// outTradeNo: 商户订单号 (与transactionID二选一)
// outRefundNo: 商户退款单号 (唯一)
// totalFee: 订单总金额 (单位：分)
// refundFee: 退款金额 (单位：分)
// refundDesc: 退款原因
func (__obf_c41ee8d1f15e7935 *Payment) Refund(__obf_3aa821703573a064 model.RefundRequest) (string, error) {

	if __obf_3aa821703573a064.TransactionID == "" && __obf_3aa821703573a064.OutTradeNo == "" {
		return "", errors.New("transaction_id or out_trade_no must be provided")
	}

	__obf_3aa821703573a064.XMLRequest = model.XMLRequest{
		AppID:    __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.AppID,
		MchID:    __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_3aa821703573a064.SignType = __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.SignType // 使用配置中的签名类型
	__obf_3aa821703573a064.NotifyURL = __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.NotifyURL
	__obf_3aa821703573a064.XMLRequest.Sign = internal.CreateSign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_3aa821703573a064.ToMap()) // 签名放入请求结构体中

	__obf_fcd98faaf4c4f1c5, __obf_b554e824306c499a := __obf_c41ee8d1f15e7935.__obf_b663a22cdfdd82bd.PostXML(internal.PaymentRefundURL, &__obf_3aa821703573a064)
	if __obf_b554e824306c499a != nil {
		return "", fmt.Errorf("refund request failed: %w", __obf_b554e824306c499a)
	}

	// 将原始响应体解析为map，用于验签
	__obf_b5df4e88652c4f07, __obf_b554e824306c499a := internal.XMLToMap(__obf_fcd98faaf4c4f1c5)
	if __obf_b554e824306c499a != nil {
		return "", fmt.Errorf("parse refund response XML to map failed: %w", __obf_b554e824306c499a)
	}

	// 验签
	if !internal.VerifySign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_b5df4e88652c4f07) {
		return "", errors.New("refund response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_962b232fe731106e struct {
		model.XMLResponse
		model.RefundResponse
	}
	__obf_b554e824306c499a = xml.Unmarshal(__obf_fcd98faaf4c4f1c5, &__obf_962b232fe731106e)
	if __obf_b554e824306c499a != nil {
		return "", fmt.Errorf("unmarshal refund response XML to struct failed: %w", __obf_b554e824306c499a)
	}

	if __obf_962b232fe731106e.ReturnCode != model.SUCCESS {
		return "", fmt.Errorf("refund API return error: %s - %s", __obf_962b232fe731106e.ReturnCode, __obf_962b232fe731106e.ReturnMsg)
	}
	if __obf_962b232fe731106e.ResultCode != model.SUCCESS {
		return "", fmt.Errorf("refund API business error: %s - %s (detail: %s)", __obf_962b232fe731106e.ErrCode, __obf_962b232fe731106e.ErrCodeDes, __obf_962b232fe731106e.ReturnMsg)
	}

	return string(__obf_fcd98faaf4c4f1c5), nil
}

func (__obf_c41ee8d1f15e7935 *Payment) PayNotify(__obf_6828bc69153a6ac0 []byte, __obf_442926e44d48bbc6 func(model.PayNotifyRequest) error) error {

	// 6. 将原始XML解析为结构体 (用于更方便地访问字段)
	var __obf_3aa821703573a064 model.PayNotifyRequest
	__obf_b554e824306c499a := xml.Unmarshal(__obf_6828bc69153a6ac0, &__obf_3aa821703573a064)
	if __obf_b554e824306c499a != nil {
		return fmt.Errorf("Payment Notify: Unmarshal XML to struct failed: %v, Body: %s", __obf_b554e824306c499a, string(__obf_6828bc69153a6ac0))
	}

	// 即使业务失败，如果通信成功 (ReturnCode是SUCCESS)，也应该返回 SUCCESS 给微信，
	// 告诉微信你已经收到通知，避免微信重复发送。
	// 只有在通信本身出错时 (如读取Body失败，验签失败等) 才返回 FAIL

	if __obf_3aa821703573a064.ReturnCode != model.SUCCESS {
		return fmt.Errorf("Payment Notify: ReturnCode is not SUCCESS: %s - %s", __obf_3aa821703573a064.ReturnCode, __obf_3aa821703573a064.ReturnMsg)
	}

	// 7. 检查业务结果 (ResultCode)
	if __obf_3aa821703573a064.ResultCode == model.SUCCESS {
		// 5. 验签
		// VerifySign 会在内部移除 reqMap 中的 "sign" 字段，所以传入前需要复制或注意。
		// 这里直接传入 reqMap 是可以的，因为验签后我们通常不再使用这个 map 进行后续操作。
		if !internal.VerifySign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_3aa821703573a064.ToMap()) {
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
		return __obf_442926e44d48bbc6(__obf_3aa821703573a064)
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
func ReturnMessage(__obf_b8bdd9a3be7e14dd, __obf_fb49adbb93a92b51 string) []byte {
	__obf_b5df4e88652c4f07 := map[string]string{
		"return_code": __obf_b8bdd9a3be7e14dd,
		"return_msg":  __obf_fb49adbb93a92b51,
	}
	return internal.MapToXML(__obf_b5df4e88652c4f07)
}

func (__obf_c41ee8d1f15e7935 *Payment) TransBank(__obf_3aa821703573a064 model.TransBankRequest) (*model.TransBankResponse, error) {

	if __obf_3aa821703573a064.PartnerTradeNo == "" && __obf_3aa821703573a064.EncBankNo == "" && __obf_3aa821703573a064.EncTrueName == "" && __obf_3aa821703573a064.BankCode == "" {
		return nil, errors.New("请提供商户付款单号、收款方银行卡号、收款方用户名、收款方开户行")
	}

	__obf_3aa821703573a064.XMLRequest = model.XMLRequest{
		AppID:    __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.AppID,
		MchID:    __obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.MchID,
		NonceStr: internal.GenerateNonceStr(),
	}
	__obf_3aa821703573a064.XMLRequest.Sign = internal.CreateSign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_3aa821703573a064.ToMap()) // 签名放入请求结构体中

	__obf_fcd98faaf4c4f1c5, __obf_b554e824306c499a := __obf_c41ee8d1f15e7935.__obf_b663a22cdfdd82bd.PostXML(internal.PaymentTransURL, &__obf_3aa821703573a064)
	if __obf_b554e824306c499a != nil {
		return nil, fmt.Errorf("pay_bank request failed: %w", __obf_b554e824306c499a)
	}

	// 将原始响应体解析为map，用于验签
	__obf_b5df4e88652c4f07, __obf_b554e824306c499a := internal.XMLToMap(__obf_fcd98faaf4c4f1c5)
	if __obf_b554e824306c499a != nil {
		return nil, fmt.Errorf("parse pay_bank response XML to map failed: %w", __obf_b554e824306c499a)
	}

	// 验签
	if !internal.VerifySign(__obf_c41ee8d1f15e7935.__obf_ce7faf1fa8664d06.APIKey, __obf_b5df4e88652c4f07) {
		return nil, errors.New("pay_bank response signature verification failed")
	}

	// 验签通过后，将原始响应体解析为结构体

	var __obf_962b232fe731106e struct {
		model.XMLResponse
		model.TransBankResponse
	}
	__obf_b554e824306c499a = xml.Unmarshal(__obf_fcd98faaf4c4f1c5, &__obf_962b232fe731106e)
	if __obf_b554e824306c499a != nil {
		return nil, fmt.Errorf("unmarshal pay_bank response XML to struct failed: %w", __obf_b554e824306c499a)
	}

	if __obf_962b232fe731106e.ReturnCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API return error: %s - %s", __obf_962b232fe731106e.ReturnCode, __obf_962b232fe731106e.ReturnMsg)
	}
	if __obf_962b232fe731106e.ResultCode != model.SUCCESS {
		return nil, fmt.Errorf("pay_bank API business error: %s - %s (detail: %s)", __obf_962b232fe731106e.ErrCode, __obf_962b232fe731106e.ErrCodeDes, __obf_962b232fe731106e.ReturnMsg)
	}

	return &__obf_962b232fe731106e.TransBankResponse, nil
}
