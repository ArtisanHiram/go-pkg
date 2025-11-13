package __obf_88600f275b498cc3

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_f298e519fb172c5e []string, __obf_1833864396e85913 string) []string {
	var __obf_7379f399b3a3643c []string

	for _, __obf_efcb6cce75accfd5 := range __obf_f298e519fb172c5e {
		if filepath.Ext(__obf_efcb6cce75accfd5) == __obf_1833864396e85913 {
			__obf_7379f399b3a3643c = append(__obf_7379f399b3a3643c, __obf_efcb6cce75accfd5)
		}
	}

	return __obf_7379f399b3a3643c
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_417bb20aa533c678 string, __obf_c9c8ed081fd47786 string, __obf_f298e519fb172c5e []*model.InMemoryFile) (__obf_8a19d4b0d82f707f string, __obf_d1f9f1122872ce17 []string, __obf_d4738a11896b5dbb error) {
	// Create temp dir
	__obf_8a19d4b0d82f707f, __obf_d4738a11896b5dbb = os.MkdirTemp(__obf_417bb20aa533c678, __obf_c9c8ed081fd47786) //fmt.Sprintf("%s-*", lang))
	if __obf_d4738a11896b5dbb != nil {
		return
	}

	for _, __obf_efcb6cce75accfd5 := range __obf_f298e519fb172c5e {
		__obf_d1f9f1122872ce17 = append(__obf_d1f9f1122872ce17, __obf_efcb6cce75accfd5.Name)
		__obf_d4738a11896b5dbb = __obf_f7143ae905cab244(__obf_8a19d4b0d82f707f, __obf_efcb6cce75accfd5)
		if __obf_d4738a11896b5dbb != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_f7143ae905cab244(__obf_8a19d4b0d82f707f string, __obf_efcb6cce75accfd5 *model.InMemoryFile) error {
	// Get absolute path to file inside basePath
	// Create all parent dirs
	__obf_d4738a11896b5dbb := os.MkdirAll(__obf_8a19d4b0d82f707f, os.ModePerm)
	if __obf_d4738a11896b5dbb != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_d4738a11896b5dbb.Error())
	}

	// Write file to disk
	__obf_d4738a11896b5dbb = os.WriteFile(filepath.Join(__obf_8a19d4b0d82f707f, __obf_efcb6cce75accfd5.Name), []byte(__obf_efcb6cce75accfd5.Content), 0664)
	if __obf_d4738a11896b5dbb != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_d4738a11896b5dbb.Error())
	}

	return nil
}

func PrintResult(__obf_52e98ae4a2a3e15c, __obf_ac3166771ec28436 string, __obf_d4738a11896b5dbb error) {
	__obf_5c33dbce17cea990 := &model.Result{
		Stdout: __obf_52e98ae4a2a3e15c,
		Stderr: __obf_ac3166771ec28436,
		Error: func(error) string {
			if __obf_d4738a11896b5dbb != nil {
				return __obf_d4738a11896b5dbb.Error()
			}
			return ""
		}(__obf_d4738a11896b5dbb),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_5c33dbce17cea990)
}
