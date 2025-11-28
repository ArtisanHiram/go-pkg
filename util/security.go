package __obf_077bcefb098a89af

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

func Sha256Hex(__obf_3811c39ee1400927 []byte) string {
	__obf_9c55061c9e6cb555 := sha256.Sum256(__obf_3811c39ee1400927)
	return hex.EncodeToString(__obf_9c55061c9e6cb555[:])
}

func HashPassword(__obf_740e61837585bac4 string) (string, error) {
	bytes, __obf_224415a75b186177 := bcrypt.GenerateFromPassword([]byte(__obf_740e61837585bac4), bcrypt.DefaultCost)
	return string(bytes), __obf_224415a75b186177
}

func CheckPasswordHash(__obf_740e61837585bac4, __obf_2b7f25642ee1177a string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_2b7f25642ee1177a), []byte(__obf_740e61837585bac4)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_2b69c849afdb21c3 string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	__obf_9c55061c9e6cb555 := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_9c55061c9e6cb555.Write([]byte(__obf_2b69c849afdb21c3))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	__obf_ba19cdd221caf0ea := __obf_9c55061c9e6cb555.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_ba19cdd221caf0ea)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_1da8943d4f351337, __obf_0af462578a0d3820 []byte) (__obf_f19de3d25da94be7 []byte, __obf_224415a75b186177 error) {
	__obf_1f567a49727eca19, __obf_224415a75b186177 := aes.NewCipher(__obf_0af462578a0d3820)
	if __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}
	__obf_de23193c818d7c55 := make([]byte, aes.BlockSize)
	if _, __obf_224415a75b186177 := io.ReadFull(rand.Reader, __obf_de23193c818d7c55); __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}
	__obf_f19de3d25da94be7 = make([]byte, len(__obf_1da8943d4f351337))
	cipher.NewCBCDecrypter(__obf_1f567a49727eca19, __obf_de23193c818d7c55).CryptBlocks(__obf_f19de3d25da94be7, __obf_1da8943d4f351337)

	return PKCS7UnPad(__obf_f19de3d25da94be7), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_547e3f206d29d5b3 []byte) []byte {
	__obf_85c14d42f927cfdf := len(__obf_547e3f206d29d5b3)
	__obf_1d85fd8d303718b0 := int(__obf_547e3f206d29d5b3[__obf_85c14d42f927cfdf-1])
	return __obf_547e3f206d29d5b3[:__obf_85c14d42f927cfdf-__obf_1d85fd8d303718b0]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_1da8943d4f351337 []byte, __obf_0af462578a0d3820 []byte) ([]byte, error) {
	__obf_1be1c6a485485b05 := len(__obf_0af462578a0d3820)
	if len(__obf_1da8943d4f351337)%__obf_1be1c6a485485b05 != 0 {
		__obf_1da8943d4f351337 = PKCS7Pad(__obf_1da8943d4f351337, __obf_1be1c6a485485b05)
	}

	__obf_1f567a49727eca19, __obf_224415a75b186177 := aes.NewCipher(__obf_0af462578a0d3820)
	if __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}

	__obf_de23193c818d7c55 := make([]byte, aes.BlockSize)
	if _, __obf_224415a75b186177 := io.ReadFull(rand.Reader, __obf_de23193c818d7c55); __obf_224415a75b186177 != nil {
		return nil, __obf_224415a75b186177
	}

	__obf_f19de3d25da94be7 := make([]byte, len(__obf_1da8943d4f351337))
	cipher.NewCBCEncrypter(__obf_1f567a49727eca19, __obf_de23193c818d7c55).CryptBlocks(__obf_f19de3d25da94be7, __obf_1da8943d4f351337)

	return __obf_f19de3d25da94be7, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_547e3f206d29d5b3 []byte, __obf_b58231598f8a57f3 int) []byte {
	if __obf_b58231598f8a57f3 < 1<<1 || __obf_b58231598f8a57f3 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_1d85fd8d303718b0 := __obf_b58231598f8a57f3 - len(__obf_547e3f206d29d5b3)%__obf_b58231598f8a57f3
	__obf_c30cfd15664952dc := bytes.Repeat([]byte{byte(__obf_1d85fd8d303718b0)}, __obf_1d85fd8d303718b0)
	return append(__obf_547e3f206d29d5b3, __obf_c30cfd15664952dc...)
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(__obf_2b69c849afdb21c3 ...string) string {
	sort.Strings(__obf_2b69c849afdb21c3)
	__obf_9c55061c9e6cb555 := sha1.New()
	__obf_9c55061c9e6cb555.Write([]byte(strings.Join(__obf_2b69c849afdb21c3, "")))
	return fmt.Sprintf("%x", __obf_9c55061c9e6cb555.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_2b69c849afdb21c3 ...string) string {
	sort.Strings(__obf_2b69c849afdb21c3)
	__obf_9c55061c9e6cb555 := md5.New()
	__obf_9c55061c9e6cb555.Write([]byte(strings.Join(__obf_2b69c849afdb21c3, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_9c55061c9e6cb555.Sum(nil)))
}
