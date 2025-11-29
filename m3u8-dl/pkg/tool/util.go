package __obf_61bc591a775c8985

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(__obf_e636267f55df438a ...string) (string, error) {
	__obf_0a3e8c2a05960a43, __obf_48065909968843ec := filepath.Abs(filepath.Dir(os.Args[0]))
	if __obf_48065909968843ec != nil {
		return "", __obf_48065909968843ec
	}
	__obf_980149e3a41644c9 := strings.Replace(__obf_0a3e8c2a05960a43, "\\", "/", -1)
	__obf_4a95416e6d172727 := filepath.Join(__obf_e636267f55df438a...)
	__obf_4a95416e6d172727 = filepath.Join(__obf_980149e3a41644c9, __obf_4a95416e6d172727)
	return __obf_4a95416e6d172727, nil
}

func ResolveURL(__obf_8a27d47dd36df03f *url.URL, __obf_980149e3a41644c9 string) string {
	if strings.HasPrefix(__obf_980149e3a41644c9, "https://") || strings.HasPrefix(__obf_980149e3a41644c9, "http://") {
		return __obf_980149e3a41644c9
	}
	var __obf_ef13cecbc398015d string
	if strings.Index(__obf_980149e3a41644c9, "/") == 0 {
		__obf_ef13cecbc398015d = __obf_8a27d47dd36df03f.Scheme + "://" + __obf_8a27d47dd36df03f.Host
	} else {
		__obf_3e441270ad20ac28 := __obf_8a27d47dd36df03f.String()
		__obf_ef13cecbc398015d = __obf_3e441270ad20ac28[0:strings.LastIndex(__obf_3e441270ad20ac28, "/")]
	}
	return __obf_ef13cecbc398015d + path.Join("/", __obf_980149e3a41644c9)
}

func DrawProgressBar(__obf_8d375f91e1c1904a string, __obf_23c3920436c03fcd float32, __obf_554351a19bd71b15 int, __obf_56307ed5705bc8d2 ...string) {
	__obf_8926187c7b103473 := int(__obf_23c3920436c03fcd * float32(__obf_554351a19bd71b15))
	__obf_c54ee8a73158c0cf := fmt.Sprintf("[%s] %s%*s %6.2f%% %s", __obf_8d375f91e1c1904a, strings.Repeat("■", __obf_8926187c7b103473), __obf_554351a19bd71b15-__obf_8926187c7b103473, "", __obf_23c3920436c03fcd*100, strings.Join(__obf_56307ed5705bc8d2, ""))
	fmt.Print("\r" + __obf_c54ee8a73158c0cf)
}
