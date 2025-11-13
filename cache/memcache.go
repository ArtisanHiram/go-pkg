package __obf_f4d8558b35981340

import (
	"time"

	util "github.com/ArtisanHiram/go-pkg/util"
	"github.com/bradfitz/gomemcache/memcache"
)

// Cache Memcache adapter.
type Memcache struct {
	__obf_3f4f5b63b71880fe *memcache.Client
}

func (__obf_291d8f01436de6d5 *Memcache) Has(__obf_6208471da03258dd string) bool {
	_, __obf_2d5e4b42e300bedd := __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Get(__obf_6208471da03258dd)
	return __obf_2d5e4b42e300bedd == nil
}

// Delete implements Cache.
func (__obf_291d8f01436de6d5 *Memcache) Delete(__obf_1261d18b5941c6bd string) {
	panic("unimplemented")
}

// NewMemCache creates a new memcache adapter.
func NewMemCache(__obf_1555516b7fecc02a []string) Cache {
	return &Memcache{memcache.New(__obf_1555516b7fecc02a...)}
}

// Get get value from memcache.
func (__obf_291d8f01436de6d5 *Memcache) Get(__obf_6208471da03258dd string) (__obf_9621ffad1aef0ed3 []byte, __obf_2d5e4b42e300bedd error) {
	if __obf_a7d9bc40b5640158, __obf_2d5e4b42e300bedd := __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Get(__obf_6208471da03258dd); __obf_2d5e4b42e300bedd == nil {

		return __obf_a7d9bc40b5640158.Value, nil
	} else {
		return nil, __obf_2d5e4b42e300bedd
	}
}

func (__obf_291d8f01436de6d5 *Memcache) Set(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any, __obf_1e67295112a2a62d time.Duration) error {
	__obf_a7d9bc40b5640158 := memcache.Item{Key: __obf_6208471da03258dd, Expiration: int32(__obf_1e67295112a2a62d / time.Second)}
	var __obf_2d5e4b42e300bedd error
	__obf_a7d9bc40b5640158.Value, __obf_2d5e4b42e300bedd = util.AnyToBytes(__obf_20a54be727e70f3c)
	if __obf_2d5e4b42e300bedd != nil {
		return __obf_2d5e4b42e300bedd
	}
	return __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Set(&__obf_a7d9bc40b5640158)
}

func (__obf_291d8f01436de6d5 *Memcache) Add(__obf_6208471da03258dd string, __obf_20a54be727e70f3c any) error {
	__obf_a7d9bc40b5640158 := memcache.Item{Key: __obf_6208471da03258dd}

	var __obf_2d5e4b42e300bedd error
	__obf_a7d9bc40b5640158.Value, __obf_2d5e4b42e300bedd = util.AnyToBytes(__obf_20a54be727e70f3c)
	if __obf_2d5e4b42e300bedd != nil {
		return __obf_2d5e4b42e300bedd
	}

	return __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Set(&__obf_a7d9bc40b5640158)
}

// ClearAll clears all cache in memcache.
func (__obf_291d8f01436de6d5 *Memcache) Clear() error {
	return __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.FlushAll()
}

// Incr increases counter.
func (__obf_291d8f01436de6d5 *Memcache) Incr(__obf_6208471da03258dd string) error {
	_, __obf_2d5e4b42e300bedd := __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Increment(__obf_6208471da03258dd, 1)
	return __obf_2d5e4b42e300bedd
}

// Decr decreases counter.
func (__obf_291d8f01436de6d5 *Memcache) Decr(__obf_6208471da03258dd string) error {
	_, __obf_2d5e4b42e300bedd := __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Decrement(__obf_6208471da03258dd, 1)
	return __obf_2d5e4b42e300bedd
}

func (__obf_291d8f01436de6d5 *Memcache) Renew(__obf_6208471da03258dd string, __obf_1e67295112a2a62d time.Duration) error {
	if __obf_a7d9bc40b5640158, __obf_2d5e4b42e300bedd := __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Get(__obf_6208471da03258dd); __obf_2d5e4b42e300bedd == nil {
		return __obf_291d8f01436de6d5.Set(__obf_6208471da03258dd, __obf_a7d9bc40b5640158, __obf_1e67295112a2a62d)
	} else {
		return __obf_2d5e4b42e300bedd
	}
}

func (__obf_291d8f01436de6d5 *Memcache) IsExist(__obf_6208471da03258dd string) bool {
	_, __obf_2d5e4b42e300bedd := __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Get(__obf_6208471da03258dd)
	return __obf_2d5e4b42e300bedd == nil
}

func (__obf_291d8f01436de6d5 *Memcache) Remove(__obf_6208471da03258dd string) error {
	return __obf_291d8f01436de6d5.__obf_3f4f5b63b71880fe.Delete(__obf_6208471da03258dd)
}
