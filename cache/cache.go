package __obf_32d927a1e06b7e71

import (
	"errors"
	util "github.com/ArtisanHiram/go-pkg/util"
	"time"
)

var (
	ErrCacheEmpty   = errors.New("cache is empty")
	ErrKeyNotExists = errors.New("the key '%s' is not exists")
)

type Expirable interface {
	CreatedAt() int64
	ExpiresIn() int
}

// type Cache interface {
// 	Get(key string) any
// 	Set(key string, val any, timeout time.Duration) error
// 	IsExist(key string) bool
// 	Delete(key string) error
// }

type Cache interface {
	Add(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any) error
	Set(__obf_13113b328a6972ad string, __obf_a967d4ebf1531f4c any, __obf_481cbd2caaded2ed time.Duration) error
	Get(__obf_13113b328a6972ad string) ([]byte, error)
	Delete(__obf_e5c9d9e65a3fa384 string)
	Remove(__obf_13113b328a6972ad string) error
	Clear() error
	Has(__obf_13113b328a6972ad string) bool
}

func Get[T any](__obf_747f01f8ed45c5b0 Cache, __obf_13113b328a6972ad string) (__obf_220d26a30e454710 T, __obf_b9653b0201c59d0c error) {
	var __obf_cfabf75f0b35f309 []byte
	__obf_cfabf75f0b35f309, __obf_b9653b0201c59d0c = __obf_747f01f8ed45c5b0.Get(__obf_13113b328a6972ad)
	if __obf_b9653b0201c59d0c != nil {
		return
	}

	return util.BytesToAny[T](__obf_cfabf75f0b35f309)
}
