package __obf_28053862c3d647b8

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

var __obf_d6c7aab7517ba912 int64 = 0
var __obf_c58c1da4a528a84e int64 = time.Now().UnixNano() / 1e6

func __obf_2a2a3aa08c34cdb2() string {

	__obf_30bbdf520e0c12af := make([]byte, 64)
	__obf_30bbdf520e0c12af = __obf_30bbdf520e0c12af[:runtime.Stack(__obf_30bbdf520e0c12af, false)]
	__obf_30bbdf520e0c12af = bytes.TrimPrefix(__obf_30bbdf520e0c12af, []byte("goroutine "))
	__obf_30bbdf520e0c12af = __obf_30bbdf520e0c12af[:bytes.IndexByte(__obf_30bbdf520e0c12af, ' ')]
	__obf_2155d5d299e7dfe9, _ := strconv.ParseUint(string(__obf_30bbdf520e0c12af), 10, 64)

	__obf_d0ca0551eac68bcb := time.Now().UnixNano() / 1e6
	__obf_2c7d3a90d1d2085f := atomic.AddInt64(&__obf_d6c7aab7517ba912, 1)
	__obf_ef4bb1ea1519b12e := mrand.Int63()
	__obf_cb219f176a3f76ed := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_c58c1da4a528a84e, __obf_2155d5d299e7dfe9, __obf_d0ca0551eac68bcb, __obf_2c7d3a90d1d2085f, __obf_ef4bb1ea1519b12e)
	__obf_33697329ad3aa86d := md5.New()
	__obf_33697329ad3aa86d.Write([]byte(__obf_cb219f176a3f76ed))
	return hex.EncodeToString(__obf_33697329ad3aa86d.Sum(nil))
}

// generateSignature 生成签名
func __obf_ae740a1c7efd2892(__obf_6bda852866144048 string, __obf_ead40cbd1507b913 string) string {
	var __obf_81e749202e3486d0 bytes.Buffer
	__obf_81e749202e3486d0.WriteString("GET&")
	__obf_81e749202e3486d0.WriteString(__obf_05b630c8ad786500("/"))
	__obf_81e749202e3486d0.WriteString("&")
	__obf_81e749202e3486d0.WriteString(__obf_05b630c8ad786500(__obf_6bda852866144048))

	__obf_33697329ad3aa86d := hmac.New(sha1.New, []byte(__obf_ead40cbd1507b913+"&"))
	__obf_33697329ad3aa86d.Write(__obf_81e749202e3486d0.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_33697329ad3aa86d.Sum(nil))
}
