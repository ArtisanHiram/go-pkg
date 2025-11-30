package __obf_08ee9322ff6adb83

import (
	"bytes"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Cache interface {
	Set(__obf_60ee4500ce182178 context.Context, __obf_f221301302ae8617 string, __obf_144a98476b036621 []byte) error
	Get(__obf_60ee4500ce182178 context.Context, __obf_f221301302ae8617 string) ([]byte, error)
}

type DirCache struct {
	__obf_273e2c2474c36944 string
}

func NewDirCache(__obf_273e2c2474c36944 string) *DirCache {
	return &DirCache{__obf_273e2c2474c36944}
}

func (__obf_9e2d2c517a75075a *DirCache) __obf_79a83c519b116489(__obf_f221301302ae8617 string) string {
	return filepath.Join(__obf_9e2d2c517a75075a.__obf_273e2c2474c36944, __obf_f221301302ae8617)
}

func (__obf_9e2d2c517a75075a *DirCache) Get(__obf_60ee4500ce182178 context.Context, __obf_f221301302ae8617 string) ([]byte, error) {
	__obf_489cc4007c169b0b, __obf_dd2d47498586b1c8 := os.Open(__obf_9e2d2c517a75075a.__obf_79a83c519b116489(__obf_f221301302ae8617))
	if __obf_dd2d47498586b1c8 != nil {
		if os.IsNotExist(__obf_dd2d47498586b1c8) {
			return nil, nil
		}
		return nil, __obf_dd2d47498586b1c8
	}
	defer __obf_489cc4007c169b0b.Close()
	__obf_891592dc366130ba := new(bytes.Buffer)
	if _, __obf_dd2d47498586b1c8 = io.Copy(__obf_891592dc366130ba, __obf_489cc4007c169b0b); __obf_dd2d47498586b1c8 != nil {
		return nil, __obf_dd2d47498586b1c8
	}
	return __obf_891592dc366130ba.Bytes(), nil
}

func (__obf_9e2d2c517a75075a *DirCache) Set(__obf_60ee4500ce182178 context.Context, __obf_f221301302ae8617 string, __obf_144a98476b036621 []byte) error {
	if __obf_dd2d47498586b1c8 := os.MkdirAll(__obf_9e2d2c517a75075a.__obf_273e2c2474c36944, 0777); __obf_dd2d47498586b1c8 != nil {
		log.Printf("Could not create cache dir %v: %v", __obf_9e2d2c517a75075a.__obf_273e2c2474c36944, __obf_dd2d47498586b1c8)
		return __obf_dd2d47498586b1c8
	}
	__obf_d20ce5c30915c6c6, __obf_dd2d47498586b1c8 := os.CreateTemp(__obf_9e2d2c517a75075a.__obf_273e2c2474c36944, __obf_f221301302ae8617+".*")
	if __obf_dd2d47498586b1c8 != nil {
		log.Printf("Could not create cache file %v: %v", __obf_d20ce5c30915c6c6, __obf_dd2d47498586b1c8)
		return __obf_dd2d47498586b1c8
	}
	if _, __obf_dd2d47498586b1c8 := io.Copy(__obf_d20ce5c30915c6c6, bytes.NewReader(__obf_144a98476b036621)); __obf_dd2d47498586b1c8 != nil {
		log.Printf("Could not write cache file %v: %v", __obf_d20ce5c30915c6c6, __obf_dd2d47498586b1c8)
		__obf_d20ce5c30915c6c6.
			Close()
		os.Remove(__obf_d20ce5c30915c6c6.Name())
		return __obf_dd2d47498586b1c8
	}
	if __obf_dd2d47498586b1c8 = __obf_d20ce5c30915c6c6.Close(); __obf_dd2d47498586b1c8 != nil {
		log.Printf("Could not close cache file %v: %v", __obf_d20ce5c30915c6c6, __obf_dd2d47498586b1c8)
		os.Remove(__obf_d20ce5c30915c6c6.Name())
		return __obf_dd2d47498586b1c8
	}
	if __obf_dd2d47498586b1c8 = os.Rename(__obf_d20ce5c30915c6c6.Name(), __obf_9e2d2c517a75075a.__obf_79a83c519b116489(__obf_f221301302ae8617)); __obf_dd2d47498586b1c8 != nil {
		log.Printf("Could not move cache file %v: %v", __obf_d20ce5c30915c6c6, __obf_dd2d47498586b1c8)
		os.Remove(__obf_d20ce5c30915c6c6.Name())
		return __obf_dd2d47498586b1c8
	}
	return nil
}
