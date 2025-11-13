package __obf_6e18fdc479ab921c

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
	__obf_71df7225c18a0382 := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_a58eb5d4b5a8098e := 32
	__obf_240c8d67cc789dde := make([]byte, __obf_a58eb5d4b5a8098e)
	for __obf_9e7a2284c4e7740a := range __obf_240c8d67cc789dde {
		__obf_240c8d67cc789dde[__obf_9e7a2284c4e7740a] = __obf_71df7225c18a0382[rand.IntN(len(__obf_71df7225c18a0382))]
	}
	return string(__obf_240c8d67cc789dde)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_7d05748024931c5b string, __obf_2c4cb6923ad79841 map[string]string) string {
	__obf_04c008f688dc4fd2 := make([]string, 0, len(__obf_2c4cb6923ad79841))
	for __obf_4dde8a7421899bf2 := range __obf_2c4cb6923ad79841 {
		if __obf_4dde8a7421899bf2 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_04c008f688dc4fd2 = append(__obf_04c008f688dc4fd2, __obf_4dde8a7421899bf2)
	}
	sort.Strings(__obf_04c008f688dc4fd2) // 字典序排序

	var __obf_ae1f1d82d5a47aa2 bytes.Buffer
	for _, __obf_4dde8a7421899bf2 := range __obf_04c008f688dc4fd2 {
		if __obf_2c4cb6923ad79841[__obf_4dde8a7421899bf2] != "" { // 空值不参与签名
			__obf_ae1f1d82d5a47aa2.WriteString(__obf_4dde8a7421899bf2)
			__obf_ae1f1d82d5a47aa2.WriteString("=")
			__obf_ae1f1d82d5a47aa2.WriteString(__obf_2c4cb6923ad79841[__obf_4dde8a7421899bf2])
			__obf_ae1f1d82d5a47aa2.WriteString("&")
		}
	}
	__obf_ae1f1d82d5a47aa2.WriteString("key=")
	__obf_ae1f1d82d5a47aa2.WriteString(__obf_7d05748024931c5b)

	var __obf_17169145900b4209 hash.Hash
	if __obf_2c4cb6923ad79841["sign_type"] == SIGN_HMACSHA256 {
		__obf_17169145900b4209 = hmac.New(sha256.New, []byte(__obf_7d05748024931c5b))
	} else {
		__obf_17169145900b4209 = md5.New()
	}
	__obf_17169145900b4209.Write(__obf_ae1f1d82d5a47aa2.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_17169145900b4209.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_7d05748024931c5b string, __obf_2c4cb6923ad79841 map[string]string) bool {
	__obf_85961a7534ad6499, __obf_4b0213d3f8492406 := __obf_2c4cb6923ad79841["sign"]
	if !__obf_4b0213d3f8492406 {
		return false // 没有sign字段
	}
	delete(__obf_2c4cb6923ad79841, "sign") // 签名时需要移除sign字段

	__obf_591bb6088f59f3cd := CreateSign(__obf_7d05748024931c5b, __obf_2c4cb6923ad79841)

	return __obf_591bb6088f59f3cd == __obf_85961a7534ad6499
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_a371d1762a30244f map[string]string) []byte {
	var __obf_ae1f1d82d5a47aa2 bytes.Buffer
	__obf_ae1f1d82d5a47aa2.WriteString("<xml>")
	for __obf_4dde8a7421899bf2, __obf_de9cd12b999685f6 := range __obf_a371d1762a30244f {
		__obf_ae1f1d82d5a47aa2.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_4dde8a7421899bf2, __obf_de9cd12b999685f6, __obf_4dde8a7421899bf2))
	}
	__obf_ae1f1d82d5a47aa2.WriteString("</xml>")
	return __obf_ae1f1d82d5a47aa2.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_1723697995a8daa8 []byte) (map[string]string, error) {
	__obf_0b6edda211c6f9b9 := xml.NewDecoder(bytes.NewReader(__obf_1723697995a8daa8))
	__obf_a371d1762a30244f := make(map[string]string)
	for {
		__obf_22df8ae893454b92, __obf_d92b02392b1fb408 := __obf_0b6edda211c6f9b9.Token()
		if __obf_d92b02392b1fb408 != nil {
			if __obf_d92b02392b1fb408.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_d92b02392b1fb408)
		}

		switch __obf_22c755954d3bab17 := __obf_22df8ae893454b92.(type) {
		case xml.StartElement:
			if __obf_22c755954d3bab17.Name.Local != "xml" { // 跳过根元素
				var __obf_baf3d7341c144032 string
				if __obf_d92b02392b1fb408 := __obf_0b6edda211c6f9b9.DecodeElement(&__obf_baf3d7341c144032, &__obf_22c755954d3bab17); __obf_d92b02392b1fb408 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_d92b02392b1fb408)
				}
				__obf_a371d1762a30244f[__obf_22c755954d3bab17.Name.Local] = __obf_baf3d7341c144032
			}
		}
	}
	return __obf_a371d1762a30244f, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_a0452422924e72f1, __obf_465421ac22acd5f8 []byte, __obf_10234e8a0ea34853, __obf_527d58298a2fb588 string) (__obf_4c9fbd9669d00b4e []byte, __obf_d92b02392b1fb408 error) {
	defer func() {
		if __obf_aee3e27789fcdec8 := recover(); __obf_aee3e27789fcdec8 != nil {
			__obf_d92b02392b1fb408 = fmt.Errorf("panic error: err=%v", __obf_aee3e27789fcdec8)
			return
		}
	}()
	var __obf_57cc787ea702af69 []byte
	__obf_57cc787ea702af69, __obf_d92b02392b1fb408 = __obf_f4151043fbaf8b28(__obf_527d58298a2fb588)
	if __obf_d92b02392b1fb408 != nil {
		panic(__obf_d92b02392b1fb408)
	}
	__obf_f690c0713fc53061 := AESEncryptMsg(__obf_a0452422924e72f1, __obf_465421ac22acd5f8, __obf_10234e8a0ea34853, __obf_57cc787ea702af69)
	__obf_4c9fbd9669d00b4e = []byte(base64.StdEncoding.EncodeToString(__obf_f690c0713fc53061))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_a0452422924e72f1, __obf_465421ac22acd5f8 []byte, __obf_10234e8a0ea34853 string, __obf_527d58298a2fb588 []byte) (__obf_f690c0713fc53061 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	__obf_ac32f5e99d0c6dc5 := 20 + len(__obf_465421ac22acd5f8)
	__obf_9a2704a151a38811 := __obf_ac32f5e99d0c6dc5 + len(__obf_10234e8a0ea34853)
	__obf_a24467d27904cefe := BlockSize - __obf_9a2704a151a38811&BlockMask
	__obf_590301beea4edb07 := __obf_9a2704a151a38811 + __obf_a24467d27904cefe

	__obf_64b611f81f3d62be := make([]byte, __obf_590301beea4edb07)

	// 拼接
	copy(__obf_64b611f81f3d62be[:16], __obf_a0452422924e72f1)
	__obf_7c935821c746316a(__obf_64b611f81f3d62be[16:20], uint32(len(__obf_465421ac22acd5f8)))
	copy(__obf_64b611f81f3d62be[20:], __obf_465421ac22acd5f8)
	copy(__obf_64b611f81f3d62be[__obf_ac32f5e99d0c6dc5:], __obf_10234e8a0ea34853)

	// PKCS#7 补位
	for __obf_9e7a2284c4e7740a := __obf_9a2704a151a38811; __obf_9e7a2284c4e7740a < __obf_590301beea4edb07; __obf_9e7a2284c4e7740a++ {
		__obf_64b611f81f3d62be[__obf_9e7a2284c4e7740a] = byte(__obf_a24467d27904cefe)
	}

	// 加密
	__obf_237cbb22504fcae4, __obf_d92b02392b1fb408 := aes.NewCipher(__obf_527d58298a2fb588)
	if __obf_d92b02392b1fb408 != nil {
		panic(__obf_d92b02392b1fb408)
	}
	__obf_beef76cffdb7cc34 := cipher.NewCBCEncrypter(__obf_237cbb22504fcae4, __obf_527d58298a2fb588[:16])
	__obf_beef76cffdb7cc34.CryptBlocks(__obf_64b611f81f3d62be, __obf_64b611f81f3d62be)

	__obf_f690c0713fc53061 = __obf_64b611f81f3d62be
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_10234e8a0ea34853, __obf_3c9e27606a22482f, __obf_527d58298a2fb588 string) (__obf_a0452422924e72f1, __obf_cc585699d3481965 []byte, __obf_d92b02392b1fb408 error) {
	defer func() {
		if __obf_aee3e27789fcdec8 := recover(); __obf_aee3e27789fcdec8 != nil {
			__obf_d92b02392b1fb408 = fmt.Errorf("panic error: err=%v", __obf_aee3e27789fcdec8)
			return
		}
	}()
	var __obf_7587a8ac8b686de5, __obf_57cc787ea702af69, __obf_9706d833e1488807 []byte
	__obf_7587a8ac8b686de5, __obf_d92b02392b1fb408 = base64.StdEncoding.DecodeString(__obf_3c9e27606a22482f)
	if __obf_d92b02392b1fb408 != nil {
		return
	}
	__obf_57cc787ea702af69, __obf_d92b02392b1fb408 = __obf_f4151043fbaf8b28(__obf_527d58298a2fb588)
	if __obf_d92b02392b1fb408 != nil {
		panic(__obf_d92b02392b1fb408)
	}
	__obf_a0452422924e72f1, __obf_cc585699d3481965, __obf_9706d833e1488807, __obf_d92b02392b1fb408 = AESDecryptMsg(__obf_7587a8ac8b686de5, __obf_57cc787ea702af69)
	if __obf_d92b02392b1fb408 != nil {
		__obf_d92b02392b1fb408 = fmt.Errorf("消息解密失败,%v", __obf_d92b02392b1fb408)
		return
	}
	if __obf_10234e8a0ea34853 != string(__obf_9706d833e1488807) {
		__obf_d92b02392b1fb408 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_f4151043fbaf8b28(__obf_16f1f49f4f751540 string) (__obf_57cc787ea702af69 []byte, __obf_d92b02392b1fb408 error) {
	if len(__obf_16f1f49f4f751540) != 43 {
		__obf_d92b02392b1fb408 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_57cc787ea702af69, __obf_d92b02392b1fb408 = base64.StdEncoding.DecodeString(__obf_16f1f49f4f751540 + "=")
	if __obf_d92b02392b1fb408 != nil {
		return
	}
	if len(__obf_57cc787ea702af69) != 32 {
		__obf_d92b02392b1fb408 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_f690c0713fc53061 []byte, __obf_527d58298a2fb588 []byte) (__obf_a0452422924e72f1, __obf_465421ac22acd5f8, __obf_10234e8a0ea34853 []byte, __obf_d92b02392b1fb408 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_f690c0713fc53061) < BlockSize {
		__obf_d92b02392b1fb408 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_f690c0713fc53061))
		return
	}
	if len(__obf_f690c0713fc53061)&BlockMask != 0 {
		__obf_d92b02392b1fb408 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_f690c0713fc53061))
		return
	}

	__obf_64b611f81f3d62be := make([]byte, len(__obf_f690c0713fc53061)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	__obf_237cbb22504fcae4, __obf_d92b02392b1fb408 := aes.NewCipher(__obf_527d58298a2fb588)
	if __obf_d92b02392b1fb408 != nil {
		panic(__obf_d92b02392b1fb408)
	}
	__obf_beef76cffdb7cc34 := cipher.NewCBCDecrypter(__obf_237cbb22504fcae4, __obf_527d58298a2fb588[:16])
	__obf_beef76cffdb7cc34.CryptBlocks(__obf_64b611f81f3d62be, __obf_f690c0713fc53061)

	// PKCS#7 去除补位
	__obf_a24467d27904cefe := int(__obf_64b611f81f3d62be[len(__obf_64b611f81f3d62be)-1])
	if __obf_a24467d27904cefe < 1 || __obf_a24467d27904cefe > BlockSize {
		__obf_d92b02392b1fb408 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_a24467d27904cefe)
		return
	}
	__obf_64b611f81f3d62be = __obf_64b611f81f3d62be[:len(__obf_64b611f81f3d62be)-__obf_a24467d27904cefe]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_64b611f81f3d62be) <= 20 {
		__obf_d92b02392b1fb408 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_64b611f81f3d62be))
		return
	}
	__obf_a84a63ca7f075809 := int(__obf_9aa8f8364c6bb300(__obf_64b611f81f3d62be[16:20]))
	if __obf_a84a63ca7f075809 < 0 {
		__obf_d92b02392b1fb408 = fmt.Errorf("incorrect msg length: %d", __obf_a84a63ca7f075809)
		return
	}
	__obf_ac32f5e99d0c6dc5 := 20 + __obf_a84a63ca7f075809
	if len(__obf_64b611f81f3d62be) <= __obf_ac32f5e99d0c6dc5 {
		__obf_d92b02392b1fb408 = fmt.Errorf("msg length too large: %d", __obf_a84a63ca7f075809)
		return
	}

	__obf_a0452422924e72f1 = __obf_64b611f81f3d62be[:16:20]
	__obf_465421ac22acd5f8 = __obf_64b611f81f3d62be[20:__obf_ac32f5e99d0c6dc5:__obf_ac32f5e99d0c6dc5]
	__obf_10234e8a0ea34853 = __obf_64b611f81f3d62be[__obf_ac32f5e99d0c6dc5:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_7c935821c746316a(__obf_3d7ae7058bcc335e []byte, __obf_0fa77f03015f34b2 uint32) {
	__obf_3d7ae7058bcc335e[0] = byte(__obf_0fa77f03015f34b2 >> 24)
	__obf_3d7ae7058bcc335e[1] = byte(__obf_0fa77f03015f34b2 >> 16)
	__obf_3d7ae7058bcc335e[2] = byte(__obf_0fa77f03015f34b2 >> 8)
	__obf_3d7ae7058bcc335e[3] = byte(__obf_0fa77f03015f34b2)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_9aa8f8364c6bb300(__obf_3d7ae7058bcc335e []byte) (__obf_0fa77f03015f34b2 uint32) {
	return uint32(__obf_3d7ae7058bcc335e[0])<<24 |
		uint32(__obf_3d7ae7058bcc335e[1])<<16 |
		uint32(__obf_3d7ae7058bcc335e[2])<<8 |
		uint32(__obf_3d7ae7058bcc335e[3])
}
