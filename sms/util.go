package __obf_7147dbcba87e70eb

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	mrand "math/rand"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
)

var __obf_fd855db66f81f7ca int64 = 0
var __obf_40d168af725e40ae int64 = time.Now().UnixNano() / 1e6

func __obf_c1c5408641580711() string {
	__obf_14fc605f8bf651d8 := make([]byte, 64)
	__obf_14fc605f8bf651d8 = __obf_14fc605f8bf651d8[:runtime.Stack(__obf_14fc605f8bf651d8, false)]
	__obf_14fc605f8bf651d8 = bytes.TrimPrefix(__obf_14fc605f8bf651d8, []byte("goroutine "))
	__obf_14fc605f8bf651d8 = __obf_14fc605f8bf651d8[:bytes.IndexByte(__obf_14fc605f8bf651d8, ' ')]
	__obf_b4d67b0c062a97d4, _ := strconv.ParseUint(string(__obf_14fc605f8bf651d8), 10, 64)
	__obf_16492a0bcb1ab9d2 := time.Now().UnixNano() / 1e6
	__obf_c8f294e8b5f412f6 := atomic.AddInt64(&__obf_fd855db66f81f7ca, 1)
	__obf_28e1c88c8653f5e5 := mrand.Int63()
	__obf_4e104c7ba0640d1c := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_40d168af725e40ae, __obf_b4d67b0c062a97d4, __obf_16492a0bcb1ab9d2, __obf_c8f294e8b5f412f6, __obf_28e1c88c8653f5e5)
	__obf_d61d8d66afea80ff := md5.New()
	__obf_d61d8d66afea80ff.
		Write([]byte(__obf_4e104c7ba0640d1c))
	return hex.EncodeToString(__obf_d61d8d66afea80ff.Sum(nil))
}

// generateSignature 生成签名
func __obf_63582ab28b26163a(__obf_70535f4bc6e4828c string, __obf_2d230803be94b4a9 string) string {
	var __obf_5a98a0f9ed47d503 bytes.Buffer
	__obf_5a98a0f9ed47d503.
		WriteString("GET&")
	__obf_5a98a0f9ed47d503.
		WriteString(__obf_b3ef92f29440b2e9("/"))
	__obf_5a98a0f9ed47d503.
		WriteString("&")
	__obf_5a98a0f9ed47d503.
		WriteString(__obf_b3ef92f29440b2e9(__obf_70535f4bc6e4828c))
	__obf_d61d8d66afea80ff := hmac.New(sha1.New, []byte(__obf_2d230803be94b4a9+"&"))
	__obf_d61d8d66afea80ff.
		Write(__obf_5a98a0f9ed47d503.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_d61d8d66afea80ff.Sum(nil))
}
