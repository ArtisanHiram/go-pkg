package __obf_0d56cba5b7a269cc

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
func PKCS7UnPadding(__obf_dc64b2ff0623a28e []byte) []byte {
	__obf_2e779cb089cd7914 := len(__obf_dc64b2ff0623a28e)
	__obf_b14dc7bbb7e3552e := int(__obf_dc64b2ff0623a28e[__obf_2e779cb089cd7914-1])
	return __obf_dc64b2ff0623a28e[:(__obf_2e779cb089cd7914 - __obf_b14dc7bbb7e3552e)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_3a4525da6f86d1db(__obf_c645f01f5ce133fe, __obf_1c48d9e190fa58ac, __obf_9d7f7cf680f6e99b []byte) ([]byte, error) {
	__obf_bca2a6aab9d7b6cf, __obf_b8d74ed2e252a551 := aes.NewCipher(__obf_c645f01f5ce133fe)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_b8d74ed2e252a551)
	}

	__obf_5a9789dcb06c8a85 := __obf_bca2a6aab9d7b6cf.BlockSize()
	if len(__obf_1c48d9e190fa58ac) != __obf_5a9789dcb06c8a85 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_5a9789dcb06c8a85)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_9d7f7cf680f6e99b)%__obf_5a9789dcb06c8a85 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	__obf_33649538bc9fd655 := cipher.NewCBCDecrypter(__obf_bca2a6aab9d7b6cf, __obf_1c48d9e190fa58ac)
	__obf_28884f3db1224916 := make([]byte, len(__obf_9d7f7cf680f6e99b))
	__obf_33649538bc9fd655.CryptBlocks(__obf_28884f3db1224916, __obf_9d7f7cf680f6e99b)

	__obf_28884f3db1224916 = PKCS7UnPadding(__obf_28884f3db1224916)

	return __obf_28884f3db1224916, nil
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
func DecryptWXData(__obf_b3c0412230c4eb4c, __obf_cd74e7cf285f3919, __obf_1c48d9e190fa58ac, __obf_667c72ca0ddfad40 string) (*DecryptedData, error) {
	__obf_4153f759a4c18a07, __obf_b8d74ed2e252a551 := base64.StdEncoding.DecodeString(__obf_b3c0412230c4eb4c)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_b8d74ed2e252a551)
	}
	__obf_6bea72b72cca5bcc, __obf_b8d74ed2e252a551 := base64.StdEncoding.DecodeString(__obf_cd74e7cf285f3919)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_b8d74ed2e252a551)
	}
	IV, __obf_b8d74ed2e252a551 := base64.StdEncoding.DecodeString(__obf_1c48d9e190fa58ac)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_b8d74ed2e252a551)
	}

	// 调用解密函数
	__obf_034f2a7d3ec06845, __obf_b8d74ed2e252a551 := __obf_3a4525da6f86d1db(__obf_4153f759a4c18a07, IV, __obf_6bea72b72cca5bcc)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_b8d74ed2e252a551)
	}

	var __obf_50f63fd2e219c82b DecryptedData
	__obf_b8d74ed2e252a551 = json.Unmarshal(__obf_034f2a7d3ec06845, &__obf_50f63fd2e219c82b)
	if __obf_b8d74ed2e252a551 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_b8d74ed2e252a551)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_50f63fd2e219c82b.Watermark.AppID != __obf_667c72ca0ddfad40 {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_50f63fd2e219c82b, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_8a648f5b60541c74 []byte, __obf_5a9789dcb06c8a85 int) []byte {
	__obf_e6f3ca833f6e2a4e := __obf_5a9789dcb06c8a85 - len(__obf_8a648f5b60541c74)%__obf_5a9789dcb06c8a85
	__obf_ea4e4bcf4feab016 := bytes.Repeat([]byte{byte(__obf_e6f3ca833f6e2a4e)}, __obf_e6f3ca833f6e2a4e)
	return append(__obf_8a648f5b60541c74, __obf_ea4e4bcf4feab016...)
}
