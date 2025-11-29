package __obf_6fd4fe34e3f784df

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
	Add(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any) error
	Set(__obf_ca986a38d1f8fbbb string, __obf_e1d0258259217620 any, __obf_bfbeaaea582ee60a time.Duration) error
	Get(__obf_ca986a38d1f8fbbb string) ([]byte, error)
	Delete(__obf_c53a458d6b02aeb9 string)
	Remove(__obf_ca986a38d1f8fbbb string) error
	Clear() error
	Has(__obf_ca986a38d1f8fbbb string) bool
}

func Get[T any](__obf_ce39efbc9e52b6ba Cache, __obf_ca986a38d1f8fbbb string) (__obf_1b74b25b0882bf4e T, __obf_0f22edc65e150c28 error) {
	var __obf_f20a3eb6a5a1b0dc []byte
	__obf_f20a3eb6a5a1b0dc, __obf_0f22edc65e150c28 = __obf_ce39efbc9e52b6ba.Get(__obf_ca986a38d1f8fbbb)
	if __obf_0f22edc65e150c28 != nil {
		return
	}

	return util.BytesToAny[T](__obf_f20a3eb6a5a1b0dc)
}
