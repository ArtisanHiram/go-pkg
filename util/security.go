package __obf_69af70389b6622a3

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

func Sha256Hex(__obf_55286557d75f8aca []byte) string {
	__obf_d885756428514f2e := sha256.Sum256(__obf_55286557d75f8aca)
	return hex.EncodeToString(__obf_d885756428514f2e[:])
}

func HashPassword(__obf_4c91d5ce8562fb6d string) (string, error) {
	bytes, __obf_e812cfd1203ed4f3 := bcrypt.GenerateFromPassword([]byte(__obf_4c91d5ce8562fb6d), bcrypt.DefaultCost)
	return string(bytes), __obf_e812cfd1203ed4f3
}

func CheckPasswordHash(__obf_4c91d5ce8562fb6d, __obf_2bdcf41eeb4dfb94 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_2bdcf41eeb4dfb94), []byte(__obf_4c91d5ce8562fb6d)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_5331eba287d21166 string) string {
	__obf_d885756428514f2e := // The pattern for generating a hash is `sha1.New()`,
		// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
		// Here we start with a new hash.
		sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_d885756428514f2e.Write([]byte(__obf_5331eba287d21166))
	__obf_6f8b25ebb4a0487e := // This gets the finalized hash result as a byte
		// slice. The argument to `Sum` can be used to append
		// to an existing byte slice: it usually isn't needed.
		__obf_d885756428514f2e.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_6f8b25ebb4a0487e)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_6fb8112d553c288e, __obf_10d6974a5cfc5e1f []byte) (__obf_566697b3f9c80e4b []byte, __obf_e812cfd1203ed4f3 error) {
	__obf_591c6c6b789101f3, __obf_e812cfd1203ed4f3 := aes.NewCipher(__obf_10d6974a5cfc5e1f)
	if __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	__obf_d123856c0ce33d9d := make([]byte, aes.BlockSize)
	if _, __obf_e812cfd1203ed4f3 := io.ReadFull(rand.Reader, __obf_d123856c0ce33d9d); __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	__obf_566697b3f9c80e4b = make([]byte, len(__obf_6fb8112d553c288e))
	cipher.NewCBCDecrypter(__obf_591c6c6b789101f3, __obf_d123856c0ce33d9d).CryptBlocks(__obf_566697b3f9c80e4b, __obf_6fb8112d553c288e)

	return PKCS7UnPad(__obf_566697b3f9c80e4b), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_406cdae9ca413415 []byte) []byte {
	__obf_9020047b0bb52202 := len(__obf_406cdae9ca413415)
	__obf_f98e27b75761eded := int(__obf_406cdae9ca413415[__obf_9020047b0bb52202-1])
	return __obf_406cdae9ca413415[:__obf_9020047b0bb52202-__obf_f98e27b75761eded]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_6fb8112d553c288e []byte, __obf_10d6974a5cfc5e1f []byte) ([]byte, error) {
	__obf_5028ac00fc46e50f := len(__obf_10d6974a5cfc5e1f)
	if len(__obf_6fb8112d553c288e)%__obf_5028ac00fc46e50f != 0 {
		__obf_6fb8112d553c288e = PKCS7Pad(__obf_6fb8112d553c288e, __obf_5028ac00fc46e50f)
	}
	__obf_591c6c6b789101f3, __obf_e812cfd1203ed4f3 := aes.NewCipher(__obf_10d6974a5cfc5e1f)
	if __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	__obf_d123856c0ce33d9d := make([]byte, aes.BlockSize)
	if _, __obf_e812cfd1203ed4f3 := io.ReadFull(rand.Reader, __obf_d123856c0ce33d9d); __obf_e812cfd1203ed4f3 != nil {
		return nil, __obf_e812cfd1203ed4f3
	}
	__obf_566697b3f9c80e4b := make([]byte, len(__obf_6fb8112d553c288e))
	cipher.NewCBCEncrypter(__obf_591c6c6b789101f3, __obf_d123856c0ce33d9d).CryptBlocks(__obf_566697b3f9c80e4b, __obf_6fb8112d553c288e)

	return __obf_566697b3f9c80e4b, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_406cdae9ca413415 []byte, __obf_7c025c60ec8d7f56 int) []byte {
	if __obf_7c025c60ec8d7f56 < 1<<1 || __obf_7c025c60ec8d7f56 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_f98e27b75761eded := __obf_7c025c60ec8d7f56 - len(__obf_406cdae9ca413415)%__obf_7c025c60ec8d7f56
	__obf_d273e0583c2d1544 := bytes.Repeat([]byte{byte(__obf_f98e27b75761eded)}, __obf_f98e27b75761eded)
	return append(__obf_406cdae9ca413415,

		// SortSha1 排序并sha1，主要用于计算signature
		__obf_d273e0583c2d1544...)
}

func SortSha1(__obf_5331eba287d21166 ...string) string {
	sort.Strings(__obf_5331eba287d21166)
	__obf_d885756428514f2e := sha1.New()
	__obf_d885756428514f2e.
		Write([]byte(strings.Join(__obf_5331eba287d21166, "")))
	return fmt.Sprintf("%x", __obf_d885756428514f2e.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_5331eba287d21166 ...string) string {
	sort.Strings(__obf_5331eba287d21166)
	__obf_d885756428514f2e := md5.New()
	__obf_d885756428514f2e.
		Write([]byte(strings.Join(__obf_5331eba287d21166, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_d885756428514f2e.Sum(nil)))
}
