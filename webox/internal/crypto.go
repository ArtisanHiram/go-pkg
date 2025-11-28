package __obf_6a74725dc3694218

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
func PKCS7UnPadding(__obf_bbf1643760982270 []byte) []byte {
	__obf_da394bd28ba30c31 := len(__obf_bbf1643760982270)
	__obf_94ca6b623b2865fa := int(__obf_bbf1643760982270[__obf_da394bd28ba30c31-1])
	return __obf_bbf1643760982270[:(__obf_da394bd28ba30c31 - __obf_94ca6b623b2865fa)]
}

// decrypt 对称解密
// key：AES密钥，必须是16、24或32字节
// iv：初始向量，必须是16字节
// cipherText：待解密的密文
func __obf_895bcd4316ac4d51(__obf_a5717b59e2d4a940, __obf_99a31164af350413, __obf_82416cea1d486431 []byte) ([]byte, error) {
	__obf_b22ea7160f9e6175, __obf_b5433c9ceabad42d := aes.NewCipher(__obf_a5717b59e2d4a940)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("create AES cipher failed: %w", __obf_b5433c9ceabad42d)
	}

	__obf_ea0e3784e085a24e := __obf_b22ea7160f9e6175.BlockSize()
	if len(__obf_99a31164af350413) != __obf_ea0e3784e085a24e {
		return nil, fmt.Errorf("iv length must be %d bytes", __obf_ea0e3784e085a24e)
	}

	// 密文长度必须是分组的整数倍
	if len(__obf_82416cea1d486431)%__obf_ea0e3784e085a24e != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	__obf_09b4c555ad0085dd := cipher.NewCBCDecrypter(__obf_b22ea7160f9e6175, __obf_99a31164af350413)
	__obf_ccdc4dfcc8a2d487 := make([]byte, len(__obf_82416cea1d486431))
	__obf_09b4c555ad0085dd.CryptBlocks(__obf_ccdc4dfcc8a2d487, __obf_82416cea1d486431)

	__obf_ccdc4dfcc8a2d487 = PKCS7UnPadding(__obf_ccdc4dfcc8a2d487)

	return __obf_ccdc4dfcc8a2d487, nil
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
func DecryptWXData(__obf_8e85d36abff7f307, __obf_473cdea6ff67ff49, __obf_99a31164af350413, __obf_6e118b932609e07d string) (*DecryptedData, error) {
	__obf_a9548c252331feaa, __obf_b5433c9ceabad42d := base64.StdEncoding.DecodeString(__obf_8e85d36abff7f307)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("decode sessionKey failed: %w", __obf_b5433c9ceabad42d)
	}
	__obf_4e44c7f94fd3cb65, __obf_b5433c9ceabad42d := base64.StdEncoding.DecodeString(__obf_473cdea6ff67ff49)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("decode encryptedData failed: %w", __obf_b5433c9ceabad42d)
	}
	IV, __obf_b5433c9ceabad42d := base64.StdEncoding.DecodeString(__obf_99a31164af350413)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("decode iv failed: %w", __obf_b5433c9ceabad42d)
	}

	// 调用解密函数
	__obf_eb780ff9bf2c0f36, __obf_b5433c9ceabad42d := __obf_895bcd4316ac4d51(__obf_a9548c252331feaa, IV, __obf_4e44c7f94fd3cb65)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("decrypt data failed: %w", __obf_b5433c9ceabad42d)
	}

	var __obf_343c2aee43bfce9f DecryptedData
	__obf_b5433c9ceabad42d = json.Unmarshal(__obf_eb780ff9bf2c0f36, &__obf_343c2aee43bfce9f)
	if __obf_b5433c9ceabad42d != nil {
		return nil, fmt.Errorf("unmarshal decrypted data failed: %w", __obf_b5433c9ceabad42d)
	}

	// 验证AppID是否一致，防止恶意攻击或数据错用
	if __obf_343c2aee43bfce9f.Watermark.AppID != __obf_6e118b932609e07d {
		return nil, errors.New("app_id in watermark does not match current app_id")
	}

	return &__obf_343c2aee43bfce9f, nil
}

// PKCS7Padding PKCS7填充
func PKCS7Padding(__obf_310f083e674cdc61 []byte, __obf_ea0e3784e085a24e int) []byte {
	__obf_eb9f261bde505434 := __obf_ea0e3784e085a24e - len(__obf_310f083e674cdc61)%__obf_ea0e3784e085a24e
	__obf_24af6223cb9c60b8 := bytes.Repeat([]byte{byte(__obf_eb9f261bde505434)}, __obf_eb9f261bde505434)
	return append(__obf_310f083e674cdc61, __obf_24af6223cb9c60b8...)
}
