package __obf_c0d7bb2e04898f29

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
func PKCS7UnPadding(__obf_ec0b802a7c3b0b84 []byte) []byte {
	__obf_c5ab5df2052940ae := len(__obf_ec0b802a7c3b0b84)
	__obf_40caeac0a5f72563 := int(__obf_ec0b802a7c3b0b84[__obf_c5ab5df2052940ae-1])
	return __obf_ec0b802a7c3b0b84[:(__obf_c5ab5df2052940ae - __obf_40caeac0a5f72563)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_44db3e596ec030ca(__obf_9732b72b1933d195, __obf_5c8c1cf07aeb8ea0, __obf_301e914878944cf9 []byte) ([]byte, error) {
	__obf_255c57c66bf4666e, __obf_fb6d611ef291e586 := aes.NewCipher(__obf_9732b72b1933d195)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_fb6d611ef291e586)
	}

	__obf_55d9bdb8821e9b06 := __obf_255c57c66bf4666e.BlockSize()
	if len(__obf_5c8c1cf07aeb8ea0) != __obf_55d9bdb8821e9b06 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_55d9bdb8821e9b06)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_301e914878944cf9)%__obf_55d9bdb8821e9b06 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	__obf_03558c620be6f23d := cipher.NewCBCDecrypter(__obf_255c57c66bf4666e, __obf_5c8c1cf07aeb8ea0)
	__obf_59e182d5f23e07c4 := make([]byte, len(__obf_301e914878944cf9))
	__obf_03558c620be6f23d.CryptBlocks(__obf_59e182d5f23e07c4, __obf_301e914878944cf9)

	__obf_59e182d5f23e07c4 = PKCS7UnPadding(__obf_59e182d5f23e07c4)

	return __obf_59e182d5f23e07c4, nil
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
func DecryptWXData(__obf_4037d0971c53023f, __obf_0b11e63c837812a5, __obf_5c8c1cf07aeb8ea0, __obf_6ea411bfa0bdb6dc string) (*DecryptedData, error) {
	__obf_3806dfb1e5ce3ca8, __obf_fb6d611ef291e586 := base64.StdEncoding.DecodeString(__obf_4037d0971c53023f)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_fb6d611ef291e586)
	}
	__obf_848c66347b289b90, __obf_fb6d611ef291e586 := base64.StdEncoding.DecodeString(__obf_0b11e63c837812a5)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_fb6d611ef291e586)
	}
	IV, __obf_fb6d611ef291e586 := base64.StdEncoding.DecodeString(__obf_5c8c1cf07aeb8ea0)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_fb6d611ef291e586)
	}

	// 调用解密函数
	__obf_ade91b5b823e2f80, __obf_fb6d611ef291e586 := __obf_44db3e596ec030ca(__obf_3806dfb1e5ce3ca8, IV, __obf_848c66347b289b90)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_fb6d611ef291e586)
	}

	var __obf_c313e5576bf980c7 DecryptedData
	__obf_fb6d611ef291e586 = json.Unmarshal(__obf_ade91b5b823e2f80, &__obf_c313e5576bf980c7)
	if __obf_fb6d611ef291e586 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_fb6d611ef291e586)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_c313e5576bf980c7.Watermark.AppID != __obf_6ea411bfa0bdb6dc {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_c313e5576bf980c7, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_2351672914163ab0 []byte, __obf_55d9bdb8821e9b06 int) []byte {
	__obf_324358ee4d6d09b9 := __obf_55d9bdb8821e9b06 - len(__obf_2351672914163ab0)%__obf_55d9bdb8821e9b06
	__obf_a9975069d947fdef := bytes.Repeat([]byte{byte(__obf_324358ee4d6d09b9)}, __obf_324358ee4d6d09b9)
	return append(__obf_2351672914163ab0, __obf_a9975069d947fdef...)
}
