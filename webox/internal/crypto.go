package __obf_18f8d68b9095d0e0

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
func PKCS7UnPadding(__obf_870163b5b3d35e1e []byte) []byte {
	__obf_7e52c1900ee9de1e := len(__obf_870163b5b3d35e1e)
	__obf_187676f1038f31f4 := int(__obf_870163b5b3d35e1e[__obf_7e52c1900ee9de1e-1])
	return __obf_870163b5b3d35e1e[:(__obf_7e52c1900ee9de1e - __obf_187676f1038f31f4)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_3ae481e4ebd31b77(__obf_74e39c87e7c65a0b, __obf_6d21d5c7da6155cf, __obf_c1be1f54e3a65d3a []byte) ([]byte, error) {
	__obf_6b1f50188faf4395, __obf_f45e33988e88e5a2 := aes.NewCipher(__obf_74e39c87e7c65a0b)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_353b15d68456d3a2 := __obf_6b1f50188faf4395.BlockSize()
	if len(__obf_6d21d5c7da6155cf) != __obf_353b15d68456d3a2 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_353b15d68456d3a2)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_c1be1f54e3a65d3a)%__obf_353b15d68456d3a2 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	__obf_9a43bc2ee75371b3 := cipher.NewCBCDecrypter(__obf_6b1f50188faf4395, __obf_6d21d5c7da6155cf)
	__obf_2a3f9ef0d90702f0 := make([]byte, len(__obf_c1be1f54e3a65d3a))
	__obf_9a43bc2ee75371b3.
		CryptBlocks(__obf_2a3f9ef0d90702f0, __obf_c1be1f54e3a65d3a)
	__obf_2a3f9ef0d90702f0 = PKCS7UnPadding(__obf_2a3f9ef0d90702f0)

	return __obf_2a3f9ef0d90702f0, nil
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
func DecryptWXData(__obf_2f0b14e2824bb92f, __obf_f5a6a65c703bfa12, __obf_6d21d5c7da6155cf, __obf_64ec3c55bc656ff0 string) (*DecryptedData, error) {
	__obf_bd4b6c03e14da7f0, __obf_f45e33988e88e5a2 := base64.StdEncoding.DecodeString(__obf_2f0b14e2824bb92f)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_c61b621ed0d9b7f5, __obf_f45e33988e88e5a2 := base64.StdEncoding.DecodeString(__obf_f5a6a65c703bfa12)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_f45e33988e88e5a2)
	}
	IV, __obf_f45e33988e88e5a2 := base64.StdEncoding.DecodeString(__obf_6d21d5c7da6155cf)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_f45e33988e88e5a2)
	}
	__obf_8a6e6957fa93c280,

		// 调用解密函数
		__obf_f45e33988e88e5a2 := __obf_3ae481e4ebd31b77(__obf_bd4b6c03e14da7f0, IV, __obf_c61b621ed0d9b7f5)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_f45e33988e88e5a2)
	}

	var __obf_2a0c5f16e13ec5e3 DecryptedData
	__obf_f45e33988e88e5a2 = json.Unmarshal(__obf_8a6e6957fa93c280, &__obf_2a0c5f16e13ec5e3)
	if __obf_f45e33988e88e5a2 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_f45e33988e88e5a2)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_2a0c5f16e13ec5e3.Watermark.AppID != __obf_64ec3c55bc656ff0 {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_2a0c5f16e13ec5e3, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_e363473771864094 []byte, __obf_353b15d68456d3a2 int) []byte {
	__obf_e816c6d487cea6ef := __obf_353b15d68456d3a2 - len(__obf_e363473771864094)%__obf_353b15d68456d3a2
	__obf_97e5fb481d5d8002 := bytes.Repeat([]byte{byte(__obf_e816c6d487cea6ef)}, __obf_e816c6d487cea6ef)
	return append(__obf_e363473771864094, __obf_97e5fb481d5d8002...)
}
