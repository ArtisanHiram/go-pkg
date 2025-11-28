package __obf_ab412e39d1ca7e25

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_04dbf20f09eca552 []string, __obf_cc6eebebe4f75911 string) []string {
	var __obf_9910a6d34da70390 []string

	for _, __obf_1c4ae3cbcf5144ba := range __obf_04dbf20f09eca552 {
		if filepath.Ext(__obf_1c4ae3cbcf5144ba) == __obf_cc6eebebe4f75911 {
			__obf_9910a6d34da70390 = append(__obf_9910a6d34da70390, __obf_1c4ae3cbcf5144ba)
		}
	}

	return __obf_9910a6d34da70390
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_040e24649c9c4b94 string, __obf_15f8b057156b65ad string, __obf_04dbf20f09eca552 []*model.InMemoryFile) (__obf_01bb8703ce75f555 string, __obf_885f81497f4bf213 []string, __obf_e5418f9487faf990 error) {
	// Create temp dir
	__obf_01bb8703ce75f555, __obf_e5418f9487faf990 = os.MkdirTemp(__obf_040e24649c9c4b94, __obf_15f8b057156b65ad) //fmt.Sprintf("%s-*", lang))
	if __obf_e5418f9487faf990 != nil {
		return
	}

	for _, __obf_1c4ae3cbcf5144ba := range __obf_04dbf20f09eca552 {
		__obf_885f81497f4bf213 = append(__obf_885f81497f4bf213, __obf_1c4ae3cbcf5144ba.Name)
		__obf_e5418f9487faf990 = __obf_ef98323d635faff1(__obf_01bb8703ce75f555, __obf_1c4ae3cbcf5144ba)
		if __obf_e5418f9487faf990 != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_ef98323d635faff1(__obf_01bb8703ce75f555 string, __obf_1c4ae3cbcf5144ba *model.InMemoryFile) error {
	// Get absolute path to file inside basePath
	// Create all parent dirs
	__obf_e5418f9487faf990 := os.MkdirAll(__obf_01bb8703ce75f555, os.ModePerm)
	if __obf_e5418f9487faf990 != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_e5418f9487faf990.Error())
	}

	// Write file to disk
	__obf_e5418f9487faf990 = os.WriteFile(filepath.Join(__obf_01bb8703ce75f555, __obf_1c4ae3cbcf5144ba.Name), []byte(__obf_1c4ae3cbcf5144ba.Content), 0664)
	if __obf_e5418f9487faf990 != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_e5418f9487faf990.Error())
	}

	return nil
}

func PrintResult(__obf_ca98ea1cd2d76371, __obf_9880b6037615b6cb string, __obf_e5418f9487faf990 error) {
	__obf_c9c18e3110085347 := &model.Result{
		Stdout: __obf_ca98ea1cd2d76371,
		Stderr: __obf_9880b6037615b6cb,
		Error: func(error) string {
			if __obf_e5418f9487faf990 != nil {
				return __obf_e5418f9487faf990.Error()
			}
			return ""
		}(__obf_e5418f9487faf990),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_c9c18e3110085347)
}
