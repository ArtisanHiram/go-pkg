package __obf_a05682be1c6caf18

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
	Add(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any) error
	Set(__obf_a0e56915a05e8d99 string, __obf_6030b68d8a95172f any, __obf_c8947601a564de1b time.Duration) error
	Get(__obf_a0e56915a05e8d99 string) ([]byte, error)
	Delete(__obf_715bf114d0b286fb string)
	Remove(__obf_a0e56915a05e8d99 string) error
	Clear() error
	Has(__obf_a0e56915a05e8d99 string) bool
}

func Get[T any](__obf_5f31467ce646077a Cache, __obf_a0e56915a05e8d99 string) (__obf_ee7aeee3ef89ade3 T, __obf_1086d0f8dfd43ca1 error) {
	var __obf_a28d56867747d577 []byte
	__obf_a28d56867747d577, __obf_1086d0f8dfd43ca1 = __obf_5f31467ce646077a.Get(__obf_a0e56915a05e8d99)
	if __obf_1086d0f8dfd43ca1 != nil {
		return
	}

	return util.BytesToAny[T](__obf_a28d56867747d577)
}
