package __obf_1d74373d49f374d2

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_635abb04ff07917b []string, __obf_c6d8bd217ca16d32 string) []string {
	var __obf_ace257632b98c5f3 []string

	for _, __obf_a09eddc765544020 := range __obf_635abb04ff07917b {
		if filepath.Ext(__obf_a09eddc765544020) == __obf_c6d8bd217ca16d32 {
			__obf_ace257632b98c5f3 = append(__obf_ace257632b98c5f3, __obf_a09eddc765544020)
		}
	}

	return __obf_ace257632b98c5f3
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_a55f1adbc9cabaf6 string, __obf_0ba12a9eb3799358 string, __obf_635abb04ff07917b []*model.InMemoryFile) (__obf_df715018490607c1 string, __obf_6cf8265bef635bbf []string, __obf_400971760165a73e error) {
	__obf_df715018490607c1,
		// Create temp dir
		__obf_400971760165a73e = os.MkdirTemp(__obf_a55f1adbc9cabaf6, //fmt.Sprintf("%s-*", lang))
		__obf_0ba12a9eb3799358)
	if __obf_400971760165a73e != nil {
		return
	}

	for _, __obf_a09eddc765544020 := range __obf_635abb04ff07917b {
		__obf_6cf8265bef635bbf = append(__obf_6cf8265bef635bbf, __obf_a09eddc765544020.Name)
		__obf_400971760165a73e = __obf_0d67693d45d450fb(__obf_df715018490607c1, __obf_a09eddc765544020)
		if __obf_400971760165a73e != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_0d67693d45d450fb(__obf_df715018490607c1 string, __obf_a09eddc765544020 *model.InMemoryFile) error {
	__obf_400971760165a73e := // Get absolute path to file inside basePath
		// Create all parent dirs
		os.MkdirAll(__obf_df715018490607c1, os.ModePerm)
	if __obf_400971760165a73e != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_400971760165a73e.Error())
	}
	__obf_400971760165a73e = // Write file to disk
		os.WriteFile(filepath.Join(__obf_df715018490607c1, __obf_a09eddc765544020.Name), []byte(__obf_a09eddc765544020.Content), 0664)
	if __obf_400971760165a73e != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_400971760165a73e.Error())
	}

	return nil
}

func PrintResult(__obf_4fff1c5b00feff37, __obf_1e74ede6e2a56154 string, __obf_400971760165a73e error) {
	__obf_05349262ed4ce361 := &model.Result{
		Stdout: __obf_4fff1c5b00feff37,
		Stderr: __obf_1e74ede6e2a56154,
		Error: func(error) string {
			if __obf_400971760165a73e != nil {
				return __obf_400971760165a73e.Error()
			}
			return ""
		}(__obf_400971760165a73e),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_05349262ed4ce361)
}
