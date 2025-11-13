package __obf_6a2a96d13fe16231

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_9c8b77b544eb3c05 string, __obf_68d7187f32e653f2 string, __obf_1c999056e167eaae []string) (string, string, error) {
	__obf_c37540905f7c483e := "main.out"
	__obf_8423e90076928fbe, __obf_fcd50f835fda3286, __obf_d4be3326e9cac003 := cmd.Run(__obf_9c8b77b544eb3c05, "rustc", "-o", __obf_c37540905f7c483e, __obf_1c999056e167eaae[0])
	if __obf_d4be3326e9cac003 != nil {
		return __obf_8423e90076928fbe, __obf_fcd50f835fda3286, __obf_d4be3326e9cac003
	}

	return cmd.RunStdin(__obf_9c8b77b544eb3c05, __obf_68d7187f32e653f2, __obf_c37540905f7c483e)
}
