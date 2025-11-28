package __obf_52bdfa18dc226ac6

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

var __obf_a753196986384a9f int64 = 0
var __obf_689b757be897f6ac int64 = time.Now().UnixNano() / 1e6

func __obf_126da17986d8b304() string {

	__obf_43f093a485273158 := make([]byte, 64)
	__obf_43f093a485273158 = __obf_43f093a485273158[:runtime.Stack(__obf_43f093a485273158, false)]
	__obf_43f093a485273158 = bytes.TrimPrefix(__obf_43f093a485273158, []byte("goroutine "))
	__obf_43f093a485273158 = __obf_43f093a485273158[:bytes.IndexByte(__obf_43f093a485273158, ' ')]
	__obf_65e3a6ca9e3a48a9, _ := strconv.ParseUint(string(__obf_43f093a485273158), 10, 64)

	__obf_9ded9b5a03b80c69 := time.Now().UnixNano() / 1e6
	__obf_a20803a449c5a185 := atomic.AddInt64(&__obf_a753196986384a9f, 1)
	__obf_fd837f41361d11d2 := mrand.Int63()
	__obf_8837993ac980126c := fmt.Sprintf("%d-%d-%d-%d-%d", __obf_689b757be897f6ac, __obf_65e3a6ca9e3a48a9, __obf_9ded9b5a03b80c69, __obf_a20803a449c5a185, __obf_fd837f41361d11d2)
	__obf_2ba09f2d84fb0e3a := md5.New()
	__obf_2ba09f2d84fb0e3a.Write([]byte(__obf_8837993ac980126c))
	return hex.EncodeToString(__obf_2ba09f2d84fb0e3a.Sum(nil))
}

// generateSignature 生成签名
func __obf_7c84f133c92d8b15(__obf_4d1db89a54b812c0 string, __obf_2cda83fc4f45e711 string) string {
	var __obf_59041dd12ee742a7 bytes.Buffer
	__obf_59041dd12ee742a7.WriteString("GET&")
	__obf_59041dd12ee742a7.WriteString(__obf_c6816d7254c31423("/"))
	__obf_59041dd12ee742a7.WriteString("&")
	__obf_59041dd12ee742a7.WriteString(__obf_c6816d7254c31423(__obf_4d1db89a54b812c0))

	__obf_2ba09f2d84fb0e3a := hmac.New(sha1.New, []byte(__obf_2cda83fc4f45e711+"&"))
	__obf_2ba09f2d84fb0e3a.Write(__obf_59041dd12ee742a7.Bytes())
	return base64.StdEncoding.EncodeToString(__obf_2ba09f2d84fb0e3a.Sum(nil))
}
