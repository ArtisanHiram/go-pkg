package __obf_fdd00a584ec28b9f

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_c66827c7546c516c string, __obf_7eaf35a11f231116 string, __obf_3b1c285f664273c9 []string) (string, string, error) {
	__obf_5395d44bc4ab54fd := "main.out"
	__obf_7c546b1509b82226, __obf_bbcd1fc3ee8a7a57, __obf_e46b3aa50fc681bd := cmd.Run(__obf_c66827c7546c516c, "rustc", "-o", __obf_5395d44bc4ab54fd, __obf_3b1c285f664273c9[0])
	if __obf_e46b3aa50fc681bd != nil {
		return __obf_7c546b1509b82226, __obf_bbcd1fc3ee8a7a57, __obf_e46b3aa50fc681bd
	}

	return cmd.RunStdin(__obf_c66827c7546c516c, __obf_7eaf35a11f231116, __obf_5395d44bc4ab54fd)
}
