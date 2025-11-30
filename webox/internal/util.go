package __obf_abe10294567bade2

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
	__obf_dd50ddf79e12c9be := "abcdefghijklmnopqrstuvwxyz0123456789"
	__obf_c4d3ca237ede8853 := 32
	__obf_4c1868560aff3869 := make([]byte, __obf_c4d3ca237ede8853)
	for __obf_8324582c8d29e92f := range __obf_4c1868560aff3869 {
		__obf_4c1868560aff3869[__obf_8324582c8d29e92f] = __obf_dd50ddf79e12c9be[rand.IntN(len(__obf_dd50ddf79e12c9be))]
	}
	return string(__obf_4c1868560aff3869)
}

// CreateSign 为微信支付请求生成签名
// 签名规则：
// 1. 参数名ASCII码从小到大排序（字典序）
// 2. 使用URL键值对的格式（即key1=value1&key2=value2...）拼接成字符串
// 3. 在字符串最后拼接上 &key=APIKey
// 4. 对拼接好的字符串进行MD5加密，得到大写签名
func CreateSign(__obf_6ac433a3ac03fe1b string, __obf_4c338de14b6a6627 map[string]string) string {
	__obf_af373e8aecc99f30 := make([]string, 0, len(__obf_4c338de14b6a6627))
	for __obf_685e68f3451766f3 := range __obf_4c338de14b6a6627 {
		if __obf_685e68f3451766f3 == "sign" { // 签名参数不参与签名计算
			continue
		}
		__obf_af373e8aecc99f30 = append(__obf_af373e8aecc99f30, __obf_685e68f3451766f3)
	}
	sort.Strings(__obf_af373e8aecc99f30) // 字典序排序

	var __obf_a3c74837d6f72b93 bytes.Buffer
	for _, __obf_685e68f3451766f3 := range __obf_af373e8aecc99f30 {
		if __obf_4c338de14b6a6627[__obf_685e68f3451766f3] != "" {
			__obf_a3c74837d6f72b93. // 空值不参与签名
						WriteString(__obf_685e68f3451766f3)
			__obf_a3c74837d6f72b93.
				WriteString("=")
			__obf_a3c74837d6f72b93.
				WriteString(__obf_4c338de14b6a6627[__obf_685e68f3451766f3])
			__obf_a3c74837d6f72b93.
				WriteString("&")
		}
	}
	__obf_a3c74837d6f72b93.
		WriteString("key=")
	__obf_a3c74837d6f72b93.
		WriteString(__obf_6ac433a3ac03fe1b)

	var __obf_417107f550f14e2c hash.Hash
	if __obf_4c338de14b6a6627["sign_type"] == SIGN_HMACSHA256 {
		__obf_417107f550f14e2c = hmac.New(sha256.New, []byte(__obf_6ac433a3ac03fe1b))
	} else {
		__obf_417107f550f14e2c = md5.New()
	}
	__obf_417107f550f14e2c.
		Write(__obf_a3c74837d6f72b93.Bytes())
	// fmt.Println("创建签名：", buf.String())
	return strings.ToUpper(hex.EncodeToString(__obf_417107f550f14e2c.Sum(nil)))
}

// VerifySign 验证微信支付响应或回调的签名
// params: 包含所有参数的map，其中应包含"sign"字段
// apiKey: 商户API密钥
func VerifySign(__obf_6ac433a3ac03fe1b string, __obf_4c338de14b6a6627 map[string]string) bool {
	__obf_52c86176472b73fd, __obf_6e392fab483ffd08 := __obf_4c338de14b6a6627["sign"]
	if !__obf_6e392fab483ffd08 {
		return false // 没有sign字段
	}
	delete(__obf_4c338de14b6a6627, "sign")
	__obf_6c89934e64d4c833 := // 签名时需要移除sign字段

		CreateSign(__obf_6ac433a3ac03fe1b, __obf_4c338de14b6a6627)

	return __obf_6c89934e64d4c833 == __obf_52c86176472b73fd
}

// MapToXML 使用map来构建XML请求，适用于灵活的参数
func MapToXML(__obf_ee2c872ef85c5b21 map[string]string) []byte {
	var __obf_a3c74837d6f72b93 bytes.Buffer
	__obf_a3c74837d6f72b93.
		WriteString("<xml>")
	for __obf_685e68f3451766f3, __obf_7ef4adf3e75f965e := range __obf_ee2c872ef85c5b21 {
		__obf_a3c74837d6f72b93.
			WriteString(fmt.Sprintf("<%s><![CDATA[%s]]></%s>", __obf_685e68f3451766f3, __obf_7ef4adf3e75f965e, __obf_685e68f3451766f3))
	}
	__obf_a3c74837d6f72b93.
		WriteString("</xml>")
	return __obf_a3c74837d6f72b93.Bytes()
}

// XMLToMap 将XML字符串解析为map[string]string
func XMLToMap(__obf_b8c9404f9cff1593 []byte) (map[string]string, error) {
	__obf_8ec4b4e87b316fd4 := xml.NewDecoder(bytes.NewReader(__obf_b8c9404f9cff1593))
	__obf_ee2c872ef85c5b21 := make(map[string]string)
	for {
		__obf_32b80744ca58c61c, __obf_02af51ab07084000 := __obf_8ec4b4e87b316fd4.Token()
		if __obf_02af51ab07084000 != nil {
			if __obf_02af51ab07084000.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("decode XML token failed: %w", __obf_02af51ab07084000)
		}

		switch __obf_047bad33031b4f38 := __obf_32b80744ca58c61c.(type) {
		case xml.StartElement:
			if __obf_047bad33031b4f38.Name.Local != "xml" { // 跳过根元素
				var __obf_f30b84aa25ea2632 string
				if __obf_02af51ab07084000 := __obf_8ec4b4e87b316fd4.DecodeElement(&__obf_f30b84aa25ea2632, &__obf_047bad33031b4f38); __obf_02af51ab07084000 != nil {
					return nil, fmt.Errorf("decode XML element failed: %w", __obf_02af51ab07084000)
				}
				__obf_ee2c872ef85c5b21[__obf_047bad33031b4f38.Name.Local] = __obf_f30b84aa25ea2632
			}
		}
	}
	return __obf_ee2c872ef85c5b21, nil
}

// EncryptMsg 加密消息
func EncryptMsg(__obf_7d4f90ac6ebea097, __obf_37068d23d996d5dd []byte, __obf_d7d23d6d36d0093b, __obf_7db5d45df7d59e10 string) (__obf_e8feb47b77130570 []byte, __obf_02af51ab07084000 error) {
	defer func() {
		if __obf_2f7f340198b510d8 := recover(); __obf_2f7f340198b510d8 != nil {
			__obf_02af51ab07084000 = fmt.Errorf("panic error: err=%v", __obf_2f7f340198b510d8)
			return
		}
	}()
	var __obf_2c9789f1f6d4df1a []byte
	__obf_2c9789f1f6d4df1a, __obf_02af51ab07084000 = __obf_96e615ccbec8449f(__obf_7db5d45df7d59e10)
	if __obf_02af51ab07084000 != nil {
		panic(__obf_02af51ab07084000)
	}
	__obf_a1ae217992a1ccec := AESEncryptMsg(__obf_7d4f90ac6ebea097, __obf_37068d23d996d5dd, __obf_d7d23d6d36d0093b, __obf_2c9789f1f6d4df1a)
	__obf_e8feb47b77130570 = []byte(base64.StdEncoding.EncodeToString(__obf_a1ae217992a1ccec))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESEncryptMsg(__obf_7d4f90ac6ebea097, __obf_37068d23d996d5dd []byte, __obf_d7d23d6d36d0093b string, __obf_7db5d45df7d59e10 []byte) (__obf_a1ae217992a1ccec []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)
	__obf_ce623e6926cb682c := 20 + len(__obf_37068d23d996d5dd)
	__obf_289986fbeb4beee1 := __obf_ce623e6926cb682c + len(__obf_d7d23d6d36d0093b)
	__obf_2eb343c77f621d55 := BlockSize - __obf_289986fbeb4beee1&BlockMask
	__obf_76ce1c6fb444b944 := __obf_289986fbeb4beee1 + __obf_2eb343c77f621d55
	__obf_dd5032c155e4a98e := make([]byte, __obf_76ce1c6fb444b944)

	// 拼接
	copy(__obf_dd5032c155e4a98e[:16], __obf_7d4f90ac6ebea097)
	__obf_6a3807747bdb5477(__obf_dd5032c155e4a98e[16:20], uint32(len(__obf_37068d23d996d5dd)))
	copy(__obf_dd5032c155e4a98e[20:], __obf_37068d23d996d5dd)
	copy(__obf_dd5032c155e4a98e[__obf_ce623e6926cb682c:], __obf_d7d23d6d36d0093b)

	// PKCS#7 补位
	for __obf_8324582c8d29e92f := __obf_289986fbeb4beee1; __obf_8324582c8d29e92f < __obf_76ce1c6fb444b944; __obf_8324582c8d29e92f++ {
		__obf_dd5032c155e4a98e[__obf_8324582c8d29e92f] = byte(__obf_2eb343c77f621d55)
	}
	__obf_7c87d86ff2b2ece6,

		// 加密
		__obf_02af51ab07084000 := aes.NewCipher(__obf_7db5d45df7d59e10)
	if __obf_02af51ab07084000 != nil {
		panic(__obf_02af51ab07084000)
	}
	__obf_1d4f4676515a51ea := cipher.NewCBCEncrypter(__obf_7c87d86ff2b2ece6, __obf_7db5d45df7d59e10[:16])
	__obf_1d4f4676515a51ea.
		CryptBlocks(__obf_dd5032c155e4a98e, __obf_dd5032c155e4a98e)
	__obf_a1ae217992a1ccec = __obf_dd5032c155e4a98e
	return
}

// DecryptMsg 消息解密
func DecryptMsg(__obf_d7d23d6d36d0093b, __obf_02f4d34adc54c238, __obf_7db5d45df7d59e10 string) (__obf_7d4f90ac6ebea097, __obf_f303784bb3a48954 []byte, __obf_02af51ab07084000 error) {
	defer func() {
		if __obf_2f7f340198b510d8 := recover(); __obf_2f7f340198b510d8 != nil {
			__obf_02af51ab07084000 = fmt.Errorf("panic error: err=%v", __obf_2f7f340198b510d8)
			return
		}
	}()
	var __obf_a50c7651df5c743a, __obf_2c9789f1f6d4df1a, __obf_368abe72164def5b []byte
	__obf_a50c7651df5c743a, __obf_02af51ab07084000 = base64.StdEncoding.DecodeString(__obf_02f4d34adc54c238)
	if __obf_02af51ab07084000 != nil {
		return
	}
	__obf_2c9789f1f6d4df1a, __obf_02af51ab07084000 = __obf_96e615ccbec8449f(__obf_7db5d45df7d59e10)
	if __obf_02af51ab07084000 != nil {
		panic(__obf_02af51ab07084000)
	}
	__obf_7d4f90ac6ebea097, __obf_f303784bb3a48954, __obf_368abe72164def5b, __obf_02af51ab07084000 = AESDecryptMsg(__obf_a50c7651df5c743a, __obf_2c9789f1f6d4df1a)
	if __obf_02af51ab07084000 != nil {
		__obf_02af51ab07084000 = fmt.Errorf("消息解密失败,%v", __obf_02af51ab07084000)
		return
	}
	if __obf_d7d23d6d36d0093b != string(__obf_368abe72164def5b) {
		__obf_02af51ab07084000 = fmt.Errorf("消息解密校验APPID失败")
		return
	}
	return
}

func __obf_96e615ccbec8449f(__obf_b80cec427f5543d5 string) (__obf_2c9789f1f6d4df1a []byte, __obf_02af51ab07084000 error) {
	if len(__obf_b80cec427f5543d5) != 43 {
		__obf_02af51ab07084000 = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	__obf_2c9789f1f6d4df1a, __obf_02af51ab07084000 = base64.StdEncoding.DecodeString(__obf_b80cec427f5543d5 + "=")
	if __obf_02af51ab07084000 != nil {
		return
	}
	if len(__obf_2c9789f1f6d4df1a) != 32 {
		__obf_02af51ab07084000 = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
// 参考：github.com/chanxuehong/wechat.v2
func AESDecryptMsg(__obf_a1ae217992a1ccec []byte, __obf_7db5d45df7d59e10 []byte) (__obf_7d4f90ac6ebea097, __obf_37068d23d996d5dd, __obf_d7d23d6d36d0093b []byte, __obf_02af51ab07084000 error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(__obf_a1ae217992a1ccec) < BlockSize {
		__obf_02af51ab07084000 = fmt.Errorf("the length of ciphertext too short: %d", len(__obf_a1ae217992a1ccec))
		return
	}
	if len(__obf_a1ae217992a1ccec)&BlockMask != 0 {
		__obf_02af51ab07084000 = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(__obf_a1ae217992a1ccec))
		return
	}
	__obf_dd5032c155e4a98e := make([]byte, len(__obf_a1ae217992a1ccec))
	__obf_7c87d86ff2b2ece6, // len(plaintext) >= BLOCK_SIZE
		__obf_02af51ab07084000 := // 解密
		aes.NewCipher(__obf_7db5d45df7d59e10)
	if __obf_02af51ab07084000 != nil {
		panic(__obf_02af51ab07084000)
	}
	__obf_1d4f4676515a51ea := cipher.NewCBCDecrypter(__obf_7c87d86ff2b2ece6, __obf_7db5d45df7d59e10[:16])
	__obf_1d4f4676515a51ea.
		CryptBlocks(__obf_dd5032c155e4a98e, __obf_a1ae217992a1ccec)
	__obf_2eb343c77f621d55 := // PKCS#7 去除补位
		int(__obf_dd5032c155e4a98e[len(__obf_dd5032c155e4a98e)-1])
	if __obf_2eb343c77f621d55 < 1 || __obf_2eb343c77f621d55 > BlockSize {
		__obf_02af51ab07084000 = fmt.Errorf("the amount to pad is incorrect: %d", __obf_2eb343c77f621d55)
		return
	}
	__obf_dd5032c155e4a98e = __obf_dd5032c155e4a98e[:len(__obf_dd5032c155e4a98e)-__obf_2eb343c77f621d55]

	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(__obf_dd5032c155e4a98e) <= 20 {
		__obf_02af51ab07084000 = fmt.Errorf("plaintext too short, the length is %d", len(__obf_dd5032c155e4a98e))
		return
	}
	__obf_27f666aff77b13a9 := int(__obf_79396da1185205f7(__obf_dd5032c155e4a98e[16:20]))
	if __obf_27f666aff77b13a9 < 0 {
		__obf_02af51ab07084000 = fmt.Errorf("incorrect msg length: %d", __obf_27f666aff77b13a9)
		return
	}
	__obf_ce623e6926cb682c := 20 + __obf_27f666aff77b13a9
	if len(__obf_dd5032c155e4a98e) <= __obf_ce623e6926cb682c {
		__obf_02af51ab07084000 = fmt.Errorf("msg length too large: %d", __obf_27f666aff77b13a9)
		return
	}
	__obf_7d4f90ac6ebea097 = __obf_dd5032c155e4a98e[:16:20]
	__obf_37068d23d996d5dd = __obf_dd5032c155e4a98e[20:__obf_ce623e6926cb682c:__obf_ce623e6926cb682c]
	__obf_d7d23d6d36d0093b = __obf_dd5032c155e4a98e[__obf_ce623e6926cb682c:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func __obf_6a3807747bdb5477(__obf_9ef7c9385ad4f6b5 []byte, __obf_c539d66e88884c35 uint32) {
	__obf_9ef7c9385ad4f6b5[0] = byte(__obf_c539d66e88884c35 >> 24)
	__obf_9ef7c9385ad4f6b5[1] = byte(__obf_c539d66e88884c35 >> 16)
	__obf_9ef7c9385ad4f6b5[2] = byte(__obf_c539d66e88884c35 >> 8)
	__obf_9ef7c9385ad4f6b5[3] = byte(__obf_c539d66e88884c35)
}

// 从 4 字节的网络字节序里解析出整数
func __obf_79396da1185205f7(__obf_9ef7c9385ad4f6b5 []byte) (__obf_c539d66e88884c35 uint32) {
	return uint32(__obf_9ef7c9385ad4f6b5[0])<<24 |
		uint32(__obf_9ef7c9385ad4f6b5[1])<<16 |
		uint32(__obf_9ef7c9385ad4f6b5[2])<<8 |
		uint32(__obf_9ef7c9385ad4f6b5[3])
}
