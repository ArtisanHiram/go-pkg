package __obf_90dd9b56c0f1bd65

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_9b5717753d5f5860 context.Context, __obf_592914ec2d7c3fe8 string, __obf_d5a5fb2fa6393b0d []byte) error
	Get(__obf_9b5717753d5f5860 context.Context, __obf_592914ec2d7c3fe8 string) ([]byte, error)
}

type DirCache struct {
	__obf_9b7ae180e21623e5 string
}

func NewDirCache(__obf_9b7ae180e21623e5 string) *DirCache {
	return &DirCache{__obf_9b7ae180e21623e5}
}

func (__obf_36cbab84bad25054 *DirCache) __obf_68ffd9d48b17efcc(__obf_592914ec2d7c3fe8 string) string {
	return filepath.Join(__obf_36cbab84bad25054.__obf_9b7ae180e21623e5, __obf_592914ec2d7c3fe8)
}

func (__obf_36cbab84bad25054 *DirCache) Get(__obf_9b5717753d5f5860 context.Context, __obf_592914ec2d7c3fe8 string) ([]byte, error) {
	__obf_17a4cba300dc959c, __obf_dadfdd29cd0d4fe8 := os.Open(__obf_36cbab84bad25054.__obf_68ffd9d48b17efcc(__obf_592914ec2d7c3fe8))
	if __obf_dadfdd29cd0d4fe8 != nil {
		if os.IsNotExist(__obf_dadfdd29cd0d4fe8) {
			return nil, nil
		}
		return nil, __obf_dadfdd29cd0d4fe8
	}
	defer __obf_17a4cba300dc959c.Close()
	__obf_ab9168b9655f4867 := new(bytes.Buffer)
	if _, __obf_dadfdd29cd0d4fe8 = io.Copy(__obf_ab9168b9655f4867, __obf_17a4cba300dc959c); __obf_dadfdd29cd0d4fe8 != nil {
		return nil, __obf_dadfdd29cd0d4fe8
	}
	return __obf_ab9168b9655f4867.Bytes(), nil
}

func (__obf_36cbab84bad25054 *DirCache) Set(__obf_9b5717753d5f5860 context.Context, __obf_592914ec2d7c3fe8 string, __obf_d5a5fb2fa6393b0d []byte) error {
	if __obf_dadfdd29cd0d4fe8 := os.MkdirAll(__obf_36cbab84bad25054.__obf_9b7ae180e21623e5, 0777); __obf_dadfdd29cd0d4fe8 != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_36cbab84bad25054.__obf_9b7ae180e21623e5, __obf_dadfdd29cd0d4fe8)
		return __obf_dadfdd29cd0d4fe8
	}
	__obf_a5561431aa6ea788, __obf_dadfdd29cd0d4fe8 := os.CreateTemp(__obf_36cbab84bad25054.__obf_9b7ae180e21623e5, __obf_592914ec2d7c3fe8+".*")
	if __obf_dadfdd29cd0d4fe8 != nil {
		log.Printf("Could not create cache file %v: %v", __obf_a5561431aa6ea788, __obf_dadfdd29cd0d4fe8)
		return __obf_dadfdd29cd0d4fe8
	}
	if _, __obf_dadfdd29cd0d4fe8 := io.Copy(__obf_a5561431aa6ea788, bytes.NewReader(__obf_d5a5fb2fa6393b0d)); __obf_dadfdd29cd0d4fe8 != nil {
		log.Printf("Could not write cache file %v: %v", __obf_a5561431aa6ea788, __obf_dadfdd29cd0d4fe8)
		__obf_a5561431aa6ea788.
			Close()
		os.Remove(__obf_a5561431aa6ea788.Name())
		return __obf_dadfdd29cd0d4fe8
	}
	if __obf_dadfdd29cd0d4fe8 = __obf_a5561431aa6ea788.Close(); __obf_dadfdd29cd0d4fe8 != nil {
		log.Printf("Could not close cache file %v: %v", __obf_a5561431aa6ea788, __obf_dadfdd29cd0d4fe8)
		os.Remove(__obf_a5561431aa6ea788.Name())
		return __obf_dadfdd29cd0d4fe8
	}
	if __obf_dadfdd29cd0d4fe8 = os.Rename(__obf_a5561431aa6ea788.Name(), __obf_36cbab84bad25054.__obf_68ffd9d48b17efcc(__obf_592914ec2d7c3fe8)); __obf_dadfdd29cd0d4fe8 != nil {
		log.Printf("Could not move cache file %v: %v", __obf_a5561431aa6ea788, __obf_dadfdd29cd0d4fe8)
		os.Remove(__obf_a5561431aa6ea788.Name())
		return __obf_dadfdd29cd0d4fe8
	}
	return nil
}
