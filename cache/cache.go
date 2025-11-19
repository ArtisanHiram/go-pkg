package __obf_62df4de078c8d208

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
	Add(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any) error
	Set(__obf_4aecf3c737bbe5e8 string, __obf_676d75836c094b83 any, __obf_24d92102aed9ce02 time.Duration) error
	Get(__obf_4aecf3c737bbe5e8 string) ([]byte, error)
	Delete(__obf_0c00b35780000392 string)
	Remove(__obf_4aecf3c737bbe5e8 string) error
	Clear() error
	Has(__obf_4aecf3c737bbe5e8 string) bool
}

func Get[T any](__obf_eb6273cfd663c098 Cache, __obf_4aecf3c737bbe5e8 string) (__obf_dbf9fa51e995bfee T, __obf_68f2d1855583b558 error) {
	var __obf_0eefc8c70e7b9364 []byte
	__obf_0eefc8c70e7b9364, __obf_68f2d1855583b558 = __obf_eb6273cfd663c098.Get(__obf_4aecf3c737bbe5e8)
	if __obf_68f2d1855583b558 != nil {
		return
	}

	return util.BytesToAny[T](__obf_0eefc8c70e7b9364)
}
