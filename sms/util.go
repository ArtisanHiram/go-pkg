package __obf_721a4aff228e6809

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

var __obf_afca0ee1f21c74f9 int64 = 0
var __obf_0b692cc7cb51168c int64 = time.Now().UnixNano() / 1e6

func __obf_5fc7e9a6b9ce2950() string {

	__obf_50a119c22c210e18 := make([]byte, 64)
	__obf_50a119c22c210e18 = __obf_50a119c22c210e18[:runtime.Stack(__obf_50a119c22c210e18, false)]
	__obf_50a119c22c210e18 = bytes.TrimPrefix(__obf_50a119c22c210e18, []byte("goroutine "))
	__obf_50a119c22c210e18 = __obf_50a119c22c210e18[:bytes.IndexByte(__obf_50a119c22c210e18, ' ')]
	__obf_bab8f3fc87db33a5, _ := strconv.ParseUint(string(__obf_50a119c22c210e18), 10, 64)

	__obf_88c91fa4ad717a74 := time.Now().UnixNano() / 1e6
	__obf_a16000cd5f05bf61 := atomic.AddInt64(&__obf_afca0ee1f21c74f9, 1)
	__obf_8e6ac0cbf8ea5b53 := mrand.Int63()
	__obf_032eaa5d5ad2adf2 := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_0b692cc7cb51168c, __obf_bab8f3fc87db33a5, __obf_88c91fa4ad717a74, __obf_a16000cd5f05bf61, __obf_8e6ac0cbf8ea5b53)
	__obf_e967ae08d4f819fa := md5.New()
	__obf_e967ae08d4f819fa.Write([]byte(__obf_032eaa5d5ad2adf2))
	return hex.EncodeToString(__obf_e967ae08d4f819fa.Sum(nil))
}

// generateSignature 生成签名
func __obf_7312697087077e2b(__obf_14d9c902120ed57d string, __obf_3481be3dcb25fb5d string) string {
	var __obf_d5c391136fe9049b bytes.Buffer
	__obf_d5c391136fe9049b.WriteString("GET&")
	__obf_d5c391136fe9049b.WriteString(__obf_ec014832ed2718dc("/"))
	__obf_d5c391136fe9049b.WriteString("&")
	__obf_d5c391136fe9049b.WriteString(__obf_ec014832ed2718dc(__obf_14d9c902120ed57d))

	__obf_e967ae08d4f819fa := hmac.New(sha1.New, []byte(__obf_3481be3dcb25fb5d+"&"))
	__obf_e967ae08d4f819fa.Write(__obf_d5c391136fe9049b.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_e967ae08d4f819fa.Sum(nil))
}
