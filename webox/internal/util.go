package __obf_ca26b678cc0b7836

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
	__obf_d4ee26a1ea1fe7f7 := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_fa90f02e738a820d := 32
	__obf_d054c51e72a98a15 := make([]byte, __obf_fa90f02e738a820d)
	for __obf_fd44f72ebc028f54 := range __obf_d054c51e72a98a15 {
		__obf_d054c51e72a98a15[__obf_fd44f72ebc028f54] = __obf_d4ee26a1ea1fe7f7[rand.IntN(len(__obf_d4ee26a1ea1fe7f7))]
	}
	return string(__obf_d054c51e72a98a15)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_5f5b087d39fd7d5c string, __obf_8b50f88d82c9f24a map[string]string) string {
	__obf_7837f90c830aa476 := make([]string, 0, len(__obf_8b50f88d82c9f24a))
	for __obf_9a8f89bd5b4d1a17 := range __obf_8b50f88d82c9f24a {
		if __obf_9a8f89bd5b4d1a17 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_7837f90c830aa476 = append(__obf_7837f90c830aa476, __obf_9a8f89bd5b4d1a17)
	}
	sort.Strings(__obf_7837f90c830aa476) // 字典序排序

	var __obf_3753748412bf6190 bytes.Buffer
	for _, __obf_9a8f89bd5b4d1a17 := range __obf_7837f90c830aa476 {
		if __obf_8b50f88d82c9f24a[__obf_9a8f89bd5b4d1a17] != "" {
			__obf_3753748412bf6190. // 空值不参与签名
						WriteString(__obf_9a8f89bd5b4d1a17)
			__obf_3753748412bf6190.
				WriteString("=")
			__obf_3753748412bf6190.
				WriteString(__obf_8b50f88d82c9f24a[__obf_9a8f89bd5b4d1a17])
			__obf_3753748412bf6190.
				WriteString("&")
		}
	}
	__obf_3753748412bf6190.
		WriteString("key=")
	__obf_3753748412bf6190.
		WriteString(__obf_5f5b087d39fd7d5c)

	var __obf_9c851bf9be6d4fc5 hash.Hash
	if __obf_8b50f88d82c9f24a["sign_type"] == SIGN_HMACSHA256 {
		__obf_9c851bf9be6d4fc5 = hmac.New(sha256.New, []byte(__obf_5f5b087d39fd7d5c))
	} else {
		__obf_9c851bf9be6d4fc5 = md5.New()
	}
	__obf_9c851bf9be6d4fc5.
		Write(__obf_3753748412bf6190.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_9c851bf9be6d4fc5.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_5f5b087d39fd7d5c string, __obf_8b50f88d82c9f24a map[string]string) bool {
	__obf_31dbedc4d1f7cb42, __obf_7a842e41650679b4 := __obf_8b50f88d82c9f24a["sign"]
	if !__obf_7a842e41650679b4 {
		return false // 没有sign字段
	}
	delete(__obf_8b50f88d82c9f24a, "sign")
	__obf_ac901a84d81557fa := // 签名时需要移除sign字段

		CreateSign(__obf_5f5b087d39fd7d5c, __obf_8b50f88d82c9f24a)

	return __obf_ac901a84d81557fa == __obf_31dbedc4d1f7cb42
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_3548c160ae91e2ec map[string]string) []byte {
	var __obf_3753748412bf6190 bytes.Buffer
	__obf_3753748412bf6190.
		WriteString("<xml>")
	for __obf_9a8f89bd5b4d1a17, __obf_52524f49eda93617 := range __obf_3548c160ae91e2ec {
		__obf_3753748412bf6190.
			WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_9a8f89bd5b4d1a17, __obf_52524f49eda93617, __obf_9a8f89bd5b4d1a17))
	}
	__obf_3753748412bf6190.
		WriteString("</xml>")
	return __obf_3753748412bf6190.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_3c5c3f878ca019ea []byte) (map[string]string, error) {
	__obf_d334916abd9d21b4 := xml.NewDecoder(bytes.NewReader(__obf_3c5c3f878ca019ea))
	__obf_3548c160ae91e2ec := make(map[string]string)
	for {
		__obf_3b97dd0887889eb6, __obf_c3e83fee90661fe5 := __obf_d334916abd9d21b4.Token()
		if __obf_c3e83fee90661fe5 != nil {
			if __obf_c3e83fee90661fe5.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_c3e83fee90661fe5)
		}

		switch __obf_73c079348059d6e9 := __obf_3b97dd0887889eb6.(type) {
		case xml.StartElement:
			if __obf_73c079348059d6e9.Name.Local != "xml" { // 跳过根元素
				var __obf_24d7228587ad7ce1 string
				if __obf_c3e83fee90661fe5 := __obf_d334916abd9d21b4.DecodeElement(&__obf_24d7228587ad7ce1, &__obf_73c079348059d6e9); __obf_c3e83fee90661fe5 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_c3e83fee90661fe5)
				}
				__obf_3548c160ae91e2ec[__obf_73c079348059d6e9.Name.Local] = __obf_24d7228587ad7ce1
			}
		}
	}
	return __obf_3548c160ae91e2ec, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_4f164b7f750ece7d, __obf_94c1e4b107006643 []byte, __obf_7cd83025f59c00e1, __obf_52c103b74129ae20 string) (__obf_f707bcc6eb9b62a6 []byte, __obf_c3e83fee90661fe5 error) {
	defer func() {
		if __obf_3c6e2ab6a79f5047 := recover(); __obf_3c6e2ab6a79f5047 != nil {
			__obf_c3e83fee90661fe5 = fmt.Errorf("panic error: err=%v", __obf_3c6e2ab6a79f5047)
			return
		}
	}()
	var __obf_9fc2e8cc9640f833 []byte
	__obf_9fc2e8cc9640f833, __obf_c3e83fee90661fe5 = __obf_71c3bd4e71135461(__obf_52c103b74129ae20)
	if __obf_c3e83fee90661fe5 != nil {
		panic(__obf_c3e83fee90661fe5)
	}
	__obf_f4161864736f8982 := AESEncryptMsg(__obf_4f164b7f750ece7d, __obf_94c1e4b107006643, __obf_7cd83025f59c00e1, __obf_9fc2e8cc9640f833)
	__obf_f707bcc6eb9b62a6 = []byte(base64.StdEncoding.EncodeToString(__obf_f4161864736f8982))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_4f164b7f750ece7d, __obf_94c1e4b107006643 []byte, __obf_7cd83025f59c00e1 string, __obf_52c103b74129ae20 []byte) (__obf_f4161864736f8982 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)
	__obf_68b0b0b28ffceb9b := 20 + len(__obf_94c1e4b107006643)
	__obf_caa3b425d918a10b := __obf_68b0b0b28ffceb9b + len(__obf_7cd83025f59c00e1)
	__obf_78e7f5dda3f2b5e4 := BlockSize - __obf_caa3b425d918a10b&BlockMask
	__obf_eb810e7afe3e4acb := __obf_caa3b425d918a10b + __obf_78e7f5dda3f2b5e4
	__obf_51f361b696b212aa := make([]byte, __obf_eb810e7afe3e4acb)

	// 拼接
	copy(__obf_51f361b696b212aa[:16], __obf_4f164b7f750ece7d)
	__obf_635e7e91f0a945f9(__obf_51f361b696b212aa[16:20], uint32(len(__obf_94c1e4b107006643)))
	copy(__obf_51f361b696b212aa[20:], __obf_94c1e4b107006643)
	copy(__obf_51f361b696b212aa[__obf_68b0b0b28ffceb9b:], __obf_7cd83025f59c00e1)

	// PKCS#7 补位
	for __obf_fd44f72ebc028f54 := __obf_caa3b425d918a10b; __obf_fd44f72ebc028f54 < __obf_eb810e7afe3e4acb; __obf_fd44f72ebc028f54++ {
		__obf_51f361b696b212aa[__obf_fd44f72ebc028f54] = byte(__obf_78e7f5dda3f2b5e4)
	}
	__obf_ff17270f2d6d21bd,

		// 加密
		__obf_c3e83fee90661fe5 := aes.NewCipher(__obf_52c103b74129ae20)
	if __obf_c3e83fee90661fe5 != nil {
		panic(__obf_c3e83fee90661fe5)
	}
	__obf_9b7cda664361fde9 := cipher.NewCBCEncrypter(__obf_ff17270f2d6d21bd, __obf_52c103b74129ae20[:16])
	__obf_9b7cda664361fde9.
		CryptBlocks(__obf_51f361b696b212aa, __obf_51f361b696b212aa)
	__obf_f4161864736f8982 = __obf_51f361b696b212aa
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_7cd83025f59c00e1, __obf_e6bc7cd5bde3f21e, __obf_52c103b74129ae20 string) (__obf_4f164b7f750ece7d, __obf_596cc40ebe5975a6 []byte, __obf_c3e83fee90661fe5 error) {
	defer func() {
		if __obf_3c6e2ab6a79f5047 := recover(); __obf_3c6e2ab6a79f5047 != nil {
			__obf_c3e83fee90661fe5 = fmt.Errorf("panic error: err=%v", __obf_3c6e2ab6a79f5047)
			return
		}
	}()
	var __obf_a765272084bc60e8, __obf_9fc2e8cc9640f833, __obf_da087699eb96f1e8 []byte
	__obf_a765272084bc60e8, __obf_c3e83fee90661fe5 = base64.StdEncoding.DecodeString(__obf_e6bc7cd5bde3f21e)
	if __obf_c3e83fee90661fe5 != nil {
		return
	}
	__obf_9fc2e8cc9640f833, __obf_c3e83fee90661fe5 = __obf_71c3bd4e71135461(__obf_52c103b74129ae20)
	if __obf_c3e83fee90661fe5 != nil {
		panic(__obf_c3e83fee90661fe5)
	}
	__obf_4f164b7f750ece7d, __obf_596cc40ebe5975a6, __obf_da087699eb96f1e8, __obf_c3e83fee90661fe5 = AESDecryptMsg(__obf_a765272084bc60e8, __obf_9fc2e8cc9640f833)
	if __obf_c3e83fee90661fe5 != nil {
		__obf_c3e83fee90661fe5 = fmt.Errorf("消息解密失败,%v", __obf_c3e83fee90661fe5)
		return
	}
	if __obf_7cd83025f59c00e1 != string(__obf_da087699eb96f1e8) {
		__obf_c3e83fee90661fe5 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_71c3bd4e71135461(__obf_ea9bd62925c055a2 string) (__obf_9fc2e8cc9640f833 []byte, __obf_c3e83fee90661fe5 error) {
	if len(__obf_ea9bd62925c055a2) != 43 {
		__obf_c3e83fee90661fe5 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_9fc2e8cc9640f833, __obf_c3e83fee90661fe5 = base64.StdEncoding.DecodeString(__obf_ea9bd62925c055a2 + "=")
	if __obf_c3e83fee90661fe5 != nil {
		return
	}
	if len(__obf_9fc2e8cc9640f833) != 32 {
		__obf_c3e83fee90661fe5 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_f4161864736f8982 []byte, __obf_52c103b74129ae20 []byte) (__obf_4f164b7f750ece7d, __obf_94c1e4b107006643, __obf_7cd83025f59c00e1 []byte, __obf_c3e83fee90661fe5 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_f4161864736f8982) < BlockSize {
		__obf_c3e83fee90661fe5 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_f4161864736f8982))
		return
	}
	if len(__obf_f4161864736f8982)&BlockMask != 0 {
		__obf_c3e83fee90661fe5 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_f4161864736f8982))
		return
	}
	__obf_51f361b696b212aa := make([]byte, len(__obf_f4161864736f8982))
	__obf_ff17270f2d6d21bd, // len(plaintext) >= BLOCK_SIZE
		__obf_c3e83fee90661fe5 := // 解密
		aes.NewCipher(__obf_52c103b74129ae20)
	if __obf_c3e83fee90661fe5 != nil {
		panic(__obf_c3e83fee90661fe5)
	}
	__obf_9b7cda664361fde9 := cipher.NewCBCDecrypter(__obf_ff17270f2d6d21bd, __obf_52c103b74129ae20[:16])
	__obf_9b7cda664361fde9.
		CryptBlocks(__obf_51f361b696b212aa, __obf_f4161864736f8982)
	__obf_78e7f5dda3f2b5e4 := // PKCS#7 去除补位
		int(__obf_51f361b696b212aa[len(__obf_51f361b696b212aa)-1])
	if __obf_78e7f5dda3f2b5e4 < 1 || __obf_78e7f5dda3f2b5e4 > BlockSize {
		__obf_c3e83fee90661fe5 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_78e7f5dda3f2b5e4)
		return
	}
	__obf_51f361b696b212aa = __obf_51f361b696b212aa[:len(__obf_51f361b696b212aa)-__obf_78e7f5dda3f2b5e4]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_51f361b696b212aa) <= 20 {
		__obf_c3e83fee90661fe5 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_51f361b696b212aa))
		return
	}
	__obf_01a6d4bb92f4ec56 := int(__obf_2943d1eaa1257e0b(__obf_51f361b696b212aa[16:20]))
	if __obf_01a6d4bb92f4ec56 < 0 {
		__obf_c3e83fee90661fe5 = fmt.Errorf("incorrect msg length: %d", __obf_01a6d4bb92f4ec56)
		return
	}
	__obf_68b0b0b28ffceb9b := 20 + __obf_01a6d4bb92f4ec56
	if len(__obf_51f361b696b212aa) <= __obf_68b0b0b28ffceb9b {
		__obf_c3e83fee90661fe5 = fmt.Errorf("msg length too large: %d", __obf_01a6d4bb92f4ec56)
		return
	}
	__obf_4f164b7f750ece7d = __obf_51f361b696b212aa[:16:20]
	__obf_94c1e4b107006643 = __obf_51f361b696b212aa[20:__obf_68b0b0b28ffceb9b:__obf_68b0b0b28ffceb9b]
	__obf_7cd83025f59c00e1 = __obf_51f361b696b212aa[__obf_68b0b0b28ffceb9b:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_635e7e91f0a945f9(__obf_fd4bb9e790dedbd9 []byte, __obf_b076c3f82a21350f uint32) {
	__obf_fd4bb9e790dedbd9[0] = byte(__obf_b076c3f82a21350f >> 24)
	__obf_fd4bb9e790dedbd9[1] = byte(__obf_b076c3f82a21350f >> 16)
	__obf_fd4bb9e790dedbd9[2] = byte(__obf_b076c3f82a21350f >> 8)
	__obf_fd4bb9e790dedbd9[3] = byte(__obf_b076c3f82a21350f)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_2943d1eaa1257e0b(__obf_fd4bb9e790dedbd9 []byte) (__obf_b076c3f82a21350f uint32) {
	return uint32(__obf_fd4bb9e790dedbd9[0])<<24 |
		uint32(__obf_fd4bb9e790dedbd9[1])<<16 |
		uint32(__obf_fd4bb9e790dedbd9[2])<<8 |
		uint32(__obf_fd4bb9e790dedbd9[3])
}
