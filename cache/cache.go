package __obf_79e7502d8563d901

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
	Add(__obf_50994613b7653a88 string, __obf_e043381e0771feca any) error
	Set(__obf_50994613b7653a88 string, __obf_e043381e0771feca any, __obf_4ff6d8752a4fea92 time.Duration) error
	Get(__obf_50994613b7653a88 string) ([]byte, error)
	Delete(__obf_bac226386ec983d7 string)
	Remove(__obf_50994613b7653a88 string) error
	Clear() error
	Has(__obf_50994613b7653a88 string) bool
}

func Get[T any](__obf_38cd74f1c4e39ef3 Cache, __obf_50994613b7653a88 string) (__obf_83c2f84d0ae1360e T, __obf_8429338405294247 error) {
	var __obf_e859322e84749c37 []byte
	__obf_e859322e84749c37, __obf_8429338405294247 = __obf_38cd74f1c4e39ef3.Get(__obf_50994613b7653a88)
	if __obf_8429338405294247 != nil {
		return
	}

	return util.BytesToAny[T](__obf_e859322e84749c37)
}
