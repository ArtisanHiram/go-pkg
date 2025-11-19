package __obf_dc577bef9ac5c6b8

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_5d1af4120f069bb9 context.Context, __obf_85503b8843d1609c string, __obf_866a0608e4aff093 []byte) error
	Get(__obf_5d1af4120f069bb9 context.Context, __obf_85503b8843d1609c string) ([]byte, error)
}

type DirCache struct {
	__obf_ef682768d47e5b4e string
}

func NewDirCache(__obf_ef682768d47e5b4e string) *DirCache {
	return &DirCache{__obf_ef682768d47e5b4e}
}

func (__obf_da693a20c0580a55 *DirCache) __obf_0434299a335e7440(__obf_85503b8843d1609c string) string {
	return filepath.Join(__obf_da693a20c0580a55.__obf_ef682768d47e5b4e, __obf_85503b8843d1609c)
}

func (__obf_da693a20c0580a55 *DirCache) Get(__obf_5d1af4120f069bb9 context.Context, __obf_85503b8843d1609c string) ([]byte, error) {
	__obf_cf3312d5b91d1c43, __obf_14a3b54d08bd697d := os.Open(__obf_da693a20c0580a55.__obf_0434299a335e7440(__obf_85503b8843d1609c))
	if __obf_14a3b54d08bd697d != nil {
		if os.IsNotExist(__obf_14a3b54d08bd697d) {
			return nil, nil
		}
		return nil, __obf_14a3b54d08bd697d
	}
	defer __obf_cf3312d5b91d1c43.Close()
	__obf_f5cb3c85b79131f3 := new(bytes.Buffer)
	if _, __obf_14a3b54d08bd697d = io.Copy(__obf_f5cb3c85b79131f3, __obf_cf3312d5b91d1c43); __obf_14a3b54d08bd697d != nil {
		return nil, __obf_14a3b54d08bd697d
	}
	return __obf_f5cb3c85b79131f3.Bytes(), nil
}

func (__obf_da693a20c0580a55 *DirCache) Set(__obf_5d1af4120f069bb9 context.Context, __obf_85503b8843d1609c string, __obf_866a0608e4aff093 []byte) error {
	if __obf_14a3b54d08bd697d := os.MkdirAll(__obf_da693a20c0580a55.__obf_ef682768d47e5b4e, 0777); __obf_14a3b54d08bd697d != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_da693a20c0580a55.__obf_ef682768d47e5b4e, __obf_14a3b54d08bd697d)
		return __obf_14a3b54d08bd697d
	}
	__obf_d8ab4b546a1d7288, __obf_14a3b54d08bd697d := os.CreateTemp(__obf_da693a20c0580a55.__obf_ef682768d47e5b4e, __obf_85503b8843d1609c+".*")
	if __obf_14a3b54d08bd697d != nil {
		log.Printf("Could not create cache file %v: %v", __obf_d8ab4b546a1d7288, __obf_14a3b54d08bd697d)
		return __obf_14a3b54d08bd697d
	}
	if _, __obf_14a3b54d08bd697d := io.Copy(__obf_d8ab4b546a1d7288, bytes.NewReader(__obf_866a0608e4aff093)); __obf_14a3b54d08bd697d != nil {
		log.Printf("Could not write cache file %v: %v", __obf_d8ab4b546a1d7288, __obf_14a3b54d08bd697d)
		__obf_d8ab4b546a1d7288.Close()
		os.Remove(__obf_d8ab4b546a1d7288.Name())
		return __obf_14a3b54d08bd697d
	}
	if __obf_14a3b54d08bd697d = __obf_d8ab4b546a1d7288.Close(); __obf_14a3b54d08bd697d != nil {
		log.Printf("Could not close cache file %v: %v", __obf_d8ab4b546a1d7288, __obf_14a3b54d08bd697d)
		os.Remove(__obf_d8ab4b546a1d7288.Name())
		return __obf_14a3b54d08bd697d
	}
	if __obf_14a3b54d08bd697d = os.Rename(__obf_d8ab4b546a1d7288.Name(), __obf_da693a20c0580a55.__obf_0434299a335e7440(__obf_85503b8843d1609c)); __obf_14a3b54d08bd697d != nil {
		log.Printf("Could not move cache file %v: %v", __obf_d8ab4b546a1d7288, __obf_14a3b54d08bd697d)
		os.Remove(__obf_d8ab4b546a1d7288.Name())
		return __obf_14a3b54d08bd697d
	}
	return nil
}
