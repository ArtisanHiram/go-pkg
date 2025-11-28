package __obf_85fcc7bd65d30170

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

func Sha256Hex(__obf_7fcfc8a8085d93d2 []byte) string {
	__obf_4682ee4eeb8a789b := sha256.Sum256(__obf_7fcfc8a8085d93d2)
	return hex.EncodeToString(__obf_4682ee4eeb8a789b[:])
}

func HashPassword(__obf_9b91bd68963d3da8 string) (string, error) {
	bytes, __obf_80a78f1bfaca43d3 := bcrypt.GenerateFromPassword([]byte(__obf_9b91bd68963d3da8), bcrypt.DefaultCost)
	return string(bytes), __obf_80a78f1bfaca43d3
}

func CheckPasswordHash(__obf_9b91bd68963d3da8, __obf_92ceda2782a53d63 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_92ceda2782a53d63), []byte(__obf_9b91bd68963d3da8)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_d90c40a5354bf79c string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	__obf_4682ee4eeb8a789b := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_4682ee4eeb8a789b.Write([]byte(__obf_d90c40a5354bf79c))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	__obf_da86966b011dc473 := __obf_4682ee4eeb8a789b.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_da86966b011dc473)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_58ac4ace6141f6fb, __obf_592ca3f7d124f0fd []byte) (__obf_bdbac1049fb4cb5a []byte, __obf_80a78f1bfaca43d3 error) {
	__obf_98e2a0923bd6c60e, __obf_80a78f1bfaca43d3 := aes.NewCipher(__obf_592ca3f7d124f0fd)
	if __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}
	__obf_37b6baacb37739b8 := make([]byte, aes.BlockSize)
	if _, __obf_80a78f1bfaca43d3 := io.ReadFull(rand.Reader, __obf_37b6baacb37739b8); __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}
	__obf_bdbac1049fb4cb5a = make([]byte, len(__obf_58ac4ace6141f6fb))
	cipher.NewCBCDecrypter(__obf_98e2a0923bd6c60e, __obf_37b6baacb37739b8).CryptBlocks(__obf_bdbac1049fb4cb5a, __obf_58ac4ace6141f6fb)

	return PKCS7UnPad(__obf_bdbac1049fb4cb5a), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_7fdc431a0b68ea2f []byte) []byte {
	__obf_11bffd4ed6d8c009 := len(__obf_7fdc431a0b68ea2f)
	__obf_8968c1bc526160c0 := int(__obf_7fdc431a0b68ea2f[__obf_11bffd4ed6d8c009-1])
	return __obf_7fdc431a0b68ea2f[:__obf_11bffd4ed6d8c009-__obf_8968c1bc526160c0]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_58ac4ace6141f6fb []byte, __obf_592ca3f7d124f0fd []byte) ([]byte, error) {
	__obf_3b9db108b96306b1 := len(__obf_592ca3f7d124f0fd)
	if len(__obf_58ac4ace6141f6fb)%__obf_3b9db108b96306b1 != 0 {
		__obf_58ac4ace6141f6fb = PKCS7Pad(__obf_58ac4ace6141f6fb, __obf_3b9db108b96306b1)
	}

	__obf_98e2a0923bd6c60e, __obf_80a78f1bfaca43d3 := aes.NewCipher(__obf_592ca3f7d124f0fd)
	if __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}

	__obf_37b6baacb37739b8 := make([]byte, aes.BlockSize)
	if _, __obf_80a78f1bfaca43d3 := io.ReadFull(rand.Reader, __obf_37b6baacb37739b8); __obf_80a78f1bfaca43d3 != nil {
		return nil, __obf_80a78f1bfaca43d3
	}

	__obf_bdbac1049fb4cb5a := make([]byte, len(__obf_58ac4ace6141f6fb))
	cipher.NewCBCEncrypter(__obf_98e2a0923bd6c60e, __obf_37b6baacb37739b8).CryptBlocks(__obf_bdbac1049fb4cb5a, __obf_58ac4ace6141f6fb)

	return __obf_bdbac1049fb4cb5a, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_7fdc431a0b68ea2f []byte, __obf_799103b07957c964 int) []byte {
	if __obf_799103b07957c964 < 1<<1 || __obf_799103b07957c964 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_8968c1bc526160c0 := __obf_799103b07957c964 - len(__obf_7fdc431a0b68ea2f)%__obf_799103b07957c964
	__obf_a3bc223923d799a7 := bytes.Repeat([]byte{byte(__obf_8968c1bc526160c0)}, __obf_8968c1bc526160c0)
	return append(__obf_7fdc431a0b68ea2f, __obf_a3bc223923d799a7...)
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(__obf_d90c40a5354bf79c ...string) string {
	sort.Strings(__obf_d90c40a5354bf79c)
	__obf_4682ee4eeb8a789b := sha1.New()
	__obf_4682ee4eeb8a789b.Write([]byte(strings.Join(__obf_d90c40a5354bf79c, "")))
	return fmt.Sprintf("%x", __obf_4682ee4eeb8a789b.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_d90c40a5354bf79c ...string) string {
	sort.Strings(__obf_d90c40a5354bf79c)
	__obf_4682ee4eeb8a789b := md5.New()
	__obf_4682ee4eeb8a789b.Write([]byte(strings.Join(__obf_d90c40a5354bf79c, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_4682ee4eeb8a789b.Sum(nil)))
}
