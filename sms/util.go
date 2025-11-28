package __obf_cb62198a5f8c0e2c

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

var __obf_21b071081aae0d7a int64 = 0
var __obf_0a5c8fde46581f74 int64 = time.Now().UnixNano() / 1e6

func __obf_1f0e3362918c2912() string {

	__obf_bd05f014be491fb8 := make([]byte, 64)
	__obf_bd05f014be491fb8 = __obf_bd05f014be491fb8[:runtime.Stack(__obf_bd05f014be491fb8, false)]
	__obf_bd05f014be491fb8 = bytes.TrimPrefix(__obf_bd05f014be491fb8, []byte("goroutine "))
	__obf_bd05f014be491fb8 = __obf_bd05f014be491fb8[:bytes.IndexByte(__obf_bd05f014be491fb8, ' ')]
	__obf_983b389d4bce2818, _ := strconv.ParseUint(string(__obf_bd05f014be491fb8), 10, 64)

	__obf_a457d50604ca7404 := time.Now().UnixNano() / 1e6
	__obf_03231da4f5015b66 := atomic.AddInt64(&__obf_21b071081aae0d7a, 1)
	__obf_5be48821346d59c8 := mrand.Int63()
	__obf_88cd55b9ab3ed8af := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_0a5c8fde46581f74, __obf_983b389d4bce2818, __obf_a457d50604ca7404, __obf_03231da4f5015b66, __obf_5be48821346d59c8)
	__obf_58eae71206ec5018 := md5.New()
	__obf_58eae71206ec5018.Write([]byte(__obf_88cd55b9ab3ed8af))
	return hex.EncodeToString(__obf_58eae71206ec5018.Sum(nil))
}

// generateSignature 生成签名
func __obf_e993b87f82f2044b(__obf_2346704bb5e01c49 string, __obf_af4b82e4375e005a string) string {
	var __obf_cde918037b664289 bytes.Buffer
	__obf_cde918037b664289.WriteString("GET&")
	__obf_cde918037b664289.WriteString(__obf_fdce8fc899639dcd("/"))
	__obf_cde918037b664289.WriteString("&")
	__obf_cde918037b664289.WriteString(__obf_fdce8fc899639dcd(__obf_2346704bb5e01c49))

	__obf_58eae71206ec5018 := hmac.New(sha1.New, []byte(__obf_af4b82e4375e005a+"&"))
	__obf_58eae71206ec5018.Write(__obf_cde918037b664289.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_58eae71206ec5018.Sum(nil))
}
