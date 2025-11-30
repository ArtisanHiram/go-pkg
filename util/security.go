package __obf_e13b701bec415b48

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

func Sha256Hex(__obf_2f1ab84c85030098 []byte) string {
	__obf_779a9a3e583b3f28 := sha256.Sum256(__obf_2f1ab84c85030098)
	return hex.EncodeToString(__obf_779a9a3e583b3f28[:])
}

func HashPassword(__obf_74cb32db90c38c02 string) (string, error) {
	bytes, __obf_b93c04d14e5c503f := bcrypt.GenerateFromPassword([]byte(__obf_74cb32db90c38c02), bcrypt.DefaultCost)
	return string(bytes), __obf_b93c04d14e5c503f
}

func CheckPasswordHash(__obf_74cb32db90c38c02, __obf_d68eff92d5e14300 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_d68eff92d5e14300), []byte(__obf_74cb32db90c38c02)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_a9c9200b03759058 string) string {
	__obf_779a9a3e583b3f28 := // The pattern for generating a hash is `sha1.New()`,
		// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
		// Here we start with a new hash.
		sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_779a9a3e583b3f28.Write([]byte(__obf_a9c9200b03759058))
	__obf_787651573548af89 := // This gets the finalized hash result as a byte
		// slice. The argument to `Sum` can be used to append
		// to an existing byte slice: it usually isn't needed.
		__obf_779a9a3e583b3f28.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_787651573548af89)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_f7fcca366fccd9d2, __obf_88ef59ecffef7a5b []byte) (__obf_6004f7d37045bf33 []byte, __obf_b93c04d14e5c503f error) {
	__obf_19919e3a8d3689fa, __obf_b93c04d14e5c503f := aes.NewCipher(__obf_88ef59ecffef7a5b)
	if __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	__obf_c27f64ca31406548 := make([]byte, aes.BlockSize)
	if _, __obf_b93c04d14e5c503f := io.ReadFull(rand.Reader, __obf_c27f64ca31406548); __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	__obf_6004f7d37045bf33 = make([]byte, len(__obf_f7fcca366fccd9d2))
	cipher.NewCBCDecrypter(__obf_19919e3a8d3689fa, __obf_c27f64ca31406548).CryptBlocks(__obf_6004f7d37045bf33, __obf_f7fcca366fccd9d2)

	return PKCS7UnPad(__obf_6004f7d37045bf33), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_ecc61abd9bcbb1bb []byte) []byte {
	__obf_fdba24dfa958082c := len(__obf_ecc61abd9bcbb1bb)
	__obf_b72c47846fc8a12e := int(__obf_ecc61abd9bcbb1bb[__obf_fdba24dfa958082c-1])
	return __obf_ecc61abd9bcbb1bb[:__obf_fdba24dfa958082c-__obf_b72c47846fc8a12e]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_f7fcca366fccd9d2 []byte, __obf_88ef59ecffef7a5b []byte) ([]byte, error) {
	__obf_ccfb9e5b04e5d653 := len(__obf_88ef59ecffef7a5b)
	if len(__obf_f7fcca366fccd9d2)%__obf_ccfb9e5b04e5d653 != 0 {
		__obf_f7fcca366fccd9d2 = PKCS7Pad(__obf_f7fcca366fccd9d2, __obf_ccfb9e5b04e5d653)
	}
	__obf_19919e3a8d3689fa, __obf_b93c04d14e5c503f := aes.NewCipher(__obf_88ef59ecffef7a5b)
	if __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	__obf_c27f64ca31406548 := make([]byte, aes.BlockSize)
	if _, __obf_b93c04d14e5c503f := io.ReadFull(rand.Reader, __obf_c27f64ca31406548); __obf_b93c04d14e5c503f != nil {
		return nil, __obf_b93c04d14e5c503f
	}
	__obf_6004f7d37045bf33 := make([]byte, len(__obf_f7fcca366fccd9d2))
	cipher.NewCBCEncrypter(__obf_19919e3a8d3689fa, __obf_c27f64ca31406548).CryptBlocks(__obf_6004f7d37045bf33, __obf_f7fcca366fccd9d2)

	return __obf_6004f7d37045bf33, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_ecc61abd9bcbb1bb []byte, __obf_cbe548df200cfe33 int) []byte {
	if __obf_cbe548df200cfe33 < 1<<1 || __obf_cbe548df200cfe33 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_b72c47846fc8a12e := __obf_cbe548df200cfe33 - len(__obf_ecc61abd9bcbb1bb)%__obf_cbe548df200cfe33
	__obf_e506e7b74a36a9a7 := bytes.Repeat([]byte{byte(__obf_b72c47846fc8a12e)}, __obf_b72c47846fc8a12e)
	return append(__obf_ecc61abd9bcbb1bb,

		// SortSha1 排序并sha1，主要用于计算signature
		__obf_e506e7b74a36a9a7...)
}

func SortSha1(__obf_a9c9200b03759058 ...string) string {
	sort.Strings(__obf_a9c9200b03759058)
	__obf_779a9a3e583b3f28 := sha1.New()
	__obf_779a9a3e583b3f28.
		Write([]byte(strings.Join(__obf_a9c9200b03759058, "")))
	return fmt.Sprintf("%x", __obf_779a9a3e583b3f28.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_a9c9200b03759058 ...string) string {
	sort.Strings(__obf_a9c9200b03759058)
	__obf_779a9a3e583b3f28 := md5.New()
	__obf_779a9a3e583b3f28.
		Write([]byte(strings.Join(__obf_a9c9200b03759058, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_779a9a3e583b3f28.Sum(nil)))
}
