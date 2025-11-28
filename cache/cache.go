package __obf_8281010a3a6d2ab0

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
	Add(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any) error
	Set(__obf_805b9182f4a01a43 string, __obf_a31843a6764aecf9 any, __obf_e2e393700839b297 time.Duration) error
	Get(__obf_805b9182f4a01a43 string) ([]byte, error)
	Delete(__obf_c2739d84a72f2e49 string)
	Remove(__obf_805b9182f4a01a43 string) error
	Clear() error
	Has(__obf_805b9182f4a01a43 string) bool
}

func Get[T any](__obf_5bb7b87538df600c Cache, __obf_805b9182f4a01a43 string) (__obf_0b4c113b57caf372 T, __obf_7444330fd960f532 error) {
	var __obf_4f07d7a3752cb867 []byte
	__obf_4f07d7a3752cb867, __obf_7444330fd960f532 = __obf_5bb7b87538df600c.Get(__obf_805b9182f4a01a43)
	if __obf_7444330fd960f532 != nil {
		return
	}

	return util.BytesToAny[T](__obf_4f07d7a3752cb867)
}
