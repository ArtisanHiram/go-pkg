package __obf_f60326fd90eb13d9

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_b580c2427f47a43f context.Context, __obf_6a74cec141e92747 string, __obf_df059c9a1d1e8ec1 []byte) error
	Get(__obf_b580c2427f47a43f context.Context, __obf_6a74cec141e92747 string) ([]byte, error)
}

type DirCache struct {
	__obf_97f00cf8dd461989 string
}

func NewDirCache(__obf_97f00cf8dd461989 string) *DirCache {
	return &DirCache{__obf_97f00cf8dd461989}
}

func (__obf_6ff7f1828f3842ea *DirCache) __obf_9803e1b7613a2e9b(__obf_6a74cec141e92747 string) string {
	return filepath.Join(__obf_6ff7f1828f3842ea.__obf_97f00cf8dd461989, __obf_6a74cec141e92747)
}

func (__obf_6ff7f1828f3842ea *DirCache) Get(__obf_b580c2427f47a43f context.Context, __obf_6a74cec141e92747 string) ([]byte, error) {
	__obf_2c7a65b2dac56e8a, __obf_083e5685f93ba07c := os.Open(__obf_6ff7f1828f3842ea.__obf_9803e1b7613a2e9b(__obf_6a74cec141e92747))
	if __obf_083e5685f93ba07c != nil {
		if os.IsNotExist(__obf_083e5685f93ba07c) {
			return nil, nil
		}
		return nil, __obf_083e5685f93ba07c
	}
	defer __obf_2c7a65b2dac56e8a.Close()
	__obf_a2286bef14c440b9 := new(bytes.Buffer)
	if _, __obf_083e5685f93ba07c = io.Copy(__obf_a2286bef14c440b9, __obf_2c7a65b2dac56e8a); __obf_083e5685f93ba07c != nil {
		return nil, __obf_083e5685f93ba07c
	}
	return __obf_a2286bef14c440b9.Bytes(), nil
}

func (__obf_6ff7f1828f3842ea *DirCache) Set(__obf_b580c2427f47a43f context.Context, __obf_6a74cec141e92747 string, __obf_df059c9a1d1e8ec1 []byte) error {
	if __obf_083e5685f93ba07c := os.MkdirAll(__obf_6ff7f1828f3842ea.__obf_97f00cf8dd461989, 0777); __obf_083e5685f93ba07c != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_6ff7f1828f3842ea.__obf_97f00cf8dd461989, __obf_083e5685f93ba07c)
		return __obf_083e5685f93ba07c
	}
	__obf_4ae91c6dc2f1a9e9, __obf_083e5685f93ba07c := os.CreateTemp(__obf_6ff7f1828f3842ea.__obf_97f00cf8dd461989, __obf_6a74cec141e92747+".*")
	if __obf_083e5685f93ba07c != nil {
		log.Printf("Could not create cache file %v: %v", __obf_4ae91c6dc2f1a9e9, __obf_083e5685f93ba07c)
		return __obf_083e5685f93ba07c
	}
	if _, __obf_083e5685f93ba07c := io.Copy(__obf_4ae91c6dc2f1a9e9, bytes.NewReader(__obf_df059c9a1d1e8ec1)); __obf_083e5685f93ba07c != nil {
		log.Printf("Could not write cache file %v: %v", __obf_4ae91c6dc2f1a9e9, __obf_083e5685f93ba07c)
		__obf_4ae91c6dc2f1a9e9.
			Close()
		os.Remove(__obf_4ae91c6dc2f1a9e9.Name())
		return __obf_083e5685f93ba07c
	}
	if __obf_083e5685f93ba07c = __obf_4ae91c6dc2f1a9e9.Close(); __obf_083e5685f93ba07c != nil {
		log.Printf("Could not close cache file %v: %v", __obf_4ae91c6dc2f1a9e9, __obf_083e5685f93ba07c)
		os.Remove(__obf_4ae91c6dc2f1a9e9.Name())
		return __obf_083e5685f93ba07c
	}
	if __obf_083e5685f93ba07c = os.Rename(__obf_4ae91c6dc2f1a9e9.Name(), __obf_6ff7f1828f3842ea.__obf_9803e1b7613a2e9b(__obf_6a74cec141e92747)); __obf_083e5685f93ba07c != nil {
		log.Printf("Could not move cache file %v: %v", __obf_4ae91c6dc2f1a9e9, __obf_083e5685f93ba07c)
		os.Remove(__obf_4ae91c6dc2f1a9e9.Name())
		return __obf_083e5685f93ba07c
	}
	return nil
}
