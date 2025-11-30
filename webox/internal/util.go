package __obf_b5bcac367b722f8a

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
	__obf_28feb4fcd3d138af := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_8ff1fddf399d9d1a := 32
	__obf_cebdb062df719721 := make([]byte, __obf_8ff1fddf399d9d1a)
	for __obf_b212faf1ba253a88 := range __obf_cebdb062df719721 {
		__obf_cebdb062df719721[__obf_b212faf1ba253a88] = __obf_28feb4fcd3d138af[rand.IntN(len(__obf_28feb4fcd3d138af))]
	}
	return string(__obf_cebdb062df719721)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_5612d471f3d22e1e string, __obf_6ef442f1c29fd43f map[string]string) string {
	__obf_314a3bf869802343 := make([]string, 0, len(__obf_6ef442f1c29fd43f))
	for __obf_a0009a4f0921242e := range __obf_6ef442f1c29fd43f {
		if __obf_a0009a4f0921242e == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_314a3bf869802343 = append(__obf_314a3bf869802343, __obf_a0009a4f0921242e)
	}
	sort.Strings(__obf_314a3bf869802343) // 字典序排序

	var __obf_22c2c589af75bceb bytes.Buffer
	for _, __obf_a0009a4f0921242e := range __obf_314a3bf869802343 {
		if __obf_6ef442f1c29fd43f[__obf_a0009a4f0921242e] != "" {
			__obf_22c2c589af75bceb. // 空值不参与签名
						WriteString(__obf_a0009a4f0921242e)
			__obf_22c2c589af75bceb.
				WriteString("=")
			__obf_22c2c589af75bceb.
				WriteString(__obf_6ef442f1c29fd43f[__obf_a0009a4f0921242e])
			__obf_22c2c589af75bceb.
				WriteString("&")
		}
	}
	__obf_22c2c589af75bceb.
		WriteString("key=")
	__obf_22c2c589af75bceb.
		WriteString(__obf_5612d471f3d22e1e)

	var __obf_b04e5ea504c75d18 hash.Hash
	if __obf_6ef442f1c29fd43f["sign_type"] == SIGN_HMACSHA256 {
		__obf_b04e5ea504c75d18 = hmac.New(sha256.New, []byte(__obf_5612d471f3d22e1e))
	} else {
		__obf_b04e5ea504c75d18 = md5.New()
	}
	__obf_b04e5ea504c75d18.
		Write(__obf_22c2c589af75bceb.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_b04e5ea504c75d18.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_5612d471f3d22e1e string, __obf_6ef442f1c29fd43f map[string]string) bool {
	__obf_ec60dc0d7e292fde, __obf_da7823c54aed28c0 := __obf_6ef442f1c29fd43f["sign"]
	if !__obf_da7823c54aed28c0 {
		return false // 没有sign字段
	}
	delete(__obf_6ef442f1c29fd43f, "sign")
	__obf_315dc1c0407720c2 := // 签名时需要移除sign字段

		CreateSign(__obf_5612d471f3d22e1e, __obf_6ef442f1c29fd43f)

	return __obf_315dc1c0407720c2 == __obf_ec60dc0d7e292fde
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_814ec131f39316bf map[string]string) []byte {
	var __obf_22c2c589af75bceb bytes.Buffer
	__obf_22c2c589af75bceb.
		WriteString("<xml>")
	for __obf_a0009a4f0921242e, __obf_70d42ebd71a0d3b2 := range __obf_814ec131f39316bf {
		__obf_22c2c589af75bceb.
			WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_a0009a4f0921242e, __obf_70d42ebd71a0d3b2, __obf_a0009a4f0921242e))
	}
	__obf_22c2c589af75bceb.
		WriteString("</xml>")
	return __obf_22c2c589af75bceb.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_b216b34a283d6b83 []byte) (map[string]string, error) {
	__obf_4a6b7df04c906313 := xml.NewDecoder(bytes.NewReader(__obf_b216b34a283d6b83))
	__obf_814ec131f39316bf := make(map[string]string)
	for {
		__obf_5b7fb132e8734819, __obf_d9b19f08bff09faf := __obf_4a6b7df04c906313.Token()
		if __obf_d9b19f08bff09faf != nil {
			if __obf_d9b19f08bff09faf.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_d9b19f08bff09faf)
		}

		switch __obf_8ffe89eec17a14ed := __obf_5b7fb132e8734819.(type) {
		case xml.StartElement:
			if __obf_8ffe89eec17a14ed.Name.Local != "xml" { // 跳过根元素
				var __obf_69fa2f0b289189dd string
				if __obf_d9b19f08bff09faf := __obf_4a6b7df04c906313.DecodeElement(&__obf_69fa2f0b289189dd, &__obf_8ffe89eec17a14ed); __obf_d9b19f08bff09faf != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_d9b19f08bff09faf)
				}
				__obf_814ec131f39316bf[__obf_8ffe89eec17a14ed.Name.Local] = __obf_69fa2f0b289189dd
			}
		}
	}
	return __obf_814ec131f39316bf, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_1a8b597c2a6b451c, __obf_e8aac65c3b092c48 []byte, __obf_130053b95e1ae28e, __obf_819b883bbd4b7b1a string) (__obf_f02230eaa64db880 []byte, __obf_d9b19f08bff09faf error) {
	defer func() {
		if __obf_a21318954cdf63e3 := recover(); __obf_a21318954cdf63e3 != nil {
			__obf_d9b19f08bff09faf = fmt.Errorf("panic error: err=%v", __obf_a21318954cdf63e3)
			return
		}
	}()
	var __obf_eb194fea2e30bbb4 []byte
	__obf_eb194fea2e30bbb4, __obf_d9b19f08bff09faf = __obf_d9b3d1b65722fad4(__obf_819b883bbd4b7b1a)
	if __obf_d9b19f08bff09faf != nil {
		panic(__obf_d9b19f08bff09faf)
	}
	__obf_2e353653f52b4b5b := AESEncryptMsg(__obf_1a8b597c2a6b451c, __obf_e8aac65c3b092c48, __obf_130053b95e1ae28e, __obf_eb194fea2e30bbb4)
	__obf_f02230eaa64db880 = []byte(base64.StdEncoding.EncodeToString(__obf_2e353653f52b4b5b))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_1a8b597c2a6b451c, __obf_e8aac65c3b092c48 []byte, __obf_130053b95e1ae28e string, __obf_819b883bbd4b7b1a []byte) (__obf_2e353653f52b4b5b []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)
	__obf_40ef524e717b314b := 20 + len(__obf_e8aac65c3b092c48)
	__obf_f535496453a57f84 := __obf_40ef524e717b314b + len(__obf_130053b95e1ae28e)
	__obf_eb5031eef93a603d := BlockSize - __obf_f535496453a57f84&BlockMask
	__obf_6e882e0ee1f63edf := __obf_f535496453a57f84 + __obf_eb5031eef93a603d
	__obf_9988ac51e603ec89 := make([]byte, __obf_6e882e0ee1f63edf)

	// 拼接
	copy(__obf_9988ac51e603ec89[:16], __obf_1a8b597c2a6b451c)
	__obf_55604a055b096813(__obf_9988ac51e603ec89[16:20], uint32(len(__obf_e8aac65c3b092c48)))
	copy(__obf_9988ac51e603ec89[20:], __obf_e8aac65c3b092c48)
	copy(__obf_9988ac51e603ec89[__obf_40ef524e717b314b:], __obf_130053b95e1ae28e)

	// PKCS#7 补位
	for __obf_b212faf1ba253a88 := __obf_f535496453a57f84; __obf_b212faf1ba253a88 < __obf_6e882e0ee1f63edf; __obf_b212faf1ba253a88++ {
		__obf_9988ac51e603ec89[__obf_b212faf1ba253a88] = byte(__obf_eb5031eef93a603d)
	}
	__obf_850a472cd2c60962,

		// 加密
		__obf_d9b19f08bff09faf := aes.NewCipher(__obf_819b883bbd4b7b1a)
	if __obf_d9b19f08bff09faf != nil {
		panic(__obf_d9b19f08bff09faf)
	}
	__obf_79c31f2cb1529b76 := cipher.NewCBCEncrypter(__obf_850a472cd2c60962, __obf_819b883bbd4b7b1a[:16])
	__obf_79c31f2cb1529b76.
		CryptBlocks(__obf_9988ac51e603ec89, __obf_9988ac51e603ec89)
	__obf_2e353653f52b4b5b = __obf_9988ac51e603ec89
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_130053b95e1ae28e, __obf_2e8adb82acaa0ab7, __obf_819b883bbd4b7b1a string) (__obf_1a8b597c2a6b451c, __obf_0bae1112d834471f []byte, __obf_d9b19f08bff09faf error) {
	defer func() {
		if __obf_a21318954cdf63e3 := recover(); __obf_a21318954cdf63e3 != nil {
			__obf_d9b19f08bff09faf = fmt.Errorf("panic error: err=%v", __obf_a21318954cdf63e3)
			return
		}
	}()
	var __obf_83403a880bc1a9b9, __obf_eb194fea2e30bbb4, __obf_bfe7d30e7f3f3d3d []byte
	__obf_83403a880bc1a9b9, __obf_d9b19f08bff09faf = base64.StdEncoding.DecodeString(__obf_2e8adb82acaa0ab7)
	if __obf_d9b19f08bff09faf != nil {
		return
	}
	__obf_eb194fea2e30bbb4, __obf_d9b19f08bff09faf = __obf_d9b3d1b65722fad4(__obf_819b883bbd4b7b1a)
	if __obf_d9b19f08bff09faf != nil {
		panic(__obf_d9b19f08bff09faf)
	}
	__obf_1a8b597c2a6b451c, __obf_0bae1112d834471f, __obf_bfe7d30e7f3f3d3d, __obf_d9b19f08bff09faf = AESDecryptMsg(__obf_83403a880bc1a9b9, __obf_eb194fea2e30bbb4)
	if __obf_d9b19f08bff09faf != nil {
		__obf_d9b19f08bff09faf = fmt.Errorf("消息解密失败,%v", __obf_d9b19f08bff09faf)
		return
	}
	if __obf_130053b95e1ae28e != string(__obf_bfe7d30e7f3f3d3d) {
		__obf_d9b19f08bff09faf = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_d9b3d1b65722fad4(__obf_05786ed4affd7637 string) (__obf_eb194fea2e30bbb4 []byte, __obf_d9b19f08bff09faf error) {
	if len(__obf_05786ed4affd7637) != 43 {
		__obf_d9b19f08bff09faf = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_eb194fea2e30bbb4, __obf_d9b19f08bff09faf = base64.StdEncoding.DecodeString(__obf_05786ed4affd7637 + "=")
	if __obf_d9b19f08bff09faf != nil {
		return
	}
	if len(__obf_eb194fea2e30bbb4) != 32 {
		__obf_d9b19f08bff09faf = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_2e353653f52b4b5b []byte, __obf_819b883bbd4b7b1a []byte) (__obf_1a8b597c2a6b451c, __obf_e8aac65c3b092c48, __obf_130053b95e1ae28e []byte, __obf_d9b19f08bff09faf error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_2e353653f52b4b5b) < BlockSize {
		__obf_d9b19f08bff09faf = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_2e353653f52b4b5b))
		return
	}
	if len(__obf_2e353653f52b4b5b)&BlockMask != 0 {
		__obf_d9b19f08bff09faf = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_2e353653f52b4b5b))
		return
	}
	__obf_9988ac51e603ec89 := make([]byte, len(__obf_2e353653f52b4b5b))
	__obf_850a472cd2c60962, // len(plaintext) >= BLOCK_SIZE
		__obf_d9b19f08bff09faf := // 解密
		aes.NewCipher(__obf_819b883bbd4b7b1a)
	if __obf_d9b19f08bff09faf != nil {
		panic(__obf_d9b19f08bff09faf)
	}
	__obf_79c31f2cb1529b76 := cipher.NewCBCDecrypter(__obf_850a472cd2c60962, __obf_819b883bbd4b7b1a[:16])
	__obf_79c31f2cb1529b76.
		CryptBlocks(__obf_9988ac51e603ec89, __obf_2e353653f52b4b5b)
	__obf_eb5031eef93a603d := // PKCS#7 去除补位
		int(__obf_9988ac51e603ec89[len(__obf_9988ac51e603ec89)-1])
	if __obf_eb5031eef93a603d < 1 || __obf_eb5031eef93a603d > BlockSize {
		__obf_d9b19f08bff09faf = fmt.Errorf("the amount to pad is incorrect: %d", __obf_eb5031eef93a603d)
		return
	}
	__obf_9988ac51e603ec89 = __obf_9988ac51e603ec89[:len(__obf_9988ac51e603ec89)-__obf_eb5031eef93a603d]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_9988ac51e603ec89) <= 20 {
		__obf_d9b19f08bff09faf = fmt.Errorf("plaintext too short, the length is %d", len(__obf_9988ac51e603ec89))
		return
	}
	__obf_fe5a93af9f4a711f := int(__obf_0d040528695b1021(__obf_9988ac51e603ec89[16:20]))
	if __obf_fe5a93af9f4a711f < 0 {
		__obf_d9b19f08bff09faf = fmt.Errorf("incorrect msg length: %d", __obf_fe5a93af9f4a711f)
		return
	}
	__obf_40ef524e717b314b := 20 + __obf_fe5a93af9f4a711f
	if len(__obf_9988ac51e603ec89) <= __obf_40ef524e717b314b {
		__obf_d9b19f08bff09faf = fmt.Errorf("msg length too large: %d", __obf_fe5a93af9f4a711f)
		return
	}
	__obf_1a8b597c2a6b451c = __obf_9988ac51e603ec89[:16:20]
	__obf_e8aac65c3b092c48 = __obf_9988ac51e603ec89[20:__obf_40ef524e717b314b:__obf_40ef524e717b314b]
	__obf_130053b95e1ae28e = __obf_9988ac51e603ec89[__obf_40ef524e717b314b:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_55604a055b096813(__obf_665ebfa5834c1051 []byte, __obf_f9df8e793bc650fc uint32) {
	__obf_665ebfa5834c1051[0] = byte(__obf_f9df8e793bc650fc >> 24)
	__obf_665ebfa5834c1051[1] = byte(__obf_f9df8e793bc650fc >> 16)
	__obf_665ebfa5834c1051[2] = byte(__obf_f9df8e793bc650fc >> 8)
	__obf_665ebfa5834c1051[3] = byte(__obf_f9df8e793bc650fc)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_0d040528695b1021(__obf_665ebfa5834c1051 []byte) (__obf_f9df8e793bc650fc uint32) {
	return uint32(__obf_665ebfa5834c1051[0])<<24 |
		uint32(__obf_665ebfa5834c1051[1])<<16 |
		uint32(__obf_665ebfa5834c1051[2])<<8 |
		uint32(__obf_665ebfa5834c1051[3])
}
