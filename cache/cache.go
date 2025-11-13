package __obf_f4d8558b35981340

import (
	"errors"
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
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
	Add(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any) error
	Set(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any, __obf_1e67295112a2a62d time.Duration) error
	Get(__obf_6208471da03258dd string) ([]byte, error)
	Delete(__obf_1261d18b5941c6bd string)
	Remove(__obf_6208471da03258dd string) error
	Clear() error
	Has(__obf_6208471da03258dd string) bool
}

func Get[T any](__obf_16acb12ec004bad4 Cache, __obf_6208471da03258dd string) (__obf_75d2c979de8603cc T, __obf_06e4a4e6e9269bd7 error) {
	var __obf_aa4cb039e7b11552 []byte
	__obf_aa4cb039e7b11552, __obf_06e4a4e6e9269bd7 = __obf_16acb12ec004bad4.Get(__obf_6208471da03258dd)
	if __obf_06e4a4e6e9269bd7 != nil {
		return
	}

	return util.BytesToAny[T](__obf_aa4cb039e7b11552)
}
