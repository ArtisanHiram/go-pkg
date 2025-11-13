package __obf_813b65e0329aecbf

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

func Sha256Hex(__obf_243fff1d60c134b7 []byte) string {
	__obf_9033a830731ade2f := sha256.Sum256(__obf_243fff1d60c134b7)
	return hex.EncodeToString(__obf_9033a830731ade2f[:])
}

func HashPassword(__obf_808f912c748b75c1 string) (string, error) {
	bytes, __obf_54101b1325683820 := bcrypt.GenerateFromPassword([]byte(__obf_808f912c748b75c1), bcrypt.DefaultCost)
	return string(bytes), __obf_54101b1325683820
}

func CheckPasswordHash(__obf_808f912c748b75c1, __obf_7d187af32c0a96e1 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_7d187af32c0a96e1), []byte(__obf_808f912c748b75c1)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_3b0cfd8394bed5fb string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	__obf_9033a830731ade2f := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_9033a830731ade2f.Write([]byte(__obf_3b0cfd8394bed5fb))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	__obf_be9ee4c737c5afaf := __obf_9033a830731ade2f.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_be9ee4c737c5afaf)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_36f511fc7589036d, __obf_3610eaa580a77615 []byte) (__obf_071db135d0378af1 []byte, __obf_54101b1325683820 error) {
	__obf_ae39597f2ffeafce, __obf_54101b1325683820 := aes.NewCipher(__obf_3610eaa580a77615)
	if __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}
	__obf_944655a1bc13e6e9 := make([]byte, aes.BlockSize)
	if _, __obf_54101b1325683820 := io.ReadFull(rand.Reader, __obf_944655a1bc13e6e9); __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}
	__obf_071db135d0378af1 = make([]byte, len(__obf_36f511fc7589036d))
	cipher.NewCBCDecrypter(__obf_ae39597f2ffeafce, __obf_944655a1bc13e6e9).CryptBlocks(__obf_071db135d0378af1, __obf_36f511fc7589036d)

	return PKCS7UnPad(__obf_071db135d0378af1), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_48afc57ded5346ca []byte) []byte {
	__obf_d38fc4dc3ee3e9cd := len(__obf_48afc57ded5346ca)
	__obf_fe8147a1e3bf5fe5 := int(__obf_48afc57ded5346ca[__obf_d38fc4dc3ee3e9cd-1])
	return __obf_48afc57ded5346ca[:__obf_d38fc4dc3ee3e9cd-__obf_fe8147a1e3bf5fe5]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_36f511fc7589036d []byte, __obf_3610eaa580a77615 []byte) ([]byte, error) {
	__obf_12f59ca81d2169f7 := len(__obf_3610eaa580a77615)
	if len(__obf_36f511fc7589036d)%__obf_12f59ca81d2169f7 != 0 {
		__obf_36f511fc7589036d = PKCS7Pad(__obf_36f511fc7589036d, __obf_12f59ca81d2169f7)
	}

	__obf_ae39597f2ffeafce, __obf_54101b1325683820 := aes.NewCipher(__obf_3610eaa580a77615)
	if __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}

	__obf_944655a1bc13e6e9 := make([]byte, aes.BlockSize)
	if _, __obf_54101b1325683820 := io.ReadFull(rand.Reader, __obf_944655a1bc13e6e9); __obf_54101b1325683820 != nil {
		return nil, __obf_54101b1325683820
	}

	__obf_071db135d0378af1 := make([]byte, len(__obf_36f511fc7589036d))
	cipher.NewCBCEncrypter(__obf_ae39597f2ffeafce, __obf_944655a1bc13e6e9).CryptBlocks(__obf_071db135d0378af1, __obf_36f511fc7589036d)

	return __obf_071db135d0378af1, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_48afc57ded5346ca []byte, __obf_ec20d74e5460860b int) []byte {
	if __obf_ec20d74e5460860b < 1<<1 || __obf_ec20d74e5460860b >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_fe8147a1e3bf5fe5 := __obf_ec20d74e5460860b - len(__obf_48afc57ded5346ca)%__obf_ec20d74e5460860b
	__obf_51798d30e01e9984 := bytes.Repeat([]byte{byte(__obf_fe8147a1e3bf5fe5)}, __obf_fe8147a1e3bf5fe5)
	return append(__obf_48afc57ded5346ca, __obf_51798d30e01e9984...)
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(__obf_3b0cfd8394bed5fb ...string) string {
	sort.Strings(__obf_3b0cfd8394bed5fb)
	__obf_9033a830731ade2f := sha1.New()
	__obf_9033a830731ade2f.Write([]byte(strings.Join(__obf_3b0cfd8394bed5fb, "")))
	return fmt.Sprintf("%x", __obf_9033a830731ade2f.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_3b0cfd8394bed5fb ...string) string {
	sort.Strings(__obf_3b0cfd8394bed5fb)
	__obf_9033a830731ade2f := md5.New()
	__obf_9033a830731ade2f.Write([]byte(strings.Join(__obf_3b0cfd8394bed5fb, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_9033a830731ade2f.Sum(nil)))
}
