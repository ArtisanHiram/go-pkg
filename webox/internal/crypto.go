package __obf_2e4735ec379515a2

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
func PKCS7UnPadding(__obf_957991dbf92ca323 []byte) []byte {
	__obf_9dc9e70a2a87d3aa := len(__obf_957991dbf92ca323)
	__obf_2edb4c2ba1440fa9 := int(__obf_957991dbf92ca323[__obf_9dc9e70a2a87d3aa-1])
	return __obf_957991dbf92ca323[:(__obf_9dc9e70a2a87d3aa - __obf_2edb4c2ba1440fa9)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_449cad122bd223dc(__obf_3e0be123922aa44c, __obf_a8eb078891e59567, __obf_2b66cf8707765adb []byte) ([]byte, error) {
	__obf_50f635d7a6a89ae3, __obf_647657e89e0327f9 := aes.NewCipher(__obf_3e0be123922aa44c)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_647657e89e0327f9)
	}

	__obf_ebb9a89e31128025 := __obf_50f635d7a6a89ae3.BlockSize()
	if len(__obf_a8eb078891e59567) != __obf_ebb9a89e31128025 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_ebb9a89e31128025)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_2b66cf8707765adb)%__obf_ebb9a89e31128025 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	__obf_0af340685f35dbb6 := cipher.NewCBCDecrypter(__obf_50f635d7a6a89ae3, __obf_a8eb078891e59567)
	__obf_6994d8a891c67423 := make([]byte, len(__obf_2b66cf8707765adb))
	__obf_0af340685f35dbb6.CryptBlocks(__obf_6994d8a891c67423, __obf_2b66cf8707765adb)

	__obf_6994d8a891c67423 = PKCS7UnPadding(__obf_6994d8a891c67423)

	return __obf_6994d8a891c67423, nil
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
func DecryptWXData(__obf_8d34601e5baf4142, __obf_21165a24c61b546c, __obf_a8eb078891e59567, __obf_39c705b664062fa4 string) (*DecryptedData, error) {
	__obf_42bbc077fa84cd66, __obf_647657e89e0327f9 := base64.StdEncoding.DecodeString(__obf_8d34601e5baf4142)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_647657e89e0327f9)
	}
	__obf_91d6bc00fb7c435c, __obf_647657e89e0327f9 := base64.StdEncoding.DecodeString(__obf_21165a24c61b546c)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_647657e89e0327f9)
	}
	IV, __obf_647657e89e0327f9 := base64.StdEncoding.DecodeString(__obf_a8eb078891e59567)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_647657e89e0327f9)
	}

	// 调用解密函数
	__obf_e7d2394fab304353, __obf_647657e89e0327f9 := __obf_449cad122bd223dc(__obf_42bbc077fa84cd66, IV, __obf_91d6bc00fb7c435c)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_647657e89e0327f9)
	}

	var __obf_b7a057dcb85221de DecryptedData
	__obf_647657e89e0327f9 = json.Unmarshal(__obf_e7d2394fab304353, &__obf_b7a057dcb85221de)
	if __obf_647657e89e0327f9 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_647657e89e0327f9)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_b7a057dcb85221de.Watermark.AppID != __obf_39c705b664062fa4 {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_b7a057dcb85221de, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_182f2879ef35efd3 []byte, __obf_ebb9a89e31128025 int) []byte {
	__obf_7d2b32b4f8b09b1e := __obf_ebb9a89e31128025 - len(__obf_182f2879ef35efd3)%__obf_ebb9a89e31128025
	__obf_54c706dae9390469 := bytes.Repeat([]byte{byte(__obf_7d2b32b4f8b09b1e)}, __obf_7d2b32b4f8b09b1e)
	return append(__obf_182f2879ef35efd3, __obf_54c706dae9390469...)
}
