package __obf_6b306490bf7221d3

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
	Add(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any) error
	Set(__obf_fa3a380c35ada5d9 string, __obf_17a4af14df8bae4f any, __obf_a15e4803e6944fc5 time.Duration) error
	Get(__obf_fa3a380c35ada5d9 string) ([]byte, error)
	Delete(__obf_7e6d8e57607bcb1d string)
	Remove(__obf_fa3a380c35ada5d9 string) error
	Clear() error
	Has(__obf_fa3a380c35ada5d9 string) bool
}

func Get[T any](__obf_8c00c4e6dd3d519c Cache, __obf_fa3a380c35ada5d9 string) (__obf_381ded77785e2fe5 T, __obf_be37eed188feee9f error) {
	var __obf_713a8d65058609de []byte
	__obf_713a8d65058609de, __obf_be37eed188feee9f = __obf_8c00c4e6dd3d519c.Get(__obf_fa3a380c35ada5d9)
	if __obf_be37eed188feee9f != nil {
		return
	}

	return util.BytesToAny[T](__obf_713a8d65058609de)
}
