package __obf_ca26b678cc0b7836

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
func PKCS7UnPadding(__obf_4e0bc07420f56507 []byte) []byte {
	__obf_fa90f02e738a820d := len(__obf_4e0bc07420f56507)
	__obf_d95ff98d62760c08 := int(__obf_4e0bc07420f56507[__obf_fa90f02e738a820d-1])
	return __obf_4e0bc07420f56507[:(__obf_fa90f02e738a820d - __obf_d95ff98d62760c08)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_a1a12c496757f095(__obf_9fc2e8cc9640f833, __obf_3e9997940b14353e, __obf_44d8a3a68f381b41 []byte) ([]byte, error) {
	__obf_ff17270f2d6d21bd, __obf_c3e83fee90661fe5 := aes.NewCipher(__obf_9fc2e8cc9640f833)
	if __obf_c3e83fee90661fe5 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_c3e83fee90661fe5)
	}
	__obf_dfb909c8bff966b9 := __obf_ff17270f2d6d21bd.BlockSize()
	if len(__obf_3e9997940b14353e) != __obf_dfb909c8bff966b9 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_dfb909c8bff966b9)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_44d8a3a68f381b41)%__obf_dfb909c8bff966b9 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	__obf_f11ad58c651cc535 := cipher.NewCBCDecrypter(__obf_ff17270f2d6d21bd, __obf_3e9997940b14353e)
	__obf_da078e15c9a4521d := make([]byte, len(__obf_44d8a3a68f381b41))
	__obf_f11ad58c651cc535.
		CryptBlocks(__obf_da078e15c9a4521d, __obf_44d8a3a68f381b41)
	__obf_da078e15c9a4521d = PKCS7UnPadding(__obf_da078e15c9a4521d)

	return __obf_da078e15c9a4521d, nil
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
func DecryptWXData(__obf_14720b2832e859ae, __obf_f570a0974b38fd60, __obf_3e9997940b14353e, __obf_0d37b485cacb2dde string) (*DecryptedData, error) {
	__obf_e41e5a50ecb0df07, __obf_c3e83fee90661fe5 := base64.StdEncoding.DecodeString(__obf_14720b2832e859ae)
	if __obf_c3e83fee90661fe5 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_c3e83fee90661fe5)
	}
	__obf_758a346969013319, __obf_c3e83fee90661fe5 := base64.StdEncoding.DecodeString(__obf_f570a0974b38fd60)
	if __obf_c3e83fee90661fe5 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_c3e83fee90661fe5)
	}
	IV, __obf_c3e83fee90661fe5 := base64.StdEncoding.DecodeString(__obf_3e9997940b14353e)
	if __obf_c3e83fee90661fe5 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_c3e83fee90661fe5)
	}
	__obf_c90f6dc6de65ad97,

		// 调用解密函数
		__obf_c3e83fee90661fe5 := __obf_a1a12c496757f095(__obf_e41e5a50ecb0df07, IV, __obf_758a346969013319)
	if __obf_c3e83fee90661fe5 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_c3e83fee90661fe5)
	}

	var __obf_3548c160ae91e2ec DecryptedData
	__obf_c3e83fee90661fe5 = json.Unmarshal(__obf_c90f6dc6de65ad97, &__obf_3548c160ae91e2ec)
	if __obf_c3e83fee90661fe5 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_c3e83fee90661fe5)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_3548c160ae91e2ec.Watermark.AppID != __obf_0d37b485cacb2dde {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_3548c160ae91e2ec, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_f4161864736f8982 []byte, __obf_dfb909c8bff966b9 int) []byte {
	__obf_83652eaf6d85db23 := __obf_dfb909c8bff966b9 - len(__obf_f4161864736f8982)%__obf_dfb909c8bff966b9
	__obf_d29d34f165e7da83 := bytes.Repeat([]byte{byte(__obf_83652eaf6d85db23)}, __obf_83652eaf6d85db23)
	return append(__obf_f4161864736f8982, __obf_d29d34f165e7da83...)
}
