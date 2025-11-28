package __obf_c0d7bb2e04898f29

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
	__obf_e609852fb000c62c := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_c5ab5df2052940ae := 32
	__obf_b5530a24779d84f1 := make([]byte, __obf_c5ab5df2052940ae)
	for __obf_8a51ac22f0a1ffb6 := range __obf_b5530a24779d84f1 {
		__obf_b5530a24779d84f1[__obf_8a51ac22f0a1ffb6] = __obf_e609852fb000c62c[rand.IntN(len(__obf_e609852fb000c62c))]
	}
	return string(__obf_b5530a24779d84f1)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_a961c6319d695f3c string, __obf_e29bd3e6054f1250 map[string]string) string {
	__obf_6deafab9a7d6c818 := make([]string, 0, len(__obf_e29bd3e6054f1250))
	for __obf_f8faaa11642d2945 := range __obf_e29bd3e6054f1250 {
		if __obf_f8faaa11642d2945 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_6deafab9a7d6c818 = append(__obf_6deafab9a7d6c818, __obf_f8faaa11642d2945)
	}
	sort.Strings(__obf_6deafab9a7d6c818) // 字典序排序

	var __obf_3a8531fbac8cf56e bytes.Buffer
	for _, __obf_f8faaa11642d2945 := range __obf_6deafab9a7d6c818 {
		if __obf_e29bd3e6054f1250[__obf_f8faaa11642d2945] != "" { // 空值不参与签名
			__obf_3a8531fbac8cf56e.WriteString(__obf_f8faaa11642d2945)
			__obf_3a8531fbac8cf56e.WriteString("=")
			__obf_3a8531fbac8cf56e.WriteString(__obf_e29bd3e6054f1250[__obf_f8faaa11642d2945])
			__obf_3a8531fbac8cf56e.WriteString("&")
		}
	}
	__obf_3a8531fbac8cf56e.WriteString("key=")
	__obf_3a8531fbac8cf56e.WriteString(__obf_a961c6319d695f3c)

	var __obf_c4604846c415b656 hash.Hash
	if __obf_e29bd3e6054f1250["sign_type"] == SIGN_HMACSHA256 {
		__obf_c4604846c415b656 = hmac.New(sha256.New, []byte(__obf_a961c6319d695f3c))
	} else {
		__obf_c4604846c415b656 = md5.New()
	}
	__obf_c4604846c415b656.Write(__obf_3a8531fbac8cf56e.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_c4604846c415b656.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_a961c6319d695f3c string, __obf_e29bd3e6054f1250 map[string]string) bool {
	__obf_8053d47763784db0, __obf_8e75ea473cdc6ad4 := __obf_e29bd3e6054f1250["sign"]
	if !__obf_8e75ea473cdc6ad4 {
		return false // 没有sign字段
	}
	delete(__obf_e29bd3e6054f1250, "sign") // 签名时需要移除sign字段

	__obf_399d2578207e06ad := CreateSign(__obf_a961c6319d695f3c, __obf_e29bd3e6054f1250)

	return __obf_399d2578207e06ad == __obf_8053d47763784db0
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_c313e5576bf980c7 map[string]string) []byte {
	var __obf_3a8531fbac8cf56e bytes.Buffer
	__obf_3a8531fbac8cf56e.WriteString("<xml>")
	for __obf_f8faaa11642d2945, __obf_1a8fb20ade71b8f0 := range __obf_c313e5576bf980c7 {
		__obf_3a8531fbac8cf56e.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_f8faaa11642d2945, __obf_1a8fb20ade71b8f0, __obf_f8faaa11642d2945))
	}
	__obf_3a8531fbac8cf56e.WriteString("</xml>")
	return __obf_3a8531fbac8cf56e.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_9bcf1fd50995350e []byte) (map[string]string, error) {
	__obf_8a58a21a0a5139e0 := xml.NewDecoder(bytes.NewReader(__obf_9bcf1fd50995350e))
	__obf_c313e5576bf980c7 := make(map[string]string)
	for {
		__obf_7857c58e232bd882, __obf_fb6d611ef291e586 := __obf_8a58a21a0a5139e0.Token()
		if __obf_fb6d611ef291e586 != nil {
			if __obf_fb6d611ef291e586.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_fb6d611ef291e586)
		}

		switch __obf_4a3c65f95ced4714 := __obf_7857c58e232bd882.(type) {
		case xml.StartElement:
			if __obf_4a3c65f95ced4714.Name.Local != "xml" { // 跳过根元素
				var __obf_157310ad2d16b6c8 string
				if __obf_fb6d611ef291e586 := __obf_8a58a21a0a5139e0.DecodeElement(&__obf_157310ad2d16b6c8, &__obf_4a3c65f95ced4714); __obf_fb6d611ef291e586 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_fb6d611ef291e586)
				}
				__obf_c313e5576bf980c7[__obf_4a3c65f95ced4714.Name.Local] = __obf_157310ad2d16b6c8
			}
		}
	}
	return __obf_c313e5576bf980c7, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_97b7c2c38e5bdfb9, __obf_803bf8ce7e3f6629 []byte, __obf_cf3c3f5d338025b1, __obf_c2694a71f9acb426 string) (__obf_9b0944562f9fa3ee []byte, __obf_fb6d611ef291e586 error) {
	defer func() {
		if __obf_9190d7d935dc17a0 := recover(); __obf_9190d7d935dc17a0 != nil {
			__obf_fb6d611ef291e586 = fmt.Errorf("panic error: err=%v", __obf_9190d7d935dc17a0)
			return
		}
	}()
	var __obf_9732b72b1933d195 []byte
	__obf_9732b72b1933d195, __obf_fb6d611ef291e586 = __obf_8dbd1b8a2982a387(__obf_c2694a71f9acb426)
	if __obf_fb6d611ef291e586 != nil {
		panic(__obf_fb6d611ef291e586)
	}
	__obf_2351672914163ab0 := AESEncryptMsg(__obf_97b7c2c38e5bdfb9, __obf_803bf8ce7e3f6629, __obf_cf3c3f5d338025b1, __obf_9732b72b1933d195)
	__obf_9b0944562f9fa3ee = []byte(base64.StdEncoding.EncodeToString(__obf_2351672914163ab0))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_97b7c2c38e5bdfb9, __obf_803bf8ce7e3f6629 []byte, __obf_cf3c3f5d338025b1 string, __obf_c2694a71f9acb426 []byte) (__obf_2351672914163ab0 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	__obf_96c00645dfe9e672 := 20 + len(__obf_803bf8ce7e3f6629)
	__obf_10ad211e77b7591c := __obf_96c00645dfe9e672 + len(__obf_cf3c3f5d338025b1)
	__obf_2b2d34e992ffeb98 := BlockSize - __obf_10ad211e77b7591c&BlockMask
	__obf_f95c04104695826f := __obf_10ad211e77b7591c + __obf_2b2d34e992ffeb98

	__obf_1db286939c7e8606 := make([]byte, __obf_f95c04104695826f)

	// 拼接
	copy(__obf_1db286939c7e8606[:16], __obf_97b7c2c38e5bdfb9)
	__obf_60b408581e41ac79(__obf_1db286939c7e8606[16:20], uint32(len(__obf_803bf8ce7e3f6629)))
	copy(__obf_1db286939c7e8606[20:], __obf_803bf8ce7e3f6629)
	copy(__obf_1db286939c7e8606[__obf_96c00645dfe9e672:], __obf_cf3c3f5d338025b1)

	// PKCS#7 补位
	for __obf_8a51ac22f0a1ffb6 := __obf_10ad211e77b7591c; __obf_8a51ac22f0a1ffb6 < __obf_f95c04104695826f; __obf_8a51ac22f0a1ffb6++ {
		__obf_1db286939c7e8606[__obf_8a51ac22f0a1ffb6] = byte(__obf_2b2d34e992ffeb98)
	}

	// 加密
	__obf_255c57c66bf4666e, __obf_fb6d611ef291e586 := aes.NewCipher(__obf_c2694a71f9acb426)
	if __obf_fb6d611ef291e586 != nil {
		panic(__obf_fb6d611ef291e586)
	}
	__obf_51994b5b3978a3ad := cipher.NewCBCEncrypter(__obf_255c57c66bf4666e, __obf_c2694a71f9acb426[:16])
	__obf_51994b5b3978a3ad.CryptBlocks(__obf_1db286939c7e8606, __obf_1db286939c7e8606)

	__obf_2351672914163ab0 = __obf_1db286939c7e8606
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_cf3c3f5d338025b1, __obf_817497bde43f8905, __obf_c2694a71f9acb426 string) (__obf_97b7c2c38e5bdfb9, __obf_b60b3ac0ede811b3 []byte, __obf_fb6d611ef291e586 error) {
	defer func() {
		if __obf_9190d7d935dc17a0 := recover(); __obf_9190d7d935dc17a0 != nil {
			__obf_fb6d611ef291e586 = fmt.Errorf("panic error: err=%v", __obf_9190d7d935dc17a0)
			return
		}
	}()
	var __obf_f4fb903e6c32d97d, __obf_9732b72b1933d195, __obf_26793de608635592 []byte
	__obf_f4fb903e6c32d97d, __obf_fb6d611ef291e586 = base64.StdEncoding.DecodeString(__obf_817497bde43f8905)
	if __obf_fb6d611ef291e586 != nil {
		return
	}
	__obf_9732b72b1933d195, __obf_fb6d611ef291e586 = __obf_8dbd1b8a2982a387(__obf_c2694a71f9acb426)
	if __obf_fb6d611ef291e586 != nil {
		panic(__obf_fb6d611ef291e586)
	}
	__obf_97b7c2c38e5bdfb9, __obf_b60b3ac0ede811b3, __obf_26793de608635592, __obf_fb6d611ef291e586 = AESDecryptMsg(__obf_f4fb903e6c32d97d, __obf_9732b72b1933d195)
	if __obf_fb6d611ef291e586 != nil {
		__obf_fb6d611ef291e586 = fmt.Errorf("消息解密失败,%v", __obf_fb6d611ef291e586)
		return
	}
	if __obf_cf3c3f5d338025b1 != string(__obf_26793de608635592) {
		__obf_fb6d611ef291e586 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_8dbd1b8a2982a387(__obf_827bfe1dc832ddc4 string) (__obf_9732b72b1933d195 []byte, __obf_fb6d611ef291e586 error) {
	if len(__obf_827bfe1dc832ddc4) != 43 {
		__obf_fb6d611ef291e586 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_9732b72b1933d195, __obf_fb6d611ef291e586 = base64.StdEncoding.DecodeString(__obf_827bfe1dc832ddc4 + "=")
	if __obf_fb6d611ef291e586 != nil {
		return
	}
	if len(__obf_9732b72b1933d195) != 32 {
		__obf_fb6d611ef291e586 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_2351672914163ab0 []byte, __obf_c2694a71f9acb426 []byte) (__obf_97b7c2c38e5bdfb9, __obf_803bf8ce7e3f6629, __obf_cf3c3f5d338025b1 []byte, __obf_fb6d611ef291e586 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_2351672914163ab0) < BlockSize {
		__obf_fb6d611ef291e586 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_2351672914163ab0))
		return
	}
	if len(__obf_2351672914163ab0)&BlockMask != 0 {
		__obf_fb6d611ef291e586 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_2351672914163ab0))
		return
	}

	__obf_1db286939c7e8606 := make([]byte, len(__obf_2351672914163ab0)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	__obf_255c57c66bf4666e, __obf_fb6d611ef291e586 := aes.NewCipher(__obf_c2694a71f9acb426)
	if __obf_fb6d611ef291e586 != nil {
		panic(__obf_fb6d611ef291e586)
	}
	__obf_51994b5b3978a3ad := cipher.NewCBCDecrypter(__obf_255c57c66bf4666e, __obf_c2694a71f9acb426[:16])
	__obf_51994b5b3978a3ad.CryptBlocks(__obf_1db286939c7e8606, __obf_2351672914163ab0)

	// PKCS#7 去除补位
	__obf_2b2d34e992ffeb98 := int(__obf_1db286939c7e8606[len(__obf_1db286939c7e8606)-1])
	if __obf_2b2d34e992ffeb98 < 1 || __obf_2b2d34e992ffeb98 > BlockSize {
		__obf_fb6d611ef291e586 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_2b2d34e992ffeb98)
		return
	}
	__obf_1db286939c7e8606 = __obf_1db286939c7e8606[:len(__obf_1db286939c7e8606)-__obf_2b2d34e992ffeb98]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_1db286939c7e8606) <= 20 {
		__obf_fb6d611ef291e586 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_1db286939c7e8606))
		return
	}
	__obf_b735571c300e9d24 := int(__obf_1606ac822abf6ceb(__obf_1db286939c7e8606[16:20]))
	if __obf_b735571c300e9d24 < 0 {
		__obf_fb6d611ef291e586 = fmt.Errorf("incorrect msg length: %d", __obf_b735571c300e9d24)
		return
	}
	__obf_96c00645dfe9e672 := 20 + __obf_b735571c300e9d24
	if len(__obf_1db286939c7e8606) <= __obf_96c00645dfe9e672 {
		__obf_fb6d611ef291e586 = fmt.Errorf("msg length too large: %d", __obf_b735571c300e9d24)
		return
	}

	__obf_97b7c2c38e5bdfb9 = __obf_1db286939c7e8606[:16:20]
	__obf_803bf8ce7e3f6629 = __obf_1db286939c7e8606[20:__obf_96c00645dfe9e672:__obf_96c00645dfe9e672]
	__obf_cf3c3f5d338025b1 = __obf_1db286939c7e8606[__obf_96c00645dfe9e672:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_60b408581e41ac79(__obf_b14b9697fec51afd []byte, __obf_d6f9c198ee46cb80 uint32) {
	__obf_b14b9697fec51afd[0] = byte(__obf_d6f9c198ee46cb80 >> 24)
	__obf_b14b9697fec51afd[1] = byte(__obf_d6f9c198ee46cb80 >> 16)
	__obf_b14b9697fec51afd[2] = byte(__obf_d6f9c198ee46cb80 >> 8)
	__obf_b14b9697fec51afd[3] = byte(__obf_d6f9c198ee46cb80)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_1606ac822abf6ceb(__obf_b14b9697fec51afd []byte) (__obf_d6f9c198ee46cb80 uint32) {
	return uint32(__obf_b14b9697fec51afd[0])<<24 |
		uint32(__obf_b14b9697fec51afd[1])<<16 |
		uint32(__obf_b14b9697fec51afd[2])<<8 |
		uint32(__obf_b14b9697fec51afd[3])
}
