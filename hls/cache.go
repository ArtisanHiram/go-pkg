package __obf_5028a0a829ddcdab

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_19b79b6a380e8507 context.Context, __obf_fd503f6151c05aa1 string, __obf_24fb9b8b0dd82ea1 []byte) error
	Get(__obf_19b79b6a380e8507 context.Context, __obf_fd503f6151c05aa1 string) ([]byte, error)
}

type DirCache struct {
	__obf_4b28c4ed43afcd28 string
}

func NewDirCache(__obf_4b28c4ed43afcd28 string) *DirCache {
	return &DirCache{__obf_4b28c4ed43afcd28}
}

func (__obf_16db58eb289ba0f3 *DirCache) __obf_a4f9cd8e180e081d(__obf_fd503f6151c05aa1 string) string {
	return filepath.Join(__obf_16db58eb289ba0f3.__obf_4b28c4ed43afcd28, __obf_fd503f6151c05aa1)
}

func (__obf_16db58eb289ba0f3 *DirCache) Get(__obf_19b79b6a380e8507 context.Context, __obf_fd503f6151c05aa1 string) ([]byte, error) {
	__obf_8ddbcc776441ebf5, __obf_97be5ef7ba0b1a8f := os.Open(__obf_16db58eb289ba0f3.__obf_a4f9cd8e180e081d(__obf_fd503f6151c05aa1))
	if __obf_97be5ef7ba0b1a8f != nil {
		if os.IsNotExist(__obf_97be5ef7ba0b1a8f) {
			return nil, nil
		}
		return nil, __obf_97be5ef7ba0b1a8f
	}
	defer __obf_8ddbcc776441ebf5.Close()
	__obf_a128e98fcbfe8702 := new(bytes.Buffer)
	if _, __obf_97be5ef7ba0b1a8f = io.Copy(__obf_a128e98fcbfe8702, __obf_8ddbcc776441ebf5); __obf_97be5ef7ba0b1a8f != nil {
		return nil, __obf_97be5ef7ba0b1a8f
	}
	return __obf_a128e98fcbfe8702.Bytes(), nil
}

func (__obf_16db58eb289ba0f3 *DirCache) Set(__obf_19b79b6a380e8507 context.Context, __obf_fd503f6151c05aa1 string, __obf_24fb9b8b0dd82ea1 []byte) error {
	if __obf_97be5ef7ba0b1a8f := os.MkdirAll(__obf_16db58eb289ba0f3.__obf_4b28c4ed43afcd28, 0777); __obf_97be5ef7ba0b1a8f != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_16db58eb289ba0f3.__obf_4b28c4ed43afcd28, __obf_97be5ef7ba0b1a8f)
		return __obf_97be5ef7ba0b1a8f
	}
	__obf_9700f69ab5c5121c, __obf_97be5ef7ba0b1a8f := os.CreateTemp(__obf_16db58eb289ba0f3.__obf_4b28c4ed43afcd28, __obf_fd503f6151c05aa1+".*")
	if __obf_97be5ef7ba0b1a8f != nil {
		log.Printf("Could not create cache file %v: %v", __obf_9700f69ab5c5121c, __obf_97be5ef7ba0b1a8f)
		return __obf_97be5ef7ba0b1a8f
	}
	if _, __obf_97be5ef7ba0b1a8f := io.Copy(__obf_9700f69ab5c5121c, bytes.NewReader(__obf_24fb9b8b0dd82ea1)); __obf_97be5ef7ba0b1a8f != nil {
		log.Printf("Could not write cache file %v: %v", __obf_9700f69ab5c5121c, __obf_97be5ef7ba0b1a8f)
		__obf_9700f69ab5c5121c.Close()
		os.Remove(__obf_9700f69ab5c5121c.Name())
		return __obf_97be5ef7ba0b1a8f
	}
	if __obf_97be5ef7ba0b1a8f = __obf_9700f69ab5c5121c.Close(); __obf_97be5ef7ba0b1a8f != nil {
		log.Printf("Could not close cache file %v: %v", __obf_9700f69ab5c5121c, __obf_97be5ef7ba0b1a8f)
		os.Remove(__obf_9700f69ab5c5121c.Name())
		return __obf_97be5ef7ba0b1a8f
	}
	if __obf_97be5ef7ba0b1a8f = os.Rename(__obf_9700f69ab5c5121c.Name(), __obf_16db58eb289ba0f3.__obf_a4f9cd8e180e081d(__obf_fd503f6151c05aa1)); __obf_97be5ef7ba0b1a8f != nil {
		log.Printf("Could not move cache file %v: %v", __obf_9700f69ab5c5121c, __obf_97be5ef7ba0b1a8f)
		os.Remove(__obf_9700f69ab5c5121c.Name())
		return __obf_97be5ef7ba0b1a8f
	}
	return nil
}
