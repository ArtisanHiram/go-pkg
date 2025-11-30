package __obf_72d962f6a40bc95f

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_c12bbf95a2fd00bc *memcache.Client
}

func (__obf_8636e722aa59e966 *Memcache) Has(__obf_73cb8148cbf55699 string) bool {
	_, __obf_2498adaec5f4c8f1 := __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Get(__obf_73cb8148cbf55699)
	return __obf_2498adaec5f4c8f1 == nil
}

// Delete implements Cache.
func (__obf_8636e722aa59e966 *Memcache) Delete(__obf_939173229b00ad97 string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_471ed360c284fc6c []string) Cache {
	return &Memcache{memcache.New(__obf_471ed360c284fc6c...)}
}

// Get get value from memcache.
func (__obf_8636e722aa59e966 *Memcache) Get(__obf_73cb8148cbf55699 string) (__obf_ff9943ce01534ee6 []byte, __obf_2498adaec5f4c8f1 error) {
	if __obf_cdc238f327ed4e5c, __obf_2498adaec5f4c8f1 := __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Get(__obf_73cb8148cbf55699); __obf_2498adaec5f4c8f1 == nil {

		return __obf_cdc238f327ed4e5c.Value, nil
	} else {
		return nil, __obf_2498adaec5f4c8f1
	}
}

func (__obf_8636e722aa59e966 *Memcache) Set(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any, __obf_c6c24327f124e623 time.Duration) error {
	__obf_cdc238f327ed4e5c := memcache.Item{Key: __obf_73cb8148cbf55699, Expiration: int32(__obf_c6c24327f124e623 / time.Second)}
	var __obf_2498adaec5f4c8f1 error
	__obf_cdc238f327ed4e5c.
		Value, __obf_2498adaec5f4c8f1 = util.AnyToBytes(__obf_1dbdf9c3e13789b0)
	if __obf_2498adaec5f4c8f1 != nil {
		return __obf_2498adaec5f4c8f1
	}
	return __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Set(&__obf_cdc238f327ed4e5c)
}

func (__obf_8636e722aa59e966 *Memcache) Add(__obf_73cb8148cbf55699 string, __obf_1dbdf9c3e13789b0 any) error {
	__obf_cdc238f327ed4e5c := memcache.Item{Key: __obf_73cb8148cbf55699}

	var __obf_2498adaec5f4c8f1 error
	__obf_cdc238f327ed4e5c.
		Value, __obf_2498adaec5f4c8f1 = util.AnyToBytes(__obf_1dbdf9c3e13789b0)
	if __obf_2498adaec5f4c8f1 != nil {
		return __obf_2498adaec5f4c8f1
	}

	return __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Set(&__obf_cdc238f327ed4e5c)
}

// ClearAll clears all cache in memcache.
func (__obf_8636e722aa59e966 *Memcache) Clear() error {
	return __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.FlushAll()
}

// Incr increases counter.
func (__obf_8636e722aa59e966 *Memcache) Incr(__obf_73cb8148cbf55699 string) error {
	_, __obf_2498adaec5f4c8f1 := __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Increment(__obf_73cb8148cbf55699, 1)
	return __obf_2498adaec5f4c8f1
}

// Decr decreases counter.
func (__obf_8636e722aa59e966 *Memcache) Decr(__obf_73cb8148cbf55699 string) error {
	_, __obf_2498adaec5f4c8f1 := __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Decrement(__obf_73cb8148cbf55699, 1)
	return __obf_2498adaec5f4c8f1
}

func (__obf_8636e722aa59e966 *Memcache) Renew(__obf_73cb8148cbf55699 string, __obf_c6c24327f124e623 time.Duration) error {
	if __obf_cdc238f327ed4e5c, __obf_2498adaec5f4c8f1 := __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Get(__obf_73cb8148cbf55699); __obf_2498adaec5f4c8f1 == nil {
		return __obf_8636e722aa59e966.Set(__obf_73cb8148cbf55699, __obf_cdc238f327ed4e5c, __obf_c6c24327f124e623)
	} else {
		return __obf_2498adaec5f4c8f1
	}
}

func (__obf_8636e722aa59e966 *Memcache) IsExist(__obf_73cb8148cbf55699 string) bool {
	_, __obf_2498adaec5f4c8f1 := __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Get(__obf_73cb8148cbf55699)
	return __obf_2498adaec5f4c8f1 == nil
}

func (__obf_8636e722aa59e966 *Memcache) Remove(__obf_73cb8148cbf55699 string) error {
	return __obf_8636e722aa59e966.__obf_c12bbf95a2fd00bc.Delete(__obf_73cb8148cbf55699)
}
