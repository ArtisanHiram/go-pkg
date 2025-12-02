package __obf_67ff2b45a2ee7325

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CurrentDir(__obf_7b00369797a8c02a ...string) (string, error) {
	__obf_68e8626906343965, __obf_3d6b93bb4b532726 := filepath.Abs(filepath.Dir(os.Args[0]))
	if __obf_3d6b93bb4b532726 != nil {
		return "", __obf_3d6b93bb4b532726
	}
	__obf_1fcc7e3df8247d33 := strings.Replace(__obf_68e8626906343965, "\\", "/", -1)
	__obf_2582d4138c4d465d := filepath.Join(__obf_7b00369797a8c02a...)
	__obf_2582d4138c4d465d = filepath.Join(__obf_1fcc7e3df8247d33, __obf_2582d4138c4d465d)
	return __obf_2582d4138c4d465d, nil
}

func ResolveURL(__obf_f45e47b1d7551b1e *url.URL, __obf_1fcc7e3df8247d33 string) string {
	if strings.HasPrefix(__obf_1fcc7e3df8247d33, "https://") || strings.HasPrefix(__obf_1fcc7e3df8247d33, "http://") {
		return __obf_1fcc7e3df8247d33
	}
	var __obf_976d6e9b8650037a string
	if strings.Index(__obf_1fcc7e3df8247d33, "/") == 0 {
		__obf_976d6e9b8650037a = __obf_f45e47b1d7551b1e.Scheme + "://" + __obf_f45e47b1d7551b1e.Host
	} else {
		__obf_912b94b728bd4e6e := __obf_f45e47b1d7551b1e.String()
		__obf_976d6e9b8650037a = __obf_912b94b728bd4e6e[0:strings.LastIndex(__obf_912b94b728bd4e6e, "/")]
	}
	return __obf_976d6e9b8650037a + path.Join("/", __obf_1fcc7e3df8247d33)
}

func DrawProgressBar(__obf_f9cf8eb2d25f0e30 string, __obf_40c7921868df5884 float32, __obf_846ee76dc46320fb int, __obf_15c8682cb7298295 ...string) {
	__obf_6ccfe5ac85b36caa := int(__obf_40c7921868df5884 * float32(__obf_846ee76dc46320fb))
	__obf_4d5d7a36584c9be7 := fmt.Sprintf("[%s] %s%*s %6.2f%% %s", __obf_f9cf8eb2d25f0e30, strings.Repeat("■", __obf_6ccfe5ac85b36caa), __obf_846ee76dc46320fb-__obf_6ccfe5ac85b36caa, "", __obf_40c7921868df5884*100, strings.Join(__obf_15c8682cb7298295, ""))
	fmt.Print("\r" + __obf_4d5d7a36584c9be7)
}
