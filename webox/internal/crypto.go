package __obf_abe10294567bade2

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
func PKCS7UnPadding(__obf_da82f93e176f9f47 []byte) []byte {
	__obf_c4d3ca237ede8853 := len(__obf_da82f93e176f9f47)
	__obf_d627bfc120214f6d := int(__obf_da82f93e176f9f47[__obf_c4d3ca237ede8853-1])
	return __obf_da82f93e176f9f47[:(__obf_c4d3ca237ede8853 - __obf_d627bfc120214f6d)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_0754bb2dda81fe6a(__obf_2c9789f1f6d4df1a, __obf_a3cceebfb9c43a86, __obf_0e4b33157c20f699 []byte) ([]byte, error) {
	__obf_7c87d86ff2b2ece6, __obf_02af51ab07084000 := aes.NewCipher(__obf_2c9789f1f6d4df1a)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_02af51ab07084000)
	}
	__obf_cc6b0610c1e25e7a := __obf_7c87d86ff2b2ece6.BlockSize()
	if len(__obf_a3cceebfb9c43a86) != __obf_cc6b0610c1e25e7a {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_cc6b0610c1e25e7a)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_0e4b33157c20f699)%__obf_cc6b0610c1e25e7a != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	__obf_33718994b8b643a2 := cipher.NewCBCDecrypter(__obf_7c87d86ff2b2ece6, __obf_a3cceebfb9c43a86)
	__obf_a3d199cdb4ee502c := make([]byte, len(__obf_0e4b33157c20f699))
	__obf_33718994b8b643a2.
		CryptBlocks(__obf_a3d199cdb4ee502c, __obf_0e4b33157c20f699)
	__obf_a3d199cdb4ee502c = PKCS7UnPadding(__obf_a3d199cdb4ee502c)

	return __obf_a3d199cdb4ee502c, nil
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
func DecryptWXData(__obf_1c28a3c105a8521d, __obf_7339267d6ee76cfb, __obf_a3cceebfb9c43a86, __obf_2216a66f7c6d12ec string) (*DecryptedData, error) {
	__obf_66e386e50af569f8, __obf_02af51ab07084000 := base64.StdEncoding.DecodeString(__obf_1c28a3c105a8521d)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_02af51ab07084000)
	}
	__obf_9f6f98d93b750a0a, __obf_02af51ab07084000 := base64.StdEncoding.DecodeString(__obf_7339267d6ee76cfb)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_02af51ab07084000)
	}
	IV, __obf_02af51ab07084000 := base64.StdEncoding.DecodeString(__obf_a3cceebfb9c43a86)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_02af51ab07084000)
	}
	__obf_14f9417ce717c944,

		// 调用解密函数
		__obf_02af51ab07084000 := __obf_0754bb2dda81fe6a(__obf_66e386e50af569f8, IV, __obf_9f6f98d93b750a0a)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_02af51ab07084000)
	}

	var __obf_ee2c872ef85c5b21 DecryptedData
	__obf_02af51ab07084000 = json.Unmarshal(__obf_14f9417ce717c944, &__obf_ee2c872ef85c5b21)
	if __obf_02af51ab07084000 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_02af51ab07084000)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_ee2c872ef85c5b21.Watermark.AppID != __obf_2216a66f7c6d12ec {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_ee2c872ef85c5b21, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_a1ae217992a1ccec []byte, __obf_cc6b0610c1e25e7a int) []byte {
	__obf_69cb2282b6aebc89 := __obf_cc6b0610c1e25e7a - len(__obf_a1ae217992a1ccec)%__obf_cc6b0610c1e25e7a
	__obf_b41cd02e3a896f0d := bytes.Repeat([]byte{byte(__obf_69cb2282b6aebc89)}, __obf_69cb2282b6aebc89)
	return append(__obf_a1ae217992a1ccec, __obf_b41cd02e3a896f0d...)
}
