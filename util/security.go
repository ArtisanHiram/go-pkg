package __obf_d7b39e56b82f7f57

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

func Sha256Hex(__obf_b2cb60631f053d14 []byte) string {
	__obf_b38e75999c1d89f0 := sha256.Sum256(__obf_b2cb60631f053d14)
	return hex.EncodeToString(__obf_b38e75999c1d89f0[:])
}

func HashPassword(__obf_035d148c6cb6813a string) (string, error) {
	bytes, __obf_3ccf8c32165eeb13 := bcrypt.GenerateFromPassword([]byte(__obf_035d148c6cb6813a), bcrypt.DefaultCost)
	return string(bytes), __obf_3ccf8c32165eeb13
}

func CheckPasswordHash(__obf_035d148c6cb6813a, __obf_bf2537107562b472 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_bf2537107562b472), []byte(__obf_035d148c6cb6813a)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_40384ad7789a1f57 string) string {
	__obf_b38e75999c1d89f0 := // The pattern for generating a hash is `sha1.New()`,
		// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
		// Here we start with a new hash.
		sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_b38e75999c1d89f0.Write([]byte(__obf_40384ad7789a1f57))
	__obf_08f48d35deaf11d1 := // This gets the finalized hash result as a byte
		// slice. The argument to `Sum` can be used to append
		// to an existing byte slice: it usually isn't needed.
		__obf_b38e75999c1d89f0.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_08f48d35deaf11d1)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_def516632b2b7340, __obf_6517b70edd186159 []byte) (__obf_2fb73aaa77231bda []byte, __obf_3ccf8c32165eeb13 error) {
	__obf_573a69ed53eea7db, __obf_3ccf8c32165eeb13 := aes.NewCipher(__obf_6517b70edd186159)
	if __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	__obf_c4b76ed17c888396 := make([]byte, aes.BlockSize)
	if _, __obf_3ccf8c32165eeb13 := io.ReadFull(rand.Reader, __obf_c4b76ed17c888396); __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	__obf_2fb73aaa77231bda = make([]byte, len(__obf_def516632b2b7340))
	cipher.NewCBCDecrypter(__obf_573a69ed53eea7db, __obf_c4b76ed17c888396).CryptBlocks(__obf_2fb73aaa77231bda, __obf_def516632b2b7340)

	return PKCS7UnPad(__obf_2fb73aaa77231bda), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_a92a4b7a95d20a70 []byte) []byte {
	__obf_3e19bf522273cec0 := len(__obf_a92a4b7a95d20a70)
	__obf_df133f50810b999e := int(__obf_a92a4b7a95d20a70[__obf_3e19bf522273cec0-1])
	return __obf_a92a4b7a95d20a70[:__obf_3e19bf522273cec0-__obf_df133f50810b999e]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_def516632b2b7340 []byte, __obf_6517b70edd186159 []byte) ([]byte, error) {
	__obf_9045b7561a7c656e := len(__obf_6517b70edd186159)
	if len(__obf_def516632b2b7340)%__obf_9045b7561a7c656e != 0 {
		__obf_def516632b2b7340 = PKCS7Pad(__obf_def516632b2b7340, __obf_9045b7561a7c656e)
	}
	__obf_573a69ed53eea7db, __obf_3ccf8c32165eeb13 := aes.NewCipher(__obf_6517b70edd186159)
	if __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	__obf_c4b76ed17c888396 := make([]byte, aes.BlockSize)
	if _, __obf_3ccf8c32165eeb13 := io.ReadFull(rand.Reader, __obf_c4b76ed17c888396); __obf_3ccf8c32165eeb13 != nil {
		return nil, __obf_3ccf8c32165eeb13
	}
	__obf_2fb73aaa77231bda := make([]byte, len(__obf_def516632b2b7340))
	cipher.NewCBCEncrypter(__obf_573a69ed53eea7db, __obf_c4b76ed17c888396).CryptBlocks(__obf_2fb73aaa77231bda, __obf_def516632b2b7340)

	return __obf_2fb73aaa77231bda, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_a92a4b7a95d20a70 []byte, __obf_107bc9db8db14da4 int) []byte {
	if __obf_107bc9db8db14da4 < 1<<1 || __obf_107bc9db8db14da4 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_df133f50810b999e := __obf_107bc9db8db14da4 - len(__obf_a92a4b7a95d20a70)%__obf_107bc9db8db14da4
	__obf_b1af470cb3c49052 := bytes.Repeat([]byte{byte(__obf_df133f50810b999e)}, __obf_df133f50810b999e)
	return append(__obf_a92a4b7a95d20a70,

		// SortSha1 排序并sha1，主要用于计算signature
		__obf_b1af470cb3c49052...)
}

func SortSha1(__obf_40384ad7789a1f57 ...string) string {
	sort.Strings(__obf_40384ad7789a1f57)
	__obf_b38e75999c1d89f0 := sha1.New()
	__obf_b38e75999c1d89f0.
		Write([]byte(strings.Join(__obf_40384ad7789a1f57, "")))
	return fmt.Sprintf("%x", __obf_b38e75999c1d89f0.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_40384ad7789a1f57 ...string) string {
	sort.Strings(__obf_40384ad7789a1f57)
	__obf_b38e75999c1d89f0 := md5.New()
	__obf_b38e75999c1d89f0.
		Write([]byte(strings.Join(__obf_40384ad7789a1f57, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_b38e75999c1d89f0.Sum(nil)))
}
