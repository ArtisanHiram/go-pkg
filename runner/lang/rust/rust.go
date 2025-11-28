package __obf_fd3c51f05fc3c0d7

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_a0fb63ff744999bc string, __obf_7d1b7aa4486c58d0 string, __obf_912c143b9e5f48b7 []string) (string, string, error) {
	__obf_e4e76e1ea98e4759 := "main.out"
	__obf_2f181d3083dc7fd9, __obf_d1fae138e9094ce2, __obf_835dfc37b3b0f100 := cmd.Run(__obf_a0fb63ff744999bc, "rustc", "-o", __obf_e4e76e1ea98e4759, __obf_912c143b9e5f48b7[0])
	if __obf_835dfc37b3b0f100 != nil {
		return __obf_2f181d3083dc7fd9, __obf_d1fae138e9094ce2, __obf_835dfc37b3b0f100
	}

	return cmd.RunStdin(__obf_a0fb63ff744999bc, __obf_7d1b7aa4486c58d0, __obf_e4e76e1ea98e4759)
}
