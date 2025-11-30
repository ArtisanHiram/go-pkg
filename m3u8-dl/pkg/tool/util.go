package __obf_6f9b75fc21ef8ef6

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(__obf_abb45957eddd84d3 ...string) (string, error) {
	__obf_8f59247f17b63666, __obf_8eefb2cd18f480f3 := filepath.Abs(filepath.Dir(os.Args[0]))
	if __obf_8eefb2cd18f480f3 != nil {
		return "", __obf_8eefb2cd18f480f3
	}
	__obf_9f31aeadffd42691 := strings.Replace(__obf_8f59247f17b63666, "\\", "/", -1)
	__obf_7825783a40b1ce46 := filepath.Join(__obf_abb45957eddd84d3...)
	__obf_7825783a40b1ce46 = filepath.Join(__obf_9f31aeadffd42691, __obf_7825783a40b1ce46)
	return __obf_7825783a40b1ce46, nil
}

func ResolveURL(__obf_9d1386ef1d393320 *url.URL, __obf_9f31aeadffd42691 string) string {
	if strings.HasPrefix(__obf_9f31aeadffd42691, "https://") || strings.HasPrefix(__obf_9f31aeadffd42691, "http://") {
		return __obf_9f31aeadffd42691
	}
	var __obf_2f2aa5b08d5cde94 string
	if strings.Index(__obf_9f31aeadffd42691, "/") == 0 {
		__obf_2f2aa5b08d5cde94 = __obf_9d1386ef1d393320.Scheme + "://" + __obf_9d1386ef1d393320.Host
	} else {
		__obf_63acd623a386d248 := __obf_9d1386ef1d393320.String()
		__obf_2f2aa5b08d5cde94 = __obf_63acd623a386d248[0:strings.LastIndex(__obf_63acd623a386d248, "/")]
	}
	return __obf_2f2aa5b08d5cde94 + path.Join("/", __obf_9f31aeadffd42691)
}

func DrawProgressBar(__obf_59d732801ffa3402 string, __obf_220c9bba723fc881 float32, __obf_f4fe5fc94e7a142d int, __obf_d382afec1f916b80 ...string) {
	__obf_7bb649daced6b7ad := int(__obf_220c9bba723fc881 * float32(__obf_f4fe5fc94e7a142d))
	__obf_c06a92b349b10c13 := fmt.Sprintf("[%s] %s%*s %6.2f%% %s", __obf_59d732801ffa3402, strings.Repeat("■", __obf_7bb649daced6b7ad), __obf_f4fe5fc94e7a142d-__obf_7bb649daced6b7ad, "", __obf_220c9bba723fc881*100, strings.Join(__obf_d382afec1f916b80, ""))
	fmt.Print("\r" + __obf_c06a92b349b10c13)
}
