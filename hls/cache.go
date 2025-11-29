package __obf_5441fcd9a319cf59

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_83fad17274e83f8e context.Context, __obf_d7f29c2f26668a2e string, __obf_a99577ab3e6d5819 []byte) error
	Get(__obf_83fad17274e83f8e context.Context, __obf_d7f29c2f26668a2e string) ([]byte, error)
}

type DirCache struct {
	__obf_75bf03e05fea1b50 string
}

func NewDirCache(__obf_75bf03e05fea1b50 string) *DirCache {
	return &DirCache{__obf_75bf03e05fea1b50}
}

func (__obf_6847fa8197aa6033 *DirCache) __obf_f1f7fe554c023809(__obf_d7f29c2f26668a2e string) string {
	return filepath.Join(__obf_6847fa8197aa6033.__obf_75bf03e05fea1b50, __obf_d7f29c2f26668a2e)
}

func (__obf_6847fa8197aa6033 *DirCache) Get(__obf_83fad17274e83f8e context.Context, __obf_d7f29c2f26668a2e string) ([]byte, error) {
	__obf_617a682795b315ad, __obf_e8125dea727cd4c9 := os.Open(__obf_6847fa8197aa6033.__obf_f1f7fe554c023809(__obf_d7f29c2f26668a2e))
	if __obf_e8125dea727cd4c9 != nil {
		if os.IsNotExist(__obf_e8125dea727cd4c9) {
			return nil, nil
		}
		return nil, __obf_e8125dea727cd4c9
	}
	defer __obf_617a682795b315ad.Close()
	__obf_6408a455935ea061 := new(bytes.Buffer)
	if _, __obf_e8125dea727cd4c9 = io.Copy(__obf_6408a455935ea061, __obf_617a682795b315ad); __obf_e8125dea727cd4c9 != nil {
		return nil, __obf_e8125dea727cd4c9
	}
	return __obf_6408a455935ea061.Bytes(), nil
}

func (__obf_6847fa8197aa6033 *DirCache) Set(__obf_83fad17274e83f8e context.Context, __obf_d7f29c2f26668a2e string, __obf_a99577ab3e6d5819 []byte) error {
	if __obf_e8125dea727cd4c9 := os.MkdirAll(__obf_6847fa8197aa6033.__obf_75bf03e05fea1b50, 0777); __obf_e8125dea727cd4c9 != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_6847fa8197aa6033.__obf_75bf03e05fea1b50, __obf_e8125dea727cd4c9)
		return __obf_e8125dea727cd4c9
	}
	__obf_aea876799c98c541, __obf_e8125dea727cd4c9 := os.CreateTemp(__obf_6847fa8197aa6033.__obf_75bf03e05fea1b50, __obf_d7f29c2f26668a2e+".*")
	if __obf_e8125dea727cd4c9 != nil {
		log.Printf("Could not create cache file %v: %v", __obf_aea876799c98c541, __obf_e8125dea727cd4c9)
		return __obf_e8125dea727cd4c9
	}
	if _, __obf_e8125dea727cd4c9 := io.Copy(__obf_aea876799c98c541, bytes.NewReader(__obf_a99577ab3e6d5819)); __obf_e8125dea727cd4c9 != nil {
		log.Printf("Could not write cache file %v: %v", __obf_aea876799c98c541, __obf_e8125dea727cd4c9)
		__obf_aea876799c98c541.
			Close()
		os.Remove(__obf_aea876799c98c541.Name())
		return __obf_e8125dea727cd4c9
	}
	if __obf_e8125dea727cd4c9 = __obf_aea876799c98c541.Close(); __obf_e8125dea727cd4c9 != nil {
		log.Printf("Could not close cache file %v: %v", __obf_aea876799c98c541, __obf_e8125dea727cd4c9)
		os.Remove(__obf_aea876799c98c541.Name())
		return __obf_e8125dea727cd4c9
	}
	if __obf_e8125dea727cd4c9 = os.Rename(__obf_aea876799c98c541.Name(), __obf_6847fa8197aa6033.__obf_f1f7fe554c023809(__obf_d7f29c2f26668a2e)); __obf_e8125dea727cd4c9 != nil {
		log.Printf("Could not move cache file %v: %v", __obf_aea876799c98c541, __obf_e8125dea727cd4c9)
		os.Remove(__obf_aea876799c98c541.Name())
		return __obf_e8125dea727cd4c9
	}
	return nil
}
