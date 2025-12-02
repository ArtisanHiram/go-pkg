package __obf_dfacbdcb930f08e7

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
func PKCS7UnPadding(__obf_2c586f5f2ca8bee1 []byte) []byte {
	__obf_703404ed2aca900b := len(__obf_2c586f5f2ca8bee1)
	__obf_fb100789d9a6817c := int(__obf_2c586f5f2ca8bee1[__obf_703404ed2aca900b-1])
	return __obf_2c586f5f2ca8bee1[:(__obf_703404ed2aca900b - __obf_fb100789d9a6817c)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_e222655f42972299(__obf_fdb60bda187c670a, __obf_05f2287c726298aa, __obf_1270a095a72eb45b []byte) ([]byte, error) {
	__obf_e0eb1010f1f53051, __obf_f304d4a353574273 := aes.NewCipher(__obf_fdb60bda187c670a)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_f304d4a353574273)
	}
	__obf_58facc9d1938574a := __obf_e0eb1010f1f53051.BlockSize()
	if len(__obf_05f2287c726298aa) != __obf_58facc9d1938574a {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_58facc9d1938574a)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_1270a095a72eb45b)%__obf_58facc9d1938574a != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	__obf_05d971a2026e6251 := cipher.NewCBCDecrypter(__obf_e0eb1010f1f53051, __obf_05f2287c726298aa)
	__obf_4aa802fb79b22d06 := make([]byte, len(__obf_1270a095a72eb45b))
	__obf_05d971a2026e6251.
		CryptBlocks(__obf_4aa802fb79b22d06, __obf_1270a095a72eb45b)
	__obf_4aa802fb79b22d06 = PKCS7UnPadding(__obf_4aa802fb79b22d06)

	return __obf_4aa802fb79b22d06, nil
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
func DecryptWXData(__obf_13df57b2735cdb90, __obf_e2400235f067ee6a, __obf_05f2287c726298aa, __obf_f20438a20f203203 string) (*DecryptedData, error) {
	__obf_6f281e496e3314f6, __obf_f304d4a353574273 := base64.StdEncoding.DecodeString(__obf_13df57b2735cdb90)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_f304d4a353574273)
	}
	__obf_0234040eccf5ae1b, __obf_f304d4a353574273 := base64.StdEncoding.DecodeString(__obf_e2400235f067ee6a)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_f304d4a353574273)
	}
	IV, __obf_f304d4a353574273 := base64.StdEncoding.DecodeString(__obf_05f2287c726298aa)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_f304d4a353574273)
	}
	__obf_5f92794eda895576,

		// 调用解密函数
		__obf_f304d4a353574273 := __obf_e222655f42972299(__obf_6f281e496e3314f6, IV, __obf_0234040eccf5ae1b)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_f304d4a353574273)
	}

	var __obf_87f37beff62d945f DecryptedData
	__obf_f304d4a353574273 = json.Unmarshal(__obf_5f92794eda895576, &__obf_87f37beff62d945f)
	if __obf_f304d4a353574273 != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_f304d4a353574273)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_87f37beff62d945f.Watermark.AppID != __obf_f20438a20f203203 {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_87f37beff62d945f, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_918e5a8ab8b15df6 []byte, __obf_58facc9d1938574a int) []byte {
	__obf_a785f0f107d065e7 := __obf_58facc9d1938574a - len(__obf_918e5a8ab8b15df6)%__obf_58facc9d1938574a
	__obf_d92b6781cb40a4f5 := bytes.Repeat([]byte{byte(__obf_a785f0f107d065e7)}, __obf_a785f0f107d065e7)
	return append(__obf_918e5a8ab8b15df6, __obf_d92b6781cb40a4f5...)
}
