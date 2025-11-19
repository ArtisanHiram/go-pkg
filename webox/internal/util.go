package __obf_2e4735ec379515a2

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
	__obf_7e3be24b59506824 := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_9dc9e70a2a87d3aa := 32
	__obf_e65435718d1f82ee := make([]byte, __obf_9dc9e70a2a87d3aa)
	for __obf_af33fd20ba274c0d := range __obf_e65435718d1f82ee {
		__obf_e65435718d1f82ee[__obf_af33fd20ba274c0d] = __obf_7e3be24b59506824[rand.IntN(len(__obf_7e3be24b59506824))]
	}
	return string(__obf_e65435718d1f82ee)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_1227fb4e336bf033 string, __obf_528e037dd9f2a146 map[string]string) string {
	__obf_e09717bb1898b68d := make([]string, 0, len(__obf_528e037dd9f2a146))
	for __obf_2df09d571b10e138 := range __obf_528e037dd9f2a146 {
		if __obf_2df09d571b10e138 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_e09717bb1898b68d = append(__obf_e09717bb1898b68d, __obf_2df09d571b10e138)
	}
	sort.Strings(__obf_e09717bb1898b68d) // 字典序排序

	var __obf_a970d3d4a4832ff0 bytes.Buffer
	for _, __obf_2df09d571b10e138 := range __obf_e09717bb1898b68d {
		if __obf_528e037dd9f2a146[__obf_2df09d571b10e138] != "" { // 空值不参与签名
			__obf_a970d3d4a4832ff0.WriteString(__obf_2df09d571b10e138)
			__obf_a970d3d4a4832ff0.WriteString("=")
			__obf_a970d3d4a4832ff0.WriteString(__obf_528e037dd9f2a146[__obf_2df09d571b10e138])
			__obf_a970d3d4a4832ff0.WriteString("&")
		}
	}
	__obf_a970d3d4a4832ff0.WriteString("key=")
	__obf_a970d3d4a4832ff0.WriteString(__obf_1227fb4e336bf033)

	var __obf_26e1c99388201874 hash.Hash
	if __obf_528e037dd9f2a146["sign_type"] == SIGN_HMACSHA256 {
		__obf_26e1c99388201874 = hmac.New(sha256.New, []byte(__obf_1227fb4e336bf033))
	} else {
		__obf_26e1c99388201874 = md5.New()
	}
	__obf_26e1c99388201874.Write(__obf_a970d3d4a4832ff0.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_26e1c99388201874.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_1227fb4e336bf033 string, __obf_528e037dd9f2a146 map[string]string) bool {
	__obf_33922a11c019be4a, __obf_ae947d5c502af3d2 := __obf_528e037dd9f2a146["sign"]
	if !__obf_ae947d5c502af3d2 {
		return false // 没有sign字段
	}
	delete(__obf_528e037dd9f2a146, "sign") // 签名时需要移除sign字段

	__obf_f7146cfa29722423 := CreateSign(__obf_1227fb4e336bf033, __obf_528e037dd9f2a146)

	return __obf_f7146cfa29722423 == __obf_33922a11c019be4a
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_b7a057dcb85221de map[string]string) []byte {
	var __obf_a970d3d4a4832ff0 bytes.Buffer
	__obf_a970d3d4a4832ff0.WriteString("<xml>")
	for __obf_2df09d571b10e138, __obf_62acc15bb7c68aaa := range __obf_b7a057dcb85221de {
		__obf_a970d3d4a4832ff0.WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_2df09d571b10e138, __obf_62acc15bb7c68aaa, __obf_2df09d571b10e138))
	}
	__obf_a970d3d4a4832ff0.WriteString("</xml>")
	return __obf_a970d3d4a4832ff0.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_a5b9f7862f5cb4e2 []byte) (map[string]string, error) {
	__obf_df7f8bbd2c5d1e69 := xml.NewDecoder(bytes.NewReader(__obf_a5b9f7862f5cb4e2))
	__obf_b7a057dcb85221de := make(map[string]string)
	for {
		__obf_3f7fe469ced4c7fe, __obf_647657e89e0327f9 := __obf_df7f8bbd2c5d1e69.Token()
		if __obf_647657e89e0327f9 != nil {
			if __obf_647657e89e0327f9.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_647657e89e0327f9)
		}

		switch __obf_7649c4edadbe8ec4 := __obf_3f7fe469ced4c7fe.(type) {
		case xml.StartElement:
			if __obf_7649c4edadbe8ec4.Name.Local != "xml" { // 跳过根元素
				var __obf_4d24215b51153c14 string
				if __obf_647657e89e0327f9 := __obf_df7f8bbd2c5d1e69.DecodeElement(&__obf_4d24215b51153c14, &__obf_7649c4edadbe8ec4); __obf_647657e89e0327f9 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_647657e89e0327f9)
				}
				__obf_b7a057dcb85221de[__obf_7649c4edadbe8ec4.Name.Local] = __obf_4d24215b51153c14
			}
		}
	}
	return __obf_b7a057dcb85221de, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_327180bf9c3d25eb, __obf_01b253b195f3d410 []byte, __obf_11ab2e6fd45cd7d2, __obf_962b92503955c655 string) (__obf_de7936a6fab4f2ce []byte, __obf_647657e89e0327f9 error) {
	defer func() {
		if __obf_c7e71d64280a424a := recover(); __obf_c7e71d64280a424a != nil {
			__obf_647657e89e0327f9 = fmt.Errorf("panic error: err=%v", __obf_c7e71d64280a424a)
			return
		}
	}()
	var __obf_3e0be123922aa44c []byte
	__obf_3e0be123922aa44c, __obf_647657e89e0327f9 = __obf_35f0b7d2995aecba(__obf_962b92503955c655)
	if __obf_647657e89e0327f9 != nil {
		panic(__obf_647657e89e0327f9)
	}
	__obf_182f2879ef35efd3 := AESEncryptMsg(__obf_327180bf9c3d25eb, __obf_01b253b195f3d410, __obf_11ab2e6fd45cd7d2, __obf_3e0be123922aa44c)
	__obf_de7936a6fab4f2ce = []byte(base64.StdEncoding.EncodeToString(__obf_182f2879ef35efd3))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_327180bf9c3d25eb, __obf_01b253b195f3d410 []byte, __obf_11ab2e6fd45cd7d2 string, __obf_962b92503955c655 []byte) (__obf_182f2879ef35efd3 []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	__obf_a29de3a3146af01a := 20 + len(__obf_01b253b195f3d410)
	__obf_3b9aa4d2961dbf57 := __obf_a29de3a3146af01a + len(__obf_11ab2e6fd45cd7d2)
	__obf_e8e6f1ca020c6620 := BlockSize - __obf_3b9aa4d2961dbf57&BlockMask
	__obf_1028dc197ced18ed := __obf_3b9aa4d2961dbf57 + __obf_e8e6f1ca020c6620

	__obf_a115893c12308cec := make([]byte, __obf_1028dc197ced18ed)

	// 拼接
	copy(__obf_a115893c12308cec[:16], __obf_327180bf9c3d25eb)
	__obf_35f3435d9d7ecc94(__obf_a115893c12308cec[16:20], uint32(len(__obf_01b253b195f3d410)))
	copy(__obf_a115893c12308cec[20:], __obf_01b253b195f3d410)
	copy(__obf_a115893c12308cec[__obf_a29de3a3146af01a:], __obf_11ab2e6fd45cd7d2)

	// PKCS#7 补位
	for __obf_af33fd20ba274c0d := __obf_3b9aa4d2961dbf57; __obf_af33fd20ba274c0d < __obf_1028dc197ced18ed; __obf_af33fd20ba274c0d++ {
		__obf_a115893c12308cec[__obf_af33fd20ba274c0d] = byte(__obf_e8e6f1ca020c6620)
	}

	// 加密
	__obf_50f635d7a6a89ae3, __obf_647657e89e0327f9 := aes.NewCipher(__obf_962b92503955c655)
	if __obf_647657e89e0327f9 != nil {
		panic(__obf_647657e89e0327f9)
	}
	__obf_ca4f40e9eeb45dc3 := cipher.NewCBCEncrypter(__obf_50f635d7a6a89ae3, __obf_962b92503955c655[:16])
	__obf_ca4f40e9eeb45dc3.CryptBlocks(__obf_a115893c12308cec, __obf_a115893c12308cec)

	__obf_182f2879ef35efd3 = __obf_a115893c12308cec
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_11ab2e6fd45cd7d2, __obf_557972765cf935e7, __obf_962b92503955c655 string) (__obf_327180bf9c3d25eb, __obf_72a1216967d0a2ef []byte, __obf_647657e89e0327f9 error) {
	defer func() {
		if __obf_c7e71d64280a424a := recover(); __obf_c7e71d64280a424a != nil {
			__obf_647657e89e0327f9 = fmt.Errorf("panic error: err=%v", __obf_c7e71d64280a424a)
			return
		}
	}()
	var __obf_2a26dbe38f338838, __obf_3e0be123922aa44c, __obf_37c563de595c09c9 []byte
	__obf_2a26dbe38f338838, __obf_647657e89e0327f9 = base64.StdEncoding.DecodeString(__obf_557972765cf935e7)
	if __obf_647657e89e0327f9 != nil {
		return
	}
	__obf_3e0be123922aa44c, __obf_647657e89e0327f9 = __obf_35f0b7d2995aecba(__obf_962b92503955c655)
	if __obf_647657e89e0327f9 != nil {
		panic(__obf_647657e89e0327f9)
	}
	__obf_327180bf9c3d25eb, __obf_72a1216967d0a2ef, __obf_37c563de595c09c9, __obf_647657e89e0327f9 = AESDecryptMsg(__obf_2a26dbe38f338838, __obf_3e0be123922aa44c)
	if __obf_647657e89e0327f9 != nil {
		__obf_647657e89e0327f9 = fmt.Errorf("消息解密失败,%v", __obf_647657e89e0327f9)
		return
	}
	if __obf_11ab2e6fd45cd7d2 != string(__obf_37c563de595c09c9) {
		__obf_647657e89e0327f9 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_35f0b7d2995aecba(__obf_df3f77f5b281326b string) (__obf_3e0be123922aa44c []byte, __obf_647657e89e0327f9 error) {
	if len(__obf_df3f77f5b281326b) != 43 {
		__obf_647657e89e0327f9 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_3e0be123922aa44c, __obf_647657e89e0327f9 = base64.StdEncoding.DecodeString(__obf_df3f77f5b281326b + "=")
	if __obf_647657e89e0327f9 != nil {
		return
	}
	if len(__obf_3e0be123922aa44c) != 32 {
		__obf_647657e89e0327f9 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_182f2879ef35efd3 []byte, __obf_962b92503955c655 []byte) (__obf_327180bf9c3d25eb, __obf_01b253b195f3d410, __obf_11ab2e6fd45cd7d2 []byte, __obf_647657e89e0327f9 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_182f2879ef35efd3) < BlockSize {
		__obf_647657e89e0327f9 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_182f2879ef35efd3))
		return
	}
	if len(__obf_182f2879ef35efd3)&BlockMask != 0 {
		__obf_647657e89e0327f9 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_182f2879ef35efd3))
		return
	}

	__obf_a115893c12308cec := make([]byte, len(__obf_182f2879ef35efd3)) // len(plaintext) >= BLOCK_SIZE

	// 解密
	__obf_50f635d7a6a89ae3, __obf_647657e89e0327f9 := aes.NewCipher(__obf_962b92503955c655)
	if __obf_647657e89e0327f9 != nil {
		panic(__obf_647657e89e0327f9)
	}
	__obf_ca4f40e9eeb45dc3 := cipher.NewCBCDecrypter(__obf_50f635d7a6a89ae3, __obf_962b92503955c655[:16])
	__obf_ca4f40e9eeb45dc3.CryptBlocks(__obf_a115893c12308cec, __obf_182f2879ef35efd3)

	// PKCS#7 去除补位
	__obf_e8e6f1ca020c6620 := int(__obf_a115893c12308cec[len(__obf_a115893c12308cec)-1])
	if __obf_e8e6f1ca020c6620 < 1 || __obf_e8e6f1ca020c6620 > BlockSize {
		__obf_647657e89e0327f9 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_e8e6f1ca020c6620)
		return
	}
	__obf_a115893c12308cec = __obf_a115893c12308cec[:len(__obf_a115893c12308cec)-__obf_e8e6f1ca020c6620]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_a115893c12308cec) <= 20 {
		__obf_647657e89e0327f9 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_a115893c12308cec))
		return
	}
	__obf_200a91e8e1857fde := int(__obf_afc91166319b24bf(__obf_a115893c12308cec[16:20]))
	if __obf_200a91e8e1857fde < 0 {
		__obf_647657e89e0327f9 = fmt.Errorf("incorrect msg length: %d", __obf_200a91e8e1857fde)
		return
	}
	__obf_a29de3a3146af01a := 20 + __obf_200a91e8e1857fde
	if len(__obf_a115893c12308cec) <= __obf_a29de3a3146af01a {
		__obf_647657e89e0327f9 = fmt.Errorf("msg length too large: %d", __obf_200a91e8e1857fde)
		return
	}

	__obf_327180bf9c3d25eb = __obf_a115893c12308cec[:16:20]
	__obf_01b253b195f3d410 = __obf_a115893c12308cec[20:__obf_a29de3a3146af01a:__obf_a29de3a3146af01a]
	__obf_11ab2e6fd45cd7d2 = __obf_a115893c12308cec[__obf_a29de3a3146af01a:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_35f3435d9d7ecc94(__obf_feba684fe02cc951 []byte, __obf_d0093a92adf68d39 uint32) {
	__obf_feba684fe02cc951[0] = byte(__obf_d0093a92adf68d39 >> 24)
	__obf_feba684fe02cc951[1] = byte(__obf_d0093a92adf68d39 >> 16)
	__obf_feba684fe02cc951[2] = byte(__obf_d0093a92adf68d39 >> 8)
	__obf_feba684fe02cc951[3] = byte(__obf_d0093a92adf68d39)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_afc91166319b24bf(__obf_feba684fe02cc951 []byte) (__obf_d0093a92adf68d39 uint32) {
	return uint32(__obf_feba684fe02cc951[0])<<24 |
		uint32(__obf_feba684fe02cc951[1])<<16 |
		uint32(__obf_feba684fe02cc951[2])<<8 |
		uint32(__obf_feba684fe02cc951[3])
}
