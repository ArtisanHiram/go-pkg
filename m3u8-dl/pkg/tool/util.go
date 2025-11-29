package __obf_22f65c3e757f8811

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(__obf_3a9b395b0f585b06 ...string) (string, error) {
	__obf_0336691b7acc25f3, __obf_f749646f37e10859 := filepath.Abs(filepath.Dir(os.Args[0]))
	if __obf_f749646f37e10859 != nil {
		return "", __obf_f749646f37e10859
	}
	__obf_274b0aab616567ff := strings.Replace(__obf_0336691b7acc25f3, "\\", "/", -1)
	__obf_3371341914786ca2 := filepath.Join(__obf_3a9b395b0f585b06...)
	__obf_3371341914786ca2 = filepath.Join(__obf_274b0aab616567ff, __obf_3371341914786ca2)
	return __obf_3371341914786ca2, nil
}

func ResolveURL(__obf_1e88df0f0b966b94 *url.URL, __obf_274b0aab616567ff string) string {
	if strings.HasPrefix(__obf_274b0aab616567ff, "https://") || strings.HasPrefix(__obf_274b0aab616567ff, "http://") {
		return __obf_274b0aab616567ff
	}
	var __obf_008087e22c30d01a string
	if strings.Index(__obf_274b0aab616567ff, "/") == 0 {
		__obf_008087e22c30d01a = __obf_1e88df0f0b966b94.Scheme + "://" + __obf_1e88df0f0b966b94.Host
	} else {
		__obf_8307feab292a18a0 := __obf_1e88df0f0b966b94.String()
		__obf_008087e22c30d01a = __obf_8307feab292a18a0[0:strings.LastIndex(__obf_8307feab292a18a0, "/")]
	}
	return __obf_008087e22c30d01a + path.Join("/", __obf_274b0aab616567ff)
}

func DrawProgressBar(__obf_5ed882628459088f string, __obf_01599d90584b2a96 float32, __obf_a2760c3951a1ae6d int, __obf_0b807bd08fdce7f4 ...string) {
	__obf_cadeb652b5931706 := int(__obf_01599d90584b2a96 * float32(__obf_a2760c3951a1ae6d))
	__obf_1df6d77b42222a9b := fmt.Sprintf("[%s] %s%*s %6.2f%% %s", __obf_5ed882628459088f, strings.Repeat("■", __obf_cadeb652b5931706), __obf_a2760c3951a1ae6d-__obf_cadeb652b5931706, "", __obf_01599d90584b2a96*100, strings.Join(__obf_0b807bd08fdce7f4, ""))
	fmt.Print("\r" + __obf_1df6d77b42222a9b)
}
