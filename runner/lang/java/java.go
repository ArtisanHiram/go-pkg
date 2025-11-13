package __obf_aac102c67703a1ee

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_1260d82b37f5b57f string, __obf_09e083f3f276e6d9 string, __obf_182c73ecb73a3259 []string) (string, string, error) {
	__obf_6a628d99e1f129e9, __obf_e96532c0593d57e2, __obf_d40099bc9cce47d1 := cmd.Run(__obf_1260d82b37f5b57f, "javac", __obf_182c73ecb73a3259[0])
	if __obf_d40099bc9cce47d1 != nil {
		return __obf_6a628d99e1f129e9, __obf_e96532c0593d57e2, __obf_d40099bc9cce47d1
	}
	return cmd.RunStdin(__obf_1260d82b37f5b57f, __obf_09e083f3f276e6d9, "java", __obf_612b982540bd456c(__obf_182c73ecb73a3259[0]))
}

func __obf_612b982540bd456c(__obf_95f79699ffca00fb string) string {
	__obf_89f566f21aad9684 := filepath.Ext(__obf_95f79699ffca00fb)
	return __obf_95f79699ffca00fb[0 : len(__obf_95f79699ffca00fb)-len(__obf_89f566f21aad9684)]
}
