package __obf_42bbcad92b7de1a8

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_a8cca90b5a49e090 context.Context, __obf_7d80d03709c3cc8f string, __obf_fd585faa0d2a77e4 []byte) error
	Get(__obf_a8cca90b5a49e090 context.Context, __obf_7d80d03709c3cc8f string) ([]byte, error)
}

type DirCache struct {
	__obf_fe2438cb530d0431 string
}

func NewDirCache(__obf_fe2438cb530d0431 string) *DirCache {
	return &DirCache{__obf_fe2438cb530d0431}
}

func (__obf_e94a3e6294010486 *DirCache) __obf_374830aeecebdbdd(__obf_7d80d03709c3cc8f string) string {
	return filepath.Join(__obf_e94a3e6294010486.__obf_fe2438cb530d0431, __obf_7d80d03709c3cc8f)
}

func (__obf_e94a3e6294010486 *DirCache) Get(__obf_a8cca90b5a49e090 context.Context, __obf_7d80d03709c3cc8f string) ([]byte, error) {
	__obf_72ab747420e4c9ee, __obf_59b2ec060775896c := os.Open(__obf_e94a3e6294010486.__obf_374830aeecebdbdd(__obf_7d80d03709c3cc8f))
	if __obf_59b2ec060775896c != nil {
		if os.IsNotExist(__obf_59b2ec060775896c) {
			return nil, nil
		}
		return nil, __obf_59b2ec060775896c
	}
	defer __obf_72ab747420e4c9ee.Close()
	__obf_dbea6d57990eeedb := new(bytes.Buffer)
	if _, __obf_59b2ec060775896c = io.Copy(__obf_dbea6d57990eeedb, __obf_72ab747420e4c9ee); __obf_59b2ec060775896c != nil {
		return nil, __obf_59b2ec060775896c
	}
	return __obf_dbea6d57990eeedb.Bytes(), nil
}

func (__obf_e94a3e6294010486 *DirCache) Set(__obf_a8cca90b5a49e090 context.Context, __obf_7d80d03709c3cc8f string, __obf_fd585faa0d2a77e4 []byte) error {
	if __obf_59b2ec060775896c := os.MkdirAll(__obf_e94a3e6294010486.__obf_fe2438cb530d0431, 0777); __obf_59b2ec060775896c != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_e94a3e6294010486.__obf_fe2438cb530d0431, __obf_59b2ec060775896c)
		return __obf_59b2ec060775896c
	}
	__obf_81853fbc1de95f77, __obf_59b2ec060775896c := os.CreateTemp(__obf_e94a3e6294010486.__obf_fe2438cb530d0431, __obf_7d80d03709c3cc8f+".*")
	if __obf_59b2ec060775896c != nil {
		log.Printf("Could not create cache file %v: %v", __obf_81853fbc1de95f77, __obf_59b2ec060775896c)
		return __obf_59b2ec060775896c
	}
	if _, __obf_59b2ec060775896c := io.Copy(__obf_81853fbc1de95f77, bytes.NewReader(__obf_fd585faa0d2a77e4)); __obf_59b2ec060775896c != nil {
		log.Printf("Could not write cache file %v: %v", __obf_81853fbc1de95f77, __obf_59b2ec060775896c)
		__obf_81853fbc1de95f77.Close()
		os.Remove(__obf_81853fbc1de95f77.Name())
		return __obf_59b2ec060775896c
	}
	if __obf_59b2ec060775896c = __obf_81853fbc1de95f77.Close(); __obf_59b2ec060775896c != nil {
		log.Printf("Could not close cache file %v: %v", __obf_81853fbc1de95f77, __obf_59b2ec060775896c)
		os.Remove(__obf_81853fbc1de95f77.Name())
		return __obf_59b2ec060775896c
	}
	if __obf_59b2ec060775896c = os.Rename(__obf_81853fbc1de95f77.Name(), __obf_e94a3e6294010486.__obf_374830aeecebdbdd(__obf_7d80d03709c3cc8f)); __obf_59b2ec060775896c != nil {
		log.Printf("Could not move cache file %v: %v", __obf_81853fbc1de95f77, __obf_59b2ec060775896c)
		os.Remove(__obf_81853fbc1de95f77.Name())
		return __obf_59b2ec060775896c
	}
	return nil
}
