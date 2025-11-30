package __obf_33bc7bf683859fa2

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(__obf_2c04a0ef2cf805da ...string) (string, error) {
	__obf_f15c1d0192342767, __obf_19a0c091a533826e := filepath.Abs(filepath.Dir(os.Args[0]))
	if __obf_19a0c091a533826e != nil {
		return "", __obf_19a0c091a533826e
	}
	__obf_51db33a31204fb12 := strings.Replace(__obf_f15c1d0192342767, "\\", "/", -1)
	__obf_be70937b948be643 := filepath.Join(__obf_2c04a0ef2cf805da...)
	__obf_be70937b948be643 = filepath.Join(__obf_51db33a31204fb12, __obf_be70937b948be643)
	return __obf_be70937b948be643, nil
}

func ResolveURL(__obf_001b764ff95786f9 *url.URL, __obf_51db33a31204fb12 string) string {
	if strings.HasPrefix(__obf_51db33a31204fb12, "https://") || strings.HasPrefix(__obf_51db33a31204fb12, "http://") {
		return __obf_51db33a31204fb12
	}
	var __obf_9b779f54b0be207b string
	if strings.Index(__obf_51db33a31204fb12, "/") == 0 {
		__obf_9b779f54b0be207b = __obf_001b764ff95786f9.Scheme + "://" + __obf_001b764ff95786f9.Host
	} else {
		__obf_2e06e40a5c39ff85 := __obf_001b764ff95786f9.String()
		__obf_9b779f54b0be207b = __obf_2e06e40a5c39ff85[0:strings.LastIndex(__obf_2e06e40a5c39ff85, "/")]
	}
	return __obf_9b779f54b0be207b + path.Join("/", __obf_51db33a31204fb12)
}

func DrawProgressBar(__obf_f4c01cdd0c1f3478 string, __obf_9e4a30afb6c0fbdf float32, __obf_f39a17f9001722d6 int, __obf_cb2360a22eab6b42 ...string) {
	__obf_5176562ede04f4df := int(__obf_9e4a30afb6c0fbdf * float32(__obf_f39a17f9001722d6))
	__obf_bd6126f3888d2ed7 := fmt.Sprintf("[%s] %s%*s %6.2f%% %s", __obf_f4c01cdd0c1f3478, strings.Repeat("■", __obf_5176562ede04f4df), __obf_f39a17f9001722d6-__obf_5176562ede04f4df, "", __obf_9e4a30afb6c0fbdf*100, strings.Join(__obf_cb2360a22eab6b42, ""))
	fmt.Print("\r" + __obf_bd6126f3888d2ed7)
}
