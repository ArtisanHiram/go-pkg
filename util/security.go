package __obf_f1e346e3aa5cc554

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

func Sha256Hex(__obf_44c24162abb79539 []byte) string {
	__obf_51260656a7ac8feb := sha256.Sum256(__obf_44c24162abb79539)
	return hex.EncodeToString(__obf_51260656a7ac8feb[:])
}

func HashPassword(__obf_d502aa660a734925 string) (string, error) {
	bytes, __obf_eec784b359ebf42f := bcrypt.GenerateFromPassword([]byte(__obf_d502aa660a734925), bcrypt.DefaultCost)
	return string(bytes), __obf_eec784b359ebf42f
}

func CheckPasswordHash(__obf_d502aa660a734925, __obf_6f5aede0472f1ac1 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_6f5aede0472f1ac1), []byte(__obf_d502aa660a734925)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_0115cb118880283a string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	__obf_51260656a7ac8feb := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_51260656a7ac8feb.Write([]byte(__obf_0115cb118880283a))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	__obf_5180a77656aa2d9e := __obf_51260656a7ac8feb.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_5180a77656aa2d9e)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_ac1f80c958d4d9f8, __obf_9a20564fc99ec5cb []byte) (__obf_eaf2753e8a77b6e9 []byte, __obf_eec784b359ebf42f error) {
	__obf_1f0d52f8190d8bbf, __obf_eec784b359ebf42f := aes.NewCipher(__obf_9a20564fc99ec5cb)
	if __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}
	__obf_bea1885fa152433c := make([]byte, aes.BlockSize)
	if _, __obf_eec784b359ebf42f := io.ReadFull(rand.Reader, __obf_bea1885fa152433c); __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}
	__obf_eaf2753e8a77b6e9 = make([]byte, len(__obf_ac1f80c958d4d9f8))
	cipher.NewCBCDecrypter(__obf_1f0d52f8190d8bbf, __obf_bea1885fa152433c).CryptBlocks(__obf_eaf2753e8a77b6e9, __obf_ac1f80c958d4d9f8)

	return PKCS7UnPad(__obf_eaf2753e8a77b6e9), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_b7fc63f57d283f31 []byte) []byte {
	__obf_60e733d31111c611 := len(__obf_b7fc63f57d283f31)
	__obf_13165d2fdc17a265 := int(__obf_b7fc63f57d283f31[__obf_60e733d31111c611-1])
	return __obf_b7fc63f57d283f31[:__obf_60e733d31111c611-__obf_13165d2fdc17a265]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_ac1f80c958d4d9f8 []byte, __obf_9a20564fc99ec5cb []byte) ([]byte, error) {
	__obf_6126311d2a895ba1 := len(__obf_9a20564fc99ec5cb)
	if len(__obf_ac1f80c958d4d9f8)%__obf_6126311d2a895ba1 != 0 {
		__obf_ac1f80c958d4d9f8 = PKCS7Pad(__obf_ac1f80c958d4d9f8, __obf_6126311d2a895ba1)
	}

	__obf_1f0d52f8190d8bbf, __obf_eec784b359ebf42f := aes.NewCipher(__obf_9a20564fc99ec5cb)
	if __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}

	__obf_bea1885fa152433c := make([]byte, aes.BlockSize)
	if _, __obf_eec784b359ebf42f := io.ReadFull(rand.Reader, __obf_bea1885fa152433c); __obf_eec784b359ebf42f != nil {
		return nil, __obf_eec784b359ebf42f
	}

	__obf_eaf2753e8a77b6e9 := make([]byte, len(__obf_ac1f80c958d4d9f8))
	cipher.NewCBCEncrypter(__obf_1f0d52f8190d8bbf, __obf_bea1885fa152433c).CryptBlocks(__obf_eaf2753e8a77b6e9, __obf_ac1f80c958d4d9f8)

	return __obf_eaf2753e8a77b6e9, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_b7fc63f57d283f31 []byte, __obf_a2ce69636b8ce0a0 int) []byte {
	if __obf_a2ce69636b8ce0a0 < 1<<1 || __obf_a2ce69636b8ce0a0 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_13165d2fdc17a265 := __obf_a2ce69636b8ce0a0 - len(__obf_b7fc63f57d283f31)%__obf_a2ce69636b8ce0a0
	__obf_0f00058cd1b32208 := bytes.Repeat([]byte{byte(__obf_13165d2fdc17a265)}, __obf_13165d2fdc17a265)
	return append(__obf_b7fc63f57d283f31, __obf_0f00058cd1b32208...)
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(__obf_0115cb118880283a ...string) string {
	sort.Strings(__obf_0115cb118880283a)
	__obf_51260656a7ac8feb := sha1.New()
	__obf_51260656a7ac8feb.Write([]byte(strings.Join(__obf_0115cb118880283a, "")))
	return fmt.Sprintf("%x", __obf_51260656a7ac8feb.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_0115cb118880283a ...string) string {
	sort.Strings(__obf_0115cb118880283a)
	__obf_51260656a7ac8feb := md5.New()
	__obf_51260656a7ac8feb.Write([]byte(strings.Join(__obf_0115cb118880283a, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_51260656a7ac8feb.Sum(nil)))
}
