package __obf_819ce2f3b265eaaa

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // PayNotifyRequest 微信支付结果通知请求参数
// type PayNotifyRequest struct {
// 	XMLName       xml.Name `xml:"xml"`
// 	ReturnCode    string   `xml:"return_code"`   // 返回状态码
// 	ReturnMsg     string   `xml:"return_msg"`    // 返回信息
// 	AppID         string   `xml:"appid"`         // 公众账号ID
// 	MchID         string   `xml:"mch_id"`        // 商户号
// 	NonceStr      string   `xml:"nonce_str"`     // 随机字符串
// 	Sign          string   `xml:"sign"`          // 签名
// 	ResultCode    string   `xml:"result_code"`   // 业务结果
// 	ErrCode       string   `xml:"err_code"`      // 错误代码
// 	ErrCodeDes    string   `xml:"err_code_des"`  // 错误代码描述
// 	OpenID        string   `xml:"openid"`        // 用户标识
// 	IsSubscribe   string   `xml:"is_subscribe"`  // 是否关注公众账号
// 	TradeType     string   `xml:"trade_type"`    // 交易类型
// 	BankType      string   `xml:"bank_type"`     // 付款银行
// 	TotalFee      int      `xml:"total_fee"`     // 订单总金额，单位分
// 	CashFee       int      `xml:"cash_fee"`      // 现金支付金额，单位分
// 	TransactionID string   `xml:"transaction_id"`// 微信支付订单号
// 	OutTradeNo    string   `xml:"out_trade_no"`  // 商户系统内部订单号
// 	Attach        string   `xml:"attach"`        // 商家数据包
// 	TimeEnd       string   `xml:"time_end"`      // 支付完成时间，格式为yyyyMMddHHmmss
// 	// ... 其他可能返回的字段
// }

// // PayNotifyResponse 微信支付结果通知响应给微信的参数
// type PayNotifyResponse struct {
// 	XMLName    xml.Name `xml:"xml"`
// 	ReturnCode string   `xml:"return_code"` // 成功为 SUCCESS
// 	ReturnMsg  string   `xml:"return_msg"`  // 例如 OK
// }

// // HandlePaymentNotify 是一个处理微信支付回调的HTTP处理器函数。
// // 通常这个函数会被注册到一个HTTP路由上，例如 `/wechat/pay/notify`。
// // payInstance 是支付模块的实例，用于获取配置和APIKey进行验签。
// func HandlePaymentNotify(w http.ResponseWriter, r *http.Request, payInstance *Payment) {
// 	// 1. 读取请求体
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		log.Printf("Payment Notify: Read request body failed: %v", err)
// 		writeNotifyResponse(w, "FAIL", "Invalid Body")
// 		return
// 	}
// 	defer r.Body.Close()

// 	// 2. 将XML解析为map
// 	reqMap, err := XMLToMap(body)
// 	if err != nil {
// 		log.Printf("Payment Notify: Parse XML failed: %v, Body: %s", err, string(body))
// 		writeNotifyResponse(w, "FAIL", "XML Parse Error")
// 		return
// 	}

// 	// 3. 验签
// 	if !VerifySign(reqMap, payInstance.Config.APIKey) {
// 		log.Printf("Payment Notify: Signature verification failed. Request Map: %+v", reqMap)
// 		writeNotifyResponse(w, "FAIL", "Signature Error")
// 		return
// 	}

// 	// 4. 解析为结构体 (可选，也可以直接用map处理)
// 	var notifyReq PayNotifyRequest
// 	err = xml.Unmarshal(body, ¬ifyReq)
// 	if err != nil {
// 		log.Printf("Payment Notify: Unmarshal XML to struct failed: %v, Body: %s", err, string(body))
// 		writeNotifyResponse(w, "FAIL", "XML Unmarshal Error")
// 		return
// 	}

// 	// 5. 检查业务结果
// 	if notifyReq.ReturnCode == "SUCCESS" && notifyReq.ResultCode == "SUCCESS" {
// 		log.Printf("Payment Notify Success for Order: %s, TransactionID: %s, TotalFee: %d",
// 			notifyReq.OutTradeNo, notifyReq.TransactionID, notifyReq.TotalFee)

// 		// TODO: 在这里处理你的业务逻辑
// 		// 例如：
// 		// - 根据 notifyReq.OutTradeNo 查询订单
// 		// - 验证 notifyReq.TotalFee 是否与你订单的金额一致
// 		// - 更新订单状态为已支付
// 		// - 处理商品发货、服务开通等后续操作
// 		// - 保证幂等性：避免重复处理已支付的订单

// 		// 假设业务处理成功
// 		writeNotifyResponse(w, "SUCCESS", "OK")
// 	} else {
// 		// 支付失败或其他返回码
// 		log.Printf("Payment Notify Failed. ReturnCode: %s, ReturnMsg: %s, ResultCode: %s, ErrCode: %s, ErrCodeDes: %s, OutTradeNo: %s",
// 			notifyReq.ReturnCode, notifyReq.ReturnMsg, notifyReq.ResultCode, notifyReq.ErrCode, notifyReq.ErrCodeDes, notifyReq.OutTradeNo)
// 		// 通常失败通知不需要业务处理，但仍需返回成功给微信，避免重复通知。
// 		// 如果是ReturnCode为FAIL，表示通信失败，可能需要让微信重试
// 		if notifyReq.ReturnCode == "FAIL" {
// 			writeNotifyResponse(w, "FAIL", notifyReq.ReturnMsg) // 让微信重试
// 		} else {
// 			writeNotifyResponse(w, "SUCCESS", "OK") // 业务失败，通知微信已收到
// 		}
// 	}
// }

// // writeNotifyResponse 向微信返回响应
// func writeNotifyResponse(w http.ResponseWriter, returnCode, returnMsg string) {
// 	resp := PayNotifyResponse{
// 		ReturnCode: returnCode,
// 		ReturnMsg:  returnMsg,
// 	}
// 	respBytes, err := xml.Marshal(resp)
// 	if err != nil {
// 		log.Printf("Payment Notify: Marshal response XML failed: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/xml")
// 	w.Write(respBytes)
// }
