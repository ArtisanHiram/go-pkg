package __obf_53be58cd44fa013e

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_024a5ea5c20da32e string, __obf_3a74bdb7c55d6d29 string, __obf_f360dcf08994a4dc []string) (string, string, error) {
	__obf_5ae190181e73749a, __obf_a1f1a68e27e48053, __obf_b9975fbaf2719e2c := cmd.Run(__obf_024a5ea5c20da32e, "javac", __obf_f360dcf08994a4dc[0])
	if __obf_b9975fbaf2719e2c != nil {
		return __obf_5ae190181e73749a, __obf_a1f1a68e27e48053, __obf_b9975fbaf2719e2c
	}
	return cmd.RunStdin(__obf_024a5ea5c20da32e, __obf_3a74bdb7c55d6d29, "java", __obf_388aa30c0fae7f88(__obf_f360dcf08994a4dc[0]))
}

func __obf_388aa30c0fae7f88(__obf_a7eb72d035b25b11 string) string {
	__obf_ca60dd93b2bbcc6c := filepath.Ext(__obf_a7eb72d035b25b11)
	return __obf_a7eb72d035b25b11[0 : len(__obf_a7eb72d035b25b11)-len(__obf_ca60dd93b2bbcc6c)]
}
