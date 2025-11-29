package __obf_18a392c58f8f4352

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

var __obf_8f171f98ab5da0ae int64 = 0
var __obf_ecc647e1cb0ec515 int64 = time.Now().UnixNano() / 1e6

func __obf_4a2251d2274fd09f() string {
	__obf_6eaeb0a8242c32cc := make([]byte, 64)
	__obf_6eaeb0a8242c32cc = __obf_6eaeb0a8242c32cc[:runtime.Stack(__obf_6eaeb0a8242c32cc, false)]
	__obf_6eaeb0a8242c32cc = bytes.TrimPrefix(__obf_6eaeb0a8242c32cc, []byte("goroutine "))
	__obf_6eaeb0a8242c32cc = __obf_6eaeb0a8242c32cc[:bytes.IndexByte(__obf_6eaeb0a8242c32cc, ' ')]
	__obf_646fe4d836d0f8c9, _ := strconv.ParseUint(string(__obf_6eaeb0a8242c32cc), 10, 64)
	__obf_0548a2b1b2203abb := time.Now().UnixNano() / 1e6
	__obf_dd116f596e84db63 := atomic.AddInt64(&__obf_8f171f98ab5da0ae, 1)
	__obf_885cee8e18f2c5e0 := mrand.Int63()
	__obf_214e21a826c5f135 := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_ecc647e1cb0ec515, __obf_646fe4d836d0f8c9, __obf_0548a2b1b2203abb, __obf_dd116f596e84db63, __obf_885cee8e18f2c5e0)
	__obf_010efe188a625b0f := md5.New()
	__obf_010efe188a625b0f.
		Write([]byte(__obf_214e21a826c5f135))
	return hex.EncodeToString(__obf_010efe188a625b0f.Sum(nil))
}

// generateSignature 生成签名
func __obf_1212d7c248dd52c5(__obf_6121932bb4480036 string, __obf_723833b2544b27f1 string) string {
	var __obf_75ae8c8f71f5227e bytes.Buffer
	__obf_75ae8c8f71f5227e.
		WriteString("GET&")
	__obf_75ae8c8f71f5227e.
		WriteString(__obf_786b19c67f51c958("/"))
	__obf_75ae8c8f71f5227e.
		WriteString("&")
	__obf_75ae8c8f71f5227e.
		WriteString(__obf_786b19c67f51c958(__obf_6121932bb4480036))
	__obf_010efe188a625b0f := hmac.New(sha1.New, []byte(__obf_723833b2544b27f1+"&"))
	__obf_010efe188a625b0f.
		Write(__obf_75ae8c8f71f5227e.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_010efe188a625b0f.Sum(nil))
}
