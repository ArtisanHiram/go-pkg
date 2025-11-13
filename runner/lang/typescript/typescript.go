package __obf_6ebe31c5b182b412

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_1696996a61a7d71f string, __obf_b01305601c15d0db string, __obf_d7195192b2b758e2 []string) (string, string, error) {
	__obf_615686b63aa4c140 := "main.js"

	// Find all typescript files and build compile command
	__obf_57c907d08b25907c := util.FilterByExtension(__obf_d7195192b2b758e2, ".ts")
	__obf_1291c3c5fd2781ac := append([]string{"tsc", "-out", __obf_615686b63aa4c140}, __obf_57c907d08b25907c...)

	// Compile to javascript
	__obf_c5591e320bf7c006, __obf_28ac60400999b70e, __obf_99aae9bd1c11cf86 := cmd.Run(__obf_1696996a61a7d71f, __obf_1291c3c5fd2781ac...)
	if __obf_99aae9bd1c11cf86 != nil {
		return __obf_c5591e320bf7c006, __obf_28ac60400999b70e, __obf_99aae9bd1c11cf86
	}

	return cmd.RunStdin(__obf_1696996a61a7d71f, __obf_b01305601c15d0db, "node", __obf_615686b63aa4c140)
}
