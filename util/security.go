package __obf_b81118ac905f398e

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

func Sha256Hex(__obf_f3483784ccf89cd0 []byte) string {
	__obf_d575b5eea8c78a0b := sha256.Sum256(__obf_f3483784ccf89cd0)
	return hex.EncodeToString(__obf_d575b5eea8c78a0b[:])
}

func HashPassword(__obf_d0a795e3d87d9e33 string) (string, error) {
	bytes, __obf_65d3bb18bf1bd15f := bcrypt.GenerateFromPassword([]byte(__obf_d0a795e3d87d9e33), bcrypt.DefaultCost)
	return string(bytes), __obf_65d3bb18bf1bd15f
}

func CheckPasswordHash(__obf_d0a795e3d87d9e33, __obf_72f871f23cea6dc6 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_72f871f23cea6dc6), []byte(__obf_d0a795e3d87d9e33)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_c0c7a6993b194cda string) string {
	__obf_d575b5eea8c78a0b := // The pattern for generating a hash is `sha1.New()`,
		// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
		// Here we start with a new hash.
		sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_d575b5eea8c78a0b.Write([]byte(__obf_c0c7a6993b194cda))
	__obf_c49f101869cc203a := // This gets the finalized hash result as a byte
		// slice. The argument to `Sum` can be used to append
		// to an existing byte slice: it usually isn't needed.
		__obf_d575b5eea8c78a0b.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_c49f101869cc203a)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_16b08a5a1d6464d0, __obf_826deb59f8429f3a []byte) (__obf_bb69f2ec3ff72b96 []byte, __obf_65d3bb18bf1bd15f error) {
	__obf_793194c4c40cd790, __obf_65d3bb18bf1bd15f := aes.NewCipher(__obf_826deb59f8429f3a)
	if __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	__obf_5b37e5642dc17253 := make([]byte, aes.BlockSize)
	if _, __obf_65d3bb18bf1bd15f := io.ReadFull(rand.Reader, __obf_5b37e5642dc17253); __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	__obf_bb69f2ec3ff72b96 = make([]byte, len(__obf_16b08a5a1d6464d0))
	cipher.NewCBCDecrypter(__obf_793194c4c40cd790, __obf_5b37e5642dc17253).CryptBlocks(__obf_bb69f2ec3ff72b96, __obf_16b08a5a1d6464d0)

	return PKCS7UnPad(__obf_bb69f2ec3ff72b96), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_c8007c02b7698eea []byte) []byte {
	__obf_e4668c42490030e3 := len(__obf_c8007c02b7698eea)
	__obf_ac5df1a7f44fb9dd := int(__obf_c8007c02b7698eea[__obf_e4668c42490030e3-1])
	return __obf_c8007c02b7698eea[:__obf_e4668c42490030e3-__obf_ac5df1a7f44fb9dd]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_16b08a5a1d6464d0 []byte, __obf_826deb59f8429f3a []byte) ([]byte, error) {
	__obf_b29ccd2992d5240a := len(__obf_826deb59f8429f3a)
	if len(__obf_16b08a5a1d6464d0)%__obf_b29ccd2992d5240a != 0 {
		__obf_16b08a5a1d6464d0 = PKCS7Pad(__obf_16b08a5a1d6464d0, __obf_b29ccd2992d5240a)
	}
	__obf_793194c4c40cd790, __obf_65d3bb18bf1bd15f := aes.NewCipher(__obf_826deb59f8429f3a)
	if __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	__obf_5b37e5642dc17253 := make([]byte, aes.BlockSize)
	if _, __obf_65d3bb18bf1bd15f := io.ReadFull(rand.Reader, __obf_5b37e5642dc17253); __obf_65d3bb18bf1bd15f != nil {
		return nil, __obf_65d3bb18bf1bd15f
	}
	__obf_bb69f2ec3ff72b96 := make([]byte, len(__obf_16b08a5a1d6464d0))
	cipher.NewCBCEncrypter(__obf_793194c4c40cd790, __obf_5b37e5642dc17253).CryptBlocks(__obf_bb69f2ec3ff72b96, __obf_16b08a5a1d6464d0)

	return __obf_bb69f2ec3ff72b96, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_c8007c02b7698eea []byte, __obf_347f62f1e8e0e81d int) []byte {
	if __obf_347f62f1e8e0e81d < 1<<1 || __obf_347f62f1e8e0e81d >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_ac5df1a7f44fb9dd := __obf_347f62f1e8e0e81d - len(__obf_c8007c02b7698eea)%__obf_347f62f1e8e0e81d
	__obf_aab81c1a5ca349c4 := bytes.Repeat([]byte{byte(__obf_ac5df1a7f44fb9dd)}, __obf_ac5df1a7f44fb9dd)
	return append(__obf_c8007c02b7698eea,

		// SortSha1 排序并sha1，主要用于计算signature
		__obf_aab81c1a5ca349c4...)
}

func SortSha1(__obf_c0c7a6993b194cda ...string) string {
	sort.Strings(__obf_c0c7a6993b194cda)
	__obf_d575b5eea8c78a0b := sha1.New()
	__obf_d575b5eea8c78a0b.
		Write([]byte(strings.Join(__obf_c0c7a6993b194cda, "")))
	return fmt.Sprintf("%x", __obf_d575b5eea8c78a0b.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_c0c7a6993b194cda ...string) string {
	sort.Strings(__obf_c0c7a6993b194cda)
	__obf_d575b5eea8c78a0b := md5.New()
	__obf_d575b5eea8c78a0b.
		Write([]byte(strings.Join(__obf_c0c7a6993b194cda, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_d575b5eea8c78a0b.Sum(nil)))
}
