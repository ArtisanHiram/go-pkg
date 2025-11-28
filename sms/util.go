package __obf_2457be4436ec6bd6

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

var __obf_d49c8b67ff3ccef0 int64 = 0
var __obf_3d197b81a732bdc3 int64 = time.Now().UnixNano() / 1e6

func __obf_7f5e91beda1d5df9() string {

	__obf_a0e2128bca223336 := make([]byte, 64)
	__obf_a0e2128bca223336 = __obf_a0e2128bca223336[:runtime.Stack(__obf_a0e2128bca223336, false)]
	__obf_a0e2128bca223336 = bytes.TrimPrefix(__obf_a0e2128bca223336, []byte("goroutine "))
	__obf_a0e2128bca223336 = __obf_a0e2128bca223336[:bytes.IndexByte(__obf_a0e2128bca223336, ' ')]
	__obf_fab6d04345945617, _ := strconv.ParseUint(string(__obf_a0e2128bca223336), 10, 64)

	__obf_2dfbffd1bbe6dfd8 := time.Now().UnixNano() / 1e6
	__obf_19cf0140e843155f := atomic.AddInt64(&__obf_d49c8b67ff3ccef0, 1)
	__obf_79d01bce6c9d3607 := mrand.Int63()
	__obf_7525fbc4f143315d := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_3d197b81a732bdc3, __obf_fab6d04345945617, __obf_2dfbffd1bbe6dfd8, __obf_19cf0140e843155f, __obf_79d01bce6c9d3607)
	__obf_d7b2b9cf230e1c67 := md5.New()
	__obf_d7b2b9cf230e1c67.Write([]byte(__obf_7525fbc4f143315d))
	return hex.EncodeToString(__obf_d7b2b9cf230e1c67.Sum(nil))
}

// generateSignature 生成签名
func __obf_08f1451672c192b9(__obf_4b8f7bed5d437541 string, __obf_84b05fd3511d0033 string) string {
	var __obf_bacae7d48ba9dd66 bytes.Buffer
	__obf_bacae7d48ba9dd66.WriteString("GET&")
	__obf_bacae7d48ba9dd66.WriteString(__obf_ccb162ee6cf7f68b("/"))
	__obf_bacae7d48ba9dd66.WriteString("&")
	__obf_bacae7d48ba9dd66.WriteString(__obf_ccb162ee6cf7f68b(__obf_4b8f7bed5d437541))

	__obf_d7b2b9cf230e1c67 := hmac.New(sha1.New, []byte(__obf_84b05fd3511d0033+"&"))
	__obf_d7b2b9cf230e1c67.Write(__obf_bacae7d48ba9dd66.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_d7b2b9cf230e1c67.Sum(nil))
}
