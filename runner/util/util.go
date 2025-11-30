package __obf_f3538974c96d9692

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_ee5b5991215ba7e1 []string, __obf_441e17b42f51fd59 string) []string {
	var __obf_fb017e972bd5cada []string

	for _, __obf_93d4a19d9f583cd4 := range __obf_ee5b5991215ba7e1 {
		if filepath.Ext(__obf_93d4a19d9f583cd4) == __obf_441e17b42f51fd59 {
			__obf_fb017e972bd5cada = append(__obf_fb017e972bd5cada, __obf_93d4a19d9f583cd4)
		}
	}

	return __obf_fb017e972bd5cada
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_f23ab7ecaecd38e8 string, __obf_d325d225cd585054 string, __obf_ee5b5991215ba7e1 []*model.InMemoryFile) (__obf_cf38f9f50b7aa6ec string, __obf_0843776f6650463b []string, __obf_944c4239074e324c error) {
	__obf_cf38f9f50b7aa6ec,
		// Create temp dir
		__obf_944c4239074e324c = os.MkdirTemp(__obf_f23ab7ecaecd38e8, //fmt.Sprintf("%s-*", lang))
		__obf_d325d225cd585054)
	if __obf_944c4239074e324c != nil {
		return
	}

	for _, __obf_93d4a19d9f583cd4 := range __obf_ee5b5991215ba7e1 {
		__obf_0843776f6650463b = append(__obf_0843776f6650463b, __obf_93d4a19d9f583cd4.Name)
		__obf_944c4239074e324c = __obf_4b5c1a491e30ad2b(__obf_cf38f9f50b7aa6ec, __obf_93d4a19d9f583cd4)
		if __obf_944c4239074e324c != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_4b5c1a491e30ad2b(__obf_cf38f9f50b7aa6ec string, __obf_93d4a19d9f583cd4 *model.InMemoryFile) error {
	__obf_944c4239074e324c := // Get absolute path to file inside basePath
		// Create all parent dirs
		os.MkdirAll(__obf_cf38f9f50b7aa6ec, os.ModePerm)
	if __obf_944c4239074e324c != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_944c4239074e324c.Error())
	}
	__obf_944c4239074e324c = // Write file to disk
		os.WriteFile(filepath.Join(__obf_cf38f9f50b7aa6ec, __obf_93d4a19d9f583cd4.Name), []byte(__obf_93d4a19d9f583cd4.Content), 0664)
	if __obf_944c4239074e324c != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_944c4239074e324c.Error())
	}

	return nil
}

func PrintResult(__obf_7cc308f6daa31a7f, __obf_103bb3a86c15b7e3 string, __obf_944c4239074e324c error) {
	__obf_94e65fccac9eeff4 := &model.Result{
		Stdout: __obf_7cc308f6daa31a7f,
		Stderr: __obf_103bb3a86c15b7e3,
		Error: func(error) string {
			if __obf_944c4239074e324c != nil {
				return __obf_944c4239074e324c.Error()
			}
			return ""
		}(__obf_944c4239074e324c),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_94e65fccac9eeff4)
}
