package __obf_3fcb5f297f4fa0ce

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_607acdba0f8bfb59, __obf_5439fd2b8d11cec3 string, __obf_600a99b324bb9098 []string) (string, string, error) {
	__obf_46c345598d03ac08 := append([]string{"groovy"}, __obf_600a99b324bb9098...)
	return cmd.RunStdin(__obf_607acdba0f8bfb59, __obf_5439fd2b8d11cec3, __obf_46c345598d03ac08...)
}
