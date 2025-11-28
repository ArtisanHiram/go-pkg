package __obf_a2c5231b0d2ff1f8

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_b57176d1a4eb350a []string, __obf_16eeba2174281878 string) []string {
	var __obf_35750a310cc94061 []string

	for _, __obf_d597b51a9be7c795 := range __obf_b57176d1a4eb350a {
		if filepath.Ext(__obf_d597b51a9be7c795) == __obf_16eeba2174281878 {
			__obf_35750a310cc94061 = append(__obf_35750a310cc94061, __obf_d597b51a9be7c795)
		}
	}

	return __obf_35750a310cc94061
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_1e606586bc06608f string, __obf_2ad8c77678f65c1d string, __obf_b57176d1a4eb350a []*model.InMemoryFile) (__obf_67e342ac9762cbcf string, __obf_b1078ff3d4a3ca9f []string, __obf_46954e3d09555328 error) {
	// Create temp dir
	__obf_67e342ac9762cbcf, __obf_46954e3d09555328 = os.MkdirTemp(__obf_1e606586bc06608f, __obf_2ad8c77678f65c1d) //fmt.Sprintf("%s-*", lang))
	if __obf_46954e3d09555328 != nil {
		return
	}

	for _, __obf_d597b51a9be7c795 := range __obf_b57176d1a4eb350a {
		__obf_b1078ff3d4a3ca9f = append(__obf_b1078ff3d4a3ca9f, __obf_d597b51a9be7c795.Name)
		__obf_46954e3d09555328 = __obf_b415e83eea38317d(__obf_67e342ac9762cbcf, __obf_d597b51a9be7c795)
		if __obf_46954e3d09555328 != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_b415e83eea38317d(__obf_67e342ac9762cbcf string, __obf_d597b51a9be7c795 *model.InMemoryFile) error {
	// Get absolute path to file inside basePath
	// Create all parent dirs
	__obf_46954e3d09555328 := os.MkdirAll(__obf_67e342ac9762cbcf, os.ModePerm)
	if __obf_46954e3d09555328 != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_46954e3d09555328.Error())
	}

	// Write file to disk
	__obf_46954e3d09555328 = os.WriteFile(filepath.Join(__obf_67e342ac9762cbcf, __obf_d597b51a9be7c795.Name), []byte(__obf_d597b51a9be7c795.Content), 0664)
	if __obf_46954e3d09555328 != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_46954e3d09555328.Error())
	}

	return nil
}

func PrintResult(__obf_1cbca869d143e21a, __obf_f5b2e0447dbd9c54 string, __obf_46954e3d09555328 error) {
	__obf_62fe912f5ed5a0bc := &model.Result{
		Stdout: __obf_1cbca869d143e21a,
		Stderr: __obf_f5b2e0447dbd9c54,
		Error: func(error) string {
			if __obf_46954e3d09555328 != nil {
				return __obf_46954e3d09555328.Error()
			}
			return ""
		}(__obf_46954e3d09555328),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_62fe912f5ed5a0bc)
}
