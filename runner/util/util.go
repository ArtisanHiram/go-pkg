package __obf_fd89f4a0b8c90860

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_2a825626acb897de []string, __obf_ff1f9e4092862b19 string) []string {
	var __obf_28e512949f26607a []string

	for _, __obf_2c1e16a5469323a3 := range __obf_2a825626acb897de {
		if filepath.Ext(__obf_2c1e16a5469323a3) == __obf_ff1f9e4092862b19 {
			__obf_28e512949f26607a = append(__obf_28e512949f26607a, __obf_2c1e16a5469323a3)
		}
	}

	return __obf_28e512949f26607a
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_ff7d3950ad4d9c89 string, __obf_081df6aafeead974 string, __obf_2a825626acb897de []*model.InMemoryFile) (__obf_b183147032963fd8 string, __obf_ee5b6ffa374a19af []string, __obf_bf0573411cb4a41c error) {
	__obf_b183147032963fd8,
		// Create temp dir
		__obf_bf0573411cb4a41c = os.MkdirTemp(__obf_ff7d3950ad4d9c89, //fmt.Sprintf("%s-*", lang))
		__obf_081df6aafeead974)
	if __obf_bf0573411cb4a41c != nil {
		return
	}

	for _, __obf_2c1e16a5469323a3 := range __obf_2a825626acb897de {
		__obf_ee5b6ffa374a19af = append(__obf_ee5b6ffa374a19af, __obf_2c1e16a5469323a3.Name)
		__obf_bf0573411cb4a41c = __obf_dc8c8e7d31251936(__obf_b183147032963fd8, __obf_2c1e16a5469323a3)
		if __obf_bf0573411cb4a41c != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_dc8c8e7d31251936(__obf_b183147032963fd8 string, __obf_2c1e16a5469323a3 *model.InMemoryFile) error {
	__obf_bf0573411cb4a41c := // Get absolute path to file inside basePath
		// Create all parent dirs
		os.MkdirAll(__obf_b183147032963fd8, os.ModePerm)
	if __obf_bf0573411cb4a41c != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_bf0573411cb4a41c.Error())
	}
	__obf_bf0573411cb4a41c = // Write file to disk
		os.WriteFile(filepath.Join(__obf_b183147032963fd8, __obf_2c1e16a5469323a3.Name), []byte(__obf_2c1e16a5469323a3.Content), 0664)
	if __obf_bf0573411cb4a41c != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_bf0573411cb4a41c.Error())
	}

	return nil
}

func PrintResult(__obf_ee75dd255b8c0fd3, __obf_da32b299551f2a62 string, __obf_bf0573411cb4a41c error) {
	__obf_c6c76128eadfeece := &model.Result{
		Stdout: __obf_ee75dd255b8c0fd3,
		Stderr: __obf_da32b299551f2a62,
		Error: func(error) string {
			if __obf_bf0573411cb4a41c != nil {
				return __obf_bf0573411cb4a41c.Error()
			}
			return ""
		}(__obf_bf0573411cb4a41c),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_c6c76128eadfeece)
}
