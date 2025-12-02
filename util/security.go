package __obf_8b17832dd38bb602

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

func Sha256Hex(__obf_2304d1cf2f26db23 []byte) string {
	__obf_089acd895c45c974 := sha256.Sum256(__obf_2304d1cf2f26db23)
	return hex.EncodeToString(__obf_089acd895c45c974[:])
}

func HashPassword(__obf_32dd7083c1de82fc string) (string, error) {
	bytes, __obf_17881c470b1e6626 := bcrypt.GenerateFromPassword([]byte(__obf_32dd7083c1de82fc), bcrypt.DefaultCost)
	return string(bytes), __obf_17881c470b1e6626
}

func CheckPasswordHash(__obf_32dd7083c1de82fc, __obf_38d13ae6023eded0 string) bool {
	return bcrypt.CompareHashAndPassword([]byte(__obf_38d13ae6023eded0), []byte(__obf_32dd7083c1de82fc)) == nil
}

// func Md5(str string) string {
// 	h := md5.New()
// 	h.Write([]byte(str))
// 	return hex.EncodeToString(h.Sum(nil))
// }

func Sha1Sign(__obf_d3f34fed89aa60a8 string) string {
	__obf_089acd895c45c974 := // The pattern for generating a hash is `sha1.New()`,
		// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
		// Here we start with a new hash.
		sha1.New()

	// `Write` expects bytes. If you have a string `s`,
	// use `[]byte(s)` to coerce it to bytes.
	_, _ = __obf_089acd895c45c974.Write([]byte(__obf_d3f34fed89aa60a8))
	__obf_e9e12ba560a4da6a := // This gets the finalized hash result as a byte
		// slice. The argument to `Sum` can be used to append
		// to an existing byte slice: it usually isn't needed.
		__obf_089acd895c45c974.Sum(nil)

	// SHA1 values are often printed in hex, for example
	// in git commits. Use the `%x` format verb to convert
	// a hash results to a hex string.
	return fmt.Sprintf("%x", __obf_e9e12ba560a4da6a)
}

// AesDecrypt AES-CBC解密,PKCS#7,传入密文和密钥，[]byte
func AesDecrypt(__obf_6e113e0dc133cb4d, __obf_adf3e0476606c3da []byte) (__obf_e906b0ee3883f2b3 []byte, __obf_17881c470b1e6626 error) {
	__obf_8c63d39a956dbfcd, __obf_17881c470b1e6626 := aes.NewCipher(__obf_adf3e0476606c3da)
	if __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	__obf_c27a81fa27ac089b := make([]byte, aes.BlockSize)
	if _, __obf_17881c470b1e6626 := io.ReadFull(rand.Reader, __obf_c27a81fa27ac089b); __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	__obf_e906b0ee3883f2b3 = make([]byte, len(__obf_6e113e0dc133cb4d))
	cipher.NewCBCDecrypter(__obf_8c63d39a956dbfcd, __obf_c27a81fa27ac089b).CryptBlocks(__obf_e906b0ee3883f2b3, __obf_6e113e0dc133cb4d)

	return PKCS7UnPad(__obf_e906b0ee3883f2b3), nil
}

// PKCS7UnPad PKSC#7解包
func PKCS7UnPad(__obf_30b0dd0b61edb3f7 []byte) []byte {
	__obf_ad4bb67c6362e4c5 := len(__obf_30b0dd0b61edb3f7)
	__obf_52d9aef336259fa0 := int(__obf_30b0dd0b61edb3f7[__obf_ad4bb67c6362e4c5-1])
	return __obf_30b0dd0b61edb3f7[:__obf_ad4bb67c6362e4c5-__obf_52d9aef336259fa0]
}

// AesEncrypt AES-CBC加密+PKCS#7打包，传入明文和密钥
func AesEncrypt(__obf_6e113e0dc133cb4d []byte, __obf_adf3e0476606c3da []byte) ([]byte, error) {
	__obf_e97fdc8b2a44cc99 := len(__obf_adf3e0476606c3da)
	if len(__obf_6e113e0dc133cb4d)%__obf_e97fdc8b2a44cc99 != 0 {
		__obf_6e113e0dc133cb4d = PKCS7Pad(__obf_6e113e0dc133cb4d, __obf_e97fdc8b2a44cc99)
	}
	__obf_8c63d39a956dbfcd, __obf_17881c470b1e6626 := aes.NewCipher(__obf_adf3e0476606c3da)
	if __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	__obf_c27a81fa27ac089b := make([]byte, aes.BlockSize)
	if _, __obf_17881c470b1e6626 := io.ReadFull(rand.Reader, __obf_c27a81fa27ac089b); __obf_17881c470b1e6626 != nil {
		return nil, __obf_17881c470b1e6626
	}
	__obf_e906b0ee3883f2b3 := make([]byte, len(__obf_6e113e0dc133cb4d))
	cipher.NewCBCEncrypter(__obf_8c63d39a956dbfcd, __obf_c27a81fa27ac089b).CryptBlocks(__obf_e906b0ee3883f2b3, __obf_6e113e0dc133cb4d)

	return __obf_e906b0ee3883f2b3, nil
}

// PKCS7Pad PKCS#7打包
func PKCS7Pad(__obf_30b0dd0b61edb3f7 []byte, __obf_23ec941a3d254f12 int) []byte {
	if __obf_23ec941a3d254f12 < 1<<1 || __obf_23ec941a3d254f12 >= 1<<8 {
		panic("unsupported block size")
	}
	__obf_52d9aef336259fa0 := __obf_23ec941a3d254f12 - len(__obf_30b0dd0b61edb3f7)%__obf_23ec941a3d254f12
	__obf_06b61b23e7186213 := bytes.Repeat([]byte{byte(__obf_52d9aef336259fa0)}, __obf_52d9aef336259fa0)
	return append(__obf_30b0dd0b61edb3f7,

		// SortSha1 排序并sha1，主要用于计算signature
		__obf_06b61b23e7186213...)
}

func SortSha1(__obf_d3f34fed89aa60a8 ...string) string {
	sort.Strings(__obf_d3f34fed89aa60a8)
	__obf_089acd895c45c974 := sha1.New()
	__obf_089acd895c45c974.
		Write([]byte(strings.Join(__obf_d3f34fed89aa60a8, "")))
	return fmt.Sprintf("%x", __obf_089acd895c45c974.Sum(nil))
}

// SortMd5 排序并md5，主要用于计算sign
func SortMd5(__obf_d3f34fed89aa60a8 ...string) string {
	sort.Strings(__obf_d3f34fed89aa60a8)
	__obf_089acd895c45c974 := md5.New()
	__obf_089acd895c45c974.
		Write([]byte(strings.Join(__obf_d3f34fed89aa60a8, "")))
	return strings.ToUpper(fmt.Sprintf("%x", __obf_089acd895c45c974.Sum(nil)))
}
