package __obf_79882becdca92e5e

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
	__obf_bb293d070532c35d := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_205c9a2417fbe739 := 32
	__obf_820f7f3491b70bcf := make([]byte, __obf_205c9a2417fbe739)
	for __obf_8cb08c0673c0264d := range __obf_820f7f3491b70bcf {
		__obf_820f7f3491b70bcf[__obf_8cb08c0673c0264d] = __obf_bb293d070532c35d[rand.IntN(len(__obf_bb293d070532c35d))]
	}
	return string(__obf_820f7f3491b70bcf)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_4ef45297d00e696a string, __obf_2a7f137f923aa828 map[string]string) string {
	__obf_d772a2056aeb1321 := make([]string, 0, len(__obf_2a7f137f923aa828))
	for __obf_58a71da0eb799615 := range __obf_2a7f137f923aa828 {
		if __obf_58a71da0eb799615 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_d772a2056aeb1321 = append(__obf_d772a2056aeb1321, __obf_58a71da0eb799615)
	}
	sort.Strings(__obf_d772a2056aeb1321) // 字典序排序

	var __obf_de0d08f0cf01055e bytes.Buffer
	for _, __obf_58a71da0eb799615 := range __obf_d772a2056aeb1321 {
		if __obf_2a7f137f923aa828[__obf_58a71da0eb799615] != "" { // 空值不参与签名
			__obf_de0d08f0cf01055e.WriteString(__obf_58a71da0eb799615)
			__obf_de0d08f0cf01055e.WriteString("=")
			__obf_de0d08f0cf01055e.WriteString(__obf_2a7f137f923aa828[__obf_58a71da0eb799615])
			__obf_de0d08f0cf01055e.WriteString("&")
		}
	}
	__obf_de0d08f0cf01055e.WriteString("key=")
	__obf_de0d08f0cf01055e.WriteString(__obf_4ef45297d00e696a)

	var __obf_e12193d6f00e1029 hash.Hash
	if __obf_2a7f137f923aa828["sign_type"] == SIGN_HMACSHA256 {
		__obf_e12193d6f00e1029 = hmac.New(sha256.New, []byte(__obf_4ef45297d00e696a))
	} else {
		__obf_e12193d6f00e1029 = md5.New()
	}
	__obf_e12193d6f00e1029.Write(__obf_de0d08f0cf01055e.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_e12193d6f00e1029.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_4ef45297d00e696a string, __obf_2a7f137f923aa828 map[string]string) bool {
	__obf_fb18d6fa4494ae64, __obf_76fee97201752e9f := __obf_2a7f137f923aa828["sign"]
	if !__obf_76fee97201752e9f {
		return false // 没有sign字段
	}
	delete(__obf_2a7f137f923aa828, "sign") // 签名时需要移除sign字段

	__obf_75e512cfa4e9028d := CreateSign(__obf_4ef45297d00e696a, __obf_2a7f137f923aa828)

	return __obf_75e512cfa4e9028d == __obf_fb18d6fa4494ae64
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_d292a8768a7a03ae map[string]string) []byte {
	var __obf_de0d08f0cf01055e bytes.Buffer
	__obf_de0d08f0cf01055e.WriteString("<xml>")
	for __obf_58a71da0eb799615, __obf_0be32d5ff1d7c891 := range __obf_d292a8768a7a03ae {
		__obf_de0d08f0cf01055e.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_58a71da0eb799615, __obf_0be32d5ff1d7c891, __obf_58a71da0eb799615))
	}
	__obf_de0d08f0cf01055e.WriteString("</xml>")
	return __obf_de0d08f0cf01055e.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_f3984caf6f36c65e []byte) (map[string]string, error) {
	__obf_c96b4bb1943b999c := xml.NewDecoder(bytes.NewReader(__obf_f3984caf6f36c65e))
	__obf_d292a8768a7a03ae := make(map[string]string)
	for {
		__obf_85a01bc837a021c4, __obf_d0090ed9a7614e3f := __obf_c96b4bb1943b999c.Token()
		if __obf_d0090ed9a7614e3f != nil {
			if __obf_d0090ed9a7614e3f.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_d0090ed9a7614e3f)
		}

		switch __obf_60ce27d5d37dc5e9 := __obf_85a01bc837a021c4.(type) {
		case xml.StartElement:
			if __obf_60ce27d5d37dc5e9.Name.Local != "xml" { // 跳过根元素
				var __obf_dd3eb4763018e5f4 string
				if __obf_d0090ed9a7614e3f := __obf_c96b4bb1943b999c.DecodeElement(&__obf_dd3eb4763018e5f4, &__obf_60ce27d5d37dc5e9); __obf_d0090ed9a7614e3f != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_d0090ed9a7614e3f)
				}
				__obf_d292a8768a7a03ae[__obf_60ce27d5d37dc5e9.Name.Local] = __obf_dd3eb4763018e5f4
			}
		}
	}
	return __obf_d292a8768a7a03ae, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_3a6bcbbdc5144975, __obf_c691cfba843709e6 []byte, __obf_70d861cfe16a8234, __obf_fe629001e4bc7c72 string) (__obf_82e05a7b889cb4d8 []byte, __obf_d0090ed9a7614e3f error) {
	defer func() {
		if __obf_c686e83d1e1383f0 := recover(); __obf_c686e83d1e1383f0 != nil {
			__obf_d0090ed9a7614e3f = fmt.Errorf("panic error: err=%v", __obf_c686e83d1e1383f0)
			return
		}
	}()
	var __obf_d3be629120925ff5 []byte
	__obf_d3be629120925ff5, __obf_d0090ed9a7614e3f = __obf_4d8212ae7d49ea39(__obf_fe629001e4bc7c72)
	if __obf_d0090ed9a7614e3f != nil {
		panic(__obf_d0090ed9a7614e3f)
	}
	__obf_3f5e435edf8d2673 := AESEncryptMsg(__obf_3a6bcbbdc5144975, __obf_c691cfba843709e6, __obf_70d861cfe16a8234, __obf_d3be629120925ff5)
	__obf_82e05a7b889cb4d8 = []byte(base64.StdEncoding.EncodeToString(__obf_3f5e435edf8d2673))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_3a6bcbbdc5144975, __obf_c691cfba843709e6 []byte, __obf_70d861cfe16a8234 string, __obf_fe629001e4bc7c72 []byte) (__obf_3f5e435edf8d2673 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	__obf_2b4dd10a23b78dc8 := 20 + len(__obf_c691cfba843709e6)
	__obf_ca8429092bbd5ff2 := __obf_2b4dd10a23b78dc8 + len(__obf_70d861cfe16a8234)
	__obf_1656153f9a881b3f := BlockSize - __obf_ca8429092bbd5ff2&BlockMask
	__obf_4102131a33d26954 := __obf_ca8429092bbd5ff2 + __obf_1656153f9a881b3f

	__obf_5703739e59f05831 := make([]byte, __obf_4102131a33d26954)

	// 拼接
	copy(__obf_5703739e59f05831[:16], __obf_3a6bcbbdc5144975)
	__obf_79666a1c000f3339(__obf_5703739e59f05831[16:20], uint32(len(__obf_c691cfba843709e6)))
	copy(__obf_5703739e59f05831[20:], __obf_c691cfba843709e6)
	copy(__obf_5703739e59f05831[__obf_2b4dd10a23b78dc8:], __obf_70d861cfe16a8234)

	// PKCS#7 补位
	for __obf_8cb08c0673c0264d := __obf_ca8429092bbd5ff2; __obf_8cb08c0673c0264d < __obf_4102131a33d26954; __obf_8cb08c0673c0264d++ {
		__obf_5703739e59f05831[__obf_8cb08c0673c0264d] = byte(__obf_1656153f9a881b3f)
	}

	// 加密
	__obf_8301dc39bcbe4dde, __obf_d0090ed9a7614e3f := aes.NewCipher(__obf_fe629001e4bc7c72)
	if __obf_d0090ed9a7614e3f != nil {
		panic(__obf_d0090ed9a7614e3f)
	}
	__obf_4c0f7541bb8a9f09 := cipher.NewCBCEncrypter(__obf_8301dc39bcbe4dde, __obf_fe629001e4bc7c72[:16])
	__obf_4c0f7541bb8a9f09.CryptBlocks(__obf_5703739e59f05831, __obf_5703739e59f05831)

	__obf_3f5e435edf8d2673 = __obf_5703739e59f05831
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_70d861cfe16a8234, __obf_a8fc26ca3a9e8fc2, __obf_fe629001e4bc7c72 string) (__obf_3a6bcbbdc5144975, __obf_8201875e33ff04ed []byte, __obf_d0090ed9a7614e3f error) {
	defer func() {
		if __obf_c686e83d1e1383f0 := recover(); __obf_c686e83d1e1383f0 != nil {
			__obf_d0090ed9a7614e3f = fmt.Errorf("panic error: err=%v", __obf_c686e83d1e1383f0)
			return
		}
	}()
	var __obf_e416ff9c8b6699be, __obf_d3be629120925ff5, __obf_f1ae4d5cc0e10229 []byte
	__obf_e416ff9c8b6699be, __obf_d0090ed9a7614e3f = base64.StdEncoding.DecodeString(__obf_a8fc26ca3a9e8fc2)
	if __obf_d0090ed9a7614e3f != nil {
		return
	}
	__obf_d3be629120925ff5, __obf_d0090ed9a7614e3f = __obf_4d8212ae7d49ea39(__obf_fe629001e4bc7c72)
	if __obf_d0090ed9a7614e3f != nil {
		panic(__obf_d0090ed9a7614e3f)
	}
	__obf_3a6bcbbdc5144975, __obf_8201875e33ff04ed, __obf_f1ae4d5cc0e10229, __obf_d0090ed9a7614e3f = AESDecryptMsg(__obf_e416ff9c8b6699be, __obf_d3be629120925ff5)
	if __obf_d0090ed9a7614e3f != nil {
		__obf_d0090ed9a7614e3f = fmt.Errorf("消息解密失败,%v", __obf_d0090ed9a7614e3f)
		return
	}
	if __obf_70d861cfe16a8234 != string(__obf_f1ae4d5cc0e10229) {
		__obf_d0090ed9a7614e3f = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_4d8212ae7d49ea39(__obf_efa2e7d737932209 string) (__obf_d3be629120925ff5 []byte, __obf_d0090ed9a7614e3f error) {
	if len(__obf_efa2e7d737932209) != 43 {
		__obf_d0090ed9a7614e3f = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_d3be629120925ff5, __obf_d0090ed9a7614e3f = base64.StdEncoding.DecodeString(__obf_efa2e7d737932209 + "=")
	if __obf_d0090ed9a7614e3f != nil {
		return
	}
	if len(__obf_d3be629120925ff5) != 32 {
		__obf_d0090ed9a7614e3f = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_3f5e435edf8d2673 []byte, __obf_fe629001e4bc7c72 []byte) (__obf_3a6bcbbdc5144975, __obf_c691cfba843709e6, __obf_70d861cfe16a8234 []byte, __obf_d0090ed9a7614e3f error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_3f5e435edf8d2673) < BlockSize {
		__obf_d0090ed9a7614e3f = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_3f5e435edf8d2673))
		return
	}
	if len(__obf_3f5e435edf8d2673)&BlockMask != 0 {
		__obf_d0090ed9a7614e3f = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_3f5e435edf8d2673))
		return
	}

	__obf_5703739e59f05831 := make([]byte, len(__obf_3f5e435edf8d2673)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	__obf_8301dc39bcbe4dde, __obf_d0090ed9a7614e3f := aes.NewCipher(__obf_fe629001e4bc7c72)
	if __obf_d0090ed9a7614e3f != nil {
		panic(__obf_d0090ed9a7614e3f)
	}
	__obf_4c0f7541bb8a9f09 := cipher.NewCBCDecrypter(__obf_8301dc39bcbe4dde, __obf_fe629001e4bc7c72[:16])
	__obf_4c0f7541bb8a9f09.CryptBlocks(__obf_5703739e59f05831, __obf_3f5e435edf8d2673)

	// PKCS#7 去除补位
	__obf_1656153f9a881b3f := int(__obf_5703739e59f05831[len(__obf_5703739e59f05831)-1])
	if __obf_1656153f9a881b3f < 1 || __obf_1656153f9a881b3f > BlockSize {
		__obf_d0090ed9a7614e3f = fmt.Errorf("the amount to pad is incorrect: %d", __obf_1656153f9a881b3f)
		return
	}
	__obf_5703739e59f05831 = __obf_5703739e59f05831[:len(__obf_5703739e59f05831)-__obf_1656153f9a881b3f]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_5703739e59f05831) <= 20 {
		__obf_d0090ed9a7614e3f = fmt.Errorf("plaintext too short, the length is %d", len(__obf_5703739e59f05831))
		return
	}
	__obf_2129488d9a433118 := int(__obf_f4265f7fc5b647b8(__obf_5703739e59f05831[16:20]))
	if __obf_2129488d9a433118 < 0 {
		__obf_d0090ed9a7614e3f = fmt.Errorf("incorrect msg length: %d", __obf_2129488d9a433118)
		return
	}
	__obf_2b4dd10a23b78dc8 := 20 + __obf_2129488d9a433118
	if len(__obf_5703739e59f05831) <= __obf_2b4dd10a23b78dc8 {
		__obf_d0090ed9a7614e3f = fmt.Errorf("msg length too large: %d", __obf_2129488d9a433118)
		return
	}

	__obf_3a6bcbbdc5144975 = __obf_5703739e59f05831[:16:20]
	__obf_c691cfba843709e6 = __obf_5703739e59f05831[20:__obf_2b4dd10a23b78dc8:__obf_2b4dd10a23b78dc8]
	__obf_70d861cfe16a8234 = __obf_5703739e59f05831[__obf_2b4dd10a23b78dc8:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_79666a1c000f3339(__obf_11f8c0d7be11a877 []byte, __obf_e6cde2a7f7f13646 uint32) {
	__obf_11f8c0d7be11a877[0] = byte(__obf_e6cde2a7f7f13646 >> 24)
	__obf_11f8c0d7be11a877[1] = byte(__obf_e6cde2a7f7f13646 >> 16)
	__obf_11f8c0d7be11a877[2] = byte(__obf_e6cde2a7f7f13646 >> 8)
	__obf_11f8c0d7be11a877[3] = byte(__obf_e6cde2a7f7f13646)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_f4265f7fc5b647b8(__obf_11f8c0d7be11a877 []byte) (__obf_e6cde2a7f7f13646 uint32) {
	return uint32(__obf_11f8c0d7be11a877[0])<<24 |
		uint32(__obf_11f8c0d7be11a877[1])<<16 |
		uint32(__obf_11f8c0d7be11a877[2])<<8 |
		uint32(__obf_11f8c0d7be11a877[3])
}
