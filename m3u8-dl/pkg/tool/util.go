package __obf_c37d95bc12b416d3

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(__obf_df67fb16427c4171 ...string) (string, error) {
	__obf_d2ed974242179d3d, __obf_3ce485578a57d5a7 := filepath.Abs(filepath.Dir(os.Args[0]))
	if __obf_3ce485578a57d5a7 != nil {
		return "", __obf_3ce485578a57d5a7
	}
	__obf_9690bf53e550d284 := strings.Replace(__obf_d2ed974242179d3d, "\\", "/", -1)
	__obf_3dee22a769fbf456 := filepath.Join(__obf_df67fb16427c4171...)
	__obf_3dee22a769fbf456 = filepath.Join(__obf_9690bf53e550d284, __obf_3dee22a769fbf456)
	return __obf_3dee22a769fbf456, nil
}

func ResolveURL(__obf_b67a06a6639a2065 *url.URL, __obf_9690bf53e550d284 string) string {
	if strings.HasPrefix(__obf_9690bf53e550d284, "https://") || strings.HasPrefix(__obf_9690bf53e550d284, "http://") {
		return __obf_9690bf53e550d284
	}
	var __obf_573cac9df54992dc string
	if strings.Index(__obf_9690bf53e550d284, "/") == 0 {
		__obf_573cac9df54992dc = __obf_b67a06a6639a2065.Scheme + "://" + __obf_b67a06a6639a2065.Host
	} else {
		__obf_95f78630f260bc5b := __obf_b67a06a6639a2065.String()
		__obf_573cac9df54992dc = __obf_95f78630f260bc5b[0:strings.LastIndex(__obf_95f78630f260bc5b, "/")]
	}
	return __obf_573cac9df54992dc + path.Join("/", __obf_9690bf53e550d284)
}

func DrawProgressBar(__obf_e17a005a6ed4fd66 string, __obf_62675bd19564e17c float32, __obf_66432ac9c163fdb5 int, __obf_b6757668729d0980 ...string) {
	__obf_56caa8b3a2bb325e := int(__obf_62675bd19564e17c * float32(__obf_66432ac9c163fdb5))
	__obf_f45046999054c1c9 := fmt.Sprintf("[%s] %s%*s %6.2f%% %s", __obf_e17a005a6ed4fd66, strings.Repeat("■", __obf_56caa8b3a2bb325e), __obf_66432ac9c163fdb5-__obf_56caa8b3a2bb325e, "", __obf_62675bd19564e17c*100, strings.Join(__obf_b6757668729d0980, ""))
	fmt.Print("\r" + __obf_f45046999054c1c9)
}
