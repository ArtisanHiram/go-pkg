package __obf_a307862f84d54cc6

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

var __obf_ddb437d5df6b7338 int64 = 0
var __obf_7bbdca9bbe4109ac int64 = time.Now().UnixNano() / 1e6

func __obf_a49f1d27524bcbdb() string {

	__obf_856053960379f1d3 := make([]byte, 64)
	__obf_856053960379f1d3 = __obf_856053960379f1d3[:runtime.Stack(__obf_856053960379f1d3, false)]
	__obf_856053960379f1d3 = bytes.TrimPrefix(__obf_856053960379f1d3, []byte("goroutine "))
	__obf_856053960379f1d3 = __obf_856053960379f1d3[:bytes.IndexByte(__obf_856053960379f1d3, ' ')]
	__obf_e857b55c8ad28868, _ := strconv.ParseUint(string(__obf_856053960379f1d3), 10, 64)

	__obf_7b5de4f88a4e708f := time.Now().UnixNano() / 1e6
	__obf_7c2677d114280cad := atomic.AddInt64(&__obf_ddb437d5df6b7338, 1)
	__obf_df93f6fa08360578 := mrand.Int63()
	__obf_3b91e2714fae3e9b := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_7bbdca9bbe4109ac, __obf_e857b55c8ad28868, __obf_7b5de4f88a4e708f, __obf_7c2677d114280cad, __obf_df93f6fa08360578)
	__obf_66f2bca1c6c7d88d := md5.New()
	__obf_66f2bca1c6c7d88d.Write([]byte(__obf_3b91e2714fae3e9b))
	return hex.EncodeToString(__obf_66f2bca1c6c7d88d.Sum(nil))
}

// generateSignature 生成签名
func __obf_cdacf59f3a4d15f3(__obf_47e75f2fdeec63af string, __obf_0d83fb6d65187b1d string) string {
	var __obf_d21aa78f6d95f3c1 bytes.Buffer
	__obf_d21aa78f6d95f3c1.WriteString("GET&")
	__obf_d21aa78f6d95f3c1.WriteString(__obf_03c24bcdceda96ba("/"))
	__obf_d21aa78f6d95f3c1.WriteString("&")
	__obf_d21aa78f6d95f3c1.WriteString(__obf_03c24bcdceda96ba(__obf_47e75f2fdeec63af))

	__obf_66f2bca1c6c7d88d := hmac.New(sha1.New, []byte(__obf_0d83fb6d65187b1d+"&"))
	__obf_66f2bca1c6c7d88d.Write(__obf_d21aa78f6d95f3c1.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_66f2bca1c6c7d88d.Sum(nil))
}
