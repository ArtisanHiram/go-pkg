package __obf_9e1ee87c6b054458

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
	Add(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any) error
	Set(__obf_3c13197612c6b39f string, __obf_7d0a67e130b6e9b4 any, __obf_5a47a67359bdcb92 time.Duration) error
	Get(__obf_3c13197612c6b39f string) ([]byte, error)
	Delete(__obf_14a292fb641d1b3d string)
	Remove(__obf_3c13197612c6b39f string) error
	Clear() error
	Has(__obf_3c13197612c6b39f string) bool
}

func Get[T any](__obf_e570919ee0a46ecd Cache, __obf_3c13197612c6b39f string) (__obf_9428e3da1f957acd T, __obf_4f22c5826c5e34df error) {
	var __obf_63af5e67abca5f59 []byte
	__obf_63af5e67abca5f59, __obf_4f22c5826c5e34df = __obf_e570919ee0a46ecd.Get(__obf_3c13197612c6b39f)
	if __obf_4f22c5826c5e34df != nil {
		return
	}

	return util.BytesToAny[T](__obf_63af5e67abca5f59)
}
