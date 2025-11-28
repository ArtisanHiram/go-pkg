package __obf_0d56cba5b7a269cc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"hash"
	"math/rand/v2"
	"sort"
	"strings"
)

// 微信签名算法方式

const (
	SIGN_MD5        = "MD5"
	SIGN_HMACSHA256 = "HMAC-SHA256"
)

// GenerateNonceStr 生成32位随机字符串
func GenerateNonceStr() string {
	__obf_942353a2ea318753 := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_2e779cb089cd7914 := 32
	__obf_3147a7d4a58a41c2 := make([]byte, __obf_2e779cb089cd7914)
	for __obf_7b79ebbb651cd203 := range __obf_3147a7d4a58a41c2 {
		__obf_3147a7d4a58a41c2[__obf_7b79ebbb651cd203] = __obf_942353a2ea318753[rand.IntN(len(__obf_942353a2ea318753))]
	}
	return string(__obf_3147a7d4a58a41c2)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_f163d1cbd6083f42 string, __obf_197c96d7c7100184 map[string]string) string {
	__obf_867e62e049901e6f := make([]string, 0, len(__obf_197c96d7c7100184))
	for __obf_9d120b736914ab35 := range __obf_197c96d7c7100184 {
		if __obf_9d120b736914ab35 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_867e62e049901e6f = append(__obf_867e62e049901e6f, __obf_9d120b736914ab35)
	}
	sort.Strings(__obf_867e62e049901e6f) // 字典序排序

	var __obf_11d84cb1458596ea bytes.Buffer
	for _, __obf_9d120b736914ab35 := range __obf_867e62e049901e6f {
		if __obf_197c96d7c7100184[__obf_9d120b736914ab35] != "" { // 空值不参与签名
			__obf_11d84cb1458596ea.WriteString(__obf_9d120b736914ab35)
			__obf_11d84cb1458596ea.WriteString("=")
			__obf_11d84cb1458596ea.WriteString(__obf_197c96d7c7100184[__obf_9d120b736914ab35])
			__obf_11d84cb1458596ea.WriteString("&")
		}
	}
	__obf_11d84cb1458596ea.WriteString("key=")
	__obf_11d84cb1458596ea.WriteString(__obf_f163d1cbd6083f42)

	var __obf_56351abfd3f6c6f8 hash.Hash
	if __obf_197c96d7c7100184["sign_type"] == SIGN_HMACSHA256 {
		__obf_56351abfd3f6c6f8 = hmac.New(sha256.New, []byte(__obf_f163d1cbd6083f42))
	} else {
		__obf_56351abfd3f6c6f8 = md5.New()
	}
	__obf_56351abfd3f6c6f8.Write(__obf_11d84cb1458596ea.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_56351abfd3f6c6f8.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_f163d1cbd6083f42 string, __obf_197c96d7c7100184 map[string]string) bool {
	__obf_ed55ba1e4b85e540, __obf_1b5d36c16ef12643 := __obf_197c96d7c7100184["sign"]
	if !__obf_1b5d36c16ef12643 {
		return false // 没有sign字段
	}
	delete(__obf_197c96d7c7100184, "sign") // 签名时需要移除sign字段

	__obf_38e97c164ed49741 := CreateSign(__obf_f163d1cbd6083f42, __obf_197c96d7c7100184)

	return __obf_38e97c164ed49741 == __obf_ed55ba1e4b85e540
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_50f63fd2e219c82b map[string]string) []byte {
	var __obf_11d84cb1458596ea bytes.Buffer
	__obf_11d84cb1458596ea.WriteString("<xml>")
	for __obf_9d120b736914ab35, __obf_24f87e0512d13675 := range __obf_50f63fd2e219c82b {
		__obf_11d84cb1458596ea.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_9d120b736914ab35, __obf_24f87e0512d13675, __obf_9d120b736914ab35))
	}
	__obf_11d84cb1458596ea.WriteString("</xml>")
	return __obf_11d84cb1458596ea.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_9d19eb6a8d0004d1 []byte) (map[string]string, error) {
	__obf_163ed2ce2d8f66f7 := xml.NewDecoder(bytes.NewReader(__obf_9d19eb6a8d0004d1))
	__obf_50f63fd2e219c82b := make(map[string]string)
	for {
		__obf_d96ee3f383afc7be, __obf_b8d74ed2e252a551 := __obf_163ed2ce2d8f66f7.Token()
		if __obf_b8d74ed2e252a551 != nil {
			if __obf_b8d74ed2e252a551.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_b8d74ed2e252a551)
		}

		switch __obf_602345d552ce44ce := __obf_d96ee3f383afc7be.(type) {
		case xml.StartElement:
			if __obf_602345d552ce44ce.Name.Local != "xml" { // 跳过根元素
				var __obf_e1b2c48858a2796f string
				if __obf_b8d74ed2e252a551 := __obf_163ed2ce2d8f66f7.DecodeElement(&__obf_e1b2c48858a2796f, &__obf_602345d552ce44ce); __obf_b8d74ed2e252a551 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_b8d74ed2e252a551)
				}
				__obf_50f63fd2e219c82b[__obf_602345d552ce44ce.Name.Local] = __obf_e1b2c48858a2796f
			}
		}
	}
	return __obf_50f63fd2e219c82b, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_0a47b02b4937f5a8, __obf_918f725cfa472438 []byte, __obf_b0103a38b761ea9a, __obf_17010ca3dd1070a8 string) (__obf_843c1533a9cb0f7b []byte, __obf_b8d74ed2e252a551 error) {
	defer func() {
		if __obf_8167a10fab469cd2 := recover(); __obf_8167a10fab469cd2 != nil {
			__obf_b8d74ed2e252a551 = fmt.Errorf("panic error: err=%v", __obf_8167a10fab469cd2)
			return
		}
	}()
	var __obf_c645f01f5ce133fe []byte
	__obf_c645f01f5ce133fe, __obf_b8d74ed2e252a551 = __obf_627c404a69858bc7(__obf_17010ca3dd1070a8)
	if __obf_b8d74ed2e252a551 != nil {
		panic(__obf_b8d74ed2e252a551)
	}
	__obf_8a648f5b60541c74 := AESEncryptMsg(__obf_0a47b02b4937f5a8, __obf_918f725cfa472438, __obf_b0103a38b761ea9a, __obf_c645f01f5ce133fe)
	__obf_843c1533a9cb0f7b = []byte(base64.StdEncoding.EncodeToString(__obf_8a648f5b60541c74))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_0a47b02b4937f5a8, __obf_918f725cfa472438 []byte, __obf_b0103a38b761ea9a string, __obf_17010ca3dd1070a8 []byte) (__obf_8a648f5b60541c74 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	__obf_0c96eed961b6aed1 := 20 + len(__obf_918f725cfa472438)
	__obf_687b966f9bf2f49c := __obf_0c96eed961b6aed1 + len(__obf_b0103a38b761ea9a)
	__obf_0fc771bc13ba282a := BlockSize - __obf_687b966f9bf2f49c&BlockMask
	__obf_abfb27602c9307ba := __obf_687b966f9bf2f49c + __obf_0fc771bc13ba282a

	__obf_af8b84ddd80245d1 := make([]byte, __obf_abfb27602c9307ba)

	// 拼接
	copy(__obf_af8b84ddd80245d1[:16], __obf_0a47b02b4937f5a8)
	__obf_264737b3e28dd411(__obf_af8b84ddd80245d1[16:20], uint32(len(__obf_918f725cfa472438)))
	copy(__obf_af8b84ddd80245d1[20:], __obf_918f725cfa472438)
	copy(__obf_af8b84ddd80245d1[__obf_0c96eed961b6aed1:], __obf_b0103a38b761ea9a)

	// PKCS#7 补位
	for __obf_7b79ebbb651cd203 := __obf_687b966f9bf2f49c; __obf_7b79ebbb651cd203 < __obf_abfb27602c9307ba; __obf_7b79ebbb651cd203++ {
		__obf_af8b84ddd80245d1[__obf_7b79ebbb651cd203] = byte(__obf_0fc771bc13ba282a)
	}

	// 加密
	__obf_bca2a6aab9d7b6cf, __obf_b8d74ed2e252a551 := aes.NewCipher(__obf_17010ca3dd1070a8)
	if __obf_b8d74ed2e252a551 != nil {
		panic(__obf_b8d74ed2e252a551)
	}
	__obf_d35205862fdf2b7c := cipher.NewCBCEncrypter(__obf_bca2a6aab9d7b6cf, __obf_17010ca3dd1070a8[:16])
	__obf_d35205862fdf2b7c.CryptBlocks(__obf_af8b84ddd80245d1, __obf_af8b84ddd80245d1)

	__obf_8a648f5b60541c74 = __obf_af8b84ddd80245d1
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_b0103a38b761ea9a, __obf_6f13a80565c5261b, __obf_17010ca3dd1070a8 string) (__obf_0a47b02b4937f5a8, __obf_4c807e301d64248e []byte, __obf_b8d74ed2e252a551 error) {
	defer func() {
		if __obf_8167a10fab469cd2 := recover(); __obf_8167a10fab469cd2 != nil {
			__obf_b8d74ed2e252a551 = fmt.Errorf("panic error: err=%v", __obf_8167a10fab469cd2)
			return
		}
	}()
	var __obf_06c02bc0d25a9b8f, __obf_c645f01f5ce133fe, __obf_e6d13a1e75ba89a4 []byte
	__obf_06c02bc0d25a9b8f, __obf_b8d74ed2e252a551 = base64.StdEncoding.DecodeString(__obf_6f13a80565c5261b)
	if __obf_b8d74ed2e252a551 != nil {
		return
	}
	__obf_c645f01f5ce133fe, __obf_b8d74ed2e252a551 = __obf_627c404a69858bc7(__obf_17010ca3dd1070a8)
	if __obf_b8d74ed2e252a551 != nil {
		panic(__obf_b8d74ed2e252a551)
	}
	__obf_0a47b02b4937f5a8, __obf_4c807e301d64248e, __obf_e6d13a1e75ba89a4, __obf_b8d74ed2e252a551 = AESDecryptMsg(__obf_06c02bc0d25a9b8f, __obf_c645f01f5ce133fe)
	if __obf_b8d74ed2e252a551 != nil {
		__obf_b8d74ed2e252a551 = fmt.Errorf("消息解密失败,%v", __obf_b8d74ed2e252a551)
		return
	}
	if __obf_b0103a38b761ea9a != string(__obf_e6d13a1e75ba89a4) {
		__obf_b8d74ed2e252a551 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_627c404a69858bc7(__obf_95ab41da1342e180 string) (__obf_c645f01f5ce133fe []byte, __obf_b8d74ed2e252a551 error) {
	if len(__obf_95ab41da1342e180) != 43 {
		__obf_b8d74ed2e252a551 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_c645f01f5ce133fe, __obf_b8d74ed2e252a551 = base64.StdEncoding.DecodeString(__obf_95ab41da1342e180 + "=")
	if __obf_b8d74ed2e252a551 != nil {
		return
	}
	if len(__obf_c645f01f5ce133fe) != 32 {
		__obf_b8d74ed2e252a551 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_8a648f5b60541c74 []byte, __obf_17010ca3dd1070a8 []byte) (__obf_0a47b02b4937f5a8, __obf_918f725cfa472438, __obf_b0103a38b761ea9a []byte, __obf_b8d74ed2e252a551 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_8a648f5b60541c74) < BlockSize {
		__obf_b8d74ed2e252a551 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_8a648f5b60541c74))
		return
	}
	if len(__obf_8a648f5b60541c74)&BlockMask != 0 {
		__obf_b8d74ed2e252a551 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_8a648f5b60541c74))
		return
	}

	__obf_af8b84ddd80245d1 := make([]byte, len(__obf_8a648f5b60541c74)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	__obf_bca2a6aab9d7b6cf, __obf_b8d74ed2e252a551 := aes.NewCipher(__obf_17010ca3dd1070a8)
	if __obf_b8d74ed2e252a551 != nil {
		panic(__obf_b8d74ed2e252a551)
	}
	__obf_d35205862fdf2b7c := cipher.NewCBCDecrypter(__obf_bca2a6aab9d7b6cf, __obf_17010ca3dd1070a8[:16])
	__obf_d35205862fdf2b7c.CryptBlocks(__obf_af8b84ddd80245d1, __obf_8a648f5b60541c74)

	// PKCS#7 去除补位
	__obf_0fc771bc13ba282a := int(__obf_af8b84ddd80245d1[len(__obf_af8b84ddd80245d1)-1])
	if __obf_0fc771bc13ba282a < 1 || __obf_0fc771bc13ba282a > BlockSize {
		__obf_b8d74ed2e252a551 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_0fc771bc13ba282a)
		return
	}
	__obf_af8b84ddd80245d1 = __obf_af8b84ddd80245d1[:len(__obf_af8b84ddd80245d1)-__obf_0fc771bc13ba282a]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_af8b84ddd80245d1) <= 20 {
		__obf_b8d74ed2e252a551 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_af8b84ddd80245d1))
		return
	}
	__obf_210249dfbca13816 := int(__obf_e5652574dd50a040(__obf_af8b84ddd80245d1[16:20]))
	if __obf_210249dfbca13816 < 0 {
		__obf_b8d74ed2e252a551 = fmt.Errorf("incorrect msg length: %d", __obf_210249dfbca13816)
		return
	}
	__obf_0c96eed961b6aed1 := 20 + __obf_210249dfbca13816
	if len(__obf_af8b84ddd80245d1) <= __obf_0c96eed961b6aed1 {
		__obf_b8d74ed2e252a551 = fmt.Errorf("msg length too large: %d", __obf_210249dfbca13816)
		return
	}

	__obf_0a47b02b4937f5a8 = __obf_af8b84ddd80245d1[:16:20]
	__obf_918f725cfa472438 = __obf_af8b84ddd80245d1[20:__obf_0c96eed961b6aed1:__obf_0c96eed961b6aed1]
	__obf_b0103a38b761ea9a = __obf_af8b84ddd80245d1[__obf_0c96eed961b6aed1:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_264737b3e28dd411(__obf_7e97b03ad792ae78 []byte, __obf_9b39de4bbb2aad3a uint32) {
	__obf_7e97b03ad792ae78[0] = byte(__obf_9b39de4bbb2aad3a >> 24)
	__obf_7e97b03ad792ae78[1] = byte(__obf_9b39de4bbb2aad3a >> 16)
	__obf_7e97b03ad792ae78[2] = byte(__obf_9b39de4bbb2aad3a >> 8)
	__obf_7e97b03ad792ae78[3] = byte(__obf_9b39de4bbb2aad3a)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_e5652574dd50a040(__obf_7e97b03ad792ae78 []byte) (__obf_9b39de4bbb2aad3a uint32) {
	return uint32(__obf_7e97b03ad792ae78[0])<<24 |
		uint32(__obf_7e97b03ad792ae78[1])<<16 |
		uint32(__obf_7e97b03ad792ae78[2])<<8 |
		uint32(__obf_7e97b03ad792ae78[3])
}
