package __obf_6e18fdc479ab921c

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
func PKCS7UnPadding(__obf_1684cc882f627a00 []byte) []byte {
	__obf_a58eb5d4b5a8098e := len(__obf_1684cc882f627a00)
	__obf_691ad357582ff2ee := int(__obf_1684cc882f627a00[__obf_a58eb5d4b5a8098e-1])
	return __obf_1684cc882f627a00[:(__obf_a58eb5d4b5a8098e - __obf_691ad357582ff2ee)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_a469ee5ab90f4576(__obf_57cc787ea702af69, __obf_1ffafd9c167cbb7d, __obf_00a09b30fdf88425 []byte) ([]byte, error) {
	__obf_237cbb22504fcae4, __obf_d92b02392b1fb408 := aes.NewCipher(__obf_57cc787ea702af69)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_d92b02392b1fb408)
	}

	__obf_43a62997749e7805 := __obf_237cbb22504fcae4.BlockSize()
	if len(__obf_1ffafd9c167cbb7d) != __obf_43a62997749e7805 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_43a62997749e7805)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_00a09b30fdf88425)%__obf_43a62997749e7805 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	__obf_c5c9bff448d6231f := cipher.NewCBCDecrypter(__obf_237cbb22504fcae4, __obf_1ffafd9c167cbb7d)
	__obf_a9fc807ce1dfd8d9 := make([]byte, len(__obf_00a09b30fdf88425))
	__obf_c5c9bff448d6231f.CryptBlocks(__obf_a9fc807ce1dfd8d9, __obf_00a09b30fdf88425)

	__obf_a9fc807ce1dfd8d9 = PKCS7UnPadding(__obf_a9fc807ce1dfd8d9)

	return __obf_a9fc807ce1dfd8d9, nil
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
func DecryptWXData(__obf_d9e99bfa46289b42, __obf_ed8d8abe825b806b, __obf_1ffafd9c167cbb7d, __obf_ee54cc74f1026cd3 string) (*DecryptedData, error) {
	__obf_f185e11e74919a1f, __obf_d92b02392b1fb408 := base64.StdEncoding.DecodeString(__obf_d9e99bfa46289b42)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_d92b02392b1fb408)
	}
	__obf_706f478ce351da69, __obf_d92b02392b1fb408 := base64.StdEncoding.DecodeString(__obf_ed8d8abe825b806b)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_d92b02392b1fb408)
	}
	IV, __obf_d92b02392b1fb408 := base64.StdEncoding.DecodeString(__obf_1ffafd9c167cbb7d)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_d92b02392b1fb408)
	}

	// 调用解密函数
	__obf_c73af6f65475079f, __obf_d92b02392b1fb408 := __obf_a469ee5ab90f4576(__obf_f185e11e74919a1f, IV, __obf_706f478ce351da69)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_d92b02392b1fb408)
	}

	var __obf_a371d1762a30244f DecryptedData
	__obf_d92b02392b1fb408 = json.Unmarshal(__obf_c73af6f65475079f, &__obf_a371d1762a30244f)
	if __obf_d92b02392b1fb408 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_d92b02392b1fb408)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_a371d1762a30244f.Watermark.AppID != __obf_ee54cc74f1026cd3 {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_a371d1762a30244f, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_f690c0713fc53061 []byte, __obf_43a62997749e7805 int) []byte {
	__obf_42242053aaecf140 := __obf_43a62997749e7805 - len(__obf_f690c0713fc53061)%__obf_43a62997749e7805
	__obf_9be8fe0f6b046850 := bytes.Repeat([]byte{byte(__obf_42242053aaecf140)}, __obf_42242053aaecf140)
	return append(__obf_f690c0713fc53061, __obf_9be8fe0f6b046850...)
}
