package __obf_72d962f6a40bc95f

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
	Add(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any) error
	Set(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any, __obf_c6c24327f124e623 time.Duration) error
	Get(__obf_73cb8148cbf55699 string) ([]byte, error)
	Delete(__obf_939173229b00ad97 string)
	Remove(__obf_73cb8148cbf55699 string) error
	Clear() error
	Has(__obf_73cb8148cbf55699 string) bool
}

func Get[T any](__obf_7a0c72d86f8ef759 Cache, __obf_73cb8148cbf55699 string) (__obf_1479e1d874e9dca2 T, __obf_d1618b4c9656c93c error) {
	var __obf_4daaccd05acf2938 []byte
	__obf_4daaccd05acf2938, __obf_d1618b4c9656c93c = __obf_7a0c72d86f8ef759.Get(__obf_73cb8148cbf55699)
	if __obf_d1618b4c9656c93c != nil {
		return
	}

	return util.BytesToAny[T](__obf_4daaccd05acf2938)
}
