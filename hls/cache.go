package __obf_ab25ed2437cd567a

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_378e35ef817a6881 context.Context, __obf_10900f54d35c1067 string, __obf_d47816579d247f97 []byte) error
	Get(__obf_378e35ef817a6881 context.Context, __obf_10900f54d35c1067 string) ([]byte, error)
}

type DirCache struct {
	__obf_77c0216d25f1382f string
}

func NewDirCache(__obf_77c0216d25f1382f string) *DirCache {
	return &DirCache{__obf_77c0216d25f1382f}
}

func (__obf_a1d8fa5beaa156e6 *DirCache) __obf_dd89a7384a10e422(__obf_10900f54d35c1067 string) string {
	return filepath.Join(__obf_a1d8fa5beaa156e6.__obf_77c0216d25f1382f, __obf_10900f54d35c1067)
}

func (__obf_a1d8fa5beaa156e6 *DirCache) Get(__obf_378e35ef817a6881 context.Context, __obf_10900f54d35c1067 string) ([]byte, error) {
	__obf_56707f240662d9aa, __obf_9d170493b73636ca := os.Open(__obf_a1d8fa5beaa156e6.__obf_dd89a7384a10e422(__obf_10900f54d35c1067))
	if __obf_9d170493b73636ca != nil {
		if os.IsNotExist(__obf_9d170493b73636ca) {
			return nil, nil
		}
		return nil, __obf_9d170493b73636ca
	}
	defer __obf_56707f240662d9aa.Close()
	__obf_e097ea2782772156 := new(bytes.Buffer)
	if _, __obf_9d170493b73636ca = io.Copy(__obf_e097ea2782772156, __obf_56707f240662d9aa); __obf_9d170493b73636ca != nil {
		return nil, __obf_9d170493b73636ca
	}
	return __obf_e097ea2782772156.Bytes(), nil
}

func (__obf_a1d8fa5beaa156e6 *DirCache) Set(__obf_378e35ef817a6881 context.Context, __obf_10900f54d35c1067 string, __obf_d47816579d247f97 []byte) error {
	if __obf_9d170493b73636ca := os.MkdirAll(__obf_a1d8fa5beaa156e6.__obf_77c0216d25f1382f, 0777); __obf_9d170493b73636ca != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_a1d8fa5beaa156e6.__obf_77c0216d25f1382f, __obf_9d170493b73636ca)
		return __obf_9d170493b73636ca
	}
	__obf_0c3e7aaf7d329ee7, __obf_9d170493b73636ca := os.CreateTemp(__obf_a1d8fa5beaa156e6.__obf_77c0216d25f1382f, __obf_10900f54d35c1067+".*")
	if __obf_9d170493b73636ca != nil {
		log.Printf("Could not create cache file %v: %v", __obf_0c3e7aaf7d329ee7, __obf_9d170493b73636ca)
		return __obf_9d170493b73636ca
	}
	if _, __obf_9d170493b73636ca := io.Copy(__obf_0c3e7aaf7d329ee7, bytes.NewReader(__obf_d47816579d247f97)); __obf_9d170493b73636ca != nil {
		log.Printf("Could not write cache file %v: %v", __obf_0c3e7aaf7d329ee7, __obf_9d170493b73636ca)
		__obf_0c3e7aaf7d329ee7.Close()
		os.Remove(__obf_0c3e7aaf7d329ee7.Name())
		return __obf_9d170493b73636ca
	}
	if __obf_9d170493b73636ca = __obf_0c3e7aaf7d329ee7.Close(); __obf_9d170493b73636ca != nil {
		log.Printf("Could not close cache file %v: %v", __obf_0c3e7aaf7d329ee7, __obf_9d170493b73636ca)
		os.Remove(__obf_0c3e7aaf7d329ee7.Name())
		return __obf_9d170493b73636ca
	}
	if __obf_9d170493b73636ca = os.Rename(__obf_0c3e7aaf7d329ee7.Name(), __obf_a1d8fa5beaa156e6.__obf_dd89a7384a10e422(__obf_10900f54d35c1067)); __obf_9d170493b73636ca != nil {
		log.Printf("Could not move cache file %v: %v", __obf_0c3e7aaf7d329ee7, __obf_9d170493b73636ca)
		os.Remove(__obf_0c3e7aaf7d329ee7.Name())
		return __obf_9d170493b73636ca
	}
	return nil
}
