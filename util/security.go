package __obf_083c8deafa73f533

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

func Sha256Hex(__obf_b1e85a310b2a8288 []byte) string {
	__obf_f2d2cb562d270bc5 := sha256.Sum256(__obf_b1e85a310b2a8288)
	return hex.EncodeToString(__obf_f2d2cb562d270bc5[:])
}

func HashPassword(__obf_37e867abee569719 string) (string, error) {
	bytes, __obf_ab078048114898aa := bcrypt.GenerateFromPassword([]byte(__obf_37e867abee569719), bcrypt.DefaultCost)
	return string(bytes), __obf_ab078048114898aa
}

func CheckPasswordHash(__obf_37e867abee569719, __obf_8c461311413369ed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_8c461311413369ed), []byte(__obf_37e867abee569719)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_116809f90a65bc1a string) string {
	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	__obf_f2d2cb562d270bc5 := sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_f2d2cb562d270bc5.Write([]byte(__obf_116809f90a65bc1a))

	// This gets the finalized hash result as a byte
	// slice. The argument to `Sum` can be used to append
	// to an existing byte slice: it usually isn't needed.
	__obf_ca09173ee115154f := __obf_f2d2cb562d270bc5.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_ca09173ee115154f)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_641c50999baea69d, __obf_5e4139c6b01db295 []byte) (__obf_7f1a2faf30532967 []byte, __obf_ab078048114898aa error) {
	__obf_3d6af9358365c3d2, __obf_ab078048114898aa := aes.NewCipher(__obf_5e4139c6b01db295)
	if __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}
	__obf_f6befa370f20df34 := make([]byte, aes.BlockSize)
	if _, __obf_ab078048114898aa := io.ReadFull(rand.Reader, __obf_f6befa370f20df34); __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}
	__obf_7f1a2faf30532967 = make([]byte, len(__obf_641c50999baea69d))
	cipher.NewCBCDecrypter(__obf_3d6af9358365c3d2, __obf_f6befa370f20df34).CryptBlocks(__obf_7f1a2faf30532967, __obf_641c50999baea69d)

	return PKCS7UnPad(__obf_7f1a2faf30532967), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_556ff4928f9b300e []byte) []byte {
	__obf_48154335eab5cfd3 := len(__obf_556ff4928f9b300e)
	__obf_9f675d7533014c4d := int(__obf_556ff4928f9b300e[__obf_48154335eab5cfd3-1])
	return __obf_556ff4928f9b300e[:__obf_48154335eab5cfd3-__obf_9f675d7533014c4d]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_641c50999baea69d []byte, __obf_5e4139c6b01db295 []byte) ([]byte, error) {
	__obf_4f6630328e99b83a := len(__obf_5e4139c6b01db295)
	if len(__obf_641c50999baea69d)%__obf_4f6630328e99b83a != 0 {
		__obf_641c50999baea69d = PKCS7Pad(__obf_641c50999baea69d, __obf_4f6630328e99b83a)
	}

	__obf_3d6af9358365c3d2, __obf_ab078048114898aa := aes.NewCipher(__obf_5e4139c6b01db295)
	if __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}

	__obf_f6befa370f20df34 := make([]byte, aes.BlockSize)
	if _, __obf_ab078048114898aa := io.ReadFull(rand.Reader, __obf_f6befa370f20df34); __obf_ab078048114898aa != nil {
		return nil, __obf_ab078048114898aa
	}

	__obf_7f1a2faf30532967 := make([]byte, len(__obf_641c50999baea69d))
	cipher.NewCBCEncrypter(__obf_3d6af9358365c3d2, __obf_f6befa370f20df34).CryptBlocks(__obf_7f1a2faf30532967, __obf_641c50999baea69d)

	return __obf_7f1a2faf30532967, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_556ff4928f9b300e []byte, __obf_5c53da3ff6e1f197 int) []byte {
	if __obf_5c53da3ff6e1f197 < 1<<1 || __obf_5c53da3ff6e1f197 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_9f675d7533014c4d := __obf_5c53da3ff6e1f197 - len(__obf_556ff4928f9b300e)%__obf_5c53da3ff6e1f197
	__obf_f40f0643f28ca48e := bytes.Repeat([]byte{byte(__obf_9f675d7533014c4d)}, __obf_9f675d7533014c4d)
	return append(__obf_556ff4928f9b300e, __obf_f40f0643f28ca48e...)
}

// SortSha1 排序并sha1，主要用于计算signature
func SortSha1(__obf_116809f90a65bc1a ...string) string {
	sort.Strings(__obf_116809f90a65bc1a)
	__obf_f2d2cb562d270bc5 := sha1.New()
	__obf_f2d2cb562d270bc5.Write([]byte(strings.Join(__obf_116809f90a65bc1a, "")))
	return fmt.Sprintf("%x", __obf_f2d2cb562d270bc5.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_116809f90a65bc1a ...string) string {
	sort.Strings(__obf_116809f90a65bc1a)
	__obf_f2d2cb562d270bc5 := md5.New()
	__obf_f2d2cb562d270bc5.Write([]byte(strings.Join(__obf_116809f90a65bc1a, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_f2d2cb562d270bc5.Sum(nil)))
}
