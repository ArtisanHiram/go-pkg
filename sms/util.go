package __obf_37e82024588137a9

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

var __obf_5c505b459301f4ab int64 = 0
var __obf_a14a5565434b9ed1 int64 = time.Now().UnixNano() / 1e6

func __obf_b74d37b56b7256e3() string {
	__obf_b24fa5c36f71482a := make([]byte, 64)
	__obf_b24fa5c36f71482a = __obf_b24fa5c36f71482a[:runtime.Stack(__obf_b24fa5c36f71482a, false)]
	__obf_b24fa5c36f71482a = bytes.TrimPrefix(__obf_b24fa5c36f71482a, []byte("goroutine "))
	__obf_b24fa5c36f71482a = __obf_b24fa5c36f71482a[:bytes.IndexByte(__obf_b24fa5c36f71482a, ' ')]
	__obf_42b1d29286c346b1, _ := strconv.ParseUint(string(__obf_b24fa5c36f71482a), 10, 64)
	__obf_5d88e7ae8b7cfcee := time.Now().UnixNano() / 1e6
	__obf_34dd62eb48d1e658 := atomic.AddInt64(&__obf_5c505b459301f4ab, 1)
	__obf_9db043a304497c34 := mrand.Int63()
	__obf_2df9696bd59707ad := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_a14a5565434b9ed1, __obf_42b1d29286c346b1, __obf_5d88e7ae8b7cfcee, __obf_34dd62eb48d1e658, __obf_9db043a304497c34)
	__obf_a0a870ebbbd2494b := md5.New()
	__obf_a0a870ebbbd2494b.
		Write([]byte(__obf_2df9696bd59707ad))
	return hex.EncodeToString(__obf_a0a870ebbbd2494b.Sum(nil))
}

// generateSignature 生成签名
func __obf_e5ce4633ffe9b986(__obf_d1a5d2de53e2f453 string, __obf_5e3467f51b91dd9a string) string {
	var __obf_3456da504fb3f6ff bytes.Buffer
	__obf_3456da504fb3f6ff.
		WriteString("GET&")
	__obf_3456da504fb3f6ff.
		WriteString(__obf_792a3f6a70b79377("/"))
	__obf_3456da504fb3f6ff.
		WriteString("&")
	__obf_3456da504fb3f6ff.
		WriteString(__obf_792a3f6a70b79377(__obf_d1a5d2de53e2f453))
	__obf_a0a870ebbbd2494b := hmac.New(sha1.New, []byte(__obf_5e3467f51b91dd9a+"&"))
	__obf_a0a870ebbbd2494b.
		Write(__obf_3456da504fb3f6ff.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_a0a870ebbbd2494b.Sum(nil))
}
