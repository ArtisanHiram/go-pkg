package __obf_0db3567988dbee94

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
	"path/filepath"
)

func Run(__obf_0b2adfc2f41fd172, __obf_99b7747b68041347 string, __obf_1cb27b48bc2a845c []string) (string, string, error) {
	__obf_b95b345249378829 := "main.out"

	__obf_ee11f4b6b5e3373a := append([]string{"clang++", "-std=c++11", "-o", __obf_b95b345249378829}, __obf_1cb27b48bc2a845c...)
	__obf_adfc1aac49122c99, __obf_3e76f5da14d82fa3, __obf_1ff3ef539cddaaac := cmd.Run(__obf_0b2adfc2f41fd172, __obf_ee11f4b6b5e3373a...)
	if __obf_1ff3ef539cddaaac != nil {
		return __obf_adfc1aac49122c99, __obf_3e76f5da14d82fa3, __obf_1ff3ef539cddaaac
	}

	__obf_0af0c3c964aad7d9 := filepath.Join(__obf_0b2adfc2f41fd172, __obf_b95b345249378829)
	return cmd.RunStdin(__obf_0b2adfc2f41fd172, __obf_99b7747b68041347, __obf_0af0c3c964aad7d9)
}
