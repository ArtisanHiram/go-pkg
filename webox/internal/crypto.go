package __obf_79882becdca92e5e

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
func PKCS7UnPadding(__obf_abf4396a37976f66 []byte) []byte {
	__obf_205c9a2417fbe739 := len(__obf_abf4396a37976f66)
	__obf_60e8c0c412aabe89 := int(__obf_abf4396a37976f66[__obf_205c9a2417fbe739-1])
	return __obf_abf4396a37976f66[:(__obf_205c9a2417fbe739 - __obf_60e8c0c412aabe89)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_78df44c48ecad067(__obf_d3be629120925ff5, __obf_8a61976de1ffceeb, __obf_99755d3dbf7f15da []byte) ([]byte, error) {
	__obf_8301dc39bcbe4dde, __obf_d0090ed9a7614e3f := aes.NewCipher(__obf_d3be629120925ff5)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_d0090ed9a7614e3f)
	}

	__obf_1caf3f432e198142 := __obf_8301dc39bcbe4dde.BlockSize()
	if len(__obf_8a61976de1ffceeb) != __obf_1caf3f432e198142 {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_1caf3f432e198142)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_99755d3dbf7f15da)%__obf_1caf3f432e198142 != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	__obf_8ec3a75dfa98671c := cipher.NewCBCDecrypter(__obf_8301dc39bcbe4dde, __obf_8a61976de1ffceeb)
	__obf_9a867ab14432830b := make([]byte, len(__obf_99755d3dbf7f15da))
	__obf_8ec3a75dfa98671c.CryptBlocks(__obf_9a867ab14432830b, __obf_99755d3dbf7f15da)

	__obf_9a867ab14432830b = PKCS7UnPadding(__obf_9a867ab14432830b)

	return __obf_9a867ab14432830b, nil
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
func DecryptWXData(__obf_3a4e67c12d70ddd7, __obf_e3bf590599baba98, __obf_8a61976de1ffceeb, __obf_cbfc233e57c7740c string) (*DecryptedData, error) {
	__obf_ce86165d91a5fcb6, __obf_d0090ed9a7614e3f := base64.StdEncoding.DecodeString(__obf_3a4e67c12d70ddd7)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_d0090ed9a7614e3f)
	}
	__obf_7f219b4e60155d45, __obf_d0090ed9a7614e3f := base64.StdEncoding.DecodeString(__obf_e3bf590599baba98)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_d0090ed9a7614e3f)
	}
	IV, __obf_d0090ed9a7614e3f := base64.StdEncoding.DecodeString(__obf_8a61976de1ffceeb)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_d0090ed9a7614e3f)
	}

	// 调用解密函数
	__obf_7df37ca6565da1b8, __obf_d0090ed9a7614e3f := __obf_78df44c48ecad067(__obf_ce86165d91a5fcb6, IV, __obf_7f219b4e60155d45)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_d0090ed9a7614e3f)
	}

	var __obf_d292a8768a7a03ae DecryptedData
	__obf_d0090ed9a7614e3f = json.Unmarshal(__obf_7df37ca6565da1b8, &__obf_d292a8768a7a03ae)
	if __obf_d0090ed9a7614e3f != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_d0090ed9a7614e3f)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_d292a8768a7a03ae.Watermark.AppID != __obf_cbfc233e57c7740c {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_d292a8768a7a03ae, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_3f5e435edf8d2673 []byte, __obf_1caf3f432e198142 int) []byte {
	__obf_cbcb40d3b8a1b652 := __obf_1caf3f432e198142 - len(__obf_3f5e435edf8d2673)%__obf_1caf3f432e198142
	__obf_1fcb793a446dba9e := bytes.Repeat([]byte{byte(__obf_cbcb40d3b8a1b652)}, __obf_cbcb40d3b8a1b652)
	return append(__obf_3f5e435edf8d2673, __obf_1fcb793a446dba9e...)
}
