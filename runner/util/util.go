package __obf_521b81af4f3e0fb6

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_40e39bdcaf1cf9ac []string, __obf_163ccadd52015306 string) []string {
	var __obf_1652b919c678ab16 []string

	for _, __obf_48b99c344dac76ae := range __obf_40e39bdcaf1cf9ac {
		if filepath.Ext(__obf_48b99c344dac76ae) == __obf_163ccadd52015306 {
			__obf_1652b919c678ab16 = append(__obf_1652b919c678ab16, __obf_48b99c344dac76ae)
		}
	}

	return __obf_1652b919c678ab16
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_734f4ae5f10e1373 string, __obf_94eae137dee17b2a string, __obf_40e39bdcaf1cf9ac []*model.InMemoryFile) (__obf_72fa348822559ff6 string, __obf_3dfdf78a9e849890 []string, __obf_2dcc3585c0e25eac error) {
	__obf_72fa348822559ff6,
		// Create temp dir
		__obf_2dcc3585c0e25eac = os.MkdirTemp(__obf_734f4ae5f10e1373, //fmt.Sprintf("%s-*", lang))
		__obf_94eae137dee17b2a)
	if __obf_2dcc3585c0e25eac != nil {
		return
	}

	for _, __obf_48b99c344dac76ae := range __obf_40e39bdcaf1cf9ac {
		__obf_3dfdf78a9e849890 = append(__obf_3dfdf78a9e849890, __obf_48b99c344dac76ae.Name)
		__obf_2dcc3585c0e25eac = __obf_b2dd3a380acc3ebf(__obf_72fa348822559ff6, __obf_48b99c344dac76ae)
		if __obf_2dcc3585c0e25eac != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_b2dd3a380acc3ebf(__obf_72fa348822559ff6 string, __obf_48b99c344dac76ae *model.InMemoryFile) error {
	__obf_2dcc3585c0e25eac := // Get absolute path to file inside basePath
		// Create all parent dirs
		os.MkdirAll(__obf_72fa348822559ff6, os.ModePerm)
	if __obf_2dcc3585c0e25eac != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_2dcc3585c0e25eac.Error())
	}
	__obf_2dcc3585c0e25eac = // Write file to disk
		os.WriteFile(filepath.Join(__obf_72fa348822559ff6, __obf_48b99c344dac76ae.Name), []byte(__obf_48b99c344dac76ae.Content), 0664)
	if __obf_2dcc3585c0e25eac != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_2dcc3585c0e25eac.Error())
	}

	return nil
}

func PrintResult(__obf_4e792eebe4d12263, __obf_7f8667b081261e3b string, __obf_2dcc3585c0e25eac error) {
	__obf_e016c5c63390dd8f := &model.Result{
		Stdout: __obf_4e792eebe4d12263,
		Stderr: __obf_7f8667b081261e3b,
		Error: func(error) string {
			if __obf_2dcc3585c0e25eac != nil {
				return __obf_2dcc3585c0e25eac.Error()
			}
			return ""
		}(__obf_2dcc3585c0e25eac),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_e016c5c63390dd8f)
}
