package __obf_596164137deb9992

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_7df84a7884d421b2 string, __obf_9009b827c826dff7 string, __obf_f96ee6114943a263 []string) (string, string, error) {
	__obf_e0a2165b99b4cc70 := "main.js"
	__obf_2a92ad6e632f4f6a := // Find all typescript files and build compile command
		util.FilterByExtension(__obf_f96ee6114943a263, ".ts")
	__obf_1cd43d7bd68349ac := append([]string{"tsc", "-out", __obf_e0a2165b99b4cc70}, __obf_2a92ad6e632f4f6a...)
	__obf_c2d3cfd749ddd458,

		// Compile to javascript
		__obf_b849c432d8f84e23, __obf_b74563acf9cbda6f := cmd.Run(__obf_7df84a7884d421b2, __obf_1cd43d7bd68349ac...)
	if __obf_b74563acf9cbda6f != nil {
		return __obf_c2d3cfd749ddd458, __obf_b849c432d8f84e23, __obf_b74563acf9cbda6f
	}

	return cmd.RunStdin(__obf_7df84a7884d421b2, __obf_9009b827c826dff7, "node", __obf_e0a2165b99b4cc70)
}
