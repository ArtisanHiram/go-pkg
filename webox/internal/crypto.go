package __obf_e79e1b0b12d02d0b

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

// PKCS7UnPadding 对数据进行PKCS7 unpadding
func PKCS7UnPadding(__obf_33aa7f9c2f85db51 []byte) []byte {
	__obf_049fadfb61efc62d := len(__obf_33aa7f9c2f85db51)
	__obf_20e5de2f9a1b103c := int(__obf_33aa7f9c2f85db51[__obf_049fadfb61efc62d-1])
	return __obf_33aa7f9c2f85db51[:(__obf_049fadfb61efc62d - __obf_20e5de2f9a1b103c)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_cab19e5f923709f6(__obf_7dca2452dd625e20, __obf_5efde0b6b0fdb3bf, __obf_ce83151b9faba101 []byte) ([]byte, error) {
	__obf_78c51e8a1f2ea866, __obf_d6bcde9761346c61 := aes.NewCipher(__obf_7dca2452dd625e20)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_54456233e1e810b5 := __obf_78c51e8a1f2ea866.BlockSize()
	if len(__obf_5efde0b6b0fdb3bf) != __obf_54456233e1e810b5 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_54456233e1e810b5)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_ce83151b9faba101)%__obf_54456233e1e810b5 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	__obf_d1c1856f65d2b413 := cipher.NewCBCDecrypter(__obf_78c51e8a1f2ea866, __obf_5efde0b6b0fdb3bf)
	__obf_9724283c5275c86f := make([]byte, len(__obf_ce83151b9faba101))
	__obf_d1c1856f65d2b413.
		CryptBlocks(__obf_9724283c5275c86f, __obf_ce83151b9faba101)
	__obf_9724283c5275c86f = PKCS7UnPadding(__obf_9724283c5275c86f)

	return __obf_9724283c5275c86f, nil
}

// DecryptedData 微信加密数据解密后通用结构体
type DecryptedData struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	Watermark       struct {
		AppID     string `json:"appid"`
		Timestamp int64  `json:"timestamp"`
	} `json:"watermark"`
}

// DecryptWXData 解密微信敏感数据
// sessionKey: 用户登录凭证获取的session_key
// encryptedData: 微信加密数据
// iv: 加密算法的初始向量
// appId: 小程序或公众号的AppID (用于验证水印)
func DecryptWXData(__obf_fc9b95b3a345670c, __obf_2891de108da26cfb, __obf_5efde0b6b0fdb3bf, __obf_dfc8020f36b460b6 string) (*DecryptedData, error) {
	__obf_5acbf3d8a3ba0df9, __obf_d6bcde9761346c61 := base64.StdEncoding.DecodeString(__obf_fc9b95b3a345670c)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_b8a3f7b79b7752b6, __obf_d6bcde9761346c61 := base64.StdEncoding.DecodeString(__obf_2891de108da26cfb)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_d6bcde9761346c61)
	}
	IV, __obf_d6bcde9761346c61 := base64.StdEncoding.DecodeString(__obf_5efde0b6b0fdb3bf)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_d6bcde9761346c61)
	}
	__obf_7b7020626e455cbf,

		// 调用解密函数
		__obf_d6bcde9761346c61 := __obf_cab19e5f923709f6(__obf_5acbf3d8a3ba0df9, IV, __obf_b8a3f7b79b7752b6)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_d6bcde9761346c61)
	}

	var __obf_e8c254e95f1ffc9e DecryptedData
	__obf_d6bcde9761346c61 = json.Unmarshal(__obf_7b7020626e455cbf, &__obf_e8c254e95f1ffc9e)
	if __obf_d6bcde9761346c61 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_d6bcde9761346c61)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_e8c254e95f1ffc9e.Watermark.AppID != __obf_dfc8020f36b460b6 {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_e8c254e95f1ffc9e, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_a7ede658d77da976 []byte, __obf_54456233e1e810b5 int) []byte {
	__obf_bae42abe81aceb0b := __obf_54456233e1e810b5 - len(__obf_a7ede658d77da976)%__obf_54456233e1e810b5
	__obf_0cd0df1e94d3065a := bytes.Repeat([]byte{byte(__obf_bae42abe81aceb0b)}, __obf_bae42abe81aceb0b)
	return append(__obf_a7ede658d77da976, __obf_0cd0df1e94d3065a...)
}
