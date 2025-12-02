package __obf_dfacbdcb930f08e7

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
	__obf_82c62ecf497f2cb5 := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_703404ed2aca900b := 32
	__obf_0e8f3aea5ce75142 := make([]byte, __obf_703404ed2aca900b)
	for __obf_a5e6bc4d45d78213 := range __obf_0e8f3aea5ce75142 {
		__obf_0e8f3aea5ce75142[__obf_a5e6bc4d45d78213] = __obf_82c62ecf497f2cb5[rand.IntN(len(__obf_82c62ecf497f2cb5))]
	}
	return string(__obf_0e8f3aea5ce75142)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_4ff104a5f44cc078 string, __obf_9fa876ce58e46027 map[string]string) string {
	__obf_439708949bc0cf06 := make([]string, 0, len(__obf_9fa876ce58e46027))
	for __obf_6ec0c0d7916f38a1 := range __obf_9fa876ce58e46027 {
		if __obf_6ec0c0d7916f38a1 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_439708949bc0cf06 = append(__obf_439708949bc0cf06, __obf_6ec0c0d7916f38a1)
	}
	sort.Strings(__obf_439708949bc0cf06) // 字典序排序

	var __obf_ada382eb7750507c bytes.Buffer
	for _, __obf_6ec0c0d7916f38a1 := range __obf_439708949bc0cf06 {
		if __obf_9fa876ce58e46027[__obf_6ec0c0d7916f38a1] != "" {
			__obf_ada382eb7750507c. // 空值不参与签名
						WriteString(__obf_6ec0c0d7916f38a1)
			__obf_ada382eb7750507c.
				WriteString("=")
			__obf_ada382eb7750507c.
				WriteString(__obf_9fa876ce58e46027[__obf_6ec0c0d7916f38a1])
			__obf_ada382eb7750507c.
				WriteString("&")
		}
	}
	__obf_ada382eb7750507c.
		WriteString("key=")
	__obf_ada382eb7750507c.
		WriteString(__obf_4ff104a5f44cc078)

	var __obf_23f4410a795cee28 hash.Hash
	if __obf_9fa876ce58e46027["sign_type"] == SIGN_HMACSHA256 {
		__obf_23f4410a795cee28 = hmac.New(sha256.New, []byte(__obf_4ff104a5f44cc078))
	} else {
		__obf_23f4410a795cee28 = md5.New()
	}
	__obf_23f4410a795cee28.
		Write(__obf_ada382eb7750507c.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_23f4410a795cee28.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_4ff104a5f44cc078 string, __obf_9fa876ce58e46027 map[string]string) bool {
	__obf_f812ecf8ae70ee41, __obf_77989d717746a23b := __obf_9fa876ce58e46027["sign"]
	if !__obf_77989d717746a23b {
		return false // 没有sign字段
	}
	delete(__obf_9fa876ce58e46027, "sign")
	__obf_60c34bd77c5e53f3 := // 签名时需要移除sign字段

		CreateSign(__obf_4ff104a5f44cc078, __obf_9fa876ce58e46027)

	return __obf_60c34bd77c5e53f3 == __obf_f812ecf8ae70ee41
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_87f37beff62d945f map[string]string) []byte {
	var __obf_ada382eb7750507c bytes.Buffer
	__obf_ada382eb7750507c.
		WriteString("<xml>")
	for __obf_6ec0c0d7916f38a1, __obf_dc967b23178581e0 := range __obf_87f37beff62d945f {
		__obf_ada382eb7750507c.
			WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_6ec0c0d7916f38a1, __obf_dc967b23178581e0, __obf_6ec0c0d7916f38a1))
	}
	__obf_ada382eb7750507c.
		WriteString("</xml>")
	return __obf_ada382eb7750507c.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_a7b6b5f4f8962835 []byte) (map[string]string, error) {
	__obf_fa962c8ba8fdfd03 := xml.NewDecoder(bytes.NewReader(__obf_a7b6b5f4f8962835))
	__obf_87f37beff62d945f := make(map[string]string)
	for {
		__obf_24cc1bd1cbebd895, __obf_f304d4a353574273 := __obf_fa962c8ba8fdfd03.Token()
		if __obf_f304d4a353574273 != nil {
			if __obf_f304d4a353574273.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_f304d4a353574273)
		}

		switch __obf_5ca121cc54e4abf9 := __obf_24cc1bd1cbebd895.(type) {
		case xml.StartElement:
			if __obf_5ca121cc54e4abf9.Name.Local != "xml" { // 跳过根元素
				var __obf_a5620b06eee72490 string
				if __obf_f304d4a353574273 := __obf_fa962c8ba8fdfd03.DecodeElement(&__obf_a5620b06eee72490, &__obf_5ca121cc54e4abf9); __obf_f304d4a353574273 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_f304d4a353574273)
				}
				__obf_87f37beff62d945f[__obf_5ca121cc54e4abf9.Name.Local] = __obf_a5620b06eee72490
			}
		}
	}
	return __obf_87f37beff62d945f, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_2d2834515cb01894, __obf_d14f9931158b998a []byte, __obf_aa39b26ed95be518, __obf_38a7ffb13e7ddac4 string) (__obf_9820594e4fe286cd []byte, __obf_f304d4a353574273 error) {
	defer func() {
		if __obf_0da07af351193902 := recover(); __obf_0da07af351193902 != nil {
			__obf_f304d4a353574273 = fmt.Errorf("panic error: err=%v", __obf_0da07af351193902)
			return
		}
	}()
	var __obf_fdb60bda187c670a []byte
	__obf_fdb60bda187c670a, __obf_f304d4a353574273 = __obf_384b4c21777a6271(__obf_38a7ffb13e7ddac4)
	if __obf_f304d4a353574273 != nil {
		panic(__obf_f304d4a353574273)
	}
	__obf_918e5a8ab8b15df6 := AESEncryptMsg(__obf_2d2834515cb01894, __obf_d14f9931158b998a, __obf_aa39b26ed95be518, __obf_fdb60bda187c670a)
	__obf_9820594e4fe286cd = []byte(base64.StdEncoding.EncodeToString(__obf_918e5a8ab8b15df6))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_2d2834515cb01894, __obf_d14f9931158b998a []byte, __obf_aa39b26ed95be518 string, __obf_38a7ffb13e7ddac4 []byte) (__obf_918e5a8ab8b15df6 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)
	__obf_176e52d02cf1600e := 20 + len(__obf_d14f9931158b998a)
	__obf_ec3ed1cc5349a436 := __obf_176e52d02cf1600e + len(__obf_aa39b26ed95be518)
	__obf_c3e44d5cc9b64138 := BlockSize - __obf_ec3ed1cc5349a436&BlockMask
	__obf_f76a88734a45353d := __obf_ec3ed1cc5349a436 + __obf_c3e44d5cc9b64138
	__obf_81ec0832e8e550a6 := make([]byte, __obf_f76a88734a45353d)

	// 拼接
	copy(__obf_81ec0832e8e550a6[:16], __obf_2d2834515cb01894)
	__obf_ad327755ad7782fd(__obf_81ec0832e8e550a6[16:20], uint32(len(__obf_d14f9931158b998a)))
	copy(__obf_81ec0832e8e550a6[20:], __obf_d14f9931158b998a)
	copy(__obf_81ec0832e8e550a6[__obf_176e52d02cf1600e:], __obf_aa39b26ed95be518)

	// PKCS#7 补位
	for __obf_a5e6bc4d45d78213 := __obf_ec3ed1cc5349a436; __obf_a5e6bc4d45d78213 < __obf_f76a88734a45353d; __obf_a5e6bc4d45d78213++ {
		__obf_81ec0832e8e550a6[__obf_a5e6bc4d45d78213] = byte(__obf_c3e44d5cc9b64138)
	}
	__obf_e0eb1010f1f53051,

		// 加密
		__obf_f304d4a353574273 := aes.NewCipher(__obf_38a7ffb13e7ddac4)
	if __obf_f304d4a353574273 != nil {
		panic(__obf_f304d4a353574273)
	}
	__obf_c0c3ab58fce72864 := cipher.NewCBCEncrypter(__obf_e0eb1010f1f53051, __obf_38a7ffb13e7ddac4[:16])
	__obf_c0c3ab58fce72864.
		CryptBlocks(__obf_81ec0832e8e550a6, __obf_81ec0832e8e550a6)
	__obf_918e5a8ab8b15df6 = __obf_81ec0832e8e550a6
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_aa39b26ed95be518, __obf_33d82107593c551d, __obf_38a7ffb13e7ddac4 string) (__obf_2d2834515cb01894, __obf_fb8702eccc08c704 []byte, __obf_f304d4a353574273 error) {
	defer func() {
		if __obf_0da07af351193902 := recover(); __obf_0da07af351193902 != nil {
			__obf_f304d4a353574273 = fmt.Errorf("panic error: err=%v", __obf_0da07af351193902)
			return
		}
	}()
	var __obf_c38f78e6ff8cb2ff, __obf_fdb60bda187c670a, __obf_8c6546842303c475 []byte
	__obf_c38f78e6ff8cb2ff, __obf_f304d4a353574273 = base64.StdEncoding.DecodeString(__obf_33d82107593c551d)
	if __obf_f304d4a353574273 != nil {
		return
	}
	__obf_fdb60bda187c670a, __obf_f304d4a353574273 = __obf_384b4c21777a6271(__obf_38a7ffb13e7ddac4)
	if __obf_f304d4a353574273 != nil {
		panic(__obf_f304d4a353574273)
	}
	__obf_2d2834515cb01894, __obf_fb8702eccc08c704, __obf_8c6546842303c475, __obf_f304d4a353574273 = AESDecryptMsg(__obf_c38f78e6ff8cb2ff, __obf_fdb60bda187c670a)
	if __obf_f304d4a353574273 != nil {
		__obf_f304d4a353574273 = fmt.Errorf("消息解密失败,%v", __obf_f304d4a353574273)
		return
	}
	if __obf_aa39b26ed95be518 != string(__obf_8c6546842303c475) {
		__obf_f304d4a353574273 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_384b4c21777a6271(__obf_88eab50069fa4a69 string) (__obf_fdb60bda187c670a []byte, __obf_f304d4a353574273 error) {
	if len(__obf_88eab50069fa4a69) != 43 {
		__obf_f304d4a353574273 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_fdb60bda187c670a, __obf_f304d4a353574273 = base64.StdEncoding.DecodeString(__obf_88eab50069fa4a69 + "=")
	if __obf_f304d4a353574273 != nil {
		return
	}
	if len(__obf_fdb60bda187c670a) != 32 {
		__obf_f304d4a353574273 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_918e5a8ab8b15df6 []byte, __obf_38a7ffb13e7ddac4 []byte) (__obf_2d2834515cb01894, __obf_d14f9931158b998a, __obf_aa39b26ed95be518 []byte, __obf_f304d4a353574273 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_918e5a8ab8b15df6) < BlockSize {
		__obf_f304d4a353574273 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_918e5a8ab8b15df6))
		return
	}
	if len(__obf_918e5a8ab8b15df6)&BlockMask != 0 {
		__obf_f304d4a353574273 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_918e5a8ab8b15df6))
		return
	}
	__obf_81ec0832e8e550a6 := make([]byte, len(__obf_918e5a8ab8b15df6))
	__obf_e0eb1010f1f53051, // len(plaintext) >= BLOCK_SIZE
		__obf_f304d4a353574273 := // 解密
		aes.NewCipher(__obf_38a7ffb13e7ddac4)
	if __obf_f304d4a353574273 != nil {
		panic(__obf_f304d4a353574273)
	}
	__obf_c0c3ab58fce72864 := cipher.NewCBCDecrypter(__obf_e0eb1010f1f53051, __obf_38a7ffb13e7ddac4[:16])
	__obf_c0c3ab58fce72864.
		CryptBlocks(__obf_81ec0832e8e550a6, __obf_918e5a8ab8b15df6)
	__obf_c3e44d5cc9b64138 := // PKCS#7 去除补位
		int(__obf_81ec0832e8e550a6[len(__obf_81ec0832e8e550a6)-1])
	if __obf_c3e44d5cc9b64138 < 1 || __obf_c3e44d5cc9b64138 > BlockSize {
		__obf_f304d4a353574273 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_c3e44d5cc9b64138)
		return
	}
	__obf_81ec0832e8e550a6 = __obf_81ec0832e8e550a6[:len(__obf_81ec0832e8e550a6)-__obf_c3e44d5cc9b64138]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_81ec0832e8e550a6) <= 20 {
		__obf_f304d4a353574273 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_81ec0832e8e550a6))
		return
	}
	__obf_52e77300cce06106 := int(__obf_776733ed68c0cfd6(__obf_81ec0832e8e550a6[16:20]))
	if __obf_52e77300cce06106 < 0 {
		__obf_f304d4a353574273 = fmt.Errorf("incorrect msg length: %d", __obf_52e77300cce06106)
		return
	}
	__obf_176e52d02cf1600e := 20 + __obf_52e77300cce06106
	if len(__obf_81ec0832e8e550a6) <= __obf_176e52d02cf1600e {
		__obf_f304d4a353574273 = fmt.Errorf("msg length too large: %d", __obf_52e77300cce06106)
		return
	}
	__obf_2d2834515cb01894 = __obf_81ec0832e8e550a6[:16:20]
	__obf_d14f9931158b998a = __obf_81ec0832e8e550a6[20:__obf_176e52d02cf1600e:__obf_176e52d02cf1600e]
	__obf_aa39b26ed95be518 = __obf_81ec0832e8e550a6[__obf_176e52d02cf1600e:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_ad327755ad7782fd(__obf_4ebaafd192c0e344 []byte, __obf_22500dc508c88ad1 uint32) {
	__obf_4ebaafd192c0e344[0] = byte(__obf_22500dc508c88ad1 >> 24)
	__obf_4ebaafd192c0e344[1] = byte(__obf_22500dc508c88ad1 >> 16)
	__obf_4ebaafd192c0e344[2] = byte(__obf_22500dc508c88ad1 >> 8)
	__obf_4ebaafd192c0e344[3] = byte(__obf_22500dc508c88ad1)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_776733ed68c0cfd6(__obf_4ebaafd192c0e344 []byte) (__obf_22500dc508c88ad1 uint32) {
	return uint32(__obf_4ebaafd192c0e344[0])<<24 |
		uint32(__obf_4ebaafd192c0e344[1])<<16 |
		uint32(__obf_4ebaafd192c0e344[2])<<8 |
		uint32(__obf_4ebaafd192c0e344[3])
}
