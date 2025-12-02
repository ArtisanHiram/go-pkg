package __obf_eb5e805b9fc230e3

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_5f0e28df6b1ffd43 context.Context, __obf_2b7157422f48db91 string, __obf_4b197367e37598bd []byte) error
	Get(__obf_5f0e28df6b1ffd43 context.Context, __obf_2b7157422f48db91 string) ([]byte, error)
}

type DirCache struct {
	__obf_bd9b77e582496812 string
}

func NewDirCache(__obf_bd9b77e582496812 string) *DirCache {
	return &DirCache{__obf_bd9b77e582496812}
}

func (__obf_86deccad6befa8c8 *DirCache) __obf_53f36ee7488a0627(__obf_2b7157422f48db91 string) string {
	return filepath.Join(__obf_86deccad6befa8c8.__obf_bd9b77e582496812, __obf_2b7157422f48db91)
}

func (__obf_86deccad6befa8c8 *DirCache) Get(__obf_5f0e28df6b1ffd43 context.Context, __obf_2b7157422f48db91 string) ([]byte, error) {
	__obf_22d6c43ac51329e5, __obf_2d43222a56faebee := os.Open(__obf_86deccad6befa8c8.__obf_53f36ee7488a0627(__obf_2b7157422f48db91))
	if __obf_2d43222a56faebee != nil {
		if os.IsNotExist(__obf_2d43222a56faebee) {
			return nil, nil
		}
		return nil, __obf_2d43222a56faebee
	}
	defer __obf_22d6c43ac51329e5.Close()
	__obf_4872501882288e11 := new(bytes.Buffer)
	if _, __obf_2d43222a56faebee = io.Copy(__obf_4872501882288e11, __obf_22d6c43ac51329e5); __obf_2d43222a56faebee != nil {
		return nil, __obf_2d43222a56faebee
	}
	return __obf_4872501882288e11.Bytes(), nil
}

func (__obf_86deccad6befa8c8 *DirCache) Set(__obf_5f0e28df6b1ffd43 context.Context, __obf_2b7157422f48db91 string, __obf_4b197367e37598bd []byte) error {
	if __obf_2d43222a56faebee := os.MkdirAll(__obf_86deccad6befa8c8.__obf_bd9b77e582496812, 0777); __obf_2d43222a56faebee != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_86deccad6befa8c8.__obf_bd9b77e582496812, __obf_2d43222a56faebee)
		return __obf_2d43222a56faebee
	}
	__obf_d2843fee5f089c2d, __obf_2d43222a56faebee := os.CreateTemp(__obf_86deccad6befa8c8.__obf_bd9b77e582496812, __obf_2b7157422f48db91+".*")
	if __obf_2d43222a56faebee != nil {
		log.Printf("Could not create cache file %v: %v", __obf_d2843fee5f089c2d, __obf_2d43222a56faebee)
		return __obf_2d43222a56faebee
	}
	if _, __obf_2d43222a56faebee := io.Copy(__obf_d2843fee5f089c2d, bytes.NewReader(__obf_4b197367e37598bd)); __obf_2d43222a56faebee != nil {
		log.Printf("Could not write cache file %v: %v", __obf_d2843fee5f089c2d, __obf_2d43222a56faebee)
		__obf_d2843fee5f089c2d.
			Close()
		os.Remove(__obf_d2843fee5f089c2d.Name())
		return __obf_2d43222a56faebee
	}
	if __obf_2d43222a56faebee = __obf_d2843fee5f089c2d.Close(); __obf_2d43222a56faebee != nil {
		log.Printf("Could not close cache file %v: %v", __obf_d2843fee5f089c2d, __obf_2d43222a56faebee)
		os.Remove(__obf_d2843fee5f089c2d.Name())
		return __obf_2d43222a56faebee
	}
	if __obf_2d43222a56faebee = os.Rename(__obf_d2843fee5f089c2d.Name(), __obf_86deccad6befa8c8.__obf_53f36ee7488a0627(__obf_2b7157422f48db91)); __obf_2d43222a56faebee != nil {
		log.Printf("Could not move cache file %v: %v", __obf_d2843fee5f089c2d, __obf_2d43222a56faebee)
		os.Remove(__obf_d2843fee5f089c2d.Name())
		return __obf_2d43222a56faebee
	}
	return nil
}
