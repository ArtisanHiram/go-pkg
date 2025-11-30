package __obf_ccf5c05771773c19

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_377e68d538464b13 []string, __obf_1dd7c859ed12e85e string) []string {
	var __obf_e2c12246a5a70dda []string

	for _, __obf_e7d1904717a47b49 := range __obf_377e68d538464b13 {
		if filepath.Ext(__obf_e7d1904717a47b49) == __obf_1dd7c859ed12e85e {
			__obf_e2c12246a5a70dda = append(__obf_e2c12246a5a70dda, __obf_e7d1904717a47b49)
		}
	}

	return __obf_e2c12246a5a70dda
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_475e7d34254e6cb7 string, __obf_79457d3307c2e9dd string, __obf_377e68d538464b13 []*model.InMemoryFile) (__obf_31d17a62a54c4af9 string, __obf_43ac67e5f6d61618 []string, __obf_2bfc775e01ebe40d error) {
	__obf_31d17a62a54c4af9,
		// Create temp dir
		__obf_2bfc775e01ebe40d = os.MkdirTemp(__obf_475e7d34254e6cb7, //fmt.Sprintf("%s-*", lang))
		__obf_79457d3307c2e9dd)
	if __obf_2bfc775e01ebe40d != nil {
		return
	}

	for _, __obf_e7d1904717a47b49 := range __obf_377e68d538464b13 {
		__obf_43ac67e5f6d61618 = append(__obf_43ac67e5f6d61618, __obf_e7d1904717a47b49.Name)
		__obf_2bfc775e01ebe40d = __obf_ea71c059f90399de(__obf_31d17a62a54c4af9, __obf_e7d1904717a47b49)
		if __obf_2bfc775e01ebe40d != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_ea71c059f90399de(__obf_31d17a62a54c4af9 string, __obf_e7d1904717a47b49 *model.InMemoryFile) error {
	__obf_2bfc775e01ebe40d := // Get absolute path to file inside basePath
		// Create all parent dirs
		os.MkdirAll(__obf_31d17a62a54c4af9, os.ModePerm)
	if __obf_2bfc775e01ebe40d != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_2bfc775e01ebe40d.Error())
	}
	__obf_2bfc775e01ebe40d = // Write file to disk
		os.WriteFile(filepath.Join(__obf_31d17a62a54c4af9, __obf_e7d1904717a47b49.Name), []byte(__obf_e7d1904717a47b49.Content), 0664)
	if __obf_2bfc775e01ebe40d != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_2bfc775e01ebe40d.Error())
	}

	return nil
}

func PrintResult(__obf_ee0aef34779a44da, __obf_01a5d40825d73d93 string, __obf_2bfc775e01ebe40d error) {
	__obf_b52577d39fc6aa57 := &model.Result{
		Stdout: __obf_ee0aef34779a44da,
		Stderr: __obf_01a5d40825d73d93,
		Error: func(error) string {
			if __obf_2bfc775e01ebe40d != nil {
				return __obf_2bfc775e01ebe40d.Error()
			}
			return ""
		}(__obf_2bfc775e01ebe40d),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_b52577d39fc6aa57)
}
