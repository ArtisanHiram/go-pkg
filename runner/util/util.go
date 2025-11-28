package __obf_4cdd697466e04432

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_d7d3c6b4759d6682 []string, __obf_b0d126dc161fb556 string) []string {
	var __obf_a36b8875046b6497 []string

	for _, __obf_4309320fa51cac7d := range __obf_d7d3c6b4759d6682 {
		if filepath.Ext(__obf_4309320fa51cac7d) == __obf_b0d126dc161fb556 {
			__obf_a36b8875046b6497 = append(__obf_a36b8875046b6497, __obf_4309320fa51cac7d)
		}
	}

	return __obf_a36b8875046b6497
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_4b9ab35f40a026c6 string, __obf_59d903e403627430 string, __obf_d7d3c6b4759d6682 []*model.InMemoryFile) (__obf_fe84a69df071e4f5 string, __obf_2fa9dbbf0bb9b84b []string, __obf_a41e981e29b848f0 error) {
	// Create temp dir
	__obf_fe84a69df071e4f5, __obf_a41e981e29b848f0 = os.MkdirTemp(__obf_4b9ab35f40a026c6, __obf_59d903e403627430) //fmt.Sprintf("%s-*", lang))
	if __obf_a41e981e29b848f0 != nil {
		return
	}

	for _, __obf_4309320fa51cac7d := range __obf_d7d3c6b4759d6682 {
		__obf_2fa9dbbf0bb9b84b = append(__obf_2fa9dbbf0bb9b84b, __obf_4309320fa51cac7d.Name)
		__obf_a41e981e29b848f0 = __obf_343032dea6dd926b(__obf_fe84a69df071e4f5, __obf_4309320fa51cac7d)
		if __obf_a41e981e29b848f0 != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_343032dea6dd926b(__obf_fe84a69df071e4f5 string, __obf_4309320fa51cac7d *model.InMemoryFile) error {
	// Get absolute path to file inside basePath
	// Create all parent dirs
	__obf_a41e981e29b848f0 := os.MkdirAll(__obf_fe84a69df071e4f5, os.ModePerm)
	if __obf_a41e981e29b848f0 != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_a41e981e29b848f0.Error())
	}

	// Write file to disk
	__obf_a41e981e29b848f0 = os.WriteFile(filepath.Join(__obf_fe84a69df071e4f5, __obf_4309320fa51cac7d.Name), []byte(__obf_4309320fa51cac7d.Content), 0664)
	if __obf_a41e981e29b848f0 != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_a41e981e29b848f0.Error())
	}

	return nil
}

func PrintResult(__obf_47a9095864081556, __obf_ad79a4aacf1da9e5 string, __obf_a41e981e29b848f0 error) {
	__obf_be5022f728b96ed3 := &model.Result{
		Stdout: __obf_47a9095864081556,
		Stderr: __obf_ad79a4aacf1da9e5,
		Error: func(error) string {
			if __obf_a41e981e29b848f0 != nil {
				return __obf_a41e981e29b848f0.Error()
			}
			return ""
		}(__obf_a41e981e29b848f0),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_be5022f728b96ed3)
}
