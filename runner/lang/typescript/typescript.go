package __obf_a779885fe2c8c3c0

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	util "github.com/ArtisanHiram/go-pkg/runner/util"
)

func Run(__obf_a4a9ed8b88209546 string, __obf_aa9b5e8988ada14d string, __obf_9b54458786148208 []string) (string, string, error) {
	__obf_55183b4d11242434 := "main.js"
	__obf_ba08137e557ad4c7 := // Find all typescript files and build compile command
		util.FilterByExtension(__obf_9b54458786148208, ".ts")
	__obf_ec1c9f1e78b75eaf := append([]string{"tsc", "-out", __obf_55183b4d11242434}, __obf_ba08137e557ad4c7...)
	__obf_3e25117b795392c1,

		// Compile to javascript
		__obf_bcbf4e2f1bba8fc1, __obf_bc05fd5db54eddfe := cmd.Run(__obf_a4a9ed8b88209546, __obf_ec1c9f1e78b75eaf...)
	if __obf_bc05fd5db54eddfe != nil {
		return __obf_3e25117b795392c1, __obf_bcbf4e2f1bba8fc1, __obf_bc05fd5db54eddfe
	}

	return cmd.RunStdin(__obf_a4a9ed8b88209546, __obf_aa9b5e8988ada14d, "node", __obf_55183b4d11242434)
}
