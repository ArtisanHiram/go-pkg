package __obf_b28324f38df50634

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_c2c8a52eb04b2873 context.Context, __obf_272d13eb9342c812 string, __obf_2466d31be5bb5404 []byte) error
	Get(__obf_c2c8a52eb04b2873 context.Context, __obf_272d13eb9342c812 string) ([]byte, error)
}

type DirCache struct {
	__obf_ef284eab1db4670a string
}

func NewDirCache(__obf_ef284eab1db4670a string) *DirCache {
	return &DirCache{__obf_ef284eab1db4670a}
}

func (__obf_732b021ced848703 *DirCache) __obf_d7b6f7562156aea5(__obf_272d13eb9342c812 string) string {
	return filepath.Join(__obf_732b021ced848703.__obf_ef284eab1db4670a, __obf_272d13eb9342c812)
}

func (__obf_732b021ced848703 *DirCache) Get(__obf_c2c8a52eb04b2873 context.Context, __obf_272d13eb9342c812 string) ([]byte, error) {
	__obf_a2e9fb1ef0c6fd12, __obf_a4e073529832bad2 := os.Open(__obf_732b021ced848703.__obf_d7b6f7562156aea5(__obf_272d13eb9342c812))
	if __obf_a4e073529832bad2 != nil {
		if os.IsNotExist(__obf_a4e073529832bad2) {
			return nil, nil
		}
		return nil, __obf_a4e073529832bad2
	}
	defer __obf_a2e9fb1ef0c6fd12.Close()
	__obf_a716135a97d66cdf := new(bytes.Buffer)
	if _, __obf_a4e073529832bad2 = io.Copy(__obf_a716135a97d66cdf, __obf_a2e9fb1ef0c6fd12); __obf_a4e073529832bad2 != nil {
		return nil, __obf_a4e073529832bad2
	}
	return __obf_a716135a97d66cdf.Bytes(), nil
}

func (__obf_732b021ced848703 *DirCache) Set(__obf_c2c8a52eb04b2873 context.Context, __obf_272d13eb9342c812 string, __obf_2466d31be5bb5404 []byte) error {
	if __obf_a4e073529832bad2 := os.MkdirAll(__obf_732b021ced848703.__obf_ef284eab1db4670a, 0777); __obf_a4e073529832bad2 != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_732b021ced848703.__obf_ef284eab1db4670a, __obf_a4e073529832bad2)
		return __obf_a4e073529832bad2
	}
	__obf_780097e2e1ff1b0d, __obf_a4e073529832bad2 := os.CreateTemp(__obf_732b021ced848703.__obf_ef284eab1db4670a, __obf_272d13eb9342c812+".*")
	if __obf_a4e073529832bad2 != nil {
		log.Printf("Could not create cache file %v: %v", __obf_780097e2e1ff1b0d, __obf_a4e073529832bad2)
		return __obf_a4e073529832bad2
	}
	if _, __obf_a4e073529832bad2 := io.Copy(__obf_780097e2e1ff1b0d, bytes.NewReader(__obf_2466d31be5bb5404)); __obf_a4e073529832bad2 != nil {
		log.Printf("Could not write cache file %v: %v", __obf_780097e2e1ff1b0d, __obf_a4e073529832bad2)
		__obf_780097e2e1ff1b0d.Close()
		os.Remove(__obf_780097e2e1ff1b0d.Name())
		return __obf_a4e073529832bad2
	}
	if __obf_a4e073529832bad2 = __obf_780097e2e1ff1b0d.Close(); __obf_a4e073529832bad2 != nil {
		log.Printf("Could not close cache file %v: %v", __obf_780097e2e1ff1b0d, __obf_a4e073529832bad2)
		os.Remove(__obf_780097e2e1ff1b0d.Name())
		return __obf_a4e073529832bad2
	}
	if __obf_a4e073529832bad2 = os.Rename(__obf_780097e2e1ff1b0d.Name(), __obf_732b021ced848703.__obf_d7b6f7562156aea5(__obf_272d13eb9342c812)); __obf_a4e073529832bad2 != nil {
		log.Printf("Could not move cache file %v: %v", __obf_780097e2e1ff1b0d, __obf_a4e073529832bad2)
		os.Remove(__obf_780097e2e1ff1b0d.Name())
		return __obf_a4e073529832bad2
	}
	return nil
}
