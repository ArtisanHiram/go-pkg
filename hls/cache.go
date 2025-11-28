package __obf_acea4ab24a824c18

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_efbbc02bedef0e2d context.Context, __obf_f1b61571a6fd5b78 string, __obf_3f431fa47d37be2d []byte) error
	Get(__obf_efbbc02bedef0e2d context.Context, __obf_f1b61571a6fd5b78 string) ([]byte, error)
}

type DirCache struct {
	__obf_f2f454838e7353a1 string
}

func NewDirCache(__obf_f2f454838e7353a1 string) *DirCache {
	return &DirCache{__obf_f2f454838e7353a1}
}

func (__obf_54aa13fbe1d0156d *DirCache) __obf_85102010d97a0c56(__obf_f1b61571a6fd5b78 string) string {
	return filepath.Join(__obf_54aa13fbe1d0156d.__obf_f2f454838e7353a1, __obf_f1b61571a6fd5b78)
}

func (__obf_54aa13fbe1d0156d *DirCache) Get(__obf_efbbc02bedef0e2d context.Context, __obf_f1b61571a6fd5b78 string) ([]byte, error) {
	__obf_944727d4e31f21db, __obf_3698cbf06506c070 := os.Open(__obf_54aa13fbe1d0156d.__obf_85102010d97a0c56(__obf_f1b61571a6fd5b78))
	if __obf_3698cbf06506c070 != nil {
		if os.IsNotExist(__obf_3698cbf06506c070) {
			return nil, nil
		}
		return nil, __obf_3698cbf06506c070
	}
	defer __obf_944727d4e31f21db.Close()
	__obf_3317e77c6d5e7689 := new(bytes.Buffer)
	if _, __obf_3698cbf06506c070 = io.Copy(__obf_3317e77c6d5e7689, __obf_944727d4e31f21db); __obf_3698cbf06506c070 != nil {
		return nil, __obf_3698cbf06506c070
	}
	return __obf_3317e77c6d5e7689.Bytes(), nil
}

func (__obf_54aa13fbe1d0156d *DirCache) Set(__obf_efbbc02bedef0e2d context.Context, __obf_f1b61571a6fd5b78 string, __obf_3f431fa47d37be2d []byte) error {
	if __obf_3698cbf06506c070 := os.MkdirAll(__obf_54aa13fbe1d0156d.__obf_f2f454838e7353a1, 0777); __obf_3698cbf06506c070 != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_54aa13fbe1d0156d.__obf_f2f454838e7353a1, __obf_3698cbf06506c070)
		return __obf_3698cbf06506c070
	}
	__obf_9481ff7e22b1408d, __obf_3698cbf06506c070 := os.CreateTemp(__obf_54aa13fbe1d0156d.__obf_f2f454838e7353a1, __obf_f1b61571a6fd5b78+".*")
	if __obf_3698cbf06506c070 != nil {
		log.Printf("Could not create cache file %v: %v", __obf_9481ff7e22b1408d, __obf_3698cbf06506c070)
		return __obf_3698cbf06506c070
	}
	if _, __obf_3698cbf06506c070 := io.Copy(__obf_9481ff7e22b1408d, bytes.NewReader(__obf_3f431fa47d37be2d)); __obf_3698cbf06506c070 != nil {
		log.Printf("Could not write cache file %v: %v", __obf_9481ff7e22b1408d, __obf_3698cbf06506c070)
		__obf_9481ff7e22b1408d.Close()
		os.Remove(__obf_9481ff7e22b1408d.Name())
		return __obf_3698cbf06506c070
	}
	if __obf_3698cbf06506c070 = __obf_9481ff7e22b1408d.Close(); __obf_3698cbf06506c070 != nil {
		log.Printf("Could not close cache file %v: %v", __obf_9481ff7e22b1408d, __obf_3698cbf06506c070)
		os.Remove(__obf_9481ff7e22b1408d.Name())
		return __obf_3698cbf06506c070
	}
	if __obf_3698cbf06506c070 = os.Rename(__obf_9481ff7e22b1408d.Name(), __obf_54aa13fbe1d0156d.__obf_85102010d97a0c56(__obf_f1b61571a6fd5b78)); __obf_3698cbf06506c070 != nil {
		log.Printf("Could not move cache file %v: %v", __obf_9481ff7e22b1408d, __obf_3698cbf06506c070)
		os.Remove(__obf_9481ff7e22b1408d.Name())
		return __obf_3698cbf06506c070
	}
	return nil
}
