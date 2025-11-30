package __obf_f6c3dc60ecf20509

import (
	"encoding/json"
	"fmt"
	model "github.com/ArtisanHiram/go-pkg/runner/model"
	"os"
	"path/filepath"
)

func FilterByExtension(__obf_37a5b9d32ba2e789 []string, __obf_75d07abd3f2b728c string) []string {
	var __obf_1a31e9d44f3a463e []string

	for _, __obf_c344fa97859fd2b8 := range __obf_37a5b9d32ba2e789 {
		if filepath.Ext(__obf_c344fa97859fd2b8) == __obf_75d07abd3f2b728c {
			__obf_1a31e9d44f3a463e = append(__obf_1a31e9d44f3a463e, __obf_c344fa97859fd2b8)
		}
	}

	return __obf_1a31e9d44f3a463e
}

// Writes files to disk, returns list of absolute filepaths
func WriteFiles(__obf_59ace91ad9054b9e string, __obf_7105f3e2a697c584 string, __obf_37a5b9d32ba2e789 []*model.InMemoryFile) (__obf_4ce4ebae3765069f string, __obf_7f48400a7ee255fb []string, __obf_30c55ba125df4fa0 error) {
	__obf_4ce4ebae3765069f,
		// Create temp dir
		__obf_30c55ba125df4fa0 = os.MkdirTemp(__obf_59ace91ad9054b9e, //fmt.Sprintf("%s-*", lang))
		__obf_7105f3e2a697c584)
	if __obf_30c55ba125df4fa0 != nil {
		return
	}

	for _, __obf_c344fa97859fd2b8 := range __obf_37a5b9d32ba2e789 {
		__obf_7f48400a7ee255fb = append(__obf_7f48400a7ee255fb, __obf_c344fa97859fd2b8.Name)
		__obf_30c55ba125df4fa0 = __obf_7507eff15a02036b(__obf_4ce4ebae3765069f, __obf_c344fa97859fd2b8)
		if __obf_30c55ba125df4fa0 != nil {
			return
		}
	}
	return
}

// Writes a single file to disk
func __obf_7507eff15a02036b(__obf_4ce4ebae3765069f string, __obf_c344fa97859fd2b8 *model.InMemoryFile) error {
	__obf_30c55ba125df4fa0 := // Get absolute path to file inside basePath
		// Create all parent dirs
		os.MkdirAll(__obf_4ce4ebae3765069f, os.ModePerm)
	if __obf_30c55ba125df4fa0 != nil {
		return fmt.Errorf("MkdirAll failed: %s", __obf_30c55ba125df4fa0.Error())
	}
	__obf_30c55ba125df4fa0 = // Write file to disk
		os.WriteFile(filepath.Join(__obf_4ce4ebae3765069f, __obf_c344fa97859fd2b8.Name), []byte(__obf_c344fa97859fd2b8.Content), 0664)
	if __obf_30c55ba125df4fa0 != nil {
		return fmt.Errorf("WriteFile failed: %s", __obf_30c55ba125df4fa0.Error())
	}

	return nil
}

func PrintResult(__obf_ede12909df604f5e, __obf_f04a0d8902d07d2d string, __obf_30c55ba125df4fa0 error) {
	__obf_862e4ebca818c6bd := &model.Result{
		Stdout: __obf_ede12909df604f5e,
		Stderr: __obf_f04a0d8902d07d2d,
		Error: func(error) string {
			if __obf_30c55ba125df4fa0 != nil {
				return __obf_30c55ba125df4fa0.Error()
			}
			return ""
		}(__obf_30c55ba125df4fa0),
	}
	json.NewEncoder(os.Stdout).Encode(__obf_862e4ebca818c6bd)
}
