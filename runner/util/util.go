package __obf_ec9d7914a439392a

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_90abc88922b5ede6 []string, __obf_d3ad5ed83f1f130d string) []string {
	var __obf_6e69d47e5313fbb5 []string

	for _, __obf_acee41a77c558128 := range __obf_90abc88922b5ede6 {
		if filepath.Ext(__obf_acee41a77c558128) == __obf_d3ad5ed83f1f130d {
			__obf_6e69d47e5313fbb5 = append(__obf_6e69d47e5313fbb5, __obf_acee41a77c558128)
		}
	}

	return __obf_6e69d47e5313fbb5
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_b61fccab5a760df3 string, __obf_ddd60739524d343e string, __obf_90abc88922b5ede6 []*model.InMemoryFile) (__obf_70c05b53bc24e97a string, __obf_355ee014d91538c0 []string, __obf_e73c5c3dfebd52fb error) {
	// Create temp dir
	__obf_70c05b53bc24e97a, __obf_e73c5c3dfebd52fb = os.MkdirTemp(__obf_b61fccab5a760df3, __obf_ddd60739524d343e) //fmt.Sprintf("%s-*", lang))
	if __obf_e73c5c3dfebd52fb != nil {
		return
	}

	for _, __obf_acee41a77c558128 := range __obf_90abc88922b5ede6 {
		__obf_355ee014d91538c0 = append(__obf_355ee014d91538c0, __obf_acee41a77c558128.Name)
		__obf_e73c5c3dfebd52fb = __obf_ccd2c0048e8beac7(__obf_70c05b53bc24e97a, __obf_acee41a77c558128)
		if __obf_e73c5c3dfebd52fb != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_ccd2c0048e8beac7(__obf_70c05b53bc24e97a string, __obf_acee41a77c558128 *model.InMemoryFile) error {
	// Get absolute path to file inside basePath
	// Create all parent dirs
	__obf_e73c5c3dfebd52fb := os.MkdirAll(__obf_70c05b53bc24e97a, os.ModePerm)
	if __obf_e73c5c3dfebd52fb != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_e73c5c3dfebd52fb.Error())
	}

	// Write file to disk
	__obf_e73c5c3dfebd52fb = os.WriteFile(filepath.Join(__obf_70c05b53bc24e97a, __obf_acee41a77c558128.Name), []byte(__obf_acee41a77c558128.Content), 0664)
	if __obf_e73c5c3dfebd52fb != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_e73c5c3dfebd52fb.Error())
	}

	return nil
}

func PrintResult(__obf_34bd87285db2e858, __obf_b0cc1d6f82789b37 string, __obf_e73c5c3dfebd52fb error) {
	__obf_3559bac86f838a58 := &model.Result{
		Stdout: __obf_34bd87285db2e858,
		Stderr: __obf_b0cc1d6f82789b37,
		Error: func(error) string {
			if __obf_e73c5c3dfebd52fb != nil {
				return __obf_e73c5c3dfebd52fb.Error()
			}
			return ""
		}(__obf_e73c5c3dfebd52fb),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_3559bac86f838a58)
}
