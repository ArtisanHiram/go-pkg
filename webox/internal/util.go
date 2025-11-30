package __obf_e79e1b0b12d02d0b

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
	__obf_5c29fefb6113fa4a := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_049fadfb61efc62d := 32
	__obf_61a2594eea4ae9c0 := make([]byte, __obf_049fadfb61efc62d)
	for __obf_33d1c6b9ec98c730 := range __obf_61a2594eea4ae9c0 {
		__obf_61a2594eea4ae9c0[__obf_33d1c6b9ec98c730] = __obf_5c29fefb6113fa4a[rand.IntN(len(__obf_5c29fefb6113fa4a))]
	}
	return string(__obf_61a2594eea4ae9c0)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_434013c08d7e96ff string, __obf_349df6ea515f02e1 map[string]string) string {
	__obf_b87bc669672328e4 := make([]string, 0, len(__obf_349df6ea515f02e1))
	for __obf_8004953ba26bf329 := range __obf_349df6ea515f02e1 {
		if __obf_8004953ba26bf329 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_b87bc669672328e4 = append(__obf_b87bc669672328e4, __obf_8004953ba26bf329)
	}
	sort.Strings(__obf_b87bc669672328e4) // 字典序排序

	var __obf_974eeae9b46c93f1 bytes.Buffer
	for _, __obf_8004953ba26bf329 := range __obf_b87bc669672328e4 {
		if __obf_349df6ea515f02e1[__obf_8004953ba26bf329] != "" {
			__obf_974eeae9b46c93f1. // 空值不参与签名
						WriteString(__obf_8004953ba26bf329)
			__obf_974eeae9b46c93f1.
				WriteString("=")
			__obf_974eeae9b46c93f1.
				WriteString(__obf_349df6ea515f02e1[__obf_8004953ba26bf329])
			__obf_974eeae9b46c93f1.
				WriteString("&")
		}
	}
	__obf_974eeae9b46c93f1.
		WriteString("key=")
	__obf_974eeae9b46c93f1.
		WriteString(__obf_434013c08d7e96ff)

	var __obf_983b33218441aeeb hash.Hash
	if __obf_349df6ea515f02e1["sign_type"] == SIGN_HMACSHA256 {
		__obf_983b33218441aeeb = hmac.New(sha256.New, []byte(__obf_434013c08d7e96ff))
	} else {
		__obf_983b33218441aeeb = md5.New()
	}
	__obf_983b33218441aeeb.
		Write(__obf_974eeae9b46c93f1.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_983b33218441aeeb.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_434013c08d7e96ff string, __obf_349df6ea515f02e1 map[string]string) bool {
	__obf_b55be4625826e956, __obf_e604613b3e6dc6cf := __obf_349df6ea515f02e1["sign"]
	if !__obf_e604613b3e6dc6cf {
		return false // 没有sign字段
	}
	delete(__obf_349df6ea515f02e1, "sign")
	__obf_7faa55095c72cfd3 := // 签名时需要移除sign字段

		CreateSign(__obf_434013c08d7e96ff, __obf_349df6ea515f02e1)

	return __obf_7faa55095c72cfd3 == __obf_b55be4625826e956
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_e8c254e95f1ffc9e map[string]string) []byte {
	var __obf_974eeae9b46c93f1 bytes.Buffer
	__obf_974eeae9b46c93f1.
		WriteString("<xml>")
	for __obf_8004953ba26bf329, __obf_2ad96ff371650fa9 := range __obf_e8c254e95f1ffc9e {
		__obf_974eeae9b46c93f1.
			WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_8004953ba26bf329, __obf_2ad96ff371650fa9, __obf_8004953ba26bf329))
	}
	__obf_974eeae9b46c93f1.
		WriteString("</xml>")
	return __obf_974eeae9b46c93f1.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_208b6368e1923c12 []byte) (map[string]string, error) {
	__obf_5cae921f379521aa := xml.NewDecoder(bytes.NewReader(__obf_208b6368e1923c12))
	__obf_e8c254e95f1ffc9e := make(map[string]string)
	for {
		__obf_3165fe535c548d0a, __obf_d6bcde9761346c61 := __obf_5cae921f379521aa.Token()
		if __obf_d6bcde9761346c61 != nil {
			if __obf_d6bcde9761346c61.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_d6bcde9761346c61)
		}

		switch __obf_f6e7ebb2b44b6445 := __obf_3165fe535c548d0a.(type) {
		case xml.StartElement:
			if __obf_f6e7ebb2b44b6445.Name.Local != "xml" { // 跳过根元素
				var __obf_7149453815d49dec string
				if __obf_d6bcde9761346c61 := __obf_5cae921f379521aa.DecodeElement(&__obf_7149453815d49dec, &__obf_f6e7ebb2b44b6445); __obf_d6bcde9761346c61 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_d6bcde9761346c61)
				}
				__obf_e8c254e95f1ffc9e[__obf_f6e7ebb2b44b6445.Name.Local] = __obf_7149453815d49dec
			}
		}
	}
	return __obf_e8c254e95f1ffc9e, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_1ae149d5839818f1, __obf_7c6444a58ba68f42 []byte, __obf_5e1666c18a4df0aa, __obf_143005d8e3eba02d string) (__obf_b62846ff51788094 []byte, __obf_d6bcde9761346c61 error) {
	defer func() {
		if __obf_eeaddc22e30f34f6 := recover(); __obf_eeaddc22e30f34f6 != nil {
			__obf_d6bcde9761346c61 = fmt.Errorf("panic error: err=%v", __obf_eeaddc22e30f34f6)
			return
		}
	}()
	var __obf_7dca2452dd625e20 []byte
	__obf_7dca2452dd625e20, __obf_d6bcde9761346c61 = __obf_ed31cc08c741f1cb(__obf_143005d8e3eba02d)
	if __obf_d6bcde9761346c61 != nil {
		panic(__obf_d6bcde9761346c61)
	}
	__obf_a7ede658d77da976 := AESEncryptMsg(__obf_1ae149d5839818f1, __obf_7c6444a58ba68f42, __obf_5e1666c18a4df0aa, __obf_7dca2452dd625e20)
	__obf_b62846ff51788094 = []byte(base64.StdEncoding.EncodeToString(__obf_a7ede658d77da976))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_1ae149d5839818f1, __obf_7c6444a58ba68f42 []byte, __obf_5e1666c18a4df0aa string, __obf_143005d8e3eba02d []byte) (__obf_a7ede658d77da976 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)
	__obf_1aa946ad14b68289 := 20 + len(__obf_7c6444a58ba68f42)
	__obf_cdf97f9ee643ab58 := __obf_1aa946ad14b68289 + len(__obf_5e1666c18a4df0aa)
	__obf_51bc49e015ec3843 := BlockSize - __obf_cdf97f9ee643ab58&BlockMask
	__obf_fc7b4df0f8413832 := __obf_cdf97f9ee643ab58 + __obf_51bc49e015ec3843
	__obf_d85e7a22f93decb3 := make([]byte, __obf_fc7b4df0f8413832)

	// 拼接
	copy(__obf_d85e7a22f93decb3[:16], __obf_1ae149d5839818f1)
	__obf_273989ebd12999e1(__obf_d85e7a22f93decb3[16:20], uint32(len(__obf_7c6444a58ba68f42)))
	copy(__obf_d85e7a22f93decb3[20:], __obf_7c6444a58ba68f42)
	copy(__obf_d85e7a22f93decb3[__obf_1aa946ad14b68289:], __obf_5e1666c18a4df0aa)

	// PKCS#7 补位
	for __obf_33d1c6b9ec98c730 := __obf_cdf97f9ee643ab58; __obf_33d1c6b9ec98c730 < __obf_fc7b4df0f8413832; __obf_33d1c6b9ec98c730++ {
		__obf_d85e7a22f93decb3[__obf_33d1c6b9ec98c730] = byte(__obf_51bc49e015ec3843)
	}
	__obf_78c51e8a1f2ea866,

		// 加密
		__obf_d6bcde9761346c61 := aes.NewCipher(__obf_143005d8e3eba02d)
	if __obf_d6bcde9761346c61 != nil {
		panic(__obf_d6bcde9761346c61)
	}
	__obf_60e008e2cef713de := cipher.NewCBCEncrypter(__obf_78c51e8a1f2ea866, __obf_143005d8e3eba02d[:16])
	__obf_60e008e2cef713de.
		CryptBlocks(__obf_d85e7a22f93decb3, __obf_d85e7a22f93decb3)
	__obf_a7ede658d77da976 = __obf_d85e7a22f93decb3
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_5e1666c18a4df0aa, __obf_6f800b635c41f634, __obf_143005d8e3eba02d string) (__obf_1ae149d5839818f1, __obf_a21cf7bb974d0655 []byte, __obf_d6bcde9761346c61 error) {
	defer func() {
		if __obf_eeaddc22e30f34f6 := recover(); __obf_eeaddc22e30f34f6 != nil {
			__obf_d6bcde9761346c61 = fmt.Errorf("panic error: err=%v", __obf_eeaddc22e30f34f6)
			return
		}
	}()
	var __obf_6efbc9be32122ace, __obf_7dca2452dd625e20, __obf_beef9095bf1d9edb []byte
	__obf_6efbc9be32122ace, __obf_d6bcde9761346c61 = base64.StdEncoding.DecodeString(__obf_6f800b635c41f634)
	if __obf_d6bcde9761346c61 != nil {
		return
	}
	__obf_7dca2452dd625e20, __obf_d6bcde9761346c61 = __obf_ed31cc08c741f1cb(__obf_143005d8e3eba02d)
	if __obf_d6bcde9761346c61 != nil {
		panic(__obf_d6bcde9761346c61)
	}
	__obf_1ae149d5839818f1, __obf_a21cf7bb974d0655, __obf_beef9095bf1d9edb, __obf_d6bcde9761346c61 = AESDecryptMsg(__obf_6efbc9be32122ace, __obf_7dca2452dd625e20)
	if __obf_d6bcde9761346c61 != nil {
		__obf_d6bcde9761346c61 = fmt.Errorf("消息解密失败,%v", __obf_d6bcde9761346c61)
		return
	}
	if __obf_5e1666c18a4df0aa != string(__obf_beef9095bf1d9edb) {
		__obf_d6bcde9761346c61 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_ed31cc08c741f1cb(__obf_724765d2e748d926 string) (__obf_7dca2452dd625e20 []byte, __obf_d6bcde9761346c61 error) {
	if len(__obf_724765d2e748d926) != 43 {
		__obf_d6bcde9761346c61 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_7dca2452dd625e20, __obf_d6bcde9761346c61 = base64.StdEncoding.DecodeString(__obf_724765d2e748d926 + "=")
	if __obf_d6bcde9761346c61 != nil {
		return
	}
	if len(__obf_7dca2452dd625e20) != 32 {
		__obf_d6bcde9761346c61 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_a7ede658d77da976 []byte, __obf_143005d8e3eba02d []byte) (__obf_1ae149d5839818f1, __obf_7c6444a58ba68f42, __obf_5e1666c18a4df0aa []byte, __obf_d6bcde9761346c61 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_a7ede658d77da976) < BlockSize {
		__obf_d6bcde9761346c61 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_a7ede658d77da976))
		return
	}
	if len(__obf_a7ede658d77da976)&BlockMask != 0 {
		__obf_d6bcde9761346c61 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_a7ede658d77da976))
		return
	}
	__obf_d85e7a22f93decb3 := make([]byte, len(__obf_a7ede658d77da976))
	__obf_78c51e8a1f2ea866, // len(plaintext) >= BLOCK_SIZE
		__obf_d6bcde9761346c61 := // 解密
		aes.NewCipher(__obf_143005d8e3eba02d)
	if __obf_d6bcde9761346c61 != nil {
		panic(__obf_d6bcde9761346c61)
	}
	__obf_60e008e2cef713de := cipher.NewCBCDecrypter(__obf_78c51e8a1f2ea866, __obf_143005d8e3eba02d[:16])
	__obf_60e008e2cef713de.
		CryptBlocks(__obf_d85e7a22f93decb3, __obf_a7ede658d77da976)
	__obf_51bc49e015ec3843 := // PKCS#7 去除补位
		int(__obf_d85e7a22f93decb3[len(__obf_d85e7a22f93decb3)-1])
	if __obf_51bc49e015ec3843 < 1 || __obf_51bc49e015ec3843 > BlockSize {
		__obf_d6bcde9761346c61 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_51bc49e015ec3843)
		return
	}
	__obf_d85e7a22f93decb3 = __obf_d85e7a22f93decb3[:len(__obf_d85e7a22f93decb3)-__obf_51bc49e015ec3843]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_d85e7a22f93decb3) <= 20 {
		__obf_d6bcde9761346c61 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_d85e7a22f93decb3))
		return
	}
	__obf_dfe342826161d3b1 := int(__obf_92ca0f252284dc7b(__obf_d85e7a22f93decb3[16:20]))
	if __obf_dfe342826161d3b1 < 0 {
		__obf_d6bcde9761346c61 = fmt.Errorf("incorrect msg length: %d", __obf_dfe342826161d3b1)
		return
	}
	__obf_1aa946ad14b68289 := 20 + __obf_dfe342826161d3b1
	if len(__obf_d85e7a22f93decb3) <= __obf_1aa946ad14b68289 {
		__obf_d6bcde9761346c61 = fmt.Errorf("msg length too large: %d", __obf_dfe342826161d3b1)
		return
	}
	__obf_1ae149d5839818f1 = __obf_d85e7a22f93decb3[:16:20]
	__obf_7c6444a58ba68f42 = __obf_d85e7a22f93decb3[20:__obf_1aa946ad14b68289:__obf_1aa946ad14b68289]
	__obf_5e1666c18a4df0aa = __obf_d85e7a22f93decb3[__obf_1aa946ad14b68289:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_273989ebd12999e1(__obf_5ccfbf1a565023f7 []byte, __obf_07ea9c75768fe319 uint32) {
	__obf_5ccfbf1a565023f7[0] = byte(__obf_07ea9c75768fe319 >> 24)
	__obf_5ccfbf1a565023f7[1] = byte(__obf_07ea9c75768fe319 >> 16)
	__obf_5ccfbf1a565023f7[2] = byte(__obf_07ea9c75768fe319 >> 8)
	__obf_5ccfbf1a565023f7[3] = byte(__obf_07ea9c75768fe319)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_92ca0f252284dc7b(__obf_5ccfbf1a565023f7 []byte) (__obf_07ea9c75768fe319 uint32) {
	return uint32(__obf_5ccfbf1a565023f7[0])<<24 |
		uint32(__obf_5ccfbf1a565023f7[1])<<16 |
		uint32(__obf_5ccfbf1a565023f7[2])<<8 |
		uint32(__obf_5ccfbf1a565023f7[3])
}
