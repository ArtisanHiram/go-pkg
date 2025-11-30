package __obf_3b8640e918b7e3ff

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
	Add(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any) error
	Set(__obf_3405c14f70aaa4d0 string, __obf_13458218654a7f13 any, __obf_433bae995cf2f1ae time.Duration) error
	Get(__obf_3405c14f70aaa4d0 string) ([]byte, error)
	Delete(__obf_4298344c079705aa string)
	Remove(__obf_3405c14f70aaa4d0 string) error
	Clear() error
	Has(__obf_3405c14f70aaa4d0 string) bool
}

func Get[T any](__obf_c12b90239d6cd470 Cache, __obf_3405c14f70aaa4d0 string) (__obf_e9b0c648aeaf7005 T, __obf_43b169a18e2edb05 error) {
	var __obf_1c457b66502703a4 []byte
	__obf_1c457b66502703a4, __obf_43b169a18e2edb05 = __obf_c12b90239d6cd470.Get(__obf_3405c14f70aaa4d0)
	if __obf_43b169a18e2edb05 != nil {
		return
	}

	return util.BytesToAny[T](__obf_1c457b66502703a4)
}
