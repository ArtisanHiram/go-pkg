package __obf_dc51e1c30a41549a

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

var __obf_8a31d14f53a88e53 int64 = 0
var __obf_27400ba249d479b0 int64 = time.Now().UnixNano() / 1e6

func __obf_97dd4bb2856c2c3b() string {
	__obf_199e588fe0fa9f9c := make([]byte, 64)
	__obf_199e588fe0fa9f9c = __obf_199e588fe0fa9f9c[:runtime.Stack(__obf_199e588fe0fa9f9c, false)]
	__obf_199e588fe0fa9f9c = bytes.TrimPrefix(__obf_199e588fe0fa9f9c, []byte("goroutine "))
	__obf_199e588fe0fa9f9c = __obf_199e588fe0fa9f9c[:bytes.IndexByte(__obf_199e588fe0fa9f9c, ' ')]
	__obf_e04e12f52f6fb738, _ := strconv.ParseUint(string(__obf_199e588fe0fa9f9c), 10, 64)
	__obf_331bac21aaa21853 := time.Now().UnixNano() / 1e6
	__obf_e2eac0d0e09ea65f := atomic.AddInt64(&__obf_8a31d14f53a88e53, 1)
	__obf_fc9281bdfcd960f9 := mrand.Int63()
	__obf_864e58597aff8abf := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_27400ba249d479b0, __obf_e04e12f52f6fb738, __obf_331bac21aaa21853, __obf_e2eac0d0e09ea65f, __obf_fc9281bdfcd960f9)
	__obf_d36fd251b480dc63 := md5.New()
	__obf_d36fd251b480dc63.
		Write([]byte(__obf_864e58597aff8abf))
	return hex.EncodeToString(__obf_d36fd251b480dc63.Sum(nil))
}

// generateSignature 生成签名
func __obf_2a1bcfd3101707aa(__obf_ab08003662ad475f string, __obf_fbec4915c4cb08dc string) string {
	var __obf_a5b853582a010c81 bytes.Buffer
	__obf_a5b853582a010c81.
		WriteString("GET&")
	__obf_a5b853582a010c81.
		WriteString(__obf_11b188aaefdfc57a("/"))
	__obf_a5b853582a010c81.
		WriteString("&")
	__obf_a5b853582a010c81.
		WriteString(__obf_11b188aaefdfc57a(__obf_ab08003662ad475f))
	__obf_d36fd251b480dc63 := hmac.New(sha1.New, []byte(__obf_fbec4915c4cb08dc+"&"))
	__obf_d36fd251b480dc63.
		Write(__obf_a5b853582a010c81.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_d36fd251b480dc63.Sum(nil))
}
