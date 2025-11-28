package __obf_d984cff8712b1ee6

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

func Sha256Hex(__obf_f7b8f4634cd465ff []byte) string {
	__obf_b7e8c784ee951d79 := sha256.Sum256(__obf_f7b8f4634cd465ff)
	return hex.EncodeToString(__obf_b7e8c784ee951d79[:])
}

func HashPassword(__obf_7e75a6bf7081eda0 string) (string, error) {
	bytes, __obf_44ebf351b5f3fef8 := bcrypt.GenerateFromPassword([]byte(__obf_7e75a6bf7081eda0), bcrypt.DefaultCost)
	return string(bytes), __obf_44ebf351b5f3fef8
}

func CheckPasswordHash(__obf_7e75a6bf7081eda0, __obf_ab5b8c2b3baf441d string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_ab5b8c2b3baf441d), []byte(__obf_7e75a6bf7081eda0)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_1dbb0366e8ff201b string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	__obf_b7e8c784ee951d79 := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_b7e8c784ee951d79.Write([]byte(__obf_1dbb0366e8ff201b))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	__obf_5febfd837fc0b8bd := __obf_b7e8c784ee951d79.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_5febfd837fc0b8bd)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_43053eac9d2321a5, __obf_1c173f4352ac0b13 []byte) (__obf_4ac08e5d9beb7887 []byte, __obf_44ebf351b5f3fef8 error) {
	__obf_4869569fac6a6f82, __obf_44ebf351b5f3fef8 := aes.NewCipher(__obf_1c173f4352ac0b13)
	if __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}
	__obf_5498933ce7bb9665 := make([]byte, aes.BlockSize)
	if _, __obf_44ebf351b5f3fef8 := io.ReadFull(rand.Reader, __obf_5498933ce7bb9665); __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}
	__obf_4ac08e5d9beb7887 = make([]byte, len(__obf_43053eac9d2321a5))
	cipher.NewCBCDecrypter(__obf_4869569fac6a6f82, __obf_5498933ce7bb9665).CryptBlocks(__obf_4ac08e5d9beb7887, __obf_43053eac9d2321a5)

	return PKCS7UnPad(__obf_4ac08e5d9beb7887), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_e05a449c386ca1b9 []byte) []byte {
	__obf_910e5a981464d071 := len(__obf_e05a449c386ca1b9)
	__obf_0c9a85ca561a5b63 := int(__obf_e05a449c386ca1b9[__obf_910e5a981464d071-1])
	return __obf_e05a449c386ca1b9[:__obf_910e5a981464d071-__obf_0c9a85ca561a5b63]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_43053eac9d2321a5 []byte, __obf_1c173f4352ac0b13 []byte) ([]byte, error) {
	__obf_7bff7adf95417505 := len(__obf_1c173f4352ac0b13)
	if len(__obf_43053eac9d2321a5)%__obf_7bff7adf95417505 != 0 {
		__obf_43053eac9d2321a5 = PKCS7Pad(__obf_43053eac9d2321a5, __obf_7bff7adf95417505)
	}

	__obf_4869569fac6a6f82, __obf_44ebf351b5f3fef8 := aes.NewCipher(__obf_1c173f4352ac0b13)
	if __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}

	__obf_5498933ce7bb9665 := make([]byte, aes.BlockSize)
	if _, __obf_44ebf351b5f3fef8 := io.ReadFull(rand.Reader, __obf_5498933ce7bb9665); __obf_44ebf351b5f3fef8 != nil {
		return nil, __obf_44ebf351b5f3fef8
	}

	__obf_4ac08e5d9beb7887 := make([]byte, len(__obf_43053eac9d2321a5))
	cipher.NewCBCEncrypter(__obf_4869569fac6a6f82, __obf_5498933ce7bb9665).CryptBlocks(__obf_4ac08e5d9beb7887, __obf_43053eac9d2321a5)

	return __obf_4ac08e5d9beb7887, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_e05a449c386ca1b9 []byte, __obf_ef6bccb9518a2055 int) []byte {
	if __obf_ef6bccb9518a2055 < 1<<1 || __obf_ef6bccb9518a2055 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_0c9a85ca561a5b63 := __obf_ef6bccb9518a2055 - len(__obf_e05a449c386ca1b9)%__obf_ef6bccb9518a2055
	__obf_b9eebfe67804bbd0 := bytes.Repeat([]byte{byte(__obf_0c9a85ca561a5b63)}, __obf_0c9a85ca561a5b63)
	return append(__obf_e05a449c386ca1b9, __obf_b9eebfe67804bbd0...)
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(__obf_1dbb0366e8ff201b ...string) string {
	sort.Strings(__obf_1dbb0366e8ff201b)
	__obf_b7e8c784ee951d79 := sha1.New()
	__obf_b7e8c784ee951d79.Write([]byte(strings.Join(__obf_1dbb0366e8ff201b, "")))
	return fmt.Sprintf("%x", __obf_b7e8c784ee951d79.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_1dbb0366e8ff201b ...string) string {
	sort.Strings(__obf_1dbb0366e8ff201b)
	__obf_b7e8c784ee951d79 := md5.New()
	__obf_b7e8c784ee951d79.Write([]byte(strings.Join(__obf_1dbb0366e8ff201b, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_b7e8c784ee951d79.Sum(nil)))
}
