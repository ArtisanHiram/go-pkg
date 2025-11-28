package __obf_6a74725dc3694218

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
	__obf_ad18ae480e4998ee := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_da394bd28ba30c31 := 32
	__obf_711d145e026f1af0 := make([]byte, __obf_da394bd28ba30c31)
	for __obf_0732f02da321886d := range __obf_711d145e026f1af0 {
		__obf_711d145e026f1af0[__obf_0732f02da321886d] = __obf_ad18ae480e4998ee[rand.IntN(len(__obf_ad18ae480e4998ee))]
	}
	return string(__obf_711d145e026f1af0)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_8e9228c0af1b6128 string, __obf_d53e5b3eaa475bad map[string]string) string {
	__obf_c48507a5d3e43e34 := make([]string, 0, len(__obf_d53e5b3eaa475bad))
	for __obf_5014165c4e4961f3 := range __obf_d53e5b3eaa475bad {
		if __obf_5014165c4e4961f3 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_c48507a5d3e43e34 = append(__obf_c48507a5d3e43e34, __obf_5014165c4e4961f3)
	}
	sort.Strings(__obf_c48507a5d3e43e34) // 字典序排序

	var __obf_c7b9b95f94914450 bytes.Buffer
	for _, __obf_5014165c4e4961f3 := range __obf_c48507a5d3e43e34 {
		if __obf_d53e5b3eaa475bad[__obf_5014165c4e4961f3] != "" { // 空值不参与签名
			__obf_c7b9b95f94914450.WriteString(__obf_5014165c4e4961f3)
			__obf_c7b9b95f94914450.WriteString("=")
			__obf_c7b9b95f94914450.WriteString(__obf_d53e5b3eaa475bad[__obf_5014165c4e4961f3])
			__obf_c7b9b95f94914450.WriteString("&")
		}
	}
	__obf_c7b9b95f94914450.WriteString("key=")
	__obf_c7b9b95f94914450.WriteString(__obf_8e9228c0af1b6128)

	var __obf_136c1d0bbcf8b363 hash.Hash
	if __obf_d53e5b3eaa475bad["sign_type"] == SIGN_HMACSHA256 {
		__obf_136c1d0bbcf8b363 = hmac.New(sha256.New, []byte(__obf_8e9228c0af1b6128))
	} else {
		__obf_136c1d0bbcf8b363 = md5.New()
	}
	__obf_136c1d0bbcf8b363.Write(__obf_c7b9b95f94914450.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_136c1d0bbcf8b363.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_8e9228c0af1b6128 string, __obf_d53e5b3eaa475bad map[string]string) bool {
	__obf_3c23c7538a357bf1, __obf_7c533048f9093f95 := __obf_d53e5b3eaa475bad["sign"]
	if !__obf_7c533048f9093f95 {
		return false // 没有sign字段
	}
	delete(__obf_d53e5b3eaa475bad, "sign") // 签名时需要移除sign字段

	__obf_90c21b9b3b7c7bab := CreateSign(__obf_8e9228c0af1b6128, __obf_d53e5b3eaa475bad)

	return __obf_90c21b9b3b7c7bab == __obf_3c23c7538a357bf1
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_343c2aee43bfce9f map[string]string) []byte {
	var __obf_c7b9b95f94914450 bytes.Buffer
	__obf_c7b9b95f94914450.WriteString("<xml>")
	for __obf_5014165c4e4961f3, __obf_2f5f1559bbfb1282 := range __obf_343c2aee43bfce9f {
		__obf_c7b9b95f94914450.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_5014165c4e4961f3, __obf_2f5f1559bbfb1282, __obf_5014165c4e4961f3))
	}
	__obf_c7b9b95f94914450.WriteString("</xml>")
	return __obf_c7b9b95f94914450.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_ad9d5e9c78187397 []byte) (map[string]string, error) {
	__obf_a98a65a15a683d11 := xml.NewDecoder(bytes.NewReader(__obf_ad9d5e9c78187397))
	__obf_343c2aee43bfce9f := make(map[string]string)
	for {
		__obf_cb09769d41a7767b, __obf_b5433c9ceabad42d := __obf_a98a65a15a683d11.Token()
		if __obf_b5433c9ceabad42d != nil {
			if __obf_b5433c9ceabad42d.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_b5433c9ceabad42d)
		}

		switch __obf_48e372102a8e1774 := __obf_cb09769d41a7767b.(type) {
		case xml.StartElement:
			if __obf_48e372102a8e1774.Name.Local != "xml" { // 跳过根元素
				var __obf_57f55a999a9eb547 string
				if __obf_b5433c9ceabad42d := __obf_a98a65a15a683d11.DecodeElement(&__obf_57f55a999a9eb547, &__obf_48e372102a8e1774); __obf_b5433c9ceabad42d != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_b5433c9ceabad42d)
				}
				__obf_343c2aee43bfce9f[__obf_48e372102a8e1774.Name.Local] = __obf_57f55a999a9eb547
			}
		}
	}
	return __obf_343c2aee43bfce9f, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_aa08cd69da157758, __obf_202e1585b3593a95 []byte, __obf_d14316c0e8f1161d, __obf_8ac6328020d68b33 string) (__obf_34c0ba8ba9042ff4 []byte, __obf_b5433c9ceabad42d error) {
	defer func() {
		if __obf_fb95f2217278210a := recover(); __obf_fb95f2217278210a != nil {
			__obf_b5433c9ceabad42d = fmt.Errorf("panic error: err=%v", __obf_fb95f2217278210a)
			return
		}
	}()
	var __obf_a5717b59e2d4a940 []byte
	__obf_a5717b59e2d4a940, __obf_b5433c9ceabad42d = __obf_6bfcaeb1d890c3b1(__obf_8ac6328020d68b33)
	if __obf_b5433c9ceabad42d != nil {
		panic(__obf_b5433c9ceabad42d)
	}
	__obf_310f083e674cdc61 := AESEncryptMsg(__obf_aa08cd69da157758, __obf_202e1585b3593a95, __obf_d14316c0e8f1161d, __obf_a5717b59e2d4a940)
	__obf_34c0ba8ba9042ff4 = []byte(base64.StdEncoding.EncodeToString(__obf_310f083e674cdc61))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_aa08cd69da157758, __obf_202e1585b3593a95 []byte, __obf_d14316c0e8f1161d string, __obf_8ac6328020d68b33 []byte) (__obf_310f083e674cdc61 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	__obf_e364731281229a8a := 20 + len(__obf_202e1585b3593a95)
	__obf_9aea1d7b75b9583b := __obf_e364731281229a8a + len(__obf_d14316c0e8f1161d)
	__obf_e045f4a83d3a644a := BlockSize - __obf_9aea1d7b75b9583b&BlockMask
	__obf_4294cd0b8bcfaee7 := __obf_9aea1d7b75b9583b + __obf_e045f4a83d3a644a

	__obf_2a90715d3bcd9a4a := make([]byte, __obf_4294cd0b8bcfaee7)

	// 拼接
	copy(__obf_2a90715d3bcd9a4a[:16], __obf_aa08cd69da157758)
	__obf_7dc25118c399c2f9(__obf_2a90715d3bcd9a4a[16:20], uint32(len(__obf_202e1585b3593a95)))
	copy(__obf_2a90715d3bcd9a4a[20:], __obf_202e1585b3593a95)
	copy(__obf_2a90715d3bcd9a4a[__obf_e364731281229a8a:], __obf_d14316c0e8f1161d)

	// PKCS#7 补位
	for __obf_0732f02da321886d := __obf_9aea1d7b75b9583b; __obf_0732f02da321886d < __obf_4294cd0b8bcfaee7; __obf_0732f02da321886d++ {
		__obf_2a90715d3bcd9a4a[__obf_0732f02da321886d] = byte(__obf_e045f4a83d3a644a)
	}

	// 加密
	__obf_b22ea7160f9e6175, __obf_b5433c9ceabad42d := aes.NewCipher(__obf_8ac6328020d68b33)
	if __obf_b5433c9ceabad42d != nil {
		panic(__obf_b5433c9ceabad42d)
	}
	__obf_289b78f341e6cfa3 := cipher.NewCBCEncrypter(__obf_b22ea7160f9e6175, __obf_8ac6328020d68b33[:16])
	__obf_289b78f341e6cfa3.CryptBlocks(__obf_2a90715d3bcd9a4a, __obf_2a90715d3bcd9a4a)

	__obf_310f083e674cdc61 = __obf_2a90715d3bcd9a4a
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_d14316c0e8f1161d, __obf_9b4ea2400c47ccd6, __obf_8ac6328020d68b33 string) (__obf_aa08cd69da157758, __obf_e4d3162e3312dfd9 []byte, __obf_b5433c9ceabad42d error) {
	defer func() {
		if __obf_fb95f2217278210a := recover(); __obf_fb95f2217278210a != nil {
			__obf_b5433c9ceabad42d = fmt.Errorf("panic error: err=%v", __obf_fb95f2217278210a)
			return
		}
	}()
	var __obf_2029f67ee8e997fc, __obf_a5717b59e2d4a940, __obf_84c90adac5fa494d []byte
	__obf_2029f67ee8e997fc, __obf_b5433c9ceabad42d = base64.StdEncoding.DecodeString(__obf_9b4ea2400c47ccd6)
	if __obf_b5433c9ceabad42d != nil {
		return
	}
	__obf_a5717b59e2d4a940, __obf_b5433c9ceabad42d = __obf_6bfcaeb1d890c3b1(__obf_8ac6328020d68b33)
	if __obf_b5433c9ceabad42d != nil {
		panic(__obf_b5433c9ceabad42d)
	}
	__obf_aa08cd69da157758, __obf_e4d3162e3312dfd9, __obf_84c90adac5fa494d, __obf_b5433c9ceabad42d = AESDecryptMsg(__obf_2029f67ee8e997fc, __obf_a5717b59e2d4a940)
	if __obf_b5433c9ceabad42d != nil {
		__obf_b5433c9ceabad42d = fmt.Errorf("消息解密失败,%v", __obf_b5433c9ceabad42d)
		return
	}
	if __obf_d14316c0e8f1161d != string(__obf_84c90adac5fa494d) {
		__obf_b5433c9ceabad42d = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_6bfcaeb1d890c3b1(__obf_aaf77dc3753bfd92 string) (__obf_a5717b59e2d4a940 []byte, __obf_b5433c9ceabad42d error) {
	if len(__obf_aaf77dc3753bfd92) != 43 {
		__obf_b5433c9ceabad42d = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_a5717b59e2d4a940, __obf_b5433c9ceabad42d = base64.StdEncoding.DecodeString(__obf_aaf77dc3753bfd92 + "=")
	if __obf_b5433c9ceabad42d != nil {
		return
	}
	if len(__obf_a5717b59e2d4a940) != 32 {
		__obf_b5433c9ceabad42d = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_310f083e674cdc61 []byte, __obf_8ac6328020d68b33 []byte) (__obf_aa08cd69da157758, __obf_202e1585b3593a95, __obf_d14316c0e8f1161d []byte, __obf_b5433c9ceabad42d error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_310f083e674cdc61) < BlockSize {
		__obf_b5433c9ceabad42d = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_310f083e674cdc61))
		return
	}
	if len(__obf_310f083e674cdc61)&BlockMask != 0 {
		__obf_b5433c9ceabad42d = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_310f083e674cdc61))
		return
	}

	__obf_2a90715d3bcd9a4a := make([]byte, len(__obf_310f083e674cdc61)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	__obf_b22ea7160f9e6175, __obf_b5433c9ceabad42d := aes.NewCipher(__obf_8ac6328020d68b33)
	if __obf_b5433c9ceabad42d != nil {
		panic(__obf_b5433c9ceabad42d)
	}
	__obf_289b78f341e6cfa3 := cipher.NewCBCDecrypter(__obf_b22ea7160f9e6175, __obf_8ac6328020d68b33[:16])
	__obf_289b78f341e6cfa3.CryptBlocks(__obf_2a90715d3bcd9a4a, __obf_310f083e674cdc61)

	// PKCS#7 去除补位
	__obf_e045f4a83d3a644a := int(__obf_2a90715d3bcd9a4a[len(__obf_2a90715d3bcd9a4a)-1])
	if __obf_e045f4a83d3a644a < 1 || __obf_e045f4a83d3a644a > BlockSize {
		__obf_b5433c9ceabad42d = fmt.Errorf("the amount to pad is incorrect: %d", __obf_e045f4a83d3a644a)
		return
	}
	__obf_2a90715d3bcd9a4a = __obf_2a90715d3bcd9a4a[:len(__obf_2a90715d3bcd9a4a)-__obf_e045f4a83d3a644a]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_2a90715d3bcd9a4a) <= 20 {
		__obf_b5433c9ceabad42d = fmt.Errorf("plaintext too short, the length is %d", len(__obf_2a90715d3bcd9a4a))
		return
	}
	__obf_eb8962ad70986290 := int(__obf_fc595525f44d921e(__obf_2a90715d3bcd9a4a[16:20]))
	if __obf_eb8962ad70986290 < 0 {
		__obf_b5433c9ceabad42d = fmt.Errorf("incorrect msg length: %d", __obf_eb8962ad70986290)
		return
	}
	__obf_e364731281229a8a := 20 + __obf_eb8962ad70986290
	if len(__obf_2a90715d3bcd9a4a) <= __obf_e364731281229a8a {
		__obf_b5433c9ceabad42d = fmt.Errorf("msg length too large: %d", __obf_eb8962ad70986290)
		return
	}

	__obf_aa08cd69da157758 = __obf_2a90715d3bcd9a4a[:16:20]
	__obf_202e1585b3593a95 = __obf_2a90715d3bcd9a4a[20:__obf_e364731281229a8a:__obf_e364731281229a8a]
	__obf_d14316c0e8f1161d = __obf_2a90715d3bcd9a4a[__obf_e364731281229a8a:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_7dc25118c399c2f9(__obf_0b93c531565bd567 []byte, __obf_3b6d1cd0ef8bf591 uint32) {
	__obf_0b93c531565bd567[0] = byte(__obf_3b6d1cd0ef8bf591 >> 24)
	__obf_0b93c531565bd567[1] = byte(__obf_3b6d1cd0ef8bf591 >> 16)
	__obf_0b93c531565bd567[2] = byte(__obf_3b6d1cd0ef8bf591 >> 8)
	__obf_0b93c531565bd567[3] = byte(__obf_3b6d1cd0ef8bf591)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_fc595525f44d921e(__obf_0b93c531565bd567 []byte) (__obf_3b6d1cd0ef8bf591 uint32) {
	return uint32(__obf_0b93c531565bd567[0])<<24 |
		uint32(__obf_0b93c531565bd567[1])<<16 |
		uint32(__obf_0b93c531565bd567[2])<<8 |
		uint32(__obf_0b93c531565bd567[3])
}
