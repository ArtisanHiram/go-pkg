package __obf_a4d8029cfe468374

import (
	cmd "github.com/ArtisanHiram/go-pkg/runner/cmd"
)

func Run(__obf_12f24ef3bd59c7da string, __obf_6037c7acfd4fc66a string, __obf_61e83bdae8dc9cb3 []string) (string, string, error) {
	return cmd.RunStdin(__obf_12f24ef3bd59c7da, __obf_6037c7acfd4fc66a, "python3", __obf_61e83bdae8dc9cb3[0])
}
