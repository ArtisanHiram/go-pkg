package __obf_f4378aa3a7aae5ec

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_54d8c3ea5b9fca26 []string, __obf_dfd53aaecff9cb90 string) []string {
	var __obf_9c50e9b72247d02a []string

	for _, __obf_41c9ced678465d3c := range __obf_54d8c3ea5b9fca26 {
		if filepath.Ext(__obf_41c9ced678465d3c) == __obf_dfd53aaecff9cb90 {
			__obf_9c50e9b72247d02a = append(__obf_9c50e9b72247d02a, __obf_41c9ced678465d3c)
		}
	}

	return __obf_9c50e9b72247d02a
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_925194751c7c8dda string, __obf_29ab4592cb1642f9 string, __obf_54d8c3ea5b9fca26 []*model.InMemoryFile) (__obf_a42a989679fb6389 string, __obf_63be94827d68e088 []string, __obf_0e82b985c01a0b75 error) {
	// Create temp dir
	__obf_a42a989679fb6389, __obf_0e82b985c01a0b75 = os.MkdirTemp(__obf_925194751c7c8dda, __obf_29ab4592cb1642f9) //fmt.Sprintf("%s-*", lang))
	if __obf_0e82b985c01a0b75 != nil {
		return
	}

	for _, __obf_41c9ced678465d3c := range __obf_54d8c3ea5b9fca26 {
		__obf_63be94827d68e088 = append(__obf_63be94827d68e088, __obf_41c9ced678465d3c.Name)
		__obf_0e82b985c01a0b75 = __obf_2fb4bbef05f211e6(__obf_a42a989679fb6389, __obf_41c9ced678465d3c)
		if __obf_0e82b985c01a0b75 != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_2fb4bbef05f211e6(__obf_a42a989679fb6389 string, __obf_41c9ced678465d3c *model.InMemoryFile) error {
	// Get absolute path to file inside basePath
	// Create all parent dirs
	__obf_0e82b985c01a0b75 := os.MkdirAll(__obf_a42a989679fb6389, os.ModePerm)
	if __obf_0e82b985c01a0b75 != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_0e82b985c01a0b75.Error())
	}

	// Write file to disk
	__obf_0e82b985c01a0b75 = os.WriteFile(filepath.Join(__obf_a42a989679fb6389, __obf_41c9ced678465d3c.Name), []byte(__obf_41c9ced678465d3c.Content), 0664)
	if __obf_0e82b985c01a0b75 != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_0e82b985c01a0b75.Error())
	}

	return nil
}

func PrintResult(__obf_f39e3382682e14de, __obf_d9f20533d5c04984 string, __obf_0e82b985c01a0b75 error) {
	__obf_9a1387515840e190 := &model.Result{
		Stdout: __obf_f39e3382682e14de,
		Stderr: __obf_d9f20533d5c04984,
		Error: func(error) string {
			if __obf_0e82b985c01a0b75 != nil {
				return __obf_0e82b985c01a0b75.Error()
			}
			return ""
		}(__obf_0e82b985c01a0b75),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_9a1387515840e190)
}
