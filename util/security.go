package __obf_426da37e60cac670

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Sha256Hex(__obf_1871a36bb793083c []byte) string {
	__obf_6d8eeb12774a8172 := sha256.Sum256(__obf_1871a36bb793083c)
	return hex.EncodeToString(__obf_6d8eeb12774a8172[:])
}

func HashPassword(__obf_7c04f6c9dd1626ad string) (string, error) {
	bytes, __obf_74916b80241ef1ff := bcrypt.GenerateFromPassword([]byte(__obf_7c04f6c9dd1626ad), bcrypt.DefaultCost)
	return string(bytes), __obf_74916b80241ef1ff
}

func CheckPasswordHash(__obf_7c04f6c9dd1626ad, __obf_0ac404817cf3f3ab string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_0ac404817cf3f3ab), []byte(__obf_7c04f6c9dd1626ad)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_b69dd8ffbe131543 string) string {
	__obf_6d8eeb12774a8172 := // The pattern for generating a hash is `sha1.New()`,
		// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
		// Here we start with a new hash.
		sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_6d8eeb12774a8172.Write([]byte(__obf_b69dd8ffbe131543))
	__obf_d4dc4b2dbd842180 := // This gets the finalized hash result as a byte
		// slice. The argument to `Sum` can be used to append
		// to an existing byte slice: it usually isn't needed.
		__obf_6d8eeb12774a8172.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_d4dc4b2dbd842180)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_18b7ce7a2be7b8cb, __obf_fe32fb81ebf474d0 []byte) (__obf_eac92198fd2cd08f []byte, __obf_74916b80241ef1ff error) {
	__obf_fd2c8e7c67255446, __obf_74916b80241ef1ff := aes.NewCipher(__obf_fe32fb81ebf474d0)
	if __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	__obf_721da1fff0595313 := make([]byte, aes.BlockSize)
	if _, __obf_74916b80241ef1ff := io.ReadFull(rand.Reader, __obf_721da1fff0595313); __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	__obf_eac92198fd2cd08f = make([]byte, len(__obf_18b7ce7a2be7b8cb))
	cipher.NewCBCDecrypter(__obf_fd2c8e7c67255446, __obf_721da1fff0595313).CryptBlocks(__obf_eac92198fd2cd08f, __obf_18b7ce7a2be7b8cb)

	return PKCS7UnPad(__obf_eac92198fd2cd08f), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_d9279f250666f55f []byte) []byte {
	__obf_a5eeba35e2de3b47 := len(__obf_d9279f250666f55f)
	__obf_385fab5f579dbc67 := int(__obf_d9279f250666f55f[__obf_a5eeba35e2de3b47-1])
	return __obf_d9279f250666f55f[:__obf_a5eeba35e2de3b47-__obf_385fab5f579dbc67]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_18b7ce7a2be7b8cb []byte, __obf_fe32fb81ebf474d0 []byte) ([]byte, error) {
	__obf_35fc606f1e48ca40 := len(__obf_fe32fb81ebf474d0)
	if len(__obf_18b7ce7a2be7b8cb)%__obf_35fc606f1e48ca40 != 0 {
		__obf_18b7ce7a2be7b8cb = PKCS7Pad(__obf_18b7ce7a2be7b8cb, __obf_35fc606f1e48ca40)
	}
	__obf_fd2c8e7c67255446, __obf_74916b80241ef1ff := aes.NewCipher(__obf_fe32fb81ebf474d0)
	if __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	__obf_721da1fff0595313 := make([]byte, aes.BlockSize)
	if _, __obf_74916b80241ef1ff := io.ReadFull(rand.Reader, __obf_721da1fff0595313); __obf_74916b80241ef1ff != nil {
		return nil, __obf_74916b80241ef1ff
	}
	__obf_eac92198fd2cd08f := make([]byte, len(__obf_18b7ce7a2be7b8cb))
	cipher.NewCBCEncrypter(__obf_fd2c8e7c67255446, __obf_721da1fff0595313).CryptBlocks(__obf_eac92198fd2cd08f, __obf_18b7ce7a2be7b8cb)

	return __obf_eac92198fd2cd08f, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_d9279f250666f55f []byte, __obf_b75290ba126e4cfd int) []byte {
	if __obf_b75290ba126e4cfd < 1<<1 || __obf_b75290ba126e4cfd >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_385fab5f579dbc67 := __obf_b75290ba126e4cfd - len(__obf_d9279f250666f55f)%__obf_b75290ba126e4cfd
	__obf_9623a5ae55b6829c := bytes.Repeat([]byte{byte(__obf_385fab5f579dbc67)}, __obf_385fab5f579dbc67)
	return append(__obf_d9279f250666f55f,

		// SortSha1 排序并sha1，主要用于计算signature
		__obf_9623a5ae55b6829c...)
}

func SortSha1(__obf_b69dd8ffbe131543 ...string) string {
	sort.Strings(__obf_b69dd8ffbe131543)
	__obf_6d8eeb12774a8172 := sha1.New()
	__obf_6d8eeb12774a8172.
		Write([]byte(strings.Join(__obf_b69dd8ffbe131543, "")))
	return fmt.Sprintf("%x", __obf_6d8eeb12774a8172.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_b69dd8ffbe131543 ...string) string {
	sort.Strings(__obf_b69dd8ffbe131543)
	__obf_6d8eeb12774a8172 := md5.New()
	__obf_6d8eeb12774a8172.
		Write([]byte(strings.Join(__obf_b69dd8ffbe131543, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_6d8eeb12774a8172.Sum(nil)))
}
