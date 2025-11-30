package __obf_b5bcac367b722f8a

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
func PKCS7UnPadding(__obf_6e67bb94b5e1eda9 []byte) []byte {
	__obf_8ff1fddf399d9d1a := len(__obf_6e67bb94b5e1eda9)
	__obf_3c7dec4d404e555d := int(__obf_6e67bb94b5e1eda9[__obf_8ff1fddf399d9d1a-1])
	return __obf_6e67bb94b5e1eda9[:(__obf_8ff1fddf399d9d1a - __obf_3c7dec4d404e555d)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_b41cbe95e885c9d6(__obf_eb194fea2e30bbb4, __obf_bdc3e8ed12071b9c, __obf_fbb6499f31e56b06 []byte) ([]byte, error) {
	__obf_850a472cd2c60962, __obf_d9b19f08bff09faf := aes.NewCipher(__obf_eb194fea2e30bbb4)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_c649c9bdaba503d4 := __obf_850a472cd2c60962.BlockSize()
	if len(__obf_bdc3e8ed12071b9c) != __obf_c649c9bdaba503d4 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_c649c9bdaba503d4)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_fbb6499f31e56b06)%__obf_c649c9bdaba503d4 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	__obf_39b8bd4458c3a28b := cipher.NewCBCDecrypter(__obf_850a472cd2c60962, __obf_bdc3e8ed12071b9c)
	__obf_1932b5428dd9ce45 := make([]byte, len(__obf_fbb6499f31e56b06))
	__obf_39b8bd4458c3a28b.
		CryptBlocks(__obf_1932b5428dd9ce45, __obf_fbb6499f31e56b06)
	__obf_1932b5428dd9ce45 = PKCS7UnPadding(__obf_1932b5428dd9ce45)

	return __obf_1932b5428dd9ce45, nil
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
func DecryptWXData(__obf_3c11a9b6f8e01ec7, __obf_0fe591c71c6c8f52, __obf_bdc3e8ed12071b9c, __obf_be8acd839654090d string) (*DecryptedData, error) {
	__obf_f971d79eb18762ed, __obf_d9b19f08bff09faf := base64.StdEncoding.DecodeString(__obf_3c11a9b6f8e01ec7)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_24f30a01b3736152, __obf_d9b19f08bff09faf := base64.StdEncoding.DecodeString(__obf_0fe591c71c6c8f52)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_d9b19f08bff09faf)
	}
	IV, __obf_d9b19f08bff09faf := base64.StdEncoding.DecodeString(__obf_bdc3e8ed12071b9c)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_d9b19f08bff09faf)
	}
	__obf_068c126df94fdd8d,

		// 调用解密函数
		__obf_d9b19f08bff09faf := __obf_b41cbe95e885c9d6(__obf_f971d79eb18762ed, IV, __obf_24f30a01b3736152)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_d9b19f08bff09faf)
	}

	var __obf_814ec131f39316bf DecryptedData
	__obf_d9b19f08bff09faf = json.Unmarshal(__obf_068c126df94fdd8d, &__obf_814ec131f39316bf)
	if __obf_d9b19f08bff09faf != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_d9b19f08bff09faf)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_814ec131f39316bf.Watermark.AppID != __obf_be8acd839654090d {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_814ec131f39316bf, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_2e353653f52b4b5b []byte, __obf_c649c9bdaba503d4 int) []byte {
	__obf_680c147ce2ed092d := __obf_c649c9bdaba503d4 - len(__obf_2e353653f52b4b5b)%__obf_c649c9bdaba503d4
	__obf_a590cee5d8a7c3a5 := bytes.Repeat([]byte{byte(__obf_680c147ce2ed092d)}, __obf_680c147ce2ed092d)
	return append(__obf_2e353653f52b4b5b, __obf_a590cee5d8a7c3a5...)
}
