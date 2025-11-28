package __obf_2f5c14e012cec42e

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
	Add(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any) error
	Set(__obf_f0b3ebeba048b5e4 string, __obf_aee38dd09d94cdd5 any, __obf_6459304da6a62a40 time.Duration) error
	Get(__obf_f0b3ebeba048b5e4 string) ([]byte, error)
	Delete(__obf_68047b361b381cb2 string)
	Remove(__obf_f0b3ebeba048b5e4 string) error
	Clear() error
	Has(__obf_f0b3ebeba048b5e4 string) bool
}

func Get[T any](__obf_9ee8e7cabd1ca66a Cache, __obf_f0b3ebeba048b5e4 string) (__obf_77863c2b4b36b1f6 T, __obf_5bdf35895068ad96 error) {
	var __obf_e7b5d7bde45f2010 []byte
	__obf_e7b5d7bde45f2010, __obf_5bdf35895068ad96 = __obf_9ee8e7cabd1ca66a.Get(__obf_f0b3ebeba048b5e4)
	if __obf_5bdf35895068ad96 != nil {
		return
	}

	return util.BytesToAny[T](__obf_e7b5d7bde45f2010)
}
