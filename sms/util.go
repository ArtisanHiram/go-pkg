package __obf_dd37d42fbda9c938

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

var __obf_0dc0f7f7815ff8f3 int64 = 0
var __obf_5c16393dbf84701d int64 = time.Now().UnixNano() / 1e6

func __obf_4d1620aae5ae13dc() string {
	__obf_a1dfd1f4475c7fd7 := make([]byte, 64)
	__obf_a1dfd1f4475c7fd7 = __obf_a1dfd1f4475c7fd7[:runtime.Stack(__obf_a1dfd1f4475c7fd7, false)]
	__obf_a1dfd1f4475c7fd7 = bytes.TrimPrefix(__obf_a1dfd1f4475c7fd7, []byte("goroutine "))
	__obf_a1dfd1f4475c7fd7 = __obf_a1dfd1f4475c7fd7[:bytes.IndexByte(__obf_a1dfd1f4475c7fd7, ' ')]
	__obf_bb4268257c71c625, _ := strconv.ParseUint(string(__obf_a1dfd1f4475c7fd7), 10, 64)
	__obf_2f8cf4b082a1a8f7 := time.Now().UnixNano() / 1e6
	__obf_5e58d88853dbe0da := atomic.AddInt64(&__obf_0dc0f7f7815ff8f3, 1)
	__obf_5c1d9f76a1d1ce48 := mrand.Int63()
	__obf_3488a3aae5ba4a1e := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_5c16393dbf84701d, __obf_bb4268257c71c625, __obf_2f8cf4b082a1a8f7, __obf_5e58d88853dbe0da, __obf_5c1d9f76a1d1ce48)
	__obf_254919d2f385878b := md5.New()
	__obf_254919d2f385878b.
		Write([]byte(__obf_3488a3aae5ba4a1e))
	return hex.EncodeToString(__obf_254919d2f385878b.Sum(nil))
}

// generateSignature 生成签名
func __obf_3072ced250ac589d(__obf_3d4a1793e73d0a22 string, __obf_9796f2cbd157ab60 string) string {
	var __obf_f9eadb7782e8b573 bytes.Buffer
	__obf_f9eadb7782e8b573.
		WriteString("GET&")
	__obf_f9eadb7782e8b573.
		WriteString(__obf_be8d53df8820d385("/"))
	__obf_f9eadb7782e8b573.
		WriteString("&")
	__obf_f9eadb7782e8b573.
		WriteString(__obf_be8d53df8820d385(__obf_3d4a1793e73d0a22))
	__obf_254919d2f385878b := hmac.New(sha1.New, []byte(__obf_9796f2cbd157ab60+"&"))
	__obf_254919d2f385878b.
		Write(__obf_f9eadb7782e8b573.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_254919d2f385878b.Sum(nil))
}
