package __obf_6ff082bd539c7df0

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_5254a2186cef3f5f context.Context, __obf_5efc61a6b0304a66 string, __obf_872f3f3c6f81cac3 []byte) error
	Get(__obf_5254a2186cef3f5f context.Context, __obf_5efc61a6b0304a66 string) ([]byte, error)
}

type DirCache struct {
	__obf_7e66a22be900b14f string
}

func NewDirCache(__obf_7e66a22be900b14f string) *DirCache {
	return &DirCache{__obf_7e66a22be900b14f}
}

func (__obf_9ef0bd5feccf129a *DirCache) __obf_a4870b06a752c911(__obf_5efc61a6b0304a66 string) string {
	return filepath.Join(__obf_9ef0bd5feccf129a.__obf_7e66a22be900b14f, __obf_5efc61a6b0304a66)
}

func (__obf_9ef0bd5feccf129a *DirCache) Get(__obf_5254a2186cef3f5f context.Context, __obf_5efc61a6b0304a66 string) ([]byte, error) {
	__obf_f6e6713da7daf44f, __obf_ccdf867ff6e9ddb3 := os.Open(__obf_9ef0bd5feccf129a.__obf_a4870b06a752c911(__obf_5efc61a6b0304a66))
	if __obf_ccdf867ff6e9ddb3 != nil {
		if os.IsNotExist(__obf_ccdf867ff6e9ddb3) {
			return nil, nil
		}
		return nil, __obf_ccdf867ff6e9ddb3
	}
	defer __obf_f6e6713da7daf44f.Close()
	__obf_7140f9b71e52894f := new(bytes.Buffer)
	if _, __obf_ccdf867ff6e9ddb3 = io.Copy(__obf_7140f9b71e52894f, __obf_f6e6713da7daf44f); __obf_ccdf867ff6e9ddb3 != nil {
		return nil, __obf_ccdf867ff6e9ddb3
	}
	return __obf_7140f9b71e52894f.Bytes(), nil
}

func (__obf_9ef0bd5feccf129a *DirCache) Set(__obf_5254a2186cef3f5f context.Context, __obf_5efc61a6b0304a66 string, __obf_872f3f3c6f81cac3 []byte) error {
	if __obf_ccdf867ff6e9ddb3 := os.MkdirAll(__obf_9ef0bd5feccf129a.__obf_7e66a22be900b14f, 0777); __obf_ccdf867ff6e9ddb3 != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_9ef0bd5feccf129a.__obf_7e66a22be900b14f, __obf_ccdf867ff6e9ddb3)
		return __obf_ccdf867ff6e9ddb3
	}
	__obf_a1303d5511b53b60, __obf_ccdf867ff6e9ddb3 := os.CreateTemp(__obf_9ef0bd5feccf129a.__obf_7e66a22be900b14f, __obf_5efc61a6b0304a66+".*")
	if __obf_ccdf867ff6e9ddb3 != nil {
		log.Printf("Could not create cache file %v: %v", __obf_a1303d5511b53b60, __obf_ccdf867ff6e9ddb3)
		return __obf_ccdf867ff6e9ddb3
	}
	if _, __obf_ccdf867ff6e9ddb3 := io.Copy(__obf_a1303d5511b53b60, bytes.NewReader(__obf_872f3f3c6f81cac3)); __obf_ccdf867ff6e9ddb3 != nil {
		log.Printf("Could not write cache file %v: %v", __obf_a1303d5511b53b60, __obf_ccdf867ff6e9ddb3)
		__obf_a1303d5511b53b60.
			Close()
		os.Remove(__obf_a1303d5511b53b60.Name())
		return __obf_ccdf867ff6e9ddb3
	}
	if __obf_ccdf867ff6e9ddb3 = __obf_a1303d5511b53b60.Close(); __obf_ccdf867ff6e9ddb3 != nil {
		log.Printf("Could not close cache file %v: %v", __obf_a1303d5511b53b60, __obf_ccdf867ff6e9ddb3)
		os.Remove(__obf_a1303d5511b53b60.Name())
		return __obf_ccdf867ff6e9ddb3
	}
	if __obf_ccdf867ff6e9ddb3 = os.Rename(__obf_a1303d5511b53b60.Name(), __obf_9ef0bd5feccf129a.__obf_a4870b06a752c911(__obf_5efc61a6b0304a66)); __obf_ccdf867ff6e9ddb3 != nil {
		log.Printf("Could not move cache file %v: %v", __obf_a1303d5511b53b60, __obf_ccdf867ff6e9ddb3)
		os.Remove(__obf_a1303d5511b53b60.Name())
		return __obf_ccdf867ff6e9ddb3
	}
	return nil
}
