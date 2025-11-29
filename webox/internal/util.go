package __obf_18f8d68b9095d0e0

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
	__obf_61c7fa232f248179 := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_7e52c1900ee9de1e := 32
	__obf_4376ccafcbcc7e02 := make([]byte, __obf_7e52c1900ee9de1e)
	for __obf_a8f62e06e2afcf88 := range __obf_4376ccafcbcc7e02 {
		__obf_4376ccafcbcc7e02[__obf_a8f62e06e2afcf88] = __obf_61c7fa232f248179[rand.IntN(len(__obf_61c7fa232f248179))]
	}
	return string(__obf_4376ccafcbcc7e02)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_2aadf62169813c75 string, __obf_661c4ec7efdcfec5 map[string]string) string {
	__obf_dacccb656e7b4b25 := make([]string, 0, len(__obf_661c4ec7efdcfec5))
	for __obf_01074596fa5e7fd7 := range __obf_661c4ec7efdcfec5 {
		if __obf_01074596fa5e7fd7 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_dacccb656e7b4b25 = append(__obf_dacccb656e7b4b25, __obf_01074596fa5e7fd7)
	}
	sort.Strings(__obf_dacccb656e7b4b25) // 字典序排序

	var __obf_2c6a21aea46644d8 bytes.Buffer
	for _, __obf_01074596fa5e7fd7 := range __obf_dacccb656e7b4b25 {
		if __obf_661c4ec7efdcfec5[__obf_01074596fa5e7fd7] != "" {
			__obf_2c6a21aea46644d8. // 空值不参与签名
						WriteString(__obf_01074596fa5e7fd7)
			__obf_2c6a21aea46644d8.
				WriteString("=")
			__obf_2c6a21aea46644d8.
				WriteString(__obf_661c4ec7efdcfec5[__obf_01074596fa5e7fd7])
			__obf_2c6a21aea46644d8.
				WriteString("&")
		}
	}
	__obf_2c6a21aea46644d8.
		WriteString("key=")
	__obf_2c6a21aea46644d8.
		WriteString(__obf_2aadf62169813c75)

	var __obf_37448d443ba87549 hash.Hash
	if __obf_661c4ec7efdcfec5["sign_type"] == SIGN_HMACSHA256 {
		__obf_37448d443ba87549 = hmac.New(sha256.New, []byte(__obf_2aadf62169813c75))
	} else {
		__obf_37448d443ba87549 = md5.New()
	}
	__obf_37448d443ba87549.
		Write(__obf_2c6a21aea46644d8.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_37448d443ba87549.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_2aadf62169813c75 string, __obf_661c4ec7efdcfec5 map[string]string) bool {
	__obf_0eb5adfd5cf9690a, __obf_a375daf68c2a3268 := __obf_661c4ec7efdcfec5["sign"]
	if !__obf_a375daf68c2a3268 {
		return false // 没有sign字段
	}
	delete(__obf_661c4ec7efdcfec5, "sign")
	__obf_c156c78e787cfb3b := // 签名时需要移除sign字段

		CreateSign(__obf_2aadf62169813c75, __obf_661c4ec7efdcfec5)

	return __obf_c156c78e787cfb3b == __obf_0eb5adfd5cf9690a
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_2a0c5f16e13ec5e3 map[string]string) []byte {
	var __obf_2c6a21aea46644d8 bytes.Buffer
	__obf_2c6a21aea46644d8.
		WriteString("<xml>")
	for __obf_01074596fa5e7fd7, __obf_4b612b03fe7ab91c := range __obf_2a0c5f16e13ec5e3 {
		__obf_2c6a21aea46644d8.
			WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_01074596fa5e7fd7, __obf_4b612b03fe7ab91c, __obf_01074596fa5e7fd7))
	}
	__obf_2c6a21aea46644d8.
		WriteString("</xml>")
	return __obf_2c6a21aea46644d8.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_005c08a3be877955 []byte) (map[string]string, error) {
	__obf_b020e77ecdf2329a := xml.NewDecoder(bytes.NewReader(__obf_005c08a3be877955))
	__obf_2a0c5f16e13ec5e3 := make(map[string]string)
	for {
		__obf_d64a828c5268f3c7, __obf_f45e33988e88e5a2 := __obf_b020e77ecdf2329a.Token()
		if __obf_f45e33988e88e5a2 != nil {
			if __obf_f45e33988e88e5a2.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_f45e33988e88e5a2)
		}

		switch __obf_02bd975eb13f45d2 := __obf_d64a828c5268f3c7.(type) {
		case xml.StartElement:
			if __obf_02bd975eb13f45d2.Name.Local != "xml" { // 跳过根元素
				var __obf_d225bd4d793c0823 string
				if __obf_f45e33988e88e5a2 := __obf_b020e77ecdf2329a.DecodeElement(&__obf_d225bd4d793c0823, &__obf_02bd975eb13f45d2); __obf_f45e33988e88e5a2 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_f45e33988e88e5a2)
				}
				__obf_2a0c5f16e13ec5e3[__obf_02bd975eb13f45d2.Name.Local] = __obf_d225bd4d793c0823
			}
		}
	}
	return __obf_2a0c5f16e13ec5e3, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_bb442d4d36203a1d, __obf_3dfab2ee76818119 []byte, __obf_d35b8915f140b44d, __obf_a51f215a93568121 string) (__obf_d4c2a698139ab514 []byte, __obf_f45e33988e88e5a2 error) {
	defer func() {
		if __obf_5a4b339451be9661 := recover(); __obf_5a4b339451be9661 != nil {
			__obf_f45e33988e88e5a2 = fmt.Errorf("panic error: err=%v", __obf_5a4b339451be9661)
			return
		}
	}()
	var __obf_74e39c87e7c65a0b []byte
	__obf_74e39c87e7c65a0b, __obf_f45e33988e88e5a2 = __obf_34eb618228f1a2fd(__obf_a51f215a93568121)
	if __obf_f45e33988e88e5a2 != nil {
		panic(__obf_f45e33988e88e5a2)
	}
	__obf_e363473771864094 := AESEncryptMsg(__obf_bb442d4d36203a1d, __obf_3dfab2ee76818119, __obf_d35b8915f140b44d, __obf_74e39c87e7c65a0b)
	__obf_d4c2a698139ab514 = []byte(base64.StdEncoding.EncodeToString(__obf_e363473771864094))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_bb442d4d36203a1d, __obf_3dfab2ee76818119 []byte, __obf_d35b8915f140b44d string, __obf_a51f215a93568121 []byte) (__obf_e363473771864094 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)
	__obf_8cc8c2b8a4e9c389 := 20 + len(__obf_3dfab2ee76818119)
	__obf_dd592b3f3975123f := __obf_8cc8c2b8a4e9c389 + len(__obf_d35b8915f140b44d)
	__obf_90012175e06937d0 := BlockSize - __obf_dd592b3f3975123f&BlockMask
	__obf_0c08ec106dbf22e6 := __obf_dd592b3f3975123f + __obf_90012175e06937d0
	__obf_a8003223d6de554b := make([]byte, __obf_0c08ec106dbf22e6)

	// 拼接
	copy(__obf_a8003223d6de554b[:16], __obf_bb442d4d36203a1d)
	__obf_1de0f2788193b136(__obf_a8003223d6de554b[16:20], uint32(len(__obf_3dfab2ee76818119)))
	copy(__obf_a8003223d6de554b[20:], __obf_3dfab2ee76818119)
	copy(__obf_a8003223d6de554b[__obf_8cc8c2b8a4e9c389:], __obf_d35b8915f140b44d)

	// PKCS#7 补位
	for __obf_a8f62e06e2afcf88 := __obf_dd592b3f3975123f; __obf_a8f62e06e2afcf88 < __obf_0c08ec106dbf22e6; __obf_a8f62e06e2afcf88++ {
		__obf_a8003223d6de554b[__obf_a8f62e06e2afcf88] = byte(__obf_90012175e06937d0)
	}
	__obf_6b1f50188faf4395,

		// 加密
		__obf_f45e33988e88e5a2 := aes.NewCipher(__obf_a51f215a93568121)
	if __obf_f45e33988e88e5a2 != nil {
		panic(__obf_f45e33988e88e5a2)
	}
	__obf_6c2e4abaf2a746d5 := cipher.NewCBCEncrypter(__obf_6b1f50188faf4395, __obf_a51f215a93568121[:16])
	__obf_6c2e4abaf2a746d5.
		CryptBlocks(__obf_a8003223d6de554b, __obf_a8003223d6de554b)
	__obf_e363473771864094 = __obf_a8003223d6de554b
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_d35b8915f140b44d, __obf_33aa3e8c0d953a73, __obf_a51f215a93568121 string) (__obf_bb442d4d36203a1d, __obf_8b2a44d2698d9d7a []byte, __obf_f45e33988e88e5a2 error) {
	defer func() {
		if __obf_5a4b339451be9661 := recover(); __obf_5a4b339451be9661 != nil {
			__obf_f45e33988e88e5a2 = fmt.Errorf("panic error: err=%v", __obf_5a4b339451be9661)
			return
		}
	}()
	var __obf_4c7a241147250e26, __obf_74e39c87e7c65a0b, __obf_123ff268e1032e3d []byte
	__obf_4c7a241147250e26, __obf_f45e33988e88e5a2 = base64.StdEncoding.DecodeString(__obf_33aa3e8c0d953a73)
	if __obf_f45e33988e88e5a2 != nil {
		return
	}
	__obf_74e39c87e7c65a0b, __obf_f45e33988e88e5a2 = __obf_34eb618228f1a2fd(__obf_a51f215a93568121)
	if __obf_f45e33988e88e5a2 != nil {
		panic(__obf_f45e33988e88e5a2)
	}
	__obf_bb442d4d36203a1d, __obf_8b2a44d2698d9d7a, __obf_123ff268e1032e3d, __obf_f45e33988e88e5a2 = AESDecryptMsg(__obf_4c7a241147250e26, __obf_74e39c87e7c65a0b)
	if __obf_f45e33988e88e5a2 != nil {
		__obf_f45e33988e88e5a2 = fmt.Errorf("消息解密失败,%v", __obf_f45e33988e88e5a2)
		return
	}
	if __obf_d35b8915f140b44d != string(__obf_123ff268e1032e3d) {
		__obf_f45e33988e88e5a2 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_34eb618228f1a2fd(__obf_297a0083412110c6 string) (__obf_74e39c87e7c65a0b []byte, __obf_f45e33988e88e5a2 error) {
	if len(__obf_297a0083412110c6) != 43 {
		__obf_f45e33988e88e5a2 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_74e39c87e7c65a0b, __obf_f45e33988e88e5a2 = base64.StdEncoding.DecodeString(__obf_297a0083412110c6 + "=")
	if __obf_f45e33988e88e5a2 != nil {
		return
	}
	if len(__obf_74e39c87e7c65a0b) != 32 {
		__obf_f45e33988e88e5a2 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_e363473771864094 []byte, __obf_a51f215a93568121 []byte) (__obf_bb442d4d36203a1d, __obf_3dfab2ee76818119, __obf_d35b8915f140b44d []byte, __obf_f45e33988e88e5a2 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_e363473771864094) < BlockSize {
		__obf_f45e33988e88e5a2 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_e363473771864094))
		return
	}
	if len(__obf_e363473771864094)&BlockMask != 0 {
		__obf_f45e33988e88e5a2 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_e363473771864094))
		return
	}
	__obf_a8003223d6de554b := make([]byte, len(__obf_e363473771864094))
	__obf_6b1f50188faf4395, // len(plaintext) >= BLOCK_SIZE
		__obf_f45e33988e88e5a2 := // 解密
		aes.NewCipher(__obf_a51f215a93568121)
	if __obf_f45e33988e88e5a2 != nil {
		panic(__obf_f45e33988e88e5a2)
	}
	__obf_6c2e4abaf2a746d5 := cipher.NewCBCDecrypter(__obf_6b1f50188faf4395, __obf_a51f215a93568121[:16])
	__obf_6c2e4abaf2a746d5.
		CryptBlocks(__obf_a8003223d6de554b, __obf_e363473771864094)
	__obf_90012175e06937d0 := // PKCS#7 去除补位
		int(__obf_a8003223d6de554b[len(__obf_a8003223d6de554b)-1])
	if __obf_90012175e06937d0 < 1 || __obf_90012175e06937d0 > BlockSize {
		__obf_f45e33988e88e5a2 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_90012175e06937d0)
		return
	}
	__obf_a8003223d6de554b = __obf_a8003223d6de554b[:len(__obf_a8003223d6de554b)-__obf_90012175e06937d0]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_a8003223d6de554b) <= 20 {
		__obf_f45e33988e88e5a2 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_a8003223d6de554b))
		return
	}
	__obf_2cea4d58bfd878ec := int(__obf_023277ff901b886f(__obf_a8003223d6de554b[16:20]))
	if __obf_2cea4d58bfd878ec < 0 {
		__obf_f45e33988e88e5a2 = fmt.Errorf("incorrect msg length: %d", __obf_2cea4d58bfd878ec)
		return
	}
	__obf_8cc8c2b8a4e9c389 := 20 + __obf_2cea4d58bfd878ec
	if len(__obf_a8003223d6de554b) <= __obf_8cc8c2b8a4e9c389 {
		__obf_f45e33988e88e5a2 = fmt.Errorf("msg length too large: %d", __obf_2cea4d58bfd878ec)
		return
	}
	__obf_bb442d4d36203a1d = __obf_a8003223d6de554b[:16:20]
	__obf_3dfab2ee76818119 = __obf_a8003223d6de554b[20:__obf_8cc8c2b8a4e9c389:__obf_8cc8c2b8a4e9c389]
	__obf_d35b8915f140b44d = __obf_a8003223d6de554b[__obf_8cc8c2b8a4e9c389:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_1de0f2788193b136(__obf_78b2843ff8aa88a3 []byte, __obf_76ff9a2218821ac1 uint32) {
	__obf_78b2843ff8aa88a3[0] = byte(__obf_76ff9a2218821ac1 >> 24)
	__obf_78b2843ff8aa88a3[1] = byte(__obf_76ff9a2218821ac1 >> 16)
	__obf_78b2843ff8aa88a3[2] = byte(__obf_76ff9a2218821ac1 >> 8)
	__obf_78b2843ff8aa88a3[3] = byte(__obf_76ff9a2218821ac1)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_023277ff901b886f(__obf_78b2843ff8aa88a3 []byte) (__obf_76ff9a2218821ac1 uint32) {
	return uint32(__obf_78b2843ff8aa88a3[0])<<24 |
		uint32(__obf_78b2843ff8aa88a3[1])<<16 |
		uint32(__obf_78b2843ff8aa88a3[2])<<8 |
		uint32(__obf_78b2843ff8aa88a3[3])
}
