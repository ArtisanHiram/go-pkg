package __obf_0f0912cb4961947f

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

var __obf_877eb8a0b5039297 int64 = 0
var __obf_62b02b0165f64bae int64 = time.Now().UnixNano() / 1e6

func __obf_d01954fe5f7f19ca() string {
	__obf_2e053c5802479b30 := make([]byte, 64)
	__obf_2e053c5802479b30 = __obf_2e053c5802479b30[:runtime.Stack(__obf_2e053c5802479b30, false)]
	__obf_2e053c5802479b30 = bytes.TrimPrefix(__obf_2e053c5802479b30, []byte("goroutine "))
	__obf_2e053c5802479b30 = __obf_2e053c5802479b30[:bytes.IndexByte(__obf_2e053c5802479b30, ' ')]
	__obf_9302d45641f8ec37, _ := strconv.ParseUint(string(__obf_2e053c5802479b30), 10, 64)
	__obf_e57707f639cfabd8 := time.Now().UnixNano() / 1e6
	__obf_b05170b2d8482caf := atomic.AddInt64(&__obf_877eb8a0b5039297, 1)
	__obf_22da43146360a889 := mrand.Int63()
	__obf_2f8ae38c3f0de822 := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_62b02b0165f64bae, __obf_9302d45641f8ec37, __obf_e57707f639cfabd8, __obf_b05170b2d8482caf, __obf_22da43146360a889)
	__obf_a49911eaf1212263 := md5.New()
	__obf_a49911eaf1212263.
		Write([]byte(__obf_2f8ae38c3f0de822))
	return hex.EncodeToString(__obf_a49911eaf1212263.Sum(nil))
}

// generateSignature 生成签名
func __obf_911e53cade1302f7(__obf_c2b27722a4309195 string, __obf_dd94e101d195a7c7 string) string {
	var __obf_1ef253f8062a91a1 bytes.Buffer
	__obf_1ef253f8062a91a1.
		WriteString("GET&")
	__obf_1ef253f8062a91a1.
		WriteString(__obf_4e6893075e8516f5("/"))
	__obf_1ef253f8062a91a1.
		WriteString("&")
	__obf_1ef253f8062a91a1.
		WriteString(__obf_4e6893075e8516f5(__obf_c2b27722a4309195))
	__obf_a49911eaf1212263 := hmac.New(sha1.New, []byte(__obf_dd94e101d195a7c7+"&"))
	__obf_a49911eaf1212263.
		Write(__obf_1ef253f8062a91a1.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_a49911eaf1212263.Sum(nil))
}
